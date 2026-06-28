package internal

import (
	"bytes"
	"errors"
	"fmt"
	"strings"

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

func (e EnumDefine) String() string {
	buff := bytes.NewBufferString("")

	fmt.Fprintf(
		buff, "//go:generate stringer -type %s -linecomment\ntype %s int32\n\n",
		e.Name, e.Name,
	)
	fmt.Fprintf(buff, "const (\n")
	for _, m := range e.Members {
		if len(m.Comments) < 1 {
			v := strings.Split(m.Name, "_")
			m.Comments = append(m.Comments, v[len(v)-1])
		}

		fmt.Fprintf(
			buff, "\t%s %s = %d // %s\n",
			m.Name, e.Name, m.Value,
			strings.Join(m.Comments, " "),
		)
	}
	buff.WriteString(")\n")

	return buff.String()
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
