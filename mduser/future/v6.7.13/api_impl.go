package v6_7_13

/*
#cgo CFLAGS: -I. -I${SRCDIR} -I${SRCDIR}/../../../dependencies/future/v6.7.13/
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

	"github.com/frozenpine/ctp4go"
	"github.com/frozenpine/ctp4go/thost"
	"github.com/frozenpine/ctp4go/thost/future"
)

var (
	// 确保Api封装完整实现了thost中的接口签名
	_ future.MdApi = &ThostFtdcMdApi{}
)

func CreateThostFtdcMdApi(
	libPath string, FlowPath string, IsUsingUdp bool, IsMulticast bool, IsProductionMode bool,
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
			ctp4go.DecodeGBK(([]byte)(C.GoString(msg))),
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
			ctp4go.DecodeGBK(([]byte)(C.GoString(msg))),
		)
	}

	if versioner := C.dlsym(lib, versionFnName); versioner == nil {
		msg := C.dlerror()

		slog.Error(
			"thost mduser api version fn not found",
			slog.String("error", ctp4go.DecodeGBK(([]byte)(C.GoString(msg)))),
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
		pszFlowPath, C.bool(IsUsingUdp), C.bool(IsMulticast), C.bool(IsProductionMode),
	)

	if instance == nil {
		return nil, fmt.Errorf(
			"%w: thost mduser api[%s] create failed",
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
	spiPtr  *ThostFtdcMdSpi
}

func (api *ThostFtdcMdApi) GetApiVersion() string {
	return api.version
}

func (api *ThostFtdcMdApi) Release() {
	if api.apiPtr == nil {
		slog.Error(
			"thost mduser api not initialized",
			slog.String("caller", "Release"),
		)
		return
	}

	slog.Info(
		"executing thost mduser api",
		slog.String("caller", "Release"),
	)

	defer func() {
		if api.spiPtr != nil {
			api.spiPtr.Pinner.Unpin()
			api.spiPtr = nil
		}

		C.dlclose(api.lib)
	}()

	C.CallRelease(
		api.apiPtr.vtable.CThostFtdcMdApiVTable_Release,
		unsafe.Pointer(api.apiPtr),
	)

	slog.Info(
		"thost mduser api executed",
		slog.String("caller", "Release"),
	)

}

func (api *ThostFtdcMdApi) Init() {
	if api.apiPtr == nil {
		slog.Error(
			"thost mduser api not initialized",
			slog.String("caller", "Init"),
		)
		return
	}

	slog.Info(
		"executing thost mduser api",
		slog.String("caller", "Init"),
	)

	C.CallInit(
		api.apiPtr.vtable.CThostFtdcMdApiVTable_Init,
		unsafe.Pointer(api.apiPtr),
	)

	slog.Info(
		"thost mduser api executed",
		slog.String("caller", "Init"),
	)

}

func (api *ThostFtdcMdApi) Join() int {
	if api.apiPtr == nil {
		slog.Error(
			"thost mduser api not initialized",
			slog.String("caller", "Join"),
		)
		return -255
	}

	slog.Info(
		"executing thost mduser api",
		slog.String("caller", "Join"),
	)

	rtn := C.CallJoin(
		api.apiPtr.vtable.CThostFtdcMdApiVTable_Join,
		unsafe.Pointer(api.apiPtr),
	)

	slog.Info(
		"thost mduser api executed",
		slog.String("caller", "Join"),
	)

	return int(rtn)
}

func (api *ThostFtdcMdApi) GetTradingDay() string {
	if api.apiPtr == nil {
		slog.Error(
			"thost mduser api not initialized",
			slog.String("caller", "GetTradingDay"),
		)
		var dummy string
		return dummy
	}

	slog.Info(
		"executing thost mduser api",
		slog.String("caller", "GetTradingDay"),
	)

	rtn := C.CallGetTradingDay(
		api.apiPtr.vtable.CThostFtdcMdApiVTable_GetTradingDay,
		unsafe.Pointer(api.apiPtr),
	)

	slog.Info(
		"thost mduser api executed",
		slog.String("caller", "GetTradingDay"),
	)

	return C.GoString(rtn)
}

func (api *ThostFtdcMdApi) RegisterFront(FrontAddress string) {
	if api.apiPtr == nil {
		slog.Error(
			"thost mduser api not initialized",
			slog.String("caller", "RegisterFront"),
		)
		return
	}

	slog.Info(
		"executing thost mduser api",
		slog.String("caller", "RegisterFront"),
	)

	pszFrontAddress := C.CString(FrontAddress)
	defer C.free(unsafe.Pointer(pszFrontAddress))

	C.CallRegisterFront(
		api.apiPtr.vtable.CThostFtdcMdApiVTable_RegisterFront,
		unsafe.Pointer(api.apiPtr),
		pszFrontAddress,
	)

	slog.Info(
		"thost mduser api executed",
		slog.String("caller", "RegisterFront"),
	)

}

func (api *ThostFtdcMdApi) RegisterNameServer(NsAddress string) {
	if api.apiPtr == nil {
		slog.Error(
			"thost mduser api not initialized",
			slog.String("caller", "RegisterNameServer"),
		)
		return
	}

	slog.Info(
		"executing thost mduser api",
		slog.String("caller", "RegisterNameServer"),
	)

	pszNsAddress := C.CString(NsAddress)
	defer C.free(unsafe.Pointer(pszNsAddress))

	C.CallRegisterNameServer(
		api.apiPtr.vtable.CThostFtdcMdApiVTable_RegisterNameServer,
		unsafe.Pointer(api.apiPtr),
		pszNsAddress,
	)

	slog.Info(
		"thost mduser api executed",
		slog.String("caller", "RegisterNameServer"),
	)

}

func (api *ThostFtdcMdApi) RegisterFensUserInfo(FensUserInfo *future.CThostFtdcFensUserInfoField) {
	if api.apiPtr == nil {
		slog.Error(
			"thost mduser api not initialized",
			slog.String("caller", "RegisterFensUserInfo"),
		)
		return
	}

	slog.Info(
		"executing thost mduser api",
		slog.String("caller", "RegisterFensUserInfo"),
	)

	C.CallRegisterFensUserInfo(
		api.apiPtr.vtable.CThostFtdcMdApiVTable_RegisterFensUserInfo,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcFensUserInfoField)(unsafe.Pointer(FensUserInfo)),
	)

	slog.Info(
		"thost mduser api executed",
		slog.String("caller", "RegisterFensUserInfo"),
	)

}

func (api *ThostFtdcMdApi) RegisterSpi(Spi future.MdSpi) {
	if api.apiPtr == nil {
		slog.Error(
			"thost mduser api not initialized",
			slog.String("caller", "RegisterSpi"),
		)
		return
	}

	slog.Info(
		"executing thost mduser api",
		slog.String("caller", "RegisterSpi"),
	)

	api.spiPtr = new(ThostFtdcMdSpi)
	api.spiPtr.callback = Spi

	api.spiPtr.Pinner.Pin(unsafe.Pointer(api.spiPtr))
	slog.Info(
		"thost mduser spi created & pinned",
		slog.Any("spi", Spi),
		slog.Any("pinned", api.spiPtr),
	)

	pSpi := (*C.CThostFtdcMdSpiExt)(C.malloc(
		C.sizeof_CThostFtdcMdSpiExt,
	))
	pSpi.vtable = spiCVtablePtr
	pSpi.spi = unsafe.Pointer(api.spiPtr)

	C.CallRegisterSpi(
		api.apiPtr.vtable.CThostFtdcMdApiVTable_RegisterSpi,
		unsafe.Pointer(api.apiPtr),
		unsafe.Pointer(pSpi),
	)

	slog.Info(
		"thost mduser api executed",
		slog.String("caller", "RegisterSpi"),
	)

}

func (api *ThostFtdcMdApi) SubscribeMarketData(InstrumentID ...string) int {
	if api.apiPtr == nil {
		slog.Error(
			"thost mduser api not initialized",
			slog.String("caller", "SubscribeMarketData"),
		)
		return -255
	}

	slog.Info(
		"executing thost mduser api",
		slog.String("caller", "SubscribeMarketData"),
	)

	ppInstrumentID := make([]*C.char, len(InstrumentID))
	for idx, v := range InstrumentID {
		ppInstrumentID[idx] = C.CString(v)
		defer C.free(unsafe.Pointer(ppInstrumentID[idx]))
	}

	rtn := C.CallSubscribeMarketData(
		api.apiPtr.vtable.CThostFtdcMdApiVTable_SubscribeMarketData,
		unsafe.Pointer(api.apiPtr),
		(**C.char)(unsafe.Pointer(&ppInstrumentID[0])), C.int(len(InstrumentID)),
	)

	slog.Info(
		"thost mduser api executed",
		slog.String("caller", "SubscribeMarketData"),
	)

	return int(rtn)
}

func (api *ThostFtdcMdApi) UnSubscribeMarketData(InstrumentID ...string) int {
	if api.apiPtr == nil {
		slog.Error(
			"thost mduser api not initialized",
			slog.String("caller", "UnSubscribeMarketData"),
		)
		return -255
	}

	slog.Info(
		"executing thost mduser api",
		slog.String("caller", "UnSubscribeMarketData"),
	)

	ppInstrumentID := make([]*C.char, len(InstrumentID))
	for idx, v := range InstrumentID {
		ppInstrumentID[idx] = C.CString(v)
		defer C.free(unsafe.Pointer(ppInstrumentID[idx]))
	}

	rtn := C.CallUnSubscribeMarketData(
		api.apiPtr.vtable.CThostFtdcMdApiVTable_UnSubscribeMarketData,
		unsafe.Pointer(api.apiPtr),
		(**C.char)(unsafe.Pointer(&ppInstrumentID[0])), C.int(len(InstrumentID)),
	)

	slog.Info(
		"thost mduser api executed",
		slog.String("caller", "UnSubscribeMarketData"),
	)

	return int(rtn)
}

func (api *ThostFtdcMdApi) SubscribeForQuoteRsp(InstrumentID ...string) int {
	if api.apiPtr == nil {
		slog.Error(
			"thost mduser api not initialized",
			slog.String("caller", "SubscribeForQuoteRsp"),
		)
		return -255
	}

	slog.Info(
		"executing thost mduser api",
		slog.String("caller", "SubscribeForQuoteRsp"),
	)

	ppInstrumentID := make([]*C.char, len(InstrumentID))
	for idx, v := range InstrumentID {
		ppInstrumentID[idx] = C.CString(v)
		defer C.free(unsafe.Pointer(ppInstrumentID[idx]))
	}

	rtn := C.CallSubscribeForQuoteRsp(
		api.apiPtr.vtable.CThostFtdcMdApiVTable_SubscribeForQuoteRsp,
		unsafe.Pointer(api.apiPtr),
		(**C.char)(unsafe.Pointer(&ppInstrumentID[0])), C.int(len(InstrumentID)),
	)

	slog.Info(
		"thost mduser api executed",
		slog.String("caller", "SubscribeForQuoteRsp"),
	)

	return int(rtn)
}

func (api *ThostFtdcMdApi) UnSubscribeForQuoteRsp(InstrumentID ...string) int {
	if api.apiPtr == nil {
		slog.Error(
			"thost mduser api not initialized",
			slog.String("caller", "UnSubscribeForQuoteRsp"),
		)
		return -255
	}

	slog.Info(
		"executing thost mduser api",
		slog.String("caller", "UnSubscribeForQuoteRsp"),
	)

	ppInstrumentID := make([]*C.char, len(InstrumentID))
	for idx, v := range InstrumentID {
		ppInstrumentID[idx] = C.CString(v)
		defer C.free(unsafe.Pointer(ppInstrumentID[idx]))
	}

	rtn := C.CallUnSubscribeForQuoteRsp(
		api.apiPtr.vtable.CThostFtdcMdApiVTable_UnSubscribeForQuoteRsp,
		unsafe.Pointer(api.apiPtr),
		(**C.char)(unsafe.Pointer(&ppInstrumentID[0])), C.int(len(InstrumentID)),
	)

	slog.Info(
		"thost mduser api executed",
		slog.String("caller", "UnSubscribeForQuoteRsp"),
	)

	return int(rtn)
}

func (api *ThostFtdcMdApi) ReqUserLogin(ReqUserLoginField *future.CThostFtdcReqUserLoginField, RequestID int) int {
	if api.apiPtr == nil {
		slog.Error(
			"thost mduser api not initialized",
			slog.String("caller", "ReqUserLogin"),
		)
		return -255
	}

	slog.Info(
		"executing thost mduser api",
		slog.String("caller", "ReqUserLogin"),
	)

	rtn := C.CallReqUserLogin(
		api.apiPtr.vtable.CThostFtdcMdApiVTable_ReqUserLogin,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcReqUserLoginField)(unsafe.Pointer(ReqUserLoginField)), C.int(RequestID),
	)

	slog.Info(
		"thost mduser api executed",
		slog.String("caller", "ReqUserLogin"),
	)

	return int(rtn)
}

func (api *ThostFtdcMdApi) ReqUserLogout(UserLogout *future.CThostFtdcUserLogoutField, RequestID int) int {
	if api.apiPtr == nil {
		slog.Error(
			"thost mduser api not initialized",
			slog.String("caller", "ReqUserLogout"),
		)
		return -255
	}

	slog.Info(
		"executing thost mduser api",
		slog.String("caller", "ReqUserLogout"),
	)

	rtn := C.CallReqUserLogout(
		api.apiPtr.vtable.CThostFtdcMdApiVTable_ReqUserLogout,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcUserLogoutField)(unsafe.Pointer(UserLogout)), C.int(RequestID),
	)

	slog.Info(
		"thost mduser api executed",
		slog.String("caller", "ReqUserLogout"),
	)

	return int(rtn)
}

func (api *ThostFtdcMdApi) ReqQryMulticastInstrument(QryMulticastInstrument *future.CThostFtdcQryMulticastInstrumentField, RequestID int) int {
	if api.apiPtr == nil {
		slog.Error(
			"thost mduser api not initialized",
			slog.String("caller", "ReqQryMulticastInstrument"),
		)
		return -255
	}

	slog.Info(
		"executing thost mduser api",
		slog.String("caller", "ReqQryMulticastInstrument"),
	)

	rtn := C.CallReqQryMulticastInstrument(
		api.apiPtr.vtable.CThostFtdcMdApiVTable_ReqQryMulticastInstrument,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQryMulticastInstrumentField)(unsafe.Pointer(QryMulticastInstrument)), C.int(RequestID),
	)

	slog.Info(
		"thost mduser api executed",
		slog.String("caller", "ReqQryMulticastInstrument"),
	)

	return int(rtn)
}
