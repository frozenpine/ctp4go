package v6_7_13

/*
#cgo CFLAGS: -I. -I${SRCDIR} -I${SRCDIR}/../../dependencies/future/v6.7.13/
#cgo LDFLAGS: -ldl

#include "td_spi_helper.h"
*/
import "C"
import (
	"context"
	"log/slog"
	"runtime"
	"unsafe"

	"github.com/frozenpine/ctp4go/thost"
)

var (
	spiCVtablePtr *C.CThostFtdcTraderSpiVTable
)

func init() {
	spiCVtablePtr = (*C.CThostFtdcTraderSpiVTable)(C.malloc(C.sizeof_CThostFtdcTraderSpiVTable))

	spiCVtablePtr.CThostFtdcTraderSpi_OnFrontConnected = (C.OnFrontConnected)(
		unsafe.Pointer(C.COnFrontConnected))

	spiCVtablePtr.CThostFtdcTraderSpi_OnFrontDisconnected = (C.OnFrontDisconnected)(
		unsafe.Pointer(C.COnFrontDisconnected))

	spiCVtablePtr.CThostFtdcTraderSpi_OnHeartBeatWarning = (C.OnHeartBeatWarning)(
		unsafe.Pointer(C.COnHeartBeatWarning))

	spiCVtablePtr.CThostFtdcTraderSpi_OnRspAuthenticate = (C.OnRspAuthenticate)(
		unsafe.Pointer(C.COnRspAuthenticate))

	spiCVtablePtr.CThostFtdcTraderSpi_OnRtnPrivateSeqNo = (C.OnRtnPrivateSeqNo)(
		unsafe.Pointer(C.COnRtnPrivateSeqNo))

	spiCVtablePtr.CThostFtdcTraderSpi_OnRspUserLogin = (C.OnRspUserLogin)(
		unsafe.Pointer(C.COnRspUserLogin))

	spiCVtablePtr.CThostFtdcTraderSpi_OnRspUserLogout = (C.OnRspUserLogout)(
		unsafe.Pointer(C.COnRspUserLogout))

	spiCVtablePtr.CThostFtdcTraderSpi_OnRspUserPasswordUpdate = (C.OnRspUserPasswordUpdate)(
		unsafe.Pointer(C.COnRspUserPasswordUpdate))

	spiCVtablePtr.CThostFtdcTraderSpi_OnRspTradingAccountPasswordUpdate = (C.OnRspTradingAccountPasswordUpdate)(
		unsafe.Pointer(C.COnRspTradingAccountPasswordUpdate))

	spiCVtablePtr.CThostFtdcTraderSpi_OnRspUserAuthMethod = (C.OnRspUserAuthMethod)(
		unsafe.Pointer(C.COnRspUserAuthMethod))

	spiCVtablePtr.CThostFtdcTraderSpi_OnRspGenUserCaptcha = (C.OnRspGenUserCaptcha)(
		unsafe.Pointer(C.COnRspGenUserCaptcha))

	spiCVtablePtr.CThostFtdcTraderSpi_OnRspGenUserText = (C.OnRspGenUserText)(
		unsafe.Pointer(C.COnRspGenUserText))

	spiCVtablePtr.CThostFtdcTraderSpi_OnRspOrderInsert = (C.OnRspOrderInsert)(
		unsafe.Pointer(C.COnRspOrderInsert))

	spiCVtablePtr.CThostFtdcTraderSpi_OnRspParkedOrderInsert = (C.OnRspParkedOrderInsert)(
		unsafe.Pointer(C.COnRspParkedOrderInsert))

	spiCVtablePtr.CThostFtdcTraderSpi_OnRspParkedOrderAction = (C.OnRspParkedOrderAction)(
		unsafe.Pointer(C.COnRspParkedOrderAction))

	spiCVtablePtr.CThostFtdcTraderSpi_OnRspOrderAction = (C.OnRspOrderAction)(
		unsafe.Pointer(C.COnRspOrderAction))

	spiCVtablePtr.CThostFtdcTraderSpi_OnRspQryMaxOrderVolume = (C.OnRspQryMaxOrderVolume)(
		unsafe.Pointer(C.COnRspQryMaxOrderVolume))

	spiCVtablePtr.CThostFtdcTraderSpi_OnRspSettlementInfoConfirm = (C.OnRspSettlementInfoConfirm)(
		unsafe.Pointer(C.COnRspSettlementInfoConfirm))

	spiCVtablePtr.CThostFtdcTraderSpi_OnRspRemoveParkedOrder = (C.OnRspRemoveParkedOrder)(
		unsafe.Pointer(C.COnRspRemoveParkedOrder))

	spiCVtablePtr.CThostFtdcTraderSpi_OnRspRemoveParkedOrderAction = (C.OnRspRemoveParkedOrderAction)(
		unsafe.Pointer(C.COnRspRemoveParkedOrderAction))

	spiCVtablePtr.CThostFtdcTraderSpi_OnRspExecOrderInsert = (C.OnRspExecOrderInsert)(
		unsafe.Pointer(C.COnRspExecOrderInsert))

	spiCVtablePtr.CThostFtdcTraderSpi_OnRspExecOrderAction = (C.OnRspExecOrderAction)(
		unsafe.Pointer(C.COnRspExecOrderAction))

	spiCVtablePtr.CThostFtdcTraderSpi_OnRspForQuoteInsert = (C.OnRspForQuoteInsert)(
		unsafe.Pointer(C.COnRspForQuoteInsert))

	spiCVtablePtr.CThostFtdcTraderSpi_OnRspQuoteInsert = (C.OnRspQuoteInsert)(
		unsafe.Pointer(C.COnRspQuoteInsert))

	spiCVtablePtr.CThostFtdcTraderSpi_OnRspQuoteAction = (C.OnRspQuoteAction)(
		unsafe.Pointer(C.COnRspQuoteAction))

	spiCVtablePtr.CThostFtdcTraderSpi_OnRspBatchOrderAction = (C.OnRspBatchOrderAction)(
		unsafe.Pointer(C.COnRspBatchOrderAction))

	spiCVtablePtr.CThostFtdcTraderSpi_OnRspOptionSelfCloseInsert = (C.OnRspOptionSelfCloseInsert)(
		unsafe.Pointer(C.COnRspOptionSelfCloseInsert))

	spiCVtablePtr.CThostFtdcTraderSpi_OnRspOptionSelfCloseAction = (C.OnRspOptionSelfCloseAction)(
		unsafe.Pointer(C.COnRspOptionSelfCloseAction))

	spiCVtablePtr.CThostFtdcTraderSpi_OnRspCombActionInsert = (C.OnRspCombActionInsert)(
		unsafe.Pointer(C.COnRspCombActionInsert))

	spiCVtablePtr.CThostFtdcTraderSpi_OnRspQryOrder = (C.OnRspQryOrder)(
		unsafe.Pointer(C.COnRspQryOrder))

	spiCVtablePtr.CThostFtdcTraderSpi_OnRspQryTrade = (C.OnRspQryTrade)(
		unsafe.Pointer(C.COnRspQryTrade))

	spiCVtablePtr.CThostFtdcTraderSpi_OnRspQryInvestorPosition = (C.OnRspQryInvestorPosition)(
		unsafe.Pointer(C.COnRspQryInvestorPosition))

	spiCVtablePtr.CThostFtdcTraderSpi_OnRspQryTradingAccount = (C.OnRspQryTradingAccount)(
		unsafe.Pointer(C.COnRspQryTradingAccount))

	spiCVtablePtr.CThostFtdcTraderSpi_OnRspQryInvestor = (C.OnRspQryInvestor)(
		unsafe.Pointer(C.COnRspQryInvestor))

	spiCVtablePtr.CThostFtdcTraderSpi_OnRspQryTradingCode = (C.OnRspQryTradingCode)(
		unsafe.Pointer(C.COnRspQryTradingCode))

	spiCVtablePtr.CThostFtdcTraderSpi_OnRspQryInstrumentMarginRate = (C.OnRspQryInstrumentMarginRate)(
		unsafe.Pointer(C.COnRspQryInstrumentMarginRate))

	spiCVtablePtr.CThostFtdcTraderSpi_OnRspQryInstrumentCommissionRate = (C.OnRspQryInstrumentCommissionRate)(
		unsafe.Pointer(C.COnRspQryInstrumentCommissionRate))

	spiCVtablePtr.CThostFtdcTraderSpi_OnRspQryUserSession = (C.OnRspQryUserSession)(
		unsafe.Pointer(C.COnRspQryUserSession))

	spiCVtablePtr.CThostFtdcTraderSpi_OnRspQryExchange = (C.OnRspQryExchange)(
		unsafe.Pointer(C.COnRspQryExchange))

	spiCVtablePtr.CThostFtdcTraderSpi_OnRspQryProduct = (C.OnRspQryProduct)(
		unsafe.Pointer(C.COnRspQryProduct))

	spiCVtablePtr.CThostFtdcTraderSpi_OnRspQryInstrument = (C.OnRspQryInstrument)(
		unsafe.Pointer(C.COnRspQryInstrument))

	spiCVtablePtr.CThostFtdcTraderSpi_OnRspQryDepthMarketData = (C.OnRspQryDepthMarketData)(
		unsafe.Pointer(C.COnRspQryDepthMarketData))

	spiCVtablePtr.CThostFtdcTraderSpi_OnRspQryTraderOffer = (C.OnRspQryTraderOffer)(
		unsafe.Pointer(C.COnRspQryTraderOffer))

	spiCVtablePtr.CThostFtdcTraderSpi_OnRspQrySettlementInfo = (C.OnRspQrySettlementInfo)(
		unsafe.Pointer(C.COnRspQrySettlementInfo))

	spiCVtablePtr.CThostFtdcTraderSpi_OnRspQryTransferBank = (C.OnRspQryTransferBank)(
		unsafe.Pointer(C.COnRspQryTransferBank))

	spiCVtablePtr.CThostFtdcTraderSpi_OnRspQryInvestorPositionDetail = (C.OnRspQryInvestorPositionDetail)(
		unsafe.Pointer(C.COnRspQryInvestorPositionDetail))

	spiCVtablePtr.CThostFtdcTraderSpi_OnRspQryNotice = (C.OnRspQryNotice)(
		unsafe.Pointer(C.COnRspQryNotice))

	spiCVtablePtr.CThostFtdcTraderSpi_OnRspQrySettlementInfoConfirm = (C.OnRspQrySettlementInfoConfirm)(
		unsafe.Pointer(C.COnRspQrySettlementInfoConfirm))

	spiCVtablePtr.CThostFtdcTraderSpi_OnRspQryInvestorPositionCombineDetail = (C.OnRspQryInvestorPositionCombineDetail)(
		unsafe.Pointer(C.COnRspQryInvestorPositionCombineDetail))

	spiCVtablePtr.CThostFtdcTraderSpi_OnRspQryCFMMCTradingAccountKey = (C.OnRspQryCFMMCTradingAccountKey)(
		unsafe.Pointer(C.COnRspQryCFMMCTradingAccountKey))

	spiCVtablePtr.CThostFtdcTraderSpi_OnRspQryEWarrantOffset = (C.OnRspQryEWarrantOffset)(
		unsafe.Pointer(C.COnRspQryEWarrantOffset))

	spiCVtablePtr.CThostFtdcTraderSpi_OnRspQryInvestorProductGroupMargin = (C.OnRspQryInvestorProductGroupMargin)(
		unsafe.Pointer(C.COnRspQryInvestorProductGroupMargin))

	spiCVtablePtr.CThostFtdcTraderSpi_OnRspQryExchangeMarginRate = (C.OnRspQryExchangeMarginRate)(
		unsafe.Pointer(C.COnRspQryExchangeMarginRate))

	spiCVtablePtr.CThostFtdcTraderSpi_OnRspQryExchangeMarginRateAdjust = (C.OnRspQryExchangeMarginRateAdjust)(
		unsafe.Pointer(C.COnRspQryExchangeMarginRateAdjust))

	spiCVtablePtr.CThostFtdcTraderSpi_OnRspQryExchangeRate = (C.OnRspQryExchangeRate)(
		unsafe.Pointer(C.COnRspQryExchangeRate))

	spiCVtablePtr.CThostFtdcTraderSpi_OnRspQrySecAgentACIDMap = (C.OnRspQrySecAgentACIDMap)(
		unsafe.Pointer(C.COnRspQrySecAgentACIDMap))

	spiCVtablePtr.CThostFtdcTraderSpi_OnRspQryProductExchRate = (C.OnRspQryProductExchRate)(
		unsafe.Pointer(C.COnRspQryProductExchRate))

	spiCVtablePtr.CThostFtdcTraderSpi_OnRspQryProductGroup = (C.OnRspQryProductGroup)(
		unsafe.Pointer(C.COnRspQryProductGroup))

	spiCVtablePtr.CThostFtdcTraderSpi_OnRspQryMMInstrumentCommissionRate = (C.OnRspQryMMInstrumentCommissionRate)(
		unsafe.Pointer(C.COnRspQryMMInstrumentCommissionRate))

	spiCVtablePtr.CThostFtdcTraderSpi_OnRspQryMMOptionInstrCommRate = (C.OnRspQryMMOptionInstrCommRate)(
		unsafe.Pointer(C.COnRspQryMMOptionInstrCommRate))

	spiCVtablePtr.CThostFtdcTraderSpi_OnRspQryInstrumentOrderCommRate = (C.OnRspQryInstrumentOrderCommRate)(
		unsafe.Pointer(C.COnRspQryInstrumentOrderCommRate))

	spiCVtablePtr.CThostFtdcTraderSpi_OnRspQrySecAgentTradingAccount = (C.OnRspQrySecAgentTradingAccount)(
		unsafe.Pointer(C.COnRspQrySecAgentTradingAccount))

	spiCVtablePtr.CThostFtdcTraderSpi_OnRspQrySecAgentCheckMode = (C.OnRspQrySecAgentCheckMode)(
		unsafe.Pointer(C.COnRspQrySecAgentCheckMode))

	spiCVtablePtr.CThostFtdcTraderSpi_OnRspQrySecAgentTradeInfo = (C.OnRspQrySecAgentTradeInfo)(
		unsafe.Pointer(C.COnRspQrySecAgentTradeInfo))

	spiCVtablePtr.CThostFtdcTraderSpi_OnRspQryOptionInstrTradeCost = (C.OnRspQryOptionInstrTradeCost)(
		unsafe.Pointer(C.COnRspQryOptionInstrTradeCost))

	spiCVtablePtr.CThostFtdcTraderSpi_OnRspQryOptionInstrCommRate = (C.OnRspQryOptionInstrCommRate)(
		unsafe.Pointer(C.COnRspQryOptionInstrCommRate))

	spiCVtablePtr.CThostFtdcTraderSpi_OnRspQryExecOrder = (C.OnRspQryExecOrder)(
		unsafe.Pointer(C.COnRspQryExecOrder))

	spiCVtablePtr.CThostFtdcTraderSpi_OnRspQryForQuote = (C.OnRspQryForQuote)(
		unsafe.Pointer(C.COnRspQryForQuote))

	spiCVtablePtr.CThostFtdcTraderSpi_OnRspQryQuote = (C.OnRspQryQuote)(
		unsafe.Pointer(C.COnRspQryQuote))

	spiCVtablePtr.CThostFtdcTraderSpi_OnRspQryOptionSelfClose = (C.OnRspQryOptionSelfClose)(
		unsafe.Pointer(C.COnRspQryOptionSelfClose))

	spiCVtablePtr.CThostFtdcTraderSpi_OnRspQryInvestUnit = (C.OnRspQryInvestUnit)(
		unsafe.Pointer(C.COnRspQryInvestUnit))

	spiCVtablePtr.CThostFtdcTraderSpi_OnRspQryCombInstrumentGuard = (C.OnRspQryCombInstrumentGuard)(
		unsafe.Pointer(C.COnRspQryCombInstrumentGuard))

	spiCVtablePtr.CThostFtdcTraderSpi_OnRspQryCombAction = (C.OnRspQryCombAction)(
		unsafe.Pointer(C.COnRspQryCombAction))

	spiCVtablePtr.CThostFtdcTraderSpi_OnRspQryTransferSerial = (C.OnRspQryTransferSerial)(
		unsafe.Pointer(C.COnRspQryTransferSerial))

	spiCVtablePtr.CThostFtdcTraderSpi_OnRspQryAccountregister = (C.OnRspQryAccountregister)(
		unsafe.Pointer(C.COnRspQryAccountregister))

	spiCVtablePtr.CThostFtdcTraderSpi_OnRspError = (C.OnRspError)(
		unsafe.Pointer(C.COnRspError))

	spiCVtablePtr.CThostFtdcTraderSpi_OnRtnOrder = (C.OnRtnOrder)(
		unsafe.Pointer(C.COnRtnOrder))

	spiCVtablePtr.CThostFtdcTraderSpi_OnRtnTrade = (C.OnRtnTrade)(
		unsafe.Pointer(C.COnRtnTrade))

	spiCVtablePtr.CThostFtdcTraderSpi_OnErrRtnOrderInsert = (C.OnErrRtnOrderInsert)(
		unsafe.Pointer(C.COnErrRtnOrderInsert))

	spiCVtablePtr.CThostFtdcTraderSpi_OnErrRtnOrderAction = (C.OnErrRtnOrderAction)(
		unsafe.Pointer(C.COnErrRtnOrderAction))

	spiCVtablePtr.CThostFtdcTraderSpi_OnRtnInstrumentStatus = (C.OnRtnInstrumentStatus)(
		unsafe.Pointer(C.COnRtnInstrumentStatus))

	spiCVtablePtr.CThostFtdcTraderSpi_OnRtnBulletin = (C.OnRtnBulletin)(
		unsafe.Pointer(C.COnRtnBulletin))

	spiCVtablePtr.CThostFtdcTraderSpi_OnRtnTradingNotice = (C.OnRtnTradingNotice)(
		unsafe.Pointer(C.COnRtnTradingNotice))

	spiCVtablePtr.CThostFtdcTraderSpi_OnRtnErrorConditionalOrder = (C.OnRtnErrorConditionalOrder)(
		unsafe.Pointer(C.COnRtnErrorConditionalOrder))

	spiCVtablePtr.CThostFtdcTraderSpi_OnRtnExecOrder = (C.OnRtnExecOrder)(
		unsafe.Pointer(C.COnRtnExecOrder))

	spiCVtablePtr.CThostFtdcTraderSpi_OnErrRtnExecOrderInsert = (C.OnErrRtnExecOrderInsert)(
		unsafe.Pointer(C.COnErrRtnExecOrderInsert))

	spiCVtablePtr.CThostFtdcTraderSpi_OnErrRtnExecOrderAction = (C.OnErrRtnExecOrderAction)(
		unsafe.Pointer(C.COnErrRtnExecOrderAction))

	spiCVtablePtr.CThostFtdcTraderSpi_OnErrRtnForQuoteInsert = (C.OnErrRtnForQuoteInsert)(
		unsafe.Pointer(C.COnErrRtnForQuoteInsert))

	spiCVtablePtr.CThostFtdcTraderSpi_OnRtnQuote = (C.OnRtnQuote)(
		unsafe.Pointer(C.COnRtnQuote))

	spiCVtablePtr.CThostFtdcTraderSpi_OnErrRtnQuoteInsert = (C.OnErrRtnQuoteInsert)(
		unsafe.Pointer(C.COnErrRtnQuoteInsert))

	spiCVtablePtr.CThostFtdcTraderSpi_OnErrRtnQuoteAction = (C.OnErrRtnQuoteAction)(
		unsafe.Pointer(C.COnErrRtnQuoteAction))

	spiCVtablePtr.CThostFtdcTraderSpi_OnRtnForQuoteRsp = (C.OnRtnForQuoteRsp)(
		unsafe.Pointer(C.COnRtnForQuoteRsp))

	spiCVtablePtr.CThostFtdcTraderSpi_OnRtnCFMMCTradingAccountToken = (C.OnRtnCFMMCTradingAccountToken)(
		unsafe.Pointer(C.COnRtnCFMMCTradingAccountToken))

	spiCVtablePtr.CThostFtdcTraderSpi_OnErrRtnBatchOrderAction = (C.OnErrRtnBatchOrderAction)(
		unsafe.Pointer(C.COnErrRtnBatchOrderAction))

	spiCVtablePtr.CThostFtdcTraderSpi_OnRtnOptionSelfClose = (C.OnRtnOptionSelfClose)(
		unsafe.Pointer(C.COnRtnOptionSelfClose))

	spiCVtablePtr.CThostFtdcTraderSpi_OnErrRtnOptionSelfCloseInsert = (C.OnErrRtnOptionSelfCloseInsert)(
		unsafe.Pointer(C.COnErrRtnOptionSelfCloseInsert))

	spiCVtablePtr.CThostFtdcTraderSpi_OnErrRtnOptionSelfCloseAction = (C.OnErrRtnOptionSelfCloseAction)(
		unsafe.Pointer(C.COnErrRtnOptionSelfCloseAction))

	spiCVtablePtr.CThostFtdcTraderSpi_OnRtnCombAction = (C.OnRtnCombAction)(
		unsafe.Pointer(C.COnRtnCombAction))

	spiCVtablePtr.CThostFtdcTraderSpi_OnErrRtnCombActionInsert = (C.OnErrRtnCombActionInsert)(
		unsafe.Pointer(C.COnErrRtnCombActionInsert))

	spiCVtablePtr.CThostFtdcTraderSpi_OnRspQryContractBank = (C.OnRspQryContractBank)(
		unsafe.Pointer(C.COnRspQryContractBank))

	spiCVtablePtr.CThostFtdcTraderSpi_OnRspQryParkedOrder = (C.OnRspQryParkedOrder)(
		unsafe.Pointer(C.COnRspQryParkedOrder))

	spiCVtablePtr.CThostFtdcTraderSpi_OnRspQryParkedOrderAction = (C.OnRspQryParkedOrderAction)(
		unsafe.Pointer(C.COnRspQryParkedOrderAction))

	spiCVtablePtr.CThostFtdcTraderSpi_OnRspQryTradingNotice = (C.OnRspQryTradingNotice)(
		unsafe.Pointer(C.COnRspQryTradingNotice))

	spiCVtablePtr.CThostFtdcTraderSpi_OnRspQryBrokerTradingParams = (C.OnRspQryBrokerTradingParams)(
		unsafe.Pointer(C.COnRspQryBrokerTradingParams))

	spiCVtablePtr.CThostFtdcTraderSpi_OnRspQryBrokerTradingAlgos = (C.OnRspQryBrokerTradingAlgos)(
		unsafe.Pointer(C.COnRspQryBrokerTradingAlgos))

	spiCVtablePtr.CThostFtdcTraderSpi_OnRspQueryCFMMCTradingAccountToken = (C.OnRspQueryCFMMCTradingAccountToken)(
		unsafe.Pointer(C.COnRspQueryCFMMCTradingAccountToken))

	spiCVtablePtr.CThostFtdcTraderSpi_OnRtnFromBankToFutureByBank = (C.OnRtnFromBankToFutureByBank)(
		unsafe.Pointer(C.COnRtnFromBankToFutureByBank))

	spiCVtablePtr.CThostFtdcTraderSpi_OnRtnFromFutureToBankByBank = (C.OnRtnFromFutureToBankByBank)(
		unsafe.Pointer(C.COnRtnFromFutureToBankByBank))

	spiCVtablePtr.CThostFtdcTraderSpi_OnRtnRepealFromBankToFutureByBank = (C.OnRtnRepealFromBankToFutureByBank)(
		unsafe.Pointer(C.COnRtnRepealFromBankToFutureByBank))

	spiCVtablePtr.CThostFtdcTraderSpi_OnRtnRepealFromFutureToBankByBank = (C.OnRtnRepealFromFutureToBankByBank)(
		unsafe.Pointer(C.COnRtnRepealFromFutureToBankByBank))

	spiCVtablePtr.CThostFtdcTraderSpi_OnRtnFromBankToFutureByFuture = (C.OnRtnFromBankToFutureByFuture)(
		unsafe.Pointer(C.COnRtnFromBankToFutureByFuture))

	spiCVtablePtr.CThostFtdcTraderSpi_OnRtnFromFutureToBankByFuture = (C.OnRtnFromFutureToBankByFuture)(
		unsafe.Pointer(C.COnRtnFromFutureToBankByFuture))

	spiCVtablePtr.CThostFtdcTraderSpi_OnRtnRepealFromBankToFutureByFutureManual = (C.OnRtnRepealFromBankToFutureByFutureManual)(
		unsafe.Pointer(C.COnRtnRepealFromBankToFutureByFutureManual))

	spiCVtablePtr.CThostFtdcTraderSpi_OnRtnRepealFromFutureToBankByFutureManual = (C.OnRtnRepealFromFutureToBankByFutureManual)(
		unsafe.Pointer(C.COnRtnRepealFromFutureToBankByFutureManual))

	spiCVtablePtr.CThostFtdcTraderSpi_OnRtnQueryBankBalanceByFuture = (C.OnRtnQueryBankBalanceByFuture)(
		unsafe.Pointer(C.COnRtnQueryBankBalanceByFuture))

	spiCVtablePtr.CThostFtdcTraderSpi_OnErrRtnBankToFutureByFuture = (C.OnErrRtnBankToFutureByFuture)(
		unsafe.Pointer(C.COnErrRtnBankToFutureByFuture))

	spiCVtablePtr.CThostFtdcTraderSpi_OnErrRtnFutureToBankByFuture = (C.OnErrRtnFutureToBankByFuture)(
		unsafe.Pointer(C.COnErrRtnFutureToBankByFuture))

	spiCVtablePtr.CThostFtdcTraderSpi_OnErrRtnRepealBankToFutureByFutureManual = (C.OnErrRtnRepealBankToFutureByFutureManual)(
		unsafe.Pointer(C.COnErrRtnRepealBankToFutureByFutureManual))

	spiCVtablePtr.CThostFtdcTraderSpi_OnErrRtnRepealFutureToBankByFutureManual = (C.OnErrRtnRepealFutureToBankByFutureManual)(
		unsafe.Pointer(C.COnErrRtnRepealFutureToBankByFutureManual))

	spiCVtablePtr.CThostFtdcTraderSpi_OnErrRtnQueryBankBalanceByFuture = (C.OnErrRtnQueryBankBalanceByFuture)(
		unsafe.Pointer(C.COnErrRtnQueryBankBalanceByFuture))

	spiCVtablePtr.CThostFtdcTraderSpi_OnRtnRepealFromBankToFutureByFuture = (C.OnRtnRepealFromBankToFutureByFuture)(
		unsafe.Pointer(C.COnRtnRepealFromBankToFutureByFuture))

	spiCVtablePtr.CThostFtdcTraderSpi_OnRtnRepealFromFutureToBankByFuture = (C.OnRtnRepealFromFutureToBankByFuture)(
		unsafe.Pointer(C.COnRtnRepealFromFutureToBankByFuture))

	spiCVtablePtr.CThostFtdcTraderSpi_OnRspFromBankToFutureByFuture = (C.OnRspFromBankToFutureByFuture)(
		unsafe.Pointer(C.COnRspFromBankToFutureByFuture))

	spiCVtablePtr.CThostFtdcTraderSpi_OnRspFromFutureToBankByFuture = (C.OnRspFromFutureToBankByFuture)(
		unsafe.Pointer(C.COnRspFromFutureToBankByFuture))

	spiCVtablePtr.CThostFtdcTraderSpi_OnRspQueryBankAccountMoneyByFuture = (C.OnRspQueryBankAccountMoneyByFuture)(
		unsafe.Pointer(C.COnRspQueryBankAccountMoneyByFuture))

	spiCVtablePtr.CThostFtdcTraderSpi_OnRtnOpenAccountByBank = (C.OnRtnOpenAccountByBank)(
		unsafe.Pointer(C.COnRtnOpenAccountByBank))

	spiCVtablePtr.CThostFtdcTraderSpi_OnRtnCancelAccountByBank = (C.OnRtnCancelAccountByBank)(
		unsafe.Pointer(C.COnRtnCancelAccountByBank))

	spiCVtablePtr.CThostFtdcTraderSpi_OnRtnChangeAccountByBank = (C.OnRtnChangeAccountByBank)(
		unsafe.Pointer(C.COnRtnChangeAccountByBank))

	spiCVtablePtr.CThostFtdcTraderSpi_OnRspQryClassifiedInstrument = (C.OnRspQryClassifiedInstrument)(
		unsafe.Pointer(C.COnRspQryClassifiedInstrument))

	spiCVtablePtr.CThostFtdcTraderSpi_OnRspQryCombPromotionParam = (C.OnRspQryCombPromotionParam)(
		unsafe.Pointer(C.COnRspQryCombPromotionParam))

	spiCVtablePtr.CThostFtdcTraderSpi_OnRspQryRiskSettleInvstPosition = (C.OnRspQryRiskSettleInvstPosition)(
		unsafe.Pointer(C.COnRspQryRiskSettleInvstPosition))

	spiCVtablePtr.CThostFtdcTraderSpi_OnRspQryRiskSettleProductStatus = (C.OnRspQryRiskSettleProductStatus)(
		unsafe.Pointer(C.COnRspQryRiskSettleProductStatus))

	spiCVtablePtr.CThostFtdcTraderSpi_OnRspQrySPBMFutureParameter = (C.OnRspQrySPBMFutureParameter)(
		unsafe.Pointer(C.COnRspQrySPBMFutureParameter))

	spiCVtablePtr.CThostFtdcTraderSpi_OnRspQrySPBMOptionParameter = (C.OnRspQrySPBMOptionParameter)(
		unsafe.Pointer(C.COnRspQrySPBMOptionParameter))

	spiCVtablePtr.CThostFtdcTraderSpi_OnRspQrySPBMIntraParameter = (C.OnRspQrySPBMIntraParameter)(
		unsafe.Pointer(C.COnRspQrySPBMIntraParameter))

	spiCVtablePtr.CThostFtdcTraderSpi_OnRspQrySPBMInterParameter = (C.OnRspQrySPBMInterParameter)(
		unsafe.Pointer(C.COnRspQrySPBMInterParameter))

	spiCVtablePtr.CThostFtdcTraderSpi_OnRspQrySPBMPortfDefinition = (C.OnRspQrySPBMPortfDefinition)(
		unsafe.Pointer(C.COnRspQrySPBMPortfDefinition))

	spiCVtablePtr.CThostFtdcTraderSpi_OnRspQrySPBMInvestorPortfDef = (C.OnRspQrySPBMInvestorPortfDef)(
		unsafe.Pointer(C.COnRspQrySPBMInvestorPortfDef))

	spiCVtablePtr.CThostFtdcTraderSpi_OnRspQryInvestorPortfMarginRatio = (C.OnRspQryInvestorPortfMarginRatio)(
		unsafe.Pointer(C.COnRspQryInvestorPortfMarginRatio))

	spiCVtablePtr.CThostFtdcTraderSpi_OnRspQryInvestorProdSPBMDetail = (C.OnRspQryInvestorProdSPBMDetail)(
		unsafe.Pointer(C.COnRspQryInvestorProdSPBMDetail))

	spiCVtablePtr.CThostFtdcTraderSpi_OnRspQryInvestorCommoditySPMMMargin = (C.OnRspQryInvestorCommoditySPMMMargin)(
		unsafe.Pointer(C.COnRspQryInvestorCommoditySPMMMargin))

	spiCVtablePtr.CThostFtdcTraderSpi_OnRspQryInvestorCommodityGroupSPMMMargin = (C.OnRspQryInvestorCommodityGroupSPMMMargin)(
		unsafe.Pointer(C.COnRspQryInvestorCommodityGroupSPMMMargin))

	spiCVtablePtr.CThostFtdcTraderSpi_OnRspQrySPMMInstParam = (C.OnRspQrySPMMInstParam)(
		unsafe.Pointer(C.COnRspQrySPMMInstParam))

	spiCVtablePtr.CThostFtdcTraderSpi_OnRspQrySPMMProductParam = (C.OnRspQrySPMMProductParam)(
		unsafe.Pointer(C.COnRspQrySPMMProductParam))

	spiCVtablePtr.CThostFtdcTraderSpi_OnRspQrySPBMAddOnInterParameter = (C.OnRspQrySPBMAddOnInterParameter)(
		unsafe.Pointer(C.COnRspQrySPBMAddOnInterParameter))

	spiCVtablePtr.CThostFtdcTraderSpi_OnRspQryRCAMSCombProductInfo = (C.OnRspQryRCAMSCombProductInfo)(
		unsafe.Pointer(C.COnRspQryRCAMSCombProductInfo))

	spiCVtablePtr.CThostFtdcTraderSpi_OnRspQryRCAMSInstrParameter = (C.OnRspQryRCAMSInstrParameter)(
		unsafe.Pointer(C.COnRspQryRCAMSInstrParameter))

	spiCVtablePtr.CThostFtdcTraderSpi_OnRspQryRCAMSIntraParameter = (C.OnRspQryRCAMSIntraParameter)(
		unsafe.Pointer(C.COnRspQryRCAMSIntraParameter))

	spiCVtablePtr.CThostFtdcTraderSpi_OnRspQryRCAMSInterParameter = (C.OnRspQryRCAMSInterParameter)(
		unsafe.Pointer(C.COnRspQryRCAMSInterParameter))

	spiCVtablePtr.CThostFtdcTraderSpi_OnRspQryRCAMSShortOptAdjustParam = (C.OnRspQryRCAMSShortOptAdjustParam)(
		unsafe.Pointer(C.COnRspQryRCAMSShortOptAdjustParam))

	spiCVtablePtr.CThostFtdcTraderSpi_OnRspQryRCAMSInvestorCombPosition = (C.OnRspQryRCAMSInvestorCombPosition)(
		unsafe.Pointer(C.COnRspQryRCAMSInvestorCombPosition))

	spiCVtablePtr.CThostFtdcTraderSpi_OnRspQryInvestorProdRCAMSMargin = (C.OnRspQryInvestorProdRCAMSMargin)(
		unsafe.Pointer(C.COnRspQryInvestorProdRCAMSMargin))

	spiCVtablePtr.CThostFtdcTraderSpi_OnRspQryRULEInstrParameter = (C.OnRspQryRULEInstrParameter)(
		unsafe.Pointer(C.COnRspQryRULEInstrParameter))

	spiCVtablePtr.CThostFtdcTraderSpi_OnRspQryRULEIntraParameter = (C.OnRspQryRULEIntraParameter)(
		unsafe.Pointer(C.COnRspQryRULEIntraParameter))

	spiCVtablePtr.CThostFtdcTraderSpi_OnRspQryRULEInterParameter = (C.OnRspQryRULEInterParameter)(
		unsafe.Pointer(C.COnRspQryRULEInterParameter))

	spiCVtablePtr.CThostFtdcTraderSpi_OnRspQryInvestorProdRULEMargin = (C.OnRspQryInvestorProdRULEMargin)(
		unsafe.Pointer(C.COnRspQryInvestorProdRULEMargin))

	spiCVtablePtr.CThostFtdcTraderSpi_OnRspQryInvestorPortfSetting = (C.OnRspQryInvestorPortfSetting)(
		unsafe.Pointer(C.COnRspQryInvestorPortfSetting))

	spiCVtablePtr.CThostFtdcTraderSpi_OnRspQryInvestorInfoCommRec = (C.OnRspQryInvestorInfoCommRec)(
		unsafe.Pointer(C.COnRspQryInvestorInfoCommRec))

	spiCVtablePtr.CThostFtdcTraderSpi_OnRspQryCombLeg = (C.OnRspQryCombLeg)(
		unsafe.Pointer(C.COnRspQryCombLeg))

	spiCVtablePtr.CThostFtdcTraderSpi_OnRspOffsetSetting = (C.OnRspOffsetSetting)(
		unsafe.Pointer(C.COnRspOffsetSetting))

	spiCVtablePtr.CThostFtdcTraderSpi_OnRspCancelOffsetSetting = (C.OnRspCancelOffsetSetting)(
		unsafe.Pointer(C.COnRspCancelOffsetSetting))

	spiCVtablePtr.CThostFtdcTraderSpi_OnRtnOffsetSetting = (C.OnRtnOffsetSetting)(
		unsafe.Pointer(C.COnRtnOffsetSetting))

	spiCVtablePtr.CThostFtdcTraderSpi_OnErrRtnOffsetSetting = (C.OnErrRtnOffsetSetting)(
		unsafe.Pointer(C.COnErrRtnOffsetSetting))

	spiCVtablePtr.CThostFtdcTraderSpi_OnErrRtnCancelOffsetSetting = (C.OnErrRtnCancelOffsetSetting)(
		unsafe.Pointer(C.COnErrRtnCancelOffsetSetting))

	spiCVtablePtr.CThostFtdcTraderSpi_OnRspQryOffsetSetting = (C.OnRspQryOffsetSetting)(
		unsafe.Pointer(C.COnRspQryOffsetSetting))

	spiCVtablePtr.CThostFtdcTraderSpi_OnRspGenSMSCode = (C.OnRspGenSMSCode)(
		unsafe.Pointer(C.COnRspGenSMSCode))

	spiCVtablePtr.CThostFtdcTraderSpi_OnRspSpdApply = (C.OnRspSpdApply)(
		unsafe.Pointer(C.COnRspSpdApply))

	spiCVtablePtr.CThostFtdcTraderSpi_OnRspSpdApplyAction = (C.OnRspSpdApplyAction)(
		unsafe.Pointer(C.COnRspSpdApplyAction))

	spiCVtablePtr.CThostFtdcTraderSpi_OnRspQrySpdApply = (C.OnRspQrySpdApply)(
		unsafe.Pointer(C.COnRspQrySpdApply))

	spiCVtablePtr.CThostFtdcTraderSpi_OnRtnSpdApply = (C.OnRtnSpdApply)(
		unsafe.Pointer(C.COnRtnSpdApply))

	spiCVtablePtr.CThostFtdcTraderSpi_OnErrRtnSpdApply = (C.OnErrRtnSpdApply)(
		unsafe.Pointer(C.COnErrRtnSpdApply))

	spiCVtablePtr.CThostFtdcTraderSpi_OnErrRtnSpdApplyAction = (C.OnErrRtnSpdApplyAction)(
		unsafe.Pointer(C.COnErrRtnSpdApplyAction))

	spiCVtablePtr.CThostFtdcTraderSpi_OnRspHedgeCfm = (C.OnRspHedgeCfm)(
		unsafe.Pointer(C.COnRspHedgeCfm))

	spiCVtablePtr.CThostFtdcTraderSpi_OnRspHedgeCfmAction = (C.OnRspHedgeCfmAction)(
		unsafe.Pointer(C.COnRspHedgeCfmAction))

	spiCVtablePtr.CThostFtdcTraderSpi_OnRspQryHedgeCfm = (C.OnRspQryHedgeCfm)(
		unsafe.Pointer(C.COnRspQryHedgeCfm))

	spiCVtablePtr.CThostFtdcTraderSpi_OnRtnHedgeCfm = (C.OnRtnHedgeCfm)(
		unsafe.Pointer(C.COnRtnHedgeCfm))

	spiCVtablePtr.CThostFtdcTraderSpi_OnErrRtnHedgeCfm = (C.OnErrRtnHedgeCfm)(
		unsafe.Pointer(C.COnErrRtnHedgeCfm))

	spiCVtablePtr.CThostFtdcTraderSpi_OnErrRtnHedgeCfmAction = (C.OnErrRtnHedgeCfmAction)(
		unsafe.Pointer(C.COnErrRtnHedgeCfmAction))

}

