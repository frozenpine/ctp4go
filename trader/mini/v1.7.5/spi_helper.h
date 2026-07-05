#pragma once
#ifndef TRADER_V175_SPI_HELPER_H
#define TRADER_V175_SPI_HELPER_H

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

// converted from ..\dependencies\mini\v1.7.5\ThostFtdcTraderApi.h
// function call ptr
typedef void (*OnFrontConnected)(void *this);

typedef void (*OnFrontDisconnected)(void *this, int nReason);

typedef void (*OnHeartBeatWarning)(void *this, int nTimeLapse);

typedef void (*OnRspSubscribeFlowCtrlWarning)(void *this, struct CThostFtdcSpecificTraderField* pRspSubscribeTraderField, struct CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast);

typedef void (*OnRspUnSubscribeFlowCtrlWarning)(void *this, struct CThostFtdcSpecificTraderField* pRspSubscribeTraderField, struct CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast);

typedef void (*OnRspAuthenticate)(void *this, struct CThostFtdcRspAuthenticateField* pRspAuthenticateField, struct CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast);

typedef void (*OnRspUserLogin)(void *this, struct CThostFtdcRspUserLoginField* pRspUserLogin, struct CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast);

typedef void (*OnRspUserLogout)(void *this, struct CThostFtdcUserLogoutField* pUserLogout, struct CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast);

typedef void (*OnRspOrderInsert)(void *this, struct CThostFtdcInputOrderField* pInputOrder, struct CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast);

typedef void (*OnRspOrderAction)(void *this, struct CThostFtdcInputOrderActionField* pInputOrderAction, struct CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast);

typedef void (*OnRspMKBatchOrderAction)(void *this, struct CThostFtdcMKInputOrderActionField* pMKInputOrderAction, struct CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast);

typedef void (*OnRspExecOrderInsert)(void *this, struct CThostFtdcInputExecOrderField* pInputExecOrder, struct CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast);

typedef void (*OnRspExecOrderAction)(void *this, struct CThostFtdcInputExecOrderActionField* pInputExecOrderAction, struct CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast);

typedef void (*OnRspForQuoteInsert)(void *this, struct CThostFtdcInputForQuoteField* pInputForQuote, struct CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast);

typedef void (*OnRspQuoteInsert)(void *this, struct CThostFtdcInputQuoteField* pInputQuote, struct CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast);

typedef void (*OnRspQuoteAction)(void *this, struct CThostFtdcInputQuoteActionField* pInputQuoteAction, struct CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast);

typedef void (*OnRspBatchOrderAction)(void *this, struct CThostFtdcInputBatchOrderActionField* pInputBatchOrderAction, struct CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast);

typedef void (*OnRspOptionSelfCloseInsert)(void *this, struct CThostFtdcInputOptionSelfCloseField* pInputOptionSelfClose, struct CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast);

typedef void (*OnRspOptionSelfCloseAction)(void *this, struct CThostFtdcInputOptionSelfCloseActionField* pInputOptionSelfCloseAction, struct CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast);

typedef void (*OnRspCombActionInsert)(void *this, struct CThostFtdcInputCombActionField* pInputCombAction, struct CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast);

typedef void (*OnRspQryOrder)(void *this, struct CThostFtdcOrderField* pOrder, struct CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast);

typedef void (*OnRspQryTrade)(void *this, struct CThostFtdcTradeField* pTrade, struct CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast);

typedef void (*OnRspQryInvestorPosition)(void *this, struct CThostFtdcInvestorPositionField* pInvestorPosition, struct CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast);

typedef void (*OnRspQryTradingAccount)(void *this, struct CThostFtdcTradingAccountField* pTradingAccount, struct CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast);

typedef void (*OnRspQryInvestor)(void *this, struct CThostFtdcInvestorField* pInvestor, struct CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast);

typedef void (*OnRspQryTradingCode)(void *this, struct CThostFtdcTradingCodeField* pTradingCode, struct CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast);

typedef void (*OnRspQryInstrumentMarginRate)(void *this, struct CThostFtdcInstrumentMarginRateField* pInstrumentMarginRate, struct CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast);

typedef void (*OnRspQryInstrumentCommissionRate)(void *this, struct CThostFtdcInstrumentCommissionRateField* pInstrumentCommissionRate, struct CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast);

