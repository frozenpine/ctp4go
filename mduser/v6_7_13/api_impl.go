package v6_7_13

/*
#cgo CFLAGS: -I. -I${SRCDIR} -I${SRCDIR}/../../dependencies/future/v6.7.13/
#cgo LDFLAGS: -ldl

#include "mduser_api_helper.h"
#include "mduser_spi_helper.h"
*/
import "C"
import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"unsafe"

	"github.com/frozenpine/ctp4go/thost"
	"github.com/frozenpine/ctp4go/thost/types"
)

var (
	_ thost.MdApi = &ThostFtdcMduserApi{}
)

func CreateThostFtdcMduserApi(
	libPath, flowPath string, isUdp, isMulti, isProduct bool,
) (*ThostFtdcMduserApi, error) {
	if libPath == "" || flowPath == "" {
		return nil, fmt.Errorf(
			"%w: invalid create args", thost.ErrInvalidArgs,
		)
	}

	libPath, err := filepath.Abs(libPath)
	if err != nil {
		return nil, errors.Join(thost.ErrInvalidArgs, err)
	}

	flowPath, err = filepath.Abs(flowPath)
	if err = os.MkdirAll(flowPath, os.ModePerm); err != nil {
		return nil, errors.Join(thost.ErrInvalidArgs, err)
	}
	if !strings.HasSuffix(flowPath, "/") {
		flowPath += "/"
	}

	libCPath := C.CString(libPath)
	defer C.free(unsafe.Pointer(libCPath))

	flowCPath := C.CString(flowPath)
	defer C.free(unsafe.Pointer(flowCPath))

	slog.Info(
		"try to load thost trader api lib",
		slog.String("lib_path", libPath),
	)

	lib := C.dlopen(libCPath, C.RTLD_LAZY)
	if lib == nil {
		msg := C.dlerror()

		return nil, fmt.Errorf(
			"%w: %s", thost.ErrLibOpenFailed,
			types.DecodeGBK(([]byte)(C.GoString(msg))),
		)
	}

	slog.Info(
		"thost mduser api lib opened",
		slog.String("lib_path", libPath),
	)

	var (
		createFnName  = C.CString(MDUSER_CREATE_FN_NAME)
		versionFnName = C.CString(MDUSER_VERSION_FN_NAME)

		apiVer string
	)
	defer func() {
		C.free(unsafe.Pointer(createFnName))
		C.free(unsafe.Pointer(versionFnName))
	}()

	creator := C.dlsym(lib, createFnName)
	if creator == nil {
		msg := C.dlerror()

		return nil, fmt.Errorf(
			"%w: %s", thost.ErrLibSymbolNotFound,
			types.DecodeGBK(([]byte)(C.GoString(msg))),
		)
	}

	if versioner := C.dlsym(lib, versionFnName); versioner == nil {
		msg := C.dlerror()

		slog.Error(
			"thost mduser api version fn not found",
			slog.String("error", types.DecodeGBK(([]byte)(C.GoString(msg)))),
		)
	} else {
		apiVer = C.GoString(C.CallGetApiVersion(
			(C.GetApiVersion)(versioner),
		))
	}

	slog.Info(
		"thost trader api creator found, try to create api instance",
	)

	instance := C.CallCreateFtdcMdApi(
		C.CreateFtdcMdApi(creator), flowCPath,
		C.bool(isUdp), C.bool(isMulti), C.bool(isProduct),
	)
	if instance == nil {
		return nil, fmt.Errorf(
			"%w: thost trader api[%s] test mode[%t] create failed",
			thost.ErrApiCreateFailed, libPath, isProduct,
		)
	}

	return &ThostFtdcMduserApi{
		apiPtr:  (*C.CThostFtdcMduserApiExt)(instance),
		version: apiVer,
		lib:     lib,
	}, nil
}

type ThostFtdcMduserApi struct {
	lib unsafe.Pointer

	version     string
	apiPtr      *C.CThostFtdcMduserApiExt
	spiPtr      *ThostFtdcMduserSpi
	releaseOnce sync.Once
}

func (api *ThostFtdcMduserApi) GetApiVersion() string {
	return api.version
}

func (api *ThostFtdcMduserApi) Release() {
	api.releaseOnce.Do(func() {
		slog.Info("executing thost mduser api Release")

		defer func() {
			if api.spiPtr != nil {
				api.spiPtr.Pinner.Unpin()
				api.spiPtr = nil
			}

			C.dlclose(api.lib)
		}()

		C.CallRelease(
			api.apiPtr.vtable.CThostFtdcMduserApiVTable_Release,
			unsafe.Pointer(&api.apiPtr),
		)

		slog.Info("thost mduser api Release executed")
	})
}

func (api *ThostFtdcMduserApi) Init() {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"executing thost mduser api Init",
	)

	C.CallInit(
		api.apiPtr.vtable.CThostFtdcMduserApiVTable_Init,
		unsafe.Pointer(api.apiPtr),
	)

	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"thost mduser api Init executed",
	)
}