type ThostFtdcTraderSpi struct {
	runtime.Pinner
	callback thost.TraderSpi
}

//export CgoOnFrontConnected
func CgoOnFrontConnected(this unsafe.Pointer) {
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
func CgoOnFrontDisconnected(this unsafe.Pointer, nReason C.int) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnFrontDisconnected called",
		slog.Any("this", this),
		slog.Int("reason", int(nReason)),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnFrontDisconnected(int(nReason))
}

//export CgoOnHeartBeatWarning
func CgoOnHeartBeatWarning(this unsafe.Pointer, nTimeLapse C.int) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnHeartBeatWarning called",
		slog.Any("this", this),
		slog.Int("time_lapse", int(nTimeLapse)),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnHeartBeatWarning(int(nTimeLapse))
}

//export CgoOnRspAuthenticate
func CgoOnRspAuthenticate(
	this unsafe.Pointer,
	pRspAuthenticateField *C.struct_CThostFtdcRspAuthenticateField,
	pRspInfo *C.struct_CThostFtdcRspInfoField,
	nRequestID C.int, bIsLast C.bool,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRspAuthenticate called",
		slog.Any("this", this),
		slog.Int("requst", int(nRequestID)),
		slog.Bool("is_last", bool(bIsLast)),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnRspAuthenticate(
		(*thost.CThostFtdcRspAuthenticateField)(unsafe.Pointer(pRspAuthenticateField)),
		(*thost.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
		int(nRequestID), bool(bIsLast),
	)
}

//export CgoOnRtnPrivateSeqNo
func CgoOnRtnPrivateSeqNo(this unsafe.Pointer, nSeqNo C.int) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRtnPrivateSeqNo called",
		slog.Any("this", this),
		slog.Int("seq_no", int(nSeqNo)),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnRtnPrivateSeqNo(int(nSeqNo))
}

