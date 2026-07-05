#include "spi_helper.h"
extern void CgoOnFrontConnected(void *this);

extern void CgoOnFrontDisconnected(void *this, int nReason);

extern void CgoOnHeartBeatWarning(void *this, int nTimeLapse);

extern void CgoOnRspSubscribeFlowCtrlWarning(void *this, struct CThostFtdcSpecificTraderField* pRspSubscribeTraderField, struct CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast);

extern void CgoOnRspUnSubscribeFlowCtrlWarning(void *this, struct CThostFtdcSpecificTraderField* pRspSubscribeTraderField, struct CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast);

extern void CgoOnRspAuthenticate(void *this, struct CThostFtdcRspAuthenticateField* pRspAuthenticateField, struct CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast);

extern void CgoOnRspUserLogin(void *this, struct CThostFtdcRspUserLoginField* pRspUserLogin, struct CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast);

extern void CgoOnRspUserLogout(void *this, struct CThostFtdcUserLogoutField* pUserLogout, struct CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast);

extern void CgoOnRspOrderInsert(void *this, struct CThostFtdcInputOrderField* pInputOrder, struct CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast);

extern void CgoOnRspOrderAction(void *this, struct CThostFtdcInputOrderActionField* pInputOrderAction, struct CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast);

extern void CgoOnRspMKBatchOrderAction(void *this, struct CThostFtdcMKInputOrderActionField* pMKInputOrderAction, struct CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast);

extern void CgoOnRspExecOrderInsert(void *this, struct CThostFtdcInputExecOrderField* pInputExecOrder, struct CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast);

extern void CgoOnRspExecOrderAction(void *this, struct CThostFtdcInputExecOrderActionField* pInputExecOrderAction, struct CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast);

extern void CgoOnRspForQuoteInsert(void *this, struct CThostFtdcInputForQuoteField* pInputForQuote, struct CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast);

extern void CgoOnRspQuoteInsert(void *this, struct CThostFtdcInputQuoteField* pInputQuote, struct CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast);

extern void CgoOnRspQuoteAction(void *this, struct CThostFtdcInputQuoteActionField* pInputQuoteAction, struct CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast);

extern void CgoOnRspBatchOrderAction(void *this, struct CThostFtdcInputBatchOrderActionField* pInputBatchOrderAction, struct CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast);

extern void CgoOnRspOptionSelfCloseInsert(void *this, struct CThostFtdcInputOptionSelfCloseField* pInputOptionSelfClose, struct CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast);

extern void CgoOnRspOptionSelfCloseAction(void *this, struct CThostFtdcInputOptionSelfCloseActionField* pInputOptionSelfCloseAction, struct CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast);

extern void CgoOnRspCombActionInsert(void *this, struct CThostFtdcInputCombActionField* pInputCombAction, struct CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast);

extern void CgoOnRspQryOrder(void *this, struct CThostFtdcOrderField* pOrder, struct CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast);

extern void CgoOnRspQryTrade(void *this, struct CThostFtdcTradeField* pTrade, struct CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast);

extern void CgoOnRspQryInvestorPosition(void *this, struct CThostFtdcInvestorPositionField* pInvestorPosition, struct CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast);

extern void CgoOnRspQryTradingAccount(void *this, struct CThostFtdcTradingAccountField* pTradingAccount, struct CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast);

extern void CgoOnRspQryInvestor(void *this, struct CThostFtdcInvestorField* pInvestor, struct CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast);

extern void CgoOnRspQryTradingCode(void *this, struct CThostFtdcTradingCodeField* pTradingCode, struct CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast);

extern void CgoOnRspQryInstrumentMarginRate(void *this, struct CThostFtdcInstrumentMarginRateField* pInstrumentMarginRate, struct CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast);

extern void CgoOnRspQryInstrumentCommissionRate(void *this, struct CThostFtdcInstrumentCommissionRateField* pInstrumentCommissionRate, struct CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast);

extern void CgoOnRspQryExchange(void *this, struct CThostFtdcExchangeField* pExchange, struct CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast);

extern void CgoOnRspQryProduct(void *this, struct CThostFtdcProductField* pProduct, struct CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast);

