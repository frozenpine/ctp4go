package parser

import (
	"bytes"
	"fmt"

	"github.com/go-clang/clang-v15/clang"
)

type Param struct {
	// 参数名
	Name string
	// 参数类型
	Type string
	// 是否指针
	IsPointer bool
	// 是否const修饰
	IsConst bool
	// 是否结构体
	IsStruct bool
	// 是否不定长数组
	IsArray bool
	// 不定长数组的长度参数名
	ArrSizeName string

	Fn *ClsMethod
}

func (p Param) String() string {
	buff := bytes.NewBufferString("")
	if p.IsConst {
		buff.WriteString("const ")
	}

	buff.WriteString(p.Type)
	if p.IsPointer {
		buff.WriteString("*")
	}
	if p.IsArray {
		fmt.Fprintf(buff, " [%s]", p.ArrSizeName)
	}
	if p.Name != "" {
		fmt.Fprintf(buff, " %s", p.Name)
	}

	return buff.String()
}

type ClsMethod struct {
	baseDefine

	// Linux的编译装饰名
	MangledName string
	// Windows的编译装饰名
	WinMangledName string

	// 入参列表
	Params []*Param
	// 返回参数
	Rtn *Param

	// 是否虚函数
	IsVirtual bool
	// 是否静态函数
	IsStatic bool

	Cls *ClassDefine
}

func (fn *ClsMethod) ParseParam(cursor *clang.Cursor) {
	methodType := cursor.Type()
	rtnType := methodType.ResultType().CanonicalType()

	if rtnType.Spelling() != "void" {
		fn.Rtn = &Param{
			IsConst: rtnType.IsConstQualifiedType(),
		}
		switch rtnType.Kind() {
		case clang.Type_Pointer:
			fn.Rtn.IsPointer = true

			ptrType := rtnType.PointeeType()

			switch ptrType.Kind() {
			case clang.Type_Record:
				fn.Rtn.IsStruct = true
				under := ptrType.Declaration()
				fn.Rtn.Type = under.Spelling()
			default:
				fn.Rtn.Type = ptrType.Kind().String()
			}
		default:
			fn.Rtn.Type = rtnType.Kind().String()
		}
	}

	var param *Param

	for i := range cursor.NumArguments() {
		arg := cursor.Argument(uint32(i))
		argType := arg.Type().CanonicalType()

		if param == nil {
			param = &Param{
				Name:    arg.Spelling(),
				IsConst: argType.IsConstQualifiedType(),
			}
		}

		switch argType.Kind() {
		case clang.Type_Pointer:
			param.IsPointer = true

			ptrType := argType.PointeeType()

			switch ptrType.Kind() {
			case clang.Type_Record:
				param.IsStruct = true
				under := ptrType.Declaration()
				param.Type = under.Spelling()
			default:
				param.Type = ptrType.Kind().String()
			}
		case clang.Type_IncompleteArray:
			eleType := argType.ArrayElementType()

			if eleType.Kind() == clang.Type_Pointer {
				param.Type = eleType.PointeeType().Kind().String()
			} else {
				param.Type = eleType.Kind().String()
			}

			param.IsArray = true
			param.IsPointer = true
			continue
		default:
			if param.IsArray {
				param.ArrSizeName = arg.Spelling()
			} else {
				param.Type = argType.Kind().String()
			}
		}

		param.Fn = fn
		fn.Params = append(fn.Params, param)
		param = nil
	}
}

func (fn ClsMethod) String() string {
	buff := bytes.NewBufferString("")

	fmt.Fprintf(buff, "%s\n", fn.Comments)

	if fn.IsStatic {
		buff.WriteString("static ")
	}

	if fn.Rtn != nil {
		fmt.Fprintf(buff, "%+v ", fn.Rtn)
	}

	fmt.Fprintf(buff, "%s(", fn.Name)

	for idx, p := range fn.Params {
		if idx > 0 {
			buff.WriteString(", ")
		}
		fmt.Fprintf(buff, "%+v", p)
	}
	buff.WriteString(")\n")

	return buff.String()
}

type ClassDefine struct {
	baseDefine

	// 类静态函数列表
	Statics []*ClsMethod
	// 类成员方法列表
	Methods []*ClsMethod
}

func (c ClassDefine) String() string {
	buff := bytes.NewBufferString("")

	fmt.Fprintf(buff, "class %s\n", c.Name)

	for idx, m := range c.Statics {
		if idx > 0 {
			buff.WriteString("\n")
		}

		fmt.Fprintf(buff, "%s", m)
	}

	for _, m := range c.Methods {
		fmt.Fprintf(buff, "\n%s", m)
	}

	return buff.String()
}

func (c *ClassDefine) walkMethods(cursor, parent clang.Cursor) clang.ChildVisitResult {
	switch cursor.Kind() {
	case clang.Cursor_CXXMethod:
		methodName := cursor.Spelling()

		method := ClsMethod{
			baseDefine: baseDefine{
				Name:     methodName,
				Comments: ParseComment(cursor.ParsedComment()),
			},
			MangledName: cursor.Mangling(),
			IsVirtual:   cursor.CXXMethod_IsVirtual(),
			IsStatic:    cursor.CXXMethod_IsStatic(),
		}

		method.ParseParam(&cursor)

		method.Cls = c
		if method.IsStatic {
			c.Statics = append(c.Statics, &method)
		} else {
			c.Methods = append(c.Methods, &method)
		}
	}

	return clang.ChildVisit_Continue
}

func ParseClass(cursor *clang.Cursor) (*ClassDefine, error) {
	define := ClassDefine{
		baseDefine: baseDefine{
			Name:     cursor.DisplayName(),
			Comments: ParseComment(cursor.ParsedComment()),
		},
	}

	cursor.Visit(define.walkMethods)

	return &define, nil
}

func (e *entry) ParseClass(cursor *clang.Cursor) (*ClassDefine, error) {
	define, err := ParseClass(cursor)
	if err != nil {
		return nil, err
	}

	switch define.Name {
	case e.sdk.apiName:
		if e.apiClass != nil {
			return nil, fmt.Errorf("api class duplicated: %+v", define)
		}

		e.apiClass = define

		for _, fn := range define.Statics {
			switch fn.Name {
			case e.sdk.createCallName:
				e.createCall = fn
			case e.sdk.versionCallName:
				e.versionCall = fn
			}
		}
	case e.sdk.spiName:
		if e.spiClass != nil {
			return nil, fmt.Errorf("spi class duplicated: %+v", define)
		}

		e.spiClass = define
	}

	return define, nil
}
