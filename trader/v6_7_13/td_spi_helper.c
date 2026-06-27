#include "td_spi_helper.h"

extern void CgoOnFrontConnected(void *this);

extern void CgoOnFrontDisconnected(void *this, int nReason);

extern void CgoOnHeartBeatWarning(void *this, int nTimeLapse);

extern void CgoOnRspAuthenticate(void *this, struct CThostFtdcRspAuthenticateField *pRspAuthenticateField, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

extern void CgoOnRtnPrivateSeqNo(void *this, int nSeqNo);

extern void CgoOnRspUserLogin(void *this, struct CThostFtdcRspUserLoginField *pRspUserLogin, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

extern void CgoOnRspUserLogout(void *this, struct CThostFtdcUserLogoutField *pUserLogout, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

extern void CgoOnRspUserPasswordUpdate(void *this, struct CThostFtdcUserPasswordUpdateField *pUserPasswordUpdate, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

extern void CgoOnRspTradingAccountPasswordUpdate(void *this, struct CThostFtdcTradingAccountPasswordUpdateField *pTradingAccountPasswordUpdate, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

extern void CgoOnRspUserAuthMethod(void *this, struct CThostFtdcRspUserAuthMethodField *pRspUserAuthMethod, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

extern void CgoOnRspGenUserCaptcha(void *this, struct CThostFtdcRspGenUserCaptchaField *pRspGenUserCaptcha, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

extern void CgoOnRspGenUserText(void *this, struct CThostFtdcRspGenUserTextField *pRspGenUserText, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

extern void CgoOnRspOrderInsert(void *this, struct CThostFtdcInputOrderField *pInputOrder, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

extern void CgoOnRspParkedOrderInsert(void *this, struct CThostFtdcParkedOrderField *pParkedOrder, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

extern void CgoOnRspParkedOrderAction(void *this, struct CThostFtdcParkedOrderActionField *pParkedOrderAction, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

extern void CgoOnRspOrderAction(void *this, struct CThostFtdcInputOrderActionField *pInputOrderAction, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

extern void CgoOnRspQryMaxOrderVolume(void *this, struct CThostFtdcQryMaxOrderVolumeField *pQryMaxOrderVolume, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

extern void CgoOnRspSettlementInfoConfirm(void *this, struct CThostFtdcSettlementInfoConfirmField *pSettlementInfoConfirm, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

extern void CgoOnRspRemoveParkedOrder(void *this, struct CThostFtdcRemoveParkedOrderField *pRemoveParkedOrder, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

extern void CgoOnRspRemoveParkedOrderAction(void *this, struct CThostFtdcRemoveParkedOrderActionField *pRemoveParkedOrderAction, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

extern void CgoOnRspExecOrderInsert(void *this, struct CThostFtdcInputExecOrderField *pInputExecOrder, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

extern void CgoOnRspExecOrderAction(void *this, struct CThostFtdcInputExecOrderActionField *pInputExecOrderAction, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

extern void CgoOnRspForQuoteInsert(void *this, struct CThostFtdcInputForQuoteField *pInputForQuote, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

extern void CgoOnRspQuoteInsert(void *this, struct CThostFtdcInputQuoteField *pInputQuote, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

extern void CgoOnRspQuoteAction(void *this, struct CThostFtdcInputQuoteActionField *pInputQuoteAction, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

extern void CgoOnRspBatchOrderAction(void *this, struct CThostFtdcInputBatchOrderActionField *pInputBatchOrderAction, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

extern void CgoOnRspOptionSelfCloseInsert(void *this, struct CThostFtdcInputOptionSelfCloseField *pInputOptionSelfClose, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

extern void CgoOnRspOptionSelfCloseAction(void *this, struct CThostFtdcInputOptionSelfCloseActionField *pInputOptionSelfCloseAction, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

extern void CgoOnRspCombActionInsert(void *this, struct CThostFtdcInputCombActionField *pInputCombAction, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

extern void CgoOnRspQryOrder(void *this, struct CThostFtdcOrderField *pOrder, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

extern void CgoOnRspQryTrade(void *this, struct CThostFtdcTradeField *pTrade, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

extern void CgoOnRspQryInvestorPosition(void *this, struct CThostFtdcInvestorPositionField *pInvestorPosition, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

extern void CgoOnRspQryTradingAccount(void *this, struct CThostFtdcTradingAccountField *pTradingAccount, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

extern void CgoOnRspQryInvestor(void *this, struct CThostFtdcInvestorField *pInvestor, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

extern void CgoOnRspQryTradingCode(void *this, struct CThostFtdcTradingCodeField *pTradingCode, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

extern void CgoOnRspQryInstrumentMarginRate(void *this, struct CThostFtdcInstrumentMarginRateField *pInstrumentMarginRate, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

extern void CgoOnRspQryInstrumentCommissionRate(void *this, struct CThostFtdcInstrumentCommissionRateField *pInstrumentCommissionRate, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

extern void CgoOnRspQryUserSession(void *this, struct CThostFtdcUserSessionField *pUserSession, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

extern void CgoOnRspQryExchange(void *this, struct CThostFtdcExchangeField *pExchange, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

extern void CgoOnRspQryProduct(void *this, struct CThostFtdcProductField *pProduct, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

extern void CgoOnRspQryInstrument(void *this, struct CThostFtdcInstrumentField *pInstrument, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

extern void CgoOnRspQryDepthMarketData(void *this, struct CThostFtdcDepthMarketDataField *pDepthMarketData, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

extern void CgoOnRspQryTraderOffer(void *this, struct CThostFtdcTraderOfferField *pTraderOffer, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

extern void CgoOnRspQrySettlementInfo(void *this, struct CThostFtdcSettlementInfoField *pSettlementInfo, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

extern void CgoOnRspQryTransferBank(void *this, struct CThostFtdcTransferBankField *pTransferBank, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

extern void CgoOnRspQryInvestorPositionDetail(void *this, struct CThostFtdcInvestorPositionDetailField *pInvestorPositionDetail, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

extern void CgoOnRspQryNotice(void *this, struct CThostFtdcNoticeField *pNotice, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

extern void CgoOnRspQrySettlementInfoConfirm(void *this, struct CThostFtdcSettlementInfoConfirmField *pSettlementInfoConfirm, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

extern void CgoOnRspQryInvestorPositionCombineDetail(void *this, struct CThostFtdcInvestorPositionCombineDetailField *pInvestorPositionCombineDetail, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

extern void CgoOnRspQryCFMMCTradingAccountKey(void *this, struct CThostFtdcCFMMCTradingAccountKeyField *pCFMMCTradingAccountKey, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

extern void CgoOnRspQryEWarrantOffset(void *this, struct CThostFtdcEWarrantOffsetField *pEWarrantOffset, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

extern void CgoOnRspQryInvestorProductGroupMargin(void *this, struct CThostFtdcInvestorProductGroupMarginField *pInvestorProductGroupMargin, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

extern void CgoOnRspQryExchangeMarginRate(void *this, struct CThostFtdcExchangeMarginRateField *pExchangeMarginRate, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

extern void CgoOnRspQryExchangeMarginRateAdjust(void *this, struct CThostFtdcExchangeMarginRateAdjustField *pExchangeMarginRateAdjust, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

extern void CgoOnRspQryExchangeRate(void *this, struct CThostFtdcExchangeRateField *pExchangeRate, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

extern void CgoOnRspQrySecAgentACIDMap(void *this, struct CThostFtdcSecAgentACIDMapField *pSecAgentACIDMap, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

extern void CgoOnRspQryProductExchRate(void *this, struct CThostFtdcProductExchRateField *pProductExchRate, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

extern void CgoOnRspQryProductGroup(void *this, struct CThostFtdcProductGroupField *pProductGroup, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

extern void CgoOnRspQryMMInstrumentCommissionRate(void *this, struct CThostFtdcMMInstrumentCommissionRateField *pMMInstrumentCommissionRate, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

extern void CgoOnRspQryMMOptionInstrCommRate(void *this, struct CThostFtdcMMOptionInstrCommRateField *pMMOptionInstrCommRate, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

extern void CgoOnRspQryInstrumentOrderCommRate(void *this, struct CThostFtdcInstrumentOrderCommRateField *pInstrumentOrderCommRate, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

extern void CgoOnRspQrySecAgentTradingAccount(void *this, struct CThostFtdcTradingAccountField *pTradingAccount, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

extern void CgoOnRspQrySecAgentCheckMode(void *this, struct CThostFtdcSecAgentCheckModeField *pSecAgentCheckMode, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

extern void CgoOnRspQrySecAgentTradeInfo(void *this, struct CThostFtdcSecAgentTradeInfoField *pSecAgentTradeInfo, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

extern void CgoOnRspQryOptionInstrTradeCost(void *this, struct CThostFtdcOptionInstrTradeCostField *pOptionInstrTradeCost, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

extern void CgoOnRspQryOptionInstrCommRate(void *this, struct CThostFtdcOptionInstrCommRateField *pOptionInstrCommRate, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

extern void CgoOnRspQryExecOrder(void *this, struct CThostFtdcExecOrderField *pExecOrder, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

extern void CgoOnRspQryForQuote(void *this, struct CThostFtdcForQuoteField *pForQuote, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

extern void CgoOnRspQryQuote(void *this, struct CThostFtdcQuoteField *pQuote, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

extern void CgoOnRspQryOptionSelfClose(void *this, struct CThostFtdcOptionSelfCloseField *pOptionSelfClose, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

extern void CgoOnRspQryInvestUnit(void *this, struct CThostFtdcInvestUnitField *pInvestUnit, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

extern void CgoOnRspQryCombInstrumentGuard(void *this, struct CThostFtdcCombInstrumentGuardField *pCombInstrumentGuard, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

extern void CgoOnRspQryCombAction(void *this, struct CThostFtdcCombActionField *pCombAction, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

extern void CgoOnRspQryTransferSerial(void *this, struct CThostFtdcTransferSerialField *pTransferSerial, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

extern void CgoOnRspQryAccountregister(void *this, struct CThostFtdcAccountregisterField *pAccountregister, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

extern void CgoOnRspError(void *this, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

extern void CgoOnRtnOrder(void *this, struct CThostFtdcOrderField *pOrder);

extern void CgoOnRtnTrade(void *this, struct CThostFtdcTradeField *pTrade);

extern void CgoOnErrRtnOrderInsert(void *this, struct CThostFtdcInputOrderField *pInputOrder, struct CThostFtdcRspInfoField *pRspInfo);

extern void CgoOnErrRtnOrderAction(void *this, struct CThostFtdcOrderActionField *pOrderAction, struct CThostFtdcRspInfoField *pRspInfo);

extern void CgoOnRtnInstrumentStatus(void *this, struct CThostFtdcInstrumentStatusField *pInstrumentStatus);

extern void CgoOnRtnBulletin(void *this, struct CThostFtdcBulletinField *pBulletin);

extern void CgoOnRtnTradingNotice(void *this, struct CThostFtdcTradingNoticeInfoField *pTradingNoticeInfo);

extern void CgoOnRtnErrorConditionalOrder(void *this, struct CThostFtdcErrorConditionalOrderField *pErrorConditionalOrder);

extern void CgoOnRtnExecOrder(void *this, struct CThostFtdcExecOrderField *pExecOrder);

extern void CgoOnErrRtnExecOrderInsert(void *this, struct CThostFtdcInputExecOrderField *pInputExecOrder, struct CThostFtdcRspInfoField *pRspInfo);

extern void CgoOnErrRtnExecOrderAction(void *this, struct CThostFtdcExecOrderActionField *pExecOrderAction, struct CThostFtdcRspInfoField *pRspInfo);

extern void CgoOnErrRtnForQuoteInsert(void *this, struct CThostFtdcInputForQuoteField *pInputForQuote, struct CThostFtdcRspInfoField *pRspInfo);

extern void CgoOnRtnQuote(void *this, struct CThostFtdcQuoteField *pQuote);

extern void CgoOnErrRtnQuoteInsert(void *this, struct CThostFtdcInputQuoteField *pInputQuote, struct CThostFtdcRspInfoField *pRspInfo);

extern void CgoOnErrRtnQuoteAction(void *this, struct CThostFtdcQuoteActionField *pQuoteAction, struct CThostFtdcRspInfoField *pRspInfo);

extern void CgoOnRtnForQuoteRsp(void *this, struct CThostFtdcForQuoteRspField *pForQuoteRsp);

extern void CgoOnRtnCFMMCTradingAccountToken(void *this, struct CThostFtdcCFMMCTradingAccountTokenField *pCFMMCTradingAccountToken);

extern void CgoOnErrRtnBatchOrderAction(void *this, struct CThostFtdcBatchOrderActionField *pBatchOrderAction, struct CThostFtdcRspInfoField *pRspInfo);

extern void CgoOnRtnOptionSelfClose(void *this, struct CThostFtdcOptionSelfCloseField *pOptionSelfClose);

extern void CgoOnErrRtnOptionSelfCloseInsert(void *this, struct CThostFtdcInputOptionSelfCloseField *pInputOptionSelfClose, struct CThostFtdcRspInfoField *pRspInfo);

extern void CgoOnErrRtnOptionSelfCloseAction(void *this, struct CThostFtdcOptionSelfCloseActionField *pOptionSelfCloseAction, struct CThostFtdcRspInfoField *pRspInfo);

extern void CgoOnRtnCombAction(void *this, struct CThostFtdcCombActionField *pCombAction);

extern void CgoOnErrRtnCombActionInsert(void *this, struct CThostFtdcInputCombActionField *pInputCombAction, struct CThostFtdcRspInfoField *pRspInfo);

extern void CgoOnRspQryContractBank(void *this, struct CThostFtdcContractBankField *pContractBank, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

extern void CgoOnRspQryParkedOrder(void *this, struct CThostFtdcParkedOrderField *pParkedOrder, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

extern void CgoOnRspQryParkedOrderAction(void *this, struct CThostFtdcParkedOrderActionField *pParkedOrderAction, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

extern void CgoOnRspQryTradingNotice(void *this, struct CThostFtdcTradingNoticeField *pTradingNotice, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

extern void CgoOnRspQryBrokerTradingParams(void *this, struct CThostFtdcBrokerTradingParamsField *pBrokerTradingParams, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

extern void CgoOnRspQryBrokerTradingAlgos(void *this, struct CThostFtdcBrokerTradingAlgosField *pBrokerTradingAlgos, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

extern void CgoOnRspQueryCFMMCTradingAccountToken(void *this, struct CThostFtdcQueryCFMMCTradingAccountTokenField *pQueryCFMMCTradingAccountToken, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

extern void CgoOnRtnFromBankToFutureByBank(void *this, struct CThostFtdcRspTransferField *pRspTransfer);

extern void CgoOnRtnFromFutureToBankByBank(void *this, struct CThostFtdcRspTransferField *pRspTransfer);

extern void CgoOnRtnRepealFromBankToFutureByBank(void *this, struct CThostFtdcRspRepealField *pRspRepeal);

extern void CgoOnRtnRepealFromFutureToBankByBank(void *this, struct CThostFtdcRspRepealField *pRspRepeal);

extern void CgoOnRtnFromBankToFutureByFuture(void *this, struct CThostFtdcRspTransferField *pRspTransfer);

extern void CgoOnRtnFromFutureToBankByFuture(void *this, struct CThostFtdcRspTransferField *pRspTransfer);

extern void CgoOnRtnRepealFromBankToFutureByFutureManual(void *this, struct CThostFtdcRspRepealField *pRspRepeal);

extern void CgoOnRtnRepealFromFutureToBankByFutureManual(void *this, struct CThostFtdcRspRepealField *pRspRepeal);

extern void CgoOnRtnQueryBankBalanceByFuture(void *this, struct CThostFtdcNotifyQueryAccountField *pNotifyQueryAccount);

extern void CgoOnErrRtnBankToFutureByFuture(void *this, struct CThostFtdcReqTransferField *pReqTransfer, struct CThostFtdcRspInfoField *pRspInfo);

extern void CgoOnErrRtnFutureToBankByFuture(void *this, struct CThostFtdcReqTransferField *pReqTransfer, struct CThostFtdcRspInfoField *pRspInfo);

extern void CgoOnErrRtnRepealBankToFutureByFutureManual(void *this, struct CThostFtdcReqRepealField *pReqRepeal, struct CThostFtdcRspInfoField *pRspInfo);

extern void CgoOnErrRtnRepealFutureToBankByFutureManual(void *this, struct CThostFtdcReqRepealField *pReqRepeal, struct CThostFtdcRspInfoField *pRspInfo);

extern void CgoOnErrRtnQueryBankBalanceByFuture(void *this, struct CThostFtdcReqQueryAccountField *pReqQueryAccount, struct CThostFtdcRspInfoField *pRspInfo);

extern void CgoOnRtnRepealFromBankToFutureByFuture(void *this, struct CThostFtdcRspRepealField *pRspRepeal);

extern void CgoOnRtnRepealFromFutureToBankByFuture(void *this, struct CThostFtdcRspRepealField *pRspRepeal);

extern void CgoOnRspFromBankToFutureByFuture(void *this, struct CThostFtdcReqTransferField *pReqTransfer, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

extern void CgoOnRspFromFutureToBankByFuture(void *this, struct CThostFtdcReqTransferField *pReqTransfer, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

extern void CgoOnRspQueryBankAccountMoneyByFuture(void *this, struct CThostFtdcReqQueryAccountField *pReqQueryAccount, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

extern void CgoOnRtnOpenAccountByBank(void *this, struct CThostFtdcOpenAccountField *pOpenAccount);

extern void CgoOnRtnCancelAccountByBank(void *this, struct CThostFtdcCancelAccountField *pCancelAccount);

extern void CgoOnRtnChangeAccountByBank(void *this, struct CThostFtdcChangeAccountField *pChangeAccount);

extern void CgoOnRspQryClassifiedInstrument(void *this, struct CThostFtdcInstrumentField *pInstrument, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

extern void CgoOnRspQryCombPromotionParam(void *this, struct CThostFtdcCombPromotionParamField *pCombPromotionParam, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

extern void CgoOnRspQryRiskSettleInvstPosition(void *this, struct CThostFtdcRiskSettleInvstPositionField *pRiskSettleInvstPosition, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

extern void CgoOnRspQryRiskSettleProductStatus(void *this, struct CThostFtdcRiskSettleProductStatusField *pRiskSettleProductStatus, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

extern void CgoOnRspQrySPBMFutureParameter(void *this, struct CThostFtdcSPBMFutureParameterField *pSPBMFutureParameter, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

extern void CgoOnRspQrySPBMOptionParameter(void *this, struct CThostFtdcSPBMOptionParameterField *pSPBMOptionParameter, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

extern void CgoOnRspQrySPBMIntraParameter(void *this, struct CThostFtdcSPBMIntraParameterField *pSPBMIntraParameter, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

extern void CgoOnRspQrySPBMInterParameter(void *this, struct CThostFtdcSPBMInterParameterField *pSPBMInterParameter, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

extern void CgoOnRspQrySPBMPortfDefinition(void *this, struct CThostFtdcSPBMPortfDefinitionField *pSPBMPortfDefinition, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

extern void CgoOnRspQrySPBMInvestorPortfDef(void *this, struct CThostFtdcSPBMInvestorPortfDefField *pSPBMInvestorPortfDef, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

extern void CgoOnRspQryInvestorPortfMarginRatio(void *this, struct CThostFtdcInvestorPortfMarginRatioField *pInvestorPortfMarginRatio, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

extern void CgoOnRspQryInvestorProdSPBMDetail(void *this, struct CThostFtdcInvestorProdSPBMDetailField *pInvestorProdSPBMDetail, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

extern void CgoOnRspQryInvestorCommoditySPMMMargin(void *this, struct CThostFtdcInvestorCommoditySPMMMarginField *pInvestorCommoditySPMMMargin, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

extern void CgoOnRspQryInvestorCommodityGroupSPMMMargin(void *this, struct CThostFtdcInvestorCommodityGroupSPMMMarginField *pInvestorCommodityGroupSPMMMargin, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

extern void CgoOnRspQrySPMMInstParam(void *this, struct CThostFtdcSPMMInstParamField *pSPMMInstParam, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

extern void CgoOnRspQrySPMMProductParam(void *this, struct CThostFtdcSPMMProductParamField *pSPMMProductParam, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

extern void CgoOnRspQrySPBMAddOnInterParameter(void *this, struct CThostFtdcSPBMAddOnInterParameterField *pSPBMAddOnInterParameter, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

extern void CgoOnRspQryRCAMSCombProductInfo(void *this, struct CThostFtdcRCAMSCombProductInfoField *pRCAMSCombProductInfo, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

extern void CgoOnRspQryRCAMSInstrParameter(void *this, struct CThostFtdcRCAMSInstrParameterField *pRCAMSInstrParameter, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

extern void CgoOnRspQryRCAMSIntraParameter(void *this, struct CThostFtdcRCAMSIntraParameterField *pRCAMSIntraParameter, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

extern void CgoOnRspQryRCAMSInterParameter(void *this, struct CThostFtdcRCAMSInterParameterField *pRCAMSInterParameter, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

extern void CgoOnRspQryRCAMSShortOptAdjustParam(void *this, struct CThostFtdcRCAMSShortOptAdjustParamField *pRCAMSShortOptAdjustParam, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

extern void CgoOnRspQryRCAMSInvestorCombPosition(void *this, struct CThostFtdcRCAMSInvestorCombPositionField *pRCAMSInvestorCombPosition, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

extern void CgoOnRspQryInvestorProdRCAMSMargin(void *this, struct CThostFtdcInvestorProdRCAMSMarginField *pInvestorProdRCAMSMargin, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

extern void CgoOnRspQryRULEInstrParameter(void *this, struct CThostFtdcRULEInstrParameterField *pRULEInstrParameter, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

extern void CgoOnRspQryRULEIntraParameter(void *this, struct CThostFtdcRULEIntraParameterField *pRULEIntraParameter, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

extern void CgoOnRspQryRULEInterParameter(void *this, struct CThostFtdcRULEInterParameterField *pRULEInterParameter, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

extern void CgoOnRspQryInvestorProdRULEMargin(void *this, struct CThostFtdcInvestorProdRULEMarginField *pInvestorProdRULEMargin, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

extern void CgoOnRspQryInvestorPortfSetting(void *this, struct CThostFtdcInvestorPortfSettingField *pInvestorPortfSetting, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

extern void CgoOnRspQryInvestorInfoCommRec(void *this, struct CThostFtdcInvestorInfoCommRecField *pInvestorInfoCommRec, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

extern void CgoOnRspQryCombLeg(void *this, struct CThostFtdcCombLegField *pCombLeg, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

extern void CgoOnRspOffsetSetting(void *this, struct CThostFtdcInputOffsetSettingField *pInputOffsetSetting, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

extern void CgoOnRspCancelOffsetSetting(void *this, struct CThostFtdcInputOffsetSettingField *pInputOffsetSetting, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

extern void CgoOnRtnOffsetSetting(void *this, struct CThostFtdcOffsetSettingField *pOffsetSetting);

extern void CgoOnErrRtnOffsetSetting(void *this, struct CThostFtdcInputOffsetSettingField *pInputOffsetSetting, struct CThostFtdcRspInfoField *pRspInfo);

extern void CgoOnErrRtnCancelOffsetSetting(void *this, struct CThostFtdcCancelOffsetSettingField *pCancelOffsetSetting, struct CThostFtdcRspInfoField *pRspInfo);

extern void CgoOnRspQryOffsetSetting(void *this, struct CThostFtdcOffsetSettingField *pOffsetSetting, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

extern void CgoOnRspGenSMSCode(void *this, struct CThostFtdcRspGenSMSCodeField *pRspGenSMSCode, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

extern void CgoOnRspSpdApply(void *this, struct CThostFtdcInputSpdApplyField *pInputSpdApply, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

extern void CgoOnRspSpdApplyAction(void *this, struct CThostFtdcInputSpdApplyActionField *pInputSpdApplyAction, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

extern void CgoOnRspQrySpdApply(void *this, struct CThostFtdcSpdApplyField *pSpdApply, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

extern void CgoOnRtnSpdApply(void *this, struct CThostFtdcSpdApplyField *pSpdApply);

extern void CgoOnErrRtnSpdApply(void *this, struct CThostFtdcInputSpdApplyField *pInputSpdApply, struct CThostFtdcRspInfoField *pRspInfo);

extern void CgoOnErrRtnSpdApplyAction(void *this, struct CThostFtdcSpdApplyActionField *pSpdApplyAction, struct CThostFtdcRspInfoField *pRspInfo);

extern void CgoOnRspHedgeCfm(void *this, struct CThostFtdcInputHedgeCfmField *pInputHedgeCfm, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

extern void CgoOnRspHedgeCfmAction(void *this, struct CThostFtdcInputHedgeCfmActionField *pInputHedgeCfmAction, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

extern void CgoOnRspQryHedgeCfm(void *this, struct CThostFtdcHedgeCfmField *pHedgeCfm, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

extern void CgoOnRtnHedgeCfm(void *this, struct CThostFtdcHedgeCfmField *pHedgeCfm);

extern void CgoOnErrRtnHedgeCfm(void *this, struct CThostFtdcInputHedgeCfmField *pInputHedgeCfm, struct CThostFtdcRspInfoField *pRspInfo);

extern void CgoOnErrRtnHedgeCfmAction(void *this, struct CThostFtdcHedgeCfmActionField *pHedgeCfmAction, struct CThostFtdcRspInfoField *pRspInfo);

void COnFrontConnected(void *this)
{
    CgoOnFrontConnected(this);
}

void COnFrontDisconnected(void *this, int nReason)
{
    CgoOnFrontDisconnected(this, nReason);
}

void COnHeartBeatWarning(void *this, int nTimeLapse)
{
    CgoOnHeartBeatWarning(this, nTimeLapse);
}

void COnRspAuthenticate(void *this, struct CThostFtdcRspAuthenticateField *pRspAuthenticateField, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast)
{
    CgoOnRspAuthenticate(this, pRspAuthenticateField, pRspInfo, nRequestID, bIsLast);
}

void COnRtnPrivateSeqNo(void *this, int nSeqNo)
{
    CgoOnRtnPrivateSeqNo(this, nSeqNo);
}

void COnRspUserLogin(void *this, struct CThostFtdcRspUserLoginField *pRspUserLogin, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast)
{
    CgoOnRspUserLogin(this, pRspUserLogin, pRspInfo, nRequestID, bIsLast);
}

void COnRspUserLogout(void *this, struct CThostFtdcUserLogoutField *pUserLogout, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast)
{
    CgoOnRspUserLogout(this, pUserLogout, pRspInfo, nRequestID, bIsLast);
}

void COnRspUserPasswordUpdate(void *this, struct CThostFtdcUserPasswordUpdateField *pUserPasswordUpdate, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast)
{
    CgoOnRspUserPasswordUpdate(this, pUserPasswordUpdate, pRspInfo, nRequestID, bIsLast);
}

void COnRspTradingAccountPasswordUpdate(void *this, struct CThostFtdcTradingAccountPasswordUpdateField *pTradingAccountPasswordUpdate, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast)
{
    CgoOnRspTradingAccountPasswordUpdate(this, pTradingAccountPasswordUpdate, pRspInfo, nRequestID, bIsLast);
}

void COnRspUserAuthMethod(void *this, struct CThostFtdcRspUserAuthMethodField *pRspUserAuthMethod, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast)
{
    CgoOnRspUserAuthMethod(this, pRspUserAuthMethod, pRspInfo, nRequestID, bIsLast);
}

void COnRspGenUserCaptcha(void *this, struct CThostFtdcRspGenUserCaptchaField *pRspGenUserCaptcha, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast)
{
    CgoOnRspGenUserCaptcha(this, pRspGenUserCaptcha, pRspInfo, nRequestID, bIsLast);
}

void COnRspGenUserText(void *this, struct CThostFtdcRspGenUserTextField *pRspGenUserText, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast)
{
    CgoOnRspGenUserText(this, pRspGenUserText, pRspInfo, nRequestID, bIsLast);
}

void COnRspOrderInsert(void *this, struct CThostFtdcInputOrderField *pInputOrder, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast)
{
    CgoOnRspOrderInsert(this, pInputOrder, pRspInfo, nRequestID, bIsLast);
}

void COnRspParkedOrderInsert(void *this, struct CThostFtdcParkedOrderField *pParkedOrder, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast)
{
    CgoOnRspParkedOrderInsert(this, pParkedOrder, pRspInfo, nRequestID, bIsLast);
}

void COnRspParkedOrderAction(void *this, struct CThostFtdcParkedOrderActionField *pParkedOrderAction, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast)
{
    CgoOnRspParkedOrderAction(this, pParkedOrderAction, pRspInfo, nRequestID, bIsLast);
}

void COnRspOrderAction(void *this, struct CThostFtdcInputOrderActionField *pInputOrderAction, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast)
{
    CgoOnRspOrderAction(this, pInputOrderAction, pRspInfo, nRequestID, bIsLast);
}

void COnRspQryMaxOrderVolume(void *this, struct CThostFtdcQryMaxOrderVolumeField *pQryMaxOrderVolume, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast)
{
    CgoOnRspQryMaxOrderVolume(this, pQryMaxOrderVolume, pRspInfo, nRequestID, bIsLast);
}

void COnRspSettlementInfoConfirm(void *this, struct CThostFtdcSettlementInfoConfirmField *pSettlementInfoConfirm, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast)
{
    CgoOnRspSettlementInfoConfirm(this, pSettlementInfoConfirm, pRspInfo, nRequestID, bIsLast);
}

void COnRspRemoveParkedOrder(void *this, struct CThostFtdcRemoveParkedOrderField *pRemoveParkedOrder, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast)
{
    CgoOnRspRemoveParkedOrder(this, pRemoveParkedOrder, pRspInfo, nRequestID, bIsLast);
}

void COnRspRemoveParkedOrderAction(void *this, struct CThostFtdcRemoveParkedOrderActionField *pRemoveParkedOrderAction, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast)
{
    CgoOnRspRemoveParkedOrderAction(this, pRemoveParkedOrderAction, pRspInfo, nRequestID, bIsLast);
}

void COnRspExecOrderInsert(void *this, struct CThostFtdcInputExecOrderField *pInputExecOrder, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast)
{
    CgoOnRspExecOrderInsert(this, pInputExecOrder, pRspInfo, nRequestID, bIsLast);
}

void COnRspExecOrderAction(void *this, struct CThostFtdcInputExecOrderActionField *pInputExecOrderAction, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast)
{
    CgoOnRspExecOrderAction(this, pInputExecOrderAction, pRspInfo, nRequestID, bIsLast);
}

void COnRspForQuoteInsert(void *this, struct CThostFtdcInputForQuoteField *pInputForQuote, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast)
{
    CgoOnRspForQuoteInsert(this, pInputForQuote, pRspInfo, nRequestID, bIsLast);
}

void COnRspQuoteInsert(void *this, struct CThostFtdcInputQuoteField *pInputQuote, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast)
{
    CgoOnRspQuoteInsert(this, pInputQuote, pRspInfo, nRequestID, bIsLast);
}

void COnRspQuoteAction(void *this, struct CThostFtdcInputQuoteActionField *pInputQuoteAction, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast)
{
    CgoOnRspQuoteAction(this, pInputQuoteAction, pRspInfo, nRequestID, bIsLast);
}

void COnRspBatchOrderAction(void *this, struct CThostFtdcInputBatchOrderActionField *pInputBatchOrderAction, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast)
{
    CgoOnRspBatchOrderAction(this, pInputBatchOrderAction, pRspInfo, nRequestID, bIsLast);
}

void COnRspOptionSelfCloseInsert(void *this, struct CThostFtdcInputOptionSelfCloseField *pInputOptionSelfClose, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast)
{
    CgoOnRspOptionSelfCloseInsert(this, pInputOptionSelfClose, pRspInfo, nRequestID, bIsLast);
}

void COnRspOptionSelfCloseAction(void *this, struct CThostFtdcInputOptionSelfCloseActionField *pInputOptionSelfCloseAction, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast)
{
    CgoOnRspOptionSelfCloseAction(this, pInputOptionSelfCloseAction, pRspInfo, nRequestID, bIsLast);
}

void COnRspCombActionInsert(void *this, struct CThostFtdcInputCombActionField *pInputCombAction, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast)
{
    CgoOnRspCombActionInsert(this, pInputCombAction, pRspInfo, nRequestID, bIsLast);
}

void COnRspQryOrder(void *this, struct CThostFtdcOrderField *pOrder, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast)
{
    CgoOnRspQryOrder(this, pOrder, pRspInfo, nRequestID, bIsLast);
}

void COnRspQryTrade(void *this, struct CThostFtdcTradeField *pTrade, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast)
{
    CgoOnRspQryTrade(this, pTrade, pRspInfo, nRequestID, bIsLast);
}

void COnRspQryInvestorPosition(void *this, struct CThostFtdcInvestorPositionField *pInvestorPosition, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast)
{
    CgoOnRspQryInvestorPosition(this, pInvestorPosition, pRspInfo, nRequestID, bIsLast);
}

void COnRspQryTradingAccount(void *this, struct CThostFtdcTradingAccountField *pTradingAccount, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast)
{
    CgoOnRspQryTradingAccount(this, pTradingAccount, pRspInfo, nRequestID, bIsLast);
}

void COnRspQryInvestor(void *this, struct CThostFtdcInvestorField *pInvestor, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast)
{
    CgoOnRspQryInvestor(this, pInvestor, pRspInfo, nRequestID, bIsLast);
}

void COnRspQryTradingCode(void *this, struct CThostFtdcTradingCodeField *pTradingCode, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast)
{
    CgoOnRspQryTradingCode(this, pTradingCode, pRspInfo, nRequestID, bIsLast);
}

void COnRspQryInstrumentMarginRate(void *this, struct CThostFtdcInstrumentMarginRateField *pInstrumentMarginRate, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast)
{
    CgoOnRspQryInstrumentMarginRate(this, pInstrumentMarginRate, pRspInfo, nRequestID, bIsLast);
}

void COnRspQryInstrumentCommissionRate(void *this, struct CThostFtdcInstrumentCommissionRateField *pInstrumentCommissionRate, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast)
{
    CgoOnRspQryInstrumentCommissionRate(this, pInstrumentCommissionRate, pRspInfo, nRequestID, bIsLast);
}

void COnRspQryUserSession(void *this, struct CThostFtdcUserSessionField *pUserSession, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast)
{
    CgoOnRspQryUserSession(this, pUserSession, pRspInfo, nRequestID, bIsLast);
}

void COnRspQryExchange(void *this, struct CThostFtdcExchangeField *pExchange, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast)
{
    CgoOnRspQryExchange(this, pExchange, pRspInfo, nRequestID, bIsLast);
}

void COnRspQryProduct(void *this, struct CThostFtdcProductField *pProduct, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast)
{
    CgoOnRspQryProduct(this, pProduct, pRspInfo, nRequestID, bIsLast);
}

void COnRspQryInstrument(void *this, struct CThostFtdcInstrumentField *pInstrument, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast)
{
    CgoOnRspQryInstrument(this, pInstrument, pRspInfo, nRequestID, bIsLast);
}

void COnRspQryDepthMarketData(void *this, struct CThostFtdcDepthMarketDataField *pDepthMarketData, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast)
{
    CgoOnRspQryDepthMarketData(this, pDepthMarketData, pRspInfo, nRequestID, bIsLast);
}

void COnRspQryTraderOffer(void *this, struct CThostFtdcTraderOfferField *pTraderOffer, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast)
{
    CgoOnRspQryTraderOffer(this, pTraderOffer, pRspInfo, nRequestID, bIsLast);
}

void COnRspQrySettlementInfo(void *this, struct CThostFtdcSettlementInfoField *pSettlementInfo, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast)
{
    CgoOnRspQrySettlementInfo(this, pSettlementInfo, pRspInfo, nRequestID, bIsLast);
}

void COnRspQryTransferBank(void *this, struct CThostFtdcTransferBankField *pTransferBank, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast)
{
    CgoOnRspQryTransferBank(this, pTransferBank, pRspInfo, nRequestID, bIsLast);
}

void COnRspQryInvestorPositionDetail(void *this, struct CThostFtdcInvestorPositionDetailField *pInvestorPositionDetail, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast)
{
    CgoOnRspQryInvestorPositionDetail(this, pInvestorPositionDetail, pRspInfo, nRequestID, bIsLast);
}

void COnRspQryNotice(void *this, struct CThostFtdcNoticeField *pNotice, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast)
{
    CgoOnRspQryNotice(this, pNotice, pRspInfo, nRequestID, bIsLast);
}

void COnRspQrySettlementInfoConfirm(void *this, struct CThostFtdcSettlementInfoConfirmField *pSettlementInfoConfirm, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast)
{
    CgoOnRspQrySettlementInfoConfirm(this, pSettlementInfoConfirm, pRspInfo, nRequestID, bIsLast);
}

void COnRspQryInvestorPositionCombineDetail(void *this, struct CThostFtdcInvestorPositionCombineDetailField *pInvestorPositionCombineDetail, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast)
{
    CgoOnRspQryInvestorPositionCombineDetail(this, pInvestorPositionCombineDetail, pRspInfo, nRequestID, bIsLast);
}

void COnRspQryCFMMCTradingAccountKey(void *this, struct CThostFtdcCFMMCTradingAccountKeyField *pCFMMCTradingAccountKey, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast)
{
    CgoOnRspQryCFMMCTradingAccountKey(this, pCFMMCTradingAccountKey, pRspInfo, nRequestID, bIsLast);
}

void COnRspQryEWarrantOffset(void *this, struct CThostFtdcEWarrantOffsetField *pEWarrantOffset, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast)
{
    CgoOnRspQryEWarrantOffset(this, pEWarrantOffset, pRspInfo, nRequestID, bIsLast);
}

void COnRspQryInvestorProductGroupMargin(void *this, struct CThostFtdcInvestorProductGroupMarginField *pInvestorProductGroupMargin, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast)
{
    CgoOnRspQryInvestorProductGroupMargin(this, pInvestorProductGroupMargin, pRspInfo, nRequestID, bIsLast);
}

void COnRspQryExchangeMarginRate(void *this, struct CThostFtdcExchangeMarginRateField *pExchangeMarginRate, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast)
{
    CgoOnRspQryExchangeMarginRate(this, pExchangeMarginRate, pRspInfo, nRequestID, bIsLast);
}

void COnRspQryExchangeMarginRateAdjust(void *this, struct CThostFtdcExchangeMarginRateAdjustField *pExchangeMarginRateAdjust, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast)
{
    CgoOnRspQryExchangeMarginRateAdjust(this, pExchangeMarginRateAdjust, pRspInfo, nRequestID, bIsLast);
}

void COnRspQryExchangeRate(void *this, struct CThostFtdcExchangeRateField *pExchangeRate, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast)
{
    CgoOnRspQryExchangeRate(this, pExchangeRate, pRspInfo, nRequestID, bIsLast);
}

void COnRspQrySecAgentACIDMap(void *this, struct CThostFtdcSecAgentACIDMapField *pSecAgentACIDMap, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast)
{
    CgoOnRspQrySecAgentACIDMap(this, pSecAgentACIDMap, pRspInfo, nRequestID, bIsLast);
}

void COnRspQryProductExchRate(void *this, struct CThostFtdcProductExchRateField *pProductExchRate, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast)
{
    CgoOnRspQryProductExchRate(this, pProductExchRate, pRspInfo, nRequestID, bIsLast);
}

void COnRspQryProductGroup(void *this, struct CThostFtdcProductGroupField *pProductGroup, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast)
{
    CgoOnRspQryProductGroup(this, pProductGroup, pRspInfo, nRequestID, bIsLast);
}

void COnRspQryMMInstrumentCommissionRate(void *this, struct CThostFtdcMMInstrumentCommissionRateField *pMMInstrumentCommissionRate, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast)
{
    CgoOnRspQryMMInstrumentCommissionRate(this, pMMInstrumentCommissionRate, pRspInfo, nRequestID, bIsLast);
}

void COnRspQryMMOptionInstrCommRate(void *this, struct CThostFtdcMMOptionInstrCommRateField *pMMOptionInstrCommRate, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast)
{
    CgoOnRspQryMMOptionInstrCommRate(this, pMMOptionInstrCommRate, pRspInfo, nRequestID, bIsLast);
}

void COnRspQryInstrumentOrderCommRate(void *this, struct CThostFtdcInstrumentOrderCommRateField *pInstrumentOrderCommRate, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast)
{
    CgoOnRspQryInstrumentOrderCommRate(this, pInstrumentOrderCommRate, pRspInfo, nRequestID, bIsLast);
}

void COnRspQrySecAgentTradingAccount(void *this, struct CThostFtdcTradingAccountField *pTradingAccount, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast)
{
    CgoOnRspQrySecAgentTradingAccount(this, pTradingAccount, pRspInfo, nRequestID, bIsLast);
}

void COnRspQrySecAgentCheckMode(void *this, struct CThostFtdcSecAgentCheckModeField *pSecAgentCheckMode, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast)
{
    CgoOnRspQrySecAgentCheckMode(this, pSecAgentCheckMode, pRspInfo, nRequestID, bIsLast);
}

void COnRspQrySecAgentTradeInfo(void *this, struct CThostFtdcSecAgentTradeInfoField *pSecAgentTradeInfo, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast)
{
    CgoOnRspQrySecAgentTradeInfo(this, pSecAgentTradeInfo, pRspInfo, nRequestID, bIsLast);
}

void COnRspQryOptionInstrTradeCost(void *this, struct CThostFtdcOptionInstrTradeCostField *pOptionInstrTradeCost, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast)
{
    CgoOnRspQryOptionInstrTradeCost(this, pOptionInstrTradeCost, pRspInfo, nRequestID, bIsLast);
}

void COnRspQryOptionInstrCommRate(void *this, struct CThostFtdcOptionInstrCommRateField *pOptionInstrCommRate, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast)
{
    CgoOnRspQryOptionInstrCommRate(this, pOptionInstrCommRate, pRspInfo, nRequestID, bIsLast);
}

void COnRspQryExecOrder(void *this, struct CThostFtdcExecOrderField *pExecOrder, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast)
{
    CgoOnRspQryExecOrder(this, pExecOrder, pRspInfo, nRequestID, bIsLast);
}

void COnRspQryForQuote(void *this, struct CThostFtdcForQuoteField *pForQuote, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast)
{
    CgoOnRspQryForQuote(this, pForQuote, pRspInfo, nRequestID, bIsLast);
}

void COnRspQryQuote(void *this, struct CThostFtdcQuoteField *pQuote, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast)
{
    CgoOnRspQryQuote(this, pQuote, pRspInfo, nRequestID, bIsLast);
}

void COnRspQryOptionSelfClose(void *this, struct CThostFtdcOptionSelfCloseField *pOptionSelfClose, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast)
{
    CgoOnRspQryOptionSelfClose(this, pOptionSelfClose, pRspInfo, nRequestID, bIsLast);
}

void COnRspQryInvestUnit(void *this, struct CThostFtdcInvestUnitField *pInvestUnit, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast)
{
    CgoOnRspQryInvestUnit(this, pInvestUnit, pRspInfo, nRequestID, bIsLast);
}

void COnRspQryCombInstrumentGuard(void *this, struct CThostFtdcCombInstrumentGuardField *pCombInstrumentGuard, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast)
{
    CgoOnRspQryCombInstrumentGuard(this, pCombInstrumentGuard, pRspInfo, nRequestID, bIsLast);
}

void COnRspQryCombAction(void *this, struct CThostFtdcCombActionField *pCombAction, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast)
{
    CgoOnRspQryCombAction(this, pCombAction, pRspInfo, nRequestID, bIsLast);
}

void COnRspQryTransferSerial(void *this, struct CThostFtdcTransferSerialField *pTransferSerial, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast)
{
    CgoOnRspQryTransferSerial(this, pTransferSerial, pRspInfo, nRequestID, bIsLast);
}

void COnRspQryAccountregister(void *this, struct CThostFtdcAccountregisterField *pAccountregister, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast)
{
    CgoOnRspQryAccountregister(this, pAccountregister, pRspInfo, nRequestID, bIsLast);
}

void COnRspError(void *this, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast)
{
    CgoOnRspError(this, pRspInfo, nRequestID, bIsLast);
}

void COnRtnOrder(void *this, struct CThostFtdcOrderField *pOrder)
{
    CgoOnRtnOrder(this, pOrder);
}

void COnRtnTrade(void *this, struct CThostFtdcTradeField *pTrade)
{
    CgoOnRtnTrade(this, pTrade);
}

void COnErrRtnOrderInsert(void *this, struct CThostFtdcInputOrderField *pInputOrder, struct CThostFtdcRspInfoField *pRspInfo)
{
    CgoOnErrRtnOrderInsert(this, pInputOrder, pRspInfo);
}

void COnErrRtnOrderAction(void *this, struct CThostFtdcOrderActionField *pOrderAction, struct CThostFtdcRspInfoField *pRspInfo)
{
    CgoOnErrRtnOrderAction(this, pOrderAction, pRspInfo);
}

void COnRtnInstrumentStatus(void *this, struct CThostFtdcInstrumentStatusField *pInstrumentStatus)
{
    CgoOnRtnInstrumentStatus(this, pInstrumentStatus);
}

void COnRtnBulletin(void *this, struct CThostFtdcBulletinField *pBulletin)
{
    CgoOnRtnBulletin(this, pBulletin);
}

void COnRtnTradingNotice(void *this, struct CThostFtdcTradingNoticeInfoField *pTradingNoticeInfo)
{
    CgoOnRtnTradingNotice(this, pTradingNoticeInfo);
}

void COnRtnErrorConditionalOrder(void *this, struct CThostFtdcErrorConditionalOrderField *pErrorConditionalOrder)
{
    CgoOnRtnErrorConditionalOrder(this, pErrorConditionalOrder);
}

void COnRtnExecOrder(void *this, struct CThostFtdcExecOrderField *pExecOrder)
{
    CgoOnRtnExecOrder(this, pExecOrder);
}

void COnErrRtnExecOrderInsert(void *this, struct CThostFtdcInputExecOrderField *pInputExecOrder, struct CThostFtdcRspInfoField *pRspInfo)
{
    CgoOnErrRtnExecOrderInsert(this, pInputExecOrder, pRspInfo);
}

void COnErrRtnExecOrderAction(void *this, struct CThostFtdcExecOrderActionField *pExecOrderAction, struct CThostFtdcRspInfoField *pRspInfo)
{
    CgoOnErrRtnExecOrderAction(this, pExecOrderAction, pRspInfo);
}

void COnErrRtnForQuoteInsert(void *this, struct CThostFtdcInputForQuoteField *pInputForQuote, struct CThostFtdcRspInfoField *pRspInfo)
{
    CgoOnErrRtnForQuoteInsert(this, pInputForQuote, pRspInfo);
}

void COnRtnQuote(void *this, struct CThostFtdcQuoteField *pQuote)
{
    CgoOnRtnQuote(this, pQuote);
}

void COnErrRtnQuoteInsert(void *this, struct CThostFtdcInputQuoteField *pInputQuote, struct CThostFtdcRspInfoField *pRspInfo)
{
    CgoOnErrRtnQuoteInsert(this, pInputQuote, pRspInfo);
}

void COnErrRtnQuoteAction(void *this, struct CThostFtdcQuoteActionField *pQuoteAction, struct CThostFtdcRspInfoField *pRspInfo)
{
    CgoOnErrRtnQuoteAction(this, pQuoteAction, pRspInfo);
}

void COnRtnCFMMCTradingAccountToken(void *this, struct CThostFtdcCFMMCTradingAccountTokenField *pCFMMCTradingAccountToken)
{
    CgoOnRtnCFMMCTradingAccountToken(this, pCFMMCTradingAccountToken);
}

void COnErrRtnBatchOrderAction(void *this, struct CThostFtdcBatchOrderActionField *pBatchOrderAction, struct CThostFtdcRspInfoField *pRspInfo)
{
    CgoOnErrRtnBatchOrderAction(this, pBatchOrderAction, pRspInfo);
}

void COnRtnOptionSelfClose(void *this, struct CThostFtdcOptionSelfCloseField *pOptionSelfClose)
{
    CgoOnRtnOptionSelfClose(this, pOptionSelfClose);
}

void COnErrRtnOptionSelfCloseInsert(void *this, struct CThostFtdcInputOptionSelfCloseField *pInputOptionSelfClose, struct CThostFtdcRspInfoField *pRspInfo)
{
    CgoOnErrRtnOptionSelfCloseInsert(this, pInputOptionSelfClose, pRspInfo);
}

void COnErrRtnOptionSelfCloseAction(void *this, struct CThostFtdcOptionSelfCloseActionField *pOptionSelfCloseAction, struct CThostFtdcRspInfoField *pRspInfo)
{
    CgoOnErrRtnOptionSelfCloseAction(this, pOptionSelfCloseAction, pRspInfo);
}

void COnRtnCombAction(void *this, struct CThostFtdcCombActionField *pCombAction)
{
    CgoOnRtnCombAction(this, pCombAction);
}

void COnErrRtnCombActionInsert(void *this, struct CThostFtdcInputCombActionField *pInputCombAction, struct CThostFtdcRspInfoField *pRspInfo)
{
    CgoOnErrRtnCombActionInsert(this, pInputCombAction, pRspInfo);
}

void COnRspQryContractBank(void *this, struct CThostFtdcContractBankField *pContractBank, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast)
{
    CgoOnRspQryContractBank(this, pContractBank, pRspInfo, nRequestID, bIsLast);
}

void COnRspQryParkedOrder(void *this, struct CThostFtdcParkedOrderField *pParkedOrder, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast)
{
    CgoOnRspQryParkedOrder(this, pParkedOrder, pRspInfo, nRequestID, bIsLast);
}

void COnRspQryParkedOrderAction(void *this, struct CThostFtdcParkedOrderActionField *pParkedOrderAction, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast)
{
    CgoOnRspQryParkedOrderAction(this, pParkedOrderAction, pRspInfo, nRequestID, bIsLast);
}

void COnRspQryTradingNotice(void *this, struct CThostFtdcTradingNoticeField *pTradingNotice, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast)
{
    CgoOnRspQryTradingNotice(this, pTradingNotice, pRspInfo, nRequestID, bIsLast);
}

void COnRspQryBrokerTradingParams(void *this, struct CThostFtdcBrokerTradingParamsField *pBrokerTradingParams, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast)
{
    CgoOnRspQryBrokerTradingParams(this, pBrokerTradingParams, pRspInfo, nRequestID, bIsLast);
}

void COnRspQryBrokerTradingAlgos(void *this, struct CThostFtdcBrokerTradingAlgosField *pBrokerTradingAlgos, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast)
{
    CgoOnRspQryBrokerTradingAlgos(this, pBrokerTradingAlgos, pRspInfo, nRequestID, bIsLast);
}

void COnRspQueryCFMMCTradingAccountToken(void *this, struct CThostFtdcQueryCFMMCTradingAccountTokenField *pQueryCFMMCTradingAccountToken, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast)
{
    CgoOnRspQueryCFMMCTradingAccountToken(this, pQueryCFMMCTradingAccountToken, pRspInfo, nRequestID, bIsLast);
}

void COnRtnFromBankToFutureByBank(void *this, struct CThostFtdcRspTransferField *pRspTransfer)
{
    CgoOnRtnFromBankToFutureByBank(this, pRspTransfer);
}

void COnRtnFromFutureToBankByBank(void *this, struct CThostFtdcRspTransferField *pRspTransfer)
{
    CgoOnRtnFromFutureToBankByBank(this, pRspTransfer);
}

void COnRtnRepealFromBankToFutureByBank(void *this, struct CThostFtdcRspRepealField *pRspRepeal)
{
    CgoOnRtnRepealFromBankToFutureByBank(this, pRspRepeal);
}

void COnRtnRepealFromFutureToBankByBank(void *this, struct CThostFtdcRspRepealField *pRspRepeal)
{
    CgoOnRtnRepealFromFutureToBankByBank(this, pRspRepeal);
}

void COnRtnFromBankToFutureByFuture(void *this, struct CThostFtdcRspTransferField *pRspTransfer)
{
    CgoOnRtnFromBankToFutureByFuture(this, pRspTransfer);
}

void COnRtnFromFutureToBankByFuture(void *this, struct CThostFtdcRspTransferField *pRspTransfer)
{
    CgoOnRtnFromFutureToBankByFuture(this, pRspTransfer);
}

void COnRtnRepealFromBankToFutureByFutureManual(void *this, struct CThostFtdcRspRepealField *pRspRepeal)
{
    CgoOnRtnRepealFromBankToFutureByFutureManual(this, pRspRepeal);
}

void COnRtnRepealFromFutureToBankByFutureManual(void *this, struct CThostFtdcRspRepealField *pRspRepeal)
{
    CgoOnRtnRepealFromFutureToBankByFutureManual(this, pRspRepeal);
}

void COnRtnQueryBankBalanceByFuture(void *this, struct CThostFtdcNotifyQueryAccountField *pNotifyQueryAccount)
{
    CgoOnRtnQueryBankBalanceByFuture(this, pNotifyQueryAccount);
}

void COnErrRtnBankToFutureByFuture(void *this, struct CThostFtdcReqTransferField *pReqTransfer, struct CThostFtdcRspInfoField *pRspInfo)
{
    CgoOnErrRtnBankToFutureByFuture(this, pReqTransfer, pRspInfo);
}

void COnErrRtnFutureToBankByFuture(void *this, struct CThostFtdcReqTransferField *pReqTransfer, struct CThostFtdcRspInfoField *pRspInfo)
{
    CgoOnErrRtnFutureToBankByFuture(this, pReqTransfer, pRspInfo);
}

void COnErrRtnRepealBankToFutureByFutureManual(void *this, struct CThostFtdcReqRepealField *pReqRepeal, struct CThostFtdcRspInfoField *pRspInfo)
{
    CgoOnErrRtnRepealBankToFutureByFutureManual(this, pReqRepeal, pRspInfo);
}

void COnErrRtnRepealFutureToBankByFutureManual(void *this, struct CThostFtdcReqRepealField *pReqRepeal, struct CThostFtdcRspInfoField *pRspInfo)
{
    CgoOnErrRtnRepealFutureToBankByFutureManual(this, pReqRepeal, pRspInfo);
}

void COnErrRtnQueryBankBalanceByFuture(void *this, struct CThostFtdcReqQueryAccountField *pReqQueryAccount, struct CThostFtdcRspInfoField *pRspInfo)
{
    CgoOnErrRtnQueryBankBalanceByFuture(this, pReqQueryAccount, pRspInfo);
}

void COnRtnRepealFromBankToFutureByFuture(void *this, struct CThostFtdcRspRepealField *pRspRepeal)
{
    CgoOnRtnRepealFromBankToFutureByFuture(this, pRspRepeal);
}

void COnRtnRepealFromFutureToBankByFuture(void *this, struct CThostFtdcRspRepealField *pRspRepeal)
{
    CgoOnRtnRepealFromFutureToBankByFuture(this, pRspRepeal);
}

void COnRspFromBankToFutureByFuture(void *this, struct CThostFtdcReqTransferField *pReqTransfer, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast)
{
    CgoOnRspFromBankToFutureByFuture(this, pReqTransfer, pRspInfo, nRequestID, bIsLast);
}

void COnRspFromFutureToBankByFuture(void *this, struct CThostFtdcReqTransferField *pReqTransfer, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast)
{
    CgoOnRspFromFutureToBankByFuture(this, pReqTransfer, pRspInfo, nRequestID, bIsLast);
}

void COnRspQueryBankAccountMoneyByFuture(void *this, struct CThostFtdcReqQueryAccountField *pReqQueryAccount, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast)
{
    CgoOnRspQueryBankAccountMoneyByFuture(this, pReqQueryAccount, pRspInfo, nRequestID, bIsLast);
}

void COnRtnOpenAccountByBank(void *this, struct CThostFtdcOpenAccountField *pOpenAccount)
{
    CgoOnRtnOpenAccountByBank(this, pOpenAccount);
}

void COnRtnCancelAccountByBank(void *this, struct CThostFtdcCancelAccountField *pCancelAccount)
{
    CgoOnRtnCancelAccountByBank(this, pCancelAccount);
}

void COnRtnChangeAccountByBank(void *this, struct CThostFtdcChangeAccountField *pChangeAccount)
{
    CgoOnRtnChangeAccountByBank(this, pChangeAccount);
}

void COnRspQryClassifiedInstrument(void *this, struct CThostFtdcInstrumentField *pInstrument, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast)
{
    CgoOnRspQryClassifiedInstrument(this, pInstrument, pRspInfo, nRequestID, bIsLast);
}

void COnRspQryCombPromotionParam(void *this, struct CThostFtdcCombPromotionParamField *pCombPromotionParam, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast)
{
    CgoOnRspQryCombPromotionParam(this, pCombPromotionParam, pRspInfo, nRequestID, bIsLast);
}

void COnRspQryRiskSettleInvstPosition(void *this, struct CThostFtdcRiskSettleInvstPositionField *pRiskSettleInvstPosition, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast)
{
    CgoOnRspQryRiskSettleInvstPosition(this, pRiskSettleInvstPosition, pRspInfo, nRequestID, bIsLast);
}

void COnRspQryRiskSettleProductStatus(void *this, struct CThostFtdcRiskSettleProductStatusField *pRiskSettleProductStatus, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast)
{
    CgoOnRspQryRiskSettleProductStatus(this, pRiskSettleProductStatus, pRspInfo, nRequestID, bIsLast);
}

void COnRspQrySPBMFutureParameter(void *this, struct CThostFtdcSPBMFutureParameterField *pSPBMFutureParameter, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast)
{
    CgoOnRspQrySPBMFutureParameter(this, pSPBMFutureParameter, pRspInfo, nRequestID, bIsLast);
}

void COnRspQrySPBMOptionParameter(void *this, struct CThostFtdcSPBMOptionParameterField *pSPBMOptionParameter, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast)
{
    CgoOnRspQrySPBMOptionParameter(this, pSPBMOptionParameter, pRspInfo, nRequestID, bIsLast);
}

void COnRspQrySPBMIntraParameter(void *this, struct CThostFtdcSPBMIntraParameterField *pSPBMIntraParameter, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast)
{
    CgoOnRspQrySPBMIntraParameter(this, pSPBMIntraParameter, pRspInfo, nRequestID, bIsLast);
}

void COnRspQrySPBMInterParameter(void *this, struct CThostFtdcSPBMInterParameterField *pSPBMInterParameter, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast)
{
    CgoOnRspQrySPBMInterParameter(this, pSPBMInterParameter, pRspInfo, nRequestID, bIsLast);
}

void COnRspQrySPBMPortfDefinition(void *this, struct CThostFtdcSPBMPortfDefinitionField *pSPBMPortfDefinition, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast)
{
    CgoOnRspQrySPBMPortfDefinition(this, pSPBMPortfDefinition, pRspInfo, nRequestID, bIsLast);
}

void COnRspQrySPBMInvestorPortfDef(void *this, struct CThostFtdcSPBMInvestorPortfDefField *pSPBMInvestorPortfDef, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast)
{
    CgoOnRspQrySPBMInvestorPortfDef(this, pSPBMInvestorPortfDef, pRspInfo, nRequestID, bIsLast);
}

void COnRspQryInvestorPortfMarginRatio(void *this, struct CThostFtdcInvestorPortfMarginRatioField *pInvestorPortfMarginRatio, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast)
{
    CgoOnRspQryInvestorPortfMarginRatio(this, pInvestorPortfMarginRatio, pRspInfo, nRequestID, bIsLast);
}

void COnRspQryInvestorProdSPBMDetail(void *this, struct CThostFtdcInvestorProdSPBMDetailField *pInvestorProdSPBMDetail, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast)
{
    CgoOnRspQryInvestorProdSPBMDetail(this, pInvestorProdSPBMDetail, pRspInfo, nRequestID, bIsLast);
}

void COnRspQryInvestorCommoditySPMMMargin(void *this, struct CThostFtdcInvestorCommoditySPMMMarginField *pInvestorCommoditySPMMMargin, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast)
{
    CgoOnRspQryInvestorCommoditySPMMMargin(this, pInvestorCommoditySPMMMargin, pRspInfo, nRequestID, bIsLast);
}

void COnRspQryInvestorCommodityGroupSPMMMargin(void *this, struct CThostFtdcInvestorCommodityGroupSPMMMarginField *pInvestorCommodityGroupSPMMMargin, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast)
{
    CgoOnRspQryInvestorCommodityGroupSPMMMargin(this, pInvestorCommodityGroupSPMMMargin, pRspInfo, nRequestID, bIsLast);
}

void COnRspQrySPMMInstParam(void *this, struct CThostFtdcSPMMInstParamField *pSPMMInstParam, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast)
{
    CgoOnRspQrySPMMInstParam(this, pSPMMInstParam, pRspInfo, nRequestID, bIsLast);
}

void COnRspQrySPMMProductParam(void *this, struct CThostFtdcSPMMProductParamField *pSPMMProductParam, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast)
{
    CgoOnRspQrySPMMProductParam(this, pSPMMProductParam, pRspInfo, nRequestID, bIsLast);
}

void COnRspQrySPBMAddOnInterParameter(void *this, struct CThostFtdcSPBMAddOnInterParameterField *pSPBMAddOnInterParameter, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast)
{
    CgoOnRspQrySPBMAddOnInterParameter(this, pSPBMAddOnInterParameter, pRspInfo, nRequestID, bIsLast);
}

void COnRspQryRCAMSCombProductInfo(void *this, struct CThostFtdcRCAMSCombProductInfoField *pRCAMSCombProductInfo, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast)
{
    CgoOnRspQryRCAMSCombProductInfo(this, pRCAMSCombProductInfo, pRspInfo, nRequestID, bIsLast);
}

void COnRspQryRCAMSInstrParameter(void *this, struct CThostFtdcRCAMSInstrParameterField *pRCAMSInstrParameter, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast)
{
    CgoOnRspQryRCAMSInstrParameter(this, pRCAMSInstrParameter, pRspInfo, nRequestID, bIsLast);
}

void COnRspQryRCAMSIntraParameter(void *this, struct CThostFtdcRCAMSIntraParameterField *pRCAMSIntraParameter, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast)
{
    CgoOnRspQryRCAMSIntraParameter(this, pRCAMSIntraParameter, pRspInfo, nRequestID, bIsLast);
}

void COnRspQryRCAMSInterParameter(void *this, struct CThostFtdcRCAMSInterParameterField *pRCAMSInterParameter, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast)
{
    CgoOnRspQryRCAMSInterParameter(this, pRCAMSInterParameter, pRspInfo, nRequestID, bIsLast);
}

void COnRspQryRCAMSShortOptAdjustParam(void *this, struct CThostFtdcRCAMSShortOptAdjustParamField *pRCAMSShortOptAdjustParam, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast)
{
    CgoOnRspQryRCAMSShortOptAdjustParam(this, pRCAMSShortOptAdjustParam, pRspInfo, nRequestID, bIsLast);
}

void COnRspQryRCAMSInvestorCombPosition(void *this, struct CThostFtdcRCAMSInvestorCombPositionField *pRCAMSInvestorCombPosition, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast)
{
    CgoOnRspQryRCAMSInvestorCombPosition(this, pRCAMSInvestorCombPosition, pRspInfo, nRequestID, bIsLast);
}

void COnRspQryInvestorProdRCAMSMargin(void *this, struct CThostFtdcInvestorProdRCAMSMarginField *pInvestorProdRCAMSMargin, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast)
{
    CgoOnRspQryInvestorProdRCAMSMargin(this, pInvestorProdRCAMSMargin, pRspInfo, nRequestID, bIsLast);
}

void COnRspQryRULEInstrParameter(void *this, struct CThostFtdcRULEInstrParameterField *pRULEInstrParameter, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast)
{
    CgoOnRspQryRULEInstrParameter(this, pRULEInstrParameter, pRspInfo, nRequestID, bIsLast);
}

void COnRspQryRULEIntraParameter(void *this, struct CThostFtdcRULEIntraParameterField *pRULEIntraParameter, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast)
{
    CgoOnRspQryRULEIntraParameter(this, pRULEIntraParameter, pRspInfo, nRequestID, bIsLast);
}

void COnRspQryRULEInterParameter(void *this, struct CThostFtdcRULEInterParameterField *pRULEInterParameter, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast)
{
    CgoOnRspQryRULEInterParameter(this, pRULEInterParameter, pRspInfo, nRequestID, bIsLast);
}

void COnRspQryInvestorProdRULEMargin(void *this, struct CThostFtdcInvestorProdRULEMarginField *pInvestorProdRULEMargin, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast)
{
    CgoOnRspQryInvestorProdRULEMargin(this, pInvestorProdRULEMargin, pRspInfo, nRequestID, bIsLast);
}

void COnRspQryInvestorPortfSetting(void *this, struct CThostFtdcInvestorPortfSettingField *pInvestorPortfSetting, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast)
{
    CgoOnRspQryInvestorPortfSetting(this, pInvestorPortfSetting, pRspInfo, nRequestID, bIsLast);
}

void COnRspQryInvestorInfoCommRec(void *this, struct CThostFtdcInvestorInfoCommRecField *pInvestorInfoCommRec, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast)
{
    CgoOnRspQryInvestorInfoCommRec(this, pInvestorInfoCommRec, pRspInfo, nRequestID, bIsLast);
}

void COnRspQryCombLeg(void *this, struct CThostFtdcCombLegField *pCombLeg, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast)
{
    CgoOnRspQryCombLeg(this, pCombLeg, pRspInfo, nRequestID, bIsLast);
}

void COnRspOffsetSetting(void *this, struct CThostFtdcInputOffsetSettingField *pInputOffsetSetting, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast)
{
    CgoOnRspOffsetSetting(this, pInputOffsetSetting, pRspInfo, nRequestID, bIsLast);
}

void COnRspCancelOffsetSetting(void *this, struct CThostFtdcInputOffsetSettingField *pInputOffsetSetting, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast)
{
    CgoOnRspCancelOffsetSetting(this, pInputOffsetSetting, pRspInfo, nRequestID, bIsLast);
}

void COnRtnOffsetSetting(void *this, struct CThostFtdcOffsetSettingField *pOffsetSetting)
{
    CgoOnRtnOffsetSetting(this, pOffsetSetting);
}

void COnErrRtnOffsetSetting(void *this, struct CThostFtdcInputOffsetSettingField *pInputOffsetSetting, struct CThostFtdcRspInfoField *pRspInfo)
{
    CgoOnErrRtnOffsetSetting(this, pInputOffsetSetting, pRspInfo);
}

void COnErrRtnCancelOffsetSetting(void *this, struct CThostFtdcCancelOffsetSettingField *pCancelOffsetSetting, struct CThostFtdcRspInfoField *pRspInfo)
{
    CgoOnErrRtnCancelOffsetSetting(this, pCancelOffsetSetting, pRspInfo);
}

void COnRspQryOffsetSetting(void *this, struct CThostFtdcOffsetSettingField *pOffsetSetting, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast)
{
    CgoOnRspQryOffsetSetting(this, pOffsetSetting, pRspInfo, nRequestID, bIsLast);
}

void COnRspGenSMSCode(void *this, struct CThostFtdcRspGenSMSCodeField *pRspGenSMSCode, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast)
{
    CgoOnRspGenSMSCode(this, pRspGenSMSCode, pRspInfo, nRequestID, bIsLast);
}

void COnRspSpdApply(void *this, struct CThostFtdcInputSpdApplyField *pInputSpdApply, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast)
{
    CgoOnRspSpdApply(this, pInputSpdApply, pRspInfo, nRequestID, bIsLast);
}

void COnRspSpdApplyAction(void *this, struct CThostFtdcInputSpdApplyActionField *pInputSpdApplyAction, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast)
{
    CgoOnRspSpdApplyAction(this, pInputSpdApplyAction, pRspInfo, nRequestID, bIsLast);
}

void COnRspQrySpdApply(void *this, struct CThostFtdcSpdApplyField *pSpdApply, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast)
{
    CgoOnRspQrySpdApply(this, pSpdApply, pRspInfo, nRequestID, bIsLast);
}

void COnRtnSpdApply(void *this, struct CThostFtdcSpdApplyField *pSpdApply)
{
    CgoOnRtnSpdApply(this, pSpdApply);
}

void COnErrRtnSpdApply(void *this, struct CThostFtdcInputSpdApplyField *pInputSpdApply, struct CThostFtdcRspInfoField *pRspInfo)
{
    CgoOnErrRtnSpdApply(this, pInputSpdApply, pRspInfo);
}

void COnErrRtnSpdApplyAction(void *this, struct CThostFtdcSpdApplyActionField *pSpdApplyAction, struct CThostFtdcRspInfoField *pRspInfo)
{
    CgoOnErrRtnSpdApplyAction(this, pSpdApplyAction, pRspInfo);
}

void COnRspHedgeCfm(void *this, struct CThostFtdcInputHedgeCfmField *pInputHedgeCfm, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast)
{
    CgoOnRspHedgeCfm(this, pInputHedgeCfm, pRspInfo, nRequestID, bIsLast);
}

void COnRspHedgeCfmAction(void *this, struct CThostFtdcInputHedgeCfmActionField *pInputHedgeCfmAction, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast)
{
    CgoOnRspHedgeCfmAction(this, pInputHedgeCfmAction, pRspInfo, nRequestID, bIsLast);
}

void COnRspQryHedgeCfm(void *this, struct CThostFtdcHedgeCfmField *pHedgeCfm, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast)
{
    CgoOnRspQryHedgeCfm(this, pHedgeCfm, pRspInfo, nRequestID, bIsLast);
}

void COnRtnHedgeCfm(void *this, struct CThostFtdcHedgeCfmField *pHedgeCfm)
{
    CgoOnRtnHedgeCfm(this, pHedgeCfm);
}

void COnErrRtnHedgeCfm(void *this, struct CThostFtdcInputHedgeCfmField *pInputHedgeCfm, struct CThostFtdcRspInfoField *pRspInfo)
{
    CgoOnErrRtnHedgeCfm(this, pInputHedgeCfm, pRspInfo);
}

void COnErrRtnHedgeCfmAction(void *this, struct CThostFtdcHedgeCfmActionField *pHedgeCfmAction, struct CThostFtdcRspInfoField *pRspInfo)
{
    CgoOnErrRtnHedgeCfmAction(this, pHedgeCfmAction, pRspInfo);
}