extern void CgoOnRspQryInstrument(void *this, struct CThostFtdcInstrumentField* pInstrument, struct CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast);

extern void CgoOnRspQryCombInstrument(void *this, struct CThostFtdcCombInstrumentField* pCombInstrument, struct CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast);

extern void CgoOnRspQryRCAMSInvestorProdMargin(void *this, struct CThostFtdcRCAMSInvestorProdMarginField* pRCAMSInvestorProdMargin, struct CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast);

extern void CgoOnRspQryRCAMSInvestorCombPosition(void *this, struct CThostFtdcRCAMSInvestorCombPositionField* pRCAMSInvestorCombPosition, struct CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast);

extern void CgoOnRspQryCombAction(void *this, struct CThostFtdcCombActionField* pCombAction, struct CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast);

extern void CgoOnRspQryInvestorPositionForComb(void *this, struct CThostFtdcInvestorPositionForCombField* pForComb, struct CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast);

extern void CgoOnRspQryDepthMarketData(void *this, struct CThostFtdcDepthMarketDataField* pDepthMarketData, struct CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast);

extern void CgoOnRspQryInstrumentStatus(void *this, struct CThostFtdcInstrumentStatusField* pInstrumentStatus, struct CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast);

extern void CgoOnRspQryInvestorPositionDetail(void *this, struct CThostFtdcInvestorPositionDetailField* pInvestorPositionDetail, struct CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast);

extern void CgoOnRspQryExchangeMarginRate(void *this, struct CThostFtdcExchangeMarginRateField* pExchangeMarginRate, struct CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast);

extern void CgoOnRspQryExchangeMarginRateAdjust(void *this, struct CThostFtdcExchangeMarginRateAdjustField* pExchangeMarginRateAdjust, struct CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast);

extern void CgoOnRspQryOptionInstrTradeCost(void *this, struct CThostFtdcOptionInstrTradeCostField* pOptionInstrTradeCost, struct CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast);

extern void CgoOnRspQryOptionInstrCommRate(void *this, struct CThostFtdcOptionInstrCommRateField* pOptionInstrCommRate, struct CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast);

extern void CgoOnRspQryExecOrder(void *this, struct CThostFtdcExecOrderField* pExecOrder, struct CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast);

extern void CgoOnRspQryForQuote(void *this, struct CThostFtdcForQuoteField* pForQuote, struct CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast);

extern void CgoOnRspQryForQuoteParam(void *this, struct CThostFtdcForQuoteParamField* pForQuoteParam, struct CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast);

extern void CgoOnRspQryInvestorProdSPBMDetail(void *this, struct CThostFtdcInvestorProdSPBMDetailField* pInvestorProdSPBMDetail, struct CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast);

extern void CgoOnRspQrySPMMInvestorCommodityGroupMargin(void *this, struct CThostFtdcSPMMInvestorCommodityGroupMarginField* pSPMMInvestorCommodityGroupMargin, struct CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast);

extern void CgoOnRspQryRULEInvestorProdMargin(void *this, struct CThostFtdcRULEInvestorProdMarginField* pRULEInvestorProdMargin, struct CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast);

extern void CgoOnRspQryTraderOffer(void *this, struct CThostFtdcTraderOfferField* pTraderOffer, struct CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast);

extern void CgoOnRspQryQuote(void *this, struct CThostFtdcQuoteField* pQuote, struct CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast);

extern void CgoOnRspQryOptionSelfClose(void *this, struct CThostFtdcOptionSelfCloseField* pOptionSelfClose, struct CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast);

extern void CgoOnRspQryControlParam(void *this, struct CThostFtdcControlParamField* pControlParam, struct CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast);

extern void CgoOnRspQryOffsetSetting(void *this, struct CThostFtdcOffsetSettingField* pOffsetSetting, struct CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast);

extern void CgoOnRspError(void *this, struct CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast);

extern void CgoOnRtnOrder(void *this, struct CThostFtdcOrderField* pOrder);

extern void CgoOnRtnTrade(void *this, struct CThostFtdcTradeField* pTrade);

extern void CgoOnErrRtnOrderInsert(void *this, struct CThostFtdcInputOrderField* pInputOrder, struct CThostFtdcRspInfoField* pRspInfo);

