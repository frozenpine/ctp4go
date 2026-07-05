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

	"github.com/frozenpine/ctp4go/thost"
	"github.com/frozenpine/ctp4go/thost/mini"
	"github.com/frozenpine/ctp4go/thost/mini/types"
)

// var (
//     // 确保Api封装完整实现了thost中的接口签名
//     _ mini.TraderApi = &ThostFtdcTraderApi{}
// )

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
			types.DecodeGBK(([]byte)(C.GoString(msg))),
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
	slog.Info("executing thost trader api Release")

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

	slog.Info("thost trader api Release executed")

}

func (api *ThostFtdcTraderApi) Init(Continuous bool) {
	slog.Info("executing thost trader api Init")

	C.CallInit(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_Init,
		unsafe.Pointer(api.apiPtr),
		C.bool(Continuous),
	)

	slog.Info("thost trader api Init executed")

}

func (api *ThostFtdcTraderApi) Join() int {
	slog.Info("executing thost trader api Join")

	rtn := C.CallJoin(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_Join,
		unsafe.Pointer(api.apiPtr),
	)

	slog.Info("thost trader api Join executed")

	return int(rtn)
}

func (api *ThostFtdcTraderApi) GetTradingDay() string {
	slog.Info("executing thost trader api GetTradingDay")

	rtn := C.CallGetTradingDay(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_GetTradingDay,
		unsafe.Pointer(api.apiPtr),
	)

	slog.Info("thost trader api GetTradingDay executed")

	return C.GoString(rtn)
}

func (api *ThostFtdcTraderApi) RegisterFront(FrontAddress string) {
	slog.Info("executing thost trader api RegisterFront")

	pszFrontAddress := C.CString(FrontAddress)
	defer C.free(unsafe.Pointer(pszFrontAddress))

	C.CallRegisterFront(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_RegisterFront,
		unsafe.Pointer(api.apiPtr),
		pszFrontAddress,
	)

	slog.Info("thost trader api RegisterFront executed")

}

func (api *ThostFtdcTraderApi) RegisterSpi(Spi mini.TraderSpi) {
	slog.Info("executing thost trader api RegisterSpi")

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

	slog.Info("thost trader api RegisterSpi executed")

}

func (api *ThostFtdcTraderApi) SubscribePrivateTopic(ResumeType int) {
	slog.Info("executing thost trader api SubscribePrivateTopic")

	C.CallSubscribePrivateTopic(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_SubscribePrivateTopic,
		unsafe.Pointer(api.apiPtr),
		C.int(ResumeType),
	)

	slog.Info("thost trader api SubscribePrivateTopic executed")

}

func (api *ThostFtdcTraderApi) SubscribePublicTopic(ResumeType int) {
	slog.Info("executing thost trader api SubscribePublicTopic")

	C.CallSubscribePublicTopic(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_SubscribePublicTopic,
		unsafe.Pointer(api.apiPtr),
		C.int(ResumeType),
	)

	slog.Info("thost trader api SubscribePublicTopic executed")

}

func (api *ThostFtdcTraderApi) SubscribeFlowCtrlWarning(TraderID ...string) int {
	slog.Info("executing thost trader api SubscribeFlowCtrlWarning")

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

	slog.Info("thost trader api SubscribeFlowCtrlWarning executed")

	return int(rtn)
}

