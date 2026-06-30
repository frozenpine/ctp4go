package parser

import (
	"bytes"
	"fmt"

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

	fmt.Fprintf(buff, "%s\n", e.Comments)
	fmt.Fprintf(buff, "enum %s {\n", e.Name)
	for _, m := range e.Members {
		fmt.Fprintf(
			buff, "\t%s = %d %s\n",
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

func ParseEnum(cursor *clang.Cursor) (*EnumDefine, error) {
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

func (e *entry) ParseEnum(cursor *clang.Cursor) (*EnumDefine, error) {
	define, err := ParseEnum(cursor)
	if err != nil {
		return nil, err
	}

	if _, exist := e.enumCache[define.Name]; exist {
		return nil, fmt.Errorf("enum duplicated: %+v", define)
	}

	e.enumCache[define.Name] = define

	return define, nil
}
