package mini

type TraderApi interface {
	// 删除接口对象本身
	//   @remark  不再使用本接口对象时,调用该函数删除接口对象
	Release()

	// 初始化
	//   @param bContinuous 为true表示线程不休眠
	//   @remark  初始化运行环境,只有调用后,接口才开始工作
	Init(Continuous bool)

	// 等待接口线程结束运行
	//   @return  线程退出代码
	Join() int

	// 获取当前交易日
	// 获取到的交易日
	//   @remark  只有登录成功后,才能得到正确的交易日
	GetTradingDay() string

	// 注册前置机网络地址
	//   @param pszFrontAddress：前置机网络地址。
	//   @remark  网络地址的格式为：“protocol:ipaddress:port”，如：”tcp:127.0.0.1:17001”。
	//   @remark  “tcp”代表传输协议，“127.0.0.1”代表服务器地址。”17001”代表服务器端口号。
	RegisterFront(FrontAddress string)

	// 注册回调接口
	//   @param pSpi 派生自回调接口类的实例
	RegisterSpi(Spi TraderSpi)

	// 订阅私有流。
	//   @param nResumeType 私有流重传方式
	//     THOST_TERT_RESTART:从本交易日开始重传
	//     THOST_TERT_RESUME:从上次收到的续传
	//     THOST_TERT_QUICK:只传送登录后私有流的内容
	//   @remark  该方法要在Init方法前调用。若不调用则不会收到私有流的数据。
	SubscribePrivateTopic(ResumeType int)

	// 订阅公共流。
	//   @param nResumeType 公共流重传方式
	//     THOST_TERT_RESTART:从本交易日开始重传
	//     THOST_TERT_RESUME:从上次收到的续传
	//     THOST_TERT_QUICK:只传送登录后公共流的内容
	//   @remark  该方法要在Init方法前调用。若不调用则不会收到公共流的数据。
	SubscribePublicTopic(ResumeType int)

	// 订阅交易所流控警告
	//   @remark  该方法必须在登录成功之后调用。
	SubscribeFlowCtrlWarning(TraderID ...string) int

	// 取消订阅交易所流控警告
	//   @remark  该方法必须在登录成功之后调用。
	UnSubscribeFlowCtrlWarning(TraderID ...string) int

	// 客户端认证请求
	ReqAuthenticate(ReqAuthenticateField *CThostFtdcReqAuthenticateField, RequestID int) int

	// 用户登录请求
	ReqUserLogin(ReqUserLoginField *CThostFtdcReqUserLoginField, RequestID int) int

	// 用户加密登录请求
	ReqUserLoginEncrypt(ReqUserLoginField *CThostFtdcReqUserLoginField, RequestID int) int

	// 登出请求
	ReqUserLogout(UserLogout *CThostFtdcUserLogoutField, RequestID int) int

	// 报单录入请求
	ReqOrderInsert(InputOrder *CThostFtdcInputOrderField, RequestID int) int

	// 报单操作请求
	ReqOrderAction(InputOrderAction *CThostFtdcInputOrderActionField, RequestID int) int

	// 做市商批量报单操作请求
	ReqMKBatchOrderAction(MKInputOrderAction *CThostFtdcMKInputOrderActionField, RequestID int) int

	// 执行宣告录入请求
	ReqExecOrderInsert(InputExecOrder *CThostFtdcInputExecOrderField, RequestID int) int

	// 执行宣告操作请求
	ReqExecOrderAction(InputExecOrderAction *CThostFtdcInputExecOrderActionField, RequestID int) int

	// 询价录入请求
	ReqForQuoteInsert(InputForQuote *CThostFtdcInputForQuoteField, RequestID int) int

	// 报价录入请求
	ReqQuoteInsert(InputQuote *CThostFtdcInputQuoteField, RequestID int) int

	// 报价操作请求
	ReqQuoteAction(InputQuoteAction *CThostFtdcInputQuoteActionField, RequestID int) int

	// 批量报单操作请求
	ReqBatchOrderAction(InputBatchOrderAction *CThostFtdcInputBatchOrderActionField, RequestID int) int

	// 期权自对冲录入请求
	ReqOptionSelfCloseInsert(InputOptionSelfClose *CThostFtdcInputOptionSelfCloseField, RequestID int) int

	// 期权自对冲操作请求
	ReqOptionSelfCloseAction(InputOptionSelfCloseAction *CThostFtdcInputOptionSelfCloseActionField, RequestID int) int

	// 申请组合录入请求
	ReqCombActionInsert(InputCombAction *CThostFtdcInputCombActionField, RequestID int) int

	// 订阅资金变动推送
	ReqSubscribeFundChange(RequestID int) int

	// 取消订阅资金变动推送
	ReqUnSubscribeFundChange(RequestID int) int

	// 对冲设置请求
	ReqOffsetSetting(InputOffsetSetting *CThostFtdcInputOffsetSettingField, RequestID int) int

	// 撤销对冲设置请求
	ReqCancelOffsetSetting(InputOffsetSetting *CThostFtdcInputOffsetSettingField, RequestID int) int

	// 请求查询报单
	ReqQryOrder(QryOrder *CThostFtdcQryOrderField, RequestID int) int

	// 请求查询成交
	ReqQryTrade(QryTrade *CThostFtdcQryTradeField, RequestID int) int

	// 请求查询投资者持仓
	ReqQryInvestorPosition(QryInvestorPosition *CThostFtdcQryInvestorPositionField, RequestID int) int

	// 请求查询资金账户
	ReqQryTradingAccount(QryTradingAccount *CThostFtdcQryTradingAccountField, RequestID int) int

	// 请求查询投资者
	ReqQryInvestor(QryInvestor *CThostFtdcQryInvestorField, RequestID int) int

	// 请求查询交易编码
	ReqQryTradingCode(QryTradingCode *CThostFtdcQryTradingCodeField, RequestID int) int

	// 请求查询合约保证金率
	ReqQryInstrumentMarginRate(QryInstrumentMarginRate *CThostFtdcQryInstrumentMarginRateField, RequestID int) int

	// 请求查询合约手续费率
	ReqQryInstrumentCommissionRate(QryInstrumentCommissionRate *CThostFtdcQryInstrumentCommissionRateField, RequestID int) int

	// 请求查询交易所
	ReqQryExchange(QryExchange *CThostFtdcQryExchangeField, RequestID int) int

	// 请求查询产品
	ReqQryProduct(QryProduct *CThostFtdcQryProductField, RequestID int) int

	// 请求查询合约
	ReqQryInstrument(QryInstrument *CThostFtdcQryInstrumentField, RequestID int) int

	// 请求查询申请组合合约
	ReqQryCombInstrument(QryCombInstrument *CThostFtdcQryCombInstrumentField, RequestID int) int

	// 查询投资者RCAMS组合保证金
	ReqQryRCAMSInvestorProdMargin(QryRCAMSInvestorProdMargin *CThostFtdcQryRCAMSInvestorProdMarginField, RequestID int) int

	// 查询RCAMS策略组合持仓
	ReqQryRCAMSInvestorCombPosition(QryRCAMSInvestorCombPosition *CThostFtdcQryRCAMSInvestorCombPositionField, RequestID int) int

	// 请求单腿持仓汇总
	ReqQryInvestorPositionForComb(QryIPForComb *CThostFtdcQryInvestorPositionForCombField, RequestID int) int

	// 请求查询申请组合
	ReqQryCombAction(QryCombAction *CThostFtdcQryCombActionField, RequestID int) int

	// 请求查询行情
	ReqQryDepthMarketData(QryDepthMarketData *CThostFtdcQryDepthMarketDataField, RequestID int) int

	// 请求查询期权自对冲
	ReqQryOptionSelfClose(QryOptionSelfClose *CThostFtdcQryOptionSelfCloseField, RequestID int) int

	// 请求查询合约状态
	ReqQryInstrumentStatus(QryInstrumentStatus *CThostFtdcQryInstrumentStatusField, RequestID int) int

	// 请求查询投资者持仓明细
	// 持仓明细不再维护，请勿以查询结果为准
	ReqQryInvestorPositionDetail(QryInvestorPositionDetail *CThostFtdcQryInvestorPositionDetailField, RequestID int) int

	// 请求查询交易所保证金率
	ReqQryExchangeMarginRate(QryExchangeMarginRate *CThostFtdcQryExchangeMarginRateField, RequestID int) int

	// 请求查询交易所调整保证金率
	ReqQryExchangeMarginRateAdjust(QryExchangeMarginRateAdjust *CThostFtdcQryExchangeMarginRateAdjustField, RequestID int) int

	// 请求查询期权交易成本
	ReqQryOptionInstrTradeCost(QryOptionInstrTradeCost *CThostFtdcQryOptionInstrTradeCostField, RequestID int) int

	// 请求查询期权合约手续费
	ReqQryOptionInstrCommRate(QryOptionInstrCommRate *CThostFtdcQryOptionInstrCommRateField, RequestID int) int

	// 请求查询执行宣告
	ReqQryExecOrder(QryExecOrder *CThostFtdcQryExecOrderField, RequestID int) int

	// 请求查询询价
	ReqQryForQuote(QryForQuote *CThostFtdcQryForQuoteField, RequestID int) int

	// 请求查询报价
	ReqQryQuote(QryQuote *CThostFtdcQryQuoteField, RequestID int) int

	// 请求查询申报费率
	ReqQryInstrumentOrderCommRate(QryInstrumentOrderCommRate *CThostFtdcQryInstrumentOrderCommRateField, RequestID int) int

	// 请求查询询价价差
	ReqQryForQuoteParam(QryForQuoteParam *CThostFtdcQryForQuoteParamField, RequestID int) int

	// 请求查询交易员报盘机
	ReqQryTraderOffer(QryTraderOffer *CThostFtdcQryTraderOfferField, RequestID int) int

	// 请求查询投资者SPBM品种明细
	ReqQryInvestorProdSPBMDetail(QryInvestorProdSPBMDetail *CThostFtdcQryInvestorProdSPBMDetailField, RequestID int) int

	// 请求查询投资者SPMM商品群保证金明细
	ReqQrySPMMInvestorCommodityGroupMargin(QrySPMMInvestorCommodityGroupMargin *CThostFtdcQrySPMMInvestorCommodityGroupMarginField, RequestID int) int

	// 请求查询投资者RULE保证金明细
	ReqQryRULEInvestorProdMargin(QryRULEInvestorProdMargin *CThostFtdcQryRULEInvestorProdMarginField, RequestID int) int

	// 请求查询系统功能开关设置
	ReqQryControlParam(QryControlParam *CThostFtdcQryControlParamField, RequestID int) int

	// 请求查询对冲设置
	ReqQryOffsetSetting(QryOffsetSetting *CThostFtdcQryOffsetSettingField, RequestID int) int

	// 用户口令更新请求
	ReqUserPasswordUpdate(UserPasswordUpdate *CThostFtdcUserPasswordUpdateField, RequestID int) int
}
