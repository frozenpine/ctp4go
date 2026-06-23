package thost

import "github.com/frozenpine/ctp4go/thost/types"

// 信息分发
type CThostFtdcDisseminationField struct {
	//  序列系列号
	SequenceSeries types.TThostFtdcSequenceSeriesType
	//  序列号
	SequenceNo types.TThostFtdcSequenceNoType
}

// 用户登录请求
type CThostFtdcReqUserLoginField struct {
	//  交易日
	TradingDay types.TThostFtdcDateType
	//  经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	//  用户代码
	UserID types.TThostFtdcUserIDType
	//  密码
	Password types.TThostFtdcPasswordType
	//  用户端产品信息
	UserProductInfo types.TThostFtdcProductInfoType
	//  接口端产品信息
	InterfaceProductInfo types.TThostFtdcProductInfoType
	//  协议信息
	ProtocolInfo types.TThostFtdcProtocolInfoType
	//  Mac地址
	MacAddress types.TThostFtdcMacAddressType
	//  动态密码
	OneTimePassword types.TThostFtdcPasswordType
	//  保留的无效字段
	Reserve1 types.TThostFtdcOldIPAddressType
	//  登录备注
	LoginRemark types.TThostFtdcLoginRemarkType
	//  终端IP端口
	ClientIPPort types.TThostFtdcIPPortType
	//  终端IP地址
	ClientIPAddress types.TThostFtdcIPAddressType
	//  短信验证码
	SMSCode types.TThostFtdcSMSCodeType
}

// 用户登录应答
type CThostFtdcRspUserLoginField struct {
	//  交易日
	TradingDay types.TThostFtdcDateType
	//  登录成功时间
	LoginTime types.TThostFtdcTimeType
	//  经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	//  用户代码
	UserID types.TThostFtdcUserIDType
	//  交易系统名称
	SystemName types.TThostFtdcSystemNameType
	//  前置编号
	FrontID types.TThostFtdcFrontIDType
	//  会话编号
	SessionID types.TThostFtdcSessionIDType
	//  最大报单引用
	MaxOrderRef types.TThostFtdcOrderRefType
	//  上期所时间
	SHFETime types.TThostFtdcTimeType
	//  大商所时间
	DCETime types.TThostFtdcTimeType
	//  郑商所时间
	CZCETime types.TThostFtdcTimeType
	//  中金所时间
	FFEXTime types.TThostFtdcTimeType
	//  能源中心时间
	INETime types.TThostFtdcTimeType
	//  后台版本信息
	SysVersion types.TThostFtdcSysVersionType
	//  广期所时间
	GFEXTime types.TThostFtdcTimeType
	//  当前登录中心号
	LoginDRIdentityID types.TThostFtdcDRIdentityIDType
	//  用户所属中心号
	UserDRIdentityID types.TThostFtdcDRIdentityIDType
	//  上次登陆时间
	LastLoginTime types.TThostFtdcDateTimeType
	//  预留信息
	ReserveInfo types.TThostFtdcReserveInfoType
}

// 用户登出请求
type CThostFtdcUserLogoutField struct {
	//  经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	//  用户代码
	UserID types.TThostFtdcUserIDType
}

// 强制交易员退出
type CThostFtdcForceUserLogoutField struct {
	//  经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	//  用户代码
	UserID types.TThostFtdcUserIDType
}

// 客户端认证请求
type CThostFtdcReqAuthenticateField struct {
	//  经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	//  用户代码
	UserID types.TThostFtdcUserIDType
	//  用户端产品信息
	UserProductInfo types.TThostFtdcProductInfoType
	//  认证码
	AuthCode types.TThostFtdcAuthCodeType
	//  App代码
	AppID types.TThostFtdcAppIDType
}

// 客户端认证响应
type CThostFtdcRspAuthenticateField struct {
	//  经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	//  用户代码
	UserID types.TThostFtdcUserIDType
	//  用户端产品信息
	UserProductInfo types.TThostFtdcProductInfoType
	//  App代码
	AppID types.TThostFtdcAppIDType
	//  App类型
	AppType types.TThostFtdcAppTypeType
}

// 客户端认证信息
type CThostFtdcAuthenticationInfoField struct {
	//  经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	//  用户代码
	UserID types.TThostFtdcUserIDType
	//  用户端产品信息
	UserProductInfo types.TThostFtdcProductInfoType
	//  认证信息
	AuthInfo types.TThostFtdcAuthInfoType
	//  是否为认证结果
	IsResult types.TThostFtdcBoolType
	//  App代码
	AppID types.TThostFtdcAppIDType
	//  App类型
	AppType types.TThostFtdcAppTypeType
	//  保留的无效字段
	Reserve1 types.TThostFtdcOldIPAddressType
	//  终端IP地址
	ClientIPAddress types.TThostFtdcIPAddressType
}

// 用户登录应答2
type CThostFtdcRspUserLogin2Field struct {
	//  交易日
	TradingDay types.TThostFtdcDateType
	//  登录成功时间
	LoginTime types.TThostFtdcTimeType
	//  经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	//  用户代码
	UserID types.TThostFtdcUserIDType
	//  交易系统名称
	SystemName types.TThostFtdcSystemNameType
	//  前置编号
	FrontID types.TThostFtdcFrontIDType
	//  会话编号
	SessionID types.TThostFtdcSessionIDType
	//  最大报单引用
	MaxOrderRef types.TThostFtdcOrderRefType
	//  上期所时间
	SHFETime types.TThostFtdcTimeType
	//  大商所时间
	DCETime types.TThostFtdcTimeType
	//  郑商所时间
	CZCETime types.TThostFtdcTimeType
	//  中金所时间
	FFEXTime types.TThostFtdcTimeType
	//  能源中心时间
	INETime types.TThostFtdcTimeType
	//  随机串
	RandomString types.TThostFtdcRandomStringType
}

// 银期转帐报文头
type CThostFtdcTransferHeaderField struct {
	//  版本号，常量，1.0
	Version types.TThostFtdcVersionType
	//  交易代码，必填
	TradeCode types.TThostFtdcTradeCodeType
	//  交易日期，必填，格式：yyyymmdd
	TradeDate types.TThostFtdcTradeDateType
	//  交易时间，必填，格式：hhmmss
	TradeTime types.TThostFtdcTradeTimeType
	//  发起方流水号，N/A
	TradeSerial types.TThostFtdcTradeSerialType
	//  期货公司代码，必填
	FutureID types.TThostFtdcFutureIDType
	//  银行代码，根据查询银行得到，必填
	BankID types.TThostFtdcBankIDType
	//  银行分中心代码，根据查询银行得到，必填
	BankBrchID types.TThostFtdcBankBrchIDType
	//  操作员，N/A
	OperNo types.TThostFtdcOperNoType
	//  交易设备类型，N/A
	DeviceID types.TThostFtdcDeviceIDType
	//  记录数，N/A
	RecordNum types.TThostFtdcRecordNumType
	//  会话编号，N/A
	SessionID types.TThostFtdcSessionIDType
	//  请求编号，N/A
	RequestID types.TThostFtdcRequestIDType
}

// 银行资金转期货请求，TradeCode=202001
type CThostFtdcTransferBankToFutureReqField struct {
	//  期货资金账户
	FutureAccount types.TThostFtdcAccountIDType
	//  密码标志
	FuturePwdFlag types.TThostFtdcFuturePwdFlagType
	//  密码
	FutureAccPwd types.TThostFtdcFutureAccPwdType
	//  转账金额
	TradeAmt types.TThostFtdcMoneyType
	//  客户手续费
	CustFee types.TThostFtdcMoneyType
	//  币种：RMB-人民币 USD-美圆 HKD-港元
	CurrencyCode types.TThostFtdcCurrencyCodeType
}

// 银行资金转期货请求响应
type CThostFtdcTransferBankToFutureRspField struct {
	//  响应代码
	RetCode types.TThostFtdcRetCodeType
	//  响应信息
	RetInfo types.TThostFtdcRetInfoType
	//  资金账户
	FutureAccount types.TThostFtdcAccountIDType
	//  转帐金额
	TradeAmt types.TThostFtdcMoneyType
	//  应收客户手续费
	CustFee types.TThostFtdcMoneyType
	//  币种
	CurrencyCode types.TThostFtdcCurrencyCodeType
}

// 期货资金转银行请求，TradeCode=202002
type CThostFtdcTransferFutureToBankReqField struct {
	//  期货资金账户
	FutureAccount types.TThostFtdcAccountIDType
	//  密码标志
	FuturePwdFlag types.TThostFtdcFuturePwdFlagType
	//  密码
	FutureAccPwd types.TThostFtdcFutureAccPwdType
	//  转账金额
	TradeAmt types.TThostFtdcMoneyType
	//  客户手续费
	CustFee types.TThostFtdcMoneyType
	//  币种：RMB-人民币 USD-美圆 HKD-港元
	CurrencyCode types.TThostFtdcCurrencyCodeType
}

// 期货资金转银行请求响应
type CThostFtdcTransferFutureToBankRspField struct {
	//  响应代码
	RetCode types.TThostFtdcRetCodeType
	//  响应信息
	RetInfo types.TThostFtdcRetInfoType
	//  资金账户
	FutureAccount types.TThostFtdcAccountIDType
	//  转帐金额
	TradeAmt types.TThostFtdcMoneyType
	//  应收客户手续费
	CustFee types.TThostFtdcMoneyType
	//  币种
	CurrencyCode types.TThostFtdcCurrencyCodeType
}

// 查询银行资金请求，TradeCode=204002
type CThostFtdcTransferQryBankReqField struct {
	//  期货资金账户
	FutureAccount types.TThostFtdcAccountIDType
	//  密码标志
	FuturePwdFlag types.TThostFtdcFuturePwdFlagType
	//  密码
	FutureAccPwd types.TThostFtdcFutureAccPwdType
	//  币种：RMB-人民币 USD-美圆 HKD-港元
	CurrencyCode types.TThostFtdcCurrencyCodeType
}

// 查询银行资金请求响应
type CThostFtdcTransferQryBankRspField struct {
	//  响应代码
	RetCode types.TThostFtdcRetCodeType
	//  响应信息
	RetInfo types.TThostFtdcRetInfoType
	//  资金账户
	FutureAccount types.TThostFtdcAccountIDType
	//  银行余额
	TradeAmt types.TThostFtdcMoneyType
	//  银行可用余额
	UseAmt types.TThostFtdcMoneyType
	//  银行可取余额
	FetchAmt types.TThostFtdcMoneyType
	//  币种
	CurrencyCode types.TThostFtdcCurrencyCodeType
}

// 查询银行交易明细请求，TradeCode=204999
type CThostFtdcTransferQryDetailReqField struct {
	//  期货资金账户
	FutureAccount types.TThostFtdcAccountIDType
}

// 查询银行交易明细请求响应
type CThostFtdcTransferQryDetailRspField struct {
	//  交易日期
	TradeDate types.TThostFtdcDateType
	//  交易时间
	TradeTime types.TThostFtdcTradeTimeType
	//  交易代码
	TradeCode types.TThostFtdcTradeCodeType
	//  期货流水号
	FutureSerial types.TThostFtdcTradeSerialNoType
	//  期货公司代码
	FutureID types.TThostFtdcFutureIDType
	//  资金帐号
	FutureAccount types.TThostFtdcFutureAccountType
	//  银行流水号
	BankSerial types.TThostFtdcTradeSerialNoType
	//  银行代码
	BankID types.TThostFtdcBankIDType
	//  银行分中心代码
	BankBrchID types.TThostFtdcBankBrchIDType
	//  银行账号
	BankAccount types.TThostFtdcBankAccountType
	//  证件号码
	CertCode types.TThostFtdcCertCodeType
	//  货币代码
	CurrencyCode types.TThostFtdcCurrencyCodeType
	//  发生金额
	TxAmount types.TThostFtdcMoneyType
	//  有效标志
	Flag types.TThostFtdcTransferValidFlagType
}

// 响应信息
type CThostFtdcRspInfoField struct {
	//  错误代码
	ErrorID types.TThostFtdcErrorIDType
	//  错误信息
	ErrorMsg types.TThostFtdcErrorMsgType
}

// 交易所
type CThostFtdcExchangeField struct {
	//  交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	//  交易所名称
	ExchangeName types.TThostFtdcExchangeNameType
	//  交易所属性
	ExchangeProperty types.TThostFtdcExchangePropertyType
}

// 产品
type CThostFtdcProductField struct {
	//  保留的无效字段
	Reserve1 types.TThostFtdcOldInstrumentIDType
	//  产品名称
	ProductName types.TThostFtdcProductNameType
	//  交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	//  产品类型
	ProductClass types.TThostFtdcProductClassType
	//  合约数量乘数
	VolumeMultiple types.TThostFtdcVolumeMultipleType
	//  最小变动价位
	PriceTick types.TThostFtdcPriceType
	//  市价单最大下单量
	MaxMarketOrderVolume types.TThostFtdcVolumeType
	//  市价单最小下单量
	MinMarketOrderVolume types.TThostFtdcVolumeType
	//  限价单最大下单量
	MaxLimitOrderVolume types.TThostFtdcVolumeType
	//  限价单最小下单量
	MinLimitOrderVolume types.TThostFtdcVolumeType
	//  持仓类型
	PositionType types.TThostFtdcPositionTypeType
	//  持仓日期类型
	PositionDateType types.TThostFtdcPositionDateTypeType
	//  平仓处理类型
	CloseDealType types.TThostFtdcCloseDealTypeType
	//  交易币种类型
	TradeCurrencyID types.TThostFtdcCurrencyIDType
	//  质押资金可用范围
	MortgageFundUseRange types.TThostFtdcMortgageFundUseRangeType
	//  保留的无效字段
	Reserve2 types.TThostFtdcOldInstrumentIDType
	//  合约基础商品乘数
	UnderlyingMultiple types.TThostFtdcUnderlyingMultipleType
	//  产品代码
	ProductID types.TThostFtdcInstrumentIDType
	//  交易所产品代码
	ExchangeProductID types.TThostFtdcInstrumentIDType
	//  开仓量限制粒度
	OpenLimitControlLevel types.TThostFtdcOpenLimitControlLevelType
	//  报单频率控制粒度
	OrderFreqControlLevel types.TThostFtdcOrderFreqControlLevelType
}

// 合约
type CThostFtdcInstrumentField struct {
	//  保留的无效字段
	Reserve1 types.TThostFtdcOldInstrumentIDType
	//  交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	//  合约名称
	InstrumentName types.TThostFtdcInstrumentNameType
	//  保留的无效字段
	Reserve2 types.TThostFtdcOldExchangeInstIDType
	//  保留的无效字段
	Reserve3 types.TThostFtdcOldInstrumentIDType
	//  产品类型
	ProductClass types.TThostFtdcProductClassType
	//  交割年份
	DeliveryYear types.TThostFtdcYearType
	//  交割月
	DeliveryMonth types.TThostFtdcMonthType
	//  市价单最大下单量
	MaxMarketOrderVolume types.TThostFtdcVolumeType
	//  市价单最小下单量
	MinMarketOrderVolume types.TThostFtdcVolumeType
	//  限价单最大下单量
	MaxLimitOrderVolume types.TThostFtdcVolumeType
	//  限价单最小下单量
	MinLimitOrderVolume types.TThostFtdcVolumeType
	//  合约数量乘数
	VolumeMultiple types.TThostFtdcVolumeMultipleType
	//  最小变动价位
	PriceTick types.TThostFtdcPriceType
	//  创建日
	CreateDate types.TThostFtdcDateType
	//  上市日
	OpenDate types.TThostFtdcDateType
	//  到期日
	ExpireDate types.TThostFtdcDateType
	//  开始交割日
	StartDelivDate types.TThostFtdcDateType
	//  结束交割日
	EndDelivDate types.TThostFtdcDateType
	//  合约生命周期状态
	InstLifePhase types.TThostFtdcInstLifePhaseType
	//  当前是否交易
	IsTrading types.TThostFtdcBoolType
	//  持仓类型
	PositionType types.TThostFtdcPositionTypeType
	//  持仓日期类型
	PositionDateType types.TThostFtdcPositionDateTypeType
	//  多头保证金率
	LongMarginRatio types.TThostFtdcRatioType
	//  空头保证金率
	ShortMarginRatio types.TThostFtdcRatioType
	//  是否使用大额单边保证金算法
	MaxMarginSideAlgorithm types.TThostFtdcMaxMarginSideAlgorithmType
	//  保留的无效字段
	Reserve4 types.TThostFtdcOldInstrumentIDType
	//  执行价
	StrikePrice types.TThostFtdcPriceType
	//  期权类型
	OptionsType types.TThostFtdcOptionsTypeType
	//  合约基础商品乘数
	UnderlyingMultiple types.TThostFtdcUnderlyingMultipleType
	//  组合类型
	CombinationType types.TThostFtdcCombinationTypeType
	//  合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
	//  合约在交易所的代码
	ExchangeInstID types.TThostFtdcExchangeInstIDType
	//  产品代码
	ProductID types.TThostFtdcInstrumentIDType
	//  基础商品代码
	UnderlyingInstrID types.TThostFtdcInstrumentIDType
}

// 经纪公司
type CThostFtdcBrokerField struct {
	//  经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	//  经纪公司简称
	BrokerAbbr types.TThostFtdcBrokerAbbrType
	//  经纪公司名称
	BrokerName types.TThostFtdcBrokerNameType
	//  是否活跃
	IsActive types.TThostFtdcBoolType
}

// 交易所交易员
type CThostFtdcTraderField struct {
	//  交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	//  交易所交易员代码
	TraderID types.TThostFtdcTraderIDType
	//  会员代码
	ParticipantID types.TThostFtdcParticipantIDType
	//  密码
	Password types.TThostFtdcPasswordType
	//  安装数量
	InstallCount types.TThostFtdcInstallCountType
	//  经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	//  撤单时选择席位算法
	OrderCancelAlg types.TThostFtdcOrderCancelAlgType
	//  交易报盘安装数量
	TradeInstallCount types.TThostFtdcInstallCountType
	//  行情报盘安装数量
	MDInstallCount types.TThostFtdcInstallCountType
}

// 投资者
type CThostFtdcInvestorField struct {
	//  投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	//  经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	//  投资者分组代码
	InvestorGroupID types.TThostFtdcInvestorIDType
	//  投资者名称
	InvestorName types.TThostFtdcPartyNameType
	//  证件类型
	IdentifiedCardType types.TThostFtdcIdCardTypeType
	//  证件号码
	IdentifiedCardNo types.TThostFtdcIdentifiedCardNoType
	//  是否活跃
	IsActive types.TThostFtdcBoolType
	//  联系电话
	Telephone types.TThostFtdcTelephoneType
	//  通讯地址
	Address types.TThostFtdcAddressType
	//  开户日期
	OpenDate types.TThostFtdcDateType
	//  手机
	Mobile types.TThostFtdcMobileType
	//  手续费率模板代码
	CommModelID types.TThostFtdcInvestorIDType
	//  保证金率模板代码
	MarginModelID types.TThostFtdcInvestorIDType
	//  是否频率控制
	IsOrderFreq types.TThostFtdcEnumBoolType
	//  是否开仓限制
	IsOpenVolLimit types.TThostFtdcEnumBoolType
}

// 交易编码
type CThostFtdcTradingCodeField struct {
	//  投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	//  经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	//  交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	//  客户代码
	ClientID types.TThostFtdcClientIDType
	//  是否活跃
	IsActive types.TThostFtdcBoolType
	//  交易编码类型
	ClientIDType types.TThostFtdcClientIDTypeType
	//  营业部编号
	BranchID types.TThostFtdcBranchIDType
	//  业务类型
	BizType types.TThostFtdcBizTypeType
	//  投资单元代码
	InvestUnitID types.TThostFtdcInvestUnitIDType
}

// 会员编码和经纪公司编码对照表
type CThostFtdcPartBrokerField struct {
	//  经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	//  交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	//  会员代码
	ParticipantID types.TThostFtdcParticipantIDType
	//  是否活跃
	IsActive types.TThostFtdcBoolType
}

// 管理用户
type CThostFtdcSuperUserField struct {
	//  用户代码
	UserID types.TThostFtdcUserIDType
	//  用户名称
	UserName types.TThostFtdcUserNameType
	//  密码
	Password types.TThostFtdcPasswordType
	//  是否活跃
	IsActive types.TThostFtdcBoolType
}

// 管理用户功能权限
type CThostFtdcSuperUserFunctionField struct {
	//  用户代码
	UserID types.TThostFtdcUserIDType
	//  功能代码
	FunctionCode types.TThostFtdcFunctionCodeType
}

// 投资者组
type CThostFtdcInvestorGroupField struct {
	//  经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	//  投资者分组代码
	InvestorGroupID types.TThostFtdcInvestorIDType
	//  投资者分组名称
	InvestorGroupName types.TThostFtdcInvestorGroupNameType
}

// 资金账户
type CThostFtdcTradingAccountField struct {
	//  经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	//  投资者帐号
	AccountID types.TThostFtdcAccountIDType
	//  上次质押金额
	PreMortgage types.TThostFtdcMoneyType
	//  上次信用额度
	PreCredit types.TThostFtdcMoneyType
	//  上次存款额
	PreDeposit types.TThostFtdcMoneyType
	//  上次结算准备金
	PreBalance types.TThostFtdcMoneyType
	//  上次占用的保证金
	PreMargin types.TThostFtdcMoneyType
	//  利息基数
	InterestBase types.TThostFtdcMoneyType
	//  利息收入
	Interest types.TThostFtdcMoneyType
	//  入金金额
	Deposit types.TThostFtdcMoneyType
	//  出金金额
	Withdraw types.TThostFtdcMoneyType
	//  冻结的保证金
	FrozenMargin types.TThostFtdcMoneyType
	//  冻结的资金
	FrozenCash types.TThostFtdcMoneyType
	//  冻结的手续费
	FrozenCommission types.TThostFtdcMoneyType
	//  当前保证金总额
	CurrMargin types.TThostFtdcMoneyType
	//  资金差额
	CashIn types.TThostFtdcMoneyType
	//  手续费
	Commission types.TThostFtdcMoneyType
	//  平仓盈亏
	CloseProfit types.TThostFtdcMoneyType
	//  持仓盈亏
	PositionProfit types.TThostFtdcMoneyType
	//  期货结算准备金
	Balance types.TThostFtdcMoneyType
	//  可用资金
	Available types.TThostFtdcMoneyType
	//  可取资金
	WithdrawQuota types.TThostFtdcMoneyType
	//  基本准备金
	Reserve types.TThostFtdcMoneyType
	//  交易日
	TradingDay types.TThostFtdcDateType
	//  结算编号
	SettlementID types.TThostFtdcSettlementIDType
	//  信用额度
	Credit types.TThostFtdcMoneyType
	//  质押金额
	Mortgage types.TThostFtdcMoneyType
	//  交易所保证金
	ExchangeMargin types.TThostFtdcMoneyType
	//  投资者交割保证金
	DeliveryMargin types.TThostFtdcMoneyType
	//  交易所交割保证金
	ExchangeDeliveryMargin types.TThostFtdcMoneyType
	//  保底期货结算准备金
	ReserveBalance types.TThostFtdcMoneyType
	//  币种代码
	CurrencyID types.TThostFtdcCurrencyIDType
	//  上次货币质入金额
	PreFundMortgageIn types.TThostFtdcMoneyType
	//  上次货币质出金额
	PreFundMortgageOut types.TThostFtdcMoneyType
	//  货币质入金额
	FundMortgageIn types.TThostFtdcMoneyType
	//  货币质出金额
	FundMortgageOut types.TThostFtdcMoneyType
	//  货币质押余额
	FundMortgageAvailable types.TThostFtdcMoneyType
	//  可质押货币金额
	MortgageableFund types.TThostFtdcMoneyType
	//  特殊产品占用保证金
	SpecProductMargin types.TThostFtdcMoneyType
	//  特殊产品冻结保证金
	SpecProductFrozenMargin types.TThostFtdcMoneyType
	//  特殊产品手续费
	SpecProductCommission types.TThostFtdcMoneyType
	//  特殊产品冻结手续费
	SpecProductFrozenCommission types.TThostFtdcMoneyType
	//  特殊产品持仓盈亏
	SpecProductPositionProfit types.TThostFtdcMoneyType
	//  特殊产品平仓盈亏
	SpecProductCloseProfit types.TThostFtdcMoneyType
	//  根据持仓盈亏算法计算的特殊产品持仓盈亏
	SpecProductPositionProfitByAlg types.TThostFtdcMoneyType
	//  特殊产品交易所保证金
	SpecProductExchangeMargin types.TThostFtdcMoneyType
	//  业务类型
	BizType types.TThostFtdcBizTypeType
	//  延时换汇冻结金额
	FrozenSwap types.TThostFtdcMoneyType
	//  剩余换汇额度
	RemainSwap types.TThostFtdcMoneyType
	//  期权市值
	OptionValue types.TThostFtdcMoneyType
}

// 投资者持仓
type CThostFtdcInvestorPositionField struct {
	//  保留的无效字段
	Reserve1 types.TThostFtdcOldInstrumentIDType
	//  经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	//  投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	//  持仓多空方向
	PosiDirection types.TThostFtdcPosiDirectionType
	//  投机套保标志
	HedgeFlag types.TThostFtdcHedgeFlagType
	//  持仓日期
	PositionDate types.TThostFtdcPositionDateType
	//  上日持仓
	YdPosition types.TThostFtdcVolumeType
	//  今日持仓
	Position types.TThostFtdcVolumeType
	//  多头冻结
	LongFrozen types.TThostFtdcVolumeType
	//  空头冻结
	ShortFrozen types.TThostFtdcVolumeType
	//  开仓冻结金额
	LongFrozenAmount types.TThostFtdcMoneyType
	//  开仓冻结金额
	ShortFrozenAmount types.TThostFtdcMoneyType
	//  开仓量
	OpenVolume types.TThostFtdcVolumeType
	//  平仓量
	CloseVolume types.TThostFtdcVolumeType
	//  开仓金额
	OpenAmount types.TThostFtdcMoneyType
	//  平仓金额
	CloseAmount types.TThostFtdcMoneyType
	//  持仓成本
	PositionCost types.TThostFtdcMoneyType
	//  上次占用的保证金
	PreMargin types.TThostFtdcMoneyType
	//  占用的保证金
	UseMargin types.TThostFtdcMoneyType
	//  冻结的保证金
	FrozenMargin types.TThostFtdcMoneyType
	//  冻结的资金
	FrozenCash types.TThostFtdcMoneyType
	//  冻结的手续费
	FrozenCommission types.TThostFtdcMoneyType
	//  资金差额
	CashIn types.TThostFtdcMoneyType
	//  手续费
	Commission types.TThostFtdcMoneyType
	//  平仓盈亏
	CloseProfit types.TThostFtdcMoneyType
	//  持仓盈亏
	PositionProfit types.TThostFtdcMoneyType
	//  上次结算价
	PreSettlementPrice types.TThostFtdcPriceType
	//  本次结算价
	SettlementPrice types.TThostFtdcPriceType
	//  交易日
	TradingDay types.TThostFtdcDateType
	//  结算编号
	SettlementID types.TThostFtdcSettlementIDType
	//  开仓成本
	OpenCost types.TThostFtdcMoneyType
	//  交易所保证金
	ExchangeMargin types.TThostFtdcMoneyType
	//  组合成交形成的持仓
	CombPosition types.TThostFtdcVolumeType
	//  组合多头冻结
	CombLongFrozen types.TThostFtdcVolumeType
	//  组合空头冻结
	CombShortFrozen types.TThostFtdcVolumeType
	//  逐日盯市平仓盈亏
	CloseProfitByDate types.TThostFtdcMoneyType
	//  逐笔对冲平仓盈亏
	CloseProfitByTrade types.TThostFtdcMoneyType
	//  今日持仓
	TodayPosition types.TThostFtdcVolumeType
	//  保证金率
	MarginRateByMoney types.TThostFtdcRatioType
	//  保证金率(按手数)
	MarginRateByVolume types.TThostFtdcRatioType
	//  执行冻结
	StrikeFrozen types.TThostFtdcVolumeType
	//  执行冻结金额
	StrikeFrozenAmount types.TThostFtdcMoneyType
	//  放弃执行冻结
	AbandonFrozen types.TThostFtdcVolumeType
	//  交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	//  执行冻结的昨仓
	YdStrikeFrozen types.TThostFtdcVolumeType
	//  投资单元代码
	InvestUnitID types.TThostFtdcInvestUnitIDType
	//  持仓成本差值
	PositionCostOffset types.TThostFtdcMoneyType
	//  tas持仓手数
	TasPosition types.TThostFtdcVolumeType
	//  tas持仓成本
	TasPositionCost types.TThostFtdcMoneyType
	//  合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
	//  期权市值
	OptionValue types.TThostFtdcMoneyType
}

// 合约保证金率
type CThostFtdcInstrumentMarginRateField struct {
	//  保留的无效字段
	Reserve1 types.TThostFtdcOldInstrumentIDType
	//  投资者范围
	InvestorRange types.TThostFtdcInvestorRangeType
	//  经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	//  投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	//  投机套保标志
	HedgeFlag types.TThostFtdcHedgeFlagType
	//  多头保证金率
	LongMarginRatioByMoney types.TThostFtdcRatioType
	//  多头保证金费
	LongMarginRatioByVolume types.TThostFtdcMoneyType
	//  空头保证金率
	ShortMarginRatioByMoney types.TThostFtdcRatioType
	//  空头保证金费
	ShortMarginRatioByVolume types.TThostFtdcMoneyType
	//  是否相对交易所收取
	IsRelative types.TThostFtdcBoolType
	//  交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	//  投资单元代码
	InvestUnitID types.TThostFtdcInvestUnitIDType
	//  合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
}

// 合约手续费率
type CThostFtdcInstrumentCommissionRateField struct {
	//  保留的无效字段
	Reserve1 types.TThostFtdcOldInstrumentIDType
	//  投资者范围
	InvestorRange types.TThostFtdcInvestorRangeType
	//  经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	//  投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	//  开仓手续费率
	OpenRatioByMoney types.TThostFtdcRatioType
	//  开仓手续费
	OpenRatioByVolume types.TThostFtdcRatioType
	//  平仓手续费率
	CloseRatioByMoney types.TThostFtdcRatioType
	//  平仓手续费
	CloseRatioByVolume types.TThostFtdcRatioType
	//  平今手续费率
	CloseTodayRatioByMoney types.TThostFtdcRatioType
	//  平今手续费
	CloseTodayRatioByVolume types.TThostFtdcRatioType
	//  交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	//  业务类型
	BizType types.TThostFtdcBizTypeType
	//  投资单元代码
	InvestUnitID types.TThostFtdcInvestUnitIDType
	//  合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
}

// 深度行情
type CThostFtdcDepthMarketDataField struct {
	//  交易日
	TradingDay types.TThostFtdcDateType
	//  保留的无效字段
	Reserve1 types.TThostFtdcOldInstrumentIDType
	//  交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	//  保留的无效字段
	Reserve2 types.TThostFtdcOldExchangeInstIDType
	//  最新价
	LastPrice types.TThostFtdcPriceType
	//  上次结算价
	PreSettlementPrice types.TThostFtdcPriceType
	//  昨收盘
	PreClosePrice types.TThostFtdcPriceType
	//  昨持仓量
	PreOpenInterest types.TThostFtdcLargeVolumeType
	//  今开盘
	OpenPrice types.TThostFtdcPriceType
	//  最高价
	HighestPrice types.TThostFtdcPriceType
	//  最低价
	LowestPrice types.TThostFtdcPriceType
	//  数量
	Volume types.TThostFtdcVolumeType
	//  成交金额
	Turnover types.TThostFtdcMoneyType
	//  持仓量
	OpenInterest types.TThostFtdcLargeVolumeType
	//  今收盘
	ClosePrice types.TThostFtdcPriceType
	//  本次结算价
	SettlementPrice types.TThostFtdcPriceType
	//  涨停板价
	UpperLimitPrice types.TThostFtdcPriceType
	//  跌停板价
	LowerLimitPrice types.TThostFtdcPriceType
	//  昨虚实度
	PreDelta types.TThostFtdcRatioType
	//  今虚实度
	CurrDelta types.TThostFtdcRatioType
	//  最后修改时间
	UpdateTime types.TThostFtdcTimeType
	//  最后修改毫秒
	UpdateMillisec types.TThostFtdcMillisecType
	//  申买价一
	BidPrice1 types.TThostFtdcPriceType
	//  申买量一
	BidVolume1 types.TThostFtdcVolumeType
	//  申卖价一
	AskPrice1 types.TThostFtdcPriceType
	//  申卖量一
	AskVolume1 types.TThostFtdcVolumeType
	//  申买价二
	BidPrice2 types.TThostFtdcPriceType
	//  申买量二
	BidVolume2 types.TThostFtdcVolumeType
	//  申卖价二
	AskPrice2 types.TThostFtdcPriceType
	//  申卖量二
	AskVolume2 types.TThostFtdcVolumeType
	//  申买价三
	BidPrice3 types.TThostFtdcPriceType
	//  申买量三
	BidVolume3 types.TThostFtdcVolumeType
	//  申卖价三
	AskPrice3 types.TThostFtdcPriceType
	//  申卖量三
	AskVolume3 types.TThostFtdcVolumeType
	//  申买价四
	BidPrice4 types.TThostFtdcPriceType
	//  申买量四
	BidVolume4 types.TThostFtdcVolumeType
	//  申卖价四
	AskPrice4 types.TThostFtdcPriceType
	//  申卖量四
	AskVolume4 types.TThostFtdcVolumeType
	//  申买价五
	BidPrice5 types.TThostFtdcPriceType
	//  申买量五
	BidVolume5 types.TThostFtdcVolumeType
	//  申卖价五
	AskPrice5 types.TThostFtdcPriceType
	//  申卖量五
	AskVolume5 types.TThostFtdcVolumeType
	//  当日均价
	AveragePrice types.TThostFtdcPriceType
	//  业务日期
	ActionDay types.TThostFtdcDateType
	//  合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
	//  合约在交易所的代码
	ExchangeInstID types.TThostFtdcExchangeInstIDType
	//  上带价
	BandingUpperPrice types.TThostFtdcPriceType
	//  下带价
	BandingLowerPrice types.TThostFtdcPriceType
}

// 投资者合约交易权限
type CThostFtdcInstrumentTradingRightField struct {
	//  保留的无效字段
	Reserve1 types.TThostFtdcOldInstrumentIDType
	//  投资者范围
	InvestorRange types.TThostFtdcInvestorRangeType
	//  经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	//  投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	//  交易权限
	TradingRight types.TThostFtdcTradingRightType
	//  合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
}

// 经纪公司用户
type CThostFtdcBrokerUserField struct {
	//  经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	//  用户代码
	UserID types.TThostFtdcUserIDType
	//  用户名称
	UserName types.TThostFtdcUserNameType
	//  用户类型
	UserType types.TThostFtdcUserTypeType
	//  是否活跃
	IsActive types.TThostFtdcBoolType
	//  是否使用令牌
	IsUsingOTP types.TThostFtdcBoolType
	//  是否强制终端认证
	IsAuthForce types.TThostFtdcBoolType
}

// 经纪公司用户口令
type CThostFtdcBrokerUserPasswordField struct {
	//  经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	//  用户代码
	UserID types.TThostFtdcUserIDType
	//  密码
	Password types.TThostFtdcPasswordType
	//  上次修改时间
	LastUpdateTime types.TThostFtdcDateTimeType
	//  上次登陆时间
	LastLoginTime types.TThostFtdcDateTimeType
	//  密码过期时间
	ExpireDate types.TThostFtdcDateType
	//  弱密码过期时间
	WeakExpireDate types.TThostFtdcDateType
}

// 经纪公司用户功能权限
type CThostFtdcBrokerUserFunctionField struct {
	//  经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	//  用户代码
	UserID types.TThostFtdcUserIDType
	//  经纪公司功能代码
	BrokerFunctionCode types.TThostFtdcBrokerFunctionCodeType
}

// 交易所交易员报盘机
type CThostFtdcTraderOfferField struct {
	//  交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	//  交易所交易员代码
	TraderID types.TThostFtdcTraderIDType
	//  会员代码
	ParticipantID types.TThostFtdcParticipantIDType
	//  密码
	Password types.TThostFtdcPasswordType
	//  安装编号
	InstallID types.TThostFtdcInstallIDType
	//  本地报单编号
	OrderLocalID types.TThostFtdcOrderLocalIDType
	//  交易所交易员连接状态
	TraderConnectStatus types.TThostFtdcTraderConnectStatusType
	//  发出连接请求的日期
	ConnectRequestDate types.TThostFtdcDateType
	//  发出连接请求的时间
	ConnectRequestTime types.TThostFtdcTimeType
	//  上次报告日期
	LastReportDate types.TThostFtdcDateType
	//  上次报告时间
	LastReportTime types.TThostFtdcTimeType
	//  完成连接日期
	ConnectDate types.TThostFtdcDateType
	//  完成连接时间
	ConnectTime types.TThostFtdcTimeType
	//  启动日期
	StartDate types.TThostFtdcDateType
	//  启动时间
	StartTime types.TThostFtdcTimeType
	//  交易日
	TradingDay types.TThostFtdcDateType
	//  经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	//  本席位最大成交编号
	MaxTradeID types.TThostFtdcTradeIDType
	//  本席位最大报单备拷
	MaxOrderMessageReference types.TThostFtdcReturnCodeType
	//  撤单时选择席位算法
	OrderCancelAlg types.TThostFtdcOrderCancelAlgType
}

// 投资者结算结果
type CThostFtdcSettlementInfoField struct {
	//  交易日
	TradingDay types.TThostFtdcDateType
	//  结算编号
	SettlementID types.TThostFtdcSettlementIDType
	//  经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	//  投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	//  序号
	SequenceNo types.TThostFtdcSequenceNoType
	//  消息正文
	Content types.TThostFtdcContentType
	//  投资者帐号
	AccountID types.TThostFtdcAccountIDType
	//  币种代码
	CurrencyID types.TThostFtdcCurrencyIDType
}

// 合约保证金率调整
type CThostFtdcInstrumentMarginRateAdjustField struct {
	//  保留的无效字段
	Reserve1 types.TThostFtdcOldInstrumentIDType
	//  投资者范围
	InvestorRange types.TThostFtdcInvestorRangeType
	//  经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	//  投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	//  投机套保标志
	HedgeFlag types.TThostFtdcHedgeFlagType
	//  多头保证金率
	LongMarginRatioByMoney types.TThostFtdcRatioType
	//  多头保证金费
	LongMarginRatioByVolume types.TThostFtdcMoneyType
	//  空头保证金率
	ShortMarginRatioByMoney types.TThostFtdcRatioType
	//  空头保证金费
	ShortMarginRatioByVolume types.TThostFtdcMoneyType
	//  是否相对交易所收取
	IsRelative types.TThostFtdcBoolType
	//  合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
}

// 交易所保证金率
type CThostFtdcExchangeMarginRateField struct {
	//  经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	//  保留的无效字段
	Reserve1 types.TThostFtdcOldInstrumentIDType
	//  投机套保标志
	HedgeFlag types.TThostFtdcHedgeFlagType
	//  多头保证金率
	LongMarginRatioByMoney types.TThostFtdcRatioType
	//  多头保证金费
	LongMarginRatioByVolume types.TThostFtdcMoneyType
	//  空头保证金率
	ShortMarginRatioByMoney types.TThostFtdcRatioType
	//  空头保证金费
	ShortMarginRatioByVolume types.TThostFtdcMoneyType
	//  交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	//  合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
}

// 交易所保证金率调整
type CThostFtdcExchangeMarginRateAdjustField struct {
	//  经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	//  保留的无效字段
	Reserve1 types.TThostFtdcOldInstrumentIDType
	//  投机套保标志
	HedgeFlag types.TThostFtdcHedgeFlagType
	//  跟随交易所投资者多头保证金率
	LongMarginRatioByMoney types.TThostFtdcRatioType
	//  跟随交易所投资者多头保证金费
	LongMarginRatioByVolume types.TThostFtdcMoneyType
	//  跟随交易所投资者空头保证金率
	ShortMarginRatioByMoney types.TThostFtdcRatioType
	//  跟随交易所投资者空头保证金费
	ShortMarginRatioByVolume types.TThostFtdcMoneyType
	//  交易所多头保证金率
	ExchLongMarginRatioByMoney types.TThostFtdcRatioType
	//  交易所多头保证金费
	ExchLongMarginRatioByVolume types.TThostFtdcMoneyType
	//  交易所空头保证金率
	ExchShortMarginRatioByMoney types.TThostFtdcRatioType
	//  交易所空头保证金费
	ExchShortMarginRatioByVolume types.TThostFtdcMoneyType
	//  不跟随交易所投资者多头保证金率
	NoLongMarginRatioByMoney types.TThostFtdcRatioType
	//  不跟随交易所投资者多头保证金费
	NoLongMarginRatioByVolume types.TThostFtdcMoneyType
	//  不跟随交易所投资者空头保证金率
	NoShortMarginRatioByMoney types.TThostFtdcRatioType
	//  不跟随交易所投资者空头保证金费
	NoShortMarginRatioByVolume types.TThostFtdcMoneyType
	//  合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
}

// 汇率
type CThostFtdcExchangeRateField struct {
	//  经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	//  源币种
	FromCurrencyID types.TThostFtdcCurrencyIDType
	//  源币种单位数量
	FromCurrencyUnit types.TThostFtdcCurrencyUnitType
	//  目标币种
	ToCurrencyID types.TThostFtdcCurrencyIDType
	//  汇率
	ExchangeRate types.TThostFtdcExchangeRateType
}

// 结算引用
type CThostFtdcSettlementRefField struct {
	//  交易日
	TradingDay types.TThostFtdcDateType
	//  结算编号
	SettlementID types.TThostFtdcSettlementIDType
}

// 当前时间
type CThostFtdcCurrentTimeField struct {
	//  当前交易日
	CurrDate types.TThostFtdcDateType
	//  当前时间
	CurrTime types.TThostFtdcTimeType
	//  当前时间（毫秒）
	CurrMillisec types.TThostFtdcMillisecType
	//  自然日期
	ActionDay types.TThostFtdcDateType
}

// 通讯阶段
type CThostFtdcCommPhaseField struct {
	//  交易日
	TradingDay types.TThostFtdcDateType
	//  通讯时段编号
	CommPhaseNo types.TThostFtdcCommPhaseNoType
	//  系统编号
	SystemID types.TThostFtdcSystemIDType
}

// 登录信息
type CThostFtdcLoginInfoField struct {
	//  前置编号
	FrontID types.TThostFtdcFrontIDType
	//  会话编号
	SessionID types.TThostFtdcSessionIDType
	//  经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	//  用户代码
	UserID types.TThostFtdcUserIDType
	//  登录日期
	LoginDate types.TThostFtdcDateType
	//  登录时间
	LoginTime types.TThostFtdcTimeType
	//  保留的无效字段
	Reserve1 types.TThostFtdcOldIPAddressType
	//  用户端产品信息
	UserProductInfo types.TThostFtdcProductInfoType
	//  接口端产品信息
	InterfaceProductInfo types.TThostFtdcProductInfoType
	//  协议信息
	ProtocolInfo types.TThostFtdcProtocolInfoType
	//  系统名称
	SystemName types.TThostFtdcSystemNameType
	//  密码,已弃用
	PasswordDeprecated types.TThostFtdcPasswordType
	//  最大报单引用
	MaxOrderRef types.TThostFtdcOrderRefType
	//  上期所时间
	SHFETime types.TThostFtdcTimeType
	//  大商所时间
	DCETime types.TThostFtdcTimeType
	//  郑商所时间
	CZCETime types.TThostFtdcTimeType
	//  中金所时间
	FFEXTime types.TThostFtdcTimeType
	//  Mac地址
	MacAddress types.TThostFtdcMacAddressType
	//  动态密码
	OneTimePassword types.TThostFtdcPasswordType
	//  能源中心时间
	INETime types.TThostFtdcTimeType
	//  查询时是否需要流控
	IsQryControl types.TThostFtdcBoolType
	//  登录备注
	LoginRemark types.TThostFtdcLoginRemarkType
	//  密码
	Password types.TThostFtdcPasswordType
	//  IP地址
	IPAddress types.TThostFtdcIPAddressType
}

// 登录信息
type CThostFtdcLogoutAllField struct {
	//  前置编号
	FrontID types.TThostFtdcFrontIDType
	//  会话编号
	SessionID types.TThostFtdcSessionIDType
	//  系统名称
	SystemName types.TThostFtdcSystemNameType
}

// 前置状态
type CThostFtdcFrontStatusField struct {
	//  前置编号
	FrontID types.TThostFtdcFrontIDType
	//  上次报告日期
	LastReportDate types.TThostFtdcDateType
	//  上次报告时间
	LastReportTime types.TThostFtdcTimeType
	//  是否活跃
	IsActive types.TThostFtdcBoolType
}

// 用户口令变更
type CThostFtdcUserPasswordUpdateField struct {
	//  经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	//  用户代码
	UserID types.TThostFtdcUserIDType
	//  原来的口令
	OldPassword types.TThostFtdcPasswordType
	//  新的口令
	NewPassword types.TThostFtdcPasswordType
}

// 输入报单
type CThostFtdcInputOrderField struct {
	//  经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	//  投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	//  保留的无效字段
	Reserve1 types.TThostFtdcOldInstrumentIDType
	//  报单引用
	OrderRef types.TThostFtdcOrderRefType
	//  用户代码
	UserID types.TThostFtdcUserIDType
	//  报单价格条件
	OrderPriceType types.TThostFtdcOrderPriceTypeType
	//  买卖方向
	Direction types.TThostFtdcDirectionType
	//  组合开平标志
	CombOffsetFlag types.TThostFtdcCombOffsetFlagType
	//  组合投机套保标志
	CombHedgeFlag types.TThostFtdcCombHedgeFlagType
	//  价格
	LimitPrice types.TThostFtdcPriceType
	//  数量
	VolumeTotalOriginal types.TThostFtdcVolumeType
	//  有效期类型
	TimeCondition types.TThostFtdcTimeConditionType
	//  GTD日期
	GTDDate types.TThostFtdcDateType
	//  成交量类型
	VolumeCondition types.TThostFtdcVolumeConditionType
	//  最小成交量
	MinVolume types.TThostFtdcVolumeType
	//  触发条件
	ContingentCondition types.TThostFtdcContingentConditionType
	//  止损价
	StopPrice types.TThostFtdcPriceType
	//  强平原因
	ForceCloseReason types.TThostFtdcForceCloseReasonType
	//  自动挂起标志
	IsAutoSuspend types.TThostFtdcBoolType
	//  业务单元
	BusinessUnit types.TThostFtdcBusinessUnitType
	//  请求编号
	RequestID types.TThostFtdcRequestIDType
	//  用户强平标志
	UserForceClose types.TThostFtdcBoolType
	//  互换单标志
	IsSwapOrder types.TThostFtdcBoolType
	//  交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	//  投资单元代码
	InvestUnitID types.TThostFtdcInvestUnitIDType
	//  资金账号
	AccountID types.TThostFtdcAccountIDType
	//  币种代码
	CurrencyID types.TThostFtdcCurrencyIDType
	//  交易编码
	ClientID types.TThostFtdcClientIDType
	//  保留的无效字段
	Reserve2 types.TThostFtdcOldIPAddressType
	//  Mac地址
	MacAddress types.TThostFtdcMacAddressType
	//  合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
	//  IP地址
	IPAddress types.TThostFtdcIPAddressType
	//  报单回显字段
	OrderMemo types.TThostFtdcOrderMemoType
	//  session上请求计数 api自动维护
	SessionReqSeq types.TThostFtdcSequenceNo12Type
}

// 报单
type CThostFtdcOrderField struct {
	//  经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	//  投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	//  保留的无效字段
	Reserve1 types.TThostFtdcOldInstrumentIDType
	//  报单引用
	OrderRef types.TThostFtdcOrderRefType
	//  用户代码
	UserID types.TThostFtdcUserIDType
	//  报单价格条件
	OrderPriceType types.TThostFtdcOrderPriceTypeType
	//  买卖方向
	Direction types.TThostFtdcDirectionType
	//  组合开平标志
	CombOffsetFlag types.TThostFtdcCombOffsetFlagType
	//  组合投机套保标志
	CombHedgeFlag types.TThostFtdcCombHedgeFlagType
	//  价格
	LimitPrice types.TThostFtdcPriceType
	//  数量
	VolumeTotalOriginal types.TThostFtdcVolumeType
	//  有效期类型
	TimeCondition types.TThostFtdcTimeConditionType
	//  GTD日期
	GTDDate types.TThostFtdcDateType
	//  成交量类型
	VolumeCondition types.TThostFtdcVolumeConditionType
	//  最小成交量
	MinVolume types.TThostFtdcVolumeType
	//  触发条件
	ContingentCondition types.TThostFtdcContingentConditionType
	//  止损价
	StopPrice types.TThostFtdcPriceType
	//  强平原因
	ForceCloseReason types.TThostFtdcForceCloseReasonType
	//  自动挂起标志
	IsAutoSuspend types.TThostFtdcBoolType
	//  业务单元
	BusinessUnit types.TThostFtdcBusinessUnitType
	//  请求编号
	RequestID types.TThostFtdcRequestIDType
	//  本地报单编号
	OrderLocalID types.TThostFtdcOrderLocalIDType
	//  交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	//  会员代码
	ParticipantID types.TThostFtdcParticipantIDType
	//  客户代码
	ClientID types.TThostFtdcClientIDType
	//  保留的无效字段
	Reserve2 types.TThostFtdcOldExchangeInstIDType
	//  交易所交易员代码
	TraderID types.TThostFtdcTraderIDType
	//  安装编号
	InstallID types.TThostFtdcInstallIDType
	//  报单提交状态
	OrderSubmitStatus types.TThostFtdcOrderSubmitStatusType
	//  报单提示序号
	NotifySequence types.TThostFtdcSequenceNoType
	//  交易日
	TradingDay types.TThostFtdcDateType
	//  结算编号
	SettlementID types.TThostFtdcSettlementIDType
	//  报单编号
	OrderSysID types.TThostFtdcOrderSysIDType
	//  报单来源
	OrderSource types.TThostFtdcOrderSourceType
	//  报单状态
	OrderStatus types.TThostFtdcOrderStatusType
	//  报单类型
	OrderType types.TThostFtdcOrderTypeType
	//  今成交数量
	VolumeTraded types.TThostFtdcVolumeType
	//  剩余数量
	VolumeTotal types.TThostFtdcVolumeType
	//  报单日期
	InsertDate types.TThostFtdcDateType
	//  委托时间
	InsertTime types.TThostFtdcTimeType
	//  激活时间
	ActiveTime types.TThostFtdcTimeType
	//  挂起时间
	SuspendTime types.TThostFtdcTimeType
	//  最后修改时间
	UpdateTime types.TThostFtdcTimeType
	//  撤销时间
	CancelTime types.TThostFtdcTimeType
	//  最后修改交易所交易员代码
	ActiveTraderID types.TThostFtdcTraderIDType
	//  结算会员编号
	ClearingPartID types.TThostFtdcParticipantIDType
	//  序号
	SequenceNo types.TThostFtdcSequenceNoType
	//  前置编号
	FrontID types.TThostFtdcFrontIDType
	//  会话编号
	SessionID types.TThostFtdcSessionIDType
	//  用户端产品信息
	UserProductInfo types.TThostFtdcProductInfoType
	//  状态信息
	StatusMsg types.TThostFtdcErrorMsgType
	//  用户强平标志
	UserForceClose types.TThostFtdcBoolType
	//  操作用户代码
	ActiveUserID types.TThostFtdcUserIDType
	//  经纪公司报单编号
	BrokerOrderSeq types.TThostFtdcSequenceNoType
	//  相关报单
	RelativeOrderSysID types.TThostFtdcOrderSysIDType
	//  郑商所成交数量
	ZCETotalTradedVolume types.TThostFtdcVolumeType
	//  互换单标志
	IsSwapOrder types.TThostFtdcBoolType
	//  营业部编号
	BranchID types.TThostFtdcBranchIDType
	//  投资单元代码
	InvestUnitID types.TThostFtdcInvestUnitIDType
	//  资金账号
	AccountID types.TThostFtdcAccountIDType
	//  币种代码
	CurrencyID types.TThostFtdcCurrencyIDType
	//  保留的无效字段
	Reserve3 types.TThostFtdcOldIPAddressType
	//  Mac地址
	MacAddress types.TThostFtdcMacAddressType
	//  合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
	//  合约在交易所的代码
	ExchangeInstID types.TThostFtdcExchangeInstIDType
	//  IP地址
	IPAddress types.TThostFtdcIPAddressType
	//  报单回显字段
	OrderMemo types.TThostFtdcOrderMemoType
	//  session上请求计数 api自动维护
	SessionReqSeq types.TThostFtdcSequenceNo12Type
}

// 交易所报单
type CThostFtdcExchangeOrderField struct {
	//  报单价格条件
	OrderPriceType types.TThostFtdcOrderPriceTypeType
	//  买卖方向
	Direction types.TThostFtdcDirectionType
	//  组合开平标志
	CombOffsetFlag types.TThostFtdcCombOffsetFlagType
	//  组合投机套保标志
	CombHedgeFlag types.TThostFtdcCombHedgeFlagType
	//  价格
	LimitPrice types.TThostFtdcPriceType
	//  数量
	VolumeTotalOriginal types.TThostFtdcVolumeType
	//  有效期类型
	TimeCondition types.TThostFtdcTimeConditionType
	//  GTD日期
	GTDDate types.TThostFtdcDateType
	//  成交量类型
	VolumeCondition types.TThostFtdcVolumeConditionType
	//  最小成交量
	MinVolume types.TThostFtdcVolumeType
	//  触发条件
	ContingentCondition types.TThostFtdcContingentConditionType
	//  止损价
	StopPrice types.TThostFtdcPriceType
	//  强平原因
	ForceCloseReason types.TThostFtdcForceCloseReasonType
	//  自动挂起标志
	IsAutoSuspend types.TThostFtdcBoolType
	//  业务单元
	BusinessUnit types.TThostFtdcBusinessUnitType
	//  请求编号
	RequestID types.TThostFtdcRequestIDType
	//  本地报单编号
	OrderLocalID types.TThostFtdcOrderLocalIDType
	//  交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	//  会员代码
	ParticipantID types.TThostFtdcParticipantIDType
	//  客户代码
	ClientID types.TThostFtdcClientIDType
	//  保留的无效字段
	Reserve1 types.TThostFtdcOldExchangeInstIDType
	//  交易所交易员代码
	TraderID types.TThostFtdcTraderIDType
	//  安装编号
	InstallID types.TThostFtdcInstallIDType
	//  报单提交状态
	OrderSubmitStatus types.TThostFtdcOrderSubmitStatusType
	//  报单提示序号
	NotifySequence types.TThostFtdcSequenceNoType
	//  交易日
	TradingDay types.TThostFtdcDateType
	//  结算编号
	SettlementID types.TThostFtdcSettlementIDType
	//  报单编号
	OrderSysID types.TThostFtdcOrderSysIDType
	//  报单来源
	OrderSource types.TThostFtdcOrderSourceType
	//  报单状态
	OrderStatus types.TThostFtdcOrderStatusType
	//  报单类型
	OrderType types.TThostFtdcOrderTypeType
	//  今成交数量
	VolumeTraded types.TThostFtdcVolumeType
	//  剩余数量
	VolumeTotal types.TThostFtdcVolumeType
	//  报单日期
	InsertDate types.TThostFtdcDateType
	//  委托时间
	InsertTime types.TThostFtdcTimeType
	//  激活时间
	ActiveTime types.TThostFtdcTimeType
	//  挂起时间
	SuspendTime types.TThostFtdcTimeType
	//  最后修改时间
	UpdateTime types.TThostFtdcTimeType
	//  撤销时间
	CancelTime types.TThostFtdcTimeType
	//  最后修改交易所交易员代码
	ActiveTraderID types.TThostFtdcTraderIDType
	//  结算会员编号
	ClearingPartID types.TThostFtdcParticipantIDType
	//  序号
	SequenceNo types.TThostFtdcSequenceNoType
	//  营业部编号
	BranchID types.TThostFtdcBranchIDType
	//  保留的无效字段
	Reserve2 types.TThostFtdcOldIPAddressType
	//  Mac地址
	MacAddress types.TThostFtdcMacAddressType
	//  合约在交易所的代码
	ExchangeInstID types.TThostFtdcExchangeInstIDType
	//  IP地址
	IPAddress types.TThostFtdcIPAddressType
}

// 交易所报单插入失败
type CThostFtdcExchangeOrderInsertErrorField struct {
	//  交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	//  会员代码
	ParticipantID types.TThostFtdcParticipantIDType
	//  交易所交易员代码
	TraderID types.TThostFtdcTraderIDType
	//  安装编号
	InstallID types.TThostFtdcInstallIDType
	//  本地报单编号
	OrderLocalID types.TThostFtdcOrderLocalIDType
	//  错误代码
	ErrorID types.TThostFtdcErrorIDType
	//  错误信息
	ErrorMsg types.TThostFtdcErrorMsgType
}

// 输入报单操作
type CThostFtdcInputOrderActionField struct {
	//  经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	//  投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	//  报单操作引用
	OrderActionRef types.TThostFtdcOrderActionRefType
	//  报单引用
	OrderRef types.TThostFtdcOrderRefType
	//  请求编号
	RequestID types.TThostFtdcRequestIDType
	//  前置编号
	FrontID types.TThostFtdcFrontIDType
	//  会话编号
	SessionID types.TThostFtdcSessionIDType
	//  交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	//  报单编号
	OrderSysID types.TThostFtdcOrderSysIDType
	//  操作标志
	ActionFlag types.TThostFtdcActionFlagType
	//  价格
	LimitPrice types.TThostFtdcPriceType
	//  数量变化
	VolumeChange types.TThostFtdcVolumeType
	//  用户代码
	UserID types.TThostFtdcUserIDType
	//  保留的无效字段
	Reserve1 types.TThostFtdcOldInstrumentIDType
	//  投资单元代码
	InvestUnitID types.TThostFtdcInvestUnitIDType
	//  保留的无效字段
	Reserve2 types.TThostFtdcOldIPAddressType
	//  Mac地址
	MacAddress types.TThostFtdcMacAddressType
	//  合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
	//  IP地址
	IPAddress types.TThostFtdcIPAddressType
	//  报单回显字段
	OrderMemo types.TThostFtdcOrderMemoType
	//  session上请求计数 api自动维护
	SessionReqSeq types.TThostFtdcSequenceNo12Type
}

// 报单操作
type CThostFtdcOrderActionField struct {
	//  经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	//  投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	//  报单操作引用
	OrderActionRef types.TThostFtdcOrderActionRefType
	//  报单引用
	OrderRef types.TThostFtdcOrderRefType
	//  请求编号
	RequestID types.TThostFtdcRequestIDType
	//  前置编号
	FrontID types.TThostFtdcFrontIDType
	//  会话编号
	SessionID types.TThostFtdcSessionIDType
	//  交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	//  报单编号
	OrderSysID types.TThostFtdcOrderSysIDType
	//  操作标志
	ActionFlag types.TThostFtdcActionFlagType
	//  价格
	LimitPrice types.TThostFtdcPriceType
	//  数量变化
	VolumeChange types.TThostFtdcVolumeType
	//  操作日期
	ActionDate types.TThostFtdcDateType
	//  操作时间
	ActionTime types.TThostFtdcTimeType
	//  交易所交易员代码
	TraderID types.TThostFtdcTraderIDType
	//  安装编号
	InstallID types.TThostFtdcInstallIDType
	//  本地报单编号
	OrderLocalID types.TThostFtdcOrderLocalIDType
	//  操作本地编号
	ActionLocalID types.TThostFtdcOrderLocalIDType
	//  会员代码
	ParticipantID types.TThostFtdcParticipantIDType
	//  客户代码
	ClientID types.TThostFtdcClientIDType
	//  业务单元
	BusinessUnit types.TThostFtdcBusinessUnitType
	//  报单操作状态
	OrderActionStatus types.TThostFtdcOrderActionStatusType
	//  用户代码
	UserID types.TThostFtdcUserIDType
	//  状态信息
	StatusMsg types.TThostFtdcErrorMsgType
	//  保留的无效字段
	Reserve1 types.TThostFtdcOldInstrumentIDType
	//  营业部编号
	BranchID types.TThostFtdcBranchIDType
	//  投资单元代码
	InvestUnitID types.TThostFtdcInvestUnitIDType
	//  保留的无效字段
	Reserve2 types.TThostFtdcOldIPAddressType
	//  Mac地址
	MacAddress types.TThostFtdcMacAddressType
	//  合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
	//  IP地址
	IPAddress types.TThostFtdcIPAddressType
	//  报单回显字段
	OrderMemo types.TThostFtdcOrderMemoType
	//  session上请求计数 api自动维护
	SessionReqSeq types.TThostFtdcSequenceNo12Type
}

// 交易所报单操作
type CThostFtdcExchangeOrderActionField struct {
	//  交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	//  报单编号
	OrderSysID types.TThostFtdcOrderSysIDType
	//  操作标志
	ActionFlag types.TThostFtdcActionFlagType
	//  价格
	LimitPrice types.TThostFtdcPriceType
	//  数量变化
	VolumeChange types.TThostFtdcVolumeType
	//  操作日期
	ActionDate types.TThostFtdcDateType
	//  操作时间
	ActionTime types.TThostFtdcTimeType
	//  交易所交易员代码
	TraderID types.TThostFtdcTraderIDType
	//  安装编号
	InstallID types.TThostFtdcInstallIDType
	//  本地报单编号
	OrderLocalID types.TThostFtdcOrderLocalIDType
	//  操作本地编号
	ActionLocalID types.TThostFtdcOrderLocalIDType
	//  会员代码
	ParticipantID types.TThostFtdcParticipantIDType
	//  客户代码
	ClientID types.TThostFtdcClientIDType
	//  业务单元
	BusinessUnit types.TThostFtdcBusinessUnitType
	//  报单操作状态
	OrderActionStatus types.TThostFtdcOrderActionStatusType
	//  用户代码
	UserID types.TThostFtdcUserIDType
	//  营业部编号
	BranchID types.TThostFtdcBranchIDType
	//  保留的无效字段
	Reserve1 types.TThostFtdcOldIPAddressType
	//  Mac地址
	MacAddress types.TThostFtdcMacAddressType
	//  IP地址
	IPAddress types.TThostFtdcIPAddressType
}

// 交易所报单操作失败
type CThostFtdcExchangeOrderActionErrorField struct {
	//  交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	//  报单编号
	OrderSysID types.TThostFtdcOrderSysIDType
	//  交易所交易员代码
	TraderID types.TThostFtdcTraderIDType
	//  安装编号
	InstallID types.TThostFtdcInstallIDType
	//  本地报单编号
	OrderLocalID types.TThostFtdcOrderLocalIDType
	//  操作本地编号
	ActionLocalID types.TThostFtdcOrderLocalIDType
	//  错误代码
	ErrorID types.TThostFtdcErrorIDType
	//  错误信息
	ErrorMsg types.TThostFtdcErrorMsgType
}

// 交易所成交
type CThostFtdcExchangeTradeField struct {
	//  交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	//  成交编号
	TradeID types.TThostFtdcTradeIDType
	//  买卖方向
	Direction types.TThostFtdcDirectionType
	//  报单编号
	OrderSysID types.TThostFtdcOrderSysIDType
	//  会员代码
	ParticipantID types.TThostFtdcParticipantIDType
	//  客户代码
	ClientID types.TThostFtdcClientIDType
	//  交易角色
	TradingRole types.TThostFtdcTradingRoleType
	//  保留的无效字段
	Reserve1 types.TThostFtdcOldExchangeInstIDType
	//  开平标志
	OffsetFlag types.TThostFtdcOffsetFlagType
	//  投机套保标志
	HedgeFlag types.TThostFtdcHedgeFlagType
	//  价格
	Price types.TThostFtdcPriceType
	//  数量
	Volume types.TThostFtdcVolumeType
	//  成交时期
	TradeDate types.TThostFtdcDateType
	//  成交时间
	TradeTime types.TThostFtdcTimeType
	//  成交类型
	TradeType types.TThostFtdcTradeTypeType
	//  成交价来源
	PriceSource types.TThostFtdcPriceSourceType
	//  交易所交易员代码
	TraderID types.TThostFtdcTraderIDType
	//  本地报单编号
	OrderLocalID types.TThostFtdcOrderLocalIDType
	//  结算会员编号
	ClearingPartID types.TThostFtdcParticipantIDType
	//  业务单元
	BusinessUnit types.TThostFtdcBusinessUnitType
	//  序号
	SequenceNo types.TThostFtdcSequenceNoType
	//  成交来源
	TradeSource types.TThostFtdcTradeSourceType
	//  合约在交易所的代码
	ExchangeInstID types.TThostFtdcExchangeInstIDType
}

// 成交
type CThostFtdcTradeField struct {
	//  经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	//  投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	//  保留的无效字段
	Reserve1 types.TThostFtdcOldInstrumentIDType
	//  报单引用
	OrderRef types.TThostFtdcOrderRefType
	//  用户代码
	UserID types.TThostFtdcUserIDType
	//  交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	//  成交编号
	TradeID types.TThostFtdcTradeIDType
	//  买卖方向
	Direction types.TThostFtdcDirectionType
	//  报单编号
	OrderSysID types.TThostFtdcOrderSysIDType
	//  会员代码
	ParticipantID types.TThostFtdcParticipantIDType
	//  客户代码
	ClientID types.TThostFtdcClientIDType
	//  交易角色
	TradingRole types.TThostFtdcTradingRoleType
	//  保留的无效字段
	Reserve2 types.TThostFtdcOldExchangeInstIDType
	//  开平标志
	OffsetFlag types.TThostFtdcOffsetFlagType
	//  投机套保标志
	HedgeFlag types.TThostFtdcHedgeFlagType
	//  价格
	Price types.TThostFtdcPriceType
	//  数量
	Volume types.TThostFtdcVolumeType
	//  成交时期
	TradeDate types.TThostFtdcDateType
	//  成交时间
	TradeTime types.TThostFtdcTimeType
	//  成交类型
	TradeType types.TThostFtdcTradeTypeType
	//  成交价来源
	PriceSource types.TThostFtdcPriceSourceType
	//  交易所交易员代码
	TraderID types.TThostFtdcTraderIDType
	//  本地报单编号
	OrderLocalID types.TThostFtdcOrderLocalIDType
	//  结算会员编号
	ClearingPartID types.TThostFtdcParticipantIDType
	//  业务单元
	BusinessUnit types.TThostFtdcBusinessUnitType
	//  序号
	SequenceNo types.TThostFtdcSequenceNoType
	//  交易日
	TradingDay types.TThostFtdcDateType
	//  结算编号
	SettlementID types.TThostFtdcSettlementIDType
	//  经纪公司报单编号
	BrokerOrderSeq types.TThostFtdcSequenceNoType
	//  成交来源
	TradeSource types.TThostFtdcTradeSourceType
	//  投资单元代码
	InvestUnitID types.TThostFtdcInvestUnitIDType
	//  合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
	//  合约在交易所的代码
	ExchangeInstID types.TThostFtdcExchangeInstIDType
}

// 用户会话
type CThostFtdcUserSessionField struct {
	//  前置编号
	FrontID types.TThostFtdcFrontIDType
	//  会话编号
	SessionID types.TThostFtdcSessionIDType
	//  经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	//  用户代码
	UserID types.TThostFtdcUserIDType
	//  登录日期
	LoginDate types.TThostFtdcDateType
	//  登录时间
	LoginTime types.TThostFtdcTimeType
	//  保留的无效字段
	Reserve1 types.TThostFtdcOldIPAddressType
	//  用户端产品信息
	UserProductInfo types.TThostFtdcProductInfoType
	//  接口端产品信息
	InterfaceProductInfo types.TThostFtdcProductInfoType
	//  协议信息
	ProtocolInfo types.TThostFtdcProtocolInfoType
	//  Mac地址
	MacAddress types.TThostFtdcMacAddressType
	//  登录备注
	LoginRemark types.TThostFtdcLoginRemarkType
	//  IP地址
	IPAddress types.TThostFtdcIPAddressType
}

// 查询最大报单数量
type CThostFtdcQryMaxOrderVolumeField struct {
	//  经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	//  投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	//  保留的无效字段
	Reserve1 types.TThostFtdcOldInstrumentIDType
	//  买卖方向
	Direction types.TThostFtdcDirectionType
	//  开平标志
	OffsetFlag types.TThostFtdcOffsetFlagType
	//  投机套保标志
	HedgeFlag types.TThostFtdcHedgeFlagType
	//  最大允许报单数量
	MaxVolume types.TThostFtdcVolumeType
	//  交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	//  投资单元代码
	InvestUnitID types.TThostFtdcInvestUnitIDType
	//  合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
}

// 投资者结算结果确认信息
type CThostFtdcSettlementInfoConfirmField struct {
	//  经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	//  投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	//  确认日期
	ConfirmDate types.TThostFtdcDateType
	//  确认时间
	ConfirmTime types.TThostFtdcTimeType
	//  结算编号
	SettlementID types.TThostFtdcSettlementIDType
	//  投资者帐号
	AccountID types.TThostFtdcAccountIDType
	//  币种代码
	CurrencyID types.TThostFtdcCurrencyIDType
}

// 出入金同步
type CThostFtdcSyncDepositField struct {
	//  出入金流水号
	DepositSeqNo types.TThostFtdcDepositSeqNoType
	//  经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	//  投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	//  入金金额
	Deposit types.TThostFtdcMoneyType
	//  是否强制进行
	IsForce types.TThostFtdcBoolType
	//  币种代码
	CurrencyID types.TThostFtdcCurrencyIDType
	//  是否是个股期权内转
	IsFromSopt types.TThostFtdcBoolType
	//  资金密码
	TradingPassword types.TThostFtdcPasswordType
	//  是否二级代理商的内转
	IsSecAgentTranfer types.TThostFtdcBoolType
}

// 货币质押同步
type CThostFtdcSyncFundMortgageField struct {
	//  货币质押流水号
	MortgageSeqNo types.TThostFtdcDepositSeqNoType
	//  经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	//  投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	//  源币种
	FromCurrencyID types.TThostFtdcCurrencyIDType
	//  质押金额
	MortgageAmount types.TThostFtdcMoneyType
	//  目标币种
	ToCurrencyID types.TThostFtdcCurrencyIDType
}

// 经纪公司同步
type CThostFtdcBrokerSyncField struct {
	//  经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
}

// 正在同步中的投资者
type CThostFtdcSyncingInvestorField struct {
	//  投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	//  经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	//  投资者分组代码
	InvestorGroupID types.TThostFtdcInvestorIDType
	//  投资者名称
	InvestorName types.TThostFtdcPartyNameType
	//  证件类型
	IdentifiedCardType types.TThostFtdcIdCardTypeType
	//  证件号码
	IdentifiedCardNo types.TThostFtdcIdentifiedCardNoType
	//  是否活跃
	IsActive types.TThostFtdcBoolType
	//  联系电话
	Telephone types.TThostFtdcTelephoneType
	//  通讯地址
	Address types.TThostFtdcAddressType
	//  开户日期
	OpenDate types.TThostFtdcDateType
	//  手机
	Mobile types.TThostFtdcMobileType
	//  手续费率模板代码
	CommModelID types.TThostFtdcInvestorIDType
	//  保证金率模板代码
	MarginModelID types.TThostFtdcInvestorIDType
	//  是否频率控制
	IsOrderFreq types.TThostFtdcEnumBoolType
	//  是否开仓限制
	IsOpenVolLimit types.TThostFtdcEnumBoolType
}

// 正在同步中的交易代码
type CThostFtdcSyncingTradingCodeField struct {
	//  投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	//  经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	//  交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	//  客户代码
	ClientID types.TThostFtdcClientIDType
	//  是否活跃
	IsActive types.TThostFtdcBoolType
	//  交易编码类型
	ClientIDType types.TThostFtdcClientIDTypeType
}

// 正在同步中的投资者分组
type CThostFtdcSyncingInvestorGroupField struct {
	//  经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	//  投资者分组代码
	InvestorGroupID types.TThostFtdcInvestorIDType
	//  投资者分组名称
	InvestorGroupName types.TThostFtdcInvestorGroupNameType
}

// 正在同步中的交易账号
type CThostFtdcSyncingTradingAccountField struct {
	//  经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	//  投资者帐号
	AccountID types.TThostFtdcAccountIDType
	//  上次质押金额
	PreMortgage types.TThostFtdcMoneyType
	//  上次信用额度
	PreCredit types.TThostFtdcMoneyType
	//  上次存款额
	PreDeposit types.TThostFtdcMoneyType
	//  上次结算准备金
	PreBalance types.TThostFtdcMoneyType
	//  上次占用的保证金
	PreMargin types.TThostFtdcMoneyType
	//  利息基数
	InterestBase types.TThostFtdcMoneyType
	//  利息收入
	Interest types.TThostFtdcMoneyType
	//  入金金额
	Deposit types.TThostFtdcMoneyType
	//  出金金额
	Withdraw types.TThostFtdcMoneyType
	//  冻结的保证金
	FrozenMargin types.TThostFtdcMoneyType
	//  冻结的资金
	FrozenCash types.TThostFtdcMoneyType
	//  冻结的手续费
	FrozenCommission types.TThostFtdcMoneyType
	//  当前保证金总额
	CurrMargin types.TThostFtdcMoneyType
	//  资金差额
	CashIn types.TThostFtdcMoneyType
	//  手续费
	Commission types.TThostFtdcMoneyType
	//  平仓盈亏
	CloseProfit types.TThostFtdcMoneyType
	//  持仓盈亏
	PositionProfit types.TThostFtdcMoneyType
	//  期货结算准备金
	Balance types.TThostFtdcMoneyType
	//  可用资金
	Available types.TThostFtdcMoneyType
	//  可取资金
	WithdrawQuota types.TThostFtdcMoneyType
	//  基本准备金
	Reserve types.TThostFtdcMoneyType
	//  交易日
	TradingDay types.TThostFtdcDateType
	//  结算编号
	SettlementID types.TThostFtdcSettlementIDType
	//  信用额度
	Credit types.TThostFtdcMoneyType
	//  质押金额
	Mortgage types.TThostFtdcMoneyType
	//  交易所保证金
	ExchangeMargin types.TThostFtdcMoneyType
	//  投资者交割保证金
	DeliveryMargin types.TThostFtdcMoneyType
	//  交易所交割保证金
	ExchangeDeliveryMargin types.TThostFtdcMoneyType
	//  保底期货结算准备金
	ReserveBalance types.TThostFtdcMoneyType
	//  币种代码
	CurrencyID types.TThostFtdcCurrencyIDType
	//  上次货币质入金额
	PreFundMortgageIn types.TThostFtdcMoneyType
	//  上次货币质出金额
	PreFundMortgageOut types.TThostFtdcMoneyType
	//  货币质入金额
	FundMortgageIn types.TThostFtdcMoneyType
	//  货币质出金额
	FundMortgageOut types.TThostFtdcMoneyType
	//  货币质押余额
	FundMortgageAvailable types.TThostFtdcMoneyType
	//  可质押货币金额
	MortgageableFund types.TThostFtdcMoneyType
	//  特殊产品占用保证金
	SpecProductMargin types.TThostFtdcMoneyType
	//  特殊产品冻结保证金
	SpecProductFrozenMargin types.TThostFtdcMoneyType
	//  特殊产品手续费
	SpecProductCommission types.TThostFtdcMoneyType
	//  特殊产品冻结手续费
	SpecProductFrozenCommission types.TThostFtdcMoneyType
	//  特殊产品持仓盈亏
	SpecProductPositionProfit types.TThostFtdcMoneyType
	//  特殊产品平仓盈亏
	SpecProductCloseProfit types.TThostFtdcMoneyType
	//  根据持仓盈亏算法计算的特殊产品持仓盈亏
	SpecProductPositionProfitByAlg types.TThostFtdcMoneyType
	//  特殊产品交易所保证金
	SpecProductExchangeMargin types.TThostFtdcMoneyType
	//  延时换汇冻结金额
	FrozenSwap types.TThostFtdcMoneyType
	//  剩余换汇额度
	RemainSwap types.TThostFtdcMoneyType
	//  期权市值
	OptionValue types.TThostFtdcMoneyType
}

// 正在同步中的投资者持仓
type CThostFtdcSyncingInvestorPositionField struct {
	//  保留的无效字段
	Reserve1 types.TThostFtdcOldInstrumentIDType
	//  经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	//  投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	//  持仓多空方向
	PosiDirection types.TThostFtdcPosiDirectionType
	//  投机套保标志
	HedgeFlag types.TThostFtdcHedgeFlagType
	//  持仓日期
	PositionDate types.TThostFtdcPositionDateType
	//  上日持仓
	YdPosition types.TThostFtdcVolumeType
	//  今日持仓
	Position types.TThostFtdcVolumeType
	//  多头冻结
	LongFrozen types.TThostFtdcVolumeType
	//  空头冻结
	ShortFrozen types.TThostFtdcVolumeType
	//  开仓冻结金额
	LongFrozenAmount types.TThostFtdcMoneyType
	//  开仓冻结金额
	ShortFrozenAmount types.TThostFtdcMoneyType
	//  开仓量
	OpenVolume types.TThostFtdcVolumeType
	//  平仓量
	CloseVolume types.TThostFtdcVolumeType
	//  开仓金额
	OpenAmount types.TThostFtdcMoneyType
	//  平仓金额
	CloseAmount types.TThostFtdcMoneyType
	//  持仓成本
	PositionCost types.TThostFtdcMoneyType
	//  上次占用的保证金
	PreMargin types.TThostFtdcMoneyType
	//  占用的保证金
	UseMargin types.TThostFtdcMoneyType
	//  冻结的保证金
	FrozenMargin types.TThostFtdcMoneyType
	//  冻结的资金
	FrozenCash types.TThostFtdcMoneyType
	//  冻结的手续费
	FrozenCommission types.TThostFtdcMoneyType
	//  资金差额
	CashIn types.TThostFtdcMoneyType
	//  手续费
	Commission types.TThostFtdcMoneyType
	//  平仓盈亏
	CloseProfit types.TThostFtdcMoneyType
	//  持仓盈亏
	PositionProfit types.TThostFtdcMoneyType
	//  上次结算价
	PreSettlementPrice types.TThostFtdcPriceType
	//  本次结算价
	SettlementPrice types.TThostFtdcPriceType
	//  交易日
	TradingDay types.TThostFtdcDateType
	//  结算编号
	SettlementID types.TThostFtdcSettlementIDType
	//  开仓成本
	OpenCost types.TThostFtdcMoneyType
	//  交易所保证金
	ExchangeMargin types.TThostFtdcMoneyType
	//  组合成交形成的持仓
	CombPosition types.TThostFtdcVolumeType
	//  组合多头冻结
	CombLongFrozen types.TThostFtdcVolumeType
	//  组合空头冻结
	CombShortFrozen types.TThostFtdcVolumeType
	//  逐日盯市平仓盈亏
	CloseProfitByDate types.TThostFtdcMoneyType
	//  逐笔对冲平仓盈亏
	CloseProfitByTrade types.TThostFtdcMoneyType
	//  今日持仓
	TodayPosition types.TThostFtdcVolumeType
	//  保证金率
	MarginRateByMoney types.TThostFtdcRatioType
	//  保证金率(按手数)
	MarginRateByVolume types.TThostFtdcRatioType
	//  执行冻结
	StrikeFrozen types.TThostFtdcVolumeType
	//  执行冻结金额
	StrikeFrozenAmount types.TThostFtdcMoneyType
	//  放弃执行冻结
	AbandonFrozen types.TThostFtdcVolumeType
	//  交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	//  执行冻结的昨仓
	YdStrikeFrozen types.TThostFtdcVolumeType
	//  投资单元代码
	InvestUnitID types.TThostFtdcInvestUnitIDType
	//  持仓成本差值
	PositionCostOffset types.TThostFtdcMoneyType
	//  tas持仓手数
	TasPosition types.TThostFtdcVolumeType
	//  tas持仓成本
	TasPositionCost types.TThostFtdcMoneyType
	//  合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
}

// 正在同步中的合约保证金率
type CThostFtdcSyncingInstrumentMarginRateField struct {
	//  保留的无效字段
	Reserve1 types.TThostFtdcOldInstrumentIDType
	//  投资者范围
	InvestorRange types.TThostFtdcInvestorRangeType
	//  经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	//  投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	//  投机套保标志
	HedgeFlag types.TThostFtdcHedgeFlagType
	//  多头保证金率
	LongMarginRatioByMoney types.TThostFtdcRatioType
	//  多头保证金费
	LongMarginRatioByVolume types.TThostFtdcMoneyType
	//  空头保证金率
	ShortMarginRatioByMoney types.TThostFtdcRatioType
	//  空头保证金费
	ShortMarginRatioByVolume types.TThostFtdcMoneyType
	//  是否相对交易所收取
	IsRelative types.TThostFtdcBoolType
	//  合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
}

// 正在同步中的合约手续费率
type CThostFtdcSyncingInstrumentCommissionRateField struct {
	//  保留的无效字段
	Reserve1 types.TThostFtdcOldInstrumentIDType
	//  投资者范围
	InvestorRange types.TThostFtdcInvestorRangeType
	//  经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	//  投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	//  开仓手续费率
	OpenRatioByMoney types.TThostFtdcRatioType
	//  开仓手续费
	OpenRatioByVolume types.TThostFtdcRatioType
	//  平仓手续费率
	CloseRatioByMoney types.TThostFtdcRatioType
	//  平仓手续费
	CloseRatioByVolume types.TThostFtdcRatioType
	//  平今手续费率
	CloseTodayRatioByMoney types.TThostFtdcRatioType
	//  平今手续费
	CloseTodayRatioByVolume types.TThostFtdcRatioType
	//  合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
}

// 正在同步中的合约交易权限
type CThostFtdcSyncingInstrumentTradingRightField struct {
	//  保留的无效字段
	Reserve1 types.TThostFtdcOldInstrumentIDType
	//  投资者范围
	InvestorRange types.TThostFtdcInvestorRangeType
	//  经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	//  投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	//  交易权限
	TradingRight types.TThostFtdcTradingRightType
	//  合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
}

// 查询报单
type CThostFtdcQryOrderField struct {
	//  经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	//  投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	//  保留的无效字段
	Reserve1 types.TThostFtdcOldInstrumentIDType
	//  交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	//  报单编号
	OrderSysID types.TThostFtdcOrderSysIDType
	//  开始时间
	InsertTimeStart types.TThostFtdcTimeType
	//  结束时间
	InsertTimeEnd types.TThostFtdcTimeType
	//  投资单元代码
	InvestUnitID types.TThostFtdcInvestUnitIDType
	//  合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
}

// 查询成交
type CThostFtdcQryTradeField struct {
	//  经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	//  投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	//  保留的无效字段
	Reserve1 types.TThostFtdcOldInstrumentIDType
	//  交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	//  成交编号
	TradeID types.TThostFtdcTradeIDType
	//  开始时间
	TradeTimeStart types.TThostFtdcTimeType
	//  结束时间
	TradeTimeEnd types.TThostFtdcTimeType
	//  投资单元代码
	InvestUnitID types.TThostFtdcInvestUnitIDType
	//  合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
}

// 查询投资者持仓
type CThostFtdcQryInvestorPositionField struct {
	//  经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	//  投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	//  保留的无效字段
	Reserve1 types.TThostFtdcOldInstrumentIDType
	//  交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	//  投资单元代码
	InvestUnitID types.TThostFtdcInvestUnitIDType
	//  合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
}

// 查询资金账户
type CThostFtdcQryTradingAccountField struct {
	//  经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	//  投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	//  币种代码
	CurrencyID types.TThostFtdcCurrencyIDType
	//  业务类型
	BizType types.TThostFtdcBizTypeType
	//  投资者帐号
	AccountID types.TThostFtdcAccountIDType
}

// 查询投资者
type CThostFtdcQryInvestorField struct {
	//  经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	//  投资者代码
	InvestorID types.TThostFtdcInvestorIDType
}

// 查询交易编码
type CThostFtdcQryTradingCodeField struct {
	//  经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	//  投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	//  交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	//  客户代码
	ClientID types.TThostFtdcClientIDType
	//  交易编码类型
	ClientIDType types.TThostFtdcClientIDTypeType
	//  投资单元代码
	InvestUnitID types.TThostFtdcInvestUnitIDType
}

// 查询投资者组
type CThostFtdcQryInvestorGroupField struct {
	//  经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
}

// 查询合约保证金率
type CThostFtdcQryInstrumentMarginRateField struct {
	//  经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	//  投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	//  保留的无效字段
	Reserve1 types.TThostFtdcOldInstrumentIDType
	//  投机套保标志
	HedgeFlag types.TThostFtdcHedgeFlagType
	//  交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	//  投资单元代码
	InvestUnitID types.TThostFtdcInvestUnitIDType
	//  合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
}

// 查询手续费率
type CThostFtdcQryInstrumentCommissionRateField struct {
	//  经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	//  投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	//  保留的无效字段
	Reserve1 types.TThostFtdcOldInstrumentIDType
	//  交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	//  投资单元代码
	InvestUnitID types.TThostFtdcInvestUnitIDType
	//  合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
}

// 查询合约交易权限
type CThostFtdcQryInstrumentTradingRightField struct {
	//  经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	//  投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	//  保留的无效字段
	Reserve1 types.TThostFtdcOldInstrumentIDType
	//  合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
}

// 查询经纪公司
type CThostFtdcQryBrokerField struct {
	//  经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
}

// 查询交易员
type CThostFtdcQryTraderField struct {
	//  交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	//  会员代码
	ParticipantID types.TThostFtdcParticipantIDType
	//  交易所交易员代码
	TraderID types.TThostFtdcTraderIDType
}

// 查询管理用户功能权限
type CThostFtdcQrySuperUserFunctionField struct {
	//  用户代码
	UserID types.TThostFtdcUserIDType
}

// 查询用户会话
type CThostFtdcQryUserSessionField struct {
	//  前置编号
	FrontID types.TThostFtdcFrontIDType
	//  会话编号
	SessionID types.TThostFtdcSessionIDType
	//  经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	//  用户代码
	UserID types.TThostFtdcUserIDType
}

// 查询经纪公司会员代码
type CThostFtdcQryPartBrokerField struct {
	//  交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	//  经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	//  会员代码
	ParticipantID types.TThostFtdcParticipantIDType
}

// 查询前置状态
type CThostFtdcQryFrontStatusField struct {
	//  前置编号
	FrontID types.TThostFtdcFrontIDType
}

// 查询交易所报单
type CThostFtdcQryExchangeOrderField struct {
	//  会员代码
	ParticipantID types.TThostFtdcParticipantIDType
	//  客户代码
	ClientID types.TThostFtdcClientIDType
	//  保留的无效字段
	Reserve1 types.TThostFtdcOldExchangeInstIDType
	//  交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	//  交易所交易员代码
	TraderID types.TThostFtdcTraderIDType
	//  合约在交易所的代码
	ExchangeInstID types.TThostFtdcExchangeInstIDType
}

// 查询报单操作
type CThostFtdcQryOrderActionField struct {
	//  经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	//  投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	//  交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
}

// 查询交易所报单操作
type CThostFtdcQryExchangeOrderActionField struct {
	//  会员代码
	ParticipantID types.TThostFtdcParticipantIDType
	//  客户代码
	ClientID types.TThostFtdcClientIDType
	//  交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	//  交易所交易员代码
	TraderID types.TThostFtdcTraderIDType
}

// 查询管理用户
type CThostFtdcQrySuperUserField struct {
	//  用户代码
	UserID types.TThostFtdcUserIDType
}

// 查询交易所
type CThostFtdcQryExchangeField struct {
	//  交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
}

// 查询产品
type CThostFtdcQryProductField struct {
	//  保留的无效字段
	Reserve1 types.TThostFtdcOldInstrumentIDType
	//  产品类型
	ProductClass types.TThostFtdcProductClassType
	//  交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	//  产品代码
	ProductID types.TThostFtdcInstrumentIDType
}

// 查询合约
type CThostFtdcQryInstrumentField struct {
	//  保留的无效字段
	Reserve1 types.TThostFtdcOldInstrumentIDType
	//  交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	//  保留的无效字段
	Reserve2 types.TThostFtdcOldExchangeInstIDType
	//  保留的无效字段
	Reserve3 types.TThostFtdcOldInstrumentIDType
	//  合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
	//  合约在交易所的代码
	ExchangeInstID types.TThostFtdcExchangeInstIDType
	//  产品代码
	ProductID types.TThostFtdcInstrumentIDType
}

// 查询行情
type CThostFtdcQryDepthMarketDataField struct {
	//  保留的无效字段
	Reserve1 types.TThostFtdcOldInstrumentIDType
	//  交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	//  合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
	//  产品类型
	ProductClass types.TThostFtdcProductClassType
}

// 查询经纪公司用户
type CThostFtdcQryBrokerUserField struct {
	//  经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	//  用户代码
	UserID types.TThostFtdcUserIDType
}

// 查询经纪公司用户权限
type CThostFtdcQryBrokerUserFunctionField struct {
	//  经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	//  用户代码
	UserID types.TThostFtdcUserIDType
}

// 查询交易员报盘机
type CThostFtdcQryTraderOfferField struct {
	//  交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	//  会员代码
	ParticipantID types.TThostFtdcParticipantIDType
	//  交易所交易员代码
	TraderID types.TThostFtdcTraderIDType
}

// 查询出入金流水
type CThostFtdcQrySyncDepositField struct {
	//  经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	//  出入金流水号
	DepositSeqNo types.TThostFtdcDepositSeqNoType
}

// 查询投资者结算结果
type CThostFtdcQrySettlementInfoField struct {
	//  经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	//  投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	//  交易日
	TradingDay types.TThostFtdcDateType
	//  投资者帐号
	AccountID types.TThostFtdcAccountIDType
	//  币种代码
	CurrencyID types.TThostFtdcCurrencyIDType
}

// 查询交易所保证金率
type CThostFtdcQryExchangeMarginRateField struct {
	//  经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	//  保留的无效字段
	Reserve1 types.TThostFtdcOldInstrumentIDType
	//  投机套保标志
	HedgeFlag types.TThostFtdcHedgeFlagType
	//  交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	//  合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
}

// 查询交易所调整保证金率
type CThostFtdcQryExchangeMarginRateAdjustField struct {
	//  经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	//  保留的无效字段
	Reserve1 types.TThostFtdcOldInstrumentIDType
	//  投机套保标志
	HedgeFlag types.TThostFtdcHedgeFlagType
	//  合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
}

// 查询汇率
type CThostFtdcQryExchangeRateField struct {
	//  经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	//  源币种
	FromCurrencyID types.TThostFtdcCurrencyIDType
	//  目标币种
	ToCurrencyID types.TThostFtdcCurrencyIDType
}

// 查询货币质押流水
type CThostFtdcQrySyncFundMortgageField struct {
	//  经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	//  货币质押流水号
	MortgageSeqNo types.TThostFtdcDepositSeqNoType
}

// 查询报单
type CThostFtdcQryHisOrderField struct {
	//  经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	//  投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	//  保留的无效字段
	Reserve1 types.TThostFtdcOldInstrumentIDType
	//  交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	//  报单编号
	OrderSysID types.TThostFtdcOrderSysIDType
	//  开始时间
	InsertTimeStart types.TThostFtdcTimeType
	//  结束时间
	InsertTimeEnd types.TThostFtdcTimeType
	//  交易日
	TradingDay types.TThostFtdcDateType
	//  结算编号
	SettlementID types.TThostFtdcSettlementIDType
	//  合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
}

// 当前期权合约最小保证金
type CThostFtdcOptionInstrMiniMarginField struct {
	//  保留的无效字段
	Reserve1 types.TThostFtdcOldInstrumentIDType
	//  投资者范围
	InvestorRange types.TThostFtdcInvestorRangeType
	//  经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	//  投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	//  单位（手）期权合约最小保证金
	MinMargin types.TThostFtdcMoneyType
	//  取值方式
	ValueMethod types.TThostFtdcValueMethodType
	//  是否跟随交易所收取
	IsRelative types.TThostFtdcBoolType
	//  合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
}

// 当前期权合约保证金调整系数
type CThostFtdcOptionInstrMarginAdjustField struct {
	//  保留的无效字段
	Reserve1 types.TThostFtdcOldInstrumentIDType
	//  投资者范围
	InvestorRange types.TThostFtdcInvestorRangeType
	//  经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	//  投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	//  投机空头保证金调整系数
	SShortMarginRatioByMoney types.TThostFtdcRatioType
	//  投机空头保证金调整系数
	SShortMarginRatioByVolume types.TThostFtdcMoneyType
	//  保值空头保证金调整系数
	HShortMarginRatioByMoney types.TThostFtdcRatioType
	//  保值空头保证金调整系数
	HShortMarginRatioByVolume types.TThostFtdcMoneyType
	//  套利空头保证金调整系数
	AShortMarginRatioByMoney types.TThostFtdcRatioType
	//  套利空头保证金调整系数
	AShortMarginRatioByVolume types.TThostFtdcMoneyType
	//  是否跟随交易所收取
	IsRelative types.TThostFtdcBoolType
	//  做市商空头保证金调整系数
	MShortMarginRatioByMoney types.TThostFtdcRatioType
	//  做市商空头保证金调整系数
	MShortMarginRatioByVolume types.TThostFtdcMoneyType
	//  合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
}

// 当前期权合约手续费的详细内容
type CThostFtdcOptionInstrCommRateField struct {
	//  保留的无效字段
	Reserve1 types.TThostFtdcOldInstrumentIDType
	//  投资者范围
	InvestorRange types.TThostFtdcInvestorRangeType
	//  经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	//  投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	//  开仓手续费率
	OpenRatioByMoney types.TThostFtdcRatioType
	//  开仓手续费
	OpenRatioByVolume types.TThostFtdcRatioType
	//  平仓手续费率
	CloseRatioByMoney types.TThostFtdcRatioType
	//  平仓手续费
	CloseRatioByVolume types.TThostFtdcRatioType
	//  平今手续费率
	CloseTodayRatioByMoney types.TThostFtdcRatioType
	//  平今手续费
	CloseTodayRatioByVolume types.TThostFtdcRatioType
	//  执行手续费率
	StrikeRatioByMoney types.TThostFtdcRatioType
	//  执行手续费
	StrikeRatioByVolume types.TThostFtdcRatioType
	//  交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	//  投资单元代码
	InvestUnitID types.TThostFtdcInvestUnitIDType
	//  合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
}

// 期权交易成本
type CThostFtdcOptionInstrTradeCostField struct {
	//  经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	//  投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	//  保留的无效字段
	Reserve1 types.TThostFtdcOldInstrumentIDType
	//  投机套保标志
	HedgeFlag types.TThostFtdcHedgeFlagType
	//  期权合约保证金不变部分
	FixedMargin types.TThostFtdcMoneyType
	//  期权合约最小保证金
	MiniMargin types.TThostFtdcMoneyType
	//  期权合约权利金
	Royalty types.TThostFtdcMoneyType
	//  交易所期权合约保证金不变部分
	ExchFixedMargin types.TThostFtdcMoneyType
	//  交易所期权合约最小保证金
	ExchMiniMargin types.TThostFtdcMoneyType
	//  交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	//  投资单元代码
	InvestUnitID types.TThostFtdcInvestUnitIDType
	//  合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
}

// 期权交易成本查询
type CThostFtdcQryOptionInstrTradeCostField struct {
	//  经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	//  投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	//  保留的无效字段
	Reserve1 types.TThostFtdcOldInstrumentIDType
	//  投机套保标志
	HedgeFlag types.TThostFtdcHedgeFlagType
	//  期权合约报价
	InputPrice types.TThostFtdcPriceType
	//  标的价格,填0则用昨结算价
	UnderlyingPrice types.TThostFtdcPriceType
	//  交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	//  投资单元代码
	InvestUnitID types.TThostFtdcInvestUnitIDType
	//  合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
}

// 期权手续费率查询
type CThostFtdcQryOptionInstrCommRateField struct {
	//  经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	//  投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	//  保留的无效字段
	Reserve1 types.TThostFtdcOldInstrumentIDType
	//  交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	//  投资单元代码
	InvestUnitID types.TThostFtdcInvestUnitIDType
	//  合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
}

// 股指现货指数
type CThostFtdcIndexPriceField struct {
	//  经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	//  保留的无效字段
	Reserve1 types.TThostFtdcOldInstrumentIDType
	//  指数现货收盘价
	ClosePrice types.TThostFtdcPriceType
	//  合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
}

// 输入的执行宣告
type CThostFtdcInputExecOrderField struct {
	//  经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	//  投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	//  保留的无效字段
	Reserve1 types.TThostFtdcOldInstrumentIDType
	//  执行宣告引用
	ExecOrderRef types.TThostFtdcOrderRefType
	//  用户代码
	UserID types.TThostFtdcUserIDType
	//  数量
	Volume types.TThostFtdcVolumeType
	//  请求编号
	RequestID types.TThostFtdcRequestIDType
	//  业务单元
	BusinessUnit types.TThostFtdcBusinessUnitType
	//  开平标志
	OffsetFlag types.TThostFtdcOffsetFlagType
	//  投机套保标志
	HedgeFlag types.TThostFtdcHedgeFlagType
	//  执行类型
	ActionType types.TThostFtdcActionTypeType
	//  保留头寸申请的持仓方向
	PosiDirection types.TThostFtdcPosiDirectionType
	//  期权行权后是否保留期货头寸的标记,该字段已废弃
	ReservePositionFlag types.TThostFtdcExecOrderPositionFlagType
	//  期权行权后生成的头寸是否自动平仓
	CloseFlag types.TThostFtdcExecOrderCloseFlagType
	//  交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	//  投资单元代码
	InvestUnitID types.TThostFtdcInvestUnitIDType
	//  资金账号
	AccountID types.TThostFtdcAccountIDType
	//  币种代码
	CurrencyID types.TThostFtdcCurrencyIDType
	//  交易编码
	ClientID types.TThostFtdcClientIDType
	//  保留的无效字段
	Reserve2 types.TThostFtdcOldIPAddressType
	//  Mac地址
	MacAddress types.TThostFtdcMacAddressType
	//  合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
	//  IP地址
	IPAddress types.TThostFtdcIPAddressType
}

// 输入执行宣告操作
type CThostFtdcInputExecOrderActionField struct {
	//  经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	//  投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	//  执行宣告操作引用
	ExecOrderActionRef types.TThostFtdcOrderActionRefType
	//  执行宣告引用
	ExecOrderRef types.TThostFtdcOrderRefType
	//  请求编号
	RequestID types.TThostFtdcRequestIDType
	//  前置编号
	FrontID types.TThostFtdcFrontIDType
	//  会话编号
	SessionID types.TThostFtdcSessionIDType
	//  交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	//  执行宣告操作编号
	ExecOrderSysID types.TThostFtdcExecOrderSysIDType
	//  操作标志
	ActionFlag types.TThostFtdcActionFlagType
	//  用户代码
	UserID types.TThostFtdcUserIDType
	//  保留的无效字段
	Reserve1 types.TThostFtdcOldInstrumentIDType
	//  投资单元代码
	InvestUnitID types.TThostFtdcInvestUnitIDType
	//  保留的无效字段
	Reserve2 types.TThostFtdcOldIPAddressType
	//  Mac地址
	MacAddress types.TThostFtdcMacAddressType
	//  合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
	//  IP地址
	IPAddress types.TThostFtdcIPAddressType
}

// 执行宣告
type CThostFtdcExecOrderField struct {
	//  经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	//  投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	//  保留的无效字段
	Reserve1 types.TThostFtdcOldInstrumentIDType
	//  执行宣告引用
	ExecOrderRef types.TThostFtdcOrderRefType
	//  用户代码
	UserID types.TThostFtdcUserIDType
	//  数量
	Volume types.TThostFtdcVolumeType
	//  请求编号
	RequestID types.TThostFtdcRequestIDType
	//  业务单元
	BusinessUnit types.TThostFtdcBusinessUnitType
	//  开平标志
	OffsetFlag types.TThostFtdcOffsetFlagType
	//  投机套保标志
	HedgeFlag types.TThostFtdcHedgeFlagType
	//  执行类型
	ActionType types.TThostFtdcActionTypeType
	//  保留头寸申请的持仓方向
	PosiDirection types.TThostFtdcPosiDirectionType
	//  期权行权后是否保留期货头寸的标记,该字段已废弃
	ReservePositionFlag types.TThostFtdcExecOrderPositionFlagType
	//  期权行权后生成的头寸是否自动平仓
	CloseFlag types.TThostFtdcExecOrderCloseFlagType
	//  本地执行宣告编号
	ExecOrderLocalID types.TThostFtdcOrderLocalIDType
	//  交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	//  会员代码
	ParticipantID types.TThostFtdcParticipantIDType
	//  客户代码
	ClientID types.TThostFtdcClientIDType
	//  保留的无效字段
	Reserve2 types.TThostFtdcOldExchangeInstIDType
	//  交易所交易员代码
	TraderID types.TThostFtdcTraderIDType
	//  安装编号
	InstallID types.TThostFtdcInstallIDType
	//  执行宣告提交状态
	OrderSubmitStatus types.TThostFtdcOrderSubmitStatusType
	//  报单提示序号
	NotifySequence types.TThostFtdcSequenceNoType
	//  交易日
	TradingDay types.TThostFtdcDateType
	//  结算编号
	SettlementID types.TThostFtdcSettlementIDType
	//  执行宣告编号
	ExecOrderSysID types.TThostFtdcExecOrderSysIDType
	//  报单日期
	InsertDate types.TThostFtdcDateType
	//  插入时间
	InsertTime types.TThostFtdcTimeType
	//  撤销时间
	CancelTime types.TThostFtdcTimeType
	//  执行结果
	ExecResult types.TThostFtdcExecResultType
	//  结算会员编号
	ClearingPartID types.TThostFtdcParticipantIDType
	//  序号
	SequenceNo types.TThostFtdcSequenceNoType
	//  前置编号
	FrontID types.TThostFtdcFrontIDType
	//  会话编号
	SessionID types.TThostFtdcSessionIDType
	//  用户端产品信息
	UserProductInfo types.TThostFtdcProductInfoType
	//  状态信息
	StatusMsg types.TThostFtdcErrorMsgType
	//  操作用户代码
	ActiveUserID types.TThostFtdcUserIDType
	//  经纪公司报单编号
	BrokerExecOrderSeq types.TThostFtdcSequenceNoType
	//  营业部编号
	BranchID types.TThostFtdcBranchIDType
	//  投资单元代码
	InvestUnitID types.TThostFtdcInvestUnitIDType
	//  资金账号
	AccountID types.TThostFtdcAccountIDType
	//  币种代码
	CurrencyID types.TThostFtdcCurrencyIDType
	//  保留的无效字段
	Reserve3 types.TThostFtdcOldIPAddressType
	//  Mac地址
	MacAddress types.TThostFtdcMacAddressType
	//  合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
	//  合约在交易所的代码
	ExchangeInstID types.TThostFtdcExchangeInstIDType
	//  IP地址
	IPAddress types.TThostFtdcIPAddressType
}

// 执行宣告操作
type CThostFtdcExecOrderActionField struct {
	//  经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	//  投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	//  执行宣告操作引用
	ExecOrderActionRef types.TThostFtdcOrderActionRefType
	//  执行宣告引用
	ExecOrderRef types.TThostFtdcOrderRefType
	//  请求编号
	RequestID types.TThostFtdcRequestIDType
	//  前置编号
	FrontID types.TThostFtdcFrontIDType
	//  会话编号
	SessionID types.TThostFtdcSessionIDType
	//  交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	//  执行宣告操作编号
	ExecOrderSysID types.TThostFtdcExecOrderSysIDType
	//  操作标志
	ActionFlag types.TThostFtdcActionFlagType
	//  操作日期
	ActionDate types.TThostFtdcDateType
	//  操作时间
	ActionTime types.TThostFtdcTimeType
	//  交易所交易员代码
	TraderID types.TThostFtdcTraderIDType
	//  安装编号
	InstallID types.TThostFtdcInstallIDType
	//  本地执行宣告编号
	ExecOrderLocalID types.TThostFtdcOrderLocalIDType
	//  操作本地编号
	ActionLocalID types.TThostFtdcOrderLocalIDType
	//  会员代码
	ParticipantID types.TThostFtdcParticipantIDType
	//  客户代码
	ClientID types.TThostFtdcClientIDType
	//  业务单元
	BusinessUnit types.TThostFtdcBusinessUnitType
	//  报单操作状态
	OrderActionStatus types.TThostFtdcOrderActionStatusType
	//  用户代码
	UserID types.TThostFtdcUserIDType
	//  执行类型
	ActionType types.TThostFtdcActionTypeType
	//  状态信息
	StatusMsg types.TThostFtdcErrorMsgType
	//  保留的无效字段
	Reserve1 types.TThostFtdcOldInstrumentIDType
	//  营业部编号
	BranchID types.TThostFtdcBranchIDType
	//  投资单元代码
	InvestUnitID types.TThostFtdcInvestUnitIDType
	//  保留的无效字段
	Reserve2 types.TThostFtdcOldIPAddressType
	//  Mac地址
	MacAddress types.TThostFtdcMacAddressType
	//  合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
	//  IP地址
	IPAddress types.TThostFtdcIPAddressType
}

// 执行宣告查询
type CThostFtdcQryExecOrderField struct {
	//  经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	//  投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	//  保留的无效字段
	Reserve1 types.TThostFtdcOldInstrumentIDType
	//  交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	//  执行宣告编号
	ExecOrderSysID types.TThostFtdcExecOrderSysIDType
	//  开始时间
	InsertTimeStart types.TThostFtdcTimeType
	//  结束时间
	InsertTimeEnd types.TThostFtdcTimeType
	//  合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
}

// 交易所执行宣告信息
type CThostFtdcExchangeExecOrderField struct {
	//  数量
	Volume types.TThostFtdcVolumeType
	//  请求编号
	RequestID types.TThostFtdcRequestIDType
	//  业务单元
	BusinessUnit types.TThostFtdcBusinessUnitType
	//  开平标志
	OffsetFlag types.TThostFtdcOffsetFlagType
	//  投机套保标志
	HedgeFlag types.TThostFtdcHedgeFlagType
	//  执行类型
	ActionType types.TThostFtdcActionTypeType
	//  保留头寸申请的持仓方向
	PosiDirection types.TThostFtdcPosiDirectionType
	//  期权行权后是否保留期货头寸的标记,该字段已废弃
	ReservePositionFlag types.TThostFtdcExecOrderPositionFlagType
	//  期权行权后生成的头寸是否自动平仓
	CloseFlag types.TThostFtdcExecOrderCloseFlagType
	//  本地执行宣告编号
	ExecOrderLocalID types.TThostFtdcOrderLocalIDType
	//  交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	//  会员代码
	ParticipantID types.TThostFtdcParticipantIDType
	//  客户代码
	ClientID types.TThostFtdcClientIDType
	//  保留的无效字段
	Reserve1 types.TThostFtdcOldExchangeInstIDType
	//  交易所交易员代码
	TraderID types.TThostFtdcTraderIDType
	//  安装编号
	InstallID types.TThostFtdcInstallIDType
	//  执行宣告提交状态
	OrderSubmitStatus types.TThostFtdcOrderSubmitStatusType
	//  报单提示序号
	NotifySequence types.TThostFtdcSequenceNoType
	//  交易日
	TradingDay types.TThostFtdcDateType
	//  结算编号
	SettlementID types.TThostFtdcSettlementIDType
	//  执行宣告编号
	ExecOrderSysID types.TThostFtdcExecOrderSysIDType
	//  报单日期
	InsertDate types.TThostFtdcDateType
	//  插入时间
	InsertTime types.TThostFtdcTimeType
	//  撤销时间
	CancelTime types.TThostFtdcTimeType
	//  执行结果
	ExecResult types.TThostFtdcExecResultType
	//  结算会员编号
	ClearingPartID types.TThostFtdcParticipantIDType
	//  序号
	SequenceNo types.TThostFtdcSequenceNoType
	//  营业部编号
	BranchID types.TThostFtdcBranchIDType
	//  保留的无效字段
	Reserve2 types.TThostFtdcOldIPAddressType
	//  Mac地址
	MacAddress types.TThostFtdcMacAddressType
	//  合约在交易所的代码
	ExchangeInstID types.TThostFtdcExchangeInstIDType
	//  IP地址
	IPAddress types.TThostFtdcIPAddressType
}

// 交易所执行宣告查询
type CThostFtdcQryExchangeExecOrderField struct {
	//  会员代码
	ParticipantID types.TThostFtdcParticipantIDType
	//  客户代码
	ClientID types.TThostFtdcClientIDType
	//  保留的无效字段
	Reserve1 types.TThostFtdcOldExchangeInstIDType
	//  交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	//  交易所交易员代码
	TraderID types.TThostFtdcTraderIDType
	//  合约在交易所的代码
	ExchangeInstID types.TThostFtdcExchangeInstIDType
}

// 执行宣告操作查询
type CThostFtdcQryExecOrderActionField struct {
	//  经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	//  投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	//  交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
}

// 交易所执行宣告操作
type CThostFtdcExchangeExecOrderActionField struct {
	//  交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	//  执行宣告操作编号
	ExecOrderSysID types.TThostFtdcExecOrderSysIDType
	//  操作标志
	ActionFlag types.TThostFtdcActionFlagType
	//  操作日期
	ActionDate types.TThostFtdcDateType
	//  操作时间
	ActionTime types.TThostFtdcTimeType
	//  交易所交易员代码
	TraderID types.TThostFtdcTraderIDType
	//  安装编号
	InstallID types.TThostFtdcInstallIDType
	//  本地执行宣告编号
	ExecOrderLocalID types.TThostFtdcOrderLocalIDType
	//  操作本地编号
	ActionLocalID types.TThostFtdcOrderLocalIDType
	//  会员代码
	ParticipantID types.TThostFtdcParticipantIDType
	//  客户代码
	ClientID types.TThostFtdcClientIDType
	//  业务单元
	BusinessUnit types.TThostFtdcBusinessUnitType
	//  报单操作状态
	OrderActionStatus types.TThostFtdcOrderActionStatusType
	//  用户代码
	UserID types.TThostFtdcUserIDType
	//  执行类型
	ActionType types.TThostFtdcActionTypeType
	//  营业部编号
	BranchID types.TThostFtdcBranchIDType
	//  保留的无效字段
	Reserve1 types.TThostFtdcOldIPAddressType
	//  Mac地址
	MacAddress types.TThostFtdcMacAddressType
	//  保留的无效字段
	Reserve2 types.TThostFtdcOldExchangeInstIDType
	//  数量
	Volume types.TThostFtdcVolumeType
	//  IP地址
	IPAddress types.TThostFtdcIPAddressType
	//  合约在交易所的代码
	ExchangeInstID types.TThostFtdcExchangeInstIDType
}

// 交易所执行宣告操作查询
type CThostFtdcQryExchangeExecOrderActionField struct {
	//  会员代码
	ParticipantID types.TThostFtdcParticipantIDType
	//  客户代码
	ClientID types.TThostFtdcClientIDType
	//  交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	//  交易所交易员代码
	TraderID types.TThostFtdcTraderIDType
}

// 错误执行宣告
type CThostFtdcErrExecOrderField struct {
	//  经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	//  投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	//  保留的无效字段
	Reserve1 types.TThostFtdcOldInstrumentIDType
	//  执行宣告引用
	ExecOrderRef types.TThostFtdcOrderRefType
	//  用户代码
	UserID types.TThostFtdcUserIDType
	//  数量
	Volume types.TThostFtdcVolumeType
	//  请求编号
	RequestID types.TThostFtdcRequestIDType
	//  业务单元
	BusinessUnit types.TThostFtdcBusinessUnitType
	//  开平标志
	OffsetFlag types.TThostFtdcOffsetFlagType
	//  投机套保标志
	HedgeFlag types.TThostFtdcHedgeFlagType
	//  执行类型
	ActionType types.TThostFtdcActionTypeType
	//  保留头寸申请的持仓方向
	PosiDirection types.TThostFtdcPosiDirectionType
	//  期权行权后是否保留期货头寸的标记,该字段已废弃
	ReservePositionFlag types.TThostFtdcExecOrderPositionFlagType
	//  期权行权后生成的头寸是否自动平仓
	CloseFlag types.TThostFtdcExecOrderCloseFlagType
	//  交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	//  投资单元代码
	InvestUnitID types.TThostFtdcInvestUnitIDType
	//  资金账号
	AccountID types.TThostFtdcAccountIDType
	//  币种代码
	CurrencyID types.TThostFtdcCurrencyIDType
	//  交易编码
	ClientID types.TThostFtdcClientIDType
	//  保留的无效字段
	Reserve2 types.TThostFtdcOldIPAddressType
	//  Mac地址
	MacAddress types.TThostFtdcMacAddressType
	//  错误代码
	ErrorID types.TThostFtdcErrorIDType
	//  错误信息
	ErrorMsg types.TThostFtdcErrorMsgType
	//  合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
	//  IP地址
	IPAddress types.TThostFtdcIPAddressType
}

// 查询错误执行宣告
type CThostFtdcQryErrExecOrderField struct {
	//  经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	//  投资者代码
	InvestorID types.TThostFtdcInvestorIDType
}

// 错误执行宣告操作
type CThostFtdcErrExecOrderActionField struct {
	//  经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	//  投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	//  执行宣告操作引用
	ExecOrderActionRef types.TThostFtdcOrderActionRefType
	//  执行宣告引用
	ExecOrderRef types.TThostFtdcOrderRefType
	//  请求编号
	RequestID types.TThostFtdcRequestIDType
	//  前置编号
	FrontID types.TThostFtdcFrontIDType
	//  会话编号
	SessionID types.TThostFtdcSessionIDType
	//  交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	//  执行宣告操作编号
	ExecOrderSysID types.TThostFtdcExecOrderSysIDType
	//  操作标志
	ActionFlag types.TThostFtdcActionFlagType
	//  用户代码
	UserID types.TThostFtdcUserIDType
	//  保留的无效字段
	Reserve1 types.TThostFtdcOldInstrumentIDType
	//  投资单元代码
	InvestUnitID types.TThostFtdcInvestUnitIDType
	//  保留的无效字段
	Reserve2 types.TThostFtdcOldIPAddressType
	//  Mac地址
	MacAddress types.TThostFtdcMacAddressType
	//  错误代码
	ErrorID types.TThostFtdcErrorIDType
	//  错误信息
	ErrorMsg types.TThostFtdcErrorMsgType
	//  合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
	//  IP地址
	IPAddress types.TThostFtdcIPAddressType
}

// 查询错误执行宣告操作
type CThostFtdcQryErrExecOrderActionField struct {
	//  经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	//  投资者代码
	InvestorID types.TThostFtdcInvestorIDType
}

// 投资者期权合约交易权限
type CThostFtdcOptionInstrTradingRightField struct {
	//  保留的无效字段
	Reserve1 types.TThostFtdcOldInstrumentIDType
	//  投资者范围
	InvestorRange types.TThostFtdcInvestorRangeType
	//  经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	//  投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	//  买卖方向
	Direction types.TThostFtdcDirectionType
	//  交易权限
	TradingRight types.TThostFtdcTradingRightType
	//  合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
}

// 查询期权合约交易权限
type CThostFtdcQryOptionInstrTradingRightField struct {
	//  经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	//  投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	//  保留的无效字段
	Reserve1 types.TThostFtdcOldInstrumentIDType
	//  买卖方向
	Direction types.TThostFtdcDirectionType
	//  合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
}

// 输入的询价
type CThostFtdcInputForQuoteField struct {
	//  经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	//  投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	//  保留的无效字段
	Reserve1 types.TThostFtdcOldInstrumentIDType
	//  询价引用
	ForQuoteRef types.TThostFtdcOrderRefType
	//  用户代码
	UserID types.TThostFtdcUserIDType
	//  交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	//  投资单元代码
	InvestUnitID types.TThostFtdcInvestUnitIDType
	//  保留的无效字段
	Reserve2 types.TThostFtdcOldIPAddressType
	//  Mac地址
	MacAddress types.TThostFtdcMacAddressType
	//  合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
	//  IP地址
	IPAddress types.TThostFtdcIPAddressType
}

// 询价
type CThostFtdcForQuoteField struct {
	//  经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	//  投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	//  保留的无效字段
	Reserve1 types.TThostFtdcOldInstrumentIDType
	//  询价引用
	ForQuoteRef types.TThostFtdcOrderRefType
	//  用户代码
	UserID types.TThostFtdcUserIDType
	//  本地询价编号
	ForQuoteLocalID types.TThostFtdcOrderLocalIDType
	//  交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	//  会员代码
	ParticipantID types.TThostFtdcParticipantIDType
	//  客户代码
	ClientID types.TThostFtdcClientIDType
	//  保留的无效字段
	Reserve2 types.TThostFtdcOldExchangeInstIDType
	//  交易所交易员代码
	TraderID types.TThostFtdcTraderIDType
	//  安装编号
	InstallID types.TThostFtdcInstallIDType
	//  报单日期
	InsertDate types.TThostFtdcDateType
	//  插入时间
	InsertTime types.TThostFtdcTimeType
	//  询价状态
	ForQuoteStatus types.TThostFtdcForQuoteStatusType
	//  前置编号
	FrontID types.TThostFtdcFrontIDType
	//  会话编号
	SessionID types.TThostFtdcSessionIDType
	//  状态信息
	StatusMsg types.TThostFtdcErrorMsgType
	//  操作用户代码
	ActiveUserID types.TThostFtdcUserIDType
	//  经纪公司询价编号
	BrokerForQutoSeq types.TThostFtdcSequenceNoType
	//  投资单元代码
	InvestUnitID types.TThostFtdcInvestUnitIDType
	//  保留的无效字段
	Reserve3 types.TThostFtdcOldIPAddressType
	//  Mac地址
	MacAddress types.TThostFtdcMacAddressType
	//  合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
	//  合约在交易所的代码
	ExchangeInstID types.TThostFtdcExchangeInstIDType
	//  IP地址
	IPAddress types.TThostFtdcIPAddressType
}

// 询价查询
type CThostFtdcQryForQuoteField struct {
	//  经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	//  投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	//  保留的无效字段
	Reserve1 types.TThostFtdcOldInstrumentIDType
	//  交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	//  开始时间
	InsertTimeStart types.TThostFtdcTimeType
	//  结束时间
	InsertTimeEnd types.TThostFtdcTimeType
	//  投资单元代码
	InvestUnitID types.TThostFtdcInvestUnitIDType
	//  合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
}

// 交易所询价信息
type CThostFtdcExchangeForQuoteField struct {
	//  本地询价编号
	ForQuoteLocalID types.TThostFtdcOrderLocalIDType
	//  交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	//  会员代码
	ParticipantID types.TThostFtdcParticipantIDType
	//  客户代码
	ClientID types.TThostFtdcClientIDType
	//  保留的无效字段
	Reserve1 types.TThostFtdcOldExchangeInstIDType
	//  交易所交易员代码
	TraderID types.TThostFtdcTraderIDType
	//  安装编号
	InstallID types.TThostFtdcInstallIDType
	//  报单日期
	InsertDate types.TThostFtdcDateType
	//  插入时间
	InsertTime types.TThostFtdcTimeType
	//  询价状态
	ForQuoteStatus types.TThostFtdcForQuoteStatusType
	//  保留的无效字段
	Reserve2 types.TThostFtdcOldIPAddressType
	//  Mac地址
	MacAddress types.TThostFtdcMacAddressType
	//  合约在交易所的代码
	ExchangeInstID types.TThostFtdcExchangeInstIDType
	//  IP地址
	IPAddress types.TThostFtdcIPAddressType
}

// 交易所询价查询
type CThostFtdcQryExchangeForQuoteField struct {
	//  会员代码
	ParticipantID types.TThostFtdcParticipantIDType
	//  客户代码
	ClientID types.TThostFtdcClientIDType
	//  保留的无效字段
	Reserve1 types.TThostFtdcOldExchangeInstIDType
	//  交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	//  交易所交易员代码
	TraderID types.TThostFtdcTraderIDType
	//  合约在交易所的代码
	ExchangeInstID types.TThostFtdcExchangeInstIDType
}

// 输入的报价
type CThostFtdcInputQuoteField struct {
	//  经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	//  投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	//  保留的无效字段
	Reserve1 types.TThostFtdcOldInstrumentIDType
	//  报价引用
	QuoteRef types.TThostFtdcOrderRefType
	//  用户代码
	UserID types.TThostFtdcUserIDType
	//  卖价格
	AskPrice types.TThostFtdcPriceType
	//  买价格
	BidPrice types.TThostFtdcPriceType
	//  卖数量
	AskVolume types.TThostFtdcVolumeType
	//  买数量
	BidVolume types.TThostFtdcVolumeType
	//  请求编号
	RequestID types.TThostFtdcRequestIDType
	//  业务单元
	BusinessUnit types.TThostFtdcBusinessUnitType
	//  卖开平标志
	AskOffsetFlag types.TThostFtdcOffsetFlagType
	//  买开平标志
	BidOffsetFlag types.TThostFtdcOffsetFlagType
	//  卖投机套保标志
	AskHedgeFlag types.TThostFtdcHedgeFlagType
	//  买投机套保标志
	BidHedgeFlag types.TThostFtdcHedgeFlagType
	//  衍生卖报单引用
	AskOrderRef types.TThostFtdcOrderRefType
	//  衍生买报单引用
	BidOrderRef types.TThostFtdcOrderRefType
	//  应价编号
	ForQuoteSysID types.TThostFtdcOrderSysIDType
	//  交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	//  投资单元代码
	InvestUnitID types.TThostFtdcInvestUnitIDType
	//  交易编码
	ClientID types.TThostFtdcClientIDType
	//  保留的无效字段
	Reserve2 types.TThostFtdcOldIPAddressType
	//  Mac地址
	MacAddress types.TThostFtdcMacAddressType
	//  合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
	//  IP地址
	IPAddress types.TThostFtdcIPAddressType
	//  被顶单编号
	ReplaceSysID types.TThostFtdcOrderSysIDType
	//  有效期类型
	TimeCondition types.TThostFtdcTimeConditionType
	//  报单回显字段
	OrderMemo types.TThostFtdcOrderMemoType
	//  session上请求计数 api自动维护
	SessionReqSeq types.TThostFtdcSequenceNo12Type
}

// 输入报价操作
type CThostFtdcInputQuoteActionField struct {
	//  经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	//  投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	//  报价操作引用
	QuoteActionRef types.TThostFtdcOrderActionRefType
	//  报价引用
	QuoteRef types.TThostFtdcOrderRefType
	//  请求编号
	RequestID types.TThostFtdcRequestIDType
	//  前置编号
	FrontID types.TThostFtdcFrontIDType
	//  会话编号
	SessionID types.TThostFtdcSessionIDType
	//  交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	//  报价操作编号
	QuoteSysID types.TThostFtdcOrderSysIDType
	//  操作标志
	ActionFlag types.TThostFtdcActionFlagType
	//  用户代码
	UserID types.TThostFtdcUserIDType
	//  保留的无效字段
	Reserve1 types.TThostFtdcOldInstrumentIDType
	//  投资单元代码
	InvestUnitID types.TThostFtdcInvestUnitIDType
	//  交易编码
	ClientID types.TThostFtdcClientIDType
	//  保留的无效字段
	Reserve2 types.TThostFtdcOldIPAddressType
	//  Mac地址
	MacAddress types.TThostFtdcMacAddressType
	//  合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
	//  IP地址
	IPAddress types.TThostFtdcIPAddressType
	//  报单回显字段
	OrderMemo types.TThostFtdcOrderMemoType
	//  session上请求计数 api自动维护
	SessionReqSeq types.TThostFtdcSequenceNo12Type
}

// 报价
type CThostFtdcQuoteField struct {
	//  经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	//  投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	//  保留的无效字段
	Reserve1 types.TThostFtdcOldInstrumentIDType
	//  报价引用
	QuoteRef types.TThostFtdcOrderRefType
	//  用户代码
	UserID types.TThostFtdcUserIDType
	//  卖价格
	AskPrice types.TThostFtdcPriceType
	//  买价格
	BidPrice types.TThostFtdcPriceType
	//  卖数量
	AskVolume types.TThostFtdcVolumeType
	//  买数量
	BidVolume types.TThostFtdcVolumeType
	//  请求编号
	RequestID types.TThostFtdcRequestIDType
	//  业务单元
	BusinessUnit types.TThostFtdcBusinessUnitType
	//  卖开平标志
	AskOffsetFlag types.TThostFtdcOffsetFlagType
	//  买开平标志
	BidOffsetFlag types.TThostFtdcOffsetFlagType
	//  卖投机套保标志
	AskHedgeFlag types.TThostFtdcHedgeFlagType
	//  买投机套保标志
	BidHedgeFlag types.TThostFtdcHedgeFlagType
	//  本地报价编号
	QuoteLocalID types.TThostFtdcOrderLocalIDType
	//  交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	//  会员代码
	ParticipantID types.TThostFtdcParticipantIDType
	//  客户代码
	ClientID types.TThostFtdcClientIDType
	//  保留的无效字段
	Reserve2 types.TThostFtdcOldExchangeInstIDType
	//  交易所交易员代码
	TraderID types.TThostFtdcTraderIDType
	//  安装编号
	InstallID types.TThostFtdcInstallIDType
	//  报价提示序号
	NotifySequence types.TThostFtdcSequenceNoType
	//  报价提交状态
	OrderSubmitStatus types.TThostFtdcOrderSubmitStatusType
	//  交易日
	TradingDay types.TThostFtdcDateType
	//  结算编号
	SettlementID types.TThostFtdcSettlementIDType
	//  报价编号
	QuoteSysID types.TThostFtdcOrderSysIDType
	//  报单日期
	InsertDate types.TThostFtdcDateType
	//  插入时间
	InsertTime types.TThostFtdcTimeType
	//  撤销时间
	CancelTime types.TThostFtdcTimeType
	//  报价状态
	QuoteStatus types.TThostFtdcOrderStatusType
	//  结算会员编号
	ClearingPartID types.TThostFtdcParticipantIDType
	//  序号
	SequenceNo types.TThostFtdcSequenceNoType
	//  卖方报单编号
	AskOrderSysID types.TThostFtdcOrderSysIDType
	//  买方报单编号
	BidOrderSysID types.TThostFtdcOrderSysIDType
	//  前置编号
	FrontID types.TThostFtdcFrontIDType
	//  会话编号
	SessionID types.TThostFtdcSessionIDType
	//  用户端产品信息
	UserProductInfo types.TThostFtdcProductInfoType
	//  状态信息
	StatusMsg types.TThostFtdcErrorMsgType
	//  操作用户代码
	ActiveUserID types.TThostFtdcUserIDType
	//  经纪公司报价编号
	BrokerQuoteSeq types.TThostFtdcSequenceNoType
	//  衍生卖报单引用
	AskOrderRef types.TThostFtdcOrderRefType
	//  衍生买报单引用
	BidOrderRef types.TThostFtdcOrderRefType
	//  应价编号
	ForQuoteSysID types.TThostFtdcOrderSysIDType
	//  营业部编号
	BranchID types.TThostFtdcBranchIDType
	//  投资单元代码
	InvestUnitID types.TThostFtdcInvestUnitIDType
	//  资金账号
	AccountID types.TThostFtdcAccountIDType
	//  币种代码
	CurrencyID types.TThostFtdcCurrencyIDType
	//  保留的无效字段
	Reserve3 types.TThostFtdcOldIPAddressType
	//  Mac地址
	MacAddress types.TThostFtdcMacAddressType
	//  合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
	//  合约在交易所的代码
	ExchangeInstID types.TThostFtdcExchangeInstIDType
	//  IP地址
	IPAddress types.TThostFtdcIPAddressType
	//  被顶单编号
	ReplaceSysID types.TThostFtdcOrderSysIDType
	//  有效期类型
	TimeCondition types.TThostFtdcTimeConditionType
	//  报单回显字段
	OrderMemo types.TThostFtdcOrderMemoType
	//  session上请求计数 api自动维护
	SessionReqSeq types.TThostFtdcSequenceNo12Type
}

// 报价操作
type CThostFtdcQuoteActionField struct {
	//  经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	//  投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	//  报价操作引用
	QuoteActionRef types.TThostFtdcOrderActionRefType
	//  报价引用
	QuoteRef types.TThostFtdcOrderRefType
	//  请求编号
	RequestID types.TThostFtdcRequestIDType
	//  前置编号
	FrontID types.TThostFtdcFrontIDType
	//  会话编号
	SessionID types.TThostFtdcSessionIDType
	//  交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	//  报价操作编号
	QuoteSysID types.TThostFtdcOrderSysIDType
	//  操作标志
	ActionFlag types.TThostFtdcActionFlagType
	//  操作日期
	ActionDate types.TThostFtdcDateType
	//  操作时间
	ActionTime types.TThostFtdcTimeType
	//  交易所交易员代码
	TraderID types.TThostFtdcTraderIDType
	//  安装编号
	InstallID types.TThostFtdcInstallIDType
	//  本地报价编号
	QuoteLocalID types.TThostFtdcOrderLocalIDType
	//  操作本地编号
	ActionLocalID types.TThostFtdcOrderLocalIDType
	//  会员代码
	ParticipantID types.TThostFtdcParticipantIDType
	//  客户代码
	ClientID types.TThostFtdcClientIDType
	//  业务单元
	BusinessUnit types.TThostFtdcBusinessUnitType
	//  报单操作状态
	OrderActionStatus types.TThostFtdcOrderActionStatusType
	//  用户代码
	UserID types.TThostFtdcUserIDType
	//  状态信息
	StatusMsg types.TThostFtdcErrorMsgType
	//  保留的无效字段
	Reserve1 types.TThostFtdcOldInstrumentIDType
	//  营业部编号
	BranchID types.TThostFtdcBranchIDType
	//  投资单元代码
	InvestUnitID types.TThostFtdcInvestUnitIDType
	//  保留的无效字段
	Reserve2 types.TThostFtdcOldIPAddressType
	//  Mac地址
	MacAddress types.TThostFtdcMacAddressType
	//  合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
	//  IP地址
	IPAddress types.TThostFtdcIPAddressType
	//  报单回显字段
	OrderMemo types.TThostFtdcOrderMemoType
	//  session上请求计数 api自动维护
	SessionReqSeq types.TThostFtdcSequenceNo12Type
}

// 报价查询
type CThostFtdcQryQuoteField struct {
	//  经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	//  投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	//  保留的无效字段
	Reserve1 types.TThostFtdcOldInstrumentIDType
	//  交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	//  报价编号
	QuoteSysID types.TThostFtdcOrderSysIDType
	//  开始时间
	InsertTimeStart types.TThostFtdcTimeType
	//  结束时间
	InsertTimeEnd types.TThostFtdcTimeType
	//  投资单元代码
	InvestUnitID types.TThostFtdcInvestUnitIDType
	//  合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
}

// 交易所报价信息
type CThostFtdcExchangeQuoteField struct {
	//  卖价格
	AskPrice types.TThostFtdcPriceType
	//  买价格
	BidPrice types.TThostFtdcPriceType
	//  卖数量
	AskVolume types.TThostFtdcVolumeType
	//  买数量
	BidVolume types.TThostFtdcVolumeType
	//  请求编号
	RequestID types.TThostFtdcRequestIDType
	//  业务单元
	BusinessUnit types.TThostFtdcBusinessUnitType
	//  卖开平标志
	AskOffsetFlag types.TThostFtdcOffsetFlagType
	//  买开平标志
	BidOffsetFlag types.TThostFtdcOffsetFlagType
	//  卖投机套保标志
	AskHedgeFlag types.TThostFtdcHedgeFlagType
	//  买投机套保标志
	BidHedgeFlag types.TThostFtdcHedgeFlagType
	//  本地报价编号
	QuoteLocalID types.TThostFtdcOrderLocalIDType
	//  交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	//  会员代码
	ParticipantID types.TThostFtdcParticipantIDType
	//  客户代码
	ClientID types.TThostFtdcClientIDType
	//  保留的无效字段
	Reserve1 types.TThostFtdcOldExchangeInstIDType
	//  交易所交易员代码
	TraderID types.TThostFtdcTraderIDType
	//  安装编号
	InstallID types.TThostFtdcInstallIDType
	//  报价提示序号
	NotifySequence types.TThostFtdcSequenceNoType
	//  报价提交状态
	OrderSubmitStatus types.TThostFtdcOrderSubmitStatusType
	//  交易日
	TradingDay types.TThostFtdcDateType
	//  结算编号
	SettlementID types.TThostFtdcSettlementIDType
	//  报价编号
	QuoteSysID types.TThostFtdcOrderSysIDType
	//  报单日期
	InsertDate types.TThostFtdcDateType
	//  插入时间
	InsertTime types.TThostFtdcTimeType
	//  撤销时间
	CancelTime types.TThostFtdcTimeType
	//  报价状态
	QuoteStatus types.TThostFtdcOrderStatusType
	//  结算会员编号
	ClearingPartID types.TThostFtdcParticipantIDType
	//  序号
	SequenceNo types.TThostFtdcSequenceNoType
	//  卖方报单编号
	AskOrderSysID types.TThostFtdcOrderSysIDType
	//  买方报单编号
	BidOrderSysID types.TThostFtdcOrderSysIDType
	//  应价编号
	ForQuoteSysID types.TThostFtdcOrderSysIDType
	//  营业部编号
	BranchID types.TThostFtdcBranchIDType
	//  保留的无效字段
	Reserve2 types.TThostFtdcOldIPAddressType
	//  Mac地址
	MacAddress types.TThostFtdcMacAddressType
	//  合约在交易所的代码
	ExchangeInstID types.TThostFtdcExchangeInstIDType
	//  IP地址
	IPAddress types.TThostFtdcIPAddressType
	//  有效期类型
	TimeCondition types.TThostFtdcTimeConditionType
}

// 交易所报价查询
type CThostFtdcQryExchangeQuoteField struct {
	//  会员代码
	ParticipantID types.TThostFtdcParticipantIDType
	//  客户代码
	ClientID types.TThostFtdcClientIDType
	//  保留的无效字段
	Reserve1 types.TThostFtdcOldExchangeInstIDType
	//  交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	//  交易所交易员代码
	TraderID types.TThostFtdcTraderIDType
	//  合约在交易所的代码
	ExchangeInstID types.TThostFtdcExchangeInstIDType
}

// 报价操作查询
type CThostFtdcQryQuoteActionField struct {
	//  经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	//  投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	//  交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
}

// 交易所报价操作
type CThostFtdcExchangeQuoteActionField struct {
	//  交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	//  报价操作编号
	QuoteSysID types.TThostFtdcOrderSysIDType
	//  操作标志
	ActionFlag types.TThostFtdcActionFlagType
	//  操作日期
	ActionDate types.TThostFtdcDateType
	//  操作时间
	ActionTime types.TThostFtdcTimeType
	//  交易所交易员代码
	TraderID types.TThostFtdcTraderIDType
	//  安装编号
	InstallID types.TThostFtdcInstallIDType
	//  本地报价编号
	QuoteLocalID types.TThostFtdcOrderLocalIDType
	//  操作本地编号
	ActionLocalID types.TThostFtdcOrderLocalIDType
	//  会员代码
	ParticipantID types.TThostFtdcParticipantIDType
	//  客户代码
	ClientID types.TThostFtdcClientIDType
	//  业务单元
	BusinessUnit types.TThostFtdcBusinessUnitType
	//  报单操作状态
	OrderActionStatus types.TThostFtdcOrderActionStatusType
	//  用户代码
	UserID types.TThostFtdcUserIDType
	//  保留的无效字段
	Reserve1 types.TThostFtdcOldIPAddressType
	//  Mac地址
	MacAddress types.TThostFtdcMacAddressType
	//  IP地址
	IPAddress types.TThostFtdcIPAddressType
}

// 交易所报价操作查询
type CThostFtdcQryExchangeQuoteActionField struct {
	//  会员代码
	ParticipantID types.TThostFtdcParticipantIDType
	//  客户代码
	ClientID types.TThostFtdcClientIDType
	//  交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	//  交易所交易员代码
	TraderID types.TThostFtdcTraderIDType
}

// 期权合约delta值
type CThostFtdcOptionInstrDeltaField struct {
	//  保留的无效字段
	Reserve1 types.TThostFtdcOldInstrumentIDType
	//  投资者范围
	InvestorRange types.TThostFtdcInvestorRangeType
	//  经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	//  投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	//  Delta值
	Delta types.TThostFtdcRatioType
	//  合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
}

// 发给做市商的询价请求
type CThostFtdcForQuoteRspField struct {
	//  交易日
	TradingDay types.TThostFtdcDateType
	//  保留的无效字段
	Reserve1 types.TThostFtdcOldInstrumentIDType
	//  询价编号
	ForQuoteSysID types.TThostFtdcOrderSysIDType
	//  询价时间
	ForQuoteTime types.TThostFtdcTimeType
	//  业务日期
	ActionDay types.TThostFtdcDateType
	//  交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	//  合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
}

// 当前期权合约执行偏移值的详细内容
type CThostFtdcStrikeOffsetField struct {
	//  保留的无效字段
	Reserve1 types.TThostFtdcOldInstrumentIDType
	//  投资者范围
	InvestorRange types.TThostFtdcInvestorRangeType
	//  经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	//  投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	//  执行偏移值
	Offset types.TThostFtdcMoneyType
	//  执行偏移类型
	OffsetType types.TThostFtdcStrikeOffsetTypeType
	//  合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
}

// 期权执行偏移值查询
type CThostFtdcQryStrikeOffsetField struct {
	//  经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	//  投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	//  保留的无效字段
	Reserve1 types.TThostFtdcOldInstrumentIDType
	//  合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
}

// 输入批量报单操作
type CThostFtdcInputBatchOrderActionField struct {
	//  经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	//  投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	//  报单操作引用
	OrderActionRef types.TThostFtdcOrderActionRefType
	//  请求编号
	RequestID types.TThostFtdcRequestIDType
	//  前置编号
	FrontID types.TThostFtdcFrontIDType
	//  会话编号
	SessionID types.TThostFtdcSessionIDType
	//  交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	//  用户代码
	UserID types.TThostFtdcUserIDType
	//  投资单元代码
	InvestUnitID types.TThostFtdcInvestUnitIDType
	//  保留的无效字段
	Reserve1 types.TThostFtdcOldIPAddressType
	//  Mac地址
	MacAddress types.TThostFtdcMacAddressType
	//  IP地址
	IPAddress types.TThostFtdcIPAddressType
}

// 批量报单操作
type CThostFtdcBatchOrderActionField struct {
	//  经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	//  投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	//  报单操作引用
	OrderActionRef types.TThostFtdcOrderActionRefType
	//  请求编号
	RequestID types.TThostFtdcRequestIDType
	//  前置编号
	FrontID types.TThostFtdcFrontIDType
	//  会话编号
	SessionID types.TThostFtdcSessionIDType
	//  交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	//  操作日期
	ActionDate types.TThostFtdcDateType
	//  操作时间
	ActionTime types.TThostFtdcTimeType
	//  交易所交易员代码
	TraderID types.TThostFtdcTraderIDType
	//  安装编号
	InstallID types.TThostFtdcInstallIDType
	//  操作本地编号
	ActionLocalID types.TThostFtdcOrderLocalIDType
	//  会员代码
	ParticipantID types.TThostFtdcParticipantIDType
	//  客户代码
	ClientID types.TThostFtdcClientIDType
	//  业务单元
	BusinessUnit types.TThostFtdcBusinessUnitType
	//  报单操作状态
	OrderActionStatus types.TThostFtdcOrderActionStatusType
	//  用户代码
	UserID types.TThostFtdcUserIDType
	//  状态信息
	StatusMsg types.TThostFtdcErrorMsgType
	//  投资单元代码
	InvestUnitID types.TThostFtdcInvestUnitIDType
	//  保留的无效字段
	Reserve1 types.TThostFtdcOldIPAddressType
	//  Mac地址
	MacAddress types.TThostFtdcMacAddressType
	//  IP地址
	IPAddress types.TThostFtdcIPAddressType
}

// 交易所批量报单操作
type CThostFtdcExchangeBatchOrderActionField struct {
	//  交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	//  操作日期
	ActionDate types.TThostFtdcDateType
	//  操作时间
	ActionTime types.TThostFtdcTimeType
	//  交易所交易员代码
	TraderID types.TThostFtdcTraderIDType
	//  安装编号
	InstallID types.TThostFtdcInstallIDType
	//  操作本地编号
	ActionLocalID types.TThostFtdcOrderLocalIDType
	//  会员代码
	ParticipantID types.TThostFtdcParticipantIDType
	//  客户代码
	ClientID types.TThostFtdcClientIDType
	//  业务单元
	BusinessUnit types.TThostFtdcBusinessUnitType
	//  报单操作状态
	OrderActionStatus types.TThostFtdcOrderActionStatusType
	//  用户代码
	UserID types.TThostFtdcUserIDType
	//  保留的无效字段
	Reserve1 types.TThostFtdcOldIPAddressType
	//  Mac地址
	MacAddress types.TThostFtdcMacAddressType
	//  IP地址
	IPAddress types.TThostFtdcIPAddressType
}

// 查询批量报单操作
type CThostFtdcQryBatchOrderActionField struct {
	//  经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	//  投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	//  交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
}

// 组合合约安全系数
type CThostFtdcCombInstrumentGuardField struct {
	//  经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	//  保留的无效字段
	Reserve1 types.TThostFtdcOldInstrumentIDType
	//
	GuarantRatio types.TThostFtdcRatioType
	//  交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	//  合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
}

// 组合合约安全系数查询
type CThostFtdcQryCombInstrumentGuardField struct {
	//  经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	//  保留的无效字段
	Reserve1 types.TThostFtdcOldInstrumentIDType
	//  交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	//  合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
}

// 输入的申请组合
type CThostFtdcInputCombActionField struct {
	//  经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	//  投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	//  保留的无效字段
	Reserve1 types.TThostFtdcOldInstrumentIDType
	//  组合引用
	CombActionRef types.TThostFtdcOrderRefType
	//  用户代码
	UserID types.TThostFtdcUserIDType
	//  买卖方向
	Direction types.TThostFtdcDirectionType
	//  数量
	Volume types.TThostFtdcVolumeType
	//  组合指令方向
	CombDirection types.TThostFtdcCombDirectionType
	//  投机套保标志
	HedgeFlag types.TThostFtdcHedgeFlagType
	//  交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	//  保留的无效字段
	Reserve2 types.TThostFtdcOldIPAddressType
	//  Mac地址
	MacAddress types.TThostFtdcMacAddressType
	//  投资单元代码
	InvestUnitID types.TThostFtdcInvestUnitIDType
	//  前置编号
	FrontID types.TThostFtdcFrontIDType
	//  会话编号
	SessionID types.TThostFtdcSessionIDType
	//  合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
	//  IP地址
	IPAddress types.TThostFtdcIPAddressType
}

// 申请组合
type CThostFtdcCombActionField struct {
	//  经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	//  投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	//  保留的无效字段
	Reserve1 types.TThostFtdcOldInstrumentIDType
	//  组合引用
	CombActionRef types.TThostFtdcOrderRefType
	//  用户代码
	UserID types.TThostFtdcUserIDType
	//  买卖方向
	Direction types.TThostFtdcDirectionType
	//  数量
	Volume types.TThostFtdcVolumeType
	//  组合指令方向
	CombDirection types.TThostFtdcCombDirectionType
	//  投机套保标志
	HedgeFlag types.TThostFtdcHedgeFlagType
	//  本地申请组合编号
	ActionLocalID types.TThostFtdcOrderLocalIDType
	//  交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	//  会员代码
	ParticipantID types.TThostFtdcParticipantIDType
	//  客户代码
	ClientID types.TThostFtdcClientIDType
	//  保留的无效字段
	Reserve2 types.TThostFtdcOldExchangeInstIDType
	//  交易所交易员代码
	TraderID types.TThostFtdcTraderIDType
	//  安装编号
	InstallID types.TThostFtdcInstallIDType
	//  组合状态
	ActionStatus types.TThostFtdcOrderActionStatusType
	//  报单提示序号
	NotifySequence types.TThostFtdcSequenceNoType
	//  交易日
	TradingDay types.TThostFtdcDateType
	//  结算编号
	SettlementID types.TThostFtdcSettlementIDType
	//  序号
	SequenceNo types.TThostFtdcSequenceNoType
	//  前置编号
	FrontID types.TThostFtdcFrontIDType
	//  会话编号
	SessionID types.TThostFtdcSessionIDType
	//  用户端产品信息
	UserProductInfo types.TThostFtdcProductInfoType
	//  状态信息
	StatusMsg types.TThostFtdcErrorMsgType
	//  保留的无效字段
	Reserve3 types.TThostFtdcOldIPAddressType
	//  Mac地址
	MacAddress types.TThostFtdcMacAddressType
	//  组合编号
	ComTradeID types.TThostFtdcTradeIDType
	//  营业部编号
	BranchID types.TThostFtdcBranchIDType
	//  投资单元代码
	InvestUnitID types.TThostFtdcInvestUnitIDType
	//  合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
	//  合约在交易所的代码
	ExchangeInstID types.TThostFtdcExchangeInstIDType
	//  IP地址
	IPAddress types.TThostFtdcIPAddressType
}

// 申请组合查询
type CThostFtdcQryCombActionField struct {
	//  经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	//  投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	//  保留的无效字段
	Reserve1 types.TThostFtdcOldInstrumentIDType
	//  交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	//  投资单元代码
	InvestUnitID types.TThostFtdcInvestUnitIDType
	//  合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
}

// 交易所申请组合信息
type CThostFtdcExchangeCombActionField struct {
	//  买卖方向
	Direction types.TThostFtdcDirectionType
	//  数量
	Volume types.TThostFtdcVolumeType
	//  组合指令方向
	CombDirection types.TThostFtdcCombDirectionType
	//  投机套保标志
	HedgeFlag types.TThostFtdcHedgeFlagType
	//  本地申请组合编号
	ActionLocalID types.TThostFtdcOrderLocalIDType
	//  交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	//  会员代码
	ParticipantID types.TThostFtdcParticipantIDType
	//  客户代码
	ClientID types.TThostFtdcClientIDType
	//  保留的无效字段
	Reserve1 types.TThostFtdcOldExchangeInstIDType
	//  交易所交易员代码
	TraderID types.TThostFtdcTraderIDType
	//  安装编号
	InstallID types.TThostFtdcInstallIDType
	//  组合状态
	ActionStatus types.TThostFtdcOrderActionStatusType
	//  报单提示序号
	NotifySequence types.TThostFtdcSequenceNoType
	//  交易日
	TradingDay types.TThostFtdcDateType
	//  结算编号
	SettlementID types.TThostFtdcSettlementIDType
	//  序号
	SequenceNo types.TThostFtdcSequenceNoType
	//  保留的无效字段
	Reserve2 types.TThostFtdcOldIPAddressType
	//  Mac地址
	MacAddress types.TThostFtdcMacAddressType
	//  组合编号
	ComTradeID types.TThostFtdcTradeIDType
	//  营业部编号
	BranchID types.TThostFtdcBranchIDType
	//  合约在交易所的代码
	ExchangeInstID types.TThostFtdcExchangeInstIDType
	//  IP地址
	IPAddress types.TThostFtdcIPAddressType
}

// 交易所申请组合查询
type CThostFtdcQryExchangeCombActionField struct {
	//  会员代码
	ParticipantID types.TThostFtdcParticipantIDType
	//  客户代码
	ClientID types.TThostFtdcClientIDType
	//  保留的无效字段
	Reserve1 types.TThostFtdcOldExchangeInstIDType
	//  交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	//  交易所交易员代码
	TraderID types.TThostFtdcTraderIDType
	//  合约在交易所的代码
	ExchangeInstID types.TThostFtdcExchangeInstIDType
}

// 产品报价汇率
type CThostFtdcProductExchRateField struct {
	//  保留的无效字段
	Reserve1 types.TThostFtdcOldInstrumentIDType
	//  报价币种类型
	QuoteCurrencyID types.TThostFtdcCurrencyIDType
	//  汇率
	ExchangeRate types.TThostFtdcExchangeRateType
	//  交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	//  产品代码
	ProductID types.TThostFtdcInstrumentIDType
}

// 产品报价汇率查询
type CThostFtdcQryProductExchRateField struct {
	//  保留的无效字段
	Reserve1 types.TThostFtdcOldInstrumentIDType
	//  交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	//  产品代码
	ProductID types.TThostFtdcInstrumentIDType
}

// 查询询价价差参数
type CThostFtdcQryForQuoteParamField struct {
	//  经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	//  保留的无效字段
	Reserve1 types.TThostFtdcOldInstrumentIDType
	//  交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	//  合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
}

// 询价价差参数
type CThostFtdcForQuoteParamField struct {
	//  经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	//  保留的无效字段
	Reserve1 types.TThostFtdcOldInstrumentIDType
	//  交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	//  最新价
	LastPrice types.TThostFtdcPriceType
	//  价差
	PriceInterval types.TThostFtdcPriceType
	//  合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
}

// 当前做市商期权合约手续费的详细内容
type CThostFtdcMMOptionInstrCommRateField struct {
	//  保留的无效字段
	Reserve1 types.TThostFtdcOldInstrumentIDType
	//  投资者范围
	InvestorRange types.TThostFtdcInvestorRangeType
	//  经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	//  投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	//  开仓手续费率
	OpenRatioByMoney types.TThostFtdcRatioType
	//  开仓手续费
	OpenRatioByVolume types.TThostFtdcRatioType
	//  平仓手续费率
	CloseRatioByMoney types.TThostFtdcRatioType
	//  平仓手续费
	CloseRatioByVolume types.TThostFtdcRatioType
	//  平今手续费率
	CloseTodayRatioByMoney types.TThostFtdcRatioType
	//  平今手续费
	CloseTodayRatioByVolume types.TThostFtdcRatioType
	//  执行手续费率
	StrikeRatioByMoney types.TThostFtdcRatioType
	//  执行手续费
	StrikeRatioByVolume types.TThostFtdcRatioType
	//  合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
}

// 做市商期权手续费率查询
type CThostFtdcQryMMOptionInstrCommRateField struct {
	//  经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	//  投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	//  保留的无效字段
	Reserve1 types.TThostFtdcOldInstrumentIDType
	//  合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
}

// 做市商合约手续费率
type CThostFtdcMMInstrumentCommissionRateField struct {
	//  保留的无效字段
	Reserve1 types.TThostFtdcOldInstrumentIDType
	//  投资者范围
	InvestorRange types.TThostFtdcInvestorRangeType
	//  经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	//  投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	//  开仓手续费率
	OpenRatioByMoney types.TThostFtdcRatioType
	//  开仓手续费
	OpenRatioByVolume types.TThostFtdcRatioType
	//  平仓手续费率
	CloseRatioByMoney types.TThostFtdcRatioType
	//  平仓手续费
	CloseRatioByVolume types.TThostFtdcRatioType
	//  平今手续费率
	CloseTodayRatioByMoney types.TThostFtdcRatioType
	//  平今手续费
	CloseTodayRatioByVolume types.TThostFtdcRatioType
	//  合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
}

// 查询做市商合约手续费率
type CThostFtdcQryMMInstrumentCommissionRateField struct {
	//  经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	//  投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	//  保留的无效字段
	Reserve1 types.TThostFtdcOldInstrumentIDType
	//  合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
}

// 当前报单手续费的详细内容
type CThostFtdcInstrumentOrderCommRateField struct {
	//  保留的无效字段
	Reserve1 types.TThostFtdcOldInstrumentIDType
	//  投资者范围
	InvestorRange types.TThostFtdcInvestorRangeType
	//  经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	//  投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	//  投机套保标志
	HedgeFlag types.TThostFtdcHedgeFlagType
	//  报单手续费
	OrderCommByVolume types.TThostFtdcRatioType
	//  撤单手续费
	OrderActionCommByVolume types.TThostFtdcRatioType
	//  交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	//  投资单元代码
	InvestUnitID types.TThostFtdcInvestUnitIDType
	//  合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
	//  报单手续费
	OrderCommByTrade types.TThostFtdcRatioType
	//  撤单手续费
	OrderActionCommByTrade types.TThostFtdcRatioType
}

// 报单手续费率查询
type CThostFtdcQryInstrumentOrderCommRateField struct {
	//  经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	//  投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	//  保留的无效字段
	Reserve1 types.TThostFtdcOldInstrumentIDType
	//  合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
}

// 交易参数
type CThostFtdcTradeParamField struct {
	//  经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	//  参数代码
	TradeParamID types.TThostFtdcTradeParamIDType
	//  参数代码值
	TradeParamValue types.TThostFtdcSettlementParamValueType
	//  备注
	Memo types.TThostFtdcMemoType
}

// 合约保证金率调整
type CThostFtdcInstrumentMarginRateULField struct {
	//  保留的无效字段
	Reserve1 types.TThostFtdcOldInstrumentIDType
	//  投资者范围
	InvestorRange types.TThostFtdcInvestorRangeType
	//  经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	//  投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	//  投机套保标志
	HedgeFlag types.TThostFtdcHedgeFlagType
	//  多头保证金率
	LongMarginRatioByMoney types.TThostFtdcRatioType
	//  多头保证金费
	LongMarginRatioByVolume types.TThostFtdcMoneyType
	//  空头保证金率
	ShortMarginRatioByMoney types.TThostFtdcRatioType
	//  空头保证金费
	ShortMarginRatioByVolume types.TThostFtdcMoneyType
	//  合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
}

// 期货持仓限制参数
type CThostFtdcFutureLimitPosiParamField struct {
	//  投资者范围
	InvestorRange types.TThostFtdcInvestorRangeType
	//  经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	//  投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	//  保留的无效字段
	Reserve1 types.TThostFtdcOldInstrumentIDType
	//  当日投机开仓数量限制
	SpecOpenVolume types.TThostFtdcVolumeType
	//  当日套利开仓数量限制
	ArbiOpenVolume types.TThostFtdcVolumeType
	//  当日投机+套利开仓数量限制
	OpenVolume types.TThostFtdcVolumeType
	//  产品代码
	ProductID types.TThostFtdcInstrumentIDType
}

// 禁止登录IP
type CThostFtdcLoginForbiddenIPField struct {
	//  保留的无效字段
	Reserve1 types.TThostFtdcOldIPAddressType
	//  IP地址
	IPAddress types.TThostFtdcIPAddressType
}

// IP列表
type CThostFtdcIPListField struct {
	//  保留的无效字段
	Reserve1 types.TThostFtdcOldIPAddressType
	//  是否白名单
	IsWhite types.TThostFtdcBoolType
	//  IP地址
	IPAddress types.TThostFtdcIPAddressType
}

// 输入的期权自对冲
type CThostFtdcInputOptionSelfCloseField struct {
	//  经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	//  投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	//  保留的无效字段
	Reserve1 types.TThostFtdcOldInstrumentIDType
	//  期权自对冲引用
	OptionSelfCloseRef types.TThostFtdcOrderRefType
	//  用户代码
	UserID types.TThostFtdcUserIDType
	//  数量
	Volume types.TThostFtdcVolumeType
	//  请求编号
	RequestID types.TThostFtdcRequestIDType
	//  业务单元
	BusinessUnit types.TThostFtdcBusinessUnitType
	//  投机套保标志
	HedgeFlag types.TThostFtdcHedgeFlagType
	//  期权行权的头寸是否自对冲
	OptSelfCloseFlag types.TThostFtdcOptSelfCloseFlagType
	//  交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	//  投资单元代码
	InvestUnitID types.TThostFtdcInvestUnitIDType
	//  资金账号
	AccountID types.TThostFtdcAccountIDType
	//  币种代码
	CurrencyID types.TThostFtdcCurrencyIDType
	//  交易编码
	ClientID types.TThostFtdcClientIDType
	//  保留的无效字段
	Reserve2 types.TThostFtdcOldIPAddressType
	//  Mac地址
	MacAddress types.TThostFtdcMacAddressType
	//  合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
	//  IP地址
	IPAddress types.TThostFtdcIPAddressType
}

// 输入期权自对冲操作
type CThostFtdcInputOptionSelfCloseActionField struct {
	//  经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	//  投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	//  期权自对冲操作引用
	OptionSelfCloseActionRef types.TThostFtdcOrderActionRefType
	//  期权自对冲引用
	OptionSelfCloseRef types.TThostFtdcOrderRefType
	//  请求编号
	RequestID types.TThostFtdcRequestIDType
	//  前置编号
	FrontID types.TThostFtdcFrontIDType
	//  会话编号
	SessionID types.TThostFtdcSessionIDType
	//  交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	//  期权自对冲操作编号
	OptionSelfCloseSysID types.TThostFtdcOrderSysIDType
	//  操作标志
	ActionFlag types.TThostFtdcActionFlagType
	//  用户代码
	UserID types.TThostFtdcUserIDType
	//  保留的无效字段
	Reserve1 types.TThostFtdcOldInstrumentIDType
	//  投资单元代码
	InvestUnitID types.TThostFtdcInvestUnitIDType
	//  保留的无效字段
	Reserve2 types.TThostFtdcOldIPAddressType
	//  Mac地址
	MacAddress types.TThostFtdcMacAddressType
	//  合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
	//  IP地址
	IPAddress types.TThostFtdcIPAddressType
}

// 期权自对冲
type CThostFtdcOptionSelfCloseField struct {
	//  经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	//  投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	//  保留的无效字段
	Reserve1 types.TThostFtdcOldInstrumentIDType
	//  期权自对冲引用
	OptionSelfCloseRef types.TThostFtdcOrderRefType
	//  用户代码
	UserID types.TThostFtdcUserIDType
	//  数量
	Volume types.TThostFtdcVolumeType
	//  请求编号
	RequestID types.TThostFtdcRequestIDType
	//  业务单元
	BusinessUnit types.TThostFtdcBusinessUnitType
	//  投机套保标志
	HedgeFlag types.TThostFtdcHedgeFlagType
	//  期权行权的头寸是否自对冲
	OptSelfCloseFlag types.TThostFtdcOptSelfCloseFlagType
	//  本地期权自对冲编号
	OptionSelfCloseLocalID types.TThostFtdcOrderLocalIDType
	//  交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	//  会员代码
	ParticipantID types.TThostFtdcParticipantIDType
	//  客户代码
	ClientID types.TThostFtdcClientIDType
	//  保留的无效字段
	Reserve2 types.TThostFtdcOldExchangeInstIDType
	//  交易所交易员代码
	TraderID types.TThostFtdcTraderIDType
	//  安装编号
	InstallID types.TThostFtdcInstallIDType
	//  期权自对冲提交状态
	OrderSubmitStatus types.TThostFtdcOrderSubmitStatusType
	//  报单提示序号
	NotifySequence types.TThostFtdcSequenceNoType
	//  交易日
	TradingDay types.TThostFtdcDateType
	//  结算编号
	SettlementID types.TThostFtdcSettlementIDType
	//  期权自对冲编号
	OptionSelfCloseSysID types.TThostFtdcOrderSysIDType
	//  报单日期
	InsertDate types.TThostFtdcDateType
	//  插入时间
	InsertTime types.TThostFtdcTimeType
	//  撤销时间
	CancelTime types.TThostFtdcTimeType
	//  自对冲结果
	ExecResult types.TThostFtdcExecResultType
	//  结算会员编号
	ClearingPartID types.TThostFtdcParticipantIDType
	//  序号
	SequenceNo types.TThostFtdcSequenceNoType
	//  前置编号
	FrontID types.TThostFtdcFrontIDType
	//  会话编号
	SessionID types.TThostFtdcSessionIDType
	//  用户端产品信息
	UserProductInfo types.TThostFtdcProductInfoType
	//  状态信息
	StatusMsg types.TThostFtdcErrorMsgType
	//  操作用户代码
	ActiveUserID types.TThostFtdcUserIDType
	//  经纪公司报单编号
	BrokerOptionSelfCloseSeq types.TThostFtdcSequenceNoType
	//  营业部编号
	BranchID types.TThostFtdcBranchIDType
	//  投资单元代码
	InvestUnitID types.TThostFtdcInvestUnitIDType
	//  资金账号
	AccountID types.TThostFtdcAccountIDType
	//  币种代码
	CurrencyID types.TThostFtdcCurrencyIDType
	//  保留的无效字段
	Reserve3 types.TThostFtdcOldIPAddressType
	//  Mac地址
	MacAddress types.TThostFtdcMacAddressType
	//  合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
	//  合约在交易所的代码
	ExchangeInstID types.TThostFtdcExchangeInstIDType
	//  IP地址
	IPAddress types.TThostFtdcIPAddressType
}

// 期权自对冲操作
type CThostFtdcOptionSelfCloseActionField struct {
	//  经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	//  投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	//  期权自对冲操作引用
	OptionSelfCloseActionRef types.TThostFtdcOrderActionRefType
	//  期权自对冲引用
	OptionSelfCloseRef types.TThostFtdcOrderRefType
	//  请求编号
	RequestID types.TThostFtdcRequestIDType
	//  前置编号
	FrontID types.TThostFtdcFrontIDType
	//  会话编号
	SessionID types.TThostFtdcSessionIDType
	//  交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	//  期权自对冲操作编号
	OptionSelfCloseSysID types.TThostFtdcOrderSysIDType
	//  操作标志
	ActionFlag types.TThostFtdcActionFlagType
	//  操作日期
	ActionDate types.TThostFtdcDateType
	//  操作时间
	ActionTime types.TThostFtdcTimeType
	//  交易所交易员代码
	TraderID types.TThostFtdcTraderIDType
	//  安装编号
	InstallID types.TThostFtdcInstallIDType
	//  本地期权自对冲编号
	OptionSelfCloseLocalID types.TThostFtdcOrderLocalIDType
	//  操作本地编号
	ActionLocalID types.TThostFtdcOrderLocalIDType
	//  会员代码
	ParticipantID types.TThostFtdcParticipantIDType
	//  客户代码
	ClientID types.TThostFtdcClientIDType
	//  业务单元
	BusinessUnit types.TThostFtdcBusinessUnitType
	//  报单操作状态
	OrderActionStatus types.TThostFtdcOrderActionStatusType
	//  用户代码
	UserID types.TThostFtdcUserIDType
	//  状态信息
	StatusMsg types.TThostFtdcErrorMsgType
	//  保留的无效字段
	Reserve1 types.TThostFtdcOldInstrumentIDType
	//  营业部编号
	BranchID types.TThostFtdcBranchIDType
	//  投资单元代码
	InvestUnitID types.TThostFtdcInvestUnitIDType
	//  保留的无效字段
	Reserve2 types.TThostFtdcOldIPAddressType
	//  Mac地址
	MacAddress types.TThostFtdcMacAddressType
	//  合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
	//  IP地址
	IPAddress types.TThostFtdcIPAddressType
}

// 期权自对冲查询
type CThostFtdcQryOptionSelfCloseField struct {
	//  经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	//  投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	//  保留的无效字段
	Reserve1 types.TThostFtdcOldInstrumentIDType
	//  交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	//  期权自对冲编号
	OptionSelfCloseSysID types.TThostFtdcOrderSysIDType
	//  开始时间
	InsertTimeStart types.TThostFtdcTimeType
	//  结束时间
	InsertTimeEnd types.TThostFtdcTimeType
	//  合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
}

// 交易所期权自对冲信息
type CThostFtdcExchangeOptionSelfCloseField struct {
	//  数量
	Volume types.TThostFtdcVolumeType
	//  请求编号
	RequestID types.TThostFtdcRequestIDType
	//  业务单元
	BusinessUnit types.TThostFtdcBusinessUnitType
	//  投机套保标志
	HedgeFlag types.TThostFtdcHedgeFlagType
	//  期权行权的头寸是否自对冲
	OptSelfCloseFlag types.TThostFtdcOptSelfCloseFlagType
	//  本地期权自对冲编号
	OptionSelfCloseLocalID types.TThostFtdcOrderLocalIDType
	//  交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	//  会员代码
	ParticipantID types.TThostFtdcParticipantIDType
	//  客户代码
	ClientID types.TThostFtdcClientIDType
	//  保留的无效字段
	Reserve1 types.TThostFtdcOldExchangeInstIDType
	//  交易所交易员代码
	TraderID types.TThostFtdcTraderIDType
	//  安装编号
	InstallID types.TThostFtdcInstallIDType
	//  期权自对冲提交状态
	OrderSubmitStatus types.TThostFtdcOrderSubmitStatusType
	//  报单提示序号
	NotifySequence types.TThostFtdcSequenceNoType
	//  交易日
	TradingDay types.TThostFtdcDateType
	//  结算编号
	SettlementID types.TThostFtdcSettlementIDType
	//  期权自对冲编号
	OptionSelfCloseSysID types.TThostFtdcOrderSysIDType
	//  报单日期
	InsertDate types.TThostFtdcDateType
	//  插入时间
	InsertTime types.TThostFtdcTimeType
	//  撤销时间
	CancelTime types.TThostFtdcTimeType
	//  自对冲结果
	ExecResult types.TThostFtdcExecResultType
	//  结算会员编号
	ClearingPartID types.TThostFtdcParticipantIDType
	//  序号
	SequenceNo types.TThostFtdcSequenceNoType
	//  营业部编号
	BranchID types.TThostFtdcBranchIDType
	//  保留的无效字段
	Reserve2 types.TThostFtdcOldIPAddressType
	//  Mac地址
	MacAddress types.TThostFtdcMacAddressType
	//  合约在交易所的代码
	ExchangeInstID types.TThostFtdcExchangeInstIDType
	//  IP地址
	IPAddress types.TThostFtdcIPAddressType
}

// 期权自对冲操作查询
type CThostFtdcQryOptionSelfCloseActionField struct {
	//  经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	//  投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	//  交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
}

// 交易所期权自对冲操作
type CThostFtdcExchangeOptionSelfCloseActionField struct {
	//  交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	//  期权自对冲操作编号
	OptionSelfCloseSysID types.TThostFtdcOrderSysIDType
	//  操作标志
	ActionFlag types.TThostFtdcActionFlagType
	//  操作日期
	ActionDate types.TThostFtdcDateType
	//  操作时间
	ActionTime types.TThostFtdcTimeType
	//  交易所交易员代码
	TraderID types.TThostFtdcTraderIDType
	//  安装编号
	InstallID types.TThostFtdcInstallIDType
	//  本地期权自对冲编号
	OptionSelfCloseLocalID types.TThostFtdcOrderLocalIDType
	//  操作本地编号
	ActionLocalID types.TThostFtdcOrderLocalIDType
	//  会员代码
	ParticipantID types.TThostFtdcParticipantIDType
	//  客户代码
	ClientID types.TThostFtdcClientIDType
	//  业务单元
	BusinessUnit types.TThostFtdcBusinessUnitType
	//  报单操作状态
	OrderActionStatus types.TThostFtdcOrderActionStatusType
	//  用户代码
	UserID types.TThostFtdcUserIDType
	//  营业部编号
	BranchID types.TThostFtdcBranchIDType
	//  保留的无效字段
	Reserve1 types.TThostFtdcOldIPAddressType
	//  Mac地址
	MacAddress types.TThostFtdcMacAddressType
	//  保留的无效字段
	Reserve2 types.TThostFtdcOldExchangeInstIDType
	//  期权行权的头寸是否自对冲
	OptSelfCloseFlag types.TThostFtdcOptSelfCloseFlagType
	//  IP地址
	IPAddress types.TThostFtdcIPAddressType
	//  合约在交易所的代码
	ExchangeInstID types.TThostFtdcExchangeInstIDType
}

// 延时换汇同步
type CThostFtdcSyncDelaySwapField struct {
	//  换汇流水号
	DelaySwapSeqNo types.TThostFtdcDepositSeqNoType
	//  经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	//  投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	//  源币种
	FromCurrencyID types.TThostFtdcCurrencyIDType
	//  源金额
	FromAmount types.TThostFtdcMoneyType
	//  源换汇冻结金额(可用冻结)
	FromFrozenSwap types.TThostFtdcMoneyType
	//  源剩余换汇额度(可提冻结)
	FromRemainSwap types.TThostFtdcMoneyType
	//  目标币种
	ToCurrencyID types.TThostFtdcCurrencyIDType
	//  目标金额
	ToAmount types.TThostFtdcMoneyType
	//  是否手工换汇
	IsManualSwap types.TThostFtdcBoolType
	//  是否将所有外币的剩余换汇额度设置为0
	IsAllRemainSetZero types.TThostFtdcBoolType
}

// 查询延时换汇同步
type CThostFtdcQrySyncDelaySwapField struct {
	//  经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	//  延时换汇流水号
	DelaySwapSeqNo types.TThostFtdcDepositSeqNoType
}

// 投资单元
type CThostFtdcInvestUnitField struct {
	//  经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	//  投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	//  投资单元代码
	InvestUnitID types.TThostFtdcInvestUnitIDType
	//  投资者单元名称
	InvestorUnitName types.TThostFtdcPartyNameType
	//  投资者分组代码
	InvestorGroupID types.TThostFtdcInvestorIDType
	//  手续费率模板代码
	CommModelID types.TThostFtdcInvestorIDType
	//  保证金率模板代码
	MarginModelID types.TThostFtdcInvestorIDType
	//  资金账号
	AccountID types.TThostFtdcAccountIDType
	//  币种代码
	CurrencyID types.TThostFtdcCurrencyIDType
}

// 查询投资单元
type CThostFtdcQryInvestUnitField struct {
	//  经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	//  投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	//  投资单元代码
	InvestUnitID types.TThostFtdcInvestUnitIDType
}

// 二级代理商资金校验模式
type CThostFtdcSecAgentCheckModeField struct {
	//  投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	//  经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	//  币种
	CurrencyID types.TThostFtdcCurrencyIDType
	//  境外中介机构资金帐号
	BrokerSecAgentID types.TThostFtdcAccountIDType
	//  是否需要校验自己的资金账户
	CheckSelfAccount types.TThostFtdcBoolType
}

// 二级代理商信息
type CThostFtdcSecAgentTradeInfoField struct {
	//  经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	//  境外中介机构资金帐号
	BrokerSecAgentID types.TThostFtdcAccountIDType
	//  投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	//  二级代理商姓名
	LongCustomerName types.TThostFtdcLongIndividualNameType
}

// 市场行情
type CThostFtdcMarketDataField struct {
	//  交易日
	TradingDay types.TThostFtdcDateType
	//  保留的无效字段
	Reserve1 types.TThostFtdcOldInstrumentIDType
	//  交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	//  保留的无效字段
	Reserve2 types.TThostFtdcOldExchangeInstIDType
	//  最新价
	LastPrice types.TThostFtdcPriceType
	//  上次结算价
	PreSettlementPrice types.TThostFtdcPriceType
	//  昨收盘
	PreClosePrice types.TThostFtdcPriceType
	//  昨持仓量
	PreOpenInterest types.TThostFtdcLargeVolumeType
	//  今开盘
	OpenPrice types.TThostFtdcPriceType
	//  最高价
	HighestPrice types.TThostFtdcPriceType
	//  最低价
	LowestPrice types.TThostFtdcPriceType
	//  数量
	Volume types.TThostFtdcVolumeType
	//  成交金额
	Turnover types.TThostFtdcMoneyType
	//  持仓量
	OpenInterest types.TThostFtdcLargeVolumeType
	//  今收盘
	ClosePrice types.TThostFtdcPriceType
	//  本次结算价
	SettlementPrice types.TThostFtdcPriceType
	//  涨停板价
	UpperLimitPrice types.TThostFtdcPriceType
	//  跌停板价
	LowerLimitPrice types.TThostFtdcPriceType
	//  昨虚实度
	PreDelta types.TThostFtdcRatioType
	//  今虚实度
	CurrDelta types.TThostFtdcRatioType
	//  最后修改时间
	UpdateTime types.TThostFtdcTimeType
	//  最后修改毫秒
	UpdateMillisec types.TThostFtdcMillisecType
	//  业务日期
	ActionDay types.TThostFtdcDateType
	//  合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
	//  合约在交易所的代码
	ExchangeInstID types.TThostFtdcExchangeInstIDType
}

// 行情基础属性
type CThostFtdcMarketDataBaseField struct {
	//  交易日
	TradingDay types.TThostFtdcDateType
	//  上次结算价
	PreSettlementPrice types.TThostFtdcPriceType
	//  昨收盘
	PreClosePrice types.TThostFtdcPriceType
	//  昨持仓量
	PreOpenInterest types.TThostFtdcLargeVolumeType
	//  昨虚实度
	PreDelta types.TThostFtdcRatioType
}

// 行情静态属性
type CThostFtdcMarketDataStaticField struct {
	//  今开盘
	OpenPrice types.TThostFtdcPriceType
	//  最高价
	HighestPrice types.TThostFtdcPriceType
	//  最低价
	LowestPrice types.TThostFtdcPriceType
	//  今收盘
	ClosePrice types.TThostFtdcPriceType
	//  涨停板价
	UpperLimitPrice types.TThostFtdcPriceType
	//  跌停板价
	LowerLimitPrice types.TThostFtdcPriceType
	//  本次结算价
	SettlementPrice types.TThostFtdcPriceType
	//  今虚实度
	CurrDelta types.TThostFtdcRatioType
}

// 行情最新成交属性
type CThostFtdcMarketDataLastMatchField struct {
	//  最新价
	LastPrice types.TThostFtdcPriceType
	//  数量
	Volume types.TThostFtdcVolumeType
	//  成交金额
	Turnover types.TThostFtdcMoneyType
	//  持仓量
	OpenInterest types.TThostFtdcLargeVolumeType
}

// 行情最优价属性
type CThostFtdcMarketDataBestPriceField struct {
	//  申买价一
	BidPrice1 types.TThostFtdcPriceType
	//  申买量一
	BidVolume1 types.TThostFtdcVolumeType
	//  申卖价一
	AskPrice1 types.TThostFtdcPriceType
	//  申卖量一
	AskVolume1 types.TThostFtdcVolumeType
}

// 行情申买二、三属性
type CThostFtdcMarketDataBid23Field struct {
	//  申买价二
	BidPrice2 types.TThostFtdcPriceType
	//  申买量二
	BidVolume2 types.TThostFtdcVolumeType
	//  申买价三
	BidPrice3 types.TThostFtdcPriceType
	//  申买量三
	BidVolume3 types.TThostFtdcVolumeType
}

// 行情申卖二、三属性
type CThostFtdcMarketDataAsk23Field struct {
	//  申卖价二
	AskPrice2 types.TThostFtdcPriceType
	//  申卖量二
	AskVolume2 types.TThostFtdcVolumeType
	//  申卖价三
	AskPrice3 types.TThostFtdcPriceType
	//  申卖量三
	AskVolume3 types.TThostFtdcVolumeType
}

// 行情申买四、五属性
type CThostFtdcMarketDataBid45Field struct {
	//  申买价四
	BidPrice4 types.TThostFtdcPriceType
	//  申买量四
	BidVolume4 types.TThostFtdcVolumeType
	//  申买价五
	BidPrice5 types.TThostFtdcPriceType
	//  申买量五
	BidVolume5 types.TThostFtdcVolumeType
}

// 行情申卖四、五属性
type CThostFtdcMarketDataAsk45Field struct {
	//  申卖价四
	AskPrice4 types.TThostFtdcPriceType
	//  申卖量四
	AskVolume4 types.TThostFtdcVolumeType
	//  申卖价五
	AskPrice5 types.TThostFtdcPriceType
	//  申卖量五
	AskVolume5 types.TThostFtdcVolumeType
}

// 行情更新时间属性
type CThostFtdcMarketDataUpdateTimeField struct {
	//  保留的无效字段
	Reserve1 types.TThostFtdcOldInstrumentIDType
	//  最后修改时间
	UpdateTime types.TThostFtdcTimeType
	//  最后修改毫秒
	UpdateMillisec types.TThostFtdcMillisecType
	//  业务日期
	ActionDay types.TThostFtdcDateType
	//  合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
}

// 行情上下带价
type CThostFtdcMarketDataBandingPriceField struct {
	//  上带价
	BandingUpperPrice types.TThostFtdcPriceType
	//  下带价
	BandingLowerPrice types.TThostFtdcPriceType
}

// 行情交易所代码属性
type CThostFtdcMarketDataExchangeField struct {
	//  交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
}

// 指定的合约
type CThostFtdcSpecificInstrumentField struct {
	//  保留的无效字段
	Reserve1 types.TThostFtdcOldInstrumentIDType
	//  合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
}

// 合约状态
type CThostFtdcInstrumentStatusField struct {
	//  交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	//  保留的无效字段
	Reserve1 types.TThostFtdcOldExchangeInstIDType
	//  结算组代码
	SettlementGroupID types.TThostFtdcSettlementGroupIDType
	//  保留的无效字段
	Reserve2 types.TThostFtdcOldInstrumentIDType
	//  合约交易状态
	InstrumentStatus types.TThostFtdcInstrumentStatusType
	//  交易阶段编号
	TradingSegmentSN types.TThostFtdcTradingSegmentSNType
	//  进入本状态时间
	EnterTime types.TThostFtdcTimeType
	//  进入本状态原因
	EnterReason types.TThostFtdcInstStatusEnterReasonType
	//  合约在交易所的代码
	ExchangeInstID types.TThostFtdcExchangeInstIDType
	//  合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
}

// 查询合约状态
type CThostFtdcQryInstrumentStatusField struct {
	//  交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	//  保留的无效字段
	Reserve1 types.TThostFtdcOldExchangeInstIDType
	//  合约在交易所的代码
	ExchangeInstID types.TThostFtdcExchangeInstIDType
}

// 投资者账户
type CThostFtdcInvestorAccountField struct {
	//  经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	//  投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	//  投资者帐号
	AccountID types.TThostFtdcAccountIDType
	//  币种代码
	CurrencyID types.TThostFtdcCurrencyIDType
}

// 浮动盈亏算法
type CThostFtdcPositionProfitAlgorithmField struct {
	//  经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	//  投资者帐号
	AccountID types.TThostFtdcAccountIDType
	//  盈亏算法
	Algorithm types.TThostFtdcAlgorithmType
	//  备注
	Memo types.TThostFtdcMemoType
	//  币种代码
	CurrencyID types.TThostFtdcCurrencyIDType
}

// 会员资金折扣
type CThostFtdcDiscountField struct {
	//  经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	//  投资者范围
	InvestorRange types.TThostFtdcInvestorRangeType
	//  投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	//  资金折扣比例
	Discount types.TThostFtdcRatioType
}

// 查询转帐银行
type CThostFtdcQryTransferBankField struct {
	//  银行代码
	BankID types.TThostFtdcBankIDType
	//  银行分中心代码
	BankBrchID types.TThostFtdcBankBrchIDType
}

// 转帐银行
type CThostFtdcTransferBankField struct {
	//  银行代码
	BankID types.TThostFtdcBankIDType
	//  银行分中心代码
	BankBrchID types.TThostFtdcBankBrchIDType
	//  银行名称
	BankName types.TThostFtdcBankNameType
	//  是否活跃
	IsActive types.TThostFtdcBoolType
}

// 查询投资者持仓明细
type CThostFtdcQryInvestorPositionDetailField struct {
	//  经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	//  投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	//  保留的无效字段
	Reserve1 types.TThostFtdcOldInstrumentIDType
	//  交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	//  投资单元代码
	InvestUnitID types.TThostFtdcInvestUnitIDType
	//  合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
}

// 投资者持仓明细
type CThostFtdcInvestorPositionDetailField struct {
	//  保留的无效字段
	Reserve1 types.TThostFtdcOldInstrumentIDType
	//  经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	//  投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	//  投机套保标志
	HedgeFlag types.TThostFtdcHedgeFlagType
	//  买卖
	Direction types.TThostFtdcDirectionType
	//  开仓日期
	OpenDate types.TThostFtdcDateType
	//  成交编号
	TradeID types.TThostFtdcTradeIDType
	//  数量
	Volume types.TThostFtdcVolumeType
	//  开仓价
	OpenPrice types.TThostFtdcPriceType
	//  交易日
	TradingDay types.TThostFtdcDateType
	//  结算编号
	SettlementID types.TThostFtdcSettlementIDType
	//  成交类型
	TradeType types.TThostFtdcTradeTypeType
	//  保留的无效字段
	Reserve2 types.TThostFtdcOldInstrumentIDType
	//  交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	//  逐日盯市平仓盈亏
	CloseProfitByDate types.TThostFtdcMoneyType
	//  逐笔对冲平仓盈亏
	CloseProfitByTrade types.TThostFtdcMoneyType
	//  逐日盯市持仓盈亏
	PositionProfitByDate types.TThostFtdcMoneyType
	//  逐笔对冲持仓盈亏
	PositionProfitByTrade types.TThostFtdcMoneyType
	//  投资者保证金
	Margin types.TThostFtdcMoneyType
	//  交易所保证金
	ExchMargin types.TThostFtdcMoneyType
	//  保证金率
	MarginRateByMoney types.TThostFtdcRatioType
	//  保证金率(按手数)
	MarginRateByVolume types.TThostFtdcRatioType
	//  昨结算价
	LastSettlementPrice types.TThostFtdcPriceType
	//  结算价
	SettlementPrice types.TThostFtdcPriceType
	//  平仓量
	CloseVolume types.TThostFtdcVolumeType
	//  平仓金额
	CloseAmount types.TThostFtdcMoneyType
	//  先开先平剩余数量
	TimeFirstVolume types.TThostFtdcVolumeType
	//  投资单元代码
	InvestUnitID types.TThostFtdcInvestUnitIDType
	//  特殊持仓标志
	SpecPosiType types.TThostFtdcSpecPosiTypeType
	//  合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
	//  组合合约代码
	CombInstrumentID types.TThostFtdcInstrumentIDType
}

// 资金账户口令域
type CThostFtdcTradingAccountPasswordField struct {
	//  经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	//  投资者帐号
	AccountID types.TThostFtdcAccountIDType
	//  密码
	Password types.TThostFtdcPasswordType
	//  币种代码
	CurrencyID types.TThostFtdcCurrencyIDType
}

// 交易所行情报盘机
type CThostFtdcMDTraderOfferField struct {
	//  交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	//  交易所交易员代码
	TraderID types.TThostFtdcTraderIDType
	//  会员代码
	ParticipantID types.TThostFtdcParticipantIDType
	//  密码
	Password types.TThostFtdcPasswordType
	//  安装编号
	InstallID types.TThostFtdcInstallIDType
	//  本地报单编号
	OrderLocalID types.TThostFtdcOrderLocalIDType
	//  交易所交易员连接状态
	TraderConnectStatus types.TThostFtdcTraderConnectStatusType
	//  发出连接请求的日期
	ConnectRequestDate types.TThostFtdcDateType
	//  发出连接请求的时间
	ConnectRequestTime types.TThostFtdcTimeType
	//  上次报告日期
	LastReportDate types.TThostFtdcDateType
	//  上次报告时间
	LastReportTime types.TThostFtdcTimeType
	//  完成连接日期
	ConnectDate types.TThostFtdcDateType
	//  完成连接时间
	ConnectTime types.TThostFtdcTimeType
	//  启动日期
	StartDate types.TThostFtdcDateType
	//  启动时间
	StartTime types.TThostFtdcTimeType
	//  交易日
	TradingDay types.TThostFtdcDateType
	//  经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	//  本席位最大成交编号
	MaxTradeID types.TThostFtdcTradeIDType
	//  本席位最大报单备拷
	MaxOrderMessageReference types.TThostFtdcReturnCodeType
	//  撤单时选择席位算法
	OrderCancelAlg types.TThostFtdcOrderCancelAlgType
}

// 查询行情报盘机
type CThostFtdcQryMDTraderOfferField struct {
	//  交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	//  会员代码
	ParticipantID types.TThostFtdcParticipantIDType
	//  交易所交易员代码
	TraderID types.TThostFtdcTraderIDType
}

// 查询客户通知
type CThostFtdcQryNoticeField struct {
	//  经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
}

// 客户通知
type CThostFtdcNoticeField struct {
	//  经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	//  消息正文
	Content types.TThostFtdcContentType
	//  经纪公司通知内容序列号
	SequenceLabel types.TThostFtdcSequenceLabelType
}

// 用户权限
type CThostFtdcUserRightField struct {
	//  经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	//  用户代码
	UserID types.TThostFtdcUserIDType
	//  客户权限类型
	UserRightType types.TThostFtdcUserRightTypeType
	//  是否禁止
	IsForbidden types.TThostFtdcBoolType
}

// 查询结算信息确认域
type CThostFtdcQrySettlementInfoConfirmField struct {
	//  经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	//  投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	//  投资者帐号
	AccountID types.TThostFtdcAccountIDType
	//  币种代码
	CurrencyID types.TThostFtdcCurrencyIDType
}

// 装载结算信息
type CThostFtdcLoadSettlementInfoField struct {
	//  经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
}

// 经纪公司可提资金算法表
type CThostFtdcBrokerWithdrawAlgorithmField struct {
	//  经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	//  可提资金算法
	WithdrawAlgorithm types.TThostFtdcAlgorithmType
	//  资金使用率
	UsingRatio types.TThostFtdcRatioType
	//  可提是否包含平仓盈利
	IncludeCloseProfit types.TThostFtdcIncludeCloseProfitType
	//  本日无仓且无成交客户是否受可提比例限制
	AllWithoutTrade types.TThostFtdcAllWithoutTradeType
	//  可用是否包含平仓盈利
	AvailIncludeCloseProfit types.TThostFtdcIncludeCloseProfitType
	//  是否启用用户事件
	IsBrokerUserEvent types.TThostFtdcBoolType
	//  币种代码
	CurrencyID types.TThostFtdcCurrencyIDType
	//  货币质押比率
	FundMortgageRatio types.TThostFtdcRatioType
	//  权益算法
	BalanceAlgorithm types.TThostFtdcBalanceAlgorithmType
}

// 资金账户口令变更域
type CThostFtdcTradingAccountPasswordUpdateV1Field struct {
	//  经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	//  投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	//  原来的口令
	OldPassword types.TThostFtdcPasswordType
	//  新的口令
	NewPassword types.TThostFtdcPasswordType
}

// 资金账户口令变更域
type CThostFtdcTradingAccountPasswordUpdateField struct {
	//  经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	//  投资者帐号
	AccountID types.TThostFtdcAccountIDType
	//  原来的口令
	OldPassword types.TThostFtdcPasswordType
	//  新的口令
	NewPassword types.TThostFtdcPasswordType
	//  币种代码
	CurrencyID types.TThostFtdcCurrencyIDType
}

// 查询组合合约分腿
type CThostFtdcQryCombinationLegField struct {
	//  保留的无效字段
	Reserve1 types.TThostFtdcOldInstrumentIDType
	//  单腿编号
	LegID types.TThostFtdcLegIDType
	//  保留的无效字段
	Reserve2 types.TThostFtdcOldInstrumentIDType
	//  组合合约代码
	CombInstrumentID types.TThostFtdcInstrumentIDType
	//  单腿合约代码
	LegInstrumentID types.TThostFtdcInstrumentIDType
}

// 查询组合合约分腿
type CThostFtdcQrySyncStatusField struct {
	//  交易日
	TradingDay types.TThostFtdcDateType
}

// 组合交易合约的单腿
type CThostFtdcCombinationLegField struct {
	//  保留的无效字段
	Reserve1 types.TThostFtdcOldInstrumentIDType
	//  单腿编号
	LegID types.TThostFtdcLegIDType
	//  保留的无效字段
	Reserve2 types.TThostFtdcOldInstrumentIDType
	//  买卖方向
	Direction types.TThostFtdcDirectionType
	//  单腿乘数
	LegMultiple types.TThostFtdcLegMultipleType
	//  派生层数
	ImplyLevel types.TThostFtdcImplyLevelType
	//  组合合约代码
	CombInstrumentID types.TThostFtdcInstrumentIDType
	//  单腿合约代码
	LegInstrumentID types.TThostFtdcInstrumentIDType
}

// 数据同步状态
type CThostFtdcSyncStatusField struct {
	//  交易日
	TradingDay types.TThostFtdcDateType
	//  数据同步状态
	DataSyncStatus types.TThostFtdcDataSyncStatusType
}

// 查询联系人
type CThostFtdcQryLinkManField struct {
	//  经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	//  投资者代码
	InvestorID types.TThostFtdcInvestorIDType
}

// 联系人
type CThostFtdcLinkManField struct {
	//  经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	//  投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	//  联系人类型
	PersonType types.TThostFtdcPersonTypeType
	//  证件类型
	IdentifiedCardType types.TThostFtdcIdCardTypeType
	//  证件号码
	IdentifiedCardNo types.TThostFtdcIdentifiedCardNoType
	//  名称
	PersonName types.TThostFtdcPartyNameType
	//  联系电话
	Telephone types.TThostFtdcTelephoneType
	//  通讯地址
	Address types.TThostFtdcAddressType
	//  邮政编码
	ZipCode types.TThostFtdcZipCodeType
	//  优先级
	Priority types.TThostFtdcPriorityType
	//  开户邮政编码
	UOAZipCode types.TThostFtdcUOAZipCodeType
	//  全称
	PersonFullName types.TThostFtdcInvestorFullNameType
}

// 查询经纪公司用户事件
type CThostFtdcQryBrokerUserEventField struct {
	//  经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	//  用户代码
	UserID types.TThostFtdcUserIDType
	//  用户事件类型
	UserEventType types.TThostFtdcUserEventTypeType
}

// 查询经纪公司用户事件
type CThostFtdcBrokerUserEventField struct {
	//  经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	//  用户代码
	UserID types.TThostFtdcUserIDType
	//  用户事件类型
	UserEventType types.TThostFtdcUserEventTypeType
	//  用户事件序号
	EventSequenceNo types.TThostFtdcSequenceNoType
	//  事件发生日期
	EventDate types.TThostFtdcDateType
	//  事件发生时间
	EventTime types.TThostFtdcTimeType
	//  用户事件信息
	UserEventInfo types.TThostFtdcUserEventInfoType
	//  投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	//  保留的无效字段
	Reserve1 types.TThostFtdcOldInstrumentIDType
	//  合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
	//  交易中心代码
	DRIdentityID types.TThostFtdcDRIdentityIDType
	//  交易日
	TradingDay types.TThostFtdcDateType
}

// 查询签约银行请求
type CThostFtdcQryContractBankField struct {
	//  经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	//  银行代码
	BankID types.TThostFtdcBankIDType
	//  银行分中心代码
	BankBrchID types.TThostFtdcBankBrchIDType
}

// 查询签约银行响应
type CThostFtdcContractBankField struct {
	//  经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	//  银行代码
	BankID types.TThostFtdcBankIDType
	//  银行分中心代码
	BankBrchID types.TThostFtdcBankBrchIDType
	//  银行名称
	BankName types.TThostFtdcBankNameType
	//  上报csrc的银行代码
	CsrcBankID types.TThostFtdcBankIDType
}

// 投资者组合持仓明细
type CThostFtdcInvestorPositionCombineDetailField struct {
	//  交易日
	TradingDay types.TThostFtdcDateType
	//  开仓日期
	OpenDate types.TThostFtdcDateType
	//  交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	//  结算编号
	SettlementID types.TThostFtdcSettlementIDType
	//  经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	//  投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	//  组合编号
	ComTradeID types.TThostFtdcTradeIDType
	//  撮合编号
	TradeID types.TThostFtdcTradeIDType
	//  保留的无效字段
	Reserve1 types.TThostFtdcOldInstrumentIDType
	//  投机套保标志
	HedgeFlag types.TThostFtdcHedgeFlagType
	//  买卖
	Direction types.TThostFtdcDirectionType
	//  持仓量
	TotalAmt types.TThostFtdcVolumeType
	//  投资者保证金
	Margin types.TThostFtdcMoneyType
	//  交易所保证金
	ExchMargin types.TThostFtdcMoneyType
	//  保证金率
	MarginRateByMoney types.TThostFtdcRatioType
	//  保证金率(按手数)
	MarginRateByVolume types.TThostFtdcRatioType
	//  单腿编号
	LegID types.TThostFtdcLegIDType
	//  单腿乘数
	LegMultiple types.TThostFtdcLegMultipleType
	//  保留的无效字段
	Reserve2 types.TThostFtdcOldInstrumentIDType
	//  成交组号
	TradeGroupID types.TThostFtdcTradeGroupIDType
	//  投资单元代码
	InvestUnitID types.TThostFtdcInvestUnitIDType
	//  合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
	//  组合持仓合约编码
	CombInstrumentID types.TThostFtdcInstrumentIDType
}

// 预埋单
type CThostFtdcParkedOrderField struct {
	//  经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	//  投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	//  保留的无效字段
	Reserve1 types.TThostFtdcOldInstrumentIDType
	//  报单引用
	OrderRef types.TThostFtdcOrderRefType
	//  用户代码
	UserID types.TThostFtdcUserIDType
	//  报单价格条件
	OrderPriceType types.TThostFtdcOrderPriceTypeType
	//  买卖方向
	Direction types.TThostFtdcDirectionType
	//  组合开平标志
	CombOffsetFlag types.TThostFtdcCombOffsetFlagType
	//  组合投机套保标志
	CombHedgeFlag types.TThostFtdcCombHedgeFlagType
	//  价格
	LimitPrice types.TThostFtdcPriceType
	//  数量
	VolumeTotalOriginal types.TThostFtdcVolumeType
	//  有效期类型
	TimeCondition types.TThostFtdcTimeConditionType
	//  GTD日期
	GTDDate types.TThostFtdcDateType
	//  成交量类型
	VolumeCondition types.TThostFtdcVolumeConditionType
	//  最小成交量
	MinVolume types.TThostFtdcVolumeType
	//  触发条件
	ContingentCondition types.TThostFtdcContingentConditionType
	//  止损价
	StopPrice types.TThostFtdcPriceType
	//  强平原因
	ForceCloseReason types.TThostFtdcForceCloseReasonType
	//  自动挂起标志
	IsAutoSuspend types.TThostFtdcBoolType
	//  业务单元
	BusinessUnit types.TThostFtdcBusinessUnitType
	//  请求编号
	RequestID types.TThostFtdcRequestIDType
	//  用户强平标志
	UserForceClose types.TThostFtdcBoolType
	//  交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	//  预埋报单编号
	ParkedOrderID types.TThostFtdcParkedOrderIDType
	//  用户类型
	UserType types.TThostFtdcUserTypeType
	//  预埋单状态
	Status types.TThostFtdcParkedOrderStatusType
	//  错误代码
	ErrorID types.TThostFtdcErrorIDType
	//  错误信息
	ErrorMsg types.TThostFtdcErrorMsgType
	//  互换单标志
	IsSwapOrder types.TThostFtdcBoolType
	//  资金账号
	AccountID types.TThostFtdcAccountIDType
	//  币种代码
	CurrencyID types.TThostFtdcCurrencyIDType
	//  交易编码
	ClientID types.TThostFtdcClientIDType
	//  投资单元代码
	InvestUnitID types.TThostFtdcInvestUnitIDType
	//  保留的无效字段
	Reserve2 types.TThostFtdcOldIPAddressType
	//  Mac地址
	MacAddress types.TThostFtdcMacAddressType
	//  合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
	//  IP地址
	IPAddress types.TThostFtdcIPAddressType
}

// 输入预埋单操作
type CThostFtdcParkedOrderActionField struct {
	//  经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	//  投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	//  报单操作引用
	OrderActionRef types.TThostFtdcOrderActionRefType
	//  报单引用
	OrderRef types.TThostFtdcOrderRefType
	//  请求编号
	RequestID types.TThostFtdcRequestIDType
	//  前置编号
	FrontID types.TThostFtdcFrontIDType
	//  会话编号
	SessionID types.TThostFtdcSessionIDType
	//  交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	//  报单编号
	OrderSysID types.TThostFtdcOrderSysIDType
	//  操作标志
	ActionFlag types.TThostFtdcActionFlagType
	//  价格
	LimitPrice types.TThostFtdcPriceType
	//  数量变化
	VolumeChange types.TThostFtdcVolumeType
	//  用户代码
	UserID types.TThostFtdcUserIDType
	//  保留的无效字段
	Reserve1 types.TThostFtdcOldInstrumentIDType
	//  预埋撤单单编号
	ParkedOrderActionID types.TThostFtdcParkedOrderActionIDType
	//  用户类型
	UserType types.TThostFtdcUserTypeType
	//  预埋撤单状态
	Status types.TThostFtdcParkedOrderStatusType
	//  错误代码
	ErrorID types.TThostFtdcErrorIDType
	//  错误信息
	ErrorMsg types.TThostFtdcErrorMsgType
	//  投资单元代码
	InvestUnitID types.TThostFtdcInvestUnitIDType
	//  保留的无效字段
	Reserve2 types.TThostFtdcOldIPAddressType
	//  Mac地址
	MacAddress types.TThostFtdcMacAddressType
	//  合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
	//  IP地址
	IPAddress types.TThostFtdcIPAddressType
}

// 查询预埋单
type CThostFtdcQryParkedOrderField struct {
	//  经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	//  投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	//  保留的无效字段
	Reserve1 types.TThostFtdcOldInstrumentIDType
	//  交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	//  投资单元代码
	InvestUnitID types.TThostFtdcInvestUnitIDType
	//  合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
}

// 查询预埋撤单
type CThostFtdcQryParkedOrderActionField struct {
	//  经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	//  投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	//  保留的无效字段
	Reserve1 types.TThostFtdcOldInstrumentIDType
	//  交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	//  投资单元代码
	InvestUnitID types.TThostFtdcInvestUnitIDType
	//  合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
}

// 删除预埋单
type CThostFtdcRemoveParkedOrderField struct {
	//  经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	//  投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	//  预埋报单编号
	ParkedOrderID types.TThostFtdcParkedOrderIDType
	//  投资单元代码
	InvestUnitID types.TThostFtdcInvestUnitIDType
}

// 删除预埋撤单
type CThostFtdcRemoveParkedOrderActionField struct {
	//  经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	//  投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	//  预埋撤单编号
	ParkedOrderActionID types.TThostFtdcParkedOrderActionIDType
	//  投资单元代码
	InvestUnitID types.TThostFtdcInvestUnitIDType
}

// 经纪公司可提资金算法表
type CThostFtdcInvestorWithdrawAlgorithmField struct {
	//  经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	//  投资者范围
	InvestorRange types.TThostFtdcInvestorRangeType
	//  投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	//  可提资金比例
	UsingRatio types.TThostFtdcRatioType
	//  币种代码
	CurrencyID types.TThostFtdcCurrencyIDType
	//  货币质押比率
	FundMortgageRatio types.TThostFtdcRatioType
}

// 查询组合持仓明细
type CThostFtdcQryInvestorPositionCombineDetailField struct {
	//  经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	//  投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	//  保留的无效字段
	Reserve1 types.TThostFtdcOldInstrumentIDType
	//  交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	//  投资单元代码
	InvestUnitID types.TThostFtdcInvestUnitIDType
	//  组合持仓合约编码
	CombInstrumentID types.TThostFtdcInstrumentIDType
}

// 成交均价
type CThostFtdcMarketDataAveragePriceField struct {
	//  当日均价
	AveragePrice types.TThostFtdcPriceType
}

// 校验投资者密码
type CThostFtdcVerifyInvestorPasswordField struct {
	//  经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	//  投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	//  密码
	Password types.TThostFtdcPasswordType
}

// 用户IP
type CThostFtdcUserIPField struct {
	//  经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	//  用户代码
	UserID types.TThostFtdcUserIDType
	//  保留的无效字段
	Reserve1 types.TThostFtdcOldIPAddressType
	//  保留的无效字段
	Reserve2 types.TThostFtdcOldIPAddressType
	//  Mac地址
	MacAddress types.TThostFtdcMacAddressType
	//  IP地址
	IPAddress types.TThostFtdcIPAddressType
	//  IP地址掩码
	IPMask types.TThostFtdcIPAddressType
}

// 用户事件通知信息
type CThostFtdcTradingNoticeInfoField struct {
	//  经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	//  投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	//  发送时间
	SendTime types.TThostFtdcTimeType
	//  消息正文
	FieldContent types.TThostFtdcContentType
	//  序列系列号
	SequenceSeries types.TThostFtdcSequenceSeriesType
	//  序列号
	SequenceNo types.TThostFtdcSequenceNoType
	//  投资单元代码
	InvestUnitID types.TThostFtdcInvestUnitIDType
}

// 用户事件通知
type CThostFtdcTradingNoticeField struct {
	//  经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	//  投资者范围
	InvestorRange types.TThostFtdcInvestorRangeType
	//  投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	//  序列系列号
	SequenceSeries types.TThostFtdcSequenceSeriesType
	//  用户代码
	UserID types.TThostFtdcUserIDType
	//  发送时间
	SendTime types.TThostFtdcTimeType
	//  序列号
	SequenceNo types.TThostFtdcSequenceNoType
	//  消息正文
	FieldContent types.TThostFtdcContentType
	//  投资单元代码
	InvestUnitID types.TThostFtdcInvestUnitIDType
}

// 查询交易事件通知
type CThostFtdcQryTradingNoticeField struct {
	//  经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	//  投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	//  投资单元代码
	InvestUnitID types.TThostFtdcInvestUnitIDType
}

// 查询错误报单
type CThostFtdcQryErrOrderField struct {
	//  经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	//  投资者代码
	InvestorID types.TThostFtdcInvestorIDType
}

// 错误报单
type CThostFtdcErrOrderField struct {
	//  经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	//  投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	//  保留的无效字段
	Reserve1 types.TThostFtdcOldInstrumentIDType
	//  报单引用
	OrderRef types.TThostFtdcOrderRefType
	//  用户代码
	UserID types.TThostFtdcUserIDType
	//  报单价格条件
	OrderPriceType types.TThostFtdcOrderPriceTypeType
	//  买卖方向
	Direction types.TThostFtdcDirectionType
	//  组合开平标志
	CombOffsetFlag types.TThostFtdcCombOffsetFlagType
	//  组合投机套保标志
	CombHedgeFlag types.TThostFtdcCombHedgeFlagType
	//  价格
	LimitPrice types.TThostFtdcPriceType
	//  数量
	VolumeTotalOriginal types.TThostFtdcVolumeType
	//  有效期类型
	TimeCondition types.TThostFtdcTimeConditionType
	//  GTD日期
	GTDDate types.TThostFtdcDateType
	//  成交量类型
	VolumeCondition types.TThostFtdcVolumeConditionType
	//  最小成交量
	MinVolume types.TThostFtdcVolumeType
	//  触发条件
	ContingentCondition types.TThostFtdcContingentConditionType
	//  止损价
	StopPrice types.TThostFtdcPriceType
	//  强平原因
	ForceCloseReason types.TThostFtdcForceCloseReasonType
	//  自动挂起标志
	IsAutoSuspend types.TThostFtdcBoolType
	//  业务单元
	BusinessUnit types.TThostFtdcBusinessUnitType
	//  请求编号
	RequestID types.TThostFtdcRequestIDType
	//  用户强平标志
	UserForceClose types.TThostFtdcBoolType
	//  错误代码
	ErrorID types.TThostFtdcErrorIDType
	//  错误信息
	ErrorMsg types.TThostFtdcErrorMsgType
	//  互换单标志
	IsSwapOrder types.TThostFtdcBoolType
	//  交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	//  投资单元代码
	InvestUnitID types.TThostFtdcInvestUnitIDType
	//  资金账号
	AccountID types.TThostFtdcAccountIDType
	//  币种代码
	CurrencyID types.TThostFtdcCurrencyIDType
	//  交易编码
	ClientID types.TThostFtdcClientIDType
	//  保留的无效字段
	Reserve2 types.TThostFtdcOldIPAddressType
	//  Mac地址
	MacAddress types.TThostFtdcMacAddressType
	//  合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
	//  IP地址
	IPAddress types.TThostFtdcIPAddressType
	//  报单回显字段
	OrderMemo types.TThostFtdcOrderMemoType
	//  session上请求计数 api自动维护
	SessionReqSeq types.TThostFtdcSequenceNo12Type
}

// 查询错误报单操作
type CThostFtdcErrorConditionalOrderField struct {
	//  经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	//  投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	//  保留的无效字段
	Reserve1 types.TThostFtdcOldInstrumentIDType
	//  报单引用
	OrderRef types.TThostFtdcOrderRefType
	//  用户代码
	UserID types.TThostFtdcUserIDType
	//  报单价格条件
	OrderPriceType types.TThostFtdcOrderPriceTypeType
	//  买卖方向
	Direction types.TThostFtdcDirectionType
	//  组合开平标志
	CombOffsetFlag types.TThostFtdcCombOffsetFlagType
	//  组合投机套保标志
	CombHedgeFlag types.TThostFtdcCombHedgeFlagType
	//  价格
	LimitPrice types.TThostFtdcPriceType
	//  数量
	VolumeTotalOriginal types.TThostFtdcVolumeType
	//  有效期类型
	TimeCondition types.TThostFtdcTimeConditionType
	//  GTD日期
	GTDDate types.TThostFtdcDateType
	//  成交量类型
	VolumeCondition types.TThostFtdcVolumeConditionType
	//  最小成交量
	MinVolume types.TThostFtdcVolumeType
	//  触发条件
	ContingentCondition types.TThostFtdcContingentConditionType
	//  止损价
	StopPrice types.TThostFtdcPriceType
	//  强平原因
	ForceCloseReason types.TThostFtdcForceCloseReasonType
	//  自动挂起标志
	IsAutoSuspend types.TThostFtdcBoolType
	//  业务单元
	BusinessUnit types.TThostFtdcBusinessUnitType
	//  请求编号
	RequestID types.TThostFtdcRequestIDType
	//  本地报单编号
	OrderLocalID types.TThostFtdcOrderLocalIDType
	//  交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	//  会员代码
	ParticipantID types.TThostFtdcParticipantIDType
	//  客户代码
	ClientID types.TThostFtdcClientIDType
	//  保留的无效字段
	Reserve2 types.TThostFtdcOldExchangeInstIDType
	//  交易所交易员代码
	TraderID types.TThostFtdcTraderIDType
	//  安装编号
	InstallID types.TThostFtdcInstallIDType
	//  报单提交状态
	OrderSubmitStatus types.TThostFtdcOrderSubmitStatusType
	//  报单提示序号
	NotifySequence types.TThostFtdcSequenceNoType
	//  交易日
	TradingDay types.TThostFtdcDateType
	//  结算编号
	SettlementID types.TThostFtdcSettlementIDType
	//  报单编号
	OrderSysID types.TThostFtdcOrderSysIDType
	//  报单来源
	OrderSource types.TThostFtdcOrderSourceType
	//  报单状态
	OrderStatus types.TThostFtdcOrderStatusType
	//  报单类型
	OrderType types.TThostFtdcOrderTypeType
	//  今成交数量
	VolumeTraded types.TThostFtdcVolumeType
	//  剩余数量
	VolumeTotal types.TThostFtdcVolumeType
	//  报单日期
	InsertDate types.TThostFtdcDateType
	//  委托时间
	InsertTime types.TThostFtdcTimeType
	//  激活时间
	ActiveTime types.TThostFtdcTimeType
	//  挂起时间
	SuspendTime types.TThostFtdcTimeType
	//  最后修改时间
	UpdateTime types.TThostFtdcTimeType
	//  撤销时间
	CancelTime types.TThostFtdcTimeType
	//  最后修改交易所交易员代码
	ActiveTraderID types.TThostFtdcTraderIDType
	//  结算会员编号
	ClearingPartID types.TThostFtdcParticipantIDType
	//  序号
	SequenceNo types.TThostFtdcSequenceNoType
	//  前置编号
	FrontID types.TThostFtdcFrontIDType
	//  会话编号
	SessionID types.TThostFtdcSessionIDType
	//  用户端产品信息
	UserProductInfo types.TThostFtdcProductInfoType
	//  状态信息
	StatusMsg types.TThostFtdcErrorMsgType
	//  用户强平标志
	UserForceClose types.TThostFtdcBoolType
	//  操作用户代码
	ActiveUserID types.TThostFtdcUserIDType
	//  经纪公司报单编号
	BrokerOrderSeq types.TThostFtdcSequenceNoType
	//  相关报单
	RelativeOrderSysID types.TThostFtdcOrderSysIDType
	//  郑商所成交数量
	ZCETotalTradedVolume types.TThostFtdcVolumeType
	//  错误代码
	ErrorID types.TThostFtdcErrorIDType
	//  错误信息
	ErrorMsg types.TThostFtdcErrorMsgType
	//  互换单标志
	IsSwapOrder types.TThostFtdcBoolType
	//  营业部编号
	BranchID types.TThostFtdcBranchIDType
	//  投资单元代码
	InvestUnitID types.TThostFtdcInvestUnitIDType
	//  资金账号
	AccountID types.TThostFtdcAccountIDType
	//  币种代码
	CurrencyID types.TThostFtdcCurrencyIDType
	//  保留的无效字段
	Reserve3 types.TThostFtdcOldIPAddressType
	//  Mac地址
	MacAddress types.TThostFtdcMacAddressType
	//  合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
	//  合约在交易所的代码
	ExchangeInstID types.TThostFtdcExchangeInstIDType
	//  IP地址
	IPAddress types.TThostFtdcIPAddressType
}

// 查询错误报单操作
type CThostFtdcQryErrOrderActionField struct {
	//  经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	//  投资者代码
	InvestorID types.TThostFtdcInvestorIDType
}

// 错误报单操作
type CThostFtdcErrOrderActionField struct {
	//  经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	//  投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	//  报单操作引用
	OrderActionRef types.TThostFtdcOrderActionRefType
	//  报单引用
	OrderRef types.TThostFtdcOrderRefType
	//  请求编号
	RequestID types.TThostFtdcRequestIDType
	//  前置编号
	FrontID types.TThostFtdcFrontIDType
	//  会话编号
	SessionID types.TThostFtdcSessionIDType
	//  交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	//  报单编号
	OrderSysID types.TThostFtdcOrderSysIDType
	//  操作标志
	ActionFlag types.TThostFtdcActionFlagType
	//  价格
	LimitPrice types.TThostFtdcPriceType
	//  数量变化
	VolumeChange types.TThostFtdcVolumeType
	//  操作日期
	ActionDate types.TThostFtdcDateType
	//  操作时间
	ActionTime types.TThostFtdcTimeType
	//  交易所交易员代码
	TraderID types.TThostFtdcTraderIDType
	//  安装编号
	InstallID types.TThostFtdcInstallIDType
	//  本地报单编号
	OrderLocalID types.TThostFtdcOrderLocalIDType
	//  操作本地编号
	ActionLocalID types.TThostFtdcOrderLocalIDType
	//  会员代码
	ParticipantID types.TThostFtdcParticipantIDType
	//  客户代码
	ClientID types.TThostFtdcClientIDType
	//  业务单元
	BusinessUnit types.TThostFtdcBusinessUnitType
	//  报单操作状态
	OrderActionStatus types.TThostFtdcOrderActionStatusType
	//  用户代码
	UserID types.TThostFtdcUserIDType
	//  状态信息
	StatusMsg types.TThostFtdcErrorMsgType
	//  保留的无效字段
	Reserve1 types.TThostFtdcOldInstrumentIDType
	//  营业部编号
	BranchID types.TThostFtdcBranchIDType
	//  投资单元代码
	InvestUnitID types.TThostFtdcInvestUnitIDType
	//  保留的无效字段
	Reserve2 types.TThostFtdcOldIPAddressType
	//  Mac地址
	MacAddress types.TThostFtdcMacAddressType
	//  错误代码
	ErrorID types.TThostFtdcErrorIDType
	//  错误信息
	ErrorMsg types.TThostFtdcErrorMsgType
	//  合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
	//  IP地址
	IPAddress types.TThostFtdcIPAddressType
	//  报单回显字段
	OrderMemo types.TThostFtdcOrderMemoType
	//  session上请求计数 api自动维护
	SessionReqSeq types.TThostFtdcSequenceNo12Type
}

// 查询交易所状态
type CThostFtdcQryExchangeSequenceField struct {
	//  交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
}

// 交易所状态
type CThostFtdcExchangeSequenceField struct {
	//  交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	//  序号
	SequenceNo types.TThostFtdcSequenceNoType
	//  合约交易状态
	MarketStatus types.TThostFtdcInstrumentStatusType
}

// 根据价格查询最大报单数量
type CThostFtdcQryMaxOrderVolumeWithPriceField struct {
	//  经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	//  投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	//  保留的无效字段
	Reserve1 types.TThostFtdcOldInstrumentIDType
	//  买卖方向
	Direction types.TThostFtdcDirectionType
	//  开平标志
	OffsetFlag types.TThostFtdcOffsetFlagType
	//  投机套保标志
	HedgeFlag types.TThostFtdcHedgeFlagType
	//  最大允许报单数量
	MaxVolume types.TThostFtdcVolumeType
	//  报单价格
	Price types.TThostFtdcPriceType
	//  交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	//  投资单元代码
	InvestUnitID types.TThostFtdcInvestUnitIDType
	//  合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
}

// 查询经纪公司交易参数
type CThostFtdcQryBrokerTradingParamsField struct {
	//  经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	//  投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	//  币种代码
	CurrencyID types.TThostFtdcCurrencyIDType
	//  投资者帐号
	AccountID types.TThostFtdcAccountIDType
}

// 经纪公司交易参数
type CThostFtdcBrokerTradingParamsField struct {
	//  经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	//  投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	//  保证金价格类型
	MarginPriceType types.TThostFtdcMarginPriceTypeType
	//  盈亏算法
	Algorithm types.TThostFtdcAlgorithmType
	//  可用是否包含平仓盈利
	AvailIncludeCloseProfit types.TThostFtdcIncludeCloseProfitType
	//  币种代码
	CurrencyID types.TThostFtdcCurrencyIDType
	//  期权权利金价格类型
	OptionRoyaltyPriceType types.TThostFtdcOptionRoyaltyPriceTypeType
	//  投资者帐号
	AccountID types.TThostFtdcAccountIDType
}

// 查询经纪公司交易算法
type CThostFtdcQryBrokerTradingAlgosField struct {
	//  经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	//  交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	//  保留的无效字段
	Reserve1 types.TThostFtdcOldInstrumentIDType
	//  合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
}

// 经纪公司交易算法
type CThostFtdcBrokerTradingAlgosField struct {
	//  经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	//  交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	//  保留的无效字段
	Reserve1 types.TThostFtdcOldInstrumentIDType
	//  持仓处理算法编号
	HandlePositionAlgoID types.TThostFtdcHandlePositionAlgoIDType
	//  寻找保证金率算法编号
	FindMarginRateAlgoID types.TThostFtdcFindMarginRateAlgoIDType
	//  资金处理算法编号
	HandleTradingAccountAlgoID types.TThostFtdcHandleTradingAccountAlgoIDType
	//  合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
}

// 查询经纪公司资金
type CThostFtdcQueryBrokerDepositField struct {
	//  经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	//  交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
}

// 经纪公司资金
type CThostFtdcBrokerDepositField struct {
	//  交易日期
	TradingDay types.TThostFtdcTradeDateType
	//  经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	//  会员代码
	ParticipantID types.TThostFtdcParticipantIDType
	//  交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	//  上次结算准备金
	PreBalance types.TThostFtdcMoneyType
	//  当前保证金总额
	CurrMargin types.TThostFtdcMoneyType
	//  平仓盈亏
	CloseProfit types.TThostFtdcMoneyType
	//  期货结算准备金
	Balance types.TThostFtdcMoneyType
	//  入金金额
	Deposit types.TThostFtdcMoneyType
	//  出金金额
	Withdraw types.TThostFtdcMoneyType
	//  可提资金
	Available types.TThostFtdcMoneyType
	//  基本准备金
	Reserve types.TThostFtdcMoneyType
	//  冻结的保证金
	FrozenMargin types.TThostFtdcMoneyType
}

// 查询保证金监管系统经纪公司密钥
type CThostFtdcQryCFMMCBrokerKeyField struct {
	//  经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
}

// 保证金监管系统经纪公司密钥
type CThostFtdcCFMMCBrokerKeyField struct {
	//  经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	//  经纪公司统一编码
	ParticipantID types.TThostFtdcParticipantIDType
	//  密钥生成日期
	CreateDate types.TThostFtdcDateType
	//  密钥生成时间
	CreateTime types.TThostFtdcTimeType
	//  密钥编号
	KeyID types.TThostFtdcSequenceNoType
	//  动态密钥
	CurrentKey types.TThostFtdcCFMMCKeyType
	//  动态密钥类型
	KeyKind types.TThostFtdcCFMMCKeyKindType
}

// 保证金监管系统经纪公司资金账户密钥
type CThostFtdcCFMMCTradingAccountKeyField struct {
	//  经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	//  经纪公司统一编码
	ParticipantID types.TThostFtdcParticipantIDType
	//  投资者帐号
	AccountID types.TThostFtdcAccountIDType
	//  密钥编号
	KeyID types.TThostFtdcSequenceNoType
	//  动态密钥
	CurrentKey types.TThostFtdcCFMMCKeyType
}

// 请求查询保证金监管系统经纪公司资金账户密钥
type CThostFtdcQryCFMMCTradingAccountKeyField struct {
	//  经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	//  投资者代码
	InvestorID types.TThostFtdcInvestorIDType
}

// 用户动态令牌参数
type CThostFtdcBrokerUserOTPParamField struct {
	//  经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	//  用户代码
	UserID types.TThostFtdcUserIDType
	//  动态令牌提供商
	OTPVendorsID types.TThostFtdcOTPVendorsIDType
	//  动态令牌序列号
	SerialNumber types.TThostFtdcSerialNumberType
	//  令牌密钥
	AuthKey types.TThostFtdcAuthKeyType
	//  漂移值
	LastDrift types.TThostFtdcLastDriftType
	//  成功值
	LastSuccess types.TThostFtdcLastSuccessType
	//  动态令牌类型
	OTPType types.TThostFtdcOTPTypeType
}

// 手工同步用户动态令牌
type CThostFtdcManualSyncBrokerUserOTPField struct {
	//  经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	//  用户代码
	UserID types.TThostFtdcUserIDType
	//  动态令牌类型
	OTPType types.TThostFtdcOTPTypeType
	//  第一个动态密码
	FirstOTP types.TThostFtdcPasswordType
	//  第二个动态密码
	SecondOTP types.TThostFtdcPasswordType
}

// 投资者手续费率模板
type CThostFtdcCommRateModelField struct {
	//  经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	//  手续费率模板代码
	CommModelID types.TThostFtdcInvestorIDType
	//  模板名称
	CommModelName types.TThostFtdcCommModelNameType
}

// 请求查询投资者手续费率模板
type CThostFtdcQryCommRateModelField struct {
	//  经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	//  手续费率模板代码
	CommModelID types.TThostFtdcInvestorIDType
}

// 投资者保证金率模板
type CThostFtdcMarginModelField struct {
	//  经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	//  保证金率模板代码
	MarginModelID types.TThostFtdcInvestorIDType
	//  模板名称
	MarginModelName types.TThostFtdcCommModelNameType
}

// 请求查询投资者保证金率模板
type CThostFtdcQryMarginModelField struct {
	//  经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	//  保证金率模板代码
	MarginModelID types.TThostFtdcInvestorIDType
}

// 仓单折抵信息
type CThostFtdcEWarrantOffsetField struct {
	//  交易日期
	TradingDay types.TThostFtdcTradeDateType
	//  经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	//  投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	//  交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	//  保留的无效字段
	Reserve1 types.TThostFtdcOldInstrumentIDType
	//  买卖方向
	Direction types.TThostFtdcDirectionType
	//  投机套保标志
	HedgeFlag types.TThostFtdcHedgeFlagType
	//  数量
	Volume types.TThostFtdcVolumeType
	//  投资单元代码
	InvestUnitID types.TThostFtdcInvestUnitIDType
	//  合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
}

// 查询仓单折抵信息
type CThostFtdcQryEWarrantOffsetField struct {
	//  经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	//  投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	//  交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	//  保留的无效字段
	Reserve1 types.TThostFtdcOldInstrumentIDType
	//  投资单元代码
	InvestUnitID types.TThostFtdcInvestUnitIDType
	//  合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
}

// 查询投资者品种/跨品种保证金
type CThostFtdcQryInvestorProductGroupMarginField struct {
	//  经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	//  投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	//  保留的无效字段
	Reserve1 types.TThostFtdcOldInstrumentIDType
	//  投机套保标志
	HedgeFlag types.TThostFtdcHedgeFlagType
	//  交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	//  投资单元代码
	InvestUnitID types.TThostFtdcInvestUnitIDType
	//  品种/跨品种标示
	ProductGroupID types.TThostFtdcInstrumentIDType
}

// 投资者品种/跨品种保证金
type CThostFtdcInvestorProductGroupMarginField struct {
	//  保留的无效字段
	Reserve1 types.TThostFtdcOldInstrumentIDType
	//  经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	//  投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	//  交易日
	TradingDay types.TThostFtdcDateType
	//  结算编号
	SettlementID types.TThostFtdcSettlementIDType
	//  冻结的保证金
	FrozenMargin types.TThostFtdcMoneyType
	//  多头冻结的保证金
	LongFrozenMargin types.TThostFtdcMoneyType
	//  空头冻结的保证金
	ShortFrozenMargin types.TThostFtdcMoneyType
	//  占用的保证金
	UseMargin types.TThostFtdcMoneyType
	//  多头保证金
	LongUseMargin types.TThostFtdcMoneyType
	//  空头保证金
	ShortUseMargin types.TThostFtdcMoneyType
	//  交易所保证金
	ExchMargin types.TThostFtdcMoneyType
	//  交易所多头保证金
	LongExchMargin types.TThostFtdcMoneyType
	//  交易所空头保证金
	ShortExchMargin types.TThostFtdcMoneyType
	//  平仓盈亏
	CloseProfit types.TThostFtdcMoneyType
	//  冻结的手续费
	FrozenCommission types.TThostFtdcMoneyType
	//  手续费
	Commission types.TThostFtdcMoneyType
	//  冻结的资金
	FrozenCash types.TThostFtdcMoneyType
	//  资金差额
	CashIn types.TThostFtdcMoneyType
	//  持仓盈亏
	PositionProfit types.TThostFtdcMoneyType
	//  折抵总金额
	OffsetAmount types.TThostFtdcMoneyType
	//  多头折抵总金额
	LongOffsetAmount types.TThostFtdcMoneyType
	//  空头折抵总金额
	ShortOffsetAmount types.TThostFtdcMoneyType
	//  交易所折抵总金额
	ExchOffsetAmount types.TThostFtdcMoneyType
	//  交易所多头折抵总金额
	LongExchOffsetAmount types.TThostFtdcMoneyType
	//  交易所空头折抵总金额
	ShortExchOffsetAmount types.TThostFtdcMoneyType
	//  投机套保标志
	HedgeFlag types.TThostFtdcHedgeFlagType
	//  交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	//  投资单元代码
	InvestUnitID types.TThostFtdcInvestUnitIDType
	//  品种/跨品种标示
	ProductGroupID types.TThostFtdcInstrumentIDType
}

// 查询监控中心用户令牌
type CThostFtdcQueryCFMMCTradingAccountTokenField struct {
	//  经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	//  投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	//  投资单元代码
	InvestUnitID types.TThostFtdcInvestUnitIDType
}

// 监控中心用户令牌
type CThostFtdcCFMMCTradingAccountTokenField struct {
	//  经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	//  经纪公司统一编码
	ParticipantID types.TThostFtdcParticipantIDType
	//  投资者帐号
	AccountID types.TThostFtdcAccountIDType
	//  密钥编号
	KeyID types.TThostFtdcSequenceNoType
	//  动态令牌
	Token types.TThostFtdcCFMMCTokenType
}

// 查询产品组
type CThostFtdcQryProductGroupField struct {
	//  保留的无效字段
	Reserve1 types.TThostFtdcOldInstrumentIDType
	//  交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	//  产品代码
	ProductID types.TThostFtdcInstrumentIDType
}

// 投资者品种/跨品种保证金产品组
type CThostFtdcProductGroupField struct {
	//  保留的无效字段
	Reserve1 types.TThostFtdcOldInstrumentIDType
	//  交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	//  保留的无效字段
	Reserve2 types.TThostFtdcOldInstrumentIDType
	//  产品代码
	ProductID types.TThostFtdcInstrumentIDType
	//  产品组代码
	ProductGroupID types.TThostFtdcInstrumentIDType
}

// 交易所公告
type CThostFtdcBulletinField struct {
	//  交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	//  交易日
	TradingDay types.TThostFtdcDateType
	//  公告编号
	BulletinID types.TThostFtdcBulletinIDType
	//  序列号
	SequenceNo types.TThostFtdcSequenceNoType
	//  公告类型
	NewsType types.TThostFtdcNewsTypeType
	//  紧急程度
	NewsUrgency types.TThostFtdcNewsUrgencyType
	//  发送时间
	SendTime types.TThostFtdcTimeType
	//  消息摘要
	Abstract types.TThostFtdcAbstractType
	//  消息来源
	ComeFrom types.TThostFtdcComeFromType
	//  消息正文
	Content types.TThostFtdcContentType
	//  WEB地址
	URLLink types.TThostFtdcURLLinkType
	//  市场代码
	MarketID types.TThostFtdcMarketIDType
}

// 查询交易所公告
type CThostFtdcQryBulletinField struct {
	//  交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	//  公告编号
	BulletinID types.TThostFtdcBulletinIDType
	//  序列号
	SequenceNo types.TThostFtdcSequenceNoType
	//  公告类型
	NewsType types.TThostFtdcNewsTypeType
	//  紧急程度
	NewsUrgency types.TThostFtdcNewsUrgencyType
}

// MulticastInstrument
type CThostFtdcMulticastInstrumentField struct {
	//  主题号
	TopicID types.TThostFtdcInstallIDType
	//  保留的无效字段
	Reserve1 types.TThostFtdcOldInstrumentIDType
	//  合约编号
	InstrumentNo types.TThostFtdcInstallIDType
	//  基准价
	CodePrice types.TThostFtdcPriceType
	//  合约数量乘数
	VolumeMultiple types.TThostFtdcVolumeMultipleType
	//  最小变动价位
	PriceTick types.TThostFtdcPriceType
	//  合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
}

// QryMulticastInstrument
type CThostFtdcQryMulticastInstrumentField struct {
	//  主题号
	TopicID types.TThostFtdcInstallIDType
	//  保留的无效字段
	Reserve1 types.TThostFtdcOldInstrumentIDType
	//  合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
}

// App客户端权限分配
type CThostFtdcAppIDAuthAssignField struct {
	//  经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	//  App代码
	AppID types.TThostFtdcAppIDType
	//  交易中心代码
	DRIdentityID types.TThostFtdcDRIdentityIDType
}

// 转帐开户请求
type CThostFtdcReqOpenAccountField struct {
	//  业务功能码
	TradeCode types.TThostFtdcTradeCodeType
	//  银行代码
	BankID types.TThostFtdcBankIDType
	//  银行分支机构代码
	BankBranchID types.TThostFtdcBankBrchIDType
	//  期商代码
	BrokerID types.TThostFtdcBrokerIDType
	//  期商分支机构代码
	BrokerBranchID types.TThostFtdcFutureBranchIDType
	//  交易日期
	TradeDate types.TThostFtdcTradeDateType
	//  交易时间
	TradeTime types.TThostFtdcTradeTimeType
	//  银行流水号
	BankSerial types.TThostFtdcBankSerialType
	//  交易系统日期
	TradingDay types.TThostFtdcTradeDateType
	//  银期平台消息流水号
	PlateSerial types.TThostFtdcSerialType
	//  最后分片标志
	LastFragment types.TThostFtdcLastFragmentType
	//  会话号
	SessionID types.TThostFtdcSessionIDType
	//  客户姓名
	CustomerName types.TThostFtdcIndividualNameType
	//  证件类型
	IdCardType types.TThostFtdcIdCardTypeType
	//  证件号码
	IdentifiedCardNo types.TThostFtdcIdentifiedCardNoType
	//  性别
	Gender types.TThostFtdcGenderType
	//  国家代码
	CountryCode types.TThostFtdcCountryCodeType
	//  客户类型
	CustType types.TThostFtdcCustTypeType
	//  地址
	Address types.TThostFtdcAddressType
	//  邮编
	ZipCode types.TThostFtdcZipCodeType
	//  电话号码
	Telephone types.TThostFtdcTelephoneType
	//  手机
	MobilePhone types.TThostFtdcMobilePhoneType
	//  传真
	Fax types.TThostFtdcFaxType
	//  电子邮件
	EMail types.TThostFtdcEMailType
	//  资金账户状态
	MoneyAccountStatus types.TThostFtdcMoneyAccountStatusType
	//  银行帐号
	BankAccount types.TThostFtdcBankAccountType
	//  银行密码
	BankPassWord types.TThostFtdcPasswordType
	//  投资者帐号
	AccountID types.TThostFtdcAccountIDType
	//  期货密码
	Password types.TThostFtdcPasswordType
	//  安装编号
	InstallID types.TThostFtdcInstallIDType
	//  验证客户证件号码标志
	VerifyCertNoFlag types.TThostFtdcYesNoIndicatorType
	//  币种代码
	CurrencyID types.TThostFtdcCurrencyIDType
	//  汇钞标志
	CashExchangeCode types.TThostFtdcCashExchangeCodeType
	//  摘要
	Digest types.TThostFtdcDigestType
	//  银行帐号类型
	BankAccType types.TThostFtdcBankAccTypeType
	//  渠道标志
	DeviceID types.TThostFtdcDeviceIDType
	//  期货单位帐号类型
	BankSecuAccType types.TThostFtdcBankAccTypeType
	//  期货公司银行编码
	BrokerIDByBank types.TThostFtdcBankCodingForFutureType
	//  期货单位帐号
	BankSecuAcc types.TThostFtdcBankAccountType
	//  银行密码标志
	BankPwdFlag types.TThostFtdcPwdFlagType
	//  期货资金密码核对标志
	SecuPwdFlag types.TThostFtdcPwdFlagType
	//  交易柜员
	OperNo types.TThostFtdcOperNoType
	//  交易ID
	TID types.TThostFtdcTIDType
	//  用户标识
	UserID types.TThostFtdcUserIDType
	//  长客户姓名
	LongCustomerName types.TThostFtdcLongIndividualNameType
}

// 转帐销户请求
type CThostFtdcReqCancelAccountField struct {
	//  业务功能码
	TradeCode types.TThostFtdcTradeCodeType
	//  银行代码
	BankID types.TThostFtdcBankIDType
	//  银行分支机构代码
	BankBranchID types.TThostFtdcBankBrchIDType
	//  期商代码
	BrokerID types.TThostFtdcBrokerIDType
	//  期商分支机构代码
	BrokerBranchID types.TThostFtdcFutureBranchIDType
	//  交易日期
	TradeDate types.TThostFtdcTradeDateType
	//  交易时间
	TradeTime types.TThostFtdcTradeTimeType
	//  银行流水号
	BankSerial types.TThostFtdcBankSerialType
	//  交易系统日期
	TradingDay types.TThostFtdcTradeDateType
	//  银期平台消息流水号
	PlateSerial types.TThostFtdcSerialType
	//  最后分片标志
	LastFragment types.TThostFtdcLastFragmentType
	//  会话号
	SessionID types.TThostFtdcSessionIDType
	//  客户姓名
	CustomerName types.TThostFtdcIndividualNameType
	//  证件类型
	IdCardType types.TThostFtdcIdCardTypeType
	//  证件号码
	IdentifiedCardNo types.TThostFtdcIdentifiedCardNoType
	//  性别
	Gender types.TThostFtdcGenderType
	//  国家代码
	CountryCode types.TThostFtdcCountryCodeType
	//  客户类型
	CustType types.TThostFtdcCustTypeType
	//  地址
	Address types.TThostFtdcAddressType
	//  邮编
	ZipCode types.TThostFtdcZipCodeType
	//  电话号码
	Telephone types.TThostFtdcTelephoneType
	//  手机
	MobilePhone types.TThostFtdcMobilePhoneType
	//  传真
	Fax types.TThostFtdcFaxType
	//  电子邮件
	EMail types.TThostFtdcEMailType
	//  资金账户状态
	MoneyAccountStatus types.TThostFtdcMoneyAccountStatusType
	//  银行帐号
	BankAccount types.TThostFtdcBankAccountType
	//  银行密码
	BankPassWord types.TThostFtdcPasswordType
	//  投资者帐号
	AccountID types.TThostFtdcAccountIDType
	//  期货密码
	Password types.TThostFtdcPasswordType
	//  安装编号
	InstallID types.TThostFtdcInstallIDType
	//  验证客户证件号码标志
	VerifyCertNoFlag types.TThostFtdcYesNoIndicatorType
	//  币种代码
	CurrencyID types.TThostFtdcCurrencyIDType
	//  汇钞标志
	CashExchangeCode types.TThostFtdcCashExchangeCodeType
	//  摘要
	Digest types.TThostFtdcDigestType
	//  银行帐号类型
	BankAccType types.TThostFtdcBankAccTypeType
	//  渠道标志
	DeviceID types.TThostFtdcDeviceIDType
	//  期货单位帐号类型
	BankSecuAccType types.TThostFtdcBankAccTypeType
	//  期货公司银行编码
	BrokerIDByBank types.TThostFtdcBankCodingForFutureType
	//  期货单位帐号
	BankSecuAcc types.TThostFtdcBankAccountType
	//  银行密码标志
	BankPwdFlag types.TThostFtdcPwdFlagType
	//  期货资金密码核对标志
	SecuPwdFlag types.TThostFtdcPwdFlagType
	//  交易柜员
	OperNo types.TThostFtdcOperNoType
	//  交易ID
	TID types.TThostFtdcTIDType
	//  用户标识
	UserID types.TThostFtdcUserIDType
	//  长客户姓名
	LongCustomerName types.TThostFtdcLongIndividualNameType
}

// 变更银行账户请求
type CThostFtdcReqChangeAccountField struct {
	//  业务功能码
	TradeCode types.TThostFtdcTradeCodeType
	//  银行代码
	BankID types.TThostFtdcBankIDType
	//  银行分支机构代码
	BankBranchID types.TThostFtdcBankBrchIDType
	//  期商代码
	BrokerID types.TThostFtdcBrokerIDType
	//  期商分支机构代码
	BrokerBranchID types.TThostFtdcFutureBranchIDType
	//  交易日期
	TradeDate types.TThostFtdcTradeDateType
	//  交易时间
	TradeTime types.TThostFtdcTradeTimeType
	//  银行流水号
	BankSerial types.TThostFtdcBankSerialType
	//  交易系统日期
	TradingDay types.TThostFtdcTradeDateType
	//  银期平台消息流水号
	PlateSerial types.TThostFtdcSerialType
	//  最后分片标志
	LastFragment types.TThostFtdcLastFragmentType
	//  会话号
	SessionID types.TThostFtdcSessionIDType
	//  客户姓名
	CustomerName types.TThostFtdcIndividualNameType
	//  证件类型
	IdCardType types.TThostFtdcIdCardTypeType
	//  证件号码
	IdentifiedCardNo types.TThostFtdcIdentifiedCardNoType
	//  性别
	Gender types.TThostFtdcGenderType
	//  国家代码
	CountryCode types.TThostFtdcCountryCodeType
	//  客户类型
	CustType types.TThostFtdcCustTypeType
	//  地址
	Address types.TThostFtdcAddressType
	//  邮编
	ZipCode types.TThostFtdcZipCodeType
	//  电话号码
	Telephone types.TThostFtdcTelephoneType
	//  手机
	MobilePhone types.TThostFtdcMobilePhoneType
	//  传真
	Fax types.TThostFtdcFaxType
	//  电子邮件
	EMail types.TThostFtdcEMailType
	//  资金账户状态
	MoneyAccountStatus types.TThostFtdcMoneyAccountStatusType
	//  银行帐号
	BankAccount types.TThostFtdcBankAccountType
	//  银行密码
	BankPassWord types.TThostFtdcPasswordType
	//  新银行帐号
	NewBankAccount types.TThostFtdcBankAccountType
	//  新银行密码
	NewBankPassWord types.TThostFtdcPasswordType
	//  投资者帐号
	AccountID types.TThostFtdcAccountIDType
	//  期货密码
	Password types.TThostFtdcPasswordType
	//  银行帐号类型
	BankAccType types.TThostFtdcBankAccTypeType
	//  安装编号
	InstallID types.TThostFtdcInstallIDType
	//  验证客户证件号码标志
	VerifyCertNoFlag types.TThostFtdcYesNoIndicatorType
	//  币种代码
	CurrencyID types.TThostFtdcCurrencyIDType
	//  期货公司银行编码
	BrokerIDByBank types.TThostFtdcBankCodingForFutureType
	//  银行密码标志
	BankPwdFlag types.TThostFtdcPwdFlagType
	//  期货资金密码核对标志
	SecuPwdFlag types.TThostFtdcPwdFlagType
	//  交易ID
	TID types.TThostFtdcTIDType
	//  摘要
	Digest types.TThostFtdcDigestType
	//  长客户姓名
	LongCustomerName types.TThostFtdcLongIndividualNameType
}

// 转账请求
type CThostFtdcReqTransferField struct {
	//  业务功能码
	TradeCode types.TThostFtdcTradeCodeType
	//  银行代码
	BankID types.TThostFtdcBankIDType
	//  银行分支机构代码
	BankBranchID types.TThostFtdcBankBrchIDType
	//  期商代码
	BrokerID types.TThostFtdcBrokerIDType
	//  期商分支机构代码
	BrokerBranchID types.TThostFtdcFutureBranchIDType
	//  交易日期
	TradeDate types.TThostFtdcTradeDateType
	//  交易时间
	TradeTime types.TThostFtdcTradeTimeType
	//  银行流水号
	BankSerial types.TThostFtdcBankSerialType
	//  交易系统日期
	TradingDay types.TThostFtdcTradeDateType
	//  银期平台消息流水号
	PlateSerial types.TThostFtdcSerialType
	//  最后分片标志
	LastFragment types.TThostFtdcLastFragmentType
	//  会话号
	SessionID types.TThostFtdcSessionIDType
	//  客户姓名
	CustomerName types.TThostFtdcIndividualNameType
	//  证件类型
	IdCardType types.TThostFtdcIdCardTypeType
	//  证件号码
	IdentifiedCardNo types.TThostFtdcIdentifiedCardNoType
	//  客户类型
	CustType types.TThostFtdcCustTypeType
	//  银行帐号
	BankAccount types.TThostFtdcBankAccountType
	//  银行密码
	BankPassWord types.TThostFtdcPasswordType
	//  投资者帐号
	AccountID types.TThostFtdcAccountIDType
	//  期货密码
	Password types.TThostFtdcPasswordType
	//  安装编号
	InstallID types.TThostFtdcInstallIDType
	//  期货公司流水号
	FutureSerial types.TThostFtdcFutureSerialType
	//  用户标识
	UserID types.TThostFtdcUserIDType
	//  验证客户证件号码标志
	VerifyCertNoFlag types.TThostFtdcYesNoIndicatorType
	//  币种代码
	CurrencyID types.TThostFtdcCurrencyIDType
	//  转帐金额
	TradeAmount types.TThostFtdcTradeAmountType
	//  期货可取金额
	FutureFetchAmount types.TThostFtdcTradeAmountType
	//  费用支付标志
	FeePayFlag types.TThostFtdcFeePayFlagType
	//  应收客户费用
	CustFee types.TThostFtdcCustFeeType
	//  应收期货公司费用
	BrokerFee types.TThostFtdcFutureFeeType
	//  发送方给接收方的消息
	Message types.TThostFtdcAddInfoType
	//  摘要
	Digest types.TThostFtdcDigestType
	//  银行帐号类型
	BankAccType types.TThostFtdcBankAccTypeType
	//  渠道标志
	DeviceID types.TThostFtdcDeviceIDType
	//  期货单位帐号类型
	BankSecuAccType types.TThostFtdcBankAccTypeType
	//  期货公司银行编码
	BrokerIDByBank types.TThostFtdcBankCodingForFutureType
	//  期货单位帐号
	BankSecuAcc types.TThostFtdcBankAccountType
	//  银行密码标志
	BankPwdFlag types.TThostFtdcPwdFlagType
	//  期货资金密码核对标志
	SecuPwdFlag types.TThostFtdcPwdFlagType
	//  交易柜员
	OperNo types.TThostFtdcOperNoType
	//  请求编号
	RequestID types.TThostFtdcRequestIDType
	//  交易ID
	TID types.TThostFtdcTIDType
	//  转账交易状态
	TransferStatus types.TThostFtdcTransferStatusType
	//  长客户姓名
	LongCustomerName types.TThostFtdcLongIndividualNameType
}

// 银行发起银行资金转期货响应
type CThostFtdcRspTransferField struct {
	//  业务功能码
	TradeCode types.TThostFtdcTradeCodeType
	//  银行代码
	BankID types.TThostFtdcBankIDType
	//  银行分支机构代码
	BankBranchID types.TThostFtdcBankBrchIDType
	//  期商代码
	BrokerID types.TThostFtdcBrokerIDType
	//  期商分支机构代码
	BrokerBranchID types.TThostFtdcFutureBranchIDType
	//  交易日期
	TradeDate types.TThostFtdcTradeDateType
	//  交易时间
	TradeTime types.TThostFtdcTradeTimeType
	//  银行流水号
	BankSerial types.TThostFtdcBankSerialType
	//  交易系统日期
	TradingDay types.TThostFtdcTradeDateType
	//  银期平台消息流水号
	PlateSerial types.TThostFtdcSerialType
	//  最后分片标志
	LastFragment types.TThostFtdcLastFragmentType
	//  会话号
	SessionID types.TThostFtdcSessionIDType
	//  客户姓名
	CustomerName types.TThostFtdcIndividualNameType
	//  证件类型
	IdCardType types.TThostFtdcIdCardTypeType
	//  证件号码
	IdentifiedCardNo types.TThostFtdcIdentifiedCardNoType
	//  客户类型
	CustType types.TThostFtdcCustTypeType
	//  银行帐号
	BankAccount types.TThostFtdcBankAccountType
	//  银行密码
	BankPassWord types.TThostFtdcPasswordType
	//  投资者帐号
	AccountID types.TThostFtdcAccountIDType
	//  期货密码
	Password types.TThostFtdcPasswordType
	//  安装编号
	InstallID types.TThostFtdcInstallIDType
	//  期货公司流水号
	FutureSerial types.TThostFtdcFutureSerialType
	//  用户标识
	UserID types.TThostFtdcUserIDType
	//  验证客户证件号码标志
	VerifyCertNoFlag types.TThostFtdcYesNoIndicatorType
	//  币种代码
	CurrencyID types.TThostFtdcCurrencyIDType
	//  转帐金额
	TradeAmount types.TThostFtdcTradeAmountType
	//  期货可取金额
	FutureFetchAmount types.TThostFtdcTradeAmountType
	//  费用支付标志
	FeePayFlag types.TThostFtdcFeePayFlagType
	//  应收客户费用
	CustFee types.TThostFtdcCustFeeType
	//  应收期货公司费用
	BrokerFee types.TThostFtdcFutureFeeType
	//  发送方给接收方的消息
	Message types.TThostFtdcAddInfoType
	//  摘要
	Digest types.TThostFtdcDigestType
	//  银行帐号类型
	BankAccType types.TThostFtdcBankAccTypeType
	//  渠道标志
	DeviceID types.TThostFtdcDeviceIDType
	//  期货单位帐号类型
	BankSecuAccType types.TThostFtdcBankAccTypeType
	//  期货公司银行编码
	BrokerIDByBank types.TThostFtdcBankCodingForFutureType
	//  期货单位帐号
	BankSecuAcc types.TThostFtdcBankAccountType
	//  银行密码标志
	BankPwdFlag types.TThostFtdcPwdFlagType
	//  期货资金密码核对标志
	SecuPwdFlag types.TThostFtdcPwdFlagType
	//  交易柜员
	OperNo types.TThostFtdcOperNoType
	//  请求编号
	RequestID types.TThostFtdcRequestIDType
	//  交易ID
	TID types.TThostFtdcTIDType
	//  转账交易状态
	TransferStatus types.TThostFtdcTransferStatusType
	//  错误代码
	ErrorID types.TThostFtdcErrorIDType
	//  错误信息
	ErrorMsg types.TThostFtdcErrorMsgType
	//  长客户姓名
	LongCustomerName types.TThostFtdcLongIndividualNameType
}

// 冲正请求
type CThostFtdcReqRepealField struct {
	//  冲正时间间隔
	RepealTimeInterval types.TThostFtdcRepealTimeIntervalType
	//  已经冲正次数
	RepealedTimes types.TThostFtdcRepealedTimesType
	//  银行冲正标志
	BankRepealFlag types.TThostFtdcBankRepealFlagType
	//  期商冲正标志
	BrokerRepealFlag types.TThostFtdcBrokerRepealFlagType
	//  被冲正平台流水号
	PlateRepealSerial types.TThostFtdcPlateSerialType
	//  被冲正银行流水号
	BankRepealSerial types.TThostFtdcBankSerialType
	//  被冲正期货流水号
	FutureRepealSerial types.TThostFtdcFutureSerialType
	//  业务功能码
	TradeCode types.TThostFtdcTradeCodeType
	//  银行代码
	BankID types.TThostFtdcBankIDType
	//  银行分支机构代码
	BankBranchID types.TThostFtdcBankBrchIDType
	//  期商代码
	BrokerID types.TThostFtdcBrokerIDType
	//  期商分支机构代码
	BrokerBranchID types.TThostFtdcFutureBranchIDType
	//  交易日期
	TradeDate types.TThostFtdcTradeDateType
	//  交易时间
	TradeTime types.TThostFtdcTradeTimeType
	//  银行流水号
	BankSerial types.TThostFtdcBankSerialType
	//  交易系统日期
	TradingDay types.TThostFtdcTradeDateType
	//  银期平台消息流水号
	PlateSerial types.TThostFtdcSerialType
	//  最后分片标志
	LastFragment types.TThostFtdcLastFragmentType
	//  会话号
	SessionID types.TThostFtdcSessionIDType
	//  客户姓名
	CustomerName types.TThostFtdcIndividualNameType
	//  证件类型
	IdCardType types.TThostFtdcIdCardTypeType
	//  证件号码
	IdentifiedCardNo types.TThostFtdcIdentifiedCardNoType
	//  客户类型
	CustType types.TThostFtdcCustTypeType
	//  银行帐号
	BankAccount types.TThostFtdcBankAccountType
	//  银行密码
	BankPassWord types.TThostFtdcPasswordType
	//  投资者帐号
	AccountID types.TThostFtdcAccountIDType
	//  期货密码
	Password types.TThostFtdcPasswordType
	//  安装编号
	InstallID types.TThostFtdcInstallIDType
	//  期货公司流水号
	FutureSerial types.TThostFtdcFutureSerialType
	//  用户标识
	UserID types.TThostFtdcUserIDType
	//  验证客户证件号码标志
	VerifyCertNoFlag types.TThostFtdcYesNoIndicatorType
	//  币种代码
	CurrencyID types.TThostFtdcCurrencyIDType
	//  转帐金额
	TradeAmount types.TThostFtdcTradeAmountType
	//  期货可取金额
	FutureFetchAmount types.TThostFtdcTradeAmountType
	//  费用支付标志
	FeePayFlag types.TThostFtdcFeePayFlagType
	//  应收客户费用
	CustFee types.TThostFtdcCustFeeType
	//  应收期货公司费用
	BrokerFee types.TThostFtdcFutureFeeType
	//  发送方给接收方的消息
	Message types.TThostFtdcAddInfoType
	//  摘要
	Digest types.TThostFtdcDigestType
	//  银行帐号类型
	BankAccType types.TThostFtdcBankAccTypeType
	//  渠道标志
	DeviceID types.TThostFtdcDeviceIDType
	//  期货单位帐号类型
	BankSecuAccType types.TThostFtdcBankAccTypeType
	//  期货公司银行编码
	BrokerIDByBank types.TThostFtdcBankCodingForFutureType
	//  期货单位帐号
	BankSecuAcc types.TThostFtdcBankAccountType
	//  银行密码标志
	BankPwdFlag types.TThostFtdcPwdFlagType
	//  期货资金密码核对标志
	SecuPwdFlag types.TThostFtdcPwdFlagType
	//  交易柜员
	OperNo types.TThostFtdcOperNoType
	//  请求编号
	RequestID types.TThostFtdcRequestIDType
	//  交易ID
	TID types.TThostFtdcTIDType
	//  转账交易状态
	TransferStatus types.TThostFtdcTransferStatusType
	//  长客户姓名
	LongCustomerName types.TThostFtdcLongIndividualNameType
}

// 冲正响应
type CThostFtdcRspRepealField struct {
	//  冲正时间间隔
	RepealTimeInterval types.TThostFtdcRepealTimeIntervalType
	//  已经冲正次数
	RepealedTimes types.TThostFtdcRepealedTimesType
	//  银行冲正标志
	BankRepealFlag types.TThostFtdcBankRepealFlagType
	//  期商冲正标志
	BrokerRepealFlag types.TThostFtdcBrokerRepealFlagType
	//  被冲正平台流水号
	PlateRepealSerial types.TThostFtdcPlateSerialType
	//  被冲正银行流水号
	BankRepealSerial types.TThostFtdcBankSerialType
	//  被冲正期货流水号
	FutureRepealSerial types.TThostFtdcFutureSerialType
	//  业务功能码
	TradeCode types.TThostFtdcTradeCodeType
	//  银行代码
	BankID types.TThostFtdcBankIDType
	//  银行分支机构代码
	BankBranchID types.TThostFtdcBankBrchIDType
	//  期商代码
	BrokerID types.TThostFtdcBrokerIDType
	//  期商分支机构代码
	BrokerBranchID types.TThostFtdcFutureBranchIDType
	//  交易日期
	TradeDate types.TThostFtdcTradeDateType
	//  交易时间
	TradeTime types.TThostFtdcTradeTimeType
	//  银行流水号
	BankSerial types.TThostFtdcBankSerialType
	//  交易系统日期
	TradingDay types.TThostFtdcTradeDateType
	//  银期平台消息流水号
	PlateSerial types.TThostFtdcSerialType
	//  最后分片标志
	LastFragment types.TThostFtdcLastFragmentType
	//  会话号
	SessionID types.TThostFtdcSessionIDType
	//  客户姓名
	CustomerName types.TThostFtdcIndividualNameType
	//  证件类型
	IdCardType types.TThostFtdcIdCardTypeType
	//  证件号码
	IdentifiedCardNo types.TThostFtdcIdentifiedCardNoType
	//  客户类型
	CustType types.TThostFtdcCustTypeType
	//  银行帐号
	BankAccount types.TThostFtdcBankAccountType
	//  银行密码
	BankPassWord types.TThostFtdcPasswordType
	//  投资者帐号
	AccountID types.TThostFtdcAccountIDType
	//  期货密码
	Password types.TThostFtdcPasswordType
	//  安装编号
	InstallID types.TThostFtdcInstallIDType
	//  期货公司流水号
	FutureSerial types.TThostFtdcFutureSerialType
	//  用户标识
	UserID types.TThostFtdcUserIDType
	//  验证客户证件号码标志
	VerifyCertNoFlag types.TThostFtdcYesNoIndicatorType
	//  币种代码
	CurrencyID types.TThostFtdcCurrencyIDType
	//  转帐金额
	TradeAmount types.TThostFtdcTradeAmountType
	//  期货可取金额
	FutureFetchAmount types.TThostFtdcTradeAmountType
	//  费用支付标志
	FeePayFlag types.TThostFtdcFeePayFlagType
	//  应收客户费用
	CustFee types.TThostFtdcCustFeeType
	//  应收期货公司费用
	BrokerFee types.TThostFtdcFutureFeeType
	//  发送方给接收方的消息
	Message types.TThostFtdcAddInfoType
	//  摘要
	Digest types.TThostFtdcDigestType
	//  银行帐号类型
	BankAccType types.TThostFtdcBankAccTypeType
	//  渠道标志
	DeviceID types.TThostFtdcDeviceIDType
	//  期货单位帐号类型
	BankSecuAccType types.TThostFtdcBankAccTypeType
	//  期货公司银行编码
	BrokerIDByBank types.TThostFtdcBankCodingForFutureType
	//  期货单位帐号
	BankSecuAcc types.TThostFtdcBankAccountType
	//  银行密码标志
	BankPwdFlag types.TThostFtdcPwdFlagType
	//  期货资金密码核对标志
	SecuPwdFlag types.TThostFtdcPwdFlagType
	//  交易柜员
	OperNo types.TThostFtdcOperNoType
	//  请求编号
	RequestID types.TThostFtdcRequestIDType
	//  交易ID
	TID types.TThostFtdcTIDType
	//  转账交易状态
	TransferStatus types.TThostFtdcTransferStatusType
	//  错误代码
	ErrorID types.TThostFtdcErrorIDType
	//  错误信息
	ErrorMsg types.TThostFtdcErrorMsgType
	//  长客户姓名
	LongCustomerName types.TThostFtdcLongIndividualNameType
}

// 查询账户信息请求
type CThostFtdcReqQueryAccountField struct {
	//  业务功能码
	TradeCode types.TThostFtdcTradeCodeType
	//  银行代码
	BankID types.TThostFtdcBankIDType
	//  银行分支机构代码
	BankBranchID types.TThostFtdcBankBrchIDType
	//  期商代码
	BrokerID types.TThostFtdcBrokerIDType
	//  期商分支机构代码
	BrokerBranchID types.TThostFtdcFutureBranchIDType
	//  交易日期
	TradeDate types.TThostFtdcTradeDateType
	//  交易时间
	TradeTime types.TThostFtdcTradeTimeType
	//  银行流水号
	BankSerial types.TThostFtdcBankSerialType
	//  交易系统日期
	TradingDay types.TThostFtdcTradeDateType
	//  银期平台消息流水号
	PlateSerial types.TThostFtdcSerialType
	//  最后分片标志
	LastFragment types.TThostFtdcLastFragmentType
	//  会话号
	SessionID types.TThostFtdcSessionIDType
	//  客户姓名
	CustomerName types.TThostFtdcIndividualNameType
	//  证件类型
	IdCardType types.TThostFtdcIdCardTypeType
	//  证件号码
	IdentifiedCardNo types.TThostFtdcIdentifiedCardNoType
	//  客户类型
	CustType types.TThostFtdcCustTypeType
	//  银行帐号
	BankAccount types.TThostFtdcBankAccountType
	//  银行密码
	BankPassWord types.TThostFtdcPasswordType
	//  投资者帐号
	AccountID types.TThostFtdcAccountIDType
	//  期货密码
	Password types.TThostFtdcPasswordType
	//  期货公司流水号
	FutureSerial types.TThostFtdcFutureSerialType
	//  安装编号
	InstallID types.TThostFtdcInstallIDType
	//  用户标识
	UserID types.TThostFtdcUserIDType
	//  验证客户证件号码标志
	VerifyCertNoFlag types.TThostFtdcYesNoIndicatorType
	//  币种代码
	CurrencyID types.TThostFtdcCurrencyIDType
	//  摘要
	Digest types.TThostFtdcDigestType
	//  银行帐号类型
	BankAccType types.TThostFtdcBankAccTypeType
	//  渠道标志
	DeviceID types.TThostFtdcDeviceIDType
	//  期货单位帐号类型
	BankSecuAccType types.TThostFtdcBankAccTypeType
	//  期货公司银行编码
	BrokerIDByBank types.TThostFtdcBankCodingForFutureType
	//  期货单位帐号
	BankSecuAcc types.TThostFtdcBankAccountType
	//  银行密码标志
	BankPwdFlag types.TThostFtdcPwdFlagType
	//  期货资金密码核对标志
	SecuPwdFlag types.TThostFtdcPwdFlagType
	//  交易柜员
	OperNo types.TThostFtdcOperNoType
	//  请求编号
	RequestID types.TThostFtdcRequestIDType
	//  交易ID
	TID types.TThostFtdcTIDType
	//  长客户姓名
	LongCustomerName types.TThostFtdcLongIndividualNameType
}

// 查询账户信息响应
type CThostFtdcRspQueryAccountField struct {
	//  业务功能码
	TradeCode types.TThostFtdcTradeCodeType
	//  银行代码
	BankID types.TThostFtdcBankIDType
	//  银行分支机构代码
	BankBranchID types.TThostFtdcBankBrchIDType
	//  期商代码
	BrokerID types.TThostFtdcBrokerIDType
	//  期商分支机构代码
	BrokerBranchID types.TThostFtdcFutureBranchIDType
	//  交易日期
	TradeDate types.TThostFtdcTradeDateType
	//  交易时间
	TradeTime types.TThostFtdcTradeTimeType
	//  银行流水号
	BankSerial types.TThostFtdcBankSerialType
	//  交易系统日期
	TradingDay types.TThostFtdcTradeDateType
	//  银期平台消息流水号
	PlateSerial types.TThostFtdcSerialType
	//  最后分片标志
	LastFragment types.TThostFtdcLastFragmentType
	//  会话号
	SessionID types.TThostFtdcSessionIDType
	//  客户姓名
	CustomerName types.TThostFtdcIndividualNameType
	//  证件类型
	IdCardType types.TThostFtdcIdCardTypeType
	//  证件号码
	IdentifiedCardNo types.TThostFtdcIdentifiedCardNoType
	//  客户类型
	CustType types.TThostFtdcCustTypeType
	//  银行帐号
	BankAccount types.TThostFtdcBankAccountType
	//  银行密码
	BankPassWord types.TThostFtdcPasswordType
	//  投资者帐号
	AccountID types.TThostFtdcAccountIDType
	//  期货密码
	Password types.TThostFtdcPasswordType
	//  期货公司流水号
	FutureSerial types.TThostFtdcFutureSerialType
	//  安装编号
	InstallID types.TThostFtdcInstallIDType
	//  用户标识
	UserID types.TThostFtdcUserIDType
	//  验证客户证件号码标志
	VerifyCertNoFlag types.TThostFtdcYesNoIndicatorType
	//  币种代码
	CurrencyID types.TThostFtdcCurrencyIDType
	//  摘要
	Digest types.TThostFtdcDigestType
	//  银行帐号类型
	BankAccType types.TThostFtdcBankAccTypeType
	//  渠道标志
	DeviceID types.TThostFtdcDeviceIDType
	//  期货单位帐号类型
	BankSecuAccType types.TThostFtdcBankAccTypeType
	//  期货公司银行编码
	BrokerIDByBank types.TThostFtdcBankCodingForFutureType
	//  期货单位帐号
	BankSecuAcc types.TThostFtdcBankAccountType
	//  银行密码标志
	BankPwdFlag types.TThostFtdcPwdFlagType
	//  期货资金密码核对标志
	SecuPwdFlag types.TThostFtdcPwdFlagType
	//  交易柜员
	OperNo types.TThostFtdcOperNoType
	//  请求编号
	RequestID types.TThostFtdcRequestIDType
	//  交易ID
	TID types.TThostFtdcTIDType
	//  银行可用金额
	BankUseAmount types.TThostFtdcTradeAmountType
	//  银行可取金额
	BankFetchAmount types.TThostFtdcTradeAmountType
	//  长客户姓名
	LongCustomerName types.TThostFtdcLongIndividualNameType
}

// 期商签到签退
type CThostFtdcFutureSignIOField struct {
	//  业务功能码
	TradeCode types.TThostFtdcTradeCodeType
	//  银行代码
	BankID types.TThostFtdcBankIDType
	//  银行分支机构代码
	BankBranchID types.TThostFtdcBankBrchIDType
	//  期商代码
	BrokerID types.TThostFtdcBrokerIDType
	//  期商分支机构代码
	BrokerBranchID types.TThostFtdcFutureBranchIDType
	//  交易日期
	TradeDate types.TThostFtdcTradeDateType
	//  交易时间
	TradeTime types.TThostFtdcTradeTimeType
	//  银行流水号
	BankSerial types.TThostFtdcBankSerialType
	//  交易系统日期
	TradingDay types.TThostFtdcTradeDateType
	//  银期平台消息流水号
	PlateSerial types.TThostFtdcSerialType
	//  最后分片标志
	LastFragment types.TThostFtdcLastFragmentType
	//  会话号
	SessionID types.TThostFtdcSessionIDType
	//  安装编号
	InstallID types.TThostFtdcInstallIDType
	//  用户标识
	UserID types.TThostFtdcUserIDType
	//  摘要
	Digest types.TThostFtdcDigestType
	//  币种代码
	CurrencyID types.TThostFtdcCurrencyIDType
	//  渠道标志
	DeviceID types.TThostFtdcDeviceIDType
	//  期货公司银行编码
	BrokerIDByBank types.TThostFtdcBankCodingForFutureType
	//  交易柜员
	OperNo types.TThostFtdcOperNoType
	//  请求编号
	RequestID types.TThostFtdcRequestIDType
	//  交易ID
	TID types.TThostFtdcTIDType
}

// 期商签到响应
type CThostFtdcRspFutureSignInField struct {
	//  业务功能码
	TradeCode types.TThostFtdcTradeCodeType
	//  银行代码
	BankID types.TThostFtdcBankIDType
	//  银行分支机构代码
	BankBranchID types.TThostFtdcBankBrchIDType
	//  期商代码
	BrokerID types.TThostFtdcBrokerIDType
	//  期商分支机构代码
	BrokerBranchID types.TThostFtdcFutureBranchIDType
	//  交易日期
	TradeDate types.TThostFtdcTradeDateType
	//  交易时间
	TradeTime types.TThostFtdcTradeTimeType
	//  银行流水号
	BankSerial types.TThostFtdcBankSerialType
	//  交易系统日期
	TradingDay types.TThostFtdcTradeDateType
	//  银期平台消息流水号
	PlateSerial types.TThostFtdcSerialType
	//  最后分片标志
	LastFragment types.TThostFtdcLastFragmentType
	//  会话号
	SessionID types.TThostFtdcSessionIDType
	//  安装编号
	InstallID types.TThostFtdcInstallIDType
	//  用户标识
	UserID types.TThostFtdcUserIDType
	//  摘要
	Digest types.TThostFtdcDigestType
	//  币种代码
	CurrencyID types.TThostFtdcCurrencyIDType
	//  渠道标志
	DeviceID types.TThostFtdcDeviceIDType
	//  期货公司银行编码
	BrokerIDByBank types.TThostFtdcBankCodingForFutureType
	//  交易柜员
	OperNo types.TThostFtdcOperNoType
	//  请求编号
	RequestID types.TThostFtdcRequestIDType
	//  交易ID
	TID types.TThostFtdcTIDType
	//  错误代码
	ErrorID types.TThostFtdcErrorIDType
	//  错误信息
	ErrorMsg types.TThostFtdcErrorMsgType
	//  PIN密钥
	PinKey types.TThostFtdcPasswordKeyType
	//  MAC密钥
	MacKey types.TThostFtdcPasswordKeyType
}

// 期商签退请求
type CThostFtdcReqFutureSignOutField struct {
	//  业务功能码
	TradeCode types.TThostFtdcTradeCodeType
	//  银行代码
	BankID types.TThostFtdcBankIDType
	//  银行分支机构代码
	BankBranchID types.TThostFtdcBankBrchIDType
	//  期商代码
	BrokerID types.TThostFtdcBrokerIDType
	//  期商分支机构代码
	BrokerBranchID types.TThostFtdcFutureBranchIDType
	//  交易日期
	TradeDate types.TThostFtdcTradeDateType
	//  交易时间
	TradeTime types.TThostFtdcTradeTimeType
	//  银行流水号
	BankSerial types.TThostFtdcBankSerialType
	//  交易系统日期
	TradingDay types.TThostFtdcTradeDateType
	//  银期平台消息流水号
	PlateSerial types.TThostFtdcSerialType
	//  最后分片标志
	LastFragment types.TThostFtdcLastFragmentType
	//  会话号
	SessionID types.TThostFtdcSessionIDType
	//  安装编号
	InstallID types.TThostFtdcInstallIDType
	//  用户标识
	UserID types.TThostFtdcUserIDType
	//  摘要
	Digest types.TThostFtdcDigestType
	//  币种代码
	CurrencyID types.TThostFtdcCurrencyIDType
	//  渠道标志
	DeviceID types.TThostFtdcDeviceIDType
	//  期货公司银行编码
	BrokerIDByBank types.TThostFtdcBankCodingForFutureType
	//  交易柜员
	OperNo types.TThostFtdcOperNoType
	//  请求编号
	RequestID types.TThostFtdcRequestIDType
	//  交易ID
	TID types.TThostFtdcTIDType
}

// 期商签退响应
type CThostFtdcRspFutureSignOutField struct {
	//  业务功能码
	TradeCode types.TThostFtdcTradeCodeType
	//  银行代码
	BankID types.TThostFtdcBankIDType
	//  银行分支机构代码
	BankBranchID types.TThostFtdcBankBrchIDType
	//  期商代码
	BrokerID types.TThostFtdcBrokerIDType
	//  期商分支机构代码
	BrokerBranchID types.TThostFtdcFutureBranchIDType
	//  交易日期
	TradeDate types.TThostFtdcTradeDateType
	//  交易时间
	TradeTime types.TThostFtdcTradeTimeType
	//  银行流水号
	BankSerial types.TThostFtdcBankSerialType
	//  交易系统日期
	TradingDay types.TThostFtdcTradeDateType
	//  银期平台消息流水号
	PlateSerial types.TThostFtdcSerialType
	//  最后分片标志
	LastFragment types.TThostFtdcLastFragmentType
	//  会话号
	SessionID types.TThostFtdcSessionIDType
	//  安装编号
	InstallID types.TThostFtdcInstallIDType
	//  用户标识
	UserID types.TThostFtdcUserIDType
	//  摘要
	Digest types.TThostFtdcDigestType
	//  币种代码
	CurrencyID types.TThostFtdcCurrencyIDType
	//  渠道标志
	DeviceID types.TThostFtdcDeviceIDType
	//  期货公司银行编码
	BrokerIDByBank types.TThostFtdcBankCodingForFutureType
	//  交易柜员
	OperNo types.TThostFtdcOperNoType
	//  请求编号
	RequestID types.TThostFtdcRequestIDType
	//  交易ID
	TID types.TThostFtdcTIDType
	//  错误代码
	ErrorID types.TThostFtdcErrorIDType
	//  错误信息
	ErrorMsg types.TThostFtdcErrorMsgType
}

// 查询指定流水号的交易结果请求
type CThostFtdcReqQueryTradeResultBySerialField struct {
	//  业务功能码
	TradeCode types.TThostFtdcTradeCodeType
	//  银行代码
	BankID types.TThostFtdcBankIDType
	//  银行分支机构代码
	BankBranchID types.TThostFtdcBankBrchIDType
	//  期商代码
	BrokerID types.TThostFtdcBrokerIDType
	//  期商分支机构代码
	BrokerBranchID types.TThostFtdcFutureBranchIDType
	//  交易日期
	TradeDate types.TThostFtdcTradeDateType
	//  交易时间
	TradeTime types.TThostFtdcTradeTimeType
	//  银行流水号
	BankSerial types.TThostFtdcBankSerialType
	//  交易系统日期
	TradingDay types.TThostFtdcTradeDateType
	//  银期平台消息流水号
	PlateSerial types.TThostFtdcSerialType
	//  最后分片标志
	LastFragment types.TThostFtdcLastFragmentType
	//  会话号
	SessionID types.TThostFtdcSessionIDType
	//  流水号
	Reference types.TThostFtdcSerialType
	//  本流水号发布者的机构类型
	RefrenceIssureType types.TThostFtdcInstitutionTypeType
	//  本流水号发布者机构编码
	RefrenceIssure types.TThostFtdcOrganCodeType
	//  客户姓名
	CustomerName types.TThostFtdcIndividualNameType
	//  证件类型
	IdCardType types.TThostFtdcIdCardTypeType
	//  证件号码
	IdentifiedCardNo types.TThostFtdcIdentifiedCardNoType
	//  客户类型
	CustType types.TThostFtdcCustTypeType
	//  银行帐号
	BankAccount types.TThostFtdcBankAccountType
	//  银行密码
	BankPassWord types.TThostFtdcPasswordType
	//  投资者帐号
	AccountID types.TThostFtdcAccountIDType
	//  期货密码
	Password types.TThostFtdcPasswordType
	//  币种代码
	CurrencyID types.TThostFtdcCurrencyIDType
	//  转帐金额
	TradeAmount types.TThostFtdcTradeAmountType
	//  摘要
	Digest types.TThostFtdcDigestType
	//  长客户姓名
	LongCustomerName types.TThostFtdcLongIndividualNameType
}

// 查询指定流水号的交易结果响应
type CThostFtdcRspQueryTradeResultBySerialField struct {
	//  业务功能码
	TradeCode types.TThostFtdcTradeCodeType
	//  银行代码
	BankID types.TThostFtdcBankIDType
	//  银行分支机构代码
	BankBranchID types.TThostFtdcBankBrchIDType
	//  期商代码
	BrokerID types.TThostFtdcBrokerIDType
	//  期商分支机构代码
	BrokerBranchID types.TThostFtdcFutureBranchIDType
	//  交易日期
	TradeDate types.TThostFtdcTradeDateType
	//  交易时间
	TradeTime types.TThostFtdcTradeTimeType
	//  银行流水号
	BankSerial types.TThostFtdcBankSerialType
	//  交易系统日期
	TradingDay types.TThostFtdcTradeDateType
	//  银期平台消息流水号
	PlateSerial types.TThostFtdcSerialType
	//  最后分片标志
	LastFragment types.TThostFtdcLastFragmentType
	//  会话号
	SessionID types.TThostFtdcSessionIDType
	//  错误代码
	ErrorID types.TThostFtdcErrorIDType
	//  错误信息
	ErrorMsg types.TThostFtdcErrorMsgType
	//  流水号
	Reference types.TThostFtdcSerialType
	//  本流水号发布者的机构类型
	RefrenceIssureType types.TThostFtdcInstitutionTypeType
	//  本流水号发布者机构编码
	RefrenceIssure types.TThostFtdcOrganCodeType
	//  原始返回代码
	OriginReturnCode types.TThostFtdcReturnCodeType
	//  原始返回码描述
	OriginDescrInfoForReturnCode types.TThostFtdcDescrInfoForReturnCodeType
	//  银行帐号
	BankAccount types.TThostFtdcBankAccountType
	//  银行密码
	BankPassWord types.TThostFtdcPasswordType
	//  投资者帐号
	AccountID types.TThostFtdcAccountIDType
	//  期货密码
	Password types.TThostFtdcPasswordType
	//  币种代码
	CurrencyID types.TThostFtdcCurrencyIDType
	//  转帐金额
	TradeAmount types.TThostFtdcTradeAmountType
	//  摘要
	Digest types.TThostFtdcDigestType
}

// 日终文件就绪请求
type CThostFtdcReqDayEndFileReadyField struct {
	//  业务功能码
	TradeCode types.TThostFtdcTradeCodeType
	//  银行代码
	BankID types.TThostFtdcBankIDType
	//  银行分支机构代码
	BankBranchID types.TThostFtdcBankBrchIDType
	//  期商代码
	BrokerID types.TThostFtdcBrokerIDType
	//  期商分支机构代码
	BrokerBranchID types.TThostFtdcFutureBranchIDType
	//  交易日期
	TradeDate types.TThostFtdcTradeDateType
	//  交易时间
	TradeTime types.TThostFtdcTradeTimeType
	//  银行流水号
	BankSerial types.TThostFtdcBankSerialType
	//  交易系统日期
	TradingDay types.TThostFtdcTradeDateType
	//  银期平台消息流水号
	PlateSerial types.TThostFtdcSerialType
	//  最后分片标志
	LastFragment types.TThostFtdcLastFragmentType
	//  会话号
	SessionID types.TThostFtdcSessionIDType
	//  文件业务功能
	FileBusinessCode types.TThostFtdcFileBusinessCodeType
	//  摘要
	Digest types.TThostFtdcDigestType
}

// 返回结果
type CThostFtdcReturnResultField struct {
	//  返回代码
	ReturnCode types.TThostFtdcReturnCodeType
	//  返回码描述
	DescrInfoForReturnCode types.TThostFtdcDescrInfoForReturnCodeType
}

// 验证期货资金密码
type CThostFtdcVerifyFuturePasswordField struct {
	//  业务功能码
	TradeCode types.TThostFtdcTradeCodeType
	//  银行代码
	BankID types.TThostFtdcBankIDType
	//  银行分支机构代码
	BankBranchID types.TThostFtdcBankBrchIDType
	//  期商代码
	BrokerID types.TThostFtdcBrokerIDType
	//  期商分支机构代码
	BrokerBranchID types.TThostFtdcFutureBranchIDType
	//  交易日期
	TradeDate types.TThostFtdcTradeDateType
	//  交易时间
	TradeTime types.TThostFtdcTradeTimeType
	//  银行流水号
	BankSerial types.TThostFtdcBankSerialType
	//  交易系统日期
	TradingDay types.TThostFtdcTradeDateType
	//  银期平台消息流水号
	PlateSerial types.TThostFtdcSerialType
	//  最后分片标志
	LastFragment types.TThostFtdcLastFragmentType
	//  会话号
	SessionID types.TThostFtdcSessionIDType
	//  投资者帐号
	AccountID types.TThostFtdcAccountIDType
	//  期货密码
	Password types.TThostFtdcPasswordType
	//  银行帐号
	BankAccount types.TThostFtdcBankAccountType
	//  银行密码
	BankPassWord types.TThostFtdcPasswordType
	//  安装编号
	InstallID types.TThostFtdcInstallIDType
	//  交易ID
	TID types.TThostFtdcTIDType
	//  币种代码
	CurrencyID types.TThostFtdcCurrencyIDType
}

// 验证客户信息
type CThostFtdcVerifyCustInfoField struct {
	//  客户姓名
	CustomerName types.TThostFtdcIndividualNameType
	//  证件类型
	IdCardType types.TThostFtdcIdCardTypeType
	//  证件号码
	IdentifiedCardNo types.TThostFtdcIdentifiedCardNoType
	//  客户类型
	CustType types.TThostFtdcCustTypeType
	//  长客户姓名
	LongCustomerName types.TThostFtdcLongIndividualNameType
}

// 验证期货资金密码和客户信息
type CThostFtdcVerifyFuturePasswordAndCustInfoField struct {
	//  客户姓名
	CustomerName types.TThostFtdcIndividualNameType
	//  证件类型
	IdCardType types.TThostFtdcIdCardTypeType
	//  证件号码
	IdentifiedCardNo types.TThostFtdcIdentifiedCardNoType
	//  客户类型
	CustType types.TThostFtdcCustTypeType
	//  投资者帐号
	AccountID types.TThostFtdcAccountIDType
	//  期货密码
	Password types.TThostFtdcPasswordType
	//  币种代码
	CurrencyID types.TThostFtdcCurrencyIDType
	//  长客户姓名
	LongCustomerName types.TThostFtdcLongIndividualNameType
}

// 验证期货资金密码和客户信息
type CThostFtdcDepositResultInformField struct {
	//  出入金流水号，该流水号为银期报盘返回的流水号
	DepositSeqNo types.TThostFtdcDepositSeqNoType
	//  经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	//  投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	//  入金金额
	Deposit types.TThostFtdcMoneyType
	//  请求编号
	RequestID types.TThostFtdcRequestIDType
	//  返回代码
	ReturnCode types.TThostFtdcReturnCodeType
	//  返回码描述
	DescrInfoForReturnCode types.TThostFtdcDescrInfoForReturnCodeType
}

// 交易核心向银期报盘发出密钥同步请求
type CThostFtdcReqSyncKeyField struct {
	//  业务功能码
	TradeCode types.TThostFtdcTradeCodeType
	//  银行代码
	BankID types.TThostFtdcBankIDType
	//  银行分支机构代码
	BankBranchID types.TThostFtdcBankBrchIDType
	//  期商代码
	BrokerID types.TThostFtdcBrokerIDType
	//  期商分支机构代码
	BrokerBranchID types.TThostFtdcFutureBranchIDType
	//  交易日期
	TradeDate types.TThostFtdcTradeDateType
	//  交易时间
	TradeTime types.TThostFtdcTradeTimeType
	//  银行流水号
	BankSerial types.TThostFtdcBankSerialType
	//  交易系统日期
	TradingDay types.TThostFtdcTradeDateType
	//  银期平台消息流水号
	PlateSerial types.TThostFtdcSerialType
	//  最后分片标志
	LastFragment types.TThostFtdcLastFragmentType
	//  会话号
	SessionID types.TThostFtdcSessionIDType
	//  安装编号
	InstallID types.TThostFtdcInstallIDType
	//  用户标识
	UserID types.TThostFtdcUserIDType
	//  交易核心给银期报盘的消息
	Message types.TThostFtdcAddInfoType
	//  渠道标志
	DeviceID types.TThostFtdcDeviceIDType
	//  期货公司银行编码
	BrokerIDByBank types.TThostFtdcBankCodingForFutureType
	//  交易柜员
	OperNo types.TThostFtdcOperNoType
	//  请求编号
	RequestID types.TThostFtdcRequestIDType
	//  交易ID
	TID types.TThostFtdcTIDType
}

// 交易核心向银期报盘发出密钥同步响应
type CThostFtdcRspSyncKeyField struct {
	//  业务功能码
	TradeCode types.TThostFtdcTradeCodeType
	//  银行代码
	BankID types.TThostFtdcBankIDType
	//  银行分支机构代码
	BankBranchID types.TThostFtdcBankBrchIDType
	//  期商代码
	BrokerID types.TThostFtdcBrokerIDType
	//  期商分支机构代码
	BrokerBranchID types.TThostFtdcFutureBranchIDType
	//  交易日期
	TradeDate types.TThostFtdcTradeDateType
	//  交易时间
	TradeTime types.TThostFtdcTradeTimeType
	//  银行流水号
	BankSerial types.TThostFtdcBankSerialType
	//  交易系统日期
	TradingDay types.TThostFtdcTradeDateType
	//  银期平台消息流水号
	PlateSerial types.TThostFtdcSerialType
	//  最后分片标志
	LastFragment types.TThostFtdcLastFragmentType
	//  会话号
	SessionID types.TThostFtdcSessionIDType
	//  安装编号
	InstallID types.TThostFtdcInstallIDType
	//  用户标识
	UserID types.TThostFtdcUserIDType
	//  交易核心给银期报盘的消息
	Message types.TThostFtdcAddInfoType
	//  渠道标志
	DeviceID types.TThostFtdcDeviceIDType
	//  期货公司银行编码
	BrokerIDByBank types.TThostFtdcBankCodingForFutureType
	//  交易柜员
	OperNo types.TThostFtdcOperNoType
	//  请求编号
	RequestID types.TThostFtdcRequestIDType
	//  交易ID
	TID types.TThostFtdcTIDType
	//  错误代码
	ErrorID types.TThostFtdcErrorIDType
	//  错误信息
	ErrorMsg types.TThostFtdcErrorMsgType
}

// 查询账户信息通知
type CThostFtdcNotifyQueryAccountField struct {
	//  业务功能码
	TradeCode types.TThostFtdcTradeCodeType
	//  银行代码
	BankID types.TThostFtdcBankIDType
	//  银行分支机构代码
	BankBranchID types.TThostFtdcBankBrchIDType
	//  期商代码
	BrokerID types.TThostFtdcBrokerIDType
	//  期商分支机构代码
	BrokerBranchID types.TThostFtdcFutureBranchIDType
	//  交易日期
	TradeDate types.TThostFtdcTradeDateType
	//  交易时间
	TradeTime types.TThostFtdcTradeTimeType
	//  银行流水号
	BankSerial types.TThostFtdcBankSerialType
	//  交易系统日期
	TradingDay types.TThostFtdcTradeDateType
	//  银期平台消息流水号
	PlateSerial types.TThostFtdcSerialType
	//  最后分片标志
	LastFragment types.TThostFtdcLastFragmentType
	//  会话号
	SessionID types.TThostFtdcSessionIDType
	//  客户姓名
	CustomerName types.TThostFtdcIndividualNameType
	//  证件类型
	IdCardType types.TThostFtdcIdCardTypeType
	//  证件号码
	IdentifiedCardNo types.TThostFtdcIdentifiedCardNoType
	//  客户类型
	CustType types.TThostFtdcCustTypeType
	//  银行帐号
	BankAccount types.TThostFtdcBankAccountType
	//  银行密码
	BankPassWord types.TThostFtdcPasswordType
	//  投资者帐号
	AccountID types.TThostFtdcAccountIDType
	//  期货密码
	Password types.TThostFtdcPasswordType
	//  期货公司流水号
	FutureSerial types.TThostFtdcFutureSerialType
	//  安装编号
	InstallID types.TThostFtdcInstallIDType
	//  用户标识
	UserID types.TThostFtdcUserIDType
	//  验证客户证件号码标志
	VerifyCertNoFlag types.TThostFtdcYesNoIndicatorType
	//  币种代码
	CurrencyID types.TThostFtdcCurrencyIDType
	//  摘要
	Digest types.TThostFtdcDigestType
	//  银行帐号类型
	BankAccType types.TThostFtdcBankAccTypeType
	//  渠道标志
	DeviceID types.TThostFtdcDeviceIDType
	//  期货单位帐号类型
	BankSecuAccType types.TThostFtdcBankAccTypeType
	//  期货公司银行编码
	BrokerIDByBank types.TThostFtdcBankCodingForFutureType
	//  期货单位帐号
	BankSecuAcc types.TThostFtdcBankAccountType
	//  银行密码标志
	BankPwdFlag types.TThostFtdcPwdFlagType
	//  期货资金密码核对标志
	SecuPwdFlag types.TThostFtdcPwdFlagType
	//  交易柜员
	OperNo types.TThostFtdcOperNoType
	//  请求编号
	RequestID types.TThostFtdcRequestIDType
	//  交易ID
	TID types.TThostFtdcTIDType
	//  银行可用金额
	BankUseAmount types.TThostFtdcTradeAmountType
	//  银行可取金额
	BankFetchAmount types.TThostFtdcTradeAmountType
	//  错误代码
	ErrorID types.TThostFtdcErrorIDType
	//  错误信息
	ErrorMsg types.TThostFtdcErrorMsgType
	//  长客户姓名
	LongCustomerName types.TThostFtdcLongIndividualNameType
}

// 银期转账交易流水表
type CThostFtdcTransferSerialField struct {
	//  平台流水号
	PlateSerial types.TThostFtdcPlateSerialType
	//  交易发起方日期
	TradeDate types.TThostFtdcTradeDateType
	//  交易日期
	TradingDay types.TThostFtdcDateType
	//  交易时间
	TradeTime types.TThostFtdcTradeTimeType
	//  交易代码
	TradeCode types.TThostFtdcTradeCodeType
	//  会话编号
	SessionID types.TThostFtdcSessionIDType
	//  银行编码
	BankID types.TThostFtdcBankIDType
	//  银行分支机构编码
	BankBranchID types.TThostFtdcBankBrchIDType
	//  银行帐号类型
	BankAccType types.TThostFtdcBankAccTypeType
	//  银行帐号
	BankAccount types.TThostFtdcBankAccountType
	//  银行流水号
	BankSerial types.TThostFtdcBankSerialType
	//  期货公司编码
	BrokerID types.TThostFtdcBrokerIDType
	//  期商分支机构代码
	BrokerBranchID types.TThostFtdcFutureBranchIDType
	//  期货公司帐号类型
	FutureAccType types.TThostFtdcFutureAccTypeType
	//  投资者帐号
	AccountID types.TThostFtdcAccountIDType
	//  投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	//  期货公司流水号
	FutureSerial types.TThostFtdcFutureSerialType
	//  证件类型
	IdCardType types.TThostFtdcIdCardTypeType
	//  证件号码
	IdentifiedCardNo types.TThostFtdcIdentifiedCardNoType
	//  币种代码
	CurrencyID types.TThostFtdcCurrencyIDType
	//  交易金额
	TradeAmount types.TThostFtdcTradeAmountType
	//  应收客户费用
	CustFee types.TThostFtdcCustFeeType
	//  应收期货公司费用
	BrokerFee types.TThostFtdcFutureFeeType
	//  有效标志
	AvailabilityFlag types.TThostFtdcAvailabilityFlagType
	//  操作员
	OperatorCode types.TThostFtdcOperatorCodeType
	//  新银行帐号
	BankNewAccount types.TThostFtdcBankAccountType
	//  错误代码
	ErrorID types.TThostFtdcErrorIDType
	//  错误信息
	ErrorMsg types.TThostFtdcErrorMsgType
}

// 请求查询转帐流水
type CThostFtdcQryTransferSerialField struct {
	//  经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	//  投资者帐号
	AccountID types.TThostFtdcAccountIDType
	//  银行编码
	BankID types.TThostFtdcBankIDType
	//  币种代码
	CurrencyID types.TThostFtdcCurrencyIDType
}

// 期商签到通知
type CThostFtdcNotifyFutureSignInField struct {
	//  业务功能码
	TradeCode types.TThostFtdcTradeCodeType
	//  银行代码
	BankID types.TThostFtdcBankIDType
	//  银行分支机构代码
	BankBranchID types.TThostFtdcBankBrchIDType
	//  期商代码
	BrokerID types.TThostFtdcBrokerIDType
	//  期商分支机构代码
	BrokerBranchID types.TThostFtdcFutureBranchIDType
	//  交易日期
	TradeDate types.TThostFtdcTradeDateType
	//  交易时间
	TradeTime types.TThostFtdcTradeTimeType
	//  银行流水号
	BankSerial types.TThostFtdcBankSerialType
	//  交易系统日期
	TradingDay types.TThostFtdcTradeDateType
	//  银期平台消息流水号
	PlateSerial types.TThostFtdcSerialType
	//  最后分片标志
	LastFragment types.TThostFtdcLastFragmentType
	//  会话号
	SessionID types.TThostFtdcSessionIDType
	//  安装编号
	InstallID types.TThostFtdcInstallIDType
	//  用户标识
	UserID types.TThostFtdcUserIDType
	//  摘要
	Digest types.TThostFtdcDigestType
	//  币种代码
	CurrencyID types.TThostFtdcCurrencyIDType
	//  渠道标志
	DeviceID types.TThostFtdcDeviceIDType
	//  期货公司银行编码
	BrokerIDByBank types.TThostFtdcBankCodingForFutureType
	//  交易柜员
	OperNo types.TThostFtdcOperNoType
	//  请求编号
	RequestID types.TThostFtdcRequestIDType
	//  交易ID
	TID types.TThostFtdcTIDType
	//  错误代码
	ErrorID types.TThostFtdcErrorIDType
	//  错误信息
	ErrorMsg types.TThostFtdcErrorMsgType
	//  PIN密钥
	PinKey types.TThostFtdcPasswordKeyType
	//  MAC密钥
	MacKey types.TThostFtdcPasswordKeyType
}

// 期商签退通知
type CThostFtdcNotifyFutureSignOutField struct {
	//  业务功能码
	TradeCode types.TThostFtdcTradeCodeType
	//  银行代码
	BankID types.TThostFtdcBankIDType
	//  银行分支机构代码
	BankBranchID types.TThostFtdcBankBrchIDType
	//  期商代码
	BrokerID types.TThostFtdcBrokerIDType
	//  期商分支机构代码
	BrokerBranchID types.TThostFtdcFutureBranchIDType
	//  交易日期
	TradeDate types.TThostFtdcTradeDateType
	//  交易时间
	TradeTime types.TThostFtdcTradeTimeType
	//  银行流水号
	BankSerial types.TThostFtdcBankSerialType
	//  交易系统日期
	TradingDay types.TThostFtdcTradeDateType
	//  银期平台消息流水号
	PlateSerial types.TThostFtdcSerialType
	//  最后分片标志
	LastFragment types.TThostFtdcLastFragmentType
	//  会话号
	SessionID types.TThostFtdcSessionIDType
	//  安装编号
	InstallID types.TThostFtdcInstallIDType
	//  用户标识
	UserID types.TThostFtdcUserIDType
	//  摘要
	Digest types.TThostFtdcDigestType
	//  币种代码
	CurrencyID types.TThostFtdcCurrencyIDType
	//  渠道标志
	DeviceID types.TThostFtdcDeviceIDType
	//  期货公司银行编码
	BrokerIDByBank types.TThostFtdcBankCodingForFutureType
	//  交易柜员
	OperNo types.TThostFtdcOperNoType
	//  请求编号
	RequestID types.TThostFtdcRequestIDType
	//  交易ID
	TID types.TThostFtdcTIDType
	//  错误代码
	ErrorID types.TThostFtdcErrorIDType
	//  错误信息
	ErrorMsg types.TThostFtdcErrorMsgType
}

// 交易核心向银期报盘发出密钥同步处理结果的通知
type CThostFtdcNotifySyncKeyField struct {
	//  业务功能码
	TradeCode types.TThostFtdcTradeCodeType
	//  银行代码
	BankID types.TThostFtdcBankIDType
	//  银行分支机构代码
	BankBranchID types.TThostFtdcBankBrchIDType
	//  期商代码
	BrokerID types.TThostFtdcBrokerIDType
	//  期商分支机构代码
	BrokerBranchID types.TThostFtdcFutureBranchIDType
	//  交易日期
	TradeDate types.TThostFtdcTradeDateType
	//  交易时间
	TradeTime types.TThostFtdcTradeTimeType
	//  银行流水号
	BankSerial types.TThostFtdcBankSerialType
	//  交易系统日期
	TradingDay types.TThostFtdcTradeDateType
	//  银期平台消息流水号
	PlateSerial types.TThostFtdcSerialType
	//  最后分片标志
	LastFragment types.TThostFtdcLastFragmentType
	//  会话号
	SessionID types.TThostFtdcSessionIDType
	//  安装编号
	InstallID types.TThostFtdcInstallIDType
	//  用户标识
	UserID types.TThostFtdcUserIDType
	//  交易核心给银期报盘的消息
	Message types.TThostFtdcAddInfoType
	//  渠道标志
	DeviceID types.TThostFtdcDeviceIDType
	//  期货公司银行编码
	BrokerIDByBank types.TThostFtdcBankCodingForFutureType
	//  交易柜员
	OperNo types.TThostFtdcOperNoType
	//  请求编号
	RequestID types.TThostFtdcRequestIDType
	//  交易ID
	TID types.TThostFtdcTIDType
	//  错误代码
	ErrorID types.TThostFtdcErrorIDType
	//  错误信息
	ErrorMsg types.TThostFtdcErrorMsgType
}

// 请求查询银期签约关系
type CThostFtdcQryAccountregisterField struct {
	//  经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	//  投资者帐号
	AccountID types.TThostFtdcAccountIDType
	//  银行编码
	BankID types.TThostFtdcBankIDType
	//  银行分支机构编码
	BankBranchID types.TThostFtdcBankBrchIDType
	//  币种代码
	CurrencyID types.TThostFtdcCurrencyIDType
}

// 客户开销户信息表
type CThostFtdcAccountregisterField struct {
	//  交易日期
	TradeDay types.TThostFtdcTradeDateType
	//  银行编码
	BankID types.TThostFtdcBankIDType
	//  银行分支机构编码
	BankBranchID types.TThostFtdcBankBrchIDType
	//  银行帐号
	BankAccount types.TThostFtdcBankAccountType
	//  期货公司编码
	BrokerID types.TThostFtdcBrokerIDType
	//  期货公司分支机构编码
	BrokerBranchID types.TThostFtdcFutureBranchIDType
	//  投资者帐号
	AccountID types.TThostFtdcAccountIDType
	//  证件类型
	IdCardType types.TThostFtdcIdCardTypeType
	//  证件号码
	IdentifiedCardNo types.TThostFtdcIdentifiedCardNoType
	//  客户姓名
	CustomerName types.TThostFtdcIndividualNameType
	//  币种代码
	CurrencyID types.TThostFtdcCurrencyIDType
	//  开销户类别
	OpenOrDestroy types.TThostFtdcOpenOrDestroyType
	//  签约日期
	RegDate types.TThostFtdcTradeDateType
	//  解约日期
	OutDate types.TThostFtdcTradeDateType
	//  交易ID
	TID types.TThostFtdcTIDType
	//  客户类型
	CustType types.TThostFtdcCustTypeType
	//  银行帐号类型
	BankAccType types.TThostFtdcBankAccTypeType
	//  长客户姓名
	LongCustomerName types.TThostFtdcLongIndividualNameType
}

// 银期开户信息
type CThostFtdcOpenAccountField struct {
	//  业务功能码
	TradeCode types.TThostFtdcTradeCodeType
	//  银行代码
	BankID types.TThostFtdcBankIDType
	//  银行分支机构代码
	BankBranchID types.TThostFtdcBankBrchIDType
	//  期商代码
	BrokerID types.TThostFtdcBrokerIDType
	//  期商分支机构代码
	BrokerBranchID types.TThostFtdcFutureBranchIDType
	//  交易日期
	TradeDate types.TThostFtdcTradeDateType
	//  交易时间
	TradeTime types.TThostFtdcTradeTimeType
	//  银行流水号
	BankSerial types.TThostFtdcBankSerialType
	//  交易系统日期
	TradingDay types.TThostFtdcTradeDateType
	//  银期平台消息流水号
	PlateSerial types.TThostFtdcSerialType
	//  最后分片标志
	LastFragment types.TThostFtdcLastFragmentType
	//  会话号
	SessionID types.TThostFtdcSessionIDType
	//  客户姓名
	CustomerName types.TThostFtdcIndividualNameType
	//  证件类型
	IdCardType types.TThostFtdcIdCardTypeType
	//  证件号码
	IdentifiedCardNo types.TThostFtdcIdentifiedCardNoType
	//  性别
	Gender types.TThostFtdcGenderType
	//  国家代码
	CountryCode types.TThostFtdcCountryCodeType
	//  客户类型
	CustType types.TThostFtdcCustTypeType
	//  地址
	Address types.TThostFtdcAddressType
	//  邮编
	ZipCode types.TThostFtdcZipCodeType
	//  电话号码
	Telephone types.TThostFtdcTelephoneType
	//  手机
	MobilePhone types.TThostFtdcMobilePhoneType
	//  传真
	Fax types.TThostFtdcFaxType
	//  电子邮件
	EMail types.TThostFtdcEMailType
	//  资金账户状态
	MoneyAccountStatus types.TThostFtdcMoneyAccountStatusType
	//  银行帐号
	BankAccount types.TThostFtdcBankAccountType
	//  银行密码
	BankPassWord types.TThostFtdcPasswordType
	//  投资者帐号
	AccountID types.TThostFtdcAccountIDType
	//  期货密码
	Password types.TThostFtdcPasswordType
	//  安装编号
	InstallID types.TThostFtdcInstallIDType
	//  验证客户证件号码标志
	VerifyCertNoFlag types.TThostFtdcYesNoIndicatorType
	//  币种代码
	CurrencyID types.TThostFtdcCurrencyIDType
	//  汇钞标志
	CashExchangeCode types.TThostFtdcCashExchangeCodeType
	//  摘要
	Digest types.TThostFtdcDigestType
	//  银行帐号类型
	BankAccType types.TThostFtdcBankAccTypeType
	//  渠道标志
	DeviceID types.TThostFtdcDeviceIDType
	//  期货单位帐号类型
	BankSecuAccType types.TThostFtdcBankAccTypeType
	//  期货公司银行编码
	BrokerIDByBank types.TThostFtdcBankCodingForFutureType
	//  期货单位帐号
	BankSecuAcc types.TThostFtdcBankAccountType
	//  银行密码标志
	BankPwdFlag types.TThostFtdcPwdFlagType
	//  期货资金密码核对标志
	SecuPwdFlag types.TThostFtdcPwdFlagType
	//  交易柜员
	OperNo types.TThostFtdcOperNoType
	//  交易ID
	TID types.TThostFtdcTIDType
	//  用户标识
	UserID types.TThostFtdcUserIDType
	//  错误代码
	ErrorID types.TThostFtdcErrorIDType
	//  错误信息
	ErrorMsg types.TThostFtdcErrorMsgType
	//  长客户姓名
	LongCustomerName types.TThostFtdcLongIndividualNameType
}

// 银期销户信息
type CThostFtdcCancelAccountField struct {
	//  业务功能码
	TradeCode types.TThostFtdcTradeCodeType
	//  银行代码
	BankID types.TThostFtdcBankIDType
	//  银行分支机构代码
	BankBranchID types.TThostFtdcBankBrchIDType
	//  期商代码
	BrokerID types.TThostFtdcBrokerIDType
	//  期商分支机构代码
	BrokerBranchID types.TThostFtdcFutureBranchIDType
	//  交易日期
	TradeDate types.TThostFtdcTradeDateType
	//  交易时间
	TradeTime types.TThostFtdcTradeTimeType
	//  银行流水号
	BankSerial types.TThostFtdcBankSerialType
	//  交易系统日期
	TradingDay types.TThostFtdcTradeDateType
	//  银期平台消息流水号
	PlateSerial types.TThostFtdcSerialType
	//  最后分片标志
	LastFragment types.TThostFtdcLastFragmentType
	//  会话号
	SessionID types.TThostFtdcSessionIDType
	//  客户姓名
	CustomerName types.TThostFtdcIndividualNameType
	//  证件类型
	IdCardType types.TThostFtdcIdCardTypeType
	//  证件号码
	IdentifiedCardNo types.TThostFtdcIdentifiedCardNoType
	//  性别
	Gender types.TThostFtdcGenderType
	//  国家代码
	CountryCode types.TThostFtdcCountryCodeType
	//  客户类型
	CustType types.TThostFtdcCustTypeType
	//  地址
	Address types.TThostFtdcAddressType
	//  邮编
	ZipCode types.TThostFtdcZipCodeType
	//  电话号码
	Telephone types.TThostFtdcTelephoneType
	//  手机
	MobilePhone types.TThostFtdcMobilePhoneType
	//  传真
	Fax types.TThostFtdcFaxType
	//  电子邮件
	EMail types.TThostFtdcEMailType
	//  资金账户状态
	MoneyAccountStatus types.TThostFtdcMoneyAccountStatusType
	//  银行帐号
	BankAccount types.TThostFtdcBankAccountType
	//  银行密码
	BankPassWord types.TThostFtdcPasswordType
	//  投资者帐号
	AccountID types.TThostFtdcAccountIDType
	//  期货密码
	Password types.TThostFtdcPasswordType
	//  安装编号
	InstallID types.TThostFtdcInstallIDType
	//  验证客户证件号码标志
	VerifyCertNoFlag types.TThostFtdcYesNoIndicatorType
	//  币种代码
	CurrencyID types.TThostFtdcCurrencyIDType
	//  汇钞标志
	CashExchangeCode types.TThostFtdcCashExchangeCodeType
	//  摘要
	Digest types.TThostFtdcDigestType
	//  银行帐号类型
	BankAccType types.TThostFtdcBankAccTypeType
	//  渠道标志
	DeviceID types.TThostFtdcDeviceIDType
	//  期货单位帐号类型
	BankSecuAccType types.TThostFtdcBankAccTypeType
	//  期货公司银行编码
	BrokerIDByBank types.TThostFtdcBankCodingForFutureType
	//  期货单位帐号
	BankSecuAcc types.TThostFtdcBankAccountType
	//  银行密码标志
	BankPwdFlag types.TThostFtdcPwdFlagType
	//  期货资金密码核对标志
	SecuPwdFlag types.TThostFtdcPwdFlagType
	//  交易柜员
	OperNo types.TThostFtdcOperNoType
	//  交易ID
	TID types.TThostFtdcTIDType
	//  用户标识
	UserID types.TThostFtdcUserIDType
	//  错误代码
	ErrorID types.TThostFtdcErrorIDType
	//  错误信息
	ErrorMsg types.TThostFtdcErrorMsgType
	//  长客户姓名
	LongCustomerName types.TThostFtdcLongIndividualNameType
}

// 银期变更银行账号信息
type CThostFtdcChangeAccountField struct {
	//  业务功能码
	TradeCode types.TThostFtdcTradeCodeType
	//  银行代码
	BankID types.TThostFtdcBankIDType
	//  银行分支机构代码
	BankBranchID types.TThostFtdcBankBrchIDType
	//  期商代码
	BrokerID types.TThostFtdcBrokerIDType
	//  期商分支机构代码
	BrokerBranchID types.TThostFtdcFutureBranchIDType
	//  交易日期
	TradeDate types.TThostFtdcTradeDateType
	//  交易时间
	TradeTime types.TThostFtdcTradeTimeType
	//  银行流水号
	BankSerial types.TThostFtdcBankSerialType
	//  交易系统日期
	TradingDay types.TThostFtdcTradeDateType
	//  银期平台消息流水号
	PlateSerial types.TThostFtdcSerialType
	//  最后分片标志
	LastFragment types.TThostFtdcLastFragmentType
	//  会话号
	SessionID types.TThostFtdcSessionIDType
	//  客户姓名
	CustomerName types.TThostFtdcIndividualNameType
	//  证件类型
	IdCardType types.TThostFtdcIdCardTypeType
	//  证件号码
	IdentifiedCardNo types.TThostFtdcIdentifiedCardNoType
	//  性别
	Gender types.TThostFtdcGenderType
	//  国家代码
	CountryCode types.TThostFtdcCountryCodeType
	//  客户类型
	CustType types.TThostFtdcCustTypeType
	//  地址
	Address types.TThostFtdcAddressType
	//  邮编
	ZipCode types.TThostFtdcZipCodeType
	//  电话号码
	Telephone types.TThostFtdcTelephoneType
	//  手机
	MobilePhone types.TThostFtdcMobilePhoneType
	//  传真
	Fax types.TThostFtdcFaxType
	//  电子邮件
	EMail types.TThostFtdcEMailType
	//  资金账户状态
	MoneyAccountStatus types.TThostFtdcMoneyAccountStatusType
	//  银行帐号
	BankAccount types.TThostFtdcBankAccountType
	//  银行密码
	BankPassWord types.TThostFtdcPasswordType
	//  新银行帐号
	NewBankAccount types.TThostFtdcBankAccountType
	//  新银行密码
	NewBankPassWord types.TThostFtdcPasswordType
	//  投资者帐号
	AccountID types.TThostFtdcAccountIDType
	//  期货密码
	Password types.TThostFtdcPasswordType
	//  银行帐号类型
	BankAccType types.TThostFtdcBankAccTypeType
	//  安装编号
	InstallID types.TThostFtdcInstallIDType
	//  验证客户证件号码标志
	VerifyCertNoFlag types.TThostFtdcYesNoIndicatorType
	//  币种代码
	CurrencyID types.TThostFtdcCurrencyIDType
	//  期货公司银行编码
	BrokerIDByBank types.TThostFtdcBankCodingForFutureType
	//  银行密码标志
	BankPwdFlag types.TThostFtdcPwdFlagType
	//  期货资金密码核对标志
	SecuPwdFlag types.TThostFtdcPwdFlagType
	//  交易ID
	TID types.TThostFtdcTIDType
	//  摘要
	Digest types.TThostFtdcDigestType
	//  错误代码
	ErrorID types.TThostFtdcErrorIDType
	//  错误信息
	ErrorMsg types.TThostFtdcErrorMsgType
	//  长客户姓名
	LongCustomerName types.TThostFtdcLongIndividualNameType
}

// 二级代理操作员银期权限
type CThostFtdcSecAgentACIDMapField struct {
	//  经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	//  用户代码
	UserID types.TThostFtdcUserIDType
	//  资金账户
	AccountID types.TThostFtdcAccountIDType
	//  币种
	CurrencyID types.TThostFtdcCurrencyIDType
	//  境外中介机构资金帐号
	BrokerSecAgentID types.TThostFtdcAccountIDType
}

// 二级代理操作员银期权限查询
type CThostFtdcQrySecAgentACIDMapField struct {
	//  经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	//  用户代码
	UserID types.TThostFtdcUserIDType
	//  资金账户
	AccountID types.TThostFtdcAccountIDType
	//  币种
	CurrencyID types.TThostFtdcCurrencyIDType
}

// 灾备中心交易权限
type CThostFtdcUserRightsAssignField struct {
	//  应用单元代码
	BrokerID types.TThostFtdcBrokerIDType
	//  用户代码
	UserID types.TThostFtdcUserIDType
	//  交易中心代码
	DRIdentityID types.TThostFtdcDRIdentityIDType
}

// 经济公司是否有在本标示的交易权限
type CThostFtdcBrokerUserRightAssignField struct {
	//  应用单元代码
	BrokerID types.TThostFtdcBrokerIDType
	//  交易中心代码
	DRIdentityID types.TThostFtdcDRIdentityIDType
	//  能否交易
	Tradeable types.TThostFtdcBoolType
}

// 灾备交易转换报文
type CThostFtdcDRTransferField struct {
	//  原交易中心代码
	OrigDRIdentityID types.TThostFtdcDRIdentityIDType
	//  目标交易中心代码
	DestDRIdentityID types.TThostFtdcDRIdentityIDType
	//  原应用单元代码
	OrigBrokerID types.TThostFtdcBrokerIDType
	//  目标易用单元代码
	DestBrokerID types.TThostFtdcBrokerIDType
}

// Fens用户信息
type CThostFtdcFensUserInfoField struct {
	//  经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	//  用户代码
	UserID types.TThostFtdcUserIDType
	//  登录模式
	LoginMode types.TThostFtdcLoginModeType
}

// 当前银期所属交易中心
type CThostFtdcCurrTransferIdentityField struct {
	//  交易中心代码
	IdentityID types.TThostFtdcDRIdentityIDType
}

// 禁止登录用户
type CThostFtdcLoginForbiddenUserField struct {
	//  经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	//  用户代码
	UserID types.TThostFtdcUserIDType
	//  保留的无效字段
	Reserve1 types.TThostFtdcOldIPAddressType
	//  IP地址
	IPAddress types.TThostFtdcIPAddressType
}

// 查询禁止登录用户
type CThostFtdcQryLoginForbiddenUserField struct {
	//  经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	//  用户代码
	UserID types.TThostFtdcUserIDType
}

// 资金账户基本准备金
type CThostFtdcTradingAccountReserveField struct {
	//  经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	//  投资者帐号
	AccountID types.TThostFtdcAccountIDType
	//  基本准备金
	Reserve types.TThostFtdcMoneyType
	//  币种代码
	CurrencyID types.TThostFtdcCurrencyIDType
}

// 查询禁止登录IP
type CThostFtdcQryLoginForbiddenIPField struct {
	//  保留的无效字段
	Reserve1 types.TThostFtdcOldIPAddressType
	//  IP地址
	IPAddress types.TThostFtdcIPAddressType
}

// 查询IP列表
type CThostFtdcQryIPListField struct {
	//  保留的无效字段
	Reserve1 types.TThostFtdcOldIPAddressType
	//  IP地址
	IPAddress types.TThostFtdcIPAddressType
}

// 查询用户下单权限分配表
type CThostFtdcQryUserRightsAssignField struct {
	//  应用单元代码
	BrokerID types.TThostFtdcBrokerIDType
	//  用户代码
	UserID types.TThostFtdcUserIDType
}

// 银期预约开户确认请求
type CThostFtdcReserveOpenAccountConfirmField struct {
	//  业务功能码
	TradeCode types.TThostFtdcTradeCodeType
	//  银行代码
	BankID types.TThostFtdcBankIDType
	//  银行分支机构代码
	BankBranchID types.TThostFtdcBankBrchIDType
	//  期商代码
	BrokerID types.TThostFtdcBrokerIDType
	//  期商分支机构代码
	BrokerBranchID types.TThostFtdcFutureBranchIDType
	//  交易日期
	TradeDate types.TThostFtdcTradeDateType
	//  交易时间
	TradeTime types.TThostFtdcTradeTimeType
	//  银行流水号
	BankSerial types.TThostFtdcBankSerialType
	//  交易系统日期
	TradingDay types.TThostFtdcTradeDateType
	//  银期平台消息流水号
	PlateSerial types.TThostFtdcSerialType
	//  最后分片标志
	LastFragment types.TThostFtdcLastFragmentType
	//  会话号
	SessionID types.TThostFtdcSessionIDType
	//  客户姓名
	CustomerName types.TThostFtdcLongIndividualNameType
	//  证件类型
	IdCardType types.TThostFtdcIdCardTypeType
	//  证件号码
	IdentifiedCardNo types.TThostFtdcIdentifiedCardNoType
	//  性别
	Gender types.TThostFtdcGenderType
	//  国家代码
	CountryCode types.TThostFtdcCountryCodeType
	//  客户类型
	CustType types.TThostFtdcCustTypeType
	//  地址
	Address types.TThostFtdcAddressType
	//  邮编
	ZipCode types.TThostFtdcZipCodeType
	//  电话号码
	Telephone types.TThostFtdcTelephoneType
	//  手机
	MobilePhone types.TThostFtdcMobilePhoneType
	//  传真
	Fax types.TThostFtdcFaxType
	//  电子邮件
	EMail types.TThostFtdcEMailType
	//  资金账户状态
	MoneyAccountStatus types.TThostFtdcMoneyAccountStatusType
	//  银行帐号
	BankAccount types.TThostFtdcBankAccountType
	//  银行密码
	BankPassWord types.TThostFtdcPasswordType
	//  安装编号
	InstallID types.TThostFtdcInstallIDType
	//  验证客户证件号码标志
	VerifyCertNoFlag types.TThostFtdcYesNoIndicatorType
	//  币种代码
	CurrencyID types.TThostFtdcCurrencyIDType
	//  摘要
	Digest types.TThostFtdcDigestType
	//  银行帐号类型
	BankAccType types.TThostFtdcBankAccTypeType
	//  期货公司银行编码
	BrokerIDByBank types.TThostFtdcBankCodingForFutureType
	//  交易ID
	TID types.TThostFtdcTIDType
	//  投资者帐号
	AccountID types.TThostFtdcAccountIDType
	//  期货密码
	Password types.TThostFtdcPasswordType
	//  预约开户银行流水号
	BankReserveOpenSeq types.TThostFtdcBankSerialType
	//  预约开户日期
	BookDate types.TThostFtdcTradeDateType
	//  预约开户验证密码
	BookPsw types.TThostFtdcPasswordType
	//  错误代码
	ErrorID types.TThostFtdcErrorIDType
	//  错误信息
	ErrorMsg types.TThostFtdcErrorMsgType
}

// 银期预约开户
type CThostFtdcReserveOpenAccountField struct {
	//  业务功能码
	TradeCode types.TThostFtdcTradeCodeType
	//  银行代码
	BankID types.TThostFtdcBankIDType
	//  银行分支机构代码
	BankBranchID types.TThostFtdcBankBrchIDType
	//  期商代码
	BrokerID types.TThostFtdcBrokerIDType
	//  期商分支机构代码
	BrokerBranchID types.TThostFtdcFutureBranchIDType
	//  交易日期
	TradeDate types.TThostFtdcTradeDateType
	//  交易时间
	TradeTime types.TThostFtdcTradeTimeType
	//  银行流水号
	BankSerial types.TThostFtdcBankSerialType
	//  交易系统日期
	TradingDay types.TThostFtdcTradeDateType
	//  银期平台消息流水号
	PlateSerial types.TThostFtdcSerialType
	//  最后分片标志
	LastFragment types.TThostFtdcLastFragmentType
	//  会话号
	SessionID types.TThostFtdcSessionIDType
	//  客户姓名
	CustomerName types.TThostFtdcLongIndividualNameType
	//  证件类型
	IdCardType types.TThostFtdcIdCardTypeType
	//  证件号码
	IdentifiedCardNo types.TThostFtdcIdentifiedCardNoType
	//  性别
	Gender types.TThostFtdcGenderType
	//  国家代码
	CountryCode types.TThostFtdcCountryCodeType
	//  客户类型
	CustType types.TThostFtdcCustTypeType
	//  地址
	Address types.TThostFtdcAddressType
	//  邮编
	ZipCode types.TThostFtdcZipCodeType
	//  电话号码
	Telephone types.TThostFtdcTelephoneType
	//  手机
	MobilePhone types.TThostFtdcMobilePhoneType
	//  传真
	Fax types.TThostFtdcFaxType
	//  电子邮件
	EMail types.TThostFtdcEMailType
	//  资金账户状态
	MoneyAccountStatus types.TThostFtdcMoneyAccountStatusType
	//  银行帐号
	BankAccount types.TThostFtdcBankAccountType
	//  银行密码
	BankPassWord types.TThostFtdcPasswordType
	//  安装编号
	InstallID types.TThostFtdcInstallIDType
	//  验证客户证件号码标志
	VerifyCertNoFlag types.TThostFtdcYesNoIndicatorType
	//  币种代码
	CurrencyID types.TThostFtdcCurrencyIDType
	//  摘要
	Digest types.TThostFtdcDigestType
	//  银行帐号类型
	BankAccType types.TThostFtdcBankAccTypeType
	//  期货公司银行编码
	BrokerIDByBank types.TThostFtdcBankCodingForFutureType
	//  交易ID
	TID types.TThostFtdcTIDType
	//  预约开户状态
	ReserveOpenAccStas types.TThostFtdcReserveOpenAccStasType
	//  错误代码
	ErrorID types.TThostFtdcErrorIDType
	//  错误信息
	ErrorMsg types.TThostFtdcErrorMsgType
}

// 银行账户属性
type CThostFtdcAccountPropertyField struct {
	//  经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	//  投资者帐号
	AccountID types.TThostFtdcAccountIDType
	//  银行统一标识类型
	BankID types.TThostFtdcBankIDType
	//  银行账户
	BankAccount types.TThostFtdcBankAccountType
	//  银行账户的开户人名称
	OpenName types.TThostFtdcInvestorFullNameType
	//  银行账户的开户行
	OpenBank types.TThostFtdcOpenBankType
	//  是否活跃
	IsActive types.TThostFtdcBoolType
	//  账户来源
	AccountSourceType types.TThostFtdcAccountSourceTypeType
	//  开户日期
	OpenDate types.TThostFtdcDateType
	//  注销日期
	CancelDate types.TThostFtdcDateType
	//  录入员代码
	OperatorID types.TThostFtdcOperatorIDType
	//  录入日期
	OperateDate types.TThostFtdcDateType
	//  录入时间
	OperateTime types.TThostFtdcTimeType
	//  币种代码
	CurrencyID types.TThostFtdcCurrencyIDType
}

// 查询当前交易中心
type CThostFtdcQryCurrDRIdentityField struct {
	//  交易中心代码
	DRIdentityID types.TThostFtdcDRIdentityIDType
}

// 当前交易中心
type CThostFtdcCurrDRIdentityField struct {
	//  交易中心代码
	DRIdentityID types.TThostFtdcDRIdentityIDType
}

// 查询二级代理商资金校验模式
type CThostFtdcQrySecAgentCheckModeField struct {
	//  经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	//  投资者代码
	InvestorID types.TThostFtdcInvestorIDType
}

// 查询二级代理商信息
type CThostFtdcQrySecAgentTradeInfoField struct {
	//  经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	//  境外中介机构资金帐号
	BrokerSecAgentID types.TThostFtdcAccountIDType
}

// 用户发出获取安全安全登陆方法请求
type CThostFtdcReqUserAuthMethodField struct {
	//  交易日
	TradingDay types.TThostFtdcDateType
	//  经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	//  用户代码
	UserID types.TThostFtdcUserIDType
}

// 用户发出获取安全安全登陆方法回复
type CThostFtdcRspUserAuthMethodField struct {
	//  当前可以用的认证模式
	UsableAuthMethod types.TThostFtdcCurrentAuthMethodType
}

// 用户发出获取安全安全登陆方法请求
type CThostFtdcReqGenUserCaptchaField struct {
	//  交易日
	TradingDay types.TThostFtdcDateType
	//  经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	//  用户代码
	UserID types.TThostFtdcUserIDType
}

// 生成的图片验证码信息
type CThostFtdcRspGenUserCaptchaField struct {
	//  经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	//  用户代码
	UserID types.TThostFtdcUserIDType
	//  图片信息长度
	CaptchaInfoLen types.TThostFtdcCaptchaInfoLenType
	//  图片信息
	CaptchaInfo types.TThostFtdcCaptchaInfoType
}

// 用户发出获取安全安全登陆方法请求
type CThostFtdcReqGenUserTextField struct {
	//  交易日
	TradingDay types.TThostFtdcDateType
	//  经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	//  用户代码
	UserID types.TThostFtdcUserIDType
}

// 短信验证码生成的回复
type CThostFtdcRspGenUserTextField struct {
	//  短信验证码序号
	UserTextSeq types.TThostFtdcUserTextSeqType
}

// 用户发出带图形验证码的登录请求请求
type CThostFtdcReqUserLoginWithCaptchaField struct {
	//  交易日
	TradingDay types.TThostFtdcDateType
	//  经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	//  用户代码
	UserID types.TThostFtdcUserIDType
	//  密码
	Password types.TThostFtdcPasswordType
	//  用户端产品信息
	UserProductInfo types.TThostFtdcProductInfoType
	//  接口端产品信息
	InterfaceProductInfo types.TThostFtdcProductInfoType
	//  协议信息
	ProtocolInfo types.TThostFtdcProtocolInfoType
	//  Mac地址
	MacAddress types.TThostFtdcMacAddressType
	//  保留的无效字段
	Reserve1 types.TThostFtdcOldIPAddressType
	//  登录备注
	LoginRemark types.TThostFtdcLoginRemarkType
	//  图形验证码的文字内容
	Captcha types.TThostFtdcPasswordType
	//  终端IP端口
	ClientIPPort types.TThostFtdcIPPortType
	//  终端IP地址
	ClientIPAddress types.TThostFtdcIPAddressType
}

// 用户发出带短信验证码的登录请求请求
type CThostFtdcReqUserLoginWithTextField struct {
	//  交易日
	TradingDay types.TThostFtdcDateType
	//  经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	//  用户代码
	UserID types.TThostFtdcUserIDType
	//  密码
	Password types.TThostFtdcPasswordType
	//  用户端产品信息
	UserProductInfo types.TThostFtdcProductInfoType
	//  接口端产品信息
	InterfaceProductInfo types.TThostFtdcProductInfoType
	//  协议信息
	ProtocolInfo types.TThostFtdcProtocolInfoType
	//  Mac地址
	MacAddress types.TThostFtdcMacAddressType
	//  保留的无效字段
	Reserve1 types.TThostFtdcOldIPAddressType
	//  登录备注
	LoginRemark types.TThostFtdcLoginRemarkType
	//  短信验证码文字内容
	Text types.TThostFtdcPasswordType
	//  终端IP端口
	ClientIPPort types.TThostFtdcIPPortType
	//  终端IP地址
	ClientIPAddress types.TThostFtdcIPAddressType
}

// 用户发出带动态验证码的登录请求请求
type CThostFtdcReqUserLoginWithOTPField struct {
	//  交易日
	TradingDay types.TThostFtdcDateType
	//  经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	//  用户代码
	UserID types.TThostFtdcUserIDType
	//  密码
	Password types.TThostFtdcPasswordType
	//  用户端产品信息
	UserProductInfo types.TThostFtdcProductInfoType
	//  接口端产品信息
	InterfaceProductInfo types.TThostFtdcProductInfoType
	//  协议信息
	ProtocolInfo types.TThostFtdcProtocolInfoType
	//  Mac地址
	MacAddress types.TThostFtdcMacAddressType
	//  保留的无效字段
	Reserve1 types.TThostFtdcOldIPAddressType
	//  登录备注
	LoginRemark types.TThostFtdcLoginRemarkType
	//  OTP密码
	OTPPassword types.TThostFtdcPasswordType
	//  终端IP端口
	ClientIPPort types.TThostFtdcIPPortType
	//  终端IP地址
	ClientIPAddress types.TThostFtdcIPAddressType
}

// api握手请求
type CThostFtdcReqApiHandshakeField struct {
	//  api与front通信密钥版本号
	CryptoKeyVersion types.TThostFtdcCryptoKeyVersionType
}

// front发给api的握手回复
type CThostFtdcRspApiHandshakeField struct {
	//  握手回复数据长度
	FrontHandshakeDataLen types.TThostFtdcHandshakeDataLenType
	//  握手回复数据
	FrontHandshakeData types.TThostFtdcHandshakeDataType
	//  API认证是否开启
	IsApiAuthEnabled types.TThostFtdcBoolType
}

// api给front的验证key的请求
type CThostFtdcReqVerifyApiKeyField struct {
	//  握手回复数据长度
	ApiHandshakeDataLen types.TThostFtdcHandshakeDataLenType
	//  握手回复数据
	ApiHandshakeData types.TThostFtdcHandshakeDataType
}

// 操作员组织架构关系
type CThostFtdcDepartmentUserField struct {
	//  经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	//  用户代码
	UserID types.TThostFtdcUserIDType
	//  投资者范围
	InvestorRange types.TThostFtdcDepartmentRangeType
	//  投资者代码
	InvestorID types.TThostFtdcInvestorIDType
}

// 查询频率，每秒查询比数
type CThostFtdcQueryFreqField struct {
	//  查询频率
	QueryFreq types.TThostFtdcQueryFreqType
	//  FTD频率
	FTDPkgFreq types.TThostFtdcQueryFreqType
}

// 禁止认证IP
type CThostFtdcAuthForbiddenIPField struct {
	//  IP地址
	IPAddress types.TThostFtdcIPAddressType
}

// 查询禁止认证IP
type CThostFtdcQryAuthForbiddenIPField struct {
	//  IP地址
	IPAddress types.TThostFtdcIPAddressType
}

// 换汇可提冻结
type CThostFtdcSyncDelaySwapFrozenField struct {
	//  换汇流水号
	DelaySwapSeqNo types.TThostFtdcDepositSeqNoType
	//  经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	//  投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	//  源币种
	FromCurrencyID types.TThostFtdcCurrencyIDType
	//  源剩余换汇额度(可提冻结)
	FromRemainSwap types.TThostFtdcMoneyType
	//  是否手工换汇
	IsManualSwap types.TThostFtdcBoolType
}

// 用户系统信息
type CThostFtdcUserSystemInfoField struct {
	//  经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	//  用户代码
	UserID types.TThostFtdcUserIDType
	//  用户端系统内部信息长度
	ClientSystemInfoLen types.TThostFtdcSystemInfoLenType
	//  用户端系统内部信息
	ClientSystemInfo types.TThostFtdcClientSystemInfoType
	//  保留的无效字段
	Reserve1 types.TThostFtdcOldIPAddressType
	//  终端IP端口
	ClientIPPort types.TThostFtdcIPPortType
	//  登录成功时间
	ClientLoginTime types.TThostFtdcTimeType
	//  App代码
	ClientAppID types.TThostFtdcAppIDType
	//  用户公网IP
	ClientPublicIP types.TThostFtdcIPAddressType
	//  客户登录备注2
	ClientLoginRemark types.TThostFtdcClientLoginRemarkType
	//  客户终端的MAC等标识
	MAC types.TThostFtdcDeviceTagType
}

// 终端用户绑定信息
type CThostFtdcAuthUserIDField struct {
	//  经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	//  App代码
	AppID types.TThostFtdcAppIDType
	//  用户代码
	UserID types.TThostFtdcUserIDType
	//  校验类型
	AuthType types.TThostFtdcAuthTypeType
}

// 用户IP绑定信息
type CThostFtdcAuthIPField struct {
	//  经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	//  App代码
	AppID types.TThostFtdcAppIDType
	//  用户代码
	IPAddress types.TThostFtdcIPAddressType
}

// 查询分类合约
type CThostFtdcQryClassifiedInstrumentField struct {
	//  合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
	//  交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	//  合约在交易所的代码
	ExchangeInstID types.TThostFtdcExchangeInstIDType
	//  产品代码
	ProductID types.TThostFtdcInstrumentIDType
	//  合约交易状态
	TradingType types.TThostFtdcTradingTypeType
	//  合约分类类型
	ClassType types.TThostFtdcClassTypeType
}

// 查询组合优惠比例
type CThostFtdcQryCombPromotionParamField struct {
	//  交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	//  合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
}

// 组合优惠比例
type CThostFtdcCombPromotionParamField struct {
	//  交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	//  合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
	//  投机套保标志
	CombHedgeFlag types.TThostFtdcCombHedgeFlagType
	//  期权组合保证金比例
	Xparameter types.TThostFtdcDiscountRatioType
}

// 国密用户登录请求
type CThostFtdcReqUserLoginSMField struct {
	//  交易日
	TradingDay types.TThostFtdcDateType
	//  经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	//  用户代码
	UserID types.TThostFtdcUserIDType
	//  密码
	Password types.TThostFtdcPasswordType
	//  用户端产品信息
	UserProductInfo types.TThostFtdcProductInfoType
	//  接口端产品信息
	InterfaceProductInfo types.TThostFtdcProductInfoType
	//  协议信息
	ProtocolInfo types.TThostFtdcProtocolInfoType
	//  Mac地址
	MacAddress types.TThostFtdcMacAddressType
	//  动态密码
	OneTimePassword types.TThostFtdcPasswordType
	//  保留的无效字段
	Reserve1 types.TThostFtdcOldIPAddressType
	//  登录备注
	LoginRemark types.TThostFtdcLoginRemarkType
	//  终端IP端口
	ClientIPPort types.TThostFtdcIPPortType
	//  终端IP地址
	ClientIPAddress types.TThostFtdcIPAddressType
	//  经纪公司名称
	BrokerName types.TThostFtdcBrokerNameType
	//  认证码
	AuthCode types.TThostFtdcAuthCodeType
	//  App代码
	AppID types.TThostFtdcAppIDType
	//  PIN码
	PIN types.TThostFtdcPasswordType
}

// 投资者风险结算持仓查询
type CThostFtdcQryRiskSettleInvstPositionField struct {
	//  经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	//  投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	//  合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
}

// 风险结算产品查询
type CThostFtdcQryRiskSettleProductStatusField struct {
	//  产品代码
	ProductID types.TThostFtdcInstrumentIDType
}

// 投资者风险结算持仓
type CThostFtdcRiskSettleInvstPositionField struct {
	//  合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
	//  经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	//  投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	//  持仓多空方向
	PosiDirection types.TThostFtdcPosiDirectionType
	//  投机套保标志
	HedgeFlag types.TThostFtdcHedgeFlagType
	//  持仓日期
	PositionDate types.TThostFtdcPositionDateType
	//  上日持仓
	YdPosition types.TThostFtdcVolumeType
	//  今日持仓
	Position types.TThostFtdcVolumeType
	//  多头冻结
	LongFrozen types.TThostFtdcVolumeType
	//  空头冻结
	ShortFrozen types.TThostFtdcVolumeType
	//  开仓冻结金额
	LongFrozenAmount types.TThostFtdcMoneyType
	//  开仓冻结金额
	ShortFrozenAmount types.TThostFtdcMoneyType
	//  开仓量
	OpenVolume types.TThostFtdcVolumeType
	//  平仓量
	CloseVolume types.TThostFtdcVolumeType
	//  开仓金额
	OpenAmount types.TThostFtdcMoneyType
	//  平仓金额
	CloseAmount types.TThostFtdcMoneyType
	//  持仓成本
	PositionCost types.TThostFtdcMoneyType
	//  上次占用的保证金
	PreMargin types.TThostFtdcMoneyType
	//  占用的保证金
	UseMargin types.TThostFtdcMoneyType
	//  冻结的保证金
	FrozenMargin types.TThostFtdcMoneyType
	//  冻结的资金
	FrozenCash types.TThostFtdcMoneyType
	//  冻结的手续费
	FrozenCommission types.TThostFtdcMoneyType
	//  资金差额
	CashIn types.TThostFtdcMoneyType
	//  手续费
	Commission types.TThostFtdcMoneyType
	//  平仓盈亏
	CloseProfit types.TThostFtdcMoneyType
	//  持仓盈亏
	PositionProfit types.TThostFtdcMoneyType
	//  上次结算价
	PreSettlementPrice types.TThostFtdcPriceType
	//  本次结算价
	SettlementPrice types.TThostFtdcPriceType
	//  交易日
	TradingDay types.TThostFtdcDateType
	//  结算编号
	SettlementID types.TThostFtdcSettlementIDType
	//  开仓成本
	OpenCost types.TThostFtdcMoneyType
	//  交易所保证金
	ExchangeMargin types.TThostFtdcMoneyType
	//  组合成交形成的持仓
	CombPosition types.TThostFtdcVolumeType
	//  组合多头冻结
	CombLongFrozen types.TThostFtdcVolumeType
	//  组合空头冻结
	CombShortFrozen types.TThostFtdcVolumeType
	//  逐日盯市平仓盈亏
	CloseProfitByDate types.TThostFtdcMoneyType
	//  逐笔对冲平仓盈亏
	CloseProfitByTrade types.TThostFtdcMoneyType
	//  今日持仓
	TodayPosition types.TThostFtdcVolumeType
	//  保证金率
	MarginRateByMoney types.TThostFtdcRatioType
	//  保证金率(按手数)
	MarginRateByVolume types.TThostFtdcRatioType
	//  执行冻结
	StrikeFrozen types.TThostFtdcVolumeType
	//  执行冻结金额
	StrikeFrozenAmount types.TThostFtdcMoneyType
	//  放弃执行冻结
	AbandonFrozen types.TThostFtdcVolumeType
	//  交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	//  执行冻结的昨仓
	YdStrikeFrozen types.TThostFtdcVolumeType
	//  投资单元代码
	InvestUnitID types.TThostFtdcInvestUnitIDType
	//  持仓成本差值
	PositionCostOffset types.TThostFtdcMoneyType
	//  tas持仓手数
	TasPosition types.TThostFtdcVolumeType
	//  tas持仓成本
	TasPositionCost types.TThostFtdcMoneyType
}

// 风险品种
type CThostFtdcRiskSettleProductStatusField struct {
	//  交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	//  产品编号
	ProductID types.TThostFtdcInstrumentIDType
	//  产品结算状态
	ProductStatus types.TThostFtdcProductStatusType
}

// 风险结算追平信息
type CThostFtdcSyncDeltaInfoField struct {
	//  追平序号
	SyncDeltaSequenceNo types.TThostFtdcSequenceNoType
	//  追平状态
	SyncDeltaStatus types.TThostFtdcSyncDeltaStatusType
	//  追平描述
	SyncDescription types.TThostFtdcSyncDescriptionType
	//  是否只有资金追平
	IsOnlyTrdDelta types.TThostFtdcBoolType
}

// 风险结算追平产品信息
type CThostFtdcSyncDeltaProductStatusField struct {
	//  追平序号
	SyncDeltaSequenceNo types.TThostFtdcSequenceNoType
	//  交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	//  产品代码
	ProductID types.TThostFtdcInstrumentIDType
	//  是否允许交易
	ProductStatus types.TThostFtdcProductStatusType
}

// 风险结算追平持仓明细
type CThostFtdcSyncDeltaInvstPosDtlField struct {
	//  合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
	//  经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	//  投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	//  投机套保标志
	HedgeFlag types.TThostFtdcHedgeFlagType
	//  买卖
	Direction types.TThostFtdcDirectionType
	//  开仓日期
	OpenDate types.TThostFtdcDateType
	//  成交编号
	TradeID types.TThostFtdcTradeIDType
	//  数量
	Volume types.TThostFtdcVolumeType
	//  开仓价
	OpenPrice types.TThostFtdcPriceType
	//  交易日
	TradingDay types.TThostFtdcDateType
	//  结算编号
	SettlementID types.TThostFtdcSettlementIDType
	//  成交类型
	TradeType types.TThostFtdcTradeTypeType
	//  组合合约代码
	CombInstrumentID types.TThostFtdcInstrumentIDType
	//  交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	//  逐日盯市平仓盈亏
	CloseProfitByDate types.TThostFtdcMoneyType
	//  逐笔对冲平仓盈亏
	CloseProfitByTrade types.TThostFtdcMoneyType
	//  逐日盯市持仓盈亏
	PositionProfitByDate types.TThostFtdcMoneyType
	//  逐笔对冲持仓盈亏
	PositionProfitByTrade types.TThostFtdcMoneyType
	//  投资者保证金
	Margin types.TThostFtdcMoneyType
	//  交易所保证金
	ExchMargin types.TThostFtdcMoneyType
	//  保证金率
	MarginRateByMoney types.TThostFtdcRatioType
	//  保证金率(按手数)
	MarginRateByVolume types.TThostFtdcRatioType
	//  昨结算价
	LastSettlementPrice types.TThostFtdcPriceType
	//  结算价
	SettlementPrice types.TThostFtdcPriceType
	//  平仓量
	CloseVolume types.TThostFtdcVolumeType
	//  平仓金额
	CloseAmount types.TThostFtdcMoneyType
	//  先开先平剩余数量
	TimeFirstVolume types.TThostFtdcVolumeType
	//  特殊持仓标志
	SpecPosiType types.TThostFtdcSpecPosiTypeType
	//  操作标志
	ActionDirection types.TThostFtdcActionDirectionType
	//  追平序号
	SyncDeltaSequenceNo types.TThostFtdcSequenceNoType
}

// 风险结算追平组合持仓明细
type CThostFtdcSyncDeltaInvstPosCombDtlField struct {
	//  交易日
	TradingDay types.TThostFtdcDateType
	//  开仓日期
	OpenDate types.TThostFtdcDateType
	//  交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	//  结算编号
	SettlementID types.TThostFtdcSettlementIDType
	//  经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	//  投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	//  组合编号
	ComTradeID types.TThostFtdcTradeIDType
	//  撮合编号
	TradeID types.TThostFtdcTradeIDType
	//  合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
	//  投机套保标志
	HedgeFlag types.TThostFtdcHedgeFlagType
	//  买卖
	Direction types.TThostFtdcDirectionType
	//  持仓量
	TotalAmt types.TThostFtdcVolumeType
	//  投资者保证金
	Margin types.TThostFtdcMoneyType
	//  交易所保证金
	ExchMargin types.TThostFtdcMoneyType
	//  保证金率
	MarginRateByMoney types.TThostFtdcRatioType
	//  保证金率(按手数)
	MarginRateByVolume types.TThostFtdcRatioType
	//  单腿编号
	LegID types.TThostFtdcLegIDType
	//  单腿乘数
	LegMultiple types.TThostFtdcLegMultipleType
	//  成交组号
	TradeGroupID types.TThostFtdcTradeGroupIDType
	//  操作标志
	ActionDirection types.TThostFtdcActionDirectionType
	//  追平序号
	SyncDeltaSequenceNo types.TThostFtdcSequenceNoType
}

// 风险结算追平资金
type CThostFtdcSyncDeltaTradingAccountField struct {
	//  经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	//  投资者帐号
	AccountID types.TThostFtdcAccountIDType
	//  上次质押金额
	PreMortgage types.TThostFtdcMoneyType
	//  上次信用额度
	PreCredit types.TThostFtdcMoneyType
	//  上次存款额
	PreDeposit types.TThostFtdcMoneyType
	//  上次结算准备金
	PreBalance types.TThostFtdcMoneyType
	//  上次占用的保证金
	PreMargin types.TThostFtdcMoneyType
	//  利息基数
	InterestBase types.TThostFtdcMoneyType
	//  利息收入
	Interest types.TThostFtdcMoneyType
	//  入金金额
	Deposit types.TThostFtdcMoneyType
	//  出金金额
	Withdraw types.TThostFtdcMoneyType
	//  冻结的保证金
	FrozenMargin types.TThostFtdcMoneyType
	//  冻结的资金
	FrozenCash types.TThostFtdcMoneyType
	//  冻结的手续费
	FrozenCommission types.TThostFtdcMoneyType
	//  当前保证金总额
	CurrMargin types.TThostFtdcMoneyType
	//  资金差额
	CashIn types.TThostFtdcMoneyType
	//  手续费
	Commission types.TThostFtdcMoneyType
	//  平仓盈亏
	CloseProfit types.TThostFtdcMoneyType
	//  持仓盈亏
	PositionProfit types.TThostFtdcMoneyType
	//  期货结算准备金
	Balance types.TThostFtdcMoneyType
	//  可用资金
	Available types.TThostFtdcMoneyType
	//  可取资金
	WithdrawQuota types.TThostFtdcMoneyType
	//  基本准备金
	Reserve types.TThostFtdcMoneyType
	//  交易日
	TradingDay types.TThostFtdcDateType
	//  结算编号
	SettlementID types.TThostFtdcSettlementIDType
	//  信用额度
	Credit types.TThostFtdcMoneyType
	//  质押金额
	Mortgage types.TThostFtdcMoneyType
	//  交易所保证金
	ExchangeMargin types.TThostFtdcMoneyType
	//  投资者交割保证金
	DeliveryMargin types.TThostFtdcMoneyType
	//  交易所交割保证金
	ExchangeDeliveryMargin types.TThostFtdcMoneyType
	//  保底期货结算准备金
	ReserveBalance types.TThostFtdcMoneyType
	//  币种代码
	CurrencyID types.TThostFtdcCurrencyIDType
	//  上次货币质入金额
	PreFundMortgageIn types.TThostFtdcMoneyType
	//  上次货币质出金额
	PreFundMortgageOut types.TThostFtdcMoneyType
	//  货币质入金额
	FundMortgageIn types.TThostFtdcMoneyType
	//  货币质出金额
	FundMortgageOut types.TThostFtdcMoneyType
	//  货币质押余额
	FundMortgageAvailable types.TThostFtdcMoneyType
	//  可质押货币金额
	MortgageableFund types.TThostFtdcMoneyType
	//  特殊产品占用保证金
	SpecProductMargin types.TThostFtdcMoneyType
	//  特殊产品冻结保证金
	SpecProductFrozenMargin types.TThostFtdcMoneyType
	//  特殊产品手续费
	SpecProductCommission types.TThostFtdcMoneyType
	//  特殊产品冻结手续费
	SpecProductFrozenCommission types.TThostFtdcMoneyType
	//  特殊产品持仓盈亏
	SpecProductPositionProfit types.TThostFtdcMoneyType
	//  特殊产品平仓盈亏
	SpecProductCloseProfit types.TThostFtdcMoneyType
	//  根据持仓盈亏算法计算的特殊产品持仓盈亏
	SpecProductPositionProfitByAlg types.TThostFtdcMoneyType
	//  特殊产品交易所保证金
	SpecProductExchangeMargin types.TThostFtdcMoneyType
	//  延时换汇冻结金额
	FrozenSwap types.TThostFtdcMoneyType
	//  剩余换汇额度
	RemainSwap types.TThostFtdcMoneyType
	//  期权市值
	OptionValue types.TThostFtdcMoneyType
	//  追平序号
	SyncDeltaSequenceNo types.TThostFtdcSequenceNoType
}

// 投资者风险结算总保证金
type CThostFtdcSyncDeltaInitInvstMarginField struct {
	//  经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	//  投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	//  追平前总风险保证金
	LastRiskTotalInvstMargin types.TThostFtdcMoneyType
	//  追平前交易所总风险保证金
	LastRiskTotalExchMargin types.TThostFtdcMoneyType
	//  本次追平品种总保证金
	ThisSyncInvstMargin types.TThostFtdcMoneyType
	//  本次追平品种交易所总保证金
	ThisSyncExchMargin types.TThostFtdcMoneyType
	//  本次未追平品种总保证金
	RemainRiskInvstMargin types.TThostFtdcMoneyType
	//  本次未追平品种交易所总保证金
	RemainRiskExchMargin types.TThostFtdcMoneyType
	//  追平前总特殊产品风险保证金
	LastRiskSpecTotalInvstMargin types.TThostFtdcMoneyType
	//  追平前总特殊产品交易所风险保证金
	LastRiskSpecTotalExchMargin types.TThostFtdcMoneyType
	//  本次追平品种特殊产品总保证金
	ThisSyncSpecInvstMargin types.TThostFtdcMoneyType
	//  本次追平品种特殊产品交易所总保证金
	ThisSyncSpecExchMargin types.TThostFtdcMoneyType
	//  本次未追平品种特殊产品总保证金
	RemainRiskSpecInvstMargin types.TThostFtdcMoneyType
	//  本次未追平品种特殊产品交易所总保证金
	RemainRiskSpecExchMargin types.TThostFtdcMoneyType
	//  追平序号
	SyncDeltaSequenceNo types.TThostFtdcSequenceNoType
}

// 风险结算追平组合优先级
type CThostFtdcSyncDeltaDceCombInstrumentField struct {
	//  合约代码
	CombInstrumentID types.TThostFtdcInstrumentIDType
	//  交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	//  合约在交易所的代码
	ExchangeInstID types.TThostFtdcExchangeInstIDType
	//  成交组号
	TradeGroupID types.TThostFtdcTradeGroupIDType
	//  投机套保标志
	CombHedgeFlag types.TThostFtdcHedgeFlagType
	//  组合类型
	CombinationType types.TThostFtdcDceCombinationTypeType
	//  买卖
	Direction types.TThostFtdcDirectionType
	//  产品代码
	ProductID types.TThostFtdcInstrumentIDType
	//  期权组合保证金比例
	Xparameter types.TThostFtdcDiscountRatioType
	//  操作标志
	ActionDirection types.TThostFtdcActionDirectionType
	//  追平序号
	SyncDeltaSequenceNo types.TThostFtdcSequenceNoType
}

// 风险结算追平投资者期货保证金率
type CThostFtdcSyncDeltaInvstMarginRateField struct {
	//  合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
	//  投资者范围
	InvestorRange types.TThostFtdcInvestorRangeType
	//  经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	//  投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	//  投机套保标志
	HedgeFlag types.TThostFtdcHedgeFlagType
	//  多头保证金率
	LongMarginRatioByMoney types.TThostFtdcRatioType
	//  多头保证金费
	LongMarginRatioByVolume types.TThostFtdcMoneyType
	//  空头保证金率
	ShortMarginRatioByMoney types.TThostFtdcRatioType
	//  空头保证金费
	ShortMarginRatioByVolume types.TThostFtdcMoneyType
	//  是否相对交易所收取
	IsRelative types.TThostFtdcBoolType
	//  操作标志
	ActionDirection types.TThostFtdcActionDirectionType
	//  追平序号
	SyncDeltaSequenceNo types.TThostFtdcSequenceNoType
}

// 风险结算追平交易所期货保证金率
type CThostFtdcSyncDeltaExchMarginRateField struct {
	//  经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	//  合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
	//  投机套保标志
	HedgeFlag types.TThostFtdcHedgeFlagType
	//  多头保证金率
	LongMarginRatioByMoney types.TThostFtdcRatioType
	//  多头保证金费
	LongMarginRatioByVolume types.TThostFtdcMoneyType
	//  空头保证金率
	ShortMarginRatioByMoney types.TThostFtdcRatioType
	//  空头保证金费
	ShortMarginRatioByVolume types.TThostFtdcMoneyType
	//  操作标志
	ActionDirection types.TThostFtdcActionDirectionType
	//  追平序号
	SyncDeltaSequenceNo types.TThostFtdcSequenceNoType
}

// 风险结算追平中金现货期权交易所保证金率
type CThostFtdcSyncDeltaOptExchMarginField struct {
	//  经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	//  合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
	//  投机空头保证金调整系数
	SShortMarginRatioByMoney types.TThostFtdcRatioType
	//  投机空头保证金调整系数
	SShortMarginRatioByVolume types.TThostFtdcMoneyType
	//  保值空头保证金调整系数
	HShortMarginRatioByMoney types.TThostFtdcRatioType
	//  保值空头保证金调整系数
	HShortMarginRatioByVolume types.TThostFtdcMoneyType
	//  套利空头保证金调整系数
	AShortMarginRatioByMoney types.TThostFtdcRatioType
	//  套利空头保证金调整系数
	AShortMarginRatioByVolume types.TThostFtdcMoneyType
	//  做市商空头保证金调整系数
	MShortMarginRatioByMoney types.TThostFtdcRatioType
	//  做市商空头保证金调整系数
	MShortMarginRatioByVolume types.TThostFtdcMoneyType
	//  操作标志
	ActionDirection types.TThostFtdcActionDirectionType
	//  追平序号
	SyncDeltaSequenceNo types.TThostFtdcSequenceNoType
}

// 风险结算追平中金现货期权投资者保证金率
type CThostFtdcSyncDeltaOptInvstMarginField struct {
	//  合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
	//  投资者范围
	InvestorRange types.TThostFtdcInvestorRangeType
	//  经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	//  投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	//  投机空头保证金调整系数
	SShortMarginRatioByMoney types.TThostFtdcRatioType
	//  投机空头保证金调整系数
	SShortMarginRatioByVolume types.TThostFtdcMoneyType
	//  保值空头保证金调整系数
	HShortMarginRatioByMoney types.TThostFtdcRatioType
	//  保值空头保证金调整系数
	HShortMarginRatioByVolume types.TThostFtdcMoneyType
	//  套利空头保证金调整系数
	AShortMarginRatioByMoney types.TThostFtdcRatioType
	//  套利空头保证金调整系数
	AShortMarginRatioByVolume types.TThostFtdcMoneyType
	//  是否跟随交易所收取
	IsRelative types.TThostFtdcBoolType
	//  做市商空头保证金调整系数
	MShortMarginRatioByMoney types.TThostFtdcRatioType
	//  做市商空头保证金调整系数
	MShortMarginRatioByVolume types.TThostFtdcMoneyType
	//  操作标志
	ActionDirection types.TThostFtdcActionDirectionType
	//  追平序号
	SyncDeltaSequenceNo types.TThostFtdcSequenceNoType
}

// 风险结算追平期权标的调整保证金率
type CThostFtdcSyncDeltaInvstMarginRateULField struct {
	//  合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
	//  投资者范围
	InvestorRange types.TThostFtdcInvestorRangeType
	//  经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	//  投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	//  投机套保标志
	HedgeFlag types.TThostFtdcHedgeFlagType
	//  多头保证金率
	LongMarginRatioByMoney types.TThostFtdcRatioType
	//  多头保证金费
	LongMarginRatioByVolume types.TThostFtdcMoneyType
	//  空头保证金率
	ShortMarginRatioByMoney types.TThostFtdcRatioType
	//  空头保证金费
	ShortMarginRatioByVolume types.TThostFtdcMoneyType
	//  操作标志
	ActionDirection types.TThostFtdcActionDirectionType
	//  追平序号
	SyncDeltaSequenceNo types.TThostFtdcSequenceNoType
}

// 风险结算追平期权手续费率
type CThostFtdcSyncDeltaOptInvstCommRateField struct {
	//  合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
	//  投资者范围
	InvestorRange types.TThostFtdcInvestorRangeType
	//  经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	//  投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	//  开仓手续费率
	OpenRatioByMoney types.TThostFtdcRatioType
	//  开仓手续费
	OpenRatioByVolume types.TThostFtdcRatioType
	//  平仓手续费率
	CloseRatioByMoney types.TThostFtdcRatioType
	//  平仓手续费
	CloseRatioByVolume types.TThostFtdcRatioType
	//  平今手续费率
	CloseTodayRatioByMoney types.TThostFtdcRatioType
	//  平今手续费
	CloseTodayRatioByVolume types.TThostFtdcRatioType
	//  执行手续费率
	StrikeRatioByMoney types.TThostFtdcRatioType
	//  执行手续费
	StrikeRatioByVolume types.TThostFtdcRatioType
	//  操作标志
	ActionDirection types.TThostFtdcActionDirectionType
	//  追平序号
	SyncDeltaSequenceNo types.TThostFtdcSequenceNoType
}

// 风险结算追平期货手续费率
type CThostFtdcSyncDeltaInvstCommRateField struct {
	//  合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
	//  投资者范围
	InvestorRange types.TThostFtdcInvestorRangeType
	//  经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	//  投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	//  开仓手续费率
	OpenRatioByMoney types.TThostFtdcRatioType
	//  开仓手续费
	OpenRatioByVolume types.TThostFtdcRatioType
	//  平仓手续费率
	CloseRatioByMoney types.TThostFtdcRatioType
	//  平仓手续费
	CloseRatioByVolume types.TThostFtdcRatioType
	//  平今手续费率
	CloseTodayRatioByMoney types.TThostFtdcRatioType
	//  平今手续费
	CloseTodayRatioByVolume types.TThostFtdcRatioType
	//  操作标志
	ActionDirection types.TThostFtdcActionDirectionType
	//  追平序号
	SyncDeltaSequenceNo types.TThostFtdcSequenceNoType
}

// 风险结算追平交叉汇率
type CThostFtdcSyncDeltaProductExchRateField struct {
	//  产品代码
	ProductID types.TThostFtdcInstrumentIDType
	//  报价币种类型
	QuoteCurrencyID types.TThostFtdcCurrencyIDType
	//  汇率
	ExchangeRate types.TThostFtdcExchangeRateType
	//  操作标志
	ActionDirection types.TThostFtdcActionDirectionType
	//  追平序号
	SyncDeltaSequenceNo types.TThostFtdcSequenceNoType
}

// 风险结算追平行情
type CThostFtdcSyncDeltaDepthMarketDataField struct {
	//  交易日
	TradingDay types.TThostFtdcDateType
	//  合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
	//  交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	//  合约在交易所的代码
	ExchangeInstID types.TThostFtdcExchangeInstIDType
	//  最新价
	LastPrice types.TThostFtdcPriceType
	//  上次结算价
	PreSettlementPrice types.TThostFtdcPriceType
	//  昨收盘
	PreClosePrice types.TThostFtdcPriceType
	//  昨持仓量
	PreOpenInterest types.TThostFtdcLargeVolumeType
	//  今开盘
	OpenPrice types.TThostFtdcPriceType
	//  最高价
	HighestPrice types.TThostFtdcPriceType
	//  最低价
	LowestPrice types.TThostFtdcPriceType
	//  数量
	Volume types.TThostFtdcVolumeType
	//  成交金额
	Turnover types.TThostFtdcMoneyType
	//  持仓量
	OpenInterest types.TThostFtdcLargeVolumeType
	//  今收盘
	ClosePrice types.TThostFtdcPriceType
	//  本次结算价
	SettlementPrice types.TThostFtdcPriceType
	//  涨停板价
	UpperLimitPrice types.TThostFtdcPriceType
	//  跌停板价
	LowerLimitPrice types.TThostFtdcPriceType
	//  昨虚实度
	PreDelta types.TThostFtdcRatioType
	//  今虚实度
	CurrDelta types.TThostFtdcRatioType
	//  最后修改时间
	UpdateTime types.TThostFtdcTimeType
	//  最后修改毫秒
	UpdateMillisec types.TThostFtdcMillisecType
	//  申买价一
	BidPrice1 types.TThostFtdcPriceType
	//  申买量一
	BidVolume1 types.TThostFtdcVolumeType
	//  申卖价一
	AskPrice1 types.TThostFtdcPriceType
	//  申卖量一
	AskVolume1 types.TThostFtdcVolumeType
	//  申买价二
	BidPrice2 types.TThostFtdcPriceType
	//  申买量二
	BidVolume2 types.TThostFtdcVolumeType
	//  申卖价二
	AskPrice2 types.TThostFtdcPriceType
	//  申卖量二
	AskVolume2 types.TThostFtdcVolumeType
	//  申买价三
	BidPrice3 types.TThostFtdcPriceType
	//  申买量三
	BidVolume3 types.TThostFtdcVolumeType
	//  申卖价三
	AskPrice3 types.TThostFtdcPriceType
	//  申卖量三
	AskVolume3 types.TThostFtdcVolumeType
	//  申买价四
	BidPrice4 types.TThostFtdcPriceType
	//  申买量四
	BidVolume4 types.TThostFtdcVolumeType
	//  申卖价四
	AskPrice4 types.TThostFtdcPriceType
	//  申卖量四
	AskVolume4 types.TThostFtdcVolumeType
	//  申买价五
	BidPrice5 types.TThostFtdcPriceType
	//  申买量五
	BidVolume5 types.TThostFtdcVolumeType
	//  申卖价五
	AskPrice5 types.TThostFtdcPriceType
	//  申卖量五
	AskVolume5 types.TThostFtdcVolumeType
	//  当日均价
	AveragePrice types.TThostFtdcPriceType
	//  业务日期
	ActionDay types.TThostFtdcDateType
	//  上带价
	BandingUpperPrice types.TThostFtdcPriceType
	//  下带价
	BandingLowerPrice types.TThostFtdcPriceType
	//  操作标志
	ActionDirection types.TThostFtdcActionDirectionType
	//  追平序号
	SyncDeltaSequenceNo types.TThostFtdcSequenceNoType
}

// 风险结算追平现货指数
type CThostFtdcSyncDeltaIndexPriceField struct {
	//  经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	//  合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
	//  指数现货收盘价
	ClosePrice types.TThostFtdcPriceType
	//  操作标志
	ActionDirection types.TThostFtdcActionDirectionType
	//  追平序号
	SyncDeltaSequenceNo types.TThostFtdcSequenceNoType
}

// 风险结算追平仓单折抵
type CThostFtdcSyncDeltaEWarrantOffsetField struct {
	//  交易日期
	TradingDay types.TThostFtdcTradeDateType
	//  经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	//  投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	//  交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	//  合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
	//  买卖方向
	Direction types.TThostFtdcDirectionType
	//  投机套保标志
	HedgeFlag types.TThostFtdcHedgeFlagType
	//  数量
	Volume types.TThostFtdcVolumeType
	//  操作标志
	ActionDirection types.TThostFtdcActionDirectionType
	//  追平序号
	SyncDeltaSequenceNo types.TThostFtdcSequenceNoType
}

// SPBM期货合约保证金参数
type CThostFtdcSPBMFutureParameterField struct {
	//  交易日
	TradingDay types.TThostFtdcDateType
	//  交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	//  合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
	//  品种代码
	ProdFamilyCode types.TThostFtdcInstrumentIDType
	//  期货合约因子
	Cvf types.TThostFtdcVolumeMultipleType
	//  阶段标识
	TimeRange types.TThostFtdcTimeRangeType
	//  品种保证金标准
	MarginRate types.TThostFtdcRatioType
	//  期货合约内部对锁仓费率折扣比例
	LockRateX types.TThostFtdcRatioType
	//  提高保证金标准
	AddOnRate types.TThostFtdcRatioType
	//  昨结算价
	PreSettlementPrice types.TThostFtdcPriceType
	//  期货合约内部对锁仓附加费率折扣比例
	AddOnLockRateX2 types.TThostFtdcRatioType
}

// SPBM期权合约保证金参数
type CThostFtdcSPBMOptionParameterField struct {
	//  交易日
	TradingDay types.TThostFtdcDateType
	//  交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	//  合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
	//  品种代码
	ProdFamilyCode types.TThostFtdcInstrumentIDType
	//  期权合约因子
	Cvf types.TThostFtdcVolumeMultipleType
	//  期权冲抵价格
	DownPrice types.TThostFtdcPriceType
	//  Delta值
	Delta types.TThostFtdcDeltaType
	//  卖方期权风险转换最低值
	SlimiDelta types.TThostFtdcDeltaType
	//  昨结算价
	PreSettlementPrice types.TThostFtdcPriceType
}

// SPBM品种内对锁仓折扣参数
type CThostFtdcSPBMIntraParameterField struct {
	//  交易日
	TradingDay types.TThostFtdcDateType
	//  交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	//  品种代码
	ProdFamilyCode types.TThostFtdcInstrumentIDType
	//  品种内合约间对锁仓费率折扣比例
	IntraRateY types.TThostFtdcRatioType
	//  品种内合约间对锁仓附加费率折扣比例
	AddOnIntraRateY2 types.TThostFtdcRatioType
}

// SPBM跨品种抵扣参数
type CThostFtdcSPBMInterParameterField struct {
	//  交易日
	TradingDay types.TThostFtdcDateType
	//  交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	//  优先级
	SpreadId types.TThostFtdcSpreadIdType
	//  品种间对锁仓费率折扣比例
	InterRateZ types.TThostFtdcRatioType
	//  第一腿构成品种
	Leg1ProdFamilyCode types.TThostFtdcInstrumentIDType
	//  第二腿构成品种
	Leg2ProdFamilyCode types.TThostFtdcInstrumentIDType
}

// 同步SPBM参数结束
type CThostFtdcSyncSPBMParameterEndField struct {
	//  交易日
	TradingDay types.TThostFtdcDateType
}

// SPBM期货合约保证金参数查询
type CThostFtdcQrySPBMFutureParameterField struct {
	//  交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	//  合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
	//  品种代码
	ProdFamilyCode types.TThostFtdcInstrumentIDType
}

// SPBM期权合约保证金参数查询
type CThostFtdcQrySPBMOptionParameterField struct {
	//  交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	//  合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
	//  品种代码
	ProdFamilyCode types.TThostFtdcInstrumentIDType
}

// SPBM品种内对锁仓折扣参数查询
type CThostFtdcQrySPBMIntraParameterField struct {
	//  交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	//  品种代码
	ProdFamilyCode types.TThostFtdcInstrumentIDType
}

// SPBM跨品种抵扣参数查询
type CThostFtdcQrySPBMInterParameterField struct {
	//  交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	//  第一腿构成品种
	Leg1ProdFamilyCode types.TThostFtdcInstrumentIDType
	//  第二腿构成品种
	Leg2ProdFamilyCode types.TThostFtdcInstrumentIDType
}

// 组合保证金套餐
type CThostFtdcSPBMPortfDefinitionField struct {
	//  交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	//  组合保证金套餐代码
	PortfolioDefID types.TThostFtdcPortfolioDefIDType
	//  品种代码
	ProdFamilyCode types.TThostFtdcInstrumentIDType
	//  是否启用SPBM
	IsSPBM types.TThostFtdcBoolType
}

// 投资者套餐选择
type CThostFtdcSPBMInvestorPortfDefField struct {
	//  交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	//  经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	//  投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	//  组合保证金套餐代码
	PortfolioDefID types.TThostFtdcPortfolioDefIDType
}

// 投资者新型组合保证金系数
type CThostFtdcInvestorPortfMarginRatioField struct {
	//  投资者范围
	InvestorRange types.TThostFtdcInvestorRangeType
	//  经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	//  投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	//  交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	//  会员对投资者收取的保证金和交易所对投资者收取的保证金的比例
	MarginRatio types.TThostFtdcRatioType
	//  产品群代码
	ProductGroupID types.TThostFtdcProductIDType
}

// 组合保证金套餐查询
type CThostFtdcQrySPBMPortfDefinitionField struct {
	//  交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	//  组合保证金套餐代码
	PortfolioDefID types.TThostFtdcPortfolioDefIDType
	//  品种代码
	ProdFamilyCode types.TThostFtdcInstrumentIDType
}

// 投资者套餐选择查询
type CThostFtdcQrySPBMInvestorPortfDefField struct {
	//  交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	//  经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	//  投资者代码
	InvestorID types.TThostFtdcInvestorIDType
}

// 投资者新型组合保证金系数查询
type CThostFtdcQryInvestorPortfMarginRatioField struct {
	//  经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	//  投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	//  交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	//  产品群代码
	ProductGroupID types.TThostFtdcProductIDType
}

// 投资者产品SPBM明细
type CThostFtdcInvestorProdSPBMDetailField struct {
	//  交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	//  经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	//  投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	//  品种代码
	ProdFamilyCode types.TThostFtdcInstrumentIDType
	//  合约内对锁保证金
	IntraInstrMargin types.TThostFtdcMoneyType
	//  买归集保证金
	BCollectingMargin types.TThostFtdcMoneyType
	//  卖归集保证金
	SCollectingMargin types.TThostFtdcMoneyType
	//  品种内合约间对锁保证金
	IntraProdMargin types.TThostFtdcMoneyType
	//  净保证金
	NetMargin types.TThostFtdcMoneyType
	//  产品间对锁保证金
	InterProdMargin types.TThostFtdcMoneyType
	//  裸保证金
	SingleMargin types.TThostFtdcMoneyType
	//  附加保证金
	AddOnMargin types.TThostFtdcMoneyType
	//  交割月保证金
	DeliveryMargin types.TThostFtdcMoneyType
	//  看涨期权最低风险
	CallOptionMinRisk types.TThostFtdcMoneyType
	//  看跌期权最低风险
	PutOptionMinRisk types.TThostFtdcMoneyType
	//  卖方期权最低风险
	OptionMinRisk types.TThostFtdcMoneyType
	//  买方期权冲抵价值
	OptionValueOffset types.TThostFtdcMoneyType
	//  卖方期权权利金
	OptionRoyalty types.TThostFtdcMoneyType
	//  价值冲抵
	RealOptionValueOffset types.TThostFtdcMoneyType
	//  保证金
	Margin types.TThostFtdcMoneyType
	//  交易所保证金
	ExchMargin types.TThostFtdcMoneyType
}

// 投资者产品SPBM明细查询
type CThostFtdcQryInvestorProdSPBMDetailField struct {
	//  交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	//  经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	//  投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	//  品种代码
	ProdFamilyCode types.TThostFtdcInstrumentIDType
}

// 组保交易参数设置
type CThostFtdcPortfTradeParamSettingField struct {
	//  交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	//  经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	//  投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	//  新型组保算法
	Portfolio types.TThostFtdcPortfolioType
	//  撤单是否验资
	IsActionVerify types.TThostFtdcBoolType
	//  平仓是否验资
	IsCloseVerify types.TThostFtdcBoolType
}

// 投资者交易权限设置
type CThostFtdcInvestorTradingRightField struct {
	//  经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	//  投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	//  交易权限
	InvstTradingRight types.TThostFtdcInvstTradingRightType
}

// 质押配比参数
type CThostFtdcMortgageParamField struct {
	//  经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	//  投资者帐号
	AccountID types.TThostFtdcAccountIDType
	//  质押配比系数
	MortgageBalance types.TThostFtdcRatioType
	//  开仓是否验证质押配比
	CheckMortgageRatio types.TThostFtdcBoolType
}

// 可提控制参数
type CThostFtdcWithDrawParamField struct {
	//  经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	//  投资者帐号
	AccountID types.TThostFtdcAccountIDType
	//  参数代码
	WithDrawParamID types.TThostFtdcWithDrawParamIDType
	//  参数代码值
	WithDrawParamValue types.TThostFtdcWithDrawParamValueType
}

// Thost终端用户功能权限
type CThostFtdcThostUserFunctionField struct {
	//  经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	//  用户代码
	UserID types.TThostFtdcUserIDType
	//  Thost终端功能代码
	ThostFunctionCode types.TThostFtdcThostFunctionCodeType
}

// Thost终端用户功能权限查询
type CThostFtdcQryThostUserFunctionField struct {
	//  经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	//  用户代码
	UserID types.TThostFtdcUserIDType
}

// SPBM附加跨品种抵扣参数
type CThostFtdcSPBMAddOnInterParameterField struct {
	//  交易日
	TradingDay types.TThostFtdcDateType
	//  交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	//  优先级
	SpreadId types.TThostFtdcSpreadIdType
	//  品种间对锁仓附加费率折扣比例
	AddOnInterRateZ2 types.TThostFtdcRatioType
	//  第一腿构成品种
	Leg1ProdFamilyCode types.TThostFtdcInstrumentIDType
	//  第二腿构成品种
	Leg2ProdFamilyCode types.TThostFtdcInstrumentIDType
}

// SPBM附加跨品种抵扣参数查询
type CThostFtdcQrySPBMAddOnInterParameterField struct {
	//  交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	//  第一腿构成品种
	Leg1ProdFamilyCode types.TThostFtdcInstrumentIDType
	//  第二腿构成品种
	Leg2ProdFamilyCode types.TThostFtdcInstrumentIDType
}

// 投资者商品组SPMM记录查询
type CThostFtdcQryInvestorCommoditySPMMMarginField struct {
	//  经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	//  投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	//  商品组代码
	CommodityID types.TThostFtdcSPMMProductIDType
}

// 投资者商品群SPMM记录查询
type CThostFtdcQryInvestorCommodityGroupSPMMMarginField struct {
	//  经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	//  投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	//  商品群代码
	CommodityGroupID types.TThostFtdcSPMMProductIDType
}

// SPMM合约参数查询
type CThostFtdcQrySPMMInstParamField struct {
	//  合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
}

// SPMM产品参数查询
type CThostFtdcQrySPMMProductParamField struct {
	//  产品代码
	ProductID types.TThostFtdcSPMMProductIDType
}

// 投资者商品组SPMM记录
type CThostFtdcInvestorCommoditySPMMMarginField struct {
	//  交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	//  经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	//  投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	//  商品组代码
	CommodityID types.TThostFtdcSPMMProductIDType
	//  优惠仓位应收保证金
	MarginBeforeDiscount types.TThostFtdcMoneyType
	//  不优惠仓位应收保证金
	MarginNoDiscount types.TThostFtdcMoneyType
	//  多头实仓风险
	LongPosRisk types.TThostFtdcMoneyType
	//  多头开仓冻结风险
	LongOpenFrozenRisk types.TThostFtdcMoneyType
	//  多头被平冻结风险
	LongCloseFrozenRisk types.TThostFtdcMoneyType
	//  空头实仓风险
	ShortPosRisk types.TThostFtdcMoneyType
	//  空头开仓冻结风险
	ShortOpenFrozenRisk types.TThostFtdcMoneyType
	//  空头被平冻结风险
	ShortCloseFrozenRisk types.TThostFtdcMoneyType
	//  SPMM品种内跨期优惠系数
	IntraCommodityRate types.TThostFtdcSPMMDiscountRatioType
	//  SPMM期权优惠系数
	OptionDiscountRate types.TThostFtdcSPMMDiscountRatioType
	//  实仓对冲优惠金额
	PosDiscount types.TThostFtdcMoneyType
	//  开仓报单对冲优惠金额
	OpenFrozenDiscount types.TThostFtdcMoneyType
	//  品种风险净头
	NetRisk types.TThostFtdcMoneyType
	//  平仓冻结保证金
	CloseFrozenMargin types.TThostFtdcMoneyType
	//  冻结的手续费
	FrozenCommission types.TThostFtdcMoneyType
	//  手续费
	Commission types.TThostFtdcMoneyType
	//  冻结的资金
	FrozenCash types.TThostFtdcMoneyType
	//  资金差额
	CashIn types.TThostFtdcMoneyType
	//  行权冻结资金
	StrikeFrozenMargin types.TThostFtdcMoneyType
}

// 投资者商品群SPMM记录
type CThostFtdcInvestorCommodityGroupSPMMMarginField struct {
	//  交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	//  经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	//  投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	//  商品群代码
	CommodityGroupID types.TThostFtdcSPMMProductIDType
	//  优惠仓位应收保证金
	MarginBeforeDiscount types.TThostFtdcMoneyType
	//  不优惠仓位应收保证金
	MarginNoDiscount types.TThostFtdcMoneyType
	//  多头风险
	LongRisk types.TThostFtdcMoneyType
	//  空头风险
	ShortRisk types.TThostFtdcMoneyType
	//  商品群平仓冻结保证金
	CloseFrozenMargin types.TThostFtdcMoneyType
	//  SPMM跨品种优惠系数
	InterCommodityRate types.TThostFtdcSPMMDiscountRatioType
	//  商品群最小保证金比例
	MiniMarginRatio types.TThostFtdcSPMMDiscountRatioType
	//  投资者保证金和交易所保证金的比例
	AdjustRatio types.TThostFtdcRatioType
	//  SPMM品种内优惠汇总
	IntraCommodityDiscount types.TThostFtdcMoneyType
	//  SPMM跨品种优惠
	InterCommodityDiscount types.TThostFtdcMoneyType
	//  交易所保证金
	ExchMargin types.TThostFtdcMoneyType
	//  投资者保证金
	InvestorMargin types.TThostFtdcMoneyType
	//  冻结的手续费
	FrozenCommission types.TThostFtdcMoneyType
	//  手续费
	Commission types.TThostFtdcMoneyType
	//  冻结的资金
	FrozenCash types.TThostFtdcMoneyType
	//  资金差额
	CashIn types.TThostFtdcMoneyType
	//  行权冻结资金
	StrikeFrozenMargin types.TThostFtdcMoneyType
}

// SPMM合约参数
type CThostFtdcSPMMInstParamField struct {
	//  交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	//  合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
	//  SPMM合约保证金算法
	InstMarginCalID types.TThostFtdcInstMarginCalIDType
	//  商品组代码
	CommodityID types.TThostFtdcSPMMProductIDType
	//  商品群代码
	CommodityGroupID types.TThostFtdcSPMMProductIDType
}

// SPMM产品参数
type CThostFtdcSPMMProductParamField struct {
	//  交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	//  产品代码
	ProductID types.TThostFtdcSPMMProductIDType
	//  商品组代码
	CommodityID types.TThostFtdcSPMMProductIDType
	//  商品群代码
	CommodityGroupID types.TThostFtdcSPMMProductIDType
}

// 席位与交易中心对应关系维护查询
type CThostFtdcQryTraderAssignField struct {
	//  交易员代码
	TraderID types.TThostFtdcTraderIDType
}

// 席位与交易中心对应关系
type CThostFtdcTraderAssignField struct {
	//  应用单元代码
	BrokerID types.TThostFtdcBrokerIDType
	//  交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	//  交易所交易员代码
	TraderID types.TThostFtdcTraderIDType
	//  会员代码
	ParticipantID types.TThostFtdcParticipantIDType
	//  交易中心代码
	DRIdentityID types.TThostFtdcDRIdentityIDType
}

// 投资者申报费阶梯收取设置
type CThostFtdcInvestorInfoCntSettingField struct {
	//  交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	//  经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	//  投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	//  商品代码
	ProductID types.TThostFtdcProductIDType
	//  是否收取申报费
	IsCalInfoComm types.TThostFtdcBoolType
	//  是否限制信息量
	IsLimitInfoMax types.TThostFtdcBoolType
	//  信息量限制笔数
	InfoMaxLimit types.TThostFtdcVolumeType
}

// RCAMS产品组合信息
type CThostFtdcRCAMSCombProductInfoField struct {
	//  交易日
	TradingDay types.TThostFtdcDateType
	//  交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	//  产品代码
	ProductID types.TThostFtdcProductIDType
	//  商品组代码
	CombProductID types.TThostFtdcProductIDType
	//  商品群代码
	ProductGroupID types.TThostFtdcProductIDType
}

// RCAMS同合约风险对冲参数
type CThostFtdcRCAMSInstrParameterField struct {
	//  交易日
	TradingDay types.TThostFtdcDateType
	//  交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	//  产品代码
	ProductID types.TThostFtdcProductIDType
	//  同合约风险对冲比率
	HedgeRate types.TThostFtdcHedgeRateType
}

// RCAMS品种内风险对冲参数
type CThostFtdcRCAMSIntraParameterField struct {
	//  交易日
	TradingDay types.TThostFtdcDateType
	//  交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	//  产品组合代码
	CombProductID types.TThostFtdcProductIDType
	//  品种内对冲比率
	HedgeRate types.TThostFtdcHedgeRateType
}

// RCAMS跨品种风险折抵参数
type CThostFtdcRCAMSInterParameterField struct {
	//  交易日
	TradingDay types.TThostFtdcDateType
	//  交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	//  商品群代码
	ProductGroupID types.TThostFtdcProductIDType
	//  优先级
	Priority types.TThostFtdcRCAMSPriorityType
	//  折抵率
	CreditRate types.TThostFtdcHedgeRateType
	//  产品组合代码1
	CombProduct1 types.TThostFtdcProductIDType
	//  产品组合代码2
	CombProduct2 types.TThostFtdcProductIDType
}

// RCAMS空头期权风险调整参数
type CThostFtdcRCAMSShortOptAdjustParamField struct {
	//  交易日
	TradingDay types.TThostFtdcDateType
	//  交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	//  产品组合代码
	CombProductID types.TThostFtdcProductIDType
	//  投套标志
	HedgeFlag types.TThostFtdcHedgeFlagType
	//  空头期权风险调整标准
	AdjustValue types.TThostFtdcAdjustValueType
}

// RCAMS策略组合持仓
type CThostFtdcRCAMSInvestorCombPositionField struct {
	//  交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	//  经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	//  投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	//  合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
	//  投套标志
	HedgeFlag types.TThostFtdcHedgeFlagType
	//  持仓多空方向
	PosiDirection types.TThostFtdcPosiDirectionType
	//  组合合约代码
	CombInstrumentID types.TThostFtdcInstrumentIDType
	//  单腿编号
	LegID types.TThostFtdcLegIDType
	//  交易所组合合约代码
	ExchangeInstID types.TThostFtdcExchangeInstIDType
	//  持仓量
	TotalAmt types.TThostFtdcVolumeType
	//  交易所保证金
	ExchMargin types.TThostFtdcMoneyType
	//  投资者保证金
	Margin types.TThostFtdcMoneyType
}

// 投资者品种RCAMS保证金
type CThostFtdcInvestorProdRCAMSMarginField struct {
	//  交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	//  经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	//  投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	//  产品组合代码
	CombProductID types.TThostFtdcProductIDType
	//  投套标志
	HedgeFlag types.TThostFtdcHedgeFlagType
	//  商品群代码
	ProductGroupID types.TThostFtdcProductIDType
	//  品种组合前风险
	RiskBeforeDiscount types.TThostFtdcMoneyType
	//  同合约对冲风险
	IntraInstrRisk types.TThostFtdcMoneyType
	//  品种买持仓风险
	BPosRisk types.TThostFtdcMoneyType
	//  品种卖持仓风险
	SPosRisk types.TThostFtdcMoneyType
	//  品种内对冲风险
	IntraProdRisk types.TThostFtdcMoneyType
	//  品种净持仓风险
	NetRisk types.TThostFtdcMoneyType
	//  品种间对冲风险
	InterProdRisk types.TThostFtdcMoneyType
	//  空头期权风险调整
	ShortOptRiskAdj types.TThostFtdcMoneyType
	//  空头期权权利金
	OptionRoyalty types.TThostFtdcMoneyType
	//  大边组合平仓冻结保证金
	MMSACloseFrozenMargin types.TThostFtdcMoneyType
	//  策略组合平仓/行权冻结保证金
	CloseCombFrozenMargin types.TThostFtdcMoneyType
	//  平仓/行权冻结保证金
	CloseFrozenMargin types.TThostFtdcMoneyType
	//  大边组合开仓冻结保证金
	MMSAOpenFrozenMargin types.TThostFtdcMoneyType
	//  交割月期货开仓冻结保证金
	DeliveryOpenFrozenMargin types.TThostFtdcMoneyType
	//  开仓冻结保证金
	OpenFrozenMargin types.TThostFtdcMoneyType
	//  投资者冻结保证金
	UseFrozenMargin types.TThostFtdcMoneyType
	//  大边组合交易所持仓保证金
	MMSAExchMargin types.TThostFtdcMoneyType
	//  交割月期货交易所持仓保证金
	DeliveryExchMargin types.TThostFtdcMoneyType
	//  策略组合交易所保证金
	CombExchMargin types.TThostFtdcMoneyType
	//  交易所持仓保证金
	ExchMargin types.TThostFtdcMoneyType
	//  投资者持仓保证金
	UseMargin types.TThostFtdcMoneyType
}

// RCAMS产品组合信息查询
type CThostFtdcQryRCAMSCombProductInfoField struct {
	//  产品代码
	ProductID types.TThostFtdcProductIDType
	//  商品组代码
	CombProductID types.TThostFtdcProductIDType
	//  商品群代码
	ProductGroupID types.TThostFtdcProductIDType
}

// RCAMS同合约风险对冲参数查询
type CThostFtdcQryRCAMSInstrParameterField struct {
	//  产品代码
	ProductID types.TThostFtdcProductIDType
}

// RCAMS品种内风险对冲参数查询
type CThostFtdcQryRCAMSIntraParameterField struct {
	//  产品组合代码
	CombProductID types.TThostFtdcProductIDType
}

// RCAMS跨品种风险折抵参数查询
type CThostFtdcQryRCAMSInterParameterField struct {
	//  商品群代码
	ProductGroupID types.TThostFtdcProductIDType
	//  产品组合代码1
	CombProduct1 types.TThostFtdcProductIDType
	//  产品组合代码2
	CombProduct2 types.TThostFtdcProductIDType
}

// RCAMS空头期权风险调整参数查询
type CThostFtdcQryRCAMSShortOptAdjustParamField struct {
	//  产品组合代码
	CombProductID types.TThostFtdcProductIDType
}

// RCAMS策略组合持仓查询
type CThostFtdcQryRCAMSInvestorCombPositionField struct {
	//  经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	//  投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	//  合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
	//  组合合约代码
	CombInstrumentID types.TThostFtdcInstrumentIDType
}

// 投资者品种RCAMS保证金查询
type CThostFtdcQryInvestorProdRCAMSMarginField struct {
	//  经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	//  投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	//  产品组合代码
	CombProductID types.TThostFtdcProductIDType
	//  商品群代码
	ProductGroupID types.TThostFtdcProductIDType
}

// RULE合约保证金参数
type CThostFtdcRULEInstrParameterField struct {
	//  交易日
	TradingDay types.TThostFtdcDateType
	//  交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	//  合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
	//  合约类型
	InstrumentClass types.TThostFtdcInstrumentClassType
	//  标准合约
	StdInstrumentID types.TThostFtdcInstrumentIDType
	//  投机买折算系数
	BSpecRatio types.TThostFtdcRatioType
	//  投机卖折算系数
	SSpecRatio types.TThostFtdcRatioType
	//  套保买折算系数
	BHedgeRatio types.TThostFtdcRatioType
	//  套保卖折算系数
	SHedgeRatio types.TThostFtdcRatioType
	//  买附加风险保证金
	BAddOnMargin types.TThostFtdcMoneyType
	//  卖附加风险保证金
	SAddOnMargin types.TThostFtdcMoneyType
	//  商品群号
	CommodityGroupID types.TThostFtdcCommodityGroupIDType
}

// RULE品种内对锁仓折扣参数
type CThostFtdcRULEIntraParameterField struct {
	//  交易日
	TradingDay types.TThostFtdcDateType
	//  交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	//  品种代码
	ProdFamilyCode types.TThostFtdcInstrumentIDType
	//  标准合约
	StdInstrumentID types.TThostFtdcInstrumentIDType
	//  标准合约保证金
	StdInstrMargin types.TThostFtdcMoneyType
	//  一般月份合约组合保证金系数
	UsualIntraRate types.TThostFtdcRatioType
	//  临近交割合约组合保证金系数
	DeliveryIntraRate types.TThostFtdcRatioType
}

// RULE跨品种抵扣参数
type CThostFtdcRULEInterParameterField struct {
	//  交易日
	TradingDay types.TThostFtdcDateType
	//  交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	//  优先级
	SpreadId types.TThostFtdcSpreadIdType
	//  品种间对锁仓费率折扣比例
	InterRate types.TThostFtdcRatioType
	//  第一腿构成品种
	Leg1ProdFamilyCode types.TThostFtdcInstrumentIDType
	//  第二腿构成品种
	Leg2ProdFamilyCode types.TThostFtdcInstrumentIDType
	//  腿1比例系数
	Leg1PropFactor types.TThostFtdcCommonIntType
	//  腿2比例系数
	Leg2PropFactor types.TThostFtdcCommonIntType
	//  商品群号
	CommodityGroupID types.TThostFtdcCommodityGroupIDType
	//  商品群名称
	CommodityGroupName types.TThostFtdcInstrumentNameType
}

// RULE合约保证金参数查询
type CThostFtdcQryRULEInstrParameterField struct {
	//  交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	//  合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
}

// RULE品种内对锁仓折扣参数查询
type CThostFtdcQryRULEIntraParameterField struct {
	//  交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	//  品种代码
	ProdFamilyCode types.TThostFtdcInstrumentIDType
}

// RULE跨品种抵扣参数查询
type CThostFtdcQryRULEInterParameterField struct {
	//  交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	//  第一腿构成品种
	Leg1ProdFamilyCode types.TThostFtdcInstrumentIDType
	//  第二腿构成品种
	Leg2ProdFamilyCode types.TThostFtdcInstrumentIDType
	//  商品群号
	CommodityGroupID types.TThostFtdcCommodityGroupIDType
}

// 投资者产品RULE保证金
type CThostFtdcInvestorProdRULEMarginField struct {
	//  交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	//  经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	//  投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	//  品种代码
	ProdFamilyCode types.TThostFtdcInstrumentIDType
	//  合约类型
	InstrumentClass types.TThostFtdcInstrumentClassType
	//  商品群号
	CommodityGroupID types.TThostFtdcCommodityGroupIDType
	//  买标准持仓
	BStdPosition types.TThostFtdcStdPositionType
	//  卖标准持仓
	SStdPosition types.TThostFtdcStdPositionType
	//  买标准开仓冻结
	BStdOpenFrozen types.TThostFtdcStdPositionType
	//  卖标准开仓冻结
	SStdOpenFrozen types.TThostFtdcStdPositionType
	//  买标准平仓冻结
	BStdCloseFrozen types.TThostFtdcStdPositionType
	//  卖标准平仓冻结
	SStdCloseFrozen types.TThostFtdcStdPositionType
	//  品种内对冲标准持仓
	IntraProdStdPosition types.TThostFtdcStdPositionType
	//  品种内单腿标准持仓
	NetStdPosition types.TThostFtdcStdPositionType
	//  品种间对冲标准持仓
	InterProdStdPosition types.TThostFtdcStdPositionType
	//  单腿标准持仓
	SingleStdPosition types.TThostFtdcStdPositionType
	//  品种内对锁保证金
	IntraProdMargin types.TThostFtdcMoneyType
	//  品种间对锁保证金
	InterProdMargin types.TThostFtdcMoneyType
	//  跨品种单腿保证金
	SingleMargin types.TThostFtdcMoneyType
	//  非组合合约保证金
	NonCombMargin types.TThostFtdcMoneyType
	//  附加保证金
	AddOnMargin types.TThostFtdcMoneyType
	//  交易所保证金
	ExchMargin types.TThostFtdcMoneyType
	//  附加冻结保证金
	AddOnFrozenMargin types.TThostFtdcMoneyType
	//  开仓冻结保证金
	OpenFrozenMargin types.TThostFtdcMoneyType
	//  平仓冻结保证金
	CloseFrozenMargin types.TThostFtdcMoneyType
	//  品种保证金
	Margin types.TThostFtdcMoneyType
	//  冻结保证金
	FrozenMargin types.TThostFtdcMoneyType
}

// 投资者产品RULE保证金查询
type CThostFtdcQryInvestorProdRULEMarginField struct {
	//  交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	//  经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	//  投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	//  品种代码
	ProdFamilyCode types.TThostFtdcInstrumentIDType
	//  商品群号
	CommodityGroupID types.TThostFtdcCommodityGroupIDType
}

// 风险结算追平SPBM组合保证金套餐
type CThostFtdcSyncDeltaSPBMPortfDefinitionField struct {
	//  交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	//  组合保证金套餐代码
	PortfolioDefID types.TThostFtdcPortfolioDefIDType
	//  品种代码
	ProdFamilyCode types.TThostFtdcInstrumentIDType
	//  是否启用SPBM
	IsSPBM types.TThostFtdcBoolType
	//  操作标志
	ActionDirection types.TThostFtdcActionDirectionType
	//  追平序号
	SyncDeltaSequenceNo types.TThostFtdcSequenceNoType
}

// 风险结算追平投资者SPBM套餐选择
type CThostFtdcSyncDeltaSPBMInvstPortfDefField struct {
	//  交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	//  经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	//  投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	//  组合保证金套餐代码
	PortfolioDefID types.TThostFtdcPortfolioDefIDType
	//  操作标志
	ActionDirection types.TThostFtdcActionDirectionType
	//  追平序号
	SyncDeltaSequenceNo types.TThostFtdcSequenceNoType
}

// 风险结算追平SPBM期货合约保证金参数
type CThostFtdcSyncDeltaSPBMFutureParameterField struct {
	//  交易日
	TradingDay types.TThostFtdcDateType
	//  交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	//  合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
	//  品种代码
	ProdFamilyCode types.TThostFtdcInstrumentIDType
	//  期货合约因子
	Cvf types.TThostFtdcVolumeMultipleType
	//  阶段标识
	TimeRange types.TThostFtdcTimeRangeType
	//  品种保证金标准
	MarginRate types.TThostFtdcRatioType
	//  期货合约内部对锁仓费率折扣比例
	LockRateX types.TThostFtdcRatioType
	//  提高保证金标准
	AddOnRate types.TThostFtdcRatioType
	//  昨结算价
	PreSettlementPrice types.TThostFtdcPriceType
	//  期货合约内部对锁仓附加费率折扣比例
	AddOnLockRateX2 types.TThostFtdcRatioType
	//  操作标志
	ActionDirection types.TThostFtdcActionDirectionType
	//  追平序号
	SyncDeltaSequenceNo types.TThostFtdcSequenceNoType
}

// 风险结算追平SPBM期权合约保证金参数
type CThostFtdcSyncDeltaSPBMOptionParameterField struct {
	//  交易日
	TradingDay types.TThostFtdcDateType
	//  交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	//  合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
	//  品种代码
	ProdFamilyCode types.TThostFtdcInstrumentIDType
	//  期权合约因子
	Cvf types.TThostFtdcVolumeMultipleType
	//  期权冲抵价格
	DownPrice types.TThostFtdcPriceType
	//  Delta值
	Delta types.TThostFtdcDeltaType
	//  卖方期权风险转换最低值
	SlimiDelta types.TThostFtdcDeltaType
	//  昨结算价
	PreSettlementPrice types.TThostFtdcPriceType
	//  操作标志
	ActionDirection types.TThostFtdcActionDirectionType
	//  追平序号
	SyncDeltaSequenceNo types.TThostFtdcSequenceNoType
}

// 风险结算追平SPBM品种内对锁仓折扣参数
type CThostFtdcSyncDeltaSPBMIntraParameterField struct {
	//  交易日
	TradingDay types.TThostFtdcDateType
	//  交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	//  品种代码
	ProdFamilyCode types.TThostFtdcInstrumentIDType
	//  品种内合约间对锁仓费率折扣比例
	IntraRateY types.TThostFtdcRatioType
	//  品种内合约间对锁仓附加费率折扣比例
	AddOnIntraRateY2 types.TThostFtdcRatioType
	//  操作标志
	ActionDirection types.TThostFtdcActionDirectionType
	//  追平序号
	SyncDeltaSequenceNo types.TThostFtdcSequenceNoType
}

// 风险结算追平SPBM跨品种抵扣参数
type CThostFtdcSyncDeltaSPBMInterParameterField struct {
	//  交易日
	TradingDay types.TThostFtdcDateType
	//  交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	//  优先级
	SpreadId types.TThostFtdcSpreadIdType
	//  品种间对锁仓费率折扣比例
	InterRateZ types.TThostFtdcRatioType
	//  第一腿构成品种
	Leg1ProdFamilyCode types.TThostFtdcInstrumentIDType
	//  第二腿构成品种
	Leg2ProdFamilyCode types.TThostFtdcInstrumentIDType
	//  操作标志
	ActionDirection types.TThostFtdcActionDirectionType
	//  追平序号
	SyncDeltaSequenceNo types.TThostFtdcSequenceNoType
}

// 风险结算追平SPBM附加跨品种抵扣参数
type CThostFtdcSyncDeltaSPBMAddOnInterParamField struct {
	//  交易日
	TradingDay types.TThostFtdcDateType
	//  交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	//  优先级
	SpreadId types.TThostFtdcSpreadIdType
	//  品种间对锁仓附加费率折扣比例
	AddOnInterRateZ2 types.TThostFtdcRatioType
	//  第一腿构成品种
	Leg1ProdFamilyCode types.TThostFtdcInstrumentIDType
	//  第二腿构成品种
	Leg2ProdFamilyCode types.TThostFtdcInstrumentIDType
	//  操作标志
	ActionDirection types.TThostFtdcActionDirectionType
	//  追平序号
	SyncDeltaSequenceNo types.TThostFtdcSequenceNoType
}

// 风险结算追平SPMM合约参数
type CThostFtdcSyncDeltaSPMMInstParamField struct {
	//  交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	//  合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
	//  SPMM合约保证金算法
	InstMarginCalID types.TThostFtdcInstMarginCalIDType
	//  商品组代码
	CommodityID types.TThostFtdcSPMMProductIDType
	//  商品群代码
	CommodityGroupID types.TThostFtdcSPMMProductIDType
	//  操作标志
	ActionDirection types.TThostFtdcActionDirectionType
	//  追平序号
	SyncDeltaSequenceNo types.TThostFtdcSequenceNoType
}

// 风险结算追平SPMM产品相关参数
type CThostFtdcSyncDeltaSPMMProductParamField struct {
	//  交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	//  产品代码
	ProductID types.TThostFtdcSPMMProductIDType
	//  商品组代码
	CommodityID types.TThostFtdcSPMMProductIDType
	//  商品群代码
	CommodityGroupID types.TThostFtdcSPMMProductIDType
	//  操作标志
	ActionDirection types.TThostFtdcActionDirectionType
	//  追平序号
	SyncDeltaSequenceNo types.TThostFtdcSequenceNoType
}

// 风险结算追平投资者SPMM模板选择
type CThostFtdcSyncDeltaInvestorSPMMModelField struct {
	//  交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	//  经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	//  投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	//  SPMM模板ID
	SPMMModelID types.TThostFtdcSPMMModelIDType
	//  操作标志
	ActionDirection types.TThostFtdcActionDirectionType
	//  追平序号
	SyncDeltaSequenceNo types.TThostFtdcSequenceNoType
}

// 风险结算追平SPMM模板参数设置
type CThostFtdcSyncDeltaSPMMModelParamField struct {
	//  交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	//  SPMM模板ID
	SPMMModelID types.TThostFtdcSPMMModelIDType
	//  商品群代码
	CommodityGroupID types.TThostFtdcSPMMProductIDType
	//  SPMM品种内跨期优惠系数
	IntraCommodityRate types.TThostFtdcSPMMDiscountRatioType
	//  SPMM品种间优惠系数
	InterCommodityRate types.TThostFtdcSPMMDiscountRatioType
	//  SPMM期权优惠系数
	OptionDiscountRate types.TThostFtdcSPMMDiscountRatioType
	//  商品群最小保证金比例
	MiniMarginRatio types.TThostFtdcSPMMDiscountRatioType
	//  操作标志
	ActionDirection types.TThostFtdcActionDirectionType
	//  追平序号
	SyncDeltaSequenceNo types.TThostFtdcSequenceNoType
}

// 风险结算追平RCAMS产品组合信息
type CThostFtdcSyncDeltaRCAMSCombProdInfoField struct {
	//  交易日
	TradingDay types.TThostFtdcDateType
	//  交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	//  产品代码
	ProductID types.TThostFtdcProductIDType
	//  商品组代码
	CombProductID types.TThostFtdcProductIDType
	//  商品群代码
	ProductGroupID types.TThostFtdcProductIDType
	//  操作标志
	ActionDirection types.TThostFtdcActionDirectionType
	//  追平序号
	SyncDeltaSequenceNo types.TThostFtdcSequenceNoType
}

// 风险结算追平RCAMS同合约风险对冲参数
type CThostFtdcSyncDeltaRCAMSInstrParameterField struct {
	//  交易日
	TradingDay types.TThostFtdcDateType
	//  交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	//  产品代码
	ProductID types.TThostFtdcProductIDType
	//  同合约风险对冲比率
	HedgeRate types.TThostFtdcHedgeRateType
	//  操作标志
	ActionDirection types.TThostFtdcActionDirectionType
	//  追平序号
	SyncDeltaSequenceNo types.TThostFtdcSequenceNoType
}

// 风险结算追平RCAMS品种内风险对冲参数
type CThostFtdcSyncDeltaRCAMSIntraParameterField struct {
	//  交易日
	TradingDay types.TThostFtdcDateType
	//  交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	//  产品组合代码
	CombProductID types.TThostFtdcProductIDType
	//  品种内对冲比率
	HedgeRate types.TThostFtdcHedgeRateType
	//  操作标志
	ActionDirection types.TThostFtdcActionDirectionType
	//  追平序号
	SyncDeltaSequenceNo types.TThostFtdcSequenceNoType
}

// 风险结算追平RCAMS跨品种风险折抵参数
type CThostFtdcSyncDeltaRCAMSInterParameterField struct {
	//  交易日
	TradingDay types.TThostFtdcDateType
	//  交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	//  商品群代码
	ProductGroupID types.TThostFtdcProductIDType
	//  优先级
	Priority types.TThostFtdcRCAMSPriorityType
	//  折抵率
	CreditRate types.TThostFtdcHedgeRateType
	//  产品组合代码1
	CombProduct1 types.TThostFtdcProductIDType
	//  产品组合代码2
	CombProduct2 types.TThostFtdcProductIDType
	//  操作标志
	ActionDirection types.TThostFtdcActionDirectionType
	//  追平序号
	SyncDeltaSequenceNo types.TThostFtdcSequenceNoType
}

// 风险结算追平RCAMS空头期权风险调整参数
type CThostFtdcSyncDeltaRCAMSSOptAdjParamField struct {
	//  交易日
	TradingDay types.TThostFtdcDateType
	//  交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	//  产品组合代码
	CombProductID types.TThostFtdcProductIDType
	//  投套标志
	HedgeFlag types.TThostFtdcHedgeFlagType
	//  空头期权风险调整标准
	AdjustValue types.TThostFtdcAdjustValueType
	//  操作标志
	ActionDirection types.TThostFtdcActionDirectionType
	//  追平序号
	SyncDeltaSequenceNo types.TThostFtdcSequenceNoType
}

// 风险结算追平RCAMS策略组合规则明细
type CThostFtdcSyncDeltaRCAMSCombRuleDtlField struct {
	//  交易日
	TradingDay types.TThostFtdcDateType
	//  交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	//  策略产品
	ProdGroup types.TThostFtdcProductIDType
	//  策略id
	RuleId types.TThostFtdcRuleIdType
	//  优先级
	Priority types.TThostFtdcRCAMSPriorityType
	//  投套标志
	HedgeFlag types.TThostFtdcHedgeFlagType
	//  组合保证金标准
	CombMargin types.TThostFtdcMoneyType
	//  交易所组合合约代码
	ExchangeInstID types.TThostFtdcExchangeInstIDType
	//  单腿编号
	LegID types.TThostFtdcLegIDType
	//  单腿合约代码
	LegInstrumentID types.TThostFtdcInstrumentIDType
	//  买卖方向
	Direction types.TThostFtdcDirectionType
	//  单腿乘数
	LegMultiple types.TThostFtdcLegMultipleType
	//  操作标志
	ActionDirection types.TThostFtdcActionDirectionType
	//  追平序号
	SyncDeltaSequenceNo types.TThostFtdcSequenceNoType
}

// 风险结算追平RCAMS策略组合持仓
type CThostFtdcSyncDeltaRCAMSInvstCombPosField struct {
	//  交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	//  经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	//  投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	//  合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
	//  投套标志
	HedgeFlag types.TThostFtdcHedgeFlagType
	//  持仓多空方向
	PosiDirection types.TThostFtdcPosiDirectionType
	//  组合合约代码
	CombInstrumentID types.TThostFtdcInstrumentIDType
	//  单腿编号
	LegID types.TThostFtdcLegIDType
	//  交易所组合合约代码
	ExchangeInstID types.TThostFtdcExchangeInstIDType
	//  持仓量
	TotalAmt types.TThostFtdcVolumeType
	//  交易所保证金
	ExchMargin types.TThostFtdcMoneyType
	//  投资者保证金
	Margin types.TThostFtdcMoneyType
	//  操作标志
	ActionDirection types.TThostFtdcActionDirectionType
	//  追平序号
	SyncDeltaSequenceNo types.TThostFtdcSequenceNoType
}

// 风险结算追平RULE合约保证金参数
type CThostFtdcSyncDeltaRULEInstrParameterField struct {
	//  交易日
	TradingDay types.TThostFtdcDateType
	//  交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	//  合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
	//  合约类型
	InstrumentClass types.TThostFtdcInstrumentClassType
	//  标准合约
	StdInstrumentID types.TThostFtdcInstrumentIDType
	//  投机买折算系数
	BSpecRatio types.TThostFtdcRatioType
	//  投机卖折算系数
	SSpecRatio types.TThostFtdcRatioType
	//  套保买折算系数
	BHedgeRatio types.TThostFtdcRatioType
	//  套保卖折算系数
	SHedgeRatio types.TThostFtdcRatioType
	//  买附加风险保证金
	BAddOnMargin types.TThostFtdcMoneyType
	//  卖附加风险保证金
	SAddOnMargin types.TThostFtdcMoneyType
	//  商品群号
	CommodityGroupID types.TThostFtdcCommodityGroupIDType
	//  操作标志
	ActionDirection types.TThostFtdcActionDirectionType
	//  追平序号
	SyncDeltaSequenceNo types.TThostFtdcSequenceNoType
}

// 风险结算追平RULE品种内对锁仓折扣参数
type CThostFtdcSyncDeltaRULEIntraParameterField struct {
	//  交易日
	TradingDay types.TThostFtdcDateType
	//  交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	//  品种代码
	ProdFamilyCode types.TThostFtdcInstrumentIDType
	//  标准合约
	StdInstrumentID types.TThostFtdcInstrumentIDType
	//  标准合约保证金
	StdInstrMargin types.TThostFtdcMoneyType
	//  一般月份合约组合保证金系数
	UsualIntraRate types.TThostFtdcRatioType
	//  临近交割合约组合保证金系数
	DeliveryIntraRate types.TThostFtdcRatioType
	//  操作标志
	ActionDirection types.TThostFtdcActionDirectionType
	//  追平序号
	SyncDeltaSequenceNo types.TThostFtdcSequenceNoType
}

// 风险结算追平RULE跨品种抵扣参数
type CThostFtdcSyncDeltaRULEInterParameterField struct {
	//  交易日
	TradingDay types.TThostFtdcDateType
	//  交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	//  优先级
	SpreadId types.TThostFtdcSpreadIdType
	//  品种间对锁仓费率折扣比例
	InterRate types.TThostFtdcRatioType
	//  第一腿构成品种
	Leg1ProdFamilyCode types.TThostFtdcInstrumentIDType
	//  第二腿构成品种
	Leg2ProdFamilyCode types.TThostFtdcInstrumentIDType
	//  腿1比例系数
	Leg1PropFactor types.TThostFtdcCommonIntType
	//  腿2比例系数
	Leg2PropFactor types.TThostFtdcCommonIntType
	//  商品群号
	CommodityGroupID types.TThostFtdcCommodityGroupIDType
	//  商品群名称
	CommodityGroupName types.TThostFtdcInstrumentNameType
	//  操作标志
	ActionDirection types.TThostFtdcActionDirectionType
	//  追平序号
	SyncDeltaSequenceNo types.TThostFtdcSequenceNoType
}

// 服务地址参数
type CThostFtdcIpAddrParamField struct {
	//  经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	//  服务地址
	Address types.TThostFtdcIpAddrType
	//  交易中心代码
	DRIdentityID types.TThostFtdcDRIdentityIDType
	//  交易中心名称
	DRIdentityName types.TThostFtdcDRIdentityNameType
	//  交易地址OR行情地址
	AddrSrvMode types.TThostFtdcAddrSrvModeType
	//  地址版本
	AddrVer types.TThostFtdcAddrVerType
	//  服务地址编号
	AddrNo types.TThostFtdcCommonIntType
	//  服务地址名称
	AddrName types.TThostFtdcAddrNameType
	//  是否是国密地址
	IsSM types.TThostFtdcBoolType
	//  是否是内网地址
	IsLocalAddr types.TThostFtdcBoolType
	//  地址补充信息
	Remark types.TThostFtdcAddrRemarkType
	//  站点
	Site types.TThostFtdcSiteType
	//  网络运营商
	NetOperator types.TThostFtdcNetOperatorType
}

// 服务地址参数查询
type CThostFtdcQryIpAddrParamField struct {
	//  经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
}

// 服务地址参数
type CThostFtdcTGIpAddrParamField struct {
	//  经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	//  用户代码
	UserID types.TThostFtdcUserIDType
	//  服务地址
	Address types.TThostFtdcIpAddrType
	//  交易中心代码
	DRIdentityID types.TThostFtdcDRIdentityIDType
	//  交易中心名称
	DRIdentityName types.TThostFtdcDRIdentityNameType
	//  交易地址OR行情地址
	AddrSrvMode types.TThostFtdcAddrSrvModeType
	//  地址版本
	AddrVer types.TThostFtdcAddrVerType
	//  服务地址编号
	AddrNo types.TThostFtdcCommonIntType
	//  服务地址名称
	AddrName types.TThostFtdcAddrNameType
	//  是否是国密地址
	IsSM types.TThostFtdcBoolType
	//  是否是内网地址
	IsLocalAddr types.TThostFtdcBoolType
	//  地址补充信息
	Remark types.TThostFtdcAddrRemarkType
	//  站点
	Site types.TThostFtdcSiteType
	//  网络运营商
	NetOperator types.TThostFtdcNetOperatorType
}

// 服务地址参数查询
type CThostFtdcQryTGIpAddrParamField struct {
	//  经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	//  用户代码
	UserID types.TThostFtdcUserIDType
	//  App代码
	AppID types.TThostFtdcAppIDType
}

// TGate会话查询状态
type CThostFtdcTGSessionQryStatusField struct {
	//  最近30s的查询频率
	LastQryFreq types.TThostFtdcCommonIntType
	//  查询状态
	QryStatus types.TThostFtdcTGSessionQryStatusType
}

// 内网地址配置
type CThostFtdcLocalAddrConfigField struct {
	//  经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	//  对端地址
	PeerAddr types.TThostFtdcIpAddrType
	//  子网掩码
	NetMask types.TThostFtdcIpAddrType
	//  交易中心代码
	DRIdentityID types.TThostFtdcDRIdentityIDType
	//  内网服务地址
	LocalAddress types.TThostFtdcIpAddrType
}

// 内网地址配置查询
type CThostFtdcQryLocalAddrConfigField struct {
	//  经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
}

// 次席查询银行资金帐户信息请求
type CThostFtdcReqQueryBankAccountBySecField struct {
	//  业务功能码
	TradeCode types.TThostFtdcTradeCodeType
	//  银行代码
	BankID types.TThostFtdcBankIDType
	//  银行分支机构代码
	BankBranchID types.TThostFtdcBankBrchIDType
	//  期商代码
	BrokerID types.TThostFtdcBrokerIDType
	//  期商分支机构代码
	BrokerBranchID types.TThostFtdcFutureBranchIDType
	//  交易日期
	TradeDate types.TThostFtdcTradeDateType
	//  交易时间
	TradeTime types.TThostFtdcTradeTimeType
	//  银行流水号
	BankSerial types.TThostFtdcBankSerialType
	//  交易系统日期
	TradingDay types.TThostFtdcTradeDateType
	//  银期平台消息流水号
	PlateSerial types.TThostFtdcSerialType
	//  最后分片标志
	LastFragment types.TThostFtdcLastFragmentType
	//  会话号
	SessionID types.TThostFtdcSessionIDType
	//  客户姓名
	CustomerName types.TThostFtdcIndividualNameType
	//  证件类型
	IdCardType types.TThostFtdcIdCardTypeType
	//  证件号码
	IdentifiedCardNo types.TThostFtdcIdentifiedCardNoType
	//  客户类型
	CustType types.TThostFtdcCustTypeType
	//  银行帐号
	BankAccount types.TThostFtdcBankAccountType
	//  银行密码
	BankPassWord types.TThostFtdcPasswordType
	//  投资者帐号
	AccountID types.TThostFtdcAccountIDType
	//  期货密码
	Password types.TThostFtdcPasswordType
	//  期货公司流水号
	FutureSerial types.TThostFtdcFutureSerialType
	//  安装编号
	InstallID types.TThostFtdcInstallIDType
	//  用户标识
	UserID types.TThostFtdcUserIDType
	//  验证客户证件号码标志
	VerifyCertNoFlag types.TThostFtdcYesNoIndicatorType
	//  币种代码
	CurrencyID types.TThostFtdcCurrencyIDType
	//  摘要
	Digest types.TThostFtdcDigestType
	//  银行帐号类型
	BankAccType types.TThostFtdcBankAccTypeType
	//  渠道标志
	DeviceID types.TThostFtdcDeviceIDType
	//  期货单位帐号类型
	BankSecuAccType types.TThostFtdcBankAccTypeType
	//  期货公司银行编码
	BrokerIDByBank types.TThostFtdcBankCodingForFutureType
	//  期货单位帐号
	BankSecuAcc types.TThostFtdcBankAccountType
	//  银行密码标志
	BankPwdFlag types.TThostFtdcPwdFlagType
	//  期货资金密码核对标志
	SecuPwdFlag types.TThostFtdcPwdFlagType
	//  交易柜员
	OperNo types.TThostFtdcOperNoType
	//  请求编号
	RequestID types.TThostFtdcRequestIDType
	//  交易ID
	TID types.TThostFtdcTIDType
	//  长客户姓名
	LongCustomerName types.TThostFtdcLongIndividualNameType
	//  交易中心代码
	DRIdentityID types.TThostFtdcDRIdentityIDType
	//  次中心发起转账期货公司流水号
	SecFutureSerial types.TThostFtdcFutureSerialType
}

// 次席查询银行资金帐户信息回报
type CThostFtdcRspQueryBankAccountBySecField struct {
	//  业务功能码
	TradeCode types.TThostFtdcTradeCodeType
	//  银行代码
	BankID types.TThostFtdcBankIDType
	//  银行分支机构代码
	BankBranchID types.TThostFtdcBankBrchIDType
	//  期商代码
	BrokerID types.TThostFtdcBrokerIDType
	//  期商分支机构代码
	BrokerBranchID types.TThostFtdcFutureBranchIDType
	//  交易日期
	TradeDate types.TThostFtdcTradeDateType
	//  交易时间
	TradeTime types.TThostFtdcTradeTimeType
	//  银行流水号
	BankSerial types.TThostFtdcBankSerialType
	//  交易系统日期
	TradingDay types.TThostFtdcTradeDateType
	//  银期平台消息流水号
	PlateSerial types.TThostFtdcSerialType
	//  最后分片标志
	LastFragment types.TThostFtdcLastFragmentType
	//  会话号
	SessionID types.TThostFtdcSessionIDType
	//  客户姓名
	CustomerName types.TThostFtdcIndividualNameType
	//  证件类型
	IdCardType types.TThostFtdcIdCardTypeType
	//  证件号码
	IdentifiedCardNo types.TThostFtdcIdentifiedCardNoType
	//  客户类型
	CustType types.TThostFtdcCustTypeType
	//  银行帐号
	BankAccount types.TThostFtdcBankAccountType
	//  银行密码
	BankPassWord types.TThostFtdcPasswordType
	//  投资者帐号
	AccountID types.TThostFtdcAccountIDType
	//  期货密码
	Password types.TThostFtdcPasswordType
	//  期货公司流水号
	FutureSerial types.TThostFtdcFutureSerialType
	//  安装编号
	InstallID types.TThostFtdcInstallIDType
	//  用户标识
	UserID types.TThostFtdcUserIDType
	//  验证客户证件号码标志
	VerifyCertNoFlag types.TThostFtdcYesNoIndicatorType
	//  币种代码
	CurrencyID types.TThostFtdcCurrencyIDType
	//  摘要
	Digest types.TThostFtdcDigestType
	//  银行帐号类型
	BankAccType types.TThostFtdcBankAccTypeType
	//  渠道标志
	DeviceID types.TThostFtdcDeviceIDType
	//  期货单位帐号类型
	BankSecuAccType types.TThostFtdcBankAccTypeType
	//  期货公司银行编码
	BrokerIDByBank types.TThostFtdcBankCodingForFutureType
	//  期货单位帐号
	BankSecuAcc types.TThostFtdcBankAccountType
	//  银行密码标志
	BankPwdFlag types.TThostFtdcPwdFlagType
	//  期货资金密码核对标志
	SecuPwdFlag types.TThostFtdcPwdFlagType
	//  交易柜员
	OperNo types.TThostFtdcOperNoType
	//  请求编号
	RequestID types.TThostFtdcRequestIDType
	//  交易ID
	TID types.TThostFtdcTIDType
	//  银行可用金额
	BankUseAmount types.TThostFtdcTradeAmountType
	//  银行可取金额
	BankFetchAmount types.TThostFtdcTradeAmountType
	//  长客户姓名
	LongCustomerName types.TThostFtdcLongIndividualNameType
	//  交易中心代码
	DRIdentityID types.TThostFtdcDRIdentityIDType
	//  次中心发起转账期货公司流水号
	SecFutureSerial types.TThostFtdcFutureSerialType
}

// 次中心发起的转帐交易
type CThostFtdcReqTransferBySecField struct {
	//  业务功能码
	TradeCode types.TThostFtdcTradeCodeType
	//  银行代码
	BankID types.TThostFtdcBankIDType
	//  银行分支机构代码
	BankBranchID types.TThostFtdcBankBrchIDType
	//  期商代码
	BrokerID types.TThostFtdcBrokerIDType
	//  期商分支机构代码
	BrokerBranchID types.TThostFtdcFutureBranchIDType
	//  交易日期
	TradeDate types.TThostFtdcTradeDateType
	//  交易时间
	TradeTime types.TThostFtdcTradeTimeType
	//  银行流水号
	BankSerial types.TThostFtdcBankSerialType
	//  交易系统日期
	TradingDay types.TThostFtdcTradeDateType
	//  银期平台消息流水号
	PlateSerial types.TThostFtdcSerialType
	//  最后分片标志
	LastFragment types.TThostFtdcLastFragmentType
	//  会话号
	SessionID types.TThostFtdcSessionIDType
	//  客户姓名
	CustomerName types.TThostFtdcIndividualNameType
	//  证件类型
	IdCardType types.TThostFtdcIdCardTypeType
	//  证件号码
	IdentifiedCardNo types.TThostFtdcIdentifiedCardNoType
	//  客户类型
	CustType types.TThostFtdcCustTypeType
	//  银行帐号
	BankAccount types.TThostFtdcBankAccountType
	//  银行密码
	BankPassWord types.TThostFtdcPasswordType
	//  投资者帐号
	AccountID types.TThostFtdcAccountIDType
	//  期货密码
	Password types.TThostFtdcPasswordType
	//  安装编号
	InstallID types.TThostFtdcInstallIDType
	//  期货公司流水号
	FutureSerial types.TThostFtdcFutureSerialType
	//  用户标识
	UserID types.TThostFtdcUserIDType
	//  验证客户证件号码标志
	VerifyCertNoFlag types.TThostFtdcYesNoIndicatorType
	//  币种代码
	CurrencyID types.TThostFtdcCurrencyIDType
	//  转帐金额
	TradeAmount types.TThostFtdcTradeAmountType
	//  期货可取金额
	FutureFetchAmount types.TThostFtdcTradeAmountType
	//  费用支付标志
	FeePayFlag types.TThostFtdcFeePayFlagType
	//  应收客户费用
	CustFee types.TThostFtdcCustFeeType
	//  应收期货公司费用
	BrokerFee types.TThostFtdcFutureFeeType
	//  发送方给接收方的消息
	Message types.TThostFtdcAddInfoType
	//  摘要
	Digest types.TThostFtdcDigestType
	//  银行帐号类型
	BankAccType types.TThostFtdcBankAccTypeType
	//  渠道标志
	DeviceID types.TThostFtdcDeviceIDType
	//  期货单位帐号类型
	BankSecuAccType types.TThostFtdcBankAccTypeType
	//  期货公司银行编码
	BrokerIDByBank types.TThostFtdcBankCodingForFutureType
	//  期货单位帐号
	BankSecuAcc types.TThostFtdcBankAccountType
	//  银行密码标志
	BankPwdFlag types.TThostFtdcPwdFlagType
	//  期货资金密码核对标志
	SecuPwdFlag types.TThostFtdcPwdFlagType
	//  交易柜员
	OperNo types.TThostFtdcOperNoType
	//  请求编号
	RequestID types.TThostFtdcRequestIDType
	//  交易ID
	TID types.TThostFtdcTIDType
	//  转账交易状态
	TransferStatus types.TThostFtdcTransferStatusType
	//  长客户姓名
	LongCustomerName types.TThostFtdcLongIndividualNameType
	//  交易中心代码
	DRIdentityID types.TThostFtdcDRIdentityIDType
	//  次中心发起转账期货公司流水号
	SecFutureSerial types.TThostFtdcFutureSerialType
}

// 次中心发起的转帐交易回报
type CThostFtdcRspTransferBySecField struct {
	//  业务功能码
	TradeCode types.TThostFtdcTradeCodeType
	//  银行代码
	BankID types.TThostFtdcBankIDType
	//  银行分支机构代码
	BankBranchID types.TThostFtdcBankBrchIDType
	//  期商代码
	BrokerID types.TThostFtdcBrokerIDType
	//  期商分支机构代码
	BrokerBranchID types.TThostFtdcFutureBranchIDType
	//  交易日期
	TradeDate types.TThostFtdcTradeDateType
	//  交易时间
	TradeTime types.TThostFtdcTradeTimeType
	//  银行流水号
	BankSerial types.TThostFtdcBankSerialType
	//  交易系统日期
	TradingDay types.TThostFtdcTradeDateType
	//  银期平台消息流水号
	PlateSerial types.TThostFtdcSerialType
	//  最后分片标志
	LastFragment types.TThostFtdcLastFragmentType
	//  会话号
	SessionID types.TThostFtdcSessionIDType
	//  客户姓名
	CustomerName types.TThostFtdcIndividualNameType
	//  证件类型
	IdCardType types.TThostFtdcIdCardTypeType
	//  证件号码
	IdentifiedCardNo types.TThostFtdcIdentifiedCardNoType
	//  客户类型
	CustType types.TThostFtdcCustTypeType
	//  银行帐号
	BankAccount types.TThostFtdcBankAccountType
	//  银行密码
	BankPassWord types.TThostFtdcPasswordType
	//  投资者帐号
	AccountID types.TThostFtdcAccountIDType
	//  期货密码
	Password types.TThostFtdcPasswordType
	//  安装编号
	InstallID types.TThostFtdcInstallIDType
	//  期货公司流水号
	FutureSerial types.TThostFtdcFutureSerialType
	//  用户标识
	UserID types.TThostFtdcUserIDType
	//  验证客户证件号码标志
	VerifyCertNoFlag types.TThostFtdcYesNoIndicatorType
	//  币种代码
	CurrencyID types.TThostFtdcCurrencyIDType
	//  转帐金额
	TradeAmount types.TThostFtdcTradeAmountType
	//  期货可取金额
	FutureFetchAmount types.TThostFtdcTradeAmountType
	//  费用支付标志
	FeePayFlag types.TThostFtdcFeePayFlagType
	//  应收客户费用
	CustFee types.TThostFtdcCustFeeType
	//  应收期货公司费用
	BrokerFee types.TThostFtdcFutureFeeType
	//  发送方给接收方的消息
	Message types.TThostFtdcAddInfoType
	//  摘要
	Digest types.TThostFtdcDigestType
	//  银行帐号类型
	BankAccType types.TThostFtdcBankAccTypeType
	//  渠道标志
	DeviceID types.TThostFtdcDeviceIDType
	//  期货单位帐号类型
	BankSecuAccType types.TThostFtdcBankAccTypeType
	//  期货公司银行编码
	BrokerIDByBank types.TThostFtdcBankCodingForFutureType
	//  期货单位帐号
	BankSecuAcc types.TThostFtdcBankAccountType
	//  银行密码标志
	BankPwdFlag types.TThostFtdcPwdFlagType
	//  期货资金密码核对标志
	SecuPwdFlag types.TThostFtdcPwdFlagType
	//  交易柜员
	OperNo types.TThostFtdcOperNoType
	//  请求编号
	RequestID types.TThostFtdcRequestIDType
	//  交易ID
	TID types.TThostFtdcTIDType
	//  转账交易状态
	TransferStatus types.TThostFtdcTransferStatusType
	//  错误代码
	ErrorID types.TThostFtdcErrorIDType
	//  错误信息
	ErrorMsg types.TThostFtdcErrorMsgType
	//  长客户姓名
	LongCustomerName types.TThostFtdcLongIndividualNameType
	//  交易中心代码
	DRIdentityID types.TThostFtdcDRIdentityIDType
	//  次中心发起转账期货公司流水号
	SecFutureSerial types.TThostFtdcFutureSerialType
}

// 退出紧急状态参数
type CThostFtdcExitEmergencyField struct {
	//  经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
}

// 新组保保证金系数投资者模板对应关系
type CThostFtdcInvestorPortfMarginModelField struct {
	//  经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	//  投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	//  保证金系数模板
	MarginModelID types.TThostFtdcInvestorIDType
}

// 投资者新组保设置
type CThostFtdcInvestorPortfSettingField struct {
	//  交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	//  经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	//  投资者编号
	InvestorID types.TThostFtdcInvestorIDType
	//  投机套保标志
	HedgeFlag types.TThostFtdcHedgeFlagType
	//  是否开启新组保
	UsePortf types.TThostFtdcBoolType
}

// 投资者新组保设置查询
type CThostFtdcQryInvestorPortfSettingField struct {
	//  交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	//  经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	//  投资者编号
	InvestorID types.TThostFtdcInvestorIDType
}

// 来自次席的用户口令变更
type CThostFtdcUserPasswordUpdateFromSecField struct {
	//  经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	//  用户代码
	UserID types.TThostFtdcUserIDType
	//  原来的口令
	OldPassword types.TThostFtdcPasswordType
	//  新的口令
	NewPassword types.TThostFtdcPasswordType
	//  次席的交易中心代码
	FromSec types.TThostFtdcDRIdentityIDType
}

// 来自次席的结算结果确认
type CThostFtdcSettlementInfoConfirmFromSecField struct {
	//  经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	//  投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	//  确认日期
	ConfirmDate types.TThostFtdcDateType
	//  确认时间
	ConfirmTime types.TThostFtdcTimeType
	//  次席的交易中心代码
	FromSec types.TThostFtdcDRIdentityIDType
}

// 来自次席的资金账户口令变更
type CThostFtdcTradingAccountPasswordUpdateFromSecField struct {
	//  经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	//  投资者帐号
	AccountID types.TThostFtdcAccountIDType
	//  原来的口令
	OldPassword types.TThostFtdcPasswordType
	//  新的口令
	NewPassword types.TThostFtdcPasswordType
	//  币种代码
	CurrencyID types.TThostFtdcCurrencyIDType
	//  次席的交易中心代码
	FromSec types.TThostFtdcDRIdentityIDType
}

// 风控禁止的合约交易权限
type CThostFtdcRiskForbiddenRightField struct {
	//  经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	//  投资者编号
	InvestorID types.TThostFtdcInvestorIDType
	//  合约/产品代码
	InstrumentID types.TThostFtdcInstrumentIDType
	//  用户代码
	UserID types.TThostFtdcUserIDType
}

// 投资者申报费阶梯收取记录
type CThostFtdcInvestorInfoCommRecField struct {
	//  交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	//  经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	//  投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	//  商品代码
	InstrumentID types.TThostFtdcInstrumentIDType
	//  报单总笔数
	OrderCount types.TThostFtdcVolumeType
	//  撤单总笔数
	OrderActionCount types.TThostFtdcVolumeType
	//  询价总次数
	ForQuoteCnt types.TThostFtdcVolumeType
	//  申报费
	InfoComm types.TThostFtdcMoneyType
	//  是否期权系列
	IsOptSeries types.TThostFtdcBoolType
	//  品种代码
	ProductID types.TThostFtdcProductIDType
	//  信息量总量
	InfoCnt types.TThostFtdcVolumeType
}

// 投资者申报费阶梯收取记录查询
type CThostFtdcQryInvestorInfoCommRecField struct {
	//  投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	//  商品代码
	InstrumentID types.TThostFtdcInstrumentIDType
	//  经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
}

// 组合腿信息
type CThostFtdcCombLegField struct {
	//  组合合约代码
	CombInstrumentID types.TThostFtdcInstrumentIDType
	//  单腿编号
	LegID types.TThostFtdcLegIDType
	//  单腿合约代码
	LegInstrumentID types.TThostFtdcInstrumentIDType
	//  买卖方向
	Direction types.TThostFtdcDirectionType
	//  单腿乘数
	LegMultiple types.TThostFtdcLegMultipleType
	//  派生层数
	ImplyLevel types.TThostFtdcImplyLevelType
}

// 组合腿信息查询
type CThostFtdcQryCombLegField struct {
	//  单腿合约代码
	LegInstrumentID types.TThostFtdcInstrumentIDType
}

// 输入的对冲设置
type CThostFtdcInputOffsetSettingField struct {
	//  经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	//  投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	//  合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
	//  标的期货合约代码
	UnderlyingInstrID types.TThostFtdcInstrumentIDType
	//  产品代码
	ProductID types.TThostFtdcProductIDType
	//  对冲类型
	OffsetType types.TThostFtdcOffsetTypeType
	//  申请对冲的合约数量
	Volume types.TThostFtdcVolumeType
	//  是否对冲
	IsOffset types.TThostFtdcBoolType
	//  请求编号
	RequestID types.TThostFtdcRequestIDType
	//  用户代码
	UserID types.TThostFtdcUserIDType
	//  交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	//  IP地址
	IPAddress types.TThostFtdcIPAddressType
	//  Mac地址
	MacAddress types.TThostFtdcMacAddressType
}

// 对冲设置
type CThostFtdcOffsetSettingField struct {
	//  经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	//  投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	//  合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
	//  标的期货合约代码
	UnderlyingInstrID types.TThostFtdcInstrumentIDType
	//  产品代码
	ProductID types.TThostFtdcProductIDType
	//  对冲类型
	OffsetType types.TThostFtdcOffsetTypeType
	//  申请对冲的合约数量
	Volume types.TThostFtdcVolumeType
	//  是否对冲
	IsOffset types.TThostFtdcBoolType
	//  请求编号
	RequestID types.TThostFtdcRequestIDType
	//  用户代码
	UserID types.TThostFtdcUserIDType
	//  交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	//  IP地址
	IPAddress types.TThostFtdcIPAddressType
	//  Mac地址
	MacAddress types.TThostFtdcMacAddressType
	//  交易所合约代码
	ExchangeInstID types.TThostFtdcExchangeInstIDType
	//  交易所期权系列号
	ExchangeSerialNo types.TThostFtdcExchangeInstIDType
	//  交易所产品代码
	ExchangeProductID types.TThostFtdcProductIDType
	//  会员代码
	ParticipantID types.TThostFtdcParticipantIDType
	//  客户代码
	ClientID types.TThostFtdcClientIDType
	//  交易所交易员代码
	TraderID types.TThostFtdcTraderIDType
	//  安装编号
	InstallID types.TThostFtdcInstallIDType
	//  对冲提交状态
	OrderSubmitStatus types.TThostFtdcOrderSubmitStatusType
	//  交易日
	TradingDay types.TThostFtdcDateType
	//  结算编号
	SettlementID types.TThostFtdcSettlementIDType
	//  报单日期
	InsertDate types.TThostFtdcDateType
	//  插入时间
	InsertTime types.TThostFtdcTimeType
	//  撤销时间
	CancelTime types.TThostFtdcTimeType
	//  对冲设置结果
	ExecResult types.TThostFtdcExecResultType
	//  序号
	SequenceNo types.TThostFtdcSequenceNoType
	//  前置编号
	FrontID types.TThostFtdcFrontIDType
	//  会话编号
	SessionID types.TThostFtdcSessionIDType
	//  状态信息
	StatusMsg types.TThostFtdcErrorMsgType
	//  操作用户代码
	ActiveUserID types.TThostFtdcUserIDType
	//  经纪公司报单编号
	BrokerOffsetSettingSeq types.TThostFtdcSequenceNoType
	//  申请来源
	ApplySrc types.TThostFtdcApplySrcType
}

// 撤销对冲设置
type CThostFtdcCancelOffsetSettingField struct {
	//  经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	//  投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	//  合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
	//  标的期货合约代码
	UnderlyingInstrID types.TThostFtdcInstrumentIDType
	//  产品代码
	ProductID types.TThostFtdcProductIDType
	//  对冲类型
	OffsetType types.TThostFtdcOffsetTypeType
	//  申请对冲的合约数量
	Volume types.TThostFtdcVolumeType
	//  是否对冲
	IsOffset types.TThostFtdcBoolType
	//  请求编号
	RequestID types.TThostFtdcRequestIDType
	//  用户代码
	UserID types.TThostFtdcUserIDType
	//  交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	//  IP地址
	IPAddress types.TThostFtdcIPAddressType
	//  Mac地址
	MacAddress types.TThostFtdcMacAddressType
	//  交易所合约代码
	ExchangeInstID types.TThostFtdcExchangeInstIDType
	//  交易所期权系列号
	ExchangeSerialNo types.TThostFtdcExchangeInstIDType
	//  交易所产品代码
	ExchangeProductID types.TThostFtdcProductIDType
	//  交易所交易员代码
	TraderID types.TThostFtdcTraderIDType
	//  安装编号
	InstallID types.TThostFtdcInstallIDType
	//  会员代码
	ParticipantID types.TThostFtdcParticipantIDType
	//  客户代码
	ClientID types.TThostFtdcClientIDType
	//  报单操作状态
	OrderActionStatus types.TThostFtdcOrderActionStatusType
	//  状态信息
	StatusMsg types.TThostFtdcErrorMsgType
	//  操作本地编号
	ActionLocalID types.TThostFtdcOrderLocalIDType
	//  操作日期
	ActionDate types.TThostFtdcDateType
	//  操作时间
	ActionTime types.TThostFtdcTimeType
}

// 查询对冲设置
type CThostFtdcQryOffsetSettingField struct {
	//  经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	//  投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	//  产品代码
	ProductID types.TThostFtdcProductIDType
	//  对冲类型
	OffsetType types.TThostFtdcOffsetTypeType
}

// 服务地址和AppID的关系
type CThostFtdcAddrAppIDRelationField struct {
	//  经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	//  服务地址
	Address types.TThostFtdcIpAddrType
	//  交易中心代码
	DRIdentityID types.TThostFtdcDRIdentityIDType
	//  App代码
	AppID types.TThostFtdcAppIDType
}

// 服务地址和AppID的关系查询
type CThostFtdcQryAddrAppIDRelationField struct {
	//  经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
}

// 微信小程序等用户系统信息
type CThostFtdcWechatUserSystemInfoField struct {
	//  经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	//  用户代码
	UserID types.TThostFtdcUserIDType
	//  微信小程序等用户端系统内部信息长度
	WechatCltSysInfoLen types.TThostFtdcSystemInfoLenType
	//  微信小程序等用户端系统内部信息
	WechatCltSysInfo types.TThostFtdcClientSystemInfoType
	//  终端IP端口
	ClientIPPort types.TThostFtdcIPPortType
	//  登录成功时间
	ClientLoginTime types.TThostFtdcTimeType
	//  App代码
	ClientAppID types.TThostFtdcAppIDType
	//  用户公网IP
	ClientPublicIP types.TThostFtdcIPAddressType
	//  客户登录备注2
	ClientLoginRemark types.TThostFtdcClientLoginRemarkType
}

// 投资者预留信息
type CThostFtdcInvestorReserveInfoField struct {
	//  经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	//  用户代码
	UserID types.TThostFtdcUserIDType
	//  预留信息
	ReserveInfo types.TThostFtdcReserveInfoType
}

// 查询组织架构投资者对应关系
type CThostFtdcQryInvestorDepartmentFlatField struct {
	//  经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
}

// 组织架构投资者对应关系
type CThostFtdcInvestorDepartmentFlatField struct {
	//  经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	//  投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	//  组织架构代码
	DepartmentID types.TThostFtdcInvestorIDType
}

// 查询操作员组织架构关系
type CThostFtdcQryDepartmentUserField struct {
	//  经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
}

// 前置信息
type CThostFtdcFrontInfoField struct {
	//  前置地址
	FrontAddr types.TThostFtdcAddressType
	//  查询流控
	QryFreq types.TThostFtdcQueryFreqType
	//  FTD流控
	FTDPkgFreq types.TThostFtdcQueryFreqType
}

// 申请短信验证码请求
type CThostFtdcReqGenSMSCodeField struct {
	//  经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	//  用户代码
	UserID types.TThostFtdcUserIDType
	//  手机号
	Mobile types.TThostFtdcSMSPhoneType
}

// 申请短信验证码响应
type CThostFtdcRspGenSMSCodeField struct {
	//  经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	//  用户代码
	UserID types.TThostFtdcUserIDType
	//  生成时间
	GenTime types.TThostFtdcTimeType
}

// 短信验证信息通知
type CThostFtdcSMSVerifyInfoFromSecField struct {
	//  经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	//  经纪公司简称
	BrokerAbbr types.TThostFtdcBrokerAbbrType
	//  用户代码
	UserID types.TThostFtdcUserIDType
	//  手机号
	Mobile types.TThostFtdcSMSPhoneType
	//  短信验证码
	SMSCode types.TThostFtdcSMSCodeType
	//  验证码创建日期
	CreateDate types.TThostFtdcDateType
	//  验证码创建时间
	CreateTime types.TThostFtdcTimeType
	//  验证码是否被使用过
	IsUsed types.TThostFtdcBoolType
	//  次席的交易中心代码
	FromSec types.TThostFtdcDRIdentityIDType
}

// 登录验证设置
type CThostFtdcSMSVerifyConfigField struct {
	//  用户代码
	UserID types.TThostFtdcUserIDType
	//  经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	//  手机号
	Mobile types.TThostFtdcSMSPhoneType
	//  是否启用短信验证
	UseSMSVerify types.TThostFtdcBoolType
}

// 短信验证信息通知
type CThostFtdcSMSVerifyInfoField struct {
	//  验证码创建时间
	CreateTime types.TThostFtdcTimeType
	//  手机号
	Mobile types.TThostFtdcSMSPhoneType
	//  短信验证信息内容
	SMSContent types.TThostFtdcSMSContentType
}

// 套利确认输入基本信息
type CThostFtdcInputSpdApplyField struct {
	//  经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	//  用户代码
	UserID types.TThostFtdcUserIDType
	//  投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	//  交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	//  合约代码
	FirstLegInstrumentID types.TThostFtdcInstrumentIDType
	//  合约代码
	SecondLegInstrumentID types.TThostFtdcInstrumentIDType
	//  数量
	Volume types.TThostFtdcVolumeType
	//  买卖方向
	Direction types.TThostFtdcDirectionType
	//  组合定单类型
	CmbType types.TThostFtdcCmbTypeType
	//  请求编号
	RequestID types.TThostFtdcRequestIDType
	//  报单引用
	OrderRef types.TThostFtdcOrderRefType
	//  IP地址
	IPAddress types.TThostFtdcIPAddressType
	//  Mac地址
	MacAddress types.TThostFtdcMacAddressType
}

// 套保确认输入基本信息
type CThostFtdcInputHedgeCfmField struct {
	//  经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	//  用户代码
	UserID types.TThostFtdcUserIDType
	//  投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	//  交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	//  合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
	//  数量
	Volume types.TThostFtdcVolumeType
	//  买卖方向
	Direction types.TThostFtdcDirectionType
	//  请求编号
	RequestID types.TThostFtdcRequestIDType
	//  报单引用
	OrderRef types.TThostFtdcOrderRefType
	//  IP地址
	IPAddress types.TThostFtdcIPAddressType
	//  Mac地址
	MacAddress types.TThostFtdcMacAddressType
}

// 套利申请回报
type CThostFtdcSpdApplyField struct {
	//  经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	//  投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	//  合约代码
	FirstLegInstrumentID types.TThostFtdcInstrumentIDType
	//  合约代码
	SecondLegInstrumentID types.TThostFtdcInstrumentIDType
	//  用户代码
	UserID types.TThostFtdcUserIDType
	//  数量
	Volume types.TThostFtdcVolumeType
	//  买卖方向
	Direction types.TThostFtdcDirectionType
	//  请求编号
	RequestID types.TThostFtdcRequestIDType
	//  前置编号
	FrontID types.TThostFtdcFrontIDType
	//  会话编号
	SessionID types.TThostFtdcSessionIDType
	//  报单引用
	OrderRef types.TThostFtdcOrderRefType
	//  操作用户代码
	ActiveUserID types.TThostFtdcUserIDType
	//  经纪公司报单编号
	BrokerOrderSeq types.TThostFtdcSequenceNoType
	//  报单编号
	OrderSysID types.TThostFtdcOrderSysIDType
	//  申请状态
	ApplyStatus types.TThostFtdcApplyStatusType
	//  序号
	SequenceNo types.TThostFtdcSequenceNoType
	//  报单日期
	InsertDate types.TThostFtdcDateType
	//  委托时间
	InsertTime types.TThostFtdcTimeType
	//  撤销时间
	CancelTime types.TThostFtdcTimeType
	//  本地报单编号
	OrderLocalID types.TThostFtdcOrderLocalIDType
	//  交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	//  会员代码
	ParticipantID types.TThostFtdcParticipantIDType
	//  客户代码
	ClientID types.TThostFtdcClientIDType
	//  合约在交易所的代码
	ExchangeInstID types.TThostFtdcExchangeInstIDType
	//  交易所交易员代码
	TraderID types.TThostFtdcTraderIDType
	//  安装编号
	InstallID types.TThostFtdcInstallIDType
	//  报单提交状态
	OrderSubmitStatus types.TThostFtdcOrderSubmitStatusType
	//  报单提示序号
	NotifySequence types.TThostFtdcSequenceNoType
	//  交易日
	TradingDay types.TThostFtdcDateType
	//  结算编号
	SettlementID types.TThostFtdcSettlementIDType
	//  IP地址
	IPAddress types.TThostFtdcIPAddressType
	//  Mac地址
	MacAddress types.TThostFtdcMacAddressType
	//  组合定单类型
	CmbType types.TThostFtdcCmbTypeType
	//  状态信息
	StatusMsg types.TThostFtdcErrorMsgType
}

// 套保申请回报
type CThostFtdcHedgeCfmField struct {
	//  经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	//  投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	//  合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
	//  用户代码
	UserID types.TThostFtdcUserIDType
	//  数量
	Volume types.TThostFtdcVolumeType
	//  买卖方向
	Direction types.TThostFtdcDirectionType
	//  请求编号
	RequestID types.TThostFtdcRequestIDType
	//  前置编号
	FrontID types.TThostFtdcFrontIDType
	//  会话编号
	SessionID types.TThostFtdcSessionIDType
	//  报单引用
	OrderRef types.TThostFtdcOrderRefType
	//  操作用户代码
	ActiveUserID types.TThostFtdcUserIDType
	//  经纪公司报单编号
	BrokerOrderSeq types.TThostFtdcSequenceNoType
	//  报单编号
	OrderSysID types.TThostFtdcOrderSysIDType
	//  申请状态
	ApplyStatus types.TThostFtdcApplyStatusType
	//  序号
	SequenceNo types.TThostFtdcSequenceNoType
	//  成功处理数量
	DealVolume types.TThostFtdcVolumeType
	//  报单日期
	InsertDate types.TThostFtdcDateType
	//  委托时间
	InsertTime types.TThostFtdcTimeType
	//  撤销时间
	CancelTime types.TThostFtdcTimeType
	//  日期
	ReqDate types.TThostFtdcDateType
	//  本地报单编号
	OrderLocalID types.TThostFtdcOrderLocalIDType
	//  交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	//  会员代码
	ParticipantID types.TThostFtdcParticipantIDType
	//  客户代码
	ClientID types.TThostFtdcClientIDType
	//  合约在交易所的代码
	ExchangeInstID types.TThostFtdcExchangeInstIDType
	//  交易所交易员代码
	TraderID types.TThostFtdcTraderIDType
	//  安装编号
	InstallID types.TThostFtdcInstallIDType
	//  报单提交状态
	OrderSubmitStatus types.TThostFtdcOrderSubmitStatusType
	//  报单提示序号
	NotifySequence types.TThostFtdcSequenceNoType
	//  交易日
	TradingDay types.TThostFtdcDateType
	//  结算编号
	SettlementID types.TThostFtdcSettlementIDType
	//  状态信息
	StatusMsg types.TThostFtdcErrorMsgType
	//  IP地址
	IPAddress types.TThostFtdcIPAddressType
	//  Mac地址
	MacAddress types.TThostFtdcMacAddressType
}

// 套利套保申请查询
type CThostFtdcQrySpdApplyField struct {
	//  经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	//  投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	//  交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	//  报单编号
	OrderSysID types.TThostFtdcOrderSysIDType
	//  第一腿合约编码
	FirstLegInstrumentID types.TThostFtdcExchangeInstIDType
	//  第二腿合约编码
	SecondLegInstrumentID types.TThostFtdcExchangeInstIDType
}

// 套利套保申请查询
type CThostFtdcQryHedgeCfmField struct {
	//  经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	//  投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	//  交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	//  报单编号
	OrderSysID types.TThostFtdcOrderSysIDType
	//  合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
}

// 套利申请撤销
type CThostFtdcInputSpdApplyActionField struct {
	//  经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	//  用户代码
	UserID types.TThostFtdcUserIDType
	//  投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	//  交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	//  合同编号
	OrderSysID types.TThostFtdcOrderSysIDType
	//  报单引用
	OrderRef types.TThostFtdcOrderRefType
	//  前置编号
	FrontID types.TThostFtdcFrontIDType
	//  会话编号
	SessionID types.TThostFtdcSessionIDType
	//  请求编号
	RequestID types.TThostFtdcRequestIDType
	//  IP地址
	IPAddress types.TThostFtdcIPAddressType
	//  Mac地址
	MacAddress types.TThostFtdcMacAddressType
}

// 套保申请撤销
type CThostFtdcInputHedgeCfmActionField struct {
	//  经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	//  用户代码
	UserID types.TThostFtdcUserIDType
	//  投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	//  交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	//  合同编号
	OrderSysID types.TThostFtdcOrderSysIDType
	//  报单引用
	OrderRef types.TThostFtdcOrderRefType
	//  前置编号
	FrontID types.TThostFtdcFrontIDType
	//  会话编号
	SessionID types.TThostFtdcSessionIDType
	//  请求编号
	RequestID types.TThostFtdcRequestIDType
	//  IP地址
	IPAddress types.TThostFtdcIPAddressType
	//  Mac地址
	MacAddress types.TThostFtdcMacAddressType
}

// 套利申请撤销回报
type CThostFtdcSpdApplyActionField struct {
	//  经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	//  投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	//  操作日期
	ActionDate types.TThostFtdcDateType
	//  操作时间
	ActionTime types.TThostFtdcTimeType
	//  交易所交易员代码
	TraderID types.TThostFtdcTraderIDType
	//  安装编号
	InstallID types.TThostFtdcInstallIDType
	//  本地报单编号
	OrderLocalID types.TThostFtdcOrderLocalIDType
	//  操作本地编号
	ActionLocalID types.TThostFtdcOrderLocalIDType
	//  会员代码
	ParticipantID types.TThostFtdcParticipantIDType
	//  客户代码
	ClientID types.TThostFtdcClientIDType
	//  报单操作状态
	OrderActionStatus types.TThostFtdcOrderActionStatusType
	//  用户代码
	UserID types.TThostFtdcUserIDType
	//  交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	//  合同编号
	OrderSysID types.TThostFtdcOrderSysIDType
	//  请求编号
	RequestID types.TThostFtdcRequestIDType
	//  状态信息
	StatusMsg types.TThostFtdcErrorMsgType
	//  报单引用
	OrderRef types.TThostFtdcOrderRefType
	//  前置编号
	FrontID types.TThostFtdcFrontIDType
	//  会话编号
	SessionID types.TThostFtdcSessionIDType
	//  IP地址
	IPAddress types.TThostFtdcIPAddressType
	//  Mac地址
	MacAddress types.TThostFtdcMacAddressType
}

// 套保申请撤销回报
type CThostFtdcHedgeCfmActionField struct {
	//  经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	//  投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	//  操作日期
	ActionDate types.TThostFtdcDateType
	//  操作时间
	ActionTime types.TThostFtdcTimeType
	//  交易所交易员代码
	TraderID types.TThostFtdcTraderIDType
	//  安装编号
	InstallID types.TThostFtdcInstallIDType
	//  本地报单编号
	OrderLocalID types.TThostFtdcOrderLocalIDType
	//  操作本地编号
	ActionLocalID types.TThostFtdcOrderLocalIDType
	//  会员代码
	ParticipantID types.TThostFtdcParticipantIDType
	//  客户代码
	ClientID types.TThostFtdcClientIDType
	//  报单操作状态
	OrderActionStatus types.TThostFtdcOrderActionStatusType
	//  用户代码
	UserID types.TThostFtdcUserIDType
	//  交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	//  合同编号
	OrderSysID types.TThostFtdcOrderSysIDType
	//  请求编号
	RequestID types.TThostFtdcRequestIDType
	//  状态信息
	StatusMsg types.TThostFtdcErrorMsgType
	//  报单引用
	OrderRef types.TThostFtdcOrderRefType
	//  前置编号
	FrontID types.TThostFtdcFrontIDType
	//  会话编号
	SessionID types.TThostFtdcSessionIDType
	//  IP地址
	IPAddress types.TThostFtdcIPAddressType
	//  Mac地址
	MacAddress types.TThostFtdcMacAddressType
}