typedef void (*OnRspQryExchange)(void *this, struct CThostFtdcExchangeField* pExchange, struct CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast);

typedef void (*OnRspQryProduct)(void *this, struct CThostFtdcProductField* pProduct, struct CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast);

typedef void (*OnRspQryInstrument)(void *this, struct CThostFtdcInstrumentField* pInstrument, struct CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast);

typedef void (*OnRspQryCombInstrument)(void *this, struct CThostFtdcCombInstrumentField* pCombInstrument, struct CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast);

typedef void (*OnRspQryRCAMSInvestorProdMargin)(void *this, struct CThostFtdcRCAMSInvestorProdMarginField* pRCAMSInvestorProdMargin, struct CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast);

typedef void (*OnRspQryRCAMSInvestorCombPosition)(void *this, struct CThostFtdcRCAMSInvestorCombPositionField* pRCAMSInvestorCombPosition, struct CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast);

typedef void (*OnRspQryCombAction)(void *this, struct CThostFtdcCombActionField* pCombAction, struct CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast);

typedef void (*OnRspQryInvestorPositionForComb)(void *this, struct CThostFtdcInvestorPositionForCombField* pForComb, struct CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast);

typedef void (*OnRspQryDepthMarketData)(void *this, struct CThostFtdcDepthMarketDataField* pDepthMarketData, struct CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast);

typedef void (*OnRspQryInstrumentStatus)(void *this, struct CThostFtdcInstrumentStatusField* pInstrumentStatus, struct CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast);

typedef void (*OnRspQryInvestorPositionDetail)(void *this, struct CThostFtdcInvestorPositionDetailField* pInvestorPositionDetail, struct CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast);

typedef void (*OnRspQryExchangeMarginRate)(void *this, struct CThostFtdcExchangeMarginRateField* pExchangeMarginRate, struct CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast);

typedef void (*OnRspQryExchangeMarginRateAdjust)(void *this, struct CThostFtdcExchangeMarginRateAdjustField* pExchangeMarginRateAdjust, struct CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast);

typedef void (*OnRspQryOptionInstrTradeCost)(void *this, struct CThostFtdcOptionInstrTradeCostField* pOptionInstrTradeCost, struct CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast);

typedef void (*OnRspQryOptionInstrCommRate)(void *this, struct CThostFtdcOptionInstrCommRateField* pOptionInstrCommRate, struct CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast);

typedef void (*OnRspQryExecOrder)(void *this, struct CThostFtdcExecOrderField* pExecOrder, struct CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast);

typedef void (*OnRspQryForQuote)(void *this, struct CThostFtdcForQuoteField* pForQuote, struct CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast);

typedef void (*OnRspQryForQuoteParam)(void *this, struct CThostFtdcForQuoteParamField* pForQuoteParam, struct CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast);

typedef void (*OnRspQryInvestorProdSPBMDetail)(void *this, struct CThostFtdcInvestorProdSPBMDetailField* pInvestorProdSPBMDetail, struct CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast);

typedef void (*OnRspQrySPMMInvestorCommodityGroupMargin)(void *this, struct CThostFtdcSPMMInvestorCommodityGroupMarginField* pSPMMInvestorCommodityGroupMargin, struct CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast);

typedef void (*OnRspQryRULEInvestorProdMargin)(void *this, struct CThostFtdcRULEInvestorProdMarginField* pRULEInvestorProdMargin, struct CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast);

typedef void (*OnRspQryTraderOffer)(void *this, struct CThostFtdcTraderOfferField* pTraderOffer, struct CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast);

typedef void (*OnRspQryQuote)(void *this, struct CThostFtdcQuoteField* pQuote, struct CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast);

typedef void (*OnRspQryOptionSelfClose)(void *this, struct CThostFtdcOptionSelfCloseField* pOptionSelfClose, struct CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast);

typedef void (*OnRspQryControlParam)(void *this, struct CThostFtdcControlParamField* pControlParam, struct CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast);

typedef void (*OnRspQryOffsetSetting)(void *this, struct CThostFtdcOffsetSettingField* pOffsetSetting, struct CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast);

typedef void (*OnRspError)(void *this, struct CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast);

