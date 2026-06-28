package internal

import (
	"strings"

	"github.com/go-clang/clang-v15/clang"
)

func ParseComment(comment clang.Comment) []string {
	var comments []string

	cKind := comment.Kind()

	switch cKind {
	case clang.Comment_FullComment:
	case clang.Comment_Paragraph:
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
			comments, ParseComment(comment.Child(i))...,
		)
	}

	return comments
}
