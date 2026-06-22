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
	CREATE_API_WIN   = "?CreateFtdcTraderApi@CThostFtdcTraderApi@@SAPEAV1@PEBD_N@Z"
	CREATE_API_LINUX = "_ZN19CThostFtdcTraderApi19CreateFtdcTraderApiEPKcb"

	API_VER_WIN   = "?GetApiVersion@CThostFtdcTraderApi@@SAPEBDXZ"
	API_VER_LINUX = "_ZN19CThostFtdcTraderApi13GetApiVersionEv"
)

func CreateThostFtdcTraderApi(
	libPath string, isTest bool,
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
		C.CreateFtdcTraderApi(creator), libCPath, C.bool(isTest),
	)
	if instance == nil {
		return nil, fmt.Errorf(
			"%w: thost trader api[%s] test mode[%t] create failed",
			thost.ErrApiCreateFailed, libPath, isTest,
		)
	}

	api := &ThostFtdcTraderApi{
		apiPtr:  (*C.CThostFtdcTraderApiExt)(instance),
		version: apiVer,
		lib:     lib,
	}

	runtime.SetFinalizer(api, func(ins *ThostFtdcTraderApi) {
		ins.Release()

		ins.spiPtr.Pinner.Unpin()
		ins.spiPtr = nil

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
		slog.Error("register spi is nil")
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

func (api *ThostFtdcTraderApi) ReqUserLogin(
	pReqUserLoginField *thost.CThostFtdcReqUserLoginField, nRequestID int,
) int {
	slog.Log(context.Background(), slog.LevelDebug-2,
		"executing thost trader api ReqUserLogin",
	)

	rtn := C.CallReqUserLogin(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqUserLogin,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcReqUserLoginField)(unsafe.Pointer(pReqUserLoginField)),
		C.int(nRequestID),
	)

	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"thost trader api ReqUserLogin executed",
	)

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqUserLogout(
	pUserLogout *thost.CThostFtdcUserLogoutField, nRequestID int,
) int {
	slog.Log(context.Background(), slog.LevelDebug-2,
		"executing thost trader api ReqUserLogout",
	)

	rtn := C.CallReqUserLogout(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqUserLogout,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcUserLogoutField)(unsafe.Pointer(pUserLogout)),
		C.int(nRequestID),
	)

	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"thost trader api ReqUserLogout executed",
	)

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqUserPasswordUpdate(
	pUserPasswordUpdate *thost.CThostFtdcUserPasswordUpdateField, nRequestID int,
) int {
	slog.Log(context.Background(), slog.LevelDebug-2,
		"executing thost trader api ReqUserPasswordUpdate",
	)

	rtn := C.CallReqUserPasswordUpdate(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqUserPasswordUpdate,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcUserPasswordUpdateField)(unsafe.Pointer(pUserPasswordUpdate)),
		C.int(nRequestID),
	)

	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"thost trader api ReqUserPasswordUpdate executed",
	)

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqTradingAccountPasswordUpdate(
	pTradingAccountPasswordUpdate *thost.CThostFtdcTradingAccountPasswordUpdateField, nRequestID int,
) int {
	slog.Log(context.Background(), slog.LevelDebug-2,
		"executing thost trader api ReqTradingAccountPasswordUpdate",
	)

	rtn := C.CallReqTradingAccountPasswordUpdate(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqTradingAccountPasswordUpdate,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcTradingAccountPasswordUpdateField)(unsafe.Pointer(pTradingAccountPasswordUpdate)),
		C.int(nRequestID),
	)

	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"thost trader api ReqTradingAccountPasswordUpdate executed",
	)

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqUserAuthMethod(
	pReqUserAuthMethod *thost.CThostFtdcReqUserAuthMethodField, nRequestID int,
) int {
	slog.Log(context.Background(), slog.LevelDebug-2,
		"executing thost trader api ReqUserAuthMethod",
	)

	rtn := C.CallReqUserAuthMethod(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqUserAuthMethod,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcReqUserAuthMethodField)(unsafe.Pointer(pReqUserAuthMethod)),
		C.int(nRequestID),
	)

	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"thost trader api ReqUserAuthMethod executed",
	)

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqGenUserCaptcha(
	pReqGenUserCaptcha *thost.CThostFtdcReqGenUserCaptchaField, nRequestID int,
) int {
	slog.Log(context.Background(), slog.LevelDebug-2,
		"executing thost trader api ReqGenUserCaptcha",
	)

	rtn := C.CallReqGenUserCaptcha(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqGenUserCaptcha,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcReqGenUserCaptchaField)(unsafe.Pointer(pReqGenUserCaptcha)),
		C.int(nRequestID),
	)

	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"thost trader api ReqGenUserCaptcha executed",
	)

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqGenUserText(
	pReqGenUserText *thost.CThostFtdcReqGenUserTextField, nRequestID int,
) int {
	slog.Log(context.Background(), slog.LevelDebug-2,
		"executing thost trader api ReqGenUserText",
	)

	rtn := C.CallReqGenUserText(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqGenUserText,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcReqGenUserTextField)(unsafe.Pointer(pReqGenUserText)),
		C.int(nRequestID),
	)

	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"thost trader api ReqGenUserText executed",
	)

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqUserLoginWithCaptcha(
	pReqUserLoginWithCaptcha *thost.CThostFtdcReqUserLoginWithCaptchaField, nRequestID int,
) int {
	slog.Log(context.Background(), slog.LevelDebug-2,
		"executing thost trader api ReqUserLoginWithCaptcha",
	)

	rtn := C.CallReqUserLoginWithCaptcha(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqUserLoginWithCaptcha,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcReqUserLoginWithCaptchaField)(unsafe.Pointer(pReqUserLoginWithCaptcha)),
		C.int(nRequestID),
	)

	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"thost trader api ReqUserLoginWithCaptcha executed",
	)

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqUserLoginWithText(
	pReqUserLoginWithText *thost.CThostFtdcReqUserLoginWithTextField, nRequestID int,
) int {
	slog.Log(context.Background(), slog.LevelDebug-2,
		"executing thost trader api ReqUserLoginWithText",
	)

	rtn := C.CallReqUserLoginWithText(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqUserLoginWithText,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcReqUserLoginWithTextField)(unsafe.Pointer(pReqUserLoginWithText)),
		C.int(nRequestID),
	)

	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"thost trader api ReqUserLoginWithText executed",
	)

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqUserLoginWithOTP(
	pReqUserLoginWithOTP *thost.CThostFtdcReqUserLoginWithOTPField, nRequestID int,
) int {
	slog.Log(context.Background(), slog.LevelDebug-2,
		"executing thost trader api ReqUserLoginWithOTP",
	)

	rtn := C.CallReqUserLoginWithOTP(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqUserLoginWithOTP,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcReqUserLoginWithOTPField)(unsafe.Pointer(pReqUserLoginWithOTP)),
		C.int(nRequestID),
	)

	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"thost trader api ReqUserLoginWithOTP executed",
	)

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqOrderInsert(
	pInputOrder *thost.CThostFtdcInputOrderField, nRequestID int,
) int {
	slog.Log(context.Background(), slog.LevelDebug-2,
		"executing thost trader api ReqOrderInsert",
	)

	rtn := C.CallReqOrderInsert(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqOrderInsert,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcInputOrderField)(unsafe.Pointer(pInputOrder)),
		C.int(nRequestID),
	)

	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"thost trader api ReqOrderInsert executed",
	)

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqParkedOrderInsert(
	pParkedOrder *thost.CThostFtdcParkedOrderField, nRequestID int,
) int {
	slog.Log(context.Background(), slog.LevelDebug-2,
		"executing thost trader api ReqParkedOrderInsert",
	)

	rtn := C.CallReqParkedOrderInsert(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqParkedOrderInsert,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcParkedOrderField)(unsafe.Pointer(pParkedOrder)),
		C.int(nRequestID),
	)

	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"thost trader api ReqParkedOrderInsert executed",
	)

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqParkedOrderAction(
	pParkedOrderAction *thost.CThostFtdcParkedOrderActionField, nRequestID int,
) int {
	slog.Log(context.Background(), slog.LevelDebug-2,
		"executing thost trader api ReqParkedOrderAction",
	)

	rtn := C.CallReqParkedOrderAction(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqParkedOrderAction,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcParkedOrderActionField)(unsafe.Pointer(pParkedOrderAction)),
		C.int(nRequestID),
	)

	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"thost trader api ReqParkedOrderAction executed",
	)

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqOrderAction(
	pInputOrderAction *thost.CThostFtdcInputOrderActionField, nRequestID int,
) int {
	slog.Log(context.Background(), slog.LevelDebug-2,
		"executing thost trader api ReqOrderAction",
	)

	rtn := C.CallReqOrderAction(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqOrderAction,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcInputOrderActionField)(unsafe.Pointer(pInputOrderAction)),
		C.int(nRequestID),
	)

	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"thost trader api ReqOrderAction executed",
	)

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQryMaxOrderVolume(
	pQryMaxOrderVolume *thost.CThostFtdcQryMaxOrderVolumeField, nRequestID int,
) int {
	slog.Log(context.Background(), slog.LevelDebug-2,
		"executing thost trader api ReqQryMaxOrderVolume",
	)

	rtn := C.CallReqQryMaxOrderVolume(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQryMaxOrderVolume,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQryMaxOrderVolumeField)(unsafe.Pointer(pQryMaxOrderVolume)),
		C.int(nRequestID),
	)

	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"thost trader api ReqQryMaxOrderVolume executed",
	)

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqSettlementInfoConfirm(
	pSettlementInfoConfirm *thost.CThostFtdcSettlementInfoConfirmField, nRequestID int,
) int {
	slog.Log(context.Background(), slog.LevelDebug-2,
		"executing thost trader api ReqSettlementInfoConfirm",
	)

	rtn := C.CallReqSettlementInfoConfirm(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqSettlementInfoConfirm,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcSettlementInfoConfirmField)(unsafe.Pointer(pSettlementInfoConfirm)),
		C.int(nRequestID),
	)

	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"thost trader api ReqSettlementInfoConfirm executed",
	)

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqRemoveParkedOrder(
	pRemoveParkedOrder *thost.CThostFtdcRemoveParkedOrderField, nRequestID int,
) int {
	slog.Log(context.Background(), slog.LevelDebug-2,
		"executing thost trader api ReqRemoveParkedOrder",
	)

	rtn := C.CallReqRemoveParkedOrder(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqRemoveParkedOrder,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcRemoveParkedOrderField)(unsafe.Pointer(pRemoveParkedOrder)),
		C.int(nRequestID),
	)

	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"thost trader api ReqRemoveParkedOrder executed",
	)

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqRemoveParkedOrderAction(
	pRemoveParkedOrderAction *thost.CThostFtdcRemoveParkedOrderActionField, nRequestID int,
) int {
	slog.Log(context.Background(), slog.LevelDebug-2,
		"executing thost trader api ReqRemoveParkedOrderAction",
	)

	rtn := C.CallReqRemoveParkedOrderAction(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqRemoveParkedOrderAction,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcRemoveParkedOrderActionField)(unsafe.Pointer(pRemoveParkedOrderAction)),
		C.int(nRequestID),
	)

	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"thost trader api ReqRemoveParkedOrderAction executed",
	)

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqExecOrderInsert(
	pInputExecOrder *thost.CThostFtdcInputExecOrderField, nRequestID int,
) int {
	slog.Log(context.Background(), slog.LevelDebug-2,
		"executing thost trader api ReqExecOrderInsert",
	)

	rtn := C.CallReqExecOrderInsert(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqExecOrderInsert,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcInputExecOrderField)(unsafe.Pointer(pInputExecOrder)),
		C.int(nRequestID),
	)

	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"thost trader api ReqExecOrderInsert executed",
	)

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqExecOrderAction(
	pInputExecOrderAction *thost.CThostFtdcInputExecOrderActionField, nRequestID int,
) int {
	slog.Log(context.Background(), slog.LevelDebug-2,
		"executing thost trader api ReqExecOrderAction",
	)

	rtn := C.CallReqExecOrderAction(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqExecOrderAction,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcInputExecOrderActionField)(unsafe.Pointer(pInputExecOrderAction)),
		C.int(nRequestID),
	)

	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"thost trader api ReqExecOrderAction executed",
	)

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqForQuoteInsert(
	pInputForQuote *thost.CThostFtdcInputForQuoteField, nRequestID int,
) int {
	slog.Log(context.Background(), slog.LevelDebug-2,
		"executing thost trader api ReqForQuoteInsert",
	)

	rtn := C.CallReqForQuoteInsert(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqForQuoteInsert,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcInputForQuoteField)(unsafe.Pointer(pInputForQuote)),
		C.int(nRequestID),
	)

	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"thost trader api ReqForQuoteInsert executed",
	)

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQuoteInsert(
	pInputQuote *thost.CThostFtdcInputQuoteField, nRequestID int,
) int {
	slog.Log(context.Background(), slog.LevelDebug-2,
		"executing thost trader api ReqQuoteInsert",
	)

	rtn := C.CallReqQuoteInsert(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQuoteInsert,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcInputQuoteField)(unsafe.Pointer(pInputQuote)),
		C.int(nRequestID),
	)

	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"thost trader api ReqQuoteInsert executed",
	)

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQuoteAction(
	pInputQuoteAction *thost.CThostFtdcInputQuoteActionField, nRequestID int,
) int {
	slog.Log(context.Background(), slog.LevelDebug-2,
		"executing thost trader api ReqQuoteAction",
	)

	rtn := C.CallReqQuoteAction(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQuoteAction,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcInputQuoteActionField)(unsafe.Pointer(pInputQuoteAction)),
		C.int(nRequestID),
	)

	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"thost trader api ReqQuoteAction executed",
	)

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqBatchOrderAction(
	pInputBatchOrderAction *thost.CThostFtdcInputBatchOrderActionField, nRequestID int,
) int {
	slog.Log(context.Background(), slog.LevelDebug-2,
		"executing thost trader api ReqBatchOrderAction",
	)

	rtn := C.CallReqBatchOrderAction(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqBatchOrderAction,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcInputBatchOrderActionField)(unsafe.Pointer(pInputBatchOrderAction)),
		C.int(nRequestID),
	)

	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"thost trader api ReqBatchOrderAction executed",
	)

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqOptionSelfCloseInsert(
	pInputOptionSelfClose *thost.CThostFtdcInputOptionSelfCloseField, nRequestID int,
) int {
	slog.Log(context.Background(), slog.LevelDebug-2,
		"executing thost trader api ReqOptionSelfCloseInsert",
	)

	rtn := C.CallReqOptionSelfCloseInsert(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqOptionSelfCloseInsert,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcInputOptionSelfCloseField)(unsafe.Pointer(pInputOptionSelfClose)),
		C.int(nRequestID),
	)

	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"thost trader api ReqOptionSelfCloseInsert executed",
	)

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqOptionSelfCloseAction(
	pInputOptionSelfCloseAction *thost.CThostFtdcInputOptionSelfCloseActionField, nRequestID int,
) int {
	slog.Log(context.Background(), slog.LevelDebug-2,
		"executing thost trader api ReqOptionSelfCloseAction",
	)

	rtn := C.CallReqOptionSelfCloseAction(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqOptionSelfCloseAction,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcInputOptionSelfCloseActionField)(unsafe.Pointer(pInputOptionSelfCloseAction)),
		C.int(nRequestID),
	)

	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"thost trader api ReqOptionSelfCloseAction executed",
	)

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqCombActionInsert(
	pInputCombAction *thost.CThostFtdcInputCombActionField, nRequestID int,
) int {
	slog.Log(context.Background(), slog.LevelDebug-2,
		"executing thost trader api ReqCombActionInsert",
	)

	rtn := C.CallReqCombActionInsert(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqCombActionInsert,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcInputCombActionField)(unsafe.Pointer(pInputCombAction)),
		C.int(nRequestID),
	)

	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"thost trader api ReqCombActionInsert executed",
	)

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQryOrder(
	pQryOrder *thost.CThostFtdcQryOrderField, nRequestID int,
) int {
	slog.Log(context.Background(), slog.LevelDebug-2,
		"executing thost trader api ReqQryOrder",
	)

	rtn := C.CallReqQryOrder(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQryOrder,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQryOrderField)(unsafe.Pointer(pQryOrder)),
		C.int(nRequestID),
	)

	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"thost trader api ReqQryOrder executed",
	)

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQryTrade(
	pQryTrade *thost.CThostFtdcQryTradeField, nRequestID int,
) int {
	slog.Log(context.Background(), slog.LevelDebug-2,
		"executing thost trader api ReqQryTrade",
	)

	rtn := C.CallReqQryTrade(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQryTrade,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQryTradeField)(unsafe.Pointer(pQryTrade)),
		C.int(nRequestID),
	)

	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"thost trader api ReqQryTrade executed",
	)

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQryInvestorPosition(
	pQryInvestorPosition *thost.CThostFtdcQryInvestorPositionField, nRequestID int,
) int {
	slog.Log(context.Background(), slog.LevelDebug-2,
		"executing thost trader api ReqQryInvestorPosition",
	)

	rtn := C.CallReqQryInvestorPosition(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQryInvestorPosition,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQryInvestorPositionField)(unsafe.Pointer(pQryInvestorPosition)),
		C.int(nRequestID),
	)

	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"thost trader api ReqQryInvestorPosition executed",
	)

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQryTradingAccount(
	pQryTradingAccount *thost.CThostFtdcQryTradingAccountField, nRequestID int,
) int {
	slog.Log(context.Background(), slog.LevelDebug-2,
		"executing thost trader api ReqQryTradingAccount",
	)

	rtn := C.CallReqQryTradingAccount(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQryTradingAccount,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQryTradingAccountField)(unsafe.Pointer(pQryTradingAccount)),
		C.int(nRequestID),
	)

	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"thost trader api ReqQryTradingAccount executed",
	)

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQryInvestor(
	pQryInvestor *thost.CThostFtdcQryInvestorField, nRequestID int,
) int {
	slog.Log(context.Background(), slog.LevelDebug-2,
		"executing thost trader api ReqQryInvestor",
	)

	rtn := C.CallReqQryInvestor(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQryInvestor,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQryInvestorField)(unsafe.Pointer(pQryInvestor)),
		C.int(nRequestID),
	)

	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"thost trader api ReqQryInvestor executed",
	)

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQryTradingCode(
	pQryTradingCode *thost.CThostFtdcQryTradingCodeField, nRequestID int,
) int {
	slog.Log(context.Background(), slog.LevelDebug-2,
		"executing thost trader api ReqQryTradingCode",
	)

	rtn := C.CallReqQryTradingCode(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQryTradingCode,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQryTradingCodeField)(unsafe.Pointer(pQryTradingCode)),
		C.int(nRequestID),
	)

	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"thost trader api ReqQryTradingCode executed",
	)

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQryInstrumentMarginRate(
	pQryInstrumentMarginRate *thost.CThostFtdcQryInstrumentMarginRateField, nRequestID int,
) int {
	slog.Log(context.Background(), slog.LevelDebug-2,
		"executing thost trader api ReqQryInstrumentMarginRate",
	)

	rtn := C.CallReqQryInstrumentMarginRate(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQryInstrumentMarginRate,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQryInstrumentMarginRateField)(unsafe.Pointer(pQryInstrumentMarginRate)),
		C.int(nRequestID),
	)

	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"thost trader api ReqQryInstrumentMarginRate executed",
	)

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQryInstrumentCommissionRate(
	pQryInstrumentCommissionRate *thost.CThostFtdcQryInstrumentCommissionRateField, nRequestID int,
) int {
	slog.Log(context.Background(), slog.LevelDebug-2,
		"executing thost trader api ReqQryInstrumentCommissionRate",
	)

	rtn := C.CallReqQryInstrumentCommissionRate(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQryInstrumentCommissionRate,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQryInstrumentCommissionRateField)(unsafe.Pointer(pQryInstrumentCommissionRate)),
		C.int(nRequestID),
	)

	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"thost trader api ReqQryInstrumentCommissionRate executed",
	)

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQryUserSession(
	pQryUserSession *thost.CThostFtdcQryUserSessionField, nRequestID int,
) int {
	slog.Log(context.Background(), slog.LevelDebug-2,
		"executing thost trader api ReqQryUserSession",
	)

	rtn := C.CallReqQryUserSession(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQryUserSession,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQryUserSessionField)(unsafe.Pointer(pQryUserSession)),
		C.int(nRequestID),
	)

	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"thost trader api ReqQryUserSession executed",
	)

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQryExchange(
	pQryExchange *thost.CThostFtdcQryExchangeField, nRequestID int,
) int {
	slog.Log(context.Background(), slog.LevelDebug-2,
		"executing thost trader api ReqQryExchange",
	)

	rtn := C.CallReqQryExchange(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQryExchange,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQryExchangeField)(unsafe.Pointer(pQryExchange)),
		C.int(nRequestID),
	)

	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"thost trader api ReqQryExchange executed",
	)

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQryProduct(
	pQryProduct *thost.CThostFtdcQryProductField, nRequestID int,
) int {
	slog.Log(context.Background(), slog.LevelDebug-2,
		"executing thost trader api ReqQryProduct",
	)

	rtn := C.CallReqQryProduct(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQryProduct,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQryProductField)(unsafe.Pointer(pQryProduct)),
		C.int(nRequestID),
	)

	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"thost trader api ReqQryProduct executed",
	)

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQryInstrument(
	pQryInstrument *thost.CThostFtdcQryInstrumentField, nRequestID int,
) int {
	slog.Log(context.Background(), slog.LevelDebug-2,
		"executing thost trader api ReqQryInstrument",
	)

	rtn := C.CallReqQryInstrument(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQryInstrument,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQryInstrumentField)(unsafe.Pointer(pQryInstrument)),
		C.int(nRequestID),
	)

	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"thost trader api ReqQryInstrument executed",
	)

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQryDepthMarketData(
	pQryDepthMarketData *thost.CThostFtdcQryDepthMarketDataField, nRequestID int,
) int {
	slog.Log(context.Background(), slog.LevelDebug-2,
		"executing thost trader api ReqQryDepthMarketData",
	)

	rtn := C.CallReqQryDepthMarketData(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQryDepthMarketData,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQryDepthMarketDataField)(unsafe.Pointer(pQryDepthMarketData)),
		C.int(nRequestID),
	)

	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"thost trader api ReqQryDepthMarketData executed",
	)

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQryTraderOffer(
	pQryTraderOffer *thost.CThostFtdcQryTraderOfferField, nRequestID int,
) int {
	slog.Log(context.Background(), slog.LevelDebug-2,
		"executing thost trader api ReqQryTraderOffer",
	)

	rtn := C.CallReqQryTraderOffer(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQryTraderOffer,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQryTraderOfferField)(unsafe.Pointer(pQryTraderOffer)),
		C.int(nRequestID),
	)

	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"thost trader api ReqQryTraderOffer executed",
	)

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQrySettlementInfo(
	pQrySettlementInfo *thost.CThostFtdcQrySettlementInfoField, nRequestID int,
) int {
	slog.Log(context.Background(), slog.LevelDebug-2,
		"executing thost trader api ReqQrySettlementInfo",
	)

	rtn := C.CallReqQrySettlementInfo(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQrySettlementInfo,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQrySettlementInfoField)(unsafe.Pointer(pQrySettlementInfo)),
		C.int(nRequestID),
	)

	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"thost trader api ReqQrySettlementInfo executed",
	)

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQryTransferBank(
	pQryTransferBank *thost.CThostFtdcQryTransferBankField, nRequestID int,
) int {
	slog.Log(context.Background(), slog.LevelDebug-2,
		"executing thost trader api ReqQryTransferBank",
	)

	rtn := C.CallReqQryTransferBank(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQryTransferBank,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQryTransferBankField)(unsafe.Pointer(pQryTransferBank)),
		C.int(nRequestID),
	)

	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"thost trader api ReqQryTransferBank executed",
	)

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQryInvestorPositionDetail(
	pQryInvestorPositionDetail *thost.CThostFtdcQryInvestorPositionDetailField, nRequestID int,
) int {
	slog.Log(context.Background(), slog.LevelDebug-2,
		"executing thost trader api ReqQryInvestorPositionDetail",
	)

	rtn := C.CallReqQryInvestorPositionDetail(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQryInvestorPositionDetail,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQryInvestorPositionDetailField)(unsafe.Pointer(pQryInvestorPositionDetail)),
		C.int(nRequestID),
	)

	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"thost trader api ReqQryInvestorPositionDetail executed",
	)

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQryNotice(
	pQryNotice *thost.CThostFtdcQryNoticeField, nRequestID int,
) int {
	slog.Log(context.Background(), slog.LevelDebug-2,
		"executing thost trader api ReqQryNotice",
	)

	rtn := C.CallReqQryNotice(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQryNotice,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQryNoticeField)(unsafe.Pointer(pQryNotice)),
		C.int(nRequestID),
	)

	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"thost trader api ReqQryNotice executed",
	)

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQrySettlementInfoConfirm(
	pQrySettlementInfoConfirm *thost.CThostFtdcQrySettlementInfoConfirmField, nRequestID int,
) int {
	slog.Log(context.Background(), slog.LevelDebug-2,
		"executing thost trader api ReqQrySettlementInfoConfirm",
	)

	rtn := C.CallReqQrySettlementInfoConfirm(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQrySettlementInfoConfirm,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQrySettlementInfoConfirmField)(unsafe.Pointer(pQrySettlementInfoConfirm)),
		C.int(nRequestID),
	)

	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"thost trader api ReqQrySettlementInfoConfirm executed",
	)

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQryInvestorPositionCombineDetail(
	pQryInvestorPositionCombineDetail *thost.CThostFtdcQryInvestorPositionCombineDetailField, nRequestID int,
) int {
	slog.Log(context.Background(), slog.LevelDebug-2,
		"executing thost trader api ReqQryInvestorPositionCombineDetail",
	)

	rtn := C.CallReqQryInvestorPositionCombineDetail(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQryInvestorPositionCombineDetail,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQryInvestorPositionCombineDetailField)(unsafe.Pointer(pQryInvestorPositionCombineDetail)),
		C.int(nRequestID),
	)

	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"thost trader api ReqQryInvestorPositionCombineDetail executed",
	)

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQryCFMMCTradingAccountKey(
	pQryCFMMCTradingAccountKey *thost.CThostFtdcQryCFMMCTradingAccountKeyField, nRequestID int,
) int {
	slog.Log(context.Background(), slog.LevelDebug-2,
		"executing thost trader api ReqQryCFMMCTradingAccountKey",
	)

	rtn := C.CallReqQryCFMMCTradingAccountKey(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQryCFMMCTradingAccountKey,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQryCFMMCTradingAccountKeyField)(unsafe.Pointer(pQryCFMMCTradingAccountKey)),
		C.int(nRequestID),
	)

	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"thost trader api ReqQryCFMMCTradingAccountKey executed",
	)

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQryEWarrantOffset(
	pQryEWarrantOffset *thost.CThostFtdcQryEWarrantOffsetField, nRequestID int,
) int {
	slog.Log(context.Background(), slog.LevelDebug-2,
		"executing thost trader api ReqQryEWarrantOffset",
	)

	rtn := C.CallReqQryEWarrantOffset(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQryEWarrantOffset,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQryEWarrantOffsetField)(unsafe.Pointer(pQryEWarrantOffset)),
		C.int(nRequestID),
	)

	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"thost trader api ReqQryEWarrantOffset executed",
	)

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQryInvestorProductGroupMargin(
	pQryInvestorProductGroupMargin *thost.CThostFtdcQryInvestorProductGroupMarginField, nRequestID int,
) int {
	slog.Log(context.Background(), slog.LevelDebug-2,
		"executing thost trader api ReqQryInvestorProductGroupMargin",
	)

	rtn := C.CallReqQryInvestorProductGroupMargin(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQryInvestorProductGroupMargin,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQryInvestorProductGroupMarginField)(unsafe.Pointer(pQryInvestorProductGroupMargin)),
		C.int(nRequestID),
	)

	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"thost trader api ReqQryInvestorProductGroupMargin executed",
	)

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQryExchangeMarginRate(
	pQryExchangeMarginRate *thost.CThostFtdcQryExchangeMarginRateField, nRequestID int,
) int {
	slog.Log(context.Background(), slog.LevelDebug-2,
		"executing thost trader api ReqQryExchangeMarginRate",
	)

	rtn := C.CallReqQryExchangeMarginRate(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQryExchangeMarginRate,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQryExchangeMarginRateField)(unsafe.Pointer(pQryExchangeMarginRate)),
		C.int(nRequestID),
	)

	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"thost trader api ReqQryExchangeMarginRate executed",
	)

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQryExchangeMarginRateAdjust(
	pQryExchangeMarginRateAdjust *thost.CThostFtdcQryExchangeMarginRateAdjustField, nRequestID int,
) int {
	slog.Log(context.Background(), slog.LevelDebug-2,
		"executing thost trader api ReqQryExchangeMarginRateAdjust",
	)

	rtn := C.CallReqQryExchangeMarginRateAdjust(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQryExchangeMarginRateAdjust,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQryExchangeMarginRateAdjustField)(unsafe.Pointer(pQryExchangeMarginRateAdjust)),
		C.int(nRequestID),
	)

	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"thost trader api ReqQryExchangeMarginRateAdjust executed",
	)

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQryExchangeRate(
	pQryExchangeRate *thost.CThostFtdcQryExchangeRateField, nRequestID int,
) int {
	slog.Log(context.Background(), slog.LevelDebug-2,
		"executing thost trader api ReqQryExchangeRate",
	)

	rtn := C.CallReqQryExchangeRate(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQryExchangeRate,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQryExchangeRateField)(unsafe.Pointer(pQryExchangeRate)),
		C.int(nRequestID),
	)

	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"thost trader api ReqQryExchangeRate executed",
	)

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQrySecAgentACIDMap(
	pQrySecAgentACIDMap *thost.CThostFtdcQrySecAgentACIDMapField, nRequestID int,
) int {
	slog.Log(context.Background(), slog.LevelDebug-2,
		"executing thost trader api ReqQrySecAgentACIDMap",
	)

	rtn := C.CallReqQrySecAgentACIDMap(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQrySecAgentACIDMap,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQrySecAgentACIDMapField)(unsafe.Pointer(pQrySecAgentACIDMap)),
		C.int(nRequestID),
	)

	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"thost trader api ReqQrySecAgentACIDMap executed",
	)

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQryProductExchRate(
	pQryProductExchRate *thost.CThostFtdcQryProductExchRateField, nRequestID int,
) int {
	slog.Log(context.Background(), slog.LevelDebug-2,
		"executing thost trader api ReqQryProductExchRate",
	)

	rtn := C.CallReqQryProductExchRate(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQryProductExchRate,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQryProductExchRateField)(unsafe.Pointer(pQryProductExchRate)),
		C.int(nRequestID),
	)

	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"thost trader api ReqQryProductExchRate executed",
	)

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQryProductGroup(
	pQryProductGroup *thost.CThostFtdcQryProductGroupField, nRequestID int,
) int {
	slog.Log(context.Background(), slog.LevelDebug-2,
		"executing thost trader api ReqQryProductGroup",
	)

	rtn := C.CallReqQryProductGroup(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQryProductGroup,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQryProductGroupField)(unsafe.Pointer(pQryProductGroup)),
		C.int(nRequestID),
	)

	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"thost trader api ReqQryProductGroup executed",
	)

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQryMMInstrumentCommissionRate(
	pQryMMInstrumentCommissionRate *thost.CThostFtdcQryMMInstrumentCommissionRateField, nRequestID int,
) int {
	slog.Log(context.Background(), slog.LevelDebug-2,
		"executing thost trader api ReqQryMMInstrumentCommissionRate",
	)

	rtn := C.CallReqQryMMInstrumentCommissionRate(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQryMMInstrumentCommissionRate,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQryMMInstrumentCommissionRateField)(unsafe.Pointer(pQryMMInstrumentCommissionRate)),
		C.int(nRequestID),
	)

	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"thost trader api ReqQryMMInstrumentCommissionRate executed",
	)

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQryMMOptionInstrCommRate(
	pQryMMOptionInstrCommRate *thost.CThostFtdcQryMMOptionInstrCommRateField, nRequestID int,
) int {
	slog.Log(context.Background(), slog.LevelDebug-2,
		"executing thost trader api ReqQryMMOptionInstrCommRate",
	)

	rtn := C.CallReqQryMMOptionInstrCommRate(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQryMMOptionInstrCommRate,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQryMMOptionInstrCommRateField)(unsafe.Pointer(pQryMMOptionInstrCommRate)),
		C.int(nRequestID),
	)

	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"thost trader api ReqQryMMOptionInstrCommRate executed",
	)

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQryInstrumentOrderCommRate(
	pQryInstrumentOrderCommRate *thost.CThostFtdcQryInstrumentOrderCommRateField, nRequestID int,
) int {
	slog.Log(context.Background(), slog.LevelDebug-2,
		"executing thost trader api ReqQryInstrumentOrderCommRate",
	)

	rtn := C.CallReqQryInstrumentOrderCommRate(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQryInstrumentOrderCommRate,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQryInstrumentOrderCommRateField)(unsafe.Pointer(pQryInstrumentOrderCommRate)),
		C.int(nRequestID),
	)

	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"thost trader api ReqQryInstrumentOrderCommRate executed",
	)

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQrySecAgentTradingAccount(
	pQryTradingAccount *thost.CThostFtdcQryTradingAccountField, nRequestID int,
) int {
	slog.Log(context.Background(), slog.LevelDebug-2,
		"executing thost trader api ReqQrySecAgentTradingAccount",
	)

	rtn := C.CallReqQrySecAgentTradingAccount(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQrySecAgentTradingAccount,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQryTradingAccountField)(unsafe.Pointer(pQryTradingAccount)),
		C.int(nRequestID),
	)

	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"thost trader api ReqQrySecAgentTradingAccount executed",
	)

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQrySecAgentCheckMode(
	pQrySecAgentCheckMode *thost.CThostFtdcQrySecAgentCheckModeField, nRequestID int,
) int {
	slog.Log(context.Background(), slog.LevelDebug-2,
		"executing thost trader api ReqQrySecAgentCheckMode",
	)

	rtn := C.CallReqQrySecAgentCheckMode(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQrySecAgentCheckMode,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQrySecAgentCheckModeField)(unsafe.Pointer(pQrySecAgentCheckMode)),
		C.int(nRequestID),
	)

	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"thost trader api ReqQrySecAgentCheckMode executed",
	)

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQrySecAgentTradeInfo(
	pQrySecAgentTradeInfo *thost.CThostFtdcQrySecAgentTradeInfoField, nRequestID int,
) int {
	slog.Log(context.Background(), slog.LevelDebug-2,
		"executing thost trader api ReqQrySecAgentTradeInfo",
	)

	rtn := C.CallReqQrySecAgentTradeInfo(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQrySecAgentTradeInfo,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQrySecAgentTradeInfoField)(unsafe.Pointer(pQrySecAgentTradeInfo)),
		C.int(nRequestID),
	)

	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"thost trader api ReqQrySecAgentTradeInfo executed",
	)

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQryOptionInstrTradeCost(
	pQryOptionInstrTradeCost *thost.CThostFtdcQryOptionInstrTradeCostField, nRequestID int,
) int {
	slog.Log(context.Background(), slog.LevelDebug-2,
		"executing thost trader api ReqQryOptionInstrTradeCost",
	)

	rtn := C.CallReqQryOptionInstrTradeCost(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQryOptionInstrTradeCost,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQryOptionInstrTradeCostField)(unsafe.Pointer(pQryOptionInstrTradeCost)),
		C.int(nRequestID),
	)

	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"thost trader api ReqQryOptionInstrTradeCost executed",
	)

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQryOptionInstrCommRate(
	pQryOptionInstrCommRate *thost.CThostFtdcQryOptionInstrCommRateField, nRequestID int,
) int {
	slog.Log(context.Background(), slog.LevelDebug-2,
		"executing thost trader api ReqQryOptionInstrCommRate",
	)

	rtn := C.CallReqQryOptionInstrCommRate(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQryOptionInstrCommRate,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQryOptionInstrCommRateField)(unsafe.Pointer(pQryOptionInstrCommRate)),
		C.int(nRequestID),
	)

	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"thost trader api ReqQryOptionInstrCommRate executed",
	)

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQryExecOrder(
	pQryExecOrder *thost.CThostFtdcQryExecOrderField, nRequestID int,
) int {
	slog.Log(context.Background(), slog.LevelDebug-2,
		"executing thost trader api ReqQryExecOrder",
	)

	rtn := C.CallReqQryExecOrder(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQryExecOrder,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQryExecOrderField)(unsafe.Pointer(pQryExecOrder)),
		C.int(nRequestID),
	)

	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"thost trader api ReqQryExecOrder executed",
	)

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQryForQuote(
	pQryForQuote *thost.CThostFtdcQryForQuoteField, nRequestID int,
) int {
	slog.Log(context.Background(), slog.LevelDebug-2,
		"executing thost trader api ReqQryForQuote",
	)

	rtn := C.CallReqQryForQuote(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQryForQuote,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQryForQuoteField)(unsafe.Pointer(pQryForQuote)),
		C.int(nRequestID),
	)

	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"thost trader api ReqQryForQuote executed",
	)

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQryQuote(
	pQryQuote *thost.CThostFtdcQryQuoteField, nRequestID int,
) int {
	slog.Log(context.Background(), slog.LevelDebug-2,
		"executing thost trader api ReqQryQuote",
	)

	rtn := C.CallReqQryQuote(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQryQuote,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQryQuoteField)(unsafe.Pointer(pQryQuote)),
		C.int(nRequestID),
	)

	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"thost trader api ReqQryQuote executed",
	)

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQryOptionSelfClose(
	pQryOptionSelfClose *thost.CThostFtdcQryOptionSelfCloseField, nRequestID int,
) int {
	slog.Log(context.Background(), slog.LevelDebug-2,
		"executing thost trader api ReqQryOptionSelfClose",
	)

	rtn := C.CallReqQryOptionSelfClose(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQryOptionSelfClose,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQryOptionSelfCloseField)(unsafe.Pointer(pQryOptionSelfClose)),
		C.int(nRequestID),
	)

	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"thost trader api ReqQryOptionSelfClose executed",
	)

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQryInvestUnit(
	pQryInvestUnit *thost.CThostFtdcQryInvestUnitField, nRequestID int,
) int {
	slog.Log(context.Background(), slog.LevelDebug-2,
		"executing thost trader api ReqQryInvestUnit",
	)

	rtn := C.CallReqQryInvestUnit(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQryInvestUnit,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQryInvestUnitField)(unsafe.Pointer(pQryInvestUnit)),
		C.int(nRequestID),
	)

	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"thost trader api ReqQryInvestUnit executed",
	)

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQryCombInstrumentGuard(
	pQryCombInstrumentGuard *thost.CThostFtdcQryCombInstrumentGuardField, nRequestID int,
) int {
	slog.Log(context.Background(), slog.LevelDebug-2,
		"executing thost trader api ReqQryCombInstrumentGuard",
	)

	rtn := C.CallReqQryCombInstrumentGuard(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQryCombInstrumentGuard,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQryCombInstrumentGuardField)(unsafe.Pointer(pQryCombInstrumentGuard)),
		C.int(nRequestID),
	)

	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"thost trader api ReqQryCombInstrumentGuard executed",
	)

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQryCombAction(
	pQryCombAction *thost.CThostFtdcQryCombActionField, nRequestID int,
) int {
	slog.Log(context.Background(), slog.LevelDebug-2,
		"executing thost trader api ReqQryCombAction",
	)

	rtn := C.CallReqQryCombAction(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQryCombAction,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQryCombActionField)(unsafe.Pointer(pQryCombAction)),
		C.int(nRequestID),
	)

	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"thost trader api ReqQryCombAction executed",
	)

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQryTransferSerial(
	pQryTransferSerial *thost.CThostFtdcQryTransferSerialField, nRequestID int,
) int {
	slog.Log(context.Background(), slog.LevelDebug-2,
		"executing thost trader api ReqQryTransferSerial",
	)

	rtn := C.CallReqQryTransferSerial(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQryTransferSerial,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQryTransferSerialField)(unsafe.Pointer(pQryTransferSerial)),
		C.int(nRequestID),
	)

	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"thost trader api ReqQryTransferSerial executed",
	)

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQryAccountregister(
	pQryAccountregister *thost.CThostFtdcQryAccountregisterField, nRequestID int,
) int {
	slog.Log(context.Background(), slog.LevelDebug-2,
		"executing thost trader api ReqQryAccountregister",
	)

	rtn := C.CallReqQryAccountregister(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQryAccountregister,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQryAccountregisterField)(unsafe.Pointer(pQryAccountregister)),
		C.int(nRequestID),
	)

	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"thost trader api ReqQryAccountregister executed",
	)

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQryContractBank(
	pQryContractBank *thost.CThostFtdcQryContractBankField, nRequestID int,
) int {
	slog.Log(context.Background(), slog.LevelDebug-2,
		"executing thost trader api ReqQryContractBank",
	)

	rtn := C.CallReqQryContractBank(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQryContractBank,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQryContractBankField)(unsafe.Pointer(pQryContractBank)),
		C.int(nRequestID),
	)

	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"thost trader api ReqQryContractBank executed",
	)

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQryParkedOrder(
	pQryParkedOrder *thost.CThostFtdcQryParkedOrderField, nRequestID int,
) int {
	slog.Log(context.Background(), slog.LevelDebug-2,
		"executing thost trader api ReqQryParkedOrder",
	)

	rtn := C.CallReqQryParkedOrder(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQryParkedOrder,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQryParkedOrderField)(unsafe.Pointer(pQryParkedOrder)),
		C.int(nRequestID),
	)

	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"thost trader api ReqQryParkedOrder executed",
	)

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQryParkedOrderAction(
	pQryParkedOrderAction *thost.CThostFtdcQryParkedOrderActionField, nRequestID int,
) int {
	slog.Log(context.Background(), slog.LevelDebug-2,
		"executing thost trader api ReqQryParkedOrderAction",
	)

	rtn := C.CallReqQryParkedOrderAction(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQryParkedOrderAction,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQryParkedOrderActionField)(unsafe.Pointer(pQryParkedOrderAction)),
		C.int(nRequestID),
	)

	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"thost trader api ReqQryParkedOrderAction executed",
	)

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQryTradingNotice(
	pQryTradingNotice *thost.CThostFtdcQryTradingNoticeField, nRequestID int,
) int {
	slog.Log(context.Background(), slog.LevelDebug-2,
		"executing thost trader api ReqQryTradingNotice",
	)

	rtn := C.CallReqQryTradingNotice(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQryTradingNotice,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQryTradingNoticeField)(unsafe.Pointer(pQryTradingNotice)),
		C.int(nRequestID),
	)

	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"thost trader api ReqQryTradingNotice executed",
	)

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQryBrokerTradingParams(
	pQryBrokerTradingParams *thost.CThostFtdcQryBrokerTradingParamsField, nRequestID int,
) int {
	slog.Log(context.Background(), slog.LevelDebug-2,
		"executing thost trader api ReqQryBrokerTradingParams",
	)

	rtn := C.CallReqQryBrokerTradingParams(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQryBrokerTradingParams,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQryBrokerTradingParamsField)(unsafe.Pointer(pQryBrokerTradingParams)),
		C.int(nRequestID),
	)

	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"thost trader api ReqQryBrokerTradingParams executed",
	)

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQryBrokerTradingAlgos(
	pQryBrokerTradingAlgos *thost.CThostFtdcQryBrokerTradingAlgosField, nRequestID int,
) int {
	slog.Log(context.Background(), slog.LevelDebug-2,
		"executing thost trader api ReqQryBrokerTradingAlgos",
	)

	rtn := C.CallReqQryBrokerTradingAlgos(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQryBrokerTradingAlgos,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQryBrokerTradingAlgosField)(unsafe.Pointer(pQryBrokerTradingAlgos)),
		C.int(nRequestID),
	)

	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"thost trader api ReqQryBrokerTradingAlgos executed",
	)

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQueryCFMMCTradingAccountToken(
	pQueryCFMMCTradingAccountToken *thost.CThostFtdcQueryCFMMCTradingAccountTokenField, nRequestID int,
) int {
	slog.Log(context.Background(), slog.LevelDebug-2,
		"executing thost trader api ReqQueryCFMMCTradingAccountToken",
	)

	rtn := C.CallReqQueryCFMMCTradingAccountToken(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQueryCFMMCTradingAccountToken,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQueryCFMMCTradingAccountTokenField)(unsafe.Pointer(pQueryCFMMCTradingAccountToken)),
		C.int(nRequestID),
	)

	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"thost trader api ReqQueryCFMMCTradingAccountToken executed",
	)

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqFromBankToFutureByFuture(
	pReqTransfer *thost.CThostFtdcReqTransferField, nRequestID int,
) int {
	slog.Log(context.Background(), slog.LevelDebug-2,
		"executing thost trader api ReqFromBankToFutureByFuture",
	)

	rtn := C.CallReqFromBankToFutureByFuture(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqFromBankToFutureByFuture,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcReqTransferField)(unsafe.Pointer(pReqTransfer)),
		C.int(nRequestID),
	)

	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"thost trader api ReqFromBankToFutureByFuture executed",
	)

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqFromFutureToBankByFuture(
	pReqTransfer *thost.CThostFtdcReqTransferField, nRequestID int,
) int {
	slog.Log(context.Background(), slog.LevelDebug-2,
		"executing thost trader api ReqFromFutureToBankByFuture",
	)

	rtn := C.CallReqFromFutureToBankByFuture(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqFromFutureToBankByFuture,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcReqTransferField)(unsafe.Pointer(pReqTransfer)),
		C.int(nRequestID),
	)

	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"thost trader api ReqFromFutureToBankByFuture executed",
	)

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQueryBankAccountMoneyByFuture(
	pReqQueryAccount *thost.CThostFtdcReqQueryAccountField, nRequestID int,
) int {
	slog.Log(context.Background(), slog.LevelDebug-2,
		"executing thost trader api ReqQueryBankAccountMoneyByFuture",
	)

	rtn := C.CallReqQueryBankAccountMoneyByFuture(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQueryBankAccountMoneyByFuture,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcReqQueryAccountField)(unsafe.Pointer(pReqQueryAccount)),
		C.int(nRequestID),
	)

	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"thost trader api ReqQueryBankAccountMoneyByFuture executed",
	)

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQryClassifiedInstrument(
	pQryClassifiedInstrument *thost.CThostFtdcQryClassifiedInstrumentField, nRequestID int,
) int {
	slog.Log(context.Background(), slog.LevelDebug-2,
		"executing thost trader api ReqQryClassifiedInstrument",
	)

	rtn := C.CallReqQryClassifiedInstrument(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQryClassifiedInstrument,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQryClassifiedInstrumentField)(unsafe.Pointer(pQryClassifiedInstrument)),
		C.int(nRequestID),
	)

	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"thost trader api ReqQryClassifiedInstrument executed",
	)

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQryCombPromotionParam(
	pQryCombPromotionParam *thost.CThostFtdcQryCombPromotionParamField, nRequestID int,
) int {
	slog.Log(context.Background(), slog.LevelDebug-2,
		"executing thost trader api ReqQryCombPromotionParam",
	)

	rtn := C.CallReqQryCombPromotionParam(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQryCombPromotionParam,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQryCombPromotionParamField)(unsafe.Pointer(pQryCombPromotionParam)),
		C.int(nRequestID),
	)

	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"thost trader api ReqQryCombPromotionParam executed",
	)

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQryRiskSettleInvstPosition(
	pQryRiskSettleInvstPosition *thost.CThostFtdcQryRiskSettleInvstPositionField, nRequestID int,
) int {
	slog.Log(context.Background(), slog.LevelDebug-2,
		"executing thost trader api ReqQryRiskSettleInvstPosition",
	)

	rtn := C.CallReqQryRiskSettleInvstPosition(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQryRiskSettleInvstPosition,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQryRiskSettleInvstPositionField)(unsafe.Pointer(pQryRiskSettleInvstPosition)),
		C.int(nRequestID),
	)

	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"thost trader api ReqQryRiskSettleInvstPosition executed",
	)

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQryRiskSettleProductStatus(
	pQryRiskSettleProductStatus *thost.CThostFtdcQryRiskSettleProductStatusField, nRequestID int,
) int {
	slog.Log(context.Background(), slog.LevelDebug-2,
		"executing thost trader api ReqQryRiskSettleProductStatus",
	)

	rtn := C.CallReqQryRiskSettleProductStatus(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQryRiskSettleProductStatus,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQryRiskSettleProductStatusField)(unsafe.Pointer(pQryRiskSettleProductStatus)),
		C.int(nRequestID),
	)

	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"thost trader api ReqQryRiskSettleProductStatus executed",
	)

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQrySPBMFutureParameter(
	pQrySPBMFutureParameter *thost.CThostFtdcQrySPBMFutureParameterField, nRequestID int,
) int {
	slog.Log(context.Background(), slog.LevelDebug-2,
		"executing thost trader api ReqQrySPBMFutureParameter",
	)

	rtn := C.CallReqQrySPBMFutureParameter(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQrySPBMFutureParameter,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQrySPBMFutureParameterField)(unsafe.Pointer(pQrySPBMFutureParameter)),
		C.int(nRequestID),
	)

	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"thost trader api ReqQrySPBMFutureParameter executed",
	)

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQrySPBMOptionParameter(
	pQrySPBMOptionParameter *thost.CThostFtdcQrySPBMOptionParameterField, nRequestID int,
) int {
	slog.Log(context.Background(), slog.LevelDebug-2,
		"executing thost trader api ReqQrySPBMOptionParameter",
	)

	rtn := C.CallReqQrySPBMOptionParameter(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQrySPBMOptionParameter,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQrySPBMOptionParameterField)(unsafe.Pointer(pQrySPBMOptionParameter)),
		C.int(nRequestID),
	)

	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"thost trader api ReqQrySPBMOptionParameter executed",
	)

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQrySPBMIntraParameter(
	pQrySPBMIntraParameter *thost.CThostFtdcQrySPBMIntraParameterField, nRequestID int,
) int {
	slog.Log(context.Background(), slog.LevelDebug-2,
		"executing thost trader api ReqQrySPBMIntraParameter",
	)

	rtn := C.CallReqQrySPBMIntraParameter(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQrySPBMIntraParameter,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQrySPBMIntraParameterField)(unsafe.Pointer(pQrySPBMIntraParameter)),
		C.int(nRequestID),
	)

	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"thost trader api ReqQrySPBMIntraParameter executed",
	)

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQrySPBMInterParameter(
	pQrySPBMInterParameter *thost.CThostFtdcQrySPBMInterParameterField, nRequestID int,
) int {
	slog.Log(context.Background(), slog.LevelDebug-2,
		"executing thost trader api ReqQrySPBMInterParameter",
	)

	rtn := C.CallReqQrySPBMInterParameter(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQrySPBMInterParameter,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQrySPBMInterParameterField)(unsafe.Pointer(pQrySPBMInterParameter)),
		C.int(nRequestID),
	)

	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"thost trader api ReqQrySPBMInterParameter executed",
	)

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQrySPBMPortfDefinition(
	pQrySPBMPortfDefinition *thost.CThostFtdcQrySPBMPortfDefinitionField, nRequestID int,
) int {
	slog.Log(context.Background(), slog.LevelDebug-2,
		"executing thost trader api ReqQrySPBMPortfDefinition",
	)

	rtn := C.CallReqQrySPBMPortfDefinition(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQrySPBMPortfDefinition,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQrySPBMPortfDefinitionField)(unsafe.Pointer(pQrySPBMPortfDefinition)),
		C.int(nRequestID),
	)

	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"thost trader api ReqQrySPBMPortfDefinition executed",
	)

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQrySPBMInvestorPortfDef(
	pQrySPBMInvestorPortfDef *thost.CThostFtdcQrySPBMInvestorPortfDefField, nRequestID int,
) int {
	slog.Log(context.Background(), slog.LevelDebug-2,
		"executing thost trader api ReqQrySPBMInvestorPortfDef",
	)

	rtn := C.CallReqQrySPBMInvestorPortfDef(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQrySPBMInvestorPortfDef,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQrySPBMInvestorPortfDefField)(unsafe.Pointer(pQrySPBMInvestorPortfDef)),
		C.int(nRequestID),
	)

	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"thost trader api ReqQrySPBMInvestorPortfDef executed",
	)

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQryInvestorPortfMarginRatio(
	pQryInvestorPortfMarginRatio *thost.CThostFtdcQryInvestorPortfMarginRatioField, nRequestID int,
) int {
	slog.Log(context.Background(), slog.LevelDebug-2,
		"executing thost trader api ReqQryInvestorPortfMarginRatio",
	)

	rtn := C.CallReqQryInvestorPortfMarginRatio(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQryInvestorPortfMarginRatio,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQryInvestorPortfMarginRatioField)(unsafe.Pointer(pQryInvestorPortfMarginRatio)),
		C.int(nRequestID),
	)

	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"thost trader api ReqQryInvestorPortfMarginRatio executed",
	)

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQryInvestorProdSPBMDetail(
	pQryInvestorProdSPBMDetail *thost.CThostFtdcQryInvestorProdSPBMDetailField, nRequestID int,
) int {
	slog.Log(context.Background(), slog.LevelDebug-2,
		"executing thost trader api ReqQryInvestorProdSPBMDetail",
	)

	rtn := C.CallReqQryInvestorProdSPBMDetail(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQryInvestorProdSPBMDetail,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQryInvestorProdSPBMDetailField)(unsafe.Pointer(pQryInvestorProdSPBMDetail)),
		C.int(nRequestID),
	)

	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"thost trader api ReqQryInvestorProdSPBMDetail executed",
	)

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQryInvestorCommoditySPMMMargin(
	pQryInvestorCommoditySPMMMargin *thost.CThostFtdcQryInvestorCommoditySPMMMarginField, nRequestID int,
) int {
	slog.Log(context.Background(), slog.LevelDebug-2,
		"executing thost trader api ReqQryInvestorCommoditySPMMMargin",
	)

	rtn := C.CallReqQryInvestorCommoditySPMMMargin(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQryInvestorCommoditySPMMMargin,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQryInvestorCommoditySPMMMarginField)(unsafe.Pointer(pQryInvestorCommoditySPMMMargin)),
		C.int(nRequestID),
	)

	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"thost trader api ReqQryInvestorCommoditySPMMMargin executed",
	)

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQryInvestorCommodityGroupSPMMMargin(
	pQryInvestorCommodityGroupSPMMMargin *thost.CThostFtdcQryInvestorCommodityGroupSPMMMarginField, nRequestID int,
) int {
	slog.Log(context.Background(), slog.LevelDebug-2,
		"executing thost trader api ReqQryInvestorCommodityGroupSPMMMargin",
	)

	rtn := C.CallReqQryInvestorCommodityGroupSPMMMargin(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQryInvestorCommodityGroupSPMMMargin,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQryInvestorCommodityGroupSPMMMarginField)(unsafe.Pointer(pQryInvestorCommodityGroupSPMMMargin)),
		C.int(nRequestID),
	)

	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"thost trader api ReqQryInvestorCommodityGroupSPMMMargin executed",
	)

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQrySPMMInstParam(
	pQrySPMMInstParam *thost.CThostFtdcQrySPMMInstParamField, nRequestID int,
) int {
	slog.Log(context.Background(), slog.LevelDebug-2,
		"executing thost trader api ReqQrySPMMInstParam",
	)

	rtn := C.CallReqQrySPMMInstParam(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQrySPMMInstParam,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQrySPMMInstParamField)(unsafe.Pointer(pQrySPMMInstParam)),
		C.int(nRequestID),
	)

	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"thost trader api ReqQrySPMMInstParam executed",
	)

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQrySPMMProductParam(
	pQrySPMMProductParam *thost.CThostFtdcQrySPMMProductParamField, nRequestID int,
) int {
	slog.Log(context.Background(), slog.LevelDebug-2,
		"executing thost trader api ReqQrySPMMProductParam",
	)

	rtn := C.CallReqQrySPMMProductParam(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQrySPMMProductParam,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQrySPMMProductParamField)(unsafe.Pointer(pQrySPMMProductParam)),
		C.int(nRequestID),
	)

	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"thost trader api ReqQrySPMMProductParam executed",
	)

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQrySPBMAddOnInterParameter(
	pQrySPBMAddOnInterParameter *thost.CThostFtdcQrySPBMAddOnInterParameterField, nRequestID int,
) int {
	slog.Log(context.Background(), slog.LevelDebug-2,
		"executing thost trader api ReqQrySPBMAddOnInterParameter",
	)

	rtn := C.CallReqQrySPBMAddOnInterParameter(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQrySPBMAddOnInterParameter,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQrySPBMAddOnInterParameterField)(unsafe.Pointer(pQrySPBMAddOnInterParameter)),
		C.int(nRequestID),
	)

	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"thost trader api ReqQrySPBMAddOnInterParameter executed",
	)

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQryRCAMSCombProductInfo(
	pQryRCAMSCombProductInfo *thost.CThostFtdcQryRCAMSCombProductInfoField, nRequestID int,
) int {
	slog.Log(context.Background(), slog.LevelDebug-2,
		"executing thost trader api ReqQryRCAMSCombProductInfo",
	)

	rtn := C.CallReqQryRCAMSCombProductInfo(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQryRCAMSCombProductInfo,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQryRCAMSCombProductInfoField)(unsafe.Pointer(pQryRCAMSCombProductInfo)),
		C.int(nRequestID),
	)

	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"thost trader api ReqQryRCAMSCombProductInfo executed",
	)

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQryRCAMSInstrParameter(
	pQryRCAMSInstrParameter *thost.CThostFtdcQryRCAMSInstrParameterField, nRequestID int,
) int {
	slog.Log(context.Background(), slog.LevelDebug-2,
		"executing thost trader api ReqQryRCAMSInstrParameter",
	)

	rtn := C.CallReqQryRCAMSInstrParameter(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQryRCAMSInstrParameter,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQryRCAMSInstrParameterField)(unsafe.Pointer(pQryRCAMSInstrParameter)),
		C.int(nRequestID),
	)

	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"thost trader api ReqQryRCAMSInstrParameter executed",
	)

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQryRCAMSIntraParameter(
	pQryRCAMSIntraParameter *thost.CThostFtdcQryRCAMSIntraParameterField, nRequestID int,
) int {
	slog.Log(context.Background(), slog.LevelDebug-2,
		"executing thost trader api ReqQryRCAMSIntraParameter",
	)

	rtn := C.CallReqQryRCAMSIntraParameter(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQryRCAMSIntraParameter,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQryRCAMSIntraParameterField)(unsafe.Pointer(pQryRCAMSIntraParameter)),
		C.int(nRequestID),
	)

	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"thost trader api ReqQryRCAMSIntraParameter executed",
	)

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQryRCAMSInterParameter(
	pQryRCAMSInterParameter *thost.CThostFtdcQryRCAMSInterParameterField, nRequestID int,
) int {
	slog.Log(context.Background(), slog.LevelDebug-2,
		"executing thost trader api ReqQryRCAMSInterParameter",
	)

	rtn := C.CallReqQryRCAMSInterParameter(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQryRCAMSInterParameter,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQryRCAMSInterParameterField)(unsafe.Pointer(pQryRCAMSInterParameter)),
		C.int(nRequestID),
	)

	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"thost trader api ReqQryRCAMSInterParameter executed",
	)

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQryRCAMSShortOptAdjustParam(
	pQryRCAMSShortOptAdjustParam *thost.CThostFtdcQryRCAMSShortOptAdjustParamField, nRequestID int,
) int {
	slog.Log(context.Background(), slog.LevelDebug-2,
		"executing thost trader api ReqQryRCAMSShortOptAdjustParam",
	)

	rtn := C.CallReqQryRCAMSShortOptAdjustParam(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQryRCAMSShortOptAdjustParam,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQryRCAMSShortOptAdjustParamField)(unsafe.Pointer(pQryRCAMSShortOptAdjustParam)),
		C.int(nRequestID),
	)

	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"thost trader api ReqQryRCAMSShortOptAdjustParam executed",
	)

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQryRCAMSInvestorCombPosition(
	pQryRCAMSInvestorCombPosition *thost.CThostFtdcQryRCAMSInvestorCombPositionField, nRequestID int,
) int {
	slog.Log(context.Background(), slog.LevelDebug-2,
		"executing thost trader api ReqQryRCAMSInvestorCombPosition",
	)

	rtn := C.CallReqQryRCAMSInvestorCombPosition(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQryRCAMSInvestorCombPosition,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQryRCAMSInvestorCombPositionField)(unsafe.Pointer(pQryRCAMSInvestorCombPosition)),
		C.int(nRequestID),
	)

	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"thost trader api ReqQryRCAMSInvestorCombPosition executed",
	)

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQryInvestorProdRCAMSMargin(
	pQryInvestorProdRCAMSMargin *thost.CThostFtdcQryInvestorProdRCAMSMarginField, nRequestID int,
) int {
	slog.Log(context.Background(), slog.LevelDebug-2,
		"executing thost trader api ReqQryInvestorProdRCAMSMargin",
	)

	rtn := C.CallReqQryInvestorProdRCAMSMargin(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQryInvestorProdRCAMSMargin,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQryInvestorProdRCAMSMarginField)(unsafe.Pointer(pQryInvestorProdRCAMSMargin)),
		C.int(nRequestID),
	)

	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"thost trader api ReqQryInvestorProdRCAMSMargin executed",
	)

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQryRULEInstrParameter(
	pQryRULEInstrParameter *thost.CThostFtdcQryRULEInstrParameterField, nRequestID int,
) int {
	slog.Log(context.Background(), slog.LevelDebug-2,
		"executing thost trader api ReqQryRULEInstrParameter",
	)

	rtn := C.CallReqQryRULEInstrParameter(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQryRULEInstrParameter,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQryRULEInstrParameterField)(unsafe.Pointer(pQryRULEInstrParameter)),
		C.int(nRequestID),
	)

	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"thost trader api ReqQryRULEInstrParameter executed",
	)

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQryRULEIntraParameter(
	pQryRULEIntraParameter *thost.CThostFtdcQryRULEIntraParameterField, nRequestID int,
) int {
	slog.Log(context.Background(), slog.LevelDebug-2,
		"executing thost trader api ReqQryRULEIntraParameter",
	)

	rtn := C.CallReqQryRULEIntraParameter(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQryRULEIntraParameter,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQryRULEIntraParameterField)(unsafe.Pointer(pQryRULEIntraParameter)),
		C.int(nRequestID),
	)

	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"thost trader api ReqQryRULEIntraParameter executed",
	)

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQryRULEInterParameter(
	pQryRULEInterParameter *thost.CThostFtdcQryRULEInterParameterField, nRequestID int,
) int {
	slog.Log(context.Background(), slog.LevelDebug-2,
		"executing thost trader api ReqQryRULEInterParameter",
	)

	rtn := C.CallReqQryRULEInterParameter(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQryRULEInterParameter,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQryRULEInterParameterField)(unsafe.Pointer(pQryRULEInterParameter)),
		C.int(nRequestID),
	)

	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"thost trader api ReqQryRULEInterParameter executed",
	)

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQryInvestorProdRULEMargin(
	pQryInvestorProdRULEMargin *thost.CThostFtdcQryInvestorProdRULEMarginField, nRequestID int,
) int {
	slog.Log(context.Background(), slog.LevelDebug-2,
		"executing thost trader api ReqQryInvestorProdRULEMargin",
	)

	rtn := C.CallReqQryInvestorProdRULEMargin(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQryInvestorProdRULEMargin,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQryInvestorProdRULEMarginField)(unsafe.Pointer(pQryInvestorProdRULEMargin)),
		C.int(nRequestID),
	)

	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"thost trader api ReqQryInvestorProdRULEMargin executed",
	)

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQryInvestorPortfSetting(
	pQryInvestorPortfSetting *thost.CThostFtdcQryInvestorPortfSettingField, nRequestID int,
) int {
	slog.Log(context.Background(), slog.LevelDebug-2,
		"executing thost trader api ReqQryInvestorPortfSetting",
	)

	rtn := C.CallReqQryInvestorPortfSetting(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQryInvestorPortfSetting,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQryInvestorPortfSettingField)(unsafe.Pointer(pQryInvestorPortfSetting)),
		C.int(nRequestID),
	)

	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"thost trader api ReqQryInvestorPortfSetting executed",
	)

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQryInvestorInfoCommRec(
	pQryInvestorInfoCommRec *thost.CThostFtdcQryInvestorInfoCommRecField, nRequestID int,
) int {
	slog.Log(context.Background(), slog.LevelDebug-2,
		"executing thost trader api ReqQryInvestorInfoCommRec",
	)

	rtn := C.CallReqQryInvestorInfoCommRec(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQryInvestorInfoCommRec,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQryInvestorInfoCommRecField)(unsafe.Pointer(pQryInvestorInfoCommRec)),
		C.int(nRequestID),
	)

	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"thost trader api ReqQryInvestorInfoCommRec executed",
	)

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQryCombLeg(
	pQryCombLeg *thost.CThostFtdcQryCombLegField, nRequestID int,
) int {
	slog.Log(context.Background(), slog.LevelDebug-2,
		"executing thost trader api ReqQryCombLeg",
	)

	rtn := C.CallReqQryCombLeg(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQryCombLeg,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQryCombLegField)(unsafe.Pointer(pQryCombLeg)),
		C.int(nRequestID),
	)

	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"thost trader api ReqQryCombLeg executed",
	)

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqOffsetSetting(
	pInputOffsetSetting *thost.CThostFtdcInputOffsetSettingField, nRequestID int,
) int {
	slog.Log(context.Background(), slog.LevelDebug-2,
		"executing thost trader api ReqOffsetSetting",
	)

	rtn := C.CallReqOffsetSetting(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqOffsetSetting,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcInputOffsetSettingField)(unsafe.Pointer(pInputOffsetSetting)),
		C.int(nRequestID),
	)

	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"thost trader api ReqOffsetSetting executed",
	)

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqCancelOffsetSetting(
	pInputOffsetSetting *thost.CThostFtdcInputOffsetSettingField, nRequestID int,
) int {
	slog.Log(context.Background(), slog.LevelDebug-2,
		"executing thost trader api ReqCancelOffsetSetting",
	)

	rtn := C.CallReqCancelOffsetSetting(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqCancelOffsetSetting,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcInputOffsetSettingField)(unsafe.Pointer(pInputOffsetSetting)),
		C.int(nRequestID),
	)

	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"thost trader api ReqCancelOffsetSetting executed",
	)

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQryOffsetSetting(
	pQryOffsetSetting *thost.CThostFtdcQryOffsetSettingField, nRequestID int,
) int {
	slog.Log(context.Background(), slog.LevelDebug-2,
		"executing thost trader api ReqQryOffsetSetting",
	)

	rtn := C.CallReqQryOffsetSetting(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQryOffsetSetting,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQryOffsetSettingField)(unsafe.Pointer(pQryOffsetSetting)),
		C.int(nRequestID),
	)

	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"thost trader api ReqQryOffsetSetting executed",
	)

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqGenSMSCode(
	pReqGenSMSCode *thost.CThostFtdcReqGenSMSCodeField, nRequestID int,
) int {
	slog.Log(context.Background(), slog.LevelDebug-2,
		"executing thost trader api ReqGenSMSCode",
	)

	rtn := C.CallReqGenSMSCode(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqGenSMSCode,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcReqGenSMSCodeField)(unsafe.Pointer(pReqGenSMSCode)),
		C.int(nRequestID),
	)

	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"thost trader api ReqGenSMSCode executed",
	)

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqSpdApply(
	pInputSpdApply *thost.CThostFtdcInputSpdApplyField, nRequestID int,
) int {
	slog.Log(context.Background(), slog.LevelDebug-2,
		"executing thost trader api ReqSpdApply",
	)

	rtn := C.CallReqSpdApply(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqSpdApply,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcInputSpdApplyField)(unsafe.Pointer(pInputSpdApply)),
		C.int(nRequestID),
	)

	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"thost trader api ReqSpdApply executed",
	)

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqSpdApplyAction(
	pInputSpdApplyAction *thost.CThostFtdcInputSpdApplyActionField, nRequestID int,
) int {
	slog.Log(context.Background(), slog.LevelDebug-2,
		"executing thost trader api ReqSpdApplyAction",
	)

	rtn := C.CallReqSpdApplyAction(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqSpdApplyAction,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcInputSpdApplyActionField)(unsafe.Pointer(pInputSpdApplyAction)),
		C.int(nRequestID),
	)

	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"thost trader api ReqSpdApplyAction executed",
	)

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQrySpdApply(
	pQrySpdApply *thost.CThostFtdcQrySpdApplyField, nRequestID int,
) int {
	slog.Log(context.Background(), slog.LevelDebug-2,
		"executing thost trader api ReqQrySpdApply",
	)

	rtn := C.CallReqQrySpdApply(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQrySpdApply,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQrySpdApplyField)(unsafe.Pointer(pQrySpdApply)),
		C.int(nRequestID),
	)

	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"thost trader api ReqQrySpdApply executed",
	)

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqHedgeCfm(
	pInputHedgeCfm *thost.CThostFtdcInputHedgeCfmField, nRequestID int,
) int {
	slog.Log(context.Background(), slog.LevelDebug-2,
		"executing thost trader api ReqHedgeCfm",
	)

	rtn := C.CallReqHedgeCfm(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqHedgeCfm,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcInputHedgeCfmField)(unsafe.Pointer(pInputHedgeCfm)),
		C.int(nRequestID),
	)

	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"thost trader api ReqHedgeCfm executed",
	)

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqHedgeCfmAction(
	pInputHedgeCfmAction *thost.CThostFtdcInputHedgeCfmActionField, nRequestID int,
) int {
	slog.Log(context.Background(), slog.LevelDebug-2,
		"executing thost trader api ReqHedgeCfmAction",
	)

	rtn := C.CallReqHedgeCfmAction(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqHedgeCfmAction,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcInputHedgeCfmActionField)(unsafe.Pointer(pInputHedgeCfmAction)),
		C.int(nRequestID),
	)

	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"thost trader api ReqHedgeCfmAction executed",
	)

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQryHedgeCfm(
	pQryHedgeCfm *thost.CThostFtdcQryHedgeCfmField, nRequestID int,
) int {
	slog.Log(context.Background(), slog.LevelDebug-2,
		"executing thost trader api ReqQryHedgeCfm",
	)

	rtn := C.CallReqQryHedgeCfm(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQryHedgeCfm,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQryHedgeCfmField)(unsafe.Pointer(pQryHedgeCfm)),
		C.int(nRequestID),
	)

	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"thost trader api ReqQryHedgeCfm executed",
	)

	return int(rtn)
}
