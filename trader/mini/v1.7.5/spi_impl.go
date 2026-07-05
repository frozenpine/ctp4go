package v1_7_5

/*
#cgo CFLAGS: -I. -I${SRCDIR} -I${SRCDIR}/../../../dependencies/mini/v1.7.5/
#cgo LDFLAGS: -ldl

#include "spi_helper.h"
*/
import "C"
import (
	"context"
	"log/slog"
	"runtime"
	"unsafe"

	"github.com/frozenpine/ctp4go/thost/mini"
)

var (
	// 全局回调函数虚表
	spiCVtablePtr *C.CThostFtdcTraderSpiVTable
)

func init() {
	// C端为虚表分配内存
	spiCVtablePtr = (*C.CThostFtdcTraderSpiVTable)(C.malloc(
		C.sizeof_CThostFtdcTraderSpiVTable))

	spiCVtablePtr.CThostFtdcTraderSpiVTable_OnFrontConnected = (C.OnFrontConnected)(
		unsafe.Pointer(C.COnFrontConnected))

	spiCVtablePtr.CThostFtdcTraderSpiVTable_OnFrontDisconnected = (C.OnFrontDisconnected)(
		unsafe.Pointer(C.COnFrontDisconnected))

	spiCVtablePtr.CThostFtdcTraderSpiVTable_OnHeartBeatWarning = (C.OnHeartBeatWarning)(
		unsafe.Pointer(C.COnHeartBeatWarning))

	spiCVtablePtr.CThostFtdcTraderSpiVTable_OnRspSubscribeFlowCtrlWarning = (C.OnRspSubscribeFlowCtrlWarning)(
		unsafe.Pointer(C.COnRspSubscribeFlowCtrlWarning))

	spiCVtablePtr.CThostFtdcTraderSpiVTable_OnRspUnSubscribeFlowCtrlWarning = (C.OnRspUnSubscribeFlowCtrlWarning)(
		unsafe.Pointer(C.COnRspUnSubscribeFlowCtrlWarning))

	spiCVtablePtr.CThostFtdcTraderSpiVTable_OnRspAuthenticate = (C.OnRspAuthenticate)(
		unsafe.Pointer(C.COnRspAuthenticate))

	spiCVtablePtr.CThostFtdcTraderSpiVTable_OnRspUserLogin = (C.OnRspUserLogin)(
		unsafe.Pointer(C.COnRspUserLogin))

	spiCVtablePtr.CThostFtdcTraderSpiVTable_OnRspUserLogout = (C.OnRspUserLogout)(
		unsafe.Pointer(C.COnRspUserLogout))

	spiCVtablePtr.CThostFtdcTraderSpiVTable_OnRspOrderInsert = (C.OnRspOrderInsert)(
		unsafe.Pointer(C.COnRspOrderInsert))

	spiCVtablePtr.CThostFtdcTraderSpiVTable_OnRspOrderAction = (C.OnRspOrderAction)(
		unsafe.Pointer(C.COnRspOrderAction))

	spiCVtablePtr.CThostFtdcTraderSpiVTable_OnRspMKBatchOrderAction = (C.OnRspMKBatchOrderAction)(
		unsafe.Pointer(C.COnRspMKBatchOrderAction))

	spiCVtablePtr.CThostFtdcTraderSpiVTable_OnRspExecOrderInsert = (C.OnRspExecOrderInsert)(
		unsafe.Pointer(C.COnRspExecOrderInsert))

	spiCVtablePtr.CThostFtdcTraderSpiVTable_OnRspExecOrderAction = (C.OnRspExecOrderAction)(
		unsafe.Pointer(C.COnRspExecOrderAction))

	spiCVtablePtr.CThostFtdcTraderSpiVTable_OnRspForQuoteInsert = (C.OnRspForQuoteInsert)(
		unsafe.Pointer(C.COnRspForQuoteInsert))

	spiCVtablePtr.CThostFtdcTraderSpiVTable_OnRspQuoteInsert = (C.OnRspQuoteInsert)(
		unsafe.Pointer(C.COnRspQuoteInsert))

	spiCVtablePtr.CThostFtdcTraderSpiVTable_OnRspQuoteAction = (C.OnRspQuoteAction)(
		unsafe.Pointer(C.COnRspQuoteAction))

	spiCVtablePtr.CThostFtdcTraderSpiVTable_OnRspBatchOrderAction = (C.OnRspBatchOrderAction)(
		unsafe.Pointer(C.COnRspBatchOrderAction))

	spiCVtablePtr.CThostFtdcTraderSpiVTable_OnRspOptionSelfCloseInsert = (C.OnRspOptionSelfCloseInsert)(
		unsafe.Pointer(C.COnRspOptionSelfCloseInsert))

	spiCVtablePtr.CThostFtdcTraderSpiVTable_OnRspOptionSelfCloseAction = (C.OnRspOptionSelfCloseAction)(
		unsafe.Pointer(C.COnRspOptionSelfCloseAction))

	spiCVtablePtr.CThostFtdcTraderSpiVTable_OnRspCombActionInsert = (C.OnRspCombActionInsert)(
		unsafe.Pointer(C.COnRspCombActionInsert))

	spiCVtablePtr.CThostFtdcTraderSpiVTable_OnRspQryOrder = (C.OnRspQryOrder)(
		unsafe.Pointer(C.COnRspQryOrder))

	spiCVtablePtr.CThostFtdcTraderSpiVTable_OnRspQryTrade = (C.OnRspQryTrade)(
		unsafe.Pointer(C.COnRspQryTrade))

	spiCVtablePtr.CThostFtdcTraderSpiVTable_OnRspQryInvestorPosition = (C.OnRspQryInvestorPosition)(
		unsafe.Pointer(C.COnRspQryInvestorPosition))

	spiCVtablePtr.CThostFtdcTraderSpiVTable_OnRspQryTradingAccount = (C.OnRspQryTradingAccount)(
		unsafe.Pointer(C.COnRspQryTradingAccount))

	spiCVtablePtr.CThostFtdcTraderSpiVTable_OnRspQryInvestor = (C.OnRspQryInvestor)(
		unsafe.Pointer(C.COnRspQryInvestor))

	spiCVtablePtr.CThostFtdcTraderSpiVTable_OnRspQryTradingCode = (C.OnRspQryTradingCode)(
		unsafe.Pointer(C.COnRspQryTradingCode))

	spiCVtablePtr.CThostFtdcTraderSpiVTable_OnRspQryInstrumentMarginRate = (C.OnRspQryInstrumentMarginRate)(
		unsafe.Pointer(C.COnRspQryInstrumentMarginRate))

	spiCVtablePtr.CThostFtdcTraderSpiVTable_OnRspQryInstrumentCommissionRate = (C.OnRspQryInstrumentCommissionRate)(
		unsafe.Pointer(C.COnRspQryInstrumentCommissionRate))

	spiCVtablePtr.CThostFtdcTraderSpiVTable_OnRspQryExchange = (C.OnRspQryExchange)(
		unsafe.Pointer(C.COnRspQryExchange))

	spiCVtablePtr.CThostFtdcTraderSpiVTable_OnRspQryProduct = (C.OnRspQryProduct)(
		unsafe.Pointer(C.COnRspQryProduct))

	spiCVtablePtr.CThostFtdcTraderSpiVTable_OnRspQryInstrument = (C.OnRspQryInstrument)(
		unsafe.Pointer(C.COnRspQryInstrument))

	spiCVtablePtr.CThostFtdcTraderSpiVTable_OnRspQryCombInstrument = (C.OnRspQryCombInstrument)(
		unsafe.Pointer(C.COnRspQryCombInstrument))

	spiCVtablePtr.CThostFtdcTraderSpiVTable_OnRspQryRCAMSInvestorProdMargin = (C.OnRspQryRCAMSInvestorProdMargin)(
		unsafe.Pointer(C.COnRspQryRCAMSInvestorProdMargin))

	spiCVtablePtr.CThostFtdcTraderSpiVTable_OnRspQryRCAMSInvestorCombPosition = (C.OnRspQryRCAMSInvestorCombPosition)(
		unsafe.Pointer(C.COnRspQryRCAMSInvestorCombPosition))

	spiCVtablePtr.CThostFtdcTraderSpiVTable_OnRspQryCombAction = (C.OnRspQryCombAction)(
		unsafe.Pointer(C.COnRspQryCombAction))

	spiCVtablePtr.CThostFtdcTraderSpiVTable_OnRspQryInvestorPositionForComb = (C.OnRspQryInvestorPositionForComb)(
		unsafe.Pointer(C.COnRspQryInvestorPositionForComb))

	spiCVtablePtr.CThostFtdcTraderSpiVTable_OnRspQryDepthMarketData = (C.OnRspQryDepthMarketData)(
		unsafe.Pointer(C.COnRspQryDepthMarketData))

	spiCVtablePtr.CThostFtdcTraderSpiVTable_OnRspQryInstrumentStatus = (C.OnRspQryInstrumentStatus)(
		unsafe.Pointer(C.COnRspQryInstrumentStatus))

	spiCVtablePtr.CThostFtdcTraderSpiVTable_OnRspQryInvestorPositionDetail = (C.OnRspQryInvestorPositionDetail)(
		unsafe.Pointer(C.COnRspQryInvestorPositionDetail))

	spiCVtablePtr.CThostFtdcTraderSpiVTable_OnRspQryExchangeMarginRate = (C.OnRspQryExchangeMarginRate)(
		unsafe.Pointer(C.COnRspQryExchangeMarginRate))

	spiCVtablePtr.CThostFtdcTraderSpiVTable_OnRspQryExchangeMarginRateAdjust = (C.OnRspQryExchangeMarginRateAdjust)(
		unsafe.Pointer(C.COnRspQryExchangeMarginRateAdjust))

	spiCVtablePtr.CThostFtdcTraderSpiVTable_OnRspQryOptionInstrTradeCost = (C.OnRspQryOptionInstrTradeCost)(
		unsafe.Pointer(C.COnRspQryOptionInstrTradeCost))

	spiCVtablePtr.CThostFtdcTraderSpiVTable_OnRspQryOptionInstrCommRate = (C.OnRspQryOptionInstrCommRate)(
		unsafe.Pointer(C.COnRspQryOptionInstrCommRate))

	spiCVtablePtr.CThostFtdcTraderSpiVTable_OnRspQryExecOrder = (C.OnRspQryExecOrder)(
		unsafe.Pointer(C.COnRspQryExecOrder))

	spiCVtablePtr.CThostFtdcTraderSpiVTable_OnRspQryForQuote = (C.OnRspQryForQuote)(
		unsafe.Pointer(C.COnRspQryForQuote))

	spiCVtablePtr.CThostFtdcTraderSpiVTable_OnRspQryForQuoteParam = (C.OnRspQryForQuoteParam)(
		unsafe.Pointer(C.COnRspQryForQuoteParam))

	spiCVtablePtr.CThostFtdcTraderSpiVTable_OnRspQryInvestorProdSPBMDetail = (C.OnRspQryInvestorProdSPBMDetail)(
		unsafe.Pointer(C.COnRspQryInvestorProdSPBMDetail))

	spiCVtablePtr.CThostFtdcTraderSpiVTable_OnRspQrySPMMInvestorCommodityGroupMargin = (C.OnRspQrySPMMInvestorCommodityGroupMargin)(
		unsafe.Pointer(C.COnRspQrySPMMInvestorCommodityGroupMargin))

	spiCVtablePtr.CThostFtdcTraderSpiVTable_OnRspQryRULEInvestorProdMargin = (C.OnRspQryRULEInvestorProdMargin)(
		unsafe.Pointer(C.COnRspQryRULEInvestorProdMargin))

	spiCVtablePtr.CThostFtdcTraderSpiVTable_OnRspQryTraderOffer = (C.OnRspQryTraderOffer)(
		unsafe.Pointer(C.COnRspQryTraderOffer))

	spiCVtablePtr.CThostFtdcTraderSpiVTable_OnRspQryQuote = (C.OnRspQryQuote)(
		unsafe.Pointer(C.COnRspQryQuote))

	spiCVtablePtr.CThostFtdcTraderSpiVTable_OnRspQryOptionSelfClose = (C.OnRspQryOptionSelfClose)(
		unsafe.Pointer(C.COnRspQryOptionSelfClose))

	spiCVtablePtr.CThostFtdcTraderSpiVTable_OnRspQryControlParam = (C.OnRspQryControlParam)(
		unsafe.Pointer(C.COnRspQryControlParam))

	spiCVtablePtr.CThostFtdcTraderSpiVTable_OnRspQryOffsetSetting = (C.OnRspQryOffsetSetting)(
		unsafe.Pointer(C.COnRspQryOffsetSetting))

	spiCVtablePtr.CThostFtdcTraderSpiVTable_OnRspError = (C.OnRspError)(
		unsafe.Pointer(C.COnRspError))

	spiCVtablePtr.CThostFtdcTraderSpiVTable_OnRtnOrder = (C.OnRtnOrder)(
		unsafe.Pointer(C.COnRtnOrder))

	spiCVtablePtr.CThostFtdcTraderSpiVTable_OnRtnTrade = (C.OnRtnTrade)(
		unsafe.Pointer(C.COnRtnTrade))

	spiCVtablePtr.CThostFtdcTraderSpiVTable_OnErrRtnOrderInsert = (C.OnErrRtnOrderInsert)(
		unsafe.Pointer(C.COnErrRtnOrderInsert))

	spiCVtablePtr.CThostFtdcTraderSpiVTable_OnErrRtnOrderAction = (C.OnErrRtnOrderAction)(
		unsafe.Pointer(C.COnErrRtnOrderAction))

	spiCVtablePtr.CThostFtdcTraderSpiVTable_OnRtnInstrumentStatus = (C.OnRtnInstrumentStatus)(
		unsafe.Pointer(C.COnRtnInstrumentStatus))

	spiCVtablePtr.CThostFtdcTraderSpiVTable_OnRtnExecOrder = (C.OnRtnExecOrder)(
		unsafe.Pointer(C.COnRtnExecOrder))

	spiCVtablePtr.CThostFtdcTraderSpiVTable_OnErrRtnExecOrderInsert = (C.OnErrRtnExecOrderInsert)(
		unsafe.Pointer(C.COnErrRtnExecOrderInsert))

	spiCVtablePtr.CThostFtdcTraderSpiVTable_OnErrRtnExecOrderAction = (C.OnErrRtnExecOrderAction)(
		unsafe.Pointer(C.COnErrRtnExecOrderAction))

	spiCVtablePtr.CThostFtdcTraderSpiVTable_OnErrRtnForQuoteInsert = (C.OnErrRtnForQuoteInsert)(
		unsafe.Pointer(C.COnErrRtnForQuoteInsert))

	spiCVtablePtr.CThostFtdcTraderSpiVTable_OnRtnQuote = (C.OnRtnQuote)(
		unsafe.Pointer(C.COnRtnQuote))

	spiCVtablePtr.CThostFtdcTraderSpiVTable_OnErrRtnQuoteInsert = (C.OnErrRtnQuoteInsert)(
		unsafe.Pointer(C.COnErrRtnQuoteInsert))

	spiCVtablePtr.CThostFtdcTraderSpiVTable_OnErrRtnQuoteAction = (C.OnErrRtnQuoteAction)(
		unsafe.Pointer(C.COnErrRtnQuoteAction))

	spiCVtablePtr.CThostFtdcTraderSpiVTable_OnRtnForQuoteRsp = (C.OnRtnForQuoteRsp)(
		unsafe.Pointer(C.COnRtnForQuoteRsp))

	spiCVtablePtr.CThostFtdcTraderSpiVTable_OnErrRtnBatchOrderAction = (C.OnErrRtnBatchOrderAction)(
		unsafe.Pointer(C.COnErrRtnBatchOrderAction))

	spiCVtablePtr.CThostFtdcTraderSpiVTable_OnRtnOptionSelfClose = (C.OnRtnOptionSelfClose)(
		unsafe.Pointer(C.COnRtnOptionSelfClose))

	spiCVtablePtr.CThostFtdcTraderSpiVTable_OnErrRtnOptionSelfCloseInsert = (C.OnErrRtnOptionSelfCloseInsert)(
		unsafe.Pointer(C.COnErrRtnOptionSelfCloseInsert))

	spiCVtablePtr.CThostFtdcTraderSpiVTable_OnErrRtnOptionSelfCloseAction = (C.OnErrRtnOptionSelfCloseAction)(
		unsafe.Pointer(C.COnErrRtnOptionSelfCloseAction))

	spiCVtablePtr.CThostFtdcTraderSpiVTable_OnRtnCombAction = (C.OnRtnCombAction)(
		unsafe.Pointer(C.COnRtnCombAction))

	spiCVtablePtr.CThostFtdcTraderSpiVTable_OnRspQryInstrumentOrderCommRate = (C.OnRspQryInstrumentOrderCommRate)(
		unsafe.Pointer(C.COnRspQryInstrumentOrderCommRate))

	spiCVtablePtr.CThostFtdcTraderSpiVTable_OnRtnFlowCtrlWarning = (C.OnRtnFlowCtrlWarning)(
		unsafe.Pointer(C.COnRtnFlowCtrlWarning))

	spiCVtablePtr.CThostFtdcTraderSpiVTable_OnRspSubscribeFundChange = (C.OnRspSubscribeFundChange)(
		unsafe.Pointer(C.COnRspSubscribeFundChange))

	spiCVtablePtr.CThostFtdcTraderSpiVTable_OnRspUnSubscribeFundChange = (C.OnRspUnSubscribeFundChange)(
		unsafe.Pointer(C.COnRspUnSubscribeFundChange))

	spiCVtablePtr.CThostFtdcTraderSpiVTable_OnRtnFundChange = (C.OnRtnFundChange)(
		unsafe.Pointer(C.COnRtnFundChange))

	spiCVtablePtr.CThostFtdcTraderSpiVTable_OnRspOffsetSetting = (C.OnRspOffsetSetting)(
		unsafe.Pointer(C.COnRspOffsetSetting))

	spiCVtablePtr.CThostFtdcTraderSpiVTable_OnRspCancelOffsetSetting = (C.OnRspCancelOffsetSetting)(
		unsafe.Pointer(C.COnRspCancelOffsetSetting))

	spiCVtablePtr.CThostFtdcTraderSpiVTable_OnRtnOffsetSetting = (C.OnRtnOffsetSetting)(
		unsafe.Pointer(C.COnRtnOffsetSetting))

	spiCVtablePtr.CThostFtdcTraderSpiVTable_OnErrRtnCancelOffsetSetting = (C.OnErrRtnCancelOffsetSetting)(
		unsafe.Pointer(C.COnErrRtnCancelOffsetSetting))

	spiCVtablePtr.CThostFtdcTraderSpiVTable_OnRspUserPasswordUpdate = (C.OnRspUserPasswordUpdate)(
		unsafe.Pointer(C.COnRspUserPasswordUpdate))

}