extern void CgoOnErrRtnOrderAction(void *this, struct CThostFtdcOrderActionField* pOrderAction, struct CThostFtdcRspInfoField* pRspInfo);

extern void CgoOnRtnInstrumentStatus(void *this, struct CThostFtdcInstrumentStatusField* pInstrumentStatus);

extern void CgoOnRtnExecOrder(void *this, struct CThostFtdcExecOrderField* pExecOrder);

extern void CgoOnErrRtnExecOrderInsert(void *this, struct CThostFtdcInputExecOrderField* pInputExecOrder, struct CThostFtdcRspInfoField* pRspInfo);

extern void CgoOnErrRtnExecOrderAction(void *this, struct CThostFtdcExecOrderActionField* pExecOrderAction, struct CThostFtdcRspInfoField* pRspInfo);

extern void CgoOnErrRtnForQuoteInsert(void *this, struct CThostFtdcInputForQuoteField* pInputForQuote, struct CThostFtdcRspInfoField* pRspInfo);

extern void CgoOnRtnQuote(void *this, struct CThostFtdcQuoteField* pQuote);

extern void CgoOnErrRtnQuoteInsert(void *this, struct CThostFtdcInputQuoteField* pInputQuote, struct CThostFtdcRspInfoField* pRspInfo);

extern void CgoOnErrRtnQuoteAction(void *this, struct CThostFtdcQuoteActionField* pQuoteAction, struct CThostFtdcRspInfoField* pRspInfo);

extern void CgoOnRtnForQuoteRsp(void *this, struct CThostFtdcForQuoteRspField* pForQuoteRsp);

extern void CgoOnErrRtnBatchOrderAction(void *this, struct CThostFtdcBatchOrderActionField* pBatchOrderAction, struct CThostFtdcRspInfoField* pRspInfo);

extern void CgoOnRtnOptionSelfClose(void *this, struct CThostFtdcOptionSelfCloseField* pOptionSelfClose);

extern void CgoOnErrRtnOptionSelfCloseInsert(void *this, struct CThostFtdcInputOptionSelfCloseField* pInputOptionSelfClose, struct CThostFtdcRspInfoField* pRspInfo);

extern void CgoOnErrRtnOptionSelfCloseAction(void *this, struct CThostFtdcOptionSelfCloseActionField* pOptionSelfCloseAction, struct CThostFtdcRspInfoField* pRspInfo);

extern void CgoOnRtnCombAction(void *this, struct CThostFtdcCombActionField* pCombAction);

extern void CgoOnRspQryInstrumentOrderCommRate(void *this, struct CThostFtdcInstrumentOrderCommRateField* pInstrumentOrderCommRate, struct CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast);

extern void CgoOnRtnFlowCtrlWarning(void *this, struct CThostFtdcFlowCtrlWarningField* pFlowCtrlWarning);

extern void CgoOnRspSubscribeFundChange(void *this, struct CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast);

extern void CgoOnRspUnSubscribeFundChange(void *this, struct CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast);

extern void CgoOnRtnFundChange(void *this, struct CThostFtdcTradingAccountField* pTradingAccount);

extern void CgoOnRspOffsetSetting(void *this, struct CThostFtdcInputOffsetSettingField* pInputOffsetSetting, struct CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast);

extern void CgoOnRspCancelOffsetSetting(void *this, struct CThostFtdcInputOffsetSettingField* pInputOffsetSetting, struct CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast);

extern void CgoOnRtnOffsetSetting(void *this, struct CThostFtdcOffsetSettingField* pOffsetSetting);

extern void CgoOnErrRtnCancelOffsetSetting(void *this, struct CThostFtdcCancelOffsetSettingField* pCancelOffsetSetting, struct CThostFtdcRspInfoField* pRspInfo);

extern void CgoOnRspUserPasswordUpdate(void *this, struct CThostFtdcUserPasswordUpdateField* pUserPasswordUpdate, struct CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast);

void COnFrontConnected(void *this)
{
    return CgoOnFrontConnected(this);
}

void COnFrontDisconnected(void *this, int nReason)
{
    return CgoOnFrontDisconnected(this, nReason);
}

void COnHeartBeatWarning(void *this, int nTimeLapse)
{
    return CgoOnHeartBeatWarning(this, nTimeLapse);
}

