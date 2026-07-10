package v1_7_5

/*
#cgo CFLAGS: -I. -I${SRCDIR} -I${SRCDIR}/../../../dependencies/mini/v1.7.5/
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
	"github.com/frozenpine/ctp4go/thost/mini"
)

var (
	// 确保Api封装完整实现了thost中的接口签名
	_ mini.TraderApi = &ThostFtdcTraderApi{}
)

func CreateThostFtdcTraderApi(
	libPath string, FlowPath string,
) (*ThostFtdcTraderApi, error) {
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
		"try to load thost trader api lib",
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
		"thost trader api lib opened",
		slog.String("lib_path", libPath),
	)

	var (
		createFnName  = C.CString(TRADER_CREATE_FN_NAME)
		versionFnName = C.CString(TRADER_VERSION_FN_NAME)

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

	instance := C.CallCreateFtdcTraderApi(
		C.CreateFtdcTraderApi(creator),
		pszFlowPath,
	)

	if instance == nil {
		return nil, fmt.Errorf(
			"%w: thost trader api[%s] create failed",
			thost.ErrApiCreateFailed, libPath,
		)
	}

	return &ThostFtdcTraderApi{
		apiPtr:  (*C.CThostFtdcTraderApiExt)(instance),
		version: apiVer,
		lib:     lib,
	}, nil
}

type ThostFtdcTraderApi struct {
	lib unsafe.Pointer

	version string
	apiPtr  *C.CThostFtdcTraderApiExt
	spiPtr  *ThostFtdcTraderSpi
}

func (api *ThostFtdcTraderApi) GetApiVersion() string {
	return api.version
}

func (api *ThostFtdcTraderApi) Release() {
	if api.apiPtr == nil {
		slog.Error(
			"thost trader api not initialized",
			slog.String("caller", "Release"),
		)
		return
	}

	slog.Info(
		"executing thost trader api",
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
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_Release,
		unsafe.Pointer(api.apiPtr),
	)

	slog.Info(
		"thost trader api executed",
		slog.String("caller", "Release"),
	)

}

func (api *ThostFtdcTraderApi) Init(Continuous bool) {
	if api.apiPtr == nil {
		slog.Error(
			"thost trader api not initialized",
			slog.String("caller", "Init"),
		)
		return
	}

	slog.Info(
		"executing thost trader api",
		slog.String("caller", "Init"),
	)

	C.CallInit(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_Init,
		unsafe.Pointer(api.apiPtr),
		C.bool(Continuous),
	)

	slog.Info(
		"thost trader api executed",
		slog.String("caller", "Init"),
	)

}

func (api *ThostFtdcTraderApi) Join() int {
	if api.apiPtr == nil {
		slog.Error(
			"thost trader api not initialized",
			slog.String("caller", "Join"),
		)
		return -255
	}

	slog.Info(
		"executing thost trader api",
		slog.String("caller", "Join"),
	)

	rtn := C.CallJoin(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_Join,
		unsafe.Pointer(api.apiPtr),
	)

	slog.Info(
		"thost trader api executed",
		slog.String("caller", "Join"),
	)

	return int(rtn)
}

func (api *ThostFtdcTraderApi) GetTradingDay() string {
	if api.apiPtr == nil {
		slog.Error(
			"thost trader api not initialized",
			slog.String("caller", "GetTradingDay"),
		)
		var dummy string
		return dummy
	}

	slog.Info(
		"executing thost trader api",
		slog.String("caller", "GetTradingDay"),
	)

	rtn := C.CallGetTradingDay(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_GetTradingDay,
		unsafe.Pointer(api.apiPtr),
	)

	slog.Info(
		"thost trader api executed",
		slog.String("caller", "GetTradingDay"),
	)

	return C.GoString(rtn)
}

func (api *ThostFtdcTraderApi) RegisterFront(FrontAddress string) {
	if api.apiPtr == nil {
		slog.Error(
			"thost trader api not initialized",
			slog.String("caller", "RegisterFront"),
		)
		return
	}

	slog.Info(
		"executing thost trader api",
		slog.String("caller", "RegisterFront"),
	)

	pszFrontAddress := C.CString(FrontAddress)
	defer C.free(unsafe.Pointer(pszFrontAddress))

	C.CallRegisterFront(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_RegisterFront,
		unsafe.Pointer(api.apiPtr),
		pszFrontAddress,
	)

	slog.Info(
		"thost trader api executed",
		slog.String("caller", "RegisterFront"),
	)

}

func (api *ThostFtdcTraderApi) RegisterSpi(Spi mini.TraderSpi) {
	if api.apiPtr == nil {
		slog.Error(
			"thost trader api not initialized",
			slog.String("caller", "RegisterSpi"),
		)
		return
	}

	slog.Info(
		"executing thost trader api",
		slog.String("caller", "RegisterSpi"),
	)

	api.spiPtr = new(ThostFtdcTraderSpi)
	api.spiPtr.callback = Spi

	api.spiPtr.Pinner.Pin(unsafe.Pointer(api.spiPtr))
	slog.Info(
		"thost trader spi created & pinned",
		slog.Any("spi", Spi),
		slog.Any("pinned", api.spiPtr),
	)

	pSpi := (*C.CThostFtdcTraderSpiExt)(C.malloc(
		C.sizeof_CThostFtdcTraderSpiExt,
	))
	pSpi.vtable = spiCVtablePtr
	pSpi.spi = unsafe.Pointer(api.spiPtr)

	C.CallRegisterSpi(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_RegisterSpi,
		unsafe.Pointer(api.apiPtr),
		unsafe.Pointer(pSpi),
	)

	slog.Info(
		"thost trader api executed",
		slog.String("caller", "RegisterSpi"),
	)

}

func (api *ThostFtdcTraderApi) SubscribePrivateTopic(ResumeType int) {
	if api.apiPtr == nil {
		slog.Error(
			"thost trader api not initialized",
			slog.String("caller", "SubscribePrivateTopic"),
		)
		return
	}

	slog.Info(
		"executing thost trader api",
		slog.String("caller", "SubscribePrivateTopic"),
	)

	C.CallSubscribePrivateTopic(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_SubscribePrivateTopic,
		unsafe.Pointer(api.apiPtr),
		C.int(ResumeType),
	)

	slog.Info(
		"thost trader api executed",
		slog.String("caller", "SubscribePrivateTopic"),
	)

}

func (api *ThostFtdcTraderApi) SubscribePublicTopic(ResumeType int) {
	if api.apiPtr == nil {
		slog.Error(
			"thost trader api not initialized",
			slog.String("caller", "SubscribePublicTopic"),
		)
		return
	}

	slog.Info(
		"executing thost trader api",
		slog.String("caller", "SubscribePublicTopic"),
	)

	C.CallSubscribePublicTopic(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_SubscribePublicTopic,
		unsafe.Pointer(api.apiPtr),
		C.int(ResumeType),
	)

	slog.Info(
		"thost trader api executed",
		slog.String("caller", "SubscribePublicTopic"),
	)

}

func (api *ThostFtdcTraderApi) SubscribeFlowCtrlWarning(TraderID ...string) int {
	if api.apiPtr == nil {
		slog.Error(
			"thost trader api not initialized",
			slog.String("caller", "SubscribeFlowCtrlWarning"),
		)
		return -255
	}

	slog.Info(
		"executing thost trader api",
		slog.String("caller", "SubscribeFlowCtrlWarning"),
	)

	ppTraderID := make([]*C.char, len(TraderID))
	for idx, v := range TraderID {
		ppTraderID[idx] = C.CString(v)
		defer C.free(unsafe.Pointer(ppTraderID[idx]))
	}

	rtn := C.CallSubscribeFlowCtrlWarning(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_SubscribeFlowCtrlWarning,
		unsafe.Pointer(api.apiPtr),
		(**C.char)(unsafe.Pointer(&ppTraderID[0])), C.int(len(TraderID)),
	)

	slog.Info(
		"thost trader api executed",
		slog.String("caller", "SubscribeFlowCtrlWarning"),
	)

	return int(rtn)
}

func (api *ThostFtdcTraderApi) UnSubscribeFlowCtrlWarning(TraderID ...string) int {
	if api.apiPtr == nil {
		slog.Error(
			"thost trader api not initialized",
			slog.String("caller", "UnSubscribeFlowCtrlWarning"),
		)
		return -255
	}

	slog.Info(
		"executing thost trader api",
		slog.String("caller", "UnSubscribeFlowCtrlWarning"),
	)

	ppTraderID := make([]*C.char, len(TraderID))
	for idx, v := range TraderID {
		ppTraderID[idx] = C.CString(v)
		defer C.free(unsafe.Pointer(ppTraderID[idx]))
	}

	rtn := C.CallUnSubscribeFlowCtrlWarning(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_UnSubscribeFlowCtrlWarning,
		unsafe.Pointer(api.apiPtr),
		(**C.char)(unsafe.Pointer(&ppTraderID[0])), C.int(len(TraderID)),
	)

	slog.Info(
		"thost trader api executed",
		slog.String("caller", "UnSubscribeFlowCtrlWarning"),
	)

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqAuthenticate(ReqAuthenticateField *mini.CThostFtdcReqAuthenticateField, RequestID int) int {
	if api.apiPtr == nil {
		slog.Error(
			"thost trader api not initialized",
			slog.String("caller", "ReqAuthenticate"),
		)
		return -255
	}

	slog.Info(
		"executing thost trader api",
		slog.String("caller", "ReqAuthenticate"),
	)

	rtn := C.CallReqAuthenticate(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqAuthenticate,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcReqAuthenticateField)(unsafe.Pointer(ReqAuthenticateField)), C.int(RequestID),
	)

	slog.Info(
		"thost trader api executed",
		slog.String("caller", "ReqAuthenticate"),
	)

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqUserLogin(ReqUserLoginField *mini.CThostFtdcReqUserLoginField, RequestID int) int {
	if api.apiPtr == nil {
		slog.Error(
			"thost trader api not initialized",
			slog.String("caller", "ReqUserLogin"),
		)
		return -255
	}

	slog.Info(
		"executing thost trader api",
		slog.String("caller", "ReqUserLogin"),
	)

	rtn := C.CallReqUserLogin(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqUserLogin,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcReqUserLoginField)(unsafe.Pointer(ReqUserLoginField)), C.int(RequestID),
	)

	slog.Info(
		"thost trader api executed",
		slog.String("caller", "ReqUserLogin"),
	)

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqUserLoginEncrypt(ReqUserLoginField *mini.CThostFtdcReqUserLoginField, RequestID int) int {
	if api.apiPtr == nil {
		slog.Error(
			"thost trader api not initialized",
			slog.String("caller", "ReqUserLoginEncrypt"),
		)
		return -255
	}

	slog.Info(
		"executing thost trader api",
		slog.String("caller", "ReqUserLoginEncrypt"),
	)

	rtn := C.CallReqUserLoginEncrypt(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqUserLoginEncrypt,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcReqUserLoginField)(unsafe.Pointer(ReqUserLoginField)), C.int(RequestID),
	)

	slog.Info(
		"thost trader api executed",
		slog.String("caller", "ReqUserLoginEncrypt"),
	)

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqUserLogout(UserLogout *mini.CThostFtdcUserLogoutField, RequestID int) int {
	if api.apiPtr == nil {
		slog.Error(
			"thost trader api not initialized",
			slog.String("caller", "ReqUserLogout"),
		)
		return -255
	}

	slog.Info(
		"executing thost trader api",
		slog.String("caller", "ReqUserLogout"),
	)

	rtn := C.CallReqUserLogout(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqUserLogout,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcUserLogoutField)(unsafe.Pointer(UserLogout)), C.int(RequestID),
	)

	slog.Info(
		"thost trader api executed",
		slog.String("caller", "ReqUserLogout"),
	)

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqOrderInsert(InputOrder *mini.CThostFtdcInputOrderField, RequestID int) int {
	if api.apiPtr == nil {
		slog.Error(
			"thost trader api not initialized",
			slog.String("caller", "ReqOrderInsert"),
		)
		return -255
	}

	slog.Info(
		"executing thost trader api",
		slog.String("caller", "ReqOrderInsert"),
	)

	rtn := C.CallReqOrderInsert(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqOrderInsert,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcInputOrderField)(unsafe.Pointer(InputOrder)), C.int(RequestID),
	)

	slog.Info(
		"thost trader api executed",
		slog.String("caller", "ReqOrderInsert"),
	)

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqOrderAction(InputOrderAction *mini.CThostFtdcInputOrderActionField, RequestID int) int {
	if api.apiPtr == nil {
		slog.Error(
			"thost trader api not initialized",
			slog.String("caller", "ReqOrderAction"),
		)
		return -255
	}

	slog.Info(
		"executing thost trader api",
		slog.String("caller", "ReqOrderAction"),
	)

	rtn := C.CallReqOrderAction(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqOrderAction,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcInputOrderActionField)(unsafe.Pointer(InputOrderAction)), C.int(RequestID),
	)

	slog.Info(
		"thost trader api executed",
		slog.String("caller", "ReqOrderAction"),
	)

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqMKBatchOrderAction(MKInputOrderAction *mini.CThostFtdcMKInputOrderActionField, RequestID int) int {
	if api.apiPtr == nil {
		slog.Error(
			"thost trader api not initialized",
			slog.String("caller", "ReqMKBatchOrderAction"),
		)
		return -255
	}

	slog.Info(
		"executing thost trader api",
		slog.String("caller", "ReqMKBatchOrderAction"),
	)

	rtn := C.CallReqMKBatchOrderAction(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqMKBatchOrderAction,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcMKInputOrderActionField)(unsafe.Pointer(MKInputOrderAction)), C.int(RequestID),
	)

	slog.Info(
		"thost trader api executed",
		slog.String("caller", "ReqMKBatchOrderAction"),
	)

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqExecOrderInsert(InputExecOrder *mini.CThostFtdcInputExecOrderField, RequestID int) int {
	if api.apiPtr == nil {
		slog.Error(
			"thost trader api not initialized",
			slog.String("caller", "ReqExecOrderInsert"),
		)
		return -255
	}

	slog.Info(
		"executing thost trader api",
		slog.String("caller", "ReqExecOrderInsert"),
	)

	rtn := C.CallReqExecOrderInsert(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqExecOrderInsert,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcInputExecOrderField)(unsafe.Pointer(InputExecOrder)), C.int(RequestID),
	)

	slog.Info(
		"thost trader api executed",
		slog.String("caller", "ReqExecOrderInsert"),
	)

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqExecOrderAction(InputExecOrderAction *mini.CThostFtdcInputExecOrderActionField, RequestID int) int {
	if api.apiPtr == nil {
		slog.Error(
			"thost trader api not initialized",
			slog.String("caller", "ReqExecOrderAction"),
		)
		return -255
	}

	slog.Info(
		"executing thost trader api",
		slog.String("caller", "ReqExecOrderAction"),
	)

	rtn := C.CallReqExecOrderAction(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqExecOrderAction,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcInputExecOrderActionField)(unsafe.Pointer(InputExecOrderAction)), C.int(RequestID),
	)

	slog.Info(
		"thost trader api executed",
		slog.String("caller", "ReqExecOrderAction"),
	)

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqForQuoteInsert(InputForQuote *mini.CThostFtdcInputForQuoteField, RequestID int) int {
	if api.apiPtr == nil {
		slog.Error(
			"thost trader api not initialized",
			slog.String("caller", "ReqForQuoteInsert"),
		)
		return -255
	}

	slog.Info(
		"executing thost trader api",
		slog.String("caller", "ReqForQuoteInsert"),
	)

	rtn := C.CallReqForQuoteInsert(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqForQuoteInsert,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcInputForQuoteField)(unsafe.Pointer(InputForQuote)), C.int(RequestID),
	)

	slog.Info(
		"thost trader api executed",
		slog.String("caller", "ReqForQuoteInsert"),
	)

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQuoteInsert(InputQuote *mini.CThostFtdcInputQuoteField, RequestID int) int {
	if api.apiPtr == nil {
		slog.Error(
			"thost trader api not initialized",
			slog.String("caller", "ReqQuoteInsert"),
		)
		return -255
	}

	slog.Info(
		"executing thost trader api",
		slog.String("caller", "ReqQuoteInsert"),
	)

	rtn := C.CallReqQuoteInsert(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQuoteInsert,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcInputQuoteField)(unsafe.Pointer(InputQuote)), C.int(RequestID),
	)

	slog.Info(
		"thost trader api executed",
		slog.String("caller", "ReqQuoteInsert"),
	)

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQuoteAction(InputQuoteAction *mini.CThostFtdcInputQuoteActionField, RequestID int) int {
	if api.apiPtr == nil {
		slog.Error(
			"thost trader api not initialized",
			slog.String("caller", "ReqQuoteAction"),
		)
		return -255
	}

	slog.Info(
		"executing thost trader api",
		slog.String("caller", "ReqQuoteAction"),
	)

	rtn := C.CallReqQuoteAction(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQuoteAction,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcInputQuoteActionField)(unsafe.Pointer(InputQuoteAction)), C.int(RequestID),
	)

	slog.Info(
		"thost trader api executed",
		slog.String("caller", "ReqQuoteAction"),
	)

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqBatchOrderAction(InputBatchOrderAction *mini.CThostFtdcInputBatchOrderActionField, RequestID int) int {
	if api.apiPtr == nil {
		slog.Error(
			"thost trader api not initialized",
			slog.String("caller", "ReqBatchOrderAction"),
		)
		return -255
	}

	slog.Info(
		"executing thost trader api",
		slog.String("caller", "ReqBatchOrderAction"),
	)

	rtn := C.CallReqBatchOrderAction(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqBatchOrderAction,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcInputBatchOrderActionField)(unsafe.Pointer(InputBatchOrderAction)), C.int(RequestID),
	)

	slog.Info(
		"thost trader api executed",
		slog.String("caller", "ReqBatchOrderAction"),
	)

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqOptionSelfCloseInsert(InputOptionSelfClose *mini.CThostFtdcInputOptionSelfCloseField, RequestID int) int {
	if api.apiPtr == nil {
		slog.Error(
			"thost trader api not initialized",
			slog.String("caller", "ReqOptionSelfCloseInsert"),
		)
		return -255
	}

	slog.Info(
		"executing thost trader api",
		slog.String("caller", "ReqOptionSelfCloseInsert"),
	)

	rtn := C.CallReqOptionSelfCloseInsert(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqOptionSelfCloseInsert,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcInputOptionSelfCloseField)(unsafe.Pointer(InputOptionSelfClose)), C.int(RequestID),
	)

	slog.Info(
		"thost trader api executed",
		slog.String("caller", "ReqOptionSelfCloseInsert"),
	)

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqOptionSelfCloseAction(InputOptionSelfCloseAction *mini.CThostFtdcInputOptionSelfCloseActionField, RequestID int) int {
	if api.apiPtr == nil {
		slog.Error(
			"thost trader api not initialized",
			slog.String("caller", "ReqOptionSelfCloseAction"),
		)
		return -255
	}

	slog.Info(
		"executing thost trader api",
		slog.String("caller", "ReqOptionSelfCloseAction"),
	)

	rtn := C.CallReqOptionSelfCloseAction(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqOptionSelfCloseAction,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcInputOptionSelfCloseActionField)(unsafe.Pointer(InputOptionSelfCloseAction)), C.int(RequestID),
	)

	slog.Info(
		"thost trader api executed",
		slog.String("caller", "ReqOptionSelfCloseAction"),
	)

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqCombActionInsert(InputCombAction *mini.CThostFtdcInputCombActionField, RequestID int) int {
	if api.apiPtr == nil {
		slog.Error(
			"thost trader api not initialized",
			slog.String("caller", "ReqCombActionInsert"),
		)
		return -255
	}

	slog.Info(
		"executing thost trader api",
		slog.String("caller", "ReqCombActionInsert"),
	)

	rtn := C.CallReqCombActionInsert(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqCombActionInsert,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcInputCombActionField)(unsafe.Pointer(InputCombAction)), C.int(RequestID),
	)

	slog.Info(
		"thost trader api executed",
		slog.String("caller", "ReqCombActionInsert"),
	)

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqSubscribeFundChange(RequestID int) int {
	if api.apiPtr == nil {
		slog.Error(
			"thost trader api not initialized",
			slog.String("caller", "ReqSubscribeFundChange"),
		)
		return -255
	}

	slog.Info(
		"executing thost trader api",
		slog.String("caller", "ReqSubscribeFundChange"),
	)

	rtn := C.CallReqSubscribeFundChange(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqSubscribeFundChange,
		unsafe.Pointer(api.apiPtr),
		C.int(RequestID),
	)

	slog.Info(
		"thost trader api executed",
		slog.String("caller", "ReqSubscribeFundChange"),
	)

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqUnSubscribeFundChange(RequestID int) int {
	if api.apiPtr == nil {
		slog.Error(
			"thost trader api not initialized",
			slog.String("caller", "ReqUnSubscribeFundChange"),
		)
		return -255
	}

	slog.Info(
		"executing thost trader api",
		slog.String("caller", "ReqUnSubscribeFundChange"),
	)

	rtn := C.CallReqUnSubscribeFundChange(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqUnSubscribeFundChange,
		unsafe.Pointer(api.apiPtr),
		C.int(RequestID),
	)

	slog.Info(
		"thost trader api executed",
		slog.String("caller", "ReqUnSubscribeFundChange"),
	)

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqOffsetSetting(InputOffsetSetting *mini.CThostFtdcInputOffsetSettingField, RequestID int) int {
	if api.apiPtr == nil {
		slog.Error(
			"thost trader api not initialized",
			slog.String("caller", "ReqOffsetSetting"),
		)
		return -255
	}

	slog.Info(
		"executing thost trader api",
		slog.String("caller", "ReqOffsetSetting"),
	)

	rtn := C.CallReqOffsetSetting(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqOffsetSetting,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcInputOffsetSettingField)(unsafe.Pointer(InputOffsetSetting)), C.int(RequestID),
	)

	slog.Info(
		"thost trader api executed",
		slog.String("caller", "ReqOffsetSetting"),
	)

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqCancelOffsetSetting(InputOffsetSetting *mini.CThostFtdcInputOffsetSettingField, RequestID int) int {
	if api.apiPtr == nil {
		slog.Error(
			"thost trader api not initialized",
			slog.String("caller", "ReqCancelOffsetSetting"),
		)
		return -255
	}

	slog.Info(
		"executing thost trader api",
		slog.String("caller", "ReqCancelOffsetSetting"),
	)

	rtn := C.CallReqCancelOffsetSetting(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqCancelOffsetSetting,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcInputOffsetSettingField)(unsafe.Pointer(InputOffsetSetting)), C.int(RequestID),
	)

	slog.Info(
		"thost trader api executed",
		slog.String("caller", "ReqCancelOffsetSetting"),
	)

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQryOrder(QryOrder *mini.CThostFtdcQryOrderField, RequestID int) int {
	if api.apiPtr == nil {
		slog.Error(
			"thost trader api not initialized",
			slog.String("caller", "ReqQryOrder"),
		)
		return -255
	}

	slog.Info(
		"executing thost trader api",
		slog.String("caller", "ReqQryOrder"),
	)

	rtn := C.CallReqQryOrder(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQryOrder,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQryOrderField)(unsafe.Pointer(QryOrder)), C.int(RequestID),
	)

	slog.Info(
		"thost trader api executed",
		slog.String("caller", "ReqQryOrder"),
	)

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQryTrade(QryTrade *mini.CThostFtdcQryTradeField, RequestID int) int {
	if api.apiPtr == nil {
		slog.Error(
			"thost trader api not initialized",
			slog.String("caller", "ReqQryTrade"),
		)
		return -255
	}

	slog.Info(
		"executing thost trader api",
		slog.String("caller", "ReqQryTrade"),
	)

	rtn := C.CallReqQryTrade(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQryTrade,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQryTradeField)(unsafe.Pointer(QryTrade)), C.int(RequestID),
	)

	slog.Info(
		"thost trader api executed",
		slog.String("caller", "ReqQryTrade"),
	)

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQryInvestorPosition(QryInvestorPosition *mini.CThostFtdcQryInvestorPositionField, RequestID int) int {
	if api.apiPtr == nil {
		slog.Error(
			"thost trader api not initialized",
			slog.String("caller", "ReqQryInvestorPosition"),
		)
		return -255
	}

	slog.Info(
		"executing thost trader api",
		slog.String("caller", "ReqQryInvestorPosition"),
	)

	rtn := C.CallReqQryInvestorPosition(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQryInvestorPosition,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQryInvestorPositionField)(unsafe.Pointer(QryInvestorPosition)), C.int(RequestID),
	)

	slog.Info(
		"thost trader api executed",
		slog.String("caller", "ReqQryInvestorPosition"),
	)

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQryTradingAccount(QryTradingAccount *mini.CThostFtdcQryTradingAccountField, RequestID int) int {
	if api.apiPtr == nil {
		slog.Error(
			"thost trader api not initialized",
			slog.String("caller", "ReqQryTradingAccount"),
		)
		return -255
	}

	slog.Info(
		"executing thost trader api",
		slog.String("caller", "ReqQryTradingAccount"),
	)

	rtn := C.CallReqQryTradingAccount(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQryTradingAccount,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQryTradingAccountField)(unsafe.Pointer(QryTradingAccount)), C.int(RequestID),
	)

	slog.Info(
		"thost trader api executed",
		slog.String("caller", "ReqQryTradingAccount"),
	)

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQryInvestor(QryInvestor *mini.CThostFtdcQryInvestorField, RequestID int) int {
	if api.apiPtr == nil {
		slog.Error(
			"thost trader api not initialized",
			slog.String("caller", "ReqQryInvestor"),
		)
		return -255
	}

	slog.Info(
		"executing thost trader api",
		slog.String("caller", "ReqQryInvestor"),
	)

	rtn := C.CallReqQryInvestor(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQryInvestor,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQryInvestorField)(unsafe.Pointer(QryInvestor)), C.int(RequestID),
	)

	slog.Info(
		"thost trader api executed",
		slog.String("caller", "ReqQryInvestor"),
	)

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQryTradingCode(QryTradingCode *mini.CThostFtdcQryTradingCodeField, RequestID int) int {
	if api.apiPtr == nil {
		slog.Error(
			"thost trader api not initialized",
			slog.String("caller", "ReqQryTradingCode"),
		)
		return -255
	}

	slog.Info(
		"executing thost trader api",
		slog.String("caller", "ReqQryTradingCode"),
	)

	rtn := C.CallReqQryTradingCode(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQryTradingCode,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQryTradingCodeField)(unsafe.Pointer(QryTradingCode)), C.int(RequestID),
	)

	slog.Info(
		"thost trader api executed",
		slog.String("caller", "ReqQryTradingCode"),
	)

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQryInstrumentMarginRate(QryInstrumentMarginRate *mini.CThostFtdcQryInstrumentMarginRateField, RequestID int) int {
	if api.apiPtr == nil {
		slog.Error(
			"thost trader api not initialized",
			slog.String("caller", "ReqQryInstrumentMarginRate"),
		)
		return -255
	}

	slog.Info(
		"executing thost trader api",
		slog.String("caller", "ReqQryInstrumentMarginRate"),
	)

	rtn := C.CallReqQryInstrumentMarginRate(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQryInstrumentMarginRate,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQryInstrumentMarginRateField)(unsafe.Pointer(QryInstrumentMarginRate)), C.int(RequestID),
	)

	slog.Info(
		"thost trader api executed",
		slog.String("caller", "ReqQryInstrumentMarginRate"),
	)

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQryInstrumentCommissionRate(QryInstrumentCommissionRate *mini.CThostFtdcQryInstrumentCommissionRateField, RequestID int) int {
	if api.apiPtr == nil {
		slog.Error(
			"thost trader api not initialized",
			slog.String("caller", "ReqQryInstrumentCommissionRate"),
		)
		return -255
	}

	slog.Info(
		"executing thost trader api",
		slog.String("caller", "ReqQryInstrumentCommissionRate"),
	)

	rtn := C.CallReqQryInstrumentCommissionRate(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQryInstrumentCommissionRate,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQryInstrumentCommissionRateField)(unsafe.Pointer(QryInstrumentCommissionRate)), C.int(RequestID),
	)

	slog.Info(
		"thost trader api executed",
		slog.String("caller", "ReqQryInstrumentCommissionRate"),
	)

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQryExchange(QryExchange *mini.CThostFtdcQryExchangeField, RequestID int) int {
	if api.apiPtr == nil {
		slog.Error(
			"thost trader api not initialized",
			slog.String("caller", "ReqQryExchange"),
		)
		return -255
	}

	slog.Info(
		"executing thost trader api",
		slog.String("caller", "ReqQryExchange"),
	)

	rtn := C.CallReqQryExchange(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQryExchange,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQryExchangeField)(unsafe.Pointer(QryExchange)), C.int(RequestID),
	)

	slog.Info(
		"thost trader api executed",
		slog.String("caller", "ReqQryExchange"),
	)

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQryProduct(QryProduct *mini.CThostFtdcQryProductField, RequestID int) int {
	if api.apiPtr == nil {
		slog.Error(
			"thost trader api not initialized",
			slog.String("caller", "ReqQryProduct"),
		)
		return -255
	}

	slog.Info(
		"executing thost trader api",
		slog.String("caller", "ReqQryProduct"),
	)

	rtn := C.CallReqQryProduct(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQryProduct,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQryProductField)(unsafe.Pointer(QryProduct)), C.int(RequestID),
	)

	slog.Info(
		"thost trader api executed",
		slog.String("caller", "ReqQryProduct"),
	)

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQryInstrument(QryInstrument *mini.CThostFtdcQryInstrumentField, RequestID int) int {
	if api.apiPtr == nil {
		slog.Error(
			"thost trader api not initialized",
			slog.String("caller", "ReqQryInstrument"),
		)
		return -255
	}

	slog.Info(
		"executing thost trader api",
		slog.String("caller", "ReqQryInstrument"),
	)

	rtn := C.CallReqQryInstrument(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQryInstrument,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQryInstrumentField)(unsafe.Pointer(QryInstrument)), C.int(RequestID),
	)

	slog.Info(
		"thost trader api executed",
		slog.String("caller", "ReqQryInstrument"),
	)

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQryCombInstrument(QryCombInstrument *mini.CThostFtdcQryCombInstrumentField, RequestID int) int {
	if api.apiPtr == nil {
		slog.Error(
			"thost trader api not initialized",
			slog.String("caller", "ReqQryCombInstrument"),
		)
		return -255
	}

	slog.Info(
		"executing thost trader api",
		slog.String("caller", "ReqQryCombInstrument"),
	)

	rtn := C.CallReqQryCombInstrument(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQryCombInstrument,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQryCombInstrumentField)(unsafe.Pointer(QryCombInstrument)), C.int(RequestID),
	)

	slog.Info(
		"thost trader api executed",
		slog.String("caller", "ReqQryCombInstrument"),
	)

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQryRCAMSInvestorProdMargin(QryRCAMSInvestorProdMargin *mini.CThostFtdcQryRCAMSInvestorProdMarginField, RequestID int) int {
	if api.apiPtr == nil {
		slog.Error(
			"thost trader api not initialized",
			slog.String("caller", "ReqQryRCAMSInvestorProdMargin"),
		)
		return -255
	}

	slog.Info(
		"executing thost trader api",
		slog.String("caller", "ReqQryRCAMSInvestorProdMargin"),
	)

	rtn := C.CallReqQryRCAMSInvestorProdMargin(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQryRCAMSInvestorProdMargin,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQryRCAMSInvestorProdMarginField)(unsafe.Pointer(QryRCAMSInvestorProdMargin)), C.int(RequestID),
	)

	slog.Info(
		"thost trader api executed",
		slog.String("caller", "ReqQryRCAMSInvestorProdMargin"),
	)

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQryRCAMSInvestorCombPosition(QryRCAMSInvestorCombPosition *mini.CThostFtdcQryRCAMSInvestorCombPositionField, RequestID int) int {
	if api.apiPtr == nil {
		slog.Error(
			"thost trader api not initialized",
			slog.String("caller", "ReqQryRCAMSInvestorCombPosition"),
		)
		return -255
	}

	slog.Info(
		"executing thost trader api",
		slog.String("caller", "ReqQryRCAMSInvestorCombPosition"),
	)

	rtn := C.CallReqQryRCAMSInvestorCombPosition(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQryRCAMSInvestorCombPosition,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQryRCAMSInvestorCombPositionField)(unsafe.Pointer(QryRCAMSInvestorCombPosition)), C.int(RequestID),
	)

	slog.Info(
		"thost trader api executed",
		slog.String("caller", "ReqQryRCAMSInvestorCombPosition"),
	)

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQryInvestorPositionForComb(QryIPForComb *mini.CThostFtdcQryInvestorPositionForCombField, RequestID int) int {
	if api.apiPtr == nil {
		slog.Error(
			"thost trader api not initialized",
			slog.String("caller", "ReqQryInvestorPositionForComb"),
		)
		return -255
	}

	slog.Info(
		"executing thost trader api",
		slog.String("caller", "ReqQryInvestorPositionForComb"),
	)

	rtn := C.CallReqQryInvestorPositionForComb(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQryInvestorPositionForComb,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQryInvestorPositionForCombField)(unsafe.Pointer(QryIPForComb)), C.int(RequestID),
	)

	slog.Info(
		"thost trader api executed",
		slog.String("caller", "ReqQryInvestorPositionForComb"),
	)

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQryCombAction(QryCombAction *mini.CThostFtdcQryCombActionField, RequestID int) int {
	if api.apiPtr == nil {
		slog.Error(
			"thost trader api not initialized",
			slog.String("caller", "ReqQryCombAction"),
		)
		return -255
	}

	slog.Info(
		"executing thost trader api",
		slog.String("caller", "ReqQryCombAction"),
	)

	rtn := C.CallReqQryCombAction(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQryCombAction,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQryCombActionField)(unsafe.Pointer(QryCombAction)), C.int(RequestID),
	)

	slog.Info(
		"thost trader api executed",
		slog.String("caller", "ReqQryCombAction"),
	)

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQryDepthMarketData(QryDepthMarketData *mini.CThostFtdcQryDepthMarketDataField, RequestID int) int {
	if api.apiPtr == nil {
		slog.Error(
			"thost trader api not initialized",
			slog.String("caller", "ReqQryDepthMarketData"),
		)
		return -255
	}

	slog.Info(
		"executing thost trader api",
		slog.String("caller", "ReqQryDepthMarketData"),
	)

	rtn := C.CallReqQryDepthMarketData(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQryDepthMarketData,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQryDepthMarketDataField)(unsafe.Pointer(QryDepthMarketData)), C.int(RequestID),
	)

	slog.Info(
		"thost trader api executed",
		slog.String("caller", "ReqQryDepthMarketData"),
	)

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQryOptionSelfClose(QryOptionSelfClose *mini.CThostFtdcQryOptionSelfCloseField, RequestID int) int {
	if api.apiPtr == nil {
		slog.Error(
			"thost trader api not initialized",
			slog.String("caller", "ReqQryOptionSelfClose"),
		)
		return -255
	}

	slog.Info(
		"executing thost trader api",
		slog.String("caller", "ReqQryOptionSelfClose"),
	)

	rtn := C.CallReqQryOptionSelfClose(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQryOptionSelfClose,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQryOptionSelfCloseField)(unsafe.Pointer(QryOptionSelfClose)), C.int(RequestID),
	)

	slog.Info(
		"thost trader api executed",
		slog.String("caller", "ReqQryOptionSelfClose"),
	)

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQryInstrumentStatus(QryInstrumentStatus *mini.CThostFtdcQryInstrumentStatusField, RequestID int) int {
	if api.apiPtr == nil {
		slog.Error(
			"thost trader api not initialized",
			slog.String("caller", "ReqQryInstrumentStatus"),
		)
		return -255
	}

	slog.Info(
		"executing thost trader api",
		slog.String("caller", "ReqQryInstrumentStatus"),
	)

	rtn := C.CallReqQryInstrumentStatus(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQryInstrumentStatus,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQryInstrumentStatusField)(unsafe.Pointer(QryInstrumentStatus)), C.int(RequestID),
	)

	slog.Info(
		"thost trader api executed",
		slog.String("caller", "ReqQryInstrumentStatus"),
	)

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQryInvestorPositionDetail(QryInvestorPositionDetail *mini.CThostFtdcQryInvestorPositionDetailField, RequestID int) int {
	if api.apiPtr == nil {
		slog.Error(
			"thost trader api not initialized",
			slog.String("caller", "ReqQryInvestorPositionDetail"),
		)
		return -255
	}

	slog.Info(
		"executing thost trader api",
		slog.String("caller", "ReqQryInvestorPositionDetail"),
	)

	rtn := C.CallReqQryInvestorPositionDetail(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQryInvestorPositionDetail,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQryInvestorPositionDetailField)(unsafe.Pointer(QryInvestorPositionDetail)), C.int(RequestID),
	)

	slog.Info(
		"thost trader api executed",
		slog.String("caller", "ReqQryInvestorPositionDetail"),
	)

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQryExchangeMarginRate(QryExchangeMarginRate *mini.CThostFtdcQryExchangeMarginRateField, RequestID int) int {
	if api.apiPtr == nil {
		slog.Error(
			"thost trader api not initialized",
			slog.String("caller", "ReqQryExchangeMarginRate"),
		)
		return -255
	}

	slog.Info(
		"executing thost trader api",
		slog.String("caller", "ReqQryExchangeMarginRate"),
	)

	rtn := C.CallReqQryExchangeMarginRate(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQryExchangeMarginRate,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQryExchangeMarginRateField)(unsafe.Pointer(QryExchangeMarginRate)), C.int(RequestID),
	)

	slog.Info(
		"thost trader api executed",
		slog.String("caller", "ReqQryExchangeMarginRate"),
	)

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQryExchangeMarginRateAdjust(QryExchangeMarginRateAdjust *mini.CThostFtdcQryExchangeMarginRateAdjustField, RequestID int) int {
	if api.apiPtr == nil {
		slog.Error(
			"thost trader api not initialized",
			slog.String("caller", "ReqQryExchangeMarginRateAdjust"),
		)
		return -255
	}

	slog.Info(
		"executing thost trader api",
		slog.String("caller", "ReqQryExchangeMarginRateAdjust"),
	)

	rtn := C.CallReqQryExchangeMarginRateAdjust(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQryExchangeMarginRateAdjust,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQryExchangeMarginRateAdjustField)(unsafe.Pointer(QryExchangeMarginRateAdjust)), C.int(RequestID),
	)

	slog.Info(
		"thost trader api executed",
		slog.String("caller", "ReqQryExchangeMarginRateAdjust"),
	)

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQryOptionInstrTradeCost(QryOptionInstrTradeCost *mini.CThostFtdcQryOptionInstrTradeCostField, RequestID int) int {
	if api.apiPtr == nil {
		slog.Error(
			"thost trader api not initialized",
			slog.String("caller", "ReqQryOptionInstrTradeCost"),
		)
		return -255
	}

	slog.Info(
		"executing thost trader api",
		slog.String("caller", "ReqQryOptionInstrTradeCost"),
	)

	rtn := C.CallReqQryOptionInstrTradeCost(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQryOptionInstrTradeCost,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQryOptionInstrTradeCostField)(unsafe.Pointer(QryOptionInstrTradeCost)), C.int(RequestID),
	)

	slog.Info(
		"thost trader api executed",
		slog.String("caller", "ReqQryOptionInstrTradeCost"),
	)

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQryOptionInstrCommRate(QryOptionInstrCommRate *mini.CThostFtdcQryOptionInstrCommRateField, RequestID int) int {
	if api.apiPtr == nil {
		slog.Error(
			"thost trader api not initialized",
			slog.String("caller", "ReqQryOptionInstrCommRate"),
		)
		return -255
	}

	slog.Info(
		"executing thost trader api",
		slog.String("caller", "ReqQryOptionInstrCommRate"),
	)

	rtn := C.CallReqQryOptionInstrCommRate(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQryOptionInstrCommRate,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQryOptionInstrCommRateField)(unsafe.Pointer(QryOptionInstrCommRate)), C.int(RequestID),
	)

	slog.Info(
		"thost trader api executed",
		slog.String("caller", "ReqQryOptionInstrCommRate"),
	)

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQryExecOrder(QryExecOrder *mini.CThostFtdcQryExecOrderField, RequestID int) int {
	if api.apiPtr == nil {
		slog.Error(
			"thost trader api not initialized",
			slog.String("caller", "ReqQryExecOrder"),
		)
		return -255
	}

	slog.Info(
		"executing thost trader api",
		slog.String("caller", "ReqQryExecOrder"),
	)

	rtn := C.CallReqQryExecOrder(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQryExecOrder,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQryExecOrderField)(unsafe.Pointer(QryExecOrder)), C.int(RequestID),
	)

	slog.Info(
		"thost trader api executed",
		slog.String("caller", "ReqQryExecOrder"),
	)

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQryForQuote(QryForQuote *mini.CThostFtdcQryForQuoteField, RequestID int) int {
	if api.apiPtr == nil {
		slog.Error(
			"thost trader api not initialized",
			slog.String("caller", "ReqQryForQuote"),
		)
		return -255
	}

	slog.Info(
		"executing thost trader api",
		slog.String("caller", "ReqQryForQuote"),
	)

	rtn := C.CallReqQryForQuote(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQryForQuote,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQryForQuoteField)(unsafe.Pointer(QryForQuote)), C.int(RequestID),
	)

	slog.Info(
		"thost trader api executed",
		slog.String("caller", "ReqQryForQuote"),
	)

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQryQuote(QryQuote *mini.CThostFtdcQryQuoteField, RequestID int) int {
	if api.apiPtr == nil {
		slog.Error(
			"thost trader api not initialized",
			slog.String("caller", "ReqQryQuote"),
		)
		return -255
	}

	slog.Info(
		"executing thost trader api",
		slog.String("caller", "ReqQryQuote"),
	)

	rtn := C.CallReqQryQuote(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQryQuote,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQryQuoteField)(unsafe.Pointer(QryQuote)), C.int(RequestID),
	)

	slog.Info(
		"thost trader api executed",
		slog.String("caller", "ReqQryQuote"),
	)

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQryInstrumentOrderCommRate(QryInstrumentOrderCommRate *mini.CThostFtdcQryInstrumentOrderCommRateField, RequestID int) int {
	if api.apiPtr == nil {
		slog.Error(
			"thost trader api not initialized",
			slog.String("caller", "ReqQryInstrumentOrderCommRate"),
		)
		return -255
	}

	slog.Info(
		"executing thost trader api",
		slog.String("caller", "ReqQryInstrumentOrderCommRate"),
	)

	rtn := C.CallReqQryInstrumentOrderCommRate(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQryInstrumentOrderCommRate,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQryInstrumentOrderCommRateField)(unsafe.Pointer(QryInstrumentOrderCommRate)), C.int(RequestID),
	)

	slog.Info(
		"thost trader api executed",
		slog.String("caller", "ReqQryInstrumentOrderCommRate"),
	)

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQryForQuoteParam(QryForQuoteParam *mini.CThostFtdcQryForQuoteParamField, RequestID int) int {
	if api.apiPtr == nil {
		slog.Error(
			"thost trader api not initialized",
			slog.String("caller", "ReqQryForQuoteParam"),
		)
		return -255
	}

	slog.Info(
		"executing thost trader api",
		slog.String("caller", "ReqQryForQuoteParam"),
	)

	rtn := C.CallReqQryForQuoteParam(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQryForQuoteParam,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQryForQuoteParamField)(unsafe.Pointer(QryForQuoteParam)), C.int(RequestID),
	)

	slog.Info(
		"thost trader api executed",
		slog.String("caller", "ReqQryForQuoteParam"),
	)

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQryTraderOffer(QryTraderOffer *mini.CThostFtdcQryTraderOfferField, RequestID int) int {
	if api.apiPtr == nil {
		slog.Error(
			"thost trader api not initialized",
			slog.String("caller", "ReqQryTraderOffer"),
		)
		return -255
	}

	slog.Info(
		"executing thost trader api",
		slog.String("caller", "ReqQryTraderOffer"),
	)

	rtn := C.CallReqQryTraderOffer(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQryTraderOffer,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQryTraderOfferField)(unsafe.Pointer(QryTraderOffer)), C.int(RequestID),
	)

	slog.Info(
		"thost trader api executed",
		slog.String("caller", "ReqQryTraderOffer"),
	)

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQryInvestorProdSPBMDetail(QryInvestorProdSPBMDetail *mini.CThostFtdcQryInvestorProdSPBMDetailField, RequestID int) int {
	if api.apiPtr == nil {
		slog.Error(
			"thost trader api not initialized",
			slog.String("caller", "ReqQryInvestorProdSPBMDetail"),
		)
		return -255
	}

	slog.Info(
		"executing thost trader api",
		slog.String("caller", "ReqQryInvestorProdSPBMDetail"),
	)

	rtn := C.CallReqQryInvestorProdSPBMDetail(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQryInvestorProdSPBMDetail,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQryInvestorProdSPBMDetailField)(unsafe.Pointer(QryInvestorProdSPBMDetail)), C.int(RequestID),
	)

	slog.Info(
		"thost trader api executed",
		slog.String("caller", "ReqQryInvestorProdSPBMDetail"),
	)

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQrySPMMInvestorCommodityGroupMargin(QrySPMMInvestorCommodityGroupMargin *mini.CThostFtdcQrySPMMInvestorCommodityGroupMarginField, RequestID int) int {
	if api.apiPtr == nil {
		slog.Error(
			"thost trader api not initialized",
			slog.String("caller", "ReqQrySPMMInvestorCommodityGroupMargin"),
		)
		return -255
	}

	slog.Info(
		"executing thost trader api",
		slog.String("caller", "ReqQrySPMMInvestorCommodityGroupMargin"),
	)

	rtn := C.CallReqQrySPMMInvestorCommodityGroupMargin(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQrySPMMInvestorCommodityGroupMargin,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQrySPMMInvestorCommodityGroupMarginField)(unsafe.Pointer(QrySPMMInvestorCommodityGroupMargin)), C.int(RequestID),
	)

	slog.Info(
		"thost trader api executed",
		slog.String("caller", "ReqQrySPMMInvestorCommodityGroupMargin"),
	)

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQryRULEInvestorProdMargin(QryRULEInvestorProdMargin *mini.CThostFtdcQryRULEInvestorProdMarginField, RequestID int) int {
	if api.apiPtr == nil {
		slog.Error(
			"thost trader api not initialized",
			slog.String("caller", "ReqQryRULEInvestorProdMargin"),
		)
		return -255
	}

	slog.Info(
		"executing thost trader api",
		slog.String("caller", "ReqQryRULEInvestorProdMargin"),
	)

	rtn := C.CallReqQryRULEInvestorProdMargin(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQryRULEInvestorProdMargin,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQryRULEInvestorProdMarginField)(unsafe.Pointer(QryRULEInvestorProdMargin)), C.int(RequestID),
	)

	slog.Info(
		"thost trader api executed",
		slog.String("caller", "ReqQryRULEInvestorProdMargin"),
	)

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQryControlParam(QryControlParam *mini.CThostFtdcQryControlParamField, RequestID int) int {
	if api.apiPtr == nil {
		slog.Error(
			"thost trader api not initialized",
			slog.String("caller", "ReqQryControlParam"),
		)
		return -255
	}

	slog.Info(
		"executing thost trader api",
		slog.String("caller", "ReqQryControlParam"),
	)

	rtn := C.CallReqQryControlParam(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQryControlParam,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQryControlParamField)(unsafe.Pointer(QryControlParam)), C.int(RequestID),
	)

	slog.Info(
		"thost trader api executed",
		slog.String("caller", "ReqQryControlParam"),
	)

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQryOffsetSetting(QryOffsetSetting *mini.CThostFtdcQryOffsetSettingField, RequestID int) int {
	if api.apiPtr == nil {
		slog.Error(
			"thost trader api not initialized",
			slog.String("caller", "ReqQryOffsetSetting"),
		)
		return -255
	}

	slog.Info(
		"executing thost trader api",
		slog.String("caller", "ReqQryOffsetSetting"),
	)

	rtn := C.CallReqQryOffsetSetting(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQryOffsetSetting,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQryOffsetSettingField)(unsafe.Pointer(QryOffsetSetting)), C.int(RequestID),
	)

	slog.Info(
		"thost trader api executed",
		slog.String("caller", "ReqQryOffsetSetting"),
	)

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqUserPasswordUpdate(UserPasswordUpdate *mini.CThostFtdcUserPasswordUpdateField, RequestID int) int {
	if api.apiPtr == nil {
		slog.Error(
			"thost trader api not initialized",
			slog.String("caller", "ReqUserPasswordUpdate"),
		)
		return -255
	}

	slog.Info(
		"executing thost trader api",
		slog.String("caller", "ReqUserPasswordUpdate"),
	)

	rtn := C.CallReqUserPasswordUpdate(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqUserPasswordUpdate,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcUserPasswordUpdateField)(unsafe.Pointer(UserPasswordUpdate)), C.int(RequestID),
	)

	slog.Info(
		"thost trader api executed",
		slog.String("caller", "ReqUserPasswordUpdate"),
	)

	return int(rtn)
}