func (api *ThostFtdcMduserApi) Join() int {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"executing thost mduser api Join",
	)

	rtn := C.CallJoin(
		api.apiPtr.vtable.CThostFtdcMduserApiVTable_Join,
		unsafe.Pointer(api.apiPtr),
	)

	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"thost mduser api Join executed",
	)

	return int(rtn)
}

func (api *ThostFtdcMduserApi) GetTradingDay() string {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"executing thost mdser api GetTradingDay",
	)

	rtn := C.CallGetTradingDay(
		api.apiPtr.vtable.CThostFtdcMduserApiVTable_GetTradingDay,
		unsafe.Pointer(api.apiPtr),
	)

	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"thost mduser api GetTradingDay executed",
	)

	return C.GoString(rtn)
}

func (api *ThostFtdcMduserApi) RegisterFront(pszFrontAddress string) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"executing thost mduser api RegisterFront",
		slog.String("front", pszFrontAddress),
	)

	front := C.CString(pszFrontAddress)
	defer C.free(unsafe.Pointer(front))

	C.CallRegisterFront(
		api.apiPtr.vtable.CThostFtdcMduserApiVTable_RegisterFront,
		unsafe.Pointer(api.apiPtr), front,
	)

	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"thost mduser api RegisterFront executed",
	)
}

func (api *ThostFtdcMduserApi) RegisterNameServer(pszNsAddress string) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"executing thost mduser api RegisterNameServer",
		slog.String("nameserver", pszNsAddress),
	)

	ns := C.CString(pszNsAddress)
	defer C.free(unsafe.Pointer(ns))

	C.CallRegisterNameServer(
		api.apiPtr.vtable.CThostFtdcMduserApiVTable_RegisterNameServer,
		unsafe.Pointer(api.apiPtr), ns,
	)

	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"thost mduser api RegisterNameServer executed",
	)
}

func (api *ThostFtdcMduserApi) RegisterFensUserInfo(
	pFensUserInfo *thost.CThostFtdcFensUserInfoField,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"executing thost mduser api RegisterFensUserInfo",
	)

	C.CallRegisterFensUserInfo(
		api.apiPtr.vtable.CThostFtdcMduserApiVTable_RegisterFensUserInfo,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcFensUserInfoField)(unsafe.Pointer(pFensUserInfo)),
	)

	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"thost mduser api RegisterFensUserInfo executed",
	)
}

func (api *ThostFtdcMduserApi) RegisterSpi(pSpi thost.MdSpi) {
	if pSpi == nil {
		slog.Error("register spi is nil")
		return
	}

	api.spiPtr = new(ThostFtdcMduserSpi)
	api.spiPtr.callback = pSpi

	api.spiPtr.Pinner.Pin(unsafe.Pointer(api.spiPtr))
	slog.Info(
		"thost mduser spi created & pinned",
		slog.Any("spi", pSpi),
		slog.Any("pinned", api.spiPtr),
	)

	cSpi := C.malloc(C.sizeof_CThostFtdcMduserSpiExt)
	(*C.CThostFtdcMduserSpiExt)(cSpi).vtable = spiCVtablePtr
	(*C.CThostFtdcMduserSpiExt)(cSpi).spi = unsafe.Pointer(api.spiPtr)

	slog.Info(
		"registering thost mduser spi",
		slog.Any("spi", pSpi),
		slog.Any("pinned", api.spiPtr),
	)

	C.CallRegisterSpi(
		api.apiPtr.vtable.CThostFtdcMduserApiVTable_RegisterSpi,
		unsafe.Pointer(api.apiPtr), cSpi,
	)

	slog.Info(
		"thost mduser spi registered",
		slog.Any("spi", pSpi),
		slog.Any("pinned", api.spiPtr),
	)
}

func (api *ThostFtdcMduserApi) SubscribeMarketData(
	ppInstrumentID ...string,
) int {
	if len(ppInstrumentID) < 1 {
		slog.Warn(
			"no subscribe instruments specified when SubscribeMarketData",
		)
		return 0
	}

	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"executing thost mduser api SubscribeMarketData",
	)

	buff := make([]*C.char, len(ppInstrumentID))
	for idx, v := range ppInstrumentID {
		buff[idx] = C.CString(v)
	}
	defer func() {
		for _, p := range buff {
			C.free(unsafe.Pointer(p))
		}
	}()

	rtn := C.CallSubscribeMarketData(
		api.apiPtr.vtable.CThostFtdcMduserApiVTable_SubscribeMarketData,
		unsafe.Pointer(api.apiPtr),
		(**C.char)(unsafe.Pointer(&buff[0])),
		C.int(len(ppInstrumentID)),
	)

	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"thost mduser api SubscribeMarketData executed",
	)

	return int(rtn)
}

