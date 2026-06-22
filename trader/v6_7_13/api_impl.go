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

// / 删除接口对象本身
// /@remark 不再使用本接口对象时,调用该函数删除接口对象
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

// / 初始化
// /@remark 初始化运行环境,只有调用后,接口才开始工作
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

// / 等待接口线程结束运行
// /@return 线程退出代码
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

// / 获取当前交易日
// /@retrun 获取到的交易日
// /@remark 只有登录成功后,才能得到正确的交易日
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

// / 获取已连接的前置的信息
// /  @param pFrontInfo：输入输出参数，用于存储获取到的前置信息，不能为空
// /  @remark 连接成功后，可获取正确的前置地址信息
// /  @remark 登录成功后，可获取正确的前置流控信息
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

// / 注册前置机网络地址
// /@param pszFrontAddress：前置机网络地址。
// /@remark 网络地址的格式为：“protocol://ipaddress:port”，如：”tcp://127.0.0.1:17001”。
// /@remark “tcp”代表传输协议，“127.0.0.1”代表服务器地址。”17001”代表服务器端口号。
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

// / 注册名字服务器网络地址
// /@param pszNsAddress：名字服务器网络地址。
// /@remark 网络地址的格式为：“protocol://ipaddress:port”，如：”tcp://127.0.0.1:12001”。
// /@remark “tcp”代表传输协议，“127.0.0.1”代表服务器地址。”12001”代表服务器端口号。
// /@remark RegisterNameServer优先于RegisterFront
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

