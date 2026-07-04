package v6_7_13

/*
#cgo CFLAGS: -I. -I${SRCDIR} -I${SRCDIR}/../../../dependencies/future/v6.7.13/
#cgo LDFLAGS: -ldl

#include "spi_helper.h"
*/
import "C"
import (
	"context"
	"log/slog"
	"runtime"
	"unsafe"

	"github.com/frozenpine/ctp4go/thost/future"
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

	spiCVtablePtr.CThostFtdcTraderSpiVTable_OnRspAuthenticate = (C.OnRspAuthenticate)(
		unsafe.Pointer(C.COnRspAuthenticate))

	spiCVtablePtr.CThostFtdcTraderSpiVTable_OnRtnPrivateSeqNo = (C.OnRtnPrivateSeqNo)(
		unsafe.Pointer(C.COnRtnPrivateSeqNo))

	spiCVtablePtr.CThostFtdcTraderSpiVTable_OnRspUserLogin = (C.OnRspUserLogin)(
		unsafe.Pointer(C.COnRspUserLogin))

	spiCVtablePtr.CThostFtdcTraderSpiVTable_OnRspUserLogout = (C.OnRspUserLogout)(
		unsafe.Pointer(C.COnRspUserLogout))

	spiCVtablePtr.CThostFtdcTraderSpiVTable_OnRspUserPasswordUpdate = (C.OnRspUserPasswordUpdate)(
		unsafe.Pointer(C.COnRspUserPasswordUpdate))

	spiCVtablePtr.CThostFtdcTraderSpiVTable_OnRspTradingAccountPasswordUpdate = (C.OnRspTradingAccountPasswordUpdate)(
		unsafe.Pointer(C.COnRspTradingAccountPasswordUpdate))

	spiCVtablePtr.CThostFtdcTraderSpiVTable_OnRspUserAuthMethod = (C.OnRspUserAuthMethod)(
		unsafe.Pointer(C.COnRspUserAuthMethod))

	spiCVtablePtr.CThostFtdcTraderSpiVTable_OnRspGenUserCaptcha = (C.OnRspGenUserCaptcha)(
		unsafe.Pointer(C.COnRspGenUserCaptcha))

	spiCVtablePtr.CThostFtdcTraderSpiVTable_OnRspGenUserText = (C.OnRspGenUserText)(
		unsafe.Pointer(C.COnRspGenUserText))

	spiCVtablePtr.CThostFtdcTraderSpiVTable_OnRspOrderInsert = (C.OnRspOrderInsert)(
		unsafe.Pointer(C.COnRspOrderInsert))

	spiCVtablePtr.CThostFtdcTraderSpiVTable_OnRspParkedOrderInsert = (C.OnRspParkedOrderInsert)(
		unsafe.Pointer(C.COnRspParkedOrderInsert))

	spiCVtablePtr.CThostFtdcTraderSpiVTable_OnRspParkedOrderAction = (C.OnRspParkedOrderAction)(
		unsafe.Pointer(C.COnRspParkedOrderAction))

	spiCVtablePtr.CThostFtdcTraderSpiVTable_OnRspOrderAction = (C.OnRspOrderAction)(
		unsafe.Pointer(C.COnRspOrderAction))

	spiCVtablePtr.CThostFtdcTraderSpiVTable_OnRspQryMaxOrderVolume = (C.OnRspQryMaxOrderVolume)(
		unsafe.Pointer(C.COnRspQryMaxOrderVolume))

	spiCVtablePtr.CThostFtdcTraderSpiVTable_OnRspSettlementInfoConfirm = (C.OnRspSettlementInfoConfirm)(
		unsafe.Pointer(C.COnRspSettlementInfoConfirm))

	spiCVtablePtr.CThostFtdcTraderSpiVTable_OnRspRemoveParkedOrder = (C.OnRspRemoveParkedOrder)(
		unsafe.Pointer(C.COnRspRemoveParkedOrder))

	spiCVtablePtr.CThostFtdcTraderSpiVTable_OnRspRemoveParkedOrderAction = (C.OnRspRemoveParkedOrderAction)(
		unsafe.Pointer(C.COnRspRemoveParkedOrderAction))

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

	spiCVtablePtr.CThostFtdcTraderSpiVTable_OnRspQryUserSession = (C.OnRspQryUserSession)(
		unsafe.Pointer(C.COnRspQryUserSession))

	spiCVtablePtr.CThostFtdcTraderSpiVTable_OnRspQryExchange = (C.OnRspQryExchange)(
		unsafe.Pointer(C.COnRspQryExchange))

	spiCVtablePtr.CThostFtdcTraderSpiVTable_OnRspQryProduct = (C.OnRspQryProduct)(
		unsafe.Pointer(C.COnRspQryProduct))

	spiCVtablePtr.CThostFtdcTraderSpiVTable_OnRspQryInstrument = (C.OnRspQryInstrument)(
		unsafe.Pointer(C.COnRspQryInstrument))

	spiCVtablePtr.CThostFtdcTraderSpiVTable_OnRspQryDepthMarketData = (C.OnRspQryDepthMarketData)(
		unsafe.Pointer(C.COnRspQryDepthMarketData))

	spiCVtablePtr.CThostFtdcTraderSpiVTable_OnRspQryTraderOffer = (C.OnRspQryTraderOffer)(
		unsafe.Pointer(C.COnRspQryTraderOffer))

	spiCVtablePtr.CThostFtdcTraderSpiVTable_OnRspQrySettlementInfo = (C.OnRspQrySettlementInfo)(
		unsafe.Pointer(C.COnRspQrySettlementInfo))

	spiCVtablePtr.CThostFtdcTraderSpiVTable_OnRspQryTransferBank = (C.OnRspQryTransferBank)(
		unsafe.Pointer(C.COnRspQryTransferBank))

	spiCVtablePtr.CThostFtdcTraderSpiVTable_OnRspQryInvestorPositionDetail = (C.OnRspQryInvestorPositionDetail)(
		unsafe.Pointer(C.COnRspQryInvestorPositionDetail))

	spiCVtablePtr.CThostFtdcTraderSpiVTable_OnRspQryNotice = (C.OnRspQryNotice)(
		unsafe.Pointer(C.COnRspQryNotice))

	spiCVtablePtr.CThostFtdcTraderSpiVTable_OnRspQrySettlementInfoConfirm = (C.OnRspQrySettlementInfoConfirm)(
		unsafe.Pointer(C.COnRspQrySettlementInfoConfirm))

	spiCVtablePtr.CThostFtdcTraderSpiVTable_OnRspQryInvestorPositionCombineDetail = (C.OnRspQryInvestorPositionCombineDetail)(
		unsafe.Pointer(C.COnRspQryInvestorPositionCombineDetail))

	spiCVtablePtr.CThostFtdcTraderSpiVTable_OnRspQryCFMMCTradingAccountKey = (C.OnRspQryCFMMCTradingAccountKey)(
		unsafe.Pointer(C.COnRspQryCFMMCTradingAccountKey))

	spiCVtablePtr.CThostFtdcTraderSpiVTable_OnRspQryEWarrantOffset = (C.OnRspQryEWarrantOffset)(
		unsafe.Pointer(C.COnRspQryEWarrantOffset))

	spiCVtablePtr.CThostFtdcTraderSpiVTable_OnRspQryInvestorProductGroupMargin = (C.OnRspQryInvestorProductGroupMargin)(
		unsafe.Pointer(C.COnRspQryInvestorProductGroupMargin))

	spiCVtablePtr.CThostFtdcTraderSpiVTable_OnRspQryExchangeMarginRate = (C.OnRspQryExchangeMarginRate)(
		unsafe.Pointer(C.COnRspQryExchangeMarginRate))

	spiCVtablePtr.CThostFtdcTraderSpiVTable_OnRspQryExchangeMarginRateAdjust = (C.OnRspQryExchangeMarginRateAdjust)(
		unsafe.Pointer(C.COnRspQryExchangeMarginRateAdjust))

	spiCVtablePtr.CThostFtdcTraderSpiVTable_OnRspQryExchangeRate = (C.OnRspQryExchangeRate)(
		unsafe.Pointer(C.COnRspQryExchangeRate))

	spiCVtablePtr.CThostFtdcTraderSpiVTable_OnRspQrySecAgentACIDMap = (C.OnRspQrySecAgentACIDMap)(
		unsafe.Pointer(C.COnRspQrySecAgentACIDMap))

	spiCVtablePtr.CThostFtdcTraderSpiVTable_OnRspQryProductExchRate = (C.OnRspQryProductExchRate)(
		unsafe.Pointer(C.COnRspQryProductExchRate))

	spiCVtablePtr.CThostFtdcTraderSpiVTable_OnRspQryProductGroup = (C.OnRspQryProductGroup)(
		unsafe.Pointer(C.COnRspQryProductGroup))

	spiCVtablePtr.CThostFtdcTraderSpiVTable_OnRspQryMMInstrumentCommissionRate = (C.OnRspQryMMInstrumentCommissionRate)(
		unsafe.Pointer(C.COnRspQryMMInstrumentCommissionRate))

	spiCVtablePtr.CThostFtdcTraderSpiVTable_OnRspQryMMOptionInstrCommRate = (C.OnRspQryMMOptionInstrCommRate)(
		unsafe.Pointer(C.COnRspQryMMOptionInstrCommRate))

	spiCVtablePtr.CThostFtdcTraderSpiVTable_OnRspQryInstrumentOrderCommRate = (C.OnRspQryInstrumentOrderCommRate)(
		unsafe.Pointer(C.COnRspQryInstrumentOrderCommRate))

	spiCVtablePtr.CThostFtdcTraderSpiVTable_OnRspQrySecAgentTradingAccount = (C.OnRspQrySecAgentTradingAccount)(
		unsafe.Pointer(C.COnRspQrySecAgentTradingAccount))

	spiCVtablePtr.CThostFtdcTraderSpiVTable_OnRspQrySecAgentCheckMode = (C.OnRspQrySecAgentCheckMode)(
		unsafe.Pointer(C.COnRspQrySecAgentCheckMode))

	spiCVtablePtr.CThostFtdcTraderSpiVTable_OnRspQrySecAgentTradeInfo = (C.OnRspQrySecAgentTradeInfo)(
		unsafe.Pointer(C.COnRspQrySecAgentTradeInfo))

	spiCVtablePtr.CThostFtdcTraderSpiVTable_OnRspQryOptionInstrTradeCost = (C.OnRspQryOptionInstrTradeCost)(
		unsafe.Pointer(C.COnRspQryOptionInstrTradeCost))

	spiCVtablePtr.CThostFtdcTraderSpiVTable_OnRspQryOptionInstrCommRate = (C.OnRspQryOptionInstrCommRate)(
		unsafe.Pointer(C.COnRspQryOptionInstrCommRate))

	spiCVtablePtr.CThostFtdcTraderSpiVTable_OnRspQryExecOrder = (C.OnRspQryExecOrder)(
		unsafe.Pointer(C.COnRspQryExecOrder))

	spiCVtablePtr.CThostFtdcTraderSpiVTable_OnRspQryForQuote = (C.OnRspQryForQuote)(
		unsafe.Pointer(C.COnRspQryForQuote))

	spiCVtablePtr.CThostFtdcTraderSpiVTable_OnRspQryQuote = (C.OnRspQryQuote)(
		unsafe.Pointer(C.COnRspQryQuote))

	spiCVtablePtr.CThostFtdcTraderSpiVTable_OnRspQryOptionSelfClose = (C.OnRspQryOptionSelfClose)(
		unsafe.Pointer(C.COnRspQryOptionSelfClose))

	spiCVtablePtr.CThostFtdcTraderSpiVTable_OnRspQryInvestUnit = (C.OnRspQryInvestUnit)(
		unsafe.Pointer(C.COnRspQryInvestUnit))

	spiCVtablePtr.CThostFtdcTraderSpiVTable_OnRspQryCombInstrumentGuard = (C.OnRspQryCombInstrumentGuard)(
		unsafe.Pointer(C.COnRspQryCombInstrumentGuard))

	spiCVtablePtr.CThostFtdcTraderSpiVTable_OnRspQryCombAction = (C.OnRspQryCombAction)(
		unsafe.Pointer(C.COnRspQryCombAction))

	spiCVtablePtr.CThostFtdcTraderSpiVTable_OnRspQryTransferSerial = (C.OnRspQryTransferSerial)(
		unsafe.Pointer(C.COnRspQryTransferSerial))

	spiCVtablePtr.CThostFtdcTraderSpiVTable_OnRspQryAccountregister = (C.OnRspQryAccountregister)(
		unsafe.Pointer(C.COnRspQryAccountregister))

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

	spiCVtablePtr.CThostFtdcTraderSpiVTable_OnRtnBulletin = (C.OnRtnBulletin)(
		unsafe.Pointer(C.COnRtnBulletin))

	spiCVtablePtr.CThostFtdcTraderSpiVTable_OnRtnTradingNotice = (C.OnRtnTradingNotice)(
		unsafe.Pointer(C.COnRtnTradingNotice))

	spiCVtablePtr.CThostFtdcTraderSpiVTable_OnRtnErrorConditionalOrder = (C.OnRtnErrorConditionalOrder)(
		unsafe.Pointer(C.COnRtnErrorConditionalOrder))

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

	spiCVtablePtr.CThostFtdcTraderSpiVTable_OnRtnCFMMCTradingAccountToken = (C.OnRtnCFMMCTradingAccountToken)(
		unsafe.Pointer(C.COnRtnCFMMCTradingAccountToken))

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

	spiCVtablePtr.CThostFtdcTraderSpiVTable_OnErrRtnCombActionInsert = (C.OnErrRtnCombActionInsert)(
		unsafe.Pointer(C.COnErrRtnCombActionInsert))

	spiCVtablePtr.CThostFtdcTraderSpiVTable_OnRspQryContractBank = (C.OnRspQryContractBank)(
		unsafe.Pointer(C.COnRspQryContractBank))

	spiCVtablePtr.CThostFtdcTraderSpiVTable_OnRspQryParkedOrder = (C.OnRspQryParkedOrder)(
		unsafe.Pointer(C.COnRspQryParkedOrder))

	spiCVtablePtr.CThostFtdcTraderSpiVTable_OnRspQryParkedOrderAction = (C.OnRspQryParkedOrderAction)(
		unsafe.Pointer(C.COnRspQryParkedOrderAction))

	spiCVtablePtr.CThostFtdcTraderSpiVTable_OnRspQryTradingNotice = (C.OnRspQryTradingNotice)(
		unsafe.Pointer(C.COnRspQryTradingNotice))

	spiCVtablePtr.CThostFtdcTraderSpiVTable_OnRspQryBrokerTradingParams = (C.OnRspQryBrokerTradingParams)(
		unsafe.Pointer(C.COnRspQryBrokerTradingParams))

	spiCVtablePtr.CThostFtdcTraderSpiVTable_OnRspQryBrokerTradingAlgos = (C.OnRspQryBrokerTradingAlgos)(
		unsafe.Pointer(C.COnRspQryBrokerTradingAlgos))

	spiCVtablePtr.CThostFtdcTraderSpiVTable_OnRspQueryCFMMCTradingAccountToken = (C.OnRspQueryCFMMCTradingAccountToken)(
		unsafe.Pointer(C.COnRspQueryCFMMCTradingAccountToken))

	spiCVtablePtr.CThostFtdcTraderSpiVTable_OnRtnFromBankToFutureByBank = (C.OnRtnFromBankToFutureByBank)(
		unsafe.Pointer(C.COnRtnFromBankToFutureByBank))

	spiCVtablePtr.CThostFtdcTraderSpiVTable_OnRtnFromFutureToBankByBank = (C.OnRtnFromFutureToBankByBank)(
		unsafe.Pointer(C.COnRtnFromFutureToBankByBank))

	spiCVtablePtr.CThostFtdcTraderSpiVTable_OnRtnRepealFromBankToFutureByBank = (C.OnRtnRepealFromBankToFutureByBank)(
		unsafe.Pointer(C.COnRtnRepealFromBankToFutureByBank))

	spiCVtablePtr.CThostFtdcTraderSpiVTable_OnRtnRepealFromFutureToBankByBank = (C.OnRtnRepealFromFutureToBankByBank)(
		unsafe.Pointer(C.COnRtnRepealFromFutureToBankByBank))

	spiCVtablePtr.CThostFtdcTraderSpiVTable_OnRtnFromBankToFutureByFuture = (C.OnRtnFromBankToFutureByFuture)(
		unsafe.Pointer(C.COnRtnFromBankToFutureByFuture))

	spiCVtablePtr.CThostFtdcTraderSpiVTable_OnRtnFromFutureToBankByFuture = (C.OnRtnFromFutureToBankByFuture)(
		unsafe.Pointer(C.COnRtnFromFutureToBankByFuture))

	spiCVtablePtr.CThostFtdcTraderSpiVTable_OnRtnRepealFromBankToFutureByFutureManual = (C.OnRtnRepealFromBankToFutureByFutureManual)(
		unsafe.Pointer(C.COnRtnRepealFromBankToFutureByFutureManual))

	spiCVtablePtr.CThostFtdcTraderSpiVTable_OnRtnRepealFromFutureToBankByFutureManual = (C.OnRtnRepealFromFutureToBankByFutureManual)(
		unsafe.Pointer(C.COnRtnRepealFromFutureToBankByFutureManual))

	spiCVtablePtr.CThostFtdcTraderSpiVTable_OnRtnQueryBankBalanceByFuture = (C.OnRtnQueryBankBalanceByFuture)(
		unsafe.Pointer(C.COnRtnQueryBankBalanceByFuture))

	spiCVtablePtr.CThostFtdcTraderSpiVTable_OnErrRtnBankToFutureByFuture = (C.OnErrRtnBankToFutureByFuture)(
		unsafe.Pointer(C.COnErrRtnBankToFutureByFuture))

	spiCVtablePtr.CThostFtdcTraderSpiVTable_OnErrRtnFutureToBankByFuture = (C.OnErrRtnFutureToBankByFuture)(
		unsafe.Pointer(C.COnErrRtnFutureToBankByFuture))

	spiCVtablePtr.CThostFtdcTraderSpiVTable_OnErrRtnRepealBankToFutureByFutureManual = (C.OnErrRtnRepealBankToFutureByFutureManual)(
		unsafe.Pointer(C.COnErrRtnRepealBankToFutureByFutureManual))

	spiCVtablePtr.CThostFtdcTraderSpiVTable_OnErrRtnRepealFutureToBankByFutureManual = (C.OnErrRtnRepealFutureToBankByFutureManual)(
		unsafe.Pointer(C.COnErrRtnRepealFutureToBankByFutureManual))

	spiCVtablePtr.CThostFtdcTraderSpiVTable_OnErrRtnQueryBankBalanceByFuture = (C.OnErrRtnQueryBankBalanceByFuture)(
		unsafe.Pointer(C.COnErrRtnQueryBankBalanceByFuture))

	spiCVtablePtr.CThostFtdcTraderSpiVTable_OnRtnRepealFromBankToFutureByFuture = (C.OnRtnRepealFromBankToFutureByFuture)(
		unsafe.Pointer(C.COnRtnRepealFromBankToFutureByFuture))

	spiCVtablePtr.CThostFtdcTraderSpiVTable_OnRtnRepealFromFutureToBankByFuture = (C.OnRtnRepealFromFutureToBankByFuture)(
		unsafe.Pointer(C.COnRtnRepealFromFutureToBankByFuture))

	spiCVtablePtr.CThostFtdcTraderSpiVTable_OnRspFromBankToFutureByFuture = (C.OnRspFromBankToFutureByFuture)(
		unsafe.Pointer(C.COnRspFromBankToFutureByFuture))

	spiCVtablePtr.CThostFtdcTraderSpiVTable_OnRspFromFutureToBankByFuture = (C.OnRspFromFutureToBankByFuture)(
		unsafe.Pointer(C.COnRspFromFutureToBankByFuture))

	spiCVtablePtr.CThostFtdcTraderSpiVTable_OnRspQueryBankAccountMoneyByFuture = (C.OnRspQueryBankAccountMoneyByFuture)(
		unsafe.Pointer(C.COnRspQueryBankAccountMoneyByFuture))

	spiCVtablePtr.CThostFtdcTraderSpiVTable_OnRtnOpenAccountByBank = (C.OnRtnOpenAccountByBank)(
		unsafe.Pointer(C.COnRtnOpenAccountByBank))

	spiCVtablePtr.CThostFtdcTraderSpiVTable_OnRtnCancelAccountByBank = (C.OnRtnCancelAccountByBank)(
		unsafe.Pointer(C.COnRtnCancelAccountByBank))

	spiCVtablePtr.CThostFtdcTraderSpiVTable_OnRtnChangeAccountByBank = (C.OnRtnChangeAccountByBank)(
		unsafe.Pointer(C.COnRtnChangeAccountByBank))

	spiCVtablePtr.CThostFtdcTraderSpiVTable_OnRspQryClassifiedInstrument = (C.OnRspQryClassifiedInstrument)(
		unsafe.Pointer(C.COnRspQryClassifiedInstrument))

	spiCVtablePtr.CThostFtdcTraderSpiVTable_OnRspQryCombPromotionParam = (C.OnRspQryCombPromotionParam)(
		unsafe.Pointer(C.COnRspQryCombPromotionParam))

	spiCVtablePtr.CThostFtdcTraderSpiVTable_OnRspQryRiskSettleInvstPosition = (C.OnRspQryRiskSettleInvstPosition)(
		unsafe.Pointer(C.COnRspQryRiskSettleInvstPosition))

	spiCVtablePtr.CThostFtdcTraderSpiVTable_OnRspQryRiskSettleProductStatus = (C.OnRspQryRiskSettleProductStatus)(
		unsafe.Pointer(C.COnRspQryRiskSettleProductStatus))

	spiCVtablePtr.CThostFtdcTraderSpiVTable_OnRspQrySPBMFutureParameter = (C.OnRspQrySPBMFutureParameter)(
		unsafe.Pointer(C.COnRspQrySPBMFutureParameter))

	spiCVtablePtr.CThostFtdcTraderSpiVTable_OnRspQrySPBMOptionParameter = (C.OnRspQrySPBMOptionParameter)(
		unsafe.Pointer(C.COnRspQrySPBMOptionParameter))

	spiCVtablePtr.CThostFtdcTraderSpiVTable_OnRspQrySPBMIntraParameter = (C.OnRspQrySPBMIntraParameter)(
		unsafe.Pointer(C.COnRspQrySPBMIntraParameter))

	spiCVtablePtr.CThostFtdcTraderSpiVTable_OnRspQrySPBMInterParameter = (C.OnRspQrySPBMInterParameter)(
		unsafe.Pointer(C.COnRspQrySPBMInterParameter))

	spiCVtablePtr.CThostFtdcTraderSpiVTable_OnRspQrySPBMPortfDefinition = (C.OnRspQrySPBMPortfDefinition)(
		unsafe.Pointer(C.COnRspQrySPBMPortfDefinition))

	spiCVtablePtr.CThostFtdcTraderSpiVTable_OnRspQrySPBMInvestorPortfDef = (C.OnRspQrySPBMInvestorPortfDef)(
		unsafe.Pointer(C.COnRspQrySPBMInvestorPortfDef))

	spiCVtablePtr.CThostFtdcTraderSpiVTable_OnRspQryInvestorPortfMarginRatio = (C.OnRspQryInvestorPortfMarginRatio)(
		unsafe.Pointer(C.COnRspQryInvestorPortfMarginRatio))

	spiCVtablePtr.CThostFtdcTraderSpiVTable_OnRspQryInvestorProdSPBMDetail = (C.OnRspQryInvestorProdSPBMDetail)(
		unsafe.Pointer(C.COnRspQryInvestorProdSPBMDetail))

	spiCVtablePtr.CThostFtdcTraderSpiVTable_OnRspQryInvestorCommoditySPMMMargin = (C.OnRspQryInvestorCommoditySPMMMargin)(
		unsafe.Pointer(C.COnRspQryInvestorCommoditySPMMMargin))

	spiCVtablePtr.CThostFtdcTraderSpiVTable_OnRspQryInvestorCommodityGroupSPMMMargin = (C.OnRspQryInvestorCommodityGroupSPMMMargin)(
		unsafe.Pointer(C.COnRspQryInvestorCommodityGroupSPMMMargin))

	spiCVtablePtr.CThostFtdcTraderSpiVTable_OnRspQrySPMMInstParam = (C.OnRspQrySPMMInstParam)(
		unsafe.Pointer(C.COnRspQrySPMMInstParam))

	spiCVtablePtr.CThostFtdcTraderSpiVTable_OnRspQrySPMMProductParam = (C.OnRspQrySPMMProductParam)(
		unsafe.Pointer(C.COnRspQrySPMMProductParam))

	spiCVtablePtr.CThostFtdcTraderSpiVTable_OnRspQrySPBMAddOnInterParameter = (C.OnRspQrySPBMAddOnInterParameter)(
		unsafe.Pointer(C.COnRspQrySPBMAddOnInterParameter))

	spiCVtablePtr.CThostFtdcTraderSpiVTable_OnRspQryRCAMSCombProductInfo = (C.OnRspQryRCAMSCombProductInfo)(
		unsafe.Pointer(C.COnRspQryRCAMSCombProductInfo))

	spiCVtablePtr.CThostFtdcTraderSpiVTable_OnRspQryRCAMSInstrParameter = (C.OnRspQryRCAMSInstrParameter)(
		unsafe.Pointer(C.COnRspQryRCAMSInstrParameter))

	spiCVtablePtr.CThostFtdcTraderSpiVTable_OnRspQryRCAMSIntraParameter = (C.OnRspQryRCAMSIntraParameter)(
		unsafe.Pointer(C.COnRspQryRCAMSIntraParameter))

	spiCVtablePtr.CThostFtdcTraderSpiVTable_OnRspQryRCAMSInterParameter = (C.OnRspQryRCAMSInterParameter)(
		unsafe.Pointer(C.COnRspQryRCAMSInterParameter))

	spiCVtablePtr.CThostFtdcTraderSpiVTable_OnRspQryRCAMSShortOptAdjustParam = (C.OnRspQryRCAMSShortOptAdjustParam)(
		unsafe.Pointer(C.COnRspQryRCAMSShortOptAdjustParam))

	spiCVtablePtr.CThostFtdcTraderSpiVTable_OnRspQryRCAMSInvestorCombPosition = (C.OnRspQryRCAMSInvestorCombPosition)(
		unsafe.Pointer(C.COnRspQryRCAMSInvestorCombPosition))

	spiCVtablePtr.CThostFtdcTraderSpiVTable_OnRspQryInvestorProdRCAMSMargin = (C.OnRspQryInvestorProdRCAMSMargin)(
		unsafe.Pointer(C.COnRspQryInvestorProdRCAMSMargin))

	spiCVtablePtr.CThostFtdcTraderSpiVTable_OnRspQryRULEInstrParameter = (C.OnRspQryRULEInstrParameter)(
		unsafe.Pointer(C.COnRspQryRULEInstrParameter))

	spiCVtablePtr.CThostFtdcTraderSpiVTable_OnRspQryRULEIntraParameter = (C.OnRspQryRULEIntraParameter)(
		unsafe.Pointer(C.COnRspQryRULEIntraParameter))

	spiCVtablePtr.CThostFtdcTraderSpiVTable_OnRspQryRULEInterParameter = (C.OnRspQryRULEInterParameter)(
		unsafe.Pointer(C.COnRspQryRULEInterParameter))

	spiCVtablePtr.CThostFtdcTraderSpiVTable_OnRspQryInvestorProdRULEMargin = (C.OnRspQryInvestorProdRULEMargin)(
		unsafe.Pointer(C.COnRspQryInvestorProdRULEMargin))

	spiCVtablePtr.CThostFtdcTraderSpiVTable_OnRspQryInvestorPortfSetting = (C.OnRspQryInvestorPortfSetting)(
		unsafe.Pointer(C.COnRspQryInvestorPortfSetting))

	spiCVtablePtr.CThostFtdcTraderSpiVTable_OnRspQryInvestorInfoCommRec = (C.OnRspQryInvestorInfoCommRec)(
		unsafe.Pointer(C.COnRspQryInvestorInfoCommRec))

	spiCVtablePtr.CThostFtdcTraderSpiVTable_OnRspQryCombLeg = (C.OnRspQryCombLeg)(
		unsafe.Pointer(C.COnRspQryCombLeg))

	spiCVtablePtr.CThostFtdcTraderSpiVTable_OnRspOffsetSetting = (C.OnRspOffsetSetting)(
		unsafe.Pointer(C.COnRspOffsetSetting))

	spiCVtablePtr.CThostFtdcTraderSpiVTable_OnRspCancelOffsetSetting = (C.OnRspCancelOffsetSetting)(
		unsafe.Pointer(C.COnRspCancelOffsetSetting))

	spiCVtablePtr.CThostFtdcTraderSpiVTable_OnRtnOffsetSetting = (C.OnRtnOffsetSetting)(
		unsafe.Pointer(C.COnRtnOffsetSetting))

	spiCVtablePtr.CThostFtdcTraderSpiVTable_OnErrRtnOffsetSetting = (C.OnErrRtnOffsetSetting)(
		unsafe.Pointer(C.COnErrRtnOffsetSetting))

	spiCVtablePtr.CThostFtdcTraderSpiVTable_OnErrRtnCancelOffsetSetting = (C.OnErrRtnCancelOffsetSetting)(
		unsafe.Pointer(C.COnErrRtnCancelOffsetSetting))

	spiCVtablePtr.CThostFtdcTraderSpiVTable_OnRspQryOffsetSetting = (C.OnRspQryOffsetSetting)(
		unsafe.Pointer(C.COnRspQryOffsetSetting))

	spiCVtablePtr.CThostFtdcTraderSpiVTable_OnRspGenSMSCode = (C.OnRspGenSMSCode)(
		unsafe.Pointer(C.COnRspGenSMSCode))

	spiCVtablePtr.CThostFtdcTraderSpiVTable_OnRspSpdApply = (C.OnRspSpdApply)(
		unsafe.Pointer(C.COnRspSpdApply))

	spiCVtablePtr.CThostFtdcTraderSpiVTable_OnRspSpdApplyAction = (C.OnRspSpdApplyAction)(
		unsafe.Pointer(C.COnRspSpdApplyAction))

	spiCVtablePtr.CThostFtdcTraderSpiVTable_OnRspQrySpdApply = (C.OnRspQrySpdApply)(
		unsafe.Pointer(C.COnRspQrySpdApply))

	spiCVtablePtr.CThostFtdcTraderSpiVTable_OnRtnSpdApply = (C.OnRtnSpdApply)(
		unsafe.Pointer(C.COnRtnSpdApply))

	spiCVtablePtr.CThostFtdcTraderSpiVTable_OnErrRtnSpdApply = (C.OnErrRtnSpdApply)(
		unsafe.Pointer(C.COnErrRtnSpdApply))

	spiCVtablePtr.CThostFtdcTraderSpiVTable_OnErrRtnSpdApplyAction = (C.OnErrRtnSpdApplyAction)(
		unsafe.Pointer(C.COnErrRtnSpdApplyAction))

	spiCVtablePtr.CThostFtdcTraderSpiVTable_OnRspHedgeCfm = (C.OnRspHedgeCfm)(
		unsafe.Pointer(C.COnRspHedgeCfm))

	spiCVtablePtr.CThostFtdcTraderSpiVTable_OnRspHedgeCfmAction = (C.OnRspHedgeCfmAction)(
		unsafe.Pointer(C.COnRspHedgeCfmAction))

	spiCVtablePtr.CThostFtdcTraderSpiVTable_OnRspQryHedgeCfm = (C.OnRspQryHedgeCfm)(
		unsafe.Pointer(C.COnRspQryHedgeCfm))

	spiCVtablePtr.CThostFtdcTraderSpiVTable_OnRtnHedgeCfm = (C.OnRtnHedgeCfm)(
		unsafe.Pointer(C.COnRtnHedgeCfm))

	spiCVtablePtr.CThostFtdcTraderSpiVTable_OnErrRtnHedgeCfm = (C.OnErrRtnHedgeCfm)(
		unsafe.Pointer(C.COnErrRtnHedgeCfm))

	spiCVtablePtr.CThostFtdcTraderSpiVTable_OnErrRtnHedgeCfmAction = (C.OnErrRtnHedgeCfmAction)(
		unsafe.Pointer(C.COnErrRtnHedgeCfmAction))

}

