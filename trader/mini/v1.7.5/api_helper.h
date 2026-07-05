#pragma once
#ifndef TRADER_V175_API_HELPER_H
#define TRADER_V175_API_HELPER_H

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
typedef void (*Release)(void *this);

typedef void (*Init)(void *this, bool bContinuous);

typedef int (*Join)(void *this);

typedef char* (*GetTradingDay)(void *this);

typedef void (*RegisterFront)(void *this, char* pszFrontAddress);

typedef void (*RegisterSpi)(void *this, void *pSpi);

typedef void (*SubscribePrivateTopic)(void *this, int nResumeType);

typedef void (*SubscribePublicTopic)(void *this, int nResumeType);

typedef int (*SubscribeFlowCtrlWarning)(void *this, char* ppTraderID[], int nCount);

typedef int (*UnSubscribeFlowCtrlWarning)(void *this, char* ppTraderID[], int nCount);

typedef int (*ReqAuthenticate)(void *this, struct CThostFtdcReqAuthenticateField* pReqAuthenticateField, int nRequestID);

typedef int (*ReqUserLogin)(void *this, struct CThostFtdcReqUserLoginField* pReqUserLoginField, int nRequestID);

typedef int (*ReqUserLoginEncrypt)(void *this, struct CThostFtdcReqUserLoginField* pReqUserLoginField, int nRequestID);

typedef int (*ReqUserLogout)(void *this, struct CThostFtdcUserLogoutField* pUserLogout, int nRequestID);

typedef int (*ReqOrderInsert)(void *this, struct CThostFtdcInputOrderField* pInputOrder, int nRequestID);

typedef int (*ReqOrderAction)(void *this, struct CThostFtdcInputOrderActionField* pInputOrderAction, int nRequestID);

typedef int (*ReqMKBatchOrderAction)(void *this, struct CThostFtdcMKInputOrderActionField* pMKInputOrderAction, int nRequestID);

typedef int (*ReqExecOrderInsert)(void *this, struct CThostFtdcInputExecOrderField* pInputExecOrder, int nRequestID);

typedef int (*ReqExecOrderAction)(void *this, struct CThostFtdcInputExecOrderActionField* pInputExecOrderAction, int nRequestID);

typedef int (*ReqForQuoteInsert)(void *this, struct CThostFtdcInputForQuoteField* pInputForQuote, int nRequestID);

typedef int (*ReqQuoteInsert)(void *this, struct CThostFtdcInputQuoteField* pInputQuote, int nRequestID);

typedef int (*ReqQuoteAction)(void *this, struct CThostFtdcInputQuoteActionField* pInputQuoteAction, int nRequestID);

typedef int (*ReqBatchOrderAction)(void *this, struct CThostFtdcInputBatchOrderActionField* pInputBatchOrderAction, int nRequestID);

typedef int (*ReqOptionSelfCloseInsert)(void *this, struct CThostFtdcInputOptionSelfCloseField* pInputOptionSelfClose, int nRequestID);

typedef int (*ReqOptionSelfCloseAction)(void *this, struct CThostFtdcInputOptionSelfCloseActionField* pInputOptionSelfCloseAction, int nRequestID);

typedef int (*ReqCombActionInsert)(void *this, struct CThostFtdcInputCombActionField* pInputCombAction, int nRequestID);

typedef int (*ReqSubscribeFundChange)(void *this, int nRequestID);

typedef int (*ReqUnSubscribeFundChange)(void *this, int nRequestID);

typedef int (*ReqOffsetSetting)(void *this, struct CThostFtdcInputOffsetSettingField* pInputOffsetSetting, int nRequestID);

typedef int (*ReqCancelOffsetSetting)(void *this, struct CThostFtdcInputOffsetSettingField* pInputOffsetSetting, int nRequestID);

typedef int (*ReqQryOrder)(void *this, struct CThostFtdcQryOrderField* pQryOrder, int nRequestID);

typedef int (*ReqQryTrade)(void *this, struct CThostFtdcQryTradeField* pQryTrade, int nRequestID);

