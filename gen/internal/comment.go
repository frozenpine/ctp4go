package internal

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/go-clang/clang-v15/clang"
)

type ParamComment struct {
	Cmd         string
	ArgName     string
	Description string
	Values      []string
}

type CommentDefine struct {
	Summary      []string
	ParamComment []*ParamComment
}

func (c *CommentDefine) trimComments() {
	trimmed := make([]string, 0, len(c.Summary))

	for _, c := range c.Summary {
		if v := strings.TrimSpace(
			strings.ReplaceAll(c, "/", ""),
		); v != "" {
			trimmed = append(trimmed, v)
		}
	}

	c.Summary = trimmed
}

func (c CommentDefine) String() string {
	buff := bytes.NewBufferString("")

	if len(c.Summary) < 1 {
		fmt.Fprintf(buff, "// No comment")
	} else {
		for idx, v := range c.Summary {
			if idx > 0 {
				fmt.Fprintf(buff, "\n// %s", v)
			} else {
				fmt.Fprintf(buff, "// %s", v)
			}
		}
	}

	for _, p := range c.ParamComment {
		fmt.Fprintf(buff, "\n//   @%s", p.Cmd)
		if p.ArgName != "" {
			fmt.Fprintf(buff, " %s", p.ArgName)
		}
		if p.Description != "" {
			fmt.Fprintf(buff, " %s", p.Description)
		}
		for _, v := range p.Values {
			fmt.Fprintf(buff, "\n//     %s", v)
		}
	}

	return buff.String()
}

func recursiveComment(comment clang.Comment) []string {
	var comments []string

	cKind := comment.Kind()

	switch cKind {
	case clang.Comment_Paragraph:
	case clang.Comment_BlockCommand:
		cmd := comment.BlockCommandComment_getCommandName()
		comments = append(
			comments, fmt.Sprintf(
				"@%s = ", strings.TrimSpace(cmd)))
	case clang.Comment_ParamCommand:
		comments = append(
			comments, fmt.Sprintf(
				"@param = %s",
				comment.ParamCommandComment_getParamName(),
			))
	case clang.Comment_Text:
		if c := strings.TrimSpace(strings.ReplaceAll(
			comment.TextComment_getText(), "/", "",
		)); c != "" {
			return []string{c}
		} else {
			return nil
		}
	default:
		return nil
	}

	for i := range comment.NumChildren() {
		comments = append(
			comments, recursiveComment(comment.Child(i))...,
		)
	}

	return comments
}

func ParseComment(c clang.Comment) CommentDefine {
	var define CommentDefine

	switch c.Kind() {
	case clang.Comment_FullComment:
		for i := range c.NumChildren() {
			var param *ParamComment

			for idx, c := range recursiveComment(c.Child(i)) {
				if after, ok := strings.CutPrefix(c, "@"); ok {
					values := strings.Split(after, " = ")
					param = &ParamComment{
						Cmd: values[0],
					}
					if param.Cmd == "param" {
						param.ArgName = values[len(values)-1]
					}
				} else if param != nil {
					if idx == 1 {
						param.Description = c
					} else {
						param.Values = append(param.Values, c)
					}
				} else {
					define.Summary = append(define.Summary, c)
				}
			}

			if param != nil {
				define.ParamComment = append(define.ParamComment, param)
			}
		}
	}

	define.trimComments()

	return define
}
