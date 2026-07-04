package future

type TraderApi interface {
	// 删除接口对象本身
	//   @remark  不再使用本接口对象时,调用该函数删除接口对象
	Release()

	// 初始化
	//   @remark  初始化运行环境,只有调用后,接口才开始工作
	Init()

	// 等待接口线程结束运行
	//   @return  线程退出代码
	Join() int

	// 获取当前交易日
	// 获取到的交易日
	//   @remark  只有登录成功后,才能得到正确的交易日
	GetTradingDay() string

	// 获取已连接的前置的信息
	//   @param pFrontInfo：输入输出参数，用于存储获取到的前置信息，不能为空
	//   @remark  连接成功后，可获取正确的前置地址信息
	//   @remark  登录成功后，可获取正确的前置流控信息
	GetFrontInfo(FrontInfo *CThostFtdcFrontInfoField)

	// 注册前置机网络地址
	//   @param pszFrontAddress：前置机网络地址。
	//   @remark  网络地址的格式为：“protocol:ipaddress:port”，如：”tcp:127.0.0.1:17001”。
	//   @remark  “tcp”代表传输协议，“127.0.0.1”代表服务器地址。”17001”代表服务器端口号。
	RegisterFront(FrontAddress string)

	// 注册名字服务器网络地址
	//   @param pszNsAddress：名字服务器网络地址。
	//   @remark  网络地址的格式为：“protocol:ipaddress:port”，如：”tcp:127.0.0.1:12001”。
	//   @remark  “tcp”代表传输协议，“127.0.0.1”代表服务器地址。”12001”代表服务器端口号。
	//   @remark  RegisterNameServer优先于RegisterFront
	RegisterNameServer(NsAddress string)

	// 注册名字服务器用户信息
	//   @param pFensUserInfo：用户信息。
	RegisterFensUserInfo(FensUserInfo *CThostFtdcFensUserInfoField)

	// 注册回调接口
	//   @param pSpi 派生自回调接口类的实例
	RegisterSpi(Spi TraderSpi)

	// 订阅私有流。
	//   @param nResumeType 私有流重传方式
	//     THOST_TERT_RESTART:从本交易日开始重传
	//     THOST_TERT_RESUME:从上次收到的续传
	//     THOST_TERT_QUICK:只传送登录后私有流的内容
	//     THOST_TERT_RESUME_FROM_SEQ_NO:从指定序号开始重传，序号从1开始
	//   @param nSeqNo 私有流序号，只在THOST_TERT_RESUME_FROM_SEQ_NO模式下有效
	//   @remark  该方法要在Init方法前调用。若不调用则不会收到私有流的数据。
	SubscribePrivateTopic(ResumeType int, SeqNo int)

	// 订阅公共流。
	//   @param nResumeType 公共流重传方式
	//     THOST_TERT_RESTART:从本交易日开始重传
	//     THOST_TERT_RESUME:从上次收到的续传
	//     THOST_TERT_QUICK:只传送登录后公共流的内容
	//     THOST_TERT_NONE:取消订阅公共流
	//   @remark  该方法要在Init方法前调用。若不调用则不会收到公共流的数据。
	SubscribePublicTopic(ResumeType int)

	// 客户端认证请求
	ReqAuthenticate(ReqAuthenticateField *CThostFtdcReqAuthenticateField, RequestID int) int

	// 注册用户终端信息，用于中继服务器多连接模式
	// 需要在终端认证成功后，用户登录前调用该接口
	RegisterUserSystemInfo(UserSystemInfo *CThostFtdcUserSystemInfoField) int

	// 上报用户终端信息，用于中继服务器操作员登录模式
	// 操作员登录后，可以多次调用该接口上报客户信息
	SubmitUserSystemInfo(UserSystemInfo *CThostFtdcUserSystemInfoField) int

	// 注册用户终端信息，用于中继服务器多连接模式.用于微信小程序等应用上报信息.
	RegisterWechatUserSystemInfo(UserSystemInfo *CThostFtdcWechatUserSystemInfoField) int

	// 上报用户终端信息，用于中继服务器操作员登录模式.用于微信小程序等应用上报信息.
	SubmitWechatUserSystemInfo(UserSystemInfo *CThostFtdcWechatUserSystemInfoField) int

	// 用户登录请求
	ReqUserLogin(ReqUserLoginField *CThostFtdcReqUserLoginField, RequestID int) int

	// 登出请求
	ReqUserLogout(UserLogout *CThostFtdcUserLogoutField, RequestID int) int

	// 用户口令更新请求
	ReqUserPasswordUpdate(UserPasswordUpdate *CThostFtdcUserPasswordUpdateField, RequestID int) int

	// 资金账户口令更新请求
	ReqTradingAccountPasswordUpdate(TradingAccountPasswordUpdate *CThostFtdcTradingAccountPasswordUpdateField, RequestID int) int

	// 查询用户当前支持的认证模式
	ReqUserAuthMethod(ReqUserAuthMethod *CThostFtdcReqUserAuthMethodField, RequestID int) int

	// 用户发出获取图形验证码请求
	ReqGenUserCaptcha(ReqGenUserCaptcha *CThostFtdcReqGenUserCaptchaField, RequestID int) int

	// 用户发出获取短信验证码请求
	ReqGenUserText(ReqGenUserText *CThostFtdcReqGenUserTextField, RequestID int) int

	// 用户发出带有图片验证码的登陆请求
	ReqUserLoginWithCaptcha(ReqUserLoginWithCaptcha *CThostFtdcReqUserLoginWithCaptchaField, RequestID int) int

	// 用户发出带有短信验证码的登陆请求
	ReqUserLoginWithText(ReqUserLoginWithText *CThostFtdcReqUserLoginWithTextField, RequestID int) int

	// 用户发出带有动态口令的登陆请求
	ReqUserLoginWithOTP(ReqUserLoginWithOTP *CThostFtdcReqUserLoginWithOTPField, RequestID int) int

	// 报单录入请求
	ReqOrderInsert(InputOrder *CThostFtdcInputOrderField, RequestID int) int

	// 预埋单录入请求
	ReqParkedOrderInsert(ParkedOrder *CThostFtdcParkedOrderField, RequestID int) int

	// 预埋撤单录入请求
	ReqParkedOrderAction(ParkedOrderAction *CThostFtdcParkedOrderActionField, RequestID int) int

	// 报单操作请求
	ReqOrderAction(InputOrderAction *CThostFtdcInputOrderActionField, RequestID int) int

	// 查询最大报单数量请求
	ReqQryMaxOrderVolume(QryMaxOrderVolume *CThostFtdcQryMaxOrderVolumeField, RequestID int) int

	// 投资者结算结果确认
	ReqSettlementInfoConfirm(SettlementInfoConfirm *CThostFtdcSettlementInfoConfirmField, RequestID int) int

	// 请求删除预埋单
	ReqRemoveParkedOrder(RemoveParkedOrder *CThostFtdcRemoveParkedOrderField, RequestID int) int

	// 请求删除预埋撤单
	ReqRemoveParkedOrderAction(RemoveParkedOrderAction *CThostFtdcRemoveParkedOrderActionField, RequestID int) int

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

	// 请求查询用户会话
	ReqQryUserSession(QryUserSession *CThostFtdcQryUserSessionField, RequestID int) int

	// 请求查询交易所
	ReqQryExchange(QryExchange *CThostFtdcQryExchangeField, RequestID int) int

	// 请求查询产品
	ReqQryProduct(QryProduct *CThostFtdcQryProductField, RequestID int) int

	// 请求查询合约
	ReqQryInstrument(QryInstrument *CThostFtdcQryInstrumentField, RequestID int) int

	// 请求查询行情
	ReqQryDepthMarketData(QryDepthMarketData *CThostFtdcQryDepthMarketDataField, RequestID int) int

	// 请求查询交易员报盘机
	ReqQryTraderOffer(QryTraderOffer *CThostFtdcQryTraderOfferField, RequestID int) int

	// 请求查询投资者结算结果
	ReqQrySettlementInfo(QrySettlementInfo *CThostFtdcQrySettlementInfoField, RequestID int) int

	// 请求查询转帐银行
	ReqQryTransferBank(QryTransferBank *CThostFtdcQryTransferBankField, RequestID int) int

	// 请求查询投资者持仓明细
	ReqQryInvestorPositionDetail(QryInvestorPositionDetail *CThostFtdcQryInvestorPositionDetailField, RequestID int) int

	// 请求查询客户通知
	ReqQryNotice(QryNotice *CThostFtdcQryNoticeField, RequestID int) int

	// 请求查询结算信息确认
	ReqQrySettlementInfoConfirm(QrySettlementInfoConfirm *CThostFtdcQrySettlementInfoConfirmField, RequestID int) int

	// 请求查询投资者持仓明细
	ReqQryInvestorPositionCombineDetail(QryInvestorPositionCombineDetail *CThostFtdcQryInvestorPositionCombineDetailField, RequestID int) int

	// 请求查询保证金监管系统经纪公司资金账户密钥
	ReqQryCFMMCTradingAccountKey(QryCFMMCTradingAccountKey *CThostFtdcQryCFMMCTradingAccountKeyField, RequestID int) int

	// 请求查询仓单折抵信息
	ReqQryEWarrantOffset(QryEWarrantOffset *CThostFtdcQryEWarrantOffsetField, RequestID int) int

	// 请求查询投资者品种跨品种保证金
	ReqQryInvestorProductGroupMargin(QryInvestorProductGroupMargin *CThostFtdcQryInvestorProductGroupMarginField, RequestID int) int

	// 请求查询交易所保证金率
	ReqQryExchangeMarginRate(QryExchangeMarginRate *CThostFtdcQryExchangeMarginRateField, RequestID int) int

	// 请求查询交易所调整保证金率
	ReqQryExchangeMarginRateAdjust(QryExchangeMarginRateAdjust *CThostFtdcQryExchangeMarginRateAdjustField, RequestID int) int

	// 请求查询汇率
	ReqQryExchangeRate(QryExchangeRate *CThostFtdcQryExchangeRateField, RequestID int) int

	// 请求查询二级代理操作员银期权限
	ReqQrySecAgentACIDMap(QrySecAgentACIDMap *CThostFtdcQrySecAgentACIDMapField, RequestID int) int

	// 请求查询产品报价汇率
	ReqQryProductExchRate(QryProductExchRate *CThostFtdcQryProductExchRateField, RequestID int) int

	// 请求查询产品组
	ReqQryProductGroup(QryProductGroup *CThostFtdcQryProductGroupField, RequestID int) int

	// 请求查询做市商合约手续费率
	ReqQryMMInstrumentCommissionRate(QryMMInstrumentCommissionRate *CThostFtdcQryMMInstrumentCommissionRateField, RequestID int) int

	// 请求查询做市商期权合约手续费
	ReqQryMMOptionInstrCommRate(QryMMOptionInstrCommRate *CThostFtdcQryMMOptionInstrCommRateField, RequestID int) int

	// 请求查询报单手续费
	ReqQryInstrumentOrderCommRate(QryInstrumentOrderCommRate *CThostFtdcQryInstrumentOrderCommRateField, RequestID int) int

	// 请求查询资金账户
	ReqQrySecAgentTradingAccount(QryTradingAccount *CThostFtdcQryTradingAccountField, RequestID int) int

	// 请求查询二级代理商资金校验模式
	ReqQrySecAgentCheckMode(QrySecAgentCheckMode *CThostFtdcQrySecAgentCheckModeField, RequestID int) int

	// 请求查询二级代理商信息
	ReqQrySecAgentTradeInfo(QrySecAgentTradeInfo *CThostFtdcQrySecAgentTradeInfoField, RequestID int) int

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

	// 请求查询期权自对冲
	ReqQryOptionSelfClose(QryOptionSelfClose *CThostFtdcQryOptionSelfCloseField, RequestID int) int

	// 请求查询投资单元
	ReqQryInvestUnit(QryInvestUnit *CThostFtdcQryInvestUnitField, RequestID int) int

	// 请求查询组合合约安全系数
	ReqQryCombInstrumentGuard(QryCombInstrumentGuard *CThostFtdcQryCombInstrumentGuardField, RequestID int) int

	// 请求查询申请组合
	ReqQryCombAction(QryCombAction *CThostFtdcQryCombActionField, RequestID int) int

	// 请求查询转帐流水
	ReqQryTransferSerial(QryTransferSerial *CThostFtdcQryTransferSerialField, RequestID int) int

	// 请求查询银期签约关系
	ReqQryAccountregister(QryAccountregister *CThostFtdcQryAccountregisterField, RequestID int) int

	// 请求查询签约银行
	ReqQryContractBank(QryContractBank *CThostFtdcQryContractBankField, RequestID int) int

	// 请求查询预埋单
	ReqQryParkedOrder(QryParkedOrder *CThostFtdcQryParkedOrderField, RequestID int) int

	// 请求查询预埋撤单
	ReqQryParkedOrderAction(QryParkedOrderAction *CThostFtdcQryParkedOrderActionField, RequestID int) int

	// 请求查询交易通知
	ReqQryTradingNotice(QryTradingNotice *CThostFtdcQryTradingNoticeField, RequestID int) int

	// 请求查询经纪公司交易参数
	ReqQryBrokerTradingParams(QryBrokerTradingParams *CThostFtdcQryBrokerTradingParamsField, RequestID int) int

	// 请求查询经纪公司交易算法
	ReqQryBrokerTradingAlgos(QryBrokerTradingAlgos *CThostFtdcQryBrokerTradingAlgosField, RequestID int) int

	// 请求查询监控中心用户令牌
	ReqQueryCFMMCTradingAccountToken(QueryCFMMCTradingAccountToken *CThostFtdcQueryCFMMCTradingAccountTokenField, RequestID int) int

	// 期货发起银行资金转期货请求
	ReqFromBankToFutureByFuture(ReqTransfer *CThostFtdcReqTransferField, RequestID int) int

	// 期货发起期货资金转银行请求
	ReqFromFutureToBankByFuture(ReqTransfer *CThostFtdcReqTransferField, RequestID int) int

	// 期货发起查询银行余额请求
	ReqQueryBankAccountMoneyByFuture(ReqQueryAccount *CThostFtdcReqQueryAccountField, RequestID int) int

	// 请求查询分类合约
	ReqQryClassifiedInstrument(QryClassifiedInstrument *CThostFtdcQryClassifiedInstrumentField, RequestID int) int

	// 请求组合优惠比例
	ReqQryCombPromotionParam(QryCombPromotionParam *CThostFtdcQryCombPromotionParamField, RequestID int) int

	// 投资者风险结算持仓查询
	ReqQryRiskSettleInvstPosition(QryRiskSettleInvstPosition *CThostFtdcQryRiskSettleInvstPositionField, RequestID int) int

	// 风险结算产品查询
	ReqQryRiskSettleProductStatus(QryRiskSettleProductStatus *CThostFtdcQryRiskSettleProductStatusField, RequestID int) int

	// SPBM期货合约参数查询
	ReqQrySPBMFutureParameter(QrySPBMFutureParameter *CThostFtdcQrySPBMFutureParameterField, RequestID int) int

	// SPBM期权合约参数查询
	ReqQrySPBMOptionParameter(QrySPBMOptionParameter *CThostFtdcQrySPBMOptionParameterField, RequestID int) int

	// SPBM品种内对锁仓折扣参数查询
	ReqQrySPBMIntraParameter(QrySPBMIntraParameter *CThostFtdcQrySPBMIntraParameterField, RequestID int) int

	// SPBM跨品种抵扣参数查询
	ReqQrySPBMInterParameter(QrySPBMInterParameter *CThostFtdcQrySPBMInterParameterField, RequestID int) int

	// SPBM组合保证金套餐查询
	ReqQrySPBMPortfDefinition(QrySPBMPortfDefinition *CThostFtdcQrySPBMPortfDefinitionField, RequestID int) int

	// 投资者SPBM套餐选择查询
	ReqQrySPBMInvestorPortfDef(QrySPBMInvestorPortfDef *CThostFtdcQrySPBMInvestorPortfDefField, RequestID int) int

	// 投资者新型组合保证金系数查询
	ReqQryInvestorPortfMarginRatio(QryInvestorPortfMarginRatio *CThostFtdcQryInvestorPortfMarginRatioField, RequestID int) int

	// 投资者产品SPBM明细查询
	ReqQryInvestorProdSPBMDetail(QryInvestorProdSPBMDetail *CThostFtdcQryInvestorProdSPBMDetailField, RequestID int) int

	// 投资者商品组SPMM记录查询
	ReqQryInvestorCommoditySPMMMargin(QryInvestorCommoditySPMMMargin *CThostFtdcQryInvestorCommoditySPMMMarginField, RequestID int) int

	// 投资者商品群SPMM记录查询
	ReqQryInvestorCommodityGroupSPMMMargin(QryInvestorCommodityGroupSPMMMargin *CThostFtdcQryInvestorCommodityGroupSPMMMarginField, RequestID int) int

	// SPMM合约参数查询
	ReqQrySPMMInstParam(QrySPMMInstParam *CThostFtdcQrySPMMInstParamField, RequestID int) int

	// SPMM产品参数查询
	ReqQrySPMMProductParam(QrySPMMProductParam *CThostFtdcQrySPMMProductParamField, RequestID int) int

	// SPBM附加跨品种抵扣参数查询
	ReqQrySPBMAddOnInterParameter(QrySPBMAddOnInterParameter *CThostFtdcQrySPBMAddOnInterParameterField, RequestID int) int

	// RCAMS产品组合信息查询
	ReqQryRCAMSCombProductInfo(QryRCAMSCombProductInfo *CThostFtdcQryRCAMSCombProductInfoField, RequestID int) int

	// RCAMS同合约风险对冲参数查询
	ReqQryRCAMSInstrParameter(QryRCAMSInstrParameter *CThostFtdcQryRCAMSInstrParameterField, RequestID int) int

	// RCAMS品种内风险对冲参数查询
	ReqQryRCAMSIntraParameter(QryRCAMSIntraParameter *CThostFtdcQryRCAMSIntraParameterField, RequestID int) int

	// RCAMS跨品种风险折抵参数查询
	ReqQryRCAMSInterParameter(QryRCAMSInterParameter *CThostFtdcQryRCAMSInterParameterField, RequestID int) int

	// RCAMS空头期权风险调整参数查询
	ReqQryRCAMSShortOptAdjustParam(QryRCAMSShortOptAdjustParam *CThostFtdcQryRCAMSShortOptAdjustParamField, RequestID int) int

	// RCAMS策略组合持仓查询
	ReqQryRCAMSInvestorCombPosition(QryRCAMSInvestorCombPosition *CThostFtdcQryRCAMSInvestorCombPositionField, RequestID int) int

	// 投资者品种RCAMS保证金查询
	ReqQryInvestorProdRCAMSMargin(QryInvestorProdRCAMSMargin *CThostFtdcQryInvestorProdRCAMSMarginField, RequestID int) int

	// RULE合约保证金参数查询
	ReqQryRULEInstrParameter(QryRULEInstrParameter *CThostFtdcQryRULEInstrParameterField, RequestID int) int

	// RULE品种内对锁仓折扣参数查询
	ReqQryRULEIntraParameter(QryRULEIntraParameter *CThostFtdcQryRULEIntraParameterField, RequestID int) int

	// RULE跨品种抵扣参数查询
	ReqQryRULEInterParameter(QryRULEInterParameter *CThostFtdcQryRULEInterParameterField, RequestID int) int

	// 投资者产品RULE保证金查询
	ReqQryInvestorProdRULEMargin(QryInvestorProdRULEMargin *CThostFtdcQryInvestorProdRULEMarginField, RequestID int) int

	// 投资者新型组合保证金开关查询
	ReqQryInvestorPortfSetting(QryInvestorPortfSetting *CThostFtdcQryInvestorPortfSettingField, RequestID int) int

	// 投资者申报费阶梯收取记录查询
	ReqQryInvestorInfoCommRec(QryInvestorInfoCommRec *CThostFtdcQryInvestorInfoCommRecField, RequestID int) int

	// 组合腿信息查询
	ReqQryCombLeg(QryCombLeg *CThostFtdcQryCombLegField, RequestID int) int

	// 对冲设置请求
	ReqOffsetSetting(InputOffsetSetting *CThostFtdcInputOffsetSettingField, RequestID int) int

	// 对冲设置撤销请求
	ReqCancelOffsetSetting(InputOffsetSetting *CThostFtdcInputOffsetSettingField, RequestID int) int

	// 投资者对冲设置查询
	ReqQryOffsetSetting(QryOffsetSetting *CThostFtdcQryOffsetSettingField, RequestID int) int

	// 申请短信验证码请求
	ReqGenSMSCode(ReqGenSMSCode *CThostFtdcReqGenSMSCodeField, RequestID int) int

	// 套利确认请求
	ReqSpdApply(InputSpdApply *CThostFtdcInputSpdApplyField, RequestID int) int

	// 套利确认撤销请求
	ReqSpdApplyAction(InputSpdApplyAction *CThostFtdcInputSpdApplyActionField, RequestID int) int

	// 套利确认查询请求
	ReqQrySpdApply(QrySpdApply *CThostFtdcQrySpdApplyField, RequestID int) int

	// 套保确认请求
	ReqHedgeCfm(InputHedgeCfm *CThostFtdcInputHedgeCfmField, RequestID int) int

	// 套保确认撤销请求
	ReqHedgeCfmAction(InputHedgeCfmAction *CThostFtdcInputHedgeCfmActionField, RequestID int) int

	// 套保确认查询请求
	ReqQryHedgeCfm(QryHedgeCfm *CThostFtdcQryHedgeCfmField, RequestID int) int
}
