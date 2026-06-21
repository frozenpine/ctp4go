#include "td_api_helper.h"

void CallRelease(Release fn, void *this)
{
    fn(this);
}

void CallInit(Init fn, void *this)
{
    fn(this);
}

int CallJoin(Join fn, void *this)
{
    return fn(this);
}

const char *CallGetTradingDay(GetTradingDay fn, void *this)
{
    return fn(this);
}

void CallGetFrontInfo(GetFrontInfo fn, void *this, struct CThostFtdcFrontInfoField *pFrontInfo)
{
    return fn(this, pFrontInfo);
}

void CallRegisterFront(RegisterFront fn, void *this, char *pszFrontAddress)
{
    return fn(this, pszFrontAddress);
}

void CallRegisterNameServer(RegisterNameServer fn, void *this, char *pszNsAddress)
{
    return fn(this, pszNsAddress);
}

void CallRegisterFensUserInfo(RegisterFensUserInfo fn, void *this, struct CThostFtdcFensUserInfoField *pFensUserInfo)
{
    return fn(this, pFensUserInfo);
}

void CallRegisterSpi(RegisterSpi fn, void *this, void *pSpi)
{
    return fn(this, pSpi);
}

void CallSubscribePrivateTopic(SubscribePrivateTopic fn, void *this, THOST_TE_RESUME_TYPE nResumeType, int nSeqNo)
{
    return fn(this, nResumeType, nSeqNo);
}

void CallSubscribePublicTopic(SubscribePublicTopic fn, void *this, THOST_TE_RESUME_TYPE nResumeType)
{
    return fn(this, nResumeType);
}

int CallReqAuthenticate(ReqAuthenticate fn, void *this, struct CThostFtdcReqAuthenticateField *pReqAuthenticateField, int nRequestID)
{
    return fn(this, pReqAuthenticateField, nRequestID);
}

int CallRegisterUserSystemInfo(RegisterUserSystemInfo fn, void *this, struct CThostFtdcUserSystemInfoField *pUserSystemInfo)
{
    return fn(this, pUserSystemInfo);
}

int CallSubmitUserSystemInfo(SubmitUserSystemInfo fn, void *this, struct CThostFtdcUserSystemInfoField *pUserSystemInfo)
{
    return fn(this, pUserSystemInfo);
}

int CallRegisterWechatUserSystemInfo(RegisterWechatUserSystemInfo fn, void *this, struct CThostFtdcWechatUserSystemInfoField *pUserSystemInfo)
{
    return fn(this, pUserSystemInfo);
}

int CallSubmitWechatUserSystemInfo(SubmitWechatUserSystemInfo fn, void *this, struct CThostFtdcWechatUserSystemInfoField *pUserSystemInfo)
{
    return fn(this, pUserSystemInfo);
}

int CallReqUserLogin(ReqUserLogin fn, void *this, struct CThostFtdcReqUserLoginField *pReqUserLoginField, int nRequestID)
{
    return fn(this, pReqUserLoginField, nRequestID);
}

int CallReqUserLogout(ReqUserLogout fn, void *this, struct CThostFtdcUserLogoutField *pUserLogout, int nRequestID)
{
    return fn(this, pUserLogout, nRequestID);
}

int CallReqUserPasswordUpdate(ReqUserPasswordUpdate fn, void *this, struct CThostFtdcUserPasswordUpdateField *pUserPasswordUpdate, int nRequestID)
{
    return fn(this, pUserPasswordUpdate, nRequestID);
}

int CallReqTradingAccountPasswordUpdate(ReqTradingAccountPasswordUpdate fn, void *this, struct CThostFtdcTradingAccountPasswordUpdateField *pTradingAccountPasswordUpdate, int nRequestID)
{
    return fn(this, pTradingAccountPasswordUpdate, nRequestID);
}

int CallReqUserAuthMethod(ReqUserAuthMethod fn, void *this, struct CThostFtdcReqUserAuthMethodField *pReqUserAuthMethod, int nRequestID)
{
    return fn(this, pReqUserAuthMethod, nRequestID);
}

int CallReqGenUserCaptcha(ReqGenUserCaptcha fn, void *this, struct CThostFtdcReqGenUserCaptchaField *pReqGenUserCaptcha, int nRequestID)
{
    return fn(this, pReqGenUserCaptcha, nRequestID);
}

int CallReqGenUserText(ReqGenUserText fn, void *this, struct CThostFtdcReqGenUserTextField *pReqGenUserText, int nRequestID)
{
    return fn(this, pReqGenUserText, nRequestID);
}

