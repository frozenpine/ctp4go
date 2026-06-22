#pragma once

#ifndef TRADER_V6713_API_HELPER_H
#define TRADER_V6713_API_HELPER_H

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

#define THOST_TE_RESUME_TYPE int

    typedef void (*Release)(void *this);

    typedef void (*Init)(void *this);

    typedef int (*Join)(void *this);

    typedef const char *(*GetTradingDay)(void *this);

    typedef void (*GetFrontInfo)(void *this, struct CThostFtdcFrontInfoField *pFrontInfo);

    typedef void (*RegisterFront)(void *this, char *pszFrontAddress);

    typedef void (*RegisterNameServer)(void *this, char *pszNsAddress);

    typedef void (*RegisterFensUserInfo)(void *this, struct CThostFtdcFensUserInfoField *pFensUserInfo);

    typedef void (*RegisterSpi)(void *this, void *pSpi);

    typedef void (*SubscribePrivateTopic)(void *this, THOST_TE_RESUME_TYPE nResumeType, int nSeqNo);

    typedef void (*SubscribePublicTopic)(void *this, THOST_TE_RESUME_TYPE nResumeType);

    typedef int (*ReqAuthenticate)(void *this, struct CThostFtdcReqAuthenticateField *pReqAuthenticateField, int nRequestID);

    typedef int (*RegisterUserSystemInfo)(void *this, struct CThostFtdcUserSystemInfoField *pUserSystemInfo);

    typedef int (*SubmitUserSystemInfo)(void *this, struct CThostFtdcUserSystemInfoField *pUserSystemInfo);

    typedef int (*RegisterWechatUserSystemInfo)(void *this, struct CThostFtdcWechatUserSystemInfoField *pUserSystemInfo);

    typedef int (*SubmitWechatUserSystemInfo)(void *this, struct CThostFtdcWechatUserSystemInfoField *pUserSystemInfo);

    typedef int (*ReqUserLogin)(void *this, struct CThostFtdcReqUserLoginField *pReqUserLoginField, int nRequestID);

    typedef int (*ReqUserLogout)(void *this, struct CThostFtdcUserLogoutField *pUserLogout, int nRequestID);

    typedef int (*ReqUserPasswordUpdate)(void *this, struct CThostFtdcUserPasswordUpdateField *pUserPasswordUpdate, int nRequestID);

    typedef int (*ReqTradingAccountPasswordUpdate)(void *this, struct CThostFtdcTradingAccountPasswordUpdateField *pTradingAccountPasswordUpdate, int nRequestID);

    typedef int (*ReqUserAuthMethod)(void *this, struct CThostFtdcReqUserAuthMethodField *pReqUserAuthMethod, int nRequestID);

    typedef int (*ReqGenUserCaptcha)(void *this, struct CThostFtdcReqGenUserCaptchaField *pReqGenUserCaptcha, int nRequestID);

    typedef int (*ReqGenUserText)(void *this, struct CThostFtdcReqGenUserTextField *pReqGenUserText, int nRequestID);

    typedef int (*ReqUserLoginWithCaptcha)(void *this, struct CThostFtdcReqUserLoginWithCaptchaField *pReqUserLoginWithCaptcha, int nRequestID);

    typedef int (*ReqUserLoginWithText)(void *this, struct CThostFtdcReqUserLoginWithTextField *pReqUserLoginWithText, int nRequestID);

    typedef int (*ReqUserLoginWithOTP)(void *this, struct CThostFtdcReqUserLoginWithOTPField *pReqUserLoginWithOTP, int nRequestID);

    typedef int (*ReqOrderInsert)(void *this, struct CThostFtdcInputOrderField *pInputOrder, int nRequestID);

    typedef int (*ReqParkedOrderInsert)(void *this, struct CThostFtdcParkedOrderField *pParkedOrder, int nRequestID);

    typedef int (*ReqParkedOrderAction)(void *this, struct CThostFtdcParkedOrderActionField *pParkedOrderAction, int nRequestID);

    typedef int (*ReqOrderAction)(void *this, struct CThostFtdcInputOrderActionField *pInputOrderAction, int nRequestID);

    typedef int (*ReqQryMaxOrderVolume)(void *this, struct CThostFtdcQryMaxOrderVolumeField *pQryMaxOrderVolume, int nRequestID);

    typedef int (*ReqSettlementInfoConfirm)(void *this, struct CThostFtdcSettlementInfoConfirmField *pSettlementInfoConfirm, int nRequestID);

    typedef int (*ReqRemoveParkedOrder)(void *this, struct CThostFtdcRemoveParkedOrderField *pRemoveParkedOrder, int nRequestID);

    typedef int (*ReqRemoveParkedOrderAction)(void *this, struct CThostFtdcRemoveParkedOrderActionField *pRemoveParkedOrderAction, int nRequestID);

    typedef int (*ReqExecOrderInsert)(void *this, struct CThostFtdcInputExecOrderField *pInputExecOrder, int nRequestID);

    typedef int (*ReqExecOrderAction)(void *this, struct CThostFtdcInputExecOrderActionField *pInputExecOrderAction, int nRequestID);

    typedef int (*ReqForQuoteInsert)(void *this, struct CThostFtdcInputForQuoteField *pInputForQuote, int nRequestID);

    typedef int (*ReqQuoteInsert)(void *this, struct CThostFtdcInputQuoteField *pInputQuote, int nRequestID);

    typedef int (*ReqQuoteAction)(void *this, struct CThostFtdcInputQuoteActionField *pInputQuoteAction, int nRequestID);

    typedef int (*ReqBatchOrderAction)(void *this, struct CThostFtdcInputBatchOrderActionField *pInputBatchOrderAction, int nRequestID);

    typedef int (*ReqOptionSelfCloseInsert)(void *this, struct CThostFtdcInputOptionSelfCloseField *pInputOptionSelfClose, int nRequestID);

    typedef int (*ReqOptionSelfCloseAction)(void *this, struct CThostFtdcInputOptionSelfCloseActionField *pInputOptionSelfCloseAction, int nRequestID);

    typedef int (*ReqCombActionInsert)(void *this, struct CThostFtdcInputCombActionField *pInputCombAction, int nRequestID);

    typedef int (*ReqQryOrder)(void *this, struct CThostFtdcQryOrderField *pQryOrder, int nRequestID);

    typedef int (*ReqQryTrade)(void *this, struct CThostFtdcQryTradeField *pQryTrade, int nRequestID);

    typedef int (*ReqQryInvestorPosition)(void *this, struct CThostFtdcQryInvestorPositionField *pQryInvestorPosition, int nRequestID);

    typedef int (*ReqQryTradingAccount)(void *this, struct CThostFtdcQryTradingAccountField *pQryTradingAccount, int nRequestID);

    typedef int (*ReqQryInvestor)(void *this, struct CThostFtdcQryInvestorField *pQryInvestor, int nRequestID);

    typedef int (*ReqQryTradingCode)(void *this, struct CThostFtdcQryTradingCodeField *pQryTradingCode, int nRequestID);

    typedef int (*ReqQryInstrumentMarginRate)(void *this, struct CThostFtdcQryInstrumentMarginRateField *pQryInstrumentMarginRate, int nRequestID);

    typedef int (*ReqQryInstrumentCommissionRate)(void *this, struct CThostFtdcQryInstrumentCommissionRateField *pQryInstrumentCommissionRate, int nRequestID);

    typedef int (*ReqQryUserSession)(void *this, struct CThostFtdcQryUserSessionField *pQryUserSession, int nRequestID);

    typedef int (*ReqQryExchange)(void *this, struct CThostFtdcQryExchangeField *pQryExchange, int nRequestID);

    typedef int (*ReqQryProduct)(void *this, struct CThostFtdcQryProductField *pQryProduct, int nRequestID);

    typedef int (*ReqQryInstrument)(void *this, struct CThostFtdcQryInstrumentField *pQryInstrument, int nRequestID);

    typedef int (*ReqQryDepthMarketData)(void *this, struct CThostFtdcQryDepthMarketDataField *pQryDepthMarketData, int nRequestID);

    typedef int (*ReqQryTraderOffer)(void *this, struct CThostFtdcQryTraderOfferField *pQryTraderOffer, int nRequestID);

    typedef int (*ReqQrySettlementInfo)(void *this, struct CThostFtdcQrySettlementInfoField *pQrySettlementInfo, int nRequestID);

    typedef int (*ReqQryTransferBank)(void *this, struct CThostFtdcQryTransferBankField *pQryTransferBank, int nRequestID);

    typedef int (*ReqQryInvestorPositionDetail)(void *this, struct CThostFtdcQryInvestorPositionDetailField *pQryInvestorPositionDetail, int nRequestID);

    typedef int (*ReqQryNotice)(void *this, struct CThostFtdcQryNoticeField *pQryNotice, int nRequestID);

    typedef int (*ReqQrySettlementInfoConfirm)(void *this, struct CThostFtdcQrySettlementInfoConfirmField *pQrySettlementInfoConfirm, int nRequestID);

    typedef int (*ReqQryInvestorPositionCombineDetail)(void *this, struct CThostFtdcQryInvestorPositionCombineDetailField *pQryInvestorPositionCombineDetail, int nRequestID);

    typedef int (*ReqQryCFMMCTradingAccountKey)(void *this, struct CThostFtdcQryCFMMCTradingAccountKeyField *pQryCFMMCTradingAccountKey, int nRequestID);

    typedef int (*ReqQryEWarrantOffset)(void *this, struct CThostFtdcQryEWarrantOffsetField *pQryEWarrantOffset, int nRequestID);

    typedef int (*ReqQryInvestorProductGroupMargin)(void *this, struct CThostFtdcQryInvestorProductGroupMarginField *pQryInvestorProductGroupMargin, int nRequestID);

    typedef int (*ReqQryExchangeMarginRate)(void *this, struct CThostFtdcQryExchangeMarginRateField *pQryExchangeMarginRate, int nRequestID);

    typedef int (*ReqQryExchangeMarginRateAdjust)(void *this, struct CThostFtdcQryExchangeMarginRateAdjustField *pQryExchangeMarginRateAdjust, int nRequestID);

    typedef int (*ReqQryExchangeRate)(void *this, struct CThostFtdcQryExchangeRateField *pQryExchangeRate, int nRequestID);

    typedef int (*ReqQrySecAgentACIDMap)(void *this, struct CThostFtdcQrySecAgentACIDMapField *pQrySecAgentACIDMap, int nRequestID);

    typedef int (*ReqQryProductExchRate)(void *this, struct CThostFtdcQryProductExchRateField *pQryProductExchRate, int nRequestID);

    typedef int (*ReqQryProductGroup)(void *this, struct CThostFtdcQryProductGroupField *pQryProductGroup, int nRequestID);

    typedef int (*ReqQryMMInstrumentCommissionRate)(void *this, struct CThostFtdcQryMMInstrumentCommissionRateField *pQryMMInstrumentCommissionRate, int nRequestID);

    typedef int (*ReqQryMMOptionInstrCommRate)(void *this, struct CThostFtdcQryMMOptionInstrCommRateField *pQryMMOptionInstrCommRate, int nRequestID);

    typedef int (*ReqQryInstrumentOrderCommRate)(void *this, struct CThostFtdcQryInstrumentOrderCommRateField *pQryInstrumentOrderCommRate, int nRequestID);

    typedef int (*ReqQrySecAgentTradingAccount)(void *this, struct CThostFtdcQryTradingAccountField *pQryTradingAccount, int nRequestID);

    typedef int (*ReqQrySecAgentCheckMode)(void *this, struct CThostFtdcQrySecAgentCheckModeField *pQrySecAgentCheckMode, int nRequestID);

    typedef int (*ReqQrySecAgentTradeInfo)(void *this, struct CThostFtdcQrySecAgentTradeInfoField *pQrySecAgentTradeInfo, int nRequestID);

    typedef int (*ReqQryOptionInstrTradeCost)(void *this, struct CThostFtdcQryOptionInstrTradeCostField *pQryOptionInstrTradeCost, int nRequestID);

    typedef int (*ReqQryOptionInstrCommRate)(void *this, struct CThostFtdcQryOptionInstrCommRateField *pQryOptionInstrCommRate, int nRequestID);

    typedef int (*ReqQryExecOrder)(void *this, struct CThostFtdcQryExecOrderField *pQryExecOrder, int nRequestID);

    typedef int (*ReqQryForQuote)(void *this, struct CThostFtdcQryForQuoteField *pQryForQuote, int nRequestID);

    typedef int (*ReqQryQuote)(void *this, struct CThostFtdcQryQuoteField *pQryQuote, int nRequestID);

    typedef int (*ReqQryOptionSelfClose)(void *this, struct CThostFtdcQryOptionSelfCloseField *pQryOptionSelfClose, int nRequestID);

    typedef int (*ReqQryInvestUnit)(void *this, struct CThostFtdcQryInvestUnitField *pQryInvestUnit, int nRequestID);

    typedef int (*ReqQryCombInstrumentGuard)(void *this, struct CThostFtdcQryCombInstrumentGuardField *pQryCombInstrumentGuard, int nRequestID);

    typedef int (*ReqQryCombAction)(void *this, struct CThostFtdcQryCombActionField *pQryCombAction, int nRequestID);

    typedef int (*ReqQryTransferSerial)(void *this, struct CThostFtdcQryTransferSerialField *pQryTransferSerial, int nRequestID);

    typedef int (*ReqQryAccountregister)(void *this, struct CThostFtdcQryAccountregisterField *pQryAccountregister, int nRequestID);

    typedef int (*ReqQryContractBank)(void *this, struct CThostFtdcQryContractBankField *pQryContractBank, int nRequestID);

    typedef int (*ReqQryParkedOrder)(void *this, struct CThostFtdcQryParkedOrderField *pQryParkedOrder, int nRequestID);

    typedef int (*ReqQryParkedOrderAction)(void *this, struct CThostFtdcQryParkedOrderActionField *pQryParkedOrderAction, int nRequestID);

    typedef int (*ReqQryTradingNotice)(void *this, struct CThostFtdcQryTradingNoticeField *pQryTradingNotice, int nRequestID);

    typedef int (*ReqQryBrokerTradingParams)(void *this, struct CThostFtdcQryBrokerTradingParamsField *pQryBrokerTradingParams, int nRequestID);

    typedef int (*ReqQryBrokerTradingAlgos)(void *this, struct CThostFtdcQryBrokerTradingAlgosField *pQryBrokerTradingAlgos, int nRequestID);

    typedef int (*ReqQueryCFMMCTradingAccountToken)(void *this, struct CThostFtdcQueryCFMMCTradingAccountTokenField *pQueryCFMMCTradingAccountToken, int nRequestID);

    typedef int (*ReqFromBankToFutureByFuture)(void *this, struct CThostFtdcReqTransferField *pReqTransfer, int nRequestID);

    typedef int (*ReqFromFutureToBankByFuture)(void *this, struct CThostFtdcReqTransferField *pReqTransfer, int nRequestID);

    typedef int (*ReqQueryBankAccountMoneyByFuture)(void *this, struct CThostFtdcReqQueryAccountField *pReqQueryAccount, int nRequestID);

    typedef int (*ReqQryClassifiedInstrument)(void *this, struct CThostFtdcQryClassifiedInstrumentField *pQryClassifiedInstrument, int nRequestID);

    typedef int (*ReqQryCombPromotionParam)(void *this, struct CThostFtdcQryCombPromotionParamField *pQryCombPromotionParam, int nRequestID);

    typedef int (*ReqQryRiskSettleInvstPosition)(void *this, struct CThostFtdcQryRiskSettleInvstPositionField *pQryRiskSettleInvstPosition, int nRequestID);

    typedef int (*ReqQryRiskSettleProductStatus)(void *this, struct CThostFtdcQryRiskSettleProductStatusField *pQryRiskSettleProductStatus, int nRequestID);

    typedef int (*ReqQrySPBMFutureParameter)(void *this, struct CThostFtdcQrySPBMFutureParameterField *pQrySPBMFutureParameter, int nRequestID);

    typedef int (*ReqQrySPBMOptionParameter)(void *this, struct CThostFtdcQrySPBMOptionParameterField *pQrySPBMOptionParameter, int nRequestID);

    typedef int (*ReqQrySPBMIntraParameter)(void *this, struct CThostFtdcQrySPBMIntraParameterField *pQrySPBMIntraParameter, int nRequestID);

    typedef int (*ReqQrySPBMInterParameter)(void *this, struct CThostFtdcQrySPBMInterParameterField *pQrySPBMInterParameter, int nRequestID);

    typedef int (*ReqQrySPBMPortfDefinition)(void *this, struct CThostFtdcQrySPBMPortfDefinitionField *pQrySPBMPortfDefinition, int nRequestID);

    typedef int (*ReqQrySPBMInvestorPortfDef)(void *this, struct CThostFtdcQrySPBMInvestorPortfDefField *pQrySPBMInvestorPortfDef, int nRequestID);

    typedef int (*ReqQryInvestorPortfMarginRatio)(void *this, struct CThostFtdcQryInvestorPortfMarginRatioField *pQryInvestorPortfMarginRatio, int nRequestID);

    typedef int (*ReqQryInvestorProdSPBMDetail)(void *this, struct CThostFtdcQryInvestorProdSPBMDetailField *pQryInvestorProdSPBMDetail, int nRequestID);

    typedef int (*ReqQryInvestorCommoditySPMMMargin)(void *this, struct CThostFtdcQryInvestorCommoditySPMMMarginField *pQryInvestorCommoditySPMMMargin, int nRequestID);

    typedef int (*ReqQryInvestorCommodityGroupSPMMMargin)(void *this, struct CThostFtdcQryInvestorCommodityGroupSPMMMarginField *pQryInvestorCommodityGroupSPMMMargin, int nRequestID);

    typedef int (*ReqQrySPMMInstParam)(void *this, struct CThostFtdcQrySPMMInstParamField *pQrySPMMInstParam, int nRequestID);

    typedef int (*ReqQrySPMMProductParam)(void *this, struct CThostFtdcQrySPMMProductParamField *pQrySPMMProductParam, int nRequestID);

    typedef int (*ReqQrySPBMAddOnInterParameter)(void *this, struct CThostFtdcQrySPBMAddOnInterParameterField *pQrySPBMAddOnInterParameter, int nRequestID);

    typedef int (*ReqQryRCAMSCombProductInfo)(void *this, struct CThostFtdcQryRCAMSCombProductInfoField *pQryRCAMSCombProductInfo, int nRequestID);

    typedef int (*ReqQryRCAMSInstrParameter)(void *this, struct CThostFtdcQryRCAMSInstrParameterField *pQryRCAMSInstrParameter, int nRequestID);

    typedef int (*ReqQryRCAMSIntraParameter)(void *this, struct CThostFtdcQryRCAMSIntraParameterField *pQryRCAMSIntraParameter, int nRequestID);

    typedef int (*ReqQryRCAMSInterParameter)(void *this, struct CThostFtdcQryRCAMSInterParameterField *pQryRCAMSInterParameter, int nRequestID);

    typedef int (*ReqQryRCAMSShortOptAdjustParam)(void *this, struct CThostFtdcQryRCAMSShortOptAdjustParamField *pQryRCAMSShortOptAdjustParam, int nRequestID);

    typedef int (*ReqQryRCAMSInvestorCombPosition)(void *this, struct CThostFtdcQryRCAMSInvestorCombPositionField *pQryRCAMSInvestorCombPosition, int nRequestID);

    typedef int (*ReqQryInvestorProdRCAMSMargin)(void *this, struct CThostFtdcQryInvestorProdRCAMSMarginField *pQryInvestorProdRCAMSMargin, int nRequestID);

    typedef int (*ReqQryRULEInstrParameter)(void *this, struct CThostFtdcQryRULEInstrParameterField *pQryRULEInstrParameter, int nRequestID);

    typedef int (*ReqQryRULEIntraParameter)(void *this, struct CThostFtdcQryRULEIntraParameterField *pQryRULEIntraParameter, int nRequestID);

    typedef int (*ReqQryRULEInterParameter)(void *this, struct CThostFtdcQryRULEInterParameterField *pQryRULEInterParameter, int nRequestID);

    typedef int (*ReqQryInvestorProdRULEMargin)(void *this, struct CThostFtdcQryInvestorProdRULEMarginField *pQryInvestorProdRULEMargin, int nRequestID);

    typedef int (*ReqQryInvestorPortfSetting)(void *this, struct CThostFtdcQryInvestorPortfSettingField *pQryInvestorPortfSetting, int nRequestID);

    typedef int (*ReqQryInvestorInfoCommRec)(void *this, struct CThostFtdcQryInvestorInfoCommRecField *pQryInvestorInfoCommRec, int nRequestID);

    typedef int (*ReqQryCombLeg)(void *this, struct CThostFtdcQryCombLegField *pQryCombLeg, int nRequestID);

    typedef int (*ReqOffsetSetting)(void *this, struct CThostFtdcInputOffsetSettingField *pInputOffsetSetting, int nRequestID);

    typedef int (*ReqCancelOffsetSetting)(void *this, struct CThostFtdcInputOffsetSettingField *pInputOffsetSetting, int nRequestID);

    typedef int (*ReqQryOffsetSetting)(void *this, struct CThostFtdcQryOffsetSettingField *pQryOffsetSetting, int nRequestID);

    typedef int (*ReqGenSMSCode)(void *this, struct CThostFtdcReqGenSMSCodeField *pReqGenSMSCode, int nRequestID);

    typedef int (*ReqSpdApply)(void *this, struct CThostFtdcInputSpdApplyField *pInputSpdApply, int nRequestID);

    typedef int (*ReqSpdApplyAction)(void *this, struct CThostFtdcInputSpdApplyActionField *pInputSpdApplyAction, int nRequestID);

    typedef int (*ReqQrySpdApply)(void *this, struct CThostFtdcQrySpdApplyField *pQrySpdApply, int nRequestID);

    typedef int (*ReqHedgeCfm)(void *this, struct CThostFtdcInputHedgeCfmField *pInputHedgeCfm, int nRequestID);

    typedef int (*ReqHedgeCfmAction)(void *this, struct CThostFtdcInputHedgeCfmActionField *pInputHedgeCfmAction, int nRequestID);

    typedef int (*ReqQryHedgeCfm)(void *this, struct CThostFtdcQryHedgeCfmField *pQryHedgeCfm, int nRequestID);

    typedef struct
    {
        Release CThostFtdcTraderApiVTable_Release;
        Init CThostFtdcTraderApiVTable_Init;
        Join CThostFtdcTraderApiVTable_Join;
        GetTradingDay CThostFtdcTraderApiVTable_GetTradingDay;
        GetFrontInfo CThostFtdcTraderApiVTable_GetFrontInfo;
        RegisterFront CThostFtdcTraderApiVTable_RegisterFront;
        RegisterNameServer CThostFtdcTraderApiVTable_RegisterNameServer;
        RegisterFensUserInfo CThostFtdcTraderApiVTable_RegisterFensUserInfo;
        RegisterSpi CThostFtdcTraderApiVTable_RegisterSpi;
        SubscribePrivateTopic CThostFtdcTraderApiVTable_SubscribePrivateTopic;
        SubscribePublicTopic CThostFtdcTraderApiVTable_SubscribePublicTopic;
        ReqAuthenticate CThostFtdcTraderApiVTable_ReqAuthenticate;
        RegisterUserSystemInfo CThostFtdcTraderApiVTable_RegisterUserSystemInfo;
        SubmitUserSystemInfo CThostFtdcTraderApiVTable_SubmitUserSystemInfo;
        RegisterWechatUserSystemInfo CThostFtdcTraderApiVTable_RegisterWechatUserSystemInfo;
        SubmitWechatUserSystemInfo CThostFtdcTraderApiVTable_SubmitWechatUserSystemInfo;
        ReqUserLogin CThostFtdcTraderApiVTable_ReqUserLogin;
        ReqUserLogout CThostFtdcTraderApiVTable_ReqUserLogout;
        ReqUserPasswordUpdate CThostFtdcTraderApiVTable_ReqUserPasswordUpdate;
        ReqTradingAccountPasswordUpdate CThostFtdcTraderApiVTable_ReqTradingAccountPasswordUpdate;
        ReqUserAuthMethod CThostFtdcTraderApiVTable_ReqUserAuthMethod;
        ReqGenUserCaptcha CThostFtdcTraderApiVTable_ReqGenUserCaptcha;
        ReqGenUserText CThostFtdcTraderApiVTable_ReqGenUserText;
        ReqUserLoginWithCaptcha CThostFtdcTraderApiVTable_ReqUserLoginWithCaptcha;
        ReqUserLoginWithText CThostFtdcTraderApiVTable_ReqUserLoginWithText;
        ReqUserLoginWithOTP CThostFtdcTraderApiVTable_ReqUserLoginWithOTP;
        ReqOrderInsert CThostFtdcTraderApiVTable_ReqOrderInsert;
        ReqParkedOrderInsert CThostFtdcTraderApiVTable_ReqParkedOrderInsert;
        ReqParkedOrderAction CThostFtdcTraderApiVTable_ReqParkedOrderAction;
        ReqOrderAction CThostFtdcTraderApiVTable_ReqOrderAction;
        ReqQryMaxOrderVolume CThostFtdcTraderApiVTable_ReqQryMaxOrderVolume;
        ReqSettlementInfoConfirm CThostFtdcTraderApiVTable_ReqSettlementInfoConfirm;
        ReqRemoveParkedOrder CThostFtdcTraderApiVTable_ReqRemoveParkedOrder;
        ReqRemoveParkedOrderAction CThostFtdcTraderApiVTable_ReqRemoveParkedOrderAction;
        ReqExecOrderInsert CThostFtdcTraderApiVTable_ReqExecOrderInsert;
        ReqExecOrderAction CThostFtdcTraderApiVTable_ReqExecOrderAction;
        ReqForQuoteInsert CThostFtdcTraderApiVTable_ReqForQuoteInsert;
        ReqQuoteInsert CThostFtdcTraderApiVTable_ReqQuoteInsert;
        ReqQuoteAction CThostFtdcTraderApiVTable_ReqQuoteAction;
        ReqBatchOrderAction CThostFtdcTraderApiVTable_ReqBatchOrderAction;
        ReqOptionSelfCloseInsert CThostFtdcTraderApiVTable_ReqOptionSelfCloseInsert;
        ReqOptionSelfCloseAction CThostFtdcTraderApiVTable_ReqOptionSelfCloseAction;
        ReqCombActionInsert CThostFtdcTraderApiVTable_ReqCombActionInsert;
        ReqQryOrder CThostFtdcTraderApiVTable_ReqQryOrder;
        ReqQryTrade CThostFtdcTraderApiVTable_ReqQryTrade;
        ReqQryInvestorPosition CThostFtdcTraderApiVTable_ReqQryInvestorPosition;
        ReqQryTradingAccount CThostFtdcTraderApiVTable_ReqQryTradingAccount;
        ReqQryInvestor CThostFtdcTraderApiVTable_ReqQryInvestor;
        ReqQryTradingCode CThostFtdcTraderApiVTable_ReqQryTradingCode;
        ReqQryInstrumentMarginRate CThostFtdcTraderApiVTable_ReqQryInstrumentMarginRate;
        ReqQryInstrumentCommissionRate CThostFtdcTraderApiVTable_ReqQryInstrumentCommissionRate;
        ReqQryUserSession CThostFtdcTraderApiVTable_ReqQryUserSession;
        ReqQryExchange CThostFtdcTraderApiVTable_ReqQryExchange;
        ReqQryProduct CThostFtdcTraderApiVTable_ReqQryProduct;
        ReqQryInstrument CThostFtdcTraderApiVTable_ReqQryInstrument;
        ReqQryDepthMarketData CThostFtdcTraderApiVTable_ReqQryDepthMarketData;
        ReqQryTraderOffer CThostFtdcTraderApiVTable_ReqQryTraderOffer;
        ReqQrySettlementInfo CThostFtdcTraderApiVTable_ReqQrySettlementInfo;
        ReqQryTransferBank CThostFtdcTraderApiVTable_ReqQryTransferBank;
        ReqQryInvestorPositionDetail CThostFtdcTraderApiVTable_ReqQryInvestorPositionDetail;
        ReqQryNotice CThostFtdcTraderApiVTable_ReqQryNotice;
        ReqQrySettlementInfoConfirm CThostFtdcTraderApiVTable_ReqQrySettlementInfoConfirm;
        ReqQryInvestorPositionCombineDetail CThostFtdcTraderApiVTable_ReqQryInvestorPositionCombineDetail;
        ReqQryCFMMCTradingAccountKey CThostFtdcTraderApiVTable_ReqQryCFMMCTradingAccountKey;
        ReqQryEWarrantOffset CThostFtdcTraderApiVTable_ReqQryEWarrantOffset;
        ReqQryInvestorProductGroupMargin CThostFtdcTraderApiVTable_ReqQryInvestorProductGroupMargin;
        ReqQryExchangeMarginRate CThostFtdcTraderApiVTable_ReqQryExchangeMarginRate;
        ReqQryExchangeMarginRateAdjust CThostFtdcTraderApiVTable_ReqQryExchangeMarginRateAdjust;
        ReqQryExchangeRate CThostFtdcTraderApiVTable_ReqQryExchangeRate;
        ReqQrySecAgentACIDMap CThostFtdcTraderApiVTable_ReqQrySecAgentACIDMap;
        ReqQryProductExchRate CThostFtdcTraderApiVTable_ReqQryProductExchRate;
        ReqQryProductGroup CThostFtdcTraderApiVTable_ReqQryProductGroup;
        ReqQryMMInstrumentCommissionRate CThostFtdcTraderApiVTable_ReqQryMMInstrumentCommissionRate;
        ReqQryMMOptionInstrCommRate CThostFtdcTraderApiVTable_ReqQryMMOptionInstrCommRate;
        ReqQryInstrumentOrderCommRate CThostFtdcTraderApiVTable_ReqQryInstrumentOrderCommRate;
        ReqQrySecAgentTradingAccount CThostFtdcTraderApiVTable_ReqQrySecAgentTradingAccount;
        ReqQrySecAgentCheckMode CThostFtdcTraderApiVTable_ReqQrySecAgentCheckMode;
        ReqQrySecAgentTradeInfo CThostFtdcTraderApiVTable_ReqQrySecAgentTradeInfo;
        ReqQryOptionInstrTradeCost CThostFtdcTraderApiVTable_ReqQryOptionInstrTradeCost;
        ReqQryOptionInstrCommRate CThostFtdcTraderApiVTable_ReqQryOptionInstrCommRate;
        ReqQryExecOrder CThostFtdcTraderApiVTable_ReqQryExecOrder;
        ReqQryForQuote CThostFtdcTraderApiVTable_ReqQryForQuote;
        ReqQryQuote CThostFtdcTraderApiVTable_ReqQryQuote;
        ReqQryOptionSelfClose CThostFtdcTraderApiVTable_ReqQryOptionSelfClose;
        ReqQryInvestUnit CThostFtdcTraderApiVTable_ReqQryInvestUnit;
        ReqQryCombInstrumentGuard CThostFtdcTraderApiVTable_ReqQryCombInstrumentGuard;
        ReqQryCombAction CThostFtdcTraderApiVTable_ReqQryCombAction;
        ReqQryTransferSerial CThostFtdcTraderApiVTable_ReqQryTransferSerial;
        ReqQryAccountregister CThostFtdcTraderApiVTable_ReqQryAccountregister;
        ReqQryContractBank CThostFtdcTraderApiVTable_ReqQryContractBank;
        ReqQryParkedOrder CThostFtdcTraderApiVTable_ReqQryParkedOrder;
        ReqQryParkedOrderAction CThostFtdcTraderApiVTable_ReqQryParkedOrderAction;
        ReqQryTradingNotice CThostFtdcTraderApiVTable_ReqQryTradingNotice;
        ReqQryBrokerTradingParams CThostFtdcTraderApiVTable_ReqQryBrokerTradingParams;
        ReqQryBrokerTradingAlgos CThostFtdcTraderApiVTable_ReqQryBrokerTradingAlgos;
        ReqQueryCFMMCTradingAccountToken CThostFtdcTraderApiVTable_ReqQueryCFMMCTradingAccountToken;
        ReqFromBankToFutureByFuture CThostFtdcTraderApiVTable_ReqFromBankToFutureByFuture;
        ReqFromFutureToBankByFuture CThostFtdcTraderApiVTable_ReqFromFutureToBankByFuture;
        ReqQueryBankAccountMoneyByFuture CThostFtdcTraderApiVTable_ReqQueryBankAccountMoneyByFuture;
        ReqQryClassifiedInstrument CThostFtdcTraderApiVTable_ReqQryClassifiedInstrument;
        ReqQryCombPromotionParam CThostFtdcTraderApiVTable_ReqQryCombPromotionParam;
        ReqQryRiskSettleInvstPosition CThostFtdcTraderApiVTable_ReqQryRiskSettleInvstPosition;
        ReqQryRiskSettleProductStatus CThostFtdcTraderApiVTable_ReqQryRiskSettleProductStatus;
        ReqQrySPBMFutureParameter CThostFtdcTraderApiVTable_ReqQrySPBMFutureParameter;
        ReqQrySPBMOptionParameter CThostFtdcTraderApiVTable_ReqQrySPBMOptionParameter;
        ReqQrySPBMIntraParameter CThostFtdcTraderApiVTable_ReqQrySPBMIntraParameter;
        ReqQrySPBMInterParameter CThostFtdcTraderApiVTable_ReqQrySPBMInterParameter;
        ReqQrySPBMPortfDefinition CThostFtdcTraderApiVTable_ReqQrySPBMPortfDefinition;
        ReqQrySPBMInvestorPortfDef CThostFtdcTraderApiVTable_ReqQrySPBMInvestorPortfDef;
        ReqQryInvestorPortfMarginRatio CThostFtdcTraderApiVTable_ReqQryInvestorPortfMarginRatio;
        ReqQryInvestorProdSPBMDetail CThostFtdcTraderApiVTable_ReqQryInvestorProdSPBMDetail;
        ReqQryInvestorCommoditySPMMMargin CThostFtdcTraderApiVTable_ReqQryInvestorCommoditySPMMMargin;
        ReqQryInvestorCommodityGroupSPMMMargin CThostFtdcTraderApiVTable_ReqQryInvestorCommodityGroupSPMMMargin;
        ReqQrySPMMInstParam CThostFtdcTraderApiVTable_ReqQrySPMMInstParam;
        ReqQrySPMMProductParam CThostFtdcTraderApiVTable_ReqQrySPMMProductParam;
        ReqQrySPBMAddOnInterParameter CThostFtdcTraderApiVTable_ReqQrySPBMAddOnInterParameter;
        ReqQryRCAMSCombProductInfo CThostFtdcTraderApiVTable_ReqQryRCAMSCombProductInfo;
        ReqQryRCAMSInstrParameter CThostFtdcTraderApiVTable_ReqQryRCAMSInstrParameter;
        ReqQryRCAMSIntraParameter CThostFtdcTraderApiVTable_ReqQryRCAMSIntraParameter;
        ReqQryRCAMSInterParameter CThostFtdcTraderApiVTable_ReqQryRCAMSInterParameter;
        ReqQryRCAMSShortOptAdjustParam CThostFtdcTraderApiVTable_ReqQryRCAMSShortOptAdjustParam;
        ReqQryRCAMSInvestorCombPosition CThostFtdcTraderApiVTable_ReqQryRCAMSInvestorCombPosition;
        ReqQryInvestorProdRCAMSMargin CThostFtdcTraderApiVTable_ReqQryInvestorProdRCAMSMargin;
        ReqQryRULEInstrParameter CThostFtdcTraderApiVTable_ReqQryRULEInstrParameter;
        ReqQryRULEIntraParameter CThostFtdcTraderApiVTable_ReqQryRULEIntraParameter;
        ReqQryRULEInterParameter CThostFtdcTraderApiVTable_ReqQryRULEInterParameter;
        ReqQryInvestorProdRULEMargin CThostFtdcTraderApiVTable_ReqQryInvestorProdRULEMargin;
        ReqQryInvestorPortfSetting CThostFtdcTraderApiVTable_ReqQryInvestorPortfSetting;
        ReqQryInvestorInfoCommRec CThostFtdcTraderApiVTable_ReqQryInvestorInfoCommRec;
        ReqQryCombLeg CThostFtdcTraderApiVTable_ReqQryCombLeg;
        ReqOffsetSetting CThostFtdcTraderApiVTable_ReqOffsetSetting;
        ReqCancelOffsetSetting CThostFtdcTraderApiVTable_ReqCancelOffsetSetting;
        ReqQryOffsetSetting CThostFtdcTraderApiVTable_ReqQryOffsetSetting;
        ReqGenSMSCode CThostFtdcTraderApiVTable_ReqGenSMSCode;
        ReqSpdApply CThostFtdcTraderApiVTable_ReqSpdApply;
        ReqSpdApplyAction CThostFtdcTraderApiVTable_ReqSpdApplyAction;
        ReqQrySpdApply CThostFtdcTraderApiVTable_ReqQrySpdApply;
        ReqHedgeCfm CThostFtdcTraderApiVTable_ReqHedgeCfm;
        ReqHedgeCfmAction CThostFtdcTraderApiVTable_ReqHedgeCfmAction;
        ReqQryHedgeCfm CThostFtdcTraderApiVTable_ReqQryHedgeCfm;
    } CThostFtdcTraderApiVTable;

    typedef struct
    {
        CThostFtdcTraderApiVTable *vtable;
    } CThostFtdcTraderApiExt;

    typedef void *(*CreateFtdcTraderApi)(const char *pszFlowPath, bool bIsProductionMode);

    typedef const char *(*GetApiVersion)();

    void *CallCreateFtdcTraderApi(CreateFtdcTraderApi fn, const char *pszFlowPath, bool bIsProductionMode);

    const char *CallGetApiVersion(GetApiVersion fn);

    void CallRelease(Release fn, void *this);

    void CallInit(Init fn, void *this);

    int CallJoin(Join fn, void *this);

    const char *CallGetTradingDay(GetTradingDay fn, void *this);

    void CallGetFrontInfo(GetFrontInfo fn, void *this, struct CThostFtdcFrontInfoField *pFrontInfo);

    void CallRegisterFront(RegisterFront fn, void *this, char *pszFrontAddress);

    void CallRegisterNameServer(RegisterNameServer fn, void *this, char *pszNsAddress);

    void CallRegisterFensUserInfo(RegisterFensUserInfo fn, void *this, struct CThostFtdcFensUserInfoField *pFensUserInfo);

    void CallRegisterSpi(RegisterSpi fn, void *this, void *pSpi);

    void CallSubscribePrivateTopic(SubscribePrivateTopic fn, void *this, THOST_TE_RESUME_TYPE nResumeType, int nSeqNo);

    void CallSubscribePublicTopic(SubscribePublicTopic fn, void *this, THOST_TE_RESUME_TYPE nResumeType);

    int CallReqAuthenticate(ReqAuthenticate fn, void *this, struct CThostFtdcReqAuthenticateField *pReqAuthenticateField, int nRequestID);

    int CallRegisterUserSystemInfo(RegisterUserSystemInfo fn, void *this, struct CThostFtdcUserSystemInfoField *pUserSystemInfo);

    int CallSubmitUserSystemInfo(SubmitUserSystemInfo fn, void *this, struct CThostFtdcUserSystemInfoField *pUserSystemInfo);

    int CallRegisterWechatUserSystemInfo(RegisterWechatUserSystemInfo fn, void *this, struct CThostFtdcWechatUserSystemInfoField *pUserSystemInfo);

    int CallSubmitWechatUserSystemInfo(SubmitWechatUserSystemInfo fn, void *this, struct CThostFtdcWechatUserSystemInfoField *pUserSystemInfo);

    int CallReqUserLogin(ReqUserLogin fn, void *this, struct CThostFtdcReqUserLoginField *pReqUserLoginField, int nRequestID);

    int CallReqUserLogout(ReqUserLogout fn, void *this, struct CThostFtdcUserLogoutField *pUserLogout, int nRequestID);

    int CallReqUserPasswordUpdate(ReqUserPasswordUpdate fn, void *this, struct CThostFtdcUserPasswordUpdateField *pUserPasswordUpdate, int nRequestID);

    int CallReqTradingAccountPasswordUpdate(ReqTradingAccountPasswordUpdate fn, void *this, struct CThostFtdcTradingAccountPasswordUpdateField *pTradingAccountPasswordUpdate, int nRequestID);

    int CallReqUserAuthMethod(ReqUserAuthMethod fn, void *this, struct CThostFtdcReqUserAuthMethodField *pReqUserAuthMethod, int nRequestID);

    int CallReqGenUserCaptcha(ReqGenUserCaptcha fn, void *this, struct CThostFtdcReqGenUserCaptchaField *pReqGenUserCaptcha, int nRequestID);

    int CallReqGenUserText(ReqGenUserText fn, void *this, struct CThostFtdcReqGenUserTextField *pReqGenUserText, int nRequestID);

    int CallReqUserLoginWithCaptcha(ReqUserLoginWithCaptcha fn, void *this, struct CThostFtdcReqUserLoginWithCaptchaField *pReqUserLoginWithCaptcha, int nRequestID);

    int CallReqUserLoginWithText(ReqUserLoginWithText fn, void *this, struct CThostFtdcReqUserLoginWithTextField *pReqUserLoginWithText, int nRequestID);

    int CallReqUserLoginWithOTP(ReqUserLoginWithOTP fn, void *this, struct CThostFtdcReqUserLoginWithOTPField *pReqUserLoginWithOTP, int nRequestID);

    int CallReqOrderInsert(ReqOrderInsert fn, void *this, struct CThostFtdcInputOrderField *pInputOrder, int nRequestID);

    int CallReqParkedOrderInsert(ReqParkedOrderInsert fn, void *this, struct CThostFtdcParkedOrderField *pParkedOrder, int nRequestID);

    int CallReqParkedOrderAction(ReqParkedOrderAction fn, void *this, struct CThostFtdcParkedOrderActionField *pParkedOrderAction, int nRequestID);

    int CallReqOrderAction(ReqOrderAction fn, void *this, struct CThostFtdcInputOrderActionField *pInputOrderAction, int nRequestID);

    int CallReqQryMaxOrderVolume(ReqQryMaxOrderVolume fn, void *this, struct CThostFtdcQryMaxOrderVolumeField *pQryMaxOrderVolume, int nRequestID);

    int CallReqSettlementInfoConfirm(ReqSettlementInfoConfirm fn, void *this, struct CThostFtdcSettlementInfoConfirmField *pSettlementInfoConfirm, int nRequestID);

    int CallReqRemoveParkedOrder(ReqRemoveParkedOrder fn, void *this, struct CThostFtdcRemoveParkedOrderField *pRemoveParkedOrder, int nRequestID);

    int CallReqRemoveParkedOrderAction(ReqRemoveParkedOrderAction fn, void *this, struct CThostFtdcRemoveParkedOrderActionField *pRemoveParkedOrderAction, int nRequestID);

    int CallReqExecOrderInsert(ReqExecOrderInsert fn, void *this, struct CThostFtdcInputExecOrderField *pInputExecOrder, int nRequestID);

    int CallReqExecOrderAction(ReqExecOrderAction fn, void *this, struct CThostFtdcInputExecOrderActionField *pInputExecOrderAction, int nRequestID);

    int CallReqForQuoteInsert(ReqForQuoteInsert fn, void *this, struct CThostFtdcInputForQuoteField *pInputForQuote, int nRequestID);

    int CallReqQuoteInsert(ReqQuoteInsert fn, void *this, struct CThostFtdcInputQuoteField *pInputQuote, int nRequestID);

    int CallReqQuoteAction(ReqQuoteAction fn, void *this, struct CThostFtdcInputQuoteActionField *pInputQuoteAction, int nRequestID);

    int CallReqBatchOrderAction(ReqBatchOrderAction fn, void *this, struct CThostFtdcInputBatchOrderActionField *pInputBatchOrderAction, int nRequestID);

    int CallReqOptionSelfCloseInsert(ReqOptionSelfCloseInsert fn, void *this, struct CThostFtdcInputOptionSelfCloseField *pInputOptionSelfClose, int nRequestID);

    int CallReqOptionSelfCloseAction(ReqOptionSelfCloseAction fn, void *this, struct CThostFtdcInputOptionSelfCloseActionField *pInputOptionSelfCloseAction, int nRequestID);

    int CallReqCombActionInsert(ReqCombActionInsert fn, void *this, struct CThostFtdcInputCombActionField *pInputCombAction, int nRequestID);

    int CallReqQryOrder(ReqQryOrder fn, void *this, struct CThostFtdcQryOrderField *pQryOrder, int nRequestID);

    int CallReqQryTrade(ReqQryTrade fn, void *this, struct CThostFtdcQryTradeField *pQryTrade, int nRequestID);

    int CallReqQryInvestorPosition(ReqQryInvestorPosition fn, void *this, struct CThostFtdcQryInvestorPositionField *pQryInvestorPosition, int nRequestID);

    int CallReqQryTradingAccount(ReqQryTradingAccount fn, void *this, struct CThostFtdcQryTradingAccountField *pQryTradingAccount, int nRequestID);

    int CallReqQryInvestor(ReqQryInvestor fn, void *this, struct CThostFtdcQryInvestorField *pQryInvestor, int nRequestID);

    int CallReqQryTradingCode(ReqQryTradingCode fn, void *this, struct CThostFtdcQryTradingCodeField *pQryTradingCode, int nRequestID);

    int CallReqQryInstrumentMarginRate(ReqQryInstrumentMarginRate fn, void *this, struct CThostFtdcQryInstrumentMarginRateField *pQryInstrumentMarginRate, int nRequestID);

    int CallReqQryInstrumentCommissionRate(ReqQryInstrumentCommissionRate fn, void *this, struct CThostFtdcQryInstrumentCommissionRateField *pQryInstrumentCommissionRate, int nRequestID);

    int CallReqQryUserSession(ReqQryUserSession fn, void *this, struct CThostFtdcQryUserSessionField *pQryUserSession, int nRequestID);

    int CallReqQryExchange(ReqQryExchange fn, void *this, struct CThostFtdcQryExchangeField *pQryExchange, int nRequestID);

    int CallReqQryProduct(ReqQryProduct fn, void *this, struct CThostFtdcQryProductField *pQryProduct, int nRequestID);

    int CallReqQryInstrument(ReqQryInstrument fn, void *this, struct CThostFtdcQryInstrumentField *pQryInstrument, int nRequestID);

    int CallReqQryDepthMarketData(ReqQryDepthMarketData fn, void *this, struct CThostFtdcQryDepthMarketDataField *pQryDepthMarketData, int nRequestID);

    int CallReqQryTraderOffer(ReqQryTraderOffer fn, void *this, struct CThostFtdcQryTraderOfferField *pQryTraderOffer, int nRequestID);

    int CallReqQrySettlementInfo(ReqQrySettlementInfo fn, void *this, struct CThostFtdcQrySettlementInfoField *pQrySettlementInfo, int nRequestID);

    int CallReqQryTransferBank(ReqQryTransferBank fn, void *this, struct CThostFtdcQryTransferBankField *pQryTransferBank, int nRequestID);

    int CallReqQryInvestorPositionDetail(ReqQryInvestorPositionDetail fn, void *this, struct CThostFtdcQryInvestorPositionDetailField *pQryInvestorPositionDetail, int nRequestID);

    int CallReqQryNotice(ReqQryNotice fn, void *this, struct CThostFtdcQryNoticeField *pQryNotice, int nRequestID);

    int CallReqQrySettlementInfoConfirm(ReqQrySettlementInfoConfirm fn, void *this, struct CThostFtdcQrySettlementInfoConfirmField *pQrySettlementInfoConfirm, int nRequestID);

    int CallReqQryInvestorPositionCombineDetail(ReqQryInvestorPositionCombineDetail fn, void *this, struct CThostFtdcQryInvestorPositionCombineDetailField *pQryInvestorPositionCombineDetail, int nRequestID);

    int CallReqQryCFMMCTradingAccountKey(ReqQryCFMMCTradingAccountKey fn, void *this, struct CThostFtdcQryCFMMCTradingAccountKeyField *pQryCFMMCTradingAccountKey, int nRequestID);

    int CallReqQryEWarrantOffset(ReqQryEWarrantOffset fn, void *this, struct CThostFtdcQryEWarrantOffsetField *pQryEWarrantOffset, int nRequestID);

    int CallReqQryInvestorProductGroupMargin(ReqQryInvestorProductGroupMargin fn, void *this, struct CThostFtdcQryInvestorProductGroupMarginField *pQryInvestorProductGroupMargin, int nRequestID);

    int CallReqQryExchangeMarginRate(ReqQryExchangeMarginRate fn, void *this, struct CThostFtdcQryExchangeMarginRateField *pQryExchangeMarginRate, int nRequestID);

    int CallReqQryExchangeMarginRateAdjust(ReqQryExchangeMarginRateAdjust fn, void *this, struct CThostFtdcQryExchangeMarginRateAdjustField *pQryExchangeMarginRateAdjust, int nRequestID);

    int CallReqQryExchangeRate(ReqQryExchangeRate fn, void *this, struct CThostFtdcQryExchangeRateField *pQryExchangeRate, int nRequestID);

    int CallReqQrySecAgentACIDMap(ReqQrySecAgentACIDMap fn, void *this, struct CThostFtdcQrySecAgentACIDMapField *pQrySecAgentACIDMap, int nRequestID);

    int CallReqQryProductExchRate(ReqQryProductExchRate fn, void *this, struct CThostFtdcQryProductExchRateField *pQryProductExchRate, int nRequestID);

    int CallReqQryProductGroup(ReqQryProductGroup fn, void *this, struct CThostFtdcQryProductGroupField *pQryProductGroup, int nRequestID);

    int CallReqQryMMInstrumentCommissionRate(ReqQryMMInstrumentCommissionRate fn, void *this, struct CThostFtdcQryMMInstrumentCommissionRateField *pQryMMInstrumentCommissionRate, int nRequestID);

    int CallReqQryMMOptionInstrCommRate(ReqQryMMOptionInstrCommRate fn, void *this, struct CThostFtdcQryMMOptionInstrCommRateField *pQryMMOptionInstrCommRate, int nRequestID);

    int CallReqQryInstrumentOrderCommRate(ReqQryInstrumentOrderCommRate fn, void *this, struct CThostFtdcQryInstrumentOrderCommRateField *pQryInstrumentOrderCommRate, int nRequestID);

    int CallReqQrySecAgentTradingAccount(ReqQrySecAgentTradingAccount fn, void *this, struct CThostFtdcQryTradingAccountField *pQryTradingAccount, int nRequestID);

    int CallReqQrySecAgentCheckMode(ReqQrySecAgentCheckMode fn, void *this, struct CThostFtdcQrySecAgentCheckModeField *pQrySecAgentCheckMode, int nRequestID);

    int CallReqQrySecAgentTradeInfo(ReqQrySecAgentTradeInfo fn, void *this, struct CThostFtdcQrySecAgentTradeInfoField *pQrySecAgentTradeInfo, int nRequestID);

    int CallReqQryOptionInstrTradeCost(ReqQryOptionInstrTradeCost fn, void *this, struct CThostFtdcQryOptionInstrTradeCostField *pQryOptionInstrTradeCost, int nRequestID);

    int CallReqQryOptionInstrCommRate(ReqQryOptionInstrCommRate fn, void *this, struct CThostFtdcQryOptionInstrCommRateField *pQryOptionInstrCommRate, int nRequestID);

    int CallReqQryExecOrder(ReqQryExecOrder fn, void *this, struct CThostFtdcQryExecOrderField *pQryExecOrder, int nRequestID);

    int CallReqQryForQuote(ReqQryForQuote fn, void *this, struct CThostFtdcQryForQuoteField *pQryForQuote, int nRequestID);

    int CallReqQryQuote(ReqQryQuote fn, void *this, struct CThostFtdcQryQuoteField *pQryQuote, int nRequestID);

    int CallReqQryOptionSelfClose(ReqQryOptionSelfClose fn, void *this, struct CThostFtdcQryOptionSelfCloseField *pQryOptionSelfClose, int nRequestID);

    int CallReqQryInvestUnit(ReqQryInvestUnit fn, void *this, struct CThostFtdcQryInvestUnitField *pQryInvestUnit, int nRequestID);

    int CallReqQryCombInstrumentGuard(ReqQryCombInstrumentGuard fn, void *this, struct CThostFtdcQryCombInstrumentGuardField *pQryCombInstrumentGuard, int nRequestID);

    int CallReqQryCombAction(ReqQryCombAction fn, void *this, struct CThostFtdcQryCombActionField *pQryCombAction, int nRequestID);

    int CallReqQryTransferSerial(ReqQryTransferSerial fn, void *this, struct CThostFtdcQryTransferSerialField *pQryTransferSerial, int nRequestID);

    int CallReqQryAccountregister(ReqQryAccountregister fn, void *this, struct CThostFtdcQryAccountregisterField *pQryAccountregister, int nRequestID);

    int CallReqQryContractBank(ReqQryContractBank fn, void *this, struct CThostFtdcQryContractBankField *pQryContractBank, int nRequestID);

    int CallReqQryParkedOrder(ReqQryParkedOrder fn, void *this, struct CThostFtdcQryParkedOrderField *pQryParkedOrder, int nRequestID);

    int CallReqQryParkedOrderAction(ReqQryParkedOrderAction fn, void *this, struct CThostFtdcQryParkedOrderActionField *pQryParkedOrderAction, int nRequestID);

    int CallReqQryTradingNotice(ReqQryTradingNotice fn, void *this, struct CThostFtdcQryTradingNoticeField *pQryTradingNotice, int nRequestID);

    int CallReqQryBrokerTradingParams(ReqQryBrokerTradingParams fn, void *this, struct CThostFtdcQryBrokerTradingParamsField *pQryBrokerTradingParams, int nRequestID);

    int CallReqQryBrokerTradingAlgos(ReqQryBrokerTradingAlgos fn, void *this, struct CThostFtdcQryBrokerTradingAlgosField *pQryBrokerTradingAlgos, int nRequestID);

    int CallReqQueryCFMMCTradingAccountToken(ReqQueryCFMMCTradingAccountToken fn, void *this, struct CThostFtdcQueryCFMMCTradingAccountTokenField *pQueryCFMMCTradingAccountToken, int nRequestID);

    int CallReqFromBankToFutureByFuture(ReqFromBankToFutureByFuture fn, void *this, struct CThostFtdcReqTransferField *pReqTransfer, int nRequestID);

    int CallReqFromFutureToBankByFuture(ReqFromFutureToBankByFuture fn, void *this, struct CThostFtdcReqTransferField *pReqTransfer, int nRequestID);

    int CallReqQueryBankAccountMoneyByFuture(ReqQueryBankAccountMoneyByFuture fn, void *this, struct CThostFtdcReqQueryAccountField *pReqQueryAccount, int nRequestID);

    int CallReqQryClassifiedInstrument(ReqQryClassifiedInstrument fn, void *this, struct CThostFtdcQryClassifiedInstrumentField *pQryClassifiedInstrument, int nRequestID);

    int CallReqQryCombPromotionParam(ReqQryCombPromotionParam fn, void *this, struct CThostFtdcQryCombPromotionParamField *pQryCombPromotionParam, int nRequestID);

    int CallReqQryRiskSettleInvstPosition(ReqQryRiskSettleInvstPosition fn, void *this, struct CThostFtdcQryRiskSettleInvstPositionField *pQryRiskSettleInvstPosition, int nRequestID);

    int CallReqQryRiskSettleProductStatus(ReqQryRiskSettleProductStatus fn, void *this, struct CThostFtdcQryRiskSettleProductStatusField *pQryRiskSettleProductStatus, int nRequestID);

    int CallReqQrySPBMFutureParameter(ReqQrySPBMFutureParameter fn, void *this, struct CThostFtdcQrySPBMFutureParameterField *pQrySPBMFutureParameter, int nRequestID);

    int CallReqQrySPBMOptionParameter(ReqQrySPBMOptionParameter fn, void *this, struct CThostFtdcQrySPBMOptionParameterField *pQrySPBMOptionParameter, int nRequestID);

    int CallReqQrySPBMIntraParameter(ReqQrySPBMIntraParameter fn, void *this, struct CThostFtdcQrySPBMIntraParameterField *pQrySPBMIntraParameter, int nRequestID);

    int CallReqQrySPBMInterParameter(ReqQrySPBMInterParameter fn, void *this, struct CThostFtdcQrySPBMInterParameterField *pQrySPBMInterParameter, int nRequestID);

    int CallReqQrySPBMPortfDefinition(ReqQrySPBMPortfDefinition fn, void *this, struct CThostFtdcQrySPBMPortfDefinitionField *pQrySPBMPortfDefinition, int nRequestID);

    int CallReqQrySPBMInvestorPortfDef(ReqQrySPBMInvestorPortfDef fn, void *this, struct CThostFtdcQrySPBMInvestorPortfDefField *pQrySPBMInvestorPortfDef, int nRequestID);

    int CallReqQryInvestorPortfMarginRatio(ReqQryInvestorPortfMarginRatio fn, void *this, struct CThostFtdcQryInvestorPortfMarginRatioField *pQryInvestorPortfMarginRatio, int nRequestID);

    int CallReqQryInvestorProdSPBMDetail(ReqQryInvestorProdSPBMDetail fn, void *this, struct CThostFtdcQryInvestorProdSPBMDetailField *pQryInvestorProdSPBMDetail, int nRequestID);

    int CallReqQryInvestorCommoditySPMMMargin(ReqQryInvestorCommoditySPMMMargin fn, void *this, struct CThostFtdcQryInvestorCommoditySPMMMarginField *pQryInvestorCommoditySPMMMargin, int nRequestID);

    int CallReqQryInvestorCommodityGroupSPMMMargin(ReqQryInvestorCommodityGroupSPMMMargin fn, void *this, struct CThostFtdcQryInvestorCommodityGroupSPMMMarginField *pQryInvestorCommodityGroupSPMMMargin, int nRequestID);

    int CallReqQrySPMMInstParam(ReqQrySPMMInstParam fn, void *this, struct CThostFtdcQrySPMMInstParamField *pQrySPMMInstParam, int nRequestID);

    int CallReqQrySPMMProductParam(ReqQrySPMMProductParam fn, void *this, struct CThostFtdcQrySPMMProductParamField *pQrySPMMProductParam, int nRequestID);

    int CallReqQrySPBMAddOnInterParameter(ReqQrySPBMAddOnInterParameter fn, void *this, struct CThostFtdcQrySPBMAddOnInterParameterField *pQrySPBMAddOnInterParameter, int nRequestID);

    int CallReqQryRCAMSCombProductInfo(ReqQryRCAMSCombProductInfo fn, void *this, struct CThostFtdcQryRCAMSCombProductInfoField *pQryRCAMSCombProductInfo, int nRequestID);

    int CallReqQryRCAMSInstrParameter(ReqQryRCAMSInstrParameter fn, void *this, struct CThostFtdcQryRCAMSInstrParameterField *pQryRCAMSInstrParameter, int nRequestID);

    int CallReqQryRCAMSIntraParameter(ReqQryRCAMSIntraParameter fn, void *this, struct CThostFtdcQryRCAMSIntraParameterField *pQryRCAMSIntraParameter, int nRequestID);

    int CallReqQryRCAMSInterParameter(ReqQryRCAMSInterParameter fn, void *this, struct CThostFtdcQryRCAMSInterParameterField *pQryRCAMSInterParameter, int nRequestID);

    int CallReqQryRCAMSShortOptAdjustParam(ReqQryRCAMSShortOptAdjustParam fn, void *this, struct CThostFtdcQryRCAMSShortOptAdjustParamField *pQryRCAMSShortOptAdjustParam, int nRequestID);

    int CallReqQryRCAMSInvestorCombPosition(ReqQryRCAMSInvestorCombPosition fn, void *this, struct CThostFtdcQryRCAMSInvestorCombPositionField *pQryRCAMSInvestorCombPosition, int nRequestID);

    int CallReqQryInvestorProdRCAMSMargin(ReqQryInvestorProdRCAMSMargin fn, void *this, struct CThostFtdcQryInvestorProdRCAMSMarginField *pQryInvestorProdRCAMSMargin, int nRequestID);

    int CallReqQryRULEInstrParameter(ReqQryRULEInstrParameter fn, void *this, struct CThostFtdcQryRULEInstrParameterField *pQryRULEInstrParameter, int nRequestID);

    int CallReqQryRULEIntraParameter(ReqQryRULEIntraParameter fn, void *this, struct CThostFtdcQryRULEIntraParameterField *pQryRULEIntraParameter, int nRequestID);

    int CallReqQryRULEInterParameter(ReqQryRULEInterParameter fn, void *this, struct CThostFtdcQryRULEInterParameterField *pQryRULEInterParameter, int nRequestID);

    int CallReqQryInvestorProdRULEMargin(ReqQryInvestorProdRULEMargin fn, void *this, struct CThostFtdcQryInvestorProdRULEMarginField *pQryInvestorProdRULEMargin, int nRequestID);

    int CallReqQryInvestorPortfSetting(ReqQryInvestorPortfSetting fn, void *this, struct CThostFtdcQryInvestorPortfSettingField *pQryInvestorPortfSetting, int nRequestID);

    int CallReqQryInvestorInfoCommRec(ReqQryInvestorInfoCommRec fn, void *this, struct CThostFtdcQryInvestorInfoCommRecField *pQryInvestorInfoCommRec, int nRequestID);

    int CallReqQryCombLeg(ReqQryCombLeg fn, void *this, struct CThostFtdcQryCombLegField *pQryCombLeg, int nRequestID);

    int CallReqOffsetSetting(ReqOffsetSetting fn, void *this, struct CThostFtdcInputOffsetSettingField *pInputOffsetSetting, int nRequestID);

    int CallReqCancelOffsetSetting(ReqCancelOffsetSetting fn, void *this, struct CThostFtdcInputOffsetSettingField *pInputOffsetSetting, int nRequestID);

    int CallReqQryOffsetSetting(ReqQryOffsetSetting fn, void *this, struct CThostFtdcQryOffsetSettingField *pQryOffsetSetting, int nRequestID);

    int CallReqGenSMSCode(ReqGenSMSCode fn, void *this, struct CThostFtdcReqGenSMSCodeField *pReqGenSMSCode, int nRequestID);

    int CallReqSpdApply(ReqSpdApply fn, void *this, struct CThostFtdcInputSpdApplyField *pInputSpdApply, int nRequestID);

    int CallReqSpdApplyAction(ReqSpdApplyAction fn, void *this, struct CThostFtdcInputSpdApplyActionField *pInputSpdApplyAction, int nRequestID);

    int CallReqQrySpdApply(ReqQrySpdApply fn, void *this, struct CThostFtdcQrySpdApplyField *pQrySpdApply, int nRequestID);

    int CallReqHedgeCfm(ReqHedgeCfm fn, void *this, struct CThostFtdcInputHedgeCfmField *pInputHedgeCfm, int nRequestID);

    int CallReqHedgeCfmAction(ReqHedgeCfmAction fn, void *this, struct CThostFtdcInputHedgeCfmActionField *pInputHedgeCfmAction, int nRequestID);

    int CallReqQryHedgeCfm(ReqQryHedgeCfm fn, void *this, struct CThostFtdcQryHedgeCfmField *pQryHedgeCfm, int nRequestID);

#ifdef __cplusplus
}
#endif
#endif