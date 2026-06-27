package main

// #cgo pkg-config: llvm
import "C"

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"

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

	// clang.TranslationUnit_None
	// clang.TranslationUnit_CXXChainedPCH
	tu := idx.ParseTranslationUnit(
		sdkFile, []string{}, nil, clang.TranslationUnit_CXXChainedPCH,
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
			// enumName := cursor.DisplayName()
			// fmt.Fprintf(os.Stdout, "%s\n", enumName)
		case clang.Cursor_TypedefDecl:
			typeName := cursor.DisplayName()

			typeComments := parseComment(cursor.ParsedComment())

			underType := cursor.TypedefDeclUnderlyingType()
			underKind := underType.Kind()

			switch underKind {
			case clang.Type_ConstantArray:
				arrSize := underType.ArraySize()
				elemType := underType.ArrayElementType()
				elemKind := elemType.Kind()

				fmt.Fprintf(
					os.Stdout, "%s\n[%d:%d:%d] %s: [%d]%s\n\n",
					strings.Join(typeComments, "\n"),
					line, column, offset, typeName,
					arrSize, elemKind.String(),
				)
			}
		case clang.Cursor_StructDecl:
			// structName := cursor.DisplayName()

			// fmt.Fprintf(os.Stdout, "%s\n", structName)
		}

		return clang.ChildVisit_Recurse
	})
}
