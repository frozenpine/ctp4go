package parser

import (
	"bytes"
	"fmt"
	"os"
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
	baseDefine

	Underlying  UnderType
	MacroDefine *MacroGroup
}

func (t TypedefDefine) String() string {
	buff := bytes.NewBufferString("")

	fmt.Fprintf(buff, "%s\n", t.Comments)

	fmt.Fprintf(
		buff, "typedef %s %s\n",
		t.Name, t.Underlying,
	)

	if t.MacroDefine != nil {
		for _, field := range t.MacroDefine.Defines {
			fmt.Fprintf(
				buff, "\t%s %s\n", field.Name, field.Token,
			)
		}
	}

	return buff.String()
}

func (t TypedefDefine) HasDefine() bool {
	return t.MacroDefine != nil
}

func ParseTypedef(cursor *clang.Cursor) (*TypedefDefine, error) {
	define := TypedefDefine{
		baseDefine: baseDefine{
			Name:     cursor.DisplayName(),
			Comments: ParseComment(cursor.ParsedComment()),
		},
	}
	define.Underlying.typ = cursor.TypedefDeclUnderlyingType()
	define.Underlying.kind = define.Underlying.typ.Kind()

	switch define.Underlying.kind {
	case clang.Type_ConstantArray:
		elemType := define.Underlying.typ.ArrayElementType()

		define.Underlying.Name = elemType.Kind().String()
		define.Underlying.Size = define.Underlying.typ.ArraySize()
	case clang.Type_Char_S, clang.Type_Char_U,
		clang.Type_Double, clang.Type_Int,
		clang.Type_Short, clang.Type_UShort,
		clang.Type_Long, clang.Type_ULong,
		clang.Type_LongLong:
		define.Underlying.Name = define.Underlying.kind.String()
	default:
		fmt.Fprintf(
			os.Stderr, "unsupported typedef: %s %s",
			define.Name, define.Underlying.String(),
		)
	}

	return &define, nil
}

func (e *Entry) ParseTypedef(cursor *clang.Cursor) (*TypedefDefine, error) {
	define, err := ParseTypedef(cursor)
	if err != nil {
		return nil, err
	}

	macroTypeName := strings.Replace(
		define.Name, "TThost", "T", 1,
	)

	if g, exist := e.defineCache[macroTypeName]; exist {
		define.Comments = g.Comments
		define.MacroDefine = g
	}

	if err := e.typeCache.Append(define.Name, define); err != nil {
		old := e.typeCache.Get(define.Name)
		if define.Underlying.Name != old.Underlying.Name {
			return nil, fmt.Errorf("typedef conflicted: %+v", define)
		} else {
			fmt.Fprintf(
				os.Stderr, "typedef duplicated: \n%+v\n", define,
			)
		}
	}

	return define, nil
}
