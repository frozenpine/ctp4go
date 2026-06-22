package v6_7_13

/*
#cgo CFLAGS: -I. -I${SRCDIR} -I${SRCDIR}/../../dependencies/future/v6.7.13/
#cgo LDFLAGS: -ldl

#include "td_api_helper.h"
#include "td_spi_helper.h"
*/
import "C"
import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"path/filepath"
	"runtime"
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
}

func (api *ThostFtdcTraderApi) Release() {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"executing thost trader api Release",
	)

	C.CallRelease(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_Release,
		unsafe.Pointer(api.apiPtr),
	)

	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"thost trader api Release executed",
	)
}

func (api *ThostFtdcTraderApi) Init() {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"executing thost trader api Init",
	)

	C.CallInit(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_Init,
		unsafe.Pointer(api.apiPtr),
	)

	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"thost trader api Init executed",
	)
}

func (api *ThostFtdcTraderApi) Join() int {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"executing thost trader api Join",
	)

	rtn := C.CallJoin(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_Join,
		unsafe.Pointer(api.apiPtr),
	)

	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"thost trader api Join executed",
	)

	return int(rtn)
}

func (api *ThostFtdcTraderApi) GetTradingDay() string {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"executing thost trader api GetTradingDay",
	)

	rtn := C.CallGetTradingDay(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_GetTradingDay,
		unsafe.Pointer(api.apiPtr),
	)

	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"thost trader api GetTradingDay executed",
	)

	return C.GoString(rtn)
}

func (api *ThostFtdcTraderApi) GetFrontInfo(pFrontInfo *thost.CThostFtdcFrontInfoField) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"executing thost trader api GetFrontInfo",
	)

	C.CallGetFrontInfo(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_GetFrontInfo,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcFrontInfoField)(unsafe.Pointer(pFrontInfo)),
	)

	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"thost trader api GetFrontInfo executed",
	)
}

func (api *ThostFtdcTraderApi) RegisterFront(pszFrontAddress string) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"executing thost trader api RegisterFront",
		slog.String("front", pszFrontAddress),
	)

	front := C.CString(pszFrontAddress)
	defer C.free(unsafe.Pointer(front))

	C.CallRegisterFront(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_RegisterFront,
		unsafe.Pointer(api.apiPtr), front,
	)

	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"thost trader api RegisterFront executed",
	)
}

func (api *ThostFtdcTraderApi) RegisterNameServer(pszNsAddress string) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"executing thost trader api RegisterNameServer",
		slog.String("nameserver", pszNsAddress),
	)

	ns := C.CString(pszNsAddress)
	defer C.free(unsafe.Pointer(ns))

	C.CallRegisterNameServer(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_RegisterNameServer,
		unsafe.Pointer(api.apiPtr), ns,
	)

	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"thost trader api RegisterNameServer executed",
	)
}

func (api *ThostFtdcTraderApi) RegisterFensUserInfo(
	pFensUserInfo *thost.CThostFtdcFensUserInfoField,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"executing thost trader api RegisterFensUserInfo",
	)

	C.CallRegisterFensUserInfo(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_RegisterFensUserInfo,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcFensUserInfoField)(unsafe.Pointer(pFensUserInfo)),
	)

	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"thost trader api RegisterFensUserInfo executed",
	)
}

func (api *ThostFtdcTraderApi) RegisterSpi(pSpi thost.TraderSpi) {
	if pSpi == nil {
		slog.Error("%w: spi is nil", thost.ErrInvalidArgs)
		return
	}

	api.spiPtr = new(ThostFtdcTraderSpi)
	api.spiPtr.callback = pSpi

	api.spiPtr.Pinner.Pin(unsafe.Pointer(api.spiPtr))
	slog.Info(
		"thost trader spi created & pinned",
		slog.Any("spi", pSpi),
		slog.Any("pinned", api.spiPtr),
	)

	cSpi := C.malloc(C.sizeof_CThostFtdcTraderSpiExt)
	(*C.CThostFtdcTraderSpiExt)(cSpi).vtable = spiCVtablePtr
	(*C.CThostFtdcTraderSpiExt)(cSpi).spi = unsafe.Pointer(api.spiPtr)

	slog.Info(
		"registering thost trader spi",
		slog.Any("spi", pSpi),
		slog.Any("pinned", api.spiPtr),
	)

	C.CallRegisterSpi(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_RegisterSpi,
		unsafe.Pointer(api.apiPtr), cSpi,
	)

	slog.Info(
		"thost trader spi registered",
		slog.Any("spi", pSpi),
		slog.Any("pinned", api.spiPtr),
	)
}

