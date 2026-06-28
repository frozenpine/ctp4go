package internal

import (
	"errors"

	"github.com/go-clang/clang-v15/clang"
)

type UnderType struct {
	Name string
	Size int64
}

type TypedefDefine struct {
	Name       string
	Comments   []string
	Underlying UnderType
}

func ParseTypedef(cursor *clang.Cursor) (*TypedefDefine, error) {
	if cursor.Kind() != clang.Cursor_TypedefDecl {
		return nil, errors.New("not typedef type")
	}

	define := TypedefDefine{
		Name:     cursor.DisplayName(),
		Comments: ParseComment(cursor.ParsedComment()),
	}

	underType := cursor.TypedefDeclUnderlyingType()
	underKind := underType.Kind()

	switch underKind {
	case clang.Type_ConstantArray:
		define.Underlying.Size = underType.ArraySize()
		elemType := underType.ArrayElementType()
		define.Underlying.Name = elemType.Kind().String()
	}

	return &define, nil
}
