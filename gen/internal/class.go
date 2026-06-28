package internal

import (
	"bytes"
	"fmt"

	"github.com/go-clang/clang-v15/clang"
)

type Param struct {
	Name string
	Type string

	DefaultValue any
}

type ClsMethod struct {
	baseDefine

	MangledName string

	Params []*Param

	IsVirtual bool
	IsStatic  bool
}

func (fn *ClsMethod) walkParam(cursor, parent clang.Cursor) clang.ChildVisitResult {

	return clang.ChildVisit_Continue
}

type ClassDefine struct {
	baseDefine

	Methods []*ClsMethod
}

func (c ClassDefine) String() string {
	buff := bytes.NewBufferString("")

	for _, m := range c.Methods {
		fmt.Fprintf(buff, "%s\n", m.Comments)

		if !m.IsStatic {
			fmt.Fprintf(buff, "%s(%s* this, ", m.Name, c.Name)
		}

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

		cursor.Visit(method.walkParam)

		c.Methods = append(c.Methods, &method)
	}

	return clang.ChildVisit_Continue
}

func (e *entry) ParseClass(cursor *clang.Cursor) (*ClassDefine, error) {
	define := ClassDefine{
		baseDefine: baseDefine{
			Name:     cursor.DisplayName(),
			Comments: ParseComment(cursor.ParsedComment()),
		},
	}

	cursor.Visit(define.walkMethods)

	return &define, nil
}