type ThostFtdcTraderSpi struct {
	runtime.Pinner
	callback future.TraderSpi
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
		(*future.CThostFtdcRspAuthenticateField)(unsafe.Pointer(pRspAuthenticateField)),
		(*future.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
		int(nRequestID),
		bool(bIsLast),
	)
}

//export CgoOnRtnPrivateSeqNo
func CgoOnRtnPrivateSeqNo(
	this unsafe.Pointer,
	nSeqNo C.int,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRtnPrivateSeqNo called",
		slog.Any("this", this),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnRtnPrivateSeqNo(
		int(nSeqNo),
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
		(*future.CThostFtdcRspUserLoginField)(unsafe.Pointer(pRspUserLogin)),
		(*future.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
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
		(*future.CThostFtdcUserLogoutField)(unsafe.Pointer(pUserLogout)),
		(*future.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
		int(nRequestID),
		bool(bIsLast),
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
		(*future.CThostFtdcUserPasswordUpdateField)(unsafe.Pointer(pUserPasswordUpdate)),
		(*future.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
		int(nRequestID),
		bool(bIsLast),
	)
}

//export CgoOnRspTradingAccountPasswordUpdate
func CgoOnRspTradingAccountPasswordUpdate(
	this unsafe.Pointer,
	pTradingAccountPasswordUpdate *C.struct_CThostFtdcTradingAccountPasswordUpdateField,
	pRspInfo *C.struct_CThostFtdcRspInfoField,
	nRequestID C.int,
	bIsLast C.bool,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRspTradingAccountPasswordUpdate called",
		slog.Any("this", this),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnRspTradingAccountPasswordUpdate(
		(*future.CThostFtdcTradingAccountPasswordUpdateField)(unsafe.Pointer(pTradingAccountPasswordUpdate)),
		(*future.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
		int(nRequestID),
		bool(bIsLast),
	)
}

//export CgoOnRspUserAuthMethod
func CgoOnRspUserAuthMethod(
	this unsafe.Pointer,
	pRspUserAuthMethod *C.struct_CThostFtdcRspUserAuthMethodField,
	pRspInfo *C.struct_CThostFtdcRspInfoField,
	nRequestID C.int,
	bIsLast C.bool,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRspUserAuthMethod called",
		slog.Any("this", this),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnRspUserAuthMethod(
		(*future.CThostFtdcRspUserAuthMethodField)(unsafe.Pointer(pRspUserAuthMethod)),
		(*future.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
		int(nRequestID),
		bool(bIsLast),
	)
}

//export CgoOnRspGenUserCaptcha
func CgoOnRspGenUserCaptcha(
	this unsafe.Pointer,
	pRspGenUserCaptcha *C.struct_CThostFtdcRspGenUserCaptchaField,
	pRspInfo *C.struct_CThostFtdcRspInfoField,
	nRequestID C.int,
	bIsLast C.bool,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRspGenUserCaptcha called",
		slog.Any("this", this),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnRspGenUserCaptcha(
		(*future.CThostFtdcRspGenUserCaptchaField)(unsafe.Pointer(pRspGenUserCaptcha)),
		(*future.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
		int(nRequestID),
		bool(bIsLast),
	)
}

//export CgoOnRspGenUserText
func CgoOnRspGenUserText(
	this unsafe.Pointer,
	pRspGenUserText *C.struct_CThostFtdcRspGenUserTextField,
	pRspInfo *C.struct_CThostFtdcRspInfoField,
	nRequestID C.int,
	bIsLast C.bool,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRspGenUserText called",
		slog.Any("this", this),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnRspGenUserText(
		(*future.CThostFtdcRspGenUserTextField)(unsafe.Pointer(pRspGenUserText)),
		(*future.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
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
		(*future.CThostFtdcInputOrderField)(unsafe.Pointer(pInputOrder)),
		(*future.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
		int(nRequestID),
		bool(bIsLast),
	)
}

//export CgoOnRspParkedOrderInsert
func CgoOnRspParkedOrderInsert(
	this unsafe.Pointer,
	pParkedOrder *C.struct_CThostFtdcParkedOrderField,
	pRspInfo *C.struct_CThostFtdcRspInfoField,
	nRequestID C.int,
	bIsLast C.bool,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRspParkedOrderInsert called",
		slog.Any("this", this),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnRspParkedOrderInsert(
		(*future.CThostFtdcParkedOrderField)(unsafe.Pointer(pParkedOrder)),
		(*future.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
		int(nRequestID),
		bool(bIsLast),
	)
}

//export CgoOnRspParkedOrderAction
func CgoOnRspParkedOrderAction(
	this unsafe.Pointer,
	pParkedOrderAction *C.struct_CThostFtdcParkedOrderActionField,
	pRspInfo *C.struct_CThostFtdcRspInfoField,
	nRequestID C.int,
	bIsLast C.bool,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRspParkedOrderAction called",
		slog.Any("this", this),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnRspParkedOrderAction(
		(*future.CThostFtdcParkedOrderActionField)(unsafe.Pointer(pParkedOrderAction)),
		(*future.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
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
		(*future.CThostFtdcInputOrderActionField)(unsafe.Pointer(pInputOrderAction)),
		(*future.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
		int(nRequestID),
		bool(bIsLast),
	)
}

//export CgoOnRspQryMaxOrderVolume
func CgoOnRspQryMaxOrderVolume(
	this unsafe.Pointer,
	pQryMaxOrderVolume *C.struct_CThostFtdcQryMaxOrderVolumeField,
	pRspInfo *C.struct_CThostFtdcRspInfoField,
	nRequestID C.int,
	bIsLast C.bool,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRspQryMaxOrderVolume called",
		slog.Any("this", this),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnRspQryMaxOrderVolume(
		(*future.CThostFtdcQryMaxOrderVolumeField)(unsafe.Pointer(pQryMaxOrderVolume)),
		(*future.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
		int(nRequestID),
		bool(bIsLast),
	)
}

//export CgoOnRspSettlementInfoConfirm
func CgoOnRspSettlementInfoConfirm(
	this unsafe.Pointer,
	pSettlementInfoConfirm *C.struct_CThostFtdcSettlementInfoConfirmField,
	pRspInfo *C.struct_CThostFtdcRspInfoField,
	nRequestID C.int,
	bIsLast C.bool,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRspSettlementInfoConfirm called",
		slog.Any("this", this),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnRspSettlementInfoConfirm(
		(*future.CThostFtdcSettlementInfoConfirmField)(unsafe.Pointer(pSettlementInfoConfirm)),
		(*future.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
		int(nRequestID),
		bool(bIsLast),
	)
}

//export CgoOnRspRemoveParkedOrder
func CgoOnRspRemoveParkedOrder(
	this unsafe.Pointer,
	pRemoveParkedOrder *C.struct_CThostFtdcRemoveParkedOrderField,
	pRspInfo *C.struct_CThostFtdcRspInfoField,
	nRequestID C.int,
	bIsLast C.bool,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRspRemoveParkedOrder called",
		slog.Any("this", this),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnRspRemoveParkedOrder(
		(*future.CThostFtdcRemoveParkedOrderField)(unsafe.Pointer(pRemoveParkedOrder)),
		(*future.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
		int(nRequestID),
		bool(bIsLast),
	)
}

//export CgoOnRspRemoveParkedOrderAction
func CgoOnRspRemoveParkedOrderAction(
	this unsafe.Pointer,
	pRemoveParkedOrderAction *C.struct_CThostFtdcRemoveParkedOrderActionField,
	pRspInfo *C.struct_CThostFtdcRspInfoField,
	nRequestID C.int,
	bIsLast C.bool,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRspRemoveParkedOrderAction called",
		slog.Any("this", this),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnRspRemoveParkedOrderAction(
		(*future.CThostFtdcRemoveParkedOrderActionField)(unsafe.Pointer(pRemoveParkedOrderAction)),
		(*future.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
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
		(*future.CThostFtdcInputExecOrderField)(unsafe.Pointer(pInputExecOrder)),
		(*future.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
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
		(*future.CThostFtdcInputExecOrderActionField)(unsafe.Pointer(pInputExecOrderAction)),
		(*future.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
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
		(*future.CThostFtdcInputForQuoteField)(unsafe.Pointer(pInputForQuote)),
		(*future.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
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
		(*future.CThostFtdcInputQuoteField)(unsafe.Pointer(pInputQuote)),
		(*future.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
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
		(*future.CThostFtdcInputQuoteActionField)(unsafe.Pointer(pInputQuoteAction)),
		(*future.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
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
		(*future.CThostFtdcInputBatchOrderActionField)(unsafe.Pointer(pInputBatchOrderAction)),
		(*future.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
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
		(*future.CThostFtdcInputOptionSelfCloseField)(unsafe.Pointer(pInputOptionSelfClose)),
		(*future.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
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
		(*future.CThostFtdcInputOptionSelfCloseActionField)(unsafe.Pointer(pInputOptionSelfCloseAction)),
		(*future.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
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
		(*future.CThostFtdcInputCombActionField)(unsafe.Pointer(pInputCombAction)),
		(*future.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
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
		(*future.CThostFtdcOrderField)(unsafe.Pointer(pOrder)),
		(*future.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
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
		(*future.CThostFtdcTradeField)(unsafe.Pointer(pTrade)),
		(*future.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
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
		(*future.CThostFtdcInvestorPositionField)(unsafe.Pointer(pInvestorPosition)),
		(*future.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
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
		(*future.CThostFtdcTradingAccountField)(unsafe.Pointer(pTradingAccount)),
		(*future.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
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
		(*future.CThostFtdcInvestorField)(unsafe.Pointer(pInvestor)),
		(*future.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
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
		(*future.CThostFtdcTradingCodeField)(unsafe.Pointer(pTradingCode)),
		(*future.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
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
		(*future.CThostFtdcInstrumentMarginRateField)(unsafe.Pointer(pInstrumentMarginRate)),
		(*future.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
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
		(*future.CThostFtdcInstrumentCommissionRateField)(unsafe.Pointer(pInstrumentCommissionRate)),
		(*future.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
		int(nRequestID),
		bool(bIsLast),
	)
}

//export CgoOnRspQryUserSession
func CgoOnRspQryUserSession(
	this unsafe.Pointer,
	pUserSession *C.struct_CThostFtdcUserSessionField,
	pRspInfo *C.struct_CThostFtdcRspInfoField,
	nRequestID C.int,
	bIsLast C.bool,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRspQryUserSession called",
		slog.Any("this", this),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnRspQryUserSession(
		(*future.CThostFtdcUserSessionField)(unsafe.Pointer(pUserSession)),
		(*future.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
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
		(*future.CThostFtdcExchangeField)(unsafe.Pointer(pExchange)),
		(*future.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
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
		(*future.CThostFtdcProductField)(unsafe.Pointer(pProduct)),
		(*future.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
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
		(*future.CThostFtdcInstrumentField)(unsafe.Pointer(pInstrument)),
		(*future.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
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
		(*future.CThostFtdcDepthMarketDataField)(unsafe.Pointer(pDepthMarketData)),
		(*future.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
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
		(*future.CThostFtdcTraderOfferField)(unsafe.Pointer(pTraderOffer)),
		(*future.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
		int(nRequestID),
		bool(bIsLast),
	)
}

//export CgoOnRspQrySettlementInfo
func CgoOnRspQrySettlementInfo(
	this unsafe.Pointer,
	pSettlementInfo *C.struct_CThostFtdcSettlementInfoField,
	pRspInfo *C.struct_CThostFtdcRspInfoField,
	nRequestID C.int,
	bIsLast C.bool,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRspQrySettlementInfo called",
		slog.Any("this", this),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnRspQrySettlementInfo(
		(*future.CThostFtdcSettlementInfoField)(unsafe.Pointer(pSettlementInfo)),
		(*future.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
		int(nRequestID),
		bool(bIsLast),
	)
}

//export CgoOnRspQryTransferBank
func CgoOnRspQryTransferBank(
	this unsafe.Pointer,
	pTransferBank *C.struct_CThostFtdcTransferBankField,
	pRspInfo *C.struct_CThostFtdcRspInfoField,
	nRequestID C.int,
	bIsLast C.bool,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRspQryTransferBank called",
		slog.Any("this", this),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnRspQryTransferBank(
		(*future.CThostFtdcTransferBankField)(unsafe.Pointer(pTransferBank)),
		(*future.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
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
		(*future.CThostFtdcInvestorPositionDetailField)(unsafe.Pointer(pInvestorPositionDetail)),
		(*future.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
		int(nRequestID),
		bool(bIsLast),
	)
}

//export CgoOnRspQryNotice
func CgoOnRspQryNotice(
	this unsafe.Pointer,
	pNotice *C.struct_CThostFtdcNoticeField,
	pRspInfo *C.struct_CThostFtdcRspInfoField,
	nRequestID C.int,
	bIsLast C.bool,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRspQryNotice called",
		slog.Any("this", this),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnRspQryNotice(
		(*future.CThostFtdcNoticeField)(unsafe.Pointer(pNotice)),
		(*future.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
		int(nRequestID),
		bool(bIsLast),
	)
}

//export CgoOnRspQrySettlementInfoConfirm
func CgoOnRspQrySettlementInfoConfirm(
	this unsafe.Pointer,
	pSettlementInfoConfirm *C.struct_CThostFtdcSettlementInfoConfirmField,
	pRspInfo *C.struct_CThostFtdcRspInfoField,
	nRequestID C.int,
	bIsLast C.bool,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRspQrySettlementInfoConfirm called",
		slog.Any("this", this),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnRspQrySettlementInfoConfirm(
		(*future.CThostFtdcSettlementInfoConfirmField)(unsafe.Pointer(pSettlementInfoConfirm)),
		(*future.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
		int(nRequestID),
		bool(bIsLast),
	)
}

//export CgoOnRspQryInvestorPositionCombineDetail
func CgoOnRspQryInvestorPositionCombineDetail(
	this unsafe.Pointer,
	pInvestorPositionCombineDetail *C.struct_CThostFtdcInvestorPositionCombineDetailField,
	pRspInfo *C.struct_CThostFtdcRspInfoField,
	nRequestID C.int,
	bIsLast C.bool,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRspQryInvestorPositionCombineDetail called",
		slog.Any("this", this),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnRspQryInvestorPositionCombineDetail(
		(*future.CThostFtdcInvestorPositionCombineDetailField)(unsafe.Pointer(pInvestorPositionCombineDetail)),
		(*future.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
		int(nRequestID),
		bool(bIsLast),
	)
}

//export CgoOnRspQryCFMMCTradingAccountKey
func CgoOnRspQryCFMMCTradingAccountKey(
	this unsafe.Pointer,
	pCFMMCTradingAccountKey *C.struct_CThostFtdcCFMMCTradingAccountKeyField,
	pRspInfo *C.struct_CThostFtdcRspInfoField,
	nRequestID C.int,
	bIsLast C.bool,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRspQryCFMMCTradingAccountKey called",
		slog.Any("this", this),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnRspQryCFMMCTradingAccountKey(
		(*future.CThostFtdcCFMMCTradingAccountKeyField)(unsafe.Pointer(pCFMMCTradingAccountKey)),
		(*future.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
		int(nRequestID),
		bool(bIsLast),
	)
}

//export CgoOnRspQryEWarrantOffset
func CgoOnRspQryEWarrantOffset(
	this unsafe.Pointer,
	pEWarrantOffset *C.struct_CThostFtdcEWarrantOffsetField,
	pRspInfo *C.struct_CThostFtdcRspInfoField,
	nRequestID C.int,
	bIsLast C.bool,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRspQryEWarrantOffset called",
		slog.Any("this", this),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnRspQryEWarrantOffset(
		(*future.CThostFtdcEWarrantOffsetField)(unsafe.Pointer(pEWarrantOffset)),
		(*future.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
		int(nRequestID),
		bool(bIsLast),
	)
}

//export CgoOnRspQryInvestorProductGroupMargin
func CgoOnRspQryInvestorProductGroupMargin(
	this unsafe.Pointer,
	pInvestorProductGroupMargin *C.struct_CThostFtdcInvestorProductGroupMarginField,
	pRspInfo *C.struct_CThostFtdcRspInfoField,
	nRequestID C.int,
	bIsLast C.bool,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRspQryInvestorProductGroupMargin called",
		slog.Any("this", this),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnRspQryInvestorProductGroupMargin(
		(*future.CThostFtdcInvestorProductGroupMarginField)(unsafe.Pointer(pInvestorProductGroupMargin)),
		(*future.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
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
		(*future.CThostFtdcExchangeMarginRateField)(unsafe.Pointer(pExchangeMarginRate)),
		(*future.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
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
		(*future.CThostFtdcExchangeMarginRateAdjustField)(unsafe.Pointer(pExchangeMarginRateAdjust)),
		(*future.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
		int(nRequestID),
		bool(bIsLast),
	)
}

//export CgoOnRspQryExchangeRate
func CgoOnRspQryExchangeRate(
	this unsafe.Pointer,
	pExchangeRate *C.struct_CThostFtdcExchangeRateField,
	pRspInfo *C.struct_CThostFtdcRspInfoField,
	nRequestID C.int,
	bIsLast C.bool,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRspQryExchangeRate called",
		slog.Any("this", this),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnRspQryExchangeRate(
		(*future.CThostFtdcExchangeRateField)(unsafe.Pointer(pExchangeRate)),
		(*future.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
		int(nRequestID),
		bool(bIsLast),
	)
}

//export CgoOnRspQrySecAgentACIDMap
func CgoOnRspQrySecAgentACIDMap(
	this unsafe.Pointer,
	pSecAgentACIDMap *C.struct_CThostFtdcSecAgentACIDMapField,
	pRspInfo *C.struct_CThostFtdcRspInfoField,
	nRequestID C.int,
	bIsLast C.bool,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRspQrySecAgentACIDMap called",
		slog.Any("this", this),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnRspQrySecAgentACIDMap(
		(*future.CThostFtdcSecAgentACIDMapField)(unsafe.Pointer(pSecAgentACIDMap)),
		(*future.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
		int(nRequestID),
		bool(bIsLast),
	)
}

//export CgoOnRspQryProductExchRate
func CgoOnRspQryProductExchRate(
	this unsafe.Pointer,
	pProductExchRate *C.struct_CThostFtdcProductExchRateField,
	pRspInfo *C.struct_CThostFtdcRspInfoField,
	nRequestID C.int,
	bIsLast C.bool,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRspQryProductExchRate called",
		slog.Any("this", this),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnRspQryProductExchRate(
		(*future.CThostFtdcProductExchRateField)(unsafe.Pointer(pProductExchRate)),
		(*future.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
		int(nRequestID),
		bool(bIsLast),
	)
}

//export CgoOnRspQryProductGroup
func CgoOnRspQryProductGroup(
	this unsafe.Pointer,
	pProductGroup *C.struct_CThostFtdcProductGroupField,
	pRspInfo *C.struct_CThostFtdcRspInfoField,
	nRequestID C.int,
	bIsLast C.bool,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRspQryProductGroup called",
		slog.Any("this", this),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnRspQryProductGroup(
		(*future.CThostFtdcProductGroupField)(unsafe.Pointer(pProductGroup)),
		(*future.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
		int(nRequestID),
		bool(bIsLast),
	)
}

//export CgoOnRspQryMMInstrumentCommissionRate
func CgoOnRspQryMMInstrumentCommissionRate(
	this unsafe.Pointer,
	pMMInstrumentCommissionRate *C.struct_CThostFtdcMMInstrumentCommissionRateField,
	pRspInfo *C.struct_CThostFtdcRspInfoField,
	nRequestID C.int,
	bIsLast C.bool,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRspQryMMInstrumentCommissionRate called",
		slog.Any("this", this),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnRspQryMMInstrumentCommissionRate(
		(*future.CThostFtdcMMInstrumentCommissionRateField)(unsafe.Pointer(pMMInstrumentCommissionRate)),
		(*future.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
		int(nRequestID),
		bool(bIsLast),
	)
}

//export CgoOnRspQryMMOptionInstrCommRate
func CgoOnRspQryMMOptionInstrCommRate(
	this unsafe.Pointer,
	pMMOptionInstrCommRate *C.struct_CThostFtdcMMOptionInstrCommRateField,
	pRspInfo *C.struct_CThostFtdcRspInfoField,
	nRequestID C.int,
	bIsLast C.bool,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRspQryMMOptionInstrCommRate called",
		slog.Any("this", this),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnRspQryMMOptionInstrCommRate(
		(*future.CThostFtdcMMOptionInstrCommRateField)(unsafe.Pointer(pMMOptionInstrCommRate)),
		(*future.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
		int(nRequestID),
		bool(bIsLast),
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
		(*future.CThostFtdcInstrumentOrderCommRateField)(unsafe.Pointer(pInstrumentOrderCommRate)),
		(*future.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
		int(nRequestID),
		bool(bIsLast),
	)
}

//export CgoOnRspQrySecAgentTradingAccount
func CgoOnRspQrySecAgentTradingAccount(
	this unsafe.Pointer,
	pTradingAccount *C.struct_CThostFtdcTradingAccountField,
	pRspInfo *C.struct_CThostFtdcRspInfoField,
	nRequestID C.int,
	bIsLast C.bool,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRspQrySecAgentTradingAccount called",
		slog.Any("this", this),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnRspQrySecAgentTradingAccount(
		(*future.CThostFtdcTradingAccountField)(unsafe.Pointer(pTradingAccount)),
		(*future.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
		int(nRequestID),
		bool(bIsLast),
	)
}

//export CgoOnRspQrySecAgentCheckMode
func CgoOnRspQrySecAgentCheckMode(
	this unsafe.Pointer,
	pSecAgentCheckMode *C.struct_CThostFtdcSecAgentCheckModeField,
	pRspInfo *C.struct_CThostFtdcRspInfoField,
	nRequestID C.int,
	bIsLast C.bool,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRspQrySecAgentCheckMode called",
		slog.Any("this", this),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnRspQrySecAgentCheckMode(
		(*future.CThostFtdcSecAgentCheckModeField)(unsafe.Pointer(pSecAgentCheckMode)),
		(*future.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
		int(nRequestID),
		bool(bIsLast),
	)
}

//export CgoOnRspQrySecAgentTradeInfo
func CgoOnRspQrySecAgentTradeInfo(
	this unsafe.Pointer,
	pSecAgentTradeInfo *C.struct_CThostFtdcSecAgentTradeInfoField,
	pRspInfo *C.struct_CThostFtdcRspInfoField,
	nRequestID C.int,
	bIsLast C.bool,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRspQrySecAgentTradeInfo called",
		slog.Any("this", this),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnRspQrySecAgentTradeInfo(
		(*future.CThostFtdcSecAgentTradeInfoField)(unsafe.Pointer(pSecAgentTradeInfo)),
		(*future.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
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
		(*future.CThostFtdcOptionInstrTradeCostField)(unsafe.Pointer(pOptionInstrTradeCost)),
		(*future.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
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
		(*future.CThostFtdcOptionInstrCommRateField)(unsafe.Pointer(pOptionInstrCommRate)),
		(*future.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
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
		(*future.CThostFtdcExecOrderField)(unsafe.Pointer(pExecOrder)),
		(*future.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
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
		(*future.CThostFtdcForQuoteField)(unsafe.Pointer(pForQuote)),
		(*future.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
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
		(*future.CThostFtdcQuoteField)(unsafe.Pointer(pQuote)),
		(*future.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
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
		(*future.CThostFtdcOptionSelfCloseField)(unsafe.Pointer(pOptionSelfClose)),
		(*future.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
		int(nRequestID),
		bool(bIsLast),
	)
}

//export CgoOnRspQryInvestUnit
func CgoOnRspQryInvestUnit(
	this unsafe.Pointer,
	pInvestUnit *C.struct_CThostFtdcInvestUnitField,
	pRspInfo *C.struct_CThostFtdcRspInfoField,
	nRequestID C.int,
	bIsLast C.bool,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRspQryInvestUnit called",
		slog.Any("this", this),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnRspQryInvestUnit(
		(*future.CThostFtdcInvestUnitField)(unsafe.Pointer(pInvestUnit)),
		(*future.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
		int(nRequestID),
		bool(bIsLast),
	)
}

//export CgoOnRspQryCombInstrumentGuard
func CgoOnRspQryCombInstrumentGuard(
	this unsafe.Pointer,
	pCombInstrumentGuard *C.struct_CThostFtdcCombInstrumentGuardField,
	pRspInfo *C.struct_CThostFtdcRspInfoField,
	nRequestID C.int,
	bIsLast C.bool,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRspQryCombInstrumentGuard called",
		slog.Any("this", this),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnRspQryCombInstrumentGuard(
		(*future.CThostFtdcCombInstrumentGuardField)(unsafe.Pointer(pCombInstrumentGuard)),
		(*future.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
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
		(*future.CThostFtdcCombActionField)(unsafe.Pointer(pCombAction)),
		(*future.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
		int(nRequestID),
		bool(bIsLast),
	)
}

//export CgoOnRspQryTransferSerial
func CgoOnRspQryTransferSerial(
	this unsafe.Pointer,
	pTransferSerial *C.struct_CThostFtdcTransferSerialField,
	pRspInfo *C.struct_CThostFtdcRspInfoField,
	nRequestID C.int,
	bIsLast C.bool,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRspQryTransferSerial called",
		slog.Any("this", this),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnRspQryTransferSerial(
		(*future.CThostFtdcTransferSerialField)(unsafe.Pointer(pTransferSerial)),
		(*future.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
		int(nRequestID),
		bool(bIsLast),
	)
}

//export CgoOnRspQryAccountregister
func CgoOnRspQryAccountregister(
	this unsafe.Pointer,
	pAccountregister *C.struct_CThostFtdcAccountregisterField,
	pRspInfo *C.struct_CThostFtdcRspInfoField,
	nRequestID C.int,
	bIsLast C.bool,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRspQryAccountregister called",
		slog.Any("this", this),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnRspQryAccountregister(
		(*future.CThostFtdcAccountregisterField)(unsafe.Pointer(pAccountregister)),
		(*future.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
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
		(*future.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
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
		(*future.CThostFtdcOrderField)(unsafe.Pointer(pOrder)),
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
		(*future.CThostFtdcTradeField)(unsafe.Pointer(pTrade)),
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
		(*future.CThostFtdcInputOrderField)(unsafe.Pointer(pInputOrder)),
		(*future.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
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
		(*future.CThostFtdcOrderActionField)(unsafe.Pointer(pOrderAction)),
		(*future.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
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
		(*future.CThostFtdcInstrumentStatusField)(unsafe.Pointer(pInstrumentStatus)),
	)
}

//export CgoOnRtnBulletin
func CgoOnRtnBulletin(
	this unsafe.Pointer,
	pBulletin *C.struct_CThostFtdcBulletinField,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRtnBulletin called",
		slog.Any("this", this),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnRtnBulletin(
		(*future.CThostFtdcBulletinField)(unsafe.Pointer(pBulletin)),
	)
}

//export CgoOnRtnTradingNotice
func CgoOnRtnTradingNotice(
	this unsafe.Pointer,
	pTradingNoticeInfo *C.struct_CThostFtdcTradingNoticeInfoField,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRtnTradingNotice called",
		slog.Any("this", this),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnRtnTradingNotice(
		(*future.CThostFtdcTradingNoticeInfoField)(unsafe.Pointer(pTradingNoticeInfo)),
	)
}

//export CgoOnRtnErrorConditionalOrder
func CgoOnRtnErrorConditionalOrder(
	this unsafe.Pointer,
	pErrorConditionalOrder *C.struct_CThostFtdcErrorConditionalOrderField,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRtnErrorConditionalOrder called",
		slog.Any("this", this),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnRtnErrorConditionalOrder(
		(*future.CThostFtdcErrorConditionalOrderField)(unsafe.Pointer(pErrorConditionalOrder)),
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
		(*future.CThostFtdcExecOrderField)(unsafe.Pointer(pExecOrder)),
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
		(*future.CThostFtdcInputExecOrderField)(unsafe.Pointer(pInputExecOrder)),
		(*future.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
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
		(*future.CThostFtdcExecOrderActionField)(unsafe.Pointer(pExecOrderAction)),
		(*future.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
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
		(*future.CThostFtdcInputForQuoteField)(unsafe.Pointer(pInputForQuote)),
		(*future.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
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
		(*future.CThostFtdcQuoteField)(unsafe.Pointer(pQuote)),
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
		(*future.CThostFtdcInputQuoteField)(unsafe.Pointer(pInputQuote)),
		(*future.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
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
		(*future.CThostFtdcQuoteActionField)(unsafe.Pointer(pQuoteAction)),
		(*future.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
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
		(*future.CThostFtdcForQuoteRspField)(unsafe.Pointer(pForQuoteRsp)),
	)
}

//export CgoOnRtnCFMMCTradingAccountToken
func CgoOnRtnCFMMCTradingAccountToken(
	this unsafe.Pointer,
	pCFMMCTradingAccountToken *C.struct_CThostFtdcCFMMCTradingAccountTokenField,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRtnCFMMCTradingAccountToken called",
		slog.Any("this", this),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnRtnCFMMCTradingAccountToken(
		(*future.CThostFtdcCFMMCTradingAccountTokenField)(unsafe.Pointer(pCFMMCTradingAccountToken)),
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
		(*future.CThostFtdcBatchOrderActionField)(unsafe.Pointer(pBatchOrderAction)),
		(*future.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
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
		(*future.CThostFtdcOptionSelfCloseField)(unsafe.Pointer(pOptionSelfClose)),
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
		(*future.CThostFtdcInputOptionSelfCloseField)(unsafe.Pointer(pInputOptionSelfClose)),
		(*future.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
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
		(*future.CThostFtdcOptionSelfCloseActionField)(unsafe.Pointer(pOptionSelfCloseAction)),
		(*future.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
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
		(*future.CThostFtdcCombActionField)(unsafe.Pointer(pCombAction)),
	)
}

//export CgoOnErrRtnCombActionInsert
func CgoOnErrRtnCombActionInsert(
	this unsafe.Pointer,
	pInputCombAction *C.struct_CThostFtdcInputCombActionField,
	pRspInfo *C.struct_CThostFtdcRspInfoField,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnErrRtnCombActionInsert called",
		slog.Any("this", this),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnErrRtnCombActionInsert(
		(*future.CThostFtdcInputCombActionField)(unsafe.Pointer(pInputCombAction)),
		(*future.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
	)
}

//export CgoOnRspQryContractBank
func CgoOnRspQryContractBank(
	this unsafe.Pointer,
	pContractBank *C.struct_CThostFtdcContractBankField,
	pRspInfo *C.struct_CThostFtdcRspInfoField,
	nRequestID C.int,
	bIsLast C.bool,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRspQryContractBank called",
		slog.Any("this", this),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnRspQryContractBank(
		(*future.CThostFtdcContractBankField)(unsafe.Pointer(pContractBank)),
		(*future.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
		int(nRequestID),
		bool(bIsLast),
	)
}

//export CgoOnRspQryParkedOrder
func CgoOnRspQryParkedOrder(
	this unsafe.Pointer,
	pParkedOrder *C.struct_CThostFtdcParkedOrderField,
	pRspInfo *C.struct_CThostFtdcRspInfoField,
	nRequestID C.int,
	bIsLast C.bool,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRspQryParkedOrder called",
		slog.Any("this", this),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnRspQryParkedOrder(
		(*future.CThostFtdcParkedOrderField)(unsafe.Pointer(pParkedOrder)),
		(*future.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
		int(nRequestID),
		bool(bIsLast),
	)
}

//export CgoOnRspQryParkedOrderAction
func CgoOnRspQryParkedOrderAction(
	this unsafe.Pointer,
	pParkedOrderAction *C.struct_CThostFtdcParkedOrderActionField,
	pRspInfo *C.struct_CThostFtdcRspInfoField,
	nRequestID C.int,
	bIsLast C.bool,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRspQryParkedOrderAction called",
		slog.Any("this", this),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnRspQryParkedOrderAction(
		(*future.CThostFtdcParkedOrderActionField)(unsafe.Pointer(pParkedOrderAction)),
		(*future.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
		int(nRequestID),
		bool(bIsLast),
	)
}

//export CgoOnRspQryTradingNotice
func CgoOnRspQryTradingNotice(
	this unsafe.Pointer,
	pTradingNotice *C.struct_CThostFtdcTradingNoticeField,
	pRspInfo *C.struct_CThostFtdcRspInfoField,
	nRequestID C.int,
	bIsLast C.bool,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRspQryTradingNotice called",
		slog.Any("this", this),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnRspQryTradingNotice(
		(*future.CThostFtdcTradingNoticeField)(unsafe.Pointer(pTradingNotice)),
		(*future.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
		int(nRequestID),
		bool(bIsLast),
	)
}

//export CgoOnRspQryBrokerTradingParams
func CgoOnRspQryBrokerTradingParams(
	this unsafe.Pointer,
	pBrokerTradingParams *C.struct_CThostFtdcBrokerTradingParamsField,
	pRspInfo *C.struct_CThostFtdcRspInfoField,
	nRequestID C.int,
	bIsLast C.bool,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRspQryBrokerTradingParams called",
		slog.Any("this", this),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnRspQryBrokerTradingParams(
		(*future.CThostFtdcBrokerTradingParamsField)(unsafe.Pointer(pBrokerTradingParams)),
		(*future.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
		int(nRequestID),
		bool(bIsLast),
	)
}

//export CgoOnRspQryBrokerTradingAlgos
func CgoOnRspQryBrokerTradingAlgos(
	this unsafe.Pointer,
	pBrokerTradingAlgos *C.struct_CThostFtdcBrokerTradingAlgosField,
	pRspInfo *C.struct_CThostFtdcRspInfoField,
	nRequestID C.int,
	bIsLast C.bool,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRspQryBrokerTradingAlgos called",
		slog.Any("this", this),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnRspQryBrokerTradingAlgos(
		(*future.CThostFtdcBrokerTradingAlgosField)(unsafe.Pointer(pBrokerTradingAlgos)),
		(*future.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
		int(nRequestID),
		bool(bIsLast),
	)
}

//export CgoOnRspQueryCFMMCTradingAccountToken
func CgoOnRspQueryCFMMCTradingAccountToken(
	this unsafe.Pointer,
	pQueryCFMMCTradingAccountToken *C.struct_CThostFtdcQueryCFMMCTradingAccountTokenField,
	pRspInfo *C.struct_CThostFtdcRspInfoField,
	nRequestID C.int,
	bIsLast C.bool,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRspQueryCFMMCTradingAccountToken called",
		slog.Any("this", this),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnRspQueryCFMMCTradingAccountToken(
		(*future.CThostFtdcQueryCFMMCTradingAccountTokenField)(unsafe.Pointer(pQueryCFMMCTradingAccountToken)),
		(*future.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
		int(nRequestID),
		bool(bIsLast),
	)
}

//export CgoOnRtnFromBankToFutureByBank
func CgoOnRtnFromBankToFutureByBank(
	this unsafe.Pointer,
	pRspTransfer *C.struct_CThostFtdcRspTransferField,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRtnFromBankToFutureByBank called",
		slog.Any("this", this),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnRtnFromBankToFutureByBank(
		(*future.CThostFtdcRspTransferField)(unsafe.Pointer(pRspTransfer)),
	)
}

//export CgoOnRtnFromFutureToBankByBank
func CgoOnRtnFromFutureToBankByBank(
	this unsafe.Pointer,
	pRspTransfer *C.struct_CThostFtdcRspTransferField,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRtnFromFutureToBankByBank called",
		slog.Any("this", this),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnRtnFromFutureToBankByBank(
		(*future.CThostFtdcRspTransferField)(unsafe.Pointer(pRspTransfer)),
	)
}

//export CgoOnRtnRepealFromBankToFutureByBank
func CgoOnRtnRepealFromBankToFutureByBank(
	this unsafe.Pointer,
	pRspRepeal *C.struct_CThostFtdcRspRepealField,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRtnRepealFromBankToFutureByBank called",
		slog.Any("this", this),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnRtnRepealFromBankToFutureByBank(
		(*future.CThostFtdcRspRepealField)(unsafe.Pointer(pRspRepeal)),
	)
}

//export CgoOnRtnRepealFromFutureToBankByBank
func CgoOnRtnRepealFromFutureToBankByBank(
	this unsafe.Pointer,
	pRspRepeal *C.struct_CThostFtdcRspRepealField,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRtnRepealFromFutureToBankByBank called",
		slog.Any("this", this),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnRtnRepealFromFutureToBankByBank(
		(*future.CThostFtdcRspRepealField)(unsafe.Pointer(pRspRepeal)),
	)
}

//export CgoOnRtnFromBankToFutureByFuture
func CgoOnRtnFromBankToFutureByFuture(
	this unsafe.Pointer,
	pRspTransfer *C.struct_CThostFtdcRspTransferField,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRtnFromBankToFutureByFuture called",
		slog.Any("this", this),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnRtnFromBankToFutureByFuture(
		(*future.CThostFtdcRspTransferField)(unsafe.Pointer(pRspTransfer)),
	)
}

//export CgoOnRtnFromFutureToBankByFuture
func CgoOnRtnFromFutureToBankByFuture(
	this unsafe.Pointer,
	pRspTransfer *C.struct_CThostFtdcRspTransferField,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRtnFromFutureToBankByFuture called",
		slog.Any("this", this),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnRtnFromFutureToBankByFuture(
		(*future.CThostFtdcRspTransferField)(unsafe.Pointer(pRspTransfer)),
	)
}

//export CgoOnRtnRepealFromBankToFutureByFutureManual
func CgoOnRtnRepealFromBankToFutureByFutureManual(
	this unsafe.Pointer,
	pRspRepeal *C.struct_CThostFtdcRspRepealField,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRtnRepealFromBankToFutureByFutureManual called",
		slog.Any("this", this),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnRtnRepealFromBankToFutureByFutureManual(
		(*future.CThostFtdcRspRepealField)(unsafe.Pointer(pRspRepeal)),
	)
}

//export CgoOnRtnRepealFromFutureToBankByFutureManual
func CgoOnRtnRepealFromFutureToBankByFutureManual(
	this unsafe.Pointer,
	pRspRepeal *C.struct_CThostFtdcRspRepealField,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRtnRepealFromFutureToBankByFutureManual called",
		slog.Any("this", this),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnRtnRepealFromFutureToBankByFutureManual(
		(*future.CThostFtdcRspRepealField)(unsafe.Pointer(pRspRepeal)),
	)
}

//export CgoOnRtnQueryBankBalanceByFuture
func CgoOnRtnQueryBankBalanceByFuture(
	this unsafe.Pointer,
	pNotifyQueryAccount *C.struct_CThostFtdcNotifyQueryAccountField,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRtnQueryBankBalanceByFuture called",
		slog.Any("this", this),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnRtnQueryBankBalanceByFuture(
		(*future.CThostFtdcNotifyQueryAccountField)(unsafe.Pointer(pNotifyQueryAccount)),
	)
}

//export CgoOnErrRtnBankToFutureByFuture
func CgoOnErrRtnBankToFutureByFuture(
	this unsafe.Pointer,
	pReqTransfer *C.struct_CThostFtdcReqTransferField,
	pRspInfo *C.struct_CThostFtdcRspInfoField,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnErrRtnBankToFutureByFuture called",
		slog.Any("this", this),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnErrRtnBankToFutureByFuture(
		(*future.CThostFtdcReqTransferField)(unsafe.Pointer(pReqTransfer)),
		(*future.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
	)
}

//export CgoOnErrRtnFutureToBankByFuture
func CgoOnErrRtnFutureToBankByFuture(
	this unsafe.Pointer,
	pReqTransfer *C.struct_CThostFtdcReqTransferField,
	pRspInfo *C.struct_CThostFtdcRspInfoField,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnErrRtnFutureToBankByFuture called",
		slog.Any("this", this),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnErrRtnFutureToBankByFuture(
		(*future.CThostFtdcReqTransferField)(unsafe.Pointer(pReqTransfer)),
		(*future.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
	)
}

//export CgoOnErrRtnRepealBankToFutureByFutureManual
func CgoOnErrRtnRepealBankToFutureByFutureManual(
	this unsafe.Pointer,
	pReqRepeal *C.struct_CThostFtdcReqRepealField,
	pRspInfo *C.struct_CThostFtdcRspInfoField,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnErrRtnRepealBankToFutureByFutureManual called",
		slog.Any("this", this),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnErrRtnRepealBankToFutureByFutureManual(
		(*future.CThostFtdcReqRepealField)(unsafe.Pointer(pReqRepeal)),
		(*future.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
	)
}

//export CgoOnErrRtnRepealFutureToBankByFutureManual
func CgoOnErrRtnRepealFutureToBankByFutureManual(
	this unsafe.Pointer,
	pReqRepeal *C.struct_CThostFtdcReqRepealField,
	pRspInfo *C.struct_CThostFtdcRspInfoField,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnErrRtnRepealFutureToBankByFutureManual called",
		slog.Any("this", this),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnErrRtnRepealFutureToBankByFutureManual(
		(*future.CThostFtdcReqRepealField)(unsafe.Pointer(pReqRepeal)),
		(*future.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
	)
}

//export CgoOnErrRtnQueryBankBalanceByFuture
func CgoOnErrRtnQueryBankBalanceByFuture(
	this unsafe.Pointer,
	pReqQueryAccount *C.struct_CThostFtdcReqQueryAccountField,
	pRspInfo *C.struct_CThostFtdcRspInfoField,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnErrRtnQueryBankBalanceByFuture called",
		slog.Any("this", this),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnErrRtnQueryBankBalanceByFuture(
		(*future.CThostFtdcReqQueryAccountField)(unsafe.Pointer(pReqQueryAccount)),
		(*future.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
	)
}

//export CgoOnRtnRepealFromBankToFutureByFuture
func CgoOnRtnRepealFromBankToFutureByFuture(
	this unsafe.Pointer,
	pRspRepeal *C.struct_CThostFtdcRspRepealField,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRtnRepealFromBankToFutureByFuture called",
		slog.Any("this", this),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnRtnRepealFromBankToFutureByFuture(
		(*future.CThostFtdcRspRepealField)(unsafe.Pointer(pRspRepeal)),
	)
}

//export CgoOnRtnRepealFromFutureToBankByFuture
func CgoOnRtnRepealFromFutureToBankByFuture(
	this unsafe.Pointer,
	pRspRepeal *C.struct_CThostFtdcRspRepealField,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRtnRepealFromFutureToBankByFuture called",
		slog.Any("this", this),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnRtnRepealFromFutureToBankByFuture(
		(*future.CThostFtdcRspRepealField)(unsafe.Pointer(pRspRepeal)),
	)
}

//export CgoOnRspFromBankToFutureByFuture
func CgoOnRspFromBankToFutureByFuture(
	this unsafe.Pointer,
	pReqTransfer *C.struct_CThostFtdcReqTransferField,
	pRspInfo *C.struct_CThostFtdcRspInfoField,
	nRequestID C.int,
	bIsLast C.bool,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRspFromBankToFutureByFuture called",
		slog.Any("this", this),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnRspFromBankToFutureByFuture(
		(*future.CThostFtdcReqTransferField)(unsafe.Pointer(pReqTransfer)),
		(*future.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
		int(nRequestID),
		bool(bIsLast),
	)
}

//export CgoOnRspFromFutureToBankByFuture
func CgoOnRspFromFutureToBankByFuture(
	this unsafe.Pointer,
	pReqTransfer *C.struct_CThostFtdcReqTransferField,
	pRspInfo *C.struct_CThostFtdcRspInfoField,
	nRequestID C.int,
	bIsLast C.bool,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRspFromFutureToBankByFuture called",
		slog.Any("this", this),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnRspFromFutureToBankByFuture(
		(*future.CThostFtdcReqTransferField)(unsafe.Pointer(pReqTransfer)),
		(*future.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
		int(nRequestID),
		bool(bIsLast),
	)
}

//export CgoOnRspQueryBankAccountMoneyByFuture
func CgoOnRspQueryBankAccountMoneyByFuture(
	this unsafe.Pointer,
	pReqQueryAccount *C.struct_CThostFtdcReqQueryAccountField,
	pRspInfo *C.struct_CThostFtdcRspInfoField,
	nRequestID C.int,
	bIsLast C.bool,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRspQueryBankAccountMoneyByFuture called",
		slog.Any("this", this),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnRspQueryBankAccountMoneyByFuture(
		(*future.CThostFtdcReqQueryAccountField)(unsafe.Pointer(pReqQueryAccount)),
		(*future.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
		int(nRequestID),
		bool(bIsLast),
	)
}

//export CgoOnRtnOpenAccountByBank
func CgoOnRtnOpenAccountByBank(
	this unsafe.Pointer,
	pOpenAccount *C.struct_CThostFtdcOpenAccountField,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRtnOpenAccountByBank called",
		slog.Any("this", this),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnRtnOpenAccountByBank(
		(*future.CThostFtdcOpenAccountField)(unsafe.Pointer(pOpenAccount)),
	)
}

//export CgoOnRtnCancelAccountByBank
func CgoOnRtnCancelAccountByBank(
	this unsafe.Pointer,
	pCancelAccount *C.struct_CThostFtdcCancelAccountField,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRtnCancelAccountByBank called",
		slog.Any("this", this),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnRtnCancelAccountByBank(
		(*future.CThostFtdcCancelAccountField)(unsafe.Pointer(pCancelAccount)),
	)
}

//export CgoOnRtnChangeAccountByBank
func CgoOnRtnChangeAccountByBank(
	this unsafe.Pointer,
	pChangeAccount *C.struct_CThostFtdcChangeAccountField,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRtnChangeAccountByBank called",
		slog.Any("this", this),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnRtnChangeAccountByBank(
		(*future.CThostFtdcChangeAccountField)(unsafe.Pointer(pChangeAccount)),
	)
}

//export CgoOnRspQryClassifiedInstrument
func CgoOnRspQryClassifiedInstrument(
	this unsafe.Pointer,
	pInstrument *C.struct_CThostFtdcInstrumentField,
	pRspInfo *C.struct_CThostFtdcRspInfoField,
	nRequestID C.int,
	bIsLast C.bool,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRspQryClassifiedInstrument called",
		slog.Any("this", this),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnRspQryClassifiedInstrument(
		(*future.CThostFtdcInstrumentField)(unsafe.Pointer(pInstrument)),
		(*future.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
		int(nRequestID),
		bool(bIsLast),
	)
}

//export CgoOnRspQryCombPromotionParam
func CgoOnRspQryCombPromotionParam(
	this unsafe.Pointer,
	pCombPromotionParam *C.struct_CThostFtdcCombPromotionParamField,
	pRspInfo *C.struct_CThostFtdcRspInfoField,
	nRequestID C.int,
	bIsLast C.bool,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRspQryCombPromotionParam called",
		slog.Any("this", this),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnRspQryCombPromotionParam(
		(*future.CThostFtdcCombPromotionParamField)(unsafe.Pointer(pCombPromotionParam)),
		(*future.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
		int(nRequestID),
		bool(bIsLast),
	)
}

//export CgoOnRspQryRiskSettleInvstPosition
func CgoOnRspQryRiskSettleInvstPosition(
	this unsafe.Pointer,
	pRiskSettleInvstPosition *C.struct_CThostFtdcRiskSettleInvstPositionField,
	pRspInfo *C.struct_CThostFtdcRspInfoField,
	nRequestID C.int,
	bIsLast C.bool,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRspQryRiskSettleInvstPosition called",
		slog.Any("this", this),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnRspQryRiskSettleInvstPosition(
		(*future.CThostFtdcRiskSettleInvstPositionField)(unsafe.Pointer(pRiskSettleInvstPosition)),
		(*future.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
		int(nRequestID),
		bool(bIsLast),
	)
}

//export CgoOnRspQryRiskSettleProductStatus
func CgoOnRspQryRiskSettleProductStatus(
	this unsafe.Pointer,
	pRiskSettleProductStatus *C.struct_CThostFtdcRiskSettleProductStatusField,
	pRspInfo *C.struct_CThostFtdcRspInfoField,
	nRequestID C.int,
	bIsLast C.bool,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRspQryRiskSettleProductStatus called",
		slog.Any("this", this),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnRspQryRiskSettleProductStatus(
		(*future.CThostFtdcRiskSettleProductStatusField)(unsafe.Pointer(pRiskSettleProductStatus)),
		(*future.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
		int(nRequestID),
		bool(bIsLast),
	)
}

//export CgoOnRspQrySPBMFutureParameter
func CgoOnRspQrySPBMFutureParameter(
	this unsafe.Pointer,
	pSPBMFutureParameter *C.struct_CThostFtdcSPBMFutureParameterField,
	pRspInfo *C.struct_CThostFtdcRspInfoField,
	nRequestID C.int,
	bIsLast C.bool,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRspQrySPBMFutureParameter called",
		slog.Any("this", this),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnRspQrySPBMFutureParameter(
		(*future.CThostFtdcSPBMFutureParameterField)(unsafe.Pointer(pSPBMFutureParameter)),
		(*future.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
		int(nRequestID),
		bool(bIsLast),
	)
}

//export CgoOnRspQrySPBMOptionParameter
func CgoOnRspQrySPBMOptionParameter(
	this unsafe.Pointer,
	pSPBMOptionParameter *C.struct_CThostFtdcSPBMOptionParameterField,
	pRspInfo *C.struct_CThostFtdcRspInfoField,
	nRequestID C.int,
	bIsLast C.bool,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRspQrySPBMOptionParameter called",
		slog.Any("this", this),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnRspQrySPBMOptionParameter(
		(*future.CThostFtdcSPBMOptionParameterField)(unsafe.Pointer(pSPBMOptionParameter)),
		(*future.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
		int(nRequestID),
		bool(bIsLast),
	)
}

//export CgoOnRspQrySPBMIntraParameter
func CgoOnRspQrySPBMIntraParameter(
	this unsafe.Pointer,
	pSPBMIntraParameter *C.struct_CThostFtdcSPBMIntraParameterField,
	pRspInfo *C.struct_CThostFtdcRspInfoField,
	nRequestID C.int,
	bIsLast C.bool,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRspQrySPBMIntraParameter called",
		slog.Any("this", this),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnRspQrySPBMIntraParameter(
		(*future.CThostFtdcSPBMIntraParameterField)(unsafe.Pointer(pSPBMIntraParameter)),
		(*future.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
		int(nRequestID),
		bool(bIsLast),
	)
}

//export CgoOnRspQrySPBMInterParameter
func CgoOnRspQrySPBMInterParameter(
	this unsafe.Pointer,
	pSPBMInterParameter *C.struct_CThostFtdcSPBMInterParameterField,
	pRspInfo *C.struct_CThostFtdcRspInfoField,
	nRequestID C.int,
	bIsLast C.bool,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRspQrySPBMInterParameter called",
		slog.Any("this", this),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnRspQrySPBMInterParameter(
		(*future.CThostFtdcSPBMInterParameterField)(unsafe.Pointer(pSPBMInterParameter)),
		(*future.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
		int(nRequestID),
		bool(bIsLast),
	)
}

//export CgoOnRspQrySPBMPortfDefinition
func CgoOnRspQrySPBMPortfDefinition(
	this unsafe.Pointer,
	pSPBMPortfDefinition *C.struct_CThostFtdcSPBMPortfDefinitionField,
	pRspInfo *C.struct_CThostFtdcRspInfoField,
	nRequestID C.int,
	bIsLast C.bool,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRspQrySPBMPortfDefinition called",
		slog.Any("this", this),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnRspQrySPBMPortfDefinition(
		(*future.CThostFtdcSPBMPortfDefinitionField)(unsafe.Pointer(pSPBMPortfDefinition)),
		(*future.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
		int(nRequestID),
		bool(bIsLast),
	)
}

//export CgoOnRspQrySPBMInvestorPortfDef
func CgoOnRspQrySPBMInvestorPortfDef(
	this unsafe.Pointer,
	pSPBMInvestorPortfDef *C.struct_CThostFtdcSPBMInvestorPortfDefField,
	pRspInfo *C.struct_CThostFtdcRspInfoField,
	nRequestID C.int,
	bIsLast C.bool,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRspQrySPBMInvestorPortfDef called",
		slog.Any("this", this),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnRspQrySPBMInvestorPortfDef(
		(*future.CThostFtdcSPBMInvestorPortfDefField)(unsafe.Pointer(pSPBMInvestorPortfDef)),
		(*future.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
		int(nRequestID),
		bool(bIsLast),
	)
}

//export CgoOnRspQryInvestorPortfMarginRatio
func CgoOnRspQryInvestorPortfMarginRatio(
	this unsafe.Pointer,
	pInvestorPortfMarginRatio *C.struct_CThostFtdcInvestorPortfMarginRatioField,
	pRspInfo *C.struct_CThostFtdcRspInfoField,
	nRequestID C.int,
	bIsLast C.bool,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRspQryInvestorPortfMarginRatio called",
		slog.Any("this", this),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnRspQryInvestorPortfMarginRatio(
		(*future.CThostFtdcInvestorPortfMarginRatioField)(unsafe.Pointer(pInvestorPortfMarginRatio)),
		(*future.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
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
		(*future.CThostFtdcInvestorProdSPBMDetailField)(unsafe.Pointer(pInvestorProdSPBMDetail)),
		(*future.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
		int(nRequestID),
		bool(bIsLast),
	)
}

//export CgoOnRspQryInvestorCommoditySPMMMargin
func CgoOnRspQryInvestorCommoditySPMMMargin(
	this unsafe.Pointer,
	pInvestorCommoditySPMMMargin *C.struct_CThostFtdcInvestorCommoditySPMMMarginField,
	pRspInfo *C.struct_CThostFtdcRspInfoField,
	nRequestID C.int,
	bIsLast C.bool,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRspQryInvestorCommoditySPMMMargin called",
		slog.Any("this", this),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnRspQryInvestorCommoditySPMMMargin(
		(*future.CThostFtdcInvestorCommoditySPMMMarginField)(unsafe.Pointer(pInvestorCommoditySPMMMargin)),
		(*future.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
		int(nRequestID),
		bool(bIsLast),
	)
}

//export CgoOnRspQryInvestorCommodityGroupSPMMMargin
func CgoOnRspQryInvestorCommodityGroupSPMMMargin(
	this unsafe.Pointer,
	pInvestorCommodityGroupSPMMMargin *C.struct_CThostFtdcInvestorCommodityGroupSPMMMarginField,
	pRspInfo *C.struct_CThostFtdcRspInfoField,
	nRequestID C.int,
	bIsLast C.bool,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRspQryInvestorCommodityGroupSPMMMargin called",
		slog.Any("this", this),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnRspQryInvestorCommodityGroupSPMMMargin(
		(*future.CThostFtdcInvestorCommodityGroupSPMMMarginField)(unsafe.Pointer(pInvestorCommodityGroupSPMMMargin)),
		(*future.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
		int(nRequestID),
		bool(bIsLast),
	)
}

//export CgoOnRspQrySPMMInstParam
func CgoOnRspQrySPMMInstParam(
	this unsafe.Pointer,
	pSPMMInstParam *C.struct_CThostFtdcSPMMInstParamField,
	pRspInfo *C.struct_CThostFtdcRspInfoField,
	nRequestID C.int,
	bIsLast C.bool,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRspQrySPMMInstParam called",
		slog.Any("this", this),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnRspQrySPMMInstParam(
		(*future.CThostFtdcSPMMInstParamField)(unsafe.Pointer(pSPMMInstParam)),
		(*future.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
		int(nRequestID),
		bool(bIsLast),
	)
}

//export CgoOnRspQrySPMMProductParam
func CgoOnRspQrySPMMProductParam(
	this unsafe.Pointer,
	pSPMMProductParam *C.struct_CThostFtdcSPMMProductParamField,
	pRspInfo *C.struct_CThostFtdcRspInfoField,
	nRequestID C.int,
	bIsLast C.bool,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRspQrySPMMProductParam called",
		slog.Any("this", this),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnRspQrySPMMProductParam(
		(*future.CThostFtdcSPMMProductParamField)(unsafe.Pointer(pSPMMProductParam)),
		(*future.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
		int(nRequestID),
		bool(bIsLast),
	)
}

//export CgoOnRspQrySPBMAddOnInterParameter
func CgoOnRspQrySPBMAddOnInterParameter(
	this unsafe.Pointer,
	pSPBMAddOnInterParameter *C.struct_CThostFtdcSPBMAddOnInterParameterField,
	pRspInfo *C.struct_CThostFtdcRspInfoField,
	nRequestID C.int,
	bIsLast C.bool,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRspQrySPBMAddOnInterParameter called",
		slog.Any("this", this),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnRspQrySPBMAddOnInterParameter(
		(*future.CThostFtdcSPBMAddOnInterParameterField)(unsafe.Pointer(pSPBMAddOnInterParameter)),
		(*future.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
		int(nRequestID),
		bool(bIsLast),
	)
}

//export CgoOnRspQryRCAMSCombProductInfo
func CgoOnRspQryRCAMSCombProductInfo(
	this unsafe.Pointer,
	pRCAMSCombProductInfo *C.struct_CThostFtdcRCAMSCombProductInfoField,
	pRspInfo *C.struct_CThostFtdcRspInfoField,
	nRequestID C.int,
	bIsLast C.bool,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRspQryRCAMSCombProductInfo called",
		slog.Any("this", this),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnRspQryRCAMSCombProductInfo(
		(*future.CThostFtdcRCAMSCombProductInfoField)(unsafe.Pointer(pRCAMSCombProductInfo)),
		(*future.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
		int(nRequestID),
		bool(bIsLast),
	)
}

//export CgoOnRspQryRCAMSInstrParameter
func CgoOnRspQryRCAMSInstrParameter(
	this unsafe.Pointer,
	pRCAMSInstrParameter *C.struct_CThostFtdcRCAMSInstrParameterField,
	pRspInfo *C.struct_CThostFtdcRspInfoField,
	nRequestID C.int,
	bIsLast C.bool,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRspQryRCAMSInstrParameter called",
		slog.Any("this", this),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnRspQryRCAMSInstrParameter(
		(*future.CThostFtdcRCAMSInstrParameterField)(unsafe.Pointer(pRCAMSInstrParameter)),
		(*future.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
		int(nRequestID),
		bool(bIsLast),
	)
}

//export CgoOnRspQryRCAMSIntraParameter
func CgoOnRspQryRCAMSIntraParameter(
	this unsafe.Pointer,
	pRCAMSIntraParameter *C.struct_CThostFtdcRCAMSIntraParameterField,
	pRspInfo *C.struct_CThostFtdcRspInfoField,
	nRequestID C.int,
	bIsLast C.bool,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRspQryRCAMSIntraParameter called",
		slog.Any("this", this),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnRspQryRCAMSIntraParameter(
		(*future.CThostFtdcRCAMSIntraParameterField)(unsafe.Pointer(pRCAMSIntraParameter)),
		(*future.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
		int(nRequestID),
		bool(bIsLast),
	)
}

//export CgoOnRspQryRCAMSInterParameter
func CgoOnRspQryRCAMSInterParameter(
	this unsafe.Pointer,
	pRCAMSInterParameter *C.struct_CThostFtdcRCAMSInterParameterField,
	pRspInfo *C.struct_CThostFtdcRspInfoField,
	nRequestID C.int,
	bIsLast C.bool,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRspQryRCAMSInterParameter called",
		slog.Any("this", this),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnRspQryRCAMSInterParameter(
		(*future.CThostFtdcRCAMSInterParameterField)(unsafe.Pointer(pRCAMSInterParameter)),
		(*future.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
		int(nRequestID),
		bool(bIsLast),
	)
}

//export CgoOnRspQryRCAMSShortOptAdjustParam
func CgoOnRspQryRCAMSShortOptAdjustParam(
	this unsafe.Pointer,
	pRCAMSShortOptAdjustParam *C.struct_CThostFtdcRCAMSShortOptAdjustParamField,
	pRspInfo *C.struct_CThostFtdcRspInfoField,
	nRequestID C.int,
	bIsLast C.bool,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRspQryRCAMSShortOptAdjustParam called",
		slog.Any("this", this),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnRspQryRCAMSShortOptAdjustParam(
		(*future.CThostFtdcRCAMSShortOptAdjustParamField)(unsafe.Pointer(pRCAMSShortOptAdjustParam)),
		(*future.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
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
		(*future.CThostFtdcRCAMSInvestorCombPositionField)(unsafe.Pointer(pRCAMSInvestorCombPosition)),
		(*future.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
		int(nRequestID),
		bool(bIsLast),
	)
}

//export CgoOnRspQryInvestorProdRCAMSMargin
func CgoOnRspQryInvestorProdRCAMSMargin(
	this unsafe.Pointer,
	pInvestorProdRCAMSMargin *C.struct_CThostFtdcInvestorProdRCAMSMarginField,
	pRspInfo *C.struct_CThostFtdcRspInfoField,
	nRequestID C.int,
	bIsLast C.bool,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRspQryInvestorProdRCAMSMargin called",
		slog.Any("this", this),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnRspQryInvestorProdRCAMSMargin(
		(*future.CThostFtdcInvestorProdRCAMSMarginField)(unsafe.Pointer(pInvestorProdRCAMSMargin)),
		(*future.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
		int(nRequestID),
		bool(bIsLast),
	)
}

//export CgoOnRspQryRULEInstrParameter
func CgoOnRspQryRULEInstrParameter(
	this unsafe.Pointer,
	pRULEInstrParameter *C.struct_CThostFtdcRULEInstrParameterField,
	pRspInfo *C.struct_CThostFtdcRspInfoField,
	nRequestID C.int,
	bIsLast C.bool,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRspQryRULEInstrParameter called",
		slog.Any("this", this),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnRspQryRULEInstrParameter(
		(*future.CThostFtdcRULEInstrParameterField)(unsafe.Pointer(pRULEInstrParameter)),
		(*future.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
		int(nRequestID),
		bool(bIsLast),
	)
}

//export CgoOnRspQryRULEIntraParameter
func CgoOnRspQryRULEIntraParameter(
	this unsafe.Pointer,
	pRULEIntraParameter *C.struct_CThostFtdcRULEIntraParameterField,
	pRspInfo *C.struct_CThostFtdcRspInfoField,
	nRequestID C.int,
	bIsLast C.bool,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRspQryRULEIntraParameter called",
		slog.Any("this", this),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnRspQryRULEIntraParameter(
		(*future.CThostFtdcRULEIntraParameterField)(unsafe.Pointer(pRULEIntraParameter)),
		(*future.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
		int(nRequestID),
		bool(bIsLast),
	)
}

//export CgoOnRspQryRULEInterParameter
func CgoOnRspQryRULEInterParameter(
	this unsafe.Pointer,
	pRULEInterParameter *C.struct_CThostFtdcRULEInterParameterField,
	pRspInfo *C.struct_CThostFtdcRspInfoField,
	nRequestID C.int,
	bIsLast C.bool,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRspQryRULEInterParameter called",
		slog.Any("this", this),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnRspQryRULEInterParameter(
		(*future.CThostFtdcRULEInterParameterField)(unsafe.Pointer(pRULEInterParameter)),
		(*future.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
		int(nRequestID),
		bool(bIsLast),
	)
}

//export CgoOnRspQryInvestorProdRULEMargin
func CgoOnRspQryInvestorProdRULEMargin(
	this unsafe.Pointer,
	pInvestorProdRULEMargin *C.struct_CThostFtdcInvestorProdRULEMarginField,
	pRspInfo *C.struct_CThostFtdcRspInfoField,
	nRequestID C.int,
	bIsLast C.bool,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRspQryInvestorProdRULEMargin called",
		slog.Any("this", this),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnRspQryInvestorProdRULEMargin(
		(*future.CThostFtdcInvestorProdRULEMarginField)(unsafe.Pointer(pInvestorProdRULEMargin)),
		(*future.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
		int(nRequestID),
		bool(bIsLast),
	)
}

//export CgoOnRspQryInvestorPortfSetting
func CgoOnRspQryInvestorPortfSetting(
	this unsafe.Pointer,
	pInvestorPortfSetting *C.struct_CThostFtdcInvestorPortfSettingField,
	pRspInfo *C.struct_CThostFtdcRspInfoField,
	nRequestID C.int,
	bIsLast C.bool,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRspQryInvestorPortfSetting called",
		slog.Any("this", this),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnRspQryInvestorPortfSetting(
		(*future.CThostFtdcInvestorPortfSettingField)(unsafe.Pointer(pInvestorPortfSetting)),
		(*future.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
		int(nRequestID),
		bool(bIsLast),
	)
}

//export CgoOnRspQryInvestorInfoCommRec
func CgoOnRspQryInvestorInfoCommRec(
	this unsafe.Pointer,
	pInvestorInfoCommRec *C.struct_CThostFtdcInvestorInfoCommRecField,
	pRspInfo *C.struct_CThostFtdcRspInfoField,
	nRequestID C.int,
	bIsLast C.bool,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRspQryInvestorInfoCommRec called",
		slog.Any("this", this),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnRspQryInvestorInfoCommRec(
		(*future.CThostFtdcInvestorInfoCommRecField)(unsafe.Pointer(pInvestorInfoCommRec)),
		(*future.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
		int(nRequestID),
		bool(bIsLast),
	)
}

//export CgoOnRspQryCombLeg
func CgoOnRspQryCombLeg(
	this unsafe.Pointer,
	pCombLeg *C.struct_CThostFtdcCombLegField,
	pRspInfo *C.struct_CThostFtdcRspInfoField,
	nRequestID C.int,
	bIsLast C.bool,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRspQryCombLeg called",
		slog.Any("this", this),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnRspQryCombLeg(
		(*future.CThostFtdcCombLegField)(unsafe.Pointer(pCombLeg)),
		(*future.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
		int(nRequestID),
		bool(bIsLast),
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
		(*future.CThostFtdcInputOffsetSettingField)(unsafe.Pointer(pInputOffsetSetting)),
		(*future.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
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
		(*future.CThostFtdcInputOffsetSettingField)(unsafe.Pointer(pInputOffsetSetting)),
		(*future.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
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
		(*future.CThostFtdcOffsetSettingField)(unsafe.Pointer(pOffsetSetting)),
	)
}

//export CgoOnErrRtnOffsetSetting
func CgoOnErrRtnOffsetSetting(
	this unsafe.Pointer,
	pInputOffsetSetting *C.struct_CThostFtdcInputOffsetSettingField,
	pRspInfo *C.struct_CThostFtdcRspInfoField,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnErrRtnOffsetSetting called",
		slog.Any("this", this),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnErrRtnOffsetSetting(
		(*future.CThostFtdcInputOffsetSettingField)(unsafe.Pointer(pInputOffsetSetting)),
		(*future.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
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
		(*future.CThostFtdcCancelOffsetSettingField)(unsafe.Pointer(pCancelOffsetSetting)),
		(*future.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
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
		(*future.CThostFtdcOffsetSettingField)(unsafe.Pointer(pOffsetSetting)),
		(*future.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
		int(nRequestID),
		bool(bIsLast),
	)
}

//export CgoOnRspGenSMSCode
func CgoOnRspGenSMSCode(
	this unsafe.Pointer,
	pRspGenSMSCode *C.struct_CThostFtdcRspGenSMSCodeField,
	pRspInfo *C.struct_CThostFtdcRspInfoField,
	nRequestID C.int,
	bIsLast C.bool,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRspGenSMSCode called",
		slog.Any("this", this),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnRspGenSMSCode(
		(*future.CThostFtdcRspGenSMSCodeField)(unsafe.Pointer(pRspGenSMSCode)),
		(*future.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
		int(nRequestID),
		bool(bIsLast),
	)
}

//export CgoOnRspSpdApply
func CgoOnRspSpdApply(
	this unsafe.Pointer,
	pInputSpdApply *C.struct_CThostFtdcInputSpdApplyField,
	pRspInfo *C.struct_CThostFtdcRspInfoField,
	nRequestID C.int,
	bIsLast C.bool,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRspSpdApply called",
		slog.Any("this", this),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnRspSpdApply(
		(*future.CThostFtdcInputSpdApplyField)(unsafe.Pointer(pInputSpdApply)),
		(*future.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
		int(nRequestID),
		bool(bIsLast),
	)
}

//export CgoOnRspSpdApplyAction
func CgoOnRspSpdApplyAction(
	this unsafe.Pointer,
	pInputSpdApplyAction *C.struct_CThostFtdcInputSpdApplyActionField,
	pRspInfo *C.struct_CThostFtdcRspInfoField,
	nRequestID C.int,
	bIsLast C.bool,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRspSpdApplyAction called",
		slog.Any("this", this),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnRspSpdApplyAction(
		(*future.CThostFtdcInputSpdApplyActionField)(unsafe.Pointer(pInputSpdApplyAction)),
		(*future.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
		int(nRequestID),
		bool(bIsLast),
	)
}

//export CgoOnRspQrySpdApply
func CgoOnRspQrySpdApply(
	this unsafe.Pointer,
	pSpdApply *C.struct_CThostFtdcSpdApplyField,
	pRspInfo *C.struct_CThostFtdcRspInfoField,
	nRequestID C.int,
	bIsLast C.bool,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRspQrySpdApply called",
		slog.Any("this", this),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnRspQrySpdApply(
		(*future.CThostFtdcSpdApplyField)(unsafe.Pointer(pSpdApply)),
		(*future.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
		int(nRequestID),
		bool(bIsLast),
	)
}

//export CgoOnRtnSpdApply
func CgoOnRtnSpdApply(
	this unsafe.Pointer,
	pSpdApply *C.struct_CThostFtdcSpdApplyField,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRtnSpdApply called",
		slog.Any("this", this),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnRtnSpdApply(
		(*future.CThostFtdcSpdApplyField)(unsafe.Pointer(pSpdApply)),
	)
}

//export CgoOnErrRtnSpdApply
func CgoOnErrRtnSpdApply(
	this unsafe.Pointer,
	pInputSpdApply *C.struct_CThostFtdcInputSpdApplyField,
	pRspInfo *C.struct_CThostFtdcRspInfoField,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnErrRtnSpdApply called",
		slog.Any("this", this),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnErrRtnSpdApply(
		(*future.CThostFtdcInputSpdApplyField)(unsafe.Pointer(pInputSpdApply)),
		(*future.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
	)
}

//export CgoOnErrRtnSpdApplyAction
func CgoOnErrRtnSpdApplyAction(
	this unsafe.Pointer,
	pSpdApplyAction *C.struct_CThostFtdcSpdApplyActionField,
	pRspInfo *C.struct_CThostFtdcRspInfoField,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnErrRtnSpdApplyAction called",
		slog.Any("this", this),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnErrRtnSpdApplyAction(
		(*future.CThostFtdcSpdApplyActionField)(unsafe.Pointer(pSpdApplyAction)),
		(*future.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
	)
}

//export CgoOnRspHedgeCfm
func CgoOnRspHedgeCfm(
	this unsafe.Pointer,
	pInputHedgeCfm *C.struct_CThostFtdcInputHedgeCfmField,
	pRspInfo *C.struct_CThostFtdcRspInfoField,
	nRequestID C.int,
	bIsLast C.bool,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRspHedgeCfm called",
		slog.Any("this", this),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnRspHedgeCfm(
		(*future.CThostFtdcInputHedgeCfmField)(unsafe.Pointer(pInputHedgeCfm)),
		(*future.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
		int(nRequestID),
		bool(bIsLast),
	)
}

//export CgoOnRspHedgeCfmAction
func CgoOnRspHedgeCfmAction(
	this unsafe.Pointer,
	pInputHedgeCfmAction *C.struct_CThostFtdcInputHedgeCfmActionField,
	pRspInfo *C.struct_CThostFtdcRspInfoField,
	nRequestID C.int,
	bIsLast C.bool,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRspHedgeCfmAction called",
		slog.Any("this", this),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnRspHedgeCfmAction(
		(*future.CThostFtdcInputHedgeCfmActionField)(unsafe.Pointer(pInputHedgeCfmAction)),
		(*future.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
		int(nRequestID),
		bool(bIsLast),
	)
}

//export CgoOnRspQryHedgeCfm
func CgoOnRspQryHedgeCfm(
	this unsafe.Pointer,
	pHedgeCfm *C.struct_CThostFtdcHedgeCfmField,
	pRspInfo *C.struct_CThostFtdcRspInfoField,
	nRequestID C.int,
	bIsLast C.bool,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRspQryHedgeCfm called",
		slog.Any("this", this),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnRspQryHedgeCfm(
		(*future.CThostFtdcHedgeCfmField)(unsafe.Pointer(pHedgeCfm)),
		(*future.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
		int(nRequestID),
		bool(bIsLast),
	)
}

//export CgoOnRtnHedgeCfm
func CgoOnRtnHedgeCfm(
	this unsafe.Pointer,
	pHedgeCfm *C.struct_CThostFtdcHedgeCfmField,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRtnHedgeCfm called",
		slog.Any("this", this),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnRtnHedgeCfm(
		(*future.CThostFtdcHedgeCfmField)(unsafe.Pointer(pHedgeCfm)),
	)
}

//export CgoOnErrRtnHedgeCfm
func CgoOnErrRtnHedgeCfm(
	this unsafe.Pointer,
	pInputHedgeCfm *C.struct_CThostFtdcInputHedgeCfmField,
	pRspInfo *C.struct_CThostFtdcRspInfoField,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnErrRtnHedgeCfm called",
		slog.Any("this", this),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnErrRtnHedgeCfm(
		(*future.CThostFtdcInputHedgeCfmField)(unsafe.Pointer(pInputHedgeCfm)),
		(*future.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
	)
}

//export CgoOnErrRtnHedgeCfmAction
func CgoOnErrRtnHedgeCfmAction(
	this unsafe.Pointer,
	pHedgeCfmAction *C.struct_CThostFtdcHedgeCfmActionField,
	pRspInfo *C.struct_CThostFtdcRspInfoField,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnErrRtnHedgeCfmAction called",
		slog.Any("this", this),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnErrRtnHedgeCfmAction(
		(*future.CThostFtdcHedgeCfmActionField)(unsafe.Pointer(pHedgeCfmAction)),
		(*future.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
	)
}