func (api *ThostFtdcTraderApi) SubscribePrivateTopic(
	nResumeType thost.THOST_TE_RESUME_TYPE, nSeqNo int,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"executing thost trader api SubscribePrivateTopic",
	)

	C.CallSubscribePrivateTopic(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_SubscribePrivateTopic,
		unsafe.Pointer(api.apiPtr), C.int(nResumeType), C.int(nSeqNo),
	)

	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"thost trader api SubscribePrivateTopic executed",
	)
}

func (api *ThostFtdcTraderApi) SubscribePublicTopic(
	nResumeType thost.THOST_TE_RESUME_TYPE,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"executing thost trader api SubscribePublicTopic",
	)

	C.CallSubscribePublicTopic(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_SubscribePublicTopic,
		unsafe.Pointer(api.apiPtr), C.int(nResumeType),
	)

	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"thost trader api SubscribePublicTopic executed",
	)
}

func (api *ThostFtdcTraderApi) ReqAuthenticate(
	pReqAuthenticateField *thost.CThostFtdcReqAuthenticateField, nRequestID int,
) int {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"executing thost trader api ReqAuthenticate",
	)

	rtn := C.CallReqAuthenticate(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqAuthenticate,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcReqAuthenticateField)(unsafe.Pointer(pReqAuthenticateField)),
		C.int(nRequestID),
	)

	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"thost trader api ReqAuthenticate executed",
	)

	return int(rtn)
}

func (api *ThostFtdcTraderApi) RegisterUserSystemInfo(
	pUserSystemInfo *thost.CThostFtdcUserSystemInfoField,
) int {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"executing thost trader api RegisterUserSystemInfo",
	)

	rtn := C.CallRegisterUserSystemInfo(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_RegisterUserSystemInfo,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcUserSystemInfoField)(unsafe.Pointer(pUserSystemInfo)),
	)

	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"thost trader api RegisterUserSystemInfo executed",
	)

	return int(rtn)
}

func (api *ThostFtdcTraderApi) SubmitUserSystemInfo(
	pUserSystemInfo *thost.CThostFtdcUserSystemInfoField,
) int {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"executing thost trader api SubmitUserSystemInfo",
	)

	rtn := C.CallSubmitUserSystemInfo(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_SubmitUserSystemInfo,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcUserSystemInfoField)(unsafe.Pointer(pUserSystemInfo)),
	)

	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"thost trader api SubmitUserSystemInfo executed",
	)

	return int(rtn)
}

func (api *ThostFtdcTraderApi) RegisterWechatUserSystemInfo(
	pUserSystemInfo *thost.CThostFtdcWechatUserSystemInfoField,
) int {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"executing thost trader api RegisterWechatUserSystemInfo",
	)

	rtn := C.CallRegisterWechatUserSystemInfo(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_RegisterWechatUserSystemInfo,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcWechatUserSystemInfoField)(unsafe.Pointer(pUserSystemInfo)),
	)

	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"thost trader api RegisterWechatUserSystemInfo executed",
	)

	return int(rtn)
}

func (api *ThostFtdcTraderApi) SubmitWechatUserSystemInfo(
	pUserSystemInfo *thost.CThostFtdcWechatUserSystemInfoField,
) int {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"executing thost trader api SubmitWechatUserSystemInfo",
	)

	rtn := C.CallSubmitWechatUserSystemInfo(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_SubmitWechatUserSystemInfo,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcWechatUserSystemInfoField)(unsafe.Pointer(pUserSystemInfo)),
	)

	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"thost trader api SubmitWechatUserSystemInfo executed",
	)

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqUserLogin(pReqUserLoginField *thost.CThostFtdcReqUserLoginField, nRequestID int) int {
}