//export CgoOnRspUserLogin
func CgoOnRspUserLogin(
	this unsafe.Pointer,
	pRspUserLogin *C.struct_CThostFtdcRspUserLoginField,
	pRspInfo *C.struct_CThostFtdcRspInfoField,
	nRequestID C.int, bIsLast C.bool,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRspUserLogin called",
		slog.Any("this", this),
		slog.Int("requst", int(nRequestID)),
		slog.Bool("is_last", bool(bIsLast)),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnRspUserLogin(
		(*thost.CThostFtdcRspUserLoginField)(unsafe.Pointer(pRspUserLogin)),
		(*thost.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
		int(nRequestID), bool(bIsLast),
	)
}

//export CgoOnRspUserLogout
func CgoOnRspUserLogout(
	this unsafe.Pointer,
	pUserLogout *C.struct_CThostFtdcUserLogoutField,
	pRspInfo *C.struct_CThostFtdcRspInfoField,
	nRequestID C.int, bIsLast C.bool,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRspUserLogout called",
		slog.Any("this", this),
		slog.Int("requst", int(nRequestID)),
		slog.Bool("is_last", bool(bIsLast)),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnRspUserLogout(
		(*thost.CThostFtdcUserLogoutField)(unsafe.Pointer(pUserLogout)),
		(*thost.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
		int(nRequestID), bool(bIsLast),
	)
}

//export CgoOnRspUserPasswordUpdate
func CgoOnRspUserPasswordUpdate(
	this unsafe.Pointer,
	pUserPasswordUpdate *C.struct_CThostFtdcUserPasswordUpdateField,
	pRspInfo *C.struct_CThostFtdcRspInfoField,
	nRequestID C.int, bIsLast C.bool,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRspUserPasswordUpdate called",
		slog.Any("this", this),
		slog.Int("requst", int(nRequestID)),
		slog.Bool("is_last", bool(bIsLast)),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnRspUserPasswordUpdate(
		(*thost.CThostFtdcUserPasswordUpdateField)(unsafe.Pointer(pUserPasswordUpdate)),
		(*thost.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
		int(nRequestID), bool(bIsLast),
	)
}

//export CgoOnRspTradingAccountPasswordUpdate
func CgoOnRspTradingAccountPasswordUpdate(
	this unsafe.Pointer,
	pTradingAccountPasswordUpdate *C.struct_CThostFtdcTradingAccountPasswordUpdateField,
	pRspInfo *C.struct_CThostFtdcRspInfoField,
	nRequestID C.int, bIsLast C.bool,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRspTradingAccountPasswordUpdate called",
		slog.Any("this", this),
		slog.Int("requst", int(nRequestID)),
		slog.Bool("is_last", bool(bIsLast)),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnRspTradingAccountPasswordUpdate(
		(*thost.CThostFtdcTradingAccountPasswordUpdateField)(unsafe.Pointer(pTradingAccountPasswordUpdate)),
		(*thost.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
		int(nRequestID), bool(bIsLast),
	)
}

//export CgoOnRspUserAuthMethod
func CgoOnRspUserAuthMethod(
	this unsafe.Pointer,
	pRspUserAuthMethod *C.struct_CThostFtdcRspUserAuthMethodField,
	pRspInfo *C.struct_CThostFtdcRspInfoField,
	nRequestID C.int, bIsLast C.bool,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRspUserAuthMethod called",
		slog.Any("this", this),
		slog.Int("requst", int(nRequestID)),
		slog.Bool("is_last", bool(bIsLast)),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnRspUserAuthMethod(
		(*thost.CThostFtdcRspUserAuthMethodField)(unsafe.Pointer(pRspUserAuthMethod)),
		(*thost.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
		int(nRequestID), bool(bIsLast),
	)
}

//export CgoOnRspGenUserCaptcha
func CgoOnRspGenUserCaptcha(
	this unsafe.Pointer,
	pRspGenUserCaptcha *C.struct_CThostFtdcRspGenUserCaptchaField,
	pRspInfo *C.struct_CThostFtdcRspInfoField,
	nRequestID C.int, bIsLast C.bool,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRspGenUserCaptcha called",
		slog.Any("this", this),
		slog.Int("requst", int(nRequestID)),
		slog.Bool("is_last", bool(bIsLast)),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnRspGenUserCaptcha(
		(*thost.CThostFtdcRspGenUserCaptchaField)(unsafe.Pointer(pRspGenUserCaptcha)),
		(*thost.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
		int(nRequestID), bool(bIsLast),
	)
}

//export CgoOnRspGenUserText
func CgoOnRspGenUserText(
	this unsafe.Pointer,
	pRspGenUserText *C.struct_CThostFtdcRspGenUserTextField,
	pRspInfo *C.struct_CThostFtdcRspInfoField,
	nRequestID C.int, bIsLast C.bool,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRspGenUserText called",
		slog.Any("this", this),
		slog.Int("requst", int(nRequestID)),
		slog.Bool("is_last", bool(bIsLast)),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnRspGenUserText(
		(*thost.CThostFtdcRspGenUserTextField)(unsafe.Pointer(pRspGenUserText)),
		(*thost.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
		int(nRequestID), bool(bIsLast),
	)
}

//export CgoOnRspOrderInsert
func CgoOnRspOrderInsert(
	this unsafe.Pointer,
	pInputOrder *C.struct_CThostFtdcInputOrderField,
	pRspInfo *C.struct_CThostFtdcRspInfoField,
	nRequestID C.int, bIsLast C.bool,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRspOrderInsert called",
		slog.Any("this", this),
		slog.Int("requst", int(nRequestID)),
		slog.Bool("is_last", bool(bIsLast)),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnRspOrderInsert(
		(*thost.CThostFtdcInputOrderField)(unsafe.Pointer(pInputOrder)),
		(*thost.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
		int(nRequestID), bool(bIsLast),
	)
}

//export CgoOnRspParkedOrderInsert
func CgoOnRspParkedOrderInsert(
	this unsafe.Pointer,
	pParkedOrder *C.struct_CThostFtdcParkedOrderField,
	pRspInfo *C.struct_CThostFtdcRspInfoField,
	nRequestID C.int, bIsLast C.bool,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRspParkedOrderInsert called",
		slog.Any("this", this),
		slog.Int("requst", int(nRequestID)),
		slog.Bool("is_last", bool(bIsLast)),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnRspParkedOrderInsert(
		(*thost.CThostFtdcParkedOrderField)(unsafe.Pointer(pParkedOrder)),
		(*thost.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
		int(nRequestID), bool(bIsLast),
	)
}

//export CgoOnRspParkedOrderAction
func CgoOnRspParkedOrderAction(
	this unsafe.Pointer,
	pParkedOrderAction *C.struct_CThostFtdcParkedOrderActionField,
	pRspInfo *C.struct_CThostFtdcRspInfoField,
	nRequestID C.int, bIsLast C.bool,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRspParkedOrderAction called",
		slog.Any("this", this),
		slog.Int("requst", int(nRequestID)),
		slog.Bool("is_last", bool(bIsLast)),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnRspParkedOrderAction(
		(*thost.CThostFtdcParkedOrderActionField)(unsafe.Pointer(pParkedOrderAction)),
		(*thost.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
		int(nRequestID), bool(bIsLast),
	)
}

//export CgoOnRspOrderAction
func CgoOnRspOrderAction(
	this unsafe.Pointer,
	pInputOrderAction *C.struct_CThostFtdcInputOrderActionField,
	pRspInfo *C.struct_CThostFtdcRspInfoField,
	nRequestID C.int, bIsLast C.bool,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRspOrderAction called",
		slog.Any("this", this),
		slog.Int("requst", int(nRequestID)),
		slog.Bool("is_last", bool(bIsLast)),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnRspOrderAction(
		(*thost.CThostFtdcInputOrderActionField)(unsafe.Pointer(pInputOrderAction)),
		(*thost.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
		int(nRequestID), bool(bIsLast),
	)
}

//export CgoOnRspQryMaxOrderVolume
func CgoOnRspQryMaxOrderVolume(
	this unsafe.Pointer,
	pQryMaxOrderVolume *C.struct_CThostFtdcQryMaxOrderVolumeField,
	pRspInfo *C.struct_CThostFtdcRspInfoField,
	nRequestID C.int, bIsLast C.bool,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRspQryMaxOrderVolume called",
		slog.Any("this", this),
		slog.Int("request", int(nRequestID)),
		slog.Bool("is_last", bool(bIsLast)),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnRspQryMaxOrderVolume(
		(*thost.CThostFtdcQryMaxOrderVolumeField)(unsafe.Pointer(pQryMaxOrderVolume)),
		(*thost.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
		int(nRequestID), bool(bIsLast),
	)
}

//export CgoOnRspSettlementInfoConfirm
func CgoOnRspSettlementInfoConfirm(
	this unsafe.Pointer,
	pSettlementInfoConfirm *C.struct_CThostFtdcSettlementInfoConfirmField,
	pRspInfo *C.struct_CThostFtdcRspInfoField,
	nRequestID C.int, bIsLast C.bool,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRspSettlementInfoConfirm called",
		slog.Any("this", this),
		slog.Int("request", int(nRequestID)),
		slog.Bool("is_last", bool(bIsLast)),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnRspSettlementInfoConfirm(
		(*thost.CThostFtdcSettlementInfoConfirmField)(unsafe.Pointer(pSettlementInfoConfirm)),
		(*thost.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
		int(nRequestID), bool(bIsLast),
	)
}

//export CgoOnRspRemoveParkedOrder
func CgoOnRspRemoveParkedOrder(
	this unsafe.Pointer,
	pRemoveParkedOrder *C.struct_CThostFtdcRemoveParkedOrderField,
	pRspInfo *C.struct_CThostFtdcRspInfoField,
	nRequestID C.int, bIsLast C.bool,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRspRemoveParkedOrder called",
		slog.Any("this", this),
		slog.Int("request", int(nRequestID)),
		slog.Bool("is_last", bool(bIsLast)),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnRspRemoveParkedOrder(
		(*thost.CThostFtdcRemoveParkedOrderField)(unsafe.Pointer(pRemoveParkedOrder)),
		(*thost.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
		int(nRequestID), bool(bIsLast),
	)
}

//export CgoOnRspRemoveParkedOrderAction
func CgoOnRspRemoveParkedOrderAction(
	this unsafe.Pointer,
	pRemoveParkedOrderAction *C.struct_CThostFtdcRemoveParkedOrderActionField,
	pRspInfo *C.struct_CThostFtdcRspInfoField,
	nRequestID C.int, bIsLast C.bool,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRspRemoveParkedOrderAction called",
		slog.Any("this", this),
		slog.Int("request", int(nRequestID)),
		slog.Bool("is_last", bool(bIsLast)),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnRspRemoveParkedOrderAction(
		(*thost.CThostFtdcRemoveParkedOrderActionField)(unsafe.Pointer(pRemoveParkedOrderAction)),
		(*thost.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
		int(nRequestID), bool(bIsLast),
	)
}

//export CgoOnRspExecOrderInsert
func CgoOnRspExecOrderInsert(
	this unsafe.Pointer,
	pInputExecOrder *C.struct_CThostFtdcInputExecOrderField,
	pRspInfo *C.struct_CThostFtdcRspInfoField,
	nRequestID C.int, bIsLast C.bool,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRspExecOrderInsert called",
		slog.Any("this", this),
		slog.Int("request", int(nRequestID)),
		slog.Bool("is_last", bool(bIsLast)),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnRspExecOrderInsert(
		(*thost.CThostFtdcInputExecOrderField)(unsafe.Pointer(pInputExecOrder)),
		(*thost.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
		int(nRequestID), bool(bIsLast),
	)
}

//export CgoOnRspExecOrderAction
func CgoOnRspExecOrderAction(
	this unsafe.Pointer,
	pInputExecOrderAction *C.struct_CThostFtdcInputExecOrderActionField,
	pRspInfo *C.struct_CThostFtdcRspInfoField,
	nRequestID C.int, bIsLast C.bool,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRspExecOrderAction called",
		slog.Any("this", this),
		slog.Int("request", int(nRequestID)),
		slog.Bool("is_last", bool(bIsLast)),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnRspExecOrderAction(
		(*thost.CThostFtdcInputExecOrderActionField)(unsafe.Pointer(pInputExecOrderAction)),
		(*thost.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
		int(nRequestID), bool(bIsLast),
	)
}

//export CgoOnRspForQuoteInsert
func CgoOnRspForQuoteInsert(
	this unsafe.Pointer,
	pInputForQuote *C.struct_CThostFtdcInputForQuoteField,
	pRspInfo *C.struct_CThostFtdcRspInfoField,
	nRequestID C.int, bIsLast C.bool,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRspForQuoteInsert called",
		slog.Any("this", this),
		slog.Int("request", int(nRequestID)),
		slog.Bool("is_last", bool(bIsLast)),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnRspForQuoteInsert(
		(*thost.CThostFtdcInputForQuoteField)(unsafe.Pointer(pInputForQuote)),
		(*thost.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
		int(nRequestID), bool(bIsLast),
	)
}

//export CgoOnRspQuoteInsert
func CgoOnRspQuoteInsert(
	this unsafe.Pointer,
	pInputQuote *C.struct_CThostFtdcInputQuoteField,
	pRspInfo *C.struct_CThostFtdcRspInfoField,
	nRequestID C.int, bIsLast C.bool,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRspQuoteInsert called",
		slog.Any("this", this),
		slog.Int("request", int(nRequestID)),
		slog.Bool("is_last", bool(bIsLast)),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnRspQuoteInsert(
		(*thost.CThostFtdcInputQuoteField)(unsafe.Pointer(pInputQuote)),
		(*thost.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
		int(nRequestID), bool(bIsLast),
	)
}

//export CgoOnRspQuoteAction
func CgoOnRspQuoteAction(
	this unsafe.Pointer,
	pInputQuoteAction *C.struct_CThostFtdcInputQuoteActionField,
	pRspInfo *C.struct_CThostFtdcRspInfoField,
	nRequestID C.int, bIsLast C.bool,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRspQuoteAction called",
		slog.Any("this", this),
		slog.Int("request", int(nRequestID)),
		slog.Bool("is_last", bool(bIsLast)),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnRspQuoteAction(
		(*thost.CThostFtdcInputQuoteActionField)(unsafe.Pointer(pInputQuoteAction)),
		(*thost.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
		int(nRequestID), bool(bIsLast),
	)
}

//export CgoOnRspBatchOrderAction
func CgoOnRspBatchOrderAction(
	this unsafe.Pointer,
	pInputBatchOrderAction *C.struct_CThostFtdcInputBatchOrderActionField,
	pRspInfo *C.struct_CThostFtdcRspInfoField,
	nRequestID C.int, bIsLast C.bool,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRspBatchOrderAction called",
		slog.Any("this", this),
		slog.Int("request", int(nRequestID)),
		slog.Bool("is_last", bool(bIsLast)),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnRspBatchOrderAction(
		(*thost.CThostFtdcInputBatchOrderActionField)(unsafe.Pointer(pInputBatchOrderAction)),
		(*thost.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
		int(nRequestID), bool(bIsLast),
	)
}

//export CgoOnRspOptionSelfCloseInsert
func CgoOnRspOptionSelfCloseInsert(
	this unsafe.Pointer,
	pInputOptionSelfClose *C.struct_CThostFtdcInputOptionSelfCloseField,
	pRspInfo *C.struct_CThostFtdcRspInfoField,
	nRequestID C.int, bIsLast C.bool,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRspOptionSelfCloseInsert called",
		slog.Any("this", this),
		slog.Int("request", int(nRequestID)),
		slog.Bool("is_last", bool(bIsLast)),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnRspOptionSelfCloseInsert(
		(*thost.CThostFtdcInputOptionSelfCloseField)(unsafe.Pointer(pInputOptionSelfClose)),
		(*thost.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
		int(nRequestID), bool(bIsLast),
	)
}

//export CgoOnRspOptionSelfCloseAction
func CgoOnRspOptionSelfCloseAction(
	this unsafe.Pointer,
	pInputOptionSelfCloseAction *C.struct_CThostFtdcInputOptionSelfCloseActionField,
	pRspInfo *C.struct_CThostFtdcRspInfoField,
	nRequestID C.int, bIsLast C.bool,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRspOptionSelfCloseAction called",
		slog.Any("this", this),
		slog.Int("request", int(nRequestID)),
		slog.Bool("is_last", bool(bIsLast)),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnRspOptionSelfCloseAction(
		(*thost.CThostFtdcInputOptionSelfCloseActionField)(unsafe.Pointer(pInputOptionSelfCloseAction)),
		(*thost.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
		int(nRequestID), bool(bIsLast),
	)
}

//export CgoOnRspCombActionInsert
func CgoOnRspCombActionInsert(
	this unsafe.Pointer,
	pInputCombAction *C.struct_CThostFtdcInputCombActionField,
	pRspInfo *C.struct_CThostFtdcRspInfoField,
	nRequestID C.int, bIsLast C.bool,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRspCombActionInsert called",
		slog.Any("this", this),
		slog.Int("request", int(nRequestID)),
		slog.Bool("is_last", bool(bIsLast)),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnRspCombActionInsert(
		(*thost.CThostFtdcInputCombActionField)(unsafe.Pointer(pInputCombAction)),
		(*thost.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
		int(nRequestID), bool(bIsLast),
	)
}

//export CgoOnRspQryOrder
func CgoOnRspQryOrder(
	this unsafe.Pointer,
	pOrder *C.struct_CThostFtdcOrderField,
	pRspInfo *C.struct_CThostFtdcRspInfoField,
	nRequestID C.int, bIsLast C.bool,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRspQryOrder called",
		slog.Any("this", this),
		slog.Int("request", int(nRequestID)),
		slog.Bool("is_last", bool(bIsLast)),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnRspQryOrder(
		(*thost.CThostFtdcOrderField)(unsafe.Pointer(pOrder)),
		(*thost.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
		int(nRequestID), bool(bIsLast),
	)
}

//export CgoOnRspQryTrade
func CgoOnRspQryTrade(
	this unsafe.Pointer,
	pTrade *C.struct_CThostFtdcTradeField,
	pRspInfo *C.struct_CThostFtdcRspInfoField,
	nRequestID C.int, bIsLast C.bool,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRspQryTrade called",
		slog.Any("this", this),
		slog.Int("request", int(nRequestID)),
		slog.Bool("is_last", bool(bIsLast)),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnRspQryTrade(
		(*thost.CThostFtdcTradeField)(unsafe.Pointer(pTrade)),
		(*thost.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
		int(nRequestID), bool(bIsLast),
	)
}

//export CgoOnRspQryInvestorPosition
func CgoOnRspQryInvestorPosition(
	this unsafe.Pointer,
	pInvestorPosition *C.struct_CThostFtdcInvestorPositionField,
	pRspInfo *C.struct_CThostFtdcRspInfoField,
	nRequestID C.int, bIsLast C.bool,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRspQryInvestorPosition called",
		slog.Any("this", this),
		slog.Int("request", int(nRequestID)),
		slog.Bool("is_last", bool(bIsLast)),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnRspQryInvestorPosition(
		(*thost.CThostFtdcInvestorPositionField)(unsafe.Pointer(pInvestorPosition)),
		(*thost.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
		int(nRequestID), bool(bIsLast),
	)
}

//export CgoOnRspQryTradingAccount
func CgoOnRspQryTradingAccount(
	this unsafe.Pointer,
	pTradingAccount *C.struct_CThostFtdcTradingAccountField,
	pRspInfo *C.struct_CThostFtdcRspInfoField,
	nRequestID C.int, bIsLast C.bool,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRspQryTradingAccount called",
		slog.Any("this", this),
		slog.Int("request", int(nRequestID)),
		slog.Bool("is_last", bool(bIsLast)),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnRspQryTradingAccount(
		(*thost.CThostFtdcTradingAccountField)(unsafe.Pointer(pTradingAccount)),
		(*thost.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
		int(nRequestID), bool(bIsLast),
	)
}

//export CgoOnRspQryInvestor
func CgoOnRspQryInvestor(
	this unsafe.Pointer,
	pInvestor *C.struct_CThostFtdcInvestorField,
	pRspInfo *C.struct_CThostFtdcRspInfoField,
	nRequestID C.int, bIsLast C.bool,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRspQryInvestor called",
		slog.Any("this", this),
		slog.Int("request", int(nRequestID)),
		slog.Bool("is_last", bool(bIsLast)),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnRspQryInvestor(
		(*thost.CThostFtdcInvestorField)(unsafe.Pointer(pInvestor)),
		(*thost.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
		int(nRequestID), bool(bIsLast),
	)
}

//export CgoOnRspQryTradingCode
func CgoOnRspQryTradingCode(
	this unsafe.Pointer,
	pTradingCode *C.struct_CThostFtdcTradingCodeField,
	pRspInfo *C.struct_CThostFtdcRspInfoField,
	nRequestID C.int, bIsLast C.bool,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRspQryTradingCode called",
		slog.Any("this", this),
		slog.Int("request", int(nRequestID)),
		slog.Bool("is_last", bool(bIsLast)),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnRspQryTradingCode(
		(*thost.CThostFtdcTradingCodeField)(unsafe.Pointer(pTradingCode)),
		(*thost.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
		int(nRequestID), bool(bIsLast),
	)
}

//export CgoOnRspQryInstrumentMarginRate
func CgoOnRspQryInstrumentMarginRate(
	this unsafe.Pointer,
	pInstrumentMarginRate *C.struct_CThostFtdcInstrumentMarginRateField,
	pRspInfo *C.struct_CThostFtdcRspInfoField,
	nRequestID C.int, bIsLast C.bool,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRspQryInstrumentMarginRate called",
		slog.Any("this", this),
		slog.Int("request", int(nRequestID)),
		slog.Bool("is_last", bool(bIsLast)),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnRspQryInstrumentMarginRate(
		(*thost.CThostFtdcInstrumentMarginRateField)(unsafe.Pointer(pInstrumentMarginRate)),
		(*thost.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
		int(nRequestID), bool(bIsLast),
	)
}

//export CgoOnRspQryInstrumentCommissionRate
func CgoOnRspQryInstrumentCommissionRate(
	this unsafe.Pointer,
	pInstrumentCommissionRate *C.struct_CThostFtdcInstrumentCommissionRateField,
	pRspInfo *C.struct_CThostFtdcRspInfoField,
	nRequestID C.int, bIsLast C.bool,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRspQryInstrumentCommissionRate called",
		slog.Any("this", this),
		slog.Int("request", int(nRequestID)),
		slog.Bool("is_last", bool(bIsLast)),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnRspQryInstrumentCommissionRate(
		(*thost.CThostFtdcInstrumentCommissionRateField)(unsafe.Pointer(pInstrumentCommissionRate)),
		(*thost.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
		int(nRequestID), bool(bIsLast),
	)
}

//export CgoOnRspQryUserSession
func CgoOnRspQryUserSession(
	this unsafe.Pointer,
	pUserSession *C.struct_CThostFtdcUserSessionField,
	pRspInfo *C.struct_CThostFtdcRspInfoField,
	nRequestID C.int, bIsLast C.bool,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRspQryUserSession called",
		slog.Any("this", this),
		slog.Int("request", int(nRequestID)),
		slog.Bool("is_last", bool(bIsLast)),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnRspQryUserSession(
		(*thost.CThostFtdcUserSessionField)(unsafe.Pointer(pUserSession)),
		(*thost.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
		int(nRequestID), bool(bIsLast),
	)
}

//export CgoOnRspQryExchange
func CgoOnRspQryExchange(
	this unsafe.Pointer,
	pExchange *C.struct_CThostFtdcExchangeField,
	pRspInfo *C.struct_CThostFtdcRspInfoField,
	nRequestID C.int, bIsLast C.bool,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRspQryExchange called",
		slog.Any("this", this),
		slog.Int("request", int(nRequestID)),
		slog.Bool("is_last", bool(bIsLast)),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnRspQryExchange(
		(*thost.CThostFtdcExchangeField)(unsafe.Pointer(pExchange)),
		(*thost.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
		int(nRequestID), bool(bIsLast),
	)
}

//export CgoOnRspQryProduct
func CgoOnRspQryProduct(
	this unsafe.Pointer,
	pProduct *C.struct_CThostFtdcProductField,
	pRspInfo *C.struct_CThostFtdcRspInfoField,
	nRequestID C.int, bIsLast C.bool,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRspQryProduct called",
		slog.Any("this", this),
		slog.Int("request", int(nRequestID)),
		slog.Bool("is_last", bool(bIsLast)),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnRspQryProduct(
		(*thost.CThostFtdcProductField)(unsafe.Pointer(pProduct)),
		(*thost.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
		int(nRequestID), bool(bIsLast),
	)
}

//export CgoOnRspQryInstrument
func CgoOnRspQryInstrument(
	this unsafe.Pointer,
	pInstrument *C.struct_CThostFtdcInstrumentField,
	pRspInfo *C.struct_CThostFtdcRspInfoField,
	nRequestID C.int, bIsLast C.bool,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRspQryInstrument called",
		slog.Any("this", this),
		slog.Int("request", int(nRequestID)),
		slog.Bool("is_last", bool(bIsLast)),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnRspQryInstrument(
		(*thost.CThostFtdcInstrumentField)(unsafe.Pointer(pInstrument)),
		(*thost.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
		int(nRequestID), bool(bIsLast),
	)
}

//export CgoOnRspQryDepthMarketData
func CgoOnRspQryDepthMarketData(
	this unsafe.Pointer,
	pDepthMarketData *C.struct_CThostFtdcDepthMarketDataField,
	pRspInfo *C.struct_CThostFtdcRspInfoField,
	nRequestID C.int, bIsLast C.bool,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRspQryDepthMarketData called",
		slog.Any("this", this),
		slog.Int("request", int(nRequestID)),
		slog.Bool("is_last", bool(bIsLast)),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnRspQryDepthMarketData(
		(*thost.CThostFtdcDepthMarketDataField)(unsafe.Pointer(pDepthMarketData)),
		(*thost.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
		int(nRequestID), bool(bIsLast),
	)
}

//export CgoOnRspQryTraderOffer
func CgoOnRspQryTraderOffer(
	this unsafe.Pointer,
	pTraderOffer *C.struct_CThostFtdcTraderOfferField,
	pRspInfo *C.struct_CThostFtdcRspInfoField,
	nRequestID C.int, bIsLast C.bool,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRspQryTraderOffer called",
		slog.Any("this", this),
		slog.Int("request", int(nRequestID)),
		slog.Bool("is_last", bool(bIsLast)),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnRspQryTraderOffer(
		(*thost.CThostFtdcTraderOfferField)(unsafe.Pointer(pTraderOffer)),
		(*thost.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
		int(nRequestID), bool(bIsLast),
	)
}

//export CgoOnRspQrySettlementInfo
func CgoOnRspQrySettlementInfo(
	this unsafe.Pointer,
	pSettlementInfo *C.struct_CThostFtdcSettlementInfoField,
	pRspInfo *C.struct_CThostFtdcRspInfoField,
	nRequestID C.int, bIsLast C.bool,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRspQrySettlementInfo called",
		slog.Any("this", this),
		slog.Int("request", int(nRequestID)),
		slog.Bool("is_last", bool(bIsLast)),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnRspQrySettlementInfo(
		(*thost.CThostFtdcSettlementInfoField)(unsafe.Pointer(pSettlementInfo)),
		(*thost.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
		int(nRequestID), bool(bIsLast),
	)
}

//export CgoOnRspQryTransferBank
func CgoOnRspQryTransferBank(
	this unsafe.Pointer,
	pTransferBank *C.struct_CThostFtdcTransferBankField,
	pRspInfo *C.struct_CThostFtdcRspInfoField,
	nRequestID C.int, bIsLast C.bool,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRspQryTransferBank called",
		slog.Any("this", this),
		slog.Int("request", int(nRequestID)),
		slog.Bool("is_last", bool(bIsLast)),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnRspQryTransferBank(
		(*thost.CThostFtdcTransferBankField)(unsafe.Pointer(pTransferBank)),
		(*thost.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
		int(nRequestID), bool(bIsLast),
	)
}

//export CgoOnRspQryInvestorPositionDetail
func CgoOnRspQryInvestorPositionDetail(
	this unsafe.Pointer,
	pInvestorPositionDetail *C.struct_CThostFtdcInvestorPositionDetailField,
	pRspInfo *C.struct_CThostFtdcRspInfoField,
	nRequestID C.int, bIsLast C.bool,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRspQryInvestorPositionDetail called",
		slog.Any("this", this),
		slog.Int("request", int(nRequestID)),
		slog.Bool("is_last", bool(bIsLast)),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnRspQryInvestorPositionDetail(
		(*thost.CThostFtdcInvestorPositionDetailField)(unsafe.Pointer(pInvestorPositionDetail)),
		(*thost.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
		int(nRequestID), bool(bIsLast),
	)
}

//export CgoOnRspQryNotice
func CgoOnRspQryNotice(
	this unsafe.Pointer,
	pNotice *C.struct_CThostFtdcNoticeField,
	pRspInfo *C.struct_CThostFtdcRspInfoField,
	nRequestID C.int, bIsLast C.bool,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRspQryNotice called",
		slog.Any("this", this),
		slog.Int("request", int(nRequestID)),
		slog.Bool("is_last", bool(bIsLast)),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnRspQryNotice(
		(*thost.CThostFtdcNoticeField)(unsafe.Pointer(pNotice)),
		(*thost.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
		int(nRequestID), bool(bIsLast),
	)
}

//export CgoOnRspQrySettlementInfoConfirm
func CgoOnRspQrySettlementInfoConfirm(
	this unsafe.Pointer,
	pSettlementInfoConfirm *C.struct_CThostFtdcSettlementInfoConfirmField,
	pRspInfo *C.struct_CThostFtdcRspInfoField,
	nRequestID C.int, bIsLast C.bool,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRspQrySettlementInfoConfirm called",
		slog.Any("this", this),
		slog.Int("request", int(nRequestID)),
		slog.Bool("is_last", bool(bIsLast)),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnRspQrySettlementInfoConfirm(
		(*thost.CThostFtdcSettlementInfoConfirmField)(unsafe.Pointer(pSettlementInfoConfirm)),
		(*thost.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
		int(nRequestID), bool(bIsLast),
	)
}

//export CgoOnRspQryInvestorPositionCombineDetail
func CgoOnRspQryInvestorPositionCombineDetail(
	this unsafe.Pointer,
	pInvestorPositionCombineDetail *C.struct_CThostFtdcInvestorPositionCombineDetailField,
	pRspInfo *C.struct_CThostFtdcRspInfoField,
	nRequestID C.int, bIsLast C.bool,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRspQryInvestorPositionCombineDetail called",
		slog.Any("this", this),
		slog.Int("request", int(nRequestID)),
		slog.Bool("is_last", bool(bIsLast)),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnRspQryInvestorPositionCombineDetail(
		(*thost.CThostFtdcInvestorPositionCombineDetailField)(unsafe.Pointer(pInvestorPositionCombineDetail)),
		(*thost.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
		int(nRequestID), bool(bIsLast),
	)
}

//export CgoOnRspQryCFMMCTradingAccountKey
func CgoOnRspQryCFMMCTradingAccountKey(
	this unsafe.Pointer,
	pCFMMCTradingAccountKey *C.struct_CThostFtdcCFMMCTradingAccountKeyField,
	pRspInfo *C.struct_CThostFtdcRspInfoField,
	nRequestID C.int, bIsLast C.bool,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRspQryCFMMCTradingAccountKey called",
		slog.Any("this", this),
		slog.Int("request", int(nRequestID)),
		slog.Bool("is_last", bool(bIsLast)),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnRspQryCFMMCTradingAccountKey(
		(*thost.CThostFtdcCFMMCTradingAccountKeyField)(unsafe.Pointer(pCFMMCTradingAccountKey)),
		(*thost.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
		int(nRequestID), bool(bIsLast),
	)
}

//export CgoOnRspQryEWarrantOffset
func CgoOnRspQryEWarrantOffset(
	this unsafe.Pointer,
	pEWarrantOffset *C.struct_CThostFtdcEWarrantOffsetField,
	pRspInfo *C.struct_CThostFtdcRspInfoField,
	nRequestID C.int, bIsLast C.bool,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRspQryEWarrantOffset called",
		slog.Any("this", this),
		slog.Int("request", int(nRequestID)),
		slog.Bool("is_last", bool(bIsLast)),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnRspQryEWarrantOffset(
		(*thost.CThostFtdcEWarrantOffsetField)(unsafe.Pointer(pEWarrantOffset)),
		(*thost.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
		int(nRequestID), bool(bIsLast),
	)
}

//export CgoOnRspQryInvestorProductGroupMargin
func CgoOnRspQryInvestorProductGroupMargin(
	this unsafe.Pointer,
	pInvestorProductGroupMargin *C.struct_CThostFtdcInvestorProductGroupMarginField,
	pRspInfo *C.struct_CThostFtdcRspInfoField,
	nRequestID C.int, bIsLast C.bool,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRspQryInvestorProductGroupMargin called",
		slog.Any("this", this),
		slog.Int("request", int(nRequestID)),
		slog.Bool("is_last", bool(bIsLast)),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnRspQryInvestorProductGroupMargin(
		(*thost.CThostFtdcInvestorProductGroupMarginField)(unsafe.Pointer(pInvestorProductGroupMargin)),
		(*thost.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
		int(nRequestID), bool(bIsLast),
	)
}

//export CgoOnRspQryExchangeMarginRate
func CgoOnRspQryExchangeMarginRate(
	this unsafe.Pointer,
	pExchangeMarginRate *C.struct_CThostFtdcExchangeMarginRateField,
	pRspInfo *C.struct_CThostFtdcRspInfoField,
	nRequestID C.int, bIsLast C.bool,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRspQryExchangeMarginRate called",
		slog.Any("this", this),
		slog.Int("request", int(nRequestID)),
		slog.Bool("is_last", bool(bIsLast)),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnRspQryExchangeMarginRate(
		(*thost.CThostFtdcExchangeMarginRateField)(unsafe.Pointer(pExchangeMarginRate)),
		(*thost.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
		int(nRequestID), bool(bIsLast),
	)
}

//export CgoOnRspQryExchangeMarginRateAdjust
func CgoOnRspQryExchangeMarginRateAdjust(
	this unsafe.Pointer,
	pExchangeMarginRateAdjust *C.struct_CThostFtdcExchangeMarginRateAdjustField,
	pRspInfo *C.struct_CThostFtdcRspInfoField,
	nRequestID C.int, bIsLast C.bool,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRspQryExchangeMarginRateAdjust called",
		slog.Any("this", this),
		slog.Int("request", int(nRequestID)),
		slog.Bool("is_last", bool(bIsLast)),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnRspQryExchangeMarginRateAdjust(
		(*thost.CThostFtdcExchangeMarginRateAdjustField)(unsafe.Pointer(pExchangeMarginRateAdjust)),
		(*thost.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
		int(nRequestID), bool(bIsLast),
	)
}

//export CgoOnRspQryExchangeRate
func CgoOnRspQryExchangeRate(
	this unsafe.Pointer,
	pExchangeRate *C.struct_CThostFtdcExchangeRateField,
	pRspInfo *C.struct_CThostFtdcRspInfoField,
	nRequestID C.int, bIsLast C.bool,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRspQryExchangeRate called",
		slog.Any("this", this),
		slog.Int("request", int(nRequestID)),
		slog.Bool("is_last", bool(bIsLast)),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnRspQryExchangeRate(
		(*thost.CThostFtdcExchangeRateField)(unsafe.Pointer(pExchangeRate)),
		(*thost.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
		int(nRequestID), bool(bIsLast),
	)
}

//export CgoOnRspQrySecAgentACIDMap
func CgoOnRspQrySecAgentACIDMap(
	this unsafe.Pointer,
	pSecAgentACIDMap *C.struct_CThostFtdcSecAgentACIDMapField,
	pRspInfo *C.struct_CThostFtdcRspInfoField,
	nRequestID C.int, bIsLast C.bool,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRspQrySecAgentACIDMap called",
		slog.Any("this", this),
		slog.Int("request", int(nRequestID)),
		slog.Bool("is_last", bool(bIsLast)),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnRspQrySecAgentACIDMap(
		(*thost.CThostFtdcSecAgentACIDMapField)(unsafe.Pointer(pSecAgentACIDMap)),
		(*thost.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
		int(nRequestID), bool(bIsLast),
	)
}

//export CgoOnRspQryProductExchRate
func CgoOnRspQryProductExchRate(
	this unsafe.Pointer,
	pProductExchRate *C.struct_CThostFtdcProductExchRateField,
	pRspInfo *C.struct_CThostFtdcRspInfoField,
	nRequestID C.int, bIsLast C.bool,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRspQryProductExchRate called",
		slog.Any("this", this),
		slog.Int("request", int(nRequestID)),
		slog.Bool("is_last", bool(bIsLast)),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnRspQryProductExchRate(
		(*thost.CThostFtdcProductExchRateField)(unsafe.Pointer(pProductExchRate)),
		(*thost.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
		int(nRequestID), bool(bIsLast),
	)
}

//export CgoOnRspQryProductGroup
func CgoOnRspQryProductGroup(
	this unsafe.Pointer,
	pProductGroup *C.struct_CThostFtdcProductGroupField,
	pRspInfo *C.struct_CThostFtdcRspInfoField,
	nRequestID C.int, bIsLast C.bool,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRspQryProductGroup called",
		slog.Any("this", this),
		slog.Int("request", int(nRequestID)),
		slog.Bool("is_last", bool(bIsLast)),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnRspQryProductGroup(
		(*thost.CThostFtdcProductGroupField)(unsafe.Pointer(pProductGroup)),
		(*thost.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
		int(nRequestID), bool(bIsLast),
	)
}

//export CgoOnRspQryMMInstrumentCommissionRate
func CgoOnRspQryMMInstrumentCommissionRate(
	this unsafe.Pointer,
	pMMInstrumentCommissionRate *C.struct_CThostFtdcMMInstrumentCommissionRateField,
	pRspInfo *C.struct_CThostFtdcRspInfoField,
	nRequestID C.int, bIsLast C.bool,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRspQryMMInstrumentCommissionRate called",
		slog.Any("this", this),
		slog.Int("request", int(nRequestID)),
		slog.Bool("is_last", bool(bIsLast)),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnRspQryMMInstrumentCommissionRate(
		(*thost.CThostFtdcMMInstrumentCommissionRateField)(unsafe.Pointer(pMMInstrumentCommissionRate)),
		(*thost.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
		int(nRequestID), bool(bIsLast),
	)
}

//export CgoOnRspQryMMOptionInstrCommRate
func CgoOnRspQryMMOptionInstrCommRate(
	this unsafe.Pointer,
	pMMOptionInstrCommRate *C.struct_CThostFtdcMMOptionInstrCommRateField,
	pRspInfo *C.struct_CThostFtdcRspInfoField,
	nRequestID C.int, bIsLast C.bool,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRspQryMMOptionInstrCommRate called",
		slog.Any("this", this),
		slog.Int("request", int(nRequestID)),
		slog.Bool("is_last", bool(bIsLast)),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnRspQryMMOptionInstrCommRate(
		(*thost.CThostFtdcMMOptionInstrCommRateField)(unsafe.Pointer(pMMOptionInstrCommRate)),
		(*thost.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
		int(nRequestID), bool(bIsLast),
	)
}

//export CgoOnRspQryInstrumentOrderCommRate
func CgoOnRspQryInstrumentOrderCommRate(
	this unsafe.Pointer,
	pInstrumentOrderCommRate *C.struct_CThostFtdcInstrumentOrderCommRateField,
	pRspInfo *C.struct_CThostFtdcRspInfoField,
	nRequestID C.int, bIsLast C.bool,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRspQryInstrumentOrderCommRate called",
		slog.Any("this", this),
		slog.Int("request", int(nRequestID)),
		slog.Bool("is_last", bool(bIsLast)),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnRspQryInstrumentOrderCommRate(
		(*thost.CThostFtdcInstrumentOrderCommRateField)(unsafe.Pointer(pInstrumentOrderCommRate)),
		(*thost.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
		int(nRequestID), bool(bIsLast),
	)
}

//export CgoOnRspQrySecAgentTradingAccount
func CgoOnRspQrySecAgentTradingAccount(
	this unsafe.Pointer,
	pTradingAccount *C.struct_CThostFtdcTradingAccountField,
	pRspInfo *C.struct_CThostFtdcRspInfoField,
	nRequestID C.int, bIsLast C.bool,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRspQrySecAgentTradingAccount called",
		slog.Any("this", this),
		slog.Int("request", int(nRequestID)),
		slog.Bool("is_last", bool(bIsLast)),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnRspQrySecAgentTradingAccount(
		(*thost.CThostFtdcTradingAccountField)(unsafe.Pointer(pTradingAccount)),
		(*thost.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
		int(nRequestID), bool(bIsLast),
	)
}

//export CgoOnRspQrySecAgentCheckMode
func CgoOnRspQrySecAgentCheckMode(
	this unsafe.Pointer,
	pSecAgentCheckMode *C.struct_CThostFtdcSecAgentCheckModeField,
	pRspInfo *C.struct_CThostFtdcRspInfoField,
	nRequestID C.int, bIsLast C.bool,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRspQrySecAgentCheckMode called",
		slog.Any("this", this),
		slog.Int("request", int(nRequestID)),
		slog.Bool("is_last", bool(bIsLast)),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnRspQrySecAgentCheckMode(
		(*thost.CThostFtdcSecAgentCheckModeField)(unsafe.Pointer(pSecAgentCheckMode)),
		(*thost.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
		int(nRequestID), bool(bIsLast),
	)
}

//export CgoOnRspQrySecAgentTradeInfo
func CgoOnRspQrySecAgentTradeInfo(
	this unsafe.Pointer,
	pSecAgentTradeInfo *C.struct_CThostFtdcSecAgentTradeInfoField,
	pRspInfo *C.struct_CThostFtdcRspInfoField,
	nRequestID C.int, bIsLast C.bool,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRspQrySecAgentTradeInfo called",
		slog.Any("this", this),
		slog.Int("request", int(nRequestID)),
		slog.Bool("is_last", bool(bIsLast)),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnRspQrySecAgentTradeInfo(
		(*thost.CThostFtdcSecAgentTradeInfoField)(unsafe.Pointer(pSecAgentTradeInfo)),
		(*thost.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
		int(nRequestID), bool(bIsLast),
	)
}

//export CgoOnRspQryOptionInstrTradeCost
func CgoOnRspQryOptionInstrTradeCost(
	this unsafe.Pointer,
	pOptionInstrTradeCost *C.struct_CThostFtdcOptionInstrTradeCostField,
	pRspInfo *C.struct_CThostFtdcRspInfoField,
	nRequestID C.int, bIsLast C.bool,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRspQryOptionInstrTradeCost called",
		slog.Any("this", this),
		slog.Int("request", int(nRequestID)),
		slog.Bool("is_last", bool(bIsLast)),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnRspQryOptionInstrTradeCost(
		(*thost.CThostFtdcOptionInstrTradeCostField)(unsafe.Pointer(pOptionInstrTradeCost)),
		(*thost.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
		int(nRequestID), bool(bIsLast),
	)
}

//export CgoOnRspQryOptionInstrCommRate
func CgoOnRspQryOptionInstrCommRate(
	this unsafe.Pointer,
	pOptionInstrCommRate *C.struct_CThostFtdcOptionInstrCommRateField,
	pRspInfo *C.struct_CThostFtdcRspInfoField,
	nRequestID C.int, bIsLast C.bool,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRspQryOptionInstrCommRate called",
		slog.Any("this", this),
		slog.Int("request", int(nRequestID)),
		slog.Bool("is_last", bool(bIsLast)),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnRspQryOptionInstrCommRate(
		(*thost.CThostFtdcOptionInstrCommRateField)(unsafe.Pointer(pOptionInstrCommRate)),
		(*thost.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
		int(nRequestID), bool(bIsLast),
	)
}

//export CgoOnRspQryExecOrder
func CgoOnRspQryExecOrder(
	this unsafe.Pointer,
	pExecOrder *C.struct_CThostFtdcExecOrderField,
	pRspInfo *C.struct_CThostFtdcRspInfoField,
	nRequestID C.int, bIsLast C.bool,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRspQryExecOrder called",
		slog.Any("this", this),
		slog.Int("request", int(nRequestID)),
		slog.Bool("is_last", bool(bIsLast)),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnRspQryExecOrder(
		(*thost.CThostFtdcExecOrderField)(unsafe.Pointer(pExecOrder)),
		(*thost.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
		int(nRequestID), bool(bIsLast),
	)
}

//export CgoOnRspQryForQuote
func CgoOnRspQryForQuote(
	this unsafe.Pointer,
	pForQuote *C.struct_CThostFtdcForQuoteField,
	pRspInfo *C.struct_CThostFtdcRspInfoField,
	nRequestID C.int, bIsLast C.bool,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRspQryForQuote called",
		slog.Any("this", this),
		slog.Int("request", int(nRequestID)),
		slog.Bool("is_last", bool(bIsLast)),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnRspQryForQuote(
		(*thost.CThostFtdcForQuoteField)(unsafe.Pointer(pForQuote)),
		(*thost.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
		int(nRequestID), bool(bIsLast),
	)
}

//export CgoOnRspQryQuote
func CgoOnRspQryQuote(
	this unsafe.Pointer,
	pQuote *C.struct_CThostFtdcQuoteField,
	pRspInfo *C.struct_CThostFtdcRspInfoField,
	nRequestID C.int, bIsLast C.bool,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRspQryQuote called",
		slog.Any("this", this),
		slog.Int("request", int(nRequestID)),
		slog.Bool("is_last", bool(bIsLast)),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnRspQryQuote(
		(*thost.CThostFtdcQuoteField)(unsafe.Pointer(pQuote)),
		(*thost.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
		int(nRequestID), bool(bIsLast),
	)
}

//export CgoOnRspQryOptionSelfClose
func CgoOnRspQryOptionSelfClose(
	this unsafe.Pointer,
	pOptionSelfClose *C.struct_CThostFtdcOptionSelfCloseField,
	pRspInfo *C.struct_CThostFtdcRspInfoField,
	nRequestID C.int, bIsLast C.bool,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRspQryOptionSelfClose called",
		slog.Any("this", this),
		slog.Int("request", int(nRequestID)),
		slog.Bool("is_last", bool(bIsLast)),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnRspQryOptionSelfClose(
		(*thost.CThostFtdcOptionSelfCloseField)(unsafe.Pointer(pOptionSelfClose)),
		(*thost.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
		int(nRequestID), bool(bIsLast),
	)
}

//export CgoOnRspQryInvestUnit
func CgoOnRspQryInvestUnit(
	this unsafe.Pointer,
	pInvestUnit *C.struct_CThostFtdcInvestUnitField,
	pRspInfo *C.struct_CThostFtdcRspInfoField,
	nRequestID C.int, bIsLast C.bool,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRspQryInvestUnit called",
		slog.Any("this", this),
		slog.Int("request", int(nRequestID)),
		slog.Bool("is_last", bool(bIsLast)),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnRspQryInvestUnit(
		(*thost.CThostFtdcInvestUnitField)(unsafe.Pointer(pInvestUnit)),
		(*thost.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
		int(nRequestID), bool(bIsLast),
	)
}

//export CgoOnRspQryCombInstrumentGuard
func CgoOnRspQryCombInstrumentGuard(
	this unsafe.Pointer,
	pCombInstrumentGuard *C.struct_CThostFtdcCombInstrumentGuardField,
	pRspInfo *C.struct_CThostFtdcRspInfoField,
	nRequestID C.int, bIsLast C.bool,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRspQryCombInstrumentGuard called",
		slog.Any("this", this),
		slog.Int("request", int(nRequestID)),
		slog.Bool("is_last", bool(bIsLast)),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnRspQryCombInstrumentGuard(
		(*thost.CThostFtdcCombInstrumentGuardField)(unsafe.Pointer(pCombInstrumentGuard)),
		(*thost.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
		int(nRequestID), bool(bIsLast),
	)
}

//export CgoOnRspQryCombAction
func CgoOnRspQryCombAction(
	this unsafe.Pointer,
	pCombAction *C.struct_CThostFtdcCombActionField,
	pRspInfo *C.struct_CThostFtdcRspInfoField,
	nRequestID C.int, bIsLast C.bool,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRspQryCombAction called",
		slog.Any("this", this),
		slog.Int("request", int(nRequestID)),
		slog.Bool("is_last", bool(bIsLast)),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnRspQryCombAction(
		(*thost.CThostFtdcCombActionField)(unsafe.Pointer(pCombAction)),
		(*thost.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
		int(nRequestID), bool(bIsLast),
	)
}

//export CgoOnRspQryTransferSerial
func CgoOnRspQryTransferSerial(
	this unsafe.Pointer,
	pTransferSerial *C.struct_CThostFtdcTransferSerialField,
	pRspInfo *C.struct_CThostFtdcRspInfoField,
	nRequestID C.int, bIsLast C.bool,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRspQryTransferSerial called",
		slog.Any("this", this),
		slog.Int("request", int(nRequestID)),
		slog.Bool("is_last", bool(bIsLast)),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnRspQryTransferSerial(
		(*thost.CThostFtdcTransferSerialField)(unsafe.Pointer(pTransferSerial)),
		(*thost.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
		int(nRequestID), bool(bIsLast),
	)
}

//export CgoOnRspQryAccountregister
func CgoOnRspQryAccountregister(
	this unsafe.Pointer,
	pAccountregister *C.struct_CThostFtdcAccountregisterField,
	pRspInfo *C.struct_CThostFtdcRspInfoField,
	nRequestID C.int, bIsLast C.bool,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRspQryAccountregister called",
		slog.Any("this", this),
		slog.Int("request", int(nRequestID)),
		slog.Bool("is_last", bool(bIsLast)),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnRspQryAccountregister(
		(*thost.CThostFtdcAccountregisterField)(unsafe.Pointer(pAccountregister)),
		(*thost.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
		int(nRequestID), bool(bIsLast),
	)
}

//export CgoOnRspError
func CgoOnRspError(
	this unsafe.Pointer,
	pRspInfo *C.struct_CThostFtdcRspInfoField,
	nRequestID C.int, bIsLast C.bool,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRspError called",
		slog.Any("this", this),
		slog.Int("error_id", int(pRspInfo.ErrorID)),
		slog.String("error_msg", thost.DecodeGBK(unsafe.Slice(
			(*byte)(unsafe.Pointer(&pRspInfo.ErrorMsg[0])),
			len(pRspInfo.ErrorMsg),
		))),
		slog.Int("request", int(nRequestID)),
		slog.Bool("is_last", bool(bIsLast)),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnRspError(
		(*thost.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
		int(nRequestID), bool(bIsLast),
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
		(*thost.CThostFtdcOrderField)(unsafe.Pointer(pOrder)),
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
		(*thost.CThostFtdcTradeField)(unsafe.Pointer(pTrade)),
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
		slog.Int("error_id", int(pRspInfo.ErrorID)),
		slog.String("error_msg", thost.DecodeGBK(unsafe.Slice(
			(*byte)(unsafe.Pointer(&pRspInfo.ErrorMsg[0])),
			len(pRspInfo.ErrorMsg),
		))),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnErrRtnOrderInsert(
		(*thost.CThostFtdcInputOrderField)(unsafe.Pointer(pInputOrder)),
		(*thost.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
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
		slog.Int("error_id", int(pRspInfo.ErrorID)),
		slog.String("error_msg", thost.DecodeGBK(unsafe.Slice(
			(*byte)(unsafe.Pointer(&pRspInfo.ErrorMsg[0])),
			len(pRspInfo.ErrorMsg),
		))),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnErrRtnOrderAction(
		(*thost.CThostFtdcOrderActionField)(unsafe.Pointer(pOrderAction)),
		(*thost.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
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
		(*thost.CThostFtdcInstrumentStatusField)(unsafe.Pointer(pInstrumentStatus)),
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
		(*thost.CThostFtdcBulletinField)(unsafe.Pointer(pBulletin)),
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
		(*thost.CThostFtdcTradingNoticeInfoField)(unsafe.Pointer(pTradingNoticeInfo)),
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
		(*thost.CThostFtdcErrorConditionalOrderField)(unsafe.Pointer(pErrorConditionalOrder)),
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
		(*thost.CThostFtdcExecOrderField)(unsafe.Pointer(pExecOrder)),
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
		slog.Int("error_id", int(pRspInfo.ErrorID)),
		slog.String("error_msg", thost.DecodeGBK(unsafe.Slice(
			(*byte)(unsafe.Pointer(&pRspInfo.ErrorMsg[0])),
			len(pRspInfo.ErrorMsg),
		))),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnErrRtnExecOrderInsert(
		(*thost.CThostFtdcInputExecOrderField)(unsafe.Pointer(pInputExecOrder)),
		(*thost.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
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
		slog.Int("error_id", int(pRspInfo.ErrorID)),
		slog.String("error_msg", thost.DecodeGBK(unsafe.Slice(
			(*byte)(unsafe.Pointer(&pRspInfo.ErrorMsg[0])),
			len(pRspInfo.ErrorMsg),
		))),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnErrRtnExecOrderAction(
		(*thost.CThostFtdcExecOrderActionField)(unsafe.Pointer(pExecOrderAction)),
		(*thost.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
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
		slog.Int("error_id", int(pRspInfo.ErrorID)),
		slog.String("error_msg", thost.DecodeGBK(unsafe.Slice(
			(*byte)(unsafe.Pointer(&pRspInfo.ErrorMsg[0])),
			len(pRspInfo.ErrorMsg),
		))),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnErrRtnForQuoteInsert(
		(*thost.CThostFtdcInputForQuoteField)(unsafe.Pointer(pInputForQuote)),
		(*thost.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
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
		(*thost.CThostFtdcQuoteField)(unsafe.Pointer(pQuote)),
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
		slog.Int("error_id", int(pRspInfo.ErrorID)),
		slog.String("error_msg", thost.DecodeGBK(unsafe.Slice(
			(*byte)(unsafe.Pointer(&pRspInfo.ErrorMsg[0])),
			len(pRspInfo.ErrorMsg),
		))),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnErrRtnQuoteInsert(
		(*thost.CThostFtdcInputQuoteField)(unsafe.Pointer(pInputQuote)),
		(*thost.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
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
		slog.Int("error_id", int(pRspInfo.ErrorID)),
		slog.String("error_msg", thost.DecodeGBK(unsafe.Slice(
			(*byte)(unsafe.Pointer(&pRspInfo.ErrorMsg[0])),
			len(pRspInfo.ErrorMsg),
		))),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnErrRtnQuoteAction(
		(*thost.CThostFtdcQuoteActionField)(unsafe.Pointer(pQuoteAction)),
		(*thost.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
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
		(*thost.CThostFtdcForQuoteRspField)(unsafe.Pointer(pForQuoteRsp)),
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
		(*thost.CThostFtdcCFMMCTradingAccountTokenField)(unsafe.Pointer(pCFMMCTradingAccountToken)),
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
		slog.Int("error_id", int(pRspInfo.ErrorID)),
		slog.String("error_msg", thost.DecodeGBK(unsafe.Slice(
			(*byte)(unsafe.Pointer(&pRspInfo.ErrorMsg[0])),
			len(pRspInfo.ErrorMsg),
		))),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnErrRtnBatchOrderAction(
		(*thost.CThostFtdcBatchOrderActionField)(unsafe.Pointer(pBatchOrderAction)),
		(*thost.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
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
		(*thost.CThostFtdcOptionSelfCloseField)(unsafe.Pointer(pOptionSelfClose)),
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
		slog.Int("error_id", int(pRspInfo.ErrorID)),
		slog.String("error_msg", thost.DecodeGBK(unsafe.Slice(
			(*byte)(unsafe.Pointer(&pRspInfo.ErrorMsg[0])),
			len(pRspInfo.ErrorMsg),
		))),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnErrRtnOptionSelfCloseInsert(
		(*thost.CThostFtdcInputOptionSelfCloseField)(unsafe.Pointer(pInputOptionSelfClose)),
		(*thost.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
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
		slog.Int("error_id", int(pRspInfo.ErrorID)),
		slog.String("error_msg", thost.DecodeGBK(unsafe.Slice(
			(*byte)(unsafe.Pointer(&pRspInfo.ErrorMsg[0])),
			len(pRspInfo.ErrorMsg),
		))),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnErrRtnOptionSelfCloseAction(
		(*thost.CThostFtdcOptionSelfCloseActionField)(unsafe.Pointer(pOptionSelfCloseAction)),
		(*thost.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
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
		(*thost.CThostFtdcCombActionField)(unsafe.Pointer(pCombAction)),
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
		slog.Int("error_id", int(pRspInfo.ErrorID)),
		slog.String("error_msg", thost.DecodeGBK(unsafe.Slice(
			(*byte)(unsafe.Pointer(&pRspInfo.ErrorMsg[0])),
			len(pRspInfo.ErrorMsg),
		))),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnErrRtnCombActionInsert(
		(*thost.CThostFtdcInputCombActionField)(unsafe.Pointer(pInputCombAction)),
		(*thost.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
	)
}

//export CgoOnRspQryContractBank
func CgoOnRspQryContractBank(
	this unsafe.Pointer,
	pContractBank *C.struct_CThostFtdcContractBankField,
	pRspInfo *C.struct_CThostFtdcRspInfoField,
	nRequestID C.int, bIsLast C.bool,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRspQryContractBank called",
		slog.Any("this", this),
		slog.Int("request", int(nRequestID)),
		slog.Bool("is_last", bool(bIsLast)),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnRspQryContractBank(
		(*thost.CThostFtdcContractBankField)(unsafe.Pointer(pContractBank)),
		(*thost.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
		int(nRequestID), bool(bIsLast),
	)
}

//export CgoOnRspQryParkedOrder
func CgoOnRspQryParkedOrder(
	this unsafe.Pointer,
	pParkedOrder *C.struct_CThostFtdcParkedOrderField,
	pRspInfo *C.struct_CThostFtdcRspInfoField,
	nRequestID C.int, bIsLast C.bool,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRspQryParkedOrder called",
		slog.Any("this", this),
		slog.Int("request", int(nRequestID)),
		slog.Bool("is_last", bool(bIsLast)),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnRspQryParkedOrder(
		(*thost.CThostFtdcParkedOrderField)(unsafe.Pointer(pParkedOrder)),
		(*thost.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
		int(nRequestID), bool(bIsLast),
	)
}

//export CgoOnRspQryParkedOrderAction
func CgoOnRspQryParkedOrderAction(
	this unsafe.Pointer,
	pParkedOrderAction *C.struct_CThostFtdcParkedOrderActionField,
	pRspInfo *C.struct_CThostFtdcRspInfoField,
	nRequestID C.int, bIsLast C.bool,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRspQryParkedOrderAction called",
		slog.Any("this", this),
		slog.Int("request", int(nRequestID)),
		slog.Bool("is_last", bool(bIsLast)),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnRspQryParkedOrderAction(
		(*thost.CThostFtdcParkedOrderActionField)(unsafe.Pointer(pParkedOrderAction)),
		(*thost.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
		int(nRequestID), bool(bIsLast),
	)
}

//export CgoOnRspQryTradingNotice
func CgoOnRspQryTradingNotice(
	this unsafe.Pointer,
	pTradingNotice *C.struct_CThostFtdcTradingNoticeField,
	pRspInfo *C.struct_CThostFtdcRspInfoField,
	nRequestID C.int, bIsLast C.bool,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRspQryTradingNotice called",
		slog.Any("this", this),
		slog.Int("request", int(nRequestID)),
		slog.Bool("is_last", bool(bIsLast)),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnRspQryTradingNotice(
		(*thost.CThostFtdcTradingNoticeField)(unsafe.Pointer(pTradingNotice)),
		(*thost.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
		int(nRequestID), bool(bIsLast),
	)
}

//export CgoOnRspQryBrokerTradingParams
func CgoOnRspQryBrokerTradingParams(
	this unsafe.Pointer,
	pBrokerTradingParams *C.struct_CThostFtdcBrokerTradingParamsField,
	pRspInfo *C.struct_CThostFtdcRspInfoField,
	nRequestID C.int, bIsLast C.bool,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRspQryBrokerTradingParams called",
		slog.Any("this", this),
		slog.Int("request", int(nRequestID)),
		slog.Bool("is_last", bool(bIsLast)),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnRspQryBrokerTradingParams(
		(*thost.CThostFtdcBrokerTradingParamsField)(unsafe.Pointer(pBrokerTradingParams)),
		(*thost.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
		int(nRequestID), bool(bIsLast),
	)
}

//export CgoOnRspQryBrokerTradingAlgos
func CgoOnRspQryBrokerTradingAlgos(
	this unsafe.Pointer,
	pBrokerTradingAlgos *C.struct_CThostFtdcBrokerTradingAlgosField,
	pRspInfo *C.struct_CThostFtdcRspInfoField,
	nRequestID C.int, bIsLast C.bool,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRspQryBrokerTradingAlgos called",
		slog.Any("this", this),
		slog.Int("request", int(nRequestID)),
		slog.Bool("is_last", bool(bIsLast)),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnRspQryBrokerTradingAlgos(
		(*thost.CThostFtdcBrokerTradingAlgosField)(unsafe.Pointer(pBrokerTradingAlgos)),
		(*thost.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
		int(nRequestID), bool(bIsLast),
	)
}

//export CgoOnRspQueryCFMMCTradingAccountToken
func CgoOnRspQueryCFMMCTradingAccountToken(
	this unsafe.Pointer,
	pQueryCFMMCTradingAccountToken *C.struct_CThostFtdcQueryCFMMCTradingAccountTokenField,
	pRspInfo *C.struct_CThostFtdcRspInfoField,
	nRequestID C.int, bIsLast C.bool,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRspQueryCFMMCTradingAccountToken called",
		slog.Any("this", this),
		slog.Int("request", int(nRequestID)),
		slog.Bool("is_last", bool(bIsLast)),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnRspQueryCFMMCTradingAccountToken(
		(*thost.CThostFtdcQueryCFMMCTradingAccountTokenField)(unsafe.Pointer(pQueryCFMMCTradingAccountToken)),
		(*thost.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
		int(nRequestID), bool(bIsLast),
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
		(*thost.CThostFtdcRspTransferField)(unsafe.Pointer(pRspTransfer)),
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
		(*thost.CThostFtdcRspTransferField)(unsafe.Pointer(pRspTransfer)),
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
		(*thost.CThostFtdcRspRepealField)(unsafe.Pointer(pRspRepeal)),
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
		(*thost.CThostFtdcRspRepealField)(unsafe.Pointer(pRspRepeal)),
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
		(*thost.CThostFtdcRspTransferField)(unsafe.Pointer(pRspTransfer)),
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
		(*thost.CThostFtdcRspTransferField)(unsafe.Pointer(pRspTransfer)),
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
		(*thost.CThostFtdcRspRepealField)(unsafe.Pointer(pRspRepeal)),
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
		(*thost.CThostFtdcRspRepealField)(unsafe.Pointer(pRspRepeal)),
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
		(*thost.CThostFtdcNotifyQueryAccountField)(unsafe.Pointer(pNotifyQueryAccount)),
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
		slog.Int("error_id", int(pRspInfo.ErrorID)),
		slog.String("error_msg", thost.DecodeGBK(unsafe.Slice(
			(*byte)(unsafe.Pointer(&pRspInfo.ErrorMsg[0])),
			len(pRspInfo.ErrorMsg),
		))),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnErrRtnBankToFutureByFuture(
		(*thost.CThostFtdcReqTransferField)(unsafe.Pointer(pReqTransfer)),
		(*thost.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
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
		slog.Int("error_id", int(pRspInfo.ErrorID)),
		slog.String("error_msg", thost.DecodeGBK(unsafe.Slice(
			(*byte)(unsafe.Pointer(&pRspInfo.ErrorMsg[0])),
			len(pRspInfo.ErrorMsg),
		))),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnErrRtnFutureToBankByFuture(
		(*thost.CThostFtdcReqTransferField)(unsafe.Pointer(pReqTransfer)),
		(*thost.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
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
		slog.Int("error_id", int(pRspInfo.ErrorID)),
		slog.String("error_msg", thost.DecodeGBK(unsafe.Slice(
			(*byte)(unsafe.Pointer(&pRspInfo.ErrorMsg[0])),
			len(pRspInfo.ErrorMsg),
		))),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnErrRtnRepealBankToFutureByFutureManual(
		(*thost.CThostFtdcReqRepealField)(unsafe.Pointer(pReqRepeal)),
		(*thost.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
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
		slog.Int("error_id", int(pRspInfo.ErrorID)),
		slog.String("error_msg", thost.DecodeGBK(unsafe.Slice(
			(*byte)(unsafe.Pointer(&pRspInfo.ErrorMsg[0])),
			len(pRspInfo.ErrorMsg),
		))),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnErrRtnRepealFutureToBankByFutureManual(
		(*thost.CThostFtdcReqRepealField)(unsafe.Pointer(pReqRepeal)),
		(*thost.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
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
		slog.Int("error_id", int(pRspInfo.ErrorID)),
		slog.String("error_msg", thost.DecodeGBK(unsafe.Slice(
			(*byte)(unsafe.Pointer(&pRspInfo.ErrorMsg[0])),
			len(pRspInfo.ErrorMsg),
		))),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnErrRtnQueryBankBalanceByFuture(
		(*thost.CThostFtdcReqQueryAccountField)(unsafe.Pointer(pReqQueryAccount)),
		(*thost.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
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
		(*thost.CThostFtdcRspRepealField)(unsafe.Pointer(pRspRepeal)),
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
		(*thost.CThostFtdcRspRepealField)(unsafe.Pointer(pRspRepeal)),
	)
}

//export CgoOnRspFromBankToFutureByFuture
func CgoOnRspFromBankToFutureByFuture(
	this unsafe.Pointer,
	pReqTransfer *C.struct_CThostFtdcReqTransferField,
	pRspInfo *C.struct_CThostFtdcRspInfoField,
	nRequestID C.int, bIsLast C.bool,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRspFromBankToFutureByFuture called",
		slog.Any("this", this),
		slog.Int("request", int(nRequestID)),
		slog.Bool("is_last", bool(bIsLast)),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnRspFromBankToFutureByFuture(
		(*thost.CThostFtdcReqTransferField)(unsafe.Pointer(pReqTransfer)),
		(*thost.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
		int(nRequestID), bool(bIsLast),
	)
}

//export CgoOnRspFromFutureToBankByFuture
func CgoOnRspFromFutureToBankByFuture(
	this unsafe.Pointer,
	pReqTransfer *C.struct_CThostFtdcReqTransferField,
	pRspInfo *C.struct_CThostFtdcRspInfoField,
	nRequestID C.int, bIsLast C.bool,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRspFromFutureToBankByFuture called",
		slog.Any("this", this),
		slog.Int("request", int(nRequestID)),
		slog.Bool("is_last", bool(bIsLast)),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnRspFromFutureToBankByFuture(
		(*thost.CThostFtdcReqTransferField)(unsafe.Pointer(pReqTransfer)),
		(*thost.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
		int(nRequestID), bool(bIsLast),
	)
}

//export CgoOnRspQueryBankAccountMoneyByFuture
func CgoOnRspQueryBankAccountMoneyByFuture(
	this unsafe.Pointer,
	pReqQueryAccount *C.struct_CThostFtdcReqQueryAccountField,
	pRspInfo *C.struct_CThostFtdcRspInfoField,
	nRequestID C.int, bIsLast C.bool,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRspQueryBankAccountMoneyByFuture called",
		slog.Any("this", this),
		slog.Int("request", int(nRequestID)),
		slog.Bool("is_last", bool(bIsLast)),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnRspQueryBankAccountMoneyByFuture(
		(*thost.CThostFtdcReqQueryAccountField)(unsafe.Pointer(pReqQueryAccount)),
		(*thost.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
		int(nRequestID), bool(bIsLast),
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
		(*thost.CThostFtdcOpenAccountField)(unsafe.Pointer(pOpenAccount)),
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
		(*thost.CThostFtdcCancelAccountField)(unsafe.Pointer(pCancelAccount)),
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
		(*thost.CThostFtdcChangeAccountField)(unsafe.Pointer(pChangeAccount)),
	)
}

//export CgoOnRspQryClassifiedInstrument
func CgoOnRspQryClassifiedInstrument(
	this unsafe.Pointer,
	pInstrument *C.struct_CThostFtdcInstrumentField,
	pRspInfo *C.struct_CThostFtdcRspInfoField,
	nRequestID C.int, bIsLast C.bool,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRspQryClassifiedInstrument called",
		slog.Any("this", this),
		slog.Int("request", int(nRequestID)),
		slog.Bool("is_last", bool(bIsLast)),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnRspQryClassifiedInstrument(
		(*thost.CThostFtdcInstrumentField)(unsafe.Pointer(pInstrument)),
		(*thost.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
		int(nRequestID), bool(bIsLast),
	)
}

//export CgoOnRspQryCombPromotionParam
func CgoOnRspQryCombPromotionParam(
	this unsafe.Pointer,
	pCombPromotionParam *C.struct_CThostFtdcCombPromotionParamField,
	pRspInfo *C.struct_CThostFtdcRspInfoField,
	nRequestID C.int, bIsLast C.bool,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRspQryCombPromotionParam called",
		slog.Any("this", this),
		slog.Int("request", int(nRequestID)),
		slog.Bool("is_last", bool(bIsLast)),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnRspQryCombPromotionParam(
		(*thost.CThostFtdcCombPromotionParamField)(unsafe.Pointer(pCombPromotionParam)),
		(*thost.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
		int(nRequestID), bool(bIsLast),
	)
}

//export CgoOnRspQryRiskSettleInvstPosition
func CgoOnRspQryRiskSettleInvstPosition(
	this unsafe.Pointer,
	pRiskSettleInvstPosition *C.struct_CThostFtdcRiskSettleInvstPositionField,
	pRspInfo *C.struct_CThostFtdcRspInfoField,
	nRequestID C.int, bIsLast C.bool,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRspQryRiskSettleInvstPosition called",
		slog.Any("this", this),
		slog.Int("request", int(nRequestID)),
		slog.Bool("is_last", bool(bIsLast)),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnRspQryRiskSettleInvstPosition(
		(*thost.CThostFtdcRiskSettleInvstPositionField)(unsafe.Pointer(pRiskSettleInvstPosition)),
		(*thost.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
		int(nRequestID), bool(bIsLast),
	)
}

//export CgoOnRspQryRiskSettleProductStatus
func CgoOnRspQryRiskSettleProductStatus(
	this unsafe.Pointer,
	pRiskSettleProductStatus *C.struct_CThostFtdcRiskSettleProductStatusField,
	pRspInfo *C.struct_CThostFtdcRspInfoField,
	nRequestID C.int, bIsLast C.bool,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRspQryRiskSettleProductStatus called",
		slog.Any("this", this),
		slog.Int("request", int(nRequestID)),
		slog.Bool("is_last", bool(bIsLast)),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnRspQryRiskSettleProductStatus(
		(*thost.CThostFtdcRiskSettleProductStatusField)(unsafe.Pointer(pRiskSettleProductStatus)),
		(*thost.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
		int(nRequestID), bool(bIsLast),
	)
}

//export CgoOnRspQrySPBMFutureParameter
func CgoOnRspQrySPBMFutureParameter(
	this unsafe.Pointer,
	pSPBMFutureParameter *C.struct_CThostFtdcSPBMFutureParameterField,
	pRspInfo *C.struct_CThostFtdcRspInfoField,
	nRequestID C.int, bIsLast C.bool,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRspQrySPBMFutureParameter called",
		slog.Any("this", this),
		slog.Int("request", int(nRequestID)),
		slog.Bool("is_last", bool(bIsLast)),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnRspQrySPBMFutureParameter(
		(*thost.CThostFtdcSPBMFutureParameterField)(unsafe.Pointer(pSPBMFutureParameter)),
		(*thost.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
		int(nRequestID), bool(bIsLast),
	)
}

//export CgoOnRspQrySPBMOptionParameter
func CgoOnRspQrySPBMOptionParameter(
	this unsafe.Pointer,
	pSPBMOptionParameter *C.struct_CThostFtdcSPBMOptionParameterField,
	pRspInfo *C.struct_CThostFtdcRspInfoField,
	nRequestID C.int, bIsLast C.bool,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRspQrySPBMOptionParameter called",
		slog.Any("this", this),
		slog.Int("request", int(nRequestID)),
		slog.Bool("is_last", bool(bIsLast)),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnRspQrySPBMOptionParameter(
		(*thost.CThostFtdcSPBMOptionParameterField)(unsafe.Pointer(pSPBMOptionParameter)),
		(*thost.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
		int(nRequestID), bool(bIsLast),
	)
}

//export CgoOnRspQrySPBMIntraParameter
func CgoOnRspQrySPBMIntraParameter(
	this unsafe.Pointer,
	pSPBMIntraParameter *C.struct_CThostFtdcSPBMIntraParameterField,
	pRspInfo *C.struct_CThostFtdcRspInfoField,
	nRequestID C.int, bIsLast C.bool,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRspQrySPBMIntraParameter called",
		slog.Any("this", this),
		slog.Int("request", int(nRequestID)),
		slog.Bool("is_last", bool(bIsLast)),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnRspQrySPBMIntraParameter(
		(*thost.CThostFtdcSPBMIntraParameterField)(unsafe.Pointer(pSPBMIntraParameter)),
		(*thost.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
		int(nRequestID), bool(bIsLast),
	)
}

//export CgoOnRspQrySPBMInterParameter
func CgoOnRspQrySPBMInterParameter(
	this unsafe.Pointer,
	pSPBMInterParameter *C.struct_CThostFtdcSPBMInterParameterField,
	pRspInfo *C.struct_CThostFtdcRspInfoField,
	nRequestID C.int, bIsLast C.bool,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRspQrySPBMInterParameter called",
		slog.Any("this", this),
		slog.Int("request", int(nRequestID)),
		slog.Bool("is_last", bool(bIsLast)),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnRspQrySPBMInterParameter(
		(*thost.CThostFtdcSPBMInterParameterField)(unsafe.Pointer(pSPBMInterParameter)),
		(*thost.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
		int(nRequestID), bool(bIsLast),
	)
}

//export CgoOnRspQrySPBMPortfDefinition
func CgoOnRspQrySPBMPortfDefinition(
	this unsafe.Pointer,
	pSPBMPortfDefinition *C.struct_CThostFtdcSPBMPortfDefinitionField,
	pRspInfo *C.struct_CThostFtdcRspInfoField,
	nRequestID C.int, bIsLast C.bool,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRspQrySPBMPortfDefinition called",
		slog.Any("this", this),
		slog.Int("request", int(nRequestID)),
		slog.Bool("is_last", bool(bIsLast)),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnRspQrySPBMPortfDefinition(
		(*thost.CThostFtdcSPBMPortfDefinitionField)(unsafe.Pointer(pSPBMPortfDefinition)),
		(*thost.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
		int(nRequestID), bool(bIsLast),
	)
}

//export CgoOnRspQrySPBMInvestorPortfDef
func CgoOnRspQrySPBMInvestorPortfDef(
	this unsafe.Pointer,
	pSPBMInvestorPortfDef *C.struct_CThostFtdcSPBMInvestorPortfDefField,
	pRspInfo *C.struct_CThostFtdcRspInfoField,
	nRequestID C.int, bIsLast C.bool,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRspQrySPBMInvestorPortfDef called",
		slog.Any("this", this),
		slog.Int("request", int(nRequestID)),
		slog.Bool("is_last", bool(bIsLast)),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnRspQrySPBMInvestorPortfDef(
		(*thost.CThostFtdcSPBMInvestorPortfDefField)(unsafe.Pointer(pSPBMInvestorPortfDef)),
		(*thost.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
		int(nRequestID), bool(bIsLast),
	)
}

//export CgoOnRspQryInvestorPortfMarginRatio
func CgoOnRspQryInvestorPortfMarginRatio(
	this unsafe.Pointer,
	pInvestorPortfMarginRatio *C.struct_CThostFtdcInvestorPortfMarginRatioField,
	pRspInfo *C.struct_CThostFtdcRspInfoField,
	nRequestID C.int, bIsLast C.bool,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRspQryInvestorPortfMarginRatio called",
		slog.Any("this", this),
		slog.Int("request", int(nRequestID)),
		slog.Bool("is_last", bool(bIsLast)),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnRspQryInvestorPortfMarginRatio(
		(*thost.CThostFtdcInvestorPortfMarginRatioField)(unsafe.Pointer(pInvestorPortfMarginRatio)),
		(*thost.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
		int(nRequestID), bool(bIsLast),
	)
}

//export CgoOnRspQryInvestorProdSPBMDetail
func CgoOnRspQryInvestorProdSPBMDetail(
	this unsafe.Pointer,
	pInvestorProdSPBMDetail *C.struct_CThostFtdcInvestorProdSPBMDetailField,
	pRspInfo *C.struct_CThostFtdcRspInfoField,
	nRequestID C.int, bIsLast C.bool,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRspQryInvestorProdSPBMDetail called",
		slog.Any("this", this),
		slog.Int("request", int(nRequestID)),
		slog.Bool("is_last", bool(bIsLast)),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnRspQryInvestorProdSPBMDetail(
		(*thost.CThostFtdcInvestorProdSPBMDetailField)(unsafe.Pointer(pInvestorProdSPBMDetail)),
		(*thost.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
		int(nRequestID), bool(bIsLast),
	)
}

//export CgoOnRspQryInvestorCommoditySPMMMargin
func CgoOnRspQryInvestorCommoditySPMMMargin(
	this unsafe.Pointer,
	pInvestorCommoditySPMMMargin *C.struct_CThostFtdcInvestorCommoditySPMMMarginField,
	pRspInfo *C.struct_CThostFtdcRspInfoField,
	nRequestID C.int, bIsLast C.bool,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRspQryInvestorCommoditySPMMMargin called",
		slog.Any("this", this),
		slog.Int("request", int(nRequestID)),
		slog.Bool("is_last", bool(bIsLast)),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnRspQryInvestorCommoditySPMMMargin(
		(*thost.CThostFtdcInvestorCommoditySPMMMarginField)(unsafe.Pointer(pInvestorCommoditySPMMMargin)),
		(*thost.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
		int(nRequestID), bool(bIsLast),
	)
}

//export CgoOnRspQryInvestorCommodityGroupSPMMMargin
func CgoOnRspQryInvestorCommodityGroupSPMMMargin(
	this unsafe.Pointer,
	pInvestorCommodityGroupSPMMMargin *C.struct_CThostFtdcInvestorCommodityGroupSPMMMarginField,
	pRspInfo *C.struct_CThostFtdcRspInfoField,
	nRequestID C.int, bIsLast C.bool,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRspQryInvestorCommodityGroupSPMMMargin called",
		slog.Any("this", this),
		slog.Int("request", int(nRequestID)),
		slog.Bool("is_last", bool(bIsLast)),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnRspQryInvestorCommodityGroupSPMMMargin(
		(*thost.CThostFtdcInvestorCommodityGroupSPMMMarginField)(unsafe.Pointer(pInvestorCommodityGroupSPMMMargin)),
		(*thost.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
		int(nRequestID), bool(bIsLast),
	)
}

//export CgoOnRspQrySPMMInstParam
func CgoOnRspQrySPMMInstParam(
	this unsafe.Pointer,
	pSPMMInstParam *C.struct_CThostFtdcSPMMInstParamField,
	pRspInfo *C.struct_CThostFtdcRspInfoField,
	nRequestID C.int, bIsLast C.bool,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRspQrySPMMInstParam called",
		slog.Any("this", this),
		slog.Int("request", int(nRequestID)),
		slog.Bool("is_last", bool(bIsLast)),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnRspQrySPMMInstParam(
		(*thost.CThostFtdcSPMMInstParamField)(unsafe.Pointer(pSPMMInstParam)),
		(*thost.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
		int(nRequestID), bool(bIsLast),
	)
}

//export CgoOnRspQrySPMMProductParam
func CgoOnRspQrySPMMProductParam(
	this unsafe.Pointer,
	pSPMMProductParam *C.struct_CThostFtdcSPMMProductParamField,
	pRspInfo *C.struct_CThostFtdcRspInfoField,
	nRequestID C.int, bIsLast C.bool,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRspQrySPMMProductParam called",
		slog.Any("this", this),
		slog.Int("request", int(nRequestID)),
		slog.Bool("is_last", bool(bIsLast)),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnRspQrySPMMProductParam(
		(*thost.CThostFtdcSPMMProductParamField)(unsafe.Pointer(pSPMMProductParam)),
		(*thost.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
		int(nRequestID), bool(bIsLast),
	)
}

//export CgoOnRspQrySPBMAddOnInterParameter
func CgoOnRspQrySPBMAddOnInterParameter(
	this unsafe.Pointer,
	pSPBMAddOnInterParameter *C.struct_CThostFtdcSPBMAddOnInterParameterField,
	pRspInfo *C.struct_CThostFtdcRspInfoField,
	nRequestID C.int, bIsLast C.bool,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRspQrySPBMAddOnInterParameter called",
		slog.Any("this", this),
		slog.Int("request", int(nRequestID)),
		slog.Bool("is_last", bool(bIsLast)),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnRspQrySPBMAddOnInterParameter(
		(*thost.CThostFtdcSPBMAddOnInterParameterField)(unsafe.Pointer(pSPBMAddOnInterParameter)),
		(*thost.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
		int(nRequestID), bool(bIsLast),
	)
}

//export CgoOnRspQryRCAMSCombProductInfo
func CgoOnRspQryRCAMSCombProductInfo(
	this unsafe.Pointer,
	pRCAMSCombProductInfo *C.struct_CThostFtdcRCAMSCombProductInfoField,
	pRspInfo *C.struct_CThostFtdcRspInfoField,
	nRequestID C.int, bIsLast C.bool,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRspQryRCAMSCombProductInfo called",
		slog.Any("this", this),
		slog.Int("request", int(nRequestID)),
		slog.Bool("is_last", bool(bIsLast)),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnRspQryRCAMSCombProductInfo(
		(*thost.CThostFtdcRCAMSCombProductInfoField)(unsafe.Pointer(pRCAMSCombProductInfo)),
		(*thost.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
		int(nRequestID), bool(bIsLast),
	)
}

//export CgoOnRspQryRCAMSInstrParameter
func CgoOnRspQryRCAMSInstrParameter(
	this unsafe.Pointer,
	pRCAMSInstrParameter *C.struct_CThostFtdcRCAMSInstrParameterField,
	pRspInfo *C.struct_CThostFtdcRspInfoField,
	nRequestID C.int, bIsLast C.bool,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRspQryRCAMSInstrParameter called",
		slog.Any("this", this),
		slog.Int("request", int(nRequestID)),
		slog.Bool("is_last", bool(bIsLast)),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnRspQryRCAMSInstrParameter(
		(*thost.CThostFtdcRCAMSInstrParameterField)(unsafe.Pointer(pRCAMSInstrParameter)),
		(*thost.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
		int(nRequestID), bool(bIsLast),
	)
}

//export CgoOnRspQryRCAMSIntraParameter
func CgoOnRspQryRCAMSIntraParameter(
	this unsafe.Pointer,
	pRCAMSIntraParameter *C.struct_CThostFtdcRCAMSIntraParameterField,
	pRspInfo *C.struct_CThostFtdcRspInfoField,
	nRequestID C.int, bIsLast C.bool,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRspQryRCAMSIntraParameter called",
		slog.Any("this", this),
		slog.Int("request", int(nRequestID)),
		slog.Bool("is_last", bool(bIsLast)),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnRspQryRCAMSIntraParameter(
		(*thost.CThostFtdcRCAMSIntraParameterField)(unsafe.Pointer(pRCAMSIntraParameter)),
		(*thost.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
		int(nRequestID), bool(bIsLast),
	)
}

//export CgoOnRspQryRCAMSInterParameter
func CgoOnRspQryRCAMSInterParameter(
	this unsafe.Pointer,
	pRCAMSInterParameter *C.struct_CThostFtdcRCAMSInterParameterField,
	pRspInfo *C.struct_CThostFtdcRspInfoField,
	nRequestID C.int, bIsLast C.bool,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRspQryRCAMSInterParameter called",
		slog.Any("this", this),
		slog.Int("request", int(nRequestID)),
		slog.Bool("is_last", bool(bIsLast)),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnRspQryRCAMSInterParameter(
		(*thost.CThostFtdcRCAMSInterParameterField)(unsafe.Pointer(pRCAMSInterParameter)),
		(*thost.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
		int(nRequestID), bool(bIsLast),
	)
}

//export CgoOnRspQryRCAMSShortOptAdjustParam
func CgoOnRspQryRCAMSShortOptAdjustParam(
	this unsafe.Pointer,
	pRCAMSShortOptAdjustParam *C.struct_CThostFtdcRCAMSShortOptAdjustParamField,
	pRspInfo *C.struct_CThostFtdcRspInfoField,
	nRequestID C.int, bIsLast C.bool,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRspQryRCAMSShortOptAdjustParam called",
		slog.Any("this", this),
		slog.Int("request", int(nRequestID)),
		slog.Bool("is_last", bool(bIsLast)),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnRspQryRCAMSShortOptAdjustParam(
		(*thost.CThostFtdcRCAMSShortOptAdjustParamField)(unsafe.Pointer(pRCAMSShortOptAdjustParam)),
		(*thost.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
		int(nRequestID), bool(bIsLast),
	)
}

//export CgoOnRspQryRCAMSInvestorCombPosition
func CgoOnRspQryRCAMSInvestorCombPosition(
	this unsafe.Pointer,
	pRCAMSInvestorCombPosition *C.struct_CThostFtdcRCAMSInvestorCombPositionField,
	pRspInfo *C.struct_CThostFtdcRspInfoField,
	nRequestID C.int, bIsLast C.bool,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRspQryRCAMSInvestorCombPosition called",
		slog.Any("this", this),
		slog.Int("request", int(nRequestID)),
		slog.Bool("is_last", bool(bIsLast)),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnRspQryRCAMSInvestorCombPosition(
		(*thost.CThostFtdcRCAMSInvestorCombPositionField)(unsafe.Pointer(pRCAMSInvestorCombPosition)),
		(*thost.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
		int(nRequestID), bool(bIsLast),
	)
}

//export CgoOnRspQryInvestorProdRCAMSMargin
func CgoOnRspQryInvestorProdRCAMSMargin(
	this unsafe.Pointer,
	pInvestorProdRCAMSMargin *C.struct_CThostFtdcInvestorProdRCAMSMarginField,
	pRspInfo *C.struct_CThostFtdcRspInfoField,
	nRequestID C.int, bIsLast C.bool,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRspQryInvestorProdRCAMSMargin called",
		slog.Any("this", this),
		slog.Int("request", int(nRequestID)),
		slog.Bool("is_last", bool(bIsLast)),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnRspQryInvestorProdRCAMSMargin(
		(*thost.CThostFtdcInvestorProdRCAMSMarginField)(unsafe.Pointer(pInvestorProdRCAMSMargin)),
		(*thost.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
		int(nRequestID), bool(bIsLast),
	)
}

//export CgoOnRspQryRULEInstrParameter
func CgoOnRspQryRULEInstrParameter(
	this unsafe.Pointer,
	pRULEInstrParameter *C.struct_CThostFtdcRULEInstrParameterField,
	pRspInfo *C.struct_CThostFtdcRspInfoField,
	nRequestID C.int, bIsLast C.bool,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRspQryRULEInstrParameter called",
		slog.Any("this", this),
		slog.Int("request", int(nRequestID)),
		slog.Bool("is_last", bool(bIsLast)),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnRspQryRULEInstrParameter(
		(*thost.CThostFtdcRULEInstrParameterField)(unsafe.Pointer(pRULEInstrParameter)),
		(*thost.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
		int(nRequestID), bool(bIsLast),
	)
}

//export CgoOnRspQryRULEIntraParameter
func CgoOnRspQryRULEIntraParameter(
	this unsafe.Pointer,
	pRULEIntraParameter *C.struct_CThostFtdcRULEIntraParameterField,
	pRspInfo *C.struct_CThostFtdcRspInfoField,
	nRequestID C.int, bIsLast C.bool,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRspQryRULEIntraParameter called",
		slog.Any("this", this),
		slog.Int("request", int(nRequestID)),
		slog.Bool("is_last", bool(bIsLast)),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnRspQryRULEIntraParameter(
		(*thost.CThostFtdcRULEIntraParameterField)(unsafe.Pointer(pRULEIntraParameter)),
		(*thost.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
		int(nRequestID), bool(bIsLast),
	)
}

//export CgoOnRspQryRULEInterParameter
func CgoOnRspQryRULEInterParameter(
	this unsafe.Pointer,
	pRULEInterParameter *C.struct_CThostFtdcRULEInterParameterField,
	pRspInfo *C.struct_CThostFtdcRspInfoField,
	nRequestID C.int, bIsLast C.bool,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRspQryRULEInterParameter called",
		slog.Any("this", this),
		slog.Int("request", int(nRequestID)),
		slog.Bool("is_last", bool(bIsLast)),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnRspQryRULEInterParameter(
		(*thost.CThostFtdcRULEInterParameterField)(unsafe.Pointer(pRULEInterParameter)),
		(*thost.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
		int(nRequestID), bool(bIsLast),
	)
}

//export CgoOnRspQryInvestorProdRULEMargin
func CgoOnRspQryInvestorProdRULEMargin(
	this unsafe.Pointer,
	pInvestorProdRULEMargin *C.struct_CThostFtdcInvestorProdRULEMarginField,
	pRspInfo *C.struct_CThostFtdcRspInfoField,
	nRequestID C.int, bIsLast C.bool,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRspQryInvestorProdRULEMargin called",
		slog.Any("this", this),
		slog.Int("request", int(nRequestID)),
		slog.Bool("is_last", bool(bIsLast)),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnRspQryInvestorProdRULEMargin(
		(*thost.CThostFtdcInvestorProdRULEMarginField)(unsafe.Pointer(pInvestorProdRULEMargin)),
		(*thost.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
		int(nRequestID), bool(bIsLast),
	)
}

//export CgoOnRspQryInvestorPortfSetting
func CgoOnRspQryInvestorPortfSetting(
	this unsafe.Pointer,
	pInvestorPortfSetting *C.struct_CThostFtdcInvestorPortfSettingField,
	pRspInfo *C.struct_CThostFtdcRspInfoField,
	nRequestID C.int, bIsLast C.bool,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRspQryInvestorPortfSetting called",
		slog.Any("this", this),
		slog.Int("request", int(nRequestID)),
		slog.Bool("is_last", bool(bIsLast)),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnRspQryInvestorPortfSetting(
		(*thost.CThostFtdcInvestorPortfSettingField)(unsafe.Pointer(pInvestorPortfSetting)),
		(*thost.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
		int(nRequestID), bool(bIsLast),
	)
}

//export CgoOnRspQryInvestorInfoCommRec
func CgoOnRspQryInvestorInfoCommRec(
	this unsafe.Pointer,
	pInvestorInfoCommRec *C.struct_CThostFtdcInvestorInfoCommRecField,
	pRspInfo *C.struct_CThostFtdcRspInfoField,
	nRequestID C.int, bIsLast C.bool,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRspQryInvestorInfoCommRec called",
		slog.Any("this", this),
		slog.Int("request", int(nRequestID)),
		slog.Bool("is_last", bool(bIsLast)),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnRspQryInvestorInfoCommRec(
		(*thost.CThostFtdcInvestorInfoCommRecField)(unsafe.Pointer(pInvestorInfoCommRec)),
		(*thost.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
		int(nRequestID), bool(bIsLast),
	)
}

//export CgoOnRspQryCombLeg
func CgoOnRspQryCombLeg(
	this unsafe.Pointer,
	pCombLeg *C.struct_CThostFtdcCombLegField,
	pRspInfo *C.struct_CThostFtdcRspInfoField,
	nRequestID C.int, bIsLast C.bool,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRspQryCombLeg called",
		slog.Any("this", this),
		slog.Int("request", int(nRequestID)),
		slog.Bool("is_last", bool(bIsLast)),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnRspQryCombLeg(
		(*thost.CThostFtdcCombLegField)(unsafe.Pointer(pCombLeg)),
		(*thost.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
		int(nRequestID), bool(bIsLast),
	)
}

//export CgoOnRspOffsetSetting
func CgoOnRspOffsetSetting(
	this unsafe.Pointer,
	pInputOffsetSetting *C.struct_CThostFtdcInputOffsetSettingField,
	pRspInfo *C.struct_CThostFtdcRspInfoField,
	nRequestID C.int, bIsLast C.bool,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRspOffsetSetting called",
		slog.Any("this", this),
		slog.Int("request", int(nRequestID)),
		slog.Bool("is_last", bool(bIsLast)),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnRspOffsetSetting(
		(*thost.CThostFtdcInputOffsetSettingField)(unsafe.Pointer(pInputOffsetSetting)),
		(*thost.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
		int(nRequestID), bool(bIsLast),
	)
}

//export CgoOnRspCancelOffsetSetting
func CgoOnRspCancelOffsetSetting(
	this unsafe.Pointer,
	pInputOffsetSetting *C.struct_CThostFtdcInputOffsetSettingField,
	pRspInfo *C.struct_CThostFtdcRspInfoField,
	nRequestID C.int, bIsLast C.bool,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRspCancelOffsetSetting called",
		slog.Any("this", this),
		slog.Int("request", int(nRequestID)),
		slog.Bool("is_last", bool(bIsLast)),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnRspCancelOffsetSetting(
		(*thost.CThostFtdcInputOffsetSettingField)(unsafe.Pointer(pInputOffsetSetting)),
		(*thost.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
		int(nRequestID), bool(bIsLast),
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
		(*thost.CThostFtdcOffsetSettingField)(unsafe.Pointer(pOffsetSetting)),
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
		slog.Int("error_id", int(pRspInfo.ErrorID)),
		slog.String("error_msg", thost.DecodeGBK(unsafe.Slice(
			(*byte)(unsafe.Pointer(&pRspInfo.ErrorMsg[0])),
			len(pRspInfo.ErrorMsg),
		))),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnErrRtnOffsetSetting(
		(*thost.CThostFtdcInputOffsetSettingField)(unsafe.Pointer(pInputOffsetSetting)),
		(*thost.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
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
		slog.Int("error_id", int(pRspInfo.ErrorID)),
		slog.String("error_msg", thost.DecodeGBK(unsafe.Slice(
			(*byte)(unsafe.Pointer(&pRspInfo.ErrorMsg[0])),
			len(pRspInfo.ErrorMsg),
		))),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnErrRtnCancelOffsetSetting(
		(*thost.CThostFtdcCancelOffsetSettingField)(unsafe.Pointer(pCancelOffsetSetting)),
		(*thost.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
	)
}

//export CgoOnRspQryOffsetSetting
func CgoOnRspQryOffsetSetting(
	this unsafe.Pointer,
	pOffsetSetting *C.struct_CThostFtdcOffsetSettingField,
	pRspInfo *C.struct_CThostFtdcRspInfoField,
	nRequestID C.int, bIsLast C.bool,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRspQryOffsetSetting called",
		slog.Any("this", this),
		slog.Int("request", int(nRequestID)),
		slog.Bool("is_last", bool(bIsLast)),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnRspQryOffsetSetting(
		(*thost.CThostFtdcOffsetSettingField)(unsafe.Pointer(pOffsetSetting)),
		(*thost.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
		int(nRequestID), bool(bIsLast),
	)
}

//export CgoOnRspGenSMSCode
func CgoOnRspGenSMSCode(
	this unsafe.Pointer,
	pRspGenSMSCode *C.struct_CThostFtdcRspGenSMSCodeField,
	pRspInfo *C.struct_CThostFtdcRspInfoField,
	nRequestID C.int, bIsLast C.bool,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRspGenSMSCode called",
		slog.Any("this", this),
		slog.Int("request", int(nRequestID)),
		slog.Bool("is_last", bool(bIsLast)),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnRspGenSMSCode(
		(*thost.CThostFtdcRspGenSMSCodeField)(unsafe.Pointer(pRspGenSMSCode)),
		(*thost.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
		int(nRequestID), bool(bIsLast),
	)
}

//export CgoOnRspSpdApply
func CgoOnRspSpdApply(
	this unsafe.Pointer,
	pInputSpdApply *C.struct_CThostFtdcInputSpdApplyField,
	pRspInfo *C.struct_CThostFtdcRspInfoField,
	nRequestID C.int, bIsLast C.bool,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRspSpdApply called",
		slog.Any("this", this),
		slog.Int("request", int(nRequestID)),
		slog.Bool("is_last", bool(bIsLast)),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnRspSpdApply(
		(*thost.CThostFtdcInputSpdApplyField)(unsafe.Pointer(pInputSpdApply)),
		(*thost.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
		int(nRequestID), bool(bIsLast),
	)
}

//export CgoOnRspSpdApplyAction
func CgoOnRspSpdApplyAction(
	this unsafe.Pointer,
	pInputSpdApplyAction *C.struct_CThostFtdcInputSpdApplyActionField,
	pRspInfo *C.struct_CThostFtdcRspInfoField,
	nRequestID C.int, bIsLast C.bool,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRspSpdApplyAction called",
		slog.Any("this", this),
		slog.Int("request", int(nRequestID)),
		slog.Bool("is_last", bool(bIsLast)),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnRspSpdApplyAction(
		(*thost.CThostFtdcInputSpdApplyActionField)(unsafe.Pointer(pInputSpdApplyAction)),
		(*thost.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
		int(nRequestID), bool(bIsLast),
	)
}

//export CgoOnRspQrySpdApply
func CgoOnRspQrySpdApply(
	this unsafe.Pointer,
	pSpdApply *C.struct_CThostFtdcSpdApplyField,
	pRspInfo *C.struct_CThostFtdcRspInfoField,
	nRequestID C.int, bIsLast C.bool,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRspQrySpdApply called",
		slog.Any("this", this),
		slog.Int("request", int(nRequestID)),
		slog.Bool("is_last", bool(bIsLast)),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnRspQrySpdApply(
		(*thost.CThostFtdcSpdApplyField)(unsafe.Pointer(pSpdApply)),
		(*thost.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
		int(nRequestID), bool(bIsLast),
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
		(*thost.CThostFtdcSpdApplyField)(unsafe.Pointer(pSpdApply)),
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
		slog.Int("error_id", int(pRspInfo.ErrorID)),
		slog.String("error_msg", thost.DecodeGBK(unsafe.Slice(
			(*byte)(unsafe.Pointer(&pRspInfo.ErrorMsg[0])),
			len(pRspInfo.ErrorMsg),
		))),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnErrRtnSpdApply(
		(*thost.CThostFtdcInputSpdApplyField)(unsafe.Pointer(pInputSpdApply)),
		(*thost.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
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
		slog.Int("error_id", int(pRspInfo.ErrorID)),
		slog.String("error_msg", thost.DecodeGBK(unsafe.Slice(
			(*byte)(unsafe.Pointer(&pRspInfo.ErrorMsg[0])),
			len(pRspInfo.ErrorMsg),
		))),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnErrRtnSpdApplyAction(
		(*thost.CThostFtdcSpdApplyActionField)(unsafe.Pointer(pSpdApplyAction)),
		(*thost.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
	)
}

//export CgoOnRspHedgeCfm
func CgoOnRspHedgeCfm(
	this unsafe.Pointer,
	pInputHedgeCfm *C.struct_CThostFtdcInputHedgeCfmField,
	pRspInfo *C.struct_CThostFtdcRspInfoField,
	nRequestID C.int, bIsLast C.bool,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRspHedgeCfm called",
		slog.Any("this", this),
		slog.Int("request", int(nRequestID)),
		slog.Bool("is_last", bool(bIsLast)),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnRspHedgeCfm(
		(*thost.CThostFtdcInputHedgeCfmField)(unsafe.Pointer(pInputHedgeCfm)),
		(*thost.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
		int(nRequestID), bool(bIsLast),
	)
}

//export CgoOnRspHedgeCfmAction
func CgoOnRspHedgeCfmAction(
	this unsafe.Pointer,
	pInputHedgeCfmAction *C.struct_CThostFtdcInputHedgeCfmActionField,
	pRspInfo *C.struct_CThostFtdcRspInfoField,
	nRequestID C.int, bIsLast C.bool,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRspHedgeCfmAction called",
		slog.Any("this", this),
		slog.Int("request", int(nRequestID)),
		slog.Bool("is_last", bool(bIsLast)),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnRspHedgeCfmAction(
		(*thost.CThostFtdcInputHedgeCfmActionField)(unsafe.Pointer(pInputHedgeCfmAction)),
		(*thost.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
		int(nRequestID), bool(bIsLast),
	)
}

//export CgoOnRspQryHedgeCfm
func CgoOnRspQryHedgeCfm(
	this unsafe.Pointer,
	pHedgeCfm *C.struct_CThostFtdcHedgeCfmField,
	pRspInfo *C.struct_CThostFtdcRspInfoField,
	nRequestID C.int, bIsLast C.bool,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"CgoOnRspQryHedgeCfm called",
		slog.Any("this", this),
		slog.Int("request", int(nRequestID)),
		slog.Bool("is_last", bool(bIsLast)),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnRspQryHedgeCfm(
		(*thost.CThostFtdcHedgeCfmField)(unsafe.Pointer(pHedgeCfm)),
		(*thost.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
		int(nRequestID), bool(bIsLast),
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
		(*thost.CThostFtdcHedgeCfmField)(unsafe.Pointer(pHedgeCfm)),
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
		slog.Int("error_id", int(pRspInfo.ErrorID)),
		slog.String("error_msg", thost.DecodeGBK(unsafe.Slice(
			(*byte)(unsafe.Pointer(&pRspInfo.ErrorMsg[0])),
			len(pRspInfo.ErrorMsg),
		))),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnErrRtnHedgeCfm(
		(*thost.CThostFtdcInputHedgeCfmField)(unsafe.Pointer(pInputHedgeCfm)),
		(*thost.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
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
		slog.Int("error_id", int(pRspInfo.ErrorID)),
		slog.String("error_msg", thost.DecodeGBK(unsafe.Slice(
			(*byte)(unsafe.Pointer(&pRspInfo.ErrorMsg[0])),
			len(pRspInfo.ErrorMsg),
		))),
	)

	(*ThostFtdcTraderSpi)(
		(*C.CThostFtdcTraderSpiExt)(this).spi,
	).callback.OnErrRtnHedgeCfmAction(
		(*thost.CThostFtdcHedgeCfmActionField)(unsafe.Pointer(pHedgeCfmAction)),
		(*thost.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
	)
}
