package internal

import (
	"errors"

	"github.com/go-clang/clang-v15/clang"
)

type ClassDefine struct {
	baseDefine
}

func (e *entry) ParseClass(cursor *clang.Cursor) (*ClassDefine, error) {
	if cursor.Kind() != clang.Cursor_ClassDecl {
		return nil, errors.New("not class type")
	}

	define := ClassDefine{}

	return &define, nil
}
