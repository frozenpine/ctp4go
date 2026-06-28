package internal

import (
	"bytes"
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

	fmt.Fprintf(buff, "%s\n", s.Comments)

	fmt.Fprintf(buff, "struct %s {\n", s.Name)
	for _, f := range s.Fields {
		fmt.Fprintf(
			buff, "\t%s\n\t%s %s\n",
			f.Comments, f.Name, f.Type,
		)
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
		s.Fields = append(s.Fields, field)
	}

	return clang.ChildVisit_Continue
}

func (e *entry) ParseStruct(cursor *clang.Cursor) (*StructDefine, error) {
	define := StructDefine{
		baseDefine: baseDefine{
			Name:     cursor.DisplayName(),
			Comments: ParseComment(cursor.ParsedComment()),
		},
	}

	cursor.Visit(define.walkFields)

	return &define, nil
}