typedef int (*ReqQryInvestorPosition)(void *this, struct CThostFtdcQryInvestorPositionField* pQryInvestorPosition, int nRequestID);

typedef int (*ReqQryTradingAccount)(void *this, struct CThostFtdcQryTradingAccountField* pQryTradingAccount, int nRequestID);

typedef int (*ReqQryInvestor)(void *this, struct CThostFtdcQryInvestorField* pQryInvestor, int nRequestID);

typedef int (*ReqQryTradingCode)(void *this, struct CThostFtdcQryTradingCodeField* pQryTradingCode, int nRequestID);

typedef int (*ReqQryInstrumentMarginRate)(void *this, struct CThostFtdcQryInstrumentMarginRateField* pQryInstrumentMarginRate, int nRequestID);

typedef int (*ReqQryInstrumentCommissionRate)(void *this, struct CThostFtdcQryInstrumentCommissionRateField* pQryInstrumentCommissionRate, int nRequestID);

typedef int (*ReqQryExchange)(void *this, struct CThostFtdcQryExchangeField* pQryExchange, int nRequestID);

typedef int (*ReqQryProduct)(void *this, struct CThostFtdcQryProductField* pQryProduct, int nRequestID);

typedef int (*ReqQryInstrument)(void *this, struct CThostFtdcQryInstrumentField* pQryInstrument, int nRequestID);

typedef int (*ReqQryCombInstrument)(void *this, struct CThostFtdcQryCombInstrumentField* pQryCombInstrument, int nRequestID);

typedef int (*ReqQryRCAMSInvestorProdMargin)(void *this, struct CThostFtdcQryRCAMSInvestorProdMarginField* pQryRCAMSInvestorProdMargin, int nRequestID);

typedef int (*ReqQryRCAMSInvestorCombPosition)(void *this, struct CThostFtdcQryRCAMSInvestorCombPositionField* pQryRCAMSInvestorCombPosition, int nRequestID);

typedef int (*ReqQryInvestorPositionForComb)(void *this, struct CThostFtdcQryInvestorPositionForCombField* pQryIPForComb, int nRequestID);

typedef int (*ReqQryCombAction)(void *this, struct CThostFtdcQryCombActionField* pQryCombAction, int nRequestID);

typedef int (*ReqQryDepthMarketData)(void *this, struct CThostFtdcQryDepthMarketDataField* pQryDepthMarketData, int nRequestID);

typedef int (*ReqQryOptionSelfClose)(void *this, struct CThostFtdcQryOptionSelfCloseField* pQryOptionSelfClose, int nRequestID);

typedef int (*ReqQryInstrumentStatus)(void *this, struct CThostFtdcQryInstrumentStatusField* pQryInstrumentStatus, int nRequestID);

typedef int (*ReqQryInvestorPositionDetail)(void *this, struct CThostFtdcQryInvestorPositionDetailField* pQryInvestorPositionDetail, int nRequestID);

typedef int (*ReqQryExchangeMarginRate)(void *this, struct CThostFtdcQryExchangeMarginRateField* pQryExchangeMarginRate, int nRequestID);

typedef int (*ReqQryExchangeMarginRateAdjust)(void *this, struct CThostFtdcQryExchangeMarginRateAdjustField* pQryExchangeMarginRateAdjust, int nRequestID);

typedef int (*ReqQryOptionInstrTradeCost)(void *this, struct CThostFtdcQryOptionInstrTradeCostField* pQryOptionInstrTradeCost, int nRequestID);

typedef int (*ReqQryOptionInstrCommRate)(void *this, struct CThostFtdcQryOptionInstrCommRateField* pQryOptionInstrCommRate, int nRequestID);

typedef int (*ReqQryExecOrder)(void *this, struct CThostFtdcQryExecOrderField* pQryExecOrder, int nRequestID);

typedef int (*ReqQryForQuote)(void *this, struct CThostFtdcQryForQuoteField* pQryForQuote, int nRequestID);

typedef int (*ReqQryQuote)(void *this, struct CThostFtdcQryQuoteField* pQryQuote, int nRequestID);

