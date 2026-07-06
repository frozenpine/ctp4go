package mini

import "github.com/frozenpine/ctp4go/thost/mini/types"

// 信息分发
type CThostFtdcDisseminationField struct {
	// 序列系列号
	SequenceSeries types.TThostFtdcSequenceSeriesType
	// 序列号
	SequenceNo types.TThostFtdcSequenceNoType
}

// 用户登录请求
type CThostFtdcReqUserLoginField struct {
	// 交易日
	TradingDay types.TThostFtdcDateType
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 用户代码
	UserID types.TThostFtdcUserIDType
	// 密码
	Password types.TThostFtdcPasswordType
	// 用户端产品信息
	UserProductInfo types.TThostFtdcProductInfoType
	// 接口端产品信息
	InterfaceProductInfo types.TThostFtdcProductInfoType
	// 协议信息
	ProtocolInfo types.TThostFtdcProtocolInfoType
	// Mac地址
	MacAddress types.TThostFtdcMacAddressType
	// 动态密码
	OneTimePassword types.TThostFtdcPasswordType
	// 终端IP地址
	ClientIPAddress types.TThostFtdcIPAddressType
	// 登录备注
	LoginRemark types.TThostFtdcLoginRemarkType
}

// 用户登录应答
type CThostFtdcRspUserLoginField struct {
	// 交易日
	TradingDay types.TThostFtdcDateType
	// 登录成功时间
	LoginTime types.TThostFtdcTimeType
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 用户代码
	UserID types.TThostFtdcUserIDType
	// 交易系统名称
	SystemName types.TThostFtdcSystemNameType
	// 前置编号
	FrontID types.TThostFtdcFrontIDType
	// 会话编号
	SessionID types.TThostFtdcSessionIDType
	// 最大报单引用
	MaxOrderRef types.TThostFtdcOrderRefType
	// 上期所时间
	SHFETime types.TThostFtdcTimeType
	// 大商所时间
	DCETime types.TThostFtdcTimeType
	// 郑商所时间
	CZCETime types.TThostFtdcTimeType
	// 中金所时间
	FFEXTime types.TThostFtdcTimeType
	// 能源中心时间
	INETime types.TThostFtdcTimeType
	// 广期所时间
	GFEXTime types.TThostFtdcTimeType
}

// 用户登出请求
type CThostFtdcUserLogoutField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 用户代码
	UserID types.TThostFtdcUserIDType
}

// 强制交易员退出
type CThostFtdcForceUserLogoutField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 用户代码
	UserID types.TThostFtdcUserIDType
}

// 客户端认证请求
type CThostFtdcReqAuthenticateField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 用户代码
	UserID types.TThostFtdcUserIDType
	// 用户端产品信息
	UserProductInfo types.TThostFtdcProductInfoType
	// 认证码
	AuthCode types.TThostFtdcAuthCodeType
	// App代码
	AppID types.TThostFtdcClientAppIDType
}

// 客户端认证响应
type CThostFtdcRspAuthenticateField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 用户代码
	UserID types.TThostFtdcUserIDType
	// 用户端产品信息
	UserProductInfo types.TThostFtdcProductInfoType
	// App代码
	AppID types.TThostFtdcClientAppIDType
}

// 客户端认证信息
type CThostFtdcAuthenticationInfoField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 用户代码
	UserID types.TThostFtdcUserIDType
	// 用户端产品信息
	UserProductInfo types.TThostFtdcProductInfoType
	// 时间戳
	TimeStamp types.TThostFtdcAuthInfoType
	// 认证信息
	AuthInfo types.TThostFtdcAuthInfoType
	// 是否为认证结果
	IsResult types.TThostFtdcBoolType
	// App代码
	AppID types.TThostFtdcClientAppIDType
}

// 银期转帐报文头
type CThostFtdcTransferHeaderField struct {
	// 版本号，常量，1.0
	Version types.TThostFtdcVersionType
	// 交易代码，必填
	TradeCode types.TThostFtdcTradeCodeType
	// 交易日期，必填，格式：yyyymmdd
	TradeDate types.TThostFtdcTradeDateType
	// 交易时间，必填，格式：hhmmss
	TradeTime types.TThostFtdcTradeTimeType
	// 发起方流水号，NA
	TradeSerial types.TThostFtdcTradeSerialType
	// 期货公司代码，必填
	FutureID types.TThostFtdcFutureIDType
	// 银行代码，根据查询银行得到，必填
	BankID types.TThostFtdcBankIDType
	// 银行分中心代码，根据查询银行得到，必填
	BankBrchID types.TThostFtdcBankBrchIDType
	// 操作员，NA
	OperNo types.TThostFtdcOperNoType
	// 交易设备类型，NA
	DeviceID types.TThostFtdcDeviceIDType
	// 记录数，NA
	RecordNum types.TThostFtdcRecordNumType
	// 会话编号，NA
	SessionID types.TThostFtdcSessionIDType
	// 请求编号，NA
	RequestID types.TThostFtdcRequestIDType
}

// 银行资金转期货请求，TradeCode=202001
type CThostFtdcTransferBankToFutureReqField struct {
	// 期货资金账户
	FutureAccount types.TThostFtdcAccountIDType
	// 密码标志
	FuturePwdFlag types.TThostFtdcFuturePwdFlagType
	// 密码
	FutureAccPwd types.TThostFtdcFutureAccPwdType
	// 转账金额
	TradeAmt types.TThostFtdcMoneyType
	// 客户手续费
	CustFee types.TThostFtdcMoneyType
	// 币种：RMB-人民币 USD-美圆 HKD-港元
	CurrencyCode types.TThostFtdcCurrencyCodeType
}

// 银行资金转期货请求响应
type CThostFtdcTransferBankToFutureRspField struct {
	// 响应代码
	RetCode types.TThostFtdcRetCodeType
	// 响应信息
	RetInfo types.TThostFtdcRetInfoType
	// 资金账户
	FutureAccount types.TThostFtdcAccountIDType
	// 转帐金额
	TradeAmt types.TThostFtdcMoneyType
	// 应收客户手续费
	CustFee types.TThostFtdcMoneyType
	// 币种
	CurrencyCode types.TThostFtdcCurrencyCodeType
}

// 期货资金转银行请求，TradeCode=202002
type CThostFtdcTransferFutureToBankReqField struct {
	// 期货资金账户
	FutureAccount types.TThostFtdcAccountIDType
	// 密码标志
	FuturePwdFlag types.TThostFtdcFuturePwdFlagType
	// 密码
	FutureAccPwd types.TThostFtdcFutureAccPwdType
	// 转账金额
	TradeAmt types.TThostFtdcMoneyType
	// 客户手续费
	CustFee types.TThostFtdcMoneyType
	// 币种：RMB-人民币 USD-美圆 HKD-港元
	CurrencyCode types.TThostFtdcCurrencyCodeType
}

// 期货资金转银行请求响应
type CThostFtdcTransferFutureToBankRspField struct {
	// 响应代码
	RetCode types.TThostFtdcRetCodeType
	// 响应信息
	RetInfo types.TThostFtdcRetInfoType
	// 资金账户
	FutureAccount types.TThostFtdcAccountIDType
	// 转帐金额
	TradeAmt types.TThostFtdcMoneyType
	// 应收客户手续费
	CustFee types.TThostFtdcMoneyType
	// 币种
	CurrencyCode types.TThostFtdcCurrencyCodeType
}

// 查询银行资金请求，TradeCode=204002
type CThostFtdcTransferQryBankReqField struct {
	// 期货资金账户
	FutureAccount types.TThostFtdcAccountIDType
	// 密码标志
	FuturePwdFlag types.TThostFtdcFuturePwdFlagType
	// 密码
	FutureAccPwd types.TThostFtdcFutureAccPwdType
	// 币种：RMB-人民币 USD-美圆 HKD-港元
	CurrencyCode types.TThostFtdcCurrencyCodeType
}

// 查询银行资金请求响应
type CThostFtdcTransferQryBankRspField struct {
	// 响应代码
	RetCode types.TThostFtdcRetCodeType
	// 响应信息
	RetInfo types.TThostFtdcRetInfoType
	// 资金账户
	FutureAccount types.TThostFtdcAccountIDType
	// 银行余额
	TradeAmt types.TThostFtdcMoneyType
	// 银行可用余额
	UseAmt types.TThostFtdcMoneyType
	// 银行可取余额
	FetchAmt types.TThostFtdcMoneyType
	// 币种
	CurrencyCode types.TThostFtdcCurrencyCodeType
}

// 查询银行交易明细请求，TradeCode=204999
type CThostFtdcTransferQryDetailReqField struct {
	// 期货资金账户
	FutureAccount types.TThostFtdcAccountIDType
}

// 查询银行交易明细请求响应
type CThostFtdcTransferQryDetailRspField struct {
	// 交易日期
	TradeDate types.TThostFtdcDateType
	// 交易时间
	TradeTime types.TThostFtdcTradeTimeType
	// 交易代码
	TradeCode types.TThostFtdcTradeCodeType
	// 期货流水号
	FutureSerial types.TThostFtdcTradeSerialNoType
	// 期货公司代码
	FutureID types.TThostFtdcFutureIDType
	// 资金帐号
	FutureAccount types.TThostFtdcFutureAccountType
	// 银行流水号
	BankSerial types.TThostFtdcTradeSerialNoType
	// 银行代码
	BankID types.TThostFtdcBankIDType
	// 银行分中心代码
	BankBrchID types.TThostFtdcBankBrchIDType
	// 银行账号
	BankAccount types.TThostFtdcBankAccountType
	// 证件号码
	CertCode types.TThostFtdcCertCodeType
	// 货币代码
	CurrencyCode types.TThostFtdcCurrencyCodeType
	// 发生金额
	TxAmount types.TThostFtdcMoneyType
	// 有效标志
	Flag types.TThostFtdcTransferValidFlagType
}

// 响应信息
type CThostFtdcRspInfoField struct {
	// 错误代码
	ErrorID types.TThostFtdcErrorIDType
	// 错误信息
	ErrorMsg types.TThostFtdcErrorMsgType
	// 批量交易成功记录数
	RecordCount types.TThostFtdcRecordCountType
}

// 交易所
type CThostFtdcExchangeField struct {
	// 交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	// 交易所名称
	ExchangeName types.TThostFtdcExchangeNameType
	// 交易所属性
	ExchangeProperty types.TThostFtdcExchangePropertyType
}

// 产品
type CThostFtdcProductField struct {
	// 产品代码
	ProductID types.TThostFtdcInstrumentIDType
	// 产品名称
	ProductName types.TThostFtdcProductNameType
	// 交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	// 产品类型
	ProductClass types.TThostFtdcProductClassType
	// 合约数量乘数
	VolumeMultiple types.TThostFtdcVolumeMultipleType
	// 最小变动价位
	PriceTick types.TThostFtdcPriceType
	// 市价单最大下单量
	MaxMarketOrderVolume types.TThostFtdcVolumeType
	// 市价单最小下单量
	MinMarketOrderVolume types.TThostFtdcVolumeType
	// 限价单最大下单量
	MaxLimitOrderVolume types.TThostFtdcVolumeType
	// 限价单最小下单量
	MinLimitOrderVolume types.TThostFtdcVolumeType
	// 持仓类型
	PositionType types.TThostFtdcPositionTypeType
	// 持仓日期类型
	PositionDateType types.TThostFtdcPositionDateTypeType
	// 平仓处理类型
	CloseDealType types.TThostFtdcCloseDealTypeType
	// 交易币种类型
	TradeCurrencyID types.TThostFtdcCurrencyIDType
	// 质押资金可用范围
	MortgageFundUseRange types.TThostFtdcMortgageFundUseRangeType
	// 交易所产品代码
	ExchangeProductID types.TThostFtdcInstrumentIDType
	// 合约基础商品乘数
	UnderlyingMultiple types.TThostFtdcUnderlyingMultipleType
}

// 合约
type CThostFtdcInstrumentField struct {
	// 合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
	// 交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	// 合约名称
	InstrumentName types.TThostFtdcInstrumentNameType
	// 合约在交易所的代码
	ExchangeInstID types.TThostFtdcExchangeInstIDType
	// 产品代码
	ProductID types.TThostFtdcInstrumentIDType
	// 产品类型
	ProductClass types.TThostFtdcProductClassType
	// 交割年份
	DeliveryYear types.TThostFtdcYearType
	// 交割月
	DeliveryMonth types.TThostFtdcMonthType
	// 市价单最大下单量
	MaxMarketOrderVolume types.TThostFtdcVolumeType
	// 市价单最小下单量
	MinMarketOrderVolume types.TThostFtdcVolumeType
	// 限价单最大下单量
	MaxLimitOrderVolume types.TThostFtdcVolumeType
	// 限价单最小下单量
	MinLimitOrderVolume types.TThostFtdcVolumeType
	// 合约数量乘数
	VolumeMultiple types.TThostFtdcVolumeMultipleType
	// 最小变动价位
	PriceTick types.TThostFtdcPriceType
	// 创建日
	CreateDate types.TThostFtdcDateType
	// 上市日
	OpenDate types.TThostFtdcDateType
	// 到期日
	ExpireDate types.TThostFtdcDateType
	// 开始交割日
	StartDelivDate types.TThostFtdcDateType
	// 结束交割日
	EndDelivDate types.TThostFtdcDateType
	// 合约生命周期状态
	InstLifePhase types.TThostFtdcInstLifePhaseType
	// 当前是否交易
	IsTrading types.TThostFtdcBoolType
	// 持仓类型
	PositionType types.TThostFtdcPositionTypeType
	// 持仓日期类型
	PositionDateType types.TThostFtdcPositionDateTypeType
	// 多头保证金率
	LongMarginRatio types.TThostFtdcRatioType
	// 空头保证金率
	ShortMarginRatio types.TThostFtdcRatioType
	// 是否使用大额单边保证金算法
	MaxMarginSideAlgorithm types.TThostFtdcMaxMarginSideAlgorithmType
	// 基础商品代码
	UnderlyingInstrID types.TThostFtdcInstrumentIDType
	// 执行价
	StrikePrice types.TThostFtdcPriceType
	// 期权类型
	OptionsType types.TThostFtdcOptionsTypeType
	// 合约基础商品乘数
	UnderlyingMultiple types.TThostFtdcUnderlyingMultipleType
	// 组合类型
	CombinationType types.TThostFtdcCombinationTypeType
}

// 申请组合合约信息
type CThostFtdcCombInstrumentField struct {
	// 合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
	// 交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	// 组合类型
	CombinationType types.TThostFtdcDceCombinationTypeType
	// 产品代码
	ProductID types.TThostFtdcInstrumentIDType
	// 保证金优惠比例
	Xparameter types.TThostFtdcRatioType
}

// 投资者RCAMS组合保证金信息
type CThostFtdcRCAMSInvestorProdMarginField struct {
	// 交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	// 商品组代码
	CombProductID types.TThostFtdcInstrumentIDType
	// 投机套保标志
	HedgeFlag types.TThostFtdcHedgeFlagType
	// 商品群代码
	ProductGroupID types.TThostFtdcInstrumentIDType
	// 品种组合前风险
	RiskBeforeDiscount types.TThostFtdcMoneyType
	// 同合约对冲风险
	IntraInstrRisk types.TThostFtdcMoneyType
	// 品种买持仓风险
	BPosRisk types.TThostFtdcMoneyType
	// 品种卖持仓风险
	SPosRisk types.TThostFtdcMoneyType
	// 品种内对冲风险
	IntraProdRisk types.TThostFtdcMoneyType
	// 品种净持仓风险
	NetRisk types.TThostFtdcMoneyType
	// 品种间对冲风险
	InterProdRisk types.TThostFtdcMoneyType
	// 空头期权权利金
	OptionRoyalty types.TThostFtdcMoneyType
	// 交割月期货开仓冻结保证金
	DeliveryOpenFrozenMargin types.TThostFtdcMoneyType
	// 开仓冻结保证金
	OpenFrozenMargin types.TThostFtdcMoneyType
	// 投资者冻结保证金
	UseFrozenMargin types.TThostFtdcMoneyType
	// 投资者冻结保证金
	MMSAExchMargin types.TThostFtdcMoneyType
	// 交割月期货交易所持仓保证金
	DeliveryExchMargin types.TThostFtdcMoneyType
	// 策略组合交易所保证金
	CombExchMargin types.TThostFtdcMoneyType
	// 交易所持仓保证金
	ExchMargin types.TThostFtdcMoneyType
	// 投资者持仓保证金
	UseMargin types.TThostFtdcMoneyType
}

// RCAMS策略组合持仓信息
type CThostFtdcRCAMSInvestorCombPositionField struct {
	// 交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	// 合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
	// 投机套保标志
	HedgeFlag types.TThostFtdcHedgeFlagType
	// 持仓多空方向
	PosiDirection types.TThostFtdcPosiDirectionType
	// 组合合约代码
	CombInstrumentID types.TThostFtdcInstrumentIDType
	// LegID
	LegID types.TThostFtdcLegIDType
	// 交易所组合合约代码
	ExchangeInstID types.TThostFtdcInstrumentIDType
	// 数量
	TotalAmt types.TThostFtdcVolumeType
	// 交易所持仓保证金
	ExchMargin types.TThostFtdcMoneyType
	// 投资者保证金
	Margin types.TThostFtdcMoneyType
}

// 经纪公司
type CThostFtdcBrokerField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 经纪公司简称
	BrokerAbbr types.TThostFtdcBrokerAbbrType
	// 经纪公司名称
	BrokerName types.TThostFtdcBrokerNameType
	// 是否活跃
	IsActive types.TThostFtdcBoolType
}

// 交易所交易员
type CThostFtdcTraderField struct {
	// 交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	// 交易所交易员代码
	TraderID types.TThostFtdcTraderIDType
	// 会员代码
	ParticipantID types.TThostFtdcParticipantIDType
	// 密码
	Password types.TThostFtdcPasswordType
	// 安装数量
	InstallCount types.TThostFtdcInstallCountType
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
}

// 投资者
type CThostFtdcInvestorField struct {
	// 投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 投资者分组代码
	InvestorGroupID types.TThostFtdcInvestorIDType
	// 投资者名称
	InvestorName types.TThostFtdcPartyNameType
	// 证件类型
	IdentifiedCardType types.TThostFtdcIdCardTypeType
	// 证件号码
	IdentifiedCardNo types.TThostFtdcIdentifiedCardNoType
	// 是否活跃
	IsActive types.TThostFtdcBoolType
	// 联系电话
	Telephone types.TThostFtdcTelephoneType
	// 通讯地址
	Address types.TThostFtdcAddressType
	// 开户日期
	OpenDate types.TThostFtdcDateType
	// 手机
	Mobile types.TThostFtdcMobileType
	// 手续费率模板代码
	CommModelID types.TThostFtdcInvestorIDType
	// 保证金率模板代码
	MarginModelID types.TThostFtdcInvestorIDType
}

// 交易编码
type CThostFtdcTradingCodeField struct {
	// 投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	// 客户代码
	ClientID types.TThostFtdcClientIDType
	// 是否活跃
	IsActive types.TThostFtdcBoolType
	// 交易编码类型
	ClientIDType types.TThostFtdcClientIDTypeType
}

// 会员编码和经纪公司编码对照表
type CThostFtdcPartBrokerField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	// 会员代码
	ParticipantID types.TThostFtdcParticipantIDType
	// 是否活跃
	IsActive types.TThostFtdcBoolType
}

// 管理用户
type CThostFtdcSuperUserField struct {
	// 用户代码
	UserID types.TThostFtdcUserIDType
	// 用户名称
	UserName types.TThostFtdcUserNameType
	// 密码
	Password types.TThostFtdcPasswordType
	// 是否活跃
	IsActive types.TThostFtdcBoolType
}

// 管理用户功能权限
type CThostFtdcSuperUserFunctionField struct {
	// 用户代码
	UserID types.TThostFtdcUserIDType
	// 功能代码
	FunctionCode types.TThostFtdcFunctionCodeType
}

// 投资者组
type CThostFtdcInvestorGroupField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 投资者分组代码
	InvestorGroupID types.TThostFtdcInvestorIDType
	// 投资者分组名称
	InvestorGroupName types.TThostFtdcInvestorGroupNameType
}

// 资金账户
type CThostFtdcTradingAccountField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 投资者帐号
	AccountID types.TThostFtdcAccountIDType
	// 上次质押金额
	PreMortgage types.TThostFtdcMoneyType
	// 上次信用额度
	PreCredit types.TThostFtdcMoneyType
	// 上次存款额
	PreDeposit types.TThostFtdcMoneyType
	// 上次结算准备金
	PreBalance types.TThostFtdcMoneyType
	// 上次占用的保证金
	PreMargin types.TThostFtdcMoneyType
	// 利息基数
	InterestBase types.TThostFtdcMoneyType
	// 利息收入
	Interest types.TThostFtdcMoneyType
	// 入金金额
	Deposit types.TThostFtdcMoneyType
	// 出金金额
	Withdraw types.TThostFtdcMoneyType
	// 冻结的保证金
	FrozenMargin types.TThostFtdcMoneyType
	// 冻结的资金
	FrozenCash types.TThostFtdcMoneyType
	// 冻结的手续费
	FrozenCommission types.TThostFtdcMoneyType
	// 当前保证金总额
	CurrMargin types.TThostFtdcMoneyType
	// 资金差额
	CashIn types.TThostFtdcMoneyType
	// 手续费
	Commission types.TThostFtdcMoneyType
	// 平仓盈亏
	CloseProfit types.TThostFtdcMoneyType
	// 持仓盈亏
	PositionProfit types.TThostFtdcMoneyType
	// 期货结算准备金
	Balance types.TThostFtdcMoneyType
	// 可用资金
	Available types.TThostFtdcMoneyType
	// 可取资金
	WithdrawQuota types.TThostFtdcMoneyType
	// 基本准备金
	Reserve types.TThostFtdcMoneyType
	// 交易日
	TradingDay types.TThostFtdcDateType
	// 结算编号
	SettlementID types.TThostFtdcSettlementIDType
	// 信用额度
	Credit types.TThostFtdcMoneyType
	// 质押金额
	Mortgage types.TThostFtdcMoneyType
	// 交易所保证金
	ExchangeMargin types.TThostFtdcMoneyType
	// 投资者交割保证金
	DeliveryMargin types.TThostFtdcMoneyType
	// 交易所交割保证金
	ExchangeDeliveryMargin types.TThostFtdcMoneyType
	// 保底期货结算准备金
	ReserveBalance types.TThostFtdcMoneyType
	// 币种代码
	CurrencyID types.TThostFtdcCurrencyIDType
	// 上次货币质入金额
	PreFundMortgageIn types.TThostFtdcMoneyType
	// 上次货币质出金额
	PreFundMortgageOut types.TThostFtdcMoneyType
	// 货币质入金额
	FundMortgageIn types.TThostFtdcMoneyType
	// 货币质出金额
	FundMortgageOut types.TThostFtdcMoneyType
	// 货币质押余额
	FundMortgageAvailable types.TThostFtdcMoneyType
	// 可质押货币金额
	MortgageableFund types.TThostFtdcMoneyType
	// 特殊产品占用保证金
	SpecProductMargin types.TThostFtdcMoneyType
	// 特殊产品冻结保证金
	SpecProductFrozenMargin types.TThostFtdcMoneyType
	// 特殊产品手续费
	SpecProductCommission types.TThostFtdcMoneyType
	// 特殊产品冻结手续费
	SpecProductFrozenCommission types.TThostFtdcMoneyType
	// 特殊产品持仓盈亏
	SpecProductPositionProfit types.TThostFtdcMoneyType
	// 特殊产品平仓盈亏
	SpecProductCloseProfit types.TThostFtdcMoneyType
	// 根据持仓盈亏算法计算的特殊产品持仓盈亏
	SpecProductPositionProfitByAlg types.TThostFtdcMoneyType
	// 特殊产品交易所保证金
	SpecProductExchangeMargin types.TThostFtdcMoneyType
	// 分仓冻结资金
	FrozenPartition types.TThostFtdcMoneyType
}

// 投资者持仓
type CThostFtdcInvestorPositionField struct {
	// 合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	// 持仓多空方向
	PosiDirection types.TThostFtdcPosiDirectionType
	// 投机套保标志
	HedgeFlag types.TThostFtdcHedgeFlagType
	// 持仓日期
	PositionDate types.TThostFtdcPositionDateType
	// 上日持仓
	YdPosition types.TThostFtdcVolumeType
	// 今日持仓
	Position types.TThostFtdcVolumeType
	// 多头冻结
	LongFrozen types.TThostFtdcVolumeType
	// 空头冻结
	ShortFrozen types.TThostFtdcVolumeType
	// 开仓冻结金额
	LongFrozenAmount types.TThostFtdcMoneyType
	// 开仓冻结金额
	ShortFrozenAmount types.TThostFtdcMoneyType
	// 开仓量
	OpenVolume types.TThostFtdcVolumeType
	// 平仓量
	CloseVolume types.TThostFtdcVolumeType
	// 开仓金额
	OpenAmount types.TThostFtdcMoneyType
	// 平仓金额
	CloseAmount types.TThostFtdcMoneyType
	// 持仓成本
	PositionCost types.TThostFtdcMoneyType
	// 上次占用的保证金
	PreMargin types.TThostFtdcMoneyType
	// 占用的保证金
	UseMargin types.TThostFtdcMoneyType
	// 冻结的保证金
	FrozenMargin types.TThostFtdcMoneyType
	// 冻结的资金
	FrozenCash types.TThostFtdcMoneyType
	// 冻结的手续费
	FrozenCommission types.TThostFtdcMoneyType
	// 资金差额
	CashIn types.TThostFtdcMoneyType
	// 手续费
	Commission types.TThostFtdcMoneyType
	// 平仓盈亏
	CloseProfit types.TThostFtdcMoneyType
	// 持仓盈亏
	PositionProfit types.TThostFtdcMoneyType
	// 上次结算价
	PreSettlementPrice types.TThostFtdcPriceType
	// 本次结算价
	SettlementPrice types.TThostFtdcPriceType
	// 交易日
	TradingDay types.TThostFtdcDateType
	// 结算编号
	SettlementID types.TThostFtdcSettlementIDType
	// 开仓成本
	OpenCost types.TThostFtdcMoneyType
	// 交易所保证金
	ExchangeMargin types.TThostFtdcMoneyType
	// 组合成交形成的持仓
	CombPosition types.TThostFtdcVolumeType
	// 组合多头冻结
	CombLongFrozen types.TThostFtdcVolumeType
	// 组合空头冻结
	CombShortFrozen types.TThostFtdcVolumeType
	// 逐日盯市平仓盈亏
	CloseProfitByDate types.TThostFtdcMoneyType
	// 逐笔对冲平仓盈亏
	CloseProfitByTrade types.TThostFtdcMoneyType
	// 今日持仓
	TodayPosition types.TThostFtdcVolumeType
	// 保证金率
	MarginRateByMoney types.TThostFtdcRatioType
	// 保证金率(按手数)
	MarginRateByVolume types.TThostFtdcRatioType
	// 执行冻结
	StrikeFrozen types.TThostFtdcVolumeType
	// 执行冻结金额
	StrikeFrozenAmount types.TThostFtdcMoneyType
	// 放弃执行冻结
	AbandonFrozen types.TThostFtdcVolumeType
}

// 合约保证金率
type CThostFtdcInstrumentMarginRateField struct {
	// 合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
	// 投资者范围
	InvestorRange types.TThostFtdcInvestorRangeType
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	// 投机套保标志
	HedgeFlag types.TThostFtdcHedgeFlagType
	// 多头保证金率
	LongMarginRatioByMoney types.TThostFtdcRatioType
	// 多头保证金费
	LongMarginRatioByVolume types.TThostFtdcMoneyType
	// 空头保证金率
	ShortMarginRatioByMoney types.TThostFtdcRatioType
	// 空头保证金费
	ShortMarginRatioByVolume types.TThostFtdcMoneyType
	// 是否相对交易所收取
	IsRelative types.TThostFtdcBoolType
}

// 合约手续费率
type CThostFtdcInstrumentCommissionRateField struct {
	// 合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
	// 投资者范围
	InvestorRange types.TThostFtdcInvestorRangeType
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	// 开仓手续费率
	OpenRatioByMoney types.TThostFtdcRatioType
	// 开仓手续费
	OpenRatioByVolume types.TThostFtdcRatioType
	// 平仓手续费率
	CloseRatioByMoney types.TThostFtdcRatioType
	// 平仓手续费
	CloseRatioByVolume types.TThostFtdcRatioType
	// 平今手续费率
	CloseTodayRatioByMoney types.TThostFtdcRatioType
	// 平今手续费
	CloseTodayRatioByVolume types.TThostFtdcRatioType
}

// SPBM商品组保证金明细
type CThostFtdcInvestorProdSPBMDetailField struct {
	// 交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	// 品种代码
	ProdFamilyCode types.TThostFtdcInstrumentIDType
	// 合约内对锁保证金
	IntraInstrMargin types.TThostFtdcMoneyType
	// 买方向归集保证金
	BCollectingMargin types.TThostFtdcMoneyType
	// 卖方向归集保证金
	SCollectingMargin types.TThostFtdcMoneyType
	// 品种内合约间对锁保证金
	IntraProdMargin types.TThostFtdcMoneyType
	// 净保证金
	NetMargin types.TThostFtdcMoneyType
	// 产品间对锁保证金
	InterProdMargin types.TThostFtdcMoneyType
	// 裸保证金
	SingleMargin types.TThostFtdcMoneyType
	// 附加保证金
	AddOnMargin types.TThostFtdcMoneyType
	// 交割保证金
	DeliveryMargin types.TThostFtdcMoneyType
	// 卖方期权最低风险
	OptionMinRisk types.TThostFtdcMoneyType
	// 价值冲抵
	RealOptionValueOffset types.TThostFtdcMoneyType
	// 保证金
	Margin types.TThostFtdcMoneyType
	// 交易所保证金
	ExchMargin types.TThostFtdcMoneyType
}

// SPMM商品群保证金明细
type CThostFtdcSPMMInvestorCommodityGroupMarginField struct {
	// 交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	// 商品群代码
	CommodityGroupID types.TThostFtdcSPMMProductIDType
	// 优惠仓位应收保证金
	MarginBeforeDiscount types.TThostFtdcMoneyType
	// 不优惠仓位应收保证金
	MarginNoDiscount types.TThostFtdcMoneyType
	// 多头风险
	LongRisk types.TThostFtdcMoneyType
	// 空头风险
	ShortRisk types.TThostFtdcMoneyType
	// 商品群平仓冻结保证金
	CloseFrozenMargin types.TThostFtdcMoneyType
	// SPMM跨品种优惠系数
	InterCommodityRate types.TThostFtdcMoneyType
	// 商品群最小保证金比例
	MiniMarginRatio types.TThostFtdcMoneyType
	// 投资者保证金和交易所保证金的比例
	AdjustRatio types.TThostFtdcMoneyType
	// SPMM品种内优惠汇总
	IntraCommodityDiscount types.TThostFtdcMoneyType
	// SPMM跨品种优惠
	InterCommodityDiscount types.TThostFtdcMoneyType
	// 交易所保证金
	ExchMargin types.TThostFtdcMoneyType
	// 投资者保证金
	InvestorMargin types.TThostFtdcMoneyType
	// 冻结的手续费
	FrozenCommission types.TThostFtdcMoneyType
	// 手续费
	Commission types.TThostFtdcMoneyType
	// 冻结的资金
	FrozenCash types.TThostFtdcMoneyType
	// 资金差额
	CashIn types.TThostFtdcMoneyType
	// 行权冻结资金
	StrikeFrozenMargin types.TThostFtdcMoneyType
}

type CThostFtdcRULEInvestorProdMarginField struct {
	// 交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	// 品种代码
	ProdFamilyCode types.TThostFtdcProductIDType
	// 合约类型
	InstrumentClass types.TThostFtdcRULEInstrClassType
	// 商品群号
	CommodityGroupID types.TThostFtdcRULECommodityGroupIDType
	// 品种内对锁保证金
	IntraProdMargin types.TThostFtdcMoneyType
	// 品种间对锁保证金
	InterProdMargin types.TThostFtdcMoneyType
	// 跨品种单腿保证金
	SingleMargin types.TThostFtdcMoneyType
	// 非组合合约保证金
	NonCombMargin types.TThostFtdcMoneyType
	// 附加保证金
	AddOnMargin types.TThostFtdcMoneyType
	// 交易所保证金
	ExchMargin types.TThostFtdcMoneyType
	// 开仓冻结保证金
	OpenFrozenMargin types.TThostFtdcMoneyType
	// 平仓冻结保证金
	CloseFrozenMargin types.TThostFtdcMoneyType
	// 品种保证金
	Margin types.TThostFtdcMoneyType
	// 冻结保证金
	FrozenMargin types.TThostFtdcMoneyType
}

// 深度行情
type CThostFtdcDepthMarketDataField struct {
	// 交易日
	TradingDay types.TThostFtdcDateType
	// 合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
	// 交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	// 合约在交易所的代码
	ExchangeInstID types.TThostFtdcExchangeInstIDType
	// 最新价
	LastPrice types.TThostFtdcPriceType
	// 上次结算价
	PreSettlementPrice types.TThostFtdcPriceType
	// 昨收盘
	PreClosePrice types.TThostFtdcPriceType
	// 昨持仓量
	PreOpenInterest types.TThostFtdcLargeVolumeType
	// 今开盘
	OpenPrice types.TThostFtdcPriceType
	// 最高价
	HighestPrice types.TThostFtdcPriceType
	// 最低价
	LowestPrice types.TThostFtdcPriceType
	// 数量
	Volume types.TThostFtdcVolumeType
	// 成交金额
	Turnover types.TThostFtdcMoneyType
	// 持仓量
	OpenInterest types.TThostFtdcLargeVolumeType
	// 今收盘
	ClosePrice types.TThostFtdcPriceType
	// 本次结算价
	SettlementPrice types.TThostFtdcPriceType
	// 涨停板价
	UpperLimitPrice types.TThostFtdcPriceType
	// 跌停板价
	LowerLimitPrice types.TThostFtdcPriceType
	// 昨虚实度
	PreDelta types.TThostFtdcRatioType
	// 今虚实度
	CurrDelta types.TThostFtdcRatioType
	// 最后修改时间
	UpdateTime types.TThostFtdcTimeType
	// 最后修改毫秒
	UpdateMillisec types.TThostFtdcMillisecType
	// 申买价一
	BidPrice1 types.TThostFtdcPriceType
	// 申买量一
	BidVolume1 types.TThostFtdcVolumeType
	// 申卖价一
	AskPrice1 types.TThostFtdcPriceType
	// 申卖量一
	AskVolume1 types.TThostFtdcVolumeType
	// 申买价二
	BidPrice2 types.TThostFtdcPriceType
	// 申买量二
	BidVolume2 types.TThostFtdcVolumeType
	// 申卖价二
	AskPrice2 types.TThostFtdcPriceType
	// 申卖量二
	AskVolume2 types.TThostFtdcVolumeType
	// 申买价三
	BidPrice3 types.TThostFtdcPriceType
	// 申买量三
	BidVolume3 types.TThostFtdcVolumeType
	// 申卖价三
	AskPrice3 types.TThostFtdcPriceType
	// 申卖量三
	AskVolume3 types.TThostFtdcVolumeType
	// 申买价四
	BidPrice4 types.TThostFtdcPriceType
	// 申买量四
	BidVolume4 types.TThostFtdcVolumeType
	// 申卖价四
	AskPrice4 types.TThostFtdcPriceType
	// 申卖量四
	AskVolume4 types.TThostFtdcVolumeType
	// 申买价五
	BidPrice5 types.TThostFtdcPriceType
	// 申买量五
	BidVolume5 types.TThostFtdcVolumeType
	// 申卖价五
	AskPrice5 types.TThostFtdcPriceType
	// 申卖量五
	AskVolume5 types.TThostFtdcVolumeType
	// 当日均价
	AveragePrice types.TThostFtdcPriceType
	// 业务日期
	ActionDay types.TThostFtdcDateType
	// 上带价
	BandingUpperPrice types.TThostFtdcPriceType
	// 下带价
	BandingLowerPrice types.TThostFtdcPriceType
}

// 投资者合约交易权限
type CThostFtdcInstrumentTradingRightField struct {
	// 合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
	// 投资者范围
	InvestorRange types.TThostFtdcInvestorRangeType
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	// 交易权限
	TradingRight types.TThostFtdcTradingRightType
}

// 经纪公司用户
type CThostFtdcBrokerUserField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 用户代码
	UserID types.TThostFtdcUserIDType
	// 用户名称
	UserName types.TThostFtdcUserNameType
	// 用户类型
	UserType types.TThostFtdcUserTypeType
	// 是否活跃
	IsActive types.TThostFtdcBoolType
	// 是否使用令牌
	IsUsingOTP types.TThostFtdcBoolType
}

// 交易开关设置
type CThostFtdcControlParamField struct {
	// 投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	// 开关代码
	ControlParamID types.TThostFtdcControlParamIDType
	// 开关设置
	ControlParamValue types.TThostFtdcControlParamValueType
	// 备注说明
	Memo types.TThostFtdcMemoType
}

// 经纪公司用户口令
type CThostFtdcBrokerUserPasswordField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 用户代码
	UserID types.TThostFtdcUserIDType
	// 密码
	Password types.TThostFtdcPasswordType
}