void COnRspSubscribeFlowCtrlWarning(void *this, struct CThostFtdcSpecificTraderField* pRspSubscribeTraderField, struct CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast)
{
    return CgoOnRspSubscribeFlowCtrlWarning(this, pRspSubscribeTraderField, pRspInfo, nRequestID, bIsLast);
}

void COnRspUnSubscribeFlowCtrlWarning(void *this, struct CThostFtdcSpecificTraderField* pRspSubscribeTraderField, struct CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast)
{
    return CgoOnRspUnSubscribeFlowCtrlWarning(this, pRspSubscribeTraderField, pRspInfo, nRequestID, bIsLast);
}

void COnRspAuthenticate(void *this, struct CThostFtdcRspAuthenticateField* pRspAuthenticateField, struct CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast)
{
    return CgoOnRspAuthenticate(this, pRspAuthenticateField, pRspInfo, nRequestID, bIsLast);
}

void COnRspUserLogin(void *this, struct CThostFtdcRspUserLoginField* pRspUserLogin, struct CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast)
{
    return CgoOnRspUserLogin(this, pRspUserLogin, pRspInfo, nRequestID, bIsLast);
}

void COnRspUserLogout(void *this, struct CThostFtdcUserLogoutField* pUserLogout, struct CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast)
{
    return CgoOnRspUserLogout(this, pUserLogout, pRspInfo, nRequestID, bIsLast);
}

void COnRspOrderInsert(void *this, struct CThostFtdcInputOrderField* pInputOrder, struct CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast)
{
    return CgoOnRspOrderInsert(this, pInputOrder, pRspInfo, nRequestID, bIsLast);
}

void COnRspOrderAction(void *this, struct CThostFtdcInputOrderActionField* pInputOrderAction, struct CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast)
{
    return CgoOnRspOrderAction(this, pInputOrderAction, pRspInfo, nRequestID, bIsLast);
}

void COnRspMKBatchOrderAction(void *this, struct CThostFtdcMKInputOrderActionField* pMKInputOrderAction, struct CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast)
{
    return CgoOnRspMKBatchOrderAction(this, pMKInputOrderAction, pRspInfo, nRequestID, bIsLast);
}

void COnRspExecOrderInsert(void *this, struct CThostFtdcInputExecOrderField* pInputExecOrder, struct CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast)
{
    return CgoOnRspExecOrderInsert(this, pInputExecOrder, pRspInfo, nRequestID, bIsLast);
}

void COnRspExecOrderAction(void *this, struct CThostFtdcInputExecOrderActionField* pInputExecOrderAction, struct CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast)
{
    return CgoOnRspExecOrderAction(this, pInputExecOrderAction, pRspInfo, nRequestID, bIsLast);
}

void COnRspForQuoteInsert(void *this, struct CThostFtdcInputForQuoteField* pInputForQuote, struct CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast)
{
    return CgoOnRspForQuoteInsert(this, pInputForQuote, pRspInfo, nRequestID, bIsLast);
}

void COnRspQuoteInsert(void *this, struct CThostFtdcInputQuoteField* pInputQuote, struct CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast)
{
    return CgoOnRspQuoteInsert(this, pInputQuote, pRspInfo, nRequestID, bIsLast);
}

void COnRspQuoteAction(void *this, struct CThostFtdcInputQuoteActionField* pInputQuoteAction, struct CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast)
{
    return CgoOnRspQuoteAction(this, pInputQuoteAction, pRspInfo, nRequestID, bIsLast);
}

void COnRspBatchOrderAction(void *this, struct CThostFtdcInputBatchOrderActionField* pInputBatchOrderAction, struct CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast)
{
    return CgoOnRspBatchOrderAction(this, pInputBatchOrderAction, pRspInfo, nRequestID, bIsLast);
}

void COnRspOptionSelfCloseInsert(void *this, struct CThostFtdcInputOptionSelfCloseField* pInputOptionSelfClose, struct CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast)
{
    return CgoOnRspOptionSelfCloseInsert(this, pInputOptionSelfClose, pRspInfo, nRequestID, bIsLast);
}