int CallReqUserLoginWithCaptcha(ReqUserLoginWithCaptcha fn, void *this, struct CThostFtdcReqUserLoginWithCaptchaField *pReqUserLoginWithCaptcha, int nRequestID)
{
    return fn(this, pReqUserLoginWithCaptcha, nRequestID);
}

int CallReqUserLoginWithText(ReqUserLoginWithText fn, void *this, struct CThostFtdcReqUserLoginWithTextField *pReqUserLoginWithText, int nRequestID)
{
    return fn(this, pReqUserLoginWithText, nRequestID);
}

int CallReqUserLoginWithOTP(ReqUserLoginWithOTP fn, void *this, struct CThostFtdcReqUserLoginWithOTPField *pReqUserLoginWithOTP, int nRequestID)
{
    return fn(this, pReqUserLoginWithOTP, nRequestID);
}

int CallReqOrderInsert(ReqOrderInsert fn, void *this, struct CThostFtdcInputOrderField *pInputOrder, int nRequestID)
{
    return fn(this, pInputOrder, nRequestID);
}

int CallReqParkedOrderInsert(ReqParkedOrderInsert fn, void *this, struct CThostFtdcParkedOrderField *pParkedOrder, int nRequestID)
{
    return fn(this, pParkedOrder, nRequestID);
}

int CallReqParkedOrderAction(ReqParkedOrderAction fn, void *this, struct CThostFtdcParkedOrderActionField *pParkedOrderAction, int nRequestID)
{
    return fn(this, pParkedOrderAction, nRequestID);
}

int CallReqOrderAction(ReqOrderAction fn, void *this, struct CThostFtdcInputOrderActionField *pInputOrderAction, int nRequestID)
{
    return fn(this, pInputOrderAction, nRequestID);
}

int CallReqQryMaxOrderVolume(ReqQryMaxOrderVolume fn, void *this, struct CThostFtdcQryMaxOrderVolumeField *pQryMaxOrderVolume, int nRequestID)
{
    return fn(this, pQryMaxOrderVolume, nRequestID);
}

int CallReqSettlementInfoConfirm(ReqSettlementInfoConfirm fn, void *this, struct CThostFtdcSettlementInfoConfirmField *pSettlementInfoConfirm, int nRequestID)
{
    return fn(this, pSettlementInfoConfirm, nRequestID);
}

int CallReqRemoveParkedOrder(ReqRemoveParkedOrder fn, void *this, struct CThostFtdcRemoveParkedOrderField *pRemoveParkedOrder, int nRequestID)
{
    return fn(this, pRemoveParkedOrder, nRequestID);
}

int CallReqRemoveParkedOrderAction(ReqRemoveParkedOrderAction fn, void *this, struct CThostFtdcRemoveParkedOrderActionField *pRemoveParkedOrderAction, int nRequestID)
{
    return fn(this, pRemoveParkedOrderAction, nRequestID);
}

int CallReqExecOrderInsert(ReqExecOrderInsert fn, void *this, struct CThostFtdcInputExecOrderField *pInputExecOrder, int nRequestID)
{
    return fn(this, pInputExecOrder, nRequestID);
}

int CallReqExecOrderAction(ReqExecOrderAction fn, void *this, struct CThostFtdcInputExecOrderActionField *pInputExecOrderAction, int nRequestID)
{
    return fn(this, pInputExecOrderAction, nRequestID);
}

int CallReqForQuoteInsert(ReqForQuoteInsert fn, void *this, struct CThostFtdcInputForQuoteField *pInputForQuote, int nRequestID)
{
    return fn(this, pInputForQuote, nRequestID);
}

int CallReqQuoteInsert(ReqQuoteInsert fn, void *this, struct CThostFtdcInputQuoteField *pInputQuote, int nRequestID)
{
    return fn(this, pInputQuote, nRequestID);
}

int CallReqQuoteAction(ReqQuoteAction fn, void *this, struct CThostFtdcInputQuoteActionField *pInputQuoteAction, int nRequestID)
{
    return fn(this, pInputQuoteAction, nRequestID);
}

int CallReqBatchOrderAction(ReqBatchOrderAction fn, void *this, struct CThostFtdcInputBatchOrderActionField *pInputBatchOrderAction, int nRequestID)
{
    return fn(this, pInputBatchOrderAction, nRequestID);
}

int CallReqOptionSelfCloseInsert(ReqOptionSelfCloseInsert fn, void *this, struct CThostFtdcInputOptionSelfCloseField *pInputOptionSelfClose, int nRequestID)
{
    return fn(this, pInputOptionSelfClose, nRequestID);
}

