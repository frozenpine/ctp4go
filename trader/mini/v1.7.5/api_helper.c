#include "api_helper.h"

void* CallCreateFtdcTraderApi(CreateFtdcTraderApi fn, char* pszFlowPath)
{
    return fn(pszFlowPath);
}

char* CallGetApiVersion(GetApiVersion fn)
{
    return fn();
}

void CallRelease(Release fn, void *this)
{
    return fn(this);
}

void CallInit(Init fn, void *this, bool bContinuous)
{
    return fn(this, bContinuous);
}

int CallJoin(Join fn, void *this)
{
    return fn(this);
}

char* CallGetTradingDay(GetTradingDay fn, void *this)
{
    return fn(this);
}

void CallRegisterFront(RegisterFront fn, void *this, char* pszFrontAddress)
{
    return fn(this, pszFrontAddress);
}

void CallRegisterSpi(RegisterSpi fn, void *this, void *pSpi)
{
    return fn(this, pSpi);
}

void CallSubscribePrivateTopic(SubscribePrivateTopic fn, void *this, int nResumeType)
{
    return fn(this, nResumeType);
}

void CallSubscribePublicTopic(SubscribePublicTopic fn, void *this, int nResumeType)
{
    return fn(this, nResumeType);
}

int CallSubscribeFlowCtrlWarning(SubscribeFlowCtrlWarning fn, void *this, char* ppTraderID[], int nCount)
{
    return fn(this, ppTraderID, nCount);
}

int CallUnSubscribeFlowCtrlWarning(UnSubscribeFlowCtrlWarning fn, void *this, char* ppTraderID[], int nCount)
{
    return fn(this, ppTraderID, nCount);
}

int CallReqAuthenticate(ReqAuthenticate fn, void *this, struct CThostFtdcReqAuthenticateField* pReqAuthenticateField, int nRequestID)
{
    return fn(this, pReqAuthenticateField, nRequestID);
}

int CallReqUserLogin(ReqUserLogin fn, void *this, struct CThostFtdcReqUserLoginField* pReqUserLoginField, int nRequestID)
{
    return fn(this, pReqUserLoginField, nRequestID);
}

int CallReqUserLoginEncrypt(ReqUserLoginEncrypt fn, void *this, struct CThostFtdcReqUserLoginField* pReqUserLoginField, int nRequestID)
{
    return fn(this, pReqUserLoginField, nRequestID);
}

int CallReqUserLogout(ReqUserLogout fn, void *this, struct CThostFtdcUserLogoutField* pUserLogout, int nRequestID)
{
    return fn(this, pUserLogout, nRequestID);
}

int CallReqOrderInsert(ReqOrderInsert fn, void *this, struct CThostFtdcInputOrderField* pInputOrder, int nRequestID)
{
    return fn(this, pInputOrder, nRequestID);
}

int CallReqOrderAction(ReqOrderAction fn, void *this, struct CThostFtdcInputOrderActionField* pInputOrderAction, int nRequestID)
{
    return fn(this, pInputOrderAction, nRequestID);
}

int CallReqMKBatchOrderAction(ReqMKBatchOrderAction fn, void *this, struct CThostFtdcMKInputOrderActionField* pMKInputOrderAction, int nRequestID)
{
    return fn(this, pMKInputOrderAction, nRequestID);
}

int CallReqExecOrderInsert(ReqExecOrderInsert fn, void *this, struct CThostFtdcInputExecOrderField* pInputExecOrder, int nRequestID)
{
    return fn(this, pInputExecOrder, nRequestID);
}

int CallReqExecOrderAction(ReqExecOrderAction fn, void *this, struct CThostFtdcInputExecOrderActionField* pInputExecOrderAction, int nRequestID)
{
    return fn(this, pInputExecOrderAction, nRequestID);
}

int CallReqForQuoteInsert(ReqForQuoteInsert fn, void *this, struct CThostFtdcInputForQuoteField* pInputForQuote, int nRequestID)
{
    return fn(this, pInputForQuote, nRequestID);
}

