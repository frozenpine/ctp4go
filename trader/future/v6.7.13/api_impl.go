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

	"github.com/frozenpine/ctp4go/thost"
	"github.com/frozenpine/ctp4go/thost/future"
	"github.com/frozenpine/ctp4go/thost/future/types"
)

// var (
//     // 确保Api封装完整实现了thost中的接口签名
//     _ future.TraderApi = &ThostFtdcTraderApi{}
// )

func CreateThostFtdcTraderApi(
	libPath string, FlowPath string, IsProductionMode bool,
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
		pszFlowPath, C.bool(IsProductionMode),
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

func (api *ThostFtdcTraderApi) Init() {
	slog.Info("executing thost trader api Init")

	C.CallInit(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_Init,
		unsafe.Pointer(api.apiPtr),
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

func (api *ThostFtdcTraderApi) GetFrontInfo(FrontInfo *future.CThostFtdcFrontInfoField) {
	slog.Info("executing thost trader api GetFrontInfo")

	C.CallGetFrontInfo(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_GetFrontInfo,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcFrontInfoField)(unsafe.Pointer(FrontInfo)),
	)

	slog.Info("thost trader api GetFrontInfo executed")

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

func (api *ThostFtdcTraderApi) RegisterNameServer(NsAddress string) {
	slog.Info("executing thost trader api RegisterNameServer")

	pszNsAddress := C.CString(NsAddress)
	defer C.free(unsafe.Pointer(pszNsAddress))

	C.CallRegisterNameServer(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_RegisterNameServer,
		unsafe.Pointer(api.apiPtr),
		pszNsAddress,
	)

	slog.Info("thost trader api RegisterNameServer executed")

}

func (api *ThostFtdcTraderApi) RegisterFensUserInfo(FensUserInfo *future.CThostFtdcFensUserInfoField) {
	slog.Info("executing thost trader api RegisterFensUserInfo")

	C.CallRegisterFensUserInfo(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_RegisterFensUserInfo,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcFensUserInfoField)(unsafe.Pointer(FensUserInfo)),
	)

	slog.Info("thost trader api RegisterFensUserInfo executed")

}

func (api *ThostFtdcTraderApi) RegisterSpi(Spi future.TraderSpi) {
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

func (api *ThostFtdcTraderApi) SubscribePrivateTopic(ResumeType int, SeqNo int) {
	slog.Info("executing thost trader api SubscribePrivateTopic")

	C.CallSubscribePrivateTopic(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_SubscribePrivateTopic,
		unsafe.Pointer(api.apiPtr),
		C.int(ResumeType), C.int(SeqNo),
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

func (api *ThostFtdcTraderApi) ReqAuthenticate(ReqAuthenticateField *future.CThostFtdcReqAuthenticateField, RequestID int) int {
	slog.Info("executing thost trader api ReqAuthenticate")

	rtn := C.CallReqAuthenticate(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqAuthenticate,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcReqAuthenticateField)(unsafe.Pointer(ReqAuthenticateField)), C.int(RequestID),
	)

	slog.Info("thost trader api ReqAuthenticate executed")

	return int(rtn)
}

func (api *ThostFtdcTraderApi) RegisterUserSystemInfo(UserSystemInfo *future.CThostFtdcUserSystemInfoField) int {
	slog.Info("executing thost trader api RegisterUserSystemInfo")

	rtn := C.CallRegisterUserSystemInfo(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_RegisterUserSystemInfo,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcUserSystemInfoField)(unsafe.Pointer(UserSystemInfo)),
	)

	slog.Info("thost trader api RegisterUserSystemInfo executed")

	return int(rtn)
}

func (api *ThostFtdcTraderApi) SubmitUserSystemInfo(UserSystemInfo *future.CThostFtdcUserSystemInfoField) int {
	slog.Info("executing thost trader api SubmitUserSystemInfo")

	rtn := C.CallSubmitUserSystemInfo(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_SubmitUserSystemInfo,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcUserSystemInfoField)(unsafe.Pointer(UserSystemInfo)),
	)

	slog.Info("thost trader api SubmitUserSystemInfo executed")

	return int(rtn)
}

func (api *ThostFtdcTraderApi) RegisterWechatUserSystemInfo(UserSystemInfo *future.CThostFtdcWechatUserSystemInfoField) int {
	slog.Info("executing thost trader api RegisterWechatUserSystemInfo")

	rtn := C.CallRegisterWechatUserSystemInfo(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_RegisterWechatUserSystemInfo,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcWechatUserSystemInfoField)(unsafe.Pointer(UserSystemInfo)),
	)

	slog.Info("thost trader api RegisterWechatUserSystemInfo executed")

	return int(rtn)
}

func (api *ThostFtdcTraderApi) SubmitWechatUserSystemInfo(UserSystemInfo *future.CThostFtdcWechatUserSystemInfoField) int {
	slog.Info("executing thost trader api SubmitWechatUserSystemInfo")

	rtn := C.CallSubmitWechatUserSystemInfo(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_SubmitWechatUserSystemInfo,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcWechatUserSystemInfoField)(unsafe.Pointer(UserSystemInfo)),
	)

	slog.Info("thost trader api SubmitWechatUserSystemInfo executed")

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqUserLogin(ReqUserLoginField *future.CThostFtdcReqUserLoginField, RequestID int) int {
	slog.Info("executing thost trader api ReqUserLogin")

	rtn := C.CallReqUserLogin(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqUserLogin,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcReqUserLoginField)(unsafe.Pointer(ReqUserLoginField)), C.int(RequestID),
	)

	slog.Info("thost trader api ReqUserLogin executed")

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqUserLogout(UserLogout *future.CThostFtdcUserLogoutField, RequestID int) int {
	slog.Info("executing thost trader api ReqUserLogout")

	rtn := C.CallReqUserLogout(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqUserLogout,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcUserLogoutField)(unsafe.Pointer(UserLogout)), C.int(RequestID),
	)

	slog.Info("thost trader api ReqUserLogout executed")

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqUserPasswordUpdate(UserPasswordUpdate *future.CThostFtdcUserPasswordUpdateField, RequestID int) int {
	slog.Info("executing thost trader api ReqUserPasswordUpdate")

	rtn := C.CallReqUserPasswordUpdate(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqUserPasswordUpdate,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcUserPasswordUpdateField)(unsafe.Pointer(UserPasswordUpdate)), C.int(RequestID),
	)

	slog.Info("thost trader api ReqUserPasswordUpdate executed")

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqTradingAccountPasswordUpdate(TradingAccountPasswordUpdate *future.CThostFtdcTradingAccountPasswordUpdateField, RequestID int) int {
	slog.Info("executing thost trader api ReqTradingAccountPasswordUpdate")

	rtn := C.CallReqTradingAccountPasswordUpdate(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqTradingAccountPasswordUpdate,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcTradingAccountPasswordUpdateField)(unsafe.Pointer(TradingAccountPasswordUpdate)), C.int(RequestID),
	)

	slog.Info("thost trader api ReqTradingAccountPasswordUpdate executed")

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqUserAuthMethod(ReqUserAuthMethod *future.CThostFtdcReqUserAuthMethodField, RequestID int) int {
	slog.Info("executing thost trader api ReqUserAuthMethod")

	rtn := C.CallReqUserAuthMethod(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqUserAuthMethod,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcReqUserAuthMethodField)(unsafe.Pointer(ReqUserAuthMethod)), C.int(RequestID),
	)

	slog.Info("thost trader api ReqUserAuthMethod executed")

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqGenUserCaptcha(ReqGenUserCaptcha *future.CThostFtdcReqGenUserCaptchaField, RequestID int) int {
	slog.Info("executing thost trader api ReqGenUserCaptcha")

	rtn := C.CallReqGenUserCaptcha(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqGenUserCaptcha,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcReqGenUserCaptchaField)(unsafe.Pointer(ReqGenUserCaptcha)), C.int(RequestID),
	)

	slog.Info("thost trader api ReqGenUserCaptcha executed")

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqGenUserText(ReqGenUserText *future.CThostFtdcReqGenUserTextField, RequestID int) int {
	slog.Info("executing thost trader api ReqGenUserText")

	rtn := C.CallReqGenUserText(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqGenUserText,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcReqGenUserTextField)(unsafe.Pointer(ReqGenUserText)), C.int(RequestID),
	)

	slog.Info("thost trader api ReqGenUserText executed")

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqUserLoginWithCaptcha(ReqUserLoginWithCaptcha *future.CThostFtdcReqUserLoginWithCaptchaField, RequestID int) int {
	slog.Info("executing thost trader api ReqUserLoginWithCaptcha")

	rtn := C.CallReqUserLoginWithCaptcha(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqUserLoginWithCaptcha,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcReqUserLoginWithCaptchaField)(unsafe.Pointer(ReqUserLoginWithCaptcha)), C.int(RequestID),
	)

	slog.Info("thost trader api ReqUserLoginWithCaptcha executed")

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqUserLoginWithText(ReqUserLoginWithText *future.CThostFtdcReqUserLoginWithTextField, RequestID int) int {
	slog.Info("executing thost trader api ReqUserLoginWithText")

	rtn := C.CallReqUserLoginWithText(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqUserLoginWithText,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcReqUserLoginWithTextField)(unsafe.Pointer(ReqUserLoginWithText)), C.int(RequestID),
	)

	slog.Info("thost trader api ReqUserLoginWithText executed")

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqUserLoginWithOTP(ReqUserLoginWithOTP *future.CThostFtdcReqUserLoginWithOTPField, RequestID int) int {
	slog.Info("executing thost trader api ReqUserLoginWithOTP")

	rtn := C.CallReqUserLoginWithOTP(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqUserLoginWithOTP,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcReqUserLoginWithOTPField)(unsafe.Pointer(ReqUserLoginWithOTP)), C.int(RequestID),
	)

	slog.Info("thost trader api ReqUserLoginWithOTP executed")

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqOrderInsert(InputOrder *future.CThostFtdcInputOrderField, RequestID int) int {
	slog.Info("executing thost trader api ReqOrderInsert")

	rtn := C.CallReqOrderInsert(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqOrderInsert,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcInputOrderField)(unsafe.Pointer(InputOrder)), C.int(RequestID),
	)

	slog.Info("thost trader api ReqOrderInsert executed")

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqParkedOrderInsert(ParkedOrder *future.CThostFtdcParkedOrderField, RequestID int) int {
	slog.Info("executing thost trader api ReqParkedOrderInsert")

	rtn := C.CallReqParkedOrderInsert(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqParkedOrderInsert,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcParkedOrderField)(unsafe.Pointer(ParkedOrder)), C.int(RequestID),
	)

	slog.Info("thost trader api ReqParkedOrderInsert executed")

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqParkedOrderAction(ParkedOrderAction *future.CThostFtdcParkedOrderActionField, RequestID int) int {
	slog.Info("executing thost trader api ReqParkedOrderAction")

	rtn := C.CallReqParkedOrderAction(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqParkedOrderAction,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcParkedOrderActionField)(unsafe.Pointer(ParkedOrderAction)), C.int(RequestID),
	)

	slog.Info("thost trader api ReqParkedOrderAction executed")

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqOrderAction(InputOrderAction *future.CThostFtdcInputOrderActionField, RequestID int) int {
	slog.Info("executing thost trader api ReqOrderAction")

	rtn := C.CallReqOrderAction(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqOrderAction,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcInputOrderActionField)(unsafe.Pointer(InputOrderAction)), C.int(RequestID),
	)

	slog.Info("thost trader api ReqOrderAction executed")

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQryMaxOrderVolume(QryMaxOrderVolume *future.CThostFtdcQryMaxOrderVolumeField, RequestID int) int {
	slog.Info("executing thost trader api ReqQryMaxOrderVolume")

	rtn := C.CallReqQryMaxOrderVolume(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQryMaxOrderVolume,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQryMaxOrderVolumeField)(unsafe.Pointer(QryMaxOrderVolume)), C.int(RequestID),
	)

	slog.Info("thost trader api ReqQryMaxOrderVolume executed")

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqSettlementInfoConfirm(SettlementInfoConfirm *future.CThostFtdcSettlementInfoConfirmField, RequestID int) int {
	slog.Info("executing thost trader api ReqSettlementInfoConfirm")

	rtn := C.CallReqSettlementInfoConfirm(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqSettlementInfoConfirm,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcSettlementInfoConfirmField)(unsafe.Pointer(SettlementInfoConfirm)), C.int(RequestID),
	)

	slog.Info("thost trader api ReqSettlementInfoConfirm executed")

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqRemoveParkedOrder(RemoveParkedOrder *future.CThostFtdcRemoveParkedOrderField, RequestID int) int {
	slog.Info("executing thost trader api ReqRemoveParkedOrder")

	rtn := C.CallReqRemoveParkedOrder(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqRemoveParkedOrder,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcRemoveParkedOrderField)(unsafe.Pointer(RemoveParkedOrder)), C.int(RequestID),
	)

	slog.Info("thost trader api ReqRemoveParkedOrder executed")

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqRemoveParkedOrderAction(RemoveParkedOrderAction *future.CThostFtdcRemoveParkedOrderActionField, RequestID int) int {
	slog.Info("executing thost trader api ReqRemoveParkedOrderAction")

	rtn := C.CallReqRemoveParkedOrderAction(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqRemoveParkedOrderAction,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcRemoveParkedOrderActionField)(unsafe.Pointer(RemoveParkedOrderAction)), C.int(RequestID),
	)

	slog.Info("thost trader api ReqRemoveParkedOrderAction executed")

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqExecOrderInsert(InputExecOrder *future.CThostFtdcInputExecOrderField, RequestID int) int {
	slog.Info("executing thost trader api ReqExecOrderInsert")

	rtn := C.CallReqExecOrderInsert(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqExecOrderInsert,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcInputExecOrderField)(unsafe.Pointer(InputExecOrder)), C.int(RequestID),
	)

	slog.Info("thost trader api ReqExecOrderInsert executed")

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqExecOrderAction(InputExecOrderAction *future.CThostFtdcInputExecOrderActionField, RequestID int) int {
	slog.Info("executing thost trader api ReqExecOrderAction")

	rtn := C.CallReqExecOrderAction(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqExecOrderAction,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcInputExecOrderActionField)(unsafe.Pointer(InputExecOrderAction)), C.int(RequestID),
	)

	slog.Info("thost trader api ReqExecOrderAction executed")

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqForQuoteInsert(InputForQuote *future.CThostFtdcInputForQuoteField, RequestID int) int {
	slog.Info("executing thost trader api ReqForQuoteInsert")

	rtn := C.CallReqForQuoteInsert(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqForQuoteInsert,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcInputForQuoteField)(unsafe.Pointer(InputForQuote)), C.int(RequestID),
	)

	slog.Info("thost trader api ReqForQuoteInsert executed")

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQuoteInsert(InputQuote *future.CThostFtdcInputQuoteField, RequestID int) int {
	slog.Info("executing thost trader api ReqQuoteInsert")

	rtn := C.CallReqQuoteInsert(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQuoteInsert,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcInputQuoteField)(unsafe.Pointer(InputQuote)), C.int(RequestID),
	)

	slog.Info("thost trader api ReqQuoteInsert executed")

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQuoteAction(InputQuoteAction *future.CThostFtdcInputQuoteActionField, RequestID int) int {
	slog.Info("executing thost trader api ReqQuoteAction")

	rtn := C.CallReqQuoteAction(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQuoteAction,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcInputQuoteActionField)(unsafe.Pointer(InputQuoteAction)), C.int(RequestID),
	)

	slog.Info("thost trader api ReqQuoteAction executed")

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqBatchOrderAction(InputBatchOrderAction *future.CThostFtdcInputBatchOrderActionField, RequestID int) int {
	slog.Info("executing thost trader api ReqBatchOrderAction")

	rtn := C.CallReqBatchOrderAction(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqBatchOrderAction,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcInputBatchOrderActionField)(unsafe.Pointer(InputBatchOrderAction)), C.int(RequestID),
	)

	slog.Info("thost trader api ReqBatchOrderAction executed")

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqOptionSelfCloseInsert(InputOptionSelfClose *future.CThostFtdcInputOptionSelfCloseField, RequestID int) int {
	slog.Info("executing thost trader api ReqOptionSelfCloseInsert")

	rtn := C.CallReqOptionSelfCloseInsert(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqOptionSelfCloseInsert,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcInputOptionSelfCloseField)(unsafe.Pointer(InputOptionSelfClose)), C.int(RequestID),
	)

	slog.Info("thost trader api ReqOptionSelfCloseInsert executed")

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqOptionSelfCloseAction(InputOptionSelfCloseAction *future.CThostFtdcInputOptionSelfCloseActionField, RequestID int) int {
	slog.Info("executing thost trader api ReqOptionSelfCloseAction")

	rtn := C.CallReqOptionSelfCloseAction(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqOptionSelfCloseAction,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcInputOptionSelfCloseActionField)(unsafe.Pointer(InputOptionSelfCloseAction)), C.int(RequestID),
	)

	slog.Info("thost trader api ReqOptionSelfCloseAction executed")

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqCombActionInsert(InputCombAction *future.CThostFtdcInputCombActionField, RequestID int) int {
	slog.Info("executing thost trader api ReqCombActionInsert")

	rtn := C.CallReqCombActionInsert(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqCombActionInsert,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcInputCombActionField)(unsafe.Pointer(InputCombAction)), C.int(RequestID),
	)

	slog.Info("thost trader api ReqCombActionInsert executed")

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQryOrder(QryOrder *future.CThostFtdcQryOrderField, RequestID int) int {
	slog.Info("executing thost trader api ReqQryOrder")

	rtn := C.CallReqQryOrder(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQryOrder,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQryOrderField)(unsafe.Pointer(QryOrder)), C.int(RequestID),
	)

	slog.Info("thost trader api ReqQryOrder executed")

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQryTrade(QryTrade *future.CThostFtdcQryTradeField, RequestID int) int {
	slog.Info("executing thost trader api ReqQryTrade")

	rtn := C.CallReqQryTrade(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQryTrade,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQryTradeField)(unsafe.Pointer(QryTrade)), C.int(RequestID),
	)

	slog.Info("thost trader api ReqQryTrade executed")

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQryInvestorPosition(QryInvestorPosition *future.CThostFtdcQryInvestorPositionField, RequestID int) int {
	slog.Info("executing thost trader api ReqQryInvestorPosition")

	rtn := C.CallReqQryInvestorPosition(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQryInvestorPosition,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQryInvestorPositionField)(unsafe.Pointer(QryInvestorPosition)), C.int(RequestID),
	)

	slog.Info("thost trader api ReqQryInvestorPosition executed")

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQryTradingAccount(QryTradingAccount *future.CThostFtdcQryTradingAccountField, RequestID int) int {
	slog.Info("executing thost trader api ReqQryTradingAccount")

	rtn := C.CallReqQryTradingAccount(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQryTradingAccount,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQryTradingAccountField)(unsafe.Pointer(QryTradingAccount)), C.int(RequestID),
	)

	slog.Info("thost trader api ReqQryTradingAccount executed")

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQryInvestor(QryInvestor *future.CThostFtdcQryInvestorField, RequestID int) int {
	slog.Info("executing thost trader api ReqQryInvestor")

	rtn := C.CallReqQryInvestor(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQryInvestor,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQryInvestorField)(unsafe.Pointer(QryInvestor)), C.int(RequestID),
	)

	slog.Info("thost trader api ReqQryInvestor executed")

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQryTradingCode(QryTradingCode *future.CThostFtdcQryTradingCodeField, RequestID int) int {
	slog.Info("executing thost trader api ReqQryTradingCode")

	rtn := C.CallReqQryTradingCode(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQryTradingCode,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQryTradingCodeField)(unsafe.Pointer(QryTradingCode)), C.int(RequestID),
	)

	slog.Info("thost trader api ReqQryTradingCode executed")

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQryInstrumentMarginRate(QryInstrumentMarginRate *future.CThostFtdcQryInstrumentMarginRateField, RequestID int) int {
	slog.Info("executing thost trader api ReqQryInstrumentMarginRate")

	rtn := C.CallReqQryInstrumentMarginRate(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQryInstrumentMarginRate,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQryInstrumentMarginRateField)(unsafe.Pointer(QryInstrumentMarginRate)), C.int(RequestID),
	)

	slog.Info("thost trader api ReqQryInstrumentMarginRate executed")

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQryInstrumentCommissionRate(QryInstrumentCommissionRate *future.CThostFtdcQryInstrumentCommissionRateField, RequestID int) int {
	slog.Info("executing thost trader api ReqQryInstrumentCommissionRate")

	rtn := C.CallReqQryInstrumentCommissionRate(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQryInstrumentCommissionRate,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQryInstrumentCommissionRateField)(unsafe.Pointer(QryInstrumentCommissionRate)), C.int(RequestID),
	)

	slog.Info("thost trader api ReqQryInstrumentCommissionRate executed")

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQryUserSession(QryUserSession *future.CThostFtdcQryUserSessionField, RequestID int) int {
	slog.Info("executing thost trader api ReqQryUserSession")

	rtn := C.CallReqQryUserSession(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQryUserSession,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQryUserSessionField)(unsafe.Pointer(QryUserSession)), C.int(RequestID),
	)

	slog.Info("thost trader api ReqQryUserSession executed")

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQryExchange(QryExchange *future.CThostFtdcQryExchangeField, RequestID int) int {
	slog.Info("executing thost trader api ReqQryExchange")

	rtn := C.CallReqQryExchange(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQryExchange,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQryExchangeField)(unsafe.Pointer(QryExchange)), C.int(RequestID),
	)

	slog.Info("thost trader api ReqQryExchange executed")

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQryProduct(QryProduct *future.CThostFtdcQryProductField, RequestID int) int {
	slog.Info("executing thost trader api ReqQryProduct")

	rtn := C.CallReqQryProduct(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQryProduct,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQryProductField)(unsafe.Pointer(QryProduct)), C.int(RequestID),
	)

	slog.Info("thost trader api ReqQryProduct executed")

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQryInstrument(QryInstrument *future.CThostFtdcQryInstrumentField, RequestID int) int {
	slog.Info("executing thost trader api ReqQryInstrument")

	rtn := C.CallReqQryInstrument(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQryInstrument,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQryInstrumentField)(unsafe.Pointer(QryInstrument)), C.int(RequestID),
	)

	slog.Info("thost trader api ReqQryInstrument executed")

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQryDepthMarketData(QryDepthMarketData *future.CThostFtdcQryDepthMarketDataField, RequestID int) int {
	slog.Info("executing thost trader api ReqQryDepthMarketData")

	rtn := C.CallReqQryDepthMarketData(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQryDepthMarketData,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQryDepthMarketDataField)(unsafe.Pointer(QryDepthMarketData)), C.int(RequestID),
	)

	slog.Info("thost trader api ReqQryDepthMarketData executed")

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQryTraderOffer(QryTraderOffer *future.CThostFtdcQryTraderOfferField, RequestID int) int {
	slog.Info("executing thost trader api ReqQryTraderOffer")

	rtn := C.CallReqQryTraderOffer(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQryTraderOffer,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQryTraderOfferField)(unsafe.Pointer(QryTraderOffer)), C.int(RequestID),
	)

	slog.Info("thost trader api ReqQryTraderOffer executed")

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQrySettlementInfo(QrySettlementInfo *future.CThostFtdcQrySettlementInfoField, RequestID int) int {
	slog.Info("executing thost trader api ReqQrySettlementInfo")

	rtn := C.CallReqQrySettlementInfo(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQrySettlementInfo,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQrySettlementInfoField)(unsafe.Pointer(QrySettlementInfo)), C.int(RequestID),
	)

	slog.Info("thost trader api ReqQrySettlementInfo executed")

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQryTransferBank(QryTransferBank *future.CThostFtdcQryTransferBankField, RequestID int) int {
	slog.Info("executing thost trader api ReqQryTransferBank")

	rtn := C.CallReqQryTransferBank(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQryTransferBank,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQryTransferBankField)(unsafe.Pointer(QryTransferBank)), C.int(RequestID),
	)

	slog.Info("thost trader api ReqQryTransferBank executed")

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQryInvestorPositionDetail(QryInvestorPositionDetail *future.CThostFtdcQryInvestorPositionDetailField, RequestID int) int {
	slog.Info("executing thost trader api ReqQryInvestorPositionDetail")

	rtn := C.CallReqQryInvestorPositionDetail(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQryInvestorPositionDetail,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQryInvestorPositionDetailField)(unsafe.Pointer(QryInvestorPositionDetail)), C.int(RequestID),
	)

	slog.Info("thost trader api ReqQryInvestorPositionDetail executed")

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQryNotice(QryNotice *future.CThostFtdcQryNoticeField, RequestID int) int {
	slog.Info("executing thost trader api ReqQryNotice")

	rtn := C.CallReqQryNotice(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQryNotice,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQryNoticeField)(unsafe.Pointer(QryNotice)), C.int(RequestID),
	)

	slog.Info("thost trader api ReqQryNotice executed")

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQrySettlementInfoConfirm(QrySettlementInfoConfirm *future.CThostFtdcQrySettlementInfoConfirmField, RequestID int) int {
	slog.Info("executing thost trader api ReqQrySettlementInfoConfirm")

	rtn := C.CallReqQrySettlementInfoConfirm(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQrySettlementInfoConfirm,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQrySettlementInfoConfirmField)(unsafe.Pointer(QrySettlementInfoConfirm)), C.int(RequestID),
	)

	slog.Info("thost trader api ReqQrySettlementInfoConfirm executed")

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQryInvestorPositionCombineDetail(QryInvestorPositionCombineDetail *future.CThostFtdcQryInvestorPositionCombineDetailField, RequestID int) int {
	slog.Info("executing thost trader api ReqQryInvestorPositionCombineDetail")

	rtn := C.CallReqQryInvestorPositionCombineDetail(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQryInvestorPositionCombineDetail,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQryInvestorPositionCombineDetailField)(unsafe.Pointer(QryInvestorPositionCombineDetail)), C.int(RequestID),
	)

	slog.Info("thost trader api ReqQryInvestorPositionCombineDetail executed")

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQryCFMMCTradingAccountKey(QryCFMMCTradingAccountKey *future.CThostFtdcQryCFMMCTradingAccountKeyField, RequestID int) int {
	slog.Info("executing thost trader api ReqQryCFMMCTradingAccountKey")

	rtn := C.CallReqQryCFMMCTradingAccountKey(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQryCFMMCTradingAccountKey,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQryCFMMCTradingAccountKeyField)(unsafe.Pointer(QryCFMMCTradingAccountKey)), C.int(RequestID),
	)

	slog.Info("thost trader api ReqQryCFMMCTradingAccountKey executed")

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQryEWarrantOffset(QryEWarrantOffset *future.CThostFtdcQryEWarrantOffsetField, RequestID int) int {
	slog.Info("executing thost trader api ReqQryEWarrantOffset")

	rtn := C.CallReqQryEWarrantOffset(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQryEWarrantOffset,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQryEWarrantOffsetField)(unsafe.Pointer(QryEWarrantOffset)), C.int(RequestID),
	)

	slog.Info("thost trader api ReqQryEWarrantOffset executed")

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQryInvestorProductGroupMargin(QryInvestorProductGroupMargin *future.CThostFtdcQryInvestorProductGroupMarginField, RequestID int) int {
	slog.Info("executing thost trader api ReqQryInvestorProductGroupMargin")

	rtn := C.CallReqQryInvestorProductGroupMargin(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQryInvestorProductGroupMargin,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQryInvestorProductGroupMarginField)(unsafe.Pointer(QryInvestorProductGroupMargin)), C.int(RequestID),
	)

	slog.Info("thost trader api ReqQryInvestorProductGroupMargin executed")

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQryExchangeMarginRate(QryExchangeMarginRate *future.CThostFtdcQryExchangeMarginRateField, RequestID int) int {
	slog.Info("executing thost trader api ReqQryExchangeMarginRate")

	rtn := C.CallReqQryExchangeMarginRate(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQryExchangeMarginRate,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQryExchangeMarginRateField)(unsafe.Pointer(QryExchangeMarginRate)), C.int(RequestID),
	)

	slog.Info("thost trader api ReqQryExchangeMarginRate executed")

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQryExchangeMarginRateAdjust(QryExchangeMarginRateAdjust *future.CThostFtdcQryExchangeMarginRateAdjustField, RequestID int) int {
	slog.Info("executing thost trader api ReqQryExchangeMarginRateAdjust")

	rtn := C.CallReqQryExchangeMarginRateAdjust(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQryExchangeMarginRateAdjust,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQryExchangeMarginRateAdjustField)(unsafe.Pointer(QryExchangeMarginRateAdjust)), C.int(RequestID),
	)

	slog.Info("thost trader api ReqQryExchangeMarginRateAdjust executed")

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQryExchangeRate(QryExchangeRate *future.CThostFtdcQryExchangeRateField, RequestID int) int {
	slog.Info("executing thost trader api ReqQryExchangeRate")

	rtn := C.CallReqQryExchangeRate(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQryExchangeRate,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQryExchangeRateField)(unsafe.Pointer(QryExchangeRate)), C.int(RequestID),
	)

	slog.Info("thost trader api ReqQryExchangeRate executed")

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQrySecAgentACIDMap(QrySecAgentACIDMap *future.CThostFtdcQrySecAgentACIDMapField, RequestID int) int {
	slog.Info("executing thost trader api ReqQrySecAgentACIDMap")

	rtn := C.CallReqQrySecAgentACIDMap(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQrySecAgentACIDMap,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQrySecAgentACIDMapField)(unsafe.Pointer(QrySecAgentACIDMap)), C.int(RequestID),
	)

	slog.Info("thost trader api ReqQrySecAgentACIDMap executed")

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQryProductExchRate(QryProductExchRate *future.CThostFtdcQryProductExchRateField, RequestID int) int {
	slog.Info("executing thost trader api ReqQryProductExchRate")

	rtn := C.CallReqQryProductExchRate(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQryProductExchRate,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQryProductExchRateField)(unsafe.Pointer(QryProductExchRate)), C.int(RequestID),
	)

	slog.Info("thost trader api ReqQryProductExchRate executed")

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQryProductGroup(QryProductGroup *future.CThostFtdcQryProductGroupField, RequestID int) int {
	slog.Info("executing thost trader api ReqQryProductGroup")

	rtn := C.CallReqQryProductGroup(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQryProductGroup,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQryProductGroupField)(unsafe.Pointer(QryProductGroup)), C.int(RequestID),
	)

	slog.Info("thost trader api ReqQryProductGroup executed")

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQryMMInstrumentCommissionRate(QryMMInstrumentCommissionRate *future.CThostFtdcQryMMInstrumentCommissionRateField, RequestID int) int {
	slog.Info("executing thost trader api ReqQryMMInstrumentCommissionRate")

	rtn := C.CallReqQryMMInstrumentCommissionRate(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQryMMInstrumentCommissionRate,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQryMMInstrumentCommissionRateField)(unsafe.Pointer(QryMMInstrumentCommissionRate)), C.int(RequestID),
	)

	slog.Info("thost trader api ReqQryMMInstrumentCommissionRate executed")

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQryMMOptionInstrCommRate(QryMMOptionInstrCommRate *future.CThostFtdcQryMMOptionInstrCommRateField, RequestID int) int {
	slog.Info("executing thost trader api ReqQryMMOptionInstrCommRate")

	rtn := C.CallReqQryMMOptionInstrCommRate(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQryMMOptionInstrCommRate,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQryMMOptionInstrCommRateField)(unsafe.Pointer(QryMMOptionInstrCommRate)), C.int(RequestID),
	)

	slog.Info("thost trader api ReqQryMMOptionInstrCommRate executed")

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQryInstrumentOrderCommRate(QryInstrumentOrderCommRate *future.CThostFtdcQryInstrumentOrderCommRateField, RequestID int) int {
	slog.Info("executing thost trader api ReqQryInstrumentOrderCommRate")

	rtn := C.CallReqQryInstrumentOrderCommRate(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQryInstrumentOrderCommRate,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQryInstrumentOrderCommRateField)(unsafe.Pointer(QryInstrumentOrderCommRate)), C.int(RequestID),
	)

	slog.Info("thost trader api ReqQryInstrumentOrderCommRate executed")

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQrySecAgentTradingAccount(QryTradingAccount *future.CThostFtdcQryTradingAccountField, RequestID int) int {
	slog.Info("executing thost trader api ReqQrySecAgentTradingAccount")

	rtn := C.CallReqQrySecAgentTradingAccount(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQrySecAgentTradingAccount,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQryTradingAccountField)(unsafe.Pointer(QryTradingAccount)), C.int(RequestID),
	)

	slog.Info("thost trader api ReqQrySecAgentTradingAccount executed")

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQrySecAgentCheckMode(QrySecAgentCheckMode *future.CThostFtdcQrySecAgentCheckModeField, RequestID int) int {
	slog.Info("executing thost trader api ReqQrySecAgentCheckMode")

	rtn := C.CallReqQrySecAgentCheckMode(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQrySecAgentCheckMode,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQrySecAgentCheckModeField)(unsafe.Pointer(QrySecAgentCheckMode)), C.int(RequestID),
	)

	slog.Info("thost trader api ReqQrySecAgentCheckMode executed")

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQrySecAgentTradeInfo(QrySecAgentTradeInfo *future.CThostFtdcQrySecAgentTradeInfoField, RequestID int) int {
	slog.Info("executing thost trader api ReqQrySecAgentTradeInfo")

	rtn := C.CallReqQrySecAgentTradeInfo(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQrySecAgentTradeInfo,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQrySecAgentTradeInfoField)(unsafe.Pointer(QrySecAgentTradeInfo)), C.int(RequestID),
	)

	slog.Info("thost trader api ReqQrySecAgentTradeInfo executed")

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQryOptionInstrTradeCost(QryOptionInstrTradeCost *future.CThostFtdcQryOptionInstrTradeCostField, RequestID int) int {
	slog.Info("executing thost trader api ReqQryOptionInstrTradeCost")

	rtn := C.CallReqQryOptionInstrTradeCost(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQryOptionInstrTradeCost,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQryOptionInstrTradeCostField)(unsafe.Pointer(QryOptionInstrTradeCost)), C.int(RequestID),
	)

	slog.Info("thost trader api ReqQryOptionInstrTradeCost executed")

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQryOptionInstrCommRate(QryOptionInstrCommRate *future.CThostFtdcQryOptionInstrCommRateField, RequestID int) int {
	slog.Info("executing thost trader api ReqQryOptionInstrCommRate")

	rtn := C.CallReqQryOptionInstrCommRate(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQryOptionInstrCommRate,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQryOptionInstrCommRateField)(unsafe.Pointer(QryOptionInstrCommRate)), C.int(RequestID),
	)

	slog.Info("thost trader api ReqQryOptionInstrCommRate executed")

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQryExecOrder(QryExecOrder *future.CThostFtdcQryExecOrderField, RequestID int) int {
	slog.Info("executing thost trader api ReqQryExecOrder")

	rtn := C.CallReqQryExecOrder(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQryExecOrder,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQryExecOrderField)(unsafe.Pointer(QryExecOrder)), C.int(RequestID),
	)

	slog.Info("thost trader api ReqQryExecOrder executed")

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQryForQuote(QryForQuote *future.CThostFtdcQryForQuoteField, RequestID int) int {
	slog.Info("executing thost trader api ReqQryForQuote")

	rtn := C.CallReqQryForQuote(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQryForQuote,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQryForQuoteField)(unsafe.Pointer(QryForQuote)), C.int(RequestID),
	)

	slog.Info("thost trader api ReqQryForQuote executed")

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQryQuote(QryQuote *future.CThostFtdcQryQuoteField, RequestID int) int {
	slog.Info("executing thost trader api ReqQryQuote")

	rtn := C.CallReqQryQuote(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQryQuote,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQryQuoteField)(unsafe.Pointer(QryQuote)), C.int(RequestID),
	)

	slog.Info("thost trader api ReqQryQuote executed")

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQryOptionSelfClose(QryOptionSelfClose *future.CThostFtdcQryOptionSelfCloseField, RequestID int) int {
	slog.Info("executing thost trader api ReqQryOptionSelfClose")

	rtn := C.CallReqQryOptionSelfClose(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQryOptionSelfClose,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQryOptionSelfCloseField)(unsafe.Pointer(QryOptionSelfClose)), C.int(RequestID),
	)

	slog.Info("thost trader api ReqQryOptionSelfClose executed")

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQryInvestUnit(QryInvestUnit *future.CThostFtdcQryInvestUnitField, RequestID int) int {
	slog.Info("executing thost trader api ReqQryInvestUnit")

	rtn := C.CallReqQryInvestUnit(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQryInvestUnit,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQryInvestUnitField)(unsafe.Pointer(QryInvestUnit)), C.int(RequestID),
	)

	slog.Info("thost trader api ReqQryInvestUnit executed")

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQryCombInstrumentGuard(QryCombInstrumentGuard *future.CThostFtdcQryCombInstrumentGuardField, RequestID int) int {
	slog.Info("executing thost trader api ReqQryCombInstrumentGuard")

	rtn := C.CallReqQryCombInstrumentGuard(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQryCombInstrumentGuard,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQryCombInstrumentGuardField)(unsafe.Pointer(QryCombInstrumentGuard)), C.int(RequestID),
	)

	slog.Info("thost trader api ReqQryCombInstrumentGuard executed")

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQryCombAction(QryCombAction *future.CThostFtdcQryCombActionField, RequestID int) int {
	slog.Info("executing thost trader api ReqQryCombAction")

	rtn := C.CallReqQryCombAction(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQryCombAction,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQryCombActionField)(unsafe.Pointer(QryCombAction)), C.int(RequestID),
	)

	slog.Info("thost trader api ReqQryCombAction executed")

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQryTransferSerial(QryTransferSerial *future.CThostFtdcQryTransferSerialField, RequestID int) int {
	slog.Info("executing thost trader api ReqQryTransferSerial")

	rtn := C.CallReqQryTransferSerial(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQryTransferSerial,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQryTransferSerialField)(unsafe.Pointer(QryTransferSerial)), C.int(RequestID),
	)

	slog.Info("thost trader api ReqQryTransferSerial executed")

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQryAccountregister(QryAccountregister *future.CThostFtdcQryAccountregisterField, RequestID int) int {
	slog.Info("executing thost trader api ReqQryAccountregister")

	rtn := C.CallReqQryAccountregister(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQryAccountregister,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQryAccountregisterField)(unsafe.Pointer(QryAccountregister)), C.int(RequestID),
	)

	slog.Info("thost trader api ReqQryAccountregister executed")

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQryContractBank(QryContractBank *future.CThostFtdcQryContractBankField, RequestID int) int {
	slog.Info("executing thost trader api ReqQryContractBank")

	rtn := C.CallReqQryContractBank(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQryContractBank,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQryContractBankField)(unsafe.Pointer(QryContractBank)), C.int(RequestID),
	)

	slog.Info("thost trader api ReqQryContractBank executed")

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQryParkedOrder(QryParkedOrder *future.CThostFtdcQryParkedOrderField, RequestID int) int {
	slog.Info("executing thost trader api ReqQryParkedOrder")

	rtn := C.CallReqQryParkedOrder(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQryParkedOrder,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQryParkedOrderField)(unsafe.Pointer(QryParkedOrder)), C.int(RequestID),
	)

	slog.Info("thost trader api ReqQryParkedOrder executed")

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQryParkedOrderAction(QryParkedOrderAction *future.CThostFtdcQryParkedOrderActionField, RequestID int) int {
	slog.Info("executing thost trader api ReqQryParkedOrderAction")

	rtn := C.CallReqQryParkedOrderAction(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQryParkedOrderAction,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQryParkedOrderActionField)(unsafe.Pointer(QryParkedOrderAction)), C.int(RequestID),
	)

	slog.Info("thost trader api ReqQryParkedOrderAction executed")

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQryTradingNotice(QryTradingNotice *future.CThostFtdcQryTradingNoticeField, RequestID int) int {
	slog.Info("executing thost trader api ReqQryTradingNotice")

	rtn := C.CallReqQryTradingNotice(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQryTradingNotice,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQryTradingNoticeField)(unsafe.Pointer(QryTradingNotice)), C.int(RequestID),
	)

	slog.Info("thost trader api ReqQryTradingNotice executed")

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQryBrokerTradingParams(QryBrokerTradingParams *future.CThostFtdcQryBrokerTradingParamsField, RequestID int) int {
	slog.Info("executing thost trader api ReqQryBrokerTradingParams")

	rtn := C.CallReqQryBrokerTradingParams(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQryBrokerTradingParams,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQryBrokerTradingParamsField)(unsafe.Pointer(QryBrokerTradingParams)), C.int(RequestID),
	)

	slog.Info("thost trader api ReqQryBrokerTradingParams executed")

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQryBrokerTradingAlgos(QryBrokerTradingAlgos *future.CThostFtdcQryBrokerTradingAlgosField, RequestID int) int {
	slog.Info("executing thost trader api ReqQryBrokerTradingAlgos")

	rtn := C.CallReqQryBrokerTradingAlgos(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQryBrokerTradingAlgos,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQryBrokerTradingAlgosField)(unsafe.Pointer(QryBrokerTradingAlgos)), C.int(RequestID),
	)

	slog.Info("thost trader api ReqQryBrokerTradingAlgos executed")

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQueryCFMMCTradingAccountToken(QueryCFMMCTradingAccountToken *future.CThostFtdcQueryCFMMCTradingAccountTokenField, RequestID int) int {
	slog.Info("executing thost trader api ReqQueryCFMMCTradingAccountToken")

	rtn := C.CallReqQueryCFMMCTradingAccountToken(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQueryCFMMCTradingAccountToken,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQueryCFMMCTradingAccountTokenField)(unsafe.Pointer(QueryCFMMCTradingAccountToken)), C.int(RequestID),
	)

	slog.Info("thost trader api ReqQueryCFMMCTradingAccountToken executed")

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqFromBankToFutureByFuture(ReqTransfer *future.CThostFtdcReqTransferField, RequestID int) int {
	slog.Info("executing thost trader api ReqFromBankToFutureByFuture")

	rtn := C.CallReqFromBankToFutureByFuture(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqFromBankToFutureByFuture,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcReqTransferField)(unsafe.Pointer(ReqTransfer)), C.int(RequestID),
	)

	slog.Info("thost trader api ReqFromBankToFutureByFuture executed")

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqFromFutureToBankByFuture(ReqTransfer *future.CThostFtdcReqTransferField, RequestID int) int {
	slog.Info("executing thost trader api ReqFromFutureToBankByFuture")

	rtn := C.CallReqFromFutureToBankByFuture(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqFromFutureToBankByFuture,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcReqTransferField)(unsafe.Pointer(ReqTransfer)), C.int(RequestID),
	)

	slog.Info("thost trader api ReqFromFutureToBankByFuture executed")

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQueryBankAccountMoneyByFuture(ReqQueryAccount *future.CThostFtdcReqQueryAccountField, RequestID int) int {
	slog.Info("executing thost trader api ReqQueryBankAccountMoneyByFuture")

	rtn := C.CallReqQueryBankAccountMoneyByFuture(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQueryBankAccountMoneyByFuture,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcReqQueryAccountField)(unsafe.Pointer(ReqQueryAccount)), C.int(RequestID),
	)

	slog.Info("thost trader api ReqQueryBankAccountMoneyByFuture executed")

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQryClassifiedInstrument(QryClassifiedInstrument *future.CThostFtdcQryClassifiedInstrumentField, RequestID int) int {
	slog.Info("executing thost trader api ReqQryClassifiedInstrument")

	rtn := C.CallReqQryClassifiedInstrument(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQryClassifiedInstrument,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQryClassifiedInstrumentField)(unsafe.Pointer(QryClassifiedInstrument)), C.int(RequestID),
	)

	slog.Info("thost trader api ReqQryClassifiedInstrument executed")

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQryCombPromotionParam(QryCombPromotionParam *future.CThostFtdcQryCombPromotionParamField, RequestID int) int {
	slog.Info("executing thost trader api ReqQryCombPromotionParam")

	rtn := C.CallReqQryCombPromotionParam(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQryCombPromotionParam,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQryCombPromotionParamField)(unsafe.Pointer(QryCombPromotionParam)), C.int(RequestID),
	)

	slog.Info("thost trader api ReqQryCombPromotionParam executed")

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQryRiskSettleInvstPosition(QryRiskSettleInvstPosition *future.CThostFtdcQryRiskSettleInvstPositionField, RequestID int) int {
	slog.Info("executing thost trader api ReqQryRiskSettleInvstPosition")

	rtn := C.CallReqQryRiskSettleInvstPosition(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQryRiskSettleInvstPosition,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQryRiskSettleInvstPositionField)(unsafe.Pointer(QryRiskSettleInvstPosition)), C.int(RequestID),
	)

	slog.Info("thost trader api ReqQryRiskSettleInvstPosition executed")

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQryRiskSettleProductStatus(QryRiskSettleProductStatus *future.CThostFtdcQryRiskSettleProductStatusField, RequestID int) int {
	slog.Info("executing thost trader api ReqQryRiskSettleProductStatus")

	rtn := C.CallReqQryRiskSettleProductStatus(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQryRiskSettleProductStatus,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQryRiskSettleProductStatusField)(unsafe.Pointer(QryRiskSettleProductStatus)), C.int(RequestID),
	)

	slog.Info("thost trader api ReqQryRiskSettleProductStatus executed")

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQrySPBMFutureParameter(QrySPBMFutureParameter *future.CThostFtdcQrySPBMFutureParameterField, RequestID int) int {
	slog.Info("executing thost trader api ReqQrySPBMFutureParameter")

	rtn := C.CallReqQrySPBMFutureParameter(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQrySPBMFutureParameter,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQrySPBMFutureParameterField)(unsafe.Pointer(QrySPBMFutureParameter)), C.int(RequestID),
	)

	slog.Info("thost trader api ReqQrySPBMFutureParameter executed")

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQrySPBMOptionParameter(QrySPBMOptionParameter *future.CThostFtdcQrySPBMOptionParameterField, RequestID int) int {
	slog.Info("executing thost trader api ReqQrySPBMOptionParameter")

	rtn := C.CallReqQrySPBMOptionParameter(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQrySPBMOptionParameter,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQrySPBMOptionParameterField)(unsafe.Pointer(QrySPBMOptionParameter)), C.int(RequestID),
	)

	slog.Info("thost trader api ReqQrySPBMOptionParameter executed")

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQrySPBMIntraParameter(QrySPBMIntraParameter *future.CThostFtdcQrySPBMIntraParameterField, RequestID int) int {
	slog.Info("executing thost trader api ReqQrySPBMIntraParameter")

	rtn := C.CallReqQrySPBMIntraParameter(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQrySPBMIntraParameter,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQrySPBMIntraParameterField)(unsafe.Pointer(QrySPBMIntraParameter)), C.int(RequestID),
	)

	slog.Info("thost trader api ReqQrySPBMIntraParameter executed")

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQrySPBMInterParameter(QrySPBMInterParameter *future.CThostFtdcQrySPBMInterParameterField, RequestID int) int {
	slog.Info("executing thost trader api ReqQrySPBMInterParameter")

	rtn := C.CallReqQrySPBMInterParameter(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQrySPBMInterParameter,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQrySPBMInterParameterField)(unsafe.Pointer(QrySPBMInterParameter)), C.int(RequestID),
	)

	slog.Info("thost trader api ReqQrySPBMInterParameter executed")

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQrySPBMPortfDefinition(QrySPBMPortfDefinition *future.CThostFtdcQrySPBMPortfDefinitionField, RequestID int) int {
	slog.Info("executing thost trader api ReqQrySPBMPortfDefinition")

	rtn := C.CallReqQrySPBMPortfDefinition(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQrySPBMPortfDefinition,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQrySPBMPortfDefinitionField)(unsafe.Pointer(QrySPBMPortfDefinition)), C.int(RequestID),
	)

	slog.Info("thost trader api ReqQrySPBMPortfDefinition executed")

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQrySPBMInvestorPortfDef(QrySPBMInvestorPortfDef *future.CThostFtdcQrySPBMInvestorPortfDefField, RequestID int) int {
	slog.Info("executing thost trader api ReqQrySPBMInvestorPortfDef")

	rtn := C.CallReqQrySPBMInvestorPortfDef(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQrySPBMInvestorPortfDef,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQrySPBMInvestorPortfDefField)(unsafe.Pointer(QrySPBMInvestorPortfDef)), C.int(RequestID),
	)

	slog.Info("thost trader api ReqQrySPBMInvestorPortfDef executed")

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQryInvestorPortfMarginRatio(QryInvestorPortfMarginRatio *future.CThostFtdcQryInvestorPortfMarginRatioField, RequestID int) int {
	slog.Info("executing thost trader api ReqQryInvestorPortfMarginRatio")

	rtn := C.CallReqQryInvestorPortfMarginRatio(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQryInvestorPortfMarginRatio,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQryInvestorPortfMarginRatioField)(unsafe.Pointer(QryInvestorPortfMarginRatio)), C.int(RequestID),
	)

	slog.Info("thost trader api ReqQryInvestorPortfMarginRatio executed")

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQryInvestorProdSPBMDetail(QryInvestorProdSPBMDetail *future.CThostFtdcQryInvestorProdSPBMDetailField, RequestID int) int {
	slog.Info("executing thost trader api ReqQryInvestorProdSPBMDetail")

	rtn := C.CallReqQryInvestorProdSPBMDetail(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQryInvestorProdSPBMDetail,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQryInvestorProdSPBMDetailField)(unsafe.Pointer(QryInvestorProdSPBMDetail)), C.int(RequestID),
	)

	slog.Info("thost trader api ReqQryInvestorProdSPBMDetail executed")

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQryInvestorCommoditySPMMMargin(QryInvestorCommoditySPMMMargin *future.CThostFtdcQryInvestorCommoditySPMMMarginField, RequestID int) int {
	slog.Info("executing thost trader api ReqQryInvestorCommoditySPMMMargin")

	rtn := C.CallReqQryInvestorCommoditySPMMMargin(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQryInvestorCommoditySPMMMargin,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQryInvestorCommoditySPMMMarginField)(unsafe.Pointer(QryInvestorCommoditySPMMMargin)), C.int(RequestID),
	)

	slog.Info("thost trader api ReqQryInvestorCommoditySPMMMargin executed")

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQryInvestorCommodityGroupSPMMMargin(QryInvestorCommodityGroupSPMMMargin *future.CThostFtdcQryInvestorCommodityGroupSPMMMarginField, RequestID int) int {
	slog.Info("executing thost trader api ReqQryInvestorCommodityGroupSPMMMargin")

	rtn := C.CallReqQryInvestorCommodityGroupSPMMMargin(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQryInvestorCommodityGroupSPMMMargin,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQryInvestorCommodityGroupSPMMMarginField)(unsafe.Pointer(QryInvestorCommodityGroupSPMMMargin)), C.int(RequestID),
	)

	slog.Info("thost trader api ReqQryInvestorCommodityGroupSPMMMargin executed")

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQrySPMMInstParam(QrySPMMInstParam *future.CThostFtdcQrySPMMInstParamField, RequestID int) int {
	slog.Info("executing thost trader api ReqQrySPMMInstParam")

	rtn := C.CallReqQrySPMMInstParam(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQrySPMMInstParam,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQrySPMMInstParamField)(unsafe.Pointer(QrySPMMInstParam)), C.int(RequestID),
	)

	slog.Info("thost trader api ReqQrySPMMInstParam executed")

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQrySPMMProductParam(QrySPMMProductParam *future.CThostFtdcQrySPMMProductParamField, RequestID int) int {
	slog.Info("executing thost trader api ReqQrySPMMProductParam")

	rtn := C.CallReqQrySPMMProductParam(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQrySPMMProductParam,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQrySPMMProductParamField)(unsafe.Pointer(QrySPMMProductParam)), C.int(RequestID),
	)

	slog.Info("thost trader api ReqQrySPMMProductParam executed")

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQrySPBMAddOnInterParameter(QrySPBMAddOnInterParameter *future.CThostFtdcQrySPBMAddOnInterParameterField, RequestID int) int {
	slog.Info("executing thost trader api ReqQrySPBMAddOnInterParameter")

	rtn := C.CallReqQrySPBMAddOnInterParameter(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQrySPBMAddOnInterParameter,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQrySPBMAddOnInterParameterField)(unsafe.Pointer(QrySPBMAddOnInterParameter)), C.int(RequestID),
	)

	slog.Info("thost trader api ReqQrySPBMAddOnInterParameter executed")

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQryRCAMSCombProductInfo(QryRCAMSCombProductInfo *future.CThostFtdcQryRCAMSCombProductInfoField, RequestID int) int {
	slog.Info("executing thost trader api ReqQryRCAMSCombProductInfo")

	rtn := C.CallReqQryRCAMSCombProductInfo(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQryRCAMSCombProductInfo,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQryRCAMSCombProductInfoField)(unsafe.Pointer(QryRCAMSCombProductInfo)), C.int(RequestID),
	)

	slog.Info("thost trader api ReqQryRCAMSCombProductInfo executed")

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQryRCAMSInstrParameter(QryRCAMSInstrParameter *future.CThostFtdcQryRCAMSInstrParameterField, RequestID int) int {
	slog.Info("executing thost trader api ReqQryRCAMSInstrParameter")

	rtn := C.CallReqQryRCAMSInstrParameter(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQryRCAMSInstrParameter,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQryRCAMSInstrParameterField)(unsafe.Pointer(QryRCAMSInstrParameter)), C.int(RequestID),
	)

	slog.Info("thost trader api ReqQryRCAMSInstrParameter executed")

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQryRCAMSIntraParameter(QryRCAMSIntraParameter *future.CThostFtdcQryRCAMSIntraParameterField, RequestID int) int {
	slog.Info("executing thost trader api ReqQryRCAMSIntraParameter")

	rtn := C.CallReqQryRCAMSIntraParameter(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQryRCAMSIntraParameter,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQryRCAMSIntraParameterField)(unsafe.Pointer(QryRCAMSIntraParameter)), C.int(RequestID),
	)

	slog.Info("thost trader api ReqQryRCAMSIntraParameter executed")

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQryRCAMSInterParameter(QryRCAMSInterParameter *future.CThostFtdcQryRCAMSInterParameterField, RequestID int) int {
	slog.Info("executing thost trader api ReqQryRCAMSInterParameter")

	rtn := C.CallReqQryRCAMSInterParameter(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQryRCAMSInterParameter,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQryRCAMSInterParameterField)(unsafe.Pointer(QryRCAMSInterParameter)), C.int(RequestID),
	)

	slog.Info("thost trader api ReqQryRCAMSInterParameter executed")

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQryRCAMSShortOptAdjustParam(QryRCAMSShortOptAdjustParam *future.CThostFtdcQryRCAMSShortOptAdjustParamField, RequestID int) int {
	slog.Info("executing thost trader api ReqQryRCAMSShortOptAdjustParam")

	rtn := C.CallReqQryRCAMSShortOptAdjustParam(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQryRCAMSShortOptAdjustParam,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQryRCAMSShortOptAdjustParamField)(unsafe.Pointer(QryRCAMSShortOptAdjustParam)), C.int(RequestID),
	)

	slog.Info("thost trader api ReqQryRCAMSShortOptAdjustParam executed")

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQryRCAMSInvestorCombPosition(QryRCAMSInvestorCombPosition *future.CThostFtdcQryRCAMSInvestorCombPositionField, RequestID int) int {
	slog.Info("executing thost trader api ReqQryRCAMSInvestorCombPosition")

	rtn := C.CallReqQryRCAMSInvestorCombPosition(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQryRCAMSInvestorCombPosition,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQryRCAMSInvestorCombPositionField)(unsafe.Pointer(QryRCAMSInvestorCombPosition)), C.int(RequestID),
	)

	slog.Info("thost trader api ReqQryRCAMSInvestorCombPosition executed")

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQryInvestorProdRCAMSMargin(QryInvestorProdRCAMSMargin *future.CThostFtdcQryInvestorProdRCAMSMarginField, RequestID int) int {
	slog.Info("executing thost trader api ReqQryInvestorProdRCAMSMargin")

	rtn := C.CallReqQryInvestorProdRCAMSMargin(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQryInvestorProdRCAMSMargin,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQryInvestorProdRCAMSMarginField)(unsafe.Pointer(QryInvestorProdRCAMSMargin)), C.int(RequestID),
	)

	slog.Info("thost trader api ReqQryInvestorProdRCAMSMargin executed")

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQryRULEInstrParameter(QryRULEInstrParameter *future.CThostFtdcQryRULEInstrParameterField, RequestID int) int {
	slog.Info("executing thost trader api ReqQryRULEInstrParameter")

	rtn := C.CallReqQryRULEInstrParameter(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQryRULEInstrParameter,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQryRULEInstrParameterField)(unsafe.Pointer(QryRULEInstrParameter)), C.int(RequestID),
	)

	slog.Info("thost trader api ReqQryRULEInstrParameter executed")

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQryRULEIntraParameter(QryRULEIntraParameter *future.CThostFtdcQryRULEIntraParameterField, RequestID int) int {
	slog.Info("executing thost trader api ReqQryRULEIntraParameter")

	rtn := C.CallReqQryRULEIntraParameter(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQryRULEIntraParameter,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQryRULEIntraParameterField)(unsafe.Pointer(QryRULEIntraParameter)), C.int(RequestID),
	)

	slog.Info("thost trader api ReqQryRULEIntraParameter executed")

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQryRULEInterParameter(QryRULEInterParameter *future.CThostFtdcQryRULEInterParameterField, RequestID int) int {
	slog.Info("executing thost trader api ReqQryRULEInterParameter")

	rtn := C.CallReqQryRULEInterParameter(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQryRULEInterParameter,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQryRULEInterParameterField)(unsafe.Pointer(QryRULEInterParameter)), C.int(RequestID),
	)

	slog.Info("thost trader api ReqQryRULEInterParameter executed")

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQryInvestorProdRULEMargin(QryInvestorProdRULEMargin *future.CThostFtdcQryInvestorProdRULEMarginField, RequestID int) int {
	slog.Info("executing thost trader api ReqQryInvestorProdRULEMargin")

	rtn := C.CallReqQryInvestorProdRULEMargin(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQryInvestorProdRULEMargin,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQryInvestorProdRULEMarginField)(unsafe.Pointer(QryInvestorProdRULEMargin)), C.int(RequestID),
	)

	slog.Info("thost trader api ReqQryInvestorProdRULEMargin executed")

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQryInvestorPortfSetting(QryInvestorPortfSetting *future.CThostFtdcQryInvestorPortfSettingField, RequestID int) int {
	slog.Info("executing thost trader api ReqQryInvestorPortfSetting")

	rtn := C.CallReqQryInvestorPortfSetting(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQryInvestorPortfSetting,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQryInvestorPortfSettingField)(unsafe.Pointer(QryInvestorPortfSetting)), C.int(RequestID),
	)

	slog.Info("thost trader api ReqQryInvestorPortfSetting executed")

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQryInvestorInfoCommRec(QryInvestorInfoCommRec *future.CThostFtdcQryInvestorInfoCommRecField, RequestID int) int {
	slog.Info("executing thost trader api ReqQryInvestorInfoCommRec")

	rtn := C.CallReqQryInvestorInfoCommRec(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQryInvestorInfoCommRec,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQryInvestorInfoCommRecField)(unsafe.Pointer(QryInvestorInfoCommRec)), C.int(RequestID),
	)

	slog.Info("thost trader api ReqQryInvestorInfoCommRec executed")

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQryCombLeg(QryCombLeg *future.CThostFtdcQryCombLegField, RequestID int) int {
	slog.Info("executing thost trader api ReqQryCombLeg")

	rtn := C.CallReqQryCombLeg(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQryCombLeg,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQryCombLegField)(unsafe.Pointer(QryCombLeg)), C.int(RequestID),
	)

	slog.Info("thost trader api ReqQryCombLeg executed")

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqOffsetSetting(InputOffsetSetting *future.CThostFtdcInputOffsetSettingField, RequestID int) int {
	slog.Info("executing thost trader api ReqOffsetSetting")

	rtn := C.CallReqOffsetSetting(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqOffsetSetting,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcInputOffsetSettingField)(unsafe.Pointer(InputOffsetSetting)), C.int(RequestID),
	)

	slog.Info("thost trader api ReqOffsetSetting executed")

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqCancelOffsetSetting(InputOffsetSetting *future.CThostFtdcInputOffsetSettingField, RequestID int) int {
	slog.Info("executing thost trader api ReqCancelOffsetSetting")

	rtn := C.CallReqCancelOffsetSetting(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqCancelOffsetSetting,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcInputOffsetSettingField)(unsafe.Pointer(InputOffsetSetting)), C.int(RequestID),
	)

	slog.Info("thost trader api ReqCancelOffsetSetting executed")

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQryOffsetSetting(QryOffsetSetting *future.CThostFtdcQryOffsetSettingField, RequestID int) int {
	slog.Info("executing thost trader api ReqQryOffsetSetting")

	rtn := C.CallReqQryOffsetSetting(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQryOffsetSetting,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQryOffsetSettingField)(unsafe.Pointer(QryOffsetSetting)), C.int(RequestID),
	)

	slog.Info("thost trader api ReqQryOffsetSetting executed")

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqGenSMSCode(ReqGenSMSCode *future.CThostFtdcReqGenSMSCodeField, RequestID int) int {
	slog.Info("executing thost trader api ReqGenSMSCode")

	rtn := C.CallReqGenSMSCode(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqGenSMSCode,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcReqGenSMSCodeField)(unsafe.Pointer(ReqGenSMSCode)), C.int(RequestID),
	)

	slog.Info("thost trader api ReqGenSMSCode executed")

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqSpdApply(InputSpdApply *future.CThostFtdcInputSpdApplyField, RequestID int) int {
	slog.Info("executing thost trader api ReqSpdApply")

	rtn := C.CallReqSpdApply(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqSpdApply,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcInputSpdApplyField)(unsafe.Pointer(InputSpdApply)), C.int(RequestID),
	)

	slog.Info("thost trader api ReqSpdApply executed")

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqSpdApplyAction(InputSpdApplyAction *future.CThostFtdcInputSpdApplyActionField, RequestID int) int {
	slog.Info("executing thost trader api ReqSpdApplyAction")

	rtn := C.CallReqSpdApplyAction(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqSpdApplyAction,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcInputSpdApplyActionField)(unsafe.Pointer(InputSpdApplyAction)), C.int(RequestID),
	)

	slog.Info("thost trader api ReqSpdApplyAction executed")

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQrySpdApply(QrySpdApply *future.CThostFtdcQrySpdApplyField, RequestID int) int {
	slog.Info("executing thost trader api ReqQrySpdApply")

	rtn := C.CallReqQrySpdApply(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQrySpdApply,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQrySpdApplyField)(unsafe.Pointer(QrySpdApply)), C.int(RequestID),
	)

	slog.Info("thost trader api ReqQrySpdApply executed")

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqHedgeCfm(InputHedgeCfm *future.CThostFtdcInputHedgeCfmField, RequestID int) int {
	slog.Info("executing thost trader api ReqHedgeCfm")

	rtn := C.CallReqHedgeCfm(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqHedgeCfm,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcInputHedgeCfmField)(unsafe.Pointer(InputHedgeCfm)), C.int(RequestID),
	)

	slog.Info("thost trader api ReqHedgeCfm executed")

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqHedgeCfmAction(InputHedgeCfmAction *future.CThostFtdcInputHedgeCfmActionField, RequestID int) int {
	slog.Info("executing thost trader api ReqHedgeCfmAction")

	rtn := C.CallReqHedgeCfmAction(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqHedgeCfmAction,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcInputHedgeCfmActionField)(unsafe.Pointer(InputHedgeCfmAction)), C.int(RequestID),
	)

	slog.Info("thost trader api ReqHedgeCfmAction executed")

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQryHedgeCfm(QryHedgeCfm *future.CThostFtdcQryHedgeCfmField, RequestID int) int {
	slog.Info("executing thost trader api ReqQryHedgeCfm")

	rtn := C.CallReqQryHedgeCfm(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQryHedgeCfm,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQryHedgeCfmField)(unsafe.Pointer(QryHedgeCfm)), C.int(RequestID),
	)

	slog.Info("thost trader api ReqQryHedgeCfm executed")

	return int(rtn)
}