func (api *ThostFtdcTraderApi) ReqUserLogout(pUserLogout *thost.CThostFtdcUserLogoutField, nRequestID int) int {
}

func (api *ThostFtdcTraderApi) ReqUserPasswordUpdate(pUserPasswordUpdate *thost.CThostFtdcUserPasswordUpdateField, nRequestID int) int {
}

func (api *ThostFtdcTraderApi) ReqTradingAccountPasswordUpdate(pTradingAccountPasswordUpdate *thost.CThostFtdcTradingAccountPasswordUpdateField, nRequestID int) int {
}

func (api *ThostFtdcTraderApi) ReqUserAuthMethod(pReqUserAuthMethod *thost.CThostFtdcReqUserAuthMethodField, nRequestID int) int {
}

func (api *ThostFtdcTraderApi) ReqGenUserCaptcha(pReqGenUserCaptcha *thost.CThostFtdcReqGenUserCaptchaField, nRequestID int) int {
}

func (api *ThostFtdcTraderApi) ReqGenUserText(pReqGenUserText *thost.CThostFtdcReqGenUserTextField, nRequestID int) int {
}

func (api *ThostFtdcTraderApi) ReqUserLoginWithCaptcha(pReqUserLoginWithCaptcha *thost.CThostFtdcReqUserLoginWithCaptchaField, nRequestID int) int {
}

func (api *ThostFtdcTraderApi) ReqUserLoginWithText(pReqUserLoginWithText *thost.CThostFtdcReqUserLoginWithTextField, nRequestID int) int {
}

func (api *ThostFtdcTraderApi) ReqUserLoginWithOTP(pReqUserLoginWithOTP *thost.CThostFtdcReqUserLoginWithOTPField, nRequestID int) int {
}

func (api *ThostFtdcTraderApi) ReqOrderInsert(pInputOrder *thost.CThostFtdcInputOrderField, nRequestID int) int {
}

func (api *ThostFtdcTraderApi) ReqParkedOrderInsert(pParkedOrder *thost.CThostFtdcParkedOrderField, nRequestID int) int {
}

func (api *ThostFtdcTraderApi) ReqParkedOrderAction(pParkedOrderAction *thost.CThostFtdcParkedOrderActionField, nRequestID int) int {
}

func (api *ThostFtdcTraderApi) ReqOrderAction(pInputOrderAction *thost.CThostFtdcInputOrderActionField, nRequestID int) int {
}

func (api *ThostFtdcTraderApi) ReqQryMaxOrderVolume(pQryMaxOrderVolume *thost.CThostFtdcQryMaxOrderVolumeField, nRequestID int) int {
}

func (api *ThostFtdcTraderApi) ReqSettlementInfoConfirm(pSettlementInfoConfirm *thost.CThostFtdcSettlementInfoConfirmField, nRequestID int) int {
}

func (api *ThostFtdcTraderApi) ReqRemoveParkedOrder(pRemoveParkedOrder *thost.CThostFtdcRemoveParkedOrderField, nRequestID int) int {
}

func (api *ThostFtdcTraderApi) ReqRemoveParkedOrderAction(pRemoveParkedOrderAction *thost.CThostFtdcRemoveParkedOrderActionField, nRequestID int) int {
}

func (api *ThostFtdcTraderApi) ReqExecOrderInsert(pInputExecOrder *thost.CThostFtdcInputExecOrderField, nRequestID int) int {
}

func (api *ThostFtdcTraderApi) ReqExecOrderAction(pInputExecOrderAction *thost.CThostFtdcInputExecOrderActionField, nRequestID int) int {
}

func (api *ThostFtdcTraderApi) ReqForQuoteInsert(pInputForQuote *thost.CThostFtdcInputForQuoteField, nRequestID int) int {
}

func (api *ThostFtdcTraderApi) ReqQuoteInsert(pInputQuote *thost.CThostFtdcInputQuoteField, nRequestID int) int {
}

