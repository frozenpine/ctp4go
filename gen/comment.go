package main

import "github.com/go-clang/clang-v15/clang"

func parseComment(comment clang.Comment) []string {
	var comments []string

	cKind := comment.Kind()

	switch cKind {
	case clang.Comment_FullComment:
	case clang.Comment_Paragraph:
	case clang.Comment_Text:
		return []string{comment.TextComment_getText()}
	default:
		return nil
	}

	for i := range comment.NumChildren() {
		comments = append(
			comments, parseComment(comment.Child(i))...,
		)
	}

	return comments
}
