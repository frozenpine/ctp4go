package main

// #cgo pkg-config: llvm
import "C"

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/frozenpine/ctp4go/gen/internal"
	"github.com/go-clang/clang-v15/clang"
)

var (
	entryBase string
	catagory  string
	version   string
	ctpType   string

	dataOut string
	sdkOut  string

	headerFileName map[string]string = map[string]string{
		"mduser": "ThostFtdcMdApi.h",
		"trader": "ThostFtdcTraderApi.h",
	}
)

func init() {
	flag.StringVar(&entryBase, "entry", "", "Api headers base DIR")
	flag.StringVar(&catagory, "catagory", "future", "CTP system's catagory")
	flag.StringVar(&version, "version", "", "CTP system's version")
	flag.StringVar(&ctpType, "type", "mduser", "CTP api type")

	flag.StringVar(&dataOut, "data", "", "Types & structures output DIR")
	flag.StringVar(&sdkOut, "sdk", "", "SDK wapper output DIR")

	flag.Parse()
}

func main() {
	idx := clang.NewIndex(0, 0)
	defer idx.Dispose()

	fileBase := filepath.Join(entryBase, catagory, version)
	sdkFile := filepath.Join(
		fileBase, headerFileName[ctpType],
	)

	// clang.TranslationUnit_DetailedPreprocessingRecord 宏解析及展开
	// clang.TranslationUnit_CXXChainedPCH
	tu := idx.ParseTranslationUnit(
		sdkFile, []string{}, nil,
		clang.TranslationUnit_CXXChainedPCH|
			clang.TranslationUnit_DetailedPreprocessingRecord,
	)
	if !tu.IsValid() {
		fmt.Fprintf(os.Stderr, "Parse header file failed: %s", sdkFile)
		return
	}
	defer tu.Dispose()

	cursor := tu.TranslationUnitCursor()
	cursor.Visit(func(cursor, parent clang.Cursor) (status clang.ChildVisitResult) {
		loc := cursor.Location()
		file, line, column, offset := loc.FileLocation()

		if !strings.HasPrefix(file.Name(), fileBase) {
			return clang.Visit_Continue
		}

		kind := cursor.Kind()

		switch kind {
		case clang.Cursor_EnumDecl:
			if enum, err := internal.ParseEnum(&cursor); err != nil {
				fmt.Fprintf(os.Stderr, "enum parse failed: %+v", err)
			} else {
				fmt.Fprintf(
					os.Stdout, "%s\n[%d:%d:%d] %s: %s\n",
					strings.Join(enum.Comments, "\n"),
					line, column, offset, enum.Name, enum.Type,
				)
				for _, m := range enum.Members {
					fmt.Fprintf(os.Stdout, "\t%s: %d\n", m.Name, m.Value)
				}
				fmt.Fprintf(os.Stdout, "\n")
			}
		case clang.Cursor_TypedefDecl:
			if typedef, err := internal.ParseTypedef(&cursor); err != nil {
				fmt.Fprintf(os.Stderr, "typedef parse failed: %+v", err)
			} else {
				fmt.Fprintf(
					os.Stdout, "%s\n[%d:%d:%d] %s: [%d]%s\n\n",
					strings.Join(typedef.Comments, "\n"),
					line, column, offset, typedef.Name,
					typedef.Underlying.Size, typedef.Underlying.Name,
				)
			}
		case clang.Cursor_StructDecl:
			if stru, err := internal.ParseStruct(&cursor); err != nil {
				fmt.Fprintf(
					os.Stderr, "struct parse failed: %+v", err,
				)
			} else {
				fmt.Fprintf(
					os.Stdout, "%s\n%s:\n",
					strings.Join(stru.Comments, "\n"),
					stru.Name,
				)
				for _, f := range stru.Fields {
					fmt.Fprintf(
						os.Stdout, "\t%s\n\t%s: %s\n",
						strings.Join(f.Comments, "\n"),
						f.Name, f.Type,
					)
				}
				fmt.Fprintf(os.Stdout, "\n")
			}
		}

		return clang.ChildVisit_Continue
	})
}
