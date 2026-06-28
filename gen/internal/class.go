package internal

import (
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
}

func (fn *ClsMethod) walkParam(cursor, parent clang.Cursor) clang.ChildVisitResult {

	return clang.ChildVisit_Continue
}

type ClassDefine struct {
	baseDefine

	Methods []*ClsMethod
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