func (api *ThostFtdcTraderApi) ReqQuoteAction(pInputQuoteAction *thost.CThostFtdcInputQuoteActionField, nRequestID int) int {
}

func (api *ThostFtdcTraderApi) ReqBatchOrderAction(pInputBatchOrderAction *thost.CThostFtdcInputBatchOrderActionField, nRequestID int) int {
}

func (api *ThostFtdcTraderApi) ReqOptionSelfCloseInsert(pInputOptionSelfClose *thost.CThostFtdcInputOptionSelfCloseField, nRequestID int) int {
}

func (api *ThostFtdcTraderApi) ReqOptionSelfCloseAction(pInputOptionSelfCloseAction *thost.CThostFtdcInputOptionSelfCloseActionField, nRequestID int) int {
}

func (api *ThostFtdcTraderApi) ReqCombActionInsert(pInputCombAction *thost.CThostFtdcInputCombActionField, nRequestID int) int {
}

func (api *ThostFtdcTraderApi) ReqQryOrder(pQryOrder *thost.CThostFtdcQryOrderField, nRequestID int) int {
}

func (api *ThostFtdcTraderApi) ReqQryTrade(pQryTrade *thost.CThostFtdcQryTradeField, nRequestID int) int {
}

func (api *ThostFtdcTraderApi) ReqQryInvestorPosition(pQryInvestorPosition *thost.CThostFtdcQryInvestorPositionField, nRequestID int) int {
}

func (api *ThostFtdcTraderApi) ReqQryTradingAccount(pQryTradingAccount *thost.CThostFtdcQryTradingAccountField, nRequestID int) int {
}

func (api *ThostFtdcTraderApi) ReqQryInvestor(pQryInvestor *thost.CThostFtdcQryInvestorField, nRequestID int) int {
}

func (api *ThostFtdcTraderApi) ReqQryTradingCode(pQryTradingCode *thost.CThostFtdcQryTradingCodeField, nRequestID int) int {
}

func (api *ThostFtdcTraderApi) ReqQryInstrumentMarginRate(pQryInstrumentMarginRate *thost.CThostFtdcQryInstrumentMarginRateField, nRequestID int) int {
}

func (api *ThostFtdcTraderApi) ReqQryInstrumentCommissionRate(pQryInstrumentCommissionRate *thost.CThostFtdcQryInstrumentCommissionRateField, nRequestID int) int {
}

func (api *ThostFtdcTraderApi) ReqQryUserSession(pQryUserSession *thost.CThostFtdcQryUserSessionField, nRequestID int) int {
}

func (api *ThostFtdcTraderApi) ReqQryExchange(pQryExchange *thost.CThostFtdcQryExchangeField, nRequestID int) int {
}

func (api *ThostFtdcTraderApi) ReqQryProduct(pQryProduct *thost.CThostFtdcQryProductField, nRequestID int) int {
}

func (api *ThostFtdcTraderApi) ReqQryInstrument(pQryInstrument *thost.CThostFtdcQryInstrumentField, nRequestID int) int {
}

func (api *ThostFtdcTraderApi) ReqQryDepthMarketData(pQryDepthMarketData *thost.CThostFtdcQryDepthMarketDataField, nRequestID int) int {
}

func (api *ThostFtdcTraderApi) ReqQryTraderOffer(pQryTraderOffer *thost.CThostFtdcQryTraderOfferField, nRequestID int) int {
}

func (api *ThostFtdcTraderApi) ReqQrySettlementInfo(pQrySettlementInfo *thost.CThostFtdcQrySettlementInfoField, nRequestID int) int {
}

func (api *ThostFtdcTraderApi) ReqQryTransferBank(pQryTransferBank *thost.CThostFtdcQryTransferBankField, nRequestID int) int {
}

func (api *ThostFtdcTraderApi) ReqQryInvestorPositionDetail(pQryInvestorPositionDetail *thost.CThostFtdcQryInvestorPositionDetailField, nRequestID int) int {
}