int CallReqOptionSelfCloseAction(ReqOptionSelfCloseAction fn, void *this, struct CThostFtdcInputOptionSelfCloseActionField *pInputOptionSelfCloseAction, int nRequestID)
{
    return fn(this, pInputOptionSelfCloseAction, nRequestID);
}

int CallReqCombActionInsert(ReqCombActionInsert fn, void *this, struct CThostFtdcInputCombActionField *pInputCombAction, int nRequestID)
{
    return fn(this, pInputCombAction, nRequestID);
}

int CallReqQryOrder(ReqQryOrder fn, void *this, struct CThostFtdcQryOrderField *pQryOrder, int nRequestID)
{
    return fn(this, pQryOrder, nRequestID);
}

int CallReqQryTrade(ReqQryTrade fn, void *this, struct CThostFtdcQryTradeField *pQryTrade, int nRequestID)
{
    return fn(this, pQryTrade, nRequestID);
}

int CallReqQryInvestorPosition(ReqQryInvestorPosition fn, void *this, struct CThostFtdcQryInvestorPositionField *pQryInvestorPosition, int nRequestID)
{
    return fn(this, pQryInvestorPosition, nRequestID);
}

int CallReqQryTradingAccount(ReqQryTradingAccount fn, void *this, struct CThostFtdcQryTradingAccountField *pQryTradingAccount, int nRequestID)
{
    return fn(this, pQryTradingAccount, nRequestID);
}

int CallReqQryInvestor(ReqQryInvestor fn, void *this, struct CThostFtdcQryInvestorField *pQryInvestor, int nRequestID)
{
    return fn(this, pQryInvestor, nRequestID);
}

int CallReqQryTradingCode(ReqQryTradingCode fn, void *this, struct CThostFtdcQryTradingCodeField *pQryTradingCode, int nRequestID)
{
    return fn(this, pQryTradingCode, nRequestID);
}

int CallReqQryInstrumentMarginRate(ReqQryInstrumentMarginRate fn, void *this, struct CThostFtdcQryInstrumentMarginRateField *pQryInstrumentMarginRate, int nRequestID)
{
    return fn(this, pQryInstrumentMarginRate, nRequestID);
}

int CallReqQryInstrumentCommissionRate(ReqQryInstrumentCommissionRate fn, void *this, struct CThostFtdcQryInstrumentCommissionRateField *pQryInstrumentCommissionRate, int nRequestID)
{
    return fn(this, pQryInstrumentCommissionRate, nRequestID);
}

int CallReqQryUserSession(ReqQryUserSession fn, void *this, struct CThostFtdcQryUserSessionField *pQryUserSession, int nRequestID)
{
    return fn(this, pQryUserSession, nRequestID);
}

int CallReqQryExchange(ReqQryExchange fn, void *this, struct CThostFtdcQryExchangeField *pQryExchange, int nRequestID)
{
    return fn(this, pQryExchange, nRequestID);
}

int CallReqQryProduct(ReqQryProduct fn, void *this, struct CThostFtdcQryProductField *pQryProduct, int nRequestID)
{
    return fn(this, pQryProduct, nRequestID);
}

int CallReqQryInstrument(ReqQryInstrument fn, void *this, struct CThostFtdcQryInstrumentField *pQryInstrument, int nRequestID)
{
    return fn(this, pQryInstrument, nRequestID);
}

int CallReqQryDepthMarketData(ReqQryDepthMarketData fn, void *this, struct CThostFtdcQryDepthMarketDataField *pQryDepthMarketData, int nRequestID)
{
    return fn(this, pQryDepthMarketData, nRequestID);
}

int CallReqQryTraderOffer(ReqQryTraderOffer fn, void *this, struct CThostFtdcQryTraderOfferField *pQryTraderOffer, int nRequestID)
{
    return fn(this, pQryTraderOffer, nRequestID);
}

int CallReqQrySettlementInfo(ReqQrySettlementInfo fn, void *this, struct CThostFtdcQrySettlementInfoField *pQrySettlementInfo, int nRequestID)
{
    return fn(this, pQrySettlementInfo, nRequestID);
}

int CallReqQryTransferBank(ReqQryTransferBank fn, void *this, struct CThostFtdcQryTransferBankField *pQryTransferBank, int nRequestID)
{
    return fn(this, pQryTransferBank, nRequestID);
}

int CallReqQryInvestorPositionDetail(ReqQryInvestorPositionDetail fn, void *this, struct CThostFtdcQryInvestorPositionDetailField *pQryInvestorPositionDetail, int nRequestID)
{
    return fn(this, pQryInvestorPositionDetail, nRequestID);
}

