package mini

import (
	"fmt"
	"strings"

	"github.com/frozenpine/ctp4go/thost/mini/types"
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
	// 终端IP地址
	ClientIPAddress types.TThostFtdcIPAddressType
	// 登录备注
	LoginRemark types.TThostFtdcLoginRemarkType
}

func (d CThostFtdcReqUserLoginField) Type() string { return "CThostFtdcReqUserLoginField" }

func (d CThostFtdcReqUserLoginField) String() string {
	var builder strings.Builder
	builder.Grow(267)

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
	fmt.Fprintf(&builder, ", ClientIPAddress=%+v", d.ClientIPAddress)
	fmt.Fprintf(&builder, ", LoginRemark=%+v", d.LoginRemark)

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
	// 广期所时间
	GFEXTime types.TThostFtdcTimeType
}

func (d CThostFtdcRspUserLoginField) Type() string { return "CThostFtdcRspUserLoginField" }

func (d CThostFtdcRspUserLoginField) String() string {
	var builder strings.Builder
	builder.Grow(283)

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
	fmt.Fprintf(&builder, ", GFEXTime=%+v", d.GFEXTime)

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
	AppID types.TThostFtdcClientAppIDType
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
	AppID types.TThostFtdcClientAppIDType
}

func (d CThostFtdcRspAuthenticateField) Type() string { return "CThostFtdcRspAuthenticateField" }