typedef void (*OnRtnOrder)(void *this, struct CThostFtdcOrderField* pOrder);

typedef void (*OnRtnTrade)(void *this, struct CThostFtdcTradeField* pTrade);

typedef void (*OnErrRtnOrderInsert)(void *this, struct CThostFtdcInputOrderField* pInputOrder, struct CThostFtdcRspInfoField* pRspInfo);

typedef void (*OnErrRtnOrderAction)(void *this, struct CThostFtdcOrderActionField* pOrderAction, struct CThostFtdcRspInfoField* pRspInfo);

typedef void (*OnRtnInstrumentStatus)(void *this, struct CThostFtdcInstrumentStatusField* pInstrumentStatus);

typedef void (*OnRtnExecOrder)(void *this, struct CThostFtdcExecOrderField* pExecOrder);

typedef void (*OnErrRtnExecOrderInsert)(void *this, struct CThostFtdcInputExecOrderField* pInputExecOrder, struct CThostFtdcRspInfoField* pRspInfo);

typedef void (*OnErrRtnExecOrderAction)(void *this, struct CThostFtdcExecOrderActionField* pExecOrderAction, struct CThostFtdcRspInfoField* pRspInfo);

typedef void (*OnErrRtnForQuoteInsert)(void *this, struct CThostFtdcInputForQuoteField* pInputForQuote, struct CThostFtdcRspInfoField* pRspInfo);

typedef void (*OnRtnQuote)(void *this, struct CThostFtdcQuoteField* pQuote);

typedef void (*OnErrRtnQuoteInsert)(void *this, struct CThostFtdcInputQuoteField* pInputQuote, struct CThostFtdcRspInfoField* pRspInfo);

typedef void (*OnErrRtnQuoteAction)(void *this, struct CThostFtdcQuoteActionField* pQuoteAction, struct CThostFtdcRspInfoField* pRspInfo);

typedef void (*OnRtnForQuoteRsp)(void *this, struct CThostFtdcForQuoteRspField* pForQuoteRsp);

typedef void (*OnErrRtnBatchOrderAction)(void *this, struct CThostFtdcBatchOrderActionField* pBatchOrderAction, struct CThostFtdcRspInfoField* pRspInfo);

typedef void (*OnRtnOptionSelfClose)(void *this, struct CThostFtdcOptionSelfCloseField* pOptionSelfClose);

typedef void (*OnErrRtnOptionSelfCloseInsert)(void *this, struct CThostFtdcInputOptionSelfCloseField* pInputOptionSelfClose, struct CThostFtdcRspInfoField* pRspInfo);

typedef void (*OnErrRtnOptionSelfCloseAction)(void *this, struct CThostFtdcOptionSelfCloseActionField* pOptionSelfCloseAction, struct CThostFtdcRspInfoField* pRspInfo);

typedef void (*OnRtnCombAction)(void *this, struct CThostFtdcCombActionField* pCombAction);

typedef void (*OnRspQryInstrumentOrderCommRate)(void *this, struct CThostFtdcInstrumentOrderCommRateField* pInstrumentOrderCommRate, struct CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast);

typedef void (*OnRtnFlowCtrlWarning)(void *this, struct CThostFtdcFlowCtrlWarningField* pFlowCtrlWarning);

typedef void (*OnRspSubscribeFundChange)(void *this, struct CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast);

typedef void (*OnRspUnSubscribeFundChange)(void *this, struct CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast);

typedef void (*OnRtnFundChange)(void *this, struct CThostFtdcTradingAccountField* pTradingAccount);

typedef void (*OnRspOffsetSetting)(void *this, struct CThostFtdcInputOffsetSettingField* pInputOffsetSetting, struct CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast);

typedef void (*OnRspCancelOffsetSetting)(void *this, struct CThostFtdcInputOffsetSettingField* pInputOffsetSetting, struct CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast);

typedef void (*OnRtnOffsetSetting)(void *this, struct CThostFtdcOffsetSettingField* pOffsetSetting);

typedef void (*OnErrRtnCancelOffsetSetting)(void *this, struct CThostFtdcCancelOffsetSettingField* pCancelOffsetSetting, struct CThostFtdcRspInfoField* pRspInfo);