func (api *ThostFtdcMduserApi) UnSubscribeMarketData(
	ppInstrumentID ...string,
) int {
	if len(ppInstrumentID) < 1 {
		slog.Warn(
			"no subscribe instruments specified when UnSubscribeMarketData",
		)
		return 0
	}

	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"executing thost mduser api UnSubscribeMarketData",
	)

	buff := make([]*C.char, len(ppInstrumentID))
	for idx, v := range ppInstrumentID {
		buff[idx] = C.CString(v)
	}
	defer func() {
		for _, p := range buff {
			C.free(unsafe.Pointer(p))
		}
	}()

	rtn := C.CallUnSubscribeMarketData(
		api.apiPtr.vtable.CThostFtdcMduserApiVTable_UnSubscribeMarketData,
		unsafe.Pointer(api.apiPtr),
		(**C.char)(unsafe.Pointer(&buff[0])),
		C.int(len(ppInstrumentID)),
	)

	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"thost mduser api UnSubscribeMarketData executed",
	)

	return int(rtn)
}

func (api *ThostFtdcMduserApi) SubscribeForQuoteRsp(
	ppInstrumentID ...string,
) int {
	if len(ppInstrumentID) < 1 {
		slog.Warn(
			"no subscribe instruments specified when SubscribeForQuoteRsp",
		)
		return 0
	}

	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"executing thost mduser api SubscribeForQuoteRsp",
	)

	buff := make([]*C.char, len(ppInstrumentID))
	for idx, v := range ppInstrumentID {
		buff[idx] = C.CString(v)
	}
	defer func() {
		for _, p := range buff {
			C.free(unsafe.Pointer(p))
		}
	}()

	rtn := C.CallSubscribeForQuoteRsp(
		api.apiPtr.vtable.CThostFtdcMduserApiVTable_SubscribeForQuoteRsp,
		unsafe.Pointer(api.apiPtr),
		(**C.char)(unsafe.Pointer(&buff[0])),
		C.int(len(ppInstrumentID)),
	)

	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"thost mduser api SubscribeForQuoteRsp executed",
	)

	return int(rtn)
}

func (api *ThostFtdcMduserApi) UnSubscribeForQuoteRsp(
	ppInstrumentID ...string,
) int {
	if len(ppInstrumentID) < 1 {
		slog.Warn(
			"no subscribe instruments specified when UnSubscribeForQuoteRsp",
		)
		return 0
	}

	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"executing thost mduser api UnSubscribeForQuoteRsp",
	)

	buff := make([]*C.char, len(ppInstrumentID))
	for idx, v := range ppInstrumentID {
		buff[idx] = C.CString(v)
	}
	defer func() {
		for _, p := range buff {
			C.free(unsafe.Pointer(p))
		}
	}()

	rtn := C.CallUnSubscribeForQuoteRsp(
		api.apiPtr.vtable.CThostFtdcMduserApiVTable_UnSubscribeForQuoteRsp,
		unsafe.Pointer(api.apiPtr),
		(**C.char)(unsafe.Pointer(&buff[0])),
		C.int(len(ppInstrumentID)),
	)

	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"thost mduser api UnSubscribeForQuoteRsp executed",
	)

	return int(rtn)
}

func (api *ThostFtdcMduserApi) ReqUserLogin(
	pReqUserLoginField *thost.CThostFtdcReqUserLoginField, nRequestID int,
) int {
	slog.Log(context.Background(), slog.LevelDebug-2,
		"executing thost mduser api ReqUserLogin",
	)

	rtn := C.CallReqUserLogin(
		api.apiPtr.vtable.CThostFtdcMduserApiVTable_ReqUserLogin,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcReqUserLoginField)(unsafe.Pointer(
			pReqUserLoginField)),
		C.int(nRequestID),
	)

	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"thost mduser api ReqUserLogin executed",
	)

	return int(rtn)
}

func (api *ThostFtdcMduserApi) ReqUserLogout(
	pUserLogout *thost.CThostFtdcUserLogoutField, nRequestID int,
) int {
	slog.Log(context.Background(), slog.LevelDebug-2,
		"executing thost mduser api ReqUserLogout",
	)

	rtn := C.CallReqUserLogout(
		api.apiPtr.vtable.CThostFtdcMduserApiVTable_ReqUserLogout,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcUserLogoutField)(unsafe.Pointer(pUserLogout)),
		C.int(nRequestID),
	)

	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"thost mduser api ReqUserLogout executed",
	)

	return int(rtn)
}

func (api *ThostFtdcMduserApi) ReqQryMulticastInstrument(
	pQryMulticastInstrument *thost.CThostFtdcQryMulticastInstrumentField,
	nRequestID int,
) int {
	slog.Log(context.Background(), slog.LevelDebug-2,
		"executing thost mduser api ReqQryMulticastInstrument",
	)

	rtn := C.CallReqQryMulticastInstrument(
		api.apiPtr.vtable.CThostFtdcMduserApiVTable_ReqQryMulticastInstrument,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQryMulticastInstrumentField)(unsafe.Pointer(
			pQryMulticastInstrument)),
		C.int(nRequestID),
	)

	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"thost mduser api ReqQryMulticastInstrument executed",
	)

	return int(rtn)
}