// 经纪公司用户功能权限
type CThostFtdcBrokerUserFunctionField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 用户代码
	UserID types.TThostFtdcUserIDType
	// 经纪公司功能代码
	BrokerFunctionCode types.TThostFtdcBrokerFunctionCodeType
}

// 交易所交易员报盘机
type CThostFtdcTraderOfferField struct {
	// 交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	// 交易所交易员代码
	TraderID types.TThostFtdcTraderIDType
	// 会员代码
	ParticipantID types.TThostFtdcParticipantIDType
	// 密码
	Password types.TThostFtdcPasswordType
	// 安装编号
	InstallID types.TThostFtdcInstallIDType
	// 本地报单编号
	OrderLocalID types.TThostFtdcOrderLocalIDType
	// 交易所交易员连接状态
	TraderConnectStatus types.TThostFtdcTraderConnectStatusType
	// 发出连接请求的日期
	ConnectRequestDate types.TThostFtdcDateType
	// 发出连接请求的时间
	ConnectRequestTime types.TThostFtdcTimeType
	// 上次报告日期
	LastReportDate types.TThostFtdcDateType
	// 上次报告时间
	LastReportTime types.TThostFtdcTimeType
	// 完成连接日期
	ConnectDate types.TThostFtdcDateType
	// 完成连接时间
	ConnectTime types.TThostFtdcTimeType
	// 启动日期
	StartDate types.TThostFtdcDateType
	// 启动时间
	StartTime types.TThostFtdcTimeType
	// 交易日
	TradingDay types.TThostFtdcDateType
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 本席位最大成交编号
	MaxTradeID types.TThostFtdcTradeIDType
	// 本席位最大报单备拷
	MaxOrderMessageReference types.TThostFtdcReturnCodeType
}

// 投资者结算结果
type CThostFtdcSettlementInfoField struct {
	// 交易日
	TradingDay types.TThostFtdcDateType
	// 结算编号
	SettlementID types.TThostFtdcSettlementIDType
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	// 序号
	SequenceNo types.TThostFtdcSequenceNoType
	// 消息正文
	Content types.TThostFtdcContentType
}

// 合约保证金率调整
type CThostFtdcInstrumentMarginRateAdjustField struct {
	// 合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
	// 投资者范围
	InvestorRange types.TThostFtdcInvestorRangeType
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	// 投机套保标志
	HedgeFlag types.TThostFtdcHedgeFlagType
	// 多头保证金率
	LongMarginRatioByMoney types.TThostFtdcRatioType
	// 多头保证金费
	LongMarginRatioByVolume types.TThostFtdcMoneyType
	// 空头保证金率
	ShortMarginRatioByMoney types.TThostFtdcRatioType
	// 空头保证金费
	ShortMarginRatioByVolume types.TThostFtdcMoneyType
	// 是否相对交易所收取
	IsRelative types.TThostFtdcBoolType
}

// 交易所保证金率
type CThostFtdcExchangeMarginRateField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
	// 投机套保标志
	HedgeFlag types.TThostFtdcHedgeFlagType
	// 多头保证金率
	LongMarginRatioByMoney types.TThostFtdcRatioType
	// 多头保证金费
	LongMarginRatioByVolume types.TThostFtdcMoneyType
	// 空头保证金率
	ShortMarginRatioByMoney types.TThostFtdcRatioType
	// 空头保证金费
	ShortMarginRatioByVolume types.TThostFtdcMoneyType
}

// 交易所保证金率调整
type CThostFtdcExchangeMarginRateAdjustField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
	// 投机套保标志
	HedgeFlag types.TThostFtdcHedgeFlagType
	// 跟随交易所投资者多头保证金率
	LongMarginRatioByMoney types.TThostFtdcRatioType
	// 跟随交易所投资者多头保证金费
	LongMarginRatioByVolume types.TThostFtdcMoneyType
	// 跟随交易所投资者空头保证金率
	ShortMarginRatioByMoney types.TThostFtdcRatioType
	// 跟随交易所投资者空头保证金费
	ShortMarginRatioByVolume types.TThostFtdcMoneyType
	// 交易所多头保证金率
	ExchLongMarginRatioByMoney types.TThostFtdcRatioType
	// 交易所多头保证金费
	ExchLongMarginRatioByVolume types.TThostFtdcMoneyType
	// 交易所空头保证金率
	ExchShortMarginRatioByMoney types.TThostFtdcRatioType
	// 交易所空头保证金费
	ExchShortMarginRatioByVolume types.TThostFtdcMoneyType
	// 不跟随交易所投资者多头保证金率
	NoLongMarginRatioByMoney types.TThostFtdcRatioType
	// 不跟随交易所投资者多头保证金费
	NoLongMarginRatioByVolume types.TThostFtdcMoneyType
	// 不跟随交易所投资者空头保证金率
	NoShortMarginRatioByMoney types.TThostFtdcRatioType
	// 不跟随交易所投资者空头保证金费
	NoShortMarginRatioByVolume types.TThostFtdcMoneyType
}

// 汇率
type CThostFtdcExchangeRateField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 源币种
	FromCurrencyID types.TThostFtdcCurrencyIDType
	// 源币种单位数量
	FromCurrencyUnit types.TThostFtdcCurrencyUnitType
	// 目标币种
	ToCurrencyID types.TThostFtdcCurrencyIDType
	// 汇率
	ExchangeRate types.TThostFtdcExchangeRateType
}

// 结算引用
type CThostFtdcSettlementRefField struct {
	// 交易日
	TradingDay types.TThostFtdcDateType
	// 结算编号
	SettlementID types.TThostFtdcSettlementIDType
}

// 当前时间
type CThostFtdcCurrentTimeField struct {
	// 当前日期
	CurrDate types.TThostFtdcDateType
	// 当前时间
	CurrTime types.TThostFtdcTimeType
	// 当前时间（毫秒）
	CurrMillisec types.TThostFtdcMillisecType
	// 业务日期
	ActionDay types.TThostFtdcDateType
}

// 通讯阶段
type CThostFtdcCommPhaseField struct {
	// 交易日
	TradingDay types.TThostFtdcDateType
	// 通讯时段编号
	CommPhaseNo types.TThostFtdcCommPhaseNoType
	// 系统编号
	SystemID types.TThostFtdcSystemIDType
}

// 登录信息
type CThostFtdcLoginInfoField struct {
	// 前置编号
	FrontID types.TThostFtdcFrontIDType
	// 会话编号
	SessionID types.TThostFtdcSessionIDType
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 用户代码
	UserID types.TThostFtdcUserIDType
	// 登录日期
	LoginDate types.TThostFtdcDateType
	// 登录时间
	LoginTime types.TThostFtdcTimeType
	// IP地址
	IPAddress types.TThostFtdcIPAddressType
	// 用户端产品信息
	UserProductInfo types.TThostFtdcProductInfoType
	// 接口端产品信息
	InterfaceProductInfo types.TThostFtdcProductInfoType
	// 协议信息
	ProtocolInfo types.TThostFtdcProtocolInfoType
	// 系统名称
	SystemName types.TThostFtdcSystemNameType
	// 密码
	Password types.TThostFtdcPasswordType
	// 最大报单引用
	MaxOrderRef types.TThostFtdcOrderRefType
	// 上期所时间
	SHFETime types.TThostFtdcTimeType
	// 大商所时间
	DCETime types.TThostFtdcTimeType
	// 郑商所时间
	CZCETime types.TThostFtdcTimeType
	// 中金所时间
	FFEXTime types.TThostFtdcTimeType
	// Mac地址
	MacAddress types.TThostFtdcMacAddressType
	// 动态密码
	OneTimePassword types.TThostFtdcPasswordType
	// 能源中心时间
	INETime types.TThostFtdcTimeType
	// 查询时是否需要流控
	IsQryControl types.TThostFtdcBoolType
	// 登录备注
	LoginRemark types.TThostFtdcLoginRemarkType
}

// 登录信息
type CThostFtdcLogoutAllField struct {
	// 前置编号
	FrontID types.TThostFtdcFrontIDType
	// 会话编号
	SessionID types.TThostFtdcSessionIDType
	// 系统名称
	SystemName types.TThostFtdcSystemNameType
}

// 前置状态
type CThostFtdcFrontStatusField struct {
	// 前置编号
	FrontID types.TThostFtdcFrontIDType
	// 上次报告日期
	LastReportDate types.TThostFtdcDateType
	// 上次报告时间
	LastReportTime types.TThostFtdcTimeType
	// 是否活跃
	IsActive types.TThostFtdcBoolType
}

// 用户口令变更
type CThostFtdcUserPasswordUpdateField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 用户代码
	UserID types.TThostFtdcUserIDType
	// 原来的口令
	OldPassword types.TThostFtdcPasswordType
	// 新的口令
	NewPassword types.TThostFtdcPasswordType
}

// 输入报单
type CThostFtdcInputOrderField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	// 合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
	// 报单引用
	OrderRef types.TThostFtdcOrderRefType
	// 用户代码
	UserID types.TThostFtdcUserIDType
	// 报单价格条件
	OrderPriceType types.TThostFtdcOrderPriceTypeType
	// 买卖方向
	Direction types.TThostFtdcDirectionType
	// 组合开平标志
	CombOffsetFlag types.TThostFtdcCombOffsetFlagType
	// 组合投机套保标志
	CombHedgeFlag types.TThostFtdcCombHedgeFlagType
	// 价格
	LimitPrice types.TThostFtdcPriceType
	// 数量
	VolumeTotalOriginal types.TThostFtdcVolumeType
	// 有效期类型
	TimeCondition types.TThostFtdcTimeConditionType
	// GTD日期
	GTDDate types.TThostFtdcDateType
	// 成交量类型
	VolumeCondition types.TThostFtdcVolumeConditionType
	// 最小成交量
	MinVolume types.TThostFtdcVolumeType
	// 触发条件
	ContingentCondition types.TThostFtdcContingentConditionType
	// 止损价
	StopPrice types.TThostFtdcPriceType
	// 强平原因
	ForceCloseReason types.TThostFtdcForceCloseReasonType
	// 自动挂起标志
	IsAutoSuspend types.TThostFtdcBoolType
	// 业务单元
	BusinessUnit types.TThostFtdcBusinessUnitType
	// 请求编号
	RequestID types.TThostFtdcRequestIDType
	// 用户强评标志
	UserForceClose types.TThostFtdcBoolType
	// 互换单标志
	IsSwapOrder types.TThostFtdcBoolType
	// 交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	// 投资单元代码
	InvestUnitID types.TThostFtdcInvestUnitIDType
	// 资金账号
	AccountID types.TThostFtdcAccountIDType
	// 币种代码
	CurrencyID types.TThostFtdcCurrencyIDType
	// 交易编码
	ClientID types.TThostFtdcClientIDType
	// IP地址
	IPAddress types.TThostFtdcIPAddressType
	// Mac地址
	MacAddress types.TThostFtdcMacAddressType
	// 自定义字段
	CustomOrderRef types.TThostFtdcCustomOrderRefType
	// 报单标志
	OrderFlag types.TThostFtdcOrderFlagType
}

// 报单
type CThostFtdcOrderField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	// 合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
	// 报单引用
	OrderRef types.TThostFtdcOrderRefType
	// 用户代码
	UserID types.TThostFtdcUserIDType
	// 报单价格条件
	OrderPriceType types.TThostFtdcOrderPriceTypeType
	// 买卖方向
	Direction types.TThostFtdcDirectionType
	// 组合开平标志
	CombOffsetFlag types.TThostFtdcCombOffsetFlagType
	// 组合投机套保标志
	CombHedgeFlag types.TThostFtdcCombHedgeFlagType
	// 价格
	LimitPrice types.TThostFtdcPriceType
	// 数量
	VolumeTotalOriginal types.TThostFtdcVolumeType
	// 有效期类型
	TimeCondition types.TThostFtdcTimeConditionType
	// GTD日期
	GTDDate types.TThostFtdcDateType
	// 成交量类型
	VolumeCondition types.TThostFtdcVolumeConditionType
	// 最小成交量
	MinVolume types.TThostFtdcVolumeType
	// 触发条件
	ContingentCondition types.TThostFtdcContingentConditionType
	// 止损价
	StopPrice types.TThostFtdcPriceType
	// 强平原因
	ForceCloseReason types.TThostFtdcForceCloseReasonType
	// 自动挂起标志
	IsAutoSuspend types.TThostFtdcBoolType
	// 业务单元
	BusinessUnit types.TThostFtdcBusinessUnitType
	// 请求编号
	RequestID types.TThostFtdcRequestIDType
	// 本地报单编号
	OrderLocalID types.TThostFtdcOrderLocalIDType
	// 交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	// 会员代码
	ParticipantID types.TThostFtdcParticipantIDType
	// 客户代码
	ClientID types.TThostFtdcClientIDType
	// 合约在交易所的代码
	ExchangeInstID types.TThostFtdcExchangeInstIDType
	// 交易所交易员代码
	TraderID types.TThostFtdcTraderIDType
	// 安装编号
	InstallID types.TThostFtdcInstallIDType
	// 报单提交状态
	OrderSubmitStatus types.TThostFtdcOrderSubmitStatusType
	// 报单提示序号
	NotifySequence types.TThostFtdcSequenceNoType
	// 交易日
	TradingDay types.TThostFtdcDateType
	// 结算编号
	SettlementID types.TThostFtdcSettlementIDType
	// 报单编号
	OrderSysID types.TThostFtdcOrderSysIDType
	// 报单来源
	OrderSource types.TThostFtdcOrderSourceType
	// 报单状态
	OrderStatus types.TThostFtdcOrderStatusType
	// 报单类型
	OrderType types.TThostFtdcOrderTypeType
	// 今成交数量
	VolumeTraded types.TThostFtdcVolumeType
	// 剩余数量
	VolumeTotal types.TThostFtdcVolumeType
	// 报单日期
	InsertDate types.TThostFtdcDateType
	// 委托时间
	InsertTime types.TThostFtdcTimeType
	// 激活时间
	ActiveTime types.TThostFtdcTimeType
	// 挂起时间
	SuspendTime types.TThostFtdcTimeType
	// 最后修改时间
	UpdateTime types.TThostFtdcTimeType
	// 撤销时间
	CancelTime types.TThostFtdcTimeType
	// 最后修改交易所交易员代码
	ActiveTraderID types.TThostFtdcTraderIDType
	// 结算会员编号
	ClearingPartID types.TThostFtdcParticipantIDType
	// 序号
	SequenceNo types.TThostFtdcSequenceNoType
	// 前置编号
	FrontID types.TThostFtdcFrontIDType
	// 会话编号
	SessionID types.TThostFtdcSessionIDType
	// 用户端产品信息
	UserProductInfo types.TThostFtdcProductInfoType
	// 状态信息
	StatusMsg types.TThostFtdcErrorMsgType
	// 用户强评标志
	UserForceClose types.TThostFtdcBoolType
	// 操作用户代码
	ActiveUserID types.TThostFtdcUserIDType
	// 经纪公司报单编号
	BrokerOrderSeq types.TThostFtdcSequenceNoType
	// 相关报单
	RelativeOrderSysID types.TThostFtdcOrderSysIDType
	// 郑商所成交数量
	ZCETotalTradedVolume types.TThostFtdcVolumeType
	// 互换单标志
	IsSwapOrder types.TThostFtdcBoolType
	// 营业部编号
	BranchID types.TThostFtdcBranchIDType
	// 投资单元代码
	InvestUnitID types.TThostFtdcInvestUnitIDType
	// 资金账号
	AccountID types.TThostFtdcAccountIDType
	// 币种代码
	CurrencyID types.TThostFtdcCurrencyIDType
	// IP地址
	IPAddress types.TThostFtdcIPAddressType
	// Mac地址
	MacAddress types.TThostFtdcMacAddressType
	// 自定义字段
	CustomOrderRef types.TThostFtdcCustomOrderRefType
	// 成交均价
	TradeAvgPrice types.TThostFtdcPriceType
}

// 交易所报单
type CThostFtdcExchangeOrderField struct {
	// 报单价格条件
	OrderPriceType types.TThostFtdcOrderPriceTypeType
	// 买卖方向
	Direction types.TThostFtdcDirectionType
	// 组合开平标志
	CombOffsetFlag types.TThostFtdcCombOffsetFlagType
	// 组合投机套保标志
	CombHedgeFlag types.TThostFtdcCombHedgeFlagType
	// 价格
	LimitPrice types.TThostFtdcPriceType
	// 数量
	VolumeTotalOriginal types.TThostFtdcVolumeType
	// 有效期类型
	TimeCondition types.TThostFtdcTimeConditionType
	// GTD日期
	GTDDate types.TThostFtdcDateType
	// 成交量类型
	VolumeCondition types.TThostFtdcVolumeConditionType
	// 最小成交量
	MinVolume types.TThostFtdcVolumeType
	// 触发条件
	ContingentCondition types.TThostFtdcContingentConditionType
	// 止损价
	StopPrice types.TThostFtdcPriceType
	// 强平原因
	ForceCloseReason types.TThostFtdcForceCloseReasonType
	// 自动挂起标志
	IsAutoSuspend types.TThostFtdcBoolType
	// 业务单元
	BusinessUnit types.TThostFtdcBusinessUnitType
	// 请求编号
	RequestID types.TThostFtdcRequestIDType
	// 本地报单编号
	OrderLocalID types.TThostFtdcOrderLocalIDType
	// 交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	// 会员代码
	ParticipantID types.TThostFtdcParticipantIDType
	// 客户代码
	ClientID types.TThostFtdcClientIDType
	// 合约在交易所的代码
	ExchangeInstID types.TThostFtdcExchangeInstIDType
	// 交易所交易员代码
	TraderID types.TThostFtdcTraderIDType
	// 安装编号
	InstallID types.TThostFtdcInstallIDType
	// 报单提交状态
	OrderSubmitStatus types.TThostFtdcOrderSubmitStatusType
	// 报单提示序号
	NotifySequence types.TThostFtdcSequenceNoType
	// 交易日
	TradingDay types.TThostFtdcDateType
	// 结算编号
	SettlementID types.TThostFtdcSettlementIDType
	// 报单编号
	OrderSysID types.TThostFtdcOrderSysIDType
	// 报单来源
	OrderSource types.TThostFtdcOrderSourceType
	// 报单状态
	OrderStatus types.TThostFtdcOrderStatusType
	// 报单类型
	OrderType types.TThostFtdcOrderTypeType
	// 今成交数量
	VolumeTraded types.TThostFtdcVolumeType
	// 剩余数量
	VolumeTotal types.TThostFtdcVolumeType
	// 报单日期
	InsertDate types.TThostFtdcDateType
	// 委托时间
	InsertTime types.TThostFtdcTimeType
	// 激活时间
	ActiveTime types.TThostFtdcTimeType
	// 挂起时间
	SuspendTime types.TThostFtdcTimeType
	// 最后修改时间
	UpdateTime types.TThostFtdcTimeType
	// 撤销时间
	CancelTime types.TThostFtdcTimeType
	// 最后修改交易所交易员代码
	ActiveTraderID types.TThostFtdcTraderIDType
	// 结算会员编号
	ClearingPartID types.TThostFtdcParticipantIDType
	// 序号
	SequenceNo types.TThostFtdcSequenceNoType
	// 营业部编号
	BranchID types.TThostFtdcBranchIDType
	// IP地址
	IPAddress types.TThostFtdcIPAddressType
	// Mac地址
	MacAddress types.TThostFtdcMacAddressType
}

// 交易所报单插入失败
type CThostFtdcExchangeOrderInsertErrorField struct {
	// 交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	// 会员代码
	ParticipantID types.TThostFtdcParticipantIDType
	// 交易所交易员代码
	TraderID types.TThostFtdcTraderIDType
	// 安装编号
	InstallID types.TThostFtdcInstallIDType
	// 本地报单编号
	OrderLocalID types.TThostFtdcOrderLocalIDType
	// 错误代码
	ErrorID types.TThostFtdcErrorIDType
	// 错误信息
	ErrorMsg types.TThostFtdcErrorMsgType
}

// 输入报单操作
type CThostFtdcInputOrderActionField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	// 报单操作引用
	OrderActionRef types.TThostFtdcOrderActionRefType
	// 报单引用
	OrderRef types.TThostFtdcOrderRefType
	// 请求编号
	RequestID types.TThostFtdcRequestIDType
	// 前置编号
	FrontID types.TThostFtdcFrontIDType
	// 会话编号
	SessionID types.TThostFtdcSessionIDType
	// 交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	// 报单编号
	OrderSysID types.TThostFtdcOrderSysIDType
	// 操作标志
	ActionFlag types.TThostFtdcActionFlagType
	// 价格
	LimitPrice types.TThostFtdcPriceType
	// 数量变化
	VolumeChange types.TThostFtdcVolumeType
	// 用户代码
	UserID types.TThostFtdcUserIDType
	// 合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
	// 投资单元代码
	InvestUnitID types.TThostFtdcInvestUnitIDType
	// IP地址
	IPAddress types.TThostFtdcIPAddressType
	// Mac地址
	MacAddress types.TThostFtdcMacAddressType
}

// 做市商输入报单操作
type CThostFtdcMKInputOrderActionField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	// 报单操作引用
	OrderActionRef types.TThostFtdcOrderActionRefType
	// 请求编号
	RequestID types.TThostFtdcRequestIDType
	// 前置编号
	FrontID types.TThostFtdcFrontIDType
	// 会话编号
	SessionID types.TThostFtdcSessionIDType
	// 交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	// 操作标志
	ActionFlag types.TThostFtdcActionFlagType
	// 操作模式
	ActionMode types.TThostFtdcMKActionModeType
	// 下限
	LowerLimit types.TThostFtdcLowerLimitType
	// 上限
	UpperLimit types.TThostFtdcUpperLimitType
	// 用户代码
	UserID types.TThostFtdcUserIDType
	// 产品代码
	ProductID types.TThostFtdcInstrumentIDType
	// 合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
	// 报单状态
	OrderStatus types.TThostFtdcOrderStatusType
	// 投资单元代码
	InvestUnitID types.TThostFtdcInvestUnitIDType
	// IP地址
	IPAddress types.TThostFtdcIPAddressType
	// Mac地址
	MacAddress types.TThostFtdcMacAddressType
}

// 报单操作
type CThostFtdcOrderActionField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	// 报单操作引用
	OrderActionRef types.TThostFtdcOrderActionRefType
	// 报单引用
	OrderRef types.TThostFtdcOrderRefType
	// 请求编号
	RequestID types.TThostFtdcRequestIDType
	// 前置编号
	FrontID types.TThostFtdcFrontIDType
	// 会话编号
	SessionID types.TThostFtdcSessionIDType
	// 交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	// 报单编号
	OrderSysID types.TThostFtdcOrderSysIDType
	// 操作标志
	ActionFlag types.TThostFtdcActionFlagType
	// 价格
	LimitPrice types.TThostFtdcPriceType
	// 数量变化
	VolumeChange types.TThostFtdcVolumeType
	// 操作日期
	ActionDate types.TThostFtdcDateType
	// 操作时间
	ActionTime types.TThostFtdcTimeType
	// 交易所交易员代码
	TraderID types.TThostFtdcTraderIDType
	// 安装编号
	InstallID types.TThostFtdcInstallIDType
	// 本地报单编号
	OrderLocalID types.TThostFtdcOrderLocalIDType
	// 操作本地编号
	ActionLocalID types.TThostFtdcOrderLocalIDType
	// 会员代码
	ParticipantID types.TThostFtdcParticipantIDType
	// 客户代码
	ClientID types.TThostFtdcClientIDType
	// 业务单元
	BusinessUnit types.TThostFtdcBusinessUnitType
	// 报单操作状态
	OrderActionStatus types.TThostFtdcOrderActionStatusType
	// 用户代码
	UserID types.TThostFtdcUserIDType
	// 状态信息
	StatusMsg types.TThostFtdcErrorMsgType
	// 合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
	// 营业部编号
	BranchID types.TThostFtdcBranchIDType
	// 投资单元代码
	InvestUnitID types.TThostFtdcInvestUnitIDType
	// IP地址
	IPAddress types.TThostFtdcIPAddressType
	// Mac地址
	MacAddress types.TThostFtdcMacAddressType
}

// 交易所报单操作
type CThostFtdcExchangeOrderActionField struct {
	// 交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	// 报单编号
	OrderSysID types.TThostFtdcOrderSysIDType
	// 操作标志
	ActionFlag types.TThostFtdcActionFlagType
	// 价格
	LimitPrice types.TThostFtdcPriceType
	// 数量变化
	VolumeChange types.TThostFtdcVolumeType
	// 操作日期
	ActionDate types.TThostFtdcDateType
	// 操作时间
	ActionTime types.TThostFtdcTimeType
	// 交易所交易员代码
	TraderID types.TThostFtdcTraderIDType
	// 安装编号
	InstallID types.TThostFtdcInstallIDType
	// 本地报单编号
	OrderLocalID types.TThostFtdcOrderLocalIDType
	// 操作本地编号
	ActionLocalID types.TThostFtdcOrderLocalIDType
	// 会员代码
	ParticipantID types.TThostFtdcParticipantIDType
	// 客户代码
	ClientID types.TThostFtdcClientIDType
	// 业务单元
	BusinessUnit types.TThostFtdcBusinessUnitType
	// 报单操作状态
	OrderActionStatus types.TThostFtdcOrderActionStatusType
	// 用户代码
	UserID types.TThostFtdcUserIDType
	// 营业部编号
	BranchID types.TThostFtdcBranchIDType
	// IP地址
	IPAddress types.TThostFtdcIPAddressType
	// Mac地址
	MacAddress types.TThostFtdcMacAddressType
}

// 交易所报单操作失败
type CThostFtdcExchangeOrderActionErrorField struct {
	// 交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	// 报单编号
	OrderSysID types.TThostFtdcOrderSysIDType
	// 交易所交易员代码
	TraderID types.TThostFtdcTraderIDType
	// 安装编号
	InstallID types.TThostFtdcInstallIDType
	// 本地报单编号
	OrderLocalID types.TThostFtdcOrderLocalIDType
	// 操作本地编号
	ActionLocalID types.TThostFtdcOrderLocalIDType
	// 错误代码
	ErrorID types.TThostFtdcErrorIDType
	// 错误信息
	ErrorMsg types.TThostFtdcErrorMsgType
}

// 交易所成交
type CThostFtdcExchangeTradeField struct {
	// 交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	// 成交编号
	TradeID types.TThostFtdcTradeIDType
	// 买卖方向
	Direction types.TThostFtdcDirectionType
	// 报单编号
	OrderSysID types.TThostFtdcOrderSysIDType
	// 会员代码
	ParticipantID types.TThostFtdcParticipantIDType
	// 客户代码
	ClientID types.TThostFtdcClientIDType
	// 交易角色
	TradingRole types.TThostFtdcTradingRoleType
	// 合约在交易所的代码
	ExchangeInstID types.TThostFtdcExchangeInstIDType
	// 开平标志
	OffsetFlag types.TThostFtdcOffsetFlagType
	// 投机套保标志
	HedgeFlag types.TThostFtdcHedgeFlagType
	// 价格
	Price types.TThostFtdcPriceType
	// 数量
	Volume types.TThostFtdcVolumeType
	// 成交时期
	TradeDate types.TThostFtdcDateType
	// 成交时间
	TradeTime types.TThostFtdcTimeType
	// 成交类型
	TradeType types.TThostFtdcTradeTypeType
	// 成交价来源
	PriceSource types.TThostFtdcPriceSourceType
	// 交易所交易员代码
	TraderID types.TThostFtdcTraderIDType
	// 本地报单编号
	OrderLocalID types.TThostFtdcOrderLocalIDType
	// 结算会员编号
	ClearingPartID types.TThostFtdcParticipantIDType
	// 业务单元
	BusinessUnit types.TThostFtdcBusinessUnitType
	// 序号
	SequenceNo types.TThostFtdcSequenceNoType
	// 成交来源
	TradeSource types.TThostFtdcTradeSourceType
}

// 成交
type CThostFtdcTradeField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	// 合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
	// 报单引用
	OrderRef types.TThostFtdcOrderRefType
	// 用户代码
	UserID types.TThostFtdcUserIDType
	// 交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	// 成交编号
	TradeID types.TThostFtdcTradeIDType
	// 买卖方向
	Direction types.TThostFtdcDirectionType
	// 报单编号
	OrderSysID types.TThostFtdcOrderSysIDType
	// 会员代码
	ParticipantID types.TThostFtdcParticipantIDType
	// 客户代码
	ClientID types.TThostFtdcClientIDType
	// 交易角色
	TradingRole types.TThostFtdcTradingRoleType
	// 合约在交易所的代码
	ExchangeInstID types.TThostFtdcExchangeInstIDType
	// 开平标志
	OffsetFlag types.TThostFtdcOffsetFlagType
	// 投机套保标志
	HedgeFlag types.TThostFtdcHedgeFlagType
	// 价格
	Price types.TThostFtdcPriceType
	// 数量
	Volume types.TThostFtdcVolumeType
	// 成交时期
	TradeDate types.TThostFtdcDateType
	// 成交时间
	TradeTime types.TThostFtdcTimeType
	// 成交类型
	TradeType types.TThostFtdcTradeTypeType
	// 成交价来源
	PriceSource types.TThostFtdcPriceSourceType
	// 交易所交易员代码
	TraderID types.TThostFtdcTraderIDType
	// 本地报单编号
	OrderLocalID types.TThostFtdcOrderLocalIDType
	// 结算会员编号
	ClearingPartID types.TThostFtdcParticipantIDType
	// 业务单元
	BusinessUnit types.TThostFtdcBusinessUnitType
	// 序号
	SequenceNo types.TThostFtdcSequenceNoType
	// 交易日
	TradingDay types.TThostFtdcDateType
	// 结算编号
	SettlementID types.TThostFtdcSettlementIDType
	// 经纪公司报单编号
	BrokerOrderSeq types.TThostFtdcSequenceNoType
	// 成交来源
	TradeSource types.TThostFtdcTradeSourceType
}

// 用户会话
type CThostFtdcUserSessionField struct {
	// 前置编号
	FrontID types.TThostFtdcFrontIDType
	// 会话编号
	SessionID types.TThostFtdcSessionIDType
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 用户代码
	UserID types.TThostFtdcUserIDType
	// 登录日期
	LoginDate types.TThostFtdcDateType
	// 登录时间
	LoginTime types.TThostFtdcTimeType
	// IP地址
	IPAddress types.TThostFtdcIPAddressType
	// 用户端产品信息
	UserProductInfo types.TThostFtdcProductInfoType
	// 接口端产品信息
	InterfaceProductInfo types.TThostFtdcProductInfoType
	// 协议信息
	ProtocolInfo types.TThostFtdcProtocolInfoType
	// Mac地址
	MacAddress types.TThostFtdcMacAddressType
	// 登录备注
	LoginRemark types.TThostFtdcLoginRemarkType
}

// 查询最大报单数量
type CThostFtdcQueryMaxOrderVolumeField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	// 合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
	// 买卖方向
	Direction types.TThostFtdcDirectionType
	// 开平标志
	OffsetFlag types.TThostFtdcOffsetFlagType
	// 投机套保标志
	HedgeFlag types.TThostFtdcHedgeFlagType
	// 最大允许报单数量
	MaxVolume types.TThostFtdcVolumeType
}

// 投资者结算结果确认信息
type CThostFtdcSettlementInfoConfirmField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	// 确认日期
	ConfirmDate types.TThostFtdcDateType
	// 确认时间
	ConfirmTime types.TThostFtdcTimeType
}

// 出入金同步
type CThostFtdcSyncDepositField struct {
	// 出入金流水号
	DepositSeqNo types.TThostFtdcDepositSeqNoType
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	// 入金金额
	Deposit types.TThostFtdcMoneyType
	// 是否强制进行
	IsForce types.TThostFtdcBoolType
	// 币种代码
	CurrencyID types.TThostFtdcCurrencyIDType
}

// 货币质押同步
type CThostFtdcSyncFundMortgageField struct {
	// 货币质押流水号
	MortgageSeqNo types.TThostFtdcDepositSeqNoType
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	// 源币种
	FromCurrencyID types.TThostFtdcCurrencyIDType
	// 质押金额
	MortgageAmount types.TThostFtdcMoneyType
	// 目标币种
	ToCurrencyID types.TThostFtdcCurrencyIDType
}

// 核心间资金转移
type CThostFtdcTransFundField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 用户代码
	UserID types.TThostFtdcUserIDType
	// 投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	// 投资单元代码
	InvestUnitID types.TThostFtdcInvestUnitIDType
	// 投资者帐号
	AccountID types.TThostFtdcAccountIDType
	// 出金的核心地址
	DepositKernel types.TThostFtdcAddressAndPortType
	// 入金的核心地址
	IncomingKernel types.TThostFtdcAddressAndPortType
	// 币种代码
	CurrencyID types.TThostFtdcCurrencyIDType
	// 转账金额
	Amount types.TThostFtdcMoneyType
}

// 经纪公司同步
type CThostFtdcBrokerSyncField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
}

// 正在同步中的投资者
type CThostFtdcSyncingInvestorField struct {
	// 投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 投资者分组代码
	InvestorGroupID types.TThostFtdcInvestorIDType
	// 投资者名称
	InvestorName types.TThostFtdcPartyNameType
	// 证件类型
	IdentifiedCardType types.TThostFtdcIdCardTypeType
	// 证件号码
	IdentifiedCardNo types.TThostFtdcIdentifiedCardNoType
	// 是否活跃
	IsActive types.TThostFtdcBoolType
	// 联系电话
	Telephone types.TThostFtdcTelephoneType
	// 通讯地址
	Address types.TThostFtdcAddressType
	// 开户日期
	OpenDate types.TThostFtdcDateType
	// 手机
	Mobile types.TThostFtdcMobileType
	// 手续费率模板代码
	CommModelID types.TThostFtdcInvestorIDType
	// 保证金率模板代码
	MarginModelID types.TThostFtdcInvestorIDType
}

// 正在同步中的交易代码
type CThostFtdcSyncingTradingCodeField struct {
	// 投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	// 客户代码
	ClientID types.TThostFtdcClientIDType
	// 是否活跃
	IsActive types.TThostFtdcBoolType
	// 交易编码类型
	ClientIDType types.TThostFtdcClientIDTypeType
}

// 正在同步中的投资者分组
type CThostFtdcSyncingInvestorGroupField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 投资者分组代码
	InvestorGroupID types.TThostFtdcInvestorIDType
	// 投资者分组名称
	InvestorGroupName types.TThostFtdcInvestorGroupNameType
}

// 正在同步中的交易账号
type CThostFtdcSyncingTradingAccountField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 投资者帐号
	AccountID types.TThostFtdcAccountIDType
	// 上次质押金额
	PreMortgage types.TThostFtdcMoneyType
	// 上次信用额度
	PreCredit types.TThostFtdcMoneyType
	// 上次存款额
	PreDeposit types.TThostFtdcMoneyType
	// 上次结算准备金
	PreBalance types.TThostFtdcMoneyType
	// 上次占用的保证金
	PreMargin types.TThostFtdcMoneyType
	// 利息基数
	InterestBase types.TThostFtdcMoneyType
	// 利息收入
	Interest types.TThostFtdcMoneyType
	// 入金金额
	Deposit types.TThostFtdcMoneyType
	// 出金金额
	Withdraw types.TThostFtdcMoneyType
	// 冻结的保证金
	FrozenMargin types.TThostFtdcMoneyType
	// 冻结的资金
	FrozenCash types.TThostFtdcMoneyType
	// 冻结的手续费
	FrozenCommission types.TThostFtdcMoneyType
	// 当前保证金总额
	CurrMargin types.TThostFtdcMoneyType
	// 资金差额
	CashIn types.TThostFtdcMoneyType
	// 手续费
	Commission types.TThostFtdcMoneyType
	// 平仓盈亏
	CloseProfit types.TThostFtdcMoneyType
	// 持仓盈亏
	PositionProfit types.TThostFtdcMoneyType
	// 期货结算准备金
	Balance types.TThostFtdcMoneyType
	// 可用资金
	Available types.TThostFtdcMoneyType
	// 可取资金
	WithdrawQuota types.TThostFtdcMoneyType
	// 基本准备金
	Reserve types.TThostFtdcMoneyType
	// 交易日
	TradingDay types.TThostFtdcDateType
	// 结算编号
	SettlementID types.TThostFtdcSettlementIDType
	// 信用额度
	Credit types.TThostFtdcMoneyType
	// 质押金额
	Mortgage types.TThostFtdcMoneyType
	// 交易所保证金
	ExchangeMargin types.TThostFtdcMoneyType
	// 投资者交割保证金
	DeliveryMargin types.TThostFtdcMoneyType
	// 交易所交割保证金
	ExchangeDeliveryMargin types.TThostFtdcMoneyType
	// 保底期货结算准备金
	ReserveBalance types.TThostFtdcMoneyType
	// 币种代码
	CurrencyID types.TThostFtdcCurrencyIDType
	// 上次货币质入金额
	PreFundMortgageIn types.TThostFtdcMoneyType
	// 上次货币质出金额
	PreFundMortgageOut types.TThostFtdcMoneyType
	// 货币质入金额
	FundMortgageIn types.TThostFtdcMoneyType
	// 货币质出金额
	FundMortgageOut types.TThostFtdcMoneyType
	// 货币质押余额
	FundMortgageAvailable types.TThostFtdcMoneyType
	// 可质押货币金额
	MortgageableFund types.TThostFtdcMoneyType
	// 特殊产品占用保证金
	SpecProductMargin types.TThostFtdcMoneyType
	// 特殊产品冻结保证金
	SpecProductFrozenMargin types.TThostFtdcMoneyType
	// 特殊产品手续费
	SpecProductCommission types.TThostFtdcMoneyType
	// 特殊产品冻结手续费
	SpecProductFrozenCommission types.TThostFtdcMoneyType
	// 特殊产品持仓盈亏
	SpecProductPositionProfit types.TThostFtdcMoneyType
	// 特殊产品平仓盈亏
	SpecProductCloseProfit types.TThostFtdcMoneyType
	// 根据持仓盈亏算法计算的特殊产品持仓盈亏
	SpecProductPositionProfitByAlg types.TThostFtdcMoneyType
	// 特殊产品交易所保证金
	SpecProductExchangeMargin types.TThostFtdcMoneyType
}