typedef void (*OnRspUserPasswordUpdate)(void *this, struct CThostFtdcUserPasswordUpdateField* pUserPasswordUpdate, struct CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast);


// helper function callback for cgo
void COnFrontConnected(void *this);

void COnFrontDisconnected(void *this, int nReason);

void COnHeartBeatWarning(void *this, int nTimeLapse);

void COnRspSubscribeFlowCtrlWarning(void *this, struct CThostFtdcSpecificTraderField* pRspSubscribeTraderField, struct CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast);

void COnRspUnSubscribeFlowCtrlWarning(void *this, struct CThostFtdcSpecificTraderField* pRspSubscribeTraderField, struct CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast);

void COnRspAuthenticate(void *this, struct CThostFtdcRspAuthenticateField* pRspAuthenticateField, struct CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast);

void COnRspUserLogin(void *this, struct CThostFtdcRspUserLoginField* pRspUserLogin, struct CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast);

void COnRspUserLogout(void *this, struct CThostFtdcUserLogoutField* pUserLogout, struct CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast);

void COnRspOrderInsert(void *this, struct CThostFtdcInputOrderField* pInputOrder, struct CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast);

void COnRspOrderAction(void *this, struct CThostFtdcInputOrderActionField* pInputOrderAction, struct CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast);

void COnRspMKBatchOrderAction(void *this, struct CThostFtdcMKInputOrderActionField* pMKInputOrderAction, struct CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast);

void COnRspExecOrderInsert(void *this, struct CThostFtdcInputExecOrderField* pInputExecOrder, struct CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast);

void COnRspExecOrderAction(void *this, struct CThostFtdcInputExecOrderActionField* pInputExecOrderAction, struct CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast);

void COnRspForQuoteInsert(void *this, struct CThostFtdcInputForQuoteField* pInputForQuote, struct CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast);

void COnRspQuoteInsert(void *this, struct CThostFtdcInputQuoteField* pInputQuote, struct CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast);

void COnRspQuoteAction(void *this, struct CThostFtdcInputQuoteActionField* pInputQuoteAction, struct CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast);

void COnRspBatchOrderAction(void *this, struct CThostFtdcInputBatchOrderActionField* pInputBatchOrderAction, struct CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast);

void COnRspOptionSelfCloseInsert(void *this, struct CThostFtdcInputOptionSelfCloseField* pInputOptionSelfClose, struct CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast);

void COnRspOptionSelfCloseAction(void *this, struct CThostFtdcInputOptionSelfCloseActionField* pInputOptionSelfCloseAction, struct CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast);

void COnRspCombActionInsert(void *this, struct CThostFtdcInputCombActionField* pInputCombAction, struct CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast);

void COnRspQryOrder(void *this, struct CThostFtdcOrderField* pOrder, struct CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast);

void COnRspQryTrade(void *this, struct CThostFtdcTradeField* pTrade, struct CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast);

void COnRspQryInvestorPosition(void *this, struct CThostFtdcInvestorPositionField* pInvestorPosition, struct CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast);

void COnRspQryTradingAccount(void *this, struct CThostFtdcTradingAccountField* pTradingAccount, struct CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast);

void COnRspQryInvestor(void *this, struct CThostFtdcInvestorField* pInvestor, struct CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast);

void COnRspQryTradingCode(void *this, struct CThostFtdcTradingCodeField* pTradingCode, struct CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast);

void COnRspQryInstrumentMarginRate(void *this, struct CThostFtdcInstrumentMarginRateField* pInstrumentMarginRate, struct CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast);

void COnRspQryInstrumentCommissionRate(void *this, struct CThostFtdcInstrumentCommissionRateField* pInstrumentCommissionRate, struct CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast);

void COnRspQryExchange(void *this, struct CThostFtdcExchangeField* pExchange, struct CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast);

void COnRspQryProduct(void *this, struct CThostFtdcProductField* pProduct, struct CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast);

void COnRspQryInstrument(void *this, struct CThostFtdcInstrumentField* pInstrument, struct CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast);

void COnRspQryCombInstrument(void *this, struct CThostFtdcCombInstrumentField* pCombInstrument, struct CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast);

void COnRspQryRCAMSInvestorProdMargin(void *this, struct CThostFtdcRCAMSInvestorProdMarginField* pRCAMSInvestorProdMargin, struct CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast);

