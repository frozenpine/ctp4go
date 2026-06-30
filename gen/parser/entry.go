package parser

// #cgo pkg-config: llvm
import "C"

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/go-clang/clang-v15/clang"
)

type sdkName string
type platform string

const (
	Trader sdkName = "trader"
	Mduser sdkName = "mduser"

	PlatFuture platform = "future"
	PlatMini   platform = "mini"
	PlatETF    platform = "etf"

	DefaultDefinePrefix = "THOST_FTDC_"
)

var (
	sdkHdrFileName map[sdkName]string = map[sdkName]string{
		Trader: "ThostFtdcTraderApi.h",
		Mduser: "ThostFtdcMdApi.h",
	}

	sdkSpiName map[sdkName]string = map[sdkName]string{
		Mduser: "CThostFtdcMdSpi",
		Trader: "CThostFtdcTraderSpi",
	}

	sdkApiName map[sdkName]string = map[sdkName]string{
		Mduser: "CThostFtdcMdApi",
		Trader: "CThostFtdcTraderApi",
	}
)

type baseDefine struct {
	Name     string
	Comments CommentDefine
}

func (d baseDefine) HasComments() bool {
	return len(d.Comments.Summary) > 0 || len(d.Comments.ParamComment) > 0
}

type parseOpt func(*entry) error
type ParseOptions []parseOpt

type sdkOpt func(*sdkInfo) error
type SdkOptions []sdkOpt

func WithHdrFileName(n string) sdkOpt {
	return func(e *sdkInfo) error {
		if n == "" {
			return errors.New("invalid header file name")
		}

		e.hdrFileName = n
		return nil
	}
}

func WithApiName(n string, to ...string) sdkOpt {
	return func(e *sdkInfo) error {
		if n == "" {
			return errors.New("invalid api name")
		}

		e.apiName = n

		switch len(to) {
		case 0:
			e.apiExtName = e.apiName + "Ext"
			e.createCallName = strings.Replace(
				sdkApiName[e.name], "CThost", "Create", 1)
		case 2:
			e.apiExtName = to[1]
			fallthrough
		case 1:
			e.createCallName = to[0]
		default:
			return errors.New("invalid to args")
		}

		return nil
	}
}

func WithSpiName(n string, to ...string) sdkOpt {
	return func(e *sdkInfo) error {
		if n == "" {
			return errors.New("invalid spi name")
		}

		e.spiName = n

		switch len(to) {
		case 0:
			e.spiExtName = e.spiName + "Ext"
		case 1:
			e.spiExtName = to[0]
		default:
			return errors.New("invalid to args")
		}

		return nil
	}
}

func WithVersion(ver string, fn ...string) sdkOpt {
	return func(pc *sdkInfo) error {
		if ver == "" {
			return errors.New("invalid version")
		}

		pc.ver = ver

		if len(fn) > 0 {
			pc.versionCallName = fn[0]
		}

		return nil
	}
}

func WithSDK(d string, options ...sdkOpt) parseOpt {
	return func(e *entry) error {
		v := sdkName(d)
		switch v {
		case Trader, Mduser:
		default:
			return errors.New("invalid sdk")
		}

		e.sdk.name = v

		for _, opt := range append([]sdkOpt{
			WithHdrFileName(sdkHdrFileName[v]),
			WithApiName(sdkApiName[v]),
			WithSpiName(sdkSpiName[v]),
		}, options...) {
			if opt == nil {
				continue
			}

			if err := opt(&e.sdk); err != nil {
				return err
			}
		}

		return nil
	}
}