// 正在同步中的投资者持仓
type CThostFtdcSyncingInvestorPositionField struct {
	// 合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	// 持仓多空方向
	PosiDirection types.TThostFtdcPosiDirectionType
	// 投机套保标志
	HedgeFlag types.TThostFtdcHedgeFlagType
	// 持仓日期
	PositionDate types.TThostFtdcPositionDateType
	// 上日持仓
	YdPosition types.TThostFtdcVolumeType
	// 今日持仓
	Position types.TThostFtdcVolumeType
	// 多头冻结
	LongFrozen types.TThostFtdcVolumeType
	// 空头冻结
	ShortFrozen types.TThostFtdcVolumeType
	// 开仓冻结金额
	LongFrozenAmount types.TThostFtdcMoneyType
	// 开仓冻结金额
	ShortFrozenAmount types.TThostFtdcMoneyType
	// 开仓量
	OpenVolume types.TThostFtdcVolumeType
	// 平仓量
	CloseVolume types.TThostFtdcVolumeType
	// 开仓金额
	OpenAmount types.TThostFtdcMoneyType
	// 平仓金额
	CloseAmount types.TThostFtdcMoneyType
	// 持仓成本
	PositionCost types.TThostFtdcMoneyType
	// 上次占用的保证金
	PreMargin types.TThostFtdcMoneyType
	// 占用的保证金
	UseMargin types.TThostFtdcMoneyType
	// 冻结的保证金
	FrozenMargin types.TThostFtdcMoneyType
	// 冻结的资金
	FrozenCash types.TThostFtdcMoneyType
	// 冻结的手续费
	FrozenCommission types.TThostFtdcMoneyType
	// 资金差额
	CashIn types.TThostFtdcMoneyType
	// 手续费
	Commission types.TThostFtdcMoneyType
	// 平仓盈亏
	CloseProfit types.TThostFtdcMoneyType
	// 持仓盈亏
	PositionProfit types.TThostFtdcMoneyType
	// 上次结算价
	PreSettlementPrice types.TThostFtdcPriceType
	// 本次结算价
	SettlementPrice types.TThostFtdcPriceType
	// 交易日
	TradingDay types.TThostFtdcDateType
	// 结算编号
	SettlementID types.TThostFtdcSettlementIDType
	// 开仓成本
	OpenCost types.TThostFtdcMoneyType
	// 交易所保证金
	ExchangeMargin types.TThostFtdcMoneyType
	// 组合成交形成的持仓
	CombPosition types.TThostFtdcVolumeType
	// 组合多头冻结
	CombLongFrozen types.TThostFtdcVolumeType
	// 组合空头冻结
	CombShortFrozen types.TThostFtdcVolumeType
	// 逐日盯市平仓盈亏
	CloseProfitByDate types.TThostFtdcMoneyType
	// 逐笔对冲平仓盈亏
	CloseProfitByTrade types.TThostFtdcMoneyType
	// 今日持仓
	TodayPosition types.TThostFtdcVolumeType
	// 保证金率
	MarginRateByMoney types.TThostFtdcRatioType
	// 保证金率(按手数)
	MarginRateByVolume types.TThostFtdcRatioType
	// 执行冻结
	StrikeFrozen types.TThostFtdcVolumeType
	// 执行冻结金额
	StrikeFrozenAmount types.TThostFtdcMoneyType
	// 放弃执行冻结
	AbandonFrozen types.TThostFtdcVolumeType
}

// 正在同步中的合约保证金率
type CThostFtdcSyncingInstrumentMarginRateField struct {
	// 合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
	// 投资者范围
	InvestorRange types.TThostFtdcInvestorRangeType
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	// 投机套保标志
	HedgeFlag types.TThostFtdcHedgeFlagType
	// 多头保证金率
	LongMarginRatioByMoney types.TThostFtdcRatioType
	// 多头保证金费
	LongMarginRatioByVolume types.TThostFtdcMoneyType
	// 空头保证金率
	ShortMarginRatioByMoney types.TThostFtdcRatioType
	// 空头保证金费
	ShortMarginRatioByVolume types.TThostFtdcMoneyType
	// 是否相对交易所收取
	IsRelative types.TThostFtdcBoolType
}

// 正在同步中的合约手续费率
type CThostFtdcSyncingInstrumentCommissionRateField struct {
	// 合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
	// 投资者范围
	InvestorRange types.TThostFtdcInvestorRangeType
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	// 开仓手续费率
	OpenRatioByMoney types.TThostFtdcRatioType
	// 开仓手续费
	OpenRatioByVolume types.TThostFtdcRatioType
	// 平仓手续费率
	CloseRatioByMoney types.TThostFtdcRatioType
	// 平仓手续费
	CloseRatioByVolume types.TThostFtdcRatioType
	// 平今手续费率
	CloseTodayRatioByMoney types.TThostFtdcRatioType
	// 平今手续费
	CloseTodayRatioByVolume types.TThostFtdcRatioType
}

// 正在同步中的合约交易权限
type CThostFtdcSyncingInstrumentTradingRightField struct {
	// 合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
	// 投资者范围
	InvestorRange types.TThostFtdcInvestorRangeType
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	// 交易权限
	TradingRight types.TThostFtdcTradingRightType
}

// 查询报单
type CThostFtdcQryOrderField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	// 合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
	// 交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	// 报单编号
	OrderSysID types.TThostFtdcOrderSysIDType
	// 开始时间
	InsertTimeStart types.TThostFtdcTimeType
	// 结束时间
	InsertTimeEnd types.TThostFtdcTimeType
}

// 查询成交
type CThostFtdcQryTradeField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	// 合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
	// 交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	// 成交编号
	TradeID types.TThostFtdcTradeIDType
	// 开始时间
	TradeTimeStart types.TThostFtdcTimeType
	// 结束时间
	TradeTimeEnd types.TThostFtdcTimeType
}

// 查询投资者持仓
type CThostFtdcQryInvestorPositionField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	// 合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
}

// 查询资金账户
type CThostFtdcQryTradingAccountField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	// 币种代码
	CurrencyID types.TThostFtdcCurrencyIDType
}

// 查询投资者
type CThostFtdcQryInvestorField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID types.TThostFtdcInvestorIDType
}

// 查询交易编码
type CThostFtdcQryTradingCodeField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	// 交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	// 客户代码
	ClientID types.TThostFtdcClientIDType
	// 交易编码类型
	ClientIDType types.TThostFtdcClientIDTypeType
}

// 查询投资者组
type CThostFtdcQryInvestorGroupField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
}

// 查询合约保证金率
type CThostFtdcQryInstrumentMarginRateField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	// 合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
	// 投机套保标志
	HedgeFlag types.TThostFtdcHedgeFlagType
}

// 查询手续费率
type CThostFtdcQryInstrumentCommissionRateField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	// 合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
}

// 查询合约交易权限
type CThostFtdcQryInstrumentTradingRightField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	// 合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
}

// 查询经纪公司
type CThostFtdcQryBrokerField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
}

// 查询交易员
type CThostFtdcQryTraderField struct {
	// 交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	// 会员代码
	ParticipantID types.TThostFtdcParticipantIDType
	// 交易所交易员代码
	TraderID types.TThostFtdcTraderIDType
}

// 查询管理用户功能权限
type CThostFtdcQrySuperUserFunctionField struct {
	// 用户代码
	UserID types.TThostFtdcUserIDType
}

// 查询用户会话
type CThostFtdcQryUserSessionField struct {
	// 前置编号
	FrontID types.TThostFtdcFrontIDType
	// 会话编号
	SessionID types.TThostFtdcSessionIDType
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 用户代码
	UserID types.TThostFtdcUserIDType
}

// 查询经纪公司会员代码
type CThostFtdcQryPartBrokerField struct {
	// 交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 会员代码
	ParticipantID types.TThostFtdcParticipantIDType
}

// 查询前置状态
type CThostFtdcQryFrontStatusField struct {
	// 前置编号
	FrontID types.TThostFtdcFrontIDType
}

// 查询交易所报单
type CThostFtdcQryExchangeOrderField struct {
	// 会员代码
	ParticipantID types.TThostFtdcParticipantIDType
	// 客户代码
	ClientID types.TThostFtdcClientIDType
	// 合约在交易所的代码
	ExchangeInstID types.TThostFtdcExchangeInstIDType
	// 交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	// 交易所交易员代码
	TraderID types.TThostFtdcTraderIDType
}

// 查询报单操作
type CThostFtdcQryOrderActionField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	// 交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
}

// 查询交易所报单操作
type CThostFtdcQryExchangeOrderActionField struct {
	// 会员代码
	ParticipantID types.TThostFtdcParticipantIDType
	// 客户代码
	ClientID types.TThostFtdcClientIDType
	// 交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	// 交易所交易员代码
	TraderID types.TThostFtdcTraderIDType
}

// 查询管理用户
type CThostFtdcQrySuperUserField struct {
	// 用户代码
	UserID types.TThostFtdcUserIDType
}

// 查询交易所
type CThostFtdcQryExchangeField struct {
	// 交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
}

// 查询产品
type CThostFtdcQryProductField struct {
	// 产品代码
	ProductID types.TThostFtdcInstrumentIDType
	// 产品类型
	ProductClass types.TThostFtdcProductClassType
}

// 查询合约
type CThostFtdcQryInstrumentField struct {
	// 合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
	// 交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	// 合约在交易所的代码
	ExchangeInstID types.TThostFtdcExchangeInstIDType
	// 产品代码
	ProductID types.TThostFtdcInstrumentIDType
}

// 查询申请组合合约
type CThostFtdcQryCombInstrumentField struct {
	// 交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	// 产品代码
	ProductID types.TThostFtdcInstrumentIDType
}

// 查询投资者RCAMS组合保证金
type CThostFtdcQryRCAMSInvestorProdMarginField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	// 产品代码
	CombProductID types.TThostFtdcInstrumentIDType
}

// 查询RCAMS策略组合持仓
type CThostFtdcQryRCAMSInvestorCombPositionField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	// 单腿合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
}

// 查询行情
type CThostFtdcQryDepthMarketDataField struct {
	// 合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
}

// 查询经纪公司用户
type CThostFtdcQryBrokerUserField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 用户代码
	UserID types.TThostFtdcUserIDType
}

// 查询经纪公司用户
type CThostFtdcQryControlParamField struct {
	// 投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	// 开关代码
	ControlParamID types.TThostFtdcControlParamIDType
}

// 查询经纪公司用户权限
type CThostFtdcQryBrokerUserFunctionField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 用户代码
	UserID types.TThostFtdcUserIDType
}

// 查询交易员报盘机
type CThostFtdcQryTraderOfferField struct {
	// 交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	// 交易所交易员代码
	TraderID types.TThostFtdcTraderIDType
}

// 查询出入金流水
type CThostFtdcQrySyncDepositField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 出入金流水号
	DepositSeqNo types.TThostFtdcDepositSeqNoType
}

// 查询投资者结算结果
type CThostFtdcQrySettlementInfoField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	// 交易日
	TradingDay types.TThostFtdcDateType
}

// 查询交易所保证金率
type CThostFtdcQryExchangeMarginRateField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
	// 投机套保标志
	HedgeFlag types.TThostFtdcHedgeFlagType
}

// 查询交易所调整保证金率
type CThostFtdcQryExchangeMarginRateAdjustField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
	// 投机套保标志
	HedgeFlag types.TThostFtdcHedgeFlagType
}

// 查询汇率
type CThostFtdcQryExchangeRateField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 源币种
	FromCurrencyID types.TThostFtdcCurrencyIDType
	// 目标币种
	ToCurrencyID types.TThostFtdcCurrencyIDType
}

// 查询货币质押流水
type CThostFtdcQrySyncFundMortgageField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 货币质押流水号
	MortgageSeqNo types.TThostFtdcDepositSeqNoType
}

// 查询报单
type CThostFtdcQryHisOrderField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	// 合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
	// 交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	// 报单编号
	OrderSysID types.TThostFtdcOrderSysIDType
	// 开始时间
	InsertTimeStart types.TThostFtdcTimeType
	// 结束时间
	InsertTimeEnd types.TThostFtdcTimeType
	// 交易日
	TradingDay types.TThostFtdcDateType
	// 结算编号
	SettlementID types.TThostFtdcSettlementIDType
}

// 当前期权合约最小保证金
type CThostFtdcOptionInstrMiniMarginField struct {
	// 合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
	// 投资者范围
	InvestorRange types.TThostFtdcInvestorRangeType
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	// 单位（手）期权合约最小保证金
	MinMargin types.TThostFtdcMoneyType
	// 取值方式
	ValueMethod types.TThostFtdcValueMethodType
	// 是否跟随交易所收取
	IsRelative types.TThostFtdcBoolType
}

// 当前期权合约保证金调整系数
type CThostFtdcOptionInstrMarginAdjustField struct {
	// 合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
	// 投资者范围
	InvestorRange types.TThostFtdcInvestorRangeType
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	// 投机空头保证金调整系数
	SShortMarginRatioByMoney types.TThostFtdcRatioType
	// 投机空头保证金调整系数
	SShortMarginRatioByVolume types.TThostFtdcMoneyType
	// 保值空头保证金调整系数
	HShortMarginRatioByMoney types.TThostFtdcRatioType
	// 保值空头保证金调整系数
	HShortMarginRatioByVolume types.TThostFtdcMoneyType
	// 套利空头保证金调整系数
	AShortMarginRatioByMoney types.TThostFtdcRatioType
	// 套利空头保证金调整系数
	AShortMarginRatioByVolume types.TThostFtdcMoneyType
	// 是否跟随交易所收取
	IsRelative types.TThostFtdcBoolType
	// 做市商空头保证金调整系数
	MShortMarginRatioByMoney types.TThostFtdcRatioType
	// 做市商空头保证金调整系数
	MShortMarginRatioByVolume types.TThostFtdcMoneyType
}

// 当前期权合约手续费的详细内容
type CThostFtdcOptionInstrCommRateField struct {
	// 合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
	// 投资者范围
	InvestorRange types.TThostFtdcInvestorRangeType
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	// 开仓手续费率
	OpenRatioByMoney types.TThostFtdcRatioType
	// 开仓手续费
	OpenRatioByVolume types.TThostFtdcRatioType
	// 平仓手续费率
	CloseRatioByMoney types.TThostFtdcRatioType
	// 平仓手续费
	CloseRatioByVolume types.TThostFtdcRatioType
	// 平今手续费率
	CloseTodayRatioByMoney types.TThostFtdcRatioType
	// 平今手续费
	CloseTodayRatioByVolume types.TThostFtdcRatioType
	// 执行手续费率
	StrikeRatioByMoney types.TThostFtdcRatioType
	// 执行手续费
	StrikeRatioByVolume types.TThostFtdcRatioType
}

// 期权交易成本
type CThostFtdcOptionInstrTradeCostField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	// 合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
	// 投机套保标志
	HedgeFlag types.TThostFtdcHedgeFlagType
	// 期权合约保证金不变部分
	FixedMargin types.TThostFtdcMoneyType
	// 期权合约最小保证金
	MiniMargin types.TThostFtdcMoneyType
	// 期权合约权利金
	Royalty types.TThostFtdcMoneyType
	// 交易所期权合约保证金不变部分
	ExchFixedMargin types.TThostFtdcMoneyType
	// 交易所期权合约最小保证金
	ExchMiniMargin types.TThostFtdcMoneyType
}

// 期权交易成本查询
type CThostFtdcQryOptionInstrTradeCostField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	// 合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
	// 投机套保标志
	HedgeFlag types.TThostFtdcHedgeFlagType
	// 期权合约报价
	InputPrice types.TThostFtdcPriceType
	// 标的价格,填0则用昨结算价
	UnderlyingPrice types.TThostFtdcPriceType
}

// 期权手续费率查询
type CThostFtdcQryOptionInstrCommRateField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	// 合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
}

// 股指现货指数
type CThostFtdcIndexPriceField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
	// 指数现货收盘价
	ClosePrice types.TThostFtdcPriceType
}

// 输入的执行宣告
type CThostFtdcInputExecOrderField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	// 合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
	// 执行宣告引用
	ExecOrderRef types.TThostFtdcOrderRefType
	// 用户代码
	UserID types.TThostFtdcUserIDType
	// 数量
	Volume types.TThostFtdcVolumeType
	// 请求编号
	RequestID types.TThostFtdcRequestIDType
	// 业务单元
	BusinessUnit types.TThostFtdcBusinessUnitType
	// 开平标志
	OffsetFlag types.TThostFtdcOffsetFlagType
	// 投机套保标志
	HedgeFlag types.TThostFtdcHedgeFlagType
	// 执行类型
	ActionType types.TThostFtdcActionTypeType
	// 保留头寸申请的持仓方向
	PosiDirection types.TThostFtdcPosiDirectionType
	// 期权行权后是否保留期货头寸的标记,该字段已废弃
	ReservePositionFlag types.TThostFtdcExecOrderPositionFlagType
	// 期权行权后生成的头寸是否自动平仓
	CloseFlag types.TThostFtdcExecOrderCloseFlagType
	// 交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	// 投资单元代码
	InvestUnitID types.TThostFtdcInvestUnitIDType
	// 资金账号
	AccountID types.TThostFtdcAccountIDType
	// 币种代码
	CurrencyID types.TThostFtdcCurrencyIDType
	// 交易编码
	ClientID types.TThostFtdcClientIDType
	// IP地址
	IPAddress types.TThostFtdcIPAddressType
	// Mac地址
	MacAddress types.TThostFtdcMacAddressType
}

// 输入执行宣告操作
type CThostFtdcInputExecOrderActionField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	// 执行宣告操作引用
	ExecOrderActionRef types.TThostFtdcOrderActionRefType
	// 执行宣告引用
	ExecOrderRef types.TThostFtdcOrderRefType
	// 请求编号
	RequestID types.TThostFtdcRequestIDType
	// 前置编号
	FrontID types.TThostFtdcFrontIDType
	// 会话编号
	SessionID types.TThostFtdcSessionIDType
	// 交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	// 执行宣告操作编号
	ExecOrderSysID types.TThostFtdcExecOrderSysIDType
	// 操作标志
	ActionFlag types.TThostFtdcActionFlagType
	// 用户代码
	UserID types.TThostFtdcUserIDType
	// 合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
	// 投资单元代码
	InvestUnitID types.TThostFtdcInvestUnitIDType
	// IP地址
	IPAddress types.TThostFtdcIPAddressType
	// Mac地址
	MacAddress types.TThostFtdcMacAddressType
}

// 执行宣告
type CThostFtdcExecOrderField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	// 合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
	// 执行宣告引用
	ExecOrderRef types.TThostFtdcOrderRefType
	// 用户代码
	UserID types.TThostFtdcUserIDType
	// 数量
	Volume types.TThostFtdcVolumeType
	// 请求编号
	RequestID types.TThostFtdcRequestIDType
	// 业务单元
	BusinessUnit types.TThostFtdcBusinessUnitType
	// 开平标志
	OffsetFlag types.TThostFtdcOffsetFlagType
	// 投机套保标志
	HedgeFlag types.TThostFtdcHedgeFlagType
	// 执行类型
	ActionType types.TThostFtdcActionTypeType
	// 保留头寸申请的持仓方向
	PosiDirection types.TThostFtdcPosiDirectionType
	// 期权行权后是否保留期货头寸的标记
	ReservePositionFlag types.TThostFtdcExecOrderPositionFlagType
	// 期权行权后生成的头寸是否自动平仓
	CloseFlag types.TThostFtdcExecOrderCloseFlagType
	// 本地执行宣告编号
	ExecOrderLocalID types.TThostFtdcOrderLocalIDType
	// 交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	// 会员代码
	ParticipantID types.TThostFtdcParticipantIDType
	// 客户代码
	ClientID types.TThostFtdcClientIDType
	// 合约在交易所的代码
	ExchangeInstID types.TThostFtdcExchangeInstIDType
	// 交易所交易员代码
	TraderID types.TThostFtdcTraderIDType
	// 安装编号
	InstallID types.TThostFtdcInstallIDType
	// 执行宣告提交状态
	OrderSubmitStatus types.TThostFtdcOrderSubmitStatusType
	// 报单提示序号
	NotifySequence types.TThostFtdcSequenceNoType
	// 交易日
	TradingDay types.TThostFtdcDateType
	// 结算编号
	SettlementID types.TThostFtdcSettlementIDType
	// 执行宣告编号
	ExecOrderSysID types.TThostFtdcExecOrderSysIDType
	// 报单日期
	InsertDate types.TThostFtdcDateType
	// 插入时间
	InsertTime types.TThostFtdcTimeType
	// 撤销时间
	CancelTime types.TThostFtdcTimeType
	// 执行结果
	ExecResult types.TThostFtdcExecResultType
	// 结算会员编号
	ClearingPartID types.TThostFtdcParticipantIDType
	// 序号
	SequenceNo types.TThostFtdcSequenceNoType
	// 前置编号
	FrontID types.TThostFtdcFrontIDType
	// 会话编号
	SessionID types.TThostFtdcSessionIDType
	// 用户端产品信息
	UserProductInfo types.TThostFtdcProductInfoType
	// 状态信息
	StatusMsg types.TThostFtdcErrorMsgType
	// 操作用户代码
	ActiveUserID types.TThostFtdcUserIDType
	// 经纪公司报单编号
	BrokerExecOrderSeq types.TThostFtdcSequenceNoType
	// 营业部编号
	BranchID types.TThostFtdcBranchIDType
	// 投资单元代码
	InvestUnitID types.TThostFtdcInvestUnitIDType
	// 资金账号
	AccountID types.TThostFtdcAccountIDType
	// 币种代码
	CurrencyID types.TThostFtdcCurrencyIDType
	// IP地址
	IPAddress types.TThostFtdcIPAddressType
	// Mac地址
	MacAddress types.TThostFtdcMacAddressType
}

// 执行宣告操作
type CThostFtdcExecOrderActionField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	// 执行宣告操作引用
	ExecOrderActionRef types.TThostFtdcOrderActionRefType
	// 执行宣告引用
	ExecOrderRef types.TThostFtdcOrderRefType
	// 请求编号
	RequestID types.TThostFtdcRequestIDType
	// 前置编号
	FrontID types.TThostFtdcFrontIDType
	// 会话编号
	SessionID types.TThostFtdcSessionIDType
	// 交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	// 执行宣告操作编号
	ExecOrderSysID types.TThostFtdcExecOrderSysIDType
	// 操作标志
	ActionFlag types.TThostFtdcActionFlagType
	// 操作日期
	ActionDate types.TThostFtdcDateType
	// 操作时间
	ActionTime types.TThostFtdcTimeType
	// 交易所交易员代码
	TraderID types.TThostFtdcTraderIDType
	// 安装编号
	InstallID types.TThostFtdcInstallIDType
	// 本地执行宣告编号
	ExecOrderLocalID types.TThostFtdcOrderLocalIDType
	// 操作本地编号
	ActionLocalID types.TThostFtdcOrderLocalIDType
	// 会员代码
	ParticipantID types.TThostFtdcParticipantIDType
	// 客户代码
	ClientID types.TThostFtdcClientIDType
	// 业务单元
	BusinessUnit types.TThostFtdcBusinessUnitType
	// 报单操作状态
	OrderActionStatus types.TThostFtdcOrderActionStatusType
	// 用户代码
	UserID types.TThostFtdcUserIDType
	// 执行类型
	ActionType types.TThostFtdcActionTypeType
	// 状态信息
	StatusMsg types.TThostFtdcErrorMsgType
	// 合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
	// 营业部编号
	BranchID types.TThostFtdcBranchIDType
	// 投资单元代码
	InvestUnitID types.TThostFtdcInvestUnitIDType
	// IP地址
	IPAddress types.TThostFtdcIPAddressType
	// Mac地址
	MacAddress types.TThostFtdcMacAddressType
}

// 执行宣告查询
type CThostFtdcQryExecOrderField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	// 合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
	// 交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	// 执行宣告编号
	ExecOrderSysID types.TThostFtdcExecOrderSysIDType
	// 开始时间
	InsertTimeStart types.TThostFtdcTimeType
	// 结束时间
	InsertTimeEnd types.TThostFtdcTimeType
}

// 交易所执行宣告信息
type CThostFtdcExchangeExecOrderField struct {
	// 数量
	Volume types.TThostFtdcVolumeType
	// 请求编号
	RequestID types.TThostFtdcRequestIDType
	// 业务单元
	BusinessUnit types.TThostFtdcBusinessUnitType
	// 开平标志
	OffsetFlag types.TThostFtdcOffsetFlagType
	// 投机套保标志
	HedgeFlag types.TThostFtdcHedgeFlagType
	// 执行类型
	ActionType types.TThostFtdcActionTypeType
	// 保留头寸申请的持仓方向
	PosiDirection types.TThostFtdcPosiDirectionType
	// 期权行权后是否保留期货头寸的标记
	ReservePositionFlag types.TThostFtdcExecOrderPositionFlagType
	// 期权行权后生成的头寸是否自动平仓
	CloseFlag types.TThostFtdcExecOrderCloseFlagType
	// 本地执行宣告编号
	ExecOrderLocalID types.TThostFtdcOrderLocalIDType
	// 交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	// 会员代码
	ParticipantID types.TThostFtdcParticipantIDType
	// 客户代码
	ClientID types.TThostFtdcClientIDType
	// 合约在交易所的代码
	ExchangeInstID types.TThostFtdcExchangeInstIDType
	// 交易所交易员代码
	TraderID types.TThostFtdcTraderIDType
	// 安装编号
	InstallID types.TThostFtdcInstallIDType
	// 执行宣告提交状态
	OrderSubmitStatus types.TThostFtdcOrderSubmitStatusType
	// 报单提示序号
	NotifySequence types.TThostFtdcSequenceNoType
	// 交易日
	TradingDay types.TThostFtdcDateType
	// 结算编号
	SettlementID types.TThostFtdcSettlementIDType
	// 执行宣告编号
	ExecOrderSysID types.TThostFtdcExecOrderSysIDType
	// 报单日期
	InsertDate types.TThostFtdcDateType
	// 插入时间
	InsertTime types.TThostFtdcTimeType
	// 撤销时间
	CancelTime types.TThostFtdcTimeType
	// 执行结果
	ExecResult types.TThostFtdcExecResultType
	// 结算会员编号
	ClearingPartID types.TThostFtdcParticipantIDType
	// 序号
	SequenceNo types.TThostFtdcSequenceNoType
	// 营业部编号
	BranchID types.TThostFtdcBranchIDType
	// IP地址
	IPAddress types.TThostFtdcIPAddressType
	// Mac地址
	MacAddress types.TThostFtdcMacAddressType
}

// 交易所执行宣告查询
type CThostFtdcQryExchangeExecOrderField struct {
	// 会员代码
	ParticipantID types.TThostFtdcParticipantIDType
	// 客户代码
	ClientID types.TThostFtdcClientIDType
	// 合约在交易所的代码
	ExchangeInstID types.TThostFtdcExchangeInstIDType
	// 交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	// 交易所交易员代码
	TraderID types.TThostFtdcTraderIDType
}

// 执行宣告操作查询
type CThostFtdcQryExecOrderActionField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	// 交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
}

// 交易所执行宣告操作
type CThostFtdcExchangeExecOrderActionField struct {
	// 交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	// 执行宣告操作编号
	ExecOrderSysID types.TThostFtdcExecOrderSysIDType
	// 操作标志
	ActionFlag types.TThostFtdcActionFlagType
	// 操作日期
	ActionDate types.TThostFtdcDateType
	// 操作时间
	ActionTime types.TThostFtdcTimeType
	// 交易所交易员代码
	TraderID types.TThostFtdcTraderIDType
	// 安装编号
	InstallID types.TThostFtdcInstallIDType
	// 本地执行宣告编号
	ExecOrderLocalID types.TThostFtdcOrderLocalIDType
	// 操作本地编号
	ActionLocalID types.TThostFtdcOrderLocalIDType
	// 会员代码
	ParticipantID types.TThostFtdcParticipantIDType
	// 客户代码
	ClientID types.TThostFtdcClientIDType
	// 业务单元
	BusinessUnit types.TThostFtdcBusinessUnitType
	// 报单操作状态
	OrderActionStatus types.TThostFtdcOrderActionStatusType
	// 用户代码
	UserID types.TThostFtdcUserIDType
	// 执行类型
	ActionType types.TThostFtdcActionTypeType
	// 营业部编号
	BranchID types.TThostFtdcBranchIDType
	// IP地址
	IPAddress types.TThostFtdcIPAddressType
	// Mac地址
	MacAddress types.TThostFtdcMacAddressType
}

// 交易所执行宣告操作查询
type CThostFtdcQryExchangeExecOrderActionField struct {
	// 会员代码
	ParticipantID types.TThostFtdcParticipantIDType
	// 客户代码
	ClientID types.TThostFtdcClientIDType
	// 交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	// 交易所交易员代码
	TraderID types.TThostFtdcTraderIDType
}

// 错误执行宣告
type CThostFtdcErrExecOrderField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	// 合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
	// 执行宣告引用
	ExecOrderRef types.TThostFtdcOrderRefType
	// 用户代码
	UserID types.TThostFtdcUserIDType
	// 数量
	Volume types.TThostFtdcVolumeType
	// 请求编号
	RequestID types.TThostFtdcRequestIDType
	// 业务单元
	BusinessUnit types.TThostFtdcBusinessUnitType
	// 开平标志
	OffsetFlag types.TThostFtdcOffsetFlagType
	// 投机套保标志
	HedgeFlag types.TThostFtdcHedgeFlagType
	// 执行类型
	ActionType types.TThostFtdcActionTypeType
	// 保留头寸申请的持仓方向
	PosiDirection types.TThostFtdcPosiDirectionType
	// 期权行权后是否保留期货头寸的标记
	ReservePositionFlag types.TThostFtdcExecOrderPositionFlagType
	// 期权行权后生成的头寸是否自动平仓
	CloseFlag types.TThostFtdcExecOrderCloseFlagType
	// 交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	// 投资单元代码
	InvestUnitID types.TThostFtdcInvestUnitIDType
	// 资金账号
	AccountID types.TThostFtdcAccountIDType
	// 币种代码
	CurrencyID types.TThostFtdcCurrencyIDType
	// 交易编码
	ClientID types.TThostFtdcClientIDType
	// IP地址
	IPAddress types.TThostFtdcIPAddressType
	// Mac地址
	MacAddress types.TThostFtdcMacAddressType
	// 错误代码
	ErrorID types.TThostFtdcErrorIDType
	// 错误信息
	ErrorMsg types.TThostFtdcErrorMsgType
}

// 查询错误执行宣告
type CThostFtdcQryErrExecOrderField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID types.TThostFtdcInvestorIDType
}

// 错误执行宣告操作
type CThostFtdcErrExecOrderActionField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	// 执行宣告操作引用
	ExecOrderActionRef types.TThostFtdcOrderActionRefType
	// 执行宣告引用
	ExecOrderRef types.TThostFtdcOrderRefType
	// 请求编号
	RequestID types.TThostFtdcRequestIDType
	// 前置编号
	FrontID types.TThostFtdcFrontIDType
	// 会话编号
	SessionID types.TThostFtdcSessionIDType
	// 交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	// 执行宣告操作编号
	ExecOrderSysID types.TThostFtdcExecOrderSysIDType
	// 操作标志
	ActionFlag types.TThostFtdcActionFlagType
	// 用户代码
	UserID types.TThostFtdcUserIDType
	// 合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
	// 投资单元代码
	InvestUnitID types.TThostFtdcInvestUnitIDType
	// IP地址
	IPAddress types.TThostFtdcIPAddressType
	// Mac地址
	MacAddress types.TThostFtdcMacAddressType
	// 错误代码
	ErrorID types.TThostFtdcErrorIDType
	// 错误信息
	ErrorMsg types.TThostFtdcErrorMsgType
}

// 查询错误执行宣告操作
type CThostFtdcQryErrExecOrderActionField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID types.TThostFtdcInvestorIDType
}

// 投资者期权合约交易权限
type CThostFtdcOptionInstrTradingRightField struct {
	// 合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
	// 投资者范围
	InvestorRange types.TThostFtdcInvestorRangeType
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	// 买卖方向
	Direction types.TThostFtdcDirectionType
	// 交易权限
	TradingRight types.TThostFtdcTradingRightType
}

// 查询期权合约交易权限
type CThostFtdcQryOptionInstrTradingRightField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	// 合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
	// 买卖方向
	Direction types.TThostFtdcDirectionType
}

// 输入的询价
type CThostFtdcInputForQuoteField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	// 合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
	// 询价引用
	ForQuoteRef types.TThostFtdcOrderRefType
	// 用户代码
	UserID types.TThostFtdcUserIDType
	// 交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	// 投资单元代码
	InvestUnitID types.TThostFtdcInvestUnitIDType
	// IP地址
	IPAddress types.TThostFtdcIPAddressType
	// Mac地址
	MacAddress types.TThostFtdcMacAddressType
	// 业务单元
	BusinessUnit types.TThostFtdcBusinessUnitType
}

// 询价
type CThostFtdcForQuoteField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	// 合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
	// 询价引用
	ForQuoteRef types.TThostFtdcOrderRefType
	// 用户代码
	UserID types.TThostFtdcUserIDType
	// 本地询价编号
	ForQuoteLocalID types.TThostFtdcOrderLocalIDType
	// 交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	// 会员代码
	ParticipantID types.TThostFtdcParticipantIDType
	// 客户代码
	ClientID types.TThostFtdcClientIDType
	// 合约在交易所的代码
	ExchangeInstID types.TThostFtdcExchangeInstIDType
	// 交易所交易员代码
	TraderID types.TThostFtdcTraderIDType
	// 安装编号
	InstallID types.TThostFtdcInstallIDType
	// 报单日期
	InsertDate types.TThostFtdcDateType
	// 插入时间
	InsertTime types.TThostFtdcTimeType
	// 询价状态
	ForQuoteStatus types.TThostFtdcForQuoteStatusType
	// 前置编号
	FrontID types.TThostFtdcFrontIDType
	// 会话编号
	SessionID types.TThostFtdcSessionIDType
	// 状态信息
	StatusMsg types.TThostFtdcErrorMsgType
	// 操作用户代码
	ActiveUserID types.TThostFtdcUserIDType
	// 经纪公司询价编号
	BrokerForQutoSeq types.TThostFtdcSequenceNoType
	// 投资单元代码
	InvestUnitID types.TThostFtdcInvestUnitIDType
	// IP地址
	IPAddress types.TThostFtdcIPAddressType
	// Mac地址
	MacAddress types.TThostFtdcMacAddressType
	// 业务单元
	BusinessUnit types.TThostFtdcBusinessUnitType
}

// 询价查询
type CThostFtdcQryForQuoteField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	// 合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
	// 交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	// 开始时间
	InsertTimeStart types.TThostFtdcTimeType
	// 结束时间
	InsertTimeEnd types.TThostFtdcTimeType
}

// 交易所询价信息
type CThostFtdcExchangeForQuoteField struct {
	// 本地询价编号
	ForQuoteLocalID types.TThostFtdcOrderLocalIDType
	// 交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	// 会员代码
	ParticipantID types.TThostFtdcParticipantIDType
	// 客户代码
	ClientID types.TThostFtdcClientIDType
	// 合约在交易所的代码
	ExchangeInstID types.TThostFtdcExchangeInstIDType
	// 交易所交易员代码
	TraderID types.TThostFtdcTraderIDType
	// 安装编号
	InstallID types.TThostFtdcInstallIDType
	// 报单日期
	InsertDate types.TThostFtdcDateType
	// 插入时间
	InsertTime types.TThostFtdcTimeType
	// 询价状态
	ForQuoteStatus types.TThostFtdcForQuoteStatusType
	// IP地址
	IPAddress types.TThostFtdcIPAddressType
	// Mac地址
	MacAddress types.TThostFtdcMacAddressType
}

// 交易所询价查询
type CThostFtdcQryExchangeForQuoteField struct {
	// 会员代码
	ParticipantID types.TThostFtdcParticipantIDType
	// 客户代码
	ClientID types.TThostFtdcClientIDType
	// 合约在交易所的代码
	ExchangeInstID types.TThostFtdcExchangeInstIDType
	// 交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	// 交易所交易员代码
	TraderID types.TThostFtdcTraderIDType
}