func (d CThostFtdcRspAuthenticateField) String() string {
	var builder strings.Builder
	builder.Grow(104)

	builder.WriteString("CThostFtdcRspAuthenticateField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", UserID=%+v", d.UserID)
	fmt.Fprintf(&builder, ", UserProductInfo=%+v", d.UserProductInfo)
	fmt.Fprintf(&builder, ", AppID=%+v", d.AppID)

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
	// 时间戳
	TimeStamp types.TThostFtdcAuthInfoType
	// 认证信息
	AuthInfo types.TThostFtdcAuthInfoType
	// 是否为认证结果
	IsResult types.TThostFtdcBoolType
	// App代码
	AppID types.TThostFtdcClientAppIDType
}

func (d CThostFtdcAuthenticationInfoField) Type() string { return "CThostFtdcAuthenticationInfoField" }

func (d CThostFtdcAuthenticationInfoField) String() string {
	var builder strings.Builder
	builder.Grow(162)

	builder.WriteString("CThostFtdcAuthenticationInfoField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", UserID=%+v", d.UserID)
	fmt.Fprintf(&builder, ", UserProductInfo=%+v", d.UserProductInfo)
	fmt.Fprintf(&builder, ", TimeStamp=%+v", d.TimeStamp)
	fmt.Fprintf(&builder, ", AuthInfo=%+v", d.AuthInfo)
	fmt.Fprintf(&builder, ", IsResult=%+v", d.IsResult)
	fmt.Fprintf(&builder, ", AppID=%+v", d.AppID)

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
	// 批量交易成功记录数
	RecordCount types.TThostFtdcRecordCountType
}

func (d CThostFtdcRspInfoField) Type() string { return "CThostFtdcRspInfoField" }

func (d CThostFtdcRspInfoField) String() string {
	var builder strings.Builder
	builder.Grow(78)

	builder.WriteString("CThostFtdcRspInfoField{")
	fmt.Fprintf(&builder, "ErrorID=%+v", d.ErrorID)
	fmt.Fprintf(&builder, ", ErrorMsg=%+v", d.ErrorMsg)
	fmt.Fprintf(&builder, ", RecordCount=%+v", d.RecordCount)

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

func (d CThostFtdcProductField) Type() string { return "CThostFtdcProductField" }

func (d CThostFtdcProductField) String() string {
	var builder strings.Builder
	builder.Grow(446)

	builder.WriteString("CThostFtdcProductField{")
	fmt.Fprintf(&builder, "ProductID=%+v", d.ProductID)
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
	fmt.Fprintf(&builder, ", ExchangeProductID=%+v", d.ExchangeProductID)
	fmt.Fprintf(&builder, ", UnderlyingMultiple=%+v", d.UnderlyingMultiple)

	builder.WriteByte('}')

	return builder.String()
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

func (d CThostFtdcInstrumentField) Type() string { return "CThostFtdcInstrumentField" }

func (d CThostFtdcInstrumentField) String() string {
	var builder strings.Builder
	builder.Grow(761)

	builder.WriteString("CThostFtdcInstrumentField{")
	fmt.Fprintf(&builder, "InstrumentID=%+v", d.InstrumentID)
	fmt.Fprintf(&builder, ", ExchangeID=%+v", d.ExchangeID)
	fmt.Fprintf(&builder, ", InstrumentName=%+v", d.InstrumentName)
	fmt.Fprintf(&builder, ", ExchangeInstID=%+v", d.ExchangeInstID)
	fmt.Fprintf(&builder, ", ProductID=%+v", d.ProductID)
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
	fmt.Fprintf(&builder, ", UnderlyingInstrID=%+v", d.UnderlyingInstrID)
	fmt.Fprintf(&builder, ", StrikePrice=%+v", d.StrikePrice)
	fmt.Fprintf(&builder, ", OptionsType=%+v", d.OptionsType)
	fmt.Fprintf(&builder, ", UnderlyingMultiple=%+v", d.UnderlyingMultiple)
	fmt.Fprintf(&builder, ", CombinationType=%+v", d.CombinationType)

	builder.WriteByte('}')

	return builder.String()
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

func (d CThostFtdcCombInstrumentField) Type() string { return "CThostFtdcCombInstrumentField" }

func (d CThostFtdcCombInstrumentField) String() string {
	var builder strings.Builder
	builder.Grow(135)

	builder.WriteString("CThostFtdcCombInstrumentField{")
	fmt.Fprintf(&builder, "InstrumentID=%+v", d.InstrumentID)
	fmt.Fprintf(&builder, ", ExchangeID=%+v", d.ExchangeID)
	fmt.Fprintf(&builder, ", CombinationType=%+v", d.CombinationType)
	fmt.Fprintf(&builder, ", ProductID=%+v", d.ProductID)
	fmt.Fprintf(&builder, ", Xparameter=%+v", d.Xparameter)

	builder.WriteByte('}')

	return builder.String()
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

func (d CThostFtdcRCAMSInvestorProdMarginField) Type() string {
	return "CThostFtdcRCAMSInvestorProdMarginField"
}

func (d CThostFtdcRCAMSInvestorProdMarginField) String() string {
	var builder strings.Builder
	builder.Grow(536)

	builder.WriteString("CThostFtdcRCAMSInvestorProdMarginField{")
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
	fmt.Fprintf(&builder, ", OptionRoyalty=%+v", d.OptionRoyalty)
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
}

func (d CThostFtdcTraderField) Type() string { return "CThostFtdcTraderField" }

func (d CThostFtdcTraderField) String() string {
	var builder strings.Builder
	builder.Grow(140)

	builder.WriteString("CThostFtdcTraderField{")
	fmt.Fprintf(&builder, "ExchangeID=%+v", d.ExchangeID)
	fmt.Fprintf(&builder, ", TraderID=%+v", d.TraderID)
	fmt.Fprintf(&builder, ", ParticipantID=%+v", d.ParticipantID)
	fmt.Fprintf(&builder, ", Password=%+v", d.Password)
	fmt.Fprintf(&builder, ", InstallCount=%+v", d.InstallCount)
	fmt.Fprintf(&builder, ", BrokerID=%+v", d.BrokerID)

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
}

func (d CThostFtdcInvestorField) Type() string { return "CThostFtdcInvestorField" }

func (d CThostFtdcInvestorField) String() string {
	var builder strings.Builder
	builder.Grow(294)

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
}

func (d CThostFtdcTradingCodeField) Type() string { return "CThostFtdcTradingCodeField" }

func (d CThostFtdcTradingCodeField) String() string {
	var builder strings.Builder
	builder.Grow(142)

	builder.WriteString("CThostFtdcTradingCodeField{")
	fmt.Fprintf(&builder, "InvestorID=%+v", d.InvestorID)
	fmt.Fprintf(&builder, ", BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", ExchangeID=%+v", d.ExchangeID)
	fmt.Fprintf(&builder, ", ClientID=%+v", d.ClientID)
	fmt.Fprintf(&builder, ", IsActive=%+v", d.IsActive)
	fmt.Fprintf(&builder, ", ClientIDType=%+v", d.ClientIDType)

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
	// 分仓冻结资金
	FrozenPartition types.TThostFtdcMoneyType
}

func (d CThostFtdcTradingAccountField) Type() string { return "CThostFtdcTradingAccountField" }

func (d CThostFtdcTradingAccountField) String() string {
	var builder strings.Builder
	builder.Grow(1141)

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
	fmt.Fprintf(&builder, ", FrozenPartition=%+v", d.FrozenPartition)

	builder.WriteByte('}')

	return builder.String()
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

func (d CThostFtdcInvestorPositionField) Type() string { return "CThostFtdcInvestorPositionField" }

func (d CThostFtdcInvestorPositionField) String() string {
	var builder strings.Builder
	builder.Grow(992)

	builder.WriteString("CThostFtdcInvestorPositionField{")
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

	builder.WriteByte('}')

	return builder.String()
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

func (d CThostFtdcInstrumentMarginRateField) Type() string {
	return "CThostFtdcInstrumentMarginRateField"
}

func (d CThostFtdcInstrumentMarginRateField) String() string {
	var builder strings.Builder
	builder.Grow(289)

	builder.WriteString("CThostFtdcInstrumentMarginRateField{")
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

	builder.WriteByte('}')

	return builder.String()
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

func (d CThostFtdcInstrumentCommissionRateField) Type() string {
	return "CThostFtdcInstrumentCommissionRateField"
}

func (d CThostFtdcInstrumentCommissionRateField) String() string {
	var builder strings.Builder
	builder.Grow(295)

	builder.WriteString("CThostFtdcInstrumentCommissionRateField{")
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

	builder.WriteByte('}')

	return builder.String()
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

func (d CThostFtdcInvestorProdSPBMDetailField) Type() string {
	return "CThostFtdcInvestorProdSPBMDetailField"
}

func (d CThostFtdcInvestorProdSPBMDetailField) String() string {
	var builder strings.Builder
	builder.Grow(425)

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
	fmt.Fprintf(&builder, ", OptionMinRisk=%+v", d.OptionMinRisk)
	fmt.Fprintf(&builder, ", RealOptionValueOffset=%+v", d.RealOptionValueOffset)
	fmt.Fprintf(&builder, ", Margin=%+v", d.Margin)
	fmt.Fprintf(&builder, ", ExchMargin=%+v", d.ExchMargin)

	builder.WriteByte('}')

	return builder.String()
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

func (d CThostFtdcSPMMInvestorCommodityGroupMarginField) Type() string {
	return "CThostFtdcSPMMInvestorCommodityGroupMarginField"
}

func (d CThostFtdcSPMMInvestorCommodityGroupMarginField) String() string {
	var builder strings.Builder
	builder.Grow(543)

	builder.WriteString("CThostFtdcSPMMInvestorCommodityGroupMarginField{")
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

func (d CThostFtdcRULEInvestorProdMarginField) Type() string {
	return "CThostFtdcRULEInvestorProdMarginField"
}

func (d CThostFtdcRULEInvestorProdMarginField) String() string {
	var builder strings.Builder
	builder.Grow(397)

	builder.WriteString("CThostFtdcRULEInvestorProdMarginField{")
	fmt.Fprintf(&builder, "ExchangeID=%+v", d.ExchangeID)
	fmt.Fprintf(&builder, ", BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", InvestorID=%+v", d.InvestorID)
	fmt.Fprintf(&builder, ", ProdFamilyCode=%+v", d.ProdFamilyCode)
	fmt.Fprintf(&builder, ", InstrumentClass=%+v", d.InstrumentClass)
	fmt.Fprintf(&builder, ", CommodityGroupID=%+v", d.CommodityGroupID)
	fmt.Fprintf(&builder, ", IntraProdMargin=%+v", d.IntraProdMargin)
	fmt.Fprintf(&builder, ", InterProdMargin=%+v", d.InterProdMargin)
	fmt.Fprintf(&builder, ", SingleMargin=%+v", d.SingleMargin)
	fmt.Fprintf(&builder, ", NonCombMargin=%+v", d.NonCombMargin)
	fmt.Fprintf(&builder, ", AddOnMargin=%+v", d.AddOnMargin)
	fmt.Fprintf(&builder, ", ExchMargin=%+v", d.ExchMargin)
	fmt.Fprintf(&builder, ", OpenFrozenMargin=%+v", d.OpenFrozenMargin)
	fmt.Fprintf(&builder, ", CloseFrozenMargin=%+v", d.CloseFrozenMargin)
	fmt.Fprintf(&builder, ", Margin=%+v", d.Margin)
	fmt.Fprintf(&builder, ", FrozenMargin=%+v", d.FrozenMargin)

	builder.WriteByte('}')

	return builder.String()
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

func (d CThostFtdcDepthMarketDataField) Type() string { return "CThostFtdcDepthMarketDataField" }

func (d CThostFtdcDepthMarketDataField) String() string {
	var builder strings.Builder
	builder.Grow(990)

	builder.WriteString("CThostFtdcDepthMarketDataField{")
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

	builder.WriteByte('}')

	return builder.String()
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

func (d CThostFtdcInstrumentTradingRightField) Type() string {
	return "CThostFtdcInstrumentTradingRightField"
}

func (d CThostFtdcInstrumentTradingRightField) String() string {
	var builder strings.Builder
	builder.Grow(142)

	builder.WriteString("CThostFtdcInstrumentTradingRightField{")
	fmt.Fprintf(&builder, "InstrumentID=%+v", d.InstrumentID)
	fmt.Fprintf(&builder, ", InvestorRange=%+v", d.InvestorRange)
	fmt.Fprintf(&builder, ", BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", InvestorID=%+v", d.InvestorID)
	fmt.Fprintf(&builder, ", TradingRight=%+v", d.TradingRight)

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
}

func (d CThostFtdcBrokerUserField) Type() string { return "CThostFtdcBrokerUserField" }

func (d CThostFtdcBrokerUserField) String() string {
	var builder strings.Builder
	builder.Grow(133)

	builder.WriteString("CThostFtdcBrokerUserField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", UserID=%+v", d.UserID)
	fmt.Fprintf(&builder, ", UserName=%+v", d.UserName)
	fmt.Fprintf(&builder, ", UserType=%+v", d.UserType)
	fmt.Fprintf(&builder, ", IsActive=%+v", d.IsActive)
	fmt.Fprintf(&builder, ", IsUsingOTP=%+v", d.IsUsingOTP)

	builder.WriteByte('}')

	return builder.String()
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

func (d CThostFtdcControlParamField) Type() string { return "CThostFtdcControlParamField" }

func (d CThostFtdcControlParamField) String() string {
	var builder strings.Builder
	builder.Grow(112)

	builder.WriteString("CThostFtdcControlParamField{")
	fmt.Fprintf(&builder, "InvestorID=%+v", d.InvestorID)
	fmt.Fprintf(&builder, ", ControlParamID=%+v", d.ControlParamID)
	fmt.Fprintf(&builder, ", ControlParamValue=%+v", d.ControlParamValue)
	fmt.Fprintf(&builder, ", Memo=%+v", d.Memo)

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
}

func (d CThostFtdcBrokerUserPasswordField) Type() string { return "CThostFtdcBrokerUserPasswordField" }

func (d CThostFtdcBrokerUserPasswordField) String() string {
	var builder strings.Builder
	builder.Grow(85)

	builder.WriteString("CThostFtdcBrokerUserPasswordField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", UserID=%+v", d.UserID)
	fmt.Fprintf(&builder, ", Password=%+v", d.Password)

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
}

func (d CThostFtdcTraderOfferField) Type() string { return "CThostFtdcTraderOfferField" }

func (d CThostFtdcTraderOfferField) String() string {
	var builder strings.Builder
	builder.Grow(451)

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
}

func (d CThostFtdcSettlementInfoField) Type() string { return "CThostFtdcSettlementInfoField" }

func (d CThostFtdcSettlementInfoField) String() string {
	var builder strings.Builder
	builder.Grow(146)

	builder.WriteString("CThostFtdcSettlementInfoField{")
	fmt.Fprintf(&builder, "TradingDay=%+v", d.TradingDay)
	fmt.Fprintf(&builder, ", SettlementID=%+v", d.SettlementID)
	fmt.Fprintf(&builder, ", BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", InvestorID=%+v", d.InvestorID)
	fmt.Fprintf(&builder, ", SequenceNo=%+v", d.SequenceNo)
	fmt.Fprintf(&builder, ", Content=%+v", d.Content)

	builder.WriteByte('}')

	return builder.String()
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

func (d CThostFtdcInstrumentMarginRateAdjustField) Type() string {
	return "CThostFtdcInstrumentMarginRateAdjustField"
}

func (d CThostFtdcInstrumentMarginRateAdjustField) String() string {
	var builder strings.Builder
	builder.Grow(295)

	builder.WriteString("CThostFtdcInstrumentMarginRateAdjustField{")
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

	builder.WriteByte('}')

	return builder.String()
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

func (d CThostFtdcExchangeMarginRateField) Type() string { return "CThostFtdcExchangeMarginRateField" }

func (d CThostFtdcExchangeMarginRateField) String() string {
	var builder strings.Builder
	builder.Grow(224)

	builder.WriteString("CThostFtdcExchangeMarginRateField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", InstrumentID=%+v", d.InstrumentID)
	fmt.Fprintf(&builder, ", HedgeFlag=%+v", d.HedgeFlag)
	fmt.Fprintf(&builder, ", LongMarginRatioByMoney=%+v", d.LongMarginRatioByMoney)
	fmt.Fprintf(&builder, ", LongMarginRatioByVolume=%+v", d.LongMarginRatioByVolume)
	fmt.Fprintf(&builder, ", ShortMarginRatioByMoney=%+v", d.ShortMarginRatioByMoney)
	fmt.Fprintf(&builder, ", ShortMarginRatioByVolume=%+v", d.ShortMarginRatioByVolume)

	builder.WriteByte('}')

	return builder.String()
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

func (d CThostFtdcExchangeMarginRateAdjustField) Type() string {
	return "CThostFtdcExchangeMarginRateAdjustField"
}

func (d CThostFtdcExchangeMarginRateAdjustField) String() string {
	var builder strings.Builder
	builder.Grow(518)

	builder.WriteString("CThostFtdcExchangeMarginRateAdjustField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", InstrumentID=%+v", d.InstrumentID)
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
	// 当前日期
	CurrDate types.TThostFtdcDateType
	// 当前时间
	CurrTime types.TThostFtdcTimeType
	// 当前时间（毫秒）
	CurrMillisec types.TThostFtdcMillisecType
	// 业务日期
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

func (d CThostFtdcLoginInfoField) Type() string { return "CThostFtdcLoginInfoField" }

func (d CThostFtdcLoginInfoField) String() string {
	var builder strings.Builder
	builder.Grow(463)

	builder.WriteString("CThostFtdcLoginInfoField{")
	fmt.Fprintf(&builder, "FrontID=%+v", d.FrontID)
	fmt.Fprintf(&builder, ", SessionID=%+v", d.SessionID)
	fmt.Fprintf(&builder, ", BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", UserID=%+v", d.UserID)
	fmt.Fprintf(&builder, ", LoginDate=%+v", d.LoginDate)
	fmt.Fprintf(&builder, ", LoginTime=%+v", d.LoginTime)
	fmt.Fprintf(&builder, ", IPAddress=%+v", d.IPAddress)
	fmt.Fprintf(&builder, ", UserProductInfo=%+v", d.UserProductInfo)
	fmt.Fprintf(&builder, ", InterfaceProductInfo=%+v", d.InterfaceProductInfo)
	fmt.Fprintf(&builder, ", ProtocolInfo=%+v", d.ProtocolInfo)
	fmt.Fprintf(&builder, ", SystemName=%+v", d.SystemName)
	fmt.Fprintf(&builder, ", Password=%+v", d.Password)
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

func (d CThostFtdcInputOrderField) Type() string { return "CThostFtdcInputOrderField" }

func (d CThostFtdcInputOrderField) String() string {
	var builder strings.Builder
	builder.Grow(706)

	builder.WriteString("CThostFtdcInputOrderField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", InvestorID=%+v", d.InvestorID)
	fmt.Fprintf(&builder, ", InstrumentID=%+v", d.InstrumentID)
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
	fmt.Fprintf(&builder, ", IPAddress=%+v", d.IPAddress)
	fmt.Fprintf(&builder, ", MacAddress=%+v", d.MacAddress)
	fmt.Fprintf(&builder, ", CustomOrderRef=%+v", d.CustomOrderRef)
	fmt.Fprintf(&builder, ", OrderFlag=%+v", d.OrderFlag)

	builder.WriteByte('}')

	return builder.String()
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

func (d CThostFtdcOrderField) Type() string { return "CThostFtdcOrderField" }

func (d CThostFtdcOrderField) String() string {
	var builder strings.Builder
	builder.Grow(1419)

	builder.WriteString("CThostFtdcOrderField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", InvestorID=%+v", d.InvestorID)
	fmt.Fprintf(&builder, ", InstrumentID=%+v", d.InstrumentID)
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
	fmt.Fprintf(&builder, ", ExchangeInstID=%+v", d.ExchangeInstID)
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
	fmt.Fprintf(&builder, ", IPAddress=%+v", d.IPAddress)
	fmt.Fprintf(&builder, ", MacAddress=%+v", d.MacAddress)
	fmt.Fprintf(&builder, ", CustomOrderRef=%+v", d.CustomOrderRef)
	fmt.Fprintf(&builder, ", TradeAvgPrice=%+v", d.TradeAvgPrice)

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

func (d CThostFtdcExchangeOrderField) Type() string { return "CThostFtdcExchangeOrderField" }

func (d CThostFtdcExchangeOrderField) String() string {
	var builder strings.Builder
	builder.Grow(996)

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
	fmt.Fprintf(&builder, ", ExchangeInstID=%+v", d.ExchangeInstID)
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
	fmt.Fprintf(&builder, ", IPAddress=%+v", d.IPAddress)
	fmt.Fprintf(&builder, ", MacAddress=%+v", d.MacAddress)

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
	// 合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
	// 投资单元代码
	InvestUnitID types.TThostFtdcInvestUnitIDType
	// IP地址
	IPAddress types.TThostFtdcIPAddressType
	// Mac地址
	MacAddress types.TThostFtdcMacAddressType
}

func (d CThostFtdcInputOrderActionField) Type() string { return "CThostFtdcInputOrderActionField" }

func (d CThostFtdcInputOrderActionField) String() string {
	var builder strings.Builder
	builder.Grow(367)

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
	fmt.Fprintf(&builder, ", InstrumentID=%+v", d.InstrumentID)
	fmt.Fprintf(&builder, ", InvestUnitID=%+v", d.InvestUnitID)
	fmt.Fprintf(&builder, ", IPAddress=%+v", d.IPAddress)
	fmt.Fprintf(&builder, ", MacAddress=%+v", d.MacAddress)

	builder.WriteByte('}')

	return builder.String()
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

func (d CThostFtdcMKInputOrderActionField) Type() string { return "CThostFtdcMKInputOrderActionField" }

func (d CThostFtdcMKInputOrderActionField) String() string {
	var builder strings.Builder
	builder.Grow(389)

	builder.WriteString("CThostFtdcMKInputOrderActionField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", InvestorID=%+v", d.InvestorID)
	fmt.Fprintf(&builder, ", OrderActionRef=%+v", d.OrderActionRef)
	fmt.Fprintf(&builder, ", RequestID=%+v", d.RequestID)
	fmt.Fprintf(&builder, ", FrontID=%+v", d.FrontID)
	fmt.Fprintf(&builder, ", SessionID=%+v", d.SessionID)
	fmt.Fprintf(&builder, ", ExchangeID=%+v", d.ExchangeID)
	fmt.Fprintf(&builder, ", ActionFlag=%+v", d.ActionFlag)
	fmt.Fprintf(&builder, ", ActionMode=%+v", d.ActionMode)
	fmt.Fprintf(&builder, ", LowerLimit=%+v", d.LowerLimit)
	fmt.Fprintf(&builder, ", UpperLimit=%+v", d.UpperLimit)
	fmt.Fprintf(&builder, ", UserID=%+v", d.UserID)
	fmt.Fprintf(&builder, ", ProductID=%+v", d.ProductID)
	fmt.Fprintf(&builder, ", InstrumentID=%+v", d.InstrumentID)
	fmt.Fprintf(&builder, ", OrderStatus=%+v", d.OrderStatus)
	fmt.Fprintf(&builder, ", InvestUnitID=%+v", d.InvestUnitID)
	fmt.Fprintf(&builder, ", IPAddress=%+v", d.IPAddress)
	fmt.Fprintf(&builder, ", MacAddress=%+v", d.MacAddress)

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

func (d CThostFtdcOrderActionField) Type() string { return "CThostFtdcOrderActionField" }

func (d CThostFtdcOrderActionField) String() string {
	var builder strings.Builder
	builder.Grow(611)

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
	fmt.Fprintf(&builder, ", InstrumentID=%+v", d.InstrumentID)
	fmt.Fprintf(&builder, ", BranchID=%+v", d.BranchID)
	fmt.Fprintf(&builder, ", InvestUnitID=%+v", d.InvestUnitID)
	fmt.Fprintf(&builder, ", IPAddress=%+v", d.IPAddress)
	fmt.Fprintf(&builder, ", MacAddress=%+v", d.MacAddress)

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
	// IP地址
	IPAddress types.TThostFtdcIPAddressType
	// Mac地址
	MacAddress types.TThostFtdcMacAddressType
}

func (d CThostFtdcExchangeOrderActionField) Type() string {
	return "CThostFtdcExchangeOrderActionField"
}

func (d CThostFtdcExchangeOrderActionField) String() string {
	var builder strings.Builder
	builder.Grow(421)

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
	fmt.Fprintf(&builder, ", IPAddress=%+v", d.IPAddress)
	fmt.Fprintf(&builder, ", MacAddress=%+v", d.MacAddress)

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

func (d CThostFtdcExchangeTradeField) Type() string { return "CThostFtdcExchangeTradeField" }

func (d CThostFtdcExchangeTradeField) String() string {
	var builder strings.Builder
	builder.Grow(465)

	builder.WriteString("CThostFtdcExchangeTradeField{")
	fmt.Fprintf(&builder, "ExchangeID=%+v", d.ExchangeID)
	fmt.Fprintf(&builder, ", TradeID=%+v", d.TradeID)
	fmt.Fprintf(&builder, ", Direction=%+v", d.Direction)
	fmt.Fprintf(&builder, ", OrderSysID=%+v", d.OrderSysID)
	fmt.Fprintf(&builder, ", ParticipantID=%+v", d.ParticipantID)
	fmt.Fprintf(&builder, ", ClientID=%+v", d.ClientID)
	fmt.Fprintf(&builder, ", TradingRole=%+v", d.TradingRole)
	fmt.Fprintf(&builder, ", ExchangeInstID=%+v", d.ExchangeInstID)
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

	builder.WriteByte('}')

	return builder.String()
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

func (d CThostFtdcTradeField) Type() string { return "CThostFtdcTradeField" }

func (d CThostFtdcTradeField) String() string {
	var builder strings.Builder
	builder.Grow(617)

	builder.WriteString("CThostFtdcTradeField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", InvestorID=%+v", d.InvestorID)
	fmt.Fprintf(&builder, ", InstrumentID=%+v", d.InstrumentID)
	fmt.Fprintf(&builder, ", OrderRef=%+v", d.OrderRef)
	fmt.Fprintf(&builder, ", UserID=%+v", d.UserID)
	fmt.Fprintf(&builder, ", ExchangeID=%+v", d.ExchangeID)
	fmt.Fprintf(&builder, ", TradeID=%+v", d.TradeID)
	fmt.Fprintf(&builder, ", Direction=%+v", d.Direction)
	fmt.Fprintf(&builder, ", OrderSysID=%+v", d.OrderSysID)
	fmt.Fprintf(&builder, ", ParticipantID=%+v", d.ParticipantID)
	fmt.Fprintf(&builder, ", ClientID=%+v", d.ClientID)
	fmt.Fprintf(&builder, ", TradingRole=%+v", d.TradingRole)
	fmt.Fprintf(&builder, ", ExchangeInstID=%+v", d.ExchangeInstID)
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

func (d CThostFtdcUserSessionField) Type() string { return "CThostFtdcUserSessionField" }

func (d CThostFtdcUserSessionField) String() string {
	var builder strings.Builder
	builder.Grow(271)

	builder.WriteString("CThostFtdcUserSessionField{")
	fmt.Fprintf(&builder, "FrontID=%+v", d.FrontID)
	fmt.Fprintf(&builder, ", SessionID=%+v", d.SessionID)
	fmt.Fprintf(&builder, ", BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", UserID=%+v", d.UserID)
	fmt.Fprintf(&builder, ", LoginDate=%+v", d.LoginDate)
	fmt.Fprintf(&builder, ", LoginTime=%+v", d.LoginTime)
	fmt.Fprintf(&builder, ", IPAddress=%+v", d.IPAddress)
	fmt.Fprintf(&builder, ", UserProductInfo=%+v", d.UserProductInfo)
	fmt.Fprintf(&builder, ", InterfaceProductInfo=%+v", d.InterfaceProductInfo)
	fmt.Fprintf(&builder, ", ProtocolInfo=%+v", d.ProtocolInfo)
	fmt.Fprintf(&builder, ", MacAddress=%+v", d.MacAddress)
	fmt.Fprintf(&builder, ", LoginRemark=%+v", d.LoginRemark)

	builder.WriteByte('}')

	return builder.String()
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

func (d CThostFtdcQueryMaxOrderVolumeField) Type() string {
	return "CThostFtdcQueryMaxOrderVolumeField"
}

func (d CThostFtdcQueryMaxOrderVolumeField) String() string {
	var builder strings.Builder
	builder.Grow(171)

	builder.WriteString("CThostFtdcQueryMaxOrderVolumeField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", InvestorID=%+v", d.InvestorID)
	fmt.Fprintf(&builder, ", InstrumentID=%+v", d.InstrumentID)
	fmt.Fprintf(&builder, ", Direction=%+v", d.Direction)
	fmt.Fprintf(&builder, ", OffsetFlag=%+v", d.OffsetFlag)
	fmt.Fprintf(&builder, ", HedgeFlag=%+v", d.HedgeFlag)
	fmt.Fprintf(&builder, ", MaxVolume=%+v", d.MaxVolume)

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
}

func (d CThostFtdcSettlementInfoConfirmField) Type() string {
	return "CThostFtdcSettlementInfoConfirmField"
}

func (d CThostFtdcSettlementInfoConfirmField) String() string {
	var builder strings.Builder
	builder.Grow(116)

	builder.WriteString("CThostFtdcSettlementInfoConfirmField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", InvestorID=%+v", d.InvestorID)
	fmt.Fprintf(&builder, ", ConfirmDate=%+v", d.ConfirmDate)
	fmt.Fprintf(&builder, ", ConfirmTime=%+v", d.ConfirmTime)

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
}

func (d CThostFtdcSyncDepositField) Type() string { return "CThostFtdcSyncDepositField" }

func (d CThostFtdcSyncDepositField) String() string {
	var builder strings.Builder
	builder.Grow(140)

	builder.WriteString("CThostFtdcSyncDepositField{")
	fmt.Fprintf(&builder, "DepositSeqNo=%+v", d.DepositSeqNo)
	fmt.Fprintf(&builder, ", BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", InvestorID=%+v", d.InvestorID)
	fmt.Fprintf(&builder, ", Deposit=%+v", d.Deposit)
	fmt.Fprintf(&builder, ", IsForce=%+v", d.IsForce)
	fmt.Fprintf(&builder, ", CurrencyID=%+v", d.CurrencyID)

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

func (d CThostFtdcTransFundField) Type() string { return "CThostFtdcTransFundField" }

func (d CThostFtdcTransFundField) String() string {
	var builder strings.Builder
	builder.Grow(202)

	builder.WriteString("CThostFtdcTransFundField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", UserID=%+v", d.UserID)
	fmt.Fprintf(&builder, ", InvestorID=%+v", d.InvestorID)
	fmt.Fprintf(&builder, ", InvestUnitID=%+v", d.InvestUnitID)
	fmt.Fprintf(&builder, ", AccountID=%+v", d.AccountID)
	fmt.Fprintf(&builder, ", DepositKernel=%+v", d.DepositKernel)
	fmt.Fprintf(&builder, ", IncomingKernel=%+v", d.IncomingKernel)
	fmt.Fprintf(&builder, ", CurrencyID=%+v", d.CurrencyID)
	fmt.Fprintf(&builder, ", Amount=%+v", d.Amount)

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
}

func (d CThostFtdcSyncingInvestorField) Type() string { return "CThostFtdcSyncingInvestorField" }

func (d CThostFtdcSyncingInvestorField) String() string {
	var builder strings.Builder
	builder.Grow(301)

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
}

func (d CThostFtdcSyncingTradingAccountField) Type() string {
	return "CThostFtdcSyncingTradingAccountField"
}

func (d CThostFtdcSyncingTradingAccountField) String() string {
	var builder strings.Builder
	builder.Grow(1123)

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

	builder.WriteByte('}')

	return builder.String()
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

func (d CThostFtdcSyncingInvestorPositionField) Type() string {
	return "CThostFtdcSyncingInvestorPositionField"
}

func (d CThostFtdcSyncingInvestorPositionField) String() string {
	var builder strings.Builder
	builder.Grow(999)

	builder.WriteString("CThostFtdcSyncingInvestorPositionField{")
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

	builder.WriteByte('}')

	return builder.String()
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

func (d CThostFtdcSyncingInstrumentMarginRateField) Type() string {
	return "CThostFtdcSyncingInstrumentMarginRateField"
}

func (d CThostFtdcSyncingInstrumentMarginRateField) String() string {
	var builder strings.Builder
	builder.Grow(296)

	builder.WriteString("CThostFtdcSyncingInstrumentMarginRateField{")
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

	builder.WriteByte('}')

	return builder.String()
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

func (d CThostFtdcSyncingInstrumentCommissionRateField) Type() string {
	return "CThostFtdcSyncingInstrumentCommissionRateField"
}

func (d CThostFtdcSyncingInstrumentCommissionRateField) String() string {
	var builder strings.Builder
	builder.Grow(302)

	builder.WriteString("CThostFtdcSyncingInstrumentCommissionRateField{")
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

	builder.WriteByte('}')

	return builder.String()
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

func (d CThostFtdcSyncingInstrumentTradingRightField) Type() string {
	return "CThostFtdcSyncingInstrumentTradingRightField"
}

func (d CThostFtdcSyncingInstrumentTradingRightField) String() string {
	var builder strings.Builder
	builder.Grow(149)

	builder.WriteString("CThostFtdcSyncingInstrumentTradingRightField{")
	fmt.Fprintf(&builder, "InstrumentID=%+v", d.InstrumentID)
	fmt.Fprintf(&builder, ", InvestorRange=%+v", d.InvestorRange)
	fmt.Fprintf(&builder, ", BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", InvestorID=%+v", d.InvestorID)
	fmt.Fprintf(&builder, ", TradingRight=%+v", d.TradingRight)

	builder.WriteByte('}')

	return builder.String()
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

func (d CThostFtdcQryOrderField) Type() string { return "CThostFtdcQryOrderField" }

func (d CThostFtdcQryOrderField) String() string {
	var builder strings.Builder
	builder.Grow(171)

	builder.WriteString("CThostFtdcQryOrderField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", InvestorID=%+v", d.InvestorID)
	fmt.Fprintf(&builder, ", InstrumentID=%+v", d.InstrumentID)
	fmt.Fprintf(&builder, ", ExchangeID=%+v", d.ExchangeID)
	fmt.Fprintf(&builder, ", OrderSysID=%+v", d.OrderSysID)
	fmt.Fprintf(&builder, ", InsertTimeStart=%+v", d.InsertTimeStart)
	fmt.Fprintf(&builder, ", InsertTimeEnd=%+v", d.InsertTimeEnd)

	builder.WriteByte('}')

	return builder.String()
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

func (d CThostFtdcQryTradeField) Type() string { return "CThostFtdcQryTradeField" }

func (d CThostFtdcQryTradeField) String() string {
	var builder strings.Builder
	builder.Grow(166)

	builder.WriteString("CThostFtdcQryTradeField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", InvestorID=%+v", d.InvestorID)
	fmt.Fprintf(&builder, ", InstrumentID=%+v", d.InstrumentID)
	fmt.Fprintf(&builder, ", ExchangeID=%+v", d.ExchangeID)
	fmt.Fprintf(&builder, ", TradeID=%+v", d.TradeID)
	fmt.Fprintf(&builder, ", TradeTimeStart=%+v", d.TradeTimeStart)
	fmt.Fprintf(&builder, ", TradeTimeEnd=%+v", d.TradeTimeEnd)

	builder.WriteByte('}')

	return builder.String()
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

func (d CThostFtdcQryInvestorPositionField) Type() string {
	return "CThostFtdcQryInvestorPositionField"
}

func (d CThostFtdcQryInvestorPositionField) String() string {
	var builder strings.Builder
	builder.Grow(94)

	builder.WriteString("CThostFtdcQryInvestorPositionField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", InvestorID=%+v", d.InvestorID)
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
}

func (d CThostFtdcQryTradingAccountField) Type() string { return "CThostFtdcQryTradingAccountField" }

func (d CThostFtdcQryTradingAccountField) String() string {
	var builder strings.Builder
	builder.Grow(90)

	builder.WriteString("CThostFtdcQryTradingAccountField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", InvestorID=%+v", d.InvestorID)
	fmt.Fprintf(&builder, ", CurrencyID=%+v", d.CurrencyID)

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
}

func (d CThostFtdcQryTradingCodeField) Type() string { return "CThostFtdcQryTradingCodeField" }

func (d CThostFtdcQryTradingCodeField) String() string {
	var builder strings.Builder
	builder.Grow(127)

	builder.WriteString("CThostFtdcQryTradingCodeField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", InvestorID=%+v", d.InvestorID)
	fmt.Fprintf(&builder, ", ExchangeID=%+v", d.ExchangeID)
	fmt.Fprintf(&builder, ", ClientID=%+v", d.ClientID)
	fmt.Fprintf(&builder, ", ClientIDType=%+v", d.ClientIDType)

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
	// 合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
	// 投机套保标志
	HedgeFlag types.TThostFtdcHedgeFlagType
}

func (d CThostFtdcQryInstrumentMarginRateField) Type() string {
	return "CThostFtdcQryInstrumentMarginRateField"
}

func (d CThostFtdcQryInstrumentMarginRateField) String() string {
	var builder strings.Builder
	builder.Grow(117)

	builder.WriteString("CThostFtdcQryInstrumentMarginRateField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", InvestorID=%+v", d.InvestorID)
	fmt.Fprintf(&builder, ", InstrumentID=%+v", d.InstrumentID)
	fmt.Fprintf(&builder, ", HedgeFlag=%+v", d.HedgeFlag)

	builder.WriteByte('}')

	return builder.String()
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

func (d CThostFtdcQryInstrumentCommissionRateField) Type() string {
	return "CThostFtdcQryInstrumentCommissionRateField"
}

func (d CThostFtdcQryInstrumentCommissionRateField) String() string {
	var builder strings.Builder
	builder.Grow(102)

	builder.WriteString("CThostFtdcQryInstrumentCommissionRateField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", InvestorID=%+v", d.InvestorID)
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
	// 合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
}

func (d CThostFtdcQryInstrumentTradingRightField) Type() string {
	return "CThostFtdcQryInstrumentTradingRightField"
}

func (d CThostFtdcQryInstrumentTradingRightField) String() string {
	var builder strings.Builder
	builder.Grow(100)

	builder.WriteString("CThostFtdcQryInstrumentTradingRightField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", InvestorID=%+v", d.InvestorID)
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
	// 合约在交易所的代码
	ExchangeInstID types.TThostFtdcExchangeInstIDType
	// 交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	// 交易所交易员代码
	TraderID types.TThostFtdcTraderIDType
}

func (d CThostFtdcQryExchangeOrderField) Type() string { return "CThostFtdcQryExchangeOrderField" }

func (d CThostFtdcQryExchangeOrderField) String() string {
	var builder strings.Builder
	builder.Grow(134)

	builder.WriteString("CThostFtdcQryExchangeOrderField{")
	fmt.Fprintf(&builder, "ParticipantID=%+v", d.ParticipantID)
	fmt.Fprintf(&builder, ", ClientID=%+v", d.ClientID)
	fmt.Fprintf(&builder, ", ExchangeInstID=%+v", d.ExchangeInstID)
	fmt.Fprintf(&builder, ", ExchangeID=%+v", d.ExchangeID)
	fmt.Fprintf(&builder, ", TraderID=%+v", d.TraderID)

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
	// 产品代码
	ProductID types.TThostFtdcInstrumentIDType
	// 产品类型
	ProductClass types.TThostFtdcProductClassType
}

func (d CThostFtdcQryProductField) Type() string { return "CThostFtdcQryProductField" }

func (d CThostFtdcQryProductField) String() string {
	var builder strings.Builder
	builder.Grow(66)

	builder.WriteString("CThostFtdcQryProductField{")
	fmt.Fprintf(&builder, "ProductID=%+v", d.ProductID)
	fmt.Fprintf(&builder, ", ProductClass=%+v", d.ProductClass)

	builder.WriteByte('}')

	return builder.String()
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

func (d CThostFtdcQryInstrumentField) Type() string { return "CThostFtdcQryInstrumentField" }

func (d CThostFtdcQryInstrumentField) String() string {
	var builder strings.Builder
	builder.Grow(113)

	builder.WriteString("CThostFtdcQryInstrumentField{")
	fmt.Fprintf(&builder, "InstrumentID=%+v", d.InstrumentID)
	fmt.Fprintf(&builder, ", ExchangeID=%+v", d.ExchangeID)
	fmt.Fprintf(&builder, ", ExchangeInstID=%+v", d.ExchangeInstID)
	fmt.Fprintf(&builder, ", ProductID=%+v", d.ProductID)

	builder.WriteByte('}')

	return builder.String()
}

// 查询申请组合合约
type CThostFtdcQryCombInstrumentField struct {
	// 交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	// 产品代码
	ProductID types.TThostFtdcInstrumentIDType
}

func (d CThostFtdcQryCombInstrumentField) Type() string { return "CThostFtdcQryCombInstrumentField" }

func (d CThostFtdcQryCombInstrumentField) String() string {
	var builder strings.Builder
	builder.Grow(71)

	builder.WriteString("CThostFtdcQryCombInstrumentField{")
	fmt.Fprintf(&builder, "ExchangeID=%+v", d.ExchangeID)
	fmt.Fprintf(&builder, ", ProductID=%+v", d.ProductID)

	builder.WriteByte('}')

	return builder.String()
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

func (d CThostFtdcQryRCAMSInvestorProdMarginField) Type() string {
	return "CThostFtdcQryRCAMSInvestorProdMarginField"
}

func (d CThostFtdcQryRCAMSInvestorProdMarginField) String() string {
	var builder strings.Builder
	builder.Grow(102)

	builder.WriteString("CThostFtdcQryRCAMSInvestorProdMarginField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", InvestorID=%+v", d.InvestorID)
	fmt.Fprintf(&builder, ", CombProductID=%+v", d.CombProductID)

	builder.WriteByte('}')

	return builder.String()
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

func (d CThostFtdcQryRCAMSInvestorCombPositionField) Type() string {
	return "CThostFtdcQryRCAMSInvestorCombPositionField"
}

func (d CThostFtdcQryRCAMSInvestorCombPositionField) String() string {
	var builder strings.Builder
	builder.Grow(103)

	builder.WriteString("CThostFtdcQryRCAMSInvestorCombPositionField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", InvestorID=%+v", d.InvestorID)
	fmt.Fprintf(&builder, ", InstrumentID=%+v", d.InstrumentID)

	builder.WriteByte('}')

	return builder.String()
}

// 查询行情
type CThostFtdcQryDepthMarketDataField struct {
	// 合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
}

func (d CThostFtdcQryDepthMarketDataField) Type() string { return "CThostFtdcQryDepthMarketDataField" }

func (d CThostFtdcQryDepthMarketDataField) String() string {
	var builder strings.Builder
	builder.Grow(55)

	builder.WriteString("CThostFtdcQryDepthMarketDataField{")
	fmt.Fprintf(&builder, "InstrumentID=%+v", d.InstrumentID)

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

// 查询经纪公司用户
type CThostFtdcQryControlParamField struct {
	// 投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	// 开关代码
	ControlParamID types.TThostFtdcControlParamIDType
}

func (d CThostFtdcQryControlParamField) Type() string { return "CThostFtdcQryControlParamField" }

func (d CThostFtdcQryControlParamField) String() string {
	var builder strings.Builder
	builder.Grow(74)

	builder.WriteString("CThostFtdcQryControlParamField{")
	fmt.Fprintf(&builder, "InvestorID=%+v", d.InvestorID)
	fmt.Fprintf(&builder, ", ControlParamID=%+v", d.ControlParamID)

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
	// 交易所交易员代码
	TraderID types.TThostFtdcTraderIDType
}

func (d CThostFtdcQryTraderOfferField) Type() string { return "CThostFtdcQryTraderOfferField" }

func (d CThostFtdcQryTraderOfferField) String() string {
	var builder strings.Builder
	builder.Grow(67)

	builder.WriteString("CThostFtdcQryTraderOfferField{")
	fmt.Fprintf(&builder, "ExchangeID=%+v", d.ExchangeID)
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
}

func (d CThostFtdcQrySettlementInfoField) Type() string { return "CThostFtdcQrySettlementInfoField" }

func (d CThostFtdcQrySettlementInfoField) String() string {
	var builder strings.Builder
	builder.Grow(90)

	builder.WriteString("CThostFtdcQrySettlementInfoField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", InvestorID=%+v", d.InvestorID)
	fmt.Fprintf(&builder, ", TradingDay=%+v", d.TradingDay)

	builder.WriteByte('}')

	return builder.String()
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

func (d CThostFtdcQryExchangeMarginRateField) Type() string {
	return "CThostFtdcQryExchangeMarginRateField"
}

func (d CThostFtdcQryExchangeMarginRateField) String() string {
	var builder strings.Builder
	builder.Grow(95)

	builder.WriteString("CThostFtdcQryExchangeMarginRateField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", InstrumentID=%+v", d.InstrumentID)
	fmt.Fprintf(&builder, ", HedgeFlag=%+v", d.HedgeFlag)

	builder.WriteByte('}')

	return builder.String()
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

func (d CThostFtdcQryExchangeMarginRateAdjustField) Type() string {
	return "CThostFtdcQryExchangeMarginRateAdjustField"
}

func (d CThostFtdcQryExchangeMarginRateAdjustField) String() string {
	var builder strings.Builder
	builder.Grow(101)

	builder.WriteString("CThostFtdcQryExchangeMarginRateAdjustField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", InstrumentID=%+v", d.InstrumentID)
	fmt.Fprintf(&builder, ", HedgeFlag=%+v", d.HedgeFlag)

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

func (d CThostFtdcQryHisOrderField) Type() string { return "CThostFtdcQryHisOrderField" }

func (d CThostFtdcQryHisOrderField) String() string {
	var builder strings.Builder
	builder.Grow(216)

	builder.WriteString("CThostFtdcQryHisOrderField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", InvestorID=%+v", d.InvestorID)
	fmt.Fprintf(&builder, ", InstrumentID=%+v", d.InstrumentID)
	fmt.Fprintf(&builder, ", ExchangeID=%+v", d.ExchangeID)
	fmt.Fprintf(&builder, ", OrderSysID=%+v", d.OrderSysID)
	fmt.Fprintf(&builder, ", InsertTimeStart=%+v", d.InsertTimeStart)
	fmt.Fprintf(&builder, ", InsertTimeEnd=%+v", d.InsertTimeEnd)
	fmt.Fprintf(&builder, ", TradingDay=%+v", d.TradingDay)
	fmt.Fprintf(&builder, ", SettlementID=%+v", d.SettlementID)

	builder.WriteByte('}')

	return builder.String()
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

func (d CThostFtdcOptionInstrMiniMarginField) Type() string {
	return "CThostFtdcOptionInstrMiniMarginField"
}

func (d CThostFtdcOptionInstrMiniMarginField) String() string {
	var builder strings.Builder
	builder.Grow(179)

	builder.WriteString("CThostFtdcOptionInstrMiniMarginField{")
	fmt.Fprintf(&builder, "InstrumentID=%+v", d.InstrumentID)
	fmt.Fprintf(&builder, ", InvestorRange=%+v", d.InvestorRange)
	fmt.Fprintf(&builder, ", BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", InvestorID=%+v", d.InvestorID)
	fmt.Fprintf(&builder, ", MinMargin=%+v", d.MinMargin)
	fmt.Fprintf(&builder, ", ValueMethod=%+v", d.ValueMethod)
	fmt.Fprintf(&builder, ", IsRelative=%+v", d.IsRelative)

	builder.WriteByte('}')

	return builder.String()
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

func (d CThostFtdcOptionInstrMarginAdjustField) Type() string {
	return "CThostFtdcOptionInstrMarginAdjustField"
}

func (d CThostFtdcOptionInstrMarginAdjustField) String() string {
	var builder strings.Builder
	builder.Grow(417)

	builder.WriteString("CThostFtdcOptionInstrMarginAdjustField{")
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

	builder.WriteByte('}')

	return builder.String()
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

func (d CThostFtdcOptionInstrCommRateField) Type() string {
	return "CThostFtdcOptionInstrCommRateField"
}

func (d CThostFtdcOptionInstrCommRateField) String() string {
	var builder strings.Builder
	builder.Grow(347)

	builder.WriteString("CThostFtdcOptionInstrCommRateField{")
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

	builder.WriteByte('}')

	return builder.String()
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

func (d CThostFtdcOptionInstrTradeCostField) Type() string {
	return "CThostFtdcOptionInstrTradeCostField"
}

func (d CThostFtdcOptionInstrTradeCostField) String() string {
	var builder strings.Builder
	builder.Grow(221)

	builder.WriteString("CThostFtdcOptionInstrTradeCostField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", InvestorID=%+v", d.InvestorID)
	fmt.Fprintf(&builder, ", InstrumentID=%+v", d.InstrumentID)
	fmt.Fprintf(&builder, ", HedgeFlag=%+v", d.HedgeFlag)
	fmt.Fprintf(&builder, ", FixedMargin=%+v", d.FixedMargin)
	fmt.Fprintf(&builder, ", MiniMargin=%+v", d.MiniMargin)
	fmt.Fprintf(&builder, ", Royalty=%+v", d.Royalty)
	fmt.Fprintf(&builder, ", ExchFixedMargin=%+v", d.ExchFixedMargin)
	fmt.Fprintf(&builder, ", ExchMiniMargin=%+v", d.ExchMiniMargin)

	builder.WriteByte('}')

	return builder.String()
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

func (d CThostFtdcQryOptionInstrTradeCostField) Type() string {
	return "CThostFtdcQryOptionInstrTradeCostField"
}

func (d CThostFtdcQryOptionInstrTradeCostField) String() string {
	var builder strings.Builder
	builder.Grow(162)

	builder.WriteString("CThostFtdcQryOptionInstrTradeCostField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", InvestorID=%+v", d.InvestorID)
	fmt.Fprintf(&builder, ", InstrumentID=%+v", d.InstrumentID)
	fmt.Fprintf(&builder, ", HedgeFlag=%+v", d.HedgeFlag)
	fmt.Fprintf(&builder, ", InputPrice=%+v", d.InputPrice)
	fmt.Fprintf(&builder, ", UnderlyingPrice=%+v", d.UnderlyingPrice)

	builder.WriteByte('}')

	return builder.String()
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

func (d CThostFtdcQryOptionInstrCommRateField) Type() string {
	return "CThostFtdcQryOptionInstrCommRateField"
}

func (d CThostFtdcQryOptionInstrCommRateField) String() string {
	var builder strings.Builder
	builder.Grow(97)

	builder.WriteString("CThostFtdcQryOptionInstrCommRateField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", InvestorID=%+v", d.InvestorID)
	fmt.Fprintf(&builder, ", InstrumentID=%+v", d.InstrumentID)

	builder.WriteByte('}')

	return builder.String()
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

func (d CThostFtdcIndexPriceField) Type() string { return "CThostFtdcIndexPriceField" }

func (d CThostFtdcIndexPriceField) String() string {
	var builder strings.Builder
	builder.Grow(85)

	builder.WriteString("CThostFtdcIndexPriceField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", InstrumentID=%+v", d.InstrumentID)
	fmt.Fprintf(&builder, ", ClosePrice=%+v", d.ClosePrice)

	builder.WriteByte('}')

	return builder.String()
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

func (d CThostFtdcInputExecOrderField) Type() string { return "CThostFtdcInputExecOrderField" }

func (d CThostFtdcInputExecOrderField) String() string {
	var builder strings.Builder
	builder.Grow(452)

	builder.WriteString("CThostFtdcInputExecOrderField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", InvestorID=%+v", d.InvestorID)
	fmt.Fprintf(&builder, ", InstrumentID=%+v", d.InstrumentID)
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
	fmt.Fprintf(&builder, ", IPAddress=%+v", d.IPAddress)
	fmt.Fprintf(&builder, ", MacAddress=%+v", d.MacAddress)

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
	// 合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
	// 投资单元代码
	InvestUnitID types.TThostFtdcInvestUnitIDType
	// IP地址
	IPAddress types.TThostFtdcIPAddressType
	// Mac地址
	MacAddress types.TThostFtdcMacAddressType
}

func (d CThostFtdcInputExecOrderActionField) Type() string {
	return "CThostFtdcInputExecOrderActionField"
}

func (d CThostFtdcInputExecOrderActionField) String() string {
	var builder strings.Builder
	builder.Grow(341)

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
	fmt.Fprintf(&builder, ", InstrumentID=%+v", d.InstrumentID)
	fmt.Fprintf(&builder, ", InvestUnitID=%+v", d.InvestUnitID)
	fmt.Fprintf(&builder, ", IPAddress=%+v", d.IPAddress)
	fmt.Fprintf(&builder, ", MacAddress=%+v", d.MacAddress)

	builder.WriteByte('}')

	return builder.String()
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

func (d CThostFtdcExecOrderField) Type() string { return "CThostFtdcExecOrderField" }

func (d CThostFtdcExecOrderField) String() string {
	var builder strings.Builder
	builder.Grow(946)

	builder.WriteString("CThostFtdcExecOrderField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", InvestorID=%+v", d.InvestorID)
	fmt.Fprintf(&builder, ", InstrumentID=%+v", d.InstrumentID)
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
	fmt.Fprintf(&builder, ", ExchangeInstID=%+v", d.ExchangeInstID)
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
	fmt.Fprintf(&builder, ", IPAddress=%+v", d.IPAddress)
	fmt.Fprintf(&builder, ", MacAddress=%+v", d.MacAddress)

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

func (d CThostFtdcExecOrderActionField) Type() string { return "CThostFtdcExecOrderActionField" }

func (d CThostFtdcExecOrderActionField) String() string {
	var builder strings.Builder
	builder.Grow(609)

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
	fmt.Fprintf(&builder, ", InstrumentID=%+v", d.InstrumentID)
	fmt.Fprintf(&builder, ", BranchID=%+v", d.BranchID)
	fmt.Fprintf(&builder, ", InvestUnitID=%+v", d.InvestUnitID)
	fmt.Fprintf(&builder, ", IPAddress=%+v", d.IPAddress)
	fmt.Fprintf(&builder, ", MacAddress=%+v", d.MacAddress)

	builder.WriteByte('}')

	return builder.String()
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

func (d CThostFtdcQryExecOrderField) Type() string { return "CThostFtdcQryExecOrderField" }

func (d CThostFtdcQryExecOrderField) String() string {
	var builder strings.Builder
	builder.Grow(179)

	builder.WriteString("CThostFtdcQryExecOrderField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", InvestorID=%+v", d.InvestorID)
	fmt.Fprintf(&builder, ", InstrumentID=%+v", d.InstrumentID)
	fmt.Fprintf(&builder, ", ExchangeID=%+v", d.ExchangeID)
	fmt.Fprintf(&builder, ", ExecOrderSysID=%+v", d.ExecOrderSysID)
	fmt.Fprintf(&builder, ", InsertTimeStart=%+v", d.InsertTimeStart)
	fmt.Fprintf(&builder, ", InsertTimeEnd=%+v", d.InsertTimeEnd)

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

func (d CThostFtdcExchangeExecOrderField) Type() string { return "CThostFtdcExchangeExecOrderField" }

func (d CThostFtdcExchangeExecOrderField) String() string {
	var builder strings.Builder
	builder.Grow(665)

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
	fmt.Fprintf(&builder, ", ExchangeInstID=%+v", d.ExchangeInstID)
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
	fmt.Fprintf(&builder, ", IPAddress=%+v", d.IPAddress)
	fmt.Fprintf(&builder, ", MacAddress=%+v", d.MacAddress)

	builder.WriteByte('}')

	return builder.String()
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

func (d CThostFtdcQryExchangeExecOrderField) Type() string {
	return "CThostFtdcQryExchangeExecOrderField"
}

func (d CThostFtdcQryExchangeExecOrderField) String() string {
	var builder strings.Builder
	builder.Grow(138)

	builder.WriteString("CThostFtdcQryExchangeExecOrderField{")
	fmt.Fprintf(&builder, "ParticipantID=%+v", d.ParticipantID)
	fmt.Fprintf(&builder, ", ClientID=%+v", d.ClientID)
	fmt.Fprintf(&builder, ", ExchangeInstID=%+v", d.ExchangeInstID)
	fmt.Fprintf(&builder, ", ExchangeID=%+v", d.ExchangeID)
	fmt.Fprintf(&builder, ", TraderID=%+v", d.TraderID)

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
	// IP地址
	IPAddress types.TThostFtdcIPAddressType
	// Mac地址
	MacAddress types.TThostFtdcMacAddressType
}

func (d CThostFtdcExchangeExecOrderActionField) Type() string {
	return "CThostFtdcExchangeExecOrderActionField"
}

func (d CThostFtdcExchangeExecOrderActionField) String() string {
	var builder strings.Builder
	builder.Grow(411)

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
	fmt.Fprintf(&builder, ", IPAddress=%+v", d.IPAddress)
	fmt.Fprintf(&builder, ", MacAddress=%+v", d.MacAddress)

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

func (d CThostFtdcErrExecOrderField) Type() string { return "CThostFtdcErrExecOrderField" }

func (d CThostFtdcErrExecOrderField) String() string {
	var builder strings.Builder
	builder.Grow(485)

	builder.WriteString("CThostFtdcErrExecOrderField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", InvestorID=%+v", d.InvestorID)
	fmt.Fprintf(&builder, ", InstrumentID=%+v", d.InstrumentID)
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
	fmt.Fprintf(&builder, ", IPAddress=%+v", d.IPAddress)
	fmt.Fprintf(&builder, ", MacAddress=%+v", d.MacAddress)
	fmt.Fprintf(&builder, ", ErrorID=%+v", d.ErrorID)
	fmt.Fprintf(&builder, ", ErrorMsg=%+v", d.ErrorMsg)

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

func (d CThostFtdcErrExecOrderActionField) Type() string { return "CThostFtdcErrExecOrderActionField" }

func (d CThostFtdcErrExecOrderActionField) String() string {
	var builder strings.Builder
	builder.Grow(374)

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
	fmt.Fprintf(&builder, ", InstrumentID=%+v", d.InstrumentID)
	fmt.Fprintf(&builder, ", InvestUnitID=%+v", d.InvestUnitID)
	fmt.Fprintf(&builder, ", IPAddress=%+v", d.IPAddress)
	fmt.Fprintf(&builder, ", MacAddress=%+v", d.MacAddress)
	fmt.Fprintf(&builder, ", ErrorID=%+v", d.ErrorID)
	fmt.Fprintf(&builder, ", ErrorMsg=%+v", d.ErrorMsg)

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

func (d CThostFtdcOptionInstrTradingRightField) Type() string {
	return "CThostFtdcOptionInstrTradingRightField"
}

func (d CThostFtdcOptionInstrTradingRightField) String() string {
	var builder strings.Builder
	builder.Grow(162)

	builder.WriteString("CThostFtdcOptionInstrTradingRightField{")
	fmt.Fprintf(&builder, "InstrumentID=%+v", d.InstrumentID)
	fmt.Fprintf(&builder, ", InvestorRange=%+v", d.InvestorRange)
	fmt.Fprintf(&builder, ", BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", InvestorID=%+v", d.InvestorID)
	fmt.Fprintf(&builder, ", Direction=%+v", d.Direction)
	fmt.Fprintf(&builder, ", TradingRight=%+v", d.TradingRight)

	builder.WriteByte('}')

	return builder.String()
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

func (d CThostFtdcQryOptionInstrTradingRightField) Type() string {
	return "CThostFtdcQryOptionInstrTradingRightField"
}

func (d CThostFtdcQryOptionInstrTradingRightField) String() string {
	var builder strings.Builder
	builder.Grow(120)

	builder.WriteString("CThostFtdcQryOptionInstrTradingRightField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", InvestorID=%+v", d.InvestorID)
	fmt.Fprintf(&builder, ", InstrumentID=%+v", d.InstrumentID)
	fmt.Fprintf(&builder, ", Direction=%+v", d.Direction)

	builder.WriteByte('}')

	return builder.String()
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

func (d CThostFtdcInputForQuoteField) Type() string { return "CThostFtdcInputForQuoteField" }

func (d CThostFtdcInputForQuoteField) String() string {
	var builder strings.Builder
	builder.Grow(228)

	builder.WriteString("CThostFtdcInputForQuoteField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", InvestorID=%+v", d.InvestorID)
	fmt.Fprintf(&builder, ", InstrumentID=%+v", d.InstrumentID)
	fmt.Fprintf(&builder, ", ForQuoteRef=%+v", d.ForQuoteRef)
	fmt.Fprintf(&builder, ", UserID=%+v", d.UserID)
	fmt.Fprintf(&builder, ", ExchangeID=%+v", d.ExchangeID)
	fmt.Fprintf(&builder, ", InvestUnitID=%+v", d.InvestUnitID)
	fmt.Fprintf(&builder, ", IPAddress=%+v", d.IPAddress)
	fmt.Fprintf(&builder, ", MacAddress=%+v", d.MacAddress)
	fmt.Fprintf(&builder, ", BusinessUnit=%+v", d.BusinessUnit)

	builder.WriteByte('}')

	return builder.String()
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

func (d CThostFtdcForQuoteField) Type() string { return "CThostFtdcForQuoteField" }

func (d CThostFtdcForQuoteField) String() string {
	var builder strings.Builder
	builder.Grow(517)

	builder.WriteString("CThostFtdcForQuoteField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", InvestorID=%+v", d.InvestorID)
	fmt.Fprintf(&builder, ", InstrumentID=%+v", d.InstrumentID)
	fmt.Fprintf(&builder, ", ForQuoteRef=%+v", d.ForQuoteRef)
	fmt.Fprintf(&builder, ", UserID=%+v", d.UserID)
	fmt.Fprintf(&builder, ", ForQuoteLocalID=%+v", d.ForQuoteLocalID)
	fmt.Fprintf(&builder, ", ExchangeID=%+v", d.ExchangeID)
	fmt.Fprintf(&builder, ", ParticipantID=%+v", d.ParticipantID)
	fmt.Fprintf(&builder, ", ClientID=%+v", d.ClientID)
	fmt.Fprintf(&builder, ", ExchangeInstID=%+v", d.ExchangeInstID)
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
	fmt.Fprintf(&builder, ", IPAddress=%+v", d.IPAddress)
	fmt.Fprintf(&builder, ", MacAddress=%+v", d.MacAddress)
	fmt.Fprintf(&builder, ", BusinessUnit=%+v", d.BusinessUnit)

	builder.WriteByte('}')

	return builder.String()
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

func (d CThostFtdcQryForQuoteField) Type() string { return "CThostFtdcQryForQuoteField" }

func (d CThostFtdcQryForQuoteField) String() string {
	var builder strings.Builder
	builder.Grow(154)

	builder.WriteString("CThostFtdcQryForQuoteField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", InvestorID=%+v", d.InvestorID)
	fmt.Fprintf(&builder, ", InstrumentID=%+v", d.InstrumentID)
	fmt.Fprintf(&builder, ", ExchangeID=%+v", d.ExchangeID)
	fmt.Fprintf(&builder, ", InsertTimeStart=%+v", d.InsertTimeStart)
	fmt.Fprintf(&builder, ", InsertTimeEnd=%+v", d.InsertTimeEnd)

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

func (d CThostFtdcExchangeForQuoteField) Type() string { return "CThostFtdcExchangeForQuoteField" }

func (d CThostFtdcExchangeForQuoteField) String() string {
	var builder strings.Builder
	builder.Grow(281)

	builder.WriteString("CThostFtdcExchangeForQuoteField{")
	fmt.Fprintf(&builder, "ForQuoteLocalID=%+v", d.ForQuoteLocalID)
	fmt.Fprintf(&builder, ", ExchangeID=%+v", d.ExchangeID)
	fmt.Fprintf(&builder, ", ParticipantID=%+v", d.ParticipantID)
	fmt.Fprintf(&builder, ", ClientID=%+v", d.ClientID)
	fmt.Fprintf(&builder, ", ExchangeInstID=%+v", d.ExchangeInstID)
	fmt.Fprintf(&builder, ", TraderID=%+v", d.TraderID)
	fmt.Fprintf(&builder, ", InstallID=%+v", d.InstallID)
	fmt.Fprintf(&builder, ", InsertDate=%+v", d.InsertDate)
	fmt.Fprintf(&builder, ", InsertTime=%+v", d.InsertTime)
	fmt.Fprintf(&builder, ", ForQuoteStatus=%+v", d.ForQuoteStatus)
	fmt.Fprintf(&builder, ", IPAddress=%+v", d.IPAddress)
	fmt.Fprintf(&builder, ", MacAddress=%+v", d.MacAddress)

	builder.WriteByte('}')

	return builder.String()
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

func (d CThostFtdcQryExchangeForQuoteField) Type() string {
	return "CThostFtdcQryExchangeForQuoteField"
}

func (d CThostFtdcQryExchangeForQuoteField) String() string {
	var builder strings.Builder
	builder.Grow(137)

	builder.WriteString("CThostFtdcQryExchangeForQuoteField{")
	fmt.Fprintf(&builder, "ParticipantID=%+v", d.ParticipantID)
	fmt.Fprintf(&builder, ", ClientID=%+v", d.ClientID)
	fmt.Fprintf(&builder, ", ExchangeInstID=%+v", d.ExchangeInstID)
	fmt.Fprintf(&builder, ", ExchangeID=%+v", d.ExchangeID)
	fmt.Fprintf(&builder, ", TraderID=%+v", d.TraderID)

	builder.WriteByte('}')

	return builder.String()
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

func (d CThostFtdcInputQuoteField) Type() string { return "CThostFtdcInputQuoteField" }

func (d CThostFtdcInputQuoteField) String() string {
	var builder strings.Builder
	builder.Grow(557)

	builder.WriteString("CThostFtdcInputQuoteField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", InvestorID=%+v", d.InvestorID)
	fmt.Fprintf(&builder, ", InstrumentID=%+v", d.InstrumentID)
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
	fmt.Fprintf(&builder, ", IPAddress=%+v", d.IPAddress)
	fmt.Fprintf(&builder, ", MacAddress=%+v", d.MacAddress)
	fmt.Fprintf(&builder, ", ReplaceSysID=%+v", d.ReplaceSysID)
	fmt.Fprintf(&builder, ", CustomQuoteRef=%+v", d.CustomQuoteRef)
	fmt.Fprintf(&builder, ", TimeCondition=%+v", d.TimeCondition)

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

func (d CThostFtdcInputQuoteActionField) Type() string { return "CThostFtdcInputQuoteActionField" }

func (d CThostFtdcInputQuoteActionField) String() string {
	var builder strings.Builder
	builder.Grow(343)

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
	fmt.Fprintf(&builder, ", InstrumentID=%+v", d.InstrumentID)
	fmt.Fprintf(&builder, ", InvestUnitID=%+v", d.InvestUnitID)
	fmt.Fprintf(&builder, ", ClientID=%+v", d.ClientID)
	fmt.Fprintf(&builder, ", IPAddress=%+v", d.IPAddress)
	fmt.Fprintf(&builder, ", MacAddress=%+v", d.MacAddress)

	builder.WriteByte('}')

	return builder.String()
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

func (d CThostFtdcQuoteField) Type() string { return "CThostFtdcQuoteField" }

func (d CThostFtdcQuoteField) String() string {
	var builder strings.Builder
	builder.Grow(1125)

	builder.WriteString("CThostFtdcQuoteField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", InvestorID=%+v", d.InvestorID)
	fmt.Fprintf(&builder, ", InstrumentID=%+v", d.InstrumentID)
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
	fmt.Fprintf(&builder, ", ExchangeInstID=%+v", d.ExchangeInstID)
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
	fmt.Fprintf(&builder, ", IPAddress=%+v", d.IPAddress)
	fmt.Fprintf(&builder, ", MacAddress=%+v", d.MacAddress)
	fmt.Fprintf(&builder, ", ReplaceSysID=%+v", d.ReplaceSysID)
	fmt.Fprintf(&builder, ", CustomQuoteRef=%+v", d.CustomQuoteRef)
	fmt.Fprintf(&builder, ", TimeCondition=%+v", d.TimeCondition)

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

func (d CThostFtdcQuoteActionField) Type() string { return "CThostFtdcQuoteActionField" }

func (d CThostFtdcQuoteActionField) String() string {
	var builder strings.Builder
	builder.Grow(569)

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
	fmt.Fprintf(&builder, ", InstrumentID=%+v", d.InstrumentID)
	fmt.Fprintf(&builder, ", BranchID=%+v", d.BranchID)
	fmt.Fprintf(&builder, ", InvestUnitID=%+v", d.InvestUnitID)
	fmt.Fprintf(&builder, ", IPAddress=%+v", d.IPAddress)
	fmt.Fprintf(&builder, ", MacAddress=%+v", d.MacAddress)

	builder.WriteByte('}')

	return builder.String()
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

func (d CThostFtdcQryQuoteField) Type() string { return "CThostFtdcQryQuoteField" }

func (d CThostFtdcQryQuoteField) String() string {
	var builder strings.Builder
	builder.Grow(171)

	builder.WriteString("CThostFtdcQryQuoteField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", InvestorID=%+v", d.InvestorID)
	fmt.Fprintf(&builder, ", InstrumentID=%+v", d.InstrumentID)
	fmt.Fprintf(&builder, ", ExchangeID=%+v", d.ExchangeID)
	fmt.Fprintf(&builder, ", QuoteSysID=%+v", d.QuoteSysID)
	fmt.Fprintf(&builder, ", InsertTimeStart=%+v", d.InsertTimeStart)
	fmt.Fprintf(&builder, ", InsertTimeEnd=%+v", d.InsertTimeEnd)

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

func (d CThostFtdcExchangeQuoteField) Type() string { return "CThostFtdcExchangeQuoteField" }

func (d CThostFtdcExchangeQuoteField) String() string {
	var builder strings.Builder
	builder.Grow(741)

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
	fmt.Fprintf(&builder, ", ExchangeInstID=%+v", d.ExchangeInstID)
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
	fmt.Fprintf(&builder, ", IPAddress=%+v", d.IPAddress)
	fmt.Fprintf(&builder, ", MacAddress=%+v", d.MacAddress)

	builder.WriteByte('}')

	return builder.String()
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

func (d CThostFtdcQryExchangeQuoteField) Type() string { return "CThostFtdcQryExchangeQuoteField" }

func (d CThostFtdcQryExchangeQuoteField) String() string {
	var builder strings.Builder
	builder.Grow(134)

	builder.WriteString("CThostFtdcQryExchangeQuoteField{")
	fmt.Fprintf(&builder, "ParticipantID=%+v", d.ParticipantID)
	fmt.Fprintf(&builder, ", ClientID=%+v", d.ClientID)
	fmt.Fprintf(&builder, ", ExchangeInstID=%+v", d.ExchangeInstID)
	fmt.Fprintf(&builder, ", ExchangeID=%+v", d.ExchangeID)
	fmt.Fprintf(&builder, ", TraderID=%+v", d.TraderID)

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
	// IP地址
	IPAddress types.TThostFtdcIPAddressType
	// Mac地址
	MacAddress types.TThostFtdcMacAddressType
}

func (d CThostFtdcExchangeQuoteActionField) Type() string {
	return "CThostFtdcExchangeQuoteActionField"
}

func (d CThostFtdcExchangeQuoteActionField) String() string {
	var builder strings.Builder
	builder.Grow(361)

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
	fmt.Fprintf(&builder, ", IPAddress=%+v", d.IPAddress)
	fmt.Fprintf(&builder, ", MacAddress=%+v", d.MacAddress)

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

func (d CThostFtdcOptionInstrDeltaField) Type() string { return "CThostFtdcOptionInstrDeltaField" }

func (d CThostFtdcOptionInstrDeltaField) String() string {
	var builder strings.Builder
	builder.Grow(129)

	builder.WriteString("CThostFtdcOptionInstrDeltaField{")
	fmt.Fprintf(&builder, "InstrumentID=%+v", d.InstrumentID)
	fmt.Fprintf(&builder, ", InvestorRange=%+v", d.InvestorRange)
	fmt.Fprintf(&builder, ", BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", InvestorID=%+v", d.InvestorID)
	fmt.Fprintf(&builder, ", Delta=%+v", d.Delta)

	builder.WriteByte('}')

	return builder.String()
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

func (d CThostFtdcForQuoteRspField) Type() string { return "CThostFtdcForQuoteRspField" }

func (d CThostFtdcForQuoteRspField) String() string {
	var builder strings.Builder
	builder.Grow(152)

	builder.WriteString("CThostFtdcForQuoteRspField{")
	fmt.Fprintf(&builder, "TradingDay=%+v", d.TradingDay)
	fmt.Fprintf(&builder, ", InstrumentID=%+v", d.InstrumentID)
	fmt.Fprintf(&builder, ", ForQuoteSysID=%+v", d.ForQuoteSysID)
	fmt.Fprintf(&builder, ", ForQuoteTime=%+v", d.ForQuoteTime)
	fmt.Fprintf(&builder, ", ActionDay=%+v", d.ActionDay)
	fmt.Fprintf(&builder, ", ExchangeID=%+v", d.ExchangeID)

	builder.WriteByte('}')

	return builder.String()
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

func (d CThostFtdcStrikeOffsetField) Type() string { return "CThostFtdcStrikeOffsetField" }

func (d CThostFtdcStrikeOffsetField) String() string {
	var builder strings.Builder
	builder.Grow(146)

	builder.WriteString("CThostFtdcStrikeOffsetField{")
	fmt.Fprintf(&builder, "InstrumentID=%+v", d.InstrumentID)
	fmt.Fprintf(&builder, ", InvestorRange=%+v", d.InvestorRange)
	fmt.Fprintf(&builder, ", BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", InvestorID=%+v", d.InvestorID)
	fmt.Fprintf(&builder, ", Offset=%+v", d.Offset)
	fmt.Fprintf(&builder, ", OffsetType=%+v", d.OffsetType)

	builder.WriteByte('}')

	return builder.String()
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

func (d CThostFtdcQryStrikeOffsetField) Type() string { return "CThostFtdcQryStrikeOffsetField" }

func (d CThostFtdcQryStrikeOffsetField) String() string {
	var builder strings.Builder
	builder.Grow(90)

	builder.WriteString("CThostFtdcQryStrikeOffsetField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", InvestorID=%+v", d.InvestorID)
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
	// IP地址
	IPAddress types.TThostFtdcIPAddressType
	// Mac地址
	MacAddress types.TThostFtdcMacAddressType
}

func (d CThostFtdcInputBatchOrderActionField) Type() string {
	return "CThostFtdcInputBatchOrderActionField"
}

func (d CThostFtdcInputBatchOrderActionField) String() string {
	var builder strings.Builder
	builder.Grow(250)

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
	fmt.Fprintf(&builder, ", IPAddress=%+v", d.IPAddress)
	fmt.Fprintf(&builder, ", MacAddress=%+v", d.MacAddress)

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
	// IP地址
	IPAddress types.TThostFtdcIPAddressType
	// Mac地址
	MacAddress types.TThostFtdcMacAddressType
}

func (d CThostFtdcBatchOrderActionField) Type() string { return "CThostFtdcBatchOrderActionField" }

func (d CThostFtdcBatchOrderActionField) String() string {
	var builder strings.Builder
	builder.Grow(454)

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
	fmt.Fprintf(&builder, ", IPAddress=%+v", d.IPAddress)
	fmt.Fprintf(&builder, ", MacAddress=%+v", d.MacAddress)

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
	// IP地址
	IPAddress types.TThostFtdcIPAddressType
	// Mac地址
	MacAddress types.TThostFtdcMacAddressType
}

func (d CThostFtdcExchangeBatchOrderActionField) Type() string {
	return "CThostFtdcExchangeBatchOrderActionField"
}

func (d CThostFtdcExchangeBatchOrderActionField) String() string {
	var builder strings.Builder
	builder.Grow(304)

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
	fmt.Fprintf(&builder, ", IPAddress=%+v", d.IPAddress)
	fmt.Fprintf(&builder, ", MacAddress=%+v", d.MacAddress)

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

func (d CThostFtdcInputOptionSelfCloseField) Type() string {
	return "CThostFtdcInputOptionSelfCloseField"
}

func (d CThostFtdcInputOptionSelfCloseField) String() string {
	var builder strings.Builder
	builder.Grow(379)

	builder.WriteString("CThostFtdcInputOptionSelfCloseField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", InvestorID=%+v", d.InvestorID)
	fmt.Fprintf(&builder, ", InstrumentID=%+v", d.InstrumentID)
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
	fmt.Fprintf(&builder, ", IPAddress=%+v", d.IPAddress)
	fmt.Fprintf(&builder, ", MacAddress=%+v", d.MacAddress)

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
	// 合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
	// 投资单元代码
	InvestUnitID types.TThostFtdcInvestUnitIDType
	// IP地址
	IPAddress types.TThostFtdcIPAddressType
	// Mac地址
	MacAddress types.TThostFtdcMacAddressType
}

func (d CThostFtdcInputOptionSelfCloseActionField) Type() string {
	return "CThostFtdcInputOptionSelfCloseActionField"
}

func (d CThostFtdcInputOptionSelfCloseActionField) String() string {
	var builder strings.Builder
	builder.Grow(365)

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
	fmt.Fprintf(&builder, ", InstrumentID=%+v", d.InstrumentID)
	fmt.Fprintf(&builder, ", InvestUnitID=%+v", d.InvestUnitID)
	fmt.Fprintf(&builder, ", IPAddress=%+v", d.IPAddress)
	fmt.Fprintf(&builder, ", MacAddress=%+v", d.MacAddress)

	builder.WriteByte('}')

	return builder.String()
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

func (d CThostFtdcOptionSelfCloseField) Type() string { return "CThostFtdcOptionSelfCloseField" }

func (d CThostFtdcOptionSelfCloseField) String() string {
	var builder strings.Builder
	builder.Grow(891)

	builder.WriteString("CThostFtdcOptionSelfCloseField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", InvestorID=%+v", d.InvestorID)
	fmt.Fprintf(&builder, ", InstrumentID=%+v", d.InstrumentID)
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
	fmt.Fprintf(&builder, ", ExchangeInstID=%+v", d.ExchangeInstID)
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
	fmt.Fprintf(&builder, ", IPAddress=%+v", d.IPAddress)
	fmt.Fprintf(&builder, ", MacAddress=%+v", d.MacAddress)

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

func (d CThostFtdcOptionSelfCloseActionField) Type() string {
	return "CThostFtdcOptionSelfCloseActionField"
}

func (d CThostFtdcOptionSelfCloseActionField) String() string {
	var builder strings.Builder
	builder.Grow(619)

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
	fmt.Fprintf(&builder, ", InstrumentID=%+v", d.InstrumentID)
	fmt.Fprintf(&builder, ", BranchID=%+v", d.BranchID)
	fmt.Fprintf(&builder, ", InvestUnitID=%+v", d.InvestUnitID)
	fmt.Fprintf(&builder, ", IPAddress=%+v", d.IPAddress)
	fmt.Fprintf(&builder, ", MacAddress=%+v", d.MacAddress)

	builder.WriteByte('}')

	return builder.String()
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

func (d CThostFtdcQryOptionSelfCloseField) Type() string { return "CThostFtdcQryOptionSelfCloseField" }

func (d CThostFtdcQryOptionSelfCloseField) String() string {
	var builder strings.Builder
	builder.Grow(191)

	builder.WriteString("CThostFtdcQryOptionSelfCloseField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", InvestorID=%+v", d.InvestorID)
	fmt.Fprintf(&builder, ", InstrumentID=%+v", d.InstrumentID)
	fmt.Fprintf(&builder, ", ExchangeID=%+v", d.ExchangeID)
	fmt.Fprintf(&builder, ", OptionSelfCloseSysID=%+v", d.OptionSelfCloseSysID)
	fmt.Fprintf(&builder, ", InsertTimeStart=%+v", d.InsertTimeStart)
	fmt.Fprintf(&builder, ", InsertTimeEnd=%+v", d.InsertTimeEnd)

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

func (d CThostFtdcExchangeOptionSelfCloseField) Type() string {
	return "CThostFtdcExchangeOptionSelfCloseField"
}

func (d CThostFtdcExchangeOptionSelfCloseField) String() string {
	var builder strings.Builder
	builder.Grow(598)

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
	fmt.Fprintf(&builder, ", ExchangeInstID=%+v", d.ExchangeInstID)
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
	fmt.Fprintf(&builder, ", IPAddress=%+v", d.IPAddress)
	fmt.Fprintf(&builder, ", MacAddress=%+v", d.MacAddress)

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
	// IP地址
	IPAddress types.TThostFtdcIPAddressType
	// Mac地址
	MacAddress types.TThostFtdcMacAddressType
}

func (d CThostFtdcExchangeOptionSelfCloseActionField) Type() string {
	return "CThostFtdcExchangeOptionSelfCloseActionField"
}

func (d CThostFtdcExchangeOptionSelfCloseActionField) String() string {
	var builder strings.Builder
	builder.Grow(409)

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
	fmt.Fprintf(&builder, ", IPAddress=%+v", d.IPAddress)
	fmt.Fprintf(&builder, ", MacAddress=%+v", d.MacAddress)

	builder.WriteByte('}')

	return builder.String()
}

// 组合合约安全系数
type CThostFtdcCombInstrumentGuardField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
	GuarantRatio types.TThostFtdcRatioType
}

func (d CThostFtdcCombInstrumentGuardField) Type() string {
	return "CThostFtdcCombInstrumentGuardField"
}

func (d CThostFtdcCombInstrumentGuardField) String() string {
	var builder strings.Builder
	builder.Grow(96)

	builder.WriteString("CThostFtdcCombInstrumentGuardField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", InstrumentID=%+v", d.InstrumentID)
	fmt.Fprintf(&builder, ", GuarantRatio=%+v", d.GuarantRatio)

	builder.WriteByte('}')

	return builder.String()
}

// 组合合约安全系数查询
type CThostFtdcQryCombInstrumentGuardField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
}

func (d CThostFtdcQryCombInstrumentGuardField) Type() string {
	return "CThostFtdcQryCombInstrumentGuardField"
}

func (d CThostFtdcQryCombInstrumentGuardField) String() string {
	var builder strings.Builder
	builder.Grow(77)

	builder.WriteString("CThostFtdcQryCombInstrumentGuardField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", InstrumentID=%+v", d.InstrumentID)

	builder.WriteByte('}')

	return builder.String()
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

func (d CThostFtdcInputHedgeConfirmField) Type() string { return "CThostFtdcInputHedgeConfirmField" }

func (d CThostFtdcInputHedgeConfirmField) String() string {
	var builder strings.Builder
	builder.Grow(345)

	builder.WriteString("CThostFtdcInputHedgeConfirmField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", InvestorID=%+v", d.InvestorID)
	fmt.Fprintf(&builder, ", InstrumentID=%+v", d.InstrumentID)
	fmt.Fprintf(&builder, ", OrderRef=%+v", d.OrderRef)
	fmt.Fprintf(&builder, ", UserID=%+v", d.UserID)
	fmt.Fprintf(&builder, ", Direction=%+v", d.Direction)
	fmt.Fprintf(&builder, ", Volume=%+v", d.Volume)
	fmt.Fprintf(&builder, ", RequestType=%+v", d.RequestType)
	fmt.Fprintf(&builder, ", HedgeFlag=%+v", d.HedgeFlag)
	fmt.Fprintf(&builder, ", BusinessUnit=%+v", d.BusinessUnit)
	fmt.Fprintf(&builder, ", ExchangeID=%+v", d.ExchangeID)
	fmt.Fprintf(&builder, ", SecondLeg=%+v", d.SecondLeg)
	fmt.Fprintf(&builder, ", CombinationType=%+v", d.CombinationType)
	fmt.Fprintf(&builder, ", InvestUnitID=%+v", d.InvestUnitID)
	fmt.Fprintf(&builder, ", FrontID=%+v", d.FrontID)
	fmt.Fprintf(&builder, ", SessionID=%+v", d.SessionID)

	builder.WriteByte('}')

	return builder.String()
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

func (d CThostFtdcInputHedgeConfirmActionField) Type() string {
	return "CThostFtdcInputHedgeConfirmActionField"
}

func (d CThostFtdcInputHedgeConfirmActionField) String() string {
	var builder strings.Builder
	builder.Grow(353)

	builder.WriteString("CThostFtdcInputHedgeConfirmActionField{")
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
	fmt.Fprintf(&builder, ", UserID=%+v", d.UserID)
	fmt.Fprintf(&builder, ", RequestType=%+v", d.RequestType)
	fmt.Fprintf(&builder, ", InstrumentID=%+v", d.InstrumentID)
	fmt.Fprintf(&builder, ", InvestUnitID=%+v", d.InvestUnitID)
	fmt.Fprintf(&builder, ", IPAddress=%+v", d.IPAddress)
	fmt.Fprintf(&builder, ", MacAddress=%+v", d.MacAddress)

	builder.WriteByte('}')

	return builder.String()
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

func (d CThostFtdcHedgeConfirmField) Type() string { return "CThostFtdcHedgeConfirmField" }

func (d CThostFtdcHedgeConfirmField) String() string {
	var builder strings.Builder
	builder.Grow(529)

	builder.WriteString("CThostFtdcHedgeConfirmField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", InvestorID=%+v", d.InvestorID)
	fmt.Fprintf(&builder, ", InstrumentID=%+v", d.InstrumentID)
	fmt.Fprintf(&builder, ", OrderRef=%+v", d.OrderRef)
	fmt.Fprintf(&builder, ", UserID=%+v", d.UserID)
	fmt.Fprintf(&builder, ", Direction=%+v", d.Direction)
	fmt.Fprintf(&builder, ", Volume=%+v", d.Volume)
	fmt.Fprintf(&builder, ", OrderSysID=%+v", d.OrderSysID)
	fmt.Fprintf(&builder, ", RequestType=%+v", d.RequestType)
	fmt.Fprintf(&builder, ", HedgeFlag=%+v", d.HedgeFlag)
	fmt.Fprintf(&builder, ", BusinessUnit=%+v", d.BusinessUnit)
	fmt.Fprintf(&builder, ", OrderLocalID=%+v", d.OrderLocalID)
	fmt.Fprintf(&builder, ", TraderID=%+v", d.TraderID)
	fmt.Fprintf(&builder, ", ExchangeID=%+v", d.ExchangeID)
	fmt.Fprintf(&builder, ", ParticipantID=%+v", d.ParticipantID)
	fmt.Fprintf(&builder, ", ClientID=%+v", d.ClientID)
	fmt.Fprintf(&builder, ", SecondLeg=%+v", d.SecondLeg)
	fmt.Fprintf(&builder, ", RequestStatus=%+v", d.RequestStatus)
	fmt.Fprintf(&builder, ", CombInstrumentID=%+v", d.CombInstrumentID)
	fmt.Fprintf(&builder, ", StatusMsg=%+v", d.StatusMsg)
	fmt.Fprintf(&builder, ", FrontID=%+v", d.FrontID)
	fmt.Fprintf(&builder, ", SessionID=%+v", d.SessionID)
	fmt.Fprintf(&builder, ", InvestUnitID=%+v", d.InvestUnitID)
	fmt.Fprintf(&builder, ", CombinationType=%+v", d.CombinationType)
	fmt.Fprintf(&builder, ", InsertDate=%+v", d.InsertDate)

	builder.WriteByte('}')

	return builder.String()
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

func (d CThostFtdcHedgeConfirmActionField) Type() string { return "CThostFtdcHedgeConfirmActionField" }

func (d CThostFtdcHedgeConfirmActionField) String() string {
	var builder strings.Builder
	builder.Grow(602)

	builder.WriteString("CThostFtdcHedgeConfirmActionField{")
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
	fmt.Fprintf(&builder, ", TraderID=%+v", d.TraderID)
	fmt.Fprintf(&builder, ", InstallID=%+v", d.InstallID)
	fmt.Fprintf(&builder, ", OrderLocalID=%+v", d.OrderLocalID)
	fmt.Fprintf(&builder, ", ActionLocalID=%+v", d.ActionLocalID)
	fmt.Fprintf(&builder, ", ParticipantID=%+v", d.ParticipantID)
	fmt.Fprintf(&builder, ", ClientID=%+v", d.ClientID)
	fmt.Fprintf(&builder, ", BusinessUnit=%+v", d.BusinessUnit)
	fmt.Fprintf(&builder, ", OrderActionStatus=%+v", d.OrderActionStatus)
	fmt.Fprintf(&builder, ", UserID=%+v", d.UserID)
	fmt.Fprintf(&builder, ", RequestType=%+v", d.RequestType)
	fmt.Fprintf(&builder, ", StatusMsg=%+v", d.StatusMsg)
	fmt.Fprintf(&builder, ", InstrumentID=%+v", d.InstrumentID)
	fmt.Fprintf(&builder, ", SecondLeg=%+v", d.SecondLeg)
	fmt.Fprintf(&builder, ", CombInstrumentID=%+v", d.CombInstrumentID)
	fmt.Fprintf(&builder, ", BranchID=%+v", d.BranchID)
	fmt.Fprintf(&builder, ", InvestUnitID=%+v", d.InvestUnitID)
	fmt.Fprintf(&builder, ", IPAddress=%+v", d.IPAddress)
	fmt.Fprintf(&builder, ", MacAddress=%+v", d.MacAddress)

	builder.WriteByte('}')

	return builder.String()
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

func (d CThostFtdcInputCombActionField) Type() string { return "CThostFtdcInputCombActionField" }

func (d CThostFtdcInputCombActionField) String() string {
	var builder strings.Builder
	builder.Grow(345)

	builder.WriteString("CThostFtdcInputCombActionField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", InvestorID=%+v", d.InvestorID)
	fmt.Fprintf(&builder, ", InstrumentID=%+v", d.InstrumentID)
	fmt.Fprintf(&builder, ", CombActionRef=%+v", d.CombActionRef)
	fmt.Fprintf(&builder, ", UserID=%+v", d.UserID)
	fmt.Fprintf(&builder, ", Direction=%+v", d.Direction)
	fmt.Fprintf(&builder, ", Volume=%+v", d.Volume)
	fmt.Fprintf(&builder, ", CombDirection=%+v", d.CombDirection)
	fmt.Fprintf(&builder, ", HedgeFlag=%+v", d.HedgeFlag)
	fmt.Fprintf(&builder, ", BusinessUnit=%+v", d.BusinessUnit)
	fmt.Fprintf(&builder, ", ExchangeID=%+v", d.ExchangeID)
	fmt.Fprintf(&builder, ", IPAddress=%+v", d.IPAddress)
	fmt.Fprintf(&builder, ", MacAddress=%+v", d.MacAddress)
	fmt.Fprintf(&builder, ", InvestUnitID=%+v", d.InvestUnitID)
	fmt.Fprintf(&builder, ", FrontID=%+v", d.FrontID)
	fmt.Fprintf(&builder, ", SessionID=%+v", d.SessionID)

	builder.WriteByte('}')

	return builder.String()
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

func (d CThostFtdcCombActionField) Type() string { return "CThostFtdcCombActionField" }

func (d CThostFtdcCombActionField) String() string {
	var builder strings.Builder
	builder.Grow(615)

	builder.WriteString("CThostFtdcCombActionField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", InvestorID=%+v", d.InvestorID)
	fmt.Fprintf(&builder, ", InstrumentID=%+v", d.InstrumentID)
	fmt.Fprintf(&builder, ", CombActionRef=%+v", d.CombActionRef)
	fmt.Fprintf(&builder, ", UserID=%+v", d.UserID)
	fmt.Fprintf(&builder, ", Direction=%+v", d.Direction)
	fmt.Fprintf(&builder, ", Volume=%+v", d.Volume)
	fmt.Fprintf(&builder, ", CombDirection=%+v", d.CombDirection)
	fmt.Fprintf(&builder, ", HedgeFlag=%+v", d.HedgeFlag)
	fmt.Fprintf(&builder, ", BusinessUnit=%+v", d.BusinessUnit)
	fmt.Fprintf(&builder, ", ActionLocalID=%+v", d.ActionLocalID)
	fmt.Fprintf(&builder, ", ExchangeID=%+v", d.ExchangeID)
	fmt.Fprintf(&builder, ", ParticipantID=%+v", d.ParticipantID)
	fmt.Fprintf(&builder, ", ClientID=%+v", d.ClientID)
	fmt.Fprintf(&builder, ", ExchangeInstID=%+v", d.ExchangeInstID)
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
	fmt.Fprintf(&builder, ", IPAddress=%+v", d.IPAddress)
	fmt.Fprintf(&builder, ", MacAddress=%+v", d.MacAddress)
	fmt.Fprintf(&builder, ", ComTradeID=%+v", d.ComTradeID)

	builder.WriteByte('}')

	return builder.String()
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

func (d CThostFtdcInvestorPositionForCombField) Type() string {
	return "CThostFtdcInvestorPositionForCombField"
}

func (d CThostFtdcInvestorPositionForCombField) String() string {
	var builder strings.Builder
	builder.Grow(350)

	builder.WriteString("CThostFtdcInvestorPositionForCombField{")
	fmt.Fprintf(&builder, "ExchangeID=%+v", d.ExchangeID)
	fmt.Fprintf(&builder, ", BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", InvestorID=%+v", d.InvestorID)
	fmt.Fprintf(&builder, ", LegInstrumentID=%+v", d.LegInstrumentID)
	fmt.Fprintf(&builder, ", LegHedgeFlag=%+v", d.LegHedgeFlag)
	fmt.Fprintf(&builder, ", LegDirection=%+v", d.LegDirection)
	fmt.Fprintf(&builder, ", TotalAmt=%+v", d.TotalAmt)
	fmt.Fprintf(&builder, ", LegID=%+v", d.LegID)
	fmt.Fprintf(&builder, ", TradeGroupID=%+v", d.TradeGroupID)
	fmt.Fprintf(&builder, ", CombInstrumentID=%+v", d.CombInstrumentID)
	fmt.Fprintf(&builder, ", CombHedgeFlag=%+v", d.CombHedgeFlag)
	fmt.Fprintf(&builder, ", CombinationType=%+v", d.CombinationType)
	fmt.Fprintf(&builder, ", CombTradeID=%+v", d.CombTradeID)
	fmt.Fprintf(&builder, ", OpenDate=%+v", d.OpenDate)
	fmt.Fprintf(&builder, ", TradeID=%+v", d.TradeID)

	builder.WriteByte('}')

	return builder.String()
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

func (d CThostFtdcQryCombActionField) Type() string { return "CThostFtdcQryCombActionField" }

func (d CThostFtdcQryCombActionField) String() string {
	var builder strings.Builder
	builder.Grow(108)

	builder.WriteString("CThostFtdcQryCombActionField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", InvestorID=%+v", d.InvestorID)
	fmt.Fprintf(&builder, ", InstrumentID=%+v", d.InstrumentID)
	fmt.Fprintf(&builder, ", ExchangeID=%+v", d.ExchangeID)

	builder.WriteByte('}')

	return builder.String()
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

func (d CThostFtdcQryHedgeConfirmField) Type() string { return "CThostFtdcQryHedgeConfirmField" }

func (d CThostFtdcQryHedgeConfirmField) String() string {
	var builder strings.Builder
	builder.Grow(110)

	builder.WriteString("CThostFtdcQryHedgeConfirmField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", InvestorID=%+v", d.InvestorID)
	fmt.Fprintf(&builder, ", InstrumentID=%+v", d.InstrumentID)
	fmt.Fprintf(&builder, ", ExchangeID=%+v", d.ExchangeID)

	builder.WriteByte('}')

	return builder.String()
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

func (d CThostFtdcQryInvestorPositionForCombField) Type() string {
	return "CThostFtdcQryInvestorPositionForCombField"
}

func (d CThostFtdcQryInvestorPositionForCombField) String() string {
	var builder strings.Builder
	builder.Grow(124)

	builder.WriteString("CThostFtdcQryInvestorPositionForCombField{")
	fmt.Fprintf(&builder, "ExchangeID=%+v", d.ExchangeID)
	fmt.Fprintf(&builder, ", BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", InvestorID=%+v", d.InvestorID)
	fmt.Fprintf(&builder, ", LegInstrumentID=%+v", d.LegInstrumentID)

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

func (d CThostFtdcExchangeCombActionField) Type() string { return "CThostFtdcExchangeCombActionField" }

func (d CThostFtdcExchangeCombActionField) String() string {
	var builder strings.Builder
	builder.Grow(402)

	builder.WriteString("CThostFtdcExchangeCombActionField{")
	fmt.Fprintf(&builder, "Direction=%+v", d.Direction)
	fmt.Fprintf(&builder, ", Volume=%+v", d.Volume)
	fmt.Fprintf(&builder, ", CombDirection=%+v", d.CombDirection)
	fmt.Fprintf(&builder, ", HedgeFlag=%+v", d.HedgeFlag)
	fmt.Fprintf(&builder, ", ActionLocalID=%+v", d.ActionLocalID)
	fmt.Fprintf(&builder, ", ExchangeID=%+v", d.ExchangeID)
	fmt.Fprintf(&builder, ", ParticipantID=%+v", d.ParticipantID)
	fmt.Fprintf(&builder, ", ClientID=%+v", d.ClientID)
	fmt.Fprintf(&builder, ", ExchangeInstID=%+v", d.ExchangeInstID)
	fmt.Fprintf(&builder, ", TraderID=%+v", d.TraderID)
	fmt.Fprintf(&builder, ", InstallID=%+v", d.InstallID)
	fmt.Fprintf(&builder, ", ActionStatus=%+v", d.ActionStatus)
	fmt.Fprintf(&builder, ", NotifySequence=%+v", d.NotifySequence)
	fmt.Fprintf(&builder, ", TradingDay=%+v", d.TradingDay)
	fmt.Fprintf(&builder, ", SettlementID=%+v", d.SettlementID)
	fmt.Fprintf(&builder, ", SequenceNo=%+v", d.SequenceNo)
	fmt.Fprintf(&builder, ", IPAddress=%+v", d.IPAddress)
	fmt.Fprintf(&builder, ", MacAddress=%+v", d.MacAddress)

	builder.WriteByte('}')

	return builder.String()
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

func (d CThostFtdcQryExchangeCombActionField) Type() string {
	return "CThostFtdcQryExchangeCombActionField"
}

func (d CThostFtdcQryExchangeCombActionField) String() string {
	var builder strings.Builder
	builder.Grow(139)

	builder.WriteString("CThostFtdcQryExchangeCombActionField{")
	fmt.Fprintf(&builder, "ParticipantID=%+v", d.ParticipantID)
	fmt.Fprintf(&builder, ", ClientID=%+v", d.ClientID)
	fmt.Fprintf(&builder, ", ExchangeInstID=%+v", d.ExchangeInstID)
	fmt.Fprintf(&builder, ", ExchangeID=%+v", d.ExchangeID)
	fmt.Fprintf(&builder, ", TraderID=%+v", d.TraderID)

	builder.WriteByte('}')

	return builder.String()
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

func (d CThostFtdcProductExchRateField) Type() string { return "CThostFtdcProductExchRateField" }

func (d CThostFtdcProductExchRateField) String() string {
	var builder strings.Builder
	builder.Grow(96)

	builder.WriteString("CThostFtdcProductExchRateField{")
	fmt.Fprintf(&builder, "ProductID=%+v", d.ProductID)
	fmt.Fprintf(&builder, ", QuoteCurrencyID=%+v", d.QuoteCurrencyID)
	fmt.Fprintf(&builder, ", ExchangeRate=%+v", d.ExchangeRate)

	builder.WriteByte('}')

	return builder.String()
}

// 产品报价汇率查询
type CThostFtdcQryProductExchRateField struct {
	// 产品代码
	ProductID types.TThostFtdcInstrumentIDType
}

func (d CThostFtdcQryProductExchRateField) Type() string { return "CThostFtdcQryProductExchRateField" }

func (d CThostFtdcQryProductExchRateField) String() string {
	var builder strings.Builder
	builder.Grow(52)

	builder.WriteString("CThostFtdcQryProductExchRateField{")
	fmt.Fprintf(&builder, "ProductID=%+v", d.ProductID)

	builder.WriteByte('}')

	return builder.String()
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

func (d CThostFtdcQryForQuoteParamField) Type() string { return "CThostFtdcQryForQuoteParamField" }

func (d CThostFtdcQryForQuoteParamField) String() string {
	var builder strings.Builder
	builder.Grow(91)

	builder.WriteString("CThostFtdcQryForQuoteParamField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", InstrumentID=%+v", d.InstrumentID)
	fmt.Fprintf(&builder, ", ExchangeID=%+v", d.ExchangeID)

	builder.WriteByte('}')

	return builder.String()
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

// 查询SPMM商品群明细
type CThostFtdcQrySPMMInvestorCommodityGroupMarginField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID types.TThostFtdcInvestorIDType
	// 商品群代码
	CommodityGroupID types.TThostFtdcSPMMProductIDType
}

func (d CThostFtdcQrySPMMInvestorCommodityGroupMarginField) Type() string {
	return "CThostFtdcQrySPMMInvestorCommodityGroupMarginField"
}

func (d CThostFtdcQrySPMMInvestorCommodityGroupMarginField) String() string {
	var builder strings.Builder
	builder.Grow(114)

	builder.WriteString("CThostFtdcQrySPMMInvestorCommodityGroupMarginField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", InvestorID=%+v", d.InvestorID)
	fmt.Fprintf(&builder, ", CommodityGroupID=%+v", d.CommodityGroupID)

	builder.WriteByte('}')

	return builder.String()
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

func (d CThostFtdcQryRULEInvestorProdMarginField) Type() string {
	return "CThostFtdcQryRULEInvestorProdMarginField"
}

func (d CThostFtdcQryRULEInvestorProdMarginField) String() string {
	var builder strings.Builder
	builder.Grow(148)

	builder.WriteString("CThostFtdcQryRULEInvestorProdMarginField{")
	fmt.Fprintf(&builder, "ExchangeID=%+v", d.ExchangeID)
	fmt.Fprintf(&builder, ", BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", InvestorID=%+v", d.InvestorID)
	fmt.Fprintf(&builder, ", ProdFamilyCode=%+v", d.ProdFamilyCode)
	fmt.Fprintf(&builder, ", CommodityGroupID=%+v", d.CommodityGroupID)

	builder.WriteByte('}')

	return builder.String()
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

func (d CThostFtdcForQuoteParamField) Type() string { return "CThostFtdcForQuoteParamField" }

func (d CThostFtdcForQuoteParamField) String() string {
	var builder strings.Builder
	builder.Grow(130)

	builder.WriteString("CThostFtdcForQuoteParamField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", InstrumentID=%+v", d.InstrumentID)
	fmt.Fprintf(&builder, ", ExchangeID=%+v", d.ExchangeID)
	fmt.Fprintf(&builder, ", LastPrice=%+v", d.LastPrice)
	fmt.Fprintf(&builder, ", PriceInterval=%+v", d.PriceInterval)

	builder.WriteByte('}')

	return builder.String()
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

func (d CThostFtdcMMOptionInstrCommRateField) Type() string {
	return "CThostFtdcMMOptionInstrCommRateField"
}

func (d CThostFtdcMMOptionInstrCommRateField) String() string {
	var builder strings.Builder
	builder.Grow(349)

	builder.WriteString("CThostFtdcMMOptionInstrCommRateField{")
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

	builder.WriteByte('}')

	return builder.String()
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

func (d CThostFtdcQryMMOptionInstrCommRateField) Type() string {
	return "CThostFtdcQryMMOptionInstrCommRateField"
}

func (d CThostFtdcQryMMOptionInstrCommRateField) String() string {
	var builder strings.Builder
	builder.Grow(99)

	builder.WriteString("CThostFtdcQryMMOptionInstrCommRateField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", InvestorID=%+v", d.InvestorID)
	fmt.Fprintf(&builder, ", InstrumentID=%+v", d.InstrumentID)

	builder.WriteByte('}')

	return builder.String()
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

func (d CThostFtdcMMInstrumentCommissionRateField) Type() string {
	return "CThostFtdcMMInstrumentCommissionRateField"
}

func (d CThostFtdcMMInstrumentCommissionRateField) String() string {
	var builder strings.Builder
	builder.Grow(297)

	builder.WriteString("CThostFtdcMMInstrumentCommissionRateField{")
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

	builder.WriteByte('}')

	return builder.String()
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

func (d CThostFtdcQryMMInstrumentCommissionRateField) Type() string {
	return "CThostFtdcQryMMInstrumentCommissionRateField"
}

func (d CThostFtdcQryMMInstrumentCommissionRateField) String() string {
	var builder strings.Builder
	builder.Grow(104)

	builder.WriteString("CThostFtdcQryMMInstrumentCommissionRateField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", InvestorID=%+v", d.InvestorID)
	fmt.Fprintf(&builder, ", InstrumentID=%+v", d.InstrumentID)

	builder.WriteByte('}')

	return builder.String()
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

func (d CThostFtdcInstrumentOrderCommRateField) Type() string {
	return "CThostFtdcInstrumentOrderCommRateField"
}

func (d CThostFtdcInstrumentOrderCommRateField) String() string {
	var builder strings.Builder
	builder.Grow(200)

	builder.WriteString("CThostFtdcInstrumentOrderCommRateField{")
	fmt.Fprintf(&builder, "InstrumentID=%+v", d.InstrumentID)
	fmt.Fprintf(&builder, ", InvestorRange=%+v", d.InvestorRange)
	fmt.Fprintf(&builder, ", BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", InvestorID=%+v", d.InvestorID)
	fmt.Fprintf(&builder, ", HedgeFlag=%+v", d.HedgeFlag)
	fmt.Fprintf(&builder, ", OrderCommByVolume=%+v", d.OrderCommByVolume)
	fmt.Fprintf(&builder, ", OrderActionCommByVolume=%+v", d.OrderActionCommByVolume)

	builder.WriteByte('}')

	return builder.String()
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

func (d CThostFtdcQryInstrumentOrderCommRateField) Type() string {
	return "CThostFtdcQryInstrumentOrderCommRateField"
}

func (d CThostFtdcQryInstrumentOrderCommRateField) String() string {
	var builder strings.Builder
	builder.Grow(120)

	builder.WriteString("CThostFtdcQryInstrumentOrderCommRateField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", InvestorID=%+v", d.InvestorID)
	fmt.Fprintf(&builder, ", InstrumentID=%+v", d.InstrumentID)
	fmt.Fprintf(&builder, ", HedgeFlag=%+v", d.HedgeFlag)

	builder.WriteByte('}')

	return builder.String()
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

func (d CThostFtdcMarketDataField) Type() string { return "CThostFtdcMarketDataField" }

func (d CThostFtdcMarketDataField) String() string {
	var builder strings.Builder
	builder.Grow(519)

	builder.WriteString("CThostFtdcMarketDataField{")
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
	fmt.Fprintf(&builder, ", ActionDay=%+v", d.ActionDay)

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
	// 合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
	// 最后修改时间
	UpdateTime types.TThostFtdcTimeType
	// 最后修改毫秒
	UpdateMillisec types.TThostFtdcMillisecType
	// 业务日期
	ActionDay types.TThostFtdcDateType
}

func (d CThostFtdcMarketDataUpdateTimeField) Type() string {
	return "CThostFtdcMarketDataUpdateTimeField"
}

func (d CThostFtdcMarketDataUpdateTimeField) String() string {
	var builder strings.Builder
	builder.Grow(120)

	builder.WriteString("CThostFtdcMarketDataUpdateTimeField{")
	fmt.Fprintf(&builder, "InstrumentID=%+v", d.InstrumentID)
	fmt.Fprintf(&builder, ", UpdateTime=%+v", d.UpdateTime)
	fmt.Fprintf(&builder, ", UpdateMillisec=%+v", d.UpdateMillisec)
	fmt.Fprintf(&builder, ", ActionDay=%+v", d.ActionDay)

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
	// 合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
}

func (d CThostFtdcSpecificInstrumentField) Type() string { return "CThostFtdcSpecificInstrumentField" }

func (d CThostFtdcSpecificInstrumentField) String() string {
	var builder strings.Builder
	builder.Grow(55)

	builder.WriteString("CThostFtdcSpecificInstrumentField{")
	fmt.Fprintf(&builder, "InstrumentID=%+v", d.InstrumentID)

	builder.WriteByte('}')

	return builder.String()
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

func (d CThostFtdcInstrumentStatusField) Type() string { return "CThostFtdcInstrumentStatusField" }

func (d CThostFtdcInstrumentStatusField) String() string {
	var builder strings.Builder
	builder.Grow(216)

	builder.WriteString("CThostFtdcInstrumentStatusField{")
	fmt.Fprintf(&builder, "ExchangeID=%+v", d.ExchangeID)
	fmt.Fprintf(&builder, ", ExchangeInstID=%+v", d.ExchangeInstID)
	fmt.Fprintf(&builder, ", SettlementGroupID=%+v", d.SettlementGroupID)
	fmt.Fprintf(&builder, ", InstrumentID=%+v", d.InstrumentID)
	fmt.Fprintf(&builder, ", InstrumentStatus=%+v", d.InstrumentStatus)
	fmt.Fprintf(&builder, ", TradingSegmentSN=%+v", d.TradingSegmentSN)
	fmt.Fprintf(&builder, ", EnterTime=%+v", d.EnterTime)
	fmt.Fprintf(&builder, ", EnterReason=%+v", d.EnterReason)

	builder.WriteByte('}')

	return builder.String()
}

// 查询合约状态
type CThostFtdcQryInstrumentStatusField struct {
	// 交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
	// 合约在交易所的代码
	ExchangeInstID types.TThostFtdcExchangeInstIDType
}

func (d CThostFtdcQryInstrumentStatusField) Type() string {
	return "CThostFtdcQryInstrumentStatusField"
}

func (d CThostFtdcQryInstrumentStatusField) String() string {
	var builder strings.Builder
	builder.Grow(78)

	builder.WriteString("CThostFtdcQryInstrumentStatusField{")
	fmt.Fprintf(&builder, "ExchangeID=%+v", d.ExchangeID)
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
	// 合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
}

func (d CThostFtdcQryInvestorPositionDetailField) Type() string {
	return "CThostFtdcQryInvestorPositionDetailField"
}

func (d CThostFtdcQryInvestorPositionDetailField) String() string {
	var builder strings.Builder
	builder.Grow(100)

	builder.WriteString("CThostFtdcQryInvestorPositionDetailField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", InvestorID=%+v", d.InvestorID)
	fmt.Fprintf(&builder, ", InstrumentID=%+v", d.InstrumentID)

	builder.WriteByte('}')

	return builder.String()
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

func (d CThostFtdcInvestorPositionDetailField) Type() string {
	return "CThostFtdcInvestorPositionDetailField"
}

func (d CThostFtdcInvestorPositionDetailField) String() string {
	var builder strings.Builder
	builder.Grow(615)

	builder.WriteString("CThostFtdcInvestorPositionDetailField{")
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
}

func (d CThostFtdcMDTraderOfferField) Type() string { return "CThostFtdcMDTraderOfferField" }

func (d CThostFtdcMDTraderOfferField) String() string {
	var builder strings.Builder
	builder.Grow(453)

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
}

func (d CThostFtdcQrySettlementInfoConfirmField) Type() string {
	return "CThostFtdcQrySettlementInfoConfirmField"
}

func (d CThostFtdcQrySettlementInfoConfirmField) String() string {
	var builder strings.Builder
	builder.Grow(77)

	builder.WriteString("CThostFtdcQrySettlementInfoConfirmField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", InvestorID=%+v", d.InvestorID)

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
	// 组合合约代码
	CombInstrumentID types.TThostFtdcInstrumentIDType
	// 单腿编号
	LegID types.TThostFtdcLegIDType
	// 单腿合约代码
	LegInstrumentID types.TThostFtdcInstrumentIDType
}

func (d CThostFtdcQryCombinationLegField) Type() string { return "CThostFtdcQryCombinationLegField" }

func (d CThostFtdcQryCombinationLegField) String() string {
	var builder strings.Builder
	builder.Grow(98)

	builder.WriteString("CThostFtdcQryCombinationLegField{")
	fmt.Fprintf(&builder, "CombInstrumentID=%+v", d.CombInstrumentID)
	fmt.Fprintf(&builder, ", LegID=%+v", d.LegID)
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

func (d CThostFtdcCombinationLegField) Type() string { return "CThostFtdcCombinationLegField" }

func (d CThostFtdcCombinationLegField) String() string {
	var builder strings.Builder
	builder.Grow(155)

	builder.WriteString("CThostFtdcCombinationLegField{")
	fmt.Fprintf(&builder, "CombInstrumentID=%+v", d.CombInstrumentID)
	fmt.Fprintf(&builder, ", LegID=%+v", d.LegID)
	fmt.Fprintf(&builder, ", LegInstrumentID=%+v", d.LegInstrumentID)
	fmt.Fprintf(&builder, ", Direction=%+v", d.Direction)
	fmt.Fprintf(&builder, ", LegMultiple=%+v", d.LegMultiple)
	fmt.Fprintf(&builder, ", ImplyLevel=%+v", d.ImplyLevel)

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
	// 合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
}

func (d CThostFtdcBrokerUserEventField) Type() string { return "CThostFtdcBrokerUserEventField" }

func (d CThostFtdcBrokerUserEventField) String() string {
	var builder strings.Builder
	builder.Grow(215)

	builder.WriteString("CThostFtdcBrokerUserEventField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", UserID=%+v", d.UserID)
	fmt.Fprintf(&builder, ", UserEventType=%+v", d.UserEventType)
	fmt.Fprintf(&builder, ", EventSequenceNo=%+v", d.EventSequenceNo)
	fmt.Fprintf(&builder, ", EventDate=%+v", d.EventDate)
	fmt.Fprintf(&builder, ", EventTime=%+v", d.EventTime)
	fmt.Fprintf(&builder, ", UserEventInfo=%+v", d.UserEventInfo)
	fmt.Fprintf(&builder, ", InvestorID=%+v", d.InvestorID)
	fmt.Fprintf(&builder, ", InstrumentID=%+v", d.InstrumentID)

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
}

func (d CThostFtdcContractBankField) Type() string { return "CThostFtdcContractBankField" }

func (d CThostFtdcContractBankField) String() string {
	var builder strings.Builder
	builder.Grow(99)

	builder.WriteString("CThostFtdcContractBankField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", BankID=%+v", d.BankID)
	fmt.Fprintf(&builder, ", BankBrchID=%+v", d.BankBrchID)
	fmt.Fprintf(&builder, ", BankName=%+v", d.BankName)

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

func (d CThostFtdcInvestorPositionCombineDetailField) Type() string {
	return "CThostFtdcInvestorPositionCombineDetailField"
}

func (d CThostFtdcInvestorPositionCombineDetailField) String() string {
	var builder strings.Builder
	builder.Grow(452)

	builder.WriteString("CThostFtdcInvestorPositionCombineDetailField{")
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
	fmt.Fprintf(&builder, ", CombInstrumentID=%+v", d.CombInstrumentID)
	fmt.Fprintf(&builder, ", TradeGroupID=%+v", d.TradeGroupID)

	builder.WriteByte('}')

	return builder.String()
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

func (d CThostFtdcParkedOrderField) Type() string { return "CThostFtdcParkedOrderField" }

func (d CThostFtdcParkedOrderField) String() string {
	var builder strings.Builder
	builder.Grow(756)

	builder.WriteString("CThostFtdcParkedOrderField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", InvestorID=%+v", d.InvestorID)
	fmt.Fprintf(&builder, ", InstrumentID=%+v", d.InstrumentID)
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
	fmt.Fprintf(&builder, ", IPAddress=%+v", d.IPAddress)
	fmt.Fprintf(&builder, ", MacAddress=%+v", d.MacAddress)

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

func (d CThostFtdcParkedOrderActionField) Type() string { return "CThostFtdcParkedOrderActionField" }

func (d CThostFtdcParkedOrderActionField) String() string {
	var builder strings.Builder
	builder.Grow(466)

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
	fmt.Fprintf(&builder, ", InstrumentID=%+v", d.InstrumentID)
	fmt.Fprintf(&builder, ", ParkedOrderActionID=%+v", d.ParkedOrderActionID)
	fmt.Fprintf(&builder, ", UserType=%+v", d.UserType)
	fmt.Fprintf(&builder, ", Status=%+v", d.Status)
	fmt.Fprintf(&builder, ", ErrorID=%+v", d.ErrorID)
	fmt.Fprintf(&builder, ", ErrorMsg=%+v", d.ErrorMsg)
	fmt.Fprintf(&builder, ", InvestUnitID=%+v", d.InvestUnitID)
	fmt.Fprintf(&builder, ", IPAddress=%+v", d.IPAddress)
	fmt.Fprintf(&builder, ", MacAddress=%+v", d.MacAddress)

	builder.WriteByte('}')

	return builder.String()
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

func (d CThostFtdcQryParkedOrderField) Type() string { return "CThostFtdcQryParkedOrderField" }

func (d CThostFtdcQryParkedOrderField) String() string {
	var builder strings.Builder
	builder.Grow(109)

	builder.WriteString("CThostFtdcQryParkedOrderField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", InvestorID=%+v", d.InvestorID)
	fmt.Fprintf(&builder, ", InstrumentID=%+v", d.InstrumentID)
	fmt.Fprintf(&builder, ", ExchangeID=%+v", d.ExchangeID)

	builder.WriteByte('}')

	return builder.String()
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

func (d CThostFtdcQryParkedOrderActionField) Type() string {
	return "CThostFtdcQryParkedOrderActionField"
}

func (d CThostFtdcQryParkedOrderActionField) String() string {
	var builder strings.Builder
	builder.Grow(115)

	builder.WriteString("CThostFtdcQryParkedOrderActionField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", InvestorID=%+v", d.InvestorID)
	fmt.Fprintf(&builder, ", InstrumentID=%+v", d.InstrumentID)
	fmt.Fprintf(&builder, ", ExchangeID=%+v", d.ExchangeID)

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
}

func (d CThostFtdcRemoveParkedOrderField) Type() string { return "CThostFtdcRemoveParkedOrderField" }

func (d CThostFtdcRemoveParkedOrderField) String() string {
	var builder strings.Builder
	builder.Grow(93)

	builder.WriteString("CThostFtdcRemoveParkedOrderField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", InvestorID=%+v", d.InvestorID)
	fmt.Fprintf(&builder, ", ParkedOrderID=%+v", d.ParkedOrderID)

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
}

func (d CThostFtdcRemoveParkedOrderActionField) Type() string {
	return "CThostFtdcRemoveParkedOrderActionField"
}

func (d CThostFtdcRemoveParkedOrderActionField) String() string {
	var builder strings.Builder
	builder.Grow(105)

	builder.WriteString("CThostFtdcRemoveParkedOrderActionField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", InvestorID=%+v", d.InvestorID)
	fmt.Fprintf(&builder, ", ParkedOrderActionID=%+v", d.ParkedOrderActionID)

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
	// 组合持仓合约编码
	CombInstrumentID types.TThostFtdcInstrumentIDType
}

func (d CThostFtdcQryInvestorPositionCombineDetailField) Type() string {
	return "CThostFtdcQryInvestorPositionCombineDetailField"
}

func (d CThostFtdcQryInvestorPositionCombineDetailField) String() string {
	var builder strings.Builder
	builder.Grow(111)

	builder.WriteString("CThostFtdcQryInvestorPositionCombineDetailField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", InvestorID=%+v", d.InvestorID)
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
	// IP地址
	IPAddress types.TThostFtdcIPAddressType
	// IP地址掩码
	IPMask types.TThostFtdcIPAddressType
	// Mac地址
	MacAddress types.TThostFtdcMacAddressType
}

func (d CThostFtdcUserIPField) Type() string { return "CThostFtdcUserIPField" }

func (d CThostFtdcUserIPField) String() string {
	var builder strings.Builder
	builder.Grow(110)

	builder.WriteString("CThostFtdcUserIPField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", UserID=%+v", d.UserID)
	fmt.Fprintf(&builder, ", IPAddress=%+v", d.IPAddress)
	fmt.Fprintf(&builder, ", IPMask=%+v", d.IPMask)
	fmt.Fprintf(&builder, ", MacAddress=%+v", d.MacAddress)

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
}

func (d CThostFtdcTradingNoticeInfoField) Type() string { return "CThostFtdcTradingNoticeInfoField" }

func (d CThostFtdcTradingNoticeInfoField) String() string {
	var builder strings.Builder
	builder.Grow(154)

	builder.WriteString("CThostFtdcTradingNoticeInfoField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", InvestorID=%+v", d.InvestorID)
	fmt.Fprintf(&builder, ", SendTime=%+v", d.SendTime)
	fmt.Fprintf(&builder, ", FieldContent=%+v", d.FieldContent)
	fmt.Fprintf(&builder, ", SequenceSeries=%+v", d.SequenceSeries)
	fmt.Fprintf(&builder, ", SequenceNo=%+v", d.SequenceNo)

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
}

func (d CThostFtdcTradingNoticeField) Type() string { return "CThostFtdcTradingNoticeField" }

func (d CThostFtdcTradingNoticeField) String() string {
	var builder strings.Builder
	builder.Grow(189)

	builder.WriteString("CThostFtdcTradingNoticeField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", InvestorRange=%+v", d.InvestorRange)
	fmt.Fprintf(&builder, ", InvestorID=%+v", d.InvestorID)
	fmt.Fprintf(&builder, ", SequenceSeries=%+v", d.SequenceSeries)
	fmt.Fprintf(&builder, ", UserID=%+v", d.UserID)
	fmt.Fprintf(&builder, ", SendTime=%+v", d.SendTime)
	fmt.Fprintf(&builder, ", SequenceNo=%+v", d.SequenceNo)
	fmt.Fprintf(&builder, ", FieldContent=%+v", d.FieldContent)

	builder.WriteByte('}')

	return builder.String()
}

// 查询交易事件通知
type CThostFtdcQryTradingNoticeField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID types.TThostFtdcInvestorIDType
}

func (d CThostFtdcQryTradingNoticeField) Type() string { return "CThostFtdcQryTradingNoticeField" }

func (d CThostFtdcQryTradingNoticeField) String() string {
	var builder strings.Builder
	builder.Grow(69)

	builder.WriteString("CThostFtdcQryTradingNoticeField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", InvestorID=%+v", d.InvestorID)

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

func (d CThostFtdcErrOrderField) Type() string { return "CThostFtdcErrOrderField" }

func (d CThostFtdcErrOrderField) String() string {
	var builder strings.Builder
	builder.Grow(696)

	builder.WriteString("CThostFtdcErrOrderField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", InvestorID=%+v", d.InvestorID)
	fmt.Fprintf(&builder, ", InstrumentID=%+v", d.InstrumentID)
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
	fmt.Fprintf(&builder, ", IPAddress=%+v", d.IPAddress)
	fmt.Fprintf(&builder, ", MacAddress=%+v", d.MacAddress)

	builder.WriteByte('}')

	return builder.String()
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

func (d CThostFtdcErrorConditionalOrderField) Type() string {
	return "CThostFtdcErrorConditionalOrderField"
}

func (d CThostFtdcErrorConditionalOrderField) String() string {
	var builder strings.Builder
	builder.Grow(1423)

	builder.WriteString("CThostFtdcErrorConditionalOrderField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", InvestorID=%+v", d.InvestorID)
	fmt.Fprintf(&builder, ", InstrumentID=%+v", d.InstrumentID)
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
	fmt.Fprintf(&builder, ", ExchangeInstID=%+v", d.ExchangeInstID)
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
	fmt.Fprintf(&builder, ", IPAddress=%+v", d.IPAddress)
	fmt.Fprintf(&builder, ", MacAddress=%+v", d.MacAddress)

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

func (d CThostFtdcErrOrderActionField) Type() string { return "CThostFtdcErrOrderActionField" }

func (d CThostFtdcErrOrderActionField) String() string {
	var builder strings.Builder
	builder.Grow(649)

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
	fmt.Fprintf(&builder, ", InstrumentID=%+v", d.InstrumentID)
	fmt.Fprintf(&builder, ", BranchID=%+v", d.BranchID)
	fmt.Fprintf(&builder, ", InvestUnitID=%+v", d.InvestUnitID)
	fmt.Fprintf(&builder, ", IPAddress=%+v", d.IPAddress)
	fmt.Fprintf(&builder, ", MacAddress=%+v", d.MacAddress)
	fmt.Fprintf(&builder, ", ErrorID=%+v", d.ErrorID)
	fmt.Fprintf(&builder, ", ErrorMsg=%+v", d.ErrorMsg)

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

func (d CThostFtdcQueryMaxOrderVolumeWithPriceField) Type() string {
	return "CThostFtdcQueryMaxOrderVolumeWithPriceField"
}

func (d CThostFtdcQueryMaxOrderVolumeWithPriceField) String() string {
	var builder strings.Builder
	builder.Grow(195)

	builder.WriteString("CThostFtdcQueryMaxOrderVolumeWithPriceField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", InvestorID=%+v", d.InvestorID)
	fmt.Fprintf(&builder, ", InstrumentID=%+v", d.InstrumentID)
	fmt.Fprintf(&builder, ", Direction=%+v", d.Direction)
	fmt.Fprintf(&builder, ", OffsetFlag=%+v", d.OffsetFlag)
	fmt.Fprintf(&builder, ", HedgeFlag=%+v", d.HedgeFlag)
	fmt.Fprintf(&builder, ", MaxVolume=%+v", d.MaxVolume)
	fmt.Fprintf(&builder, ", Price=%+v", d.Price)

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
}

func (d CThostFtdcQryBrokerTradingParamsField) Type() string {
	return "CThostFtdcQryBrokerTradingParamsField"
}

func (d CThostFtdcQryBrokerTradingParamsField) String() string {
	var builder strings.Builder
	builder.Grow(95)

	builder.WriteString("CThostFtdcQryBrokerTradingParamsField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", InvestorID=%+v", d.InvestorID)
	fmt.Fprintf(&builder, ", CurrencyID=%+v", d.CurrencyID)

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
}

func (d CThostFtdcBrokerTradingParamsField) Type() string {
	return "CThostFtdcBrokerTradingParamsField"
}

func (d CThostFtdcBrokerTradingParamsField) String() string {
	var builder strings.Builder
	builder.Grow(201)

	builder.WriteString("CThostFtdcBrokerTradingParamsField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", InvestorID=%+v", d.InvestorID)
	fmt.Fprintf(&builder, ", MarginPriceType=%+v", d.MarginPriceType)
	fmt.Fprintf(&builder, ", Algorithm=%+v", d.Algorithm)
	fmt.Fprintf(&builder, ", AvailIncludeCloseProfit=%+v", d.AvailIncludeCloseProfit)
	fmt.Fprintf(&builder, ", CurrencyID=%+v", d.CurrencyID)
	fmt.Fprintf(&builder, ", OptionRoyaltyPriceType=%+v", d.OptionRoyaltyPriceType)

	builder.WriteByte('}')

	return builder.String()
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

func (d CThostFtdcQryBrokerTradingAlgosField) Type() string {
	return "CThostFtdcQryBrokerTradingAlgosField"
}

func (d CThostFtdcQryBrokerTradingAlgosField) String() string {
	var builder strings.Builder
	builder.Grow(96)

	builder.WriteString("CThostFtdcQryBrokerTradingAlgosField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", ExchangeID=%+v", d.ExchangeID)
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
	// 合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
	// 持仓处理算法编号
	HandlePositionAlgoID types.TThostFtdcHandlePositionAlgoIDType
	// 寻找保证金率算法编号
	FindMarginRateAlgoID types.TThostFtdcFindMarginRateAlgoIDType
	// 资金处理算法编号
	HandleTradingAccountAlgoID types.TThostFtdcHandleTradingAccountAlgoIDType
}

func (d CThostFtdcBrokerTradingAlgosField) Type() string { return "CThostFtdcBrokerTradingAlgosField" }

func (d CThostFtdcBrokerTradingAlgosField) String() string {
	var builder strings.Builder
	builder.Grow(189)

	builder.WriteString("CThostFtdcBrokerTradingAlgosField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", ExchangeID=%+v", d.ExchangeID)
	fmt.Fprintf(&builder, ", InstrumentID=%+v", d.InstrumentID)
	fmt.Fprintf(&builder, ", HandlePositionAlgoID=%+v", d.HandlePositionAlgoID)
	fmt.Fprintf(&builder, ", FindMarginRateAlgoID=%+v", d.FindMarginRateAlgoID)
	fmt.Fprintf(&builder, ", HandleTradingAccountAlgoID=%+v", d.HandleTradingAccountAlgoID)

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
	// 合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
	// 买卖方向
	Direction types.TThostFtdcDirectionType
	// 投机套保标志
	HedgeFlag types.TThostFtdcHedgeFlagType
	// 数量
	Volume types.TThostFtdcVolumeType
}

func (d CThostFtdcEWarrantOffsetField) Type() string { return "CThostFtdcEWarrantOffsetField" }

func (d CThostFtdcEWarrantOffsetField) String() string {
	var builder strings.Builder
	builder.Grow(183)

	builder.WriteString("CThostFtdcEWarrantOffsetField{")
	fmt.Fprintf(&builder, "TradingDay=%+v", d.TradingDay)
	fmt.Fprintf(&builder, ", BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", InvestorID=%+v", d.InvestorID)
	fmt.Fprintf(&builder, ", ExchangeID=%+v", d.ExchangeID)
	fmt.Fprintf(&builder, ", InstrumentID=%+v", d.InstrumentID)
	fmt.Fprintf(&builder, ", Direction=%+v", d.Direction)
	fmt.Fprintf(&builder, ", HedgeFlag=%+v", d.HedgeFlag)
	fmt.Fprintf(&builder, ", Volume=%+v", d.Volume)

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
	// 合约代码
	InstrumentID types.TThostFtdcInstrumentIDType
}

func (d CThostFtdcQryEWarrantOffsetField) Type() string { return "CThostFtdcQryEWarrantOffsetField" }

func (d CThostFtdcQryEWarrantOffsetField) String() string {
	var builder strings.Builder
	builder.Grow(112)

	builder.WriteString("CThostFtdcQryEWarrantOffsetField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", InvestorID=%+v", d.InvestorID)
	fmt.Fprintf(&builder, ", ExchangeID=%+v", d.ExchangeID)
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
	// 品种跨品种标示
	ProductGroupID types.TThostFtdcInstrumentIDType
	// 投机套保标志
	HedgeFlag types.TThostFtdcHedgeFlagType
}

func (d CThostFtdcQryInvestorProductGroupMarginField) Type() string {
	return "CThostFtdcQryInvestorProductGroupMarginField"
}

func (d CThostFtdcQryInvestorProductGroupMarginField) String() string {
	var builder strings.Builder
	builder.Grow(125)

	builder.WriteString("CThostFtdcQryInvestorProductGroupMarginField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", InvestorID=%+v", d.InvestorID)
	fmt.Fprintf(&builder, ", ProductGroupID=%+v", d.ProductGroupID)
	fmt.Fprintf(&builder, ", HedgeFlag=%+v", d.HedgeFlag)

	builder.WriteByte('}')

	return builder.String()
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

func (d CThostFtdcInvestorProductGroupMarginField) Type() string {
	return "CThostFtdcInvestorProductGroupMarginField"
}

func (d CThostFtdcInvestorProductGroupMarginField) String() string {
	var builder strings.Builder
	builder.Grow(663)

	builder.WriteString("CThostFtdcInvestorProductGroupMarginField{")
	fmt.Fprintf(&builder, "ProductGroupID=%+v", d.ProductGroupID)
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

	builder.WriteByte('}')

	return builder.String()
}

// 查询监控中心用户令牌
type CThostFtdcQueryCFMMCTradingAccountTokenField struct {
	// 经纪公司代码
	BrokerID types.TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID types.TThostFtdcInvestorIDType
}

func (d CThostFtdcQueryCFMMCTradingAccountTokenField) Type() string {
	return "CThostFtdcQueryCFMMCTradingAccountTokenField"
}

func (d CThostFtdcQueryCFMMCTradingAccountTokenField) String() string {
	var builder strings.Builder
	builder.Grow(82)

	builder.WriteString("CThostFtdcQueryCFMMCTradingAccountTokenField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", InvestorID=%+v", d.InvestorID)

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
	// 产品代码
	ProductID types.TThostFtdcInstrumentIDType
	// 交易所代码
	ExchangeID types.TThostFtdcExchangeIDType
}

func (d CThostFtdcQryProductGroupField) Type() string { return "CThostFtdcQryProductGroupField" }

func (d CThostFtdcQryProductGroupField) String() string {
	var builder strings.Builder
	builder.Grow(69)

	builder.WriteString("CThostFtdcQryProductGroupField{")
	fmt.Fprintf(&builder, "ProductID=%+v", d.ProductID)
	fmt.Fprintf(&builder, ", ExchangeID=%+v", d.ExchangeID)

	builder.WriteByte('}')

	return builder.String()
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

func (d CThostFtdcProductGroupField) Type() string { return "CThostFtdcProductGroupField" }

func (d CThostFtdcProductGroupField) String() string {
	var builder strings.Builder
	builder.Grow(90)

	builder.WriteString("CThostFtdcProductGroupField{")
	fmt.Fprintf(&builder, "ProductID=%+v", d.ProductID)
	fmt.Fprintf(&builder, ", ExchangeID=%+v", d.ExchangeID)
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

func (d CThostFtdcReqOpenAccountField) Type() string { return "CThostFtdcReqOpenAccountField" }

func (d CThostFtdcReqOpenAccountField) String() string {
	var builder strings.Builder
	builder.Grow(904)

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
}

func (d CThostFtdcReqCancelAccountField) Type() string { return "CThostFtdcReqCancelAccountField" }

func (d CThostFtdcReqCancelAccountField) String() string {
	var builder strings.Builder
	builder.Grow(906)

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
}

func (d CThostFtdcReqChangeAccountField) Type() string { return "CThostFtdcReqChangeAccountField" }

func (d CThostFtdcReqChangeAccountField) String() string {
	var builder strings.Builder
	builder.Grow(833)

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
}

func (d CThostFtdcReqTransferField) Type() string { return "CThostFtdcReqTransferField" }

func (d CThostFtdcReqTransferField) String() string {
	var builder strings.Builder
	builder.Grow(894)

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
}

func (d CThostFtdcRspTransferField) Type() string { return "CThostFtdcRspTransferField" }

func (d CThostFtdcRspTransferField) String() string {
	var builder strings.Builder
	builder.Grow(929)

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
}

func (d CThostFtdcReqRepealField) Type() string { return "CThostFtdcReqRepealField" }

func (d CThostFtdcReqRepealField) String() string {
	var builder strings.Builder
	builder.Grow(1074)

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
}

func (d CThostFtdcRspRepealField) Type() string { return "CThostFtdcRspRepealField" }

func (d CThostFtdcRspRepealField) String() string {
	var builder strings.Builder
	builder.Grow(1109)

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
}

func (d CThostFtdcReqQueryAccountField) Type() string { return "CThostFtdcReqQueryAccountField" }

func (d CThostFtdcReqQueryAccountField) String() string {
	var builder strings.Builder
	builder.Grow(753)

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
}

func (d CThostFtdcRspQueryAccountField) Type() string { return "CThostFtdcRspQueryAccountField" }

func (d CThostFtdcRspQueryAccountField) String() string {
	var builder strings.Builder
	builder.Grow(801)

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
}

func (d CThostFtdcReqQueryTradeResultBySerialField) Type() string {
	return "CThostFtdcReqQueryTradeResultBySerialField"
}

func (d CThostFtdcReqQueryTradeResultBySerialField) String() string {
	var builder strings.Builder
	builder.Grow(575)

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
}

func (d CThostFtdcVerifyCustInfoField) Type() string { return "CThostFtdcVerifyCustInfoField" }

func (d CThostFtdcVerifyCustInfoField) String() string {
	var builder strings.Builder
	builder.Grow(115)

	builder.WriteString("CThostFtdcVerifyCustInfoField{")
	fmt.Fprintf(&builder, "CustomerName=%+v", d.CustomerName)
	fmt.Fprintf(&builder, ", IdCardType=%+v", d.IdCardType)
	fmt.Fprintf(&builder, ", IdentifiedCardNo=%+v", d.IdentifiedCardNo)
	fmt.Fprintf(&builder, ", CustType=%+v", d.CustType)

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
}

func (d CThostFtdcVerifyFuturePasswordAndCustInfoField) Type() string {
	return "CThostFtdcVerifyFuturePasswordAndCustInfoField"
}

func (d CThostFtdcVerifyFuturePasswordAndCustInfoField) String() string {
	var builder strings.Builder
	builder.Grow(189)

	builder.WriteString("CThostFtdcVerifyFuturePasswordAndCustInfoField{")
	fmt.Fprintf(&builder, "CustomerName=%+v", d.CustomerName)
	fmt.Fprintf(&builder, ", IdCardType=%+v", d.IdCardType)
	fmt.Fprintf(&builder, ", IdentifiedCardNo=%+v", d.IdentifiedCardNo)
	fmt.Fprintf(&builder, ", CustType=%+v", d.CustType)
	fmt.Fprintf(&builder, ", AccountID=%+v", d.AccountID)
	fmt.Fprintf(&builder, ", Password=%+v", d.Password)
	fmt.Fprintf(&builder, ", CurrencyID=%+v", d.CurrencyID)

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
}

func (d CThostFtdcNotifyQueryAccountField) Type() string { return "CThostFtdcNotifyQueryAccountField" }

func (d CThostFtdcNotifyQueryAccountField) String() string {
	var builder strings.Builder
	builder.Grow(839)

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
}

func (d CThostFtdcAccountregisterField) Type() string { return "CThostFtdcAccountregisterField" }

func (d CThostFtdcAccountregisterField) String() string {
	var builder strings.Builder
	builder.Grow(365)

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
}

func (d CThostFtdcOpenAccountField) Type() string { return "CThostFtdcOpenAccountField" }

func (d CThostFtdcOpenAccountField) String() string {
	var builder strings.Builder
	builder.Grow(936)

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
}

func (d CThostFtdcCancelAccountField) Type() string { return "CThostFtdcCancelAccountField" }

func (d CThostFtdcCancelAccountField) String() string {
	var builder strings.Builder
	builder.Grow(938)

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
}

func (d CThostFtdcChangeAccountField) Type() string { return "CThostFtdcChangeAccountField" }

func (d CThostFtdcChangeAccountField) String() string {
	var builder strings.Builder
	builder.Grow(865)

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
	// IP地址
	IPAddress types.TThostFtdcIPAddressType
}

func (d CThostFtdcLoginForbiddenUserField) Type() string { return "CThostFtdcLoginForbiddenUserField" }

func (d CThostFtdcLoginForbiddenUserField) String() string {
	var builder strings.Builder
	builder.Grow(86)

	builder.WriteString("CThostFtdcLoginForbiddenUserField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", UserID=%+v", d.UserID)
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

// UDP组播组信息
type CThostFtdcMulticastGroupInfoField struct {
	// 组播组IP地址
	GroupIP types.TThostFtdcIPAddressType
	// 组播组IP端口
	GroupPort types.TThostFtdcIPPortType
	// 源地址
	SourceIP types.TThostFtdcIPAddressType
}

func (d CThostFtdcMulticastGroupInfoField) Type() string { return "CThostFtdcMulticastGroupInfoField" }

func (d CThostFtdcMulticastGroupInfoField) String() string {
	var builder strings.Builder
	builder.Grow(87)

	builder.WriteString("CThostFtdcMulticastGroupInfoField{")
	fmt.Fprintf(&builder, "GroupIP=%+v", d.GroupIP)
	fmt.Fprintf(&builder, ", GroupPort=%+v", d.GroupPort)
	fmt.Fprintf(&builder, ", SourceIP=%+v", d.SourceIP)

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

func (d CThostFtdcMBLMarketDataField) Type() string { return "CThostFtdcMBLMarketDataField" }

func (d CThostFtdcMBLMarketDataField) String() string {
	var builder strings.Builder
	builder.Grow(120)

	builder.WriteString("CThostFtdcMBLMarketDataField{")
	fmt.Fprintf(&builder, "InstrumentID=%+v", d.InstrumentID)
	fmt.Fprintf(&builder, ", ExchangeID=%+v", d.ExchangeID)
	fmt.Fprintf(&builder, ", Direction=%+v", d.Direction)
	fmt.Fprintf(&builder, ", Price=%+v", d.Price)
	fmt.Fprintf(&builder, ", Volume=%+v", d.Volume)

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
	// 用户公网IP
	ClientPublicIP types.TThostFtdcIPAddressType
	// 终端IP端口
	ClientIPPort types.TThostFtdcIPPortType
	// 登录成功时间
	ClientLoginTime types.TThostFtdcTimeType
	// App代码
	ClientAppID types.TThostFtdcClientAppIDType
}

func (d CThostFtdcUserSystemInfoField) Type() string { return "CThostFtdcUserSystemInfoField" }

func (d CThostFtdcUserSystemInfoField) String() string {
	var builder strings.Builder
	builder.Grow(210)

	builder.WriteString("CThostFtdcUserSystemInfoField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", UserID=%+v", d.UserID)
	fmt.Fprintf(&builder, ", ClientSystemInfoLen=%+v", d.ClientSystemInfoLen)
	fmt.Fprintf(&builder, ", ClientSystemInfo=%+v", d.ClientSystemInfo)
	fmt.Fprintf(&builder, ", ClientPublicIP=%+v", d.ClientPublicIP)
	fmt.Fprintf(&builder, ", ClientIPPort=%+v", d.ClientIPPort)
	fmt.Fprintf(&builder, ", ClientLoginTime=%+v", d.ClientLoginTime)
	fmt.Fprintf(&builder, ", ClientAppID=%+v", d.ClientAppID)

	builder.WriteByte('}')

	return builder.String()
}

// 指定的席位
type CThostFtdcSpecificTraderField struct {
	// 席位代码
	TraderID types.TThostFtdcTraderIDType
}

func (d CThostFtdcSpecificTraderField) Type() string { return "CThostFtdcSpecificTraderField" }

func (d CThostFtdcSpecificTraderField) String() string {
	var builder strings.Builder
	builder.Grow(47)

	builder.WriteString("CThostFtdcSpecificTraderField{")
	fmt.Fprintf(&builder, "TraderID=%+v", d.TraderID)

	builder.WriteByte('}')

	return builder.String()
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

func (d CThostFtdcFlowCtrlWarningField) Type() string { return "CThostFtdcFlowCtrlWarningField" }

func (d CThostFtdcFlowCtrlWarningField) String() string {
	var builder strings.Builder
	builder.Grow(102)

	builder.WriteString("CThostFtdcFlowCtrlWarningField{")
	fmt.Fprintf(&builder, "TraderID=%+v", d.TraderID)
	fmt.Fprintf(&builder, ", TriggerTime=%+v", d.TriggerTime)
	fmt.Fprintf(&builder, ", Tgid=%+v", d.Tgid)
	fmt.Fprintf(&builder, ", CurPkgCnt=%+v", d.CurPkgCnt)

	builder.WriteByte('}')

	return builder.String()
}

// 订阅资金变动应答
type CThostFtdcRequestIDEntityField struct {
	// 席位代码
	RequestID types.TThostFtdcRequestIDType
}

func (d CThostFtdcRequestIDEntityField) Type() string { return "CThostFtdcRequestIDEntityField" }

func (d CThostFtdcRequestIDEntityField) String() string {
	var builder strings.Builder
	builder.Grow(49)

	builder.WriteString("CThostFtdcRequestIDEntityField{")
	fmt.Fprintf(&builder, "RequestID=%+v", d.RequestID)

	builder.WriteByte('}')

	return builder.String()
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

func (d CThostFtdcQryInvestorInfoCntSettingField) Type() string {
	return "CThostFtdcQryInvestorInfoCntSettingField"
}

func (d CThostFtdcQryInvestorInfoCntSettingField) String() string {
	var builder strings.Builder
	builder.Grow(99)

	builder.WriteString("CThostFtdcQryInvestorInfoCntSettingField{")
	fmt.Fprintf(&builder, "ExchangeID=%+v", d.ExchangeID)
	fmt.Fprintf(&builder, ", InvestorID=%+v", d.InvestorID)
	fmt.Fprintf(&builder, ", ProductID=%+v", d.ProductID)

	builder.WriteByte('}')

	return builder.String()
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

func (d CThostFtdcUpdRiskForbiddenRightField) Type() string {
	return "CThostFtdcUpdRiskForbiddenRightField"
}

func (d CThostFtdcUpdRiskForbiddenRightField) String() string {
	var builder strings.Builder
	builder.Grow(153)

	builder.WriteString("CThostFtdcUpdRiskForbiddenRightField{")
	fmt.Fprintf(&builder, "BrokerID=%+v", d.BrokerID)
	fmt.Fprintf(&builder, ", InvestorID=%+v", d.InvestorID)
	fmt.Fprintf(&builder, ", InstrumentID=%+v", d.InstrumentID)
	fmt.Fprintf(&builder, ", UserID=%+v", d.UserID)
	fmt.Fprintf(&builder, ", EventType=%+v", d.EventType)
	fmt.Fprintf(&builder, ", TradingRight=%+v", d.TradingRight)

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
	// 业务单元
	BusinessUnit types.TThostFtdcBusinessUnitType
	// IP地址
	IPAddress types.TThostFtdcIPAddressType
	// Mac地址
	MacAddress types.TThostFtdcMacAddressType
}

func (d CThostFtdcInputOffsetSettingField) Type() string { return "CThostFtdcInputOffsetSettingField" }

func (d CThostFtdcInputOffsetSettingField) String() string {
	var builder strings.Builder
	builder.Grow(309)

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
	fmt.Fprintf(&builder, ", BusinessUnit=%+v", d.BusinessUnit)
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

func (d CThostFtdcOffsetSettingField) Type() string { return "CThostFtdcOffsetSettingField" }

func (d CThostFtdcOffsetSettingField) String() string {
	var builder strings.Builder
	builder.Grow(755)

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
	fmt.Fprintf(&builder, ", BusinessUnit=%+v", d.BusinessUnit)
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

// 查询IP接入控制
type CThostFtdcQryIPUserACLField struct {
	// 校验模式
	Mode types.TThostFtdcIPACLCheckModeType
	// IP地址
	IPAddress types.TThostFtdcIPAddressType
	// 用户代码
	UserID types.TThostFtdcUserIDType
}

func (d CThostFtdcQryIPUserACLField) Type() string { return "CThostFtdcQryIPUserACLField" }

func (d CThostFtdcQryIPUserACLField) String() string {
	var builder strings.Builder
	builder.Grow(76)

	builder.WriteString("CThostFtdcQryIPUserACLField{")
	fmt.Fprintf(&builder, "Mode=%+v", d.Mode)
	fmt.Fprintf(&builder, ", IPAddress=%+v", d.IPAddress)
	fmt.Fprintf(&builder, ", UserID=%+v", d.UserID)

	builder.WriteByte('}')

	return builder.String()
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

func (d CThostFtdcIPUserACLField) Type() string { return "CThostFtdcIPUserACLField" }

func (d CThostFtdcIPUserACLField) String() string {
	var builder strings.Builder
	builder.Grow(128)

	builder.WriteString("CThostFtdcIPUserACLField{")
	fmt.Fprintf(&builder, "Mode=%+v", d.Mode)
	fmt.Fprintf(&builder, ", IPAddress=%+v", d.IPAddress)
	fmt.Fprintf(&builder, ", UserID=%+v", d.UserID)
	fmt.Fprintf(&builder, ", Threshold=%+v", d.Threshold)
	fmt.Fprintf(&builder, ", Count=%+v", d.Count)
	fmt.Fprintf(&builder, ", IsForbidden=%+v", d.IsForbidden)

	builder.WriteByte('}')

	return builder.String()
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

func (d CThostFtdcUpdIPUserACLField) Type() string { return "CThostFtdcUpdIPUserACLField" }

func (d CThostFtdcUpdIPUserACLField) String() string {
	var builder strings.Builder
	builder.Grow(130)

	builder.WriteString("CThostFtdcUpdIPUserACLField{")
	fmt.Fprintf(&builder, "Mode=%+v", d.Mode)
	fmt.Fprintf(&builder, ", IPAddress=%+v", d.IPAddress)
	fmt.Fprintf(&builder, ", UserID=%+v", d.UserID)
	fmt.Fprintf(&builder, ", Threshold=%+v", d.Threshold)
	fmt.Fprintf(&builder, ", EventType=%+v", d.EventType)
	fmt.Fprintf(&builder, ", IsSave=%+v", d.IsSave)

	builder.WriteByte('}')

	return builder.String()
}