typedef int (*ReqQryInstrumentOrderCommRate)(void *this, struct CThostFtdcQryInstrumentOrderCommRateField* pQryInstrumentOrderCommRate, int nRequestID);

typedef int (*ReqQryForQuoteParam)(void *this, struct CThostFtdcQryForQuoteParamField* pQryForQuoteParam, int nRequestID);

typedef int (*ReqQryTraderOffer)(void *this, struct CThostFtdcQryTraderOfferField* pQryTraderOffer, int nRequestID);

typedef int (*ReqQryInvestorProdSPBMDetail)(void *this, struct CThostFtdcQryInvestorProdSPBMDetailField* pQryInvestorProdSPBMDetail, int nRequestID);

typedef int (*ReqQrySPMMInvestorCommodityGroupMargin)(void *this, struct CThostFtdcQrySPMMInvestorCommodityGroupMarginField* pQrySPMMInvestorCommodityGroupMargin, int nRequestID);

typedef int (*ReqQryRULEInvestorProdMargin)(void *this, struct CThostFtdcQryRULEInvestorProdMarginField* pQryRULEInvestorProdMargin, int nRequestID);

typedef int (*ReqQryControlParam)(void *this, struct CThostFtdcQryControlParamField* pQryControlParam, int nRequestID);

typedef int (*ReqQryOffsetSetting)(void *this, struct CThostFtdcQryOffsetSettingField* pQryOffsetSetting, int nRequestID);

typedef int (*ReqUserPasswordUpdate)(void *this, struct CThostFtdcUserPasswordUpdateField* pUserPasswordUpdate, int nRequestID);