void COnRspOptionSelfCloseAction(void *this, struct CThostFtdcInputOptionSelfCloseActionField* pInputOptionSelfCloseAction, struct CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast)
{
    return CgoOnRspOptionSelfCloseAction(this, pInputOptionSelfCloseAction, pRspInfo, nRequestID, bIsLast);
}

void COnRspCombActionInsert(void *this, struct CThostFtdcInputCombActionField* pInputCombAction, struct CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast)
{
    return CgoOnRspCombActionInsert(this, pInputCombAction, pRspInfo, nRequestID, bIsLast);
}

void COnRspQryOrder(void *this, struct CThostFtdcOrderField* pOrder, struct CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast)
{
    return CgoOnRspQryOrder(this, pOrder, pRspInfo, nRequestID, bIsLast);
}

void COnRspQryTrade(void *this, struct CThostFtdcTradeField* pTrade, struct CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast)
{
    return CgoOnRspQryTrade(this, pTrade, pRspInfo, nRequestID, bIsLast);
}

void COnRspQryInvestorPosition(void *this, struct CThostFtdcInvestorPositionField* pInvestorPosition, struct CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast)
{
    return CgoOnRspQryInvestorPosition(this, pInvestorPosition, pRspInfo, nRequestID, bIsLast);
}

void COnRspQryTradingAccount(void *this, struct CThostFtdcTradingAccountField* pTradingAccount, struct CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast)
{
    return CgoOnRspQryTradingAccount(this, pTradingAccount, pRspInfo, nRequestID, bIsLast);
}

void COnRspQryInvestor(void *this, struct CThostFtdcInvestorField* pInvestor, struct CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast)
{
    return CgoOnRspQryInvestor(this, pInvestor, pRspInfo, nRequestID, bIsLast);
}

void COnRspQryTradingCode(void *this, struct CThostFtdcTradingCodeField* pTradingCode, struct CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast)
{
    return CgoOnRspQryTradingCode(this, pTradingCode, pRspInfo, nRequestID, bIsLast);
}

void COnRspQryInstrumentMarginRate(void *this, struct CThostFtdcInstrumentMarginRateField* pInstrumentMarginRate, struct CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast)
{
    return CgoOnRspQryInstrumentMarginRate(this, pInstrumentMarginRate, pRspInfo, nRequestID, bIsLast);
}

void COnRspQryInstrumentCommissionRate(void *this, struct CThostFtdcInstrumentCommissionRateField* pInstrumentCommissionRate, struct CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast)
{
    return CgoOnRspQryInstrumentCommissionRate(this, pInstrumentCommissionRate, pRspInfo, nRequestID, bIsLast);
}

void COnRspQryExchange(void *this, struct CThostFtdcExchangeField* pExchange, struct CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast)
{
    return CgoOnRspQryExchange(this, pExchange, pRspInfo, nRequestID, bIsLast);
}

void COnRspQryProduct(void *this, struct CThostFtdcProductField* pProduct, struct CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast)
{
    return CgoOnRspQryProduct(this, pProduct, pRspInfo, nRequestID, bIsLast);
}

void COnRspQryInstrument(void *this, struct CThostFtdcInstrumentField* pInstrument, struct CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast)
{
    return CgoOnRspQryInstrument(this, pInstrument, pRspInfo, nRequestID, bIsLast);
}

void COnRspQryCombInstrument(void *this, struct CThostFtdcCombInstrumentField* pCombInstrument, struct CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast)
{
    return CgoOnRspQryCombInstrument(this, pCombInstrument, pRspInfo, nRequestID, bIsLast);
}

void COnRspQryRCAMSInvestorProdMargin(void *this, struct CThostFtdcRCAMSInvestorProdMarginField* pRCAMSInvestorProdMargin, struct CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast)
{
    return CgoOnRspQryRCAMSInvestorProdMargin(this, pRCAMSInvestorProdMargin, pRspInfo, nRequestID, bIsLast);
}

void COnRspQryRCAMSInvestorCombPosition(void *this, struct CThostFtdcRCAMSInvestorCombPositionField* pRCAMSInvestorCombPosition, struct CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast)
{
    return CgoOnRspQryRCAMSInvestorCombPosition(this, pRCAMSInvestorCombPosition, pRspInfo, nRequestID, bIsLast);
}