func (api *ThostFtdcTraderApi) UnSubscribeFlowCtrlWarning(TraderID ...string) int {
	slog.Info("executing thost trader api UnSubscribeFlowCtrlWarning")

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

	slog.Info("thost trader api UnSubscribeFlowCtrlWarning executed")

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqAuthenticate(ReqAuthenticateField *mini.CThostFtdcReqAuthenticateField, RequestID int) int {
	slog.Info("executing thost trader api ReqAuthenticate")

	rtn := C.CallReqAuthenticate(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqAuthenticate,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcReqAuthenticateField)(unsafe.Pointer(ReqAuthenticateField)), C.int(RequestID),
	)

	slog.Info("thost trader api ReqAuthenticate executed")

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqUserLogin(ReqUserLoginField *mini.CThostFtdcReqUserLoginField, RequestID int) int {
	slog.Info("executing thost trader api ReqUserLogin")

	rtn := C.CallReqUserLogin(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqUserLogin,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcReqUserLoginField)(unsafe.Pointer(ReqUserLoginField)), C.int(RequestID),
	)

	slog.Info("thost trader api ReqUserLogin executed")

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqUserLoginEncrypt(ReqUserLoginField *mini.CThostFtdcReqUserLoginField, RequestID int) int {
	slog.Info("executing thost trader api ReqUserLoginEncrypt")

	rtn := C.CallReqUserLoginEncrypt(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqUserLoginEncrypt,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcReqUserLoginField)(unsafe.Pointer(ReqUserLoginField)), C.int(RequestID),
	)

	slog.Info("thost trader api ReqUserLoginEncrypt executed")

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqUserLogout(UserLogout *mini.CThostFtdcUserLogoutField, RequestID int) int {
	slog.Info("executing thost trader api ReqUserLogout")

	rtn := C.CallReqUserLogout(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqUserLogout,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcUserLogoutField)(unsafe.Pointer(UserLogout)), C.int(RequestID),
	)

	slog.Info("thost trader api ReqUserLogout executed")

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqOrderInsert(InputOrder *mini.CThostFtdcInputOrderField, RequestID int) int {
	slog.Info("executing thost trader api ReqOrderInsert")

	rtn := C.CallReqOrderInsert(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqOrderInsert,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcInputOrderField)(unsafe.Pointer(InputOrder)), C.int(RequestID),
	)

	slog.Info("thost trader api ReqOrderInsert executed")

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqOrderAction(InputOrderAction *mini.CThostFtdcInputOrderActionField, RequestID int) int {
	slog.Info("executing thost trader api ReqOrderAction")

	rtn := C.CallReqOrderAction(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqOrderAction,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcInputOrderActionField)(unsafe.Pointer(InputOrderAction)), C.int(RequestID),
	)

	slog.Info("thost trader api ReqOrderAction executed")

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqMKBatchOrderAction(MKInputOrderAction *mini.CThostFtdcMKInputOrderActionField, RequestID int) int {
	slog.Info("executing thost trader api ReqMKBatchOrderAction")

	rtn := C.CallReqMKBatchOrderAction(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqMKBatchOrderAction,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcMKInputOrderActionField)(unsafe.Pointer(MKInputOrderAction)), C.int(RequestID),
	)

	slog.Info("thost trader api ReqMKBatchOrderAction executed")

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqExecOrderInsert(InputExecOrder *mini.CThostFtdcInputExecOrderField, RequestID int) int {
	slog.Info("executing thost trader api ReqExecOrderInsert")

	rtn := C.CallReqExecOrderInsert(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqExecOrderInsert,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcInputExecOrderField)(unsafe.Pointer(InputExecOrder)), C.int(RequestID),
	)

	slog.Info("thost trader api ReqExecOrderInsert executed")

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqExecOrderAction(InputExecOrderAction *mini.CThostFtdcInputExecOrderActionField, RequestID int) int {
	slog.Info("executing thost trader api ReqExecOrderAction")

	rtn := C.CallReqExecOrderAction(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqExecOrderAction,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcInputExecOrderActionField)(unsafe.Pointer(InputExecOrderAction)), C.int(RequestID),
	)

	slog.Info("thost trader api ReqExecOrderAction executed")

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqForQuoteInsert(InputForQuote *mini.CThostFtdcInputForQuoteField, RequestID int) int {
	slog.Info("executing thost trader api ReqForQuoteInsert")

	rtn := C.CallReqForQuoteInsert(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqForQuoteInsert,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcInputForQuoteField)(unsafe.Pointer(InputForQuote)), C.int(RequestID),
	)

	slog.Info("thost trader api ReqForQuoteInsert executed")

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQuoteInsert(InputQuote *mini.CThostFtdcInputQuoteField, RequestID int) int {
	slog.Info("executing thost trader api ReqQuoteInsert")

	rtn := C.CallReqQuoteInsert(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQuoteInsert,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcInputQuoteField)(unsafe.Pointer(InputQuote)), C.int(RequestID),
	)

	slog.Info("thost trader api ReqQuoteInsert executed")

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQuoteAction(InputQuoteAction *mini.CThostFtdcInputQuoteActionField, RequestID int) int {
	slog.Info("executing thost trader api ReqQuoteAction")

	rtn := C.CallReqQuoteAction(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQuoteAction,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcInputQuoteActionField)(unsafe.Pointer(InputQuoteAction)), C.int(RequestID),
	)

	slog.Info("thost trader api ReqQuoteAction executed")

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqBatchOrderAction(InputBatchOrderAction *mini.CThostFtdcInputBatchOrderActionField, RequestID int) int {
	slog.Info("executing thost trader api ReqBatchOrderAction")

	rtn := C.CallReqBatchOrderAction(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqBatchOrderAction,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcInputBatchOrderActionField)(unsafe.Pointer(InputBatchOrderAction)), C.int(RequestID),
	)

	slog.Info("thost trader api ReqBatchOrderAction executed")

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqOptionSelfCloseInsert(InputOptionSelfClose *mini.CThostFtdcInputOptionSelfCloseField, RequestID int) int {
	slog.Info("executing thost trader api ReqOptionSelfCloseInsert")

	rtn := C.CallReqOptionSelfCloseInsert(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqOptionSelfCloseInsert,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcInputOptionSelfCloseField)(unsafe.Pointer(InputOptionSelfClose)), C.int(RequestID),
	)

	slog.Info("thost trader api ReqOptionSelfCloseInsert executed")

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqOptionSelfCloseAction(InputOptionSelfCloseAction *mini.CThostFtdcInputOptionSelfCloseActionField, RequestID int) int {
	slog.Info("executing thost trader api ReqOptionSelfCloseAction")

	rtn := C.CallReqOptionSelfCloseAction(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqOptionSelfCloseAction,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcInputOptionSelfCloseActionField)(unsafe.Pointer(InputOptionSelfCloseAction)), C.int(RequestID),
	)

	slog.Info("thost trader api ReqOptionSelfCloseAction executed")

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqCombActionInsert(InputCombAction *mini.CThostFtdcInputCombActionField, RequestID int) int {
	slog.Info("executing thost trader api ReqCombActionInsert")

	rtn := C.CallReqCombActionInsert(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqCombActionInsert,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcInputCombActionField)(unsafe.Pointer(InputCombAction)), C.int(RequestID),
	)

	slog.Info("thost trader api ReqCombActionInsert executed")

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqSubscribeFundChange(RequestID int) int {
	slog.Info("executing thost trader api ReqSubscribeFundChange")

	rtn := C.CallReqSubscribeFundChange(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqSubscribeFundChange,
		unsafe.Pointer(api.apiPtr),
		C.int(RequestID),
	)

	slog.Info("thost trader api ReqSubscribeFundChange executed")

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqUnSubscribeFundChange(RequestID int) int {
	slog.Info("executing thost trader api ReqUnSubscribeFundChange")

	rtn := C.CallReqUnSubscribeFundChange(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqUnSubscribeFundChange,
		unsafe.Pointer(api.apiPtr),
		C.int(RequestID),
	)

	slog.Info("thost trader api ReqUnSubscribeFundChange executed")

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqOffsetSetting(InputOffsetSetting *mini.CThostFtdcInputOffsetSettingField, RequestID int) int {
	slog.Info("executing thost trader api ReqOffsetSetting")

	rtn := C.CallReqOffsetSetting(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqOffsetSetting,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcInputOffsetSettingField)(unsafe.Pointer(InputOffsetSetting)), C.int(RequestID),
	)

	slog.Info("thost trader api ReqOffsetSetting executed")

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqCancelOffsetSetting(InputOffsetSetting *mini.CThostFtdcInputOffsetSettingField, RequestID int) int {
	slog.Info("executing thost trader api ReqCancelOffsetSetting")

	rtn := C.CallReqCancelOffsetSetting(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqCancelOffsetSetting,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcInputOffsetSettingField)(unsafe.Pointer(InputOffsetSetting)), C.int(RequestID),
	)

	slog.Info("thost trader api ReqCancelOffsetSetting executed")

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQryOrder(QryOrder *mini.CThostFtdcQryOrderField, RequestID int) int {
	slog.Info("executing thost trader api ReqQryOrder")

	rtn := C.CallReqQryOrder(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQryOrder,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQryOrderField)(unsafe.Pointer(QryOrder)), C.int(RequestID),
	)

	slog.Info("thost trader api ReqQryOrder executed")

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQryTrade(QryTrade *mini.CThostFtdcQryTradeField, RequestID int) int {
	slog.Info("executing thost trader api ReqQryTrade")

	rtn := C.CallReqQryTrade(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQryTrade,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQryTradeField)(unsafe.Pointer(QryTrade)), C.int(RequestID),
	)

	slog.Info("thost trader api ReqQryTrade executed")

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQryInvestorPosition(QryInvestorPosition *mini.CThostFtdcQryInvestorPositionField, RequestID int) int {
	slog.Info("executing thost trader api ReqQryInvestorPosition")

	rtn := C.CallReqQryInvestorPosition(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQryInvestorPosition,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQryInvestorPositionField)(unsafe.Pointer(QryInvestorPosition)), C.int(RequestID),
	)

	slog.Info("thost trader api ReqQryInvestorPosition executed")

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQryTradingAccount(QryTradingAccount *mini.CThostFtdcQryTradingAccountField, RequestID int) int {
	slog.Info("executing thost trader api ReqQryTradingAccount")

	rtn := C.CallReqQryTradingAccount(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQryTradingAccount,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQryTradingAccountField)(unsafe.Pointer(QryTradingAccount)), C.int(RequestID),
	)

	slog.Info("thost trader api ReqQryTradingAccount executed")

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQryInvestor(QryInvestor *mini.CThostFtdcQryInvestorField, RequestID int) int {
	slog.Info("executing thost trader api ReqQryInvestor")

	rtn := C.CallReqQryInvestor(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQryInvestor,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQryInvestorField)(unsafe.Pointer(QryInvestor)), C.int(RequestID),
	)

	slog.Info("thost trader api ReqQryInvestor executed")

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQryTradingCode(QryTradingCode *mini.CThostFtdcQryTradingCodeField, RequestID int) int {
	slog.Info("executing thost trader api ReqQryTradingCode")

	rtn := C.CallReqQryTradingCode(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQryTradingCode,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQryTradingCodeField)(unsafe.Pointer(QryTradingCode)), C.int(RequestID),
	)

	slog.Info("thost trader api ReqQryTradingCode executed")

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQryInstrumentMarginRate(QryInstrumentMarginRate *mini.CThostFtdcQryInstrumentMarginRateField, RequestID int) int {
	slog.Info("executing thost trader api ReqQryInstrumentMarginRate")

	rtn := C.CallReqQryInstrumentMarginRate(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQryInstrumentMarginRate,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQryInstrumentMarginRateField)(unsafe.Pointer(QryInstrumentMarginRate)), C.int(RequestID),
	)

	slog.Info("thost trader api ReqQryInstrumentMarginRate executed")

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQryInstrumentCommissionRate(QryInstrumentCommissionRate *mini.CThostFtdcQryInstrumentCommissionRateField, RequestID int) int {
	slog.Info("executing thost trader api ReqQryInstrumentCommissionRate")

	rtn := C.CallReqQryInstrumentCommissionRate(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQryInstrumentCommissionRate,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQryInstrumentCommissionRateField)(unsafe.Pointer(QryInstrumentCommissionRate)), C.int(RequestID),
	)

	slog.Info("thost trader api ReqQryInstrumentCommissionRate executed")

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQryExchange(QryExchange *mini.CThostFtdcQryExchangeField, RequestID int) int {
	slog.Info("executing thost trader api ReqQryExchange")

	rtn := C.CallReqQryExchange(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQryExchange,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQryExchangeField)(unsafe.Pointer(QryExchange)), C.int(RequestID),
	)

	slog.Info("thost trader api ReqQryExchange executed")

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQryProduct(QryProduct *mini.CThostFtdcQryProductField, RequestID int) int {
	slog.Info("executing thost trader api ReqQryProduct")

	rtn := C.CallReqQryProduct(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQryProduct,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQryProductField)(unsafe.Pointer(QryProduct)), C.int(RequestID),
	)

	slog.Info("thost trader api ReqQryProduct executed")

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQryInstrument(QryInstrument *mini.CThostFtdcQryInstrumentField, RequestID int) int {
	slog.Info("executing thost trader api ReqQryInstrument")

	rtn := C.CallReqQryInstrument(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQryInstrument,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQryInstrumentField)(unsafe.Pointer(QryInstrument)), C.int(RequestID),
	)

	slog.Info("thost trader api ReqQryInstrument executed")

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQryCombInstrument(QryCombInstrument *mini.CThostFtdcQryCombInstrumentField, RequestID int) int {
	slog.Info("executing thost trader api ReqQryCombInstrument")

	rtn := C.CallReqQryCombInstrument(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQryCombInstrument,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQryCombInstrumentField)(unsafe.Pointer(QryCombInstrument)), C.int(RequestID),
	)

	slog.Info("thost trader api ReqQryCombInstrument executed")

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQryRCAMSInvestorProdMargin(QryRCAMSInvestorProdMargin *mini.CThostFtdcQryRCAMSInvestorProdMarginField, RequestID int) int {
	slog.Info("executing thost trader api ReqQryRCAMSInvestorProdMargin")

	rtn := C.CallReqQryRCAMSInvestorProdMargin(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQryRCAMSInvestorProdMargin,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQryRCAMSInvestorProdMarginField)(unsafe.Pointer(QryRCAMSInvestorProdMargin)), C.int(RequestID),
	)

	slog.Info("thost trader api ReqQryRCAMSInvestorProdMargin executed")

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQryRCAMSInvestorCombPosition(QryRCAMSInvestorCombPosition *mini.CThostFtdcQryRCAMSInvestorCombPositionField, RequestID int) int {
	slog.Info("executing thost trader api ReqQryRCAMSInvestorCombPosition")

	rtn := C.CallReqQryRCAMSInvestorCombPosition(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQryRCAMSInvestorCombPosition,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQryRCAMSInvestorCombPositionField)(unsafe.Pointer(QryRCAMSInvestorCombPosition)), C.int(RequestID),
	)

	slog.Info("thost trader api ReqQryRCAMSInvestorCombPosition executed")

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQryInvestorPositionForComb(QryIPForComb *mini.CThostFtdcQryInvestorPositionForCombField, RequestID int) int {
	slog.Info("executing thost trader api ReqQryInvestorPositionForComb")

	rtn := C.CallReqQryInvestorPositionForComb(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQryInvestorPositionForComb,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQryInvestorPositionForCombField)(unsafe.Pointer(QryIPForComb)), C.int(RequestID),
	)

	slog.Info("thost trader api ReqQryInvestorPositionForComb executed")

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQryCombAction(QryCombAction *mini.CThostFtdcQryCombActionField, RequestID int) int {
	slog.Info("executing thost trader api ReqQryCombAction")

	rtn := C.CallReqQryCombAction(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQryCombAction,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQryCombActionField)(unsafe.Pointer(QryCombAction)), C.int(RequestID),
	)

	slog.Info("thost trader api ReqQryCombAction executed")

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQryDepthMarketData(QryDepthMarketData *mini.CThostFtdcQryDepthMarketDataField, RequestID int) int {
	slog.Info("executing thost trader api ReqQryDepthMarketData")

	rtn := C.CallReqQryDepthMarketData(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQryDepthMarketData,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQryDepthMarketDataField)(unsafe.Pointer(QryDepthMarketData)), C.int(RequestID),
	)

	slog.Info("thost trader api ReqQryDepthMarketData executed")

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQryOptionSelfClose(QryOptionSelfClose *mini.CThostFtdcQryOptionSelfCloseField, RequestID int) int {
	slog.Info("executing thost trader api ReqQryOptionSelfClose")

	rtn := C.CallReqQryOptionSelfClose(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQryOptionSelfClose,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQryOptionSelfCloseField)(unsafe.Pointer(QryOptionSelfClose)), C.int(RequestID),
	)

	slog.Info("thost trader api ReqQryOptionSelfClose executed")

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQryInstrumentStatus(QryInstrumentStatus *mini.CThostFtdcQryInstrumentStatusField, RequestID int) int {
	slog.Info("executing thost trader api ReqQryInstrumentStatus")

	rtn := C.CallReqQryInstrumentStatus(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQryInstrumentStatus,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQryInstrumentStatusField)(unsafe.Pointer(QryInstrumentStatus)), C.int(RequestID),
	)

	slog.Info("thost trader api ReqQryInstrumentStatus executed")

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQryInvestorPositionDetail(QryInvestorPositionDetail *mini.CThostFtdcQryInvestorPositionDetailField, RequestID int) int {
	slog.Info("executing thost trader api ReqQryInvestorPositionDetail")

	rtn := C.CallReqQryInvestorPositionDetail(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQryInvestorPositionDetail,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQryInvestorPositionDetailField)(unsafe.Pointer(QryInvestorPositionDetail)), C.int(RequestID),
	)

	slog.Info("thost trader api ReqQryInvestorPositionDetail executed")

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQryExchangeMarginRate(QryExchangeMarginRate *mini.CThostFtdcQryExchangeMarginRateField, RequestID int) int {
	slog.Info("executing thost trader api ReqQryExchangeMarginRate")

	rtn := C.CallReqQryExchangeMarginRate(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQryExchangeMarginRate,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQryExchangeMarginRateField)(unsafe.Pointer(QryExchangeMarginRate)), C.int(RequestID),
	)

	slog.Info("thost trader api ReqQryExchangeMarginRate executed")

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQryExchangeMarginRateAdjust(QryExchangeMarginRateAdjust *mini.CThostFtdcQryExchangeMarginRateAdjustField, RequestID int) int {
	slog.Info("executing thost trader api ReqQryExchangeMarginRateAdjust")

	rtn := C.CallReqQryExchangeMarginRateAdjust(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQryExchangeMarginRateAdjust,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQryExchangeMarginRateAdjustField)(unsafe.Pointer(QryExchangeMarginRateAdjust)), C.int(RequestID),
	)

	slog.Info("thost trader api ReqQryExchangeMarginRateAdjust executed")

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQryOptionInstrTradeCost(QryOptionInstrTradeCost *mini.CThostFtdcQryOptionInstrTradeCostField, RequestID int) int {
	slog.Info("executing thost trader api ReqQryOptionInstrTradeCost")

	rtn := C.CallReqQryOptionInstrTradeCost(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQryOptionInstrTradeCost,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQryOptionInstrTradeCostField)(unsafe.Pointer(QryOptionInstrTradeCost)), C.int(RequestID),
	)

	slog.Info("thost trader api ReqQryOptionInstrTradeCost executed")

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQryOptionInstrCommRate(QryOptionInstrCommRate *mini.CThostFtdcQryOptionInstrCommRateField, RequestID int) int {
	slog.Info("executing thost trader api ReqQryOptionInstrCommRate")

	rtn := C.CallReqQryOptionInstrCommRate(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQryOptionInstrCommRate,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQryOptionInstrCommRateField)(unsafe.Pointer(QryOptionInstrCommRate)), C.int(RequestID),
	)

	slog.Info("thost trader api ReqQryOptionInstrCommRate executed")

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQryExecOrder(QryExecOrder *mini.CThostFtdcQryExecOrderField, RequestID int) int {
	slog.Info("executing thost trader api ReqQryExecOrder")

	rtn := C.CallReqQryExecOrder(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQryExecOrder,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQryExecOrderField)(unsafe.Pointer(QryExecOrder)), C.int(RequestID),
	)

	slog.Info("thost trader api ReqQryExecOrder executed")

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQryForQuote(QryForQuote *mini.CThostFtdcQryForQuoteField, RequestID int) int {
	slog.Info("executing thost trader api ReqQryForQuote")

	rtn := C.CallReqQryForQuote(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQryForQuote,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQryForQuoteField)(unsafe.Pointer(QryForQuote)), C.int(RequestID),
	)

	slog.Info("thost trader api ReqQryForQuote executed")

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQryQuote(QryQuote *mini.CThostFtdcQryQuoteField, RequestID int) int {
	slog.Info("executing thost trader api ReqQryQuote")

	rtn := C.CallReqQryQuote(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQryQuote,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQryQuoteField)(unsafe.Pointer(QryQuote)), C.int(RequestID),
	)

	slog.Info("thost trader api ReqQryQuote executed")

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQryInstrumentOrderCommRate(QryInstrumentOrderCommRate *mini.CThostFtdcQryInstrumentOrderCommRateField, RequestID int) int {
	slog.Info("executing thost trader api ReqQryInstrumentOrderCommRate")

	rtn := C.CallReqQryInstrumentOrderCommRate(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQryInstrumentOrderCommRate,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQryInstrumentOrderCommRateField)(unsafe.Pointer(QryInstrumentOrderCommRate)), C.int(RequestID),
	)

	slog.Info("thost trader api ReqQryInstrumentOrderCommRate executed")

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQryForQuoteParam(QryForQuoteParam *mini.CThostFtdcQryForQuoteParamField, RequestID int) int {
	slog.Info("executing thost trader api ReqQryForQuoteParam")

	rtn := C.CallReqQryForQuoteParam(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQryForQuoteParam,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQryForQuoteParamField)(unsafe.Pointer(QryForQuoteParam)), C.int(RequestID),
	)

	slog.Info("thost trader api ReqQryForQuoteParam executed")

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQryTraderOffer(QryTraderOffer *mini.CThostFtdcQryTraderOfferField, RequestID int) int {
	slog.Info("executing thost trader api ReqQryTraderOffer")

	rtn := C.CallReqQryTraderOffer(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQryTraderOffer,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQryTraderOfferField)(unsafe.Pointer(QryTraderOffer)), C.int(RequestID),
	)

	slog.Info("thost trader api ReqQryTraderOffer executed")

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQryInvestorProdSPBMDetail(QryInvestorProdSPBMDetail *mini.CThostFtdcQryInvestorProdSPBMDetailField, RequestID int) int {
	slog.Info("executing thost trader api ReqQryInvestorProdSPBMDetail")

	rtn := C.CallReqQryInvestorProdSPBMDetail(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQryInvestorProdSPBMDetail,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQryInvestorProdSPBMDetailField)(unsafe.Pointer(QryInvestorProdSPBMDetail)), C.int(RequestID),
	)

	slog.Info("thost trader api ReqQryInvestorProdSPBMDetail executed")

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQrySPMMInvestorCommodityGroupMargin(QrySPMMInvestorCommodityGroupMargin *mini.CThostFtdcQrySPMMInvestorCommodityGroupMarginField, RequestID int) int {
	slog.Info("executing thost trader api ReqQrySPMMInvestorCommodityGroupMargin")

	rtn := C.CallReqQrySPMMInvestorCommodityGroupMargin(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQrySPMMInvestorCommodityGroupMargin,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQrySPMMInvestorCommodityGroupMarginField)(unsafe.Pointer(QrySPMMInvestorCommodityGroupMargin)), C.int(RequestID),
	)

	slog.Info("thost trader api ReqQrySPMMInvestorCommodityGroupMargin executed")

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQryRULEInvestorProdMargin(QryRULEInvestorProdMargin *mini.CThostFtdcQryRULEInvestorProdMarginField, RequestID int) int {
	slog.Info("executing thost trader api ReqQryRULEInvestorProdMargin")

	rtn := C.CallReqQryRULEInvestorProdMargin(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQryRULEInvestorProdMargin,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQryRULEInvestorProdMarginField)(unsafe.Pointer(QryRULEInvestorProdMargin)), C.int(RequestID),
	)

	slog.Info("thost trader api ReqQryRULEInvestorProdMargin executed")

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQryControlParam(QryControlParam *mini.CThostFtdcQryControlParamField, RequestID int) int {
	slog.Info("executing thost trader api ReqQryControlParam")

	rtn := C.CallReqQryControlParam(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQryControlParam,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQryControlParamField)(unsafe.Pointer(QryControlParam)), C.int(RequestID),
	)

	slog.Info("thost trader api ReqQryControlParam executed")

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQryOffsetSetting(QryOffsetSetting *mini.CThostFtdcQryOffsetSettingField, RequestID int) int {
	slog.Info("executing thost trader api ReqQryOffsetSetting")

	rtn := C.CallReqQryOffsetSetting(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQryOffsetSetting,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQryOffsetSettingField)(unsafe.Pointer(QryOffsetSetting)), C.int(RequestID),
	)

	slog.Info("thost trader api ReqQryOffsetSetting executed")

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqUserPasswordUpdate(UserPasswordUpdate *mini.CThostFtdcUserPasswordUpdateField, RequestID int) int {
	slog.Info("executing thost trader api ReqUserPasswordUpdate")

	rtn := C.CallReqUserPasswordUpdate(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqUserPasswordUpdate,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcUserPasswordUpdateField)(unsafe.Pointer(UserPasswordUpdate)), C.int(RequestID),
	)

	slog.Info("thost trader api ReqUserPasswordUpdate executed")

	return int(rtn)
}
