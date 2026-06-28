package internal

import (
	"errors"

	"github.com/go-clang/clang-v15/clang"
)

type EnumMember struct {
	Name     string
	Value    int64
	Comments []string
}

type EnumDefine struct {
	Name     string
	Type     string
	Comments []string
	Members  []EnumMember
}

func (e *EnumDefine) walkChilds(cursor, parent clang.Cursor) clang.ChildVisitResult {
	memKind := cursor.Kind()
	switch memKind {
	case clang.Cursor_EnumConstantDecl:
		e.Members = append(
			e.Members, EnumMember{
				Name:     cursor.Spelling(),
				Value:    cursor.EnumConstantDeclValue(),
				Comments: ParseComment(cursor.ParsedComment()),
			})
	}

	return clang.ChildVisit_Continue
}

func ParseEnum(cursor *clang.Cursor) (*EnumDefine, error) {
	if cursor.Kind() != clang.Cursor_EnumDecl {
		return nil, errors.New("not enum type")
	}

	define := EnumDefine{
		Name:     cursor.DisplayName(),
		Type:     cursor.EnumDeclIntegerType().Kind().String(),
		Comments: ParseComment(cursor.ParsedComment()),
	}

	cursor.Visit(define.walkChilds)

	return &define, nil
}