// api methods all vtable
typedef struct
{
    Release CThostFtdcTraderApiVTable_Release;

    Init CThostFtdcTraderApiVTable_Init;

    Join CThostFtdcTraderApiVTable_Join;

    GetTradingDay CThostFtdcTraderApiVTable_GetTradingDay;

    RegisterFront CThostFtdcTraderApiVTable_RegisterFront;

    RegisterSpi CThostFtdcTraderApiVTable_RegisterSpi;

    SubscribePrivateTopic CThostFtdcTraderApiVTable_SubscribePrivateTopic;

    SubscribePublicTopic CThostFtdcTraderApiVTable_SubscribePublicTopic;

    SubscribeFlowCtrlWarning CThostFtdcTraderApiVTable_SubscribeFlowCtrlWarning;

    UnSubscribeFlowCtrlWarning CThostFtdcTraderApiVTable_UnSubscribeFlowCtrlWarning;

    ReqAuthenticate CThostFtdcTraderApiVTable_ReqAuthenticate;

    ReqUserLogin CThostFtdcTraderApiVTable_ReqUserLogin;

    ReqUserLoginEncrypt CThostFtdcTraderApiVTable_ReqUserLoginEncrypt;

    ReqUserLogout CThostFtdcTraderApiVTable_ReqUserLogout;

    ReqOrderInsert CThostFtdcTraderApiVTable_ReqOrderInsert;

    ReqOrderAction CThostFtdcTraderApiVTable_ReqOrderAction;

    ReqMKBatchOrderAction CThostFtdcTraderApiVTable_ReqMKBatchOrderAction;

    ReqExecOrderInsert CThostFtdcTraderApiVTable_ReqExecOrderInsert;

    ReqExecOrderAction CThostFtdcTraderApiVTable_ReqExecOrderAction;

    ReqForQuoteInsert CThostFtdcTraderApiVTable_ReqForQuoteInsert;

    ReqQuoteInsert CThostFtdcTraderApiVTable_ReqQuoteInsert;

    ReqQuoteAction CThostFtdcTraderApiVTable_ReqQuoteAction;

    ReqBatchOrderAction CThostFtdcTraderApiVTable_ReqBatchOrderAction;

    ReqOptionSelfCloseInsert CThostFtdcTraderApiVTable_ReqOptionSelfCloseInsert;

    ReqOptionSelfCloseAction CThostFtdcTraderApiVTable_ReqOptionSelfCloseAction;

    ReqCombActionInsert CThostFtdcTraderApiVTable_ReqCombActionInsert;

    ReqSubscribeFundChange CThostFtdcTraderApiVTable_ReqSubscribeFundChange;

    ReqUnSubscribeFundChange CThostFtdcTraderApiVTable_ReqUnSubscribeFundChange;

    ReqOffsetSetting CThostFtdcTraderApiVTable_ReqOffsetSetting;

    ReqCancelOffsetSetting CThostFtdcTraderApiVTable_ReqCancelOffsetSetting;

    ReqQryOrder CThostFtdcTraderApiVTable_ReqQryOrder;

    ReqQryTrade CThostFtdcTraderApiVTable_ReqQryTrade;

    ReqQryInvestorPosition CThostFtdcTraderApiVTable_ReqQryInvestorPosition;

    ReqQryTradingAccount CThostFtdcTraderApiVTable_ReqQryTradingAccount;

    ReqQryInvestor CThostFtdcTraderApiVTable_ReqQryInvestor;

    ReqQryTradingCode CThostFtdcTraderApiVTable_ReqQryTradingCode;

    ReqQryInstrumentMarginRate CThostFtdcTraderApiVTable_ReqQryInstrumentMarginRate;

    ReqQryInstrumentCommissionRate CThostFtdcTraderApiVTable_ReqQryInstrumentCommissionRate;

    ReqQryExchange CThostFtdcTraderApiVTable_ReqQryExchange;

    ReqQryProduct CThostFtdcTraderApiVTable_ReqQryProduct;

    ReqQryInstrument CThostFtdcTraderApiVTable_ReqQryInstrument;

    ReqQryCombInstrument CThostFtdcTraderApiVTable_ReqQryCombInstrument;

    ReqQryRCAMSInvestorProdMargin CThostFtdcTraderApiVTable_ReqQryRCAMSInvestorProdMargin;

    ReqQryRCAMSInvestorCombPosition CThostFtdcTraderApiVTable_ReqQryRCAMSInvestorCombPosition;

    ReqQryInvestorPositionForComb CThostFtdcTraderApiVTable_ReqQryInvestorPositionForComb;

    ReqQryCombAction CThostFtdcTraderApiVTable_ReqQryCombAction;

    ReqQryDepthMarketData CThostFtdcTraderApiVTable_ReqQryDepthMarketData;

    ReqQryOptionSelfClose CThostFtdcTraderApiVTable_ReqQryOptionSelfClose;

    ReqQryInstrumentStatus CThostFtdcTraderApiVTable_ReqQryInstrumentStatus;

    ReqQryInvestorPositionDetail CThostFtdcTraderApiVTable_ReqQryInvestorPositionDetail;

    ReqQryExchangeMarginRate CThostFtdcTraderApiVTable_ReqQryExchangeMarginRate;

    ReqQryExchangeMarginRateAdjust CThostFtdcTraderApiVTable_ReqQryExchangeMarginRateAdjust;

    ReqQryOptionInstrTradeCost CThostFtdcTraderApiVTable_ReqQryOptionInstrTradeCost;

    ReqQryOptionInstrCommRate CThostFtdcTraderApiVTable_ReqQryOptionInstrCommRate;

    ReqQryExecOrder CThostFtdcTraderApiVTable_ReqQryExecOrder;

    ReqQryForQuote CThostFtdcTraderApiVTable_ReqQryForQuote;

    ReqQryQuote CThostFtdcTraderApiVTable_ReqQryQuote;

    ReqQryInstrumentOrderCommRate CThostFtdcTraderApiVTable_ReqQryInstrumentOrderCommRate;

    ReqQryForQuoteParam CThostFtdcTraderApiVTable_ReqQryForQuoteParam;

    ReqQryTraderOffer CThostFtdcTraderApiVTable_ReqQryTraderOffer;

    ReqQryInvestorProdSPBMDetail CThostFtdcTraderApiVTable_ReqQryInvestorProdSPBMDetail;

    ReqQrySPMMInvestorCommodityGroupMargin CThostFtdcTraderApiVTable_ReqQrySPMMInvestorCommodityGroupMargin;

    ReqQryRULEInvestorProdMargin CThostFtdcTraderApiVTable_ReqQryRULEInvestorProdMargin;

    ReqQryControlParam CThostFtdcTraderApiVTable_ReqQryControlParam;

    ReqQryOffsetSetting CThostFtdcTraderApiVTable_ReqQryOffsetSetting;

    ReqUserPasswordUpdate CThostFtdcTraderApiVTable_ReqUserPasswordUpdate;

} CThostFtdcTraderApiVTable;