func (api *ThostFtdcTraderApi) ReqQryNotice(pQryNotice *thost.CThostFtdcQryNoticeField, nRequestID int) int {
}

func (api *ThostFtdcTraderApi) ReqQrySettlementInfoConfirm(pQrySettlementInfoConfirm *thost.CThostFtdcQrySettlementInfoConfirmField, nRequestID int) int {
}

func (api *ThostFtdcTraderApi) ReqQryInvestorPositionCombineDetail(pQryInvestorPositionCombineDetail *thost.CThostFtdcQryInvestorPositionCombineDetailField, nRequestID int) int {
}

func (api *ThostFtdcTraderApi) ReqQryCFMMCTradingAccountKey(pQryCFMMCTradingAccountKey *thost.CThostFtdcQryCFMMCTradingAccountKeyField, nRequestID int) int {
}

func (api *ThostFtdcTraderApi) ReqQryEWarrantOffset(pQryEWarrantOffset *thost.CThostFtdcQryEWarrantOffsetField, nRequestID int) int {
}

func (api *ThostFtdcTraderApi) ReqQryInvestorProductGroupMargin(pQryInvestorProductGroupMargin *thost.CThostFtdcQryInvestorProductGroupMarginField, nRequestID int) int {
}

func (api *ThostFtdcTraderApi) ReqQryExchangeMarginRate(pQryExchangeMarginRate *thost.CThostFtdcQryExchangeMarginRateField, nRequestID int) int {
}

func (api *ThostFtdcTraderApi) ReqQryExchangeMarginRateAdjust(pQryExchangeMarginRateAdjust *thost.CThostFtdcQryExchangeMarginRateAdjustField, nRequestID int) int {
}

func (api *ThostFtdcTraderApi) ReqQryExchangeRate(pQryExchangeRate *thost.CThostFtdcQryExchangeRateField, nRequestID int) int {
}

func (api *ThostFtdcTraderApi) ReqQrySecAgentACIDMap(pQrySecAgentACIDMap *thost.CThostFtdcQrySecAgentACIDMapField, nRequestID int) int {
}

func (api *ThostFtdcTraderApi) ReqQryProductExchRate(pQryProductExchRate *thost.CThostFtdcQryProductExchRateField, nRequestID int) int {
}

func (api *ThostFtdcTraderApi) ReqQryProductGroup(pQryProductGroup *thost.CThostFtdcQryProductGroupField, nRequestID int) int {
}

func (api *ThostFtdcTraderApi) ReqQryMMInstrumentCommissionRate(pQryMMInstrumentCommissionRate *thost.CThostFtdcQryMMInstrumentCommissionRateField, nRequestID int) int {
}

func (api *ThostFtdcTraderApi) ReqQryMMOptionInstrCommRate(pQryMMOptionInstrCommRate *thost.CThostFtdcQryMMOptionInstrCommRateField, nRequestID int) int {
}

func (api *ThostFtdcTraderApi) ReqQryInstrumentOrderCommRate(pQryInstrumentOrderCommRate *thost.CThostFtdcQryInstrumentOrderCommRateField, nRequestID int) int {
}

func (api *ThostFtdcTraderApi) ReqQrySecAgentTradingAccount(pQryTradingAccount *thost.CThostFtdcQryTradingAccountField, nRequestID int) int {
}

func (api *ThostFtdcTraderApi) ReqQrySecAgentCheckMode(pQrySecAgentCheckMode *thost.CThostFtdcQrySecAgentCheckModeField, nRequestID int) int {
}

func (api *ThostFtdcTraderApi) ReqQrySecAgentTradeInfo(pQrySecAgentTradeInfo *thost.CThostFtdcQrySecAgentTradeInfoField, nRequestID int) int {
}

func (api *ThostFtdcTraderApi) ReqQryOptionInstrTradeCost(pQryOptionInstrTradeCost *thost.CThostFtdcQryOptionInstrTradeCostField, nRequestID int) int {
}

func (api *ThostFtdcTraderApi) ReqQryOptionInstrCommRate(pQryOptionInstrCommRate *thost.CThostFtdcQryOptionInstrCommRateField, nRequestID int) int {
}

