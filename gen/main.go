package main

// #cgo pkg-config: llvm
import "C"

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

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

	var entry internal.Entry
	if err := entry.Parse(sdkFile); err != nil {
		fmt.Fprintf(os.Stderr, "parse failed: %+v\n", err)
	} else {
		fmt.Fprintf(os.Stdout, "entry file parsed: %s\n", sdkFile)
	}
}