// cpp api instance C wrapper
typedef struct
{
    CThostFtdcTraderApiVTable *vtable;
} CThostFtdcTraderApiExt;

// static functions to create cpp api instance & get version

typedef void* (*CreateFtdcTraderApi)(char* pszFlowPath);
void* CallCreateFtdcTraderApi(CreateFtdcTraderApi fn, char* pszFlowPath);

typedef char* (*GetApiVersion)();
char* CallGetApiVersion(GetApiVersion fn);


// helper function call for cgo
void CallRelease(Release fn, void *this);

void CallInit(Init fn, void *this, bool bContinuous);

int CallJoin(Join fn, void *this);

char* CallGetTradingDay(GetTradingDay fn, void *this);

void CallRegisterFront(RegisterFront fn, void *this, char* pszFrontAddress);

void CallRegisterSpi(RegisterSpi fn, void *this, void *pSpi);

void CallSubscribePrivateTopic(SubscribePrivateTopic fn, void *this, int nResumeType);

void CallSubscribePublicTopic(SubscribePublicTopic fn, void *this, int nResumeType);

int CallSubscribeFlowCtrlWarning(SubscribeFlowCtrlWarning fn, void *this, char* ppTraderID[], int nCount);

int CallUnSubscribeFlowCtrlWarning(UnSubscribeFlowCtrlWarning fn, void *this, char* ppTraderID[], int nCount);

int CallReqAuthenticate(ReqAuthenticate fn, void *this, struct CThostFtdcReqAuthenticateField* pReqAuthenticateField, int nRequestID);

int CallReqUserLogin(ReqUserLogin fn, void *this, struct CThostFtdcReqUserLoginField* pReqUserLoginField, int nRequestID);

int CallReqUserLoginEncrypt(ReqUserLoginEncrypt fn, void *this, struct CThostFtdcReqUserLoginField* pReqUserLoginField, int nRequestID);

int CallReqUserLogout(ReqUserLogout fn, void *this, struct CThostFtdcUserLogoutField* pUserLogout, int nRequestID);

int CallReqOrderInsert(ReqOrderInsert fn, void *this, struct CThostFtdcInputOrderField* pInputOrder, int nRequestID);

int CallReqOrderAction(ReqOrderAction fn, void *this, struct CThostFtdcInputOrderActionField* pInputOrderAction, int nRequestID);

int CallReqMKBatchOrderAction(ReqMKBatchOrderAction fn, void *this, struct CThostFtdcMKInputOrderActionField* pMKInputOrderAction, int nRequestID);

int CallReqExecOrderInsert(ReqExecOrderInsert fn, void *this, struct CThostFtdcInputExecOrderField* pInputExecOrder, int nRequestID);

int CallReqExecOrderAction(ReqExecOrderAction fn, void *this, struct CThostFtdcInputExecOrderActionField* pInputExecOrderAction, int nRequestID);

int CallReqForQuoteInsert(ReqForQuoteInsert fn, void *this, struct CThostFtdcInputForQuoteField* pInputForQuote, int nRequestID);

int CallReqQuoteInsert(ReqQuoteInsert fn, void *this, struct CThostFtdcInputQuoteField* pInputQuote, int nRequestID);

int CallReqQuoteAction(ReqQuoteAction fn, void *this, struct CThostFtdcInputQuoteActionField* pInputQuoteAction, int nRequestID);

int CallReqBatchOrderAction(ReqBatchOrderAction fn, void *this, struct CThostFtdcInputBatchOrderActionField* pInputBatchOrderAction, int nRequestID);