func (api *ThostFtdcTraderApi) ReqQryExecOrder(pQryExecOrder *thost.CThostFtdcQryExecOrderField, nRequestID int) int {
}

func (api *ThostFtdcTraderApi) ReqQryForQuote(pQryForQuote *thost.CThostFtdcQryForQuoteField, nRequestID int) int {
}

func (api *ThostFtdcTraderApi) ReqQryQuote(pQryQuote *thost.CThostFtdcQryQuoteField, nRequestID int) int {
}

func (api *ThostFtdcTraderApi) ReqQryOptionSelfClose(pQryOptionSelfClose *thost.CThostFtdcQryOptionSelfCloseField, nRequestID int) int {
}

func (api *ThostFtdcTraderApi) ReqQryInvestUnit(pQryInvestUnit *thost.CThostFtdcQryInvestUnitField, nRequestID int) int {
}

func (api *ThostFtdcTraderApi) ReqQryCombInstrumentGuard(pQryCombInstrumentGuard *thost.CThostFtdcQryCombInstrumentGuardField, nRequestID int) int {
}

func (api *ThostFtdcTraderApi) ReqQryCombAction(pQryCombAction *thost.CThostFtdcQryCombActionField, nRequestID int) int {
}

func (api *ThostFtdcTraderApi) ReqQryTransferSerial(pQryTransferSerial *thost.CThostFtdcQryTransferSerialField, nRequestID int) int {
}

func (api *ThostFtdcTraderApi) ReqQryAccountregister(pQryAccountregister *thost.CThostFtdcQryAccountregisterField, nRequestID int) int {
}

func (api *ThostFtdcTraderApi) ReqQryContractBank(pQryContractBank *thost.CThostFtdcQryContractBankField, nRequestID int) int {
}

func (api *ThostFtdcTraderApi) ReqQryParkedOrder(pQryParkedOrder *thost.CThostFtdcQryParkedOrderField, nRequestID int) int {
}

func (api *ThostFtdcTraderApi) ReqQryParkedOrderAction(pQryParkedOrderAction *thost.CThostFtdcQryParkedOrderActionField, nRequestID int) int {
}

func (api *ThostFtdcTraderApi) ReqQryTradingNotice(pQryTradingNotice *thost.CThostFtdcQryTradingNoticeField, nRequestID int) int {
}

func (api *ThostFtdcTraderApi) ReqQryBrokerTradingParams(pQryBrokerTradingParams *thost.CThostFtdcQryBrokerTradingParamsField, nRequestID int) int {
}

func (api *ThostFtdcTraderApi) ReqQryBrokerTradingAlgos(pQryBrokerTradingAlgos *thost.CThostFtdcQryBrokerTradingAlgosField, nRequestID int) int {
}

func (api *ThostFtdcTraderApi) ReqQueryCFMMCTradingAccountToken(pQueryCFMMCTradingAccountToken *thost.CThostFtdcQueryCFMMCTradingAccountTokenField, nRequestID int) int {
}

func (api *ThostFtdcTraderApi) ReqFromBankToFutureByFuture(pReqTransfer *thost.CThostFtdcReqTransferField, nRequestID int) int {
}

func (api *ThostFtdcTraderApi) ReqFromFutureToBankByFuture(pReqTransfer *thost.CThostFtdcReqTransferField, nRequestID int) int {
}

func (api *ThostFtdcTraderApi) ReqQueryBankAccountMoneyByFuture(pReqQueryAccount *thost.CThostFtdcReqQueryAccountField, nRequestID int) int {
}

func (api *ThostFtdcTraderApi) ReqQryClassifiedInstrument(pQryClassifiedInstrument *thost.CThostFtdcQryClassifiedInstrumentField, nRequestID int) int {
}

func (api *ThostFtdcTraderApi) ReqQryCombPromotionParam(pQryCombPromotionParam *thost.CThostFtdcQryCombPromotionParamField, nRequestID int) int {
}

