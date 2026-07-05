package mini

type TraderSpi interface {
    // 当客户端与交易后台建立起通信连接时（还未登录前），该方法被调用。
    OnFrontConnected()

    // 当客户端与交易后台通信连接断开时，该方法被调用。当发生这个情况后，API会自动重新连接，客户端可不做处理。
    //   @param nReason 错误原因
    //     -3	关闭连接
    //     -4	网络读失败
    //     -5	网络写失败
    //     -6	读订阅流水请求出错
    //     -7	序列号错误
    //     -8	读心跳出错
    //     -9	错误的网络包大小
    OnFrontDisconnected(Reason int)

    // 心跳超时警告。当长时间未收到报文时，该方法被调用。
    //   @param nTimeLapse 距离上次接收报文的时间
    OnHeartBeatWarning(TimeLapse int)

    // 订阅流控警告应答
    OnRspSubscribeFlowCtrlWarning(RspSubscribeTraderField *CThostFtdcSpecificTraderField, RspInfo *CThostFtdcRspInfoField, RequestID int, IsLast bool)

    // 取消订阅流控警告应答
    OnRspUnSubscribeFlowCtrlWarning(RspSubscribeTraderField *CThostFtdcSpecificTraderField, RspInfo *CThostFtdcRspInfoField, RequestID int, IsLast bool)

    // 客户端认证响应
    OnRspAuthenticate(RspAuthenticateField *CThostFtdcRspAuthenticateField, RspInfo *CThostFtdcRspInfoField, RequestID int, IsLast bool)

    // 登录请求响应
    OnRspUserLogin(RspUserLogin *CThostFtdcRspUserLoginField, RspInfo *CThostFtdcRspInfoField, RequestID int, IsLast bool)

    // 登出请求响应
    OnRspUserLogout(UserLogout *CThostFtdcUserLogoutField, RspInfo *CThostFtdcRspInfoField, RequestID int, IsLast bool)

    // 报单录入请求响应
    OnRspOrderInsert(InputOrder *CThostFtdcInputOrderField, RspInfo *CThostFtdcRspInfoField, RequestID int, IsLast bool)

    // 报单操作请求响应
    OnRspOrderAction(InputOrderAction *CThostFtdcInputOrderActionField, RspInfo *CThostFtdcRspInfoField, RequestID int, IsLast bool)

    // 做市商批量报单操作请求响应
    OnRspMKBatchOrderAction(MKInputOrderAction *CThostFtdcMKInputOrderActionField, RspInfo *CThostFtdcRspInfoField, RequestID int, IsLast bool)

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

    // 请求查询交易所响应
    OnRspQryExchange(Exchange *CThostFtdcExchangeField, RspInfo *CThostFtdcRspInfoField, RequestID int, IsLast bool)

    // 请求查询产品响应
    OnRspQryProduct(Product *CThostFtdcProductField, RspInfo *CThostFtdcRspInfoField, RequestID int, IsLast bool)

    // 请求查询合约响应
    OnRspQryInstrument(Instrument *CThostFtdcInstrumentField, RspInfo *CThostFtdcRspInfoField, RequestID int, IsLast bool)

    // 请求查询申请组合合约响应
    OnRspQryCombInstrument(CombInstrument *CThostFtdcCombInstrumentField, RspInfo *CThostFtdcRspInfoField, RequestID int, IsLast bool)

    // 请求查询投资者RCAMS组合保证金响应
    OnRspQryRCAMSInvestorProdMargin(RCAMSInvestorProdMargin *CThostFtdcRCAMSInvestorProdMarginField, RspInfo *CThostFtdcRspInfoField, RequestID int, IsLast bool)

    // 请求查询RCAMS策略组合持仓响应
    OnRspQryRCAMSInvestorCombPosition(RCAMSInvestorCombPosition *CThostFtdcRCAMSInvestorCombPositionField, RspInfo *CThostFtdcRspInfoField, RequestID int, IsLast bool)

    // 请求查询申请组合响应
    OnRspQryCombAction(CombAction *CThostFtdcCombActionField, RspInfo *CThostFtdcRspInfoField, RequestID int, IsLast bool)

    // 请求查询组合单腿汇总表响应
    OnRspQryInvestorPositionForComb(ForComb *CThostFtdcInvestorPositionForCombField, RspInfo *CThostFtdcRspInfoField, RequestID int, IsLast bool)

    // 请求查询行情响应
    OnRspQryDepthMarketData(DepthMarketData *CThostFtdcDepthMarketDataField, RspInfo *CThostFtdcRspInfoField, RequestID int, IsLast bool)

    // 请求查询合约状态响应
    OnRspQryInstrumentStatus(InstrumentStatus *CThostFtdcInstrumentStatusField, RspInfo *CThostFtdcRspInfoField, RequestID int, IsLast bool)

    // 请求查询投资者持仓明细响应
    OnRspQryInvestorPositionDetail(InvestorPositionDetail *CThostFtdcInvestorPositionDetailField, RspInfo *CThostFtdcRspInfoField, RequestID int, IsLast bool)

    // 请求查询交易所保证金率响应
    OnRspQryExchangeMarginRate(ExchangeMarginRate *CThostFtdcExchangeMarginRateField, RspInfo *CThostFtdcRspInfoField, RequestID int, IsLast bool)

    // 请求查询交易所调整保证金率响应
    OnRspQryExchangeMarginRateAdjust(ExchangeMarginRateAdjust *CThostFtdcExchangeMarginRateAdjustField, RspInfo *CThostFtdcRspInfoField, RequestID int, IsLast bool)

    // 请求查询期权交易成本响应
    OnRspQryOptionInstrTradeCost(OptionInstrTradeCost *CThostFtdcOptionInstrTradeCostField, RspInfo *CThostFtdcRspInfoField, RequestID int, IsLast bool)

    // 请求查询期权合约手续费响应
    OnRspQryOptionInstrCommRate(OptionInstrCommRate *CThostFtdcOptionInstrCommRateField, RspInfo *CThostFtdcRspInfoField, RequestID int, IsLast bool)

    // 请求查询执行宣告响应
    OnRspQryExecOrder(ExecOrder *CThostFtdcExecOrderField, RspInfo *CThostFtdcRspInfoField, RequestID int, IsLast bool)

    // 请求查询询价响应
    OnRspQryForQuote(ForQuote *CThostFtdcForQuoteField, RspInfo *CThostFtdcRspInfoField, RequestID int, IsLast bool)

    // 请求查询询价价差响应
    OnRspQryForQuoteParam(ForQuoteParam *CThostFtdcForQuoteParamField, RspInfo *CThostFtdcRspInfoField, RequestID int, IsLast bool)

    // 请求查询投资者SPBM品种明细响应
    OnRspQryInvestorProdSPBMDetail(InvestorProdSPBMDetail *CThostFtdcInvestorProdSPBMDetailField, RspInfo *CThostFtdcRspInfoField, RequestID int, IsLast bool)

    // 请求查询投资者SPMM商品群明细响应
    OnRspQrySPMMInvestorCommodityGroupMargin(SPMMInvestorCommodityGroupMargin *CThostFtdcSPMMInvestorCommodityGroupMarginField, RspInfo *CThostFtdcRspInfoField, RequestID int, IsLast bool)

    // 请求查询投资者RULE保证金响应
    OnRspQryRULEInvestorProdMargin(RULEInvestorProdMargin *CThostFtdcRULEInvestorProdMarginField, RspInfo *CThostFtdcRspInfoField, RequestID int, IsLast bool)

    // 请求查询交易员报盘机响应
    OnRspQryTraderOffer(TraderOffer *CThostFtdcTraderOfferField, RspInfo *CThostFtdcRspInfoField, RequestID int, IsLast bool)

    // 请求查询报价响应
    OnRspQryQuote(Quote *CThostFtdcQuoteField, RspInfo *CThostFtdcRspInfoField, RequestID int, IsLast bool)

    // 请求查询期权自对冲响应
    OnRspQryOptionSelfClose(OptionSelfClose *CThostFtdcOptionSelfCloseField, RspInfo *CThostFtdcRspInfoField, RequestID int, IsLast bool)

    // 请求查询交易开关设置响应
    OnRspQryControlParam(ControlParam *CThostFtdcControlParamField, RspInfo *CThostFtdcRspInfoField, RequestID int, IsLast bool)

    // 请求查询对冲设置响应
    OnRspQryOffsetSetting(OffsetSetting *CThostFtdcOffsetSettingField, RspInfo *CThostFtdcRspInfoField, RequestID int, IsLast bool)

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

    // 请求查询申报费响应
    OnRspQryInstrumentOrderCommRate(InstrumentOrderCommRate *CThostFtdcInstrumentOrderCommRateField, RspInfo *CThostFtdcRspInfoField, RequestID int, IsLast bool)

    // 交易所席位流控警告
    OnRtnFlowCtrlWarning(FlowCtrlWarning *CThostFtdcFlowCtrlWarningField)

    // 订阅资金变动应答
    OnRspSubscribeFundChange(RspInfo *CThostFtdcRspInfoField, RequestID int, IsLast bool)

    // 取消订阅资金变动应答
    OnRspUnSubscribeFundChange(RspInfo *CThostFtdcRspInfoField, RequestID int, IsLast bool)

    // 资金变动推送
    OnRtnFundChange(TradingAccount *CThostFtdcTradingAccountField)

    // 对冲设置请求响应
    OnRspOffsetSetting(InputOffsetSetting *CThostFtdcInputOffsetSettingField, RspInfo *CThostFtdcRspInfoField, RequestID int, IsLast bool)

    // 对冲设置撤销请求响应
    OnRspCancelOffsetSetting(InputOffsetSetting *CThostFtdcInputOffsetSettingField, RspInfo *CThostFtdcRspInfoField, RequestID int, IsLast bool)

    // 对冲设置通知
    OnRtnOffsetSetting(OffsetSetting *CThostFtdcOffsetSettingField)

    // 对冲设置撤销错误回报
    OnErrRtnCancelOffsetSetting(CancelOffsetSetting *CThostFtdcCancelOffsetSettingField, RspInfo *CThostFtdcRspInfoField)

    // 用户口令更新请求响应
    OnRspUserPasswordUpdate(UserPasswordUpdate *CThostFtdcUserPasswordUpdateField, RspInfo *CThostFtdcRspInfoField, RequestID int, IsLast bool)

}