// 输入的报价
type CThostFtdcInputQuoteField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	// 合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
	// 报价引用
	QuoteRef types.TThostFtdcOrderRefType
	// 用户代码
	UserID types.TThostFtdcUserIDType
	// 卖价格
	AskPrice types.TThostFtdcPriceType
	// 买价格
	BidPrice types.TThostFtdcPriceType
	// 卖数量
	AskVolume types.TThostFtdcVolumeType
	// 买数量
	BidVolume types.TThostFtdcVolumeType
	// 请求编号
	RequestID types.TThostFtdcRequestIDType
	// 业务单元
	BusinessUnit types.TThostFtdcBusinessUnitType
	// 卖开平标志
	AskOffsetFlag types.TThostFtdcOffsetFlagType
	// 买开平标志
	BidOffsetFlag types.TThostFtdcOffsetFlagType
	// 卖投机套保标志
	AskHedgeFlag types.TThostFtdcHedgeFlagType
	// 买投机套保标志
	BidHedgeFlag types.TThostFtdcHedgeFlagType
	// 衍生卖报单引用
	AskOrderRef types.TThostFtdcOrderRefType
	// 衍生买报单引用
	BidOrderRef types.TThostFtdcOrderRefType
	// 应价编号
	ForQuoteSysID types.TThostFtdcOrderSysIDType
	// 交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	// 投资单元代码
	InvestUnitID types.TThostFtdcInvestUnitIDType
	// 交易编码
	ClientID types.TThostFtdcClientIDType
	// IP地址
	IPAddress types.TThostFtdcIPAddressType
	// Mac地址
	MacAddress types.TThostFtdcMacAddressType
	// 顶单编号
	ReplaceSysID types.TThostFtdcOrderSysIDType
	// 自定义字段
	CustomQuoteRef types.TThostFtdcCustomOrderRefType
	// 有效期类型
	TimeCondition types.TThostFtdcTimeConditionType
}

// 输入报价操作
type CThostFtdcInputQuoteActionField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	// 报价操作引用
	QuoteActionRef types.TThostFtdcOrderActionRefType
	// 报价引用
	QuoteRef types.TThostFtdcOrderRefType
	// 请求编号
	RequestID types.TThostFtdcRequestIDType
	// 前置编号
	FrontID types.TThostFtdcFrontIDType
	// 会话编号
	SessionID types.TThostFtdcSessionIDType
	// 交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	// 报价操作编号
	QuoteSysID types.TThostFtdcOrderSysIDType
	// 操作标志
	ActionFlag types.TThostFtdcActionFlagType
	// 用户代码
	UserID types.TThostFtdcUserIDType
	// 合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
	// 投资单元代码
	InvestUnitID types.TThostFtdcInvestUnitIDType
	// 交易编码
	ClientID types.TThostFtdcClientIDType
	// IP地址
	IPAddress types.TThostFtdcIPAddressType
	// Mac地址
	MacAddress types.TThostFtdcMacAddressType
}

// 报价
type CThostFtdcQuoteField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	// 合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
	// 报价引用
	QuoteRef types.TThostFtdcOrderRefType
	// 用户代码
	UserID types.TThostFtdcUserIDType
	// 卖价格
	AskPrice types.TThostFtdcPriceType
	// 买价格
	BidPrice types.TThostFtdcPriceType
	// 卖数量
	AskVolume types.TThostFtdcVolumeType
	// 买数量
	BidVolume types.TThostFtdcVolumeType
	// 请求编号
	RequestID types.TThostFtdcRequestIDType
	// 业务单元
	BusinessUnit types.TThostFtdcBusinessUnitType
	// 卖开平标志
	AskOffsetFlag types.TThostFtdcOffsetFlagType
	// 买开平标志
	BidOffsetFlag types.TThostFtdcOffsetFlagType
	// 卖投机套保标志
	AskHedgeFlag types.TThostFtdcHedgeFlagType
	// 买投机套保标志
	BidHedgeFlag types.TThostFtdcHedgeFlagType
	// 本地报价编号
	QuoteLocalID types.TThostFtdcOrderLocalIDType
	// 交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	// 会员代码
	ParticipantID types.TThostFtdcParticipantIDType
	// 客户代码
	ClientID types.TThostFtdcClientIDType
	// 合约在交易所的代码
	ExchangeInstID types.TThostFtdcExchangeInstIDType
	// 交易所交易员代码
	TraderID types.TThostFtdcTraderIDType
	// 安装编号
	InstallID types.TThostFtdcInstallIDType
	// 报价提示序号
	NotifySequence types.TThostFtdcSequenceNoType
	// 报价提交状态
	OrderSubmitStatus types.TThostFtdcOrderSubmitStatusType
	// 交易日
	TradingDay types.TThostFtdcDateType
	// 结算编号
	SettlementID types.TThostFtdcSettlementIDType
	// 报价编号
	QuoteSysID types.TThostFtdcOrderSysIDType
	// 报单日期
	InsertDate types.TThostFtdcDateType
	// 插入时间
	InsertTime types.TThostFtdcTimeType
	// 撤销时间
	CancelTime types.TThostFtdcTimeType
	// 报价状态
	QuoteStatus types.TThostFtdcOrderStatusType
	// 结算会员编号
	ClearingPartID types.TThostFtdcParticipantIDType
	// 序号
	SequenceNo types.TThostFtdcSequenceNoType
	// 卖方报单编号
	AskOrderSysID types.TThostFtdcOrderSysIDType
	// 买方报单编号
	BidOrderSysID types.TThostFtdcOrderSysIDType
	// 前置编号
	FrontID types.TThostFtdcFrontIDType
	// 会话编号
	SessionID types.TThostFtdcSessionIDType
	// 用户端产品信息
	UserProductInfo types.TThostFtdcProductInfoType
	// 状态信息
	StatusMsg types.TThostFtdcErrorMsgType
	// 操作用户代码
	ActiveUserID types.TThostFtdcUserIDType
	// 经纪公司报价编号
	BrokerQuoteSeq types.TThostFtdcSequenceNoType
	// 衍生卖报单引用
	AskOrderRef types.TThostFtdcOrderRefType
	// 衍生买报单引用
	BidOrderRef types.TThostFtdcOrderRefType
	// 应价编号
	ForQuoteSysID types.TThostFtdcOrderSysIDType
	// 营业部编号
	BranchID types.TThostFtdcBranchIDType
	// 投资单元代码
	InvestUnitID types.TThostFtdcInvestUnitIDType
	// 资金账号
	AccountID types.TThostFtdcAccountIDType
	// 币种代码
	CurrencyID types.TThostFtdcCurrencyIDType
	// IP地址
	IPAddress types.TThostFtdcIPAddressType
	// Mac地址
	MacAddress types.TThostFtdcMacAddressType
	// 顶单编号
	ReplaceSysID types.TThostFtdcOrderSysIDType
	// 自定义字段
	CustomQuoteRef types.TThostFtdcCustomOrderRefType
	// 有效期类型
	TimeCondition types.TThostFtdcTimeConditionType
}

// 报价操作
type CThostFtdcQuoteActionField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	// 报价操作引用
	QuoteActionRef types.TThostFtdcOrderActionRefType
	// 报价引用
	QuoteRef types.TThostFtdcOrderRefType
	// 请求编号
	RequestID types.TThostFtdcRequestIDType
	// 前置编号
	FrontID types.TThostFtdcFrontIDType
	// 会话编号
	SessionID types.TThostFtdcSessionIDType
	// 交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	// 报价操作编号
	QuoteSysID types.TThostFtdcOrderSysIDType
	// 操作标志
	ActionFlag types.TThostFtdcActionFlagType
	// 操作日期
	ActionDate types.TThostFtdcDateType
	// 操作时间
	ActionTime types.TThostFtdcTimeType
	// 交易所交易员代码
	TraderID types.TThostFtdcTraderIDType
	// 安装编号
	InstallID types.TThostFtdcInstallIDType
	// 本地报价编号
	QuoteLocalID types.TThostFtdcOrderLocalIDType
	// 操作本地编号
	ActionLocalID types.TThostFtdcOrderLocalIDType
	// 会员代码
	ParticipantID types.TThostFtdcParticipantIDType
	// 客户代码
	ClientID types.TThostFtdcClientIDType
	// 业务单元
	BusinessUnit types.TThostFtdcBusinessUnitType
	// 报单操作状态
	OrderActionStatus types.TThostFtdcOrderActionStatusType
	// 用户代码
	UserID types.TThostFtdcUserIDType
	// 状态信息
	StatusMsg types.TThostFtdcErrorMsgType
	// 合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
	// 营业部编号
	BranchID types.TThostFtdcBranchIDType
	// 投资单元代码
	InvestUnitID types.TThostFtdcInvestUnitIDType
	// IP地址
	IPAddress types.TThostFtdcIPAddressType
	// Mac地址
	MacAddress types.TThostFtdcMacAddressType
}

// 报价查询
type CThostFtdcQryQuoteField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	// 合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
	// 交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	// 报价编号
	QuoteSysID types.TThostFtdcOrderSysIDType
	// 开始时间
	InsertTimeStart types.TThostFtdcTimeType
	// 结束时间
	InsertTimeEnd types.TThostFtdcTimeType
}

// 交易所报价信息
type CThostFtdcExchangeQuoteField struct {
	// 卖价格
	AskPrice types.TThostFtdcPriceType
	// 买价格
	BidPrice types.TThostFtdcPriceType
	// 卖数量
	AskVolume types.TThostFtdcVolumeType
	// 买数量
	BidVolume types.TThostFtdcVolumeType
	// 请求编号
	RequestID types.TThostFtdcRequestIDType
	// 业务单元
	BusinessUnit types.TThostFtdcBusinessUnitType
	// 卖开平标志
	AskOffsetFlag types.TThostFtdcOffsetFlagType
	// 买开平标志
	BidOffsetFlag types.TThostFtdcOffsetFlagType
	// 卖投机套保标志
	AskHedgeFlag types.TThostFtdcHedgeFlagType
	// 买投机套保标志
	BidHedgeFlag types.TThostFtdcHedgeFlagType
	// 本地报价编号
	QuoteLocalID types.TThostFtdcOrderLocalIDType
	// 交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	// 会员代码
	ParticipantID types.TThostFtdcParticipantIDType
	// 客户代码
	ClientID types.TThostFtdcClientIDType
	// 合约在交易所的代码
	ExchangeInstID types.TThostFtdcExchangeInstIDType
	// 交易所交易员代码
	TraderID types.TThostFtdcTraderIDType
	// 安装编号
	InstallID types.TThostFtdcInstallIDType
	// 报价提示序号
	NotifySequence types.TThostFtdcSequenceNoType
	// 报价提交状态
	OrderSubmitStatus types.TThostFtdcOrderSubmitStatusType
	// 交易日
	TradingDay types.TThostFtdcDateType
	// 结算编号
	SettlementID types.TThostFtdcSettlementIDType
	// 报价编号
	QuoteSysID types.TThostFtdcOrderSysIDType
	// 报单日期
	InsertDate types.TThostFtdcDateType
	// 插入时间
	InsertTime types.TThostFtdcTimeType
	// 撤销时间
	CancelTime types.TThostFtdcTimeType
	// 报价状态
	QuoteStatus types.TThostFtdcOrderStatusType
	// 结算会员编号
	ClearingPartID types.TThostFtdcParticipantIDType
	// 序号
	SequenceNo types.TThostFtdcSequenceNoType
	// 卖方报单编号
	AskOrderSysID types.TThostFtdcOrderSysIDType
	// 买方报单编号
	BidOrderSysID types.TThostFtdcOrderSysIDType
	// 应价编号
	ForQuoteSysID types.TThostFtdcOrderSysIDType
	// 营业部编号
	BranchID types.TThostFtdcBranchIDType
	// IP地址
	IPAddress types.TThostFtdcIPAddressType
	// Mac地址
	MacAddress types.TThostFtdcMacAddressType
}

// 交易所报价查询
type CThostFtdcQryExchangeQuoteField struct {
	// 会员代码
	ParticipantID types.TThostFtdcParticipantIDType
	// 客户代码
	ClientID types.TThostFtdcClientIDType
	// 合约在交易所的代码
	ExchangeInstID types.TThostFtdcExchangeInstIDType
	// 交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	// 交易所交易员代码
	TraderID types.TThostFtdcTraderIDType
}

// 报价操作查询
type CThostFtdcQryQuoteActionField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	// 交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
}

// 交易所报价操作
type CThostFtdcExchangeQuoteActionField struct {
	// 交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	// 报价操作编号
	QuoteSysID types.TThostFtdcOrderSysIDType
	// 操作标志
	ActionFlag types.TThostFtdcActionFlagType
	// 操作日期
	ActionDate types.TThostFtdcDateType
	// 操作时间
	ActionTime types.TThostFtdcTimeType
	// 交易所交易员代码
	TraderID types.TThostFtdcTraderIDType
	// 安装编号
	InstallID types.TThostFtdcInstallIDType
	// 本地报价编号
	QuoteLocalID types.TThostFtdcOrderLocalIDType
	// 操作本地编号
	ActionLocalID types.TThostFtdcOrderLocalIDType
	// 会员代码
	ParticipantID types.TThostFtdcParticipantIDType
	// 客户代码
	ClientID types.TThostFtdcClientIDType
	// 业务单元
	BusinessUnit types.TThostFtdcBusinessUnitType
	// 报单操作状态
	OrderActionStatus types.TThostFtdcOrderActionStatusType
	// 用户代码
	UserID types.TThostFtdcUserIDType
	// IP地址
	IPAddress types.TThostFtdcIPAddressType
	// Mac地址
	MacAddress types.TThostFtdcMacAddressType
}

// 交易所报价操作查询
type CThostFtdcQryExchangeQuoteActionField struct {
	// 会员代码
	ParticipantID types.TThostFtdcParticipantIDType
	// 客户代码
	ClientID types.TThostFtdcClientIDType
	// 交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	// 交易所交易员代码
	TraderID types.TThostFtdcTraderIDType
}

// 期权合约delta值
type CThostFtdcOptionInstrDeltaField struct {
	// 合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
	// 投资者范围
	InvestorRange types.TThostFtdcInvestorRangeType
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	// Delta值
	Delta types.TThostFtdcRatioType
}

// 发给做市商的询价请求
type CThostFtdcForQuoteRspField struct {
	// 交易日
	TradingDay types.TThostFtdcDateType
	// 合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
	// 询价编号
	ForQuoteSysID types.TThostFtdcOrderSysIDType
	// 询价时间
	ForQuoteTime types.TThostFtdcTimeType
	// 业务日期
	ActionDay types.TThostFtdcDateType
	// 交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
}

// 当前期权合约执行偏移值的详细内容
type CThostFtdcStrikeOffsetField struct {
	// 合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
	// 投资者范围
	InvestorRange types.TThostFtdcInvestorRangeType
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	// 执行偏移值
	Offset types.TThostFtdcMoneyType
	// 执行偏移类型
	OffsetType types.TThostFtdcStrikeOffsetTypeType
}

// 期权执行偏移值查询
type CThostFtdcQryStrikeOffsetField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	// 合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
}

// 输入批量报单操作
type CThostFtdcInputBatchOrderActionField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	// 报单操作引用
	OrderActionRef types.TThostFtdcOrderActionRefType
	// 请求编号
	RequestID types.TThostFtdcRequestIDType
	// 前置编号
	FrontID types.TThostFtdcFrontIDType
	// 会话编号
	SessionID types.TThostFtdcSessionIDType
	// 交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	// 用户代码
	UserID types.TThostFtdcUserIDType
	// 投资单元代码
	InvestUnitID types.TThostFtdcInvestUnitIDType
	// IP地址
	IPAddress types.TThostFtdcIPAddressType
	// Mac地址
	MacAddress types.TThostFtdcMacAddressType
}

// 批量报单操作
type CThostFtdcBatchOrderActionField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	// 报单操作引用
	OrderActionRef types.TThostFtdcOrderActionRefType
	// 请求编号
	RequestID types.TThostFtdcRequestIDType
	// 前置编号
	FrontID types.TThostFtdcFrontIDType
	// 会话编号
	SessionID types.TThostFtdcSessionIDType
	// 交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	// 操作日期
	ActionDate types.TThostFtdcDateType
	// 操作时间
	ActionTime types.TThostFtdcTimeType
	// 交易所交易员代码
	TraderID types.TThostFtdcTraderIDType
	// 安装编号
	InstallID types.TThostFtdcInstallIDType
	// 操作本地编号
	ActionLocalID types.TThostFtdcOrderLocalIDType
	// 会员代码
	ParticipantID types.TThostFtdcParticipantIDType
	// 客户代码
	ClientID types.TThostFtdcClientIDType
	// 业务单元
	BusinessUnit types.TThostFtdcBusinessUnitType
	// 报单操作状态
	OrderActionStatus types.TThostFtdcOrderActionStatusType
	// 用户代码
	UserID types.TThostFtdcUserIDType
	// 状态信息
	StatusMsg types.TThostFtdcErrorMsgType
	// 投资单元代码
	InvestUnitID types.TThostFtdcInvestUnitIDType
	// IP地址
	IPAddress types.TThostFtdcIPAddressType
	// Mac地址
	MacAddress types.TThostFtdcMacAddressType
}

// 交易所批量报单操作
type CThostFtdcExchangeBatchOrderActionField struct {
	// 交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	// 操作日期
	ActionDate types.TThostFtdcDateType
	// 操作时间
	ActionTime types.TThostFtdcTimeType
	// 交易所交易员代码
	TraderID types.TThostFtdcTraderIDType
	// 安装编号
	InstallID types.TThostFtdcInstallIDType
	// 操作本地编号
	ActionLocalID types.TThostFtdcOrderLocalIDType
	// 会员代码
	ParticipantID types.TThostFtdcParticipantIDType
	// 客户代码
	ClientID types.TThostFtdcClientIDType
	// 业务单元
	BusinessUnit types.TThostFtdcBusinessUnitType
	// 报单操作状态
	OrderActionStatus types.TThostFtdcOrderActionStatusType
	// 用户代码
	UserID types.TThostFtdcUserIDType
	// IP地址
	IPAddress types.TThostFtdcIPAddressType
	// Mac地址
	MacAddress types.TThostFtdcMacAddressType
}

// 查询批量报单操作
type CThostFtdcQryBatchOrderActionField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	// 交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
}

// 输入的期权自对冲
type CThostFtdcInputOptionSelfCloseField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	// 合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
	// 期权自对冲引用
	OptionSelfCloseRef types.TThostFtdcOrderRefType
	// 用户代码
	UserID types.TThostFtdcUserIDType
	// 数量
	Volume types.TThostFtdcVolumeType
	// 请求编号
	RequestID types.TThostFtdcRequestIDType
	// 业务单元
	BusinessUnit types.TThostFtdcBusinessUnitType
	// 投机套保标志
	HedgeFlag types.TThostFtdcHedgeFlagType
	// 期权行权的头寸是否自对冲
	OptSelfCloseFlag types.TThostFtdcOptSelfCloseFlagType
	// 交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	// 投资单元代码
	InvestUnitID types.TThostFtdcInvestUnitIDType
	// 资金账号
	AccountID types.TThostFtdcAccountIDType
	// 币种代码
	CurrencyID types.TThostFtdcCurrencyIDType
	// 交易编码
	ClientID types.TThostFtdcClientIDType
	// IP地址
	IPAddress types.TThostFtdcIPAddressType
	// Mac地址
	MacAddress types.TThostFtdcMacAddressType
}

// 输入期权自对冲操作
type CThostFtdcInputOptionSelfCloseActionField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	// 期权自对冲操作引用
	OptionSelfCloseActionRef types.TThostFtdcOrderActionRefType
	// 期权自对冲引用
	OptionSelfCloseRef types.TThostFtdcOrderRefType
	// 请求编号
	RequestID types.TThostFtdcRequestIDType
	// 前置编号
	FrontID types.TThostFtdcFrontIDType
	// 会话编号
	SessionID types.TThostFtdcSessionIDType
	// 交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	// 期权自对冲操作编号
	OptionSelfCloseSysID types.TThostFtdcOrderSysIDType
	// 操作标志
	ActionFlag types.TThostFtdcActionFlagType
	// 用户代码
	UserID types.TThostFtdcUserIDType
	// 合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
	// 投资单元代码
	InvestUnitID types.TThostFtdcInvestUnitIDType
	// IP地址
	IPAddress types.TThostFtdcIPAddressType
	// Mac地址
	MacAddress types.TThostFtdcMacAddressType
}

// 期权自对冲
type CThostFtdcOptionSelfCloseField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	// 合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
	// 期权自对冲引用
	OptionSelfCloseRef types.TThostFtdcOrderRefType
	// 用户代码
	UserID types.TThostFtdcUserIDType
	// 数量
	Volume types.TThostFtdcVolumeType
	// 请求编号
	RequestID types.TThostFtdcRequestIDType
	// 业务单元
	BusinessUnit types.TThostFtdcBusinessUnitType
	// 投机套保标志
	HedgeFlag types.TThostFtdcHedgeFlagType
	// 期权行权的头寸是否自对冲
	OptSelfCloseFlag types.TThostFtdcOptSelfCloseFlagType
	// 本地期权自对冲编号
	OptionSelfCloseLocalID types.TThostFtdcOrderLocalIDType
	// 交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	// 会员代码
	ParticipantID types.TThostFtdcParticipantIDType
	// 客户代码
	ClientID types.TThostFtdcClientIDType
	// 合约在交易所的代码
	ExchangeInstID types.TThostFtdcExchangeInstIDType
	// 交易所交易员代码
	TraderID types.TThostFtdcTraderIDType
	// 安装编号
	InstallID types.TThostFtdcInstallIDType
	// 期权自对冲提交状态
	OrderSubmitStatus types.TThostFtdcOrderSubmitStatusType
	// 报单提示序号
	NotifySequence types.TThostFtdcSequenceNoType
	// 交易日
	TradingDay types.TThostFtdcDateType
	// 结算编号
	SettlementID types.TThostFtdcSettlementIDType
	// 期权自对冲编号
	OptionSelfCloseSysID types.TThostFtdcOrderSysIDType
	// 报单日期
	InsertDate types.TThostFtdcDateType
	// 插入时间
	InsertTime types.TThostFtdcTimeType
	// 撤销时间
	CancelTime types.TThostFtdcTimeType
	// 自对冲结果
	ExecResult types.TThostFtdcExecResultType
	// 结算会员编号
	ClearingPartID types.TThostFtdcParticipantIDType
	// 序号
	SequenceNo types.TThostFtdcSequenceNoType
	// 前置编号
	FrontID types.TThostFtdcFrontIDType
	// 会话编号
	SessionID types.TThostFtdcSessionIDType
	// 用户端产品信息
	UserProductInfo types.TThostFtdcProductInfoType
	// 状态信息
	StatusMsg types.TThostFtdcErrorMsgType
	// 操作用户代码
	ActiveUserID types.TThostFtdcUserIDType
	// 经纪公司报单编号
	BrokerOptionSelfCloseSeq types.TThostFtdcSequenceNoType
	// 营业部编号
	BranchID types.TThostFtdcBranchIDType
	// 投资单元代码
	InvestUnitID types.TThostFtdcInvestUnitIDType
	// 资金账号
	AccountID types.TThostFtdcAccountIDType
	// 币种代码
	CurrencyID types.TThostFtdcCurrencyIDType
	// IP地址
	IPAddress types.TThostFtdcIPAddressType
	// Mac地址
	MacAddress types.TThostFtdcMacAddressType
}

// 期权自对冲操作
type CThostFtdcOptionSelfCloseActionField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	// 期权自对冲操作引用
	OptionSelfCloseActionRef types.TThostFtdcOrderActionRefType
	// 期权自对冲引用
	OptionSelfCloseRef types.TThostFtdcOrderRefType
	// 请求编号
	RequestID types.TThostFtdcRequestIDType
	// 前置编号
	FrontID types.TThostFtdcFrontIDType
	// 会话编号
	SessionID types.TThostFtdcSessionIDType
	// 交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	// 期权自对冲操作编号
	OptionSelfCloseSysID types.TThostFtdcOrderSysIDType
	// 操作标志
	ActionFlag types.TThostFtdcActionFlagType
	// 操作日期
	ActionDate types.TThostFtdcDateType
	// 操作时间
	ActionTime types.TThostFtdcTimeType
	// 交易所交易员代码
	TraderID types.TThostFtdcTraderIDType
	// 安装编号
	InstallID types.TThostFtdcInstallIDType
	// 本地期权自对冲编号
	OptionSelfCloseLocalID types.TThostFtdcOrderLocalIDType
	// 操作本地编号
	ActionLocalID types.TThostFtdcOrderLocalIDType
	// 会员代码
	ParticipantID types.TThostFtdcParticipantIDType
	// 客户代码
	ClientID types.TThostFtdcClientIDType
	// 业务单元
	BusinessUnit types.TThostFtdcBusinessUnitType
	// 报单操作状态
	OrderActionStatus types.TThostFtdcOrderActionStatusType
	// 用户代码
	UserID types.TThostFtdcUserIDType
	// 状态信息
	StatusMsg types.TThostFtdcErrorMsgType
	// 合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
	// 营业部编号
	BranchID types.TThostFtdcBranchIDType
	// 投资单元代码
	InvestUnitID types.TThostFtdcInvestUnitIDType
	// IP地址
	IPAddress types.TThostFtdcIPAddressType
	// Mac地址
	MacAddress types.TThostFtdcMacAddressType
}

// 期权自对冲查询
type CThostFtdcQryOptionSelfCloseField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	// 合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
	// 交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	// 期权自对冲编号
	OptionSelfCloseSysID types.TThostFtdcOrderSysIDType
	// 开始时间
	InsertTimeStart types.TThostFtdcTimeType
	// 结束时间
	InsertTimeEnd types.TThostFtdcTimeType
}

// 交易所期权自对冲信息
type CThostFtdcExchangeOptionSelfCloseField struct {
	// 数量
	Volume types.TThostFtdcVolumeType
	// 请求编号
	RequestID types.TThostFtdcRequestIDType
	// 业务单元
	BusinessUnit types.TThostFtdcBusinessUnitType
	// 投机套保标志
	HedgeFlag types.TThostFtdcHedgeFlagType
	// 期权行权的头寸是否自对冲
	OptSelfCloseFlag types.TThostFtdcOptSelfCloseFlagType
	// 本地期权自对冲编号
	OptionSelfCloseLocalID types.TThostFtdcOrderLocalIDType
	// 交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	// 会员代码
	ParticipantID types.TThostFtdcParticipantIDType
	// 客户代码
	ClientID types.TThostFtdcClientIDType
	// 合约在交易所的代码
	ExchangeInstID types.TThostFtdcExchangeInstIDType
	// 交易所交易员代码
	TraderID types.TThostFtdcTraderIDType
	// 安装编号
	InstallID types.TThostFtdcInstallIDType
	// 期权自对冲提交状态
	OrderSubmitStatus types.TThostFtdcOrderSubmitStatusType
	// 报单提示序号
	NotifySequence types.TThostFtdcSequenceNoType
	// 交易日
	TradingDay types.TThostFtdcDateType
	// 结算编号
	SettlementID types.TThostFtdcSettlementIDType
	// 期权自对冲编号
	OptionSelfCloseSysID types.TThostFtdcOrderSysIDType
	// 报单日期
	InsertDate types.TThostFtdcDateType
	// 插入时间
	InsertTime types.TThostFtdcTimeType
	// 撤销时间
	CancelTime types.TThostFtdcTimeType
	// 自对冲结果
	ExecResult types.TThostFtdcExecResultType
	// 结算会员编号
	ClearingPartID types.TThostFtdcParticipantIDType
	// 序号
	SequenceNo types.TThostFtdcSequenceNoType
	// 营业部编号
	BranchID types.TThostFtdcBranchIDType
	// IP地址
	IPAddress types.TThostFtdcIPAddressType
	// Mac地址
	MacAddress types.TThostFtdcMacAddressType
}

// 期权自对冲操作查询
type CThostFtdcQryOptionSelfCloseActionField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	// 交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
}

// 交易所期权自对冲操作
type CThostFtdcExchangeOptionSelfCloseActionField struct {
	// 交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	// 期权自对冲操作编号
	OptionSelfCloseSysID types.TThostFtdcOrderSysIDType
	// 操作标志
	ActionFlag types.TThostFtdcActionFlagType
	// 操作日期
	ActionDate types.TThostFtdcDateType
	// 操作时间
	ActionTime types.TThostFtdcTimeType
	// 交易所交易员代码
	TraderID types.TThostFtdcTraderIDType
	// 安装编号
	InstallID types.TThostFtdcInstallIDType
	// 本地期权自对冲编号
	OptionSelfCloseLocalID types.TThostFtdcOrderLocalIDType
	// 操作本地编号
	ActionLocalID types.TThostFtdcOrderLocalIDType
	// 会员代码
	ParticipantID types.TThostFtdcParticipantIDType
	// 客户代码
	ClientID types.TThostFtdcClientIDType
	// 业务单元
	BusinessUnit types.TThostFtdcBusinessUnitType
	// 报单操作状态
	OrderActionStatus types.TThostFtdcOrderActionStatusType
	// 用户代码
	UserID types.TThostFtdcUserIDType
	// 营业部编号
	BranchID types.TThostFtdcBranchIDType
	// IP地址
	IPAddress types.TThostFtdcIPAddressType
	// Mac地址
	MacAddress types.TThostFtdcMacAddressType
}

// 组合合约安全系数
type CThostFtdcCombInstrumentGuardField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
	GuarantRatio types.TThostFtdcRatioType
}

// 组合合约安全系数查询
type CThostFtdcQryCombInstrumentGuardField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
}

// 输入的套利套保确认请求
type CThostFtdcInputHedgeConfirmField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	// 合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
	// 组合引用
	OrderRef types.TThostFtdcOrderRefType
	// 用户代码
	UserID types.TThostFtdcUserIDType
	// 买卖方向
	Direction types.TThostFtdcDirectionType
	// 数量
	Volume types.TThostFtdcVolumeType
	// 请求类型
	RequestType types.TThostFtdcRequestTypeType
	// 投机套保标志
	HedgeFlag types.TThostFtdcHedgeFlagType
	// 业务单元
	BusinessUnit types.TThostFtdcBusinessUnitType
	// 交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	// 第二腿合约代码
	SecondLeg types.TThostFtdcInstrumentIDType
	// 组合类型
	CombinationType types.TThostFtdcCzceCombinationTypeType
	// 投资单元代码
	InvestUnitID types.TThostFtdcInvestUnitIDType
	// 前置编号
	FrontID types.TThostFtdcFrontIDType
	// 会话编号
	SessionID types.TThostFtdcSessionIDType
}

// 输入套利套保确认操作
type CThostFtdcInputHedgeConfirmActionField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	// 报单操作引用
	OrderActionRef types.TThostFtdcOrderActionRefType
	// 报单引用
	OrderRef types.TThostFtdcOrderRefType
	// 请求编号
	RequestID types.TThostFtdcRequestIDType
	// 前置编号
	FrontID types.TThostFtdcFrontIDType
	// 会话编号
	SessionID types.TThostFtdcSessionIDType
	// 交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	// 报单编号
	OrderSysID types.TThostFtdcOrderSysIDType
	// 操作标志
	ActionFlag types.TThostFtdcActionFlagType
	// 用户代码
	UserID types.TThostFtdcUserIDType
	// 请求类型
	RequestType types.TThostFtdcRequestTypeType
	// 合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
	// 投资单元代码
	InvestUnitID types.TThostFtdcInvestUnitIDType
	// IP地址
	IPAddress types.TThostFtdcIPAddressType
	// Mac地址
	MacAddress types.TThostFtdcMacAddressType
}

// 套利套保确认
type CThostFtdcHedgeConfirmField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	// 合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
	// 组合引用
	OrderRef types.TThostFtdcOrderRefType
	// 用户代码
	UserID types.TThostFtdcUserIDType
	// 买卖方向
	Direction types.TThostFtdcDirectionType
	// 数量
	Volume types.TThostFtdcVolumeType
	// 合同编号
	OrderSysID types.TThostFtdcOrderSysIDType
	// 请求类型
	RequestType types.TThostFtdcRequestTypeType
	// 投机套保标志
	HedgeFlag types.TThostFtdcHedgeFlagType
	// 业务单元
	BusinessUnit types.TThostFtdcBusinessUnitType
	// 本地定单号
	OrderLocalID types.TThostFtdcOrderLocalIDType
	// 交易所交易员代码
	TraderID types.TThostFtdcTraderIDType
	// 交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	// 交易所会员代码
	ParticipantID types.TThostFtdcParticipantIDType
	// 投资者交易编码
	ClientID types.TThostFtdcClientIDType
	// 第二腿合约代码
	SecondLeg types.TThostFtdcInstrumentIDType
	// 请求状态
	RequestStatus types.TThostFtdcRequestStatusType
	// 组合合约编号
	CombInstrumentID types.TThostFtdcInstrumentIDType
	// 状态信息
	StatusMsg types.TThostFtdcErrorMsgType
	// 前置编号
	FrontID types.TThostFtdcFrontIDType
	// 会话编号
	SessionID types.TThostFtdcSessionIDType
	// 投资单元代码
	InvestUnitID types.TThostFtdcInvestUnitIDType
	// 组合类型
	CombinationType types.TThostFtdcCzceCombinationTypeType
	// 报单日期
	InsertDate types.TThostFtdcDateType
}

// 套利套保确认撤销
type CThostFtdcHedgeConfirmActionField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	// 投资者说明的对报单操作的唯一引用
	OrderActionRef types.TThostFtdcOrderActionRefType
	// 组合引用
	OrderRef types.TThostFtdcOrderRefType
	// 请求编号
	RequestID types.TThostFtdcRequestIDType
	// 前置编号
	FrontID types.TThostFtdcFrontIDType
	// 会话编号
	SessionID types.TThostFtdcSessionIDType
	// 交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	// 报单操作编号
	OrderSysID types.TThostFtdcOrderSysIDType
	// 操作标志
	ActionFlag types.TThostFtdcActionFlagType
	// 交易所交易员代码
	TraderID types.TThostFtdcTraderIDType
	// 交易所交易员报盘机编号
	InstallID types.TThostFtdcInstallIDType
	// 本地报单编号
	OrderLocalID types.TThostFtdcOrderLocalIDType
	// 操作本地编号
	ActionLocalID types.TThostFtdcOrderLocalIDType
	// 每一位会员在系统中的编码，编号唯一且遵循交易所制定的编码规则
	ParticipantID types.TThostFtdcParticipantIDType
	// 客户在系统中的编号，编号唯一且遵循交易所制定的编码规则
	ClientID types.TThostFtdcClientIDType
	// 业务单元
	BusinessUnit types.TThostFtdcBusinessUnitType
	// 报单操作状态
	OrderActionStatus types.TThostFtdcOrderActionStatusType
	// 用户代码
	UserID types.TThostFtdcUserIDType
	// 请求类型
	RequestType types.TThostFtdcRequestTypeType
	// 状态信息
	StatusMsg types.TThostFtdcErrorMsgType
	// 合约在系统中的编号
	InstrumentID types.TThostFtdcInstrumentIDType
	// 第二腿合约代码
	SecondLeg types.TThostFtdcInstrumentIDType
	// 组合合约的编号
	CombInstrumentID types.TThostFtdcInstrumentIDType
	// 营业部编号
	BranchID types.TThostFtdcBranchIDType
	// 投资单元代码
	InvestUnitID types.TThostFtdcInvestUnitIDType
	// IP地址
	IPAddress types.TThostFtdcIPAddressType
	// Mac地址
	MacAddress types.TThostFtdcMacAddressType
}

// 输入的申请组合
type CThostFtdcInputCombActionField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	// 合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
	// 组合引用
	CombActionRef types.TThostFtdcOrderRefType
	// 用户代码
	UserID types.TThostFtdcUserIDType
	// 买卖方向
	Direction types.TThostFtdcDirectionType
	// 数量
	Volume types.TThostFtdcVolumeType
	// 组合指令方向
	CombDirection types.TThostFtdcCombDirectionType
	// 投机套保标志
	HedgeFlag types.TThostFtdcHedgeFlagType
	// 业务单元
	BusinessUnit types.TThostFtdcBusinessUnitType
	// 交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	// IP地址
	IPAddress types.TThostFtdcIPAddressType
	// Mac地址
	MacAddress types.TThostFtdcMacAddressType
	// 投资单元代码
	InvestUnitID types.TThostFtdcInvestUnitIDType
	// 前置编号
	FrontID types.TThostFtdcFrontIDType
	// 会话编号
	SessionID types.TThostFtdcSessionIDType
}