type ThostFtdcTraderSpi struct {
	runtime.Pinner
	callback mini.TraderSpi
}

//export CgoOnFrontConnected
func CgoOnFrontConnected(
	this unsafe.Pointer,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnFrontConnected called",
		slog.Any("this", this),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnFrontConnected()
}

//export CgoOnFrontDisconnected
func CgoOnFrontDisconnected(
	this unsafe.Pointer,
	nReason C.int,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnFrontDisconnected called",
		slog.Any("this", this),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnFrontDisconnected(
		int(nReason),
	)
}

//export CgoOnHeartBeatWarning
func CgoOnHeartBeatWarning(
	this unsafe.Pointer,
	nTimeLapse C.int,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnHeartBeatWarning called",
		slog.Any("this", this),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnHeartBeatWarning(
		int(nTimeLapse),
	)
}

//export CgoOnRspSubscribeFlowCtrlWarning
func CgoOnRspSubscribeFlowCtrlWarning(
	this unsafe.Pointer,
	pRspSubscribeTraderField *C.struct_CThostFtdcSpecificTraderField,
	pRspInfo *C.struct_CThostFtdcRspInfoField,
	nRequestID C.int,
	bIsLast C.bool,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRspSubscribeFlowCtrlWarning called",
		slog.Any("this", this),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnRspSubscribeFlowCtrlWarning(
		(*mini.CThostFtdcSpecificTraderField)(unsafe.Pointer(pRspSubscribeTraderField)),
		(*mini.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
		int(nRequestID),
		bool(bIsLast),
	)
}

//export CgoOnRspUnSubscribeFlowCtrlWarning
func CgoOnRspUnSubscribeFlowCtrlWarning(
	this unsafe.Pointer,
	pRspSubscribeTraderField *C.struct_CThostFtdcSpecificTraderField,
	pRspInfo *C.struct_CThostFtdcRspInfoField,
	nRequestID C.int,
	bIsLast C.bool,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRspUnSubscribeFlowCtrlWarning called",
		slog.Any("this", this),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnRspUnSubscribeFlowCtrlWarning(
		(*mini.CThostFtdcSpecificTraderField)(unsafe.Pointer(pRspSubscribeTraderField)),
		(*mini.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
		int(nRequestID),
		bool(bIsLast),
	)
}

//export CgoOnRspAuthenticate
func CgoOnRspAuthenticate(
	this unsafe.Pointer,
	pRspAuthenticateField *C.struct_CThostFtdcRspAuthenticateField,
	pRspInfo *C.struct_CThostFtdcRspInfoField,
	nRequestID C.int,
	bIsLast C.bool,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRspAuthenticate called",
		slog.Any("this", this),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnRspAuthenticate(
		(*mini.CThostFtdcRspAuthenticateField)(unsafe.Pointer(pRspAuthenticateField)),
		(*mini.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
		int(nRequestID),
		bool(bIsLast),
	)
}

//export CgoOnRspUserLogin
func CgoOnRspUserLogin(
	this unsafe.Pointer,
	pRspUserLogin *C.struct_CThostFtdcRspUserLoginField,
	pRspInfo *C.struct_CThostFtdcRspInfoField,
	nRequestID C.int,
	bIsLast C.bool,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRspUserLogin called",
		slog.Any("this", this),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnRspUserLogin(
		(*mini.CThostFtdcRspUserLoginField)(unsafe.Pointer(pRspUserLogin)),
		(*mini.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
		int(nRequestID),
		bool(bIsLast),
	)
}

//export CgoOnRspUserLogout
func CgoOnRspUserLogout(
	this unsafe.Pointer,
	pUserLogout *C.struct_CThostFtdcUserLogoutField,
	pRspInfo *C.struct_CThostFtdcRspInfoField,
	nRequestID C.int,
	bIsLast C.bool,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRspUserLogout called",
		slog.Any("this", this),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnRspUserLogout(
		(*mini.CThostFtdcUserLogoutField)(unsafe.Pointer(pUserLogout)),
		(*mini.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
		int(nRequestID),
		bool(bIsLast),
	)
}

//export CgoOnRspOrderInsert
func CgoOnRspOrderInsert(
	this unsafe.Pointer,
	pInputOrder *C.struct_CThostFtdcInputOrderField,
	pRspInfo *C.struct_CThostFtdcRspInfoField,
	nRequestID C.int,
	bIsLast C.bool,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRspOrderInsert called",
		slog.Any("this", this),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnRspOrderInsert(
		(*mini.CThostFtdcInputOrderField)(unsafe.Pointer(pInputOrder)),
		(*mini.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
		int(nRequestID),
		bool(bIsLast),
	)
}

//export CgoOnRspOrderAction
func CgoOnRspOrderAction(
	this unsafe.Pointer,
	pInputOrderAction *C.struct_CThostFtdcInputOrderActionField,
	pRspInfo *C.struct_CThostFtdcRspInfoField,
	nRequestID C.int,
	bIsLast C.bool,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRspOrderAction called",
		slog.Any("this", this),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnRspOrderAction(
		(*mini.CThostFtdcInputOrderActionField)(unsafe.Pointer(pInputOrderAction)),
		(*mini.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
		int(nRequestID),
		bool(bIsLast),
	)
}

//export CgoOnRspMKBatchOrderAction
func CgoOnRspMKBatchOrderAction(
	this unsafe.Pointer,
	pMKInputOrderAction *C.struct_CThostFtdcMKInputOrderActionField,
	pRspInfo *C.struct_CThostFtdcRspInfoField,
	nRequestID C.int,
	bIsLast C.bool,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRspMKBatchOrderAction called",
		slog.Any("this", this),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnRspMKBatchOrderAction(
		(*mini.CThostFtdcMKInputOrderActionField)(unsafe.Pointer(pMKInputOrderAction)),
		(*mini.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
		int(nRequestID),
		bool(bIsLast),
	)
}

//export CgoOnRspExecOrderInsert
func CgoOnRspExecOrderInsert(
	this unsafe.Pointer,
	pInputExecOrder *C.struct_CThostFtdcInputExecOrderField,
	pRspInfo *C.struct_CThostFtdcRspInfoField,
	nRequestID C.int,
	bIsLast C.bool,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRspExecOrderInsert called",
		slog.Any("this", this),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnRspExecOrderInsert(
		(*mini.CThostFtdcInputExecOrderField)(unsafe.Pointer(pInputExecOrder)),
		(*mini.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
		int(nRequestID),
		bool(bIsLast),
	)
}

//export CgoOnRspExecOrderAction
func CgoOnRspExecOrderAction(
	this unsafe.Pointer,
	pInputExecOrderAction *C.struct_CThostFtdcInputExecOrderActionField,
	pRspInfo *C.struct_CThostFtdcRspInfoField,
	nRequestID C.int,
	bIsLast C.bool,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRspExecOrderAction called",
		slog.Any("this", this),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnRspExecOrderAction(
		(*mini.CThostFtdcInputExecOrderActionField)(unsafe.Pointer(pInputExecOrderAction)),
		(*mini.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
		int(nRequestID),
		bool(bIsLast),
	)
}

//export CgoOnRspForQuoteInsert
func CgoOnRspForQuoteInsert(
	this unsafe.Pointer,
	pInputForQuote *C.struct_CThostFtdcInputForQuoteField,
	pRspInfo *C.struct_CThostFtdcRspInfoField,
	nRequestID C.int,
	bIsLast C.bool,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRspForQuoteInsert called",
		slog.Any("this", this),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnRspForQuoteInsert(
		(*mini.CThostFtdcInputForQuoteField)(unsafe.Pointer(pInputForQuote)),
		(*mini.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
		int(nRequestID),
		bool(bIsLast),
	)
}

//export CgoOnRspQuoteInsert
func CgoOnRspQuoteInsert(
	this unsafe.Pointer,
	pInputQuote *C.struct_CThostFtdcInputQuoteField,
	pRspInfo *C.struct_CThostFtdcRspInfoField,
	nRequestID C.int,
	bIsLast C.bool,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRspQuoteInsert called",
		slog.Any("this", this),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnRspQuoteInsert(
		(*mini.CThostFtdcInputQuoteField)(unsafe.Pointer(pInputQuote)),
		(*mini.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
		int(nRequestID),
		bool(bIsLast),
	)
}

//export CgoOnRspQuoteAction
func CgoOnRspQuoteAction(
	this unsafe.Pointer,
	pInputQuoteAction *C.struct_CThostFtdcInputQuoteActionField,
	pRspInfo *C.struct_CThostFtdcRspInfoField,
	nRequestID C.int,
	bIsLast C.bool,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRspQuoteAction called",
		slog.Any("this", this),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnRspQuoteAction(
		(*mini.CThostFtdcInputQuoteActionField)(unsafe.Pointer(pInputQuoteAction)),
		(*mini.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
		int(nRequestID),
		bool(bIsLast),
	)
}

//export CgoOnRspBatchOrderAction
func CgoOnRspBatchOrderAction(
	this unsafe.Pointer,
	pInputBatchOrderAction *C.struct_CThostFtdcInputBatchOrderActionField,
	pRspInfo *C.struct_CThostFtdcRspInfoField,
	nRequestID C.int,
	bIsLast C.bool,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRspBatchOrderAction called",
		slog.Any("this", this),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnRspBatchOrderAction(
		(*mini.CThostFtdcInputBatchOrderActionField)(unsafe.Pointer(pInputBatchOrderAction)),
		(*mini.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
		int(nRequestID),
		bool(bIsLast),
	)
}

//export CgoOnRspOptionSelfCloseInsert
func CgoOnRspOptionSelfCloseInsert(
	this unsafe.Pointer,
	pInputOptionSelfClose *C.struct_CThostFtdcInputOptionSelfCloseField,
	pRspInfo *C.struct_CThostFtdcRspInfoField,
	nRequestID C.int,
	bIsLast C.bool,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRspOptionSelfCloseInsert called",
		slog.Any("this", this),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnRspOptionSelfCloseInsert(
		(*mini.CThostFtdcInputOptionSelfCloseField)(unsafe.Pointer(pInputOptionSelfClose)),
		(*mini.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
		int(nRequestID),
		bool(bIsLast),
	)
}

//export CgoOnRspOptionSelfCloseAction
func CgoOnRspOptionSelfCloseAction(
	this unsafe.Pointer,
	pInputOptionSelfCloseAction *C.struct_CThostFtdcInputOptionSelfCloseActionField,
	pRspInfo *C.struct_CThostFtdcRspInfoField,
	nRequestID C.int,
	bIsLast C.bool,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRspOptionSelfCloseAction called",
		slog.Any("this", this),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnRspOptionSelfCloseAction(
		(*mini.CThostFtdcInputOptionSelfCloseActionField)(unsafe.Pointer(pInputOptionSelfCloseAction)),
		(*mini.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
		int(nRequestID),
		bool(bIsLast),
	)
}

//export CgoOnRspCombActionInsert
func CgoOnRspCombActionInsert(
	this unsafe.Pointer,
	pInputCombAction *C.struct_CThostFtdcInputCombActionField,
	pRspInfo *C.struct_CThostFtdcRspInfoField,
	nRequestID C.int,
	bIsLast C.bool,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRspCombActionInsert called",
		slog.Any("this", this),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnRspCombActionInsert(
		(*mini.CThostFtdcInputCombActionField)(unsafe.Pointer(pInputCombAction)),
		(*mini.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
		int(nRequestID),
		bool(bIsLast),
	)
}

//export CgoOnRspQryOrder
func CgoOnRspQryOrder(
	this unsafe.Pointer,
	pOrder *C.struct_CThostFtdcOrderField,
	pRspInfo *C.struct_CThostFtdcRspInfoField,
	nRequestID C.int,
	bIsLast C.bool,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRspQryOrder called",
		slog.Any("this", this),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnRspQryOrder(
		(*mini.CThostFtdcOrderField)(unsafe.Pointer(pOrder)),
		(*mini.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
		int(nRequestID),
		bool(bIsLast),
	)
}

//export CgoOnRspQryTrade
func CgoOnRspQryTrade(
	this unsafe.Pointer,
	pTrade *C.struct_CThostFtdcTradeField,
	pRspInfo *C.struct_CThostFtdcRspInfoField,
	nRequestID C.int,
	bIsLast C.bool,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRspQryTrade called",
		slog.Any("this", this),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnRspQryTrade(
		(*mini.CThostFtdcTradeField)(unsafe.Pointer(pTrade)),
		(*mini.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
		int(nRequestID),
		bool(bIsLast),
	)
}

//export CgoOnRspQryInvestorPosition
func CgoOnRspQryInvestorPosition(
	this unsafe.Pointer,
	pInvestorPosition *C.struct_CThostFtdcInvestorPositionField,
	pRspInfo *C.struct_CThostFtdcRspInfoField,
	nRequestID C.int,
	bIsLast C.bool,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRspQryInvestorPosition called",
		slog.Any("this", this),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnRspQryInvestorPosition(
		(*mini.CThostFtdcInvestorPositionField)(unsafe.Pointer(pInvestorPosition)),
		(*mini.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
		int(nRequestID),
		bool(bIsLast),
	)
}

//export CgoOnRspQryTradingAccount
func CgoOnRspQryTradingAccount(
	this unsafe.Pointer,
	pTradingAccount *C.struct_CThostFtdcTradingAccountField,
	pRspInfo *C.struct_CThostFtdcRspInfoField,
	nRequestID C.int,
	bIsLast C.bool,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRspQryTradingAccount called",
		slog.Any("this", this),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnRspQryTradingAccount(
		(*mini.CThostFtdcTradingAccountField)(unsafe.Pointer(pTradingAccount)),
		(*mini.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
		int(nRequestID),
		bool(bIsLast),
	)
}

//export CgoOnRspQryInvestor
func CgoOnRspQryInvestor(
	this unsafe.Pointer,
	pInvestor *C.struct_CThostFtdcInvestorField,
	pRspInfo *C.struct_CThostFtdcRspInfoField,
	nRequestID C.int,
	bIsLast C.bool,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRspQryInvestor called",
		slog.Any("this", this),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnRspQryInvestor(
		(*mini.CThostFtdcInvestorField)(unsafe.Pointer(pInvestor)),
		(*mini.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
		int(nRequestID),
		bool(bIsLast),
	)
}

//export CgoOnRspQryTradingCode
func CgoOnRspQryTradingCode(
	this unsafe.Pointer,
	pTradingCode *C.struct_CThostFtdcTradingCodeField,
	pRspInfo *C.struct_CThostFtdcRspInfoField,
	nRequestID C.int,
	bIsLast C.bool,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRspQryTradingCode called",
		slog.Any("this", this),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnRspQryTradingCode(
		(*mini.CThostFtdcTradingCodeField)(unsafe.Pointer(pTradingCode)),
		(*mini.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
		int(nRequestID),
		bool(bIsLast),
	)
}

//export CgoOnRspQryInstrumentMarginRate
func CgoOnRspQryInstrumentMarginRate(
	this unsafe.Pointer,
	pInstrumentMarginRate *C.struct_CThostFtdcInstrumentMarginRateField,
	pRspInfo *C.struct_CThostFtdcRspInfoField,
	nRequestID C.int,
	bIsLast C.bool,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRspQryInstrumentMarginRate called",
		slog.Any("this", this),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnRspQryInstrumentMarginRate(
		(*mini.CThostFtdcInstrumentMarginRateField)(unsafe.Pointer(pInstrumentMarginRate)),
		(*mini.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
		int(nRequestID),
		bool(bIsLast),
	)
}

//export CgoOnRspQryInstrumentCommissionRate
func CgoOnRspQryInstrumentCommissionRate(
	this unsafe.Pointer,
	pInstrumentCommissionRate *C.struct_CThostFtdcInstrumentCommissionRateField,
	pRspInfo *C.struct_CThostFtdcRspInfoField,
	nRequestID C.int,
	bIsLast C.bool,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRspQryInstrumentCommissionRate called",
		slog.Any("this", this),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnRspQryInstrumentCommissionRate(
		(*mini.CThostFtdcInstrumentCommissionRateField)(unsafe.Pointer(pInstrumentCommissionRate)),
		(*mini.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
		int(nRequestID),
		bool(bIsLast),
	)
}

//export CgoOnRspQryExchange
func CgoOnRspQryExchange(
	this unsafe.Pointer,
	pExchange *C.struct_CThostFtdcExchangeField,
	pRspInfo *C.struct_CThostFtdcRspInfoField,
	nRequestID C.int,
	bIsLast C.bool,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRspQryExchange called",
		slog.Any("this", this),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnRspQryExchange(
		(*mini.CThostFtdcExchangeField)(unsafe.Pointer(pExchange)),
		(*mini.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
		int(nRequestID),
		bool(bIsLast),
	)
}

//export CgoOnRspQryProduct
func CgoOnRspQryProduct(
	this unsafe.Pointer,
	pProduct *C.struct_CThostFtdcProductField,
	pRspInfo *C.struct_CThostFtdcRspInfoField,
	nRequestID C.int,
	bIsLast C.bool,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRspQryProduct called",
		slog.Any("this", this),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnRspQryProduct(
		(*mini.CThostFtdcProductField)(unsafe.Pointer(pProduct)),
		(*mini.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
		int(nRequestID),
		bool(bIsLast),
	)
}

//export CgoOnRspQryInstrument
func CgoOnRspQryInstrument(
	this unsafe.Pointer,
	pInstrument *C.struct_CThostFtdcInstrumentField,
	pRspInfo *C.struct_CThostFtdcRspInfoField,
	nRequestID C.int,
	bIsLast C.bool,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRspQryInstrument called",
		slog.Any("this", this),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnRspQryInstrument(
		(*mini.CThostFtdcInstrumentField)(unsafe.Pointer(pInstrument)),
		(*mini.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
		int(nRequestID),
		bool(bIsLast),
	)
}

//export CgoOnRspQryCombInstrument
func CgoOnRspQryCombInstrument(
	this unsafe.Pointer,
	pCombInstrument *C.struct_CThostFtdcCombInstrumentField,
	pRspInfo *C.struct_CThostFtdcRspInfoField,
	nRequestID C.int,
	bIsLast C.bool,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRspQryCombInstrument called",
		slog.Any("this", this),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnRspQryCombInstrument(
		(*mini.CThostFtdcCombInstrumentField)(unsafe.Pointer(pCombInstrument)),
		(*mini.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
		int(nRequestID),
		bool(bIsLast),
	)
}

//export CgoOnRspQryRCAMSInvestorProdMargin
func CgoOnRspQryRCAMSInvestorProdMargin(
	this unsafe.Pointer,
	pRCAMSInvestorProdMargin *C.struct_CThostFtdcRCAMSInvestorProdMarginField,
	pRspInfo *C.struct_CThostFtdcRspInfoField,
	nRequestID C.int,
	bIsLast C.bool,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRspQryRCAMSInvestorProdMargin called",
		slog.Any("this", this),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnRspQryRCAMSInvestorProdMargin(
		(*mini.CThostFtdcRCAMSInvestorProdMarginField)(unsafe.Pointer(pRCAMSInvestorProdMargin)),
		(*mini.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
		int(nRequestID),
		bool(bIsLast),
	)
}

//export CgoOnRspQryRCAMSInvestorCombPosition
func CgoOnRspQryRCAMSInvestorCombPosition(
	this unsafe.Pointer,
	pRCAMSInvestorCombPosition *C.struct_CThostFtdcRCAMSInvestorCombPositionField,
	pRspInfo *C.struct_CThostFtdcRspInfoField,
	nRequestID C.int,
	bIsLast C.bool,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRspQryRCAMSInvestorCombPosition called",
		slog.Any("this", this),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnRspQryRCAMSInvestorCombPosition(
		(*mini.CThostFtdcRCAMSInvestorCombPositionField)(unsafe.Pointer(pRCAMSInvestorCombPosition)),
		(*mini.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
		int(nRequestID),
		bool(bIsLast),
	)
}

//export CgoOnRspQryCombAction
func CgoOnRspQryCombAction(
	this unsafe.Pointer,
	pCombAction *C.struct_CThostFtdcCombActionField,
	pRspInfo *C.struct_CThostFtdcRspInfoField,
	nRequestID C.int,
	bIsLast C.bool,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRspQryCombAction called",
		slog.Any("this", this),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnRspQryCombAction(
		(*mini.CThostFtdcCombActionField)(unsafe.Pointer(pCombAction)),
		(*mini.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
		int(nRequestID),
		bool(bIsLast),
	)
}

//export CgoOnRspQryInvestorPositionForComb
func CgoOnRspQryInvestorPositionForComb(
	this unsafe.Pointer,
	pForComb *C.struct_CThostFtdcInvestorPositionForCombField,
	pRspInfo *C.struct_CThostFtdcRspInfoField,
	nRequestID C.int,
	bIsLast C.bool,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRspQryInvestorPositionForComb called",
		slog.Any("this", this),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnRspQryInvestorPositionForComb(
		(*mini.CThostFtdcInvestorPositionForCombField)(unsafe.Pointer(pForComb)),
		(*mini.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
		int(nRequestID),
		bool(bIsLast),
	)
}

//export CgoOnRspQryDepthMarketData
func CgoOnRspQryDepthMarketData(
	this unsafe.Pointer,
	pDepthMarketData *C.struct_CThostFtdcDepthMarketDataField,
	pRspInfo *C.struct_CThostFtdcRspInfoField,
	nRequestID C.int,
	bIsLast C.bool,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRspQryDepthMarketData called",
		slog.Any("this", this),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnRspQryDepthMarketData(
		(*mini.CThostFtdcDepthMarketDataField)(unsafe.Pointer(pDepthMarketData)),
		(*mini.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
		int(nRequestID),
		bool(bIsLast),
	)
}

//export CgoOnRspQryInstrumentStatus
func CgoOnRspQryInstrumentStatus(
	this unsafe.Pointer,
	pInstrumentStatus *C.struct_CThostFtdcInstrumentStatusField,
	pRspInfo *C.struct_CThostFtdcRspInfoField,
	nRequestID C.int,
	bIsLast C.bool,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRspQryInstrumentStatus called",
		slog.Any("this", this),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnRspQryInstrumentStatus(
		(*mini.CThostFtdcInstrumentStatusField)(unsafe.Pointer(pInstrumentStatus)),
		(*mini.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
		int(nRequestID),
		bool(bIsLast),
	)
}

//export CgoOnRspQryInvestorPositionDetail
func CgoOnRspQryInvestorPositionDetail(
	this unsafe.Pointer,
	pInvestorPositionDetail *C.struct_CThostFtdcInvestorPositionDetailField,
	pRspInfo *C.struct_CThostFtdcRspInfoField,
	nRequestID C.int,
	bIsLast C.bool,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRspQryInvestorPositionDetail called",
		slog.Any("this", this),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnRspQryInvestorPositionDetail(
		(*mini.CThostFtdcInvestorPositionDetailField)(unsafe.Pointer(pInvestorPositionDetail)),
		(*mini.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
		int(nRequestID),
		bool(bIsLast),
	)
}

//export CgoOnRspQryExchangeMarginRate
func CgoOnRspQryExchangeMarginRate(
	this unsafe.Pointer,
	pExchangeMarginRate *C.struct_CThostFtdcExchangeMarginRateField,
	pRspInfo *C.struct_CThostFtdcRspInfoField,
	nRequestID C.int,
	bIsLast C.bool,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRspQryExchangeMarginRate called",
		slog.Any("this", this),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnRspQryExchangeMarginRate(
		(*mini.CThostFtdcExchangeMarginRateField)(unsafe.Pointer(pExchangeMarginRate)),
		(*mini.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
		int(nRequestID),
		bool(bIsLast),
	)
}

//export CgoOnRspQryExchangeMarginRateAdjust
func CgoOnRspQryExchangeMarginRateAdjust(
	this unsafe.Pointer,
	pExchangeMarginRateAdjust *C.struct_CThostFtdcExchangeMarginRateAdjustField,
	pRspInfo *C.struct_CThostFtdcRspInfoField,
	nRequestID C.int,
	bIsLast C.bool,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRspQryExchangeMarginRateAdjust called",
		slog.Any("this", this),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnRspQryExchangeMarginRateAdjust(
		(*mini.CThostFtdcExchangeMarginRateAdjustField)(unsafe.Pointer(pExchangeMarginRateAdjust)),
		(*mini.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
		int(nRequestID),
		bool(bIsLast),
	)
}

//export CgoOnRspQryOptionInstrTradeCost
func CgoOnRspQryOptionInstrTradeCost(
	this unsafe.Pointer,
	pOptionInstrTradeCost *C.struct_CThostFtdcOptionInstrTradeCostField,
	pRspInfo *C.struct_CThostFtdcRspInfoField,
	nRequestID C.int,
	bIsLast C.bool,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRspQryOptionInstrTradeCost called",
		slog.Any("this", this),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnRspQryOptionInstrTradeCost(
		(*mini.CThostFtdcOptionInstrTradeCostField)(unsafe.Pointer(pOptionInstrTradeCost)),
		(*mini.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
		int(nRequestID),
		bool(bIsLast),
	)
}

//export CgoOnRspQryOptionInstrCommRate
func CgoOnRspQryOptionInstrCommRate(
	this unsafe.Pointer,
	pOptionInstrCommRate *C.struct_CThostFtdcOptionInstrCommRateField,
	pRspInfo *C.struct_CThostFtdcRspInfoField,
	nRequestID C.int,
	bIsLast C.bool,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRspQryOptionInstrCommRate called",
		slog.Any("this", this),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnRspQryOptionInstrCommRate(
		(*mini.CThostFtdcOptionInstrCommRateField)(unsafe.Pointer(pOptionInstrCommRate)),
		(*mini.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
		int(nRequestID),
		bool(bIsLast),
	)
}

//export CgoOnRspQryExecOrder
func CgoOnRspQryExecOrder(
	this unsafe.Pointer,
	pExecOrder *C.struct_CThostFtdcExecOrderField,
	pRspInfo *C.struct_CThostFtdcRspInfoField,
	nRequestID C.int,
	bIsLast C.bool,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRspQryExecOrder called",
		slog.Any("this", this),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnRspQryExecOrder(
		(*mini.CThostFtdcExecOrderField)(unsafe.Pointer(pExecOrder)),
		(*mini.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
		int(nRequestID),
		bool(bIsLast),
	)
}

//export CgoOnRspQryForQuote
func CgoOnRspQryForQuote(
	this unsafe.Pointer,
	pForQuote *C.struct_CThostFtdcForQuoteField,
	pRspInfo *C.struct_CThostFtdcRspInfoField,
	nRequestID C.int,
	bIsLast C.bool,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRspQryForQuote called",
		slog.Any("this", this),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnRspQryForQuote(
		(*mini.CThostFtdcForQuoteField)(unsafe.Pointer(pForQuote)),
		(*mini.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
		int(nRequestID),
		bool(bIsLast),
	)
}

//export CgoOnRspQryForQuoteParam
func CgoOnRspQryForQuoteParam(
	this unsafe.Pointer,
	pForQuoteParam *C.struct_CThostFtdcForQuoteParamField,
	pRspInfo *C.struct_CThostFtdcRspInfoField,
	nRequestID C.int,
	bIsLast C.bool,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRspQryForQuoteParam called",
		slog.Any("this", this),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnRspQryForQuoteParam(
		(*mini.CThostFtdcForQuoteParamField)(unsafe.Pointer(pForQuoteParam)),
		(*mini.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
		int(nRequestID),
		bool(bIsLast),
	)
}

//export CgoOnRspQryInvestorProdSPBMDetail
func CgoOnRspQryInvestorProdSPBMDetail(
	this unsafe.Pointer,
	pInvestorProdSPBMDetail *C.struct_CThostFtdcInvestorProdSPBMDetailField,
	pRspInfo *C.struct_CThostFtdcRspInfoField,
	nRequestID C.int,
	bIsLast C.bool,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRspQryInvestorProdSPBMDetail called",
		slog.Any("this", this),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnRspQryInvestorProdSPBMDetail(
		(*mini.CThostFtdcInvestorProdSPBMDetailField)(unsafe.Pointer(pInvestorProdSPBMDetail)),
		(*mini.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
		int(nRequestID),
		bool(bIsLast),
	)
}

//export CgoOnRspQrySPMMInvestorCommodityGroupMargin
func CgoOnRspQrySPMMInvestorCommodityGroupMargin(
	this unsafe.Pointer,
	pSPMMInvestorCommodityGroupMargin *C.struct_CThostFtdcSPMMInvestorCommodityGroupMarginField,
	pRspInfo *C.struct_CThostFtdcRspInfoField,
	nRequestID C.int,
	bIsLast C.bool,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRspQrySPMMInvestorCommodityGroupMargin called",
		slog.Any("this", this),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnRspQrySPMMInvestorCommodityGroupMargin(
		(*mini.CThostFtdcSPMMInvestorCommodityGroupMarginField)(unsafe.Pointer(pSPMMInvestorCommodityGroupMargin)),
		(*mini.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
		int(nRequestID),
		bool(bIsLast),
	)
}

//export CgoOnRspQryRULEInvestorProdMargin
func CgoOnRspQryRULEInvestorProdMargin(
	this unsafe.Pointer,
	pRULEInvestorProdMargin *C.struct_CThostFtdcRULEInvestorProdMarginField,
	pRspInfo *C.struct_CThostFtdcRspInfoField,
	nRequestID C.int,
	bIsLast C.bool,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRspQryRULEInvestorProdMargin called",
		slog.Any("this", this),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnRspQryRULEInvestorProdMargin(
		(*mini.CThostFtdcRULEInvestorProdMarginField)(unsafe.Pointer(pRULEInvestorProdMargin)),
		(*mini.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
		int(nRequestID),
		bool(bIsLast),
	)
}

//export CgoOnRspQryTraderOffer
func CgoOnRspQryTraderOffer(
	this unsafe.Pointer,
	pTraderOffer *C.struct_CThostFtdcTraderOfferField,
	pRspInfo *C.struct_CThostFtdcRspInfoField,
	nRequestID C.int,
	bIsLast C.bool,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRspQryTraderOffer called",
		slog.Any("this", this),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnRspQryTraderOffer(
		(*mini.CThostFtdcTraderOfferField)(unsafe.Pointer(pTraderOffer)),
		(*mini.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
		int(nRequestID),
		bool(bIsLast),
	)
}

//export CgoOnRspQryQuote
func CgoOnRspQryQuote(
	this unsafe.Pointer,
	pQuote *C.struct_CThostFtdcQuoteField,
	pRspInfo *C.struct_CThostFtdcRspInfoField,
	nRequestID C.int,
	bIsLast C.bool,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRspQryQuote called",
		slog.Any("this", this),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnRspQryQuote(
		(*mini.CThostFtdcQuoteField)(unsafe.Pointer(pQuote)),
		(*mini.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
		int(nRequestID),
		bool(bIsLast),
	)
}

//export CgoOnRspQryOptionSelfClose
func CgoOnRspQryOptionSelfClose(
	this unsafe.Pointer,
	pOptionSelfClose *C.struct_CThostFtdcOptionSelfCloseField,
	pRspInfo *C.struct_CThostFtdcRspInfoField,
	nRequestID C.int,
	bIsLast C.bool,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRspQryOptionSelfClose called",
		slog.Any("this", this),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnRspQryOptionSelfClose(
		(*mini.CThostFtdcOptionSelfCloseField)(unsafe.Pointer(pOptionSelfClose)),
		(*mini.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
		int(nRequestID),
		bool(bIsLast),
	)
}

//export CgoOnRspQryControlParam
func CgoOnRspQryControlParam(
	this unsafe.Pointer,
	pControlParam *C.struct_CThostFtdcControlParamField,
	pRspInfo *C.struct_CThostFtdcRspInfoField,
	nRequestID C.int,
	bIsLast C.bool,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRspQryControlParam called",
		slog.Any("this", this),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnRspQryControlParam(
		(*mini.CThostFtdcControlParamField)(unsafe.Pointer(pControlParam)),
		(*mini.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
		int(nRequestID),
		bool(bIsLast),
	)
}

//export CgoOnRspQryOffsetSetting
func CgoOnRspQryOffsetSetting(
	this unsafe.Pointer,
	pOffsetSetting *C.struct_CThostFtdcOffsetSettingField,
	pRspInfo *C.struct_CThostFtdcRspInfoField,
	nRequestID C.int,
	bIsLast C.bool,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRspQryOffsetSetting called",
		slog.Any("this", this),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnRspQryOffsetSetting(
		(*mini.CThostFtdcOffsetSettingField)(unsafe.Pointer(pOffsetSetting)),
		(*mini.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
		int(nRequestID),
		bool(bIsLast),
	)
}

//export CgoOnRspError
func CgoOnRspError(
	this unsafe.Pointer,
	pRspInfo *C.struct_CThostFtdcRspInfoField,
	nRequestID C.int,
	bIsLast C.bool,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRspError called",
		slog.Any("this", this),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnRspError(
		(*mini.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
		int(nRequestID),
		bool(bIsLast),
	)
}

//export CgoOnRtnOrder
func CgoOnRtnOrder(
	this unsafe.Pointer,
	pOrder *C.struct_CThostFtdcOrderField,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRtnOrder called",
		slog.Any("this", this),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnRtnOrder(
		(*mini.CThostFtdcOrderField)(unsafe.Pointer(pOrder)),
	)
}

//export CgoOnRtnTrade
func CgoOnRtnTrade(
	this unsafe.Pointer,
	pTrade *C.struct_CThostFtdcTradeField,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRtnTrade called",
		slog.Any("this", this),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnRtnTrade(
		(*mini.CThostFtdcTradeField)(unsafe.Pointer(pTrade)),
	)
}

//export CgoOnErrRtnOrderInsert
func CgoOnErrRtnOrderInsert(
	this unsafe.Pointer,
	pInputOrder *C.struct_CThostFtdcInputOrderField,
	pRspInfo *C.struct_CThostFtdcRspInfoField,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnErrRtnOrderInsert called",
		slog.Any("this", this),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnErrRtnOrderInsert(
		(*mini.CThostFtdcInputOrderField)(unsafe.Pointer(pInputOrder)),
		(*mini.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
	)
}

//export CgoOnErrRtnOrderAction
func CgoOnErrRtnOrderAction(
	this unsafe.Pointer,
	pOrderAction *C.struct_CThostFtdcOrderActionField,
	pRspInfo *C.struct_CThostFtdcRspInfoField,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnErrRtnOrderAction called",
		slog.Any("this", this),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnErrRtnOrderAction(
		(*mini.CThostFtdcOrderActionField)(unsafe.Pointer(pOrderAction)),
		(*mini.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
	)
}

//export CgoOnRtnInstrumentStatus
func CgoOnRtnInstrumentStatus(
	this unsafe.Pointer,
	pInstrumentStatus *C.struct_CThostFtdcInstrumentStatusField,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRtnInstrumentStatus called",
		slog.Any("this", this),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnRtnInstrumentStatus(
		(*mini.CThostFtdcInstrumentStatusField)(unsafe.Pointer(pInstrumentStatus)),
	)
}

//export CgoOnRtnExecOrder
func CgoOnRtnExecOrder(
	this unsafe.Pointer,
	pExecOrder *C.struct_CThostFtdcExecOrderField,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRtnExecOrder called",
		slog.Any("this", this),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnRtnExecOrder(
		(*mini.CThostFtdcExecOrderField)(unsafe.Pointer(pExecOrder)),
	)
}

//export CgoOnErrRtnExecOrderInsert
func CgoOnErrRtnExecOrderInsert(
	this unsafe.Pointer,
	pInputExecOrder *C.struct_CThostFtdcInputExecOrderField,
	pRspInfo *C.struct_CThostFtdcRspInfoField,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnErrRtnExecOrderInsert called",
		slog.Any("this", this),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnErrRtnExecOrderInsert(
		(*mini.CThostFtdcInputExecOrderField)(unsafe.Pointer(pInputExecOrder)),
		(*mini.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
	)
}

//export CgoOnErrRtnExecOrderAction
func CgoOnErrRtnExecOrderAction(
	this unsafe.Pointer,
	pExecOrderAction *C.struct_CThostFtdcExecOrderActionField,
	pRspInfo *C.struct_CThostFtdcRspInfoField,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnErrRtnExecOrderAction called",
		slog.Any("this", this),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnErrRtnExecOrderAction(
		(*mini.CThostFtdcExecOrderActionField)(unsafe.Pointer(pExecOrderAction)),
		(*mini.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
	)
}

//export CgoOnErrRtnForQuoteInsert
func CgoOnErrRtnForQuoteInsert(
	this unsafe.Pointer,
	pInputForQuote *C.struct_CThostFtdcInputForQuoteField,
	pRspInfo *C.struct_CThostFtdcRspInfoField,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnErrRtnForQuoteInsert called",
		slog.Any("this", this),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnErrRtnForQuoteInsert(
		(*mini.CThostFtdcInputForQuoteField)(unsafe.Pointer(pInputForQuote)),
		(*mini.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
	)
}

//export CgoOnRtnQuote
func CgoOnRtnQuote(
	this unsafe.Pointer,
	pQuote *C.struct_CThostFtdcQuoteField,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRtnQuote called",
		slog.Any("this", this),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnRtnQuote(
		(*mini.CThostFtdcQuoteField)(unsafe.Pointer(pQuote)),
	)
}

//export CgoOnErrRtnQuoteInsert
func CgoOnErrRtnQuoteInsert(
	this unsafe.Pointer,
	pInputQuote *C.struct_CThostFtdcInputQuoteField,
	pRspInfo *C.struct_CThostFtdcRspInfoField,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnErrRtnQuoteInsert called",
		slog.Any("this", this),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnErrRtnQuoteInsert(
		(*mini.CThostFtdcInputQuoteField)(unsafe.Pointer(pInputQuote)),
		(*mini.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
	)
}

//export CgoOnErrRtnQuoteAction
func CgoOnErrRtnQuoteAction(
	this unsafe.Pointer,
	pQuoteAction *C.struct_CThostFtdcQuoteActionField,
	pRspInfo *C.struct_CThostFtdcRspInfoField,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnErrRtnQuoteAction called",
		slog.Any("this", this),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnErrRtnQuoteAction(
		(*mini.CThostFtdcQuoteActionField)(unsafe.Pointer(pQuoteAction)),
		(*mini.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
	)
}

//export CgoOnRtnForQuoteRsp
func CgoOnRtnForQuoteRsp(
	this unsafe.Pointer,
	pForQuoteRsp *C.struct_CThostFtdcForQuoteRspField,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRtnForQuoteRsp called",
		slog.Any("this", this),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnRtnForQuoteRsp(
		(*mini.CThostFtdcForQuoteRspField)(unsafe.Pointer(pForQuoteRsp)),
	)
}

//export CgoOnErrRtnBatchOrderAction
func CgoOnErrRtnBatchOrderAction(
	this unsafe.Pointer,
	pBatchOrderAction *C.struct_CThostFtdcBatchOrderActionField,
	pRspInfo *C.struct_CThostFtdcRspInfoField,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnErrRtnBatchOrderAction called",
		slog.Any("this", this),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnErrRtnBatchOrderAction(
		(*mini.CThostFtdcBatchOrderActionField)(unsafe.Pointer(pBatchOrderAction)),
		(*mini.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
	)
}

//export CgoOnRtnOptionSelfClose
func CgoOnRtnOptionSelfClose(
	this unsafe.Pointer,
	pOptionSelfClose *C.struct_CThostFtdcOptionSelfCloseField,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRtnOptionSelfClose called",
		slog.Any("this", this),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnRtnOptionSelfClose(
		(*mini.CThostFtdcOptionSelfCloseField)(unsafe.Pointer(pOptionSelfClose)),
	)
}

//export CgoOnErrRtnOptionSelfCloseInsert
func CgoOnErrRtnOptionSelfCloseInsert(
	this unsafe.Pointer,
	pInputOptionSelfClose *C.struct_CThostFtdcInputOptionSelfCloseField,
	pRspInfo *C.struct_CThostFtdcRspInfoField,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnErrRtnOptionSelfCloseInsert called",
		slog.Any("this", this),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnErrRtnOptionSelfCloseInsert(
		(*mini.CThostFtdcInputOptionSelfCloseField)(unsafe.Pointer(pInputOptionSelfClose)),
		(*mini.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
	)
}

//export CgoOnErrRtnOptionSelfCloseAction
func CgoOnErrRtnOptionSelfCloseAction(
	this unsafe.Pointer,
	pOptionSelfCloseAction *C.struct_CThostFtdcOptionSelfCloseActionField,
	pRspInfo *C.struct_CThostFtdcRspInfoField,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnErrRtnOptionSelfCloseAction called",
		slog.Any("this", this),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnErrRtnOptionSelfCloseAction(
		(*mini.CThostFtdcOptionSelfCloseActionField)(unsafe.Pointer(pOptionSelfCloseAction)),
		(*mini.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
	)
}

//export CgoOnRtnCombAction
func CgoOnRtnCombAction(
	this unsafe.Pointer,
	pCombAction *C.struct_CThostFtdcCombActionField,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRtnCombAction called",
		slog.Any("this", this),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnRtnCombAction(
		(*mini.CThostFtdcCombActionField)(unsafe.Pointer(pCombAction)),
	)
}

//export CgoOnRspQryInstrumentOrderCommRate
func CgoOnRspQryInstrumentOrderCommRate(
	this unsafe.Pointer,
	pInstrumentOrderCommRate *C.struct_CThostFtdcInstrumentOrderCommRateField,
	pRspInfo *C.struct_CThostFtdcRspInfoField,
	nRequestID C.int,
	bIsLast C.bool,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRspQryInstrumentOrderCommRate called",
		slog.Any("this", this),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnRspQryInstrumentOrderCommRate(
		(*mini.CThostFtdcInstrumentOrderCommRateField)(unsafe.Pointer(pInstrumentOrderCommRate)),
		(*mini.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
		int(nRequestID),
		bool(bIsLast),
	)
}

//export CgoOnRtnFlowCtrlWarning
func CgoOnRtnFlowCtrlWarning(
	this unsafe.Pointer,
	pFlowCtrlWarning *C.struct_CThostFtdcFlowCtrlWarningField,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRtnFlowCtrlWarning called",
		slog.Any("this", this),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnRtnFlowCtrlWarning(
		(*mini.CThostFtdcFlowCtrlWarningField)(unsafe.Pointer(pFlowCtrlWarning)),
	)
}

//export CgoOnRspSubscribeFundChange
func CgoOnRspSubscribeFundChange(
	this unsafe.Pointer,
	pRspInfo *C.struct_CThostFtdcRspInfoField,
	nRequestID C.int,
	bIsLast C.bool,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRspSubscribeFundChange called",
		slog.Any("this", this),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnRspSubscribeFundChange(
		(*mini.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
		int(nRequestID),
		bool(bIsLast),
	)
}

//export CgoOnRspUnSubscribeFundChange
func CgoOnRspUnSubscribeFundChange(
	this unsafe.Pointer,
	pRspInfo *C.struct_CThostFtdcRspInfoField,
	nRequestID C.int,
	bIsLast C.bool,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRspUnSubscribeFundChange called",
		slog.Any("this", this),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnRspUnSubscribeFundChange(
		(*mini.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
		int(nRequestID),
		bool(bIsLast),
	)
}

//export CgoOnRtnFundChange
func CgoOnRtnFundChange(
	this unsafe.Pointer,
	pTradingAccount *C.struct_CThostFtdcTradingAccountField,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRtnFundChange called",
		slog.Any("this", this),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnRtnFundChange(
		(*mini.CThostFtdcTradingAccountField)(unsafe.Pointer(pTradingAccount)),
	)
}

//export CgoOnRspOffsetSetting
func CgoOnRspOffsetSetting(
	this unsafe.Pointer,
	pInputOffsetSetting *C.struct_CThostFtdcInputOffsetSettingField,
	pRspInfo *C.struct_CThostFtdcRspInfoField,
	nRequestID C.int,
	bIsLast C.bool,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRspOffsetSetting called",
		slog.Any("this", this),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnRspOffsetSetting(
		(*mini.CThostFtdcInputOffsetSettingField)(unsafe.Pointer(pInputOffsetSetting)),
		(*mini.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
		int(nRequestID),
		bool(bIsLast),
	)
}

//export CgoOnRspCancelOffsetSetting
func CgoOnRspCancelOffsetSetting(
	this unsafe.Pointer,
	pInputOffsetSetting *C.struct_CThostFtdcInputOffsetSettingField,
	pRspInfo *C.struct_CThostFtdcRspInfoField,
	nRequestID C.int,
	bIsLast C.bool,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRspCancelOffsetSetting called",
		slog.Any("this", this),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnRspCancelOffsetSetting(
		(*mini.CThostFtdcInputOffsetSettingField)(unsafe.Pointer(pInputOffsetSetting)),
		(*mini.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
		int(nRequestID),
		bool(bIsLast),
	)
}

//export CgoOnRtnOffsetSetting
func CgoOnRtnOffsetSetting(
	this unsafe.Pointer,
	pOffsetSetting *C.struct_CThostFtdcOffsetSettingField,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRtnOffsetSetting called",
		slog.Any("this", this),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnRtnOffsetSetting(
		(*mini.CThostFtdcOffsetSettingField)(unsafe.Pointer(pOffsetSetting)),
	)
}

//export CgoOnErrRtnCancelOffsetSetting
func CgoOnErrRtnCancelOffsetSetting(
	this unsafe.Pointer,
	pCancelOffsetSetting *C.struct_CThostFtdcCancelOffsetSettingField,
	pRspInfo *C.struct_CThostFtdcRspInfoField,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnErrRtnCancelOffsetSetting called",
		slog.Any("this", this),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnErrRtnCancelOffsetSetting(
		(*mini.CThostFtdcCancelOffsetSettingField)(unsafe.Pointer(pCancelOffsetSetting)),
		(*mini.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
	)
}

//export CgoOnRspUserPasswordUpdate
func CgoOnRspUserPasswordUpdate(
	this unsafe.Pointer,
	pUserPasswordUpdate *C.struct_CThostFtdcUserPasswordUpdateField,
	pRspInfo *C.struct_CThostFtdcRspInfoField,
	nRequestID C.int,
	bIsLast C.bool,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRspUserPasswordUpdate called",
		slog.Any("this", this),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnRspUserPasswordUpdate(
		(*mini.CThostFtdcUserPasswordUpdateField)(unsafe.Pointer(pUserPasswordUpdate)),
		(*mini.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
		int(nRequestID),
		bool(bIsLast),
	)
}
