package future

import (
	"fmt"
	"strings"

	"github.com/frozenpine/ctp4go/thost/future/types"
)

// 信息分发
type CThostFtdcDisseminationField struct {
	// 序列系列号
	SequenceSeries types.TThostFtdcSequenceSeriesType
	// 序列号
	SequenceNo types.TThostFtdcSequenceNoType
}

func (d CThostFtdcDisseminationField) Type() string { return "CThostFtdcDisseminationField" }

func (d CThostFtdcDisseminationField) String() string {
	var builder strings.Builder
	builder.Grow(72)

	builder.WriteString("CThostFtdcDisseminationField{")
	fmt.Fprintf(&builder, "SequenceSeries=%+v", d.SequenceSeries)
	fmt.Fprintf(&builder, ", SequenceNo=%+v", d.SequenceNo)

	builder.WriteByte('}')

	return builder.String()
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
	// 保留的无效字段
	reserve1 types.TThostFtdcOldIPAddressType
	// 登录备注
	LoginRemark types.TThostFtdcLoginRemarkType
	// 终端IP端口
	ClientIPPort types.TThostFtdcIPPortType
	// 终端IP地址
	ClientIPAddress types.TThostFtdcIPAddressType
	// 短信验证码
	SMSCode types.TThostFtdcSMSCodeType
}

func (d CThostFtdcReqUserLoginField) Type() string { return "CThostFtdcReqUserLoginField" }

func (d CThostFtdcReqUserLoginField) String() string {
	var builder strings.Builder
	builder.Grow(324)

	builder.WriteString("CThostFtdcReqUserLoginField{")
	fmt.Fprintf(&builder, "TradingDay=%+v", d.TradingDay)
	fmt.Fprintf(&builder, ", BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", UserID=%+v", d.UserID)
	fmt.Fprintf(&builder, ", Password=%+v", d.Password)
	fmt.Fprintf(&builder, ", UserProductInfo=%+v", d.UserProductInfo)
	fmt.Fprintf(&builder, ", InterfaceProductInfo=%+v", d.InterfaceProductInfo)
	fmt.Fprintf(&builder, ", ProtocolInfo=%+v", d.ProtocolInfo)
	fmt.Fprintf(&builder, ", MacAddress=%+v", d.MacAddress)
	fmt.Fprintf(&builder, ", OneTimePassword=%+v", d.OneTimePassword)
	fmt.Fprintf(&builder, ", reserve1=%+v", d.reserve1)
	fmt.Fprintf(&builder, ", LoginRemark=%+v", d.LoginRemark)
	fmt.Fprintf(&builder, ", ClientIPPort=%+v", d.ClientIPPort)
	fmt.Fprintf(&builder, ", ClientIPAddress=%+v", d.ClientIPAddress)
	fmt.Fprintf(&builder, ", SMSCode=%+v", d.SMSCode)

	builder.WriteByte('}')

	return builder.String()
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
	// 后台版本信息
	SysVersion types.TThostFtdcSysVersionType
	// 广期所时间
	GFEXTime types.TThostFtdcTimeType
	// 当前登录中心号
	LoginDRIdentityID types.TThostFtdcDRIdentityIDType
	// 用户所属中心号
	UserDRIdentityID types.TThostFtdcDRIdentityIDType
	// 上次登陆时间
	LastLoginTime types.TThostFtdcDateTimeType
	// 预留信息
	ReserveInfo types.TThostFtdcReserveInfoType
}

func (d CThostFtdcRspUserLoginField) Type() string { return "CThostFtdcRspUserLoginField" }

func (d CThostFtdcRspUserLoginField) String() string {
	var builder strings.Builder
	builder.Grow(400)

	builder.WriteString("CThostFtdcRspUserLoginField{")
	fmt.Fprintf(&builder, "TradingDay=%+v", d.TradingDay)
	fmt.Fprintf(&builder, ", LoginTime=%+v", d.LoginTime)
	fmt.Fprintf(&builder, ", BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", UserID=%+v", d.UserID)
	fmt.Fprintf(&builder, ", SystemName=%+v", d.SystemName)
	fmt.Fprintf(&builder, ", FrontID=%+v", d.FrontID)
	fmt.Fprintf(&builder, ", SessionID=%+v", d.SessionID)
	fmt.Fprintf(&builder, ", MaxOrderRef=%+v", d.MaxOrderRef)
	fmt.Fprintf(&builder, ", SHFETime=%+v", d.SHFETime)
	fmt.Fprintf(&builder, ", DCETime=%+v", d.DCETime)
	fmt.Fprintf(&builder, ", CZCETime=%+v", d.CZCETime)
	fmt.Fprintf(&builder, ", FFEXTime=%+v", d.FFEXTime)
	fmt.Fprintf(&builder, ", INETime=%+v", d.INETime)
	fmt.Fprintf(&builder, ", SysVersion=%+v", d.SysVersion)
	fmt.Fprintf(&builder, ", GFEXTime=%+v", d.GFEXTime)
	fmt.Fprintf(&builder, ", LoginDRIdentityID=%+v", d.LoginDRIdentityID)
	fmt.Fprintf(&builder, ", UserDRIdentityID=%+v", d.UserDRIdentityID)
	fmt.Fprintf(&builder, ", LastLoginTime=%+v", d.LastLoginTime)
	fmt.Fprintf(&builder, ", ReserveInfo=%+v", d.ReserveInfo)

	builder.WriteByte('}')

	return builder.String()
}

// 用户登出请求
type CThostFtdcUserLogoutField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 用户代码
	UserID types.TThostFtdcUserIDType
}

func (d CThostFtdcUserLogoutField) Type() string { return "CThostFtdcUserLogoutField" }

func (d CThostFtdcUserLogoutField) String() string {
	var builder strings.Builder
	builder.Grow(59)

	builder.WriteString("CThostFtdcUserLogoutField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", UserID=%+v", d.UserID)

	builder.WriteByte('}')

	return builder.String()
}

// 强制交易员退出
type CThostFtdcForceUserLogoutField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 用户代码
	UserID types.TThostFtdcUserIDType
}

func (d CThostFtdcForceUserLogoutField) Type() string { return "CThostFtdcForceUserLogoutField" }

func (d CThostFtdcForceUserLogoutField) String() string {
	var builder strings.Builder
	builder.Grow(64)

	builder.WriteString("CThostFtdcForceUserLogoutField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", UserID=%+v", d.UserID)

	builder.WriteByte('}')

	return builder.String()
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
	AppID types.TThostFtdcAppIDType
}

func (d CThostFtdcReqAuthenticateField) Type() string { return "CThostFtdcReqAuthenticateField" }

func (d CThostFtdcReqAuthenticateField) String() string {
	var builder strings.Builder
	builder.Grow(122)

	builder.WriteString("CThostFtdcReqAuthenticateField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", UserID=%+v", d.UserID)
	fmt.Fprintf(&builder, ", UserProductInfo=%+v", d.UserProductInfo)
	fmt.Fprintf(&builder, ", AuthCode=%+v", d.AuthCode)
	fmt.Fprintf(&builder, ", AppID=%+v", d.AppID)

	builder.WriteByte('}')

	return builder.String()
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
	AppID types.TThostFtdcAppIDType
	// App类型
	AppType types.TThostFtdcAppTypeType
}

func (d CThostFtdcRspAuthenticateField) Type() string { return "CThostFtdcRspAuthenticateField" }

func (d CThostFtdcRspAuthenticateField) String() string {
	var builder strings.Builder
	builder.Grow(121)

	builder.WriteString("CThostFtdcRspAuthenticateField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", UserID=%+v", d.UserID)
	fmt.Fprintf(&builder, ", UserProductInfo=%+v", d.UserProductInfo)
	fmt.Fprintf(&builder, ", AppID=%+v", d.AppID)
	fmt.Fprintf(&builder, ", AppType=%+v", d.AppType)

	builder.WriteByte('}')

	return builder.String()
}

// 客户端认证信息
type CThostFtdcAuthenticationInfoField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 用户代码
	UserID types.TThostFtdcUserIDType
	// 用户端产品信息
	UserProductInfo types.TThostFtdcProductInfoType
	// 认证信息
	AuthInfo types.TThostFtdcAuthInfoType
	// 是否为认证结果
	IsResult types.TThostFtdcBoolType
	// App代码
	AppID types.TThostFtdcAppIDType
	// App类型
	AppType types.TThostFtdcAppTypeType
	// 保留的无效字段
	reserve1 types.TThostFtdcOldIPAddressType
	// 终端IP地址
	ClientIPAddress types.TThostFtdcIPAddressType
}

func (d CThostFtdcAuthenticationInfoField) Type() string { return "CThostFtdcAuthenticationInfoField" }

func (d CThostFtdcAuthenticationInfoField) String() string {
	var builder strings.Builder
	builder.Grow(203)

	builder.WriteString("CThostFtdcAuthenticationInfoField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", UserID=%+v", d.UserID)
	fmt.Fprintf(&builder, ", UserProductInfo=%+v", d.UserProductInfo)
	fmt.Fprintf(&builder, ", AuthInfo=%+v", d.AuthInfo)
	fmt.Fprintf(&builder, ", IsResult=%+v", d.IsResult)
	fmt.Fprintf(&builder, ", AppID=%+v", d.AppID)
	fmt.Fprintf(&builder, ", AppType=%+v", d.AppType)
	fmt.Fprintf(&builder, ", reserve1=%+v", d.reserve1)
	fmt.Fprintf(&builder, ", ClientIPAddress=%+v", d.ClientIPAddress)

	builder.WriteByte('}')

	return builder.String()
}

// 用户登录应答2
type CThostFtdcRspUserLogin2Field struct {
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
	// 随机串
	RandomString types.TThostFtdcRandomStringType
}

func (d CThostFtdcRspUserLogin2Field) Type() string { return "CThostFtdcRspUserLogin2Field" }

func (d CThostFtdcRspUserLogin2Field) String() string {
	var builder strings.Builder
	builder.Grow(288)

	builder.WriteString("CThostFtdcRspUserLogin2Field{")
	fmt.Fprintf(&builder, "TradingDay=%+v", d.TradingDay)
	fmt.Fprintf(&builder, ", LoginTime=%+v", d.LoginTime)
	fmt.Fprintf(&builder, ", BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", UserID=%+v", d.UserID)
	fmt.Fprintf(&builder, ", SystemName=%+v", d.SystemName)
	fmt.Fprintf(&builder, ", FrontID=%+v", d.FrontID)
	fmt.Fprintf(&builder, ", SessionID=%+v", d.SessionID)
	fmt.Fprintf(&builder, ", MaxOrderRef=%+v", d.MaxOrderRef)
	fmt.Fprintf(&builder, ", SHFETime=%+v", d.SHFETime)
	fmt.Fprintf(&builder, ", DCETime=%+v", d.DCETime)
	fmt.Fprintf(&builder, ", CZCETime=%+v", d.CZCETime)
	fmt.Fprintf(&builder, ", FFEXTime=%+v", d.FFEXTime)
	fmt.Fprintf(&builder, ", INETime=%+v", d.INETime)
	fmt.Fprintf(&builder, ", RandomString=%+v", d.RandomString)

	builder.WriteByte('}')

	return builder.String()
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

func (d CThostFtdcTransferHeaderField) Type() string { return "CThostFtdcTransferHeaderField" }

func (d CThostFtdcTransferHeaderField) String() string {
	var builder strings.Builder
	builder.Grow(269)

	builder.WriteString("CThostFtdcTransferHeaderField{")
	fmt.Fprintf(&builder, "Version=%+v", d.Version)
	fmt.Fprintf(&builder, ", TradeCode=%+v", d.TradeCode)
	fmt.Fprintf(&builder, ", TradeDate=%+v", d.TradeDate)
	fmt.Fprintf(&builder, ", TradeTime=%+v", d.TradeTime)
	fmt.Fprintf(&builder, ", TradeSerial=%+v", d.TradeSerial)
	fmt.Fprintf(&builder, ", FutureID=%+v", d.FutureID)
	fmt.Fprintf(&builder, ", BankID=%+v", d.BankID)
	fmt.Fprintf(&builder, ", BankBrchID=%+v", d.BankBrchID)
	fmt.Fprintf(&builder, ", OperNo=%+v", d.OperNo)
	fmt.Fprintf(&builder, ", DeviceID=%+v", d.DeviceID)
	fmt.Fprintf(&builder, ", RecordNum=%+v", d.RecordNum)
	fmt.Fprintf(&builder, ", SessionID=%+v", d.SessionID)
	fmt.Fprintf(&builder, ", RequestID=%+v", d.RequestID)

	builder.WriteByte('}')

	return builder.String()
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

func (d CThostFtdcTransferBankToFutureReqField) Type() string {
	return "CThostFtdcTransferBankToFutureReqField"
}

func (d CThostFtdcTransferBankToFutureReqField) String() string {
	var builder strings.Builder
	builder.Grow(163)

	builder.WriteString("CThostFtdcTransferBankToFutureReqField{")
	fmt.Fprintf(&builder, "FutureAccount=%+v", d.FutureAccount)
	fmt.Fprintf(&builder, ", FuturePwdFlag=%+v", d.FuturePwdFlag)
	fmt.Fprintf(&builder, ", FutureAccPwd=%+v", d.FutureAccPwd)
	fmt.Fprintf(&builder, ", TradeAmt=%+v", d.TradeAmt)
	fmt.Fprintf(&builder, ", CustFee=%+v", d.CustFee)
	fmt.Fprintf(&builder, ", CurrencyCode=%+v", d.CurrencyCode)

	builder.WriteByte('}')

	return builder.String()
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

func (d CThostFtdcTransferBankToFutureRspField) Type() string {
	return "CThostFtdcTransferBankToFutureRspField"
}

func (d CThostFtdcTransferBankToFutureRspField) String() string {
	var builder strings.Builder
	builder.Grow(152)

	builder.WriteString("CThostFtdcTransferBankToFutureRspField{")
	fmt.Fprintf(&builder, "RetCode=%+v", d.RetCode)
	fmt.Fprintf(&builder, ", RetInfo=%+v", d.RetInfo)
	fmt.Fprintf(&builder, ", FutureAccount=%+v", d.FutureAccount)
	fmt.Fprintf(&builder, ", TradeAmt=%+v", d.TradeAmt)
	fmt.Fprintf(&builder, ", CustFee=%+v", d.CustFee)
	fmt.Fprintf(&builder, ", CurrencyCode=%+v", d.CurrencyCode)

	builder.WriteByte('}')

	return builder.String()
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

func (d CThostFtdcTransferFutureToBankReqField) Type() string {
	return "CThostFtdcTransferFutureToBankReqField"
}

func (d CThostFtdcTransferFutureToBankReqField) String() string {
	var builder strings.Builder
	builder.Grow(163)

	builder.WriteString("CThostFtdcTransferFutureToBankReqField{")
	fmt.Fprintf(&builder, "FutureAccount=%+v", d.FutureAccount)
	fmt.Fprintf(&builder, ", FuturePwdFlag=%+v", d.FuturePwdFlag)
	fmt.Fprintf(&builder, ", FutureAccPwd=%+v", d.FutureAccPwd)
	fmt.Fprintf(&builder, ", TradeAmt=%+v", d.TradeAmt)
	fmt.Fprintf(&builder, ", CustFee=%+v", d.CustFee)
	fmt.Fprintf(&builder, ", CurrencyCode=%+v", d.CurrencyCode)

	builder.WriteByte('}')

	return builder.String()
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

func (d CThostFtdcTransferFutureToBankRspField) Type() string {
	return "CThostFtdcTransferFutureToBankRspField"
}

func (d CThostFtdcTransferFutureToBankRspField) String() string {
	var builder strings.Builder
	builder.Grow(152)

	builder.WriteString("CThostFtdcTransferFutureToBankRspField{")
	fmt.Fprintf(&builder, "RetCode=%+v", d.RetCode)
	fmt.Fprintf(&builder, ", RetInfo=%+v", d.RetInfo)
	fmt.Fprintf(&builder, ", FutureAccount=%+v", d.FutureAccount)
	fmt.Fprintf(&builder, ", TradeAmt=%+v", d.TradeAmt)
	fmt.Fprintf(&builder, ", CustFee=%+v", d.CustFee)
	fmt.Fprintf(&builder, ", CurrencyCode=%+v", d.CurrencyCode)

	builder.WriteByte('}')

	return builder.String()
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

func (d CThostFtdcTransferQryBankReqField) Type() string { return "CThostFtdcTransferQryBankReqField" }

func (d CThostFtdcTransferQryBankReqField) String() string {
	var builder strings.Builder
	builder.Grow(123)

	builder.WriteString("CThostFtdcTransferQryBankReqField{")
	fmt.Fprintf(&builder, "FutureAccount=%+v", d.FutureAccount)
	fmt.Fprintf(&builder, ", FuturePwdFlag=%+v", d.FuturePwdFlag)
	fmt.Fprintf(&builder, ", FutureAccPwd=%+v", d.FutureAccPwd)
	fmt.Fprintf(&builder, ", CurrencyCode=%+v", d.CurrencyCode)

	builder.WriteByte('}')

	return builder.String()
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

func (d CThostFtdcTransferQryBankRspField) Type() string { return "CThostFtdcTransferQryBankRspField" }

func (d CThostFtdcTransferQryBankRspField) String() string {
	var builder strings.Builder
	builder.Grow(164)

	builder.WriteString("CThostFtdcTransferQryBankRspField{")
	fmt.Fprintf(&builder, "RetCode=%+v", d.RetCode)
	fmt.Fprintf(&builder, ", RetInfo=%+v", d.RetInfo)
	fmt.Fprintf(&builder, ", FutureAccount=%+v", d.FutureAccount)
	fmt.Fprintf(&builder, ", TradeAmt=%+v", d.TradeAmt)
	fmt.Fprintf(&builder, ", UseAmt=%+v", d.UseAmt)
	fmt.Fprintf(&builder, ", FetchAmt=%+v", d.FetchAmt)
	fmt.Fprintf(&builder, ", CurrencyCode=%+v", d.CurrencyCode)

	builder.WriteByte('}')

	return builder.String()
}

// 查询银行交易明细请求，TradeCode=204999
type CThostFtdcTransferQryDetailReqField struct {
	// 期货资金账户
	FutureAccount types.TThostFtdcAccountIDType
}

func (d CThostFtdcTransferQryDetailReqField) Type() string {
	return "CThostFtdcTransferQryDetailReqField"
}

func (d CThostFtdcTransferQryDetailReqField) String() string {
	var builder strings.Builder
	builder.Grow(58)

	builder.WriteString("CThostFtdcTransferQryDetailReqField{")
	fmt.Fprintf(&builder, "FutureAccount=%+v", d.FutureAccount)

	builder.WriteByte('}')

	return builder.String()
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

func (d CThostFtdcTransferQryDetailRspField) Type() string {
	return "CThostFtdcTransferQryDetailRspField"
}

func (d CThostFtdcTransferQryDetailRspField) String() string {
	var builder strings.Builder
	builder.Grow(304)

	builder.WriteString("CThostFtdcTransferQryDetailRspField{")
	fmt.Fprintf(&builder, "TradeDate=%+v", d.TradeDate)
	fmt.Fprintf(&builder, ", TradeTime=%+v", d.TradeTime)
	fmt.Fprintf(&builder, ", TradeCode=%+v", d.TradeCode)
	fmt.Fprintf(&builder, ", FutureSerial=%+v", d.FutureSerial)
	fmt.Fprintf(&builder, ", FutureID=%+v", d.FutureID)
	fmt.Fprintf(&builder, ", FutureAccount=%+v", d.FutureAccount)
	fmt.Fprintf(&builder, ", BankSerial=%+v", d.BankSerial)
	fmt.Fprintf(&builder, ", BankID=%+v", d.BankID)
	fmt.Fprintf(&builder, ", BankBrchID=%+v", d.BankBrchID)
	fmt.Fprintf(&builder, ", BankAccount=%+v", d.BankAccount)
	fmt.Fprintf(&builder, ", CertCode=%+v", d.CertCode)
	fmt.Fprintf(&builder, ", CurrencyCode=%+v", d.CurrencyCode)
	fmt.Fprintf(&builder, ", TxAmount=%+v", d.TxAmount)
	fmt.Fprintf(&builder, ", Flag=%+v", d.Flag)

	builder.WriteByte('}')

	return builder.String()
}

// 响应信息
type CThostFtdcRspInfoField struct {
	// 错误代码
	ErrorID types.TThostFtdcErrorIDType
	// 错误信息
	ErrorMsg types.TThostFtdcErrorMsgType
}

func (d CThostFtdcRspInfoField) Type() string { return "CThostFtdcRspInfoField" }

func (d CThostFtdcRspInfoField) String() string {
	var builder strings.Builder
	builder.Grow(57)

	builder.WriteString("CThostFtdcRspInfoField{")
	fmt.Fprintf(&builder, "ErrorID=%+v", d.ErrorID)
	fmt.Fprintf(&builder, ", ErrorMsg=%+v", d.ErrorMsg)

	builder.WriteByte('}')

	return builder.String()
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

func (d CThostFtdcExchangeField) Type() string { return "CThostFtdcExchangeField" }

func (d CThostFtdcExchangeField) String() string {
	var builder strings.Builder
	builder.Grow(91)

	builder.WriteString("CThostFtdcExchangeField{")
	fmt.Fprintf(&builder, "ExchangeID=%+v", d.ExchangeID)
	fmt.Fprintf(&builder, ", ExchangeName=%+v", d.ExchangeName)
	fmt.Fprintf(&builder, ", ExchangeProperty=%+v", d.ExchangeProperty)

	builder.WriteByte('}')

	return builder.String()
}

// 产品
type CThostFtdcProductField struct {
	// 保留的无效字段
	reserve1 types.TThostFtdcOldInstrumentIDType
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
	// 保留的无效字段
	reserve2 types.TThostFtdcOldInstrumentIDType
	// 合约基础商品乘数
	UnderlyingMultiple types.TThostFtdcUnderlyingMultipleType
	// 产品代码
	ProductID types.TThostFtdcInstrumentIDType
	// 交易所产品代码
	ExchangeProductID types.TThostFtdcInstrumentIDType
	// 开仓量限制粒度
	OpenLimitControlLevel types.TThostFtdcOpenLimitControlLevelType
	// 报单频率控制粒度
	OrderFreqControlLevel types.TThostFtdcOrderFreqControlLevelType
}

func (d CThostFtdcProductField) Type() string { return "CThostFtdcProductField" }

func (d CThostFtdcProductField) String() string {
	var builder strings.Builder
	builder.Grow(544)

	builder.WriteString("CThostFtdcProductField{")
	fmt.Fprintf(&builder, "reserve1=%+v", d.reserve1)
	fmt.Fprintf(&builder, ", ProductName=%+v", d.ProductName)
	fmt.Fprintf(&builder, ", ExchangeID=%+v", d.ExchangeID)
	fmt.Fprintf(&builder, ", ProductClass=%+v", d.ProductClass)
	fmt.Fprintf(&builder, ", VolumeMultiple=%+v", d.VolumeMultiple)
	fmt.Fprintf(&builder, ", PriceTick=%+v", d.PriceTick)
	fmt.Fprintf(&builder, ", MaxMarketOrderVolume=%+v", d.MaxMarketOrderVolume)
	fmt.Fprintf(&builder, ", MinMarketOrderVolume=%+v", d.MinMarketOrderVolume)
	fmt.Fprintf(&builder, ", MaxLimitOrderVolume=%+v", d.MaxLimitOrderVolume)
	fmt.Fprintf(&builder, ", MinLimitOrderVolume=%+v", d.MinLimitOrderVolume)
	fmt.Fprintf(&builder, ", PositionType=%+v", d.PositionType)
	fmt.Fprintf(&builder, ", PositionDateType=%+v", d.PositionDateType)
	fmt.Fprintf(&builder, ", CloseDealType=%+v", d.CloseDealType)
	fmt.Fprintf(&builder, ", TradeCurrencyID=%+v", d.TradeCurrencyID)
	fmt.Fprintf(&builder, ", MortgageFundUseRange=%+v", d.MortgageFundUseRange)
	fmt.Fprintf(&builder, ", reserve2=%+v", d.reserve2)
	fmt.Fprintf(&builder, ", UnderlyingMultiple=%+v", d.UnderlyingMultiple)
	fmt.Fprintf(&builder, ", ProductID=%+v", d.ProductID)
	fmt.Fprintf(&builder, ", ExchangeProductID=%+v", d.ExchangeProductID)
	fmt.Fprintf(&builder, ", OpenLimitControlLevel=%+v", d.OpenLimitControlLevel)
	fmt.Fprintf(&builder, ", OrderFreqControlLevel=%+v", d.OrderFreqControlLevel)

	builder.WriteByte('}')

	return builder.String()
}

// 合约
type CThostFtdcInstrumentField struct {
	// 保留的无效字段
	reserve1 types.TThostFtdcOldInstrumentIDType
	// 交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	// 合约名称
	InstrumentName types.TThostFtdcInstrumentNameType
	// 保留的无效字段
	reserve2 types.TThostFtdcOldExchangeInstIDType
	// 保留的无效字段
	reserve3 types.TThostFtdcOldInstrumentIDType
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
	// 保留的无效字段
	reserve4 types.TThostFtdcOldInstrumentIDType
	// 执行价
	StrikePrice types.TThostFtdcPriceType
	// 期权类型
	OptionsType types.TThostFtdcOptionsTypeType
	// 合约基础商品乘数
	UnderlyingMultiple types.TThostFtdcUnderlyingMultipleType
	// 组合类型
	CombinationType types.TThostFtdcCombinationTypeType
	// 合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
	// 合约在交易所的代码
	ExchangeInstID types.TThostFtdcExchangeInstIDType
	// 产品代码
	ProductID types.TThostFtdcInstrumentIDType
	// 基础商品代码
	UnderlyingInstrID types.TThostFtdcInstrumentIDType
}

func (d CThostFtdcInstrumentField) Type() string { return "CThostFtdcInstrumentField" }

func (d CThostFtdcInstrumentField) String() string {
	var builder strings.Builder
	builder.Grow(833)

	builder.WriteString("CThostFtdcInstrumentField{")
	fmt.Fprintf(&builder, "reserve1=%+v", d.reserve1)
	fmt.Fprintf(&builder, ", ExchangeID=%+v", d.ExchangeID)
	fmt.Fprintf(&builder, ", InstrumentName=%+v", d.InstrumentName)
	fmt.Fprintf(&builder, ", reserve2=%+v", d.reserve2)
	fmt.Fprintf(&builder, ", reserve3=%+v", d.reserve3)
	fmt.Fprintf(&builder, ", ProductClass=%+v", d.ProductClass)
	fmt.Fprintf(&builder, ", DeliveryYear=%+v", d.DeliveryYear)
	fmt.Fprintf(&builder, ", DeliveryMonth=%+v", d.DeliveryMonth)
	fmt.Fprintf(&builder, ", MaxMarketOrderVolume=%+v", d.MaxMarketOrderVolume)
	fmt.Fprintf(&builder, ", MinMarketOrderVolume=%+v", d.MinMarketOrderVolume)
	fmt.Fprintf(&builder, ", MaxLimitOrderVolume=%+v", d.MaxLimitOrderVolume)
	fmt.Fprintf(&builder, ", MinLimitOrderVolume=%+v", d.MinLimitOrderVolume)
	fmt.Fprintf(&builder, ", VolumeMultiple=%+v", d.VolumeMultiple)
	fmt.Fprintf(&builder, ", PriceTick=%+v", d.PriceTick)
	fmt.Fprintf(&builder, ", CreateDate=%+v", d.CreateDate)
	fmt.Fprintf(&builder, ", OpenDate=%+v", d.OpenDate)
	fmt.Fprintf(&builder, ", ExpireDate=%+v", d.ExpireDate)
	fmt.Fprintf(&builder, ", StartDelivDate=%+v", d.StartDelivDate)
	fmt.Fprintf(&builder, ", EndDelivDate=%+v", d.EndDelivDate)
	fmt.Fprintf(&builder, ", InstLifePhase=%+v", d.InstLifePhase)
	fmt.Fprintf(&builder, ", IsTrading=%+v", d.IsTrading)
	fmt.Fprintf(&builder, ", PositionType=%+v", d.PositionType)
	fmt.Fprintf(&builder, ", PositionDateType=%+v", d.PositionDateType)
	fmt.Fprintf(&builder, ", LongMarginRatio=%+v", d.LongMarginRatio)
	fmt.Fprintf(&builder, ", ShortMarginRatio=%+v", d.ShortMarginRatio)
	fmt.Fprintf(&builder, ", MaxMarginSideAlgorithm=%+v", d.MaxMarginSideAlgorithm)
	fmt.Fprintf(&builder, ", reserve4=%+v", d.reserve4)
	fmt.Fprintf(&builder, ", StrikePrice=%+v", d.StrikePrice)
	fmt.Fprintf(&builder, ", OptionsType=%+v", d.OptionsType)
	fmt.Fprintf(&builder, ", UnderlyingMultiple=%+v", d.UnderlyingMultiple)
	fmt.Fprintf(&builder, ", CombinationType=%+v", d.CombinationType)
	fmt.Fprintf(&builder, ", InstrumentID=%+v", d.InstrumentID)
	fmt.Fprintf(&builder, ", ExchangeInstID=%+v", d.ExchangeInstID)
	fmt.Fprintf(&builder, ", ProductID=%+v", d.ProductID)
	fmt.Fprintf(&builder, ", UnderlyingInstrID=%+v", d.UnderlyingInstrID)

	builder.WriteByte('}')

	return builder.String()
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

func (d CThostFtdcBrokerField) Type() string { return "CThostFtdcBrokerField" }

func (d CThostFtdcBrokerField) String() string {
	var builder strings.Builder
	builder.Grow(97)

	builder.WriteString("CThostFtdcBrokerField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", BrokerAbbr=%+v", d.BrokerAbbr)
	fmt.Fprintf(&builder, ", BrokerName=%+v", d.BrokerName)
	fmt.Fprintf(&builder, ", IsActive=%+v", d.IsActive)

	builder.WriteByte('}')

	return builder.String()
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
	// 撤单时选择席位算法
	OrderCancelAlg types.TThostFtdcOrderCancelAlgType
	// 交易报盘安装数量
	TradeInstallCount types.TThostFtdcInstallCountType
	// 行情报盘安装数量
	MDInstallCount types.TThostFtdcInstallCountType
}

func (d CThostFtdcTraderField) Type() string { return "CThostFtdcTraderField" }

func (d CThostFtdcTraderField) String() string {
	var builder strings.Builder
	builder.Grow(215)

	builder.WriteString("CThostFtdcTraderField{")
	fmt.Fprintf(&builder, "ExchangeID=%+v", d.ExchangeID)
	fmt.Fprintf(&builder, ", TraderID=%+v", d.TraderID)
	fmt.Fprintf(&builder, ", ParticipantID=%+v", d.ParticipantID)
	fmt.Fprintf(&builder, ", Password=%+v", d.Password)
	fmt.Fprintf(&builder, ", InstallCount=%+v", d.InstallCount)
	fmt.Fprintf(&builder, ", BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", OrderCancelAlg=%+v", d.OrderCancelAlg)
	fmt.Fprintf(&builder, ", TradeInstallCount=%+v", d.TradeInstallCount)
	fmt.Fprintf(&builder, ", MDInstallCount=%+v", d.MDInstallCount)

	builder.WriteByte('}')

	return builder.String()
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
	// 是否频率控制
	IsOrderFreq types.TThostFtdcEnumBoolType
	// 是否开仓限制
	IsOpenVolLimit types.TThostFtdcEnumBoolType
}

func (d CThostFtdcInvestorField) Type() string { return "CThostFtdcInvestorField" }

func (d CThostFtdcInvestorField) String() string {
	var builder strings.Builder
	builder.Grow(339)

	builder.WriteString("CThostFtdcInvestorField{")
	fmt.Fprintf(&builder, "InvestorID=%+v", d.InvestorID)
	fmt.Fprintf(&builder, ", BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", InvestorGroupID=%+v", d.InvestorGroupID)
	fmt.Fprintf(&builder, ", InvestorName=%+v", d.InvestorName)
	fmt.Fprintf(&builder, ", IdentifiedCardType=%+v", d.IdentifiedCardType)
	fmt.Fprintf(&builder, ", IdentifiedCardNo=%+v", d.IdentifiedCardNo)
	fmt.Fprintf(&builder, ", IsActive=%+v", d.IsActive)
	fmt.Fprintf(&builder, ", Telephone=%+v", d.Telephone)
	fmt.Fprintf(&builder, ", Address=%+v", d.Address)
	fmt.Fprintf(&builder, ", OpenDate=%+v", d.OpenDate)
	fmt.Fprintf(&builder, ", Mobile=%+v", d.Mobile)
	fmt.Fprintf(&builder, ", CommModelID=%+v", d.CommModelID)
	fmt.Fprintf(&builder, ", MarginModelID=%+v", d.MarginModelID)
	fmt.Fprintf(&builder, ", IsOrderFreq=%+v", d.IsOrderFreq)
	fmt.Fprintf(&builder, ", IsOpenVolLimit=%+v", d.IsOpenVolLimit)

	builder.WriteByte('}')

	return builder.String()
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
	// 营业部编号
	BranchID types.TThostFtdcBranchIDType
	// 业务类型
	BizType types.TThostFtdcBizTypeType
	// 投资单元代码
	InvestUnitID types.TThostFtdcInvestUnitIDType
}

func (d CThostFtdcTradingCodeField) Type() string { return "CThostFtdcTradingCodeField" }

func (d CThostFtdcTradingCodeField) String() string {
	var builder strings.Builder
	builder.Grow(199)

	builder.WriteString("CThostFtdcTradingCodeField{")
	fmt.Fprintf(&builder, "InvestorID=%+v", d.InvestorID)
	fmt.Fprintf(&builder, ", BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", ExchangeID=%+v", d.ExchangeID)
	fmt.Fprintf(&builder, ", ClientID=%+v", d.ClientID)
	fmt.Fprintf(&builder, ", IsActive=%+v", d.IsActive)
	fmt.Fprintf(&builder, ", ClientIDType=%+v", d.ClientIDType)
	fmt.Fprintf(&builder, ", BranchID=%+v", d.BranchID)
	fmt.Fprintf(&builder, ", BizType=%+v", d.BizType)
	fmt.Fprintf(&builder, ", InvestUnitID=%+v", d.InvestUnitID)

	builder.WriteByte('}')

	return builder.String()
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

func (d CThostFtdcPartBrokerField) Type() string { return "CThostFtdcPartBrokerField" }

func (d CThostFtdcPartBrokerField) String() string {
	var builder strings.Builder
	builder.Grow(104)

	builder.WriteString("CThostFtdcPartBrokerField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", ExchangeID=%+v", d.ExchangeID)
	fmt.Fprintf(&builder, ", ParticipantID=%+v", d.ParticipantID)
	fmt.Fprintf(&builder, ", IsActive=%+v", d.IsActive)

	builder.WriteByte('}')

	return builder.String()
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

func (d CThostFtdcSuperUserField) Type() string { return "CThostFtdcSuperUserField" }

func (d CThostFtdcSuperUserField) String() string {
	var builder strings.Builder
	builder.Grow(94)

	builder.WriteString("CThostFtdcSuperUserField{")
	fmt.Fprintf(&builder, "UserID=%+v", d.UserID)
	fmt.Fprintf(&builder, ", UserName=%+v", d.UserName)
	fmt.Fprintf(&builder, ", Password=%+v", d.Password)
	fmt.Fprintf(&builder, ", IsActive=%+v", d.IsActive)

	builder.WriteByte('}')

	return builder.String()
}

// 管理用户功能权限
type CThostFtdcSuperUserFunctionField struct {
	// 用户代码
	UserID types.TThostFtdcUserIDType
	// 功能代码
	FunctionCode types.TThostFtdcFunctionCodeType
}

func (d CThostFtdcSuperUserFunctionField) Type() string { return "CThostFtdcSuperUserFunctionField" }

func (d CThostFtdcSuperUserFunctionField) String() string {
	var builder strings.Builder
	builder.Grow(70)

	builder.WriteString("CThostFtdcSuperUserFunctionField{")
	fmt.Fprintf(&builder, "UserID=%+v", d.UserID)
	fmt.Fprintf(&builder, ", FunctionCode=%+v", d.FunctionCode)

	builder.WriteByte('}')

	return builder.String()
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

func (d CThostFtdcInvestorGroupField) Type() string { return "CThostFtdcInvestorGroupField" }

func (d CThostFtdcInvestorGroupField) String() string {
	var builder strings.Builder
	builder.Grow(98)

	builder.WriteString("CThostFtdcInvestorGroupField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", InvestorGroupID=%+v", d.InvestorGroupID)
	fmt.Fprintf(&builder, ", InvestorGroupName=%+v", d.InvestorGroupName)

	builder.WriteByte('}')

	return builder.String()
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
	// 业务类型
	BizType types.TThostFtdcBizTypeType
	// 延时换汇冻结金额
	FrozenSwap types.TThostFtdcMoneyType
	// 剩余换汇额度
	RemainSwap types.TThostFtdcMoneyType
	// 期权市值
	OptionValue types.TThostFtdcMoneyType
}

func (d CThostFtdcTradingAccountField) Type() string { return "CThostFtdcTradingAccountField" }

func (d CThostFtdcTradingAccountField) String() string {
	var builder strings.Builder
	builder.Grow(1194)

	builder.WriteString("CThostFtdcTradingAccountField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", AccountID=%+v", d.AccountID)
	fmt.Fprintf(&builder, ", PreMortgage=%+v", d.PreMortgage)
	fmt.Fprintf(&builder, ", PreCredit=%+v", d.PreCredit)
	fmt.Fprintf(&builder, ", PreDeposit=%+v", d.PreDeposit)
	fmt.Fprintf(&builder, ", PreBalance=%+v", d.PreBalance)
	fmt.Fprintf(&builder, ", PreMargin=%+v", d.PreMargin)
	fmt.Fprintf(&builder, ", InterestBase=%+v", d.InterestBase)
	fmt.Fprintf(&builder, ", Interest=%+v", d.Interest)
	fmt.Fprintf(&builder, ", Deposit=%+v", d.Deposit)
	fmt.Fprintf(&builder, ", Withdraw=%+v", d.Withdraw)
	fmt.Fprintf(&builder, ", FrozenMargin=%+v", d.FrozenMargin)
	fmt.Fprintf(&builder, ", FrozenCash=%+v", d.FrozenCash)
	fmt.Fprintf(&builder, ", FrozenCommission=%+v", d.FrozenCommission)
	fmt.Fprintf(&builder, ", CurrMargin=%+v", d.CurrMargin)
	fmt.Fprintf(&builder, ", CashIn=%+v", d.CashIn)
	fmt.Fprintf(&builder, ", Commission=%+v", d.Commission)
	fmt.Fprintf(&builder, ", CloseProfit=%+v", d.CloseProfit)
	fmt.Fprintf(&builder, ", PositionProfit=%+v", d.PositionProfit)
	fmt.Fprintf(&builder, ", Balance=%+v", d.Balance)
	fmt.Fprintf(&builder, ", Available=%+v", d.Available)
	fmt.Fprintf(&builder, ", WithdrawQuota=%+v", d.WithdrawQuota)
	fmt.Fprintf(&builder, ", Reserve=%+v", d.Reserve)
	fmt.Fprintf(&builder, ", TradingDay=%+v", d.TradingDay)
	fmt.Fprintf(&builder, ", SettlementID=%+v", d.SettlementID)
	fmt.Fprintf(&builder, ", Credit=%+v", d.Credit)
	fmt.Fprintf(&builder, ", Mortgage=%+v", d.Mortgage)
	fmt.Fprintf(&builder, ", ExchangeMargin=%+v", d.ExchangeMargin)
	fmt.Fprintf(&builder, ", DeliveryMargin=%+v", d.DeliveryMargin)
	fmt.Fprintf(&builder, ", ExchangeDeliveryMargin=%+v", d.ExchangeDeliveryMargin)
	fmt.Fprintf(&builder, ", ReserveBalance=%+v", d.ReserveBalance)
	fmt.Fprintf(&builder, ", CurrencyID=%+v", d.CurrencyID)
	fmt.Fprintf(&builder, ", PreFundMortgageIn=%+v", d.PreFundMortgageIn)
	fmt.Fprintf(&builder, ", PreFundMortgageOut=%+v", d.PreFundMortgageOut)
	fmt.Fprintf(&builder, ", FundMortgageIn=%+v", d.FundMortgageIn)
	fmt.Fprintf(&builder, ", FundMortgageOut=%+v", d.FundMortgageOut)
	fmt.Fprintf(&builder, ", FundMortgageAvailable=%+v", d.FundMortgageAvailable)
	fmt.Fprintf(&builder, ", MortgageableFund=%+v", d.MortgageableFund)
	fmt.Fprintf(&builder, ", SpecProductMargin=%+v", d.SpecProductMargin)
	fmt.Fprintf(&builder, ", SpecProductFrozenMargin=%+v", d.SpecProductFrozenMargin)
	fmt.Fprintf(&builder, ", SpecProductCommission=%+v", d.SpecProductCommission)
	fmt.Fprintf(&builder, ", SpecProductFrozenCommission=%+v", d.SpecProductFrozenCommission)
	fmt.Fprintf(&builder, ", SpecProductPositionProfit=%+v", d.SpecProductPositionProfit)
	fmt.Fprintf(&builder, ", SpecProductCloseProfit=%+v", d.SpecProductCloseProfit)
	fmt.Fprintf(&builder, ", SpecProductPositionProfitByAlg=%+v", d.SpecProductPositionProfitByAlg)
	fmt.Fprintf(&builder, ", SpecProductExchangeMargin=%+v", d.SpecProductExchangeMargin)
	fmt.Fprintf(&builder, ", BizType=%+v", d.BizType)
	fmt.Fprintf(&builder, ", FrozenSwap=%+v", d.FrozenSwap)
	fmt.Fprintf(&builder, ", RemainSwap=%+v", d.RemainSwap)
	fmt.Fprintf(&builder, ", OptionValue=%+v", d.OptionValue)

	builder.WriteByte('}')

	return builder.String()
}

// 投资者持仓
type CThostFtdcInvestorPositionField struct {
	// 保留的无效字段
	reserve1 types.TThostFtdcOldInstrumentIDType
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
	// 交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	// 执行冻结的昨仓
	YdStrikeFrozen types.TThostFtdcVolumeType
	// 投资单元代码
	InvestUnitID types.TThostFtdcInvestUnitIDType
	// 持仓成本差值
	PositionCostOffset types.TThostFtdcMoneyType
	// tas持仓手数
	TasPosition types.TThostFtdcVolumeType
	// tas持仓成本
	TasPositionCost types.TThostFtdcMoneyType
	// 合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
	// 期权市值
	OptionValue types.TThostFtdcMoneyType
}

func (d CThostFtdcInvestorPositionField) Type() string { return "CThostFtdcInvestorPositionField" }

func (d CThostFtdcInvestorPositionField) String() string {
	var builder strings.Builder
	builder.Grow(1171)

	builder.WriteString("CThostFtdcInvestorPositionField{")
	fmt.Fprintf(&builder, "reserve1=%+v", d.reserve1)
	fmt.Fprintf(&builder, ", BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", InvestorID=%+v", d.InvestorID)
	fmt.Fprintf(&builder, ", PosiDirection=%+v", d.PosiDirection)
	fmt.Fprintf(&builder, ", HedgeFlag=%+v", d.HedgeFlag)
	fmt.Fprintf(&builder, ", PositionDate=%+v", d.PositionDate)
	fmt.Fprintf(&builder, ", YdPosition=%+v", d.YdPosition)
	fmt.Fprintf(&builder, ", Position=%+v", d.Position)
	fmt.Fprintf(&builder, ", LongFrozen=%+v", d.LongFrozen)
	fmt.Fprintf(&builder, ", ShortFrozen=%+v", d.ShortFrozen)
	fmt.Fprintf(&builder, ", LongFrozenAmount=%+v", d.LongFrozenAmount)
	fmt.Fprintf(&builder, ", ShortFrozenAmount=%+v", d.ShortFrozenAmount)
	fmt.Fprintf(&builder, ", OpenVolume=%+v", d.OpenVolume)
	fmt.Fprintf(&builder, ", CloseVolume=%+v", d.CloseVolume)
	fmt.Fprintf(&builder, ", OpenAmount=%+v", d.OpenAmount)
	fmt.Fprintf(&builder, ", CloseAmount=%+v", d.CloseAmount)
	fmt.Fprintf(&builder, ", PositionCost=%+v", d.PositionCost)
	fmt.Fprintf(&builder, ", PreMargin=%+v", d.PreMargin)
	fmt.Fprintf(&builder, ", UseMargin=%+v", d.UseMargin)
	fmt.Fprintf(&builder, ", FrozenMargin=%+v", d.FrozenMargin)
	fmt.Fprintf(&builder, ", FrozenCash=%+v", d.FrozenCash)
	fmt.Fprintf(&builder, ", FrozenCommission=%+v", d.FrozenCommission)
	fmt.Fprintf(&builder, ", CashIn=%+v", d.CashIn)
	fmt.Fprintf(&builder, ", Commission=%+v", d.Commission)
	fmt.Fprintf(&builder, ", CloseProfit=%+v", d.CloseProfit)
	fmt.Fprintf(&builder, ", PositionProfit=%+v", d.PositionProfit)
	fmt.Fprintf(&builder, ", PreSettlementPrice=%+v", d.PreSettlementPrice)
	fmt.Fprintf(&builder, ", SettlementPrice=%+v", d.SettlementPrice)
	fmt.Fprintf(&builder, ", TradingDay=%+v", d.TradingDay)
	fmt.Fprintf(&builder, ", SettlementID=%+v", d.SettlementID)
	fmt.Fprintf(&builder, ", OpenCost=%+v", d.OpenCost)
	fmt.Fprintf(&builder, ", ExchangeMargin=%+v", d.ExchangeMargin)
	fmt.Fprintf(&builder, ", CombPosition=%+v", d.CombPosition)
	fmt.Fprintf(&builder, ", CombLongFrozen=%+v", d.CombLongFrozen)
	fmt.Fprintf(&builder, ", CombShortFrozen=%+v", d.CombShortFrozen)
	fmt.Fprintf(&builder, ", CloseProfitByDate=%+v", d.CloseProfitByDate)
	fmt.Fprintf(&builder, ", CloseProfitByTrade=%+v", d.CloseProfitByTrade)
	fmt.Fprintf(&builder, ", TodayPosition=%+v", d.TodayPosition)
	fmt.Fprintf(&builder, ", MarginRateByMoney=%+v", d.MarginRateByMoney)
	fmt.Fprintf(&builder, ", MarginRateByVolume=%+v", d.MarginRateByVolume)
	fmt.Fprintf(&builder, ", StrikeFrozen=%+v", d.StrikeFrozen)
	fmt.Fprintf(&builder, ", StrikeFrozenAmount=%+v", d.StrikeFrozenAmount)
	fmt.Fprintf(&builder, ", AbandonFrozen=%+v", d.AbandonFrozen)
	fmt.Fprintf(&builder, ", ExchangeID=%+v", d.ExchangeID)
	fmt.Fprintf(&builder, ", YdStrikeFrozen=%+v", d.YdStrikeFrozen)
	fmt.Fprintf(&builder, ", InvestUnitID=%+v", d.InvestUnitID)
	fmt.Fprintf(&builder, ", PositionCostOffset=%+v", d.PositionCostOffset)
	fmt.Fprintf(&builder, ", TasPosition=%+v", d.TasPosition)
	fmt.Fprintf(&builder, ", TasPositionCost=%+v", d.TasPositionCost)
	fmt.Fprintf(&builder, ", InstrumentID=%+v", d.InstrumentID)
	fmt.Fprintf(&builder, ", OptionValue=%+v", d.OptionValue)

	builder.WriteByte('}')

	return builder.String()
}

// 合约保证金率
type CThostFtdcInstrumentMarginRateField struct {
	// 保留的无效字段
	reserve1 types.TThostFtdcOldInstrumentIDType
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
	// 交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	// 投资单元代码
	InvestUnitID types.TThostFtdcInvestUnitIDType
	// 合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
}

func (d CThostFtdcInstrumentMarginRateField) Type() string {
	return "CThostFtdcInstrumentMarginRateField"
}

func (d CThostFtdcInstrumentMarginRateField) String() string {
	var builder strings.Builder
	builder.Grow(349)

	builder.WriteString("CThostFtdcInstrumentMarginRateField{")
	fmt.Fprintf(&builder, "reserve1=%+v", d.reserve1)
	fmt.Fprintf(&builder, ", InvestorRange=%+v", d.InvestorRange)
	fmt.Fprintf(&builder, ", BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", InvestorID=%+v", d.InvestorID)
	fmt.Fprintf(&builder, ", HedgeFlag=%+v", d.HedgeFlag)
	fmt.Fprintf(&builder, ", LongMarginRatioByMoney=%+v", d.LongMarginRatioByMoney)
	fmt.Fprintf(&builder, ", LongMarginRatioByVolume=%+v", d.LongMarginRatioByVolume)
	fmt.Fprintf(&builder, ", ShortMarginRatioByMoney=%+v", d.ShortMarginRatioByMoney)
	fmt.Fprintf(&builder, ", ShortMarginRatioByVolume=%+v", d.ShortMarginRatioByVolume)
	fmt.Fprintf(&builder, ", IsRelative=%+v", d.IsRelative)
	fmt.Fprintf(&builder, ", ExchangeID=%+v", d.ExchangeID)
	fmt.Fprintf(&builder, ", InvestUnitID=%+v", d.InvestUnitID)
	fmt.Fprintf(&builder, ", InstrumentID=%+v", d.InstrumentID)

	builder.WriteByte('}')

	return builder.String()
}

// 合约手续费率
type CThostFtdcInstrumentCommissionRateField struct {
	// 保留的无效字段
	reserve1 types.TThostFtdcOldInstrumentIDType
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
	// 交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	// 业务类型
	BizType types.TThostFtdcBizTypeType
	// 投资单元代码
	InvestUnitID types.TThostFtdcInvestUnitIDType
	// 合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
}

func (d CThostFtdcInstrumentCommissionRateField) Type() string {
	return "CThostFtdcInstrumentCommissionRateField"
}

func (d CThostFtdcInstrumentCommissionRateField) String() string {
	var builder strings.Builder
	builder.Grow(372)

	builder.WriteString("CThostFtdcInstrumentCommissionRateField{")
	fmt.Fprintf(&builder, "reserve1=%+v", d.reserve1)
	fmt.Fprintf(&builder, ", InvestorRange=%+v", d.InvestorRange)
	fmt.Fprintf(&builder, ", BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", InvestorID=%+v", d.InvestorID)
	fmt.Fprintf(&builder, ", OpenRatioByMoney=%+v", d.OpenRatioByMoney)
	fmt.Fprintf(&builder, ", OpenRatioByVolume=%+v", d.OpenRatioByVolume)
	fmt.Fprintf(&builder, ", CloseRatioByMoney=%+v", d.CloseRatioByMoney)
	fmt.Fprintf(&builder, ", CloseRatioByVolume=%+v", d.CloseRatioByVolume)
	fmt.Fprintf(&builder, ", CloseTodayRatioByMoney=%+v", d.CloseTodayRatioByMoney)
	fmt.Fprintf(&builder, ", CloseTodayRatioByVolume=%+v", d.CloseTodayRatioByVolume)
	fmt.Fprintf(&builder, ", ExchangeID=%+v", d.ExchangeID)
	fmt.Fprintf(&builder, ", BizType=%+v", d.BizType)
	fmt.Fprintf(&builder, ", InvestUnitID=%+v", d.InvestUnitID)
	fmt.Fprintf(&builder, ", InstrumentID=%+v", d.InstrumentID)

	builder.WriteByte('}')

	return builder.String()
}

// 深度行情
type CThostFtdcDepthMarketDataField struct {
	// 交易日
	TradingDay types.TThostFtdcDateType
	// 保留的无效字段
	reserve1 types.TThostFtdcOldInstrumentIDType
	// 交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	// 保留的无效字段
	reserve2 types.TThostFtdcOldExchangeInstIDType
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
	// 合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
	// 合约在交易所的代码
	ExchangeInstID types.TThostFtdcExchangeInstIDType
	// 上带价
	BandingUpperPrice types.TThostFtdcPriceType
	// 下带价
	BandingLowerPrice types.TThostFtdcPriceType
}

func (d CThostFtdcDepthMarketDataField) Type() string { return "CThostFtdcDepthMarketDataField" }

func (d CThostFtdcDepthMarketDataField) String() string {
	var builder strings.Builder
	builder.Grow(1026)

	builder.WriteString("CThostFtdcDepthMarketDataField{")
	fmt.Fprintf(&builder, "TradingDay=%+v", d.TradingDay)
	fmt.Fprintf(&builder, ", reserve1=%+v", d.reserve1)
	fmt.Fprintf(&builder, ", ExchangeID=%+v", d.ExchangeID)
	fmt.Fprintf(&builder, ", reserve2=%+v", d.reserve2)
	fmt.Fprintf(&builder, ", LastPrice=%+v", d.LastPrice)
	fmt.Fprintf(&builder, ", PreSettlementPrice=%+v", d.PreSettlementPrice)
	fmt.Fprintf(&builder, ", PreClosePrice=%+v", d.PreClosePrice)
	fmt.Fprintf(&builder, ", PreOpenInterest=%+v", d.PreOpenInterest)
	fmt.Fprintf(&builder, ", OpenPrice=%+v", d.OpenPrice)
	fmt.Fprintf(&builder, ", HighestPrice=%+v", d.HighestPrice)
	fmt.Fprintf(&builder, ", LowestPrice=%+v", d.LowestPrice)
	fmt.Fprintf(&builder, ", Volume=%+v", d.Volume)
	fmt.Fprintf(&builder, ", Turnover=%+v", d.Turnover)
	fmt.Fprintf(&builder, ", OpenInterest=%+v", d.OpenInterest)
	fmt.Fprintf(&builder, ", ClosePrice=%+v", d.ClosePrice)
	fmt.Fprintf(&builder, ", SettlementPrice=%+v", d.SettlementPrice)
	fmt.Fprintf(&builder, ", UpperLimitPrice=%+v", d.UpperLimitPrice)
	fmt.Fprintf(&builder, ", LowerLimitPrice=%+v", d.LowerLimitPrice)
	fmt.Fprintf(&builder, ", PreDelta=%+v", d.PreDelta)
	fmt.Fprintf(&builder, ", CurrDelta=%+v", d.CurrDelta)
	fmt.Fprintf(&builder, ", UpdateTime=%+v", d.UpdateTime)
	fmt.Fprintf(&builder, ", UpdateMillisec=%+v", d.UpdateMillisec)
	fmt.Fprintf(&builder, ", BidPrice1=%+v", d.BidPrice1)
	fmt.Fprintf(&builder, ", BidVolume1=%+v", d.BidVolume1)
	fmt.Fprintf(&builder, ", AskPrice1=%+v", d.AskPrice1)
	fmt.Fprintf(&builder, ", AskVolume1=%+v", d.AskVolume1)
	fmt.Fprintf(&builder, ", BidPrice2=%+v", d.BidPrice2)
	fmt.Fprintf(&builder, ", BidVolume2=%+v", d.BidVolume2)
	fmt.Fprintf(&builder, ", AskPrice2=%+v", d.AskPrice2)
	fmt.Fprintf(&builder, ", AskVolume2=%+v", d.AskVolume2)
	fmt.Fprintf(&builder, ", BidPrice3=%+v", d.BidPrice3)
	fmt.Fprintf(&builder, ", BidVolume3=%+v", d.BidVolume3)
	fmt.Fprintf(&builder, ", AskPrice3=%+v", d.AskPrice3)
	fmt.Fprintf(&builder, ", AskVolume3=%+v", d.AskVolume3)
	fmt.Fprintf(&builder, ", BidPrice4=%+v", d.BidPrice4)
	fmt.Fprintf(&builder, ", BidVolume4=%+v", d.BidVolume4)
	fmt.Fprintf(&builder, ", AskPrice4=%+v", d.AskPrice4)
	fmt.Fprintf(&builder, ", AskVolume4=%+v", d.AskVolume4)
	fmt.Fprintf(&builder, ", BidPrice5=%+v", d.BidPrice5)
	fmt.Fprintf(&builder, ", BidVolume5=%+v", d.BidVolume5)
	fmt.Fprintf(&builder, ", AskPrice5=%+v", d.AskPrice5)
	fmt.Fprintf(&builder, ", AskVolume5=%+v", d.AskVolume5)
	fmt.Fprintf(&builder, ", AveragePrice=%+v", d.AveragePrice)
	fmt.Fprintf(&builder, ", ActionDay=%+v", d.ActionDay)
	fmt.Fprintf(&builder, ", InstrumentID=%+v", d.InstrumentID)
	fmt.Fprintf(&builder, ", ExchangeInstID=%+v", d.ExchangeInstID)
	fmt.Fprintf(&builder, ", BandingUpperPrice=%+v", d.BandingUpperPrice)
	fmt.Fprintf(&builder, ", BandingLowerPrice=%+v", d.BandingLowerPrice)

	builder.WriteByte('}')

	return builder.String()
}

// 投资者合约交易权限
type CThostFtdcInstrumentTradingRightField struct {
	// 保留的无效字段
	reserve1 types.TThostFtdcOldInstrumentIDType
	// 投资者范围
	InvestorRange types.TThostFtdcInvestorRangeType
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	// 交易权限
	TradingRight types.TThostFtdcTradingRightType
	// 合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
}

func (d CThostFtdcInstrumentTradingRightField) Type() string {
	return "CThostFtdcInstrumentTradingRightField"
}

func (d CThostFtdcInstrumentTradingRightField) String() string {
	var builder strings.Builder
	builder.Grow(160)

	builder.WriteString("CThostFtdcInstrumentTradingRightField{")
	fmt.Fprintf(&builder, "reserve1=%+v", d.reserve1)
	fmt.Fprintf(&builder, ", InvestorRange=%+v", d.InvestorRange)
	fmt.Fprintf(&builder, ", BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", InvestorID=%+v", d.InvestorID)
	fmt.Fprintf(&builder, ", TradingRight=%+v", d.TradingRight)
	fmt.Fprintf(&builder, ", InstrumentID=%+v", d.InstrumentID)

	builder.WriteByte('}')

	return builder.String()
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
	// 是否强制终端认证
	IsAuthForce types.TThostFtdcBoolType
}

func (d CThostFtdcBrokerUserField) Type() string { return "CThostFtdcBrokerUserField" }

func (d CThostFtdcBrokerUserField) String() string {
	var builder strings.Builder
	builder.Grow(154)

	builder.WriteString("CThostFtdcBrokerUserField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", UserID=%+v", d.UserID)
	fmt.Fprintf(&builder, ", UserName=%+v", d.UserName)
	fmt.Fprintf(&builder, ", UserType=%+v", d.UserType)
	fmt.Fprintf(&builder, ", IsActive=%+v", d.IsActive)
	fmt.Fprintf(&builder, ", IsUsingOTP=%+v", d.IsUsingOTP)
	fmt.Fprintf(&builder, ", IsAuthForce=%+v", d.IsAuthForce)

	builder.WriteByte('}')

	return builder.String()
}

// 经纪公司用户口令
type CThostFtdcBrokerUserPasswordField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 用户代码
	UserID types.TThostFtdcUserIDType
	// 密码
	Password types.TThostFtdcPasswordType
	// 上次修改时间
	LastUpdateTime types.TThostFtdcDateTimeType
	// 上次登陆时间
	LastLoginTime types.TThostFtdcDateTimeType
	// 密码过期时间
	ExpireDate types.TThostFtdcDateType
	// 弱密码过期时间
	WeakExpireDate types.TThostFtdcDateType
}

func (d CThostFtdcBrokerUserPasswordField) Type() string { return "CThostFtdcBrokerUserPasswordField" }

func (d CThostFtdcBrokerUserPasswordField) String() string {
	var builder strings.Builder
	builder.Grow(176)

	builder.WriteString("CThostFtdcBrokerUserPasswordField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", UserID=%+v", d.UserID)
	fmt.Fprintf(&builder, ", Password=%+v", d.Password)
	fmt.Fprintf(&builder, ", LastUpdateTime=%+v", d.LastUpdateTime)
	fmt.Fprintf(&builder, ", LastLoginTime=%+v", d.LastLoginTime)
	fmt.Fprintf(&builder, ", ExpireDate=%+v", d.ExpireDate)
	fmt.Fprintf(&builder, ", WeakExpireDate=%+v", d.WeakExpireDate)

	builder.WriteByte('}')

	return builder.String()
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

func (d CThostFtdcBrokerUserFunctionField) Type() string { return "CThostFtdcBrokerUserFunctionField" }

func (d CThostFtdcBrokerUserFunctionField) String() string {
	var builder strings.Builder
	builder.Grow(95)

	builder.WriteString("CThostFtdcBrokerUserFunctionField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", UserID=%+v", d.UserID)
	fmt.Fprintf(&builder, ", BrokerFunctionCode=%+v", d.BrokerFunctionCode)

	builder.WriteByte('}')

	return builder.String()
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
	// 撤单时选择席位算法
	OrderCancelAlg types.TThostFtdcOrderCancelAlgType
}

func (d CThostFtdcTraderOfferField) Type() string { return "CThostFtdcTraderOfferField" }

func (d CThostFtdcTraderOfferField) String() string {
	var builder strings.Builder
	builder.Grow(475)

	builder.WriteString("CThostFtdcTraderOfferField{")
	fmt.Fprintf(&builder, "ExchangeID=%+v", d.ExchangeID)
	fmt.Fprintf(&builder, ", TraderID=%+v", d.TraderID)
	fmt.Fprintf(&builder, ", ParticipantID=%+v", d.ParticipantID)
	fmt.Fprintf(&builder, ", Password=%+v", d.Password)
	fmt.Fprintf(&builder, ", InstallID=%+v", d.InstallID)
	fmt.Fprintf(&builder, ", OrderLocalID=%+v", d.OrderLocalID)
	fmt.Fprintf(&builder, ", TraderConnectStatus=%+v", d.TraderConnectStatus)
	fmt.Fprintf(&builder, ", ConnectRequestDate=%+v", d.ConnectRequestDate)
	fmt.Fprintf(&builder, ", ConnectRequestTime=%+v", d.ConnectRequestTime)
	fmt.Fprintf(&builder, ", LastReportDate=%+v", d.LastReportDate)
	fmt.Fprintf(&builder, ", LastReportTime=%+v", d.LastReportTime)
	fmt.Fprintf(&builder, ", ConnectDate=%+v", d.ConnectDate)
	fmt.Fprintf(&builder, ", ConnectTime=%+v", d.ConnectTime)
	fmt.Fprintf(&builder, ", StartDate=%+v", d.StartDate)
	fmt.Fprintf(&builder, ", StartTime=%+v", d.StartTime)
	fmt.Fprintf(&builder, ", TradingDay=%+v", d.TradingDay)
	fmt.Fprintf(&builder, ", BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", MaxTradeID=%+v", d.MaxTradeID)
	fmt.Fprintf(&builder, ", MaxOrderMessageReference=%+v", d.MaxOrderMessageReference)
	fmt.Fprintf(&builder, ", OrderCancelAlg=%+v", d.OrderCancelAlg)

	builder.WriteByte('}')

	return builder.String()
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
	// 投资者帐号
	AccountID types.TThostFtdcAccountIDType
	// 币种代码
	CurrencyID types.TThostFtdcCurrencyIDType
}

func (d CThostFtdcSettlementInfoField) Type() string { return "CThostFtdcSettlementInfoField" }

func (d CThostFtdcSettlementInfoField) String() string {
	var builder strings.Builder
	builder.Grow(185)

	builder.WriteString("CThostFtdcSettlementInfoField{")
	fmt.Fprintf(&builder, "TradingDay=%+v", d.TradingDay)
	fmt.Fprintf(&builder, ", SettlementID=%+v", d.SettlementID)
	fmt.Fprintf(&builder, ", BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", InvestorID=%+v", d.InvestorID)
	fmt.Fprintf(&builder, ", SequenceNo=%+v", d.SequenceNo)
	fmt.Fprintf(&builder, ", Content=%+v", d.Content)
	fmt.Fprintf(&builder, ", AccountID=%+v", d.AccountID)
	fmt.Fprintf(&builder, ", CurrencyID=%+v", d.CurrencyID)

	builder.WriteByte('}')

	return builder.String()
}

// 合约保证金率调整
type CThostFtdcInstrumentMarginRateAdjustField struct {
	// 保留的无效字段
	reserve1 types.TThostFtdcOldInstrumentIDType
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
	// 合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
}

func (d CThostFtdcInstrumentMarginRateAdjustField) Type() string {
	return "CThostFtdcInstrumentMarginRateAdjustField"
}

func (d CThostFtdcInstrumentMarginRateAdjustField) String() string {
	var builder strings.Builder
	builder.Grow(313)

	builder.WriteString("CThostFtdcInstrumentMarginRateAdjustField{")
	fmt.Fprintf(&builder, "reserve1=%+v", d.reserve1)
	fmt.Fprintf(&builder, ", InvestorRange=%+v", d.InvestorRange)
	fmt.Fprintf(&builder, ", BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", InvestorID=%+v", d.InvestorID)
	fmt.Fprintf(&builder, ", HedgeFlag=%+v", d.HedgeFlag)
	fmt.Fprintf(&builder, ", LongMarginRatioByMoney=%+v", d.LongMarginRatioByMoney)
	fmt.Fprintf(&builder, ", LongMarginRatioByVolume=%+v", d.LongMarginRatioByVolume)
	fmt.Fprintf(&builder, ", ShortMarginRatioByMoney=%+v", d.ShortMarginRatioByMoney)
	fmt.Fprintf(&builder, ", ShortMarginRatioByVolume=%+v", d.ShortMarginRatioByVolume)
	fmt.Fprintf(&builder, ", IsRelative=%+v", d.IsRelative)
	fmt.Fprintf(&builder, ", InstrumentID=%+v", d.InstrumentID)

	builder.WriteByte('}')

	return builder.String()
}

// 交易所保证金率
type CThostFtdcExchangeMarginRateField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 保留的无效字段
	reserve1 types.TThostFtdcOldInstrumentIDType
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
	// 交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	// 合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
}

func (d CThostFtdcExchangeMarginRateField) Type() string { return "CThostFtdcExchangeMarginRateField" }

func (d CThostFtdcExchangeMarginRateField) String() string {
	var builder strings.Builder
	builder.Grow(262)

	builder.WriteString("CThostFtdcExchangeMarginRateField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", reserve1=%+v", d.reserve1)
	fmt.Fprintf(&builder, ", HedgeFlag=%+v", d.HedgeFlag)
	fmt.Fprintf(&builder, ", LongMarginRatioByMoney=%+v", d.LongMarginRatioByMoney)
	fmt.Fprintf(&builder, ", LongMarginRatioByVolume=%+v", d.LongMarginRatioByVolume)
	fmt.Fprintf(&builder, ", ShortMarginRatioByMoney=%+v", d.ShortMarginRatioByMoney)
	fmt.Fprintf(&builder, ", ShortMarginRatioByVolume=%+v", d.ShortMarginRatioByVolume)
	fmt.Fprintf(&builder, ", ExchangeID=%+v", d.ExchangeID)
	fmt.Fprintf(&builder, ", InstrumentID=%+v", d.InstrumentID)

	builder.WriteByte('}')

	return builder.String()
}

// 交易所保证金率调整
type CThostFtdcExchangeMarginRateAdjustField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 保留的无效字段
	reserve1 types.TThostFtdcOldInstrumentIDType
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
	// 合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
}

func (d CThostFtdcExchangeMarginRateAdjustField) Type() string {
	return "CThostFtdcExchangeMarginRateAdjustField"
}

func (d CThostFtdcExchangeMarginRateAdjustField) String() string {
	var builder strings.Builder
	builder.Grow(536)

	builder.WriteString("CThostFtdcExchangeMarginRateAdjustField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", reserve1=%+v", d.reserve1)
	fmt.Fprintf(&builder, ", HedgeFlag=%+v", d.HedgeFlag)
	fmt.Fprintf(&builder, ", LongMarginRatioByMoney=%+v", d.LongMarginRatioByMoney)
	fmt.Fprintf(&builder, ", LongMarginRatioByVolume=%+v", d.LongMarginRatioByVolume)
	fmt.Fprintf(&builder, ", ShortMarginRatioByMoney=%+v", d.ShortMarginRatioByMoney)
	fmt.Fprintf(&builder, ", ShortMarginRatioByVolume=%+v", d.ShortMarginRatioByVolume)
	fmt.Fprintf(&builder, ", ExchLongMarginRatioByMoney=%+v", d.ExchLongMarginRatioByMoney)
	fmt.Fprintf(&builder, ", ExchLongMarginRatioByVolume=%+v", d.ExchLongMarginRatioByVolume)
	fmt.Fprintf(&builder, ", ExchShortMarginRatioByMoney=%+v", d.ExchShortMarginRatioByMoney)
	fmt.Fprintf(&builder, ", ExchShortMarginRatioByVolume=%+v", d.ExchShortMarginRatioByVolume)
	fmt.Fprintf(&builder, ", NoLongMarginRatioByMoney=%+v", d.NoLongMarginRatioByMoney)
	fmt.Fprintf(&builder, ", NoLongMarginRatioByVolume=%+v", d.NoLongMarginRatioByVolume)
	fmt.Fprintf(&builder, ", NoShortMarginRatioByMoney=%+v", d.NoShortMarginRatioByMoney)
	fmt.Fprintf(&builder, ", NoShortMarginRatioByVolume=%+v", d.NoShortMarginRatioByVolume)
	fmt.Fprintf(&builder, ", InstrumentID=%+v", d.InstrumentID)

	builder.WriteByte('}')

	return builder.String()
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

func (d CThostFtdcExchangeRateField) Type() string { return "CThostFtdcExchangeRateField" }

func (d CThostFtdcExchangeRateField) String() string {
	var builder strings.Builder
	builder.Grow(139)

	builder.WriteString("CThostFtdcExchangeRateField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", FromCurrencyID=%+v", d.FromCurrencyID)
	fmt.Fprintf(&builder, ", FromCurrencyUnit=%+v", d.FromCurrencyUnit)
	fmt.Fprintf(&builder, ", ToCurrencyID=%+v", d.ToCurrencyID)
	fmt.Fprintf(&builder, ", ExchangeRate=%+v", d.ExchangeRate)

	builder.WriteByte('}')

	return builder.String()
}

// 结算引用
type CThostFtdcSettlementRefField struct {
	// 交易日
	TradingDay types.TThostFtdcDateType
	// 结算编号
	SettlementID types.TThostFtdcSettlementIDType
}

func (d CThostFtdcSettlementRefField) Type() string { return "CThostFtdcSettlementRefField" }

func (d CThostFtdcSettlementRefField) String() string {
	var builder strings.Builder
	builder.Grow(70)

	builder.WriteString("CThostFtdcSettlementRefField{")
	fmt.Fprintf(&builder, "TradingDay=%+v", d.TradingDay)
	fmt.Fprintf(&builder, ", SettlementID=%+v", d.SettlementID)

	builder.WriteByte('}')

	return builder.String()
}

// 当前时间
type CThostFtdcCurrentTimeField struct {
	// 当前交易日
	CurrDate types.TThostFtdcDateType
	// 当前时间
	CurrTime types.TThostFtdcTimeType
	// 当前时间（毫秒）
	CurrMillisec types.TThostFtdcMillisecType
	// 自然日期
	ActionDay types.TThostFtdcDateType
}

func (d CThostFtdcCurrentTimeField) Type() string { return "CThostFtdcCurrentTimeField" }

func (d CThostFtdcCurrentTimeField) String() string {
	var builder strings.Builder
	builder.Grow(103)

	builder.WriteString("CThostFtdcCurrentTimeField{")
	fmt.Fprintf(&builder, "CurrDate=%+v", d.CurrDate)
	fmt.Fprintf(&builder, ", CurrTime=%+v", d.CurrTime)
	fmt.Fprintf(&builder, ", CurrMillisec=%+v", d.CurrMillisec)
	fmt.Fprintf(&builder, ", ActionDay=%+v", d.ActionDay)

	builder.WriteByte('}')

	return builder.String()
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

func (d CThostFtdcCommPhaseField) Type() string { return "CThostFtdcCommPhaseField" }

func (d CThostFtdcCommPhaseField) String() string {
	var builder strings.Builder
	builder.Grow(83)

	builder.WriteString("CThostFtdcCommPhaseField{")
	fmt.Fprintf(&builder, "TradingDay=%+v", d.TradingDay)
	fmt.Fprintf(&builder, ", CommPhaseNo=%+v", d.CommPhaseNo)
	fmt.Fprintf(&builder, ", SystemID=%+v", d.SystemID)

	builder.WriteByte('}')

	return builder.String()
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
	// 保留的无效字段
	reserve1 types.TThostFtdcOldIPAddressType
	// 用户端产品信息
	UserProductInfo types.TThostFtdcProductInfoType
	// 接口端产品信息
	InterfaceProductInfo types.TThostFtdcProductInfoType
	// 协议信息
	ProtocolInfo types.TThostFtdcProtocolInfoType
	// 系统名称
	SystemName types.TThostFtdcSystemNameType
	// 密码,已弃用
	PasswordDeprecated types.TThostFtdcPasswordType
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
	// 密码
	Password types.TThostFtdcPasswordType
	// IP地址
	IPAddress types.TThostFtdcIPAddressType
}

func (d CThostFtdcLoginInfoField) Type() string { return "CThostFtdcLoginInfoField" }

func (d CThostFtdcLoginInfoField) String() string {
	var builder strings.Builder
	builder.Grow(509)

	builder.WriteString("CThostFtdcLoginInfoField{")
	fmt.Fprintf(&builder, "FrontID=%+v", d.FrontID)
	fmt.Fprintf(&builder, ", SessionID=%+v", d.SessionID)
	fmt.Fprintf(&builder, ", BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", UserID=%+v", d.UserID)
	fmt.Fprintf(&builder, ", LoginDate=%+v", d.LoginDate)
	fmt.Fprintf(&builder, ", LoginTime=%+v", d.LoginTime)
	fmt.Fprintf(&builder, ", reserve1=%+v", d.reserve1)
	fmt.Fprintf(&builder, ", UserProductInfo=%+v", d.UserProductInfo)
	fmt.Fprintf(&builder, ", InterfaceProductInfo=%+v", d.InterfaceProductInfo)
	fmt.Fprintf(&builder, ", ProtocolInfo=%+v", d.ProtocolInfo)
	fmt.Fprintf(&builder, ", SystemName=%+v", d.SystemName)
	fmt.Fprintf(&builder, ", PasswordDeprecated=%+v", d.PasswordDeprecated)
	fmt.Fprintf(&builder, ", MaxOrderRef=%+v", d.MaxOrderRef)
	fmt.Fprintf(&builder, ", SHFETime=%+v", d.SHFETime)
	fmt.Fprintf(&builder, ", DCETime=%+v", d.DCETime)
	fmt.Fprintf(&builder, ", CZCETime=%+v", d.CZCETime)
	fmt.Fprintf(&builder, ", FFEXTime=%+v", d.FFEXTime)
	fmt.Fprintf(&builder, ", MacAddress=%+v", d.MacAddress)
	fmt.Fprintf(&builder, ", OneTimePassword=%+v", d.OneTimePassword)
	fmt.Fprintf(&builder, ", INETime=%+v", d.INETime)
	fmt.Fprintf(&builder, ", IsQryControl=%+v", d.IsQryControl)
	fmt.Fprintf(&builder, ", LoginRemark=%+v", d.LoginRemark)
	fmt.Fprintf(&builder, ", Password=%+v", d.Password)
	fmt.Fprintf(&builder, ", IPAddress=%+v", d.IPAddress)

	builder.WriteByte('}')

	return builder.String()
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

func (d CThostFtdcLogoutAllField) Type() string { return "CThostFtdcLogoutAllField" }

func (d CThostFtdcLogoutAllField) String() string {
	var builder strings.Builder
	builder.Grow(80)

	builder.WriteString("CThostFtdcLogoutAllField{")
	fmt.Fprintf(&builder, "FrontID=%+v", d.FrontID)
	fmt.Fprintf(&builder, ", SessionID=%+v", d.SessionID)
	fmt.Fprintf(&builder, ", SystemName=%+v", d.SystemName)

	builder.WriteByte('}')

	return builder.String()
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

func (d CThostFtdcFrontStatusField) Type() string { return "CThostFtdcFrontStatusField" }

func (d CThostFtdcFrontStatusField) String() string {
	var builder strings.Builder
	builder.Grow(109)

	builder.WriteString("CThostFtdcFrontStatusField{")
	fmt.Fprintf(&builder, "FrontID=%+v", d.FrontID)
	fmt.Fprintf(&builder, ", LastReportDate=%+v", d.LastReportDate)
	fmt.Fprintf(&builder, ", LastReportTime=%+v", d.LastReportTime)
	fmt.Fprintf(&builder, ", IsActive=%+v", d.IsActive)

	builder.WriteByte('}')

	return builder.String()
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

func (d CThostFtdcUserPasswordUpdateField) Type() string { return "CThostFtdcUserPasswordUpdateField" }

func (d CThostFtdcUserPasswordUpdateField) String() string {
	var builder strings.Builder
	builder.Grow(109)

	builder.WriteString("CThostFtdcUserPasswordUpdateField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", UserID=%+v", d.UserID)
	fmt.Fprintf(&builder, ", OldPassword=%+v", d.OldPassword)
	fmt.Fprintf(&builder, ", NewPassword=%+v", d.NewPassword)

	builder.WriteByte('}')

	return builder.String()
}

// 输入报单
type CThostFtdcInputOrderField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	// 保留的无效字段
	reserve1 types.TThostFtdcOldInstrumentIDType
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
	// 用户强平标志
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
	// 保留的无效字段
	reserve2 types.TThostFtdcOldIPAddressType
	// Mac地址
	MacAddress types.TThostFtdcMacAddressType
	// 合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
	// IP地址
	IPAddress types.TThostFtdcIPAddressType
	// 报单回显字段
	OrderMemo types.TThostFtdcOrderMemoType
	// session上请求计数 api自动维护
	SessionReqSeq types.TThostFtdcSequenceNo12Type
}

func (d CThostFtdcInputOrderField) Type() string { return "CThostFtdcInputOrderField" }

func (d CThostFtdcInputOrderField) String() string {
	var builder strings.Builder
	builder.Grow(741)

	builder.WriteString("CThostFtdcInputOrderField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", InvestorID=%+v", d.InvestorID)
	fmt.Fprintf(&builder, ", reserve1=%+v", d.reserve1)
	fmt.Fprintf(&builder, ", OrderRef=%+v", d.OrderRef)
	fmt.Fprintf(&builder, ", UserID=%+v", d.UserID)
	fmt.Fprintf(&builder, ", OrderPriceType=%+v", d.OrderPriceType)
	fmt.Fprintf(&builder, ", Direction=%+v", d.Direction)
	fmt.Fprintf(&builder, ", CombOffsetFlag=%+v", d.CombOffsetFlag)
	fmt.Fprintf(&builder, ", CombHedgeFlag=%+v", d.CombHedgeFlag)
	fmt.Fprintf(&builder, ", LimitPrice=%+v", d.LimitPrice)
	fmt.Fprintf(&builder, ", VolumeTotalOriginal=%+v", d.VolumeTotalOriginal)
	fmt.Fprintf(&builder, ", TimeCondition=%+v", d.TimeCondition)
	fmt.Fprintf(&builder, ", GTDDate=%+v", d.GTDDate)
	fmt.Fprintf(&builder, ", VolumeCondition=%+v", d.VolumeCondition)
	fmt.Fprintf(&builder, ", MinVolume=%+v", d.MinVolume)
	fmt.Fprintf(&builder, ", ContingentCondition=%+v", d.ContingentCondition)
	fmt.Fprintf(&builder, ", StopPrice=%+v", d.StopPrice)
	fmt.Fprintf(&builder, ", ForceCloseReason=%+v", d.ForceCloseReason)
	fmt.Fprintf(&builder, ", IsAutoSuspend=%+v", d.IsAutoSuspend)
	fmt.Fprintf(&builder, ", BusinessUnit=%+v", d.BusinessUnit)
	fmt.Fprintf(&builder, ", RequestID=%+v", d.RequestID)
	fmt.Fprintf(&builder, ", UserForceClose=%+v", d.UserForceClose)
	fmt.Fprintf(&builder, ", IsSwapOrder=%+v", d.IsSwapOrder)
	fmt.Fprintf(&builder, ", ExchangeID=%+v", d.ExchangeID)
	fmt.Fprintf(&builder, ", InvestUnitID=%+v", d.InvestUnitID)
	fmt.Fprintf(&builder, ", AccountID=%+v", d.AccountID)
	fmt.Fprintf(&builder, ", CurrencyID=%+v", d.CurrencyID)
	fmt.Fprintf(&builder, ", ClientID=%+v", d.ClientID)
	fmt.Fprintf(&builder, ", reserve2=%+v", d.reserve2)
	fmt.Fprintf(&builder, ", MacAddress=%+v", d.MacAddress)
	fmt.Fprintf(&builder, ", InstrumentID=%+v", d.InstrumentID)
	fmt.Fprintf(&builder, ", IPAddress=%+v", d.IPAddress)
	fmt.Fprintf(&builder, ", OrderMemo=%+v", d.OrderMemo)
	fmt.Fprintf(&builder, ", SessionReqSeq=%+v", d.SessionReqSeq)

	builder.WriteByte('}')

	return builder.String()
}

// 报单
type CThostFtdcOrderField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	// 保留的无效字段
	reserve1 types.TThostFtdcOldInstrumentIDType
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
	// 保留的无效字段
	reserve2 types.TThostFtdcOldExchangeInstIDType
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
	// 用户强平标志
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
	// 保留的无效字段
	reserve3 types.TThostFtdcOldIPAddressType
	// Mac地址
	MacAddress types.TThostFtdcMacAddressType
	// 合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
	// 合约在交易所的代码
	ExchangeInstID types.TThostFtdcExchangeInstIDType
	// IP地址
	IPAddress types.TThostFtdcIPAddressType
	// 报单回显字段
	OrderMemo types.TThostFtdcOrderMemoType
	// session上请求计数 api自动维护
	SessionReqSeq types.TThostFtdcSequenceNo12Type
}

func (d CThostFtdcOrderField) Type() string { return "CThostFtdcOrderField" }

func (d CThostFtdcOrderField) String() string {
	var builder strings.Builder
	builder.Grow(1468)

	builder.WriteString("CThostFtdcOrderField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", InvestorID=%+v", d.InvestorID)
	fmt.Fprintf(&builder, ", reserve1=%+v", d.reserve1)
	fmt.Fprintf(&builder, ", OrderRef=%+v", d.OrderRef)
	fmt.Fprintf(&builder, ", UserID=%+v", d.UserID)
	fmt.Fprintf(&builder, ", OrderPriceType=%+v", d.OrderPriceType)
	fmt.Fprintf(&builder, ", Direction=%+v", d.Direction)
	fmt.Fprintf(&builder, ", CombOffsetFlag=%+v", d.CombOffsetFlag)
	fmt.Fprintf(&builder, ", CombHedgeFlag=%+v", d.CombHedgeFlag)
	fmt.Fprintf(&builder, ", LimitPrice=%+v", d.LimitPrice)
	fmt.Fprintf(&builder, ", VolumeTotalOriginal=%+v", d.VolumeTotalOriginal)
	fmt.Fprintf(&builder, ", TimeCondition=%+v", d.TimeCondition)
	fmt.Fprintf(&builder, ", GTDDate=%+v", d.GTDDate)
	fmt.Fprintf(&builder, ", VolumeCondition=%+v", d.VolumeCondition)
	fmt.Fprintf(&builder, ", MinVolume=%+v", d.MinVolume)
	fmt.Fprintf(&builder, ", ContingentCondition=%+v", d.ContingentCondition)
	fmt.Fprintf(&builder, ", StopPrice=%+v", d.StopPrice)
	fmt.Fprintf(&builder, ", ForceCloseReason=%+v", d.ForceCloseReason)
	fmt.Fprintf(&builder, ", IsAutoSuspend=%+v", d.IsAutoSuspend)
	fmt.Fprintf(&builder, ", BusinessUnit=%+v", d.BusinessUnit)
	fmt.Fprintf(&builder, ", RequestID=%+v", d.RequestID)
	fmt.Fprintf(&builder, ", OrderLocalID=%+v", d.OrderLocalID)
	fmt.Fprintf(&builder, ", ExchangeID=%+v", d.ExchangeID)
	fmt.Fprintf(&builder, ", ParticipantID=%+v", d.ParticipantID)
	fmt.Fprintf(&builder, ", ClientID=%+v", d.ClientID)
	fmt.Fprintf(&builder, ", reserve2=%+v", d.reserve2)
	fmt.Fprintf(&builder, ", TraderID=%+v", d.TraderID)
	fmt.Fprintf(&builder, ", InstallID=%+v", d.InstallID)
	fmt.Fprintf(&builder, ", OrderSubmitStatus=%+v", d.OrderSubmitStatus)
	fmt.Fprintf(&builder, ", NotifySequence=%+v", d.NotifySequence)
	fmt.Fprintf(&builder, ", TradingDay=%+v", d.TradingDay)
	fmt.Fprintf(&builder, ", SettlementID=%+v", d.SettlementID)
	fmt.Fprintf(&builder, ", OrderSysID=%+v", d.OrderSysID)
	fmt.Fprintf(&builder, ", OrderSource=%+v", d.OrderSource)
	fmt.Fprintf(&builder, ", OrderStatus=%+v", d.OrderStatus)
	fmt.Fprintf(&builder, ", OrderType=%+v", d.OrderType)
	fmt.Fprintf(&builder, ", VolumeTraded=%+v", d.VolumeTraded)
	fmt.Fprintf(&builder, ", VolumeTotal=%+v", d.VolumeTotal)
	fmt.Fprintf(&builder, ", InsertDate=%+v", d.InsertDate)
	fmt.Fprintf(&builder, ", InsertTime=%+v", d.InsertTime)
	fmt.Fprintf(&builder, ", ActiveTime=%+v", d.ActiveTime)
	fmt.Fprintf(&builder, ", SuspendTime=%+v", d.SuspendTime)
	fmt.Fprintf(&builder, ", UpdateTime=%+v", d.UpdateTime)
	fmt.Fprintf(&builder, ", CancelTime=%+v", d.CancelTime)
	fmt.Fprintf(&builder, ", ActiveTraderID=%+v", d.ActiveTraderID)
	fmt.Fprintf(&builder, ", ClearingPartID=%+v", d.ClearingPartID)
	fmt.Fprintf(&builder, ", SequenceNo=%+v", d.SequenceNo)
	fmt.Fprintf(&builder, ", FrontID=%+v", d.FrontID)
	fmt.Fprintf(&builder, ", SessionID=%+v", d.SessionID)
	fmt.Fprintf(&builder, ", UserProductInfo=%+v", d.UserProductInfo)
	fmt.Fprintf(&builder, ", StatusMsg=%+v", d.StatusMsg)
	fmt.Fprintf(&builder, ", UserForceClose=%+v", d.UserForceClose)
	fmt.Fprintf(&builder, ", ActiveUserID=%+v", d.ActiveUserID)
	fmt.Fprintf(&builder, ", BrokerOrderSeq=%+v", d.BrokerOrderSeq)
	fmt.Fprintf(&builder, ", RelativeOrderSysID=%+v", d.RelativeOrderSysID)
	fmt.Fprintf(&builder, ", ZCETotalTradedVolume=%+v", d.ZCETotalTradedVolume)
	fmt.Fprintf(&builder, ", IsSwapOrder=%+v", d.IsSwapOrder)
	fmt.Fprintf(&builder, ", BranchID=%+v", d.BranchID)
	fmt.Fprintf(&builder, ", InvestUnitID=%+v", d.InvestUnitID)
	fmt.Fprintf(&builder, ", AccountID=%+v", d.AccountID)
	fmt.Fprintf(&builder, ", CurrencyID=%+v", d.CurrencyID)
	fmt.Fprintf(&builder, ", reserve3=%+v", d.reserve3)
	fmt.Fprintf(&builder, ", MacAddress=%+v", d.MacAddress)
	fmt.Fprintf(&builder, ", InstrumentID=%+v", d.InstrumentID)
	fmt.Fprintf(&builder, ", ExchangeInstID=%+v", d.ExchangeInstID)
	fmt.Fprintf(&builder, ", IPAddress=%+v", d.IPAddress)
	fmt.Fprintf(&builder, ", OrderMemo=%+v", d.OrderMemo)
	fmt.Fprintf(&builder, ", SessionReqSeq=%+v", d.SessionReqSeq)

	builder.WriteByte('}')

	return builder.String()
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
	// 保留的无效字段
	reserve1 types.TThostFtdcOldExchangeInstIDType
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
	// 保留的无效字段
	reserve2 types.TThostFtdcOldIPAddressType
	// Mac地址
	MacAddress types.TThostFtdcMacAddressType
	// 合约在交易所的代码
	ExchangeInstID types.TThostFtdcExchangeInstIDType
	// IP地址
	IPAddress types.TThostFtdcIPAddressType
}

func (d CThostFtdcExchangeOrderField) Type() string { return "CThostFtdcExchangeOrderField" }

func (d CThostFtdcExchangeOrderField) String() string {
	var builder strings.Builder
	builder.Grow(1032)

	builder.WriteString("CThostFtdcExchangeOrderField{")
	fmt.Fprintf(&builder, "OrderPriceType=%+v", d.OrderPriceType)
	fmt.Fprintf(&builder, ", Direction=%+v", d.Direction)
	fmt.Fprintf(&builder, ", CombOffsetFlag=%+v", d.CombOffsetFlag)
	fmt.Fprintf(&builder, ", CombHedgeFlag=%+v", d.CombHedgeFlag)
	fmt.Fprintf(&builder, ", LimitPrice=%+v", d.LimitPrice)
	fmt.Fprintf(&builder, ", VolumeTotalOriginal=%+v", d.VolumeTotalOriginal)
	fmt.Fprintf(&builder, ", TimeCondition=%+v", d.TimeCondition)
	fmt.Fprintf(&builder, ", GTDDate=%+v", d.GTDDate)
	fmt.Fprintf(&builder, ", VolumeCondition=%+v", d.VolumeCondition)
	fmt.Fprintf(&builder, ", MinVolume=%+v", d.MinVolume)
	fmt.Fprintf(&builder, ", ContingentCondition=%+v", d.ContingentCondition)
	fmt.Fprintf(&builder, ", StopPrice=%+v", d.StopPrice)
	fmt.Fprintf(&builder, ", ForceCloseReason=%+v", d.ForceCloseReason)
	fmt.Fprintf(&builder, ", IsAutoSuspend=%+v", d.IsAutoSuspend)
	fmt.Fprintf(&builder, ", BusinessUnit=%+v", d.BusinessUnit)
	fmt.Fprintf(&builder, ", RequestID=%+v", d.RequestID)
	fmt.Fprintf(&builder, ", OrderLocalID=%+v", d.OrderLocalID)
	fmt.Fprintf(&builder, ", ExchangeID=%+v", d.ExchangeID)
	fmt.Fprintf(&builder, ", ParticipantID=%+v", d.ParticipantID)
	fmt.Fprintf(&builder, ", ClientID=%+v", d.ClientID)
	fmt.Fprintf(&builder, ", reserve1=%+v", d.reserve1)
	fmt.Fprintf(&builder, ", TraderID=%+v", d.TraderID)
	fmt.Fprintf(&builder, ", InstallID=%+v", d.InstallID)
	fmt.Fprintf(&builder, ", OrderSubmitStatus=%+v", d.OrderSubmitStatus)
	fmt.Fprintf(&builder, ", NotifySequence=%+v", d.NotifySequence)
	fmt.Fprintf(&builder, ", TradingDay=%+v", d.TradingDay)
	fmt.Fprintf(&builder, ", SettlementID=%+v", d.SettlementID)
	fmt.Fprintf(&builder, ", OrderSysID=%+v", d.OrderSysID)
	fmt.Fprintf(&builder, ", OrderSource=%+v", d.OrderSource)
	fmt.Fprintf(&builder, ", OrderStatus=%+v", d.OrderStatus)
	fmt.Fprintf(&builder, ", OrderType=%+v", d.OrderType)
	fmt.Fprintf(&builder, ", VolumeTraded=%+v", d.VolumeTraded)
	fmt.Fprintf(&builder, ", VolumeTotal=%+v", d.VolumeTotal)
	fmt.Fprintf(&builder, ", InsertDate=%+v", d.InsertDate)
	fmt.Fprintf(&builder, ", InsertTime=%+v", d.InsertTime)
	fmt.Fprintf(&builder, ", ActiveTime=%+v", d.ActiveTime)
	fmt.Fprintf(&builder, ", SuspendTime=%+v", d.SuspendTime)
	fmt.Fprintf(&builder, ", UpdateTime=%+v", d.UpdateTime)
	fmt.Fprintf(&builder, ", CancelTime=%+v", d.CancelTime)
	fmt.Fprintf(&builder, ", ActiveTraderID=%+v", d.ActiveTraderID)
	fmt.Fprintf(&builder, ", ClearingPartID=%+v", d.ClearingPartID)
	fmt.Fprintf(&builder, ", SequenceNo=%+v", d.SequenceNo)
	fmt.Fprintf(&builder, ", BranchID=%+v", d.BranchID)
	fmt.Fprintf(&builder, ", reserve2=%+v", d.reserve2)
	fmt.Fprintf(&builder, ", MacAddress=%+v", d.MacAddress)
	fmt.Fprintf(&builder, ", ExchangeInstID=%+v", d.ExchangeInstID)
	fmt.Fprintf(&builder, ", IPAddress=%+v", d.IPAddress)

	builder.WriteByte('}')

	return builder.String()
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

func (d CThostFtdcExchangeOrderInsertErrorField) Type() string {
	return "CThostFtdcExchangeOrderInsertErrorField"
}

func (d CThostFtdcExchangeOrderInsertErrorField) String() string {
	var builder strings.Builder
	builder.Grow(176)

	builder.WriteString("CThostFtdcExchangeOrderInsertErrorField{")
	fmt.Fprintf(&builder, "ExchangeID=%+v", d.ExchangeID)
	fmt.Fprintf(&builder, ", ParticipantID=%+v", d.ParticipantID)
	fmt.Fprintf(&builder, ", TraderID=%+v", d.TraderID)
	fmt.Fprintf(&builder, ", InstallID=%+v", d.InstallID)
	fmt.Fprintf(&builder, ", OrderLocalID=%+v", d.OrderLocalID)
	fmt.Fprintf(&builder, ", ErrorID=%+v", d.ErrorID)
	fmt.Fprintf(&builder, ", ErrorMsg=%+v", d.ErrorMsg)

	builder.WriteByte('}')

	return builder.String()
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
	// 保留的无效字段
	reserve1 types.TThostFtdcOldInstrumentIDType
	// 投资单元代码
	InvestUnitID types.TThostFtdcInvestUnitIDType
	// 保留的无效字段
	reserve2 types.TThostFtdcOldIPAddressType
	// Mac地址
	MacAddress types.TThostFtdcMacAddressType
	// 合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
	// IP地址
	IPAddress types.TThostFtdcIPAddressType
	// 报单回显字段
	OrderMemo types.TThostFtdcOrderMemoType
	// session上请求计数 api自动维护
	SessionReqSeq types.TThostFtdcSequenceNo12Type
}

func (d CThostFtdcInputOrderActionField) Type() string { return "CThostFtdcInputOrderActionField" }

func (d CThostFtdcInputOrderActionField) String() string {
	var builder strings.Builder
	builder.Grow(445)

	builder.WriteString("CThostFtdcInputOrderActionField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", InvestorID=%+v", d.InvestorID)
	fmt.Fprintf(&builder, ", OrderActionRef=%+v", d.OrderActionRef)
	fmt.Fprintf(&builder, ", OrderRef=%+v", d.OrderRef)
	fmt.Fprintf(&builder, ", RequestID=%+v", d.RequestID)
	fmt.Fprintf(&builder, ", FrontID=%+v", d.FrontID)
	fmt.Fprintf(&builder, ", SessionID=%+v", d.SessionID)
	fmt.Fprintf(&builder, ", ExchangeID=%+v", d.ExchangeID)
	fmt.Fprintf(&builder, ", OrderSysID=%+v", d.OrderSysID)
	fmt.Fprintf(&builder, ", ActionFlag=%+v", d.ActionFlag)
	fmt.Fprintf(&builder, ", LimitPrice=%+v", d.LimitPrice)
	fmt.Fprintf(&builder, ", VolumeChange=%+v", d.VolumeChange)
	fmt.Fprintf(&builder, ", UserID=%+v", d.UserID)
	fmt.Fprintf(&builder, ", reserve1=%+v", d.reserve1)
	fmt.Fprintf(&builder, ", InvestUnitID=%+v", d.InvestUnitID)
	fmt.Fprintf(&builder, ", reserve2=%+v", d.reserve2)
	fmt.Fprintf(&builder, ", MacAddress=%+v", d.MacAddress)
	fmt.Fprintf(&builder, ", InstrumentID=%+v", d.InstrumentID)
	fmt.Fprintf(&builder, ", IPAddress=%+v", d.IPAddress)
	fmt.Fprintf(&builder, ", OrderMemo=%+v", d.OrderMemo)
	fmt.Fprintf(&builder, ", SessionReqSeq=%+v", d.SessionReqSeq)

	builder.WriteByte('}')

	return builder.String()
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
	// 保留的无效字段
	reserve1 types.TThostFtdcOldInstrumentIDType
	// 营业部编号
	BranchID types.TThostFtdcBranchIDType
	// 投资单元代码
	InvestUnitID types.TThostFtdcInvestUnitIDType
	// 保留的无效字段
	reserve2 types.TThostFtdcOldIPAddressType
	// Mac地址
	MacAddress types.TThostFtdcMacAddressType
	// 合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
	// IP地址
	IPAddress types.TThostFtdcIPAddressType
	// 报单回显字段
	OrderMemo types.TThostFtdcOrderMemoType
	// session上请求计数 api自动维护
	SessionReqSeq types.TThostFtdcSequenceNo12Type
}

func (d CThostFtdcOrderActionField) Type() string { return "CThostFtdcOrderActionField" }

func (d CThostFtdcOrderActionField) String() string {
	var builder strings.Builder
	builder.Grow(689)

	builder.WriteString("CThostFtdcOrderActionField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", InvestorID=%+v", d.InvestorID)
	fmt.Fprintf(&builder, ", OrderActionRef=%+v", d.OrderActionRef)
	fmt.Fprintf(&builder, ", OrderRef=%+v", d.OrderRef)
	fmt.Fprintf(&builder, ", RequestID=%+v", d.RequestID)
	fmt.Fprintf(&builder, ", FrontID=%+v", d.FrontID)
	fmt.Fprintf(&builder, ", SessionID=%+v", d.SessionID)
	fmt.Fprintf(&builder, ", ExchangeID=%+v", d.ExchangeID)
	fmt.Fprintf(&builder, ", OrderSysID=%+v", d.OrderSysID)
	fmt.Fprintf(&builder, ", ActionFlag=%+v", d.ActionFlag)
	fmt.Fprintf(&builder, ", LimitPrice=%+v", d.LimitPrice)
	fmt.Fprintf(&builder, ", VolumeChange=%+v", d.VolumeChange)
	fmt.Fprintf(&builder, ", ActionDate=%+v", d.ActionDate)
	fmt.Fprintf(&builder, ", ActionTime=%+v", d.ActionTime)
	fmt.Fprintf(&builder, ", TraderID=%+v", d.TraderID)
	fmt.Fprintf(&builder, ", InstallID=%+v", d.InstallID)
	fmt.Fprintf(&builder, ", OrderLocalID=%+v", d.OrderLocalID)
	fmt.Fprintf(&builder, ", ActionLocalID=%+v", d.ActionLocalID)
	fmt.Fprintf(&builder, ", ParticipantID=%+v", d.ParticipantID)
	fmt.Fprintf(&builder, ", ClientID=%+v", d.ClientID)
	fmt.Fprintf(&builder, ", BusinessUnit=%+v", d.BusinessUnit)
	fmt.Fprintf(&builder, ", OrderActionStatus=%+v", d.OrderActionStatus)
	fmt.Fprintf(&builder, ", UserID=%+v", d.UserID)
	fmt.Fprintf(&builder, ", StatusMsg=%+v", d.StatusMsg)
	fmt.Fprintf(&builder, ", reserve1=%+v", d.reserve1)
	fmt.Fprintf(&builder, ", BranchID=%+v", d.BranchID)
	fmt.Fprintf(&builder, ", InvestUnitID=%+v", d.InvestUnitID)
	fmt.Fprintf(&builder, ", reserve2=%+v", d.reserve2)
	fmt.Fprintf(&builder, ", MacAddress=%+v", d.MacAddress)
	fmt.Fprintf(&builder, ", InstrumentID=%+v", d.InstrumentID)
	fmt.Fprintf(&builder, ", IPAddress=%+v", d.IPAddress)
	fmt.Fprintf(&builder, ", OrderMemo=%+v", d.OrderMemo)
	fmt.Fprintf(&builder, ", SessionReqSeq=%+v", d.SessionReqSeq)

	builder.WriteByte('}')

	return builder.String()
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
	// 保留的无效字段
	reserve1 types.TThostFtdcOldIPAddressType
	// Mac地址
	MacAddress types.TThostFtdcMacAddressType
	// IP地址
	IPAddress types.TThostFtdcIPAddressType
}

func (d CThostFtdcExchangeOrderActionField) Type() string {
	return "CThostFtdcExchangeOrderActionField"
}

func (d CThostFtdcExchangeOrderActionField) String() string {
	var builder strings.Builder
	builder.Grow(439)

	builder.WriteString("CThostFtdcExchangeOrderActionField{")
	fmt.Fprintf(&builder, "ExchangeID=%+v", d.ExchangeID)
	fmt.Fprintf(&builder, ", OrderSysID=%+v", d.OrderSysID)
	fmt.Fprintf(&builder, ", ActionFlag=%+v", d.ActionFlag)
	fmt.Fprintf(&builder, ", LimitPrice=%+v", d.LimitPrice)
	fmt.Fprintf(&builder, ", VolumeChange=%+v", d.VolumeChange)
	fmt.Fprintf(&builder, ", ActionDate=%+v", d.ActionDate)
	fmt.Fprintf(&builder, ", ActionTime=%+v", d.ActionTime)
	fmt.Fprintf(&builder, ", TraderID=%+v", d.TraderID)
	fmt.Fprintf(&builder, ", InstallID=%+v", d.InstallID)
	fmt.Fprintf(&builder, ", OrderLocalID=%+v", d.OrderLocalID)
	fmt.Fprintf(&builder, ", ActionLocalID=%+v", d.ActionLocalID)
	fmt.Fprintf(&builder, ", ParticipantID=%+v", d.ParticipantID)
	fmt.Fprintf(&builder, ", ClientID=%+v", d.ClientID)
	fmt.Fprintf(&builder, ", BusinessUnit=%+v", d.BusinessUnit)
	fmt.Fprintf(&builder, ", OrderActionStatus=%+v", d.OrderActionStatus)
	fmt.Fprintf(&builder, ", UserID=%+v", d.UserID)
	fmt.Fprintf(&builder, ", BranchID=%+v", d.BranchID)
	fmt.Fprintf(&builder, ", reserve1=%+v", d.reserve1)
	fmt.Fprintf(&builder, ", MacAddress=%+v", d.MacAddress)
	fmt.Fprintf(&builder, ", IPAddress=%+v", d.IPAddress)

	builder.WriteByte('}')

	return builder.String()
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

func (d CThostFtdcExchangeOrderActionErrorField) Type() string {
	return "CThostFtdcExchangeOrderActionErrorField"
}

func (d CThostFtdcExchangeOrderActionErrorField) String() string {
	var builder strings.Builder
	builder.Grow(196)

	builder.WriteString("CThostFtdcExchangeOrderActionErrorField{")
	fmt.Fprintf(&builder, "ExchangeID=%+v", d.ExchangeID)
	fmt.Fprintf(&builder, ", OrderSysID=%+v", d.OrderSysID)
	fmt.Fprintf(&builder, ", TraderID=%+v", d.TraderID)
	fmt.Fprintf(&builder, ", InstallID=%+v", d.InstallID)
	fmt.Fprintf(&builder, ", OrderLocalID=%+v", d.OrderLocalID)
	fmt.Fprintf(&builder, ", ActionLocalID=%+v", d.ActionLocalID)
	fmt.Fprintf(&builder, ", ErrorID=%+v", d.ErrorID)
	fmt.Fprintf(&builder, ", ErrorMsg=%+v", d.ErrorMsg)

	builder.WriteByte('}')

	return builder.String()
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
	// 保留的无效字段
	reserve1 types.TThostFtdcOldExchangeInstIDType
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
	// 合约在交易所的代码
	ExchangeInstID types.TThostFtdcExchangeInstIDType
}

func (d CThostFtdcExchangeTradeField) Type() string { return "CThostFtdcExchangeTradeField" }

func (d CThostFtdcExchangeTradeField) String() string {
	var builder strings.Builder
	builder.Grow(483)

	builder.WriteString("CThostFtdcExchangeTradeField{")
	fmt.Fprintf(&builder, "ExchangeID=%+v", d.ExchangeID)
	fmt.Fprintf(&builder, ", TradeID=%+v", d.TradeID)
	fmt.Fprintf(&builder, ", Direction=%+v", d.Direction)
	fmt.Fprintf(&builder, ", OrderSysID=%+v", d.OrderSysID)
	fmt.Fprintf(&builder, ", ParticipantID=%+v", d.ParticipantID)
	fmt.Fprintf(&builder, ", ClientID=%+v", d.ClientID)
	fmt.Fprintf(&builder, ", TradingRole=%+v", d.TradingRole)
	fmt.Fprintf(&builder, ", reserve1=%+v", d.reserve1)
	fmt.Fprintf(&builder, ", OffsetFlag=%+v", d.OffsetFlag)
	fmt.Fprintf(&builder, ", HedgeFlag=%+v", d.HedgeFlag)
	fmt.Fprintf(&builder, ", Price=%+v", d.Price)
	fmt.Fprintf(&builder, ", Volume=%+v", d.Volume)
	fmt.Fprintf(&builder, ", TradeDate=%+v", d.TradeDate)
	fmt.Fprintf(&builder, ", TradeTime=%+v", d.TradeTime)
	fmt.Fprintf(&builder, ", TradeType=%+v", d.TradeType)
	fmt.Fprintf(&builder, ", PriceSource=%+v", d.PriceSource)
	fmt.Fprintf(&builder, ", TraderID=%+v", d.TraderID)
	fmt.Fprintf(&builder, ", OrderLocalID=%+v", d.OrderLocalID)
	fmt.Fprintf(&builder, ", ClearingPartID=%+v", d.ClearingPartID)
	fmt.Fprintf(&builder, ", BusinessUnit=%+v", d.BusinessUnit)
	fmt.Fprintf(&builder, ", SequenceNo=%+v", d.SequenceNo)
	fmt.Fprintf(&builder, ", TradeSource=%+v", d.TradeSource)
	fmt.Fprintf(&builder, ", ExchangeInstID=%+v", d.ExchangeInstID)

	builder.WriteByte('}')

	return builder.String()
}

// 成交
type CThostFtdcTradeField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	// 保留的无效字段
	reserve1 types.TThostFtdcOldInstrumentIDType
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
	// 保留的无效字段
	reserve2 types.TThostFtdcOldExchangeInstIDType
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
	// 投资单元代码
	InvestUnitID types.TThostFtdcInvestUnitIDType
	// 合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
	// 合约在交易所的代码
	ExchangeInstID types.TThostFtdcExchangeInstIDType
}

func (d CThostFtdcTradeField) Type() string { return "CThostFtdcTradeField" }

func (d CThostFtdcTradeField) String() string {
	var builder strings.Builder
	builder.Grow(675)

	builder.WriteString("CThostFtdcTradeField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", InvestorID=%+v", d.InvestorID)
	fmt.Fprintf(&builder, ", reserve1=%+v", d.reserve1)
	fmt.Fprintf(&builder, ", OrderRef=%+v", d.OrderRef)
	fmt.Fprintf(&builder, ", UserID=%+v", d.UserID)
	fmt.Fprintf(&builder, ", ExchangeID=%+v", d.ExchangeID)
	fmt.Fprintf(&builder, ", TradeID=%+v", d.TradeID)
	fmt.Fprintf(&builder, ", Direction=%+v", d.Direction)
	fmt.Fprintf(&builder, ", OrderSysID=%+v", d.OrderSysID)
	fmt.Fprintf(&builder, ", ParticipantID=%+v", d.ParticipantID)
	fmt.Fprintf(&builder, ", ClientID=%+v", d.ClientID)
	fmt.Fprintf(&builder, ", TradingRole=%+v", d.TradingRole)
	fmt.Fprintf(&builder, ", reserve2=%+v", d.reserve2)
	fmt.Fprintf(&builder, ", OffsetFlag=%+v", d.OffsetFlag)
	fmt.Fprintf(&builder, ", HedgeFlag=%+v", d.HedgeFlag)
	fmt.Fprintf(&builder, ", Price=%+v", d.Price)
	fmt.Fprintf(&builder, ", Volume=%+v", d.Volume)
	fmt.Fprintf(&builder, ", TradeDate=%+v", d.TradeDate)
	fmt.Fprintf(&builder, ", TradeTime=%+v", d.TradeTime)
	fmt.Fprintf(&builder, ", TradeType=%+v", d.TradeType)
	fmt.Fprintf(&builder, ", PriceSource=%+v", d.PriceSource)
	fmt.Fprintf(&builder, ", TraderID=%+v", d.TraderID)
	fmt.Fprintf(&builder, ", OrderLocalID=%+v", d.OrderLocalID)
	fmt.Fprintf(&builder, ", ClearingPartID=%+v", d.ClearingPartID)
	fmt.Fprintf(&builder, ", BusinessUnit=%+v", d.BusinessUnit)
	fmt.Fprintf(&builder, ", SequenceNo=%+v", d.SequenceNo)
	fmt.Fprintf(&builder, ", TradingDay=%+v", d.TradingDay)
	fmt.Fprintf(&builder, ", SettlementID=%+v", d.SettlementID)
	fmt.Fprintf(&builder, ", BrokerOrderSeq=%+v", d.BrokerOrderSeq)
	fmt.Fprintf(&builder, ", TradeSource=%+v", d.TradeSource)
	fmt.Fprintf(&builder, ", InvestUnitID=%+v", d.InvestUnitID)
	fmt.Fprintf(&builder, ", InstrumentID=%+v", d.InstrumentID)
	fmt.Fprintf(&builder, ", ExchangeInstID=%+v", d.ExchangeInstID)

	builder.WriteByte('}')

	return builder.String()
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
	// 保留的无效字段
	reserve1 types.TThostFtdcOldIPAddressType
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
	// IP地址
	IPAddress types.TThostFtdcIPAddressType
}

func (d CThostFtdcUserSessionField) Type() string { return "CThostFtdcUserSessionField" }

func (d CThostFtdcUserSessionField) String() string {
	var builder strings.Builder
	builder.Grow(289)

	builder.WriteString("CThostFtdcUserSessionField{")
	fmt.Fprintf(&builder, "FrontID=%+v", d.FrontID)
	fmt.Fprintf(&builder, ", SessionID=%+v", d.SessionID)
	fmt.Fprintf(&builder, ", BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", UserID=%+v", d.UserID)
	fmt.Fprintf(&builder, ", LoginDate=%+v", d.LoginDate)
	fmt.Fprintf(&builder, ", LoginTime=%+v", d.LoginTime)
	fmt.Fprintf(&builder, ", reserve1=%+v", d.reserve1)
	fmt.Fprintf(&builder, ", UserProductInfo=%+v", d.UserProductInfo)
	fmt.Fprintf(&builder, ", InterfaceProductInfo=%+v", d.InterfaceProductInfo)
	fmt.Fprintf(&builder, ", ProtocolInfo=%+v", d.ProtocolInfo)
	fmt.Fprintf(&builder, ", MacAddress=%+v", d.MacAddress)
	fmt.Fprintf(&builder, ", LoginRemark=%+v", d.LoginRemark)
	fmt.Fprintf(&builder, ", IPAddress=%+v", d.IPAddress)

	builder.WriteByte('}')

	return builder.String()
}

// 查询最大报单数量
type CThostFtdcQryMaxOrderVolumeField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	// 保留的无效字段
	reserve1 types.TThostFtdcOldInstrumentIDType
	// 买卖方向
	Direction types.TThostFtdcDirectionType
	// 开平标志
	OffsetFlag types.TThostFtdcOffsetFlagType
	// 投机套保标志
	HedgeFlag types.TThostFtdcHedgeFlagType
	// 最大允许报单数量
	MaxVolume types.TThostFtdcVolumeType
	// 交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	// 投资单元代码
	InvestUnitID types.TThostFtdcInvestUnitIDType
	// 合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
}

func (d CThostFtdcQryMaxOrderVolumeField) Type() string { return "CThostFtdcQryMaxOrderVolumeField" }

func (d CThostFtdcQryMaxOrderVolumeField) String() string {
	var builder strings.Builder
	builder.Grow(229)

	builder.WriteString("CThostFtdcQryMaxOrderVolumeField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", InvestorID=%+v", d.InvestorID)
	fmt.Fprintf(&builder, ", reserve1=%+v", d.reserve1)
	fmt.Fprintf(&builder, ", Direction=%+v", d.Direction)
	fmt.Fprintf(&builder, ", OffsetFlag=%+v", d.OffsetFlag)
	fmt.Fprintf(&builder, ", HedgeFlag=%+v", d.HedgeFlag)
	fmt.Fprintf(&builder, ", MaxVolume=%+v", d.MaxVolume)
	fmt.Fprintf(&builder, ", ExchangeID=%+v", d.ExchangeID)
	fmt.Fprintf(&builder, ", InvestUnitID=%+v", d.InvestUnitID)
	fmt.Fprintf(&builder, ", InstrumentID=%+v", d.InstrumentID)

	builder.WriteByte('}')

	return builder.String()
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
	// 结算编号
	SettlementID types.TThostFtdcSettlementIDType
	// 投资者帐号
	AccountID types.TThostFtdcAccountIDType
	// 币种代码
	CurrencyID types.TThostFtdcCurrencyIDType
}

func (d CThostFtdcSettlementInfoConfirmField) Type() string {
	return "CThostFtdcSettlementInfoConfirmField"
}

func (d CThostFtdcSettlementInfoConfirmField) String() string {
	var builder strings.Builder
	builder.Grow(177)

	builder.WriteString("CThostFtdcSettlementInfoConfirmField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", InvestorID=%+v", d.InvestorID)
	fmt.Fprintf(&builder, ", ConfirmDate=%+v", d.ConfirmDate)
	fmt.Fprintf(&builder, ", ConfirmTime=%+v", d.ConfirmTime)
	fmt.Fprintf(&builder, ", SettlementID=%+v", d.SettlementID)
	fmt.Fprintf(&builder, ", AccountID=%+v", d.AccountID)
	fmt.Fprintf(&builder, ", CurrencyID=%+v", d.CurrencyID)

	builder.WriteByte('}')

	return builder.String()
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
	// 是否是个股期权内转
	IsFromSopt types.TThostFtdcBoolType
	// 资金密码
	TradingPassword types.TThostFtdcPasswordType
	// 是否二级代理商的内转
	IsSecAgentTranfer types.TThostFtdcBoolType
}

func (d CThostFtdcSyncDepositField) Type() string { return "CThostFtdcSyncDepositField" }

func (d CThostFtdcSyncDepositField) String() string {
	var builder strings.Builder
	builder.Grow(212)

	builder.WriteString("CThostFtdcSyncDepositField{")
	fmt.Fprintf(&builder, "DepositSeqNo=%+v", d.DepositSeqNo)
	fmt.Fprintf(&builder, ", BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", InvestorID=%+v", d.InvestorID)
	fmt.Fprintf(&builder, ", Deposit=%+v", d.Deposit)
	fmt.Fprintf(&builder, ", IsForce=%+v", d.IsForce)
	fmt.Fprintf(&builder, ", CurrencyID=%+v", d.CurrencyID)
	fmt.Fprintf(&builder, ", IsFromSopt=%+v", d.IsFromSopt)
	fmt.Fprintf(&builder, ", TradingPassword=%+v", d.TradingPassword)
	fmt.Fprintf(&builder, ", IsSecAgentTranfer=%+v", d.IsSecAgentTranfer)

	builder.WriteByte('}')

	return builder.String()
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

func (d CThostFtdcSyncFundMortgageField) Type() string { return "CThostFtdcSyncFundMortgageField" }

func (d CThostFtdcSyncFundMortgageField) String() string {
	var builder strings.Builder
	builder.Grow(162)

	builder.WriteString("CThostFtdcSyncFundMortgageField{")
	fmt.Fprintf(&builder, "MortgageSeqNo=%+v", d.MortgageSeqNo)
	fmt.Fprintf(&builder, ", BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", InvestorID=%+v", d.InvestorID)
	fmt.Fprintf(&builder, ", FromCurrencyID=%+v", d.FromCurrencyID)
	fmt.Fprintf(&builder, ", MortgageAmount=%+v", d.MortgageAmount)
	fmt.Fprintf(&builder, ", ToCurrencyID=%+v", d.ToCurrencyID)

	builder.WriteByte('}')

	return builder.String()
}

// 经纪公司同步
type CThostFtdcBrokerSyncField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
}

func (d CThostFtdcBrokerSyncField) Type() string { return "CThostFtdcBrokerSyncField" }

func (d CThostFtdcBrokerSyncField) String() string {
	var builder strings.Builder
	builder.Grow(43)

	builder.WriteString("CThostFtdcBrokerSyncField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)

	builder.WriteByte('}')

	return builder.String()
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
	// 是否频率控制
	IsOrderFreq types.TThostFtdcEnumBoolType
	// 是否开仓限制
	IsOpenVolLimit types.TThostFtdcEnumBoolType
}

func (d CThostFtdcSyncingInvestorField) Type() string { return "CThostFtdcSyncingInvestorField" }

func (d CThostFtdcSyncingInvestorField) String() string {
	var builder strings.Builder
	builder.Grow(346)

	builder.WriteString("CThostFtdcSyncingInvestorField{")
	fmt.Fprintf(&builder, "InvestorID=%+v", d.InvestorID)
	fmt.Fprintf(&builder, ", BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", InvestorGroupID=%+v", d.InvestorGroupID)
	fmt.Fprintf(&builder, ", InvestorName=%+v", d.InvestorName)
	fmt.Fprintf(&builder, ", IdentifiedCardType=%+v", d.IdentifiedCardType)
	fmt.Fprintf(&builder, ", IdentifiedCardNo=%+v", d.IdentifiedCardNo)
	fmt.Fprintf(&builder, ", IsActive=%+v", d.IsActive)
	fmt.Fprintf(&builder, ", Telephone=%+v", d.Telephone)
	fmt.Fprintf(&builder, ", Address=%+v", d.Address)
	fmt.Fprintf(&builder, ", OpenDate=%+v", d.OpenDate)
	fmt.Fprintf(&builder, ", Mobile=%+v", d.Mobile)
	fmt.Fprintf(&builder, ", CommModelID=%+v", d.CommModelID)
	fmt.Fprintf(&builder, ", MarginModelID=%+v", d.MarginModelID)
	fmt.Fprintf(&builder, ", IsOrderFreq=%+v", d.IsOrderFreq)
	fmt.Fprintf(&builder, ", IsOpenVolLimit=%+v", d.IsOpenVolLimit)

	builder.WriteByte('}')

	return builder.String()
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

func (d CThostFtdcSyncingTradingCodeField) Type() string { return "CThostFtdcSyncingTradingCodeField" }

func (d CThostFtdcSyncingTradingCodeField) String() string {
	var builder strings.Builder
	builder.Grow(149)

	builder.WriteString("CThostFtdcSyncingTradingCodeField{")
	fmt.Fprintf(&builder, "InvestorID=%+v", d.InvestorID)
	fmt.Fprintf(&builder, ", BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", ExchangeID=%+v", d.ExchangeID)
	fmt.Fprintf(&builder, ", ClientID=%+v", d.ClientID)
	fmt.Fprintf(&builder, ", IsActive=%+v", d.IsActive)
	fmt.Fprintf(&builder, ", ClientIDType=%+v", d.ClientIDType)

	builder.WriteByte('}')

	return builder.String()
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

func (d CThostFtdcSyncingInvestorGroupField) Type() string {
	return "CThostFtdcSyncingInvestorGroupField"
}

func (d CThostFtdcSyncingInvestorGroupField) String() string {
	var builder strings.Builder
	builder.Grow(105)

	builder.WriteString("CThostFtdcSyncingInvestorGroupField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", InvestorGroupID=%+v", d.InvestorGroupID)
	fmt.Fprintf(&builder, ", InvestorGroupName=%+v", d.InvestorGroupName)

	builder.WriteByte('}')

	return builder.String()
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
	// 延时换汇冻结金额
	FrozenSwap types.TThostFtdcMoneyType
	// 剩余换汇额度
	RemainSwap types.TThostFtdcMoneyType
	// 期权市值
	OptionValue types.TThostFtdcMoneyType
}

func (d CThostFtdcSyncingTradingAccountField) Type() string {
	return "CThostFtdcSyncingTradingAccountField"
}

func (d CThostFtdcSyncingTradingAccountField) String() string {
	var builder strings.Builder
	builder.Grow(1184)

	builder.WriteString("CThostFtdcSyncingTradingAccountField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", AccountID=%+v", d.AccountID)
	fmt.Fprintf(&builder, ", PreMortgage=%+v", d.PreMortgage)
	fmt.Fprintf(&builder, ", PreCredit=%+v", d.PreCredit)
	fmt.Fprintf(&builder, ", PreDeposit=%+v", d.PreDeposit)
	fmt.Fprintf(&builder, ", PreBalance=%+v", d.PreBalance)
	fmt.Fprintf(&builder, ", PreMargin=%+v", d.PreMargin)
	fmt.Fprintf(&builder, ", InterestBase=%+v", d.InterestBase)
	fmt.Fprintf(&builder, ", Interest=%+v", d.Interest)
	fmt.Fprintf(&builder, ", Deposit=%+v", d.Deposit)
	fmt.Fprintf(&builder, ", Withdraw=%+v", d.Withdraw)
	fmt.Fprintf(&builder, ", FrozenMargin=%+v", d.FrozenMargin)
	fmt.Fprintf(&builder, ", FrozenCash=%+v", d.FrozenCash)
	fmt.Fprintf(&builder, ", FrozenCommission=%+v", d.FrozenCommission)
	fmt.Fprintf(&builder, ", CurrMargin=%+v", d.CurrMargin)
	fmt.Fprintf(&builder, ", CashIn=%+v", d.CashIn)
	fmt.Fprintf(&builder, ", Commission=%+v", d.Commission)
	fmt.Fprintf(&builder, ", CloseProfit=%+v", d.CloseProfit)
	fmt.Fprintf(&builder, ", PositionProfit=%+v", d.PositionProfit)
	fmt.Fprintf(&builder, ", Balance=%+v", d.Balance)
	fmt.Fprintf(&builder, ", Available=%+v", d.Available)
	fmt.Fprintf(&builder, ", WithdrawQuota=%+v", d.WithdrawQuota)
	fmt.Fprintf(&builder, ", Reserve=%+v", d.Reserve)
	fmt.Fprintf(&builder, ", TradingDay=%+v", d.TradingDay)
	fmt.Fprintf(&builder, ", SettlementID=%+v", d.SettlementID)
	fmt.Fprintf(&builder, ", Credit=%+v", d.Credit)
	fmt.Fprintf(&builder, ", Mortgage=%+v", d.Mortgage)
	fmt.Fprintf(&builder, ", ExchangeMargin=%+v", d.ExchangeMargin)
	fmt.Fprintf(&builder, ", DeliveryMargin=%+v", d.DeliveryMargin)
	fmt.Fprintf(&builder, ", ExchangeDeliveryMargin=%+v", d.ExchangeDeliveryMargin)
	fmt.Fprintf(&builder, ", ReserveBalance=%+v", d.ReserveBalance)
	fmt.Fprintf(&builder, ", CurrencyID=%+v", d.CurrencyID)
	fmt.Fprintf(&builder, ", PreFundMortgageIn=%+v", d.PreFundMortgageIn)
	fmt.Fprintf(&builder, ", PreFundMortgageOut=%+v", d.PreFundMortgageOut)
	fmt.Fprintf(&builder, ", FundMortgageIn=%+v", d.FundMortgageIn)
	fmt.Fprintf(&builder, ", FundMortgageOut=%+v", d.FundMortgageOut)
	fmt.Fprintf(&builder, ", FundMortgageAvailable=%+v", d.FundMortgageAvailable)
	fmt.Fprintf(&builder, ", MortgageableFund=%+v", d.MortgageableFund)
	fmt.Fprintf(&builder, ", SpecProductMargin=%+v", d.SpecProductMargin)
	fmt.Fprintf(&builder, ", SpecProductFrozenMargin=%+v", d.SpecProductFrozenMargin)
	fmt.Fprintf(&builder, ", SpecProductCommission=%+v", d.SpecProductCommission)
	fmt.Fprintf(&builder, ", SpecProductFrozenCommission=%+v", d.SpecProductFrozenCommission)
	fmt.Fprintf(&builder, ", SpecProductPositionProfit=%+v", d.SpecProductPositionProfit)
	fmt.Fprintf(&builder, ", SpecProductCloseProfit=%+v", d.SpecProductCloseProfit)
	fmt.Fprintf(&builder, ", SpecProductPositionProfitByAlg=%+v", d.SpecProductPositionProfitByAlg)
	fmt.Fprintf(&builder, ", SpecProductExchangeMargin=%+v", d.SpecProductExchangeMargin)
	fmt.Fprintf(&builder, ", FrozenSwap=%+v", d.FrozenSwap)
	fmt.Fprintf(&builder, ", RemainSwap=%+v", d.RemainSwap)
	fmt.Fprintf(&builder, ", OptionValue=%+v", d.OptionValue)

	builder.WriteByte('}')

	return builder.String()
}

// 正在同步中的投资者持仓
type CThostFtdcSyncingInvestorPositionField struct {
	// 保留的无效字段
	reserve1 types.TThostFtdcOldInstrumentIDType
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
	// 交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	// 执行冻结的昨仓
	YdStrikeFrozen types.TThostFtdcVolumeType
	// 投资单元代码
	InvestUnitID types.TThostFtdcInvestUnitIDType
	// 持仓成本差值
	PositionCostOffset types.TThostFtdcMoneyType
	// tas持仓手数
	TasPosition types.TThostFtdcVolumeType
	// tas持仓成本
	TasPositionCost types.TThostFtdcMoneyType
	// 合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
}

func (d CThostFtdcSyncingInvestorPositionField) Type() string {
	return "CThostFtdcSyncingInvestorPositionField"
}

func (d CThostFtdcSyncingInvestorPositionField) String() string {
	var builder strings.Builder
	builder.Grow(1157)

	builder.WriteString("CThostFtdcSyncingInvestorPositionField{")
	fmt.Fprintf(&builder, "reserve1=%+v", d.reserve1)
	fmt.Fprintf(&builder, ", BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", InvestorID=%+v", d.InvestorID)
	fmt.Fprintf(&builder, ", PosiDirection=%+v", d.PosiDirection)
	fmt.Fprintf(&builder, ", HedgeFlag=%+v", d.HedgeFlag)
	fmt.Fprintf(&builder, ", PositionDate=%+v", d.PositionDate)
	fmt.Fprintf(&builder, ", YdPosition=%+v", d.YdPosition)
	fmt.Fprintf(&builder, ", Position=%+v", d.Position)
	fmt.Fprintf(&builder, ", LongFrozen=%+v", d.LongFrozen)
	fmt.Fprintf(&builder, ", ShortFrozen=%+v", d.ShortFrozen)
	fmt.Fprintf(&builder, ", LongFrozenAmount=%+v", d.LongFrozenAmount)
	fmt.Fprintf(&builder, ", ShortFrozenAmount=%+v", d.ShortFrozenAmount)
	fmt.Fprintf(&builder, ", OpenVolume=%+v", d.OpenVolume)
	fmt.Fprintf(&builder, ", CloseVolume=%+v", d.CloseVolume)
	fmt.Fprintf(&builder, ", OpenAmount=%+v", d.OpenAmount)
	fmt.Fprintf(&builder, ", CloseAmount=%+v", d.CloseAmount)
	fmt.Fprintf(&builder, ", PositionCost=%+v", d.PositionCost)
	fmt.Fprintf(&builder, ", PreMargin=%+v", d.PreMargin)
	fmt.Fprintf(&builder, ", UseMargin=%+v", d.UseMargin)
	fmt.Fprintf(&builder, ", FrozenMargin=%+v", d.FrozenMargin)
	fmt.Fprintf(&builder, ", FrozenCash=%+v", d.FrozenCash)
	fmt.Fprintf(&builder, ", FrozenCommission=%+v", d.FrozenCommission)
	fmt.Fprintf(&builder, ", CashIn=%+v", d.CashIn)
	fmt.Fprintf(&builder, ", Commission=%+v", d.Commission)
	fmt.Fprintf(&builder, ", CloseProfit=%+v", d.CloseProfit)
	fmt.Fprintf(&builder, ", PositionProfit=%+v", d.PositionProfit)
	fmt.Fprintf(&builder, ", PreSettlementPrice=%+v", d.PreSettlementPrice)
	fmt.Fprintf(&builder, ", SettlementPrice=%+v", d.SettlementPrice)
	fmt.Fprintf(&builder, ", TradingDay=%+v", d.TradingDay)
	fmt.Fprintf(&builder, ", SettlementID=%+v", d.SettlementID)
	fmt.Fprintf(&builder, ", OpenCost=%+v", d.OpenCost)
	fmt.Fprintf(&builder, ", ExchangeMargin=%+v", d.ExchangeMargin)
	fmt.Fprintf(&builder, ", CombPosition=%+v", d.CombPosition)
	fmt.Fprintf(&builder, ", CombLongFrozen=%+v", d.CombLongFrozen)
	fmt.Fprintf(&builder, ", CombShortFrozen=%+v", d.CombShortFrozen)
	fmt.Fprintf(&builder, ", CloseProfitByDate=%+v", d.CloseProfitByDate)
	fmt.Fprintf(&builder, ", CloseProfitByTrade=%+v", d.CloseProfitByTrade)
	fmt.Fprintf(&builder, ", TodayPosition=%+v", d.TodayPosition)
	fmt.Fprintf(&builder, ", MarginRateByMoney=%+v", d.MarginRateByMoney)
	fmt.Fprintf(&builder, ", MarginRateByVolume=%+v", d.MarginRateByVolume)
	fmt.Fprintf(&builder, ", StrikeFrozen=%+v", d.StrikeFrozen)
	fmt.Fprintf(&builder, ", StrikeFrozenAmount=%+v", d.StrikeFrozenAmount)
	fmt.Fprintf(&builder, ", AbandonFrozen=%+v", d.AbandonFrozen)
	fmt.Fprintf(&builder, ", ExchangeID=%+v", d.ExchangeID)
	fmt.Fprintf(&builder, ", YdStrikeFrozen=%+v", d.YdStrikeFrozen)
	fmt.Fprintf(&builder, ", InvestUnitID=%+v", d.InvestUnitID)
	fmt.Fprintf(&builder, ", PositionCostOffset=%+v", d.PositionCostOffset)
	fmt.Fprintf(&builder, ", TasPosition=%+v", d.TasPosition)
	fmt.Fprintf(&builder, ", TasPositionCost=%+v", d.TasPositionCost)
	fmt.Fprintf(&builder, ", InstrumentID=%+v", d.InstrumentID)

	builder.WriteByte('}')

	return builder.String()
}

// 正在同步中的合约保证金率
type CThostFtdcSyncingInstrumentMarginRateField struct {
	// 保留的无效字段
	reserve1 types.TThostFtdcOldInstrumentIDType
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
	// 合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
}

func (d CThostFtdcSyncingInstrumentMarginRateField) Type() string {
	return "CThostFtdcSyncingInstrumentMarginRateField"
}

func (d CThostFtdcSyncingInstrumentMarginRateField) String() string {
	var builder strings.Builder
	builder.Grow(314)

	builder.WriteString("CThostFtdcSyncingInstrumentMarginRateField{")
	fmt.Fprintf(&builder, "reserve1=%+v", d.reserve1)
	fmt.Fprintf(&builder, ", InvestorRange=%+v", d.InvestorRange)
	fmt.Fprintf(&builder, ", BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", InvestorID=%+v", d.InvestorID)
	fmt.Fprintf(&builder, ", HedgeFlag=%+v", d.HedgeFlag)
	fmt.Fprintf(&builder, ", LongMarginRatioByMoney=%+v", d.LongMarginRatioByMoney)
	fmt.Fprintf(&builder, ", LongMarginRatioByVolume=%+v", d.LongMarginRatioByVolume)
	fmt.Fprintf(&builder, ", ShortMarginRatioByMoney=%+v", d.ShortMarginRatioByMoney)
	fmt.Fprintf(&builder, ", ShortMarginRatioByVolume=%+v", d.ShortMarginRatioByVolume)
	fmt.Fprintf(&builder, ", IsRelative=%+v", d.IsRelative)
	fmt.Fprintf(&builder, ", InstrumentID=%+v", d.InstrumentID)

	builder.WriteByte('}')

	return builder.String()
}

// 正在同步中的合约手续费率
type CThostFtdcSyncingInstrumentCommissionRateField struct {
	// 保留的无效字段
	reserve1 types.TThostFtdcOldInstrumentIDType
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
	// 合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
}

func (d CThostFtdcSyncingInstrumentCommissionRateField) Type() string {
	return "CThostFtdcSyncingInstrumentCommissionRateField"
}

func (d CThostFtdcSyncingInstrumentCommissionRateField) String() string {
	var builder strings.Builder
	builder.Grow(320)

	builder.WriteString("CThostFtdcSyncingInstrumentCommissionRateField{")
	fmt.Fprintf(&builder, "reserve1=%+v", d.reserve1)
	fmt.Fprintf(&builder, ", InvestorRange=%+v", d.InvestorRange)
	fmt.Fprintf(&builder, ", BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", InvestorID=%+v", d.InvestorID)
	fmt.Fprintf(&builder, ", OpenRatioByMoney=%+v", d.OpenRatioByMoney)
	fmt.Fprintf(&builder, ", OpenRatioByVolume=%+v", d.OpenRatioByVolume)
	fmt.Fprintf(&builder, ", CloseRatioByMoney=%+v", d.CloseRatioByMoney)
	fmt.Fprintf(&builder, ", CloseRatioByVolume=%+v", d.CloseRatioByVolume)
	fmt.Fprintf(&builder, ", CloseTodayRatioByMoney=%+v", d.CloseTodayRatioByMoney)
	fmt.Fprintf(&builder, ", CloseTodayRatioByVolume=%+v", d.CloseTodayRatioByVolume)
	fmt.Fprintf(&builder, ", InstrumentID=%+v", d.InstrumentID)

	builder.WriteByte('}')

	return builder.String()
}

// 正在同步中的合约交易权限
type CThostFtdcSyncingInstrumentTradingRightField struct {
	// 保留的无效字段
	reserve1 types.TThostFtdcOldInstrumentIDType
	// 投资者范围
	InvestorRange types.TThostFtdcInvestorRangeType
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	// 交易权限
	TradingRight types.TThostFtdcTradingRightType
	// 合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
}

func (d CThostFtdcSyncingInstrumentTradingRightField) Type() string {
	return "CThostFtdcSyncingInstrumentTradingRightField"
}

func (d CThostFtdcSyncingInstrumentTradingRightField) String() string {
	var builder strings.Builder
	builder.Grow(167)

	builder.WriteString("CThostFtdcSyncingInstrumentTradingRightField{")
	fmt.Fprintf(&builder, "reserve1=%+v", d.reserve1)
	fmt.Fprintf(&builder, ", InvestorRange=%+v", d.InvestorRange)
	fmt.Fprintf(&builder, ", BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", InvestorID=%+v", d.InvestorID)
	fmt.Fprintf(&builder, ", TradingRight=%+v", d.TradingRight)
	fmt.Fprintf(&builder, ", InstrumentID=%+v", d.InstrumentID)

	builder.WriteByte('}')

	return builder.String()
}

// 查询报单
type CThostFtdcQryOrderField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	// 保留的无效字段
	reserve1 types.TThostFtdcOldInstrumentIDType
	// 交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	// 报单编号
	OrderSysID types.TThostFtdcOrderSysIDType
	// 开始时间
	InsertTimeStart types.TThostFtdcTimeType
	// 结束时间
	InsertTimeEnd types.TThostFtdcTimeType
	// 投资单元代码
	InvestUnitID types.TThostFtdcInvestUnitIDType
	// 合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
}

func (d CThostFtdcQryOrderField) Type() string { return "CThostFtdcQryOrderField" }

func (d CThostFtdcQryOrderField) String() string {
	var builder strings.Builder
	builder.Grow(211)

	builder.WriteString("CThostFtdcQryOrderField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", InvestorID=%+v", d.InvestorID)
	fmt.Fprintf(&builder, ", reserve1=%+v", d.reserve1)
	fmt.Fprintf(&builder, ", ExchangeID=%+v", d.ExchangeID)
	fmt.Fprintf(&builder, ", OrderSysID=%+v", d.OrderSysID)
	fmt.Fprintf(&builder, ", InsertTimeStart=%+v", d.InsertTimeStart)
	fmt.Fprintf(&builder, ", InsertTimeEnd=%+v", d.InsertTimeEnd)
	fmt.Fprintf(&builder, ", InvestUnitID=%+v", d.InvestUnitID)
	fmt.Fprintf(&builder, ", InstrumentID=%+v", d.InstrumentID)

	builder.WriteByte('}')

	return builder.String()
}

// 查询成交
type CThostFtdcQryTradeField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	// 保留的无效字段
	reserve1 types.TThostFtdcOldInstrumentIDType
	// 交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	// 成交编号
	TradeID types.TThostFtdcTradeIDType
	// 开始时间
	TradeTimeStart types.TThostFtdcTimeType
	// 结束时间
	TradeTimeEnd types.TThostFtdcTimeType
	// 投资单元代码
	InvestUnitID types.TThostFtdcInvestUnitIDType
	// 合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
}

func (d CThostFtdcQryTradeField) Type() string { return "CThostFtdcQryTradeField" }

func (d CThostFtdcQryTradeField) String() string {
	var builder strings.Builder
	builder.Grow(206)

	builder.WriteString("CThostFtdcQryTradeField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", InvestorID=%+v", d.InvestorID)
	fmt.Fprintf(&builder, ", reserve1=%+v", d.reserve1)
	fmt.Fprintf(&builder, ", ExchangeID=%+v", d.ExchangeID)
	fmt.Fprintf(&builder, ", TradeID=%+v", d.TradeID)
	fmt.Fprintf(&builder, ", TradeTimeStart=%+v", d.TradeTimeStart)
	fmt.Fprintf(&builder, ", TradeTimeEnd=%+v", d.TradeTimeEnd)
	fmt.Fprintf(&builder, ", InvestUnitID=%+v", d.InvestUnitID)
	fmt.Fprintf(&builder, ", InstrumentID=%+v", d.InstrumentID)

	builder.WriteByte('}')

	return builder.String()
}

// 查询投资者持仓
type CThostFtdcQryInvestorPositionField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	// 保留的无效字段
	reserve1 types.TThostFtdcOldInstrumentIDType
	// 交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	// 投资单元代码
	InvestUnitID types.TThostFtdcInvestUnitIDType
	// 合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
}

func (d CThostFtdcQryInvestorPositionField) Type() string {
	return "CThostFtdcQryInvestorPositionField"
}

func (d CThostFtdcQryInvestorPositionField) String() string {
	var builder strings.Builder
	builder.Grow(154)

	builder.WriteString("CThostFtdcQryInvestorPositionField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", InvestorID=%+v", d.InvestorID)
	fmt.Fprintf(&builder, ", reserve1=%+v", d.reserve1)
	fmt.Fprintf(&builder, ", ExchangeID=%+v", d.ExchangeID)
	fmt.Fprintf(&builder, ", InvestUnitID=%+v", d.InvestUnitID)
	fmt.Fprintf(&builder, ", InstrumentID=%+v", d.InstrumentID)

	builder.WriteByte('}')

	return builder.String()
}

// 查询资金账户
type CThostFtdcQryTradingAccountField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	// 币种代码
	CurrencyID types.TThostFtdcCurrencyIDType
	// 业务类型
	BizType types.TThostFtdcBizTypeType
	// 投资者帐号
	AccountID types.TThostFtdcAccountIDType
}

func (d CThostFtdcQryTradingAccountField) Type() string { return "CThostFtdcQryTradingAccountField" }

func (d CThostFtdcQryTradingAccountField) String() string {
	var builder strings.Builder
	builder.Grow(126)

	builder.WriteString("CThostFtdcQryTradingAccountField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", InvestorID=%+v", d.InvestorID)
	fmt.Fprintf(&builder, ", CurrencyID=%+v", d.CurrencyID)
	fmt.Fprintf(&builder, ", BizType=%+v", d.BizType)
	fmt.Fprintf(&builder, ", AccountID=%+v", d.AccountID)

	builder.WriteByte('}')

	return builder.String()
}

// 查询投资者
type CThostFtdcQryInvestorField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID types.TThostFtdcInvestorIDType
}

func (d CThostFtdcQryInvestorField) Type() string { return "CThostFtdcQryInvestorField" }

func (d CThostFtdcQryInvestorField) String() string {
	var builder strings.Builder
	builder.Grow(64)

	builder.WriteString("CThostFtdcQryInvestorField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", InvestorID=%+v", d.InvestorID)

	builder.WriteByte('}')

	return builder.String()
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
	// 投资单元代码
	InvestUnitID types.TThostFtdcInvestUnitIDType
}

func (d CThostFtdcQryTradingCodeField) Type() string { return "CThostFtdcQryTradingCodeField" }

func (d CThostFtdcQryTradingCodeField) String() string {
	var builder strings.Builder
	builder.Grow(149)

	builder.WriteString("CThostFtdcQryTradingCodeField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", InvestorID=%+v", d.InvestorID)
	fmt.Fprintf(&builder, ", ExchangeID=%+v", d.ExchangeID)
	fmt.Fprintf(&builder, ", ClientID=%+v", d.ClientID)
	fmt.Fprintf(&builder, ", ClientIDType=%+v", d.ClientIDType)
	fmt.Fprintf(&builder, ", InvestUnitID=%+v", d.InvestUnitID)

	builder.WriteByte('}')

	return builder.String()
}

// 查询投资者组
type CThostFtdcQryInvestorGroupField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
}

func (d CThostFtdcQryInvestorGroupField) Type() string { return "CThostFtdcQryInvestorGroupField" }

func (d CThostFtdcQryInvestorGroupField) String() string {
	var builder strings.Builder
	builder.Grow(49)

	builder.WriteString("CThostFtdcQryInvestorGroupField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)

	builder.WriteByte('}')

	return builder.String()
}

// 查询合约保证金率
type CThostFtdcQryInstrumentMarginRateField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	// 保留的无效字段
	reserve1 types.TThostFtdcOldInstrumentIDType
	// 投机套保标志
	HedgeFlag types.TThostFtdcHedgeFlagType
	// 交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	// 投资单元代码
	InvestUnitID types.TThostFtdcInvestUnitIDType
	// 合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
}

func (d CThostFtdcQryInstrumentMarginRateField) Type() string {
	return "CThostFtdcQryInstrumentMarginRateField"
}

func (d CThostFtdcQryInstrumentMarginRateField) String() string {
	var builder strings.Builder
	builder.Grow(177)

	builder.WriteString("CThostFtdcQryInstrumentMarginRateField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", InvestorID=%+v", d.InvestorID)
	fmt.Fprintf(&builder, ", reserve1=%+v", d.reserve1)
	fmt.Fprintf(&builder, ", HedgeFlag=%+v", d.HedgeFlag)
	fmt.Fprintf(&builder, ", ExchangeID=%+v", d.ExchangeID)
	fmt.Fprintf(&builder, ", InvestUnitID=%+v", d.InvestUnitID)
	fmt.Fprintf(&builder, ", InstrumentID=%+v", d.InstrumentID)

	builder.WriteByte('}')

	return builder.String()
}

// 查询手续费率
type CThostFtdcQryInstrumentCommissionRateField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	// 保留的无效字段
	reserve1 types.TThostFtdcOldInstrumentIDType
	// 交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	// 投资单元代码
	InvestUnitID types.TThostFtdcInvestUnitIDType
	// 合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
}

func (d CThostFtdcQryInstrumentCommissionRateField) Type() string {
	return "CThostFtdcQryInstrumentCommissionRateField"
}

func (d CThostFtdcQryInstrumentCommissionRateField) String() string {
	var builder strings.Builder
	builder.Grow(162)

	builder.WriteString("CThostFtdcQryInstrumentCommissionRateField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", InvestorID=%+v", d.InvestorID)
	fmt.Fprintf(&builder, ", reserve1=%+v", d.reserve1)
	fmt.Fprintf(&builder, ", ExchangeID=%+v", d.ExchangeID)
	fmt.Fprintf(&builder, ", InvestUnitID=%+v", d.InvestUnitID)
	fmt.Fprintf(&builder, ", InstrumentID=%+v", d.InstrumentID)

	builder.WriteByte('}')

	return builder.String()
}

// 查询合约交易权限
type CThostFtdcQryInstrumentTradingRightField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	// 保留的无效字段
	reserve1 types.TThostFtdcOldInstrumentIDType
	// 合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
}

func (d CThostFtdcQryInstrumentTradingRightField) Type() string {
	return "CThostFtdcQryInstrumentTradingRightField"
}

func (d CThostFtdcQryInstrumentTradingRightField) String() string {
	var builder strings.Builder
	builder.Grow(118)

	builder.WriteString("CThostFtdcQryInstrumentTradingRightField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", InvestorID=%+v", d.InvestorID)
	fmt.Fprintf(&builder, ", reserve1=%+v", d.reserve1)
	fmt.Fprintf(&builder, ", InstrumentID=%+v", d.InstrumentID)

	builder.WriteByte('}')

	return builder.String()
}

// 查询经纪公司
type CThostFtdcQryBrokerField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
}

func (d CThostFtdcQryBrokerField) Type() string { return "CThostFtdcQryBrokerField" }

func (d CThostFtdcQryBrokerField) String() string {
	var builder strings.Builder
	builder.Grow(42)

	builder.WriteString("CThostFtdcQryBrokerField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)

	builder.WriteByte('}')

	return builder.String()
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

func (d CThostFtdcQryTraderField) Type() string { return "CThostFtdcQryTraderField" }

func (d CThostFtdcQryTraderField) String() string {
	var builder strings.Builder
	builder.Grow(85)

	builder.WriteString("CThostFtdcQryTraderField{")
	fmt.Fprintf(&builder, "ExchangeID=%+v", d.ExchangeID)
	fmt.Fprintf(&builder, ", ParticipantID=%+v", d.ParticipantID)
	fmt.Fprintf(&builder, ", TraderID=%+v", d.TraderID)

	builder.WriteByte('}')

	return builder.String()
}

// 查询管理用户功能权限
type CThostFtdcQrySuperUserFunctionField struct {
	// 用户代码
	UserID types.TThostFtdcUserIDType
}

func (d CThostFtdcQrySuperUserFunctionField) Type() string {
	return "CThostFtdcQrySuperUserFunctionField"
}

func (d CThostFtdcQrySuperUserFunctionField) String() string {
	var builder strings.Builder
	builder.Grow(51)

	builder.WriteString("CThostFtdcQrySuperUserFunctionField{")
	fmt.Fprintf(&builder, "UserID=%+v", d.UserID)

	builder.WriteByte('}')

	return builder.String()
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

func (d CThostFtdcQryUserSessionField) Type() string { return "CThostFtdcQryUserSessionField" }

func (d CThostFtdcQryUserSessionField) String() string {
	var builder strings.Builder
	builder.Grow(99)

	builder.WriteString("CThostFtdcQryUserSessionField{")
	fmt.Fprintf(&builder, "FrontID=%+v", d.FrontID)
	fmt.Fprintf(&builder, ", SessionID=%+v", d.SessionID)
	fmt.Fprintf(&builder, ", BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", UserID=%+v", d.UserID)

	builder.WriteByte('}')

	return builder.String()
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

func (d CThostFtdcQryPartBrokerField) Type() string { return "CThostFtdcQryPartBrokerField" }

func (d CThostFtdcQryPartBrokerField) String() string {
	var builder strings.Builder
	builder.Grow(89)

	builder.WriteString("CThostFtdcQryPartBrokerField{")
	fmt.Fprintf(&builder, "ExchangeID=%+v", d.ExchangeID)
	fmt.Fprintf(&builder, ", BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", ParticipantID=%+v", d.ParticipantID)

	builder.WriteByte('}')

	return builder.String()
}

// 查询前置状态
type CThostFtdcQryFrontStatusField struct {
	// 前置编号
	FrontID types.TThostFtdcFrontIDType
}

func (d CThostFtdcQryFrontStatusField) Type() string { return "CThostFtdcQryFrontStatusField" }

func (d CThostFtdcQryFrontStatusField) String() string {
	var builder strings.Builder
	builder.Grow(46)

	builder.WriteString("CThostFtdcQryFrontStatusField{")
	fmt.Fprintf(&builder, "FrontID=%+v", d.FrontID)

	builder.WriteByte('}')

	return builder.String()
}

// 查询交易所报单
type CThostFtdcQryExchangeOrderField struct {
	// 会员代码
	ParticipantID types.TThostFtdcParticipantIDType
	// 客户代码
	ClientID types.TThostFtdcClientIDType
	// 保留的无效字段
	reserve1 types.TThostFtdcOldExchangeInstIDType
	// 交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	// 交易所交易员代码
	TraderID types.TThostFtdcTraderIDType
	// 合约在交易所的代码
	ExchangeInstID types.TThostFtdcExchangeInstIDType
}

func (d CThostFtdcQryExchangeOrderField) Type() string { return "CThostFtdcQryExchangeOrderField" }

func (d CThostFtdcQryExchangeOrderField) String() string {
	var builder strings.Builder
	builder.Grow(152)

	builder.WriteString("CThostFtdcQryExchangeOrderField{")
	fmt.Fprintf(&builder, "ParticipantID=%+v", d.ParticipantID)
	fmt.Fprintf(&builder, ", ClientID=%+v", d.ClientID)
	fmt.Fprintf(&builder, ", reserve1=%+v", d.reserve1)
	fmt.Fprintf(&builder, ", ExchangeID=%+v", d.ExchangeID)
	fmt.Fprintf(&builder, ", TraderID=%+v", d.TraderID)
	fmt.Fprintf(&builder, ", ExchangeInstID=%+v", d.ExchangeInstID)

	builder.WriteByte('}')

	return builder.String()
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

func (d CThostFtdcQryOrderActionField) Type() string { return "CThostFtdcQryOrderActionField" }

func (d CThostFtdcQryOrderActionField) String() string {
	var builder strings.Builder
	builder.Grow(87)

	builder.WriteString("CThostFtdcQryOrderActionField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", InvestorID=%+v", d.InvestorID)
	fmt.Fprintf(&builder, ", ExchangeID=%+v", d.ExchangeID)

	builder.WriteByte('}')

	return builder.String()
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

func (d CThostFtdcQryExchangeOrderActionField) Type() string {
	return "CThostFtdcQryExchangeOrderActionField"
}

func (d CThostFtdcQryExchangeOrderActionField) String() string {
	var builder strings.Builder
	builder.Grow(116)

	builder.WriteString("CThostFtdcQryExchangeOrderActionField{")
	fmt.Fprintf(&builder, "ParticipantID=%+v", d.ParticipantID)
	fmt.Fprintf(&builder, ", ClientID=%+v", d.ClientID)
	fmt.Fprintf(&builder, ", ExchangeID=%+v", d.ExchangeID)
	fmt.Fprintf(&builder, ", TraderID=%+v", d.TraderID)

	builder.WriteByte('}')

	return builder.String()
}

// 查询管理用户
type CThostFtdcQrySuperUserField struct {
	// 用户代码
	UserID types.TThostFtdcUserIDType
}

func (d CThostFtdcQrySuperUserField) Type() string { return "CThostFtdcQrySuperUserField" }

func (d CThostFtdcQrySuperUserField) String() string {
	var builder strings.Builder
	builder.Grow(43)

	builder.WriteString("CThostFtdcQrySuperUserField{")
	fmt.Fprintf(&builder, "UserID=%+v", d.UserID)

	builder.WriteByte('}')

	return builder.String()
}

// 查询交易所
type CThostFtdcQryExchangeField struct {
	// 交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
}

func (d CThostFtdcQryExchangeField) Type() string { return "CThostFtdcQryExchangeField" }

func (d CThostFtdcQryExchangeField) String() string {
	var builder strings.Builder
	builder.Grow(46)

	builder.WriteString("CThostFtdcQryExchangeField{")
	fmt.Fprintf(&builder, "ExchangeID=%+v", d.ExchangeID)

	builder.WriteByte('}')

	return builder.String()
}

// 查询产品
type CThostFtdcQryProductField struct {
	// 保留的无效字段
	reserve1 types.TThostFtdcOldInstrumentIDType
	// 产品类型
	ProductClass types.TThostFtdcProductClassType
	// 交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	// 产品代码
	ProductID types.TThostFtdcInstrumentIDType
}

func (d CThostFtdcQryProductField) Type() string { return "CThostFtdcQryProductField" }

func (d CThostFtdcQryProductField) String() string {
	var builder strings.Builder
	builder.Grow(104)

	builder.WriteString("CThostFtdcQryProductField{")
	fmt.Fprintf(&builder, "reserve1=%+v", d.reserve1)
	fmt.Fprintf(&builder, ", ProductClass=%+v", d.ProductClass)
	fmt.Fprintf(&builder, ", ExchangeID=%+v", d.ExchangeID)
	fmt.Fprintf(&builder, ", ProductID=%+v", d.ProductID)

	builder.WriteByte('}')

	return builder.String()
}

// 查询合约
type CThostFtdcQryInstrumentField struct {
	// 保留的无效字段
	reserve1 types.TThostFtdcOldInstrumentIDType
	// 交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	// 保留的无效字段
	reserve2 types.TThostFtdcOldExchangeInstIDType
	// 保留的无效字段
	reserve3 types.TThostFtdcOldInstrumentIDType
	// 合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
	// 合约在交易所的代码
	ExchangeInstID types.TThostFtdcExchangeInstIDType
	// 产品代码
	ProductID types.TThostFtdcInstrumentIDType
}

func (d CThostFtdcQryInstrumentField) Type() string { return "CThostFtdcQryInstrumentField" }

func (d CThostFtdcQryInstrumentField) String() string {
	var builder strings.Builder
	builder.Grow(167)

	builder.WriteString("CThostFtdcQryInstrumentField{")
	fmt.Fprintf(&builder, "reserve1=%+v", d.reserve1)
	fmt.Fprintf(&builder, ", ExchangeID=%+v", d.ExchangeID)
	fmt.Fprintf(&builder, ", reserve2=%+v", d.reserve2)
	fmt.Fprintf(&builder, ", reserve3=%+v", d.reserve3)
	fmt.Fprintf(&builder, ", InstrumentID=%+v", d.InstrumentID)
	fmt.Fprintf(&builder, ", ExchangeInstID=%+v", d.ExchangeInstID)
	fmt.Fprintf(&builder, ", ProductID=%+v", d.ProductID)

	builder.WriteByte('}')

	return builder.String()
}

// 查询行情
type CThostFtdcQryDepthMarketDataField struct {
	// 保留的无效字段
	reserve1 types.TThostFtdcOldInstrumentIDType
	// 交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	// 合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
	// 产品类型
	ProductClass types.TThostFtdcProductClassType
}

func (d CThostFtdcQryDepthMarketDataField) Type() string { return "CThostFtdcQryDepthMarketDataField" }

func (d CThostFtdcQryDepthMarketDataField) String() string {
	var builder strings.Builder
	builder.Grow(115)

	builder.WriteString("CThostFtdcQryDepthMarketDataField{")
	fmt.Fprintf(&builder, "reserve1=%+v", d.reserve1)
	fmt.Fprintf(&builder, ", ExchangeID=%+v", d.ExchangeID)
	fmt.Fprintf(&builder, ", InstrumentID=%+v", d.InstrumentID)
	fmt.Fprintf(&builder, ", ProductClass=%+v", d.ProductClass)

	builder.WriteByte('}')

	return builder.String()
}

// 查询经纪公司用户
type CThostFtdcQryBrokerUserField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 用户代码
	UserID types.TThostFtdcUserIDType
}

func (d CThostFtdcQryBrokerUserField) Type() string { return "CThostFtdcQryBrokerUserField" }

func (d CThostFtdcQryBrokerUserField) String() string {
	var builder strings.Builder
	builder.Grow(62)

	builder.WriteString("CThostFtdcQryBrokerUserField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", UserID=%+v", d.UserID)

	builder.WriteByte('}')

	return builder.String()
}

// 查询经纪公司用户权限
type CThostFtdcQryBrokerUserFunctionField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 用户代码
	UserID types.TThostFtdcUserIDType
}

func (d CThostFtdcQryBrokerUserFunctionField) Type() string {
	return "CThostFtdcQryBrokerUserFunctionField"
}

func (d CThostFtdcQryBrokerUserFunctionField) String() string {
	var builder strings.Builder
	builder.Grow(70)

	builder.WriteString("CThostFtdcQryBrokerUserFunctionField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", UserID=%+v", d.UserID)

	builder.WriteByte('}')

	return builder.String()
}

// 查询交易员报盘机
type CThostFtdcQryTraderOfferField struct {
	// 交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	// 会员代码
	ParticipantID types.TThostFtdcParticipantIDType
	// 交易所交易员代码
	TraderID types.TThostFtdcTraderIDType
}

func (d CThostFtdcQryTraderOfferField) Type() string { return "CThostFtdcQryTraderOfferField" }

func (d CThostFtdcQryTraderOfferField) String() string {
	var builder strings.Builder
	builder.Grow(90)

	builder.WriteString("CThostFtdcQryTraderOfferField{")
	fmt.Fprintf(&builder, "ExchangeID=%+v", d.ExchangeID)
	fmt.Fprintf(&builder, ", ParticipantID=%+v", d.ParticipantID)
	fmt.Fprintf(&builder, ", TraderID=%+v", d.TraderID)

	builder.WriteByte('}')

	return builder.String()
}

// 查询出入金流水
type CThostFtdcQrySyncDepositField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 出入金流水号
	DepositSeqNo types.TThostFtdcDepositSeqNoType
}

func (d CThostFtdcQrySyncDepositField) Type() string { return "CThostFtdcQrySyncDepositField" }

func (d CThostFtdcQrySyncDepositField) String() string {
	var builder strings.Builder
	builder.Grow(69)

	builder.WriteString("CThostFtdcQrySyncDepositField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", DepositSeqNo=%+v", d.DepositSeqNo)

	builder.WriteByte('}')

	return builder.String()
}

// 查询投资者结算结果
type CThostFtdcQrySettlementInfoField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	// 交易日
	TradingDay types.TThostFtdcDateType
	// 投资者帐号
	AccountID types.TThostFtdcAccountIDType
	// 币种代码
	CurrencyID types.TThostFtdcCurrencyIDType
}

func (d CThostFtdcQrySettlementInfoField) Type() string { return "CThostFtdcQrySettlementInfoField" }

func (d CThostFtdcQrySettlementInfoField) String() string {
	var builder strings.Builder
	builder.Grow(129)

	builder.WriteString("CThostFtdcQrySettlementInfoField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", InvestorID=%+v", d.InvestorID)
	fmt.Fprintf(&builder, ", TradingDay=%+v", d.TradingDay)
	fmt.Fprintf(&builder, ", AccountID=%+v", d.AccountID)
	fmt.Fprintf(&builder, ", CurrencyID=%+v", d.CurrencyID)

	builder.WriteByte('}')

	return builder.String()
}

// 查询交易所保证金率
type CThostFtdcQryExchangeMarginRateField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 保留的无效字段
	reserve1 types.TThostFtdcOldInstrumentIDType
	// 投机套保标志
	HedgeFlag types.TThostFtdcHedgeFlagType
	// 交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	// 合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
}

func (d CThostFtdcQryExchangeMarginRateField) Type() string {
	return "CThostFtdcQryExchangeMarginRateField"
}

func (d CThostFtdcQryExchangeMarginRateField) String() string {
	var builder strings.Builder
	builder.Grow(133)

	builder.WriteString("CThostFtdcQryExchangeMarginRateField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", reserve1=%+v", d.reserve1)
	fmt.Fprintf(&builder, ", HedgeFlag=%+v", d.HedgeFlag)
	fmt.Fprintf(&builder, ", ExchangeID=%+v", d.ExchangeID)
	fmt.Fprintf(&builder, ", InstrumentID=%+v", d.InstrumentID)

	builder.WriteByte('}')

	return builder.String()
}

// 查询交易所调整保证金率
type CThostFtdcQryExchangeMarginRateAdjustField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 保留的无效字段
	reserve1 types.TThostFtdcOldInstrumentIDType
	// 投机套保标志
	HedgeFlag types.TThostFtdcHedgeFlagType
	// 合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
}

func (d CThostFtdcQryExchangeMarginRateAdjustField) Type() string {
	return "CThostFtdcQryExchangeMarginRateAdjustField"
}

func (d CThostFtdcQryExchangeMarginRateAdjustField) String() string {
	var builder strings.Builder
	builder.Grow(119)

	builder.WriteString("CThostFtdcQryExchangeMarginRateAdjustField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", reserve1=%+v", d.reserve1)
	fmt.Fprintf(&builder, ", HedgeFlag=%+v", d.HedgeFlag)
	fmt.Fprintf(&builder, ", InstrumentID=%+v", d.InstrumentID)

	builder.WriteByte('}')

	return builder.String()
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

func (d CThostFtdcQryExchangeRateField) Type() string { return "CThostFtdcQryExchangeRateField" }

func (d CThostFtdcQryExchangeRateField) String() string {
	var builder strings.Builder
	builder.Grow(94)

	builder.WriteString("CThostFtdcQryExchangeRateField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", FromCurrencyID=%+v", d.FromCurrencyID)
	fmt.Fprintf(&builder, ", ToCurrencyID=%+v", d.ToCurrencyID)

	builder.WriteByte('}')

	return builder.String()
}

// 查询货币质押流水
type CThostFtdcQrySyncFundMortgageField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 货币质押流水号
	MortgageSeqNo types.TThostFtdcDepositSeqNoType
}

func (d CThostFtdcQrySyncFundMortgageField) Type() string {
	return "CThostFtdcQrySyncFundMortgageField"
}

func (d CThostFtdcQrySyncFundMortgageField) String() string {
	var builder strings.Builder
	builder.Grow(75)

	builder.WriteString("CThostFtdcQrySyncFundMortgageField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", MortgageSeqNo=%+v", d.MortgageSeqNo)

	builder.WriteByte('}')

	return builder.String()
}

// 查询报单
type CThostFtdcQryHisOrderField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	// 保留的无效字段
	reserve1 types.TThostFtdcOldInstrumentIDType
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
	// 合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
}

func (d CThostFtdcQryHisOrderField) Type() string { return "CThostFtdcQryHisOrderField" }

func (d CThostFtdcQryHisOrderField) String() string {
	var builder strings.Builder
	builder.Grow(234)

	builder.WriteString("CThostFtdcQryHisOrderField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", InvestorID=%+v", d.InvestorID)
	fmt.Fprintf(&builder, ", reserve1=%+v", d.reserve1)
	fmt.Fprintf(&builder, ", ExchangeID=%+v", d.ExchangeID)
	fmt.Fprintf(&builder, ", OrderSysID=%+v", d.OrderSysID)
	fmt.Fprintf(&builder, ", InsertTimeStart=%+v", d.InsertTimeStart)
	fmt.Fprintf(&builder, ", InsertTimeEnd=%+v", d.InsertTimeEnd)
	fmt.Fprintf(&builder, ", TradingDay=%+v", d.TradingDay)
	fmt.Fprintf(&builder, ", SettlementID=%+v", d.SettlementID)
	fmt.Fprintf(&builder, ", InstrumentID=%+v", d.InstrumentID)

	builder.WriteByte('}')

	return builder.String()
}

// 当前期权合约最小保证金
type CThostFtdcOptionInstrMiniMarginField struct {
	// 保留的无效字段
	reserve1 types.TThostFtdcOldInstrumentIDType
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
	// 合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
}

func (d CThostFtdcOptionInstrMiniMarginField) Type() string {
	return "CThostFtdcOptionInstrMiniMarginField"
}

func (d CThostFtdcOptionInstrMiniMarginField) String() string {
	var builder strings.Builder
	builder.Grow(197)

	builder.WriteString("CThostFtdcOptionInstrMiniMarginField{")
	fmt.Fprintf(&builder, "reserve1=%+v", d.reserve1)
	fmt.Fprintf(&builder, ", InvestorRange=%+v", d.InvestorRange)
	fmt.Fprintf(&builder, ", BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", InvestorID=%+v", d.InvestorID)
	fmt.Fprintf(&builder, ", MinMargin=%+v", d.MinMargin)
	fmt.Fprintf(&builder, ", ValueMethod=%+v", d.ValueMethod)
	fmt.Fprintf(&builder, ", IsRelative=%+v", d.IsRelative)
	fmt.Fprintf(&builder, ", InstrumentID=%+v", d.InstrumentID)

	builder.WriteByte('}')

	return builder.String()
}

// 当前期权合约保证金调整系数
type CThostFtdcOptionInstrMarginAdjustField struct {
	// 保留的无效字段
	reserve1 types.TThostFtdcOldInstrumentIDType
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
	// 合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
}

func (d CThostFtdcOptionInstrMarginAdjustField) Type() string {
	return "CThostFtdcOptionInstrMarginAdjustField"
}

func (d CThostFtdcOptionInstrMarginAdjustField) String() string {
	var builder strings.Builder
	builder.Grow(435)

	builder.WriteString("CThostFtdcOptionInstrMarginAdjustField{")
	fmt.Fprintf(&builder, "reserve1=%+v", d.reserve1)
	fmt.Fprintf(&builder, ", InvestorRange=%+v", d.InvestorRange)
	fmt.Fprintf(&builder, ", BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", InvestorID=%+v", d.InvestorID)
	fmt.Fprintf(&builder, ", SShortMarginRatioByMoney=%+v", d.SShortMarginRatioByMoney)
	fmt.Fprintf(&builder, ", SShortMarginRatioByVolume=%+v", d.SShortMarginRatioByVolume)
	fmt.Fprintf(&builder, ", HShortMarginRatioByMoney=%+v", d.HShortMarginRatioByMoney)
	fmt.Fprintf(&builder, ", HShortMarginRatioByVolume=%+v", d.HShortMarginRatioByVolume)
	fmt.Fprintf(&builder, ", AShortMarginRatioByMoney=%+v", d.AShortMarginRatioByMoney)
	fmt.Fprintf(&builder, ", AShortMarginRatioByVolume=%+v", d.AShortMarginRatioByVolume)
	fmt.Fprintf(&builder, ", IsRelative=%+v", d.IsRelative)
	fmt.Fprintf(&builder, ", MShortMarginRatioByMoney=%+v", d.MShortMarginRatioByMoney)
	fmt.Fprintf(&builder, ", MShortMarginRatioByVolume=%+v", d.MShortMarginRatioByVolume)
	fmt.Fprintf(&builder, ", InstrumentID=%+v", d.InstrumentID)

	builder.WriteByte('}')

	return builder.String()
}

// 当前期权合约手续费的详细内容
type CThostFtdcOptionInstrCommRateField struct {
	// 保留的无效字段
	reserve1 types.TThostFtdcOldInstrumentIDType
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
	// 交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	// 投资单元代码
	InvestUnitID types.TThostFtdcInvestUnitIDType
	// 合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
}

func (d CThostFtdcOptionInstrCommRateField) Type() string {
	return "CThostFtdcOptionInstrCommRateField"
}

func (d CThostFtdcOptionInstrCommRateField) String() string {
	var builder strings.Builder
	builder.Grow(407)

	builder.WriteString("CThostFtdcOptionInstrCommRateField{")
	fmt.Fprintf(&builder, "reserve1=%+v", d.reserve1)
	fmt.Fprintf(&builder, ", InvestorRange=%+v", d.InvestorRange)
	fmt.Fprintf(&builder, ", BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", InvestorID=%+v", d.InvestorID)
	fmt.Fprintf(&builder, ", OpenRatioByMoney=%+v", d.OpenRatioByMoney)
	fmt.Fprintf(&builder, ", OpenRatioByVolume=%+v", d.OpenRatioByVolume)
	fmt.Fprintf(&builder, ", CloseRatioByMoney=%+v", d.CloseRatioByMoney)
	fmt.Fprintf(&builder, ", CloseRatioByVolume=%+v", d.CloseRatioByVolume)
	fmt.Fprintf(&builder, ", CloseTodayRatioByMoney=%+v", d.CloseTodayRatioByMoney)
	fmt.Fprintf(&builder, ", CloseTodayRatioByVolume=%+v", d.CloseTodayRatioByVolume)
	fmt.Fprintf(&builder, ", StrikeRatioByMoney=%+v", d.StrikeRatioByMoney)
	fmt.Fprintf(&builder, ", StrikeRatioByVolume=%+v", d.StrikeRatioByVolume)
	fmt.Fprintf(&builder, ", ExchangeID=%+v", d.ExchangeID)
	fmt.Fprintf(&builder, ", InvestUnitID=%+v", d.InvestUnitID)
	fmt.Fprintf(&builder, ", InstrumentID=%+v", d.InstrumentID)

	builder.WriteByte('}')

	return builder.String()
}

// 期权交易成本
type CThostFtdcOptionInstrTradeCostField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	// 保留的无效字段
	reserve1 types.TThostFtdcOldInstrumentIDType
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
	// 交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	// 投资单元代码
	InvestUnitID types.TThostFtdcInvestUnitIDType
	// 合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
}

func (d CThostFtdcOptionInstrTradeCostField) Type() string {
	return "CThostFtdcOptionInstrTradeCostField"
}

func (d CThostFtdcOptionInstrTradeCostField) String() string {
	var builder strings.Builder
	builder.Grow(281)

	builder.WriteString("CThostFtdcOptionInstrTradeCostField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", InvestorID=%+v", d.InvestorID)
	fmt.Fprintf(&builder, ", reserve1=%+v", d.reserve1)
	fmt.Fprintf(&builder, ", HedgeFlag=%+v", d.HedgeFlag)
	fmt.Fprintf(&builder, ", FixedMargin=%+v", d.FixedMargin)
	fmt.Fprintf(&builder, ", MiniMargin=%+v", d.MiniMargin)
	fmt.Fprintf(&builder, ", Royalty=%+v", d.Royalty)
	fmt.Fprintf(&builder, ", ExchFixedMargin=%+v", d.ExchFixedMargin)
	fmt.Fprintf(&builder, ", ExchMiniMargin=%+v", d.ExchMiniMargin)
	fmt.Fprintf(&builder, ", ExchangeID=%+v", d.ExchangeID)
	fmt.Fprintf(&builder, ", InvestUnitID=%+v", d.InvestUnitID)
	fmt.Fprintf(&builder, ", InstrumentID=%+v", d.InstrumentID)

	builder.WriteByte('}')

	return builder.String()
}

// 期权交易成本查询
type CThostFtdcQryOptionInstrTradeCostField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	// 保留的无效字段
	reserve1 types.TThostFtdcOldInstrumentIDType
	// 投机套保标志
	HedgeFlag types.TThostFtdcHedgeFlagType
	// 期权合约报价
	InputPrice types.TThostFtdcPriceType
	// 标的价格,填0则用昨结算价
	UnderlyingPrice types.TThostFtdcPriceType
	// 交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	// 投资单元代码
	InvestUnitID types.TThostFtdcInvestUnitIDType
	// 合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
}

func (d CThostFtdcQryOptionInstrTradeCostField) Type() string {
	return "CThostFtdcQryOptionInstrTradeCostField"
}

func (d CThostFtdcQryOptionInstrTradeCostField) String() string {
	var builder strings.Builder
	builder.Grow(222)

	builder.WriteString("CThostFtdcQryOptionInstrTradeCostField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", InvestorID=%+v", d.InvestorID)
	fmt.Fprintf(&builder, ", reserve1=%+v", d.reserve1)
	fmt.Fprintf(&builder, ", HedgeFlag=%+v", d.HedgeFlag)
	fmt.Fprintf(&builder, ", InputPrice=%+v", d.InputPrice)
	fmt.Fprintf(&builder, ", UnderlyingPrice=%+v", d.UnderlyingPrice)
	fmt.Fprintf(&builder, ", ExchangeID=%+v", d.ExchangeID)
	fmt.Fprintf(&builder, ", InvestUnitID=%+v", d.InvestUnitID)
	fmt.Fprintf(&builder, ", InstrumentID=%+v", d.InstrumentID)

	builder.WriteByte('}')

	return builder.String()
}

// 期权手续费率查询
type CThostFtdcQryOptionInstrCommRateField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	// 保留的无效字段
	reserve1 types.TThostFtdcOldInstrumentIDType
	// 交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	// 投资单元代码
	InvestUnitID types.TThostFtdcInvestUnitIDType
	// 合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
}

func (d CThostFtdcQryOptionInstrCommRateField) Type() string {
	return "CThostFtdcQryOptionInstrCommRateField"
}

func (d CThostFtdcQryOptionInstrCommRateField) String() string {
	var builder strings.Builder
	builder.Grow(157)

	builder.WriteString("CThostFtdcQryOptionInstrCommRateField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", InvestorID=%+v", d.InvestorID)
	fmt.Fprintf(&builder, ", reserve1=%+v", d.reserve1)
	fmt.Fprintf(&builder, ", ExchangeID=%+v", d.ExchangeID)
	fmt.Fprintf(&builder, ", InvestUnitID=%+v", d.InvestUnitID)
	fmt.Fprintf(&builder, ", InstrumentID=%+v", d.InstrumentID)

	builder.WriteByte('}')

	return builder.String()
}

// 股指现货指数
type CThostFtdcIndexPriceField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 保留的无效字段
	reserve1 types.TThostFtdcOldInstrumentIDType
	// 指数现货收盘价
	ClosePrice types.TThostFtdcPriceType
	// 合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
}

func (d CThostFtdcIndexPriceField) Type() string { return "CThostFtdcIndexPriceField" }

func (d CThostFtdcIndexPriceField) String() string {
	var builder strings.Builder
	builder.Grow(103)

	builder.WriteString("CThostFtdcIndexPriceField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", reserve1=%+v", d.reserve1)
	fmt.Fprintf(&builder, ", ClosePrice=%+v", d.ClosePrice)
	fmt.Fprintf(&builder, ", InstrumentID=%+v", d.InstrumentID)

	builder.WriteByte('}')

	return builder.String()
}

// 输入的执行宣告
type CThostFtdcInputExecOrderField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	// 保留的无效字段
	reserve1 types.TThostFtdcOldInstrumentIDType
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
	// 保留的无效字段
	reserve2 types.TThostFtdcOldIPAddressType
	// Mac地址
	MacAddress types.TThostFtdcMacAddressType
	// 合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
	// IP地址
	IPAddress types.TThostFtdcIPAddressType
}

func (d CThostFtdcInputExecOrderField) Type() string { return "CThostFtdcInputExecOrderField" }

func (d CThostFtdcInputExecOrderField) String() string {
	var builder strings.Builder
	builder.Grow(488)

	builder.WriteString("CThostFtdcInputExecOrderField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", InvestorID=%+v", d.InvestorID)
	fmt.Fprintf(&builder, ", reserve1=%+v", d.reserve1)
	fmt.Fprintf(&builder, ", ExecOrderRef=%+v", d.ExecOrderRef)
	fmt.Fprintf(&builder, ", UserID=%+v", d.UserID)
	fmt.Fprintf(&builder, ", Volume=%+v", d.Volume)
	fmt.Fprintf(&builder, ", RequestID=%+v", d.RequestID)
	fmt.Fprintf(&builder, ", BusinessUnit=%+v", d.BusinessUnit)
	fmt.Fprintf(&builder, ", OffsetFlag=%+v", d.OffsetFlag)
	fmt.Fprintf(&builder, ", HedgeFlag=%+v", d.HedgeFlag)
	fmt.Fprintf(&builder, ", ActionType=%+v", d.ActionType)
	fmt.Fprintf(&builder, ", PosiDirection=%+v", d.PosiDirection)
	fmt.Fprintf(&builder, ", ReservePositionFlag=%+v", d.ReservePositionFlag)
	fmt.Fprintf(&builder, ", CloseFlag=%+v", d.CloseFlag)
	fmt.Fprintf(&builder, ", ExchangeID=%+v", d.ExchangeID)
	fmt.Fprintf(&builder, ", InvestUnitID=%+v", d.InvestUnitID)
	fmt.Fprintf(&builder, ", AccountID=%+v", d.AccountID)
	fmt.Fprintf(&builder, ", CurrencyID=%+v", d.CurrencyID)
	fmt.Fprintf(&builder, ", ClientID=%+v", d.ClientID)
	fmt.Fprintf(&builder, ", reserve2=%+v", d.reserve2)
	fmt.Fprintf(&builder, ", MacAddress=%+v", d.MacAddress)
	fmt.Fprintf(&builder, ", InstrumentID=%+v", d.InstrumentID)
	fmt.Fprintf(&builder, ", IPAddress=%+v", d.IPAddress)

	builder.WriteByte('}')

	return builder.String()
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
	// 保留的无效字段
	reserve1 types.TThostFtdcOldInstrumentIDType
	// 投资单元代码
	InvestUnitID types.TThostFtdcInvestUnitIDType
	// 保留的无效字段
	reserve2 types.TThostFtdcOldIPAddressType
	// Mac地址
	MacAddress types.TThostFtdcMacAddressType
	// 合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
	// IP地址
	IPAddress types.TThostFtdcIPAddressType
}

func (d CThostFtdcInputExecOrderActionField) Type() string {
	return "CThostFtdcInputExecOrderActionField"
}

func (d CThostFtdcInputExecOrderActionField) String() string {
	var builder strings.Builder
	builder.Grow(377)

	builder.WriteString("CThostFtdcInputExecOrderActionField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", InvestorID=%+v", d.InvestorID)
	fmt.Fprintf(&builder, ", ExecOrderActionRef=%+v", d.ExecOrderActionRef)
	fmt.Fprintf(&builder, ", ExecOrderRef=%+v", d.ExecOrderRef)
	fmt.Fprintf(&builder, ", RequestID=%+v", d.RequestID)
	fmt.Fprintf(&builder, ", FrontID=%+v", d.FrontID)
	fmt.Fprintf(&builder, ", SessionID=%+v", d.SessionID)
	fmt.Fprintf(&builder, ", ExchangeID=%+v", d.ExchangeID)
	fmt.Fprintf(&builder, ", ExecOrderSysID=%+v", d.ExecOrderSysID)
	fmt.Fprintf(&builder, ", ActionFlag=%+v", d.ActionFlag)
	fmt.Fprintf(&builder, ", UserID=%+v", d.UserID)
	fmt.Fprintf(&builder, ", reserve1=%+v", d.reserve1)
	fmt.Fprintf(&builder, ", InvestUnitID=%+v", d.InvestUnitID)
	fmt.Fprintf(&builder, ", reserve2=%+v", d.reserve2)
	fmt.Fprintf(&builder, ", MacAddress=%+v", d.MacAddress)
	fmt.Fprintf(&builder, ", InstrumentID=%+v", d.InstrumentID)
	fmt.Fprintf(&builder, ", IPAddress=%+v", d.IPAddress)

	builder.WriteByte('}')

	return builder.String()
}

// 执行宣告
type CThostFtdcExecOrderField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	// 保留的无效字段
	reserve1 types.TThostFtdcOldInstrumentIDType
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
	// 本地执行宣告编号
	ExecOrderLocalID types.TThostFtdcOrderLocalIDType
	// 交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	// 会员代码
	ParticipantID types.TThostFtdcParticipantIDType
	// 客户代码
	ClientID types.TThostFtdcClientIDType
	// 保留的无效字段
	reserve2 types.TThostFtdcOldExchangeInstIDType
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
	// 保留的无效字段
	reserve3 types.TThostFtdcOldIPAddressType
	// Mac地址
	MacAddress types.TThostFtdcMacAddressType
	// 合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
	// 合约在交易所的代码
	ExchangeInstID types.TThostFtdcExchangeInstIDType
	// IP地址
	IPAddress types.TThostFtdcIPAddressType
}

func (d CThostFtdcExecOrderField) Type() string { return "CThostFtdcExecOrderField" }

func (d CThostFtdcExecOrderField) String() string {
	var builder strings.Builder
	builder.Grow(1000)

	builder.WriteString("CThostFtdcExecOrderField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", InvestorID=%+v", d.InvestorID)
	fmt.Fprintf(&builder, ", reserve1=%+v", d.reserve1)
	fmt.Fprintf(&builder, ", ExecOrderRef=%+v", d.ExecOrderRef)
	fmt.Fprintf(&builder, ", UserID=%+v", d.UserID)
	fmt.Fprintf(&builder, ", Volume=%+v", d.Volume)
	fmt.Fprintf(&builder, ", RequestID=%+v", d.RequestID)
	fmt.Fprintf(&builder, ", BusinessUnit=%+v", d.BusinessUnit)
	fmt.Fprintf(&builder, ", OffsetFlag=%+v", d.OffsetFlag)
	fmt.Fprintf(&builder, ", HedgeFlag=%+v", d.HedgeFlag)
	fmt.Fprintf(&builder, ", ActionType=%+v", d.ActionType)
	fmt.Fprintf(&builder, ", PosiDirection=%+v", d.PosiDirection)
	fmt.Fprintf(&builder, ", ReservePositionFlag=%+v", d.ReservePositionFlag)
	fmt.Fprintf(&builder, ", CloseFlag=%+v", d.CloseFlag)
	fmt.Fprintf(&builder, ", ExecOrderLocalID=%+v", d.ExecOrderLocalID)
	fmt.Fprintf(&builder, ", ExchangeID=%+v", d.ExchangeID)
	fmt.Fprintf(&builder, ", ParticipantID=%+v", d.ParticipantID)
	fmt.Fprintf(&builder, ", ClientID=%+v", d.ClientID)
	fmt.Fprintf(&builder, ", reserve2=%+v", d.reserve2)
	fmt.Fprintf(&builder, ", TraderID=%+v", d.TraderID)
	fmt.Fprintf(&builder, ", InstallID=%+v", d.InstallID)
	fmt.Fprintf(&builder, ", OrderSubmitStatus=%+v", d.OrderSubmitStatus)
	fmt.Fprintf(&builder, ", NotifySequence=%+v", d.NotifySequence)
	fmt.Fprintf(&builder, ", TradingDay=%+v", d.TradingDay)
	fmt.Fprintf(&builder, ", SettlementID=%+v", d.SettlementID)
	fmt.Fprintf(&builder, ", ExecOrderSysID=%+v", d.ExecOrderSysID)
	fmt.Fprintf(&builder, ", InsertDate=%+v", d.InsertDate)
	fmt.Fprintf(&builder, ", InsertTime=%+v", d.InsertTime)
	fmt.Fprintf(&builder, ", CancelTime=%+v", d.CancelTime)
	fmt.Fprintf(&builder, ", ExecResult=%+v", d.ExecResult)
	fmt.Fprintf(&builder, ", ClearingPartID=%+v", d.ClearingPartID)
	fmt.Fprintf(&builder, ", SequenceNo=%+v", d.SequenceNo)
	fmt.Fprintf(&builder, ", FrontID=%+v", d.FrontID)
	fmt.Fprintf(&builder, ", SessionID=%+v", d.SessionID)
	fmt.Fprintf(&builder, ", UserProductInfo=%+v", d.UserProductInfo)
	fmt.Fprintf(&builder, ", StatusMsg=%+v", d.StatusMsg)
	fmt.Fprintf(&builder, ", ActiveUserID=%+v", d.ActiveUserID)
	fmt.Fprintf(&builder, ", BrokerExecOrderSeq=%+v", d.BrokerExecOrderSeq)
	fmt.Fprintf(&builder, ", BranchID=%+v", d.BranchID)
	fmt.Fprintf(&builder, ", InvestUnitID=%+v", d.InvestUnitID)
	fmt.Fprintf(&builder, ", AccountID=%+v", d.AccountID)
	fmt.Fprintf(&builder, ", CurrencyID=%+v", d.CurrencyID)
	fmt.Fprintf(&builder, ", reserve3=%+v", d.reserve3)
	fmt.Fprintf(&builder, ", MacAddress=%+v", d.MacAddress)
	fmt.Fprintf(&builder, ", InstrumentID=%+v", d.InstrumentID)
	fmt.Fprintf(&builder, ", ExchangeInstID=%+v", d.ExchangeInstID)
	fmt.Fprintf(&builder, ", IPAddress=%+v", d.IPAddress)

	builder.WriteByte('}')

	return builder.String()
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
	// 保留的无效字段
	reserve1 types.TThostFtdcOldInstrumentIDType
	// 营业部编号
	BranchID types.TThostFtdcBranchIDType
	// 投资单元代码
	InvestUnitID types.TThostFtdcInvestUnitIDType
	// 保留的无效字段
	reserve2 types.TThostFtdcOldIPAddressType
	// Mac地址
	MacAddress types.TThostFtdcMacAddressType
	// 合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
	// IP地址
	IPAddress types.TThostFtdcIPAddressType
}

func (d CThostFtdcExecOrderActionField) Type() string { return "CThostFtdcExecOrderActionField" }

func (d CThostFtdcExecOrderActionField) String() string {
	var builder strings.Builder
	builder.Grow(645)

	builder.WriteString("CThostFtdcExecOrderActionField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", InvestorID=%+v", d.InvestorID)
	fmt.Fprintf(&builder, ", ExecOrderActionRef=%+v", d.ExecOrderActionRef)
	fmt.Fprintf(&builder, ", ExecOrderRef=%+v", d.ExecOrderRef)
	fmt.Fprintf(&builder, ", RequestID=%+v", d.RequestID)
	fmt.Fprintf(&builder, ", FrontID=%+v", d.FrontID)
	fmt.Fprintf(&builder, ", SessionID=%+v", d.SessionID)
	fmt.Fprintf(&builder, ", ExchangeID=%+v", d.ExchangeID)
	fmt.Fprintf(&builder, ", ExecOrderSysID=%+v", d.ExecOrderSysID)
	fmt.Fprintf(&builder, ", ActionFlag=%+v", d.ActionFlag)
	fmt.Fprintf(&builder, ", ActionDate=%+v", d.ActionDate)
	fmt.Fprintf(&builder, ", ActionTime=%+v", d.ActionTime)
	fmt.Fprintf(&builder, ", TraderID=%+v", d.TraderID)
	fmt.Fprintf(&builder, ", InstallID=%+v", d.InstallID)
	fmt.Fprintf(&builder, ", ExecOrderLocalID=%+v", d.ExecOrderLocalID)
	fmt.Fprintf(&builder, ", ActionLocalID=%+v", d.ActionLocalID)
	fmt.Fprintf(&builder, ", ParticipantID=%+v", d.ParticipantID)
	fmt.Fprintf(&builder, ", ClientID=%+v", d.ClientID)
	fmt.Fprintf(&builder, ", BusinessUnit=%+v", d.BusinessUnit)
	fmt.Fprintf(&builder, ", OrderActionStatus=%+v", d.OrderActionStatus)
	fmt.Fprintf(&builder, ", UserID=%+v", d.UserID)
	fmt.Fprintf(&builder, ", ActionType=%+v", d.ActionType)
	fmt.Fprintf(&builder, ", StatusMsg=%+v", d.StatusMsg)
	fmt.Fprintf(&builder, ", reserve1=%+v", d.reserve1)
	fmt.Fprintf(&builder, ", BranchID=%+v", d.BranchID)
	fmt.Fprintf(&builder, ", InvestUnitID=%+v", d.InvestUnitID)
	fmt.Fprintf(&builder, ", reserve2=%+v", d.reserve2)
	fmt.Fprintf(&builder, ", MacAddress=%+v", d.MacAddress)
	fmt.Fprintf(&builder, ", InstrumentID=%+v", d.InstrumentID)
	fmt.Fprintf(&builder, ", IPAddress=%+v", d.IPAddress)

	builder.WriteByte('}')

	return builder.String()
}

// 执行宣告查询
type CThostFtdcQryExecOrderField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	// 保留的无效字段
	reserve1 types.TThostFtdcOldInstrumentIDType
	// 交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	// 执行宣告编号
	ExecOrderSysID types.TThostFtdcExecOrderSysIDType
	// 开始时间
	InsertTimeStart types.TThostFtdcTimeType
	// 结束时间
	InsertTimeEnd types.TThostFtdcTimeType
	// 合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
}

func (d CThostFtdcQryExecOrderField) Type() string { return "CThostFtdcQryExecOrderField" }

func (d CThostFtdcQryExecOrderField) String() string {
	var builder strings.Builder
	builder.Grow(197)

	builder.WriteString("CThostFtdcQryExecOrderField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", InvestorID=%+v", d.InvestorID)
	fmt.Fprintf(&builder, ", reserve1=%+v", d.reserve1)
	fmt.Fprintf(&builder, ", ExchangeID=%+v", d.ExchangeID)
	fmt.Fprintf(&builder, ", ExecOrderSysID=%+v", d.ExecOrderSysID)
	fmt.Fprintf(&builder, ", InsertTimeStart=%+v", d.InsertTimeStart)
	fmt.Fprintf(&builder, ", InsertTimeEnd=%+v", d.InsertTimeEnd)
	fmt.Fprintf(&builder, ", InstrumentID=%+v", d.InstrumentID)

	builder.WriteByte('}')

	return builder.String()
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
	// 期权行权后是否保留期货头寸的标记,该字段已废弃
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
	// 保留的无效字段
	reserve1 types.TThostFtdcOldExchangeInstIDType
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
	// 保留的无效字段
	reserve2 types.TThostFtdcOldIPAddressType
	// Mac地址
	MacAddress types.TThostFtdcMacAddressType
	// 合约在交易所的代码
	ExchangeInstID types.TThostFtdcExchangeInstIDType
	// IP地址
	IPAddress types.TThostFtdcIPAddressType
}

func (d CThostFtdcExchangeExecOrderField) Type() string { return "CThostFtdcExchangeExecOrderField" }

func (d CThostFtdcExchangeExecOrderField) String() string {
	var builder strings.Builder
	builder.Grow(701)

	builder.WriteString("CThostFtdcExchangeExecOrderField{")
	fmt.Fprintf(&builder, "Volume=%+v", d.Volume)
	fmt.Fprintf(&builder, ", RequestID=%+v", d.RequestID)
	fmt.Fprintf(&builder, ", BusinessUnit=%+v", d.BusinessUnit)
	fmt.Fprintf(&builder, ", OffsetFlag=%+v", d.OffsetFlag)
	fmt.Fprintf(&builder, ", HedgeFlag=%+v", d.HedgeFlag)
	fmt.Fprintf(&builder, ", ActionType=%+v", d.ActionType)
	fmt.Fprintf(&builder, ", PosiDirection=%+v", d.PosiDirection)
	fmt.Fprintf(&builder, ", ReservePositionFlag=%+v", d.ReservePositionFlag)
	fmt.Fprintf(&builder, ", CloseFlag=%+v", d.CloseFlag)
	fmt.Fprintf(&builder, ", ExecOrderLocalID=%+v", d.ExecOrderLocalID)
	fmt.Fprintf(&builder, ", ExchangeID=%+v", d.ExchangeID)
	fmt.Fprintf(&builder, ", ParticipantID=%+v", d.ParticipantID)
	fmt.Fprintf(&builder, ", ClientID=%+v", d.ClientID)
	fmt.Fprintf(&builder, ", reserve1=%+v", d.reserve1)
	fmt.Fprintf(&builder, ", TraderID=%+v", d.TraderID)
	fmt.Fprintf(&builder, ", InstallID=%+v", d.InstallID)
	fmt.Fprintf(&builder, ", OrderSubmitStatus=%+v", d.OrderSubmitStatus)
	fmt.Fprintf(&builder, ", NotifySequence=%+v", d.NotifySequence)
	fmt.Fprintf(&builder, ", TradingDay=%+v", d.TradingDay)
	fmt.Fprintf(&builder, ", SettlementID=%+v", d.SettlementID)
	fmt.Fprintf(&builder, ", ExecOrderSysID=%+v", d.ExecOrderSysID)
	fmt.Fprintf(&builder, ", InsertDate=%+v", d.InsertDate)
	fmt.Fprintf(&builder, ", InsertTime=%+v", d.InsertTime)
	fmt.Fprintf(&builder, ", CancelTime=%+v", d.CancelTime)
	fmt.Fprintf(&builder, ", ExecResult=%+v", d.ExecResult)
	fmt.Fprintf(&builder, ", ClearingPartID=%+v", d.ClearingPartID)
	fmt.Fprintf(&builder, ", SequenceNo=%+v", d.SequenceNo)
	fmt.Fprintf(&builder, ", BranchID=%+v", d.BranchID)
	fmt.Fprintf(&builder, ", reserve2=%+v", d.reserve2)
	fmt.Fprintf(&builder, ", MacAddress=%+v", d.MacAddress)
	fmt.Fprintf(&builder, ", ExchangeInstID=%+v", d.ExchangeInstID)
	fmt.Fprintf(&builder, ", IPAddress=%+v", d.IPAddress)

	builder.WriteByte('}')

	return builder.String()
}

// 交易所执行宣告查询
type CThostFtdcQryExchangeExecOrderField struct {
	// 会员代码
	ParticipantID types.TThostFtdcParticipantIDType
	// 客户代码
	ClientID types.TThostFtdcClientIDType
	// 保留的无效字段
	reserve1 types.TThostFtdcOldExchangeInstIDType
	// 交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	// 交易所交易员代码
	TraderID types.TThostFtdcTraderIDType
	// 合约在交易所的代码
	ExchangeInstID types.TThostFtdcExchangeInstIDType
}

func (d CThostFtdcQryExchangeExecOrderField) Type() string {
	return "CThostFtdcQryExchangeExecOrderField"
}

func (d CThostFtdcQryExchangeExecOrderField) String() string {
	var builder strings.Builder
	builder.Grow(156)

	builder.WriteString("CThostFtdcQryExchangeExecOrderField{")
	fmt.Fprintf(&builder, "ParticipantID=%+v", d.ParticipantID)
	fmt.Fprintf(&builder, ", ClientID=%+v", d.ClientID)
	fmt.Fprintf(&builder, ", reserve1=%+v", d.reserve1)
	fmt.Fprintf(&builder, ", ExchangeID=%+v", d.ExchangeID)
	fmt.Fprintf(&builder, ", TraderID=%+v", d.TraderID)
	fmt.Fprintf(&builder, ", ExchangeInstID=%+v", d.ExchangeInstID)

	builder.WriteByte('}')

	return builder.String()
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

func (d CThostFtdcQryExecOrderActionField) Type() string { return "CThostFtdcQryExecOrderActionField" }

func (d CThostFtdcQryExecOrderActionField) String() string {
	var builder strings.Builder
	builder.Grow(91)

	builder.WriteString("CThostFtdcQryExecOrderActionField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", InvestorID=%+v", d.InvestorID)
	fmt.Fprintf(&builder, ", ExchangeID=%+v", d.ExchangeID)

	builder.WriteByte('}')

	return builder.String()
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
	// 保留的无效字段
	reserve1 types.TThostFtdcOldIPAddressType
	// Mac地址
	MacAddress types.TThostFtdcMacAddressType
	// 保留的无效字段
	reserve2 types.TThostFtdcOldExchangeInstIDType
	// 数量
	Volume types.TThostFtdcVolumeType
	// IP地址
	IPAddress types.TThostFtdcIPAddressType
	// 合约在交易所的代码
	ExchangeInstID types.TThostFtdcExchangeInstIDType
}

func (d CThostFtdcExchangeExecOrderActionField) Type() string {
	return "CThostFtdcExchangeExecOrderActionField"
}

func (d CThostFtdcExchangeExecOrderActionField) String() string {
	var builder strings.Builder
	builder.Grow(487)

	builder.WriteString("CThostFtdcExchangeExecOrderActionField{")
	fmt.Fprintf(&builder, "ExchangeID=%+v", d.ExchangeID)
	fmt.Fprintf(&builder, ", ExecOrderSysID=%+v", d.ExecOrderSysID)
	fmt.Fprintf(&builder, ", ActionFlag=%+v", d.ActionFlag)
	fmt.Fprintf(&builder, ", ActionDate=%+v", d.ActionDate)
	fmt.Fprintf(&builder, ", ActionTime=%+v", d.ActionTime)
	fmt.Fprintf(&builder, ", TraderID=%+v", d.TraderID)
	fmt.Fprintf(&builder, ", InstallID=%+v", d.InstallID)
	fmt.Fprintf(&builder, ", ExecOrderLocalID=%+v", d.ExecOrderLocalID)
	fmt.Fprintf(&builder, ", ActionLocalID=%+v", d.ActionLocalID)
	fmt.Fprintf(&builder, ", ParticipantID=%+v", d.ParticipantID)
	fmt.Fprintf(&builder, ", ClientID=%+v", d.ClientID)
	fmt.Fprintf(&builder, ", BusinessUnit=%+v", d.BusinessUnit)
	fmt.Fprintf(&builder, ", OrderActionStatus=%+v", d.OrderActionStatus)
	fmt.Fprintf(&builder, ", UserID=%+v", d.UserID)
	fmt.Fprintf(&builder, ", ActionType=%+v", d.ActionType)
	fmt.Fprintf(&builder, ", BranchID=%+v", d.BranchID)
	fmt.Fprintf(&builder, ", reserve1=%+v", d.reserve1)
	fmt.Fprintf(&builder, ", MacAddress=%+v", d.MacAddress)
	fmt.Fprintf(&builder, ", reserve2=%+v", d.reserve2)
	fmt.Fprintf(&builder, ", Volume=%+v", d.Volume)
	fmt.Fprintf(&builder, ", IPAddress=%+v", d.IPAddress)
	fmt.Fprintf(&builder, ", ExchangeInstID=%+v", d.ExchangeInstID)

	builder.WriteByte('}')

	return builder.String()
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

func (d CThostFtdcQryExchangeExecOrderActionField) Type() string {
	return "CThostFtdcQryExchangeExecOrderActionField"
}

func (d CThostFtdcQryExchangeExecOrderActionField) String() string {
	var builder strings.Builder
	builder.Grow(120)

	builder.WriteString("CThostFtdcQryExchangeExecOrderActionField{")
	fmt.Fprintf(&builder, "ParticipantID=%+v", d.ParticipantID)
	fmt.Fprintf(&builder, ", ClientID=%+v", d.ClientID)
	fmt.Fprintf(&builder, ", ExchangeID=%+v", d.ExchangeID)
	fmt.Fprintf(&builder, ", TraderID=%+v", d.TraderID)

	builder.WriteByte('}')

	return builder.String()
}

// 错误执行宣告
type CThostFtdcErrExecOrderField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	// 保留的无效字段
	reserve1 types.TThostFtdcOldInstrumentIDType
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
	// 保留的无效字段
	reserve2 types.TThostFtdcOldIPAddressType
	// Mac地址
	MacAddress types.TThostFtdcMacAddressType
	// 错误代码
	ErrorID types.TThostFtdcErrorIDType
	// 错误信息
	ErrorMsg types.TThostFtdcErrorMsgType
	// 合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
	// IP地址
	IPAddress types.TThostFtdcIPAddressType
}

func (d CThostFtdcErrExecOrderField) Type() string { return "CThostFtdcErrExecOrderField" }

func (d CThostFtdcErrExecOrderField) String() string {
	var builder strings.Builder
	builder.Grow(521)

	builder.WriteString("CThostFtdcErrExecOrderField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", InvestorID=%+v", d.InvestorID)
	fmt.Fprintf(&builder, ", reserve1=%+v", d.reserve1)
	fmt.Fprintf(&builder, ", ExecOrderRef=%+v", d.ExecOrderRef)
	fmt.Fprintf(&builder, ", UserID=%+v", d.UserID)
	fmt.Fprintf(&builder, ", Volume=%+v", d.Volume)
	fmt.Fprintf(&builder, ", RequestID=%+v", d.RequestID)
	fmt.Fprintf(&builder, ", BusinessUnit=%+v", d.BusinessUnit)
	fmt.Fprintf(&builder, ", OffsetFlag=%+v", d.OffsetFlag)
	fmt.Fprintf(&builder, ", HedgeFlag=%+v", d.HedgeFlag)
	fmt.Fprintf(&builder, ", ActionType=%+v", d.ActionType)
	fmt.Fprintf(&builder, ", PosiDirection=%+v", d.PosiDirection)
	fmt.Fprintf(&builder, ", ReservePositionFlag=%+v", d.ReservePositionFlag)
	fmt.Fprintf(&builder, ", CloseFlag=%+v", d.CloseFlag)
	fmt.Fprintf(&builder, ", ExchangeID=%+v", d.ExchangeID)
	fmt.Fprintf(&builder, ", InvestUnitID=%+v", d.InvestUnitID)
	fmt.Fprintf(&builder, ", AccountID=%+v", d.AccountID)
	fmt.Fprintf(&builder, ", CurrencyID=%+v", d.CurrencyID)
	fmt.Fprintf(&builder, ", ClientID=%+v", d.ClientID)
	fmt.Fprintf(&builder, ", reserve2=%+v", d.reserve2)
	fmt.Fprintf(&builder, ", MacAddress=%+v", d.MacAddress)
	fmt.Fprintf(&builder, ", ErrorID=%+v", d.ErrorID)
	fmt.Fprintf(&builder, ", ErrorMsg=%+v", d.ErrorMsg)
	fmt.Fprintf(&builder, ", InstrumentID=%+v", d.InstrumentID)
	fmt.Fprintf(&builder, ", IPAddress=%+v", d.IPAddress)

	builder.WriteByte('}')

	return builder.String()
}

// 查询错误执行宣告
type CThostFtdcQryErrExecOrderField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID types.TThostFtdcInvestorIDType
}

func (d CThostFtdcQryErrExecOrderField) Type() string { return "CThostFtdcQryErrExecOrderField" }

func (d CThostFtdcQryErrExecOrderField) String() string {
	var builder strings.Builder
	builder.Grow(68)

	builder.WriteString("CThostFtdcQryErrExecOrderField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", InvestorID=%+v", d.InvestorID)

	builder.WriteByte('}')

	return builder.String()
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
	// 保留的无效字段
	reserve1 types.TThostFtdcOldInstrumentIDType
	// 投资单元代码
	InvestUnitID types.TThostFtdcInvestUnitIDType
	// 保留的无效字段
	reserve2 types.TThostFtdcOldIPAddressType
	// Mac地址
	MacAddress types.TThostFtdcMacAddressType
	// 错误代码
	ErrorID types.TThostFtdcErrorIDType
	// 错误信息
	ErrorMsg types.TThostFtdcErrorMsgType
	// 合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
	// IP地址
	IPAddress types.TThostFtdcIPAddressType
}

func (d CThostFtdcErrExecOrderActionField) Type() string { return "CThostFtdcErrExecOrderActionField" }

func (d CThostFtdcErrExecOrderActionField) String() string {
	var builder strings.Builder
	builder.Grow(410)

	builder.WriteString("CThostFtdcErrExecOrderActionField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", InvestorID=%+v", d.InvestorID)
	fmt.Fprintf(&builder, ", ExecOrderActionRef=%+v", d.ExecOrderActionRef)
	fmt.Fprintf(&builder, ", ExecOrderRef=%+v", d.ExecOrderRef)
	fmt.Fprintf(&builder, ", RequestID=%+v", d.RequestID)
	fmt.Fprintf(&builder, ", FrontID=%+v", d.FrontID)
	fmt.Fprintf(&builder, ", SessionID=%+v", d.SessionID)
	fmt.Fprintf(&builder, ", ExchangeID=%+v", d.ExchangeID)
	fmt.Fprintf(&builder, ", ExecOrderSysID=%+v", d.ExecOrderSysID)
	fmt.Fprintf(&builder, ", ActionFlag=%+v", d.ActionFlag)
	fmt.Fprintf(&builder, ", UserID=%+v", d.UserID)
	fmt.Fprintf(&builder, ", reserve1=%+v", d.reserve1)
	fmt.Fprintf(&builder, ", InvestUnitID=%+v", d.InvestUnitID)
	fmt.Fprintf(&builder, ", reserve2=%+v", d.reserve2)
	fmt.Fprintf(&builder, ", MacAddress=%+v", d.MacAddress)
	fmt.Fprintf(&builder, ", ErrorID=%+v", d.ErrorID)
	fmt.Fprintf(&builder, ", ErrorMsg=%+v", d.ErrorMsg)
	fmt.Fprintf(&builder, ", InstrumentID=%+v", d.InstrumentID)
	fmt.Fprintf(&builder, ", IPAddress=%+v", d.IPAddress)

	builder.WriteByte('}')

	return builder.String()
}

// 查询错误执行宣告操作
type CThostFtdcQryErrExecOrderActionField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID types.TThostFtdcInvestorIDType
}

func (d CThostFtdcQryErrExecOrderActionField) Type() string {
	return "CThostFtdcQryErrExecOrderActionField"
}

func (d CThostFtdcQryErrExecOrderActionField) String() string {
	var builder strings.Builder
	builder.Grow(74)

	builder.WriteString("CThostFtdcQryErrExecOrderActionField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", InvestorID=%+v", d.InvestorID)

	builder.WriteByte('}')

	return builder.String()
}

// 投资者期权合约交易权限
type CThostFtdcOptionInstrTradingRightField struct {
	// 保留的无效字段
	reserve1 types.TThostFtdcOldInstrumentIDType
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
	// 合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
}

func (d CThostFtdcOptionInstrTradingRightField) Type() string {
	return "CThostFtdcOptionInstrTradingRightField"
}

func (d CThostFtdcOptionInstrTradingRightField) String() string {
	var builder strings.Builder
	builder.Grow(180)

	builder.WriteString("CThostFtdcOptionInstrTradingRightField{")
	fmt.Fprintf(&builder, "reserve1=%+v", d.reserve1)
	fmt.Fprintf(&builder, ", InvestorRange=%+v", d.InvestorRange)
	fmt.Fprintf(&builder, ", BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", InvestorID=%+v", d.InvestorID)
	fmt.Fprintf(&builder, ", Direction=%+v", d.Direction)
	fmt.Fprintf(&builder, ", TradingRight=%+v", d.TradingRight)
	fmt.Fprintf(&builder, ", InstrumentID=%+v", d.InstrumentID)

	builder.WriteByte('}')

	return builder.String()
}

// 查询期权合约交易权限
type CThostFtdcQryOptionInstrTradingRightField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	// 保留的无效字段
	reserve1 types.TThostFtdcOldInstrumentIDType
	// 买卖方向
	Direction types.TThostFtdcDirectionType
	// 合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
}

func (d CThostFtdcQryOptionInstrTradingRightField) Type() string {
	return "CThostFtdcQryOptionInstrTradingRightField"
}

func (d CThostFtdcQryOptionInstrTradingRightField) String() string {
	var builder strings.Builder
	builder.Grow(138)

	builder.WriteString("CThostFtdcQryOptionInstrTradingRightField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", InvestorID=%+v", d.InvestorID)
	fmt.Fprintf(&builder, ", reserve1=%+v", d.reserve1)
	fmt.Fprintf(&builder, ", Direction=%+v", d.Direction)
	fmt.Fprintf(&builder, ", InstrumentID=%+v", d.InstrumentID)

	builder.WriteByte('}')

	return builder.String()
}

// 输入的询价
type CThostFtdcInputForQuoteField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	// 保留的无效字段
	reserve1 types.TThostFtdcOldInstrumentIDType
	// 询价引用
	ForQuoteRef types.TThostFtdcOrderRefType
	// 用户代码
	UserID types.TThostFtdcUserIDType
	// 交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	// 投资单元代码
	InvestUnitID types.TThostFtdcInvestUnitIDType
	// 保留的无效字段
	reserve2 types.TThostFtdcOldIPAddressType
	// Mac地址
	MacAddress types.TThostFtdcMacAddressType
	// 合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
	// IP地址
	IPAddress types.TThostFtdcIPAddressType
}

func (d CThostFtdcInputForQuoteField) Type() string { return "CThostFtdcInputForQuoteField" }

func (d CThostFtdcInputForQuoteField) String() string {
	var builder strings.Builder
	builder.Grow(242)

	builder.WriteString("CThostFtdcInputForQuoteField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", InvestorID=%+v", d.InvestorID)
	fmt.Fprintf(&builder, ", reserve1=%+v", d.reserve1)
	fmt.Fprintf(&builder, ", ForQuoteRef=%+v", d.ForQuoteRef)
	fmt.Fprintf(&builder, ", UserID=%+v", d.UserID)
	fmt.Fprintf(&builder, ", ExchangeID=%+v", d.ExchangeID)
	fmt.Fprintf(&builder, ", InvestUnitID=%+v", d.InvestUnitID)
	fmt.Fprintf(&builder, ", reserve2=%+v", d.reserve2)
	fmt.Fprintf(&builder, ", MacAddress=%+v", d.MacAddress)
	fmt.Fprintf(&builder, ", InstrumentID=%+v", d.InstrumentID)
	fmt.Fprintf(&builder, ", IPAddress=%+v", d.IPAddress)

	builder.WriteByte('}')

	return builder.String()
}

// 询价
type CThostFtdcForQuoteField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	// 保留的无效字段
	reserve1 types.TThostFtdcOldInstrumentIDType
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
	// 保留的无效字段
	reserve2 types.TThostFtdcOldExchangeInstIDType
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
	// 保留的无效字段
	reserve3 types.TThostFtdcOldIPAddressType
	// Mac地址
	MacAddress types.TThostFtdcMacAddressType
	// 合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
	// 合约在交易所的代码
	ExchangeInstID types.TThostFtdcExchangeInstIDType
	// IP地址
	IPAddress types.TThostFtdcIPAddressType
}

func (d CThostFtdcForQuoteField) Type() string { return "CThostFtdcForQuoteField" }

func (d CThostFtdcForQuoteField) String() string {
	var builder strings.Builder
	builder.Grow(549)

	builder.WriteString("CThostFtdcForQuoteField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", InvestorID=%+v", d.InvestorID)
	fmt.Fprintf(&builder, ", reserve1=%+v", d.reserve1)
	fmt.Fprintf(&builder, ", ForQuoteRef=%+v", d.ForQuoteRef)
	fmt.Fprintf(&builder, ", UserID=%+v", d.UserID)
	fmt.Fprintf(&builder, ", ForQuoteLocalID=%+v", d.ForQuoteLocalID)
	fmt.Fprintf(&builder, ", ExchangeID=%+v", d.ExchangeID)
	fmt.Fprintf(&builder, ", ParticipantID=%+v", d.ParticipantID)
	fmt.Fprintf(&builder, ", ClientID=%+v", d.ClientID)
	fmt.Fprintf(&builder, ", reserve2=%+v", d.reserve2)
	fmt.Fprintf(&builder, ", TraderID=%+v", d.TraderID)
	fmt.Fprintf(&builder, ", InstallID=%+v", d.InstallID)
	fmt.Fprintf(&builder, ", InsertDate=%+v", d.InsertDate)
	fmt.Fprintf(&builder, ", InsertTime=%+v", d.InsertTime)
	fmt.Fprintf(&builder, ", ForQuoteStatus=%+v", d.ForQuoteStatus)
	fmt.Fprintf(&builder, ", FrontID=%+v", d.FrontID)
	fmt.Fprintf(&builder, ", SessionID=%+v", d.SessionID)
	fmt.Fprintf(&builder, ", StatusMsg=%+v", d.StatusMsg)
	fmt.Fprintf(&builder, ", ActiveUserID=%+v", d.ActiveUserID)
	fmt.Fprintf(&builder, ", BrokerForQutoSeq=%+v", d.BrokerForQutoSeq)
	fmt.Fprintf(&builder, ", InvestUnitID=%+v", d.InvestUnitID)
	fmt.Fprintf(&builder, ", reserve3=%+v", d.reserve3)
	fmt.Fprintf(&builder, ", MacAddress=%+v", d.MacAddress)
	fmt.Fprintf(&builder, ", InstrumentID=%+v", d.InstrumentID)
	fmt.Fprintf(&builder, ", ExchangeInstID=%+v", d.ExchangeInstID)
	fmt.Fprintf(&builder, ", IPAddress=%+v", d.IPAddress)

	builder.WriteByte('}')

	return builder.String()
}

// 询价查询
type CThostFtdcQryForQuoteField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	// 保留的无效字段
	reserve1 types.TThostFtdcOldInstrumentIDType
	// 交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	// 开始时间
	InsertTimeStart types.TThostFtdcTimeType
	// 结束时间
	InsertTimeEnd types.TThostFtdcTimeType
	// 投资单元代码
	InvestUnitID types.TThostFtdcInvestUnitIDType
	// 合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
}

func (d CThostFtdcQryForQuoteField) Type() string { return "CThostFtdcQryForQuoteField" }

func (d CThostFtdcQryForQuoteField) String() string {
	var builder strings.Builder
	builder.Grow(194)

	builder.WriteString("CThostFtdcQryForQuoteField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", InvestorID=%+v", d.InvestorID)
	fmt.Fprintf(&builder, ", reserve1=%+v", d.reserve1)
	fmt.Fprintf(&builder, ", ExchangeID=%+v", d.ExchangeID)
	fmt.Fprintf(&builder, ", InsertTimeStart=%+v", d.InsertTimeStart)
	fmt.Fprintf(&builder, ", InsertTimeEnd=%+v", d.InsertTimeEnd)
	fmt.Fprintf(&builder, ", InvestUnitID=%+v", d.InvestUnitID)
	fmt.Fprintf(&builder, ", InstrumentID=%+v", d.InstrumentID)

	builder.WriteByte('}')

	return builder.String()
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
	// 保留的无效字段
	reserve1 types.TThostFtdcOldExchangeInstIDType
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
	// 保留的无效字段
	reserve2 types.TThostFtdcOldIPAddressType
	// Mac地址
	MacAddress types.TThostFtdcMacAddressType
	// 合约在交易所的代码
	ExchangeInstID types.TThostFtdcExchangeInstIDType
	// IP地址
	IPAddress types.TThostFtdcIPAddressType
}

func (d CThostFtdcExchangeForQuoteField) Type() string { return "CThostFtdcExchangeForQuoteField" }

func (d CThostFtdcExchangeForQuoteField) String() string {
	var builder strings.Builder
	builder.Grow(317)

	builder.WriteString("CThostFtdcExchangeForQuoteField{")
	fmt.Fprintf(&builder, "ForQuoteLocalID=%+v", d.ForQuoteLocalID)
	fmt.Fprintf(&builder, ", ExchangeID=%+v", d.ExchangeID)
	fmt.Fprintf(&builder, ", ParticipantID=%+v", d.ParticipantID)
	fmt.Fprintf(&builder, ", ClientID=%+v", d.ClientID)
	fmt.Fprintf(&builder, ", reserve1=%+v", d.reserve1)
	fmt.Fprintf(&builder, ", TraderID=%+v", d.TraderID)
	fmt.Fprintf(&builder, ", InstallID=%+v", d.InstallID)
	fmt.Fprintf(&builder, ", InsertDate=%+v", d.InsertDate)
	fmt.Fprintf(&builder, ", InsertTime=%+v", d.InsertTime)
	fmt.Fprintf(&builder, ", ForQuoteStatus=%+v", d.ForQuoteStatus)
	fmt.Fprintf(&builder, ", reserve2=%+v", d.reserve2)
	fmt.Fprintf(&builder, ", MacAddress=%+v", d.MacAddress)
	fmt.Fprintf(&builder, ", ExchangeInstID=%+v", d.ExchangeInstID)
	fmt.Fprintf(&builder, ", IPAddress=%+v", d.IPAddress)

	builder.WriteByte('}')

	return builder.String()
}

// 交易所询价查询
type CThostFtdcQryExchangeForQuoteField struct {
	// 会员代码
	ParticipantID types.TThostFtdcParticipantIDType
	// 客户代码
	ClientID types.TThostFtdcClientIDType
	// 保留的无效字段
	reserve1 types.TThostFtdcOldExchangeInstIDType
	// 交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	// 交易所交易员代码
	TraderID types.TThostFtdcTraderIDType
	// 合约在交易所的代码
	ExchangeInstID types.TThostFtdcExchangeInstIDType
}

func (d CThostFtdcQryExchangeForQuoteField) Type() string {
	return "CThostFtdcQryExchangeForQuoteField"
}

func (d CThostFtdcQryExchangeForQuoteField) String() string {
	var builder strings.Builder
	builder.Grow(155)

	builder.WriteString("CThostFtdcQryExchangeForQuoteField{")
	fmt.Fprintf(&builder, "ParticipantID=%+v", d.ParticipantID)
	fmt.Fprintf(&builder, ", ClientID=%+v", d.ClientID)
	fmt.Fprintf(&builder, ", reserve1=%+v", d.reserve1)
	fmt.Fprintf(&builder, ", ExchangeID=%+v", d.ExchangeID)
	fmt.Fprintf(&builder, ", TraderID=%+v", d.TraderID)
	fmt.Fprintf(&builder, ", ExchangeInstID=%+v", d.ExchangeInstID)

	builder.WriteByte('}')

	return builder.String()
}

// 输入的报价
type CThostFtdcInputQuoteField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	// 保留的无效字段
	reserve1 types.TThostFtdcOldInstrumentIDType
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
	// 保留的无效字段
	reserve2 types.TThostFtdcOldIPAddressType
	// Mac地址
	MacAddress types.TThostFtdcMacAddressType
	// 合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
	// IP地址
	IPAddress types.TThostFtdcIPAddressType
	// 被顶单编号
	ReplaceSysID types.TThostFtdcOrderSysIDType
	// 有效期类型
	TimeCondition types.TThostFtdcTimeConditionType
	// 报单回显字段
	OrderMemo types.TThostFtdcOrderMemoType
	// session上请求计数 api自动维护
	SessionReqSeq types.TThostFtdcSequenceNo12Type
}

func (d CThostFtdcInputQuoteField) Type() string { return "CThostFtdcInputQuoteField" }

func (d CThostFtdcInputQuoteField) String() string {
	var builder strings.Builder
	builder.Grow(611)

	builder.WriteString("CThostFtdcInputQuoteField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", InvestorID=%+v", d.InvestorID)
	fmt.Fprintf(&builder, ", reserve1=%+v", d.reserve1)
	fmt.Fprintf(&builder, ", QuoteRef=%+v", d.QuoteRef)
	fmt.Fprintf(&builder, ", UserID=%+v", d.UserID)
	fmt.Fprintf(&builder, ", AskPrice=%+v", d.AskPrice)
	fmt.Fprintf(&builder, ", BidPrice=%+v", d.BidPrice)
	fmt.Fprintf(&builder, ", AskVolume=%+v", d.AskVolume)
	fmt.Fprintf(&builder, ", BidVolume=%+v", d.BidVolume)
	fmt.Fprintf(&builder, ", RequestID=%+v", d.RequestID)
	fmt.Fprintf(&builder, ", BusinessUnit=%+v", d.BusinessUnit)
	fmt.Fprintf(&builder, ", AskOffsetFlag=%+v", d.AskOffsetFlag)
	fmt.Fprintf(&builder, ", BidOffsetFlag=%+v", d.BidOffsetFlag)
	fmt.Fprintf(&builder, ", AskHedgeFlag=%+v", d.AskHedgeFlag)
	fmt.Fprintf(&builder, ", BidHedgeFlag=%+v", d.BidHedgeFlag)
	fmt.Fprintf(&builder, ", AskOrderRef=%+v", d.AskOrderRef)
	fmt.Fprintf(&builder, ", BidOrderRef=%+v", d.BidOrderRef)
	fmt.Fprintf(&builder, ", ForQuoteSysID=%+v", d.ForQuoteSysID)
	fmt.Fprintf(&builder, ", ExchangeID=%+v", d.ExchangeID)
	fmt.Fprintf(&builder, ", InvestUnitID=%+v", d.InvestUnitID)
	fmt.Fprintf(&builder, ", ClientID=%+v", d.ClientID)
	fmt.Fprintf(&builder, ", reserve2=%+v", d.reserve2)
	fmt.Fprintf(&builder, ", MacAddress=%+v", d.MacAddress)
	fmt.Fprintf(&builder, ", InstrumentID=%+v", d.InstrumentID)
	fmt.Fprintf(&builder, ", IPAddress=%+v", d.IPAddress)
	fmt.Fprintf(&builder, ", ReplaceSysID=%+v", d.ReplaceSysID)
	fmt.Fprintf(&builder, ", TimeCondition=%+v", d.TimeCondition)
	fmt.Fprintf(&builder, ", OrderMemo=%+v", d.OrderMemo)
	fmt.Fprintf(&builder, ", SessionReqSeq=%+v", d.SessionReqSeq)

	builder.WriteByte('}')

	return builder.String()
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
	// 保留的无效字段
	reserve1 types.TThostFtdcOldInstrumentIDType
	// 投资单元代码
	InvestUnitID types.TThostFtdcInvestUnitIDType
	// 交易编码
	ClientID types.TThostFtdcClientIDType
	// 保留的无效字段
	reserve2 types.TThostFtdcOldIPAddressType
	// Mac地址
	MacAddress types.TThostFtdcMacAddressType
	// 合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
	// IP地址
	IPAddress types.TThostFtdcIPAddressType
	// 报单回显字段
	OrderMemo types.TThostFtdcOrderMemoType
	// session上请求计数 api自动维护
	SessionReqSeq types.TThostFtdcSequenceNo12Type
}

func (d CThostFtdcInputQuoteActionField) Type() string { return "CThostFtdcInputQuoteActionField" }

func (d CThostFtdcInputQuoteActionField) String() string {
	var builder strings.Builder
	builder.Grow(421)

	builder.WriteString("CThostFtdcInputQuoteActionField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", InvestorID=%+v", d.InvestorID)
	fmt.Fprintf(&builder, ", QuoteActionRef=%+v", d.QuoteActionRef)
	fmt.Fprintf(&builder, ", QuoteRef=%+v", d.QuoteRef)
	fmt.Fprintf(&builder, ", RequestID=%+v", d.RequestID)
	fmt.Fprintf(&builder, ", FrontID=%+v", d.FrontID)
	fmt.Fprintf(&builder, ", SessionID=%+v", d.SessionID)
	fmt.Fprintf(&builder, ", ExchangeID=%+v", d.ExchangeID)
	fmt.Fprintf(&builder, ", QuoteSysID=%+v", d.QuoteSysID)
	fmt.Fprintf(&builder, ", ActionFlag=%+v", d.ActionFlag)
	fmt.Fprintf(&builder, ", UserID=%+v", d.UserID)
	fmt.Fprintf(&builder, ", reserve1=%+v", d.reserve1)
	fmt.Fprintf(&builder, ", InvestUnitID=%+v", d.InvestUnitID)
	fmt.Fprintf(&builder, ", ClientID=%+v", d.ClientID)
	fmt.Fprintf(&builder, ", reserve2=%+v", d.reserve2)
	fmt.Fprintf(&builder, ", MacAddress=%+v", d.MacAddress)
	fmt.Fprintf(&builder, ", InstrumentID=%+v", d.InstrumentID)
	fmt.Fprintf(&builder, ", IPAddress=%+v", d.IPAddress)
	fmt.Fprintf(&builder, ", OrderMemo=%+v", d.OrderMemo)
	fmt.Fprintf(&builder, ", SessionReqSeq=%+v", d.SessionReqSeq)

	builder.WriteByte('}')

	return builder.String()
}

// 报价
type CThostFtdcQuoteField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	// 保留的无效字段
	reserve1 types.TThostFtdcOldInstrumentIDType
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
	// 保留的无效字段
	reserve2 types.TThostFtdcOldExchangeInstIDType
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
	// 保留的无效字段
	reserve3 types.TThostFtdcOldIPAddressType
	// Mac地址
	MacAddress types.TThostFtdcMacAddressType
	// 合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
	// 合约在交易所的代码
	ExchangeInstID types.TThostFtdcExchangeInstIDType
	// IP地址
	IPAddress types.TThostFtdcIPAddressType
	// 被顶单编号
	ReplaceSysID types.TThostFtdcOrderSysIDType
	// 有效期类型
	TimeCondition types.TThostFtdcTimeConditionType
	// 报单回显字段
	OrderMemo types.TThostFtdcOrderMemoType
	// session上请求计数 api自动维护
	SessionReqSeq types.TThostFtdcSequenceNo12Type
}

func (d CThostFtdcQuoteField) Type() string { return "CThostFtdcQuoteField" }

func (d CThostFtdcQuoteField) String() string {
	var builder strings.Builder
	builder.Grow(1197)

	builder.WriteString("CThostFtdcQuoteField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", InvestorID=%+v", d.InvestorID)
	fmt.Fprintf(&builder, ", reserve1=%+v", d.reserve1)
	fmt.Fprintf(&builder, ", QuoteRef=%+v", d.QuoteRef)
	fmt.Fprintf(&builder, ", UserID=%+v", d.UserID)
	fmt.Fprintf(&builder, ", AskPrice=%+v", d.AskPrice)
	fmt.Fprintf(&builder, ", BidPrice=%+v", d.BidPrice)
	fmt.Fprintf(&builder, ", AskVolume=%+v", d.AskVolume)
	fmt.Fprintf(&builder, ", BidVolume=%+v", d.BidVolume)
	fmt.Fprintf(&builder, ", RequestID=%+v", d.RequestID)
	fmt.Fprintf(&builder, ", BusinessUnit=%+v", d.BusinessUnit)
	fmt.Fprintf(&builder, ", AskOffsetFlag=%+v", d.AskOffsetFlag)
	fmt.Fprintf(&builder, ", BidOffsetFlag=%+v", d.BidOffsetFlag)
	fmt.Fprintf(&builder, ", AskHedgeFlag=%+v", d.AskHedgeFlag)
	fmt.Fprintf(&builder, ", BidHedgeFlag=%+v", d.BidHedgeFlag)
	fmt.Fprintf(&builder, ", QuoteLocalID=%+v", d.QuoteLocalID)
	fmt.Fprintf(&builder, ", ExchangeID=%+v", d.ExchangeID)
	fmt.Fprintf(&builder, ", ParticipantID=%+v", d.ParticipantID)
	fmt.Fprintf(&builder, ", ClientID=%+v", d.ClientID)
	fmt.Fprintf(&builder, ", reserve2=%+v", d.reserve2)
	fmt.Fprintf(&builder, ", TraderID=%+v", d.TraderID)
	fmt.Fprintf(&builder, ", InstallID=%+v", d.InstallID)
	fmt.Fprintf(&builder, ", NotifySequence=%+v", d.NotifySequence)
	fmt.Fprintf(&builder, ", OrderSubmitStatus=%+v", d.OrderSubmitStatus)
	fmt.Fprintf(&builder, ", TradingDay=%+v", d.TradingDay)
	fmt.Fprintf(&builder, ", SettlementID=%+v", d.SettlementID)
	fmt.Fprintf(&builder, ", QuoteSysID=%+v", d.QuoteSysID)
	fmt.Fprintf(&builder, ", InsertDate=%+v", d.InsertDate)
	fmt.Fprintf(&builder, ", InsertTime=%+v", d.InsertTime)
	fmt.Fprintf(&builder, ", CancelTime=%+v", d.CancelTime)
	fmt.Fprintf(&builder, ", QuoteStatus=%+v", d.QuoteStatus)
	fmt.Fprintf(&builder, ", ClearingPartID=%+v", d.ClearingPartID)
	fmt.Fprintf(&builder, ", SequenceNo=%+v", d.SequenceNo)
	fmt.Fprintf(&builder, ", AskOrderSysID=%+v", d.AskOrderSysID)
	fmt.Fprintf(&builder, ", BidOrderSysID=%+v", d.BidOrderSysID)
	fmt.Fprintf(&builder, ", FrontID=%+v", d.FrontID)
	fmt.Fprintf(&builder, ", SessionID=%+v", d.SessionID)
	fmt.Fprintf(&builder, ", UserProductInfo=%+v", d.UserProductInfo)
	fmt.Fprintf(&builder, ", StatusMsg=%+v", d.StatusMsg)
	fmt.Fprintf(&builder, ", ActiveUserID=%+v", d.ActiveUserID)
	fmt.Fprintf(&builder, ", BrokerQuoteSeq=%+v", d.BrokerQuoteSeq)
	fmt.Fprintf(&builder, ", AskOrderRef=%+v", d.AskOrderRef)
	fmt.Fprintf(&builder, ", BidOrderRef=%+v", d.BidOrderRef)
	fmt.Fprintf(&builder, ", ForQuoteSysID=%+v", d.ForQuoteSysID)
	fmt.Fprintf(&builder, ", BranchID=%+v", d.BranchID)
	fmt.Fprintf(&builder, ", InvestUnitID=%+v", d.InvestUnitID)
	fmt.Fprintf(&builder, ", AccountID=%+v", d.AccountID)
	fmt.Fprintf(&builder, ", CurrencyID=%+v", d.CurrencyID)
	fmt.Fprintf(&builder, ", reserve3=%+v", d.reserve3)
	fmt.Fprintf(&builder, ", MacAddress=%+v", d.MacAddress)
	fmt.Fprintf(&builder, ", InstrumentID=%+v", d.InstrumentID)
	fmt.Fprintf(&builder, ", ExchangeInstID=%+v", d.ExchangeInstID)
	fmt.Fprintf(&builder, ", IPAddress=%+v", d.IPAddress)
	fmt.Fprintf(&builder, ", ReplaceSysID=%+v", d.ReplaceSysID)
	fmt.Fprintf(&builder, ", TimeCondition=%+v", d.TimeCondition)
	fmt.Fprintf(&builder, ", OrderMemo=%+v", d.OrderMemo)
	fmt.Fprintf(&builder, ", SessionReqSeq=%+v", d.SessionReqSeq)

	builder.WriteByte('}')

	return builder.String()
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
	// 保留的无效字段
	reserve1 types.TThostFtdcOldInstrumentIDType
	// 营业部编号
	BranchID types.TThostFtdcBranchIDType
	// 投资单元代码
	InvestUnitID types.TThostFtdcInvestUnitIDType
	// 保留的无效字段
	reserve2 types.TThostFtdcOldIPAddressType
	// Mac地址
	MacAddress types.TThostFtdcMacAddressType
	// 合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
	// IP地址
	IPAddress types.TThostFtdcIPAddressType
	// 报单回显字段
	OrderMemo types.TThostFtdcOrderMemoType
	// session上请求计数 api自动维护
	SessionReqSeq types.TThostFtdcSequenceNo12Type
}

func (d CThostFtdcQuoteActionField) Type() string { return "CThostFtdcQuoteActionField" }

func (d CThostFtdcQuoteActionField) String() string {
	var builder strings.Builder
	builder.Grow(647)

	builder.WriteString("CThostFtdcQuoteActionField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", InvestorID=%+v", d.InvestorID)
	fmt.Fprintf(&builder, ", QuoteActionRef=%+v", d.QuoteActionRef)
	fmt.Fprintf(&builder, ", QuoteRef=%+v", d.QuoteRef)
	fmt.Fprintf(&builder, ", RequestID=%+v", d.RequestID)
	fmt.Fprintf(&builder, ", FrontID=%+v", d.FrontID)
	fmt.Fprintf(&builder, ", SessionID=%+v", d.SessionID)
	fmt.Fprintf(&builder, ", ExchangeID=%+v", d.ExchangeID)
	fmt.Fprintf(&builder, ", QuoteSysID=%+v", d.QuoteSysID)
	fmt.Fprintf(&builder, ", ActionFlag=%+v", d.ActionFlag)
	fmt.Fprintf(&builder, ", ActionDate=%+v", d.ActionDate)
	fmt.Fprintf(&builder, ", ActionTime=%+v", d.ActionTime)
	fmt.Fprintf(&builder, ", TraderID=%+v", d.TraderID)
	fmt.Fprintf(&builder, ", InstallID=%+v", d.InstallID)
	fmt.Fprintf(&builder, ", QuoteLocalID=%+v", d.QuoteLocalID)
	fmt.Fprintf(&builder, ", ActionLocalID=%+v", d.ActionLocalID)
	fmt.Fprintf(&builder, ", ParticipantID=%+v", d.ParticipantID)
	fmt.Fprintf(&builder, ", ClientID=%+v", d.ClientID)
	fmt.Fprintf(&builder, ", BusinessUnit=%+v", d.BusinessUnit)
	fmt.Fprintf(&builder, ", OrderActionStatus=%+v", d.OrderActionStatus)
	fmt.Fprintf(&builder, ", UserID=%+v", d.UserID)
	fmt.Fprintf(&builder, ", StatusMsg=%+v", d.StatusMsg)
	fmt.Fprintf(&builder, ", reserve1=%+v", d.reserve1)
	fmt.Fprintf(&builder, ", BranchID=%+v", d.BranchID)
	fmt.Fprintf(&builder, ", InvestUnitID=%+v", d.InvestUnitID)
	fmt.Fprintf(&builder, ", reserve2=%+v", d.reserve2)
	fmt.Fprintf(&builder, ", MacAddress=%+v", d.MacAddress)
	fmt.Fprintf(&builder, ", InstrumentID=%+v", d.InstrumentID)
	fmt.Fprintf(&builder, ", IPAddress=%+v", d.IPAddress)
	fmt.Fprintf(&builder, ", OrderMemo=%+v", d.OrderMemo)
	fmt.Fprintf(&builder, ", SessionReqSeq=%+v", d.SessionReqSeq)

	builder.WriteByte('}')

	return builder.String()
}

// 报价查询
type CThostFtdcQryQuoteField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	// 保留的无效字段
	reserve1 types.TThostFtdcOldInstrumentIDType
	// 交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	// 报价编号
	QuoteSysID types.TThostFtdcOrderSysIDType
	// 开始时间
	InsertTimeStart types.TThostFtdcTimeType
	// 结束时间
	InsertTimeEnd types.TThostFtdcTimeType
	// 投资单元代码
	InvestUnitID types.TThostFtdcInvestUnitIDType
	// 合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
}

func (d CThostFtdcQryQuoteField) Type() string { return "CThostFtdcQryQuoteField" }

func (d CThostFtdcQryQuoteField) String() string {
	var builder strings.Builder
	builder.Grow(211)

	builder.WriteString("CThostFtdcQryQuoteField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", InvestorID=%+v", d.InvestorID)
	fmt.Fprintf(&builder, ", reserve1=%+v", d.reserve1)
	fmt.Fprintf(&builder, ", ExchangeID=%+v", d.ExchangeID)
	fmt.Fprintf(&builder, ", QuoteSysID=%+v", d.QuoteSysID)
	fmt.Fprintf(&builder, ", InsertTimeStart=%+v", d.InsertTimeStart)
	fmt.Fprintf(&builder, ", InsertTimeEnd=%+v", d.InsertTimeEnd)
	fmt.Fprintf(&builder, ", InvestUnitID=%+v", d.InvestUnitID)
	fmt.Fprintf(&builder, ", InstrumentID=%+v", d.InstrumentID)

	builder.WriteByte('}')

	return builder.String()
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
	// 保留的无效字段
	reserve1 types.TThostFtdcOldExchangeInstIDType
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
	// 保留的无效字段
	reserve2 types.TThostFtdcOldIPAddressType
	// Mac地址
	MacAddress types.TThostFtdcMacAddressType
	// 合约在交易所的代码
	ExchangeInstID types.TThostFtdcExchangeInstIDType
	// IP地址
	IPAddress types.TThostFtdcIPAddressType
	// 有效期类型
	TimeCondition types.TThostFtdcTimeConditionType
}

func (d CThostFtdcExchangeQuoteField) Type() string { return "CThostFtdcExchangeQuoteField" }

func (d CThostFtdcExchangeQuoteField) String() string {
	var builder strings.Builder
	builder.Grow(800)

	builder.WriteString("CThostFtdcExchangeQuoteField{")
	fmt.Fprintf(&builder, "AskPrice=%+v", d.AskPrice)
	fmt.Fprintf(&builder, ", BidPrice=%+v", d.BidPrice)
	fmt.Fprintf(&builder, ", AskVolume=%+v", d.AskVolume)
	fmt.Fprintf(&builder, ", BidVolume=%+v", d.BidVolume)
	fmt.Fprintf(&builder, ", RequestID=%+v", d.RequestID)
	fmt.Fprintf(&builder, ", BusinessUnit=%+v", d.BusinessUnit)
	fmt.Fprintf(&builder, ", AskOffsetFlag=%+v", d.AskOffsetFlag)
	fmt.Fprintf(&builder, ", BidOffsetFlag=%+v", d.BidOffsetFlag)
	fmt.Fprintf(&builder, ", AskHedgeFlag=%+v", d.AskHedgeFlag)
	fmt.Fprintf(&builder, ", BidHedgeFlag=%+v", d.BidHedgeFlag)
	fmt.Fprintf(&builder, ", QuoteLocalID=%+v", d.QuoteLocalID)
	fmt.Fprintf(&builder, ", ExchangeID=%+v", d.ExchangeID)
	fmt.Fprintf(&builder, ", ParticipantID=%+v", d.ParticipantID)
	fmt.Fprintf(&builder, ", ClientID=%+v", d.ClientID)
	fmt.Fprintf(&builder, ", reserve1=%+v", d.reserve1)
	fmt.Fprintf(&builder, ", TraderID=%+v", d.TraderID)
	fmt.Fprintf(&builder, ", InstallID=%+v", d.InstallID)
	fmt.Fprintf(&builder, ", NotifySequence=%+v", d.NotifySequence)
	fmt.Fprintf(&builder, ", OrderSubmitStatus=%+v", d.OrderSubmitStatus)
	fmt.Fprintf(&builder, ", TradingDay=%+v", d.TradingDay)
	fmt.Fprintf(&builder, ", SettlementID=%+v", d.SettlementID)
	fmt.Fprintf(&builder, ", QuoteSysID=%+v", d.QuoteSysID)
	fmt.Fprintf(&builder, ", InsertDate=%+v", d.InsertDate)
	fmt.Fprintf(&builder, ", InsertTime=%+v", d.InsertTime)
	fmt.Fprintf(&builder, ", CancelTime=%+v", d.CancelTime)
	fmt.Fprintf(&builder, ", QuoteStatus=%+v", d.QuoteStatus)
	fmt.Fprintf(&builder, ", ClearingPartID=%+v", d.ClearingPartID)
	fmt.Fprintf(&builder, ", SequenceNo=%+v", d.SequenceNo)
	fmt.Fprintf(&builder, ", AskOrderSysID=%+v", d.AskOrderSysID)
	fmt.Fprintf(&builder, ", BidOrderSysID=%+v", d.BidOrderSysID)
	fmt.Fprintf(&builder, ", ForQuoteSysID=%+v", d.ForQuoteSysID)
	fmt.Fprintf(&builder, ", BranchID=%+v", d.BranchID)
	fmt.Fprintf(&builder, ", reserve2=%+v", d.reserve2)
	fmt.Fprintf(&builder, ", MacAddress=%+v", d.MacAddress)
	fmt.Fprintf(&builder, ", ExchangeInstID=%+v", d.ExchangeInstID)
	fmt.Fprintf(&builder, ", IPAddress=%+v", d.IPAddress)
	fmt.Fprintf(&builder, ", TimeCondition=%+v", d.TimeCondition)

	builder.WriteByte('}')

	return builder.String()
}

// 交易所报价查询
type CThostFtdcQryExchangeQuoteField struct {
	// 会员代码
	ParticipantID types.TThostFtdcParticipantIDType
	// 客户代码
	ClientID types.TThostFtdcClientIDType
	// 保留的无效字段
	reserve1 types.TThostFtdcOldExchangeInstIDType
	// 交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	// 交易所交易员代码
	TraderID types.TThostFtdcTraderIDType
	// 合约在交易所的代码
	ExchangeInstID types.TThostFtdcExchangeInstIDType
}

func (d CThostFtdcQryExchangeQuoteField) Type() string { return "CThostFtdcQryExchangeQuoteField" }

func (d CThostFtdcQryExchangeQuoteField) String() string {
	var builder strings.Builder
	builder.Grow(152)

	builder.WriteString("CThostFtdcQryExchangeQuoteField{")
	fmt.Fprintf(&builder, "ParticipantID=%+v", d.ParticipantID)
	fmt.Fprintf(&builder, ", ClientID=%+v", d.ClientID)
	fmt.Fprintf(&builder, ", reserve1=%+v", d.reserve1)
	fmt.Fprintf(&builder, ", ExchangeID=%+v", d.ExchangeID)
	fmt.Fprintf(&builder, ", TraderID=%+v", d.TraderID)
	fmt.Fprintf(&builder, ", ExchangeInstID=%+v", d.ExchangeInstID)

	builder.WriteByte('}')

	return builder.String()
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

func (d CThostFtdcQryQuoteActionField) Type() string { return "CThostFtdcQryQuoteActionField" }

func (d CThostFtdcQryQuoteActionField) String() string {
	var builder strings.Builder
	builder.Grow(87)

	builder.WriteString("CThostFtdcQryQuoteActionField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", InvestorID=%+v", d.InvestorID)
	fmt.Fprintf(&builder, ", ExchangeID=%+v", d.ExchangeID)

	builder.WriteByte('}')

	return builder.String()
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
	// 保留的无效字段
	reserve1 types.TThostFtdcOldIPAddressType
	// Mac地址
	MacAddress types.TThostFtdcMacAddressType
	// IP地址
	IPAddress types.TThostFtdcIPAddressType
}

func (d CThostFtdcExchangeQuoteActionField) Type() string {
	return "CThostFtdcExchangeQuoteActionField"
}

func (d CThostFtdcExchangeQuoteActionField) String() string {
	var builder strings.Builder
	builder.Grow(379)

	builder.WriteString("CThostFtdcExchangeQuoteActionField{")
	fmt.Fprintf(&builder, "ExchangeID=%+v", d.ExchangeID)
	fmt.Fprintf(&builder, ", QuoteSysID=%+v", d.QuoteSysID)
	fmt.Fprintf(&builder, ", ActionFlag=%+v", d.ActionFlag)
	fmt.Fprintf(&builder, ", ActionDate=%+v", d.ActionDate)
	fmt.Fprintf(&builder, ", ActionTime=%+v", d.ActionTime)
	fmt.Fprintf(&builder, ", TraderID=%+v", d.TraderID)
	fmt.Fprintf(&builder, ", InstallID=%+v", d.InstallID)
	fmt.Fprintf(&builder, ", QuoteLocalID=%+v", d.QuoteLocalID)
	fmt.Fprintf(&builder, ", ActionLocalID=%+v", d.ActionLocalID)
	fmt.Fprintf(&builder, ", ParticipantID=%+v", d.ParticipantID)
	fmt.Fprintf(&builder, ", ClientID=%+v", d.ClientID)
	fmt.Fprintf(&builder, ", BusinessUnit=%+v", d.BusinessUnit)
	fmt.Fprintf(&builder, ", OrderActionStatus=%+v", d.OrderActionStatus)
	fmt.Fprintf(&builder, ", UserID=%+v", d.UserID)
	fmt.Fprintf(&builder, ", reserve1=%+v", d.reserve1)
	fmt.Fprintf(&builder, ", MacAddress=%+v", d.MacAddress)
	fmt.Fprintf(&builder, ", IPAddress=%+v", d.IPAddress)

	builder.WriteByte('}')

	return builder.String()
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

func (d CThostFtdcQryExchangeQuoteActionField) Type() string {
	return "CThostFtdcQryExchangeQuoteActionField"
}

func (d CThostFtdcQryExchangeQuoteActionField) String() string {
	var builder strings.Builder
	builder.Grow(116)

	builder.WriteString("CThostFtdcQryExchangeQuoteActionField{")
	fmt.Fprintf(&builder, "ParticipantID=%+v", d.ParticipantID)
	fmt.Fprintf(&builder, ", ClientID=%+v", d.ClientID)
	fmt.Fprintf(&builder, ", ExchangeID=%+v", d.ExchangeID)
	fmt.Fprintf(&builder, ", TraderID=%+v", d.TraderID)

	builder.WriteByte('}')

	return builder.String()
}

// 期权合约delta值
type CThostFtdcOptionInstrDeltaField struct {
	// 保留的无效字段
	reserve1 types.TThostFtdcOldInstrumentIDType
	// 投资者范围
	InvestorRange types.TThostFtdcInvestorRangeType
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	// Delta值
	Delta types.TThostFtdcRatioType
	// 合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
}

func (d CThostFtdcOptionInstrDeltaField) Type() string { return "CThostFtdcOptionInstrDeltaField" }

func (d CThostFtdcOptionInstrDeltaField) String() string {
	var builder strings.Builder
	builder.Grow(147)

	builder.WriteString("CThostFtdcOptionInstrDeltaField{")
	fmt.Fprintf(&builder, "reserve1=%+v", d.reserve1)
	fmt.Fprintf(&builder, ", InvestorRange=%+v", d.InvestorRange)
	fmt.Fprintf(&builder, ", BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", InvestorID=%+v", d.InvestorID)
	fmt.Fprintf(&builder, ", Delta=%+v", d.Delta)
	fmt.Fprintf(&builder, ", InstrumentID=%+v", d.InstrumentID)

	builder.WriteByte('}')

	return builder.String()
}

// 发给做市商的询价请求
type CThostFtdcForQuoteRspField struct {
	// 交易日
	TradingDay types.TThostFtdcDateType
	// 保留的无效字段
	reserve1 types.TThostFtdcOldInstrumentIDType
	// 询价编号
	ForQuoteSysID types.TThostFtdcOrderSysIDType
	// 询价时间
	ForQuoteTime types.TThostFtdcTimeType
	// 业务日期
	ActionDay types.TThostFtdcDateType
	// 交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	// 合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
}

func (d CThostFtdcForQuoteRspField) Type() string { return "CThostFtdcForQuoteRspField" }

func (d CThostFtdcForQuoteRspField) String() string {
	var builder strings.Builder
	builder.Grow(170)

	builder.WriteString("CThostFtdcForQuoteRspField{")
	fmt.Fprintf(&builder, "TradingDay=%+v", d.TradingDay)
	fmt.Fprintf(&builder, ", reserve1=%+v", d.reserve1)
	fmt.Fprintf(&builder, ", ForQuoteSysID=%+v", d.ForQuoteSysID)
	fmt.Fprintf(&builder, ", ForQuoteTime=%+v", d.ForQuoteTime)
	fmt.Fprintf(&builder, ", ActionDay=%+v", d.ActionDay)
	fmt.Fprintf(&builder, ", ExchangeID=%+v", d.ExchangeID)
	fmt.Fprintf(&builder, ", InstrumentID=%+v", d.InstrumentID)

	builder.WriteByte('}')

	return builder.String()
}

// 当前期权合约执行偏移值的详细内容
type CThostFtdcStrikeOffsetField struct {
	// 保留的无效字段
	reserve1 types.TThostFtdcOldInstrumentIDType
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
	// 合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
}

func (d CThostFtdcStrikeOffsetField) Type() string { return "CThostFtdcStrikeOffsetField" }

func (d CThostFtdcStrikeOffsetField) String() string {
	var builder strings.Builder
	builder.Grow(164)

	builder.WriteString("CThostFtdcStrikeOffsetField{")
	fmt.Fprintf(&builder, "reserve1=%+v", d.reserve1)
	fmt.Fprintf(&builder, ", InvestorRange=%+v", d.InvestorRange)
	fmt.Fprintf(&builder, ", BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", InvestorID=%+v", d.InvestorID)
	fmt.Fprintf(&builder, ", Offset=%+v", d.Offset)
	fmt.Fprintf(&builder, ", OffsetType=%+v", d.OffsetType)
	fmt.Fprintf(&builder, ", InstrumentID=%+v", d.InstrumentID)

	builder.WriteByte('}')

	return builder.String()
}

// 期权执行偏移值查询
type CThostFtdcQryStrikeOffsetField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	// 保留的无效字段
	reserve1 types.TThostFtdcOldInstrumentIDType
	// 合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
}

func (d CThostFtdcQryStrikeOffsetField) Type() string { return "CThostFtdcQryStrikeOffsetField" }

func (d CThostFtdcQryStrikeOffsetField) String() string {
	var builder strings.Builder
	builder.Grow(108)

	builder.WriteString("CThostFtdcQryStrikeOffsetField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", InvestorID=%+v", d.InvestorID)
	fmt.Fprintf(&builder, ", reserve1=%+v", d.reserve1)
	fmt.Fprintf(&builder, ", InstrumentID=%+v", d.InstrumentID)

	builder.WriteByte('}')

	return builder.String()
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
	// 保留的无效字段
	reserve1 types.TThostFtdcOldIPAddressType
	// Mac地址
	MacAddress types.TThostFtdcMacAddressType
	// IP地址
	IPAddress types.TThostFtdcIPAddressType
}

func (d CThostFtdcInputBatchOrderActionField) Type() string {
	return "CThostFtdcInputBatchOrderActionField"
}

func (d CThostFtdcInputBatchOrderActionField) String() string {
	var builder strings.Builder
	builder.Grow(268)

	builder.WriteString("CThostFtdcInputBatchOrderActionField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", InvestorID=%+v", d.InvestorID)
	fmt.Fprintf(&builder, ", OrderActionRef=%+v", d.OrderActionRef)
	fmt.Fprintf(&builder, ", RequestID=%+v", d.RequestID)
	fmt.Fprintf(&builder, ", FrontID=%+v", d.FrontID)
	fmt.Fprintf(&builder, ", SessionID=%+v", d.SessionID)
	fmt.Fprintf(&builder, ", ExchangeID=%+v", d.ExchangeID)
	fmt.Fprintf(&builder, ", UserID=%+v", d.UserID)
	fmt.Fprintf(&builder, ", InvestUnitID=%+v", d.InvestUnitID)
	fmt.Fprintf(&builder, ", reserve1=%+v", d.reserve1)
	fmt.Fprintf(&builder, ", MacAddress=%+v", d.MacAddress)
	fmt.Fprintf(&builder, ", IPAddress=%+v", d.IPAddress)

	builder.WriteByte('}')

	return builder.String()
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
	// 保留的无效字段
	reserve1 types.TThostFtdcOldIPAddressType
	// Mac地址
	MacAddress types.TThostFtdcMacAddressType
	// IP地址
	IPAddress types.TThostFtdcIPAddressType
}

func (d CThostFtdcBatchOrderActionField) Type() string { return "CThostFtdcBatchOrderActionField" }

func (d CThostFtdcBatchOrderActionField) String() string {
	var builder strings.Builder
	builder.Grow(472)

	builder.WriteString("CThostFtdcBatchOrderActionField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", InvestorID=%+v", d.InvestorID)
	fmt.Fprintf(&builder, ", OrderActionRef=%+v", d.OrderActionRef)
	fmt.Fprintf(&builder, ", RequestID=%+v", d.RequestID)
	fmt.Fprintf(&builder, ", FrontID=%+v", d.FrontID)
	fmt.Fprintf(&builder, ", SessionID=%+v", d.SessionID)
	fmt.Fprintf(&builder, ", ExchangeID=%+v", d.ExchangeID)
	fmt.Fprintf(&builder, ", ActionDate=%+v", d.ActionDate)
	fmt.Fprintf(&builder, ", ActionTime=%+v", d.ActionTime)
	fmt.Fprintf(&builder, ", TraderID=%+v", d.TraderID)
	fmt.Fprintf(&builder, ", InstallID=%+v", d.InstallID)
	fmt.Fprintf(&builder, ", ActionLocalID=%+v", d.ActionLocalID)
	fmt.Fprintf(&builder, ", ParticipantID=%+v", d.ParticipantID)
	fmt.Fprintf(&builder, ", ClientID=%+v", d.ClientID)
	fmt.Fprintf(&builder, ", BusinessUnit=%+v", d.BusinessUnit)
	fmt.Fprintf(&builder, ", OrderActionStatus=%+v", d.OrderActionStatus)
	fmt.Fprintf(&builder, ", UserID=%+v", d.UserID)
	fmt.Fprintf(&builder, ", StatusMsg=%+v", d.StatusMsg)
	fmt.Fprintf(&builder, ", InvestUnitID=%+v", d.InvestUnitID)
	fmt.Fprintf(&builder, ", reserve1=%+v", d.reserve1)
	fmt.Fprintf(&builder, ", MacAddress=%+v", d.MacAddress)
	fmt.Fprintf(&builder, ", IPAddress=%+v", d.IPAddress)

	builder.WriteByte('}')

	return builder.String()
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
	// 保留的无效字段
	reserve1 types.TThostFtdcOldIPAddressType
	// Mac地址
	MacAddress types.TThostFtdcMacAddressType
	// IP地址
	IPAddress types.TThostFtdcIPAddressType
}

func (d CThostFtdcExchangeBatchOrderActionField) Type() string {
	return "CThostFtdcExchangeBatchOrderActionField"
}

func (d CThostFtdcExchangeBatchOrderActionField) String() string {
	var builder strings.Builder
	builder.Grow(322)

	builder.WriteString("CThostFtdcExchangeBatchOrderActionField{")
	fmt.Fprintf(&builder, "ExchangeID=%+v", d.ExchangeID)
	fmt.Fprintf(&builder, ", ActionDate=%+v", d.ActionDate)
	fmt.Fprintf(&builder, ", ActionTime=%+v", d.ActionTime)
	fmt.Fprintf(&builder, ", TraderID=%+v", d.TraderID)
	fmt.Fprintf(&builder, ", InstallID=%+v", d.InstallID)
	fmt.Fprintf(&builder, ", ActionLocalID=%+v", d.ActionLocalID)
	fmt.Fprintf(&builder, ", ParticipantID=%+v", d.ParticipantID)
	fmt.Fprintf(&builder, ", ClientID=%+v", d.ClientID)
	fmt.Fprintf(&builder, ", BusinessUnit=%+v", d.BusinessUnit)
	fmt.Fprintf(&builder, ", OrderActionStatus=%+v", d.OrderActionStatus)
	fmt.Fprintf(&builder, ", UserID=%+v", d.UserID)
	fmt.Fprintf(&builder, ", reserve1=%+v", d.reserve1)
	fmt.Fprintf(&builder, ", MacAddress=%+v", d.MacAddress)
	fmt.Fprintf(&builder, ", IPAddress=%+v", d.IPAddress)

	builder.WriteByte('}')

	return builder.String()
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

func (d CThostFtdcQryBatchOrderActionField) Type() string {
	return "CThostFtdcQryBatchOrderActionField"
}

func (d CThostFtdcQryBatchOrderActionField) String() string {
	var builder strings.Builder
	builder.Grow(92)

	builder.WriteString("CThostFtdcQryBatchOrderActionField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", InvestorID=%+v", d.InvestorID)
	fmt.Fprintf(&builder, ", ExchangeID=%+v", d.ExchangeID)

	builder.WriteByte('}')

	return builder.String()
}

// 组合合约安全系数
type CThostFtdcCombInstrumentGuardField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 保留的无效字段
	reserve1     types.TThostFtdcOldInstrumentIDType
	GuarantRatio types.TThostFtdcRatioType
	// 交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	// 合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
}

func (d CThostFtdcCombInstrumentGuardField) Type() string {
	return "CThostFtdcCombInstrumentGuardField"
}

func (d CThostFtdcCombInstrumentGuardField) String() string {
	var builder strings.Builder
	builder.Grow(134)

	builder.WriteString("CThostFtdcCombInstrumentGuardField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", reserve1=%+v", d.reserve1)
	fmt.Fprintf(&builder, ", GuarantRatio=%+v", d.GuarantRatio)
	fmt.Fprintf(&builder, ", ExchangeID=%+v", d.ExchangeID)
	fmt.Fprintf(&builder, ", InstrumentID=%+v", d.InstrumentID)

	builder.WriteByte('}')

	return builder.String()
}

// 组合合约安全系数查询
type CThostFtdcQryCombInstrumentGuardField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 保留的无效字段
	reserve1 types.TThostFtdcOldInstrumentIDType
	// 交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	// 合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
}

func (d CThostFtdcQryCombInstrumentGuardField) Type() string {
	return "CThostFtdcQryCombInstrumentGuardField"
}

func (d CThostFtdcQryCombInstrumentGuardField) String() string {
	var builder strings.Builder
	builder.Grow(115)

	builder.WriteString("CThostFtdcQryCombInstrumentGuardField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", reserve1=%+v", d.reserve1)
	fmt.Fprintf(&builder, ", ExchangeID=%+v", d.ExchangeID)
	fmt.Fprintf(&builder, ", InstrumentID=%+v", d.InstrumentID)

	builder.WriteByte('}')

	return builder.String()
}

// 输入的申请组合
type CThostFtdcInputCombActionField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	// 保留的无效字段
	reserve1 types.TThostFtdcOldInstrumentIDType
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
	// 交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	// 保留的无效字段
	reserve2 types.TThostFtdcOldIPAddressType
	// Mac地址
	MacAddress types.TThostFtdcMacAddressType
	// 投资单元代码
	InvestUnitID types.TThostFtdcInvestUnitIDType
	// 前置编号
	FrontID types.TThostFtdcFrontIDType
	// 会话编号
	SessionID types.TThostFtdcSessionIDType
	// 合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
	// IP地址
	IPAddress types.TThostFtdcIPAddressType
}

func (d CThostFtdcInputCombActionField) Type() string { return "CThostFtdcInputCombActionField" }

func (d CThostFtdcInputCombActionField) String() string {
	var builder strings.Builder
	builder.Grow(359)

	builder.WriteString("CThostFtdcInputCombActionField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", InvestorID=%+v", d.InvestorID)
	fmt.Fprintf(&builder, ", reserve1=%+v", d.reserve1)
	fmt.Fprintf(&builder, ", CombActionRef=%+v", d.CombActionRef)
	fmt.Fprintf(&builder, ", UserID=%+v", d.UserID)
	fmt.Fprintf(&builder, ", Direction=%+v", d.Direction)
	fmt.Fprintf(&builder, ", Volume=%+v", d.Volume)
	fmt.Fprintf(&builder, ", CombDirection=%+v", d.CombDirection)
	fmt.Fprintf(&builder, ", HedgeFlag=%+v", d.HedgeFlag)
	fmt.Fprintf(&builder, ", ExchangeID=%+v", d.ExchangeID)
	fmt.Fprintf(&builder, ", reserve2=%+v", d.reserve2)
	fmt.Fprintf(&builder, ", MacAddress=%+v", d.MacAddress)
	fmt.Fprintf(&builder, ", InvestUnitID=%+v", d.InvestUnitID)
	fmt.Fprintf(&builder, ", FrontID=%+v", d.FrontID)
	fmt.Fprintf(&builder, ", SessionID=%+v", d.SessionID)
	fmt.Fprintf(&builder, ", InstrumentID=%+v", d.InstrumentID)
	fmt.Fprintf(&builder, ", IPAddress=%+v", d.IPAddress)

	builder.WriteByte('}')

	return builder.String()
}

// 申请组合
type CThostFtdcCombActionField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	// 保留的无效字段
	reserve1 types.TThostFtdcOldInstrumentIDType
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
	// 本地申请组合编号
	ActionLocalID types.TThostFtdcOrderLocalIDType
	// 交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	// 会员代码
	ParticipantID types.TThostFtdcParticipantIDType
	// 客户代码
	ClientID types.TThostFtdcClientIDType
	// 保留的无效字段
	reserve2 types.TThostFtdcOldExchangeInstIDType
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
	// 保留的无效字段
	reserve3 types.TThostFtdcOldIPAddressType
	// Mac地址
	MacAddress types.TThostFtdcMacAddressType
	// 组合编号
	ComTradeID types.TThostFtdcTradeIDType
	// 营业部编号
	BranchID types.TThostFtdcBranchIDType
	// 投资单元代码
	InvestUnitID types.TThostFtdcInvestUnitIDType
	// 合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
	// 合约在交易所的代码
	ExchangeInstID types.TThostFtdcExchangeInstIDType
	// IP地址
	IPAddress types.TThostFtdcIPAddressType
}

func (d CThostFtdcCombActionField) Type() string { return "CThostFtdcCombActionField" }

func (d CThostFtdcCombActionField) String() string {
	var builder strings.Builder
	builder.Grow(687)

	builder.WriteString("CThostFtdcCombActionField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", InvestorID=%+v", d.InvestorID)
	fmt.Fprintf(&builder, ", reserve1=%+v", d.reserve1)
	fmt.Fprintf(&builder, ", CombActionRef=%+v", d.CombActionRef)
	fmt.Fprintf(&builder, ", UserID=%+v", d.UserID)
	fmt.Fprintf(&builder, ", Direction=%+v", d.Direction)
	fmt.Fprintf(&builder, ", Volume=%+v", d.Volume)
	fmt.Fprintf(&builder, ", CombDirection=%+v", d.CombDirection)
	fmt.Fprintf(&builder, ", HedgeFlag=%+v", d.HedgeFlag)
	fmt.Fprintf(&builder, ", ActionLocalID=%+v", d.ActionLocalID)
	fmt.Fprintf(&builder, ", ExchangeID=%+v", d.ExchangeID)
	fmt.Fprintf(&builder, ", ParticipantID=%+v", d.ParticipantID)
	fmt.Fprintf(&builder, ", ClientID=%+v", d.ClientID)
	fmt.Fprintf(&builder, ", reserve2=%+v", d.reserve2)
	fmt.Fprintf(&builder, ", TraderID=%+v", d.TraderID)
	fmt.Fprintf(&builder, ", InstallID=%+v", d.InstallID)
	fmt.Fprintf(&builder, ", ActionStatus=%+v", d.ActionStatus)
	fmt.Fprintf(&builder, ", NotifySequence=%+v", d.NotifySequence)
	fmt.Fprintf(&builder, ", TradingDay=%+v", d.TradingDay)
	fmt.Fprintf(&builder, ", SettlementID=%+v", d.SettlementID)
	fmt.Fprintf(&builder, ", SequenceNo=%+v", d.SequenceNo)
	fmt.Fprintf(&builder, ", FrontID=%+v", d.FrontID)
	fmt.Fprintf(&builder, ", SessionID=%+v", d.SessionID)
	fmt.Fprintf(&builder, ", UserProductInfo=%+v", d.UserProductInfo)
	fmt.Fprintf(&builder, ", StatusMsg=%+v", d.StatusMsg)
	fmt.Fprintf(&builder, ", reserve3=%+v", d.reserve3)
	fmt.Fprintf(&builder, ", MacAddress=%+v", d.MacAddress)
	fmt.Fprintf(&builder, ", ComTradeID=%+v", d.ComTradeID)
	fmt.Fprintf(&builder, ", BranchID=%+v", d.BranchID)
	fmt.Fprintf(&builder, ", InvestUnitID=%+v", d.InvestUnitID)
	fmt.Fprintf(&builder, ", InstrumentID=%+v", d.InstrumentID)
	fmt.Fprintf(&builder, ", ExchangeInstID=%+v", d.ExchangeInstID)
	fmt.Fprintf(&builder, ", IPAddress=%+v", d.IPAddress)

	builder.WriteByte('}')

	return builder.String()
}

// 申请组合查询
type CThostFtdcQryCombActionField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	// 保留的无效字段
	reserve1 types.TThostFtdcOldInstrumentIDType
	// 交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	// 投资单元代码
	InvestUnitID types.TThostFtdcInvestUnitIDType
	// 合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
}

func (d CThostFtdcQryCombActionField) Type() string { return "CThostFtdcQryCombActionField" }

func (d CThostFtdcQryCombActionField) String() string {
	var builder strings.Builder
	builder.Grow(148)

	builder.WriteString("CThostFtdcQryCombActionField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", InvestorID=%+v", d.InvestorID)
	fmt.Fprintf(&builder, ", reserve1=%+v", d.reserve1)
	fmt.Fprintf(&builder, ", ExchangeID=%+v", d.ExchangeID)
	fmt.Fprintf(&builder, ", InvestUnitID=%+v", d.InvestUnitID)
	fmt.Fprintf(&builder, ", InstrumentID=%+v", d.InstrumentID)

	builder.WriteByte('}')

	return builder.String()
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
	// 保留的无效字段
	reserve1 types.TThostFtdcOldExchangeInstIDType
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
	// 保留的无效字段
	reserve2 types.TThostFtdcOldIPAddressType
	// Mac地址
	MacAddress types.TThostFtdcMacAddressType
	// 组合编号
	ComTradeID types.TThostFtdcTradeIDType
	// 营业部编号
	BranchID types.TThostFtdcBranchIDType
	// 合约在交易所的代码
	ExchangeInstID types.TThostFtdcExchangeInstIDType
	// IP地址
	IPAddress types.TThostFtdcIPAddressType
}

func (d CThostFtdcExchangeCombActionField) Type() string { return "CThostFtdcExchangeCombActionField" }

func (d CThostFtdcExchangeCombActionField) String() string {
	var builder strings.Builder
	builder.Grow(476)

	builder.WriteString("CThostFtdcExchangeCombActionField{")
	fmt.Fprintf(&builder, "Direction=%+v", d.Direction)
	fmt.Fprintf(&builder, ", Volume=%+v", d.Volume)
	fmt.Fprintf(&builder, ", CombDirection=%+v", d.CombDirection)
	fmt.Fprintf(&builder, ", HedgeFlag=%+v", d.HedgeFlag)
	fmt.Fprintf(&builder, ", ActionLocalID=%+v", d.ActionLocalID)
	fmt.Fprintf(&builder, ", ExchangeID=%+v", d.ExchangeID)
	fmt.Fprintf(&builder, ", ParticipantID=%+v", d.ParticipantID)
	fmt.Fprintf(&builder, ", ClientID=%+v", d.ClientID)
	fmt.Fprintf(&builder, ", reserve1=%+v", d.reserve1)
	fmt.Fprintf(&builder, ", TraderID=%+v", d.TraderID)
	fmt.Fprintf(&builder, ", InstallID=%+v", d.InstallID)
	fmt.Fprintf(&builder, ", ActionStatus=%+v", d.ActionStatus)
	fmt.Fprintf(&builder, ", NotifySequence=%+v", d.NotifySequence)
	fmt.Fprintf(&builder, ", TradingDay=%+v", d.TradingDay)
	fmt.Fprintf(&builder, ", SettlementID=%+v", d.SettlementID)
	fmt.Fprintf(&builder, ", SequenceNo=%+v", d.SequenceNo)
	fmt.Fprintf(&builder, ", reserve2=%+v", d.reserve2)
	fmt.Fprintf(&builder, ", MacAddress=%+v", d.MacAddress)
	fmt.Fprintf(&builder, ", ComTradeID=%+v", d.ComTradeID)
	fmt.Fprintf(&builder, ", BranchID=%+v", d.BranchID)
	fmt.Fprintf(&builder, ", ExchangeInstID=%+v", d.ExchangeInstID)
	fmt.Fprintf(&builder, ", IPAddress=%+v", d.IPAddress)

	builder.WriteByte('}')

	return builder.String()
}

// 交易所申请组合查询
type CThostFtdcQryExchangeCombActionField struct {
	// 会员代码
	ParticipantID types.TThostFtdcParticipantIDType
	// 客户代码
	ClientID types.TThostFtdcClientIDType
	// 保留的无效字段
	reserve1 types.TThostFtdcOldExchangeInstIDType
	// 交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	// 交易所交易员代码
	TraderID types.TThostFtdcTraderIDType
	// 合约在交易所的代码
	ExchangeInstID types.TThostFtdcExchangeInstIDType
}

func (d CThostFtdcQryExchangeCombActionField) Type() string {
	return "CThostFtdcQryExchangeCombActionField"
}

func (d CThostFtdcQryExchangeCombActionField) String() string {
	var builder strings.Builder
	builder.Grow(157)

	builder.WriteString("CThostFtdcQryExchangeCombActionField{")
	fmt.Fprintf(&builder, "ParticipantID=%+v", d.ParticipantID)
	fmt.Fprintf(&builder, ", ClientID=%+v", d.ClientID)
	fmt.Fprintf(&builder, ", reserve1=%+v", d.reserve1)
	fmt.Fprintf(&builder, ", ExchangeID=%+v", d.ExchangeID)
	fmt.Fprintf(&builder, ", TraderID=%+v", d.TraderID)
	fmt.Fprintf(&builder, ", ExchangeInstID=%+v", d.ExchangeInstID)

	builder.WriteByte('}')

	return builder.String()
}

// 产品报价汇率
type CThostFtdcProductExchRateField struct {
	// 保留的无效字段
	reserve1 types.TThostFtdcOldInstrumentIDType
	// 报价币种类型
	QuoteCurrencyID types.TThostFtdcCurrencyIDType
	// 汇率
	ExchangeRate types.TThostFtdcExchangeRateType
	// 交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	// 产品代码
	ProductID types.TThostFtdcInstrumentIDType
}

func (d CThostFtdcProductExchRateField) Type() string { return "CThostFtdcProductExchRateField" }

func (d CThostFtdcProductExchRateField) String() string {
	var builder strings.Builder
	builder.Grow(134)

	builder.WriteString("CThostFtdcProductExchRateField{")
	fmt.Fprintf(&builder, "reserve1=%+v", d.reserve1)
	fmt.Fprintf(&builder, ", QuoteCurrencyID=%+v", d.QuoteCurrencyID)
	fmt.Fprintf(&builder, ", ExchangeRate=%+v", d.ExchangeRate)
	fmt.Fprintf(&builder, ", ExchangeID=%+v", d.ExchangeID)
	fmt.Fprintf(&builder, ", ProductID=%+v", d.ProductID)

	builder.WriteByte('}')

	return builder.String()
}

// 产品报价汇率查询
type CThostFtdcQryProductExchRateField struct {
	// 保留的无效字段
	reserve1 types.TThostFtdcOldInstrumentIDType
	// 交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	// 产品代码
	ProductID types.TThostFtdcInstrumentIDType
}

func (d CThostFtdcQryProductExchRateField) Type() string { return "CThostFtdcQryProductExchRateField" }

func (d CThostFtdcQryProductExchRateField) String() string {
	var builder strings.Builder
	builder.Grow(90)

	builder.WriteString("CThostFtdcQryProductExchRateField{")
	fmt.Fprintf(&builder, "reserve1=%+v", d.reserve1)
	fmt.Fprintf(&builder, ", ExchangeID=%+v", d.ExchangeID)
	fmt.Fprintf(&builder, ", ProductID=%+v", d.ProductID)

	builder.WriteByte('}')

	return builder.String()
}

// 查询询价价差参数
type CThostFtdcQryForQuoteParamField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 保留的无效字段
	reserve1 types.TThostFtdcOldInstrumentIDType
	// 交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	// 合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
}

func (d CThostFtdcQryForQuoteParamField) Type() string { return "CThostFtdcQryForQuoteParamField" }

func (d CThostFtdcQryForQuoteParamField) String() string {
	var builder strings.Builder
	builder.Grow(109)

	builder.WriteString("CThostFtdcQryForQuoteParamField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", reserve1=%+v", d.reserve1)
	fmt.Fprintf(&builder, ", ExchangeID=%+v", d.ExchangeID)
	fmt.Fprintf(&builder, ", InstrumentID=%+v", d.InstrumentID)

	builder.WriteByte('}')

	return builder.String()
}

// 询价价差参数
type CThostFtdcForQuoteParamField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 保留的无效字段
	reserve1 types.TThostFtdcOldInstrumentIDType
	// 交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	// 最新价
	LastPrice types.TThostFtdcPriceType
	// 价差
	PriceInterval types.TThostFtdcPriceType
	// 合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
}

func (d CThostFtdcForQuoteParamField) Type() string { return "CThostFtdcForQuoteParamField" }

func (d CThostFtdcForQuoteParamField) String() string {
	var builder strings.Builder
	builder.Grow(148)

	builder.WriteString("CThostFtdcForQuoteParamField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", reserve1=%+v", d.reserve1)
	fmt.Fprintf(&builder, ", ExchangeID=%+v", d.ExchangeID)
	fmt.Fprintf(&builder, ", LastPrice=%+v", d.LastPrice)
	fmt.Fprintf(&builder, ", PriceInterval=%+v", d.PriceInterval)
	fmt.Fprintf(&builder, ", InstrumentID=%+v", d.InstrumentID)

	builder.WriteByte('}')

	return builder.String()
}

// 当前做市商期权合约手续费的详细内容
type CThostFtdcMMOptionInstrCommRateField struct {
	// 保留的无效字段
	reserve1 types.TThostFtdcOldInstrumentIDType
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
	// 合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
}

func (d CThostFtdcMMOptionInstrCommRateField) Type() string {
	return "CThostFtdcMMOptionInstrCommRateField"
}

func (d CThostFtdcMMOptionInstrCommRateField) String() string {
	var builder strings.Builder
	builder.Grow(367)

	builder.WriteString("CThostFtdcMMOptionInstrCommRateField{")
	fmt.Fprintf(&builder, "reserve1=%+v", d.reserve1)
	fmt.Fprintf(&builder, ", InvestorRange=%+v", d.InvestorRange)
	fmt.Fprintf(&builder, ", BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", InvestorID=%+v", d.InvestorID)
	fmt.Fprintf(&builder, ", OpenRatioByMoney=%+v", d.OpenRatioByMoney)
	fmt.Fprintf(&builder, ", OpenRatioByVolume=%+v", d.OpenRatioByVolume)
	fmt.Fprintf(&builder, ", CloseRatioByMoney=%+v", d.CloseRatioByMoney)
	fmt.Fprintf(&builder, ", CloseRatioByVolume=%+v", d.CloseRatioByVolume)
	fmt.Fprintf(&builder, ", CloseTodayRatioByMoney=%+v", d.CloseTodayRatioByMoney)
	fmt.Fprintf(&builder, ", CloseTodayRatioByVolume=%+v", d.CloseTodayRatioByVolume)
	fmt.Fprintf(&builder, ", StrikeRatioByMoney=%+v", d.StrikeRatioByMoney)
	fmt.Fprintf(&builder, ", StrikeRatioByVolume=%+v", d.StrikeRatioByVolume)
	fmt.Fprintf(&builder, ", InstrumentID=%+v", d.InstrumentID)

	builder.WriteByte('}')

	return builder.String()
}

// 做市商期权手续费率查询
type CThostFtdcQryMMOptionInstrCommRateField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	// 保留的无效字段
	reserve1 types.TThostFtdcOldInstrumentIDType
	// 合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
}

func (d CThostFtdcQryMMOptionInstrCommRateField) Type() string {
	return "CThostFtdcQryMMOptionInstrCommRateField"
}

func (d CThostFtdcQryMMOptionInstrCommRateField) String() string {
	var builder strings.Builder
	builder.Grow(117)

	builder.WriteString("CThostFtdcQryMMOptionInstrCommRateField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", InvestorID=%+v", d.InvestorID)
	fmt.Fprintf(&builder, ", reserve1=%+v", d.reserve1)
	fmt.Fprintf(&builder, ", InstrumentID=%+v", d.InstrumentID)

	builder.WriteByte('}')

	return builder.String()
}

// 做市商合约手续费率
type CThostFtdcMMInstrumentCommissionRateField struct {
	// 保留的无效字段
	reserve1 types.TThostFtdcOldInstrumentIDType
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
	// 合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
}

func (d CThostFtdcMMInstrumentCommissionRateField) Type() string {
	return "CThostFtdcMMInstrumentCommissionRateField"
}

func (d CThostFtdcMMInstrumentCommissionRateField) String() string {
	var builder strings.Builder
	builder.Grow(315)

	builder.WriteString("CThostFtdcMMInstrumentCommissionRateField{")
	fmt.Fprintf(&builder, "reserve1=%+v", d.reserve1)
	fmt.Fprintf(&builder, ", InvestorRange=%+v", d.InvestorRange)
	fmt.Fprintf(&builder, ", BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", InvestorID=%+v", d.InvestorID)
	fmt.Fprintf(&builder, ", OpenRatioByMoney=%+v", d.OpenRatioByMoney)
	fmt.Fprintf(&builder, ", OpenRatioByVolume=%+v", d.OpenRatioByVolume)
	fmt.Fprintf(&builder, ", CloseRatioByMoney=%+v", d.CloseRatioByMoney)
	fmt.Fprintf(&builder, ", CloseRatioByVolume=%+v", d.CloseRatioByVolume)
	fmt.Fprintf(&builder, ", CloseTodayRatioByMoney=%+v", d.CloseTodayRatioByMoney)
	fmt.Fprintf(&builder, ", CloseTodayRatioByVolume=%+v", d.CloseTodayRatioByVolume)
	fmt.Fprintf(&builder, ", InstrumentID=%+v", d.InstrumentID)

	builder.WriteByte('}')

	return builder.String()
}

// 查询做市商合约手续费率
type CThostFtdcQryMMInstrumentCommissionRateField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	// 保留的无效字段
	reserve1 types.TThostFtdcOldInstrumentIDType
	// 合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
}

func (d CThostFtdcQryMMInstrumentCommissionRateField) Type() string {
	return "CThostFtdcQryMMInstrumentCommissionRateField"
}

func (d CThostFtdcQryMMInstrumentCommissionRateField) String() string {
	var builder strings.Builder
	builder.Grow(122)

	builder.WriteString("CThostFtdcQryMMInstrumentCommissionRateField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", InvestorID=%+v", d.InvestorID)
	fmt.Fprintf(&builder, ", reserve1=%+v", d.reserve1)
	fmt.Fprintf(&builder, ", InstrumentID=%+v", d.InstrumentID)

	builder.WriteByte('}')

	return builder.String()
}

// 当前报单手续费的详细内容
type CThostFtdcInstrumentOrderCommRateField struct {
	// 保留的无效字段
	reserve1 types.TThostFtdcOldInstrumentIDType
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
	// 交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	// 投资单元代码
	InvestUnitID types.TThostFtdcInvestUnitIDType
	// 合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
	// 报单手续费
	OrderCommByTrade types.TThostFtdcRatioType
	// 撤单手续费
	OrderActionCommByTrade types.TThostFtdcRatioType
}

func (d CThostFtdcInstrumentOrderCommRateField) Type() string {
	return "CThostFtdcInstrumentOrderCommRateField"
}

func (d CThostFtdcInstrumentOrderCommRateField) String() string {
	var builder strings.Builder
	builder.Grow(318)

	builder.WriteString("CThostFtdcInstrumentOrderCommRateField{")
	fmt.Fprintf(&builder, "reserve1=%+v", d.reserve1)
	fmt.Fprintf(&builder, ", InvestorRange=%+v", d.InvestorRange)
	fmt.Fprintf(&builder, ", BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", InvestorID=%+v", d.InvestorID)
	fmt.Fprintf(&builder, ", HedgeFlag=%+v", d.HedgeFlag)
	fmt.Fprintf(&builder, ", OrderCommByVolume=%+v", d.OrderCommByVolume)
	fmt.Fprintf(&builder, ", OrderActionCommByVolume=%+v", d.OrderActionCommByVolume)
	fmt.Fprintf(&builder, ", ExchangeID=%+v", d.ExchangeID)
	fmt.Fprintf(&builder, ", InvestUnitID=%+v", d.InvestUnitID)
	fmt.Fprintf(&builder, ", InstrumentID=%+v", d.InstrumentID)
	fmt.Fprintf(&builder, ", OrderCommByTrade=%+v", d.OrderCommByTrade)
	fmt.Fprintf(&builder, ", OrderActionCommByTrade=%+v", d.OrderActionCommByTrade)

	builder.WriteByte('}')

	return builder.String()
}

// 报单手续费率查询
type CThostFtdcQryInstrumentOrderCommRateField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	// 保留的无效字段
	reserve1 types.TThostFtdcOldInstrumentIDType
	// 合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
}

func (d CThostFtdcQryInstrumentOrderCommRateField) Type() string {
	return "CThostFtdcQryInstrumentOrderCommRateField"
}

func (d CThostFtdcQryInstrumentOrderCommRateField) String() string {
	var builder strings.Builder
	builder.Grow(119)

	builder.WriteString("CThostFtdcQryInstrumentOrderCommRateField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", InvestorID=%+v", d.InvestorID)
	fmt.Fprintf(&builder, ", reserve1=%+v", d.reserve1)
	fmt.Fprintf(&builder, ", InstrumentID=%+v", d.InstrumentID)

	builder.WriteByte('}')

	return builder.String()
}

// 交易参数
type CThostFtdcTradeParamField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 参数代码
	TradeParamID types.TThostFtdcTradeParamIDType
	// 参数代码值
	TradeParamValue types.TThostFtdcSettlementParamValueType
	// 备注
	Memo types.TThostFtdcMemoType
}

func (d CThostFtdcTradeParamField) Type() string { return "CThostFtdcTradeParamField" }

func (d CThostFtdcTradeParamField) String() string {
	var builder strings.Builder
	builder.Grow(104)

	builder.WriteString("CThostFtdcTradeParamField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", TradeParamID=%+v", d.TradeParamID)
	fmt.Fprintf(&builder, ", TradeParamValue=%+v", d.TradeParamValue)
	fmt.Fprintf(&builder, ", Memo=%+v", d.Memo)

	builder.WriteByte('}')

	return builder.String()
}

// 合约保证金率调整
type CThostFtdcInstrumentMarginRateULField struct {
	// 保留的无效字段
	reserve1 types.TThostFtdcOldInstrumentIDType
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
	// 合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
}

func (d CThostFtdcInstrumentMarginRateULField) Type() string {
	return "CThostFtdcInstrumentMarginRateULField"
}

func (d CThostFtdcInstrumentMarginRateULField) String() string {
	var builder strings.Builder
	builder.Grow(289)

	builder.WriteString("CThostFtdcInstrumentMarginRateULField{")
	fmt.Fprintf(&builder, "reserve1=%+v", d.reserve1)
	fmt.Fprintf(&builder, ", InvestorRange=%+v", d.InvestorRange)
	fmt.Fprintf(&builder, ", BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", InvestorID=%+v", d.InvestorID)
	fmt.Fprintf(&builder, ", HedgeFlag=%+v", d.HedgeFlag)
	fmt.Fprintf(&builder, ", LongMarginRatioByMoney=%+v", d.LongMarginRatioByMoney)
	fmt.Fprintf(&builder, ", LongMarginRatioByVolume=%+v", d.LongMarginRatioByVolume)
	fmt.Fprintf(&builder, ", ShortMarginRatioByMoney=%+v", d.ShortMarginRatioByMoney)
	fmt.Fprintf(&builder, ", ShortMarginRatioByVolume=%+v", d.ShortMarginRatioByVolume)
	fmt.Fprintf(&builder, ", InstrumentID=%+v", d.InstrumentID)

	builder.WriteByte('}')

	return builder.String()
}

// 期货持仓限制参数
type CThostFtdcFutureLimitPosiParamField struct {
	// 投资者范围
	InvestorRange types.TThostFtdcInvestorRangeType
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	// 保留的无效字段
	reserve1 types.TThostFtdcOldInstrumentIDType
	// 当日投机开仓数量限制
	SpecOpenVolume types.TThostFtdcVolumeType
	// 当日套利开仓数量限制
	ArbiOpenVolume types.TThostFtdcVolumeType
	// 当日投机+套利开仓数量限制
	OpenVolume types.TThostFtdcVolumeType
	// 产品代码
	ProductID types.TThostFtdcInstrumentIDType
}

func (d CThostFtdcFutureLimitPosiParamField) Type() string {
	return "CThostFtdcFutureLimitPosiParamField"
}

func (d CThostFtdcFutureLimitPosiParamField) String() string {
	var builder strings.Builder
	builder.Grow(201)

	builder.WriteString("CThostFtdcFutureLimitPosiParamField{")
	fmt.Fprintf(&builder, "InvestorRange=%+v", d.InvestorRange)
	fmt.Fprintf(&builder, ", BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", InvestorID=%+v", d.InvestorID)
	fmt.Fprintf(&builder, ", reserve1=%+v", d.reserve1)
	fmt.Fprintf(&builder, ", SpecOpenVolume=%+v", d.SpecOpenVolume)
	fmt.Fprintf(&builder, ", ArbiOpenVolume=%+v", d.ArbiOpenVolume)
	fmt.Fprintf(&builder, ", OpenVolume=%+v", d.OpenVolume)
	fmt.Fprintf(&builder, ", ProductID=%+v", d.ProductID)

	builder.WriteByte('}')

	return builder.String()
}

// 禁止登录IP
type CThostFtdcLoginForbiddenIPField struct {
	// 保留的无效字段
	reserve1 types.TThostFtdcOldIPAddressType
	// IP地址
	IPAddress types.TThostFtdcIPAddressType
}

func (d CThostFtdcLoginForbiddenIPField) Type() string { return "CThostFtdcLoginForbiddenIPField" }

func (d CThostFtdcLoginForbiddenIPField) String() string {
	var builder strings.Builder
	builder.Grow(68)

	builder.WriteString("CThostFtdcLoginForbiddenIPField{")
	fmt.Fprintf(&builder, "reserve1=%+v", d.reserve1)
	fmt.Fprintf(&builder, ", IPAddress=%+v", d.IPAddress)

	builder.WriteByte('}')

	return builder.String()
}

// IP列表
type CThostFtdcIPListField struct {
	// 保留的无效字段
	reserve1 types.TThostFtdcOldIPAddressType
	// 是否白名单
	IsWhite types.TThostFtdcBoolType
	// IP地址
	IPAddress types.TThostFtdcIPAddressType
}

func (d CThostFtdcIPListField) Type() string { return "CThostFtdcIPListField" }

func (d CThostFtdcIPListField) String() string {
	var builder strings.Builder
	builder.Grow(75)

	builder.WriteString("CThostFtdcIPListField{")
	fmt.Fprintf(&builder, "reserve1=%+v", d.reserve1)
	fmt.Fprintf(&builder, ", IsWhite=%+v", d.IsWhite)
	fmt.Fprintf(&builder, ", IPAddress=%+v", d.IPAddress)

	builder.WriteByte('}')

	return builder.String()
}

// 输入的期权自对冲
type CThostFtdcInputOptionSelfCloseField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	// 保留的无效字段
	reserve1 types.TThostFtdcOldInstrumentIDType
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
	// 保留的无效字段
	reserve2 types.TThostFtdcOldIPAddressType
	// Mac地址
	MacAddress types.TThostFtdcMacAddressType
	// 合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
	// IP地址
	IPAddress types.TThostFtdcIPAddressType
}

func (d CThostFtdcInputOptionSelfCloseField) Type() string {
	return "CThostFtdcInputOptionSelfCloseField"
}

func (d CThostFtdcInputOptionSelfCloseField) String() string {
	var builder strings.Builder
	builder.Grow(415)

	builder.WriteString("CThostFtdcInputOptionSelfCloseField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", InvestorID=%+v", d.InvestorID)
	fmt.Fprintf(&builder, ", reserve1=%+v", d.reserve1)
	fmt.Fprintf(&builder, ", OptionSelfCloseRef=%+v", d.OptionSelfCloseRef)
	fmt.Fprintf(&builder, ", UserID=%+v", d.UserID)
	fmt.Fprintf(&builder, ", Volume=%+v", d.Volume)
	fmt.Fprintf(&builder, ", RequestID=%+v", d.RequestID)
	fmt.Fprintf(&builder, ", BusinessUnit=%+v", d.BusinessUnit)
	fmt.Fprintf(&builder, ", HedgeFlag=%+v", d.HedgeFlag)
	fmt.Fprintf(&builder, ", OptSelfCloseFlag=%+v", d.OptSelfCloseFlag)
	fmt.Fprintf(&builder, ", ExchangeID=%+v", d.ExchangeID)
	fmt.Fprintf(&builder, ", InvestUnitID=%+v", d.InvestUnitID)
	fmt.Fprintf(&builder, ", AccountID=%+v", d.AccountID)
	fmt.Fprintf(&builder, ", CurrencyID=%+v", d.CurrencyID)
	fmt.Fprintf(&builder, ", ClientID=%+v", d.ClientID)
	fmt.Fprintf(&builder, ", reserve2=%+v", d.reserve2)
	fmt.Fprintf(&builder, ", MacAddress=%+v", d.MacAddress)
	fmt.Fprintf(&builder, ", InstrumentID=%+v", d.InstrumentID)
	fmt.Fprintf(&builder, ", IPAddress=%+v", d.IPAddress)

	builder.WriteByte('}')

	return builder.String()
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
	// 保留的无效字段
	reserve1 types.TThostFtdcOldInstrumentIDType
	// 投资单元代码
	InvestUnitID types.TThostFtdcInvestUnitIDType
	// 保留的无效字段
	reserve2 types.TThostFtdcOldIPAddressType
	// Mac地址
	MacAddress types.TThostFtdcMacAddressType
	// 合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
	// IP地址
	IPAddress types.TThostFtdcIPAddressType
}

func (d CThostFtdcInputOptionSelfCloseActionField) Type() string {
	return "CThostFtdcInputOptionSelfCloseActionField"
}

func (d CThostFtdcInputOptionSelfCloseActionField) String() string {
	var builder strings.Builder
	builder.Grow(401)

	builder.WriteString("CThostFtdcInputOptionSelfCloseActionField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", InvestorID=%+v", d.InvestorID)
	fmt.Fprintf(&builder, ", OptionSelfCloseActionRef=%+v", d.OptionSelfCloseActionRef)
	fmt.Fprintf(&builder, ", OptionSelfCloseRef=%+v", d.OptionSelfCloseRef)
	fmt.Fprintf(&builder, ", RequestID=%+v", d.RequestID)
	fmt.Fprintf(&builder, ", FrontID=%+v", d.FrontID)
	fmt.Fprintf(&builder, ", SessionID=%+v", d.SessionID)
	fmt.Fprintf(&builder, ", ExchangeID=%+v", d.ExchangeID)
	fmt.Fprintf(&builder, ", OptionSelfCloseSysID=%+v", d.OptionSelfCloseSysID)
	fmt.Fprintf(&builder, ", ActionFlag=%+v", d.ActionFlag)
	fmt.Fprintf(&builder, ", UserID=%+v", d.UserID)
	fmt.Fprintf(&builder, ", reserve1=%+v", d.reserve1)
	fmt.Fprintf(&builder, ", InvestUnitID=%+v", d.InvestUnitID)
	fmt.Fprintf(&builder, ", reserve2=%+v", d.reserve2)
	fmt.Fprintf(&builder, ", MacAddress=%+v", d.MacAddress)
	fmt.Fprintf(&builder, ", InstrumentID=%+v", d.InstrumentID)
	fmt.Fprintf(&builder, ", IPAddress=%+v", d.IPAddress)

	builder.WriteByte('}')

	return builder.String()
}

// 期权自对冲
type CThostFtdcOptionSelfCloseField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	// 保留的无效字段
	reserve1 types.TThostFtdcOldInstrumentIDType
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
	// 保留的无效字段
	reserve2 types.TThostFtdcOldExchangeInstIDType
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
	// 保留的无效字段
	reserve3 types.TThostFtdcOldIPAddressType
	// Mac地址
	MacAddress types.TThostFtdcMacAddressType
	// 合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
	// 合约在交易所的代码
	ExchangeInstID types.TThostFtdcExchangeInstIDType
	// IP地址
	IPAddress types.TThostFtdcIPAddressType
}

func (d CThostFtdcOptionSelfCloseField) Type() string { return "CThostFtdcOptionSelfCloseField" }

func (d CThostFtdcOptionSelfCloseField) String() string {
	var builder strings.Builder
	builder.Grow(945)

	builder.WriteString("CThostFtdcOptionSelfCloseField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", InvestorID=%+v", d.InvestorID)
	fmt.Fprintf(&builder, ", reserve1=%+v", d.reserve1)
	fmt.Fprintf(&builder, ", OptionSelfCloseRef=%+v", d.OptionSelfCloseRef)
	fmt.Fprintf(&builder, ", UserID=%+v", d.UserID)
	fmt.Fprintf(&builder, ", Volume=%+v", d.Volume)
	fmt.Fprintf(&builder, ", RequestID=%+v", d.RequestID)
	fmt.Fprintf(&builder, ", BusinessUnit=%+v", d.BusinessUnit)
	fmt.Fprintf(&builder, ", HedgeFlag=%+v", d.HedgeFlag)
	fmt.Fprintf(&builder, ", OptSelfCloseFlag=%+v", d.OptSelfCloseFlag)
	fmt.Fprintf(&builder, ", OptionSelfCloseLocalID=%+v", d.OptionSelfCloseLocalID)
	fmt.Fprintf(&builder, ", ExchangeID=%+v", d.ExchangeID)
	fmt.Fprintf(&builder, ", ParticipantID=%+v", d.ParticipantID)
	fmt.Fprintf(&builder, ", ClientID=%+v", d.ClientID)
	fmt.Fprintf(&builder, ", reserve2=%+v", d.reserve2)
	fmt.Fprintf(&builder, ", TraderID=%+v", d.TraderID)
	fmt.Fprintf(&builder, ", InstallID=%+v", d.InstallID)
	fmt.Fprintf(&builder, ", OrderSubmitStatus=%+v", d.OrderSubmitStatus)
	fmt.Fprintf(&builder, ", NotifySequence=%+v", d.NotifySequence)
	fmt.Fprintf(&builder, ", TradingDay=%+v", d.TradingDay)
	fmt.Fprintf(&builder, ", SettlementID=%+v", d.SettlementID)
	fmt.Fprintf(&builder, ", OptionSelfCloseSysID=%+v", d.OptionSelfCloseSysID)
	fmt.Fprintf(&builder, ", InsertDate=%+v", d.InsertDate)
	fmt.Fprintf(&builder, ", InsertTime=%+v", d.InsertTime)
	fmt.Fprintf(&builder, ", CancelTime=%+v", d.CancelTime)
	fmt.Fprintf(&builder, ", ExecResult=%+v", d.ExecResult)
	fmt.Fprintf(&builder, ", ClearingPartID=%+v", d.ClearingPartID)
	fmt.Fprintf(&builder, ", SequenceNo=%+v", d.SequenceNo)
	fmt.Fprintf(&builder, ", FrontID=%+v", d.FrontID)
	fmt.Fprintf(&builder, ", SessionID=%+v", d.SessionID)
	fmt.Fprintf(&builder, ", UserProductInfo=%+v", d.UserProductInfo)
	fmt.Fprintf(&builder, ", StatusMsg=%+v", d.StatusMsg)
	fmt.Fprintf(&builder, ", ActiveUserID=%+v", d.ActiveUserID)
	fmt.Fprintf(&builder, ", BrokerOptionSelfCloseSeq=%+v", d.BrokerOptionSelfCloseSeq)
	fmt.Fprintf(&builder, ", BranchID=%+v", d.BranchID)
	fmt.Fprintf(&builder, ", InvestUnitID=%+v", d.InvestUnitID)
	fmt.Fprintf(&builder, ", AccountID=%+v", d.AccountID)
	fmt.Fprintf(&builder, ", CurrencyID=%+v", d.CurrencyID)
	fmt.Fprintf(&builder, ", reserve3=%+v", d.reserve3)
	fmt.Fprintf(&builder, ", MacAddress=%+v", d.MacAddress)
	fmt.Fprintf(&builder, ", InstrumentID=%+v", d.InstrumentID)
	fmt.Fprintf(&builder, ", ExchangeInstID=%+v", d.ExchangeInstID)
	fmt.Fprintf(&builder, ", IPAddress=%+v", d.IPAddress)

	builder.WriteByte('}')

	return builder.String()
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
	// 保留的无效字段
	reserve1 types.TThostFtdcOldInstrumentIDType
	// 营业部编号
	BranchID types.TThostFtdcBranchIDType
	// 投资单元代码
	InvestUnitID types.TThostFtdcInvestUnitIDType
	// 保留的无效字段
	reserve2 types.TThostFtdcOldIPAddressType
	// Mac地址
	MacAddress types.TThostFtdcMacAddressType
	// 合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
	// IP地址
	IPAddress types.TThostFtdcIPAddressType
}

func (d CThostFtdcOptionSelfCloseActionField) Type() string {
	return "CThostFtdcOptionSelfCloseActionField"
}

func (d CThostFtdcOptionSelfCloseActionField) String() string {
	var builder strings.Builder
	builder.Grow(655)

	builder.WriteString("CThostFtdcOptionSelfCloseActionField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", InvestorID=%+v", d.InvestorID)
	fmt.Fprintf(&builder, ", OptionSelfCloseActionRef=%+v", d.OptionSelfCloseActionRef)
	fmt.Fprintf(&builder, ", OptionSelfCloseRef=%+v", d.OptionSelfCloseRef)
	fmt.Fprintf(&builder, ", RequestID=%+v", d.RequestID)
	fmt.Fprintf(&builder, ", FrontID=%+v", d.FrontID)
	fmt.Fprintf(&builder, ", SessionID=%+v", d.SessionID)
	fmt.Fprintf(&builder, ", ExchangeID=%+v", d.ExchangeID)
	fmt.Fprintf(&builder, ", OptionSelfCloseSysID=%+v", d.OptionSelfCloseSysID)
	fmt.Fprintf(&builder, ", ActionFlag=%+v", d.ActionFlag)
	fmt.Fprintf(&builder, ", ActionDate=%+v", d.ActionDate)
	fmt.Fprintf(&builder, ", ActionTime=%+v", d.ActionTime)
	fmt.Fprintf(&builder, ", TraderID=%+v", d.TraderID)
	fmt.Fprintf(&builder, ", InstallID=%+v", d.InstallID)
	fmt.Fprintf(&builder, ", OptionSelfCloseLocalID=%+v", d.OptionSelfCloseLocalID)
	fmt.Fprintf(&builder, ", ActionLocalID=%+v", d.ActionLocalID)
	fmt.Fprintf(&builder, ", ParticipantID=%+v", d.ParticipantID)
	fmt.Fprintf(&builder, ", ClientID=%+v", d.ClientID)
	fmt.Fprintf(&builder, ", BusinessUnit=%+v", d.BusinessUnit)
	fmt.Fprintf(&builder, ", OrderActionStatus=%+v", d.OrderActionStatus)
	fmt.Fprintf(&builder, ", UserID=%+v", d.UserID)
	fmt.Fprintf(&builder, ", StatusMsg=%+v", d.StatusMsg)
	fmt.Fprintf(&builder, ", reserve1=%+v", d.reserve1)
	fmt.Fprintf(&builder, ", BranchID=%+v", d.BranchID)
	fmt.Fprintf(&builder, ", InvestUnitID=%+v", d.InvestUnitID)
	fmt.Fprintf(&builder, ", reserve2=%+v", d.reserve2)
	fmt.Fprintf(&builder, ", MacAddress=%+v", d.MacAddress)
	fmt.Fprintf(&builder, ", InstrumentID=%+v", d.InstrumentID)
	fmt.Fprintf(&builder, ", IPAddress=%+v", d.IPAddress)

	builder.WriteByte('}')

	return builder.String()
}

// 期权自对冲查询
type CThostFtdcQryOptionSelfCloseField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	// 保留的无效字段
	reserve1 types.TThostFtdcOldInstrumentIDType
	// 交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	// 期权自对冲编号
	OptionSelfCloseSysID types.TThostFtdcOrderSysIDType
	// 开始时间
	InsertTimeStart types.TThostFtdcTimeType
	// 结束时间
	InsertTimeEnd types.TThostFtdcTimeType
	// 合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
}

func (d CThostFtdcQryOptionSelfCloseField) Type() string { return "CThostFtdcQryOptionSelfCloseField" }

func (d CThostFtdcQryOptionSelfCloseField) String() string {
	var builder strings.Builder
	builder.Grow(209)

	builder.WriteString("CThostFtdcQryOptionSelfCloseField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", InvestorID=%+v", d.InvestorID)
	fmt.Fprintf(&builder, ", reserve1=%+v", d.reserve1)
	fmt.Fprintf(&builder, ", ExchangeID=%+v", d.ExchangeID)
	fmt.Fprintf(&builder, ", OptionSelfCloseSysID=%+v", d.OptionSelfCloseSysID)
	fmt.Fprintf(&builder, ", InsertTimeStart=%+v", d.InsertTimeStart)
	fmt.Fprintf(&builder, ", InsertTimeEnd=%+v", d.InsertTimeEnd)
	fmt.Fprintf(&builder, ", InstrumentID=%+v", d.InstrumentID)

	builder.WriteByte('}')

	return builder.String()
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
	// 保留的无效字段
	reserve1 types.TThostFtdcOldExchangeInstIDType
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
	// 保留的无效字段
	reserve2 types.TThostFtdcOldIPAddressType
	// Mac地址
	MacAddress types.TThostFtdcMacAddressType
	// 合约在交易所的代码
	ExchangeInstID types.TThostFtdcExchangeInstIDType
	// IP地址
	IPAddress types.TThostFtdcIPAddressType
}

func (d CThostFtdcExchangeOptionSelfCloseField) Type() string {
	return "CThostFtdcExchangeOptionSelfCloseField"
}

func (d CThostFtdcExchangeOptionSelfCloseField) String() string {
	var builder strings.Builder
	builder.Grow(634)

	builder.WriteString("CThostFtdcExchangeOptionSelfCloseField{")
	fmt.Fprintf(&builder, "Volume=%+v", d.Volume)
	fmt.Fprintf(&builder, ", RequestID=%+v", d.RequestID)
	fmt.Fprintf(&builder, ", BusinessUnit=%+v", d.BusinessUnit)
	fmt.Fprintf(&builder, ", HedgeFlag=%+v", d.HedgeFlag)
	fmt.Fprintf(&builder, ", OptSelfCloseFlag=%+v", d.OptSelfCloseFlag)
	fmt.Fprintf(&builder, ", OptionSelfCloseLocalID=%+v", d.OptionSelfCloseLocalID)
	fmt.Fprintf(&builder, ", ExchangeID=%+v", d.ExchangeID)
	fmt.Fprintf(&builder, ", ParticipantID=%+v", d.ParticipantID)
	fmt.Fprintf(&builder, ", ClientID=%+v", d.ClientID)
	fmt.Fprintf(&builder, ", reserve1=%+v", d.reserve1)
	fmt.Fprintf(&builder, ", TraderID=%+v", d.TraderID)
	fmt.Fprintf(&builder, ", InstallID=%+v", d.InstallID)
	fmt.Fprintf(&builder, ", OrderSubmitStatus=%+v", d.OrderSubmitStatus)
	fmt.Fprintf(&builder, ", NotifySequence=%+v", d.NotifySequence)
	fmt.Fprintf(&builder, ", TradingDay=%+v", d.TradingDay)
	fmt.Fprintf(&builder, ", SettlementID=%+v", d.SettlementID)
	fmt.Fprintf(&builder, ", OptionSelfCloseSysID=%+v", d.OptionSelfCloseSysID)
	fmt.Fprintf(&builder, ", InsertDate=%+v", d.InsertDate)
	fmt.Fprintf(&builder, ", InsertTime=%+v", d.InsertTime)
	fmt.Fprintf(&builder, ", CancelTime=%+v", d.CancelTime)
	fmt.Fprintf(&builder, ", ExecResult=%+v", d.ExecResult)
	fmt.Fprintf(&builder, ", ClearingPartID=%+v", d.ClearingPartID)
	fmt.Fprintf(&builder, ", SequenceNo=%+v", d.SequenceNo)
	fmt.Fprintf(&builder, ", BranchID=%+v", d.BranchID)
	fmt.Fprintf(&builder, ", reserve2=%+v", d.reserve2)
	fmt.Fprintf(&builder, ", MacAddress=%+v", d.MacAddress)
	fmt.Fprintf(&builder, ", ExchangeInstID=%+v", d.ExchangeInstID)
	fmt.Fprintf(&builder, ", IPAddress=%+v", d.IPAddress)

	builder.WriteByte('}')

	return builder.String()
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

func (d CThostFtdcQryOptionSelfCloseActionField) Type() string {
	return "CThostFtdcQryOptionSelfCloseActionField"
}

func (d CThostFtdcQryOptionSelfCloseActionField) String() string {
	var builder strings.Builder
	builder.Grow(97)

	builder.WriteString("CThostFtdcQryOptionSelfCloseActionField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", InvestorID=%+v", d.InvestorID)
	fmt.Fprintf(&builder, ", ExchangeID=%+v", d.ExchangeID)

	builder.WriteByte('}')

	return builder.String()
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
	// 保留的无效字段
	reserve1 types.TThostFtdcOldIPAddressType
	// Mac地址
	MacAddress types.TThostFtdcMacAddressType
	// 保留的无效字段
	reserve2 types.TThostFtdcOldExchangeInstIDType
	// 期权行权的头寸是否自对冲
	OptSelfCloseFlag types.TThostFtdcOptSelfCloseFlagType
	// IP地址
	IPAddress types.TThostFtdcIPAddressType
	// 合约在交易所的代码
	ExchangeInstID types.TThostFtdcExchangeInstIDType
}

func (d CThostFtdcExchangeOptionSelfCloseActionField) Type() string {
	return "CThostFtdcExchangeOptionSelfCloseActionField"
}

func (d CThostFtdcExchangeOptionSelfCloseActionField) String() string {
	var builder strings.Builder
	builder.Grow(495)

	builder.WriteString("CThostFtdcExchangeOptionSelfCloseActionField{")
	fmt.Fprintf(&builder, "ExchangeID=%+v", d.ExchangeID)
	fmt.Fprintf(&builder, ", OptionSelfCloseSysID=%+v", d.OptionSelfCloseSysID)
	fmt.Fprintf(&builder, ", ActionFlag=%+v", d.ActionFlag)
	fmt.Fprintf(&builder, ", ActionDate=%+v", d.ActionDate)
	fmt.Fprintf(&builder, ", ActionTime=%+v", d.ActionTime)
	fmt.Fprintf(&builder, ", TraderID=%+v", d.TraderID)
	fmt.Fprintf(&builder, ", InstallID=%+v", d.InstallID)
	fmt.Fprintf(&builder, ", OptionSelfCloseLocalID=%+v", d.OptionSelfCloseLocalID)
	fmt.Fprintf(&builder, ", ActionLocalID=%+v", d.ActionLocalID)
	fmt.Fprintf(&builder, ", ParticipantID=%+v", d.ParticipantID)
	fmt.Fprintf(&builder, ", ClientID=%+v", d.ClientID)
	fmt.Fprintf(&builder, ", BusinessUnit=%+v", d.BusinessUnit)
	fmt.Fprintf(&builder, ", OrderActionStatus=%+v", d.OrderActionStatus)
	fmt.Fprintf(&builder, ", UserID=%+v", d.UserID)
	fmt.Fprintf(&builder, ", BranchID=%+v", d.BranchID)
	fmt.Fprintf(&builder, ", reserve1=%+v", d.reserve1)
	fmt.Fprintf(&builder, ", MacAddress=%+v", d.MacAddress)
	fmt.Fprintf(&builder, ", reserve2=%+v", d.reserve2)
	fmt.Fprintf(&builder, ", OptSelfCloseFlag=%+v", d.OptSelfCloseFlag)
	fmt.Fprintf(&builder, ", IPAddress=%+v", d.IPAddress)
	fmt.Fprintf(&builder, ", ExchangeInstID=%+v", d.ExchangeInstID)

	builder.WriteByte('}')

	return builder.String()
}

// 延时换汇同步
type CThostFtdcSyncDelaySwapField struct {
	// 换汇流水号
	DelaySwapSeqNo types.TThostFtdcDepositSeqNoType
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	// 源币种
	FromCurrencyID types.TThostFtdcCurrencyIDType
	// 源金额
	FromAmount types.TThostFtdcMoneyType
	// 源换汇冻结金额(可用冻结)
	FromFrozenSwap types.TThostFtdcMoneyType
	// 源剩余换汇额度(可提冻结)
	FromRemainSwap types.TThostFtdcMoneyType
	// 目标币种
	ToCurrencyID types.TThostFtdcCurrencyIDType
	// 目标金额
	ToAmount types.TThostFtdcMoneyType
	// 是否手工换汇
	IsManualSwap types.TThostFtdcBoolType
	// 是否将所有外币的剩余换汇额度设置为0
	IsAllRemainSetZero types.TThostFtdcBoolType
}

func (d CThostFtdcSyncDelaySwapField) Type() string { return "CThostFtdcSyncDelaySwapField" }

func (d CThostFtdcSyncDelaySwapField) String() string {
	var builder strings.Builder
	builder.Grow(272)

	builder.WriteString("CThostFtdcSyncDelaySwapField{")
	fmt.Fprintf(&builder, "DelaySwapSeqNo=%+v", d.DelaySwapSeqNo)
	fmt.Fprintf(&builder, ", BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", InvestorID=%+v", d.InvestorID)
	fmt.Fprintf(&builder, ", FromCurrencyID=%+v", d.FromCurrencyID)
	fmt.Fprintf(&builder, ", FromAmount=%+v", d.FromAmount)
	fmt.Fprintf(&builder, ", FromFrozenSwap=%+v", d.FromFrozenSwap)
	fmt.Fprintf(&builder, ", FromRemainSwap=%+v", d.FromRemainSwap)
	fmt.Fprintf(&builder, ", ToCurrencyID=%+v", d.ToCurrencyID)
	fmt.Fprintf(&builder, ", ToAmount=%+v", d.ToAmount)
	fmt.Fprintf(&builder, ", IsManualSwap=%+v", d.IsManualSwap)
	fmt.Fprintf(&builder, ", IsAllRemainSetZero=%+v", d.IsAllRemainSetZero)

	builder.WriteByte('}')

	return builder.String()
}

// 查询延时换汇同步
type CThostFtdcQrySyncDelaySwapField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 延时换汇流水号
	DelaySwapSeqNo types.TThostFtdcDepositSeqNoType
}

func (d CThostFtdcQrySyncDelaySwapField) Type() string { return "CThostFtdcQrySyncDelaySwapField" }

func (d CThostFtdcQrySyncDelaySwapField) String() string {
	var builder strings.Builder
	builder.Grow(73)

	builder.WriteString("CThostFtdcQrySyncDelaySwapField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", DelaySwapSeqNo=%+v", d.DelaySwapSeqNo)

	builder.WriteByte('}')

	return builder.String()
}

// 投资单元
type CThostFtdcInvestUnitField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	// 投资单元代码
	InvestUnitID types.TThostFtdcInvestUnitIDType
	// 投资者单元名称
	InvestorUnitName types.TThostFtdcPartyNameType
	// 投资者分组代码
	InvestorGroupID types.TThostFtdcInvestorIDType
	// 手续费率模板代码
	CommModelID types.TThostFtdcInvestorIDType
	// 保证金率模板代码
	MarginModelID types.TThostFtdcInvestorIDType
	// 资金账号
	AccountID types.TThostFtdcAccountIDType
	// 币种代码
	CurrencyID types.TThostFtdcCurrencyIDType
}

func (d CThostFtdcInvestUnitField) Type() string { return "CThostFtdcInvestUnitField" }

func (d CThostFtdcInvestUnitField) String() string {
	var builder strings.Builder
	builder.Grow(219)

	builder.WriteString("CThostFtdcInvestUnitField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", InvestorID=%+v", d.InvestorID)
	fmt.Fprintf(&builder, ", InvestUnitID=%+v", d.InvestUnitID)
	fmt.Fprintf(&builder, ", InvestorUnitName=%+v", d.InvestorUnitName)
	fmt.Fprintf(&builder, ", InvestorGroupID=%+v", d.InvestorGroupID)
	fmt.Fprintf(&builder, ", CommModelID=%+v", d.CommModelID)
	fmt.Fprintf(&builder, ", MarginModelID=%+v", d.MarginModelID)
	fmt.Fprintf(&builder, ", AccountID=%+v", d.AccountID)
	fmt.Fprintf(&builder, ", CurrencyID=%+v", d.CurrencyID)

	builder.WriteByte('}')

	return builder.String()
}

// 查询投资单元
type CThostFtdcQryInvestUnitField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	// 投资单元代码
	InvestUnitID types.TThostFtdcInvestUnitIDType
}

func (d CThostFtdcQryInvestUnitField) Type() string { return "CThostFtdcQryInvestUnitField" }

func (d CThostFtdcQryInvestUnitField) String() string {
	var builder strings.Builder
	builder.Grow(88)

	builder.WriteString("CThostFtdcQryInvestUnitField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", InvestorID=%+v", d.InvestorID)
	fmt.Fprintf(&builder, ", InvestUnitID=%+v", d.InvestUnitID)

	builder.WriteByte('}')

	return builder.String()
}

// 二级代理商资金校验模式
type CThostFtdcSecAgentCheckModeField struct {
	// 投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 币种
	CurrencyID types.TThostFtdcCurrencyIDType
	// 境外中介机构资金帐号
	BrokerSecAgentID types.TThostFtdcAccountIDType
	// 是否需要校验自己的资金账户
	CheckSelfAccount types.TThostFtdcBoolType
}

func (d CThostFtdcSecAgentCheckModeField) Type() string { return "CThostFtdcSecAgentCheckModeField" }

func (d CThostFtdcSecAgentCheckModeField) String() string {
	var builder strings.Builder
	builder.Grow(142)

	builder.WriteString("CThostFtdcSecAgentCheckModeField{")
	fmt.Fprintf(&builder, "InvestorID=%+v", d.InvestorID)
	fmt.Fprintf(&builder, ", BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", CurrencyID=%+v", d.CurrencyID)
	fmt.Fprintf(&builder, ", BrokerSecAgentID=%+v", d.BrokerSecAgentID)
	fmt.Fprintf(&builder, ", CheckSelfAccount=%+v", d.CheckSelfAccount)

	builder.WriteByte('}')

	return builder.String()
}

// 二级代理商信息
type CThostFtdcSecAgentTradeInfoField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 境外中介机构资金帐号
	BrokerSecAgentID types.TThostFtdcAccountIDType
	// 投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	// 二级代理商姓名
	LongCustomerName types.TThostFtdcLongIndividualNameType
}

func (d CThostFtdcSecAgentTradeInfoField) Type() string { return "CThostFtdcSecAgentTradeInfoField" }

func (d CThostFtdcSecAgentTradeInfoField) String() string {
	var builder strings.Builder
	builder.Grow(122)

	builder.WriteString("CThostFtdcSecAgentTradeInfoField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", BrokerSecAgentID=%+v", d.BrokerSecAgentID)
	fmt.Fprintf(&builder, ", InvestorID=%+v", d.InvestorID)
	fmt.Fprintf(&builder, ", LongCustomerName=%+v", d.LongCustomerName)

	builder.WriteByte('}')

	return builder.String()
}

// 市场行情
type CThostFtdcMarketDataField struct {
	// 交易日
	TradingDay types.TThostFtdcDateType
	// 保留的无效字段
	reserve1 types.TThostFtdcOldInstrumentIDType
	// 交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	// 保留的无效字段
	reserve2 types.TThostFtdcOldExchangeInstIDType
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
	// 合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
	// 合约在交易所的代码
	ExchangeInstID types.TThostFtdcExchangeInstIDType
}

func (d CThostFtdcMarketDataField) Type() string { return "CThostFtdcMarketDataField" }

func (d CThostFtdcMarketDataField) String() string {
	var builder strings.Builder
	builder.Grow(555)

	builder.WriteString("CThostFtdcMarketDataField{")
	fmt.Fprintf(&builder, "TradingDay=%+v", d.TradingDay)
	fmt.Fprintf(&builder, ", reserve1=%+v", d.reserve1)
	fmt.Fprintf(&builder, ", ExchangeID=%+v", d.ExchangeID)
	fmt.Fprintf(&builder, ", reserve2=%+v", d.reserve2)
	fmt.Fprintf(&builder, ", LastPrice=%+v", d.LastPrice)
	fmt.Fprintf(&builder, ", PreSettlementPrice=%+v", d.PreSettlementPrice)
	fmt.Fprintf(&builder, ", PreClosePrice=%+v", d.PreClosePrice)
	fmt.Fprintf(&builder, ", PreOpenInterest=%+v", d.PreOpenInterest)
	fmt.Fprintf(&builder, ", OpenPrice=%+v", d.OpenPrice)
	fmt.Fprintf(&builder, ", HighestPrice=%+v", d.HighestPrice)
	fmt.Fprintf(&builder, ", LowestPrice=%+v", d.LowestPrice)
	fmt.Fprintf(&builder, ", Volume=%+v", d.Volume)
	fmt.Fprintf(&builder, ", Turnover=%+v", d.Turnover)
	fmt.Fprintf(&builder, ", OpenInterest=%+v", d.OpenInterest)
	fmt.Fprintf(&builder, ", ClosePrice=%+v", d.ClosePrice)
	fmt.Fprintf(&builder, ", SettlementPrice=%+v", d.SettlementPrice)
	fmt.Fprintf(&builder, ", UpperLimitPrice=%+v", d.UpperLimitPrice)
	fmt.Fprintf(&builder, ", LowerLimitPrice=%+v", d.LowerLimitPrice)
	fmt.Fprintf(&builder, ", PreDelta=%+v", d.PreDelta)
	fmt.Fprintf(&builder, ", CurrDelta=%+v", d.CurrDelta)
	fmt.Fprintf(&builder, ", UpdateTime=%+v", d.UpdateTime)
	fmt.Fprintf(&builder, ", UpdateMillisec=%+v", d.UpdateMillisec)
	fmt.Fprintf(&builder, ", ActionDay=%+v", d.ActionDay)
	fmt.Fprintf(&builder, ", InstrumentID=%+v", d.InstrumentID)
	fmt.Fprintf(&builder, ", ExchangeInstID=%+v", d.ExchangeInstID)

	builder.WriteByte('}')

	return builder.String()
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

func (d CThostFtdcMarketDataBaseField) Type() string { return "CThostFtdcMarketDataBaseField" }

func (d CThostFtdcMarketDataBaseField) String() string {
	var builder strings.Builder
	builder.Grow(143)

	builder.WriteString("CThostFtdcMarketDataBaseField{")
	fmt.Fprintf(&builder, "TradingDay=%+v", d.TradingDay)
	fmt.Fprintf(&builder, ", PreSettlementPrice=%+v", d.PreSettlementPrice)
	fmt.Fprintf(&builder, ", PreClosePrice=%+v", d.PreClosePrice)
	fmt.Fprintf(&builder, ", PreOpenInterest=%+v", d.PreOpenInterest)
	fmt.Fprintf(&builder, ", PreDelta=%+v", d.PreDelta)

	builder.WriteByte('}')

	return builder.String()
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

func (d CThostFtdcMarketDataStaticField) Type() string { return "CThostFtdcMarketDataStaticField" }

func (d CThostFtdcMarketDataStaticField) String() string {
	var builder strings.Builder
	builder.Grow(207)

	builder.WriteString("CThostFtdcMarketDataStaticField{")
	fmt.Fprintf(&builder, "OpenPrice=%+v", d.OpenPrice)
	fmt.Fprintf(&builder, ", HighestPrice=%+v", d.HighestPrice)
	fmt.Fprintf(&builder, ", LowestPrice=%+v", d.LowestPrice)
	fmt.Fprintf(&builder, ", ClosePrice=%+v", d.ClosePrice)
	fmt.Fprintf(&builder, ", UpperLimitPrice=%+v", d.UpperLimitPrice)
	fmt.Fprintf(&builder, ", LowerLimitPrice=%+v", d.LowerLimitPrice)
	fmt.Fprintf(&builder, ", SettlementPrice=%+v", d.SettlementPrice)
	fmt.Fprintf(&builder, ", CurrDelta=%+v", d.CurrDelta)

	builder.WriteByte('}')

	return builder.String()
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

func (d CThostFtdcMarketDataLastMatchField) Type() string {
	return "CThostFtdcMarketDataLastMatchField"
}

func (d CThostFtdcMarketDataLastMatchField) String() string {
	var builder strings.Builder
	builder.Grow(109)

	builder.WriteString("CThostFtdcMarketDataLastMatchField{")
	fmt.Fprintf(&builder, "LastPrice=%+v", d.LastPrice)
	fmt.Fprintf(&builder, ", Volume=%+v", d.Volume)
	fmt.Fprintf(&builder, ", Turnover=%+v", d.Turnover)
	fmt.Fprintf(&builder, ", OpenInterest=%+v", d.OpenInterest)

	builder.WriteByte('}')

	return builder.String()
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

func (d CThostFtdcMarketDataBestPriceField) Type() string {
	return "CThostFtdcMarketDataBestPriceField"
}

func (d CThostFtdcMarketDataBestPriceField) String() string {
	var builder strings.Builder
	builder.Grow(112)

	builder.WriteString("CThostFtdcMarketDataBestPriceField{")
	fmt.Fprintf(&builder, "BidPrice1=%+v", d.BidPrice1)
	fmt.Fprintf(&builder, ", BidVolume1=%+v", d.BidVolume1)
	fmt.Fprintf(&builder, ", AskPrice1=%+v", d.AskPrice1)
	fmt.Fprintf(&builder, ", AskVolume1=%+v", d.AskVolume1)

	builder.WriteByte('}')

	return builder.String()
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

func (d CThostFtdcMarketDataBid23Field) Type() string { return "CThostFtdcMarketDataBid23Field" }

func (d CThostFtdcMarketDataBid23Field) String() string {
	var builder strings.Builder
	builder.Grow(108)

	builder.WriteString("CThostFtdcMarketDataBid23Field{")
	fmt.Fprintf(&builder, "BidPrice2=%+v", d.BidPrice2)
	fmt.Fprintf(&builder, ", BidVolume2=%+v", d.BidVolume2)
	fmt.Fprintf(&builder, ", BidPrice3=%+v", d.BidPrice3)
	fmt.Fprintf(&builder, ", BidVolume3=%+v", d.BidVolume3)

	builder.WriteByte('}')

	return builder.String()
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

func (d CThostFtdcMarketDataAsk23Field) Type() string { return "CThostFtdcMarketDataAsk23Field" }

func (d CThostFtdcMarketDataAsk23Field) String() string {
	var builder strings.Builder
	builder.Grow(108)

	builder.WriteString("CThostFtdcMarketDataAsk23Field{")
	fmt.Fprintf(&builder, "AskPrice2=%+v", d.AskPrice2)
	fmt.Fprintf(&builder, ", AskVolume2=%+v", d.AskVolume2)
	fmt.Fprintf(&builder, ", AskPrice3=%+v", d.AskPrice3)
	fmt.Fprintf(&builder, ", AskVolume3=%+v", d.AskVolume3)

	builder.WriteByte('}')

	return builder.String()
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

func (d CThostFtdcMarketDataBid45Field) Type() string { return "CThostFtdcMarketDataBid45Field" }

func (d CThostFtdcMarketDataBid45Field) String() string {
	var builder strings.Builder
	builder.Grow(108)

	builder.WriteString("CThostFtdcMarketDataBid45Field{")
	fmt.Fprintf(&builder, "BidPrice4=%+v", d.BidPrice4)
	fmt.Fprintf(&builder, ", BidVolume4=%+v", d.BidVolume4)
	fmt.Fprintf(&builder, ", BidPrice5=%+v", d.BidPrice5)
	fmt.Fprintf(&builder, ", BidVolume5=%+v", d.BidVolume5)

	builder.WriteByte('}')

	return builder.String()
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

func (d CThostFtdcMarketDataAsk45Field) Type() string { return "CThostFtdcMarketDataAsk45Field" }

func (d CThostFtdcMarketDataAsk45Field) String() string {
	var builder strings.Builder
	builder.Grow(108)

	builder.WriteString("CThostFtdcMarketDataAsk45Field{")
	fmt.Fprintf(&builder, "AskPrice4=%+v", d.AskPrice4)
	fmt.Fprintf(&builder, ", AskVolume4=%+v", d.AskVolume4)
	fmt.Fprintf(&builder, ", AskPrice5=%+v", d.AskPrice5)
	fmt.Fprintf(&builder, ", AskVolume5=%+v", d.AskVolume5)

	builder.WriteByte('}')

	return builder.String()
}

// 行情更新时间属性
type CThostFtdcMarketDataUpdateTimeField struct {
	// 保留的无效字段
	reserve1 types.TThostFtdcOldInstrumentIDType
	// 最后修改时间
	UpdateTime types.TThostFtdcTimeType
	// 最后修改毫秒
	UpdateMillisec types.TThostFtdcMillisecType
	// 业务日期
	ActionDay types.TThostFtdcDateType
	// 合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
}

func (d CThostFtdcMarketDataUpdateTimeField) Type() string {
	return "CThostFtdcMarketDataUpdateTimeField"
}

func (d CThostFtdcMarketDataUpdateTimeField) String() string {
	var builder strings.Builder
	builder.Grow(138)

	builder.WriteString("CThostFtdcMarketDataUpdateTimeField{")
	fmt.Fprintf(&builder, "reserve1=%+v", d.reserve1)
	fmt.Fprintf(&builder, ", UpdateTime=%+v", d.UpdateTime)
	fmt.Fprintf(&builder, ", UpdateMillisec=%+v", d.UpdateMillisec)
	fmt.Fprintf(&builder, ", ActionDay=%+v", d.ActionDay)
	fmt.Fprintf(&builder, ", InstrumentID=%+v", d.InstrumentID)

	builder.WriteByte('}')

	return builder.String()
}

// 行情上下带价
type CThostFtdcMarketDataBandingPriceField struct {
	// 上带价
	BandingUpperPrice types.TThostFtdcPriceType
	// 下带价
	BandingLowerPrice types.TThostFtdcPriceType
}

func (d CThostFtdcMarketDataBandingPriceField) Type() string {
	return "CThostFtdcMarketDataBandingPriceField"
}

func (d CThostFtdcMarketDataBandingPriceField) String() string {
	var builder strings.Builder
	builder.Grow(91)

	builder.WriteString("CThostFtdcMarketDataBandingPriceField{")
	fmt.Fprintf(&builder, "BandingUpperPrice=%+v", d.BandingUpperPrice)
	fmt.Fprintf(&builder, ", BandingLowerPrice=%+v", d.BandingLowerPrice)

	builder.WriteByte('}')

	return builder.String()
}

// 行情交易所代码属性
type CThostFtdcMarketDataExchangeField struct {
	// 交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
}

func (d CThostFtdcMarketDataExchangeField) Type() string { return "CThostFtdcMarketDataExchangeField" }

func (d CThostFtdcMarketDataExchangeField) String() string {
	var builder strings.Builder
	builder.Grow(53)

	builder.WriteString("CThostFtdcMarketDataExchangeField{")
	fmt.Fprintf(&builder, "ExchangeID=%+v", d.ExchangeID)

	builder.WriteByte('}')

	return builder.String()
}

// 指定的合约
type CThostFtdcSpecificInstrumentField struct {
	// 保留的无效字段
	reserve1 types.TThostFtdcOldInstrumentIDType
	// 合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
}

func (d CThostFtdcSpecificInstrumentField) Type() string { return "CThostFtdcSpecificInstrumentField" }

func (d CThostFtdcSpecificInstrumentField) String() string {
	var builder strings.Builder
	builder.Grow(73)

	builder.WriteString("CThostFtdcSpecificInstrumentField{")
	fmt.Fprintf(&builder, "reserve1=%+v", d.reserve1)
	fmt.Fprintf(&builder, ", InstrumentID=%+v", d.InstrumentID)

	builder.WriteByte('}')

	return builder.String()
}

// 合约状态
type CThostFtdcInstrumentStatusField struct {
	// 交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	// 保留的无效字段
	reserve1 types.TThostFtdcOldExchangeInstIDType
	// 结算组代码
	SettlementGroupID types.TThostFtdcSettlementGroupIDType
	// 保留的无效字段
	reserve2 types.TThostFtdcOldInstrumentIDType
	// 合约交易状态
	InstrumentStatus types.TThostFtdcInstrumentStatusType
	// 交易阶段编号
	TradingSegmentSN types.TThostFtdcTradingSegmentSNType
	// 进入本状态时间
	EnterTime types.TThostFtdcTimeType
	// 进入本状态原因
	EnterReason types.TThostFtdcInstStatusEnterReasonType
	// 合约在交易所的代码
	ExchangeInstID types.TThostFtdcExchangeInstIDType
	// 合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
}

func (d CThostFtdcInstrumentStatusField) Type() string { return "CThostFtdcInstrumentStatusField" }

func (d CThostFtdcInstrumentStatusField) String() string {
	var builder strings.Builder
	builder.Grow(252)

	builder.WriteString("CThostFtdcInstrumentStatusField{")
	fmt.Fprintf(&builder, "ExchangeID=%+v", d.ExchangeID)
	fmt.Fprintf(&builder, ", reserve1=%+v", d.reserve1)
	fmt.Fprintf(&builder, ", SettlementGroupID=%+v", d.SettlementGroupID)
	fmt.Fprintf(&builder, ", reserve2=%+v", d.reserve2)
	fmt.Fprintf(&builder, ", InstrumentStatus=%+v", d.InstrumentStatus)
	fmt.Fprintf(&builder, ", TradingSegmentSN=%+v", d.TradingSegmentSN)
	fmt.Fprintf(&builder, ", EnterTime=%+v", d.EnterTime)
	fmt.Fprintf(&builder, ", EnterReason=%+v", d.EnterReason)
	fmt.Fprintf(&builder, ", ExchangeInstID=%+v", d.ExchangeInstID)
	fmt.Fprintf(&builder, ", InstrumentID=%+v", d.InstrumentID)

	builder.WriteByte('}')

	return builder.String()
}

// 查询合约状态
type CThostFtdcQryInstrumentStatusField struct {
	// 交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	// 保留的无效字段
	reserve1 types.TThostFtdcOldExchangeInstIDType
	// 合约在交易所的代码
	ExchangeInstID types.TThostFtdcExchangeInstIDType
}

func (d CThostFtdcQryInstrumentStatusField) Type() string {
	return "CThostFtdcQryInstrumentStatusField"
}

func (d CThostFtdcQryInstrumentStatusField) String() string {
	var builder strings.Builder
	builder.Grow(96)

	builder.WriteString("CThostFtdcQryInstrumentStatusField{")
	fmt.Fprintf(&builder, "ExchangeID=%+v", d.ExchangeID)
	fmt.Fprintf(&builder, ", reserve1=%+v", d.reserve1)
	fmt.Fprintf(&builder, ", ExchangeInstID=%+v", d.ExchangeInstID)

	builder.WriteByte('}')

	return builder.String()
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

func (d CThostFtdcInvestorAccountField) Type() string { return "CThostFtdcInvestorAccountField" }

func (d CThostFtdcInvestorAccountField) String() string {
	var builder strings.Builder
	builder.Grow(107)

	builder.WriteString("CThostFtdcInvestorAccountField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", InvestorID=%+v", d.InvestorID)
	fmt.Fprintf(&builder, ", AccountID=%+v", d.AccountID)
	fmt.Fprintf(&builder, ", CurrencyID=%+v", d.CurrencyID)

	builder.WriteByte('}')

	return builder.String()
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

func (d CThostFtdcPositionProfitAlgorithmField) Type() string {
	return "CThostFtdcPositionProfitAlgorithmField"
}

func (d CThostFtdcPositionProfitAlgorithmField) String() string {
	var builder strings.Builder
	builder.Grow(128)

	builder.WriteString("CThostFtdcPositionProfitAlgorithmField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", AccountID=%+v", d.AccountID)
	fmt.Fprintf(&builder, ", Algorithm=%+v", d.Algorithm)
	fmt.Fprintf(&builder, ", Memo=%+v", d.Memo)
	fmt.Fprintf(&builder, ", CurrencyID=%+v", d.CurrencyID)

	builder.WriteByte('}')

	return builder.String()
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

func (d CThostFtdcDiscountField) Type() string { return "CThostFtdcDiscountField" }

func (d CThostFtdcDiscountField) String() string {
	var builder strings.Builder
	builder.Grow(102)

	builder.WriteString("CThostFtdcDiscountField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", InvestorRange=%+v", d.InvestorRange)
	fmt.Fprintf(&builder, ", InvestorID=%+v", d.InvestorID)
	fmt.Fprintf(&builder, ", Discount=%+v", d.Discount)

	builder.WriteByte('}')

	return builder.String()
}

// 查询转帐银行
type CThostFtdcQryTransferBankField struct {
	// 银行代码
	BankID types.TThostFtdcBankIDType
	// 银行分中心代码
	BankBrchID types.TThostFtdcBankBrchIDType
}

func (d CThostFtdcQryTransferBankField) Type() string { return "CThostFtdcQryTransferBankField" }

func (d CThostFtdcQryTransferBankField) String() string {
	var builder strings.Builder
	builder.Grow(66)

	builder.WriteString("CThostFtdcQryTransferBankField{")
	fmt.Fprintf(&builder, "BankID=%+v", d.BankID)
	fmt.Fprintf(&builder, ", BankBrchID=%+v", d.BankBrchID)

	builder.WriteByte('}')

	return builder.String()
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

func (d CThostFtdcTransferBankField) Type() string { return "CThostFtdcTransferBankField" }

func (d CThostFtdcTransferBankField) String() string {
	var builder strings.Builder
	builder.Grow(99)

	builder.WriteString("CThostFtdcTransferBankField{")
	fmt.Fprintf(&builder, "BankID=%+v", d.BankID)
	fmt.Fprintf(&builder, ", BankBrchID=%+v", d.BankBrchID)
	fmt.Fprintf(&builder, ", BankName=%+v", d.BankName)
	fmt.Fprintf(&builder, ", IsActive=%+v", d.IsActive)

	builder.WriteByte('}')

	return builder.String()
}

// 查询投资者持仓明细
type CThostFtdcQryInvestorPositionDetailField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	// 保留的无效字段
	reserve1 types.TThostFtdcOldInstrumentIDType
	// 交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	// 投资单元代码
	InvestUnitID types.TThostFtdcInvestUnitIDType
	// 合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
}

func (d CThostFtdcQryInvestorPositionDetailField) Type() string {
	return "CThostFtdcQryInvestorPositionDetailField"
}

func (d CThostFtdcQryInvestorPositionDetailField) String() string {
	var builder strings.Builder
	builder.Grow(160)

	builder.WriteString("CThostFtdcQryInvestorPositionDetailField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", InvestorID=%+v", d.InvestorID)
	fmt.Fprintf(&builder, ", reserve1=%+v", d.reserve1)
	fmt.Fprintf(&builder, ", ExchangeID=%+v", d.ExchangeID)
	fmt.Fprintf(&builder, ", InvestUnitID=%+v", d.InvestUnitID)
	fmt.Fprintf(&builder, ", InstrumentID=%+v", d.InstrumentID)

	builder.WriteByte('}')

	return builder.String()
}

// 投资者持仓明细
type CThostFtdcInvestorPositionDetailField struct {
	// 保留的无效字段
	reserve1 types.TThostFtdcOldInstrumentIDType
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
	// 保留的无效字段
	reserve2 types.TThostFtdcOldInstrumentIDType
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
	// 先开先平剩余数量
	TimeFirstVolume types.TThostFtdcVolumeType
	// 投资单元代码
	InvestUnitID types.TThostFtdcInvestUnitIDType
	// 特殊持仓标志
	SpecPosiType types.TThostFtdcSpecPosiTypeType
	// 合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
	// 组合合约代码
	CombInstrumentID types.TThostFtdcInstrumentIDType
}

func (d CThostFtdcInvestorPositionDetailField) Type() string {
	return "CThostFtdcInvestorPositionDetailField"
}

func (d CThostFtdcInvestorPositionDetailField) String() string {
	var builder strings.Builder
	builder.Grow(720)

	builder.WriteString("CThostFtdcInvestorPositionDetailField{")
	fmt.Fprintf(&builder, "reserve1=%+v", d.reserve1)
	fmt.Fprintf(&builder, ", BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", InvestorID=%+v", d.InvestorID)
	fmt.Fprintf(&builder, ", HedgeFlag=%+v", d.HedgeFlag)
	fmt.Fprintf(&builder, ", Direction=%+v", d.Direction)
	fmt.Fprintf(&builder, ", OpenDate=%+v", d.OpenDate)
	fmt.Fprintf(&builder, ", TradeID=%+v", d.TradeID)
	fmt.Fprintf(&builder, ", Volume=%+v", d.Volume)
	fmt.Fprintf(&builder, ", OpenPrice=%+v", d.OpenPrice)
	fmt.Fprintf(&builder, ", TradingDay=%+v", d.TradingDay)
	fmt.Fprintf(&builder, ", SettlementID=%+v", d.SettlementID)
	fmt.Fprintf(&builder, ", TradeType=%+v", d.TradeType)
	fmt.Fprintf(&builder, ", reserve2=%+v", d.reserve2)
	fmt.Fprintf(&builder, ", ExchangeID=%+v", d.ExchangeID)
	fmt.Fprintf(&builder, ", CloseProfitByDate=%+v", d.CloseProfitByDate)
	fmt.Fprintf(&builder, ", CloseProfitByTrade=%+v", d.CloseProfitByTrade)
	fmt.Fprintf(&builder, ", PositionProfitByDate=%+v", d.PositionProfitByDate)
	fmt.Fprintf(&builder, ", PositionProfitByTrade=%+v", d.PositionProfitByTrade)
	fmt.Fprintf(&builder, ", Margin=%+v", d.Margin)
	fmt.Fprintf(&builder, ", ExchMargin=%+v", d.ExchMargin)
	fmt.Fprintf(&builder, ", MarginRateByMoney=%+v", d.MarginRateByMoney)
	fmt.Fprintf(&builder, ", MarginRateByVolume=%+v", d.MarginRateByVolume)
	fmt.Fprintf(&builder, ", LastSettlementPrice=%+v", d.LastSettlementPrice)
	fmt.Fprintf(&builder, ", SettlementPrice=%+v", d.SettlementPrice)
	fmt.Fprintf(&builder, ", CloseVolume=%+v", d.CloseVolume)
	fmt.Fprintf(&builder, ", CloseAmount=%+v", d.CloseAmount)
	fmt.Fprintf(&builder, ", TimeFirstVolume=%+v", d.TimeFirstVolume)
	fmt.Fprintf(&builder, ", InvestUnitID=%+v", d.InvestUnitID)
	fmt.Fprintf(&builder, ", SpecPosiType=%+v", d.SpecPosiType)
	fmt.Fprintf(&builder, ", InstrumentID=%+v", d.InstrumentID)
	fmt.Fprintf(&builder, ", CombInstrumentID=%+v", d.CombInstrumentID)

	builder.WriteByte('}')

	return builder.String()
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

func (d CThostFtdcTradingAccountPasswordField) Type() string {
	return "CThostFtdcTradingAccountPasswordField"
}

func (d CThostFtdcTradingAccountPasswordField) String() string {
	var builder strings.Builder
	builder.Grow(112)

	builder.WriteString("CThostFtdcTradingAccountPasswordField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", AccountID=%+v", d.AccountID)
	fmt.Fprintf(&builder, ", Password=%+v", d.Password)
	fmt.Fprintf(&builder, ", CurrencyID=%+v", d.CurrencyID)

	builder.WriteByte('}')

	return builder.String()
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
	// 撤单时选择席位算法
	OrderCancelAlg types.TThostFtdcOrderCancelAlgType
}

func (d CThostFtdcMDTraderOfferField) Type() string { return "CThostFtdcMDTraderOfferField" }

func (d CThostFtdcMDTraderOfferField) String() string {
	var builder strings.Builder
	builder.Grow(477)

	builder.WriteString("CThostFtdcMDTraderOfferField{")
	fmt.Fprintf(&builder, "ExchangeID=%+v", d.ExchangeID)
	fmt.Fprintf(&builder, ", TraderID=%+v", d.TraderID)
	fmt.Fprintf(&builder, ", ParticipantID=%+v", d.ParticipantID)
	fmt.Fprintf(&builder, ", Password=%+v", d.Password)
	fmt.Fprintf(&builder, ", InstallID=%+v", d.InstallID)
	fmt.Fprintf(&builder, ", OrderLocalID=%+v", d.OrderLocalID)
	fmt.Fprintf(&builder, ", TraderConnectStatus=%+v", d.TraderConnectStatus)
	fmt.Fprintf(&builder, ", ConnectRequestDate=%+v", d.ConnectRequestDate)
	fmt.Fprintf(&builder, ", ConnectRequestTime=%+v", d.ConnectRequestTime)
	fmt.Fprintf(&builder, ", LastReportDate=%+v", d.LastReportDate)
	fmt.Fprintf(&builder, ", LastReportTime=%+v", d.LastReportTime)
	fmt.Fprintf(&builder, ", ConnectDate=%+v", d.ConnectDate)
	fmt.Fprintf(&builder, ", ConnectTime=%+v", d.ConnectTime)
	fmt.Fprintf(&builder, ", StartDate=%+v", d.StartDate)
	fmt.Fprintf(&builder, ", StartTime=%+v", d.StartTime)
	fmt.Fprintf(&builder, ", TradingDay=%+v", d.TradingDay)
	fmt.Fprintf(&builder, ", BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", MaxTradeID=%+v", d.MaxTradeID)
	fmt.Fprintf(&builder, ", MaxOrderMessageReference=%+v", d.MaxOrderMessageReference)
	fmt.Fprintf(&builder, ", OrderCancelAlg=%+v", d.OrderCancelAlg)

	builder.WriteByte('}')

	return builder.String()
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

func (d CThostFtdcQryMDTraderOfferField) Type() string { return "CThostFtdcQryMDTraderOfferField" }

func (d CThostFtdcQryMDTraderOfferField) String() string {
	var builder strings.Builder
	builder.Grow(92)

	builder.WriteString("CThostFtdcQryMDTraderOfferField{")
	fmt.Fprintf(&builder, "ExchangeID=%+v", d.ExchangeID)
	fmt.Fprintf(&builder, ", ParticipantID=%+v", d.ParticipantID)
	fmt.Fprintf(&builder, ", TraderID=%+v", d.TraderID)

	builder.WriteByte('}')

	return builder.String()
}

// 查询客户通知
type CThostFtdcQryNoticeField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
}

func (d CThostFtdcQryNoticeField) Type() string { return "CThostFtdcQryNoticeField" }

func (d CThostFtdcQryNoticeField) String() string {
	var builder strings.Builder
	builder.Grow(42)

	builder.WriteString("CThostFtdcQryNoticeField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)

	builder.WriteByte('}')

	return builder.String()
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

func (d CThostFtdcNoticeField) Type() string { return "CThostFtdcNoticeField" }

func (d CThostFtdcNoticeField) String() string {
	var builder strings.Builder
	builder.Grow(79)

	builder.WriteString("CThostFtdcNoticeField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", Content=%+v", d.Content)
	fmt.Fprintf(&builder, ", SequenceLabel=%+v", d.SequenceLabel)

	builder.WriteByte('}')

	return builder.String()
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

func (d CThostFtdcUserRightField) Type() string { return "CThostFtdcUserRightField" }

func (d CThostFtdcUserRightField) String() string {
	var builder strings.Builder
	builder.Grow(102)

	builder.WriteString("CThostFtdcUserRightField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", UserID=%+v", d.UserID)
	fmt.Fprintf(&builder, ", UserRightType=%+v", d.UserRightType)
	fmt.Fprintf(&builder, ", IsForbidden=%+v", d.IsForbidden)

	builder.WriteByte('}')

	return builder.String()
}

// 查询结算信息确认域
type CThostFtdcQrySettlementInfoConfirmField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	// 投资者帐号
	AccountID types.TThostFtdcAccountIDType
	// 币种代码
	CurrencyID types.TThostFtdcCurrencyIDType
}

func (d CThostFtdcQrySettlementInfoConfirmField) Type() string {
	return "CThostFtdcQrySettlementInfoConfirmField"
}

func (d CThostFtdcQrySettlementInfoConfirmField) String() string {
	var builder strings.Builder
	builder.Grow(116)

	builder.WriteString("CThostFtdcQrySettlementInfoConfirmField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", InvestorID=%+v", d.InvestorID)
	fmt.Fprintf(&builder, ", AccountID=%+v", d.AccountID)
	fmt.Fprintf(&builder, ", CurrencyID=%+v", d.CurrencyID)

	builder.WriteByte('}')

	return builder.String()
}

// 装载结算信息
type CThostFtdcLoadSettlementInfoField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
}

func (d CThostFtdcLoadSettlementInfoField) Type() string { return "CThostFtdcLoadSettlementInfoField" }

func (d CThostFtdcLoadSettlementInfoField) String() string {
	var builder strings.Builder
	builder.Grow(51)

	builder.WriteString("CThostFtdcLoadSettlementInfoField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)

	builder.WriteByte('}')

	return builder.String()
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

func (d CThostFtdcBrokerWithdrawAlgorithmField) Type() string {
	return "CThostFtdcBrokerWithdrawAlgorithmField"
}

func (d CThostFtdcBrokerWithdrawAlgorithmField) String() string {
	var builder strings.Builder
	builder.Grow(289)

	builder.WriteString("CThostFtdcBrokerWithdrawAlgorithmField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", WithdrawAlgorithm=%+v", d.WithdrawAlgorithm)
	fmt.Fprintf(&builder, ", UsingRatio=%+v", d.UsingRatio)
	fmt.Fprintf(&builder, ", IncludeCloseProfit=%+v", d.IncludeCloseProfit)
	fmt.Fprintf(&builder, ", AllWithoutTrade=%+v", d.AllWithoutTrade)
	fmt.Fprintf(&builder, ", AvailIncludeCloseProfit=%+v", d.AvailIncludeCloseProfit)
	fmt.Fprintf(&builder, ", IsBrokerUserEvent=%+v", d.IsBrokerUserEvent)
	fmt.Fprintf(&builder, ", CurrencyID=%+v", d.CurrencyID)
	fmt.Fprintf(&builder, ", FundMortgageRatio=%+v", d.FundMortgageRatio)
	fmt.Fprintf(&builder, ", BalanceAlgorithm=%+v", d.BalanceAlgorithm)

	builder.WriteByte('}')

	return builder.String()
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

func (d CThostFtdcTradingAccountPasswordUpdateV1Field) Type() string {
	return "CThostFtdcTradingAccountPasswordUpdateV1Field"
}

func (d CThostFtdcTradingAccountPasswordUpdateV1Field) String() string {
	var builder strings.Builder
	builder.Grow(125)

	builder.WriteString("CThostFtdcTradingAccountPasswordUpdateV1Field{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", InvestorID=%+v", d.InvestorID)
	fmt.Fprintf(&builder, ", OldPassword=%+v", d.OldPassword)
	fmt.Fprintf(&builder, ", NewPassword=%+v", d.NewPassword)

	builder.WriteByte('}')

	return builder.String()
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

func (d CThostFtdcTradingAccountPasswordUpdateField) Type() string {
	return "CThostFtdcTradingAccountPasswordUpdateField"
}

func (d CThostFtdcTradingAccountPasswordUpdateField) String() string {
	var builder strings.Builder
	builder.Grow(142)

	builder.WriteString("CThostFtdcTradingAccountPasswordUpdateField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", AccountID=%+v", d.AccountID)
	fmt.Fprintf(&builder, ", OldPassword=%+v", d.OldPassword)
	fmt.Fprintf(&builder, ", NewPassword=%+v", d.NewPassword)
	fmt.Fprintf(&builder, ", CurrencyID=%+v", d.CurrencyID)

	builder.WriteByte('}')

	return builder.String()
}

// 查询组合合约分腿
type CThostFtdcQryCombinationLegField struct {
	// 保留的无效字段
	reserve1 types.TThostFtdcOldInstrumentIDType
	// 单腿编号
	LegID types.TThostFtdcLegIDType
	// 保留的无效字段
	reserve2 types.TThostFtdcOldInstrumentIDType
	// 组合合约代码
	CombInstrumentID types.TThostFtdcInstrumentIDType
	// 单腿合约代码
	LegInstrumentID types.TThostFtdcInstrumentIDType
}

func (d CThostFtdcQryCombinationLegField) Type() string { return "CThostFtdcQryCombinationLegField" }

func (d CThostFtdcQryCombinationLegField) String() string {
	var builder strings.Builder
	builder.Grow(134)

	builder.WriteString("CThostFtdcQryCombinationLegField{")
	fmt.Fprintf(&builder, "reserve1=%+v", d.reserve1)
	fmt.Fprintf(&builder, ", LegID=%+v", d.LegID)
	fmt.Fprintf(&builder, ", reserve2=%+v", d.reserve2)
	fmt.Fprintf(&builder, ", CombInstrumentID=%+v", d.CombInstrumentID)
	fmt.Fprintf(&builder, ", LegInstrumentID=%+v", d.LegInstrumentID)

	builder.WriteByte('}')

	return builder.String()
}

// 查询组合合约分腿
type CThostFtdcQrySyncStatusField struct {
	// 交易日
	TradingDay types.TThostFtdcDateType
}

func (d CThostFtdcQrySyncStatusField) Type() string { return "CThostFtdcQrySyncStatusField" }

func (d CThostFtdcQrySyncStatusField) String() string {
	var builder strings.Builder
	builder.Grow(48)

	builder.WriteString("CThostFtdcQrySyncStatusField{")
	fmt.Fprintf(&builder, "TradingDay=%+v", d.TradingDay)

	builder.WriteByte('}')

	return builder.String()
}

// 组合交易合约的单腿
type CThostFtdcCombinationLegField struct {
	// 保留的无效字段
	reserve1 types.TThostFtdcOldInstrumentIDType
	// 单腿编号
	LegID types.TThostFtdcLegIDType
	// 保留的无效字段
	reserve2 types.TThostFtdcOldInstrumentIDType
	// 买卖方向
	Direction types.TThostFtdcDirectionType
	// 单腿乘数
	LegMultiple types.TThostFtdcLegMultipleType
	// 派生层数
	ImplyLevel types.TThostFtdcImplyLevelType
	// 组合合约代码
	CombInstrumentID types.TThostFtdcInstrumentIDType
	// 单腿合约代码
	LegInstrumentID types.TThostFtdcInstrumentIDType
}

func (d CThostFtdcCombinationLegField) Type() string { return "CThostFtdcCombinationLegField" }

func (d CThostFtdcCombinationLegField) String() string {
	var builder strings.Builder
	builder.Grow(191)

	builder.WriteString("CThostFtdcCombinationLegField{")
	fmt.Fprintf(&builder, "reserve1=%+v", d.reserve1)
	fmt.Fprintf(&builder, ", LegID=%+v", d.LegID)
	fmt.Fprintf(&builder, ", reserve2=%+v", d.reserve2)
	fmt.Fprintf(&builder, ", Direction=%+v", d.Direction)
	fmt.Fprintf(&builder, ", LegMultiple=%+v", d.LegMultiple)
	fmt.Fprintf(&builder, ", ImplyLevel=%+v", d.ImplyLevel)
	fmt.Fprintf(&builder, ", CombInstrumentID=%+v", d.CombInstrumentID)
	fmt.Fprintf(&builder, ", LegInstrumentID=%+v", d.LegInstrumentID)

	builder.WriteByte('}')

	return builder.String()
}

// 数据同步状态
type CThostFtdcSyncStatusField struct {
	// 交易日
	TradingDay types.TThostFtdcDateType
	// 数据同步状态
	DataSyncStatus types.TThostFtdcDataSyncStatusType
}

func (d CThostFtdcSyncStatusField) Type() string { return "CThostFtdcSyncStatusField" }

func (d CThostFtdcSyncStatusField) String() string {
	var builder strings.Builder
	builder.Grow(69)

	builder.WriteString("CThostFtdcSyncStatusField{")
	fmt.Fprintf(&builder, "TradingDay=%+v", d.TradingDay)
	fmt.Fprintf(&builder, ", DataSyncStatus=%+v", d.DataSyncStatus)

	builder.WriteByte('}')

	return builder.String()
}

// 查询联系人
type CThostFtdcQryLinkManField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID types.TThostFtdcInvestorIDType
}

func (d CThostFtdcQryLinkManField) Type() string { return "CThostFtdcQryLinkManField" }

func (d CThostFtdcQryLinkManField) String() string {
	var builder strings.Builder
	builder.Grow(63)

	builder.WriteString("CThostFtdcQryLinkManField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", InvestorID=%+v", d.InvestorID)

	builder.WriteByte('}')

	return builder.String()
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

func (d CThostFtdcLinkManField) Type() string { return "CThostFtdcLinkManField" }

func (d CThostFtdcLinkManField) String() string {
	var builder strings.Builder
	builder.Grow(269)

	builder.WriteString("CThostFtdcLinkManField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", InvestorID=%+v", d.InvestorID)
	fmt.Fprintf(&builder, ", PersonType=%+v", d.PersonType)
	fmt.Fprintf(&builder, ", IdentifiedCardType=%+v", d.IdentifiedCardType)
	fmt.Fprintf(&builder, ", IdentifiedCardNo=%+v", d.IdentifiedCardNo)
	fmt.Fprintf(&builder, ", PersonName=%+v", d.PersonName)
	fmt.Fprintf(&builder, ", Telephone=%+v", d.Telephone)
	fmt.Fprintf(&builder, ", Address=%+v", d.Address)
	fmt.Fprintf(&builder, ", ZipCode=%+v", d.ZipCode)
	fmt.Fprintf(&builder, ", Priority=%+v", d.Priority)
	fmt.Fprintf(&builder, ", UOAZipCode=%+v", d.UOAZipCode)
	fmt.Fprintf(&builder, ", PersonFullName=%+v", d.PersonFullName)

	builder.WriteByte('}')

	return builder.String()
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

func (d CThostFtdcQryBrokerUserEventField) Type() string { return "CThostFtdcQryBrokerUserEventField" }

func (d CThostFtdcQryBrokerUserEventField) String() string {
	var builder strings.Builder
	builder.Grow(90)

	builder.WriteString("CThostFtdcQryBrokerUserEventField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", UserID=%+v", d.UserID)
	fmt.Fprintf(&builder, ", UserEventType=%+v", d.UserEventType)

	builder.WriteByte('}')

	return builder.String()
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
	// 保留的无效字段
	reserve1 types.TThostFtdcOldInstrumentIDType
	// 合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
	// 交易中心代码
	DRIdentityID types.TThostFtdcDRIdentityIDType
	// 交易日
	TradingDay types.TThostFtdcDateType
}

func (d CThostFtdcBrokerUserEventField) Type() string { return "CThostFtdcBrokerUserEventField" }

func (d CThostFtdcBrokerUserEventField) String() string {
	var builder strings.Builder
	builder.Grow(275)

	builder.WriteString("CThostFtdcBrokerUserEventField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", UserID=%+v", d.UserID)
	fmt.Fprintf(&builder, ", UserEventType=%+v", d.UserEventType)
	fmt.Fprintf(&builder, ", EventSequenceNo=%+v", d.EventSequenceNo)
	fmt.Fprintf(&builder, ", EventDate=%+v", d.EventDate)
	fmt.Fprintf(&builder, ", EventTime=%+v", d.EventTime)
	fmt.Fprintf(&builder, ", UserEventInfo=%+v", d.UserEventInfo)
	fmt.Fprintf(&builder, ", InvestorID=%+v", d.InvestorID)
	fmt.Fprintf(&builder, ", reserve1=%+v", d.reserve1)
	fmt.Fprintf(&builder, ", InstrumentID=%+v", d.InstrumentID)
	fmt.Fprintf(&builder, ", DRIdentityID=%+v", d.DRIdentityID)
	fmt.Fprintf(&builder, ", TradingDay=%+v", d.TradingDay)

	builder.WriteByte('}')

	return builder.String()
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

func (d CThostFtdcQryContractBankField) Type() string { return "CThostFtdcQryContractBankField" }

func (d CThostFtdcQryContractBankField) String() string {
	var builder strings.Builder
	builder.Grow(84)

	builder.WriteString("CThostFtdcQryContractBankField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", BankID=%+v", d.BankID)
	fmt.Fprintf(&builder, ", BankBrchID=%+v", d.BankBrchID)

	builder.WriteByte('}')

	return builder.String()
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
	// 上报csrc的银行代码
	csrcBankID types.TThostFtdcBankIDType
}

func (d CThostFtdcContractBankField) Type() string { return "CThostFtdcContractBankField" }

func (d CThostFtdcContractBankField) String() string {
	var builder strings.Builder
	builder.Grow(119)

	builder.WriteString("CThostFtdcContractBankField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", BankID=%+v", d.BankID)
	fmt.Fprintf(&builder, ", BankBrchID=%+v", d.BankBrchID)
	fmt.Fprintf(&builder, ", BankName=%+v", d.BankName)
	fmt.Fprintf(&builder, ", csrcBankID=%+v", d.csrcBankID)

	builder.WriteByte('}')

	return builder.String()
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
	// 保留的无效字段
	reserve1 types.TThostFtdcOldInstrumentIDType
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
	// 保留的无效字段
	reserve2 types.TThostFtdcOldInstrumentIDType
	// 成交组号
	TradeGroupID types.TThostFtdcTradeGroupIDType
	// 投资单元代码
	InvestUnitID types.TThostFtdcInvestUnitIDType
	// 合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
	// 组合持仓合约编码
	CombInstrumentID types.TThostFtdcInstrumentIDType
}

func (d CThostFtdcInvestorPositionCombineDetailField) Type() string {
	return "CThostFtdcInvestorPositionCombineDetailField"
}

func (d CThostFtdcInvestorPositionCombineDetailField) String() string {
	var builder strings.Builder
	builder.Grow(510)

	builder.WriteString("CThostFtdcInvestorPositionCombineDetailField{")
	fmt.Fprintf(&builder, "TradingDay=%+v", d.TradingDay)
	fmt.Fprintf(&builder, ", OpenDate=%+v", d.OpenDate)
	fmt.Fprintf(&builder, ", ExchangeID=%+v", d.ExchangeID)
	fmt.Fprintf(&builder, ", SettlementID=%+v", d.SettlementID)
	fmt.Fprintf(&builder, ", BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", InvestorID=%+v", d.InvestorID)
	fmt.Fprintf(&builder, ", ComTradeID=%+v", d.ComTradeID)
	fmt.Fprintf(&builder, ", TradeID=%+v", d.TradeID)
	fmt.Fprintf(&builder, ", reserve1=%+v", d.reserve1)
	fmt.Fprintf(&builder, ", HedgeFlag=%+v", d.HedgeFlag)
	fmt.Fprintf(&builder, ", Direction=%+v", d.Direction)
	fmt.Fprintf(&builder, ", TotalAmt=%+v", d.TotalAmt)
	fmt.Fprintf(&builder, ", Margin=%+v", d.Margin)
	fmt.Fprintf(&builder, ", ExchMargin=%+v", d.ExchMargin)
	fmt.Fprintf(&builder, ", MarginRateByMoney=%+v", d.MarginRateByMoney)
	fmt.Fprintf(&builder, ", MarginRateByVolume=%+v", d.MarginRateByVolume)
	fmt.Fprintf(&builder, ", LegID=%+v", d.LegID)
	fmt.Fprintf(&builder, ", LegMultiple=%+v", d.LegMultiple)
	fmt.Fprintf(&builder, ", reserve2=%+v", d.reserve2)
	fmt.Fprintf(&builder, ", TradeGroupID=%+v", d.TradeGroupID)
	fmt.Fprintf(&builder, ", InvestUnitID=%+v", d.InvestUnitID)
	fmt.Fprintf(&builder, ", InstrumentID=%+v", d.InstrumentID)
	fmt.Fprintf(&builder, ", CombInstrumentID=%+v", d.CombInstrumentID)

	builder.WriteByte('}')

	return builder.String()
}

// 预埋单
type CThostFtdcParkedOrderField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	// 保留的无效字段
	reserve1 types.TThostFtdcOldInstrumentIDType
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
	// 用户强平标志
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
	// 保留的无效字段
	reserve2 types.TThostFtdcOldIPAddressType
	// Mac地址
	MacAddress types.TThostFtdcMacAddressType
	// 合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
	// IP地址
	IPAddress types.TThostFtdcIPAddressType
}

func (d CThostFtdcParkedOrderField) Type() string { return "CThostFtdcParkedOrderField" }

func (d CThostFtdcParkedOrderField) String() string {
	var builder strings.Builder
	builder.Grow(792)

	builder.WriteString("CThostFtdcParkedOrderField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", InvestorID=%+v", d.InvestorID)
	fmt.Fprintf(&builder, ", reserve1=%+v", d.reserve1)
	fmt.Fprintf(&builder, ", OrderRef=%+v", d.OrderRef)
	fmt.Fprintf(&builder, ", UserID=%+v", d.UserID)
	fmt.Fprintf(&builder, ", OrderPriceType=%+v", d.OrderPriceType)
	fmt.Fprintf(&builder, ", Direction=%+v", d.Direction)
	fmt.Fprintf(&builder, ", CombOffsetFlag=%+v", d.CombOffsetFlag)
	fmt.Fprintf(&builder, ", CombHedgeFlag=%+v", d.CombHedgeFlag)
	fmt.Fprintf(&builder, ", LimitPrice=%+v", d.LimitPrice)
	fmt.Fprintf(&builder, ", VolumeTotalOriginal=%+v", d.VolumeTotalOriginal)
	fmt.Fprintf(&builder, ", TimeCondition=%+v", d.TimeCondition)
	fmt.Fprintf(&builder, ", GTDDate=%+v", d.GTDDate)
	fmt.Fprintf(&builder, ", VolumeCondition=%+v", d.VolumeCondition)
	fmt.Fprintf(&builder, ", MinVolume=%+v", d.MinVolume)
	fmt.Fprintf(&builder, ", ContingentCondition=%+v", d.ContingentCondition)
	fmt.Fprintf(&builder, ", StopPrice=%+v", d.StopPrice)
	fmt.Fprintf(&builder, ", ForceCloseReason=%+v", d.ForceCloseReason)
	fmt.Fprintf(&builder, ", IsAutoSuspend=%+v", d.IsAutoSuspend)
	fmt.Fprintf(&builder, ", BusinessUnit=%+v", d.BusinessUnit)
	fmt.Fprintf(&builder, ", RequestID=%+v", d.RequestID)
	fmt.Fprintf(&builder, ", UserForceClose=%+v", d.UserForceClose)
	fmt.Fprintf(&builder, ", ExchangeID=%+v", d.ExchangeID)
	fmt.Fprintf(&builder, ", ParkedOrderID=%+v", d.ParkedOrderID)
	fmt.Fprintf(&builder, ", UserType=%+v", d.UserType)
	fmt.Fprintf(&builder, ", Status=%+v", d.Status)
	fmt.Fprintf(&builder, ", ErrorID=%+v", d.ErrorID)
	fmt.Fprintf(&builder, ", ErrorMsg=%+v", d.ErrorMsg)
	fmt.Fprintf(&builder, ", IsSwapOrder=%+v", d.IsSwapOrder)
	fmt.Fprintf(&builder, ", AccountID=%+v", d.AccountID)
	fmt.Fprintf(&builder, ", CurrencyID=%+v", d.CurrencyID)
	fmt.Fprintf(&builder, ", ClientID=%+v", d.ClientID)
	fmt.Fprintf(&builder, ", InvestUnitID=%+v", d.InvestUnitID)
	fmt.Fprintf(&builder, ", reserve2=%+v", d.reserve2)
	fmt.Fprintf(&builder, ", MacAddress=%+v", d.MacAddress)
	fmt.Fprintf(&builder, ", InstrumentID=%+v", d.InstrumentID)
	fmt.Fprintf(&builder, ", IPAddress=%+v", d.IPAddress)

	builder.WriteByte('}')

	return builder.String()
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
	// 保留的无效字段
	reserve1 types.TThostFtdcOldInstrumentIDType
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
	// 保留的无效字段
	reserve2 types.TThostFtdcOldIPAddressType
	// Mac地址
	MacAddress types.TThostFtdcMacAddressType
	// 合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
	// IP地址
	IPAddress types.TThostFtdcIPAddressType
}

func (d CThostFtdcParkedOrderActionField) Type() string { return "CThostFtdcParkedOrderActionField" }

func (d CThostFtdcParkedOrderActionField) String() string {
	var builder strings.Builder
	builder.Grow(502)

	builder.WriteString("CThostFtdcParkedOrderActionField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", InvestorID=%+v", d.InvestorID)
	fmt.Fprintf(&builder, ", OrderActionRef=%+v", d.OrderActionRef)
	fmt.Fprintf(&builder, ", OrderRef=%+v", d.OrderRef)
	fmt.Fprintf(&builder, ", RequestID=%+v", d.RequestID)
	fmt.Fprintf(&builder, ", FrontID=%+v", d.FrontID)
	fmt.Fprintf(&builder, ", SessionID=%+v", d.SessionID)
	fmt.Fprintf(&builder, ", ExchangeID=%+v", d.ExchangeID)
	fmt.Fprintf(&builder, ", OrderSysID=%+v", d.OrderSysID)
	fmt.Fprintf(&builder, ", ActionFlag=%+v", d.ActionFlag)
	fmt.Fprintf(&builder, ", LimitPrice=%+v", d.LimitPrice)
	fmt.Fprintf(&builder, ", VolumeChange=%+v", d.VolumeChange)
	fmt.Fprintf(&builder, ", UserID=%+v", d.UserID)
	fmt.Fprintf(&builder, ", reserve1=%+v", d.reserve1)
	fmt.Fprintf(&builder, ", ParkedOrderActionID=%+v", d.ParkedOrderActionID)
	fmt.Fprintf(&builder, ", UserType=%+v", d.UserType)
	fmt.Fprintf(&builder, ", Status=%+v", d.Status)
	fmt.Fprintf(&builder, ", ErrorID=%+v", d.ErrorID)
	fmt.Fprintf(&builder, ", ErrorMsg=%+v", d.ErrorMsg)
	fmt.Fprintf(&builder, ", InvestUnitID=%+v", d.InvestUnitID)
	fmt.Fprintf(&builder, ", reserve2=%+v", d.reserve2)
	fmt.Fprintf(&builder, ", MacAddress=%+v", d.MacAddress)
	fmt.Fprintf(&builder, ", InstrumentID=%+v", d.InstrumentID)
	fmt.Fprintf(&builder, ", IPAddress=%+v", d.IPAddress)

	builder.WriteByte('}')

	return builder.String()
}

// 查询预埋单
type CThostFtdcQryParkedOrderField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	// 保留的无效字段
	reserve1 types.TThostFtdcOldInstrumentIDType
	// 交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	// 投资单元代码
	InvestUnitID types.TThostFtdcInvestUnitIDType
	// 合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
}

func (d CThostFtdcQryParkedOrderField) Type() string { return "CThostFtdcQryParkedOrderField" }

func (d CThostFtdcQryParkedOrderField) String() string {
	var builder strings.Builder
	builder.Grow(149)

	builder.WriteString("CThostFtdcQryParkedOrderField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", InvestorID=%+v", d.InvestorID)
	fmt.Fprintf(&builder, ", reserve1=%+v", d.reserve1)
	fmt.Fprintf(&builder, ", ExchangeID=%+v", d.ExchangeID)
	fmt.Fprintf(&builder, ", InvestUnitID=%+v", d.InvestUnitID)
	fmt.Fprintf(&builder, ", InstrumentID=%+v", d.InstrumentID)

	builder.WriteByte('}')

	return builder.String()
}

// 查询预埋撤单
type CThostFtdcQryParkedOrderActionField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	// 保留的无效字段
	reserve1 types.TThostFtdcOldInstrumentIDType
	// 交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	// 投资单元代码
	InvestUnitID types.TThostFtdcInvestUnitIDType
	// 合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
}

func (d CThostFtdcQryParkedOrderActionField) Type() string {
	return "CThostFtdcQryParkedOrderActionField"
}

func (d CThostFtdcQryParkedOrderActionField) String() string {
	var builder strings.Builder
	builder.Grow(155)

	builder.WriteString("CThostFtdcQryParkedOrderActionField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", InvestorID=%+v", d.InvestorID)
	fmt.Fprintf(&builder, ", reserve1=%+v", d.reserve1)
	fmt.Fprintf(&builder, ", ExchangeID=%+v", d.ExchangeID)
	fmt.Fprintf(&builder, ", InvestUnitID=%+v", d.InvestUnitID)
	fmt.Fprintf(&builder, ", InstrumentID=%+v", d.InstrumentID)

	builder.WriteByte('}')

	return builder.String()
}

// 删除预埋单
type CThostFtdcRemoveParkedOrderField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	// 预埋报单编号
	ParkedOrderID types.TThostFtdcParkedOrderIDType
	// 投资单元代码
	InvestUnitID types.TThostFtdcInvestUnitIDType
}

func (d CThostFtdcRemoveParkedOrderField) Type() string { return "CThostFtdcRemoveParkedOrderField" }

func (d CThostFtdcRemoveParkedOrderField) String() string {
	var builder strings.Builder
	builder.Grow(115)

	builder.WriteString("CThostFtdcRemoveParkedOrderField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", InvestorID=%+v", d.InvestorID)
	fmt.Fprintf(&builder, ", ParkedOrderID=%+v", d.ParkedOrderID)
	fmt.Fprintf(&builder, ", InvestUnitID=%+v", d.InvestUnitID)

	builder.WriteByte('}')

	return builder.String()
}

// 删除预埋撤单
type CThostFtdcRemoveParkedOrderActionField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	// 预埋撤单编号
	ParkedOrderActionID types.TThostFtdcParkedOrderActionIDType
	// 投资单元代码
	InvestUnitID types.TThostFtdcInvestUnitIDType
}

func (d CThostFtdcRemoveParkedOrderActionField) Type() string {
	return "CThostFtdcRemoveParkedOrderActionField"
}

func (d CThostFtdcRemoveParkedOrderActionField) String() string {
	var builder strings.Builder
	builder.Grow(127)

	builder.WriteString("CThostFtdcRemoveParkedOrderActionField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", InvestorID=%+v", d.InvestorID)
	fmt.Fprintf(&builder, ", ParkedOrderActionID=%+v", d.ParkedOrderActionID)
	fmt.Fprintf(&builder, ", InvestUnitID=%+v", d.InvestUnitID)

	builder.WriteByte('}')

	return builder.String()
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

func (d CThostFtdcInvestorWithdrawAlgorithmField) Type() string {
	return "CThostFtdcInvestorWithdrawAlgorithmField"
}

func (d CThostFtdcInvestorWithdrawAlgorithmField) String() string {
	var builder strings.Builder
	builder.Grow(168)

	builder.WriteString("CThostFtdcInvestorWithdrawAlgorithmField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", InvestorRange=%+v", d.InvestorRange)
	fmt.Fprintf(&builder, ", InvestorID=%+v", d.InvestorID)
	fmt.Fprintf(&builder, ", UsingRatio=%+v", d.UsingRatio)
	fmt.Fprintf(&builder, ", CurrencyID=%+v", d.CurrencyID)
	fmt.Fprintf(&builder, ", FundMortgageRatio=%+v", d.FundMortgageRatio)

	builder.WriteByte('}')

	return builder.String()
}

// 查询组合持仓明细
type CThostFtdcQryInvestorPositionCombineDetailField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	// 保留的无效字段
	reserve1 types.TThostFtdcOldInstrumentIDType
	// 交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	// 投资单元代码
	InvestUnitID types.TThostFtdcInvestUnitIDType
	// 组合持仓合约编码
	CombInstrumentID types.TThostFtdcInstrumentIDType
}

func (d CThostFtdcQryInvestorPositionCombineDetailField) Type() string {
	return "CThostFtdcQryInvestorPositionCombineDetailField"
}

func (d CThostFtdcQryInvestorPositionCombineDetailField) String() string {
	var builder strings.Builder
	builder.Grow(171)

	builder.WriteString("CThostFtdcQryInvestorPositionCombineDetailField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", InvestorID=%+v", d.InvestorID)
	fmt.Fprintf(&builder, ", reserve1=%+v", d.reserve1)
	fmt.Fprintf(&builder, ", ExchangeID=%+v", d.ExchangeID)
	fmt.Fprintf(&builder, ", InvestUnitID=%+v", d.InvestUnitID)
	fmt.Fprintf(&builder, ", CombInstrumentID=%+v", d.CombInstrumentID)

	builder.WriteByte('}')

	return builder.String()
}

// 成交均价
type CThostFtdcMarketDataAveragePriceField struct {
	// 当日均价
	AveragePrice types.TThostFtdcPriceType
}

func (d CThostFtdcMarketDataAveragePriceField) Type() string {
	return "CThostFtdcMarketDataAveragePriceField"
}

func (d CThostFtdcMarketDataAveragePriceField) String() string {
	var builder strings.Builder
	builder.Grow(59)

	builder.WriteString("CThostFtdcMarketDataAveragePriceField{")
	fmt.Fprintf(&builder, "AveragePrice=%+v", d.AveragePrice)

	builder.WriteByte('}')

	return builder.String()
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

func (d CThostFtdcVerifyInvestorPasswordField) Type() string {
	return "CThostFtdcVerifyInvestorPasswordField"
}

func (d CThostFtdcVerifyInvestorPasswordField) String() string {
	var builder strings.Builder
	builder.Grow(93)

	builder.WriteString("CThostFtdcVerifyInvestorPasswordField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", InvestorID=%+v", d.InvestorID)
	fmt.Fprintf(&builder, ", Password=%+v", d.Password)

	builder.WriteByte('}')

	return builder.String()
}

// 用户IP
type CThostFtdcUserIPField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 用户代码
	UserID types.TThostFtdcUserIDType
	// 保留的无效字段
	reserve1 types.TThostFtdcOldIPAddressType
	// 保留的无效字段
	reserve2 types.TThostFtdcOldIPAddressType
	// Mac地址
	MacAddress types.TThostFtdcMacAddressType
	// IP地址
	IPAddress types.TThostFtdcIPAddressType
	// IP地址掩码
	IPMask types.TThostFtdcIPAddressType
}

func (d CThostFtdcUserIPField) Type() string { return "CThostFtdcUserIPField" }

func (d CThostFtdcUserIPField) String() string {
	var builder strings.Builder
	builder.Grow(146)

	builder.WriteString("CThostFtdcUserIPField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", UserID=%+v", d.UserID)
	fmt.Fprintf(&builder, ", reserve1=%+v", d.reserve1)
	fmt.Fprintf(&builder, ", reserve2=%+v", d.reserve2)
	fmt.Fprintf(&builder, ", MacAddress=%+v", d.MacAddress)
	fmt.Fprintf(&builder, ", IPAddress=%+v", d.IPAddress)
	fmt.Fprintf(&builder, ", IPMask=%+v", d.IPMask)

	builder.WriteByte('}')

	return builder.String()
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
	// 投资单元代码
	InvestUnitID types.TThostFtdcInvestUnitIDType
}

func (d CThostFtdcTradingNoticeInfoField) Type() string { return "CThostFtdcTradingNoticeInfoField" }

func (d CThostFtdcTradingNoticeInfoField) String() string {
	var builder strings.Builder
	builder.Grow(176)

	builder.WriteString("CThostFtdcTradingNoticeInfoField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", InvestorID=%+v", d.InvestorID)
	fmt.Fprintf(&builder, ", SendTime=%+v", d.SendTime)
	fmt.Fprintf(&builder, ", FieldContent=%+v", d.FieldContent)
	fmt.Fprintf(&builder, ", SequenceSeries=%+v", d.SequenceSeries)
	fmt.Fprintf(&builder, ", SequenceNo=%+v", d.SequenceNo)
	fmt.Fprintf(&builder, ", InvestUnitID=%+v", d.InvestUnitID)

	builder.WriteByte('}')

	return builder.String()
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
	// 投资单元代码
	InvestUnitID types.TThostFtdcInvestUnitIDType
}

func (d CThostFtdcTradingNoticeField) Type() string { return "CThostFtdcTradingNoticeField" }

func (d CThostFtdcTradingNoticeField) String() string {
	var builder strings.Builder
	builder.Grow(211)

	builder.WriteString("CThostFtdcTradingNoticeField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", InvestorRange=%+v", d.InvestorRange)
	fmt.Fprintf(&builder, ", InvestorID=%+v", d.InvestorID)
	fmt.Fprintf(&builder, ", SequenceSeries=%+v", d.SequenceSeries)
	fmt.Fprintf(&builder, ", UserID=%+v", d.UserID)
	fmt.Fprintf(&builder, ", SendTime=%+v", d.SendTime)
	fmt.Fprintf(&builder, ", SequenceNo=%+v", d.SequenceNo)
	fmt.Fprintf(&builder, ", FieldContent=%+v", d.FieldContent)
	fmt.Fprintf(&builder, ", InvestUnitID=%+v", d.InvestUnitID)

	builder.WriteByte('}')

	return builder.String()
}

// 查询交易事件通知
type CThostFtdcQryTradingNoticeField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	// 投资单元代码
	InvestUnitID types.TThostFtdcInvestUnitIDType
}

func (d CThostFtdcQryTradingNoticeField) Type() string { return "CThostFtdcQryTradingNoticeField" }

func (d CThostFtdcQryTradingNoticeField) String() string {
	var builder strings.Builder
	builder.Grow(91)

	builder.WriteString("CThostFtdcQryTradingNoticeField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", InvestorID=%+v", d.InvestorID)
	fmt.Fprintf(&builder, ", InvestUnitID=%+v", d.InvestUnitID)

	builder.WriteByte('}')

	return builder.String()
}

// 查询错误报单
type CThostFtdcQryErrOrderField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID types.TThostFtdcInvestorIDType
}

func (d CThostFtdcQryErrOrderField) Type() string { return "CThostFtdcQryErrOrderField" }

func (d CThostFtdcQryErrOrderField) String() string {
	var builder strings.Builder
	builder.Grow(64)

	builder.WriteString("CThostFtdcQryErrOrderField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", InvestorID=%+v", d.InvestorID)

	builder.WriteByte('}')

	return builder.String()
}

// 错误报单
type CThostFtdcErrOrderField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	// 保留的无效字段
	reserve1 types.TThostFtdcOldInstrumentIDType
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
	// 用户强平标志
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
	// 保留的无效字段
	reserve2 types.TThostFtdcOldIPAddressType
	// Mac地址
	MacAddress types.TThostFtdcMacAddressType
	// 合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
	// IP地址
	IPAddress types.TThostFtdcIPAddressType
	// 报单回显字段
	OrderMemo types.TThostFtdcOrderMemoType
	// session上请求计数 api自动维护
	SessionReqSeq types.TThostFtdcSequenceNo12Type
}

func (d CThostFtdcErrOrderField) Type() string { return "CThostFtdcErrOrderField" }

func (d CThostFtdcErrOrderField) String() string {
	var builder strings.Builder
	builder.Grow(774)

	builder.WriteString("CThostFtdcErrOrderField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", InvestorID=%+v", d.InvestorID)
	fmt.Fprintf(&builder, ", reserve1=%+v", d.reserve1)
	fmt.Fprintf(&builder, ", OrderRef=%+v", d.OrderRef)
	fmt.Fprintf(&builder, ", UserID=%+v", d.UserID)
	fmt.Fprintf(&builder, ", OrderPriceType=%+v", d.OrderPriceType)
	fmt.Fprintf(&builder, ", Direction=%+v", d.Direction)
	fmt.Fprintf(&builder, ", CombOffsetFlag=%+v", d.CombOffsetFlag)
	fmt.Fprintf(&builder, ", CombHedgeFlag=%+v", d.CombHedgeFlag)
	fmt.Fprintf(&builder, ", LimitPrice=%+v", d.LimitPrice)
	fmt.Fprintf(&builder, ", VolumeTotalOriginal=%+v", d.VolumeTotalOriginal)
	fmt.Fprintf(&builder, ", TimeCondition=%+v", d.TimeCondition)
	fmt.Fprintf(&builder, ", GTDDate=%+v", d.GTDDate)
	fmt.Fprintf(&builder, ", VolumeCondition=%+v", d.VolumeCondition)
	fmt.Fprintf(&builder, ", MinVolume=%+v", d.MinVolume)
	fmt.Fprintf(&builder, ", ContingentCondition=%+v", d.ContingentCondition)
	fmt.Fprintf(&builder, ", StopPrice=%+v", d.StopPrice)
	fmt.Fprintf(&builder, ", ForceCloseReason=%+v", d.ForceCloseReason)
	fmt.Fprintf(&builder, ", IsAutoSuspend=%+v", d.IsAutoSuspend)
	fmt.Fprintf(&builder, ", BusinessUnit=%+v", d.BusinessUnit)
	fmt.Fprintf(&builder, ", RequestID=%+v", d.RequestID)
	fmt.Fprintf(&builder, ", UserForceClose=%+v", d.UserForceClose)
	fmt.Fprintf(&builder, ", ErrorID=%+v", d.ErrorID)
	fmt.Fprintf(&builder, ", ErrorMsg=%+v", d.ErrorMsg)
	fmt.Fprintf(&builder, ", IsSwapOrder=%+v", d.IsSwapOrder)
	fmt.Fprintf(&builder, ", ExchangeID=%+v", d.ExchangeID)
	fmt.Fprintf(&builder, ", InvestUnitID=%+v", d.InvestUnitID)
	fmt.Fprintf(&builder, ", AccountID=%+v", d.AccountID)
	fmt.Fprintf(&builder, ", CurrencyID=%+v", d.CurrencyID)
	fmt.Fprintf(&builder, ", ClientID=%+v", d.ClientID)
	fmt.Fprintf(&builder, ", reserve2=%+v", d.reserve2)
	fmt.Fprintf(&builder, ", MacAddress=%+v", d.MacAddress)
	fmt.Fprintf(&builder, ", InstrumentID=%+v", d.InstrumentID)
	fmt.Fprintf(&builder, ", IPAddress=%+v", d.IPAddress)
	fmt.Fprintf(&builder, ", OrderMemo=%+v", d.OrderMemo)
	fmt.Fprintf(&builder, ", SessionReqSeq=%+v", d.SessionReqSeq)

	builder.WriteByte('}')

	return builder.String()
}

// 查询错误报单操作
type CThostFtdcErrorConditionalOrderField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	// 保留的无效字段
	reserve1 types.TThostFtdcOldInstrumentIDType
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
	// 保留的无效字段
	reserve2 types.TThostFtdcOldExchangeInstIDType
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
	// 用户强平标志
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
	// 保留的无效字段
	reserve3 types.TThostFtdcOldIPAddressType
	// Mac地址
	MacAddress types.TThostFtdcMacAddressType
	// 合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
	// 合约在交易所的代码
	ExchangeInstID types.TThostFtdcExchangeInstIDType
	// IP地址
	IPAddress types.TThostFtdcIPAddressType
}

func (d CThostFtdcErrorConditionalOrderField) Type() string {
	return "CThostFtdcErrorConditionalOrderField"
}

func (d CThostFtdcErrorConditionalOrderField) String() string {
	var builder strings.Builder
	builder.Grow(1477)

	builder.WriteString("CThostFtdcErrorConditionalOrderField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", InvestorID=%+v", d.InvestorID)
	fmt.Fprintf(&builder, ", reserve1=%+v", d.reserve1)
	fmt.Fprintf(&builder, ", OrderRef=%+v", d.OrderRef)
	fmt.Fprintf(&builder, ", UserID=%+v", d.UserID)
	fmt.Fprintf(&builder, ", OrderPriceType=%+v", d.OrderPriceType)
	fmt.Fprintf(&builder, ", Direction=%+v", d.Direction)
	fmt.Fprintf(&builder, ", CombOffsetFlag=%+v", d.CombOffsetFlag)
	fmt.Fprintf(&builder, ", CombHedgeFlag=%+v", d.CombHedgeFlag)
	fmt.Fprintf(&builder, ", LimitPrice=%+v", d.LimitPrice)
	fmt.Fprintf(&builder, ", VolumeTotalOriginal=%+v", d.VolumeTotalOriginal)
	fmt.Fprintf(&builder, ", TimeCondition=%+v", d.TimeCondition)
	fmt.Fprintf(&builder, ", GTDDate=%+v", d.GTDDate)
	fmt.Fprintf(&builder, ", VolumeCondition=%+v", d.VolumeCondition)
	fmt.Fprintf(&builder, ", MinVolume=%+v", d.MinVolume)
	fmt.Fprintf(&builder, ", ContingentCondition=%+v", d.ContingentCondition)
	fmt.Fprintf(&builder, ", StopPrice=%+v", d.StopPrice)
	fmt.Fprintf(&builder, ", ForceCloseReason=%+v", d.ForceCloseReason)
	fmt.Fprintf(&builder, ", IsAutoSuspend=%+v", d.IsAutoSuspend)
	fmt.Fprintf(&builder, ", BusinessUnit=%+v", d.BusinessUnit)
	fmt.Fprintf(&builder, ", RequestID=%+v", d.RequestID)
	fmt.Fprintf(&builder, ", OrderLocalID=%+v", d.OrderLocalID)
	fmt.Fprintf(&builder, ", ExchangeID=%+v", d.ExchangeID)
	fmt.Fprintf(&builder, ", ParticipantID=%+v", d.ParticipantID)
	fmt.Fprintf(&builder, ", ClientID=%+v", d.ClientID)
	fmt.Fprintf(&builder, ", reserve2=%+v", d.reserve2)
	fmt.Fprintf(&builder, ", TraderID=%+v", d.TraderID)
	fmt.Fprintf(&builder, ", InstallID=%+v", d.InstallID)
	fmt.Fprintf(&builder, ", OrderSubmitStatus=%+v", d.OrderSubmitStatus)
	fmt.Fprintf(&builder, ", NotifySequence=%+v", d.NotifySequence)
	fmt.Fprintf(&builder, ", TradingDay=%+v", d.TradingDay)
	fmt.Fprintf(&builder, ", SettlementID=%+v", d.SettlementID)
	fmt.Fprintf(&builder, ", OrderSysID=%+v", d.OrderSysID)
	fmt.Fprintf(&builder, ", OrderSource=%+v", d.OrderSource)
	fmt.Fprintf(&builder, ", OrderStatus=%+v", d.OrderStatus)
	fmt.Fprintf(&builder, ", OrderType=%+v", d.OrderType)
	fmt.Fprintf(&builder, ", VolumeTraded=%+v", d.VolumeTraded)
	fmt.Fprintf(&builder, ", VolumeTotal=%+v", d.VolumeTotal)
	fmt.Fprintf(&builder, ", InsertDate=%+v", d.InsertDate)
	fmt.Fprintf(&builder, ", InsertTime=%+v", d.InsertTime)
	fmt.Fprintf(&builder, ", ActiveTime=%+v", d.ActiveTime)
	fmt.Fprintf(&builder, ", SuspendTime=%+v", d.SuspendTime)
	fmt.Fprintf(&builder, ", UpdateTime=%+v", d.UpdateTime)
	fmt.Fprintf(&builder, ", CancelTime=%+v", d.CancelTime)
	fmt.Fprintf(&builder, ", ActiveTraderID=%+v", d.ActiveTraderID)
	fmt.Fprintf(&builder, ", ClearingPartID=%+v", d.ClearingPartID)
	fmt.Fprintf(&builder, ", SequenceNo=%+v", d.SequenceNo)
	fmt.Fprintf(&builder, ", FrontID=%+v", d.FrontID)
	fmt.Fprintf(&builder, ", SessionID=%+v", d.SessionID)
	fmt.Fprintf(&builder, ", UserProductInfo=%+v", d.UserProductInfo)
	fmt.Fprintf(&builder, ", StatusMsg=%+v", d.StatusMsg)
	fmt.Fprintf(&builder, ", UserForceClose=%+v", d.UserForceClose)
	fmt.Fprintf(&builder, ", ActiveUserID=%+v", d.ActiveUserID)
	fmt.Fprintf(&builder, ", BrokerOrderSeq=%+v", d.BrokerOrderSeq)
	fmt.Fprintf(&builder, ", RelativeOrderSysID=%+v", d.RelativeOrderSysID)
	fmt.Fprintf(&builder, ", ZCETotalTradedVolume=%+v", d.ZCETotalTradedVolume)
	fmt.Fprintf(&builder, ", ErrorID=%+v", d.ErrorID)
	fmt.Fprintf(&builder, ", ErrorMsg=%+v", d.ErrorMsg)
	fmt.Fprintf(&builder, ", IsSwapOrder=%+v", d.IsSwapOrder)
	fmt.Fprintf(&builder, ", BranchID=%+v", d.BranchID)
	fmt.Fprintf(&builder, ", InvestUnitID=%+v", d.InvestUnitID)
	fmt.Fprintf(&builder, ", AccountID=%+v", d.AccountID)
	fmt.Fprintf(&builder, ", CurrencyID=%+v", d.CurrencyID)
	fmt.Fprintf(&builder, ", reserve3=%+v", d.reserve3)
	fmt.Fprintf(&builder, ", MacAddress=%+v", d.MacAddress)
	fmt.Fprintf(&builder, ", InstrumentID=%+v", d.InstrumentID)
	fmt.Fprintf(&builder, ", ExchangeInstID=%+v", d.ExchangeInstID)
	fmt.Fprintf(&builder, ", IPAddress=%+v", d.IPAddress)

	builder.WriteByte('}')

	return builder.String()
}

// 查询错误报单操作
type CThostFtdcQryErrOrderActionField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID types.TThostFtdcInvestorIDType
}

func (d CThostFtdcQryErrOrderActionField) Type() string { return "CThostFtdcQryErrOrderActionField" }

func (d CThostFtdcQryErrOrderActionField) String() string {
	var builder strings.Builder
	builder.Grow(70)

	builder.WriteString("CThostFtdcQryErrOrderActionField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", InvestorID=%+v", d.InvestorID)

	builder.WriteByte('}')

	return builder.String()
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
	// 保留的无效字段
	reserve1 types.TThostFtdcOldInstrumentIDType
	// 营业部编号
	BranchID types.TThostFtdcBranchIDType
	// 投资单元代码
	InvestUnitID types.TThostFtdcInvestUnitIDType
	// 保留的无效字段
	reserve2 types.TThostFtdcOldIPAddressType
	// Mac地址
	MacAddress types.TThostFtdcMacAddressType
	// 错误代码
	ErrorID types.TThostFtdcErrorIDType
	// 错误信息
	ErrorMsg types.TThostFtdcErrorMsgType
	// 合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
	// IP地址
	IPAddress types.TThostFtdcIPAddressType
	// 报单回显字段
	OrderMemo types.TThostFtdcOrderMemoType
	// session上请求计数 api自动维护
	SessionReqSeq types.TThostFtdcSequenceNo12Type
}

func (d CThostFtdcErrOrderActionField) Type() string { return "CThostFtdcErrOrderActionField" }

func (d CThostFtdcErrOrderActionField) String() string {
	var builder strings.Builder
	builder.Grow(727)

	builder.WriteString("CThostFtdcErrOrderActionField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", InvestorID=%+v", d.InvestorID)
	fmt.Fprintf(&builder, ", OrderActionRef=%+v", d.OrderActionRef)
	fmt.Fprintf(&builder, ", OrderRef=%+v", d.OrderRef)
	fmt.Fprintf(&builder, ", RequestID=%+v", d.RequestID)
	fmt.Fprintf(&builder, ", FrontID=%+v", d.FrontID)
	fmt.Fprintf(&builder, ", SessionID=%+v", d.SessionID)
	fmt.Fprintf(&builder, ", ExchangeID=%+v", d.ExchangeID)
	fmt.Fprintf(&builder, ", OrderSysID=%+v", d.OrderSysID)
	fmt.Fprintf(&builder, ", ActionFlag=%+v", d.ActionFlag)
	fmt.Fprintf(&builder, ", LimitPrice=%+v", d.LimitPrice)
	fmt.Fprintf(&builder, ", VolumeChange=%+v", d.VolumeChange)
	fmt.Fprintf(&builder, ", ActionDate=%+v", d.ActionDate)
	fmt.Fprintf(&builder, ", ActionTime=%+v", d.ActionTime)
	fmt.Fprintf(&builder, ", TraderID=%+v", d.TraderID)
	fmt.Fprintf(&builder, ", InstallID=%+v", d.InstallID)
	fmt.Fprintf(&builder, ", OrderLocalID=%+v", d.OrderLocalID)
	fmt.Fprintf(&builder, ", ActionLocalID=%+v", d.ActionLocalID)
	fmt.Fprintf(&builder, ", ParticipantID=%+v", d.ParticipantID)
	fmt.Fprintf(&builder, ", ClientID=%+v", d.ClientID)
	fmt.Fprintf(&builder, ", BusinessUnit=%+v", d.BusinessUnit)
	fmt.Fprintf(&builder, ", OrderActionStatus=%+v", d.OrderActionStatus)
	fmt.Fprintf(&builder, ", UserID=%+v", d.UserID)
	fmt.Fprintf(&builder, ", StatusMsg=%+v", d.StatusMsg)
	fmt.Fprintf(&builder, ", reserve1=%+v", d.reserve1)
	fmt.Fprintf(&builder, ", BranchID=%+v", d.BranchID)
	fmt.Fprintf(&builder, ", InvestUnitID=%+v", d.InvestUnitID)
	fmt.Fprintf(&builder, ", reserve2=%+v", d.reserve2)
	fmt.Fprintf(&builder, ", MacAddress=%+v", d.MacAddress)
	fmt.Fprintf(&builder, ", ErrorID=%+v", d.ErrorID)
	fmt.Fprintf(&builder, ", ErrorMsg=%+v", d.ErrorMsg)
	fmt.Fprintf(&builder, ", InstrumentID=%+v", d.InstrumentID)
	fmt.Fprintf(&builder, ", IPAddress=%+v", d.IPAddress)
	fmt.Fprintf(&builder, ", OrderMemo=%+v", d.OrderMemo)
	fmt.Fprintf(&builder, ", SessionReqSeq=%+v", d.SessionReqSeq)

	builder.WriteByte('}')

	return builder.String()
}

// 查询交易所状态
type CThostFtdcQryExchangeSequenceField struct {
	// 交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
}

func (d CThostFtdcQryExchangeSequenceField) Type() string {
	return "CThostFtdcQryExchangeSequenceField"
}

func (d CThostFtdcQryExchangeSequenceField) String() string {
	var builder strings.Builder
	builder.Grow(54)

	builder.WriteString("CThostFtdcQryExchangeSequenceField{")
	fmt.Fprintf(&builder, "ExchangeID=%+v", d.ExchangeID)

	builder.WriteByte('}')

	return builder.String()
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

func (d CThostFtdcExchangeSequenceField) Type() string { return "CThostFtdcExchangeSequenceField" }

func (d CThostFtdcExchangeSequenceField) String() string {
	var builder strings.Builder
	builder.Grow(93)

	builder.WriteString("CThostFtdcExchangeSequenceField{")
	fmt.Fprintf(&builder, "ExchangeID=%+v", d.ExchangeID)
	fmt.Fprintf(&builder, ", SequenceNo=%+v", d.SequenceNo)
	fmt.Fprintf(&builder, ", MarketStatus=%+v", d.MarketStatus)

	builder.WriteByte('}')

	return builder.String()
}

// 根据价格查询最大报单数量
type CThostFtdcQryMaxOrderVolumeWithPriceField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	// 保留的无效字段
	reserve1 types.TThostFtdcOldInstrumentIDType
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
	// 交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	// 投资单元代码
	InvestUnitID types.TThostFtdcInvestUnitIDType
	// 合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
}

func (d CThostFtdcQryMaxOrderVolumeWithPriceField) Type() string {
	return "CThostFtdcQryMaxOrderVolumeWithPriceField"
}

func (d CThostFtdcQryMaxOrderVolumeWithPriceField) String() string {
	var builder strings.Builder
	builder.Grow(253)

	builder.WriteString("CThostFtdcQryMaxOrderVolumeWithPriceField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", InvestorID=%+v", d.InvestorID)
	fmt.Fprintf(&builder, ", reserve1=%+v", d.reserve1)
	fmt.Fprintf(&builder, ", Direction=%+v", d.Direction)
	fmt.Fprintf(&builder, ", OffsetFlag=%+v", d.OffsetFlag)
	fmt.Fprintf(&builder, ", HedgeFlag=%+v", d.HedgeFlag)
	fmt.Fprintf(&builder, ", MaxVolume=%+v", d.MaxVolume)
	fmt.Fprintf(&builder, ", Price=%+v", d.Price)
	fmt.Fprintf(&builder, ", ExchangeID=%+v", d.ExchangeID)
	fmt.Fprintf(&builder, ", InvestUnitID=%+v", d.InvestUnitID)
	fmt.Fprintf(&builder, ", InstrumentID=%+v", d.InstrumentID)

	builder.WriteByte('}')

	return builder.String()
}

// 查询经纪公司交易参数
type CThostFtdcQryBrokerTradingParamsField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	// 币种代码
	CurrencyID types.TThostFtdcCurrencyIDType
	// 投资者帐号
	AccountID types.TThostFtdcAccountIDType
}

func (d CThostFtdcQryBrokerTradingParamsField) Type() string {
	return "CThostFtdcQryBrokerTradingParamsField"
}

func (d CThostFtdcQryBrokerTradingParamsField) String() string {
	var builder strings.Builder
	builder.Grow(114)

	builder.WriteString("CThostFtdcQryBrokerTradingParamsField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", InvestorID=%+v", d.InvestorID)
	fmt.Fprintf(&builder, ", CurrencyID=%+v", d.CurrencyID)
	fmt.Fprintf(&builder, ", AccountID=%+v", d.AccountID)

	builder.WriteByte('}')

	return builder.String()
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
	// 投资者帐号
	AccountID types.TThostFtdcAccountIDType
}

func (d CThostFtdcBrokerTradingParamsField) Type() string {
	return "CThostFtdcBrokerTradingParamsField"
}

func (d CThostFtdcBrokerTradingParamsField) String() string {
	var builder strings.Builder
	builder.Grow(220)

	builder.WriteString("CThostFtdcBrokerTradingParamsField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", InvestorID=%+v", d.InvestorID)
	fmt.Fprintf(&builder, ", MarginPriceType=%+v", d.MarginPriceType)
	fmt.Fprintf(&builder, ", Algorithm=%+v", d.Algorithm)
	fmt.Fprintf(&builder, ", AvailIncludeCloseProfit=%+v", d.AvailIncludeCloseProfit)
	fmt.Fprintf(&builder, ", CurrencyID=%+v", d.CurrencyID)
	fmt.Fprintf(&builder, ", OptionRoyaltyPriceType=%+v", d.OptionRoyaltyPriceType)
	fmt.Fprintf(&builder, ", AccountID=%+v", d.AccountID)

	builder.WriteByte('}')

	return builder.String()
}

// 查询经纪公司交易算法
type CThostFtdcQryBrokerTradingAlgosField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	// 保留的无效字段
	reserve1 types.TThostFtdcOldInstrumentIDType
	// 合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
}

func (d CThostFtdcQryBrokerTradingAlgosField) Type() string {
	return "CThostFtdcQryBrokerTradingAlgosField"
}

func (d CThostFtdcQryBrokerTradingAlgosField) String() string {
	var builder strings.Builder
	builder.Grow(114)

	builder.WriteString("CThostFtdcQryBrokerTradingAlgosField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", ExchangeID=%+v", d.ExchangeID)
	fmt.Fprintf(&builder, ", reserve1=%+v", d.reserve1)
	fmt.Fprintf(&builder, ", InstrumentID=%+v", d.InstrumentID)

	builder.WriteByte('}')

	return builder.String()
}

// 经纪公司交易算法
type CThostFtdcBrokerTradingAlgosField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	// 保留的无效字段
	reserve1 types.TThostFtdcOldInstrumentIDType
	// 持仓处理算法编号
	HandlePositionAlgoID types.TThostFtdcHandlePositionAlgoIDType
	// 寻找保证金率算法编号
	FindMarginRateAlgoID types.TThostFtdcFindMarginRateAlgoIDType
	// 资金处理算法编号
	HandleTradingAccountAlgoID types.TThostFtdcHandleTradingAccountAlgoIDType
	// 合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
}

func (d CThostFtdcBrokerTradingAlgosField) Type() string { return "CThostFtdcBrokerTradingAlgosField" }

func (d CThostFtdcBrokerTradingAlgosField) String() string {
	var builder strings.Builder
	builder.Grow(207)

	builder.WriteString("CThostFtdcBrokerTradingAlgosField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", ExchangeID=%+v", d.ExchangeID)
	fmt.Fprintf(&builder, ", reserve1=%+v", d.reserve1)
	fmt.Fprintf(&builder, ", HandlePositionAlgoID=%+v", d.HandlePositionAlgoID)
	fmt.Fprintf(&builder, ", FindMarginRateAlgoID=%+v", d.FindMarginRateAlgoID)
	fmt.Fprintf(&builder, ", HandleTradingAccountAlgoID=%+v", d.HandleTradingAccountAlgoID)
	fmt.Fprintf(&builder, ", InstrumentID=%+v", d.InstrumentID)

	builder.WriteByte('}')

	return builder.String()
}

// 查询经纪公司资金
type CThostFtdcQueryBrokerDepositField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
}

func (d CThostFtdcQueryBrokerDepositField) Type() string { return "CThostFtdcQueryBrokerDepositField" }

func (d CThostFtdcQueryBrokerDepositField) String() string {
	var builder strings.Builder
	builder.Grow(71)

	builder.WriteString("CThostFtdcQueryBrokerDepositField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", ExchangeID=%+v", d.ExchangeID)

	builder.WriteByte('}')

	return builder.String()
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

func (d CThostFtdcBrokerDepositField) Type() string { return "CThostFtdcBrokerDepositField" }

func (d CThostFtdcBrokerDepositField) String() string {
	var builder strings.Builder
	builder.Grow(280)

	builder.WriteString("CThostFtdcBrokerDepositField{")
	fmt.Fprintf(&builder, "TradingDay=%+v", d.TradingDay)
	fmt.Fprintf(&builder, ", BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", ParticipantID=%+v", d.ParticipantID)
	fmt.Fprintf(&builder, ", ExchangeID=%+v", d.ExchangeID)
	fmt.Fprintf(&builder, ", PreBalance=%+v", d.PreBalance)
	fmt.Fprintf(&builder, ", CurrMargin=%+v", d.CurrMargin)
	fmt.Fprintf(&builder, ", CloseProfit=%+v", d.CloseProfit)
	fmt.Fprintf(&builder, ", Balance=%+v", d.Balance)
	fmt.Fprintf(&builder, ", Deposit=%+v", d.Deposit)
	fmt.Fprintf(&builder, ", Withdraw=%+v", d.Withdraw)
	fmt.Fprintf(&builder, ", Available=%+v", d.Available)
	fmt.Fprintf(&builder, ", Reserve=%+v", d.Reserve)
	fmt.Fprintf(&builder, ", FrozenMargin=%+v", d.FrozenMargin)

	builder.WriteByte('}')

	return builder.String()
}

// 查询保证金监管系统经纪公司密钥
type CThostFtdcQryCFMMCBrokerKeyField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
}

func (d CThostFtdcQryCFMMCBrokerKeyField) Type() string { return "CThostFtdcQryCFMMCBrokerKeyField" }

func (d CThostFtdcQryCFMMCBrokerKeyField) String() string {
	var builder strings.Builder
	builder.Grow(50)

	builder.WriteString("CThostFtdcQryCFMMCBrokerKeyField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)

	builder.WriteByte('}')

	return builder.String()
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

func (d CThostFtdcCFMMCBrokerKeyField) Type() string { return "CThostFtdcCFMMCBrokerKeyField" }

func (d CThostFtdcCFMMCBrokerKeyField) String() string {
	var builder strings.Builder
	builder.Grow(162)

	builder.WriteString("CThostFtdcCFMMCBrokerKeyField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", ParticipantID=%+v", d.ParticipantID)
	fmt.Fprintf(&builder, ", CreateDate=%+v", d.CreateDate)
	fmt.Fprintf(&builder, ", CreateTime=%+v", d.CreateTime)
	fmt.Fprintf(&builder, ", KeyID=%+v", d.KeyID)
	fmt.Fprintf(&builder, ", CurrentKey=%+v", d.CurrentKey)
	fmt.Fprintf(&builder, ", KeyKind=%+v", d.KeyKind)

	builder.WriteByte('}')

	return builder.String()
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

func (d CThostFtdcCFMMCTradingAccountKeyField) Type() string {
	return "CThostFtdcCFMMCTradingAccountKeyField"
}

func (d CThostFtdcCFMMCTradingAccountKeyField) String() string {
	var builder strings.Builder
	builder.Grow(132)

	builder.WriteString("CThostFtdcCFMMCTradingAccountKeyField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", ParticipantID=%+v", d.ParticipantID)
	fmt.Fprintf(&builder, ", AccountID=%+v", d.AccountID)
	fmt.Fprintf(&builder, ", KeyID=%+v", d.KeyID)
	fmt.Fprintf(&builder, ", CurrentKey=%+v", d.CurrentKey)

	builder.WriteByte('}')

	return builder.String()
}

// 请求查询保证金监管系统经纪公司资金账户密钥
type CThostFtdcQryCFMMCTradingAccountKeyField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID types.TThostFtdcInvestorIDType
}

func (d CThostFtdcQryCFMMCTradingAccountKeyField) Type() string {
	return "CThostFtdcQryCFMMCTradingAccountKeyField"
}

func (d CThostFtdcQryCFMMCTradingAccountKeyField) String() string {
	var builder strings.Builder
	builder.Grow(78)

	builder.WriteString("CThostFtdcQryCFMMCTradingAccountKeyField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", InvestorID=%+v", d.InvestorID)

	builder.WriteByte('}')

	return builder.String()
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

func (d CThostFtdcBrokerUserOTPParamField) Type() string { return "CThostFtdcBrokerUserOTPParamField" }

func (d CThostFtdcBrokerUserOTPParamField) String() string {
	var builder strings.Builder
	builder.Grow(185)

	builder.WriteString("CThostFtdcBrokerUserOTPParamField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", UserID=%+v", d.UserID)
	fmt.Fprintf(&builder, ", OTPVendorsID=%+v", d.OTPVendorsID)
	fmt.Fprintf(&builder, ", SerialNumber=%+v", d.SerialNumber)
	fmt.Fprintf(&builder, ", AuthKey=%+v", d.AuthKey)
	fmt.Fprintf(&builder, ", LastDrift=%+v", d.LastDrift)
	fmt.Fprintf(&builder, ", LastSuccess=%+v", d.LastSuccess)
	fmt.Fprintf(&builder, ", OTPType=%+v", d.OTPType)

	builder.WriteByte('}')

	return builder.String()
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

func (d CThostFtdcManualSyncBrokerUserOTPField) Type() string {
	return "CThostFtdcManualSyncBrokerUserOTPField"
}

func (d CThostFtdcManualSyncBrokerUserOTPField) String() string {
	var builder strings.Builder
	builder.Grow(126)

	builder.WriteString("CThostFtdcManualSyncBrokerUserOTPField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", UserID=%+v", d.UserID)
	fmt.Fprintf(&builder, ", OTPType=%+v", d.OTPType)
	fmt.Fprintf(&builder, ", FirstOTP=%+v", d.FirstOTP)
	fmt.Fprintf(&builder, ", SecondOTP=%+v", d.SecondOTP)

	builder.WriteByte('}')

	return builder.String()
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

func (d CThostFtdcCommRateModelField) Type() string { return "CThostFtdcCommRateModelField" }

func (d CThostFtdcCommRateModelField) String() string {
	var builder strings.Builder
	builder.Grow(90)

	builder.WriteString("CThostFtdcCommRateModelField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", CommModelID=%+v", d.CommModelID)
	fmt.Fprintf(&builder, ", CommModelName=%+v", d.CommModelName)

	builder.WriteByte('}')

	return builder.String()
}

// 请求查询投资者手续费率模板
type CThostFtdcQryCommRateModelField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 手续费率模板代码
	CommModelID types.TThostFtdcInvestorIDType
}

func (d CThostFtdcQryCommRateModelField) Type() string { return "CThostFtdcQryCommRateModelField" }

func (d CThostFtdcQryCommRateModelField) String() string {
	var builder strings.Builder
	builder.Grow(70)

	builder.WriteString("CThostFtdcQryCommRateModelField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", CommModelID=%+v", d.CommModelID)

	builder.WriteByte('}')

	return builder.String()
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

func (d CThostFtdcMarginModelField) Type() string { return "CThostFtdcMarginModelField" }

func (d CThostFtdcMarginModelField) String() string {
	var builder strings.Builder
	builder.Grow(92)

	builder.WriteString("CThostFtdcMarginModelField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", MarginModelID=%+v", d.MarginModelID)
	fmt.Fprintf(&builder, ", MarginModelName=%+v", d.MarginModelName)

	builder.WriteByte('}')

	return builder.String()
}

// 请求查询投资者保证金率模板
type CThostFtdcQryMarginModelField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 保证金率模板代码
	MarginModelID types.TThostFtdcInvestorIDType
}

func (d CThostFtdcQryMarginModelField) Type() string { return "CThostFtdcQryMarginModelField" }

func (d CThostFtdcQryMarginModelField) String() string {
	var builder strings.Builder
	builder.Grow(70)

	builder.WriteString("CThostFtdcQryMarginModelField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", MarginModelID=%+v", d.MarginModelID)

	builder.WriteByte('}')

	return builder.String()
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
	// 保留的无效字段
	reserve1 types.TThostFtdcOldInstrumentIDType
	// 买卖方向
	Direction types.TThostFtdcDirectionType
	// 投机套保标志
	HedgeFlag types.TThostFtdcHedgeFlagType
	// 数量
	Volume types.TThostFtdcVolumeType
	// 投资单元代码
	InvestUnitID types.TThostFtdcInvestUnitIDType
	// 合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
}

func (d CThostFtdcEWarrantOffsetField) Type() string { return "CThostFtdcEWarrantOffsetField" }

func (d CThostFtdcEWarrantOffsetField) String() string {
	var builder strings.Builder
	builder.Grow(223)

	builder.WriteString("CThostFtdcEWarrantOffsetField{")
	fmt.Fprintf(&builder, "TradingDay=%+v", d.TradingDay)
	fmt.Fprintf(&builder, ", BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", InvestorID=%+v", d.InvestorID)
	fmt.Fprintf(&builder, ", ExchangeID=%+v", d.ExchangeID)
	fmt.Fprintf(&builder, ", reserve1=%+v", d.reserve1)
	fmt.Fprintf(&builder, ", Direction=%+v", d.Direction)
	fmt.Fprintf(&builder, ", HedgeFlag=%+v", d.HedgeFlag)
	fmt.Fprintf(&builder, ", Volume=%+v", d.Volume)
	fmt.Fprintf(&builder, ", InvestUnitID=%+v", d.InvestUnitID)
	fmt.Fprintf(&builder, ", InstrumentID=%+v", d.InstrumentID)

	builder.WriteByte('}')

	return builder.String()
}

// 查询仓单折抵信息
type CThostFtdcQryEWarrantOffsetField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	// 交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	// 保留的无效字段
	reserve1 types.TThostFtdcOldInstrumentIDType
	// 投资单元代码
	InvestUnitID types.TThostFtdcInvestUnitIDType
	// 合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
}

func (d CThostFtdcQryEWarrantOffsetField) Type() string { return "CThostFtdcQryEWarrantOffsetField" }

func (d CThostFtdcQryEWarrantOffsetField) String() string {
	var builder strings.Builder
	builder.Grow(152)

	builder.WriteString("CThostFtdcQryEWarrantOffsetField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", InvestorID=%+v", d.InvestorID)
	fmt.Fprintf(&builder, ", ExchangeID=%+v", d.ExchangeID)
	fmt.Fprintf(&builder, ", reserve1=%+v", d.reserve1)
	fmt.Fprintf(&builder, ", InvestUnitID=%+v", d.InvestUnitID)
	fmt.Fprintf(&builder, ", InstrumentID=%+v", d.InstrumentID)

	builder.WriteByte('}')

	return builder.String()
}

// 查询投资者品种跨品种保证金
type CThostFtdcQryInvestorProductGroupMarginField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	// 保留的无效字段
	reserve1 types.TThostFtdcOldInstrumentIDType
	// 投机套保标志
	HedgeFlag types.TThostFtdcHedgeFlagType
	// 交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	// 投资单元代码
	InvestUnitID types.TThostFtdcInvestUnitIDType
	// 品种跨品种标示
	ProductGroupID types.TThostFtdcInstrumentIDType
}

func (d CThostFtdcQryInvestorProductGroupMarginField) Type() string {
	return "CThostFtdcQryInvestorProductGroupMarginField"
}

func (d CThostFtdcQryInvestorProductGroupMarginField) String() string {
	var builder strings.Builder
	builder.Grow(185)

	builder.WriteString("CThostFtdcQryInvestorProductGroupMarginField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", InvestorID=%+v", d.InvestorID)
	fmt.Fprintf(&builder, ", reserve1=%+v", d.reserve1)
	fmt.Fprintf(&builder, ", HedgeFlag=%+v", d.HedgeFlag)
	fmt.Fprintf(&builder, ", ExchangeID=%+v", d.ExchangeID)
	fmt.Fprintf(&builder, ", InvestUnitID=%+v", d.InvestUnitID)
	fmt.Fprintf(&builder, ", ProductGroupID=%+v", d.ProductGroupID)

	builder.WriteByte('}')

	return builder.String()
}

// 投资者品种跨品种保证金
type CThostFtdcInvestorProductGroupMarginField struct {
	// 保留的无效字段
	reserve1 types.TThostFtdcOldInstrumentIDType
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
	// 交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	// 投资单元代码
	InvestUnitID types.TThostFtdcInvestUnitIDType
	// 品种跨品种标示
	ProductGroupID types.TThostFtdcInstrumentIDType
}

func (d CThostFtdcInvestorProductGroupMarginField) Type() string {
	return "CThostFtdcInvestorProductGroupMarginField"
}

func (d CThostFtdcInvestorProductGroupMarginField) String() string {
	var builder strings.Builder
	builder.Grow(723)

	builder.WriteString("CThostFtdcInvestorProductGroupMarginField{")
	fmt.Fprintf(&builder, "reserve1=%+v", d.reserve1)
	fmt.Fprintf(&builder, ", BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", InvestorID=%+v", d.InvestorID)
	fmt.Fprintf(&builder, ", TradingDay=%+v", d.TradingDay)
	fmt.Fprintf(&builder, ", SettlementID=%+v", d.SettlementID)
	fmt.Fprintf(&builder, ", FrozenMargin=%+v", d.FrozenMargin)
	fmt.Fprintf(&builder, ", LongFrozenMargin=%+v", d.LongFrozenMargin)
	fmt.Fprintf(&builder, ", ShortFrozenMargin=%+v", d.ShortFrozenMargin)
	fmt.Fprintf(&builder, ", UseMargin=%+v", d.UseMargin)
	fmt.Fprintf(&builder, ", LongUseMargin=%+v", d.LongUseMargin)
	fmt.Fprintf(&builder, ", ShortUseMargin=%+v", d.ShortUseMargin)
	fmt.Fprintf(&builder, ", ExchMargin=%+v", d.ExchMargin)
	fmt.Fprintf(&builder, ", LongExchMargin=%+v", d.LongExchMargin)
	fmt.Fprintf(&builder, ", ShortExchMargin=%+v", d.ShortExchMargin)
	fmt.Fprintf(&builder, ", CloseProfit=%+v", d.CloseProfit)
	fmt.Fprintf(&builder, ", FrozenCommission=%+v", d.FrozenCommission)
	fmt.Fprintf(&builder, ", Commission=%+v", d.Commission)
	fmt.Fprintf(&builder, ", FrozenCash=%+v", d.FrozenCash)
	fmt.Fprintf(&builder, ", CashIn=%+v", d.CashIn)
	fmt.Fprintf(&builder, ", PositionProfit=%+v", d.PositionProfit)
	fmt.Fprintf(&builder, ", OffsetAmount=%+v", d.OffsetAmount)
	fmt.Fprintf(&builder, ", LongOffsetAmount=%+v", d.LongOffsetAmount)
	fmt.Fprintf(&builder, ", ShortOffsetAmount=%+v", d.ShortOffsetAmount)
	fmt.Fprintf(&builder, ", ExchOffsetAmount=%+v", d.ExchOffsetAmount)
	fmt.Fprintf(&builder, ", LongExchOffsetAmount=%+v", d.LongExchOffsetAmount)
	fmt.Fprintf(&builder, ", ShortExchOffsetAmount=%+v", d.ShortExchOffsetAmount)
	fmt.Fprintf(&builder, ", HedgeFlag=%+v", d.HedgeFlag)
	fmt.Fprintf(&builder, ", ExchangeID=%+v", d.ExchangeID)
	fmt.Fprintf(&builder, ", InvestUnitID=%+v", d.InvestUnitID)
	fmt.Fprintf(&builder, ", ProductGroupID=%+v", d.ProductGroupID)

	builder.WriteByte('}')

	return builder.String()
}

// 查询监控中心用户令牌
type CThostFtdcQueryCFMMCTradingAccountTokenField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	// 投资单元代码
	InvestUnitID types.TThostFtdcInvestUnitIDType
}

func (d CThostFtdcQueryCFMMCTradingAccountTokenField) Type() string {
	return "CThostFtdcQueryCFMMCTradingAccountTokenField"
}

func (d CThostFtdcQueryCFMMCTradingAccountTokenField) String() string {
	var builder strings.Builder
	builder.Grow(104)

	builder.WriteString("CThostFtdcQueryCFMMCTradingAccountTokenField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", InvestorID=%+v", d.InvestorID)
	fmt.Fprintf(&builder, ", InvestUnitID=%+v", d.InvestUnitID)

	builder.WriteByte('}')

	return builder.String()
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

func (d CThostFtdcCFMMCTradingAccountTokenField) Type() string {
	return "CThostFtdcCFMMCTradingAccountTokenField"
}

func (d CThostFtdcCFMMCTradingAccountTokenField) String() string {
	var builder strings.Builder
	builder.Grow(129)

	builder.WriteString("CThostFtdcCFMMCTradingAccountTokenField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", ParticipantID=%+v", d.ParticipantID)
	fmt.Fprintf(&builder, ", AccountID=%+v", d.AccountID)
	fmt.Fprintf(&builder, ", KeyID=%+v", d.KeyID)
	fmt.Fprintf(&builder, ", Token=%+v", d.Token)

	builder.WriteByte('}')

	return builder.String()
}

// 查询产品组
type CThostFtdcQryProductGroupField struct {
	// 保留的无效字段
	reserve1 types.TThostFtdcOldInstrumentIDType
	// 交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	// 产品代码
	ProductID types.TThostFtdcInstrumentIDType
}

func (d CThostFtdcQryProductGroupField) Type() string { return "CThostFtdcQryProductGroupField" }

func (d CThostFtdcQryProductGroupField) String() string {
	var builder strings.Builder
	builder.Grow(87)

	builder.WriteString("CThostFtdcQryProductGroupField{")
	fmt.Fprintf(&builder, "reserve1=%+v", d.reserve1)
	fmt.Fprintf(&builder, ", ExchangeID=%+v", d.ExchangeID)
	fmt.Fprintf(&builder, ", ProductID=%+v", d.ProductID)

	builder.WriteByte('}')

	return builder.String()
}

// 投资者品种跨品种保证金产品组
type CThostFtdcProductGroupField struct {
	// 保留的无效字段
	reserve1 types.TThostFtdcOldInstrumentIDType
	// 交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	// 保留的无效字段
	reserve2 types.TThostFtdcOldInstrumentIDType
	// 产品代码
	ProductID types.TThostFtdcInstrumentIDType
	// 产品组代码
	ProductGroupID types.TThostFtdcInstrumentIDType
}

func (d CThostFtdcProductGroupField) Type() string { return "CThostFtdcProductGroupField" }

func (d CThostFtdcProductGroupField) String() string {
	var builder strings.Builder
	builder.Grow(126)

	builder.WriteString("CThostFtdcProductGroupField{")
	fmt.Fprintf(&builder, "reserve1=%+v", d.reserve1)
	fmt.Fprintf(&builder, ", ExchangeID=%+v", d.ExchangeID)
	fmt.Fprintf(&builder, ", reserve2=%+v", d.reserve2)
	fmt.Fprintf(&builder, ", ProductID=%+v", d.ProductID)
	fmt.Fprintf(&builder, ", ProductGroupID=%+v", d.ProductGroupID)

	builder.WriteByte('}')

	return builder.String()
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

func (d CThostFtdcBulletinField) Type() string { return "CThostFtdcBulletinField" }

func (d CThostFtdcBulletinField) String() string {
	var builder strings.Builder
	builder.Grow(248)

	builder.WriteString("CThostFtdcBulletinField{")
	fmt.Fprintf(&builder, "ExchangeID=%+v", d.ExchangeID)
	fmt.Fprintf(&builder, ", TradingDay=%+v", d.TradingDay)
	fmt.Fprintf(&builder, ", BulletinID=%+v", d.BulletinID)
	fmt.Fprintf(&builder, ", SequenceNo=%+v", d.SequenceNo)
	fmt.Fprintf(&builder, ", NewsType=%+v", d.NewsType)
	fmt.Fprintf(&builder, ", NewsUrgency=%+v", d.NewsUrgency)
	fmt.Fprintf(&builder, ", SendTime=%+v", d.SendTime)
	fmt.Fprintf(&builder, ", Abstract=%+v", d.Abstract)
	fmt.Fprintf(&builder, ", ComeFrom=%+v", d.ComeFrom)
	fmt.Fprintf(&builder, ", Content=%+v", d.Content)
	fmt.Fprintf(&builder, ", URLLink=%+v", d.URLLink)
	fmt.Fprintf(&builder, ", MarketID=%+v", d.MarketID)

	builder.WriteByte('}')

	return builder.String()
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

func (d CThostFtdcQryBulletinField) Type() string { return "CThostFtdcQryBulletinField" }

func (d CThostFtdcQryBulletinField) String() string {
	var builder strings.Builder
	builder.Grow(125)

	builder.WriteString("CThostFtdcQryBulletinField{")
	fmt.Fprintf(&builder, "ExchangeID=%+v", d.ExchangeID)
	fmt.Fprintf(&builder, ", BulletinID=%+v", d.BulletinID)
	fmt.Fprintf(&builder, ", SequenceNo=%+v", d.SequenceNo)
	fmt.Fprintf(&builder, ", NewsType=%+v", d.NewsType)
	fmt.Fprintf(&builder, ", NewsUrgency=%+v", d.NewsUrgency)

	builder.WriteByte('}')

	return builder.String()
}

// MulticastInstrument
type CThostFtdcMulticastInstrumentField struct {
	// 主题号
	TopicID types.TThostFtdcInstallIDType
	// 保留的无效字段
	reserve1 types.TThostFtdcOldInstrumentIDType
	// 合约编号
	InstrumentNo types.TThostFtdcInstallIDType
	// 基准价
	CodePrice types.TThostFtdcPriceType
	// 合约数量乘数
	VolumeMultiple types.TThostFtdcVolumeMultipleType
	// 最小变动价位
	PriceTick types.TThostFtdcPriceType
	// 合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
}

func (d CThostFtdcMulticastInstrumentField) Type() string {
	return "CThostFtdcMulticastInstrumentField"
}

func (d CThostFtdcMulticastInstrumentField) String() string {
	var builder strings.Builder
	builder.Grow(175)

	builder.WriteString("CThostFtdcMulticastInstrumentField{")
	fmt.Fprintf(&builder, "TopicID=%+v", d.TopicID)
	fmt.Fprintf(&builder, ", reserve1=%+v", d.reserve1)
	fmt.Fprintf(&builder, ", InstrumentNo=%+v", d.InstrumentNo)
	fmt.Fprintf(&builder, ", CodePrice=%+v", d.CodePrice)
	fmt.Fprintf(&builder, ", VolumeMultiple=%+v", d.VolumeMultiple)
	fmt.Fprintf(&builder, ", PriceTick=%+v", d.PriceTick)
	fmt.Fprintf(&builder, ", InstrumentID=%+v", d.InstrumentID)

	builder.WriteByte('}')

	return builder.String()
}

// QryMulticastInstrument
type CThostFtdcQryMulticastInstrumentField struct {
	// 主题号
	TopicID types.TThostFtdcInstallIDType
	// 保留的无效字段
	reserve1 types.TThostFtdcOldInstrumentIDType
	// 合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
}

func (d CThostFtdcQryMulticastInstrumentField) Type() string {
	return "CThostFtdcQryMulticastInstrumentField"
}

func (d CThostFtdcQryMulticastInstrumentField) String() string {
	var builder strings.Builder
	builder.Grow(94)

	builder.WriteString("CThostFtdcQryMulticastInstrumentField{")
	fmt.Fprintf(&builder, "TopicID=%+v", d.TopicID)
	fmt.Fprintf(&builder, ", reserve1=%+v", d.reserve1)
	fmt.Fprintf(&builder, ", InstrumentID=%+v", d.InstrumentID)

	builder.WriteByte('}')

	return builder.String()
}

// App客户端权限分配
type CThostFtdcAppIDAuthAssignField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// App代码
	AppID types.TThostFtdcAppIDType
	// 交易中心代码
	DRIdentityID types.TThostFtdcDRIdentityIDType
}

func (d CThostFtdcAppIDAuthAssignField) Type() string { return "CThostFtdcAppIDAuthAssignField" }

func (d CThostFtdcAppIDAuthAssignField) String() string {
	var builder strings.Builder
	builder.Grow(85)

	builder.WriteString("CThostFtdcAppIDAuthAssignField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", AppID=%+v", d.AppID)
	fmt.Fprintf(&builder, ", DRIdentityID=%+v", d.DRIdentityID)

	builder.WriteByte('}')

	return builder.String()
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
	// 长客户姓名
	LongCustomerName types.TThostFtdcLongIndividualNameType
}

func (d CThostFtdcReqOpenAccountField) Type() string { return "CThostFtdcReqOpenAccountField" }

func (d CThostFtdcReqOpenAccountField) String() string {
	var builder strings.Builder
	builder.Grow(930)

	builder.WriteString("CThostFtdcReqOpenAccountField{")
	fmt.Fprintf(&builder, "TradeCode=%+v", d.TradeCode)
	fmt.Fprintf(&builder, ", BankID=%+v", d.BankID)
	fmt.Fprintf(&builder, ", BankBranchID=%+v", d.BankBranchID)
	fmt.Fprintf(&builder, ", BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", BrokerBranchID=%+v", d.BrokerBranchID)
	fmt.Fprintf(&builder, ", TradeDate=%+v", d.TradeDate)
	fmt.Fprintf(&builder, ", TradeTime=%+v", d.TradeTime)
	fmt.Fprintf(&builder, ", BankSerial=%+v", d.BankSerial)
	fmt.Fprintf(&builder, ", TradingDay=%+v", d.TradingDay)
	fmt.Fprintf(&builder, ", PlateSerial=%+v", d.PlateSerial)
	fmt.Fprintf(&builder, ", LastFragment=%+v", d.LastFragment)
	fmt.Fprintf(&builder, ", SessionID=%+v", d.SessionID)
	fmt.Fprintf(&builder, ", CustomerName=%+v", d.CustomerName)
	fmt.Fprintf(&builder, ", IdCardType=%+v", d.IdCardType)
	fmt.Fprintf(&builder, ", IdentifiedCardNo=%+v", d.IdentifiedCardNo)
	fmt.Fprintf(&builder, ", Gender=%+v", d.Gender)
	fmt.Fprintf(&builder, ", CountryCode=%+v", d.CountryCode)
	fmt.Fprintf(&builder, ", CustType=%+v", d.CustType)
	fmt.Fprintf(&builder, ", Address=%+v", d.Address)
	fmt.Fprintf(&builder, ", ZipCode=%+v", d.ZipCode)
	fmt.Fprintf(&builder, ", Telephone=%+v", d.Telephone)
	fmt.Fprintf(&builder, ", MobilePhone=%+v", d.MobilePhone)
	fmt.Fprintf(&builder, ", Fax=%+v", d.Fax)
	fmt.Fprintf(&builder, ", EMail=%+v", d.EMail)
	fmt.Fprintf(&builder, ", MoneyAccountStatus=%+v", d.MoneyAccountStatus)
	fmt.Fprintf(&builder, ", BankAccount=%+v", d.BankAccount)
	fmt.Fprintf(&builder, ", BankPassWord=%+v", d.BankPassWord)
	fmt.Fprintf(&builder, ", AccountID=%+v", d.AccountID)
	fmt.Fprintf(&builder, ", Password=%+v", d.Password)
	fmt.Fprintf(&builder, ", InstallID=%+v", d.InstallID)
	fmt.Fprintf(&builder, ", VerifyCertNoFlag=%+v", d.VerifyCertNoFlag)
	fmt.Fprintf(&builder, ", CurrencyID=%+v", d.CurrencyID)
	fmt.Fprintf(&builder, ", CashExchangeCode=%+v", d.CashExchangeCode)
	fmt.Fprintf(&builder, ", Digest=%+v", d.Digest)
	fmt.Fprintf(&builder, ", BankAccType=%+v", d.BankAccType)
	fmt.Fprintf(&builder, ", DeviceID=%+v", d.DeviceID)
	fmt.Fprintf(&builder, ", BankSecuAccType=%+v", d.BankSecuAccType)
	fmt.Fprintf(&builder, ", BrokerIDByBank=%+v", d.BrokerIDByBank)
	fmt.Fprintf(&builder, ", BankSecuAcc=%+v", d.BankSecuAcc)
	fmt.Fprintf(&builder, ", BankPwdFlag=%+v", d.BankPwdFlag)
	fmt.Fprintf(&builder, ", SecuPwdFlag=%+v", d.SecuPwdFlag)
	fmt.Fprintf(&builder, ", OperNo=%+v", d.OperNo)
	fmt.Fprintf(&builder, ", TID=%+v", d.TID)
	fmt.Fprintf(&builder, ", UserID=%+v", d.UserID)
	fmt.Fprintf(&builder, ", LongCustomerName=%+v", d.LongCustomerName)

	builder.WriteByte('}')

	return builder.String()
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
	// 长客户姓名
	LongCustomerName types.TThostFtdcLongIndividualNameType
}

func (d CThostFtdcReqCancelAccountField) Type() string { return "CThostFtdcReqCancelAccountField" }

func (d CThostFtdcReqCancelAccountField) String() string {
	var builder strings.Builder
	builder.Grow(932)

	builder.WriteString("CThostFtdcReqCancelAccountField{")
	fmt.Fprintf(&builder, "TradeCode=%+v", d.TradeCode)
	fmt.Fprintf(&builder, ", BankID=%+v", d.BankID)
	fmt.Fprintf(&builder, ", BankBranchID=%+v", d.BankBranchID)
	fmt.Fprintf(&builder, ", BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", BrokerBranchID=%+v", d.BrokerBranchID)
	fmt.Fprintf(&builder, ", TradeDate=%+v", d.TradeDate)
	fmt.Fprintf(&builder, ", TradeTime=%+v", d.TradeTime)
	fmt.Fprintf(&builder, ", BankSerial=%+v", d.BankSerial)
	fmt.Fprintf(&builder, ", TradingDay=%+v", d.TradingDay)
	fmt.Fprintf(&builder, ", PlateSerial=%+v", d.PlateSerial)
	fmt.Fprintf(&builder, ", LastFragment=%+v", d.LastFragment)
	fmt.Fprintf(&builder, ", SessionID=%+v", d.SessionID)
	fmt.Fprintf(&builder, ", CustomerName=%+v", d.CustomerName)
	fmt.Fprintf(&builder, ", IdCardType=%+v", d.IdCardType)
	fmt.Fprintf(&builder, ", IdentifiedCardNo=%+v", d.IdentifiedCardNo)
	fmt.Fprintf(&builder, ", Gender=%+v", d.Gender)
	fmt.Fprintf(&builder, ", CountryCode=%+v", d.CountryCode)
	fmt.Fprintf(&builder, ", CustType=%+v", d.CustType)
	fmt.Fprintf(&builder, ", Address=%+v", d.Address)
	fmt.Fprintf(&builder, ", ZipCode=%+v", d.ZipCode)
	fmt.Fprintf(&builder, ", Telephone=%+v", d.Telephone)
	fmt.Fprintf(&builder, ", MobilePhone=%+v", d.MobilePhone)
	fmt.Fprintf(&builder, ", Fax=%+v", d.Fax)
	fmt.Fprintf(&builder, ", EMail=%+v", d.EMail)
	fmt.Fprintf(&builder, ", MoneyAccountStatus=%+v", d.MoneyAccountStatus)
	fmt.Fprintf(&builder, ", BankAccount=%+v", d.BankAccount)
	fmt.Fprintf(&builder, ", BankPassWord=%+v", d.BankPassWord)
	fmt.Fprintf(&builder, ", AccountID=%+v", d.AccountID)
	fmt.Fprintf(&builder, ", Password=%+v", d.Password)
	fmt.Fprintf(&builder, ", InstallID=%+v", d.InstallID)
	fmt.Fprintf(&builder, ", VerifyCertNoFlag=%+v", d.VerifyCertNoFlag)
	fmt.Fprintf(&builder, ", CurrencyID=%+v", d.CurrencyID)
	fmt.Fprintf(&builder, ", CashExchangeCode=%+v", d.CashExchangeCode)
	fmt.Fprintf(&builder, ", Digest=%+v", d.Digest)
	fmt.Fprintf(&builder, ", BankAccType=%+v", d.BankAccType)
	fmt.Fprintf(&builder, ", DeviceID=%+v", d.DeviceID)
	fmt.Fprintf(&builder, ", BankSecuAccType=%+v", d.BankSecuAccType)
	fmt.Fprintf(&builder, ", BrokerIDByBank=%+v", d.BrokerIDByBank)
	fmt.Fprintf(&builder, ", BankSecuAcc=%+v", d.BankSecuAcc)
	fmt.Fprintf(&builder, ", BankPwdFlag=%+v", d.BankPwdFlag)
	fmt.Fprintf(&builder, ", SecuPwdFlag=%+v", d.SecuPwdFlag)
	fmt.Fprintf(&builder, ", OperNo=%+v", d.OperNo)
	fmt.Fprintf(&builder, ", TID=%+v", d.TID)
	fmt.Fprintf(&builder, ", UserID=%+v", d.UserID)
	fmt.Fprintf(&builder, ", LongCustomerName=%+v", d.LongCustomerName)

	builder.WriteByte('}')

	return builder.String()
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
	// 长客户姓名
	LongCustomerName types.TThostFtdcLongIndividualNameType
}

func (d CThostFtdcReqChangeAccountField) Type() string { return "CThostFtdcReqChangeAccountField" }

func (d CThostFtdcReqChangeAccountField) String() string {
	var builder strings.Builder
	builder.Grow(859)

	builder.WriteString("CThostFtdcReqChangeAccountField{")
	fmt.Fprintf(&builder, "TradeCode=%+v", d.TradeCode)
	fmt.Fprintf(&builder, ", BankID=%+v", d.BankID)
	fmt.Fprintf(&builder, ", BankBranchID=%+v", d.BankBranchID)
	fmt.Fprintf(&builder, ", BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", BrokerBranchID=%+v", d.BrokerBranchID)
	fmt.Fprintf(&builder, ", TradeDate=%+v", d.TradeDate)
	fmt.Fprintf(&builder, ", TradeTime=%+v", d.TradeTime)
	fmt.Fprintf(&builder, ", BankSerial=%+v", d.BankSerial)
	fmt.Fprintf(&builder, ", TradingDay=%+v", d.TradingDay)
	fmt.Fprintf(&builder, ", PlateSerial=%+v", d.PlateSerial)
	fmt.Fprintf(&builder, ", LastFragment=%+v", d.LastFragment)
	fmt.Fprintf(&builder, ", SessionID=%+v", d.SessionID)
	fmt.Fprintf(&builder, ", CustomerName=%+v", d.CustomerName)
	fmt.Fprintf(&builder, ", IdCardType=%+v", d.IdCardType)
	fmt.Fprintf(&builder, ", IdentifiedCardNo=%+v", d.IdentifiedCardNo)
	fmt.Fprintf(&builder, ", Gender=%+v", d.Gender)
	fmt.Fprintf(&builder, ", CountryCode=%+v", d.CountryCode)
	fmt.Fprintf(&builder, ", CustType=%+v", d.CustType)
	fmt.Fprintf(&builder, ", Address=%+v", d.Address)
	fmt.Fprintf(&builder, ", ZipCode=%+v", d.ZipCode)
	fmt.Fprintf(&builder, ", Telephone=%+v", d.Telephone)
	fmt.Fprintf(&builder, ", MobilePhone=%+v", d.MobilePhone)
	fmt.Fprintf(&builder, ", Fax=%+v", d.Fax)
	fmt.Fprintf(&builder, ", EMail=%+v", d.EMail)
	fmt.Fprintf(&builder, ", MoneyAccountStatus=%+v", d.MoneyAccountStatus)
	fmt.Fprintf(&builder, ", BankAccount=%+v", d.BankAccount)
	fmt.Fprintf(&builder, ", BankPassWord=%+v", d.BankPassWord)
	fmt.Fprintf(&builder, ", NewBankAccount=%+v", d.NewBankAccount)
	fmt.Fprintf(&builder, ", NewBankPassWord=%+v", d.NewBankPassWord)
	fmt.Fprintf(&builder, ", AccountID=%+v", d.AccountID)
	fmt.Fprintf(&builder, ", Password=%+v", d.Password)
	fmt.Fprintf(&builder, ", BankAccType=%+v", d.BankAccType)
	fmt.Fprintf(&builder, ", InstallID=%+v", d.InstallID)
	fmt.Fprintf(&builder, ", VerifyCertNoFlag=%+v", d.VerifyCertNoFlag)
	fmt.Fprintf(&builder, ", CurrencyID=%+v", d.CurrencyID)
	fmt.Fprintf(&builder, ", BrokerIDByBank=%+v", d.BrokerIDByBank)
	fmt.Fprintf(&builder, ", BankPwdFlag=%+v", d.BankPwdFlag)
	fmt.Fprintf(&builder, ", SecuPwdFlag=%+v", d.SecuPwdFlag)
	fmt.Fprintf(&builder, ", TID=%+v", d.TID)
	fmt.Fprintf(&builder, ", Digest=%+v", d.Digest)
	fmt.Fprintf(&builder, ", LongCustomerName=%+v", d.LongCustomerName)

	builder.WriteByte('}')

	return builder.String()
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
	// 长客户姓名
	LongCustomerName types.TThostFtdcLongIndividualNameType
}

func (d CThostFtdcReqTransferField) Type() string { return "CThostFtdcReqTransferField" }

func (d CThostFtdcReqTransferField) String() string {
	var builder strings.Builder
	builder.Grow(920)

	builder.WriteString("CThostFtdcReqTransferField{")
	fmt.Fprintf(&builder, "TradeCode=%+v", d.TradeCode)
	fmt.Fprintf(&builder, ", BankID=%+v", d.BankID)
	fmt.Fprintf(&builder, ", BankBranchID=%+v", d.BankBranchID)
	fmt.Fprintf(&builder, ", BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", BrokerBranchID=%+v", d.BrokerBranchID)
	fmt.Fprintf(&builder, ", TradeDate=%+v", d.TradeDate)
	fmt.Fprintf(&builder, ", TradeTime=%+v", d.TradeTime)
	fmt.Fprintf(&builder, ", BankSerial=%+v", d.BankSerial)
	fmt.Fprintf(&builder, ", TradingDay=%+v", d.TradingDay)
	fmt.Fprintf(&builder, ", PlateSerial=%+v", d.PlateSerial)
	fmt.Fprintf(&builder, ", LastFragment=%+v", d.LastFragment)
	fmt.Fprintf(&builder, ", SessionID=%+v", d.SessionID)
	fmt.Fprintf(&builder, ", CustomerName=%+v", d.CustomerName)
	fmt.Fprintf(&builder, ", IdCardType=%+v", d.IdCardType)
	fmt.Fprintf(&builder, ", IdentifiedCardNo=%+v", d.IdentifiedCardNo)
	fmt.Fprintf(&builder, ", CustType=%+v", d.CustType)
	fmt.Fprintf(&builder, ", BankAccount=%+v", d.BankAccount)
	fmt.Fprintf(&builder, ", BankPassWord=%+v", d.BankPassWord)
	fmt.Fprintf(&builder, ", AccountID=%+v", d.AccountID)
	fmt.Fprintf(&builder, ", Password=%+v", d.Password)
	fmt.Fprintf(&builder, ", InstallID=%+v", d.InstallID)
	fmt.Fprintf(&builder, ", FutureSerial=%+v", d.FutureSerial)
	fmt.Fprintf(&builder, ", UserID=%+v", d.UserID)
	fmt.Fprintf(&builder, ", VerifyCertNoFlag=%+v", d.VerifyCertNoFlag)
	fmt.Fprintf(&builder, ", CurrencyID=%+v", d.CurrencyID)
	fmt.Fprintf(&builder, ", TradeAmount=%+v", d.TradeAmount)
	fmt.Fprintf(&builder, ", FutureFetchAmount=%+v", d.FutureFetchAmount)
	fmt.Fprintf(&builder, ", FeePayFlag=%+v", d.FeePayFlag)
	fmt.Fprintf(&builder, ", CustFee=%+v", d.CustFee)
	fmt.Fprintf(&builder, ", BrokerFee=%+v", d.BrokerFee)
	fmt.Fprintf(&builder, ", Message=%+v", d.Message)
	fmt.Fprintf(&builder, ", Digest=%+v", d.Digest)
	fmt.Fprintf(&builder, ", BankAccType=%+v", d.BankAccType)
	fmt.Fprintf(&builder, ", DeviceID=%+v", d.DeviceID)
	fmt.Fprintf(&builder, ", BankSecuAccType=%+v", d.BankSecuAccType)
	fmt.Fprintf(&builder, ", BrokerIDByBank=%+v", d.BrokerIDByBank)
	fmt.Fprintf(&builder, ", BankSecuAcc=%+v", d.BankSecuAcc)
	fmt.Fprintf(&builder, ", BankPwdFlag=%+v", d.BankPwdFlag)
	fmt.Fprintf(&builder, ", SecuPwdFlag=%+v", d.SecuPwdFlag)
	fmt.Fprintf(&builder, ", OperNo=%+v", d.OperNo)
	fmt.Fprintf(&builder, ", RequestID=%+v", d.RequestID)
	fmt.Fprintf(&builder, ", TID=%+v", d.TID)
	fmt.Fprintf(&builder, ", TransferStatus=%+v", d.TransferStatus)
	fmt.Fprintf(&builder, ", LongCustomerName=%+v", d.LongCustomerName)

	builder.WriteByte('}')

	return builder.String()
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
	// 长客户姓名
	LongCustomerName types.TThostFtdcLongIndividualNameType
}

func (d CThostFtdcRspTransferField) Type() string { return "CThostFtdcRspTransferField" }

func (d CThostFtdcRspTransferField) String() string {
	var builder strings.Builder
	builder.Grow(955)

	builder.WriteString("CThostFtdcRspTransferField{")
	fmt.Fprintf(&builder, "TradeCode=%+v", d.TradeCode)
	fmt.Fprintf(&builder, ", BankID=%+v", d.BankID)
	fmt.Fprintf(&builder, ", BankBranchID=%+v", d.BankBranchID)
	fmt.Fprintf(&builder, ", BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", BrokerBranchID=%+v", d.BrokerBranchID)
	fmt.Fprintf(&builder, ", TradeDate=%+v", d.TradeDate)
	fmt.Fprintf(&builder, ", TradeTime=%+v", d.TradeTime)
	fmt.Fprintf(&builder, ", BankSerial=%+v", d.BankSerial)
	fmt.Fprintf(&builder, ", TradingDay=%+v", d.TradingDay)
	fmt.Fprintf(&builder, ", PlateSerial=%+v", d.PlateSerial)
	fmt.Fprintf(&builder, ", LastFragment=%+v", d.LastFragment)
	fmt.Fprintf(&builder, ", SessionID=%+v", d.SessionID)
	fmt.Fprintf(&builder, ", CustomerName=%+v", d.CustomerName)
	fmt.Fprintf(&builder, ", IdCardType=%+v", d.IdCardType)
	fmt.Fprintf(&builder, ", IdentifiedCardNo=%+v", d.IdentifiedCardNo)
	fmt.Fprintf(&builder, ", CustType=%+v", d.CustType)
	fmt.Fprintf(&builder, ", BankAccount=%+v", d.BankAccount)
	fmt.Fprintf(&builder, ", BankPassWord=%+v", d.BankPassWord)
	fmt.Fprintf(&builder, ", AccountID=%+v", d.AccountID)
	fmt.Fprintf(&builder, ", Password=%+v", d.Password)
	fmt.Fprintf(&builder, ", InstallID=%+v", d.InstallID)
	fmt.Fprintf(&builder, ", FutureSerial=%+v", d.FutureSerial)
	fmt.Fprintf(&builder, ", UserID=%+v", d.UserID)
	fmt.Fprintf(&builder, ", VerifyCertNoFlag=%+v", d.VerifyCertNoFlag)
	fmt.Fprintf(&builder, ", CurrencyID=%+v", d.CurrencyID)
	fmt.Fprintf(&builder, ", TradeAmount=%+v", d.TradeAmount)
	fmt.Fprintf(&builder, ", FutureFetchAmount=%+v", d.FutureFetchAmount)
	fmt.Fprintf(&builder, ", FeePayFlag=%+v", d.FeePayFlag)
	fmt.Fprintf(&builder, ", CustFee=%+v", d.CustFee)
	fmt.Fprintf(&builder, ", BrokerFee=%+v", d.BrokerFee)
	fmt.Fprintf(&builder, ", Message=%+v", d.Message)
	fmt.Fprintf(&builder, ", Digest=%+v", d.Digest)
	fmt.Fprintf(&builder, ", BankAccType=%+v", d.BankAccType)
	fmt.Fprintf(&builder, ", DeviceID=%+v", d.DeviceID)
	fmt.Fprintf(&builder, ", BankSecuAccType=%+v", d.BankSecuAccType)
	fmt.Fprintf(&builder, ", BrokerIDByBank=%+v", d.BrokerIDByBank)
	fmt.Fprintf(&builder, ", BankSecuAcc=%+v", d.BankSecuAcc)
	fmt.Fprintf(&builder, ", BankPwdFlag=%+v", d.BankPwdFlag)
	fmt.Fprintf(&builder, ", SecuPwdFlag=%+v", d.SecuPwdFlag)
	fmt.Fprintf(&builder, ", OperNo=%+v", d.OperNo)
	fmt.Fprintf(&builder, ", RequestID=%+v", d.RequestID)
	fmt.Fprintf(&builder, ", TID=%+v", d.TID)
	fmt.Fprintf(&builder, ", TransferStatus=%+v", d.TransferStatus)
	fmt.Fprintf(&builder, ", ErrorID=%+v", d.ErrorID)
	fmt.Fprintf(&builder, ", ErrorMsg=%+v", d.ErrorMsg)
	fmt.Fprintf(&builder, ", LongCustomerName=%+v", d.LongCustomerName)

	builder.WriteByte('}')

	return builder.String()
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
	// 长客户姓名
	LongCustomerName types.TThostFtdcLongIndividualNameType
}

func (d CThostFtdcReqRepealField) Type() string { return "CThostFtdcReqRepealField" }

func (d CThostFtdcReqRepealField) String() string {
	var builder strings.Builder
	builder.Grow(1100)

	builder.WriteString("CThostFtdcReqRepealField{")
	fmt.Fprintf(&builder, "RepealTimeInterval=%+v", d.RepealTimeInterval)
	fmt.Fprintf(&builder, ", RepealedTimes=%+v", d.RepealedTimes)
	fmt.Fprintf(&builder, ", BankRepealFlag=%+v", d.BankRepealFlag)
	fmt.Fprintf(&builder, ", BrokerRepealFlag=%+v", d.BrokerRepealFlag)
	fmt.Fprintf(&builder, ", PlateRepealSerial=%+v", d.PlateRepealSerial)
	fmt.Fprintf(&builder, ", BankRepealSerial=%+v", d.BankRepealSerial)
	fmt.Fprintf(&builder, ", FutureRepealSerial=%+v", d.FutureRepealSerial)
	fmt.Fprintf(&builder, ", TradeCode=%+v", d.TradeCode)
	fmt.Fprintf(&builder, ", BankID=%+v", d.BankID)
	fmt.Fprintf(&builder, ", BankBranchID=%+v", d.BankBranchID)
	fmt.Fprintf(&builder, ", BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", BrokerBranchID=%+v", d.BrokerBranchID)
	fmt.Fprintf(&builder, ", TradeDate=%+v", d.TradeDate)
	fmt.Fprintf(&builder, ", TradeTime=%+v", d.TradeTime)
	fmt.Fprintf(&builder, ", BankSerial=%+v", d.BankSerial)
	fmt.Fprintf(&builder, ", TradingDay=%+v", d.TradingDay)
	fmt.Fprintf(&builder, ", PlateSerial=%+v", d.PlateSerial)
	fmt.Fprintf(&builder, ", LastFragment=%+v", d.LastFragment)
	fmt.Fprintf(&builder, ", SessionID=%+v", d.SessionID)
	fmt.Fprintf(&builder, ", CustomerName=%+v", d.CustomerName)
	fmt.Fprintf(&builder, ", IdCardType=%+v", d.IdCardType)
	fmt.Fprintf(&builder, ", IdentifiedCardNo=%+v", d.IdentifiedCardNo)
	fmt.Fprintf(&builder, ", CustType=%+v", d.CustType)
	fmt.Fprintf(&builder, ", BankAccount=%+v", d.BankAccount)
	fmt.Fprintf(&builder, ", BankPassWord=%+v", d.BankPassWord)
	fmt.Fprintf(&builder, ", AccountID=%+v", d.AccountID)
	fmt.Fprintf(&builder, ", Password=%+v", d.Password)
	fmt.Fprintf(&builder, ", InstallID=%+v", d.InstallID)
	fmt.Fprintf(&builder, ", FutureSerial=%+v", d.FutureSerial)
	fmt.Fprintf(&builder, ", UserID=%+v", d.UserID)
	fmt.Fprintf(&builder, ", VerifyCertNoFlag=%+v", d.VerifyCertNoFlag)
	fmt.Fprintf(&builder, ", CurrencyID=%+v", d.CurrencyID)
	fmt.Fprintf(&builder, ", TradeAmount=%+v", d.TradeAmount)
	fmt.Fprintf(&builder, ", FutureFetchAmount=%+v", d.FutureFetchAmount)
	fmt.Fprintf(&builder, ", FeePayFlag=%+v", d.FeePayFlag)
	fmt.Fprintf(&builder, ", CustFee=%+v", d.CustFee)
	fmt.Fprintf(&builder, ", BrokerFee=%+v", d.BrokerFee)
	fmt.Fprintf(&builder, ", Message=%+v", d.Message)
	fmt.Fprintf(&builder, ", Digest=%+v", d.Digest)
	fmt.Fprintf(&builder, ", BankAccType=%+v", d.BankAccType)
	fmt.Fprintf(&builder, ", DeviceID=%+v", d.DeviceID)
	fmt.Fprintf(&builder, ", BankSecuAccType=%+v", d.BankSecuAccType)
	fmt.Fprintf(&builder, ", BrokerIDByBank=%+v", d.BrokerIDByBank)
	fmt.Fprintf(&builder, ", BankSecuAcc=%+v", d.BankSecuAcc)
	fmt.Fprintf(&builder, ", BankPwdFlag=%+v", d.BankPwdFlag)
	fmt.Fprintf(&builder, ", SecuPwdFlag=%+v", d.SecuPwdFlag)
	fmt.Fprintf(&builder, ", OperNo=%+v", d.OperNo)
	fmt.Fprintf(&builder, ", RequestID=%+v", d.RequestID)
	fmt.Fprintf(&builder, ", TID=%+v", d.TID)
	fmt.Fprintf(&builder, ", TransferStatus=%+v", d.TransferStatus)
	fmt.Fprintf(&builder, ", LongCustomerName=%+v", d.LongCustomerName)

	builder.WriteByte('}')

	return builder.String()
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
	// 长客户姓名
	LongCustomerName types.TThostFtdcLongIndividualNameType
}

func (d CThostFtdcRspRepealField) Type() string { return "CThostFtdcRspRepealField" }

func (d CThostFtdcRspRepealField) String() string {
	var builder strings.Builder
	builder.Grow(1135)

	builder.WriteString("CThostFtdcRspRepealField{")
	fmt.Fprintf(&builder, "RepealTimeInterval=%+v", d.RepealTimeInterval)
	fmt.Fprintf(&builder, ", RepealedTimes=%+v", d.RepealedTimes)
	fmt.Fprintf(&builder, ", BankRepealFlag=%+v", d.BankRepealFlag)
	fmt.Fprintf(&builder, ", BrokerRepealFlag=%+v", d.BrokerRepealFlag)
	fmt.Fprintf(&builder, ", PlateRepealSerial=%+v", d.PlateRepealSerial)
	fmt.Fprintf(&builder, ", BankRepealSerial=%+v", d.BankRepealSerial)
	fmt.Fprintf(&builder, ", FutureRepealSerial=%+v", d.FutureRepealSerial)
	fmt.Fprintf(&builder, ", TradeCode=%+v", d.TradeCode)
	fmt.Fprintf(&builder, ", BankID=%+v", d.BankID)
	fmt.Fprintf(&builder, ", BankBranchID=%+v", d.BankBranchID)
	fmt.Fprintf(&builder, ", BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", BrokerBranchID=%+v", d.BrokerBranchID)
	fmt.Fprintf(&builder, ", TradeDate=%+v", d.TradeDate)
	fmt.Fprintf(&builder, ", TradeTime=%+v", d.TradeTime)
	fmt.Fprintf(&builder, ", BankSerial=%+v", d.BankSerial)
	fmt.Fprintf(&builder, ", TradingDay=%+v", d.TradingDay)
	fmt.Fprintf(&builder, ", PlateSerial=%+v", d.PlateSerial)
	fmt.Fprintf(&builder, ", LastFragment=%+v", d.LastFragment)
	fmt.Fprintf(&builder, ", SessionID=%+v", d.SessionID)
	fmt.Fprintf(&builder, ", CustomerName=%+v", d.CustomerName)
	fmt.Fprintf(&builder, ", IdCardType=%+v", d.IdCardType)
	fmt.Fprintf(&builder, ", IdentifiedCardNo=%+v", d.IdentifiedCardNo)
	fmt.Fprintf(&builder, ", CustType=%+v", d.CustType)
	fmt.Fprintf(&builder, ", BankAccount=%+v", d.BankAccount)
	fmt.Fprintf(&builder, ", BankPassWord=%+v", d.BankPassWord)
	fmt.Fprintf(&builder, ", AccountID=%+v", d.AccountID)
	fmt.Fprintf(&builder, ", Password=%+v", d.Password)
	fmt.Fprintf(&builder, ", InstallID=%+v", d.InstallID)
	fmt.Fprintf(&builder, ", FutureSerial=%+v", d.FutureSerial)
	fmt.Fprintf(&builder, ", UserID=%+v", d.UserID)
	fmt.Fprintf(&builder, ", VerifyCertNoFlag=%+v", d.VerifyCertNoFlag)
	fmt.Fprintf(&builder, ", CurrencyID=%+v", d.CurrencyID)
	fmt.Fprintf(&builder, ", TradeAmount=%+v", d.TradeAmount)
	fmt.Fprintf(&builder, ", FutureFetchAmount=%+v", d.FutureFetchAmount)
	fmt.Fprintf(&builder, ", FeePayFlag=%+v", d.FeePayFlag)
	fmt.Fprintf(&builder, ", CustFee=%+v", d.CustFee)
	fmt.Fprintf(&builder, ", BrokerFee=%+v", d.BrokerFee)
	fmt.Fprintf(&builder, ", Message=%+v", d.Message)
	fmt.Fprintf(&builder, ", Digest=%+v", d.Digest)
	fmt.Fprintf(&builder, ", BankAccType=%+v", d.BankAccType)
	fmt.Fprintf(&builder, ", DeviceID=%+v", d.DeviceID)
	fmt.Fprintf(&builder, ", BankSecuAccType=%+v", d.BankSecuAccType)
	fmt.Fprintf(&builder, ", BrokerIDByBank=%+v", d.BrokerIDByBank)
	fmt.Fprintf(&builder, ", BankSecuAcc=%+v", d.BankSecuAcc)
	fmt.Fprintf(&builder, ", BankPwdFlag=%+v", d.BankPwdFlag)
	fmt.Fprintf(&builder, ", SecuPwdFlag=%+v", d.SecuPwdFlag)
	fmt.Fprintf(&builder, ", OperNo=%+v", d.OperNo)
	fmt.Fprintf(&builder, ", RequestID=%+v", d.RequestID)
	fmt.Fprintf(&builder, ", TID=%+v", d.TID)
	fmt.Fprintf(&builder, ", TransferStatus=%+v", d.TransferStatus)
	fmt.Fprintf(&builder, ", ErrorID=%+v", d.ErrorID)
	fmt.Fprintf(&builder, ", ErrorMsg=%+v", d.ErrorMsg)
	fmt.Fprintf(&builder, ", LongCustomerName=%+v", d.LongCustomerName)

	builder.WriteByte('}')

	return builder.String()
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
	// 长客户姓名
	LongCustomerName types.TThostFtdcLongIndividualNameType
}

func (d CThostFtdcReqQueryAccountField) Type() string { return "CThostFtdcReqQueryAccountField" }

func (d CThostFtdcReqQueryAccountField) String() string {
	var builder strings.Builder
	builder.Grow(779)

	builder.WriteString("CThostFtdcReqQueryAccountField{")
	fmt.Fprintf(&builder, "TradeCode=%+v", d.TradeCode)
	fmt.Fprintf(&builder, ", BankID=%+v", d.BankID)
	fmt.Fprintf(&builder, ", BankBranchID=%+v", d.BankBranchID)
	fmt.Fprintf(&builder, ", BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", BrokerBranchID=%+v", d.BrokerBranchID)
	fmt.Fprintf(&builder, ", TradeDate=%+v", d.TradeDate)
	fmt.Fprintf(&builder, ", TradeTime=%+v", d.TradeTime)
	fmt.Fprintf(&builder, ", BankSerial=%+v", d.BankSerial)
	fmt.Fprintf(&builder, ", TradingDay=%+v", d.TradingDay)
	fmt.Fprintf(&builder, ", PlateSerial=%+v", d.PlateSerial)
	fmt.Fprintf(&builder, ", LastFragment=%+v", d.LastFragment)
	fmt.Fprintf(&builder, ", SessionID=%+v", d.SessionID)
	fmt.Fprintf(&builder, ", CustomerName=%+v", d.CustomerName)
	fmt.Fprintf(&builder, ", IdCardType=%+v", d.IdCardType)
	fmt.Fprintf(&builder, ", IdentifiedCardNo=%+v", d.IdentifiedCardNo)
	fmt.Fprintf(&builder, ", CustType=%+v", d.CustType)
	fmt.Fprintf(&builder, ", BankAccount=%+v", d.BankAccount)
	fmt.Fprintf(&builder, ", BankPassWord=%+v", d.BankPassWord)
	fmt.Fprintf(&builder, ", AccountID=%+v", d.AccountID)
	fmt.Fprintf(&builder, ", Password=%+v", d.Password)
	fmt.Fprintf(&builder, ", FutureSerial=%+v", d.FutureSerial)
	fmt.Fprintf(&builder, ", InstallID=%+v", d.InstallID)
	fmt.Fprintf(&builder, ", UserID=%+v", d.UserID)
	fmt.Fprintf(&builder, ", VerifyCertNoFlag=%+v", d.VerifyCertNoFlag)
	fmt.Fprintf(&builder, ", CurrencyID=%+v", d.CurrencyID)
	fmt.Fprintf(&builder, ", Digest=%+v", d.Digest)
	fmt.Fprintf(&builder, ", BankAccType=%+v", d.BankAccType)
	fmt.Fprintf(&builder, ", DeviceID=%+v", d.DeviceID)
	fmt.Fprintf(&builder, ", BankSecuAccType=%+v", d.BankSecuAccType)
	fmt.Fprintf(&builder, ", BrokerIDByBank=%+v", d.BrokerIDByBank)
	fmt.Fprintf(&builder, ", BankSecuAcc=%+v", d.BankSecuAcc)
	fmt.Fprintf(&builder, ", BankPwdFlag=%+v", d.BankPwdFlag)
	fmt.Fprintf(&builder, ", SecuPwdFlag=%+v", d.SecuPwdFlag)
	fmt.Fprintf(&builder, ", OperNo=%+v", d.OperNo)
	fmt.Fprintf(&builder, ", RequestID=%+v", d.RequestID)
	fmt.Fprintf(&builder, ", TID=%+v", d.TID)
	fmt.Fprintf(&builder, ", LongCustomerName=%+v", d.LongCustomerName)

	builder.WriteByte('}')

	return builder.String()
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
	// 长客户姓名
	LongCustomerName types.TThostFtdcLongIndividualNameType
}

func (d CThostFtdcRspQueryAccountField) Type() string { return "CThostFtdcRspQueryAccountField" }

func (d CThostFtdcRspQueryAccountField) String() string {
	var builder strings.Builder
	builder.Grow(827)

	builder.WriteString("CThostFtdcRspQueryAccountField{")
	fmt.Fprintf(&builder, "TradeCode=%+v", d.TradeCode)
	fmt.Fprintf(&builder, ", BankID=%+v", d.BankID)
	fmt.Fprintf(&builder, ", BankBranchID=%+v", d.BankBranchID)
	fmt.Fprintf(&builder, ", BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", BrokerBranchID=%+v", d.BrokerBranchID)
	fmt.Fprintf(&builder, ", TradeDate=%+v", d.TradeDate)
	fmt.Fprintf(&builder, ", TradeTime=%+v", d.TradeTime)
	fmt.Fprintf(&builder, ", BankSerial=%+v", d.BankSerial)
	fmt.Fprintf(&builder, ", TradingDay=%+v", d.TradingDay)
	fmt.Fprintf(&builder, ", PlateSerial=%+v", d.PlateSerial)
	fmt.Fprintf(&builder, ", LastFragment=%+v", d.LastFragment)
	fmt.Fprintf(&builder, ", SessionID=%+v", d.SessionID)
	fmt.Fprintf(&builder, ", CustomerName=%+v", d.CustomerName)
	fmt.Fprintf(&builder, ", IdCardType=%+v", d.IdCardType)
	fmt.Fprintf(&builder, ", IdentifiedCardNo=%+v", d.IdentifiedCardNo)
	fmt.Fprintf(&builder, ", CustType=%+v", d.CustType)
	fmt.Fprintf(&builder, ", BankAccount=%+v", d.BankAccount)
	fmt.Fprintf(&builder, ", BankPassWord=%+v", d.BankPassWord)
	fmt.Fprintf(&builder, ", AccountID=%+v", d.AccountID)
	fmt.Fprintf(&builder, ", Password=%+v", d.Password)
	fmt.Fprintf(&builder, ", FutureSerial=%+v", d.FutureSerial)
	fmt.Fprintf(&builder, ", InstallID=%+v", d.InstallID)
	fmt.Fprintf(&builder, ", UserID=%+v", d.UserID)
	fmt.Fprintf(&builder, ", VerifyCertNoFlag=%+v", d.VerifyCertNoFlag)
	fmt.Fprintf(&builder, ", CurrencyID=%+v", d.CurrencyID)
	fmt.Fprintf(&builder, ", Digest=%+v", d.Digest)
	fmt.Fprintf(&builder, ", BankAccType=%+v", d.BankAccType)
	fmt.Fprintf(&builder, ", DeviceID=%+v", d.DeviceID)
	fmt.Fprintf(&builder, ", BankSecuAccType=%+v", d.BankSecuAccType)
	fmt.Fprintf(&builder, ", BrokerIDByBank=%+v", d.BrokerIDByBank)
	fmt.Fprintf(&builder, ", BankSecuAcc=%+v", d.BankSecuAcc)
	fmt.Fprintf(&builder, ", BankPwdFlag=%+v", d.BankPwdFlag)
	fmt.Fprintf(&builder, ", SecuPwdFlag=%+v", d.SecuPwdFlag)
	fmt.Fprintf(&builder, ", OperNo=%+v", d.OperNo)
	fmt.Fprintf(&builder, ", RequestID=%+v", d.RequestID)
	fmt.Fprintf(&builder, ", TID=%+v", d.TID)
	fmt.Fprintf(&builder, ", BankUseAmount=%+v", d.BankUseAmount)
	fmt.Fprintf(&builder, ", BankFetchAmount=%+v", d.BankFetchAmount)
	fmt.Fprintf(&builder, ", LongCustomerName=%+v", d.LongCustomerName)

	builder.WriteByte('}')

	return builder.String()
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

func (d CThostFtdcFutureSignIOField) Type() string { return "CThostFtdcFutureSignIOField" }

func (d CThostFtdcFutureSignIOField) String() string {
	var builder strings.Builder
	builder.Grow(427)

	builder.WriteString("CThostFtdcFutureSignIOField{")
	fmt.Fprintf(&builder, "TradeCode=%+v", d.TradeCode)
	fmt.Fprintf(&builder, ", BankID=%+v", d.BankID)
	fmt.Fprintf(&builder, ", BankBranchID=%+v", d.BankBranchID)
	fmt.Fprintf(&builder, ", BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", BrokerBranchID=%+v", d.BrokerBranchID)
	fmt.Fprintf(&builder, ", TradeDate=%+v", d.TradeDate)
	fmt.Fprintf(&builder, ", TradeTime=%+v", d.TradeTime)
	fmt.Fprintf(&builder, ", BankSerial=%+v", d.BankSerial)
	fmt.Fprintf(&builder, ", TradingDay=%+v", d.TradingDay)
	fmt.Fprintf(&builder, ", PlateSerial=%+v", d.PlateSerial)
	fmt.Fprintf(&builder, ", LastFragment=%+v", d.LastFragment)
	fmt.Fprintf(&builder, ", SessionID=%+v", d.SessionID)
	fmt.Fprintf(&builder, ", InstallID=%+v", d.InstallID)
	fmt.Fprintf(&builder, ", UserID=%+v", d.UserID)
	fmt.Fprintf(&builder, ", Digest=%+v", d.Digest)
	fmt.Fprintf(&builder, ", CurrencyID=%+v", d.CurrencyID)
	fmt.Fprintf(&builder, ", DeviceID=%+v", d.DeviceID)
	fmt.Fprintf(&builder, ", BrokerIDByBank=%+v", d.BrokerIDByBank)
	fmt.Fprintf(&builder, ", OperNo=%+v", d.OperNo)
	fmt.Fprintf(&builder, ", RequestID=%+v", d.RequestID)
	fmt.Fprintf(&builder, ", TID=%+v", d.TID)

	builder.WriteByte('}')

	return builder.String()
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

func (d CThostFtdcRspFutureSignInField) Type() string { return "CThostFtdcRspFutureSignInField" }

func (d CThostFtdcRspFutureSignInField) String() string {
	var builder strings.Builder
	builder.Grow(497)

	builder.WriteString("CThostFtdcRspFutureSignInField{")
	fmt.Fprintf(&builder, "TradeCode=%+v", d.TradeCode)
	fmt.Fprintf(&builder, ", BankID=%+v", d.BankID)
	fmt.Fprintf(&builder, ", BankBranchID=%+v", d.BankBranchID)
	fmt.Fprintf(&builder, ", BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", BrokerBranchID=%+v", d.BrokerBranchID)
	fmt.Fprintf(&builder, ", TradeDate=%+v", d.TradeDate)
	fmt.Fprintf(&builder, ", TradeTime=%+v", d.TradeTime)
	fmt.Fprintf(&builder, ", BankSerial=%+v", d.BankSerial)
	fmt.Fprintf(&builder, ", TradingDay=%+v", d.TradingDay)
	fmt.Fprintf(&builder, ", PlateSerial=%+v", d.PlateSerial)
	fmt.Fprintf(&builder, ", LastFragment=%+v", d.LastFragment)
	fmt.Fprintf(&builder, ", SessionID=%+v", d.SessionID)
	fmt.Fprintf(&builder, ", InstallID=%+v", d.InstallID)
	fmt.Fprintf(&builder, ", UserID=%+v", d.UserID)
	fmt.Fprintf(&builder, ", Digest=%+v", d.Digest)
	fmt.Fprintf(&builder, ", CurrencyID=%+v", d.CurrencyID)
	fmt.Fprintf(&builder, ", DeviceID=%+v", d.DeviceID)
	fmt.Fprintf(&builder, ", BrokerIDByBank=%+v", d.BrokerIDByBank)
	fmt.Fprintf(&builder, ", OperNo=%+v", d.OperNo)
	fmt.Fprintf(&builder, ", RequestID=%+v", d.RequestID)
	fmt.Fprintf(&builder, ", TID=%+v", d.TID)
	fmt.Fprintf(&builder, ", ErrorID=%+v", d.ErrorID)
	fmt.Fprintf(&builder, ", ErrorMsg=%+v", d.ErrorMsg)
	fmt.Fprintf(&builder, ", PinKey=%+v", d.PinKey)
	fmt.Fprintf(&builder, ", MacKey=%+v", d.MacKey)

	builder.WriteByte('}')

	return builder.String()
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

func (d CThostFtdcReqFutureSignOutField) Type() string { return "CThostFtdcReqFutureSignOutField" }

func (d CThostFtdcReqFutureSignOutField) String() string {
	var builder strings.Builder
	builder.Grow(431)

	builder.WriteString("CThostFtdcReqFutureSignOutField{")
	fmt.Fprintf(&builder, "TradeCode=%+v", d.TradeCode)
	fmt.Fprintf(&builder, ", BankID=%+v", d.BankID)
	fmt.Fprintf(&builder, ", BankBranchID=%+v", d.BankBranchID)
	fmt.Fprintf(&builder, ", BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", BrokerBranchID=%+v", d.BrokerBranchID)
	fmt.Fprintf(&builder, ", TradeDate=%+v", d.TradeDate)
	fmt.Fprintf(&builder, ", TradeTime=%+v", d.TradeTime)
	fmt.Fprintf(&builder, ", BankSerial=%+v", d.BankSerial)
	fmt.Fprintf(&builder, ", TradingDay=%+v", d.TradingDay)
	fmt.Fprintf(&builder, ", PlateSerial=%+v", d.PlateSerial)
	fmt.Fprintf(&builder, ", LastFragment=%+v", d.LastFragment)
	fmt.Fprintf(&builder, ", SessionID=%+v", d.SessionID)
	fmt.Fprintf(&builder, ", InstallID=%+v", d.InstallID)
	fmt.Fprintf(&builder, ", UserID=%+v", d.UserID)
	fmt.Fprintf(&builder, ", Digest=%+v", d.Digest)
	fmt.Fprintf(&builder, ", CurrencyID=%+v", d.CurrencyID)
	fmt.Fprintf(&builder, ", DeviceID=%+v", d.DeviceID)
	fmt.Fprintf(&builder, ", BrokerIDByBank=%+v", d.BrokerIDByBank)
	fmt.Fprintf(&builder, ", OperNo=%+v", d.OperNo)
	fmt.Fprintf(&builder, ", RequestID=%+v", d.RequestID)
	fmt.Fprintf(&builder, ", TID=%+v", d.TID)

	builder.WriteByte('}')

	return builder.String()
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

func (d CThostFtdcRspFutureSignOutField) Type() string { return "CThostFtdcRspFutureSignOutField" }

func (d CThostFtdcRspFutureSignOutField) String() string {
	var builder strings.Builder
	builder.Grow(466)

	builder.WriteString("CThostFtdcRspFutureSignOutField{")
	fmt.Fprintf(&builder, "TradeCode=%+v", d.TradeCode)
	fmt.Fprintf(&builder, ", BankID=%+v", d.BankID)
	fmt.Fprintf(&builder, ", BankBranchID=%+v", d.BankBranchID)
	fmt.Fprintf(&builder, ", BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", BrokerBranchID=%+v", d.BrokerBranchID)
	fmt.Fprintf(&builder, ", TradeDate=%+v", d.TradeDate)
	fmt.Fprintf(&builder, ", TradeTime=%+v", d.TradeTime)
	fmt.Fprintf(&builder, ", BankSerial=%+v", d.BankSerial)
	fmt.Fprintf(&builder, ", TradingDay=%+v", d.TradingDay)
	fmt.Fprintf(&builder, ", PlateSerial=%+v", d.PlateSerial)
	fmt.Fprintf(&builder, ", LastFragment=%+v", d.LastFragment)
	fmt.Fprintf(&builder, ", SessionID=%+v", d.SessionID)
	fmt.Fprintf(&builder, ", InstallID=%+v", d.InstallID)
	fmt.Fprintf(&builder, ", UserID=%+v", d.UserID)
	fmt.Fprintf(&builder, ", Digest=%+v", d.Digest)
	fmt.Fprintf(&builder, ", CurrencyID=%+v", d.CurrencyID)
	fmt.Fprintf(&builder, ", DeviceID=%+v", d.DeviceID)
	fmt.Fprintf(&builder, ", BrokerIDByBank=%+v", d.BrokerIDByBank)
	fmt.Fprintf(&builder, ", OperNo=%+v", d.OperNo)
	fmt.Fprintf(&builder, ", RequestID=%+v", d.RequestID)
	fmt.Fprintf(&builder, ", TID=%+v", d.TID)
	fmt.Fprintf(&builder, ", ErrorID=%+v", d.ErrorID)
	fmt.Fprintf(&builder, ", ErrorMsg=%+v", d.ErrorMsg)

	builder.WriteByte('}')

	return builder.String()
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
	// 长客户姓名
	LongCustomerName types.TThostFtdcLongIndividualNameType
}

func (d CThostFtdcReqQueryTradeResultBySerialField) Type() string {
	return "CThostFtdcReqQueryTradeResultBySerialField"
}

func (d CThostFtdcReqQueryTradeResultBySerialField) String() string {
	var builder strings.Builder
	builder.Grow(601)

	builder.WriteString("CThostFtdcReqQueryTradeResultBySerialField{")
	fmt.Fprintf(&builder, "TradeCode=%+v", d.TradeCode)
	fmt.Fprintf(&builder, ", BankID=%+v", d.BankID)
	fmt.Fprintf(&builder, ", BankBranchID=%+v", d.BankBranchID)
	fmt.Fprintf(&builder, ", BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", BrokerBranchID=%+v", d.BrokerBranchID)
	fmt.Fprintf(&builder, ", TradeDate=%+v", d.TradeDate)
	fmt.Fprintf(&builder, ", TradeTime=%+v", d.TradeTime)
	fmt.Fprintf(&builder, ", BankSerial=%+v", d.BankSerial)
	fmt.Fprintf(&builder, ", TradingDay=%+v", d.TradingDay)
	fmt.Fprintf(&builder, ", PlateSerial=%+v", d.PlateSerial)
	fmt.Fprintf(&builder, ", LastFragment=%+v", d.LastFragment)
	fmt.Fprintf(&builder, ", SessionID=%+v", d.SessionID)
	fmt.Fprintf(&builder, ", Reference=%+v", d.Reference)
	fmt.Fprintf(&builder, ", RefrenceIssureType=%+v", d.RefrenceIssureType)
	fmt.Fprintf(&builder, ", RefrenceIssure=%+v", d.RefrenceIssure)
	fmt.Fprintf(&builder, ", CustomerName=%+v", d.CustomerName)
	fmt.Fprintf(&builder, ", IdCardType=%+v", d.IdCardType)
	fmt.Fprintf(&builder, ", IdentifiedCardNo=%+v", d.IdentifiedCardNo)
	fmt.Fprintf(&builder, ", CustType=%+v", d.CustType)
	fmt.Fprintf(&builder, ", BankAccount=%+v", d.BankAccount)
	fmt.Fprintf(&builder, ", BankPassWord=%+v", d.BankPassWord)
	fmt.Fprintf(&builder, ", AccountID=%+v", d.AccountID)
	fmt.Fprintf(&builder, ", Password=%+v", d.Password)
	fmt.Fprintf(&builder, ", CurrencyID=%+v", d.CurrencyID)
	fmt.Fprintf(&builder, ", TradeAmount=%+v", d.TradeAmount)
	fmt.Fprintf(&builder, ", Digest=%+v", d.Digest)
	fmt.Fprintf(&builder, ", LongCustomerName=%+v", d.LongCustomerName)

	builder.WriteByte('}')

	return builder.String()
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

func (d CThostFtdcRspQueryTradeResultBySerialField) Type() string {
	return "CThostFtdcRspQueryTradeResultBySerialField"
}

func (d CThostFtdcRspQueryTradeResultBySerialField) String() string {
	var builder strings.Builder
	builder.Grow(588)

	builder.WriteString("CThostFtdcRspQueryTradeResultBySerialField{")
	fmt.Fprintf(&builder, "TradeCode=%+v", d.TradeCode)
	fmt.Fprintf(&builder, ", BankID=%+v", d.BankID)
	fmt.Fprintf(&builder, ", BankBranchID=%+v", d.BankBranchID)
	fmt.Fprintf(&builder, ", BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", BrokerBranchID=%+v", d.BrokerBranchID)
	fmt.Fprintf(&builder, ", TradeDate=%+v", d.TradeDate)
	fmt.Fprintf(&builder, ", TradeTime=%+v", d.TradeTime)
	fmt.Fprintf(&builder, ", BankSerial=%+v", d.BankSerial)
	fmt.Fprintf(&builder, ", TradingDay=%+v", d.TradingDay)
	fmt.Fprintf(&builder, ", PlateSerial=%+v", d.PlateSerial)
	fmt.Fprintf(&builder, ", LastFragment=%+v", d.LastFragment)
	fmt.Fprintf(&builder, ", SessionID=%+v", d.SessionID)
	fmt.Fprintf(&builder, ", ErrorID=%+v", d.ErrorID)
	fmt.Fprintf(&builder, ", ErrorMsg=%+v", d.ErrorMsg)
	fmt.Fprintf(&builder, ", Reference=%+v", d.Reference)
	fmt.Fprintf(&builder, ", RefrenceIssureType=%+v", d.RefrenceIssureType)
	fmt.Fprintf(&builder, ", RefrenceIssure=%+v", d.RefrenceIssure)
	fmt.Fprintf(&builder, ", OriginReturnCode=%+v", d.OriginReturnCode)
	fmt.Fprintf(&builder, ", OriginDescrInfoForReturnCode=%+v", d.OriginDescrInfoForReturnCode)
	fmt.Fprintf(&builder, ", BankAccount=%+v", d.BankAccount)
	fmt.Fprintf(&builder, ", BankPassWord=%+v", d.BankPassWord)
	fmt.Fprintf(&builder, ", AccountID=%+v", d.AccountID)
	fmt.Fprintf(&builder, ", Password=%+v", d.Password)
	fmt.Fprintf(&builder, ", CurrencyID=%+v", d.CurrencyID)
	fmt.Fprintf(&builder, ", TradeAmount=%+v", d.TradeAmount)
	fmt.Fprintf(&builder, ", Digest=%+v", d.Digest)

	builder.WriteByte('}')

	return builder.String()
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

func (d CThostFtdcReqDayEndFileReadyField) Type() string { return "CThostFtdcReqDayEndFileReadyField" }

func (d CThostFtdcReqDayEndFileReadyField) String() string {
	var builder strings.Builder
	builder.Grow(314)

	builder.WriteString("CThostFtdcReqDayEndFileReadyField{")
	fmt.Fprintf(&builder, "TradeCode=%+v", d.TradeCode)
	fmt.Fprintf(&builder, ", BankID=%+v", d.BankID)
	fmt.Fprintf(&builder, ", BankBranchID=%+v", d.BankBranchID)
	fmt.Fprintf(&builder, ", BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", BrokerBranchID=%+v", d.BrokerBranchID)
	fmt.Fprintf(&builder, ", TradeDate=%+v", d.TradeDate)
	fmt.Fprintf(&builder, ", TradeTime=%+v", d.TradeTime)
	fmt.Fprintf(&builder, ", BankSerial=%+v", d.BankSerial)
	fmt.Fprintf(&builder, ", TradingDay=%+v", d.TradingDay)
	fmt.Fprintf(&builder, ", PlateSerial=%+v", d.PlateSerial)
	fmt.Fprintf(&builder, ", LastFragment=%+v", d.LastFragment)
	fmt.Fprintf(&builder, ", SessionID=%+v", d.SessionID)
	fmt.Fprintf(&builder, ", FileBusinessCode=%+v", d.FileBusinessCode)
	fmt.Fprintf(&builder, ", Digest=%+v", d.Digest)

	builder.WriteByte('}')

	return builder.String()
}

// 返回结果
type CThostFtdcReturnResultField struct {
	// 返回代码
	ReturnCode types.TThostFtdcReturnCodeType
	// 返回码描述
	DescrInfoForReturnCode types.TThostFtdcDescrInfoForReturnCodeType
}

func (d CThostFtdcReturnResultField) Type() string { return "CThostFtdcReturnResultField" }

func (d CThostFtdcReturnResultField) String() string {
	var builder strings.Builder
	builder.Grow(79)

	builder.WriteString("CThostFtdcReturnResultField{")
	fmt.Fprintf(&builder, "ReturnCode=%+v", d.ReturnCode)
	fmt.Fprintf(&builder, ", DescrInfoForReturnCode=%+v", d.DescrInfoForReturnCode)

	builder.WriteByte('}')

	return builder.String()
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

func (d CThostFtdcVerifyFuturePasswordField) Type() string {
	return "CThostFtdcVerifyFuturePasswordField"
}

func (d CThostFtdcVerifyFuturePasswordField) String() string {
	var builder strings.Builder
	builder.Grow(406)

	builder.WriteString("CThostFtdcVerifyFuturePasswordField{")
	fmt.Fprintf(&builder, "TradeCode=%+v", d.TradeCode)
	fmt.Fprintf(&builder, ", BankID=%+v", d.BankID)
	fmt.Fprintf(&builder, ", BankBranchID=%+v", d.BankBranchID)
	fmt.Fprintf(&builder, ", BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", BrokerBranchID=%+v", d.BrokerBranchID)
	fmt.Fprintf(&builder, ", TradeDate=%+v", d.TradeDate)
	fmt.Fprintf(&builder, ", TradeTime=%+v", d.TradeTime)
	fmt.Fprintf(&builder, ", BankSerial=%+v", d.BankSerial)
	fmt.Fprintf(&builder, ", TradingDay=%+v", d.TradingDay)
	fmt.Fprintf(&builder, ", PlateSerial=%+v", d.PlateSerial)
	fmt.Fprintf(&builder, ", LastFragment=%+v", d.LastFragment)
	fmt.Fprintf(&builder, ", SessionID=%+v", d.SessionID)
	fmt.Fprintf(&builder, ", AccountID=%+v", d.AccountID)
	fmt.Fprintf(&builder, ", Password=%+v", d.Password)
	fmt.Fprintf(&builder, ", BankAccount=%+v", d.BankAccount)
	fmt.Fprintf(&builder, ", BankPassWord=%+v", d.BankPassWord)
	fmt.Fprintf(&builder, ", InstallID=%+v", d.InstallID)
	fmt.Fprintf(&builder, ", TID=%+v", d.TID)
	fmt.Fprintf(&builder, ", CurrencyID=%+v", d.CurrencyID)

	builder.WriteByte('}')

	return builder.String()
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
	// 长客户姓名
	LongCustomerName types.TThostFtdcLongIndividualNameType
}

func (d CThostFtdcVerifyCustInfoField) Type() string { return "CThostFtdcVerifyCustInfoField" }

func (d CThostFtdcVerifyCustInfoField) String() string {
	var builder strings.Builder
	builder.Grow(141)

	builder.WriteString("CThostFtdcVerifyCustInfoField{")
	fmt.Fprintf(&builder, "CustomerName=%+v", d.CustomerName)
	fmt.Fprintf(&builder, ", IdCardType=%+v", d.IdCardType)
	fmt.Fprintf(&builder, ", IdentifiedCardNo=%+v", d.IdentifiedCardNo)
	fmt.Fprintf(&builder, ", CustType=%+v", d.CustType)
	fmt.Fprintf(&builder, ", LongCustomerName=%+v", d.LongCustomerName)

	builder.WriteByte('}')

	return builder.String()
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
	// 长客户姓名
	LongCustomerName types.TThostFtdcLongIndividualNameType
}

func (d CThostFtdcVerifyFuturePasswordAndCustInfoField) Type() string {
	return "CThostFtdcVerifyFuturePasswordAndCustInfoField"
}

func (d CThostFtdcVerifyFuturePasswordAndCustInfoField) String() string {
	var builder strings.Builder
	builder.Grow(215)

	builder.WriteString("CThostFtdcVerifyFuturePasswordAndCustInfoField{")
	fmt.Fprintf(&builder, "CustomerName=%+v", d.CustomerName)
	fmt.Fprintf(&builder, ", IdCardType=%+v", d.IdCardType)
	fmt.Fprintf(&builder, ", IdentifiedCardNo=%+v", d.IdentifiedCardNo)
	fmt.Fprintf(&builder, ", CustType=%+v", d.CustType)
	fmt.Fprintf(&builder, ", AccountID=%+v", d.AccountID)
	fmt.Fprintf(&builder, ", Password=%+v", d.Password)
	fmt.Fprintf(&builder, ", CurrencyID=%+v", d.CurrencyID)
	fmt.Fprintf(&builder, ", LongCustomerName=%+v", d.LongCustomerName)

	builder.WriteByte('}')

	return builder.String()
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

func (d CThostFtdcDepositResultInformField) Type() string {
	return "CThostFtdcDepositResultInformField"
}

func (d CThostFtdcDepositResultInformField) String() string {
	var builder strings.Builder
	builder.Grow(182)

	builder.WriteString("CThostFtdcDepositResultInformField{")
	fmt.Fprintf(&builder, "DepositSeqNo=%+v", d.DepositSeqNo)
	fmt.Fprintf(&builder, ", BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", InvestorID=%+v", d.InvestorID)
	fmt.Fprintf(&builder, ", Deposit=%+v", d.Deposit)
	fmt.Fprintf(&builder, ", RequestID=%+v", d.RequestID)
	fmt.Fprintf(&builder, ", ReturnCode=%+v", d.ReturnCode)
	fmt.Fprintf(&builder, ", DescrInfoForReturnCode=%+v", d.DescrInfoForReturnCode)

	builder.WriteByte('}')

	return builder.String()
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

func (d CThostFtdcReqSyncKeyField) Type() string { return "CThostFtdcReqSyncKeyField" }

func (d CThostFtdcReqSyncKeyField) String() string {
	var builder strings.Builder
	builder.Grow(406)

	builder.WriteString("CThostFtdcReqSyncKeyField{")
	fmt.Fprintf(&builder, "TradeCode=%+v", d.TradeCode)
	fmt.Fprintf(&builder, ", BankID=%+v", d.BankID)
	fmt.Fprintf(&builder, ", BankBranchID=%+v", d.BankBranchID)
	fmt.Fprintf(&builder, ", BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", BrokerBranchID=%+v", d.BrokerBranchID)
	fmt.Fprintf(&builder, ", TradeDate=%+v", d.TradeDate)
	fmt.Fprintf(&builder, ", TradeTime=%+v", d.TradeTime)
	fmt.Fprintf(&builder, ", BankSerial=%+v", d.BankSerial)
	fmt.Fprintf(&builder, ", TradingDay=%+v", d.TradingDay)
	fmt.Fprintf(&builder, ", PlateSerial=%+v", d.PlateSerial)
	fmt.Fprintf(&builder, ", LastFragment=%+v", d.LastFragment)
	fmt.Fprintf(&builder, ", SessionID=%+v", d.SessionID)
	fmt.Fprintf(&builder, ", InstallID=%+v", d.InstallID)
	fmt.Fprintf(&builder, ", UserID=%+v", d.UserID)
	fmt.Fprintf(&builder, ", Message=%+v", d.Message)
	fmt.Fprintf(&builder, ", DeviceID=%+v", d.DeviceID)
	fmt.Fprintf(&builder, ", BrokerIDByBank=%+v", d.BrokerIDByBank)
	fmt.Fprintf(&builder, ", OperNo=%+v", d.OperNo)
	fmt.Fprintf(&builder, ", RequestID=%+v", d.RequestID)
	fmt.Fprintf(&builder, ", TID=%+v", d.TID)

	builder.WriteByte('}')

	return builder.String()
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

func (d CThostFtdcRspSyncKeyField) Type() string { return "CThostFtdcRspSyncKeyField" }

func (d CThostFtdcRspSyncKeyField) String() string {
	var builder strings.Builder
	builder.Grow(441)

	builder.WriteString("CThostFtdcRspSyncKeyField{")
	fmt.Fprintf(&builder, "TradeCode=%+v", d.TradeCode)
	fmt.Fprintf(&builder, ", BankID=%+v", d.BankID)
	fmt.Fprintf(&builder, ", BankBranchID=%+v", d.BankBranchID)
	fmt.Fprintf(&builder, ", BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", BrokerBranchID=%+v", d.BrokerBranchID)
	fmt.Fprintf(&builder, ", TradeDate=%+v", d.TradeDate)
	fmt.Fprintf(&builder, ", TradeTime=%+v", d.TradeTime)
	fmt.Fprintf(&builder, ", BankSerial=%+v", d.BankSerial)
	fmt.Fprintf(&builder, ", TradingDay=%+v", d.TradingDay)
	fmt.Fprintf(&builder, ", PlateSerial=%+v", d.PlateSerial)
	fmt.Fprintf(&builder, ", LastFragment=%+v", d.LastFragment)
	fmt.Fprintf(&builder, ", SessionID=%+v", d.SessionID)
	fmt.Fprintf(&builder, ", InstallID=%+v", d.InstallID)
	fmt.Fprintf(&builder, ", UserID=%+v", d.UserID)
	fmt.Fprintf(&builder, ", Message=%+v", d.Message)
	fmt.Fprintf(&builder, ", DeviceID=%+v", d.DeviceID)
	fmt.Fprintf(&builder, ", BrokerIDByBank=%+v", d.BrokerIDByBank)
	fmt.Fprintf(&builder, ", OperNo=%+v", d.OperNo)
	fmt.Fprintf(&builder, ", RequestID=%+v", d.RequestID)
	fmt.Fprintf(&builder, ", TID=%+v", d.TID)
	fmt.Fprintf(&builder, ", ErrorID=%+v", d.ErrorID)
	fmt.Fprintf(&builder, ", ErrorMsg=%+v", d.ErrorMsg)

	builder.WriteByte('}')

	return builder.String()
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
	// 长客户姓名
	LongCustomerName types.TThostFtdcLongIndividualNameType
}

func (d CThostFtdcNotifyQueryAccountField) Type() string { return "CThostFtdcNotifyQueryAccountField" }

func (d CThostFtdcNotifyQueryAccountField) String() string {
	var builder strings.Builder
	builder.Grow(865)

	builder.WriteString("CThostFtdcNotifyQueryAccountField{")
	fmt.Fprintf(&builder, "TradeCode=%+v", d.TradeCode)
	fmt.Fprintf(&builder, ", BankID=%+v", d.BankID)
	fmt.Fprintf(&builder, ", BankBranchID=%+v", d.BankBranchID)
	fmt.Fprintf(&builder, ", BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", BrokerBranchID=%+v", d.BrokerBranchID)
	fmt.Fprintf(&builder, ", TradeDate=%+v", d.TradeDate)
	fmt.Fprintf(&builder, ", TradeTime=%+v", d.TradeTime)
	fmt.Fprintf(&builder, ", BankSerial=%+v", d.BankSerial)
	fmt.Fprintf(&builder, ", TradingDay=%+v", d.TradingDay)
	fmt.Fprintf(&builder, ", PlateSerial=%+v", d.PlateSerial)
	fmt.Fprintf(&builder, ", LastFragment=%+v", d.LastFragment)
	fmt.Fprintf(&builder, ", SessionID=%+v", d.SessionID)
	fmt.Fprintf(&builder, ", CustomerName=%+v", d.CustomerName)
	fmt.Fprintf(&builder, ", IdCardType=%+v", d.IdCardType)
	fmt.Fprintf(&builder, ", IdentifiedCardNo=%+v", d.IdentifiedCardNo)
	fmt.Fprintf(&builder, ", CustType=%+v", d.CustType)
	fmt.Fprintf(&builder, ", BankAccount=%+v", d.BankAccount)
	fmt.Fprintf(&builder, ", BankPassWord=%+v", d.BankPassWord)
	fmt.Fprintf(&builder, ", AccountID=%+v", d.AccountID)
	fmt.Fprintf(&builder, ", Password=%+v", d.Password)
	fmt.Fprintf(&builder, ", FutureSerial=%+v", d.FutureSerial)
	fmt.Fprintf(&builder, ", InstallID=%+v", d.InstallID)
	fmt.Fprintf(&builder, ", UserID=%+v", d.UserID)
	fmt.Fprintf(&builder, ", VerifyCertNoFlag=%+v", d.VerifyCertNoFlag)
	fmt.Fprintf(&builder, ", CurrencyID=%+v", d.CurrencyID)
	fmt.Fprintf(&builder, ", Digest=%+v", d.Digest)
	fmt.Fprintf(&builder, ", BankAccType=%+v", d.BankAccType)
	fmt.Fprintf(&builder, ", DeviceID=%+v", d.DeviceID)
	fmt.Fprintf(&builder, ", BankSecuAccType=%+v", d.BankSecuAccType)
	fmt.Fprintf(&builder, ", BrokerIDByBank=%+v", d.BrokerIDByBank)
	fmt.Fprintf(&builder, ", BankSecuAcc=%+v", d.BankSecuAcc)
	fmt.Fprintf(&builder, ", BankPwdFlag=%+v", d.BankPwdFlag)
	fmt.Fprintf(&builder, ", SecuPwdFlag=%+v", d.SecuPwdFlag)
	fmt.Fprintf(&builder, ", OperNo=%+v", d.OperNo)
	fmt.Fprintf(&builder, ", RequestID=%+v", d.RequestID)
	fmt.Fprintf(&builder, ", TID=%+v", d.TID)
	fmt.Fprintf(&builder, ", BankUseAmount=%+v", d.BankUseAmount)
	fmt.Fprintf(&builder, ", BankFetchAmount=%+v", d.BankFetchAmount)
	fmt.Fprintf(&builder, ", ErrorID=%+v", d.ErrorID)
	fmt.Fprintf(&builder, ", ErrorMsg=%+v", d.ErrorMsg)
	fmt.Fprintf(&builder, ", LongCustomerName=%+v", d.LongCustomerName)

	builder.WriteByte('}')

	return builder.String()
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

func (d CThostFtdcTransferSerialField) Type() string { return "CThostFtdcTransferSerialField" }

func (d CThostFtdcTransferSerialField) String() string {
	var builder strings.Builder
	builder.Grow(602)

	builder.WriteString("CThostFtdcTransferSerialField{")
	fmt.Fprintf(&builder, "PlateSerial=%+v", d.PlateSerial)
	fmt.Fprintf(&builder, ", TradeDate=%+v", d.TradeDate)
	fmt.Fprintf(&builder, ", TradingDay=%+v", d.TradingDay)
	fmt.Fprintf(&builder, ", TradeTime=%+v", d.TradeTime)
	fmt.Fprintf(&builder, ", TradeCode=%+v", d.TradeCode)
	fmt.Fprintf(&builder, ", SessionID=%+v", d.SessionID)
	fmt.Fprintf(&builder, ", BankID=%+v", d.BankID)
	fmt.Fprintf(&builder, ", BankBranchID=%+v", d.BankBranchID)
	fmt.Fprintf(&builder, ", BankAccType=%+v", d.BankAccType)
	fmt.Fprintf(&builder, ", BankAccount=%+v", d.BankAccount)
	fmt.Fprintf(&builder, ", BankSerial=%+v", d.BankSerial)
	fmt.Fprintf(&builder, ", BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", BrokerBranchID=%+v", d.BrokerBranchID)
	fmt.Fprintf(&builder, ", FutureAccType=%+v", d.FutureAccType)
	fmt.Fprintf(&builder, ", AccountID=%+v", d.AccountID)
	fmt.Fprintf(&builder, ", InvestorID=%+v", d.InvestorID)
	fmt.Fprintf(&builder, ", FutureSerial=%+v", d.FutureSerial)
	fmt.Fprintf(&builder, ", IdCardType=%+v", d.IdCardType)
	fmt.Fprintf(&builder, ", IdentifiedCardNo=%+v", d.IdentifiedCardNo)
	fmt.Fprintf(&builder, ", CurrencyID=%+v", d.CurrencyID)
	fmt.Fprintf(&builder, ", TradeAmount=%+v", d.TradeAmount)
	fmt.Fprintf(&builder, ", CustFee=%+v", d.CustFee)
	fmt.Fprintf(&builder, ", BrokerFee=%+v", d.BrokerFee)
	fmt.Fprintf(&builder, ", AvailabilityFlag=%+v", d.AvailabilityFlag)
	fmt.Fprintf(&builder, ", OperatorCode=%+v", d.OperatorCode)
	fmt.Fprintf(&builder, ", BankNewAccount=%+v", d.BankNewAccount)
	fmt.Fprintf(&builder, ", ErrorID=%+v", d.ErrorID)
	fmt.Fprintf(&builder, ", ErrorMsg=%+v", d.ErrorMsg)

	builder.WriteByte('}')

	return builder.String()
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

func (d CThostFtdcQryTransferSerialField) Type() string { return "CThostFtdcQryTransferSerialField" }

func (d CThostFtdcQryTransferSerialField) String() string {
	var builder strings.Builder
	builder.Grow(105)

	builder.WriteString("CThostFtdcQryTransferSerialField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", AccountID=%+v", d.AccountID)
	fmt.Fprintf(&builder, ", BankID=%+v", d.BankID)
	fmt.Fprintf(&builder, ", CurrencyID=%+v", d.CurrencyID)

	builder.WriteByte('}')

	return builder.String()
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

func (d CThostFtdcNotifyFutureSignInField) Type() string { return "CThostFtdcNotifyFutureSignInField" }

func (d CThostFtdcNotifyFutureSignInField) String() string {
	var builder strings.Builder
	builder.Grow(500)

	builder.WriteString("CThostFtdcNotifyFutureSignInField{")
	fmt.Fprintf(&builder, "TradeCode=%+v", d.TradeCode)
	fmt.Fprintf(&builder, ", BankID=%+v", d.BankID)
	fmt.Fprintf(&builder, ", BankBranchID=%+v", d.BankBranchID)
	fmt.Fprintf(&builder, ", BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", BrokerBranchID=%+v", d.BrokerBranchID)
	fmt.Fprintf(&builder, ", TradeDate=%+v", d.TradeDate)
	fmt.Fprintf(&builder, ", TradeTime=%+v", d.TradeTime)
	fmt.Fprintf(&builder, ", BankSerial=%+v", d.BankSerial)
	fmt.Fprintf(&builder, ", TradingDay=%+v", d.TradingDay)
	fmt.Fprintf(&builder, ", PlateSerial=%+v", d.PlateSerial)
	fmt.Fprintf(&builder, ", LastFragment=%+v", d.LastFragment)
	fmt.Fprintf(&builder, ", SessionID=%+v", d.SessionID)
	fmt.Fprintf(&builder, ", InstallID=%+v", d.InstallID)
	fmt.Fprintf(&builder, ", UserID=%+v", d.UserID)
	fmt.Fprintf(&builder, ", Digest=%+v", d.Digest)
	fmt.Fprintf(&builder, ", CurrencyID=%+v", d.CurrencyID)
	fmt.Fprintf(&builder, ", DeviceID=%+v", d.DeviceID)
	fmt.Fprintf(&builder, ", BrokerIDByBank=%+v", d.BrokerIDByBank)
	fmt.Fprintf(&builder, ", OperNo=%+v", d.OperNo)
	fmt.Fprintf(&builder, ", RequestID=%+v", d.RequestID)
	fmt.Fprintf(&builder, ", TID=%+v", d.TID)
	fmt.Fprintf(&builder, ", ErrorID=%+v", d.ErrorID)
	fmt.Fprintf(&builder, ", ErrorMsg=%+v", d.ErrorMsg)
	fmt.Fprintf(&builder, ", PinKey=%+v", d.PinKey)
	fmt.Fprintf(&builder, ", MacKey=%+v", d.MacKey)

	builder.WriteByte('}')

	return builder.String()
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

func (d CThostFtdcNotifyFutureSignOutField) Type() string {
	return "CThostFtdcNotifyFutureSignOutField"
}

func (d CThostFtdcNotifyFutureSignOutField) String() string {
	var builder strings.Builder
	builder.Grow(469)

	builder.WriteString("CThostFtdcNotifyFutureSignOutField{")
	fmt.Fprintf(&builder, "TradeCode=%+v", d.TradeCode)
	fmt.Fprintf(&builder, ", BankID=%+v", d.BankID)
	fmt.Fprintf(&builder, ", BankBranchID=%+v", d.BankBranchID)
	fmt.Fprintf(&builder, ", BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", BrokerBranchID=%+v", d.BrokerBranchID)
	fmt.Fprintf(&builder, ", TradeDate=%+v", d.TradeDate)
	fmt.Fprintf(&builder, ", TradeTime=%+v", d.TradeTime)
	fmt.Fprintf(&builder, ", BankSerial=%+v", d.BankSerial)
	fmt.Fprintf(&builder, ", TradingDay=%+v", d.TradingDay)
	fmt.Fprintf(&builder, ", PlateSerial=%+v", d.PlateSerial)
	fmt.Fprintf(&builder, ", LastFragment=%+v", d.LastFragment)
	fmt.Fprintf(&builder, ", SessionID=%+v", d.SessionID)
	fmt.Fprintf(&builder, ", InstallID=%+v", d.InstallID)
	fmt.Fprintf(&builder, ", UserID=%+v", d.UserID)
	fmt.Fprintf(&builder, ", Digest=%+v", d.Digest)
	fmt.Fprintf(&builder, ", CurrencyID=%+v", d.CurrencyID)
	fmt.Fprintf(&builder, ", DeviceID=%+v", d.DeviceID)
	fmt.Fprintf(&builder, ", BrokerIDByBank=%+v", d.BrokerIDByBank)
	fmt.Fprintf(&builder, ", OperNo=%+v", d.OperNo)
	fmt.Fprintf(&builder, ", RequestID=%+v", d.RequestID)
	fmt.Fprintf(&builder, ", TID=%+v", d.TID)
	fmt.Fprintf(&builder, ", ErrorID=%+v", d.ErrorID)
	fmt.Fprintf(&builder, ", ErrorMsg=%+v", d.ErrorMsg)

	builder.WriteByte('}')

	return builder.String()
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

func (d CThostFtdcNotifySyncKeyField) Type() string { return "CThostFtdcNotifySyncKeyField" }

func (d CThostFtdcNotifySyncKeyField) String() string {
	var builder strings.Builder
	builder.Grow(444)

	builder.WriteString("CThostFtdcNotifySyncKeyField{")
	fmt.Fprintf(&builder, "TradeCode=%+v", d.TradeCode)
	fmt.Fprintf(&builder, ", BankID=%+v", d.BankID)
	fmt.Fprintf(&builder, ", BankBranchID=%+v", d.BankBranchID)
	fmt.Fprintf(&builder, ", BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", BrokerBranchID=%+v", d.BrokerBranchID)
	fmt.Fprintf(&builder, ", TradeDate=%+v", d.TradeDate)
	fmt.Fprintf(&builder, ", TradeTime=%+v", d.TradeTime)
	fmt.Fprintf(&builder, ", BankSerial=%+v", d.BankSerial)
	fmt.Fprintf(&builder, ", TradingDay=%+v", d.TradingDay)
	fmt.Fprintf(&builder, ", PlateSerial=%+v", d.PlateSerial)
	fmt.Fprintf(&builder, ", LastFragment=%+v", d.LastFragment)
	fmt.Fprintf(&builder, ", SessionID=%+v", d.SessionID)
	fmt.Fprintf(&builder, ", InstallID=%+v", d.InstallID)
	fmt.Fprintf(&builder, ", UserID=%+v", d.UserID)
	fmt.Fprintf(&builder, ", Message=%+v", d.Message)
	fmt.Fprintf(&builder, ", DeviceID=%+v", d.DeviceID)
	fmt.Fprintf(&builder, ", BrokerIDByBank=%+v", d.BrokerIDByBank)
	fmt.Fprintf(&builder, ", OperNo=%+v", d.OperNo)
	fmt.Fprintf(&builder, ", RequestID=%+v", d.RequestID)
	fmt.Fprintf(&builder, ", TID=%+v", d.TID)
	fmt.Fprintf(&builder, ", ErrorID=%+v", d.ErrorID)
	fmt.Fprintf(&builder, ", ErrorMsg=%+v", d.ErrorMsg)

	builder.WriteByte('}')

	return builder.String()
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

func (d CThostFtdcQryAccountregisterField) Type() string { return "CThostFtdcQryAccountregisterField" }

func (d CThostFtdcQryAccountregisterField) String() string {
	var builder strings.Builder
	builder.Grow(128)

	builder.WriteString("CThostFtdcQryAccountregisterField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", AccountID=%+v", d.AccountID)
	fmt.Fprintf(&builder, ", BankID=%+v", d.BankID)
	fmt.Fprintf(&builder, ", BankBranchID=%+v", d.BankBranchID)
	fmt.Fprintf(&builder, ", CurrencyID=%+v", d.CurrencyID)

	builder.WriteByte('}')

	return builder.String()
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
	// 长客户姓名
	LongCustomerName types.TThostFtdcLongIndividualNameType
}

func (d CThostFtdcAccountregisterField) Type() string { return "CThostFtdcAccountregisterField" }

func (d CThostFtdcAccountregisterField) String() string {
	var builder strings.Builder
	builder.Grow(391)

	builder.WriteString("CThostFtdcAccountregisterField{")
	fmt.Fprintf(&builder, "TradeDay=%+v", d.TradeDay)
	fmt.Fprintf(&builder, ", BankID=%+v", d.BankID)
	fmt.Fprintf(&builder, ", BankBranchID=%+v", d.BankBranchID)
	fmt.Fprintf(&builder, ", BankAccount=%+v", d.BankAccount)
	fmt.Fprintf(&builder, ", BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", BrokerBranchID=%+v", d.BrokerBranchID)
	fmt.Fprintf(&builder, ", AccountID=%+v", d.AccountID)
	fmt.Fprintf(&builder, ", IdCardType=%+v", d.IdCardType)
	fmt.Fprintf(&builder, ", IdentifiedCardNo=%+v", d.IdentifiedCardNo)
	fmt.Fprintf(&builder, ", CustomerName=%+v", d.CustomerName)
	fmt.Fprintf(&builder, ", CurrencyID=%+v", d.CurrencyID)
	fmt.Fprintf(&builder, ", OpenOrDestroy=%+v", d.OpenOrDestroy)
	fmt.Fprintf(&builder, ", RegDate=%+v", d.RegDate)
	fmt.Fprintf(&builder, ", OutDate=%+v", d.OutDate)
	fmt.Fprintf(&builder, ", TID=%+v", d.TID)
	fmt.Fprintf(&builder, ", CustType=%+v", d.CustType)
	fmt.Fprintf(&builder, ", BankAccType=%+v", d.BankAccType)
	fmt.Fprintf(&builder, ", LongCustomerName=%+v", d.LongCustomerName)

	builder.WriteByte('}')

	return builder.String()
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
	// 长客户姓名
	LongCustomerName types.TThostFtdcLongIndividualNameType
}

func (d CThostFtdcOpenAccountField) Type() string { return "CThostFtdcOpenAccountField" }

func (d CThostFtdcOpenAccountField) String() string {
	var builder strings.Builder
	builder.Grow(962)

	builder.WriteString("CThostFtdcOpenAccountField{")
	fmt.Fprintf(&builder, "TradeCode=%+v", d.TradeCode)
	fmt.Fprintf(&builder, ", BankID=%+v", d.BankID)
	fmt.Fprintf(&builder, ", BankBranchID=%+v", d.BankBranchID)
	fmt.Fprintf(&builder, ", BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", BrokerBranchID=%+v", d.BrokerBranchID)
	fmt.Fprintf(&builder, ", TradeDate=%+v", d.TradeDate)
	fmt.Fprintf(&builder, ", TradeTime=%+v", d.TradeTime)
	fmt.Fprintf(&builder, ", BankSerial=%+v", d.BankSerial)
	fmt.Fprintf(&builder, ", TradingDay=%+v", d.TradingDay)
	fmt.Fprintf(&builder, ", PlateSerial=%+v", d.PlateSerial)
	fmt.Fprintf(&builder, ", LastFragment=%+v", d.LastFragment)
	fmt.Fprintf(&builder, ", SessionID=%+v", d.SessionID)
	fmt.Fprintf(&builder, ", CustomerName=%+v", d.CustomerName)
	fmt.Fprintf(&builder, ", IdCardType=%+v", d.IdCardType)
	fmt.Fprintf(&builder, ", IdentifiedCardNo=%+v", d.IdentifiedCardNo)
	fmt.Fprintf(&builder, ", Gender=%+v", d.Gender)
	fmt.Fprintf(&builder, ", CountryCode=%+v", d.CountryCode)
	fmt.Fprintf(&builder, ", CustType=%+v", d.CustType)
	fmt.Fprintf(&builder, ", Address=%+v", d.Address)
	fmt.Fprintf(&builder, ", ZipCode=%+v", d.ZipCode)
	fmt.Fprintf(&builder, ", Telephone=%+v", d.Telephone)
	fmt.Fprintf(&builder, ", MobilePhone=%+v", d.MobilePhone)
	fmt.Fprintf(&builder, ", Fax=%+v", d.Fax)
	fmt.Fprintf(&builder, ", EMail=%+v", d.EMail)
	fmt.Fprintf(&builder, ", MoneyAccountStatus=%+v", d.MoneyAccountStatus)
	fmt.Fprintf(&builder, ", BankAccount=%+v", d.BankAccount)
	fmt.Fprintf(&builder, ", BankPassWord=%+v", d.BankPassWord)
	fmt.Fprintf(&builder, ", AccountID=%+v", d.AccountID)
	fmt.Fprintf(&builder, ", Password=%+v", d.Password)
	fmt.Fprintf(&builder, ", InstallID=%+v", d.InstallID)
	fmt.Fprintf(&builder, ", VerifyCertNoFlag=%+v", d.VerifyCertNoFlag)
	fmt.Fprintf(&builder, ", CurrencyID=%+v", d.CurrencyID)
	fmt.Fprintf(&builder, ", CashExchangeCode=%+v", d.CashExchangeCode)
	fmt.Fprintf(&builder, ", Digest=%+v", d.Digest)
	fmt.Fprintf(&builder, ", BankAccType=%+v", d.BankAccType)
	fmt.Fprintf(&builder, ", DeviceID=%+v", d.DeviceID)
	fmt.Fprintf(&builder, ", BankSecuAccType=%+v", d.BankSecuAccType)
	fmt.Fprintf(&builder, ", BrokerIDByBank=%+v", d.BrokerIDByBank)
	fmt.Fprintf(&builder, ", BankSecuAcc=%+v", d.BankSecuAcc)
	fmt.Fprintf(&builder, ", BankPwdFlag=%+v", d.BankPwdFlag)
	fmt.Fprintf(&builder, ", SecuPwdFlag=%+v", d.SecuPwdFlag)
	fmt.Fprintf(&builder, ", OperNo=%+v", d.OperNo)
	fmt.Fprintf(&builder, ", TID=%+v", d.TID)
	fmt.Fprintf(&builder, ", UserID=%+v", d.UserID)
	fmt.Fprintf(&builder, ", ErrorID=%+v", d.ErrorID)
	fmt.Fprintf(&builder, ", ErrorMsg=%+v", d.ErrorMsg)
	fmt.Fprintf(&builder, ", LongCustomerName=%+v", d.LongCustomerName)

	builder.WriteByte('}')

	return builder.String()
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
	// 长客户姓名
	LongCustomerName types.TThostFtdcLongIndividualNameType
}

func (d CThostFtdcCancelAccountField) Type() string { return "CThostFtdcCancelAccountField" }

func (d CThostFtdcCancelAccountField) String() string {
	var builder strings.Builder
	builder.Grow(964)

	builder.WriteString("CThostFtdcCancelAccountField{")
	fmt.Fprintf(&builder, "TradeCode=%+v", d.TradeCode)
	fmt.Fprintf(&builder, ", BankID=%+v", d.BankID)
	fmt.Fprintf(&builder, ", BankBranchID=%+v", d.BankBranchID)
	fmt.Fprintf(&builder, ", BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", BrokerBranchID=%+v", d.BrokerBranchID)
	fmt.Fprintf(&builder, ", TradeDate=%+v", d.TradeDate)
	fmt.Fprintf(&builder, ", TradeTime=%+v", d.TradeTime)
	fmt.Fprintf(&builder, ", BankSerial=%+v", d.BankSerial)
	fmt.Fprintf(&builder, ", TradingDay=%+v", d.TradingDay)
	fmt.Fprintf(&builder, ", PlateSerial=%+v", d.PlateSerial)
	fmt.Fprintf(&builder, ", LastFragment=%+v", d.LastFragment)
	fmt.Fprintf(&builder, ", SessionID=%+v", d.SessionID)
	fmt.Fprintf(&builder, ", CustomerName=%+v", d.CustomerName)
	fmt.Fprintf(&builder, ", IdCardType=%+v", d.IdCardType)
	fmt.Fprintf(&builder, ", IdentifiedCardNo=%+v", d.IdentifiedCardNo)
	fmt.Fprintf(&builder, ", Gender=%+v", d.Gender)
	fmt.Fprintf(&builder, ", CountryCode=%+v", d.CountryCode)
	fmt.Fprintf(&builder, ", CustType=%+v", d.CustType)
	fmt.Fprintf(&builder, ", Address=%+v", d.Address)
	fmt.Fprintf(&builder, ", ZipCode=%+v", d.ZipCode)
	fmt.Fprintf(&builder, ", Telephone=%+v", d.Telephone)
	fmt.Fprintf(&builder, ", MobilePhone=%+v", d.MobilePhone)
	fmt.Fprintf(&builder, ", Fax=%+v", d.Fax)
	fmt.Fprintf(&builder, ", EMail=%+v", d.EMail)
	fmt.Fprintf(&builder, ", MoneyAccountStatus=%+v", d.MoneyAccountStatus)
	fmt.Fprintf(&builder, ", BankAccount=%+v", d.BankAccount)
	fmt.Fprintf(&builder, ", BankPassWord=%+v", d.BankPassWord)
	fmt.Fprintf(&builder, ", AccountID=%+v", d.AccountID)
	fmt.Fprintf(&builder, ", Password=%+v", d.Password)
	fmt.Fprintf(&builder, ", InstallID=%+v", d.InstallID)
	fmt.Fprintf(&builder, ", VerifyCertNoFlag=%+v", d.VerifyCertNoFlag)
	fmt.Fprintf(&builder, ", CurrencyID=%+v", d.CurrencyID)
	fmt.Fprintf(&builder, ", CashExchangeCode=%+v", d.CashExchangeCode)
	fmt.Fprintf(&builder, ", Digest=%+v", d.Digest)
	fmt.Fprintf(&builder, ", BankAccType=%+v", d.BankAccType)
	fmt.Fprintf(&builder, ", DeviceID=%+v", d.DeviceID)
	fmt.Fprintf(&builder, ", BankSecuAccType=%+v", d.BankSecuAccType)
	fmt.Fprintf(&builder, ", BrokerIDByBank=%+v", d.BrokerIDByBank)
	fmt.Fprintf(&builder, ", BankSecuAcc=%+v", d.BankSecuAcc)
	fmt.Fprintf(&builder, ", BankPwdFlag=%+v", d.BankPwdFlag)
	fmt.Fprintf(&builder, ", SecuPwdFlag=%+v", d.SecuPwdFlag)
	fmt.Fprintf(&builder, ", OperNo=%+v", d.OperNo)
	fmt.Fprintf(&builder, ", TID=%+v", d.TID)
	fmt.Fprintf(&builder, ", UserID=%+v", d.UserID)
	fmt.Fprintf(&builder, ", ErrorID=%+v", d.ErrorID)
	fmt.Fprintf(&builder, ", ErrorMsg=%+v", d.ErrorMsg)
	fmt.Fprintf(&builder, ", LongCustomerName=%+v", d.LongCustomerName)

	builder.WriteByte('}')

	return builder.String()
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
	// 长客户姓名
	LongCustomerName types.TThostFtdcLongIndividualNameType
}

func (d CThostFtdcChangeAccountField) Type() string { return "CThostFtdcChangeAccountField" }

func (d CThostFtdcChangeAccountField) String() string {
	var builder strings.Builder
	builder.Grow(891)

	builder.WriteString("CThostFtdcChangeAccountField{")
	fmt.Fprintf(&builder, "TradeCode=%+v", d.TradeCode)
	fmt.Fprintf(&builder, ", BankID=%+v", d.BankID)
	fmt.Fprintf(&builder, ", BankBranchID=%+v", d.BankBranchID)
	fmt.Fprintf(&builder, ", BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", BrokerBranchID=%+v", d.BrokerBranchID)
	fmt.Fprintf(&builder, ", TradeDate=%+v", d.TradeDate)
	fmt.Fprintf(&builder, ", TradeTime=%+v", d.TradeTime)
	fmt.Fprintf(&builder, ", BankSerial=%+v", d.BankSerial)
	fmt.Fprintf(&builder, ", TradingDay=%+v", d.TradingDay)
	fmt.Fprintf(&builder, ", PlateSerial=%+v", d.PlateSerial)
	fmt.Fprintf(&builder, ", LastFragment=%+v", d.LastFragment)
	fmt.Fprintf(&builder, ", SessionID=%+v", d.SessionID)
	fmt.Fprintf(&builder, ", CustomerName=%+v", d.CustomerName)
	fmt.Fprintf(&builder, ", IdCardType=%+v", d.IdCardType)
	fmt.Fprintf(&builder, ", IdentifiedCardNo=%+v", d.IdentifiedCardNo)
	fmt.Fprintf(&builder, ", Gender=%+v", d.Gender)
	fmt.Fprintf(&builder, ", CountryCode=%+v", d.CountryCode)
	fmt.Fprintf(&builder, ", CustType=%+v", d.CustType)
	fmt.Fprintf(&builder, ", Address=%+v", d.Address)
	fmt.Fprintf(&builder, ", ZipCode=%+v", d.ZipCode)
	fmt.Fprintf(&builder, ", Telephone=%+v", d.Telephone)
	fmt.Fprintf(&builder, ", MobilePhone=%+v", d.MobilePhone)
	fmt.Fprintf(&builder, ", Fax=%+v", d.Fax)
	fmt.Fprintf(&builder, ", EMail=%+v", d.EMail)
	fmt.Fprintf(&builder, ", MoneyAccountStatus=%+v", d.MoneyAccountStatus)
	fmt.Fprintf(&builder, ", BankAccount=%+v", d.BankAccount)
	fmt.Fprintf(&builder, ", BankPassWord=%+v", d.BankPassWord)
	fmt.Fprintf(&builder, ", NewBankAccount=%+v", d.NewBankAccount)
	fmt.Fprintf(&builder, ", NewBankPassWord=%+v", d.NewBankPassWord)
	fmt.Fprintf(&builder, ", AccountID=%+v", d.AccountID)
	fmt.Fprintf(&builder, ", Password=%+v", d.Password)
	fmt.Fprintf(&builder, ", BankAccType=%+v", d.BankAccType)
	fmt.Fprintf(&builder, ", InstallID=%+v", d.InstallID)
	fmt.Fprintf(&builder, ", VerifyCertNoFlag=%+v", d.VerifyCertNoFlag)
	fmt.Fprintf(&builder, ", CurrencyID=%+v", d.CurrencyID)
	fmt.Fprintf(&builder, ", BrokerIDByBank=%+v", d.BrokerIDByBank)
	fmt.Fprintf(&builder, ", BankPwdFlag=%+v", d.BankPwdFlag)
	fmt.Fprintf(&builder, ", SecuPwdFlag=%+v", d.SecuPwdFlag)
	fmt.Fprintf(&builder, ", TID=%+v", d.TID)
	fmt.Fprintf(&builder, ", Digest=%+v", d.Digest)
	fmt.Fprintf(&builder, ", ErrorID=%+v", d.ErrorID)
	fmt.Fprintf(&builder, ", ErrorMsg=%+v", d.ErrorMsg)
	fmt.Fprintf(&builder, ", LongCustomerName=%+v", d.LongCustomerName)

	builder.WriteByte('}')

	return builder.String()
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

func (d CThostFtdcSecAgentACIDMapField) Type() string { return "CThostFtdcSecAgentACIDMapField" }

func (d CThostFtdcSecAgentACIDMapField) String() string {
	var builder strings.Builder
	builder.Grow(129)

	builder.WriteString("CThostFtdcSecAgentACIDMapField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", UserID=%+v", d.UserID)
	fmt.Fprintf(&builder, ", AccountID=%+v", d.AccountID)
	fmt.Fprintf(&builder, ", CurrencyID=%+v", d.CurrencyID)
	fmt.Fprintf(&builder, ", BrokerSecAgentID=%+v", d.BrokerSecAgentID)

	builder.WriteByte('}')

	return builder.String()
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

func (d CThostFtdcQrySecAgentACIDMapField) Type() string { return "CThostFtdcQrySecAgentACIDMapField" }

func (d CThostFtdcQrySecAgentACIDMapField) String() string {
	var builder strings.Builder
	builder.Grow(106)

	builder.WriteString("CThostFtdcQrySecAgentACIDMapField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", UserID=%+v", d.UserID)
	fmt.Fprintf(&builder, ", AccountID=%+v", d.AccountID)
	fmt.Fprintf(&builder, ", CurrencyID=%+v", d.CurrencyID)

	builder.WriteByte('}')

	return builder.String()
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

func (d CThostFtdcUserRightsAssignField) Type() string { return "CThostFtdcUserRightsAssignField" }

func (d CThostFtdcUserRightsAssignField) String() string {
	var builder strings.Builder
	builder.Grow(87)

	builder.WriteString("CThostFtdcUserRightsAssignField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", UserID=%+v", d.UserID)
	fmt.Fprintf(&builder, ", DRIdentityID=%+v", d.DRIdentityID)

	builder.WriteByte('}')

	return builder.String()
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

func (d CThostFtdcBrokerUserRightAssignField) Type() string {
	return "CThostFtdcBrokerUserRightAssignField"
}

func (d CThostFtdcBrokerUserRightAssignField) String() string {
	var builder strings.Builder
	builder.Grow(95)

	builder.WriteString("CThostFtdcBrokerUserRightAssignField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", DRIdentityID=%+v", d.DRIdentityID)
	fmt.Fprintf(&builder, ", Tradeable=%+v", d.Tradeable)

	builder.WriteByte('}')

	return builder.String()
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

func (d CThostFtdcDRTransferField) Type() string { return "CThostFtdcDRTransferField" }

func (d CThostFtdcDRTransferField) String() string {
	var builder strings.Builder
	builder.Grow(121)

	builder.WriteString("CThostFtdcDRTransferField{")
	fmt.Fprintf(&builder, "OrigDRIdentityID=%+v", d.OrigDRIdentityID)
	fmt.Fprintf(&builder, ", DestDRIdentityID=%+v", d.DestDRIdentityID)
	fmt.Fprintf(&builder, ", OrigBrokerID=%+v", d.OrigBrokerID)
	fmt.Fprintf(&builder, ", DestBrokerID=%+v", d.DestBrokerID)

	builder.WriteByte('}')

	return builder.String()
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

func (d CThostFtdcFensUserInfoField) Type() string { return "CThostFtdcFensUserInfoField" }

func (d CThostFtdcFensUserInfoField) String() string {
	var builder strings.Builder
	builder.Grow(80)

	builder.WriteString("CThostFtdcFensUserInfoField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", UserID=%+v", d.UserID)
	fmt.Fprintf(&builder, ", LoginMode=%+v", d.LoginMode)

	builder.WriteByte('}')

	return builder.String()
}

// 当前银期所属交易中心
type CThostFtdcCurrTransferIdentityField struct {
	// 交易中心代码
	IdentityID types.TThostFtdcDRIdentityIDType
}

func (d CThostFtdcCurrTransferIdentityField) Type() string {
	return "CThostFtdcCurrTransferIdentityField"
}

func (d CThostFtdcCurrTransferIdentityField) String() string {
	var builder strings.Builder
	builder.Grow(55)

	builder.WriteString("CThostFtdcCurrTransferIdentityField{")
	fmt.Fprintf(&builder, "IdentityID=%+v", d.IdentityID)

	builder.WriteByte('}')

	return builder.String()
}

// 禁止登录用户
type CThostFtdcLoginForbiddenUserField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 用户代码
	UserID types.TThostFtdcUserIDType
	// 保留的无效字段
	reserve1 types.TThostFtdcOldIPAddressType
	// IP地址
	IPAddress types.TThostFtdcIPAddressType
}

func (d CThostFtdcLoginForbiddenUserField) Type() string { return "CThostFtdcLoginForbiddenUserField" }

func (d CThostFtdcLoginForbiddenUserField) String() string {
	var builder strings.Builder
	builder.Grow(104)

	builder.WriteString("CThostFtdcLoginForbiddenUserField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", UserID=%+v", d.UserID)
	fmt.Fprintf(&builder, ", reserve1=%+v", d.reserve1)
	fmt.Fprintf(&builder, ", IPAddress=%+v", d.IPAddress)

	builder.WriteByte('}')

	return builder.String()
}

// 查询禁止登录用户
type CThostFtdcQryLoginForbiddenUserField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 用户代码
	UserID types.TThostFtdcUserIDType
}

func (d CThostFtdcQryLoginForbiddenUserField) Type() string {
	return "CThostFtdcQryLoginForbiddenUserField"
}

func (d CThostFtdcQryLoginForbiddenUserField) String() string {
	var builder strings.Builder
	builder.Grow(70)

	builder.WriteString("CThostFtdcQryLoginForbiddenUserField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", UserID=%+v", d.UserID)

	builder.WriteByte('}')

	return builder.String()
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

func (d CThostFtdcTradingAccountReserveField) Type() string {
	return "CThostFtdcTradingAccountReserveField"
}

func (d CThostFtdcTradingAccountReserveField) String() string {
	var builder strings.Builder
	builder.Grow(110)

	builder.WriteString("CThostFtdcTradingAccountReserveField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", AccountID=%+v", d.AccountID)
	fmt.Fprintf(&builder, ", Reserve=%+v", d.Reserve)
	fmt.Fprintf(&builder, ", CurrencyID=%+v", d.CurrencyID)

	builder.WriteByte('}')

	return builder.String()
}

// 查询禁止登录IP
type CThostFtdcQryLoginForbiddenIPField struct {
	// 保留的无效字段
	reserve1 types.TThostFtdcOldIPAddressType
	// IP地址
	IPAddress types.TThostFtdcIPAddressType
}

func (d CThostFtdcQryLoginForbiddenIPField) Type() string {
	return "CThostFtdcQryLoginForbiddenIPField"
}

func (d CThostFtdcQryLoginForbiddenIPField) String() string {
	var builder strings.Builder
	builder.Grow(71)

	builder.WriteString("CThostFtdcQryLoginForbiddenIPField{")
	fmt.Fprintf(&builder, "reserve1=%+v", d.reserve1)
	fmt.Fprintf(&builder, ", IPAddress=%+v", d.IPAddress)

	builder.WriteByte('}')

	return builder.String()
}

// 查询IP列表
type CThostFtdcQryIPListField struct {
	// 保留的无效字段
	reserve1 types.TThostFtdcOldIPAddressType
	// IP地址
	IPAddress types.TThostFtdcIPAddressType
}

func (d CThostFtdcQryIPListField) Type() string { return "CThostFtdcQryIPListField" }

func (d CThostFtdcQryIPListField) String() string {
	var builder strings.Builder
	builder.Grow(61)

	builder.WriteString("CThostFtdcQryIPListField{")
	fmt.Fprintf(&builder, "reserve1=%+v", d.reserve1)
	fmt.Fprintf(&builder, ", IPAddress=%+v", d.IPAddress)

	builder.WriteByte('}')

	return builder.String()
}

// 查询用户下单权限分配表
type CThostFtdcQryUserRightsAssignField struct {
	// 应用单元代码
	BrokerID types.TThostFtdcBrokerIDType
	// 用户代码
	UserID types.TThostFtdcUserIDType
}

func (d CThostFtdcQryUserRightsAssignField) Type() string {
	return "CThostFtdcQryUserRightsAssignField"
}

func (d CThostFtdcQryUserRightsAssignField) String() string {
	var builder strings.Builder
	builder.Grow(68)

	builder.WriteString("CThostFtdcQryUserRightsAssignField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", UserID=%+v", d.UserID)

	builder.WriteByte('}')

	return builder.String()
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
	CustomerName types.TThostFtdcLongIndividualNameType
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

func (d CThostFtdcReserveOpenAccountConfirmField) Type() string {
	return "CThostFtdcReserveOpenAccountConfirmField"
}

func (d CThostFtdcReserveOpenAccountConfirmField) String() string {
	var builder strings.Builder
	builder.Grow(849)

	builder.WriteString("CThostFtdcReserveOpenAccountConfirmField{")
	fmt.Fprintf(&builder, "TradeCode=%+v", d.TradeCode)
	fmt.Fprintf(&builder, ", BankID=%+v", d.BankID)
	fmt.Fprintf(&builder, ", BankBranchID=%+v", d.BankBranchID)
	fmt.Fprintf(&builder, ", BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", BrokerBranchID=%+v", d.BrokerBranchID)
	fmt.Fprintf(&builder, ", TradeDate=%+v", d.TradeDate)
	fmt.Fprintf(&builder, ", TradeTime=%+v", d.TradeTime)
	fmt.Fprintf(&builder, ", BankSerial=%+v", d.BankSerial)
	fmt.Fprintf(&builder, ", TradingDay=%+v", d.TradingDay)
	fmt.Fprintf(&builder, ", PlateSerial=%+v", d.PlateSerial)
	fmt.Fprintf(&builder, ", LastFragment=%+v", d.LastFragment)
	fmt.Fprintf(&builder, ", SessionID=%+v", d.SessionID)
	fmt.Fprintf(&builder, ", CustomerName=%+v", d.CustomerName)
	fmt.Fprintf(&builder, ", IdCardType=%+v", d.IdCardType)
	fmt.Fprintf(&builder, ", IdentifiedCardNo=%+v", d.IdentifiedCardNo)
	fmt.Fprintf(&builder, ", Gender=%+v", d.Gender)
	fmt.Fprintf(&builder, ", CountryCode=%+v", d.CountryCode)
	fmt.Fprintf(&builder, ", CustType=%+v", d.CustType)
	fmt.Fprintf(&builder, ", Address=%+v", d.Address)
	fmt.Fprintf(&builder, ", ZipCode=%+v", d.ZipCode)
	fmt.Fprintf(&builder, ", Telephone=%+v", d.Telephone)
	fmt.Fprintf(&builder, ", MobilePhone=%+v", d.MobilePhone)
	fmt.Fprintf(&builder, ", Fax=%+v", d.Fax)
	fmt.Fprintf(&builder, ", EMail=%+v", d.EMail)
	fmt.Fprintf(&builder, ", MoneyAccountStatus=%+v", d.MoneyAccountStatus)
	fmt.Fprintf(&builder, ", BankAccount=%+v", d.BankAccount)
	fmt.Fprintf(&builder, ", BankPassWord=%+v", d.BankPassWord)
	fmt.Fprintf(&builder, ", InstallID=%+v", d.InstallID)
	fmt.Fprintf(&builder, ", VerifyCertNoFlag=%+v", d.VerifyCertNoFlag)
	fmt.Fprintf(&builder, ", CurrencyID=%+v", d.CurrencyID)
	fmt.Fprintf(&builder, ", Digest=%+v", d.Digest)
	fmt.Fprintf(&builder, ", BankAccType=%+v", d.BankAccType)
	fmt.Fprintf(&builder, ", BrokerIDByBank=%+v", d.BrokerIDByBank)
	fmt.Fprintf(&builder, ", TID=%+v", d.TID)
	fmt.Fprintf(&builder, ", AccountID=%+v", d.AccountID)
	fmt.Fprintf(&builder, ", Password=%+v", d.Password)
	fmt.Fprintf(&builder, ", BankReserveOpenSeq=%+v", d.BankReserveOpenSeq)
	fmt.Fprintf(&builder, ", BookDate=%+v", d.BookDate)
	fmt.Fprintf(&builder, ", BookPsw=%+v", d.BookPsw)
	fmt.Fprintf(&builder, ", ErrorID=%+v", d.ErrorID)
	fmt.Fprintf(&builder, ", ErrorMsg=%+v", d.ErrorMsg)

	builder.WriteByte('}')

	return builder.String()
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
	CustomerName types.TThostFtdcLongIndividualNameType
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

func (d CThostFtdcReserveOpenAccountField) Type() string { return "CThostFtdcReserveOpenAccountField" }

func (d CThostFtdcReserveOpenAccountField) String() string {
	var builder strings.Builder
	builder.Grow(770)

	builder.WriteString("CThostFtdcReserveOpenAccountField{")
	fmt.Fprintf(&builder, "TradeCode=%+v", d.TradeCode)
	fmt.Fprintf(&builder, ", BankID=%+v", d.BankID)
	fmt.Fprintf(&builder, ", BankBranchID=%+v", d.BankBranchID)
	fmt.Fprintf(&builder, ", BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", BrokerBranchID=%+v", d.BrokerBranchID)
	fmt.Fprintf(&builder, ", TradeDate=%+v", d.TradeDate)
	fmt.Fprintf(&builder, ", TradeTime=%+v", d.TradeTime)
	fmt.Fprintf(&builder, ", BankSerial=%+v", d.BankSerial)
	fmt.Fprintf(&builder, ", TradingDay=%+v", d.TradingDay)
	fmt.Fprintf(&builder, ", PlateSerial=%+v", d.PlateSerial)
	fmt.Fprintf(&builder, ", LastFragment=%+v", d.LastFragment)
	fmt.Fprintf(&builder, ", SessionID=%+v", d.SessionID)
	fmt.Fprintf(&builder, ", CustomerName=%+v", d.CustomerName)
	fmt.Fprintf(&builder, ", IdCardType=%+v", d.IdCardType)
	fmt.Fprintf(&builder, ", IdentifiedCardNo=%+v", d.IdentifiedCardNo)
	fmt.Fprintf(&builder, ", Gender=%+v", d.Gender)
	fmt.Fprintf(&builder, ", CountryCode=%+v", d.CountryCode)
	fmt.Fprintf(&builder, ", CustType=%+v", d.CustType)
	fmt.Fprintf(&builder, ", Address=%+v", d.Address)
	fmt.Fprintf(&builder, ", ZipCode=%+v", d.ZipCode)
	fmt.Fprintf(&builder, ", Telephone=%+v", d.Telephone)
	fmt.Fprintf(&builder, ", MobilePhone=%+v", d.MobilePhone)
	fmt.Fprintf(&builder, ", Fax=%+v", d.Fax)
	fmt.Fprintf(&builder, ", EMail=%+v", d.EMail)
	fmt.Fprintf(&builder, ", MoneyAccountStatus=%+v", d.MoneyAccountStatus)
	fmt.Fprintf(&builder, ", BankAccount=%+v", d.BankAccount)
	fmt.Fprintf(&builder, ", BankPassWord=%+v", d.BankPassWord)
	fmt.Fprintf(&builder, ", InstallID=%+v", d.InstallID)
	fmt.Fprintf(&builder, ", VerifyCertNoFlag=%+v", d.VerifyCertNoFlag)
	fmt.Fprintf(&builder, ", CurrencyID=%+v", d.CurrencyID)
	fmt.Fprintf(&builder, ", Digest=%+v", d.Digest)
	fmt.Fprintf(&builder, ", BankAccType=%+v", d.BankAccType)
	fmt.Fprintf(&builder, ", BrokerIDByBank=%+v", d.BrokerIDByBank)
	fmt.Fprintf(&builder, ", TID=%+v", d.TID)
	fmt.Fprintf(&builder, ", ReserveOpenAccStas=%+v", d.ReserveOpenAccStas)
	fmt.Fprintf(&builder, ", ErrorID=%+v", d.ErrorID)
	fmt.Fprintf(&builder, ", ErrorMsg=%+v", d.ErrorMsg)

	builder.WriteByte('}')

	return builder.String()
}

// 银行账户属性
type CThostFtdcAccountPropertyField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 投资者帐号
	AccountID types.TThostFtdcAccountIDType
	// 银行统一标识类型
	BankID types.TThostFtdcBankIDType
	// 银行账户
	BankAccount types.TThostFtdcBankAccountType
	// 银行账户的开户人名称
	OpenName types.TThostFtdcInvestorFullNameType
	// 银行账户的开户行
	OpenBank types.TThostFtdcOpenBankType
	// 是否活跃
	IsActive types.TThostFtdcBoolType
	// 账户来源
	AccountSourceType types.TThostFtdcAccountSourceTypeType
	// 开户日期
	OpenDate types.TThostFtdcDateType
	// 注销日期
	CancelDate types.TThostFtdcDateType
	// 录入员代码
	OperatorID types.TThostFtdcOperatorIDType
	// 录入日期
	OperateDate types.TThostFtdcDateType
	// 录入时间
	OperateTime types.TThostFtdcTimeType
	// 币种代码
	CurrencyID types.TThostFtdcCurrencyIDType
}

func (d CThostFtdcAccountPropertyField) Type() string { return "CThostFtdcAccountPropertyField" }

func (d CThostFtdcAccountPropertyField) String() string {
	var builder strings.Builder
	builder.Grow(305)

	builder.WriteString("CThostFtdcAccountPropertyField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", AccountID=%+v", d.AccountID)
	fmt.Fprintf(&builder, ", BankID=%+v", d.BankID)
	fmt.Fprintf(&builder, ", BankAccount=%+v", d.BankAccount)
	fmt.Fprintf(&builder, ", OpenName=%+v", d.OpenName)
	fmt.Fprintf(&builder, ", OpenBank=%+v", d.OpenBank)
	fmt.Fprintf(&builder, ", IsActive=%+v", d.IsActive)
	fmt.Fprintf(&builder, ", AccountSourceType=%+v", d.AccountSourceType)
	fmt.Fprintf(&builder, ", OpenDate=%+v", d.OpenDate)
	fmt.Fprintf(&builder, ", CancelDate=%+v", d.CancelDate)
	fmt.Fprintf(&builder, ", OperatorID=%+v", d.OperatorID)
	fmt.Fprintf(&builder, ", OperateDate=%+v", d.OperateDate)
	fmt.Fprintf(&builder, ", OperateTime=%+v", d.OperateTime)
	fmt.Fprintf(&builder, ", CurrencyID=%+v", d.CurrencyID)

	builder.WriteByte('}')

	return builder.String()
}

// 查询当前交易中心
type CThostFtdcQryCurrDRIdentityField struct {
	// 交易中心代码
	DRIdentityID types.TThostFtdcDRIdentityIDType
}

func (d CThostFtdcQryCurrDRIdentityField) Type() string { return "CThostFtdcQryCurrDRIdentityField" }

func (d CThostFtdcQryCurrDRIdentityField) String() string {
	var builder strings.Builder
	builder.Grow(54)

	builder.WriteString("CThostFtdcQryCurrDRIdentityField{")
	fmt.Fprintf(&builder, "DRIdentityID=%+v", d.DRIdentityID)

	builder.WriteByte('}')

	return builder.String()
}

// 当前交易中心
type CThostFtdcCurrDRIdentityField struct {
	// 交易中心代码
	DRIdentityID types.TThostFtdcDRIdentityIDType
}

func (d CThostFtdcCurrDRIdentityField) Type() string { return "CThostFtdcCurrDRIdentityField" }

func (d CThostFtdcCurrDRIdentityField) String() string {
	var builder strings.Builder
	builder.Grow(51)

	builder.WriteString("CThostFtdcCurrDRIdentityField{")
	fmt.Fprintf(&builder, "DRIdentityID=%+v", d.DRIdentityID)

	builder.WriteByte('}')

	return builder.String()
}

// 查询二级代理商资金校验模式
type CThostFtdcQrySecAgentCheckModeField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID types.TThostFtdcInvestorIDType
}

func (d CThostFtdcQrySecAgentCheckModeField) Type() string {
	return "CThostFtdcQrySecAgentCheckModeField"
}

func (d CThostFtdcQrySecAgentCheckModeField) String() string {
	var builder strings.Builder
	builder.Grow(73)

	builder.WriteString("CThostFtdcQrySecAgentCheckModeField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", InvestorID=%+v", d.InvestorID)

	builder.WriteByte('}')

	return builder.String()
}

// 查询二级代理商信息
type CThostFtdcQrySecAgentTradeInfoField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 境外中介机构资金帐号
	BrokerSecAgentID types.TThostFtdcAccountIDType
}

func (d CThostFtdcQrySecAgentTradeInfoField) Type() string {
	return "CThostFtdcQrySecAgentTradeInfoField"
}

func (d CThostFtdcQrySecAgentTradeInfoField) String() string {
	var builder strings.Builder
	builder.Grow(79)

	builder.WriteString("CThostFtdcQrySecAgentTradeInfoField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", BrokerSecAgentID=%+v", d.BrokerSecAgentID)

	builder.WriteByte('}')

	return builder.String()
}

// 用户发出获取安全安全登陆方法请求
type CThostFtdcReqUserAuthMethodField struct {
	// 交易日
	TradingDay types.TThostFtdcDateType
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 用户代码
	UserID types.TThostFtdcUserIDType
}

func (d CThostFtdcReqUserAuthMethodField) Type() string { return "CThostFtdcReqUserAuthMethodField" }

func (d CThostFtdcReqUserAuthMethodField) String() string {
	var builder strings.Builder
	builder.Grow(86)

	builder.WriteString("CThostFtdcReqUserAuthMethodField{")
	fmt.Fprintf(&builder, "TradingDay=%+v", d.TradingDay)
	fmt.Fprintf(&builder, ", BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", UserID=%+v", d.UserID)

	builder.WriteByte('}')

	return builder.String()
}

// 用户发出获取安全安全登陆方法回复
type CThostFtdcRspUserAuthMethodField struct {
	// 当前可以用的认证模式
	UsableAuthMethod types.TThostFtdcCurrentAuthMethodType
}

func (d CThostFtdcRspUserAuthMethodField) Type() string { return "CThostFtdcRspUserAuthMethodField" }

func (d CThostFtdcRspUserAuthMethodField) String() string {
	var builder strings.Builder
	builder.Grow(58)

	builder.WriteString("CThostFtdcRspUserAuthMethodField{")
	fmt.Fprintf(&builder, "UsableAuthMethod=%+v", d.UsableAuthMethod)

	builder.WriteByte('}')

	return builder.String()
}

// 用户发出获取安全安全登陆方法请求
type CThostFtdcReqGenUserCaptchaField struct {
	// 交易日
	TradingDay types.TThostFtdcDateType
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 用户代码
	UserID types.TThostFtdcUserIDType
}

func (d CThostFtdcReqGenUserCaptchaField) Type() string { return "CThostFtdcReqGenUserCaptchaField" }

func (d CThostFtdcReqGenUserCaptchaField) String() string {
	var builder strings.Builder
	builder.Grow(86)

	builder.WriteString("CThostFtdcReqGenUserCaptchaField{")
	fmt.Fprintf(&builder, "TradingDay=%+v", d.TradingDay)
	fmt.Fprintf(&builder, ", BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", UserID=%+v", d.UserID)

	builder.WriteByte('}')

	return builder.String()
}

// 生成的图片验证码信息
type CThostFtdcRspGenUserCaptchaField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 用户代码
	UserID types.TThostFtdcUserIDType
	// 图片信息长度
	CaptchaInfoLen types.TThostFtdcCaptchaInfoLenType
	// 图片信息
	CaptchaInfo types.TThostFtdcCaptchaInfoType
}

func (d CThostFtdcRspGenUserCaptchaField) Type() string { return "CThostFtdcRspGenUserCaptchaField" }

func (d CThostFtdcRspGenUserCaptchaField) String() string {
	var builder strings.Builder
	builder.Grow(111)

	builder.WriteString("CThostFtdcRspGenUserCaptchaField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", UserID=%+v", d.UserID)
	fmt.Fprintf(&builder, ", CaptchaInfoLen=%+v", d.CaptchaInfoLen)
	fmt.Fprintf(&builder, ", CaptchaInfo=%+v", d.CaptchaInfo)

	builder.WriteByte('}')

	return builder.String()
}

// 用户发出获取安全安全登陆方法请求
type CThostFtdcReqGenUserTextField struct {
	// 交易日
	TradingDay types.TThostFtdcDateType
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 用户代码
	UserID types.TThostFtdcUserIDType
}

func (d CThostFtdcReqGenUserTextField) Type() string { return "CThostFtdcReqGenUserTextField" }

func (d CThostFtdcReqGenUserTextField) String() string {
	var builder strings.Builder
	builder.Grow(83)

	builder.WriteString("CThostFtdcReqGenUserTextField{")
	fmt.Fprintf(&builder, "TradingDay=%+v", d.TradingDay)
	fmt.Fprintf(&builder, ", BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", UserID=%+v", d.UserID)

	builder.WriteByte('}')

	return builder.String()
}

// 短信验证码生成的回复
type CThostFtdcRspGenUserTextField struct {
	// 短信验证码序号
	UserTextSeq types.TThostFtdcUserTextSeqType
}

func (d CThostFtdcRspGenUserTextField) Type() string { return "CThostFtdcRspGenUserTextField" }

func (d CThostFtdcRspGenUserTextField) String() string {
	var builder strings.Builder
	builder.Grow(50)

	builder.WriteString("CThostFtdcRspGenUserTextField{")
	fmt.Fprintf(&builder, "UserTextSeq=%+v", d.UserTextSeq)

	builder.WriteByte('}')

	return builder.String()
}

// 用户发出带图形验证码的登录请求请求
type CThostFtdcReqUserLoginWithCaptchaField struct {
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
	// 保留的无效字段
	reserve1 types.TThostFtdcOldIPAddressType
	// 登录备注
	LoginRemark types.TThostFtdcLoginRemarkType
	// 图形验证码的文字内容
	Captcha types.TThostFtdcPasswordType
	// 终端IP端口
	ClientIPPort types.TThostFtdcIPPortType
	// 终端IP地址
	ClientIPAddress types.TThostFtdcIPAddressType
}

func (d CThostFtdcReqUserLoginWithCaptchaField) Type() string {
	return "CThostFtdcReqUserLoginWithCaptchaField"
}

func (d CThostFtdcReqUserLoginWithCaptchaField) String() string {
	var builder strings.Builder
	builder.Grow(310)

	builder.WriteString("CThostFtdcReqUserLoginWithCaptchaField{")
	fmt.Fprintf(&builder, "TradingDay=%+v", d.TradingDay)
	fmt.Fprintf(&builder, ", BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", UserID=%+v", d.UserID)
	fmt.Fprintf(&builder, ", Password=%+v", d.Password)
	fmt.Fprintf(&builder, ", UserProductInfo=%+v", d.UserProductInfo)
	fmt.Fprintf(&builder, ", InterfaceProductInfo=%+v", d.InterfaceProductInfo)
	fmt.Fprintf(&builder, ", ProtocolInfo=%+v", d.ProtocolInfo)
	fmt.Fprintf(&builder, ", MacAddress=%+v", d.MacAddress)
	fmt.Fprintf(&builder, ", reserve1=%+v", d.reserve1)
	fmt.Fprintf(&builder, ", LoginRemark=%+v", d.LoginRemark)
	fmt.Fprintf(&builder, ", Captcha=%+v", d.Captcha)
	fmt.Fprintf(&builder, ", ClientIPPort=%+v", d.ClientIPPort)
	fmt.Fprintf(&builder, ", ClientIPAddress=%+v", d.ClientIPAddress)

	builder.WriteByte('}')

	return builder.String()
}

// 用户发出带短信验证码的登录请求请求
type CThostFtdcReqUserLoginWithTextField struct {
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
	// 保留的无效字段
	reserve1 types.TThostFtdcOldIPAddressType
	// 登录备注
	LoginRemark types.TThostFtdcLoginRemarkType
	// 短信验证码文字内容
	Text types.TThostFtdcPasswordType
	// 终端IP端口
	ClientIPPort types.TThostFtdcIPPortType
	// 终端IP地址
	ClientIPAddress types.TThostFtdcIPAddressType
}

func (d CThostFtdcReqUserLoginWithTextField) Type() string {
	return "CThostFtdcReqUserLoginWithTextField"
}

func (d CThostFtdcReqUserLoginWithTextField) String() string {
	var builder strings.Builder
	builder.Grow(304)

	builder.WriteString("CThostFtdcReqUserLoginWithTextField{")
	fmt.Fprintf(&builder, "TradingDay=%+v", d.TradingDay)
	fmt.Fprintf(&builder, ", BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", UserID=%+v", d.UserID)
	fmt.Fprintf(&builder, ", Password=%+v", d.Password)
	fmt.Fprintf(&builder, ", UserProductInfo=%+v", d.UserProductInfo)
	fmt.Fprintf(&builder, ", InterfaceProductInfo=%+v", d.InterfaceProductInfo)
	fmt.Fprintf(&builder, ", ProtocolInfo=%+v", d.ProtocolInfo)
	fmt.Fprintf(&builder, ", MacAddress=%+v", d.MacAddress)
	fmt.Fprintf(&builder, ", reserve1=%+v", d.reserve1)
	fmt.Fprintf(&builder, ", LoginRemark=%+v", d.LoginRemark)
	fmt.Fprintf(&builder, ", Text=%+v", d.Text)
	fmt.Fprintf(&builder, ", ClientIPPort=%+v", d.ClientIPPort)
	fmt.Fprintf(&builder, ", ClientIPAddress=%+v", d.ClientIPAddress)

	builder.WriteByte('}')

	return builder.String()
}

// 用户发出带动态验证码的登录请求请求
type CThostFtdcReqUserLoginWithOTPField struct {
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
	// 保留的无效字段
	reserve1 types.TThostFtdcOldIPAddressType
	// 登录备注
	LoginRemark types.TThostFtdcLoginRemarkType
	// OTP密码
	OTPPassword types.TThostFtdcPasswordType
	// 终端IP端口
	ClientIPPort types.TThostFtdcIPPortType
	// 终端IP地址
	ClientIPAddress types.TThostFtdcIPAddressType
}

func (d CThostFtdcReqUserLoginWithOTPField) Type() string {
	return "CThostFtdcReqUserLoginWithOTPField"
}

func (d CThostFtdcReqUserLoginWithOTPField) String() string {
	var builder strings.Builder
	builder.Grow(310)

	builder.WriteString("CThostFtdcReqUserLoginWithOTPField{")
	fmt.Fprintf(&builder, "TradingDay=%+v", d.TradingDay)
	fmt.Fprintf(&builder, ", BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", UserID=%+v", d.UserID)
	fmt.Fprintf(&builder, ", Password=%+v", d.Password)
	fmt.Fprintf(&builder, ", UserProductInfo=%+v", d.UserProductInfo)
	fmt.Fprintf(&builder, ", InterfaceProductInfo=%+v", d.InterfaceProductInfo)
	fmt.Fprintf(&builder, ", ProtocolInfo=%+v", d.ProtocolInfo)
	fmt.Fprintf(&builder, ", MacAddress=%+v", d.MacAddress)
	fmt.Fprintf(&builder, ", reserve1=%+v", d.reserve1)
	fmt.Fprintf(&builder, ", LoginRemark=%+v", d.LoginRemark)
	fmt.Fprintf(&builder, ", OTPPassword=%+v", d.OTPPassword)
	fmt.Fprintf(&builder, ", ClientIPPort=%+v", d.ClientIPPort)
	fmt.Fprintf(&builder, ", ClientIPAddress=%+v", d.ClientIPAddress)

	builder.WriteByte('}')

	return builder.String()
}

// api握手请求
type CThostFtdcReqApiHandshakeField struct {
	// api与front通信密钥版本号
	CryptoKeyVersion types.TThostFtdcCryptoKeyVersionType
}

func (d CThostFtdcReqApiHandshakeField) Type() string { return "CThostFtdcReqApiHandshakeField" }

func (d CThostFtdcReqApiHandshakeField) String() string {
	var builder strings.Builder
	builder.Grow(56)

	builder.WriteString("CThostFtdcReqApiHandshakeField{")
	fmt.Fprintf(&builder, "CryptoKeyVersion=%+v", d.CryptoKeyVersion)

	builder.WriteByte('}')

	return builder.String()
}

// front发给api的握手回复
type CThostFtdcRspApiHandshakeField struct {
	// 握手回复数据长度
	FrontHandshakeDataLen types.TThostFtdcHandshakeDataLenType
	// 握手回复数据
	FrontHandshakeData types.TThostFtdcHandshakeDataType
	// API认证是否开启
	IsApiAuthEnabled types.TThostFtdcBoolType
}

func (d CThostFtdcRspApiHandshakeField) Type() string { return "CThostFtdcRspApiHandshakeField" }

func (d CThostFtdcRspApiHandshakeField) String() string {
	var builder strings.Builder
	builder.Grow(115)

	builder.WriteString("CThostFtdcRspApiHandshakeField{")
	fmt.Fprintf(&builder, "FrontHandshakeDataLen=%+v", d.FrontHandshakeDataLen)
	fmt.Fprintf(&builder, ", FrontHandshakeData=%+v", d.FrontHandshakeData)
	fmt.Fprintf(&builder, ", IsApiAuthEnabled=%+v", d.IsApiAuthEnabled)

	builder.WriteByte('}')

	return builder.String()
}

// api给front的验证key的请求
type CThostFtdcReqVerifyApiKeyField struct {
	// 握手回复数据长度
	ApiHandshakeDataLen types.TThostFtdcHandshakeDataLenType
	// 握手回复数据
	ApiHandshakeData types.TThostFtdcHandshakeDataType
}

func (d CThostFtdcReqVerifyApiKeyField) Type() string { return "CThostFtdcReqVerifyApiKeyField" }

func (d CThostFtdcReqVerifyApiKeyField) String() string {
	var builder strings.Builder
	builder.Grow(85)

	builder.WriteString("CThostFtdcReqVerifyApiKeyField{")
	fmt.Fprintf(&builder, "ApiHandshakeDataLen=%+v", d.ApiHandshakeDataLen)
	fmt.Fprintf(&builder, ", ApiHandshakeData=%+v", d.ApiHandshakeData)

	builder.WriteByte('}')

	return builder.String()
}

// 操作员组织架构关系
type CThostFtdcDepartmentUserField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 用户代码
	UserID types.TThostFtdcUserIDType
	// 投资者范围
	InvestorRange types.TThostFtdcDepartmentRangeType
	// 投资者代码
	InvestorID types.TThostFtdcInvestorIDType
}

func (d CThostFtdcDepartmentUserField) Type() string { return "CThostFtdcDepartmentUserField" }

func (d CThostFtdcDepartmentUserField) String() string {
	var builder strings.Builder
	builder.Grow(106)

	builder.WriteString("CThostFtdcDepartmentUserField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", UserID=%+v", d.UserID)
	fmt.Fprintf(&builder, ", InvestorRange=%+v", d.InvestorRange)
	fmt.Fprintf(&builder, ", InvestorID=%+v", d.InvestorID)

	builder.WriteByte('}')

	return builder.String()
}

// 查询频率，每秒查询比数
type CThostFtdcQueryFreqField struct {
	// 查询频率
	QueryFreq types.TThostFtdcQueryFreqType
	// FTD频率
	FTDPkgFreq types.TThostFtdcQueryFreqType
}

func (d CThostFtdcQueryFreqField) Type() string { return "CThostFtdcQueryFreqField" }

func (d CThostFtdcQueryFreqField) String() string {
	var builder strings.Builder
	builder.Grow(63)

	builder.WriteString("CThostFtdcQueryFreqField{")
	fmt.Fprintf(&builder, "QueryFreq=%+v", d.QueryFreq)
	fmt.Fprintf(&builder, ", FTDPkgFreq=%+v", d.FTDPkgFreq)

	builder.WriteByte('}')

	return builder.String()
}

// 禁止认证IP
type CThostFtdcAuthForbiddenIPField struct {
	// IP地址
	IPAddress types.TThostFtdcIPAddressType
}

func (d CThostFtdcAuthForbiddenIPField) Type() string { return "CThostFtdcAuthForbiddenIPField" }

func (d CThostFtdcAuthForbiddenIPField) String() string {
	var builder strings.Builder
	builder.Grow(49)

	builder.WriteString("CThostFtdcAuthForbiddenIPField{")
	fmt.Fprintf(&builder, "IPAddress=%+v", d.IPAddress)

	builder.WriteByte('}')

	return builder.String()
}

// 查询禁止认证IP
type CThostFtdcQryAuthForbiddenIPField struct {
	// IP地址
	IPAddress types.TThostFtdcIPAddressType
}

func (d CThostFtdcQryAuthForbiddenIPField) Type() string { return "CThostFtdcQryAuthForbiddenIPField" }

func (d CThostFtdcQryAuthForbiddenIPField) String() string {
	var builder strings.Builder
	builder.Grow(52)

	builder.WriteString("CThostFtdcQryAuthForbiddenIPField{")
	fmt.Fprintf(&builder, "IPAddress=%+v", d.IPAddress)

	builder.WriteByte('}')

	return builder.String()
}

// 换汇可提冻结
type CThostFtdcSyncDelaySwapFrozenField struct {
	// 换汇流水号
	DelaySwapSeqNo types.TThostFtdcDepositSeqNoType
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	// 源币种
	FromCurrencyID types.TThostFtdcCurrencyIDType
	// 源剩余换汇额度(可提冻结)
	FromRemainSwap types.TThostFtdcMoneyType
	// 是否手工换汇
	IsManualSwap types.TThostFtdcBoolType
}

func (d CThostFtdcSyncDelaySwapFrozenField) Type() string {
	return "CThostFtdcSyncDelaySwapFrozenField"
}

func (d CThostFtdcSyncDelaySwapFrozenField) String() string {
	var builder strings.Builder
	builder.Grow(166)

	builder.WriteString("CThostFtdcSyncDelaySwapFrozenField{")
	fmt.Fprintf(&builder, "DelaySwapSeqNo=%+v", d.DelaySwapSeqNo)
	fmt.Fprintf(&builder, ", BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", InvestorID=%+v", d.InvestorID)
	fmt.Fprintf(&builder, ", FromCurrencyID=%+v", d.FromCurrencyID)
	fmt.Fprintf(&builder, ", FromRemainSwap=%+v", d.FromRemainSwap)
	fmt.Fprintf(&builder, ", IsManualSwap=%+v", d.IsManualSwap)

	builder.WriteByte('}')

	return builder.String()
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
	// 保留的无效字段
	reserve1 types.TThostFtdcOldIPAddressType
	// 终端IP端口
	ClientIPPort types.TThostFtdcIPPortType
	// 登录成功时间
	ClientLoginTime types.TThostFtdcTimeType
	// App代码
	ClientAppID types.TThostFtdcAppIDType
	// 用户公网IP
	ClientPublicIP types.TThostFtdcIPAddressType
	// 客户登录备注2
	ClientLoginRemark types.TThostFtdcClientLoginRemarkType
	// 客户终端的MAC等标识
	MAC types.TThostFtdcDeviceTagType
}

func (d CThostFtdcUserSystemInfoField) Type() string { return "CThostFtdcUserSystemInfoField" }

func (d CThostFtdcUserSystemInfoField) String() string {
	var builder strings.Builder
	builder.Grow(268)

	builder.WriteString("CThostFtdcUserSystemInfoField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", UserID=%+v", d.UserID)
	fmt.Fprintf(&builder, ", ClientSystemInfoLen=%+v", d.ClientSystemInfoLen)
	fmt.Fprintf(&builder, ", ClientSystemInfo=%+v", d.ClientSystemInfo)
	fmt.Fprintf(&builder, ", reserve1=%+v", d.reserve1)
	fmt.Fprintf(&builder, ", ClientIPPort=%+v", d.ClientIPPort)
	fmt.Fprintf(&builder, ", ClientLoginTime=%+v", d.ClientLoginTime)
	fmt.Fprintf(&builder, ", ClientAppID=%+v", d.ClientAppID)
	fmt.Fprintf(&builder, ", ClientPublicIP=%+v", d.ClientPublicIP)
	fmt.Fprintf(&builder, ", ClientLoginRemark=%+v", d.ClientLoginRemark)
	fmt.Fprintf(&builder, ", MAC=%+v", d.MAC)

	builder.WriteByte('}')

	return builder.String()
}

// 终端用户绑定信息
type CThostFtdcAuthUserIDField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// App代码
	AppID types.TThostFtdcAppIDType
	// 用户代码
	UserID types.TThostFtdcUserIDType
	// 校验类型
	AuthType types.TThostFtdcAuthTypeType
}

func (d CThostFtdcAuthUserIDField) Type() string { return "CThostFtdcAuthUserIDField" }

func (d CThostFtdcAuthUserIDField) String() string {
	var builder strings.Builder
	builder.Grow(92)

	builder.WriteString("CThostFtdcAuthUserIDField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", AppID=%+v", d.AppID)
	fmt.Fprintf(&builder, ", UserID=%+v", d.UserID)
	fmt.Fprintf(&builder, ", AuthType=%+v", d.AuthType)

	builder.WriteByte('}')

	return builder.String()
}

// 用户IP绑定信息
type CThostFtdcAuthIPField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// App代码
	AppID types.TThostFtdcAppIDType
	// 用户代码
	IPAddress types.TThostFtdcIPAddressType
}

func (d CThostFtdcAuthIPField) Type() string { return "CThostFtdcAuthIPField" }

func (d CThostFtdcAuthIPField) String() string {
	var builder strings.Builder
	builder.Grow(73)

	builder.WriteString("CThostFtdcAuthIPField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", AppID=%+v", d.AppID)
	fmt.Fprintf(&builder, ", IPAddress=%+v", d.IPAddress)

	builder.WriteByte('}')

	return builder.String()
}

// 查询分类合约
type CThostFtdcQryClassifiedInstrumentField struct {
	// 合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
	// 交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	// 合约在交易所的代码
	ExchangeInstID types.TThostFtdcExchangeInstIDType
	// 产品代码
	ProductID types.TThostFtdcInstrumentIDType
	// 合约交易状态
	TradingType types.TThostFtdcTradingTypeType
	// 合约分类类型
	ClassType types.TThostFtdcClassTypeType
}

func (d CThostFtdcQryClassifiedInstrumentField) Type() string {
	return "CThostFtdcQryClassifiedInstrumentField"
}

func (d CThostFtdcQryClassifiedInstrumentField) String() string {
	var builder strings.Builder
	builder.Grow(163)

	builder.WriteString("CThostFtdcQryClassifiedInstrumentField{")
	fmt.Fprintf(&builder, "InstrumentID=%+v", d.InstrumentID)
	fmt.Fprintf(&builder, ", ExchangeID=%+v", d.ExchangeID)
	fmt.Fprintf(&builder, ", ExchangeInstID=%+v", d.ExchangeInstID)
	fmt.Fprintf(&builder, ", ProductID=%+v", d.ProductID)
	fmt.Fprintf(&builder, ", TradingType=%+v", d.TradingType)
	fmt.Fprintf(&builder, ", ClassType=%+v", d.ClassType)

	builder.WriteByte('}')

	return builder.String()
}

// 查询组合优惠比例
type CThostFtdcQryCombPromotionParamField struct {
	// 交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	// 合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
}

func (d CThostFtdcQryCombPromotionParamField) Type() string {
	return "CThostFtdcQryCombPromotionParamField"
}

func (d CThostFtdcQryCombPromotionParamField) String() string {
	var builder strings.Builder
	builder.Grow(78)

	builder.WriteString("CThostFtdcQryCombPromotionParamField{")
	fmt.Fprintf(&builder, "ExchangeID=%+v", d.ExchangeID)
	fmt.Fprintf(&builder, ", InstrumentID=%+v", d.InstrumentID)

	builder.WriteByte('}')

	return builder.String()
}

// 组合优惠比例
type CThostFtdcCombPromotionParamField struct {
	// 交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	// 合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
	// 投机套保标志
	CombHedgeFlag types.TThostFtdcCombHedgeFlagType
	// 期权组合保证金比例
	Xparameter types.TThostFtdcDiscountRatioType
}

func (d CThostFtdcCombPromotionParamField) Type() string { return "CThostFtdcCombPromotionParamField" }

func (d CThostFtdcCombPromotionParamField) String() string {
	var builder strings.Builder
	builder.Grow(118)

	builder.WriteString("CThostFtdcCombPromotionParamField{")
	fmt.Fprintf(&builder, "ExchangeID=%+v", d.ExchangeID)
	fmt.Fprintf(&builder, ", InstrumentID=%+v", d.InstrumentID)
	fmt.Fprintf(&builder, ", CombHedgeFlag=%+v", d.CombHedgeFlag)
	fmt.Fprintf(&builder, ", Xparameter=%+v", d.Xparameter)

	builder.WriteByte('}')

	return builder.String()
}

// 国密用户登录请求
type CThostFtdcReqUserLoginSMField struct {
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
	// 保留的无效字段
	reserve1 types.TThostFtdcOldIPAddressType
	// 登录备注
	LoginRemark types.TThostFtdcLoginRemarkType
	// 终端IP端口
	ClientIPPort types.TThostFtdcIPPortType
	// 终端IP地址
	ClientIPAddress types.TThostFtdcIPAddressType
	// 短信验证码
	SMSCode types.TThostFtdcSMSCodeType
	// 经纪公司名称
	BrokerName types.TThostFtdcBrokerNameType
	// 认证码
	AuthCode types.TThostFtdcAuthCodeType
	// App代码
	AppID types.TThostFtdcAppIDType
	// PIN码
	PIN types.TThostFtdcPasswordType
}

func (d CThostFtdcReqUserLoginSMField) Type() string { return "CThostFtdcReqUserLoginSMField" }

func (d CThostFtdcReqUserLoginSMField) String() string {
	var builder strings.Builder
	builder.Grow(392)

	builder.WriteString("CThostFtdcReqUserLoginSMField{")
	fmt.Fprintf(&builder, "TradingDay=%+v", d.TradingDay)
	fmt.Fprintf(&builder, ", BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", UserID=%+v", d.UserID)
	fmt.Fprintf(&builder, ", Password=%+v", d.Password)
	fmt.Fprintf(&builder, ", UserProductInfo=%+v", d.UserProductInfo)
	fmt.Fprintf(&builder, ", InterfaceProductInfo=%+v", d.InterfaceProductInfo)
	fmt.Fprintf(&builder, ", ProtocolInfo=%+v", d.ProtocolInfo)
	fmt.Fprintf(&builder, ", MacAddress=%+v", d.MacAddress)
	fmt.Fprintf(&builder, ", OneTimePassword=%+v", d.OneTimePassword)
	fmt.Fprintf(&builder, ", reserve1=%+v", d.reserve1)
	fmt.Fprintf(&builder, ", LoginRemark=%+v", d.LoginRemark)
	fmt.Fprintf(&builder, ", ClientIPPort=%+v", d.ClientIPPort)
	fmt.Fprintf(&builder, ", ClientIPAddress=%+v", d.ClientIPAddress)
	fmt.Fprintf(&builder, ", SMSCode=%+v", d.SMSCode)
	fmt.Fprintf(&builder, ", BrokerName=%+v", d.BrokerName)
	fmt.Fprintf(&builder, ", AuthCode=%+v", d.AuthCode)
	fmt.Fprintf(&builder, ", AppID=%+v", d.AppID)
	fmt.Fprintf(&builder, ", PIN=%+v", d.PIN)

	builder.WriteByte('}')

	return builder.String()
}

// 投资者风险结算持仓查询
type CThostFtdcQryRiskSettleInvstPositionField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	// 合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
}

func (d CThostFtdcQryRiskSettleInvstPositionField) Type() string {
	return "CThostFtdcQryRiskSettleInvstPositionField"
}

func (d CThostFtdcQryRiskSettleInvstPositionField) String() string {
	var builder strings.Builder
	builder.Grow(101)

	builder.WriteString("CThostFtdcQryRiskSettleInvstPositionField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", InvestorID=%+v", d.InvestorID)
	fmt.Fprintf(&builder, ", InstrumentID=%+v", d.InstrumentID)

	builder.WriteByte('}')

	return builder.String()
}

// 风险结算产品查询
type CThostFtdcQryRiskSettleProductStatusField struct {
	// 产品代码
	ProductID types.TThostFtdcInstrumentIDType
}

func (d CThostFtdcQryRiskSettleProductStatusField) Type() string {
	return "CThostFtdcQryRiskSettleProductStatusField"
}

func (d CThostFtdcQryRiskSettleProductStatusField) String() string {
	var builder strings.Builder
	builder.Grow(60)

	builder.WriteString("CThostFtdcQryRiskSettleProductStatusField{")
	fmt.Fprintf(&builder, "ProductID=%+v", d.ProductID)

	builder.WriteByte('}')

	return builder.String()
}

// 投资者风险结算持仓
type CThostFtdcRiskSettleInvstPositionField struct {
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
	// 交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	// 执行冻结的昨仓
	YdStrikeFrozen types.TThostFtdcVolumeType
	// 投资单元代码
	InvestUnitID types.TThostFtdcInvestUnitIDType
	// 持仓成本差值
	PositionCostOffset types.TThostFtdcMoneyType
	// tas持仓手数
	TasPosition types.TThostFtdcVolumeType
	// tas持仓成本
	TasPositionCost types.TThostFtdcMoneyType
}

func (d CThostFtdcRiskSettleInvstPositionField) Type() string {
	return "CThostFtdcRiskSettleInvstPositionField"
}

func (d CThostFtdcRiskSettleInvstPositionField) String() string {
	var builder strings.Builder
	builder.Grow(1139)

	builder.WriteString("CThostFtdcRiskSettleInvstPositionField{")
	fmt.Fprintf(&builder, "InstrumentID=%+v", d.InstrumentID)
	fmt.Fprintf(&builder, ", BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", InvestorID=%+v", d.InvestorID)
	fmt.Fprintf(&builder, ", PosiDirection=%+v", d.PosiDirection)
	fmt.Fprintf(&builder, ", HedgeFlag=%+v", d.HedgeFlag)
	fmt.Fprintf(&builder, ", PositionDate=%+v", d.PositionDate)
	fmt.Fprintf(&builder, ", YdPosition=%+v", d.YdPosition)
	fmt.Fprintf(&builder, ", Position=%+v", d.Position)
	fmt.Fprintf(&builder, ", LongFrozen=%+v", d.LongFrozen)
	fmt.Fprintf(&builder, ", ShortFrozen=%+v", d.ShortFrozen)
	fmt.Fprintf(&builder, ", LongFrozenAmount=%+v", d.LongFrozenAmount)
	fmt.Fprintf(&builder, ", ShortFrozenAmount=%+v", d.ShortFrozenAmount)
	fmt.Fprintf(&builder, ", OpenVolume=%+v", d.OpenVolume)
	fmt.Fprintf(&builder, ", CloseVolume=%+v", d.CloseVolume)
	fmt.Fprintf(&builder, ", OpenAmount=%+v", d.OpenAmount)
	fmt.Fprintf(&builder, ", CloseAmount=%+v", d.CloseAmount)
	fmt.Fprintf(&builder, ", PositionCost=%+v", d.PositionCost)
	fmt.Fprintf(&builder, ", PreMargin=%+v", d.PreMargin)
	fmt.Fprintf(&builder, ", UseMargin=%+v", d.UseMargin)
	fmt.Fprintf(&builder, ", FrozenMargin=%+v", d.FrozenMargin)
	fmt.Fprintf(&builder, ", FrozenCash=%+v", d.FrozenCash)
	fmt.Fprintf(&builder, ", FrozenCommission=%+v", d.FrozenCommission)
	fmt.Fprintf(&builder, ", CashIn=%+v", d.CashIn)
	fmt.Fprintf(&builder, ", Commission=%+v", d.Commission)
	fmt.Fprintf(&builder, ", CloseProfit=%+v", d.CloseProfit)
	fmt.Fprintf(&builder, ", PositionProfit=%+v", d.PositionProfit)
	fmt.Fprintf(&builder, ", PreSettlementPrice=%+v", d.PreSettlementPrice)
	fmt.Fprintf(&builder, ", SettlementPrice=%+v", d.SettlementPrice)
	fmt.Fprintf(&builder, ", TradingDay=%+v", d.TradingDay)
	fmt.Fprintf(&builder, ", SettlementID=%+v", d.SettlementID)
	fmt.Fprintf(&builder, ", OpenCost=%+v", d.OpenCost)
	fmt.Fprintf(&builder, ", ExchangeMargin=%+v", d.ExchangeMargin)
	fmt.Fprintf(&builder, ", CombPosition=%+v", d.CombPosition)
	fmt.Fprintf(&builder, ", CombLongFrozen=%+v", d.CombLongFrozen)
	fmt.Fprintf(&builder, ", CombShortFrozen=%+v", d.CombShortFrozen)
	fmt.Fprintf(&builder, ", CloseProfitByDate=%+v", d.CloseProfitByDate)
	fmt.Fprintf(&builder, ", CloseProfitByTrade=%+v", d.CloseProfitByTrade)
	fmt.Fprintf(&builder, ", TodayPosition=%+v", d.TodayPosition)
	fmt.Fprintf(&builder, ", MarginRateByMoney=%+v", d.MarginRateByMoney)
	fmt.Fprintf(&builder, ", MarginRateByVolume=%+v", d.MarginRateByVolume)
	fmt.Fprintf(&builder, ", StrikeFrozen=%+v", d.StrikeFrozen)
	fmt.Fprintf(&builder, ", StrikeFrozenAmount=%+v", d.StrikeFrozenAmount)
	fmt.Fprintf(&builder, ", AbandonFrozen=%+v", d.AbandonFrozen)
	fmt.Fprintf(&builder, ", ExchangeID=%+v", d.ExchangeID)
	fmt.Fprintf(&builder, ", YdStrikeFrozen=%+v", d.YdStrikeFrozen)
	fmt.Fprintf(&builder, ", InvestUnitID=%+v", d.InvestUnitID)
	fmt.Fprintf(&builder, ", PositionCostOffset=%+v", d.PositionCostOffset)
	fmt.Fprintf(&builder, ", TasPosition=%+v", d.TasPosition)
	fmt.Fprintf(&builder, ", TasPositionCost=%+v", d.TasPositionCost)

	builder.WriteByte('}')

	return builder.String()
}

// 风险品种
type CThostFtdcRiskSettleProductStatusField struct {
	// 交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	// 产品编号
	ProductID types.TThostFtdcInstrumentIDType
	// 产品结算状态
	ProductStatus types.TThostFtdcProductStatusType
}

func (d CThostFtdcRiskSettleProductStatusField) Type() string {
	return "CThostFtdcRiskSettleProductStatusField"
}

func (d CThostFtdcRiskSettleProductStatusField) String() string {
	var builder strings.Builder
	builder.Grow(100)

	builder.WriteString("CThostFtdcRiskSettleProductStatusField{")
	fmt.Fprintf(&builder, "ExchangeID=%+v", d.ExchangeID)
	fmt.Fprintf(&builder, ", ProductID=%+v", d.ProductID)
	fmt.Fprintf(&builder, ", ProductStatus=%+v", d.ProductStatus)

	builder.WriteByte('}')

	return builder.String()
}

// 风险结算追平信息
type CThostFtdcSyncDeltaInfoField struct {
	// 追平序号
	SyncDeltaSequenceNo types.TThostFtdcSequenceNoType
	// 追平状态
	SyncDeltaStatus types.TThostFtdcSyncDeltaStatusType
	// 追平描述
	SyncDescription types.TThostFtdcSyncDescriptionType
	// 是否只有资金追平
	IsOnlyTrdDelta types.TThostFtdcBoolType
}

func (d CThostFtdcSyncDeltaInfoField) Type() string { return "CThostFtdcSyncDeltaInfoField" }

func (d CThostFtdcSyncDeltaInfoField) String() string {
	var builder strings.Builder
	builder.Grow(131)

	builder.WriteString("CThostFtdcSyncDeltaInfoField{")
	fmt.Fprintf(&builder, "SyncDeltaSequenceNo=%+v", d.SyncDeltaSequenceNo)
	fmt.Fprintf(&builder, ", SyncDeltaStatus=%+v", d.SyncDeltaStatus)
	fmt.Fprintf(&builder, ", SyncDescription=%+v", d.SyncDescription)
	fmt.Fprintf(&builder, ", IsOnlyTrdDelta=%+v", d.IsOnlyTrdDelta)

	builder.WriteByte('}')

	return builder.String()
}

// 风险结算追平产品信息
type CThostFtdcSyncDeltaProductStatusField struct {
	// 追平序号
	SyncDeltaSequenceNo types.TThostFtdcSequenceNoType
	// 交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	// 产品代码
	ProductID types.TThostFtdcInstrumentIDType
	// 是否允许交易
	ProductStatus types.TThostFtdcProductStatusType
}

func (d CThostFtdcSyncDeltaProductStatusField) Type() string {
	return "CThostFtdcSyncDeltaProductStatusField"
}

func (d CThostFtdcSyncDeltaProductStatusField) String() string {
	var builder strings.Builder
	builder.Grow(128)

	builder.WriteString("CThostFtdcSyncDeltaProductStatusField{")
	fmt.Fprintf(&builder, "SyncDeltaSequenceNo=%+v", d.SyncDeltaSequenceNo)
	fmt.Fprintf(&builder, ", ExchangeID=%+v", d.ExchangeID)
	fmt.Fprintf(&builder, ", ProductID=%+v", d.ProductID)
	fmt.Fprintf(&builder, ", ProductStatus=%+v", d.ProductStatus)

	builder.WriteByte('}')

	return builder.String()
}

// 风险结算追平持仓明细
type CThostFtdcSyncDeltaInvstPosDtlField struct {
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
	// 先开先平剩余数量
	TimeFirstVolume types.TThostFtdcVolumeType
	// 特殊持仓标志
	SpecPosiType types.TThostFtdcSpecPosiTypeType
	// 操作标志
	ActionDirection types.TThostFtdcActionDirectionType
	// 追平序号
	SyncDeltaSequenceNo types.TThostFtdcSequenceNoType
}

func (d CThostFtdcSyncDeltaInvstPosDtlField) Type() string {
	return "CThostFtdcSyncDeltaInvstPosDtlField"
}

func (d CThostFtdcSyncDeltaInvstPosDtlField) String() string {
	var builder strings.Builder
	builder.Grow(714)

	builder.WriteString("CThostFtdcSyncDeltaInvstPosDtlField{")
	fmt.Fprintf(&builder, "InstrumentID=%+v", d.InstrumentID)
	fmt.Fprintf(&builder, ", BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", InvestorID=%+v", d.InvestorID)
	fmt.Fprintf(&builder, ", HedgeFlag=%+v", d.HedgeFlag)
	fmt.Fprintf(&builder, ", Direction=%+v", d.Direction)
	fmt.Fprintf(&builder, ", OpenDate=%+v", d.OpenDate)
	fmt.Fprintf(&builder, ", TradeID=%+v", d.TradeID)
	fmt.Fprintf(&builder, ", Volume=%+v", d.Volume)
	fmt.Fprintf(&builder, ", OpenPrice=%+v", d.OpenPrice)
	fmt.Fprintf(&builder, ", TradingDay=%+v", d.TradingDay)
	fmt.Fprintf(&builder, ", SettlementID=%+v", d.SettlementID)
	fmt.Fprintf(&builder, ", TradeType=%+v", d.TradeType)
	fmt.Fprintf(&builder, ", CombInstrumentID=%+v", d.CombInstrumentID)
	fmt.Fprintf(&builder, ", ExchangeID=%+v", d.ExchangeID)
	fmt.Fprintf(&builder, ", CloseProfitByDate=%+v", d.CloseProfitByDate)
	fmt.Fprintf(&builder, ", CloseProfitByTrade=%+v", d.CloseProfitByTrade)
	fmt.Fprintf(&builder, ", PositionProfitByDate=%+v", d.PositionProfitByDate)
	fmt.Fprintf(&builder, ", PositionProfitByTrade=%+v", d.PositionProfitByTrade)
	fmt.Fprintf(&builder, ", Margin=%+v", d.Margin)
	fmt.Fprintf(&builder, ", ExchMargin=%+v", d.ExchMargin)
	fmt.Fprintf(&builder, ", MarginRateByMoney=%+v", d.MarginRateByMoney)
	fmt.Fprintf(&builder, ", MarginRateByVolume=%+v", d.MarginRateByVolume)
	fmt.Fprintf(&builder, ", LastSettlementPrice=%+v", d.LastSettlementPrice)
	fmt.Fprintf(&builder, ", SettlementPrice=%+v", d.SettlementPrice)
	fmt.Fprintf(&builder, ", CloseVolume=%+v", d.CloseVolume)
	fmt.Fprintf(&builder, ", CloseAmount=%+v", d.CloseAmount)
	fmt.Fprintf(&builder, ", TimeFirstVolume=%+v", d.TimeFirstVolume)
	fmt.Fprintf(&builder, ", SpecPosiType=%+v", d.SpecPosiType)
	fmt.Fprintf(&builder, ", ActionDirection=%+v", d.ActionDirection)
	fmt.Fprintf(&builder, ", SyncDeltaSequenceNo=%+v", d.SyncDeltaSequenceNo)

	builder.WriteByte('}')

	return builder.String()
}

// 风险结算追平组合持仓明细
type CThostFtdcSyncDeltaInvstPosCombDtlField struct {
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
	// 成交组号
	TradeGroupID types.TThostFtdcTradeGroupIDType
	// 操作标志
	ActionDirection types.TThostFtdcActionDirectionType
	// 追平序号
	SyncDeltaSequenceNo types.TThostFtdcSequenceNoType
}

func (d CThostFtdcSyncDeltaInvstPosCombDtlField) Type() string {
	return "CThostFtdcSyncDeltaInvstPosCombDtlField"
}

func (d CThostFtdcSyncDeltaInvstPosCombDtlField) String() string {
	var builder strings.Builder
	builder.Grow(475)

	builder.WriteString("CThostFtdcSyncDeltaInvstPosCombDtlField{")
	fmt.Fprintf(&builder, "TradingDay=%+v", d.TradingDay)
	fmt.Fprintf(&builder, ", OpenDate=%+v", d.OpenDate)
	fmt.Fprintf(&builder, ", ExchangeID=%+v", d.ExchangeID)
	fmt.Fprintf(&builder, ", SettlementID=%+v", d.SettlementID)
	fmt.Fprintf(&builder, ", BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", InvestorID=%+v", d.InvestorID)
	fmt.Fprintf(&builder, ", ComTradeID=%+v", d.ComTradeID)
	fmt.Fprintf(&builder, ", TradeID=%+v", d.TradeID)
	fmt.Fprintf(&builder, ", InstrumentID=%+v", d.InstrumentID)
	fmt.Fprintf(&builder, ", HedgeFlag=%+v", d.HedgeFlag)
	fmt.Fprintf(&builder, ", Direction=%+v", d.Direction)
	fmt.Fprintf(&builder, ", TotalAmt=%+v", d.TotalAmt)
	fmt.Fprintf(&builder, ", Margin=%+v", d.Margin)
	fmt.Fprintf(&builder, ", ExchMargin=%+v", d.ExchMargin)
	fmt.Fprintf(&builder, ", MarginRateByMoney=%+v", d.MarginRateByMoney)
	fmt.Fprintf(&builder, ", MarginRateByVolume=%+v", d.MarginRateByVolume)
	fmt.Fprintf(&builder, ", LegID=%+v", d.LegID)
	fmt.Fprintf(&builder, ", LegMultiple=%+v", d.LegMultiple)
	fmt.Fprintf(&builder, ", TradeGroupID=%+v", d.TradeGroupID)
	fmt.Fprintf(&builder, ", ActionDirection=%+v", d.ActionDirection)
	fmt.Fprintf(&builder, ", SyncDeltaSequenceNo=%+v", d.SyncDeltaSequenceNo)

	builder.WriteByte('}')

	return builder.String()
}

// 风险结算追平资金
type CThostFtdcSyncDeltaTradingAccountField struct {
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
	// 延时换汇冻结金额
	FrozenSwap types.TThostFtdcMoneyType
	// 剩余换汇额度
	RemainSwap types.TThostFtdcMoneyType
	// 期权市值
	OptionValue types.TThostFtdcMoneyType
	// 追平序号
	SyncDeltaSequenceNo types.TThostFtdcSequenceNoType
}

func (d CThostFtdcSyncDeltaTradingAccountField) Type() string {
	return "CThostFtdcSyncDeltaTradingAccountField"
}

func (d CThostFtdcSyncDeltaTradingAccountField) String() string {
	var builder strings.Builder
	builder.Grow(1215)

	builder.WriteString("CThostFtdcSyncDeltaTradingAccountField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", AccountID=%+v", d.AccountID)
	fmt.Fprintf(&builder, ", PreMortgage=%+v", d.PreMortgage)
	fmt.Fprintf(&builder, ", PreCredit=%+v", d.PreCredit)
	fmt.Fprintf(&builder, ", PreDeposit=%+v", d.PreDeposit)
	fmt.Fprintf(&builder, ", PreBalance=%+v", d.PreBalance)
	fmt.Fprintf(&builder, ", PreMargin=%+v", d.PreMargin)
	fmt.Fprintf(&builder, ", InterestBase=%+v", d.InterestBase)
	fmt.Fprintf(&builder, ", Interest=%+v", d.Interest)
	fmt.Fprintf(&builder, ", Deposit=%+v", d.Deposit)
	fmt.Fprintf(&builder, ", Withdraw=%+v", d.Withdraw)
	fmt.Fprintf(&builder, ", FrozenMargin=%+v", d.FrozenMargin)
	fmt.Fprintf(&builder, ", FrozenCash=%+v", d.FrozenCash)
	fmt.Fprintf(&builder, ", FrozenCommission=%+v", d.FrozenCommission)
	fmt.Fprintf(&builder, ", CurrMargin=%+v", d.CurrMargin)
	fmt.Fprintf(&builder, ", CashIn=%+v", d.CashIn)
	fmt.Fprintf(&builder, ", Commission=%+v", d.Commission)
	fmt.Fprintf(&builder, ", CloseProfit=%+v", d.CloseProfit)
	fmt.Fprintf(&builder, ", PositionProfit=%+v", d.PositionProfit)
	fmt.Fprintf(&builder, ", Balance=%+v", d.Balance)
	fmt.Fprintf(&builder, ", Available=%+v", d.Available)
	fmt.Fprintf(&builder, ", WithdrawQuota=%+v", d.WithdrawQuota)
	fmt.Fprintf(&builder, ", Reserve=%+v", d.Reserve)
	fmt.Fprintf(&builder, ", TradingDay=%+v", d.TradingDay)
	fmt.Fprintf(&builder, ", SettlementID=%+v", d.SettlementID)
	fmt.Fprintf(&builder, ", Credit=%+v", d.Credit)
	fmt.Fprintf(&builder, ", Mortgage=%+v", d.Mortgage)
	fmt.Fprintf(&builder, ", ExchangeMargin=%+v", d.ExchangeMargin)
	fmt.Fprintf(&builder, ", DeliveryMargin=%+v", d.DeliveryMargin)
	fmt.Fprintf(&builder, ", ExchangeDeliveryMargin=%+v", d.ExchangeDeliveryMargin)
	fmt.Fprintf(&builder, ", ReserveBalance=%+v", d.ReserveBalance)
	fmt.Fprintf(&builder, ", CurrencyID=%+v", d.CurrencyID)
	fmt.Fprintf(&builder, ", PreFundMortgageIn=%+v", d.PreFundMortgageIn)
	fmt.Fprintf(&builder, ", PreFundMortgageOut=%+v", d.PreFundMortgageOut)
	fmt.Fprintf(&builder, ", FundMortgageIn=%+v", d.FundMortgageIn)
	fmt.Fprintf(&builder, ", FundMortgageOut=%+v", d.FundMortgageOut)
	fmt.Fprintf(&builder, ", FundMortgageAvailable=%+v", d.FundMortgageAvailable)
	fmt.Fprintf(&builder, ", MortgageableFund=%+v", d.MortgageableFund)
	fmt.Fprintf(&builder, ", SpecProductMargin=%+v", d.SpecProductMargin)
	fmt.Fprintf(&builder, ", SpecProductFrozenMargin=%+v", d.SpecProductFrozenMargin)
	fmt.Fprintf(&builder, ", SpecProductCommission=%+v", d.SpecProductCommission)
	fmt.Fprintf(&builder, ", SpecProductFrozenCommission=%+v", d.SpecProductFrozenCommission)
	fmt.Fprintf(&builder, ", SpecProductPositionProfit=%+v", d.SpecProductPositionProfit)
	fmt.Fprintf(&builder, ", SpecProductCloseProfit=%+v", d.SpecProductCloseProfit)
	fmt.Fprintf(&builder, ", SpecProductPositionProfitByAlg=%+v", d.SpecProductPositionProfitByAlg)
	fmt.Fprintf(&builder, ", SpecProductExchangeMargin=%+v", d.SpecProductExchangeMargin)
	fmt.Fprintf(&builder, ", FrozenSwap=%+v", d.FrozenSwap)
	fmt.Fprintf(&builder, ", RemainSwap=%+v", d.RemainSwap)
	fmt.Fprintf(&builder, ", OptionValue=%+v", d.OptionValue)
	fmt.Fprintf(&builder, ", SyncDeltaSequenceNo=%+v", d.SyncDeltaSequenceNo)

	builder.WriteByte('}')

	return builder.String()
}

// 投资者风险结算总保证金
type CThostFtdcSyncDeltaInitInvstMarginField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	// 追平前总风险保证金
	LastRiskTotalInvstMargin types.TThostFtdcMoneyType
	// 追平前交易所总风险保证金
	LastRiskTotalExchMargin types.TThostFtdcMoneyType
	// 本次追平品种总保证金
	ThisSyncInvstMargin types.TThostFtdcMoneyType
	// 本次追平品种交易所总保证金
	ThisSyncExchMargin types.TThostFtdcMoneyType
	// 本次未追平品种总保证金
	RemainRiskInvstMargin types.TThostFtdcMoneyType
	// 本次未追平品种交易所总保证金
	RemainRiskExchMargin types.TThostFtdcMoneyType
	// 追平前总特殊产品风险保证金
	LastRiskSpecTotalInvstMargin types.TThostFtdcMoneyType
	// 追平前总特殊产品交易所风险保证金
	LastRiskSpecTotalExchMargin types.TThostFtdcMoneyType
	// 本次追平品种特殊产品总保证金
	ThisSyncSpecInvstMargin types.TThostFtdcMoneyType
	// 本次追平品种特殊产品交易所总保证金
	ThisSyncSpecExchMargin types.TThostFtdcMoneyType
	// 本次未追平品种特殊产品总保证金
	RemainRiskSpecInvstMargin types.TThostFtdcMoneyType
	// 本次未追平品种特殊产品交易所总保证金
	RemainRiskSpecExchMargin types.TThostFtdcMoneyType
	// 追平序号
	SyncDeltaSequenceNo types.TThostFtdcSequenceNoType
}

func (d CThostFtdcSyncDeltaInitInvstMarginField) Type() string {
	return "CThostFtdcSyncDeltaInitInvstMarginField"
}

func (d CThostFtdcSyncDeltaInitInvstMarginField) String() string {
	var builder strings.Builder
	builder.Grow(500)

	builder.WriteString("CThostFtdcSyncDeltaInitInvstMarginField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", InvestorID=%+v", d.InvestorID)
	fmt.Fprintf(&builder, ", LastRiskTotalInvstMargin=%+v", d.LastRiskTotalInvstMargin)
	fmt.Fprintf(&builder, ", LastRiskTotalExchMargin=%+v", d.LastRiskTotalExchMargin)
	fmt.Fprintf(&builder, ", ThisSyncInvstMargin=%+v", d.ThisSyncInvstMargin)
	fmt.Fprintf(&builder, ", ThisSyncExchMargin=%+v", d.ThisSyncExchMargin)
	fmt.Fprintf(&builder, ", RemainRiskInvstMargin=%+v", d.RemainRiskInvstMargin)
	fmt.Fprintf(&builder, ", RemainRiskExchMargin=%+v", d.RemainRiskExchMargin)
	fmt.Fprintf(&builder, ", LastRiskSpecTotalInvstMargin=%+v", d.LastRiskSpecTotalInvstMargin)
	fmt.Fprintf(&builder, ", LastRiskSpecTotalExchMargin=%+v", d.LastRiskSpecTotalExchMargin)
	fmt.Fprintf(&builder, ", ThisSyncSpecInvstMargin=%+v", d.ThisSyncSpecInvstMargin)
	fmt.Fprintf(&builder, ", ThisSyncSpecExchMargin=%+v", d.ThisSyncSpecExchMargin)
	fmt.Fprintf(&builder, ", RemainRiskSpecInvstMargin=%+v", d.RemainRiskSpecInvstMargin)
	fmt.Fprintf(&builder, ", RemainRiskSpecExchMargin=%+v", d.RemainRiskSpecExchMargin)
	fmt.Fprintf(&builder, ", SyncDeltaSequenceNo=%+v", d.SyncDeltaSequenceNo)

	builder.WriteByte('}')

	return builder.String()
}

// 风险结算追平组合优先级
type CThostFtdcSyncDeltaDceCombInstrumentField struct {
	// 合约代码
	CombInstrumentID types.TThostFtdcInstrumentIDType
	// 交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	// 合约在交易所的代码
	ExchangeInstID types.TThostFtdcExchangeInstIDType
	// 成交组号
	TradeGroupID types.TThostFtdcTradeGroupIDType
	// 投机套保标志
	CombHedgeFlag types.TThostFtdcHedgeFlagType
	// 组合类型
	CombinationType types.TThostFtdcDceCombinationTypeType
	// 买卖
	Direction types.TThostFtdcDirectionType
	// 产品代码
	ProductID types.TThostFtdcInstrumentIDType
	// 期货期权组合保证金比例
	Xparameter types.TThostFtdcDiscountRatioType
	// 操作标志
	ActionDirection types.TThostFtdcActionDirectionType
	// 追平序号
	SyncDeltaSequenceNo types.TThostFtdcSequenceNoType
}

func (d CThostFtdcSyncDeltaDceCombInstrumentField) Type() string {
	return "CThostFtdcSyncDeltaDceCombInstrumentField"
}

func (d CThostFtdcSyncDeltaDceCombInstrumentField) String() string {
	var builder strings.Builder
	builder.Grow(293)

	builder.WriteString("CThostFtdcSyncDeltaDceCombInstrumentField{")
	fmt.Fprintf(&builder, "CombInstrumentID=%+v", d.CombInstrumentID)
	fmt.Fprintf(&builder, ", ExchangeID=%+v", d.ExchangeID)
	fmt.Fprintf(&builder, ", ExchangeInstID=%+v", d.ExchangeInstID)
	fmt.Fprintf(&builder, ", TradeGroupID=%+v", d.TradeGroupID)
	fmt.Fprintf(&builder, ", CombHedgeFlag=%+v", d.CombHedgeFlag)
	fmt.Fprintf(&builder, ", CombinationType=%+v", d.CombinationType)
	fmt.Fprintf(&builder, ", Direction=%+v", d.Direction)
	fmt.Fprintf(&builder, ", ProductID=%+v", d.ProductID)
	fmt.Fprintf(&builder, ", Xparameter=%+v", d.Xparameter)
	fmt.Fprintf(&builder, ", ActionDirection=%+v", d.ActionDirection)
	fmt.Fprintf(&builder, ", SyncDeltaSequenceNo=%+v", d.SyncDeltaSequenceNo)

	builder.WriteByte('}')

	return builder.String()
}

// 风险结算追平投资者期货保证金率
type CThostFtdcSyncDeltaInvstMarginRateField struct {
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
	// 操作标志
	ActionDirection types.TThostFtdcActionDirectionType
	// 追平序号
	SyncDeltaSequenceNo types.TThostFtdcSequenceNoType
}

func (d CThostFtdcSyncDeltaInvstMarginRateField) Type() string {
	return "CThostFtdcSyncDeltaInvstMarginRateField"
}

func (d CThostFtdcSyncDeltaInvstMarginRateField) String() string {
	var builder strings.Builder
	builder.Grow(347)

	builder.WriteString("CThostFtdcSyncDeltaInvstMarginRateField{")
	fmt.Fprintf(&builder, "InstrumentID=%+v", d.InstrumentID)
	fmt.Fprintf(&builder, ", InvestorRange=%+v", d.InvestorRange)
	fmt.Fprintf(&builder, ", BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", InvestorID=%+v", d.InvestorID)
	fmt.Fprintf(&builder, ", HedgeFlag=%+v", d.HedgeFlag)
	fmt.Fprintf(&builder, ", LongMarginRatioByMoney=%+v", d.LongMarginRatioByMoney)
	fmt.Fprintf(&builder, ", LongMarginRatioByVolume=%+v", d.LongMarginRatioByVolume)
	fmt.Fprintf(&builder, ", ShortMarginRatioByMoney=%+v", d.ShortMarginRatioByMoney)
	fmt.Fprintf(&builder, ", ShortMarginRatioByVolume=%+v", d.ShortMarginRatioByVolume)
	fmt.Fprintf(&builder, ", IsRelative=%+v", d.IsRelative)
	fmt.Fprintf(&builder, ", ActionDirection=%+v", d.ActionDirection)
	fmt.Fprintf(&builder, ", SyncDeltaSequenceNo=%+v", d.SyncDeltaSequenceNo)

	builder.WriteByte('}')

	return builder.String()
}

// 风险结算追平交易所期货保证金率
type CThostFtdcSyncDeltaExchMarginRateField struct {
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
	// 操作标志
	ActionDirection types.TThostFtdcActionDirectionType
	// 追平序号
	SyncDeltaSequenceNo types.TThostFtdcSequenceNoType
}

func (d CThostFtdcSyncDeltaExchMarginRateField) Type() string {
	return "CThostFtdcSyncDeltaExchMarginRateField"
}

func (d CThostFtdcSyncDeltaExchMarginRateField) String() string {
	var builder strings.Builder
	builder.Grow(283)

	builder.WriteString("CThostFtdcSyncDeltaExchMarginRateField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", InstrumentID=%+v", d.InstrumentID)
	fmt.Fprintf(&builder, ", HedgeFlag=%+v", d.HedgeFlag)
	fmt.Fprintf(&builder, ", LongMarginRatioByMoney=%+v", d.LongMarginRatioByMoney)
	fmt.Fprintf(&builder, ", LongMarginRatioByVolume=%+v", d.LongMarginRatioByVolume)
	fmt.Fprintf(&builder, ", ShortMarginRatioByMoney=%+v", d.ShortMarginRatioByMoney)
	fmt.Fprintf(&builder, ", ShortMarginRatioByVolume=%+v", d.ShortMarginRatioByVolume)
	fmt.Fprintf(&builder, ", ActionDirection=%+v", d.ActionDirection)
	fmt.Fprintf(&builder, ", SyncDeltaSequenceNo=%+v", d.SyncDeltaSequenceNo)

	builder.WriteByte('}')

	return builder.String()
}

// 风险结算追平中金现货期权交易所保证金率
type CThostFtdcSyncDeltaOptExchMarginField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
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
	// 做市商空头保证金调整系数
	MShortMarginRatioByMoney types.TThostFtdcRatioType
	// 做市商空头保证金调整系数
	MShortMarginRatioByVolume types.TThostFtdcMoneyType
	// 操作标志
	ActionDirection types.TThostFtdcActionDirectionType
	// 追平序号
	SyncDeltaSequenceNo types.TThostFtdcSequenceNoType
}

func (d CThostFtdcSyncDeltaOptExchMarginField) Type() string {
	return "CThostFtdcSyncDeltaOptExchMarginField"
}

func (d CThostFtdcSyncDeltaOptExchMarginField) String() string {
	var builder strings.Builder
	builder.Grow(407)

	builder.WriteString("CThostFtdcSyncDeltaOptExchMarginField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", InstrumentID=%+v", d.InstrumentID)
	fmt.Fprintf(&builder, ", SShortMarginRatioByMoney=%+v", d.SShortMarginRatioByMoney)
	fmt.Fprintf(&builder, ", SShortMarginRatioByVolume=%+v", d.SShortMarginRatioByVolume)
	fmt.Fprintf(&builder, ", HShortMarginRatioByMoney=%+v", d.HShortMarginRatioByMoney)
	fmt.Fprintf(&builder, ", HShortMarginRatioByVolume=%+v", d.HShortMarginRatioByVolume)
	fmt.Fprintf(&builder, ", AShortMarginRatioByMoney=%+v", d.AShortMarginRatioByMoney)
	fmt.Fprintf(&builder, ", AShortMarginRatioByVolume=%+v", d.AShortMarginRatioByVolume)
	fmt.Fprintf(&builder, ", MShortMarginRatioByMoney=%+v", d.MShortMarginRatioByMoney)
	fmt.Fprintf(&builder, ", MShortMarginRatioByVolume=%+v", d.MShortMarginRatioByVolume)
	fmt.Fprintf(&builder, ", ActionDirection=%+v", d.ActionDirection)
	fmt.Fprintf(&builder, ", SyncDeltaSequenceNo=%+v", d.SyncDeltaSequenceNo)

	builder.WriteByte('}')

	return builder.String()
}

// 风险结算追平中金现货期权投资者保证金率
type CThostFtdcSyncDeltaOptInvstMarginField struct {
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
	// 操作标志
	ActionDirection types.TThostFtdcActionDirectionType
	// 追平序号
	SyncDeltaSequenceNo types.TThostFtdcSequenceNoType
}

func (d CThostFtdcSyncDeltaOptInvstMarginField) Type() string {
	return "CThostFtdcSyncDeltaOptInvstMarginField"
}

func (d CThostFtdcSyncDeltaOptInvstMarginField) String() string {
	var builder strings.Builder
	builder.Grow(471)

	builder.WriteString("CThostFtdcSyncDeltaOptInvstMarginField{")
	fmt.Fprintf(&builder, "InstrumentID=%+v", d.InstrumentID)
	fmt.Fprintf(&builder, ", InvestorRange=%+v", d.InvestorRange)
	fmt.Fprintf(&builder, ", BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", InvestorID=%+v", d.InvestorID)
	fmt.Fprintf(&builder, ", SShortMarginRatioByMoney=%+v", d.SShortMarginRatioByMoney)
	fmt.Fprintf(&builder, ", SShortMarginRatioByVolume=%+v", d.SShortMarginRatioByVolume)
	fmt.Fprintf(&builder, ", HShortMarginRatioByMoney=%+v", d.HShortMarginRatioByMoney)
	fmt.Fprintf(&builder, ", HShortMarginRatioByVolume=%+v", d.HShortMarginRatioByVolume)
	fmt.Fprintf(&builder, ", AShortMarginRatioByMoney=%+v", d.AShortMarginRatioByMoney)
	fmt.Fprintf(&builder, ", AShortMarginRatioByVolume=%+v", d.AShortMarginRatioByVolume)
	fmt.Fprintf(&builder, ", IsRelative=%+v", d.IsRelative)
	fmt.Fprintf(&builder, ", MShortMarginRatioByMoney=%+v", d.MShortMarginRatioByMoney)
	fmt.Fprintf(&builder, ", MShortMarginRatioByVolume=%+v", d.MShortMarginRatioByVolume)
	fmt.Fprintf(&builder, ", ActionDirection=%+v", d.ActionDirection)
	fmt.Fprintf(&builder, ", SyncDeltaSequenceNo=%+v", d.SyncDeltaSequenceNo)

	builder.WriteByte('}')

	return builder.String()
}

// 风险结算追平期权标的调整保证金率
type CThostFtdcSyncDeltaInvstMarginRateULField struct {
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
	// 操作标志
	ActionDirection types.TThostFtdcActionDirectionType
	// 追平序号
	SyncDeltaSequenceNo types.TThostFtdcSequenceNoType
}

func (d CThostFtdcSyncDeltaInvstMarginRateULField) Type() string {
	return "CThostFtdcSyncDeltaInvstMarginRateULField"
}

func (d CThostFtdcSyncDeltaInvstMarginRateULField) String() string {
	var builder strings.Builder
	builder.Grow(329)

	builder.WriteString("CThostFtdcSyncDeltaInvstMarginRateULField{")
	fmt.Fprintf(&builder, "InstrumentID=%+v", d.InstrumentID)
	fmt.Fprintf(&builder, ", InvestorRange=%+v", d.InvestorRange)
	fmt.Fprintf(&builder, ", BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", InvestorID=%+v", d.InvestorID)
	fmt.Fprintf(&builder, ", HedgeFlag=%+v", d.HedgeFlag)
	fmt.Fprintf(&builder, ", LongMarginRatioByMoney=%+v", d.LongMarginRatioByMoney)
	fmt.Fprintf(&builder, ", LongMarginRatioByVolume=%+v", d.LongMarginRatioByVolume)
	fmt.Fprintf(&builder, ", ShortMarginRatioByMoney=%+v", d.ShortMarginRatioByMoney)
	fmt.Fprintf(&builder, ", ShortMarginRatioByVolume=%+v", d.ShortMarginRatioByVolume)
	fmt.Fprintf(&builder, ", ActionDirection=%+v", d.ActionDirection)
	fmt.Fprintf(&builder, ", SyncDeltaSequenceNo=%+v", d.SyncDeltaSequenceNo)

	builder.WriteByte('}')

	return builder.String()
}

// 风险结算追平期权手续费率
type CThostFtdcSyncDeltaOptInvstCommRateField struct {
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
	// 操作标志
	ActionDirection types.TThostFtdcActionDirectionType
	// 追平序号
	SyncDeltaSequenceNo types.TThostFtdcSequenceNoType
}

func (d CThostFtdcSyncDeltaOptInvstCommRateField) Type() string {
	return "CThostFtdcSyncDeltaOptInvstCommRateField"
}

func (d CThostFtdcSyncDeltaOptInvstCommRateField) String() string {
	var builder strings.Builder
	builder.Grow(407)

	builder.WriteString("CThostFtdcSyncDeltaOptInvstCommRateField{")
	fmt.Fprintf(&builder, "InstrumentID=%+v", d.InstrumentID)
	fmt.Fprintf(&builder, ", InvestorRange=%+v", d.InvestorRange)
	fmt.Fprintf(&builder, ", BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", InvestorID=%+v", d.InvestorID)
	fmt.Fprintf(&builder, ", OpenRatioByMoney=%+v", d.OpenRatioByMoney)
	fmt.Fprintf(&builder, ", OpenRatioByVolume=%+v", d.OpenRatioByVolume)
	fmt.Fprintf(&builder, ", CloseRatioByMoney=%+v", d.CloseRatioByMoney)
	fmt.Fprintf(&builder, ", CloseRatioByVolume=%+v", d.CloseRatioByVolume)
	fmt.Fprintf(&builder, ", CloseTodayRatioByMoney=%+v", d.CloseTodayRatioByMoney)
	fmt.Fprintf(&builder, ", CloseTodayRatioByVolume=%+v", d.CloseTodayRatioByVolume)
	fmt.Fprintf(&builder, ", StrikeRatioByMoney=%+v", d.StrikeRatioByMoney)
	fmt.Fprintf(&builder, ", StrikeRatioByVolume=%+v", d.StrikeRatioByVolume)
	fmt.Fprintf(&builder, ", ActionDirection=%+v", d.ActionDirection)
	fmt.Fprintf(&builder, ", SyncDeltaSequenceNo=%+v", d.SyncDeltaSequenceNo)

	builder.WriteByte('}')

	return builder.String()
}

// 风险结算追平期货手续费率
type CThostFtdcSyncDeltaInvstCommRateField struct {
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
	// 操作标志
	ActionDirection types.TThostFtdcActionDirectionType
	// 追平序号
	SyncDeltaSequenceNo types.TThostFtdcSequenceNoType
}

func (d CThostFtdcSyncDeltaInvstCommRateField) Type() string {
	return "CThostFtdcSyncDeltaInvstCommRateField"
}

func (d CThostFtdcSyncDeltaInvstCommRateField) String() string {
	var builder strings.Builder
	builder.Grow(347)

	builder.WriteString("CThostFtdcSyncDeltaInvstCommRateField{")
	fmt.Fprintf(&builder, "InstrumentID=%+v", d.InstrumentID)
	fmt.Fprintf(&builder, ", InvestorRange=%+v", d.InvestorRange)
	fmt.Fprintf(&builder, ", BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", InvestorID=%+v", d.InvestorID)
	fmt.Fprintf(&builder, ", OpenRatioByMoney=%+v", d.OpenRatioByMoney)
	fmt.Fprintf(&builder, ", OpenRatioByVolume=%+v", d.OpenRatioByVolume)
	fmt.Fprintf(&builder, ", CloseRatioByMoney=%+v", d.CloseRatioByMoney)
	fmt.Fprintf(&builder, ", CloseRatioByVolume=%+v", d.CloseRatioByVolume)
	fmt.Fprintf(&builder, ", CloseTodayRatioByMoney=%+v", d.CloseTodayRatioByMoney)
	fmt.Fprintf(&builder, ", CloseTodayRatioByVolume=%+v", d.CloseTodayRatioByVolume)
	fmt.Fprintf(&builder, ", ActionDirection=%+v", d.ActionDirection)
	fmt.Fprintf(&builder, ", SyncDeltaSequenceNo=%+v", d.SyncDeltaSequenceNo)

	builder.WriteByte('}')

	return builder.String()
}

// 风险结算追平交叉汇率
type CThostFtdcSyncDeltaProductExchRateField struct {
	// 产品代码
	ProductID types.TThostFtdcInstrumentIDType
	// 报价币种类型
	QuoteCurrencyID types.TThostFtdcCurrencyIDType
	// 汇率
	ExchangeRate types.TThostFtdcExchangeRateType
	// 操作标志
	ActionDirection types.TThostFtdcActionDirectionType
	// 追平序号
	SyncDeltaSequenceNo types.TThostFtdcSequenceNoType
}

func (d CThostFtdcSyncDeltaProductExchRateField) Type() string {
	return "CThostFtdcSyncDeltaProductExchRateField"
}

func (d CThostFtdcSyncDeltaProductExchRateField) String() string {
	var builder strings.Builder
	builder.Grow(159)

	builder.WriteString("CThostFtdcSyncDeltaProductExchRateField{")
	fmt.Fprintf(&builder, "ProductID=%+v", d.ProductID)
	fmt.Fprintf(&builder, ", QuoteCurrencyID=%+v", d.QuoteCurrencyID)
	fmt.Fprintf(&builder, ", ExchangeRate=%+v", d.ExchangeRate)
	fmt.Fprintf(&builder, ", ActionDirection=%+v", d.ActionDirection)
	fmt.Fprintf(&builder, ", SyncDeltaSequenceNo=%+v", d.SyncDeltaSequenceNo)

	builder.WriteByte('}')

	return builder.String()
}

// 风险结算追平行情
type CThostFtdcSyncDeltaDepthMarketDataField struct {
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
	// 操作标志
	ActionDirection types.TThostFtdcActionDirectionType
	// 追平序号
	SyncDeltaSequenceNo types.TThostFtdcSequenceNoType
}

func (d CThostFtdcSyncDeltaDepthMarketDataField) Type() string {
	return "CThostFtdcSyncDeltaDepthMarketDataField"
}

func (d CThostFtdcSyncDeltaDepthMarketDataField) String() string {
	var builder strings.Builder
	builder.Grow(1053)

	builder.WriteString("CThostFtdcSyncDeltaDepthMarketDataField{")
	fmt.Fprintf(&builder, "TradingDay=%+v", d.TradingDay)
	fmt.Fprintf(&builder, ", InstrumentID=%+v", d.InstrumentID)
	fmt.Fprintf(&builder, ", ExchangeID=%+v", d.ExchangeID)
	fmt.Fprintf(&builder, ", ExchangeInstID=%+v", d.ExchangeInstID)
	fmt.Fprintf(&builder, ", LastPrice=%+v", d.LastPrice)
	fmt.Fprintf(&builder, ", PreSettlementPrice=%+v", d.PreSettlementPrice)
	fmt.Fprintf(&builder, ", PreClosePrice=%+v", d.PreClosePrice)
	fmt.Fprintf(&builder, ", PreOpenInterest=%+v", d.PreOpenInterest)
	fmt.Fprintf(&builder, ", OpenPrice=%+v", d.OpenPrice)
	fmt.Fprintf(&builder, ", HighestPrice=%+v", d.HighestPrice)
	fmt.Fprintf(&builder, ", LowestPrice=%+v", d.LowestPrice)
	fmt.Fprintf(&builder, ", Volume=%+v", d.Volume)
	fmt.Fprintf(&builder, ", Turnover=%+v", d.Turnover)
	fmt.Fprintf(&builder, ", OpenInterest=%+v", d.OpenInterest)
	fmt.Fprintf(&builder, ", ClosePrice=%+v", d.ClosePrice)
	fmt.Fprintf(&builder, ", SettlementPrice=%+v", d.SettlementPrice)
	fmt.Fprintf(&builder, ", UpperLimitPrice=%+v", d.UpperLimitPrice)
	fmt.Fprintf(&builder, ", LowerLimitPrice=%+v", d.LowerLimitPrice)
	fmt.Fprintf(&builder, ", PreDelta=%+v", d.PreDelta)
	fmt.Fprintf(&builder, ", CurrDelta=%+v", d.CurrDelta)
	fmt.Fprintf(&builder, ", UpdateTime=%+v", d.UpdateTime)
	fmt.Fprintf(&builder, ", UpdateMillisec=%+v", d.UpdateMillisec)
	fmt.Fprintf(&builder, ", BidPrice1=%+v", d.BidPrice1)
	fmt.Fprintf(&builder, ", BidVolume1=%+v", d.BidVolume1)
	fmt.Fprintf(&builder, ", AskPrice1=%+v", d.AskPrice1)
	fmt.Fprintf(&builder, ", AskVolume1=%+v", d.AskVolume1)
	fmt.Fprintf(&builder, ", BidPrice2=%+v", d.BidPrice2)
	fmt.Fprintf(&builder, ", BidVolume2=%+v", d.BidVolume2)
	fmt.Fprintf(&builder, ", AskPrice2=%+v", d.AskPrice2)
	fmt.Fprintf(&builder, ", AskVolume2=%+v", d.AskVolume2)
	fmt.Fprintf(&builder, ", BidPrice3=%+v", d.BidPrice3)
	fmt.Fprintf(&builder, ", BidVolume3=%+v", d.BidVolume3)
	fmt.Fprintf(&builder, ", AskPrice3=%+v", d.AskPrice3)
	fmt.Fprintf(&builder, ", AskVolume3=%+v", d.AskVolume3)
	fmt.Fprintf(&builder, ", BidPrice4=%+v", d.BidPrice4)
	fmt.Fprintf(&builder, ", BidVolume4=%+v", d.BidVolume4)
	fmt.Fprintf(&builder, ", AskPrice4=%+v", d.AskPrice4)
	fmt.Fprintf(&builder, ", AskVolume4=%+v", d.AskVolume4)
	fmt.Fprintf(&builder, ", BidPrice5=%+v", d.BidPrice5)
	fmt.Fprintf(&builder, ", BidVolume5=%+v", d.BidVolume5)
	fmt.Fprintf(&builder, ", AskPrice5=%+v", d.AskPrice5)
	fmt.Fprintf(&builder, ", AskVolume5=%+v", d.AskVolume5)
	fmt.Fprintf(&builder, ", AveragePrice=%+v", d.AveragePrice)
	fmt.Fprintf(&builder, ", ActionDay=%+v", d.ActionDay)
	fmt.Fprintf(&builder, ", BandingUpperPrice=%+v", d.BandingUpperPrice)
	fmt.Fprintf(&builder, ", BandingLowerPrice=%+v", d.BandingLowerPrice)
	fmt.Fprintf(&builder, ", ActionDirection=%+v", d.ActionDirection)
	fmt.Fprintf(&builder, ", SyncDeltaSequenceNo=%+v", d.SyncDeltaSequenceNo)

	builder.WriteByte('}')

	return builder.String()
}

// 风险结算追平现货指数
type CThostFtdcSyncDeltaIndexPriceField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
	// 指数现货收盘价
	ClosePrice types.TThostFtdcPriceType
	// 操作标志
	ActionDirection types.TThostFtdcActionDirectionType
	// 追平序号
	SyncDeltaSequenceNo types.TThostFtdcSequenceNoType
}

func (d CThostFtdcSyncDeltaIndexPriceField) Type() string {
	return "CThostFtdcSyncDeltaIndexPriceField"
}

func (d CThostFtdcSyncDeltaIndexPriceField) String() string {
	var builder strings.Builder
	builder.Grow(148)

	builder.WriteString("CThostFtdcSyncDeltaIndexPriceField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", InstrumentID=%+v", d.InstrumentID)
	fmt.Fprintf(&builder, ", ClosePrice=%+v", d.ClosePrice)
	fmt.Fprintf(&builder, ", ActionDirection=%+v", d.ActionDirection)
	fmt.Fprintf(&builder, ", SyncDeltaSequenceNo=%+v", d.SyncDeltaSequenceNo)

	builder.WriteByte('}')

	return builder.String()
}

// 风险结算追平仓单折抵
type CThostFtdcSyncDeltaEWarrantOffsetField struct {
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
	// 操作标志
	ActionDirection types.TThostFtdcActionDirectionType
	// 追平序号
	SyncDeltaSequenceNo types.TThostFtdcSequenceNoType
}

func (d CThostFtdcSyncDeltaEWarrantOffsetField) Type() string {
	return "CThostFtdcSyncDeltaEWarrantOffsetField"
}

func (d CThostFtdcSyncDeltaEWarrantOffsetField) String() string {
	var builder strings.Builder
	builder.Grow(246)

	builder.WriteString("CThostFtdcSyncDeltaEWarrantOffsetField{")
	fmt.Fprintf(&builder, "TradingDay=%+v", d.TradingDay)
	fmt.Fprintf(&builder, ", BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", InvestorID=%+v", d.InvestorID)
	fmt.Fprintf(&builder, ", ExchangeID=%+v", d.ExchangeID)
	fmt.Fprintf(&builder, ", InstrumentID=%+v", d.InstrumentID)
	fmt.Fprintf(&builder, ", Direction=%+v", d.Direction)
	fmt.Fprintf(&builder, ", HedgeFlag=%+v", d.HedgeFlag)
	fmt.Fprintf(&builder, ", Volume=%+v", d.Volume)
	fmt.Fprintf(&builder, ", ActionDirection=%+v", d.ActionDirection)
	fmt.Fprintf(&builder, ", SyncDeltaSequenceNo=%+v", d.SyncDeltaSequenceNo)

	builder.WriteByte('}')

	return builder.String()
}

// SPBM期货合约保证金参数
type CThostFtdcSPBMFutureParameterField struct {
	// 交易日
	TradingDay types.TThostFtdcDateType
	// 交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	// 合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
	// 品种代码
	ProdFamilyCode types.TThostFtdcInstrumentIDType
	// 期货合约因子
	Cvf types.TThostFtdcVolumeMultipleType
	// 阶段标识
	TimeRange types.TThostFtdcTimeRangeType
	// 品种保证金标准
	MarginRate types.TThostFtdcRatioType
	// 期货合约内部对锁仓费率折扣比例
	LockRateX types.TThostFtdcRatioType
	// 提高保证金标准
	AddOnRate types.TThostFtdcRatioType
	// 昨结算价
	PreSettlementPrice types.TThostFtdcPriceType
	// 期货合约内部对锁仓附加费率折扣比例
	AddOnLockRateX2 types.TThostFtdcRatioType
}

func (d CThostFtdcSPBMFutureParameterField) Type() string {
	return "CThostFtdcSPBMFutureParameterField"
}

func (d CThostFtdcSPBMFutureParameterField) String() string {
	var builder strings.Builder
	builder.Grow(263)

	builder.WriteString("CThostFtdcSPBMFutureParameterField{")
	fmt.Fprintf(&builder, "TradingDay=%+v", d.TradingDay)
	fmt.Fprintf(&builder, ", ExchangeID=%+v", d.ExchangeID)
	fmt.Fprintf(&builder, ", InstrumentID=%+v", d.InstrumentID)
	fmt.Fprintf(&builder, ", ProdFamilyCode=%+v", d.ProdFamilyCode)
	fmt.Fprintf(&builder, ", Cvf=%+v", d.Cvf)
	fmt.Fprintf(&builder, ", TimeRange=%+v", d.TimeRange)
	fmt.Fprintf(&builder, ", MarginRate=%+v", d.MarginRate)
	fmt.Fprintf(&builder, ", LockRateX=%+v", d.LockRateX)
	fmt.Fprintf(&builder, ", AddOnRate=%+v", d.AddOnRate)
	fmt.Fprintf(&builder, ", PreSettlementPrice=%+v", d.PreSettlementPrice)
	fmt.Fprintf(&builder, ", AddOnLockRateX2=%+v", d.AddOnLockRateX2)

	builder.WriteByte('}')

	return builder.String()
}

// SPBM期权合约保证金参数
type CThostFtdcSPBMOptionParameterField struct {
	// 交易日
	TradingDay types.TThostFtdcDateType
	// 交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	// 合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
	// 品种代码
	ProdFamilyCode types.TThostFtdcInstrumentIDType
	// 期权合约因子
	Cvf types.TThostFtdcVolumeMultipleType
	// 期权冲抵价格
	DownPrice types.TThostFtdcPriceType
	// Delta值
	Delta types.TThostFtdcDeltaType
	// 卖方期权风险转换最低值
	SlimiDelta types.TThostFtdcDeltaType
	// 昨结算价
	PreSettlementPrice types.TThostFtdcPriceType
}

func (d CThostFtdcSPBMOptionParameterField) Type() string {
	return "CThostFtdcSPBMOptionParameterField"
}

func (d CThostFtdcSPBMOptionParameterField) String() string {
	var builder strings.Builder
	builder.Grow(215)

	builder.WriteString("CThostFtdcSPBMOptionParameterField{")
	fmt.Fprintf(&builder, "TradingDay=%+v", d.TradingDay)
	fmt.Fprintf(&builder, ", ExchangeID=%+v", d.ExchangeID)
	fmt.Fprintf(&builder, ", InstrumentID=%+v", d.InstrumentID)
	fmt.Fprintf(&builder, ", ProdFamilyCode=%+v", d.ProdFamilyCode)
	fmt.Fprintf(&builder, ", Cvf=%+v", d.Cvf)
	fmt.Fprintf(&builder, ", DownPrice=%+v", d.DownPrice)
	fmt.Fprintf(&builder, ", Delta=%+v", d.Delta)
	fmt.Fprintf(&builder, ", SlimiDelta=%+v", d.SlimiDelta)
	fmt.Fprintf(&builder, ", PreSettlementPrice=%+v", d.PreSettlementPrice)

	builder.WriteByte('}')

	return builder.String()
}

// SPBM品种内对锁仓折扣参数
type CThostFtdcSPBMIntraParameterField struct {
	// 交易日
	TradingDay types.TThostFtdcDateType
	// 交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	// 品种代码
	ProdFamilyCode types.TThostFtdcInstrumentIDType
	// 品种内合约间对锁仓费率折扣比例
	IntraRateY types.TThostFtdcRatioType
	// 品种内合约间对锁仓附加费率折扣比例
	AddOnIntraRateY2 types.TThostFtdcRatioType
}

func (d CThostFtdcSPBMIntraParameterField) Type() string { return "CThostFtdcSPBMIntraParameterField" }

func (d CThostFtdcSPBMIntraParameterField) String() string {
	var builder strings.Builder
	builder.Grow(143)

	builder.WriteString("CThostFtdcSPBMIntraParameterField{")
	fmt.Fprintf(&builder, "TradingDay=%+v", d.TradingDay)
	fmt.Fprintf(&builder, ", ExchangeID=%+v", d.ExchangeID)
	fmt.Fprintf(&builder, ", ProdFamilyCode=%+v", d.ProdFamilyCode)
	fmt.Fprintf(&builder, ", IntraRateY=%+v", d.IntraRateY)
	fmt.Fprintf(&builder, ", AddOnIntraRateY2=%+v", d.AddOnIntraRateY2)

	builder.WriteByte('}')

	return builder.String()
}

// SPBM跨品种抵扣参数
type CThostFtdcSPBMInterParameterField struct {
	// 交易日
	TradingDay types.TThostFtdcDateType
	// 交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	// 优先级
	SpreadId types.TThostFtdcSpreadIdType
	// 品种间对锁仓费率折扣比例
	InterRateZ types.TThostFtdcRatioType
	// 第一腿构成品种
	Leg1ProdFamilyCode types.TThostFtdcInstrumentIDType
	// 第二腿构成品种
	Leg2ProdFamilyCode types.TThostFtdcInstrumentIDType
}

func (d CThostFtdcSPBMInterParameterField) Type() string { return "CThostFtdcSPBMInterParameterField" }

func (d CThostFtdcSPBMInterParameterField) String() string {
	var builder strings.Builder
	builder.Grow(167)

	builder.WriteString("CThostFtdcSPBMInterParameterField{")
	fmt.Fprintf(&builder, "TradingDay=%+v", d.TradingDay)
	fmt.Fprintf(&builder, ", ExchangeID=%+v", d.ExchangeID)
	fmt.Fprintf(&builder, ", SpreadId=%+v", d.SpreadId)
	fmt.Fprintf(&builder, ", InterRateZ=%+v", d.InterRateZ)
	fmt.Fprintf(&builder, ", Leg1ProdFamilyCode=%+v", d.Leg1ProdFamilyCode)
	fmt.Fprintf(&builder, ", Leg2ProdFamilyCode=%+v", d.Leg2ProdFamilyCode)

	builder.WriteByte('}')

	return builder.String()
}

// 同步SPBM参数结束
type CThostFtdcSyncSPBMParameterEndField struct {
	// 交易日
	TradingDay types.TThostFtdcDateType
}

func (d CThostFtdcSyncSPBMParameterEndField) Type() string {
	return "CThostFtdcSyncSPBMParameterEndField"
}

func (d CThostFtdcSyncSPBMParameterEndField) String() string {
	var builder strings.Builder
	builder.Grow(55)

	builder.WriteString("CThostFtdcSyncSPBMParameterEndField{")
	fmt.Fprintf(&builder, "TradingDay=%+v", d.TradingDay)

	builder.WriteByte('}')

	return builder.String()
}

// SPBM期货合约保证金参数查询
type CThostFtdcQrySPBMFutureParameterField struct {
	// 交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	// 合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
	// 品种代码
	ProdFamilyCode types.TThostFtdcInstrumentIDType
}

func (d CThostFtdcQrySPBMFutureParameterField) Type() string {
	return "CThostFtdcQrySPBMFutureParameterField"
}

func (d CThostFtdcQrySPBMFutureParameterField) String() string {
	var builder strings.Builder
	builder.Grow(103)

	builder.WriteString("CThostFtdcQrySPBMFutureParameterField{")
	fmt.Fprintf(&builder, "ExchangeID=%+v", d.ExchangeID)
	fmt.Fprintf(&builder, ", InstrumentID=%+v", d.InstrumentID)
	fmt.Fprintf(&builder, ", ProdFamilyCode=%+v", d.ProdFamilyCode)

	builder.WriteByte('}')

	return builder.String()
}

// SPBM期权合约保证金参数查询
type CThostFtdcQrySPBMOptionParameterField struct {
	// 交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	// 合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
	// 品种代码
	ProdFamilyCode types.TThostFtdcInstrumentIDType
}

func (d CThostFtdcQrySPBMOptionParameterField) Type() string {
	return "CThostFtdcQrySPBMOptionParameterField"
}

func (d CThostFtdcQrySPBMOptionParameterField) String() string {
	var builder strings.Builder
	builder.Grow(103)

	builder.WriteString("CThostFtdcQrySPBMOptionParameterField{")
	fmt.Fprintf(&builder, "ExchangeID=%+v", d.ExchangeID)
	fmt.Fprintf(&builder, ", InstrumentID=%+v", d.InstrumentID)
	fmt.Fprintf(&builder, ", ProdFamilyCode=%+v", d.ProdFamilyCode)

	builder.WriteByte('}')

	return builder.String()
}

// SPBM品种内对锁仓折扣参数查询
type CThostFtdcQrySPBMIntraParameterField struct {
	// 交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	// 品种代码
	ProdFamilyCode types.TThostFtdcInstrumentIDType
}

func (d CThostFtdcQrySPBMIntraParameterField) Type() string {
	return "CThostFtdcQrySPBMIntraParameterField"
}

func (d CThostFtdcQrySPBMIntraParameterField) String() string {
	var builder strings.Builder
	builder.Grow(80)

	builder.WriteString("CThostFtdcQrySPBMIntraParameterField{")
	fmt.Fprintf(&builder, "ExchangeID=%+v", d.ExchangeID)
	fmt.Fprintf(&builder, ", ProdFamilyCode=%+v", d.ProdFamilyCode)

	builder.WriteByte('}')

	return builder.String()
}

// SPBM跨品种抵扣参数查询
type CThostFtdcQrySPBMInterParameterField struct {
	// 交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	// 第一腿构成品种
	Leg1ProdFamilyCode types.TThostFtdcInstrumentIDType
	// 第二腿构成品种
	Leg2ProdFamilyCode types.TThostFtdcInstrumentIDType
}

func (d CThostFtdcQrySPBMInterParameterField) Type() string {
	return "CThostFtdcQrySPBMInterParameterField"
}

func (d CThostFtdcQrySPBMInterParameterField) String() string {
	var builder strings.Builder
	builder.Grow(112)

	builder.WriteString("CThostFtdcQrySPBMInterParameterField{")
	fmt.Fprintf(&builder, "ExchangeID=%+v", d.ExchangeID)
	fmt.Fprintf(&builder, ", Leg1ProdFamilyCode=%+v", d.Leg1ProdFamilyCode)
	fmt.Fprintf(&builder, ", Leg2ProdFamilyCode=%+v", d.Leg2ProdFamilyCode)

	builder.WriteByte('}')

	return builder.String()
}

// 组合保证金套餐
type CThostFtdcSPBMPortfDefinitionField struct {
	// 交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	// 组合保证金套餐代码
	PortfolioDefID types.TThostFtdcPortfolioDefIDType
	// 品种代码
	ProdFamilyCode types.TThostFtdcInstrumentIDType
	// 是否启用SPBM
	IsSPBM types.TThostFtdcBoolType
}

func (d CThostFtdcSPBMPortfDefinitionField) Type() string {
	return "CThostFtdcSPBMPortfDefinitionField"
}

func (d CThostFtdcSPBMPortfDefinitionField) String() string {
	var builder strings.Builder
	builder.Grow(118)

	builder.WriteString("CThostFtdcSPBMPortfDefinitionField{")
	fmt.Fprintf(&builder, "ExchangeID=%+v", d.ExchangeID)
	fmt.Fprintf(&builder, ", PortfolioDefID=%+v", d.PortfolioDefID)
	fmt.Fprintf(&builder, ", ProdFamilyCode=%+v", d.ProdFamilyCode)
	fmt.Fprintf(&builder, ", IsSPBM=%+v", d.IsSPBM)

	builder.WriteByte('}')

	return builder.String()
}

// 投资者套餐选择
type CThostFtdcSPBMInvestorPortfDefField struct {
	// 交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	// 组合保证金套餐代码
	PortfolioDefID types.TThostFtdcPortfolioDefIDType
}

func (d CThostFtdcSPBMInvestorPortfDefField) Type() string {
	return "CThostFtdcSPBMInvestorPortfDefField"
}

func (d CThostFtdcSPBMInvestorPortfDefField) String() string {
	var builder strings.Builder
	builder.Grow(117)

	builder.WriteString("CThostFtdcSPBMInvestorPortfDefField{")
	fmt.Fprintf(&builder, "ExchangeID=%+v", d.ExchangeID)
	fmt.Fprintf(&builder, ", BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", InvestorID=%+v", d.InvestorID)
	fmt.Fprintf(&builder, ", PortfolioDefID=%+v", d.PortfolioDefID)

	builder.WriteByte('}')

	return builder.String()
}

// 投资者新型组合保证金系数
type CThostFtdcInvestorPortfMarginRatioField struct {
	// 投资者范围
	InvestorRange types.TThostFtdcInvestorRangeType
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	// 交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	// 会员对投资者收取的保证金和交易所对投资者收取的保证金的比例
	MarginRatio types.TThostFtdcRatioType
	// 产品群代码
	ProductGroupID types.TThostFtdcProductIDType
}

func (d CThostFtdcInvestorPortfMarginRatioField) Type() string {
	return "CThostFtdcInvestorPortfMarginRatioField"
}

func (d CThostFtdcInvestorPortfMarginRatioField) String() string {
	var builder strings.Builder
	builder.Grow(165)

	builder.WriteString("CThostFtdcInvestorPortfMarginRatioField{")
	fmt.Fprintf(&builder, "InvestorRange=%+v", d.InvestorRange)
	fmt.Fprintf(&builder, ", BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", InvestorID=%+v", d.InvestorID)
	fmt.Fprintf(&builder, ", ExchangeID=%+v", d.ExchangeID)
	fmt.Fprintf(&builder, ", MarginRatio=%+v", d.MarginRatio)
	fmt.Fprintf(&builder, ", ProductGroupID=%+v", d.ProductGroupID)

	builder.WriteByte('}')

	return builder.String()
}

// 组合保证金套餐查询
type CThostFtdcQrySPBMPortfDefinitionField struct {
	// 交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	// 组合保证金套餐代码
	PortfolioDefID types.TThostFtdcPortfolioDefIDType
	// 品种代码
	ProdFamilyCode types.TThostFtdcInstrumentIDType
}

func (d CThostFtdcQrySPBMPortfDefinitionField) Type() string {
	return "CThostFtdcQrySPBMPortfDefinitionField"
}

func (d CThostFtdcQrySPBMPortfDefinitionField) String() string {
	var builder strings.Builder
	builder.Grow(105)

	builder.WriteString("CThostFtdcQrySPBMPortfDefinitionField{")
	fmt.Fprintf(&builder, "ExchangeID=%+v", d.ExchangeID)
	fmt.Fprintf(&builder, ", PortfolioDefID=%+v", d.PortfolioDefID)
	fmt.Fprintf(&builder, ", ProdFamilyCode=%+v", d.ProdFamilyCode)

	builder.WriteByte('}')

	return builder.String()
}

// 投资者套餐选择查询
type CThostFtdcQrySPBMInvestorPortfDefField struct {
	// 交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID types.TThostFtdcInvestorIDType
}

func (d CThostFtdcQrySPBMInvestorPortfDefField) Type() string {
	return "CThostFtdcQrySPBMInvestorPortfDefField"
}

func (d CThostFtdcQrySPBMInvestorPortfDefField) String() string {
	var builder strings.Builder
	builder.Grow(96)

	builder.WriteString("CThostFtdcQrySPBMInvestorPortfDefField{")
	fmt.Fprintf(&builder, "ExchangeID=%+v", d.ExchangeID)
	fmt.Fprintf(&builder, ", BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", InvestorID=%+v", d.InvestorID)

	builder.WriteByte('}')

	return builder.String()
}

// 投资者新型组合保证金系数查询
type CThostFtdcQryInvestorPortfMarginRatioField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	// 交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	// 产品群代码
	ProductGroupID types.TThostFtdcProductIDType
}

func (d CThostFtdcQryInvestorPortfMarginRatioField) Type() string {
	return "CThostFtdcQryInvestorPortfMarginRatioField"
}

func (d CThostFtdcQryInvestorPortfMarginRatioField) String() string {
	var builder strings.Builder
	builder.Grow(124)

	builder.WriteString("CThostFtdcQryInvestorPortfMarginRatioField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", InvestorID=%+v", d.InvestorID)
	fmt.Fprintf(&builder, ", ExchangeID=%+v", d.ExchangeID)
	fmt.Fprintf(&builder, ", ProductGroupID=%+v", d.ProductGroupID)

	builder.WriteByte('}')

	return builder.String()
}

// 投资者产品SPBM明细
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
	// 买归集保证金
	BCollectingMargin types.TThostFtdcMoneyType
	// 卖归集保证金
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
	// 交割月保证金
	DeliveryMargin types.TThostFtdcMoneyType
	// 看涨期权最低风险
	CallOptionMinRisk types.TThostFtdcMoneyType
	// 看跌期权最低风险
	PutOptionMinRisk types.TThostFtdcMoneyType
	// 卖方期权最低风险
	OptionMinRisk types.TThostFtdcMoneyType
	// 买方期权冲抵价值
	OptionValueOffset types.TThostFtdcMoneyType
	// 卖方期权权利金
	OptionRoyalty types.TThostFtdcMoneyType
	// 价值冲抵
	RealOptionValueOffset types.TThostFtdcMoneyType
	// 保证金
	Margin types.TThostFtdcMoneyType
	// 交易所保证金
	ExchMargin types.TThostFtdcMoneyType
}

func (d CThostFtdcInvestorProdSPBMDetailField) Type() string {
	return "CThostFtdcInvestorProdSPBMDetailField"
}

func (d CThostFtdcInvestorProdSPBMDetailField) String() string {
	var builder strings.Builder
	builder.Grow(528)

	builder.WriteString("CThostFtdcInvestorProdSPBMDetailField{")
	fmt.Fprintf(&builder, "ExchangeID=%+v", d.ExchangeID)
	fmt.Fprintf(&builder, ", BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", InvestorID=%+v", d.InvestorID)
	fmt.Fprintf(&builder, ", ProdFamilyCode=%+v", d.ProdFamilyCode)
	fmt.Fprintf(&builder, ", IntraInstrMargin=%+v", d.IntraInstrMargin)
	fmt.Fprintf(&builder, ", BCollectingMargin=%+v", d.BCollectingMargin)
	fmt.Fprintf(&builder, ", SCollectingMargin=%+v", d.SCollectingMargin)
	fmt.Fprintf(&builder, ", IntraProdMargin=%+v", d.IntraProdMargin)
	fmt.Fprintf(&builder, ", NetMargin=%+v", d.NetMargin)
	fmt.Fprintf(&builder, ", InterProdMargin=%+v", d.InterProdMargin)
	fmt.Fprintf(&builder, ", SingleMargin=%+v", d.SingleMargin)
	fmt.Fprintf(&builder, ", AddOnMargin=%+v", d.AddOnMargin)
	fmt.Fprintf(&builder, ", DeliveryMargin=%+v", d.DeliveryMargin)
	fmt.Fprintf(&builder, ", CallOptionMinRisk=%+v", d.CallOptionMinRisk)
	fmt.Fprintf(&builder, ", PutOptionMinRisk=%+v", d.PutOptionMinRisk)
	fmt.Fprintf(&builder, ", OptionMinRisk=%+v", d.OptionMinRisk)
	fmt.Fprintf(&builder, ", OptionValueOffset=%+v", d.OptionValueOffset)
	fmt.Fprintf(&builder, ", OptionRoyalty=%+v", d.OptionRoyalty)
	fmt.Fprintf(&builder, ", RealOptionValueOffset=%+v", d.RealOptionValueOffset)
	fmt.Fprintf(&builder, ", Margin=%+v", d.Margin)
	fmt.Fprintf(&builder, ", ExchMargin=%+v", d.ExchMargin)

	builder.WriteByte('}')

	return builder.String()
}

// 投资者产品SPBM明细查询
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

func (d CThostFtdcQryInvestorProdSPBMDetailField) Type() string {
	return "CThostFtdcQryInvestorProdSPBMDetailField"
}

func (d CThostFtdcQryInvestorProdSPBMDetailField) String() string {
	var builder strings.Builder
	builder.Grow(122)

	builder.WriteString("CThostFtdcQryInvestorProdSPBMDetailField{")
	fmt.Fprintf(&builder, "ExchangeID=%+v", d.ExchangeID)
	fmt.Fprintf(&builder, ", BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", InvestorID=%+v", d.InvestorID)
	fmt.Fprintf(&builder, ", ProdFamilyCode=%+v", d.ProdFamilyCode)

	builder.WriteByte('}')

	return builder.String()
}

// 组保交易参数设置
type CThostFtdcPortfTradeParamSettingField struct {
	// 交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	// 组保算法
	Portfolio types.TThostFtdcPortfolioType
	// 撤单是否验资
	IsActionVerify types.TThostFtdcBoolType
	// 平仓是否验资
	IsCloseVerify types.TThostFtdcBoolType
}

func (d CThostFtdcPortfTradeParamSettingField) Type() string {
	return "CThostFtdcPortfTradeParamSettingField"
}

func (d CThostFtdcPortfTradeParamSettingField) String() string {
	var builder strings.Builder
	builder.Grow(161)

	builder.WriteString("CThostFtdcPortfTradeParamSettingField{")
	fmt.Fprintf(&builder, "ExchangeID=%+v", d.ExchangeID)
	fmt.Fprintf(&builder, ", BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", InvestorID=%+v", d.InvestorID)
	fmt.Fprintf(&builder, ", Portfolio=%+v", d.Portfolio)
	fmt.Fprintf(&builder, ", IsActionVerify=%+v", d.IsActionVerify)
	fmt.Fprintf(&builder, ", IsCloseVerify=%+v", d.IsCloseVerify)

	builder.WriteByte('}')

	return builder.String()
}

// 投资者交易权限设置
type CThostFtdcInvestorTradingRightField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	// 交易权限
	InvstTradingRight types.TThostFtdcInvstTradingRightType
}

func (d CThostFtdcInvestorTradingRightField) Type() string {
	return "CThostFtdcInvestorTradingRightField"
}

func (d CThostFtdcInvestorTradingRightField) String() string {
	var builder strings.Builder
	builder.Grow(100)

	builder.WriteString("CThostFtdcInvestorTradingRightField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", InvestorID=%+v", d.InvestorID)
	fmt.Fprintf(&builder, ", InvstTradingRight=%+v", d.InvstTradingRight)

	builder.WriteByte('}')

	return builder.String()
}

// 质押配比参数
type CThostFtdcMortgageParamField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 投资者帐号
	AccountID types.TThostFtdcAccountIDType
	// 质押配比系数
	MortgageBalance types.TThostFtdcRatioType
	// 开仓是否验证质押配比
	CheckMortgageRatio types.TThostFtdcBoolType
}

func (d CThostFtdcMortgageParamField) Type() string { return "CThostFtdcMortgageParamField" }

func (d CThostFtdcMortgageParamField) String() string {
	var builder strings.Builder
	builder.Grow(118)

	builder.WriteString("CThostFtdcMortgageParamField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", AccountID=%+v", d.AccountID)
	fmt.Fprintf(&builder, ", MortgageBalance=%+v", d.MortgageBalance)
	fmt.Fprintf(&builder, ", CheckMortgageRatio=%+v", d.CheckMortgageRatio)

	builder.WriteByte('}')

	return builder.String()
}

// 可提控制参数
type CThostFtdcWithDrawParamField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 投资者帐号
	AccountID types.TThostFtdcAccountIDType
	// 参数代码
	WithDrawParamID types.TThostFtdcWithDrawParamIDType
	// 参数代码值
	WithDrawParamValue types.TThostFtdcWithDrawParamValueType
}

func (d CThostFtdcWithDrawParamField) Type() string { return "CThostFtdcWithDrawParamField" }

func (d CThostFtdcWithDrawParamField) String() string {
	var builder strings.Builder
	builder.Grow(118)

	builder.WriteString("CThostFtdcWithDrawParamField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", AccountID=%+v", d.AccountID)
	fmt.Fprintf(&builder, ", WithDrawParamID=%+v", d.WithDrawParamID)
	fmt.Fprintf(&builder, ", WithDrawParamValue=%+v", d.WithDrawParamValue)

	builder.WriteByte('}')

	return builder.String()
}

// Thost终端用户功能权限
type CThostFtdcThostUserFunctionField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 用户代码
	UserID types.TThostFtdcUserIDType
	// Thost终端功能代码
	ThostFunctionCode types.TThostFtdcThostFunctionCodeType
}

func (d CThostFtdcThostUserFunctionField) Type() string { return "CThostFtdcThostUserFunctionField" }

func (d CThostFtdcThostUserFunctionField) String() string {
	var builder strings.Builder
	builder.Grow(93)

	builder.WriteString("CThostFtdcThostUserFunctionField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", UserID=%+v", d.UserID)
	fmt.Fprintf(&builder, ", ThostFunctionCode=%+v", d.ThostFunctionCode)

	builder.WriteByte('}')

	return builder.String()
}

// Thost终端用户功能权限查询
type CThostFtdcQryThostUserFunctionField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 用户代码
	UserID types.TThostFtdcUserIDType
}

func (d CThostFtdcQryThostUserFunctionField) Type() string {
	return "CThostFtdcQryThostUserFunctionField"
}

func (d CThostFtdcQryThostUserFunctionField) String() string {
	var builder strings.Builder
	builder.Grow(69)

	builder.WriteString("CThostFtdcQryThostUserFunctionField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", UserID=%+v", d.UserID)

	builder.WriteByte('}')

	return builder.String()
}

// SPBM附加跨品种抵扣参数
type CThostFtdcSPBMAddOnInterParameterField struct {
	// 交易日
	TradingDay types.TThostFtdcDateType
	// 交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	// 优先级
	SpreadId types.TThostFtdcSpreadIdType
	// 品种间对锁仓附加费率折扣比例
	AddOnInterRateZ2 types.TThostFtdcRatioType
	// 第一腿构成品种
	Leg1ProdFamilyCode types.TThostFtdcInstrumentIDType
	// 第二腿构成品种
	Leg2ProdFamilyCode types.TThostFtdcInstrumentIDType
}

func (d CThostFtdcSPBMAddOnInterParameterField) Type() string {
	return "CThostFtdcSPBMAddOnInterParameterField"
}

func (d CThostFtdcSPBMAddOnInterParameterField) String() string {
	var builder strings.Builder
	builder.Grow(178)

	builder.WriteString("CThostFtdcSPBMAddOnInterParameterField{")
	fmt.Fprintf(&builder, "TradingDay=%+v", d.TradingDay)
	fmt.Fprintf(&builder, ", ExchangeID=%+v", d.ExchangeID)
	fmt.Fprintf(&builder, ", SpreadId=%+v", d.SpreadId)
	fmt.Fprintf(&builder, ", AddOnInterRateZ2=%+v", d.AddOnInterRateZ2)
	fmt.Fprintf(&builder, ", Leg1ProdFamilyCode=%+v", d.Leg1ProdFamilyCode)
	fmt.Fprintf(&builder, ", Leg2ProdFamilyCode=%+v", d.Leg2ProdFamilyCode)

	builder.WriteByte('}')

	return builder.String()
}

// SPBM附加跨品种抵扣参数查询
type CThostFtdcQrySPBMAddOnInterParameterField struct {
	// 交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	// 第一腿构成品种
	Leg1ProdFamilyCode types.TThostFtdcInstrumentIDType
	// 第二腿构成品种
	Leg2ProdFamilyCode types.TThostFtdcInstrumentIDType
}

func (d CThostFtdcQrySPBMAddOnInterParameterField) Type() string {
	return "CThostFtdcQrySPBMAddOnInterParameterField"
}

func (d CThostFtdcQrySPBMAddOnInterParameterField) String() string {
	var builder strings.Builder
	builder.Grow(117)

	builder.WriteString("CThostFtdcQrySPBMAddOnInterParameterField{")
	fmt.Fprintf(&builder, "ExchangeID=%+v", d.ExchangeID)
	fmt.Fprintf(&builder, ", Leg1ProdFamilyCode=%+v", d.Leg1ProdFamilyCode)
	fmt.Fprintf(&builder, ", Leg2ProdFamilyCode=%+v", d.Leg2ProdFamilyCode)

	builder.WriteByte('}')

	return builder.String()
}

// 投资者商品组SPMM记录查询
type CThostFtdcQryInvestorCommoditySPMMMarginField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	// 商品组代码
	CommodityID types.TThostFtdcSPMMProductIDType
}

func (d CThostFtdcQryInvestorCommoditySPMMMarginField) Type() string {
	return "CThostFtdcQryInvestorCommoditySPMMMarginField"
}

func (d CThostFtdcQryInvestorCommoditySPMMMarginField) String() string {
	var builder strings.Builder
	builder.Grow(104)

	builder.WriteString("CThostFtdcQryInvestorCommoditySPMMMarginField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", InvestorID=%+v", d.InvestorID)
	fmt.Fprintf(&builder, ", CommodityID=%+v", d.CommodityID)

	builder.WriteByte('}')

	return builder.String()
}

// 投资者商品群SPMM记录查询
type CThostFtdcQryInvestorCommodityGroupSPMMMarginField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	// 商品群代码
	CommodityGroupID types.TThostFtdcSPMMProductIDType
}

func (d CThostFtdcQryInvestorCommodityGroupSPMMMarginField) Type() string {
	return "CThostFtdcQryInvestorCommodityGroupSPMMMarginField"
}

func (d CThostFtdcQryInvestorCommodityGroupSPMMMarginField) String() string {
	var builder strings.Builder
	builder.Grow(114)

	builder.WriteString("CThostFtdcQryInvestorCommodityGroupSPMMMarginField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", InvestorID=%+v", d.InvestorID)
	fmt.Fprintf(&builder, ", CommodityGroupID=%+v", d.CommodityGroupID)

	builder.WriteByte('}')

	return builder.String()
}

// SPMM合约参数查询
type CThostFtdcQrySPMMInstParamField struct {
	// 合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
}

func (d CThostFtdcQrySPMMInstParamField) Type() string { return "CThostFtdcQrySPMMInstParamField" }

func (d CThostFtdcQrySPMMInstParamField) String() string {
	var builder strings.Builder
	builder.Grow(53)

	builder.WriteString("CThostFtdcQrySPMMInstParamField{")
	fmt.Fprintf(&builder, "InstrumentID=%+v", d.InstrumentID)

	builder.WriteByte('}')

	return builder.String()
}

// SPMM产品参数查询
type CThostFtdcQrySPMMProductParamField struct {
	// 产品代码
	ProductID types.TThostFtdcSPMMProductIDType
}

func (d CThostFtdcQrySPMMProductParamField) Type() string {
	return "CThostFtdcQrySPMMProductParamField"
}

func (d CThostFtdcQrySPMMProductParamField) String() string {
	var builder strings.Builder
	builder.Grow(53)

	builder.WriteString("CThostFtdcQrySPMMProductParamField{")
	fmt.Fprintf(&builder, "ProductID=%+v", d.ProductID)

	builder.WriteByte('}')

	return builder.String()
}

// 投资者商品组SPMM记录
type CThostFtdcInvestorCommoditySPMMMarginField struct {
	// 交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	// 商品组代码
	CommodityID types.TThostFtdcSPMMProductIDType
	// 优惠仓位应收保证金
	MarginBeforeDiscount types.TThostFtdcMoneyType
	// 不优惠仓位应收保证金
	MarginNoDiscount types.TThostFtdcMoneyType
	// 多头实仓风险
	LongPosRisk types.TThostFtdcMoneyType
	// 多头开仓冻结风险
	LongOpenFrozenRisk types.TThostFtdcMoneyType
	// 多头被平冻结风险
	LongCloseFrozenRisk types.TThostFtdcMoneyType
	// 空头实仓风险
	ShortPosRisk types.TThostFtdcMoneyType
	// 空头开仓冻结风险
	ShortOpenFrozenRisk types.TThostFtdcMoneyType
	// 空头被平冻结风险
	ShortCloseFrozenRisk types.TThostFtdcMoneyType
	// SPMM品种内跨期优惠系数
	IntraCommodityRate types.TThostFtdcSPMMDiscountRatioType
	// SPMM期权优惠系数
	OptionDiscountRate types.TThostFtdcSPMMDiscountRatioType
	// 实仓对冲优惠金额
	PosDiscount types.TThostFtdcMoneyType
	// 开仓报单对冲优惠金额
	OpenFrozenDiscount types.TThostFtdcMoneyType
	// 品种风险净头
	NetRisk types.TThostFtdcMoneyType
	// 平仓冻结保证金
	CloseFrozenMargin types.TThostFtdcMoneyType
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

func (d CThostFtdcInvestorCommoditySPMMMarginField) Type() string {
	return "CThostFtdcInvestorCommoditySPMMMarginField"
}

func (d CThostFtdcInvestorCommoditySPMMMarginField) String() string {
	var builder strings.Builder
	builder.Grow(595)

	builder.WriteString("CThostFtdcInvestorCommoditySPMMMarginField{")
	fmt.Fprintf(&builder, "ExchangeID=%+v", d.ExchangeID)
	fmt.Fprintf(&builder, ", BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", InvestorID=%+v", d.InvestorID)
	fmt.Fprintf(&builder, ", CommodityID=%+v", d.CommodityID)
	fmt.Fprintf(&builder, ", MarginBeforeDiscount=%+v", d.MarginBeforeDiscount)
	fmt.Fprintf(&builder, ", MarginNoDiscount=%+v", d.MarginNoDiscount)
	fmt.Fprintf(&builder, ", LongPosRisk=%+v", d.LongPosRisk)
	fmt.Fprintf(&builder, ", LongOpenFrozenRisk=%+v", d.LongOpenFrozenRisk)
	fmt.Fprintf(&builder, ", LongCloseFrozenRisk=%+v", d.LongCloseFrozenRisk)
	fmt.Fprintf(&builder, ", ShortPosRisk=%+v", d.ShortPosRisk)
	fmt.Fprintf(&builder, ", ShortOpenFrozenRisk=%+v", d.ShortOpenFrozenRisk)
	fmt.Fprintf(&builder, ", ShortCloseFrozenRisk=%+v", d.ShortCloseFrozenRisk)
	fmt.Fprintf(&builder, ", IntraCommodityRate=%+v", d.IntraCommodityRate)
	fmt.Fprintf(&builder, ", OptionDiscountRate=%+v", d.OptionDiscountRate)
	fmt.Fprintf(&builder, ", PosDiscount=%+v", d.PosDiscount)
	fmt.Fprintf(&builder, ", OpenFrozenDiscount=%+v", d.OpenFrozenDiscount)
	fmt.Fprintf(&builder, ", NetRisk=%+v", d.NetRisk)
	fmt.Fprintf(&builder, ", CloseFrozenMargin=%+v", d.CloseFrozenMargin)
	fmt.Fprintf(&builder, ", FrozenCommission=%+v", d.FrozenCommission)
	fmt.Fprintf(&builder, ", Commission=%+v", d.Commission)
	fmt.Fprintf(&builder, ", FrozenCash=%+v", d.FrozenCash)
	fmt.Fprintf(&builder, ", CashIn=%+v", d.CashIn)
	fmt.Fprintf(&builder, ", StrikeFrozenMargin=%+v", d.StrikeFrozenMargin)

	builder.WriteByte('}')

	return builder.String()
}

// 投资者商品群SPMM记录
type CThostFtdcInvestorCommodityGroupSPMMMarginField struct {
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
	InterCommodityRate types.TThostFtdcSPMMDiscountRatioType
	// 商品群最小保证金比例
	MiniMarginRatio types.TThostFtdcSPMMDiscountRatioType
	// 投资者保证金和交易所保证金的比例
	AdjustRatio types.TThostFtdcRatioType
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

func (d CThostFtdcInvestorCommodityGroupSPMMMarginField) Type() string {
	return "CThostFtdcInvestorCommodityGroupSPMMMarginField"
}

func (d CThostFtdcInvestorCommodityGroupSPMMMarginField) String() string {
	var builder strings.Builder
	builder.Grow(543)

	builder.WriteString("CThostFtdcInvestorCommodityGroupSPMMMarginField{")
	fmt.Fprintf(&builder, "ExchangeID=%+v", d.ExchangeID)
	fmt.Fprintf(&builder, ", BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", InvestorID=%+v", d.InvestorID)
	fmt.Fprintf(&builder, ", CommodityGroupID=%+v", d.CommodityGroupID)
	fmt.Fprintf(&builder, ", MarginBeforeDiscount=%+v", d.MarginBeforeDiscount)
	fmt.Fprintf(&builder, ", MarginNoDiscount=%+v", d.MarginNoDiscount)
	fmt.Fprintf(&builder, ", LongRisk=%+v", d.LongRisk)
	fmt.Fprintf(&builder, ", ShortRisk=%+v", d.ShortRisk)
	fmt.Fprintf(&builder, ", CloseFrozenMargin=%+v", d.CloseFrozenMargin)
	fmt.Fprintf(&builder, ", InterCommodityRate=%+v", d.InterCommodityRate)
	fmt.Fprintf(&builder, ", MiniMarginRatio=%+v", d.MiniMarginRatio)
	fmt.Fprintf(&builder, ", AdjustRatio=%+v", d.AdjustRatio)
	fmt.Fprintf(&builder, ", IntraCommodityDiscount=%+v", d.IntraCommodityDiscount)
	fmt.Fprintf(&builder, ", InterCommodityDiscount=%+v", d.InterCommodityDiscount)
	fmt.Fprintf(&builder, ", ExchMargin=%+v", d.ExchMargin)
	fmt.Fprintf(&builder, ", InvestorMargin=%+v", d.InvestorMargin)
	fmt.Fprintf(&builder, ", FrozenCommission=%+v", d.FrozenCommission)
	fmt.Fprintf(&builder, ", Commission=%+v", d.Commission)
	fmt.Fprintf(&builder, ", FrozenCash=%+v", d.FrozenCash)
	fmt.Fprintf(&builder, ", CashIn=%+v", d.CashIn)
	fmt.Fprintf(&builder, ", StrikeFrozenMargin=%+v", d.StrikeFrozenMargin)

	builder.WriteByte('}')

	return builder.String()
}

// SPMM合约参数
type CThostFtdcSPMMInstParamField struct {
	// 交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	// 合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
	// SPMM合约保证金算法
	InstMarginCalID types.TThostFtdcInstMarginCalIDType
	// 商品组代码
	CommodityID types.TThostFtdcSPMMProductIDType
	// 商品群代码
	CommodityGroupID types.TThostFtdcSPMMProductIDType
}

func (d CThostFtdcSPMMInstParamField) Type() string { return "CThostFtdcSPMMInstParamField" }

func (d CThostFtdcSPMMInstParamField) String() string {
	var builder strings.Builder
	builder.Grow(142)

	builder.WriteString("CThostFtdcSPMMInstParamField{")
	fmt.Fprintf(&builder, "ExchangeID=%+v", d.ExchangeID)
	fmt.Fprintf(&builder, ", InstrumentID=%+v", d.InstrumentID)
	fmt.Fprintf(&builder, ", InstMarginCalID=%+v", d.InstMarginCalID)
	fmt.Fprintf(&builder, ", CommodityID=%+v", d.CommodityID)
	fmt.Fprintf(&builder, ", CommodityGroupID=%+v", d.CommodityGroupID)

	builder.WriteByte('}')

	return builder.String()
}

// SPMM产品参数
type CThostFtdcSPMMProductParamField struct {
	// 交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	// 产品代码
	ProductID types.TThostFtdcSPMMProductIDType
	// 商品组代码
	CommodityID types.TThostFtdcSPMMProductIDType
	// 商品群代码
	CommodityGroupID types.TThostFtdcSPMMProductIDType
}

func (d CThostFtdcSPMMProductParamField) Type() string { return "CThostFtdcSPMMProductParamField" }

func (d CThostFtdcSPMMProductParamField) String() string {
	var builder strings.Builder
	builder.Grow(117)

	builder.WriteString("CThostFtdcSPMMProductParamField{")
	fmt.Fprintf(&builder, "ExchangeID=%+v", d.ExchangeID)
	fmt.Fprintf(&builder, ", ProductID=%+v", d.ProductID)
	fmt.Fprintf(&builder, ", CommodityID=%+v", d.CommodityID)
	fmt.Fprintf(&builder, ", CommodityGroupID=%+v", d.CommodityGroupID)

	builder.WriteByte('}')

	return builder.String()
}

// 席位与交易中心对应关系维护查询
type CThostFtdcQryTraderAssignField struct {
	// 交易员代码
	TraderID types.TThostFtdcTraderIDType
}

func (d CThostFtdcQryTraderAssignField) Type() string { return "CThostFtdcQryTraderAssignField" }

func (d CThostFtdcQryTraderAssignField) String() string {
	var builder strings.Builder
	builder.Grow(48)

	builder.WriteString("CThostFtdcQryTraderAssignField{")
	fmt.Fprintf(&builder, "TraderID=%+v", d.TraderID)

	builder.WriteByte('}')

	return builder.String()
}

// 席位与交易中心对应关系
type CThostFtdcTraderAssignField struct {
	// 应用单元代码
	BrokerID types.TThostFtdcBrokerIDType
	// 交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	// 交易所交易员代码
	TraderID types.TThostFtdcTraderIDType
	// 会员代码
	ParticipantID types.TThostFtdcParticipantIDType
	// 交易中心代码
	DRIdentityID types.TThostFtdcDRIdentityIDType
}

func (d CThostFtdcTraderAssignField) Type() string { return "CThostFtdcTraderAssignField" }

func (d CThostFtdcTraderAssignField) String() string {
	var builder strings.Builder
	builder.Grow(128)

	builder.WriteString("CThostFtdcTraderAssignField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", ExchangeID=%+v", d.ExchangeID)
	fmt.Fprintf(&builder, ", TraderID=%+v", d.TraderID)
	fmt.Fprintf(&builder, ", ParticipantID=%+v", d.ParticipantID)
	fmt.Fprintf(&builder, ", DRIdentityID=%+v", d.DRIdentityID)

	builder.WriteByte('}')

	return builder.String()
}

// 投资者申报费阶梯收取设置
type CThostFtdcInvestorInfoCntSettingField struct {
	// 交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	// 商品代码
	ProductID types.TThostFtdcProductIDType
	// 是否收取申报费
	IsCalInfoComm types.TThostFtdcBoolType
	// 是否限制信息量
	IsLimitInfoMax types.TThostFtdcBoolType
	// 信息量限制笔数
	InfoMaxLimit types.TThostFtdcVolumeType
}

func (d CThostFtdcInvestorInfoCntSettingField) Type() string {
	return "CThostFtdcInvestorInfoCntSettingField"
}

func (d CThostFtdcInvestorInfoCntSettingField) String() string {
	var builder strings.Builder
	builder.Grow(183)

	builder.WriteString("CThostFtdcInvestorInfoCntSettingField{")
	fmt.Fprintf(&builder, "ExchangeID=%+v", d.ExchangeID)
	fmt.Fprintf(&builder, ", BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", InvestorID=%+v", d.InvestorID)
	fmt.Fprintf(&builder, ", ProductID=%+v", d.ProductID)
	fmt.Fprintf(&builder, ", IsCalInfoComm=%+v", d.IsCalInfoComm)
	fmt.Fprintf(&builder, ", IsLimitInfoMax=%+v", d.IsLimitInfoMax)
	fmt.Fprintf(&builder, ", InfoMaxLimit=%+v", d.InfoMaxLimit)

	builder.WriteByte('}')

	return builder.String()
}

// RCAMS产品组合信息
type CThostFtdcRCAMSCombProductInfoField struct {
	// 交易日
	TradingDay types.TThostFtdcDateType
	// 交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	// 产品代码
	ProductID types.TThostFtdcProductIDType
	// 商品组代码
	CombProductID types.TThostFtdcProductIDType
	// 商品群代码
	ProductGroupID types.TThostFtdcProductIDType
}

func (d CThostFtdcRCAMSCombProductInfoField) Type() string {
	return "CThostFtdcRCAMSCombProductInfoField"
}

func (d CThostFtdcRCAMSCombProductInfoField) String() string {
	var builder strings.Builder
	builder.Grow(141)

	builder.WriteString("CThostFtdcRCAMSCombProductInfoField{")
	fmt.Fprintf(&builder, "TradingDay=%+v", d.TradingDay)
	fmt.Fprintf(&builder, ", ExchangeID=%+v", d.ExchangeID)
	fmt.Fprintf(&builder, ", ProductID=%+v", d.ProductID)
	fmt.Fprintf(&builder, ", CombProductID=%+v", d.CombProductID)
	fmt.Fprintf(&builder, ", ProductGroupID=%+v", d.ProductGroupID)

	builder.WriteByte('}')

	return builder.String()
}

// RCAMS同合约风险对冲参数
type CThostFtdcRCAMSInstrParameterField struct {
	// 交易日
	TradingDay types.TThostFtdcDateType
	// 交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	// 产品代码
	ProductID types.TThostFtdcProductIDType
	// 同合约风险对冲比率
	HedgeRate types.TThostFtdcHedgeRateType
}

func (d CThostFtdcRCAMSInstrParameterField) Type() string {
	return "CThostFtdcRCAMSInstrParameterField"
}

func (d CThostFtdcRCAMSInstrParameterField) String() string {
	var builder strings.Builder
	builder.Grow(112)

	builder.WriteString("CThostFtdcRCAMSInstrParameterField{")
	fmt.Fprintf(&builder, "TradingDay=%+v", d.TradingDay)
	fmt.Fprintf(&builder, ", ExchangeID=%+v", d.ExchangeID)
	fmt.Fprintf(&builder, ", ProductID=%+v", d.ProductID)
	fmt.Fprintf(&builder, ", HedgeRate=%+v", d.HedgeRate)

	builder.WriteByte('}')

	return builder.String()
}

// RCAMS品种内风险对冲参数
type CThostFtdcRCAMSIntraParameterField struct {
	// 交易日
	TradingDay types.TThostFtdcDateType
	// 交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	// 产品组合代码
	CombProductID types.TThostFtdcProductIDType
	// 品种内对冲比率
	HedgeRate types.TThostFtdcHedgeRateType
}

func (d CThostFtdcRCAMSIntraParameterField) Type() string {
	return "CThostFtdcRCAMSIntraParameterField"
}

func (d CThostFtdcRCAMSIntraParameterField) String() string {
	var builder strings.Builder
	builder.Grow(116)

	builder.WriteString("CThostFtdcRCAMSIntraParameterField{")
	fmt.Fprintf(&builder, "TradingDay=%+v", d.TradingDay)
	fmt.Fprintf(&builder, ", ExchangeID=%+v", d.ExchangeID)
	fmt.Fprintf(&builder, ", CombProductID=%+v", d.CombProductID)
	fmt.Fprintf(&builder, ", HedgeRate=%+v", d.HedgeRate)

	builder.WriteByte('}')

	return builder.String()
}

// RCAMS跨品种风险折抵参数
type CThostFtdcRCAMSInterParameterField struct {
	// 交易日
	TradingDay types.TThostFtdcDateType
	// 交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	// 商品群代码
	ProductGroupID types.TThostFtdcProductIDType
	// 优先级
	Priority types.TThostFtdcRCAMSPriorityType
	// 折抵率
	CreditRate types.TThostFtdcHedgeRateType
	// 产品组合代码1
	CombProduct1 types.TThostFtdcProductIDType
	// 产品组合代码2
	CombProduct2 types.TThostFtdcProductIDType
}

func (d CThostFtdcRCAMSInterParameterField) Type() string {
	return "CThostFtdcRCAMSInterParameterField"
}

func (d CThostFtdcRCAMSInterParameterField) String() string {
	var builder strings.Builder
	builder.Grow(180)

	builder.WriteString("CThostFtdcRCAMSInterParameterField{")
	fmt.Fprintf(&builder, "TradingDay=%+v", d.TradingDay)
	fmt.Fprintf(&builder, ", ExchangeID=%+v", d.ExchangeID)
	fmt.Fprintf(&builder, ", ProductGroupID=%+v", d.ProductGroupID)
	fmt.Fprintf(&builder, ", Priority=%+v", d.Priority)
	fmt.Fprintf(&builder, ", CreditRate=%+v", d.CreditRate)
	fmt.Fprintf(&builder, ", CombProduct1=%+v", d.CombProduct1)
	fmt.Fprintf(&builder, ", CombProduct2=%+v", d.CombProduct2)

	builder.WriteByte('}')

	return builder.String()
}

// RCAMS空头期权风险调整参数
type CThostFtdcRCAMSShortOptAdjustParamField struct {
	// 交易日
	TradingDay types.TThostFtdcDateType
	// 交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	// 产品组合代码
	CombProductID types.TThostFtdcProductIDType
	// 投套标志
	HedgeFlag types.TThostFtdcHedgeFlagType
	// 空头期权风险调整标准
	AdjustValue types.TThostFtdcAdjustValueType
}

func (d CThostFtdcRCAMSShortOptAdjustParamField) Type() string {
	return "CThostFtdcRCAMSShortOptAdjustParamField"
}

func (d CThostFtdcRCAMSShortOptAdjustParamField) String() string {
	var builder strings.Builder
	builder.Grow(142)

	builder.WriteString("CThostFtdcRCAMSShortOptAdjustParamField{")
	fmt.Fprintf(&builder, "TradingDay=%+v", d.TradingDay)
	fmt.Fprintf(&builder, ", ExchangeID=%+v", d.ExchangeID)
	fmt.Fprintf(&builder, ", CombProductID=%+v", d.CombProductID)
	fmt.Fprintf(&builder, ", HedgeFlag=%+v", d.HedgeFlag)
	fmt.Fprintf(&builder, ", AdjustValue=%+v", d.AdjustValue)

	builder.WriteByte('}')

	return builder.String()
}

// RCAMS策略组合持仓
type CThostFtdcRCAMSInvestorCombPositionField struct {
	// 交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	// 合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
	// 投套标志
	HedgeFlag types.TThostFtdcHedgeFlagType
	// 持仓多空方向
	PosiDirection types.TThostFtdcPosiDirectionType
	// 组合合约代码
	CombInstrumentID types.TThostFtdcInstrumentIDType
	// 单腿编号
	LegID types.TThostFtdcLegIDType
	// 交易所组合合约代码
	ExchangeInstID types.TThostFtdcExchangeInstIDType
	// 持仓量
	TotalAmt types.TThostFtdcVolumeType
	// 交易所保证金
	ExchMargin types.TThostFtdcMoneyType
	// 投资者保证金
	Margin types.TThostFtdcMoneyType
}

func (d CThostFtdcRCAMSInvestorCombPositionField) Type() string {
	return "CThostFtdcRCAMSInvestorCombPositionField"
}

func (d CThostFtdcRCAMSInvestorCombPositionField) String() string {
	var builder strings.Builder
	builder.Grow(281)

	builder.WriteString("CThostFtdcRCAMSInvestorCombPositionField{")
	fmt.Fprintf(&builder, "ExchangeID=%+v", d.ExchangeID)
	fmt.Fprintf(&builder, ", BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", InvestorID=%+v", d.InvestorID)
	fmt.Fprintf(&builder, ", InstrumentID=%+v", d.InstrumentID)
	fmt.Fprintf(&builder, ", HedgeFlag=%+v", d.HedgeFlag)
	fmt.Fprintf(&builder, ", PosiDirection=%+v", d.PosiDirection)
	fmt.Fprintf(&builder, ", CombInstrumentID=%+v", d.CombInstrumentID)
	fmt.Fprintf(&builder, ", LegID=%+v", d.LegID)
	fmt.Fprintf(&builder, ", ExchangeInstID=%+v", d.ExchangeInstID)
	fmt.Fprintf(&builder, ", TotalAmt=%+v", d.TotalAmt)
	fmt.Fprintf(&builder, ", ExchMargin=%+v", d.ExchMargin)
	fmt.Fprintf(&builder, ", Margin=%+v", d.Margin)

	builder.WriteByte('}')

	return builder.String()
}

// 投资者品种RCAMS保证金
type CThostFtdcInvestorProdRCAMSMarginField struct {
	// 交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	// 产品组合代码
	CombProductID types.TThostFtdcProductIDType
	// 投套标志
	HedgeFlag types.TThostFtdcHedgeFlagType
	// 商品群代码
	ProductGroupID types.TThostFtdcProductIDType
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
	// 空头期权风险调整
	ShortOptRiskAdj types.TThostFtdcMoneyType
	// 空头期权权利金
	OptionRoyalty types.TThostFtdcMoneyType
	// 大边组合平仓冻结保证金
	MMSACloseFrozenMargin types.TThostFtdcMoneyType
	// 策略组合平仓行权冻结保证金
	CloseCombFrozenMargin types.TThostFtdcMoneyType
	// 平仓行权冻结保证金
	CloseFrozenMargin types.TThostFtdcMoneyType
	// 大边组合开仓冻结保证金
	MMSAOpenFrozenMargin types.TThostFtdcMoneyType
	// 交割月期货开仓冻结保证金
	DeliveryOpenFrozenMargin types.TThostFtdcMoneyType
	// 开仓冻结保证金
	OpenFrozenMargin types.TThostFtdcMoneyType
	// 投资者冻结保证金
	UseFrozenMargin types.TThostFtdcMoneyType
	// 大边组合交易所持仓保证金
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

func (d CThostFtdcInvestorProdRCAMSMarginField) Type() string {
	return "CThostFtdcInvestorProdRCAMSMarginField"
}

func (d CThostFtdcInvestorProdRCAMSMarginField) String() string {
	var builder strings.Builder
	builder.Grow(680)

	builder.WriteString("CThostFtdcInvestorProdRCAMSMarginField{")
	fmt.Fprintf(&builder, "ExchangeID=%+v", d.ExchangeID)
	fmt.Fprintf(&builder, ", BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", InvestorID=%+v", d.InvestorID)
	fmt.Fprintf(&builder, ", CombProductID=%+v", d.CombProductID)
	fmt.Fprintf(&builder, ", HedgeFlag=%+v", d.HedgeFlag)
	fmt.Fprintf(&builder, ", ProductGroupID=%+v", d.ProductGroupID)
	fmt.Fprintf(&builder, ", RiskBeforeDiscount=%+v", d.RiskBeforeDiscount)
	fmt.Fprintf(&builder, ", IntraInstrRisk=%+v", d.IntraInstrRisk)
	fmt.Fprintf(&builder, ", BPosRisk=%+v", d.BPosRisk)
	fmt.Fprintf(&builder, ", SPosRisk=%+v", d.SPosRisk)
	fmt.Fprintf(&builder, ", IntraProdRisk=%+v", d.IntraProdRisk)
	fmt.Fprintf(&builder, ", NetRisk=%+v", d.NetRisk)
	fmt.Fprintf(&builder, ", InterProdRisk=%+v", d.InterProdRisk)
	fmt.Fprintf(&builder, ", ShortOptRiskAdj=%+v", d.ShortOptRiskAdj)
	fmt.Fprintf(&builder, ", OptionRoyalty=%+v", d.OptionRoyalty)
	fmt.Fprintf(&builder, ", MMSACloseFrozenMargin=%+v", d.MMSACloseFrozenMargin)
	fmt.Fprintf(&builder, ", CloseCombFrozenMargin=%+v", d.CloseCombFrozenMargin)
	fmt.Fprintf(&builder, ", CloseFrozenMargin=%+v", d.CloseFrozenMargin)
	fmt.Fprintf(&builder, ", MMSAOpenFrozenMargin=%+v", d.MMSAOpenFrozenMargin)
	fmt.Fprintf(&builder, ", DeliveryOpenFrozenMargin=%+v", d.DeliveryOpenFrozenMargin)
	fmt.Fprintf(&builder, ", OpenFrozenMargin=%+v", d.OpenFrozenMargin)
	fmt.Fprintf(&builder, ", UseFrozenMargin=%+v", d.UseFrozenMargin)
	fmt.Fprintf(&builder, ", MMSAExchMargin=%+v", d.MMSAExchMargin)
	fmt.Fprintf(&builder, ", DeliveryExchMargin=%+v", d.DeliveryExchMargin)
	fmt.Fprintf(&builder, ", CombExchMargin=%+v", d.CombExchMargin)
	fmt.Fprintf(&builder, ", ExchMargin=%+v", d.ExchMargin)
	fmt.Fprintf(&builder, ", UseMargin=%+v", d.UseMargin)

	builder.WriteByte('}')

	return builder.String()
}

// RCAMS产品组合信息查询
type CThostFtdcQryRCAMSCombProductInfoField struct {
	// 产品代码
	ProductID types.TThostFtdcProductIDType
	// 商品组代码
	CombProductID types.TThostFtdcProductIDType
	// 商品群代码
	ProductGroupID types.TThostFtdcProductIDType
}

func (d CThostFtdcQryRCAMSCombProductInfoField) Type() string {
	return "CThostFtdcQryRCAMSCombProductInfoField"
}

func (d CThostFtdcQryRCAMSCombProductInfoField) String() string {
	var builder strings.Builder
	builder.Grow(104)

	builder.WriteString("CThostFtdcQryRCAMSCombProductInfoField{")
	fmt.Fprintf(&builder, "ProductID=%+v", d.ProductID)
	fmt.Fprintf(&builder, ", CombProductID=%+v", d.CombProductID)
	fmt.Fprintf(&builder, ", ProductGroupID=%+v", d.ProductGroupID)

	builder.WriteByte('}')

	return builder.String()
}

// RCAMS同合约风险对冲参数查询
type CThostFtdcQryRCAMSInstrParameterField struct {
	// 产品代码
	ProductID types.TThostFtdcProductIDType
}

func (d CThostFtdcQryRCAMSInstrParameterField) Type() string {
	return "CThostFtdcQryRCAMSInstrParameterField"
}

func (d CThostFtdcQryRCAMSInstrParameterField) String() string {
	var builder strings.Builder
	builder.Grow(56)

	builder.WriteString("CThostFtdcQryRCAMSInstrParameterField{")
	fmt.Fprintf(&builder, "ProductID=%+v", d.ProductID)

	builder.WriteByte('}')

	return builder.String()
}

// RCAMS品种内风险对冲参数查询
type CThostFtdcQryRCAMSIntraParameterField struct {
	// 产品组合代码
	CombProductID types.TThostFtdcProductIDType
}

func (d CThostFtdcQryRCAMSIntraParameterField) Type() string {
	return "CThostFtdcQryRCAMSIntraParameterField"
}

func (d CThostFtdcQryRCAMSIntraParameterField) String() string {
	var builder strings.Builder
	builder.Grow(60)

	builder.WriteString("CThostFtdcQryRCAMSIntraParameterField{")
	fmt.Fprintf(&builder, "CombProductID=%+v", d.CombProductID)

	builder.WriteByte('}')

	return builder.String()
}

// RCAMS跨品种风险折抵参数查询
type CThostFtdcQryRCAMSInterParameterField struct {
	// 商品群代码
	ProductGroupID types.TThostFtdcProductIDType
	// 产品组合代码1
	CombProduct1 types.TThostFtdcProductIDType
	// 产品组合代码2
	CombProduct2 types.TThostFtdcProductIDType
}

func (d CThostFtdcQryRCAMSInterParameterField) Type() string {
	return "CThostFtdcQryRCAMSInterParameterField"
}

func (d CThostFtdcQryRCAMSInterParameterField) String() string {
	var builder strings.Builder
	builder.Grow(105)

	builder.WriteString("CThostFtdcQryRCAMSInterParameterField{")
	fmt.Fprintf(&builder, "ProductGroupID=%+v", d.ProductGroupID)
	fmt.Fprintf(&builder, ", CombProduct1=%+v", d.CombProduct1)
	fmt.Fprintf(&builder, ", CombProduct2=%+v", d.CombProduct2)

	builder.WriteByte('}')

	return builder.String()
}

// RCAMS空头期权风险调整参数查询
type CThostFtdcQryRCAMSShortOptAdjustParamField struct {
	// 产品组合代码
	CombProductID types.TThostFtdcProductIDType
}

func (d CThostFtdcQryRCAMSShortOptAdjustParamField) Type() string {
	return "CThostFtdcQryRCAMSShortOptAdjustParamField"
}

func (d CThostFtdcQryRCAMSShortOptAdjustParamField) String() string {
	var builder strings.Builder
	builder.Grow(65)

	builder.WriteString("CThostFtdcQryRCAMSShortOptAdjustParamField{")
	fmt.Fprintf(&builder, "CombProductID=%+v", d.CombProductID)

	builder.WriteByte('}')

	return builder.String()
}

// RCAMS策略组合持仓查询
type CThostFtdcQryRCAMSInvestorCombPositionField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	// 合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
	// 组合合约代码
	CombInstrumentID types.TThostFtdcInstrumentIDType
}

func (d CThostFtdcQryRCAMSInvestorCombPositionField) Type() string {
	return "CThostFtdcQryRCAMSInvestorCombPositionField"
}

func (d CThostFtdcQryRCAMSInvestorCombPositionField) String() string {
	var builder strings.Builder
	builder.Grow(129)

	builder.WriteString("CThostFtdcQryRCAMSInvestorCombPositionField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", InvestorID=%+v", d.InvestorID)
	fmt.Fprintf(&builder, ", InstrumentID=%+v", d.InstrumentID)
	fmt.Fprintf(&builder, ", CombInstrumentID=%+v", d.CombInstrumentID)

	builder.WriteByte('}')

	return builder.String()
}

// 投资者品种RCAMS保证金查询
type CThostFtdcQryInvestorProdRCAMSMarginField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	// 产品组合代码
	CombProductID types.TThostFtdcProductIDType
	// 商品群代码
	ProductGroupID types.TThostFtdcProductIDType
}

func (d CThostFtdcQryInvestorProdRCAMSMarginField) Type() string {
	return "CThostFtdcQryInvestorProdRCAMSMarginField"
}

func (d CThostFtdcQryInvestorProdRCAMSMarginField) String() string {
	var builder strings.Builder
	builder.Grow(126)

	builder.WriteString("CThostFtdcQryInvestorProdRCAMSMarginField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", InvestorID=%+v", d.InvestorID)
	fmt.Fprintf(&builder, ", CombProductID=%+v", d.CombProductID)
	fmt.Fprintf(&builder, ", ProductGroupID=%+v", d.ProductGroupID)

	builder.WriteByte('}')

	return builder.String()
}

// RULE合约保证金参数
type CThostFtdcRULEInstrParameterField struct {
	// 交易日
	TradingDay types.TThostFtdcDateType
	// 交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	// 合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
	// 合约类型
	InstrumentClass types.TThostFtdcInstrumentClassType
	// 标准合约
	StdInstrumentID types.TThostFtdcInstrumentIDType
	// 投机买折算系数
	BSpecRatio types.TThostFtdcRatioType
	// 投机卖折算系数
	SSpecRatio types.TThostFtdcRatioType
	// 套保买折算系数
	BHedgeRatio types.TThostFtdcRatioType
	// 套保卖折算系数
	SHedgeRatio types.TThostFtdcRatioType
	// 买附加风险保证金
	BAddOnMargin types.TThostFtdcMoneyType
	// 卖附加风险保证金
	SAddOnMargin types.TThostFtdcMoneyType
	// 商品群号
	CommodityGroupID types.TThostFtdcCommodityGroupIDType
}

func (d CThostFtdcRULEInstrParameterField) Type() string { return "CThostFtdcRULEInstrParameterField" }

func (d CThostFtdcRULEInstrParameterField) String() string {
	var builder strings.Builder
	builder.Grow(297)

	builder.WriteString("CThostFtdcRULEInstrParameterField{")
	fmt.Fprintf(&builder, "TradingDay=%+v", d.TradingDay)
	fmt.Fprintf(&builder, ", ExchangeID=%+v", d.ExchangeID)
	fmt.Fprintf(&builder, ", InstrumentID=%+v", d.InstrumentID)
	fmt.Fprintf(&builder, ", InstrumentClass=%+v", d.InstrumentClass)
	fmt.Fprintf(&builder, ", StdInstrumentID=%+v", d.StdInstrumentID)
	fmt.Fprintf(&builder, ", BSpecRatio=%+v", d.BSpecRatio)
	fmt.Fprintf(&builder, ", SSpecRatio=%+v", d.SSpecRatio)
	fmt.Fprintf(&builder, ", BHedgeRatio=%+v", d.BHedgeRatio)
	fmt.Fprintf(&builder, ", SHedgeRatio=%+v", d.SHedgeRatio)
	fmt.Fprintf(&builder, ", BAddOnMargin=%+v", d.BAddOnMargin)
	fmt.Fprintf(&builder, ", SAddOnMargin=%+v", d.SAddOnMargin)
	fmt.Fprintf(&builder, ", CommodityGroupID=%+v", d.CommodityGroupID)

	builder.WriteByte('}')

	return builder.String()
}

// RULE品种内对锁仓折扣参数
type CThostFtdcRULEIntraParameterField struct {
	// 交易日
	TradingDay types.TThostFtdcDateType
	// 交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	// 品种代码
	ProdFamilyCode types.TThostFtdcInstrumentIDType
	// 标准合约
	StdInstrumentID types.TThostFtdcInstrumentIDType
	// 标准合约保证金
	StdInstrMargin types.TThostFtdcMoneyType
	// 一般月份合约组合保证金系数
	UsualIntraRate types.TThostFtdcRatioType
	// 临近交割合约组合保证金系数
	DeliveryIntraRate types.TThostFtdcRatioType
}

func (d CThostFtdcRULEIntraParameterField) Type() string { return "CThostFtdcRULEIntraParameterField" }

func (d CThostFtdcRULEIntraParameterField) String() string {
	var builder strings.Builder
	builder.Grow(197)

	builder.WriteString("CThostFtdcRULEIntraParameterField{")
	fmt.Fprintf(&builder, "TradingDay=%+v", d.TradingDay)
	fmt.Fprintf(&builder, ", ExchangeID=%+v", d.ExchangeID)
	fmt.Fprintf(&builder, ", ProdFamilyCode=%+v", d.ProdFamilyCode)
	fmt.Fprintf(&builder, ", StdInstrumentID=%+v", d.StdInstrumentID)
	fmt.Fprintf(&builder, ", StdInstrMargin=%+v", d.StdInstrMargin)
	fmt.Fprintf(&builder, ", UsualIntraRate=%+v", d.UsualIntraRate)
	fmt.Fprintf(&builder, ", DeliveryIntraRate=%+v", d.DeliveryIntraRate)

	builder.WriteByte('}')

	return builder.String()
}

// RULE跨品种抵扣参数
type CThostFtdcRULEInterParameterField struct {
	// 交易日
	TradingDay types.TThostFtdcDateType
	// 交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	// 优先级
	SpreadId types.TThostFtdcSpreadIdType
	// 品种间对锁仓费率折扣比例
	InterRate types.TThostFtdcRatioType
	// 第一腿构成品种
	Leg1ProdFamilyCode types.TThostFtdcInstrumentIDType
	// 第二腿构成品种
	Leg2ProdFamilyCode types.TThostFtdcInstrumentIDType
	// 腿1比例系数
	Leg1PropFactor types.TThostFtdcCommonIntType
	// 腿2比例系数
	Leg2PropFactor types.TThostFtdcCommonIntType
	// 商品群号
	CommodityGroupID types.TThostFtdcCommodityGroupIDType
	// 商品群名称
	CommodityGroupName types.TThostFtdcInstrumentNameType
}

func (d CThostFtdcRULEInterParameterField) Type() string { return "CThostFtdcRULEInterParameterField" }

func (d CThostFtdcRULEInterParameterField) String() string {
	var builder strings.Builder
	builder.Grow(268)

	builder.WriteString("CThostFtdcRULEInterParameterField{")
	fmt.Fprintf(&builder, "TradingDay=%+v", d.TradingDay)
	fmt.Fprintf(&builder, ", ExchangeID=%+v", d.ExchangeID)
	fmt.Fprintf(&builder, ", SpreadId=%+v", d.SpreadId)
	fmt.Fprintf(&builder, ", InterRate=%+v", d.InterRate)
	fmt.Fprintf(&builder, ", Leg1ProdFamilyCode=%+v", d.Leg1ProdFamilyCode)
	fmt.Fprintf(&builder, ", Leg2ProdFamilyCode=%+v", d.Leg2ProdFamilyCode)
	fmt.Fprintf(&builder, ", Leg1PropFactor=%+v", d.Leg1PropFactor)
	fmt.Fprintf(&builder, ", Leg2PropFactor=%+v", d.Leg2PropFactor)
	fmt.Fprintf(&builder, ", CommodityGroupID=%+v", d.CommodityGroupID)
	fmt.Fprintf(&builder, ", CommodityGroupName=%+v", d.CommodityGroupName)

	builder.WriteByte('}')

	return builder.String()
}

// RULE合约保证金参数查询
type CThostFtdcQryRULEInstrParameterField struct {
	// 交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	// 合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
}

func (d CThostFtdcQryRULEInstrParameterField) Type() string {
	return "CThostFtdcQryRULEInstrParameterField"
}

func (d CThostFtdcQryRULEInstrParameterField) String() string {
	var builder strings.Builder
	builder.Grow(78)

	builder.WriteString("CThostFtdcQryRULEInstrParameterField{")
	fmt.Fprintf(&builder, "ExchangeID=%+v", d.ExchangeID)
	fmt.Fprintf(&builder, ", InstrumentID=%+v", d.InstrumentID)

	builder.WriteByte('}')

	return builder.String()
}

// RULE品种内对锁仓折扣参数查询
type CThostFtdcQryRULEIntraParameterField struct {
	// 交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	// 品种代码
	ProdFamilyCode types.TThostFtdcInstrumentIDType
}

func (d CThostFtdcQryRULEIntraParameterField) Type() string {
	return "CThostFtdcQryRULEIntraParameterField"
}

func (d CThostFtdcQryRULEIntraParameterField) String() string {
	var builder strings.Builder
	builder.Grow(80)

	builder.WriteString("CThostFtdcQryRULEIntraParameterField{")
	fmt.Fprintf(&builder, "ExchangeID=%+v", d.ExchangeID)
	fmt.Fprintf(&builder, ", ProdFamilyCode=%+v", d.ProdFamilyCode)

	builder.WriteByte('}')

	return builder.String()
}

// RULE跨品种抵扣参数查询
type CThostFtdcQryRULEInterParameterField struct {
	// 交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	// 第一腿构成品种
	Leg1ProdFamilyCode types.TThostFtdcInstrumentIDType
	// 第二腿构成品种
	Leg2ProdFamilyCode types.TThostFtdcInstrumentIDType
	// 商品群号
	CommodityGroupID types.TThostFtdcCommodityGroupIDType
}

func (d CThostFtdcQryRULEInterParameterField) Type() string {
	return "CThostFtdcQryRULEInterParameterField"
}

func (d CThostFtdcQryRULEInterParameterField) String() string {
	var builder strings.Builder
	builder.Grow(138)

	builder.WriteString("CThostFtdcQryRULEInterParameterField{")
	fmt.Fprintf(&builder, "ExchangeID=%+v", d.ExchangeID)
	fmt.Fprintf(&builder, ", Leg1ProdFamilyCode=%+v", d.Leg1ProdFamilyCode)
	fmt.Fprintf(&builder, ", Leg2ProdFamilyCode=%+v", d.Leg2ProdFamilyCode)
	fmt.Fprintf(&builder, ", CommodityGroupID=%+v", d.CommodityGroupID)

	builder.WriteByte('}')

	return builder.String()
}

// 投资者产品RULE保证金
type CThostFtdcInvestorProdRULEMarginField struct {
	// 交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	// 品种代码
	ProdFamilyCode types.TThostFtdcInstrumentIDType
	// 合约类型
	InstrumentClass types.TThostFtdcInstrumentClassType
	// 商品群号
	CommodityGroupID types.TThostFtdcCommodityGroupIDType
	// 买标准持仓
	BStdPosition types.TThostFtdcStdPositionType
	// 卖标准持仓
	SStdPosition types.TThostFtdcStdPositionType
	// 买标准开仓冻结
	BStdOpenFrozen types.TThostFtdcStdPositionType
	// 卖标准开仓冻结
	SStdOpenFrozen types.TThostFtdcStdPositionType
	// 买标准平仓冻结
	BStdCloseFrozen types.TThostFtdcStdPositionType
	// 卖标准平仓冻结
	SStdCloseFrozen types.TThostFtdcStdPositionType
	// 品种内对冲标准持仓
	IntraProdStdPosition types.TThostFtdcStdPositionType
	// 品种内单腿标准持仓
	NetStdPosition types.TThostFtdcStdPositionType
	// 品种间对冲标准持仓
	InterProdStdPosition types.TThostFtdcStdPositionType
	// 单腿标准持仓
	SingleStdPosition types.TThostFtdcStdPositionType
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
	// 附加冻结保证金
	AddOnFrozenMargin types.TThostFtdcMoneyType
	// 开仓冻结保证金
	OpenFrozenMargin types.TThostFtdcMoneyType
	// 平仓冻结保证金
	CloseFrozenMargin types.TThostFtdcMoneyType
	// 品种保证金
	Margin types.TThostFtdcMoneyType
	// 冻结保证金
	FrozenMargin types.TThostFtdcMoneyType
}

func (d CThostFtdcInvestorProdRULEMarginField) Type() string {
	return "CThostFtdcInvestorProdRULEMarginField"
}

func (d CThostFtdcInvestorProdRULEMarginField) String() string {
	var builder strings.Builder
	builder.Grow(677)

	builder.WriteString("CThostFtdcInvestorProdRULEMarginField{")
	fmt.Fprintf(&builder, "ExchangeID=%+v", d.ExchangeID)
	fmt.Fprintf(&builder, ", BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", InvestorID=%+v", d.InvestorID)
	fmt.Fprintf(&builder, ", ProdFamilyCode=%+v", d.ProdFamilyCode)
	fmt.Fprintf(&builder, ", InstrumentClass=%+v", d.InstrumentClass)
	fmt.Fprintf(&builder, ", CommodityGroupID=%+v", d.CommodityGroupID)
	fmt.Fprintf(&builder, ", BStdPosition=%+v", d.BStdPosition)
	fmt.Fprintf(&builder, ", SStdPosition=%+v", d.SStdPosition)
	fmt.Fprintf(&builder, ", BStdOpenFrozen=%+v", d.BStdOpenFrozen)
	fmt.Fprintf(&builder, ", SStdOpenFrozen=%+v", d.SStdOpenFrozen)
	fmt.Fprintf(&builder, ", BStdCloseFrozen=%+v", d.BStdCloseFrozen)
	fmt.Fprintf(&builder, ", SStdCloseFrozen=%+v", d.SStdCloseFrozen)
	fmt.Fprintf(&builder, ", IntraProdStdPosition=%+v", d.IntraProdStdPosition)
	fmt.Fprintf(&builder, ", NetStdPosition=%+v", d.NetStdPosition)
	fmt.Fprintf(&builder, ", InterProdStdPosition=%+v", d.InterProdStdPosition)
	fmt.Fprintf(&builder, ", SingleStdPosition=%+v", d.SingleStdPosition)
	fmt.Fprintf(&builder, ", IntraProdMargin=%+v", d.IntraProdMargin)
	fmt.Fprintf(&builder, ", InterProdMargin=%+v", d.InterProdMargin)
	fmt.Fprintf(&builder, ", SingleMargin=%+v", d.SingleMargin)
	fmt.Fprintf(&builder, ", NonCombMargin=%+v", d.NonCombMargin)
	fmt.Fprintf(&builder, ", AddOnMargin=%+v", d.AddOnMargin)
	fmt.Fprintf(&builder, ", ExchMargin=%+v", d.ExchMargin)
	fmt.Fprintf(&builder, ", AddOnFrozenMargin=%+v", d.AddOnFrozenMargin)
	fmt.Fprintf(&builder, ", OpenFrozenMargin=%+v", d.OpenFrozenMargin)
	fmt.Fprintf(&builder, ", CloseFrozenMargin=%+v", d.CloseFrozenMargin)
	fmt.Fprintf(&builder, ", Margin=%+v", d.Margin)
	fmt.Fprintf(&builder, ", FrozenMargin=%+v", d.FrozenMargin)

	builder.WriteByte('}')

	return builder.String()
}

// 投资者产品RULE保证金查询
type CThostFtdcQryInvestorProdRULEMarginField struct {
	// 交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	// 品种代码
	ProdFamilyCode types.TThostFtdcInstrumentIDType
	// 商品群号
	CommodityGroupID types.TThostFtdcCommodityGroupIDType
}

func (d CThostFtdcQryInvestorProdRULEMarginField) Type() string {
	return "CThostFtdcQryInvestorProdRULEMarginField"
}

func (d CThostFtdcQryInvestorProdRULEMarginField) String() string {
	var builder strings.Builder
	builder.Grow(148)

	builder.WriteString("CThostFtdcQryInvestorProdRULEMarginField{")
	fmt.Fprintf(&builder, "ExchangeID=%+v", d.ExchangeID)
	fmt.Fprintf(&builder, ", BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", InvestorID=%+v", d.InvestorID)
	fmt.Fprintf(&builder, ", ProdFamilyCode=%+v", d.ProdFamilyCode)
	fmt.Fprintf(&builder, ", CommodityGroupID=%+v", d.CommodityGroupID)

	builder.WriteByte('}')

	return builder.String()
}

// 风险结算追平SPBM组合保证金套餐
type CThostFtdcSyncDeltaSPBMPortfDefinitionField struct {
	// 交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	// 组合保证金套餐代码
	PortfolioDefID types.TThostFtdcPortfolioDefIDType
	// 品种代码
	ProdFamilyCode types.TThostFtdcInstrumentIDType
	// 是否启用SPBM
	IsSPBM types.TThostFtdcBoolType
	// 操作标志
	ActionDirection types.TThostFtdcActionDirectionType
	// 追平序号
	SyncDeltaSequenceNo types.TThostFtdcSequenceNoType
}

func (d CThostFtdcSyncDeltaSPBMPortfDefinitionField) Type() string {
	return "CThostFtdcSyncDeltaSPBMPortfDefinitionField"
}

func (d CThostFtdcSyncDeltaSPBMPortfDefinitionField) String() string {
	var builder strings.Builder
	builder.Grow(181)

	builder.WriteString("CThostFtdcSyncDeltaSPBMPortfDefinitionField{")
	fmt.Fprintf(&builder, "ExchangeID=%+v", d.ExchangeID)
	fmt.Fprintf(&builder, ", PortfolioDefID=%+v", d.PortfolioDefID)
	fmt.Fprintf(&builder, ", ProdFamilyCode=%+v", d.ProdFamilyCode)
	fmt.Fprintf(&builder, ", IsSPBM=%+v", d.IsSPBM)
	fmt.Fprintf(&builder, ", ActionDirection=%+v", d.ActionDirection)
	fmt.Fprintf(&builder, ", SyncDeltaSequenceNo=%+v", d.SyncDeltaSequenceNo)

	builder.WriteByte('}')

	return builder.String()
}

// 风险结算追平投资者SPBM套餐选择
type CThostFtdcSyncDeltaSPBMInvstPortfDefField struct {
	// 交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	// 组合保证金套餐代码
	PortfolioDefID types.TThostFtdcPortfolioDefIDType
	// 操作标志
	ActionDirection types.TThostFtdcActionDirectionType
	// 追平序号
	SyncDeltaSequenceNo types.TThostFtdcSequenceNoType
}

func (d CThostFtdcSyncDeltaSPBMInvstPortfDefField) Type() string {
	return "CThostFtdcSyncDeltaSPBMInvstPortfDefField"
}

func (d CThostFtdcSyncDeltaSPBMInvstPortfDefField) String() string {
	var builder strings.Builder
	builder.Grow(177)

	builder.WriteString("CThostFtdcSyncDeltaSPBMInvstPortfDefField{")
	fmt.Fprintf(&builder, "ExchangeID=%+v", d.ExchangeID)
	fmt.Fprintf(&builder, ", BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", InvestorID=%+v", d.InvestorID)
	fmt.Fprintf(&builder, ", PortfolioDefID=%+v", d.PortfolioDefID)
	fmt.Fprintf(&builder, ", ActionDirection=%+v", d.ActionDirection)
	fmt.Fprintf(&builder, ", SyncDeltaSequenceNo=%+v", d.SyncDeltaSequenceNo)

	builder.WriteByte('}')

	return builder.String()
}

// 风险结算追平SPBM期货合约保证金参数
type CThostFtdcSyncDeltaSPBMFutureParameterField struct {
	// 交易日
	TradingDay types.TThostFtdcDateType
	// 交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	// 合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
	// 品种代码
	ProdFamilyCode types.TThostFtdcInstrumentIDType
	// 期货合约因子
	Cvf types.TThostFtdcVolumeMultipleType
	// 阶段标识
	TimeRange types.TThostFtdcTimeRangeType
	// 品种保证金标准
	MarginRate types.TThostFtdcRatioType
	// 期货合约内部对锁仓费率折扣比例
	LockRateX types.TThostFtdcRatioType
	// 提高保证金标准
	AddOnRate types.TThostFtdcRatioType
	// 昨结算价
	PreSettlementPrice types.TThostFtdcPriceType
	// 期货合约内部对锁仓附加费率折扣比例
	AddOnLockRateX2 types.TThostFtdcRatioType
	// 操作标志
	ActionDirection types.TThostFtdcActionDirectionType
	// 追平序号
	SyncDeltaSequenceNo types.TThostFtdcSequenceNoType
}

func (d CThostFtdcSyncDeltaSPBMFutureParameterField) Type() string {
	return "CThostFtdcSyncDeltaSPBMFutureParameterField"
}

func (d CThostFtdcSyncDeltaSPBMFutureParameterField) String() string {
	var builder strings.Builder
	builder.Grow(326)

	builder.WriteString("CThostFtdcSyncDeltaSPBMFutureParameterField{")
	fmt.Fprintf(&builder, "TradingDay=%+v", d.TradingDay)
	fmt.Fprintf(&builder, ", ExchangeID=%+v", d.ExchangeID)
	fmt.Fprintf(&builder, ", InstrumentID=%+v", d.InstrumentID)
	fmt.Fprintf(&builder, ", ProdFamilyCode=%+v", d.ProdFamilyCode)
	fmt.Fprintf(&builder, ", Cvf=%+v", d.Cvf)
	fmt.Fprintf(&builder, ", TimeRange=%+v", d.TimeRange)
	fmt.Fprintf(&builder, ", MarginRate=%+v", d.MarginRate)
	fmt.Fprintf(&builder, ", LockRateX=%+v", d.LockRateX)
	fmt.Fprintf(&builder, ", AddOnRate=%+v", d.AddOnRate)
	fmt.Fprintf(&builder, ", PreSettlementPrice=%+v", d.PreSettlementPrice)
	fmt.Fprintf(&builder, ", AddOnLockRateX2=%+v", d.AddOnLockRateX2)
	fmt.Fprintf(&builder, ", ActionDirection=%+v", d.ActionDirection)
	fmt.Fprintf(&builder, ", SyncDeltaSequenceNo=%+v", d.SyncDeltaSequenceNo)

	builder.WriteByte('}')

	return builder.String()
}

// 风险结算追平SPBM期权合约保证金参数
type CThostFtdcSyncDeltaSPBMOptionParameterField struct {
	// 交易日
	TradingDay types.TThostFtdcDateType
	// 交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	// 合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
	// 品种代码
	ProdFamilyCode types.TThostFtdcInstrumentIDType
	// 期权合约因子
	Cvf types.TThostFtdcVolumeMultipleType
	// 期权冲抵价格
	DownPrice types.TThostFtdcPriceType
	// Delta值
	Delta types.TThostFtdcDeltaType
	// 卖方期权风险转换最低值
	SlimiDelta types.TThostFtdcDeltaType
	// 昨结算价
	PreSettlementPrice types.TThostFtdcPriceType
	// 操作标志
	ActionDirection types.TThostFtdcActionDirectionType
	// 追平序号
	SyncDeltaSequenceNo types.TThostFtdcSequenceNoType
}

func (d CThostFtdcSyncDeltaSPBMOptionParameterField) Type() string {
	return "CThostFtdcSyncDeltaSPBMOptionParameterField"
}

func (d CThostFtdcSyncDeltaSPBMOptionParameterField) String() string {
	var builder strings.Builder
	builder.Grow(278)

	builder.WriteString("CThostFtdcSyncDeltaSPBMOptionParameterField{")
	fmt.Fprintf(&builder, "TradingDay=%+v", d.TradingDay)
	fmt.Fprintf(&builder, ", ExchangeID=%+v", d.ExchangeID)
	fmt.Fprintf(&builder, ", InstrumentID=%+v", d.InstrumentID)
	fmt.Fprintf(&builder, ", ProdFamilyCode=%+v", d.ProdFamilyCode)
	fmt.Fprintf(&builder, ", Cvf=%+v", d.Cvf)
	fmt.Fprintf(&builder, ", DownPrice=%+v", d.DownPrice)
	fmt.Fprintf(&builder, ", Delta=%+v", d.Delta)
	fmt.Fprintf(&builder, ", SlimiDelta=%+v", d.SlimiDelta)
	fmt.Fprintf(&builder, ", PreSettlementPrice=%+v", d.PreSettlementPrice)
	fmt.Fprintf(&builder, ", ActionDirection=%+v", d.ActionDirection)
	fmt.Fprintf(&builder, ", SyncDeltaSequenceNo=%+v", d.SyncDeltaSequenceNo)

	builder.WriteByte('}')

	return builder.String()
}

// 风险结算追平SPBM品种内对锁仓折扣参数
type CThostFtdcSyncDeltaSPBMIntraParameterField struct {
	// 交易日
	TradingDay types.TThostFtdcDateType
	// 交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	// 品种代码
	ProdFamilyCode types.TThostFtdcInstrumentIDType
	// 品种内合约间对锁仓费率折扣比例
	IntraRateY types.TThostFtdcRatioType
	// 品种内合约间对锁仓附加费率折扣比例
	AddOnIntraRateY2 types.TThostFtdcRatioType
	// 操作标志
	ActionDirection types.TThostFtdcActionDirectionType
	// 追平序号
	SyncDeltaSequenceNo types.TThostFtdcSequenceNoType
}

func (d CThostFtdcSyncDeltaSPBMIntraParameterField) Type() string {
	return "CThostFtdcSyncDeltaSPBMIntraParameterField"
}

func (d CThostFtdcSyncDeltaSPBMIntraParameterField) String() string {
	var builder strings.Builder
	builder.Grow(206)

	builder.WriteString("CThostFtdcSyncDeltaSPBMIntraParameterField{")
	fmt.Fprintf(&builder, "TradingDay=%+v", d.TradingDay)
	fmt.Fprintf(&builder, ", ExchangeID=%+v", d.ExchangeID)
	fmt.Fprintf(&builder, ", ProdFamilyCode=%+v", d.ProdFamilyCode)
	fmt.Fprintf(&builder, ", IntraRateY=%+v", d.IntraRateY)
	fmt.Fprintf(&builder, ", AddOnIntraRateY2=%+v", d.AddOnIntraRateY2)
	fmt.Fprintf(&builder, ", ActionDirection=%+v", d.ActionDirection)
	fmt.Fprintf(&builder, ", SyncDeltaSequenceNo=%+v", d.SyncDeltaSequenceNo)

	builder.WriteByte('}')

	return builder.String()
}

// 风险结算追平SPBM跨品种抵扣参数
type CThostFtdcSyncDeltaSPBMInterParameterField struct {
	// 交易日
	TradingDay types.TThostFtdcDateType
	// 交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	// 优先级
	SpreadId types.TThostFtdcSpreadIdType
	// 品种间对锁仓费率折扣比例
	InterRateZ types.TThostFtdcRatioType
	// 第一腿构成品种
	Leg1ProdFamilyCode types.TThostFtdcInstrumentIDType
	// 第二腿构成品种
	Leg2ProdFamilyCode types.TThostFtdcInstrumentIDType
	// 操作标志
	ActionDirection types.TThostFtdcActionDirectionType
	// 追平序号
	SyncDeltaSequenceNo types.TThostFtdcSequenceNoType
}

func (d CThostFtdcSyncDeltaSPBMInterParameterField) Type() string {
	return "CThostFtdcSyncDeltaSPBMInterParameterField"
}

func (d CThostFtdcSyncDeltaSPBMInterParameterField) String() string {
	var builder strings.Builder
	builder.Grow(230)

	builder.WriteString("CThostFtdcSyncDeltaSPBMInterParameterField{")
	fmt.Fprintf(&builder, "TradingDay=%+v", d.TradingDay)
	fmt.Fprintf(&builder, ", ExchangeID=%+v", d.ExchangeID)
	fmt.Fprintf(&builder, ", SpreadId=%+v", d.SpreadId)
	fmt.Fprintf(&builder, ", InterRateZ=%+v", d.InterRateZ)
	fmt.Fprintf(&builder, ", Leg1ProdFamilyCode=%+v", d.Leg1ProdFamilyCode)
	fmt.Fprintf(&builder, ", Leg2ProdFamilyCode=%+v", d.Leg2ProdFamilyCode)
	fmt.Fprintf(&builder, ", ActionDirection=%+v", d.ActionDirection)
	fmt.Fprintf(&builder, ", SyncDeltaSequenceNo=%+v", d.SyncDeltaSequenceNo)

	builder.WriteByte('}')

	return builder.String()
}

// 风险结算追平SPBM附加跨品种抵扣参数
type CThostFtdcSyncDeltaSPBMAddOnInterParamField struct {
	// 交易日
	TradingDay types.TThostFtdcDateType
	// 交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	// 优先级
	SpreadId types.TThostFtdcSpreadIdType
	// 品种间对锁仓附加费率折扣比例
	AddOnInterRateZ2 types.TThostFtdcRatioType
	// 第一腿构成品种
	Leg1ProdFamilyCode types.TThostFtdcInstrumentIDType
	// 第二腿构成品种
	Leg2ProdFamilyCode types.TThostFtdcInstrumentIDType
	// 操作标志
	ActionDirection types.TThostFtdcActionDirectionType
	// 追平序号
	SyncDeltaSequenceNo types.TThostFtdcSequenceNoType
}

func (d CThostFtdcSyncDeltaSPBMAddOnInterParamField) Type() string {
	return "CThostFtdcSyncDeltaSPBMAddOnInterParamField"
}

func (d CThostFtdcSyncDeltaSPBMAddOnInterParamField) String() string {
	var builder strings.Builder
	builder.Grow(237)

	builder.WriteString("CThostFtdcSyncDeltaSPBMAddOnInterParamField{")
	fmt.Fprintf(&builder, "TradingDay=%+v", d.TradingDay)
	fmt.Fprintf(&builder, ", ExchangeID=%+v", d.ExchangeID)
	fmt.Fprintf(&builder, ", SpreadId=%+v", d.SpreadId)
	fmt.Fprintf(&builder, ", AddOnInterRateZ2=%+v", d.AddOnInterRateZ2)
	fmt.Fprintf(&builder, ", Leg1ProdFamilyCode=%+v", d.Leg1ProdFamilyCode)
	fmt.Fprintf(&builder, ", Leg2ProdFamilyCode=%+v", d.Leg2ProdFamilyCode)
	fmt.Fprintf(&builder, ", ActionDirection=%+v", d.ActionDirection)
	fmt.Fprintf(&builder, ", SyncDeltaSequenceNo=%+v", d.SyncDeltaSequenceNo)

	builder.WriteByte('}')

	return builder.String()
}

// 风险结算追平SPMM合约参数
type CThostFtdcSyncDeltaSPMMInstParamField struct {
	// 交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	// 合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
	// SPMM合约保证金算法
	InstMarginCalID types.TThostFtdcInstMarginCalIDType
	// 商品组代码
	CommodityID types.TThostFtdcSPMMProductIDType
	// 商品群代码
	CommodityGroupID types.TThostFtdcSPMMProductIDType
	// 操作标志
	ActionDirection types.TThostFtdcActionDirectionType
	// 追平序号
	SyncDeltaSequenceNo types.TThostFtdcSequenceNoType
}

func (d CThostFtdcSyncDeltaSPMMInstParamField) Type() string {
	return "CThostFtdcSyncDeltaSPMMInstParamField"
}

func (d CThostFtdcSyncDeltaSPMMInstParamField) String() string {
	var builder strings.Builder
	builder.Grow(205)

	builder.WriteString("CThostFtdcSyncDeltaSPMMInstParamField{")
	fmt.Fprintf(&builder, "ExchangeID=%+v", d.ExchangeID)
	fmt.Fprintf(&builder, ", InstrumentID=%+v", d.InstrumentID)
	fmt.Fprintf(&builder, ", InstMarginCalID=%+v", d.InstMarginCalID)
	fmt.Fprintf(&builder, ", CommodityID=%+v", d.CommodityID)
	fmt.Fprintf(&builder, ", CommodityGroupID=%+v", d.CommodityGroupID)
	fmt.Fprintf(&builder, ", ActionDirection=%+v", d.ActionDirection)
	fmt.Fprintf(&builder, ", SyncDeltaSequenceNo=%+v", d.SyncDeltaSequenceNo)

	builder.WriteByte('}')

	return builder.String()
}

// 风险结算追平SPMM产品相关参数
type CThostFtdcSyncDeltaSPMMProductParamField struct {
	// 交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	// 产品代码
	ProductID types.TThostFtdcSPMMProductIDType
	// 商品组代码
	CommodityID types.TThostFtdcSPMMProductIDType
	// 商品群代码
	CommodityGroupID types.TThostFtdcSPMMProductIDType
	// 操作标志
	ActionDirection types.TThostFtdcActionDirectionType
	// 追平序号
	SyncDeltaSequenceNo types.TThostFtdcSequenceNoType
}

func (d CThostFtdcSyncDeltaSPMMProductParamField) Type() string {
	return "CThostFtdcSyncDeltaSPMMProductParamField"
}

func (d CThostFtdcSyncDeltaSPMMProductParamField) String() string {
	var builder strings.Builder
	builder.Grow(180)

	builder.WriteString("CThostFtdcSyncDeltaSPMMProductParamField{")
	fmt.Fprintf(&builder, "ExchangeID=%+v", d.ExchangeID)
	fmt.Fprintf(&builder, ", ProductID=%+v", d.ProductID)
	fmt.Fprintf(&builder, ", CommodityID=%+v", d.CommodityID)
	fmt.Fprintf(&builder, ", CommodityGroupID=%+v", d.CommodityGroupID)
	fmt.Fprintf(&builder, ", ActionDirection=%+v", d.ActionDirection)
	fmt.Fprintf(&builder, ", SyncDeltaSequenceNo=%+v", d.SyncDeltaSequenceNo)

	builder.WriteByte('}')

	return builder.String()
}

// 风险结算追平投资者SPMM模板选择
type CThostFtdcSyncDeltaInvestorSPMMModelField struct {
	// 交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	// SPMM模板ID
	SPMMModelID types.TThostFtdcSPMMModelIDType
	// 操作标志
	ActionDirection types.TThostFtdcActionDirectionType
	// 追平序号
	SyncDeltaSequenceNo types.TThostFtdcSequenceNoType
}

func (d CThostFtdcSyncDeltaInvestorSPMMModelField) Type() string {
	return "CThostFtdcSyncDeltaInvestorSPMMModelField"
}

func (d CThostFtdcSyncDeltaInvestorSPMMModelField) String() string {
	var builder strings.Builder
	builder.Grow(174)

	builder.WriteString("CThostFtdcSyncDeltaInvestorSPMMModelField{")
	fmt.Fprintf(&builder, "ExchangeID=%+v", d.ExchangeID)
	fmt.Fprintf(&builder, ", BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", InvestorID=%+v", d.InvestorID)
	fmt.Fprintf(&builder, ", SPMMModelID=%+v", d.SPMMModelID)
	fmt.Fprintf(&builder, ", ActionDirection=%+v", d.ActionDirection)
	fmt.Fprintf(&builder, ", SyncDeltaSequenceNo=%+v", d.SyncDeltaSequenceNo)

	builder.WriteByte('}')

	return builder.String()
}

// 风险结算追平SPMM模板参数设置
type CThostFtdcSyncDeltaSPMMModelParamField struct {
	// 交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	// SPMM模板ID
	SPMMModelID types.TThostFtdcSPMMModelIDType
	// 商品群代码
	CommodityGroupID types.TThostFtdcSPMMProductIDType
	// SPMM品种内跨期优惠系数
	IntraCommodityRate types.TThostFtdcSPMMDiscountRatioType
	// SPMM品种间优惠系数
	InterCommodityRate types.TThostFtdcSPMMDiscountRatioType
	// SPMM期权优惠系数
	OptionDiscountRate types.TThostFtdcSPMMDiscountRatioType
	// 商品群最小保证金比例
	MiniMarginRatio types.TThostFtdcSPMMDiscountRatioType
	// 操作标志
	ActionDirection types.TThostFtdcActionDirectionType
	// 追平序号
	SyncDeltaSequenceNo types.TThostFtdcSequenceNoType
}

func (d CThostFtdcSyncDeltaSPMMModelParamField) Type() string {
	return "CThostFtdcSyncDeltaSPMMModelParamField"
}

func (d CThostFtdcSyncDeltaSPMMModelParamField) String() string {
	var builder strings.Builder
	builder.Grow(268)

	builder.WriteString("CThostFtdcSyncDeltaSPMMModelParamField{")
	fmt.Fprintf(&builder, "ExchangeID=%+v", d.ExchangeID)
	fmt.Fprintf(&builder, ", SPMMModelID=%+v", d.SPMMModelID)
	fmt.Fprintf(&builder, ", CommodityGroupID=%+v", d.CommodityGroupID)
	fmt.Fprintf(&builder, ", IntraCommodityRate=%+v", d.IntraCommodityRate)
	fmt.Fprintf(&builder, ", InterCommodityRate=%+v", d.InterCommodityRate)
	fmt.Fprintf(&builder, ", OptionDiscountRate=%+v", d.OptionDiscountRate)
	fmt.Fprintf(&builder, ", MiniMarginRatio=%+v", d.MiniMarginRatio)
	fmt.Fprintf(&builder, ", ActionDirection=%+v", d.ActionDirection)
	fmt.Fprintf(&builder, ", SyncDeltaSequenceNo=%+v", d.SyncDeltaSequenceNo)

	builder.WriteByte('}')

	return builder.String()
}

// 风险结算追平RCAMS产品组合信息
type CThostFtdcSyncDeltaRCAMSCombProdInfoField struct {
	// 交易日
	TradingDay types.TThostFtdcDateType
	// 交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	// 产品代码
	ProductID types.TThostFtdcProductIDType
	// 商品组代码
	CombProductID types.TThostFtdcProductIDType
	// 商品群代码
	ProductGroupID types.TThostFtdcProductIDType
	// 操作标志
	ActionDirection types.TThostFtdcActionDirectionType
	// 追平序号
	SyncDeltaSequenceNo types.TThostFtdcSequenceNoType
}

func (d CThostFtdcSyncDeltaRCAMSCombProdInfoField) Type() string {
	return "CThostFtdcSyncDeltaRCAMSCombProdInfoField"
}

func (d CThostFtdcSyncDeltaRCAMSCombProdInfoField) String() string {
	var builder strings.Builder
	builder.Grow(201)

	builder.WriteString("CThostFtdcSyncDeltaRCAMSCombProdInfoField{")
	fmt.Fprintf(&builder, "TradingDay=%+v", d.TradingDay)
	fmt.Fprintf(&builder, ", ExchangeID=%+v", d.ExchangeID)
	fmt.Fprintf(&builder, ", ProductID=%+v", d.ProductID)
	fmt.Fprintf(&builder, ", CombProductID=%+v", d.CombProductID)
	fmt.Fprintf(&builder, ", ProductGroupID=%+v", d.ProductGroupID)
	fmt.Fprintf(&builder, ", ActionDirection=%+v", d.ActionDirection)
	fmt.Fprintf(&builder, ", SyncDeltaSequenceNo=%+v", d.SyncDeltaSequenceNo)

	builder.WriteByte('}')

	return builder.String()
}

// 风险结算追平RCAMS同合约风险对冲参数
type CThostFtdcSyncDeltaRCAMSInstrParameterField struct {
	// 交易日
	TradingDay types.TThostFtdcDateType
	// 交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	// 产品代码
	ProductID types.TThostFtdcProductIDType
	// 同合约风险对冲比率
	HedgeRate types.TThostFtdcHedgeRateType
	// 操作标志
	ActionDirection types.TThostFtdcActionDirectionType
	// 追平序号
	SyncDeltaSequenceNo types.TThostFtdcSequenceNoType
}

func (d CThostFtdcSyncDeltaRCAMSInstrParameterField) Type() string {
	return "CThostFtdcSyncDeltaRCAMSInstrParameterField"
}

func (d CThostFtdcSyncDeltaRCAMSInstrParameterField) String() string {
	var builder strings.Builder
	builder.Grow(175)

	builder.WriteString("CThostFtdcSyncDeltaRCAMSInstrParameterField{")
	fmt.Fprintf(&builder, "TradingDay=%+v", d.TradingDay)
	fmt.Fprintf(&builder, ", ExchangeID=%+v", d.ExchangeID)
	fmt.Fprintf(&builder, ", ProductID=%+v", d.ProductID)
	fmt.Fprintf(&builder, ", HedgeRate=%+v", d.HedgeRate)
	fmt.Fprintf(&builder, ", ActionDirection=%+v", d.ActionDirection)
	fmt.Fprintf(&builder, ", SyncDeltaSequenceNo=%+v", d.SyncDeltaSequenceNo)

	builder.WriteByte('}')

	return builder.String()
}

// 风险结算追平RCAMS品种内风险对冲参数
type CThostFtdcSyncDeltaRCAMSIntraParameterField struct {
	// 交易日
	TradingDay types.TThostFtdcDateType
	// 交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	// 产品组合代码
	CombProductID types.TThostFtdcProductIDType
	// 品种内对冲比率
	HedgeRate types.TThostFtdcHedgeRateType
	// 操作标志
	ActionDirection types.TThostFtdcActionDirectionType
	// 追平序号
	SyncDeltaSequenceNo types.TThostFtdcSequenceNoType
}

func (d CThostFtdcSyncDeltaRCAMSIntraParameterField) Type() string {
	return "CThostFtdcSyncDeltaRCAMSIntraParameterField"
}

func (d CThostFtdcSyncDeltaRCAMSIntraParameterField) String() string {
	var builder strings.Builder
	builder.Grow(179)

	builder.WriteString("CThostFtdcSyncDeltaRCAMSIntraParameterField{")
	fmt.Fprintf(&builder, "TradingDay=%+v", d.TradingDay)
	fmt.Fprintf(&builder, ", ExchangeID=%+v", d.ExchangeID)
	fmt.Fprintf(&builder, ", CombProductID=%+v", d.CombProductID)
	fmt.Fprintf(&builder, ", HedgeRate=%+v", d.HedgeRate)
	fmt.Fprintf(&builder, ", ActionDirection=%+v", d.ActionDirection)
	fmt.Fprintf(&builder, ", SyncDeltaSequenceNo=%+v", d.SyncDeltaSequenceNo)

	builder.WriteByte('}')

	return builder.String()
}

// 风险结算追平RCAMS跨品种风险折抵参数
type CThostFtdcSyncDeltaRCAMSInterParameterField struct {
	// 交易日
	TradingDay types.TThostFtdcDateType
	// 交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	// 商品群代码
	ProductGroupID types.TThostFtdcProductIDType
	// 优先级
	Priority types.TThostFtdcRCAMSPriorityType
	// 折抵率
	CreditRate types.TThostFtdcHedgeRateType
	// 产品组合代码1
	CombProduct1 types.TThostFtdcProductIDType
	// 产品组合代码2
	CombProduct2 types.TThostFtdcProductIDType
	// 操作标志
	ActionDirection types.TThostFtdcActionDirectionType
	// 追平序号
	SyncDeltaSequenceNo types.TThostFtdcSequenceNoType
}

func (d CThostFtdcSyncDeltaRCAMSInterParameterField) Type() string {
	return "CThostFtdcSyncDeltaRCAMSInterParameterField"
}

func (d CThostFtdcSyncDeltaRCAMSInterParameterField) String() string {
	var builder strings.Builder
	builder.Grow(243)

	builder.WriteString("CThostFtdcSyncDeltaRCAMSInterParameterField{")
	fmt.Fprintf(&builder, "TradingDay=%+v", d.TradingDay)
	fmt.Fprintf(&builder, ", ExchangeID=%+v", d.ExchangeID)
	fmt.Fprintf(&builder, ", ProductGroupID=%+v", d.ProductGroupID)
	fmt.Fprintf(&builder, ", Priority=%+v", d.Priority)
	fmt.Fprintf(&builder, ", CreditRate=%+v", d.CreditRate)
	fmt.Fprintf(&builder, ", CombProduct1=%+v", d.CombProduct1)
	fmt.Fprintf(&builder, ", CombProduct2=%+v", d.CombProduct2)
	fmt.Fprintf(&builder, ", ActionDirection=%+v", d.ActionDirection)
	fmt.Fprintf(&builder, ", SyncDeltaSequenceNo=%+v", d.SyncDeltaSequenceNo)

	builder.WriteByte('}')

	return builder.String()
}

// 风险结算追平RCAMS空头期权风险调整参数
type CThostFtdcSyncDeltaRCAMSSOptAdjParamField struct {
	// 交易日
	TradingDay types.TThostFtdcDateType
	// 交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	// 产品组合代码
	CombProductID types.TThostFtdcProductIDType
	// 投套标志
	HedgeFlag types.TThostFtdcHedgeFlagType
	// 空头期权风险调整标准
	AdjustValue types.TThostFtdcAdjustValueType
	// 操作标志
	ActionDirection types.TThostFtdcActionDirectionType
	// 追平序号
	SyncDeltaSequenceNo types.TThostFtdcSequenceNoType
}

func (d CThostFtdcSyncDeltaRCAMSSOptAdjParamField) Type() string {
	return "CThostFtdcSyncDeltaRCAMSSOptAdjParamField"
}

func (d CThostFtdcSyncDeltaRCAMSSOptAdjParamField) String() string {
	var builder strings.Builder
	builder.Grow(198)

	builder.WriteString("CThostFtdcSyncDeltaRCAMSSOptAdjParamField{")
	fmt.Fprintf(&builder, "TradingDay=%+v", d.TradingDay)
	fmt.Fprintf(&builder, ", ExchangeID=%+v", d.ExchangeID)
	fmt.Fprintf(&builder, ", CombProductID=%+v", d.CombProductID)
	fmt.Fprintf(&builder, ", HedgeFlag=%+v", d.HedgeFlag)
	fmt.Fprintf(&builder, ", AdjustValue=%+v", d.AdjustValue)
	fmt.Fprintf(&builder, ", ActionDirection=%+v", d.ActionDirection)
	fmt.Fprintf(&builder, ", SyncDeltaSequenceNo=%+v", d.SyncDeltaSequenceNo)

	builder.WriteByte('}')

	return builder.String()
}

// 风险结算追平RCAMS策略组合规则明细
type CThostFtdcSyncDeltaRCAMSCombRuleDtlField struct {
	// 交易日
	TradingDay types.TThostFtdcDateType
	// 交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	// 策略产品
	ProdGroup types.TThostFtdcProductIDType
	// 策略id
	RuleId types.TThostFtdcRuleIdType
	// 优先级
	Priority types.TThostFtdcRCAMSPriorityType
	// 投套标志
	HedgeFlag types.TThostFtdcHedgeFlagType
	// 组合保证金标准
	CombMargin types.TThostFtdcMoneyType
	// 交易所组合合约代码
	ExchangeInstID types.TThostFtdcExchangeInstIDType
	// 单腿编号
	LegID types.TThostFtdcLegIDType
	// 单腿合约代码
	LegInstrumentID types.TThostFtdcInstrumentIDType
	// 买卖方向
	Direction types.TThostFtdcDirectionType
	// 单腿乘数
	LegMultiple types.TThostFtdcLegMultipleType
	// 操作标志
	ActionDirection types.TThostFtdcActionDirectionType
	// 追平序号
	SyncDeltaSequenceNo types.TThostFtdcSequenceNoType
}

func (d CThostFtdcSyncDeltaRCAMSCombRuleDtlField) Type() string {
	return "CThostFtdcSyncDeltaRCAMSCombRuleDtlField"
}

func (d CThostFtdcSyncDeltaRCAMSCombRuleDtlField) String() string {
	var builder strings.Builder
	builder.Grow(330)

	builder.WriteString("CThostFtdcSyncDeltaRCAMSCombRuleDtlField{")
	fmt.Fprintf(&builder, "TradingDay=%+v", d.TradingDay)
	fmt.Fprintf(&builder, ", ExchangeID=%+v", d.ExchangeID)
	fmt.Fprintf(&builder, ", ProdGroup=%+v", d.ProdGroup)
	fmt.Fprintf(&builder, ", RuleId=%+v", d.RuleId)
	fmt.Fprintf(&builder, ", Priority=%+v", d.Priority)
	fmt.Fprintf(&builder, ", HedgeFlag=%+v", d.HedgeFlag)
	fmt.Fprintf(&builder, ", CombMargin=%+v", d.CombMargin)
	fmt.Fprintf(&builder, ", ExchangeInstID=%+v", d.ExchangeInstID)
	fmt.Fprintf(&builder, ", LegID=%+v", d.LegID)
	fmt.Fprintf(&builder, ", LegInstrumentID=%+v", d.LegInstrumentID)
	fmt.Fprintf(&builder, ", Direction=%+v", d.Direction)
	fmt.Fprintf(&builder, ", LegMultiple=%+v", d.LegMultiple)
	fmt.Fprintf(&builder, ", ActionDirection=%+v", d.ActionDirection)
	fmt.Fprintf(&builder, ", SyncDeltaSequenceNo=%+v", d.SyncDeltaSequenceNo)

	builder.WriteByte('}')

	return builder.String()
}

// 风险结算追平RCAMS策略组合持仓
type CThostFtdcSyncDeltaRCAMSInvstCombPosField struct {
	// 交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	// 合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
	// 投套标志
	HedgeFlag types.TThostFtdcHedgeFlagType
	// 持仓多空方向
	PosiDirection types.TThostFtdcPosiDirectionType
	// 组合合约代码
	CombInstrumentID types.TThostFtdcInstrumentIDType
	// 单腿编号
	LegID types.TThostFtdcLegIDType
	// 交易所组合合约代码
	ExchangeInstID types.TThostFtdcExchangeInstIDType
	// 持仓量
	TotalAmt types.TThostFtdcVolumeType
	// 交易所保证金
	ExchMargin types.TThostFtdcMoneyType
	// 投资者保证金
	Margin types.TThostFtdcMoneyType
	// 操作标志
	ActionDirection types.TThostFtdcActionDirectionType
	// 追平序号
	SyncDeltaSequenceNo types.TThostFtdcSequenceNoType
}

func (d CThostFtdcSyncDeltaRCAMSInvstCombPosField) Type() string {
	return "CThostFtdcSyncDeltaRCAMSInvstCombPosField"
}

func (d CThostFtdcSyncDeltaRCAMSInvstCombPosField) String() string {
	var builder strings.Builder
	builder.Grow(336)

	builder.WriteString("CThostFtdcSyncDeltaRCAMSInvstCombPosField{")
	fmt.Fprintf(&builder, "ExchangeID=%+v", d.ExchangeID)
	fmt.Fprintf(&builder, ", BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", InvestorID=%+v", d.InvestorID)
	fmt.Fprintf(&builder, ", InstrumentID=%+v", d.InstrumentID)
	fmt.Fprintf(&builder, ", HedgeFlag=%+v", d.HedgeFlag)
	fmt.Fprintf(&builder, ", PosiDirection=%+v", d.PosiDirection)
	fmt.Fprintf(&builder, ", CombInstrumentID=%+v", d.CombInstrumentID)
	fmt.Fprintf(&builder, ", LegID=%+v", d.LegID)
	fmt.Fprintf(&builder, ", ExchangeInstID=%+v", d.ExchangeInstID)
	fmt.Fprintf(&builder, ", TotalAmt=%+v", d.TotalAmt)
	fmt.Fprintf(&builder, ", ExchMargin=%+v", d.ExchMargin)
	fmt.Fprintf(&builder, ", Margin=%+v", d.Margin)
	fmt.Fprintf(&builder, ", ActionDirection=%+v", d.ActionDirection)
	fmt.Fprintf(&builder, ", SyncDeltaSequenceNo=%+v", d.SyncDeltaSequenceNo)

	builder.WriteByte('}')

	return builder.String()
}

// 风险结算追平RULE合约保证金参数
type CThostFtdcSyncDeltaRULEInstrParameterField struct {
	// 交易日
	TradingDay types.TThostFtdcDateType
	// 交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	// 合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
	// 合约类型
	InstrumentClass types.TThostFtdcInstrumentClassType
	// 标准合约
	StdInstrumentID types.TThostFtdcInstrumentIDType
	// 投机买折算系数
	BSpecRatio types.TThostFtdcRatioType
	// 投机卖折算系数
	SSpecRatio types.TThostFtdcRatioType
	// 套保买折算系数
	BHedgeRatio types.TThostFtdcRatioType
	// 套保卖折算系数
	SHedgeRatio types.TThostFtdcRatioType
	// 买附加风险保证金
	BAddOnMargin types.TThostFtdcMoneyType
	// 卖附加风险保证金
	SAddOnMargin types.TThostFtdcMoneyType
	// 商品群号
	CommodityGroupID types.TThostFtdcCommodityGroupIDType
	// 操作标志
	ActionDirection types.TThostFtdcActionDirectionType
	// 追平序号
	SyncDeltaSequenceNo types.TThostFtdcSequenceNoType
}

func (d CThostFtdcSyncDeltaRULEInstrParameterField) Type() string {
	return "CThostFtdcSyncDeltaRULEInstrParameterField"
}

func (d CThostFtdcSyncDeltaRULEInstrParameterField) String() string {
	var builder strings.Builder
	builder.Grow(360)

	builder.WriteString("CThostFtdcSyncDeltaRULEInstrParameterField{")
	fmt.Fprintf(&builder, "TradingDay=%+v", d.TradingDay)
	fmt.Fprintf(&builder, ", ExchangeID=%+v", d.ExchangeID)
	fmt.Fprintf(&builder, ", InstrumentID=%+v", d.InstrumentID)
	fmt.Fprintf(&builder, ", InstrumentClass=%+v", d.InstrumentClass)
	fmt.Fprintf(&builder, ", StdInstrumentID=%+v", d.StdInstrumentID)
	fmt.Fprintf(&builder, ", BSpecRatio=%+v", d.BSpecRatio)
	fmt.Fprintf(&builder, ", SSpecRatio=%+v", d.SSpecRatio)
	fmt.Fprintf(&builder, ", BHedgeRatio=%+v", d.BHedgeRatio)
	fmt.Fprintf(&builder, ", SHedgeRatio=%+v", d.SHedgeRatio)
	fmt.Fprintf(&builder, ", BAddOnMargin=%+v", d.BAddOnMargin)
	fmt.Fprintf(&builder, ", SAddOnMargin=%+v", d.SAddOnMargin)
	fmt.Fprintf(&builder, ", CommodityGroupID=%+v", d.CommodityGroupID)
	fmt.Fprintf(&builder, ", ActionDirection=%+v", d.ActionDirection)
	fmt.Fprintf(&builder, ", SyncDeltaSequenceNo=%+v", d.SyncDeltaSequenceNo)

	builder.WriteByte('}')

	return builder.String()
}

// 风险结算追平RULE品种内对锁仓折扣参数
type CThostFtdcSyncDeltaRULEIntraParameterField struct {
	// 交易日
	TradingDay types.TThostFtdcDateType
	// 交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	// 品种代码
	ProdFamilyCode types.TThostFtdcInstrumentIDType
	// 标准合约
	StdInstrumentID types.TThostFtdcInstrumentIDType
	// 标准合约保证金
	StdInstrMargin types.TThostFtdcMoneyType
	// 一般月份合约组合保证金系数
	UsualIntraRate types.TThostFtdcRatioType
	// 临近交割合约组合保证金系数
	DeliveryIntraRate types.TThostFtdcRatioType
	// 操作标志
	ActionDirection types.TThostFtdcActionDirectionType
	// 追平序号
	SyncDeltaSequenceNo types.TThostFtdcSequenceNoType
}

func (d CThostFtdcSyncDeltaRULEIntraParameterField) Type() string {
	return "CThostFtdcSyncDeltaRULEIntraParameterField"
}

func (d CThostFtdcSyncDeltaRULEIntraParameterField) String() string {
	var builder strings.Builder
	builder.Grow(260)

	builder.WriteString("CThostFtdcSyncDeltaRULEIntraParameterField{")
	fmt.Fprintf(&builder, "TradingDay=%+v", d.TradingDay)
	fmt.Fprintf(&builder, ", ExchangeID=%+v", d.ExchangeID)
	fmt.Fprintf(&builder, ", ProdFamilyCode=%+v", d.ProdFamilyCode)
	fmt.Fprintf(&builder, ", StdInstrumentID=%+v", d.StdInstrumentID)
	fmt.Fprintf(&builder, ", StdInstrMargin=%+v", d.StdInstrMargin)
	fmt.Fprintf(&builder, ", UsualIntraRate=%+v", d.UsualIntraRate)
	fmt.Fprintf(&builder, ", DeliveryIntraRate=%+v", d.DeliveryIntraRate)
	fmt.Fprintf(&builder, ", ActionDirection=%+v", d.ActionDirection)
	fmt.Fprintf(&builder, ", SyncDeltaSequenceNo=%+v", d.SyncDeltaSequenceNo)

	builder.WriteByte('}')

	return builder.String()
}

// 风险结算追平RULE跨品种抵扣参数
type CThostFtdcSyncDeltaRULEInterParameterField struct {
	// 交易日
	TradingDay types.TThostFtdcDateType
	// 交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	// 优先级
	SpreadId types.TThostFtdcSpreadIdType
	// 品种间对锁仓费率折扣比例
	InterRate types.TThostFtdcRatioType
	// 第一腿构成品种
	Leg1ProdFamilyCode types.TThostFtdcInstrumentIDType
	// 第二腿构成品种
	Leg2ProdFamilyCode types.TThostFtdcInstrumentIDType
	// 腿1比例系数
	Leg1PropFactor types.TThostFtdcCommonIntType
	// 腿2比例系数
	Leg2PropFactor types.TThostFtdcCommonIntType
	// 商品群号
	CommodityGroupID types.TThostFtdcCommodityGroupIDType
	// 商品群名称
	CommodityGroupName types.TThostFtdcInstrumentNameType
	// 操作标志
	ActionDirection types.TThostFtdcActionDirectionType
	// 追平序号
	SyncDeltaSequenceNo types.TThostFtdcSequenceNoType
}

func (d CThostFtdcSyncDeltaRULEInterParameterField) Type() string {
	return "CThostFtdcSyncDeltaRULEInterParameterField"
}

func (d CThostFtdcSyncDeltaRULEInterParameterField) String() string {
	var builder strings.Builder
	builder.Grow(331)

	builder.WriteString("CThostFtdcSyncDeltaRULEInterParameterField{")
	fmt.Fprintf(&builder, "TradingDay=%+v", d.TradingDay)
	fmt.Fprintf(&builder, ", ExchangeID=%+v", d.ExchangeID)
	fmt.Fprintf(&builder, ", SpreadId=%+v", d.SpreadId)
	fmt.Fprintf(&builder, ", InterRate=%+v", d.InterRate)
	fmt.Fprintf(&builder, ", Leg1ProdFamilyCode=%+v", d.Leg1ProdFamilyCode)
	fmt.Fprintf(&builder, ", Leg2ProdFamilyCode=%+v", d.Leg2ProdFamilyCode)
	fmt.Fprintf(&builder, ", Leg1PropFactor=%+v", d.Leg1PropFactor)
	fmt.Fprintf(&builder, ", Leg2PropFactor=%+v", d.Leg2PropFactor)
	fmt.Fprintf(&builder, ", CommodityGroupID=%+v", d.CommodityGroupID)
	fmt.Fprintf(&builder, ", CommodityGroupName=%+v", d.CommodityGroupName)
	fmt.Fprintf(&builder, ", ActionDirection=%+v", d.ActionDirection)
	fmt.Fprintf(&builder, ", SyncDeltaSequenceNo=%+v", d.SyncDeltaSequenceNo)

	builder.WriteByte('}')

	return builder.String()
}

// 服务地址参数
type CThostFtdcIpAddrParamField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 服务地址
	Address types.TThostFtdcIpAddrType
	// 交易中心代码
	DRIdentityID types.TThostFtdcDRIdentityIDType
	// 交易中心名称
	DRIdentityName types.TThostFtdcDRIdentityNameType
	// 交易地址OR行情地址
	AddrSrvMode types.TThostFtdcAddrSrvModeType
	// 地址版本
	AddrVer types.TThostFtdcAddrVerType
	// 服务地址编号
	AddrNo types.TThostFtdcCommonIntType
	// 服务地址名称
	AddrName types.TThostFtdcAddrNameType
	// 是否是国密地址
	IsSM types.TThostFtdcBoolType
	// 是否是内网地址
	IsLocalAddr types.TThostFtdcBoolType
	// 地址补充信息
	Remark types.TThostFtdcAddrRemarkType
	// 站点
	Site types.TThostFtdcSiteType
	// 网络运营商
	NetOperator types.TThostFtdcNetOperatorType
	// 系统名称
	SysName types.TThostFtdcAddrNameType
}

func (d CThostFtdcIpAddrParamField) Type() string { return "CThostFtdcIpAddrParamField" }

func (d CThostFtdcIpAddrParamField) String() string {
	var builder strings.Builder
	builder.Grow(282)

	builder.WriteString("CThostFtdcIpAddrParamField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", Address=%+v", d.Address)
	fmt.Fprintf(&builder, ", DRIdentityID=%+v", d.DRIdentityID)
	fmt.Fprintf(&builder, ", DRIdentityName=%+v", d.DRIdentityName)
	fmt.Fprintf(&builder, ", AddrSrvMode=%+v", d.AddrSrvMode)
	fmt.Fprintf(&builder, ", AddrVer=%+v", d.AddrVer)
	fmt.Fprintf(&builder, ", AddrNo=%+v", d.AddrNo)
	fmt.Fprintf(&builder, ", AddrName=%+v", d.AddrName)
	fmt.Fprintf(&builder, ", IsSM=%+v", d.IsSM)
	fmt.Fprintf(&builder, ", IsLocalAddr=%+v", d.IsLocalAddr)
	fmt.Fprintf(&builder, ", Remark=%+v", d.Remark)
	fmt.Fprintf(&builder, ", Site=%+v", d.Site)
	fmt.Fprintf(&builder, ", NetOperator=%+v", d.NetOperator)
	fmt.Fprintf(&builder, ", SysName=%+v", d.SysName)

	builder.WriteByte('}')

	return builder.String()
}

// 服务地址参数查询
type CThostFtdcQryIpAddrParamField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
}

func (d CThostFtdcQryIpAddrParamField) Type() string { return "CThostFtdcQryIpAddrParamField" }

func (d CThostFtdcQryIpAddrParamField) String() string {
	var builder strings.Builder
	builder.Grow(47)

	builder.WriteString("CThostFtdcQryIpAddrParamField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)

	builder.WriteByte('}')

	return builder.String()
}

// 服务地址参数
type CThostFtdcTGIpAddrParamField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 用户代码
	UserID types.TThostFtdcUserIDType
	// 服务地址
	Address types.TThostFtdcIpAddrType
	// 交易中心代码
	DRIdentityID types.TThostFtdcDRIdentityIDType
	// 交易中心名称
	DRIdentityName types.TThostFtdcDRIdentityNameType
	// 交易地址OR行情地址
	AddrSrvMode types.TThostFtdcAddrSrvModeType
	// 地址版本
	AddrVer types.TThostFtdcAddrVerType
	// 服务地址编号
	AddrNo types.TThostFtdcCommonIntType
	// 服务地址名称
	AddrName types.TThostFtdcAddrNameType
	// 是否是国密地址
	IsSM types.TThostFtdcBoolType
	// 是否是内网地址
	IsLocalAddr types.TThostFtdcBoolType
	// 地址补充信息
	Remark types.TThostFtdcAddrRemarkType
	// 站点
	Site types.TThostFtdcSiteType
	// 网络运营商
	NetOperator types.TThostFtdcNetOperatorType
	// 系统名称
	SysName types.TThostFtdcAddrNameType
}

func (d CThostFtdcTGIpAddrParamField) Type() string { return "CThostFtdcTGIpAddrParamField" }

func (d CThostFtdcTGIpAddrParamField) String() string {
	var builder strings.Builder
	builder.Grow(300)

	builder.WriteString("CThostFtdcTGIpAddrParamField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", UserID=%+v", d.UserID)
	fmt.Fprintf(&builder, ", Address=%+v", d.Address)
	fmt.Fprintf(&builder, ", DRIdentityID=%+v", d.DRIdentityID)
	fmt.Fprintf(&builder, ", DRIdentityName=%+v", d.DRIdentityName)
	fmt.Fprintf(&builder, ", AddrSrvMode=%+v", d.AddrSrvMode)
	fmt.Fprintf(&builder, ", AddrVer=%+v", d.AddrVer)
	fmt.Fprintf(&builder, ", AddrNo=%+v", d.AddrNo)
	fmt.Fprintf(&builder, ", AddrName=%+v", d.AddrName)
	fmt.Fprintf(&builder, ", IsSM=%+v", d.IsSM)
	fmt.Fprintf(&builder, ", IsLocalAddr=%+v", d.IsLocalAddr)
	fmt.Fprintf(&builder, ", Remark=%+v", d.Remark)
	fmt.Fprintf(&builder, ", Site=%+v", d.Site)
	fmt.Fprintf(&builder, ", NetOperator=%+v", d.NetOperator)
	fmt.Fprintf(&builder, ", SysName=%+v", d.SysName)

	builder.WriteByte('}')

	return builder.String()
}

// 服务地址参数查询
type CThostFtdcQryTGIpAddrParamField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 用户代码
	UserID types.TThostFtdcUserIDType
	// App代码
	AppID types.TThostFtdcAppIDType
}

func (d CThostFtdcQryTGIpAddrParamField) Type() string { return "CThostFtdcQryTGIpAddrParamField" }

func (d CThostFtdcQryTGIpAddrParamField) String() string {
	var builder strings.Builder
	builder.Grow(80)

	builder.WriteString("CThostFtdcQryTGIpAddrParamField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", UserID=%+v", d.UserID)
	fmt.Fprintf(&builder, ", AppID=%+v", d.AppID)

	builder.WriteByte('}')

	return builder.String()
}

// TGate会话查询状态
type CThostFtdcTGSessionQryStatusField struct {
	// 最近30s的查询频率
	LastQryFreq types.TThostFtdcCommonIntType
	// 查询状态
	QryStatus types.TThostFtdcTGSessionQryStatusType
}

func (d CThostFtdcTGSessionQryStatusField) Type() string { return "CThostFtdcTGSessionQryStatusField" }

func (d CThostFtdcTGSessionQryStatusField) String() string {
	var builder strings.Builder
	builder.Grow(73)

	builder.WriteString("CThostFtdcTGSessionQryStatusField{")
	fmt.Fprintf(&builder, "LastQryFreq=%+v", d.LastQryFreq)
	fmt.Fprintf(&builder, ", QryStatus=%+v", d.QryStatus)

	builder.WriteByte('}')

	return builder.String()
}

// 内网地址配置
type CThostFtdcLocalAddrConfigField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 对端地址
	PeerAddr types.TThostFtdcIpAddrType
	// 子网掩码
	NetMask types.TThostFtdcIpAddrType
	// 交易中心代码
	DRIdentityID types.TThostFtdcDRIdentityIDType
	// 内网服务地址
	LocalAddress types.TThostFtdcIpAddrType
}

func (d CThostFtdcLocalAddrConfigField) Type() string { return "CThostFtdcLocalAddrConfigField" }

func (d CThostFtdcLocalAddrConfigField) String() string {
	var builder strings.Builder
	builder.Grow(127)

	builder.WriteString("CThostFtdcLocalAddrConfigField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", PeerAddr=%+v", d.PeerAddr)
	fmt.Fprintf(&builder, ", NetMask=%+v", d.NetMask)
	fmt.Fprintf(&builder, ", DRIdentityID=%+v", d.DRIdentityID)
	fmt.Fprintf(&builder, ", LocalAddress=%+v", d.LocalAddress)

	builder.WriteByte('}')

	return builder.String()
}

// 内网地址配置查询
type CThostFtdcQryLocalAddrConfigField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
}

func (d CThostFtdcQryLocalAddrConfigField) Type() string { return "CThostFtdcQryLocalAddrConfigField" }

func (d CThostFtdcQryLocalAddrConfigField) String() string {
	var builder strings.Builder
	builder.Grow(51)

	builder.WriteString("CThostFtdcQryLocalAddrConfigField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)

	builder.WriteByte('}')

	return builder.String()
}

// 次席查询银行资金帐户信息请求
type CThostFtdcReqQueryBankAccountBySecField struct {
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
	// 长客户姓名
	LongCustomerName types.TThostFtdcLongIndividualNameType
	// 交易中心代码
	DRIdentityID types.TThostFtdcDRIdentityIDType
	// 次中心发起转账期货公司流水号
	SecFutureSerial types.TThostFtdcFutureSerialType
}

func (d CThostFtdcReqQueryBankAccountBySecField) Type() string {
	return "CThostFtdcReqQueryBankAccountBySecField"
}

func (d CThostFtdcReqQueryBankAccountBySecField) String() string {
	var builder strings.Builder
	builder.Grow(835)

	builder.WriteString("CThostFtdcReqQueryBankAccountBySecField{")
	fmt.Fprintf(&builder, "TradeCode=%+v", d.TradeCode)
	fmt.Fprintf(&builder, ", BankID=%+v", d.BankID)
	fmt.Fprintf(&builder, ", BankBranchID=%+v", d.BankBranchID)
	fmt.Fprintf(&builder, ", BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", BrokerBranchID=%+v", d.BrokerBranchID)
	fmt.Fprintf(&builder, ", TradeDate=%+v", d.TradeDate)
	fmt.Fprintf(&builder, ", TradeTime=%+v", d.TradeTime)
	fmt.Fprintf(&builder, ", BankSerial=%+v", d.BankSerial)
	fmt.Fprintf(&builder, ", TradingDay=%+v", d.TradingDay)
	fmt.Fprintf(&builder, ", PlateSerial=%+v", d.PlateSerial)
	fmt.Fprintf(&builder, ", LastFragment=%+v", d.LastFragment)
	fmt.Fprintf(&builder, ", SessionID=%+v", d.SessionID)
	fmt.Fprintf(&builder, ", CustomerName=%+v", d.CustomerName)
	fmt.Fprintf(&builder, ", IdCardType=%+v", d.IdCardType)
	fmt.Fprintf(&builder, ", IdentifiedCardNo=%+v", d.IdentifiedCardNo)
	fmt.Fprintf(&builder, ", CustType=%+v", d.CustType)
	fmt.Fprintf(&builder, ", BankAccount=%+v", d.BankAccount)
	fmt.Fprintf(&builder, ", BankPassWord=%+v", d.BankPassWord)
	fmt.Fprintf(&builder, ", AccountID=%+v", d.AccountID)
	fmt.Fprintf(&builder, ", Password=%+v", d.Password)
	fmt.Fprintf(&builder, ", FutureSerial=%+v", d.FutureSerial)
	fmt.Fprintf(&builder, ", InstallID=%+v", d.InstallID)
	fmt.Fprintf(&builder, ", UserID=%+v", d.UserID)
	fmt.Fprintf(&builder, ", VerifyCertNoFlag=%+v", d.VerifyCertNoFlag)
	fmt.Fprintf(&builder, ", CurrencyID=%+v", d.CurrencyID)
	fmt.Fprintf(&builder, ", Digest=%+v", d.Digest)
	fmt.Fprintf(&builder, ", BankAccType=%+v", d.BankAccType)
	fmt.Fprintf(&builder, ", DeviceID=%+v", d.DeviceID)
	fmt.Fprintf(&builder, ", BankSecuAccType=%+v", d.BankSecuAccType)
	fmt.Fprintf(&builder, ", BrokerIDByBank=%+v", d.BrokerIDByBank)
	fmt.Fprintf(&builder, ", BankSecuAcc=%+v", d.BankSecuAcc)
	fmt.Fprintf(&builder, ", BankPwdFlag=%+v", d.BankPwdFlag)
	fmt.Fprintf(&builder, ", SecuPwdFlag=%+v", d.SecuPwdFlag)
	fmt.Fprintf(&builder, ", OperNo=%+v", d.OperNo)
	fmt.Fprintf(&builder, ", RequestID=%+v", d.RequestID)
	fmt.Fprintf(&builder, ", TID=%+v", d.TID)
	fmt.Fprintf(&builder, ", LongCustomerName=%+v", d.LongCustomerName)
	fmt.Fprintf(&builder, ", DRIdentityID=%+v", d.DRIdentityID)
	fmt.Fprintf(&builder, ", SecFutureSerial=%+v", d.SecFutureSerial)

	builder.WriteByte('}')

	return builder.String()
}

// 次席查询银行资金帐户信息回报
type CThostFtdcRspQueryBankAccountBySecField struct {
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
	// 长客户姓名
	LongCustomerName types.TThostFtdcLongIndividualNameType
	// 交易中心代码
	DRIdentityID types.TThostFtdcDRIdentityIDType
	// 次中心发起转账期货公司流水号
	SecFutureSerial types.TThostFtdcFutureSerialType
}

func (d CThostFtdcRspQueryBankAccountBySecField) Type() string {
	return "CThostFtdcRspQueryBankAccountBySecField"
}

func (d CThostFtdcRspQueryBankAccountBySecField) String() string {
	var builder strings.Builder
	builder.Grow(883)

	builder.WriteString("CThostFtdcRspQueryBankAccountBySecField{")
	fmt.Fprintf(&builder, "TradeCode=%+v", d.TradeCode)
	fmt.Fprintf(&builder, ", BankID=%+v", d.BankID)
	fmt.Fprintf(&builder, ", BankBranchID=%+v", d.BankBranchID)
	fmt.Fprintf(&builder, ", BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", BrokerBranchID=%+v", d.BrokerBranchID)
	fmt.Fprintf(&builder, ", TradeDate=%+v", d.TradeDate)
	fmt.Fprintf(&builder, ", TradeTime=%+v", d.TradeTime)
	fmt.Fprintf(&builder, ", BankSerial=%+v", d.BankSerial)
	fmt.Fprintf(&builder, ", TradingDay=%+v", d.TradingDay)
	fmt.Fprintf(&builder, ", PlateSerial=%+v", d.PlateSerial)
	fmt.Fprintf(&builder, ", LastFragment=%+v", d.LastFragment)
	fmt.Fprintf(&builder, ", SessionID=%+v", d.SessionID)
	fmt.Fprintf(&builder, ", CustomerName=%+v", d.CustomerName)
	fmt.Fprintf(&builder, ", IdCardType=%+v", d.IdCardType)
	fmt.Fprintf(&builder, ", IdentifiedCardNo=%+v", d.IdentifiedCardNo)
	fmt.Fprintf(&builder, ", CustType=%+v", d.CustType)
	fmt.Fprintf(&builder, ", BankAccount=%+v", d.BankAccount)
	fmt.Fprintf(&builder, ", BankPassWord=%+v", d.BankPassWord)
	fmt.Fprintf(&builder, ", AccountID=%+v", d.AccountID)
	fmt.Fprintf(&builder, ", Password=%+v", d.Password)
	fmt.Fprintf(&builder, ", FutureSerial=%+v", d.FutureSerial)
	fmt.Fprintf(&builder, ", InstallID=%+v", d.InstallID)
	fmt.Fprintf(&builder, ", UserID=%+v", d.UserID)
	fmt.Fprintf(&builder, ", VerifyCertNoFlag=%+v", d.VerifyCertNoFlag)
	fmt.Fprintf(&builder, ", CurrencyID=%+v", d.CurrencyID)
	fmt.Fprintf(&builder, ", Digest=%+v", d.Digest)
	fmt.Fprintf(&builder, ", BankAccType=%+v", d.BankAccType)
	fmt.Fprintf(&builder, ", DeviceID=%+v", d.DeviceID)
	fmt.Fprintf(&builder, ", BankSecuAccType=%+v", d.BankSecuAccType)
	fmt.Fprintf(&builder, ", BrokerIDByBank=%+v", d.BrokerIDByBank)
	fmt.Fprintf(&builder, ", BankSecuAcc=%+v", d.BankSecuAcc)
	fmt.Fprintf(&builder, ", BankPwdFlag=%+v", d.BankPwdFlag)
	fmt.Fprintf(&builder, ", SecuPwdFlag=%+v", d.SecuPwdFlag)
	fmt.Fprintf(&builder, ", OperNo=%+v", d.OperNo)
	fmt.Fprintf(&builder, ", RequestID=%+v", d.RequestID)
	fmt.Fprintf(&builder, ", TID=%+v", d.TID)
	fmt.Fprintf(&builder, ", BankUseAmount=%+v", d.BankUseAmount)
	fmt.Fprintf(&builder, ", BankFetchAmount=%+v", d.BankFetchAmount)
	fmt.Fprintf(&builder, ", LongCustomerName=%+v", d.LongCustomerName)
	fmt.Fprintf(&builder, ", DRIdentityID=%+v", d.DRIdentityID)
	fmt.Fprintf(&builder, ", SecFutureSerial=%+v", d.SecFutureSerial)

	builder.WriteByte('}')

	return builder.String()
}

// 次中心发起的转帐交易
type CThostFtdcReqTransferBySecField struct {
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
	// 长客户姓名
	LongCustomerName types.TThostFtdcLongIndividualNameType
	// 交易中心代码
	DRIdentityID types.TThostFtdcDRIdentityIDType
	// 次中心发起转账期货公司流水号
	SecFutureSerial types.TThostFtdcFutureSerialType
}

func (d CThostFtdcReqTransferBySecField) Type() string { return "CThostFtdcReqTransferBySecField" }

func (d CThostFtdcReqTransferBySecField) String() string {
	var builder strings.Builder
	builder.Grow(972)

	builder.WriteString("CThostFtdcReqTransferBySecField{")
	fmt.Fprintf(&builder, "TradeCode=%+v", d.TradeCode)
	fmt.Fprintf(&builder, ", BankID=%+v", d.BankID)
	fmt.Fprintf(&builder, ", BankBranchID=%+v", d.BankBranchID)
	fmt.Fprintf(&builder, ", BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", BrokerBranchID=%+v", d.BrokerBranchID)
	fmt.Fprintf(&builder, ", TradeDate=%+v", d.TradeDate)
	fmt.Fprintf(&builder, ", TradeTime=%+v", d.TradeTime)
	fmt.Fprintf(&builder, ", BankSerial=%+v", d.BankSerial)
	fmt.Fprintf(&builder, ", TradingDay=%+v", d.TradingDay)
	fmt.Fprintf(&builder, ", PlateSerial=%+v", d.PlateSerial)
	fmt.Fprintf(&builder, ", LastFragment=%+v", d.LastFragment)
	fmt.Fprintf(&builder, ", SessionID=%+v", d.SessionID)
	fmt.Fprintf(&builder, ", CustomerName=%+v", d.CustomerName)
	fmt.Fprintf(&builder, ", IdCardType=%+v", d.IdCardType)
	fmt.Fprintf(&builder, ", IdentifiedCardNo=%+v", d.IdentifiedCardNo)
	fmt.Fprintf(&builder, ", CustType=%+v", d.CustType)
	fmt.Fprintf(&builder, ", BankAccount=%+v", d.BankAccount)
	fmt.Fprintf(&builder, ", BankPassWord=%+v", d.BankPassWord)
	fmt.Fprintf(&builder, ", AccountID=%+v", d.AccountID)
	fmt.Fprintf(&builder, ", Password=%+v", d.Password)
	fmt.Fprintf(&builder, ", InstallID=%+v", d.InstallID)
	fmt.Fprintf(&builder, ", FutureSerial=%+v", d.FutureSerial)
	fmt.Fprintf(&builder, ", UserID=%+v", d.UserID)
	fmt.Fprintf(&builder, ", VerifyCertNoFlag=%+v", d.VerifyCertNoFlag)
	fmt.Fprintf(&builder, ", CurrencyID=%+v", d.CurrencyID)
	fmt.Fprintf(&builder, ", TradeAmount=%+v", d.TradeAmount)
	fmt.Fprintf(&builder, ", FutureFetchAmount=%+v", d.FutureFetchAmount)
	fmt.Fprintf(&builder, ", FeePayFlag=%+v", d.FeePayFlag)
	fmt.Fprintf(&builder, ", CustFee=%+v", d.CustFee)
	fmt.Fprintf(&builder, ", BrokerFee=%+v", d.BrokerFee)
	fmt.Fprintf(&builder, ", Message=%+v", d.Message)
	fmt.Fprintf(&builder, ", Digest=%+v", d.Digest)
	fmt.Fprintf(&builder, ", BankAccType=%+v", d.BankAccType)
	fmt.Fprintf(&builder, ", DeviceID=%+v", d.DeviceID)
	fmt.Fprintf(&builder, ", BankSecuAccType=%+v", d.BankSecuAccType)
	fmt.Fprintf(&builder, ", BrokerIDByBank=%+v", d.BrokerIDByBank)
	fmt.Fprintf(&builder, ", BankSecuAcc=%+v", d.BankSecuAcc)
	fmt.Fprintf(&builder, ", BankPwdFlag=%+v", d.BankPwdFlag)
	fmt.Fprintf(&builder, ", SecuPwdFlag=%+v", d.SecuPwdFlag)
	fmt.Fprintf(&builder, ", OperNo=%+v", d.OperNo)
	fmt.Fprintf(&builder, ", RequestID=%+v", d.RequestID)
	fmt.Fprintf(&builder, ", TID=%+v", d.TID)
	fmt.Fprintf(&builder, ", TransferStatus=%+v", d.TransferStatus)
	fmt.Fprintf(&builder, ", LongCustomerName=%+v", d.LongCustomerName)
	fmt.Fprintf(&builder, ", DRIdentityID=%+v", d.DRIdentityID)
	fmt.Fprintf(&builder, ", SecFutureSerial=%+v", d.SecFutureSerial)

	builder.WriteByte('}')

	return builder.String()
}

// 次中心发起的转帐交易回报
type CThostFtdcRspTransferBySecField struct {
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
	// 长客户姓名
	LongCustomerName types.TThostFtdcLongIndividualNameType
	// 交易中心代码
	DRIdentityID types.TThostFtdcDRIdentityIDType
	// 次中心发起转账期货公司流水号
	SecFutureSerial types.TThostFtdcFutureSerialType
}

func (d CThostFtdcRspTransferBySecField) Type() string { return "CThostFtdcRspTransferBySecField" }

func (d CThostFtdcRspTransferBySecField) String() string {
	var builder strings.Builder
	builder.Grow(1007)

	builder.WriteString("CThostFtdcRspTransferBySecField{")
	fmt.Fprintf(&builder, "TradeCode=%+v", d.TradeCode)
	fmt.Fprintf(&builder, ", BankID=%+v", d.BankID)
	fmt.Fprintf(&builder, ", BankBranchID=%+v", d.BankBranchID)
	fmt.Fprintf(&builder, ", BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", BrokerBranchID=%+v", d.BrokerBranchID)
	fmt.Fprintf(&builder, ", TradeDate=%+v", d.TradeDate)
	fmt.Fprintf(&builder, ", TradeTime=%+v", d.TradeTime)
	fmt.Fprintf(&builder, ", BankSerial=%+v", d.BankSerial)
	fmt.Fprintf(&builder, ", TradingDay=%+v", d.TradingDay)
	fmt.Fprintf(&builder, ", PlateSerial=%+v", d.PlateSerial)
	fmt.Fprintf(&builder, ", LastFragment=%+v", d.LastFragment)
	fmt.Fprintf(&builder, ", SessionID=%+v", d.SessionID)
	fmt.Fprintf(&builder, ", CustomerName=%+v", d.CustomerName)
	fmt.Fprintf(&builder, ", IdCardType=%+v", d.IdCardType)
	fmt.Fprintf(&builder, ", IdentifiedCardNo=%+v", d.IdentifiedCardNo)
	fmt.Fprintf(&builder, ", CustType=%+v", d.CustType)
	fmt.Fprintf(&builder, ", BankAccount=%+v", d.BankAccount)
	fmt.Fprintf(&builder, ", BankPassWord=%+v", d.BankPassWord)
	fmt.Fprintf(&builder, ", AccountID=%+v", d.AccountID)
	fmt.Fprintf(&builder, ", Password=%+v", d.Password)
	fmt.Fprintf(&builder, ", InstallID=%+v", d.InstallID)
	fmt.Fprintf(&builder, ", FutureSerial=%+v", d.FutureSerial)
	fmt.Fprintf(&builder, ", UserID=%+v", d.UserID)
	fmt.Fprintf(&builder, ", VerifyCertNoFlag=%+v", d.VerifyCertNoFlag)
	fmt.Fprintf(&builder, ", CurrencyID=%+v", d.CurrencyID)
	fmt.Fprintf(&builder, ", TradeAmount=%+v", d.TradeAmount)
	fmt.Fprintf(&builder, ", FutureFetchAmount=%+v", d.FutureFetchAmount)
	fmt.Fprintf(&builder, ", FeePayFlag=%+v", d.FeePayFlag)
	fmt.Fprintf(&builder, ", CustFee=%+v", d.CustFee)
	fmt.Fprintf(&builder, ", BrokerFee=%+v", d.BrokerFee)
	fmt.Fprintf(&builder, ", Message=%+v", d.Message)
	fmt.Fprintf(&builder, ", Digest=%+v", d.Digest)
	fmt.Fprintf(&builder, ", BankAccType=%+v", d.BankAccType)
	fmt.Fprintf(&builder, ", DeviceID=%+v", d.DeviceID)
	fmt.Fprintf(&builder, ", BankSecuAccType=%+v", d.BankSecuAccType)
	fmt.Fprintf(&builder, ", BrokerIDByBank=%+v", d.BrokerIDByBank)
	fmt.Fprintf(&builder, ", BankSecuAcc=%+v", d.BankSecuAcc)
	fmt.Fprintf(&builder, ", BankPwdFlag=%+v", d.BankPwdFlag)
	fmt.Fprintf(&builder, ", SecuPwdFlag=%+v", d.SecuPwdFlag)
	fmt.Fprintf(&builder, ", OperNo=%+v", d.OperNo)
	fmt.Fprintf(&builder, ", RequestID=%+v", d.RequestID)
	fmt.Fprintf(&builder, ", TID=%+v", d.TID)
	fmt.Fprintf(&builder, ", TransferStatus=%+v", d.TransferStatus)
	fmt.Fprintf(&builder, ", ErrorID=%+v", d.ErrorID)
	fmt.Fprintf(&builder, ", ErrorMsg=%+v", d.ErrorMsg)
	fmt.Fprintf(&builder, ", LongCustomerName=%+v", d.LongCustomerName)
	fmt.Fprintf(&builder, ", DRIdentityID=%+v", d.DRIdentityID)
	fmt.Fprintf(&builder, ", SecFutureSerial=%+v", d.SecFutureSerial)

	builder.WriteByte('}')

	return builder.String()
}

// 查询银行资金帐户信息通知 要发往次席
type CThostFtdcNotifyQueryFutureAccountBySecField struct {
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
	// 长客户姓名
	LongCustomerName types.TThostFtdcLongIndividualNameType
	// 交易中心代码
	DRIdentityID types.TThostFtdcDRIdentityIDType
	// 次中心发起转账期货公司流水号
	SecFutureSerial types.TThostFtdcFutureSerialType
}

func (d CThostFtdcNotifyQueryFutureAccountBySecField) Type() string {
	return "CThostFtdcNotifyQueryFutureAccountBySecField"
}

func (d CThostFtdcNotifyQueryFutureAccountBySecField) String() string {
	var builder strings.Builder
	builder.Grow(923)

	builder.WriteString("CThostFtdcNotifyQueryFutureAccountBySecField{")
	fmt.Fprintf(&builder, "TradeCode=%+v", d.TradeCode)
	fmt.Fprintf(&builder, ", BankID=%+v", d.BankID)
	fmt.Fprintf(&builder, ", BankBranchID=%+v", d.BankBranchID)
	fmt.Fprintf(&builder, ", BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", BrokerBranchID=%+v", d.BrokerBranchID)
	fmt.Fprintf(&builder, ", TradeDate=%+v", d.TradeDate)
	fmt.Fprintf(&builder, ", TradeTime=%+v", d.TradeTime)
	fmt.Fprintf(&builder, ", BankSerial=%+v", d.BankSerial)
	fmt.Fprintf(&builder, ", TradingDay=%+v", d.TradingDay)
	fmt.Fprintf(&builder, ", PlateSerial=%+v", d.PlateSerial)
	fmt.Fprintf(&builder, ", LastFragment=%+v", d.LastFragment)
	fmt.Fprintf(&builder, ", SessionID=%+v", d.SessionID)
	fmt.Fprintf(&builder, ", CustomerName=%+v", d.CustomerName)
	fmt.Fprintf(&builder, ", IdCardType=%+v", d.IdCardType)
	fmt.Fprintf(&builder, ", IdentifiedCardNo=%+v", d.IdentifiedCardNo)
	fmt.Fprintf(&builder, ", CustType=%+v", d.CustType)
	fmt.Fprintf(&builder, ", BankAccount=%+v", d.BankAccount)
	fmt.Fprintf(&builder, ", BankPassWord=%+v", d.BankPassWord)
	fmt.Fprintf(&builder, ", AccountID=%+v", d.AccountID)
	fmt.Fprintf(&builder, ", Password=%+v", d.Password)
	fmt.Fprintf(&builder, ", FutureSerial=%+v", d.FutureSerial)
	fmt.Fprintf(&builder, ", InstallID=%+v", d.InstallID)
	fmt.Fprintf(&builder, ", UserID=%+v", d.UserID)
	fmt.Fprintf(&builder, ", VerifyCertNoFlag=%+v", d.VerifyCertNoFlag)
	fmt.Fprintf(&builder, ", CurrencyID=%+v", d.CurrencyID)
	fmt.Fprintf(&builder, ", Digest=%+v", d.Digest)
	fmt.Fprintf(&builder, ", BankAccType=%+v", d.BankAccType)
	fmt.Fprintf(&builder, ", DeviceID=%+v", d.DeviceID)
	fmt.Fprintf(&builder, ", BankSecuAccType=%+v", d.BankSecuAccType)
	fmt.Fprintf(&builder, ", BrokerIDByBank=%+v", d.BrokerIDByBank)
	fmt.Fprintf(&builder, ", BankSecuAcc=%+v", d.BankSecuAcc)
	fmt.Fprintf(&builder, ", BankPwdFlag=%+v", d.BankPwdFlag)
	fmt.Fprintf(&builder, ", SecuPwdFlag=%+v", d.SecuPwdFlag)
	fmt.Fprintf(&builder, ", OperNo=%+v", d.OperNo)
	fmt.Fprintf(&builder, ", RequestID=%+v", d.RequestID)
	fmt.Fprintf(&builder, ", TID=%+v", d.TID)
	fmt.Fprintf(&builder, ", BankUseAmount=%+v", d.BankUseAmount)
	fmt.Fprintf(&builder, ", BankFetchAmount=%+v", d.BankFetchAmount)
	fmt.Fprintf(&builder, ", ErrorID=%+v", d.ErrorID)
	fmt.Fprintf(&builder, ", ErrorMsg=%+v", d.ErrorMsg)
	fmt.Fprintf(&builder, ", LongCustomerName=%+v", d.LongCustomerName)
	fmt.Fprintf(&builder, ", DRIdentityID=%+v", d.DRIdentityID)
	fmt.Fprintf(&builder, ", SecFutureSerial=%+v", d.SecFutureSerial)

	builder.WriteByte('}')

	return builder.String()
}

// 退出紧急状态参数
type CThostFtdcExitEmergencyField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
}

func (d CThostFtdcExitEmergencyField) Type() string { return "CThostFtdcExitEmergencyField" }

func (d CThostFtdcExitEmergencyField) String() string {
	var builder strings.Builder
	builder.Grow(46)

	builder.WriteString("CThostFtdcExitEmergencyField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)

	builder.WriteByte('}')

	return builder.String()
}

// 新组保保证金系数投资者模板对应关系
type CThostFtdcInvestorPortfMarginModelField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	// 保证金系数模板
	MarginModelID types.TThostFtdcInvestorIDType
}

func (d CThostFtdcInvestorPortfMarginModelField) Type() string {
	return "CThostFtdcInvestorPortfMarginModelField"
}

func (d CThostFtdcInvestorPortfMarginModelField) String() string {
	var builder strings.Builder
	builder.Grow(100)

	builder.WriteString("CThostFtdcInvestorPortfMarginModelField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", InvestorID=%+v", d.InvestorID)
	fmt.Fprintf(&builder, ", MarginModelID=%+v", d.MarginModelID)

	builder.WriteByte('}')

	return builder.String()
}

// 投资者新组保设置
type CThostFtdcInvestorPortfSettingField struct {
	// 交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 投资者编号
	InvestorID types.TThostFtdcInvestorIDType
	// 投机套保标志
	HedgeFlag types.TThostFtdcHedgeFlagType
	// 是否开启新组保
	UsePortf types.TThostFtdcBoolType
}

func (d CThostFtdcInvestorPortfSettingField) Type() string {
	return "CThostFtdcInvestorPortfSettingField"
}

func (d CThostFtdcInvestorPortfSettingField) String() string {
	var builder strings.Builder
	builder.Grow(130)

	builder.WriteString("CThostFtdcInvestorPortfSettingField{")
	fmt.Fprintf(&builder, "ExchangeID=%+v", d.ExchangeID)
	fmt.Fprintf(&builder, ", BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", InvestorID=%+v", d.InvestorID)
	fmt.Fprintf(&builder, ", HedgeFlag=%+v", d.HedgeFlag)
	fmt.Fprintf(&builder, ", UsePortf=%+v", d.UsePortf)

	builder.WriteByte('}')

	return builder.String()
}

// 投资者新组保设置查询
type CThostFtdcQryInvestorPortfSettingField struct {
	// 交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 投资者编号
	InvestorID types.TThostFtdcInvestorIDType
}

func (d CThostFtdcQryInvestorPortfSettingField) Type() string {
	return "CThostFtdcQryInvestorPortfSettingField"
}

func (d CThostFtdcQryInvestorPortfSettingField) String() string {
	var builder strings.Builder
	builder.Grow(96)

	builder.WriteString("CThostFtdcQryInvestorPortfSettingField{")
	fmt.Fprintf(&builder, "ExchangeID=%+v", d.ExchangeID)
	fmt.Fprintf(&builder, ", BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", InvestorID=%+v", d.InvestorID)

	builder.WriteByte('}')

	return builder.String()
}

// 来自次席的用户口令变更
type CThostFtdcUserPasswordUpdateFromSecField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 用户代码
	UserID types.TThostFtdcUserIDType
	// 原来的口令
	OldPassword types.TThostFtdcPasswordType
	// 新的口令
	NewPassword types.TThostFtdcPasswordType
	// 次席的交易中心代码
	FromSec types.TThostFtdcDRIdentityIDType
}

func (d CThostFtdcUserPasswordUpdateFromSecField) Type() string {
	return "CThostFtdcUserPasswordUpdateFromSecField"
}

func (d CThostFtdcUserPasswordUpdateFromSecField) String() string {
	var builder strings.Builder
	builder.Grow(133)

	builder.WriteString("CThostFtdcUserPasswordUpdateFromSecField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", UserID=%+v", d.UserID)
	fmt.Fprintf(&builder, ", OldPassword=%+v", d.OldPassword)
	fmt.Fprintf(&builder, ", NewPassword=%+v", d.NewPassword)
	fmt.Fprintf(&builder, ", FromSec=%+v", d.FromSec)

	builder.WriteByte('}')

	return builder.String()
}

// 来自次席的结算结果确认
type CThostFtdcSettlementInfoConfirmFromSecField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	// 确认日期
	ConfirmDate types.TThostFtdcDateType
	// 确认时间
	ConfirmTime types.TThostFtdcTimeType
	// 次席的交易中心代码
	FromSec types.TThostFtdcDRIdentityIDType
}

func (d CThostFtdcSettlementInfoConfirmFromSecField) Type() string {
	return "CThostFtdcSettlementInfoConfirmFromSecField"
}

func (d CThostFtdcSettlementInfoConfirmFromSecField) String() string {
	var builder strings.Builder
	builder.Grow(140)

	builder.WriteString("CThostFtdcSettlementInfoConfirmFromSecField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", InvestorID=%+v", d.InvestorID)
	fmt.Fprintf(&builder, ", ConfirmDate=%+v", d.ConfirmDate)
	fmt.Fprintf(&builder, ", ConfirmTime=%+v", d.ConfirmTime)
	fmt.Fprintf(&builder, ", FromSec=%+v", d.FromSec)

	builder.WriteByte('}')

	return builder.String()
}

// 来自次席的资金账户口令变更
type CThostFtdcTradingAccountPasswordUpdateFromSecField struct {
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
	// 次席的交易中心代码
	FromSec types.TThostFtdcDRIdentityIDType
}

func (d CThostFtdcTradingAccountPasswordUpdateFromSecField) Type() string {
	return "CThostFtdcTradingAccountPasswordUpdateFromSecField"
}

func (d CThostFtdcTradingAccountPasswordUpdateFromSecField) String() string {
	var builder strings.Builder
	builder.Grow(166)

	builder.WriteString("CThostFtdcTradingAccountPasswordUpdateFromSecField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", AccountID=%+v", d.AccountID)
	fmt.Fprintf(&builder, ", OldPassword=%+v", d.OldPassword)
	fmt.Fprintf(&builder, ", NewPassword=%+v", d.NewPassword)
	fmt.Fprintf(&builder, ", CurrencyID=%+v", d.CurrencyID)
	fmt.Fprintf(&builder, ", FromSec=%+v", d.FromSec)

	builder.WriteByte('}')

	return builder.String()
}

// 风控禁止的合约交易权限
type CThostFtdcRiskForbiddenRightField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 投资者编号
	InvestorID types.TThostFtdcInvestorIDType
	// 合约产品代码
	InstrumentID types.TThostFtdcInstrumentIDType
	// 用户代码
	UserID types.TThostFtdcUserIDType
}

func (d CThostFtdcRiskForbiddenRightField) Type() string { return "CThostFtdcRiskForbiddenRightField" }

func (d CThostFtdcRiskForbiddenRightField) String() string {
	var builder strings.Builder
	builder.Grow(109)

	builder.WriteString("CThostFtdcRiskForbiddenRightField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", InvestorID=%+v", d.InvestorID)
	fmt.Fprintf(&builder, ", InstrumentID=%+v", d.InstrumentID)
	fmt.Fprintf(&builder, ", UserID=%+v", d.UserID)

	builder.WriteByte('}')

	return builder.String()
}

// 投资者申报费阶梯收取记录
type CThostFtdcInvestorInfoCommRecField struct {
	// 交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	// 商品代码
	InstrumentID types.TThostFtdcInstrumentIDType
	// 报单总笔数
	OrderCount types.TThostFtdcVolumeType
	// 撤单总笔数
	OrderActionCount types.TThostFtdcVolumeType
	// 询价总次数
	ForQuoteCnt types.TThostFtdcVolumeType
	// 申报费
	InfoComm types.TThostFtdcMoneyType
	// 是否期权系列
	IsOptSeries types.TThostFtdcBoolType
	// 品种代码
	ProductID types.TThostFtdcProductIDType
	// 信息量总量
	InfoCnt types.TThostFtdcVolumeType
}

func (d CThostFtdcInvestorInfoCommRecField) Type() string {
	return "CThostFtdcInvestorInfoCommRecField"
}

func (d CThostFtdcInvestorInfoCommRecField) String() string {
	var builder strings.Builder
	builder.Grow(256)

	builder.WriteString("CThostFtdcInvestorInfoCommRecField{")
	fmt.Fprintf(&builder, "ExchangeID=%+v", d.ExchangeID)
	fmt.Fprintf(&builder, ", BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", InvestorID=%+v", d.InvestorID)
	fmt.Fprintf(&builder, ", InstrumentID=%+v", d.InstrumentID)
	fmt.Fprintf(&builder, ", OrderCount=%+v", d.OrderCount)
	fmt.Fprintf(&builder, ", OrderActionCount=%+v", d.OrderActionCount)
	fmt.Fprintf(&builder, ", ForQuoteCnt=%+v", d.ForQuoteCnt)
	fmt.Fprintf(&builder, ", InfoComm=%+v", d.InfoComm)
	fmt.Fprintf(&builder, ", IsOptSeries=%+v", d.IsOptSeries)
	fmt.Fprintf(&builder, ", ProductID=%+v", d.ProductID)
	fmt.Fprintf(&builder, ", InfoCnt=%+v", d.InfoCnt)

	builder.WriteByte('}')

	return builder.String()
}

// 投资者申报费阶梯收取记录查询
type CThostFtdcQryInvestorInfoCommRecField struct {
	// 投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	// 商品代码
	InstrumentID types.TThostFtdcInstrumentIDType
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
}

func (d CThostFtdcQryInvestorInfoCommRecField) Type() string {
	return "CThostFtdcQryInvestorInfoCommRecField"
}

func (d CThostFtdcQryInvestorInfoCommRecField) String() string {
	var builder strings.Builder
	builder.Grow(97)

	builder.WriteString("CThostFtdcQryInvestorInfoCommRecField{")
	fmt.Fprintf(&builder, "InvestorID=%+v", d.InvestorID)
	fmt.Fprintf(&builder, ", InstrumentID=%+v", d.InstrumentID)
	fmt.Fprintf(&builder, ", BrokerID=%+v", d.BrokerID)

	builder.WriteByte('}')

	return builder.String()
}

// 组合腿信息
type CThostFtdcCombLegField struct {
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

func (d CThostFtdcCombLegField) Type() string { return "CThostFtdcCombLegField" }

func (d CThostFtdcCombLegField) String() string {
	var builder strings.Builder
	builder.Grow(148)

	builder.WriteString("CThostFtdcCombLegField{")
	fmt.Fprintf(&builder, "CombInstrumentID=%+v", d.CombInstrumentID)
	fmt.Fprintf(&builder, ", LegID=%+v", d.LegID)
	fmt.Fprintf(&builder, ", LegInstrumentID=%+v", d.LegInstrumentID)
	fmt.Fprintf(&builder, ", Direction=%+v", d.Direction)
	fmt.Fprintf(&builder, ", LegMultiple=%+v", d.LegMultiple)
	fmt.Fprintf(&builder, ", ImplyLevel=%+v", d.ImplyLevel)

	builder.WriteByte('}')

	return builder.String()
}

// 组合腿信息查询
type CThostFtdcQryCombLegField struct {
	// 单腿合约代码
	LegInstrumentID types.TThostFtdcInstrumentIDType
}

func (d CThostFtdcQryCombLegField) Type() string { return "CThostFtdcQryCombLegField" }

func (d CThostFtdcQryCombLegField) String() string {
	var builder strings.Builder
	builder.Grow(50)

	builder.WriteString("CThostFtdcQryCombLegField{")
	fmt.Fprintf(&builder, "LegInstrumentID=%+v", d.LegInstrumentID)

	builder.WriteByte('}')

	return builder.String()
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
	// IP地址
	IPAddress types.TThostFtdcIPAddressType
	// Mac地址
	MacAddress types.TThostFtdcMacAddressType
}

func (d CThostFtdcInputOffsetSettingField) Type() string { return "CThostFtdcInputOffsetSettingField" }

func (d CThostFtdcInputOffsetSettingField) String() string {
	var builder strings.Builder
	builder.Grow(287)

	builder.WriteString("CThostFtdcInputOffsetSettingField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", InvestorID=%+v", d.InvestorID)
	fmt.Fprintf(&builder, ", InstrumentID=%+v", d.InstrumentID)
	fmt.Fprintf(&builder, ", UnderlyingInstrID=%+v", d.UnderlyingInstrID)
	fmt.Fprintf(&builder, ", ProductID=%+v", d.ProductID)
	fmt.Fprintf(&builder, ", OffsetType=%+v", d.OffsetType)
	fmt.Fprintf(&builder, ", Volume=%+v", d.Volume)
	fmt.Fprintf(&builder, ", IsOffset=%+v", d.IsOffset)
	fmt.Fprintf(&builder, ", RequestID=%+v", d.RequestID)
	fmt.Fprintf(&builder, ", UserID=%+v", d.UserID)
	fmt.Fprintf(&builder, ", ExchangeID=%+v", d.ExchangeID)
	fmt.Fprintf(&builder, ", IPAddress=%+v", d.IPAddress)
	fmt.Fprintf(&builder, ", MacAddress=%+v", d.MacAddress)

	builder.WriteByte('}')

	return builder.String()
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

func (d CThostFtdcOffsetSettingField) Type() string { return "CThostFtdcOffsetSettingField" }

func (d CThostFtdcOffsetSettingField) String() string {
	var builder strings.Builder
	builder.Grow(733)

	builder.WriteString("CThostFtdcOffsetSettingField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", InvestorID=%+v", d.InvestorID)
	fmt.Fprintf(&builder, ", InstrumentID=%+v", d.InstrumentID)
	fmt.Fprintf(&builder, ", UnderlyingInstrID=%+v", d.UnderlyingInstrID)
	fmt.Fprintf(&builder, ", ProductID=%+v", d.ProductID)
	fmt.Fprintf(&builder, ", OffsetType=%+v", d.OffsetType)
	fmt.Fprintf(&builder, ", Volume=%+v", d.Volume)
	fmt.Fprintf(&builder, ", IsOffset=%+v", d.IsOffset)
	fmt.Fprintf(&builder, ", RequestID=%+v", d.RequestID)
	fmt.Fprintf(&builder, ", UserID=%+v", d.UserID)
	fmt.Fprintf(&builder, ", ExchangeID=%+v", d.ExchangeID)
	fmt.Fprintf(&builder, ", IPAddress=%+v", d.IPAddress)
	fmt.Fprintf(&builder, ", MacAddress=%+v", d.MacAddress)
	fmt.Fprintf(&builder, ", ExchangeInstID=%+v", d.ExchangeInstID)
	fmt.Fprintf(&builder, ", ExchangeSerialNo=%+v", d.ExchangeSerialNo)
	fmt.Fprintf(&builder, ", ExchangeProductID=%+v", d.ExchangeProductID)
	fmt.Fprintf(&builder, ", ParticipantID=%+v", d.ParticipantID)
	fmt.Fprintf(&builder, ", ClientID=%+v", d.ClientID)
	fmt.Fprintf(&builder, ", TraderID=%+v", d.TraderID)
	fmt.Fprintf(&builder, ", InstallID=%+v", d.InstallID)
	fmt.Fprintf(&builder, ", OrderSubmitStatus=%+v", d.OrderSubmitStatus)
	fmt.Fprintf(&builder, ", TradingDay=%+v", d.TradingDay)
	fmt.Fprintf(&builder, ", SettlementID=%+v", d.SettlementID)
	fmt.Fprintf(&builder, ", InsertDate=%+v", d.InsertDate)
	fmt.Fprintf(&builder, ", InsertTime=%+v", d.InsertTime)
	fmt.Fprintf(&builder, ", CancelTime=%+v", d.CancelTime)
	fmt.Fprintf(&builder, ", ExecResult=%+v", d.ExecResult)
	fmt.Fprintf(&builder, ", SequenceNo=%+v", d.SequenceNo)
	fmt.Fprintf(&builder, ", FrontID=%+v", d.FrontID)
	fmt.Fprintf(&builder, ", SessionID=%+v", d.SessionID)
	fmt.Fprintf(&builder, ", StatusMsg=%+v", d.StatusMsg)
	fmt.Fprintf(&builder, ", ActiveUserID=%+v", d.ActiveUserID)
	fmt.Fprintf(&builder, ", BrokerOffsetSettingSeq=%+v", d.BrokerOffsetSettingSeq)
	fmt.Fprintf(&builder, ", ApplySrc=%+v", d.ApplySrc)

	builder.WriteByte('}')

	return builder.String()
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

func (d CThostFtdcCancelOffsetSettingField) Type() string {
	return "CThostFtdcCancelOffsetSettingField"
}

func (d CThostFtdcCancelOffsetSettingField) String() string {
	var builder strings.Builder
	builder.Grow(552)

	builder.WriteString("CThostFtdcCancelOffsetSettingField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", InvestorID=%+v", d.InvestorID)
	fmt.Fprintf(&builder, ", InstrumentID=%+v", d.InstrumentID)
	fmt.Fprintf(&builder, ", UnderlyingInstrID=%+v", d.UnderlyingInstrID)
	fmt.Fprintf(&builder, ", ProductID=%+v", d.ProductID)
	fmt.Fprintf(&builder, ", OffsetType=%+v", d.OffsetType)
	fmt.Fprintf(&builder, ", Volume=%+v", d.Volume)
	fmt.Fprintf(&builder, ", IsOffset=%+v", d.IsOffset)
	fmt.Fprintf(&builder, ", RequestID=%+v", d.RequestID)
	fmt.Fprintf(&builder, ", UserID=%+v", d.UserID)
	fmt.Fprintf(&builder, ", ExchangeID=%+v", d.ExchangeID)
	fmt.Fprintf(&builder, ", IPAddress=%+v", d.IPAddress)
	fmt.Fprintf(&builder, ", MacAddress=%+v", d.MacAddress)
	fmt.Fprintf(&builder, ", ExchangeInstID=%+v", d.ExchangeInstID)
	fmt.Fprintf(&builder, ", ExchangeSerialNo=%+v", d.ExchangeSerialNo)
	fmt.Fprintf(&builder, ", ExchangeProductID=%+v", d.ExchangeProductID)
	fmt.Fprintf(&builder, ", TraderID=%+v", d.TraderID)
	fmt.Fprintf(&builder, ", InstallID=%+v", d.InstallID)
	fmt.Fprintf(&builder, ", ParticipantID=%+v", d.ParticipantID)
	fmt.Fprintf(&builder, ", ClientID=%+v", d.ClientID)
	fmt.Fprintf(&builder, ", OrderActionStatus=%+v", d.OrderActionStatus)
	fmt.Fprintf(&builder, ", StatusMsg=%+v", d.StatusMsg)
	fmt.Fprintf(&builder, ", ActionLocalID=%+v", d.ActionLocalID)
	fmt.Fprintf(&builder, ", ActionDate=%+v", d.ActionDate)
	fmt.Fprintf(&builder, ", ActionTime=%+v", d.ActionTime)

	builder.WriteByte('}')

	return builder.String()
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

func (d CThostFtdcQryOffsetSettingField) Type() string { return "CThostFtdcQryOffsetSettingField" }

func (d CThostFtdcQryOffsetSettingField) String() string {
	var builder strings.Builder
	builder.Grow(108)

	builder.WriteString("CThostFtdcQryOffsetSettingField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", InvestorID=%+v", d.InvestorID)
	fmt.Fprintf(&builder, ", ProductID=%+v", d.ProductID)
	fmt.Fprintf(&builder, ", OffsetType=%+v", d.OffsetType)

	builder.WriteByte('}')

	return builder.String()
}

// 服务地址和AppID的关系
type CThostFtdcAddrAppIDRelationField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 服务地址
	Address types.TThostFtdcIpAddrType
	// 交易中心代码
	DRIdentityID types.TThostFtdcDRIdentityIDType
	// App代码
	AppID types.TThostFtdcAppIDType
}

func (d CThostFtdcAddrAppIDRelationField) Type() string { return "CThostFtdcAddrAppIDRelationField" }

func (d CThostFtdcAddrAppIDRelationField) String() string {
	var builder strings.Builder
	builder.Grow(104)

	builder.WriteString("CThostFtdcAddrAppIDRelationField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", Address=%+v", d.Address)
	fmt.Fprintf(&builder, ", DRIdentityID=%+v", d.DRIdentityID)
	fmt.Fprintf(&builder, ", AppID=%+v", d.AppID)

	builder.WriteByte('}')

	return builder.String()
}

// 服务地址和AppID的关系查询
type CThostFtdcQryAddrAppIDRelationField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
}

func (d CThostFtdcQryAddrAppIDRelationField) Type() string {
	return "CThostFtdcQryAddrAppIDRelationField"
}

func (d CThostFtdcQryAddrAppIDRelationField) String() string {
	var builder strings.Builder
	builder.Grow(53)

	builder.WriteString("CThostFtdcQryAddrAppIDRelationField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)

	builder.WriteByte('}')

	return builder.String()
}

// 微信小程序等用户系统信息
type CThostFtdcWechatUserSystemInfoField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 用户代码
	UserID types.TThostFtdcUserIDType
	// 微信小程序等用户端系统内部信息长度
	WechatCltSysInfoLen types.TThostFtdcSystemInfoLenType
	// 微信小程序等用户端系统内部信息
	WechatCltSysInfo types.TThostFtdcClientSystemInfoType
	// 终端IP端口
	ClientIPPort types.TThostFtdcIPPortType
	// 登录成功时间
	ClientLoginTime types.TThostFtdcTimeType
	// App代码
	ClientAppID types.TThostFtdcAppIDType
	// 用户公网IP
	ClientPublicIP types.TThostFtdcIPAddressType
	// 客户登录备注2
	ClientLoginRemark types.TThostFtdcClientLoginRemarkType
}

func (d CThostFtdcWechatUserSystemInfoField) Type() string {
	return "CThostFtdcWechatUserSystemInfoField"
}

func (d CThostFtdcWechatUserSystemInfoField) String() string {
	var builder strings.Builder
	builder.Grow(243)

	builder.WriteString("CThostFtdcWechatUserSystemInfoField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", UserID=%+v", d.UserID)
	fmt.Fprintf(&builder, ", WechatCltSysInfoLen=%+v", d.WechatCltSysInfoLen)
	fmt.Fprintf(&builder, ", WechatCltSysInfo=%+v", d.WechatCltSysInfo)
	fmt.Fprintf(&builder, ", ClientIPPort=%+v", d.ClientIPPort)
	fmt.Fprintf(&builder, ", ClientLoginTime=%+v", d.ClientLoginTime)
	fmt.Fprintf(&builder, ", ClientAppID=%+v", d.ClientAppID)
	fmt.Fprintf(&builder, ", ClientPublicIP=%+v", d.ClientPublicIP)
	fmt.Fprintf(&builder, ", ClientLoginRemark=%+v", d.ClientLoginRemark)

	builder.WriteByte('}')

	return builder.String()
}

// 投资者预留信息
type CThostFtdcInvestorReserveInfoField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 用户代码
	UserID types.TThostFtdcUserIDType
	// 预留信息
	ReserveInfo types.TThostFtdcReserveInfoType
}

func (d CThostFtdcInvestorReserveInfoField) Type() string {
	return "CThostFtdcInvestorReserveInfoField"
}

func (d CThostFtdcInvestorReserveInfoField) String() string {
	var builder strings.Builder
	builder.Grow(89)

	builder.WriteString("CThostFtdcInvestorReserveInfoField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", UserID=%+v", d.UserID)
	fmt.Fprintf(&builder, ", ReserveInfo=%+v", d.ReserveInfo)

	builder.WriteByte('}')

	return builder.String()
}

// 查询组织架构投资者对应关系
type CThostFtdcQryInvestorDepartmentFlatField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
}

func (d CThostFtdcQryInvestorDepartmentFlatField) Type() string {
	return "CThostFtdcQryInvestorDepartmentFlatField"
}

func (d CThostFtdcQryInvestorDepartmentFlatField) String() string {
	var builder strings.Builder
	builder.Grow(58)

	builder.WriteString("CThostFtdcQryInvestorDepartmentFlatField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)

	builder.WriteByte('}')

	return builder.String()
}

// 组织架构投资者对应关系
type CThostFtdcInvestorDepartmentFlatField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	// 组织架构代码
	DepartmentID types.TThostFtdcInvestorIDType
}

func (d CThostFtdcInvestorDepartmentFlatField) Type() string {
	return "CThostFtdcInvestorDepartmentFlatField"
}

func (d CThostFtdcInvestorDepartmentFlatField) String() string {
	var builder strings.Builder
	builder.Grow(97)

	builder.WriteString("CThostFtdcInvestorDepartmentFlatField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", InvestorID=%+v", d.InvestorID)
	fmt.Fprintf(&builder, ", DepartmentID=%+v", d.DepartmentID)

	builder.WriteByte('}')

	return builder.String()
}

// 查询操作员组织架构关系
type CThostFtdcQryDepartmentUserField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
}

func (d CThostFtdcQryDepartmentUserField) Type() string { return "CThostFtdcQryDepartmentUserField" }

func (d CThostFtdcQryDepartmentUserField) String() string {
	var builder strings.Builder
	builder.Grow(50)

	builder.WriteString("CThostFtdcQryDepartmentUserField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)

	builder.WriteByte('}')

	return builder.String()
}

// App客户端认证码
type CThostFtdcAppAuthenticationCodeField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// App代码
	AppID types.TThostFtdcAppIDType
	// 认证码
	AuthCode types.TThostFtdcAuthCodeType
	// 旧认证码
	PreAuthCode types.TThostFtdcAuthCodeType
	// App类型
	AppType types.TThostFtdcAppTypeType
}

func (d CThostFtdcAppAuthenticationCodeField) Type() string {
	return "CThostFtdcAppAuthenticationCodeField"
}

func (d CThostFtdcAppAuthenticationCodeField) String() string {
	var builder strings.Builder
	builder.Grow(125)

	builder.WriteString("CThostFtdcAppAuthenticationCodeField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", AppID=%+v", d.AppID)
	fmt.Fprintf(&builder, ", AuthCode=%+v", d.AuthCode)
	fmt.Fprintf(&builder, ", PreAuthCode=%+v", d.PreAuthCode)
	fmt.Fprintf(&builder, ", AppType=%+v", d.AppType)

	builder.WriteByte('}')

	return builder.String()
}

// 客户中心权限豁免
type CThostFtdcUserDRIBypassField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 用户代码
	UserID types.TThostFtdcUserIDType
	// 交易中心代码
	DRIdentityID types.TThostFtdcDRIdentityIDType
}

func (d CThostFtdcUserDRIBypassField) Type() string { return "CThostFtdcUserDRIBypassField" }

func (d CThostFtdcUserDRIBypassField) String() string {
	var builder strings.Builder
	builder.Grow(84)

	builder.WriteString("CThostFtdcUserDRIBypassField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", UserID=%+v", d.UserID)
	fmt.Fprintf(&builder, ", DRIdentityID=%+v", d.DRIdentityID)

	builder.WriteByte('}')

	return builder.String()
}

// 申请短信验证码请求
type CThostFtdcReqGenSMSCodeField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 用户代码
	UserID types.TThostFtdcUserIDType
	// 手机号
	Mobile types.TThostFtdcSMSPhoneType
}

func (d CThostFtdcReqGenSMSCodeField) Type() string { return "CThostFtdcReqGenSMSCodeField" }

func (d CThostFtdcReqGenSMSCodeField) String() string {
	var builder strings.Builder
	builder.Grow(78)

	builder.WriteString("CThostFtdcReqGenSMSCodeField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", UserID=%+v", d.UserID)
	fmt.Fprintf(&builder, ", Mobile=%+v", d.Mobile)

	builder.WriteByte('}')

	return builder.String()
}

// 申请短信验证码响应
type CThostFtdcRspGenSMSCodeField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 用户代码
	UserID types.TThostFtdcUserIDType
	// 生成时间
	GenTime types.TThostFtdcTimeType
}

func (d CThostFtdcRspGenSMSCodeField) Type() string { return "CThostFtdcRspGenSMSCodeField" }

func (d CThostFtdcRspGenSMSCodeField) String() string {
	var builder strings.Builder
	builder.Grow(79)

	builder.WriteString("CThostFtdcRspGenSMSCodeField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", UserID=%+v", d.UserID)
	fmt.Fprintf(&builder, ", GenTime=%+v", d.GenTime)

	builder.WriteByte('}')

	return builder.String()
}

// 短信验证信息通知
type CThostFtdcSMSVerifyInfoFromSecField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 经纪公司简称
	BrokerAbbr types.TThostFtdcBrokerAbbrType
	// 用户代码
	UserID types.TThostFtdcUserIDType
	// 手机号
	Mobile types.TThostFtdcSMSPhoneType
	// 短信验证码
	SMSCode types.TThostFtdcSMSCodeType
	// 验证码创建日期
	CreateDate types.TThostFtdcDateType
	// 验证码创建时间
	CreateTime types.TThostFtdcTimeType
	// 验证码是否被使用过
	IsUsed types.TThostFtdcBoolType
	// 次席的交易中心代码
	FromSec types.TThostFtdcDRIdentityIDType
}

func (d CThostFtdcSMSVerifyInfoFromSecField) Type() string {
	return "CThostFtdcSMSVerifyInfoFromSecField"
}

func (d CThostFtdcSMSVerifyInfoFromSecField) String() string {
	var builder strings.Builder
	builder.Grow(195)

	builder.WriteString("CThostFtdcSMSVerifyInfoFromSecField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", BrokerAbbr=%+v", d.BrokerAbbr)
	fmt.Fprintf(&builder, ", UserID=%+v", d.UserID)
	fmt.Fprintf(&builder, ", Mobile=%+v", d.Mobile)
	fmt.Fprintf(&builder, ", SMSCode=%+v", d.SMSCode)
	fmt.Fprintf(&builder, ", CreateDate=%+v", d.CreateDate)
	fmt.Fprintf(&builder, ", CreateTime=%+v", d.CreateTime)
	fmt.Fprintf(&builder, ", IsUsed=%+v", d.IsUsed)
	fmt.Fprintf(&builder, ", FromSec=%+v", d.FromSec)

	builder.WriteByte('}')

	return builder.String()
}

// 登录验证设置
type CThostFtdcSMSVerifyConfigField struct {
	// 用户代码
	UserID types.TThostFtdcUserIDType
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 手机号
	Mobile types.TThostFtdcSMSPhoneType
	// 是否启用短信验证
	UseSMSVerify types.TThostFtdcBoolType
}

func (d CThostFtdcSMSVerifyConfigField) Type() string { return "CThostFtdcSMSVerifyConfigField" }

func (d CThostFtdcSMSVerifyConfigField) String() string {
	var builder strings.Builder
	builder.Grow(102)

	builder.WriteString("CThostFtdcSMSVerifyConfigField{")
	fmt.Fprintf(&builder, "UserID=%+v", d.UserID)
	fmt.Fprintf(&builder, ", BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", Mobile=%+v", d.Mobile)
	fmt.Fprintf(&builder, ", UseSMSVerify=%+v", d.UseSMSVerify)

	builder.WriteByte('}')

	return builder.String()
}

// 短信验证信息通知
type CThostFtdcSMSVerifyInfoField struct {
	// 验证码创建时间
	CreateTime types.TThostFtdcTimeType
	// 手机号
	Mobile types.TThostFtdcSMSPhoneType
	// 短信验证信息内容
	SMSContent types.TThostFtdcSMSContentType
}

func (d CThostFtdcSMSVerifyInfoField) Type() string { return "CThostFtdcSMSVerifyInfoField" }

func (d CThostFtdcSMSVerifyInfoField) String() string {
	var builder strings.Builder
	builder.Grow(84)

	builder.WriteString("CThostFtdcSMSVerifyInfoField{")
	fmt.Fprintf(&builder, "CreateTime=%+v", d.CreateTime)
	fmt.Fprintf(&builder, ", Mobile=%+v", d.Mobile)
	fmt.Fprintf(&builder, ", SMSContent=%+v", d.SMSContent)

	builder.WriteByte('}')

	return builder.String()
}

// 套利确认输入基本信息
type CThostFtdcInputSpdApplyField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 用户代码
	UserID types.TThostFtdcUserIDType
	// 投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	// 交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	// 合约代码
	FirstLegInstrumentID types.TThostFtdcInstrumentIDType
	// 合约代码
	SecondLegInstrumentID types.TThostFtdcInstrumentIDType
	// 数量
	Volume types.TThostFtdcVolumeType
	// 买卖方向
	Direction types.TThostFtdcDirectionType
	// 组合定单类型
	CmbType types.TThostFtdcCmbTypeType
	// 请求编号
	RequestID types.TThostFtdcRequestIDType
	// 报单引用
	OrderRef types.TThostFtdcOrderRefType
	// IP地址
	IPAddress types.TThostFtdcIPAddressType
	// Mac地址
	MacAddress types.TThostFtdcMacAddressType
}

func (d CThostFtdcInputSpdApplyField) Type() string { return "CThostFtdcInputSpdApplyField" }

func (d CThostFtdcInputSpdApplyField) String() string {
	var builder strings.Builder
	builder.Grow(291)

	builder.WriteString("CThostFtdcInputSpdApplyField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", UserID=%+v", d.UserID)
	fmt.Fprintf(&builder, ", InvestorID=%+v", d.InvestorID)
	fmt.Fprintf(&builder, ", ExchangeID=%+v", d.ExchangeID)
	fmt.Fprintf(&builder, ", FirstLegInstrumentID=%+v", d.FirstLegInstrumentID)
	fmt.Fprintf(&builder, ", SecondLegInstrumentID=%+v", d.SecondLegInstrumentID)
	fmt.Fprintf(&builder, ", Volume=%+v", d.Volume)
	fmt.Fprintf(&builder, ", Direction=%+v", d.Direction)
	fmt.Fprintf(&builder, ", CmbType=%+v", d.CmbType)
	fmt.Fprintf(&builder, ", RequestID=%+v", d.RequestID)
	fmt.Fprintf(&builder, ", OrderRef=%+v", d.OrderRef)
	fmt.Fprintf(&builder, ", IPAddress=%+v", d.IPAddress)
	fmt.Fprintf(&builder, ", MacAddress=%+v", d.MacAddress)

	builder.WriteByte('}')

	return builder.String()
}

// 套保确认输入基本信息
type CThostFtdcInputHedgeCfmField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 用户代码
	UserID types.TThostFtdcUserIDType
	// 投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	// 交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	// 合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
	// 数量
	Volume types.TThostFtdcVolumeType
	// 买卖方向
	Direction types.TThostFtdcDirectionType
	// 请求编号
	RequestID types.TThostFtdcRequestIDType
	// 报单引用
	OrderRef types.TThostFtdcOrderRefType
	// IP地址
	IPAddress types.TThostFtdcIPAddressType
	// Mac地址
	MacAddress types.TThostFtdcMacAddressType
}

func (d CThostFtdcInputHedgeCfmField) Type() string { return "CThostFtdcInputHedgeCfmField" }

func (d CThostFtdcInputHedgeCfmField) String() string {
	var builder strings.Builder
	builder.Grow(235)

	builder.WriteString("CThostFtdcInputHedgeCfmField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", UserID=%+v", d.UserID)
	fmt.Fprintf(&builder, ", InvestorID=%+v", d.InvestorID)
	fmt.Fprintf(&builder, ", ExchangeID=%+v", d.ExchangeID)
	fmt.Fprintf(&builder, ", InstrumentID=%+v", d.InstrumentID)
	fmt.Fprintf(&builder, ", Volume=%+v", d.Volume)
	fmt.Fprintf(&builder, ", Direction=%+v", d.Direction)
	fmt.Fprintf(&builder, ", RequestID=%+v", d.RequestID)
	fmt.Fprintf(&builder, ", OrderRef=%+v", d.OrderRef)
	fmt.Fprintf(&builder, ", IPAddress=%+v", d.IPAddress)
	fmt.Fprintf(&builder, ", MacAddress=%+v", d.MacAddress)

	builder.WriteByte('}')

	return builder.String()
}

// 套利申请回报
type CThostFtdcSpdApplyField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	// 合约代码
	FirstLegInstrumentID types.TThostFtdcInstrumentIDType
	// 合约代码
	SecondLegInstrumentID types.TThostFtdcInstrumentIDType
	// 用户代码
	UserID types.TThostFtdcUserIDType
	// 数量
	Volume types.TThostFtdcVolumeType
	// 买卖方向
	Direction types.TThostFtdcDirectionType
	// 请求编号
	RequestID types.TThostFtdcRequestIDType
	// 前置编号
	FrontID types.TThostFtdcFrontIDType
	// 会话编号
	SessionID types.TThostFtdcSessionIDType
	// 报单引用
	OrderRef types.TThostFtdcOrderRefType
	// 操作用户代码
	ActiveUserID types.TThostFtdcUserIDType
	// 经纪公司报单编号
	BrokerOrderSeq types.TThostFtdcSequenceNoType
	// 报单编号
	OrderSysID types.TThostFtdcOrderSysIDType
	// 申请状态
	ApplyStatus types.TThostFtdcApplyStatusType
	// 序号
	SequenceNo types.TThostFtdcSequenceNoType
	// 报单日期
	InsertDate types.TThostFtdcDateType
	// 委托时间
	InsertTime types.TThostFtdcTimeType
	// 撤销时间
	CancelTime types.TThostFtdcTimeType
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
	// IP地址
	IPAddress types.TThostFtdcIPAddressType
	// Mac地址
	MacAddress types.TThostFtdcMacAddressType
	// 组合定单类型
	CmbType types.TThostFtdcCmbTypeType
	// 状态信息
	StatusMsg types.TThostFtdcErrorMsgType
}

func (d CThostFtdcSpdApplyField) Type() string { return "CThostFtdcSpdApplyField" }

func (d CThostFtdcSpdApplyField) String() string {
	var builder strings.Builder
	builder.Grow(725)

	builder.WriteString("CThostFtdcSpdApplyField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", InvestorID=%+v", d.InvestorID)
	fmt.Fprintf(&builder, ", FirstLegInstrumentID=%+v", d.FirstLegInstrumentID)
	fmt.Fprintf(&builder, ", SecondLegInstrumentID=%+v", d.SecondLegInstrumentID)
	fmt.Fprintf(&builder, ", UserID=%+v", d.UserID)
	fmt.Fprintf(&builder, ", Volume=%+v", d.Volume)
	fmt.Fprintf(&builder, ", Direction=%+v", d.Direction)
	fmt.Fprintf(&builder, ", RequestID=%+v", d.RequestID)
	fmt.Fprintf(&builder, ", FrontID=%+v", d.FrontID)
	fmt.Fprintf(&builder, ", SessionID=%+v", d.SessionID)
	fmt.Fprintf(&builder, ", OrderRef=%+v", d.OrderRef)
	fmt.Fprintf(&builder, ", ActiveUserID=%+v", d.ActiveUserID)
	fmt.Fprintf(&builder, ", BrokerOrderSeq=%+v", d.BrokerOrderSeq)
	fmt.Fprintf(&builder, ", OrderSysID=%+v", d.OrderSysID)
	fmt.Fprintf(&builder, ", ApplyStatus=%+v", d.ApplyStatus)
	fmt.Fprintf(&builder, ", SequenceNo=%+v", d.SequenceNo)
	fmt.Fprintf(&builder, ", InsertDate=%+v", d.InsertDate)
	fmt.Fprintf(&builder, ", InsertTime=%+v", d.InsertTime)
	fmt.Fprintf(&builder, ", CancelTime=%+v", d.CancelTime)
	fmt.Fprintf(&builder, ", OrderLocalID=%+v", d.OrderLocalID)
	fmt.Fprintf(&builder, ", ExchangeID=%+v", d.ExchangeID)
	fmt.Fprintf(&builder, ", ParticipantID=%+v", d.ParticipantID)
	fmt.Fprintf(&builder, ", ClientID=%+v", d.ClientID)
	fmt.Fprintf(&builder, ", ExchangeInstID=%+v", d.ExchangeInstID)
	fmt.Fprintf(&builder, ", TraderID=%+v", d.TraderID)
	fmt.Fprintf(&builder, ", InstallID=%+v", d.InstallID)
	fmt.Fprintf(&builder, ", OrderSubmitStatus=%+v", d.OrderSubmitStatus)
	fmt.Fprintf(&builder, ", NotifySequence=%+v", d.NotifySequence)
	fmt.Fprintf(&builder, ", TradingDay=%+v", d.TradingDay)
	fmt.Fprintf(&builder, ", SettlementID=%+v", d.SettlementID)
	fmt.Fprintf(&builder, ", IPAddress=%+v", d.IPAddress)
	fmt.Fprintf(&builder, ", MacAddress=%+v", d.MacAddress)
	fmt.Fprintf(&builder, ", CmbType=%+v", d.CmbType)
	fmt.Fprintf(&builder, ", StatusMsg=%+v", d.StatusMsg)

	builder.WriteByte('}')

	return builder.String()
}

// 套保申请回报
type CThostFtdcHedgeCfmField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	// 合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
	// 用户代码
	UserID types.TThostFtdcUserIDType
	// 数量
	Volume types.TThostFtdcVolumeType
	// 买卖方向
	Direction types.TThostFtdcDirectionType
	// 请求编号
	RequestID types.TThostFtdcRequestIDType
	// 前置编号
	FrontID types.TThostFtdcFrontIDType
	// 会话编号
	SessionID types.TThostFtdcSessionIDType
	// 报单引用
	OrderRef types.TThostFtdcOrderRefType
	// 操作用户代码
	ActiveUserID types.TThostFtdcUserIDType
	// 经纪公司报单编号
	BrokerOrderSeq types.TThostFtdcSequenceNoType
	// 报单编号
	OrderSysID types.TThostFtdcOrderSysIDType
	// 申请状态
	ApplyStatus types.TThostFtdcApplyStatusType
	// 序号
	SequenceNo types.TThostFtdcSequenceNoType
	// 成功处理数量
	DealVolume types.TThostFtdcVolumeType
	// 报单日期
	InsertDate types.TThostFtdcDateType
	// 委托时间
	InsertTime types.TThostFtdcTimeType
	// 撤销时间
	CancelTime types.TThostFtdcTimeType
	// 日期
	ReqDate types.TThostFtdcDateType
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
	// 状态信息
	StatusMsg types.TThostFtdcErrorMsgType
	// IP地址
	IPAddress types.TThostFtdcIPAddressType
	// Mac地址
	MacAddress types.TThostFtdcMacAddressType
}

func (d CThostFtdcHedgeCfmField) Type() string { return "CThostFtdcHedgeCfmField" }

func (d CThostFtdcHedgeCfmField) String() string {
	var builder strings.Builder
	builder.Grow(706)

	builder.WriteString("CThostFtdcHedgeCfmField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", InvestorID=%+v", d.InvestorID)
	fmt.Fprintf(&builder, ", InstrumentID=%+v", d.InstrumentID)
	fmt.Fprintf(&builder, ", UserID=%+v", d.UserID)
	fmt.Fprintf(&builder, ", Volume=%+v", d.Volume)
	fmt.Fprintf(&builder, ", Direction=%+v", d.Direction)
	fmt.Fprintf(&builder, ", RequestID=%+v", d.RequestID)
	fmt.Fprintf(&builder, ", FrontID=%+v", d.FrontID)
	fmt.Fprintf(&builder, ", SessionID=%+v", d.SessionID)
	fmt.Fprintf(&builder, ", OrderRef=%+v", d.OrderRef)
	fmt.Fprintf(&builder, ", ActiveUserID=%+v", d.ActiveUserID)
	fmt.Fprintf(&builder, ", BrokerOrderSeq=%+v", d.BrokerOrderSeq)
	fmt.Fprintf(&builder, ", OrderSysID=%+v", d.OrderSysID)
	fmt.Fprintf(&builder, ", ApplyStatus=%+v", d.ApplyStatus)
	fmt.Fprintf(&builder, ", SequenceNo=%+v", d.SequenceNo)
	fmt.Fprintf(&builder, ", DealVolume=%+v", d.DealVolume)
	fmt.Fprintf(&builder, ", InsertDate=%+v", d.InsertDate)
	fmt.Fprintf(&builder, ", InsertTime=%+v", d.InsertTime)
	fmt.Fprintf(&builder, ", CancelTime=%+v", d.CancelTime)
	fmt.Fprintf(&builder, ", ReqDate=%+v", d.ReqDate)
	fmt.Fprintf(&builder, ", OrderLocalID=%+v", d.OrderLocalID)
	fmt.Fprintf(&builder, ", ExchangeID=%+v", d.ExchangeID)
	fmt.Fprintf(&builder, ", ParticipantID=%+v", d.ParticipantID)
	fmt.Fprintf(&builder, ", ClientID=%+v", d.ClientID)
	fmt.Fprintf(&builder, ", ExchangeInstID=%+v", d.ExchangeInstID)
	fmt.Fprintf(&builder, ", TraderID=%+v", d.TraderID)
	fmt.Fprintf(&builder, ", InstallID=%+v", d.InstallID)
	fmt.Fprintf(&builder, ", OrderSubmitStatus=%+v", d.OrderSubmitStatus)
	fmt.Fprintf(&builder, ", NotifySequence=%+v", d.NotifySequence)
	fmt.Fprintf(&builder, ", TradingDay=%+v", d.TradingDay)
	fmt.Fprintf(&builder, ", SettlementID=%+v", d.SettlementID)
	fmt.Fprintf(&builder, ", StatusMsg=%+v", d.StatusMsg)
	fmt.Fprintf(&builder, ", IPAddress=%+v", d.IPAddress)
	fmt.Fprintf(&builder, ", MacAddress=%+v", d.MacAddress)

	builder.WriteByte('}')

	return builder.String()
}

// 套利套保申请查询
type CThostFtdcQrySpdApplyField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	// 交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	// 报单编号
	OrderSysID types.TThostFtdcOrderSysIDType
	// 第一腿合约编码
	FirstLegInstrumentID types.TThostFtdcExchangeInstIDType
	// 第二腿合约编码
	SecondLegInstrumentID types.TThostFtdcExchangeInstIDType
}

func (d CThostFtdcQrySpdApplyField) Type() string { return "CThostFtdcQrySpdApplyField" }

func (d CThostFtdcQrySpdApplyField) String() string {
	var builder strings.Builder
	builder.Grow(165)

	builder.WriteString("CThostFtdcQrySpdApplyField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", InvestorID=%+v", d.InvestorID)
	fmt.Fprintf(&builder, ", ExchangeID=%+v", d.ExchangeID)
	fmt.Fprintf(&builder, ", OrderSysID=%+v", d.OrderSysID)
	fmt.Fprintf(&builder, ", FirstLegInstrumentID=%+v", d.FirstLegInstrumentID)
	fmt.Fprintf(&builder, ", SecondLegInstrumentID=%+v", d.SecondLegInstrumentID)

	builder.WriteByte('}')

	return builder.String()
}

// 套利套保申请查询
type CThostFtdcQryHedgeCfmField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	// 交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	// 报单编号
	OrderSysID types.TThostFtdcOrderSysIDType
	// 合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
}

func (d CThostFtdcQryHedgeCfmField) Type() string { return "CThostFtdcQryHedgeCfmField" }

func (d CThostFtdcQryHedgeCfmField) String() string {
	var builder strings.Builder
	builder.Grow(126)

	builder.WriteString("CThostFtdcQryHedgeCfmField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", InvestorID=%+v", d.InvestorID)
	fmt.Fprintf(&builder, ", ExchangeID=%+v", d.ExchangeID)
	fmt.Fprintf(&builder, ", OrderSysID=%+v", d.OrderSysID)
	fmt.Fprintf(&builder, ", InstrumentID=%+v", d.InstrumentID)

	builder.WriteByte('}')

	return builder.String()
}

// 套利申请撤销
type CThostFtdcInputSpdApplyActionField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 用户代码
	UserID types.TThostFtdcUserIDType
	// 投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	// 交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	// 合同编号
	OrderSysID types.TThostFtdcOrderSysIDType
	// 报单引用
	OrderRef types.TThostFtdcOrderRefType
	// 前置编号
	FrontID types.TThostFtdcFrontIDType
	// 会话编号
	SessionID types.TThostFtdcSessionIDType
	// 请求编号
	RequestID types.TThostFtdcRequestIDType
	// IP地址
	IPAddress types.TThostFtdcIPAddressType
	// Mac地址
	MacAddress types.TThostFtdcMacAddressType
}

func (d CThostFtdcInputSpdApplyActionField) Type() string {
	return "CThostFtdcInputSpdApplyActionField"
}

func (d CThostFtdcInputSpdApplyActionField) String() string {
	var builder strings.Builder
	builder.Grow(240)

	builder.WriteString("CThostFtdcInputSpdApplyActionField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", UserID=%+v", d.UserID)
	fmt.Fprintf(&builder, ", InvestorID=%+v", d.InvestorID)
	fmt.Fprintf(&builder, ", ExchangeID=%+v", d.ExchangeID)
	fmt.Fprintf(&builder, ", OrderSysID=%+v", d.OrderSysID)
	fmt.Fprintf(&builder, ", OrderRef=%+v", d.OrderRef)
	fmt.Fprintf(&builder, ", FrontID=%+v", d.FrontID)
	fmt.Fprintf(&builder, ", SessionID=%+v", d.SessionID)
	fmt.Fprintf(&builder, ", RequestID=%+v", d.RequestID)
	fmt.Fprintf(&builder, ", IPAddress=%+v", d.IPAddress)
	fmt.Fprintf(&builder, ", MacAddress=%+v", d.MacAddress)

	builder.WriteByte('}')

	return builder.String()
}

// 套保申请撤销
type CThostFtdcInputHedgeCfmActionField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 用户代码
	UserID types.TThostFtdcUserIDType
	// 投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	// 交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	// 合同编号
	OrderSysID types.TThostFtdcOrderSysIDType
	// 报单引用
	OrderRef types.TThostFtdcOrderRefType
	// 前置编号
	FrontID types.TThostFtdcFrontIDType
	// 会话编号
	SessionID types.TThostFtdcSessionIDType
	// 请求编号
	RequestID types.TThostFtdcRequestIDType
	// IP地址
	IPAddress types.TThostFtdcIPAddressType
	// Mac地址
	MacAddress types.TThostFtdcMacAddressType
}

func (d CThostFtdcInputHedgeCfmActionField) Type() string {
	return "CThostFtdcInputHedgeCfmActionField"
}

func (d CThostFtdcInputHedgeCfmActionField) String() string {
	var builder strings.Builder
	builder.Grow(240)

	builder.WriteString("CThostFtdcInputHedgeCfmActionField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", UserID=%+v", d.UserID)
	fmt.Fprintf(&builder, ", InvestorID=%+v", d.InvestorID)
	fmt.Fprintf(&builder, ", ExchangeID=%+v", d.ExchangeID)
	fmt.Fprintf(&builder, ", OrderSysID=%+v", d.OrderSysID)
	fmt.Fprintf(&builder, ", OrderRef=%+v", d.OrderRef)
	fmt.Fprintf(&builder, ", FrontID=%+v", d.FrontID)
	fmt.Fprintf(&builder, ", SessionID=%+v", d.SessionID)
	fmt.Fprintf(&builder, ", RequestID=%+v", d.RequestID)
	fmt.Fprintf(&builder, ", IPAddress=%+v", d.IPAddress)
	fmt.Fprintf(&builder, ", MacAddress=%+v", d.MacAddress)

	builder.WriteByte('}')

	return builder.String()
}

// 套利申请撤销回报
type CThostFtdcSpdApplyActionField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID types.TThostFtdcInvestorIDType
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
	// 报单操作状态
	OrderActionStatus types.TThostFtdcOrderActionStatusType
	// 用户代码
	UserID types.TThostFtdcUserIDType
	// 交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	// 合同编号
	OrderSysID types.TThostFtdcOrderSysIDType
	// 请求编号
	RequestID types.TThostFtdcRequestIDType
	// 状态信息
	StatusMsg types.TThostFtdcErrorMsgType
	// 报单引用
	OrderRef types.TThostFtdcOrderRefType
	// 前置编号
	FrontID types.TThostFtdcFrontIDType
	// 会话编号
	SessionID types.TThostFtdcSessionIDType
	// IP地址
	IPAddress types.TThostFtdcIPAddressType
	// Mac地址
	MacAddress types.TThostFtdcMacAddressType
}

func (d CThostFtdcSpdApplyActionField) Type() string { return "CThostFtdcSpdApplyActionField" }

func (d CThostFtdcSpdApplyActionField) String() string {
	var builder strings.Builder
	builder.Grow(444)

	builder.WriteString("CThostFtdcSpdApplyActionField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", InvestorID=%+v", d.InvestorID)
	fmt.Fprintf(&builder, ", ActionDate=%+v", d.ActionDate)
	fmt.Fprintf(&builder, ", ActionTime=%+v", d.ActionTime)
	fmt.Fprintf(&builder, ", TraderID=%+v", d.TraderID)
	fmt.Fprintf(&builder, ", InstallID=%+v", d.InstallID)
	fmt.Fprintf(&builder, ", OrderLocalID=%+v", d.OrderLocalID)
	fmt.Fprintf(&builder, ", ActionLocalID=%+v", d.ActionLocalID)
	fmt.Fprintf(&builder, ", ParticipantID=%+v", d.ParticipantID)
	fmt.Fprintf(&builder, ", ClientID=%+v", d.ClientID)
	fmt.Fprintf(&builder, ", OrderActionStatus=%+v", d.OrderActionStatus)
	fmt.Fprintf(&builder, ", UserID=%+v", d.UserID)
	fmt.Fprintf(&builder, ", ExchangeID=%+v", d.ExchangeID)
	fmt.Fprintf(&builder, ", OrderSysID=%+v", d.OrderSysID)
	fmt.Fprintf(&builder, ", RequestID=%+v", d.RequestID)
	fmt.Fprintf(&builder, ", StatusMsg=%+v", d.StatusMsg)
	fmt.Fprintf(&builder, ", OrderRef=%+v", d.OrderRef)
	fmt.Fprintf(&builder, ", FrontID=%+v", d.FrontID)
	fmt.Fprintf(&builder, ", SessionID=%+v", d.SessionID)
	fmt.Fprintf(&builder, ", IPAddress=%+v", d.IPAddress)
	fmt.Fprintf(&builder, ", MacAddress=%+v", d.MacAddress)

	builder.WriteByte('}')

	return builder.String()
}

// 套保申请撤销回报
type CThostFtdcHedgeCfmActionField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID types.TThostFtdcInvestorIDType
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
	// 报单操作状态
	OrderActionStatus types.TThostFtdcOrderActionStatusType
	// 用户代码
	UserID types.TThostFtdcUserIDType
	// 交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	// 合同编号
	OrderSysID types.TThostFtdcOrderSysIDType
	// 请求编号
	RequestID types.TThostFtdcRequestIDType
	// 状态信息
	StatusMsg types.TThostFtdcErrorMsgType
	// 报单引用
	OrderRef types.TThostFtdcOrderRefType
	// 前置编号
	FrontID types.TThostFtdcFrontIDType
	// 会话编号
	SessionID types.TThostFtdcSessionIDType
	// IP地址
	IPAddress types.TThostFtdcIPAddressType
	// Mac地址
	MacAddress types.TThostFtdcMacAddressType
}

func (d CThostFtdcHedgeCfmActionField) Type() string { return "CThostFtdcHedgeCfmActionField" }

func (d CThostFtdcHedgeCfmActionField) String() string {
	var builder strings.Builder
	builder.Grow(444)

	builder.WriteString("CThostFtdcHedgeCfmActionField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", InvestorID=%+v", d.InvestorID)
	fmt.Fprintf(&builder, ", ActionDate=%+v", d.ActionDate)
	fmt.Fprintf(&builder, ", ActionTime=%+v", d.ActionTime)
	fmt.Fprintf(&builder, ", TraderID=%+v", d.TraderID)
	fmt.Fprintf(&builder, ", InstallID=%+v", d.InstallID)
	fmt.Fprintf(&builder, ", OrderLocalID=%+v", d.OrderLocalID)
	fmt.Fprintf(&builder, ", ActionLocalID=%+v", d.ActionLocalID)
	fmt.Fprintf(&builder, ", ParticipantID=%+v", d.ParticipantID)
	fmt.Fprintf(&builder, ", ClientID=%+v", d.ClientID)
	fmt.Fprintf(&builder, ", OrderActionStatus=%+v", d.OrderActionStatus)
	fmt.Fprintf(&builder, ", UserID=%+v", d.UserID)
	fmt.Fprintf(&builder, ", ExchangeID=%+v", d.ExchangeID)
	fmt.Fprintf(&builder, ", OrderSysID=%+v", d.OrderSysID)
	fmt.Fprintf(&builder, ", RequestID=%+v", d.RequestID)
	fmt.Fprintf(&builder, ", StatusMsg=%+v", d.StatusMsg)
	fmt.Fprintf(&builder, ", OrderRef=%+v", d.OrderRef)
	fmt.Fprintf(&builder, ", FrontID=%+v", d.FrontID)
	fmt.Fprintf(&builder, ", SessionID=%+v", d.SessionID)
	fmt.Fprintf(&builder, ", IPAddress=%+v", d.IPAddress)
	fmt.Fprintf(&builder, ", MacAddress=%+v", d.MacAddress)

	builder.WriteByte('}')

	return builder.String()
}

// 前置信息
type CThostFtdcFrontInfoField struct {
	// 前置地址
	FrontAddr types.TThostFtdcAddressType
	// 查询流控
	QryFreq types.TThostFtdcQueryFreqType
	// FTD流控
	FTDPkgFreq types.TThostFtdcQueryFreqType
}

func (d CThostFtdcFrontInfoField) Type() string { return "CThostFtdcFrontInfoField" }

func (d CThostFtdcFrontInfoField) String() string {
	var builder strings.Builder
	builder.Grow(80)

	builder.WriteString("CThostFtdcFrontInfoField{")
	fmt.Fprintf(&builder, "FrontAddr=%+v", d.FrontAddr)
	fmt.Fprintf(&builder, ", QryFreq=%+v", d.QryFreq)
	fmt.Fprintf(&builder, ", FTDPkgFreq=%+v", d.FTDPkgFreq)

	builder.WriteByte('}')

	return builder.String()
}
