package internal

import (
	"errors"
	"fmt"
	"os"

	"github.com/go-clang/clang-v15/clang"
)

type parseConfig struct {
	dataOut  string
	sepTypes bool

	sdkOut string
}

type parseOpt func(*parseConfig) error

type Entry struct {
	filePath string
}

func (e *Entry) walk(cursor, parent clang.Cursor) clang.ChildVisitResult {
	kind := cursor.Kind()

	switch kind {
	case clang.Cursor_EnumDecl:
		if enum, err := ParseEnum(&cursor); err != nil {
			fmt.Fprintf(os.Stderr, "enum parse failed: %+v", err)
		} else {
			fmt.Fprintf(os.Stdout, "%s\n", enum)
		}
	case clang.Cursor_TypedefDecl:
		if typedef, err := ParseTypedef(&cursor); err != nil {
			fmt.Fprintf(os.Stderr, "typedef parse failed: %+v", err)
		} else {
			fmt.Fprintf(os.Stdout, "%s\n", typedef)
		}
	case clang.Cursor_StructDecl:
		if stru, err := ParseStruct(&cursor); err != nil {
			fmt.Fprintf(
				os.Stderr, "struct parse failed: %+v", err,
			)
		} else {
			fmt.Fprintf(os.Stdout, "%s\n", stru)
		}
	}

	return clang.ChildVisit_Continue
}

func (e *Entry) Parse(filePath string, options ...parseOpt) error {
	e.filePath = filePath

	idx := clang.NewIndex(0, 0)
	defer idx.Dispose()

	// clang.TranslationUnit_DetailedPreprocessingRecord 宏解析及展开
	// clang.TranslationUnit_CXXChainedPCH
	tu := idx.ParseTranslationUnit(
		e.filePath, []string{}, nil,
		clang.TranslationUnit_CXXChainedPCH|
			clang.TranslationUnit_DetailedPreprocessingRecord,
	)
	if !tu.IsValid() {
		return errors.New("parse entry file failed")
	}
	defer tu.Dispose()

	cursor := tu.TranslationUnitCursor()
	cursor.Visit(e.walk)

	return nil
}
