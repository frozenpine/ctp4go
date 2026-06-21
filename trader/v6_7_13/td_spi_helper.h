#pragma once

#ifndef TRADER_V6713_SPI_HELPER_H
#define TRADER_V6713_SPI_HELPER_H

#ifdef __cplusplus
extern "C"
{
#endif

#include <stdbool.h>
#include <stdint.h>
#include <stdlib.h>
#include <stdio.h>
#include <string.h>
#include <dlfcn.h>

#include "ThostFtdcUserApiStruct.h"

    typedef void (*OnFrontConnected)();

    typedef void (*OnFrontDisconnected)(int nReason);

    typedef void (*OnHeartBeatWarning)(int nTimeLapse);

    typedef void (*OnRspAuthenticate)(struct CThostFtdcRspAuthenticateField *pRspAuthenticateField, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

    typedef void (*OnRtnPrivateSeqNo)(int nSeqNo);

    typedef void (*OnRspUserLogin)(struct CThostFtdcRspUserLoginField *pRspUserLogin, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

    typedef void (*OnRspUserLogout)(struct CThostFtdcUserLogoutField *pUserLogout, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

    typedef void (*OnRspUserPasswordUpdate)(struct CThostFtdcUserPasswordUpdateField *pUserPasswordUpdate, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

    typedef void (*OnRspTradingAccountPasswordUpdate)(struct CThostFtdcTradingAccountPasswordUpdateField *pTradingAccountPasswordUpdate, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

    typedef void (*OnRspUserAuthMethod)(struct CThostFtdcRspUserAuthMethodField *pRspUserAuthMethod, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

    typedef void (*OnRspGenUserCaptcha)(struct CThostFtdcRspGenUserCaptchaField *pRspGenUserCaptcha, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

    typedef void (*OnRspGenUserText)(struct CThostFtdcRspGenUserTextField *pRspGenUserText, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

    typedef void (*OnRspOrderInsert)(struct CThostFtdcInputOrderField *pInputOrder, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

    typedef void (*OnRspParkedOrderInsert)(struct CThostFtdcParkedOrderField *pParkedOrder, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

    typedef void (*OnRspParkedOrderAction)(struct CThostFtdcParkedOrderActionField *pParkedOrderAction, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

    typedef void (*OnRspOrderAction)(struct CThostFtdcInputOrderActionField *pInputOrderAction, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

    typedef void (*OnRspQryMaxOrderVolume)(struct CThostFtdcQryMaxOrderVolumeField *pQryMaxOrderVolume, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

    typedef void (*OnRspSettlementInfoConfirm)(struct CThostFtdcSettlementInfoConfirmField *pSettlementInfoConfirm, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

    typedef void (*OnRspRemoveParkedOrder)(struct CThostFtdcRemoveParkedOrderField *pRemoveParkedOrder, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

    typedef void (*OnRspRemoveParkedOrderAction)(struct CThostFtdcRemoveParkedOrderActionField *pRemoveParkedOrderAction, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

    typedef void (*OnRspExecOrderInsert)(struct CThostFtdcInputExecOrderField *pInputExecOrder, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

    typedef void (*OnRspExecOrderAction)(struct CThostFtdcInputExecOrderActionField *pInputExecOrderAction, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

    typedef void (*OnRspForQuoteInsert)(struct CThostFtdcInputForQuoteField *pInputForQuote, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

    typedef void (*OnRspQuoteInsert)(struct CThostFtdcInputQuoteField *pInputQuote, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

    typedef void (*OnRspQuoteAction)(struct CThostFtdcInputQuoteActionField *pInputQuoteAction, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

    typedef void (*OnRspBatchOrderAction)(struct CThostFtdcInputBatchOrderActionField *pInputBatchOrderAction, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

    typedef void (*OnRspOptionSelfCloseInsert)(struct CThostFtdcInputOptionSelfCloseField *pInputOptionSelfClose, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

    typedef void (*OnRspOptionSelfCloseAction)(struct CThostFtdcInputOptionSelfCloseActionField *pInputOptionSelfCloseAction, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

    typedef void (*OnRspCombActionInsert)(struct CThostFtdcInputCombActionField *pInputCombAction, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

    typedef void (*OnRspQryOrder)(struct CThostFtdcOrderField *pOrder, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

    typedef void (*OnRspQryTrade)(struct CThostFtdcTradeField *pTrade, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

    typedef void (*OnRspQryInvestorPosition)(struct CThostFtdcInvestorPositionField *pInvestorPosition, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

    typedef void (*OnRspQryTradingAccount)(struct CThostFtdcTradingAccountField *pTradingAccount, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

    typedef void (*OnRspQryInvestor)(struct CThostFtdcInvestorField *pInvestor, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

    typedef void (*OnRspQryTradingCode)(struct CThostFtdcTradingCodeField *pTradingCode, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

    typedef void (*OnRspQryInstrumentMarginRate)(struct CThostFtdcInstrumentMarginRateField *pInstrumentMarginRate, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

    typedef void (*OnRspQryInstrumentCommissionRate)(struct CThostFtdcInstrumentCommissionRateField *pInstrumentCommissionRate, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

    typedef void (*OnRspQryUserSession)(struct CThostFtdcUserSessionField *pUserSession, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

    typedef void (*OnRspQryExchange)(struct CThostFtdcExchangeField *pExchange, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

    typedef void (*OnRspQryProduct)(struct CThostFtdcProductField *pProduct, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

    typedef void (*OnRspQryInstrument)(struct CThostFtdcInstrumentField *pInstrument, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

    typedef void (*OnRspQryDepthMarketData)(struct CThostFtdcDepthMarketDataField *pDepthMarketData, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

    typedef void (*OnRspQryTraderOffer)(struct CThostFtdcTraderOfferField *pTraderOffer, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

    typedef void (*OnRspQrySettlementInfo)(struct CThostFtdcSettlementInfoField *pSettlementInfo, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

    typedef void (*OnRspQryTransferBank)(struct CThostFtdcTransferBankField *pTransferBank, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

    typedef void (*OnRspQryInvestorPositionDetail)(struct CThostFtdcInvestorPositionDetailField *pInvestorPositionDetail, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

    typedef void (*OnRspQryNotice)(struct CThostFtdcNoticeField *pNotice, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

    typedef void (*OnRspQrySettlementInfoConfirm)(struct CThostFtdcSettlementInfoConfirmField *pSettlementInfoConfirm, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

    typedef void (*OnRspQryInvestorPositionCombineDetail)(struct CThostFtdcInvestorPositionCombineDetailField *pInvestorPositionCombineDetail, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

    typedef void (*OnRspQryCFMMCTradingAccountKey)(struct CThostFtdcCFMMCTradingAccountKeyField *pCFMMCTradingAccountKey, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

    typedef void (*OnRspQryEWarrantOffset)(struct CThostFtdcEWarrantOffsetField *pEWarrantOffset, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

    typedef void (*OnRspQryInvestorProductGroupMargin)(struct CThostFtdcInvestorProductGroupMarginField *pInvestorProductGroupMargin, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

    typedef void (*OnRspQryExchangeMarginRate)(struct CThostFtdcExchangeMarginRateField *pExchangeMarginRate, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

    typedef void (*OnRspQryExchangeMarginRateAdjust)(struct CThostFtdcExchangeMarginRateAdjustField *pExchangeMarginRateAdjust, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

    typedef void (*OnRspQryExchangeRate)(struct CThostFtdcExchangeRateField *pExchangeRate, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

    typedef void (*OnRspQrySecAgentACIDMap)(struct CThostFtdcSecAgentACIDMapField *pSecAgentACIDMap, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

    typedef void (*OnRspQryProductExchRate)(struct CThostFtdcProductExchRateField *pProductExchRate, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

    typedef void (*OnRspQryProductGroup)(struct CThostFtdcProductGroupField *pProductGroup, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

    typedef void (*OnRspQryMMInstrumentCommissionRate)(struct CThostFtdcMMInstrumentCommissionRateField *pMMInstrumentCommissionRate, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

    typedef void (*OnRspQryMMOptionInstrCommRate)(struct CThostFtdcMMOptionInstrCommRateField *pMMOptionInstrCommRate, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

    typedef void (*OnRspQryInstrumentOrderCommRate)(struct CThostFtdcInstrumentOrderCommRateField *pInstrumentOrderCommRate, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

    typedef void (*OnRspQrySecAgentTradingAccount)(struct CThostFtdcTradingAccountField *pTradingAccount, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

    typedef void (*OnRspQrySecAgentCheckMode)(struct CThostFtdcSecAgentCheckModeField *pSecAgentCheckMode, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

    typedef void (*OnRspQrySecAgentTradeInfo)(struct CThostFtdcSecAgentTradeInfoField *pSecAgentTradeInfo, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

    typedef void (*OnRspQryOptionInstrTradeCost)(struct CThostFtdcOptionInstrTradeCostField *pOptionInstrTradeCost, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

    typedef void (*OnRspQryOptionInstrCommRate)(struct CThostFtdcOptionInstrCommRateField *pOptionInstrCommRate, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

    typedef void (*OnRspQryExecOrder)(struct CThostFtdcExecOrderField *pExecOrder, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

    typedef void (*OnRspQryForQuote)(struct CThostFtdcForQuoteField *pForQuote, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

    typedef void (*OnRspQryQuote)(struct CThostFtdcQuoteField *pQuote, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

    typedef void (*OnRspQryOptionSelfClose)(struct CThostFtdcOptionSelfCloseField *pOptionSelfClose, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

    typedef void (*OnRspQryInvestUnit)(struct CThostFtdcInvestUnitField *pInvestUnit, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

    typedef void (*OnRspQryCombInstrumentGuard)(struct CThostFtdcCombInstrumentGuardField *pCombInstrumentGuard, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

    typedef void (*OnRspQryCombAction)(struct CThostFtdcCombActionField *pCombAction, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

    typedef void (*OnRspQryTransferSerial)(struct CThostFtdcTransferSerialField *pTransferSerial, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

    typedef void (*OnRspQryAccountregister)(struct CThostFtdcAccountregisterField *pAccountregister, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

    typedef void (*OnRspError)(struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

    typedef void (*OnRtnOrder)(struct CThostFtdcOrderField *pOrder);

    typedef void (*OnRtnTrade)(struct CThostFtdcTradeField *pTrade);

    typedef void (*OnErrRtnOrderInsert)(struct CThostFtdcInputOrderField *pInputOrder, struct CThostFtdcRspInfoField *pRspInfo);

    typedef void (*OnErrRtnOrderAction)(struct CThostFtdcOrderActionField *pOrderAction, struct CThostFtdcRspInfoField *pRspInfo);

    typedef void (*OnRtnInstrumentStatus)(struct CThostFtdcInstrumentStatusField *pInstrumentStatus);

    typedef void (*OnRtnBulletin)(struct CThostFtdcBulletinField *pBulletin);

    typedef void (*OnRtnTradingNotice)(struct CThostFtdcTradingNoticeInfoField *pTradingNoticeInfo);

    typedef void (*OnRtnErrorConditionalOrder)(struct CThostFtdcErrorConditionalOrderField *pErrorConditionalOrder);

    typedef void (*OnRtnExecOrder)(struct CThostFtdcExecOrderField *pExecOrder);

    typedef void (*OnErrRtnExecOrderInsert)(struct CThostFtdcInputExecOrderField *pInputExecOrder, struct CThostFtdcRspInfoField *pRspInfo);

    typedef void (*OnErrRtnExecOrderAction)(struct CThostFtdcExecOrderActionField *pExecOrderAction, struct CThostFtdcRspInfoField *pRspInfo);

    typedef void (*OnErrRtnForQuoteInsert)(struct CThostFtdcInputForQuoteField *pInputForQuote, struct CThostFtdcRspInfoField *pRspInfo);

    typedef void (*OnRtnQuote)(struct CThostFtdcQuoteField *pQuote);

    typedef void (*OnErrRtnQuoteInsert)(struct CThostFtdcInputQuoteField *pInputQuote, struct CThostFtdcRspInfoField *pRspInfo);

    typedef void (*OnErrRtnQuoteAction)(struct CThostFtdcQuoteActionField *pQuoteAction, struct CThostFtdcRspInfoField *pRspInfo);

    typedef void (*OnRtnForQuoteRsp)(struct CThostFtdcForQuoteRspField *pForQuoteRsp);

    typedef void (*OnRtnCFMMCTradingAccountToken)(struct CThostFtdcCFMMCTradingAccountTokenField *pCFMMCTradingAccountToken);

    typedef void (*OnErrRtnBatchOrderAction)(struct CThostFtdcBatchOrderActionField *pBatchOrderAction, struct CThostFtdcRspInfoField *pRspInfo);

    typedef void (*OnRtnOptionSelfClose)(struct CThostFtdcOptionSelfCloseField *pOptionSelfClose);

    typedef void (*OnErrRtnOptionSelfCloseInsert)(struct CThostFtdcInputOptionSelfCloseField *pInputOptionSelfClose, struct CThostFtdcRspInfoField *pRspInfo);

    typedef void (*OnErrRtnOptionSelfCloseAction)(struct CThostFtdcOptionSelfCloseActionField *pOptionSelfCloseAction, struct CThostFtdcRspInfoField *pRspInfo);

    typedef void (*OnRtnCombAction)(struct CThostFtdcCombActionField *pCombAction);

    typedef void (*OnErrRtnCombActionInsert)(struct CThostFtdcInputCombActionField *pInputCombAction, struct CThostFtdcRspInfoField *pRspInfo);

    typedef void (*OnRspQryContractBank)(struct CThostFtdcContractBankField *pContractBank, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

    typedef void (*OnRspQryParkedOrder)(struct CThostFtdcParkedOrderField *pParkedOrder, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

    typedef void (*OnRspQryParkedOrderAction)(struct CThostFtdcParkedOrderActionField *pParkedOrderAction, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

    typedef void (*OnRspQryTradingNotice)(struct CThostFtdcTradingNoticeField *pTradingNotice, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

    typedef void (*OnRspQryBrokerTradingParams)(struct CThostFtdcBrokerTradingParamsField *pBrokerTradingParams, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

    typedef void (*OnRspQryBrokerTradingAlgos)(struct CThostFtdcBrokerTradingAlgosField *pBrokerTradingAlgos, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

    typedef void (*OnRspQueryCFMMCTradingAccountToken)(struct CThostFtdcQueryCFMMCTradingAccountTokenField *pQueryCFMMCTradingAccountToken, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

    typedef void (*OnRtnFromBankToFutureByBank)(struct CThostFtdcRspTransferField *pRspTransfer);

    typedef void (*OnRtnFromFutureToBankByBank)(struct CThostFtdcRspTransferField *pRspTransfer);

    typedef void (*OnRtnRepealFromBankToFutureByBank)(struct CThostFtdcRspRepealField *pRspRepeal);

    typedef void (*OnRtnRepealFromFutureToBankByBank)(struct CThostFtdcRspRepealField *pRspRepeal);

    typedef void (*OnRtnFromBankToFutureByFuture)(struct CThostFtdcRspTransferField *pRspTransfer);

    typedef void (*OnRtnFromFutureToBankByFuture)(struct CThostFtdcRspTransferField *pRspTransfer);

    typedef void (*OnRtnRepealFromBankToFutureByFutureManual)(struct CThostFtdcRspRepealField *pRspRepeal);

    typedef void (*OnRtnRepealFromFutureToBankByFutureManual)(struct CThostFtdcRspRepealField *pRspRepeal);

    typedef void (*OnRtnQueryBankBalanceByFuture)(struct CThostFtdcNotifyQueryAccountField *pNotifyQueryAccount);

    typedef void (*OnErrRtnBankToFutureByFuture)(struct CThostFtdcReqTransferField *pReqTransfer, struct CThostFtdcRspInfoField *pRspInfo);

    typedef void (*OnErrRtnFutureToBankByFuture)(struct CThostFtdcReqTransferField *pReqTransfer, struct CThostFtdcRspInfoField *pRspInfo);

    typedef void (*OnErrRtnRepealBankToFutureByFutureManual)(struct CThostFtdcReqRepealField *pReqRepeal, struct CThostFtdcRspInfoField *pRspInfo);

    typedef void (*OnErrRtnRepealFutureToBankByFutureManual)(struct CThostFtdcReqRepealField *pReqRepeal, struct CThostFtdcRspInfoField *pRspInfo);

    typedef void (*OnErrRtnQueryBankBalanceByFuture)(struct CThostFtdcReqQueryAccountField *pReqQueryAccount, struct CThostFtdcRspInfoField *pRspInfo);

    typedef void (*OnRtnRepealFromBankToFutureByFuture)(struct CThostFtdcRspRepealField *pRspRepeal);

    typedef void (*OnRtnRepealFromFutureToBankByFuture)(struct CThostFtdcRspRepealField *pRspRepeal);

    typedef void (*OnRspFromBankToFutureByFuture)(struct CThostFtdcReqTransferField *pReqTransfer, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

    typedef void (*OnRspFromFutureToBankByFuture)(struct CThostFtdcReqTransferField *pReqTransfer, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

    typedef void (*OnRspQueryBankAccountMoneyByFuture)(struct CThostFtdcReqQueryAccountField *pReqQueryAccount, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

    typedef void (*OnRtnOpenAccountByBank)(struct CThostFtdcOpenAccountField *pOpenAccount);

    typedef void (*OnRtnCancelAccountByBank)(struct CThostFtdcCancelAccountField *pCancelAccount);

    typedef void (*OnRtnChangeAccountByBank)(struct CThostFtdcChangeAccountField *pChangeAccount);

    typedef void (*OnRspQryClassifiedInstrument)(struct CThostFtdcInstrumentField *pInstrument, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

    typedef void (*OnRspQryCombPromotionParam)(struct CThostFtdcCombPromotionParamField *pCombPromotionParam, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

    typedef void (*OnRspQryRiskSettleInvstPosition)(struct CThostFtdcRiskSettleInvstPositionField *pRiskSettleInvstPosition, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

    typedef void (*OnRspQryRiskSettleProductStatus)(struct CThostFtdcRiskSettleProductStatusField *pRiskSettleProductStatus, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

    typedef void (*OnRspQrySPBMFutureParameter)(struct CThostFtdcSPBMFutureParameterField *pSPBMFutureParameter, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

    typedef void (*OnRspQrySPBMOptionParameter)(struct CThostFtdcSPBMOptionParameterField *pSPBMOptionParameter, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

    typedef void (*OnRspQrySPBMIntraParameter)(struct CThostFtdcSPBMIntraParameterField *pSPBMIntraParameter, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

    typedef void (*OnRspQrySPBMInterParameter)(struct CThostFtdcSPBMInterParameterField *pSPBMInterParameter, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

    typedef void (*OnRspQrySPBMPortfDefinition)(struct CThostFtdcSPBMPortfDefinitionField *pSPBMPortfDefinition, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

    typedef void (*OnRspQrySPBMInvestorPortfDef)(struct CThostFtdcSPBMInvestorPortfDefField *pSPBMInvestorPortfDef, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

    typedef void (*OnRspQryInvestorPortfMarginRatio)(struct CThostFtdcInvestorPortfMarginRatioField *pInvestorPortfMarginRatio, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

    typedef void (*OnRspQryInvestorProdSPBMDetail)(struct CThostFtdcInvestorProdSPBMDetailField *pInvestorProdSPBMDetail, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

    typedef void (*OnRspQryInvestorCommoditySPMMMargin)(struct CThostFtdcInvestorCommoditySPMMMarginField *pInvestorCommoditySPMMMargin, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

    typedef void (*OnRspQryInvestorCommodityGroupSPMMMargin)(struct CThostFtdcInvestorCommodityGroupSPMMMarginField *pInvestorCommodityGroupSPMMMargin, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

    typedef void (*OnRspQrySPMMInstParam)(struct CThostFtdcSPMMInstParamField *pSPMMInstParam, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

    typedef void (*OnRspQrySPMMProductParam)(struct CThostFtdcSPMMProductParamField *pSPMMProductParam, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

    typedef void (*OnRspQrySPBMAddOnInterParameter)(struct CThostFtdcSPBMAddOnInterParameterField *pSPBMAddOnInterParameter, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

    typedef void (*OnRspQryRCAMSCombProductInfo)(struct CThostFtdcRCAMSCombProductInfoField *pRCAMSCombProductInfo, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

    typedef void (*OnRspQryRCAMSInstrParameter)(struct CThostFtdcRCAMSInstrParameterField *pRCAMSInstrParameter, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

    typedef void (*OnRspQryRCAMSIntraParameter)(struct CThostFtdcRCAMSIntraParameterField *pRCAMSIntraParameter, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

    typedef void (*OnRspQryRCAMSInterParameter)(struct CThostFtdcRCAMSInterParameterField *pRCAMSInterParameter, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

    typedef void (*OnRspQryRCAMSShortOptAdjustParam)(struct CThostFtdcRCAMSShortOptAdjustParamField *pRCAMSShortOptAdjustParam, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

    typedef void (*OnRspQryRCAMSInvestorCombPosition)(struct CThostFtdcRCAMSInvestorCombPositionField *pRCAMSInvestorCombPosition, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

    typedef void (*OnRspQryInvestorProdRCAMSMargin)(struct CThostFtdcInvestorProdRCAMSMarginField *pInvestorProdRCAMSMargin, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

    typedef void (*OnRspQryRULEInstrParameter)(struct CThostFtdcRULEInstrParameterField *pRULEInstrParameter, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

    typedef void (*OnRspQryRULEIntraParameter)(struct CThostFtdcRULEIntraParameterField *pRULEIntraParameter, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

    typedef void (*OnRspQryRULEInterParameter)(struct CThostFtdcRULEInterParameterField *pRULEInterParameter, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

    typedef void (*OnRspQryInvestorProdRULEMargin)(struct CThostFtdcInvestorProdRULEMarginField *pInvestorProdRULEMargin, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

    typedef void (*OnRspQryInvestorPortfSetting)(struct CThostFtdcInvestorPortfSettingField *pInvestorPortfSetting, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

    typedef void (*OnRspQryInvestorInfoCommRec)(struct CThostFtdcInvestorInfoCommRecField *pInvestorInfoCommRec, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

    typedef void (*OnRspQryCombLeg)(struct CThostFtdcCombLegField *pCombLeg, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

    typedef void (*OnRspOffsetSetting)(struct CThostFtdcInputOffsetSettingField *pInputOffsetSetting, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

    typedef void (*OnRspCancelOffsetSetting)(struct CThostFtdcInputOffsetSettingField *pInputOffsetSetting, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

    typedef void (*OnRtnOffsetSetting)(struct CThostFtdcOffsetSettingField *pOffsetSetting);

    typedef void (*OnErrRtnOffsetSetting)(struct CThostFtdcInputOffsetSettingField *pInputOffsetSetting, struct CThostFtdcRspInfoField *pRspInfo);

    typedef void (*OnErrRtnCancelOffsetSetting)(struct CThostFtdcCancelOffsetSettingField *pCancelOffsetSetting, struct CThostFtdcRspInfoField *pRspInfo);

    typedef void (*OnRspQryOffsetSetting)(struct CThostFtdcOffsetSettingField *pOffsetSetting, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

    typedef void (*OnRspGenSMSCode)(struct CThostFtdcRspGenSMSCodeField *pRspGenSMSCode, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

    typedef void (*OnRspSpdApply)(struct CThostFtdcInputSpdApplyField *pInputSpdApply, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

    typedef void (*OnRspSpdApplyAction)(struct CThostFtdcInputSpdApplyActionField *pInputSpdApplyAction, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

    typedef void (*OnRspQrySpdApply)(struct CThostFtdcSpdApplyField *pSpdApply, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

    typedef void (*OnRtnSpdApply)(struct CThostFtdcSpdApplyField *pSpdApply);

    typedef void (*OnErrRtnSpdApply)(struct CThostFtdcInputSpdApplyField *pInputSpdApply, struct CThostFtdcRspInfoField *pRspInfo);

    typedef void (*OnErrRtnSpdApplyAction)(struct CThostFtdcSpdApplyActionField *pSpdApplyAction, struct CThostFtdcRspInfoField *pRspInfo);

    typedef void (*OnRspHedgeCfm)(struct CThostFtdcInputHedgeCfmField *pInputHedgeCfm, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

    typedef void (*OnRspHedgeCfmAction)(struct CThostFtdcInputHedgeCfmActionField *pInputHedgeCfmAction, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

    typedef void (*OnRspQryHedgeCfm)(struct CThostFtdcHedgeCfmField *pHedgeCfm, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

    typedef void (*OnRtnHedgeCfm)(struct CThostFtdcHedgeCfmField *pHedgeCfm);

    typedef void (*OnErrRtnHedgeCfm)(struct CThostFtdcInputHedgeCfmField *pInputHedgeCfm, struct CThostFtdcRspInfoField *pRspInfo);

    typedef void (*OnErrRtnHedgeCfmAction)(struct CThostFtdcHedgeCfmActionField *pHedgeCfmAction, struct CThostFtdcRspInfoField *pRspInfo);

    void COnFrontConnected(void *this);

    void COnFrontDisconnected(void *this, int nReason);

    void COnHeartBeatWarning(void *this, int nTimeLapse);

    void COnRspAuthenticate(void *this, struct CThostFtdcRspAuthenticateField *pRspAuthenticateField, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

    void COnRtnPrivateSeqNo(void *this, int nSeqNo);

    void COnRspUserLogin(void *this, struct CThostFtdcRspUserLoginField *pRspUserLogin, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

    void COnRspUserLogout(void *this, struct CThostFtdcUserLogoutField *pUserLogout, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

    void COnRspUserPasswordUpdate(void *this, struct CThostFtdcUserPasswordUpdateField *pUserPasswordUpdate, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

    void COnRspTradingAccountPasswordUpdate(void *this, struct CThostFtdcTradingAccountPasswordUpdateField *pTradingAccountPasswordUpdate, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

    void COnRspUserAuthMethod(void *this, struct CThostFtdcRspUserAuthMethodField *pRspUserAuthMethod, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

    void COnRspGenUserCaptcha(void *this, struct CThostFtdcRspGenUserCaptchaField *pRspGenUserCaptcha, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

    void COnRspGenUserText(void *this, struct CThostFtdcRspGenUserTextField *pRspGenUserText, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

    void COnRspOrderInsert(void *this, struct CThostFtdcInputOrderField *pInputOrder, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

    void COnRspParkedOrderInsert(void *this, struct CThostFtdcParkedOrderField *pParkedOrder, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

    void COnRspParkedOrderAction(void *this, struct CThostFtdcParkedOrderActionField *pParkedOrderAction, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

    void COnRspOrderAction(void *this, struct CThostFtdcInputOrderActionField *pInputOrderAction, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

    void COnRspQryMaxOrderVolume(void *this, struct CThostFtdcQryMaxOrderVolumeField *pQryMaxOrderVolume, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

    void COnRspSettlementInfoConfirm(void *this, struct CThostFtdcSettlementInfoConfirmField *pSettlementInfoConfirm, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

    void COnRspRemoveParkedOrder(void *this, struct CThostFtdcRemoveParkedOrderField *pRemoveParkedOrder, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

    void COnRspRemoveParkedOrderAction(void *this, struct CThostFtdcRemoveParkedOrderActionField *pRemoveParkedOrderAction, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

    void COnRspExecOrderInsert(void *this, struct CThostFtdcInputExecOrderField *pInputExecOrder, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

    void COnRspExecOrderAction(void *this, struct CThostFtdcInputExecOrderActionField *pInputExecOrderAction, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

    void COnRspForQuoteInsert(void *this, struct CThostFtdcInputForQuoteField *pInputForQuote, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

    void COnRspQuoteInsert(void *this, struct CThostFtdcInputQuoteField *pInputQuote, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

    void COnRspQuoteAction(void *this, struct CThostFtdcInputQuoteActionField *pInputQuoteAction, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

    void COnRspBatchOrderAction(void *this, struct CThostFtdcInputBatchOrderActionField *pInputBatchOrderAction, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

    void COnRspOptionSelfCloseInsert(void *this, struct CThostFtdcInputOptionSelfCloseField *pInputOptionSelfClose, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

    void COnRspOptionSelfCloseAction(void *this, struct CThostFtdcInputOptionSelfCloseActionField *pInputOptionSelfCloseAction, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

    void COnRspCombActionInsert(void *this, struct CThostFtdcInputCombActionField *pInputCombAction, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

    void COnRspQryOrder(void *this, struct CThostFtdcOrderField *pOrder, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

    void COnRspQryTrade(void *this, struct CThostFtdcTradeField *pTrade, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

    void COnRspQryInvestorPosition(void *this, struct CThostFtdcInvestorPositionField *pInvestorPosition, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

    void COnRspQryTradingAccount(void *this, struct CThostFtdcTradingAccountField *pTradingAccount, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

    void COnRspQryInvestor(void *this, struct CThostFtdcInvestorField *pInvestor, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

    void COnRspQryTradingCode(void *this, struct CThostFtdcTradingCodeField *pTradingCode, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

    void COnRspQryInstrumentMarginRate(void *this, struct CThostFtdcInstrumentMarginRateField *pInstrumentMarginRate, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

    void COnRspQryInstrumentCommissionRate(void *this, struct CThostFtdcInstrumentCommissionRateField *pInstrumentCommissionRate, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

    void COnRspQryUserSession(void *this, struct CThostFtdcUserSessionField *pUserSession, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

    void COnRspQryExchange(void *this, struct CThostFtdcExchangeField *pExchange, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

    void COnRspQryProduct(void *this, struct CThostFtdcProductField *pProduct, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

    void COnRspQryInstrument(void *this, struct CThostFtdcInstrumentField *pInstrument, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

    void COnRspQryDepthMarketData(void *this, struct CThostFtdcDepthMarketDataField *pDepthMarketData, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

    void COnRspQryTraderOffer(void *this, struct CThostFtdcTraderOfferField *pTraderOffer, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

    void COnRspQrySettlementInfo(void *this, struct CThostFtdcSettlementInfoField *pSettlementInfo, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

    void COnRspQryTransferBank(void *this, struct CThostFtdcTransferBankField *pTransferBank, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

    void COnRspQryInvestorPositionDetail(void *this, struct CThostFtdcInvestorPositionDetailField *pInvestorPositionDetail, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

    void COnRspQryNotice(void *this, struct CThostFtdcNoticeField *pNotice, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

    void COnRspQrySettlementInfoConfirm(void *this, struct CThostFtdcSettlementInfoConfirmField *pSettlementInfoConfirm, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

    void COnRspQryInvestorPositionCombineDetail(void *this, struct CThostFtdcInvestorPositionCombineDetailField *pInvestorPositionCombineDetail, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

    void COnRspQryCFMMCTradingAccountKey(void *this, struct CThostFtdcCFMMCTradingAccountKeyField *pCFMMCTradingAccountKey, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

    void COnRspQryEWarrantOffset(void *this, struct CThostFtdcEWarrantOffsetField *pEWarrantOffset, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

    void COnRspQryInvestorProductGroupMargin(void *this, struct CThostFtdcInvestorProductGroupMarginField *pInvestorProductGroupMargin, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

    void COnRspQryExchangeMarginRate(void *this, struct CThostFtdcExchangeMarginRateField *pExchangeMarginRate, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

    void COnRspQryExchangeMarginRateAdjust(void *this, struct CThostFtdcExchangeMarginRateAdjustField *pExchangeMarginRateAdjust, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

    void COnRspQryExchangeRate(void *this, struct CThostFtdcExchangeRateField *pExchangeRate, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

    void COnRspQrySecAgentACIDMap(void *this, struct CThostFtdcSecAgentACIDMapField *pSecAgentACIDMap, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

    void COnRspQryProductExchRate(void *this, struct CThostFtdcProductExchRateField *pProductExchRate, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

    void COnRspQryProductGroup(void *this, struct CThostFtdcProductGroupField *pProductGroup, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

    void COnRspQryMMInstrumentCommissionRate(void *this, struct CThostFtdcMMInstrumentCommissionRateField *pMMInstrumentCommissionRate, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

    void COnRspQryMMOptionInstrCommRate(void *this, struct CThostFtdcMMOptionInstrCommRateField *pMMOptionInstrCommRate, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

    void COnRspQryInstrumentOrderCommRate(void *this, struct CThostFtdcInstrumentOrderCommRateField *pInstrumentOrderCommRate, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

    void COnRspQrySecAgentTradingAccount(void *this, struct CThostFtdcTradingAccountField *pTradingAccount, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

    void COnRspQrySecAgentCheckMode(void *this, struct CThostFtdcSecAgentCheckModeField *pSecAgentCheckMode, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

    void COnRspQrySecAgentTradeInfo(void *this, struct CThostFtdcSecAgentTradeInfoField *pSecAgentTradeInfo, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

    void COnRspQryOptionInstrTradeCost(void *this, struct CThostFtdcOptionInstrTradeCostField *pOptionInstrTradeCost, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

    void COnRspQryOptionInstrCommRate(void *this, struct CThostFtdcOptionInstrCommRateField *pOptionInstrCommRate, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

    void COnRspQryExecOrder(void *this, struct CThostFtdcExecOrderField *pExecOrder, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

    void COnRspQryForQuote(void *this, struct CThostFtdcForQuoteField *pForQuote, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

    void COnRspQryQuote(void *this, struct CThostFtdcQuoteField *pQuote, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

    void COnRspQryOptionSelfClose(void *this, struct CThostFtdcOptionSelfCloseField *pOptionSelfClose, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

    void COnRspQryInvestUnit(void *this, struct CThostFtdcInvestUnitField *pInvestUnit, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

    void COnRspQryCombInstrumentGuard(void *this, struct CThostFtdcCombInstrumentGuardField *pCombInstrumentGuard, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

    void COnRspQryCombAction(void *this, struct CThostFtdcCombActionField *pCombAction, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

    void COnRspQryTransferSerial(void *this, struct CThostFtdcTransferSerialField *pTransferSerial, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

    void COnRspQryAccountregister(void *this, struct CThostFtdcAccountregisterField *pAccountregister, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

    void COnRspError(void *this, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

    void COnRtnOrder(void *this, struct CThostFtdcOrderField *pOrder);

    void COnRtnTrade(void *this, struct CThostFtdcTradeField *pTrade);

    void COnErrRtnOrderInsert(void *this, struct CThostFtdcInputOrderField *pInputOrder, struct CThostFtdcRspInfoField *pRspInfo);

    void COnErrRtnOrderAction(void *this, struct CThostFtdcOrderActionField *pOrderAction, struct CThostFtdcRspInfoField *pRspInfo);

    void COnRtnInstrumentStatus(void *this, struct CThostFtdcInstrumentStatusField *pInstrumentStatus);

    void COnRtnBulletin(void *this, struct CThostFtdcBulletinField *pBulletin);

    void COnRtnTradingNotice(void *this, struct CThostFtdcTradingNoticeInfoField *pTradingNoticeInfo);

    void COnRtnErrorConditionalOrder(void *this, struct CThostFtdcErrorConditionalOrderField *pErrorConditionalOrder);

    void COnRtnExecOrder(void *this, struct CThostFtdcExecOrderField *pExecOrder);

    void COnErrRtnExecOrderInsert(void *this, struct CThostFtdcInputExecOrderField *pInputExecOrder, struct CThostFtdcRspInfoField *pRspInfo);

    void COnErrRtnExecOrderAction(void *this, struct CThostFtdcExecOrderActionField *pExecOrderAction, struct CThostFtdcRspInfoField *pRspInfo);

    void COnErrRtnForQuoteInsert(void *this, struct CThostFtdcInputForQuoteField *pInputForQuote, struct CThostFtdcRspInfoField *pRspInfo);

    void COnRtnQuote(void *this, struct CThostFtdcQuoteField *pQuote);

    void COnErrRtnQuoteInsert(void *this, struct CThostFtdcInputQuoteField *pInputQuote, struct CThostFtdcRspInfoField *pRspInfo);

    void COnErrRtnQuoteAction(void *this, struct CThostFtdcQuoteActionField *pQuoteAction, struct CThostFtdcRspInfoField *pRspInfo);

    void COnRtnForQuoteRsp(void *this, struct CThostFtdcForQuoteRspField *pForQuoteRsp);

    void COnRtnCFMMCTradingAccountToken(void *this, struct CThostFtdcCFMMCTradingAccountTokenField *pCFMMCTradingAccountToken);

    void COnErrRtnBatchOrderAction(void *this, struct CThostFtdcBatchOrderActionField *pBatchOrderAction, struct CThostFtdcRspInfoField *pRspInfo);

    void COnRtnOptionSelfClose(void *this, struct CThostFtdcOptionSelfCloseField *pOptionSelfClose);

    void COnErrRtnOptionSelfCloseInsert(void *this, struct CThostFtdcInputOptionSelfCloseField *pInputOptionSelfClose, struct CThostFtdcRspInfoField *pRspInfo);

    void COnErrRtnOptionSelfCloseAction(void *this, struct CThostFtdcOptionSelfCloseActionField *pOptionSelfCloseAction, struct CThostFtdcRspInfoField *pRspInfo);

    void COnRtnCombAction(void *this, struct CThostFtdcCombActionField *pCombAction);

    void COnErrRtnCombActionInsert(void *this, struct CThostFtdcInputCombActionField *pInputCombAction, struct CThostFtdcRspInfoField *pRspInfo);

    void COnRspQryContractBank(void *this, struct CThostFtdcContractBankField *pContractBank, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

    void COnRspQryParkedOrder(void *this, struct CThostFtdcParkedOrderField *pParkedOrder, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

    void COnRspQryParkedOrderAction(void *this, struct CThostFtdcParkedOrderActionField *pParkedOrderAction, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

    void COnRspQryTradingNotice(void *this, struct CThostFtdcTradingNoticeField *pTradingNotice, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

    void COnRspQryBrokerTradingParams(void *this, struct CThostFtdcBrokerTradingParamsField *pBrokerTradingParams, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

    void COnRspQryBrokerTradingAlgos(void *this, struct CThostFtdcBrokerTradingAlgosField *pBrokerTradingAlgos, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

    void COnRspQueryCFMMCTradingAccountToken(void *this, struct CThostFtdcQueryCFMMCTradingAccountTokenField *pQueryCFMMCTradingAccountToken, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

    void COnRtnFromBankToFutureByBank(void *this, struct CThostFtdcRspTransferField *pRspTransfer);

    void COnRtnFromFutureToBankByBank(void *this, struct CThostFtdcRspTransferField *pRspTransfer);

    void COnRtnRepealFromBankToFutureByBank(void *this, struct CThostFtdcRspRepealField *pRspRepeal);

    void COnRtnRepealFromFutureToBankByBank(void *this, struct CThostFtdcRspRepealField *pRspRepeal);

    void COnRtnFromBankToFutureByFuture(void *this, struct CThostFtdcRspTransferField *pRspTransfer);

    void COnRtnFromFutureToBankByFuture(void *this, struct CThostFtdcRspTransferField *pRspTransfer);

    void COnRtnRepealFromBankToFutureByFutureManual(void *this, struct CThostFtdcRspRepealField *pRspRepeal);

    void COnRtnRepealFromFutureToBankByFutureManual(void *this, struct CThostFtdcRspRepealField *pRspRepeal);

    void COnRtnQueryBankBalanceByFuture(void *this, struct CThostFtdcNotifyQueryAccountField *pNotifyQueryAccount);

    void COnErrRtnBankToFutureByFuture(void *this, struct CThostFtdcReqTransferField *pReqTransfer, struct CThostFtdcRspInfoField *pRspInfo);

    void COnErrRtnFutureToBankByFuture(void *this, struct CThostFtdcReqTransferField *pReqTransfer, struct CThostFtdcRspInfoField *pRspInfo);

    void COnErrRtnRepealBankToFutureByFutureManual(void *this, struct CThostFtdcReqRepealField *pReqRepeal, struct CThostFtdcRspInfoField *pRspInfo);

    void COnErrRtnRepealFutureToBankByFutureManual(void *this, struct CThostFtdcReqRepealField *pReqRepeal, struct CThostFtdcRspInfoField *pRspInfo);

    void COnErrRtnQueryBankBalanceByFuture(void *this, struct CThostFtdcReqQueryAccountField *pReqQueryAccount, struct CThostFtdcRspInfoField *pRspInfo);

    void COnRtnRepealFromBankToFutureByFuture(void *this, struct CThostFtdcRspRepealField *pRspRepeal);

    void COnRtnRepealFromFutureToBankByFuture(void *this, struct CThostFtdcRspRepealField *pRspRepeal);

    void COnRspFromBankToFutureByFuture(void *this, struct CThostFtdcReqTransferField *pReqTransfer, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

    void COnRspFromFutureToBankByFuture(void *this, struct CThostFtdcReqTransferField *pReqTransfer, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

    void COnRspQueryBankAccountMoneyByFuture(void *this, struct CThostFtdcReqQueryAccountField *pReqQueryAccount, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

    void COnRtnOpenAccountByBank(void *this, struct CThostFtdcOpenAccountField *pOpenAccount);

    void COnRtnCancelAccountByBank(void *this, struct CThostFtdcCancelAccountField *pCancelAccount);

    void COnRtnChangeAccountByBank(void *this, struct CThostFtdcChangeAccountField *pChangeAccount);

    void COnRspQryClassifiedInstrument(void *this, struct CThostFtdcInstrumentField *pInstrument, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

    void COnRspQryCombPromotionParam(void *this, struct CThostFtdcCombPromotionParamField *pCombPromotionParam, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

    void COnRspQryRiskSettleInvstPosition(void *this, struct CThostFtdcRiskSettleInvstPositionField *pRiskSettleInvstPosition, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

    void COnRspQryRiskSettleProductStatus(void *this, struct CThostFtdcRiskSettleProductStatusField *pRiskSettleProductStatus, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

    void COnRspQrySPBMFutureParameter(void *this, struct CThostFtdcSPBMFutureParameterField *pSPBMFutureParameter, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

    void COnRspQrySPBMOptionParameter(void *this, struct CThostFtdcSPBMOptionParameterField *pSPBMOptionParameter, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

    void COnRspQrySPBMIntraParameter(void *this, struct CThostFtdcSPBMIntraParameterField *pSPBMIntraParameter, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

    void COnRspQrySPBMInterParameter(void *this, struct CThostFtdcSPBMInterParameterField *pSPBMInterParameter, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

    void COnRspQrySPBMPortfDefinition(void *this, struct CThostFtdcSPBMPortfDefinitionField *pSPBMPortfDefinition, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

    void COnRspQrySPBMInvestorPortfDef(void *this, struct CThostFtdcSPBMInvestorPortfDefField *pSPBMInvestorPortfDef, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

    void COnRspQryInvestorPortfMarginRatio(void *this, struct CThostFtdcInvestorPortfMarginRatioField *pInvestorPortfMarginRatio, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

    void COnRspQryInvestorProdSPBMDetail(void *this, struct CThostFtdcInvestorProdSPBMDetailField *pInvestorProdSPBMDetail, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

    void COnRspQryInvestorCommoditySPMMMargin(void *this, struct CThostFtdcInvestorCommoditySPMMMarginField *pInvestorCommoditySPMMMargin, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

    void COnRspQryInvestorCommodityGroupSPMMMargin(void *this, struct CThostFtdcInvestorCommodityGroupSPMMMarginField *pInvestorCommodityGroupSPMMMargin, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

    void COnRspQrySPMMInstParam(void *this, struct CThostFtdcSPMMInstParamField *pSPMMInstParam, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

    void COnRspQrySPMMProductParam(void *this, struct CThostFtdcSPMMProductParamField *pSPMMProductParam, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

    void COnRspQrySPBMAddOnInterParameter(void *this, struct CThostFtdcSPBMAddOnInterParameterField *pSPBMAddOnInterParameter, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

    void COnRspQryRCAMSCombProductInfo(void *this, struct CThostFtdcRCAMSCombProductInfoField *pRCAMSCombProductInfo, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

    void COnRspQryRCAMSInstrParameter(void *this, struct CThostFtdcRCAMSInstrParameterField *pRCAMSInstrParameter, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

    void COnRspQryRCAMSIntraParameter(void *this, struct CThostFtdcRCAMSIntraParameterField *pRCAMSIntraParameter, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

    void COnRspQryRCAMSInterParameter(void *this, struct CThostFtdcRCAMSInterParameterField *pRCAMSInterParameter, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

    void COnRspQryRCAMSShortOptAdjustParam(void *this, struct CThostFtdcRCAMSShortOptAdjustParamField *pRCAMSShortOptAdjustParam, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

    void COnRspQryRCAMSInvestorCombPosition(void *this, struct CThostFtdcRCAMSInvestorCombPositionField *pRCAMSInvestorCombPosition, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

    void COnRspQryInvestorProdRCAMSMargin(void *this, struct CThostFtdcInvestorProdRCAMSMarginField *pInvestorProdRCAMSMargin, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

    void COnRspQryRULEInstrParameter(void *this, struct CThostFtdcRULEInstrParameterField *pRULEInstrParameter, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

    void COnRspQryRULEIntraParameter(void *this, struct CThostFtdcRULEIntraParameterField *pRULEIntraParameter, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

    void COnRspQryRULEInterParameter(void *this, struct CThostFtdcRULEInterParameterField *pRULEInterParameter, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

    void COnRspQryInvestorProdRULEMargin(void *this, struct CThostFtdcInvestorProdRULEMarginField *pInvestorProdRULEMargin, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

    void COnRspQryInvestorPortfSetting(void *this, struct CThostFtdcInvestorPortfSettingField *pInvestorPortfSetting, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

    void COnRspQryInvestorInfoCommRec(void *this, struct CThostFtdcInvestorInfoCommRecField *pInvestorInfoCommRec, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

    void COnRspQryCombLeg(void *this, struct CThostFtdcCombLegField *pCombLeg, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

    void COnRspOffsetSetting(void *this, struct CThostFtdcInputOffsetSettingField *pInputOffsetSetting, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

    void COnRspCancelOffsetSetting(void *this, struct CThostFtdcInputOffsetSettingField *pInputOffsetSetting, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

    void COnRtnOffsetSetting(void *this, struct CThostFtdcOffsetSettingField *pOffsetSetting);

    void COnErrRtnOffsetSetting(void *this, struct CThostFtdcInputOffsetSettingField *pInputOffsetSetting, struct CThostFtdcRspInfoField *pRspInfo);

    void COnErrRtnCancelOffsetSetting(void *this, struct CThostFtdcCancelOffsetSettingField *pCancelOffsetSetting, struct CThostFtdcRspInfoField *pRspInfo);

    void COnRspQryOffsetSetting(void *this, struct CThostFtdcOffsetSettingField *pOffsetSetting, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

    void COnRspGenSMSCode(void *this, struct CThostFtdcRspGenSMSCodeField *pRspGenSMSCode, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

    void COnRspSpdApply(void *this, struct CThostFtdcInputSpdApplyField *pInputSpdApply, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

    void COnRspSpdApplyAction(void *this, struct CThostFtdcInputSpdApplyActionField *pInputSpdApplyAction, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

    void COnRspQrySpdApply(void *this, struct CThostFtdcSpdApplyField *pSpdApply, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

    void COnRtnSpdApply(void *this, struct CThostFtdcSpdApplyField *pSpdApply);

    void COnErrRtnSpdApply(void *this, struct CThostFtdcInputSpdApplyField *pInputSpdApply, struct CThostFtdcRspInfoField *pRspInfo);

    void COnErrRtnSpdApplyAction(void *this, struct CThostFtdcSpdApplyActionField *pSpdApplyAction, struct CThostFtdcRspInfoField *pRspInfo);

    void COnRspHedgeCfm(void *this, struct CThostFtdcInputHedgeCfmField *pInputHedgeCfm, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

    void COnRspHedgeCfmAction(void *this, struct CThostFtdcInputHedgeCfmActionField *pInputHedgeCfmAction, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

    void COnRspQryHedgeCfm(void *this, struct CThostFtdcHedgeCfmField *pHedgeCfm, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

    void COnRtnHedgeCfm(void *this, struct CThostFtdcHedgeCfmField *pHedgeCfm);

    void COnErrRtnHedgeCfm(void *this, struct CThostFtdcInputHedgeCfmField *pInputHedgeCfm, struct CThostFtdcRspInfoField *pRspInfo);

    void COnErrRtnHedgeCfmAction(void *this, struct CThostFtdcHedgeCfmActionField *pHedgeCfmAction, struct CThostFtdcRspInfoField *pRspInfo);

    typedef struct
    {
        OnFrontConnected CThostFtdcTraderSpi_OnFrontConnected;
        OnFrontDisconnected CThostFtdcTraderSpi_OnFrontDisconnected;
        OnHeartBeatWarning CThostFtdcTraderSpi_OnHeartBeatWarning;
        OnRspAuthenticate CThostFtdcTraderSpi_OnRspAuthenticate;
        OnRtnPrivateSeqNo CThostFtdcTraderSpi_OnRtnPrivateSeqNo;
        OnRspUserLogin CThostFtdcTraderSpi_OnRspUserLogin;
        OnRspUserLogout CThostFtdcTraderSpi_OnRspUserLogout;
        OnRspUserPasswordUpdate CThostFtdcTraderSpi_OnRspUserPasswordUpdate;
        OnRspTradingAccountPasswordUpdate CThostFtdcTraderSpi_OnRspTradingAccountPasswordUpdate;
        OnRspUserAuthMethod CThostFtdcTraderSpi_OnRspUserAuthMethod;
        OnRspGenUserCaptcha CThostFtdcTraderSpi_OnRspGenUserCaptcha;
        OnRspGenUserText CThostFtdcTraderSpi_OnRspGenUserText;
        OnRspOrderInsert CThostFtdcTraderSpi_OnRspOrderInsert;
        OnRspParkedOrderInsert CThostFtdcTraderSpi_OnRspParkedOrderInsert;
        OnRspParkedOrderAction CThostFtdcTraderSpi_OnRspParkedOrderAction;
        OnRspOrderAction CThostFtdcTraderSpi_OnRspOrderAction;
        OnRspQryMaxOrderVolume CThostFtdcTraderSpi_OnRspQryMaxOrderVolume;
        OnRspSettlementInfoConfirm CThostFtdcTraderSpi_OnRspSettlementInfoConfirm;
        OnRspRemoveParkedOrder CThostFtdcTraderSpi_OnRspRemoveParkedOrder;
        OnRspRemoveParkedOrderAction CThostFtdcTraderSpi_OnRspRemoveParkedOrderAction;
        OnRspExecOrderInsert CThostFtdcTraderSpi_OnRspExecOrderInsert;
        OnRspExecOrderAction CThostFtdcTraderSpi_OnRspExecOrderAction;
        OnRspForQuoteInsert CThostFtdcTraderSpi_OnRspForQuoteInsert;
        OnRspQuoteInsert CThostFtdcTraderSpi_OnRspQuoteInsert;
        OnRspQuoteAction CThostFtdcTraderSpi_OnRspQuoteAction;
        OnRspBatchOrderAction CThostFtdcTraderSpi_OnRspBatchOrderAction;
        OnRspOptionSelfCloseInsert CThostFtdcTraderSpi_OnRspOptionSelfCloseInsert;
        OnRspOptionSelfCloseAction CThostFtdcTraderSpi_OnRspOptionSelfCloseAction;
        OnRspCombActionInsert CThostFtdcTraderSpi_OnRspCombActionInsert;
        OnRspQryOrder CThostFtdcTraderSpi_OnRspQryOrder;
        OnRspQryTrade CThostFtdcTraderSpi_OnRspQryTrade;
        OnRspQryInvestorPosition CThostFtdcTraderSpi_OnRspQryInvestorPosition;
        OnRspQryTradingAccount CThostFtdcTraderSpi_OnRspQryTradingAccount;
        OnRspQryInvestor CThostFtdcTraderSpi_OnRspQryInvestor;
        OnRspQryTradingCode CThostFtdcTraderSpi_OnRspQryTradingCode;
        OnRspQryInstrumentMarginRate CThostFtdcTraderSpi_OnRspQryInstrumentMarginRate;
        OnRspQryInstrumentCommissionRate CThostFtdcTraderSpi_OnRspQryInstrumentCommissionRate;
        OnRspQryUserSession CThostFtdcTraderSpi_OnRspQryUserSession;
        OnRspQryExchange CThostFtdcTraderSpi_OnRspQryExchange;
        OnRspQryProduct CThostFtdcTraderSpi_OnRspQryProduct;
        OnRspQryInstrument CThostFtdcTraderSpi_OnRspQryInstrument;
        OnRspQryDepthMarketData CThostFtdcTraderSpi_OnRspQryDepthMarketData;
        OnRspQryTraderOffer CThostFtdcTraderSpi_OnRspQryTraderOffer;
        OnRspQrySettlementInfo CThostFtdcTraderSpi_OnRspQrySettlementInfo;
        OnRspQryTransferBank CThostFtdcTraderSpi_OnRspQryTransferBank;
        OnRspQryInvestorPositionDetail CThostFtdcTraderSpi_OnRspQryInvestorPositionDetail;
        OnRspQryNotice CThostFtdcTraderSpi_OnRspQryNotice;
        OnRspQrySettlementInfoConfirm CThostFtdcTraderSpi_OnRspQrySettlementInfoConfirm;
        OnRspQryInvestorPositionCombineDetail CThostFtdcTraderSpi_OnRspQryInvestorPositionCombineDetail;
        OnRspQryCFMMCTradingAccountKey CThostFtdcTraderSpi_OnRspQryCFMMCTradingAccountKey;
        OnRspQryEWarrantOffset CThostFtdcTraderSpi_OnRspQryEWarrantOffset;
        OnRspQryInvestorProductGroupMargin CThostFtdcTraderSpi_OnRspQryInvestorProductGroupMargin;
        OnRspQryExchangeMarginRate CThostFtdcTraderSpi_OnRspQryExchangeMarginRate;
        OnRspQryExchangeMarginRateAdjust CThostFtdcTraderSpi_OnRspQryExchangeMarginRateAdjust;
        OnRspQryExchangeRate CThostFtdcTraderSpi_OnRspQryExchangeRate;
        OnRspQrySecAgentACIDMap CThostFtdcTraderSpi_OnRspQrySecAgentACIDMap;
        OnRspQryProductExchRate CThostFtdcTraderSpi_OnRspQryProductExchRate;
        OnRspQryProductGroup CThostFtdcTraderSpi_OnRspQryProductGroup;
        OnRspQryMMInstrumentCommissionRate CThostFtdcTraderSpi_OnRspQryMMInstrumentCommissionRate;
        OnRspQryMMOptionInstrCommRate CThostFtdcTraderSpi_OnRspQryMMOptionInstrCommRate;
        OnRspQryInstrumentOrderCommRate CThostFtdcTraderSpi_OnRspQryInstrumentOrderCommRate;
        OnRspQrySecAgentTradingAccount CThostFtdcTraderSpi_OnRspQrySecAgentTradingAccount;
        OnRspQrySecAgentCheckMode CThostFtdcTraderSpi_OnRspQrySecAgentCheckMode;
        OnRspQrySecAgentTradeInfo CThostFtdcTraderSpi_OnRspQrySecAgentTradeInfo;
        OnRspQryOptionInstrTradeCost CThostFtdcTraderSpi_OnRspQryOptionInstrTradeCost;
        OnRspQryOptionInstrCommRate CThostFtdcTraderSpi_OnRspQryOptionInstrCommRate;
        OnRspQryExecOrder CThostFtdcTraderSpi_OnRspQryExecOrder;
        OnRspQryForQuote CThostFtdcTraderSpi_OnRspQryForQuote;
        OnRspQryQuote CThostFtdcTraderSpi_OnRspQryQuote;
        OnRspQryOptionSelfClose CThostFtdcTraderSpi_OnRspQryOptionSelfClose;
        OnRspQryInvestUnit CThostFtdcTraderSpi_OnRspQryInvestUnit;
        OnRspQryCombInstrumentGuard CThostFtdcTraderSpi_OnRspQryCombInstrumentGuard;
        OnRspQryCombAction CThostFtdcTraderSpi_OnRspQryCombAction;
        OnRspQryTransferSerial CThostFtdcTraderSpi_OnRspQryTransferSerial;
        OnRspQryAccountregister CThostFtdcTraderSpi_OnRspQryAccountregister;
        OnRspError CThostFtdcTraderSpi_OnRspError;
        OnRtnOrder CThostFtdcTraderSpi_OnRtnOrder;
        OnRtnTrade CThostFtdcTraderSpi_OnRtnTrade;
        OnErrRtnOrderInsert CThostFtdcTraderSpi_OnErrRtnOrderInsert;
        OnErrRtnOrderAction CThostFtdcTraderSpi_OnErrRtnOrderAction;
        OnRtnInstrumentStatus CThostFtdcTraderSpi_OnRtnInstrumentStatus;
        OnRtnBulletin CThostFtdcTraderSpi_OnRtnBulletin;
        OnRtnTradingNotice CThostFtdcTraderSpi_OnRtnTradingNotice;
        OnRtnErrorConditionalOrder CThostFtdcTraderSpi_OnRtnErrorConditionalOrder;
        OnRtnExecOrder CThostFtdcTraderSpi_OnRtnExecOrder;
        OnErrRtnExecOrderInsert CThostFtdcTraderSpi_OnErrRtnExecOrderInsert;
        OnErrRtnExecOrderAction CThostFtdcTraderSpi_OnErrRtnExecOrderAction;
        OnErrRtnForQuoteInsert CThostFtdcTraderSpi_OnErrRtnForQuoteInsert;
        OnRtnQuote CThostFtdcTraderSpi_OnRtnQuote;
        OnErrRtnQuoteInsert CThostFtdcTraderSpi_OnErrRtnQuoteInsert;
        OnErrRtnQuoteAction CThostFtdcTraderSpi_OnErrRtnQuoteAction;
        OnRtnForQuoteRsp CThostFtdcTraderSpi_OnRtnForQuoteRsp;
        OnRtnCFMMCTradingAccountToken CThostFtdcTraderSpi_OnRtnCFMMCTradingAccountToken;
        OnErrRtnBatchOrderAction CThostFtdcTraderSpi_OnErrRtnBatchOrderAction;
        OnRtnOptionSelfClose CThostFtdcTraderSpi_OnRtnOptionSelfClose;
        OnErrRtnOptionSelfCloseInsert CThostFtdcTraderSpi_OnErrRtnOptionSelfCloseInsert;
        OnErrRtnOptionSelfCloseAction CThostFtdcTraderSpi_OnErrRtnOptionSelfCloseAction;
        OnRtnCombAction CThostFtdcTraderSpi_OnRtnCombAction;
        OnErrRtnCombActionInsert CThostFtdcTraderSpi_OnErrRtnCombActionInsert;
        OnRspQryContractBank CThostFtdcTraderSpi_OnRspQryContractBank;
        OnRspQryParkedOrder CThostFtdcTraderSpi_OnRspQryParkedOrder;
        OnRspQryParkedOrderAction CThostFtdcTraderSpi_OnRspQryParkedOrderAction;
        OnRspQryTradingNotice CThostFtdcTraderSpi_OnRspQryTradingNotice;
        OnRspQryBrokerTradingParams CThostFtdcTraderSpi_OnRspQryBrokerTradingParams;
        OnRspQryBrokerTradingAlgos CThostFtdcTraderSpi_OnRspQryBrokerTradingAlgos;
        OnRspQueryCFMMCTradingAccountToken CThostFtdcTraderSpi_OnRspQueryCFMMCTradingAccountToken;
        OnRtnFromBankToFutureByBank CThostFtdcTraderSpi_OnRtnFromBankToFutureByBank;
        OnRtnFromFutureToBankByBank CThostFtdcTraderSpi_OnRtnFromFutureToBankByBank;
        OnRtnRepealFromBankToFutureByBank CThostFtdcTraderSpi_OnRtnRepealFromBankToFutureByBank;
        OnRtnRepealFromFutureToBankByBank CThostFtdcTraderSpi_OnRtnRepealFromFutureToBankByBank;
        OnRtnFromBankToFutureByFuture CThostFtdcTraderSpi_OnRtnFromBankToFutureByFuture;
        OnRtnFromFutureToBankByFuture CThostFtdcTraderSpi_OnRtnFromFutureToBankByFuture;
        OnRtnRepealFromBankToFutureByFutureManual CThostFtdcTraderSpi_OnRtnRepealFromBankToFutureByFutureManual;
        OnRtnRepealFromFutureToBankByFutureManual CThostFtdcTraderSpi_OnRtnRepealFromFutureToBankByFutureManual;
        OnRtnQueryBankBalanceByFuture CThostFtdcTraderSpi_OnRtnQueryBankBalanceByFuture;
        OnErrRtnBankToFutureByFuture CThostFtdcTraderSpi_OnErrRtnBankToFutureByFuture;
        OnErrRtnFutureToBankByFuture CThostFtdcTraderSpi_OnErrRtnFutureToBankByFuture;
        OnErrRtnRepealBankToFutureByFutureManual CThostFtdcTraderSpi_OnErrRtnRepealBankToFutureByFutureManual;
        OnErrRtnRepealFutureToBankByFutureManual CThostFtdcTraderSpi_OnErrRtnRepealFutureToBankByFutureManual;
        OnErrRtnQueryBankBalanceByFuture CThostFtdcTraderSpi_OnErrRtnQueryBankBalanceByFuture;
        OnRtnRepealFromBankToFutureByFuture CThostFtdcTraderSpi_OnRtnRepealFromBankToFutureByFuture;
        OnRtnRepealFromFutureToBankByFuture CThostFtdcTraderSpi_OnRtnRepealFromFutureToBankByFuture;
        OnRspFromBankToFutureByFuture CThostFtdcTraderSpi_OnRspFromBankToFutureByFuture;
        OnRspFromFutureToBankByFuture CThostFtdcTraderSpi_OnRspFromFutureToBankByFuture;
        OnRspQueryBankAccountMoneyByFuture CThostFtdcTraderSpi_OnRspQueryBankAccountMoneyByFuture;
        OnRtnOpenAccountByBank CThostFtdcTraderSpi_OnRtnOpenAccountByBank;
        OnRtnCancelAccountByBank CThostFtdcTraderSpi_OnRtnCancelAccountByBank;
        OnRtnChangeAccountByBank CThostFtdcTraderSpi_OnRtnChangeAccountByBank;
        OnRspQryClassifiedInstrument CThostFtdcTraderSpi_OnRspQryClassifiedInstrument;
        OnRspQryCombPromotionParam CThostFtdcTraderSpi_OnRspQryCombPromotionParam;
        OnRspQryRiskSettleInvstPosition CThostFtdcTraderSpi_OnRspQryRiskSettleInvstPosition;
        OnRspQryRiskSettleProductStatus CThostFtdcTraderSpi_OnRspQryRiskSettleProductStatus;
        OnRspQrySPBMFutureParameter CThostFtdcTraderSpi_OnRspQrySPBMFutureParameter;
        OnRspQrySPBMOptionParameter CThostFtdcTraderSpi_OnRspQrySPBMOptionParameter;
        OnRspQrySPBMIntraParameter CThostFtdcTraderSpi_OnRspQrySPBMIntraParameter;
        OnRspQrySPBMInterParameter CThostFtdcTraderSpi_OnRspQrySPBMInterParameter;
        OnRspQrySPBMPortfDefinition CThostFtdcTraderSpi_OnRspQrySPBMPortfDefinition;
        OnRspQrySPBMInvestorPortfDef CThostFtdcTraderSpi_OnRspQrySPBMInvestorPortfDef;
        OnRspQryInvestorPortfMarginRatio CThostFtdcTraderSpi_OnRspQryInvestorPortfMarginRatio;
        OnRspQryInvestorProdSPBMDetail CThostFtdcTraderSpi_OnRspQryInvestorProdSPBMDetail;
        OnRspQryInvestorCommoditySPMMMargin CThostFtdcTraderSpi_OnRspQryInvestorCommoditySPMMMargin;
        OnRspQryInvestorCommodityGroupSPMMMargin CThostFtdcTraderSpi_OnRspQryInvestorCommodityGroupSPMMMargin;
        OnRspQrySPMMInstParam CThostFtdcTraderSpi_OnRspQrySPMMInstParam;
        OnRspQrySPMMProductParam CThostFtdcTraderSpi_OnRspQrySPMMProductParam;
        OnRspQrySPBMAddOnInterParameter CThostFtdcTraderSpi_OnRspQrySPBMAddOnInterParameter;
        OnRspQryRCAMSCombProductInfo CThostFtdcTraderSpi_OnRspQryRCAMSCombProductInfo;
        OnRspQryRCAMSInstrParameter CThostFtdcTraderSpi_OnRspQryRCAMSInstrParameter;
        OnRspQryRCAMSIntraParameter CThostFtdcTraderSpi_OnRspQryRCAMSIntraParameter;
        OnRspQryRCAMSInterParameter CThostFtdcTraderSpi_OnRspQryRCAMSInterParameter;
        OnRspQryRCAMSShortOptAdjustParam CThostFtdcTraderSpi_OnRspQryRCAMSShortOptAdjustParam;
        OnRspQryRCAMSInvestorCombPosition CThostFtdcTraderSpi_OnRspQryRCAMSInvestorCombPosition;
        OnRspQryInvestorProdRCAMSMargin CThostFtdcTraderSpi_OnRspQryInvestorProdRCAMSMargin;
        OnRspQryRULEInstrParameter CThostFtdcTraderSpi_OnRspQryRULEInstrParameter;
        OnRspQryRULEIntraParameter CThostFtdcTraderSpi_OnRspQryRULEIntraParameter;
        OnRspQryRULEInterParameter CThostFtdcTraderSpi_OnRspQryRULEInterParameter;
        OnRspQryInvestorProdRULEMargin CThostFtdcTraderSpi_OnRspQryInvestorProdRULEMargin;
        OnRspQryInvestorPortfSetting CThostFtdcTraderSpi_OnRspQryInvestorPortfSetting;
        OnRspQryInvestorInfoCommRec CThostFtdcTraderSpi_OnRspQryInvestorInfoCommRec;
        OnRspQryCombLeg CThostFtdcTraderSpi_OnRspQryCombLeg;
        OnRspOffsetSetting CThostFtdcTraderSpi_OnRspOffsetSetting;
        OnRspCancelOffsetSetting CThostFtdcTraderSpi_OnRspCancelOffsetSetting;
        OnRtnOffsetSetting CThostFtdcTraderSpi_OnRtnOffsetSetting;
        OnErrRtnOffsetSetting CThostFtdcTraderSpi_OnErrRtnOffsetSetting;
        OnErrRtnCancelOffsetSetting CThostFtdcTraderSpi_OnErrRtnCancelOffsetSetting;
        OnRspQryOffsetSetting CThostFtdcTraderSpi_OnRspQryOffsetSetting;
        OnRspGenSMSCode CThostFtdcTraderSpi_OnRspGenSMSCode;
        OnRspSpdApply CThostFtdcTraderSpi_OnRspSpdApply;
        OnRspSpdApplyAction CThostFtdcTraderSpi_OnRspSpdApplyAction;
        OnRspQrySpdApply CThostFtdcTraderSpi_OnRspQrySpdApply;
        OnRtnSpdApply CThostFtdcTraderSpi_OnRtnSpdApply;
        OnErrRtnSpdApply CThostFtdcTraderSpi_OnErrRtnSpdApply;
        OnErrRtnSpdApplyAction CThostFtdcTraderSpi_OnErrRtnSpdApplyAction;
        OnRspHedgeCfm CThostFtdcTraderSpi_OnRspHedgeCfm;
        OnRspHedgeCfmAction CThostFtdcTraderSpi_OnRspHedgeCfmAction;
        OnRspQryHedgeCfm CThostFtdcTraderSpi_OnRspQryHedgeCfm;
        OnRtnHedgeCfm CThostFtdcTraderSpi_OnRtnHedgeCfm;
        OnErrRtnHedgeCfm CThostFtdcTraderSpi_OnErrRtnHedgeCfm;
        OnErrRtnHedgeCfmAction CThostFtdcTraderSpi_OnErrRtnHedgeCfmAction;
    } CThostFtdcTraderSpiVTable;

    typedef struct
    {
        CThostFtdcTraderSpiVTable *vtable;
        void *spi;
    } CThostFtdcTraderSpiExt;

#ifdef __cplusplus
}
#endif

#endif