// 申请组合
type CThostFtdcCombActionField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	// 合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
	// 组合引用
	CombActionRef types.TThostFtdcOrderRefType
	// 用户代码
	UserID types.TThostFtdcUserIDType
	// 买卖方向
	Direction types.TThostFtdcDirectionType
	// 数量
	Volume types.TThostFtdcVolumeType
	// 组合指令方向
	CombDirection types.TThostFtdcCombDirectionType
	// 投机套保标志
	HedgeFlag types.TThostFtdcHedgeFlagType
	// 业务单元
	BusinessUnit types.TThostFtdcBusinessUnitType
	// 本地申请组合编号
	ActionLocalID types.TThostFtdcOrderLocalIDType
	// 交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	// 会员代码
	ParticipantID types.TThostFtdcParticipantIDType
	// 客户代码
	ClientID types.TThostFtdcClientIDType
	// 合约在交易所的代码
	ExchangeInstID types.TThostFtdcExchangeInstIDType
	// 交易所交易员代码
	TraderID types.TThostFtdcTraderIDType
	// 安装编号
	InstallID types.TThostFtdcInstallIDType
	// 组合状态
	ActionStatus types.TThostFtdcOrderActionStatusType
	// 报单提示序号
	NotifySequence types.TThostFtdcSequenceNoType
	// 交易日
	TradingDay types.TThostFtdcDateType
	// 结算编号
	SettlementID types.TThostFtdcSettlementIDType
	// 序号
	SequenceNo types.TThostFtdcSequenceNoType
	// 前置编号
	FrontID types.TThostFtdcFrontIDType
	// 会话编号
	SessionID types.TThostFtdcSessionIDType
	// 用户端产品信息
	UserProductInfo types.TThostFtdcProductInfoType
	// 状态信息
	StatusMsg types.TThostFtdcErrorMsgType
	// IP地址
	IPAddress types.TThostFtdcIPAddressType
	// Mac地址
	MacAddress types.TThostFtdcMacAddressType
	// 组合编号
	ComTradeID types.TThostFtdcTradeIDType
}

// 组合单腿汇总
type CThostFtdcInvestorPositionForCombField struct {
	// 交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	// 合约代码
	LegInstrumentID types.TThostFtdcInstrumentIDType
	// 投机套保标志
	LegHedgeFlag types.TThostFtdcHedgeFlagType
	// 买卖方向
	LegDirection types.TThostFtdcDirectionType
	// 数量
	TotalAmt types.TThostFtdcVolumeType
	// LegID
	LegID types.TThostFtdcLegIDType
	// 组合优先级
	TradeGroupID types.TThostFtdcTradeGroupIDType
	// 组合合约代码
	CombInstrumentID types.TThostFtdcInstrumentIDType
	// 组合投机套保标志
	CombHedgeFlag types.TThostFtdcHedgeFlagType
	// 组合类型
	CombinationType types.TThostFtdcCombinationTypeType
	// 组合成交编号
	CombTradeID types.TThostFtdcTradeIDType
	// 成交日期
	OpenDate types.TThostFtdcDateType
	// 成交编号
	TradeID types.TThostFtdcTradeIDType
}

// 申请组合查询
type CThostFtdcQryCombActionField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	// 合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
	// 交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
}

// 套利套保确认查询
type CThostFtdcQryHedgeConfirmField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	// 合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
	// 交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
}

// 申请组合单腿汇总查询
type CThostFtdcQryInvestorPositionForCombField struct {
	// 交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	// 单腿合约代码
	LegInstrumentID types.TThostFtdcInstrumentIDType
}

// 交易所申请组合信息
type CThostFtdcExchangeCombActionField struct {
	// 买卖方向
	Direction types.TThostFtdcDirectionType
	// 数量
	Volume types.TThostFtdcVolumeType
	// 组合指令方向
	CombDirection types.TThostFtdcCombDirectionType
	// 投机套保标志
	HedgeFlag types.TThostFtdcHedgeFlagType
	// 本地申请组合编号
	ActionLocalID types.TThostFtdcOrderLocalIDType
	// 交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	// 会员代码
	ParticipantID types.TThostFtdcParticipantIDType
	// 客户代码
	ClientID types.TThostFtdcClientIDType
	// 合约在交易所的代码
	ExchangeInstID types.TThostFtdcExchangeInstIDType
	// 交易所交易员代码
	TraderID types.TThostFtdcTraderIDType
	// 安装编号
	InstallID types.TThostFtdcInstallIDType
	// 组合状态
	ActionStatus types.TThostFtdcOrderActionStatusType
	// 报单提示序号
	NotifySequence types.TThostFtdcSequenceNoType
	// 交易日
	TradingDay types.TThostFtdcDateType
	// 结算编号
	SettlementID types.TThostFtdcSettlementIDType
	// 序号
	SequenceNo types.TThostFtdcSequenceNoType
	// IP地址
	IPAddress types.TThostFtdcIPAddressType
	// Mac地址
	MacAddress types.TThostFtdcMacAddressType
}

// 交易所申请组合查询
type CThostFtdcQryExchangeCombActionField struct {
	// 会员代码
	ParticipantID types.TThostFtdcParticipantIDType
	// 客户代码
	ClientID types.TThostFtdcClientIDType
	// 合约在交易所的代码
	ExchangeInstID types.TThostFtdcExchangeInstIDType
	// 交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	// 交易所交易员代码
	TraderID types.TThostFtdcTraderIDType
}

// 产品报价汇率
type CThostFtdcProductExchRateField struct {
	// 产品代码
	ProductID types.TThostFtdcInstrumentIDType
	// 报价币种类型
	QuoteCurrencyID types.TThostFtdcCurrencyIDType
	// 汇率
	ExchangeRate types.TThostFtdcExchangeRateType
}

// 产品报价汇率查询
type CThostFtdcQryProductExchRateField struct {
	// 产品代码
	ProductID types.TThostFtdcInstrumentIDType
}

// 查询询价价差
type CThostFtdcQryForQuoteParamField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
	// 交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
}

// 查询SPBM品种明细
type CThostFtdcQryInvestorProdSPBMDetailField struct {
	// 交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	// 品种代码
	ProdFamilyCode types.TThostFtdcInstrumentIDType
}

// 查询SPMM商品群明细
type CThostFtdcQrySPMMInvestorCommodityGroupMarginField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	// 商品群代码
	CommodityGroupID types.TThostFtdcSPMMProductIDType
}

// 查询RULE产品保证金
type CThostFtdcQryRULEInvestorProdMarginField struct {
	// 交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	// 品种代码
	ProdFamilyCode types.TThostFtdcProductIDType
	// 商品群号
	CommodityGroupID types.TThostFtdcRULECommodityGroupIDType
}

// 询价价差
type CThostFtdcForQuoteParamField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
	// 交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	// 最新价
	LastPrice types.TThostFtdcPriceType
	// 价差
	PriceInterval types.TThostFtdcPriceType
}

// 当前做市商期权合约手续费的详细内容
type CThostFtdcMMOptionInstrCommRateField struct {
	// 合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
	// 投资者范围
	InvestorRange types.TThostFtdcInvestorRangeType
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	// 开仓手续费率
	OpenRatioByMoney types.TThostFtdcRatioType
	// 开仓手续费
	OpenRatioByVolume types.TThostFtdcRatioType
	// 平仓手续费率
	CloseRatioByMoney types.TThostFtdcRatioType
	// 平仓手续费
	CloseRatioByVolume types.TThostFtdcRatioType
	// 平今手续费率
	CloseTodayRatioByMoney types.TThostFtdcRatioType
	// 平今手续费
	CloseTodayRatioByVolume types.TThostFtdcRatioType
	// 执行手续费率
	StrikeRatioByMoney types.TThostFtdcRatioType
	// 执行手续费
	StrikeRatioByVolume types.TThostFtdcRatioType
}

// 做市商期权手续费率查询
type CThostFtdcQryMMOptionInstrCommRateField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	// 合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
}

// 做市商合约手续费率
type CThostFtdcMMInstrumentCommissionRateField struct {
	// 合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
	// 投资者范围
	InvestorRange types.TThostFtdcInvestorRangeType
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	// 开仓手续费率
	OpenRatioByMoney types.TThostFtdcRatioType
	// 开仓手续费
	OpenRatioByVolume types.TThostFtdcRatioType
	// 平仓手续费率
	CloseRatioByMoney types.TThostFtdcRatioType
	// 平仓手续费
	CloseRatioByVolume types.TThostFtdcRatioType
	// 平今手续费率
	CloseTodayRatioByMoney types.TThostFtdcRatioType
	// 平今手续费
	CloseTodayRatioByVolume types.TThostFtdcRatioType
}

// 查询做市商合约手续费率
type CThostFtdcQryMMInstrumentCommissionRateField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	// 合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
}

// 当前报单手续费的详细内容
type CThostFtdcInstrumentOrderCommRateField struct {
	// 合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
	// 投资者范围
	InvestorRange types.TThostFtdcInvestorRangeType
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	// 投机套保标志
	HedgeFlag types.TThostFtdcHedgeFlagType
	// 报单手续费
	OrderCommByVolume types.TThostFtdcRatioType
	// 撤单手续费
	OrderActionCommByVolume types.TThostFtdcRatioType
}

// 申报费率查询
type CThostFtdcQryInstrumentOrderCommRateField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	// 合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
	// 投机套保标志
	HedgeFlag types.TThostFtdcHedgeFlagType
}

// 市场行情
type CThostFtdcMarketDataField struct {
	// 交易日
	TradingDay types.TThostFtdcDateType
	// 合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
	// 交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	// 合约在交易所的代码
	ExchangeInstID types.TThostFtdcExchangeInstIDType
	// 最新价
	LastPrice types.TThostFtdcPriceType
	// 上次结算价
	PreSettlementPrice types.TThostFtdcPriceType
	// 昨收盘
	PreClosePrice types.TThostFtdcPriceType
	// 昨持仓量
	PreOpenInterest types.TThostFtdcLargeVolumeType
	// 今开盘
	OpenPrice types.TThostFtdcPriceType
	// 最高价
	HighestPrice types.TThostFtdcPriceType
	// 最低价
	LowestPrice types.TThostFtdcPriceType
	// 数量
	Volume types.TThostFtdcVolumeType
	// 成交金额
	Turnover types.TThostFtdcMoneyType
	// 持仓量
	OpenInterest types.TThostFtdcLargeVolumeType
	// 今收盘
	ClosePrice types.TThostFtdcPriceType
	// 本次结算价
	SettlementPrice types.TThostFtdcPriceType
	// 涨停板价
	UpperLimitPrice types.TThostFtdcPriceType
	// 跌停板价
	LowerLimitPrice types.TThostFtdcPriceType
	// 昨虚实度
	PreDelta types.TThostFtdcRatioType
	// 今虚实度
	CurrDelta types.TThostFtdcRatioType
	// 最后修改时间
	UpdateTime types.TThostFtdcTimeType
	// 最后修改毫秒
	UpdateMillisec types.TThostFtdcMillisecType
	// 业务日期
	ActionDay types.TThostFtdcDateType
}

// 行情基础属性
type CThostFtdcMarketDataBaseField struct {
	// 交易日
	TradingDay types.TThostFtdcDateType
	// 上次结算价
	PreSettlementPrice types.TThostFtdcPriceType
	// 昨收盘
	PreClosePrice types.TThostFtdcPriceType
	// 昨持仓量
	PreOpenInterest types.TThostFtdcLargeVolumeType
	// 昨虚实度
	PreDelta types.TThostFtdcRatioType
}

// 行情静态属性
type CThostFtdcMarketDataStaticField struct {
	// 今开盘
	OpenPrice types.TThostFtdcPriceType
	// 最高价
	HighestPrice types.TThostFtdcPriceType
	// 最低价
	LowestPrice types.TThostFtdcPriceType
	// 今收盘
	ClosePrice types.TThostFtdcPriceType
	// 涨停板价
	UpperLimitPrice types.TThostFtdcPriceType
	// 跌停板价
	LowerLimitPrice types.TThostFtdcPriceType
	// 本次结算价
	SettlementPrice types.TThostFtdcPriceType
	// 今虚实度
	CurrDelta types.TThostFtdcRatioType
}

// 行情最新成交属性
type CThostFtdcMarketDataLastMatchField struct {
	// 最新价
	LastPrice types.TThostFtdcPriceType
	// 数量
	Volume types.TThostFtdcVolumeType
	// 成交金额
	Turnover types.TThostFtdcMoneyType
	// 持仓量
	OpenInterest types.TThostFtdcLargeVolumeType
}

// 行情最优价属性
type CThostFtdcMarketDataBestPriceField struct {
	// 申买价一
	BidPrice1 types.TThostFtdcPriceType
	// 申买量一
	BidVolume1 types.TThostFtdcVolumeType
	// 申卖价一
	AskPrice1 types.TThostFtdcPriceType
	// 申卖量一
	AskVolume1 types.TThostFtdcVolumeType
}

// 行情申买二、三属性
type CThostFtdcMarketDataBid23Field struct {
	// 申买价二
	BidPrice2 types.TThostFtdcPriceType
	// 申买量二
	BidVolume2 types.TThostFtdcVolumeType
	// 申买价三
	BidPrice3 types.TThostFtdcPriceType
	// 申买量三
	BidVolume3 types.TThostFtdcVolumeType
}

// 行情申卖二、三属性
type CThostFtdcMarketDataAsk23Field struct {
	// 申卖价二
	AskPrice2 types.TThostFtdcPriceType
	// 申卖量二
	AskVolume2 types.TThostFtdcVolumeType
	// 申卖价三
	AskPrice3 types.TThostFtdcPriceType
	// 申卖量三
	AskVolume3 types.TThostFtdcVolumeType
}

// 行情申买四、五属性
type CThostFtdcMarketDataBid45Field struct {
	// 申买价四
	BidPrice4 types.TThostFtdcPriceType
	// 申买量四
	BidVolume4 types.TThostFtdcVolumeType
	// 申买价五
	BidPrice5 types.TThostFtdcPriceType
	// 申买量五
	BidVolume5 types.TThostFtdcVolumeType
}

// 行情申卖四、五属性
type CThostFtdcMarketDataAsk45Field struct {
	// 申卖价四
	AskPrice4 types.TThostFtdcPriceType
	// 申卖量四
	AskVolume4 types.TThostFtdcVolumeType
	// 申卖价五
	AskPrice5 types.TThostFtdcPriceType
	// 申卖量五
	AskVolume5 types.TThostFtdcVolumeType
}

// 行情更新时间属性
type CThostFtdcMarketDataUpdateTimeField struct {
	// 合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
	// 最后修改时间
	UpdateTime types.TThostFtdcTimeType
	// 最后修改毫秒
	UpdateMillisec types.TThostFtdcMillisecType
	// 业务日期
	ActionDay types.TThostFtdcDateType
}

// 行情交易所代码属性
type CThostFtdcMarketDataExchangeField struct {
	// 交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
}

// 指定的合约
type CThostFtdcSpecificInstrumentField struct {
	// 合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
}

// 合约状态
type CThostFtdcInstrumentStatusField struct {
	// 交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	// 合约在交易所的代码
	ExchangeInstID types.TThostFtdcExchangeInstIDType
	// 结算组代码
	SettlementGroupID types.TThostFtdcSettlementGroupIDType
	// 合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
	// 合约交易状态
	InstrumentStatus types.TThostFtdcInstrumentStatusType
	// 交易阶段编号
	TradingSegmentSN types.TThostFtdcTradingSegmentSNType
	// 进入本状态时间
	EnterTime types.TThostFtdcTimeType
	// 进入本状态原因
	EnterReason types.TThostFtdcInstStatusEnterReasonType
}

// 查询合约状态
type CThostFtdcQryInstrumentStatusField struct {
	// 交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	// 合约在交易所的代码
	ExchangeInstID types.TThostFtdcExchangeInstIDType
}

// 投资者账户
type CThostFtdcInvestorAccountField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	// 投资者帐号
	AccountID types.TThostFtdcAccountIDType
	// 币种代码
	CurrencyID types.TThostFtdcCurrencyIDType
}

// 浮动盈亏算法
type CThostFtdcPositionProfitAlgorithmField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 投资者帐号
	AccountID types.TThostFtdcAccountIDType
	// 盈亏算法
	Algorithm types.TThostFtdcAlgorithmType
	// 备注
	Memo types.TThostFtdcMemoType
	// 币种代码
	CurrencyID types.TThostFtdcCurrencyIDType
}

// 会员资金折扣
type CThostFtdcDiscountField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 投资者范围
	InvestorRange types.TThostFtdcInvestorRangeType
	// 投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	// 资金折扣比例
	Discount types.TThostFtdcRatioType
}

// 查询转帐银行
type CThostFtdcQryTransferBankField struct {
	// 银行代码
	BankID types.TThostFtdcBankIDType
	// 银行分中心代码
	BankBrchID types.TThostFtdcBankBrchIDType
}

// 转帐银行
type CThostFtdcTransferBankField struct {
	// 银行代码
	BankID types.TThostFtdcBankIDType
	// 银行分中心代码
	BankBrchID types.TThostFtdcBankBrchIDType
	// 银行名称
	BankName types.TThostFtdcBankNameType
	// 是否活跃
	IsActive types.TThostFtdcBoolType
}

// 查询投资者持仓明细
type CThostFtdcQryInvestorPositionDetailField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	// 合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
}

// 投资者持仓明细
type CThostFtdcInvestorPositionDetailField struct {
	// 合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	// 投机套保标志
	HedgeFlag types.TThostFtdcHedgeFlagType
	// 买卖
	Direction types.TThostFtdcDirectionType
	// 开仓日期
	OpenDate types.TThostFtdcDateType
	// 成交编号
	TradeID types.TThostFtdcTradeIDType
	// 数量
	Volume types.TThostFtdcVolumeType
	// 开仓价
	OpenPrice types.TThostFtdcPriceType
	// 交易日
	TradingDay types.TThostFtdcDateType
	// 结算编号
	SettlementID types.TThostFtdcSettlementIDType
	// 成交类型
	TradeType types.TThostFtdcTradeTypeType
	// 组合合约代码
	CombInstrumentID types.TThostFtdcInstrumentIDType
	// 交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	// 逐日盯市平仓盈亏
	CloseProfitByDate types.TThostFtdcMoneyType
	// 逐笔对冲平仓盈亏
	CloseProfitByTrade types.TThostFtdcMoneyType
	// 逐日盯市持仓盈亏
	PositionProfitByDate types.TThostFtdcMoneyType
	// 逐笔对冲持仓盈亏
	PositionProfitByTrade types.TThostFtdcMoneyType
	// 投资者保证金
	Margin types.TThostFtdcMoneyType
	// 交易所保证金
	ExchMargin types.TThostFtdcMoneyType
	// 保证金率
	MarginRateByMoney types.TThostFtdcRatioType
	// 保证金率(按手数)
	MarginRateByVolume types.TThostFtdcRatioType
	// 昨结算价
	LastSettlementPrice types.TThostFtdcPriceType
	// 结算价
	SettlementPrice types.TThostFtdcPriceType
	// 平仓量
	CloseVolume types.TThostFtdcVolumeType
	// 平仓金额
	CloseAmount types.TThostFtdcMoneyType
}

// 资金账户口令域
type CThostFtdcTradingAccountPasswordField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 投资者帐号
	AccountID types.TThostFtdcAccountIDType
	// 密码
	Password types.TThostFtdcPasswordType
	// 币种代码
	CurrencyID types.TThostFtdcCurrencyIDType
}

// 交易所行情报盘机
type CThostFtdcMDTraderOfferField struct {
	// 交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	// 交易所交易员代码
	TraderID types.TThostFtdcTraderIDType
	// 会员代码
	ParticipantID types.TThostFtdcParticipantIDType
	// 密码
	Password types.TThostFtdcPasswordType
	// 安装编号
	InstallID types.TThostFtdcInstallIDType
	// 本地报单编号
	OrderLocalID types.TThostFtdcOrderLocalIDType
	// 交易所交易员连接状态
	TraderConnectStatus types.TThostFtdcTraderConnectStatusType
	// 发出连接请求的日期
	ConnectRequestDate types.TThostFtdcDateType
	// 发出连接请求的时间
	ConnectRequestTime types.TThostFtdcTimeType
	// 上次报告日期
	LastReportDate types.TThostFtdcDateType
	// 上次报告时间
	LastReportTime types.TThostFtdcTimeType
	// 完成连接日期
	ConnectDate types.TThostFtdcDateType
	// 完成连接时间
	ConnectTime types.TThostFtdcTimeType
	// 启动日期
	StartDate types.TThostFtdcDateType
	// 启动时间
	StartTime types.TThostFtdcTimeType
	// 交易日
	TradingDay types.TThostFtdcDateType
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 本席位最大成交编号
	MaxTradeID types.TThostFtdcTradeIDType
	// 本席位最大报单备拷
	MaxOrderMessageReference types.TThostFtdcReturnCodeType
}

// 查询行情报盘机
type CThostFtdcQryMDTraderOfferField struct {
	// 交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	// 会员代码
	ParticipantID types.TThostFtdcParticipantIDType
	// 交易所交易员代码
	TraderID types.TThostFtdcTraderIDType
}

// 查询客户通知
type CThostFtdcQryNoticeField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
}

// 客户通知
type CThostFtdcNoticeField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 消息正文
	Content types.TThostFtdcContentType
	// 经纪公司通知内容序列号
	SequenceLabel types.TThostFtdcSequenceLabelType
}

// 用户权限
type CThostFtdcUserRightField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 用户代码
	UserID types.TThostFtdcUserIDType
	// 客户权限类型
	UserRightType types.TThostFtdcUserRightTypeType
	// 是否禁止
	IsForbidden types.TThostFtdcBoolType
}

// 查询结算信息确认域
type CThostFtdcQrySettlementInfoConfirmField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID types.TThostFtdcInvestorIDType
}

// 装载结算信息
type CThostFtdcLoadSettlementInfoField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
}

// 经纪公司可提资金算法表
type CThostFtdcBrokerWithdrawAlgorithmField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 可提资金算法
	WithdrawAlgorithm types.TThostFtdcAlgorithmType
	// 资金使用率
	UsingRatio types.TThostFtdcRatioType
	// 可提是否包含平仓盈利
	IncludeCloseProfit types.TThostFtdcIncludeCloseProfitType
	// 本日无仓且无成交客户是否受可提比例限制
	AllWithoutTrade types.TThostFtdcAllWithoutTradeType
	// 可用是否包含平仓盈利
	AvailIncludeCloseProfit types.TThostFtdcIncludeCloseProfitType
	// 是否启用用户事件
	IsBrokerUserEvent types.TThostFtdcBoolType
	// 币种代码
	CurrencyID types.TThostFtdcCurrencyIDType
	// 货币质押比率
	FundMortgageRatio types.TThostFtdcRatioType
	// 权益算法
	BalanceAlgorithm types.TThostFtdcBalanceAlgorithmType
}

// 资金账户口令变更域
type CThostFtdcTradingAccountPasswordUpdateV1Field struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	// 原来的口令
	OldPassword types.TThostFtdcPasswordType
	// 新的口令
	NewPassword types.TThostFtdcPasswordType
}

// 资金账户口令变更域
type CThostFtdcTradingAccountPasswordUpdateField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 投资者帐号
	AccountID types.TThostFtdcAccountIDType
	// 原来的口令
	OldPassword types.TThostFtdcPasswordType
	// 新的口令
	NewPassword types.TThostFtdcPasswordType
	// 币种代码
	CurrencyID types.TThostFtdcCurrencyIDType
}

// 查询组合合约分腿
type CThostFtdcQryCombinationLegField struct {
	// 组合合约代码
	CombInstrumentID types.TThostFtdcInstrumentIDType
	// 单腿编号
	LegID types.TThostFtdcLegIDType
	// 单腿合约代码
	LegInstrumentID types.TThostFtdcInstrumentIDType
}

// 查询组合合约分腿
type CThostFtdcQrySyncStatusField struct {
	// 交易日
	TradingDay types.TThostFtdcDateType
}

// 组合交易合约的单腿
type CThostFtdcCombinationLegField struct {
	// 组合合约代码
	CombInstrumentID types.TThostFtdcInstrumentIDType
	// 单腿编号
	LegID types.TThostFtdcLegIDType
	// 单腿合约代码
	LegInstrumentID types.TThostFtdcInstrumentIDType
	// 买卖方向
	Direction types.TThostFtdcDirectionType
	// 单腿乘数
	LegMultiple types.TThostFtdcLegMultipleType
	// 派生层数
	ImplyLevel types.TThostFtdcImplyLevelType
}

// 数据同步状态
type CThostFtdcSyncStatusField struct {
	// 交易日
	TradingDay types.TThostFtdcDateType
	// 数据同步状态
	DataSyncStatus types.TThostFtdcDataSyncStatusType
}

// 查询联系人
type CThostFtdcQryLinkManField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID types.TThostFtdcInvestorIDType
}

// 联系人
type CThostFtdcLinkManField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	// 联系人类型
	PersonType types.TThostFtdcPersonTypeType
	// 证件类型
	IdentifiedCardType types.TThostFtdcIdCardTypeType
	// 证件号码
	IdentifiedCardNo types.TThostFtdcIdentifiedCardNoType
	// 名称
	PersonName types.TThostFtdcPartyNameType
	// 联系电话
	Telephone types.TThostFtdcTelephoneType
	// 通讯地址
	Address types.TThostFtdcAddressType
	// 邮政编码
	ZipCode types.TThostFtdcZipCodeType
	// 优先级
	Priority types.TThostFtdcPriorityType
	// 开户邮政编码
	UOAZipCode types.TThostFtdcUOAZipCodeType
	// 全称
	PersonFullName types.TThostFtdcInvestorFullNameType
}

// 查询经纪公司用户事件
type CThostFtdcQryBrokerUserEventField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 用户代码
	UserID types.TThostFtdcUserIDType
	// 用户事件类型
	UserEventType types.TThostFtdcUserEventTypeType
}

// 查询经纪公司用户事件
type CThostFtdcBrokerUserEventField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 用户代码
	UserID types.TThostFtdcUserIDType
	// 用户事件类型
	UserEventType types.TThostFtdcUserEventTypeType
	// 用户事件序号
	EventSequenceNo types.TThostFtdcSequenceNoType
	// 事件发生日期
	EventDate types.TThostFtdcDateType
	// 事件发生时间
	EventTime types.TThostFtdcTimeType
	// 用户事件信息
	UserEventInfo types.TThostFtdcUserEventInfoType
	// 投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	// 合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
}

// 查询签约银行请求
type CThostFtdcQryContractBankField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 银行代码
	BankID types.TThostFtdcBankIDType
	// 银行分中心代码
	BankBrchID types.TThostFtdcBankBrchIDType
}

// 查询签约银行响应
type CThostFtdcContractBankField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 银行代码
	BankID types.TThostFtdcBankIDType
	// 银行分中心代码
	BankBrchID types.TThostFtdcBankBrchIDType
	// 银行名称
	BankName types.TThostFtdcBankNameType
}

// 投资者组合持仓明细
type CThostFtdcInvestorPositionCombineDetailField struct {
	// 交易日
	TradingDay types.TThostFtdcDateType
	// 开仓日期
	OpenDate types.TThostFtdcDateType
	// 交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	// 结算编号
	SettlementID types.TThostFtdcSettlementIDType
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	// 组合编号
	ComTradeID types.TThostFtdcTradeIDType
	// 撮合编号
	TradeID types.TThostFtdcTradeIDType
	// 合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
	// 投机套保标志
	HedgeFlag types.TThostFtdcHedgeFlagType
	// 买卖
	Direction types.TThostFtdcDirectionType
	// 持仓量
	TotalAmt types.TThostFtdcVolumeType
	// 投资者保证金
	Margin types.TThostFtdcMoneyType
	// 交易所保证金
	ExchMargin types.TThostFtdcMoneyType
	// 保证金率
	MarginRateByMoney types.TThostFtdcRatioType
	// 保证金率(按手数)
	MarginRateByVolume types.TThostFtdcRatioType
	// 单腿编号
	LegID types.TThostFtdcLegIDType
	// 单腿乘数
	LegMultiple types.TThostFtdcLegMultipleType
	// 组合持仓合约编码
	CombInstrumentID types.TThostFtdcInstrumentIDType
	// 成交组号
	TradeGroupID types.TThostFtdcTradeGroupIDType
}

// 预埋单
type CThostFtdcParkedOrderField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	// 合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
	// 报单引用
	OrderRef types.TThostFtdcOrderRefType
	// 用户代码
	UserID types.TThostFtdcUserIDType
	// 报单价格条件
	OrderPriceType types.TThostFtdcOrderPriceTypeType
	// 买卖方向
	Direction types.TThostFtdcDirectionType
	// 组合开平标志
	CombOffsetFlag types.TThostFtdcCombOffsetFlagType
	// 组合投机套保标志
	CombHedgeFlag types.TThostFtdcCombHedgeFlagType
	// 价格
	LimitPrice types.TThostFtdcPriceType
	// 数量
	VolumeTotalOriginal types.TThostFtdcVolumeType
	// 有效期类型
	TimeCondition types.TThostFtdcTimeConditionType
	// GTD日期
	GTDDate types.TThostFtdcDateType
	// 成交量类型
	VolumeCondition types.TThostFtdcVolumeConditionType
	// 最小成交量
	MinVolume types.TThostFtdcVolumeType
	// 触发条件
	ContingentCondition types.TThostFtdcContingentConditionType
	// 止损价
	StopPrice types.TThostFtdcPriceType
	// 强平原因
	ForceCloseReason types.TThostFtdcForceCloseReasonType
	// 自动挂起标志
	IsAutoSuspend types.TThostFtdcBoolType
	// 业务单元
	BusinessUnit types.TThostFtdcBusinessUnitType
	// 请求编号
	RequestID types.TThostFtdcRequestIDType
	// 用户强评标志
	UserForceClose types.TThostFtdcBoolType
	// 交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	// 预埋报单编号
	ParkedOrderID types.TThostFtdcParkedOrderIDType
	// 用户类型
	UserType types.TThostFtdcUserTypeType
	// 预埋单状态
	Status types.TThostFtdcParkedOrderStatusType
	// 错误代码
	ErrorID types.TThostFtdcErrorIDType
	// 错误信息
	ErrorMsg types.TThostFtdcErrorMsgType
	// 互换单标志
	IsSwapOrder types.TThostFtdcBoolType
	// 资金账号
	AccountID types.TThostFtdcAccountIDType
	// 币种代码
	CurrencyID types.TThostFtdcCurrencyIDType
	// 交易编码
	ClientID types.TThostFtdcClientIDType
	// 投资单元代码
	InvestUnitID types.TThostFtdcInvestUnitIDType
	// IP地址
	IPAddress types.TThostFtdcIPAddressType
	// Mac地址
	MacAddress types.TThostFtdcMacAddressType
}

// 输入预埋单操作
type CThostFtdcParkedOrderActionField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	// 报单操作引用
	OrderActionRef types.TThostFtdcOrderActionRefType
	// 报单引用
	OrderRef types.TThostFtdcOrderRefType
	// 请求编号
	RequestID types.TThostFtdcRequestIDType
	// 前置编号
	FrontID types.TThostFtdcFrontIDType
	// 会话编号
	SessionID types.TThostFtdcSessionIDType
	// 交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	// 报单编号
	OrderSysID types.TThostFtdcOrderSysIDType
	// 操作标志
	ActionFlag types.TThostFtdcActionFlagType
	// 价格
	LimitPrice types.TThostFtdcPriceType
	// 数量变化
	VolumeChange types.TThostFtdcVolumeType
	// 用户代码
	UserID types.TThostFtdcUserIDType
	// 合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
	// 预埋撤单单编号
	ParkedOrderActionID types.TThostFtdcParkedOrderActionIDType
	// 用户类型
	UserType types.TThostFtdcUserTypeType
	// 预埋撤单状态
	Status types.TThostFtdcParkedOrderStatusType
	// 错误代码
	ErrorID types.TThostFtdcErrorIDType
	// 错误信息
	ErrorMsg types.TThostFtdcErrorMsgType
	// 投资单元代码
	InvestUnitID types.TThostFtdcInvestUnitIDType
	// IP地址
	IPAddress types.TThostFtdcIPAddressType
	// Mac地址
	MacAddress types.TThostFtdcMacAddressType
}

// 查询预埋单
type CThostFtdcQryParkedOrderField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	// 合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
	// 交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
}

// 查询预埋撤单
type CThostFtdcQryParkedOrderActionField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	// 合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
	// 交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
}

// 删除预埋单
type CThostFtdcRemoveParkedOrderField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	// 预埋报单编号
	ParkedOrderID types.TThostFtdcParkedOrderIDType
}

// 删除预埋撤单
type CThostFtdcRemoveParkedOrderActionField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	// 预埋撤单编号
	ParkedOrderActionID types.TThostFtdcParkedOrderActionIDType
}

// 经纪公司可提资金算法表
type CThostFtdcInvestorWithdrawAlgorithmField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 投资者范围
	InvestorRange types.TThostFtdcInvestorRangeType
	// 投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	// 可提资金比例
	UsingRatio types.TThostFtdcRatioType
	// 币种代码
	CurrencyID types.TThostFtdcCurrencyIDType
	// 货币质押比率
	FundMortgageRatio types.TThostFtdcRatioType
}

// 查询组合持仓明细
type CThostFtdcQryInvestorPositionCombineDetailField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	// 组合持仓合约编码
	CombInstrumentID types.TThostFtdcInstrumentIDType
}

// 成交均价
type CThostFtdcMarketDataAveragePriceField struct {
	// 当日均价
	AveragePrice types.TThostFtdcPriceType
}

// 校验投资者密码
type CThostFtdcVerifyInvestorPasswordField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	// 密码
	Password types.TThostFtdcPasswordType
}

// 用户IP
type CThostFtdcUserIPField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 用户代码
	UserID types.TThostFtdcUserIDType
	// IP地址
	IPAddress types.TThostFtdcIPAddressType
	// IP地址掩码
	IPMask types.TThostFtdcIPAddressType
	// Mac地址
	MacAddress types.TThostFtdcMacAddressType
}

// 用户事件通知信息
type CThostFtdcTradingNoticeInfoField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	// 发送时间
	SendTime types.TThostFtdcTimeType
	// 消息正文
	FieldContent types.TThostFtdcContentType
	// 序列系列号
	SequenceSeries types.TThostFtdcSequenceSeriesType
	// 序列号
	SequenceNo types.TThostFtdcSequenceNoType
}

// 用户事件通知
type CThostFtdcTradingNoticeField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 投资者范围
	InvestorRange types.TThostFtdcInvestorRangeType
	// 投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	// 序列系列号
	SequenceSeries types.TThostFtdcSequenceSeriesType
	// 用户代码
	UserID types.TThostFtdcUserIDType
	// 发送时间
	SendTime types.TThostFtdcTimeType
	// 序列号
	SequenceNo types.TThostFtdcSequenceNoType
	// 消息正文
	FieldContent types.TThostFtdcContentType
}

// 查询交易事件通知
type CThostFtdcQryTradingNoticeField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID types.TThostFtdcInvestorIDType
}

// 查询错误报单
type CThostFtdcQryErrOrderField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID types.TThostFtdcInvestorIDType
}

// 错误报单
type CThostFtdcErrOrderField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	// 合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
	// 报单引用
	OrderRef types.TThostFtdcOrderRefType
	// 用户代码
	UserID types.TThostFtdcUserIDType
	// 报单价格条件
	OrderPriceType types.TThostFtdcOrderPriceTypeType
	// 买卖方向
	Direction types.TThostFtdcDirectionType
	// 组合开平标志
	CombOffsetFlag types.TThostFtdcCombOffsetFlagType
	// 组合投机套保标志
	CombHedgeFlag types.TThostFtdcCombHedgeFlagType
	// 价格
	LimitPrice types.TThostFtdcPriceType
	// 数量
	VolumeTotalOriginal types.TThostFtdcVolumeType
	// 有效期类型
	TimeCondition types.TThostFtdcTimeConditionType
	// GTD日期
	GTDDate types.TThostFtdcDateType
	// 成交量类型
	VolumeCondition types.TThostFtdcVolumeConditionType
	// 最小成交量
	MinVolume types.TThostFtdcVolumeType
	// 触发条件
	ContingentCondition types.TThostFtdcContingentConditionType
	// 止损价
	StopPrice types.TThostFtdcPriceType
	// 强平原因
	ForceCloseReason types.TThostFtdcForceCloseReasonType
	// 自动挂起标志
	IsAutoSuspend types.TThostFtdcBoolType
	// 业务单元
	BusinessUnit types.TThostFtdcBusinessUnitType
	// 请求编号
	RequestID types.TThostFtdcRequestIDType
	// 用户强评标志
	UserForceClose types.TThostFtdcBoolType
	// 错误代码
	ErrorID types.TThostFtdcErrorIDType
	// 错误信息
	ErrorMsg types.TThostFtdcErrorMsgType
	// 互换单标志
	IsSwapOrder types.TThostFtdcBoolType
	// 交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	// 投资单元代码
	InvestUnitID types.TThostFtdcInvestUnitIDType
	// 资金账号
	AccountID types.TThostFtdcAccountIDType
	// 币种代码
	CurrencyID types.TThostFtdcCurrencyIDType
	// 交易编码
	ClientID types.TThostFtdcClientIDType
	// IP地址
	IPAddress types.TThostFtdcIPAddressType
	// Mac地址
	MacAddress types.TThostFtdcMacAddressType
}