void COnRspQryCombAction(void *this, struct CThostFtdcCombActionField* pCombAction, struct CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast)
{
    return CgoOnRspQryCombAction(this, pCombAction, pRspInfo, nRequestID, bIsLast);
}

void COnRspQryInvestorPositionForComb(void *this, struct CThostFtdcInvestorPositionForCombField* pForComb, struct CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast)
{
    return CgoOnRspQryInvestorPositionForComb(this, pForComb, pRspInfo, nRequestID, bIsLast);
}

void COnRspQryDepthMarketData(void *this, struct CThostFtdcDepthMarketDataField* pDepthMarketData, struct CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast)
{
    return CgoOnRspQryDepthMarketData(this, pDepthMarketData, pRspInfo, nRequestID, bIsLast);
}

void COnRspQryInstrumentStatus(void *this, struct CThostFtdcInstrumentStatusField* pInstrumentStatus, struct CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast)
{
    return CgoOnRspQryInstrumentStatus(this, pInstrumentStatus, pRspInfo, nRequestID, bIsLast);
}

void COnRspQryInvestorPositionDetail(void *this, struct CThostFtdcInvestorPositionDetailField* pInvestorPositionDetail, struct CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast)
{
    return CgoOnRspQryInvestorPositionDetail(this, pInvestorPositionDetail, pRspInfo, nRequestID, bIsLast);
}

void COnRspQryExchangeMarginRate(void *this, struct CThostFtdcExchangeMarginRateField* pExchangeMarginRate, struct CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast)
{
    return CgoOnRspQryExchangeMarginRate(this, pExchangeMarginRate, pRspInfo, nRequestID, bIsLast);
}

void COnRspQryExchangeMarginRateAdjust(void *this, struct CThostFtdcExchangeMarginRateAdjustField* pExchangeMarginRateAdjust, struct CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast)
{
    return CgoOnRspQryExchangeMarginRateAdjust(this, pExchangeMarginRateAdjust, pRspInfo, nRequestID, bIsLast);
}

void COnRspQryOptionInstrTradeCost(void *this, struct CThostFtdcOptionInstrTradeCostField* pOptionInstrTradeCost, struct CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast)
{
    return CgoOnRspQryOptionInstrTradeCost(this, pOptionInstrTradeCost, pRspInfo, nRequestID, bIsLast);
}

void COnRspQryOptionInstrCommRate(void *this, struct CThostFtdcOptionInstrCommRateField* pOptionInstrCommRate, struct CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast)
{
    return CgoOnRspQryOptionInstrCommRate(this, pOptionInstrCommRate, pRspInfo, nRequestID, bIsLast);
}

void COnRspQryExecOrder(void *this, struct CThostFtdcExecOrderField* pExecOrder, struct CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast)
{
    return CgoOnRspQryExecOrder(this, pExecOrder, pRspInfo, nRequestID, bIsLast);
}

void COnRspQryForQuote(void *this, struct CThostFtdcForQuoteField* pForQuote, struct CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast)
{
    return CgoOnRspQryForQuote(this, pForQuote, pRspInfo, nRequestID, bIsLast);
}

void COnRspQryForQuoteParam(void *this, struct CThostFtdcForQuoteParamField* pForQuoteParam, struct CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast)
{
    return CgoOnRspQryForQuoteParam(this, pForQuoteParam, pRspInfo, nRequestID, bIsLast);
}

void COnRspQryInvestorProdSPBMDetail(void *this, struct CThostFtdcInvestorProdSPBMDetailField* pInvestorProdSPBMDetail, struct CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast)
{
    return CgoOnRspQryInvestorProdSPBMDetail(this, pInvestorProdSPBMDetail, pRspInfo, nRequestID, bIsLast);
}

void COnRspQrySPMMInvestorCommodityGroupMargin(void *this, struct CThostFtdcSPMMInvestorCommodityGroupMarginField* pSPMMInvestorCommodityGroupMargin, struct CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast)
{
    return CgoOnRspQrySPMMInvestorCommodityGroupMargin(this, pSPMMInvestorCommodityGroupMargin, pRspInfo, nRequestID, bIsLast);
}

