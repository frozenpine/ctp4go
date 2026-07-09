package types

import (
	"math"
	"strconv"

	"github.com/frozenpine/ctp4go"
)

//go:generate stringer -type THOST_TE_RESUME_TYPE -linecomment
type THOST_TE_RESUME_TYPE int32

const (
	THOST_TERT_RESTART THOST_TE_RESUME_TYPE = 0 // RESTART
	THOST_TERT_RESUME  THOST_TE_RESUME_TYPE = 1 // RESUME
	THOST_TERT_QUICK   THOST_TE_RESUME_TYPE = 2 // QUICK
)

// TFtdcTraderIDType是一个交易所交易员代码类型
type TThostFtdcTraderIDType [21]byte

func (t TThostFtdcTraderIDType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcTraderIDType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcInvestorIDType是一个投资者代码类型
type TThostFtdcInvestorIDType [13]byte

func (t TThostFtdcInvestorIDType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcInvestorIDType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcBrokerIDType是一个经纪公司代码类型
type TThostFtdcBrokerIDType [11]byte

func (t TThostFtdcBrokerIDType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcBrokerIDType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcBrokerAbbrType是一个经纪公司简称类型
type TThostFtdcBrokerAbbrType [9]byte

func (t TThostFtdcBrokerAbbrType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcBrokerAbbrType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcBrokerNameType是一个经纪公司名称类型
type TThostFtdcBrokerNameType [81]byte

func (t TThostFtdcBrokerNameType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcBrokerNameType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcExchangeInstIDType是一个合约在交易所的代码类型
type TThostFtdcExchangeInstIDType [81]byte

func (t TThostFtdcExchangeInstIDType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcExchangeInstIDType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcOrderRefType是一个报单引用类型
type TThostFtdcOrderRefType [13]byte

func (t TThostFtdcOrderRefType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcOrderRefType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcParticipantIDType是一个会员代码类型
type TThostFtdcParticipantIDType [11]byte

func (t TThostFtdcParticipantIDType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcParticipantIDType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcUserIDType是一个用户代码类型
type TThostFtdcUserIDType [16]byte

func (t TThostFtdcUserIDType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcUserIDType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcControlParamIDType是一个开关代码类型
type TThostFtdcControlParamIDType uint8

// TFtdcControlParamValueType是一个开关设置类型
type TThostFtdcControlParamValueType [81]byte

func (t TThostFtdcControlParamValueType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcControlParamValueType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcPasswordType是一个密码类型
type TThostFtdcPasswordType [41]byte

func (t TThostFtdcPasswordType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcPasswordType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcClientIDType是一个交易编码类型
type TThostFtdcClientIDType [11]byte

func (t TThostFtdcClientIDType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcClientIDType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcInstrumentIDType是一个合约代码类型
type TThostFtdcInstrumentIDType [81]byte

func (t TThostFtdcInstrumentIDType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcInstrumentIDType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcSPMMProductIDType是一个商品群代码类型
type TThostFtdcSPMMProductIDType [41]byte

func (t TThostFtdcSPMMProductIDType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcSPMMProductIDType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcMarketIDType是一个市场代码类型
type TThostFtdcMarketIDType [31]byte

func (t TThostFtdcMarketIDType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcMarketIDType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcProductNameType是一个产品名称类型
type TThostFtdcProductNameType [21]byte

func (t TThostFtdcProductNameType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcProductNameType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcExchangeIDType是一个交易所代码类型
type TThostFtdcExchangeIDType [9]byte

func (t TThostFtdcExchangeIDType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcExchangeIDType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcExchangeNameType是一个交易所名称类型
type TThostFtdcExchangeNameType [61]byte

func (t TThostFtdcExchangeNameType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcExchangeNameType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcExchangeAbbrType是一个交易所简称类型
type TThostFtdcExchangeAbbrType [9]byte

func (t TThostFtdcExchangeAbbrType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcExchangeAbbrType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcExchangeFlagType是一个交易所标志类型
type TThostFtdcExchangeFlagType [2]byte

func (t TThostFtdcExchangeFlagType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcExchangeFlagType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcMacAddressType是一个Mac地址类型
type TThostFtdcMacAddressType [21]byte

func (t TThostFtdcMacAddressType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcMacAddressType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcSystemIDType是一个系统编号类型
type TThostFtdcSystemIDType [21]byte

func (t TThostFtdcSystemIDType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcSystemIDType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcExchangePropertyType是一个交易所属性类型
type TThostFtdcExchangePropertyType uint8

//go:generate stringer -type TThostFtdcExchangePropertyType -linecomment
const (
	THOST_FTDC_EXP_Normal          TThostFtdcExchangePropertyType = '0' // 正常
	THOST_FTDC_EXP_GenOrderByTrade TThostFtdcExchangePropertyType = '1' // 根据成交生成报单
)

// TFtdcDateType是一个日期类型
type TThostFtdcDateType [9]byte

func (t TThostFtdcDateType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcDateType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcTimeType是一个时间类型
type TThostFtdcTimeType [9]byte

func (t TThostFtdcTimeType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcTimeType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcLongTimeType是一个长时间类型
type TThostFtdcLongTimeType [13]byte

func (t TThostFtdcLongTimeType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcLongTimeType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcInstrumentNameType是一个合约名称类型
type TThostFtdcInstrumentNameType [21]byte

func (t TThostFtdcInstrumentNameType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcInstrumentNameType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcSettlementGroupIDType是一个结算组代码类型
type TThostFtdcSettlementGroupIDType [9]byte

func (t TThostFtdcSettlementGroupIDType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcSettlementGroupIDType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcOrderSysIDType是一个报单编号类型
type TThostFtdcOrderSysIDType [21]byte

func (t TThostFtdcOrderSysIDType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcOrderSysIDType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcTradeIDType是一个成交编号类型
type TThostFtdcTradeIDType [21]byte

func (t TThostFtdcTradeIDType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcTradeIDType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcCommandTypeType是一个DB命令类型类型
type TThostFtdcCommandTypeType [65]byte

func (t TThostFtdcCommandTypeType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcCommandTypeType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcIPAddressType是一个IP地址类型
type TThostFtdcIPAddressType [16]byte

func (t TThostFtdcIPAddressType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcIPAddressType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcIPPortType是一个IP端口类型
type TThostFtdcIPPortType int32

// TFtdcProductInfoType是一个产品信息类型
type TThostFtdcProductInfoType [11]byte

func (t TThostFtdcProductInfoType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcProductInfoType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcProtocolInfoType是一个协议信息类型
type TThostFtdcProtocolInfoType [11]byte

func (t TThostFtdcProtocolInfoType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcProtocolInfoType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcBusinessUnitType是一个业务单元类型
type TThostFtdcBusinessUnitType [21]byte

func (t TThostFtdcBusinessUnitType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcBusinessUnitType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcDepositSeqNoType是一个出入金流水号类型
type TThostFtdcDepositSeqNoType [15]byte

func (t TThostFtdcDepositSeqNoType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcDepositSeqNoType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcIdentifiedCardNoType是一个证件号码类型
type TThostFtdcIdentifiedCardNoType [51]byte

func (t TThostFtdcIdentifiedCardNoType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcIdentifiedCardNoType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcIdCardTypeType是一个证件类型类型
type TThostFtdcIdCardTypeType uint8

//go:generate stringer -type TThostFtdcIdCardTypeType -linecomment
const (
	THOST_FTDC_ICT_EID                    TThostFtdcIdCardTypeType = '0' // 组织机构代码
	THOST_FTDC_ICT_IDCard                 TThostFtdcIdCardTypeType = '1' // 中国公民身份证
	THOST_FTDC_ICT_OfficerIDCard          TThostFtdcIdCardTypeType = '2' // 军官证
	THOST_FTDC_ICT_PoliceIDCard           TThostFtdcIdCardTypeType = '3' // 警官证
	THOST_FTDC_ICT_SoldierIDCard          TThostFtdcIdCardTypeType = '4' // 士兵证
	THOST_FTDC_ICT_HouseholdRegister      TThostFtdcIdCardTypeType = '5' // 户口簿
	THOST_FTDC_ICT_Passport               TThostFtdcIdCardTypeType = '6' // 护照
	THOST_FTDC_ICT_TaiwanCompatriotIDCard TThostFtdcIdCardTypeType = '7' // 台胞证
	THOST_FTDC_ICT_HomeComingCard         TThostFtdcIdCardTypeType = '8' // 回乡证
	THOST_FTDC_ICT_LicenseNo              TThostFtdcIdCardTypeType = '9' // 营业执照号
	THOST_FTDC_ICT_TaxNo                  TThostFtdcIdCardTypeType = 'A' // 税务登记号当地纳税ID
	THOST_FTDC_ICT_HMMainlandTravelPermit TThostFtdcIdCardTypeType = 'B' // 港澳居民来往内地通行证
	THOST_FTDC_ICT_TwMainlandTravelPermit TThostFtdcIdCardTypeType = 'C' // 台湾居民来往大陆通行证
	THOST_FTDC_ICT_DrivingLicense         TThostFtdcIdCardTypeType = 'D' // 驾照
	THOST_FTDC_ICT_SocialID               TThostFtdcIdCardTypeType = 'F' // 当地社保ID
	THOST_FTDC_ICT_LocalID                TThostFtdcIdCardTypeType = 'G' // 当地身份证
	THOST_FTDC_ICT_BusinessRegistration   TThostFtdcIdCardTypeType = 'H' // 商业登记证
	THOST_FTDC_ICT_HKMCIDCard             TThostFtdcIdCardTypeType = 'I' // 港澳永久性居民身份证
	THOST_FTDC_ICT_AccountsPermits        TThostFtdcIdCardTypeType = 'J' // 人行开户许可证
	THOST_FTDC_ICT_OtherCard              TThostFtdcIdCardTypeType = 'x' // 其他证件
)

// TFtdcOrderLocalIDType是一个本地报单编号类型
type TThostFtdcOrderLocalIDType [13]byte

func (t TThostFtdcOrderLocalIDType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcOrderLocalIDType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcUserNameType是一个用户名称类型
type TThostFtdcUserNameType [81]byte

func (t TThostFtdcUserNameType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcUserNameType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcPartyNameType是一个参与人名称类型
type TThostFtdcPartyNameType [81]byte

func (t TThostFtdcPartyNameType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcPartyNameType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcErrorMsgType是一个错误信息类型
type TThostFtdcErrorMsgType [81]byte

func (t TThostFtdcErrorMsgType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcErrorMsgType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcFieldNameType是一个字段名类型
type TThostFtdcFieldNameType [2049]byte

func (t TThostFtdcFieldNameType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcFieldNameType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcFieldContentType是一个字段内容类型
type TThostFtdcFieldContentType [2049]byte

func (t TThostFtdcFieldContentType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcFieldContentType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcSystemNameType是一个系统名称类型
type TThostFtdcSystemNameType [41]byte

func (t TThostFtdcSystemNameType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcSystemNameType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcContentType是一个消息正文类型
type TThostFtdcContentType [501]byte

func (t TThostFtdcContentType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcContentType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcInvestorRangeType是一个投资者范围类型
type TThostFtdcInvestorRangeType uint8

//go:generate stringer -type TThostFtdcInvestorRangeType -linecomment
const (
	THOST_FTDC_IR_All    TThostFtdcInvestorRangeType = '1' // 所有
	THOST_FTDC_IR_Group  TThostFtdcInvestorRangeType = '2' // 投资者组
	THOST_FTDC_IR_Single TThostFtdcInvestorRangeType = '3' // 单一投资者
)

// TFtdcDepartmentRangeType是一个投资者范围类型
type TThostFtdcDepartmentRangeType uint8

//go:generate stringer -type TThostFtdcDepartmentRangeType -linecomment
const (
	THOST_FTDC_DR_All    TThostFtdcDepartmentRangeType = '1' // 所有
	THOST_FTDC_DR_Group  TThostFtdcDepartmentRangeType = '2' // 组织架构
	THOST_FTDC_DR_Single TThostFtdcDepartmentRangeType = '3' // 单一投资者
)

// TFtdcDataSyncStatusType是一个数据同步状态类型
type TThostFtdcDataSyncStatusType uint8

//go:generate stringer -type TThostFtdcDataSyncStatusType -linecomment
const (
	THOST_FTDC_DS_Asynchronous  TThostFtdcDataSyncStatusType = '1' // 未同步
	THOST_FTDC_DS_Synchronizing TThostFtdcDataSyncStatusType = '2' // 同步中
	THOST_FTDC_DS_Synchronized  TThostFtdcDataSyncStatusType = '3' // 已同步
)

// TFtdcBrokerDataSyncStatusType是一个经纪公司数据同步状态类型
type TThostFtdcBrokerDataSyncStatusType uint8

//go:generate stringer -type TThostFtdcBrokerDataSyncStatusType -linecomment
const (
	THOST_FTDC_BDS_Synchronized  TThostFtdcBrokerDataSyncStatusType = '1' // 已同步
	THOST_FTDC_BDS_Synchronizing TThostFtdcBrokerDataSyncStatusType = '2' // 同步中
)

// TFtdcExchangeConnectStatusType是一个交易所连接状态类型
type TThostFtdcExchangeConnectStatusType uint8

//go:generate stringer -type TThostFtdcExchangeConnectStatusType -linecomment
const (
	THOST_FTDC_ECS_NoConnection      TThostFtdcExchangeConnectStatusType = '1' // 没有任何连接
	THOST_FTDC_ECS_QryInstrumentSent TThostFtdcExchangeConnectStatusType = '2' // 已经发出合约查询请求
	THOST_FTDC_ECS_GotInformation    TThostFtdcExchangeConnectStatusType = '9' // 已经获取信息
)

// TFtdcTraderConnectStatusType是一个交易所交易员连接状态类型
type TThostFtdcTraderConnectStatusType uint8

//go:generate stringer -type TThostFtdcTraderConnectStatusType -linecomment
const (
	THOST_FTDC_TCS_NotConnected      TThostFtdcTraderConnectStatusType = '1' // 没有任何连接
	THOST_FTDC_TCS_Connected         TThostFtdcTraderConnectStatusType = '2' // 已经连接
	THOST_FTDC_TCS_QryInstrumentSent TThostFtdcTraderConnectStatusType = '3' // 已经发出合约查询请求
	THOST_FTDC_TCS_SubPrivateFlow    TThostFtdcTraderConnectStatusType = '4' // 订阅私有流
)

// TFtdcFunctionCodeType是一个功能代码类型
type TThostFtdcFunctionCodeType uint8

//go:generate stringer -type TThostFtdcFunctionCodeType -linecomment
const (
	THOST_FTDC_FC_DataAsync              TThostFtdcFunctionCodeType = '1' // 数据异步化
	THOST_FTDC_FC_ForceUserLogout        TThostFtdcFunctionCodeType = '2' // 强制用户登出
	THOST_FTDC_FC_UserPasswordUpdate     TThostFtdcFunctionCodeType = '3' // 变更管理用户口令
	THOST_FTDC_FC_BrokerPasswordUpdate   TThostFtdcFunctionCodeType = '4' // 变更经纪公司口令
	THOST_FTDC_FC_InvestorPasswordUpdate TThostFtdcFunctionCodeType = '5' // 变更投资者口令
	THOST_FTDC_FC_OrderInsert            TThostFtdcFunctionCodeType = '6' // 报单插入
	THOST_FTDC_FC_OrderAction            TThostFtdcFunctionCodeType = '7' // 报单操作
	THOST_FTDC_FC_SyncSystemData         TThostFtdcFunctionCodeType = '8' // 同步系统数据
	THOST_FTDC_FC_SyncBrokerData         TThostFtdcFunctionCodeType = '9' // 同步经纪公司数据
	THOST_FTDC_FC_BachSyncBrokerData     TThostFtdcFunctionCodeType = 'A' // 批量同步经纪公司数据
	THOST_FTDC_FC_SuperQuery             TThostFtdcFunctionCodeType = 'B' // 超级查询
	THOST_FTDC_FC_ParkedOrderInsert      TThostFtdcFunctionCodeType = 'C' // 预埋报单插入
	THOST_FTDC_FC_ParkedOrderAction      TThostFtdcFunctionCodeType = 'D' // 预埋报单操作
	THOST_FTDC_FC_SyncOTP                TThostFtdcFunctionCodeType = 'E' // 同步动态令牌
	THOST_FTDC_FC_DeleteOrder            TThostFtdcFunctionCodeType = 'F' // 删除未知单
)

// TFtdcBrokerFunctionCodeType是一个经纪公司功能代码类型
type TThostFtdcBrokerFunctionCodeType uint8

//go:generate stringer -type TThostFtdcBrokerFunctionCodeType -linecomment
const (
	THOST_FTDC_BFC_ForceUserLogout    TThostFtdcBrokerFunctionCodeType = '1' // 强制用户登出
	THOST_FTDC_BFC_UserPasswordUpdate TThostFtdcBrokerFunctionCodeType = '2' // 变更用户口令
	THOST_FTDC_BFC_SyncBrokerData     TThostFtdcBrokerFunctionCodeType = '3' // 同步经纪公司数据
	THOST_FTDC_BFC_BachSyncBrokerData TThostFtdcBrokerFunctionCodeType = '4' // 批量同步经纪公司数据
	THOST_FTDC_BFC_OrderInsert        TThostFtdcBrokerFunctionCodeType = '5' // 报单插入
	THOST_FTDC_BFC_OrderAction        TThostFtdcBrokerFunctionCodeType = '6' // 报单操作
	THOST_FTDC_BFC_AllQuery           TThostFtdcBrokerFunctionCodeType = '7' // 全部查询
	THOST_FTDC_BFC_log                TThostFtdcBrokerFunctionCodeType = 'a' // 系统功能：登入登出修改密码等
	THOST_FTDC_BFC_BaseQry            TThostFtdcBrokerFunctionCodeType = 'b' // 基本查询：查询基础数据，如合约，交易所等常量
	THOST_FTDC_BFC_TradeQry           TThostFtdcBrokerFunctionCodeType = 'c' // 交易查询：如查成交，委托
	THOST_FTDC_BFC_Trade              TThostFtdcBrokerFunctionCodeType = 'd' // 交易功能：报单，撤单
	THOST_FTDC_BFC_Virement           TThostFtdcBrokerFunctionCodeType = 'e' // 银期转账
	THOST_FTDC_BFC_Risk               TThostFtdcBrokerFunctionCodeType = 'f' // 风险监控
	THOST_FTDC_BFC_Session            TThostFtdcBrokerFunctionCodeType = 'g' // 查询管理：查询会话，踢人等
	THOST_FTDC_BFC_RiskNoticeCtl      TThostFtdcBrokerFunctionCodeType = 'h' // 风控通知控制
	THOST_FTDC_BFC_RiskNotice         TThostFtdcBrokerFunctionCodeType = 'i' // 风控通知发送
	THOST_FTDC_BFC_BrokerDeposit      TThostFtdcBrokerFunctionCodeType = 'j' // 察看经纪公司资金权限
	THOST_FTDC_BFC_QueryFund          TThostFtdcBrokerFunctionCodeType = 'k' // 资金查询
	THOST_FTDC_BFC_QueryOrder         TThostFtdcBrokerFunctionCodeType = 'l' // 报单查询
	THOST_FTDC_BFC_QueryTrade         TThostFtdcBrokerFunctionCodeType = 'm' // 成交查询
	THOST_FTDC_BFC_QueryPosition      TThostFtdcBrokerFunctionCodeType = 'n' // 持仓查询
	THOST_FTDC_BFC_QueryMarketData    TThostFtdcBrokerFunctionCodeType = 'o' // 行情查询
	THOST_FTDC_BFC_QueryUserEvent     TThostFtdcBrokerFunctionCodeType = 'p' // 用户事件查询
	THOST_FTDC_BFC_QueryRiskNotify    TThostFtdcBrokerFunctionCodeType = 'q' // 风险通知查询
	THOST_FTDC_BFC_QueryFundChange    TThostFtdcBrokerFunctionCodeType = 'r' // 出入金查询
	THOST_FTDC_BFC_QueryInvestor      TThostFtdcBrokerFunctionCodeType = 's' // 投资者信息查询
	THOST_FTDC_BFC_QueryTradingCode   TThostFtdcBrokerFunctionCodeType = 't' // 交易编码查询
	THOST_FTDC_BFC_ForceClose         TThostFtdcBrokerFunctionCodeType = 'u' // 强平
	THOST_FTDC_BFC_PressTest          TThostFtdcBrokerFunctionCodeType = 'v' // 压力测试
	THOST_FTDC_BFC_RemainCalc         TThostFtdcBrokerFunctionCodeType = 'w' // 权益反算
	THOST_FTDC_BFC_NetPositionInd     TThostFtdcBrokerFunctionCodeType = 'x' // 净持仓保证金指标
	THOST_FTDC_BFC_RiskPredict        TThostFtdcBrokerFunctionCodeType = 'y' // 风险预算
	THOST_FTDC_BFC_DataExport         TThostFtdcBrokerFunctionCodeType = 'z' // 数据导出
	THOST_FTDC_BFC_RiskTargetSetup    TThostFtdcBrokerFunctionCodeType = 'A' // 风控指标设置
	THOST_FTDC_BFC_MarketDataWarn     TThostFtdcBrokerFunctionCodeType = 'B' // 行情预警
	THOST_FTDC_BFC_QryBizNotice       TThostFtdcBrokerFunctionCodeType = 'C' // 业务通知查询
	THOST_FTDC_BFC_CfgBizNotice       TThostFtdcBrokerFunctionCodeType = 'D' // 业务通知模板设置
	THOST_FTDC_BFC_SyncOTP            TThostFtdcBrokerFunctionCodeType = 'E' // 同步动态令牌
	THOST_FTDC_BFC_SendBizNotice      TThostFtdcBrokerFunctionCodeType = 'F' // 发送业务通知
	THOST_FTDC_BFC_CfgRiskLevelStd    TThostFtdcBrokerFunctionCodeType = 'G' // 风险级别标准设置
	THOST_FTDC_BFC_TbCommand          TThostFtdcBrokerFunctionCodeType = 'H' // 交易终端应急功能
	THOST_FTDC_BFC_DeleteOrder        TThostFtdcBrokerFunctionCodeType = 'J' // 删除未知单
	THOST_FTDC_BFC_ParkedOrderInsert  TThostFtdcBrokerFunctionCodeType = 'K' // 预埋报单插入
	THOST_FTDC_BFC_ParkedOrderAction  TThostFtdcBrokerFunctionCodeType = 'L' // 预埋报单操作
	THOST_FTDC_BFC_ExecOrderNoCheck   TThostFtdcBrokerFunctionCodeType = 'M' // 资金不够仍允许行权
)

// TFtdcOrderActionStatusType是一个报单操作状态类型
type TThostFtdcOrderActionStatusType uint8

//go:generate stringer -type TThostFtdcOrderActionStatusType -linecomment
const (
	THOST_FTDC_OAS_Submitted TThostFtdcOrderActionStatusType = 'a' // 已经提交
	THOST_FTDC_OAS_Accepted  TThostFtdcOrderActionStatusType = 'b' // 已经接受
	THOST_FTDC_OAS_Rejected  TThostFtdcOrderActionStatusType = 'c' // 已经被拒绝
)

// TFtdcOrderStatusType是一个报单状态类型
type TThostFtdcOrderStatusType uint8

//go:generate stringer -type TThostFtdcOrderStatusType -linecomment
const (
	THOST_FTDC_OST_AllTraded             TThostFtdcOrderStatusType = '0' // 全部成交
	THOST_FTDC_OST_PartTradedQueueing    TThostFtdcOrderStatusType = '1' // 部分成交还在队列中
	THOST_FTDC_OST_PartTradedNotQueueing TThostFtdcOrderStatusType = '2' // 部分成交不在队列中
	THOST_FTDC_OST_NoTradeQueueing       TThostFtdcOrderStatusType = '3' // 未成交还在队列中
	THOST_FTDC_OST_NoTradeNotQueueing    TThostFtdcOrderStatusType = '4' // 未成交不在队列中
	THOST_FTDC_OST_Canceled              TThostFtdcOrderStatusType = '5' // 撤单
	THOST_FTDC_OST_Unknown               TThostFtdcOrderStatusType = 'a' // 未知
	THOST_FTDC_OST_NotTouched            TThostFtdcOrderStatusType = 'b' // 尚未触发
	THOST_FTDC_OST_Touched               TThostFtdcOrderStatusType = 'c' // 已触发
)

// TFtdcOrderSubmitStatusType是一个报单提交状态类型
type TThostFtdcOrderSubmitStatusType uint8

//go:generate stringer -type TThostFtdcOrderSubmitStatusType -linecomment
const (
	THOST_FTDC_OSS_InsertSubmitted TThostFtdcOrderSubmitStatusType = '0' // 已经提交
	THOST_FTDC_OSS_CancelSubmitted TThostFtdcOrderSubmitStatusType = '1' // 撤单已经提交
	THOST_FTDC_OSS_ModifySubmitted TThostFtdcOrderSubmitStatusType = '2' // 修改已经提交
	THOST_FTDC_OSS_Accepted        TThostFtdcOrderSubmitStatusType = '3' // 已经接受
	THOST_FTDC_OSS_InsertRejected  TThostFtdcOrderSubmitStatusType = '4' // 报单已经被拒绝
	THOST_FTDC_OSS_CancelRejected  TThostFtdcOrderSubmitStatusType = '5' // 撤单已经被拒绝
	THOST_FTDC_OSS_ModifyRejected  TThostFtdcOrderSubmitStatusType = '6' // 改单已经被拒绝
)

// TFtdcRequestStatusType是一个套利套保确认请求状态类型
type TThostFtdcRequestStatusType uint8

//go:generate stringer -type TThostFtdcRequestStatusType -linecomment
const (
	THOST_FTDC_RS_Submitted TThostFtdcRequestStatusType = 'a' // 已经提交
	THOST_FTDC_RS_Canceled  TThostFtdcRequestStatusType = 'c' // 已经撤销
	THOST_FTDC_RS_DELETE    TThostFtdcRequestStatusType = '0' // 删除
	THOST_FTDC_RS_HANGUP    TThostFtdcRequestStatusType = '1' // 挂起
	THOST_FTDC_RS_ACTIVATE  TThostFtdcRequestStatusType = '2' // 激活
	THOST_FTDC_RS_All       TThostFtdcRequestStatusType = '1' // 按所有品种
	THOST_FTDC_RS_ByProduct TThostFtdcRequestStatusType = '2' // 按品种
)

// TFtdcPositionDateType是一个持仓日期类型
type TThostFtdcPositionDateType uint8

//go:generate stringer -type TThostFtdcPositionDateType -linecomment
const (
	THOST_FTDC_PSD_Today   TThostFtdcPositionDateType = '1' // 今日持仓
	THOST_FTDC_PSD_History TThostFtdcPositionDateType = '2' // 历史持仓
)

// TFtdcPositionDateTypeType是一个持仓日期类型类型
type TThostFtdcPositionDateTypeType uint8

//go:generate stringer -type TThostFtdcPositionDateTypeType -linecomment
const (
	THOST_FTDC_PDT_UseHistory   TThostFtdcPositionDateTypeType = '1' // 使用历史持仓
	THOST_FTDC_PDT_NoUseHistory TThostFtdcPositionDateTypeType = '2' // 不使用历史持仓
)

// TFtdcTradingRoleType是一个交易角色类型
type TThostFtdcTradingRoleType uint8

//go:generate stringer -type TThostFtdcTradingRoleType -linecomment
const (
	THOST_FTDC_ER_Broker TThostFtdcTradingRoleType = '1' // 代理
	THOST_FTDC_ER_Host   TThostFtdcTradingRoleType = '2' // 自营
	THOST_FTDC_ER_Maker  TThostFtdcTradingRoleType = '3' // 做市商
)

// TFtdcProductClassType是一个产品类型类型
type TThostFtdcProductClassType uint8

//go:generate stringer -type TThostFtdcProductClassType -linecomment
const (
	THOST_FTDC_PC_Futures     TThostFtdcProductClassType = '1' // 期货
	THOST_FTDC_PC_Options     TThostFtdcProductClassType = '2' // 期货期权
	THOST_FTDC_PC_Combination TThostFtdcProductClassType = '3' // 组合
	THOST_FTDC_PC_Spot        TThostFtdcProductClassType = '4' // 即期
	THOST_FTDC_PC_EFP         TThostFtdcProductClassType = '5' // 期转现
	THOST_FTDC_PC_SpotOption  TThostFtdcProductClassType = '6' // 现货期权
	THOST_FTDC_PC_TAS         TThostFtdcProductClassType = '7' // TAS合约
)

// TFtdcInstLifePhaseType是一个合约生命周期状态类型
type TThostFtdcInstLifePhaseType uint8

//go:generate stringer -type TThostFtdcInstLifePhaseType -linecomment
const (
	THOST_FTDC_IP_NotStart TThostFtdcInstLifePhaseType = '0' // 未上市
	THOST_FTDC_IP_Started  TThostFtdcInstLifePhaseType = '1' // 上市
	THOST_FTDC_IP_Pause    TThostFtdcInstLifePhaseType = '2' // 停牌
	THOST_FTDC_IP_Expired  TThostFtdcInstLifePhaseType = '3' // 到期
)

// TFtdcDirectionType是一个买卖方向类型
type TThostFtdcDirectionType uint8

//go:generate stringer -type TThostFtdcDirectionType -linecomment
const (
	THOST_FTDC_D_Buy  TThostFtdcDirectionType = '0' // 买
	THOST_FTDC_D_Sell TThostFtdcDirectionType = '1' // 卖
)

// TFtdcPositionTypeType是一个持仓类型类型
type TThostFtdcPositionTypeType uint8

//go:generate stringer -type TThostFtdcPositionTypeType -linecomment
const (
	THOST_FTDC_PT_Net   TThostFtdcPositionTypeType = '1' // 净持仓
	THOST_FTDC_PT_Gross TThostFtdcPositionTypeType = '2' // 综合持仓
)

// TFtdcPosiDirectionType是一个持仓多空方向类型
type TThostFtdcPosiDirectionType uint8

//go:generate stringer -type TThostFtdcPosiDirectionType -linecomment
const (
	THOST_FTDC_PD_Net   TThostFtdcPosiDirectionType = '1' // 净
	THOST_FTDC_PD_Long  TThostFtdcPosiDirectionType = '2' // 多头
	THOST_FTDC_PD_Short TThostFtdcPosiDirectionType = '3' // 空头
)

// TFtdcSysSettlementStatusType是一个系统结算状态类型
type TThostFtdcSysSettlementStatusType uint8

//go:generate stringer -type TThostFtdcSysSettlementStatusType -linecomment
const (
	THOST_FTDC_SS_NonActive          TThostFtdcSysSettlementStatusType = '1' // 不活跃
	THOST_FTDC_SS_Startup            TThostFtdcSysSettlementStatusType = '2' // 启动
	THOST_FTDC_SS_Operating          TThostFtdcSysSettlementStatusType = '3' // 操作
	THOST_FTDC_SS_Settlement         TThostFtdcSysSettlementStatusType = '4' // 结算
	THOST_FTDC_SS_SettlementFinished TThostFtdcSysSettlementStatusType = '5' // 结算完成
)

// TFtdcRatioAttrType是一个费率属性类型
type TThostFtdcRatioAttrType uint8

//go:generate stringer -type TThostFtdcRatioAttrType -linecomment
const (
	THOST_FTDC_RA_Trade      TThostFtdcRatioAttrType = '0' // 交易费率
	THOST_FTDC_RA_Settlement TThostFtdcRatioAttrType = '1' // 结算费率
)

// TFtdcHedgeFlagType是一个投机套保标志类型
type TThostFtdcHedgeFlagType uint8

//go:generate stringer -type TThostFtdcHedgeFlagType -linecomment
const (
	THOST_FTDC_HF_Speculation TThostFtdcHedgeFlagType = '1' // 投机
	THOST_FTDC_HF_Arbitrage   TThostFtdcHedgeFlagType = '2' // 套利
	THOST_FTDC_HF_Hedge       TThostFtdcHedgeFlagType = '3' // 套保
	THOST_FTDC_HF_MarketMaker TThostFtdcHedgeFlagType = '5' // 做市商
	THOST_FTDC_HF_SpecHedge   TThostFtdcHedgeFlagType = '6' // 第一腿投机第二腿套保
	THOST_FTDC_HF_HedgeSpec   TThostFtdcHedgeFlagType = '7' // 第一腿套保第二腿投机
)

// TFtdcBillHedgeFlagType是一个投机套保标志类型
type TThostFtdcBillHedgeFlagType uint8

//go:generate stringer -type TThostFtdcBillHedgeFlagType -linecomment
const (
	THOST_FTDC_BHF_Speculation TThostFtdcBillHedgeFlagType = '1' // 投机
	THOST_FTDC_BHF_Arbitrage   TThostFtdcBillHedgeFlagType = '2' // 套利
	THOST_FTDC_BHF_Hedge       TThostFtdcBillHedgeFlagType = '3' // 套保
)

// TFtdcClientIDTypeType是一个交易编码类型类型
type TThostFtdcClientIDTypeType uint8

//go:generate stringer -type TThostFtdcClientIDTypeType -linecomment
const (
	THOST_FTDC_CIDT_Speculation TThostFtdcClientIDTypeType = '1' // 投机
	THOST_FTDC_CIDT_Arbitrage   TThostFtdcClientIDTypeType = '2' // 套利
	THOST_FTDC_CIDT_Hedge       TThostFtdcClientIDTypeType = '3' // 套保
	THOST_FTDC_CIDT_MarketMaker TThostFtdcClientIDTypeType = '5' // 做市商
)

// TFtdcOrderPriceTypeType是一个报单价格条件类型
type TThostFtdcOrderPriceTypeType uint8

//go:generate stringer -type TThostFtdcOrderPriceTypeType -linecomment
const (
	THOST_FTDC_OPT_AnyPrice                TThostFtdcOrderPriceTypeType = '1' // 任意价
	THOST_FTDC_OPT_LimitPrice              TThostFtdcOrderPriceTypeType = '2' // 限价
	THOST_FTDC_OPT_BestPrice               TThostFtdcOrderPriceTypeType = '3' // 最优价
	THOST_FTDC_OPT_LastPrice               TThostFtdcOrderPriceTypeType = '4' // 最新价
	THOST_FTDC_OPT_LastPricePlusOneTicks   TThostFtdcOrderPriceTypeType = '5' // 最新价浮动上浮1个ticks
	THOST_FTDC_OPT_LastPricePlusTwoTicks   TThostFtdcOrderPriceTypeType = '6' // 最新价浮动上浮2个ticks
	THOST_FTDC_OPT_LastPricePlusThreeTicks TThostFtdcOrderPriceTypeType = '7' // 最新价浮动上浮3个ticks
	THOST_FTDC_OPT_AskPrice1               TThostFtdcOrderPriceTypeType = '8' // 卖一价
	THOST_FTDC_OPT_AskPrice1PlusOneTicks   TThostFtdcOrderPriceTypeType = '9' // 卖一价浮动上浮1个ticks
	THOST_FTDC_OPT_AskPrice1PlusTwoTicks   TThostFtdcOrderPriceTypeType = 'A' // 卖一价浮动上浮2个ticks
	THOST_FTDC_OPT_AskPrice1PlusThreeTicks TThostFtdcOrderPriceTypeType = 'B' // 卖一价浮动上浮3个ticks
	THOST_FTDC_OPT_BidPrice1               TThostFtdcOrderPriceTypeType = 'C' // 买一价
	THOST_FTDC_OPT_BidPrice1PlusOneTicks   TThostFtdcOrderPriceTypeType = 'D' // 买一价浮动上浮1个ticks
	THOST_FTDC_OPT_BidPrice1PlusTwoTicks   TThostFtdcOrderPriceTypeType = 'E' // 买一价浮动上浮2个ticks
	THOST_FTDC_OPT_BidPrice1PlusThreeTicks TThostFtdcOrderPriceTypeType = 'F' // 买一价浮动上浮3个ticks
	THOST_FTDC_OPT_FiveLevelPrice          TThostFtdcOrderPriceTypeType = 'G' // 五档价
)

// TFtdcOffsetFlagType是一个开平标志类型
type TThostFtdcOffsetFlagType uint8

//go:generate stringer -type TThostFtdcOffsetFlagType -linecomment
const (
	THOST_FTDC_OF_Open            TThostFtdcOffsetFlagType = '0' // 开仓
	THOST_FTDC_OF_Close           TThostFtdcOffsetFlagType = '1' // 平仓
	THOST_FTDC_OF_ForceClose      TThostFtdcOffsetFlagType = '2' // 强平
	THOST_FTDC_OF_CloseToday      TThostFtdcOffsetFlagType = '3' // 平今
	THOST_FTDC_OF_CloseYesterday  TThostFtdcOffsetFlagType = '4' // 平昨
	THOST_FTDC_OF_ForceOff        TThostFtdcOffsetFlagType = '5' // 强减
	THOST_FTDC_OF_LocalForceClose TThostFtdcOffsetFlagType = '6' // 本地强平
)

// TFtdcForceCloseReasonType是一个强平原因类型
type TThostFtdcForceCloseReasonType uint8

//go:generate stringer -type TThostFtdcForceCloseReasonType -linecomment
const (
	THOST_FTDC_FCC_NotForceClose           TThostFtdcForceCloseReasonType = '0' // 非强平
	THOST_FTDC_FCC_LackDeposit             TThostFtdcForceCloseReasonType = '1' // 资金不足
	THOST_FTDC_FCC_ClientOverPositionLimit TThostFtdcForceCloseReasonType = '2' // 客户超仓
	THOST_FTDC_FCC_MemberOverPositionLimit TThostFtdcForceCloseReasonType = '3' // 会员超仓
	THOST_FTDC_FCC_NotMultiple             TThostFtdcForceCloseReasonType = '4' // 持仓非整数倍
	THOST_FTDC_FCC_Violation               TThostFtdcForceCloseReasonType = '5' // 违规
	THOST_FTDC_FCC_Other                   TThostFtdcForceCloseReasonType = '6' // 其它
	THOST_FTDC_FCC_PersonDeliv             TThostFtdcForceCloseReasonType = '7' // 自然人临近交割
)

// TFtdcOrderTypeType是一个报单类型类型
type TThostFtdcOrderTypeType uint8

//go:generate stringer -type TThostFtdcOrderTypeType -linecomment
const (
	THOST_FTDC_ORDT_Normal                TThostFtdcOrderTypeType = '0' // 正常
	THOST_FTDC_ORDT_DeriveFromQuote       TThostFtdcOrderTypeType = '1' // 报价衍生
	THOST_FTDC_ORDT_DeriveFromCombination TThostFtdcOrderTypeType = '2' // 组合衍生
	THOST_FTDC_ORDT_Combination           TThostFtdcOrderTypeType = '3' // 组合报单
	THOST_FTDC_ORDT_ConditionalOrder      TThostFtdcOrderTypeType = '4' // 条件单
	THOST_FTDC_ORDT_Swap                  TThostFtdcOrderTypeType = '5' // 互换单
)

// TFtdcTimeConditionType是一个有效期类型类型
type TThostFtdcTimeConditionType uint8

//go:generate stringer -type TThostFtdcTimeConditionType -linecomment
const (
	THOST_FTDC_TC_IOC TThostFtdcTimeConditionType = '1' // 立即完成，否则撤销
	THOST_FTDC_TC_GFS TThostFtdcTimeConditionType = '2' // 本节有效
	THOST_FTDC_TC_GFD TThostFtdcTimeConditionType = '3' // 当日有效
	THOST_FTDC_TC_GTD TThostFtdcTimeConditionType = '4' // 指定日期前有效
	THOST_FTDC_TC_GTC TThostFtdcTimeConditionType = '5' // 撤销前有效
	THOST_FTDC_TC_GFA TThostFtdcTimeConditionType = '6' // 集合竞价有效
)

// TFtdcVolumeConditionType是一个成交量类型类型
type TThostFtdcVolumeConditionType uint8

//go:generate stringer -type TThostFtdcVolumeConditionType -linecomment
const (
	THOST_FTDC_VC_AV TThostFtdcVolumeConditionType = '1' // 任何数量
	THOST_FTDC_VC_MV TThostFtdcVolumeConditionType = '2' // 最小数量
	THOST_FTDC_VC_CV TThostFtdcVolumeConditionType = '3' // 全部数量
)

// TFtdcContingentConditionType是一个触发条件类型
type TThostFtdcContingentConditionType uint8

//go:generate stringer -type TThostFtdcContingentConditionType -linecomment
const (
	THOST_FTDC_CC_Immediately                    TThostFtdcContingentConditionType = '1' // 立即
	THOST_FTDC_CC_Touch                          TThostFtdcContingentConditionType = '2' // 止损
	THOST_FTDC_CC_TouchProfit                    TThostFtdcContingentConditionType = '3' // 止赢
	THOST_FTDC_CC_ParkedOrder                    TThostFtdcContingentConditionType = '4' // 预埋单
	THOST_FTDC_CC_LastPriceGreaterThanStopPrice  TThostFtdcContingentConditionType = '5' // 最新价大于条件价
	THOST_FTDC_CC_LastPriceGreaterEqualStopPrice TThostFtdcContingentConditionType = '6' // 最新价大于等于条件价
	THOST_FTDC_CC_LastPriceLesserThanStopPrice   TThostFtdcContingentConditionType = '7' // 最新价小于条件价
	THOST_FTDC_CC_LastPriceLesserEqualStopPrice  TThostFtdcContingentConditionType = '8' // 最新价小于等于条件价
	THOST_FTDC_CC_AskPriceGreaterThanStopPrice   TThostFtdcContingentConditionType = '9' // 卖一价大于条件价
	THOST_FTDC_CC_AskPriceGreaterEqualStopPrice  TThostFtdcContingentConditionType = 'A' // 卖一价大于等于条件价
	THOST_FTDC_CC_AskPriceLesserThanStopPrice    TThostFtdcContingentConditionType = 'B' // 卖一价小于条件价
	THOST_FTDC_CC_AskPriceLesserEqualStopPrice   TThostFtdcContingentConditionType = 'C' // 卖一价小于等于条件价
	THOST_FTDC_CC_BidPriceGreaterThanStopPrice   TThostFtdcContingentConditionType = 'D' // 买一价大于条件价
	THOST_FTDC_CC_BidPriceGreaterEqualStopPrice  TThostFtdcContingentConditionType = 'E' // 买一价大于等于条件价
	THOST_FTDC_CC_BidPriceLesserThanStopPrice    TThostFtdcContingentConditionType = 'F' // 买一价小于条件价
	THOST_FTDC_CC_BidPriceLesserEqualStopPrice   TThostFtdcContingentConditionType = 'H' // 买一价小于等于条件价
)

// TFtdcActionFlagType是一个操作标志类型
type TThostFtdcActionFlagType uint8

//go:generate stringer -type TThostFtdcActionFlagType -linecomment
const (
	THOST_FTDC_AF_Delete TThostFtdcActionFlagType = '0' // 删除
	THOST_FTDC_AF_Modify TThostFtdcActionFlagType = '3' // 修改
)

// TFtdcMKActionModeType是一个做市商批量撤单类型
type TThostFtdcMKActionModeType uint8

//go:generate stringer -type TThostFtdcMKActionModeType -linecomment
const (
	THOST_FTDC_MKAM_Product  TThostFtdcMKActionModeType = '1' // 产品
	THOST_FTDC_MKAM_OrderRef TThostFtdcMKActionModeType = '2' // OrderRef
)

// TFtdcLowerLimitType是批量撤单下限类型
type TThostFtdcLowerLimitType [81]byte

func (t TThostFtdcLowerLimitType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcLowerLimitType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcUpperLimitType是批量撤单上限类型
type TThostFtdcUpperLimitType [81]byte

func (t TThostFtdcUpperLimitType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcUpperLimitType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcTradingRightType是一个交易权限类型
type TThostFtdcTradingRightType uint8

//go:generate stringer -type TThostFtdcTradingRightType -linecomment
const (
	THOST_FTDC_TR_Allow     TThostFtdcTradingRightType = '0' // 可以交易
	THOST_FTDC_TR_CloseOnly TThostFtdcTradingRightType = '1' // 只能平仓
	THOST_FTDC_TR_Forbidden TThostFtdcTradingRightType = '2' // 不能交易
)

// TFtdcOrderSourceType是一个报单来源类型
type TThostFtdcOrderSourceType uint8

//go:generate stringer -type TThostFtdcOrderSourceType -linecomment
const (
	THOST_FTDC_OSRC_Participant   TThostFtdcOrderSourceType = '0' // 来自参与者
	THOST_FTDC_OSRC_Administrator TThostFtdcOrderSourceType = '1' // 来自管理员
)

// TFtdcTradeTypeType是一个成交类型类型
type TThostFtdcTradeTypeType uint8

//go:generate stringer -type TThostFtdcTradeTypeType -linecomment
const (
	THOST_FTDC_TRDT_SplitCombination   TThostFtdcTradeTypeType = '#' // 组合持仓拆分为单一持仓,初始化不应包含该类型的持仓
	THOST_FTDC_TRDT_Common             TThostFtdcTradeTypeType = '0' // 普通成交
	THOST_FTDC_TRDT_OptionsExecution   TThostFtdcTradeTypeType = '1' // 期权执行
	THOST_FTDC_TRDT_OTC                TThostFtdcTradeTypeType = '2' // OTC成交
	THOST_FTDC_TRDT_EFPDerived         TThostFtdcTradeTypeType = '3' // 期转现衍生成交
	THOST_FTDC_TRDT_CombinationDerived TThostFtdcTradeTypeType = '4' // 组合衍生成交
)

// TFtdcPriceSourceType是一个成交价来源类型
type TThostFtdcPriceSourceType uint8

//go:generate stringer -type TThostFtdcPriceSourceType -linecomment
const (
	THOST_FTDC_PSRC_LastPrice TThostFtdcPriceSourceType = '0' // 前成交价
	THOST_FTDC_PSRC_Buy       TThostFtdcPriceSourceType = '1' // 买委托价
	THOST_FTDC_PSRC_Sell      TThostFtdcPriceSourceType = '2' // 卖委托价
)

// TFtdcInstrumentStatusType是一个合约交易状态类型
type TThostFtdcInstrumentStatusType uint8

//go:generate stringer -type TThostFtdcInstrumentStatusType -linecomment
const (
	THOST_FTDC_IS_BeforeTrading   TThostFtdcInstrumentStatusType = '0' // 开盘前
	THOST_FTDC_IS_NoTrading       TThostFtdcInstrumentStatusType = '1' // 非交易
	THOST_FTDC_IS_Continous       TThostFtdcInstrumentStatusType = '2' // 连续交易
	THOST_FTDC_IS_AuctionOrdering TThostFtdcInstrumentStatusType = '3' // 集合竞价报单
	THOST_FTDC_IS_AuctionBalance  TThostFtdcInstrumentStatusType = '4' // 集合竞价价格平衡
	THOST_FTDC_IS_AuctionMatch    TThostFtdcInstrumentStatusType = '5' // 集合竞价撮合
	THOST_FTDC_IS_Closed          TThostFtdcInstrumentStatusType = '6' // 收盘
)

// TFtdcInstStatusEnterReasonType是一个品种进入交易状态原因类型
type TThostFtdcInstStatusEnterReasonType uint8

//go:generate stringer -type TThostFtdcInstStatusEnterReasonType -linecomment
const (
	THOST_FTDC_IER_Automatic TThostFtdcInstStatusEnterReasonType = '1' // 自动切换
	THOST_FTDC_IER_Manual    TThostFtdcInstStatusEnterReasonType = '2' // 手动切换
	THOST_FTDC_IER_Fuse      TThostFtdcInstStatusEnterReasonType = '3' // 熔断
)

// TFtdcOrderActionRefType是一个报单操作引用类型
type TThostFtdcOrderActionRefType int32

// TFtdcInstallCountType是一个安装数量类型
type TThostFtdcInstallCountType int32

// TFtdcInstallIDType是一个安装编号类型
type TThostFtdcInstallIDType int32

// TFtdcErrorIDType是一个错误代码类型
type TThostFtdcErrorIDType int32

// TFtdcSettlementIDType是一个结算编号类型
type TThostFtdcSettlementIDType int32

// TFtdcVolumeType是一个数量类型
type TThostFtdcVolumeType int32

// TFtdcFrontIDType是一个前置编号类型
type TThostFtdcFrontIDType int32

// TFtdcSessionIDType是一个会话编号类型
type TThostFtdcSessionIDType int32

// TFtdcSequenceNoType是一个序号类型
type TThostFtdcSequenceNoType int32

// TFtdcCommandNoType是一个DB命令序号类型
type TThostFtdcCommandNoType int32

// TFtdcMillisecType是一个时间（毫秒）类型
type TThostFtdcMillisecType int32

// TFtdcVolumeMultipleType是一个合约数量乘数类型
type TThostFtdcVolumeMultipleType int32

// TFtdcTradingSegmentSNType是一个交易阶段编号类型
type TThostFtdcTradingSegmentSNType int32

// TFtdcRequestIDType是一个请求编号类型
type TThostFtdcRequestIDType int32

// TFtdcYearType是一个年份类型
type TThostFtdcYearType int32

// TFtdcMonthType是一个月份类型
type TThostFtdcMonthType int32

// TFtdcBoolType是一个布尔型类型
type TThostFtdcBoolType int32

// TFtdcPriceType是一个价格类型
type TThostFtdcPriceType float64

func (t TThostFtdcPriceType) IsValid() bool {
	return float64(t) != math.MaxFloat64
}

func (t TThostFtdcPriceType) String() string {
	if t.IsValid() {
		return strconv.FormatFloat(float64(t), 'f', 6, 64)
	} else {
		return "DBL_MAX"
	}
}

// TFtdcCombOffsetFlagType是一个组合开平标志类型
type TThostFtdcCombOffsetFlagType [5]byte

func (t TThostFtdcCombOffsetFlagType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcCombOffsetFlagType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcCombHedgeFlagType是一个组合投机套保标志类型
type TThostFtdcCombHedgeFlagType [5]byte

func (t TThostFtdcCombHedgeFlagType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcCombHedgeFlagType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcRatioType是一个比率类型
type TThostFtdcRatioType float64

func (t TThostFtdcRatioType) IsValid() bool {
	return float64(t) != math.MaxFloat64
}

func (t TThostFtdcRatioType) String() string {
	if t.IsValid() {
		return strconv.FormatFloat(float64(t), 'f', 6, 64)
	} else {
		return "DBL_MAX"
	}
}

// TFtdcMoneyType是一个资金类型
type TThostFtdcMoneyType float64

func (t TThostFtdcMoneyType) IsValid() bool {
	return float64(t) != math.MaxFloat64
}

func (t TThostFtdcMoneyType) String() string {
	if t.IsValid() {
		return strconv.FormatFloat(float64(t), 'f', 6, 64)
	} else {
		return "DBL_MAX"
	}
}

// TFtdcLargeVolumeType是一个大额数量类型
type TThostFtdcLargeVolumeType float64

func (t TThostFtdcLargeVolumeType) IsValid() bool {
	return float64(t) != math.MaxFloat64
}

func (t TThostFtdcLargeVolumeType) String() string {
	if t.IsValid() {
		return strconv.FormatFloat(float64(t), 'f', 6, 64)
	} else {
		return "DBL_MAX"
	}
}

// TFtdcSequenceSeriesType是一个序列系列号类型
type TThostFtdcSequenceSeriesType int16

// TFtdcCommPhaseNoType是一个通讯时段编号类型
type TThostFtdcCommPhaseNoType int16

// TFtdcSequenceLabelType是一个序列编号类型
type TThostFtdcSequenceLabelType [2]byte

func (t TThostFtdcSequenceLabelType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcSequenceLabelType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcUnderlyingMultipleType是一个基础商品乘数类型
type TThostFtdcUnderlyingMultipleType float64

func (t TThostFtdcUnderlyingMultipleType) IsValid() bool {
	return float64(t) != math.MaxFloat64
}

func (t TThostFtdcUnderlyingMultipleType) String() string {
	if t.IsValid() {
		return strconv.FormatFloat(float64(t), 'f', 6, 64)
	} else {
		return "DBL_MAX"
	}
}

// TFtdcPriorityType是一个优先级类型
type TThostFtdcPriorityType int32

// TFtdcContractCodeType是一个合同编号类型
type TThostFtdcContractCodeType [41]byte

func (t TThostFtdcContractCodeType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcContractCodeType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcCityType是一个市类型
type TThostFtdcCityType [51]byte

func (t TThostFtdcCityType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcCityType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcIsStockType是一个是否股民类型
type TThostFtdcIsStockType [11]byte

func (t TThostFtdcIsStockType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcIsStockType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcChannelType是一个渠道类型
type TThostFtdcChannelType [51]byte

func (t TThostFtdcChannelType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcChannelType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcAddressType是一个通讯地址类型
type TThostFtdcAddressType [101]byte

func (t TThostFtdcAddressType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcAddressType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcZipCodeType是一个邮政编码类型
type TThostFtdcZipCodeType [7]byte

func (t TThostFtdcZipCodeType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcZipCodeType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcTelephoneType是一个联系电话类型
type TThostFtdcTelephoneType [41]byte

func (t TThostFtdcTelephoneType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcTelephoneType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcFaxType是一个传真类型
type TThostFtdcFaxType [41]byte

func (t TThostFtdcFaxType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcFaxType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcMobileType是一个手机类型
type TThostFtdcMobileType [41]byte

func (t TThostFtdcMobileType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcMobileType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcEMailType是一个电子邮件类型
type TThostFtdcEMailType [41]byte

func (t TThostFtdcEMailType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcEMailType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcMemoType是一个备注类型
type TThostFtdcMemoType [161]byte

func (t TThostFtdcMemoType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcMemoType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcCompanyCodeType是一个企业代码类型
type TThostFtdcCompanyCodeType [51]byte

func (t TThostFtdcCompanyCodeType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcCompanyCodeType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcWebsiteType是一个网站地址类型
type TThostFtdcWebsiteType [51]byte

func (t TThostFtdcWebsiteType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcWebsiteType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcTaxNoType是一个税务登记号类型
type TThostFtdcTaxNoType [31]byte

func (t TThostFtdcTaxNoType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcTaxNoType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcBatchStatusType是一个处理状态类型
type TThostFtdcBatchStatusType uint8

//go:generate stringer -type TThostFtdcBatchStatusType -linecomment
const (
	THOST_FTDC_BS_NoUpload TThostFtdcBatchStatusType = '1' // 未上传
	THOST_FTDC_BS_Uploaded TThostFtdcBatchStatusType = '2' // 已上传
	THOST_FTDC_BS_Failed   TThostFtdcBatchStatusType = '3' // 审核失败
)

// TFtdcPropertyIDType是一个属性代码类型
type TThostFtdcPropertyIDType [33]byte

func (t TThostFtdcPropertyIDType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcPropertyIDType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcPropertyNameType是一个属性名称类型
type TThostFtdcPropertyNameType [65]byte

func (t TThostFtdcPropertyNameType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcPropertyNameType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcLicenseNoType是一个营业执照号类型
type TThostFtdcLicenseNoType [51]byte

func (t TThostFtdcLicenseNoType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcLicenseNoType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcAgentIDType是一个经纪人代码类型
type TThostFtdcAgentIDType [13]byte

func (t TThostFtdcAgentIDType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcAgentIDType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcAgentNameType是一个经纪人名称类型
type TThostFtdcAgentNameType [41]byte

func (t TThostFtdcAgentNameType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcAgentNameType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcAgentGroupIDType是一个经纪人组代码类型
type TThostFtdcAgentGroupIDType [13]byte

func (t TThostFtdcAgentGroupIDType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcAgentGroupIDType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcAgentGroupNameType是一个经纪人组名称类型
type TThostFtdcAgentGroupNameType [41]byte

func (t TThostFtdcAgentGroupNameType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcAgentGroupNameType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

type TThostFtdcReturnStyleType uint8

// TFtdcReturnPatternType是一个返还模式类型
type TThostFtdcReturnPatternType uint8

//go:generate stringer -type TThostFtdcReturnPatternType -linecomment
const (
	THOST_FTDC_RP_ByVolume    TThostFtdcReturnPatternType = '1' // 按成交手数
	THOST_FTDC_RP_ByFeeOnHand TThostFtdcReturnPatternType = '2' // 按留存手续费
)

// TFtdcReturnLevelType是一个返还级别类型
type TThostFtdcReturnLevelType uint8

//go:generate stringer -type TThostFtdcReturnLevelType -linecomment
const (
	THOST_FTDC_RL_Level1 TThostFtdcReturnLevelType = '1' // 级别1
	THOST_FTDC_RL_Level2 TThostFtdcReturnLevelType = '2' // 级别2
	THOST_FTDC_RL_Level3 TThostFtdcReturnLevelType = '3' // 级别3
	THOST_FTDC_RL_Level4 TThostFtdcReturnLevelType = '4' // 级别4
	THOST_FTDC_RL_Level5 TThostFtdcReturnLevelType = '5' // 级别5
	THOST_FTDC_RL_Level6 TThostFtdcReturnLevelType = '6' // 级别6
	THOST_FTDC_RL_Level7 TThostFtdcReturnLevelType = '7' // 级别7
	THOST_FTDC_RL_Level8 TThostFtdcReturnLevelType = '8' // 级别8
	THOST_FTDC_RL_Level9 TThostFtdcReturnLevelType = '9' // 级别9
)

// TFtdcReturnStandardType是一个返还标准类型
type TThostFtdcReturnStandardType uint8

//go:generate stringer -type TThostFtdcReturnStandardType -linecomment
const (
	THOST_FTDC_RSD_ByPeriod   TThostFtdcReturnStandardType = '1' // 分阶段返还
	THOST_FTDC_RSD_ByStandard TThostFtdcReturnStandardType = '2' // 按某一标准
)

// TFtdcMortgageTypeType是一个质押类型类型
type TThostFtdcMortgageTypeType uint8

//go:generate stringer -type TThostFtdcMortgageTypeType -linecomment
const (
	THOST_FTDC_MT_Out TThostFtdcMortgageTypeType = '0' // 质出
	THOST_FTDC_MT_In  TThostFtdcMortgageTypeType = '1' // 质入
)

// TFtdcInvestorSettlementParamIDType是一个投资者结算参数代码类型
type TThostFtdcInvestorSettlementParamIDType uint8

//go:generate stringer -type TThostFtdcInvestorSettlementParamIDType -linecomment
const (
	THOST_FTDC_ISPI_MortgageRatio TThostFtdcInvestorSettlementParamIDType = '4' // 质押比例
	THOST_FTDC_ISPI_MarginWay     TThostFtdcInvestorSettlementParamIDType = '5' // 保证金算法
	THOST_FTDC_ISPI_BillDeposit   TThostFtdcInvestorSettlementParamIDType = '9' // 结算单结存是否包含质押
)

// TFtdcExchangeSettlementParamIDType是一个交易所结算参数代码类型
type TThostFtdcExchangeSettlementParamIDType uint8

//go:generate stringer -type TThostFtdcExchangeSettlementParamIDType -linecomment
const (
	THOST_FTDC_ESPI_MortgageRatio      TThostFtdcExchangeSettlementParamIDType = '1' // 质押比例
	THOST_FTDC_ESPI_OtherFundItem      TThostFtdcExchangeSettlementParamIDType = '2' // 分项资金导入项
	THOST_FTDC_ESPI_OtherFundImport    TThostFtdcExchangeSettlementParamIDType = '3' // 分项资金入交易所出入金
	THOST_FTDC_ESPI_CFFEXMinPrepa      TThostFtdcExchangeSettlementParamIDType = '6' // 中金所开户最低可用金额
	THOST_FTDC_ESPI_CZCESettlementType TThostFtdcExchangeSettlementParamIDType = '7' // 郑商所结算方式
	THOST_FTDC_ESPI_ExchDelivFeeMode   TThostFtdcExchangeSettlementParamIDType = '9' // 交易所交割手续费收取方式
	THOST_FTDC_ESPI_DelivFeeMode       TThostFtdcExchangeSettlementParamIDType = '0' // 投资者交割手续费收取方式
	THOST_FTDC_ESPI_CZCEComMarginType  TThostFtdcExchangeSettlementParamIDType = 'A' // 郑商所组合持仓保证金收取方式
	THOST_FTDC_ESPI_DceComMarginType   TThostFtdcExchangeSettlementParamIDType = 'B' // 大商所套利保证金是否优惠
	THOST_FTDC_ESPI_OptOutDisCountRate TThostFtdcExchangeSettlementParamIDType = 'a' // 虚值期权保证金优惠比率
	THOST_FTDC_ESPI_OptMiniGuarantee   TThostFtdcExchangeSettlementParamIDType = 'b' // 最低保障系数
)

// TFtdcSystemParamIDType是一个系统参数代码类型
type TThostFtdcSystemParamIDType uint8

//go:generate stringer -type TThostFtdcSystemParamIDType -linecomment
const (
	THOST_FTDC_SPI_InvestorIDMinLength     TThostFtdcSystemParamIDType = '1' // 投资者代码最小长度
	THOST_FTDC_SPI_AccountIDMinLength      TThostFtdcSystemParamIDType = '2' // 投资者帐号代码最小长度
	THOST_FTDC_SPI_UserRightLogon          TThostFtdcSystemParamIDType = '3' // 投资者开户默认登录权限
	THOST_FTDC_SPI_SettlementBillTrade     TThostFtdcSystemParamIDType = '4' // 投资者交易结算单成交汇总方式
	THOST_FTDC_SPI_TradingCode             TThostFtdcSystemParamIDType = '5' // 统一开户更新交易编码方式
	THOST_FTDC_SPI_CheckFund               TThostFtdcSystemParamIDType = '6' // 结算是否判断存在未复核的出入金和分项资金
	THOST_FTDC_SPI_CommModelRight          TThostFtdcSystemParamIDType = '7' // 是否启用手续费模板数据权限
	THOST_FTDC_SPI_MarginModelRight        TThostFtdcSystemParamIDType = '9' // 是否启用保证金率模板数据权限
	THOST_FTDC_SPI_IsStandardActive        TThostFtdcSystemParamIDType = '8' // 是否规范用户才能激活
	THOST_FTDC_SPI_UploadSettlementFile    TThostFtdcSystemParamIDType = 'U' // 上传的交易所结算文件路径
	THOST_FTDC_SPI_DownloadCSRCFile        TThostFtdcSystemParamIDType = 'D' // 上报保证金监控中心文件路径
	THOST_FTDC_SPI_SettlementBillFile      TThostFtdcSystemParamIDType = 'S' // 生成的结算单文件路径
	THOST_FTDC_SPI_CSRCOthersFile          TThostFtdcSystemParamIDType = 'C' // 证监会文件标识
	THOST_FTDC_SPI_InvestorPhoto           TThostFtdcSystemParamIDType = 'P' // 投资者照片路径
	THOST_FTDC_SPI_CSRCData                TThostFtdcSystemParamIDType = 'R' // 全结经纪公司上传文件路径
	THOST_FTDC_SPI_InvestorPwdModel        TThostFtdcSystemParamIDType = 'I' // 开户密码录入方式
	THOST_FTDC_SPI_CFFEXInvestorSettleFile TThostFtdcSystemParamIDType = 'F' // 投资者中金所结算文件下载路径
	THOST_FTDC_SPI_InvestorIDType          TThostFtdcSystemParamIDType = 'a' // 投资者代码编码方式
	THOST_FTDC_SPI_FreezeMaxReMain         TThostFtdcSystemParamIDType = 'r' // 休眠户最高权益
	THOST_FTDC_SPI_IsSync                  TThostFtdcSystemParamIDType = 'A' // 手续费相关操作实时上场开关
	THOST_FTDC_SPI_RelieveOpenLimit        TThostFtdcSystemParamIDType = 'O' // 解除开仓权限限制
	THOST_FTDC_SPI_IsStandardFreeze        TThostFtdcSystemParamIDType = 'X' // 是否规范用户才能休眠
	THOST_FTDC_SPI_CZCENormalProductHedge  TThostFtdcSystemParamIDType = 'B' // 郑商所是否开放所有品种套保交易
)

// TFtdcTradeParamIDType是一个交易系统参数代码类型
type TThostFtdcTradeParamIDType uint8

//go:generate stringer -type TThostFtdcTradeParamIDType -linecomment
const (
	THOST_FTDC_TPID_EncryptionStandard      TThostFtdcTradeParamIDType = 'E' // 系统加密算法
	THOST_FTDC_TPID_RiskMode                TThostFtdcTradeParamIDType = 'R' // 系统风险算法
	THOST_FTDC_TPID_RiskModeGlobal          TThostFtdcTradeParamIDType = 'G' // 系统风险算法是否全局 0-否 1-是
	THOST_FTDC_TPID_modeEncode              TThostFtdcTradeParamIDType = 'P' // 密码加密算法
	THOST_FTDC_TPID_tickMode                TThostFtdcTradeParamIDType = 'T' // 价格小数位数参数
	THOST_FTDC_TPID_SingleUserSessionMaxNum TThostFtdcTradeParamIDType = 'S' // 用户最大会话数
	THOST_FTDC_TPID_LoginFailMaxNum         TThostFtdcTradeParamIDType = 'L' // 最大连续登录失败数
	THOST_FTDC_TPID_IsAuthForce             TThostFtdcTradeParamIDType = 'A' // 是否强制认证
	THOST_FTDC_TPID_IsPosiFreeze            TThostFtdcTradeParamIDType = 'F' // 是否冻结证券持仓
	THOST_FTDC_TPID_IsPosiLimit             TThostFtdcTradeParamIDType = 'M' // 是否限仓
	THOST_FTDC_TPID_ForQuoteTimeInterval    TThostFtdcTradeParamIDType = 'Q' // 郑商所询价时间间隔
	THOST_FTDC_TPID_IsFuturePosiLimit       TThostFtdcTradeParamIDType = 'B' // 是否期货限仓
	THOST_FTDC_TPID_IsFutureOrderFreq       TThostFtdcTradeParamIDType = 'C' // 是否期货下单频率限制
	THOST_FTDC_TPID_IsExecOrderProfit       TThostFtdcTradeParamIDType = 'H' // 行权冻结是否计算盈利
)

// TFtdcSettlementParamValueType是一个参数代码值类型
type TThostFtdcSettlementParamValueType [256]byte

func (t TThostFtdcSettlementParamValueType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcSettlementParamValueType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcCounterIDType是一个计数器代码类型
type TThostFtdcCounterIDType [33]byte

func (t TThostFtdcCounterIDType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcCounterIDType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcInvestorGroupNameType是一个投资者分组名称类型
type TThostFtdcInvestorGroupNameType [41]byte

func (t TThostFtdcInvestorGroupNameType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcInvestorGroupNameType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcBrandCodeType是一个牌号类型
type TThostFtdcBrandCodeType [257]byte

func (t TThostFtdcBrandCodeType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcBrandCodeType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcWarehouseType是一个仓库类型
type TThostFtdcWarehouseType [257]byte

func (t TThostFtdcWarehouseType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcWarehouseType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcProductDateType是一个产期类型
type TThostFtdcProductDateType [41]byte

func (t TThostFtdcProductDateType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcProductDateType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcGradeType是一个等级类型
type TThostFtdcGradeType [41]byte

func (t TThostFtdcGradeType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcGradeType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcClassifyType是一个类别类型
type TThostFtdcClassifyType [41]byte

func (t TThostFtdcClassifyType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcClassifyType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcPositionType是一个货位类型
type TThostFtdcPositionType [41]byte

func (t TThostFtdcPositionType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcPositionType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcYieldlyType是一个产地类型
type TThostFtdcYieldlyType [41]byte

func (t TThostFtdcYieldlyType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcYieldlyType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcWeightType是一个公定重量类型
type TThostFtdcWeightType [41]byte

func (t TThostFtdcWeightType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcWeightType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcSubEntryFundNoType是一个分项资金流水号类型
type TThostFtdcSubEntryFundNoType int32

// TFtdcFileIDType是一个文件标识类型
type TThostFtdcFileIDType uint8

//go:generate stringer -type TThostFtdcFileIDType -linecomment
const (
	THOST_FTDC_FI_SettlementFund            TThostFtdcFileIDType = 'F' // 资金数据
	THOST_FTDC_FI_Trade                     TThostFtdcFileIDType = 'T' // 成交数据
	THOST_FTDC_FI_InvestorPosition          TThostFtdcFileIDType = 'P' // 投资者持仓数据
	THOST_FTDC_FI_SubEntryFund              TThostFtdcFileIDType = 'O' // 投资者分项资金数据
	THOST_FTDC_FI_CZCECombinationPos        TThostFtdcFileIDType = 'C' // 组合持仓数据
	THOST_FTDC_FI_CSRCData                  TThostFtdcFileIDType = 'R' // 上报保证金监控中心数据
	THOST_FTDC_FI_CZCEClose                 TThostFtdcFileIDType = 'L' // 郑商所平仓了结数据
	THOST_FTDC_FI_CZCENoClose               TThostFtdcFileIDType = 'N' // 郑商所非平仓了结数据
	THOST_FTDC_FI_PositionDtl               TThostFtdcFileIDType = 'D' // 持仓明细数据
	THOST_FTDC_FI_OptionStrike              TThostFtdcFileIDType = 'S' // 期权执行文件
	THOST_FTDC_FI_SettlementPriceComparison TThostFtdcFileIDType = 'M' // 结算价比对文件
	THOST_FTDC_FI_NonTradePosChange         TThostFtdcFileIDType = 'B' // 上期所非持仓变动明细
)

// TFtdcFileNameType是一个文件名称类型
type TThostFtdcFileNameType [257]byte

func (t TThostFtdcFileNameType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcFileNameType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcFileTypeType是一个文件上传类型类型
type TThostFtdcFileTypeType uint8

//go:generate stringer -type TThostFtdcFileTypeType -linecomment
const (
	THOST_FTDC_FUT_Settlement TThostFtdcFileTypeType = '0' // 结算
	THOST_FTDC_FUT_Check      TThostFtdcFileTypeType = '1' // 核对
)

// TFtdcFileFormatType是一个文件格式类型
type TThostFtdcFileFormatType uint8

//go:generate stringer -type TThostFtdcFileFormatType -linecomment
const (
	THOST_FTDC_FFT_Txt TThostFtdcFileFormatType = '0' // 文本文件(.txt)
	THOST_FTDC_FFT_Zip TThostFtdcFileFormatType = '1' // 压缩文件(.zip)
	THOST_FTDC_FFT_DBF TThostFtdcFileFormatType = '2' // DBF文件(.dbf)
)

// TFtdcFileUploadStatusType是一个文件状态类型
type TThostFtdcFileUploadStatusType uint8

//go:generate stringer -type TThostFtdcFileUploadStatusType -linecomment
const (
	THOST_FTDC_FUS_SucceedUpload   TThostFtdcFileUploadStatusType = '1' // 上传成功
	THOST_FTDC_FUS_FailedUpload    TThostFtdcFileUploadStatusType = '2' // 上传失败
	THOST_FTDC_FUS_SucceedLoad     TThostFtdcFileUploadStatusType = '3' // 导入成功
	THOST_FTDC_FUS_PartSucceedLoad TThostFtdcFileUploadStatusType = '4' // 导入部分成功
	THOST_FTDC_FUS_FailedLoad      TThostFtdcFileUploadStatusType = '5' // 导入失败
)

// TFtdcTransferDirectionType是一个移仓方向类型
type TThostFtdcTransferDirectionType uint8

//go:generate stringer -type TThostFtdcTransferDirectionType -linecomment
const (
	THOST_FTDC_TD_Out TThostFtdcTransferDirectionType = '0' // 移出
	THOST_FTDC_TD_In  TThostFtdcTransferDirectionType = '1' // 移入
)

// TFtdcUploadModeType是一个上传文件类型类型
type TThostFtdcUploadModeType [21]byte

func (t TThostFtdcUploadModeType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcUploadModeType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcAccountIDType是一个投资者帐号类型
type TThostFtdcAccountIDType [13]byte

func (t TThostFtdcAccountIDType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcAccountIDType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcBankFlagType是一个银行统一标识类型类型
type TThostFtdcBankFlagType [4]byte

func (t TThostFtdcBankFlagType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcBankFlagType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcBankAccountType是一个银行账户类型
type TThostFtdcBankAccountType [41]byte

func (t TThostFtdcBankAccountType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcBankAccountType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcOpenNameType是一个银行账户的开户人名称类型
type TThostFtdcOpenNameType [61]byte

func (t TThostFtdcOpenNameType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcOpenNameType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcOpenBankType是一个银行账户的开户行类型
type TThostFtdcOpenBankType [101]byte

func (t TThostFtdcOpenBankType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcOpenBankType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcBankNameType是一个银行名称类型
type TThostFtdcBankNameType [101]byte

func (t TThostFtdcBankNameType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcBankNameType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcPublishPathType是一个发布路径类型
type TThostFtdcPublishPathType [257]byte

func (t TThostFtdcPublishPathType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcPublishPathType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcOperatorIDType是一个操作员代码类型
type TThostFtdcOperatorIDType [65]byte

func (t TThostFtdcOperatorIDType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcOperatorIDType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcMonthCountType是一个月份数量类型
type TThostFtdcMonthCountType int32

// TFtdcAdvanceMonthArrayType是一个月份提前数组类型
type TThostFtdcAdvanceMonthArrayType [13]byte

func (t TThostFtdcAdvanceMonthArrayType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcAdvanceMonthArrayType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcDateExprType是一个日期表达式类型
type TThostFtdcDateExprType [1025]byte

func (t TThostFtdcDateExprType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcDateExprType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcInstrumentIDExprType是一个合约代码表达式类型
type TThostFtdcInstrumentIDExprType [41]byte

func (t TThostFtdcInstrumentIDExprType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcInstrumentIDExprType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcInstrumentNameExprType是一个合约名称表达式类型
type TThostFtdcInstrumentNameExprType [41]byte

func (t TThostFtdcInstrumentNameExprType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcInstrumentNameExprType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcSpecialCreateRuleType是一个特殊的创建规则类型
type TThostFtdcSpecialCreateRuleType uint8

//go:generate stringer -type TThostFtdcSpecialCreateRuleType -linecomment
const (
	THOST_FTDC_SC_NoSpecialRule    TThostFtdcSpecialCreateRuleType = '0' // 没有特殊创建规则
	THOST_FTDC_SC_NoSpringFestival TThostFtdcSpecialCreateRuleType = '1' // 不包含春节
)

// TFtdcBasisPriceTypeType是一个挂牌基准价类型类型
type TThostFtdcBasisPriceTypeType uint8

//go:generate stringer -type TThostFtdcBasisPriceTypeType -linecomment
const (
	THOST_FTDC_IPT_LastSettlement TThostFtdcBasisPriceTypeType = '1' // 上一合约结算价
	THOST_FTDC_IPT_LaseClose      TThostFtdcBasisPriceTypeType = '2' // 上一合约收盘价
)

// TFtdcProductLifePhaseType是一个产品生命周期状态类型
type TThostFtdcProductLifePhaseType uint8

//go:generate stringer -type TThostFtdcProductLifePhaseType -linecomment
const (
	THOST_FTDC_PLP_Active    TThostFtdcProductLifePhaseType = '1' // 活跃
	THOST_FTDC_PLP_NonActive TThostFtdcProductLifePhaseType = '2' // 不活跃
	THOST_FTDC_PLP_Canceled  TThostFtdcProductLifePhaseType = '3' // 注销
)

// TFtdcDeliveryModeType是一个交割方式类型
type TThostFtdcDeliveryModeType uint8

//go:generate stringer -type TThostFtdcDeliveryModeType -linecomment
const (
	THOST_FTDC_DM_CashDeliv      TThostFtdcDeliveryModeType = '1' // 现金交割
	THOST_FTDC_DM_CommodityDeliv TThostFtdcDeliveryModeType = '2' // 实物交割
)

// TFtdcLogLevelType是一个日志级别类型
type TThostFtdcLogLevelType [33]byte

func (t TThostFtdcLogLevelType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcLogLevelType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcProcessNameType是一个存储过程名称类型
type TThostFtdcProcessNameType [257]byte

func (t TThostFtdcProcessNameType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcProcessNameType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcOperationMemoType是一个操作摘要类型
type TThostFtdcOperationMemoType [1025]byte

func (t TThostFtdcOperationMemoType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcOperationMemoType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcFundIOTypeType是一个出入金类型类型
type TThostFtdcFundIOTypeType uint8

//go:generate stringer -type TThostFtdcFundIOTypeType -linecomment
const (
	THOST_FTDC_FIOT_FundIO       TThostFtdcFundIOTypeType = '1' // 出入金
	THOST_FTDC_FIOT_Transfer     TThostFtdcFundIOTypeType = '2' // 银期转帐
	THOST_FTDC_FIOT_SwapCurrency TThostFtdcFundIOTypeType = '3' // 银期换汇
)

// TFtdcFundTypeType是一个资金类型类型
type TThostFtdcFundTypeType uint8

//go:generate stringer -type TThostFtdcFundTypeType -linecomment
const (
	THOST_FTDC_FT_Deposite      TThostFtdcFundTypeType = '1' // 银行存款
	THOST_FTDC_FT_ItemFund      TThostFtdcFundTypeType = '2' // 分项资金
	THOST_FTDC_FT_Company       TThostFtdcFundTypeType = '3' // 公司调整
	THOST_FTDC_FT_InnerTransfer TThostFtdcFundTypeType = '4' // 资金内转
)

// TFtdcFundDirectionType是一个出入金方向类型
type TThostFtdcFundDirectionType uint8

//go:generate stringer -type TThostFtdcFundDirectionType -linecomment
const (
	THOST_FTDC_FD_In  TThostFtdcFundDirectionType = '1' // 入金
	THOST_FTDC_FD_Out TThostFtdcFundDirectionType = '2' // 出金
)

// TFtdcFundStatusType是一个资金状态类型
type TThostFtdcFundStatusType uint8

//go:generate stringer -type TThostFtdcFundStatusType -linecomment
const (
	THOST_FTDC_FS_Record TThostFtdcFundStatusType = '1' // 已录入
	THOST_FTDC_FS_Check  TThostFtdcFundStatusType = '2' // 已复核
	THOST_FTDC_FS_Charge TThostFtdcFundStatusType = '3' // 已冲销
)

// TFtdcBillNoType是一个票据号类型
type TThostFtdcBillNoType [15]byte

func (t TThostFtdcBillNoType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcBillNoType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcBillNameType是一个票据名称类型
type TThostFtdcBillNameType [33]byte

func (t TThostFtdcBillNameType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcBillNameType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcPublishStatusType是一个发布状态类型
type TThostFtdcPublishStatusType uint8

//go:generate stringer -type TThostFtdcPublishStatusType -linecomment
const (
	THOST_FTDC_PS_None       TThostFtdcPublishStatusType = '1' // 未发布
	THOST_FTDC_PS_Publishing TThostFtdcPublishStatusType = '2' // 正在发布
	THOST_FTDC_PS_Published  TThostFtdcPublishStatusType = '3' // 已发布
)

// TFtdcEnumValueIDType是一个枚举值代码类型
type TThostFtdcEnumValueIDType [65]byte

func (t TThostFtdcEnumValueIDType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcEnumValueIDType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcEnumValueTypeType是一个枚举值类型类型
type TThostFtdcEnumValueTypeType [33]byte

func (t TThostFtdcEnumValueTypeType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcEnumValueTypeType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcEnumValueLabelType是一个枚举值名称类型
type TThostFtdcEnumValueLabelType [65]byte

func (t TThostFtdcEnumValueLabelType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcEnumValueLabelType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcEnumValueResultType是一个枚举值结果类型
type TThostFtdcEnumValueResultType [33]byte

func (t TThostFtdcEnumValueResultType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcEnumValueResultType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcSystemStatusType是一个系统状态类型
type TThostFtdcSystemStatusType uint8

//go:generate stringer -type TThostFtdcSystemStatusType -linecomment
const (
	THOST_FTDC_ES_NonActive   TThostFtdcSystemStatusType = '1' // 不活跃
	THOST_FTDC_ES_Startup     TThostFtdcSystemStatusType = '2' // 启动
	THOST_FTDC_ES_Initialize  TThostFtdcSystemStatusType = '3' // 交易开始初始化
	THOST_FTDC_ES_Initialized TThostFtdcSystemStatusType = '4' // 交易完成初始化
	THOST_FTDC_ES_Close       TThostFtdcSystemStatusType = '5' // 收市开始
	THOST_FTDC_ES_Closed      TThostFtdcSystemStatusType = '6' // 收市完成
	THOST_FTDC_ES_Settlement  TThostFtdcSystemStatusType = '7' // 结算
)

// TFtdcSettlementStatusType是一个结算状态类型
type TThostFtdcSettlementStatusType uint8

//go:generate stringer -type TThostFtdcSettlementStatusType -linecomment
const (
	THOST_FTDC_STS_Initialize    TThostFtdcSettlementStatusType = '0' // 初始
	THOST_FTDC_STS_Settlementing TThostFtdcSettlementStatusType = '1' // 结算中
	THOST_FTDC_STS_Settlemented  TThostFtdcSettlementStatusType = '2' // 已结算
	THOST_FTDC_STS_Finished      TThostFtdcSettlementStatusType = '3' // 结算完成
)

// TFtdcRangeIntTypeType是一个限定值类型类型
type TThostFtdcRangeIntTypeType [33]byte

func (t TThostFtdcRangeIntTypeType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcRangeIntTypeType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcRangeIntFromType是一个限定值下限类型
type TThostFtdcRangeIntFromType [33]byte

func (t TThostFtdcRangeIntFromType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcRangeIntFromType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcRangeIntToType是一个限定值上限类型
type TThostFtdcRangeIntToType [33]byte

func (t TThostFtdcRangeIntToType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcRangeIntToType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcFunctionIDType是一个功能代码类型
type TThostFtdcFunctionIDType [25]byte

func (t TThostFtdcFunctionIDType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcFunctionIDType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcFunctionValueCodeType是一个功能编码类型
type TThostFtdcFunctionValueCodeType [257]byte

func (t TThostFtdcFunctionValueCodeType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcFunctionValueCodeType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcFunctionNameType是一个功能名称类型
type TThostFtdcFunctionNameType [65]byte

func (t TThostFtdcFunctionNameType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcFunctionNameType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcRoleIDType是一个角色编号类型
type TThostFtdcRoleIDType [11]byte

func (t TThostFtdcRoleIDType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcRoleIDType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcRoleNameType是一个角色名称类型
type TThostFtdcRoleNameType [41]byte

func (t TThostFtdcRoleNameType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcRoleNameType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcDescriptionType是一个描述类型
type TThostFtdcDescriptionType [401]byte

func (t TThostFtdcDescriptionType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcDescriptionType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcCombineIDType是一个组合编号类型
type TThostFtdcCombineIDType [25]byte

func (t TThostFtdcCombineIDType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcCombineIDType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcCombineTypeType是一个组合类型类型
type TThostFtdcCombineTypeType [25]byte

func (t TThostFtdcCombineTypeType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcCombineTypeType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcInvestorTypeType是一个投资者类型类型
type TThostFtdcInvestorTypeType uint8

//go:generate stringer -type TThostFtdcInvestorTypeType -linecomment
const (
	THOST_FTDC_CT_Person       TThostFtdcInvestorTypeType = '0' // 自然人
	THOST_FTDC_CT_Company      TThostFtdcInvestorTypeType = '1' // 法人
	THOST_FTDC_CT_Fund         TThostFtdcInvestorTypeType = '2' // 投资基金
	THOST_FTDC_CT_SpecialOrgan TThostFtdcInvestorTypeType = '3' // 特殊法人
	THOST_FTDC_CT_Asset        TThostFtdcInvestorTypeType = '4' // 资管户
)

// TFtdcBrokerTypeType是一个经纪公司类型类型
type TThostFtdcBrokerTypeType uint8

//go:generate stringer -type TThostFtdcBrokerTypeType -linecomment
const (
	THOST_FTDC_BT_Trade       TThostFtdcBrokerTypeType = '0' // 交易会员
	THOST_FTDC_BT_TradeSettle TThostFtdcBrokerTypeType = '1' // 交易结算会员
	THOST_FTDC_BT_Request     TThostFtdcBrokerTypeType = '1' // 请求
	THOST_FTDC_BT_Response    TThostFtdcBrokerTypeType = '2' // 应答
	THOST_FTDC_BT_Notice      TThostFtdcBrokerTypeType = '3' // 通知
	THOST_FTDC_BT_Profit      TThostFtdcBrokerTypeType = '0' // 盈利
	THOST_FTDC_BT_Loss        TThostFtdcBrokerTypeType = '1' // 亏损
	THOST_FTDC_BT_Other       TThostFtdcBrokerTypeType = 'Z' // 其他
)

// TFtdcRiskLevelType是一个风险等级类型
type TThostFtdcRiskLevelType uint8

//go:generate stringer -type TThostFtdcRiskLevelType -linecomment
const (
	THOST_FTDC_FAS_Low     TThostFtdcRiskLevelType = '1' // 低风险客户
	THOST_FTDC_FAS_Normal  TThostFtdcRiskLevelType = '2' // 普通客户
	THOST_FTDC_FAS_Focus   TThostFtdcRiskLevelType = '3' // 关注客户
	THOST_FTDC_FAS_Risk    TThostFtdcRiskLevelType = '4' // 风险客户
	THOST_FTDC_FAS_ByTrade TThostFtdcRiskLevelType = '1' // 按交易收取
	THOST_FTDC_FAS_ByDeliv TThostFtdcRiskLevelType = '2' // 按交割收取
	THOST_FTDC_FAS_None    TThostFtdcRiskLevelType = '3' // 不收
	THOST_FTDC_FAS_FixFee  TThostFtdcRiskLevelType = '4' // 按指定手续费收取
)

type TThostFtdcFeeAcceptStyleType uint8

// TFtdcPasswordTypeType是一个密码类型类型
type TThostFtdcPasswordTypeType uint8

//go:generate stringer -type TThostFtdcPasswordTypeType -linecomment
const (
	THOST_FTDC_PWDT_Trade   TThostFtdcPasswordTypeType = '1' // 交易密码
	THOST_FTDC_PWDT_Account TThostFtdcPasswordTypeType = '2' // 资金密码
)

// TFtdcAlgorithmType是一个盈亏算法类型
type TThostFtdcAlgorithmType uint8

//go:generate stringer -type TThostFtdcAlgorithmType -linecomment
const (
	THOST_FTDC_AG_All      TThostFtdcAlgorithmType = '1' // 浮盈浮亏都计算
	THOST_FTDC_AG_OnlyLost TThostFtdcAlgorithmType = '2' // 浮盈不计，浮亏计
	THOST_FTDC_AG_OnlyGain TThostFtdcAlgorithmType = '3' // 浮盈计，浮亏不计
	THOST_FTDC_AG_None     TThostFtdcAlgorithmType = '4' // 浮盈浮亏都不计算
)

// TFtdcIncludeCloseProfitType是一个是否包含平仓盈利类型
type TThostFtdcIncludeCloseProfitType uint8

//go:generate stringer -type TThostFtdcIncludeCloseProfitType -linecomment
const (
	THOST_FTDC_ICP_Include    TThostFtdcIncludeCloseProfitType = '0' // 包含平仓盈利
	THOST_FTDC_ICP_NotInclude TThostFtdcIncludeCloseProfitType = '2' // 不包含平仓盈利
)

// TFtdcAllWithoutTradeType是一个是否受可提比例限制类型
type TThostFtdcAllWithoutTradeType uint8

//go:generate stringer -type TThostFtdcAllWithoutTradeType -linecomment
const (
	THOST_FTDC_AWT_Enable       TThostFtdcAllWithoutTradeType = '0' // 无仓无成交不受可提比例限制
	THOST_FTDC_AWT_Disable      TThostFtdcAllWithoutTradeType = '2' // 受可提比例限制
	THOST_FTDC_AWT_NoHoldEnable TThostFtdcAllWithoutTradeType = '3' // 无仓不受可提比例限制
)

// TFtdcCommentType是一个盈亏算法说明类型
type TThostFtdcCommentType [31]byte

func (t TThostFtdcCommentType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcCommentType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcVersionType是一个版本号类型
type TThostFtdcVersionType [4]byte

func (t TThostFtdcVersionType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcVersionType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcTradeCodeType是一个交易代码类型
type TThostFtdcTradeCodeType [7]byte

func (t TThostFtdcTradeCodeType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcTradeCodeType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcTradeDateType是一个交易日期类型
type TThostFtdcTradeDateType [9]byte

func (t TThostFtdcTradeDateType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcTradeDateType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcTradeTimeType是一个交易时间类型
type TThostFtdcTradeTimeType [9]byte

func (t TThostFtdcTradeTimeType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcTradeTimeType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcTradeSerialType是一个发起方流水号类型
type TThostFtdcTradeSerialType [9]byte

func (t TThostFtdcTradeSerialType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcTradeSerialType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcTradeSerialNoType是一个发起方流水号类型
type TThostFtdcTradeSerialNoType int32

// TFtdcFutureIDType是一个期货公司代码类型
type TThostFtdcFutureIDType [11]byte

func (t TThostFtdcFutureIDType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcFutureIDType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcBankIDType是一个银行代码类型
type TThostFtdcBankIDType [4]byte

func (t TThostFtdcBankIDType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcBankIDType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcBankBrchIDType是一个银行分中心代码类型
type TThostFtdcBankBrchIDType [5]byte

func (t TThostFtdcBankBrchIDType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcBankBrchIDType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcBankBranchIDType是一个分中心代码类型
type TThostFtdcBankBranchIDType [11]byte

func (t TThostFtdcBankBranchIDType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcBankBranchIDType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcOperNoType是一个交易柜员类型
type TThostFtdcOperNoType [17]byte

func (t TThostFtdcOperNoType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcOperNoType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcDeviceIDType是一个渠道标志类型
type TThostFtdcDeviceIDType [3]byte

func (t TThostFtdcDeviceIDType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcDeviceIDType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcRecordNumType是一个记录数类型
type TThostFtdcRecordNumType [7]byte

func (t TThostFtdcRecordNumType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcRecordNumType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcFutureAccountType是一个期货资金账号类型
type TThostFtdcFutureAccountType [22]byte

func (t TThostFtdcFutureAccountType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcFutureAccountType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcFuturePwdFlagType是一个资金密码核对标志类型
type TThostFtdcFuturePwdFlagType uint8

//go:generate stringer -type TThostFtdcFuturePwdFlagType -linecomment
const (
	THOST_FTDC_FPWD_UnCheck TThostFtdcFuturePwdFlagType = '0' // 不核对
	THOST_FTDC_FPWD_Check   TThostFtdcFuturePwdFlagType = '1' // 核对
)

// TFtdcTransferTypeType是一个银期转账类型类型
type TThostFtdcTransferTypeType uint8

//go:generate stringer -type TThostFtdcTransferTypeType -linecomment
const (
	THOST_FTDC_TT_BankToFuture TThostFtdcTransferTypeType = '0' // 银行转期货
	THOST_FTDC_TT_FutureToBank TThostFtdcTransferTypeType = '1' // 期货转银行
)

// TFtdcFutureAccPwdType是一个期货资金密码类型
type TThostFtdcFutureAccPwdType [17]byte

func (t TThostFtdcFutureAccPwdType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcFutureAccPwdType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcCurrencyCodeType是一个币种类型
type TThostFtdcCurrencyCodeType [4]byte

func (t TThostFtdcCurrencyCodeType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcCurrencyCodeType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcRetCodeType是一个响应代码类型
type TThostFtdcRetCodeType [5]byte

func (t TThostFtdcRetCodeType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcRetCodeType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcRetInfoType是一个响应信息类型
type TThostFtdcRetInfoType [129]byte

func (t TThostFtdcRetInfoType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcRetInfoType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcTradeAmtType是一个银行总余额类型
type TThostFtdcTradeAmtType [20]byte

func (t TThostFtdcTradeAmtType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcTradeAmtType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcUseAmtType是一个银行可用余额类型
type TThostFtdcUseAmtType [20]byte

func (t TThostFtdcUseAmtType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcUseAmtType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcFetchAmtType是一个银行可取余额类型
type TThostFtdcFetchAmtType [20]byte

func (t TThostFtdcFetchAmtType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcFetchAmtType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcTransferValidFlagType是一个转账有效标志类型
type TThostFtdcTransferValidFlagType uint8

//go:generate stringer -type TThostFtdcTransferValidFlagType -linecomment
const (
	THOST_FTDC_TVF_Invalid TThostFtdcTransferValidFlagType = '0' // 无效或失败
	THOST_FTDC_TVF_Valid   TThostFtdcTransferValidFlagType = '1' // 有效
	THOST_FTDC_TVF_Reverse TThostFtdcTransferValidFlagType = '2' // 冲正
)

// TFtdcCertCodeType是一个证件号码类型
type TThostFtdcCertCodeType [21]byte

func (t TThostFtdcCertCodeType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcCertCodeType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcReasonType是一个事由类型
type TThostFtdcReasonType uint8

//go:generate stringer -type TThostFtdcReasonType -linecomment
const (
	THOST_FTDC_RN_CD TThostFtdcReasonType = '0' // 错单
	THOST_FTDC_RN_ZT TThostFtdcReasonType = '1' // 资金在途
	THOST_FTDC_RN_QT TThostFtdcReasonType = '2' // 其它
)

// TFtdcFundProjectIDType是一个资金项目编号类型
type TThostFtdcFundProjectIDType [5]byte

func (t TThostFtdcFundProjectIDType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcFundProjectIDType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcSexType是一个性别类型
type TThostFtdcSexType uint8

//go:generate stringer -type TThostFtdcSexType -linecomment
const (
	THOST_FTDC_SEX_None  TThostFtdcSexType = '0' // 未知
	THOST_FTDC_SEX_Man   TThostFtdcSexType = '1' // 男
	THOST_FTDC_SEX_Woman TThostFtdcSexType = '2' // 女
)

// TFtdcProfessionType是一个职业类型
type TThostFtdcProfessionType [101]byte

func (t TThostFtdcProfessionType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcProfessionType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcNationalType是一个国籍类型
type TThostFtdcNationalType [31]byte

func (t TThostFtdcNationalType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcNationalType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcProvinceType是一个省类型
type TThostFtdcProvinceType [51]byte

func (t TThostFtdcProvinceType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcProvinceType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcRegionType是一个区类型
type TThostFtdcRegionType [16]byte

func (t TThostFtdcRegionType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcRegionType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcCountryType是一个国家类型
type TThostFtdcCountryType [16]byte

func (t TThostFtdcCountryType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcCountryType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcLicenseNOType是一个营业执照类型
type TThostFtdcLicenseNOType [33]byte

func (t TThostFtdcLicenseNOType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcLicenseNOType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcCompanyTypeType是一个企业性质类型
type TThostFtdcCompanyTypeType [16]byte

func (t TThostFtdcCompanyTypeType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcCompanyTypeType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcBusinessScopeType是一个经营范围类型
type TThostFtdcBusinessScopeType [1001]byte

func (t TThostFtdcBusinessScopeType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcBusinessScopeType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcCapitalCurrencyType是一个注册资本币种类型
type TThostFtdcCapitalCurrencyType [4]byte

func (t TThostFtdcCapitalCurrencyType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcCapitalCurrencyType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcUserTypeType是一个用户类型类型
type TThostFtdcUserTypeType uint8

//go:generate stringer -type TThostFtdcUserTypeType -linecomment
const (
	THOST_FTDC_UT_Investor  TThostFtdcUserTypeType = '0' // 投资者
	THOST_FTDC_UT_Operator  TThostFtdcUserTypeType = '1' // 操作员
	THOST_FTDC_UT_SuperUser TThostFtdcUserTypeType = '2' // 管理员
)

// TFtdcBranchIDType是一个营业部编号类型
type TThostFtdcBranchIDType [9]byte

func (t TThostFtdcBranchIDType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcBranchIDType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcRateTypeType是一个费率类型类型
type TThostFtdcRateTypeType uint8

//go:generate stringer -type TThostFtdcRateTypeType -linecomment
const (
	THOST_FTDC_RATETYPE_MarginRate TThostFtdcRateTypeType = '2' // 保证金率
)

// TFtdcNoteTypeType是一个通知类型类型
type TThostFtdcNoteTypeType uint8

//go:generate stringer -type TThostFtdcNoteTypeType -linecomment
const (
	THOST_FTDC_NOTETYPE_TradeSettleBill  TThostFtdcNoteTypeType = '1' // 交易结算单
	THOST_FTDC_NOTETYPE_TradeSettleMonth TThostFtdcNoteTypeType = '2' // 交易结算月报
	THOST_FTDC_NOTETYPE_CallMarginNotes  TThostFtdcNoteTypeType = '3' // 追加保证金通知书
	THOST_FTDC_NOTETYPE_ForceCloseNotes  TThostFtdcNoteTypeType = '4' // 强行平仓通知书
	THOST_FTDC_NOTETYPE_TradeNotes       TThostFtdcNoteTypeType = '5' // 成交通知书
	THOST_FTDC_NOTETYPE_DelivNotes       TThostFtdcNoteTypeType = '6' // 交割通知书
)

// TFtdcSettlementStyleType是一个结算单方式类型
type TThostFtdcSettlementStyleType uint8

//go:generate stringer -type TThostFtdcSettlementStyleType -linecomment
const (
	THOST_FTDC_SBS_Day    TThostFtdcSettlementStyleType = '1' // 逐日盯市
	THOST_FTDC_SBS_Volume TThostFtdcSettlementStyleType = '2' // 逐笔对冲
)

// TFtdcBrokerDNSType是一个域名类型
type TThostFtdcBrokerDNSType [256]byte

func (t TThostFtdcBrokerDNSType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcBrokerDNSType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcSentenceType是一个语句类型
type TThostFtdcSentenceType [501]byte

func (t TThostFtdcSentenceType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcSentenceType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcSettlementBillTypeType是一个结算单类型类型
type TThostFtdcSettlementBillTypeType uint8

//go:generate stringer -type TThostFtdcSettlementBillTypeType -linecomment
const (
	THOST_FTDC_ST_Day   TThostFtdcSettlementBillTypeType = '0' // 日报
	THOST_FTDC_ST_Month TThostFtdcSettlementBillTypeType = '1' // 月报
)

// TFtdcUserRightTypeType是一个客户权限类型类型
type TThostFtdcUserRightTypeType uint8

//go:generate stringer -type TThostFtdcUserRightTypeType -linecomment
const (
	THOST_FTDC_URT_Logon          TThostFtdcUserRightTypeType = '1' // 登录
	THOST_FTDC_URT_Transfer       TThostFtdcUserRightTypeType = '2' // 银期转帐
	THOST_FTDC_URT_EMail          TThostFtdcUserRightTypeType = '3' // 邮寄结算单
	THOST_FTDC_URT_Fax            TThostFtdcUserRightTypeType = '4' // 传真结算单
	THOST_FTDC_URT_ConditionOrder TThostFtdcUserRightTypeType = '5' // 条件单
)

// TFtdcMarginPriceTypeType是一个保证金价格类型类型
type TThostFtdcMarginPriceTypeType uint8

//go:generate stringer -type TThostFtdcMarginPriceTypeType -linecomment
const (
	THOST_FTDC_MPT_PreSettlementPrice TThostFtdcMarginPriceTypeType = '1' // 昨结算价
	THOST_FTDC_MPT_SettlementPrice    TThostFtdcMarginPriceTypeType = '2' // 最新价
	THOST_FTDC_MPT_AveragePrice       TThostFtdcMarginPriceTypeType = '3' // 成交均价
	THOST_FTDC_MPT_OpenPrice          TThostFtdcMarginPriceTypeType = '4' // 开仓价
)

// TFtdcBillGenStatusType是一个结算单生成状态类型
type TThostFtdcBillGenStatusType uint8

//go:generate stringer -type TThostFtdcBillGenStatusType -linecomment
const (
	THOST_FTDC_BGS_None        TThostFtdcBillGenStatusType = '0' // 未生成
	THOST_FTDC_BGS_NoGenerated TThostFtdcBillGenStatusType = '1' // 生成中
	THOST_FTDC_BGS_Generated   TThostFtdcBillGenStatusType = '2' // 已生成
)

// TFtdcAlgoTypeType是一个算法类型类型
type TThostFtdcAlgoTypeType uint8

//go:generate stringer -type TThostFtdcAlgoTypeType -linecomment
const (
	THOST_FTDC_AT_HandlePositionAlgo TThostFtdcAlgoTypeType = '1' // 持仓处理算法
	THOST_FTDC_AT_FindMarginRateAlgo TThostFtdcAlgoTypeType = '2' // 寻找保证金率算法
)

// TFtdcHandlePositionAlgoIDType是一个持仓处理算法编号类型
type TThostFtdcHandlePositionAlgoIDType uint8

//go:generate stringer -type TThostFtdcHandlePositionAlgoIDType -linecomment
const (
	THOST_FTDC_HPA_Base TThostFtdcHandlePositionAlgoIDType = '1' // 基本
	THOST_FTDC_HPA_DCE  TThostFtdcHandlePositionAlgoIDType = '2' // 大连商品交易所
	THOST_FTDC_HPA_CZCE TThostFtdcHandlePositionAlgoIDType = '3' // 郑州商品交易所
)

// TFtdcFindMarginRateAlgoIDType是一个寻找保证金率算法编号类型
type TThostFtdcFindMarginRateAlgoIDType uint8

//go:generate stringer -type TThostFtdcFindMarginRateAlgoIDType -linecomment
const (
	THOST_FTDC_FMRA_Base TThostFtdcFindMarginRateAlgoIDType = '1' // 基本
	THOST_FTDC_FMRA_DCE  TThostFtdcFindMarginRateAlgoIDType = '2' // 大连商品交易所
	THOST_FTDC_FMRA_CZCE TThostFtdcFindMarginRateAlgoIDType = '3' // 郑州商品交易所
)

// TFtdcHandleTradingAccountAlgoIDType是一个资金处理算法编号类型
type TThostFtdcHandleTradingAccountAlgoIDType uint8

//go:generate stringer -type TThostFtdcHandleTradingAccountAlgoIDType -linecomment
const (
	THOST_FTDC_HTAA_Base TThostFtdcHandleTradingAccountAlgoIDType = '1' // 基本
	THOST_FTDC_HTAA_DCE  TThostFtdcHandleTradingAccountAlgoIDType = '2' // 大连商品交易所
	THOST_FTDC_HTAA_CZCE TThostFtdcHandleTradingAccountAlgoIDType = '3' // 郑州商品交易所
)

// TFtdcPersonTypeType是一个联系人类型类型
type TThostFtdcPersonTypeType uint8

//go:generate stringer -type TThostFtdcPersonTypeType -linecomment
const (
	THOST_FTDC_PST_Order              TThostFtdcPersonTypeType = '1' // 指定下单人
	THOST_FTDC_PST_Open               TThostFtdcPersonTypeType = '2' // 开户授权人
	THOST_FTDC_PST_Fund               TThostFtdcPersonTypeType = '3' // 资金调拨人
	THOST_FTDC_PST_Settlement         TThostFtdcPersonTypeType = '4' // 结算单确认人
	THOST_FTDC_PST_Company            TThostFtdcPersonTypeType = '5' // 法人
	THOST_FTDC_PST_Corporation        TThostFtdcPersonTypeType = '6' // 法人代表
	THOST_FTDC_PST_LinkMan            TThostFtdcPersonTypeType = '7' // 投资者联系人
	THOST_FTDC_PST_Ledger             TThostFtdcPersonTypeType = '8' // 分户管理资产负责人
	THOST_FTDC_PST_Trustee            TThostFtdcPersonTypeType = '9' // 托（保）管人
	THOST_FTDC_PST_TrusteeCorporation TThostFtdcPersonTypeType = 'A' // 托（保）管机构法人代表
	THOST_FTDC_PST_TrusteeOpen        TThostFtdcPersonTypeType = 'B' // 托（保）管机构开户授权人
	THOST_FTDC_PST_TrusteeContact     TThostFtdcPersonTypeType = 'C' // 托（保）管机构联系人
	THOST_FTDC_PST_ForeignerRefer     TThostFtdcPersonTypeType = 'D' // 境外自然人参考证件
	THOST_FTDC_PST_CorporationRefer   TThostFtdcPersonTypeType = 'E' // 法人代表参考证件
)

// TFtdcQueryInvestorRangeType是一个查询范围类型
type TThostFtdcQueryInvestorRangeType uint8

//go:generate stringer -type TThostFtdcQueryInvestorRangeType -linecomment
const (
	THOST_FTDC_QIR_All    TThostFtdcQueryInvestorRangeType = '1' // 所有
	THOST_FTDC_QIR_Group  TThostFtdcQueryInvestorRangeType = '2' // 查询分类
	THOST_FTDC_QIR_Single TThostFtdcQueryInvestorRangeType = '3' // 单一投资者
)

// TFtdcInvestorRiskStatusType是一个投资者风险状态类型
type TThostFtdcInvestorRiskStatusType uint8

//go:generate stringer -type TThostFtdcInvestorRiskStatusType -linecomment
const (
	THOST_FTDC_IRS_Normal    TThostFtdcInvestorRiskStatusType = '1' // 正常
	THOST_FTDC_IRS_Warn      TThostFtdcInvestorRiskStatusType = '2' // 警告
	THOST_FTDC_IRS_Call      TThostFtdcInvestorRiskStatusType = '3' // 追保
	THOST_FTDC_IRS_Force     TThostFtdcInvestorRiskStatusType = '4' // 强平
	THOST_FTDC_IRS_Exception TThostFtdcInvestorRiskStatusType = '5' // 异常
)

// TFtdcLegIDType是一个单腿编号类型
type TThostFtdcLegIDType int32

// TFtdcLegMultipleType是一个单腿乘数类型
type TThostFtdcLegMultipleType int32

// TFtdcImplyLevelType是一个派生层数类型
type TThostFtdcImplyLevelType int32

// TFtdcClearAccountType是一个结算账户类型
type TThostFtdcClearAccountType [33]byte

func (t TThostFtdcClearAccountType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcClearAccountType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcOrganNOType是一个结算账户类型
type TThostFtdcOrganNOType [6]byte

func (t TThostFtdcOrganNOType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcOrganNOType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcClearbarchIDType是一个结算账户联行号类型
type TThostFtdcClearbarchIDType [6]byte

func (t TThostFtdcClearbarchIDType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcClearbarchIDType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcUserEventTypeType是一个用户事件类型类型
type TThostFtdcUserEventTypeType uint8

//go:generate stringer -type TThostFtdcUserEventTypeType -linecomment
const (
	THOST_FTDC_UET_Login          TThostFtdcUserEventTypeType = '1' // 登录
	THOST_FTDC_UET_Logout         TThostFtdcUserEventTypeType = '2' // 登出
	THOST_FTDC_UET_Trading        TThostFtdcUserEventTypeType = '3' // 交易成功
	THOST_FTDC_UET_TradingError   TThostFtdcUserEventTypeType = '4' // 交易失败
	THOST_FTDC_UET_UpdatePassword TThostFtdcUserEventTypeType = '5' // 修改密码
	THOST_FTDC_UET_Authenticate   TThostFtdcUserEventTypeType = '6' // 客户端认证
	THOST_FTDC_UET_Other          TThostFtdcUserEventTypeType = '9' // 其他
)

// TFtdcUserEventInfoType是一个用户事件信息类型
type TThostFtdcUserEventInfoType [1025]byte

func (t TThostFtdcUserEventInfoType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcUserEventInfoType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcCloseStyleType是一个平仓方式类型
type TThostFtdcCloseStyleType uint8

//go:generate stringer -type TThostFtdcCloseStyleType -linecomment
const (
	THOST_FTDC_ICS_Close      TThostFtdcCloseStyleType = '0' // 先开先平
	THOST_FTDC_ICS_CloseToday TThostFtdcCloseStyleType = '1' // 先平今再平昨
)

// TFtdcStatModeType是一个统计方式类型
type TThostFtdcStatModeType uint8

//go:generate stringer -type TThostFtdcStatModeType -linecomment
const (
	THOST_FTDC_SM_Non        TThostFtdcStatModeType = '0' // ----
	THOST_FTDC_SM_Instrument TThostFtdcStatModeType = '1' // 按合约统计
	THOST_FTDC_SM_Product    TThostFtdcStatModeType = '2' // 按产品统计
	THOST_FTDC_SM_Investor   TThostFtdcStatModeType = '3' // 按投资者统计
	THOST_FTDC_SM_Normal     TThostFtdcStatModeType = '1' // 正常
	THOST_FTDC_SM_Emerge     TThostFtdcStatModeType = '2' // 应急
	THOST_FTDC_SM_Restore    TThostFtdcStatModeType = '3' // 恢复
)

// TFtdcParkedOrderStatusType是一个预埋单状态类型
type TThostFtdcParkedOrderStatusType uint8

//go:generate stringer -type TThostFtdcParkedOrderStatusType -linecomment
const (
	THOST_FTDC_PAOS_NotSend TThostFtdcParkedOrderStatusType = '1' // 未发送
	THOST_FTDC_PAOS_Send    TThostFtdcParkedOrderStatusType = '2' // 已发送
	THOST_FTDC_PAOS_Deleted TThostFtdcParkedOrderStatusType = '3' // 已删除
)

// TFtdcParkedOrderIDType是一个预埋报单编号类型
type TThostFtdcParkedOrderIDType [13]byte

func (t TThostFtdcParkedOrderIDType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcParkedOrderIDType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcParkedOrderActionIDType是一个预埋撤单编号类型
type TThostFtdcParkedOrderActionIDType [13]byte

func (t TThostFtdcParkedOrderActionIDType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcParkedOrderActionIDType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcVirDealStatusType是一个处理状态类型
type TThostFtdcVirDealStatusType uint8

//go:generate stringer -type TThostFtdcVirDealStatusType -linecomment
const (
	THOST_FTDC_VDS_Dealing      TThostFtdcVirDealStatusType = '1' // 正在处理
	THOST_FTDC_VDS_DeaclSucceed TThostFtdcVirDealStatusType = '2' // 处理成功
)

// TFtdcOrgSystemIDType是一个原有系统代码类型
type TThostFtdcOrgSystemIDType uint8

//go:generate stringer -type TThostFtdcOrgSystemIDType -linecomment
const (
	THOST_FTDC_ORGS_Standard   TThostFtdcOrgSystemIDType = '0' // 综合交易平台
	THOST_FTDC_ORGS_ESunny     TThostFtdcOrgSystemIDType = '1' // 易盛系统
	THOST_FTDC_ORGS_KingStarV6 TThostFtdcOrgSystemIDType = '2' // 金仕达V6系统
)

// TFtdcVirTradeStatusType是一个交易状态类型
type TThostFtdcVirTradeStatusType uint8

//go:generate stringer -type TThostFtdcVirTradeStatusType -linecomment
const (
	THOST_FTDC_VTS_NaturalDeal  TThostFtdcVirTradeStatusType = '0' // 正常处理中
	THOST_FTDC_VTS_SucceedEnd   TThostFtdcVirTradeStatusType = '1' // 成功结束
	THOST_FTDC_VTS_FailedEND    TThostFtdcVirTradeStatusType = '2' // 失败结束
	THOST_FTDC_VTS_Exception    TThostFtdcVirTradeStatusType = '3' // 异常中
	THOST_FTDC_VTS_ManualDeal   TThostFtdcVirTradeStatusType = '4' // 已人工异常处理
	THOST_FTDC_VTS_MesException TThostFtdcVirTradeStatusType = '5' // 通讯异常 ，请人工处理
	THOST_FTDC_VTS_SysException TThostFtdcVirTradeStatusType = '6' // 系统出错，请人工处理
)

// TFtdcVirBankAccTypeType是一个银行帐户类型类型
type TThostFtdcVirBankAccTypeType uint8

//go:generate stringer -type TThostFtdcVirBankAccTypeType -linecomment
const (
	THOST_FTDC_VBAT_BankBook   TThostFtdcVirBankAccTypeType = '1' // 存折
	THOST_FTDC_VBAT_BankCard   TThostFtdcVirBankAccTypeType = '2' // 储蓄卡
	THOST_FTDC_VBAT_CreditCard TThostFtdcVirBankAccTypeType = '3' // 信用卡
)

// TFtdcVirementStatusType是一个银行帐户类型类型
type TThostFtdcVirementStatusType uint8

//go:generate stringer -type TThostFtdcVirementStatusType -linecomment
const (
	THOST_FTDC_VMS_Natural  TThostFtdcVirementStatusType = '0' // 正常
	THOST_FTDC_VMS_Canceled TThostFtdcVirementStatusType = '9' // 销户
)

// TFtdcVirementAvailAbilityType是一个有效标志类型
type TThostFtdcVirementAvailAbilityType uint8

//go:generate stringer -type TThostFtdcVirementAvailAbilityType -linecomment
const (
	THOST_FTDC_VAA_NoAvailAbility TThostFtdcVirementAvailAbilityType = '0' // 未确认
	THOST_FTDC_VAA_AvailAbility   TThostFtdcVirementAvailAbilityType = '1' // 有效
	THOST_FTDC_VAA_Repeal         TThostFtdcVirementAvailAbilityType = '2' // 冲正
)

// TFtdcVirementTradeCodeType是一个交易代码类型
type TThostFtdcVirementTradeCodeType uint8

//go:generate stringer -type TThostFtdcVirementTradeCodeType -linecomment
const (
	THOST_FTDC_VTC_BankBankToFuture   TThostFtdcVirementTradeCodeType = '1' // 银行发起银行资金转期货
	THOST_FTDC_VTC_BankFutureToBank   TThostFtdcVirementTradeCodeType = '2' // 银行发起期货资金转银行
	THOST_FTDC_VTC_FutureBankToFuture TThostFtdcVirementTradeCodeType = '3' // 期货发起银行资金转期货
	THOST_FTDC_VTC_FutureFutureToBank TThostFtdcVirementTradeCodeType = '4' // 期货发起期货资金转银行
)

// TFtdcPhotoTypeNameType是一个影像类型名称类型
type TThostFtdcPhotoTypeNameType [41]byte

func (t TThostFtdcPhotoTypeNameType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcPhotoTypeNameType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcPhotoTypeIDType是一个影像类型代码类型
type TThostFtdcPhotoTypeIDType [5]byte

func (t TThostFtdcPhotoTypeIDType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcPhotoTypeIDType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcPhotoNameType是一个影像名称类型
type TThostFtdcPhotoNameType [161]byte

func (t TThostFtdcPhotoNameType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcPhotoNameType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcTopicIDType是一个主题代码类型
type TThostFtdcTopicIDType int32

// TFtdcReportTypeIDType是一个交易报告类型标识类型
type TThostFtdcReportTypeIDType [3]byte

func (t TThostFtdcReportTypeIDType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcReportTypeIDType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcCharacterIDType是一个交易特征代码类型
type TThostFtdcCharacterIDType [5]byte

func (t TThostFtdcCharacterIDType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcCharacterIDType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcAMLParamIDType是一个参数代码类型
type TThostFtdcAMLParamIDType [21]byte

func (t TThostFtdcAMLParamIDType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcAMLParamIDType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcAMLInvestorTypeType是一个投资者类型类型
type TThostFtdcAMLInvestorTypeType [3]byte

func (t TThostFtdcAMLInvestorTypeType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcAMLInvestorTypeType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcAMLIdCardTypeType是一个证件类型类型
type TThostFtdcAMLIdCardTypeType [3]byte

func (t TThostFtdcAMLIdCardTypeType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcAMLIdCardTypeType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcAMLTradeDirectType是一个资金进出方向类型
type TThostFtdcAMLTradeDirectType [3]byte

func (t TThostFtdcAMLTradeDirectType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcAMLTradeDirectType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcAMLTradeModelType是一个资金进出方式类型
type TThostFtdcAMLTradeModelType [3]byte

func (t TThostFtdcAMLTradeModelType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcAMLTradeModelType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcAMLOpParamValueType是一个业务参数代码值类型
type TThostFtdcAMLOpParamValueType float64

func (t TThostFtdcAMLOpParamValueType) IsValid() bool {
	return float64(t) != math.MaxFloat64
}

func (t TThostFtdcAMLOpParamValueType) String() string {
	if t.IsValid() {
		return strconv.FormatFloat(float64(t), 'f', 6, 64)
	} else {
		return "DBL_MAX"
	}
}

// TFtdcAMLCustomerCardTypeType是一个客户身份证件证明文件类型类型
type TThostFtdcAMLCustomerCardTypeType [81]byte

func (t TThostFtdcAMLCustomerCardTypeType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcAMLCustomerCardTypeType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcAMLInstitutionNameType是一个金融机构网点名称类型
type TThostFtdcAMLInstitutionNameType [65]byte

func (t TThostFtdcAMLInstitutionNameType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcAMLInstitutionNameType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcAMLDistrictIDType是一个金融机构网点所在地区行政区划代码类型
type TThostFtdcAMLDistrictIDType [7]byte

func (t TThostFtdcAMLDistrictIDType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcAMLDistrictIDType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcAMLRelationShipType是一个金融机构网点与大额交易的关系类型
type TThostFtdcAMLRelationShipType [3]byte

func (t TThostFtdcAMLRelationShipType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcAMLRelationShipType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcAMLInstitutionTypeType是一个金融机构网点代码类型类型
type TThostFtdcAMLInstitutionTypeType [3]byte

func (t TThostFtdcAMLInstitutionTypeType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcAMLInstitutionTypeType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcAMLInstitutionIDType是一个金融机构网点代码类型
type TThostFtdcAMLInstitutionIDType [13]byte

func (t TThostFtdcAMLInstitutionIDType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcAMLInstitutionIDType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcAMLAccountTypeType是一个账户类型类型
type TThostFtdcAMLAccountTypeType [5]byte

func (t TThostFtdcAMLAccountTypeType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcAMLAccountTypeType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcAMLTradingTypeType是一个交易方式类型
type TThostFtdcAMLTradingTypeType [7]byte

func (t TThostFtdcAMLTradingTypeType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcAMLTradingTypeType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcAMLTransactClassType是一个涉外收支交易分类与代码类型
type TThostFtdcAMLTransactClassType [7]byte

func (t TThostFtdcAMLTransactClassType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcAMLTransactClassType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcAMLCapitalIOType是一个资金收付标识类型
type TThostFtdcAMLCapitalIOType [3]byte

func (t TThostFtdcAMLCapitalIOType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcAMLCapitalIOType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcAMLSiteType是一个交易地点类型
type TThostFtdcAMLSiteType [10]byte

func (t TThostFtdcAMLSiteType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcAMLSiteType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcAMLCapitalPurposeType是一个资金用途类型
type TThostFtdcAMLCapitalPurposeType [129]byte

func (t TThostFtdcAMLCapitalPurposeType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcAMLCapitalPurposeType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcAMLReportTypeType是一个报文类型类型
type TThostFtdcAMLReportTypeType [2]byte

func (t TThostFtdcAMLReportTypeType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcAMLReportTypeType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcAMLSerialNoType是一个编号类型
type TThostFtdcAMLSerialNoType [5]byte

func (t TThostFtdcAMLSerialNoType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcAMLSerialNoType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcAMLStatusType是一个状态类型
type TThostFtdcAMLStatusType [2]byte

func (t TThostFtdcAMLStatusType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcAMLStatusType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcAMLGenStatusType是一个Aml生成方式类型
type TThostFtdcAMLGenStatusType uint8

//go:generate stringer -type TThostFtdcAMLGenStatusType -linecomment
const (
	THOST_FTDC_GEN_Program  TThostFtdcAMLGenStatusType = '0' // 程序生成
	THOST_FTDC_GEN_HandWork TThostFtdcAMLGenStatusType = '1' // 人工生成
)

// TFtdcAMLSeqCodeType是一个业务标识号类型
type TThostFtdcAMLSeqCodeType [65]byte

func (t TThostFtdcAMLSeqCodeType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcAMLSeqCodeType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcAMLFileNameType是一个AML文件名类型
type TThostFtdcAMLFileNameType [257]byte

func (t TThostFtdcAMLFileNameType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcAMLFileNameType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcAMLMoneyType是一个反洗钱资金类型
type TThostFtdcAMLMoneyType float64

func (t TThostFtdcAMLMoneyType) IsValid() bool {
	return float64(t) != math.MaxFloat64
}

func (t TThostFtdcAMLMoneyType) String() string {
	if t.IsValid() {
		return strconv.FormatFloat(float64(t), 'f', 6, 64)
	} else {
		return "DBL_MAX"
	}
}

// TFtdcAMLFileAmountType是一个反洗钱资金类型
type TThostFtdcAMLFileAmountType int32

// TFtdcCFMMCKeyType是一个密钥类型(保证金监管)类型
type TThostFtdcCFMMCKeyType [21]byte

func (t TThostFtdcCFMMCKeyType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcCFMMCKeyType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcCFMMCTokenType是一个令牌类型(保证金监管)类型
type TThostFtdcCFMMCTokenType [21]byte

func (t TThostFtdcCFMMCTokenType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcCFMMCTokenType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcCFMMCKeyKindType是一个动态密钥类别(保证金监管)类型
type TThostFtdcCFMMCKeyKindType uint8

//go:generate stringer -type TThostFtdcCFMMCKeyKindType -linecomment
const (
	THOST_FTDC_CFMMCKK_REQUEST TThostFtdcCFMMCKeyKindType = 'R' // 主动请求更新
	THOST_FTDC_CFMMCKK_AUTO    TThostFtdcCFMMCKeyKindType = 'A' // CFMMC自动更新
	THOST_FTDC_CFMMCKK_MANUAL  TThostFtdcCFMMCKeyKindType = 'M' // CFMMC手动更新
)

// TFtdcAMLReportNameType是一个报文名称类型
type TThostFtdcAMLReportNameType [81]byte

func (t TThostFtdcAMLReportNameType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcAMLReportNameType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcIndividualNameType是一个个人姓名类型
type TThostFtdcIndividualNameType [51]byte

func (t TThostFtdcIndividualNameType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcIndividualNameType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcCurrencyIDType是一个币种代码类型
type TThostFtdcCurrencyIDType [4]byte

func (t TThostFtdcCurrencyIDType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcCurrencyIDType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcCustNumberType是一个客户编号类型
type TThostFtdcCustNumberType [36]byte

func (t TThostFtdcCustNumberType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcCustNumberType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcOrganCodeType是一个机构编码类型
type TThostFtdcOrganCodeType [36]byte

func (t TThostFtdcOrganCodeType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcOrganCodeType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcOrganNameType是一个机构名称类型
type TThostFtdcOrganNameType [71]byte

func (t TThostFtdcOrganNameType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcOrganNameType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcSuperOrganCodeType是一个上级机构编码,即期货公司总部、银行总行类型
type TThostFtdcSuperOrganCodeType [12]byte

func (t TThostFtdcSuperOrganCodeType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcSuperOrganCodeType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcSubBranchIDType是一个分支机构类型
type TThostFtdcSubBranchIDType [31]byte

func (t TThostFtdcSubBranchIDType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcSubBranchIDType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcSubBranchNameType是一个分支机构名称类型
type TThostFtdcSubBranchNameType [71]byte

func (t TThostFtdcSubBranchNameType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcSubBranchNameType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcBranchNetCodeType是一个机构网点号类型
type TThostFtdcBranchNetCodeType [31]byte

func (t TThostFtdcBranchNetCodeType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcBranchNetCodeType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcBranchNetNameType是一个机构网点名称类型
type TThostFtdcBranchNetNameType [71]byte

func (t TThostFtdcBranchNetNameType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcBranchNetNameType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcOrganFlagType是一个机构标识类型
type TThostFtdcOrganFlagType [2]byte

func (t TThostFtdcOrganFlagType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcOrganFlagType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcBankCodingForFutureType是一个银行对期货公司的编码类型
type TThostFtdcBankCodingForFutureType [33]byte

func (t TThostFtdcBankCodingForFutureType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcBankCodingForFutureType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcBankReturnCodeType是一个银行对返回码的定义类型
type TThostFtdcBankReturnCodeType [7]byte

func (t TThostFtdcBankReturnCodeType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcBankReturnCodeType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcPlateReturnCodeType是一个银期转帐平台对返回码的定义类型
type TThostFtdcPlateReturnCodeType [5]byte

func (t TThostFtdcPlateReturnCodeType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcPlateReturnCodeType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcBankSubBranchIDType是一个银行分支机构编码类型
type TThostFtdcBankSubBranchIDType [31]byte

func (t TThostFtdcBankSubBranchIDType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcBankSubBranchIDType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcFutureBranchIDType是一个期货分支机构编码类型
type TThostFtdcFutureBranchIDType [31]byte

func (t TThostFtdcFutureBranchIDType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcFutureBranchIDType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcReturnCodeType是一个返回代码类型
type TThostFtdcReturnCodeType [7]byte

func (t TThostFtdcReturnCodeType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcReturnCodeType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcOperatorCodeType是一个操作员类型
type TThostFtdcOperatorCodeType [17]byte

func (t TThostFtdcOperatorCodeType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcOperatorCodeType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcClearDepIDType是一个机构结算帐户机构号类型
type TThostFtdcClearDepIDType [6]byte

func (t TThostFtdcClearDepIDType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcClearDepIDType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcClearBrchIDType是一个机构结算帐户联行号类型
type TThostFtdcClearBrchIDType [6]byte

func (t TThostFtdcClearBrchIDType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcClearBrchIDType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcClearNameType是一个机构结算帐户名称类型
type TThostFtdcClearNameType [71]byte

func (t TThostFtdcClearNameType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcClearNameType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcBankAccountNameType是一个银行帐户名称类型
type TThostFtdcBankAccountNameType [71]byte

func (t TThostFtdcBankAccountNameType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcBankAccountNameType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcInvDepIDType是一个机构投资人账号机构号类型
type TThostFtdcInvDepIDType [6]byte

func (t TThostFtdcInvDepIDType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcInvDepIDType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcInvBrchIDType是一个机构投资人联行号类型
type TThostFtdcInvBrchIDType [6]byte

func (t TThostFtdcInvBrchIDType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcInvBrchIDType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcMessageFormatVersionType是一个信息格式版本类型
type TThostFtdcMessageFormatVersionType [36]byte

func (t TThostFtdcMessageFormatVersionType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcMessageFormatVersionType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcDigestType是一个摘要类型
type TThostFtdcDigestType [36]byte

func (t TThostFtdcDigestType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcDigestType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcAuthenticDataType是一个认证数据类型
type TThostFtdcAuthenticDataType [129]byte

func (t TThostFtdcAuthenticDataType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcAuthenticDataType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcPasswordKeyType是一个密钥类型
type TThostFtdcPasswordKeyType [129]byte

func (t TThostFtdcPasswordKeyType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcPasswordKeyType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcFutureAccountNameType是一个期货帐户名称类型
type TThostFtdcFutureAccountNameType [129]byte

func (t TThostFtdcFutureAccountNameType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcFutureAccountNameType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcMobilePhoneType是一个手机类型
type TThostFtdcMobilePhoneType [21]byte

func (t TThostFtdcMobilePhoneType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcMobilePhoneType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcFutureMainKeyType是一个期货公司主密钥类型
type TThostFtdcFutureMainKeyType [129]byte

func (t TThostFtdcFutureMainKeyType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcFutureMainKeyType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcFutureWorkKeyType是一个期货公司工作密钥类型
type TThostFtdcFutureWorkKeyType [129]byte

func (t TThostFtdcFutureWorkKeyType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcFutureWorkKeyType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcFutureTransKeyType是一个期货公司传输密钥类型
type TThostFtdcFutureTransKeyType [129]byte

func (t TThostFtdcFutureTransKeyType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcFutureTransKeyType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcBankMainKeyType是一个银行主密钥类型
type TThostFtdcBankMainKeyType [129]byte

func (t TThostFtdcBankMainKeyType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcBankMainKeyType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcBankWorkKeyType是一个银行工作密钥类型
type TThostFtdcBankWorkKeyType [129]byte

func (t TThostFtdcBankWorkKeyType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcBankWorkKeyType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcBankTransKeyType是一个银行传输密钥类型
type TThostFtdcBankTransKeyType [129]byte

func (t TThostFtdcBankTransKeyType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcBankTransKeyType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcBankServerDescriptionType是一个银行服务器描述信息类型
type TThostFtdcBankServerDescriptionType [129]byte

func (t TThostFtdcBankServerDescriptionType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcBankServerDescriptionType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcAddInfoType是一个附加信息类型
type TThostFtdcAddInfoType [129]byte

func (t TThostFtdcAddInfoType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcAddInfoType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcDescrInfoForReturnCodeType是一个返回码描述类型
type TThostFtdcDescrInfoForReturnCodeType [129]byte

func (t TThostFtdcDescrInfoForReturnCodeType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcDescrInfoForReturnCodeType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcCountryCodeType是一个国家代码类型
type TThostFtdcCountryCodeType [21]byte

func (t TThostFtdcCountryCodeType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcCountryCodeType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcSerialType是一个流水号类型
type TThostFtdcSerialType int32

// TFtdcPlateSerialType是一个平台流水号类型
type TThostFtdcPlateSerialType int32

// TFtdcBankSerialType是一个银行流水号类型
type TThostFtdcBankSerialType [13]byte

func (t TThostFtdcBankSerialType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcBankSerialType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcCorrectSerialType是一个被冲正交易流水号类型
type TThostFtdcCorrectSerialType int32

// TFtdcFutureSerialType是一个期货公司流水号类型
type TThostFtdcFutureSerialType int32

// TFtdcApplicationIDType是一个应用标识类型
type TThostFtdcApplicationIDType int32

// TFtdcBankProxyIDType是一个银行代理标识类型
type TThostFtdcBankProxyIDType int32

// TFtdcFBTCoreIDType是一个银期转帐核心系统标识类型
type TThostFtdcFBTCoreIDType int32

// TFtdcServerPortType是一个服务端口号类型
type TThostFtdcServerPortType int32

// TFtdcRepealedTimesType是一个已经冲正次数类型
type TThostFtdcRepealedTimesType int32

// TFtdcRepealTimeIntervalType是一个冲正时间间隔类型
type TThostFtdcRepealTimeIntervalType int32

// TFtdcTotalTimesType是一个每日累计转帐次数类型
type TThostFtdcTotalTimesType int32

// TFtdcFBTRequestIDType是一个请求ID类型
type TThostFtdcFBTRequestIDType int32

// TFtdcTIDType是一个交易ID类型
type TThostFtdcTIDType int32

// TFtdcTradeAmountType是一个交易金额（元）类型
type TThostFtdcTradeAmountType float64

func (t TThostFtdcTradeAmountType) IsValid() bool {
	return float64(t) != math.MaxFloat64
}

func (t TThostFtdcTradeAmountType) String() string {
	if t.IsValid() {
		return strconv.FormatFloat(float64(t), 'f', 6, 64)
	} else {
		return "DBL_MAX"
	}
}

// TFtdcCustFeeType是一个应收客户费用（元）类型
type TThostFtdcCustFeeType float64

func (t TThostFtdcCustFeeType) IsValid() bool {
	return float64(t) != math.MaxFloat64
}

func (t TThostFtdcCustFeeType) String() string {
	if t.IsValid() {
		return strconv.FormatFloat(float64(t), 'f', 6, 64)
	} else {
		return "DBL_MAX"
	}
}

// TFtdcFutureFeeType是一个应收期货公司费用（元）类型
type TThostFtdcFutureFeeType float64

func (t TThostFtdcFutureFeeType) IsValid() bool {
	return float64(t) != math.MaxFloat64
}

func (t TThostFtdcFutureFeeType) String() string {
	if t.IsValid() {
		return strconv.FormatFloat(float64(t), 'f', 6, 64)
	} else {
		return "DBL_MAX"
	}
}

// TFtdcSingleMaxAmtType是一个单笔最高限额类型
type TThostFtdcSingleMaxAmtType float64

func (t TThostFtdcSingleMaxAmtType) IsValid() bool {
	return float64(t) != math.MaxFloat64
}

func (t TThostFtdcSingleMaxAmtType) String() string {
	if t.IsValid() {
		return strconv.FormatFloat(float64(t), 'f', 6, 64)
	} else {
		return "DBL_MAX"
	}
}

// TFtdcSingleMinAmtType是一个单笔最低限额类型
type TThostFtdcSingleMinAmtType float64

func (t TThostFtdcSingleMinAmtType) IsValid() bool {
	return float64(t) != math.MaxFloat64
}

func (t TThostFtdcSingleMinAmtType) String() string {
	if t.IsValid() {
		return strconv.FormatFloat(float64(t), 'f', 6, 64)
	} else {
		return "DBL_MAX"
	}
}

// TFtdcTotalAmtType是一个每日累计转帐额度类型
type TThostFtdcTotalAmtType float64

func (t TThostFtdcTotalAmtType) IsValid() bool {
	return float64(t) != math.MaxFloat64
}

func (t TThostFtdcTotalAmtType) String() string {
	if t.IsValid() {
		return strconv.FormatFloat(float64(t), 'f', 6, 64)
	} else {
		return "DBL_MAX"
	}
}

// TFtdcCertificationTypeType是一个证件类型类型
type TThostFtdcCertificationTypeType uint8

//go:generate stringer -type TThostFtdcCertificationTypeType -linecomment
const (
	THOST_FTDC_CFT_IDCard                TThostFtdcCertificationTypeType = '0' // 身份证
	THOST_FTDC_CFT_Passport              TThostFtdcCertificationTypeType = '1' // 护照
	THOST_FTDC_CFT_OfficerIDCard         TThostFtdcCertificationTypeType = '2' // 军官证
	THOST_FTDC_CFT_SoldierIDCard         TThostFtdcCertificationTypeType = '3' // 士兵证
	THOST_FTDC_CFT_HomeComingCard        TThostFtdcCertificationTypeType = '4' // 回乡证
	THOST_FTDC_CFT_HouseholdRegister     TThostFtdcCertificationTypeType = '5' // 户口簿
	THOST_FTDC_CFT_LicenseNo             TThostFtdcCertificationTypeType = '6' // 营业执照号
	THOST_FTDC_CFT_InstitutionCodeCard   TThostFtdcCertificationTypeType = '7' // 组织机构代码证
	THOST_FTDC_CFT_TempLicenseNo         TThostFtdcCertificationTypeType = '8' // 临时营业执照号
	THOST_FTDC_CFT_NoEnterpriseLicenseNo TThostFtdcCertificationTypeType = '9' // 民办非企业登记证书
	THOST_FTDC_CFT_OtherCard             TThostFtdcCertificationTypeType = 'x' // 其他证件
	THOST_FTDC_CFT_SuperDepAgree         TThostFtdcCertificationTypeType = 'a' // 主管部门批文
)

// TFtdcFileBusinessCodeType是一个文件业务功能类型
type TThostFtdcFileBusinessCodeType uint8

//go:generate stringer -type TThostFtdcFileBusinessCodeType -linecomment
const (
	THOST_FTDC_FBC_Others                         TThostFtdcFileBusinessCodeType = '0' // 其他
	THOST_FTDC_FBC_TransferDetails                TThostFtdcFileBusinessCodeType = '1' // 转账交易明细对账
	THOST_FTDC_FBC_CustAccStatus                  TThostFtdcFileBusinessCodeType = '2' // 客户账户状态对账
	THOST_FTDC_FBC_AccountTradeDetails            TThostFtdcFileBusinessCodeType = '3' // 账户类交易明细对账
	THOST_FTDC_FBC_FutureAccountChangeInfoDetails TThostFtdcFileBusinessCodeType = '4' // 期货账户信息变更明细对账
	THOST_FTDC_FBC_CustMoneyDetail                TThostFtdcFileBusinessCodeType = '5' // 客户资金台账余额明细对账
	THOST_FTDC_FBC_CustCancelAccountInfo          TThostFtdcFileBusinessCodeType = '6' // 客户销户结息明细对账
	THOST_FTDC_FBC_CustMoneyResult                TThostFtdcFileBusinessCodeType = '7' // 客户资金余额对账结果
	THOST_FTDC_FBC_OthersExceptionResult          TThostFtdcFileBusinessCodeType = '8' // 其它对账异常结果文件
	THOST_FTDC_FBC_CustInterestNetMoneyDetails    TThostFtdcFileBusinessCodeType = '9' // 客户结息净额明细
	THOST_FTDC_FBC_CustMoneySendAndReceiveDetails TThostFtdcFileBusinessCodeType = 'a' // 客户资金交收明细
	THOST_FTDC_FBC_CorporationMoneyTotal          TThostFtdcFileBusinessCodeType = 'b' // 法人存管银行资金交收汇总
	THOST_FTDC_FBC_MainbodyMoneyTotal             TThostFtdcFileBusinessCodeType = 'c' // 主体间资金交收汇总
	THOST_FTDC_FBC_MainPartMonitorData            TThostFtdcFileBusinessCodeType = 'd' // 总分平衡监管数据
	THOST_FTDC_FBC_PreparationMoney               TThostFtdcFileBusinessCodeType = 'e' // 存管银行备付金余额
	THOST_FTDC_FBC_BankMoneyMonitorData           TThostFtdcFileBusinessCodeType = 'f' // 协办存管银行资金监管数据
)

// TFtdcCashExchangeCodeType是一个汇钞标志类型
type TThostFtdcCashExchangeCodeType uint8

//go:generate stringer -type TThostFtdcCashExchangeCodeType -linecomment
const (
	THOST_FTDC_CEC_Exchange TThostFtdcCashExchangeCodeType = '1' // 汇
	THOST_FTDC_CEC_Cash     TThostFtdcCashExchangeCodeType = '2' // 钞
)

// TFtdcYesNoIndicatorType是一个是或否标识类型
type TThostFtdcYesNoIndicatorType uint8

//go:generate stringer -type TThostFtdcYesNoIndicatorType -linecomment
const (
	THOST_FTDC_YNI_Yes TThostFtdcYesNoIndicatorType = '0' // 是
	THOST_FTDC_YNI_No  TThostFtdcYesNoIndicatorType = '1' // 否
)

// TFtdcBanlanceTypeType是一个余额类型类型
type TThostFtdcBanlanceTypeType uint8

//go:generate stringer -type TThostFtdcBanlanceTypeType -linecomment
const (
	THOST_FTDC_BLT_CurrentMoney   TThostFtdcBanlanceTypeType = '0' // 当前余额
	THOST_FTDC_BLT_UsableMoney    TThostFtdcBanlanceTypeType = '1' // 可用余额
	THOST_FTDC_BLT_FetchableMoney TThostFtdcBanlanceTypeType = '2' // 可取余额
	THOST_FTDC_BLT_FreezeMoney    TThostFtdcBanlanceTypeType = '3' // 冻结余额
)

// TFtdcGenderType是一个性别类型
type TThostFtdcGenderType uint8

//go:generate stringer -type TThostFtdcGenderType -linecomment
const (
	THOST_FTDC_GD_Unknown TThostFtdcGenderType = '0' // 未知状态
	THOST_FTDC_GD_Male    TThostFtdcGenderType = '1' // 男
	THOST_FTDC_GD_Female  TThostFtdcGenderType = '2' // 女
)

// TFtdcFeePayFlagType是一个费用支付标志类型
type TThostFtdcFeePayFlagType uint8

//go:generate stringer -type TThostFtdcFeePayFlagType -linecomment
const (
	THOST_FTDC_FPF_BEN TThostFtdcFeePayFlagType = '0' // 由受益方支付费用
	THOST_FTDC_FPF_OUR TThostFtdcFeePayFlagType = '1' // 由发送方支付费用
	THOST_FTDC_FPF_SHA TThostFtdcFeePayFlagType = '2' // 由发送方支付发起的费用，受益方支付接受的费用
)

// TFtdcPassWordKeyTypeType是一个密钥类型类型
type TThostFtdcPassWordKeyTypeType uint8

//go:generate stringer -type TThostFtdcPassWordKeyTypeType -linecomment
const (
	THOST_FTDC_PWKT_ExchangeKey TThostFtdcPassWordKeyTypeType = '0' // 交换密钥
	THOST_FTDC_PWKT_PassWordKey TThostFtdcPassWordKeyTypeType = '1' // 密码密钥
	THOST_FTDC_PWKT_MACKey      TThostFtdcPassWordKeyTypeType = '2' // MAC密钥
	THOST_FTDC_PWKT_MessageKey  TThostFtdcPassWordKeyTypeType = '3' // 报文密钥
)

// TFtdcFBTPassWordTypeType是一个密码类型类型
type TThostFtdcFBTPassWordTypeType uint8

//go:generate stringer -type TThostFtdcFBTPassWordTypeType -linecomment
const (
	THOST_FTDC_PWT_Query    TThostFtdcFBTPassWordTypeType = '0' // 查询
	THOST_FTDC_PWT_Fetch    TThostFtdcFBTPassWordTypeType = '1' // 取款
	THOST_FTDC_PWT_Transfer TThostFtdcFBTPassWordTypeType = '2' // 转帐
	THOST_FTDC_PWT_Trade    TThostFtdcFBTPassWordTypeType = '3' // 交易
)

// TFtdcFBTEncryModeType是一个加密方式类型
type TThostFtdcFBTEncryModeType uint8

//go:generate stringer -type TThostFtdcFBTEncryModeType -linecomment
const (
	THOST_FTDC_EM_NoEncry TThostFtdcFBTEncryModeType = '0' // 不加密
	THOST_FTDC_EM_DES     TThostFtdcFBTEncryModeType = '1' // DES
	THOST_FTDC_EM_3DES    TThostFtdcFBTEncryModeType = '2' // 3DES
)

// TFtdcBankRepealFlagType是一个银行冲正标志类型
type TThostFtdcBankRepealFlagType uint8

//go:generate stringer -type TThostFtdcBankRepealFlagType -linecomment
const (
	THOST_FTDC_BRF_BankNotNeedRepeal TThostFtdcBankRepealFlagType = '0' // 银行无需自动冲正
	THOST_FTDC_BRF_BankWaitingRepeal TThostFtdcBankRepealFlagType = '1' // 银行待自动冲正
	THOST_FTDC_BRF_BankBeenRepealed  TThostFtdcBankRepealFlagType = '2' // 银行已自动冲正
)

// TFtdcBrokerRepealFlagType是一个期商冲正标志类型
type TThostFtdcBrokerRepealFlagType uint8

//go:generate stringer -type TThostFtdcBrokerRepealFlagType -linecomment
const (
	THOST_FTDC_BRORF_BrokerNotNeedRepeal TThostFtdcBrokerRepealFlagType = '0' // 期商无需自动冲正
	THOST_FTDC_BRORF_BrokerWaitingRepeal TThostFtdcBrokerRepealFlagType = '1' // 期商待自动冲正
	THOST_FTDC_BRORF_BrokerBeenRepealed  TThostFtdcBrokerRepealFlagType = '2' // 期商已自动冲正
)

// TFtdcInstitutionTypeType是一个机构类别类型
type TThostFtdcInstitutionTypeType uint8

//go:generate stringer -type TThostFtdcInstitutionTypeType -linecomment
const (
	THOST_FTDC_TS_Bank   TThostFtdcInstitutionTypeType = '0' // 银行
	THOST_FTDC_TS_Future TThostFtdcInstitutionTypeType = '1' // 期商
	THOST_FTDC_TS_Store  TThostFtdcInstitutionTypeType = '2' // 券商
)

// TFtdcLastFragmentType是一个最后分片标志类型
type TThostFtdcLastFragmentType uint8

//go:generate stringer -type TThostFtdcLastFragmentType -linecomment
const (
	THOST_FTDC_LF_Yes TThostFtdcLastFragmentType = '0' // 是最后分片
	THOST_FTDC_LF_No  TThostFtdcLastFragmentType = '1' // 不是最后分片
)

// TFtdcBankAccStatusType是一个银行账户状态类型
type TThostFtdcBankAccStatusType uint8

//go:generate stringer -type TThostFtdcBankAccStatusType -linecomment
const (
	THOST_FTDC_BAS_Normal     TThostFtdcBankAccStatusType = '0' // 正常
	THOST_FTDC_BAS_Freeze     TThostFtdcBankAccStatusType = '1' // 冻结
	THOST_FTDC_BAS_ReportLoss TThostFtdcBankAccStatusType = '2' // 挂失
)

// TFtdcMoneyAccountStatusType是一个资金账户状态类型
type TThostFtdcMoneyAccountStatusType uint8

//go:generate stringer -type TThostFtdcMoneyAccountStatusType -linecomment
const (
	THOST_FTDC_MAS_Normal TThostFtdcMoneyAccountStatusType = '0' // 正常
	THOST_FTDC_MAS_Cancel TThostFtdcMoneyAccountStatusType = '1' // 销户
)

// TFtdcManageStatusType是一个存管状态类型
type TThostFtdcManageStatusType uint8

//go:generate stringer -type TThostFtdcManageStatusType -linecomment
const (
	THOST_FTDC_MSS_Point       TThostFtdcManageStatusType = '0' // 指定存管
	THOST_FTDC_MSS_PrePoint    TThostFtdcManageStatusType = '1' // 预指定
	THOST_FTDC_MSS_CancelPoint TThostFtdcManageStatusType = '2' // 撤销指定
)

// TFtdcSystemTypeType是一个应用系统类型类型
type TThostFtdcSystemTypeType uint8

//go:generate stringer -type TThostFtdcSystemTypeType -linecomment
const (
	THOST_FTDC_SYT_FutureBankTransfer TThostFtdcSystemTypeType = '0' // 银期转帐
	THOST_FTDC_SYT_StockBankTransfer  TThostFtdcSystemTypeType = '1' // 银证转帐
	THOST_FTDC_SYT_TheThirdPartStore  TThostFtdcSystemTypeType = '2' // 第三方存管
)

// TFtdcTxnEndFlagType是一个银期转帐划转结果标志类型
type TThostFtdcTxnEndFlagType uint8

//go:generate stringer -type TThostFtdcTxnEndFlagType -linecomment
const (
	THOST_FTDC_TEF_NormalProcessing             TThostFtdcTxnEndFlagType = '0' // 正常处理中
	THOST_FTDC_TEF_Success                      TThostFtdcTxnEndFlagType = '1' // 成功结束
	THOST_FTDC_TEF_Failed                       TThostFtdcTxnEndFlagType = '2' // 失败结束
	THOST_FTDC_TEF_Abnormal                     TThostFtdcTxnEndFlagType = '3' // 异常中
	THOST_FTDC_TEF_ManualProcessedForException  TThostFtdcTxnEndFlagType = '4' // 已人工异常处理
	THOST_FTDC_TEF_CommuFailedNeedManualProcess TThostFtdcTxnEndFlagType = '5' // 通讯异常 ，请人工处理
	THOST_FTDC_TEF_SysErrorNeedManualProcess    TThostFtdcTxnEndFlagType = '6' // 系统出错，请人工处理
)

// TFtdcProcessStatusType是一个银期转帐服务处理状态类型
type TThostFtdcProcessStatusType uint8

//go:generate stringer -type TThostFtdcProcessStatusType -linecomment
const (
	THOST_FTDC_PSS_NotProcess   TThostFtdcProcessStatusType = '0' // 未处理
	THOST_FTDC_PSS_StartProcess TThostFtdcProcessStatusType = '1' // 开始处理
	THOST_FTDC_PSS_Finished     TThostFtdcProcessStatusType = '2' // 处理完成
)

// TFtdcCustTypeType是一个客户类型类型
type TThostFtdcCustTypeType uint8

//go:generate stringer -type TThostFtdcCustTypeType -linecomment
const (
	THOST_FTDC_CUSTT_Person      TThostFtdcCustTypeType = '0' // 自然人
	THOST_FTDC_CUSTT_Institution TThostFtdcCustTypeType = '1' // 机构户
)

// TFtdcFBTTransferDirectionType是一个银期转帐方向类型
type TThostFtdcFBTTransferDirectionType uint8

//go:generate stringer -type TThostFtdcFBTTransferDirectionType -linecomment
const (
	THOST_FTDC_FBTTD_FromBankToFuture TThostFtdcFBTTransferDirectionType = '1' // 入金，银行转期货
	THOST_FTDC_FBTTD_FromFutureToBank TThostFtdcFBTTransferDirectionType = '2' // 出金，期货转银行
)

// TFtdcOpenOrDestroyType是一个开销户类别类型
type TThostFtdcOpenOrDestroyType uint8

//go:generate stringer -type TThostFtdcOpenOrDestroyType -linecomment
const (
	THOST_FTDC_OOD_Open    TThostFtdcOpenOrDestroyType = '1' // 开户
	THOST_FTDC_OOD_Destroy TThostFtdcOpenOrDestroyType = '0' // 销户
)

// TFtdcAvailabilityFlagType是一个有效标志类型
type TThostFtdcAvailabilityFlagType uint8

//go:generate stringer -type TThostFtdcAvailabilityFlagType -linecomment
const (
	THOST_FTDC_AVAF_Invalid TThostFtdcAvailabilityFlagType = '0' // 未确认
	THOST_FTDC_AVAF_Valid   TThostFtdcAvailabilityFlagType = '1' // 有效
	THOST_FTDC_AVAF_Repeal  TThostFtdcAvailabilityFlagType = '2' // 冲正
)

// TFtdcOrganTypeType是一个机构类型类型
type TThostFtdcOrganTypeType uint8

//go:generate stringer -type TThostFtdcOrganTypeType -linecomment
const (
	THOST_FTDC_OT_Bank           TThostFtdcOrganTypeType = '1' // 银行代理
	THOST_FTDC_OT_Future         TThostFtdcOrganTypeType = '2' // 交易前置
	THOST_FTDC_OT_PlateForm      TThostFtdcOrganTypeType = '9' // 银期转帐平台管理
	THOST_FTDC_OT_OPT_OFFSET     TThostFtdcOrganTypeType = '0' // 期权对冲
	THOST_FTDC_OT_FUT_OFFSET     TThostFtdcOrganTypeType = '1' // 期货对冲
	THOST_FTDC_OT_EXEC_OFFSET    TThostFtdcOrganTypeType = '2' // 行权后期货对冲
	THOST_FTDC_OT_PERFORM_OFFSET TThostFtdcOrganTypeType = '3' // 履约后期货对冲
)

// TFtdcOrganLevelType是一个机构级别类型
type TThostFtdcOrganLevelType uint8

//go:generate stringer -type TThostFtdcOrganLevelType -linecomment
const (
	THOST_FTDC_OL_HeadQuarters TThostFtdcOrganLevelType = '1' // 银行总行或期商总部
	THOST_FTDC_OL_Branch       TThostFtdcOrganLevelType = '2' // 银行分中心或期货公司营业部
)

// TFtdcProtocalIDType是一个协议类型类型
type TThostFtdcProtocalIDType uint8

//go:generate stringer -type TThostFtdcProtocalIDType -linecomment
const (
	THOST_FTDC_PID_FutureProtocal       TThostFtdcProtocalIDType = '0' // 期商协议
	THOST_FTDC_PID_ICBCProtocal         TThostFtdcProtocalIDType = '1' // 工行协议
	THOST_FTDC_PID_ABCProtocal          TThostFtdcProtocalIDType = '2' // 农行协议
	THOST_FTDC_PID_CBCProtocal          TThostFtdcProtocalIDType = '3' // 中国银行协议
	THOST_FTDC_PID_CCBProtocal          TThostFtdcProtocalIDType = '4' // 建行协议
	THOST_FTDC_PID_BOCOMProtocal        TThostFtdcProtocalIDType = '5' // 交行协议
	THOST_FTDC_PID_FBTPlateFormProtocal TThostFtdcProtocalIDType = 'X' // 银期转帐平台协议
)

// TFtdcConnectModeType是一个套接字连接方式类型
type TThostFtdcConnectModeType uint8

//go:generate stringer -type TThostFtdcConnectModeType -linecomment
const (
	THOST_FTDC_CM_ShortConnect TThostFtdcConnectModeType = '0' // 短连接
	THOST_FTDC_CM_LongConnect  TThostFtdcConnectModeType = '1' // 长连接
)

// TFtdcSyncModeType是一个套接字通信方式类型
type TThostFtdcSyncModeType uint8

//go:generate stringer -type TThostFtdcSyncModeType -linecomment
const (
	THOST_FTDC_SRM_ASync TThostFtdcSyncModeType = '0' // 异步
	THOST_FTDC_SRM_Sync  TThostFtdcSyncModeType = '1' // 同步
)

// TFtdcBankAccTypeType是一个银行帐号类型类型
type TThostFtdcBankAccTypeType uint8

//go:generate stringer -type TThostFtdcBankAccTypeType -linecomment
const (
	THOST_FTDC_BAT_BankBook   TThostFtdcBankAccTypeType = '1' // 银行存折
	THOST_FTDC_BAT_SavingCard TThostFtdcBankAccTypeType = '2' // 储蓄卡
	THOST_FTDC_BAT_CreditCard TThostFtdcBankAccTypeType = '3' // 信用卡
)

// TFtdcFutureAccTypeType是一个期货公司帐号类型类型
type TThostFtdcFutureAccTypeType uint8

//go:generate stringer -type TThostFtdcFutureAccTypeType -linecomment
const (
	THOST_FTDC_FAT_BankBook   TThostFtdcFutureAccTypeType = '1' // 银行存折
	THOST_FTDC_FAT_SavingCard TThostFtdcFutureAccTypeType = '2' // 储蓄卡
	THOST_FTDC_FAT_CreditCard TThostFtdcFutureAccTypeType = '3' // 信用卡
)

// TFtdcOrganStatusType是一个接入机构状态类型
type TThostFtdcOrganStatusType uint8

//go:generate stringer -type TThostFtdcOrganStatusType -linecomment
const (
	THOST_FTDC_OS_Ready            TThostFtdcOrganStatusType = '0' // 启用
	THOST_FTDC_OS_CheckIn          TThostFtdcOrganStatusType = '1' // 签到
	THOST_FTDC_OS_CheckOut         TThostFtdcOrganStatusType = '2' // 签退
	THOST_FTDC_OS_CheckFileArrived TThostFtdcOrganStatusType = '3' // 对帐文件到达
	THOST_FTDC_OS_CheckDetail      TThostFtdcOrganStatusType = '4' // 对帐
	THOST_FTDC_OS_DayEndClean      TThostFtdcOrganStatusType = '5' // 日终清理
	THOST_FTDC_OS_Invalid          TThostFtdcOrganStatusType = '9' // 注销
)

// TFtdcCCBFeeModeType是一个建行收费模式类型
type TThostFtdcCCBFeeModeType uint8

//go:generate stringer -type TThostFtdcCCBFeeModeType -linecomment
const (
	THOST_FTDC_CCBFM_ByAmount TThostFtdcCCBFeeModeType = '1' // 按金额扣收
	THOST_FTDC_CCBFM_ByMonth  TThostFtdcCCBFeeModeType = '2' // 按月扣收
)

// TFtdcCommApiTypeType是一个通讯API类型类型
type TThostFtdcCommApiTypeType uint8

//go:generate stringer -type TThostFtdcCommApiTypeType -linecomment
const (
	THOST_FTDC_CAPIT_Client  TThostFtdcCommApiTypeType = '1' // 客户端
	THOST_FTDC_CAPIT_Server  TThostFtdcCommApiTypeType = '2' // 服务端
	THOST_FTDC_CAPIT_UserApi TThostFtdcCommApiTypeType = '3' // 交易系统的UserApi
)

// TFtdcServiceIDType是一个服务编号类型
type TThostFtdcServiceIDType int32

// TFtdcServiceLineNoType是一个服务线路编号类型
type TThostFtdcServiceLineNoType int32

// TFtdcServiceNameType是一个服务名类型
type TThostFtdcServiceNameType [61]byte

func (t TThostFtdcServiceNameType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcServiceNameType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcLinkStatusType是一个连接状态类型
type TThostFtdcLinkStatusType uint8

//go:generate stringer -type TThostFtdcLinkStatusType -linecomment
const (
	THOST_FTDC_LS_Connected    TThostFtdcLinkStatusType = '1' // 已经连接
	THOST_FTDC_LS_Disconnected TThostFtdcLinkStatusType = '2' // 没有连接
)

// TFtdcCommApiPointerType是一个通讯API指针类型
type TThostFtdcCommApiPointerType int32

// TFtdcPwdFlagType是一个密码核对标志类型
type TThostFtdcPwdFlagType uint8

//go:generate stringer -type TThostFtdcPwdFlagType -linecomment
const (
	THOST_FTDC_BPWDF_NoCheck      TThostFtdcPwdFlagType = '0' // 不核对
	THOST_FTDC_BPWDF_BlankCheck   TThostFtdcPwdFlagType = '1' // 明文核对
	THOST_FTDC_BPWDF_EncryptCheck TThostFtdcPwdFlagType = '2' // 密文核对
)

// TFtdcSecuAccTypeType是一个期货帐号类型类型
type TThostFtdcSecuAccTypeType uint8

//go:generate stringer -type TThostFtdcSecuAccTypeType -linecomment
const (
	THOST_FTDC_SAT_AccountID       TThostFtdcSecuAccTypeType = '1' // 资金帐号
	THOST_FTDC_SAT_CardID          TThostFtdcSecuAccTypeType = '2' // 资金卡号
	THOST_FTDC_SAT_SHStockholderID TThostFtdcSecuAccTypeType = '3' // 上海股东帐号
	THOST_FTDC_SAT_SZStockholderID TThostFtdcSecuAccTypeType = '4' // 深圳股东帐号
)

// TFtdcTransferStatusType是一个转账交易状态类型
type TThostFtdcTransferStatusType uint8

//go:generate stringer -type TThostFtdcTransferStatusType -linecomment
const (
	THOST_FTDC_TRFS_Normal   TThostFtdcTransferStatusType = '0' // 正常
	THOST_FTDC_TRFS_Repealed TThostFtdcTransferStatusType = '1' // 被冲正
)

// TFtdcSponsorTypeType是一个发起方类型
type TThostFtdcSponsorTypeType uint8

//go:generate stringer -type TThostFtdcSponsorTypeType -linecomment
const (
	THOST_FTDC_SPTYPE_Broker TThostFtdcSponsorTypeType = '0' // 期商
	THOST_FTDC_SPTYPE_Bank   TThostFtdcSponsorTypeType = '1' // 银行
)

// TFtdcReqRspTypeType是一个请求响应类别类型
type TThostFtdcReqRspTypeType uint8

//go:generate stringer -type TThostFtdcReqRspTypeType -linecomment
const (
	THOST_FTDC_REQRSP_Request  TThostFtdcReqRspTypeType = '0' // 请求
	THOST_FTDC_REQRSP_Response TThostFtdcReqRspTypeType = '1' // 响应
)

// TFtdcFBTUserEventTypeType是一个银期转帐用户事件类型类型
type TThostFtdcFBTUserEventTypeType uint8

//go:generate stringer -type TThostFtdcFBTUserEventTypeType -linecomment
const (
	THOST_FTDC_FBTUET_SignIn                    TThostFtdcFBTUserEventTypeType = '0' // 签到
	THOST_FTDC_FBTUET_FromBankToFuture          TThostFtdcFBTUserEventTypeType = '1' // 银行转期货
	THOST_FTDC_FBTUET_FromFutureToBank          TThostFtdcFBTUserEventTypeType = '2' // 期货转银行
	THOST_FTDC_FBTUET_OpenAccount               TThostFtdcFBTUserEventTypeType = '3' // 开户
	THOST_FTDC_FBTUET_CancelAccount             TThostFtdcFBTUserEventTypeType = '4' // 销户
	THOST_FTDC_FBTUET_ChangeAccount             TThostFtdcFBTUserEventTypeType = '5' // 变更银行账户
	THOST_FTDC_FBTUET_RepealFromBankToFuture    TThostFtdcFBTUserEventTypeType = '6' // 冲正银行转期货
	THOST_FTDC_FBTUET_RepealFromFutureToBank    TThostFtdcFBTUserEventTypeType = '7' // 冲正期货转银行
	THOST_FTDC_FBTUET_QueryBankAccount          TThostFtdcFBTUserEventTypeType = '8' // 查询银行账户
	THOST_FTDC_FBTUET_QueryFutureAccount        TThostFtdcFBTUserEventTypeType = '9' // 查询期货账户
	THOST_FTDC_FBTUET_SignOut                   TThostFtdcFBTUserEventTypeType = 'A' // 签退
	THOST_FTDC_FBTUET_SyncKey                   TThostFtdcFBTUserEventTypeType = 'B' // 密钥同步
	THOST_FTDC_FBTUET_ReserveOpenAccount        TThostFtdcFBTUserEventTypeType = 'C' // 预约开户
	THOST_FTDC_FBTUET_CancelReserveOpenAccount  TThostFtdcFBTUserEventTypeType = 'D' // 撤销预约开户
	THOST_FTDC_FBTUET_ReserveOpenAccountConfirm TThostFtdcFBTUserEventTypeType = 'E' // 预约开户确认
	THOST_FTDC_FBTUET_Other                     TThostFtdcFBTUserEventTypeType = 'Z' // 其他
)

// TFtdcBankIDByBankType是一个银行自己的编码类型
type TThostFtdcBankIDByBankType [21]byte

func (t TThostFtdcBankIDByBankType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcBankIDByBankType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcBankOperNoType是一个银行操作员号类型
type TThostFtdcBankOperNoType [4]byte

func (t TThostFtdcBankOperNoType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcBankOperNoType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcBankCustNoType是一个银行客户号类型
type TThostFtdcBankCustNoType [21]byte

func (t TThostFtdcBankCustNoType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcBankCustNoType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcDBOPSeqNoType是一个递增的序列号类型
type TThostFtdcDBOPSeqNoType int32

// TFtdcTableNameType是一个FBT表名类型
type TThostFtdcTableNameType [61]byte

func (t TThostFtdcTableNameType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcTableNameType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcPKNameType是一个FBT表操作主键名类型
type TThostFtdcPKNameType [201]byte

func (t TThostFtdcPKNameType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcPKNameType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcPKValueType是一个FBT表操作主键值类型
type TThostFtdcPKValueType [501]byte

func (t TThostFtdcPKValueType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcPKValueType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcDBOperationType是一个记录操作类型类型
type TThostFtdcDBOperationType uint8

//go:generate stringer -type TThostFtdcDBOperationType -linecomment
const (
	THOST_FTDC_DBOP_Insert TThostFtdcDBOperationType = '0' // 插入
	THOST_FTDC_DBOP_Update TThostFtdcDBOperationType = '1' // 更新
	THOST_FTDC_DBOP_Delete TThostFtdcDBOperationType = '2' // 删除
)

// TFtdcSyncFlagType是一个同步标记类型
type TThostFtdcSyncFlagType uint8

//go:generate stringer -type TThostFtdcSyncFlagType -linecomment
const (
	THOST_FTDC_SYNF_Yes TThostFtdcSyncFlagType = '0' // 已同步
	THOST_FTDC_SYNF_No  TThostFtdcSyncFlagType = '1' // 未同步
)

// TFtdcTargetIDType是一个同步目标编号类型
type TThostFtdcTargetIDType [4]byte

func (t TThostFtdcTargetIDType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcTargetIDType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcSyncTypeType是一个同步类型类型
type TThostFtdcSyncTypeType uint8

//go:generate stringer -type TThostFtdcSyncTypeType -linecomment
const (
	THOST_FTDC_SYNT_OneOffSync    TThostFtdcSyncTypeType = '0' // 一次同步
	THOST_FTDC_SYNT_TimerSync     TThostFtdcSyncTypeType = '1' // 定时同步
	THOST_FTDC_SYNT_TimerFullSync TThostFtdcSyncTypeType = '2' // 定时完全同步
)

// TFtdcFBETimeType是一个各种换汇时间类型
type TThostFtdcFBETimeType [7]byte

func (t TThostFtdcFBETimeType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcFBETimeType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcFBEBankNoType是一个换汇银行行号类型
type TThostFtdcFBEBankNoType [13]byte

func (t TThostFtdcFBEBankNoType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcFBEBankNoType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcFBECertNoType是一个换汇凭证号类型
type TThostFtdcFBECertNoType [13]byte

func (t TThostFtdcFBECertNoType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcFBECertNoType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcExDirectionType是一个换汇方向类型
type TThostFtdcExDirectionType uint8

//go:generate stringer -type TThostFtdcExDirectionType -linecomment
const (
	THOST_FTDC_FBEDIR_Settlement TThostFtdcExDirectionType = '0' // 结汇
	THOST_FTDC_FBEDIR_Sale       TThostFtdcExDirectionType = '1' // 售汇
)

// TFtdcFBEBankAccountType是一个换汇银行账户类型
type TThostFtdcFBEBankAccountType [33]byte

func (t TThostFtdcFBEBankAccountType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcFBEBankAccountType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcFBEBankAccountNameType是一个换汇银行账户名类型
type TThostFtdcFBEBankAccountNameType [61]byte

func (t TThostFtdcFBEBankAccountNameType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcFBEBankAccountNameType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcFBEAmtType是一个各种换汇金额类型
type TThostFtdcFBEAmtType float64

func (t TThostFtdcFBEAmtType) IsValid() bool {
	return float64(t) != math.MaxFloat64
}

func (t TThostFtdcFBEAmtType) String() string {
	if t.IsValid() {
		return strconv.FormatFloat(float64(t), 'f', 6, 64)
	} else {
		return "DBL_MAX"
	}
}

// TFtdcFBEBusinessTypeType是一个换汇业务类型类型
type TThostFtdcFBEBusinessTypeType [3]byte

func (t TThostFtdcFBEBusinessTypeType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcFBEBusinessTypeType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcFBEPostScriptType是一个换汇附言类型
type TThostFtdcFBEPostScriptType [61]byte

func (t TThostFtdcFBEPostScriptType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcFBEPostScriptType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcFBERemarkType是一个换汇备注类型
type TThostFtdcFBERemarkType [71]byte

func (t TThostFtdcFBERemarkType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcFBERemarkType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcExRateType是一个换汇汇率类型
type TThostFtdcExRateType float64

func (t TThostFtdcExRateType) IsValid() bool {
	return float64(t) != math.MaxFloat64
}

func (t TThostFtdcExRateType) String() string {
	if t.IsValid() {
		return strconv.FormatFloat(float64(t), 'f', 6, 64)
	} else {
		return "DBL_MAX"
	}
}

// TFtdcFBEResultFlagType是一个换汇成功标志类型
type TThostFtdcFBEResultFlagType uint8

//go:generate stringer -type TThostFtdcFBEResultFlagType -linecomment
const (
	THOST_FTDC_FBERES_Success             TThostFtdcFBEResultFlagType = '0' // 成功
	THOST_FTDC_FBERES_InsufficientBalance TThostFtdcFBEResultFlagType = '1' // 账户余额不足
	THOST_FTDC_FBERES_UnknownTrading      TThostFtdcFBEResultFlagType = '8' // 交易结果未知
	THOST_FTDC_FBERES_Fail                TThostFtdcFBEResultFlagType = 'x' // 失败
)

// TFtdcFBERtnMsgType是一个换汇返回信息类型
type TThostFtdcFBERtnMsgType [61]byte

func (t TThostFtdcFBERtnMsgType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcFBERtnMsgType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcFBEExtendMsgType是一个换汇扩展信息类型
type TThostFtdcFBEExtendMsgType [61]byte

func (t TThostFtdcFBEExtendMsgType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcFBEExtendMsgType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcFBEBusinessSerialType是一个换汇记账流水号类型
type TThostFtdcFBEBusinessSerialType [31]byte

func (t TThostFtdcFBEBusinessSerialType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcFBEBusinessSerialType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcFBESystemSerialType是一个换汇流水号类型
type TThostFtdcFBESystemSerialType [21]byte

func (t TThostFtdcFBESystemSerialType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcFBESystemSerialType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcFBETotalExCntType是一个换汇交易总笔数类型
type TThostFtdcFBETotalExCntType int32

// TFtdcFBEExchStatusType是一个换汇交易状态类型
type TThostFtdcFBEExchStatusType uint8

//go:generate stringer -type TThostFtdcFBEExchStatusType -linecomment
const (
	THOST_FTDC_FBEES_Normal     TThostFtdcFBEExchStatusType = '0' // 正常
	THOST_FTDC_FBEES_ReExchange TThostFtdcFBEExchStatusType = '1' // 交易重发
)

// TFtdcFBEFileFlagType是一个换汇文件标志类型
type TThostFtdcFBEFileFlagType uint8

//go:generate stringer -type TThostFtdcFBEFileFlagType -linecomment
const (
	THOST_FTDC_FBEFG_DataPackage TThostFtdcFBEFileFlagType = '0' // 数据包
	THOST_FTDC_FBEFG_File        TThostFtdcFBEFileFlagType = '1' // 文件
)

// TFtdcFBEAlreadyTradeType是一个换汇已交易标志类型
type TThostFtdcFBEAlreadyTradeType uint8

//go:generate stringer -type TThostFtdcFBEAlreadyTradeType -linecomment
const (
	THOST_FTDC_FBEAT_NotTrade TThostFtdcFBEAlreadyTradeType = '0' // 未交易
	THOST_FTDC_FBEAT_Trade    TThostFtdcFBEAlreadyTradeType = '1' // 已交易
)

// TFtdcFBEOpenBankType是一个换汇账户开户行类型
type TThostFtdcFBEOpenBankType [61]byte

func (t TThostFtdcFBEOpenBankType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcFBEOpenBankType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcFBEUserEventTypeType是一个银期换汇用户事件类型类型
type TThostFtdcFBEUserEventTypeType uint8

//go:generate stringer -type TThostFtdcFBEUserEventTypeType -linecomment
const (
	THOST_FTDC_FBEUET_SignIn           TThostFtdcFBEUserEventTypeType = '0' // 签到
	THOST_FTDC_FBEUET_Exchange         TThostFtdcFBEUserEventTypeType = '1' // 换汇
	THOST_FTDC_FBEUET_ReExchange       TThostFtdcFBEUserEventTypeType = '2' // 换汇重发
	THOST_FTDC_FBEUET_QueryBankAccount TThostFtdcFBEUserEventTypeType = '3' // 银行账户查询
	THOST_FTDC_FBEUET_QueryExchDetial  TThostFtdcFBEUserEventTypeType = '4' // 换汇明细查询
	THOST_FTDC_FBEUET_QueryExchSummary TThostFtdcFBEUserEventTypeType = '5' // 换汇汇总查询
	THOST_FTDC_FBEUET_QueryExchRate    TThostFtdcFBEUserEventTypeType = '6' // 换汇汇率查询
	THOST_FTDC_FBEUET_CheckBankAccount TThostFtdcFBEUserEventTypeType = '7' // 对账文件通知
	THOST_FTDC_FBEUET_SignOut          TThostFtdcFBEUserEventTypeType = '8' // 签退
	THOST_FTDC_FBEUET_Other            TThostFtdcFBEUserEventTypeType = 'Z' // 其他
)

// TFtdcFBEFileNameType是一个换汇相关文件名类型
type TThostFtdcFBEFileNameType [21]byte

func (t TThostFtdcFBEFileNameType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcFBEFileNameType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcFBEBatchSerialType是一个换汇批次号类型
type TThostFtdcFBEBatchSerialType [21]byte

func (t TThostFtdcFBEBatchSerialType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcFBEBatchSerialType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcFBEReqFlagType是一个换汇发送标志类型
type TThostFtdcFBEReqFlagType uint8

//go:generate stringer -type TThostFtdcFBEReqFlagType -linecomment
const (
	THOST_FTDC_FBERF_UnProcessed TThostFtdcFBEReqFlagType = '0' // 未处理
	THOST_FTDC_FBERF_WaitSend    TThostFtdcFBEReqFlagType = '1' // 等待发送
	THOST_FTDC_FBERF_SendSuccess TThostFtdcFBEReqFlagType = '2' // 发送成功
	THOST_FTDC_FBERF_SendFailed  TThostFtdcFBEReqFlagType = '3' // 发送失败
	THOST_FTDC_FBERF_WaitReSend  TThostFtdcFBEReqFlagType = '4' // 等待重发
)

// TFtdcNotifyClassType是一个风险通知类型类型
type TThostFtdcNotifyClassType uint8

//go:generate stringer -type TThostFtdcNotifyClassType -linecomment
const (
	THOST_FTDC_NC_NOERROR   TThostFtdcNotifyClassType = '0' // 正常
	THOST_FTDC_NC_Warn      TThostFtdcNotifyClassType = '1' // 警示
	THOST_FTDC_NC_Call      TThostFtdcNotifyClassType = '2' // 追保
	THOST_FTDC_NC_Force     TThostFtdcNotifyClassType = '3' // 强平
	THOST_FTDC_NC_CHUANCANG TThostFtdcNotifyClassType = '4' // 穿仓
	THOST_FTDC_NC_Exception TThostFtdcNotifyClassType = '5' // 异常
)

// TFtdcRiskNofityInfoType是一个客户风险通知消息类型
type TThostFtdcRiskNofityInfoType [257]byte

func (t TThostFtdcRiskNofityInfoType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcRiskNofityInfoType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcForceCloseSceneIdType是一个强平场景编号类型
type TThostFtdcForceCloseSceneIdType [24]byte

func (t TThostFtdcForceCloseSceneIdType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcForceCloseSceneIdType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcForceCloseTypeType是一个强平单类型类型
type TThostFtdcForceCloseTypeType uint8

//go:generate stringer -type TThostFtdcForceCloseTypeType -linecomment
const (
	THOST_FTDC_FCT_Manual TThostFtdcForceCloseTypeType = '0' // 手工强平
	THOST_FTDC_FCT_Single TThostFtdcForceCloseTypeType = '1' // 单一投资者辅助强平
	THOST_FTDC_FCT_Group  TThostFtdcForceCloseTypeType = '2' // 批量投资者辅助强平
)

// TFtdcInstrumentIDsType是一个多个产品代码,用+分隔,如cu+zn类型
type TThostFtdcInstrumentIDsType [101]byte

func (t TThostFtdcInstrumentIDsType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcInstrumentIDsType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcRiskNotifyMethodType是一个风险通知途径类型
type TThostFtdcRiskNotifyMethodType uint8

//go:generate stringer -type TThostFtdcRiskNotifyMethodType -linecomment
const (
	THOST_FTDC_RNM_System TThostFtdcRiskNotifyMethodType = '0' // 系统通知
	THOST_FTDC_RNM_SMS    TThostFtdcRiskNotifyMethodType = '1' // 短信通知
	THOST_FTDC_RNM_EMail  TThostFtdcRiskNotifyMethodType = '2' // 邮件通知
	THOST_FTDC_RNM_Manual TThostFtdcRiskNotifyMethodType = '3' // 人工通知
)

// TFtdcRiskNotifyStatusType是一个风险通知状态类型
type TThostFtdcRiskNotifyStatusType uint8

//go:generate stringer -type TThostFtdcRiskNotifyStatusType -linecomment
const (
	THOST_FTDC_RNS_NotGen    TThostFtdcRiskNotifyStatusType = '0' // 未生成
	THOST_FTDC_RNS_Generated TThostFtdcRiskNotifyStatusType = '1' // 已生成未发送
	THOST_FTDC_RNS_SendError TThostFtdcRiskNotifyStatusType = '2' // 发送失败
	THOST_FTDC_RNS_SendOk    TThostFtdcRiskNotifyStatusType = '3' // 已发送未接收
	THOST_FTDC_RNS_Received  TThostFtdcRiskNotifyStatusType = '4' // 已接收未确认
	THOST_FTDC_RNS_Confirmed TThostFtdcRiskNotifyStatusType = '5' // 已确认
)

// TFtdcRiskUserEventType是一个风控用户操作事件类型
type TThostFtdcRiskUserEventType uint8

//go:generate stringer -type TThostFtdcRiskUserEventType -linecomment
const (
	THOST_FTDC_RUE_ExportData TThostFtdcRiskUserEventType = '0' // 导出数据
)

// TFtdcParamIDType是一个参数代码类型
type TThostFtdcParamIDType int32

// TFtdcParamNameType是一个参数名类型
type TThostFtdcParamNameType [41]byte

func (t TThostFtdcParamNameType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcParamNameType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcParamValueType是一个参数值类型
type TThostFtdcParamValueType [41]byte

func (t TThostFtdcParamValueType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcParamValueType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcConditionalOrderSortTypeType是一个条件单索引条件类型
type TThostFtdcConditionalOrderSortTypeType uint8

//go:generate stringer -type TThostFtdcConditionalOrderSortTypeType -linecomment
const (
	THOST_FTDC_COST_LastPriceAsc  TThostFtdcConditionalOrderSortTypeType = '0' // 使用最新价升序
	THOST_FTDC_COST_LastPriceDesc TThostFtdcConditionalOrderSortTypeType = '1' // 使用最新价降序
	THOST_FTDC_COST_AskPriceAsc   TThostFtdcConditionalOrderSortTypeType = '2' // 使用卖价升序
	THOST_FTDC_COST_AskPriceDesc  TThostFtdcConditionalOrderSortTypeType = '3' // 使用卖价降序
	THOST_FTDC_COST_BidPriceAsc   TThostFtdcConditionalOrderSortTypeType = '4' // 使用买价升序
	THOST_FTDC_COST_BidPriceDesc  TThostFtdcConditionalOrderSortTypeType = '5' // 使用买价降序
)

// TFtdcSendTypeType是一个报送状态类型
type TThostFtdcSendTypeType uint8

//go:generate stringer -type TThostFtdcSendTypeType -linecomment
const (
	THOST_FTDC_UOAST_NoSend    TThostFtdcSendTypeType = '0' // 未发送
	THOST_FTDC_UOAST_Sended    TThostFtdcSendTypeType = '1' // 已发送
	THOST_FTDC_UOAST_Generated TThostFtdcSendTypeType = '2' // 已生成
	THOST_FTDC_UOAST_SendFail  TThostFtdcSendTypeType = '3' // 报送失败
	THOST_FTDC_UOAST_Success   TThostFtdcSendTypeType = '4' // 接收成功
	THOST_FTDC_UOAST_Fail      TThostFtdcSendTypeType = '5' // 接收失败
	THOST_FTDC_UOAST_Cancel    TThostFtdcSendTypeType = '6' // 取消报送
)

// TFtdcClientIDStatusType是一个交易编码状态类型
type TThostFtdcClientIDStatusType uint8

//go:generate stringer -type TThostFtdcClientIDStatusType -linecomment
const (
	THOST_FTDC_UOACS_NoApply  TThostFtdcClientIDStatusType = '1' // 未申请
	THOST_FTDC_UOACS_Submited TThostFtdcClientIDStatusType = '2' // 已提交申请
	THOST_FTDC_UOACS_Sended   TThostFtdcClientIDStatusType = '3' // 已发送申请
	THOST_FTDC_UOACS_Success  TThostFtdcClientIDStatusType = '4' // 完成
	THOST_FTDC_UOACS_Refuse   TThostFtdcClientIDStatusType = '5' // 拒绝
	THOST_FTDC_UOACS_Cancel   TThostFtdcClientIDStatusType = '6' // 已撤销编码
)

// TFtdcIndustryIDType是一个行业编码类型
type TThostFtdcIndustryIDType [17]byte

func (t TThostFtdcIndustryIDType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcIndustryIDType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcQuestionIDType是一个特有信息编号类型
type TThostFtdcQuestionIDType [5]byte

func (t TThostFtdcQuestionIDType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcQuestionIDType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcQuestionContentType是一个特有信息说明类型
type TThostFtdcQuestionContentType [41]byte

func (t TThostFtdcQuestionContentType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcQuestionContentType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcOptionIDType是一个选项编号类型
type TThostFtdcOptionIDType [13]byte

func (t TThostFtdcOptionIDType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcOptionIDType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcOptionContentType是一个选项说明类型
type TThostFtdcOptionContentType [61]byte

func (t TThostFtdcOptionContentType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcOptionContentType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcQuestionTypeType是一个特有信息类型类型
type TThostFtdcQuestionTypeType uint8

//go:generate stringer -type TThostFtdcQuestionTypeType -linecomment
const (
	THOST_FTDC_QT_Radio  TThostFtdcQuestionTypeType = '1' // 单选
	THOST_FTDC_QT_Option TThostFtdcQuestionTypeType = '2' // 多选
	THOST_FTDC_QT_Blank  TThostFtdcQuestionTypeType = '3' // 填空
)

// TFtdcProcessIDType是一个业务流水号类型
type TThostFtdcProcessIDType [33]byte

func (t TThostFtdcProcessIDType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcProcessIDType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcSeqNoType是一个流水号类型
type TThostFtdcSeqNoType int32

// TFtdcUOAProcessStatusType是一个流程状态类型
type TThostFtdcUOAProcessStatusType [3]byte

func (t TThostFtdcUOAProcessStatusType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcUOAProcessStatusType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcProcessTypeType是一个流程功能类型类型
type TThostFtdcProcessTypeType [3]byte

func (t TThostFtdcProcessTypeType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcProcessTypeType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

type TThostFtdcBusinessTypeType uint8

// TFtdcCfmmcReturnCodeType是一个监控中心返回码类型
type TThostFtdcCfmmcReturnCodeType uint8

//go:generate stringer -type TThostFtdcCfmmcReturnCodeType -linecomment
const (
	THOST_FTDC_CRC_Success    TThostFtdcCfmmcReturnCodeType = '0' // 成功
	THOST_FTDC_CRC_Working    TThostFtdcCfmmcReturnCodeType = '1' // 该客户已经有流程在处理中
	THOST_FTDC_CRC_InfoFail   TThostFtdcCfmmcReturnCodeType = '2' // 监控中客户资料检查失败
	THOST_FTDC_CRC_IDCardFail TThostFtdcCfmmcReturnCodeType = '3' // 监控中实名制检查失败
	THOST_FTDC_CRC_OtherFail  TThostFtdcCfmmcReturnCodeType = '4' // 其他错误
)

// TFtdcExReturnCodeType是一个交易所返回码类型
type TThostFtdcExReturnCodeType int32

// TFtdcClientTypeType是一个客户类型类型
type TThostFtdcClientTypeType uint8

//go:generate stringer -type TThostFtdcClientTypeType -linecomment
const (
	THOST_FTDC_CfMMCCT_All          TThostFtdcClientTypeType = '0' // 所有
	THOST_FTDC_CfMMCCT_Person       TThostFtdcClientTypeType = '1' // 个人
	THOST_FTDC_CfMMCCT_Company      TThostFtdcClientTypeType = '2' // 单位
	THOST_FTDC_CfMMCCT_Other        TThostFtdcClientTypeType = '3' // 其他
	THOST_FTDC_CfMMCCT_SpecialOrgan TThostFtdcClientTypeType = '4' // 特殊法人
	THOST_FTDC_CfMMCCT_Asset        TThostFtdcClientTypeType = '5' // 资管户
)

// TFtdcExchangeIDTypeType是一个交易所编号类型
type TThostFtdcExchangeIDTypeType uint8

//go:generate stringer -type TThostFtdcExchangeIDTypeType -linecomment
const (
	THOST_FTDC_EIDT_SHFE  TThostFtdcExchangeIDTypeType = 'S' // 上海期货交易所
	THOST_FTDC_EIDT_CZCE  TThostFtdcExchangeIDTypeType = 'Z' // 郑州商品交易所
	THOST_FTDC_EIDT_DCE   TThostFtdcExchangeIDTypeType = 'D' // 大连商品交易所
	THOST_FTDC_EIDT_CFFEX TThostFtdcExchangeIDTypeType = 'J' // 中国金融期货交易所
	THOST_FTDC_EIDT_INE   TThostFtdcExchangeIDTypeType = 'N' // 上海国际能源交易中心股份有限公司
)

// TFtdcExClientIDTypeType是一个交易编码类型类型
type TThostFtdcExClientIDTypeType uint8

//go:generate stringer -type TThostFtdcExClientIDTypeType -linecomment
const (
	THOST_FTDC_ECIDT_Hedge       TThostFtdcExClientIDTypeType = '1' // 套保
	THOST_FTDC_ECIDT_Arbitrage   TThostFtdcExClientIDTypeType = '2' // 套利
	THOST_FTDC_ECIDT_Speculation TThostFtdcExClientIDTypeType = '3' // 投机
)

// TFtdcClientClassifyType是一个客户分类码类型
type TThostFtdcClientClassifyType [11]byte

func (t TThostFtdcClientClassifyType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcClientClassifyType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcUOAOrganTypeType是一个单位性质类型
type TThostFtdcUOAOrganTypeType [11]byte

func (t TThostFtdcUOAOrganTypeType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcUOAOrganTypeType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcUOACountryCodeType是一个国家代码类型
type TThostFtdcUOACountryCodeType [11]byte

func (t TThostFtdcUOACountryCodeType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcUOACountryCodeType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcAreaCodeType是一个区号类型
type TThostFtdcAreaCodeType [11]byte

func (t TThostFtdcAreaCodeType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcAreaCodeType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcFuturesIDType是一个监控中心为客户分配的代码类型
type TThostFtdcFuturesIDType [21]byte

func (t TThostFtdcFuturesIDType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcFuturesIDType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcCffmcDateType是一个日期类型
type TThostFtdcCffmcDateType [11]byte

func (t TThostFtdcCffmcDateType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcCffmcDateType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcCffmcTimeType是一个时间类型
type TThostFtdcCffmcTimeType [11]byte

func (t TThostFtdcCffmcTimeType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcCffmcTimeType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcNocIDType是一个组织机构代码类型
type TThostFtdcNocIDType [21]byte

func (t TThostFtdcNocIDType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcNocIDType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcUpdateFlagType是一个更新状态类型
type TThostFtdcUpdateFlagType uint8

//go:generate stringer -type TThostFtdcUpdateFlagType -linecomment
const (
	THOST_FTDC_UF_NoUpdate  TThostFtdcUpdateFlagType = '0' // 未更新
	THOST_FTDC_UF_Success   TThostFtdcUpdateFlagType = '1' // 更新全部信息成功
	THOST_FTDC_UF_Fail      TThostFtdcUpdateFlagType = '2' // 更新全部信息失败
	THOST_FTDC_UF_TCSuccess TThostFtdcUpdateFlagType = '3' // 更新交易编码成功
	THOST_FTDC_UF_TCFail    TThostFtdcUpdateFlagType = '4' // 更新交易编码失败
	THOST_FTDC_UF_Cancel    TThostFtdcUpdateFlagType = '5' // 已丢弃
)

// TFtdcApplyOperateIDType是一个申请动作类型
type TThostFtdcApplyOperateIDType uint8

//go:generate stringer -type TThostFtdcApplyOperateIDType -linecomment
const (
	THOST_FTDC_AOID_OpenInvestor        TThostFtdcApplyOperateIDType = '1' // 开户
	THOST_FTDC_AOID_ModifyIDCard        TThostFtdcApplyOperateIDType = '2' // 修改身份信息
	THOST_FTDC_AOID_ModifyNoIDCard      TThostFtdcApplyOperateIDType = '3' // 修改一般信息
	THOST_FTDC_AOID_ApplyTradingCode    TThostFtdcApplyOperateIDType = '4' // 申请交易编码
	THOST_FTDC_AOID_CancelTradingCode   TThostFtdcApplyOperateIDType = '5' // 撤销交易编码
	THOST_FTDC_AOID_CancelInvestor      TThostFtdcApplyOperateIDType = '6' // 销户
	THOST_FTDC_AOID_FreezeAccount       TThostFtdcApplyOperateIDType = '8' // 账户休眠
	THOST_FTDC_AOID_ActiveFreezeAccount TThostFtdcApplyOperateIDType = '9' // 激活休眠账户
)

// TFtdcApplyStatusIDType是一个申请状态类型
type TThostFtdcApplyStatusIDType uint8

//go:generate stringer -type TThostFtdcApplyStatusIDType -linecomment
const (
	THOST_FTDC_ASID_NoComplete TThostFtdcApplyStatusIDType = '1' // 未补全
	THOST_FTDC_ASID_Submited   TThostFtdcApplyStatusIDType = '2' // 已提交
	THOST_FTDC_ASID_Checked    TThostFtdcApplyStatusIDType = '3' // 已审核
	THOST_FTDC_ASID_Refused    TThostFtdcApplyStatusIDType = '4' // 已拒绝
	THOST_FTDC_ASID_Deleted    TThostFtdcApplyStatusIDType = '5' // 已删除
)

// TFtdcSendMethodType是一个发送方式类型
type TThostFtdcSendMethodType uint8

//go:generate stringer -type TThostFtdcSendMethodType -linecomment
const (
	THOST_FTDC_UOASM_ByAPI  TThostFtdcSendMethodType = '1' // 文件发送
	THOST_FTDC_UOASM_ByFile TThostFtdcSendMethodType = '2' // 电子发送
)

// TFtdcEventTypeType是一个业务操作类型类型
type TThostFtdcEventTypeType [33]byte

func (t TThostFtdcEventTypeType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcEventTypeType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcEventModeType是一个操作方法类型
type TThostFtdcEventModeType uint8

//go:generate stringer -type TThostFtdcEventModeType -linecomment
const (
	THOST_FTDC_EvM_ADD                   TThostFtdcEventModeType = '1' // 增加
	THOST_FTDC_EvM_UPDATE                TThostFtdcEventModeType = '2' // 修改
	THOST_FTDC_EvM_DELETE                TThostFtdcEventModeType = '3' // 删除
	THOST_FTDC_EvM_CHECK                 TThostFtdcEventModeType = '4' // 复核
	THOST_FTDC_EvM_COPY                  TThostFtdcEventModeType = '5' // 复制
	THOST_FTDC_EvM_CANCEL                TThostFtdcEventModeType = '6' // 注销
	THOST_FTDC_EvM_Reverse               TThostFtdcEventModeType = '7' // 冲销
	THOST_FTDC_EvM_InvestorGroupFlow     TThostFtdcEventModeType = '1' // 投资者对应投资者组设置
	THOST_FTDC_EvM_InvestorRate          TThostFtdcEventModeType = '2' // 投资者手续费率设置
	THOST_FTDC_EvM_InvestorCommRateModel TThostFtdcEventModeType = '3' // 投资者手续费率模板关系设置
)

// TFtdcUOAAutoSendType是一个统一开户申请自动发送类型
type TThostFtdcUOAAutoSendType uint8

//go:generate stringer -type TThostFtdcUOAAutoSendType -linecomment
const (
	THOST_FTDC_UOAA_ASR  TThostFtdcUOAAutoSendType = '1' // 自动发送并接收
	THOST_FTDC_UOAA_ASNR TThostFtdcUOAAutoSendType = '2' // 自动发送，不自动接收
	THOST_FTDC_UOAA_NSAR TThostFtdcUOAAutoSendType = '3' // 不自动发送，自动接收
	THOST_FTDC_UOAA_NSR  TThostFtdcUOAAutoSendType = '4' // 不自动发送，也不自动接收
)

// TFtdcQueryDepthType是一个查询深度类型
type TThostFtdcQueryDepthType int32

// TFtdcDataCenterIDType是一个数据中心代码类型
type TThostFtdcDataCenterIDType int32

type TThostFtdcFlowIDType uint8

// TFtdcCheckLevelType是一个复核级别类型
type TThostFtdcCheckLevelType uint8

//go:generate stringer -type TThostFtdcCheckLevelType -linecomment
const (
	THOST_FTDC_CL_Zero TThostFtdcCheckLevelType = '0' // 零级复核
	THOST_FTDC_CL_One  TThostFtdcCheckLevelType = '1' // 一级复核
	THOST_FTDC_CL_Two  TThostFtdcCheckLevelType = '2' // 二级复核
)

// TFtdcCheckNoType是一个操作次数类型
type TThostFtdcCheckNoType int32

// TFtdcCheckStatusType是一个复核级别类型
type TThostFtdcCheckStatusType uint8

//go:generate stringer -type TThostFtdcCheckStatusType -linecomment
const (
	THOST_FTDC_CHS_Init     TThostFtdcCheckStatusType = '0' // 未复核
	THOST_FTDC_CHS_Checking TThostFtdcCheckStatusType = '1' // 复核中
	THOST_FTDC_CHS_Checked  TThostFtdcCheckStatusType = '2' // 已复核
	THOST_FTDC_CHS_Refuse   TThostFtdcCheckStatusType = '3' // 拒绝
	THOST_FTDC_CHS_Cancel   TThostFtdcCheckStatusType = '4' // 作废
)

// TFtdcUsedStatusType是一个生效状态类型
type TThostFtdcUsedStatusType uint8

//go:generate stringer -type TThostFtdcUsedStatusType -linecomment
const (
	THOST_FTDC_CHU_Unused TThostFtdcUsedStatusType = '0' // 未生效
	THOST_FTDC_CHU_Used   TThostFtdcUsedStatusType = '1' // 已生效
	THOST_FTDC_CHU_Fail   TThostFtdcUsedStatusType = '2' // 生效失败
)

// TFtdcRateTemplateNameType是一个模型名称类型
type TThostFtdcRateTemplateNameType [61]byte

func (t TThostFtdcRateTemplateNameType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcRateTemplateNameType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcPropertyStringType是一个用于查询的投资属性字段类型
type TThostFtdcPropertyStringType [2049]byte

func (t TThostFtdcPropertyStringType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcPropertyStringType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcBankAcountOriginType是一个账户来源类型
type TThostFtdcBankAcountOriginType uint8

//go:generate stringer -type TThostFtdcBankAcountOriginType -linecomment
const (
	THOST_FTDC_BAO_ByAccProperty TThostFtdcBankAcountOriginType = '0' // 手工录入
	THOST_FTDC_BAO_ByFBTransfer  TThostFtdcBankAcountOriginType = '1' // 银期转账
)

// TFtdcMonthBillTradeSumType是一个结算单月报成交汇总方式类型
type TThostFtdcMonthBillTradeSumType uint8

//go:generate stringer -type TThostFtdcMonthBillTradeSumType -linecomment
const (
	THOST_FTDC_MBTS_ByInstrument TThostFtdcMonthBillTradeSumType = '0' // 同日同合约
	THOST_FTDC_MBTS_ByDayInsPrc  TThostFtdcMonthBillTradeSumType = '1' // 同日同合约同价格
	THOST_FTDC_MBTS_ByDayIns     TThostFtdcMonthBillTradeSumType = '2' // 同合约
)

// TFtdcFBTTradeCodeEnumType是一个银期交易代码枚举类型
type TThostFtdcFBTTradeCodeEnumType uint8

//go:generate stringer -type TThostFtdcFBTTradeCodeEnumType -linecomment
const (
	THOST_FTDC_FTC_BankLaunchBankToBroker   TThostFtdcFBTTradeCodeEnumType = '1' // 银行发起银行转期货
	THOST_FTDC_FTC_BrokerLaunchBankToBroker TThostFtdcFBTTradeCodeEnumType = '2' // 期货发起银行转期货
	THOST_FTDC_FTC_BankLaunchBrokerToBank   TThostFtdcFBTTradeCodeEnumType = '3' // 银行发起期货转银行
	THOST_FTDC_FTC_BrokerLaunchBrokerToBank TThostFtdcFBTTradeCodeEnumType = '4' // 期货发起期货转银行
)

// TFtdcRateTemplateIDType是一个模型代码类型
type TThostFtdcRateTemplateIDType [9]byte

func (t TThostFtdcRateTemplateIDType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcRateTemplateIDType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcRiskRateType是一个风险度类型
type TThostFtdcRiskRateType [21]byte

func (t TThostFtdcRiskRateType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcRiskRateType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcTimestampType是一个时间戳类型
type TThostFtdcTimestampType int32

// TFtdcInvestorIDRuleNameType是一个号段规则名称类型
type TThostFtdcInvestorIDRuleNameType [61]byte

func (t TThostFtdcInvestorIDRuleNameType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcInvestorIDRuleNameType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcInvestorIDRuleExprType是一个号段规则表达式类型
type TThostFtdcInvestorIDRuleExprType [513]byte

func (t TThostFtdcInvestorIDRuleExprType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcInvestorIDRuleExprType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcLastDriftType是一个上次OTP漂移值类型
type TThostFtdcLastDriftType int32

// TFtdcLastSuccessType是一个上次OTP成功值类型
type TThostFtdcLastSuccessType int32

// TFtdcAuthKeyType是一个令牌密钥类型
type TThostFtdcAuthKeyType [41]byte

func (t TThostFtdcAuthKeyType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcAuthKeyType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcSerialNumberType是一个序列号类型
type TThostFtdcSerialNumberType [17]byte

func (t TThostFtdcSerialNumberType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcSerialNumberType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcOTPTypeType是一个动态令牌类型类型
type TThostFtdcOTPTypeType uint8

//go:generate stringer -type TThostFtdcOTPTypeType -linecomment
const (
	THOST_FTDC_OTP_NONE TThostFtdcOTPTypeType = '0' // 无动态令牌
	THOST_FTDC_OTP_TOTP TThostFtdcOTPTypeType = '1' // 时间令牌
)

// TFtdcOTPVendorsIDType是一个动态令牌提供商类型
type TThostFtdcOTPVendorsIDType [2]byte

func (t TThostFtdcOTPVendorsIDType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcOTPVendorsIDType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcOTPVendorsNameType是一个动态令牌提供商名称类型
type TThostFtdcOTPVendorsNameType [61]byte

func (t TThostFtdcOTPVendorsNameType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcOTPVendorsNameType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcOTPStatusType是一个动态令牌状态类型
type TThostFtdcOTPStatusType uint8

//go:generate stringer -type TThostFtdcOTPStatusType -linecomment
const (
	THOST_FTDC_OTPS_Unused TThostFtdcOTPStatusType = '0' // 未使用
	THOST_FTDC_OTPS_Used   TThostFtdcOTPStatusType = '1' // 已使用
	THOST_FTDC_OTPS_Disuse TThostFtdcOTPStatusType = '2' // 注销
)

// TFtdcBrokerUserTypeType是一个经济公司用户类型类型
type TThostFtdcBrokerUserTypeType uint8

//go:generate stringer -type TThostFtdcBrokerUserTypeType -linecomment
const (
	THOST_FTDC_BUT_Investor   TThostFtdcBrokerUserTypeType = '1' // 投资者
	THOST_FTDC_BUT_BrokerUser TThostFtdcBrokerUserTypeType = '2' // 操作员
)

// TFtdcFutureTypeType是一个期货类型类型
type TThostFtdcFutureTypeType uint8

//go:generate stringer -type TThostFtdcFutureTypeType -linecomment
const (
	THOST_FTDC_FUTT_Commodity TThostFtdcFutureTypeType = '1' // 商品期货
	THOST_FTDC_FUTT_Financial TThostFtdcFutureTypeType = '2' // 金融期货
)

// TFtdcFundEventTypeType是一个资金管理操作类型类型
type TThostFtdcFundEventTypeType uint8

//go:generate stringer -type TThostFtdcFundEventTypeType -linecomment
const (
	THOST_FTDC_FET_Restriction         TThostFtdcFundEventTypeType = '0' // 转账限额
	THOST_FTDC_FET_TodayRestriction    TThostFtdcFundEventTypeType = '1' // 当日转账限额
	THOST_FTDC_FET_Transfer            TThostFtdcFundEventTypeType = '2' // 期商流水
	THOST_FTDC_FET_Credit              TThostFtdcFundEventTypeType = '3' // 资金冻结
	THOST_FTDC_FET_InvestorWithdrawAlm TThostFtdcFundEventTypeType = '4' // 投资者可提资金比例
	THOST_FTDC_FET_BankRestriction     TThostFtdcFundEventTypeType = '5' // 单个银行帐户转账限额
	THOST_FTDC_FET_Accountregister     TThostFtdcFundEventTypeType = '6' // 银期签约账户
	THOST_FTDC_FET_ExchangeFundIO      TThostFtdcFundEventTypeType = '7' // 交易所出入金
	THOST_FTDC_FET_InvestorFundIO      TThostFtdcFundEventTypeType = '8' // 投资者出入金
)

// TFtdcAccountSourceTypeType是一个资金账户来源类型
type TThostFtdcAccountSourceTypeType uint8

//go:generate stringer -type TThostFtdcAccountSourceTypeType -linecomment
const (
	THOST_FTDC_AST_FBTransfer  TThostFtdcAccountSourceTypeType = '0' // 银期同步
	THOST_FTDC_AST_ManualEntry TThostFtdcAccountSourceTypeType = '1' // 手工录入
)

// TFtdcCodeSourceTypeType是一个交易编码来源类型
type TThostFtdcCodeSourceTypeType uint8

//go:generate stringer -type TThostFtdcCodeSourceTypeType -linecomment
const (
	THOST_FTDC_CST_UnifyAccount TThostFtdcCodeSourceTypeType = '0' // 统一开户(已规范)
	THOST_FTDC_CST_ManualEntry  TThostFtdcCodeSourceTypeType = '1' // 手工录入(未规范)
)

// TFtdcUserRangeType是一个操作员范围类型
type TThostFtdcUserRangeType uint8

//go:generate stringer -type TThostFtdcUserRangeType -linecomment
const (
	THOST_FTDC_UR_All    TThostFtdcUserRangeType = '0' // 所有
	THOST_FTDC_UR_Single TThostFtdcUserRangeType = '1' // 单一操作员
)

// TFtdcTimeSpanType是一个时间跨度类型
type TThostFtdcTimeSpanType [9]byte

func (t TThostFtdcTimeSpanType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcTimeSpanType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcImportSequenceIDType是一个动态令牌导入批次编号类型
type TThostFtdcImportSequenceIDType [17]byte

func (t TThostFtdcImportSequenceIDType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcImportSequenceIDType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcByGroupType是一个交易统计表按客户统计方式类型
type TThostFtdcByGroupType uint8

//go:generate stringer -type TThostFtdcByGroupType -linecomment
const (
	THOST_FTDC_BG_Investor TThostFtdcByGroupType = '2' // 按投资者统计
	THOST_FTDC_BG_Group    TThostFtdcByGroupType = '1' // 按类统计
)

// TFtdcTradeSumStatModeType是一个交易统计表按范围统计方式类型
type TThostFtdcTradeSumStatModeType uint8

//go:generate stringer -type TThostFtdcTradeSumStatModeType -linecomment
const (
	THOST_FTDC_TSSM_Instrument TThostFtdcTradeSumStatModeType = '1' // 按合约统计
	THOST_FTDC_TSSM_Product    TThostFtdcTradeSumStatModeType = '2' // 按产品统计
	THOST_FTDC_TSSM_Exchange   TThostFtdcTradeSumStatModeType = '3' // 按交易所统计
)

// TFtdcComTypeType是一个组合成交类型类型
type TThostFtdcComTypeType int32

// TFtdcUserProductIDType是一个产品标识类型
type TThostFtdcUserProductIDType [33]byte

func (t TThostFtdcUserProductIDType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcUserProductIDType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcUserProductNameType是一个产品名称类型
type TThostFtdcUserProductNameType [65]byte

func (t TThostFtdcUserProductNameType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcUserProductNameType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcUserProductMemoType是一个产品说明类型
type TThostFtdcUserProductMemoType [129]byte

func (t TThostFtdcUserProductMemoType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcUserProductMemoType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcCSRCCancelFlagType是一个新增或变更标志类型
type TThostFtdcCSRCCancelFlagType [2]byte

func (t TThostFtdcCSRCCancelFlagType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcCSRCCancelFlagType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcCSRCDateType是一个日期类型
type TThostFtdcCSRCDateType [11]byte

func (t TThostFtdcCSRCDateType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcCSRCDateType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcCSRCInvestorNameType是一个客户名称类型
type TThostFtdcCSRCInvestorNameType [201]byte

func (t TThostFtdcCSRCInvestorNameType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcCSRCInvestorNameType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcCSRCOpenInvestorNameType是一个客户名称类型
type TThostFtdcCSRCOpenInvestorNameType [101]byte

func (t TThostFtdcCSRCOpenInvestorNameType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcCSRCOpenInvestorNameType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcCSRCInvestorIDType是一个客户代码类型
type TThostFtdcCSRCInvestorIDType [13]byte

func (t TThostFtdcCSRCInvestorIDType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcCSRCInvestorIDType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcCSRCIdentifiedCardNoType是一个证件号码类型
type TThostFtdcCSRCIdentifiedCardNoType [51]byte

func (t TThostFtdcCSRCIdentifiedCardNoType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcCSRCIdentifiedCardNoType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcCSRCClientIDType是一个交易编码类型
type TThostFtdcCSRCClientIDType [11]byte

func (t TThostFtdcCSRCClientIDType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcCSRCClientIDType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcCSRCBankFlagType是一个银行标识类型
type TThostFtdcCSRCBankFlagType [3]byte

func (t TThostFtdcCSRCBankFlagType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcCSRCBankFlagType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcCSRCBankAccountType是一个银行账户类型
type TThostFtdcCSRCBankAccountType [23]byte

func (t TThostFtdcCSRCBankAccountType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcCSRCBankAccountType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcCSRCOpenNameType是一个开户人类型
type TThostFtdcCSRCOpenNameType [401]byte

func (t TThostFtdcCSRCOpenNameType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcCSRCOpenNameType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcCSRCMemoType是一个说明类型
type TThostFtdcCSRCMemoType [101]byte

func (t TThostFtdcCSRCMemoType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcCSRCMemoType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcCSRCTimeType是一个时间类型
type TThostFtdcCSRCTimeType [11]byte

func (t TThostFtdcCSRCTimeType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcCSRCTimeType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcCSRCTradeIDType是一个成交流水号类型
type TThostFtdcCSRCTradeIDType [21]byte

func (t TThostFtdcCSRCTradeIDType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcCSRCTradeIDType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcCSRCExchangeInstIDType是一个合约代码类型
type TThostFtdcCSRCExchangeInstIDType [81]byte

func (t TThostFtdcCSRCExchangeInstIDType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcCSRCExchangeInstIDType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcCSRCMortgageNameType是一个质押品名称类型
type TThostFtdcCSRCMortgageNameType [7]byte

func (t TThostFtdcCSRCMortgageNameType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcCSRCMortgageNameType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcCSRCReasonType是一个事由类型
type TThostFtdcCSRCReasonType [3]byte

func (t TThostFtdcCSRCReasonType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcCSRCReasonType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcIsSettlementType是一个是否为非结算会员类型
type TThostFtdcIsSettlementType [2]byte

func (t TThostFtdcIsSettlementType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcIsSettlementType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcCSRCMoneyType是一个资金类型
type TThostFtdcCSRCMoneyType float64

func (t TThostFtdcCSRCMoneyType) IsValid() bool {
	return float64(t) != math.MaxFloat64
}

func (t TThostFtdcCSRCMoneyType) String() string {
	if t.IsValid() {
		return strconv.FormatFloat(float64(t), 'f', 6, 64)
	} else {
		return "DBL_MAX"
	}
}

// TFtdcCSRCPriceType是一个价格类型
type TThostFtdcCSRCPriceType float64

func (t TThostFtdcCSRCPriceType) IsValid() bool {
	return float64(t) != math.MaxFloat64
}

func (t TThostFtdcCSRCPriceType) String() string {
	if t.IsValid() {
		return strconv.FormatFloat(float64(t), 'f', 6, 64)
	} else {
		return "DBL_MAX"
	}
}

// TFtdcCSRCOptionsTypeType是一个期权类型类型
type TThostFtdcCSRCOptionsTypeType [2]byte

func (t TThostFtdcCSRCOptionsTypeType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcCSRCOptionsTypeType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcCSRCStrikePriceType是一个执行价类型
type TThostFtdcCSRCStrikePriceType float64

func (t TThostFtdcCSRCStrikePriceType) IsValid() bool {
	return float64(t) != math.MaxFloat64
}

func (t TThostFtdcCSRCStrikePriceType) String() string {
	if t.IsValid() {
		return strconv.FormatFloat(float64(t), 'f', 6, 64)
	} else {
		return "DBL_MAX"
	}
}

// TFtdcCSRCTargetProductIDType是一个标的品种类型
type TThostFtdcCSRCTargetProductIDType [3]byte

func (t TThostFtdcCSRCTargetProductIDType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcCSRCTargetProductIDType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcCSRCTargetInstrIDType是一个标的合约类型
type TThostFtdcCSRCTargetInstrIDType [31]byte

func (t TThostFtdcCSRCTargetInstrIDType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcCSRCTargetInstrIDType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcCommModelNameType是一个手续费率模板名称类型
type TThostFtdcCommModelNameType [161]byte

func (t TThostFtdcCommModelNameType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcCommModelNameType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcCommModelMemoType是一个手续费率模板备注类型
type TThostFtdcCommModelMemoType [1025]byte

func (t TThostFtdcCommModelMemoType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcCommModelMemoType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcExprSetModeType是一个日期表达式设置类型类型
type TThostFtdcExprSetModeType uint8

//go:generate stringer -type TThostFtdcExprSetModeType -linecomment
const (
	THOST_FTDC_ESM_Relative TThostFtdcExprSetModeType = '1' // 相对已有规则设置
	THOST_FTDC_ESM_Typical  TThostFtdcExprSetModeType = '2' // 典型设置
)

// TFtdcRateInvestorRangeType是一个投资者范围类型
type TThostFtdcRateInvestorRangeType uint8

//go:generate stringer -type TThostFtdcRateInvestorRangeType -linecomment
const (
	THOST_FTDC_RIR_All    TThostFtdcRateInvestorRangeType = '1' // 公司标准
	THOST_FTDC_RIR_Model  TThostFtdcRateInvestorRangeType = '2' // 模板
	THOST_FTDC_RIR_Single TThostFtdcRateInvestorRangeType = '3' // 单一投资者
)

// TFtdcAgentBrokerIDType是一个代理经纪公司代码类型
type TThostFtdcAgentBrokerIDType [13]byte

func (t TThostFtdcAgentBrokerIDType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcAgentBrokerIDType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcDRIdentityIDType是一个交易中心代码类型
type TThostFtdcDRIdentityIDType int32

// TFtdcDRIdentityNameType是一个交易中心名称类型
type TThostFtdcDRIdentityNameType [65]byte

func (t TThostFtdcDRIdentityNameType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcDRIdentityNameType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcDBLinkIDType是一个DBLink标识号类型
type TThostFtdcDBLinkIDType [31]byte

func (t TThostFtdcDBLinkIDType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcDBLinkIDType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcSyncDataStatusType是一个主次用系统数据同步状态类型
type TThostFtdcSyncDataStatusType uint8

//go:generate stringer -type TThostFtdcSyncDataStatusType -linecomment
const (
	THOST_FTDC_SDS_Initialize    TThostFtdcSyncDataStatusType = '0' // 未同步
	THOST_FTDC_SDS_Settlementing TThostFtdcSyncDataStatusType = '1' // 同步中
	THOST_FTDC_SDS_Settlemented  TThostFtdcSyncDataStatusType = '2' // 已同步
)

// TFtdcTradeSourceType是一个成交来源类型
type TThostFtdcTradeSourceType uint8

//go:generate stringer -type TThostFtdcTradeSourceType -linecomment
const (
	THOST_FTDC_TSRC_NORMAL TThostFtdcTradeSourceType = '0' // 来自交易所普通回报
	THOST_FTDC_TSRC_QUERY  TThostFtdcTradeSourceType = '1' // 来自查询
)

// TFtdcFlexStatModeType是一个产品合约统计方式类型
type TThostFtdcFlexStatModeType uint8

//go:generate stringer -type TThostFtdcFlexStatModeType -linecomment
const (
	THOST_FTDC_FSM_Product  TThostFtdcFlexStatModeType = '1' // 产品统计
	THOST_FTDC_FSM_Exchange TThostFtdcFlexStatModeType = '2' // 交易所统计
	THOST_FTDC_FSM_All      TThostFtdcFlexStatModeType = '3' // 统计所有
)

// TFtdcByInvestorRangeType是一个投资者范围统计方式类型
type TThostFtdcByInvestorRangeType uint8

//go:generate stringer -type TThostFtdcByInvestorRangeType -linecomment
const (
	THOST_FTDC_BIR_Property TThostFtdcByInvestorRangeType = '1' // 属性统计
	THOST_FTDC_BIR_All      TThostFtdcByInvestorRangeType = '2' // 统计所有
)

// TFtdcSRiskRateType是一个风险度类型
type TThostFtdcSRiskRateType [21]byte

func (t TThostFtdcSRiskRateType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcSRiskRateType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcSequenceNo12Type是一个序号类型
type TThostFtdcSequenceNo12Type int32

// TFtdcPropertyInvestorRangeType是一个投资者范围类型
type TThostFtdcPropertyInvestorRangeType uint8

//go:generate stringer -type TThostFtdcPropertyInvestorRangeType -linecomment
const (
	THOST_FTDC_PIR_All      TThostFtdcPropertyInvestorRangeType = '1' // 所有
	THOST_FTDC_PIR_Property TThostFtdcPropertyInvestorRangeType = '2' // 投资者属性
	THOST_FTDC_PIR_Single   TThostFtdcPropertyInvestorRangeType = '3' // 单一投资者
)

// TFtdcFileStatusType是一个文件状态类型
type TThostFtdcFileStatusType uint8

//go:generate stringer -type TThostFtdcFileStatusType -linecomment
const (
	THOST_FTDC_FIS_NoCreate TThostFtdcFileStatusType = '0' // 未生成
	THOST_FTDC_FIS_Created  TThostFtdcFileStatusType = '1' // 已生成
	THOST_FTDC_FIS_Failed   TThostFtdcFileStatusType = '2' // 生成失败
)

// TFtdcFileGenStyleType是一个文件生成方式类型
type TThostFtdcFileGenStyleType uint8

//go:generate stringer -type TThostFtdcFileGenStyleType -linecomment
const (
	THOST_FTDC_FGS_FileTransmit TThostFtdcFileGenStyleType = '0' // 下发
	THOST_FTDC_FGS_FileGen      TThostFtdcFileGenStyleType = '1' // 生成
)

// TFtdcSysOperModeType是一个系统日志操作方法类型
type TThostFtdcSysOperModeType uint8

//go:generate stringer -type TThostFtdcSysOperModeType -linecomment
const (
	THOST_FTDC_SoM_Add    TThostFtdcSysOperModeType = '1' // 增加
	THOST_FTDC_SoM_Update TThostFtdcSysOperModeType = '2' // 修改
	THOST_FTDC_SoM_Delete TThostFtdcSysOperModeType = '3' // 删除
	THOST_FTDC_SoM_Copy   TThostFtdcSysOperModeType = '4' // 复制
	THOST_FTDC_SoM_AcTive TThostFtdcSysOperModeType = '5' // 激活
	THOST_FTDC_SoM_CanCel TThostFtdcSysOperModeType = '6' // 注销
	THOST_FTDC_SoM_ReSet  TThostFtdcSysOperModeType = '7' // 重置
)

// TFtdcSysOperTypeType是一个系统日志操作类型类型
type TThostFtdcSysOperTypeType uint8

//go:generate stringer -type TThostFtdcSysOperTypeType -linecomment
const (
	THOST_FTDC_SoT_UpdatePassword          TThostFtdcSysOperTypeType = '0' // 修改操作员密码
	THOST_FTDC_SoT_UserDepartment          TThostFtdcSysOperTypeType = '1' // 操作员组织架构关系
	THOST_FTDC_SoT_RoleManager             TThostFtdcSysOperTypeType = '2' // 角色管理
	THOST_FTDC_SoT_RoleFunction            TThostFtdcSysOperTypeType = '3' // 角色功能设置
	THOST_FTDC_SoT_BaseParam               TThostFtdcSysOperTypeType = '4' // 基础参数设置
	THOST_FTDC_SoT_SetUserID               TThostFtdcSysOperTypeType = '5' // 设置操作员
	THOST_FTDC_SoT_SetUserRole             TThostFtdcSysOperTypeType = '6' // 用户角色设置
	THOST_FTDC_SoT_UserIpRestriction       TThostFtdcSysOperTypeType = '7' // 用户IP限制
	THOST_FTDC_SoT_DepartmentManager       TThostFtdcSysOperTypeType = '8' // 组织架构管理
	THOST_FTDC_SoT_DepartmentCopy          TThostFtdcSysOperTypeType = '9' // 组织架构向查询分类复制
	THOST_FTDC_SoT_Tradingcode             TThostFtdcSysOperTypeType = 'A' // 交易编码管理
	THOST_FTDC_SoT_InvestorStatus          TThostFtdcSysOperTypeType = 'B' // 投资者状态维护
	THOST_FTDC_SoT_InvestorAuthority       TThostFtdcSysOperTypeType = 'C' // 投资者权限管理
	THOST_FTDC_SoT_PropertySet             TThostFtdcSysOperTypeType = 'D' // 属性设置
	THOST_FTDC_SoT_ReSetInvestorPasswd     TThostFtdcSysOperTypeType = 'E' // 重置投资者密码
	THOST_FTDC_SoT_InvestorPersonalityInfo TThostFtdcSysOperTypeType = 'F' // 投资者个性信息维护
)

// TFtdcCSRCDataQueyTypeType是一个上报数据查询类型类型
type TThostFtdcCSRCDataQueyTypeType uint8

//go:generate stringer -type TThostFtdcCSRCDataQueyTypeType -linecomment
const (
	THOST_FTDC_CSRCQ_Current TThostFtdcCSRCDataQueyTypeType = '0' // 查询当前交易日报送的数据
	THOST_FTDC_CSRCQ_History TThostFtdcCSRCDataQueyTypeType = '1' // 查询历史报送的代理经纪公司的数据
)

// TFtdcFreezeStatusType是一个休眠状态类型
type TThostFtdcFreezeStatusType uint8

//go:generate stringer -type TThostFtdcFreezeStatusType -linecomment
const (
	THOST_FTDC_FRS_Normal TThostFtdcFreezeStatusType = '1' // 活跃
	THOST_FTDC_FRS_Freeze TThostFtdcFreezeStatusType = '0' // 休眠
)

// TFtdcStandardStatusType是一个规范状态类型
type TThostFtdcStandardStatusType uint8

//go:generate stringer -type TThostFtdcStandardStatusType -linecomment
const (
	THOST_FTDC_STST_Standard    TThostFtdcStandardStatusType = '0' // 已规范
	THOST_FTDC_STST_NonStandard TThostFtdcStandardStatusType = '1' // 未规范
)

// TFtdcCSRCFreezeStatusType是一个休眠状态类型
type TThostFtdcCSRCFreezeStatusType [2]byte

func (t TThostFtdcCSRCFreezeStatusType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcCSRCFreezeStatusType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcRightParamTypeType是一个配置类型类型
type TThostFtdcRightParamTypeType uint8

//go:generate stringer -type TThostFtdcRightParamTypeType -linecomment
const (
	THOST_FTDC_RPT_Freeze           TThostFtdcRightParamTypeType = '1' // 休眠户
	THOST_FTDC_RPT_FreezeActive     TThostFtdcRightParamTypeType = '2' // 激活休眠户
	THOST_FTDC_RPT_OpenLimit        TThostFtdcRightParamTypeType = '3' // 开仓权限限制
	THOST_FTDC_RPT_RelieveOpenLimit TThostFtdcRightParamTypeType = '4' // 解除开仓权限限制
)

// TFtdcRightTemplateIDType是一个模板代码类型
type TThostFtdcRightTemplateIDType [9]byte

func (t TThostFtdcRightTemplateIDType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcRightTemplateIDType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcRightTemplateNameType是一个模板名称类型
type TThostFtdcRightTemplateNameType [61]byte

func (t TThostFtdcRightTemplateNameType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcRightTemplateNameType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcDataStatusType是一个反洗钱审核表数据状态类型
type TThostFtdcDataStatusType uint8

//go:generate stringer -type TThostFtdcDataStatusType -linecomment
const (
	THOST_FTDC_AMLDS_Normal  TThostFtdcDataStatusType = '0' // 正常
	THOST_FTDC_AMLDS_Deleted TThostFtdcDataStatusType = '1' // 已删除
)

// TFtdcAMLCheckStatusType是一个审核状态类型
type TThostFtdcAMLCheckStatusType uint8

//go:generate stringer -type TThostFtdcAMLCheckStatusType -linecomment
const (
	THOST_FTDC_AMLCHS_Init         TThostFtdcAMLCheckStatusType = '0' // 未复核
	THOST_FTDC_AMLCHS_Checking     TThostFtdcAMLCheckStatusType = '1' // 复核中
	THOST_FTDC_AMLCHS_Checked      TThostFtdcAMLCheckStatusType = '2' // 已复核
	THOST_FTDC_AMLCHS_RefuseReport TThostFtdcAMLCheckStatusType = '3' // 拒绝上报
)

// TFtdcAmlDateTypeType是一个日期类型类型
type TThostFtdcAmlDateTypeType uint8

//go:generate stringer -type TThostFtdcAmlDateTypeType -linecomment
const (
	THOST_FTDC_AMLDT_DrawDay  TThostFtdcAmlDateTypeType = '0' // 检查日期
	THOST_FTDC_AMLDT_TouchDay TThostFtdcAmlDateTypeType = '1' // 发生日期
)

// TFtdcAmlCheckLevelType是一个审核级别类型
type TThostFtdcAmlCheckLevelType uint8

//go:generate stringer -type TThostFtdcAmlCheckLevelType -linecomment
const (
	THOST_FTDC_AMLCL_CheckLevel0 TThostFtdcAmlCheckLevelType = '0' // 零级审核
	THOST_FTDC_AMLCL_CheckLevel1 TThostFtdcAmlCheckLevelType = '1' // 一级审核
	THOST_FTDC_AMLCL_CheckLevel2 TThostFtdcAmlCheckLevelType = '2' // 二级审核
	THOST_FTDC_AMLCL_CheckLevel3 TThostFtdcAmlCheckLevelType = '3' // 三级审核
)

// TFtdcAmlCheckFlowType是一个反洗钱数据抽取审核流程类型
type TThostFtdcAmlCheckFlowType [2]byte

func (t TThostFtdcAmlCheckFlowType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcAmlCheckFlowType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcDataTypeType是一个数据类型类型
type TThostFtdcDataTypeType [129]byte

func (t TThostFtdcDataTypeType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcDataTypeType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcExportFileTypeType是一个导出文件类型类型
type TThostFtdcExportFileTypeType uint8

//go:generate stringer -type TThostFtdcExportFileTypeType -linecomment
const (
	THOST_FTDC_EFT_CSV   TThostFtdcExportFileTypeType = '0' // CSV
	THOST_FTDC_EFT_EXCEL TThostFtdcExportFileTypeType = '1' // Excel
	THOST_FTDC_EFT_DBF   TThostFtdcExportFileTypeType = '2' // DBF
)

// TFtdcSettleManagerTypeType是一个结算配置类型类型
type TThostFtdcSettleManagerTypeType uint8

//go:generate stringer -type TThostFtdcSettleManagerTypeType -linecomment
const (
	THOST_FTDC_SMT_Before       TThostFtdcSettleManagerTypeType = '1' // 结算前准备
	THOST_FTDC_SMT_Settlement   TThostFtdcSettleManagerTypeType = '2' // 结算
	THOST_FTDC_SMT_After        TThostFtdcSettleManagerTypeType = '3' // 结算后核对
	THOST_FTDC_SMT_Settlemented TThostFtdcSettleManagerTypeType = '4' // 结算后处理
)

// TFtdcSettleManagerIDType是一个结算配置代码类型
type TThostFtdcSettleManagerIDType [33]byte

func (t TThostFtdcSettleManagerIDType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcSettleManagerIDType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcSettleManagerNameType是一个结算配置名称类型
type TThostFtdcSettleManagerNameType [129]byte

func (t TThostFtdcSettleManagerNameType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcSettleManagerNameType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcSettleManagerLevelType是一个结算配置等级类型
type TThostFtdcSettleManagerLevelType uint8

//go:generate stringer -type TThostFtdcSettleManagerLevelType -linecomment
const (
	THOST_FTDC_SML_Must   TThostFtdcSettleManagerLevelType = '1' // 必要
	THOST_FTDC_SML_Alarm  TThostFtdcSettleManagerLevelType = '2' // 警告
	THOST_FTDC_SML_Prompt TThostFtdcSettleManagerLevelType = '3' // 提示
	THOST_FTDC_SML_Ignore TThostFtdcSettleManagerLevelType = '4' // 不检查
)

// TFtdcSettleManagerGroupType是一个模块分组类型
type TThostFtdcSettleManagerGroupType uint8

//go:generate stringer -type TThostFtdcSettleManagerGroupType -linecomment
const (
	THOST_FTDC_SMG_Exhcange TThostFtdcSettleManagerGroupType = '1' // 交易所核对
	THOST_FTDC_SMG_ASP      TThostFtdcSettleManagerGroupType = '2' // 内部核对
	THOST_FTDC_SMG_CSRC     TThostFtdcSettleManagerGroupType = '3' // 上报数据核对
)

// TFtdcCheckResultMemoType是一个核对结果说明类型
type TThostFtdcCheckResultMemoType [1025]byte

func (t TThostFtdcCheckResultMemoType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcCheckResultMemoType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcFunctionUrlType是一个功能链接类型
type TThostFtdcFunctionUrlType [1025]byte

func (t TThostFtdcFunctionUrlType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcFunctionUrlType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcAuthInfoType是一个客户端认证信息类型
type TThostFtdcAuthInfoType [129]byte

func (t TThostFtdcAuthInfoType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcAuthInfoType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcAuthCodeType是一个客户端认证码类型
type TThostFtdcAuthCodeType [17]byte

func (t TThostFtdcAuthCodeType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcAuthCodeType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcLimitUseTypeType是一个保值额度使用类型类型
type TThostFtdcLimitUseTypeType uint8

//go:generate stringer -type TThostFtdcLimitUseTypeType -linecomment
const (
	THOST_FTDC_LUT_Repeatable   TThostFtdcLimitUseTypeType = '1' // 可重复使用
	THOST_FTDC_LUT_Unrepeatable TThostFtdcLimitUseTypeType = '2' // 不可重复使用
)

// TFtdcDataResourceType是一个数据来源类型
type TThostFtdcDataResourceType uint8

//go:generate stringer -type TThostFtdcDataResourceType -linecomment
const (
	THOST_FTDC_DAR_Settle   TThostFtdcDataResourceType = '1' // 本系统
	THOST_FTDC_DAR_Exchange TThostFtdcDataResourceType = '2' // 交易所
	THOST_FTDC_DAR_CSRC     TThostFtdcDataResourceType = '3' // 报送数据
)

// TFtdcMarginTypeType是一个保证金类型类型
type TThostFtdcMarginTypeType uint8

//go:generate stringer -type TThostFtdcMarginTypeType -linecomment
const (
	THOST_FTDC_MGT_ExchMarginRate       TThostFtdcMarginTypeType = '0' // 交易所保证金率
	THOST_FTDC_MGT_InstrMarginRate      TThostFtdcMarginTypeType = '1' // 投资者保证金率
	THOST_FTDC_MGT_InstrMarginRateTrade TThostFtdcMarginTypeType = '2' // 投资者交易保证金率
)

// TFtdcActiveTypeType是一个生效类型类型
type TThostFtdcActiveTypeType uint8

//go:generate stringer -type TThostFtdcActiveTypeType -linecomment
const (
	THOST_FTDC_ACT_Intraday TThostFtdcActiveTypeType = '1' // 仅当日生效
	THOST_FTDC_ACT_Long     TThostFtdcActiveTypeType = '2' // 长期生效
)

// TFtdcMarginRateTypeType是一个冲突保证金率类型类型
type TThostFtdcMarginRateTypeType uint8

//go:generate stringer -type TThostFtdcMarginRateTypeType -linecomment
const (
	THOST_FTDC_MRT_Exchange      TThostFtdcMarginRateTypeType = '1' // 交易所保证金率
	THOST_FTDC_MRT_Investor      TThostFtdcMarginRateTypeType = '2' // 投资者保证金率
	THOST_FTDC_MRT_InvestorTrade TThostFtdcMarginRateTypeType = '3' // 投资者交易保证金率
)

// TFtdcBackUpStatusType是一个备份数据状态类型
type TThostFtdcBackUpStatusType uint8

//go:generate stringer -type TThostFtdcBackUpStatusType -linecomment
const (
	THOST_FTDC_BUS_UnBak   TThostFtdcBackUpStatusType = '0' // 未生成备份数据
	THOST_FTDC_BUS_BakUp   TThostFtdcBackUpStatusType = '1' // 备份数据生成中
	THOST_FTDC_BUS_BakUped TThostFtdcBackUpStatusType = '2' // 已生成备份数据
	THOST_FTDC_BUS_BakFail TThostFtdcBackUpStatusType = '3' // 备份数据失败
)

// TFtdcInitSettlementType是一个结算初始化状态类型
type TThostFtdcInitSettlementType uint8

//go:generate stringer -type TThostFtdcInitSettlementType -linecomment
const (
	THOST_FTDC_SIS_UnInitialize TThostFtdcInitSettlementType = '0' // 结算初始化未开始
	THOST_FTDC_SIS_Initialize   TThostFtdcInitSettlementType = '1' // 结算初始化中
	THOST_FTDC_SIS_Initialized  TThostFtdcInitSettlementType = '2' // 结算初始化完成
)

// TFtdcReportStatusType是一个报表数据生成状态类型
type TThostFtdcReportStatusType uint8

//go:generate stringer -type TThostFtdcReportStatusType -linecomment
const (
	THOST_FTDC_SRS_NoCreate   TThostFtdcReportStatusType = '0' // 未生成报表数据
	THOST_FTDC_SRS_Create     TThostFtdcReportStatusType = '1' // 报表数据生成中
	THOST_FTDC_SRS_Created    TThostFtdcReportStatusType = '2' // 已生成报表数据
	THOST_FTDC_SRS_CreateFail TThostFtdcReportStatusType = '3' // 生成报表数据失败
)

// TFtdcSaveStatusType是一个数据归档状态类型
type TThostFtdcSaveStatusType uint8

//go:generate stringer -type TThostFtdcSaveStatusType -linecomment
const (
	THOST_FTDC_SSS_UnSaveData TThostFtdcSaveStatusType = '0' // 归档未完成
	THOST_FTDC_SSS_SaveDatad  TThostFtdcSaveStatusType = '1' // 归档完成
)

// TFtdcSettArchiveStatusType是一个结算确认数据归档状态类型
type TThostFtdcSettArchiveStatusType uint8

//go:generate stringer -type TThostFtdcSettArchiveStatusType -linecomment
const (
	THOST_FTDC_SAS_UnArchived  TThostFtdcSettArchiveStatusType = '0' // 未归档数据
	THOST_FTDC_SAS_Archiving   TThostFtdcSettArchiveStatusType = '1' // 数据归档中
	THOST_FTDC_SAS_Archived    TThostFtdcSettArchiveStatusType = '2' // 已归档数据
	THOST_FTDC_SAS_ArchiveFail TThostFtdcSettArchiveStatusType = '3' // 归档数据失败
)

// TFtdcCTPTypeType是一个CTP交易系统类型类型
type TThostFtdcCTPTypeType uint8

//go:generate stringer -type TThostFtdcCTPTypeType -linecomment
const (
	THOST_FTDC_CTPT_Unkown     TThostFtdcCTPTypeType = '0' // 未知类型
	THOST_FTDC_CTPT_MainCenter TThostFtdcCTPTypeType = '1' // 主中心
	THOST_FTDC_CTPT_BackUp     TThostFtdcCTPTypeType = '2' // 备中心
)

// TFtdcToolIDType是一个工具代码类型
type TThostFtdcToolIDType [9]byte

func (t TThostFtdcToolIDType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcToolIDType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcToolNameType是一个工具名称类型
type TThostFtdcToolNameType [81]byte

func (t TThostFtdcToolNameType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcToolNameType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcCloseDealTypeType是一个平仓处理类型类型
type TThostFtdcCloseDealTypeType uint8

//go:generate stringer -type TThostFtdcCloseDealTypeType -linecomment
const (
	THOST_FTDC_CDT_Normal    TThostFtdcCloseDealTypeType = '0' // 正常
	THOST_FTDC_CDT_SpecFirst TThostFtdcCloseDealTypeType = '1' // 投机平仓优先
)

// TFtdcMortgageFundUseRangeType是一个货币质押资金可用范围类型
type TThostFtdcMortgageFundUseRangeType uint8

//go:generate stringer -type TThostFtdcMortgageFundUseRangeType -linecomment
const (
	THOST_FTDC_MFUR_None   TThostFtdcMortgageFundUseRangeType = '0' // 不能使用
	THOST_FTDC_MFUR_Margin TThostFtdcMortgageFundUseRangeType = '1' // 用于保证金
	THOST_FTDC_MFUR_All    TThostFtdcMortgageFundUseRangeType = '2' // 用于手续费、盈亏、保证金
)

// TFtdcCurrencyUnitType是一个币种单位数量类型
type TThostFtdcCurrencyUnitType float64

func (t TThostFtdcCurrencyUnitType) IsValid() bool {
	return float64(t) != math.MaxFloat64
}

func (t TThostFtdcCurrencyUnitType) String() string {
	if t.IsValid() {
		return strconv.FormatFloat(float64(t), 'f', 6, 64)
	} else {
		return "DBL_MAX"
	}
}

// TFtdcExchangeRateType是一个汇率类型
type TThostFtdcExchangeRateType float64

func (t TThostFtdcExchangeRateType) IsValid() bool {
	return float64(t) != math.MaxFloat64
}

func (t TThostFtdcExchangeRateType) String() string {
	if t.IsValid() {
		return strconv.FormatFloat(float64(t), 'f', 6, 64)
	} else {
		return "DBL_MAX"
	}
}

// TFtdcSpecProductTypeType是一个特殊产品类型类型
type TThostFtdcSpecProductTypeType uint8

//go:generate stringer -type TThostFtdcSpecProductTypeType -linecomment
const (
	THOST_FTDC_SPT_CzceHedge          TThostFtdcSpecProductTypeType = '1' // 郑商所套保产品
	THOST_FTDC_SPT_IneForeignCurrency TThostFtdcSpecProductTypeType = '2' // 货币质押产品
	THOST_FTDC_SPT_DceOpenClose       TThostFtdcSpecProductTypeType = '3' // 大连短线开平仓产品
)

// TFtdcFundMortgageTypeType是一个货币质押类型类型
type TThostFtdcFundMortgageTypeType uint8

//go:generate stringer -type TThostFtdcFundMortgageTypeType -linecomment
const (
	THOST_FTDC_FMT_Mortgage   TThostFtdcFundMortgageTypeType = '1' // 质押
	THOST_FTDC_FMT_Redemption TThostFtdcFundMortgageTypeType = '2' // 解质
)

// TFtdcAccountSettlementParamIDType是一个投资者账户结算参数代码类型
type TThostFtdcAccountSettlementParamIDType uint8

//go:generate stringer -type TThostFtdcAccountSettlementParamIDType -linecomment
const (
	THOST_FTDC_ASPI_BaseMargin     TThostFtdcAccountSettlementParamIDType = '1' // 基础保证金
	THOST_FTDC_ASPI_LowestInterest TThostFtdcAccountSettlementParamIDType = '2' // 最低权益标准
)

// TFtdcCurrencyNameType是一个币种名称类型
type TThostFtdcCurrencyNameType [31]byte

func (t TThostFtdcCurrencyNameType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcCurrencyNameType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcCurrencySignType是一个币种符号类型
type TThostFtdcCurrencySignType [4]byte

func (t TThostFtdcCurrencySignType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcCurrencySignType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcFundMortDirectionType是一个货币质押方向类型
type TThostFtdcFundMortDirectionType uint8

//go:generate stringer -type TThostFtdcFundMortDirectionType -linecomment
const (
	THOST_FTDC_FMD_In  TThostFtdcFundMortDirectionType = '1' // 货币质入
	THOST_FTDC_FMD_Out TThostFtdcFundMortDirectionType = '2' // 货币质出
)

type TThostFtdcBusinessClassType uint8

// TFtdcSwapSourceTypeType是一个换汇数据来源类型
type TThostFtdcSwapSourceTypeType uint8

//go:generate stringer -type TThostFtdcSwapSourceTypeType -linecomment
const (
	THOST_FTDC_SST_Manual    TThostFtdcSwapSourceTypeType = '0' // 手工
	THOST_FTDC_SST_Automatic TThostFtdcSwapSourceTypeType = '1' // 自动生成
)

// TFtdcCurrExDirectionType是一个换汇类型类型
type TThostFtdcCurrExDirectionType uint8

//go:generate stringer -type TThostFtdcCurrExDirectionType -linecomment
const (
	THOST_FTDC_CED_Settlement TThostFtdcCurrExDirectionType = '0' // 结汇
	THOST_FTDC_CED_Sale       TThostFtdcCurrExDirectionType = '1' // 售汇
)

// TFtdcCurrencySwapStatusType是一个申请状态类型
type TThostFtdcCurrencySwapStatusType uint8

//go:generate stringer -type TThostFtdcCurrencySwapStatusType -linecomment
const (
	THOST_FTDC_CSS_Entry   TThostFtdcCurrencySwapStatusType = '1' // 已录入
	THOST_FTDC_CSS_Approve TThostFtdcCurrencySwapStatusType = '2' // 已审核
	THOST_FTDC_CSS_Refuse  TThostFtdcCurrencySwapStatusType = '3' // 已拒绝
	THOST_FTDC_CSS_Revoke  TThostFtdcCurrencySwapStatusType = '4' // 已撤销
	THOST_FTDC_CSS_Send    TThostFtdcCurrencySwapStatusType = '5' // 已发送
	THOST_FTDC_CSS_Success TThostFtdcCurrencySwapStatusType = '6' // 换汇成功
	THOST_FTDC_CSS_Failure TThostFtdcCurrencySwapStatusType = '7' // 换汇失败
)

// TFtdcCurrExchCertNoType是一个凭证号类型
type TThostFtdcCurrExchCertNoType [13]byte

func (t TThostFtdcCurrExchCertNoType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcCurrExchCertNoType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcBatchSerialNoType是一个批次号类型
type TThostFtdcBatchSerialNoType [21]byte

func (t TThostFtdcBatchSerialNoType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcBatchSerialNoType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcReqFlagType是一个换汇发送标志类型
type TThostFtdcReqFlagType uint8

//go:generate stringer -type TThostFtdcReqFlagType -linecomment
const (
	THOST_FTDC_REQF_NoSend      TThostFtdcReqFlagType = '0' // 未发送
	THOST_FTDC_REQF_SendSuccess TThostFtdcReqFlagType = '1' // 发送成功
	THOST_FTDC_REQF_SendFailed  TThostFtdcReqFlagType = '2' // 发送失败
	THOST_FTDC_REQF_WaitReSend  TThostFtdcReqFlagType = '3' // 等待重发
)

// TFtdcResFlagType是一个换汇返回成功标志类型
type TThostFtdcResFlagType uint8

//go:generate stringer -type TThostFtdcResFlagType -linecomment
const (
	THOST_FTDC_RESF_Success      TThostFtdcResFlagType = '0' // 成功
	THOST_FTDC_RESF_InsuffiCient TThostFtdcResFlagType = '1' // 账户余额不足
	THOST_FTDC_RESF_UnKnown      TThostFtdcResFlagType = '8' // 交易结果未知
)

// TFtdcPageControlType是一个换汇页面控制类型
type TThostFtdcPageControlType [2]byte

func (t TThostFtdcPageControlType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcPageControlType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcRecordCountType是一个记录数类型
type TThostFtdcRecordCountType int32

// TFtdcCurrencySwapMemoType是一个换汇需确认信息类型
type TThostFtdcCurrencySwapMemoType [101]byte

func (t TThostFtdcCurrencySwapMemoType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcCurrencySwapMemoType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcExStatusType是一个修改状态类型
type TThostFtdcExStatusType uint8

//go:generate stringer -type TThostFtdcExStatusType -linecomment
const (
	THOST_FTDC_EXS_Before TThostFtdcExStatusType = '0' // 修改前
	THOST_FTDC_EXS_After  TThostFtdcExStatusType = '1' // 修改后
)

// TFtdcClientRegionType是一个开户客户地域类型
type TThostFtdcClientRegionType uint8

//go:generate stringer -type TThostFtdcClientRegionType -linecomment
const (
	THOST_FTDC_CR_Domestic TThostFtdcClientRegionType = '1' // 国内客户
	THOST_FTDC_CR_GMT      TThostFtdcClientRegionType = '2' // 港澳台客户
	THOST_FTDC_CR_Foreign  TThostFtdcClientRegionType = '3' // 国外客户
)

// TFtdcWorkPlaceType是一个工作单位类型
type TThostFtdcWorkPlaceType [101]byte

func (t TThostFtdcWorkPlaceType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcWorkPlaceType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcBusinessPeriodType是一个经营期限类型
type TThostFtdcBusinessPeriodType [21]byte

func (t TThostFtdcBusinessPeriodType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcBusinessPeriodType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcWebSiteType是一个网址类型
type TThostFtdcWebSiteType [101]byte

func (t TThostFtdcWebSiteType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcWebSiteType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcUOAIdCardTypeType是一个统一开户证件类型类型
type TThostFtdcUOAIdCardTypeType [3]byte

func (t TThostFtdcUOAIdCardTypeType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcUOAIdCardTypeType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcClientModeType是一个开户模式类型
type TThostFtdcClientModeType [3]byte

func (t TThostFtdcClientModeType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcClientModeType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcInvestorFullNameType是一个投资者全称类型
type TThostFtdcInvestorFullNameType [101]byte

func (t TThostFtdcInvestorFullNameType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcInvestorFullNameType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcUOABrokerIDType是一个境外中介机构ID类型
type TThostFtdcUOABrokerIDType [11]byte

func (t TThostFtdcUOABrokerIDType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcUOABrokerIDType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcUOAZipCodeType是一个邮政编码类型
type TThostFtdcUOAZipCodeType [11]byte

func (t TThostFtdcUOAZipCodeType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcUOAZipCodeType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcUOAEMailType是一个电子邮箱类型
type TThostFtdcUOAEMailType [101]byte

func (t TThostFtdcUOAEMailType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcUOAEMailType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcOldCityType是一个城市类型
type TThostFtdcOldCityType [41]byte

func (t TThostFtdcOldCityType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcOldCityType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcCorporateIdentifiedCardNoType是一个法人代表证件号码类型
type TThostFtdcCorporateIdentifiedCardNoType [101]byte

func (t TThostFtdcCorporateIdentifiedCardNoType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcCorporateIdentifiedCardNoType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcHasBoardType是一个是否有董事会类型
type TThostFtdcHasBoardType uint8

//go:generate stringer -type TThostFtdcHasBoardType -linecomment
const (
	THOST_FTDC_HB_No  TThostFtdcHasBoardType = '0' // 没有
	THOST_FTDC_HB_Yes TThostFtdcHasBoardType = '1' // 有
)

type TThostFtdcStartModeType uint8

// TFtdcTemplateTypeType是一个模型类型类型
type TThostFtdcTemplateTypeType uint8

//go:generate stringer -type TThostFtdcTemplateTypeType -linecomment
const (
	THOST_FTDC_TPT_Full      TThostFtdcTemplateTypeType = '1' // 全量
	THOST_FTDC_TPT_Increment TThostFtdcTemplateTypeType = '2' // 增量
	THOST_FTDC_TPT_BackUp    TThostFtdcTemplateTypeType = '3' // 备份
)

// TFtdcLoginModeType是一个登录模式类型
type TThostFtdcLoginModeType uint8

//go:generate stringer -type TThostFtdcLoginModeType -linecomment
const (
	THOST_FTDC_LM_Trade    TThostFtdcLoginModeType = '0' // 交易
	THOST_FTDC_LM_Transfer TThostFtdcLoginModeType = '1' // 转账
)

// TFtdcPromptTypeType是一个日历提示类型类型
type TThostFtdcPromptTypeType uint8

//go:generate stringer -type TThostFtdcPromptTypeType -linecomment
const (
	THOST_FTDC_CPT_Instrument TThostFtdcPromptTypeType = '1' // 合约上下市
	THOST_FTDC_CPT_Margin     TThostFtdcPromptTypeType = '2' // 保证金分段生效
)

// TFtdcLedgerManageIDType是一个分户管理资产编码类型
type TThostFtdcLedgerManageIDType [51]byte

func (t TThostFtdcLedgerManageIDType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcLedgerManageIDType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcInvestVarietyType是一个投资品种类型
type TThostFtdcInvestVarietyType [101]byte

func (t TThostFtdcInvestVarietyType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcInvestVarietyType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcBankAccountTypeType是一个账户类别类型
type TThostFtdcBankAccountTypeType [2]byte

func (t TThostFtdcBankAccountTypeType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcBankAccountTypeType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcLedgerManageBankType是一个开户银行类型
type TThostFtdcLedgerManageBankType [101]byte

func (t TThostFtdcLedgerManageBankType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcLedgerManageBankType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcCffexDepartmentNameType是一个开户营业部类型
type TThostFtdcCffexDepartmentNameType [101]byte

func (t TThostFtdcCffexDepartmentNameType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcCffexDepartmentNameType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcCffexDepartmentCodeType是一个营业部代码类型
type TThostFtdcCffexDepartmentCodeType [9]byte

func (t TThostFtdcCffexDepartmentCodeType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcCffexDepartmentCodeType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcHasTrusteeType是一个是否有托管人类型
type TThostFtdcHasTrusteeType uint8

//go:generate stringer -type TThostFtdcHasTrusteeType -linecomment
const (
	THOST_FTDC_HT_Yes TThostFtdcHasTrusteeType = '1' // 有
	THOST_FTDC_HT_No  TThostFtdcHasTrusteeType = '0' // 没有
)

// TFtdcCSRCMemo1Type是一个说明类型
type TThostFtdcCSRCMemo1Type [41]byte

func (t TThostFtdcCSRCMemo1Type) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcCSRCMemo1Type) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcAssetmgrCFullNameType是一个代理资产管理业务的期货公司全称类型
type TThostFtdcAssetmgrCFullNameType [101]byte

func (t TThostFtdcAssetmgrCFullNameType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcAssetmgrCFullNameType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcAssetmgrApprovalNOType是一个资产管理业务批文号类型
type TThostFtdcAssetmgrApprovalNOType [51]byte

func (t TThostFtdcAssetmgrApprovalNOType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcAssetmgrApprovalNOType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcAssetmgrMgrNameType是一个资产管理业务负责人姓名类型
type TThostFtdcAssetmgrMgrNameType [401]byte

func (t TThostFtdcAssetmgrMgrNameType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcAssetmgrMgrNameType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcAmTypeType是一个机构类型类型
type TThostFtdcAmTypeType uint8

//go:generate stringer -type TThostFtdcAmTypeType -linecomment
const (
	THOST_FTDC_AMT_Bank       TThostFtdcAmTypeType = '1' // 银行
	THOST_FTDC_AMT_Securities TThostFtdcAmTypeType = '2' // 证券公司
	THOST_FTDC_AMT_Fund       TThostFtdcAmTypeType = '3' // 基金公司
	THOST_FTDC_AMT_Insurance  TThostFtdcAmTypeType = '4' // 保险公司
	THOST_FTDC_AMT_Trust      TThostFtdcAmTypeType = '5' // 信托公司
	THOST_FTDC_AMT_Other      TThostFtdcAmTypeType = '9' // 其他
)

// TFtdcCSRCAmTypeType是一个机构类型类型
type TThostFtdcCSRCAmTypeType [5]byte

func (t TThostFtdcCSRCAmTypeType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcCSRCAmTypeType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcCSRCFundIOTypeType是一个出入金类型类型
type TThostFtdcCSRCFundIOTypeType uint8

//go:generate stringer -type TThostFtdcCSRCFundIOTypeType -linecomment
const (
	THOST_FTDC_CFIOT_FundIO       TThostFtdcCSRCFundIOTypeType = '0' // 出入金
	THOST_FTDC_CFIOT_SwapCurrency TThostFtdcCSRCFundIOTypeType = '1' // 银期换汇
)

// TFtdcCusAccountTypeType是一个结算账户类型类型
type TThostFtdcCusAccountTypeType uint8

//go:generate stringer -type TThostFtdcCusAccountTypeType -linecomment
const (
	THOST_FTDC_CAT_Futures          TThostFtdcCusAccountTypeType = '1' // 期货结算账户
	THOST_FTDC_CAT_AssetmgrFuture   TThostFtdcCusAccountTypeType = '2' // 纯期货资管业务下的资管结算账户
	THOST_FTDC_CAT_AssetmgrTrustee  TThostFtdcCusAccountTypeType = '3' // 综合类资管业务下的期货资管托管账户
	THOST_FTDC_CAT_AssetmgrTransfer TThostFtdcCusAccountTypeType = '4' // 综合类资管业务下的资金中转账户
)

// TFtdcCSRCNationalType是一个国籍类型
type TThostFtdcCSRCNationalType [4]byte

func (t TThostFtdcCSRCNationalType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcCSRCNationalType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcCSRCSecAgentIDType是一个二级代理ID类型
type TThostFtdcCSRCSecAgentIDType [11]byte

func (t TThostFtdcCSRCSecAgentIDType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcCSRCSecAgentIDType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcLanguageTypeType是一个通知语言类型类型
type TThostFtdcLanguageTypeType uint8

//go:generate stringer -type TThostFtdcLanguageTypeType -linecomment
const (
	THOST_FTDC_LT_Chinese TThostFtdcLanguageTypeType = '1' // 中文
	THOST_FTDC_LT_English TThostFtdcLanguageTypeType = '2' // 英文
)

// TFtdcAmAccountType是一个投资账户类型
type TThostFtdcAmAccountType [23]byte

func (t TThostFtdcAmAccountType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcAmAccountType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcAssetmgrClientTypeType是一个资产管理客户类型类型
type TThostFtdcAssetmgrClientTypeType uint8

//go:generate stringer -type TThostFtdcAssetmgrClientTypeType -linecomment
const (
	THOST_FTDC_AMCT_Person       TThostFtdcAssetmgrClientTypeType = '1' // 个人资管客户
	THOST_FTDC_AMCT_Organ        TThostFtdcAssetmgrClientTypeType = '2' // 单位资管客户
	THOST_FTDC_AMCT_SpecialOrgan TThostFtdcAssetmgrClientTypeType = '4' // 特殊单位资管客户
)

// TFtdcAssetmgrTypeType是一个投资类型类型
type TThostFtdcAssetmgrTypeType uint8

//go:generate stringer -type TThostFtdcAssetmgrTypeType -linecomment
const (
	THOST_FTDC_ASST_Futures      TThostFtdcAssetmgrTypeType = '3' // 期货类
	THOST_FTDC_ASST_SpecialOrgan TThostFtdcAssetmgrTypeType = '4' // 综合类
)

// TFtdcUOMType是一个计量单位类型
type TThostFtdcUOMType [11]byte

func (t TThostFtdcUOMType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcUOMType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcSHFEInstLifePhaseType是一个上期所合约生命周期状态类型
type TThostFtdcSHFEInstLifePhaseType [3]byte

func (t TThostFtdcSHFEInstLifePhaseType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcSHFEInstLifePhaseType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcSHFEProductClassType是一个产品类型类型
type TThostFtdcSHFEProductClassType [11]byte

func (t TThostFtdcSHFEProductClassType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcSHFEProductClassType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcPriceDecimalType是一个价格小数位类型
type TThostFtdcPriceDecimalType [2]byte

func (t TThostFtdcPriceDecimalType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcPriceDecimalType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcInTheMoneyFlagType是一个平值期权标志类型
type TThostFtdcInTheMoneyFlagType [2]byte

func (t TThostFtdcInTheMoneyFlagType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcInTheMoneyFlagType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcCheckInstrTypeType是一个合约比较类型类型
type TThostFtdcCheckInstrTypeType uint8

//go:generate stringer -type TThostFtdcCheckInstrTypeType -linecomment
const (
	THOST_FTDC_CIT_HasExch TThostFtdcCheckInstrTypeType = '0' // 合约交易所不存在
	THOST_FTDC_CIT_HasATP  TThostFtdcCheckInstrTypeType = '1' // 合约本系统不存在
	THOST_FTDC_CIT_HasDiff TThostFtdcCheckInstrTypeType = '2' // 合约比较不一致
)

// TFtdcDeliveryTypeType是一个交割类型类型
type TThostFtdcDeliveryTypeType uint8

//go:generate stringer -type TThostFtdcDeliveryTypeType -linecomment
const (
	THOST_FTDC_DT_HandDeliv   TThostFtdcDeliveryTypeType = '1' // 手工交割
	THOST_FTDC_DT_PersonDeliv TThostFtdcDeliveryTypeType = '2' // 到期交割
)

// TFtdcBigMoneyType是一个资金类型
type TThostFtdcBigMoneyType float64

func (t TThostFtdcBigMoneyType) IsValid() bool {
	return float64(t) != math.MaxFloat64
}

func (t TThostFtdcBigMoneyType) String() string {
	if t.IsValid() {
		return strconv.FormatFloat(float64(t), 'f', 6, 64)
	} else {
		return "DBL_MAX"
	}
}

// TFtdcMaxMarginSideAlgorithmType是一个大额单边保证金算法类型
type TThostFtdcMaxMarginSideAlgorithmType uint8

//go:generate stringer -type TThostFtdcMaxMarginSideAlgorithmType -linecomment
const (
	THOST_FTDC_MMSA_NO  TThostFtdcMaxMarginSideAlgorithmType = '0' // 不使用大额单边保证金算法
	THOST_FTDC_MMSA_YES TThostFtdcMaxMarginSideAlgorithmType = '1' // 使用大额单边保证金算法
)

// TFtdcDAClientTypeType是一个资产管理客户类型类型
type TThostFtdcDAClientTypeType uint8

//go:generate stringer -type TThostFtdcDAClientTypeType -linecomment
const (
	THOST_FTDC_CACT_Person  TThostFtdcDAClientTypeType = '0' // 自然人
	THOST_FTDC_CACT_Company TThostFtdcDAClientTypeType = '1' // 法人
	THOST_FTDC_CACT_Other   TThostFtdcDAClientTypeType = '2' // 其他
)

// TFtdcCombinInstrIDType是一个套利合约代码类型
type TThostFtdcCombinInstrIDType [61]byte

func (t TThostFtdcCombinInstrIDType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcCombinInstrIDType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcCombinSettlePriceType是一个各腿结算价类型
type TThostFtdcCombinSettlePriceType [61]byte

func (t TThostFtdcCombinSettlePriceType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcCombinSettlePriceType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcDCEPriorityType是一个优先级类型
type TThostFtdcDCEPriorityType int32

// TFtdcTradeGroupIDType是一个成交组号类型
type TThostFtdcTradeGroupIDType int32

// TFtdcIsCheckPrepaType是一个是否校验开户可用资金类型
type TThostFtdcIsCheckPrepaType int32

// TFtdcUOAAssetmgrTypeType是一个投资类型类型
type TThostFtdcUOAAssetmgrTypeType uint8

//go:generate stringer -type TThostFtdcUOAAssetmgrTypeType -linecomment
const (
	THOST_FTDC_UOAAT_Futures      TThostFtdcUOAAssetmgrTypeType = '1' // 期货类
	THOST_FTDC_UOAAT_SpecialOrgan TThostFtdcUOAAssetmgrTypeType = '2' // 综合类
)

// TFtdcDirectionEnType是一个买卖方向类型
type TThostFtdcDirectionEnType uint8

//go:generate stringer -type TThostFtdcDirectionEnType -linecomment
const (
	THOST_FTDC_DEN_Buy  TThostFtdcDirectionEnType = '0' // Buy
	THOST_FTDC_DEN_Sell TThostFtdcDirectionEnType = '1' // Sell
)

// TFtdcOffsetFlagEnType是一个开平标志类型
type TThostFtdcOffsetFlagEnType uint8

//go:generate stringer -type TThostFtdcOffsetFlagEnType -linecomment
const (
	THOST_FTDC_OFEN_Open            TThostFtdcOffsetFlagEnType = '0' // Position Opening
	THOST_FTDC_OFEN_Close           TThostFtdcOffsetFlagEnType = '1' // Position Close
	THOST_FTDC_OFEN_ForceClose      TThostFtdcOffsetFlagEnType = '2' // Forced Liquidation
	THOST_FTDC_OFEN_CloseToday      TThostFtdcOffsetFlagEnType = '3' // Close Today
	THOST_FTDC_OFEN_CloseYesterday  TThostFtdcOffsetFlagEnType = '4' // Close Prev.
	THOST_FTDC_OFEN_ForceOff        TThostFtdcOffsetFlagEnType = '5' // Forced Reduction
	THOST_FTDC_OFEN_LocalForceClose TThostFtdcOffsetFlagEnType = '6' // Local Forced Liquidation
)

// TFtdcHedgeFlagEnType是一个投机套保标志类型
type TThostFtdcHedgeFlagEnType uint8

//go:generate stringer -type TThostFtdcHedgeFlagEnType -linecomment
const (
	THOST_FTDC_HFEN_Speculation TThostFtdcHedgeFlagEnType = '1' // Speculation
	THOST_FTDC_HFEN_Arbitrage   TThostFtdcHedgeFlagEnType = '2' // Arbitrage
	THOST_FTDC_HFEN_Hedge       TThostFtdcHedgeFlagEnType = '3' // Hedge
)

// TFtdcFundIOTypeEnType是一个出入金类型类型
type TThostFtdcFundIOTypeEnType uint8

//go:generate stringer -type TThostFtdcFundIOTypeEnType -linecomment
const (
	THOST_FTDC_FIOTEN_FundIO       TThostFtdcFundIOTypeEnType = '1' // DepositWithdrawal
	THOST_FTDC_FIOTEN_Transfer     TThostFtdcFundIOTypeEnType = '2' // Bank-Futures Transfer
	THOST_FTDC_FIOTEN_SwapCurrency TThostFtdcFundIOTypeEnType = '3' // Bank-Futures FX Exchange
)

// TFtdcFundTypeEnType是一个资金类型类型
type TThostFtdcFundTypeEnType uint8

//go:generate stringer -type TThostFtdcFundTypeEnType -linecomment
const (
	THOST_FTDC_FTEN_Deposite      TThostFtdcFundTypeEnType = '1' // Bank Deposit
	THOST_FTDC_FTEN_ItemFund      TThostFtdcFundTypeEnType = '2' // PaymentFee
	THOST_FTDC_FTEN_Company       TThostFtdcFundTypeEnType = '3' // Brokerage Adj
	THOST_FTDC_FTEN_InnerTransfer TThostFtdcFundTypeEnType = '4' // Internal Transfer
)

// TFtdcFundDirectionEnType是一个出入金方向类型
type TThostFtdcFundDirectionEnType uint8

//go:generate stringer -type TThostFtdcFundDirectionEnType -linecomment
const (
	THOST_FTDC_FDEN_In  TThostFtdcFundDirectionEnType = '1' // Deposit
	THOST_FTDC_FDEN_Out TThostFtdcFundDirectionEnType = '2' // Withdrawal
)

// TFtdcFundMortDirectionEnType是一个货币质押方向类型
type TThostFtdcFundMortDirectionEnType uint8

//go:generate stringer -type TThostFtdcFundMortDirectionEnType -linecomment
const (
	THOST_FTDC_FMDEN_In  TThostFtdcFundMortDirectionEnType = '1' // Pledge
	THOST_FTDC_FMDEN_Out TThostFtdcFundMortDirectionEnType = '2' // Redemption
)

// TFtdcSwapBusinessTypeType是一个换汇业务种类类型
type TThostFtdcSwapBusinessTypeType [3]byte

func (t TThostFtdcSwapBusinessTypeType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcSwapBusinessTypeType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcOptionsTypeType是一个期权类型类型
type TThostFtdcOptionsTypeType uint8

//go:generate stringer -type TThostFtdcOptionsTypeType -linecomment
const (
	THOST_FTDC_CP_CallOptions TThostFtdcOptionsTypeType = '1' // 看涨
	THOST_FTDC_CP_PutOptions  TThostFtdcOptionsTypeType = '2' // 看跌
)

// TFtdcStrikeModeType是一个执行方式类型
type TThostFtdcStrikeModeType uint8

//go:generate stringer -type TThostFtdcStrikeModeType -linecomment
const (
	THOST_FTDC_STM_Continental TThostFtdcStrikeModeType = '0' // 欧式
	THOST_FTDC_STM_American    TThostFtdcStrikeModeType = '1' // 美式
	THOST_FTDC_STM_Bermuda     TThostFtdcStrikeModeType = '2' // 百慕大
)

// TFtdcStrikeTypeType是一个执行类型类型
type TThostFtdcStrikeTypeType uint8

//go:generate stringer -type TThostFtdcStrikeTypeType -linecomment
const (
	THOST_FTDC_STT_Hedge TThostFtdcStrikeTypeType = '0' // 自身对冲
	THOST_FTDC_STT_Match TThostFtdcStrikeTypeType = '1' // 匹配执行
)

// TFtdcApplyTypeType是一个中金所期权放弃执行申请类型类型
type TThostFtdcApplyTypeType uint8

//go:generate stringer -type TThostFtdcApplyTypeType -linecomment
const (
	THOST_FTDC_APPT_NotStrikeNum TThostFtdcApplyTypeType = '4' // 不执行数量
)

// TFtdcGiveUpDataSourceType是一个放弃执行申请数据来源类型
type TThostFtdcGiveUpDataSourceType uint8

//go:generate stringer -type TThostFtdcGiveUpDataSourceType -linecomment
const (
	THOST_FTDC_GUDS_Gen  TThostFtdcGiveUpDataSourceType = '0' // 系统生成
	THOST_FTDC_GUDS_Hand TThostFtdcGiveUpDataSourceType = '1' // 手工添加
)

// TFtdcExecOrderSysIDType是一个执行宣告系统编号类型
type TThostFtdcExecOrderSysIDType [21]byte

func (t TThostFtdcExecOrderSysIDType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcExecOrderSysIDType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcExecResultType是一个执行结果类型
type TThostFtdcExecResultType uint8

//go:generate stringer -type TThostFtdcExecResultType -linecomment
const (
	THOST_FTDC_OER_NoExec               TThostFtdcExecResultType = 'n' // 没有执行
	THOST_FTDC_OER_Canceled             TThostFtdcExecResultType = 'c' // 已经取消
	THOST_FTDC_OER_OK                   TThostFtdcExecResultType = '0' // 执行成功
	THOST_FTDC_OER_NoPosition           TThostFtdcExecResultType = '1' // 期权持仓不够
	THOST_FTDC_OER_NoDeposit            TThostFtdcExecResultType = '2' // 资金不够
	THOST_FTDC_OER_NoParticipant        TThostFtdcExecResultType = '3' // 会员不存在
	THOST_FTDC_OER_NoClient             TThostFtdcExecResultType = '4' // 客户不存在
	THOST_FTDC_OER_NoInstrument         TThostFtdcExecResultType = '6' // 合约不存在
	THOST_FTDC_OER_NoRight              TThostFtdcExecResultType = '7' // 没有执行权限
	THOST_FTDC_OER_InvalidVolume        TThostFtdcExecResultType = '8' // 不合理的数量
	THOST_FTDC_OER_NoEnoughHistoryTrade TThostFtdcExecResultType = '9' // 没有足够的历史成交
	THOST_FTDC_OER_Unknown              TThostFtdcExecResultType = 'a' // 未知
)

// TFtdcStrikeSequenceType是一个执行序号类型
type TThostFtdcStrikeSequenceType int32

// TFtdcStrikeTimeType是一个执行时间类型
type TThostFtdcStrikeTimeType [13]byte

func (t TThostFtdcStrikeTimeType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcStrikeTimeType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcCombinationTypeType是一个组合类型类型
type TThostFtdcCombinationTypeType uint8

//go:generate stringer -type TThostFtdcCombinationTypeType -linecomment
const (
	THOST_FTDC_COMBT_Future TThostFtdcCombinationTypeType = '0' // 期货组合
	THOST_FTDC_COMBT_BUL    TThostFtdcCombinationTypeType = '1' // 垂直价差BUL
	THOST_FTDC_COMBT_BER    TThostFtdcCombinationTypeType = '2' // 垂直价差BER
	THOST_FTDC_COMBT_STD    TThostFtdcCombinationTypeType = '3' // 跨式组合
	THOST_FTDC_COMBT_STG    TThostFtdcCombinationTypeType = '4' // 宽跨式组合
	THOST_FTDC_COMBT_PRT    TThostFtdcCombinationTypeType = '5' // 备兑组合
	THOST_FTDC_COMBT_CLD    TThostFtdcCombinationTypeType = '6' // 时间价差组合
)

// TFtdcDceCombinationTypeType是一个大商所组合类型类型
type TThostFtdcDceCombinationTypeType uint8

//go:generate stringer -type TThostFtdcDceCombinationTypeType -linecomment
const (
	THOST_FTDC_DCECOMBT_SPL TThostFtdcDceCombinationTypeType = '0' // 期货对锁组合
	THOST_FTDC_DCECOMBT_OPL TThostFtdcDceCombinationTypeType = '1' // 期权对锁组合
	THOST_FTDC_DCECOMBT_SP  TThostFtdcDceCombinationTypeType = '2' // 期货跨期组合
	THOST_FTDC_DCECOMBT_SPC TThostFtdcDceCombinationTypeType = '3' // 期货跨品种组合
	THOST_FTDC_DCECOMBT_BLS TThostFtdcDceCombinationTypeType = '4' // 买入期权垂直价差组合
	THOST_FTDC_DCECOMBT_BES TThostFtdcDceCombinationTypeType = '5' // 卖出期权垂直价差组合
	THOST_FTDC_DCECOMBT_CAS TThostFtdcDceCombinationTypeType = '6' // 期权日历价差组合
	THOST_FTDC_DCECOMBT_STD TThostFtdcDceCombinationTypeType = '7' // 期权跨式组合
	THOST_FTDC_DCECOMBT_STG TThostFtdcDceCombinationTypeType = '8' // 期权宽跨式组合
	THOST_FTDC_DCECOMBT_BFO TThostFtdcDceCombinationTypeType = '9' // 买入期货期权组合
	THOST_FTDC_DCECOMBT_SFO TThostFtdcDceCombinationTypeType = 'a' // 卖出期货期权组合
)

// TFtdcCzceCombinationTypeType是郑商所组合类型类型
type TThostFtdcCzceCombinationTypeType uint8

//go:generate stringer -type TThostFtdcCzceCombinationTypeType -linecomment
const (
	THOST_FTDC_CZCECOMBT_SPZ TThostFtdcCzceCombinationTypeType = '0' // 期货对锁
	THOST_FTDC_CZCECOMBT_SPD TThostFtdcCzceCombinationTypeType = '1' // 期货跨期
	THOST_FTDC_CZCECOMBT_IPS TThostFtdcCzceCombinationTypeType = '2' // 期货跨品种
	THOST_FTDC_CZCECOMBT_BUL TThostFtdcCzceCombinationTypeType = '3' // 看涨期权垂直价差
	THOST_FTDC_CZCECOMBT_BER TThostFtdcCzceCombinationTypeType = '4' // 看跌期权垂直价差
	THOST_FTDC_CZCECOMBT_BLT TThostFtdcCzceCombinationTypeType = '5' // 看涨期权水平价差
	THOST_FTDC_CZCECOMBT_BRT TThostFtdcCzceCombinationTypeType = '6' // 看跌期权水平价差
	THOST_FTDC_CZCECOMBT_STD TThostFtdcCzceCombinationTypeType = '7' // 期权跨式组合
	THOST_FTDC_CZCECOMBT_STG TThostFtdcCzceCombinationTypeType = '8' // 期权宽跨式组合
	THOST_FTDC_CZCECOMBT_PRT TThostFtdcCzceCombinationTypeType = '9' // 备兑
)

// TFtdcOptionRoyaltyPriceTypeType是一个期权权利金价格类型类型
type TThostFtdcOptionRoyaltyPriceTypeType uint8

//go:generate stringer -type TThostFtdcOptionRoyaltyPriceTypeType -linecomment
const (
	THOST_FTDC_ORPT_PreSettlementPrice TThostFtdcOptionRoyaltyPriceTypeType = '1' // 昨结算价
	THOST_FTDC_ORPT_OpenPrice          TThostFtdcOptionRoyaltyPriceTypeType = '4' // 开仓价
)

// TFtdcBalanceAlgorithmType是一个权益算法类型
type TThostFtdcBalanceAlgorithmType uint8

//go:generate stringer -type TThostFtdcBalanceAlgorithmType -linecomment
const (
	THOST_FTDC_BLAG_Default           TThostFtdcBalanceAlgorithmType = '1' // 不计算期权市值盈亏
	THOST_FTDC_BLAG_IncludeOptValLost TThostFtdcBalanceAlgorithmType = '2' // 计算期权市值亏损
)

// TFtdcActionTypeType是一个执行类型类型
type TThostFtdcActionTypeType uint8

//go:generate stringer -type TThostFtdcActionTypeType -linecomment
const (
	THOST_FTDC_ACTP_Exec    TThostFtdcActionTypeType = '1' // 执行
	THOST_FTDC_ACTP_Abandon TThostFtdcActionTypeType = '2' // 放弃
)

// TFtdcForQuoteStatusType是一个询价状态类型
type TThostFtdcForQuoteStatusType uint8

//go:generate stringer -type TThostFtdcForQuoteStatusType -linecomment
const (
	THOST_FTDC_FQST_Submitted TThostFtdcForQuoteStatusType = 'a' // 已经提交
	THOST_FTDC_FQST_Accepted  TThostFtdcForQuoteStatusType = 'b' // 已经接受
	THOST_FTDC_FQST_Rejected  TThostFtdcForQuoteStatusType = 'c' // 已经被拒绝
)

// TFtdcValueMethodType是一个取值方式类型
type TThostFtdcValueMethodType uint8

//go:generate stringer -type TThostFtdcValueMethodType -linecomment
const (
	THOST_FTDC_VM_Absolute TThostFtdcValueMethodType = '0' // 按绝对值
	THOST_FTDC_VM_Ratio    TThostFtdcValueMethodType = '1' // 按比率
)

// TFtdcExecOrderPositionFlagType是一个期权行权后是否保留期货头寸的标记类型
type TThostFtdcExecOrderPositionFlagType uint8

//go:generate stringer -type TThostFtdcExecOrderPositionFlagType -linecomment
const (
	THOST_FTDC_EOPF_Reserve   TThostFtdcExecOrderPositionFlagType = '0' // 保留
	THOST_FTDC_EOPF_UnReserve TThostFtdcExecOrderPositionFlagType = '1' // 不保留
)

// TFtdcExecOrderCloseFlagType是一个期权行权后生成的头寸是否自动平仓类型
type TThostFtdcExecOrderCloseFlagType uint8

//go:generate stringer -type TThostFtdcExecOrderCloseFlagType -linecomment
const (
	THOST_FTDC_EOCF_AutoClose  TThostFtdcExecOrderCloseFlagType = '0' // 自动平仓
	THOST_FTDC_EOCF_NotToClose TThostFtdcExecOrderCloseFlagType = '1' // 免于自动平仓
)

// TFtdcProductTypeType是一个产品类型类型
type TThostFtdcProductTypeType uint8

//go:generate stringer -type TThostFtdcProductTypeType -linecomment
const (
	THOST_FTDC_PTE_Futures TThostFtdcProductTypeType = '1' // 期货
	THOST_FTDC_PTE_Options TThostFtdcProductTypeType = '2' // 期权
)

// TFtdcCZCEUploadFileNameType是一个郑商所结算文件名类型
type TThostFtdcCZCEUploadFileNameType uint8

//go:generate stringer -type TThostFtdcCZCEUploadFileNameType -linecomment
const (
	THOST_FTDC_CUFN_CUFN_O TThostFtdcCZCEUploadFileNameType = 'O' // ^\d{8}_zz_\d{4}
	THOST_FTDC_CUFN_CUFN_T TThostFtdcCZCEUploadFileNameType = 'T' // ^\d{8}成交表
	THOST_FTDC_CUFN_CUFN_P TThostFtdcCZCEUploadFileNameType = 'P' // ^\d{8}单腿持仓表new
	THOST_FTDC_CUFN_CUFN_N TThostFtdcCZCEUploadFileNameType = 'N' // ^\d{8}非平仓了结表
	THOST_FTDC_CUFN_CUFN_L TThostFtdcCZCEUploadFileNameType = 'L' // ^\d{8}平仓表
	THOST_FTDC_CUFN_CUFN_F TThostFtdcCZCEUploadFileNameType = 'F' // ^\d{8}资金表
	THOST_FTDC_CUFN_CUFN_C TThostFtdcCZCEUploadFileNameType = 'C' // ^\d{8}组合持仓表
	THOST_FTDC_CUFN_CUFN_M TThostFtdcCZCEUploadFileNameType = 'M' // ^\d{8}保证金参数表
)

// TFtdcDCEUploadFileNameType是一个大商所结算文件名类型
type TThostFtdcDCEUploadFileNameType uint8

//go:generate stringer -type TThostFtdcDCEUploadFileNameType -linecomment
const (
	THOST_FTDC_DUFN_DUFN_O TThostFtdcDCEUploadFileNameType = 'O' // ^\d{8}_dl_\d{3}
	THOST_FTDC_DUFN_DUFN_T TThostFtdcDCEUploadFileNameType = 'T' // ^\d{8}_成交表
	THOST_FTDC_DUFN_DUFN_P TThostFtdcDCEUploadFileNameType = 'P' // ^\d{8}_持仓表
	THOST_FTDC_DUFN_DUFN_F TThostFtdcDCEUploadFileNameType = 'F' // ^\d{8}_资金结算表
	THOST_FTDC_DUFN_DUFN_C TThostFtdcDCEUploadFileNameType = 'C' // ^\d{8}_优惠组合持仓明细表
	THOST_FTDC_DUFN_DUFN_D TThostFtdcDCEUploadFileNameType = 'D' // ^\d{8}_持仓明细表
	THOST_FTDC_DUFN_DUFN_M TThostFtdcDCEUploadFileNameType = 'M' // ^\d{8}_保证金参数表
	THOST_FTDC_DUFN_DUFN_S TThostFtdcDCEUploadFileNameType = 'S' // ^\d{8}_期权执行表
)

// TFtdcSHFEUploadFileNameType是一个上期所结算文件名类型
type TThostFtdcSHFEUploadFileNameType uint8

//go:generate stringer -type TThostFtdcSHFEUploadFileNameType -linecomment
const (
	THOST_FTDC_SUFN_SUFN_O TThostFtdcSHFEUploadFileNameType = 'O' // ^\d{4}_\d{8}_\d{8}_DailyFundChg
	THOST_FTDC_SUFN_SUFN_T TThostFtdcSHFEUploadFileNameType = 'T' // ^\d{4}_\d{8}_\d{8}_Trade
	THOST_FTDC_SUFN_SUFN_P TThostFtdcSHFEUploadFileNameType = 'P' // ^\d{4}_\d{8}_\d{8}_SettlementDetail
	THOST_FTDC_SUFN_SUFN_F TThostFtdcSHFEUploadFileNameType = 'F' // ^\d{4}_\d{8}_\d{8}_Capital
)

// TFtdcCFFEXUploadFileNameType是一个中金所结算文件名类型
type TThostFtdcCFFEXUploadFileNameType uint8

//go:generate stringer -type TThostFtdcCFFEXUploadFileNameType -linecomment
const (
	THOST_FTDC_CFUFN_SUFN_T TThostFtdcCFFEXUploadFileNameType = 'T' // ^\d{4}_SG\d{1}_\d{8}_\d{1}_Trade
	THOST_FTDC_CFUFN_SUFN_P TThostFtdcCFFEXUploadFileNameType = 'P' // ^\d{4}_SG\d{1}_\d{8}_\d{1}_SettlementDetail
	THOST_FTDC_CFUFN_SUFN_F TThostFtdcCFFEXUploadFileNameType = 'F' // ^\d{4}_SG\d{1}_\d{8}_\d{1}_Capital
	THOST_FTDC_CFUFN_SUFN_S TThostFtdcCFFEXUploadFileNameType = 'S' // ^\d{4}_SG\d{1}_\d{8}_\d{1}_OptionExec
)

// TFtdcCombDirectionType是一个组合指令方向类型
type TThostFtdcCombDirectionType uint8

//go:generate stringer -type TThostFtdcCombDirectionType -linecomment
const (
	THOST_FTDC_CMDR_Comb    TThostFtdcCombDirectionType = '0' // 申请组合
	THOST_FTDC_CMDR_UnComb  TThostFtdcCombDirectionType = '1' // 申请拆分
	THOST_FTDC_CMDR_DelComb TThostFtdcCombDirectionType = '2' // 操作员删组合单
)

// TFtdcRequestTypeType是一个郑商所套利套保确认申请类型
type TThostFtdcRequestTypeType uint8

//go:generate stringer -type TThostFtdcRequestTypeType -linecomment
const (
	THOST_FTDC_CONFREQ_ARBI  TThostFtdcRequestTypeType = '1' // 套利
	THOST_FTDC_CONFREQ_HEDGE TThostFtdcRequestTypeType = '2' // 套期保值
)

// TFtdcStrikeOffsetTypeType是一个行权偏移类型类型
type TThostFtdcStrikeOffsetTypeType uint8

//go:generate stringer -type TThostFtdcStrikeOffsetTypeType -linecomment
const (
	THOST_FTDC_STOV_RealValue   TThostFtdcStrikeOffsetTypeType = '1' // 实值额
	THOST_FTDC_STOV_ProfitValue TThostFtdcStrikeOffsetTypeType = '2' // 盈利额
	THOST_FTDC_STOV_RealRatio   TThostFtdcStrikeOffsetTypeType = '3' // 实值比例
	THOST_FTDC_STOV_ProfitRatio TThostFtdcStrikeOffsetTypeType = '4' // 盈利比例
)

// TFtdcReserveOpenAccStasType是一个预约开户状态类型
type TThostFtdcReserveOpenAccStasType uint8

//go:generate stringer -type TThostFtdcReserveOpenAccStasType -linecomment
const (
	THOST_FTDC_ROAST_Processing TThostFtdcReserveOpenAccStasType = '0' // 等待处理中
	THOST_FTDC_ROAST_Cancelled  TThostFtdcReserveOpenAccStasType = '1' // 已撤销
	THOST_FTDC_ROAST_Opened     TThostFtdcReserveOpenAccStasType = '2' // 已开户
	THOST_FTDC_ROAST_Invalid    TThostFtdcReserveOpenAccStasType = '3' // 无效请求
)

// TFtdcLoginRemarkType是一个登录备注类型
type TThostFtdcLoginRemarkType [36]byte

func (t TThostFtdcLoginRemarkType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcLoginRemarkType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcInvestUnitIDType是一个投资单元代码类型
type TThostFtdcInvestUnitIDType [17]byte

func (t TThostFtdcInvestUnitIDType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcInvestUnitIDType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcBulletinIDType是一个公告编号类型
type TThostFtdcBulletinIDType int32

// TFtdcNewsTypeType是一个公告类型类型
type TThostFtdcNewsTypeType [3]byte

func (t TThostFtdcNewsTypeType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcNewsTypeType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcNewsUrgencyType是一个紧急程度类型
type TThostFtdcNewsUrgencyType uint8

// TFtdcAbstractType是一个消息摘要类型
type TThostFtdcAbstractType [81]byte

func (t TThostFtdcAbstractType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcAbstractType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcComeFromType是一个消息来源类型
type TThostFtdcComeFromType [21]byte

func (t TThostFtdcComeFromType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcComeFromType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcURLLinkType是一个WEB地址类型
type TThostFtdcURLLinkType [201]byte

func (t TThostFtdcURLLinkType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcURLLinkType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcLongIndividualNameType是一个长个人姓名类型
type TThostFtdcLongIndividualNameType [161]byte

func (t TThostFtdcLongIndividualNameType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcLongIndividualNameType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcLongFBEBankAccountNameType是一个长换汇银行账户名类型
type TThostFtdcLongFBEBankAccountNameType [161]byte

func (t TThostFtdcLongFBEBankAccountNameType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcLongFBEBankAccountNameType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TFtdcOptSelfCloseFlagType是一个期权行权的头寸是否自对冲类型
type TThostFtdcOptSelfCloseFlagType uint8

//go:generate stringer -type TThostFtdcOptSelfCloseFlagType -linecomment
const (
	THOST_FTDC_OSCF_CloseSelfOptionPosition     TThostFtdcOptSelfCloseFlagType = '1' // 自对冲期权仓位
	THOST_FTDC_OSCF_ReserveOptionPosition       TThostFtdcOptSelfCloseFlagType = '2' // 保留期权仓位
	THOST_FTDC_OSCF_SellCloseSelfFuturePosition TThostFtdcOptSelfCloseFlagType = '3' // 自对冲卖方履约后的期货仓位
	THOST_FTDC_OSCF_ReserveFuturePosition       TThostFtdcOptSelfCloseFlagType = '4' // 保留卖方履约后的期货仓位
)

// TThostFtdcAddressAndPortType是一个IP地址和端口类型
type TThostFtdcAddressAndPortType [33]byte

func (t TThostFtdcAddressAndPortType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcAddressAndPortType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TThostFtdcSystemInfoLenType是用户端系统内部信息长度类型
type TThostFtdcSystemInfoLenType int32

// TThostFtdcClientSystemInfoType是用户端系统内部信息类型
type TThostFtdcClientSystemInfoType [273]byte

func (t TThostFtdcClientSystemInfoType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcClientSystemInfoType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TThostFtdcAppIDType是App代码类型
type TThostFtdcClientAppIDType [33]byte

func (t TThostFtdcClientAppIDType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcClientAppIDType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TThostFtdcAutoCodeType是AutoCode代码类型
type TThostFtdcAutoCodeType [17]byte

func (t TThostFtdcAutoCodeType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcAutoCodeType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TThostFtdcCustomOrderRefType是报单自定义字段类型类型
type TThostFtdcCustomOrderRefType int64

// TThostFtdcRULECommodityGroupIDType是RULE商品群号类型
type TThostFtdcRULECommodityGroupIDType int32

// TThostFtdcProductIDType是一个产品代码类型
type TThostFtdcProductIDType [81]byte

func (t TThostFtdcProductIDType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcProductIDType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

type TThostFtdcRULEInstrClassType uint8

type TThostFtdcOrderFlagType [21]byte

func (t TThostFtdcOrderFlagType) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *TThostFtdcOrderFlagType) SetString(v string) int {
	return ctp4go.SetCString(([]byte)((*t)[:]), v)
}

// TThostFtdcOperateType是一个业务操作类型
type TThostFtdcOperateType uint8

type TThostFtdcOffsetTypeType uint8

// TFtdcApplySrcType是一个申请来源类型
type TThostFtdcApplySrcType uint8

//go:generate stringer -type TThostFtdcApplySrcType -linecomment
const (
	THOST_FTDC_AS_Trade  TThostFtdcApplySrcType = '0' // 交易
	THOST_FTDC_AS_Member TThostFtdcApplySrcType = '1' // 会服
)

// TFtdcIPACLCheckModeType是IP访问控制校验模式类型
type TThostFtdcIPACLCheckModeType uint8

//go:generate stringer -type TThostFtdcIPACLCheckModeType -linecomment
const (
	THOST_FTDC_IPAC_IP     TThostFtdcIPACLCheckModeType = 'I' // IP模式
	THOST_FTDC_IPAC_UserID TThostFtdcIPACLCheckModeType = 'U' // UserID模式
)
