package internal

// #cgo pkg-config: llvm
import "C"

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/go-clang/clang-v15/clang"
)

type sdk string
type platform string

const (
	Trader sdk = "trader"
	Mduser sdk = "mduser"

	PlatFuture platform = "future"
	PlatMini   platform = "mini"
	PlatETF    platform = "etf"

	DefaultDefinePrefix = "THOST_FTDC_"
)

var (
	sdkHeaderName map[sdk]string = map[sdk]string{
		Trader: "ThostFtdcTraderApi.h",
		Mduser: "ThostFtdcMdApi.h",
	}

	sdkSpiName map[sdk]string = map[sdk]string{
		Mduser: "CThostFtdcMdSpi",
		Trader: "CThostFtdcTraderSpi",
	}

	sdkApiName map[sdk]string = map[sdk]string{
		Mduser: "CThostFtdcMdApi",
		Trader: "CThostFtdcTraderApi",
	}
)

type baseDefine struct {
	Name     string
	Comments CommentDefine
}

type parseOpt func(*entry) error

func WithSDK(d string) parseOpt {
	return func(pc *entry) error {
		v := sdk(d)
		switch v {
		case Trader, Mduser:
		default:
			return errors.New("invalid sdk")
		}

		pc.sdk = v
		return nil
	}
}

func WithPlatform(plat string) parseOpt {
	return func(pc *entry) error {
		v := platform(plat)
		switch v {
		case PlatFuture, PlatMini, PlatETF:
		default:
			return errors.New("invalid platform")
		}

		pc.plat = v
		return nil
	}
}

func WithVersion(ver string) parseOpt {
	return func(pc *entry) error {
		if ver == "" {
			return errors.New("invalid version")
		}

		pc.ver = ver
		return nil
	}
}

var CTPEntry = entry{
	sdk:          Trader,
	plat:         PlatFuture,
	definePrefix: DefaultDefinePrefix,
}

type entry struct {
	baseDir   string
	entryPath string

	sdk  sdk
	plat platform
	ver  string

	hdrFileName string
	apiName     string
	spiName     string

	files        map[string]*os.File
	definePrefix string
	defineType   map[string]string
	defineCache  map[string]*MacroGroup
}

func (e *entry) EntryFile() string {
	return e.entryPath
}

func (e *entry) Parse(base string, options ...parseOpt) error {
	for _, opt := range options {
		if opt == nil {
			continue
		}

		if err := opt(e); err != nil {
			return err
		}
	}

	platVerBase := filepath.Join(base, string(e.plat), e.ver)
	e.entryPath = filepath.Join(platVerBase, sdkHeaderName[e.sdk])

	if info, err := os.Stat(e.entryPath); err != nil {
		return err
	} else if info.IsDir() {
		return errors.New("entry path is not file")
	}

	idx := clang.NewIndex(0, 0)
	defer idx.Dispose()

	// clang.TranslationUnit_DetailedPreprocessingRecord 宏解析及展开
	// clang.TranslationUnit_CXXChainedPCH
	// clang.TranslationUnit_SkipFunctionBodies 不解析函数体
	// 解析时可以传入 -CC (包含注释的宏) 和 -fparse-all-comments (解析所有注释风格)
	tu := idx.ParseTranslationUnit(
		e.entryPath, []string{
			"-x", "c++",
			"-fparse-all-comments", "-CC",
		}, nil,
		clang.TranslationUnit_CXXChainedPCH|
			clang.TranslationUnit_DetailedPreprocessingRecord|
			clang.TranslationUnit_SkipFunctionBodies,
	)
	if !tu.IsValid() {
		return errors.New("parse entry file failed")
	}
	defer tu.Dispose()

	cursor := tu.TranslationUnitCursor()
	cursor.Visit(e.walk)

	return nil
}

func (e *entry) walk(cursor, parent clang.Cursor) clang.ChildVisitResult {
	kind := cursor.Kind()

	switch kind {
	case clang.Cursor_EnumDecl:
		if enum, err := e.ParseEnum(&cursor); err != nil {
			fmt.Fprintf(os.Stderr, "enum parse failed: %+v", err)
		} else {
			fmt.Fprintf(os.Stdout, "%s\n", enum)
		}
	case clang.Cursor_TypedefDecl:
		if typedef, err := e.ParseTypedef(&cursor); err != nil {
			fmt.Fprintf(os.Stderr, "typedef parse failed: %+v", err)
		} else {
			fmt.Fprintf(os.Stdout, "%s\n", typedef)
		}
	case clang.Cursor_StructDecl:
		if stru, err := e.ParseStruct(&cursor); err != nil {
			fmt.Fprintf(
				os.Stderr, "struct parse failed: %+v", err,
			)
		} else {
			fmt.Fprintf(os.Stdout, "%s\n", stru)
		}
	case clang.Cursor_ClassDecl:
		if _, err := e.ParseClass(&cursor); err != nil {
			fmt.Fprintf(
				os.Stderr, "class parse failed: %+v", err,
			)
		} else {
			// fmt.Fprintf(os.Stdout, "%s\n", class)
		}
	case clang.Cursor_MacroDefinition:
		if _, err := e.ParseMacro(&cursor, e.definePrefix); err != nil {
			fmt.Fprintf(
				os.Stderr, "class parse failed: %+v", err,
			)
		} /*else if macro != nil {
			fmt.Fprintf(os.Stdout, "%s\n", macro)
		}*/
	}

	return clang.ChildVisit_Continue
}