int CallReqQuoteInsert(ReqQuoteInsert fn, void *this, struct CThostFtdcInputQuoteField* pInputQuote, int nRequestID)
{
    return fn(this, pInputQuote, nRequestID);
}

int CallReqQuoteAction(ReqQuoteAction fn, void *this, struct CThostFtdcInputQuoteActionField* pInputQuoteAction, int nRequestID)
{
    return fn(this, pInputQuoteAction, nRequestID);
}

int CallReqBatchOrderAction(ReqBatchOrderAction fn, void *this, struct CThostFtdcInputBatchOrderActionField* pInputBatchOrderAction, int nRequestID)
{
    return fn(this, pInputBatchOrderAction, nRequestID);
}

int CallReqOptionSelfCloseInsert(ReqOptionSelfCloseInsert fn, void *this, struct CThostFtdcInputOptionSelfCloseField* pInputOptionSelfClose, int nRequestID)
{
    return fn(this, pInputOptionSelfClose, nRequestID);
}

int CallReqOptionSelfCloseAction(ReqOptionSelfCloseAction fn, void *this, struct CThostFtdcInputOptionSelfCloseActionField* pInputOptionSelfCloseAction, int nRequestID)
{
    return fn(this, pInputOptionSelfCloseAction, nRequestID);
}

int CallReqCombActionInsert(ReqCombActionInsert fn, void *this, struct CThostFtdcInputCombActionField* pInputCombAction, int nRequestID)
{
    return fn(this, pInputCombAction, nRequestID);
}

int CallReqSubscribeFundChange(ReqSubscribeFundChange fn, void *this, int nRequestID)
{
    return fn(this, nRequestID);
}

int CallReqUnSubscribeFundChange(ReqUnSubscribeFundChange fn, void *this, int nRequestID)
{
    return fn(this, nRequestID);
}

int CallReqOffsetSetting(ReqOffsetSetting fn, void *this, struct CThostFtdcInputOffsetSettingField* pInputOffsetSetting, int nRequestID)
{
    return fn(this, pInputOffsetSetting, nRequestID);
}

int CallReqCancelOffsetSetting(ReqCancelOffsetSetting fn, void *this, struct CThostFtdcInputOffsetSettingField* pInputOffsetSetting, int nRequestID)
{
    return fn(this, pInputOffsetSetting, nRequestID);
}

int CallReqQryOrder(ReqQryOrder fn, void *this, struct CThostFtdcQryOrderField* pQryOrder, int nRequestID)
{
    return fn(this, pQryOrder, nRequestID);
}

int CallReqQryTrade(ReqQryTrade fn, void *this, struct CThostFtdcQryTradeField* pQryTrade, int nRequestID)
{
    return fn(this, pQryTrade, nRequestID);
}

int CallReqQryInvestorPosition(ReqQryInvestorPosition fn, void *this, struct CThostFtdcQryInvestorPositionField* pQryInvestorPosition, int nRequestID)
{
    return fn(this, pQryInvestorPosition, nRequestID);
}

int CallReqQryTradingAccount(ReqQryTradingAccount fn, void *this, struct CThostFtdcQryTradingAccountField* pQryTradingAccount, int nRequestID)
{
    return fn(this, pQryTradingAccount, nRequestID);
}

int CallReqQryInvestor(ReqQryInvestor fn, void *this, struct CThostFtdcQryInvestorField* pQryInvestor, int nRequestID)
{
    return fn(this, pQryInvestor, nRequestID);
}

int CallReqQryTradingCode(ReqQryTradingCode fn, void *this, struct CThostFtdcQryTradingCodeField* pQryTradingCode, int nRequestID)
{
    return fn(this, pQryTradingCode, nRequestID);
}

int CallReqQryInstrumentMarginRate(ReqQryInstrumentMarginRate fn, void *this, struct CThostFtdcQryInstrumentMarginRateField* pQryInstrumentMarginRate, int nRequestID)
{
    return fn(this, pQryInstrumentMarginRate, nRequestID);
}

