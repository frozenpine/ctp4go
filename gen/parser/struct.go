package parser

import (
	"bytes"
	"fmt"
	"os"

	"github.com/go-clang/clang-v15/clang"
)

type StructField struct {
	baseDefine

	Type       string
	IsBaseType bool
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
		if field.Type == "" {
			// 结构体成员使用了非 typedef 类型
			field.Type = cursor.Type().Kind().String()
			field.IsBaseType = true
		}
		s.Fields = append(s.Fields, field)
	default:
		fmt.Fprintf(
			os.Stderr, "unsupported field: %+v", fieldKind,
		)
	}

	return clang.ChildVisit_Continue
}

func ParseStruct(cursor *clang.Cursor) (*StructDefine, error) {
	define := StructDefine{
		baseDefine: baseDefine{
			Name:     cursor.DisplayName(),
			Comments: ParseComment(cursor.ParsedComment()),
		},
	}

	cursor.Visit(define.walkFields)

	return &define, nil
}

func (e *Entry) ParseStruct(cursor *clang.Cursor) (*StructDefine, error) {
	define, err := ParseStruct(cursor)
	if err != nil {
		return nil, err
	}

	if err := e.dataCache.Append(define.Name, define); err != nil {
		return nil, fmt.Errorf("%w: data struct duplicated %+v", err, define)
	}

	return define, nil
}
