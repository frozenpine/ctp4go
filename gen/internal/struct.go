package internal

import (
	"bytes"
	"errors"
	"fmt"

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

func (s StructDefine) String() string {
	buff := bytes.NewBufferString("")

	for _, c := range s.Comments {
		fmt.Fprintf(buff, "// %s\n", c)
	}

	fmt.Fprintf(buff, "type %s struct {\n", s.Name)
	for _, f := range s.Fields {
		for _, c := range f.Comments {
			fmt.Fprintf(buff, "\t// %s\n", c)
		}
		fmt.Fprintf(buff, "\t%s %s\n", f.Name, f.Type)
	}
	buff.WriteString("}\n")

	return buff.String()
}

func (s *StructDefine) walkFields(cursor, parent clang.Cursor) clang.ChildVisitResult {
	fieldKind := cursor.Kind()

	switch fieldKind {
	case clang.Cursor_FieldDecl:
		s.Fields = append(s.Fields, StructField{
			Name:     cursor.DisplayName(),
			Type:     cursor.Type().DefName(),
			Comments: ParseComment(cursor.ParsedComment()),
		})
	}

	return clang.ChildVisit_Continue
}

func ParseStruct(cursor *clang.Cursor) (*StructDefine, error) {
	if cursor.Kind() != clang.Cursor_StructDecl {
		return nil, errors.New("not struct type")
	}

	define := StructDefine{
		Name:     cursor.DisplayName(),
		Comments: ParseComment(cursor.ParsedComment()),
	}

	cursor.Visit(define.walkFields)

	return &define, nil
}
