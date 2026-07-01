package v6_5_1

/*
#cgo CFLAGS: -I. -I${SRCDIR} -I${SRCDIR}/../../dependencies/future/v6.5.1/
#cgo LDFLAGS: -ldl

#include "api_helper.h"
#include "spi_helper.h"
*/
import "C"
import (
	"errors"
	"fmt"
	"log/slog"
	"path/filepath"
	"unsafe"

	"github.com/frozenpine/ctp4go/thost"
	"github.com/frozenpine/ctp4go/thost/types"
)

var (
	// 确保Api封装完整实现了thost中的接口签名
	_ thost.MdApi = &ThostFtdcMdApi{}
)

func CreateThostFtdcMdApi(
	libPath string, FlowPath string, IsUsingUdp bool, IsMulticast bool,
) (*ThostFtdcMdApi, error) {
	if libPath == "" {
		return nil, fmt.Errorf(
			"%w: invalid create args", thost.ErrInvalidArgs,
		)
	}

	libPath, err := filepath.Abs(libPath)
	if err != nil {
		return nil, errors.Join(thost.ErrInvalidArgs, err)
	}

	libCPath := C.CString(libPath)
	defer C.free(unsafe.Pointer(libCPath))

	slog.Info(
		"try to load thost mduser api lib",
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

	pszFlowPath := C.CString(FlowPath)
	defer C.free(unsafe.Pointer(pszFlowPath))

	instance := C.CallCreateFtdcMdApi(
		C.CreateFtdcMdApi(creator),
		pszFlowPath, C.bool(IsUsingUdp), C.bool(IsMulticast),
	)

	if instance == nil {
		return nil, fmt.Errorf(
			"%w: thost mduser api[%s] test mode[%t] create failed",
			thost.ErrApiCreateFailed, libPath,
		)
	}

	return &ThostFtdcMdApi{
		apiPtr:  (*C.CThostFtdcMdApiExt)(instance),
		version: apiVer,
		lib:     lib,
	}, nil
}

type ThostFtdcMdApi struct {
	lib unsafe.Pointer

	version string
	apiPtr  *C.CThostFtdcMdApiExt
	// spiPtr      *ThostFtdcMdSpi
}

func (api *ThostFtdcMdApi) GetApiVersion() string {
	return api.version
}

func (api *ThostFtdcMdApi) Release() {
	slog.Info("executing thost mduser api Release")

	defer func() {
		// if api.spiPtr != nil {
		// 	api.spiPtr.Pinner.Unpin()
		// 	api.spiPtr = nil
		// }

		C.dlclose(api.lib)
	}()

	C.CallRelease(
		api.apiPtr.vtable.CThostFtdcMdApiVTable_Release,
		unsafe.Pointer(api.apiPtr),
	)

	slog.Info("thost mduser api Release executed")

}

func (api *ThostFtdcMdApi) Init() {
	slog.Info("executing thost mduser api Init")

	C.CallInit(
		api.apiPtr.vtable.CThostFtdcMdApiVTable_Init,
		unsafe.Pointer(api.apiPtr),
	)

	slog.Info("thost mduser api Init executed")

}

func (api *ThostFtdcMdApi) Join() int {
	slog.Info("executing thost mduser api Join")

	rtn := C.CallJoin(
		api.apiPtr.vtable.CThostFtdcMdApiVTable_Join,
		unsafe.Pointer(api.apiPtr),
	)

	slog.Info("thost mduser api Join executed")

	return
}

func (api *ThostFtdcMdApi) GetTradingDay() string {
	slog.Info("executing thost mduser api GetTradingDay")

	rtn := C.CallGetTradingDay(
		api.apiPtr.vtable.CThostFtdcMdApiVTable_GetTradingDay,
		unsafe.Pointer(api.apiPtr),
	)

	slog.Info("thost mduser api GetTradingDay executed")

	return
}

func (api *ThostFtdcMdApi) RegisterFront(FrontAddress string) {
	slog.Info("executing thost mduser api RegisterFront")

	pszFrontAddress := C.CString(FrontAddress)
	defer C.free(unsafe.Pointer(pszFrontAddress))

	C.CallRegisterFront(
		api.apiPtr.vtable.CThostFtdcMdApiVTable_RegisterFront,
		unsafe.Pointer(api.apiPtr),
		pszFrontAddress,
	)

	slog.Info("thost mduser api RegisterFront executed")

}

func (api *ThostFtdcMdApi) RegisterNameServer(NsAddress string) {
	slog.Info("executing thost mduser api RegisterNameServer")

	pszNsAddress := C.CString(NsAddress)
	defer C.free(unsafe.Pointer(pszNsAddress))

	C.CallRegisterNameServer(
		api.apiPtr.vtable.CThostFtdcMdApiVTable_RegisterNameServer,
		unsafe.Pointer(api.apiPtr),
		pszNsAddress,
	)

	slog.Info("thost mduser api RegisterNameServer executed")

}

func (api *ThostFtdcMdApi) RegisterFensUserInfo(FensUserInfo *thost.CThostFtdcFensUserInfoField) {
	slog.Info("executing thost mduser api RegisterFensUserInfo")

	C.CallRegisterFensUserInfo(
		api.apiPtr.vtable.CThostFtdcMdApiVTable_RegisterFensUserInfo,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcFensUserInfoField)(unsafe.Pointer(FensUserInfo)),
	)

	slog.Info("thost mduser api RegisterFensUserInfo executed")

}