int CallReqQryNotice(ReqQryNotice fn, void *this, struct CThostFtdcQryNoticeField *pQryNotice, int nRequestID)
{
    return fn(this, pQryNotice, nRequestID);
}

int CallReqQrySettlementInfoConfirm(ReqQrySettlementInfoConfirm fn, void *this, struct CThostFtdcQrySettlementInfoConfirmField *pQrySettlementInfoConfirm, int nRequestID)
{
    return fn(this, pQrySettlementInfoConfirm, nRequestID);
}

int CallReqQryInvestorPositionCombineDetail(ReqQryInvestorPositionCombineDetail fn, void *this, struct CThostFtdcQryInvestorPositionCombineDetailField *pQryInvestorPositionCombineDetail, int nRequestID)
{
    return fn(this, pQryInvestorPositionCombineDetail, nRequestID);
}

int CallReqQryCFMMCTradingAccountKey(ReqQryCFMMCTradingAccountKey fn, void *this, struct CThostFtdcQryCFMMCTradingAccountKeyField *pQryCFMMCTradingAccountKey, int nRequestID)
{
    return fn(this, pQryCFMMCTradingAccountKey, nRequestID);
}

int CallReqQryEWarrantOffset(ReqQryEWarrantOffset fn, void *this, struct CThostFtdcQryEWarrantOffsetField *pQryEWarrantOffset, int nRequestID)
{
    return fn(this, pQryEWarrantOffset, nRequestID);
}

int CallReqQryInvestorProductGroupMargin(ReqQryInvestorProductGroupMargin fn, void *this, struct CThostFtdcQryInvestorProductGroupMarginField *pQryInvestorProductGroupMargin, int nRequestID)
{
    return fn(this, pQryInvestorProductGroupMargin, nRequestID);
}

int CallReqQryExchangeMarginRate(ReqQryExchangeMarginRate fn, void *this, struct CThostFtdcQryExchangeMarginRateField *pQryExchangeMarginRate, int nRequestID)
{
    return fn(this, pQryExchangeMarginRate, nRequestID);
}

int CallReqQryExchangeMarginRateAdjust(ReqQryExchangeMarginRateAdjust fn, void *this, struct CThostFtdcQryExchangeMarginRateAdjustField *pQryExchangeMarginRateAdjust, int nRequestID)
{
    return fn(this, pQryExchangeMarginRateAdjust, nRequestID);
}

int CallReqQryExchangeRate(ReqQryExchangeRate fn, void *this, struct CThostFtdcQryExchangeRateField *pQryExchangeRate, int nRequestID)
{
    return fn(this, pQryExchangeRate, nRequestID);
}

int CallReqQrySecAgentACIDMap(ReqQrySecAgentACIDMap fn, void *this, struct CThostFtdcQrySecAgentACIDMapField *pQrySecAgentACIDMap, int nRequestID)
{
    return fn(this, pQrySecAgentACIDMap, nRequestID);
}

int CallReqQryProductExchRate(ReqQryProductExchRate fn, void *this, struct CThostFtdcQryProductExchRateField *pQryProductExchRate, int nRequestID)
{
    return fn(this, pQryProductExchRate, nRequestID);
}

int CallReqQryProductGroup(ReqQryProductGroup fn, void *this, struct CThostFtdcQryProductGroupField *pQryProductGroup, int nRequestID)
{
    return fn(this, pQryProductGroup, nRequestID);
}

int CallReqQryMMInstrumentCommissionRate(ReqQryMMInstrumentCommissionRate fn, void *this, struct CThostFtdcQryMMInstrumentCommissionRateField *pQryMMInstrumentCommissionRate, int nRequestID)
{
    return fn(this, pQryMMInstrumentCommissionRate, nRequestID);
}

int CallReqQryMMOptionInstrCommRate(ReqQryMMOptionInstrCommRate fn, void *this, struct CThostFtdcQryMMOptionInstrCommRateField *pQryMMOptionInstrCommRate, int nRequestID)
{
    return fn(this, pQryMMOptionInstrCommRate, nRequestID);
}

int CallReqQryInstrumentOrderCommRate(ReqQryInstrumentOrderCommRate fn, void *this, struct CThostFtdcQryInstrumentOrderCommRateField *pQryInstrumentOrderCommRate, int nRequestID)
{
    return fn(this, pQryInstrumentOrderCommRate, nRequestID);
}

