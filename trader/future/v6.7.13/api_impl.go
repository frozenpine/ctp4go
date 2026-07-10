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
	_ future.TraderApi = &ThostFtdcTraderApi{}
)

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

func (api *ThostFtdcTraderApi) Init() {
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

func (api *ThostFtdcTraderApi) GetFrontInfo(FrontInfo *future.CThostFtdcFrontInfoField) {
	if api.apiPtr == nil {
		slog.Error(
			"thost trader api not initialized",
			slog.String("caller", "GetFrontInfo"),
		)
		return
	}

	slog.Info(
		"executing thost trader api",
		slog.String("caller", "GetFrontInfo"),
	)

	C.CallGetFrontInfo(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_GetFrontInfo,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcFrontInfoField)(unsafe.Pointer(FrontInfo)),
	)

	slog.Info(
		"thost trader api executed",
		slog.String("caller", "GetFrontInfo"),
	)

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

func (api *ThostFtdcTraderApi) RegisterNameServer(NsAddress string) {
	if api.apiPtr == nil {
		slog.Error(
			"thost trader api not initialized",
			slog.String("caller", "RegisterNameServer"),
		)
		return
	}

	slog.Info(
		"executing thost trader api",
		slog.String("caller", "RegisterNameServer"),
	)

	pszNsAddress := C.CString(NsAddress)
	defer C.free(unsafe.Pointer(pszNsAddress))

	C.CallRegisterNameServer(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_RegisterNameServer,
		unsafe.Pointer(api.apiPtr),
		pszNsAddress,
	)

	slog.Info(
		"thost trader api executed",
		slog.String("caller", "RegisterNameServer"),
	)

}

func (api *ThostFtdcTraderApi) RegisterFensUserInfo(FensUserInfo *future.CThostFtdcFensUserInfoField) {
	if api.apiPtr == nil {
		slog.Error(
			"thost trader api not initialized",
			slog.String("caller", "RegisterFensUserInfo"),
		)
		return
	}

	slog.Info(
		"executing thost trader api",
		slog.String("caller", "RegisterFensUserInfo"),
	)

	C.CallRegisterFensUserInfo(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_RegisterFensUserInfo,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcFensUserInfoField)(unsafe.Pointer(FensUserInfo)),
	)

	slog.Info(
		"thost trader api executed",
		slog.String("caller", "RegisterFensUserInfo"),
	)

}

func (api *ThostFtdcTraderApi) RegisterSpi(Spi future.TraderSpi) {
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

func (api *ThostFtdcTraderApi) SubscribePrivateTopic(ResumeType int, SeqNo int) {
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
		C.int(ResumeType), C.int(SeqNo),
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

func (api *ThostFtdcTraderApi) ReqAuthenticate(ReqAuthenticateField *future.CThostFtdcReqAuthenticateField, RequestID int) int {
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

func (api *ThostFtdcTraderApi) RegisterUserSystemInfo(UserSystemInfo *future.CThostFtdcUserSystemInfoField) int {
	if api.apiPtr == nil {
		slog.Error(
			"thost trader api not initialized",
			slog.String("caller", "RegisterUserSystemInfo"),
		)
		return -255
	}

	slog.Info(
		"executing thost trader api",
		slog.String("caller", "RegisterUserSystemInfo"),
	)

	rtn := C.CallRegisterUserSystemInfo(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_RegisterUserSystemInfo,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcUserSystemInfoField)(unsafe.Pointer(UserSystemInfo)),
	)

	slog.Info(
		"thost trader api executed",
		slog.String("caller", "RegisterUserSystemInfo"),
	)

	return int(rtn)
}

func (api *ThostFtdcTraderApi) SubmitUserSystemInfo(UserSystemInfo *future.CThostFtdcUserSystemInfoField) int {
	if api.apiPtr == nil {
		slog.Error(
			"thost trader api not initialized",
			slog.String("caller", "SubmitUserSystemInfo"),
		)
		return -255
	}

	slog.Info(
		"executing thost trader api",
		slog.String("caller", "SubmitUserSystemInfo"),
	)

	rtn := C.CallSubmitUserSystemInfo(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_SubmitUserSystemInfo,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcUserSystemInfoField)(unsafe.Pointer(UserSystemInfo)),
	)

	slog.Info(
		"thost trader api executed",
		slog.String("caller", "SubmitUserSystemInfo"),
	)

	return int(rtn)
}

func (api *ThostFtdcTraderApi) RegisterWechatUserSystemInfo(UserSystemInfo *future.CThostFtdcWechatUserSystemInfoField) int {
	if api.apiPtr == nil {
		slog.Error(
			"thost trader api not initialized",
			slog.String("caller", "RegisterWechatUserSystemInfo"),
		)
		return -255
	}

	slog.Info(
		"executing thost trader api",
		slog.String("caller", "RegisterWechatUserSystemInfo"),
	)

	rtn := C.CallRegisterWechatUserSystemInfo(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_RegisterWechatUserSystemInfo,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcWechatUserSystemInfoField)(unsafe.Pointer(UserSystemInfo)),
	)

	slog.Info(
		"thost trader api executed",
		slog.String("caller", "RegisterWechatUserSystemInfo"),
	)

	return int(rtn)
}

