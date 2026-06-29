package internal

import (
	"bytes"
	"fmt"

	"github.com/go-clang/clang-v15/clang"
)

type Param struct {
	Name        string
	Type        string
	IsPointer   bool
	IsConst     bool
	IsArray     bool
	ArrSizeName string
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

	MangledName string

	Params []*Param
	Rtn    *Param

	IsVirtual bool
	IsStatic  bool
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

		fn.Params = append(fn.Params, param)
		param = nil
	}
}

type ClassDefine struct {
	baseDefine

	Methods []*ClsMethod
}

func (c ClassDefine) String() string {
	buff := bytes.NewBufferString("")

	fmt.Fprintf(buff, "class %s\n", c.Name)

	for idx, m := range c.Methods {
		if idx > 0 {
			buff.WriteString("\n")
		}
		fmt.Fprintf(buff, "%s\n", m.Comments)

		if m.IsStatic {
			buff.WriteString("static ")
		}

		if m.Rtn != nil {
			fmt.Fprintf(buff, "%+v ", m.Rtn)
		}

		fmt.Fprintf(buff, "%s(", m.Name)

		for idx, p := range m.Params {
			if idx > 0 {
				buff.WriteString(", ")
			}
			fmt.Fprintf(buff, "%+v", p)
		}
		buff.WriteString(")\n")
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

			IsVirtual: cursor.CXXMethod_IsVirtual(),
			IsStatic:  cursor.CXXMethod_IsStatic(),
		}

		method.ParseParam(&cursor)

		c.Methods = append(c.Methods, &method)
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
	case e.sdk.spiName:
		if e.spiClass != nil {
			return nil, fmt.Errorf("spi class duplicated: %+v", define)
		}

		e.spiClass = define
	}

	return define, nil
}