int CallReqOptionSelfCloseInsert(ReqOptionSelfCloseInsert fn, void *this, struct CThostFtdcInputOptionSelfCloseField* pInputOptionSelfClose, int nRequestID);

int CallReqOptionSelfCloseAction(ReqOptionSelfCloseAction fn, void *this, struct CThostFtdcInputOptionSelfCloseActionField* pInputOptionSelfCloseAction, int nRequestID);

int CallReqCombActionInsert(ReqCombActionInsert fn, void *this, struct CThostFtdcInputCombActionField* pInputCombAction, int nRequestID);

int CallReqSubscribeFundChange(ReqSubscribeFundChange fn, void *this, int nRequestID);

int CallReqUnSubscribeFundChange(ReqUnSubscribeFundChange fn, void *this, int nRequestID);

int CallReqOffsetSetting(ReqOffsetSetting fn, void *this, struct CThostFtdcInputOffsetSettingField* pInputOffsetSetting, int nRequestID);

int CallReqCancelOffsetSetting(ReqCancelOffsetSetting fn, void *this, struct CThostFtdcInputOffsetSettingField* pInputOffsetSetting, int nRequestID);

int CallReqQryOrder(ReqQryOrder fn, void *this, struct CThostFtdcQryOrderField* pQryOrder, int nRequestID);

int CallReqQryTrade(ReqQryTrade fn, void *this, struct CThostFtdcQryTradeField* pQryTrade, int nRequestID);

int CallReqQryInvestorPosition(ReqQryInvestorPosition fn, void *this, struct CThostFtdcQryInvestorPositionField* pQryInvestorPosition, int nRequestID);

int CallReqQryTradingAccount(ReqQryTradingAccount fn, void *this, struct CThostFtdcQryTradingAccountField* pQryTradingAccount, int nRequestID);

int CallReqQryInvestor(ReqQryInvestor fn, void *this, struct CThostFtdcQryInvestorField* pQryInvestor, int nRequestID);

int CallReqQryTradingCode(ReqQryTradingCode fn, void *this, struct CThostFtdcQryTradingCodeField* pQryTradingCode, int nRequestID);

int CallReqQryInstrumentMarginRate(ReqQryInstrumentMarginRate fn, void *this, struct CThostFtdcQryInstrumentMarginRateField* pQryInstrumentMarginRate, int nRequestID);

int CallReqQryInstrumentCommissionRate(ReqQryInstrumentCommissionRate fn, void *this, struct CThostFtdcQryInstrumentCommissionRateField* pQryInstrumentCommissionRate, int nRequestID);

int CallReqQryExchange(ReqQryExchange fn, void *this, struct CThostFtdcQryExchangeField* pQryExchange, int nRequestID);

int CallReqQryProduct(ReqQryProduct fn, void *this, struct CThostFtdcQryProductField* pQryProduct, int nRequestID);

int CallReqQryInstrument(ReqQryInstrument fn, void *this, struct CThostFtdcQryInstrumentField* pQryInstrument, int nRequestID);

int CallReqQryCombInstrument(ReqQryCombInstrument fn, void *this, struct CThostFtdcQryCombInstrumentField* pQryCombInstrument, int nRequestID);

int CallReqQryRCAMSInvestorProdMargin(ReqQryRCAMSInvestorProdMargin fn, void *this, struct CThostFtdcQryRCAMSInvestorProdMarginField* pQryRCAMSInvestorProdMargin, int nRequestID);

int CallReqQryRCAMSInvestorCombPosition(ReqQryRCAMSInvestorCombPosition fn, void *this, struct CThostFtdcQryRCAMSInvestorCombPositionField* pQryRCAMSInvestorCombPosition, int nRequestID);

int CallReqQryInvestorPositionForComb(ReqQryInvestorPositionForComb fn, void *this, struct CThostFtdcQryInvestorPositionForCombField* pQryIPForComb, int nRequestID);

int CallReqQryCombAction(ReqQryCombAction fn, void *this, struct CThostFtdcQryCombActionField* pQryCombAction, int nRequestID);

