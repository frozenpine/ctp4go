package internal

import (
	"bytes"
	"errors"
	"fmt"

	"github.com/go-clang/clang-v15/clang"
)

type UnderType struct {
	Name string
	Size int64

	typ  clang.Type
	kind clang.TypeKind
}

func (u UnderType) String() string {
	buff := bytes.NewBufferString("")

	fmt.Fprintf(buff, "[%d]byte", u.Size)

	return buff.String()
}

type TypedefDefine struct {
	Name       string
	Comments   []string
	Underlying UnderType
}

func (t TypedefDefine) String() string {
	buff := bytes.NewBufferString("")

	for _, c := range t.Comments {
		fmt.Fprintf(buff, "// %s\n", c)
	}

	fmt.Fprintf(buff, "type %s %s\n\n", t.Name, t.Underlying)

	return buff.String()
}

func ParseTypedef(cursor *clang.Cursor) (*TypedefDefine, error) {
	if cursor.Kind() != clang.Cursor_TypedefDecl {
		return nil, errors.New("not typedef type")
	}

	define := TypedefDefine{
		Name:     cursor.DisplayName(),
		Comments: ParseComment(cursor.ParsedComment()),
	}
	define.Underlying.typ = cursor.TypedefDeclUnderlyingType()
	define.Underlying.kind = define.Underlying.typ.Kind()

	switch define.Underlying.kind {
	case clang.Type_ConstantArray:
		define.Underlying.Size = define.Underlying.typ.ArraySize()
		elemType := define.Underlying.typ.ArrayElementType()
		define.Underlying.Name = elemType.Kind().String()
	}

	return &define, nil
}
