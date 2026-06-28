package internal

import (
	"bytes"
	"errors"
	"fmt"

	"github.com/go-clang/clang-v15/clang"
)

type StructField struct {
	baseDefine

	Type string
}

type StructDefine struct {
	baseDefine

	Fields []StructField
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
		field := StructField{
			baseDefine: baseDefine{
				Name:     cursor.DisplayName(),
				Comments: ParseComment(cursor.ParsedComment()),
			},
			Type: cursor.Type().DefName(),
		}
		field.trimComments()
		s.Fields = append(s.Fields, field)
	}

	return clang.ChildVisit_Continue
}

func (e *entry) ParseStruct(cursor *clang.Cursor) (*StructDefine, error) {
	if cursor.Kind() != clang.Cursor_StructDecl {
		return nil, errors.New("not struct type")
	}

	define := StructDefine{
		baseDefine: baseDefine{
			Name:     cursor.DisplayName(),
			Comments: ParseComment(cursor.ParsedComment()),
		},
	}

	cursor.Visit(define.walkFields)
	define.trimComments()

	return &define, nil
}