int CallReqQryInstrumentCommissionRate(ReqQryInstrumentCommissionRate fn, void *this, struct CThostFtdcQryInstrumentCommissionRateField* pQryInstrumentCommissionRate, int nRequestID)
{
    return fn(this, pQryInstrumentCommissionRate, nRequestID);
}

int CallReqQryExchange(ReqQryExchange fn, void *this, struct CThostFtdcQryExchangeField* pQryExchange, int nRequestID)
{
    return fn(this, pQryExchange, nRequestID);
}

int CallReqQryProduct(ReqQryProduct fn, void *this, struct CThostFtdcQryProductField* pQryProduct, int nRequestID)
{
    return fn(this, pQryProduct, nRequestID);
}

int CallReqQryInstrument(ReqQryInstrument fn, void *this, struct CThostFtdcQryInstrumentField* pQryInstrument, int nRequestID)
{
    return fn(this, pQryInstrument, nRequestID);
}

int CallReqQryCombInstrument(ReqQryCombInstrument fn, void *this, struct CThostFtdcQryCombInstrumentField* pQryCombInstrument, int nRequestID)
{
    return fn(this, pQryCombInstrument, nRequestID);
}

int CallReqQryRCAMSInvestorProdMargin(ReqQryRCAMSInvestorProdMargin fn, void *this, struct CThostFtdcQryRCAMSInvestorProdMarginField* pQryRCAMSInvestorProdMargin, int nRequestID)
{
    return fn(this, pQryRCAMSInvestorProdMargin, nRequestID);
}

int CallReqQryRCAMSInvestorCombPosition(ReqQryRCAMSInvestorCombPosition fn, void *this, struct CThostFtdcQryRCAMSInvestorCombPositionField* pQryRCAMSInvestorCombPosition, int nRequestID)
{
    return fn(this, pQryRCAMSInvestorCombPosition, nRequestID);
}

int CallReqQryInvestorPositionForComb(ReqQryInvestorPositionForComb fn, void *this, struct CThostFtdcQryInvestorPositionForCombField* pQryIPForComb, int nRequestID)
{
    return fn(this, pQryIPForComb, nRequestID);
}

int CallReqQryCombAction(ReqQryCombAction fn, void *this, struct CThostFtdcQryCombActionField* pQryCombAction, int nRequestID)
{
    return fn(this, pQryCombAction, nRequestID);
}

int CallReqQryDepthMarketData(ReqQryDepthMarketData fn, void *this, struct CThostFtdcQryDepthMarketDataField* pQryDepthMarketData, int nRequestID)
{
    return fn(this, pQryDepthMarketData, nRequestID);
}

int CallReqQryOptionSelfClose(ReqQryOptionSelfClose fn, void *this, struct CThostFtdcQryOptionSelfCloseField* pQryOptionSelfClose, int nRequestID)
{
    return fn(this, pQryOptionSelfClose, nRequestID);
}

int CallReqQryInstrumentStatus(ReqQryInstrumentStatus fn, void *this, struct CThostFtdcQryInstrumentStatusField* pQryInstrumentStatus, int nRequestID)
{
    return fn(this, pQryInstrumentStatus, nRequestID);
}

int CallReqQryInvestorPositionDetail(ReqQryInvestorPositionDetail fn, void *this, struct CThostFtdcQryInvestorPositionDetailField* pQryInvestorPositionDetail, int nRequestID)
{
    return fn(this, pQryInvestorPositionDetail, nRequestID);
}

int CallReqQryExchangeMarginRate(ReqQryExchangeMarginRate fn, void *this, struct CThostFtdcQryExchangeMarginRateField* pQryExchangeMarginRate, int nRequestID)
{
    return fn(this, pQryExchangeMarginRate, nRequestID);
}

int CallReqQryExchangeMarginRateAdjust(ReqQryExchangeMarginRateAdjust fn, void *this, struct CThostFtdcQryExchangeMarginRateAdjustField* pQryExchangeMarginRateAdjust, int nRequestID)
{
    return fn(this, pQryExchangeMarginRateAdjust, nRequestID);
}