// 查询错误报单操作
type CThostFtdcErrorConditionalOrderField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	// 合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
	// 报单引用
	OrderRef types.TThostFtdcOrderRefType
	// 用户代码
	UserID types.TThostFtdcUserIDType
	// 报单价格条件
	OrderPriceType types.TThostFtdcOrderPriceTypeType
	// 买卖方向
	Direction types.TThostFtdcDirectionType
	// 组合开平标志
	CombOffsetFlag types.TThostFtdcCombOffsetFlagType
	// 组合投机套保标志
	CombHedgeFlag types.TThostFtdcCombHedgeFlagType
	// 价格
	LimitPrice types.TThostFtdcPriceType
	// 数量
	VolumeTotalOriginal types.TThostFtdcVolumeType
	// 有效期类型
	TimeCondition types.TThostFtdcTimeConditionType
	// GTD日期
	GTDDate types.TThostFtdcDateType
	// 成交量类型
	VolumeCondition types.TThostFtdcVolumeConditionType
	// 最小成交量
	MinVolume types.TThostFtdcVolumeType
	// 触发条件
	ContingentCondition types.TThostFtdcContingentConditionType
	// 止损价
	StopPrice types.TThostFtdcPriceType
	// 强平原因
	ForceCloseReason types.TThostFtdcForceCloseReasonType
	// 自动挂起标志
	IsAutoSuspend types.TThostFtdcBoolType
	// 业务单元
	BusinessUnit types.TThostFtdcBusinessUnitType
	// 请求编号
	RequestID types.TThostFtdcRequestIDType
	// 本地报单编号
	OrderLocalID types.TThostFtdcOrderLocalIDType
	// 交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	// 会员代码
	ParticipantID types.TThostFtdcParticipantIDType
	// 客户代码
	ClientID types.TThostFtdcClientIDType
	// 合约在交易所的代码
	ExchangeInstID types.TThostFtdcExchangeInstIDType
	// 交易所交易员代码
	TraderID types.TThostFtdcTraderIDType
	// 安装编号
	InstallID types.TThostFtdcInstallIDType
	// 报单提交状态
	OrderSubmitStatus types.TThostFtdcOrderSubmitStatusType
	// 报单提示序号
	NotifySequence types.TThostFtdcSequenceNoType
	// 交易日
	TradingDay types.TThostFtdcDateType
	// 结算编号
	SettlementID types.TThostFtdcSettlementIDType
	// 报单编号
	OrderSysID types.TThostFtdcOrderSysIDType
	// 报单来源
	OrderSource types.TThostFtdcOrderSourceType
	// 报单状态
	OrderStatus types.TThostFtdcOrderStatusType
	// 报单类型
	OrderType types.TThostFtdcOrderTypeType
	// 今成交数量
	VolumeTraded types.TThostFtdcVolumeType
	// 剩余数量
	VolumeTotal types.TThostFtdcVolumeType
	// 报单日期
	InsertDate types.TThostFtdcDateType
	// 委托时间
	InsertTime types.TThostFtdcTimeType
	// 激活时间
	ActiveTime types.TThostFtdcTimeType
	// 挂起时间
	SuspendTime types.TThostFtdcTimeType
	// 最后修改时间
	UpdateTime types.TThostFtdcTimeType
	// 撤销时间
	CancelTime types.TThostFtdcTimeType
	// 最后修改交易所交易员代码
	ActiveTraderID types.TThostFtdcTraderIDType
	// 结算会员编号
	ClearingPartID types.TThostFtdcParticipantIDType
	// 序号
	SequenceNo types.TThostFtdcSequenceNoType
	// 前置编号
	FrontID types.TThostFtdcFrontIDType
	// 会话编号
	SessionID types.TThostFtdcSessionIDType
	// 用户端产品信息
	UserProductInfo types.TThostFtdcProductInfoType
	// 状态信息
	StatusMsg types.TThostFtdcErrorMsgType
	// 用户强评标志
	UserForceClose types.TThostFtdcBoolType
	// 操作用户代码
	ActiveUserID types.TThostFtdcUserIDType
	// 经纪公司报单编号
	BrokerOrderSeq types.TThostFtdcSequenceNoType
	// 相关报单
	RelativeOrderSysID types.TThostFtdcOrderSysIDType
	// 郑商所成交数量
	ZCETotalTradedVolume types.TThostFtdcVolumeType
	// 错误代码
	ErrorID types.TThostFtdcErrorIDType
	// 错误信息
	ErrorMsg types.TThostFtdcErrorMsgType
	// 互换单标志
	IsSwapOrder types.TThostFtdcBoolType
	// 营业部编号
	BranchID types.TThostFtdcBranchIDType
	// 投资单元代码
	InvestUnitID types.TThostFtdcInvestUnitIDType
	// 资金账号
	AccountID types.TThostFtdcAccountIDType
	// 币种代码
	CurrencyID types.TThostFtdcCurrencyIDType
	// IP地址
	IPAddress types.TThostFtdcIPAddressType
	// Mac地址
	MacAddress types.TThostFtdcMacAddressType
}

// 查询错误报单操作
type CThostFtdcQryErrOrderActionField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID types.TThostFtdcInvestorIDType
}

// 错误报单操作
type CThostFtdcErrOrderActionField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	// 报单操作引用
	OrderActionRef types.TThostFtdcOrderActionRefType
	// 报单引用
	OrderRef types.TThostFtdcOrderRefType
	// 请求编号
	RequestID types.TThostFtdcRequestIDType
	// 前置编号
	FrontID types.TThostFtdcFrontIDType
	// 会话编号
	SessionID types.TThostFtdcSessionIDType
	// 交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	// 报单编号
	OrderSysID types.TThostFtdcOrderSysIDType
	// 操作标志
	ActionFlag types.TThostFtdcActionFlagType
	// 价格
	LimitPrice types.TThostFtdcPriceType
	// 数量变化
	VolumeChange types.TThostFtdcVolumeType
	// 操作日期
	ActionDate types.TThostFtdcDateType
	// 操作时间
	ActionTime types.TThostFtdcTimeType
	// 交易所交易员代码
	TraderID types.TThostFtdcTraderIDType
	// 安装编号
	InstallID types.TThostFtdcInstallIDType
	// 本地报单编号
	OrderLocalID types.TThostFtdcOrderLocalIDType
	// 操作本地编号
	ActionLocalID types.TThostFtdcOrderLocalIDType
	// 会员代码
	ParticipantID types.TThostFtdcParticipantIDType
	// 客户代码
	ClientID types.TThostFtdcClientIDType
	// 业务单元
	BusinessUnit types.TThostFtdcBusinessUnitType
	// 报单操作状态
	OrderActionStatus types.TThostFtdcOrderActionStatusType
	// 用户代码
	UserID types.TThostFtdcUserIDType
	// 状态信息
	StatusMsg types.TThostFtdcErrorMsgType
	// 合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
	// 营业部编号
	BranchID types.TThostFtdcBranchIDType
	// 投资单元代码
	InvestUnitID types.TThostFtdcInvestUnitIDType
	// IP地址
	IPAddress types.TThostFtdcIPAddressType
	// Mac地址
	MacAddress types.TThostFtdcMacAddressType
	// 错误代码
	ErrorID types.TThostFtdcErrorIDType
	// 错误信息
	ErrorMsg types.TThostFtdcErrorMsgType
}

// 查询交易所状态
type CThostFtdcQryExchangeSequenceField struct {
	// 交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
}

// 交易所状态
type CThostFtdcExchangeSequenceField struct {
	// 交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	// 序号
	SequenceNo types.TThostFtdcSequenceNoType
	// 合约交易状态
	MarketStatus types.TThostFtdcInstrumentStatusType
}

// 根据价格查询最大报单数量
type CThostFtdcQueryMaxOrderVolumeWithPriceField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	// 合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
	// 买卖方向
	Direction types.TThostFtdcDirectionType
	// 开平标志
	OffsetFlag types.TThostFtdcOffsetFlagType
	// 投机套保标志
	HedgeFlag types.TThostFtdcHedgeFlagType
	// 最大允许报单数量
	MaxVolume types.TThostFtdcVolumeType
	// 报单价格
	Price types.TThostFtdcPriceType
}

// 查询经纪公司交易参数
type CThostFtdcQryBrokerTradingParamsField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	// 币种代码
	CurrencyID types.TThostFtdcCurrencyIDType
}

// 经纪公司交易参数
type CThostFtdcBrokerTradingParamsField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	// 保证金价格类型
	MarginPriceType types.TThostFtdcMarginPriceTypeType
	// 盈亏算法
	Algorithm types.TThostFtdcAlgorithmType
	// 可用是否包含平仓盈利
	AvailIncludeCloseProfit types.TThostFtdcIncludeCloseProfitType
	// 币种代码
	CurrencyID types.TThostFtdcCurrencyIDType
	// 期权权利金价格类型
	OptionRoyaltyPriceType types.TThostFtdcOptionRoyaltyPriceTypeType
}

// 查询经纪公司交易算法
type CThostFtdcQryBrokerTradingAlgosField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	// 合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
}

// 经纪公司交易算法
type CThostFtdcBrokerTradingAlgosField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	// 合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
	// 持仓处理算法编号
	HandlePositionAlgoID types.TThostFtdcHandlePositionAlgoIDType
	// 寻找保证金率算法编号
	FindMarginRateAlgoID types.TThostFtdcFindMarginRateAlgoIDType
	// 资金处理算法编号
	HandleTradingAccountAlgoID types.TThostFtdcHandleTradingAccountAlgoIDType
}

// 查询经纪公司资金
type CThostFtdcQueryBrokerDepositField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
}

// 经纪公司资金
type CThostFtdcBrokerDepositField struct {
	// 交易日期
	TradingDay types.TThostFtdcTradeDateType
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 会员代码
	ParticipantID types.TThostFtdcParticipantIDType
	// 交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	// 上次结算准备金
	PreBalance types.TThostFtdcMoneyType
	// 当前保证金总额
	CurrMargin types.TThostFtdcMoneyType
	// 平仓盈亏
	CloseProfit types.TThostFtdcMoneyType
	// 期货结算准备金
	Balance types.TThostFtdcMoneyType
	// 入金金额
	Deposit types.TThostFtdcMoneyType
	// 出金金额
	Withdraw types.TThostFtdcMoneyType
	// 可提资金
	Available types.TThostFtdcMoneyType
	// 基本准备金
	Reserve types.TThostFtdcMoneyType
	// 冻结的保证金
	FrozenMargin types.TThostFtdcMoneyType
}

// 查询保证金监管系统经纪公司密钥
type CThostFtdcQryCFMMCBrokerKeyField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
}

// 保证金监管系统经纪公司密钥
type CThostFtdcCFMMCBrokerKeyField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 经纪公司统一编码
	ParticipantID types.TThostFtdcParticipantIDType
	// 密钥生成日期
	CreateDate types.TThostFtdcDateType
	// 密钥生成时间
	CreateTime types.TThostFtdcTimeType
	// 密钥编号
	KeyID types.TThostFtdcSequenceNoType
	// 动态密钥
	CurrentKey types.TThostFtdcCFMMCKeyType
	// 动态密钥类型
	KeyKind types.TThostFtdcCFMMCKeyKindType
}

// 保证金监管系统经纪公司资金账户密钥
type CThostFtdcCFMMCTradingAccountKeyField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 经纪公司统一编码
	ParticipantID types.TThostFtdcParticipantIDType
	// 投资者帐号
	AccountID types.TThostFtdcAccountIDType
	// 密钥编号
	KeyID types.TThostFtdcSequenceNoType
	// 动态密钥
	CurrentKey types.TThostFtdcCFMMCKeyType
}

// 请求查询保证金监管系统经纪公司资金账户密钥
type CThostFtdcQryCFMMCTradingAccountKeyField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID types.TThostFtdcInvestorIDType
}

// 用户动态令牌参数
type CThostFtdcBrokerUserOTPParamField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 用户代码
	UserID types.TThostFtdcUserIDType
	// 动态令牌提供商
	OTPVendorsID types.TThostFtdcOTPVendorsIDType
	// 动态令牌序列号
	SerialNumber types.TThostFtdcSerialNumberType
	// 令牌密钥
	AuthKey types.TThostFtdcAuthKeyType
	// 漂移值
	LastDrift types.TThostFtdcLastDriftType
	// 成功值
	LastSuccess types.TThostFtdcLastSuccessType
	// 动态令牌类型
	OTPType types.TThostFtdcOTPTypeType
}

// 手工同步用户动态令牌
type CThostFtdcManualSyncBrokerUserOTPField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 用户代码
	UserID types.TThostFtdcUserIDType
	// 动态令牌类型
	OTPType types.TThostFtdcOTPTypeType
	// 第一个动态密码
	FirstOTP types.TThostFtdcPasswordType
	// 第二个动态密码
	SecondOTP types.TThostFtdcPasswordType
}

// 投资者手续费率模板
type CThostFtdcCommRateModelField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 手续费率模板代码
	CommModelID types.TThostFtdcInvestorIDType
	// 模板名称
	CommModelName types.TThostFtdcCommModelNameType
}

// 请求查询投资者手续费率模板
type CThostFtdcQryCommRateModelField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 手续费率模板代码
	CommModelID types.TThostFtdcInvestorIDType
}

// 投资者保证金率模板
type CThostFtdcMarginModelField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 保证金率模板代码
	MarginModelID types.TThostFtdcInvestorIDType
	// 模板名称
	MarginModelName types.TThostFtdcCommModelNameType
}

// 请求查询投资者保证金率模板
type CThostFtdcQryMarginModelField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 保证金率模板代码
	MarginModelID types.TThostFtdcInvestorIDType
}

// 仓单折抵信息
type CThostFtdcEWarrantOffsetField struct {
	// 交易日期
	TradingDay types.TThostFtdcTradeDateType
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	// 交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	// 合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
	// 买卖方向
	Direction types.TThostFtdcDirectionType
	// 投机套保标志
	HedgeFlag types.TThostFtdcHedgeFlagType
	// 数量
	Volume types.TThostFtdcVolumeType
}

// 查询仓单折抵信息
type CThostFtdcQryEWarrantOffsetField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	// 交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	// 合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
}

// 查询投资者品种跨品种保证金
type CThostFtdcQryInvestorProductGroupMarginField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	// 品种跨品种标示
	ProductGroupID types.TThostFtdcInstrumentIDType
	// 投机套保标志
	HedgeFlag types.TThostFtdcHedgeFlagType
}

// 投资者品种跨品种保证金
type CThostFtdcInvestorProductGroupMarginField struct {
	// 品种跨品种标示
	ProductGroupID types.TThostFtdcInstrumentIDType
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	// 交易日
	TradingDay types.TThostFtdcDateType
	// 结算编号
	SettlementID types.TThostFtdcSettlementIDType
	// 冻结的保证金
	FrozenMargin types.TThostFtdcMoneyType
	// 多头冻结的保证金
	LongFrozenMargin types.TThostFtdcMoneyType
	// 空头冻结的保证金
	ShortFrozenMargin types.TThostFtdcMoneyType
	// 占用的保证金
	UseMargin types.TThostFtdcMoneyType
	// 多头保证金
	LongUseMargin types.TThostFtdcMoneyType
	// 空头保证金
	ShortUseMargin types.TThostFtdcMoneyType
	// 交易所保证金
	ExchMargin types.TThostFtdcMoneyType
	// 交易所多头保证金
	LongExchMargin types.TThostFtdcMoneyType
	// 交易所空头保证金
	ShortExchMargin types.TThostFtdcMoneyType
	// 平仓盈亏
	CloseProfit types.TThostFtdcMoneyType
	// 冻结的手续费
	FrozenCommission types.TThostFtdcMoneyType
	// 手续费
	Commission types.TThostFtdcMoneyType
	// 冻结的资金
	FrozenCash types.TThostFtdcMoneyType
	// 资金差额
	CashIn types.TThostFtdcMoneyType
	// 持仓盈亏
	PositionProfit types.TThostFtdcMoneyType
	// 折抵总金额
	OffsetAmount types.TThostFtdcMoneyType
	// 多头折抵总金额
	LongOffsetAmount types.TThostFtdcMoneyType
	// 空头折抵总金额
	ShortOffsetAmount types.TThostFtdcMoneyType
	// 交易所折抵总金额
	ExchOffsetAmount types.TThostFtdcMoneyType
	// 交易所多头折抵总金额
	LongExchOffsetAmount types.TThostFtdcMoneyType
	// 交易所空头折抵总金额
	ShortExchOffsetAmount types.TThostFtdcMoneyType
	// 投机套保标志
	HedgeFlag types.TThostFtdcHedgeFlagType
}

// 查询监控中心用户令牌
type CThostFtdcQueryCFMMCTradingAccountTokenField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID types.TThostFtdcInvestorIDType
}

// 监控中心用户令牌
type CThostFtdcCFMMCTradingAccountTokenField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 经纪公司统一编码
	ParticipantID types.TThostFtdcParticipantIDType
	// 投资者帐号
	AccountID types.TThostFtdcAccountIDType
	// 密钥编号
	KeyID types.TThostFtdcSequenceNoType
	// 动态令牌
	Token types.TThostFtdcCFMMCTokenType
}

// 查询产品组
type CThostFtdcQryProductGroupField struct {
	// 产品代码
	ProductID types.TThostFtdcInstrumentIDType
	// 交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
}

// 投资者品种跨品种保证金产品组
type CThostFtdcProductGroupField struct {
	// 产品代码
	ProductID types.TThostFtdcInstrumentIDType
	// 交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	// 产品组代码
	ProductGroupID types.TThostFtdcInstrumentIDType
}

// 交易所公告
type CThostFtdcBulletinField struct {
	// 交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	// 交易日
	TradingDay types.TThostFtdcDateType
	// 公告编号
	BulletinID types.TThostFtdcBulletinIDType
	// 序列号
	SequenceNo types.TThostFtdcSequenceNoType
	// 公告类型
	NewsType types.TThostFtdcNewsTypeType
	// 紧急程度
	NewsUrgency types.TThostFtdcNewsUrgencyType
	// 发送时间
	SendTime types.TThostFtdcTimeType
	// 消息摘要
	Abstract types.TThostFtdcAbstractType
	// 消息来源
	ComeFrom types.TThostFtdcComeFromType
	// 消息正文
	Content types.TThostFtdcContentType
	// WEB地址
	URLLink types.TThostFtdcURLLinkType
	// 市场代码
	MarketID types.TThostFtdcMarketIDType
}

// 查询交易所公告
type CThostFtdcQryBulletinField struct {
	// 交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	// 公告编号
	BulletinID types.TThostFtdcBulletinIDType
	// 序列号
	SequenceNo types.TThostFtdcSequenceNoType
	// 公告类型
	NewsType types.TThostFtdcNewsTypeType
	// 紧急程度
	NewsUrgency types.TThostFtdcNewsUrgencyType
}

// 转帐开户请求
type CThostFtdcReqOpenAccountField struct {
	// 业务功能码
	TradeCode types.TThostFtdcTradeCodeType
	// 银行代码
	BankID types.TThostFtdcBankIDType
	// 银行分支机构代码
	BankBranchID types.TThostFtdcBankBrchIDType
	// 期商代码
	BrokerID types.TThostFtdcBrokerIDType
	// 期商分支机构代码
	BrokerBranchID types.TThostFtdcFutureBranchIDType
	// 交易日期
	TradeDate types.TThostFtdcTradeDateType
	// 交易时间
	TradeTime types.TThostFtdcTradeTimeType
	// 银行流水号
	BankSerial types.TThostFtdcBankSerialType
	// 交易系统日期
	TradingDay types.TThostFtdcTradeDateType
	// 银期平台消息流水号
	PlateSerial types.TThostFtdcSerialType
	// 最后分片标志
	LastFragment types.TThostFtdcLastFragmentType
	// 会话号
	SessionID types.TThostFtdcSessionIDType
	// 客户姓名
	CustomerName types.TThostFtdcIndividualNameType
	// 证件类型
	IdCardType types.TThostFtdcIdCardTypeType
	// 证件号码
	IdentifiedCardNo types.TThostFtdcIdentifiedCardNoType
	// 性别
	Gender types.TThostFtdcGenderType
	// 国家代码
	CountryCode types.TThostFtdcCountryCodeType
	// 客户类型
	CustType types.TThostFtdcCustTypeType
	// 地址
	Address types.TThostFtdcAddressType
	// 邮编
	ZipCode types.TThostFtdcZipCodeType
	// 电话号码
	Telephone types.TThostFtdcTelephoneType
	// 手机
	MobilePhone types.TThostFtdcMobilePhoneType
	// 传真
	Fax types.TThostFtdcFaxType
	// 电子邮件
	EMail types.TThostFtdcEMailType
	// 资金账户状态
	MoneyAccountStatus types.TThostFtdcMoneyAccountStatusType
	// 银行帐号
	BankAccount types.TThostFtdcBankAccountType
	// 银行密码
	BankPassWord types.TThostFtdcPasswordType
	// 投资者帐号
	AccountID types.TThostFtdcAccountIDType
	// 期货密码
	Password types.TThostFtdcPasswordType
	// 安装编号
	InstallID types.TThostFtdcInstallIDType
	// 验证客户证件号码标志
	VerifyCertNoFlag types.TThostFtdcYesNoIndicatorType
	// 币种代码
	CurrencyID types.TThostFtdcCurrencyIDType
	// 汇钞标志
	CashExchangeCode types.TThostFtdcCashExchangeCodeType
	// 摘要
	Digest types.TThostFtdcDigestType
	// 银行帐号类型
	BankAccType types.TThostFtdcBankAccTypeType
	// 渠道标志
	DeviceID types.TThostFtdcDeviceIDType
	// 期货单位帐号类型
	BankSecuAccType types.TThostFtdcBankAccTypeType
	// 期货公司银行编码
	BrokerIDByBank types.TThostFtdcBankCodingForFutureType
	// 期货单位帐号
	BankSecuAcc types.TThostFtdcBankAccountType
	// 银行密码标志
	BankPwdFlag types.TThostFtdcPwdFlagType
	// 期货资金密码核对标志
	SecuPwdFlag types.TThostFtdcPwdFlagType
	// 交易柜员
	OperNo types.TThostFtdcOperNoType
	// 交易ID
	TID types.TThostFtdcTIDType
	// 用户标识
	UserID types.TThostFtdcUserIDType
}

// 转帐销户请求
type CThostFtdcReqCancelAccountField struct {
	// 业务功能码
	TradeCode types.TThostFtdcTradeCodeType
	// 银行代码
	BankID types.TThostFtdcBankIDType
	// 银行分支机构代码
	BankBranchID types.TThostFtdcBankBrchIDType
	// 期商代码
	BrokerID types.TThostFtdcBrokerIDType
	// 期商分支机构代码
	BrokerBranchID types.TThostFtdcFutureBranchIDType
	// 交易日期
	TradeDate types.TThostFtdcTradeDateType
	// 交易时间
	TradeTime types.TThostFtdcTradeTimeType
	// 银行流水号
	BankSerial types.TThostFtdcBankSerialType
	// 交易系统日期
	TradingDay types.TThostFtdcTradeDateType
	// 银期平台消息流水号
	PlateSerial types.TThostFtdcSerialType
	// 最后分片标志
	LastFragment types.TThostFtdcLastFragmentType
	// 会话号
	SessionID types.TThostFtdcSessionIDType
	// 客户姓名
	CustomerName types.TThostFtdcIndividualNameType
	// 证件类型
	IdCardType types.TThostFtdcIdCardTypeType
	// 证件号码
	IdentifiedCardNo types.TThostFtdcIdentifiedCardNoType
	// 性别
	Gender types.TThostFtdcGenderType
	// 国家代码
	CountryCode types.TThostFtdcCountryCodeType
	// 客户类型
	CustType types.TThostFtdcCustTypeType
	// 地址
	Address types.TThostFtdcAddressType
	// 邮编
	ZipCode types.TThostFtdcZipCodeType
	// 电话号码
	Telephone types.TThostFtdcTelephoneType
	// 手机
	MobilePhone types.TThostFtdcMobilePhoneType
	// 传真
	Fax types.TThostFtdcFaxType
	// 电子邮件
	EMail types.TThostFtdcEMailType
	// 资金账户状态
	MoneyAccountStatus types.TThostFtdcMoneyAccountStatusType
	// 银行帐号
	BankAccount types.TThostFtdcBankAccountType
	// 银行密码
	BankPassWord types.TThostFtdcPasswordType
	// 投资者帐号
	AccountID types.TThostFtdcAccountIDType
	// 期货密码
	Password types.TThostFtdcPasswordType
	// 安装编号
	InstallID types.TThostFtdcInstallIDType
	// 验证客户证件号码标志
	VerifyCertNoFlag types.TThostFtdcYesNoIndicatorType
	// 币种代码
	CurrencyID types.TThostFtdcCurrencyIDType
	// 汇钞标志
	CashExchangeCode types.TThostFtdcCashExchangeCodeType
	// 摘要
	Digest types.TThostFtdcDigestType
	// 银行帐号类型
	BankAccType types.TThostFtdcBankAccTypeType
	// 渠道标志
	DeviceID types.TThostFtdcDeviceIDType
	// 期货单位帐号类型
	BankSecuAccType types.TThostFtdcBankAccTypeType
	// 期货公司银行编码
	BrokerIDByBank types.TThostFtdcBankCodingForFutureType
	// 期货单位帐号
	BankSecuAcc types.TThostFtdcBankAccountType
	// 银行密码标志
	BankPwdFlag types.TThostFtdcPwdFlagType
	// 期货资金密码核对标志
	SecuPwdFlag types.TThostFtdcPwdFlagType
	// 交易柜员
	OperNo types.TThostFtdcOperNoType
	// 交易ID
	TID types.TThostFtdcTIDType
	// 用户标识
	UserID types.TThostFtdcUserIDType
}

// 变更银行账户请求
type CThostFtdcReqChangeAccountField struct {
	// 业务功能码
	TradeCode types.TThostFtdcTradeCodeType
	// 银行代码
	BankID types.TThostFtdcBankIDType
	// 银行分支机构代码
	BankBranchID types.TThostFtdcBankBrchIDType
	// 期商代码
	BrokerID types.TThostFtdcBrokerIDType
	// 期商分支机构代码
	BrokerBranchID types.TThostFtdcFutureBranchIDType
	// 交易日期
	TradeDate types.TThostFtdcTradeDateType
	// 交易时间
	TradeTime types.TThostFtdcTradeTimeType
	// 银行流水号
	BankSerial types.TThostFtdcBankSerialType
	// 交易系统日期
	TradingDay types.TThostFtdcTradeDateType
	// 银期平台消息流水号
	PlateSerial types.TThostFtdcSerialType
	// 最后分片标志
	LastFragment types.TThostFtdcLastFragmentType
	// 会话号
	SessionID types.TThostFtdcSessionIDType
	// 客户姓名
	CustomerName types.TThostFtdcIndividualNameType
	// 证件类型
	IdCardType types.TThostFtdcIdCardTypeType
	// 证件号码
	IdentifiedCardNo types.TThostFtdcIdentifiedCardNoType
	// 性别
	Gender types.TThostFtdcGenderType
	// 国家代码
	CountryCode types.TThostFtdcCountryCodeType
	// 客户类型
	CustType types.TThostFtdcCustTypeType
	// 地址
	Address types.TThostFtdcAddressType
	// 邮编
	ZipCode types.TThostFtdcZipCodeType
	// 电话号码
	Telephone types.TThostFtdcTelephoneType
	// 手机
	MobilePhone types.TThostFtdcMobilePhoneType
	// 传真
	Fax types.TThostFtdcFaxType
	// 电子邮件
	EMail types.TThostFtdcEMailType
	// 资金账户状态
	MoneyAccountStatus types.TThostFtdcMoneyAccountStatusType
	// 银行帐号
	BankAccount types.TThostFtdcBankAccountType
	// 银行密码
	BankPassWord types.TThostFtdcPasswordType
	// 新银行帐号
	NewBankAccount types.TThostFtdcBankAccountType
	// 新银行密码
	NewBankPassWord types.TThostFtdcPasswordType
	// 投资者帐号
	AccountID types.TThostFtdcAccountIDType
	// 期货密码
	Password types.TThostFtdcPasswordType
	// 银行帐号类型
	BankAccType types.TThostFtdcBankAccTypeType
	// 安装编号
	InstallID types.TThostFtdcInstallIDType
	// 验证客户证件号码标志
	VerifyCertNoFlag types.TThostFtdcYesNoIndicatorType
	// 币种代码
	CurrencyID types.TThostFtdcCurrencyIDType
	// 期货公司银行编码
	BrokerIDByBank types.TThostFtdcBankCodingForFutureType
	// 银行密码标志
	BankPwdFlag types.TThostFtdcPwdFlagType
	// 期货资金密码核对标志
	SecuPwdFlag types.TThostFtdcPwdFlagType
	// 交易ID
	TID types.TThostFtdcTIDType
	// 摘要
	Digest types.TThostFtdcDigestType
}

// 转账请求
type CThostFtdcReqTransferField struct {
	// 业务功能码
	TradeCode types.TThostFtdcTradeCodeType
	// 银行代码
	BankID types.TThostFtdcBankIDType
	// 银行分支机构代码
	BankBranchID types.TThostFtdcBankBrchIDType
	// 期商代码
	BrokerID types.TThostFtdcBrokerIDType
	// 期商分支机构代码
	BrokerBranchID types.TThostFtdcFutureBranchIDType
	// 交易日期
	TradeDate types.TThostFtdcTradeDateType
	// 交易时间
	TradeTime types.TThostFtdcTradeTimeType
	// 银行流水号
	BankSerial types.TThostFtdcBankSerialType
	// 交易系统日期
	TradingDay types.TThostFtdcTradeDateType
	// 银期平台消息流水号
	PlateSerial types.TThostFtdcSerialType
	// 最后分片标志
	LastFragment types.TThostFtdcLastFragmentType
	// 会话号
	SessionID types.TThostFtdcSessionIDType
	// 客户姓名
	CustomerName types.TThostFtdcIndividualNameType
	// 证件类型
	IdCardType types.TThostFtdcIdCardTypeType
	// 证件号码
	IdentifiedCardNo types.TThostFtdcIdentifiedCardNoType
	// 客户类型
	CustType types.TThostFtdcCustTypeType
	// 银行帐号
	BankAccount types.TThostFtdcBankAccountType
	// 银行密码
	BankPassWord types.TThostFtdcPasswordType
	// 投资者帐号
	AccountID types.TThostFtdcAccountIDType
	// 期货密码
	Password types.TThostFtdcPasswordType
	// 安装编号
	InstallID types.TThostFtdcInstallIDType
	// 期货公司流水号
	FutureSerial types.TThostFtdcFutureSerialType
	// 用户标识
	UserID types.TThostFtdcUserIDType
	// 验证客户证件号码标志
	VerifyCertNoFlag types.TThostFtdcYesNoIndicatorType
	// 币种代码
	CurrencyID types.TThostFtdcCurrencyIDType
	// 转帐金额
	TradeAmount types.TThostFtdcTradeAmountType
	// 期货可取金额
	FutureFetchAmount types.TThostFtdcTradeAmountType
	// 费用支付标志
	FeePayFlag types.TThostFtdcFeePayFlagType
	// 应收客户费用
	CustFee types.TThostFtdcCustFeeType
	// 应收期货公司费用
	BrokerFee types.TThostFtdcFutureFeeType
	// 发送方给接收方的消息
	Message types.TThostFtdcAddInfoType
	// 摘要
	Digest types.TThostFtdcDigestType
	// 银行帐号类型
	BankAccType types.TThostFtdcBankAccTypeType
	// 渠道标志
	DeviceID types.TThostFtdcDeviceIDType
	// 期货单位帐号类型
	BankSecuAccType types.TThostFtdcBankAccTypeType
	// 期货公司银行编码
	BrokerIDByBank types.TThostFtdcBankCodingForFutureType
	// 期货单位帐号
	BankSecuAcc types.TThostFtdcBankAccountType
	// 银行密码标志
	BankPwdFlag types.TThostFtdcPwdFlagType
	// 期货资金密码核对标志
	SecuPwdFlag types.TThostFtdcPwdFlagType
	// 交易柜员
	OperNo types.TThostFtdcOperNoType
	// 请求编号
	RequestID types.TThostFtdcRequestIDType
	// 交易ID
	TID types.TThostFtdcTIDType
	// 转账交易状态
	TransferStatus types.TThostFtdcTransferStatusType
}

// 银行发起银行资金转期货响应
type CThostFtdcRspTransferField struct {
	// 业务功能码
	TradeCode types.TThostFtdcTradeCodeType
	// 银行代码
	BankID types.TThostFtdcBankIDType
	// 银行分支机构代码
	BankBranchID types.TThostFtdcBankBrchIDType
	// 期商代码
	BrokerID types.TThostFtdcBrokerIDType
	// 期商分支机构代码
	BrokerBranchID types.TThostFtdcFutureBranchIDType
	// 交易日期
	TradeDate types.TThostFtdcTradeDateType
	// 交易时间
	TradeTime types.TThostFtdcTradeTimeType
	// 银行流水号
	BankSerial types.TThostFtdcBankSerialType
	// 交易系统日期
	TradingDay types.TThostFtdcTradeDateType
	// 银期平台消息流水号
	PlateSerial types.TThostFtdcSerialType
	// 最后分片标志
	LastFragment types.TThostFtdcLastFragmentType
	// 会话号
	SessionID types.TThostFtdcSessionIDType
	// 客户姓名
	CustomerName types.TThostFtdcIndividualNameType
	// 证件类型
	IdCardType types.TThostFtdcIdCardTypeType
	// 证件号码
	IdentifiedCardNo types.TThostFtdcIdentifiedCardNoType
	// 客户类型
	CustType types.TThostFtdcCustTypeType
	// 银行帐号
	BankAccount types.TThostFtdcBankAccountType
	// 银行密码
	BankPassWord types.TThostFtdcPasswordType
	// 投资者帐号
	AccountID types.TThostFtdcAccountIDType
	// 期货密码
	Password types.TThostFtdcPasswordType
	// 安装编号
	InstallID types.TThostFtdcInstallIDType
	// 期货公司流水号
	FutureSerial types.TThostFtdcFutureSerialType
	// 用户标识
	UserID types.TThostFtdcUserIDType
	// 验证客户证件号码标志
	VerifyCertNoFlag types.TThostFtdcYesNoIndicatorType
	// 币种代码
	CurrencyID types.TThostFtdcCurrencyIDType
	// 转帐金额
	TradeAmount types.TThostFtdcTradeAmountType
	// 期货可取金额
	FutureFetchAmount types.TThostFtdcTradeAmountType
	// 费用支付标志
	FeePayFlag types.TThostFtdcFeePayFlagType
	// 应收客户费用
	CustFee types.TThostFtdcCustFeeType
	// 应收期货公司费用
	BrokerFee types.TThostFtdcFutureFeeType
	// 发送方给接收方的消息
	Message types.TThostFtdcAddInfoType
	// 摘要
	Digest types.TThostFtdcDigestType
	// 银行帐号类型
	BankAccType types.TThostFtdcBankAccTypeType
	// 渠道标志
	DeviceID types.TThostFtdcDeviceIDType
	// 期货单位帐号类型
	BankSecuAccType types.TThostFtdcBankAccTypeType
	// 期货公司银行编码
	BrokerIDByBank types.TThostFtdcBankCodingForFutureType
	// 期货单位帐号
	BankSecuAcc types.TThostFtdcBankAccountType
	// 银行密码标志
	BankPwdFlag types.TThostFtdcPwdFlagType
	// 期货资金密码核对标志
	SecuPwdFlag types.TThostFtdcPwdFlagType
	// 交易柜员
	OperNo types.TThostFtdcOperNoType
	// 请求编号
	RequestID types.TThostFtdcRequestIDType
	// 交易ID
	TID types.TThostFtdcTIDType
	// 转账交易状态
	TransferStatus types.TThostFtdcTransferStatusType
	// 错误代码
	ErrorID types.TThostFtdcErrorIDType
	// 错误信息
	ErrorMsg types.TThostFtdcErrorMsgType
}