func (api *ThostFtdcTraderApi) SubmitWechatUserSystemInfo(UserSystemInfo *future.CThostFtdcWechatUserSystemInfoField) int {
	if api.apiPtr == nil {
		slog.Error(
			"thost trader api not initialized",
			slog.String("caller", "SubmitWechatUserSystemInfo"),
		)
		return -255
	}

	slog.Info(
		"executing thost trader api",
		slog.String("caller", "SubmitWechatUserSystemInfo"),
	)

	rtn := C.CallSubmitWechatUserSystemInfo(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_SubmitWechatUserSystemInfo,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcWechatUserSystemInfoField)(unsafe.Pointer(UserSystemInfo)),
	)

	slog.Info(
		"thost trader api executed",
		slog.String("caller", "SubmitWechatUserSystemInfo"),
	)

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqUserLogin(ReqUserLoginField *future.CThostFtdcReqUserLoginField, RequestID int) int {
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

func (api *ThostFtdcTraderApi) ReqUserLogout(UserLogout *future.CThostFtdcUserLogoutField, RequestID int) int {
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

func (api *ThostFtdcTraderApi) ReqUserPasswordUpdate(UserPasswordUpdate *future.CThostFtdcUserPasswordUpdateField, RequestID int) int {
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

func (api *ThostFtdcTraderApi) ReqTradingAccountPasswordUpdate(TradingAccountPasswordUpdate *future.CThostFtdcTradingAccountPasswordUpdateField, RequestID int) int {
	if api.apiPtr == nil {
		slog.Error(
			"thost trader api not initialized",
			slog.String("caller", "ReqTradingAccountPasswordUpdate"),
		)
		return -255
	}

	slog.Info(
		"executing thost trader api",
		slog.String("caller", "ReqTradingAccountPasswordUpdate"),
	)

	rtn := C.CallReqTradingAccountPasswordUpdate(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqTradingAccountPasswordUpdate,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcTradingAccountPasswordUpdateField)(unsafe.Pointer(TradingAccountPasswordUpdate)), C.int(RequestID),
	)

	slog.Info(
		"thost trader api executed",
		slog.String("caller", "ReqTradingAccountPasswordUpdate"),
	)

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqUserAuthMethod(ReqUserAuthMethod *future.CThostFtdcReqUserAuthMethodField, RequestID int) int {
	if api.apiPtr == nil {
		slog.Error(
			"thost trader api not initialized",
			slog.String("caller", "ReqUserAuthMethod"),
		)
		return -255
	}

	slog.Info(
		"executing thost trader api",
		slog.String("caller", "ReqUserAuthMethod"),
	)

	rtn := C.CallReqUserAuthMethod(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqUserAuthMethod,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcReqUserAuthMethodField)(unsafe.Pointer(ReqUserAuthMethod)), C.int(RequestID),
	)

	slog.Info(
		"thost trader api executed",
		slog.String("caller", "ReqUserAuthMethod"),
	)

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqGenUserCaptcha(ReqGenUserCaptcha *future.CThostFtdcReqGenUserCaptchaField, RequestID int) int {
	if api.apiPtr == nil {
		slog.Error(
			"thost trader api not initialized",
			slog.String("caller", "ReqGenUserCaptcha"),
		)
		return -255
	}

	slog.Info(
		"executing thost trader api",
		slog.String("caller", "ReqGenUserCaptcha"),
	)

	rtn := C.CallReqGenUserCaptcha(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqGenUserCaptcha,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcReqGenUserCaptchaField)(unsafe.Pointer(ReqGenUserCaptcha)), C.int(RequestID),
	)

	slog.Info(
		"thost trader api executed",
		slog.String("caller", "ReqGenUserCaptcha"),
	)

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqGenUserText(ReqGenUserText *future.CThostFtdcReqGenUserTextField, RequestID int) int {
	if api.apiPtr == nil {
		slog.Error(
			"thost trader api not initialized",
			slog.String("caller", "ReqGenUserText"),
		)
		return -255
	}

	slog.Info(
		"executing thost trader api",
		slog.String("caller", "ReqGenUserText"),
	)

	rtn := C.CallReqGenUserText(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqGenUserText,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcReqGenUserTextField)(unsafe.Pointer(ReqGenUserText)), C.int(RequestID),
	)

	slog.Info(
		"thost trader api executed",
		slog.String("caller", "ReqGenUserText"),
	)

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqUserLoginWithCaptcha(ReqUserLoginWithCaptcha *future.CThostFtdcReqUserLoginWithCaptchaField, RequestID int) int {
	if api.apiPtr == nil {
		slog.Error(
			"thost trader api not initialized",
			slog.String("caller", "ReqUserLoginWithCaptcha"),
		)
		return -255
	}

	slog.Info(
		"executing thost trader api",
		slog.String("caller", "ReqUserLoginWithCaptcha"),
	)

	rtn := C.CallReqUserLoginWithCaptcha(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqUserLoginWithCaptcha,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcReqUserLoginWithCaptchaField)(unsafe.Pointer(ReqUserLoginWithCaptcha)), C.int(RequestID),
	)

	slog.Info(
		"thost trader api executed",
		slog.String("caller", "ReqUserLoginWithCaptcha"),
	)

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqUserLoginWithText(ReqUserLoginWithText *future.CThostFtdcReqUserLoginWithTextField, RequestID int) int {
	if api.apiPtr == nil {
		slog.Error(
			"thost trader api not initialized",
			slog.String("caller", "ReqUserLoginWithText"),
		)
		return -255
	}

	slog.Info(
		"executing thost trader api",
		slog.String("caller", "ReqUserLoginWithText"),
	)

	rtn := C.CallReqUserLoginWithText(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqUserLoginWithText,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcReqUserLoginWithTextField)(unsafe.Pointer(ReqUserLoginWithText)), C.int(RequestID),
	)

	slog.Info(
		"thost trader api executed",
		slog.String("caller", "ReqUserLoginWithText"),
	)

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqUserLoginWithOTP(ReqUserLoginWithOTP *future.CThostFtdcReqUserLoginWithOTPField, RequestID int) int {
	if api.apiPtr == nil {
		slog.Error(
			"thost trader api not initialized",
			slog.String("caller", "ReqUserLoginWithOTP"),
		)
		return -255
	}

	slog.Info(
		"executing thost trader api",
		slog.String("caller", "ReqUserLoginWithOTP"),
	)

	rtn := C.CallReqUserLoginWithOTP(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqUserLoginWithOTP,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcReqUserLoginWithOTPField)(unsafe.Pointer(ReqUserLoginWithOTP)), C.int(RequestID),
	)

	slog.Info(
		"thost trader api executed",
		slog.String("caller", "ReqUserLoginWithOTP"),
	)

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqOrderInsert(InputOrder *future.CThostFtdcInputOrderField, RequestID int) int {
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

func (api *ThostFtdcTraderApi) ReqParkedOrderInsert(ParkedOrder *future.CThostFtdcParkedOrderField, RequestID int) int {
	if api.apiPtr == nil {
		slog.Error(
			"thost trader api not initialized",
			slog.String("caller", "ReqParkedOrderInsert"),
		)
		return -255
	}

	slog.Info(
		"executing thost trader api",
		slog.String("caller", "ReqParkedOrderInsert"),
	)

	rtn := C.CallReqParkedOrderInsert(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqParkedOrderInsert,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcParkedOrderField)(unsafe.Pointer(ParkedOrder)), C.int(RequestID),
	)

	slog.Info(
		"thost trader api executed",
		slog.String("caller", "ReqParkedOrderInsert"),
	)

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqParkedOrderAction(ParkedOrderAction *future.CThostFtdcParkedOrderActionField, RequestID int) int {
	if api.apiPtr == nil {
		slog.Error(
			"thost trader api not initialized",
			slog.String("caller", "ReqParkedOrderAction"),
		)
		return -255
	}

	slog.Info(
		"executing thost trader api",
		slog.String("caller", "ReqParkedOrderAction"),
	)

	rtn := C.CallReqParkedOrderAction(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqParkedOrderAction,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcParkedOrderActionField)(unsafe.Pointer(ParkedOrderAction)), C.int(RequestID),
	)

	slog.Info(
		"thost trader api executed",
		slog.String("caller", "ReqParkedOrderAction"),
	)

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqOrderAction(InputOrderAction *future.CThostFtdcInputOrderActionField, RequestID int) int {
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

func (api *ThostFtdcTraderApi) ReqQryMaxOrderVolume(QryMaxOrderVolume *future.CThostFtdcQryMaxOrderVolumeField, RequestID int) int {
	if api.apiPtr == nil {
		slog.Error(
			"thost trader api not initialized",
			slog.String("caller", "ReqQryMaxOrderVolume"),
		)
		return -255
	}

	slog.Info(
		"executing thost trader api",
		slog.String("caller", "ReqQryMaxOrderVolume"),
	)

	rtn := C.CallReqQryMaxOrderVolume(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQryMaxOrderVolume,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQryMaxOrderVolumeField)(unsafe.Pointer(QryMaxOrderVolume)), C.int(RequestID),
	)

	slog.Info(
		"thost trader api executed",
		slog.String("caller", "ReqQryMaxOrderVolume"),
	)

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqSettlementInfoConfirm(SettlementInfoConfirm *future.CThostFtdcSettlementInfoConfirmField, RequestID int) int {
	if api.apiPtr == nil {
		slog.Error(
			"thost trader api not initialized",
			slog.String("caller", "ReqSettlementInfoConfirm"),
		)
		return -255
	}

	slog.Info(
		"executing thost trader api",
		slog.String("caller", "ReqSettlementInfoConfirm"),
	)

	rtn := C.CallReqSettlementInfoConfirm(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqSettlementInfoConfirm,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcSettlementInfoConfirmField)(unsafe.Pointer(SettlementInfoConfirm)), C.int(RequestID),
	)

	slog.Info(
		"thost trader api executed",
		slog.String("caller", "ReqSettlementInfoConfirm"),
	)

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqRemoveParkedOrder(RemoveParkedOrder *future.CThostFtdcRemoveParkedOrderField, RequestID int) int {
	if api.apiPtr == nil {
		slog.Error(
			"thost trader api not initialized",
			slog.String("caller", "ReqRemoveParkedOrder"),
		)
		return -255
	}

	slog.Info(
		"executing thost trader api",
		slog.String("caller", "ReqRemoveParkedOrder"),
	)

	rtn := C.CallReqRemoveParkedOrder(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqRemoveParkedOrder,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcRemoveParkedOrderField)(unsafe.Pointer(RemoveParkedOrder)), C.int(RequestID),
	)

	slog.Info(
		"thost trader api executed",
		slog.String("caller", "ReqRemoveParkedOrder"),
	)

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqRemoveParkedOrderAction(RemoveParkedOrderAction *future.CThostFtdcRemoveParkedOrderActionField, RequestID int) int {
	if api.apiPtr == nil {
		slog.Error(
			"thost trader api not initialized",
			slog.String("caller", "ReqRemoveParkedOrderAction"),
		)
		return -255
	}

	slog.Info(
		"executing thost trader api",
		slog.String("caller", "ReqRemoveParkedOrderAction"),
	)

	rtn := C.CallReqRemoveParkedOrderAction(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqRemoveParkedOrderAction,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcRemoveParkedOrderActionField)(unsafe.Pointer(RemoveParkedOrderAction)), C.int(RequestID),
	)

	slog.Info(
		"thost trader api executed",
		slog.String("caller", "ReqRemoveParkedOrderAction"),
	)

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqExecOrderInsert(InputExecOrder *future.CThostFtdcInputExecOrderField, RequestID int) int {
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

func (api *ThostFtdcTraderApi) ReqExecOrderAction(InputExecOrderAction *future.CThostFtdcInputExecOrderActionField, RequestID int) int {
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

func (api *ThostFtdcTraderApi) ReqForQuoteInsert(InputForQuote *future.CThostFtdcInputForQuoteField, RequestID int) int {
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

func (api *ThostFtdcTraderApi) ReqQuoteInsert(InputQuote *future.CThostFtdcInputQuoteField, RequestID int) int {
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

func (api *ThostFtdcTraderApi) ReqQuoteAction(InputQuoteAction *future.CThostFtdcInputQuoteActionField, RequestID int) int {
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

func (api *ThostFtdcTraderApi) ReqBatchOrderAction(InputBatchOrderAction *future.CThostFtdcInputBatchOrderActionField, RequestID int) int {
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

func (api *ThostFtdcTraderApi) ReqOptionSelfCloseInsert(InputOptionSelfClose *future.CThostFtdcInputOptionSelfCloseField, RequestID int) int {
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

func (api *ThostFtdcTraderApi) ReqOptionSelfCloseAction(InputOptionSelfCloseAction *future.CThostFtdcInputOptionSelfCloseActionField, RequestID int) int {
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

func (api *ThostFtdcTraderApi) ReqCombActionInsert(InputCombAction *future.CThostFtdcInputCombActionField, RequestID int) int {
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

func (api *ThostFtdcTraderApi) ReqQryOrder(QryOrder *future.CThostFtdcQryOrderField, RequestID int) int {
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

func (api *ThostFtdcTraderApi) ReqQryTrade(QryTrade *future.CThostFtdcQryTradeField, RequestID int) int {
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

func (api *ThostFtdcTraderApi) ReqQryInvestorPosition(QryInvestorPosition *future.CThostFtdcQryInvestorPositionField, RequestID int) int {
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

func (api *ThostFtdcTraderApi) ReqQryTradingAccount(QryTradingAccount *future.CThostFtdcQryTradingAccountField, RequestID int) int {
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

func (api *ThostFtdcTraderApi) ReqQryInvestor(QryInvestor *future.CThostFtdcQryInvestorField, RequestID int) int {
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

func (api *ThostFtdcTraderApi) ReqQryTradingCode(QryTradingCode *future.CThostFtdcQryTradingCodeField, RequestID int) int {
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

func (api *ThostFtdcTraderApi) ReqQryInstrumentMarginRate(QryInstrumentMarginRate *future.CThostFtdcQryInstrumentMarginRateField, RequestID int) int {
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

func (api *ThostFtdcTraderApi) ReqQryInstrumentCommissionRate(QryInstrumentCommissionRate *future.CThostFtdcQryInstrumentCommissionRateField, RequestID int) int {
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

func (api *ThostFtdcTraderApi) ReqQryUserSession(QryUserSession *future.CThostFtdcQryUserSessionField, RequestID int) int {
	if api.apiPtr == nil {
		slog.Error(
			"thost trader api not initialized",
			slog.String("caller", "ReqQryUserSession"),
		)
		return -255
	}

	slog.Info(
		"executing thost trader api",
		slog.String("caller", "ReqQryUserSession"),
	)

	rtn := C.CallReqQryUserSession(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQryUserSession,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQryUserSessionField)(unsafe.Pointer(QryUserSession)), C.int(RequestID),
	)

	slog.Info(
		"thost trader api executed",
		slog.String("caller", "ReqQryUserSession"),
	)

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQryExchange(QryExchange *future.CThostFtdcQryExchangeField, RequestID int) int {
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

func (api *ThostFtdcTraderApi) ReqQryProduct(QryProduct *future.CThostFtdcQryProductField, RequestID int) int {
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

func (api *ThostFtdcTraderApi) ReqQryInstrument(QryInstrument *future.CThostFtdcQryInstrumentField, RequestID int) int {
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

func (api *ThostFtdcTraderApi) ReqQryDepthMarketData(QryDepthMarketData *future.CThostFtdcQryDepthMarketDataField, RequestID int) int {
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

func (api *ThostFtdcTraderApi) ReqQryTraderOffer(QryTraderOffer *future.CThostFtdcQryTraderOfferField, RequestID int) int {
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

func (api *ThostFtdcTraderApi) ReqQrySettlementInfo(QrySettlementInfo *future.CThostFtdcQrySettlementInfoField, RequestID int) int {
	if api.apiPtr == nil {
		slog.Error(
			"thost trader api not initialized",
			slog.String("caller", "ReqQrySettlementInfo"),
		)
		return -255
	}

	slog.Info(
		"executing thost trader api",
		slog.String("caller", "ReqQrySettlementInfo"),
	)

	rtn := C.CallReqQrySettlementInfo(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQrySettlementInfo,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQrySettlementInfoField)(unsafe.Pointer(QrySettlementInfo)), C.int(RequestID),
	)

	slog.Info(
		"thost trader api executed",
		slog.String("caller", "ReqQrySettlementInfo"),
	)

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQryTransferBank(QryTransferBank *future.CThostFtdcQryTransferBankField, RequestID int) int {
	if api.apiPtr == nil {
		slog.Error(
			"thost trader api not initialized",
			slog.String("caller", "ReqQryTransferBank"),
		)
		return -255
	}

	slog.Info(
		"executing thost trader api",
		slog.String("caller", "ReqQryTransferBank"),
	)

	rtn := C.CallReqQryTransferBank(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQryTransferBank,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQryTransferBankField)(unsafe.Pointer(QryTransferBank)), C.int(RequestID),
	)

	slog.Info(
		"thost trader api executed",
		slog.String("caller", "ReqQryTransferBank"),
	)

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQryInvestorPositionDetail(QryInvestorPositionDetail *future.CThostFtdcQryInvestorPositionDetailField, RequestID int) int {
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

func (api *ThostFtdcTraderApi) ReqQryNotice(QryNotice *future.CThostFtdcQryNoticeField, RequestID int) int {
	if api.apiPtr == nil {
		slog.Error(
			"thost trader api not initialized",
			slog.String("caller", "ReqQryNotice"),
		)
		return -255
	}

	slog.Info(
		"executing thost trader api",
		slog.String("caller", "ReqQryNotice"),
	)

	rtn := C.CallReqQryNotice(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQryNotice,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQryNoticeField)(unsafe.Pointer(QryNotice)), C.int(RequestID),
	)

	slog.Info(
		"thost trader api executed",
		slog.String("caller", "ReqQryNotice"),
	)

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQrySettlementInfoConfirm(QrySettlementInfoConfirm *future.CThostFtdcQrySettlementInfoConfirmField, RequestID int) int {
	if api.apiPtr == nil {
		slog.Error(
			"thost trader api not initialized",
			slog.String("caller", "ReqQrySettlementInfoConfirm"),
		)
		return -255
	}

	slog.Info(
		"executing thost trader api",
		slog.String("caller", "ReqQrySettlementInfoConfirm"),
	)

	rtn := C.CallReqQrySettlementInfoConfirm(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQrySettlementInfoConfirm,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQrySettlementInfoConfirmField)(unsafe.Pointer(QrySettlementInfoConfirm)), C.int(RequestID),
	)

	slog.Info(
		"thost trader api executed",
		slog.String("caller", "ReqQrySettlementInfoConfirm"),
	)

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQryInvestorPositionCombineDetail(QryInvestorPositionCombineDetail *future.CThostFtdcQryInvestorPositionCombineDetailField, RequestID int) int {
	if api.apiPtr == nil {
		slog.Error(
			"thost trader api not initialized",
			slog.String("caller", "ReqQryInvestorPositionCombineDetail"),
		)
		return -255
	}

	slog.Info(
		"executing thost trader api",
		slog.String("caller", "ReqQryInvestorPositionCombineDetail"),
	)

	rtn := C.CallReqQryInvestorPositionCombineDetail(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQryInvestorPositionCombineDetail,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQryInvestorPositionCombineDetailField)(unsafe.Pointer(QryInvestorPositionCombineDetail)), C.int(RequestID),
	)

	slog.Info(
		"thost trader api executed",
		slog.String("caller", "ReqQryInvestorPositionCombineDetail"),
	)

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQryCFMMCTradingAccountKey(QryCFMMCTradingAccountKey *future.CThostFtdcQryCFMMCTradingAccountKeyField, RequestID int) int {
	if api.apiPtr == nil {
		slog.Error(
			"thost trader api not initialized",
			slog.String("caller", "ReqQryCFMMCTradingAccountKey"),
		)
		return -255
	}

	slog.Info(
		"executing thost trader api",
		slog.String("caller", "ReqQryCFMMCTradingAccountKey"),
	)

	rtn := C.CallReqQryCFMMCTradingAccountKey(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQryCFMMCTradingAccountKey,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQryCFMMCTradingAccountKeyField)(unsafe.Pointer(QryCFMMCTradingAccountKey)), C.int(RequestID),
	)

	slog.Info(
		"thost trader api executed",
		slog.String("caller", "ReqQryCFMMCTradingAccountKey"),
	)

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQryEWarrantOffset(QryEWarrantOffset *future.CThostFtdcQryEWarrantOffsetField, RequestID int) int {
	if api.apiPtr == nil {
		slog.Error(
			"thost trader api not initialized",
			slog.String("caller", "ReqQryEWarrantOffset"),
		)
		return -255
	}

	slog.Info(
		"executing thost trader api",
		slog.String("caller", "ReqQryEWarrantOffset"),
	)

	rtn := C.CallReqQryEWarrantOffset(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQryEWarrantOffset,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQryEWarrantOffsetField)(unsafe.Pointer(QryEWarrantOffset)), C.int(RequestID),
	)

	slog.Info(
		"thost trader api executed",
		slog.String("caller", "ReqQryEWarrantOffset"),
	)

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQryInvestorProductGroupMargin(QryInvestorProductGroupMargin *future.CThostFtdcQryInvestorProductGroupMarginField, RequestID int) int {
	if api.apiPtr == nil {
		slog.Error(
			"thost trader api not initialized",
			slog.String("caller", "ReqQryInvestorProductGroupMargin"),
		)
		return -255
	}

	slog.Info(
		"executing thost trader api",
		slog.String("caller", "ReqQryInvestorProductGroupMargin"),
	)

	rtn := C.CallReqQryInvestorProductGroupMargin(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQryInvestorProductGroupMargin,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQryInvestorProductGroupMarginField)(unsafe.Pointer(QryInvestorProductGroupMargin)), C.int(RequestID),
	)

	slog.Info(
		"thost trader api executed",
		slog.String("caller", "ReqQryInvestorProductGroupMargin"),
	)

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQryExchangeMarginRate(QryExchangeMarginRate *future.CThostFtdcQryExchangeMarginRateField, RequestID int) int {
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

func (api *ThostFtdcTraderApi) ReqQryExchangeMarginRateAdjust(QryExchangeMarginRateAdjust *future.CThostFtdcQryExchangeMarginRateAdjustField, RequestID int) int {
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

func (api *ThostFtdcTraderApi) ReqQryExchangeRate(QryExchangeRate *future.CThostFtdcQryExchangeRateField, RequestID int) int {
	if api.apiPtr == nil {
		slog.Error(
			"thost trader api not initialized",
			slog.String("caller", "ReqQryExchangeRate"),
		)
		return -255
	}

	slog.Info(
		"executing thost trader api",
		slog.String("caller", "ReqQryExchangeRate"),
	)

	rtn := C.CallReqQryExchangeRate(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQryExchangeRate,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQryExchangeRateField)(unsafe.Pointer(QryExchangeRate)), C.int(RequestID),
	)

	slog.Info(
		"thost trader api executed",
		slog.String("caller", "ReqQryExchangeRate"),
	)

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQrySecAgentACIDMap(QrySecAgentACIDMap *future.CThostFtdcQrySecAgentACIDMapField, RequestID int) int {
	if api.apiPtr == nil {
		slog.Error(
			"thost trader api not initialized",
			slog.String("caller", "ReqQrySecAgentACIDMap"),
		)
		return -255
	}

	slog.Info(
		"executing thost trader api",
		slog.String("caller", "ReqQrySecAgentACIDMap"),
	)

	rtn := C.CallReqQrySecAgentACIDMap(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQrySecAgentACIDMap,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQrySecAgentACIDMapField)(unsafe.Pointer(QrySecAgentACIDMap)), C.int(RequestID),
	)

	slog.Info(
		"thost trader api executed",
		slog.String("caller", "ReqQrySecAgentACIDMap"),
	)

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQryProductExchRate(QryProductExchRate *future.CThostFtdcQryProductExchRateField, RequestID int) int {
	if api.apiPtr == nil {
		slog.Error(
			"thost trader api not initialized",
			slog.String("caller", "ReqQryProductExchRate"),
		)
		return -255
	}

	slog.Info(
		"executing thost trader api",
		slog.String("caller", "ReqQryProductExchRate"),
	)

	rtn := C.CallReqQryProductExchRate(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQryProductExchRate,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQryProductExchRateField)(unsafe.Pointer(QryProductExchRate)), C.int(RequestID),
	)

	slog.Info(
		"thost trader api executed",
		slog.String("caller", "ReqQryProductExchRate"),
	)

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQryProductGroup(QryProductGroup *future.CThostFtdcQryProductGroupField, RequestID int) int {
	if api.apiPtr == nil {
		slog.Error(
			"thost trader api not initialized",
			slog.String("caller", "ReqQryProductGroup"),
		)
		return -255
	}

	slog.Info(
		"executing thost trader api",
		slog.String("caller", "ReqQryProductGroup"),
	)

	rtn := C.CallReqQryProductGroup(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQryProductGroup,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQryProductGroupField)(unsafe.Pointer(QryProductGroup)), C.int(RequestID),
	)

	slog.Info(
		"thost trader api executed",
		slog.String("caller", "ReqQryProductGroup"),
	)

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQryMMInstrumentCommissionRate(QryMMInstrumentCommissionRate *future.CThostFtdcQryMMInstrumentCommissionRateField, RequestID int) int {
	if api.apiPtr == nil {
		slog.Error(
			"thost trader api not initialized",
			slog.String("caller", "ReqQryMMInstrumentCommissionRate"),
		)
		return -255
	}

	slog.Info(
		"executing thost trader api",
		slog.String("caller", "ReqQryMMInstrumentCommissionRate"),
	)

	rtn := C.CallReqQryMMInstrumentCommissionRate(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQryMMInstrumentCommissionRate,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQryMMInstrumentCommissionRateField)(unsafe.Pointer(QryMMInstrumentCommissionRate)), C.int(RequestID),
	)

	slog.Info(
		"thost trader api executed",
		slog.String("caller", "ReqQryMMInstrumentCommissionRate"),
	)

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQryMMOptionInstrCommRate(QryMMOptionInstrCommRate *future.CThostFtdcQryMMOptionInstrCommRateField, RequestID int) int {
	if api.apiPtr == nil {
		slog.Error(
			"thost trader api not initialized",
			slog.String("caller", "ReqQryMMOptionInstrCommRate"),
		)
		return -255
	}

	slog.Info(
		"executing thost trader api",
		slog.String("caller", "ReqQryMMOptionInstrCommRate"),
	)

	rtn := C.CallReqQryMMOptionInstrCommRate(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQryMMOptionInstrCommRate,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQryMMOptionInstrCommRateField)(unsafe.Pointer(QryMMOptionInstrCommRate)), C.int(RequestID),
	)

	slog.Info(
		"thost trader api executed",
		slog.String("caller", "ReqQryMMOptionInstrCommRate"),
	)

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQryInstrumentOrderCommRate(QryInstrumentOrderCommRate *future.CThostFtdcQryInstrumentOrderCommRateField, RequestID int) int {
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

func (api *ThostFtdcTraderApi) ReqQrySecAgentTradingAccount(QryTradingAccount *future.CThostFtdcQryTradingAccountField, RequestID int) int {
	if api.apiPtr == nil {
		slog.Error(
			"thost trader api not initialized",
			slog.String("caller", "ReqQrySecAgentTradingAccount"),
		)
		return -255
	}

	slog.Info(
		"executing thost trader api",
		slog.String("caller", "ReqQrySecAgentTradingAccount"),
	)

	rtn := C.CallReqQrySecAgentTradingAccount(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQrySecAgentTradingAccount,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQryTradingAccountField)(unsafe.Pointer(QryTradingAccount)), C.int(RequestID),
	)

	slog.Info(
		"thost trader api executed",
		slog.String("caller", "ReqQrySecAgentTradingAccount"),
	)

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQrySecAgentCheckMode(QrySecAgentCheckMode *future.CThostFtdcQrySecAgentCheckModeField, RequestID int) int {
	if api.apiPtr == nil {
		slog.Error(
			"thost trader api not initialized",
			slog.String("caller", "ReqQrySecAgentCheckMode"),
		)
		return -255
	}

	slog.Info(
		"executing thost trader api",
		slog.String("caller", "ReqQrySecAgentCheckMode"),
	)

	rtn := C.CallReqQrySecAgentCheckMode(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQrySecAgentCheckMode,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQrySecAgentCheckModeField)(unsafe.Pointer(QrySecAgentCheckMode)), C.int(RequestID),
	)

	slog.Info(
		"thost trader api executed",
		slog.String("caller", "ReqQrySecAgentCheckMode"),
	)

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQrySecAgentTradeInfo(QrySecAgentTradeInfo *future.CThostFtdcQrySecAgentTradeInfoField, RequestID int) int {
	if api.apiPtr == nil {
		slog.Error(
			"thost trader api not initialized",
			slog.String("caller", "ReqQrySecAgentTradeInfo"),
		)
		return -255
	}

	slog.Info(
		"executing thost trader api",
		slog.String("caller", "ReqQrySecAgentTradeInfo"),
	)

	rtn := C.CallReqQrySecAgentTradeInfo(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQrySecAgentTradeInfo,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQrySecAgentTradeInfoField)(unsafe.Pointer(QrySecAgentTradeInfo)), C.int(RequestID),
	)

	slog.Info(
		"thost trader api executed",
		slog.String("caller", "ReqQrySecAgentTradeInfo"),
	)

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQryOptionInstrTradeCost(QryOptionInstrTradeCost *future.CThostFtdcQryOptionInstrTradeCostField, RequestID int) int {
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

func (api *ThostFtdcTraderApi) ReqQryOptionInstrCommRate(QryOptionInstrCommRate *future.CThostFtdcQryOptionInstrCommRateField, RequestID int) int {
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

func (api *ThostFtdcTraderApi) ReqQryExecOrder(QryExecOrder *future.CThostFtdcQryExecOrderField, RequestID int) int {
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

func (api *ThostFtdcTraderApi) ReqQryForQuote(QryForQuote *future.CThostFtdcQryForQuoteField, RequestID int) int {
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

func (api *ThostFtdcTraderApi) ReqQryQuote(QryQuote *future.CThostFtdcQryQuoteField, RequestID int) int {
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

func (api *ThostFtdcTraderApi) ReqQryOptionSelfClose(QryOptionSelfClose *future.CThostFtdcQryOptionSelfCloseField, RequestID int) int {
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

func (api *ThostFtdcTraderApi) ReqQryInvestUnit(QryInvestUnit *future.CThostFtdcQryInvestUnitField, RequestID int) int {
	if api.apiPtr == nil {
		slog.Error(
			"thost trader api not initialized",
			slog.String("caller", "ReqQryInvestUnit"),
		)
		return -255
	}

	slog.Info(
		"executing thost trader api",
		slog.String("caller", "ReqQryInvestUnit"),
	)

	rtn := C.CallReqQryInvestUnit(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQryInvestUnit,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQryInvestUnitField)(unsafe.Pointer(QryInvestUnit)), C.int(RequestID),
	)

	slog.Info(
		"thost trader api executed",
		slog.String("caller", "ReqQryInvestUnit"),
	)

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQryCombInstrumentGuard(QryCombInstrumentGuard *future.CThostFtdcQryCombInstrumentGuardField, RequestID int) int {
	if api.apiPtr == nil {
		slog.Error(
			"thost trader api not initialized",
			slog.String("caller", "ReqQryCombInstrumentGuard"),
		)
		return -255
	}

	slog.Info(
		"executing thost trader api",
		slog.String("caller", "ReqQryCombInstrumentGuard"),
	)

	rtn := C.CallReqQryCombInstrumentGuard(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQryCombInstrumentGuard,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQryCombInstrumentGuardField)(unsafe.Pointer(QryCombInstrumentGuard)), C.int(RequestID),
	)

	slog.Info(
		"thost trader api executed",
		slog.String("caller", "ReqQryCombInstrumentGuard"),
	)

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQryCombAction(QryCombAction *future.CThostFtdcQryCombActionField, RequestID int) int {
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

func (api *ThostFtdcTraderApi) ReqQryTransferSerial(QryTransferSerial *future.CThostFtdcQryTransferSerialField, RequestID int) int {
	if api.apiPtr == nil {
		slog.Error(
			"thost trader api not initialized",
			slog.String("caller", "ReqQryTransferSerial"),
		)
		return -255
	}

	slog.Info(
		"executing thost trader api",
		slog.String("caller", "ReqQryTransferSerial"),
	)

	rtn := C.CallReqQryTransferSerial(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQryTransferSerial,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQryTransferSerialField)(unsafe.Pointer(QryTransferSerial)), C.int(RequestID),
	)

	slog.Info(
		"thost trader api executed",
		slog.String("caller", "ReqQryTransferSerial"),
	)

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQryAccountregister(QryAccountregister *future.CThostFtdcQryAccountregisterField, RequestID int) int {
	if api.apiPtr == nil {
		slog.Error(
			"thost trader api not initialized",
			slog.String("caller", "ReqQryAccountregister"),
		)
		return -255
	}

	slog.Info(
		"executing thost trader api",
		slog.String("caller", "ReqQryAccountregister"),
	)

	rtn := C.CallReqQryAccountregister(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQryAccountregister,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQryAccountregisterField)(unsafe.Pointer(QryAccountregister)), C.int(RequestID),
	)

	slog.Info(
		"thost trader api executed",
		slog.String("caller", "ReqQryAccountregister"),
	)

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQryContractBank(QryContractBank *future.CThostFtdcQryContractBankField, RequestID int) int {
	if api.apiPtr == nil {
		slog.Error(
			"thost trader api not initialized",
			slog.String("caller", "ReqQryContractBank"),
		)
		return -255
	}

	slog.Info(
		"executing thost trader api",
		slog.String("caller", "ReqQryContractBank"),
	)

	rtn := C.CallReqQryContractBank(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQryContractBank,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQryContractBankField)(unsafe.Pointer(QryContractBank)), C.int(RequestID),
	)

	slog.Info(
		"thost trader api executed",
		slog.String("caller", "ReqQryContractBank"),
	)

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQryParkedOrder(QryParkedOrder *future.CThostFtdcQryParkedOrderField, RequestID int) int {
	if api.apiPtr == nil {
		slog.Error(
			"thost trader api not initialized",
			slog.String("caller", "ReqQryParkedOrder"),
		)
		return -255
	}

	slog.Info(
		"executing thost trader api",
		slog.String("caller", "ReqQryParkedOrder"),
	)

	rtn := C.CallReqQryParkedOrder(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQryParkedOrder,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQryParkedOrderField)(unsafe.Pointer(QryParkedOrder)), C.int(RequestID),
	)

	slog.Info(
		"thost trader api executed",
		slog.String("caller", "ReqQryParkedOrder"),
	)

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQryParkedOrderAction(QryParkedOrderAction *future.CThostFtdcQryParkedOrderActionField, RequestID int) int {
	if api.apiPtr == nil {
		slog.Error(
			"thost trader api not initialized",
			slog.String("caller", "ReqQryParkedOrderAction"),
		)
		return -255
	}

	slog.Info(
		"executing thost trader api",
		slog.String("caller", "ReqQryParkedOrderAction"),
	)

	rtn := C.CallReqQryParkedOrderAction(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQryParkedOrderAction,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQryParkedOrderActionField)(unsafe.Pointer(QryParkedOrderAction)), C.int(RequestID),
	)

	slog.Info(
		"thost trader api executed",
		slog.String("caller", "ReqQryParkedOrderAction"),
	)

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQryTradingNotice(QryTradingNotice *future.CThostFtdcQryTradingNoticeField, RequestID int) int {
	if api.apiPtr == nil {
		slog.Error(
			"thost trader api not initialized",
			slog.String("caller", "ReqQryTradingNotice"),
		)
		return -255
	}

	slog.Info(
		"executing thost trader api",
		slog.String("caller", "ReqQryTradingNotice"),
	)

	rtn := C.CallReqQryTradingNotice(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQryTradingNotice,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQryTradingNoticeField)(unsafe.Pointer(QryTradingNotice)), C.int(RequestID),
	)

	slog.Info(
		"thost trader api executed",
		slog.String("caller", "ReqQryTradingNotice"),
	)

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQryBrokerTradingParams(QryBrokerTradingParams *future.CThostFtdcQryBrokerTradingParamsField, RequestID int) int {
	if api.apiPtr == nil {
		slog.Error(
			"thost trader api not initialized",
			slog.String("caller", "ReqQryBrokerTradingParams"),
		)
		return -255
	}

	slog.Info(
		"executing thost trader api",
		slog.String("caller", "ReqQryBrokerTradingParams"),
	)

	rtn := C.CallReqQryBrokerTradingParams(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQryBrokerTradingParams,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQryBrokerTradingParamsField)(unsafe.Pointer(QryBrokerTradingParams)), C.int(RequestID),
	)

	slog.Info(
		"thost trader api executed",
		slog.String("caller", "ReqQryBrokerTradingParams"),
	)

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQryBrokerTradingAlgos(QryBrokerTradingAlgos *future.CThostFtdcQryBrokerTradingAlgosField, RequestID int) int {
	if api.apiPtr == nil {
		slog.Error(
			"thost trader api not initialized",
			slog.String("caller", "ReqQryBrokerTradingAlgos"),
		)
		return -255
	}

	slog.Info(
		"executing thost trader api",
		slog.String("caller", "ReqQryBrokerTradingAlgos"),
	)

	rtn := C.CallReqQryBrokerTradingAlgos(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQryBrokerTradingAlgos,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQryBrokerTradingAlgosField)(unsafe.Pointer(QryBrokerTradingAlgos)), C.int(RequestID),
	)

	slog.Info(
		"thost trader api executed",
		slog.String("caller", "ReqQryBrokerTradingAlgos"),
	)

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQueryCFMMCTradingAccountToken(QueryCFMMCTradingAccountToken *future.CThostFtdcQueryCFMMCTradingAccountTokenField, RequestID int) int {
	if api.apiPtr == nil {
		slog.Error(
			"thost trader api not initialized",
			slog.String("caller", "ReqQueryCFMMCTradingAccountToken"),
		)
		return -255
	}

	slog.Info(
		"executing thost trader api",
		slog.String("caller", "ReqQueryCFMMCTradingAccountToken"),
	)

	rtn := C.CallReqQueryCFMMCTradingAccountToken(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQueryCFMMCTradingAccountToken,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQueryCFMMCTradingAccountTokenField)(unsafe.Pointer(QueryCFMMCTradingAccountToken)), C.int(RequestID),
	)

	slog.Info(
		"thost trader api executed",
		slog.String("caller", "ReqQueryCFMMCTradingAccountToken"),
	)

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqFromBankToFutureByFuture(ReqTransfer *future.CThostFtdcReqTransferField, RequestID int) int {
	if api.apiPtr == nil {
		slog.Error(
			"thost trader api not initialized",
			slog.String("caller", "ReqFromBankToFutureByFuture"),
		)
		return -255
	}

	slog.Info(
		"executing thost trader api",
		slog.String("caller", "ReqFromBankToFutureByFuture"),
	)

	rtn := C.CallReqFromBankToFutureByFuture(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqFromBankToFutureByFuture,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcReqTransferField)(unsafe.Pointer(ReqTransfer)), C.int(RequestID),
	)

	slog.Info(
		"thost trader api executed",
		slog.String("caller", "ReqFromBankToFutureByFuture"),
	)

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqFromFutureToBankByFuture(ReqTransfer *future.CThostFtdcReqTransferField, RequestID int) int {
	if api.apiPtr == nil {
		slog.Error(
			"thost trader api not initialized",
			slog.String("caller", "ReqFromFutureToBankByFuture"),
		)
		return -255
	}

	slog.Info(
		"executing thost trader api",
		slog.String("caller", "ReqFromFutureToBankByFuture"),
	)

	rtn := C.CallReqFromFutureToBankByFuture(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqFromFutureToBankByFuture,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcReqTransferField)(unsafe.Pointer(ReqTransfer)), C.int(RequestID),
	)

	slog.Info(
		"thost trader api executed",
		slog.String("caller", "ReqFromFutureToBankByFuture"),
	)

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQueryBankAccountMoneyByFuture(ReqQueryAccount *future.CThostFtdcReqQueryAccountField, RequestID int) int {
	if api.apiPtr == nil {
		slog.Error(
			"thost trader api not initialized",
			slog.String("caller", "ReqQueryBankAccountMoneyByFuture"),
		)
		return -255
	}

	slog.Info(
		"executing thost trader api",
		slog.String("caller", "ReqQueryBankAccountMoneyByFuture"),
	)

	rtn := C.CallReqQueryBankAccountMoneyByFuture(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQueryBankAccountMoneyByFuture,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcReqQueryAccountField)(unsafe.Pointer(ReqQueryAccount)), C.int(RequestID),
	)

	slog.Info(
		"thost trader api executed",
		slog.String("caller", "ReqQueryBankAccountMoneyByFuture"),
	)

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQryClassifiedInstrument(QryClassifiedInstrument *future.CThostFtdcQryClassifiedInstrumentField, RequestID int) int {
	if api.apiPtr == nil {
		slog.Error(
			"thost trader api not initialized",
			slog.String("caller", "ReqQryClassifiedInstrument"),
		)
		return -255
	}

	slog.Info(
		"executing thost trader api",
		slog.String("caller", "ReqQryClassifiedInstrument"),
	)

	rtn := C.CallReqQryClassifiedInstrument(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQryClassifiedInstrument,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQryClassifiedInstrumentField)(unsafe.Pointer(QryClassifiedInstrument)), C.int(RequestID),
	)

	slog.Info(
		"thost trader api executed",
		slog.String("caller", "ReqQryClassifiedInstrument"),
	)

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQryCombPromotionParam(QryCombPromotionParam *future.CThostFtdcQryCombPromotionParamField, RequestID int) int {
	if api.apiPtr == nil {
		slog.Error(
			"thost trader api not initialized",
			slog.String("caller", "ReqQryCombPromotionParam"),
		)
		return -255
	}

	slog.Info(
		"executing thost trader api",
		slog.String("caller", "ReqQryCombPromotionParam"),
	)

	rtn := C.CallReqQryCombPromotionParam(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQryCombPromotionParam,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQryCombPromotionParamField)(unsafe.Pointer(QryCombPromotionParam)), C.int(RequestID),
	)

	slog.Info(
		"thost trader api executed",
		slog.String("caller", "ReqQryCombPromotionParam"),
	)

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQryRiskSettleInvstPosition(QryRiskSettleInvstPosition *future.CThostFtdcQryRiskSettleInvstPositionField, RequestID int) int {
	if api.apiPtr == nil {
		slog.Error(
			"thost trader api not initialized",
			slog.String("caller", "ReqQryRiskSettleInvstPosition"),
		)
		return -255
	}

	slog.Info(
		"executing thost trader api",
		slog.String("caller", "ReqQryRiskSettleInvstPosition"),
	)

	rtn := C.CallReqQryRiskSettleInvstPosition(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQryRiskSettleInvstPosition,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQryRiskSettleInvstPositionField)(unsafe.Pointer(QryRiskSettleInvstPosition)), C.int(RequestID),
	)

	slog.Info(
		"thost trader api executed",
		slog.String("caller", "ReqQryRiskSettleInvstPosition"),
	)

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQryRiskSettleProductStatus(QryRiskSettleProductStatus *future.CThostFtdcQryRiskSettleProductStatusField, RequestID int) int {
	if api.apiPtr == nil {
		slog.Error(
			"thost trader api not initialized",
			slog.String("caller", "ReqQryRiskSettleProductStatus"),
		)
		return -255
	}

	slog.Info(
		"executing thost trader api",
		slog.String("caller", "ReqQryRiskSettleProductStatus"),
	)

	rtn := C.CallReqQryRiskSettleProductStatus(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQryRiskSettleProductStatus,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQryRiskSettleProductStatusField)(unsafe.Pointer(QryRiskSettleProductStatus)), C.int(RequestID),
	)

	slog.Info(
		"thost trader api executed",
		slog.String("caller", "ReqQryRiskSettleProductStatus"),
	)

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQrySPBMFutureParameter(QrySPBMFutureParameter *future.CThostFtdcQrySPBMFutureParameterField, RequestID int) int {
	if api.apiPtr == nil {
		slog.Error(
			"thost trader api not initialized",
			slog.String("caller", "ReqQrySPBMFutureParameter"),
		)
		return -255
	}

	slog.Info(
		"executing thost trader api",
		slog.String("caller", "ReqQrySPBMFutureParameter"),
	)

	rtn := C.CallReqQrySPBMFutureParameter(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQrySPBMFutureParameter,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQrySPBMFutureParameterField)(unsafe.Pointer(QrySPBMFutureParameter)), C.int(RequestID),
	)

	slog.Info(
		"thost trader api executed",
		slog.String("caller", "ReqQrySPBMFutureParameter"),
	)

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQrySPBMOptionParameter(QrySPBMOptionParameter *future.CThostFtdcQrySPBMOptionParameterField, RequestID int) int {
	if api.apiPtr == nil {
		slog.Error(
			"thost trader api not initialized",
			slog.String("caller", "ReqQrySPBMOptionParameter"),
		)
		return -255
	}

	slog.Info(
		"executing thost trader api",
		slog.String("caller", "ReqQrySPBMOptionParameter"),
	)

	rtn := C.CallReqQrySPBMOptionParameter(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQrySPBMOptionParameter,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQrySPBMOptionParameterField)(unsafe.Pointer(QrySPBMOptionParameter)), C.int(RequestID),
	)

	slog.Info(
		"thost trader api executed",
		slog.String("caller", "ReqQrySPBMOptionParameter"),
	)

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQrySPBMIntraParameter(QrySPBMIntraParameter *future.CThostFtdcQrySPBMIntraParameterField, RequestID int) int {
	if api.apiPtr == nil {
		slog.Error(
			"thost trader api not initialized",
			slog.String("caller", "ReqQrySPBMIntraParameter"),
		)
		return -255
	}

	slog.Info(
		"executing thost trader api",
		slog.String("caller", "ReqQrySPBMIntraParameter"),
	)

	rtn := C.CallReqQrySPBMIntraParameter(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQrySPBMIntraParameter,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQrySPBMIntraParameterField)(unsafe.Pointer(QrySPBMIntraParameter)), C.int(RequestID),
	)

	slog.Info(
		"thost trader api executed",
		slog.String("caller", "ReqQrySPBMIntraParameter"),
	)

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQrySPBMInterParameter(QrySPBMInterParameter *future.CThostFtdcQrySPBMInterParameterField, RequestID int) int {
	if api.apiPtr == nil {
		slog.Error(
			"thost trader api not initialized",
			slog.String("caller", "ReqQrySPBMInterParameter"),
		)
		return -255
	}

	slog.Info(
		"executing thost trader api",
		slog.String("caller", "ReqQrySPBMInterParameter"),
	)

	rtn := C.CallReqQrySPBMInterParameter(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQrySPBMInterParameter,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQrySPBMInterParameterField)(unsafe.Pointer(QrySPBMInterParameter)), C.int(RequestID),
	)

	slog.Info(
		"thost trader api executed",
		slog.String("caller", "ReqQrySPBMInterParameter"),
	)

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQrySPBMPortfDefinition(QrySPBMPortfDefinition *future.CThostFtdcQrySPBMPortfDefinitionField, RequestID int) int {
	if api.apiPtr == nil {
		slog.Error(
			"thost trader api not initialized",
			slog.String("caller", "ReqQrySPBMPortfDefinition"),
		)
		return -255
	}

	slog.Info(
		"executing thost trader api",
		slog.String("caller", "ReqQrySPBMPortfDefinition"),
	)

	rtn := C.CallReqQrySPBMPortfDefinition(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQrySPBMPortfDefinition,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQrySPBMPortfDefinitionField)(unsafe.Pointer(QrySPBMPortfDefinition)), C.int(RequestID),
	)

	slog.Info(
		"thost trader api executed",
		slog.String("caller", "ReqQrySPBMPortfDefinition"),
	)

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQrySPBMInvestorPortfDef(QrySPBMInvestorPortfDef *future.CThostFtdcQrySPBMInvestorPortfDefField, RequestID int) int {
	if api.apiPtr == nil {
		slog.Error(
			"thost trader api not initialized",
			slog.String("caller", "ReqQrySPBMInvestorPortfDef"),
		)
		return -255
	}

	slog.Info(
		"executing thost trader api",
		slog.String("caller", "ReqQrySPBMInvestorPortfDef"),
	)

	rtn := C.CallReqQrySPBMInvestorPortfDef(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQrySPBMInvestorPortfDef,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQrySPBMInvestorPortfDefField)(unsafe.Pointer(QrySPBMInvestorPortfDef)), C.int(RequestID),
	)

	slog.Info(
		"thost trader api executed",
		slog.String("caller", "ReqQrySPBMInvestorPortfDef"),
	)

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQryInvestorPortfMarginRatio(QryInvestorPortfMarginRatio *future.CThostFtdcQryInvestorPortfMarginRatioField, RequestID int) int {
	if api.apiPtr == nil {
		slog.Error(
			"thost trader api not initialized",
			slog.String("caller", "ReqQryInvestorPortfMarginRatio"),
		)
		return -255
	}

	slog.Info(
		"executing thost trader api",
		slog.String("caller", "ReqQryInvestorPortfMarginRatio"),
	)

	rtn := C.CallReqQryInvestorPortfMarginRatio(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQryInvestorPortfMarginRatio,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQryInvestorPortfMarginRatioField)(unsafe.Pointer(QryInvestorPortfMarginRatio)), C.int(RequestID),
	)

	slog.Info(
		"thost trader api executed",
		slog.String("caller", "ReqQryInvestorPortfMarginRatio"),
	)

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQryInvestorProdSPBMDetail(QryInvestorProdSPBMDetail *future.CThostFtdcQryInvestorProdSPBMDetailField, RequestID int) int {
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

func (api *ThostFtdcTraderApi) ReqQryInvestorCommoditySPMMMargin(QryInvestorCommoditySPMMMargin *future.CThostFtdcQryInvestorCommoditySPMMMarginField, RequestID int) int {
	if api.apiPtr == nil {
		slog.Error(
			"thost trader api not initialized",
			slog.String("caller", "ReqQryInvestorCommoditySPMMMargin"),
		)
		return -255
	}

	slog.Info(
		"executing thost trader api",
		slog.String("caller", "ReqQryInvestorCommoditySPMMMargin"),
	)

	rtn := C.CallReqQryInvestorCommoditySPMMMargin(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQryInvestorCommoditySPMMMargin,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQryInvestorCommoditySPMMMarginField)(unsafe.Pointer(QryInvestorCommoditySPMMMargin)), C.int(RequestID),
	)

	slog.Info(
		"thost trader api executed",
		slog.String("caller", "ReqQryInvestorCommoditySPMMMargin"),
	)

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQryInvestorCommodityGroupSPMMMargin(QryInvestorCommodityGroupSPMMMargin *future.CThostFtdcQryInvestorCommodityGroupSPMMMarginField, RequestID int) int {
	if api.apiPtr == nil {
		slog.Error(
			"thost trader api not initialized",
			slog.String("caller", "ReqQryInvestorCommodityGroupSPMMMargin"),
		)
		return -255
	}

	slog.Info(
		"executing thost trader api",
		slog.String("caller", "ReqQryInvestorCommodityGroupSPMMMargin"),
	)

	rtn := C.CallReqQryInvestorCommodityGroupSPMMMargin(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQryInvestorCommodityGroupSPMMMargin,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQryInvestorCommodityGroupSPMMMarginField)(unsafe.Pointer(QryInvestorCommodityGroupSPMMMargin)), C.int(RequestID),
	)

	slog.Info(
		"thost trader api executed",
		slog.String("caller", "ReqQryInvestorCommodityGroupSPMMMargin"),
	)

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQrySPMMInstParam(QrySPMMInstParam *future.CThostFtdcQrySPMMInstParamField, RequestID int) int {
	if api.apiPtr == nil {
		slog.Error(
			"thost trader api not initialized",
			slog.String("caller", "ReqQrySPMMInstParam"),
		)
		return -255
	}

	slog.Info(
		"executing thost trader api",
		slog.String("caller", "ReqQrySPMMInstParam"),
	)

	rtn := C.CallReqQrySPMMInstParam(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQrySPMMInstParam,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQrySPMMInstParamField)(unsafe.Pointer(QrySPMMInstParam)), C.int(RequestID),
	)

	slog.Info(
		"thost trader api executed",
		slog.String("caller", "ReqQrySPMMInstParam"),
	)

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQrySPMMProductParam(QrySPMMProductParam *future.CThostFtdcQrySPMMProductParamField, RequestID int) int {
	if api.apiPtr == nil {
		slog.Error(
			"thost trader api not initialized",
			slog.String("caller", "ReqQrySPMMProductParam"),
		)
		return -255
	}

	slog.Info(
		"executing thost trader api",
		slog.String("caller", "ReqQrySPMMProductParam"),
	)

	rtn := C.CallReqQrySPMMProductParam(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQrySPMMProductParam,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQrySPMMProductParamField)(unsafe.Pointer(QrySPMMProductParam)), C.int(RequestID),
	)

	slog.Info(
		"thost trader api executed",
		slog.String("caller", "ReqQrySPMMProductParam"),
	)

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQrySPBMAddOnInterParameter(QrySPBMAddOnInterParameter *future.CThostFtdcQrySPBMAddOnInterParameterField, RequestID int) int {
	if api.apiPtr == nil {
		slog.Error(
			"thost trader api not initialized",
			slog.String("caller", "ReqQrySPBMAddOnInterParameter"),
		)
		return -255
	}

	slog.Info(
		"executing thost trader api",
		slog.String("caller", "ReqQrySPBMAddOnInterParameter"),
	)

	rtn := C.CallReqQrySPBMAddOnInterParameter(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQrySPBMAddOnInterParameter,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQrySPBMAddOnInterParameterField)(unsafe.Pointer(QrySPBMAddOnInterParameter)), C.int(RequestID),
	)

	slog.Info(
		"thost trader api executed",
		slog.String("caller", "ReqQrySPBMAddOnInterParameter"),
	)

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQryRCAMSCombProductInfo(QryRCAMSCombProductInfo *future.CThostFtdcQryRCAMSCombProductInfoField, RequestID int) int {
	if api.apiPtr == nil {
		slog.Error(
			"thost trader api not initialized",
			slog.String("caller", "ReqQryRCAMSCombProductInfo"),
		)
		return -255
	}

	slog.Info(
		"executing thost trader api",
		slog.String("caller", "ReqQryRCAMSCombProductInfo"),
	)

	rtn := C.CallReqQryRCAMSCombProductInfo(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQryRCAMSCombProductInfo,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQryRCAMSCombProductInfoField)(unsafe.Pointer(QryRCAMSCombProductInfo)), C.int(RequestID),
	)

	slog.Info(
		"thost trader api executed",
		slog.String("caller", "ReqQryRCAMSCombProductInfo"),
	)

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQryRCAMSInstrParameter(QryRCAMSInstrParameter *future.CThostFtdcQryRCAMSInstrParameterField, RequestID int) int {
	if api.apiPtr == nil {
		slog.Error(
			"thost trader api not initialized",
			slog.String("caller", "ReqQryRCAMSInstrParameter"),
		)
		return -255
	}

	slog.Info(
		"executing thost trader api",
		slog.String("caller", "ReqQryRCAMSInstrParameter"),
	)

	rtn := C.CallReqQryRCAMSInstrParameter(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQryRCAMSInstrParameter,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQryRCAMSInstrParameterField)(unsafe.Pointer(QryRCAMSInstrParameter)), C.int(RequestID),
	)

	slog.Info(
		"thost trader api executed",
		slog.String("caller", "ReqQryRCAMSInstrParameter"),
	)

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQryRCAMSIntraParameter(QryRCAMSIntraParameter *future.CThostFtdcQryRCAMSIntraParameterField, RequestID int) int {
	if api.apiPtr == nil {
		slog.Error(
			"thost trader api not initialized",
			slog.String("caller", "ReqQryRCAMSIntraParameter"),
		)
		return -255
	}

	slog.Info(
		"executing thost trader api",
		slog.String("caller", "ReqQryRCAMSIntraParameter"),
	)

	rtn := C.CallReqQryRCAMSIntraParameter(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQryRCAMSIntraParameter,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQryRCAMSIntraParameterField)(unsafe.Pointer(QryRCAMSIntraParameter)), C.int(RequestID),
	)

	slog.Info(
		"thost trader api executed",
		slog.String("caller", "ReqQryRCAMSIntraParameter"),
	)

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQryRCAMSInterParameter(QryRCAMSInterParameter *future.CThostFtdcQryRCAMSInterParameterField, RequestID int) int {
	if api.apiPtr == nil {
		slog.Error(
			"thost trader api not initialized",
			slog.String("caller", "ReqQryRCAMSInterParameter"),
		)
		return -255
	}

	slog.Info(
		"executing thost trader api",
		slog.String("caller", "ReqQryRCAMSInterParameter"),
	)

	rtn := C.CallReqQryRCAMSInterParameter(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQryRCAMSInterParameter,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQryRCAMSInterParameterField)(unsafe.Pointer(QryRCAMSInterParameter)), C.int(RequestID),
	)

	slog.Info(
		"thost trader api executed",
		slog.String("caller", "ReqQryRCAMSInterParameter"),
	)

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQryRCAMSShortOptAdjustParam(QryRCAMSShortOptAdjustParam *future.CThostFtdcQryRCAMSShortOptAdjustParamField, RequestID int) int {
	if api.apiPtr == nil {
		slog.Error(
			"thost trader api not initialized",
			slog.String("caller", "ReqQryRCAMSShortOptAdjustParam"),
		)
		return -255
	}

	slog.Info(
		"executing thost trader api",
		slog.String("caller", "ReqQryRCAMSShortOptAdjustParam"),
	)

	rtn := C.CallReqQryRCAMSShortOptAdjustParam(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQryRCAMSShortOptAdjustParam,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQryRCAMSShortOptAdjustParamField)(unsafe.Pointer(QryRCAMSShortOptAdjustParam)), C.int(RequestID),
	)

	slog.Info(
		"thost trader api executed",
		slog.String("caller", "ReqQryRCAMSShortOptAdjustParam"),
	)

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQryRCAMSInvestorCombPosition(QryRCAMSInvestorCombPosition *future.CThostFtdcQryRCAMSInvestorCombPositionField, RequestID int) int {
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

func (api *ThostFtdcTraderApi) ReqQryInvestorProdRCAMSMargin(QryInvestorProdRCAMSMargin *future.CThostFtdcQryInvestorProdRCAMSMarginField, RequestID int) int {
	if api.apiPtr == nil {
		slog.Error(
			"thost trader api not initialized",
			slog.String("caller", "ReqQryInvestorProdRCAMSMargin"),
		)
		return -255
	}

	slog.Info(
		"executing thost trader api",
		slog.String("caller", "ReqQryInvestorProdRCAMSMargin"),
	)

	rtn := C.CallReqQryInvestorProdRCAMSMargin(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQryInvestorProdRCAMSMargin,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQryInvestorProdRCAMSMarginField)(unsafe.Pointer(QryInvestorProdRCAMSMargin)), C.int(RequestID),
	)

	slog.Info(
		"thost trader api executed",
		slog.String("caller", "ReqQryInvestorProdRCAMSMargin"),
	)

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQryRULEInstrParameter(QryRULEInstrParameter *future.CThostFtdcQryRULEInstrParameterField, RequestID int) int {
	if api.apiPtr == nil {
		slog.Error(
			"thost trader api not initialized",
			slog.String("caller", "ReqQryRULEInstrParameter"),
		)
		return -255
	}

	slog.Info(
		"executing thost trader api",
		slog.String("caller", "ReqQryRULEInstrParameter"),
	)

	rtn := C.CallReqQryRULEInstrParameter(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQryRULEInstrParameter,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQryRULEInstrParameterField)(unsafe.Pointer(QryRULEInstrParameter)), C.int(RequestID),
	)

	slog.Info(
		"thost trader api executed",
		slog.String("caller", "ReqQryRULEInstrParameter"),
	)

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQryRULEIntraParameter(QryRULEIntraParameter *future.CThostFtdcQryRULEIntraParameterField, RequestID int) int {
	if api.apiPtr == nil {
		slog.Error(
			"thost trader api not initialized",
			slog.String("caller", "ReqQryRULEIntraParameter"),
		)
		return -255
	}

	slog.Info(
		"executing thost trader api",
		slog.String("caller", "ReqQryRULEIntraParameter"),
	)

	rtn := C.CallReqQryRULEIntraParameter(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQryRULEIntraParameter,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQryRULEIntraParameterField)(unsafe.Pointer(QryRULEIntraParameter)), C.int(RequestID),
	)

	slog.Info(
		"thost trader api executed",
		slog.String("caller", "ReqQryRULEIntraParameter"),
	)

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQryRULEInterParameter(QryRULEInterParameter *future.CThostFtdcQryRULEInterParameterField, RequestID int) int {
	if api.apiPtr == nil {
		slog.Error(
			"thost trader api not initialized",
			slog.String("caller", "ReqQryRULEInterParameter"),
		)
		return -255
	}

	slog.Info(
		"executing thost trader api",
		slog.String("caller", "ReqQryRULEInterParameter"),
	)

	rtn := C.CallReqQryRULEInterParameter(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQryRULEInterParameter,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQryRULEInterParameterField)(unsafe.Pointer(QryRULEInterParameter)), C.int(RequestID),
	)

	slog.Info(
		"thost trader api executed",
		slog.String("caller", "ReqQryRULEInterParameter"),
	)

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQryInvestorProdRULEMargin(QryInvestorProdRULEMargin *future.CThostFtdcQryInvestorProdRULEMarginField, RequestID int) int {
	if api.apiPtr == nil {
		slog.Error(
			"thost trader api not initialized",
			slog.String("caller", "ReqQryInvestorProdRULEMargin"),
		)
		return -255
	}

	slog.Info(
		"executing thost trader api",
		slog.String("caller", "ReqQryInvestorProdRULEMargin"),
	)

	rtn := C.CallReqQryInvestorProdRULEMargin(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQryInvestorProdRULEMargin,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQryInvestorProdRULEMarginField)(unsafe.Pointer(QryInvestorProdRULEMargin)), C.int(RequestID),
	)

	slog.Info(
		"thost trader api executed",
		slog.String("caller", "ReqQryInvestorProdRULEMargin"),
	)

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQryInvestorPortfSetting(QryInvestorPortfSetting *future.CThostFtdcQryInvestorPortfSettingField, RequestID int) int {
	if api.apiPtr == nil {
		slog.Error(
			"thost trader api not initialized",
			slog.String("caller", "ReqQryInvestorPortfSetting"),
		)
		return -255
	}

	slog.Info(
		"executing thost trader api",
		slog.String("caller", "ReqQryInvestorPortfSetting"),
	)

	rtn := C.CallReqQryInvestorPortfSetting(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQryInvestorPortfSetting,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQryInvestorPortfSettingField)(unsafe.Pointer(QryInvestorPortfSetting)), C.int(RequestID),
	)

	slog.Info(
		"thost trader api executed",
		slog.String("caller", "ReqQryInvestorPortfSetting"),
	)

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQryInvestorInfoCommRec(QryInvestorInfoCommRec *future.CThostFtdcQryInvestorInfoCommRecField, RequestID int) int {
	if api.apiPtr == nil {
		slog.Error(
			"thost trader api not initialized",
			slog.String("caller", "ReqQryInvestorInfoCommRec"),
		)
		return -255
	}

	slog.Info(
		"executing thost trader api",
		slog.String("caller", "ReqQryInvestorInfoCommRec"),
	)

	rtn := C.CallReqQryInvestorInfoCommRec(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQryInvestorInfoCommRec,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQryInvestorInfoCommRecField)(unsafe.Pointer(QryInvestorInfoCommRec)), C.int(RequestID),
	)

	slog.Info(
		"thost trader api executed",
		slog.String("caller", "ReqQryInvestorInfoCommRec"),
	)

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQryCombLeg(QryCombLeg *future.CThostFtdcQryCombLegField, RequestID int) int {
	if api.apiPtr == nil {
		slog.Error(
			"thost trader api not initialized",
			slog.String("caller", "ReqQryCombLeg"),
		)
		return -255
	}

	slog.Info(
		"executing thost trader api",
		slog.String("caller", "ReqQryCombLeg"),
	)

	rtn := C.CallReqQryCombLeg(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQryCombLeg,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQryCombLegField)(unsafe.Pointer(QryCombLeg)), C.int(RequestID),
	)

	slog.Info(
		"thost trader api executed",
		slog.String("caller", "ReqQryCombLeg"),
	)

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqOffsetSetting(InputOffsetSetting *future.CThostFtdcInputOffsetSettingField, RequestID int) int {
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

func (api *ThostFtdcTraderApi) ReqCancelOffsetSetting(InputOffsetSetting *future.CThostFtdcInputOffsetSettingField, RequestID int) int {
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

func (api *ThostFtdcTraderApi) ReqQryOffsetSetting(QryOffsetSetting *future.CThostFtdcQryOffsetSettingField, RequestID int) int {
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

func (api *ThostFtdcTraderApi) ReqGenSMSCode(ReqGenSMSCode *future.CThostFtdcReqGenSMSCodeField, RequestID int) int {
	if api.apiPtr == nil {
		slog.Error(
			"thost trader api not initialized",
			slog.String("caller", "ReqGenSMSCode"),
		)
		return -255
	}

	slog.Info(
		"executing thost trader api",
		slog.String("caller", "ReqGenSMSCode"),
	)

	rtn := C.CallReqGenSMSCode(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqGenSMSCode,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcReqGenSMSCodeField)(unsafe.Pointer(ReqGenSMSCode)), C.int(RequestID),
	)

	slog.Info(
		"thost trader api executed",
		slog.String("caller", "ReqGenSMSCode"),
	)

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqSpdApply(InputSpdApply *future.CThostFtdcInputSpdApplyField, RequestID int) int {
	if api.apiPtr == nil {
		slog.Error(
			"thost trader api not initialized",
			slog.String("caller", "ReqSpdApply"),
		)
		return -255
	}

	slog.Info(
		"executing thost trader api",
		slog.String("caller", "ReqSpdApply"),
	)

	rtn := C.CallReqSpdApply(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqSpdApply,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcInputSpdApplyField)(unsafe.Pointer(InputSpdApply)), C.int(RequestID),
	)

	slog.Info(
		"thost trader api executed",
		slog.String("caller", "ReqSpdApply"),
	)

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqSpdApplyAction(InputSpdApplyAction *future.CThostFtdcInputSpdApplyActionField, RequestID int) int {
	if api.apiPtr == nil {
		slog.Error(
			"thost trader api not initialized",
			slog.String("caller", "ReqSpdApplyAction"),
		)
		return -255
	}

	slog.Info(
		"executing thost trader api",
		slog.String("caller", "ReqSpdApplyAction"),
	)

	rtn := C.CallReqSpdApplyAction(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqSpdApplyAction,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcInputSpdApplyActionField)(unsafe.Pointer(InputSpdApplyAction)), C.int(RequestID),
	)

	slog.Info(
		"thost trader api executed",
		slog.String("caller", "ReqSpdApplyAction"),
	)

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQrySpdApply(QrySpdApply *future.CThostFtdcQrySpdApplyField, RequestID int) int {
	if api.apiPtr == nil {
		slog.Error(
			"thost trader api not initialized",
			slog.String("caller", "ReqQrySpdApply"),
		)
		return -255
	}

	slog.Info(
		"executing thost trader api",
		slog.String("caller", "ReqQrySpdApply"),
	)

	rtn := C.CallReqQrySpdApply(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQrySpdApply,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQrySpdApplyField)(unsafe.Pointer(QrySpdApply)), C.int(RequestID),
	)

	slog.Info(
		"thost trader api executed",
		slog.String("caller", "ReqQrySpdApply"),
	)

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqHedgeCfm(InputHedgeCfm *future.CThostFtdcInputHedgeCfmField, RequestID int) int {
	if api.apiPtr == nil {
		slog.Error(
			"thost trader api not initialized",
			slog.String("caller", "ReqHedgeCfm"),
		)
		return -255
	}

	slog.Info(
		"executing thost trader api",
		slog.String("caller", "ReqHedgeCfm"),
	)

	rtn := C.CallReqHedgeCfm(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqHedgeCfm,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcInputHedgeCfmField)(unsafe.Pointer(InputHedgeCfm)), C.int(RequestID),
	)

	slog.Info(
		"thost trader api executed",
		slog.String("caller", "ReqHedgeCfm"),
	)

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqHedgeCfmAction(InputHedgeCfmAction *future.CThostFtdcInputHedgeCfmActionField, RequestID int) int {
	if api.apiPtr == nil {
		slog.Error(
			"thost trader api not initialized",
			slog.String("caller", "ReqHedgeCfmAction"),
		)
		return -255
	}

	slog.Info(
		"executing thost trader api",
		slog.String("caller", "ReqHedgeCfmAction"),
	)

	rtn := C.CallReqHedgeCfmAction(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqHedgeCfmAction,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcInputHedgeCfmActionField)(unsafe.Pointer(InputHedgeCfmAction)), C.int(RequestID),
	)

	slog.Info(
		"thost trader api executed",
		slog.String("caller", "ReqHedgeCfmAction"),
	)

	return int(rtn)
}

func (api *ThostFtdcTraderApi) ReqQryHedgeCfm(QryHedgeCfm *future.CThostFtdcQryHedgeCfmField, RequestID int) int {
	if api.apiPtr == nil {
		slog.Error(
			"thost trader api not initialized",
			slog.String("caller", "ReqQryHedgeCfm"),
		)
		return -255
	}

	slog.Info(
		"executing thost trader api",
		slog.String("caller", "ReqQryHedgeCfm"),
	)

	rtn := C.CallReqQryHedgeCfm(
		api.apiPtr.vtable.CThostFtdcTraderApiVTable_ReqQryHedgeCfm,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CThostFtdcQryHedgeCfmField)(unsafe.Pointer(QryHedgeCfm)), C.int(RequestID),
	)

	slog.Info(
		"thost trader api executed",
		slog.String("caller", "ReqQryHedgeCfm"),
	)

	return int(rtn)
}