void COnRspQryRULEInvestorProdMargin(void *this, struct CThostFtdcRULEInvestorProdMarginField* pRULEInvestorProdMargin, struct CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast)
{
    return CgoOnRspQryRULEInvestorProdMargin(this, pRULEInvestorProdMargin, pRspInfo, nRequestID, bIsLast);
}

void COnRspQryTraderOffer(void *this, struct CThostFtdcTraderOfferField* pTraderOffer, struct CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast)
{
    return CgoOnRspQryTraderOffer(this, pTraderOffer, pRspInfo, nRequestID, bIsLast);
}

void COnRspQryQuote(void *this, struct CThostFtdcQuoteField* pQuote, struct CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast)
{
    return CgoOnRspQryQuote(this, pQuote, pRspInfo, nRequestID, bIsLast);
}

void COnRspQryOptionSelfClose(void *this, struct CThostFtdcOptionSelfCloseField* pOptionSelfClose, struct CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast)
{
    return CgoOnRspQryOptionSelfClose(this, pOptionSelfClose, pRspInfo, nRequestID, bIsLast);
}

void COnRspQryControlParam(void *this, struct CThostFtdcControlParamField* pControlParam, struct CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast)
{
    return CgoOnRspQryControlParam(this, pControlParam, pRspInfo, nRequestID, bIsLast);
}

void COnRspQryOffsetSetting(void *this, struct CThostFtdcOffsetSettingField* pOffsetSetting, struct CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast)
{
    return CgoOnRspQryOffsetSetting(this, pOffsetSetting, pRspInfo, nRequestID, bIsLast);
}

void COnRspError(void *this, struct CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast)
{
    return CgoOnRspError(this, pRspInfo, nRequestID, bIsLast);
}

void COnRtnOrder(void *this, struct CThostFtdcOrderField* pOrder)
{
    return CgoOnRtnOrder(this, pOrder);
}

void COnRtnTrade(void *this, struct CThostFtdcTradeField* pTrade)
{
    return CgoOnRtnTrade(this, pTrade);
}

void COnErrRtnOrderInsert(void *this, struct CThostFtdcInputOrderField* pInputOrder, struct CThostFtdcRspInfoField* pRspInfo)
{
    return CgoOnErrRtnOrderInsert(this, pInputOrder, pRspInfo);
}

void COnErrRtnOrderAction(void *this, struct CThostFtdcOrderActionField* pOrderAction, struct CThostFtdcRspInfoField* pRspInfo)
{
    return CgoOnErrRtnOrderAction(this, pOrderAction, pRspInfo);
}

void COnRtnInstrumentStatus(void *this, struct CThostFtdcInstrumentStatusField* pInstrumentStatus)
{
    return CgoOnRtnInstrumentStatus(this, pInstrumentStatus);
}

void COnRtnExecOrder(void *this, struct CThostFtdcExecOrderField* pExecOrder)
{
    return CgoOnRtnExecOrder(this, pExecOrder);
}

void COnErrRtnExecOrderInsert(void *this, struct CThostFtdcInputExecOrderField* pInputExecOrder, struct CThostFtdcRspInfoField* pRspInfo)
{
    return CgoOnErrRtnExecOrderInsert(this, pInputExecOrder, pRspInfo);
}

void COnErrRtnExecOrderAction(void *this, struct CThostFtdcExecOrderActionField* pExecOrderAction, struct CThostFtdcRspInfoField* pRspInfo)
{
    return CgoOnErrRtnExecOrderAction(this, pExecOrderAction, pRspInfo);
}

void COnErrRtnForQuoteInsert(void *this, struct CThostFtdcInputForQuoteField* pInputForQuote, struct CThostFtdcRspInfoField* pRspInfo)
{
    return CgoOnErrRtnForQuoteInsert(this, pInputForQuote, pRspInfo);
}

void COnRtnQuote(void *this, struct CThostFtdcQuoteField* pQuote)
{
    return CgoOnRtnQuote(this, pQuote);
}

void COnErrRtnQuoteInsert(void *this, struct CThostFtdcInputQuoteField* pInputQuote, struct CThostFtdcRspInfoField* pRspInfo)
{
    return CgoOnErrRtnQuoteInsert(this, pInputQuote, pRspInfo);
}