func (api *ThostFtdcTraderApi) ReqQryRiskSettleInvstPosition(pQryRiskSettleInvstPosition *thost.CThostFtdcQryRiskSettleInvstPositionField, nRequestID int) int {
}

func (api *ThostFtdcTraderApi) ReqQryRiskSettleProductStatus(pQryRiskSettleProductStatus *thost.CThostFtdcQryRiskSettleProductStatusField, nRequestID int) int {
}

func (api *ThostFtdcTraderApi) ReqQrySPBMFutureParameter(pQrySPBMFutureParameter *thost.CThostFtdcQrySPBMFutureParameterField, nRequestID int) int {
}

func (api *ThostFtdcTraderApi) ReqQrySPBMOptionParameter(pQrySPBMOptionParameter *thost.CThostFtdcQrySPBMOptionParameterField, nRequestID int) int {
}

func (api *ThostFtdcTraderApi) ReqQrySPBMIntraParameter(pQrySPBMIntraParameter *thost.CThostFtdcQrySPBMIntraParameterField, nRequestID int) int {
}

func (api *ThostFtdcTraderApi) ReqQrySPBMInterParameter(pQrySPBMInterParameter *thost.CThostFtdcQrySPBMInterParameterField, nRequestID int) int {
}

func (api *ThostFtdcTraderApi) ReqQrySPBMPortfDefinition(pQrySPBMPortfDefinition *thost.CThostFtdcQrySPBMPortfDefinitionField, nRequestID int) int {
}

func (api *ThostFtdcTraderApi) ReqQrySPBMInvestorPortfDef(pQrySPBMInvestorPortfDef *thost.CThostFtdcQrySPBMInvestorPortfDefField, nRequestID int) int {
}

func (api *ThostFtdcTraderApi) ReqQryInvestorPortfMarginRatio(pQryInvestorPortfMarginRatio *thost.CThostFtdcQryInvestorPortfMarginRatioField, nRequestID int) int {
}

func (api *ThostFtdcTraderApi) ReqQryInvestorProdSPBMDetail(pQryInvestorProdSPBMDetail *thost.CThostFtdcQryInvestorProdSPBMDetailField, nRequestID int) int {
}

func (api *ThostFtdcTraderApi) ReqQryInvestorCommoditySPMMMargin(pQryInvestorCommoditySPMMMargin *thost.CThostFtdcQryInvestorCommoditySPMMMarginField, nRequestID int) int {
}

func (api *ThostFtdcTraderApi) ReqQryInvestorCommodityGroupSPMMMargin(pQryInvestorCommodityGroupSPMMMargin *thost.CThostFtdcQryInvestorCommodityGroupSPMMMarginField, nRequestID int) int {
}

func (api *ThostFtdcTraderApi) ReqQrySPMMInstParam(pQrySPMMInstParam *thost.CThostFtdcQrySPMMInstParamField, nRequestID int) int {
}

func (api *ThostFtdcTraderApi) ReqQrySPMMProductParam(pQrySPMMProductParam *thost.CThostFtdcQrySPMMProductParamField, nRequestID int) int {
}

func (api *ThostFtdcTraderApi) ReqQrySPBMAddOnInterParameter(pQrySPBMAddOnInterParameter *thost.CThostFtdcQrySPBMAddOnInterParameterField, nRequestID int) int {
}

func (api *ThostFtdcTraderApi) ReqQryRCAMSCombProductInfo(pQryRCAMSCombProductInfo *thost.CThostFtdcQryRCAMSCombProductInfoField, nRequestID int) int {
}

func (api *ThostFtdcTraderApi) ReqQryRCAMSInstrParameter(pQryRCAMSInstrParameter *thost.CThostFtdcQryRCAMSInstrParameterField, nRequestID int) int {
}

func (api *ThostFtdcTraderApi) ReqQryRCAMSIntraParameter(pQryRCAMSIntraParameter *thost.CThostFtdcQryRCAMSIntraParameterField, nRequestID int) int {
}

func (api *ThostFtdcTraderApi) ReqQryRCAMSInterParameter(pQryRCAMSInterParameter *thost.CThostFtdcQryRCAMSInterParameterField, nRequestID int) int {
}