// / 注册名字服务器用户信息
// /@param pFensUserInfo：用户信息。
func (api *ThostFtdcTraderApi) RegisterFensUserInfo(pFensUserInfo *thost.CThostFtdcFensUserInfoField) {
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

// / 注册回调接口
// /@param pSpi 派生自回调接口类的实例
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

// / 订阅私有流。
// /@param nResumeType 私有流重传方式
// /         THOST_TERT_RESTART:从本交易日开始重传
// /         THOST_TERT_RESUME:从上次收到的续传
// /         THOST_TERT_QUICK:只传送登录后私有流的内容
// /         THOST_TERT_RESUME_FROM_SEQ_NO:从指定序号开始重传，序号从1开始
// /@param nSeqNo 私有流序号，只在THOST_TERT_RESUME_FROM_SEQ_NO模式下有效
// /@remark 该方法要在Init方法前调用。若不调用则不会收到私有流的数据。
func (api *ThostFtdcTraderApi) SubscribePrivateTopic(nResumeType thost.THOST_TE_RESUME_TYPE, nSeqNo int) {
}

// / 订阅公共流。
// /@param nResumeType 公共流重传方式
// /         THOST_TERT_RESTART:从本交易日开始重传
// /         THOST_TERT_RESUME:从上次收到的续传
// /         THOST_TERT_QUICK:只传送登录后公共流的内容
// /         THOST_TERT_NONE:取消订阅公共流
// /@remark 该方法要在Init方法前调用。若不调用则不会收到公共流的数据。
func (api *ThostFtdcTraderApi) SubscribePublicTopic(nResumeType thost.THOST_TE_RESUME_TYPE) {}

// / 客户端认证请求
func (api *ThostFtdcTraderApi) ReqAuthenticate(pReqAuthenticateField *thost.CThostFtdcReqAuthenticateField, nRequestID int) int {
}

// / 注册用户终端信息，用于中继服务器多连接模式
// / 需要在终端认证成功后，用户登录前调用该接口
func (api *ThostFtdcTraderApi) RegisterUserSystemInfo(pUserSystemInfo *thost.CThostFtdcUserSystemInfoField) int {
}

// / 上报用户终端信息，用于中继服务器操作员登录模式
// / 操作员登录后，可以多次调用该接口上报客户信息
func (api *ThostFtdcTraderApi) SubmitUserSystemInfo(pUserSystemInfo *thost.CThostFtdcUserSystemInfoField) int {
}

// / 注册用户终端信息，用于中继服务器多连接模式.用于微信小程序等应用上报信息.
func (api *ThostFtdcTraderApi) RegisterWechatUserSystemInfo(pUserSystemInfo *thost.CThostFtdcWechatUserSystemInfoField) int {
}

// / 上报用户终端信息，用于中继服务器操作员登录模式.用于微信小程序等应用上报信息.
func (api *ThostFtdcTraderApi) SubmitWechatUserSystemInfo(pUserSystemInfo *thost.CThostFtdcWechatUserSystemInfoField) int {
}

// / 用户登录请求
func (api *ThostFtdcTraderApi) ReqUserLogin(pReqUserLoginField *thost.CThostFtdcReqUserLoginField, nRequestID int) int {
}

// / 登出请求
func (api *ThostFtdcTraderApi) ReqUserLogout(pUserLogout *thost.CThostFtdcUserLogoutField, nRequestID int) int {
}

// / 用户口令更新请求
func (api *ThostFtdcTraderApi) ReqUserPasswordUpdate(pUserPasswordUpdate *thost.CThostFtdcUserPasswordUpdateField, nRequestID int) int {
}

// / 资金账户口令更新请求
func (api *ThostFtdcTraderApi) ReqTradingAccountPasswordUpdate(pTradingAccountPasswordUpdate *thost.CThostFtdcTradingAccountPasswordUpdateField, nRequestID int) int {
}

// / 查询用户当前支持的认证模式
func (api *ThostFtdcTraderApi) ReqUserAuthMethod(pReqUserAuthMethod *thost.CThostFtdcReqUserAuthMethodField, nRequestID int) int {
}

// / 用户发出获取图形验证码请求
func (api *ThostFtdcTraderApi) ReqGenUserCaptcha(pReqGenUserCaptcha *thost.CThostFtdcReqGenUserCaptchaField, nRequestID int) int {
}

// / 用户发出获取短信验证码请求
func (api *ThostFtdcTraderApi) ReqGenUserText(pReqGenUserText *thost.CThostFtdcReqGenUserTextField, nRequestID int) int {
}

// / 用户发出带有图片验证码的登陆请求
func (api *ThostFtdcTraderApi) ReqUserLoginWithCaptcha(pReqUserLoginWithCaptcha *thost.CThostFtdcReqUserLoginWithCaptchaField, nRequestID int) int {
}

// / 用户发出带有短信验证码的登陆请求
func (api *ThostFtdcTraderApi) ReqUserLoginWithText(pReqUserLoginWithText *thost.CThostFtdcReqUserLoginWithTextField, nRequestID int) int {
}

// / 用户发出带有动态口令的登陆请求
func (api *ThostFtdcTraderApi) ReqUserLoginWithOTP(pReqUserLoginWithOTP *thost.CThostFtdcReqUserLoginWithOTPField, nRequestID int) int {
}

// / 报单录入请求
func (api *ThostFtdcTraderApi) ReqOrderInsert(pInputOrder *thost.CThostFtdcInputOrderField, nRequestID int) int {
}

// / 预埋单录入请求
func (api *ThostFtdcTraderApi) ReqParkedOrderInsert(pParkedOrder *thost.CThostFtdcParkedOrderField, nRequestID int) int {
}

// / 预埋撤单录入请求
func (api *ThostFtdcTraderApi) ReqParkedOrderAction(pParkedOrderAction *thost.CThostFtdcParkedOrderActionField, nRequestID int) int {
}

// / 报单操作请求
func (api *ThostFtdcTraderApi) ReqOrderAction(pInputOrderAction *thost.CThostFtdcInputOrderActionField, nRequestID int) int {
}

// / 查询最大报单数量请求
func (api *ThostFtdcTraderApi) ReqQryMaxOrderVolume(pQryMaxOrderVolume *thost.CThostFtdcQryMaxOrderVolumeField, nRequestID int) int {
}

// / 投资者结算结果确认
func (api *ThostFtdcTraderApi) ReqSettlementInfoConfirm(pSettlementInfoConfirm *thost.CThostFtdcSettlementInfoConfirmField, nRequestID int) int {
}

// / 请求删除预埋单
func (api *ThostFtdcTraderApi) ReqRemoveParkedOrder(pRemoveParkedOrder *thost.CThostFtdcRemoveParkedOrderField, nRequestID int) int {
}

// / 请求删除预埋撤单
func (api *ThostFtdcTraderApi) ReqRemoveParkedOrderAction(pRemoveParkedOrderAction *thost.CThostFtdcRemoveParkedOrderActionField, nRequestID int) int {
}

// / 执行宣告录入请求
func (api *ThostFtdcTraderApi) ReqExecOrderInsert(pInputExecOrder *thost.CThostFtdcInputExecOrderField, nRequestID int) int {
}

// / 执行宣告操作请求
func (api *ThostFtdcTraderApi) ReqExecOrderAction(pInputExecOrderAction *thost.CThostFtdcInputExecOrderActionField, nRequestID int) int {
}

// / 询价录入请求
func (api *ThostFtdcTraderApi) ReqForQuoteInsert(pInputForQuote *thost.CThostFtdcInputForQuoteField, nRequestID int) int {
}

// / 报价录入请求
func (api *ThostFtdcTraderApi) ReqQuoteInsert(pInputQuote *thost.CThostFtdcInputQuoteField, nRequestID int) int {
}

// / 报价操作请求
func (api *ThostFtdcTraderApi) ReqQuoteAction(pInputQuoteAction *thost.CThostFtdcInputQuoteActionField, nRequestID int) int {
}

// / 批量报单操作请求
func (api *ThostFtdcTraderApi) ReqBatchOrderAction(pInputBatchOrderAction *thost.CThostFtdcInputBatchOrderActionField, nRequestID int) int {
}

// / 期权自对冲录入请求
func (api *ThostFtdcTraderApi) ReqOptionSelfCloseInsert(pInputOptionSelfClose *thost.CThostFtdcInputOptionSelfCloseField, nRequestID int) int {
}

// / 期权自对冲操作请求
func (api *ThostFtdcTraderApi) ReqOptionSelfCloseAction(pInputOptionSelfCloseAction *thost.CThostFtdcInputOptionSelfCloseActionField, nRequestID int) int {
}

// / 申请组合录入请求
func (api *ThostFtdcTraderApi) ReqCombActionInsert(pInputCombAction *thost.CThostFtdcInputCombActionField, nRequestID int) int {
}

// / 请求查询报单
func (api *ThostFtdcTraderApi) ReqQryOrder(pQryOrder *thost.CThostFtdcQryOrderField, nRequestID int) int {
}

// / 请求查询成交
func (api *ThostFtdcTraderApi) ReqQryTrade(pQryTrade *thost.CThostFtdcQryTradeField, nRequestID int) int {
}

// / 请求查询投资者持仓
func (api *ThostFtdcTraderApi) ReqQryInvestorPosition(pQryInvestorPosition *thost.CThostFtdcQryInvestorPositionField, nRequestID int) int {
}

// / 请求查询资金账户
func (api *ThostFtdcTraderApi) ReqQryTradingAccount(pQryTradingAccount *thost.CThostFtdcQryTradingAccountField, nRequestID int) int {
}

// / 请求查询投资者
func (api *ThostFtdcTraderApi) ReqQryInvestor(pQryInvestor *thost.CThostFtdcQryInvestorField, nRequestID int) int {
}

// / 请求查询交易编码
func (api *ThostFtdcTraderApi) ReqQryTradingCode(pQryTradingCode *thost.CThostFtdcQryTradingCodeField, nRequestID int) int {
}

// / 请求查询合约保证金率
func (api *ThostFtdcTraderApi) ReqQryInstrumentMarginRate(pQryInstrumentMarginRate *thost.CThostFtdcQryInstrumentMarginRateField, nRequestID int) int {
}

// / 请求查询合约手续费率
func (api *ThostFtdcTraderApi) ReqQryInstrumentCommissionRate(pQryInstrumentCommissionRate *thost.CThostFtdcQryInstrumentCommissionRateField, nRequestID int) int {
}

// / 请求查询用户会话
func (api *ThostFtdcTraderApi) ReqQryUserSession(pQryUserSession *thost.CThostFtdcQryUserSessionField, nRequestID int) int {
}

// / 请求查询交易所
func (api *ThostFtdcTraderApi) ReqQryExchange(pQryExchange *thost.CThostFtdcQryExchangeField, nRequestID int) int {
}

// / 请求查询产品
func (api *ThostFtdcTraderApi) ReqQryProduct(pQryProduct *thost.CThostFtdcQryProductField, nRequestID int) int {
}

// / 请求查询合约
func (api *ThostFtdcTraderApi) ReqQryInstrument(pQryInstrument *thost.CThostFtdcQryInstrumentField, nRequestID int) int {
}

// / 请求查询行情
func (api *ThostFtdcTraderApi) ReqQryDepthMarketData(pQryDepthMarketData *thost.CThostFtdcQryDepthMarketDataField, nRequestID int) int {
}

// / 请求查询交易员报盘机
func (api *ThostFtdcTraderApi) ReqQryTraderOffer(pQryTraderOffer *thost.CThostFtdcQryTraderOfferField, nRequestID int) int {
}

// / 请求查询投资者结算结果
func (api *ThostFtdcTraderApi) ReqQrySettlementInfo(pQrySettlementInfo *thost.CThostFtdcQrySettlementInfoField, nRequestID int) int {
}

// / 请求查询转帐银行
func (api *ThostFtdcTraderApi) ReqQryTransferBank(pQryTransferBank *thost.CThostFtdcQryTransferBankField, nRequestID int) int {
}

// / 请求查询投资者持仓明细
func (api *ThostFtdcTraderApi) ReqQryInvestorPositionDetail(pQryInvestorPositionDetail *thost.CThostFtdcQryInvestorPositionDetailField, nRequestID int) int {
}

// / 请求查询客户通知
func (api *ThostFtdcTraderApi) ReqQryNotice(pQryNotice *thost.CThostFtdcQryNoticeField, nRequestID int) int {
}

// / 请求查询结算信息确认
func (api *ThostFtdcTraderApi) ReqQrySettlementInfoConfirm(pQrySettlementInfoConfirm *thost.CThostFtdcQrySettlementInfoConfirmField, nRequestID int) int {
}

// / 请求查询投资者持仓明细
func (api *ThostFtdcTraderApi) ReqQryInvestorPositionCombineDetail(pQryInvestorPositionCombineDetail *thost.CThostFtdcQryInvestorPositionCombineDetailField, nRequestID int) int {
}

// / 请求查询保证金监管系统经纪公司资金账户密钥
func (api *ThostFtdcTraderApi) ReqQryCFMMCTradingAccountKey(pQryCFMMCTradingAccountKey *thost.CThostFtdcQryCFMMCTradingAccountKeyField, nRequestID int) int {
}

// / 请求查询仓单折抵信息
func (api *ThostFtdcTraderApi) ReqQryEWarrantOffset(pQryEWarrantOffset *thost.CThostFtdcQryEWarrantOffsetField, nRequestID int) int {
}

// / 请求查询投资者品种/跨品种保证金
func (api *ThostFtdcTraderApi) ReqQryInvestorProductGroupMargin(pQryInvestorProductGroupMargin *thost.CThostFtdcQryInvestorProductGroupMarginField, nRequestID int) int {
}

// / 请求查询交易所保证金率
func (api *ThostFtdcTraderApi) ReqQryExchangeMarginRate(pQryExchangeMarginRate *thost.CThostFtdcQryExchangeMarginRateField, nRequestID int) int {
}

// / 请求查询交易所调整保证金率
func (api *ThostFtdcTraderApi) ReqQryExchangeMarginRateAdjust(pQryExchangeMarginRateAdjust *thost.CThostFtdcQryExchangeMarginRateAdjustField, nRequestID int) int {
}

// / 请求查询汇率
func (api *ThostFtdcTraderApi) ReqQryExchangeRate(pQryExchangeRate *thost.CThostFtdcQryExchangeRateField, nRequestID int) int {
}

// / 请求查询二级代理操作员银期权限
func (api *ThostFtdcTraderApi) ReqQrySecAgentACIDMap(pQrySecAgentACIDMap *thost.CThostFtdcQrySecAgentACIDMapField, nRequestID int) int {
}

// / 请求查询产品报价汇率
func (api *ThostFtdcTraderApi) ReqQryProductExchRate(pQryProductExchRate *thost.CThostFtdcQryProductExchRateField, nRequestID int) int {
}

// / 请求查询产品组
func (api *ThostFtdcTraderApi) ReqQryProductGroup(pQryProductGroup *thost.CThostFtdcQryProductGroupField, nRequestID int) int {
}

// / 请求查询做市商合约手续费率
func (api *ThostFtdcTraderApi) ReqQryMMInstrumentCommissionRate(pQryMMInstrumentCommissionRate *thost.CThostFtdcQryMMInstrumentCommissionRateField, nRequestID int) int {
}

// / 请求查询做市商期权合约手续费
func (api *ThostFtdcTraderApi) ReqQryMMOptionInstrCommRate(pQryMMOptionInstrCommRate *thost.CThostFtdcQryMMOptionInstrCommRateField, nRequestID int) int {
}

// / 请求查询报单手续费
func (api *ThostFtdcTraderApi) ReqQryInstrumentOrderCommRate(pQryInstrumentOrderCommRate *thost.CThostFtdcQryInstrumentOrderCommRateField, nRequestID int) int {
}

// / 请求查询资金账户
func (api *ThostFtdcTraderApi) ReqQrySecAgentTradingAccount(pQryTradingAccount *thost.CThostFtdcQryTradingAccountField, nRequestID int) int {
}

// / 请求查询二级代理商资金校验模式
func (api *ThostFtdcTraderApi) ReqQrySecAgentCheckMode(pQrySecAgentCheckMode *thost.CThostFtdcQrySecAgentCheckModeField, nRequestID int) int {
}

// / 请求查询二级代理商信息
func (api *ThostFtdcTraderApi) ReqQrySecAgentTradeInfo(pQrySecAgentTradeInfo *thost.CThostFtdcQrySecAgentTradeInfoField, nRequestID int) int {
}

// / 请求查询期权交易成本
func (api *ThostFtdcTraderApi) ReqQryOptionInstrTradeCost(pQryOptionInstrTradeCost *thost.CThostFtdcQryOptionInstrTradeCostField, nRequestID int) int {
}

// / 请求查询期权合约手续费
func (api *ThostFtdcTraderApi) ReqQryOptionInstrCommRate(pQryOptionInstrCommRate *thost.CThostFtdcQryOptionInstrCommRateField, nRequestID int) int {
}

// / 请求查询执行宣告
func (api *ThostFtdcTraderApi) ReqQryExecOrder(pQryExecOrder *thost.CThostFtdcQryExecOrderField, nRequestID int) int {
}

// / 请求查询询价
func (api *ThostFtdcTraderApi) ReqQryForQuote(pQryForQuote *thost.CThostFtdcQryForQuoteField, nRequestID int) int {
}

// / 请求查询报价
func (api *ThostFtdcTraderApi) ReqQryQuote(pQryQuote *thost.CThostFtdcQryQuoteField, nRequestID int) int {
}

// / 请求查询期权自对冲
func (api *ThostFtdcTraderApi) ReqQryOptionSelfClose(pQryOptionSelfClose *thost.CThostFtdcQryOptionSelfCloseField, nRequestID int) int {
}

// / 请求查询投资单元
func (api *ThostFtdcTraderApi) ReqQryInvestUnit(pQryInvestUnit *thost.CThostFtdcQryInvestUnitField, nRequestID int) int {
}

// / 请求查询组合合约安全系数
func (api *ThostFtdcTraderApi) ReqQryCombInstrumentGuard(pQryCombInstrumentGuard *thost.CThostFtdcQryCombInstrumentGuardField, nRequestID int) int {
}

// / 请求查询申请组合
func (api *ThostFtdcTraderApi) ReqQryCombAction(pQryCombAction *thost.CThostFtdcQryCombActionField, nRequestID int) int {
}

// / 请求查询转帐流水
func (api *ThostFtdcTraderApi) ReqQryTransferSerial(pQryTransferSerial *thost.CThostFtdcQryTransferSerialField, nRequestID int) int {
}

// / 请求查询银期签约关系
func (api *ThostFtdcTraderApi) ReqQryAccountregister(pQryAccountregister *thost.CThostFtdcQryAccountregisterField, nRequestID int) int {
}

// / 请求查询签约银行
func (api *ThostFtdcTraderApi) ReqQryContractBank(pQryContractBank *thost.CThostFtdcQryContractBankField, nRequestID int) int {
}

// / 请求查询预埋单
func (api *ThostFtdcTraderApi) ReqQryParkedOrder(pQryParkedOrder *thost.CThostFtdcQryParkedOrderField, nRequestID int) int {
}

// / 请求查询预埋撤单
func (api *ThostFtdcTraderApi) ReqQryParkedOrderAction(pQryParkedOrderAction *thost.CThostFtdcQryParkedOrderActionField, nRequestID int) int {
}

// / 请求查询交易通知
func (api *ThostFtdcTraderApi) ReqQryTradingNotice(pQryTradingNotice *thost.CThostFtdcQryTradingNoticeField, nRequestID int) int {
}

// / 请求查询经纪公司交易参数
func (api *ThostFtdcTraderApi) ReqQryBrokerTradingParams(pQryBrokerTradingParams *thost.CThostFtdcQryBrokerTradingParamsField, nRequestID int) int {
}

// / 请求查询经纪公司交易算法
func (api *ThostFtdcTraderApi) ReqQryBrokerTradingAlgos(pQryBrokerTradingAlgos *thost.CThostFtdcQryBrokerTradingAlgosField, nRequestID int) int {
}

// / 请求查询监控中心用户令牌
func (api *ThostFtdcTraderApi) ReqQueryCFMMCTradingAccountToken(pQueryCFMMCTradingAccountToken *thost.CThostFtdcQueryCFMMCTradingAccountTokenField, nRequestID int) int {
}

// / 期货发起银行资金转期货请求
func (api *ThostFtdcTraderApi) ReqFromBankToFutureByFuture(pReqTransfer *thost.CThostFtdcReqTransferField, nRequestID int) int {
}

// / 期货发起期货资金转银行请求
func (api *ThostFtdcTraderApi) ReqFromFutureToBankByFuture(pReqTransfer *thost.CThostFtdcReqTransferField, nRequestID int) int {
}

// / 期货发起查询银行余额请求
func (api *ThostFtdcTraderApi) ReqQueryBankAccountMoneyByFuture(pReqQueryAccount *thost.CThostFtdcReqQueryAccountField, nRequestID int) int {
}

// / 请求查询分类合约
func (api *ThostFtdcTraderApi) ReqQryClassifiedInstrument(pQryClassifiedInstrument *thost.CThostFtdcQryClassifiedInstrumentField, nRequestID int) int {
}

// / 请求组合优惠比例
func (api *ThostFtdcTraderApi) ReqQryCombPromotionParam(pQryCombPromotionParam *thost.CThostFtdcQryCombPromotionParamField, nRequestID int) int {
}

// / 投资者风险结算持仓查询
func (api *ThostFtdcTraderApi) ReqQryRiskSettleInvstPosition(pQryRiskSettleInvstPosition *thost.CThostFtdcQryRiskSettleInvstPositionField, nRequestID int) int {
}

// / 风险结算产品查询
func (api *ThostFtdcTraderApi) ReqQryRiskSettleProductStatus(pQryRiskSettleProductStatus *thost.CThostFtdcQryRiskSettleProductStatusField, nRequestID int) int {
}

// / SPBM期货合约参数查询
func (api *ThostFtdcTraderApi) ReqQrySPBMFutureParameter(pQrySPBMFutureParameter *thost.CThostFtdcQrySPBMFutureParameterField, nRequestID int) int {
}

// / SPBM期权合约参数查询
func (api *ThostFtdcTraderApi) ReqQrySPBMOptionParameter(pQrySPBMOptionParameter *thost.CThostFtdcQrySPBMOptionParameterField, nRequestID int) int {
}

// / SPBM品种内对锁仓折扣参数查询
func (api *ThostFtdcTraderApi) ReqQrySPBMIntraParameter(pQrySPBMIntraParameter *thost.CThostFtdcQrySPBMIntraParameterField, nRequestID int) int {
}

// / SPBM跨品种抵扣参数查询
func (api *ThostFtdcTraderApi) ReqQrySPBMInterParameter(pQrySPBMInterParameter *thost.CThostFtdcQrySPBMInterParameterField, nRequestID int) int {
}

// / SPBM组合保证金套餐查询
func (api *ThostFtdcTraderApi) ReqQrySPBMPortfDefinition(pQrySPBMPortfDefinition *thost.CThostFtdcQrySPBMPortfDefinitionField, nRequestID int) int {
}

// / 投资者SPBM套餐选择查询
func (api *ThostFtdcTraderApi) ReqQrySPBMInvestorPortfDef(pQrySPBMInvestorPortfDef *thost.CThostFtdcQrySPBMInvestorPortfDefField, nRequestID int) int {
}

// / 投资者新型组合保证金系数查询
func (api *ThostFtdcTraderApi) ReqQryInvestorPortfMarginRatio(pQryInvestorPortfMarginRatio *thost.CThostFtdcQryInvestorPortfMarginRatioField, nRequestID int) int {
}

// / 投资者产品SPBM明细查询
func (api *ThostFtdcTraderApi) ReqQryInvestorProdSPBMDetail(pQryInvestorProdSPBMDetail *thost.CThostFtdcQryInvestorProdSPBMDetailField, nRequestID int) int {
}

// / 投资者商品组SPMM记录查询
func (api *ThostFtdcTraderApi) ReqQryInvestorCommoditySPMMMargin(pQryInvestorCommoditySPMMMargin *thost.CThostFtdcQryInvestorCommoditySPMMMarginField, nRequestID int) int {
}

// / 投资者商品群SPMM记录查询
func (api *ThostFtdcTraderApi) ReqQryInvestorCommodityGroupSPMMMargin(pQryInvestorCommodityGroupSPMMMargin *thost.CThostFtdcQryInvestorCommodityGroupSPMMMarginField, nRequestID int) int {
}

// / SPMM合约参数查询
func (api *ThostFtdcTraderApi) ReqQrySPMMInstParam(pQrySPMMInstParam *thost.CThostFtdcQrySPMMInstParamField, nRequestID int) int {
}

// / SPMM产品参数查询
func (api *ThostFtdcTraderApi) ReqQrySPMMProductParam(pQrySPMMProductParam *thost.CThostFtdcQrySPMMProductParamField, nRequestID int) int {
}

// / SPBM附加跨品种抵扣参数查询
func (api *ThostFtdcTraderApi) ReqQrySPBMAddOnInterParameter(pQrySPBMAddOnInterParameter *thost.CThostFtdcQrySPBMAddOnInterParameterField, nRequestID int) int {
}

// / RCAMS产品组合信息查询
func (api *ThostFtdcTraderApi) ReqQryRCAMSCombProductInfo(pQryRCAMSCombProductInfo *thost.CThostFtdcQryRCAMSCombProductInfoField, nRequestID int) int {
}

// / RCAMS同合约风险对冲参数查询
func (api *ThostFtdcTraderApi) ReqQryRCAMSInstrParameter(pQryRCAMSInstrParameter *thost.CThostFtdcQryRCAMSInstrParameterField, nRequestID int) int {
}

// / RCAMS品种内风险对冲参数查询
func (api *ThostFtdcTraderApi) ReqQryRCAMSIntraParameter(pQryRCAMSIntraParameter *thost.CThostFtdcQryRCAMSIntraParameterField, nRequestID int) int {
}

// / RCAMS跨品种风险折抵参数查询
func (api *ThostFtdcTraderApi) ReqQryRCAMSInterParameter(pQryRCAMSInterParameter *thost.CThostFtdcQryRCAMSInterParameterField, nRequestID int) int {
}

// / RCAMS空头期权风险调整参数查询
func (api *ThostFtdcTraderApi) ReqQryRCAMSShortOptAdjustParam(pQryRCAMSShortOptAdjustParam *thost.CThostFtdcQryRCAMSShortOptAdjustParamField, nRequestID int) int {
}

// / RCAMS策略组合持仓查询
func (api *ThostFtdcTraderApi) ReqQryRCAMSInvestorCombPosition(pQryRCAMSInvestorCombPosition *thost.CThostFtdcQryRCAMSInvestorCombPositionField, nRequestID int) int {
}

// / 投资者品种RCAMS保证金查询
func (api *ThostFtdcTraderApi) ReqQryInvestorProdRCAMSMargin(pQryInvestorProdRCAMSMargin *thost.CThostFtdcQryInvestorProdRCAMSMarginField, nRequestID int) int {
}

// / RULE合约保证金参数查询
func (api *ThostFtdcTraderApi) ReqQryRULEInstrParameter(pQryRULEInstrParameter *thost.CThostFtdcQryRULEInstrParameterField, nRequestID int) int {
}

// / RULE品种内对锁仓折扣参数查询
func (api *ThostFtdcTraderApi) ReqQryRULEIntraParameter(pQryRULEIntraParameter *thost.CThostFtdcQryRULEIntraParameterField, nRequestID int) int {
}

// / RULE跨品种抵扣参数查询
func (api *ThostFtdcTraderApi) ReqQryRULEInterParameter(pQryRULEInterParameter *thost.CThostFtdcQryRULEInterParameterField, nRequestID int) int {
}

// / 投资者产品RULE保证金查询
func (api *ThostFtdcTraderApi) ReqQryInvestorProdRULEMargin(pQryInvestorProdRULEMargin *thost.CThostFtdcQryInvestorProdRULEMarginField, nRequestID int) int {
}

// / 投资者新型组合保证金开关查询
func (api *ThostFtdcTraderApi) ReqQryInvestorPortfSetting(pQryInvestorPortfSetting *thost.CThostFtdcQryInvestorPortfSettingField, nRequestID int) int {
}

// / 投资者申报费阶梯收取记录查询
func (api *ThostFtdcTraderApi) ReqQryInvestorInfoCommRec(pQryInvestorInfoCommRec *thost.CThostFtdcQryInvestorInfoCommRecField, nRequestID int) int {
}

// / 组合腿信息查询
func (api *ThostFtdcTraderApi) ReqQryCombLeg(pQryCombLeg *thost.CThostFtdcQryCombLegField, nRequestID int) int {
}

// / 对冲设置请求
func (api *ThostFtdcTraderApi) ReqOffsetSetting(pInputOffsetSetting *thost.CThostFtdcInputOffsetSettingField, nRequestID int) int {
}

// / 对冲设置撤销请求
func (api *ThostFtdcTraderApi) ReqCancelOffsetSetting(pInputOffsetSetting *thost.CThostFtdcInputOffsetSettingField, nRequestID int) int {
}

// / 投资者对冲设置查询
func (api *ThostFtdcTraderApi) ReqQryOffsetSetting(pQryOffsetSetting *thost.CThostFtdcQryOffsetSettingField, nRequestID int) int {
}

// / 申请短信验证码请求
func (api *ThostFtdcTraderApi) ReqGenSMSCode(pReqGenSMSCode *thost.CThostFtdcReqGenSMSCodeField, nRequestID int) int {
}

// / 套利确认请求
func (api *ThostFtdcTraderApi) ReqSpdApply(pInputSpdApply *thost.CThostFtdcInputSpdApplyField, nRequestID int) int {
}

// / 套利确认撤销请求
func (api *ThostFtdcTraderApi) ReqSpdApplyAction(pInputSpdApplyAction *thost.CThostFtdcInputSpdApplyActionField, nRequestID int) int {
}

// / 套利确认查询请求
func (api *ThostFtdcTraderApi) ReqQrySpdApply(pQrySpdApply *thost.CThostFtdcQrySpdApplyField, nRequestID int) int {
}

// / 套保确认请求
func (api *ThostFtdcTraderApi) ReqHedgeCfm(pInputHedgeCfm *thost.CThostFtdcInputHedgeCfmField, nRequestID int) int {
}

// / 套保确认撤销请求
func (api *ThostFtdcTraderApi) ReqHedgeCfmAction(pInputHedgeCfmAction *thost.CThostFtdcInputHedgeCfmActionField, nRequestID int) int {
}

// / 套保确认查询请求
func (api *ThostFtdcTraderApi) ReqQryHedgeCfm(pQryHedgeCfm *thost.CThostFtdcQryHedgeCfmField, nRequestID int) int {
}