void COnRspQryRCAMSInvestorCombPosition(void *this, struct CThostFtdcRCAMSInvestorCombPositionField* pRCAMSInvestorCombPosition, struct CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast);

void COnRspQryCombAction(void *this, struct CThostFtdcCombActionField* pCombAction, struct CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast);

void COnRspQryInvestorPositionForComb(void *this, struct CThostFtdcInvestorPositionForCombField* pForComb, struct CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast);

void COnRspQryDepthMarketData(void *this, struct CThostFtdcDepthMarketDataField* pDepthMarketData, struct CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast);

void COnRspQryInstrumentStatus(void *this, struct CThostFtdcInstrumentStatusField* pInstrumentStatus, struct CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast);

void COnRspQryInvestorPositionDetail(void *this, struct CThostFtdcInvestorPositionDetailField* pInvestorPositionDetail, struct CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast);

void COnRspQryExchangeMarginRate(void *this, struct CThostFtdcExchangeMarginRateField* pExchangeMarginRate, struct CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast);

void COnRspQryExchangeMarginRateAdjust(void *this, struct CThostFtdcExchangeMarginRateAdjustField* pExchangeMarginRateAdjust, struct CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast);

void COnRspQryOptionInstrTradeCost(void *this, struct CThostFtdcOptionInstrTradeCostField* pOptionInstrTradeCost, struct CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast);

void COnRspQryOptionInstrCommRate(void *this, struct CThostFtdcOptionInstrCommRateField* pOptionInstrCommRate, struct CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast);

void COnRspQryExecOrder(void *this, struct CThostFtdcExecOrderField* pExecOrder, struct CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast);

void COnRspQryForQuote(void *this, struct CThostFtdcForQuoteField* pForQuote, struct CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast);

void COnRspQryForQuoteParam(void *this, struct CThostFtdcForQuoteParamField* pForQuoteParam, struct CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast);

void COnRspQryInvestorProdSPBMDetail(void *this, struct CThostFtdcInvestorProdSPBMDetailField* pInvestorProdSPBMDetail, struct CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast);

void COnRspQrySPMMInvestorCommodityGroupMargin(void *this, struct CThostFtdcSPMMInvestorCommodityGroupMarginField* pSPMMInvestorCommodityGroupMargin, struct CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast);

void COnRspQryRULEInvestorProdMargin(void *this, struct CThostFtdcRULEInvestorProdMarginField* pRULEInvestorProdMargin, struct CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast);

void COnRspQryTraderOffer(void *this, struct CThostFtdcTraderOfferField* pTraderOffer, struct CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast);

void COnRspQryQuote(void *this, struct CThostFtdcQuoteField* pQuote, struct CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast);

void COnRspQryOptionSelfClose(void *this, struct CThostFtdcOptionSelfCloseField* pOptionSelfClose, struct CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast);

void COnRspQryControlParam(void *this, struct CThostFtdcControlParamField* pControlParam, struct CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast);

void COnRspQryOffsetSetting(void *this, struct CThostFtdcOffsetSettingField* pOffsetSetting, struct CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast);

void COnRspError(void *this, struct CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast);

void COnRtnOrder(void *this, struct CThostFtdcOrderField* pOrder);

void COnRtnTrade(void *this, struct CThostFtdcTradeField* pTrade);

void COnErrRtnOrderInsert(void *this, struct CThostFtdcInputOrderField* pInputOrder, struct CThostFtdcRspInfoField* pRspInfo);

void COnErrRtnOrderAction(void *this, struct CThostFtdcOrderActionField* pOrderAction, struct CThostFtdcRspInfoField* pRspInfo);

void COnRtnInstrumentStatus(void *this, struct CThostFtdcInstrumentStatusField* pInstrumentStatus);

void COnRtnExecOrder(void *this, struct CThostFtdcExecOrderField* pExecOrder);

void COnErrRtnExecOrderInsert(void *this, struct CThostFtdcInputExecOrderField* pInputExecOrder, struct CThostFtdcRspInfoField* pRspInfo);

void COnErrRtnExecOrderAction(void *this, struct CThostFtdcExecOrderActionField* pExecOrderAction, struct CThostFtdcRspInfoField* pRspInfo);