int CallReqQrySecAgentTradingAccount(ReqQrySecAgentTradingAccount fn, void *this, struct CThostFtdcQryTradingAccountField *pQryTradingAccount, int nRequestID)
{
    return fn(this, pQryTradingAccount, nRequestID);
}

int CallReqQrySecAgentCheckMode(ReqQrySecAgentCheckMode fn, void *this, struct CThostFtdcQrySecAgentCheckModeField *pQrySecAgentCheckMode, int nRequestID)
{
    return fn(this, pQrySecAgentCheckMode, nRequestID);
}

int CallReqQrySecAgentTradeInfo(ReqQrySecAgentTradeInfo fn, void *this, struct CThostFtdcQrySecAgentTradeInfoField *pQrySecAgentTradeInfo, int nRequestID)
{
    return fn(this, pQrySecAgentTradeInfo, nRequestID);
}

int CallReqQryOptionInstrTradeCost(ReqQryOptionInstrTradeCost fn, void *this, struct CThostFtdcQryOptionInstrTradeCostField *pQryOptionInstrTradeCost, int nRequestID)
{
    return fn(this, pQryOptionInstrTradeCost, nRequestID);
}

int CallReqQryOptionInstrCommRate(ReqQryOptionInstrCommRate fn, void *this, struct CThostFtdcQryOptionInstrCommRateField *pQryOptionInstrCommRate, int nRequestID)
{
    return fn(this, pQryOptionInstrCommRate, nRequestID);
}

int CallReqQryExecOrder(ReqQryExecOrder fn, void *this, struct CThostFtdcQryExecOrderField *pQryExecOrder, int nRequestID)
{
    return fn(this, pQryExecOrder, nRequestID);
}

int CallReqQryForQuote(ReqQryForQuote fn, void *this, struct CThostFtdcQryForQuoteField *pQryForQuote, int nRequestID)
{
    return fn(this, pQryForQuote, nRequestID);
}

int CallReqQryQuote(ReqQryQuote fn, void *this, struct CThostFtdcQryQuoteField *pQryQuote, int nRequestID)
{
    return fn(this, pQryQuote, nRequestID);
}

int CallReqQryOptionSelfClose(ReqQryOptionSelfClose fn, void *this, struct CThostFtdcQryOptionSelfCloseField *pQryOptionSelfClose, int nRequestID)
{
    return fn(this, pQryOptionSelfClose, nRequestID);
}

int CallReqQryInvestUnit(ReqQryInvestUnit fn, void *this, struct CThostFtdcQryInvestUnitField *pQryInvestUnit, int nRequestID)
{
    return fn(this, pQryInvestUnit, nRequestID);
}

int CallReqQryCombInstrumentGuard(ReqQryCombInstrumentGuard fn, void *this, struct CThostFtdcQryCombInstrumentGuardField *pQryCombInstrumentGuard, int nRequestID)
{
    return fn(this, pQryCombInstrumentGuard, nRequestID);
}

int CallReqQryCombAction(ReqQryCombAction fn, void *this, struct CThostFtdcQryCombActionField *pQryCombAction, int nRequestID)
{
    return fn(this, pQryCombAction, nRequestID);
}

int CallReqQryTransferSerial(ReqQryTransferSerial fn, void *this, struct CThostFtdcQryTransferSerialField *pQryTransferSerial, int nRequestID)
{
    return fn(this, pQryTransferSerial, nRequestID);
}

int CallReqQryAccountregister(ReqQryAccountregister fn, void *this, struct CThostFtdcQryAccountregisterField *pQryAccountregister, int nRequestID)
{
    return fn(this, pQryAccountregister, nRequestID);
}

int CallReqQryContractBank(ReqQryContractBank fn, void *this, struct CThostFtdcQryContractBankField *pQryContractBank, int nRequestID)
{
    return fn(this, pQryContractBank, nRequestID);
}

int CallReqQryParkedOrder(ReqQryParkedOrder fn, void *this, struct CThostFtdcQryParkedOrderField *pQryParkedOrder, int nRequestID)
{
    return fn(this, pQryParkedOrder, nRequestID);
}

int CallReqQryParkedOrderAction(ReqQryParkedOrderAction fn, void *this, struct CThostFtdcQryParkedOrderActionField *pQryParkedOrderAction, int nRequestID)
{
    return fn(this, pQryParkedOrderAction, nRequestID);
}

int CallReqQryTradingNotice(ReqQryTradingNotice fn, void *this, struct CThostFtdcQryTradingNoticeField *pQryTradingNotice, int nRequestID)
{
    return fn(this, pQryTradingNotice, nRequestID);
}