void COnErrRtnQuoteAction(void *this, struct CThostFtdcQuoteActionField* pQuoteAction, struct CThostFtdcRspInfoField* pRspInfo)
{
    return CgoOnErrRtnQuoteAction(this, pQuoteAction, pRspInfo);
}

void COnRtnForQuoteRsp(void *this, struct CThostFtdcForQuoteRspField* pForQuoteRsp)
{
    return CgoOnRtnForQuoteRsp(this, pForQuoteRsp);
}

void COnErrRtnBatchOrderAction(void *this, struct CThostFtdcBatchOrderActionField* pBatchOrderAction, struct CThostFtdcRspInfoField* pRspInfo)
{
    return CgoOnErrRtnBatchOrderAction(this, pBatchOrderAction, pRspInfo);
}

void COnRtnOptionSelfClose(void *this, struct CThostFtdcOptionSelfCloseField* pOptionSelfClose)
{
    return CgoOnRtnOptionSelfClose(this, pOptionSelfClose);
}

void COnErrRtnOptionSelfCloseInsert(void *this, struct CThostFtdcInputOptionSelfCloseField* pInputOptionSelfClose, struct CThostFtdcRspInfoField* pRspInfo)
{
    return CgoOnErrRtnOptionSelfCloseInsert(this, pInputOptionSelfClose, pRspInfo);
}

void COnErrRtnOptionSelfCloseAction(void *this, struct CThostFtdcOptionSelfCloseActionField* pOptionSelfCloseAction, struct CThostFtdcRspInfoField* pRspInfo)
{
    return CgoOnErrRtnOptionSelfCloseAction(this, pOptionSelfCloseAction, pRspInfo);
}

void COnRtnCombAction(void *this, struct CThostFtdcCombActionField* pCombAction)
{
    return CgoOnRtnCombAction(this, pCombAction);
}

void COnRspQryInstrumentOrderCommRate(void *this, struct CThostFtdcInstrumentOrderCommRateField* pInstrumentOrderCommRate, struct CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast)
{
    return CgoOnRspQryInstrumentOrderCommRate(this, pInstrumentOrderCommRate, pRspInfo, nRequestID, bIsLast);
}

void COnRtnFlowCtrlWarning(void *this, struct CThostFtdcFlowCtrlWarningField* pFlowCtrlWarning)
{
    return CgoOnRtnFlowCtrlWarning(this, pFlowCtrlWarning);
}

void COnRspSubscribeFundChange(void *this, struct CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast)
{
    return CgoOnRspSubscribeFundChange(this, pRspInfo, nRequestID, bIsLast);
}

void COnRspUnSubscribeFundChange(void *this, struct CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast)
{
    return CgoOnRspUnSubscribeFundChange(this, pRspInfo, nRequestID, bIsLast);
}

void COnRtnFundChange(void *this, struct CThostFtdcTradingAccountField* pTradingAccount)
{
    return CgoOnRtnFundChange(this, pTradingAccount);
}

void COnRspOffsetSetting(void *this, struct CThostFtdcInputOffsetSettingField* pInputOffsetSetting, struct CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast)
{
    return CgoOnRspOffsetSetting(this, pInputOffsetSetting, pRspInfo, nRequestID, bIsLast);
}

void COnRspCancelOffsetSetting(void *this, struct CThostFtdcInputOffsetSettingField* pInputOffsetSetting, struct CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast)
{
    return CgoOnRspCancelOffsetSetting(this, pInputOffsetSetting, pRspInfo, nRequestID, bIsLast);
}

void COnRtnOffsetSetting(void *this, struct CThostFtdcOffsetSettingField* pOffsetSetting)
{
    return CgoOnRtnOffsetSetting(this, pOffsetSetting);
}

void COnErrRtnCancelOffsetSetting(void *this, struct CThostFtdcCancelOffsetSettingField* pCancelOffsetSetting, struct CThostFtdcRspInfoField* pRspInfo)
{
    return CgoOnErrRtnCancelOffsetSetting(this, pCancelOffsetSetting, pRspInfo);
}

void COnRspUserPasswordUpdate(void *this, struct CThostFtdcUserPasswordUpdateField* pUserPasswordUpdate, struct CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast)
{
    return CgoOnRspUserPasswordUpdate(this, pUserPasswordUpdate, pRspInfo, nRequestID, bIsLast);
}