void COnErrRtnForQuoteInsert(void *this, struct CThostFtdcInputForQuoteField* pInputForQuote, struct CThostFtdcRspInfoField* pRspInfo);

void COnRtnQuote(void *this, struct CThostFtdcQuoteField* pQuote);

void COnErrRtnQuoteInsert(void *this, struct CThostFtdcInputQuoteField* pInputQuote, struct CThostFtdcRspInfoField* pRspInfo);

void COnErrRtnQuoteAction(void *this, struct CThostFtdcQuoteActionField* pQuoteAction, struct CThostFtdcRspInfoField* pRspInfo);

void COnRtnForQuoteRsp(void *this, struct CThostFtdcForQuoteRspField* pForQuoteRsp);

void COnErrRtnBatchOrderAction(void *this, struct CThostFtdcBatchOrderActionField* pBatchOrderAction, struct CThostFtdcRspInfoField* pRspInfo);

void COnRtnOptionSelfClose(void *this, struct CThostFtdcOptionSelfCloseField* pOptionSelfClose);

void COnErrRtnOptionSelfCloseInsert(void *this, struct CThostFtdcInputOptionSelfCloseField* pInputOptionSelfClose, struct CThostFtdcRspInfoField* pRspInfo);

void COnErrRtnOptionSelfCloseAction(void *this, struct CThostFtdcOptionSelfCloseActionField* pOptionSelfCloseAction, struct CThostFtdcRspInfoField* pRspInfo);

void COnRtnCombAction(void *this, struct CThostFtdcCombActionField* pCombAction);

void COnRspQryInstrumentOrderCommRate(void *this, struct CThostFtdcInstrumentOrderCommRateField* pInstrumentOrderCommRate, struct CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast);

void COnRtnFlowCtrlWarning(void *this, struct CThostFtdcFlowCtrlWarningField* pFlowCtrlWarning);

void COnRspSubscribeFundChange(void *this, struct CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast);

void COnRspUnSubscribeFundChange(void *this, struct CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast);

void COnRtnFundChange(void *this, struct CThostFtdcTradingAccountField* pTradingAccount);

void COnRspOffsetSetting(void *this, struct CThostFtdcInputOffsetSettingField* pInputOffsetSetting, struct CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast);

void COnRspCancelOffsetSetting(void *this, struct CThostFtdcInputOffsetSettingField* pInputOffsetSetting, struct CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast);

void COnRtnOffsetSetting(void *this, struct CThostFtdcOffsetSettingField* pOffsetSetting);

void COnErrRtnCancelOffsetSetting(void *this, struct CThostFtdcCancelOffsetSettingField* pCancelOffsetSetting, struct CThostFtdcRspInfoField* pRspInfo);

void COnRspUserPasswordUpdate(void *this, struct CThostFtdcUserPasswordUpdateField* pUserPasswordUpdate, struct CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast);