func withBase(b string) parseOpt {
	return func(e *entry) error {
		if b == "" {
			return errors.New("invalid entry base")
		}

		e.baseDir = b
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

func WithDefPrefix(pre string) parseOpt {
	return func(e *entry) error {
		if pre == "" {
			return errors.New("invalid prefix for #DEFINE")
		}

		e.definePrefix = pre
		return nil
	}
}

func WithDebug() parseOpt {
	return func(e *entry) error {
		e.debug = true
		return nil
	}
}

var CTPEntry = entry{
	sdk: sdkInfo{
		name:        Trader,
		hdrFileName: sdkHdrFileName[Trader],
		apiName:     sdkApiName[Trader],
		apiExtName:  sdkApiName[Trader] + "Ext",
		createCallName: strings.Replace(
			sdkApiName[Trader], "CThost", "Create", 1),
		versionCallName: "GetApiVersion",
		spiName:         sdkSpiName[Trader],
		spiExtName:      sdkSpiName[Trader] + "Ext",
	},
	plat:         PlatFuture,
	definePrefix: DefaultDefinePrefix,
	defineType:   make(map[string]string),
	defineCache:  make(map[string]*MacroGroup),
	enumCache:    make(map[string]*EnumDefine),
	typeCache:    make(map[string]*TypedefDefine),
	dataCache:    make(map[string]*StructDefine),
}

type sdkInfo struct {
	name            sdkName
	ver             string
	hdrFileName     string
	apiName         string
	apiExtName      string
	createCallName  string
	versionCallName string
	spiName         string
	spiExtName      string
}

func (i sdkInfo) validate() error {
	if i.name == "" {
		return errors.New("sdk name not specified")
	}

	if i.ver == "" {
		return errors.New("sdk version not specified")
	}

	if i.hdrFileName == "" {
		return errors.New("sdk hdr file missing")
	}

	if i.apiName == "" {
		return errors.New("sdk api name missing")
	}

	if i.apiExtName == "" {
		return errors.New("sdk api extend name missing")
	}

	if i.spiName == "" {
		return errors.New("sdk spi name missing")
	}

	if i.spiExtName == "" {
		return errors.New("sdk spi extend name missing")
	}

	return nil
}

func (i sdkInfo) Name() string { return string(i.name) }

func (i sdkInfo) Version() string { return i.ver }

func (i sdkInfo) HdrFileName() string { return i.hdrFileName }

func (i sdkInfo) ApiName() string { return i.apiName }

func (i sdkInfo) ApiExtName() string { return i.apiExtName }

func (i sdkInfo) SpiName() string { return i.spiName }

func (i sdkInfo) SpiExtName() string { return i.spiExtName }

type entry struct {
	baseDir   string
	entryPath string
	debug     bool

	plat platform
	sdk  sdkInfo

	files        map[string]*os.File
	releaseOnce  sync.Once
	definePrefix string
	defineType   map[string]string
	defineCache  map[string]*MacroGroup

	typeCache   map[string]*TypedefDefine
	enumCache   map[string]*EnumDefine
	dataCache   map[string]*StructDefine
	apiClass    *ClassDefine
	createCall  *ClsMethod
	versionCall *ClsMethod
	spiClass    *ClassDefine
}

func (e *entry) ApiClass() *ClassDefine { return e.apiClass }

func (e *entry) SpiClass() *ClassDefine { return e.spiClass }

func (e *entry) CreateCall() *ClsMethod { return e.createCall }

func (e *entry) VersionCall() *ClsMethod { return e.versionCall }

func (e *entry) Sdk() sdkInfo { return e.sdk }

func (e *entry) Platform() string { return string(e.plat) }

func (e *entry) EntryFile() string {
	return e.entryPath
}

func (e *entry) validate() error {
	if e.baseDir == "" {
		return errors.New("no entry base specified")
	}

	if e.plat == "" {
		return errors.New("no platform specified")
	}

	return e.sdk.validate()
}

func (e *entry) Release() {
	e.releaseOnce.Do(func() {
		for k, v := range e.files {
			if err := v.Close(); err != nil {
				fmt.Fprintf(
					os.Stderr, "close file[%s] failed: %+v",
					k, err,
				)
			}
		}
	})
}

func (e *entry) Parse(base string, options ...parseOpt) error {
	for _, opt := range append(options, withBase(base)) {
		if opt == nil {
			continue
		}

		if err := opt(e); err != nil {
			return err
		}
	}

	if err := e.validate(); err != nil {
		return err
	}

	platVerBase := filepath.Join(e.baseDir, string(e.plat), e.sdk.ver)
	e.entryPath = filepath.Join(platVerBase, e.sdk.hdrFileName)

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
			"-std=c++11",
			"-fsyntax-only",
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

	if !cursor.Visit(e.walk) {
		return errors.New("ast walk breaked")
	}

	return nil
}

func (e *entry) walk(cursor, parent clang.Cursor) clang.ChildVisitResult {
	kind := cursor.Kind()

	switch kind {
	case clang.Cursor_EnumDecl:
		if enum, err := e.ParseEnum(&cursor); err != nil {
			fmt.Fprintf(os.Stderr, "enum parse failed: %+v", err)
			return clang.ChildVisit_Break
		} else if e.debug {
			fmt.Fprintf(os.Stdout, "%+v\n", enum)
		}
	case clang.Cursor_TypedefDecl:
		if typedef, err := e.ParseTypedef(&cursor); err != nil {
			fmt.Fprintf(os.Stderr, "typedef parse failed: %+v", err)
			return clang.ChildVisit_Break
		} else if e.debug {
			fmt.Fprintf(os.Stdout, "%+v\n", typedef)
		}
	case clang.Cursor_StructDecl:
		if data, err := e.ParseStruct(&cursor); err != nil {
			fmt.Fprintf(os.Stderr, "struct parse failed: %+v", err)
			return clang.ChildVisit_Break
		} else if e.debug {
			fmt.Fprintf(os.Stdout, "%s\n", data)
		}
	case clang.Cursor_ClassDecl:
		if class, err := e.ParseClass(&cursor); err != nil {
			fmt.Fprintf(os.Stderr, "class parse failed: %+v", err)
			return clang.ChildVisit_Break
		} else if e.debug {
			fmt.Fprintf(os.Stdout, "%s\n", class)
		}
	case clang.Cursor_MacroDefinition:
		if macro, err := e.ParseMacro(&cursor, e.definePrefix); err != nil {
			fmt.Fprintf(os.Stderr, "%+v: %+v", err, macro)
			return clang.ChildVisit_Break
		}
	}

	return clang.ChildVisit_Continue
}