int CallReqQryOptionInstrTradeCost(ReqQryOptionInstrTradeCost fn, void *this, struct CThostFtdcQryOptionInstrTradeCostField* pQryOptionInstrTradeCost, int nRequestID)
{
    return fn(this, pQryOptionInstrTradeCost, nRequestID);
}

int CallReqQryOptionInstrCommRate(ReqQryOptionInstrCommRate fn, void *this, struct CThostFtdcQryOptionInstrCommRateField* pQryOptionInstrCommRate, int nRequestID)
{
    return fn(this, pQryOptionInstrCommRate, nRequestID);
}

int CallReqQryExecOrder(ReqQryExecOrder fn, void *this, struct CThostFtdcQryExecOrderField* pQryExecOrder, int nRequestID)
{
    return fn(this, pQryExecOrder, nRequestID);
}

int CallReqQryForQuote(ReqQryForQuote fn, void *this, struct CThostFtdcQryForQuoteField* pQryForQuote, int nRequestID)
{
    return fn(this, pQryForQuote, nRequestID);
}

int CallReqQryQuote(ReqQryQuote fn, void *this, struct CThostFtdcQryQuoteField* pQryQuote, int nRequestID)
{
    return fn(this, pQryQuote, nRequestID);
}

int CallReqQryInstrumentOrderCommRate(ReqQryInstrumentOrderCommRate fn, void *this, struct CThostFtdcQryInstrumentOrderCommRateField* pQryInstrumentOrderCommRate, int nRequestID)
{
    return fn(this, pQryInstrumentOrderCommRate, nRequestID);
}

int CallReqQryForQuoteParam(ReqQryForQuoteParam fn, void *this, struct CThostFtdcQryForQuoteParamField* pQryForQuoteParam, int nRequestID)
{
    return fn(this, pQryForQuoteParam, nRequestID);
}

int CallReqQryTraderOffer(ReqQryTraderOffer fn, void *this, struct CThostFtdcQryTraderOfferField* pQryTraderOffer, int nRequestID)
{
    return fn(this, pQryTraderOffer, nRequestID);
}

int CallReqQryInvestorProdSPBMDetail(ReqQryInvestorProdSPBMDetail fn, void *this, struct CThostFtdcQryInvestorProdSPBMDetailField* pQryInvestorProdSPBMDetail, int nRequestID)
{
    return fn(this, pQryInvestorProdSPBMDetail, nRequestID);
}

int CallReqQrySPMMInvestorCommodityGroupMargin(ReqQrySPMMInvestorCommodityGroupMargin fn, void *this, struct CThostFtdcQrySPMMInvestorCommodityGroupMarginField* pQrySPMMInvestorCommodityGroupMargin, int nRequestID)
{
    return fn(this, pQrySPMMInvestorCommodityGroupMargin, nRequestID);
}

int CallReqQryRULEInvestorProdMargin(ReqQryRULEInvestorProdMargin fn, void *this, struct CThostFtdcQryRULEInvestorProdMarginField* pQryRULEInvestorProdMargin, int nRequestID)
{
    return fn(this, pQryRULEInvestorProdMargin, nRequestID);
}

int CallReqQryControlParam(ReqQryControlParam fn, void *this, struct CThostFtdcQryControlParamField* pQryControlParam, int nRequestID)
{
    return fn(this, pQryControlParam, nRequestID);
}

int CallReqQryOffsetSetting(ReqQryOffsetSetting fn, void *this, struct CThostFtdcQryOffsetSettingField* pQryOffsetSetting, int nRequestID)
{
    return fn(this, pQryOffsetSetting, nRequestID);
}

int CallReqUserPasswordUpdate(ReqUserPasswordUpdate fn, void *this, struct CThostFtdcUserPasswordUpdateField* pUserPasswordUpdate, int nRequestID)
{
    return fn(this, pUserPasswordUpdate, nRequestID);
}