int CallReqQryBrokerTradingParams(ReqQryBrokerTradingParams fn, void *this, struct CThostFtdcQryBrokerTradingParamsField *pQryBrokerTradingParams, int nRequestID)
{
    return fn(this, pQryBrokerTradingParams, nRequestID);
}

int CallReqQryBrokerTradingAlgos(ReqQryBrokerTradingAlgos fn, void *this, struct CThostFtdcQryBrokerTradingAlgosField *pQryBrokerTradingAlgos, int nRequestID)
{
    return fn(this, pQryBrokerTradingAlgos, nRequestID);
}

int CallReqQueryCFMMCTradingAccountToken(ReqQueryCFMMCTradingAccountToken fn, void *this, struct CThostFtdcQueryCFMMCTradingAccountTokenField *pQueryCFMMCTradingAccountToken, int nRequestID)
{
    return fn(this, pQueryCFMMCTradingAccountToken, nRequestID);
}

int CallReqFromBankToFutureByFuture(ReqFromBankToFutureByFuture fn, void *this, struct CThostFtdcReqTransferField *pReqTransfer, int nRequestID)
{
    return fn(this, pReqTransfer, nRequestID);
}

int CallReqFromFutureToBankByFuture(ReqFromFutureToBankByFuture fn, void *this, struct CThostFtdcReqTransferField *pReqTransfer, int nRequestID)
{
    return fn(this, pReqTransfer, nRequestID);
}

int CallReqQueryBankAccountMoneyByFuture(ReqQueryBankAccountMoneyByFuture fn, void *this, struct CThostFtdcReqQueryAccountField *pReqQueryAccount, int nRequestID)
{
    return fn(this, pReqQueryAccount, nRequestID);
}

int CallReqQryClassifiedInstrument(ReqQryClassifiedInstrument fn, void *this, struct CThostFtdcQryClassifiedInstrumentField *pQryClassifiedInstrument, int nRequestID)
{
    return fn(this, pQryClassifiedInstrument, nRequestID);
}

int CallReqQryCombPromotionParam(ReqQryCombPromotionParam fn, void *this, struct CThostFtdcQryCombPromotionParamField *pQryCombPromotionParam, int nRequestID)
{
    return fn(this, pQryCombPromotionParam, nRequestID);
}

int CallReqQryRiskSettleInvstPosition(ReqQryRiskSettleInvstPosition fn, void *this, struct CThostFtdcQryRiskSettleInvstPositionField *pQryRiskSettleInvstPosition, int nRequestID)
{
    return fn(this, pQryRiskSettleInvstPosition, nRequestID);
}

int CallReqQryRiskSettleProductStatus(ReqQryRiskSettleProductStatus fn, void *this, struct CThostFtdcQryRiskSettleProductStatusField *pQryRiskSettleProductStatus, int nRequestID)
{
    return fn(this, pQryRiskSettleProductStatus, nRequestID);
}

int CallReqQrySPBMFutureParameter(ReqQrySPBMFutureParameter fn, void *this, struct CThostFtdcQrySPBMFutureParameterField *pQrySPBMFutureParameter, int nRequestID)
{
    return fn(this, pQrySPBMFutureParameter, nRequestID);
}

int CallReqQrySPBMOptionParameter(ReqQrySPBMOptionParameter fn, void *this, struct CThostFtdcQrySPBMOptionParameterField *pQrySPBMOptionParameter, int nRequestID)
{
    return fn(this, pQrySPBMOptionParameter, nRequestID);
}

int CallReqQrySPBMIntraParameter(ReqQrySPBMIntraParameter fn, void *this, struct CThostFtdcQrySPBMIntraParameterField *pQrySPBMIntraParameter, int nRequestID)
{
    return fn(this, pQrySPBMIntraParameter, nRequestID);
}

int CallReqQrySPBMInterParameter(ReqQrySPBMInterParameter fn, void *this, struct CThostFtdcQrySPBMInterParameterField *pQrySPBMInterParameter, int nRequestID)
{
    return fn(this, pQrySPBMInterParameter, nRequestID);
}

int CallReqQrySPBMPortfDefinition(ReqQrySPBMPortfDefinition fn, void *this, struct CThostFtdcQrySPBMPortfDefinitionField *pQrySPBMPortfDefinition, int nRequestID)
{
    return fn(this, pQrySPBMPortfDefinition, nRequestID);
}

int CallReqQrySPBMInvestorPortfDef(ReqQrySPBMInvestorPortfDef fn, void *this, struct CThostFtdcQrySPBMInvestorPortfDefField *pQrySPBMInvestorPortfDef, int nRequestID)
{
    return fn(this, pQrySPBMInvestorPortfDef, nRequestID);
}

