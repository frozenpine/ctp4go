package v6_7_13

/*
#cgo CFLAGS: -I. -I${SRCDIR} -I${SRCDIR}/../../dependencies/future/v6.7.13/
#cgo LDFLAGS: -ldl

#include "td_api_helper.h"
*/
import "C"
import (
	"errors"
	"fmt"
	"log/slog"
	"path/filepath"
	"runtime"
	"sync"
	"sync/atomic"
	"unsafe"

	"github.com/frozenpine/ctp4go/thost"
)

const (
	CREATE_API_WIN   = ""
	CREATE_API_LINUX = ""

	API_VER_WIN   = ""
	API_VER_LINUX = ""
)

func CreateThostFtdcTraderApi(
	libPath, frontAddr string, options ...thost.ApiOpt,
) (*ThostFtdcTraderApi, error) {
	if libPath == "" || frontAddr == "" {
		return nil, fmt.Errorf(
			"%w: invalid create args", thost.ErrInvalidArgs,
		)
	}

	libPath, err := filepath.Abs(libPath)
	if err != nil {
		return nil, errors.Join(thost.ErrInvalidArgs, err)
	}

	cfg := thost.ApiCfg{}
	for _, opt := range options {
		if opt == nil {
			continue
		}

		if err = opt(&cfg); err != nil {
			return nil, err
		}
	}

	libCPath := C.CString(libPath)
	defer C.free(unsafe.Pointer(libCPath))

	slog.Info(
		"try to load thost trader api lib",
		slog.String("lib_path", libPath),
	)

	lib := C.dlopen(libCPath, C.RTLD_LAZY)
	if lib == nil {
		msg := C.dlerror()

		return nil, fmt.Errorf(
			"%w: %s", thost.ErrLibOpenFailed,
			thost.DecodeGBK(([]byte)(C.GoString(msg))),
		)
	}

	slog.Info(
		"thost trader api lib opened",
		slog.String("lib_path", libPath),
	)

	var (
		createFnName  *C.char
		versionFnName *C.char

		apiVer string
	)

	switch runtime.GOOS {
	case "linux":
		createFnName = C.CString(CREATE_API_LINUX)
		versionFnName = C.CString(API_VER_LINUX)
	case "windows":
		createFnName = C.CString(CREATE_API_WIN)
		versionFnName = C.CString(API_VER_LINUX)
	default:
		return nil, fmt.Errorf(
			"%w: unsupported platform", thost.ErrLibSymbolNotFound,
		)
	}
	defer func() {
		C.free(unsafe.Pointer(createFnName))
		C.free(unsafe.Pointer(versionFnName))
	}()

	creator := C.dlsym(lib, createFnName)
	if creator == nil {
		msg := C.dlerror()

		return nil, fmt.Errorf(
			"%w: %s", thost.ErrLibSymbolNotFound,
			thost.DecodeGBK(([]byte)(C.GoString(msg))),
		)
	}

	if versioner := C.dlsym(lib, versionFnName); versioner == nil {
		msg := C.dlerror()

		slog.Error(
			"thost trader api version fn not found",
			slog.String("error", thost.DecodeGBK(([]byte)(C.GoString(msg)))),
		)
	} else {
		apiVer = C.GoString(C.CallGetApiVersion(
			(C.GetApiVersion)(versioner),
		))
	}

	slog.Info(
		"thost trader api creator found, try to create api instance",
	)

	instance := C.CallCreateFtdcTraderApi(
		C.CreateFtdcTraderApi(creator), libCPath, C.bool(cfg.IsTestVersion),
	)
	if instance == nil {
		return nil, fmt.Errorf(
			"%w: thost trader api[%s] test mode[%t] create failed",
			thost.ErrApiCreateFailed, libPath, cfg.IsTestVersion,
		)
	}

	api := &ThostFtdcTraderApi{
		apiPtr:  (*C.CThostFtdcTraderApiExt)(instance),
		version: apiVer,
		lib:     lib,
	}

	runtime.SetFinalizer(api, func(ins *ThostFtdcTraderApi) {
		// TODO: api release
		C.dlclose(ins.lib)
	})

	return api, nil
}

type ThostFtdcTraderApi struct {
	lib unsafe.Pointer

	version string
	apiPtr  *C.CThostFtdcTraderApiExt
	spiPtr  *ThostFtdcTraderSpi

	registerOnce sync.Once
	reqID        atomic.Int32
}
