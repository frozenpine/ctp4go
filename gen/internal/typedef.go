package internal

import (
	"bytes"
	"errors"
	"fmt"
	"strings"

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

	if u.Size > 0 {
		fmt.Fprintf(buff, "[%d]", u.Size)
	}
	buff.WriteString(u.Name)

	return buff.String()
}

type TypedefDefine struct {
	Name     string
	Comments []string

	Underlying  UnderType
	MacroDefine *MacroGroup
}

func (t TypedefDefine) String() string {
	buff := bytes.NewBufferString("")

	for _, c := range t.Comments {
		fmt.Fprintf(buff, "// %s\n", c)
	}

	fmt.Fprintf(
		buff, "type %s %s\n",
		t.Name, t.Underlying,
	)

	if t.MacroDefine != nil {
		for _, field := range t.MacroDefine.Defines {
			fmt.Fprintf(
				buff, "\t%s %s\n", field.Name, field.Token,
			)
		}
	}

	buff.WriteString("\n")

	return buff.String()
}

func (e *entry) ParseTypedef(cursor *clang.Cursor) (*TypedefDefine, error) {
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
		elemType := define.Underlying.typ.ArrayElementType()

		define.Underlying.Name = elemType.Kind().String()
		define.Underlying.Size = define.Underlying.typ.ArraySize()
	case clang.Type_Char_S, clang.Type_Char_U:
		define.Underlying.Name = define.Underlying.kind.String()

		macroTypeName := strings.Replace(
			define.Name, "TThost", "T", 1,
		)

		if g, exist := e.defineCache[macroTypeName]; exist {
			define.Comments = g.Comments
			define.MacroDefine = g
		}
	}

	return &define, nil
}