int CallReqQryInvestorPortfMarginRatio(ReqQryInvestorPortfMarginRatio fn, void *this, struct CThostFtdcQryInvestorPortfMarginRatioField *pQryInvestorPortfMarginRatio, int nRequestID)
{
    return fn(this, pQryInvestorPortfMarginRatio, nRequestID);
}

int CallReqQryInvestorProdSPBMDetail(ReqQryInvestorProdSPBMDetail fn, void *this, struct CThostFtdcQryInvestorProdSPBMDetailField *pQryInvestorProdSPBMDetail, int nRequestID)
{
    return fn(this, pQryInvestorProdSPBMDetail, nRequestID);
}

int CallReqQryInvestorCommoditySPMMMargin(ReqQryInvestorCommoditySPMMMargin fn, void *this, struct CThostFtdcQryInvestorCommoditySPMMMarginField *pQryInvestorCommoditySPMMMargin, int nRequestID)
{
    return fn(this, pQryInvestorCommoditySPMMMargin, nRequestID);
}

int CallReqQryInvestorCommodityGroupSPMMMargin(ReqQryInvestorCommodityGroupSPMMMargin fn, void *this, struct CThostFtdcQryInvestorCommodityGroupSPMMMarginField *pQryInvestorCommodityGroupSPMMMargin, int nRequestID)
{
    return fn(this, pQryInvestorCommodityGroupSPMMMargin, nRequestID);
}

int CallReqQrySPMMInstParam(ReqQrySPMMInstParam fn, void *this, struct CThostFtdcQrySPMMInstParamField *pQrySPMMInstParam, int nRequestID)
{
    return fn(this, pQrySPMMInstParam, nRequestID);
}

int CallReqQrySPMMProductParam(ReqQrySPMMProductParam fn, void *this, struct CThostFtdcQrySPMMProductParamField *pQrySPMMProductParam, int nRequestID)
{
    return fn(this, pQrySPMMProductParam, nRequestID);
}

int CallReqQrySPBMAddOnInterParameter(ReqQrySPBMAddOnInterParameter fn, void *this, struct CThostFtdcQrySPBMAddOnInterParameterField *pQrySPBMAddOnInterParameter, int nRequestID)
{
    return fn(this, pQrySPBMAddOnInterParameter, nRequestID);
}

int CallReqQryRCAMSCombProductInfo(ReqQryRCAMSCombProductInfo fn, void *this, struct CThostFtdcQryRCAMSCombProductInfoField *pQryRCAMSCombProductInfo, int nRequestID)
{
    return fn(this, pQryRCAMSCombProductInfo, nRequestID);
}

int CallReqQryRCAMSInstrParameter(ReqQryRCAMSInstrParameter fn, void *this, struct CThostFtdcQryRCAMSInstrParameterField *pQryRCAMSInstrParameter, int nRequestID)
{
    return fn(this, pQryRCAMSInstrParameter, nRequestID);
}

int CallReqQryRCAMSIntraParameter(ReqQryRCAMSIntraParameter fn, void *this, struct CThostFtdcQryRCAMSIntraParameterField *pQryRCAMSIntraParameter, int nRequestID)
{
    return fn(this, pQryRCAMSIntraParameter, nRequestID);
}

int CallReqQryRCAMSInterParameter(ReqQryRCAMSInterParameter fn, void *this, struct CThostFtdcQryRCAMSInterParameterField *pQryRCAMSInterParameter, int nRequestID)
{
    return fn(this, pQryRCAMSInterParameter, nRequestID);
}

int CallReqQryRCAMSShortOptAdjustParam(ReqQryRCAMSShortOptAdjustParam fn, void *this, struct CThostFtdcQryRCAMSShortOptAdjustParamField *pQryRCAMSShortOptAdjustParam, int nRequestID)
{
    return fn(this, pQryRCAMSShortOptAdjustParam, nRequestID);
}

int CallReqQryRCAMSInvestorCombPosition(ReqQryRCAMSInvestorCombPosition fn, void *this, struct CThostFtdcQryRCAMSInvestorCombPositionField *pQryRCAMSInvestorCombPosition, int nRequestID)
{
    return fn(this, pQryRCAMSInvestorCombPosition, nRequestID);
}

int CallReqQryInvestorProdRCAMSMargin(ReqQryInvestorProdRCAMSMargin fn, void *this, struct CThostFtdcQryInvestorProdRCAMSMarginField *pQryInvestorProdRCAMSMargin, int nRequestID)
{
    return fn(this, pQryInvestorProdRCAMSMargin, nRequestID);
}

