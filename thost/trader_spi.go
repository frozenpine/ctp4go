package thost

type TraderSpi interface {
    // 当客户端与交易后台建立起通信连接时（还未登录前），该方法被调用。
    OnFrontConnected()

    // 当客户端与交易后台通信连接断开时，该方法被调用。当发生这个情况后，API会自动重新连接，客户端可不做处理。
    //   @param nReason 错误原因
    //     0x1001 网络读失败
    //     0x1002 网络写失败
    //     0x2001 接收心跳超时
    //     0x2002 发送心跳失败
    //     0x2003 收到错误报文
    OnFrontDisconnected(Reason int)

    // 心跳超时警告。当长时间未收到报文时，该方法被调用。
    //   @param nTimeLapse 距离上次接收报文的时间
    OnHeartBeatWarning(TimeLapse int)

    // 客户端认证响应
    OnRspAuthenticate(RspAuthenticateField *CThostFtdcRspAuthenticateField, RspInfo *CThostFtdcRspInfoField, RequestID int, IsLast bool)

    // 该方法在处理私有流之前被调用
    //   @param nSeqNo 即将被处理的私有流的序号
    OnRtnPrivateSeqNo(SeqNo int)

    // 登录请求响应
    OnRspUserLogin(RspUserLogin *CThostFtdcRspUserLoginField, RspInfo *CThostFtdcRspInfoField, RequestID int, IsLast bool)

    // 登出请求响应
    OnRspUserLogout(UserLogout *CThostFtdcUserLogoutField, RspInfo *CThostFtdcRspInfoField, RequestID int, IsLast bool)

    // 用户口令更新请求响应
    OnRspUserPasswordUpdate(UserPasswordUpdate *CThostFtdcUserPasswordUpdateField, RspInfo *CThostFtdcRspInfoField, RequestID int, IsLast bool)

    // 资金账户口令更新请求响应
    OnRspTradingAccountPasswordUpdate(TradingAccountPasswordUpdate *CThostFtdcTradingAccountPasswordUpdateField, RspInfo *CThostFtdcRspInfoField, RequestID int, IsLast bool)

    // 查询用户当前支持的认证模式的回复
    OnRspUserAuthMethod(RspUserAuthMethod *CThostFtdcRspUserAuthMethodField, RspInfo *CThostFtdcRspInfoField, RequestID int, IsLast bool)

    // 获取图形验证码请求的回复
    OnRspGenUserCaptcha(RspGenUserCaptcha *CThostFtdcRspGenUserCaptchaField, RspInfo *CThostFtdcRspInfoField, RequestID int, IsLast bool)

    // 获取短信验证码请求的回复
    OnRspGenUserText(RspGenUserText *CThostFtdcRspGenUserTextField, RspInfo *CThostFtdcRspInfoField, RequestID int, IsLast bool)

    // 报单录入请求响应
    OnRspOrderInsert(InputOrder *CThostFtdcInputOrderField, RspInfo *CThostFtdcRspInfoField, RequestID int, IsLast bool)

    // 预埋单录入请求响应
    OnRspParkedOrderInsert(ParkedOrder *CThostFtdcParkedOrderField, RspInfo *CThostFtdcRspInfoField, RequestID int, IsLast bool)

    // 预埋撤单录入请求响应
    OnRspParkedOrderAction(ParkedOrderAction *CThostFtdcParkedOrderActionField, RspInfo *CThostFtdcRspInfoField, RequestID int, IsLast bool)

    // 报单操作请求响应
    OnRspOrderAction(InputOrderAction *CThostFtdcInputOrderActionField, RspInfo *CThostFtdcRspInfoField, RequestID int, IsLast bool)

    // 查询最大报单数量响应
    OnRspQryMaxOrderVolume(QryMaxOrderVolume *CThostFtdcQryMaxOrderVolumeField, RspInfo *CThostFtdcRspInfoField, RequestID int, IsLast bool)

    // 投资者结算结果确认响应
    OnRspSettlementInfoConfirm(SettlementInfoConfirm *CThostFtdcSettlementInfoConfirmField, RspInfo *CThostFtdcRspInfoField, RequestID int, IsLast bool)

    // 删除预埋单响应
    OnRspRemoveParkedOrder(RemoveParkedOrder *CThostFtdcRemoveParkedOrderField, RspInfo *CThostFtdcRspInfoField, RequestID int, IsLast bool)

    // 删除预埋撤单响应
    OnRspRemoveParkedOrderAction(RemoveParkedOrderAction *CThostFtdcRemoveParkedOrderActionField, RspInfo *CThostFtdcRspInfoField, RequestID int, IsLast bool)

    // 执行宣告录入请求响应
    OnRspExecOrderInsert(InputExecOrder *CThostFtdcInputExecOrderField, RspInfo *CThostFtdcRspInfoField, RequestID int, IsLast bool)

    // 执行宣告操作请求响应
    OnRspExecOrderAction(InputExecOrderAction *CThostFtdcInputExecOrderActionField, RspInfo *CThostFtdcRspInfoField, RequestID int, IsLast bool)

    // 询价录入请求响应
    OnRspForQuoteInsert(InputForQuote *CThostFtdcInputForQuoteField, RspInfo *CThostFtdcRspInfoField, RequestID int, IsLast bool)

    // 报价录入请求响应
    OnRspQuoteInsert(InputQuote *CThostFtdcInputQuoteField, RspInfo *CThostFtdcRspInfoField, RequestID int, IsLast bool)

    // 报价操作请求响应
    OnRspQuoteAction(InputQuoteAction *CThostFtdcInputQuoteActionField, RspInfo *CThostFtdcRspInfoField, RequestID int, IsLast bool)

    // 批量报单操作请求响应
    OnRspBatchOrderAction(InputBatchOrderAction *CThostFtdcInputBatchOrderActionField, RspInfo *CThostFtdcRspInfoField, RequestID int, IsLast bool)

    // 期权自对冲录入请求响应
    OnRspOptionSelfCloseInsert(InputOptionSelfClose *CThostFtdcInputOptionSelfCloseField, RspInfo *CThostFtdcRspInfoField, RequestID int, IsLast bool)

    // 期权自对冲操作请求响应
    OnRspOptionSelfCloseAction(InputOptionSelfCloseAction *CThostFtdcInputOptionSelfCloseActionField, RspInfo *CThostFtdcRspInfoField, RequestID int, IsLast bool)

    // 申请组合录入请求响应
    OnRspCombActionInsert(InputCombAction *CThostFtdcInputCombActionField, RspInfo *CThostFtdcRspInfoField, RequestID int, IsLast bool)

    // 请求查询报单响应
    OnRspQryOrder(Order *CThostFtdcOrderField, RspInfo *CThostFtdcRspInfoField, RequestID int, IsLast bool)

    // 请求查询成交响应
    OnRspQryTrade(Trade *CThostFtdcTradeField, RspInfo *CThostFtdcRspInfoField, RequestID int, IsLast bool)

    // 请求查询投资者持仓响应
    OnRspQryInvestorPosition(InvestorPosition *CThostFtdcInvestorPositionField, RspInfo *CThostFtdcRspInfoField, RequestID int, IsLast bool)

    // 请求查询资金账户响应
    OnRspQryTradingAccount(TradingAccount *CThostFtdcTradingAccountField, RspInfo *CThostFtdcRspInfoField, RequestID int, IsLast bool)

    // 请求查询投资者响应
    OnRspQryInvestor(Investor *CThostFtdcInvestorField, RspInfo *CThostFtdcRspInfoField, RequestID int, IsLast bool)

    // 请求查询交易编码响应
    OnRspQryTradingCode(TradingCode *CThostFtdcTradingCodeField, RspInfo *CThostFtdcRspInfoField, RequestID int, IsLast bool)

    // 请求查询合约保证金率响应
    OnRspQryInstrumentMarginRate(InstrumentMarginRate *CThostFtdcInstrumentMarginRateField, RspInfo *CThostFtdcRspInfoField, RequestID int, IsLast bool)

    // 请求查询合约手续费率响应
    OnRspQryInstrumentCommissionRate(InstrumentCommissionRate *CThostFtdcInstrumentCommissionRateField, RspInfo *CThostFtdcRspInfoField, RequestID int, IsLast bool)

    // 请求查询用户会话响应
    OnRspQryUserSession(UserSession *CThostFtdcUserSessionField, RspInfo *CThostFtdcRspInfoField, RequestID int, IsLast bool)

    // 请求查询交易所响应
    OnRspQryExchange(Exchange *CThostFtdcExchangeField, RspInfo *CThostFtdcRspInfoField, RequestID int, IsLast bool)

    // 请求查询产品响应
    OnRspQryProduct(Product *CThostFtdcProductField, RspInfo *CThostFtdcRspInfoField, RequestID int, IsLast bool)

    // 请求查询合约响应
    OnRspQryInstrument(Instrument *CThostFtdcInstrumentField, RspInfo *CThostFtdcRspInfoField, RequestID int, IsLast bool)

    // 请求查询行情响应
    OnRspQryDepthMarketData(DepthMarketData *CThostFtdcDepthMarketDataField, RspInfo *CThostFtdcRspInfoField, RequestID int, IsLast bool)

    // 请求查询交易员报盘机响应
    OnRspQryTraderOffer(TraderOffer *CThostFtdcTraderOfferField, RspInfo *CThostFtdcRspInfoField, RequestID int, IsLast bool)

    // 请求查询投资者结算结果响应
    OnRspQrySettlementInfo(SettlementInfo *CThostFtdcSettlementInfoField, RspInfo *CThostFtdcRspInfoField, RequestID int, IsLast bool)

    // 请求查询转帐银行响应
    OnRspQryTransferBank(TransferBank *CThostFtdcTransferBankField, RspInfo *CThostFtdcRspInfoField, RequestID int, IsLast bool)

    // 请求查询投资者持仓明细响应
    OnRspQryInvestorPositionDetail(InvestorPositionDetail *CThostFtdcInvestorPositionDetailField, RspInfo *CThostFtdcRspInfoField, RequestID int, IsLast bool)

    // 请求查询客户通知响应
    OnRspQryNotice(Notice *CThostFtdcNoticeField, RspInfo *CThostFtdcRspInfoField, RequestID int, IsLast bool)

    // 请求查询结算信息确认响应
    OnRspQrySettlementInfoConfirm(SettlementInfoConfirm *CThostFtdcSettlementInfoConfirmField, RspInfo *CThostFtdcRspInfoField, RequestID int, IsLast bool)

    // 请求查询投资者持仓明细响应
    OnRspQryInvestorPositionCombineDetail(InvestorPositionCombineDetail *CThostFtdcInvestorPositionCombineDetailField, RspInfo *CThostFtdcRspInfoField, RequestID int, IsLast bool)

    // 查询保证金监管系统经纪公司资金账户密钥响应
    OnRspQryCFMMCTradingAccountKey(CFMMCTradingAccountKey *CThostFtdcCFMMCTradingAccountKeyField, RspInfo *CThostFtdcRspInfoField, RequestID int, IsLast bool)

    // 请求查询仓单折抵信息响应
    OnRspQryEWarrantOffset(EWarrantOffset *CThostFtdcEWarrantOffsetField, RspInfo *CThostFtdcRspInfoField, RequestID int, IsLast bool)

    // 请求查询投资者品种跨品种保证金响应
    OnRspQryInvestorProductGroupMargin(InvestorProductGroupMargin *CThostFtdcInvestorProductGroupMarginField, RspInfo *CThostFtdcRspInfoField, RequestID int, IsLast bool)

    // 请求查询交易所保证金率响应
    OnRspQryExchangeMarginRate(ExchangeMarginRate *CThostFtdcExchangeMarginRateField, RspInfo *CThostFtdcRspInfoField, RequestID int, IsLast bool)

    // 请求查询交易所调整保证金率响应
    OnRspQryExchangeMarginRateAdjust(ExchangeMarginRateAdjust *CThostFtdcExchangeMarginRateAdjustField, RspInfo *CThostFtdcRspInfoField, RequestID int, IsLast bool)

    // 请求查询汇率响应
    OnRspQryExchangeRate(ExchangeRate *CThostFtdcExchangeRateField, RspInfo *CThostFtdcRspInfoField, RequestID int, IsLast bool)

    // 请求查询二级代理操作员银期权限响应
    OnRspQrySecAgentACIDMap(SecAgentACIDMap *CThostFtdcSecAgentACIDMapField, RspInfo *CThostFtdcRspInfoField, RequestID int, IsLast bool)

    // 请求查询产品报价汇率
    OnRspQryProductExchRate(ProductExchRate *CThostFtdcProductExchRateField, RspInfo *CThostFtdcRspInfoField, RequestID int, IsLast bool)

    // 请求查询产品组
    OnRspQryProductGroup(ProductGroup *CThostFtdcProductGroupField, RspInfo *CThostFtdcRspInfoField, RequestID int, IsLast bool)

    // 请求查询做市商合约手续费率响应
    OnRspQryMMInstrumentCommissionRate(MMInstrumentCommissionRate *CThostFtdcMMInstrumentCommissionRateField, RspInfo *CThostFtdcRspInfoField, RequestID int, IsLast bool)

    // 请求查询做市商期权合约手续费响应
    OnRspQryMMOptionInstrCommRate(MMOptionInstrCommRate *CThostFtdcMMOptionInstrCommRateField, RspInfo *CThostFtdcRspInfoField, RequestID int, IsLast bool)

    // 请求查询报单手续费响应
    OnRspQryInstrumentOrderCommRate(InstrumentOrderCommRate *CThostFtdcInstrumentOrderCommRateField, RspInfo *CThostFtdcRspInfoField, RequestID int, IsLast bool)

    // 请求查询资金账户响应
    OnRspQrySecAgentTradingAccount(TradingAccount *CThostFtdcTradingAccountField, RspInfo *CThostFtdcRspInfoField, RequestID int, IsLast bool)

    // 请求查询二级代理商资金校验模式响应
    OnRspQrySecAgentCheckMode(SecAgentCheckMode *CThostFtdcSecAgentCheckModeField, RspInfo *CThostFtdcRspInfoField, RequestID int, IsLast bool)

    // 请求查询二级代理商信息响应
    OnRspQrySecAgentTradeInfo(SecAgentTradeInfo *CThostFtdcSecAgentTradeInfoField, RspInfo *CThostFtdcRspInfoField, RequestID int, IsLast bool)

    // 请求查询期权交易成本响应
    OnRspQryOptionInstrTradeCost(OptionInstrTradeCost *CThostFtdcOptionInstrTradeCostField, RspInfo *CThostFtdcRspInfoField, RequestID int, IsLast bool)

    // 请求查询期权合约手续费响应
    OnRspQryOptionInstrCommRate(OptionInstrCommRate *CThostFtdcOptionInstrCommRateField, RspInfo *CThostFtdcRspInfoField, RequestID int, IsLast bool)

    // 请求查询执行宣告响应
    OnRspQryExecOrder(ExecOrder *CThostFtdcExecOrderField, RspInfo *CThostFtdcRspInfoField, RequestID int, IsLast bool)

    // 请求查询询价响应
    OnRspQryForQuote(ForQuote *CThostFtdcForQuoteField, RspInfo *CThostFtdcRspInfoField, RequestID int, IsLast bool)

    // 请求查询报价响应
    OnRspQryQuote(Quote *CThostFtdcQuoteField, RspInfo *CThostFtdcRspInfoField, RequestID int, IsLast bool)

    // 请求查询期权自对冲响应
    OnRspQryOptionSelfClose(OptionSelfClose *CThostFtdcOptionSelfCloseField, RspInfo *CThostFtdcRspInfoField, RequestID int, IsLast bool)

    // 请求查询投资单元响应
    OnRspQryInvestUnit(InvestUnit *CThostFtdcInvestUnitField, RspInfo *CThostFtdcRspInfoField, RequestID int, IsLast bool)

    // 请求查询组合合约安全系数响应
    OnRspQryCombInstrumentGuard(CombInstrumentGuard *CThostFtdcCombInstrumentGuardField, RspInfo *CThostFtdcRspInfoField, RequestID int, IsLast bool)

    // 请求查询申请组合响应
    OnRspQryCombAction(CombAction *CThostFtdcCombActionField, RspInfo *CThostFtdcRspInfoField, RequestID int, IsLast bool)

    // 请求查询转帐流水响应
    OnRspQryTransferSerial(TransferSerial *CThostFtdcTransferSerialField, RspInfo *CThostFtdcRspInfoField, RequestID int, IsLast bool)

    // 请求查询银期签约关系响应
    OnRspQryAccountregister(Accountregister *CThostFtdcAccountregisterField, RspInfo *CThostFtdcRspInfoField, RequestID int, IsLast bool)

    // 错误应答
    OnRspError(RspInfo *CThostFtdcRspInfoField, RequestID int, IsLast bool)

    // 报单通知
    OnRtnOrder(Order *CThostFtdcOrderField)

    // 成交通知
    OnRtnTrade(Trade *CThostFtdcTradeField)

    // 报单录入错误回报
    OnErrRtnOrderInsert(InputOrder *CThostFtdcInputOrderField, RspInfo *CThostFtdcRspInfoField)

    // 报单操作错误回报
    OnErrRtnOrderAction(OrderAction *CThostFtdcOrderActionField, RspInfo *CThostFtdcRspInfoField)

    // 合约交易状态通知
    OnRtnInstrumentStatus(InstrumentStatus *CThostFtdcInstrumentStatusField)

    // 交易所公告通知
    OnRtnBulletin(Bulletin *CThostFtdcBulletinField)

    // 交易通知
    OnRtnTradingNotice(TradingNoticeInfo *CThostFtdcTradingNoticeInfoField)

    // 提示条件单校验错误
    OnRtnErrorConditionalOrder(ErrorConditionalOrder *CThostFtdcErrorConditionalOrderField)

    // 执行宣告通知
    OnRtnExecOrder(ExecOrder *CThostFtdcExecOrderField)

    // 执行宣告录入错误回报
    OnErrRtnExecOrderInsert(InputExecOrder *CThostFtdcInputExecOrderField, RspInfo *CThostFtdcRspInfoField)

    // 执行宣告操作错误回报
    OnErrRtnExecOrderAction(ExecOrderAction *CThostFtdcExecOrderActionField, RspInfo *CThostFtdcRspInfoField)

    // 询价录入错误回报
    OnErrRtnForQuoteInsert(InputForQuote *CThostFtdcInputForQuoteField, RspInfo *CThostFtdcRspInfoField)

    // 报价通知
    OnRtnQuote(Quote *CThostFtdcQuoteField)

    // 报价录入错误回报
    OnErrRtnQuoteInsert(InputQuote *CThostFtdcInputQuoteField, RspInfo *CThostFtdcRspInfoField)

    // 报价操作错误回报
    OnErrRtnQuoteAction(QuoteAction *CThostFtdcQuoteActionField, RspInfo *CThostFtdcRspInfoField)

    // 询价通知
    OnRtnForQuoteRsp(ForQuoteRsp *CThostFtdcForQuoteRspField)

    // 保证金监控中心用户令牌
    OnRtnCFMMCTradingAccountToken(CFMMCTradingAccountToken *CThostFtdcCFMMCTradingAccountTokenField)

    // 批量报单操作错误回报
    OnErrRtnBatchOrderAction(BatchOrderAction *CThostFtdcBatchOrderActionField, RspInfo *CThostFtdcRspInfoField)

    // 期权自对冲通知
    OnRtnOptionSelfClose(OptionSelfClose *CThostFtdcOptionSelfCloseField)

    // 期权自对冲录入错误回报
    OnErrRtnOptionSelfCloseInsert(InputOptionSelfClose *CThostFtdcInputOptionSelfCloseField, RspInfo *CThostFtdcRspInfoField)

    // 期权自对冲操作错误回报
    OnErrRtnOptionSelfCloseAction(OptionSelfCloseAction *CThostFtdcOptionSelfCloseActionField, RspInfo *CThostFtdcRspInfoField)

    // 申请组合通知
    OnRtnCombAction(CombAction *CThostFtdcCombActionField)

    // 申请组合录入错误回报
    OnErrRtnCombActionInsert(InputCombAction *CThostFtdcInputCombActionField, RspInfo *CThostFtdcRspInfoField)

    // 请求查询签约银行响应
    OnRspQryContractBank(ContractBank *CThostFtdcContractBankField, RspInfo *CThostFtdcRspInfoField, RequestID int, IsLast bool)

    // 请求查询预埋单响应
    OnRspQryParkedOrder(ParkedOrder *CThostFtdcParkedOrderField, RspInfo *CThostFtdcRspInfoField, RequestID int, IsLast bool)

    // 请求查询预埋撤单响应
    OnRspQryParkedOrderAction(ParkedOrderAction *CThostFtdcParkedOrderActionField, RspInfo *CThostFtdcRspInfoField, RequestID int, IsLast bool)

    // 请求查询交易通知响应
    OnRspQryTradingNotice(TradingNotice *CThostFtdcTradingNoticeField, RspInfo *CThostFtdcRspInfoField, RequestID int, IsLast bool)

    // 请求查询经纪公司交易参数响应
    OnRspQryBrokerTradingParams(BrokerTradingParams *CThostFtdcBrokerTradingParamsField, RspInfo *CThostFtdcRspInfoField, RequestID int, IsLast bool)

    // 请求查询经纪公司交易算法响应
    OnRspQryBrokerTradingAlgos(BrokerTradingAlgos *CThostFtdcBrokerTradingAlgosField, RspInfo *CThostFtdcRspInfoField, RequestID int, IsLast bool)

    // 请求查询监控中心用户令牌
    OnRspQueryCFMMCTradingAccountToken(QueryCFMMCTradingAccountToken *CThostFtdcQueryCFMMCTradingAccountTokenField, RspInfo *CThostFtdcRspInfoField, RequestID int, IsLast bool)

    // 银行发起银行资金转期货通知
    OnRtnFromBankToFutureByBank(RspTransfer *CThostFtdcRspTransferField)

    // 银行发起期货资金转银行通知
    OnRtnFromFutureToBankByBank(RspTransfer *CThostFtdcRspTransferField)

    // 银行发起冲正银行转期货通知
    OnRtnRepealFromBankToFutureByBank(RspRepeal *CThostFtdcRspRepealField)

    // 银行发起冲正期货转银行通知
    OnRtnRepealFromFutureToBankByBank(RspRepeal *CThostFtdcRspRepealField)

    // 期货发起银行资金转期货通知
    OnRtnFromBankToFutureByFuture(RspTransfer *CThostFtdcRspTransferField)

    // 期货发起期货资金转银行通知
    OnRtnFromFutureToBankByFuture(RspTransfer *CThostFtdcRspTransferField)

    // 系统运行时期货端手工发起冲正银行转期货请求，银行处理完毕后报盘发回的通知
    OnRtnRepealFromBankToFutureByFutureManual(RspRepeal *CThostFtdcRspRepealField)

    // 系统运行时期货端手工发起冲正期货转银行请求，银行处理完毕后报盘发回的通知
    OnRtnRepealFromFutureToBankByFutureManual(RspRepeal *CThostFtdcRspRepealField)

    // 期货发起查询银行余额通知
    OnRtnQueryBankBalanceByFuture(NotifyQueryAccount *CThostFtdcNotifyQueryAccountField)

    // 期货发起银行资金转期货错误回报
    OnErrRtnBankToFutureByFuture(ReqTransfer *CThostFtdcReqTransferField, RspInfo *CThostFtdcRspInfoField)

    // 期货发起期货资金转银行错误回报
    OnErrRtnFutureToBankByFuture(ReqTransfer *CThostFtdcReqTransferField, RspInfo *CThostFtdcRspInfoField)

    // 系统运行时期货端手工发起冲正银行转期货错误回报
    OnErrRtnRepealBankToFutureByFutureManual(ReqRepeal *CThostFtdcReqRepealField, RspInfo *CThostFtdcRspInfoField)

    // 系统运行时期货端手工发起冲正期货转银行错误回报
    OnErrRtnRepealFutureToBankByFutureManual(ReqRepeal *CThostFtdcReqRepealField, RspInfo *CThostFtdcRspInfoField)

    // 期货发起查询银行余额错误回报
    OnErrRtnQueryBankBalanceByFuture(ReqQueryAccount *CThostFtdcReqQueryAccountField, RspInfo *CThostFtdcRspInfoField)

    // 期货发起冲正银行转期货请求，银行处理完毕后报盘发回的通知
    OnRtnRepealFromBankToFutureByFuture(RspRepeal *CThostFtdcRspRepealField)

    // 期货发起冲正期货转银行请求，银行处理完毕后报盘发回的通知
    OnRtnRepealFromFutureToBankByFuture(RspRepeal *CThostFtdcRspRepealField)

    // 期货发起银行资金转期货应答
    OnRspFromBankToFutureByFuture(ReqTransfer *CThostFtdcReqTransferField, RspInfo *CThostFtdcRspInfoField, RequestID int, IsLast bool)

    // 期货发起期货资金转银行应答
    OnRspFromFutureToBankByFuture(ReqTransfer *CThostFtdcReqTransferField, RspInfo *CThostFtdcRspInfoField, RequestID int, IsLast bool)

    // 期货发起查询银行余额应答
    OnRspQueryBankAccountMoneyByFuture(ReqQueryAccount *CThostFtdcReqQueryAccountField, RspInfo *CThostFtdcRspInfoField, RequestID int, IsLast bool)

    // 银行发起银期开户通知
    OnRtnOpenAccountByBank(OpenAccount *CThostFtdcOpenAccountField)

    // 银行发起银期销户通知
    OnRtnCancelAccountByBank(CancelAccount *CThostFtdcCancelAccountField)

    // 银行发起变更银行账号通知
    OnRtnChangeAccountByBank(ChangeAccount *CThostFtdcChangeAccountField)

    // 请求查询分类合约响应
    OnRspQryClassifiedInstrument(Instrument *CThostFtdcInstrumentField, RspInfo *CThostFtdcRspInfoField, RequestID int, IsLast bool)

    // 请求组合优惠比例响应
    OnRspQryCombPromotionParam(CombPromotionParam *CThostFtdcCombPromotionParamField, RspInfo *CThostFtdcRspInfoField, RequestID int, IsLast bool)

    // 投资者风险结算持仓查询响应
    OnRspQryRiskSettleInvstPosition(RiskSettleInvstPosition *CThostFtdcRiskSettleInvstPositionField, RspInfo *CThostFtdcRspInfoField, RequestID int, IsLast bool)

    // 风险结算产品查询响应
    OnRspQryRiskSettleProductStatus(RiskSettleProductStatus *CThostFtdcRiskSettleProductStatusField, RspInfo *CThostFtdcRspInfoField, RequestID int, IsLast bool)

    // SPBM期货合约参数查询响应
    OnRspQrySPBMFutureParameter(SPBMFutureParameter *CThostFtdcSPBMFutureParameterField, RspInfo *CThostFtdcRspInfoField, RequestID int, IsLast bool)

    // SPBM期权合约参数查询响应
    OnRspQrySPBMOptionParameter(SPBMOptionParameter *CThostFtdcSPBMOptionParameterField, RspInfo *CThostFtdcRspInfoField, RequestID int, IsLast bool)

    // SPBM品种内对锁仓折扣参数查询响应
    OnRspQrySPBMIntraParameter(SPBMIntraParameter *CThostFtdcSPBMIntraParameterField, RspInfo *CThostFtdcRspInfoField, RequestID int, IsLast bool)

    // SPBM跨品种抵扣参数查询响应
    OnRspQrySPBMInterParameter(SPBMInterParameter *CThostFtdcSPBMInterParameterField, RspInfo *CThostFtdcRspInfoField, RequestID int, IsLast bool)

    // SPBM组合保证金套餐查询响应
    OnRspQrySPBMPortfDefinition(SPBMPortfDefinition *CThostFtdcSPBMPortfDefinitionField, RspInfo *CThostFtdcRspInfoField, RequestID int, IsLast bool)

    // 投资者SPBM套餐选择查询响应
    OnRspQrySPBMInvestorPortfDef(SPBMInvestorPortfDef *CThostFtdcSPBMInvestorPortfDefField, RspInfo *CThostFtdcRspInfoField, RequestID int, IsLast bool)

    // 投资者新型组合保证金系数查询响应
    OnRspQryInvestorPortfMarginRatio(InvestorPortfMarginRatio *CThostFtdcInvestorPortfMarginRatioField, RspInfo *CThostFtdcRspInfoField, RequestID int, IsLast bool)

    // 投资者产品SPBM明细查询响应
    OnRspQryInvestorProdSPBMDetail(InvestorProdSPBMDetail *CThostFtdcInvestorProdSPBMDetailField, RspInfo *CThostFtdcRspInfoField, RequestID int, IsLast bool)

    // 投资者商品组SPMM记录查询响应
    OnRspQryInvestorCommoditySPMMMargin(InvestorCommoditySPMMMargin *CThostFtdcInvestorCommoditySPMMMarginField, RspInfo *CThostFtdcRspInfoField, RequestID int, IsLast bool)

    // 投资者商品群SPMM记录查询响应
    OnRspQryInvestorCommodityGroupSPMMMargin(InvestorCommodityGroupSPMMMargin *CThostFtdcInvestorCommodityGroupSPMMMarginField, RspInfo *CThostFtdcRspInfoField, RequestID int, IsLast bool)

    // SPMM合约参数查询响应
    OnRspQrySPMMInstParam(SPMMInstParam *CThostFtdcSPMMInstParamField, RspInfo *CThostFtdcRspInfoField, RequestID int, IsLast bool)

    // SPMM产品参数查询响应
    OnRspQrySPMMProductParam(SPMMProductParam *CThostFtdcSPMMProductParamField, RspInfo *CThostFtdcRspInfoField, RequestID int, IsLast bool)

    // SPBM附加跨品种抵扣参数查询响应
    OnRspQrySPBMAddOnInterParameter(SPBMAddOnInterParameter *CThostFtdcSPBMAddOnInterParameterField, RspInfo *CThostFtdcRspInfoField, RequestID int, IsLast bool)

    // RCAMS产品组合信息查询响应
    OnRspQryRCAMSCombProductInfo(RCAMSCombProductInfo *CThostFtdcRCAMSCombProductInfoField, RspInfo *CThostFtdcRspInfoField, RequestID int, IsLast bool)

    // RCAMS同合约风险对冲参数查询响应
    OnRspQryRCAMSInstrParameter(RCAMSInstrParameter *CThostFtdcRCAMSInstrParameterField, RspInfo *CThostFtdcRspInfoField, RequestID int, IsLast bool)

    // RCAMS品种内风险对冲参数查询响应
    OnRspQryRCAMSIntraParameter(RCAMSIntraParameter *CThostFtdcRCAMSIntraParameterField, RspInfo *CThostFtdcRspInfoField, RequestID int, IsLast bool)

    // RCAMS跨品种风险折抵参数查询响应
    OnRspQryRCAMSInterParameter(RCAMSInterParameter *CThostFtdcRCAMSInterParameterField, RspInfo *CThostFtdcRspInfoField, RequestID int, IsLast bool)

    // RCAMS空头期权风险调整参数查询响应
    OnRspQryRCAMSShortOptAdjustParam(RCAMSShortOptAdjustParam *CThostFtdcRCAMSShortOptAdjustParamField, RspInfo *CThostFtdcRspInfoField, RequestID int, IsLast bool)

    // RCAMS策略组合持仓查询响应
    OnRspQryRCAMSInvestorCombPosition(RCAMSInvestorCombPosition *CThostFtdcRCAMSInvestorCombPositionField, RspInfo *CThostFtdcRspInfoField, RequestID int, IsLast bool)

    // 投资者品种RCAMS保证金查询响应
    OnRspQryInvestorProdRCAMSMargin(InvestorProdRCAMSMargin *CThostFtdcInvestorProdRCAMSMarginField, RspInfo *CThostFtdcRspInfoField, RequestID int, IsLast bool)

    // RULE合约保证金参数查询响应
    OnRspQryRULEInstrParameter(RULEInstrParameter *CThostFtdcRULEInstrParameterField, RspInfo *CThostFtdcRspInfoField, RequestID int, IsLast bool)

    // RULE品种内对锁仓折扣参数查询响应
    OnRspQryRULEIntraParameter(RULEIntraParameter *CThostFtdcRULEIntraParameterField, RspInfo *CThostFtdcRspInfoField, RequestID int, IsLast bool)

    // RULE跨品种抵扣参数查询响应
    OnRspQryRULEInterParameter(RULEInterParameter *CThostFtdcRULEInterParameterField, RspInfo *CThostFtdcRspInfoField, RequestID int, IsLast bool)

    // 投资者产品RULE保证金查询响应
    OnRspQryInvestorProdRULEMargin(InvestorProdRULEMargin *CThostFtdcInvestorProdRULEMarginField, RspInfo *CThostFtdcRspInfoField, RequestID int, IsLast bool)

    // 投资者新型组合保证金开关查询响应
    OnRspQryInvestorPortfSetting(InvestorPortfSetting *CThostFtdcInvestorPortfSettingField, RspInfo *CThostFtdcRspInfoField, RequestID int, IsLast bool)

    // 投资者申报费阶梯收取记录查询响应
    OnRspQryInvestorInfoCommRec(InvestorInfoCommRec *CThostFtdcInvestorInfoCommRecField, RspInfo *CThostFtdcRspInfoField, RequestID int, IsLast bool)

    // 组合腿信息查询响应
    OnRspQryCombLeg(CombLeg *CThostFtdcCombLegField, RspInfo *CThostFtdcRspInfoField, RequestID int, IsLast bool)

    // 对冲设置请求响应
    OnRspOffsetSetting(InputOffsetSetting *CThostFtdcInputOffsetSettingField, RspInfo *CThostFtdcRspInfoField, RequestID int, IsLast bool)

    // 对冲设置撤销请求响应
    OnRspCancelOffsetSetting(InputOffsetSetting *CThostFtdcInputOffsetSettingField, RspInfo *CThostFtdcRspInfoField, RequestID int, IsLast bool)

    // 对冲设置通知
    OnRtnOffsetSetting(OffsetSetting *CThostFtdcOffsetSettingField)

    // 对冲设置错误回报
    OnErrRtnOffsetSetting(InputOffsetSetting *CThostFtdcInputOffsetSettingField, RspInfo *CThostFtdcRspInfoField)

    // 对冲设置撤销错误回报
    OnErrRtnCancelOffsetSetting(CancelOffsetSetting *CThostFtdcCancelOffsetSettingField, RspInfo *CThostFtdcRspInfoField)

    // 投资者对冲设置查询响应
    OnRspQryOffsetSetting(OffsetSetting *CThostFtdcOffsetSettingField, RspInfo *CThostFtdcRspInfoField, RequestID int, IsLast bool)

    // 申请短信验证码响应
    OnRspGenSMSCode(RspGenSMSCode *CThostFtdcRspGenSMSCodeField, RspInfo *CThostFtdcRspInfoField, RequestID int, IsLast bool)

    // 套利确认回复
    OnRspSpdApply(InputSpdApply *CThostFtdcInputSpdApplyField, RspInfo *CThostFtdcRspInfoField, RequestID int, IsLast bool)

    // 套利确认撤销回复
    OnRspSpdApplyAction(InputSpdApplyAction *CThostFtdcInputSpdApplyActionField, RspInfo *CThostFtdcRspInfoField, RequestID int, IsLast bool)

    // 套利确认查询回复
    OnRspQrySpdApply(SpdApply *CThostFtdcSpdApplyField, RspInfo *CThostFtdcRspInfoField, RequestID int, IsLast bool)

    // 套利确认通知
    OnRtnSpdApply(SpdApply *CThostFtdcSpdApplyField)

    // 套利申请录入错误回报
    OnErrRtnSpdApply(InputSpdApply *CThostFtdcInputSpdApplyField, RspInfo *CThostFtdcRspInfoField)

    // 套利确认撤销通知
    OnErrRtnSpdApplyAction(SpdApplyAction *CThostFtdcSpdApplyActionField, RspInfo *CThostFtdcRspInfoField)

    // 套保确认回复
    OnRspHedgeCfm(InputHedgeCfm *CThostFtdcInputHedgeCfmField, RspInfo *CThostFtdcRspInfoField, RequestID int, IsLast bool)

    // 套保确认撤销回复
    OnRspHedgeCfmAction(InputHedgeCfmAction *CThostFtdcInputHedgeCfmActionField, RspInfo *CThostFtdcRspInfoField, RequestID int, IsLast bool)

    // 套保确认查询回复
    OnRspQryHedgeCfm(HedgeCfm *CThostFtdcHedgeCfmField, RspInfo *CThostFtdcRspInfoField, RequestID int, IsLast bool)

    // 套保确认通知
    OnRtnHedgeCfm(HedgeCfm *CThostFtdcHedgeCfmField)

    // 套保额度录入错误回报
    OnErrRtnHedgeCfm(InputHedgeCfm *CThostFtdcInputHedgeCfmField, RspInfo *CThostFtdcRspInfoField)

    // 套保确认撤销通知
    OnErrRtnHedgeCfmAction(HedgeCfmAction *CThostFtdcHedgeCfmActionField, RspInfo *CThostFtdcRspInfoField)

}