func (api *ThostFtdcMdApi) RegisterSpi(Spi *thost.CThostFtdcMdSpi) {
	slog.Info("executing thost mduser api RegisterSpi")

	C.CallRegisterSpi(
		api.apiPtr.vtable.CThostFtdcMdApiVTable_RegisterSpi,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcMdSpi)(unsafe.Pointer(Spi)),
	)

	slog.Info("thost mduser api RegisterSpi executed")

}

func (api *ThostFtdcMdApi) SubscribeMarketData(InstrumentID ...string) int {
	slog.Info("executing thost mduser api SubscribeMarketData")

	cntInstrumentID := len(InstrumentID)
	ppInstrumentID := C.malloc(C.int(cntInstrumentID))
	defer C.free(unsafe.Pointer(ppInstrumentID))

	rtn := C.CallSubscribeMarketData(
		api.apiPtr.vtable.CThostFtdcMdApiVTable_SubscribeMarketData,
		unsafe.Pointer(api.apiPtr),
		(**C.char)(ppInstrumentID), C.int(cntInstrumentID),
	)

	slog.Info("thost mduser api SubscribeMarketData executed")

	return
}

func (api *ThostFtdcMdApi) UnSubscribeMarketData(InstrumentID ...string) int {
	slog.Info("executing thost mduser api UnSubscribeMarketData")

	cntInstrumentID := len(InstrumentID)
	ppInstrumentID := C.malloc(C.int(cntInstrumentID))
	defer C.free(unsafe.Pointer(ppInstrumentID))

	rtn := C.CallUnSubscribeMarketData(
		api.apiPtr.vtable.CThostFtdcMdApiVTable_UnSubscribeMarketData,
		unsafe.Pointer(api.apiPtr),
		(**C.char)(ppInstrumentID), C.int(cntInstrumentID),
	)

	slog.Info("thost mduser api UnSubscribeMarketData executed")

	return
}

func (api *ThostFtdcMdApi) SubscribeForQuoteRsp(InstrumentID ...string) int {
	slog.Info("executing thost mduser api SubscribeForQuoteRsp")

	cntInstrumentID := len(InstrumentID)
	ppInstrumentID := C.malloc(C.int(cntInstrumentID))
	defer C.free(unsafe.Pointer(ppInstrumentID))

	rtn := C.CallSubscribeForQuoteRsp(
		api.apiPtr.vtable.CThostFtdcMdApiVTable_SubscribeForQuoteRsp,
		unsafe.Pointer(api.apiPtr),
		(**C.char)(ppInstrumentID), C.int(cntInstrumentID),
	)

	slog.Info("thost mduser api SubscribeForQuoteRsp executed")

	return
}

func (api *ThostFtdcMdApi) UnSubscribeForQuoteRsp(InstrumentID ...string) int {
	slog.Info("executing thost mduser api UnSubscribeForQuoteRsp")

	cntInstrumentID := len(InstrumentID)
	ppInstrumentID := C.malloc(C.int(cntInstrumentID))
	defer C.free(unsafe.Pointer(ppInstrumentID))

	rtn := C.CallUnSubscribeForQuoteRsp(
		api.apiPtr.vtable.CThostFtdcMdApiVTable_UnSubscribeForQuoteRsp,
		unsafe.Pointer(api.apiPtr),
		(**C.char)(ppInstrumentID), C.int(cntInstrumentID),
	)

	slog.Info("thost mduser api UnSubscribeForQuoteRsp executed")

	return
}

func (api *ThostFtdcMdApi) ReqUserLogin(ReqUserLoginField *thost.CThostFtdcReqUserLoginField, RequestID int) int {
	slog.Info("executing thost mduser api ReqUserLogin")

	rtn := C.CallReqUserLogin(
		api.apiPtr.vtable.CThostFtdcMdApiVTable_ReqUserLogin,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcReqUserLoginField)(unsafe.Pointer(ReqUserLoginField)), C.int(RequestID),
	)

	slog.Info("thost mduser api ReqUserLogin executed")

	return
}

func (api *ThostFtdcMdApi) ReqUserLogout(UserLogout *thost.CThostFtdcUserLogoutField, RequestID int) int {
	slog.Info("executing thost mduser api ReqUserLogout")

	rtn := C.CallReqUserLogout(
		api.apiPtr.vtable.CThostFtdcMdApiVTable_ReqUserLogout,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcUserLogoutField)(unsafe.Pointer(UserLogout)), C.int(RequestID),
	)

	slog.Info("thost mduser api ReqUserLogout executed")

	return
}

func (api *ThostFtdcMdApi) ReqQryMulticastInstrument(QryMulticastInstrument *thost.CThostFtdcQryMulticastInstrumentField, RequestID int) int {
	slog.Info("executing thost mduser api ReqQryMulticastInstrument")

	rtn := C.CallReqQryMulticastInstrument(
		api.apiPtr.vtable.CThostFtdcMdApiVTable_ReqQryMulticastInstrument,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQryMulticastInstrumentField)(unsafe.Pointer(QryMulticastInstrument)), C.int(RequestID),
	)

	slog.Info("thost mduser api ReqQryMulticastInstrument executed")

	return
}