int CallReqQryRULEInstrParameter(ReqQryRULEInstrParameter fn, void *this, struct CThostFtdcQryRULEInstrParameterField *pQryRULEInstrParameter, int nRequestID)
{
    return fn(this, pQryRULEInstrParameter, nRequestID);
}

int CallReqQryRULEIntraParameter(ReqQryRULEIntraParameter fn, void *this, struct CThostFtdcQryRULEIntraParameterField *pQryRULEIntraParameter, int nRequestID)
{
    return fn(this, pQryRULEIntraParameter, nRequestID);
}

int CallReqQryRULEInterParameter(ReqQryRULEInterParameter fn, void *this, struct CThostFtdcQryRULEInterParameterField *pQryRULEInterParameter, int nRequestID)
{
    return fn(this, pQryRULEInterParameter, nRequestID);
}

int CallReqQryInvestorProdRULEMargin(ReqQryInvestorProdRULEMargin fn, void *this, struct CThostFtdcQryInvestorProdRULEMarginField *pQryInvestorProdRULEMargin, int nRequestID)
{
    return fn(this, pQryInvestorProdRULEMargin, nRequestID);
}

int CallReqQryInvestorPortfSetting(ReqQryInvestorPortfSetting fn, void *this, struct CThostFtdcQryInvestorPortfSettingField *pQryInvestorPortfSetting, int nRequestID)
{
    return fn(this, pQryInvestorPortfSetting, nRequestID);
}

int CallReqQryInvestorInfoCommRec(ReqQryInvestorInfoCommRec fn, void *this, struct CThostFtdcQryInvestorInfoCommRecField *pQryInvestorInfoCommRec, int nRequestID)
{
    return fn(this, pQryInvestorInfoCommRec, nRequestID);
}

int CallReqQryCombLeg(ReqQryCombLeg fn, void *this, struct CThostFtdcQryCombLegField *pQryCombLeg, int nRequestID)
{
    return fn(this, pQryCombLeg, nRequestID);
}

int CallReqOffsetSetting(ReqOffsetSetting fn, void *this, struct CThostFtdcInputOffsetSettingField *pInputOffsetSetting, int nRequestID)
{
    return fn(this, pInputOffsetSetting, nRequestID);
}

int CallReqCancelOffsetSetting(ReqCancelOffsetSetting fn, void *this, struct CThostFtdcInputOffsetSettingField *pInputOffsetSetting, int nRequestID)
{
    return fn(this, pInputOffsetSetting, nRequestID);
}

int CallReqQryOffsetSetting(ReqQryOffsetSetting fn, void *this, struct CThostFtdcQryOffsetSettingField *pQryOffsetSetting, int nRequestID)
{
    return fn(this, pQryOffsetSetting, nRequestID);
}

int CallReqGenSMSCode(ReqGenSMSCode fn, void *this, struct CThostFtdcReqGenSMSCodeField *pReqGenSMSCode, int nRequestID)
{
    return fn(this, pReqGenSMSCode, nRequestID);
}

int CallReqSpdApply(ReqSpdApply fn, void *this, struct CThostFtdcInputSpdApplyField *pInputSpdApply, int nRequestID)
{
    return fn(this, pInputSpdApply, nRequestID);
}

int CallReqSpdApplyAction(ReqSpdApplyAction fn, void *this, struct CThostFtdcInputSpdApplyActionField *pInputSpdApplyAction, int nRequestID)
{
    return fn(this, pInputSpdApplyAction, nRequestID);
}

int CallReqQrySpdApply(ReqQrySpdApply fn, void *this, struct CThostFtdcQrySpdApplyField *pQrySpdApply, int nRequestID)
{
    return fn(this, pQrySpdApply, nRequestID);
}

int CallReqHedgeCfm(ReqHedgeCfm fn, void *this, struct CThostFtdcInputHdegeCfmField *pInputHedgeCfm, int nRequestID)
{
    return fn(this, pInputHedgeCfm, nRequestID);
}

int CallReqHedgeCfmAction(ReqHedgeCfmAction fn, void *this, struct CThostFtdcInputHedgeCfmActionField *pInputHedgeCfmAction, int nRequestID)
{
    return fn(this, pInputHedgeCfmAction, nRequestID);
}

int CallReqQryHedgeCfm(ReqQryHedgeCfm fn, void *this, struct CThostFtdcQryHedgeCfmField *pQryHedgeCfm, int nRequestID)
{
    return fn(this, pQryHedgeCfm, nRequestID);
}