func (api *ThostFtdcTraderApi) ReqQryRCAMSShortOptAdjustParam(pQryRCAMSShortOptAdjustParam *thost.CThostFtdcQryRCAMSShortOptAdjustParamField, nRequestID int) int {
}

func (api *ThostFtdcTraderApi) ReqQryRCAMSInvestorCombPosition(pQryRCAMSInvestorCombPosition *thost.CThostFtdcQryRCAMSInvestorCombPositionField, nRequestID int) int {
}

func (api *ThostFtdcTraderApi) ReqQryInvestorProdRCAMSMargin(pQryInvestorProdRCAMSMargin *thost.CThostFtdcQryInvestorProdRCAMSMarginField, nRequestID int) int {
}

func (api *ThostFtdcTraderApi) ReqQryRULEInstrParameter(pQryRULEInstrParameter *thost.CThostFtdcQryRULEInstrParameterField, nRequestID int) int {
}

func (api *ThostFtdcTraderApi) ReqQryRULEIntraParameter(pQryRULEIntraParameter *thost.CThostFtdcQryRULEIntraParameterField, nRequestID int) int {
}

func (api *ThostFtdcTraderApi) ReqQryRULEInterParameter(pQryRULEInterParameter *thost.CThostFtdcQryRULEInterParameterField, nRequestID int) int {
}

func (api *ThostFtdcTraderApi) ReqQryInvestorProdRULEMargin(pQryInvestorProdRULEMargin *thost.CThostFtdcQryInvestorProdRULEMarginField, nRequestID int) int {
}

func (api *ThostFtdcTraderApi) ReqQryInvestorPortfSetting(pQryInvestorPortfSetting *thost.CThostFtdcQryInvestorPortfSettingField, nRequestID int) int {
}

func (api *ThostFtdcTraderApi) ReqQryInvestorInfoCommRec(pQryInvestorInfoCommRec *thost.CThostFtdcQryInvestorInfoCommRecField, nRequestID int) int {
}

func (api *ThostFtdcTraderApi) ReqQryCombLeg(pQryCombLeg *thost.CThostFtdcQryCombLegField, nRequestID int) int {
}

func (api *ThostFtdcTraderApi) ReqOffsetSetting(pInputOffsetSetting *thost.CThostFtdcInputOffsetSettingField, nRequestID int) int {
}

func (api *ThostFtdcTraderApi) ReqCancelOffsetSetting(pInputOffsetSetting *thost.CThostFtdcInputOffsetSettingField, nRequestID int) int {
}

func (api *ThostFtdcTraderApi) ReqQryOffsetSetting(pQryOffsetSetting *thost.CThostFtdcQryOffsetSettingField, nRequestID int) int {
}

func (api *ThostFtdcTraderApi) ReqGenSMSCode(pReqGenSMSCode *thost.CThostFtdcReqGenSMSCodeField, nRequestID int) int {
}

func (api *ThostFtdcTraderApi) ReqSpdApply(pInputSpdApply *thost.CThostFtdcInputSpdApplyField, nRequestID int) int {
}

func (api *ThostFtdcTraderApi) ReqSpdApplyAction(pInputSpdApplyAction *thost.CThostFtdcInputSpdApplyActionField, nRequestID int) int {
}

func (api *ThostFtdcTraderApi) ReqQrySpdApply(pQrySpdApply *thost.CThostFtdcQrySpdApplyField, nRequestID int) int {
}

func (api *ThostFtdcTraderApi) ReqHedgeCfm(pInputHedgeCfm *thost.CThostFtdcInputHedgeCfmField, nRequestID int) int {
}

func (api *ThostFtdcTraderApi) ReqHedgeCfmAction(pInputHedgeCfmAction *thost.CThostFtdcInputHedgeCfmActionField, nRequestID int) int {
}

func (api *ThostFtdcTraderApi) ReqQryHedgeCfm(pQryHedgeCfm *thost.CThostFtdcQryHedgeCfmField, nRequestID int) int {
}