// spi callback vtable
typedef struct
{
    OnFrontConnected CThostFtdcTraderSpiVTable_OnFrontConnected;

    OnFrontDisconnected CThostFtdcTraderSpiVTable_OnFrontDisconnected;

    OnHeartBeatWarning CThostFtdcTraderSpiVTable_OnHeartBeatWarning;

    OnRspSubscribeFlowCtrlWarning CThostFtdcTraderSpiVTable_OnRspSubscribeFlowCtrlWarning;

    OnRspUnSubscribeFlowCtrlWarning CThostFtdcTraderSpiVTable_OnRspUnSubscribeFlowCtrlWarning;

    OnRspAuthenticate CThostFtdcTraderSpiVTable_OnRspAuthenticate;

    OnRspUserLogin CThostFtdcTraderSpiVTable_OnRspUserLogin;

    OnRspUserLogout CThostFtdcTraderSpiVTable_OnRspUserLogout;

    OnRspOrderInsert CThostFtdcTraderSpiVTable_OnRspOrderInsert;

    OnRspOrderAction CThostFtdcTraderSpiVTable_OnRspOrderAction;

    OnRspMKBatchOrderAction CThostFtdcTraderSpiVTable_OnRspMKBatchOrderAction;

    OnRspExecOrderInsert CThostFtdcTraderSpiVTable_OnRspExecOrderInsert;

    OnRspExecOrderAction CThostFtdcTraderSpiVTable_OnRspExecOrderAction;

    OnRspForQuoteInsert CThostFtdcTraderSpiVTable_OnRspForQuoteInsert;

    OnRspQuoteInsert CThostFtdcTraderSpiVTable_OnRspQuoteInsert;

    OnRspQuoteAction CThostFtdcTraderSpiVTable_OnRspQuoteAction;

    OnRspBatchOrderAction CThostFtdcTraderSpiVTable_OnRspBatchOrderAction;

    OnRspOptionSelfCloseInsert CThostFtdcTraderSpiVTable_OnRspOptionSelfCloseInsert;

    OnRspOptionSelfCloseAction CThostFtdcTraderSpiVTable_OnRspOptionSelfCloseAction;

    OnRspCombActionInsert CThostFtdcTraderSpiVTable_OnRspCombActionInsert;

    OnRspQryOrder CThostFtdcTraderSpiVTable_OnRspQryOrder;

    OnRspQryTrade CThostFtdcTraderSpiVTable_OnRspQryTrade;

    OnRspQryInvestorPosition CThostFtdcTraderSpiVTable_OnRspQryInvestorPosition;

    OnRspQryTradingAccount CThostFtdcTraderSpiVTable_OnRspQryTradingAccount;

    OnRspQryInvestor CThostFtdcTraderSpiVTable_OnRspQryInvestor;

    OnRspQryTradingCode CThostFtdcTraderSpiVTable_OnRspQryTradingCode;

    OnRspQryInstrumentMarginRate CThostFtdcTraderSpiVTable_OnRspQryInstrumentMarginRate;

    OnRspQryInstrumentCommissionRate CThostFtdcTraderSpiVTable_OnRspQryInstrumentCommissionRate;

    OnRspQryExchange CThostFtdcTraderSpiVTable_OnRspQryExchange;

    OnRspQryProduct CThostFtdcTraderSpiVTable_OnRspQryProduct;

    OnRspQryInstrument CThostFtdcTraderSpiVTable_OnRspQryInstrument;

    OnRspQryCombInstrument CThostFtdcTraderSpiVTable_OnRspQryCombInstrument;

    OnRspQryRCAMSInvestorProdMargin CThostFtdcTraderSpiVTable_OnRspQryRCAMSInvestorProdMargin;

    OnRspQryRCAMSInvestorCombPosition CThostFtdcTraderSpiVTable_OnRspQryRCAMSInvestorCombPosition;

    OnRspQryCombAction CThostFtdcTraderSpiVTable_OnRspQryCombAction;

    OnRspQryInvestorPositionForComb CThostFtdcTraderSpiVTable_OnRspQryInvestorPositionForComb;

    OnRspQryDepthMarketData CThostFtdcTraderSpiVTable_OnRspQryDepthMarketData;

    OnRspQryInstrumentStatus CThostFtdcTraderSpiVTable_OnRspQryInstrumentStatus;

    OnRspQryInvestorPositionDetail CThostFtdcTraderSpiVTable_OnRspQryInvestorPositionDetail;

    OnRspQryExchangeMarginRate CThostFtdcTraderSpiVTable_OnRspQryExchangeMarginRate;

    OnRspQryExchangeMarginRateAdjust CThostFtdcTraderSpiVTable_OnRspQryExchangeMarginRateAdjust;

    OnRspQryOptionInstrTradeCost CThostFtdcTraderSpiVTable_OnRspQryOptionInstrTradeCost;

    OnRspQryOptionInstrCommRate CThostFtdcTraderSpiVTable_OnRspQryOptionInstrCommRate;

    OnRspQryExecOrder CThostFtdcTraderSpiVTable_OnRspQryExecOrder;

    OnRspQryForQuote CThostFtdcTraderSpiVTable_OnRspQryForQuote;

    OnRspQryForQuoteParam CThostFtdcTraderSpiVTable_OnRspQryForQuoteParam;

    OnRspQryInvestorProdSPBMDetail CThostFtdcTraderSpiVTable_OnRspQryInvestorProdSPBMDetail;

    OnRspQrySPMMInvestorCommodityGroupMargin CThostFtdcTraderSpiVTable_OnRspQrySPMMInvestorCommodityGroupMargin;

    OnRspQryRULEInvestorProdMargin CThostFtdcTraderSpiVTable_OnRspQryRULEInvestorProdMargin;

    OnRspQryTraderOffer CThostFtdcTraderSpiVTable_OnRspQryTraderOffer;

    OnRspQryQuote CThostFtdcTraderSpiVTable_OnRspQryQuote;

    OnRspQryOptionSelfClose CThostFtdcTraderSpiVTable_OnRspQryOptionSelfClose;

    OnRspQryControlParam CThostFtdcTraderSpiVTable_OnRspQryControlParam;

    OnRspQryOffsetSetting CThostFtdcTraderSpiVTable_OnRspQryOffsetSetting;

    OnRspError CThostFtdcTraderSpiVTable_OnRspError;

    OnRtnOrder CThostFtdcTraderSpiVTable_OnRtnOrder;

    OnRtnTrade CThostFtdcTraderSpiVTable_OnRtnTrade;

    OnErrRtnOrderInsert CThostFtdcTraderSpiVTable_OnErrRtnOrderInsert;

    OnErrRtnOrderAction CThostFtdcTraderSpiVTable_OnErrRtnOrderAction;

    OnRtnInstrumentStatus CThostFtdcTraderSpiVTable_OnRtnInstrumentStatus;

    OnRtnExecOrder CThostFtdcTraderSpiVTable_OnRtnExecOrder;

    OnErrRtnExecOrderInsert CThostFtdcTraderSpiVTable_OnErrRtnExecOrderInsert;

    OnErrRtnExecOrderAction CThostFtdcTraderSpiVTable_OnErrRtnExecOrderAction;

    OnErrRtnForQuoteInsert CThostFtdcTraderSpiVTable_OnErrRtnForQuoteInsert;

    OnRtnQuote CThostFtdcTraderSpiVTable_OnRtnQuote;

    OnErrRtnQuoteInsert CThostFtdcTraderSpiVTable_OnErrRtnQuoteInsert;

    OnErrRtnQuoteAction CThostFtdcTraderSpiVTable_OnErrRtnQuoteAction;

    OnRtnForQuoteRsp CThostFtdcTraderSpiVTable_OnRtnForQuoteRsp;

    OnErrRtnBatchOrderAction CThostFtdcTraderSpiVTable_OnErrRtnBatchOrderAction;

    OnRtnOptionSelfClose CThostFtdcTraderSpiVTable_OnRtnOptionSelfClose;

    OnErrRtnOptionSelfCloseInsert CThostFtdcTraderSpiVTable_OnErrRtnOptionSelfCloseInsert;

    OnErrRtnOptionSelfCloseAction CThostFtdcTraderSpiVTable_OnErrRtnOptionSelfCloseAction;

    OnRtnCombAction CThostFtdcTraderSpiVTable_OnRtnCombAction;

    OnRspQryInstrumentOrderCommRate CThostFtdcTraderSpiVTable_OnRspQryInstrumentOrderCommRate;

    OnRtnFlowCtrlWarning CThostFtdcTraderSpiVTable_OnRtnFlowCtrlWarning;

    OnRspSubscribeFundChange CThostFtdcTraderSpiVTable_OnRspSubscribeFundChange;

    OnRspUnSubscribeFundChange CThostFtdcTraderSpiVTable_OnRspUnSubscribeFundChange;

    OnRtnFundChange CThostFtdcTraderSpiVTable_OnRtnFundChange;

    OnRspOffsetSetting CThostFtdcTraderSpiVTable_OnRspOffsetSetting;

    OnRspCancelOffsetSetting CThostFtdcTraderSpiVTable_OnRspCancelOffsetSetting;

    OnRtnOffsetSetting CThostFtdcTraderSpiVTable_OnRtnOffsetSetting;

    OnErrRtnCancelOffsetSetting CThostFtdcTraderSpiVTable_OnErrRtnCancelOffsetSetting;

    OnRspUserPasswordUpdate CThostFtdcTraderSpiVTable_OnRspUserPasswordUpdate;

} CThostFtdcTraderSpiVTable;

// cpp spi instance C wrapper
typedef struct
{
    CThostFtdcTraderSpiVTable *vtable;
    // spi instance ptr in go, ptr must be Pinned
    void *spi;
} CThostFtdcTraderSpiExt;

#ifdef __cplusplus
}
#endif

#endif