// 冲正请求
type CThostFtdcReqRepealField struct {
	// 冲正时间间隔
	RepealTimeInterval types.TThostFtdcRepealTimeIntervalType
	// 已经冲正次数
	RepealedTimes types.TThostFtdcRepealedTimesType
	// 银行冲正标志
	BankRepealFlag types.TThostFtdcBankRepealFlagType
	// 期商冲正标志
	BrokerRepealFlag types.TThostFtdcBrokerRepealFlagType
	// 被冲正平台流水号
	PlateRepealSerial types.TThostFtdcPlateSerialType
	// 被冲正银行流水号
	BankRepealSerial types.TThostFtdcBankSerialType
	// 被冲正期货流水号
	FutureRepealSerial types.TThostFtdcFutureSerialType
	// 业务功能码
	TradeCode types.TThostFtdcTradeCodeType
	// 银行代码
	BankID types.TThostFtdcBankIDType
	// 银行分支机构代码
	BankBranchID types.TThostFtdcBankBrchIDType
	// 期商代码
	BrokerID types.TThostFtdcBrokerIDType
	// 期商分支机构代码
	BrokerBranchID types.TThostFtdcFutureBranchIDType
	// 交易日期
	TradeDate types.TThostFtdcTradeDateType
	// 交易时间
	TradeTime types.TThostFtdcTradeTimeType
	// 银行流水号
	BankSerial types.TThostFtdcBankSerialType
	// 交易系统日期
	TradingDay types.TThostFtdcTradeDateType
	// 银期平台消息流水号
	PlateSerial types.TThostFtdcSerialType
	// 最后分片标志
	LastFragment types.TThostFtdcLastFragmentType
	// 会话号
	SessionID types.TThostFtdcSessionIDType
	// 客户姓名
	CustomerName types.TThostFtdcIndividualNameType
	// 证件类型
	IdCardType types.TThostFtdcIdCardTypeType
	// 证件号码
	IdentifiedCardNo types.TThostFtdcIdentifiedCardNoType
	// 客户类型
	CustType types.TThostFtdcCustTypeType
	// 银行帐号
	BankAccount types.TThostFtdcBankAccountType
	// 银行密码
	BankPassWord types.TThostFtdcPasswordType
	// 投资者帐号
	AccountID types.TThostFtdcAccountIDType
	// 期货密码
	Password types.TThostFtdcPasswordType
	// 安装编号
	InstallID types.TThostFtdcInstallIDType
	// 期货公司流水号
	FutureSerial types.TThostFtdcFutureSerialType
	// 用户标识
	UserID types.TThostFtdcUserIDType
	// 验证客户证件号码标志
	VerifyCertNoFlag types.TThostFtdcYesNoIndicatorType
	// 币种代码
	CurrencyID types.TThostFtdcCurrencyIDType
	// 转帐金额
	TradeAmount types.TThostFtdcTradeAmountType
	// 期货可取金额
	FutureFetchAmount types.TThostFtdcTradeAmountType
	// 费用支付标志
	FeePayFlag types.TThostFtdcFeePayFlagType
	// 应收客户费用
	CustFee types.TThostFtdcCustFeeType
	// 应收期货公司费用
	BrokerFee types.TThostFtdcFutureFeeType
	// 发送方给接收方的消息
	Message types.TThostFtdcAddInfoType
	// 摘要
	Digest types.TThostFtdcDigestType
	// 银行帐号类型
	BankAccType types.TThostFtdcBankAccTypeType
	// 渠道标志
	DeviceID types.TThostFtdcDeviceIDType
	// 期货单位帐号类型
	BankSecuAccType types.TThostFtdcBankAccTypeType
	// 期货公司银行编码
	BrokerIDByBank types.TThostFtdcBankCodingForFutureType
	// 期货单位帐号
	BankSecuAcc types.TThostFtdcBankAccountType
	// 银行密码标志
	BankPwdFlag types.TThostFtdcPwdFlagType
	// 期货资金密码核对标志
	SecuPwdFlag types.TThostFtdcPwdFlagType
	// 交易柜员
	OperNo types.TThostFtdcOperNoType
	// 请求编号
	RequestID types.TThostFtdcRequestIDType
	// 交易ID
	TID types.TThostFtdcTIDType
	// 转账交易状态
	TransferStatus types.TThostFtdcTransferStatusType
}

// 冲正响应
type CThostFtdcRspRepealField struct {
	// 冲正时间间隔
	RepealTimeInterval types.TThostFtdcRepealTimeIntervalType
	// 已经冲正次数
	RepealedTimes types.TThostFtdcRepealedTimesType
	// 银行冲正标志
	BankRepealFlag types.TThostFtdcBankRepealFlagType
	// 期商冲正标志
	BrokerRepealFlag types.TThostFtdcBrokerRepealFlagType
	// 被冲正平台流水号
	PlateRepealSerial types.TThostFtdcPlateSerialType
	// 被冲正银行流水号
	BankRepealSerial types.TThostFtdcBankSerialType
	// 被冲正期货流水号
	FutureRepealSerial types.TThostFtdcFutureSerialType
	// 业务功能码
	TradeCode types.TThostFtdcTradeCodeType
	// 银行代码
	BankID types.TThostFtdcBankIDType
	// 银行分支机构代码
	BankBranchID types.TThostFtdcBankBrchIDType
	// 期商代码
	BrokerID types.TThostFtdcBrokerIDType
	// 期商分支机构代码
	BrokerBranchID types.TThostFtdcFutureBranchIDType
	// 交易日期
	TradeDate types.TThostFtdcTradeDateType
	// 交易时间
	TradeTime types.TThostFtdcTradeTimeType
	// 银行流水号
	BankSerial types.TThostFtdcBankSerialType
	// 交易系统日期
	TradingDay types.TThostFtdcTradeDateType
	// 银期平台消息流水号
	PlateSerial types.TThostFtdcSerialType
	// 最后分片标志
	LastFragment types.TThostFtdcLastFragmentType
	// 会话号
	SessionID types.TThostFtdcSessionIDType
	// 客户姓名
	CustomerName types.TThostFtdcIndividualNameType
	// 证件类型
	IdCardType types.TThostFtdcIdCardTypeType
	// 证件号码
	IdentifiedCardNo types.TThostFtdcIdentifiedCardNoType
	// 客户类型
	CustType types.TThostFtdcCustTypeType
	// 银行帐号
	BankAccount types.TThostFtdcBankAccountType
	// 银行密码
	BankPassWord types.TThostFtdcPasswordType
	// 投资者帐号
	AccountID types.TThostFtdcAccountIDType
	// 期货密码
	Password types.TThostFtdcPasswordType
	// 安装编号
	InstallID types.TThostFtdcInstallIDType
	// 期货公司流水号
	FutureSerial types.TThostFtdcFutureSerialType
	// 用户标识
	UserID types.TThostFtdcUserIDType
	// 验证客户证件号码标志
	VerifyCertNoFlag types.TThostFtdcYesNoIndicatorType
	// 币种代码
	CurrencyID types.TThostFtdcCurrencyIDType
	// 转帐金额
	TradeAmount types.TThostFtdcTradeAmountType
	// 期货可取金额
	FutureFetchAmount types.TThostFtdcTradeAmountType
	// 费用支付标志
	FeePayFlag types.TThostFtdcFeePayFlagType
	// 应收客户费用
	CustFee types.TThostFtdcCustFeeType
	// 应收期货公司费用
	BrokerFee types.TThostFtdcFutureFeeType
	// 发送方给接收方的消息
	Message types.TThostFtdcAddInfoType
	// 摘要
	Digest types.TThostFtdcDigestType
	// 银行帐号类型
	BankAccType types.TThostFtdcBankAccTypeType
	// 渠道标志
	DeviceID types.TThostFtdcDeviceIDType
	// 期货单位帐号类型
	BankSecuAccType types.TThostFtdcBankAccTypeType
	// 期货公司银行编码
	BrokerIDByBank types.TThostFtdcBankCodingForFutureType
	// 期货单位帐号
	BankSecuAcc types.TThostFtdcBankAccountType
	// 银行密码标志
	BankPwdFlag types.TThostFtdcPwdFlagType
	// 期货资金密码核对标志
	SecuPwdFlag types.TThostFtdcPwdFlagType
	// 交易柜员
	OperNo types.TThostFtdcOperNoType
	// 请求编号
	RequestID types.TThostFtdcRequestIDType
	// 交易ID
	TID types.TThostFtdcTIDType
	// 转账交易状态
	TransferStatus types.TThostFtdcTransferStatusType
	// 错误代码
	ErrorID types.TThostFtdcErrorIDType
	// 错误信息
	ErrorMsg types.TThostFtdcErrorMsgType
}

// 查询账户信息请求
type CThostFtdcReqQueryAccountField struct {
	// 业务功能码
	TradeCode types.TThostFtdcTradeCodeType
	// 银行代码
	BankID types.TThostFtdcBankIDType
	// 银行分支机构代码
	BankBranchID types.TThostFtdcBankBrchIDType
	// 期商代码
	BrokerID types.TThostFtdcBrokerIDType
	// 期商分支机构代码
	BrokerBranchID types.TThostFtdcFutureBranchIDType
	// 交易日期
	TradeDate types.TThostFtdcTradeDateType
	// 交易时间
	TradeTime types.TThostFtdcTradeTimeType
	// 银行流水号
	BankSerial types.TThostFtdcBankSerialType
	// 交易系统日期
	TradingDay types.TThostFtdcTradeDateType
	// 银期平台消息流水号
	PlateSerial types.TThostFtdcSerialType
	// 最后分片标志
	LastFragment types.TThostFtdcLastFragmentType
	// 会话号
	SessionID types.TThostFtdcSessionIDType
	// 客户姓名
	CustomerName types.TThostFtdcIndividualNameType
	// 证件类型
	IdCardType types.TThostFtdcIdCardTypeType
	// 证件号码
	IdentifiedCardNo types.TThostFtdcIdentifiedCardNoType
	// 客户类型
	CustType types.TThostFtdcCustTypeType
	// 银行帐号
	BankAccount types.TThostFtdcBankAccountType
	// 银行密码
	BankPassWord types.TThostFtdcPasswordType
	// 投资者帐号
	AccountID types.TThostFtdcAccountIDType
	// 期货密码
	Password types.TThostFtdcPasswordType
	// 期货公司流水号
	FutureSerial types.TThostFtdcFutureSerialType
	// 安装编号
	InstallID types.TThostFtdcInstallIDType
	// 用户标识
	UserID types.TThostFtdcUserIDType
	// 验证客户证件号码标志
	VerifyCertNoFlag types.TThostFtdcYesNoIndicatorType
	// 币种代码
	CurrencyID types.TThostFtdcCurrencyIDType
	// 摘要
	Digest types.TThostFtdcDigestType
	// 银行帐号类型
	BankAccType types.TThostFtdcBankAccTypeType
	// 渠道标志
	DeviceID types.TThostFtdcDeviceIDType
	// 期货单位帐号类型
	BankSecuAccType types.TThostFtdcBankAccTypeType
	// 期货公司银行编码
	BrokerIDByBank types.TThostFtdcBankCodingForFutureType
	// 期货单位帐号
	BankSecuAcc types.TThostFtdcBankAccountType
	// 银行密码标志
	BankPwdFlag types.TThostFtdcPwdFlagType
	// 期货资金密码核对标志
	SecuPwdFlag types.TThostFtdcPwdFlagType
	// 交易柜员
	OperNo types.TThostFtdcOperNoType
	// 请求编号
	RequestID types.TThostFtdcRequestIDType
	// 交易ID
	TID types.TThostFtdcTIDType
}

// 查询账户信息响应
type CThostFtdcRspQueryAccountField struct {
	// 业务功能码
	TradeCode types.TThostFtdcTradeCodeType
	// 银行代码
	BankID types.TThostFtdcBankIDType
	// 银行分支机构代码
	BankBranchID types.TThostFtdcBankBrchIDType
	// 期商代码
	BrokerID types.TThostFtdcBrokerIDType
	// 期商分支机构代码
	BrokerBranchID types.TThostFtdcFutureBranchIDType
	// 交易日期
	TradeDate types.TThostFtdcTradeDateType
	// 交易时间
	TradeTime types.TThostFtdcTradeTimeType
	// 银行流水号
	BankSerial types.TThostFtdcBankSerialType
	// 交易系统日期
	TradingDay types.TThostFtdcTradeDateType
	// 银期平台消息流水号
	PlateSerial types.TThostFtdcSerialType
	// 最后分片标志
	LastFragment types.TThostFtdcLastFragmentType
	// 会话号
	SessionID types.TThostFtdcSessionIDType
	// 客户姓名
	CustomerName types.TThostFtdcIndividualNameType
	// 证件类型
	IdCardType types.TThostFtdcIdCardTypeType
	// 证件号码
	IdentifiedCardNo types.TThostFtdcIdentifiedCardNoType
	// 客户类型
	CustType types.TThostFtdcCustTypeType
	// 银行帐号
	BankAccount types.TThostFtdcBankAccountType
	// 银行密码
	BankPassWord types.TThostFtdcPasswordType
	// 投资者帐号
	AccountID types.TThostFtdcAccountIDType
	// 期货密码
	Password types.TThostFtdcPasswordType
	// 期货公司流水号
	FutureSerial types.TThostFtdcFutureSerialType
	// 安装编号
	InstallID types.TThostFtdcInstallIDType
	// 用户标识
	UserID types.TThostFtdcUserIDType
	// 验证客户证件号码标志
	VerifyCertNoFlag types.TThostFtdcYesNoIndicatorType
	// 币种代码
	CurrencyID types.TThostFtdcCurrencyIDType
	// 摘要
	Digest types.TThostFtdcDigestType
	// 银行帐号类型
	BankAccType types.TThostFtdcBankAccTypeType
	// 渠道标志
	DeviceID types.TThostFtdcDeviceIDType
	// 期货单位帐号类型
	BankSecuAccType types.TThostFtdcBankAccTypeType
	// 期货公司银行编码
	BrokerIDByBank types.TThostFtdcBankCodingForFutureType
	// 期货单位帐号
	BankSecuAcc types.TThostFtdcBankAccountType
	// 银行密码标志
	BankPwdFlag types.TThostFtdcPwdFlagType
	// 期货资金密码核对标志
	SecuPwdFlag types.TThostFtdcPwdFlagType
	// 交易柜员
	OperNo types.TThostFtdcOperNoType
	// 请求编号
	RequestID types.TThostFtdcRequestIDType
	// 交易ID
	TID types.TThostFtdcTIDType
	// 银行可用金额
	BankUseAmount types.TThostFtdcTradeAmountType
	// 银行可取金额
	BankFetchAmount types.TThostFtdcTradeAmountType
}

// 期商签到签退
type CThostFtdcFutureSignIOField struct {
	// 业务功能码
	TradeCode types.TThostFtdcTradeCodeType
	// 银行代码
	BankID types.TThostFtdcBankIDType
	// 银行分支机构代码
	BankBranchID types.TThostFtdcBankBrchIDType
	// 期商代码
	BrokerID types.TThostFtdcBrokerIDType
	// 期商分支机构代码
	BrokerBranchID types.TThostFtdcFutureBranchIDType
	// 交易日期
	TradeDate types.TThostFtdcTradeDateType
	// 交易时间
	TradeTime types.TThostFtdcTradeTimeType
	// 银行流水号
	BankSerial types.TThostFtdcBankSerialType
	// 交易系统日期
	TradingDay types.TThostFtdcTradeDateType
	// 银期平台消息流水号
	PlateSerial types.TThostFtdcSerialType
	// 最后分片标志
	LastFragment types.TThostFtdcLastFragmentType
	// 会话号
	SessionID types.TThostFtdcSessionIDType
	// 安装编号
	InstallID types.TThostFtdcInstallIDType
	// 用户标识
	UserID types.TThostFtdcUserIDType
	// 摘要
	Digest types.TThostFtdcDigestType
	// 币种代码
	CurrencyID types.TThostFtdcCurrencyIDType
	// 渠道标志
	DeviceID types.TThostFtdcDeviceIDType
	// 期货公司银行编码
	BrokerIDByBank types.TThostFtdcBankCodingForFutureType
	// 交易柜员
	OperNo types.TThostFtdcOperNoType
	// 请求编号
	RequestID types.TThostFtdcRequestIDType
	// 交易ID
	TID types.TThostFtdcTIDType
}

// 期商签到响应
type CThostFtdcRspFutureSignInField struct {
	// 业务功能码
	TradeCode types.TThostFtdcTradeCodeType
	// 银行代码
	BankID types.TThostFtdcBankIDType
	// 银行分支机构代码
	BankBranchID types.TThostFtdcBankBrchIDType
	// 期商代码
	BrokerID types.TThostFtdcBrokerIDType
	// 期商分支机构代码
	BrokerBranchID types.TThostFtdcFutureBranchIDType
	// 交易日期
	TradeDate types.TThostFtdcTradeDateType
	// 交易时间
	TradeTime types.TThostFtdcTradeTimeType
	// 银行流水号
	BankSerial types.TThostFtdcBankSerialType
	// 交易系统日期
	TradingDay types.TThostFtdcTradeDateType
	// 银期平台消息流水号
	PlateSerial types.TThostFtdcSerialType
	// 最后分片标志
	LastFragment types.TThostFtdcLastFragmentType
	// 会话号
	SessionID types.TThostFtdcSessionIDType
	// 安装编号
	InstallID types.TThostFtdcInstallIDType
	// 用户标识
	UserID types.TThostFtdcUserIDType
	// 摘要
	Digest types.TThostFtdcDigestType
	// 币种代码
	CurrencyID types.TThostFtdcCurrencyIDType
	// 渠道标志
	DeviceID types.TThostFtdcDeviceIDType
	// 期货公司银行编码
	BrokerIDByBank types.TThostFtdcBankCodingForFutureType
	// 交易柜员
	OperNo types.TThostFtdcOperNoType
	// 请求编号
	RequestID types.TThostFtdcRequestIDType
	// 交易ID
	TID types.TThostFtdcTIDType
	// 错误代码
	ErrorID types.TThostFtdcErrorIDType
	// 错误信息
	ErrorMsg types.TThostFtdcErrorMsgType
	// PIN密钥
	PinKey types.TThostFtdcPasswordKeyType
	// MAC密钥
	MacKey types.TThostFtdcPasswordKeyType
}

// 期商签退请求
type CThostFtdcReqFutureSignOutField struct {
	// 业务功能码
	TradeCode types.TThostFtdcTradeCodeType
	// 银行代码
	BankID types.TThostFtdcBankIDType
	// 银行分支机构代码
	BankBranchID types.TThostFtdcBankBrchIDType
	// 期商代码
	BrokerID types.TThostFtdcBrokerIDType
	// 期商分支机构代码
	BrokerBranchID types.TThostFtdcFutureBranchIDType
	// 交易日期
	TradeDate types.TThostFtdcTradeDateType
	// 交易时间
	TradeTime types.TThostFtdcTradeTimeType
	// 银行流水号
	BankSerial types.TThostFtdcBankSerialType
	// 交易系统日期
	TradingDay types.TThostFtdcTradeDateType
	// 银期平台消息流水号
	PlateSerial types.TThostFtdcSerialType
	// 最后分片标志
	LastFragment types.TThostFtdcLastFragmentType
	// 会话号
	SessionID types.TThostFtdcSessionIDType
	// 安装编号
	InstallID types.TThostFtdcInstallIDType
	// 用户标识
	UserID types.TThostFtdcUserIDType
	// 摘要
	Digest types.TThostFtdcDigestType
	// 币种代码
	CurrencyID types.TThostFtdcCurrencyIDType
	// 渠道标志
	DeviceID types.TThostFtdcDeviceIDType
	// 期货公司银行编码
	BrokerIDByBank types.TThostFtdcBankCodingForFutureType
	// 交易柜员
	OperNo types.TThostFtdcOperNoType
	// 请求编号
	RequestID types.TThostFtdcRequestIDType
	// 交易ID
	TID types.TThostFtdcTIDType
}

// 期商签退响应
type CThostFtdcRspFutureSignOutField struct {
	// 业务功能码
	TradeCode types.TThostFtdcTradeCodeType
	// 银行代码
	BankID types.TThostFtdcBankIDType
	// 银行分支机构代码
	BankBranchID types.TThostFtdcBankBrchIDType
	// 期商代码
	BrokerID types.TThostFtdcBrokerIDType
	// 期商分支机构代码
	BrokerBranchID types.TThostFtdcFutureBranchIDType
	// 交易日期
	TradeDate types.TThostFtdcTradeDateType
	// 交易时间
	TradeTime types.TThostFtdcTradeTimeType
	// 银行流水号
	BankSerial types.TThostFtdcBankSerialType
	// 交易系统日期
	TradingDay types.TThostFtdcTradeDateType
	// 银期平台消息流水号
	PlateSerial types.TThostFtdcSerialType
	// 最后分片标志
	LastFragment types.TThostFtdcLastFragmentType
	// 会话号
	SessionID types.TThostFtdcSessionIDType
	// 安装编号
	InstallID types.TThostFtdcInstallIDType
	// 用户标识
	UserID types.TThostFtdcUserIDType
	// 摘要
	Digest types.TThostFtdcDigestType
	// 币种代码
	CurrencyID types.TThostFtdcCurrencyIDType
	// 渠道标志
	DeviceID types.TThostFtdcDeviceIDType
	// 期货公司银行编码
	BrokerIDByBank types.TThostFtdcBankCodingForFutureType
	// 交易柜员
	OperNo types.TThostFtdcOperNoType
	// 请求编号
	RequestID types.TThostFtdcRequestIDType
	// 交易ID
	TID types.TThostFtdcTIDType
	// 错误代码
	ErrorID types.TThostFtdcErrorIDType
	// 错误信息
	ErrorMsg types.TThostFtdcErrorMsgType
}

// 查询指定流水号的交易结果请求
type CThostFtdcReqQueryTradeResultBySerialField struct {
	// 业务功能码
	TradeCode types.TThostFtdcTradeCodeType
	// 银行代码
	BankID types.TThostFtdcBankIDType
	// 银行分支机构代码
	BankBranchID types.TThostFtdcBankBrchIDType
	// 期商代码
	BrokerID types.TThostFtdcBrokerIDType
	// 期商分支机构代码
	BrokerBranchID types.TThostFtdcFutureBranchIDType
	// 交易日期
	TradeDate types.TThostFtdcTradeDateType
	// 交易时间
	TradeTime types.TThostFtdcTradeTimeType
	// 银行流水号
	BankSerial types.TThostFtdcBankSerialType
	// 交易系统日期
	TradingDay types.TThostFtdcTradeDateType
	// 银期平台消息流水号
	PlateSerial types.TThostFtdcSerialType
	// 最后分片标志
	LastFragment types.TThostFtdcLastFragmentType
	// 会话号
	SessionID types.TThostFtdcSessionIDType
	// 流水号
	Reference types.TThostFtdcSerialType
	// 本流水号发布者的机构类型
	RefrenceIssureType types.TThostFtdcInstitutionTypeType
	// 本流水号发布者机构编码
	RefrenceIssure types.TThostFtdcOrganCodeType
	// 客户姓名
	CustomerName types.TThostFtdcIndividualNameType
	// 证件类型
	IdCardType types.TThostFtdcIdCardTypeType
	// 证件号码
	IdentifiedCardNo types.TThostFtdcIdentifiedCardNoType
	// 客户类型
	CustType types.TThostFtdcCustTypeType
	// 银行帐号
	BankAccount types.TThostFtdcBankAccountType
	// 银行密码
	BankPassWord types.TThostFtdcPasswordType
	// 投资者帐号
	AccountID types.TThostFtdcAccountIDType
	// 期货密码
	Password types.TThostFtdcPasswordType
	// 币种代码
	CurrencyID types.TThostFtdcCurrencyIDType
	// 转帐金额
	TradeAmount types.TThostFtdcTradeAmountType
	// 摘要
	Digest types.TThostFtdcDigestType
}

// 查询指定流水号的交易结果响应
type CThostFtdcRspQueryTradeResultBySerialField struct {
	// 业务功能码
	TradeCode types.TThostFtdcTradeCodeType
	// 银行代码
	BankID types.TThostFtdcBankIDType
	// 银行分支机构代码
	BankBranchID types.TThostFtdcBankBrchIDType
	// 期商代码
	BrokerID types.TThostFtdcBrokerIDType
	// 期商分支机构代码
	BrokerBranchID types.TThostFtdcFutureBranchIDType
	// 交易日期
	TradeDate types.TThostFtdcTradeDateType
	// 交易时间
	TradeTime types.TThostFtdcTradeTimeType
	// 银行流水号
	BankSerial types.TThostFtdcBankSerialType
	// 交易系统日期
	TradingDay types.TThostFtdcTradeDateType
	// 银期平台消息流水号
	PlateSerial types.TThostFtdcSerialType
	// 最后分片标志
	LastFragment types.TThostFtdcLastFragmentType
	// 会话号
	SessionID types.TThostFtdcSessionIDType
	// 错误代码
	ErrorID types.TThostFtdcErrorIDType
	// 错误信息
	ErrorMsg types.TThostFtdcErrorMsgType
	// 流水号
	Reference types.TThostFtdcSerialType
	// 本流水号发布者的机构类型
	RefrenceIssureType types.TThostFtdcInstitutionTypeType
	// 本流水号发布者机构编码
	RefrenceIssure types.TThostFtdcOrganCodeType
	// 原始返回代码
	OriginReturnCode types.TThostFtdcReturnCodeType
	// 原始返回码描述
	OriginDescrInfoForReturnCode types.TThostFtdcDescrInfoForReturnCodeType
	// 银行帐号
	BankAccount types.TThostFtdcBankAccountType
	// 银行密码
	BankPassWord types.TThostFtdcPasswordType
	// 投资者帐号
	AccountID types.TThostFtdcAccountIDType
	// 期货密码
	Password types.TThostFtdcPasswordType
	// 币种代码
	CurrencyID types.TThostFtdcCurrencyIDType
	// 转帐金额
	TradeAmount types.TThostFtdcTradeAmountType
	// 摘要
	Digest types.TThostFtdcDigestType
}

// 日终文件就绪请求
type CThostFtdcReqDayEndFileReadyField struct {
	// 业务功能码
	TradeCode types.TThostFtdcTradeCodeType
	// 银行代码
	BankID types.TThostFtdcBankIDType
	// 银行分支机构代码
	BankBranchID types.TThostFtdcBankBrchIDType
	// 期商代码
	BrokerID types.TThostFtdcBrokerIDType
	// 期商分支机构代码
	BrokerBranchID types.TThostFtdcFutureBranchIDType
	// 交易日期
	TradeDate types.TThostFtdcTradeDateType
	// 交易时间
	TradeTime types.TThostFtdcTradeTimeType
	// 银行流水号
	BankSerial types.TThostFtdcBankSerialType
	// 交易系统日期
	TradingDay types.TThostFtdcTradeDateType
	// 银期平台消息流水号
	PlateSerial types.TThostFtdcSerialType
	// 最后分片标志
	LastFragment types.TThostFtdcLastFragmentType
	// 会话号
	SessionID types.TThostFtdcSessionIDType
	// 文件业务功能
	FileBusinessCode types.TThostFtdcFileBusinessCodeType
	// 摘要
	Digest types.TThostFtdcDigestType
}

// 返回结果
type CThostFtdcReturnResultField struct {
	// 返回代码
	ReturnCode types.TThostFtdcReturnCodeType
	// 返回码描述
	DescrInfoForReturnCode types.TThostFtdcDescrInfoForReturnCodeType
}

// 验证期货资金密码
type CThostFtdcVerifyFuturePasswordField struct {
	// 业务功能码
	TradeCode types.TThostFtdcTradeCodeType
	// 银行代码
	BankID types.TThostFtdcBankIDType
	// 银行分支机构代码
	BankBranchID types.TThostFtdcBankBrchIDType
	// 期商代码
	BrokerID types.TThostFtdcBrokerIDType
	// 期商分支机构代码
	BrokerBranchID types.TThostFtdcFutureBranchIDType
	// 交易日期
	TradeDate types.TThostFtdcTradeDateType
	// 交易时间
	TradeTime types.TThostFtdcTradeTimeType
	// 银行流水号
	BankSerial types.TThostFtdcBankSerialType
	// 交易系统日期
	TradingDay types.TThostFtdcTradeDateType
	// 银期平台消息流水号
	PlateSerial types.TThostFtdcSerialType
	// 最后分片标志
	LastFragment types.TThostFtdcLastFragmentType
	// 会话号
	SessionID types.TThostFtdcSessionIDType
	// 投资者帐号
	AccountID types.TThostFtdcAccountIDType
	// 期货密码
	Password types.TThostFtdcPasswordType
	// 银行帐号
	BankAccount types.TThostFtdcBankAccountType
	// 银行密码
	BankPassWord types.TThostFtdcPasswordType
	// 安装编号
	InstallID types.TThostFtdcInstallIDType
	// 交易ID
	TID types.TThostFtdcTIDType
	// 币种代码
	CurrencyID types.TThostFtdcCurrencyIDType
}

// 验证客户信息
type CThostFtdcVerifyCustInfoField struct {
	// 客户姓名
	CustomerName types.TThostFtdcIndividualNameType
	// 证件类型
	IdCardType types.TThostFtdcIdCardTypeType
	// 证件号码
	IdentifiedCardNo types.TThostFtdcIdentifiedCardNoType
	// 客户类型
	CustType types.TThostFtdcCustTypeType
}

// 验证期货资金密码和客户信息
type CThostFtdcVerifyFuturePasswordAndCustInfoField struct {
	// 客户姓名
	CustomerName types.TThostFtdcIndividualNameType
	// 证件类型
	IdCardType types.TThostFtdcIdCardTypeType
	// 证件号码
	IdentifiedCardNo types.TThostFtdcIdentifiedCardNoType
	// 客户类型
	CustType types.TThostFtdcCustTypeType
	// 投资者帐号
	AccountID types.TThostFtdcAccountIDType
	// 期货密码
	Password types.TThostFtdcPasswordType
	// 币种代码
	CurrencyID types.TThostFtdcCurrencyIDType
}

// 验证期货资金密码和客户信息
type CThostFtdcDepositResultInformField struct {
	// 出入金流水号，该流水号为银期报盘返回的流水号
	DepositSeqNo types.TThostFtdcDepositSeqNoType
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	// 入金金额
	Deposit types.TThostFtdcMoneyType
	// 请求编号
	RequestID types.TThostFtdcRequestIDType
	// 返回代码
	ReturnCode types.TThostFtdcReturnCodeType
	// 返回码描述
	DescrInfoForReturnCode types.TThostFtdcDescrInfoForReturnCodeType
}

// 交易核心向银期报盘发出密钥同步请求
type CThostFtdcReqSyncKeyField struct {
	// 业务功能码
	TradeCode types.TThostFtdcTradeCodeType
	// 银行代码
	BankID types.TThostFtdcBankIDType
	// 银行分支机构代码
	BankBranchID types.TThostFtdcBankBrchIDType
	// 期商代码
	BrokerID types.TThostFtdcBrokerIDType
	// 期商分支机构代码
	BrokerBranchID types.TThostFtdcFutureBranchIDType
	// 交易日期
	TradeDate types.TThostFtdcTradeDateType
	// 交易时间
	TradeTime types.TThostFtdcTradeTimeType
	// 银行流水号
	BankSerial types.TThostFtdcBankSerialType
	// 交易系统日期
	TradingDay types.TThostFtdcTradeDateType
	// 银期平台消息流水号
	PlateSerial types.TThostFtdcSerialType
	// 最后分片标志
	LastFragment types.TThostFtdcLastFragmentType
	// 会话号
	SessionID types.TThostFtdcSessionIDType
	// 安装编号
	InstallID types.TThostFtdcInstallIDType
	// 用户标识
	UserID types.TThostFtdcUserIDType
	// 交易核心给银期报盘的消息
	Message types.TThostFtdcAddInfoType
	// 渠道标志
	DeviceID types.TThostFtdcDeviceIDType
	// 期货公司银行编码
	BrokerIDByBank types.TThostFtdcBankCodingForFutureType
	// 交易柜员
	OperNo types.TThostFtdcOperNoType
	// 请求编号
	RequestID types.TThostFtdcRequestIDType
	// 交易ID
	TID types.TThostFtdcTIDType
}

// 交易核心向银期报盘发出密钥同步响应
type CThostFtdcRspSyncKeyField struct {
	// 业务功能码
	TradeCode types.TThostFtdcTradeCodeType
	// 银行代码
	BankID types.TThostFtdcBankIDType
	// 银行分支机构代码
	BankBranchID types.TThostFtdcBankBrchIDType
	// 期商代码
	BrokerID types.TThostFtdcBrokerIDType
	// 期商分支机构代码
	BrokerBranchID types.TThostFtdcFutureBranchIDType
	// 交易日期
	TradeDate types.TThostFtdcTradeDateType
	// 交易时间
	TradeTime types.TThostFtdcTradeTimeType
	// 银行流水号
	BankSerial types.TThostFtdcBankSerialType
	// 交易系统日期
	TradingDay types.TThostFtdcTradeDateType
	// 银期平台消息流水号
	PlateSerial types.TThostFtdcSerialType
	// 最后分片标志
	LastFragment types.TThostFtdcLastFragmentType
	// 会话号
	SessionID types.TThostFtdcSessionIDType
	// 安装编号
	InstallID types.TThostFtdcInstallIDType
	// 用户标识
	UserID types.TThostFtdcUserIDType
	// 交易核心给银期报盘的消息
	Message types.TThostFtdcAddInfoType
	// 渠道标志
	DeviceID types.TThostFtdcDeviceIDType
	// 期货公司银行编码
	BrokerIDByBank types.TThostFtdcBankCodingForFutureType
	// 交易柜员
	OperNo types.TThostFtdcOperNoType
	// 请求编号
	RequestID types.TThostFtdcRequestIDType
	// 交易ID
	TID types.TThostFtdcTIDType
	// 错误代码
	ErrorID types.TThostFtdcErrorIDType
	// 错误信息
	ErrorMsg types.TThostFtdcErrorMsgType
}

// 查询账户信息通知
type CThostFtdcNotifyQueryAccountField struct {
	// 业务功能码
	TradeCode types.TThostFtdcTradeCodeType
	// 银行代码
	BankID types.TThostFtdcBankIDType
	// 银行分支机构代码
	BankBranchID types.TThostFtdcBankBrchIDType
	// 期商代码
	BrokerID types.TThostFtdcBrokerIDType
	// 期商分支机构代码
	BrokerBranchID types.TThostFtdcFutureBranchIDType
	// 交易日期
	TradeDate types.TThostFtdcTradeDateType
	// 交易时间
	TradeTime types.TThostFtdcTradeTimeType
	// 银行流水号
	BankSerial types.TThostFtdcBankSerialType
	// 交易系统日期
	TradingDay types.TThostFtdcTradeDateType
	// 银期平台消息流水号
	PlateSerial types.TThostFtdcSerialType
	// 最后分片标志
	LastFragment types.TThostFtdcLastFragmentType
	// 会话号
	SessionID types.TThostFtdcSessionIDType
	// 客户姓名
	CustomerName types.TThostFtdcIndividualNameType
	// 证件类型
	IdCardType types.TThostFtdcIdCardTypeType
	// 证件号码
	IdentifiedCardNo types.TThostFtdcIdentifiedCardNoType
	// 客户类型
	CustType types.TThostFtdcCustTypeType
	// 银行帐号
	BankAccount types.TThostFtdcBankAccountType
	// 银行密码
	BankPassWord types.TThostFtdcPasswordType
	// 投资者帐号
	AccountID types.TThostFtdcAccountIDType
	// 期货密码
	Password types.TThostFtdcPasswordType
	// 期货公司流水号
	FutureSerial types.TThostFtdcFutureSerialType
	// 安装编号
	InstallID types.TThostFtdcInstallIDType
	// 用户标识
	UserID types.TThostFtdcUserIDType
	// 验证客户证件号码标志
	VerifyCertNoFlag types.TThostFtdcYesNoIndicatorType
	// 币种代码
	CurrencyID types.TThostFtdcCurrencyIDType
	// 摘要
	Digest types.TThostFtdcDigestType
	// 银行帐号类型
	BankAccType types.TThostFtdcBankAccTypeType
	// 渠道标志
	DeviceID types.TThostFtdcDeviceIDType
	// 期货单位帐号类型
	BankSecuAccType types.TThostFtdcBankAccTypeType
	// 期货公司银行编码
	BrokerIDByBank types.TThostFtdcBankCodingForFutureType
	// 期货单位帐号
	BankSecuAcc types.TThostFtdcBankAccountType
	// 银行密码标志
	BankPwdFlag types.TThostFtdcPwdFlagType
	// 期货资金密码核对标志
	SecuPwdFlag types.TThostFtdcPwdFlagType
	// 交易柜员
	OperNo types.TThostFtdcOperNoType
	// 请求编号
	RequestID types.TThostFtdcRequestIDType
	// 交易ID
	TID types.TThostFtdcTIDType
	// 银行可用金额
	BankUseAmount types.TThostFtdcTradeAmountType
	// 银行可取金额
	BankFetchAmount types.TThostFtdcTradeAmountType
	// 错误代码
	ErrorID types.TThostFtdcErrorIDType
	// 错误信息
	ErrorMsg types.TThostFtdcErrorMsgType
}

// 银期转账交易流水表
type CThostFtdcTransferSerialField struct {
	// 平台流水号
	PlateSerial types.TThostFtdcPlateSerialType
	// 交易发起方日期
	TradeDate types.TThostFtdcTradeDateType
	// 交易日期
	TradingDay types.TThostFtdcDateType
	// 交易时间
	TradeTime types.TThostFtdcTradeTimeType
	// 交易代码
	TradeCode types.TThostFtdcTradeCodeType
	// 会话编号
	SessionID types.TThostFtdcSessionIDType
	// 银行编码
	BankID types.TThostFtdcBankIDType
	// 银行分支机构编码
	BankBranchID types.TThostFtdcBankBrchIDType
	// 银行帐号类型
	BankAccType types.TThostFtdcBankAccTypeType
	// 银行帐号
	BankAccount types.TThostFtdcBankAccountType
	// 银行流水号
	BankSerial types.TThostFtdcBankSerialType
	// 期货公司编码
	BrokerID types.TThostFtdcBrokerIDType
	// 期商分支机构代码
	BrokerBranchID types.TThostFtdcFutureBranchIDType
	// 期货公司帐号类型
	FutureAccType types.TThostFtdcFutureAccTypeType
	// 投资者帐号
	AccountID types.TThostFtdcAccountIDType
	// 投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	// 期货公司流水号
	FutureSerial types.TThostFtdcFutureSerialType
	// 证件类型
	IdCardType types.TThostFtdcIdCardTypeType
	// 证件号码
	IdentifiedCardNo types.TThostFtdcIdentifiedCardNoType
	// 币种代码
	CurrencyID types.TThostFtdcCurrencyIDType
	// 交易金额
	TradeAmount types.TThostFtdcTradeAmountType
	// 应收客户费用
	CustFee types.TThostFtdcCustFeeType
	// 应收期货公司费用
	BrokerFee types.TThostFtdcFutureFeeType
	// 有效标志
	AvailabilityFlag types.TThostFtdcAvailabilityFlagType
	// 操作员
	OperatorCode types.TThostFtdcOperatorCodeType
	// 新银行帐号
	BankNewAccount types.TThostFtdcBankAccountType
	// 错误代码
	ErrorID types.TThostFtdcErrorIDType
	// 错误信息
	ErrorMsg types.TThostFtdcErrorMsgType
}

