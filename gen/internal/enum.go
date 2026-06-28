package internal

import (
	"bytes"
	"errors"
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
		member := EnumMember{
			baseDefine: baseDefine{
				Name:     cursor.Spelling(),
				Comments: ParseComment(cursor.ParsedComment()),
			},
			Value: cursor.EnumConstantDeclValue(),
		}
		member.trimComments()

		e.Members = append(e.Members, member)
	}

	return clang.ChildVisit_Continue
}

func (e *entry) ParseEnum(cursor *clang.Cursor) (*EnumDefine, error) {
	if cursor.Kind() != clang.Cursor_EnumDecl {
		return nil, errors.New("not enum type")
	}

	define := EnumDefine{
		baseDefine: baseDefine{
			Name:     cursor.DisplayName(),
			Comments: ParseComment(cursor.ParsedComment()),
		},

		Type: cursor.EnumDeclIntegerType().Kind().String(),
	}

	cursor.Visit(define.walkChilds)
	define.trimComments()

	return &define, nil
}
