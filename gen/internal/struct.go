package internal

import (
	"errors"

	"github.com/go-clang/clang-v15/clang"
)

type StructField struct {
	Name     string
	Type     string
	Comments []string
}

type StructDefine struct {
	Name     string
	Comments []string
	Fields   []StructField
}

func ParseStruct(cursor *clang.Cursor) (*StructDefine, error) {
	if cursor.Kind() != clang.Cursor_StructDecl {
		return nil, errors.New("not struct type")
	}

	define := StructDefine{
		Name:     cursor.DisplayName(),
		Comments: ParseComment(cursor.ParsedComment()),
	}

	cursor.Visit(func(cursor, parent clang.Cursor) (status clang.ChildVisitResult) {
		fieldKind := cursor.Kind()

		switch fieldKind {
		case clang.Cursor_FieldDecl:
			define.Fields = append(define.Fields, StructField{
				Name:     cursor.DisplayName(),
				Type:     cursor.Type().DefName(),
				Comments: ParseComment(cursor.ParsedComment()),
			})
		}

		return clang.ChildVisit_Continue
	})

	return &define, nil
}