// 请求查询转帐流水
type CThostFtdcQryTransferSerialField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 投资者帐号
	AccountID types.TThostFtdcAccountIDType
	// 银行编码
	BankID types.TThostFtdcBankIDType
	// 币种代码
	CurrencyID types.TThostFtdcCurrencyIDType
}

// 期商签到通知
type CThostFtdcNotifyFutureSignInField struct {
	// 业务功能码
	TradeCode types.TThostFtdcTradeCodeType
	// 银行代码
	BankID types.TThostFtdcBankIDType
	// 银行分支机构代码
	BankBranchID types.TThostFtdcBankBrchIDType
	// 期商代码
	BrokerID types.TThostFtdcBrokerIDType
	// 期商分支机构代码
	BrokerBranchID types.TThostFtdcFutureBranchIDType
	// 交易日期
	TradeDate types.TThostFtdcTradeDateType
	// 交易时间
	TradeTime types.TThostFtdcTradeTimeType
	// 银行流水号
	BankSerial types.TThostFtdcBankSerialType
	// 交易系统日期
	TradingDay types.TThostFtdcTradeDateType
	// 银期平台消息流水号
	PlateSerial types.TThostFtdcSerialType
	// 最后分片标志
	LastFragment types.TThostFtdcLastFragmentType
	// 会话号
	SessionID types.TThostFtdcSessionIDType
	// 安装编号
	InstallID types.TThostFtdcInstallIDType
	// 用户标识
	UserID types.TThostFtdcUserIDType
	// 摘要
	Digest types.TThostFtdcDigestType
	// 币种代码
	CurrencyID types.TThostFtdcCurrencyIDType
	// 渠道标志
	DeviceID types.TThostFtdcDeviceIDType
	// 期货公司银行编码
	BrokerIDByBank types.TThostFtdcBankCodingForFutureType
	// 交易柜员
	OperNo types.TThostFtdcOperNoType
	// 请求编号
	RequestID types.TThostFtdcRequestIDType
	// 交易ID
	TID types.TThostFtdcTIDType
	// 错误代码
	ErrorID types.TThostFtdcErrorIDType
	// 错误信息
	ErrorMsg types.TThostFtdcErrorMsgType
	// PIN密钥
	PinKey types.TThostFtdcPasswordKeyType
	// MAC密钥
	MacKey types.TThostFtdcPasswordKeyType
}

// 期商签退通知
type CThostFtdcNotifyFutureSignOutField struct {
	// 业务功能码
	TradeCode types.TThostFtdcTradeCodeType
	// 银行代码
	BankID types.TThostFtdcBankIDType
	// 银行分支机构代码
	BankBranchID types.TThostFtdcBankBrchIDType
	// 期商代码
	BrokerID types.TThostFtdcBrokerIDType
	// 期商分支机构代码
	BrokerBranchID types.TThostFtdcFutureBranchIDType
	// 交易日期
	TradeDate types.TThostFtdcTradeDateType
	// 交易时间
	TradeTime types.TThostFtdcTradeTimeType
	// 银行流水号
	BankSerial types.TThostFtdcBankSerialType
	// 交易系统日期
	TradingDay types.TThostFtdcTradeDateType
	// 银期平台消息流水号
	PlateSerial types.TThostFtdcSerialType
	// 最后分片标志
	LastFragment types.TThostFtdcLastFragmentType
	// 会话号
	SessionID types.TThostFtdcSessionIDType
	// 安装编号
	InstallID types.TThostFtdcInstallIDType
	// 用户标识
	UserID types.TThostFtdcUserIDType
	// 摘要
	Digest types.TThostFtdcDigestType
	// 币种代码
	CurrencyID types.TThostFtdcCurrencyIDType
	// 渠道标志
	DeviceID types.TThostFtdcDeviceIDType
	// 期货公司银行编码
	BrokerIDByBank types.TThostFtdcBankCodingForFutureType
	// 交易柜员
	OperNo types.TThostFtdcOperNoType
	// 请求编号
	RequestID types.TThostFtdcRequestIDType
	// 交易ID
	TID types.TThostFtdcTIDType
	// 错误代码
	ErrorID types.TThostFtdcErrorIDType
	// 错误信息
	ErrorMsg types.TThostFtdcErrorMsgType
}

// 交易核心向银期报盘发出密钥同步处理结果的通知
type CThostFtdcNotifySyncKeyField struct {
	// 业务功能码
	TradeCode types.TThostFtdcTradeCodeType
	// 银行代码
	BankID types.TThostFtdcBankIDType
	// 银行分支机构代码
	BankBranchID types.TThostFtdcBankBrchIDType
	// 期商代码
	BrokerID types.TThostFtdcBrokerIDType
	// 期商分支机构代码
	BrokerBranchID types.TThostFtdcFutureBranchIDType
	// 交易日期
	TradeDate types.TThostFtdcTradeDateType
	// 交易时间
	TradeTime types.TThostFtdcTradeTimeType
	// 银行流水号
	BankSerial types.TThostFtdcBankSerialType
	// 交易系统日期
	TradingDay types.TThostFtdcTradeDateType
	// 银期平台消息流水号
	PlateSerial types.TThostFtdcSerialType
	// 最后分片标志
	LastFragment types.TThostFtdcLastFragmentType
	// 会话号
	SessionID types.TThostFtdcSessionIDType
	// 安装编号
	InstallID types.TThostFtdcInstallIDType
	// 用户标识
	UserID types.TThostFtdcUserIDType
	// 交易核心给银期报盘的消息
	Message types.TThostFtdcAddInfoType
	// 渠道标志
	DeviceID types.TThostFtdcDeviceIDType
	// 期货公司银行编码
	BrokerIDByBank types.TThostFtdcBankCodingForFutureType
	// 交易柜员
	OperNo types.TThostFtdcOperNoType
	// 请求编号
	RequestID types.TThostFtdcRequestIDType
	// 交易ID
	TID types.TThostFtdcTIDType
	// 错误代码
	ErrorID types.TThostFtdcErrorIDType
	// 错误信息
	ErrorMsg types.TThostFtdcErrorMsgType
}

// 请求查询银期签约关系
type CThostFtdcQryAccountregisterField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 投资者帐号
	AccountID types.TThostFtdcAccountIDType
	// 银行编码
	BankID types.TThostFtdcBankIDType
	// 银行分支机构编码
	BankBranchID types.TThostFtdcBankBrchIDType
	// 币种代码
	CurrencyID types.TThostFtdcCurrencyIDType
}

// 客户开销户信息表
type CThostFtdcAccountregisterField struct {
	// 交易日期
	TradeDay types.TThostFtdcTradeDateType
	// 银行编码
	BankID types.TThostFtdcBankIDType
	// 银行分支机构编码
	BankBranchID types.TThostFtdcBankBrchIDType
	// 银行帐号
	BankAccount types.TThostFtdcBankAccountType
	// 期货公司编码
	BrokerID types.TThostFtdcBrokerIDType
	// 期货公司分支机构编码
	BrokerBranchID types.TThostFtdcFutureBranchIDType
	// 投资者帐号
	AccountID types.TThostFtdcAccountIDType
	// 证件类型
	IdCardType types.TThostFtdcIdCardTypeType
	// 证件号码
	IdentifiedCardNo types.TThostFtdcIdentifiedCardNoType
	// 客户姓名
	CustomerName types.TThostFtdcIndividualNameType
	// 币种代码
	CurrencyID types.TThostFtdcCurrencyIDType
	// 开销户类别
	OpenOrDestroy types.TThostFtdcOpenOrDestroyType
	// 签约日期
	RegDate types.TThostFtdcTradeDateType
	// 解约日期
	OutDate types.TThostFtdcTradeDateType
	// 交易ID
	TID types.TThostFtdcTIDType
	// 客户类型
	CustType types.TThostFtdcCustTypeType
	// 银行帐号类型
	BankAccType types.TThostFtdcBankAccTypeType
}

// 银期开户信息
type CThostFtdcOpenAccountField struct {
	// 业务功能码
	TradeCode types.TThostFtdcTradeCodeType
	// 银行代码
	BankID types.TThostFtdcBankIDType
	// 银行分支机构代码
	BankBranchID types.TThostFtdcBankBrchIDType
	// 期商代码
	BrokerID types.TThostFtdcBrokerIDType
	// 期商分支机构代码
	BrokerBranchID types.TThostFtdcFutureBranchIDType
	// 交易日期
	TradeDate types.TThostFtdcTradeDateType
	// 交易时间
	TradeTime types.TThostFtdcTradeTimeType
	// 银行流水号
	BankSerial types.TThostFtdcBankSerialType
	// 交易系统日期
	TradingDay types.TThostFtdcTradeDateType
	// 银期平台消息流水号
	PlateSerial types.TThostFtdcSerialType
	// 最后分片标志
	LastFragment types.TThostFtdcLastFragmentType
	// 会话号
	SessionID types.TThostFtdcSessionIDType
	// 客户姓名
	CustomerName types.TThostFtdcIndividualNameType
	// 证件类型
	IdCardType types.TThostFtdcIdCardTypeType
	// 证件号码
	IdentifiedCardNo types.TThostFtdcIdentifiedCardNoType
	// 性别
	Gender types.TThostFtdcGenderType
	// 国家代码
	CountryCode types.TThostFtdcCountryCodeType
	// 客户类型
	CustType types.TThostFtdcCustTypeType
	// 地址
	Address types.TThostFtdcAddressType
	// 邮编
	ZipCode types.TThostFtdcZipCodeType
	// 电话号码
	Telephone types.TThostFtdcTelephoneType
	// 手机
	MobilePhone types.TThostFtdcMobilePhoneType
	// 传真
	Fax types.TThostFtdcFaxType
	// 电子邮件
	EMail types.TThostFtdcEMailType
	// 资金账户状态
	MoneyAccountStatus types.TThostFtdcMoneyAccountStatusType
	// 银行帐号
	BankAccount types.TThostFtdcBankAccountType
	// 银行密码
	BankPassWord types.TThostFtdcPasswordType
	// 投资者帐号
	AccountID types.TThostFtdcAccountIDType
	// 期货密码
	Password types.TThostFtdcPasswordType
	// 安装编号
	InstallID types.TThostFtdcInstallIDType
	// 验证客户证件号码标志
	VerifyCertNoFlag types.TThostFtdcYesNoIndicatorType
	// 币种代码
	CurrencyID types.TThostFtdcCurrencyIDType
	// 汇钞标志
	CashExchangeCode types.TThostFtdcCashExchangeCodeType
	// 摘要
	Digest types.TThostFtdcDigestType
	// 银行帐号类型
	BankAccType types.TThostFtdcBankAccTypeType
	// 渠道标志
	DeviceID types.TThostFtdcDeviceIDType
	// 期货单位帐号类型
	BankSecuAccType types.TThostFtdcBankAccTypeType
	// 期货公司银行编码
	BrokerIDByBank types.TThostFtdcBankCodingForFutureType
	// 期货单位帐号
	BankSecuAcc types.TThostFtdcBankAccountType
	// 银行密码标志
	BankPwdFlag types.TThostFtdcPwdFlagType
	// 期货资金密码核对标志
	SecuPwdFlag types.TThostFtdcPwdFlagType
	// 交易柜员
	OperNo types.TThostFtdcOperNoType
	// 交易ID
	TID types.TThostFtdcTIDType
	// 用户标识
	UserID types.TThostFtdcUserIDType
	// 错误代码
	ErrorID types.TThostFtdcErrorIDType
	// 错误信息
	ErrorMsg types.TThostFtdcErrorMsgType
}

// 银期销户信息
type CThostFtdcCancelAccountField struct {
	// 业务功能码
	TradeCode types.TThostFtdcTradeCodeType
	// 银行代码
	BankID types.TThostFtdcBankIDType
	// 银行分支机构代码
	BankBranchID types.TThostFtdcBankBrchIDType
	// 期商代码
	BrokerID types.TThostFtdcBrokerIDType
	// 期商分支机构代码
	BrokerBranchID types.TThostFtdcFutureBranchIDType
	// 交易日期
	TradeDate types.TThostFtdcTradeDateType
	// 交易时间
	TradeTime types.TThostFtdcTradeTimeType
	// 银行流水号
	BankSerial types.TThostFtdcBankSerialType
	// 交易系统日期
	TradingDay types.TThostFtdcTradeDateType
	// 银期平台消息流水号
	PlateSerial types.TThostFtdcSerialType
	// 最后分片标志
	LastFragment types.TThostFtdcLastFragmentType
	// 会话号
	SessionID types.TThostFtdcSessionIDType
	// 客户姓名
	CustomerName types.TThostFtdcIndividualNameType
	// 证件类型
	IdCardType types.TThostFtdcIdCardTypeType
	// 证件号码
	IdentifiedCardNo types.TThostFtdcIdentifiedCardNoType
	// 性别
	Gender types.TThostFtdcGenderType
	// 国家代码
	CountryCode types.TThostFtdcCountryCodeType
	// 客户类型
	CustType types.TThostFtdcCustTypeType
	// 地址
	Address types.TThostFtdcAddressType
	// 邮编
	ZipCode types.TThostFtdcZipCodeType
	// 电话号码
	Telephone types.TThostFtdcTelephoneType
	// 手机
	MobilePhone types.TThostFtdcMobilePhoneType
	// 传真
	Fax types.TThostFtdcFaxType
	// 电子邮件
	EMail types.TThostFtdcEMailType
	// 资金账户状态
	MoneyAccountStatus types.TThostFtdcMoneyAccountStatusType
	// 银行帐号
	BankAccount types.TThostFtdcBankAccountType
	// 银行密码
	BankPassWord types.TThostFtdcPasswordType
	// 投资者帐号
	AccountID types.TThostFtdcAccountIDType
	// 期货密码
	Password types.TThostFtdcPasswordType
	// 安装编号
	InstallID types.TThostFtdcInstallIDType
	// 验证客户证件号码标志
	VerifyCertNoFlag types.TThostFtdcYesNoIndicatorType
	// 币种代码
	CurrencyID types.TThostFtdcCurrencyIDType
	// 汇钞标志
	CashExchangeCode types.TThostFtdcCashExchangeCodeType
	// 摘要
	Digest types.TThostFtdcDigestType
	// 银行帐号类型
	BankAccType types.TThostFtdcBankAccTypeType
	// 渠道标志
	DeviceID types.TThostFtdcDeviceIDType
	// 期货单位帐号类型
	BankSecuAccType types.TThostFtdcBankAccTypeType
	// 期货公司银行编码
	BrokerIDByBank types.TThostFtdcBankCodingForFutureType
	// 期货单位帐号
	BankSecuAcc types.TThostFtdcBankAccountType
	// 银行密码标志
	BankPwdFlag types.TThostFtdcPwdFlagType
	// 期货资金密码核对标志
	SecuPwdFlag types.TThostFtdcPwdFlagType
	// 交易柜员
	OperNo types.TThostFtdcOperNoType
	// 交易ID
	TID types.TThostFtdcTIDType
	// 用户标识
	UserID types.TThostFtdcUserIDType
	// 错误代码
	ErrorID types.TThostFtdcErrorIDType
	// 错误信息
	ErrorMsg types.TThostFtdcErrorMsgType
}

// 银期变更银行账号信息
type CThostFtdcChangeAccountField struct {
	// 业务功能码
	TradeCode types.TThostFtdcTradeCodeType
	// 银行代码
	BankID types.TThostFtdcBankIDType
	// 银行分支机构代码
	BankBranchID types.TThostFtdcBankBrchIDType
	// 期商代码
	BrokerID types.TThostFtdcBrokerIDType
	// 期商分支机构代码
	BrokerBranchID types.TThostFtdcFutureBranchIDType
	// 交易日期
	TradeDate types.TThostFtdcTradeDateType
	// 交易时间
	TradeTime types.TThostFtdcTradeTimeType
	// 银行流水号
	BankSerial types.TThostFtdcBankSerialType
	// 交易系统日期
	TradingDay types.TThostFtdcTradeDateType
	// 银期平台消息流水号
	PlateSerial types.TThostFtdcSerialType
	// 最后分片标志
	LastFragment types.TThostFtdcLastFragmentType
	// 会话号
	SessionID types.TThostFtdcSessionIDType
	// 客户姓名
	CustomerName types.TThostFtdcIndividualNameType
	// 证件类型
	IdCardType types.TThostFtdcIdCardTypeType
	// 证件号码
	IdentifiedCardNo types.TThostFtdcIdentifiedCardNoType
	// 性别
	Gender types.TThostFtdcGenderType
	// 国家代码
	CountryCode types.TThostFtdcCountryCodeType
	// 客户类型
	CustType types.TThostFtdcCustTypeType
	// 地址
	Address types.TThostFtdcAddressType
	// 邮编
	ZipCode types.TThostFtdcZipCodeType
	// 电话号码
	Telephone types.TThostFtdcTelephoneType
	// 手机
	MobilePhone types.TThostFtdcMobilePhoneType
	// 传真
	Fax types.TThostFtdcFaxType
	// 电子邮件
	EMail types.TThostFtdcEMailType
	// 资金账户状态
	MoneyAccountStatus types.TThostFtdcMoneyAccountStatusType
	// 银行帐号
	BankAccount types.TThostFtdcBankAccountType
	// 银行密码
	BankPassWord types.TThostFtdcPasswordType
	// 新银行帐号
	NewBankAccount types.TThostFtdcBankAccountType
	// 新银行密码
	NewBankPassWord types.TThostFtdcPasswordType
	// 投资者帐号
	AccountID types.TThostFtdcAccountIDType
	// 期货密码
	Password types.TThostFtdcPasswordType
	// 银行帐号类型
	BankAccType types.TThostFtdcBankAccTypeType
	// 安装编号
	InstallID types.TThostFtdcInstallIDType
	// 验证客户证件号码标志
	VerifyCertNoFlag types.TThostFtdcYesNoIndicatorType
	// 币种代码
	CurrencyID types.TThostFtdcCurrencyIDType
	// 期货公司银行编码
	BrokerIDByBank types.TThostFtdcBankCodingForFutureType
	// 银行密码标志
	BankPwdFlag types.TThostFtdcPwdFlagType
	// 期货资金密码核对标志
	SecuPwdFlag types.TThostFtdcPwdFlagType
	// 交易ID
	TID types.TThostFtdcTIDType
	// 摘要
	Digest types.TThostFtdcDigestType
	// 错误代码
	ErrorID types.TThostFtdcErrorIDType
	// 错误信息
	ErrorMsg types.TThostFtdcErrorMsgType
}

// 二级代理操作员银期权限
type CThostFtdcSecAgentACIDMapField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 用户代码
	UserID types.TThostFtdcUserIDType
	// 资金账户
	AccountID types.TThostFtdcAccountIDType
	// 币种
	CurrencyID types.TThostFtdcCurrencyIDType
	// 境外中介机构资金帐号
	BrokerSecAgentID types.TThostFtdcAccountIDType
}

// 二级代理操作员银期权限查询
type CThostFtdcQrySecAgentACIDMapField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 用户代码
	UserID types.TThostFtdcUserIDType
	// 资金账户
	AccountID types.TThostFtdcAccountIDType
	// 币种
	CurrencyID types.TThostFtdcCurrencyIDType
}

// 灾备中心交易权限
type CThostFtdcUserRightsAssignField struct {
	// 应用单元代码
	BrokerID types.TThostFtdcBrokerIDType
	// 用户代码
	UserID types.TThostFtdcUserIDType
	// 交易中心代码
	DRIdentityID types.TThostFtdcDRIdentityIDType
}

// 经济公司是否有在本标示的交易权限
type CThostFtdcBrokerUserRightAssignField struct {
	// 应用单元代码
	BrokerID types.TThostFtdcBrokerIDType
	// 交易中心代码
	DRIdentityID types.TThostFtdcDRIdentityIDType
	// 能否交易
	Tradeable types.TThostFtdcBoolType
}

// 灾备交易转换报文
type CThostFtdcDRTransferField struct {
	// 原交易中心代码
	OrigDRIdentityID types.TThostFtdcDRIdentityIDType
	// 目标交易中心代码
	DestDRIdentityID types.TThostFtdcDRIdentityIDType
	// 原应用单元代码
	OrigBrokerID types.TThostFtdcBrokerIDType
	// 目标易用单元代码
	DestBrokerID types.TThostFtdcBrokerIDType
}

// Fens用户信息
type CThostFtdcFensUserInfoField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 用户代码
	UserID types.TThostFtdcUserIDType
	// 登录模式
	LoginMode types.TThostFtdcLoginModeType
}

// 当前银期所属交易中心
type CThostFtdcCurrTransferIdentityField struct {
	// 交易中心代码
	IdentityID types.TThostFtdcDRIdentityIDType
}

// 禁止登录用户
type CThostFtdcLoginForbiddenUserField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 用户代码
	UserID types.TThostFtdcUserIDType
	// IP地址
	IPAddress types.TThostFtdcIPAddressType
}

// 查询禁止登录用户
type CThostFtdcQryLoginForbiddenUserField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 用户代码
	UserID types.TThostFtdcUserIDType
}

// UDP组播组信息
type CThostFtdcMulticastGroupInfoField struct {
	// 组播组IP地址
	GroupIP types.TThostFtdcIPAddressType
	// 组播组IP端口
	GroupPort types.TThostFtdcIPPortType
	// 源地址
	SourceIP types.TThostFtdcIPAddressType
}

// 资金账户基本准备金
type CThostFtdcTradingAccountReserveField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 投资者帐号
	AccountID types.TThostFtdcAccountIDType
	// 基本准备金
	Reserve types.TThostFtdcMoneyType
	// 币种代码
	CurrencyID types.TThostFtdcCurrencyIDType
}

// 银期预约开户确认请求
type CThostFtdcReserveOpenAccountConfirmField struct {
	// 业务功能码
	TradeCode types.TThostFtdcTradeCodeType
	// 银行代码
	BankID types.TThostFtdcBankIDType
	// 银行分支机构代码
	BankBranchID types.TThostFtdcBankBrchIDType
	// 期商代码
	BrokerID types.TThostFtdcBrokerIDType
	// 期商分支机构代码
	BrokerBranchID types.TThostFtdcFutureBranchIDType
	// 交易日期
	TradeDate types.TThostFtdcTradeDateType
	// 交易时间
	TradeTime types.TThostFtdcTradeTimeType
	// 银行流水号
	BankSerial types.TThostFtdcBankSerialType
	// 交易系统日期
	TradingDay types.TThostFtdcTradeDateType
	// 银期平台消息流水号
	PlateSerial types.TThostFtdcSerialType
	// 最后分片标志
	LastFragment types.TThostFtdcLastFragmentType
	// 会话号
	SessionID types.TThostFtdcSessionIDType
	// 客户姓名
	CustomerName types.TThostFtdcIndividualNameType
	// 证件类型
	IdCardType types.TThostFtdcIdCardTypeType
	// 证件号码
	IdentifiedCardNo types.TThostFtdcIdentifiedCardNoType
	// 性别
	Gender types.TThostFtdcGenderType
	// 国家代码
	CountryCode types.TThostFtdcCountryCodeType
	// 客户类型
	CustType types.TThostFtdcCustTypeType
	// 地址
	Address types.TThostFtdcAddressType
	// 邮编
	ZipCode types.TThostFtdcZipCodeType
	// 电话号码
	Telephone types.TThostFtdcTelephoneType
	// 手机
	MobilePhone types.TThostFtdcMobilePhoneType
	// 传真
	Fax types.TThostFtdcFaxType
	// 电子邮件
	EMail types.TThostFtdcEMailType
	// 资金账户状态
	MoneyAccountStatus types.TThostFtdcMoneyAccountStatusType
	// 银行帐号
	BankAccount types.TThostFtdcBankAccountType
	// 银行密码
	BankPassWord types.TThostFtdcPasswordType
	// 安装编号
	InstallID types.TThostFtdcInstallIDType
	// 验证客户证件号码标志
	VerifyCertNoFlag types.TThostFtdcYesNoIndicatorType
	// 币种代码
	CurrencyID types.TThostFtdcCurrencyIDType
	// 摘要
	Digest types.TThostFtdcDigestType
	// 银行帐号类型
	BankAccType types.TThostFtdcBankAccTypeType
	// 期货公司银行编码
	BrokerIDByBank types.TThostFtdcBankCodingForFutureType
	// 交易ID
	TID types.TThostFtdcTIDType
	// 投资者帐号
	AccountID types.TThostFtdcAccountIDType
	// 期货密码
	Password types.TThostFtdcPasswordType
	// 预约开户银行流水号
	BankReserveOpenSeq types.TThostFtdcBankSerialType
	// 预约开户日期
	BookDate types.TThostFtdcTradeDateType
	// 预约开户验证密码
	BookPsw types.TThostFtdcPasswordType
	// 错误代码
	ErrorID types.TThostFtdcErrorIDType
	// 错误信息
	ErrorMsg types.TThostFtdcErrorMsgType
}

// 银期预约开户
type CThostFtdcReserveOpenAccountField struct {
	// 业务功能码
	TradeCode types.TThostFtdcTradeCodeType
	// 银行代码
	BankID types.TThostFtdcBankIDType
	// 银行分支机构代码
	BankBranchID types.TThostFtdcBankBrchIDType
	// 期商代码
	BrokerID types.TThostFtdcBrokerIDType
	// 期商分支机构代码
	BrokerBranchID types.TThostFtdcFutureBranchIDType
	// 交易日期
	TradeDate types.TThostFtdcTradeDateType
	// 交易时间
	TradeTime types.TThostFtdcTradeTimeType
	// 银行流水号
	BankSerial types.TThostFtdcBankSerialType
	// 交易系统日期
	TradingDay types.TThostFtdcTradeDateType
	// 银期平台消息流水号
	PlateSerial types.TThostFtdcSerialType
	// 最后分片标志
	LastFragment types.TThostFtdcLastFragmentType
	// 会话号
	SessionID types.TThostFtdcSessionIDType
	// 客户姓名
	CustomerName types.TThostFtdcIndividualNameType
	// 证件类型
	IdCardType types.TThostFtdcIdCardTypeType
	// 证件号码
	IdentifiedCardNo types.TThostFtdcIdentifiedCardNoType
	// 性别
	Gender types.TThostFtdcGenderType
	// 国家代码
	CountryCode types.TThostFtdcCountryCodeType
	// 客户类型
	CustType types.TThostFtdcCustTypeType
	// 地址
	Address types.TThostFtdcAddressType
	// 邮编
	ZipCode types.TThostFtdcZipCodeType
	// 电话号码
	Telephone types.TThostFtdcTelephoneType
	// 手机
	MobilePhone types.TThostFtdcMobilePhoneType
	// 传真
	Fax types.TThostFtdcFaxType
	// 电子邮件
	EMail types.TThostFtdcEMailType
	// 资金账户状态
	MoneyAccountStatus types.TThostFtdcMoneyAccountStatusType
	// 银行帐号
	BankAccount types.TThostFtdcBankAccountType
	// 银行密码
	BankPassWord types.TThostFtdcPasswordType
	// 安装编号
	InstallID types.TThostFtdcInstallIDType
	// 验证客户证件号码标志
	VerifyCertNoFlag types.TThostFtdcYesNoIndicatorType
	// 币种代码
	CurrencyID types.TThostFtdcCurrencyIDType
	// 摘要
	Digest types.TThostFtdcDigestType
	// 银行帐号类型
	BankAccType types.TThostFtdcBankAccTypeType
	// 期货公司银行编码
	BrokerIDByBank types.TThostFtdcBankCodingForFutureType
	// 交易ID
	TID types.TThostFtdcTIDType
	// 预约开户状态
	ReserveOpenAccStas types.TThostFtdcReserveOpenAccStasType
	// 错误代码
	ErrorID types.TThostFtdcErrorIDType
	// 错误信息
	ErrorMsg types.TThostFtdcErrorMsgType
}

// 分价表
type CThostFtdcMBLMarketDataField struct {
	// 合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
	// 交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	// 方向
	Direction types.TThostFtdcDirectionType
	// 价格
	Price types.TThostFtdcPriceType
	// 数量
	Volume types.TThostFtdcVolumeType
}

// 用户系统信息
type CThostFtdcUserSystemInfoField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 用户代码
	UserID types.TThostFtdcUserIDType
	// 用户端系统内部信息长度
	ClientSystemInfoLen types.TThostFtdcSystemInfoLenType
	// 用户端系统内部信息
	ClientSystemInfo types.TThostFtdcClientSystemInfoType
	// 用户公网IP
	ClientPublicIP types.TThostFtdcIPAddressType
	// 终端IP端口
	ClientIPPort types.TThostFtdcIPPortType
	// 登录成功时间
	ClientLoginTime types.TThostFtdcTimeType
	// App代码
	ClientAppID types.TThostFtdcClientAppIDType
}

// 指定的席位
type CThostFtdcSpecificTraderField struct {
	// 席位代码
	TraderID types.TThostFtdcTraderIDType
}

// 流控警告通知
type CThostFtdcFlowCtrlWarningField struct {
	// 交易所交易员代码
	TraderID types.TThostFtdcTraderIDType
	// 触发时间
	TriggerTime types.TThostFtdcLongTimeType
	// tid组号
	Tgid uint32
	// 瞬时流量值
	CurPkgCnt uint32
}

// 订阅资金变动应答
type CThostFtdcRequestIDEntityField struct {
	// 席位代码
	RequestID types.TThostFtdcRequestIDType
}

// 查询投资者申报费相关设置
type CThostFtdcQryInvestorInfoCntSettingField struct {
	// 交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	// 投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	// 产品代码
	ProductID types.TThostFtdcProductIDType
}

// 投资者申报费相关设置
type CThostFtdcInvestorInfoCntSettingField struct {
	// 交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	// 产品代码
	ProductID types.TThostFtdcProductIDType
	// 是否收取申报费
	IsCalInfoComm types.TThostFtdcBoolType
	// 是否限制信息量
	IsLimitInfoMax types.TThostFtdcBoolType
	// 信息量限制笔数
	InfoMaxLimit types.TThostFtdcVolumeType
}

// 新风控支持临时关闭或开启交易权限
type CThostFtdcUpdRiskForbiddenRightField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	// 产品代码
	InstrumentID types.TThostFtdcInstrumentIDType
	// 用户代码
	UserID types.TThostFtdcUserIDType
	// 业务操作类型
	EventType types.TThostFtdcOperateType
	// 交易权限
	TradingRight types.TThostFtdcTradingRightType
}

// 输入的对冲设置
type CThostFtdcInputOffsetSettingField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	// 合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
	// 标的期货合约代码
	UnderlyingInstrID types.TThostFtdcInstrumentIDType
	// 产品代码
	ProductID types.TThostFtdcProductIDType
	// 对冲类型
	OffsetType types.TThostFtdcOffsetTypeType
	// 申请对冲的合约数量
	Volume types.TThostFtdcVolumeType
	// 是否对冲
	IsOffset types.TThostFtdcBoolType
	// 请求编号
	RequestID types.TThostFtdcRequestIDType
	// 用户代码
	UserID types.TThostFtdcUserIDType
	// 交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	// 业务单元
	BusinessUnit types.TThostFtdcBusinessUnitType
	// IP地址
	IPAddress types.TThostFtdcIPAddressType
	// Mac地址
	MacAddress types.TThostFtdcMacAddressType
}

// 对冲设置
type CThostFtdcOffsetSettingField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	// 合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
	// 标的期货合约代码
	UnderlyingInstrID types.TThostFtdcInstrumentIDType
	// 产品代码
	ProductID types.TThostFtdcProductIDType
	// 对冲类型
	OffsetType types.TThostFtdcOffsetTypeType
	// 申请对冲的合约数量
	Volume types.TThostFtdcVolumeType
	// 是否对冲
	IsOffset types.TThostFtdcBoolType
	// 请求编号
	RequestID types.TThostFtdcRequestIDType
	// 用户代码
	UserID types.TThostFtdcUserIDType
	// 交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	// 业务单元
	BusinessUnit types.TThostFtdcBusinessUnitType
	// IP地址
	IPAddress types.TThostFtdcIPAddressType
	// Mac地址
	MacAddress types.TThostFtdcMacAddressType
	// 交易所合约代码
	ExchangeInstID types.TThostFtdcExchangeInstIDType
	// 交易所期权系列号
	ExchangeSerialNo types.TThostFtdcExchangeInstIDType
	// 交易所产品代码
	ExchangeProductID types.TThostFtdcProductIDType
	// 会员代码
	ParticipantID types.TThostFtdcParticipantIDType
	// 客户代码
	ClientID types.TThostFtdcClientIDType
	// 交易所交易员代码
	TraderID types.TThostFtdcTraderIDType
	// 安装编号
	InstallID types.TThostFtdcInstallIDType
	// 对冲提交状态
	OrderSubmitStatus types.TThostFtdcOrderSubmitStatusType
	// 交易日
	TradingDay types.TThostFtdcDateType
	// 结算编号
	SettlementID types.TThostFtdcSettlementIDType
	// 报单日期
	InsertDate types.TThostFtdcDateType
	// 插入时间
	InsertTime types.TThostFtdcTimeType
	// 撤销时间
	CancelTime types.TThostFtdcTimeType
	// 对冲设置结果
	ExecResult types.TThostFtdcExecResultType
	// 序号
	SequenceNo types.TThostFtdcSequenceNoType
	// 前置编号
	FrontID types.TThostFtdcFrontIDType
	// 会话编号
	SessionID types.TThostFtdcSessionIDType
	// 状态信息
	StatusMsg types.TThostFtdcErrorMsgType
	// 操作用户代码
	ActiveUserID types.TThostFtdcUserIDType
	// 经纪公司报单编号
	BrokerOffsetSettingSeq types.TThostFtdcSequenceNoType
	// 申请来源
	ApplySrc types.TThostFtdcApplySrcType
}

// 撤销对冲设置
type CThostFtdcCancelOffsetSettingField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	// 合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
	// 标的期货合约代码
	UnderlyingInstrID types.TThostFtdcInstrumentIDType
	// 产品代码
	ProductID types.TThostFtdcProductIDType
	// 对冲类型
	OffsetType types.TThostFtdcOffsetTypeType
	// 申请对冲的合约数量
	Volume types.TThostFtdcVolumeType
	// 是否对冲
	IsOffset types.TThostFtdcBoolType
	// 请求编号
	RequestID types.TThostFtdcRequestIDType
	// 用户代码
	UserID types.TThostFtdcUserIDType
	// 交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	// IP地址
	IPAddress types.TThostFtdcIPAddressType
	// Mac地址
	MacAddress types.TThostFtdcMacAddressType
	// 交易所合约代码
	ExchangeInstID types.TThostFtdcExchangeInstIDType
	// 交易所期权系列号
	ExchangeSerialNo types.TThostFtdcExchangeInstIDType
	// 交易所产品代码
	ExchangeProductID types.TThostFtdcProductIDType
	// 交易所交易员代码
	TraderID types.TThostFtdcTraderIDType
	// 安装编号
	InstallID types.TThostFtdcInstallIDType
	// 会员代码
	ParticipantID types.TThostFtdcParticipantIDType
	// 客户代码
	ClientID types.TThostFtdcClientIDType
	// 报单操作状态
	OrderActionStatus types.TThostFtdcOrderActionStatusType
	// 状态信息
	StatusMsg types.TThostFtdcErrorMsgType
	// 操作本地编号
	ActionLocalID types.TThostFtdcOrderLocalIDType
	// 操作日期
	ActionDate types.TThostFtdcDateType
	// 操作时间
	ActionTime types.TThostFtdcTimeType
}

// 查询对冲设置
type CThostFtdcQryOffsetSettingField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	// 产品代码
	ProductID types.TThostFtdcProductIDType
	// 对冲类型
	OffsetType types.TThostFtdcOffsetTypeType
}

// 查询IP接入控制
type CThostFtdcQryIPUserACLField struct {
	// 校验模式
	Mode types.TThostFtdcIPACLCheckModeType
	// IP地址
	IPAddress types.TThostFtdcIPAddressType
	// 用户代码
	UserID types.TThostFtdcUserIDType
}

// IP接入控制的查询结果
type CThostFtdcIPUserACLField struct {
	// 校验模式
	Mode types.TThostFtdcIPACLCheckModeType
	// IP地址
	IPAddress types.TThostFtdcIPAddressType
	// 用户代码
	UserID types.TThostFtdcUserIDType
	// 阈值
	Threshold types.TThostFtdcRecordCountType
	// 错误计数
	Count types.TThostFtdcRecordCountType
	// 禁止标志
	IsForbidden types.TThostFtdcBoolType
}

// 更新IP接入控制
type CThostFtdcUpdIPUserACLField struct {
	// 校验模式
	Mode types.TThostFtdcIPACLCheckModeType
	// IP地址
	IPAddress types.TThostFtdcIPAddressType
	// 用户代码
	UserID types.TThostFtdcUserIDType
	// 阈值
	Threshold types.TThostFtdcRecordCountType
	// 业务操作类型
	EventType types.TThostFtdcOperateType
	// 是否写入文件
	IsSave types.TThostFtdcBoolType
}
