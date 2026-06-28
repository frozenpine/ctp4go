package internal

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/go-clang/clang-v15/clang"
)

type EnumMember struct {
	baseDefine

	Value int64
}

type EnumDefine struct {
	baseDefine

	Type    string
	Members []EnumMember
}

func (e EnumDefine) String() string {
	buff := bytes.NewBufferString("")

	fmt.Fprintf(buff, "enum %s {\n", e.Name)
	for _, m := range e.Members {
		if len(m.Comments.Summary) < 1 {
			v := strings.Split(m.Name, "_")
			m.Comments.Summary = append(m.Comments.Summary, v[len(v)-1])
		}

		fmt.Fprintf(
			buff, "\t%s = %d %s",
			m.Name, m.Value, m.Comments,
		)
	}
	buff.WriteString("}\n")

	return buff.String()
}

func (e *EnumDefine) walkChilds(cursor, parent clang.Cursor) clang.ChildVisitResult {
	memKind := cursor.Kind()
	switch memKind {
	case clang.Cursor_EnumConstantDecl:
		member := EnumMember{
			baseDefine: baseDefine{
				Name:     cursor.Spelling(),
				Comments: ParseComment(cursor.ParsedComment()),
			},
			Value: cursor.EnumConstantDeclValue(),
		}

		e.Members = append(e.Members, member)
	}

	return clang.ChildVisit_Continue
}

func (e *entry) ParseEnum(cursor *clang.Cursor) (*EnumDefine, error) {
	define := EnumDefine{
		baseDefine: baseDefine{
			Name:     cursor.DisplayName(),
			Comments: ParseComment(cursor.ParsedComment()),
		},

		Type: cursor.EnumDeclIntegerType().Kind().String(),
	}

	cursor.Visit(define.walkChilds)

	return &define, nil
}