int CallReqQryDepthMarketData(ReqQryDepthMarketData fn, void *this, struct CThostFtdcQryDepthMarketDataField* pQryDepthMarketData, int nRequestID);

int CallReqQryOptionSelfClose(ReqQryOptionSelfClose fn, void *this, struct CThostFtdcQryOptionSelfCloseField* pQryOptionSelfClose, int nRequestID);

int CallReqQryInstrumentStatus(ReqQryInstrumentStatus fn, void *this, struct CThostFtdcQryInstrumentStatusField* pQryInstrumentStatus, int nRequestID);

int CallReqQryInvestorPositionDetail(ReqQryInvestorPositionDetail fn, void *this, struct CThostFtdcQryInvestorPositionDetailField* pQryInvestorPositionDetail, int nRequestID);

int CallReqQryExchangeMarginRate(ReqQryExchangeMarginRate fn, void *this, struct CThostFtdcQryExchangeMarginRateField* pQryExchangeMarginRate, int nRequestID);

int CallReqQryExchangeMarginRateAdjust(ReqQryExchangeMarginRateAdjust fn, void *this, struct CThostFtdcQryExchangeMarginRateAdjustField* pQryExchangeMarginRateAdjust, int nRequestID);

int CallReqQryOptionInstrTradeCost(ReqQryOptionInstrTradeCost fn, void *this, struct CThostFtdcQryOptionInstrTradeCostField* pQryOptionInstrTradeCost, int nRequestID);

int CallReqQryOptionInstrCommRate(ReqQryOptionInstrCommRate fn, void *this, struct CThostFtdcQryOptionInstrCommRateField* pQryOptionInstrCommRate, int nRequestID);

int CallReqQryExecOrder(ReqQryExecOrder fn, void *this, struct CThostFtdcQryExecOrderField* pQryExecOrder, int nRequestID);

int CallReqQryForQuote(ReqQryForQuote fn, void *this, struct CThostFtdcQryForQuoteField* pQryForQuote, int nRequestID);

int CallReqQryQuote(ReqQryQuote fn, void *this, struct CThostFtdcQryQuoteField* pQryQuote, int nRequestID);

int CallReqQryInstrumentOrderCommRate(ReqQryInstrumentOrderCommRate fn, void *this, struct CThostFtdcQryInstrumentOrderCommRateField* pQryInstrumentOrderCommRate, int nRequestID);

int CallReqQryForQuoteParam(ReqQryForQuoteParam fn, void *this, struct CThostFtdcQryForQuoteParamField* pQryForQuoteParam, int nRequestID);

int CallReqQryTraderOffer(ReqQryTraderOffer fn, void *this, struct CThostFtdcQryTraderOfferField* pQryTraderOffer, int nRequestID);

int CallReqQryInvestorProdSPBMDetail(ReqQryInvestorProdSPBMDetail fn, void *this, struct CThostFtdcQryInvestorProdSPBMDetailField* pQryInvestorProdSPBMDetail, int nRequestID);

int CallReqQrySPMMInvestorCommodityGroupMargin(ReqQrySPMMInvestorCommodityGroupMargin fn, void *this, struct CThostFtdcQrySPMMInvestorCommodityGroupMarginField* pQrySPMMInvestorCommodityGroupMargin, int nRequestID);

int CallReqQryRULEInvestorProdMargin(ReqQryRULEInvestorProdMargin fn, void *this, struct CThostFtdcQryRULEInvestorProdMarginField* pQryRULEInvestorProdMargin, int nRequestID);

int CallReqQryControlParam(ReqQryControlParam fn, void *this, struct CThostFtdcQryControlParamField* pQryControlParam, int nRequestID);

int CallReqQryOffsetSetting(ReqQryOffsetSetting fn, void *this, struct CThostFtdcQryOffsetSettingField* pQryOffsetSetting, int nRequestID);

int CallReqUserPasswordUpdate(ReqUserPasswordUpdate fn, void *this, struct CThostFtdcUserPasswordUpdateField* pUserPasswordUpdate, int nRequestID);


#ifdef __cplusplus
}
#endif

#endif