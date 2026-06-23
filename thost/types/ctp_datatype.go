package types

import (
	"strings"
)

//go:generate stringer -type THOST_TE_RESUME_TYPE -linecomment
type THOST_TE_RESUME_TYPE int32

const (
	THOST_TERT_RESTART            THOST_TE_RESUME_TYPE = 0 // 重传
	THOST_TERT_RESUME             THOST_TE_RESUME_TYPE = 1 // 续传
	THOST_TERT_QUICK              THOST_TE_RESUME_TYPE = 2 // 快速
	THOST_TERT_NONE               THOST_TE_RESUME_TYPE = 3 // 不传输
	THOST_TERT_RESUME_FROM_SEQ_NO THOST_TE_RESUME_TYPE = 4 // 序号重传
)

// 交易所交易员代码类型
type TThostFtdcTraderIDType Byte21

// 投资者代码类型
type TThostFtdcInvestorIDType Byte13

// 经纪公司代码类型
type TThostFtdcBrokerIDType Byte11

// 经纪公司简称类型
type TThostFtdcBrokerAbbrType Byte9

// 经纪公司名称类型
type TThostFtdcBrokerNameType Byte81

// 合约在交易所的代码类型
type TThostFtdcOldExchangeInstIDType Byte31

// 合约在交易所的代码类型
type TThostFtdcExchangeInstIDType Byte81

// 报单引用类型
type TThostFtdcOrderRefType Byte13

// 会员代码类型
type TThostFtdcParticipantIDType Byte11

// 用户代码类型
type TThostFtdcUserIDType Byte16

// 密码类型
type TThostFtdcPasswordType Byte41

// 设备标识类型
type TThostFtdcDeviceTagType Byte41

// 交易编码类型
type TThostFtdcClientIDType Byte11

// 合约代码类型
type TThostFtdcInstrumentIDType Byte81

// 合约代码类型
type TThostFtdcOldInstrumentIDType Byte31

// 合约标识码类型
type TThostFtdcInstrumentCodeType Byte31

// 市场代码类型
type TThostFtdcMarketIDType Byte31

// 产品名称类型
type TThostFtdcProductNameType Byte21

// 交易所代码类型
type TThostFtdcExchangeIDType Byte9

// 交易所名称类型
type TThostFtdcExchangeNameType Byte61

// 交易所简称类型
type TThostFtdcExchangeAbbrType Byte9

// 交易所标志类型
type TThostFtdcExchangeFlagType Byte2

// Mac地址类型
type TThostFtdcMacAddressType Byte21

// 系统编号类型
type TThostFtdcSystemIDType Byte21

// 客户登录备注2类型
type TThostFtdcClientLoginRemarkType Byte151

// 交易所属性类型
//
//go:generate stringer -type TThostFtdcExchangePropertyType -linecomment
type TThostFtdcExchangePropertyType byte

const (
	THOST_FTDC_EXP_Normal          TThostFtdcExchangePropertyType = '0' // 正常
	THOST_FTDC_EXP_GenOrderByTrade TThostFtdcExchangePropertyType = '1' // 根据成交生成报单
)

// 日期类型
type TThostFtdcDateType Byte9

// 时间类型
type TThostFtdcTimeType Byte9

// 长时间类型
type TThostFtdcLongTimeType Byte13

// 合约名称类型
type TThostFtdcInstrumentNameType Byte21

// 结算组代码类型
type TThostFtdcSettlementGroupIDType Byte9

// 报单编号类型
type TThostFtdcOrderSysIDType Byte21

// 成交编号类型
type TThostFtdcTradeIDType Byte21

// DB命令类型类型
type TThostFtdcCommandTypeType Byte65

// IP地址类型
type TThostFtdcOldIPAddressType Byte16

// IP地址类型
type TThostFtdcIPAddressType Byte33

// IP端口类型
type TThostFtdcIPPortType int32

// 产品信息类型
type TThostFtdcProductInfoType Byte11

// 协议信息类型
type TThostFtdcProtocolInfoType Byte11

// 业务单元类型
type TThostFtdcBusinessUnitType Byte21

// 出入金流水号类型
type TThostFtdcDepositSeqNoType Byte15

// 证件号码类型
type TThostFtdcIdentifiedCardNoType Byte51

// 证件类型类型
//
//go:generate stringer -type TThostFtdcIdCardTypeType -linecomment
type TThostFtdcIdCardTypeType byte

const (
	THOST_FTDC_ICT_EID                     TThostFtdcIdCardTypeType = '0' // 组织机构代码
	THOST_FTDC_ICT_IDCard                  TThostFtdcIdCardTypeType = '1' // 中国公民身份证
	THOST_FTDC_ICT_OfficerIDCard           TThostFtdcIdCardTypeType = '2' // 军官证
	THOST_FTDC_ICT_PoliceIDCard            TThostFtdcIdCardTypeType = '3' // 警官证
	THOST_FTDC_ICT_SoldierIDCard           TThostFtdcIdCardTypeType = '4' // 士兵证
	THOST_FTDC_ICT_HouseholdRegister       TThostFtdcIdCardTypeType = '5' // 户口簿
	THOST_FTDC_ICT_Passport                TThostFtdcIdCardTypeType = '6' // 护照
	THOST_FTDC_ICT_TaiwanCompatriotIDCard  TThostFtdcIdCardTypeType = '7' // 台胞证
	THOST_FTDC_ICT_HomeComingCard          TThostFtdcIdCardTypeType = '8' // 回乡证
	THOST_FTDC_ICT_LicenseNo               TThostFtdcIdCardTypeType = '9' // 营业执照号
	THOST_FTDC_ICT_TaxNo                   TThostFtdcIdCardTypeType = 'A' // 税务登记号/当地纳税ID
	THOST_FTDC_ICT_HMMainlandTravelPermit  TThostFtdcIdCardTypeType = 'B' // 港澳居民来往内地通行证
	THOST_FTDC_ICT_TwMainlandTravelPermit  TThostFtdcIdCardTypeType = 'C' // 台湾居民来往大陆通行证
	THOST_FTDC_ICT_DrivingLicense          TThostFtdcIdCardTypeType = 'D' // 驾照
	THOST_FTDC_ICT_SocialID                TThostFtdcIdCardTypeType = 'F' // 当地社保ID
	THOST_FTDC_ICT_LocalID                 TThostFtdcIdCardTypeType = 'G' // 当地身份证
	THOST_FTDC_ICT_BusinessRegistration    TThostFtdcIdCardTypeType = 'H' // 商业登记证
	THOST_FTDC_ICT_HKMCIDCard              TThostFtdcIdCardTypeType = 'I' // 港澳永久性居民身份证
	THOST_FTDC_ICT_AccountsPermits         TThostFtdcIdCardTypeType = 'J' // 人行开户许可证
	THOST_FTDC_ICT_FrgPrmtRdCard           TThostFtdcIdCardTypeType = 'K' // 外国人永久居留证
	THOST_FTDC_ICT_CptMngPrdLetter         TThostFtdcIdCardTypeType = 'L' // 资管产品备案函
	THOST_FTDC_ICT_HKMCTwResidencePermit   TThostFtdcIdCardTypeType = 'M' // 港澳台居民居住证
	THOST_FTDC_ICT_UniformSocialCreditCode TThostFtdcIdCardTypeType = 'N' // 统一社会信用代码
	THOST_FTDC_ICT_CorporationCertNo       TThostFtdcIdCardTypeType = 'O' // 机构成立证明文件
	THOST_FTDC_ICT_OtherCard               TThostFtdcIdCardTypeType = 'x' // 其他证件
)

// 本地报单编号类型
type TThostFtdcOrderLocalIDType Byte13

// 用户名称类型
type TThostFtdcUserNameType Byte81

// 参与人名称类型
type TThostFtdcPartyNameType Byte81

// 错误信息类型
type TThostFtdcErrorMsgType Byte81

// 字段名类型
type TThostFtdcFieldNameType Byte2049

// 字段内容类型
type TThostFtdcFieldContentType Byte2049

// 系统名称类型
type TThostFtdcSystemNameType Byte41

// 消息正文类型
type TThostFtdcContentType Byte501

// 投资者范围类型
//
//go:generate stringer -type TThostFtdcInvestorRangeType -linecomment
type TThostFtdcInvestorRangeType byte

const (
	THOST_FTDC_IR_All    TThostFtdcInvestorRangeType = '1' // 所有
	THOST_FTDC_IR_Group  TThostFtdcInvestorRangeType = '2' // 投资者组
	THOST_FTDC_IR_Single TThostFtdcInvestorRangeType = '3' // 单一投资者
)

// 投资者范围类型
//
//go:generate stringer -type TThostFtdcDepartmentRangeType -linecomment
type TThostFtdcDepartmentRangeType byte

const (
	THOST_FTDC_DR_All    TThostFtdcDepartmentRangeType = '1' // 所有
	THOST_FTDC_DR_Group  TThostFtdcDepartmentRangeType = '2' // 组织架构
	THOST_FTDC_DR_Single TThostFtdcDepartmentRangeType = '3' // 单一投资者
)

// 数据同步状态类型
//
//go:generate stringer -type TThostFtdcDataSyncStatusType -linecomment
type TThostFtdcDataSyncStatusType byte

const (
	THOST_FTDC_DS_Asynchronous  TThostFtdcDataSyncStatusType = '1' // 未同步
	THOST_FTDC_DS_Synchronizing TThostFtdcDataSyncStatusType = '2' // 同步中
	THOST_FTDC_DS_Synchronized  TThostFtdcDataSyncStatusType = '3' // 已同步
)

// 经纪公司数据同步状态类型
//
//go:generate stringer -type TThostFtdcBrokerDataSyncStatusType -linecomment
type TThostFtdcBrokerDataSyncStatusType byte

const (
	THOST_FTDC_BDS_Synchronized  TThostFtdcBrokerDataSyncStatusType = '1' // 已同步
	THOST_FTDC_BDS_Synchronizing TThostFtdcBrokerDataSyncStatusType = '2' // 同步中
)

// 交易所连接状态类型
//
//go:generate stringer -type TThostFtdcExchangeConnectStatusType -linecomment
type TThostFtdcExchangeConnectStatusType byte

const (
	THOST_FTDC_ECS_NoConnection      TThostFtdcExchangeConnectStatusType = '1' // 没有任何连接
	THOST_FTDC_ECS_QryInstrumentSent TThostFtdcExchangeConnectStatusType = '2' // 已经发出合约查询请求
	THOST_FTDC_ECS_GotInformation    TThostFtdcExchangeConnectStatusType = '9' // 已经获取信息
)

// 交易所交易员连接状态类型
//
//go:generate stringer -type TThostFtdcTraderConnectStatusType -linecomment
type TThostFtdcTraderConnectStatusType byte

const (
	THOST_FTDC_TCS_NotConnected      TThostFtdcTraderConnectStatusType = '1' // 没有任何连接
	THOST_FTDC_TCS_Connected         TThostFtdcTraderConnectStatusType = '2' // 已经连接
	THOST_FTDC_TCS_QryInstrumentSent TThostFtdcTraderConnectStatusType = '3' // 已经发出合约查询请求
	THOST_FTDC_TCS_SubPrivateFlow    TThostFtdcTraderConnectStatusType = '4' // 订阅私有流
)

// 功能代码类型
//
//go:generate stringer -type TThostFtdcFunctionCodeType -linecomment
type TThostFtdcFunctionCodeType byte

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
	THOST_FTDC_FC_ExitEmergency          TThostFtdcFunctionCodeType = 'G' // 退出紧急状态
)

// 经纪公司功能代码类型
//
//go:generate stringer -type TThostFtdcBrokerFunctionCodeType -linecomment
type TThostFtdcBrokerFunctionCodeType byte

const (
	THOST_FTDC_BFC_ForceUserLogout    TThostFtdcBrokerFunctionCodeType = '1' // 强制用户登出
	THOST_FTDC_BFC_UserPasswordUpdate TThostFtdcBrokerFunctionCodeType = '2' // 变更用户口令
	THOST_FTDC_BFC_SyncBrokerData     TThostFtdcBrokerFunctionCodeType = '3' // 同步经纪公司数据
	THOST_FTDC_BFC_BachSyncBrokerData TThostFtdcBrokerFunctionCodeType = '4' // 批量同步经纪公司数据
	THOST_FTDC_BFC_OrderInsert        TThostFtdcBrokerFunctionCodeType = '5' // 报单插入
	THOST_FTDC_BFC_OrderAction        TThostFtdcBrokerFunctionCodeType = '6' // 报单操作
	THOST_FTDC_BFC_AllQuery           TThostFtdcBrokerFunctionCodeType = '7' // 全部查询
	THOST_FTDC_BFC_log                TThostFtdcBrokerFunctionCodeType = 'a' // 系统功能：登入/登出/修改密码等
	THOST_FTDC_BFC_BaseQry            TThostFtdcBrokerFunctionCodeType = 'b' // 基本查询：查询基础数据，如合约，交易所等常量
	THOST_FTDC_BFC_TradeQry           TThostFtdcBrokerFunctionCodeType = 'c' // 交易查询：如查成交，委托
	THOST_FTDC_BFC_Trade              TThostFtdcBrokerFunctionCodeType = 'd' // 交易功能：报单，撤单
	THOST_FTDC_BFC_Virement           TThostFtdcBrokerFunctionCodeType = 'e' // 银期转账
	THOST_FTDC_BFC_Risk               TThostFtdcBrokerFunctionCodeType = 'f' // 风险监控
	THOST_FTDC_BFC_Session            TThostFtdcBrokerFunctionCodeType = 'g' // 查询/管理：查询会话，踢人等
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
	THOST_FTDC_BFC_Designate          TThostFtdcBrokerFunctionCodeType = 'N' // 指定
	THOST_FTDC_BFC_StockDisposal      TThostFtdcBrokerFunctionCodeType = 'O' // 证券处置
	THOST_FTDC_BFC_BrokerDepositWarn  TThostFtdcBrokerFunctionCodeType = 'Q' // 席位资金预警
	THOST_FTDC_BFC_CoverWarn          TThostFtdcBrokerFunctionCodeType = 'S' // 备兑不足预警
	THOST_FTDC_BFC_PreExecOrder       TThostFtdcBrokerFunctionCodeType = 'T' // 行权试算
	THOST_FTDC_BFC_ExecOrderRisk      TThostFtdcBrokerFunctionCodeType = 'P' // 行权交收风险
	THOST_FTDC_BFC_PosiLimitWarn      TThostFtdcBrokerFunctionCodeType = 'U' // 持仓限额预警
	THOST_FTDC_BFC_QryPosiLimit       TThostFtdcBrokerFunctionCodeType = 'V' // 持仓限额查询
	THOST_FTDC_BFC_FBSign             TThostFtdcBrokerFunctionCodeType = 'W' // 银期签到签退
	THOST_FTDC_BFC_FBAccount          TThostFtdcBrokerFunctionCodeType = 'X' // 银期签约解约
)

// 报单操作状态类型
//
//go:generate stringer -type TThostFtdcOrderActionStatusType -linecomment
type TThostFtdcOrderActionStatusType byte

const (
	THOST_FTDC_OAS_Submitted TThostFtdcOrderActionStatusType = 'a' // 已经提交
	THOST_FTDC_OAS_Accepted  TThostFtdcOrderActionStatusType = 'b' // 已经接受
	THOST_FTDC_OAS_Rejected  TThostFtdcOrderActionStatusType = 'c' // 已经被拒绝
)

// 报单状态类型
//
//go:generate stringer -type TThostFtdcOrderStatusType -linecomment
type TThostFtdcOrderStatusType byte

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

// 报单提交状态类型
//
//go:generate stringer -type TThostFtdcOrderSubmitStatusType -linecomment
type TThostFtdcOrderSubmitStatusType byte

const (
	THOST_FTDC_OSS_InsertSubmitted TThostFtdcOrderSubmitStatusType = '0' // 已经提交
	THOST_FTDC_OSS_CancelSubmitted TThostFtdcOrderSubmitStatusType = '1' // 撤单已经提交
	THOST_FTDC_OSS_ModifySubmitted TThostFtdcOrderSubmitStatusType = '2' // 修改已经提交
	THOST_FTDC_OSS_Accepted        TThostFtdcOrderSubmitStatusType = '3' // 已经接受
	THOST_FTDC_OSS_InsertRejected  TThostFtdcOrderSubmitStatusType = '4' // 报单已经被拒绝
	THOST_FTDC_OSS_CancelRejected  TThostFtdcOrderSubmitStatusType = '5' // 撤单已经被拒绝
	THOST_FTDC_OSS_ModifyRejected  TThostFtdcOrderSubmitStatusType = '6' // 改单已经被拒绝
)

// 持仓日期类型
//
//go:generate stringer -type TThostFtdcPositionDateType -linecomment
type TThostFtdcPositionDateType byte

const (
	THOST_FTDC_PSD_Today   TThostFtdcPositionDateType = '1' // 今日持仓
	THOST_FTDC_PSD_History TThostFtdcPositionDateType = '2' // 历史持仓
)

// 持仓日期类型类型
//
//go:generate stringer -type TThostFtdcPositionDateTypeType -linecomment
type TThostFtdcPositionDateTypeType byte

const (
	THOST_FTDC_PDT_UseHistory   TThostFtdcPositionDateTypeType = '1' // 使用历史持仓
	THOST_FTDC_PDT_NoUseHistory TThostFtdcPositionDateTypeType = '2' // 不使用历史持仓
)

// 交易角色类型
//
//go:generate stringer -type TThostFtdcTradingRoleType -linecomment
type TThostFtdcTradingRoleType byte

const (
	THOST_FTDC_ER_Broker TThostFtdcTradingRoleType = '1' // 代理
	THOST_FTDC_ER_Host   TThostFtdcTradingRoleType = '2' // 自营
	THOST_FTDC_ER_Maker  TThostFtdcTradingRoleType = '3' // 做市商
)

// 产品类型类型
//
//go:generate stringer -type TThostFtdcProductClassType -linecomment
type TThostFtdcProductClassType byte

const (
	THOST_FTDC_PC_Futures     TThostFtdcProductClassType = '1' // 期货
	THOST_FTDC_PC_Options     TThostFtdcProductClassType = '2' // 期货期权
	THOST_FTDC_PC_Combination TThostFtdcProductClassType = '3' // 组合
	THOST_FTDC_PC_Spot        TThostFtdcProductClassType = '4' // 即期
	THOST_FTDC_PC_EFP         TThostFtdcProductClassType = '5' // 期转现
	THOST_FTDC_PC_SpotOption  TThostFtdcProductClassType = '6' // 现货期权
	THOST_FTDC_PC_TAS         TThostFtdcProductClassType = '7' // TAS合约
	THOST_FTDC_PC_MI          TThostFtdcProductClassType = 'I' // 金属指数
)

// 产品类型类型
//
//go:generate stringer -type TThostFtdcAPIProductClassType -linecomment
type TThostFtdcAPIProductClassType byte

const (
	THOST_FTDC_APC_FutureSingle  TThostFtdcAPIProductClassType = '1' // 期货单一合约
	THOST_FTDC_APC_OptionSingle  TThostFtdcAPIProductClassType = '2' // 期权单一合约
	THOST_FTDC_APC_Futures       TThostFtdcAPIProductClassType = '3' // 可交易期货(含期货组合和期货单一合约)
	THOST_FTDC_APC_Options       TThostFtdcAPIProductClassType = '4' // 可交易期权(含期权组合和期权单一合约)
	THOST_FTDC_APC_TradingComb   TThostFtdcAPIProductClassType = '5' // 可下单套利组合
	THOST_FTDC_APC_UnTradingComb TThostFtdcAPIProductClassType = '6' // 可申请的组合（可以申请的组合合约 包含可以交易的合约）
	THOST_FTDC_APC_AllTrading    TThostFtdcAPIProductClassType = '7' // 所有可以交易合约
	THOST_FTDC_APC_All           TThostFtdcAPIProductClassType = '8' // 所有合约（包含不能交易合约 慎用）
)

// 合约生命周期状态类型
//
//go:generate stringer -type TThostFtdcInstLifePhaseType -linecomment
type TThostFtdcInstLifePhaseType byte

const (
	THOST_FTDC_IP_NotStart TThostFtdcInstLifePhaseType = '0' // 未上市
	THOST_FTDC_IP_Started  TThostFtdcInstLifePhaseType = '1' // 上市
	THOST_FTDC_IP_Pause    TThostFtdcInstLifePhaseType = '2' // 停牌
	THOST_FTDC_IP_Expired  TThostFtdcInstLifePhaseType = '3' // 到期
)

// 买卖方向类型
//
//go:generate stringer -type TThostFtdcDirectionType -linecomment
type TThostFtdcDirectionType byte

const (
	THOST_FTDC_D_Buy  TThostFtdcDirectionType = '0' // 买
	THOST_FTDC_D_Sell TThostFtdcDirectionType = '1' // 卖
)

// 持仓类型类型
//
//go:generate stringer -type TThostFtdcPositionTypeType -linecomment
type TThostFtdcPositionTypeType byte

const (
	THOST_FTDC_PT_Net   TThostFtdcPositionTypeType = '1' // 净持仓
	THOST_FTDC_PT_Gross TThostFtdcPositionTypeType = '2' // 综合持仓
)

// 持仓多空方向类型
//
//go:generate stringer -type TThostFtdcPosiDirectionType -linecomment
type TThostFtdcPosiDirectionType byte

const (
	THOST_FTDC_PD_Net   TThostFtdcPosiDirectionType = '1' // 净
	THOST_FTDC_PD_Long  TThostFtdcPosiDirectionType = '2' // 多头
	THOST_FTDC_PD_Short TThostFtdcPosiDirectionType = '3' // 空头
)

// 系统结算状态类型
//
//go:generate stringer -type TThostFtdcSysSettlementStatusType -linecomment
type TThostFtdcSysSettlementStatusType byte

const (
	THOST_FTDC_SS_NonActive          TThostFtdcSysSettlementStatusType = '1' // 不活跃
	THOST_FTDC_SS_Startup            TThostFtdcSysSettlementStatusType = '2' // 启动
	THOST_FTDC_SS_Operating          TThostFtdcSysSettlementStatusType = '3' // 操作
	THOST_FTDC_SS_Settlement         TThostFtdcSysSettlementStatusType = '4' // 结算
	THOST_FTDC_SS_SettlementFinished TThostFtdcSysSettlementStatusType = '5' // 结算完成
)

// 费率属性类型
//
//go:generate stringer -type TThostFtdcRatioAttrType -linecomment
type TThostFtdcRatioAttrType byte

const (
	THOST_FTDC_RA_Trade      TThostFtdcRatioAttrType = '0' // 交易费率
	THOST_FTDC_RA_Settlement TThostFtdcRatioAttrType = '1' // 结算费率
)

// 投机套保标志类型
//
//go:generate stringer -type TThostFtdcHedgeFlagType -linecomment
type TThostFtdcHedgeFlagType byte

const (
	THOST_FTDC_HF_Speculation TThostFtdcHedgeFlagType = '1' // 投机
	THOST_FTDC_HF_Arbitrage   TThostFtdcHedgeFlagType = '2' // 套利
	THOST_FTDC_HF_Hedge       TThostFtdcHedgeFlagType = '3' // 套保
	THOST_FTDC_HF_MarketMaker TThostFtdcHedgeFlagType = '5' // 做市商
	THOST_FTDC_HF_SpecHedge   TThostFtdcHedgeFlagType = '6' // 第一腿投机第二腿套保
	THOST_FTDC_HF_HedgeSpec   TThostFtdcHedgeFlagType = '7' // 第一腿套保第二腿投机
)

// 投机套保标志类型
//
//go:generate stringer -type TThostFtdcBillHedgeFlagType -linecomment
type TThostFtdcBillHedgeFlagType byte

const (
	THOST_FTDC_BHF_Speculation TThostFtdcBillHedgeFlagType = '1' // 投机
	THOST_FTDC_BHF_Arbitrage   TThostFtdcBillHedgeFlagType = '2' // 套利
	THOST_FTDC_BHF_Hedge       TThostFtdcBillHedgeFlagType = '3' // 套保
)

// 交易编码类型类型
//
//go:generate stringer -type TThostFtdcClientIDTypeType -linecomment
type TThostFtdcClientIDTypeType byte

const (
	THOST_FTDC_CIDT_Speculation TThostFtdcClientIDTypeType = '1' // 投机
	THOST_FTDC_CIDT_Arbitrage   TThostFtdcClientIDTypeType = '2' // 套利
	THOST_FTDC_CIDT_Hedge       TThostFtdcClientIDTypeType = '3' // 套保
	THOST_FTDC_CIDT_MarketMaker TThostFtdcClientIDTypeType = '5' // 做市商
)

// 报单价格条件类型
//
//go:generate stringer -type TThostFtdcOrderPriceTypeType -linecomment
type TThostFtdcOrderPriceTypeType byte

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

// 开平标志类型
//
//go:generate stringer -type TThostFtdcOffsetFlagType -linecomment
type TThostFtdcOffsetFlagType byte

const (
	THOST_FTDC_OF_Open            TThostFtdcOffsetFlagType = '0' // 开仓
	THOST_FTDC_OF_Close           TThostFtdcOffsetFlagType = '1' // 平仓
	THOST_FTDC_OF_ForceClose      TThostFtdcOffsetFlagType = '2' // 强平
	THOST_FTDC_OF_CloseToday      TThostFtdcOffsetFlagType = '3' // 平今
	THOST_FTDC_OF_CloseYesterday  TThostFtdcOffsetFlagType = '4' // 平昨
	THOST_FTDC_OF_ForceOff        TThostFtdcOffsetFlagType = '5' // 强减
	THOST_FTDC_OF_LocalForceClose TThostFtdcOffsetFlagType = '6' // 本地强平
)

// 强平原因类型
//
//go:generate stringer -type TThostFtdcForceCloseReasonType -linecomment
type TThostFtdcForceCloseReasonType byte

const (
	THOST_FTDC_FCC_NotForceClose           TThostFtdcForceCloseReasonType = '0' // 非强平
	THOST_FTDC_FCC_LackDeposit             TThostFtdcForceCloseReasonType = '1' // 资金不足
	THOST_FTDC_FCC_ClientOverPositionLimit TThostFtdcForceCloseReasonType = '2' // 客户超仓
	THOST_FTDC_FCC_MemberOverPositionLimit TThostFtdcForceCloseReasonType = '3' // 会员超仓
	THOST_FTDC_FCC_NotMultiple             TThostFtdcForceCloseReasonType = '4' // 持仓非整数倍
	THOST_FTDC_FCC_Violation               TThostFtdcForceCloseReasonType = '5' // 违规
	THOST_FTDC_FCC_Other                   TThostFtdcForceCloseReasonType = '6' // 其它
	THOST_FTDC_FCC_PersonDeliv             TThostFtdcForceCloseReasonType = '7' // 自然人临近交割
	THOST_FTDC_FCC_Notverifycapital        TThostFtdcForceCloseReasonType = '8' // 本地强平资金不足忽略敞口
	THOST_FTDC_FCC_LocalLackDeposit        TThostFtdcForceCloseReasonType = '9' // 本地强平资金不足
	THOST_FTDC_FCC_LocalViolationNocheck   TThostFtdcForceCloseReasonType = 'a' // 本地强平违规持仓忽略敞口
	THOST_FTDC_FCC_LocalViolation          TThostFtdcForceCloseReasonType = 'b' // 本地强平违规持仓
)

// 报单类型类型
//
//go:generate stringer -type TThostFtdcOrderTypeType -linecomment
type TThostFtdcOrderTypeType byte

const (
	THOST_FTDC_ORDT_Normal                TThostFtdcOrderTypeType = '0' // 正常
	THOST_FTDC_ORDT_DeriveFromQuote       TThostFtdcOrderTypeType = '1' // 报价衍生
	THOST_FTDC_ORDT_DeriveFromCombination TThostFtdcOrderTypeType = '2' // 组合衍生
	THOST_FTDC_ORDT_Combination           TThostFtdcOrderTypeType = '3' // 组合报单
	THOST_FTDC_ORDT_ConditionalOrder      TThostFtdcOrderTypeType = '4' // 条件单
	THOST_FTDC_ORDT_Swap                  TThostFtdcOrderTypeType = '5' // 互换单
	THOST_FTDC_ORDT_DeriveFromBlockTrade  TThostFtdcOrderTypeType = '6' // 大宗交易成交衍生
	THOST_FTDC_ORDT_DeriveFromEFPTrade    TThostFtdcOrderTypeType = '7' // 期转现成交衍生
)

// 有效期类型类型
//
//go:generate stringer -type TThostFtdcTimeConditionType -linecomment
type TThostFtdcTimeConditionType byte

const (
	THOST_FTDC_TC_IOC TThostFtdcTimeConditionType = '1' // 立即完成，否则撤销
	THOST_FTDC_TC_GFS TThostFtdcTimeConditionType = '2' // 本节有效
	THOST_FTDC_TC_GFD TThostFtdcTimeConditionType = '3' // 当日有效
	THOST_FTDC_TC_GTD TThostFtdcTimeConditionType = '4' // 指定日期前有效
	THOST_FTDC_TC_GTC TThostFtdcTimeConditionType = '5' // 撤销前有效
	THOST_FTDC_TC_GFA TThostFtdcTimeConditionType = '6' // 集合竞价有效
)

// 成交量类型类型
//
//go:generate stringer -type TThostFtdcVolumeConditionType -linecomment
type TThostFtdcVolumeConditionType byte

const (
	THOST_FTDC_VC_AV TThostFtdcVolumeConditionType = '1' // 任何数量
	THOST_FTDC_VC_MV TThostFtdcVolumeConditionType = '2' // 最小数量
	THOST_FTDC_VC_CV TThostFtdcVolumeConditionType = '3' // 全部数量
)

// 触发条件类型
//
//go:generate stringer -type TThostFtdcContingentConditionType -linecomment
type TThostFtdcContingentConditionType byte

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

// 操作标志类型
//
//go:generate stringer -type TThostFtdcActionFlagType -linecomment
type TThostFtdcActionFlagType byte

const (
	THOST_FTDC_AF_Delete TThostFtdcActionFlagType = '0' // 删除
	THOST_FTDC_AF_Modify TThostFtdcActionFlagType = '3' // 修改
)

// 交易权限类型
//
//go:generate stringer -type TThostFtdcTradingRightType -linecomment
type TThostFtdcTradingRightType byte

const (
	THOST_FTDC_TR_Allow     TThostFtdcTradingRightType = '0' // 可以交易
	THOST_FTDC_TR_CloseOnly TThostFtdcTradingRightType = '1' // 只能平仓
	THOST_FTDC_TR_Forbidden TThostFtdcTradingRightType = '2' // 不能交易
)

// 报单来源类型
//
//go:generate stringer -type TThostFtdcOrderSourceType -linecomment
type TThostFtdcOrderSourceType byte

const (
	THOST_FTDC_OSRC_Participant   TThostFtdcOrderSourceType = '0' // 来自参与者
	THOST_FTDC_OSRC_Administrator TThostFtdcOrderSourceType = '1' // 来自管理员
)

// 成交类型类型
//
//go:generate stringer -type TThostFtdcTradeTypeType -linecomment
type TThostFtdcTradeTypeType byte

const (
	THOST_FTDC_TRDT_SplitCombination   TThostFtdcTradeTypeType = '#' // 组合持仓拆分为单一持仓,初始化不应包含该类型的持仓
	THOST_FTDC_TRDT_Common             TThostFtdcTradeTypeType = '0' // 普通成交
	THOST_FTDC_TRDT_OptionsExecution   TThostFtdcTradeTypeType = '1' // 期权执行
	THOST_FTDC_TRDT_OTC                TThostFtdcTradeTypeType = '2' // OTC成交
	THOST_FTDC_TRDT_EFPDerived         TThostFtdcTradeTypeType = '3' // 期转现衍生成交
	THOST_FTDC_TRDT_CombinationDerived TThostFtdcTradeTypeType = '4' // 组合衍生成交
	THOST_FTDC_TRDT_BlockTrade         TThostFtdcTradeTypeType = '5' // 大宗交易成交
)

// 特殊持仓明细标识类型
//
//go:generate stringer -type TThostFtdcSpecPosiTypeType -linecomment
type TThostFtdcSpecPosiTypeType byte

const (
	THOST_FTDC_SPOST_Common TThostFtdcSpecPosiTypeType = '#' // 普通持仓明细
	THOST_FTDC_SPOST_Tas    TThostFtdcSpecPosiTypeType = '0' // TAS合约成交产生的标的合约持仓明细
)

// 成交价来源类型
//
//go:generate stringer -type TThostFtdcPriceSourceType -linecomment
type TThostFtdcPriceSourceType byte

const (
	THOST_FTDC_PSRC_LastPrice TThostFtdcPriceSourceType = '0' // 前成交价
	THOST_FTDC_PSRC_Buy       TThostFtdcPriceSourceType = '1' // 买委托价
	THOST_FTDC_PSRC_Sell      TThostFtdcPriceSourceType = '2' // 卖委托价
	THOST_FTDC_PSRC_OTC       TThostFtdcPriceSourceType = '3' // 场外成交价
)

// 合约交易状态类型
//
//go:generate stringer -type TThostFtdcInstrumentStatusType -linecomment
type TThostFtdcInstrumentStatusType byte

const (
	THOST_FTDC_IS_BeforeTrading         TThostFtdcInstrumentStatusType = '0' // 开盘前
	THOST_FTDC_IS_NoTrading             TThostFtdcInstrumentStatusType = '1' // 非交易
	THOST_FTDC_IS_Continous             TThostFtdcInstrumentStatusType = '2' // 连续交易
	THOST_FTDC_IS_AuctionOrdering       TThostFtdcInstrumentStatusType = '3' // 集合竞价报单
	THOST_FTDC_IS_AuctionBalance        TThostFtdcInstrumentStatusType = '4' // 集合竞价价格平衡
	THOST_FTDC_IS_AuctionMatch          TThostFtdcInstrumentStatusType = '5' // 集合竞价撮合
	THOST_FTDC_IS_Closed                TThostFtdcInstrumentStatusType = '6' // 收盘
	THOST_FTDC_IS_TransactionProcessing TThostFtdcInstrumentStatusType = '7' // 交易业务处理
)

// 品种进入交易状态原因类型
//
//go:generate stringer -type TThostFtdcInstStatusEnterReasonType -linecomment
type TThostFtdcInstStatusEnterReasonType byte

const (
	THOST_FTDC_IER_Automatic TThostFtdcInstStatusEnterReasonType = '1' // 自动切换
	THOST_FTDC_IER_Manual    TThostFtdcInstStatusEnterReasonType = '2' // 手动切换
	THOST_FTDC_IER_Fuse      TThostFtdcInstStatusEnterReasonType = '3' // 熔断
)

// 报单操作引用类型
type TThostFtdcOrderActionRefType int32

// 安装数量类型
type TThostFtdcInstallCountType int32

// 安装编号类型
type TThostFtdcInstallIDType int32

// 错误代码类型
type TThostFtdcErrorIDType int32

// 结算编号类型
type TThostFtdcSettlementIDType int32

// 数量类型
type TThostFtdcVolumeType int32

// 前置编号类型
type TThostFtdcFrontIDType int32

// 会话编号类型
type TThostFtdcSessionIDType int32

// 序号类型
type TThostFtdcSequenceNoType int32

// DB命令序号类型
type TThostFtdcCommandNoType int32

// 时间（毫秒）类型
type TThostFtdcMillisecType int32

// 时间（秒）类型
type TThostFtdcSecType int32

// 合约数量乘数类型
type TThostFtdcVolumeMultipleType int32

// 交易阶段编号类型
type TThostFtdcTradingSegmentSNType int32

// 请求编号类型
type TThostFtdcRequestIDType int32

// 年份类型
type TThostFtdcYearType int32

// 月份类型
type TThostFtdcMonthType int32

// 布尔型类型
type TThostFtdcBoolType int32

// 价格类型
type TThostFtdcPriceType Double

// 组合开平标志类型
type TThostFtdcCombOffsetFlagType Byte5

// 组合投机套保标志类型
type TThostFtdcCombHedgeFlagType Byte5

// 比率类型
type TThostFtdcRatioType Double

// 资金类型
type TThostFtdcMoneyType Double

// 大额数量类型
type TThostFtdcLargeVolumeType Double

// 序列系列号类型
type TThostFtdcSequenceSeriesType int16

// 通讯时段编号类型
type TThostFtdcCommPhaseNoType int16

// 序列编号类型
type TThostFtdcSequenceLabelType Byte2

// 基础商品乘数类型
type TThostFtdcUnderlyingMultipleType Double

// 优先级类型
type TThostFtdcPriorityType int32

// 合同编号类型
type TThostFtdcContractCodeType Byte41

// 市类型
type TThostFtdcCityType Byte51

// 是否股民类型
type TThostFtdcIsStockType Byte11

// 渠道类型
type TThostFtdcChannelType Byte51

// 通讯地址类型
type TThostFtdcAddressType Byte101

// 邮政编码类型
type TThostFtdcZipCodeType Byte7

// 联系电话类型
type TThostFtdcTelephoneType Byte41

// 传真类型
type TThostFtdcFaxType Byte41

// 手机类型
type TThostFtdcMobileType Byte41

// 电子邮件类型
type TThostFtdcEMailType Byte41

// 备注类型
type TThostFtdcMemoType Byte161

// 企业代码类型
type TThostFtdcCompanyCodeType Byte51

// 网站地址类型
type TThostFtdcWebsiteType Byte51

// 税务登记号类型
type TThostFtdcTaxNoType Byte31

// 处理状态类型
//
//go:generate stringer -type TThostFtdcBatchStatusType -linecomment
type TThostFtdcBatchStatusType byte

const (
	THOST_FTDC_BS_NoUpload TThostFtdcBatchStatusType = '1' // 未上传
	THOST_FTDC_BS_Uploaded TThostFtdcBatchStatusType = '2' // 已上传
	THOST_FTDC_BS_Failed   TThostFtdcBatchStatusType = '3' // 审核失败
)

// 属性代码类型
type TThostFtdcPropertyIDType Byte33

// 属性名称类型
type TThostFtdcPropertyNameType Byte65

// 营业执照号类型
type TThostFtdcLicenseNoType Byte51

// 经纪人代码类型
type TThostFtdcAgentIDType Byte13

// 经纪人名称类型
type TThostFtdcAgentNameType Byte41

// 经纪人组代码类型
type TThostFtdcAgentGroupIDType Byte13

// 经纪人组名称类型
type TThostFtdcAgentGroupNameType Byte41

// 按品种返还方式类型
//
//go:generate stringer -type TThostFtdcReturnStyleType -linecomment
type TThostFtdcReturnStyleType byte

const (
	THOST_FTDC_RS_All       TThostFtdcReturnStyleType = '1' // 按所有品种
	THOST_FTDC_RS_ByProduct TThostFtdcReturnStyleType = '2' // 按品种
)

// 返还模式类型
//
//go:generate stringer -type TThostFtdcReturnPatternType -linecomment
type TThostFtdcReturnPatternType byte

const (
	THOST_FTDC_RP_ByVolume    TThostFtdcReturnPatternType = '1' // 按成交手数
	THOST_FTDC_RP_ByFeeOnHand TThostFtdcReturnPatternType = '2' // 按留存手续费
)

// 返还级别类型
//
//go:generate stringer -type TThostFtdcReturnLevelType -linecomment
type TThostFtdcReturnLevelType byte

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

// 返还标准类型
//
//go:generate stringer -type TThostFtdcReturnStandardType -linecomment
type TThostFtdcReturnStandardType byte

const (
	THOST_FTDC_RSD_ByPeriod   TThostFtdcReturnStandardType = '1' // 分阶段返还
	THOST_FTDC_RSD_ByStandard TThostFtdcReturnStandardType = '2' // 按某一标准
)

// 质押类型类型
//
//go:generate stringer -type TThostFtdcMortgageTypeType -linecomment
type TThostFtdcMortgageTypeType byte

const (
	THOST_FTDC_MT_Out TThostFtdcMortgageTypeType = '0' // 质出
	THOST_FTDC_MT_In  TThostFtdcMortgageTypeType = '1' // 质入
)

// 投资者结算参数代码类型
//
//go:generate stringer -type TThostFtdcInvestorSettlementParamIDType -linecomment
type TThostFtdcInvestorSettlementParamIDType byte

const (
	THOST_FTDC_ISPI_MortgageRatio TThostFtdcInvestorSettlementParamIDType = '4' // 质押比例
	THOST_FTDC_ISPI_MarginWay     TThostFtdcInvestorSettlementParamIDType = '5' // 保证金算法
	THOST_FTDC_ISPI_BillDeposit   TThostFtdcInvestorSettlementParamIDType = '9' // 结算单结存是否包含质押
)

// 交易所结算参数代码类型
//
//go:generate stringer -type TThostFtdcExchangeSettlementParamIDType -linecomment
type TThostFtdcExchangeSettlementParamIDType byte

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

// 系统参数代码类型
//
//go:generate stringer -type TThostFtdcSystemParamIDType -linecomment
type TThostFtdcSystemParamIDType byte

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

// 交易系统参数代码类型
//
//go:generate stringer -type TThostFtdcTradeParamIDType -linecomment
type TThostFtdcTradeParamIDType byte

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
	THOST_FTDC_TPID_IsCheckBankAcc          TThostFtdcTradeParamIDType = 'I' // 银期开户是否验证开户银行卡号是否是预留银行账户
	THOST_FTDC_TPID_PasswordDeadLine        TThostFtdcTradeParamIDType = 'J' // 弱密码最后修改日期
	THOST_FTDC_TPID_IsStrongPassword        TThostFtdcTradeParamIDType = 'K' // 强密码校验
	THOST_FTDC_TPID_BalanceMorgage          TThostFtdcTradeParamIDType = 'a' // 自有资金质押比
	THOST_FTDC_TPID_MinPwdLen               TThostFtdcTradeParamIDType = 'O' // 最小密码长度
	THOST_FTDC_TPID_LoginFailMaxNumForIP    TThostFtdcTradeParamIDType = 'U' // IP当日最大登陆失败次数
	THOST_FTDC_TPID_PasswordPeriod          TThostFtdcTradeParamIDType = 'V' // 密码有效期
	THOST_FTDC_TPID_PwdHistoryCmp           TThostFtdcTradeParamIDType = 'X' // 历史密码重复限制次数
	THOST_FTDC_TPID_TranferChkProperty      TThostFtdcTradeParamIDType = 'i' // 转账是否验证预留银行账户
	THOST_FTDC_TPID_TradeChkPhase           TThostFtdcTradeParamIDType = 'j' // 非交易时间异常报单校验参数
	THOST_FTDC_TPID_TradeChkPriceVol        TThostFtdcTradeParamIDType = 'k' // 其他异常报单校验参数（价格和手数）
	THOST_FTDC_TPID_NewBESMarginAlgo        TThostFtdcTradeParamIDType = 'l' // 卖出垂直价差组合新算法
)

// 参数代码值类型
type TThostFtdcSettlementParamValueType Byte256

// 计数器代码类型
type TThostFtdcCounterIDType Byte33

// 投资者分组名称类型
type TThostFtdcInvestorGroupNameType Byte41

// 牌号类型
type TThostFtdcBrandCodeType Byte257

// 仓库类型
type TThostFtdcWarehouseType Byte257

// 产期类型
type TThostFtdcProductDateType Byte41

// 等级类型
type TThostFtdcGradeType Byte41

// 类别类型
type TThostFtdcClassifyType Byte41

// 货位类型
type TThostFtdcPositionType Byte41

// 产地类型
type TThostFtdcYieldlyType Byte41

// 公定重量类型
type TThostFtdcWeightType Byte41

// 分项资金流水号类型
type TThostFtdcSubEntryFundNoType int32

// 文件标识类型
//
//go:generate stringer -type TThostFtdcFileIDType -linecomment
type TThostFtdcFileIDType byte

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

// 文件名称类型
type TThostFtdcFileNameType Byte257

// 文件上传类型类型
//
//go:generate stringer -type TThostFtdcFileTypeType -linecomment
type TThostFtdcFileTypeType byte

const (
	THOST_FTDC_FUT_Settlement TThostFtdcFileTypeType = '0' // 结算
	THOST_FTDC_FUT_Check      TThostFtdcFileTypeType = '1' // 核对
)

// 文件格式类型
//
//go:generate stringer -type TThostFtdcFileFormatType -linecomment
type TThostFtdcFileFormatType byte

const (
	THOST_FTDC_FFT_Txt TThostFtdcFileFormatType = '0' // 文本文件(.txt)
	THOST_FTDC_FFT_Zip TThostFtdcFileFormatType = '1' // 压缩文件(.zip)
	THOST_FTDC_FFT_DBF TThostFtdcFileFormatType = '2' // DBF文件(.dbf)
)

// 文件状态类型
//
//go:generate stringer -type TThostFtdcFileUploadStatusType -linecomment
type TThostFtdcFileUploadStatusType byte

const (
	THOST_FTDC_FUS_SucceedUpload   TThostFtdcFileUploadStatusType = '1' // 上传成功
	THOST_FTDC_FUS_FailedUpload    TThostFtdcFileUploadStatusType = '2' // 上传失败
	THOST_FTDC_FUS_SucceedLoad     TThostFtdcFileUploadStatusType = '3' // 导入成功
	THOST_FTDC_FUS_PartSucceedLoad TThostFtdcFileUploadStatusType = '4' // 导入部分成功
	THOST_FTDC_FUS_FailedLoad      TThostFtdcFileUploadStatusType = '5' // 导入失败
)

// 移仓方向类型
//
//go:generate stringer -type TThostFtdcTransferDirectionType -linecomment
type TThostFtdcTransferDirectionType byte

const (
	THOST_FTDC_TD_Out TThostFtdcTransferDirectionType = '0' // 移出
	THOST_FTDC_TD_In  TThostFtdcTransferDirectionType = '1' // 移入
)

// 上传文件类型类型
type TThostFtdcUploadModeType Byte21

// 投资者帐号类型
type TThostFtdcAccountIDType Byte13

// 银行统一标识类型类型
type TThostFtdcBankFlagType Byte4

// 银行账户类型
type TThostFtdcBankAccountType Byte41

// 银行账户的开户人名称类型
type TThostFtdcOpenNameType Byte61

// 银行账户的开户行类型
type TThostFtdcOpenBankType Byte101

// 银行名称类型
type TThostFtdcBankNameType Byte101

// 发布路径类型
type TThostFtdcPublishPathType Byte257

// 操作员代码类型
type TThostFtdcOperatorIDType Byte65

// 月份数量类型
type TThostFtdcMonthCountType int32

// 月份提前数组类型
type TThostFtdcAdvanceMonthArrayType Byte13

// 日期表达式类型
type TThostFtdcDateExprType Byte1025

// 合约代码表达式类型
type TThostFtdcInstrumentIDExprType Byte41

// 合约名称表达式类型
type TThostFtdcInstrumentNameExprType Byte41

// 特殊的创建规则类型
//
//go:generate stringer -type TThostFtdcSpecialCreateRuleType -linecomment
type TThostFtdcSpecialCreateRuleType byte

const (
	THOST_FTDC_SC_NoSpecialRule    TThostFtdcSpecialCreateRuleType = '0' // 没有特殊创建规则
	THOST_FTDC_SC_NoSpringFestival TThostFtdcSpecialCreateRuleType = '1' // 不包含春节
)

// 挂牌基准价类型类型
//
//go:generate stringer -type TThostFtdcBasisPriceTypeType -linecomment
type TThostFtdcBasisPriceTypeType byte

const (
	THOST_FTDC_IPT_LastSettlement TThostFtdcBasisPriceTypeType = '1' // 上一合约结算价
	THOST_FTDC_IPT_LaseClose      TThostFtdcBasisPriceTypeType = '2' // 上一合约收盘价
)

// 产品生命周期状态类型
//
//go:generate stringer -type TThostFtdcProductLifePhaseType -linecomment
type TThostFtdcProductLifePhaseType byte

const (
	THOST_FTDC_PLP_Active    TThostFtdcProductLifePhaseType = '1' // 活跃
	THOST_FTDC_PLP_NonActive TThostFtdcProductLifePhaseType = '2' // 不活跃
	THOST_FTDC_PLP_Canceled  TThostFtdcProductLifePhaseType = '3' // 注销
)

// 交割方式类型
//
//go:generate stringer -type TThostFtdcDeliveryModeType -linecomment
type TThostFtdcDeliveryModeType byte

const (
	THOST_FTDC_DM_CashDeliv      TThostFtdcDeliveryModeType = '1' // 现金交割
	THOST_FTDC_DM_CommodityDeliv TThostFtdcDeliveryModeType = '2' // 实物交割
)

// 日志级别类型
type TThostFtdcLogLevelType Byte33

// 存储过程名称类型
type TThostFtdcProcessNameType Byte257

// 操作摘要类型
type TThostFtdcOperationMemoType Byte1025

// 出入金类型类型
//
//go:generate stringer -type TThostFtdcFundIOTypeType -linecomment
type TThostFtdcFundIOTypeType byte

const (
	THOST_FTDC_FIOT_FundIO       TThostFtdcFundIOTypeType = '1' // 出入金
	THOST_FTDC_FIOT_Transfer     TThostFtdcFundIOTypeType = '2' // 银期转帐
	THOST_FTDC_FIOT_SwapCurrency TThostFtdcFundIOTypeType = '3' // 银期换汇
)

// 资金类型类型
//
//go:generate stringer -type TThostFtdcFundTypeType -linecomment
type TThostFtdcFundTypeType byte

const (
	THOST_FTDC_FT_Deposite      TThostFtdcFundTypeType = '1' // 银行存款
	THOST_FTDC_FT_ItemFund      TThostFtdcFundTypeType = '2' // 分项资金
	THOST_FTDC_FT_Company       TThostFtdcFundTypeType = '3' // 公司调整
	THOST_FTDC_FT_InnerTransfer TThostFtdcFundTypeType = '4' // 资金内转
)

// 出入金方向类型
//
//go:generate stringer -type TThostFtdcFundDirectionType -linecomment
type TThostFtdcFundDirectionType byte

const (
	THOST_FTDC_FD_In  TThostFtdcFundDirectionType = '1' // 入金
	THOST_FTDC_FD_Out TThostFtdcFundDirectionType = '2' // 出金
)

// 资金状态类型
//
//go:generate stringer -type TThostFtdcFundStatusType -linecomment
type TThostFtdcFundStatusType byte

const (
	THOST_FTDC_FS_Record TThostFtdcFundStatusType = '1' // 已录入
	THOST_FTDC_FS_Check  TThostFtdcFundStatusType = '2' // 已复核
	THOST_FTDC_FS_Charge TThostFtdcFundStatusType = '3' // 已冲销
)

// 票据号类型
type TThostFtdcBillNoType Byte15

// 票据名称类型
type TThostFtdcBillNameType Byte33

// 发布状态类型
//
//go:generate stringer -type TThostFtdcPublishStatusType -linecomment
type TThostFtdcPublishStatusType byte

const (
	THOST_FTDC_PS_None       TThostFtdcPublishStatusType = '1' // 未发布
	THOST_FTDC_PS_Publishing TThostFtdcPublishStatusType = '2' // 正在发布
	THOST_FTDC_PS_Published  TThostFtdcPublishStatusType = '3' // 已发布
)

// 枚举值代码类型
type TThostFtdcEnumValueIDType Byte65

// 枚举值类型类型
type TThostFtdcEnumValueTypeType Byte33

// 枚举值名称类型
type TThostFtdcEnumValueLabelType Byte65

// 枚举值结果类型
type TThostFtdcEnumValueResultType Byte33

// 系统状态类型
//
//go:generate stringer -type TThostFtdcSystemStatusType -linecomment
type TThostFtdcSystemStatusType byte

const (
	THOST_FTDC_ES_NonActive   TThostFtdcSystemStatusType = '1' // 不活跃
	THOST_FTDC_ES_Startup     TThostFtdcSystemStatusType = '2' // 启动
	THOST_FTDC_ES_Initialize  TThostFtdcSystemStatusType = '3' // 交易开始初始化
	THOST_FTDC_ES_Initialized TThostFtdcSystemStatusType = '4' // 交易完成初始化
	THOST_FTDC_ES_Close       TThostFtdcSystemStatusType = '5' // 收市开始
	THOST_FTDC_ES_Closed      TThostFtdcSystemStatusType = '6' // 收市完成
	THOST_FTDC_ES_Settlement  TThostFtdcSystemStatusType = '7' // 结算
)

// 结算状态类型
//
//go:generate stringer -type TThostFtdcSettlementStatusType -linecomment
type TThostFtdcSettlementStatusType byte

const (
	THOST_FTDC_STS_Initialize    TThostFtdcSettlementStatusType = '0' // 初始
	THOST_FTDC_STS_Settlementing TThostFtdcSettlementStatusType = '1' // 结算中
	THOST_FTDC_STS_Settlemented  TThostFtdcSettlementStatusType = '2' // 已结算
	THOST_FTDC_STS_Finished      TThostFtdcSettlementStatusType = '3' // 结算完成
)

// 限定值类型类型
type TThostFtdcRangeIntTypeType Byte33

// 限定值下限类型
type TThostFtdcRangeIntFromType Byte33

// 限定值上限类型
type TThostFtdcRangeIntToType Byte33

// 功能代码类型
type TThostFtdcFunctionIDType Byte25

// 功能编码类型
type TThostFtdcFunctionValueCodeType Byte257

// 功能名称类型
type TThostFtdcFunctionNameType Byte65

// 角色编号类型
type TThostFtdcRoleIDType Byte11

// 角色名称类型
type TThostFtdcRoleNameType Byte41

// 描述类型
type TThostFtdcDescriptionType Byte401

// 组合编号类型
type TThostFtdcCombineIDType Byte25

// 组合类型类型
type TThostFtdcCombineTypeType Byte25

// 投资者类型类型
//
//go:generate stringer -type TThostFtdcInvestorTypeType -linecomment
type TThostFtdcInvestorTypeType byte

const (
	THOST_FTDC_CT_Person       TThostFtdcInvestorTypeType = '0' // 自然人
	THOST_FTDC_CT_Company      TThostFtdcInvestorTypeType = '1' // 法人
	THOST_FTDC_CT_Fund         TThostFtdcInvestorTypeType = '2' // 投资基金
	THOST_FTDC_CT_SpecialOrgan TThostFtdcInvestorTypeType = '3' // 特殊法人
	THOST_FTDC_CT_Asset        TThostFtdcInvestorTypeType = '4' // 资管户
)

// 经纪公司类型类型
//
//go:generate stringer -type TThostFtdcBrokerTypeType -linecomment
type TThostFtdcBrokerTypeType byte

const (
	THOST_FTDC_BT_Trade       TThostFtdcBrokerTypeType = '0' // 交易会员
	THOST_FTDC_BT_TradeSettle TThostFtdcBrokerTypeType = '1' // 交易结算会员
)

// 风险等级类型
//
//go:generate stringer -type TThostFtdcRiskLevelType -linecomment
type TThostFtdcRiskLevelType byte

const (
	THOST_FTDC_FAS_Low    TThostFtdcRiskLevelType = '1' // 低风险客户
	THOST_FTDC_FAS_Normal TThostFtdcRiskLevelType = '2' // 普通客户
	THOST_FTDC_FAS_Focus  TThostFtdcRiskLevelType = '3' // 关注客户
	THOST_FTDC_FAS_Risk   TThostFtdcRiskLevelType = '4' // 风险客户
)

// 手续费收取方式类型
//
//go:generate stringer -type TThostFtdcFeeAcceptStyleType -linecomment
type TThostFtdcFeeAcceptStyleType byte

const (
	THOST_FTDC_FAS_ByTrade TThostFtdcFeeAcceptStyleType = '1' // 按交易收取
	THOST_FTDC_FAS_ByDeliv TThostFtdcFeeAcceptStyleType = '2' // 按交割收取
	THOST_FTDC_FAS_None    TThostFtdcFeeAcceptStyleType = '3' // 不收
	THOST_FTDC_FAS_FixFee  TThostFtdcFeeAcceptStyleType = '4' // 按指定手续费收取
)

// 密码类型类型
//
//go:generate stringer -type TThostFtdcPasswordTypeType -linecomment
type TThostFtdcPasswordTypeType byte

const (
	THOST_FTDC_PWDT_Trade   TThostFtdcPasswordTypeType = '1' // 交易密码
	THOST_FTDC_PWDT_Account TThostFtdcPasswordTypeType = '2' // 资金密码
)

// 盈亏算法类型
//
//go:generate stringer -type TThostFtdcAlgorithmType -linecomment
type TThostFtdcAlgorithmType byte

const (
	THOST_FTDC_AG_All      TThostFtdcAlgorithmType = '1' // 浮盈浮亏都计算
	THOST_FTDC_AG_OnlyLost TThostFtdcAlgorithmType = '2' // 浮盈不计，浮亏计
	THOST_FTDC_AG_OnlyGain TThostFtdcAlgorithmType = '3' // 浮盈计，浮亏不计
	THOST_FTDC_AG_None     TThostFtdcAlgorithmType = '4' // 浮盈浮亏都不计算
)

// 是否包含平仓盈利类型
//
//go:generate stringer -type TThostFtdcIncludeCloseProfitType -linecomment
type TThostFtdcIncludeCloseProfitType byte

const (
	THOST_FTDC_ICP_Include    TThostFtdcIncludeCloseProfitType = '0' // 包含平仓盈利
	THOST_FTDC_ICP_NotInclude TThostFtdcIncludeCloseProfitType = '2' // 不包含平仓盈利
)

// 是否受可提比例限制类型
//
//go:generate stringer -type TThostFtdcAllWithoutTradeType -linecomment
type TThostFtdcAllWithoutTradeType byte

const (
	THOST_FTDC_AWT_Enable       TThostFtdcAllWithoutTradeType = '0' // 无仓无成交不受可提比例限制
	THOST_FTDC_AWT_Disable      TThostFtdcAllWithoutTradeType = '2' // 受可提比例限制
	THOST_FTDC_AWT_NoHoldEnable TThostFtdcAllWithoutTradeType = '3' // 无仓不受可提比例限制
)

// 盈亏算法说明类型
type TThostFtdcCommentType Byte31

// 版本号类型
type TThostFtdcVersionType Byte4

// 交易代码类型
type TThostFtdcTradeCodeType Byte7

// 交易日期类型
type TThostFtdcTradeDateType Byte9

// 交易时间类型
type TThostFtdcTradeTimeType Byte9

// 发起方流水号类型
type TThostFtdcTradeSerialType Byte9

// 发起方流水号类型
type TThostFtdcTradeSerialNoType int32

// 期货公司代码类型
type TThostFtdcFutureIDType Byte11

// 银行代码类型
type TThostFtdcBankIDType Byte4

// 银行分中心代码类型
type TThostFtdcBankBrchIDType Byte5

// 分中心代码类型
type TThostFtdcBankBranchIDType Byte11

// 交易柜员类型
type TThostFtdcOperNoType Byte17

// 渠道标志类型
type TThostFtdcDeviceIDType Byte3

// 记录数类型
type TThostFtdcRecordNumType Byte7

// 期货资金账号类型
type TThostFtdcFutureAccountType Byte22

// 资金密码核对标志类型
//
//go:generate stringer -type TThostFtdcFuturePwdFlagType -linecomment
type TThostFtdcFuturePwdFlagType byte

const (
	THOST_FTDC_FPWD_UnCheck TThostFtdcFuturePwdFlagType = '0' // 不核对
	THOST_FTDC_FPWD_Check   TThostFtdcFuturePwdFlagType = '1' // 核对
)

// 银期转账类型类型
//
//go:generate stringer -type TThostFtdcTransferTypeType -linecomment
type TThostFtdcTransferTypeType byte

const (
	THOST_FTDC_TT_BankToFuture TThostFtdcTransferTypeType = '0' // 银行转期货
	THOST_FTDC_TT_FutureToBank TThostFtdcTransferTypeType = '1' // 期货转银行
)

// 期货资金密码类型
type TThostFtdcFutureAccPwdType Byte17

// 币种类型
type TThostFtdcCurrencyCodeType Byte4

// 响应代码类型
type TThostFtdcRetCodeType Byte5

// 响应信息类型
type TThostFtdcRetInfoType Byte129

// 银行总余额类型
type TThostFtdcTradeAmtType Byte20

// 银行可用余额类型
type TThostFtdcUseAmtType Byte20

// 银行可取余额类型
type TThostFtdcFetchAmtType Byte20

// 转账有效标志类型
//
//go:generate stringer -type TThostFtdcTransferValidFlagType -linecomment
type TThostFtdcTransferValidFlagType byte

const (
	THOST_FTDC_TVF_Invalid TThostFtdcTransferValidFlagType = '0' // 无效或失败
	THOST_FTDC_TVF_Valid   TThostFtdcTransferValidFlagType = '1' // 有效
	THOST_FTDC_TVF_Reverse TThostFtdcTransferValidFlagType = '2' // 冲正
)

// 证件号码类型
type TThostFtdcCertCodeType Byte21

// 事由类型
//
//go:generate stringer -type TThostFtdcReasonType -linecomment
type TThostFtdcReasonType byte

const (
	THOST_FTDC_RN_CD TThostFtdcReasonType = '0' // 错单
	THOST_FTDC_RN_ZT TThostFtdcReasonType = '1' // 资金在途
	THOST_FTDC_RN_QT TThostFtdcReasonType = '2' // 其它
)

// 资金项目编号类型
type TThostFtdcFundProjectIDType Byte5

// 性别类型
//
//go:generate stringer -type TThostFtdcSexType -linecomment
type TThostFtdcSexType byte

const (
	THOST_FTDC_SEX_None  TThostFtdcSexType = '0' // 未知
	THOST_FTDC_SEX_Man   TThostFtdcSexType = '1' // 男
	THOST_FTDC_SEX_Woman TThostFtdcSexType = '2' // 女
)

// 职业类型
type TThostFtdcProfessionType Byte101

// 国籍类型
type TThostFtdcNationalType Byte31

// 省类型
type TThostFtdcProvinceType Byte51

// 区类型
type TThostFtdcRegionType Byte16

// 国家类型
type TThostFtdcCountryType Byte16

// 营业执照类型
type TThostFtdcLicenseNOType Byte33

// 企业性质类型
type TThostFtdcCompanyTypeType Byte16

// 经营范围类型
type TThostFtdcBusinessScopeType Byte1001

// 注册资本币种类型
type TThostFtdcCapitalCurrencyType Byte4

// 用户类型类型
//
//go:generate stringer -type TThostFtdcUserTypeType -linecomment
type TThostFtdcUserTypeType byte

const (
	THOST_FTDC_UT_Investor  TThostFtdcUserTypeType = '0' // 投资者
	THOST_FTDC_UT_Operator  TThostFtdcUserTypeType = '1' // 操作员
	THOST_FTDC_UT_SuperUser TThostFtdcUserTypeType = '2' // 管理员
)

// 营业部编号类型
type TThostFtdcBranchIDType Byte9

// 费率类型类型
//
//go:generate stringer -type TThostFtdcRateTypeType -linecomment
type TThostFtdcRateTypeType byte

const THOST_FTDC_RATETYPE_MarginRate TThostFtdcRateTypeType = '2' // 保证金率

// 通知类型类型
//
//go:generate stringer -type TThostFtdcNoteTypeType -linecomment
type TThostFtdcNoteTypeType byte

const (
	THOST_FTDC_NOTETYPE_TradeSettleBill  TThostFtdcNoteTypeType = '1' // 交易结算单
	THOST_FTDC_NOTETYPE_TradeSettleMonth TThostFtdcNoteTypeType = '2' // 交易结算月报
	THOST_FTDC_NOTETYPE_CallMarginNotes  TThostFtdcNoteTypeType = '3' // 追加保证金通知书
	THOST_FTDC_NOTETYPE_ForceCloseNotes  TThostFtdcNoteTypeType = '4' // 强行平仓通知书
	THOST_FTDC_NOTETYPE_TradeNotes       TThostFtdcNoteTypeType = '5' // 成交通知书
	THOST_FTDC_NOTETYPE_DelivNotes       TThostFtdcNoteTypeType = '6' // 交割通知书
)

// 结算单方式类型
//
//go:generate stringer -type TThostFtdcSettlementStyleType -linecomment
type TThostFtdcSettlementStyleType byte

const (
	THOST_FTDC_SBS_Day    TThostFtdcSettlementStyleType = '1' // 逐日盯市
	THOST_FTDC_SBS_Volume TThostFtdcSettlementStyleType = '2' // 逐笔对冲
)

// 域名类型
type TThostFtdcBrokerDNSType Byte256

// 语句类型
type TThostFtdcSentenceType Byte501

// 结算单类型类型
//
//go:generate stringer -type TThostFtdcSettlementBillTypeType -linecomment
type TThostFtdcSettlementBillTypeType byte

const (
	THOST_FTDC_ST_Day   TThostFtdcSettlementBillTypeType = '0' // 日报
	THOST_FTDC_ST_Month TThostFtdcSettlementBillTypeType = '1' // 月报
)

// 客户权限类型类型
//
//go:generate stringer -type TThostFtdcUserRightTypeType -linecomment
type TThostFtdcUserRightTypeType byte

const (
	THOST_FTDC_URT_Logon          TThostFtdcUserRightTypeType = '1' // 登录
	THOST_FTDC_URT_Transfer       TThostFtdcUserRightTypeType = '2' // 银期转帐
	THOST_FTDC_URT_EMail          TThostFtdcUserRightTypeType = '3' // 邮寄结算单
	THOST_FTDC_URT_Fax            TThostFtdcUserRightTypeType = '4' // 传真结算单
	THOST_FTDC_URT_ConditionOrder TThostFtdcUserRightTypeType = '5' // 条件单
)

// 保证金价格类型类型
//
//go:generate stringer -type TThostFtdcMarginPriceTypeType -linecomment
type TThostFtdcMarginPriceTypeType byte

const (
	THOST_FTDC_MPT_PreSettlementPrice TThostFtdcMarginPriceTypeType = '1' // 昨结算价
	THOST_FTDC_MPT_SettlementPrice    TThostFtdcMarginPriceTypeType = '2' // 最新价
	THOST_FTDC_MPT_AveragePrice       TThostFtdcMarginPriceTypeType = '3' // 成交均价
	THOST_FTDC_MPT_OpenPrice          TThostFtdcMarginPriceTypeType = '4' // 开仓价
)

// 结算单生成状态类型
//
//go:generate stringer -type TThostFtdcBillGenStatusType -linecomment
type TThostFtdcBillGenStatusType byte

const (
	THOST_FTDC_BGS_None        TThostFtdcBillGenStatusType = '0' // 未生成
	THOST_FTDC_BGS_NoGenerated TThostFtdcBillGenStatusType = '1' // 生成中
	THOST_FTDC_BGS_Generated   TThostFtdcBillGenStatusType = '2' // 已生成
)

// 算法类型类型
//
//go:generate stringer -type TThostFtdcAlgoTypeType -linecomment
type TThostFtdcAlgoTypeType byte

const (
	THOST_FTDC_AT_HandlePositionAlgo TThostFtdcAlgoTypeType = '1' // 持仓处理算法
	THOST_FTDC_AT_FindMarginRateAlgo TThostFtdcAlgoTypeType = '2' // 寻找保证金率算法
)

// 持仓处理算法编号类型
//
//go:generate stringer -type TThostFtdcHandlePositionAlgoIDType -linecomment
type TThostFtdcHandlePositionAlgoIDType byte

const (
	THOST_FTDC_HPA_Base TThostFtdcHandlePositionAlgoIDType = '1' // 基本
	THOST_FTDC_HPA_DCE  TThostFtdcHandlePositionAlgoIDType = '2' // 大连商品交易所
	THOST_FTDC_HPA_CZCE TThostFtdcHandlePositionAlgoIDType = '3' // 郑州商品交易所
)

// 寻找保证金率算法编号类型
//
//go:generate stringer -type TThostFtdcFindMarginRateAlgoIDType -linecomment
type TThostFtdcFindMarginRateAlgoIDType byte

const (
	THOST_FTDC_FMRA_Base TThostFtdcFindMarginRateAlgoIDType = '1' // 基本
	THOST_FTDC_FMRA_DCE  TThostFtdcFindMarginRateAlgoIDType = '2' // 大连商品交易所
	THOST_FTDC_FMRA_CZCE TThostFtdcFindMarginRateAlgoIDType = '3' // 郑州商品交易所
)

// 资金处理算法编号类型
//
//go:generate stringer -type TThostFtdcHandleTradingAccountAlgoIDType -linecomment
type TThostFtdcHandleTradingAccountAlgoIDType byte

const (
	THOST_FTDC_HTAA_Base TThostFtdcHandleTradingAccountAlgoIDType = '1' // 基本
	THOST_FTDC_HTAA_DCE  TThostFtdcHandleTradingAccountAlgoIDType = '2' // 大连商品交易所
	THOST_FTDC_HTAA_CZCE TThostFtdcHandleTradingAccountAlgoIDType = '3' // 郑州商品交易所
)

// 联系人类型类型
//
//go:generate stringer -type TThostFtdcPersonTypeType -linecomment
type TThostFtdcPersonTypeType byte

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

// 查询范围类型
//
//go:generate stringer -type TThostFtdcQueryInvestorRangeType -linecomment
type TThostFtdcQueryInvestorRangeType byte

const (
	THOST_FTDC_QIR_All    TThostFtdcQueryInvestorRangeType = '1' // 所有
	THOST_FTDC_QIR_Group  TThostFtdcQueryInvestorRangeType = '2' // 查询分类
	THOST_FTDC_QIR_Single TThostFtdcQueryInvestorRangeType = '3' // 单一投资者
)

// 投资者风险状态类型
//
//go:generate stringer -type TThostFtdcInvestorRiskStatusType -linecomment
type TThostFtdcInvestorRiskStatusType byte

const (
	THOST_FTDC_IRS_Normal    TThostFtdcInvestorRiskStatusType = '1' // 正常
	THOST_FTDC_IRS_Warn      TThostFtdcInvestorRiskStatusType = '2' // 警告
	THOST_FTDC_IRS_Call      TThostFtdcInvestorRiskStatusType = '3' // 追保
	THOST_FTDC_IRS_Force     TThostFtdcInvestorRiskStatusType = '4' // 强平
	THOST_FTDC_IRS_Exception TThostFtdcInvestorRiskStatusType = '5' // 异常
)

// 单腿编号类型
type TThostFtdcLegIDType int32

// 单腿乘数类型
type TThostFtdcLegMultipleType int32

// 派生层数类型
type TThostFtdcImplyLevelType int32

// 结算账户类型
type TThostFtdcClearAccountType Byte33

// 结算账户类型
type TThostFtdcOrganNOType Byte6

// 结算账户联行号类型
type TThostFtdcClearbarchIDType Byte6

// 用户事件类型类型
type TThostFtdcUserEventTypeType byte

const THOST_FTDC_UET_Login TThostFtdcUserEventTypeType = '1' // 登录

const THOST_FTDC_UET_Logout TThostFtdcUserEventTypeType = '2' // 登出

const THOST_FTDC_UET_Trading TThostFtdcUserEventTypeType = '3' // CTP校验通过

const THOST_FTDC_UET_TradingError TThostFtdcUserEventTypeType = '4' // CTP校验失败

const THOST_FTDC_UET_UpdatePassword TThostFtdcUserEventTypeType = '5' // 修改密码

const THOST_FTDC_UET_Authenticate TThostFtdcUserEventTypeType = '6' // 客户端认证

const THOST_FTDC_UET_SubmitSysInfo TThostFtdcUserEventTypeType = '7' // 终端信息上报

const THOST_FTDC_UET_Transfer TThostFtdcUserEventTypeType = '8' // 转账

const THOST_FTDC_UET_Other TThostFtdcUserEventTypeType = '9' // 其他

const THOST_FTDC_UET_UpdateTradingAccountPassword TThostFtdcUserEventTypeType = 'a' // 修改资金密码

var mpTThostFtdcUserEventTypeType = map[TThostFtdcUserEventTypeType]string{'1': "THOST_FTDC_UET_Login", '2': "THOST_FTDC_UET_Logout", '3': "THOST_FTDC_UET_Trading", '4': "THOST_FTDC_UET_TradingError", '5': "THOST_FTDC_UET_UpdatePassword", '6': "THOST_FTDC_UET_Authenticate", '7': "THOST_FTDC_UET_SubmitSysInfo", '8': "THOST_FTDC_UET_Transfer", '9': "THOST_FTDC_UET_Other", 'a': "THOST_FTDC_UET_UpdateTradingAccountPassword"}

func (e TThostFtdcUserEventTypeType) String() string {
	if s, ok := mpTThostFtdcUserEventTypeType[e]; ok {
		return s[strings.LastIndex(s, "_")+1:]
	}
	return string(e) + "值未定义"
}

// 用户事件信息类型
type TThostFtdcUserEventInfoType Byte1025

// 平仓方式类型
type TThostFtdcCloseStyleType byte

const THOST_FTDC_ICS_Close TThostFtdcCloseStyleType = '0' // 先开先平

const THOST_FTDC_ICS_CloseToday TThostFtdcCloseStyleType = '1' // 先平今再平昨

var mpTThostFtdcCloseStyleType = map[TThostFtdcCloseStyleType]string{'0': "THOST_FTDC_ICS_Close", '1': "THOST_FTDC_ICS_CloseToday"}

func (e TThostFtdcCloseStyleType) String() string {
	if s, ok := mpTThostFtdcCloseStyleType[e]; ok {
		return s[strings.LastIndex(s, "_")+1:]
	}
	return string(e) + "值未定义"
}

// 统计方式类型
type TThostFtdcStatModeType byte

const THOST_FTDC_SM_Non TThostFtdcStatModeType = '0' // ----

const THOST_FTDC_SM_Instrument TThostFtdcStatModeType = '1' // 按合约统计

const THOST_FTDC_SM_Product TThostFtdcStatModeType = '2' // 按产品统计

const THOST_FTDC_SM_Investor TThostFtdcStatModeType = '3' // 按投资者统计

var mpTThostFtdcStatModeType = map[TThostFtdcStatModeType]string{'0': "THOST_FTDC_SM_Non", '1': "THOST_FTDC_SM_Instrument", '2': "THOST_FTDC_SM_Product", '3': "THOST_FTDC_SM_Investor"}

func (e TThostFtdcStatModeType) String() string {
	if s, ok := mpTThostFtdcStatModeType[e]; ok {
		return s[strings.LastIndex(s, "_")+1:]
	}
	return string(e) + "值未定义"
}

// 预埋单状态类型
type TThostFtdcParkedOrderStatusType byte

const THOST_FTDC_PAOS_NotSend TThostFtdcParkedOrderStatusType = '1' // 未发送

const THOST_FTDC_PAOS_Send TThostFtdcParkedOrderStatusType = '2' // 已发送

const THOST_FTDC_PAOS_Deleted TThostFtdcParkedOrderStatusType = '3' // 已删除

var mpTThostFtdcParkedOrderStatusType = map[TThostFtdcParkedOrderStatusType]string{'1': "THOST_FTDC_PAOS_NotSend", '2': "THOST_FTDC_PAOS_Send", '3': "THOST_FTDC_PAOS_Deleted"}

func (e TThostFtdcParkedOrderStatusType) String() string {
	if s, ok := mpTThostFtdcParkedOrderStatusType[e]; ok {
		return s[strings.LastIndex(s, "_")+1:]
	}
	return string(e) + "值未定义"
}

// 预埋报单编号类型
type TThostFtdcParkedOrderIDType Byte13

// 预埋撤单编号类型
type TThostFtdcParkedOrderActionIDType Byte13

// 处理状态类型
type TThostFtdcVirDealStatusType byte

const THOST_FTDC_VDS_Dealing TThostFtdcVirDealStatusType = '1' // 正在处理

const THOST_FTDC_VDS_DeaclSucceed TThostFtdcVirDealStatusType = '2' // 处理成功

var mpTThostFtdcVirDealStatusType = map[TThostFtdcVirDealStatusType]string{'1': "THOST_FTDC_VDS_Dealing", '2': "THOST_FTDC_VDS_DeaclSucceed"}

func (e TThostFtdcVirDealStatusType) String() string {
	if s, ok := mpTThostFtdcVirDealStatusType[e]; ok {
		return s[strings.LastIndex(s, "_")+1:]
	}
	return string(e) + "值未定义"
}

// 原有系统代码类型
type TThostFtdcOrgSystemIDType byte

const THOST_FTDC_ORGS_Standard TThostFtdcOrgSystemIDType = '0' // 综合交易平台

const THOST_FTDC_ORGS_ESunny TThostFtdcOrgSystemIDType = '1' // 易盛系统

const THOST_FTDC_ORGS_KingStarV6 TThostFtdcOrgSystemIDType = '2' // 金仕达V6系统

var mpTThostFtdcOrgSystemIDType = map[TThostFtdcOrgSystemIDType]string{'0': "THOST_FTDC_ORGS_Standard", '1': "THOST_FTDC_ORGS_ESunny", '2': "THOST_FTDC_ORGS_KingStarV6"}

func (e TThostFtdcOrgSystemIDType) String() string {
	if s, ok := mpTThostFtdcOrgSystemIDType[e]; ok {
		return s[strings.LastIndex(s, "_")+1:]
	}
	return string(e) + "值未定义"
}

// 交易状态类型
type TThostFtdcVirTradeStatusType byte

const THOST_FTDC_VTS_NaturalDeal TThostFtdcVirTradeStatusType = '0' // 正常处理中

const THOST_FTDC_VTS_SucceedEnd TThostFtdcVirTradeStatusType = '1' // 成功结束

const THOST_FTDC_VTS_FailedEND TThostFtdcVirTradeStatusType = '2' // 失败结束

const THOST_FTDC_VTS_Exception TThostFtdcVirTradeStatusType = '3' // 异常中

const THOST_FTDC_VTS_ManualDeal TThostFtdcVirTradeStatusType = '4' // 已人工异常处理

const THOST_FTDC_VTS_MesException TThostFtdcVirTradeStatusType = '5' // 通讯异常 ，请人工处理

const THOST_FTDC_VTS_SysException TThostFtdcVirTradeStatusType = '6' // 系统出错，请人工处理

var mpTThostFtdcVirTradeStatusType = map[TThostFtdcVirTradeStatusType]string{'0': "THOST_FTDC_VTS_NaturalDeal", '1': "THOST_FTDC_VTS_SucceedEnd", '2': "THOST_FTDC_VTS_FailedEND", '3': "THOST_FTDC_VTS_Exception", '4': "THOST_FTDC_VTS_ManualDeal", '5': "THOST_FTDC_VTS_MesException", '6': "THOST_FTDC_VTS_SysException"}

func (e TThostFtdcVirTradeStatusType) String() string {
	if s, ok := mpTThostFtdcVirTradeStatusType[e]; ok {
		return s[strings.LastIndex(s, "_")+1:]
	}
	return string(e) + "值未定义"
}

// 银行帐户类型类型
type TThostFtdcVirBankAccTypeType byte

const THOST_FTDC_VBAT_BankBook TThostFtdcVirBankAccTypeType = '1' // 存折

const THOST_FTDC_VBAT_BankCard TThostFtdcVirBankAccTypeType = '2' // 储蓄卡

const THOST_FTDC_VBAT_CreditCard TThostFtdcVirBankAccTypeType = '3' // 信用卡

var mpTThostFtdcVirBankAccTypeType = map[TThostFtdcVirBankAccTypeType]string{'1': "THOST_FTDC_VBAT_BankBook", '2': "THOST_FTDC_VBAT_BankCard", '3': "THOST_FTDC_VBAT_CreditCard"}

func (e TThostFtdcVirBankAccTypeType) String() string {
	if s, ok := mpTThostFtdcVirBankAccTypeType[e]; ok {
		return s[strings.LastIndex(s, "_")+1:]
	}
	return string(e) + "值未定义"
}

// 银行帐户类型类型
type TThostFtdcVirementStatusType byte

const THOST_FTDC_VMS_Natural TThostFtdcVirementStatusType = '0' // 正常

const THOST_FTDC_VMS_Canceled TThostFtdcVirementStatusType = '9' // 销户

var mpTThostFtdcVirementStatusType = map[TThostFtdcVirementStatusType]string{'0': "THOST_FTDC_VMS_Natural", '9': "THOST_FTDC_VMS_Canceled"}

func (e TThostFtdcVirementStatusType) String() string {
	if s, ok := mpTThostFtdcVirementStatusType[e]; ok {
		return s[strings.LastIndex(s, "_")+1:]
	}
	return string(e) + "值未定义"
}

// 有效标志类型
type TThostFtdcVirementAvailAbilityType byte

const THOST_FTDC_VAA_NoAvailAbility TThostFtdcVirementAvailAbilityType = '0' // 未确认

const THOST_FTDC_VAA_AvailAbility TThostFtdcVirementAvailAbilityType = '1' // 有效

const THOST_FTDC_VAA_Repeal TThostFtdcVirementAvailAbilityType = '2' // 冲正

var mpTThostFtdcVirementAvailAbilityType = map[TThostFtdcVirementAvailAbilityType]string{'0': "THOST_FTDC_VAA_NoAvailAbility", '1': "THOST_FTDC_VAA_AvailAbility", '2': "THOST_FTDC_VAA_Repeal"}

func (e TThostFtdcVirementAvailAbilityType) String() string {
	if s, ok := mpTThostFtdcVirementAvailAbilityType[e]; ok {
		return s[strings.LastIndex(s, "_")+1:]
	}
	return string(e) + "值未定义"
}

// 交易代码类型
type TThostFtdcVirementTradeCodeType string

const THOST_FTDC_VTC_BankBankToFuture TThostFtdcVirementTradeCodeType = "102001" // 银行发起银行资金转期货

const THOST_FTDC_VTC_BankFutureToBank TThostFtdcVirementTradeCodeType = "102002" // 银行发起期货资金转银行

const THOST_FTDC_VTC_FutureBankToFuture TThostFtdcVirementTradeCodeType = "202001" // 期货发起银行资金转期货

const THOST_FTDC_VTC_FutureFutureToBank TThostFtdcVirementTradeCodeType = "202002" // 期货发起期货资金转银行

var mpTThostFtdcVirementTradeCodeType = map[TThostFtdcVirementTradeCodeType]string{"102001": "THOST_FTDC_VTC_BankBankToFuture", "102002": "THOST_FTDC_VTC_BankFutureToBank", "202001": "THOST_FTDC_VTC_FutureBankToFuture", "202002": "THOST_FTDC_VTC_FutureFutureToBank"}

func (e TThostFtdcVirementTradeCodeType) String() string {
	if s, ok := mpTThostFtdcVirementTradeCodeType[e]; ok {
		return s[strings.LastIndex(s, "_")+1:]
	}
	return string(e) + "值未定义"
}

// 影像类型名称类型
type TThostFtdcPhotoTypeNameType Byte41

// 影像类型代码类型
type TThostFtdcPhotoTypeIDType Byte5

// 影像名称类型
type TThostFtdcPhotoNameType Byte161

// 主题代码类型
type TThostFtdcTopicIDType int32

// 交易报告类型标识类型
type TThostFtdcReportTypeIDType Byte3

// 交易特征代码类型
type TThostFtdcCharacterIDType Byte5

// 参数代码类型
type TThostFtdcAMLParamIDType Byte21

// 投资者类型类型
type TThostFtdcAMLInvestorTypeType Byte3

// 证件类型类型
type TThostFtdcAMLIdCardTypeType Byte3

// 资金进出方向类型
type TThostFtdcAMLTradeDirectType Byte3

// 资金进出方式类型
type TThostFtdcAMLTradeModelType Byte3

// 业务参数代码值类型
type TThostFtdcAMLOpParamValueType Double

// 客户身份证件/证明文件类型类型
type TThostFtdcAMLCustomerCardTypeType Byte81

// 金融机构网点名称类型
type TThostFtdcAMLInstitutionNameType Byte65

// 金融机构网点所在地区行政区划代码类型
type TThostFtdcAMLDistrictIDType Byte7

// 金融机构网点与大额交易的关系类型
type TThostFtdcAMLRelationShipType Byte3

// 金融机构网点代码类型类型
type TThostFtdcAMLInstitutionTypeType Byte3

// 金融机构网点代码类型
type TThostFtdcAMLInstitutionIDType Byte13

// 账户类型类型
type TThostFtdcAMLAccountTypeType Byte5

// 交易方式类型
type TThostFtdcAMLTradingTypeType Byte7

// 涉外收支交易分类与代码类型
type TThostFtdcAMLTransactClassType Byte7

// 资金收付标识类型
type TThostFtdcAMLCapitalIOType Byte3

// 交易地点类型
type TThostFtdcAMLSiteType Byte10

// 资金用途类型
type TThostFtdcAMLCapitalPurposeType Byte129

// 报文类型类型
type TThostFtdcAMLReportTypeType Byte2

// 编号类型
type TThostFtdcAMLSerialNoType Byte5

// 状态类型
type TThostFtdcAMLStatusType Byte2

// Aml生成方式类型
type TThostFtdcAMLGenStatusType byte

const THOST_FTDC_GEN_Program TThostFtdcAMLGenStatusType = '0' // 程序生成

const THOST_FTDC_GEN_HandWork TThostFtdcAMLGenStatusType = '1' // 人工生成

var mpTThostFtdcAMLGenStatusType = map[TThostFtdcAMLGenStatusType]string{'0': "THOST_FTDC_GEN_Program", '1': "THOST_FTDC_GEN_HandWork"}

func (e TThostFtdcAMLGenStatusType) String() string {
	if s, ok := mpTThostFtdcAMLGenStatusType[e]; ok {
		return s[strings.LastIndex(s, "_")+1:]
	}
	return string(e) + "值未定义"
}

// 业务标识号类型
type TThostFtdcAMLSeqCodeType Byte65

// AML文件名类型
type TThostFtdcAMLFileNameType Byte257

// 反洗钱资金类型
type TThostFtdcAMLMoneyType Double

// 反洗钱资金类型
type TThostFtdcAMLFileAmountType int32

// 密钥类型(保证金监管)类型
type TThostFtdcCFMMCKeyType Byte21

// 令牌类型(保证金监管)类型
type TThostFtdcCFMMCTokenType Byte21

// 动态密钥类别(保证金监管)类型
type TThostFtdcCFMMCKeyKindType byte

const THOST_FTDC_CFMMCKK_REQUEST TThostFtdcCFMMCKeyKindType = 'R' // 主动请求更新

const THOST_FTDC_CFMMCKK_AUTO TThostFtdcCFMMCKeyKindType = 'A' // CFMMC自动更新

const THOST_FTDC_CFMMCKK_MANUAL TThostFtdcCFMMCKeyKindType = 'M' // CFMMC手动更新

var mpTThostFtdcCFMMCKeyKindType = map[TThostFtdcCFMMCKeyKindType]string{'R': "THOST_FTDC_CFMMCKK_REQUEST", 'A': "THOST_FTDC_CFMMCKK_AUTO", 'M': "THOST_FTDC_CFMMCKK_MANUAL"}

func (e TThostFtdcCFMMCKeyKindType) String() string {
	if s, ok := mpTThostFtdcCFMMCKeyKindType[e]; ok {
		return s[strings.LastIndex(s, "_")+1:]
	}
	return string(e) + "值未定义"
}

// 报文名称类型
type TThostFtdcAMLReportNameType Byte81

// 个人姓名类型
type TThostFtdcIndividualNameType Byte51

// 币种代码类型
type TThostFtdcCurrencyIDType Byte4

// 客户编号类型
type TThostFtdcCustNumberType Byte36

// 机构编码类型
type TThostFtdcOrganCodeType Byte36

// 机构名称类型
type TThostFtdcOrganNameType Byte71

// 上级机构编码,即期货公司总部、银行总行类型
type TThostFtdcSuperOrganCodeType Byte12

// 分支机构类型
type TThostFtdcSubBranchIDType Byte31

// 分支机构名称类型
type TThostFtdcSubBranchNameType Byte71

// 机构网点号类型
type TThostFtdcBranchNetCodeType Byte31

// 机构网点名称类型
type TThostFtdcBranchNetNameType Byte71

// 机构标识类型
type TThostFtdcOrganFlagType Byte2

// 银行对期货公司的编码类型
type TThostFtdcBankCodingForFutureType Byte33

// 银行对返回码的定义类型
type TThostFtdcBankReturnCodeType Byte7

// 银期转帐平台对返回码的定义类型
type TThostFtdcPlateReturnCodeType Byte5

// 银行分支机构编码类型
type TThostFtdcBankSubBranchIDType Byte31

// 期货分支机构编码类型
type TThostFtdcFutureBranchIDType Byte31

// 返回代码类型
type TThostFtdcReturnCodeType Byte7

// 操作员类型
type TThostFtdcOperatorCodeType Byte17

// 机构结算帐户机构号类型
type TThostFtdcClearDepIDType Byte6

// 机构结算帐户联行号类型
type TThostFtdcClearBrchIDType Byte6

// 机构结算帐户名称类型
type TThostFtdcClearNameType Byte71

// 银行帐户名称类型
type TThostFtdcBankAccountNameType Byte71

// 机构投资人账号机构号类型
type TThostFtdcInvDepIDType Byte6

// 机构投资人联行号类型
type TThostFtdcInvBrchIDType Byte6

// 信息格式版本类型
type TThostFtdcMessageFormatVersionType Byte36

// 摘要类型
type TThostFtdcDigestType Byte36

// 认证数据类型
type TThostFtdcAuthenticDataType Byte129

// 密钥类型
type TThostFtdcPasswordKeyType Byte129

// 期货帐户名称类型
type TThostFtdcFutureAccountNameType Byte129

// 手机类型
type TThostFtdcMobilePhoneType Byte21

// 期货公司主密钥类型
type TThostFtdcFutureMainKeyType Byte129

// 期货公司工作密钥类型
type TThostFtdcFutureWorkKeyType Byte129

// 期货公司传输密钥类型
type TThostFtdcFutureTransKeyType Byte129

// 银行主密钥类型
type TThostFtdcBankMainKeyType Byte129

// 银行工作密钥类型
type TThostFtdcBankWorkKeyType Byte129

// 银行传输密钥类型
type TThostFtdcBankTransKeyType Byte129

// 银行服务器描述信息类型
type TThostFtdcBankServerDescriptionType Byte129

// 附加信息类型
type TThostFtdcAddInfoType Byte129

// 返回码描述类型
type TThostFtdcDescrInfoForReturnCodeType Byte129

// 国家代码类型
type TThostFtdcCountryCodeType Byte21

// 流水号类型
type TThostFtdcSerialType int32

// 平台流水号类型
type TThostFtdcPlateSerialType int32

// 银行流水号类型
type TThostFtdcBankSerialType Byte13

// 被冲正交易流水号类型
type TThostFtdcCorrectSerialType int32

// 期货公司流水号类型
type TThostFtdcFutureSerialType int32

// 应用标识类型
type TThostFtdcApplicationIDType int32

// 银行代理标识类型
type TThostFtdcBankProxyIDType int32

// 银期转帐核心系统标识类型
type TThostFtdcFBTCoreIDType int32

// 服务端口号类型
type TThostFtdcServerPortType int32

// 已经冲正次数类型
type TThostFtdcRepealedTimesType int32

// 冲正时间间隔类型
type TThostFtdcRepealTimeIntervalType int32

// 每日累计转帐次数类型
type TThostFtdcTotalTimesType int32

// 请求ID类型
type TThostFtdcFBTRequestIDType int32

// 交易ID类型
type TThostFtdcTIDType int32

// 交易金额（元）类型
type TThostFtdcTradeAmountType Double

// 应收客户费用（元）类型
type TThostFtdcCustFeeType Double

// 应收期货公司费用（元）类型
type TThostFtdcFutureFeeType Double

// 单笔最高限额类型
type TThostFtdcSingleMaxAmtType Double

// 单笔最低限额类型
type TThostFtdcSingleMinAmtType Double

// 每日累计转帐额度类型
type TThostFtdcTotalAmtType Double

// 证件类型类型
type TThostFtdcCertificationTypeType byte

const THOST_FTDC_CFT_IDCard TThostFtdcCertificationTypeType = '0' // 身份证

const THOST_FTDC_CFT_Passport TThostFtdcCertificationTypeType = '1' // 护照

const THOST_FTDC_CFT_OfficerIDCard TThostFtdcCertificationTypeType = '2' // 军官证

const THOST_FTDC_CFT_SoldierIDCard TThostFtdcCertificationTypeType = '3' // 士兵证

const THOST_FTDC_CFT_HomeComingCard TThostFtdcCertificationTypeType = '4' // 回乡证

const THOST_FTDC_CFT_HouseholdRegister TThostFtdcCertificationTypeType = '5' // 户口簿

const THOST_FTDC_CFT_LicenseNo TThostFtdcCertificationTypeType = '6' // 营业执照号

const THOST_FTDC_CFT_InstitutionCodeCard TThostFtdcCertificationTypeType = '7' // 组织机构代码证

const THOST_FTDC_CFT_TempLicenseNo TThostFtdcCertificationTypeType = '8' // 临时营业执照号

const THOST_FTDC_CFT_NoEnterpriseLicenseNo TThostFtdcCertificationTypeType = '9' // 民办非企业登记证书

const THOST_FTDC_CFT_OtherCard TThostFtdcCertificationTypeType = 'x' // 其他证件

const THOST_FTDC_CFT_SuperDepAgree TThostFtdcCertificationTypeType = 'a' // 主管部门批文

var mpTThostFtdcCertificationTypeType = map[TThostFtdcCertificationTypeType]string{'0': "THOST_FTDC_CFT_IDCard", '1': "THOST_FTDC_CFT_Passport", '2': "THOST_FTDC_CFT_OfficerIDCard", '3': "THOST_FTDC_CFT_SoldierIDCard", '4': "THOST_FTDC_CFT_HomeComingCard", '5': "THOST_FTDC_CFT_HouseholdRegister", '6': "THOST_FTDC_CFT_LicenseNo", '7': "THOST_FTDC_CFT_InstitutionCodeCard", '8': "THOST_FTDC_CFT_TempLicenseNo", '9': "THOST_FTDC_CFT_NoEnterpriseLicenseNo", 'x': "THOST_FTDC_CFT_OtherCard", 'a': "THOST_FTDC_CFT_SuperDepAgree"}

func (e TThostFtdcCertificationTypeType) String() string {
	if s, ok := mpTThostFtdcCertificationTypeType[e]; ok {
		return s[strings.LastIndex(s, "_")+1:]
	}
	return string(e) + "值未定义"
}

// 文件业务功能类型
type TThostFtdcFileBusinessCodeType byte

const THOST_FTDC_FBC_Others TThostFtdcFileBusinessCodeType = '0' // 其他

const THOST_FTDC_FBC_TransferDetails TThostFtdcFileBusinessCodeType = '1' // 转账交易明细对账

const THOST_FTDC_FBC_CustAccStatus TThostFtdcFileBusinessCodeType = '2' // 客户账户状态对账

const THOST_FTDC_FBC_AccountTradeDetails TThostFtdcFileBusinessCodeType = '3' // 账户类交易明细对账

const THOST_FTDC_FBC_FutureAccountChangeInfoDetails TThostFtdcFileBusinessCodeType = '4' // 期货账户信息变更明细对账

const THOST_FTDC_FBC_CustMoneyDetail TThostFtdcFileBusinessCodeType = '5' // 客户资金台账余额明细对账

const THOST_FTDC_FBC_CustCancelAccountInfo TThostFtdcFileBusinessCodeType = '6' // 客户销户结息明细对账

const THOST_FTDC_FBC_CustMoneyResult TThostFtdcFileBusinessCodeType = '7' // 客户资金余额对账结果

const THOST_FTDC_FBC_OthersExceptionResult TThostFtdcFileBusinessCodeType = '8' // 其它对账异常结果文件

const THOST_FTDC_FBC_CustInterestNetMoneyDetails TThostFtdcFileBusinessCodeType = '9' // 客户结息净额明细

const THOST_FTDC_FBC_CustMoneySendAndReceiveDetails TThostFtdcFileBusinessCodeType = 'a' // 客户资金交收明细

const THOST_FTDC_FBC_CorporationMoneyTotal TThostFtdcFileBusinessCodeType = 'b' // 法人存管银行资金交收汇总

const THOST_FTDC_FBC_MainbodyMoneyTotal TThostFtdcFileBusinessCodeType = 'c' // 主体间资金交收汇总

const THOST_FTDC_FBC_MainPartMonitorData TThostFtdcFileBusinessCodeType = 'd' // 总分平衡监管数据

const THOST_FTDC_FBC_PreparationMoney TThostFtdcFileBusinessCodeType = 'e' // 存管银行备付金余额

const THOST_FTDC_FBC_BankMoneyMonitorData TThostFtdcFileBusinessCodeType = 'f' // 协办存管银行资金监管数据

var mpTThostFtdcFileBusinessCodeType = map[TThostFtdcFileBusinessCodeType]string{'0': "THOST_FTDC_FBC_Others", '1': "THOST_FTDC_FBC_TransferDetails", '2': "THOST_FTDC_FBC_CustAccStatus", '3': "THOST_FTDC_FBC_AccountTradeDetails", '4': "THOST_FTDC_FBC_FutureAccountChangeInfoDetails", '5': "THOST_FTDC_FBC_CustMoneyDetail", '6': "THOST_FTDC_FBC_CustCancelAccountInfo", '7': "THOST_FTDC_FBC_CustMoneyResult", '8': "THOST_FTDC_FBC_OthersExceptionResult", '9': "THOST_FTDC_FBC_CustInterestNetMoneyDetails", 'a': "THOST_FTDC_FBC_CustMoneySendAndReceiveDetails", 'b': "THOST_FTDC_FBC_CorporationMoneyTotal", 'c': "THOST_FTDC_FBC_MainbodyMoneyTotal", 'd': "THOST_FTDC_FBC_MainPartMonitorData", 'e': "THOST_FTDC_FBC_PreparationMoney", 'f': "THOST_FTDC_FBC_BankMoneyMonitorData"}

func (e TThostFtdcFileBusinessCodeType) String() string {
	if s, ok := mpTThostFtdcFileBusinessCodeType[e]; ok {
		return s[strings.LastIndex(s, "_")+1:]
	}
	return string(e) + "值未定义"
}

// 汇钞标志类型
type TThostFtdcCashExchangeCodeType byte

const THOST_FTDC_CEC_Exchange TThostFtdcCashExchangeCodeType = '1' // 汇

const THOST_FTDC_CEC_Cash TThostFtdcCashExchangeCodeType = '2' // 钞

var mpTThostFtdcCashExchangeCodeType = map[TThostFtdcCashExchangeCodeType]string{'1': "THOST_FTDC_CEC_Exchange", '2': "THOST_FTDC_CEC_Cash"}

func (e TThostFtdcCashExchangeCodeType) String() string {
	if s, ok := mpTThostFtdcCashExchangeCodeType[e]; ok {
		return s[strings.LastIndex(s, "_")+1:]
	}
	return string(e) + "值未定义"
}

// 是或否标识类型
type TThostFtdcYesNoIndicatorType byte

const THOST_FTDC_YNI_Yes TThostFtdcYesNoIndicatorType = '0' // 是

const THOST_FTDC_YNI_No TThostFtdcYesNoIndicatorType = '1' // 否

var mpTThostFtdcYesNoIndicatorType = map[TThostFtdcYesNoIndicatorType]string{'0': "THOST_FTDC_YNI_Yes", '1': "THOST_FTDC_YNI_No"}

func (e TThostFtdcYesNoIndicatorType) String() string {
	if s, ok := mpTThostFtdcYesNoIndicatorType[e]; ok {
		return s[strings.LastIndex(s, "_")+1:]
	}
	return string(e) + "值未定义"
}

// 余额类型类型
type TThostFtdcBanlanceTypeType byte

const THOST_FTDC_BLT_CurrentMoney TThostFtdcBanlanceTypeType = '0' // 当前余额

const THOST_FTDC_BLT_UsableMoney TThostFtdcBanlanceTypeType = '1' // 可用余额

const THOST_FTDC_BLT_FetchableMoney TThostFtdcBanlanceTypeType = '2' // 可取余额

const THOST_FTDC_BLT_FreezeMoney TThostFtdcBanlanceTypeType = '3' // 冻结余额

var mpTThostFtdcBanlanceTypeType = map[TThostFtdcBanlanceTypeType]string{'0': "THOST_FTDC_BLT_CurrentMoney", '1': "THOST_FTDC_BLT_UsableMoney", '2': "THOST_FTDC_BLT_FetchableMoney", '3': "THOST_FTDC_BLT_FreezeMoney"}

func (e TThostFtdcBanlanceTypeType) String() string {
	if s, ok := mpTThostFtdcBanlanceTypeType[e]; ok {
		return s[strings.LastIndex(s, "_")+1:]
	}
	return string(e) + "值未定义"
}

// 性别类型
type TThostFtdcGenderType byte

const THOST_FTDC_GD_Unknown TThostFtdcGenderType = '0' // 未知状态

const THOST_FTDC_GD_Male TThostFtdcGenderType = '1' // 男

const THOST_FTDC_GD_Female TThostFtdcGenderType = '2' // 女

var mpTThostFtdcGenderType = map[TThostFtdcGenderType]string{'0': "THOST_FTDC_GD_Unknown", '1': "THOST_FTDC_GD_Male", '2': "THOST_FTDC_GD_Female"}

func (e TThostFtdcGenderType) String() string {
	if s, ok := mpTThostFtdcGenderType[e]; ok {
		return s[strings.LastIndex(s, "_")+1:]
	}
	return string(e) + "值未定义"
}

// 费用支付标志类型
type TThostFtdcFeePayFlagType byte

const THOST_FTDC_FPF_BEN TThostFtdcFeePayFlagType = '0' // 由受益方支付费用

const THOST_FTDC_FPF_OUR TThostFtdcFeePayFlagType = '1' // 由发送方支付费用

const THOST_FTDC_FPF_SHA TThostFtdcFeePayFlagType = '2' // 由发送方支付发起的费用，受益方支付接受的费用

var mpTThostFtdcFeePayFlagType = map[TThostFtdcFeePayFlagType]string{'0': "THOST_FTDC_FPF_BEN", '1': "THOST_FTDC_FPF_OUR", '2': "THOST_FTDC_FPF_SHA"}

func (e TThostFtdcFeePayFlagType) String() string {
	if s, ok := mpTThostFtdcFeePayFlagType[e]; ok {
		return s[strings.LastIndex(s, "_")+1:]
	}
	return string(e) + "值未定义"
}

// 密钥类型类型
type TThostFtdcPassWordKeyTypeType byte

const THOST_FTDC_PWKT_ExchangeKey TThostFtdcPassWordKeyTypeType = '0' // 交换密钥

const THOST_FTDC_PWKT_PassWordKey TThostFtdcPassWordKeyTypeType = '1' // 密码密钥

const THOST_FTDC_PWKT_MACKey TThostFtdcPassWordKeyTypeType = '2' // MAC密钥

const THOST_FTDC_PWKT_MessageKey TThostFtdcPassWordKeyTypeType = '3' // 报文密钥

var mpTThostFtdcPassWordKeyTypeType = map[TThostFtdcPassWordKeyTypeType]string{'0': "THOST_FTDC_PWKT_ExchangeKey", '1': "THOST_FTDC_PWKT_PassWordKey", '2': "THOST_FTDC_PWKT_MACKey", '3': "THOST_FTDC_PWKT_MessageKey"}

func (e TThostFtdcPassWordKeyTypeType) String() string {
	if s, ok := mpTThostFtdcPassWordKeyTypeType[e]; ok {
		return s[strings.LastIndex(s, "_")+1:]
	}
	return string(e) + "值未定义"
}

// 密码类型类型
type TThostFtdcFBTPassWordTypeType byte

const THOST_FTDC_PWT_Query TThostFtdcFBTPassWordTypeType = '0' // 查询

const THOST_FTDC_PWT_Fetch TThostFtdcFBTPassWordTypeType = '1' // 取款

const THOST_FTDC_PWT_Transfer TThostFtdcFBTPassWordTypeType = '2' // 转帐

const THOST_FTDC_PWT_Trade TThostFtdcFBTPassWordTypeType = '3' // 交易

var mpTThostFtdcFBTPassWordTypeType = map[TThostFtdcFBTPassWordTypeType]string{'0': "THOST_FTDC_PWT_Query", '1': "THOST_FTDC_PWT_Fetch", '2': "THOST_FTDC_PWT_Transfer", '3': "THOST_FTDC_PWT_Trade"}

func (e TThostFtdcFBTPassWordTypeType) String() string {
	if s, ok := mpTThostFtdcFBTPassWordTypeType[e]; ok {
		return s[strings.LastIndex(s, "_")+1:]
	}
	return string(e) + "值未定义"
}

// 加密方式类型
type TThostFtdcFBTEncryModeType byte

const THOST_FTDC_EM_NoEncry TThostFtdcFBTEncryModeType = '0' // 不加密

const THOST_FTDC_EM_DES TThostFtdcFBTEncryModeType = '1' // DES

const THOST_FTDC_EM_3DES TThostFtdcFBTEncryModeType = '2' // 3DES

var mpTThostFtdcFBTEncryModeType = map[TThostFtdcFBTEncryModeType]string{'0': "THOST_FTDC_EM_NoEncry", '1': "THOST_FTDC_EM_DES", '2': "THOST_FTDC_EM_3DES"}

func (e TThostFtdcFBTEncryModeType) String() string {
	if s, ok := mpTThostFtdcFBTEncryModeType[e]; ok {
		return s[strings.LastIndex(s, "_")+1:]
	}
	return string(e) + "值未定义"
}

// 银行冲正标志类型
type TThostFtdcBankRepealFlagType byte

const THOST_FTDC_BRF_BankNotNeedRepeal TThostFtdcBankRepealFlagType = '0' // 银行无需自动冲正

const THOST_FTDC_BRF_BankWaitingRepeal TThostFtdcBankRepealFlagType = '1' // 银行待自动冲正

const THOST_FTDC_BRF_BankBeenRepealed TThostFtdcBankRepealFlagType = '2' // 银行已自动冲正

var mpTThostFtdcBankRepealFlagType = map[TThostFtdcBankRepealFlagType]string{'0': "THOST_FTDC_BRF_BankNotNeedRepeal", '1': "THOST_FTDC_BRF_BankWaitingRepeal", '2': "THOST_FTDC_BRF_BankBeenRepealed"}

func (e TThostFtdcBankRepealFlagType) String() string {
	if s, ok := mpTThostFtdcBankRepealFlagType[e]; ok {
		return s[strings.LastIndex(s, "_")+1:]
	}
	return string(e) + "值未定义"
}

// 期商冲正标志类型
type TThostFtdcBrokerRepealFlagType byte

const THOST_FTDC_BRORF_BrokerNotNeedRepeal TThostFtdcBrokerRepealFlagType = '0' // 期商无需自动冲正

const THOST_FTDC_BRORF_BrokerWaitingRepeal TThostFtdcBrokerRepealFlagType = '1' // 期商待自动冲正

const THOST_FTDC_BRORF_BrokerBeenRepealed TThostFtdcBrokerRepealFlagType = '2' // 期商已自动冲正

var mpTThostFtdcBrokerRepealFlagType = map[TThostFtdcBrokerRepealFlagType]string{'0': "THOST_FTDC_BRORF_BrokerNotNeedRepeal", '1': "THOST_FTDC_BRORF_BrokerWaitingRepeal", '2': "THOST_FTDC_BRORF_BrokerBeenRepealed"}

func (e TThostFtdcBrokerRepealFlagType) String() string {
	if s, ok := mpTThostFtdcBrokerRepealFlagType[e]; ok {
		return s[strings.LastIndex(s, "_")+1:]
	}
	return string(e) + "值未定义"
}

// 机构类别类型
type TThostFtdcInstitutionTypeType byte

const THOST_FTDC_TS_Bank TThostFtdcInstitutionTypeType = '0' // 银行

const THOST_FTDC_TS_Future TThostFtdcInstitutionTypeType = '1' // 期商

const THOST_FTDC_TS_Store TThostFtdcInstitutionTypeType = '2' // 券商

var mpTThostFtdcInstitutionTypeType = map[TThostFtdcInstitutionTypeType]string{'0': "THOST_FTDC_TS_Bank", '1': "THOST_FTDC_TS_Future", '2': "THOST_FTDC_TS_Store"}

func (e TThostFtdcInstitutionTypeType) String() string {
	if s, ok := mpTThostFtdcInstitutionTypeType[e]; ok {
		return s[strings.LastIndex(s, "_")+1:]
	}
	return string(e) + "值未定义"
}

// 最后分片标志类型
type TThostFtdcLastFragmentType byte

const THOST_FTDC_LF_Yes TThostFtdcLastFragmentType = '0' // 是最后分片

const THOST_FTDC_LF_No TThostFtdcLastFragmentType = '1' // 不是最后分片

var mpTThostFtdcLastFragmentType = map[TThostFtdcLastFragmentType]string{'0': "THOST_FTDC_LF_Yes", '1': "THOST_FTDC_LF_No"}

func (e TThostFtdcLastFragmentType) String() string {
	if s, ok := mpTThostFtdcLastFragmentType[e]; ok {
		return s[strings.LastIndex(s, "_")+1:]
	}
	return string(e) + "值未定义"
}

// 银行账户状态类型
type TThostFtdcBankAccStatusType byte

const THOST_FTDC_BAS_Normal TThostFtdcBankAccStatusType = '0' // 正常

const THOST_FTDC_BAS_Freeze TThostFtdcBankAccStatusType = '1' // 冻结

const THOST_FTDC_BAS_ReportLoss TThostFtdcBankAccStatusType = '2' // 挂失

var mpTThostFtdcBankAccStatusType = map[TThostFtdcBankAccStatusType]string{'0': "THOST_FTDC_BAS_Normal", '1': "THOST_FTDC_BAS_Freeze", '2': "THOST_FTDC_BAS_ReportLoss"}

func (e TThostFtdcBankAccStatusType) String() string {
	if s, ok := mpTThostFtdcBankAccStatusType[e]; ok {
		return s[strings.LastIndex(s, "_")+1:]
	}
	return string(e) + "值未定义"
}

// 资金账户状态类型
type TThostFtdcMoneyAccountStatusType byte

const THOST_FTDC_MAS_Normal TThostFtdcMoneyAccountStatusType = '0' // 正常

const THOST_FTDC_MAS_Cancel TThostFtdcMoneyAccountStatusType = '1' // 销户

var mpTThostFtdcMoneyAccountStatusType = map[TThostFtdcMoneyAccountStatusType]string{'0': "THOST_FTDC_MAS_Normal", '1': "THOST_FTDC_MAS_Cancel"}

func (e TThostFtdcMoneyAccountStatusType) String() string {
	if s, ok := mpTThostFtdcMoneyAccountStatusType[e]; ok {
		return s[strings.LastIndex(s, "_")+1:]
	}
	return string(e) + "值未定义"
}

// 存管状态类型
type TThostFtdcManageStatusType byte

const THOST_FTDC_MSS_Point TThostFtdcManageStatusType = '0' // 指定存管

const THOST_FTDC_MSS_PrePoint TThostFtdcManageStatusType = '1' // 预指定

const THOST_FTDC_MSS_CancelPoint TThostFtdcManageStatusType = '2' // 撤销指定

var mpTThostFtdcManageStatusType = map[TThostFtdcManageStatusType]string{'0': "THOST_FTDC_MSS_Point", '1': "THOST_FTDC_MSS_PrePoint", '2': "THOST_FTDC_MSS_CancelPoint"}

func (e TThostFtdcManageStatusType) String() string {
	if s, ok := mpTThostFtdcManageStatusType[e]; ok {
		return s[strings.LastIndex(s, "_")+1:]
	}
	return string(e) + "值未定义"
}

// 应用系统类型类型
type TThostFtdcSystemTypeType byte

const THOST_FTDC_SYT_FutureBankTransfer TThostFtdcSystemTypeType = '0' // 银期转帐

const THOST_FTDC_SYT_StockBankTransfer TThostFtdcSystemTypeType = '1' // 银证转帐

const THOST_FTDC_SYT_TheThirdPartStore TThostFtdcSystemTypeType = '2' // 第三方存管

var mpTThostFtdcSystemTypeType = map[TThostFtdcSystemTypeType]string{'0': "THOST_FTDC_SYT_FutureBankTransfer", '1': "THOST_FTDC_SYT_StockBankTransfer", '2': "THOST_FTDC_SYT_TheThirdPartStore"}

func (e TThostFtdcSystemTypeType) String() string {
	if s, ok := mpTThostFtdcSystemTypeType[e]; ok {
		return s[strings.LastIndex(s, "_")+1:]
	}
	return string(e) + "值未定义"
}

// 银期转帐划转结果标志类型
type TThostFtdcTxnEndFlagType byte

const THOST_FTDC_TEF_NormalProcessing TThostFtdcTxnEndFlagType = '0' // 正常处理中

const THOST_FTDC_TEF_Success TThostFtdcTxnEndFlagType = '1' // 成功结束

const THOST_FTDC_TEF_Failed TThostFtdcTxnEndFlagType = '2' // 失败结束

const THOST_FTDC_TEF_Abnormal TThostFtdcTxnEndFlagType = '3' // 异常中

const THOST_FTDC_TEF_ManualProcessedForException TThostFtdcTxnEndFlagType = '4' // 已人工异常处理

const THOST_FTDC_TEF_CommuFailedNeedManualProcess TThostFtdcTxnEndFlagType = '5' // 通讯异常 ，请人工处理

const THOST_FTDC_TEF_SysErrorNeedManualProcess TThostFtdcTxnEndFlagType = '6' // 系统出错，请人工处理

var mpTThostFtdcTxnEndFlagType = map[TThostFtdcTxnEndFlagType]string{'0': "THOST_FTDC_TEF_NormalProcessing", '1': "THOST_FTDC_TEF_Success", '2': "THOST_FTDC_TEF_Failed", '3': "THOST_FTDC_TEF_Abnormal", '4': "THOST_FTDC_TEF_ManualProcessedForException", '5': "THOST_FTDC_TEF_CommuFailedNeedManualProcess", '6': "THOST_FTDC_TEF_SysErrorNeedManualProcess"}

func (e TThostFtdcTxnEndFlagType) String() string {
	if s, ok := mpTThostFtdcTxnEndFlagType[e]; ok {
		return s[strings.LastIndex(s, "_")+1:]
	}
	return string(e) + "值未定义"
}

// 银期转帐服务处理状态类型
type TThostFtdcProcessStatusType byte

const THOST_FTDC_PSS_NotProcess TThostFtdcProcessStatusType = '0' // 未处理

const THOST_FTDC_PSS_StartProcess TThostFtdcProcessStatusType = '1' // 开始处理

const THOST_FTDC_PSS_Finished TThostFtdcProcessStatusType = '2' // 处理完成

var mpTThostFtdcProcessStatusType = map[TThostFtdcProcessStatusType]string{'0': "THOST_FTDC_PSS_NotProcess", '1': "THOST_FTDC_PSS_StartProcess", '2': "THOST_FTDC_PSS_Finished"}

func (e TThostFtdcProcessStatusType) String() string {
	if s, ok := mpTThostFtdcProcessStatusType[e]; ok {
		return s[strings.LastIndex(s, "_")+1:]
	}
	return string(e) + "值未定义"
}

// 客户类型类型
type TThostFtdcCustTypeType byte

const THOST_FTDC_CUSTT_Person TThostFtdcCustTypeType = '0' // 自然人

const THOST_FTDC_CUSTT_Institution TThostFtdcCustTypeType = '1' // 机构户

var mpTThostFtdcCustTypeType = map[TThostFtdcCustTypeType]string{'0': "THOST_FTDC_CUSTT_Person", '1': "THOST_FTDC_CUSTT_Institution"}

func (e TThostFtdcCustTypeType) String() string {
	if s, ok := mpTThostFtdcCustTypeType[e]; ok {
		return s[strings.LastIndex(s, "_")+1:]
	}
	return string(e) + "值未定义"
}

// 银期转帐方向类型
type TThostFtdcFBTTransferDirectionType byte

const THOST_FTDC_FBTTD_FromBankToFuture TThostFtdcFBTTransferDirectionType = '1' // 入金，银行转期货

const THOST_FTDC_FBTTD_FromFutureToBank TThostFtdcFBTTransferDirectionType = '2' // 出金，期货转银行

var mpTThostFtdcFBTTransferDirectionType = map[TThostFtdcFBTTransferDirectionType]string{'1': "THOST_FTDC_FBTTD_FromBankToFuture", '2': "THOST_FTDC_FBTTD_FromFutureToBank"}

func (e TThostFtdcFBTTransferDirectionType) String() string {
	if s, ok := mpTThostFtdcFBTTransferDirectionType[e]; ok {
		return s[strings.LastIndex(s, "_")+1:]
	}
	return string(e) + "值未定义"
}

// 开销户类别类型
type TThostFtdcOpenOrDestroyType byte

const THOST_FTDC_OOD_Open TThostFtdcOpenOrDestroyType = '1' // 开户

const THOST_FTDC_OOD_Destroy TThostFtdcOpenOrDestroyType = '0' // 销户

var mpTThostFtdcOpenOrDestroyType = map[TThostFtdcOpenOrDestroyType]string{'1': "THOST_FTDC_OOD_Open", '0': "THOST_FTDC_OOD_Destroy"}

func (e TThostFtdcOpenOrDestroyType) String() string {
	if s, ok := mpTThostFtdcOpenOrDestroyType[e]; ok {
		return s[strings.LastIndex(s, "_")+1:]
	}
	return string(e) + "值未定义"
}

// 有效标志类型
type TThostFtdcAvailabilityFlagType byte

const THOST_FTDC_AVAF_Invalid TThostFtdcAvailabilityFlagType = '0' // 未确认

const THOST_FTDC_AVAF_Valid TThostFtdcAvailabilityFlagType = '1' // 有效

const THOST_FTDC_AVAF_Repeal TThostFtdcAvailabilityFlagType = '2' // 冲正

var mpTThostFtdcAvailabilityFlagType = map[TThostFtdcAvailabilityFlagType]string{'0': "THOST_FTDC_AVAF_Invalid", '1': "THOST_FTDC_AVAF_Valid", '2': "THOST_FTDC_AVAF_Repeal"}

func (e TThostFtdcAvailabilityFlagType) String() string {
	if s, ok := mpTThostFtdcAvailabilityFlagType[e]; ok {
		return s[strings.LastIndex(s, "_")+1:]
	}
	return string(e) + "值未定义"
}

// 机构类型类型
type TThostFtdcOrganTypeType byte

const THOST_FTDC_OT_Bank TThostFtdcOrganTypeType = '1' // 银行代理

const THOST_FTDC_OT_Future TThostFtdcOrganTypeType = '2' // 交易前置

const THOST_FTDC_OT_PlateForm TThostFtdcOrganTypeType = '9' // 银期转帐平台管理

var mpTThostFtdcOrganTypeType = map[TThostFtdcOrganTypeType]string{'1': "THOST_FTDC_OT_Bank", '2': "THOST_FTDC_OT_Future", '9': "THOST_FTDC_OT_PlateForm"}

func (e TThostFtdcOrganTypeType) String() string {
	if s, ok := mpTThostFtdcOrganTypeType[e]; ok {
		return s[strings.LastIndex(s, "_")+1:]
	}
	return string(e) + "值未定义"
}

// 机构级别类型
type TThostFtdcOrganLevelType byte

const THOST_FTDC_OL_HeadQuarters TThostFtdcOrganLevelType = '1' // 银行总行或期商总部

const THOST_FTDC_OL_Branch TThostFtdcOrganLevelType = '2' // 银行分中心或期货公司营业部

var mpTThostFtdcOrganLevelType = map[TThostFtdcOrganLevelType]string{'1': "THOST_FTDC_OL_HeadQuarters", '2': "THOST_FTDC_OL_Branch"}

func (e TThostFtdcOrganLevelType) String() string {
	if s, ok := mpTThostFtdcOrganLevelType[e]; ok {
		return s[strings.LastIndex(s, "_")+1:]
	}
	return string(e) + "值未定义"
}

// 协议类型类型
type TThostFtdcProtocalIDType byte

const THOST_FTDC_PID_FutureProtocal TThostFtdcProtocalIDType = '0' // 期商协议

const THOST_FTDC_PID_ICBCProtocal TThostFtdcProtocalIDType = '1' // 工行协议

const THOST_FTDC_PID_ABCProtocal TThostFtdcProtocalIDType = '2' // 农行协议

const THOST_FTDC_PID_CBCProtocal TThostFtdcProtocalIDType = '3' // 中国银行协议

const THOST_FTDC_PID_CCBProtocal TThostFtdcProtocalIDType = '4' // 建行协议

const THOST_FTDC_PID_BOCOMProtocal TThostFtdcProtocalIDType = '5' // 交行协议

const THOST_FTDC_PID_FBTPlateFormProtocal TThostFtdcProtocalIDType = 'X' // 银期转帐平台协议

var mpTThostFtdcProtocalIDType = map[TThostFtdcProtocalIDType]string{'0': "THOST_FTDC_PID_FutureProtocal", '1': "THOST_FTDC_PID_ICBCProtocal", '2': "THOST_FTDC_PID_ABCProtocal", '3': "THOST_FTDC_PID_CBCProtocal", '4': "THOST_FTDC_PID_CCBProtocal", '5': "THOST_FTDC_PID_BOCOMProtocal", 'X': "THOST_FTDC_PID_FBTPlateFormProtocal"}

func (e TThostFtdcProtocalIDType) String() string {
	if s, ok := mpTThostFtdcProtocalIDType[e]; ok {
		return s[strings.LastIndex(s, "_")+1:]
	}
	return string(e) + "值未定义"
}

// 套接字连接方式类型
type TThostFtdcConnectModeType byte

const THOST_FTDC_CM_ShortConnect TThostFtdcConnectModeType = '0' // 短连接

const THOST_FTDC_CM_LongConnect TThostFtdcConnectModeType = '1' // 长连接

var mpTThostFtdcConnectModeType = map[TThostFtdcConnectModeType]string{'0': "THOST_FTDC_CM_ShortConnect", '1': "THOST_FTDC_CM_LongConnect"}

func (e TThostFtdcConnectModeType) String() string {
	if s, ok := mpTThostFtdcConnectModeType[e]; ok {
		return s[strings.LastIndex(s, "_")+1:]
	}
	return string(e) + "值未定义"
}

// 套接字通信方式类型
type TThostFtdcSyncModeType byte

const THOST_FTDC_SRM_ASync TThostFtdcSyncModeType = '0' // 异步

const THOST_FTDC_SRM_Sync TThostFtdcSyncModeType = '1' // 同步

var mpTThostFtdcSyncModeType = map[TThostFtdcSyncModeType]string{'0': "THOST_FTDC_SRM_ASync", '1': "THOST_FTDC_SRM_Sync"}

func (e TThostFtdcSyncModeType) String() string {
	if s, ok := mpTThostFtdcSyncModeType[e]; ok {
		return s[strings.LastIndex(s, "_")+1:]
	}
	return string(e) + "值未定义"
}

// 银行帐号类型类型
type TThostFtdcBankAccTypeType byte

const THOST_FTDC_BAT_BankBook TThostFtdcBankAccTypeType = '1' // 银行存折

const THOST_FTDC_BAT_SavingCard TThostFtdcBankAccTypeType = '2' // 储蓄卡

const THOST_FTDC_BAT_CreditCard TThostFtdcBankAccTypeType = '3' // 信用卡

var mpTThostFtdcBankAccTypeType = map[TThostFtdcBankAccTypeType]string{'1': "THOST_FTDC_BAT_BankBook", '2': "THOST_FTDC_BAT_SavingCard", '3': "THOST_FTDC_BAT_CreditCard"}

func (e TThostFtdcBankAccTypeType) String() string {
	if s, ok := mpTThostFtdcBankAccTypeType[e]; ok {
		return s[strings.LastIndex(s, "_")+1:]
	}
	return string(e) + "值未定义"
}

// 期货公司帐号类型类型
type TThostFtdcFutureAccTypeType byte

const THOST_FTDC_FAT_BankBook TThostFtdcFutureAccTypeType = '1' // 银行存折

const THOST_FTDC_FAT_SavingCard TThostFtdcFutureAccTypeType = '2' // 储蓄卡

const THOST_FTDC_FAT_CreditCard TThostFtdcFutureAccTypeType = '3' // 信用卡

var mpTThostFtdcFutureAccTypeType = map[TThostFtdcFutureAccTypeType]string{'1': "THOST_FTDC_FAT_BankBook", '2': "THOST_FTDC_FAT_SavingCard", '3': "THOST_FTDC_FAT_CreditCard"}

func (e TThostFtdcFutureAccTypeType) String() string {
	if s, ok := mpTThostFtdcFutureAccTypeType[e]; ok {
		return s[strings.LastIndex(s, "_")+1:]
	}
	return string(e) + "值未定义"
}

// 接入机构状态类型
type TThostFtdcOrganStatusType byte

const THOST_FTDC_OS_Ready TThostFtdcOrganStatusType = '0' // 启用

const THOST_FTDC_OS_CheckIn TThostFtdcOrganStatusType = '1' // 签到

const THOST_FTDC_OS_CheckOut TThostFtdcOrganStatusType = '2' // 签退

const THOST_FTDC_OS_CheckFileArrived TThostFtdcOrganStatusType = '3' // 对帐文件到达

const THOST_FTDC_OS_CheckDetail TThostFtdcOrganStatusType = '4' // 对帐

const THOST_FTDC_OS_DayEndClean TThostFtdcOrganStatusType = '5' // 日终清理

const THOST_FTDC_OS_Invalid TThostFtdcOrganStatusType = '9' // 注销

var mpTThostFtdcOrganStatusType = map[TThostFtdcOrganStatusType]string{'0': "THOST_FTDC_OS_Ready", '1': "THOST_FTDC_OS_CheckIn", '2': "THOST_FTDC_OS_CheckOut", '3': "THOST_FTDC_OS_CheckFileArrived", '4': "THOST_FTDC_OS_CheckDetail", '5': "THOST_FTDC_OS_DayEndClean", '9': "THOST_FTDC_OS_Invalid"}

func (e TThostFtdcOrganStatusType) String() string {
	if s, ok := mpTThostFtdcOrganStatusType[e]; ok {
		return s[strings.LastIndex(s, "_")+1:]
	}
	return string(e) + "值未定义"
}

// 建行收费模式类型
type TThostFtdcCCBFeeModeType byte

const THOST_FTDC_CCBFM_ByAmount TThostFtdcCCBFeeModeType = '1' // 按金额扣收

const THOST_FTDC_CCBFM_ByMonth TThostFtdcCCBFeeModeType = '2' // 按月扣收

var mpTThostFtdcCCBFeeModeType = map[TThostFtdcCCBFeeModeType]string{'1': "THOST_FTDC_CCBFM_ByAmount", '2': "THOST_FTDC_CCBFM_ByMonth"}

func (e TThostFtdcCCBFeeModeType) String() string {
	if s, ok := mpTThostFtdcCCBFeeModeType[e]; ok {
		return s[strings.LastIndex(s, "_")+1:]
	}
	return string(e) + "值未定义"
}

// 通讯API类型类型
type TThostFtdcCommApiTypeType byte

const THOST_FTDC_CAPIT_Client TThostFtdcCommApiTypeType = '1' // 客户端

const THOST_FTDC_CAPIT_Server TThostFtdcCommApiTypeType = '2' // 服务端

const THOST_FTDC_CAPIT_UserApi TThostFtdcCommApiTypeType = '3' // 交易系统的UserApi

var mpTThostFtdcCommApiTypeType = map[TThostFtdcCommApiTypeType]string{'1': "THOST_FTDC_CAPIT_Client", '2': "THOST_FTDC_CAPIT_Server", '3': "THOST_FTDC_CAPIT_UserApi"}

func (e TThostFtdcCommApiTypeType) String() string {
	if s, ok := mpTThostFtdcCommApiTypeType[e]; ok {
		return s[strings.LastIndex(s, "_")+1:]
	}
	return string(e) + "值未定义"
}

// 服务编号类型
type TThostFtdcServiceIDType int32

// 服务线路编号类型
type TThostFtdcServiceLineNoType int32

// 服务名类型
type TThostFtdcServiceNameType Byte61

// 连接状态类型
type TThostFtdcLinkStatusType byte

const THOST_FTDC_LS_Connected TThostFtdcLinkStatusType = '1' // 已经连接

const THOST_FTDC_LS_Disconnected TThostFtdcLinkStatusType = '2' // 没有连接

var mpTThostFtdcLinkStatusType = map[TThostFtdcLinkStatusType]string{'1': "THOST_FTDC_LS_Connected", '2': "THOST_FTDC_LS_Disconnected"}

func (e TThostFtdcLinkStatusType) String() string {
	if s, ok := mpTThostFtdcLinkStatusType[e]; ok {
		return s[strings.LastIndex(s, "_")+1:]
	}
	return string(e) + "值未定义"
}

// 通讯API指针类型
type TThostFtdcCommApiPointerType int32

// 密码核对标志类型
type TThostFtdcPwdFlagType byte

const THOST_FTDC_BPWDF_NoCheck TThostFtdcPwdFlagType = '0' // 不核对

const THOST_FTDC_BPWDF_BlankCheck TThostFtdcPwdFlagType = '1' // 明文核对

const THOST_FTDC_BPWDF_EncryptCheck TThostFtdcPwdFlagType = '2' // 密文核对

var mpTThostFtdcPwdFlagType = map[TThostFtdcPwdFlagType]string{'0': "THOST_FTDC_BPWDF_NoCheck", '1': "THOST_FTDC_BPWDF_BlankCheck", '2': "THOST_FTDC_BPWDF_EncryptCheck"}

func (e TThostFtdcPwdFlagType) String() string {
	if s, ok := mpTThostFtdcPwdFlagType[e]; ok {
		return s[strings.LastIndex(s, "_")+1:]
	}
	return string(e) + "值未定义"
}

// 期货帐号类型类型
type TThostFtdcSecuAccTypeType byte

const THOST_FTDC_SAT_AccountID TThostFtdcSecuAccTypeType = '1' // 资金帐号

const THOST_FTDC_SAT_CardID TThostFtdcSecuAccTypeType = '2' // 资金卡号

const THOST_FTDC_SAT_SHStockholderID TThostFtdcSecuAccTypeType = '3' // 上海股东帐号

const THOST_FTDC_SAT_SZStockholderID TThostFtdcSecuAccTypeType = '4' // 深圳股东帐号

var mpTThostFtdcSecuAccTypeType = map[TThostFtdcSecuAccTypeType]string{'1': "THOST_FTDC_SAT_AccountID", '2': "THOST_FTDC_SAT_CardID", '3': "THOST_FTDC_SAT_SHStockholderID", '4': "THOST_FTDC_SAT_SZStockholderID"}

func (e TThostFtdcSecuAccTypeType) String() string {
	if s, ok := mpTThostFtdcSecuAccTypeType[e]; ok {
		return s[strings.LastIndex(s, "_")+1:]
	}
	return string(e) + "值未定义"
}

// 转账交易状态类型
type TThostFtdcTransferStatusType byte

const THOST_FTDC_TRFS_Normal TThostFtdcTransferStatusType = '0' // 正常

const THOST_FTDC_TRFS_Repealed TThostFtdcTransferStatusType = '1' // 被冲正

var mpTThostFtdcTransferStatusType = map[TThostFtdcTransferStatusType]string{'0': "THOST_FTDC_TRFS_Normal", '1': "THOST_FTDC_TRFS_Repealed"}

func (e TThostFtdcTransferStatusType) String() string {
	if s, ok := mpTThostFtdcTransferStatusType[e]; ok {
		return s[strings.LastIndex(s, "_")+1:]
	}
	return string(e) + "值未定义"
}

// 发起方类型
type TThostFtdcSponsorTypeType byte

const THOST_FTDC_SPTYPE_Broker TThostFtdcSponsorTypeType = '0' // 期商

const THOST_FTDC_SPTYPE_Bank TThostFtdcSponsorTypeType = '1' // 银行

var mpTThostFtdcSponsorTypeType = map[TThostFtdcSponsorTypeType]string{'0': "THOST_FTDC_SPTYPE_Broker", '1': "THOST_FTDC_SPTYPE_Bank"}

func (e TThostFtdcSponsorTypeType) String() string {
	if s, ok := mpTThostFtdcSponsorTypeType[e]; ok {
		return s[strings.LastIndex(s, "_")+1:]
	}
	return string(e) + "值未定义"
}

// 请求响应类别类型
type TThostFtdcReqRspTypeType byte

const THOST_FTDC_REQRSP_Request TThostFtdcReqRspTypeType = '0' // 请求

const THOST_FTDC_REQRSP_Response TThostFtdcReqRspTypeType = '1' // 响应

var mpTThostFtdcReqRspTypeType = map[TThostFtdcReqRspTypeType]string{'0': "THOST_FTDC_REQRSP_Request", '1': "THOST_FTDC_REQRSP_Response"}

func (e TThostFtdcReqRspTypeType) String() string {
	if s, ok := mpTThostFtdcReqRspTypeType[e]; ok {
		return s[strings.LastIndex(s, "_")+1:]
	}
	return string(e) + "值未定义"
}

// 银期转帐用户事件类型类型
type TThostFtdcFBTUserEventTypeType byte

const THOST_FTDC_FBTUET_SignIn TThostFtdcFBTUserEventTypeType = '0' // 签到

const THOST_FTDC_FBTUET_FromBankToFuture TThostFtdcFBTUserEventTypeType = '1' // 银行转期货

const THOST_FTDC_FBTUET_FromFutureToBank TThostFtdcFBTUserEventTypeType = '2' // 期货转银行

const THOST_FTDC_FBTUET_OpenAccount TThostFtdcFBTUserEventTypeType = '3' // 开户

const THOST_FTDC_FBTUET_CancelAccount TThostFtdcFBTUserEventTypeType = '4' // 销户

const THOST_FTDC_FBTUET_ChangeAccount TThostFtdcFBTUserEventTypeType = '5' // 变更银行账户

const THOST_FTDC_FBTUET_RepealFromBankToFuture TThostFtdcFBTUserEventTypeType = '6' // 冲正银行转期货

const THOST_FTDC_FBTUET_RepealFromFutureToBank TThostFtdcFBTUserEventTypeType = '7' // 冲正期货转银行

const THOST_FTDC_FBTUET_QueryBankAccount TThostFtdcFBTUserEventTypeType = '8' // 查询银行账户

const THOST_FTDC_FBTUET_QueryFutureAccount TThostFtdcFBTUserEventTypeType = '9' // 查询期货账户

const THOST_FTDC_FBTUET_SignOut TThostFtdcFBTUserEventTypeType = 'A' // 签退

const THOST_FTDC_FBTUET_SyncKey TThostFtdcFBTUserEventTypeType = 'B' // 密钥同步

const THOST_FTDC_FBTUET_ReserveOpenAccount TThostFtdcFBTUserEventTypeType = 'C' // 预约开户

const THOST_FTDC_FBTUET_CancelReserveOpenAccount TThostFtdcFBTUserEventTypeType = 'D' // 撤销预约开户

const THOST_FTDC_FBTUET_ReserveOpenAccountConfirm TThostFtdcFBTUserEventTypeType = 'E' // 预约开户确认

const THOST_FTDC_FBTUET_Other TThostFtdcFBTUserEventTypeType = 'Z' // 其他

var mpTThostFtdcFBTUserEventTypeType = map[TThostFtdcFBTUserEventTypeType]string{'0': "THOST_FTDC_FBTUET_SignIn", '1': "THOST_FTDC_FBTUET_FromBankToFuture", '2': "THOST_FTDC_FBTUET_FromFutureToBank", '3': "THOST_FTDC_FBTUET_OpenAccount", '4': "THOST_FTDC_FBTUET_CancelAccount", '5': "THOST_FTDC_FBTUET_ChangeAccount", '6': "THOST_FTDC_FBTUET_RepealFromBankToFuture", '7': "THOST_FTDC_FBTUET_RepealFromFutureToBank", '8': "THOST_FTDC_FBTUET_QueryBankAccount", '9': "THOST_FTDC_FBTUET_QueryFutureAccount", 'A': "THOST_FTDC_FBTUET_SignOut", 'B': "THOST_FTDC_FBTUET_SyncKey", 'C': "THOST_FTDC_FBTUET_ReserveOpenAccount", 'D': "THOST_FTDC_FBTUET_CancelReserveOpenAccount", 'E': "THOST_FTDC_FBTUET_ReserveOpenAccountConfirm", 'Z': "THOST_FTDC_FBTUET_Other"}

func (e TThostFtdcFBTUserEventTypeType) String() string {
	if s, ok := mpTThostFtdcFBTUserEventTypeType[e]; ok {
		return s[strings.LastIndex(s, "_")+1:]
	}
	return string(e) + "值未定义"
}

// 银行自己的编码类型
type TThostFtdcBankIDByBankType Byte21

// 银行操作员号类型
type TThostFtdcBankOperNoType Byte4

// 银行客户号类型
type TThostFtdcBankCustNoType Byte21

// 递增的序列号类型
type TThostFtdcDBOPSeqNoType int32

// FBT表名类型
type TThostFtdcTableNameType Byte61

// FBT表操作主键名类型
type TThostFtdcPKNameType Byte201

// FBT表操作主键值类型
type TThostFtdcPKValueType Byte501

// 记录操作类型类型
type TThostFtdcDBOperationType byte

const THOST_FTDC_DBOP_Insert TThostFtdcDBOperationType = '0' // 插入

const THOST_FTDC_DBOP_Update TThostFtdcDBOperationType = '1' // 更新

const THOST_FTDC_DBOP_Delete TThostFtdcDBOperationType = '2' // 删除

var mpTThostFtdcDBOperationType = map[TThostFtdcDBOperationType]string{'0': "THOST_FTDC_DBOP_Insert", '1': "THOST_FTDC_DBOP_Update", '2': "THOST_FTDC_DBOP_Delete"}

func (e TThostFtdcDBOperationType) String() string {
	if s, ok := mpTThostFtdcDBOperationType[e]; ok {
		return s[strings.LastIndex(s, "_")+1:]
	}
	return string(e) + "值未定义"
}

// 同步标记类型
type TThostFtdcSyncFlagType byte

const THOST_FTDC_SYNF_Yes TThostFtdcSyncFlagType = '0' // 已同步

const THOST_FTDC_SYNF_No TThostFtdcSyncFlagType = '1' // 未同步

var mpTThostFtdcSyncFlagType = map[TThostFtdcSyncFlagType]string{'0': "THOST_FTDC_SYNF_Yes", '1': "THOST_FTDC_SYNF_No"}

func (e TThostFtdcSyncFlagType) String() string {
	if s, ok := mpTThostFtdcSyncFlagType[e]; ok {
		return s[strings.LastIndex(s, "_")+1:]
	}
	return string(e) + "值未定义"
}

// 同步目标编号类型
type TThostFtdcTargetIDType Byte4

// 同步类型类型
type TThostFtdcSyncTypeType byte

const THOST_FTDC_SYNT_OneOffSync TThostFtdcSyncTypeType = '0' // 一次同步

const THOST_FTDC_SYNT_TimerSync TThostFtdcSyncTypeType = '1' // 定时同步

const THOST_FTDC_SYNT_TimerFullSync TThostFtdcSyncTypeType = '2' // 定时完全同步

var mpTThostFtdcSyncTypeType = map[TThostFtdcSyncTypeType]string{'0': "THOST_FTDC_SYNT_OneOffSync", '1': "THOST_FTDC_SYNT_TimerSync", '2': "THOST_FTDC_SYNT_TimerFullSync"}

func (e TThostFtdcSyncTypeType) String() string {
	if s, ok := mpTThostFtdcSyncTypeType[e]; ok {
		return s[strings.LastIndex(s, "_")+1:]
	}
	return string(e) + "值未定义"
}

// 各种换汇时间类型
type TThostFtdcFBETimeType Byte7

// 换汇银行行号类型
type TThostFtdcFBEBankNoType Byte13

// 换汇凭证号类型
type TThostFtdcFBECertNoType Byte13

// 换汇方向类型
type TThostFtdcExDirectionType byte

const THOST_FTDC_FBEDIR_Settlement TThostFtdcExDirectionType = '0' // 结汇

const THOST_FTDC_FBEDIR_Sale TThostFtdcExDirectionType = '1' // 售汇

var mpTThostFtdcExDirectionType = map[TThostFtdcExDirectionType]string{'0': "THOST_FTDC_FBEDIR_Settlement", '1': "THOST_FTDC_FBEDIR_Sale"}

func (e TThostFtdcExDirectionType) String() string {
	if s, ok := mpTThostFtdcExDirectionType[e]; ok {
		return s[strings.LastIndex(s, "_")+1:]
	}
	return string(e) + "值未定义"
}

// 换汇银行账户类型
type TThostFtdcFBEBankAccountType Byte33

// 换汇银行账户名类型
type TThostFtdcFBEBankAccountNameType Byte61

// 各种换汇金额类型
type TThostFtdcFBEAmtType Double

// 换汇业务类型类型
type TThostFtdcFBEBusinessTypeType Byte3

// 换汇附言类型
type TThostFtdcFBEPostScriptType Byte61

// 换汇备注类型
type TThostFtdcFBERemarkType Byte71

// 换汇汇率类型
type TThostFtdcExRateType Double

// 换汇成功标志类型
type TThostFtdcFBEResultFlagType byte

const THOST_FTDC_FBERES_Success TThostFtdcFBEResultFlagType = '0' // 成功

const THOST_FTDC_FBERES_InsufficientBalance TThostFtdcFBEResultFlagType = '1' // 账户余额不足

const THOST_FTDC_FBERES_UnknownTrading TThostFtdcFBEResultFlagType = '8' // 交易结果未知

const THOST_FTDC_FBERES_Fail TThostFtdcFBEResultFlagType = 'x' // 失败

var mpTThostFtdcFBEResultFlagType = map[TThostFtdcFBEResultFlagType]string{'0': "THOST_FTDC_FBERES_Success", '1': "THOST_FTDC_FBERES_InsufficientBalance", '8': "THOST_FTDC_FBERES_UnknownTrading", 'x': "THOST_FTDC_FBERES_Fail"}

func (e TThostFtdcFBEResultFlagType) String() string {
	if s, ok := mpTThostFtdcFBEResultFlagType[e]; ok {
		return s[strings.LastIndex(s, "_")+1:]
	}
	return string(e) + "值未定义"
}

// 换汇返回信息类型
type TThostFtdcFBERtnMsgType Byte61

// 换汇扩展信息类型
type TThostFtdcFBEExtendMsgType Byte61

// 换汇记账流水号类型
type TThostFtdcFBEBusinessSerialType Byte31

// 换汇流水号类型
type TThostFtdcFBESystemSerialType Byte21

// 换汇交易总笔数类型
type TThostFtdcFBETotalExCntType int32

// 换汇交易状态类型
type TThostFtdcFBEExchStatusType byte

const THOST_FTDC_FBEES_Normal TThostFtdcFBEExchStatusType = '0' // 正常

const THOST_FTDC_FBEES_ReExchange TThostFtdcFBEExchStatusType = '1' // 交易重发

var mpTThostFtdcFBEExchStatusType = map[TThostFtdcFBEExchStatusType]string{'0': "THOST_FTDC_FBEES_Normal", '1': "THOST_FTDC_FBEES_ReExchange"}

func (e TThostFtdcFBEExchStatusType) String() string {
	if s, ok := mpTThostFtdcFBEExchStatusType[e]; ok {
		return s[strings.LastIndex(s, "_")+1:]
	}
	return string(e) + "值未定义"
}

// 换汇文件标志类型
type TThostFtdcFBEFileFlagType byte

const THOST_FTDC_FBEFG_DataPackage TThostFtdcFBEFileFlagType = '0' // 数据包

const THOST_FTDC_FBEFG_File TThostFtdcFBEFileFlagType = '1' // 文件

var mpTThostFtdcFBEFileFlagType = map[TThostFtdcFBEFileFlagType]string{'0': "THOST_FTDC_FBEFG_DataPackage", '1': "THOST_FTDC_FBEFG_File"}

func (e TThostFtdcFBEFileFlagType) String() string {
	if s, ok := mpTThostFtdcFBEFileFlagType[e]; ok {
		return s[strings.LastIndex(s, "_")+1:]
	}
	return string(e) + "值未定义"
}

// 换汇已交易标志类型
type TThostFtdcFBEAlreadyTradeType byte

const THOST_FTDC_FBEAT_NotTrade TThostFtdcFBEAlreadyTradeType = '0' // 未交易

const THOST_FTDC_FBEAT_Trade TThostFtdcFBEAlreadyTradeType = '1' // 已交易

var mpTThostFtdcFBEAlreadyTradeType = map[TThostFtdcFBEAlreadyTradeType]string{'0': "THOST_FTDC_FBEAT_NotTrade", '1': "THOST_FTDC_FBEAT_Trade"}

func (e TThostFtdcFBEAlreadyTradeType) String() string {
	if s, ok := mpTThostFtdcFBEAlreadyTradeType[e]; ok {
		return s[strings.LastIndex(s, "_")+1:]
	}
	return string(e) + "值未定义"
}

// 换汇账户开户行类型
type TThostFtdcFBEOpenBankType Byte61

// 银期换汇用户事件类型类型
type TThostFtdcFBEUserEventTypeType byte

const THOST_FTDC_FBEUET_SignIn TThostFtdcFBEUserEventTypeType = '0' // 签到

const THOST_FTDC_FBEUET_Exchange TThostFtdcFBEUserEventTypeType = '1' // 换汇

const THOST_FTDC_FBEUET_ReExchange TThostFtdcFBEUserEventTypeType = '2' // 换汇重发

const THOST_FTDC_FBEUET_QueryBankAccount TThostFtdcFBEUserEventTypeType = '3' // 银行账户查询

const THOST_FTDC_FBEUET_QueryExchDetial TThostFtdcFBEUserEventTypeType = '4' // 换汇明细查询

const THOST_FTDC_FBEUET_QueryExchSummary TThostFtdcFBEUserEventTypeType = '5' // 换汇汇总查询

const THOST_FTDC_FBEUET_QueryExchRate TThostFtdcFBEUserEventTypeType = '6' // 换汇汇率查询

const THOST_FTDC_FBEUET_CheckBankAccount TThostFtdcFBEUserEventTypeType = '7' // 对账文件通知

const THOST_FTDC_FBEUET_SignOut TThostFtdcFBEUserEventTypeType = '8' // 签退

const THOST_FTDC_FBEUET_Other TThostFtdcFBEUserEventTypeType = 'Z' // 其他

var mpTThostFtdcFBEUserEventTypeType = map[TThostFtdcFBEUserEventTypeType]string{'0': "THOST_FTDC_FBEUET_SignIn", '1': "THOST_FTDC_FBEUET_Exchange", '2': "THOST_FTDC_FBEUET_ReExchange", '3': "THOST_FTDC_FBEUET_QueryBankAccount", '4': "THOST_FTDC_FBEUET_QueryExchDetial", '5': "THOST_FTDC_FBEUET_QueryExchSummary", '6': "THOST_FTDC_FBEUET_QueryExchRate", '7': "THOST_FTDC_FBEUET_CheckBankAccount", '8': "THOST_FTDC_FBEUET_SignOut", 'Z': "THOST_FTDC_FBEUET_Other"}

func (e TThostFtdcFBEUserEventTypeType) String() string {
	if s, ok := mpTThostFtdcFBEUserEventTypeType[e]; ok {
		return s[strings.LastIndex(s, "_")+1:]
	}
	return string(e) + "值未定义"
}

// 换汇相关文件名类型
type TThostFtdcFBEFileNameType Byte21

// 换汇批次号类型
type TThostFtdcFBEBatchSerialType Byte21

// 换汇发送标志类型
type TThostFtdcFBEReqFlagType byte

const THOST_FTDC_FBERF_UnProcessed TThostFtdcFBEReqFlagType = '0' // 未处理

const THOST_FTDC_FBERF_WaitSend TThostFtdcFBEReqFlagType = '1' // 等待发送

const THOST_FTDC_FBERF_SendSuccess TThostFtdcFBEReqFlagType = '2' // 发送成功

const THOST_FTDC_FBERF_SendFailed TThostFtdcFBEReqFlagType = '3' // 发送失败

const THOST_FTDC_FBERF_WaitReSend TThostFtdcFBEReqFlagType = '4' // 等待重发

var mpTThostFtdcFBEReqFlagType = map[TThostFtdcFBEReqFlagType]string{'0': "THOST_FTDC_FBERF_UnProcessed", '1': "THOST_FTDC_FBERF_WaitSend", '2': "THOST_FTDC_FBERF_SendSuccess", '3': "THOST_FTDC_FBERF_SendFailed", '4': "THOST_FTDC_FBERF_WaitReSend"}

func (e TThostFtdcFBEReqFlagType) String() string {
	if s, ok := mpTThostFtdcFBEReqFlagType[e]; ok {
		return s[strings.LastIndex(s, "_")+1:]
	}
	return string(e) + "值未定义"
}

// 风险通知类型类型
type TThostFtdcNotifyClassType byte

const THOST_FTDC_NC_NOERROR TThostFtdcNotifyClassType = '0' // 正常

const THOST_FTDC_NC_Warn TThostFtdcNotifyClassType = '1' // 警示

const THOST_FTDC_NC_Call TThostFtdcNotifyClassType = '2' // 追保

const THOST_FTDC_NC_Force TThostFtdcNotifyClassType = '3' // 强平

const THOST_FTDC_NC_CHUANCANG TThostFtdcNotifyClassType = '4' // 穿仓

const THOST_FTDC_NC_Exception TThostFtdcNotifyClassType = '5' // 异常

var mpTThostFtdcNotifyClassType = map[TThostFtdcNotifyClassType]string{'0': "THOST_FTDC_NC_NOERROR", '1': "THOST_FTDC_NC_Warn", '2': "THOST_FTDC_NC_Call", '3': "THOST_FTDC_NC_Force", '4': "THOST_FTDC_NC_CHUANCANG", '5': "THOST_FTDC_NC_Exception"}

func (e TThostFtdcNotifyClassType) String() string {
	if s, ok := mpTThostFtdcNotifyClassType[e]; ok {
		return s[strings.LastIndex(s, "_")+1:]
	}
	return string(e) + "值未定义"
}

// 客户风险通知消息类型
type TThostFtdcRiskNofityInfoType Byte257

// 强平场景编号类型
type TThostFtdcForceCloseSceneIdType Byte24

// 强平单类型类型
type TThostFtdcForceCloseTypeType byte

const THOST_FTDC_FCT_Manual TThostFtdcForceCloseTypeType = '0' // 手工强平

const THOST_FTDC_FCT_Single TThostFtdcForceCloseTypeType = '1' // 单一投资者辅助强平

const THOST_FTDC_FCT_Group TThostFtdcForceCloseTypeType = '2' // 批量投资者辅助强平

var mpTThostFtdcForceCloseTypeType = map[TThostFtdcForceCloseTypeType]string{'0': "THOST_FTDC_FCT_Manual", '1': "THOST_FTDC_FCT_Single", '2': "THOST_FTDC_FCT_Group"}

func (e TThostFtdcForceCloseTypeType) String() string {
	if s, ok := mpTThostFtdcForceCloseTypeType[e]; ok {
		return s[strings.LastIndex(s, "_")+1:]
	}
	return string(e) + "值未定义"
}

// 多个产品代码,用+分隔,如cu+zn类型
type TThostFtdcInstrumentIDsType Byte101

// 风险通知途径类型
type TThostFtdcRiskNotifyMethodType byte

const THOST_FTDC_RNM_System TThostFtdcRiskNotifyMethodType = '0' // 系统通知

const THOST_FTDC_RNM_SMS TThostFtdcRiskNotifyMethodType = '1' // 短信通知

const THOST_FTDC_RNM_EMail TThostFtdcRiskNotifyMethodType = '2' // 邮件通知

const THOST_FTDC_RNM_Manual TThostFtdcRiskNotifyMethodType = '3' // 人工通知

var mpTThostFtdcRiskNotifyMethodType = map[TThostFtdcRiskNotifyMethodType]string{'0': "THOST_FTDC_RNM_System", '1': "THOST_FTDC_RNM_SMS", '2': "THOST_FTDC_RNM_EMail", '3': "THOST_FTDC_RNM_Manual"}

func (e TThostFtdcRiskNotifyMethodType) String() string {
	if s, ok := mpTThostFtdcRiskNotifyMethodType[e]; ok {
		return s[strings.LastIndex(s, "_")+1:]
	}
	return string(e) + "值未定义"
}

// 风险通知状态类型
type TThostFtdcRiskNotifyStatusType byte

const THOST_FTDC_RNS_NotGen TThostFtdcRiskNotifyStatusType = '0' // 未生成

const THOST_FTDC_RNS_Generated TThostFtdcRiskNotifyStatusType = '1' // 已生成未发送

const THOST_FTDC_RNS_SendError TThostFtdcRiskNotifyStatusType = '2' // 发送失败

const THOST_FTDC_RNS_SendOk TThostFtdcRiskNotifyStatusType = '3' // 已发送未接收

const THOST_FTDC_RNS_Received TThostFtdcRiskNotifyStatusType = '4' // 已接收未确认

const THOST_FTDC_RNS_Confirmed TThostFtdcRiskNotifyStatusType = '5' // 已确认

var mpTThostFtdcRiskNotifyStatusType = map[TThostFtdcRiskNotifyStatusType]string{'0': "THOST_FTDC_RNS_NotGen", '1': "THOST_FTDC_RNS_Generated", '2': "THOST_FTDC_RNS_SendError", '3': "THOST_FTDC_RNS_SendOk", '4': "THOST_FTDC_RNS_Received", '5': "THOST_FTDC_RNS_Confirmed"}

func (e TThostFtdcRiskNotifyStatusType) String() string {
	if s, ok := mpTThostFtdcRiskNotifyStatusType[e]; ok {
		return s[strings.LastIndex(s, "_")+1:]
	}
	return string(e) + "值未定义"
}

// 风控用户操作事件类型
type TThostFtdcRiskUserEventType byte

const THOST_FTDC_RUE_ExportData TThostFtdcRiskUserEventType = '0' // 导出数据

var mpTThostFtdcRiskUserEventType = map[TThostFtdcRiskUserEventType]string{'0': "THOST_FTDC_RUE_ExportData"}

func (e TThostFtdcRiskUserEventType) String() string {
	if s, ok := mpTThostFtdcRiskUserEventType[e]; ok {
		return s[strings.LastIndex(s, "_")+1:]
	}
	return string(e) + "值未定义"
}

// 参数代码类型
type TThostFtdcParamIDType int32

// 参数名类型
type TThostFtdcParamNameType Byte41

// 参数值类型
type TThostFtdcParamValueType Byte41

// 条件单索引条件类型
type TThostFtdcConditionalOrderSortTypeType byte

const THOST_FTDC_COST_LastPriceAsc TThostFtdcConditionalOrderSortTypeType = '0' // 使用最新价升序

const THOST_FTDC_COST_LastPriceDesc TThostFtdcConditionalOrderSortTypeType = '1' // 使用最新价降序

const THOST_FTDC_COST_AskPriceAsc TThostFtdcConditionalOrderSortTypeType = '2' // 使用卖价升序

const THOST_FTDC_COST_AskPriceDesc TThostFtdcConditionalOrderSortTypeType = '3' // 使用卖价降序

const THOST_FTDC_COST_BidPriceAsc TThostFtdcConditionalOrderSortTypeType = '4' // 使用买价升序

const THOST_FTDC_COST_BidPriceDesc TThostFtdcConditionalOrderSortTypeType = '5' // 使用买价降序

var mpTThostFtdcConditionalOrderSortTypeType = map[TThostFtdcConditionalOrderSortTypeType]string{'0': "THOST_FTDC_COST_LastPriceAsc", '1': "THOST_FTDC_COST_LastPriceDesc", '2': "THOST_FTDC_COST_AskPriceAsc", '3': "THOST_FTDC_COST_AskPriceDesc", '4': "THOST_FTDC_COST_BidPriceAsc", '5': "THOST_FTDC_COST_BidPriceDesc"}

func (e TThostFtdcConditionalOrderSortTypeType) String() string {
	if s, ok := mpTThostFtdcConditionalOrderSortTypeType[e]; ok {
		return s[strings.LastIndex(s, "_")+1:]
	}
	return string(e) + "值未定义"
}

// 报送状态类型
type TThostFtdcSendTypeType byte

const THOST_FTDC_UOAST_NoSend TThostFtdcSendTypeType = '0' // 未发送

const THOST_FTDC_UOAST_Sended TThostFtdcSendTypeType = '1' // 已发送

const THOST_FTDC_UOAST_Generated TThostFtdcSendTypeType = '2' // 已生成

const THOST_FTDC_UOAST_SendFail TThostFtdcSendTypeType = '3' // 报送失败

const THOST_FTDC_UOAST_Success TThostFtdcSendTypeType = '4' // 接收成功

const THOST_FTDC_UOAST_Fail TThostFtdcSendTypeType = '5' // 接收失败

const THOST_FTDC_UOAST_Cancel TThostFtdcSendTypeType = '6' // 取消报送

var mpTThostFtdcSendTypeType = map[TThostFtdcSendTypeType]string{'0': "THOST_FTDC_UOAST_NoSend", '1': "THOST_FTDC_UOAST_Sended", '2': "THOST_FTDC_UOAST_Generated", '3': "THOST_FTDC_UOAST_SendFail", '4': "THOST_FTDC_UOAST_Success", '5': "THOST_FTDC_UOAST_Fail", '6': "THOST_FTDC_UOAST_Cancel"}

func (e TThostFtdcSendTypeType) String() string {
	if s, ok := mpTThostFtdcSendTypeType[e]; ok {
		return s[strings.LastIndex(s, "_")+1:]
	}
	return string(e) + "值未定义"
}

// 交易编码状态类型
type TThostFtdcClientIDStatusType byte

const THOST_FTDC_UOACS_NoApply TThostFtdcClientIDStatusType = '1' // 未申请

const THOST_FTDC_UOACS_Submited TThostFtdcClientIDStatusType = '2' // 已提交申请

const THOST_FTDC_UOACS_Sended TThostFtdcClientIDStatusType = '3' // 已发送申请

const THOST_FTDC_UOACS_Success TThostFtdcClientIDStatusType = '4' // 完成

const THOST_FTDC_UOACS_Refuse TThostFtdcClientIDStatusType = '5' // 拒绝

const THOST_FTDC_UOACS_Cancel TThostFtdcClientIDStatusType = '6' // 已撤销编码

var mpTThostFtdcClientIDStatusType = map[TThostFtdcClientIDStatusType]string{'1': "THOST_FTDC_UOACS_NoApply", '2': "THOST_FTDC_UOACS_Submited", '3': "THOST_FTDC_UOACS_Sended", '4': "THOST_FTDC_UOACS_Success", '5': "THOST_FTDC_UOACS_Refuse", '6': "THOST_FTDC_UOACS_Cancel"}

func (e TThostFtdcClientIDStatusType) String() string {
	if s, ok := mpTThostFtdcClientIDStatusType[e]; ok {
		return s[strings.LastIndex(s, "_")+1:]
	}
	return string(e) + "值未定义"
}

// 行业编码类型
type TThostFtdcIndustryIDType Byte17

// 特有信息编号类型
type TThostFtdcQuestionIDType Byte5

// 特有信息说明类型
type TThostFtdcQuestionContentType Byte41

// 选项编号类型
type TThostFtdcOptionIDType Byte13

// 选项说明类型
type TThostFtdcOptionContentType Byte61

// 特有信息类型类型
type TThostFtdcQuestionTypeType byte

const THOST_FTDC_QT_Radio TThostFtdcQuestionTypeType = '1' // 单选

const THOST_FTDC_QT_Option TThostFtdcQuestionTypeType = '2' // 多选

const THOST_FTDC_QT_Blank TThostFtdcQuestionTypeType = '3' // 填空

var mpTThostFtdcQuestionTypeType = map[TThostFtdcQuestionTypeType]string{'1': "THOST_FTDC_QT_Radio", '2': "THOST_FTDC_QT_Option", '3': "THOST_FTDC_QT_Blank"}

func (e TThostFtdcQuestionTypeType) String() string {
	if s, ok := mpTThostFtdcQuestionTypeType[e]; ok {
		return s[strings.LastIndex(s, "_")+1:]
	}
	return string(e) + "值未定义"
}

// 业务流水号类型
type TThostFtdcProcessIDType Byte33

// 流水号类型
type TThostFtdcSeqNoType int32

// 流程状态类型
type TThostFtdcUOAProcessStatusType Byte3

// 流程功能类型类型
type TThostFtdcProcessTypeType Byte3

// 业务类型类型
type TThostFtdcBusinessTypeType byte

const THOST_FTDC_BT_Request TThostFtdcBusinessTypeType = '1' // 请求

const THOST_FTDC_BT_Response TThostFtdcBusinessTypeType = '2' // 应答

const THOST_FTDC_BT_Notice TThostFtdcBusinessTypeType = '3' // 通知

var mpTThostFtdcBusinessTypeType = map[TThostFtdcBusinessTypeType]string{'1': "THOST_FTDC_BT_Request", '2': "THOST_FTDC_BT_Response", '3': "THOST_FTDC_BT_Notice"}

func (e TThostFtdcBusinessTypeType) String() string {
	if s, ok := mpTThostFtdcBusinessTypeType[e]; ok {
		return s[strings.LastIndex(s, "_")+1:]
	}
	return string(e) + "值未定义"
}

// 监控中心返回码类型
type TThostFtdcCfmmcReturnCodeType byte

const THOST_FTDC_CRC_Success TThostFtdcCfmmcReturnCodeType = '0' // 成功

const THOST_FTDC_CRC_Working TThostFtdcCfmmcReturnCodeType = '1' // 该客户已经有流程在处理中

const THOST_FTDC_CRC_InfoFail TThostFtdcCfmmcReturnCodeType = '2' // 监控中客户资料检查失败

const THOST_FTDC_CRC_IDCardFail TThostFtdcCfmmcReturnCodeType = '3' // 监控中实名制检查失败

const THOST_FTDC_CRC_OtherFail TThostFtdcCfmmcReturnCodeType = '4' // 其他错误

var mpTThostFtdcCfmmcReturnCodeType = map[TThostFtdcCfmmcReturnCodeType]string{'0': "THOST_FTDC_CRC_Success", '1': "THOST_FTDC_CRC_Working", '2': "THOST_FTDC_CRC_InfoFail", '3': "THOST_FTDC_CRC_IDCardFail", '4': "THOST_FTDC_CRC_OtherFail"}

func (e TThostFtdcCfmmcReturnCodeType) String() string {
	if s, ok := mpTThostFtdcCfmmcReturnCodeType[e]; ok {
		return s[strings.LastIndex(s, "_")+1:]
	}
	return string(e) + "值未定义"
}

// 交易所返回码类型
type TThostFtdcExReturnCodeType int32

// 客户类型类型
type TThostFtdcClientTypeType byte

const THOST_FTDC_CfMMCCT_All TThostFtdcClientTypeType = '0' // 所有

const THOST_FTDC_CfMMCCT_Person TThostFtdcClientTypeType = '1' // 个人

const THOST_FTDC_CfMMCCT_Company TThostFtdcClientTypeType = '2' // 单位

const THOST_FTDC_CfMMCCT_Other TThostFtdcClientTypeType = '3' // 其他

const THOST_FTDC_CfMMCCT_SpecialOrgan TThostFtdcClientTypeType = '4' // 特殊法人

const THOST_FTDC_CfMMCCT_Asset TThostFtdcClientTypeType = '5' // 资管户

var mpTThostFtdcClientTypeType = map[TThostFtdcClientTypeType]string{'0': "THOST_FTDC_CfMMCCT_All", '1': "THOST_FTDC_CfMMCCT_Person", '2': "THOST_FTDC_CfMMCCT_Company", '3': "THOST_FTDC_CfMMCCT_Other", '4': "THOST_FTDC_CfMMCCT_SpecialOrgan", '5': "THOST_FTDC_CfMMCCT_Asset"}

func (e TThostFtdcClientTypeType) String() string {
	if s, ok := mpTThostFtdcClientTypeType[e]; ok {
		return s[strings.LastIndex(s, "_")+1:]
	}
	return string(e) + "值未定义"
}

// 交易所编号类型
type TThostFtdcExchangeIDTypeType byte

const THOST_FTDC_EIDT_SHFE TThostFtdcExchangeIDTypeType = 'S' // 上海期货交易所

const THOST_FTDC_EIDT_CZCE TThostFtdcExchangeIDTypeType = 'Z' // 郑州商品交易所

const THOST_FTDC_EIDT_DCE TThostFtdcExchangeIDTypeType = 'D' // 大连商品交易所

const THOST_FTDC_EIDT_CFFEX TThostFtdcExchangeIDTypeType = 'J' // 中国金融期货交易所

const THOST_FTDC_EIDT_INE TThostFtdcExchangeIDTypeType = 'N' // 上海国际能源交易中心股份有限公司

var mpTThostFtdcExchangeIDTypeType = map[TThostFtdcExchangeIDTypeType]string{'S': "THOST_FTDC_EIDT_SHFE", 'Z': "THOST_FTDC_EIDT_CZCE", 'D': "THOST_FTDC_EIDT_DCE", 'J': "THOST_FTDC_EIDT_CFFEX", 'N': "THOST_FTDC_EIDT_INE"}

func (e TThostFtdcExchangeIDTypeType) String() string {
	if s, ok := mpTThostFtdcExchangeIDTypeType[e]; ok {
		return s[strings.LastIndex(s, "_")+1:]
	}
	return string(e) + "值未定义"
}

// 交易编码类型类型
type TThostFtdcExClientIDTypeType byte

const THOST_FTDC_ECIDT_Hedge TThostFtdcExClientIDTypeType = '1' // 套保

const THOST_FTDC_ECIDT_Arbitrage TThostFtdcExClientIDTypeType = '2' // 套利

const THOST_FTDC_ECIDT_Speculation TThostFtdcExClientIDTypeType = '3' // 投机

var mpTThostFtdcExClientIDTypeType = map[TThostFtdcExClientIDTypeType]string{'1': "THOST_FTDC_ECIDT_Hedge", '2': "THOST_FTDC_ECIDT_Arbitrage", '3': "THOST_FTDC_ECIDT_Speculation"}

func (e TThostFtdcExClientIDTypeType) String() string {
	if s, ok := mpTThostFtdcExClientIDTypeType[e]; ok {
		return s[strings.LastIndex(s, "_")+1:]
	}
	return string(e) + "值未定义"
}

// 客户分类码类型
type TThostFtdcClientClassifyType Byte11

// 单位性质类型
type TThostFtdcUOAOrganTypeType Byte11

// 国家代码类型
type TThostFtdcUOACountryCodeType Byte11

// 区号类型
type TThostFtdcAreaCodeType Byte11

// 监控中心为客户分配的代码类型
type TThostFtdcFuturesIDType Byte21

// 日期类型
type TThostFtdcCffmcDateType Byte11

// 时间类型
type TThostFtdcCffmcTimeType Byte11

// 组织机构代码类型
type TThostFtdcNocIDType Byte21

// 更新状态类型
type TThostFtdcUpdateFlagType byte

const THOST_FTDC_UF_NoUpdate TThostFtdcUpdateFlagType = '0' // 未更新

const THOST_FTDC_UF_Success TThostFtdcUpdateFlagType = '1' // 更新全部信息成功

const THOST_FTDC_UF_Fail TThostFtdcUpdateFlagType = '2' // 更新全部信息失败

const THOST_FTDC_UF_TCSuccess TThostFtdcUpdateFlagType = '3' // 更新交易编码成功

const THOST_FTDC_UF_TCFail TThostFtdcUpdateFlagType = '4' // 更新交易编码失败

const THOST_FTDC_UF_Cancel TThostFtdcUpdateFlagType = '5' // 已丢弃

var mpTThostFtdcUpdateFlagType = map[TThostFtdcUpdateFlagType]string{'0': "THOST_FTDC_UF_NoUpdate", '1': "THOST_FTDC_UF_Success", '2': "THOST_FTDC_UF_Fail", '3': "THOST_FTDC_UF_TCSuccess", '4': "THOST_FTDC_UF_TCFail", '5': "THOST_FTDC_UF_Cancel"}

func (e TThostFtdcUpdateFlagType) String() string {
	if s, ok := mpTThostFtdcUpdateFlagType[e]; ok {
		return s[strings.LastIndex(s, "_")+1:]
	}
	return string(e) + "值未定义"
}

// 申请动作类型
type TThostFtdcApplyOperateIDType byte

const THOST_FTDC_AOID_OpenInvestor TThostFtdcApplyOperateIDType = '1' // 开户

const THOST_FTDC_AOID_ModifyIDCard TThostFtdcApplyOperateIDType = '2' // 修改身份信息

const THOST_FTDC_AOID_ModifyNoIDCard TThostFtdcApplyOperateIDType = '3' // 修改一般信息

const THOST_FTDC_AOID_ApplyTradingCode TThostFtdcApplyOperateIDType = '4' // 申请交易编码

const THOST_FTDC_AOID_CancelTradingCode TThostFtdcApplyOperateIDType = '5' // 撤销交易编码

const THOST_FTDC_AOID_CancelInvestor TThostFtdcApplyOperateIDType = '6' // 销户

const THOST_FTDC_AOID_FreezeAccount TThostFtdcApplyOperateIDType = '8' // 账户休眠

const THOST_FTDC_AOID_ActiveFreezeAccount TThostFtdcApplyOperateIDType = '9' // 激活休眠账户

var mpTThostFtdcApplyOperateIDType = map[TThostFtdcApplyOperateIDType]string{'1': "THOST_FTDC_AOID_OpenInvestor", '2': "THOST_FTDC_AOID_ModifyIDCard", '3': "THOST_FTDC_AOID_ModifyNoIDCard", '4': "THOST_FTDC_AOID_ApplyTradingCode", '5': "THOST_FTDC_AOID_CancelTradingCode", '6': "THOST_FTDC_AOID_CancelInvestor", '8': "THOST_FTDC_AOID_FreezeAccount", '9': "THOST_FTDC_AOID_ActiveFreezeAccount"}

func (e TThostFtdcApplyOperateIDType) String() string {
	if s, ok := mpTThostFtdcApplyOperateIDType[e]; ok {
		return s[strings.LastIndex(s, "_")+1:]
	}
	return string(e) + "值未定义"
}

// 申请状态类型
type TThostFtdcApplyStatusIDType byte

const THOST_FTDC_ASID_NoComplete TThostFtdcApplyStatusIDType = '1' // 未补全

const THOST_FTDC_ASID_Submited TThostFtdcApplyStatusIDType = '2' // 已提交

const THOST_FTDC_ASID_Checked TThostFtdcApplyStatusIDType = '3' // 已审核

const THOST_FTDC_ASID_Refused TThostFtdcApplyStatusIDType = '4' // 已拒绝

const THOST_FTDC_ASID_Deleted TThostFtdcApplyStatusIDType = '5' // 已删除

var mpTThostFtdcApplyStatusIDType = map[TThostFtdcApplyStatusIDType]string{'1': "THOST_FTDC_ASID_NoComplete", '2': "THOST_FTDC_ASID_Submited", '3': "THOST_FTDC_ASID_Checked", '4': "THOST_FTDC_ASID_Refused", '5': "THOST_FTDC_ASID_Deleted"}

func (e TThostFtdcApplyStatusIDType) String() string {
	if s, ok := mpTThostFtdcApplyStatusIDType[e]; ok {
		return s[strings.LastIndex(s, "_")+1:]
	}
	return string(e) + "值未定义"
}

// 发送方式类型
type TThostFtdcSendMethodType byte

const THOST_FTDC_UOASM_ByAPI TThostFtdcSendMethodType = '1' // 文件发送

const THOST_FTDC_UOASM_ByFile TThostFtdcSendMethodType = '2' // 电子发送

var mpTThostFtdcSendMethodType = map[TThostFtdcSendMethodType]string{'1': "THOST_FTDC_UOASM_ByAPI", '2': "THOST_FTDC_UOASM_ByFile"}

func (e TThostFtdcSendMethodType) String() string {
	if s, ok := mpTThostFtdcSendMethodType[e]; ok {
		return s[strings.LastIndex(s, "_")+1:]
	}
	return string(e) + "值未定义"
}

// 业务操作类型类型
type TThostFtdcEventTypeType Byte33

// 操作方法类型
type TThostFtdcEventModeType byte

const THOST_FTDC_EvM_ADD TThostFtdcEventModeType = '1' // 增加

const THOST_FTDC_EvM_UPDATE TThostFtdcEventModeType = '2' // 修改

const THOST_FTDC_EvM_DELETE TThostFtdcEventModeType = '3' // 删除

const THOST_FTDC_EvM_CHECK TThostFtdcEventModeType = '4' // 复核

const THOST_FTDC_EvM_COPY TThostFtdcEventModeType = '5' // 复制

const THOST_FTDC_EvM_CANCEL TThostFtdcEventModeType = '6' // 注销

const THOST_FTDC_EvM_Reverse TThostFtdcEventModeType = '7' // 冲销

var mpTThostFtdcEventModeType = map[TThostFtdcEventModeType]string{'1': "THOST_FTDC_EvM_ADD", '2': "THOST_FTDC_EvM_UPDATE", '3': "THOST_FTDC_EvM_DELETE", '4': "THOST_FTDC_EvM_CHECK", '5': "THOST_FTDC_EvM_COPY", '6': "THOST_FTDC_EvM_CANCEL", '7': "THOST_FTDC_EvM_Reverse"}

func (e TThostFtdcEventModeType) String() string {
	if s, ok := mpTThostFtdcEventModeType[e]; ok {
		return s[strings.LastIndex(s, "_")+1:]
	}
	return string(e) + "值未定义"
}

// 统一开户申请自动发送类型
type TThostFtdcUOAAutoSendType byte

const THOST_FTDC_UOAA_ASR TThostFtdcUOAAutoSendType = '1' // 自动发送并接收

const THOST_FTDC_UOAA_ASNR TThostFtdcUOAAutoSendType = '2' // 自动发送，不自动接收

const THOST_FTDC_UOAA_NSAR TThostFtdcUOAAutoSendType = '3' // 不自动发送，自动接收

const THOST_FTDC_UOAA_NSR TThostFtdcUOAAutoSendType = '4' // 不自动发送，也不自动接收

var mpTThostFtdcUOAAutoSendType = map[TThostFtdcUOAAutoSendType]string{'1': "THOST_FTDC_UOAA_ASR", '2': "THOST_FTDC_UOAA_ASNR", '3': "THOST_FTDC_UOAA_NSAR", '4': "THOST_FTDC_UOAA_NSR"}

func (e TThostFtdcUOAAutoSendType) String() string {
	if s, ok := mpTThostFtdcUOAAutoSendType[e]; ok {
		return s[strings.LastIndex(s, "_")+1:]
	}
	return string(e) + "值未定义"
}

// 查询深度类型
type TThostFtdcQueryDepthType int32

// 数据中心代码类型
type TThostFtdcDataCenterIDType int32

// 流程ID类型
type TThostFtdcFlowIDType byte

const THOST_FTDC_EvM_InvestorGroupFlow TThostFtdcFlowIDType = '1' // 投资者对应投资者组设置

const THOST_FTDC_EvM_InvestorRate TThostFtdcFlowIDType = '2' // 投资者手续费率设置

const THOST_FTDC_EvM_InvestorCommRateModel TThostFtdcFlowIDType = '3' // 投资者手续费率模板关系设置

var mpTThostFtdcFlowIDType = map[TThostFtdcFlowIDType]string{'1': "THOST_FTDC_EvM_InvestorGroupFlow", '2': "THOST_FTDC_EvM_InvestorRate", '3': "THOST_FTDC_EvM_InvestorCommRateModel"}

func (e TThostFtdcFlowIDType) String() string {
	if s, ok := mpTThostFtdcFlowIDType[e]; ok {
		return s[strings.LastIndex(s, "_")+1:]
	}
	return string(e) + "值未定义"
}

// 复核级别类型
type TThostFtdcCheckLevelType byte

const THOST_FTDC_CL_Zero TThostFtdcCheckLevelType = '0' // 零级复核

const THOST_FTDC_CL_One TThostFtdcCheckLevelType = '1' // 一级复核

const THOST_FTDC_CL_Two TThostFtdcCheckLevelType = '2' // 二级复核

var mpTThostFtdcCheckLevelType = map[TThostFtdcCheckLevelType]string{'0': "THOST_FTDC_CL_Zero", '1': "THOST_FTDC_CL_One", '2': "THOST_FTDC_CL_Two"}

func (e TThostFtdcCheckLevelType) String() string {
	if s, ok := mpTThostFtdcCheckLevelType[e]; ok {
		return s[strings.LastIndex(s, "_")+1:]
	}
	return string(e) + "值未定义"
}

// 操作次数类型
type TThostFtdcCheckNoType int32

// 复核级别类型
type TThostFtdcCheckStatusType byte

const THOST_FTDC_CHS_Init TThostFtdcCheckStatusType = '0' // 未复核

const THOST_FTDC_CHS_Checking TThostFtdcCheckStatusType = '1' // 复核中

const THOST_FTDC_CHS_Checked TThostFtdcCheckStatusType = '2' // 已复核

const THOST_FTDC_CHS_Refuse TThostFtdcCheckStatusType = '3' // 拒绝

const THOST_FTDC_CHS_Cancel TThostFtdcCheckStatusType = '4' // 作废

var mpTThostFtdcCheckStatusType = map[TThostFtdcCheckStatusType]string{'0': "THOST_FTDC_CHS_Init", '1': "THOST_FTDC_CHS_Checking", '2': "THOST_FTDC_CHS_Checked", '3': "THOST_FTDC_CHS_Refuse", '4': "THOST_FTDC_CHS_Cancel"}

func (e TThostFtdcCheckStatusType) String() string {
	if s, ok := mpTThostFtdcCheckStatusType[e]; ok {
		return s[strings.LastIndex(s, "_")+1:]
	}
	return string(e) + "值未定义"
}

// 生效状态类型
type TThostFtdcUsedStatusType byte

const THOST_FTDC_CHU_Unused TThostFtdcUsedStatusType = '0' // 未生效

const THOST_FTDC_CHU_Used TThostFtdcUsedStatusType = '1' // 已生效

const THOST_FTDC_CHU_Fail TThostFtdcUsedStatusType = '2' // 生效失败

var mpTThostFtdcUsedStatusType = map[TThostFtdcUsedStatusType]string{'0': "THOST_FTDC_CHU_Unused", '1': "THOST_FTDC_CHU_Used", '2': "THOST_FTDC_CHU_Fail"}

func (e TThostFtdcUsedStatusType) String() string {
	if s, ok := mpTThostFtdcUsedStatusType[e]; ok {
		return s[strings.LastIndex(s, "_")+1:]
	}
	return string(e) + "值未定义"
}

// 模型名称类型
type TThostFtdcRateTemplateNameType Byte61

// 用于查询的投资属性字段类型
type TThostFtdcPropertyStringType Byte2049

// 账户来源类型
type TThostFtdcBankAcountOriginType byte

const THOST_FTDC_BAO_ByAccProperty TThostFtdcBankAcountOriginType = '0' // 手工录入

const THOST_FTDC_BAO_ByFBTransfer TThostFtdcBankAcountOriginType = '1' // 银期转账

var mpTThostFtdcBankAcountOriginType = map[TThostFtdcBankAcountOriginType]string{'0': "THOST_FTDC_BAO_ByAccProperty", '1': "THOST_FTDC_BAO_ByFBTransfer"}

func (e TThostFtdcBankAcountOriginType) String() string {
	if s, ok := mpTThostFtdcBankAcountOriginType[e]; ok {
		return s[strings.LastIndex(s, "_")+1:]
	}
	return string(e) + "值未定义"
}

// 结算单月报成交汇总方式类型
type TThostFtdcMonthBillTradeSumType byte

const THOST_FTDC_MBTS_ByInstrument TThostFtdcMonthBillTradeSumType = '0' // 同日同合约

const THOST_FTDC_MBTS_ByDayInsPrc TThostFtdcMonthBillTradeSumType = '1' // 同日同合约同价格

const THOST_FTDC_MBTS_ByDayIns TThostFtdcMonthBillTradeSumType = '2' // 同合约

var mpTThostFtdcMonthBillTradeSumType = map[TThostFtdcMonthBillTradeSumType]string{'0': "THOST_FTDC_MBTS_ByInstrument", '1': "THOST_FTDC_MBTS_ByDayInsPrc", '2': "THOST_FTDC_MBTS_ByDayIns"}

func (e TThostFtdcMonthBillTradeSumType) String() string {
	if s, ok := mpTThostFtdcMonthBillTradeSumType[e]; ok {
		return s[strings.LastIndex(s, "_")+1:]
	}
	return string(e) + "值未定义"
}

// 银期交易代码枚举类型
type TThostFtdcFBTTradeCodeEnumType string

const THOST_FTDC_FTC_BankLaunchBankToBroker TThostFtdcFBTTradeCodeEnumType = "102001" // 银行发起银行转期货

const THOST_FTDC_FTC_BrokerLaunchBankToBroker TThostFtdcFBTTradeCodeEnumType = "202001" // 期货发起银行转期货

const THOST_FTDC_FTC_BankLaunchBrokerToBank TThostFtdcFBTTradeCodeEnumType = "102002" // 银行发起期货转银行

const THOST_FTDC_FTC_BrokerLaunchBrokerToBank TThostFtdcFBTTradeCodeEnumType = "202002" // 期货发起期货转银行

var mpTThostFtdcFBTTradeCodeEnumType = map[TThostFtdcFBTTradeCodeEnumType]string{"102001": "THOST_FTDC_FTC_BankLaunchBankToBroker", "202001": "THOST_FTDC_FTC_BrokerLaunchBankToBroker", "102002": "THOST_FTDC_FTC_BankLaunchBrokerToBank", "202002": "THOST_FTDC_FTC_BrokerLaunchBrokerToBank"}

func (e TThostFtdcFBTTradeCodeEnumType) String() string {
	if s, ok := mpTThostFtdcFBTTradeCodeEnumType[e]; ok {
		return s[strings.LastIndex(s, "_")+1:]
	}
	return string(e) + "值未定义"
}

// 模型代码类型
type TThostFtdcRateTemplateIDType Byte9

// 风险度类型
type TThostFtdcRiskRateType Byte21

// 时间戳类型
type TThostFtdcTimestampType int32

// 号段规则名称类型
type TThostFtdcInvestorIDRuleNameType Byte61

// 号段规则表达式类型
type TThostFtdcInvestorIDRuleExprType Byte513

// 上次OTP漂移值类型
type TThostFtdcLastDriftType int32

// 上次OTP成功值类型
type TThostFtdcLastSuccessType int32

// 令牌密钥类型
type TThostFtdcAuthKeyType Byte41

// 序列号类型
type TThostFtdcSerialNumberType Byte17

// 动态令牌类型类型
type TThostFtdcOTPTypeType byte

const THOST_FTDC_OTP_NONE TThostFtdcOTPTypeType = '0' // 无动态令牌

const THOST_FTDC_OTP_TOTP TThostFtdcOTPTypeType = '1' // 时间令牌

var mpTThostFtdcOTPTypeType = map[TThostFtdcOTPTypeType]string{'0': "THOST_FTDC_OTP_NONE", '1': "THOST_FTDC_OTP_TOTP"}

func (e TThostFtdcOTPTypeType) String() string {
	if s, ok := mpTThostFtdcOTPTypeType[e]; ok {
		return s[strings.LastIndex(s, "_")+1:]
	}
	return string(e) + "值未定义"
}

// 动态令牌提供商类型
type TThostFtdcOTPVendorsIDType Byte2

// 动态令牌提供商名称类型
type TThostFtdcOTPVendorsNameType Byte61

// 动态令牌状态类型
type TThostFtdcOTPStatusType byte

const THOST_FTDC_OTPS_Unused TThostFtdcOTPStatusType = '0' // 未使用

const THOST_FTDC_OTPS_Used TThostFtdcOTPStatusType = '1' // 已使用

const THOST_FTDC_OTPS_Disuse TThostFtdcOTPStatusType = '2' // 注销

var mpTThostFtdcOTPStatusType = map[TThostFtdcOTPStatusType]string{'0': "THOST_FTDC_OTPS_Unused", '1': "THOST_FTDC_OTPS_Used", '2': "THOST_FTDC_OTPS_Disuse"}

func (e TThostFtdcOTPStatusType) String() string {
	if s, ok := mpTThostFtdcOTPStatusType[e]; ok {
		return s[strings.LastIndex(s, "_")+1:]
	}
	return string(e) + "值未定义"
}

// 经济公司用户类型类型
type TThostFtdcBrokerUserTypeType byte

const THOST_FTDC_BUT_Investor TThostFtdcBrokerUserTypeType = '1' // 投资者

const THOST_FTDC_BUT_BrokerUser TThostFtdcBrokerUserTypeType = '2' // 操作员

var mpTThostFtdcBrokerUserTypeType = map[TThostFtdcBrokerUserTypeType]string{'1': "THOST_FTDC_BUT_Investor", '2': "THOST_FTDC_BUT_BrokerUser"}

func (e TThostFtdcBrokerUserTypeType) String() string {
	if s, ok := mpTThostFtdcBrokerUserTypeType[e]; ok {
		return s[strings.LastIndex(s, "_")+1:]
	}
	return string(e) + "值未定义"
}

// 期货类型类型
type TThostFtdcFutureTypeType byte

const THOST_FTDC_FUTT_Commodity TThostFtdcFutureTypeType = '1' // 商品期货

const THOST_FTDC_FUTT_Financial TThostFtdcFutureTypeType = '2' // 金融期货

var mpTThostFtdcFutureTypeType = map[TThostFtdcFutureTypeType]string{'1': "THOST_FTDC_FUTT_Commodity", '2': "THOST_FTDC_FUTT_Financial"}

func (e TThostFtdcFutureTypeType) String() string {
	if s, ok := mpTThostFtdcFutureTypeType[e]; ok {
		return s[strings.LastIndex(s, "_")+1:]
	}
	return string(e) + "值未定义"
}

// 资金管理操作类型类型
type TThostFtdcFundEventTypeType byte

const THOST_FTDC_FET_Restriction TThostFtdcFundEventTypeType = '0' // 转账限额

const THOST_FTDC_FET_TodayRestriction TThostFtdcFundEventTypeType = '1' // 当日转账限额

const THOST_FTDC_FET_Transfer TThostFtdcFundEventTypeType = '2' // 期商流水

const THOST_FTDC_FET_Credit TThostFtdcFundEventTypeType = '3' // 资金冻结

const THOST_FTDC_FET_InvestorWithdrawAlm TThostFtdcFundEventTypeType = '4' // 投资者可提资金比例

const THOST_FTDC_FET_BankRestriction TThostFtdcFundEventTypeType = '5' // 单个银行帐户转账限额

const THOST_FTDC_FET_Accountregister TThostFtdcFundEventTypeType = '6' // 银期签约账户

const THOST_FTDC_FET_ExchangeFundIO TThostFtdcFundEventTypeType = '7' // 交易所出入金

const THOST_FTDC_FET_InvestorFundIO TThostFtdcFundEventTypeType = '8' // 投资者出入金

var mpTThostFtdcFundEventTypeType = map[TThostFtdcFundEventTypeType]string{'0': "THOST_FTDC_FET_Restriction", '1': "THOST_FTDC_FET_TodayRestriction", '2': "THOST_FTDC_FET_Transfer", '3': "THOST_FTDC_FET_Credit", '4': "THOST_FTDC_FET_InvestorWithdrawAlm", '5': "THOST_FTDC_FET_BankRestriction", '6': "THOST_FTDC_FET_Accountregister", '7': "THOST_FTDC_FET_ExchangeFundIO", '8': "THOST_FTDC_FET_InvestorFundIO"}

func (e TThostFtdcFundEventTypeType) String() string {
	if s, ok := mpTThostFtdcFundEventTypeType[e]; ok {
		return s[strings.LastIndex(s, "_")+1:]
	}
	return string(e) + "值未定义"
}

// 资金账户来源类型
type TThostFtdcAccountSourceTypeType byte

const THOST_FTDC_AST_FBTransfer TThostFtdcAccountSourceTypeType = '0' // 银期同步

const THOST_FTDC_AST_ManualEntry TThostFtdcAccountSourceTypeType = '1' // 手工录入

var mpTThostFtdcAccountSourceTypeType = map[TThostFtdcAccountSourceTypeType]string{'0': "THOST_FTDC_AST_FBTransfer", '1': "THOST_FTDC_AST_ManualEntry"}

func (e TThostFtdcAccountSourceTypeType) String() string {
	if s, ok := mpTThostFtdcAccountSourceTypeType[e]; ok {
		return s[strings.LastIndex(s, "_")+1:]
	}
	return string(e) + "值未定义"
}

// 交易编码来源类型
type TThostFtdcCodeSourceTypeType byte

const THOST_FTDC_CST_UnifyAccount TThostFtdcCodeSourceTypeType = '0' // 统一开户(已规范)

const THOST_FTDC_CST_ManualEntry TThostFtdcCodeSourceTypeType = '1' // 手工录入(未规范)

var mpTThostFtdcCodeSourceTypeType = map[TThostFtdcCodeSourceTypeType]string{'0': "THOST_FTDC_CST_UnifyAccount", '1': "THOST_FTDC_CST_ManualEntry"}

func (e TThostFtdcCodeSourceTypeType) String() string {
	if s, ok := mpTThostFtdcCodeSourceTypeType[e]; ok {
		return s[strings.LastIndex(s, "_")+1:]
	}
	return string(e) + "值未定义"
}

// 操作员范围类型
type TThostFtdcUserRangeType byte

const THOST_FTDC_UR_All TThostFtdcUserRangeType = '0' // 所有

const THOST_FTDC_UR_Single TThostFtdcUserRangeType = '1' // 单一操作员

var mpTThostFtdcUserRangeType = map[TThostFtdcUserRangeType]string{'0': "THOST_FTDC_UR_All", '1': "THOST_FTDC_UR_Single"}

func (e TThostFtdcUserRangeType) String() string {
	if s, ok := mpTThostFtdcUserRangeType[e]; ok {
		return s[strings.LastIndex(s, "_")+1:]
	}
	return string(e) + "值未定义"
}

// 时间跨度类型
type TThostFtdcTimeSpanType Byte9

// 动态令牌导入批次编号类型
type TThostFtdcImportSequenceIDType Byte17

// 交易统计表按客户统计方式类型
type TThostFtdcByGroupType byte

const THOST_FTDC_BG_Investor TThostFtdcByGroupType = '2' // 按投资者统计

const THOST_FTDC_BG_Group TThostFtdcByGroupType = '1' // 按类统计

var mpTThostFtdcByGroupType = map[TThostFtdcByGroupType]string{'2': "THOST_FTDC_BG_Investor", '1': "THOST_FTDC_BG_Group"}

func (e TThostFtdcByGroupType) String() string {
	if s, ok := mpTThostFtdcByGroupType[e]; ok {
		return s[strings.LastIndex(s, "_")+1:]
	}
	return string(e) + "值未定义"
}

// 交易统计表按范围统计方式类型
type TThostFtdcTradeSumStatModeType byte

const THOST_FTDC_TSSM_Instrument TThostFtdcTradeSumStatModeType = '1' // 按合约统计

const THOST_FTDC_TSSM_Product TThostFtdcTradeSumStatModeType = '2' // 按产品统计

const THOST_FTDC_TSSM_Exchange TThostFtdcTradeSumStatModeType = '3' // 按交易所统计

var mpTThostFtdcTradeSumStatModeType = map[TThostFtdcTradeSumStatModeType]string{'1': "THOST_FTDC_TSSM_Instrument", '2': "THOST_FTDC_TSSM_Product", '3': "THOST_FTDC_TSSM_Exchange"}

func (e TThostFtdcTradeSumStatModeType) String() string {
	if s, ok := mpTThostFtdcTradeSumStatModeType[e]; ok {
		return s[strings.LastIndex(s, "_")+1:]
	}
	return string(e) + "值未定义"
}

// 组合成交类型类型
type TThostFtdcComTypeType int32

// 产品标识类型
type TThostFtdcUserProductIDType Byte33

// 产品名称类型
type TThostFtdcUserProductNameType Byte65

// 产品说明类型
type TThostFtdcUserProductMemoType Byte129

// 新增或变更标志类型
type TThostFtdcCSRCCancelFlagType Byte2

// 日期类型
type TThostFtdcCSRCDateType Byte11

// 客户名称类型
type TThostFtdcCSRCInvestorNameType Byte201

// 客户名称类型
type TThostFtdcCSRCOpenInvestorNameType Byte101

// 客户代码类型
type TThostFtdcCSRCInvestorIDType Byte13

// 证件号码类型
type TThostFtdcCSRCIdentifiedCardNoType Byte51

// 交易编码类型
type TThostFtdcCSRCClientIDType Byte11

// 银行标识类型
type TThostFtdcCSRCBankFlagType Byte3

// 银行账户类型
type TThostFtdcCSRCBankAccountType Byte23

// 开户人类型
type TThostFtdcCSRCOpenNameType Byte401

// 说明类型
type TThostFtdcCSRCMemoType Byte101

// 时间类型
type TThostFtdcCSRCTimeType Byte11

// 成交流水号类型
type TThostFtdcCSRCTradeIDType Byte21

// 合约代码类型
type TThostFtdcCSRCExchangeInstIDType Byte31

// 质押品名称类型
type TThostFtdcCSRCMortgageNameType Byte7

// 事由类型
type TThostFtdcCSRCReasonType Byte3

// 是否为非结算会员类型
type TThostFtdcIsSettlementType Byte2

// 资金类型
type TThostFtdcCSRCMoneyType Double

// 价格类型
type TThostFtdcCSRCPriceType Double

// 期权类型类型
type TThostFtdcCSRCOptionsTypeType Byte2

// 执行价类型
type TThostFtdcCSRCStrikePriceType Double

// 标的品种类型
type TThostFtdcCSRCTargetProductIDType Byte3

// 标的合约类型
type TThostFtdcCSRCTargetInstrIDType Byte31

// 手续费率模板名称类型
type TThostFtdcCommModelNameType Byte161

// 手续费率模板备注类型
type TThostFtdcCommModelMemoType Byte1025

// 日期表达式设置类型类型
type TThostFtdcExprSetModeType byte

const THOST_FTDC_ESM_Relative TThostFtdcExprSetModeType = '1' // 相对已有规则设置

const THOST_FTDC_ESM_Typical TThostFtdcExprSetModeType = '2' // 典型设置

var mpTThostFtdcExprSetModeType = map[TThostFtdcExprSetModeType]string{'1': "THOST_FTDC_ESM_Relative", '2': "THOST_FTDC_ESM_Typical"}

func (e TThostFtdcExprSetModeType) String() string {
	if s, ok := mpTThostFtdcExprSetModeType[e]; ok {
		return s[strings.LastIndex(s, "_")+1:]
	}
	return string(e) + "值未定义"
}

// 投资者范围类型
type TThostFtdcRateInvestorRangeType byte

const THOST_FTDC_RIR_All TThostFtdcRateInvestorRangeType = '1' // 公司标准

const THOST_FTDC_RIR_Model TThostFtdcRateInvestorRangeType = '2' // 模板

const THOST_FTDC_RIR_Single TThostFtdcRateInvestorRangeType = '3' // 单一投资者

var mpTThostFtdcRateInvestorRangeType = map[TThostFtdcRateInvestorRangeType]string{'1': "THOST_FTDC_RIR_All", '2': "THOST_FTDC_RIR_Model", '3': "THOST_FTDC_RIR_Single"}

func (e TThostFtdcRateInvestorRangeType) String() string {
	if s, ok := mpTThostFtdcRateInvestorRangeType[e]; ok {
		return s[strings.LastIndex(s, "_")+1:]
	}
	return string(e) + "值未定义"
}

// 代理经纪公司代码类型
type TThostFtdcAgentBrokerIDType Byte13

// 交易中心代码类型
type TThostFtdcDRIdentityIDType int32

// 交易中心名称类型
type TThostFtdcDRIdentityNameType Byte65

// DBLink标识号类型
type TThostFtdcDBLinkIDType Byte31

// 主次用系统数据同步状态类型
type TThostFtdcSyncDataStatusType byte

const THOST_FTDC_SDS_Initialize TThostFtdcSyncDataStatusType = '0' // 未同步

const THOST_FTDC_SDS_Settlementing TThostFtdcSyncDataStatusType = '1' // 同步中

const THOST_FTDC_SDS_Settlemented TThostFtdcSyncDataStatusType = '2' // 已同步

var mpTThostFtdcSyncDataStatusType = map[TThostFtdcSyncDataStatusType]string{'0': "THOST_FTDC_SDS_Initialize", '1': "THOST_FTDC_SDS_Settlementing", '2': "THOST_FTDC_SDS_Settlemented"}

func (e TThostFtdcSyncDataStatusType) String() string {
	if s, ok := mpTThostFtdcSyncDataStatusType[e]; ok {
		return s[strings.LastIndex(s, "_")+1:]
	}
	return string(e) + "值未定义"
}

// 成交来源类型
type TThostFtdcTradeSourceType byte

const THOST_FTDC_TSRC_NORMAL TThostFtdcTradeSourceType = '0' // 来自交易所普通回报

const THOST_FTDC_TSRC_QUERY TThostFtdcTradeSourceType = '1' // 来自查询

var mpTThostFtdcTradeSourceType = map[TThostFtdcTradeSourceType]string{'0': "THOST_FTDC_TSRC_NORMAL", '1': "THOST_FTDC_TSRC_QUERY"}

func (e TThostFtdcTradeSourceType) String() string {
	if s, ok := mpTThostFtdcTradeSourceType[e]; ok {
		return s[strings.LastIndex(s, "_")+1:]
	}
	return string(e) + "值未定义"
}

// 产品合约统计方式类型
type TThostFtdcFlexStatModeType byte

const THOST_FTDC_FSM_Product TThostFtdcFlexStatModeType = '1' // 产品统计

const THOST_FTDC_FSM_Exchange TThostFtdcFlexStatModeType = '2' // 交易所统计

const THOST_FTDC_FSM_All TThostFtdcFlexStatModeType = '3' // 统计所有

var mpTThostFtdcFlexStatModeType = map[TThostFtdcFlexStatModeType]string{'1': "THOST_FTDC_FSM_Product", '2': "THOST_FTDC_FSM_Exchange", '3': "THOST_FTDC_FSM_All"}

func (e TThostFtdcFlexStatModeType) String() string {
	if s, ok := mpTThostFtdcFlexStatModeType[e]; ok {
		return s[strings.LastIndex(s, "_")+1:]
	}
	return string(e) + "值未定义"
}

// 投资者范围统计方式类型
type TThostFtdcByInvestorRangeType byte

const THOST_FTDC_BIR_Property TThostFtdcByInvestorRangeType = '1' // 属性统计

const THOST_FTDC_BIR_All TThostFtdcByInvestorRangeType = '2' // 统计所有

var mpTThostFtdcByInvestorRangeType = map[TThostFtdcByInvestorRangeType]string{'1': "THOST_FTDC_BIR_Property", '2': "THOST_FTDC_BIR_All"}

func (e TThostFtdcByInvestorRangeType) String() string {
	if s, ok := mpTThostFtdcByInvestorRangeType[e]; ok {
		return s[strings.LastIndex(s, "_")+1:]
	}
	return string(e) + "值未定义"
}

// 风险度类型
type TThostFtdcSRiskRateType Byte21

// 序号类型
type TThostFtdcSequenceNo12Type int32

// 投资者范围类型
type TThostFtdcPropertyInvestorRangeType byte

const THOST_FTDC_PIR_All TThostFtdcPropertyInvestorRangeType = '1' // 所有

const THOST_FTDC_PIR_Property TThostFtdcPropertyInvestorRangeType = '2' // 投资者属性

const THOST_FTDC_PIR_Single TThostFtdcPropertyInvestorRangeType = '3' // 单一投资者

var mpTThostFtdcPropertyInvestorRangeType = map[TThostFtdcPropertyInvestorRangeType]string{'1': "THOST_FTDC_PIR_All", '2': "THOST_FTDC_PIR_Property", '3': "THOST_FTDC_PIR_Single"}

func (e TThostFtdcPropertyInvestorRangeType) String() string {
	if s, ok := mpTThostFtdcPropertyInvestorRangeType[e]; ok {
		return s[strings.LastIndex(s, "_")+1:]
	}
	return string(e) + "值未定义"
}

// 文件状态类型
type TThostFtdcFileStatusType byte

const THOST_FTDC_FIS_NoCreate TThostFtdcFileStatusType = '0' // 未生成

const THOST_FTDC_FIS_Created TThostFtdcFileStatusType = '1' // 已生成

const THOST_FTDC_FIS_Failed TThostFtdcFileStatusType = '2' // 生成失败

var mpTThostFtdcFileStatusType = map[TThostFtdcFileStatusType]string{'0': "THOST_FTDC_FIS_NoCreate", '1': "THOST_FTDC_FIS_Created", '2': "THOST_FTDC_FIS_Failed"}

func (e TThostFtdcFileStatusType) String() string {
	if s, ok := mpTThostFtdcFileStatusType[e]; ok {
		return s[strings.LastIndex(s, "_")+1:]
	}
	return string(e) + "值未定义"
}

// 文件生成方式类型
type TThostFtdcFileGenStyleType byte

const THOST_FTDC_FGS_FileTransmit TThostFtdcFileGenStyleType = '0' // 下发

const THOST_FTDC_FGS_FileGen TThostFtdcFileGenStyleType = '1' // 生成

var mpTThostFtdcFileGenStyleType = map[TThostFtdcFileGenStyleType]string{'0': "THOST_FTDC_FGS_FileTransmit", '1': "THOST_FTDC_FGS_FileGen"}

func (e TThostFtdcFileGenStyleType) String() string {
	if s, ok := mpTThostFtdcFileGenStyleType[e]; ok {
		return s[strings.LastIndex(s, "_")+1:]
	}
	return string(e) + "值未定义"
}

// 系统日志操作方法类型
type TThostFtdcSysOperModeType byte

const THOST_FTDC_SoM_Add TThostFtdcSysOperModeType = '1' // 增加

const THOST_FTDC_SoM_Update TThostFtdcSysOperModeType = '2' // 修改

const THOST_FTDC_SoM_Delete TThostFtdcSysOperModeType = '3' // 删除

const THOST_FTDC_SoM_Copy TThostFtdcSysOperModeType = '4' // 复制

const THOST_FTDC_SoM_AcTive TThostFtdcSysOperModeType = '5' // 激活

const THOST_FTDC_SoM_CanCel TThostFtdcSysOperModeType = '6' // 注销

const THOST_FTDC_SoM_ReSet TThostFtdcSysOperModeType = '7' // 重置

var mpTThostFtdcSysOperModeType = map[TThostFtdcSysOperModeType]string{'1': "THOST_FTDC_SoM_Add", '2': "THOST_FTDC_SoM_Update", '3': "THOST_FTDC_SoM_Delete", '4': "THOST_FTDC_SoM_Copy", '5': "THOST_FTDC_SoM_AcTive", '6': "THOST_FTDC_SoM_CanCel", '7': "THOST_FTDC_SoM_ReSet"}

func (e TThostFtdcSysOperModeType) String() string {
	if s, ok := mpTThostFtdcSysOperModeType[e]; ok {
		return s[strings.LastIndex(s, "_")+1:]
	}
	return string(e) + "值未定义"
}

// 系统日志操作类型类型
type TThostFtdcSysOperTypeType byte

const THOST_FTDC_SoT_UpdatePassword TThostFtdcSysOperTypeType = '0' // 修改操作员密码

const THOST_FTDC_SoT_UserDepartment TThostFtdcSysOperTypeType = '1' // 操作员组织架构关系

const THOST_FTDC_SoT_RoleManager TThostFtdcSysOperTypeType = '2' // 角色管理

const THOST_FTDC_SoT_RoleFunction TThostFtdcSysOperTypeType = '3' // 角色功能设置

const THOST_FTDC_SoT_BaseParam TThostFtdcSysOperTypeType = '4' // 基础参数设置

const THOST_FTDC_SoT_SetUserID TThostFtdcSysOperTypeType = '5' // 设置操作员

const THOST_FTDC_SoT_SetUserRole TThostFtdcSysOperTypeType = '6' // 用户角色设置

const THOST_FTDC_SoT_UserIpRestriction TThostFtdcSysOperTypeType = '7' // 用户IP限制

const THOST_FTDC_SoT_DepartmentManager TThostFtdcSysOperTypeType = '8' // 组织架构管理

const THOST_FTDC_SoT_DepartmentCopy TThostFtdcSysOperTypeType = '9' // 组织架构向查询分类复制

const THOST_FTDC_SoT_Tradingcode TThostFtdcSysOperTypeType = 'A' // 交易编码管理

const THOST_FTDC_SoT_InvestorStatus TThostFtdcSysOperTypeType = 'B' // 投资者状态维护

const THOST_FTDC_SoT_InvestorAuthority TThostFtdcSysOperTypeType = 'C' // 投资者权限管理

const THOST_FTDC_SoT_PropertySet TThostFtdcSysOperTypeType = 'D' // 属性设置

const THOST_FTDC_SoT_ReSetInvestorPasswd TThostFtdcSysOperTypeType = 'E' // 重置投资者密码

const THOST_FTDC_SoT_InvestorPersonalityInfo TThostFtdcSysOperTypeType = 'F' // 投资者个性信息维护

var mpTThostFtdcSysOperTypeType = map[TThostFtdcSysOperTypeType]string{'0': "THOST_FTDC_SoT_UpdatePassword", '1': "THOST_FTDC_SoT_UserDepartment", '2': "THOST_FTDC_SoT_RoleManager", '3': "THOST_FTDC_SoT_RoleFunction", '4': "THOST_FTDC_SoT_BaseParam", '5': "THOST_FTDC_SoT_SetUserID", '6': "THOST_FTDC_SoT_SetUserRole", '7': "THOST_FTDC_SoT_UserIpRestriction", '8': "THOST_FTDC_SoT_DepartmentManager", '9': "THOST_FTDC_SoT_DepartmentCopy", 'A': "THOST_FTDC_SoT_Tradingcode", 'B': "THOST_FTDC_SoT_InvestorStatus", 'C': "THOST_FTDC_SoT_InvestorAuthority", 'D': "THOST_FTDC_SoT_PropertySet", 'E': "THOST_FTDC_SoT_ReSetInvestorPasswd", 'F': "THOST_FTDC_SoT_InvestorPersonalityInfo"}

func (e TThostFtdcSysOperTypeType) String() string {
	if s, ok := mpTThostFtdcSysOperTypeType[e]; ok {
		return s[strings.LastIndex(s, "_")+1:]
	}
	return string(e) + "值未定义"
}

// 上报数据查询类型类型
type TThostFtdcCSRCDataQueyTypeType byte

const THOST_FTDC_CSRCQ_Current TThostFtdcCSRCDataQueyTypeType = '0' // 查询当前交易日报送的数据

const THOST_FTDC_CSRCQ_History TThostFtdcCSRCDataQueyTypeType = '1' // 查询历史报送的代理经纪公司的数据

var mpTThostFtdcCSRCDataQueyTypeType = map[TThostFtdcCSRCDataQueyTypeType]string{'0': "THOST_FTDC_CSRCQ_Current", '1': "THOST_FTDC_CSRCQ_History"}

func (e TThostFtdcCSRCDataQueyTypeType) String() string {
	if s, ok := mpTThostFtdcCSRCDataQueyTypeType[e]; ok {
		return s[strings.LastIndex(s, "_")+1:]
	}
	return string(e) + "值未定义"
}

// 休眠状态类型
type TThostFtdcFreezeStatusType byte

const THOST_FTDC_FRS_Normal TThostFtdcFreezeStatusType = '1' // 活跃

const THOST_FTDC_FRS_Freeze TThostFtdcFreezeStatusType = '0' // 休眠

var mpTThostFtdcFreezeStatusType = map[TThostFtdcFreezeStatusType]string{'1': "THOST_FTDC_FRS_Normal", '0': "THOST_FTDC_FRS_Freeze"}

func (e TThostFtdcFreezeStatusType) String() string {
	if s, ok := mpTThostFtdcFreezeStatusType[e]; ok {
		return s[strings.LastIndex(s, "_")+1:]
	}
	return string(e) + "值未定义"
}

// 规范状态类型
type TThostFtdcStandardStatusType byte

const THOST_FTDC_STST_Standard TThostFtdcStandardStatusType = '0' // 已规范

const THOST_FTDC_STST_NonStandard TThostFtdcStandardStatusType = '1' // 未规范

var mpTThostFtdcStandardStatusType = map[TThostFtdcStandardStatusType]string{'0': "THOST_FTDC_STST_Standard", '1': "THOST_FTDC_STST_NonStandard"}

func (e TThostFtdcStandardStatusType) String() string {
	if s, ok := mpTThostFtdcStandardStatusType[e]; ok {
		return s[strings.LastIndex(s, "_")+1:]
	}
	return string(e) + "值未定义"
}

// 休眠状态类型
type TThostFtdcCSRCFreezeStatusType Byte2

// 配置类型类型
type TThostFtdcRightParamTypeType byte

const THOST_FTDC_RPT_Freeze TThostFtdcRightParamTypeType = '1' // 休眠户

const THOST_FTDC_RPT_FreezeActive TThostFtdcRightParamTypeType = '2' // 激活休眠户

const THOST_FTDC_RPT_OpenLimit TThostFtdcRightParamTypeType = '3' // 开仓权限限制

const THOST_FTDC_RPT_RelieveOpenLimit TThostFtdcRightParamTypeType = '4' // 解除开仓权限限制

var mpTThostFtdcRightParamTypeType = map[TThostFtdcRightParamTypeType]string{'1': "THOST_FTDC_RPT_Freeze", '2': "THOST_FTDC_RPT_FreezeActive", '3': "THOST_FTDC_RPT_OpenLimit", '4': "THOST_FTDC_RPT_RelieveOpenLimit"}

func (e TThostFtdcRightParamTypeType) String() string {
	if s, ok := mpTThostFtdcRightParamTypeType[e]; ok {
		return s[strings.LastIndex(s, "_")+1:]
	}
	return string(e) + "值未定义"
}

// 模板代码类型
type TThostFtdcRightTemplateIDType Byte9

// 模板名称类型
type TThostFtdcRightTemplateNameType Byte61

// 反洗钱审核表数据状态类型
type TThostFtdcDataStatusType byte

const THOST_FTDC_AMLDS_Normal TThostFtdcDataStatusType = '0' // 正常

const THOST_FTDC_AMLDS_Deleted TThostFtdcDataStatusType = '1' // 已删除

var mpTThostFtdcDataStatusType = map[TThostFtdcDataStatusType]string{'0': "THOST_FTDC_AMLDS_Normal", '1': "THOST_FTDC_AMLDS_Deleted"}

func (e TThostFtdcDataStatusType) String() string {
	if s, ok := mpTThostFtdcDataStatusType[e]; ok {
		return s[strings.LastIndex(s, "_")+1:]
	}
	return string(e) + "值未定义"
}

// 审核状态类型
type TThostFtdcAMLCheckStatusType byte

const THOST_FTDC_AMLCHS_Init TThostFtdcAMLCheckStatusType = '0' // 未复核

const THOST_FTDC_AMLCHS_Checking TThostFtdcAMLCheckStatusType = '1' // 复核中

const THOST_FTDC_AMLCHS_Checked TThostFtdcAMLCheckStatusType = '2' // 已复核

const THOST_FTDC_AMLCHS_RefuseReport TThostFtdcAMLCheckStatusType = '3' // 拒绝上报

var mpTThostFtdcAMLCheckStatusType = map[TThostFtdcAMLCheckStatusType]string{'0': "THOST_FTDC_AMLCHS_Init", '1': "THOST_FTDC_AMLCHS_Checking", '2': "THOST_FTDC_AMLCHS_Checked", '3': "THOST_FTDC_AMLCHS_RefuseReport"}

func (e TThostFtdcAMLCheckStatusType) String() string {
	if s, ok := mpTThostFtdcAMLCheckStatusType[e]; ok {
		return s[strings.LastIndex(s, "_")+1:]
	}
	return string(e) + "值未定义"
}

// 日期类型类型
type TThostFtdcAmlDateTypeType byte

const THOST_FTDC_AMLDT_DrawDay TThostFtdcAmlDateTypeType = '0' // 检查日期

const THOST_FTDC_AMLDT_TouchDay TThostFtdcAmlDateTypeType = '1' // 发生日期

var mpTThostFtdcAmlDateTypeType = map[TThostFtdcAmlDateTypeType]string{'0': "THOST_FTDC_AMLDT_DrawDay", '1': "THOST_FTDC_AMLDT_TouchDay"}

func (e TThostFtdcAmlDateTypeType) String() string {
	if s, ok := mpTThostFtdcAmlDateTypeType[e]; ok {
		return s[strings.LastIndex(s, "_")+1:]
	}
	return string(e) + "值未定义"
}

// 审核级别类型
type TThostFtdcAmlCheckLevelType byte

const THOST_FTDC_AMLCL_CheckLevel0 TThostFtdcAmlCheckLevelType = '0' // 零级审核

const THOST_FTDC_AMLCL_CheckLevel1 TThostFtdcAmlCheckLevelType = '1' // 一级审核

const THOST_FTDC_AMLCL_CheckLevel2 TThostFtdcAmlCheckLevelType = '2' // 二级审核

const THOST_FTDC_AMLCL_CheckLevel3 TThostFtdcAmlCheckLevelType = '3' // 三级审核

var mpTThostFtdcAmlCheckLevelType = map[TThostFtdcAmlCheckLevelType]string{'0': "THOST_FTDC_AMLCL_CheckLevel0", '1': "THOST_FTDC_AMLCL_CheckLevel1", '2': "THOST_FTDC_AMLCL_CheckLevel2", '3': "THOST_FTDC_AMLCL_CheckLevel3"}

func (e TThostFtdcAmlCheckLevelType) String() string {
	if s, ok := mpTThostFtdcAmlCheckLevelType[e]; ok {
		return s[strings.LastIndex(s, "_")+1:]
	}
	return string(e) + "值未定义"
}

// 反洗钱数据抽取审核流程类型
type TThostFtdcAmlCheckFlowType Byte2

// 数据类型类型
type TThostFtdcDataTypeType Byte129

// 导出文件类型类型
type TThostFtdcExportFileTypeType byte

const THOST_FTDC_EFT_CSV TThostFtdcExportFileTypeType = '0' // CSV

const THOST_FTDC_EFT_EXCEL TThostFtdcExportFileTypeType = '1' // Excel

const THOST_FTDC_EFT_DBF TThostFtdcExportFileTypeType = '2' // DBF

var mpTThostFtdcExportFileTypeType = map[TThostFtdcExportFileTypeType]string{'0': "THOST_FTDC_EFT_CSV", '1': "THOST_FTDC_EFT_EXCEL", '2': "THOST_FTDC_EFT_DBF"}

func (e TThostFtdcExportFileTypeType) String() string {
	if s, ok := mpTThostFtdcExportFileTypeType[e]; ok {
		return s[strings.LastIndex(s, "_")+1:]
	}
	return string(e) + "值未定义"
}

// 结算配置类型类型
type TThostFtdcSettleManagerTypeType byte

const THOST_FTDC_SMT_Before TThostFtdcSettleManagerTypeType = '1' // 结算前准备

const THOST_FTDC_SMT_Settlement TThostFtdcSettleManagerTypeType = '2' // 结算

const THOST_FTDC_SMT_After TThostFtdcSettleManagerTypeType = '3' // 结算后核对

const THOST_FTDC_SMT_Settlemented TThostFtdcSettleManagerTypeType = '4' // 结算后处理

var mpTThostFtdcSettleManagerTypeType = map[TThostFtdcSettleManagerTypeType]string{'1': "THOST_FTDC_SMT_Before", '2': "THOST_FTDC_SMT_Settlement", '3': "THOST_FTDC_SMT_After", '4': "THOST_FTDC_SMT_Settlemented"}

func (e TThostFtdcSettleManagerTypeType) String() string {
	if s, ok := mpTThostFtdcSettleManagerTypeType[e]; ok {
		return s[strings.LastIndex(s, "_")+1:]
	}
	return string(e) + "值未定义"
}

// 结算配置代码类型
type TThostFtdcSettleManagerIDType Byte33

// 结算配置名称类型
type TThostFtdcSettleManagerNameType Byte129

// 结算配置等级类型
type TThostFtdcSettleManagerLevelType byte

const THOST_FTDC_SML_Must TThostFtdcSettleManagerLevelType = '1' // 必要

const THOST_FTDC_SML_Alarm TThostFtdcSettleManagerLevelType = '2' // 警告

const THOST_FTDC_SML_Prompt TThostFtdcSettleManagerLevelType = '3' // 提示

const THOST_FTDC_SML_Ignore TThostFtdcSettleManagerLevelType = '4' // 不检查

var mpTThostFtdcSettleManagerLevelType = map[TThostFtdcSettleManagerLevelType]string{'1': "THOST_FTDC_SML_Must", '2': "THOST_FTDC_SML_Alarm", '3': "THOST_FTDC_SML_Prompt", '4': "THOST_FTDC_SML_Ignore"}

func (e TThostFtdcSettleManagerLevelType) String() string {
	if s, ok := mpTThostFtdcSettleManagerLevelType[e]; ok {
		return s[strings.LastIndex(s, "_")+1:]
	}
	return string(e) + "值未定义"
}

// 模块分组类型
type TThostFtdcSettleManagerGroupType byte

const THOST_FTDC_SMG_Exhcange TThostFtdcSettleManagerGroupType = '1' // 交易所核对

const THOST_FTDC_SMG_ASP TThostFtdcSettleManagerGroupType = '2' // 内部核对

const THOST_FTDC_SMG_CSRC TThostFtdcSettleManagerGroupType = '3' // 上报数据核对

var mpTThostFtdcSettleManagerGroupType = map[TThostFtdcSettleManagerGroupType]string{'1': "THOST_FTDC_SMG_Exhcange", '2': "THOST_FTDC_SMG_ASP", '3': "THOST_FTDC_SMG_CSRC"}

func (e TThostFtdcSettleManagerGroupType) String() string {
	if s, ok := mpTThostFtdcSettleManagerGroupType[e]; ok {
		return s[strings.LastIndex(s, "_")+1:]
	}
	return string(e) + "值未定义"
}

// 核对结果说明类型
type TThostFtdcCheckResultMemoType Byte1025

// 功能链接类型
type TThostFtdcFunctionUrlType Byte1025

// 客户端认证信息类型
type TThostFtdcAuthInfoType Byte129

// 客户端认证码类型
type TThostFtdcAuthCodeType Byte17

// 保值额度使用类型类型
type TThostFtdcLimitUseTypeType byte

const THOST_FTDC_LUT_Repeatable TThostFtdcLimitUseTypeType = '1' // 可重复使用

const THOST_FTDC_LUT_Unrepeatable TThostFtdcLimitUseTypeType = '2' // 不可重复使用

var mpTThostFtdcLimitUseTypeType = map[TThostFtdcLimitUseTypeType]string{'1': "THOST_FTDC_LUT_Repeatable", '2': "THOST_FTDC_LUT_Unrepeatable"}

func (e TThostFtdcLimitUseTypeType) String() string {
	if s, ok := mpTThostFtdcLimitUseTypeType[e]; ok {
		return s[strings.LastIndex(s, "_")+1:]
	}
	return string(e) + "值未定义"
}

// 数据来源类型
type TThostFtdcDataResourceType byte

const THOST_FTDC_DAR_Settle TThostFtdcDataResourceType = '1' // 本系统

const THOST_FTDC_DAR_Exchange TThostFtdcDataResourceType = '2' // 交易所

const THOST_FTDC_DAR_CSRC TThostFtdcDataResourceType = '3' // 报送数据

var mpTThostFtdcDataResourceType = map[TThostFtdcDataResourceType]string{'1': "THOST_FTDC_DAR_Settle", '2': "THOST_FTDC_DAR_Exchange", '3': "THOST_FTDC_DAR_CSRC"}

func (e TThostFtdcDataResourceType) String() string {
	if s, ok := mpTThostFtdcDataResourceType[e]; ok {
		return s[strings.LastIndex(s, "_")+1:]
	}
	return string(e) + "值未定义"
}

// 保证金类型类型
type TThostFtdcMarginTypeType byte

const THOST_FTDC_MGT_ExchMarginRate TThostFtdcMarginTypeType = '0' // 交易所保证金率

const THOST_FTDC_MGT_InstrMarginRate TThostFtdcMarginTypeType = '1' // 投资者保证金率

const THOST_FTDC_MGT_InstrMarginRateTrade TThostFtdcMarginTypeType = '2' // 投资者交易保证金率

var mpTThostFtdcMarginTypeType = map[TThostFtdcMarginTypeType]string{'0': "THOST_FTDC_MGT_ExchMarginRate", '1': "THOST_FTDC_MGT_InstrMarginRate", '2': "THOST_FTDC_MGT_InstrMarginRateTrade"}

func (e TThostFtdcMarginTypeType) String() string {
	if s, ok := mpTThostFtdcMarginTypeType[e]; ok {
		return s[strings.LastIndex(s, "_")+1:]
	}
	return string(e) + "值未定义"
}

// 生效类型类型
type TThostFtdcActiveTypeType byte

const THOST_FTDC_ACT_Intraday TThostFtdcActiveTypeType = '1' // 仅当日生效

const THOST_FTDC_ACT_Long TThostFtdcActiveTypeType = '2' // 长期生效

var mpTThostFtdcActiveTypeType = map[TThostFtdcActiveTypeType]string{'1': "THOST_FTDC_ACT_Intraday", '2': "THOST_FTDC_ACT_Long"}

func (e TThostFtdcActiveTypeType) String() string {
	if s, ok := mpTThostFtdcActiveTypeType[e]; ok {
		return s[strings.LastIndex(s, "_")+1:]
	}
	return string(e) + "值未定义"
}

// 冲突保证金率类型类型
type TThostFtdcMarginRateTypeType byte

const THOST_FTDC_MRT_Exchange TThostFtdcMarginRateTypeType = '1' // 交易所保证金率

const THOST_FTDC_MRT_Investor TThostFtdcMarginRateTypeType = '2' // 投资者保证金率

const THOST_FTDC_MRT_InvestorTrade TThostFtdcMarginRateTypeType = '3' // 投资者交易保证金率

var mpTThostFtdcMarginRateTypeType = map[TThostFtdcMarginRateTypeType]string{'1': "THOST_FTDC_MRT_Exchange", '2': "THOST_FTDC_MRT_Investor", '3': "THOST_FTDC_MRT_InvestorTrade"}

func (e TThostFtdcMarginRateTypeType) String() string {
	if s, ok := mpTThostFtdcMarginRateTypeType[e]; ok {
		return s[strings.LastIndex(s, "_")+1:]
	}
	return string(e) + "值未定义"
}

// 备份数据状态类型
type TThostFtdcBackUpStatusType byte

const THOST_FTDC_BUS_UnBak TThostFtdcBackUpStatusType = '0' // 未生成备份数据

const THOST_FTDC_BUS_BakUp TThostFtdcBackUpStatusType = '1' // 备份数据生成中

const THOST_FTDC_BUS_BakUped TThostFtdcBackUpStatusType = '2' // 已生成备份数据

const THOST_FTDC_BUS_BakFail TThostFtdcBackUpStatusType = '3' // 备份数据失败

var mpTThostFtdcBackUpStatusType = map[TThostFtdcBackUpStatusType]string{'0': "THOST_FTDC_BUS_UnBak", '1': "THOST_FTDC_BUS_BakUp", '2': "THOST_FTDC_BUS_BakUped", '3': "THOST_FTDC_BUS_BakFail"}

func (e TThostFtdcBackUpStatusType) String() string {
	if s, ok := mpTThostFtdcBackUpStatusType[e]; ok {
		return s[strings.LastIndex(s, "_")+1:]
	}
	return string(e) + "值未定义"
}

// 结算初始化状态类型
type TThostFtdcInitSettlementType byte

const THOST_FTDC_SIS_UnInitialize TThostFtdcInitSettlementType = '0' // 结算初始化未开始

const THOST_FTDC_SIS_Initialize TThostFtdcInitSettlementType = '1' // 结算初始化中

const THOST_FTDC_SIS_Initialized TThostFtdcInitSettlementType = '2' // 结算初始化完成

var mpTThostFtdcInitSettlementType = map[TThostFtdcInitSettlementType]string{'0': "THOST_FTDC_SIS_UnInitialize", '1': "THOST_FTDC_SIS_Initialize", '2': "THOST_FTDC_SIS_Initialized"}

func (e TThostFtdcInitSettlementType) String() string {
	if s, ok := mpTThostFtdcInitSettlementType[e]; ok {
		return s[strings.LastIndex(s, "_")+1:]
	}
	return string(e) + "值未定义"
}

// 报表数据生成状态类型
type TThostFtdcReportStatusType byte

const THOST_FTDC_SRS_NoCreate TThostFtdcReportStatusType = '0' // 未生成报表数据

const THOST_FTDC_SRS_Create TThostFtdcReportStatusType = '1' // 报表数据生成中

const THOST_FTDC_SRS_Created TThostFtdcReportStatusType = '2' // 已生成报表数据

const THOST_FTDC_SRS_CreateFail TThostFtdcReportStatusType = '3' // 生成报表数据失败

var mpTThostFtdcReportStatusType = map[TThostFtdcReportStatusType]string{'0': "THOST_FTDC_SRS_NoCreate", '1': "THOST_FTDC_SRS_Create", '2': "THOST_FTDC_SRS_Created", '3': "THOST_FTDC_SRS_CreateFail"}

func (e TThostFtdcReportStatusType) String() string {
	if s, ok := mpTThostFtdcReportStatusType[e]; ok {
		return s[strings.LastIndex(s, "_")+1:]
	}
	return string(e) + "值未定义"
}

// 数据归档状态类型
type TThostFtdcSaveStatusType byte

const THOST_FTDC_SSS_UnSaveData TThostFtdcSaveStatusType = '0' // 归档未完成

const THOST_FTDC_SSS_SaveDatad TThostFtdcSaveStatusType = '1' // 归档完成

var mpTThostFtdcSaveStatusType = map[TThostFtdcSaveStatusType]string{'0': "THOST_FTDC_SSS_UnSaveData", '1': "THOST_FTDC_SSS_SaveDatad"}

func (e TThostFtdcSaveStatusType) String() string {
	if s, ok := mpTThostFtdcSaveStatusType[e]; ok {
		return s[strings.LastIndex(s, "_")+1:]
	}
	return string(e) + "值未定义"
}

// 结算确认数据归档状态类型
type TThostFtdcSettArchiveStatusType byte

const THOST_FTDC_SAS_UnArchived TThostFtdcSettArchiveStatusType = '0' // 未归档数据

const THOST_FTDC_SAS_Archiving TThostFtdcSettArchiveStatusType = '1' // 数据归档中

const THOST_FTDC_SAS_Archived TThostFtdcSettArchiveStatusType = '2' // 已归档数据

const THOST_FTDC_SAS_ArchiveFail TThostFtdcSettArchiveStatusType = '3' // 归档数据失败

var mpTThostFtdcSettArchiveStatusType = map[TThostFtdcSettArchiveStatusType]string{'0': "THOST_FTDC_SAS_UnArchived", '1': "THOST_FTDC_SAS_Archiving", '2': "THOST_FTDC_SAS_Archived", '3': "THOST_FTDC_SAS_ArchiveFail"}

func (e TThostFtdcSettArchiveStatusType) String() string {
	if s, ok := mpTThostFtdcSettArchiveStatusType[e]; ok {
		return s[strings.LastIndex(s, "_")+1:]
	}
	return string(e) + "值未定义"
}

// CTP交易系统类型类型
type TThostFtdcCTPTypeType byte

const THOST_FTDC_CTPT_Unkown TThostFtdcCTPTypeType = '0' // 未知类型

const THOST_FTDC_CTPT_MainCenter TThostFtdcCTPTypeType = '1' // 主中心

const THOST_FTDC_CTPT_BackUp TThostFtdcCTPTypeType = '2' // 备中心

var mpTThostFtdcCTPTypeType = map[TThostFtdcCTPTypeType]string{'0': "THOST_FTDC_CTPT_Unkown", '1': "THOST_FTDC_CTPT_MainCenter", '2': "THOST_FTDC_CTPT_BackUp"}

func (e TThostFtdcCTPTypeType) String() string {
	if s, ok := mpTThostFtdcCTPTypeType[e]; ok {
		return s[strings.LastIndex(s, "_")+1:]
	}
	return string(e) + "值未定义"
}

// 工具代码类型
type TThostFtdcToolIDType Byte9

// 工具名称类型
type TThostFtdcToolNameType Byte81

// 平仓处理类型类型
type TThostFtdcCloseDealTypeType byte

const THOST_FTDC_CDT_Normal TThostFtdcCloseDealTypeType = '0' // 正常

const THOST_FTDC_CDT_SpecFirst TThostFtdcCloseDealTypeType = '1' // 投机平仓优先

var mpTThostFtdcCloseDealTypeType = map[TThostFtdcCloseDealTypeType]string{'0': "THOST_FTDC_CDT_Normal", '1': "THOST_FTDC_CDT_SpecFirst"}

func (e TThostFtdcCloseDealTypeType) String() string {
	if s, ok := mpTThostFtdcCloseDealTypeType[e]; ok {
		return s[strings.LastIndex(s, "_")+1:]
	}
	return string(e) + "值未定义"
}

// 货币质押资金可用范围类型
type TThostFtdcMortgageFundUseRangeType byte

const THOST_FTDC_MFUR_None TThostFtdcMortgageFundUseRangeType = '0' // 不能使用

const THOST_FTDC_MFUR_Margin TThostFtdcMortgageFundUseRangeType = '1' // 用于保证金

const THOST_FTDC_MFUR_All TThostFtdcMortgageFundUseRangeType = '2' // 用于手续费、盈亏、保证金

const THOST_FTDC_MFUR_CNY3 TThostFtdcMortgageFundUseRangeType = '3' // 人民币方案3

var mpTThostFtdcMortgageFundUseRangeType = map[TThostFtdcMortgageFundUseRangeType]string{'0': "THOST_FTDC_MFUR_None", '1': "THOST_FTDC_MFUR_Margin", '2': "THOST_FTDC_MFUR_All", '3': "THOST_FTDC_MFUR_CNY3"}

func (e TThostFtdcMortgageFundUseRangeType) String() string {
	if s, ok := mpTThostFtdcMortgageFundUseRangeType[e]; ok {
		return s[strings.LastIndex(s, "_")+1:]
	}
	return string(e) + "值未定义"
}

// 币种单位数量类型
type TThostFtdcCurrencyUnitType Double

// 汇率类型
type TThostFtdcExchangeRateType Double

// 特殊产品类型类型
type TThostFtdcSpecProductTypeType byte

const THOST_FTDC_SPT_CzceHedge TThostFtdcSpecProductTypeType = '1' // 郑商所套保产品

const THOST_FTDC_SPT_IneForeignCurrency TThostFtdcSpecProductTypeType = '2' // 货币质押产品

const THOST_FTDC_SPT_DceOpenClose TThostFtdcSpecProductTypeType = '3' // 大连短线开平仓产品

var mpTThostFtdcSpecProductTypeType = map[TThostFtdcSpecProductTypeType]string{'1': "THOST_FTDC_SPT_CzceHedge", '2': "THOST_FTDC_SPT_IneForeignCurrency", '3': "THOST_FTDC_SPT_DceOpenClose"}

func (e TThostFtdcSpecProductTypeType) String() string {
	if s, ok := mpTThostFtdcSpecProductTypeType[e]; ok {
		return s[strings.LastIndex(s, "_")+1:]
	}
	return string(e) + "值未定义"
}

// 货币质押类型类型
type TThostFtdcFundMortgageTypeType byte

const THOST_FTDC_FMT_Mortgage TThostFtdcFundMortgageTypeType = '1' // 质押

const THOST_FTDC_FMT_Redemption TThostFtdcFundMortgageTypeType = '2' // 解质

var mpTThostFtdcFundMortgageTypeType = map[TThostFtdcFundMortgageTypeType]string{'1': "THOST_FTDC_FMT_Mortgage", '2': "THOST_FTDC_FMT_Redemption"}

func (e TThostFtdcFundMortgageTypeType) String() string {
	if s, ok := mpTThostFtdcFundMortgageTypeType[e]; ok {
		return s[strings.LastIndex(s, "_")+1:]
	}
	return string(e) + "值未定义"
}

// 投资者账户结算参数代码类型
type TThostFtdcAccountSettlementParamIDType byte

const THOST_FTDC_ASPI_BaseMargin TThostFtdcAccountSettlementParamIDType = '1' // 基础保证金

const THOST_FTDC_ASPI_LowestInterest TThostFtdcAccountSettlementParamIDType = '2' // 最低权益标准

var mpTThostFtdcAccountSettlementParamIDType = map[TThostFtdcAccountSettlementParamIDType]string{'1': "THOST_FTDC_ASPI_BaseMargin", '2': "THOST_FTDC_ASPI_LowestInterest"}

func (e TThostFtdcAccountSettlementParamIDType) String() string {
	if s, ok := mpTThostFtdcAccountSettlementParamIDType[e]; ok {
		return s[strings.LastIndex(s, "_")+1:]
	}
	return string(e) + "值未定义"
}

// 币种名称类型
type TThostFtdcCurrencyNameType Byte31

// 币种符号类型
type TThostFtdcCurrencySignType Byte4

// 货币质押方向类型
type TThostFtdcFundMortDirectionType byte

const THOST_FTDC_FMD_In TThostFtdcFundMortDirectionType = '1' // 货币质入

const THOST_FTDC_FMD_Out TThostFtdcFundMortDirectionType = '2' // 货币质出

var mpTThostFtdcFundMortDirectionType = map[TThostFtdcFundMortDirectionType]string{'1': "THOST_FTDC_FMD_In", '2': "THOST_FTDC_FMD_Out"}

func (e TThostFtdcFundMortDirectionType) String() string {
	if s, ok := mpTThostFtdcFundMortDirectionType[e]; ok {
		return s[strings.LastIndex(s, "_")+1:]
	}
	return string(e) + "值未定义"
}

// 换汇类别类型
type TThostFtdcBusinessClassType byte

const THOST_FTDC_BT_Profit TThostFtdcBusinessClassType = '0' // 盈利

const THOST_FTDC_BT_Loss TThostFtdcBusinessClassType = '1' // 亏损

const THOST_FTDC_BT_Other TThostFtdcBusinessClassType = 'Z' // 其他

var mpTThostFtdcBusinessClassType = map[TThostFtdcBusinessClassType]string{'0': "THOST_FTDC_BT_Profit", '1': "THOST_FTDC_BT_Loss", 'Z': "THOST_FTDC_BT_Other"}

func (e TThostFtdcBusinessClassType) String() string {
	if s, ok := mpTThostFtdcBusinessClassType[e]; ok {
		return s[strings.LastIndex(s, "_")+1:]
	}
	return string(e) + "值未定义"
}

// 换汇数据来源类型
type TThostFtdcSwapSourceTypeType byte

const THOST_FTDC_SST_Manual TThostFtdcSwapSourceTypeType = '0' // 手工

const THOST_FTDC_SST_Automatic TThostFtdcSwapSourceTypeType = '1' // 自动生成

var mpTThostFtdcSwapSourceTypeType = map[TThostFtdcSwapSourceTypeType]string{'0': "THOST_FTDC_SST_Manual", '1': "THOST_FTDC_SST_Automatic"}

func (e TThostFtdcSwapSourceTypeType) String() string {
	if s, ok := mpTThostFtdcSwapSourceTypeType[e]; ok {
		return s[strings.LastIndex(s, "_")+1:]
	}
	return string(e) + "值未定义"
}

// 换汇类型类型
type TThostFtdcCurrExDirectionType byte

const THOST_FTDC_CED_Settlement TThostFtdcCurrExDirectionType = '0' // 结汇

const THOST_FTDC_CED_Sale TThostFtdcCurrExDirectionType = '1' // 售汇

var mpTThostFtdcCurrExDirectionType = map[TThostFtdcCurrExDirectionType]string{'0': "THOST_FTDC_CED_Settlement", '1': "THOST_FTDC_CED_Sale"}

func (e TThostFtdcCurrExDirectionType) String() string {
	if s, ok := mpTThostFtdcCurrExDirectionType[e]; ok {
		return s[strings.LastIndex(s, "_")+1:]
	}
	return string(e) + "值未定义"
}

// 申请状态类型
type TThostFtdcCurrencySwapStatusType byte

const THOST_FTDC_CSS_Entry TThostFtdcCurrencySwapStatusType = '1' // 已录入

const THOST_FTDC_CSS_Approve TThostFtdcCurrencySwapStatusType = '2' // 已审核

const THOST_FTDC_CSS_Refuse TThostFtdcCurrencySwapStatusType = '3' // 已拒绝

const THOST_FTDC_CSS_Revoke TThostFtdcCurrencySwapStatusType = '4' // 已撤销

const THOST_FTDC_CSS_Send TThostFtdcCurrencySwapStatusType = '5' // 已发送

const THOST_FTDC_CSS_Success TThostFtdcCurrencySwapStatusType = '6' // 换汇成功

const THOST_FTDC_CSS_Failure TThostFtdcCurrencySwapStatusType = '7' // 换汇失败

var mpTThostFtdcCurrencySwapStatusType = map[TThostFtdcCurrencySwapStatusType]string{'1': "THOST_FTDC_CSS_Entry", '2': "THOST_FTDC_CSS_Approve", '3': "THOST_FTDC_CSS_Refuse", '4': "THOST_FTDC_CSS_Revoke", '5': "THOST_FTDC_CSS_Send", '6': "THOST_FTDC_CSS_Success", '7': "THOST_FTDC_CSS_Failure"}

func (e TThostFtdcCurrencySwapStatusType) String() string {
	if s, ok := mpTThostFtdcCurrencySwapStatusType[e]; ok {
		return s[strings.LastIndex(s, "_")+1:]
	}
	return string(e) + "值未定义"
}

// 凭证号类型
type TThostFtdcCurrExchCertNoType Byte13

// 批次号类型
type TThostFtdcBatchSerialNoType Byte21

// 换汇发送标志类型
type TThostFtdcReqFlagType byte

const THOST_FTDC_REQF_NoSend TThostFtdcReqFlagType = '0' // 未发送

const THOST_FTDC_REQF_SendSuccess TThostFtdcReqFlagType = '1' // 发送成功

const THOST_FTDC_REQF_SendFailed TThostFtdcReqFlagType = '2' // 发送失败

const THOST_FTDC_REQF_WaitReSend TThostFtdcReqFlagType = '3' // 等待重发

var mpTThostFtdcReqFlagType = map[TThostFtdcReqFlagType]string{'0': "THOST_FTDC_REQF_NoSend", '1': "THOST_FTDC_REQF_SendSuccess", '2': "THOST_FTDC_REQF_SendFailed", '3': "THOST_FTDC_REQF_WaitReSend"}

func (e TThostFtdcReqFlagType) String() string {
	if s, ok := mpTThostFtdcReqFlagType[e]; ok {
		return s[strings.LastIndex(s, "_")+1:]
	}
	return string(e) + "值未定义"
}

// 换汇返回成功标志类型
type TThostFtdcResFlagType byte

const THOST_FTDC_RESF_Success TThostFtdcResFlagType = '0' // 成功

const THOST_FTDC_RESF_InsuffiCient TThostFtdcResFlagType = '1' // 账户余额不足

const THOST_FTDC_RESF_UnKnown TThostFtdcResFlagType = '8' // 交易结果未知

var mpTThostFtdcResFlagType = map[TThostFtdcResFlagType]string{'0': "THOST_FTDC_RESF_Success", '1': "THOST_FTDC_RESF_InsuffiCient", '8': "THOST_FTDC_RESF_UnKnown"}

func (e TThostFtdcResFlagType) String() string {
	if s, ok := mpTThostFtdcResFlagType[e]; ok {
		return s[strings.LastIndex(s, "_")+1:]
	}
	return string(e) + "值未定义"
}

// 换汇页面控制类型
type TThostFtdcPageControlType Byte2

// 记录数类型
type TThostFtdcRecordCountType int32

// 换汇需确认信息类型
type TThostFtdcCurrencySwapMemoType Byte101

// 修改状态类型
type TThostFtdcExStatusType byte

const THOST_FTDC_EXS_Before TThostFtdcExStatusType = '0' // 修改前

const THOST_FTDC_EXS_After TThostFtdcExStatusType = '1' // 修改后

var mpTThostFtdcExStatusType = map[TThostFtdcExStatusType]string{'0': "THOST_FTDC_EXS_Before", '1': "THOST_FTDC_EXS_After"}

func (e TThostFtdcExStatusType) String() string {
	if s, ok := mpTThostFtdcExStatusType[e]; ok {
		return s[strings.LastIndex(s, "_")+1:]
	}
	return string(e) + "值未定义"
}

// 开户客户地域类型
type TThostFtdcClientRegionType byte

const THOST_FTDC_CR_Domestic TThostFtdcClientRegionType = '1' // 国内客户

const THOST_FTDC_CR_GMT TThostFtdcClientRegionType = '2' // 港澳台客户

const THOST_FTDC_CR_Foreign TThostFtdcClientRegionType = '3' // 国外客户

var mpTThostFtdcClientRegionType = map[TThostFtdcClientRegionType]string{'1': "THOST_FTDC_CR_Domestic", '2': "THOST_FTDC_CR_GMT", '3': "THOST_FTDC_CR_Foreign"}

func (e TThostFtdcClientRegionType) String() string {
	if s, ok := mpTThostFtdcClientRegionType[e]; ok {
		return s[strings.LastIndex(s, "_")+1:]
	}
	return string(e) + "值未定义"
}

// 工作单位类型
type TThostFtdcWorkPlaceType Byte101

// 经营期限类型
type TThostFtdcBusinessPeriodType Byte21

// 网址类型
type TThostFtdcWebSiteType Byte101

// 统一开户证件类型类型
type TThostFtdcUOAIdCardTypeType Byte3

// 开户模式类型
type TThostFtdcClientModeType Byte3

// 投资者全称类型
type TThostFtdcInvestorFullNameType Byte101

// 境外中介机构ID类型
type TThostFtdcUOABrokerIDType Byte11

// 邮政编码类型
type TThostFtdcUOAZipCodeType Byte11

// 电子邮箱类型
type TThostFtdcUOAEMailType Byte101

// 城市类型
type TThostFtdcOldCityType Byte41

// 法人代表证件号码类型
type TThostFtdcCorporateIdentifiedCardNoType Byte101

// 是否有董事会类型
type TThostFtdcHasBoardType byte

const THOST_FTDC_HB_No TThostFtdcHasBoardType = '0' // 没有

const THOST_FTDC_HB_Yes TThostFtdcHasBoardType = '1' // 有

var mpTThostFtdcHasBoardType = map[TThostFtdcHasBoardType]string{'0': "THOST_FTDC_HB_No", '1': "THOST_FTDC_HB_Yes"}

func (e TThostFtdcHasBoardType) String() string {
	if s, ok := mpTThostFtdcHasBoardType[e]; ok {
		return s[strings.LastIndex(s, "_")+1:]
	}
	return string(e) + "值未定义"
}

// 启动模式类型
type TThostFtdcStartModeType byte

const THOST_FTDC_SM_Normal TThostFtdcStartModeType = '1' // 正常

const THOST_FTDC_SM_Emerge TThostFtdcStartModeType = '2' // 应急

const THOST_FTDC_SM_Restore TThostFtdcStartModeType = '3' // 恢复

var mpTThostFtdcStartModeType = map[TThostFtdcStartModeType]string{'1': "THOST_FTDC_SM_Normal", '2': "THOST_FTDC_SM_Emerge", '3': "THOST_FTDC_SM_Restore"}

func (e TThostFtdcStartModeType) String() string {
	if s, ok := mpTThostFtdcStartModeType[e]; ok {
		return s[strings.LastIndex(s, "_")+1:]
	}
	return string(e) + "值未定义"
}

// 模型类型类型
type TThostFtdcTemplateTypeType byte

const THOST_FTDC_TPT_Full TThostFtdcTemplateTypeType = '1' // 全量

const THOST_FTDC_TPT_Increment TThostFtdcTemplateTypeType = '2' // 增量

const THOST_FTDC_TPT_BackUp TThostFtdcTemplateTypeType = '3' // 备份

var mpTThostFtdcTemplateTypeType = map[TThostFtdcTemplateTypeType]string{'1': "THOST_FTDC_TPT_Full", '2': "THOST_FTDC_TPT_Increment", '3': "THOST_FTDC_TPT_BackUp"}

func (e TThostFtdcTemplateTypeType) String() string {
	if s, ok := mpTThostFtdcTemplateTypeType[e]; ok {
		return s[strings.LastIndex(s, "_")+1:]
	}
	return string(e) + "值未定义"
}

// 登录模式类型
type TThostFtdcLoginModeType byte

const THOST_FTDC_LM_Trade TThostFtdcLoginModeType = '0' // 交易

const THOST_FTDC_LM_Transfer TThostFtdcLoginModeType = '1' // 转账

var mpTThostFtdcLoginModeType = map[TThostFtdcLoginModeType]string{'0': "THOST_FTDC_LM_Trade", '1': "THOST_FTDC_LM_Transfer"}

func (e TThostFtdcLoginModeType) String() string {
	if s, ok := mpTThostFtdcLoginModeType[e]; ok {
		return s[strings.LastIndex(s, "_")+1:]
	}
	return string(e) + "值未定义"
}

// 日历提示类型类型
type TThostFtdcPromptTypeType byte

const THOST_FTDC_CPT_Instrument TThostFtdcPromptTypeType = '1' // 合约上下市

const THOST_FTDC_CPT_Margin TThostFtdcPromptTypeType = '2' // 保证金分段生效

var mpTThostFtdcPromptTypeType = map[TThostFtdcPromptTypeType]string{'1': "THOST_FTDC_CPT_Instrument", '2': "THOST_FTDC_CPT_Margin"}

func (e TThostFtdcPromptTypeType) String() string {
	if s, ok := mpTThostFtdcPromptTypeType[e]; ok {
		return s[strings.LastIndex(s, "_")+1:]
	}
	return string(e) + "值未定义"
}

// 分户管理资产编码类型
type TThostFtdcLedgerManageIDType Byte51

// 投资品种类型
type TThostFtdcInvestVarietyType Byte101

// 账户类别类型
type TThostFtdcBankAccountTypeType Byte2

// 开户银行类型
type TThostFtdcLedgerManageBankType Byte101

// 开户营业部类型
type TThostFtdcCffexDepartmentNameType Byte101

// 营业部代码类型
type TThostFtdcCffexDepartmentCodeType Byte9

// 是否有托管人类型
type TThostFtdcHasTrusteeType byte

const THOST_FTDC_HT_Yes TThostFtdcHasTrusteeType = '1' // 有

const THOST_FTDC_HT_No TThostFtdcHasTrusteeType = '0' // 没有

var mpTThostFtdcHasTrusteeType = map[TThostFtdcHasTrusteeType]string{'1': "THOST_FTDC_HT_Yes", '0': "THOST_FTDC_HT_No"}

func (e TThostFtdcHasTrusteeType) String() string {
	if s, ok := mpTThostFtdcHasTrusteeType[e]; ok {
		return s[strings.LastIndex(s, "_")+1:]
	}
	return string(e) + "值未定义"
}

// 说明类型
type TThostFtdcCSRCMemo1Type Byte41

// 代理资产管理业务的期货公司全称类型
type TThostFtdcAssetmgrCFullNameType Byte101

// 资产管理业务批文号类型
type TThostFtdcAssetmgrApprovalNOType Byte51

// 资产管理业务负责人姓名类型
type TThostFtdcAssetmgrMgrNameType Byte401

// 机构类型类型
type TThostFtdcAmTypeType byte

const THOST_FTDC_AMT_Bank TThostFtdcAmTypeType = '1' // 银行

const THOST_FTDC_AMT_Securities TThostFtdcAmTypeType = '2' // 证券公司

const THOST_FTDC_AMT_Fund TThostFtdcAmTypeType = '3' // 基金公司

const THOST_FTDC_AMT_Insurance TThostFtdcAmTypeType = '4' // 保险公司

const THOST_FTDC_AMT_Trust TThostFtdcAmTypeType = '5' // 信托公司

const THOST_FTDC_AMT_Other TThostFtdcAmTypeType = '9' // 其他

var mpTThostFtdcAmTypeType = map[TThostFtdcAmTypeType]string{'1': "THOST_FTDC_AMT_Bank", '2': "THOST_FTDC_AMT_Securities", '3': "THOST_FTDC_AMT_Fund", '4': "THOST_FTDC_AMT_Insurance", '5': "THOST_FTDC_AMT_Trust", '9': "THOST_FTDC_AMT_Other"}

func (e TThostFtdcAmTypeType) String() string {
	if s, ok := mpTThostFtdcAmTypeType[e]; ok {
		return s[strings.LastIndex(s, "_")+1:]
	}
	return string(e) + "值未定义"
}

// 机构类型类型
type TThostFtdcCSRCAmTypeType Byte5

// 出入金类型类型
type TThostFtdcCSRCFundIOTypeType byte

const THOST_FTDC_CFIOT_FundIO TThostFtdcCSRCFundIOTypeType = '0' // 出入金

const THOST_FTDC_CFIOT_SwapCurrency TThostFtdcCSRCFundIOTypeType = '1' // 银期换汇

var mpTThostFtdcCSRCFundIOTypeType = map[TThostFtdcCSRCFundIOTypeType]string{'0': "THOST_FTDC_CFIOT_FundIO", '1': "THOST_FTDC_CFIOT_SwapCurrency"}

func (e TThostFtdcCSRCFundIOTypeType) String() string {
	if s, ok := mpTThostFtdcCSRCFundIOTypeType[e]; ok {
		return s[strings.LastIndex(s, "_")+1:]
	}
	return string(e) + "值未定义"
}

// 结算账户类型类型
type TThostFtdcCusAccountTypeType byte

const THOST_FTDC_CAT_Futures TThostFtdcCusAccountTypeType = '1' // 期货结算账户

const THOST_FTDC_CAT_AssetmgrFuture TThostFtdcCusAccountTypeType = '2' // 纯期货资管业务下的资管结算账户

const THOST_FTDC_CAT_AssetmgrTrustee TThostFtdcCusAccountTypeType = '3' // 综合类资管业务下的期货资管托管账户

const THOST_FTDC_CAT_AssetmgrTransfer TThostFtdcCusAccountTypeType = '4' // 综合类资管业务下的资金中转账户

var mpTThostFtdcCusAccountTypeType = map[TThostFtdcCusAccountTypeType]string{'1': "THOST_FTDC_CAT_Futures", '2': "THOST_FTDC_CAT_AssetmgrFuture", '3': "THOST_FTDC_CAT_AssetmgrTrustee", '4': "THOST_FTDC_CAT_AssetmgrTransfer"}

func (e TThostFtdcCusAccountTypeType) String() string {
	if s, ok := mpTThostFtdcCusAccountTypeType[e]; ok {
		return s[strings.LastIndex(s, "_")+1:]
	}
	return string(e) + "值未定义"
}

// 国籍类型
type TThostFtdcCSRCNationalType Byte4

// 二级代理ID类型
type TThostFtdcCSRCSecAgentIDType Byte11

// 通知语言类型类型
type TThostFtdcLanguageTypeType byte

const THOST_FTDC_LT_Chinese TThostFtdcLanguageTypeType = '1' // 中文

const THOST_FTDC_LT_English TThostFtdcLanguageTypeType = '2' // 英文

var mpTThostFtdcLanguageTypeType = map[TThostFtdcLanguageTypeType]string{'1': "THOST_FTDC_LT_Chinese", '2': "THOST_FTDC_LT_English"}

func (e TThostFtdcLanguageTypeType) String() string {
	if s, ok := mpTThostFtdcLanguageTypeType[e]; ok {
		return s[strings.LastIndex(s, "_")+1:]
	}
	return string(e) + "值未定义"
}

// 投资账户类型
type TThostFtdcAmAccountType Byte23

// 资产管理客户类型类型
type TThostFtdcAssetmgrClientTypeType byte

const THOST_FTDC_AMCT_Person TThostFtdcAssetmgrClientTypeType = '1' // 个人资管客户

const THOST_FTDC_AMCT_Organ TThostFtdcAssetmgrClientTypeType = '2' // 单位资管客户

const THOST_FTDC_AMCT_SpecialOrgan TThostFtdcAssetmgrClientTypeType = '4' // 特殊单位资管客户

var mpTThostFtdcAssetmgrClientTypeType = map[TThostFtdcAssetmgrClientTypeType]string{'1': "THOST_FTDC_AMCT_Person", '2': "THOST_FTDC_AMCT_Organ", '4': "THOST_FTDC_AMCT_SpecialOrgan"}

func (e TThostFtdcAssetmgrClientTypeType) String() string {
	if s, ok := mpTThostFtdcAssetmgrClientTypeType[e]; ok {
		return s[strings.LastIndex(s, "_")+1:]
	}
	return string(e) + "值未定义"
}

// 投资类型类型
type TThostFtdcAssetmgrTypeType byte

const THOST_FTDC_ASST_Futures TThostFtdcAssetmgrTypeType = '3' // 期货类

const THOST_FTDC_ASST_SpecialOrgan TThostFtdcAssetmgrTypeType = '4' // 综合类

var mpTThostFtdcAssetmgrTypeType = map[TThostFtdcAssetmgrTypeType]string{'3': "THOST_FTDC_ASST_Futures", '4': "THOST_FTDC_ASST_SpecialOrgan"}

func (e TThostFtdcAssetmgrTypeType) String() string {
	if s, ok := mpTThostFtdcAssetmgrTypeType[e]; ok {
		return s[strings.LastIndex(s, "_")+1:]
	}
	return string(e) + "值未定义"
}

// 计量单位类型
type TThostFtdcUOMType Byte11

// 上期所合约生命周期状态类型
type TThostFtdcSHFEInstLifePhaseType Byte3

// 产品类型类型
type TThostFtdcSHFEProductClassType Byte11

// 价格小数位类型
type TThostFtdcPriceDecimalType Byte2

// 平值期权标志类型
type TThostFtdcInTheMoneyFlagType Byte2

// 合约比较类型类型
type TThostFtdcCheckInstrTypeType byte

const THOST_FTDC_CIT_HasExch TThostFtdcCheckInstrTypeType = '0' // 合约交易所不存在

const THOST_FTDC_CIT_HasATP TThostFtdcCheckInstrTypeType = '1' // 合约本系统不存在

const THOST_FTDC_CIT_HasDiff TThostFtdcCheckInstrTypeType = '2' // 合约比较不一致

var mpTThostFtdcCheckInstrTypeType = map[TThostFtdcCheckInstrTypeType]string{'0': "THOST_FTDC_CIT_HasExch", '1': "THOST_FTDC_CIT_HasATP", '2': "THOST_FTDC_CIT_HasDiff"}

func (e TThostFtdcCheckInstrTypeType) String() string {
	if s, ok := mpTThostFtdcCheckInstrTypeType[e]; ok {
		return s[strings.LastIndex(s, "_")+1:]
	}
	return string(e) + "值未定义"
}

// 交割类型类型
type TThostFtdcDeliveryTypeType byte

const THOST_FTDC_DT_HandDeliv TThostFtdcDeliveryTypeType = '1' // 手工交割

const THOST_FTDC_DT_PersonDeliv TThostFtdcDeliveryTypeType = '2' // 到期交割

var mpTThostFtdcDeliveryTypeType = map[TThostFtdcDeliveryTypeType]string{'1': "THOST_FTDC_DT_HandDeliv", '2': "THOST_FTDC_DT_PersonDeliv"}

func (e TThostFtdcDeliveryTypeType) String() string {
	if s, ok := mpTThostFtdcDeliveryTypeType[e]; ok {
		return s[strings.LastIndex(s, "_")+1:]
	}
	return string(e) + "值未定义"
}

// 资金类型
type TThostFtdcBigMoneyType Double

// 大额单边保证金算法类型
type TThostFtdcMaxMarginSideAlgorithmType byte

const THOST_FTDC_MMSA_NO TThostFtdcMaxMarginSideAlgorithmType = '0' // 不使用大额单边保证金算法

const THOST_FTDC_MMSA_YES TThostFtdcMaxMarginSideAlgorithmType = '1' // 使用大额单边保证金算法

var mpTThostFtdcMaxMarginSideAlgorithmType = map[TThostFtdcMaxMarginSideAlgorithmType]string{'0': "THOST_FTDC_MMSA_NO", '1': "THOST_FTDC_MMSA_YES"}

func (e TThostFtdcMaxMarginSideAlgorithmType) String() string {
	if s, ok := mpTThostFtdcMaxMarginSideAlgorithmType[e]; ok {
		return s[strings.LastIndex(s, "_")+1:]
	}
	return string(e) + "值未定义"
}

// 资产管理客户类型类型
type TThostFtdcDAClientTypeType byte

const THOST_FTDC_CACT_Person TThostFtdcDAClientTypeType = '0' // 自然人

const THOST_FTDC_CACT_Company TThostFtdcDAClientTypeType = '1' // 法人

const THOST_FTDC_CACT_Other TThostFtdcDAClientTypeType = '2' // 其他

var mpTThostFtdcDAClientTypeType = map[TThostFtdcDAClientTypeType]string{'0': "THOST_FTDC_CACT_Person", '1': "THOST_FTDC_CACT_Company", '2': "THOST_FTDC_CACT_Other"}

func (e TThostFtdcDAClientTypeType) String() string {
	if s, ok := mpTThostFtdcDAClientTypeType[e]; ok {
		return s[strings.LastIndex(s, "_")+1:]
	}
	return string(e) + "值未定义"
}

// 套利合约代码类型
type TThostFtdcCombinInstrIDType Byte61

// 各腿结算价类型
type TThostFtdcCombinSettlePriceType Byte61

// 优先级类型
type TThostFtdcDCEPriorityType int32

// 成交组号类型
type TThostFtdcTradeGroupIDType int32

// 是否校验开户可用资金类型
type TThostFtdcIsCheckPrepaType int32

// 投资类型类型
type TThostFtdcUOAAssetmgrTypeType byte

const THOST_FTDC_UOAAT_Futures TThostFtdcUOAAssetmgrTypeType = '1' // 期货类

const THOST_FTDC_UOAAT_SpecialOrgan TThostFtdcUOAAssetmgrTypeType = '2' // 综合类

var mpTThostFtdcUOAAssetmgrTypeType = map[TThostFtdcUOAAssetmgrTypeType]string{'1': "THOST_FTDC_UOAAT_Futures", '2': "THOST_FTDC_UOAAT_SpecialOrgan"}

func (e TThostFtdcUOAAssetmgrTypeType) String() string {
	if s, ok := mpTThostFtdcUOAAssetmgrTypeType[e]; ok {
		return s[strings.LastIndex(s, "_")+1:]
	}
	return string(e) + "值未定义"
}

// 买卖方向类型
type TThostFtdcDirectionEnType byte

const THOST_FTDC_DEN_Buy TThostFtdcDirectionEnType = '0' // Buy

const THOST_FTDC_DEN_Sell TThostFtdcDirectionEnType = '1' // Sell

var mpTThostFtdcDirectionEnType = map[TThostFtdcDirectionEnType]string{'0': "THOST_FTDC_DEN_Buy", '1': "THOST_FTDC_DEN_Sell"}

func (e TThostFtdcDirectionEnType) String() string {
	if s, ok := mpTThostFtdcDirectionEnType[e]; ok {
		return s[strings.LastIndex(s, "_")+1:]
	}
	return string(e) + "值未定义"
}

// 开平标志类型
type TThostFtdcOffsetFlagEnType byte

const THOST_FTDC_OFEN_Open TThostFtdcOffsetFlagEnType = '0' // Position Opening

const THOST_FTDC_OFEN_Close TThostFtdcOffsetFlagEnType = '1' // Position Close

const THOST_FTDC_OFEN_ForceClose TThostFtdcOffsetFlagEnType = '2' // Forced Liquidation

const THOST_FTDC_OFEN_CloseToday TThostFtdcOffsetFlagEnType = '3' // Close Today

const THOST_FTDC_OFEN_CloseYesterday TThostFtdcOffsetFlagEnType = '4' // Close Prev.

const THOST_FTDC_OFEN_ForceOff TThostFtdcOffsetFlagEnType = '5' // Forced Reduction

const THOST_FTDC_OFEN_LocalForceClose TThostFtdcOffsetFlagEnType = '6' // Local Forced Liquidation

var mpTThostFtdcOffsetFlagEnType = map[TThostFtdcOffsetFlagEnType]string{'0': "THOST_FTDC_OFEN_Open", '1': "THOST_FTDC_OFEN_Close", '2': "THOST_FTDC_OFEN_ForceClose", '3': "THOST_FTDC_OFEN_CloseToday", '4': "THOST_FTDC_OFEN_CloseYesterday", '5': "THOST_FTDC_OFEN_ForceOff", '6': "THOST_FTDC_OFEN_LocalForceClose"}

func (e TThostFtdcOffsetFlagEnType) String() string {
	if s, ok := mpTThostFtdcOffsetFlagEnType[e]; ok {
		return s[strings.LastIndex(s, "_")+1:]
	}
	return string(e) + "值未定义"
}

// 投机套保标志类型
type TThostFtdcHedgeFlagEnType byte

const THOST_FTDC_HFEN_Speculation TThostFtdcHedgeFlagEnType = '1' // Speculation

const THOST_FTDC_HFEN_Arbitrage TThostFtdcHedgeFlagEnType = '2' // Arbitrage

const THOST_FTDC_HFEN_Hedge TThostFtdcHedgeFlagEnType = '3' // Hedge

var mpTThostFtdcHedgeFlagEnType = map[TThostFtdcHedgeFlagEnType]string{'1': "THOST_FTDC_HFEN_Speculation", '2': "THOST_FTDC_HFEN_Arbitrage", '3': "THOST_FTDC_HFEN_Hedge"}

func (e TThostFtdcHedgeFlagEnType) String() string {
	if s, ok := mpTThostFtdcHedgeFlagEnType[e]; ok {
		return s[strings.LastIndex(s, "_")+1:]
	}
	return string(e) + "值未定义"
}

// 出入金类型类型
type TThostFtdcFundIOTypeEnType byte

const THOST_FTDC_FIOTEN_FundIO TThostFtdcFundIOTypeEnType = '1' // Deposit/Withdrawal

const THOST_FTDC_FIOTEN_Transfer TThostFtdcFundIOTypeEnType = '2' // Bank-Futures Transfer

const THOST_FTDC_FIOTEN_SwapCurrency TThostFtdcFundIOTypeEnType = '3' // Bank-Futures FX Exchange

var mpTThostFtdcFundIOTypeEnType = map[TThostFtdcFundIOTypeEnType]string{'1': "THOST_FTDC_FIOTEN_FundIO", '2': "THOST_FTDC_FIOTEN_Transfer", '3': "THOST_FTDC_FIOTEN_SwapCurrency"}

func (e TThostFtdcFundIOTypeEnType) String() string {
	if s, ok := mpTThostFtdcFundIOTypeEnType[e]; ok {
		return s[strings.LastIndex(s, "_")+1:]
	}
	return string(e) + "值未定义"
}

// 资金类型类型
type TThostFtdcFundTypeEnType byte

const THOST_FTDC_FTEN_Deposite TThostFtdcFundTypeEnType = '1' // Bank Deposit

const THOST_FTDC_FTEN_ItemFund TThostFtdcFundTypeEnType = '2' // Payment/Fee

const THOST_FTDC_FTEN_Company TThostFtdcFundTypeEnType = '3' // Brokerage Adj

const THOST_FTDC_FTEN_InnerTransfer TThostFtdcFundTypeEnType = '4' // Internal Transfer

var mpTThostFtdcFundTypeEnType = map[TThostFtdcFundTypeEnType]string{'1': "THOST_FTDC_FTEN_Deposite", '2': "THOST_FTDC_FTEN_ItemFund", '3': "THOST_FTDC_FTEN_Company", '4': "THOST_FTDC_FTEN_InnerTransfer"}

func (e TThostFtdcFundTypeEnType) String() string {
	if s, ok := mpTThostFtdcFundTypeEnType[e]; ok {
		return s[strings.LastIndex(s, "_")+1:]
	}
	return string(e) + "值未定义"
}

// 出入金方向类型
type TThostFtdcFundDirectionEnType byte

const THOST_FTDC_FDEN_In TThostFtdcFundDirectionEnType = '1' // Deposit

const THOST_FTDC_FDEN_Out TThostFtdcFundDirectionEnType = '2' // Withdrawal

var mpTThostFtdcFundDirectionEnType = map[TThostFtdcFundDirectionEnType]string{'1': "THOST_FTDC_FDEN_In", '2': "THOST_FTDC_FDEN_Out"}

func (e TThostFtdcFundDirectionEnType) String() string {
	if s, ok := mpTThostFtdcFundDirectionEnType[e]; ok {
		return s[strings.LastIndex(s, "_")+1:]
	}
	return string(e) + "值未定义"
}

// 货币质押方向类型
type TThostFtdcFundMortDirectionEnType byte

const THOST_FTDC_FMDEN_In TThostFtdcFundMortDirectionEnType = '1' // Pledge

const THOST_FTDC_FMDEN_Out TThostFtdcFundMortDirectionEnType = '2' // Redemption

var mpTThostFtdcFundMortDirectionEnType = map[TThostFtdcFundMortDirectionEnType]string{'1': "THOST_FTDC_FMDEN_In", '2': "THOST_FTDC_FMDEN_Out"}

func (e TThostFtdcFundMortDirectionEnType) String() string {
	if s, ok := mpTThostFtdcFundMortDirectionEnType[e]; ok {
		return s[strings.LastIndex(s, "_")+1:]
	}
	return string(e) + "值未定义"
}

// 换汇业务种类类型
type TThostFtdcSwapBusinessTypeType Byte3

// 期权类型类型
type TThostFtdcOptionsTypeType byte

const THOST_FTDC_CP_CallOptions TThostFtdcOptionsTypeType = '1' // 看涨

const THOST_FTDC_CP_PutOptions TThostFtdcOptionsTypeType = '2' // 看跌

var mpTThostFtdcOptionsTypeType = map[TThostFtdcOptionsTypeType]string{'1': "THOST_FTDC_CP_CallOptions", '2': "THOST_FTDC_CP_PutOptions"}

func (e TThostFtdcOptionsTypeType) String() string {
	if s, ok := mpTThostFtdcOptionsTypeType[e]; ok {
		return s[strings.LastIndex(s, "_")+1:]
	}
	return string(e) + "值未定义"
}

// 执行方式类型
type TThostFtdcStrikeModeType byte

const THOST_FTDC_STM_Continental TThostFtdcStrikeModeType = '0' // 欧式

const THOST_FTDC_STM_American TThostFtdcStrikeModeType = '1' // 美式

const THOST_FTDC_STM_Bermuda TThostFtdcStrikeModeType = '2' // 百慕大

var mpTThostFtdcStrikeModeType = map[TThostFtdcStrikeModeType]string{'0': "THOST_FTDC_STM_Continental", '1': "THOST_FTDC_STM_American", '2': "THOST_FTDC_STM_Bermuda"}

func (e TThostFtdcStrikeModeType) String() string {
	if s, ok := mpTThostFtdcStrikeModeType[e]; ok {
		return s[strings.LastIndex(s, "_")+1:]
	}
	return string(e) + "值未定义"
}

// 执行类型类型
type TThostFtdcStrikeTypeType byte

const THOST_FTDC_STT_Hedge TThostFtdcStrikeTypeType = '0' // 自身对冲

const THOST_FTDC_STT_Match TThostFtdcStrikeTypeType = '1' // 匹配执行

var mpTThostFtdcStrikeTypeType = map[TThostFtdcStrikeTypeType]string{'0': "THOST_FTDC_STT_Hedge", '1': "THOST_FTDC_STT_Match"}

func (e TThostFtdcStrikeTypeType) String() string {
	if s, ok := mpTThostFtdcStrikeTypeType[e]; ok {
		return s[strings.LastIndex(s, "_")+1:]
	}
	return string(e) + "值未定义"
}

// 中金所期权放弃执行申请类型类型
type TThostFtdcApplyTypeType byte

const THOST_FTDC_APPT_NotStrikeNum TThostFtdcApplyTypeType = '4' // 不执行数量

var mpTThostFtdcApplyTypeType = map[TThostFtdcApplyTypeType]string{'4': "THOST_FTDC_APPT_NotStrikeNum"}

func (e TThostFtdcApplyTypeType) String() string {
	if s, ok := mpTThostFtdcApplyTypeType[e]; ok {
		return s[strings.LastIndex(s, "_")+1:]
	}
	return string(e) + "值未定义"
}

// 放弃执行申请数据来源类型
type TThostFtdcGiveUpDataSourceType byte

const THOST_FTDC_GUDS_Gen TThostFtdcGiveUpDataSourceType = '0' // 系统生成

const THOST_FTDC_GUDS_Hand TThostFtdcGiveUpDataSourceType = '1' // 手工添加

var mpTThostFtdcGiveUpDataSourceType = map[TThostFtdcGiveUpDataSourceType]string{'0': "THOST_FTDC_GUDS_Gen", '1': "THOST_FTDC_GUDS_Hand"}

func (e TThostFtdcGiveUpDataSourceType) String() string {
	if s, ok := mpTThostFtdcGiveUpDataSourceType[e]; ok {
		return s[strings.LastIndex(s, "_")+1:]
	}
	return string(e) + "值未定义"
}

// 执行宣告系统编号类型
type TThostFtdcExecOrderSysIDType Byte21

// 执行结果类型
type TThostFtdcExecResultType byte

const THOST_FTDC_OER_NoExec TThostFtdcExecResultType = 'n' // 没有执行

const THOST_FTDC_OER_Canceled TThostFtdcExecResultType = 'c' // 已经取消

const THOST_FTDC_OER_OK TThostFtdcExecResultType = '0' // 执行成功

const THOST_FTDC_OER_NoPosition TThostFtdcExecResultType = '1' // 期权持仓不够

const THOST_FTDC_OER_NoDeposit TThostFtdcExecResultType = '2' // 资金不够

const THOST_FTDC_OER_NoParticipant TThostFtdcExecResultType = '3' // 会员不存在

const THOST_FTDC_OER_NoClient TThostFtdcExecResultType = '4' // 客户不存在

const THOST_FTDC_OER_NoInstrument TThostFtdcExecResultType = '6' // 合约不存在

const THOST_FTDC_OER_NoRight TThostFtdcExecResultType = '7' // 没有执行权限

const THOST_FTDC_OER_InvalidVolume TThostFtdcExecResultType = '8' // 不合理的数量

const THOST_FTDC_OER_NoEnoughHistoryTrade TThostFtdcExecResultType = '9' // 没有足够的历史成交

const THOST_FTDC_OER_Unknown TThostFtdcExecResultType = 'a' // 未知

var mpTThostFtdcExecResultType = map[TThostFtdcExecResultType]string{'n': "THOST_FTDC_OER_NoExec", 'c': "THOST_FTDC_OER_Canceled", '0': "THOST_FTDC_OER_OK", '1': "THOST_FTDC_OER_NoPosition", '2': "THOST_FTDC_OER_NoDeposit", '3': "THOST_FTDC_OER_NoParticipant", '4': "THOST_FTDC_OER_NoClient", '6': "THOST_FTDC_OER_NoInstrument", '7': "THOST_FTDC_OER_NoRight", '8': "THOST_FTDC_OER_InvalidVolume", '9': "THOST_FTDC_OER_NoEnoughHistoryTrade", 'a': "THOST_FTDC_OER_Unknown"}

func (e TThostFtdcExecResultType) String() string {
	if s, ok := mpTThostFtdcExecResultType[e]; ok {
		return s[strings.LastIndex(s, "_")+1:]
	}
	return string(e) + "值未定义"
}

// 执行序号类型
type TThostFtdcStrikeSequenceType int32

// 执行时间类型
type TThostFtdcStrikeTimeType Byte13

// 组合类型类型
type TThostFtdcCombinationTypeType byte

const THOST_FTDC_COMBT_Future TThostFtdcCombinationTypeType = '0' // 期货组合

const THOST_FTDC_COMBT_BUL TThostFtdcCombinationTypeType = '1' // 垂直价差BUL

const THOST_FTDC_COMBT_BER TThostFtdcCombinationTypeType = '2' // 垂直价差BER

const THOST_FTDC_COMBT_STD TThostFtdcCombinationTypeType = '3' // 跨式组合

const THOST_FTDC_COMBT_STG TThostFtdcCombinationTypeType = '4' // 宽跨式组合

const THOST_FTDC_COMBT_PRT TThostFtdcCombinationTypeType = '5' // 备兑组合

const THOST_FTDC_COMBT_CAS TThostFtdcCombinationTypeType = '6' // 时间价差组合

const THOST_FTDC_COMBT_OPL TThostFtdcCombinationTypeType = '7' // 期权对锁组合

const THOST_FTDC_COMBT_BFO TThostFtdcCombinationTypeType = '8' // 买备兑组合

const THOST_FTDC_COMBT_BLS TThostFtdcCombinationTypeType = '9' // 买入期权垂直价差组合

const THOST_FTDC_COMBT_BES TThostFtdcCombinationTypeType = 'a' // 卖出期权垂直价差组合

var mpTThostFtdcCombinationTypeType = map[TThostFtdcCombinationTypeType]string{'0': "THOST_FTDC_COMBT_Future", '1': "THOST_FTDC_COMBT_BUL", '2': "THOST_FTDC_COMBT_BER", '3': "THOST_FTDC_COMBT_STD", '4': "THOST_FTDC_COMBT_STG", '5': "THOST_FTDC_COMBT_PRT", '6': "THOST_FTDC_COMBT_CAS", '7': "THOST_FTDC_COMBT_OPL", '8': "THOST_FTDC_COMBT_BFO", '9': "THOST_FTDC_COMBT_BLS", 'a': "THOST_FTDC_COMBT_BES"}

func (e TThostFtdcCombinationTypeType) String() string {
	if s, ok := mpTThostFtdcCombinationTypeType[e]; ok {
		return s[strings.LastIndex(s, "_")+1:]
	}
	return string(e) + "值未定义"
}

// 组合类型类型
type TThostFtdcDceCombinationTypeType byte

const THOST_FTDC_DCECOMBT_SPL TThostFtdcDceCombinationTypeType = '0' // 期货对锁组合

const THOST_FTDC_DCECOMBT_OPL TThostFtdcDceCombinationTypeType = '1' // 期权对锁组合

const THOST_FTDC_DCECOMBT_SP TThostFtdcDceCombinationTypeType = '2' // 期货跨期组合

const THOST_FTDC_DCECOMBT_SPC TThostFtdcDceCombinationTypeType = '3' // 期货跨品种组合

const THOST_FTDC_DCECOMBT_BLS TThostFtdcDceCombinationTypeType = '4' // 买入期权垂直价差组合

const THOST_FTDC_DCECOMBT_BES TThostFtdcDceCombinationTypeType = '5' // 卖出期权垂直价差组合

const THOST_FTDC_DCECOMBT_CAS TThostFtdcDceCombinationTypeType = '6' // 期权日历价差组合

const THOST_FTDC_DCECOMBT_STD TThostFtdcDceCombinationTypeType = '7' // 期权跨式组合

const THOST_FTDC_DCECOMBT_STG TThostFtdcDceCombinationTypeType = '8' // 期权宽跨式组合

const THOST_FTDC_DCECOMBT_BFO TThostFtdcDceCombinationTypeType = '9' // 买入期货期权组合

const THOST_FTDC_DCECOMBT_SFO TThostFtdcDceCombinationTypeType = 'a' // 卖出期货期权组合

var mpTThostFtdcDceCombinationTypeType = map[TThostFtdcDceCombinationTypeType]string{'0': "THOST_FTDC_DCECOMBT_SPL", '1': "THOST_FTDC_DCECOMBT_OPL", '2': "THOST_FTDC_DCECOMBT_SP", '3': "THOST_FTDC_DCECOMBT_SPC", '4': "THOST_FTDC_DCECOMBT_BLS", '5': "THOST_FTDC_DCECOMBT_BES", '6': "THOST_FTDC_DCECOMBT_CAS", '7': "THOST_FTDC_DCECOMBT_STD", '8': "THOST_FTDC_DCECOMBT_STG", '9': "THOST_FTDC_DCECOMBT_BFO", 'a': "THOST_FTDC_DCECOMBT_SFO"}

func (e TThostFtdcDceCombinationTypeType) String() string {
	if s, ok := mpTThostFtdcDceCombinationTypeType[e]; ok {
		return s[strings.LastIndex(s, "_")+1:]
	}
	return string(e) + "值未定义"
}

// 期权权利金价格类型类型
type TThostFtdcOptionRoyaltyPriceTypeType byte

const THOST_FTDC_ORPT_PreSettlementPrice TThostFtdcOptionRoyaltyPriceTypeType = '1' // 昨结算价

const THOST_FTDC_ORPT_OpenPrice TThostFtdcOptionRoyaltyPriceTypeType = '4' // 开仓价

const THOST_FTDC_ORPT_MaxPreSettlementPrice TThostFtdcOptionRoyaltyPriceTypeType = '5' // 最新价与昨结算价较大值

var mpTThostFtdcOptionRoyaltyPriceTypeType = map[TThostFtdcOptionRoyaltyPriceTypeType]string{'1': "THOST_FTDC_ORPT_PreSettlementPrice", '4': "THOST_FTDC_ORPT_OpenPrice", '5': "THOST_FTDC_ORPT_MaxPreSettlementPrice"}

func (e TThostFtdcOptionRoyaltyPriceTypeType) String() string {
	if s, ok := mpTThostFtdcOptionRoyaltyPriceTypeType[e]; ok {
		return s[strings.LastIndex(s, "_")+1:]
	}
	return string(e) + "值未定义"
}

// 权益算法类型
type TThostFtdcBalanceAlgorithmType byte

const THOST_FTDC_BLAG_Default TThostFtdcBalanceAlgorithmType = '1' // 不计算期权市值盈亏

const THOST_FTDC_BLAG_IncludeOptValLost TThostFtdcBalanceAlgorithmType = '2' // 计算期权市值亏损

var mpTThostFtdcBalanceAlgorithmType = map[TThostFtdcBalanceAlgorithmType]string{'1': "THOST_FTDC_BLAG_Default", '2': "THOST_FTDC_BLAG_IncludeOptValLost"}

func (e TThostFtdcBalanceAlgorithmType) String() string {
	if s, ok := mpTThostFtdcBalanceAlgorithmType[e]; ok {
		return s[strings.LastIndex(s, "_")+1:]
	}
	return string(e) + "值未定义"
}

// 执行类型类型
type TThostFtdcActionTypeType byte

const THOST_FTDC_ACTP_Exec TThostFtdcActionTypeType = '1' // 执行

const THOST_FTDC_ACTP_Abandon TThostFtdcActionTypeType = '2' // 放弃

var mpTThostFtdcActionTypeType = map[TThostFtdcActionTypeType]string{'1': "THOST_FTDC_ACTP_Exec", '2': "THOST_FTDC_ACTP_Abandon"}

func (e TThostFtdcActionTypeType) String() string {
	if s, ok := mpTThostFtdcActionTypeType[e]; ok {
		return s[strings.LastIndex(s, "_")+1:]
	}
	return string(e) + "值未定义"
}

// 询价状态类型
type TThostFtdcForQuoteStatusType byte

const THOST_FTDC_FQST_Submitted TThostFtdcForQuoteStatusType = 'a' // 已经提交

const THOST_FTDC_FQST_Accepted TThostFtdcForQuoteStatusType = 'b' // 已经接受

const THOST_FTDC_FQST_Rejected TThostFtdcForQuoteStatusType = 'c' // 已经被拒绝

var mpTThostFtdcForQuoteStatusType = map[TThostFtdcForQuoteStatusType]string{'a': "THOST_FTDC_FQST_Submitted", 'b': "THOST_FTDC_FQST_Accepted", 'c': "THOST_FTDC_FQST_Rejected"}

func (e TThostFtdcForQuoteStatusType) String() string {
	if s, ok := mpTThostFtdcForQuoteStatusType[e]; ok {
		return s[strings.LastIndex(s, "_")+1:]
	}
	return string(e) + "值未定义"
}

// 取值方式类型
type TThostFtdcValueMethodType byte

const THOST_FTDC_VM_Absolute TThostFtdcValueMethodType = '0' // 按绝对值

const THOST_FTDC_VM_Ratio TThostFtdcValueMethodType = '1' // 按比率

var mpTThostFtdcValueMethodType = map[TThostFtdcValueMethodType]string{'0': "THOST_FTDC_VM_Absolute", '1': "THOST_FTDC_VM_Ratio"}

func (e TThostFtdcValueMethodType) String() string {
	if s, ok := mpTThostFtdcValueMethodType[e]; ok {
		return s[strings.LastIndex(s, "_")+1:]
	}
	return string(e) + "值未定义"
}

// 期权行权后是否保留期货头寸的标记类型
type TThostFtdcExecOrderPositionFlagType byte

const THOST_FTDC_EOPF_Reserve TThostFtdcExecOrderPositionFlagType = '0' // 保留

const THOST_FTDC_EOPF_UnReserve TThostFtdcExecOrderPositionFlagType = '1' // 不保留

var mpTThostFtdcExecOrderPositionFlagType = map[TThostFtdcExecOrderPositionFlagType]string{'0': "THOST_FTDC_EOPF_Reserve", '1': "THOST_FTDC_EOPF_UnReserve"}

func (e TThostFtdcExecOrderPositionFlagType) String() string {
	if s, ok := mpTThostFtdcExecOrderPositionFlagType[e]; ok {
		return s[strings.LastIndex(s, "_")+1:]
	}
	return string(e) + "值未定义"
}

// 期权行权后生成的头寸是否自动平仓类型
type TThostFtdcExecOrderCloseFlagType byte

const THOST_FTDC_EOCF_AutoClose TThostFtdcExecOrderCloseFlagType = '0' // 自动平仓

const THOST_FTDC_EOCF_NotToClose TThostFtdcExecOrderCloseFlagType = '1' // 免于自动平仓

var mpTThostFtdcExecOrderCloseFlagType = map[TThostFtdcExecOrderCloseFlagType]string{'0': "THOST_FTDC_EOCF_AutoClose", '1': "THOST_FTDC_EOCF_NotToClose"}

func (e TThostFtdcExecOrderCloseFlagType) String() string {
	if s, ok := mpTThostFtdcExecOrderCloseFlagType[e]; ok {
		return s[strings.LastIndex(s, "_")+1:]
	}
	return string(e) + "值未定义"
}

// 产品类型类型
type TThostFtdcProductTypeType byte

const THOST_FTDC_PTE_Futures TThostFtdcProductTypeType = '1' // 期货

const THOST_FTDC_PTE_Options TThostFtdcProductTypeType = '2' // 期权

var mpTThostFtdcProductTypeType = map[TThostFtdcProductTypeType]string{'1': "THOST_FTDC_PTE_Futures", '2': "THOST_FTDC_PTE_Options"}

func (e TThostFtdcProductTypeType) String() string {
	if s, ok := mpTThostFtdcProductTypeType[e]; ok {
		return s[strings.LastIndex(s, "_")+1:]
	}
	return string(e) + "值未定义"
}

// 郑商所结算文件名类型
type TThostFtdcCZCEUploadFileNameType byte

const THOST_FTDC_CUFN_CUFN_O TThostFtdcCZCEUploadFileNameType = 'O' // ^\d{8}_zz_\d{4}

const THOST_FTDC_CUFN_CUFN_T TThostFtdcCZCEUploadFileNameType = 'T' // ^\d{8}成交表

const THOST_FTDC_CUFN_CUFN_P TThostFtdcCZCEUploadFileNameType = 'P' // ^\d{8}单腿持仓表new

const THOST_FTDC_CUFN_CUFN_N TThostFtdcCZCEUploadFileNameType = 'N' // ^\d{8}非平仓了结表

const THOST_FTDC_CUFN_CUFN_L TThostFtdcCZCEUploadFileNameType = 'L' // ^\d{8}平仓表

const THOST_FTDC_CUFN_CUFN_F TThostFtdcCZCEUploadFileNameType = 'F' // ^\d{8}资金表

const THOST_FTDC_CUFN_CUFN_C TThostFtdcCZCEUploadFileNameType = 'C' // ^\d{8}组合持仓表

const THOST_FTDC_CUFN_CUFN_M TThostFtdcCZCEUploadFileNameType = 'M' // ^\d{8}保证金参数表

var mpTThostFtdcCZCEUploadFileNameType = map[TThostFtdcCZCEUploadFileNameType]string{'O': "THOST_FTDC_CUFN_CUFN_O", 'T': "THOST_FTDC_CUFN_CUFN_T", 'P': "THOST_FTDC_CUFN_CUFN_P", 'N': "THOST_FTDC_CUFN_CUFN_N", 'L': "THOST_FTDC_CUFN_CUFN_L", 'F': "THOST_FTDC_CUFN_CUFN_F", 'C': "THOST_FTDC_CUFN_CUFN_C", 'M': "THOST_FTDC_CUFN_CUFN_M"}

func (e TThostFtdcCZCEUploadFileNameType) String() string {
	if s, ok := mpTThostFtdcCZCEUploadFileNameType[e]; ok {
		return s[strings.LastIndex(s, "_")+1:]
	}
	return string(e) + "值未定义"
}

// 大商所结算文件名类型
type TThostFtdcDCEUploadFileNameType byte

const THOST_FTDC_DUFN_DUFN_O TThostFtdcDCEUploadFileNameType = 'O' // ^\d{8}_dl_\d{3}

const THOST_FTDC_DUFN_DUFN_T TThostFtdcDCEUploadFileNameType = 'T' // ^\d{8}_成交表

const THOST_FTDC_DUFN_DUFN_P TThostFtdcDCEUploadFileNameType = 'P' // ^\d{8}_持仓表

const THOST_FTDC_DUFN_DUFN_F TThostFtdcDCEUploadFileNameType = 'F' // ^\d{8}_资金结算表

const THOST_FTDC_DUFN_DUFN_C TThostFtdcDCEUploadFileNameType = 'C' // ^\d{8}_优惠组合持仓明细表

const THOST_FTDC_DUFN_DUFN_D TThostFtdcDCEUploadFileNameType = 'D' // ^\d{8}_持仓明细表

const THOST_FTDC_DUFN_DUFN_M TThostFtdcDCEUploadFileNameType = 'M' // ^\d{8}_保证金参数表

const THOST_FTDC_DUFN_DUFN_S TThostFtdcDCEUploadFileNameType = 'S' // ^\d{8}_期权执行表

var mpTThostFtdcDCEUploadFileNameType = map[TThostFtdcDCEUploadFileNameType]string{'O': "THOST_FTDC_DUFN_DUFN_O", 'T': "THOST_FTDC_DUFN_DUFN_T", 'P': "THOST_FTDC_DUFN_DUFN_P", 'F': "THOST_FTDC_DUFN_DUFN_F", 'C': "THOST_FTDC_DUFN_DUFN_C", 'D': "THOST_FTDC_DUFN_DUFN_D", 'M': "THOST_FTDC_DUFN_DUFN_M", 'S': "THOST_FTDC_DUFN_DUFN_S"}

func (e TThostFtdcDCEUploadFileNameType) String() string {
	if s, ok := mpTThostFtdcDCEUploadFileNameType[e]; ok {
		return s[strings.LastIndex(s, "_")+1:]
	}
	return string(e) + "值未定义"
}

// 上期所结算文件名类型
type TThostFtdcSHFEUploadFileNameType byte

const THOST_FTDC_SUFN_SUFN_O TThostFtdcSHFEUploadFileNameType = 'O' // ^\d{4}_\d{8}_\d{8}_DailyFundChg

const THOST_FTDC_SUFN_SUFN_T TThostFtdcSHFEUploadFileNameType = 'T' // ^\d{4}_\d{8}_\d{8}_Trade

const THOST_FTDC_SUFN_SUFN_P TThostFtdcSHFEUploadFileNameType = 'P' // ^\d{4}_\d{8}_\d{8}_SettlementDetail

const THOST_FTDC_SUFN_SUFN_F TThostFtdcSHFEUploadFileNameType = 'F' // ^\d{4}_\d{8}_\d{8}_Capital

var mpTThostFtdcSHFEUploadFileNameType = map[TThostFtdcSHFEUploadFileNameType]string{'O': "THOST_FTDC_SUFN_SUFN_O", 'T': "THOST_FTDC_SUFN_SUFN_T", 'P': "THOST_FTDC_SUFN_SUFN_P", 'F': "THOST_FTDC_SUFN_SUFN_F"}

func (e TThostFtdcSHFEUploadFileNameType) String() string {
	if s, ok := mpTThostFtdcSHFEUploadFileNameType[e]; ok {
		return s[strings.LastIndex(s, "_")+1:]
	}
	return string(e) + "值未定义"
}

// 中金所结算文件名类型
type TThostFtdcCFFEXUploadFileNameType byte

const THOST_FTDC_CFUFN_SUFN_T TThostFtdcCFFEXUploadFileNameType = 'T' // ^\d{4}_SG\d{1}_\d{8}_\d{1}_Trade

const THOST_FTDC_CFUFN_SUFN_P TThostFtdcCFFEXUploadFileNameType = 'P' // ^\d{4}_SG\d{1}_\d{8}_\d{1}_SettlementDetail

const THOST_FTDC_CFUFN_SUFN_F TThostFtdcCFFEXUploadFileNameType = 'F' // ^\d{4}_SG\d{1}_\d{8}_\d{1}_Capital

const THOST_FTDC_CFUFN_SUFN_S TThostFtdcCFFEXUploadFileNameType = 'S' // ^\d{4}_SG\d{1}_\d{8}_\d{1}_OptionExec

var mpTThostFtdcCFFEXUploadFileNameType = map[TThostFtdcCFFEXUploadFileNameType]string{'T': "THOST_FTDC_CFUFN_SUFN_T", 'P': "THOST_FTDC_CFUFN_SUFN_P", 'F': "THOST_FTDC_CFUFN_SUFN_F", 'S': "THOST_FTDC_CFUFN_SUFN_S"}

func (e TThostFtdcCFFEXUploadFileNameType) String() string {
	if s, ok := mpTThostFtdcCFFEXUploadFileNameType[e]; ok {
		return s[strings.LastIndex(s, "_")+1:]
	}
	return string(e) + "值未定义"
}

// 组合指令方向类型
type TThostFtdcCombDirectionType byte

const THOST_FTDC_CMDR_Comb TThostFtdcCombDirectionType = '0' // 申请组合

const THOST_FTDC_CMDR_UnComb TThostFtdcCombDirectionType = '1' // 申请拆分

const THOST_FTDC_CMDR_DelComb TThostFtdcCombDirectionType = '2' // 操作员删组合单

var mpTThostFtdcCombDirectionType = map[TThostFtdcCombDirectionType]string{'0': "THOST_FTDC_CMDR_Comb", '1': "THOST_FTDC_CMDR_UnComb", '2': "THOST_FTDC_CMDR_DelComb"}

func (e TThostFtdcCombDirectionType) String() string {
	if s, ok := mpTThostFtdcCombDirectionType[e]; ok {
		return s[strings.LastIndex(s, "_")+1:]
	}
	return string(e) + "值未定义"
}

// 行权偏移类型类型
type TThostFtdcStrikeOffsetTypeType byte

const THOST_FTDC_STOV_RealValue TThostFtdcStrikeOffsetTypeType = '1' // 实值额

const THOST_FTDC_STOV_ProfitValue TThostFtdcStrikeOffsetTypeType = '2' // 盈利额

const THOST_FTDC_STOV_RealRatio TThostFtdcStrikeOffsetTypeType = '3' // 实值比例

const THOST_FTDC_STOV_ProfitRatio TThostFtdcStrikeOffsetTypeType = '4' // 盈利比例

var mpTThostFtdcStrikeOffsetTypeType = map[TThostFtdcStrikeOffsetTypeType]string{'1': "THOST_FTDC_STOV_RealValue", '2': "THOST_FTDC_STOV_ProfitValue", '3': "THOST_FTDC_STOV_RealRatio", '4': "THOST_FTDC_STOV_ProfitRatio"}

func (e TThostFtdcStrikeOffsetTypeType) String() string {
	if s, ok := mpTThostFtdcStrikeOffsetTypeType[e]; ok {
		return s[strings.LastIndex(s, "_")+1:]
	}
	return string(e) + "值未定义"
}

// 预约开户状态类型
type TThostFtdcReserveOpenAccStasType byte

const THOST_FTDC_ROAST_Processing TThostFtdcReserveOpenAccStasType = '0' // 等待处理中

const THOST_FTDC_ROAST_Cancelled TThostFtdcReserveOpenAccStasType = '1' // 已撤销

const THOST_FTDC_ROAST_Opened TThostFtdcReserveOpenAccStasType = '2' // 已开户

const THOST_FTDC_ROAST_Invalid TThostFtdcReserveOpenAccStasType = '3' // 无效请求

var mpTThostFtdcReserveOpenAccStasType = map[TThostFtdcReserveOpenAccStasType]string{'0': "THOST_FTDC_ROAST_Processing", '1': "THOST_FTDC_ROAST_Cancelled", '2': "THOST_FTDC_ROAST_Opened", '3': "THOST_FTDC_ROAST_Invalid"}

func (e TThostFtdcReserveOpenAccStasType) String() string {
	if s, ok := mpTThostFtdcReserveOpenAccStasType[e]; ok {
		return s[strings.LastIndex(s, "_")+1:]
	}
	return string(e) + "值未定义"
}

// 登录备注类型
type TThostFtdcLoginRemarkType Byte36

// 投资单元代码类型
type TThostFtdcInvestUnitIDType Byte17

// 公告编号类型
type TThostFtdcBulletinIDType int32

// 公告类型类型
type TThostFtdcNewsTypeType Byte3

// 紧急程度类型
type TThostFtdcNewsUrgencyType byte

// 消息摘要类型
type TThostFtdcAbstractType Byte81

// 消息来源类型
type TThostFtdcComeFromType Byte21

// WEB地址类型
type TThostFtdcURLLinkType Byte201

// 长个人姓名类型
type TThostFtdcLongIndividualNameType Byte161

// 长换汇银行账户名类型
type TThostFtdcLongFBEBankAccountNameType Byte161

// 日期时间类型
type TThostFtdcDateTimeType Byte17

// 弱密码来源类型
type TThostFtdcWeakPasswordSourceType byte

const THOST_FTDC_WPSR_Lib TThostFtdcWeakPasswordSourceType = '1' // 弱密码库

const THOST_FTDC_WPSR_Manual TThostFtdcWeakPasswordSourceType = '2' // 手工录入

var mpTThostFtdcWeakPasswordSourceType = map[TThostFtdcWeakPasswordSourceType]string{'1': "THOST_FTDC_WPSR_Lib", '2': "THOST_FTDC_WPSR_Manual"}

func (e TThostFtdcWeakPasswordSourceType) String() string {
	if s, ok := mpTThostFtdcWeakPasswordSourceType[e]; ok {
		return s[strings.LastIndex(s, "_")+1:]
	}
	return string(e) + "值未定义"
}

// 随机串类型
type TThostFtdcRandomStringType Byte17

// 报单回显字段类型
type TThostFtdcOrderMemoType Byte13

// 期权行权的头寸是否自对冲类型
type TThostFtdcOptSelfCloseFlagType byte

const THOST_FTDC_OSCF_CloseSelfOptionPosition TThostFtdcOptSelfCloseFlagType = '1' // 自对冲期权仓位

const THOST_FTDC_OSCF_ReserveOptionPosition TThostFtdcOptSelfCloseFlagType = '2' // 保留期权仓位

const THOST_FTDC_OSCF_SellCloseSelfFuturePosition TThostFtdcOptSelfCloseFlagType = '3' // 自对冲卖方履约后的期货仓位

const THOST_FTDC_OSCF_ReserveFuturePosition TThostFtdcOptSelfCloseFlagType = '4' // 保留卖方履约后的期货仓位

var mpTThostFtdcOptSelfCloseFlagType = map[TThostFtdcOptSelfCloseFlagType]string{'1': "THOST_FTDC_OSCF_CloseSelfOptionPosition", '2': "THOST_FTDC_OSCF_ReserveOptionPosition", '3': "THOST_FTDC_OSCF_SellCloseSelfFuturePosition", '4': "THOST_FTDC_OSCF_ReserveFuturePosition"}

func (e TThostFtdcOptSelfCloseFlagType) String() string {
	if s, ok := mpTThostFtdcOptSelfCloseFlagType[e]; ok {
		return s[strings.LastIndex(s, "_")+1:]
	}
	return string(e) + "值未定义"
}

// 业务类型类型
type TThostFtdcBizTypeType byte

const THOST_FTDC_BZTP_Future TThostFtdcBizTypeType = '1' // 期货

const THOST_FTDC_BZTP_Stock TThostFtdcBizTypeType = '2' // 证券

var mpTThostFtdcBizTypeType = map[TThostFtdcBizTypeType]string{'1': "THOST_FTDC_BZTP_Future", '2': "THOST_FTDC_BZTP_Stock"}

func (e TThostFtdcBizTypeType) String() string {
	if s, ok := mpTThostFtdcBizTypeType[e]; ok {
		return s[strings.LastIndex(s, "_")+1:]
	}
	return string(e) + "值未定义"
}

// 用户App类型类型
type TThostFtdcAppTypeType byte

const THOST_FTDC_APP_TYPE_Investor TThostFtdcAppTypeType = '1' // 直连的投资者

const THOST_FTDC_APP_TYPE_InvestorRelay TThostFtdcAppTypeType = '2' // 为每个投资者都创建连接的中继

const THOST_FTDC_APP_TYPE_OperatorRelay TThostFtdcAppTypeType = '3' // 所有投资者共享一个操作员连接的中继

const THOST_FTDC_APP_TYPE_UnKnown TThostFtdcAppTypeType = '4' // 未知

var mpTThostFtdcAppTypeType = map[TThostFtdcAppTypeType]string{'1': "THOST_FTDC_APP_TYPE_Investor", '2': "THOST_FTDC_APP_TYPE_InvestorRelay", '3': "THOST_FTDC_APP_TYPE_OperatorRelay", '4': "THOST_FTDC_APP_TYPE_UnKnown"}

func (e TThostFtdcAppTypeType) String() string {
	if s, ok := mpTThostFtdcAppTypeType[e]; ok {
		return s[strings.LastIndex(s, "_")+1:]
	}
	return string(e) + "值未定义"
}

// App代码类型
type TThostFtdcAppIDType Byte33

// 系统信息长度类型
type TThostFtdcSystemInfoLenType int32

// 补充信息长度类型
type TThostFtdcAdditionalInfoLenType int32

// 交易终端系统信息类型
type TThostFtdcClientSystemInfoType Byte273

// 系统外部信息类型
type TThostFtdcAdditionalInfoType Byte261

// base64交易终端系统信息类型
type TThostFtdcBase64ClientSystemInfoType Byte365

// base64系统外部信息类型
type TThostFtdcBase64AdditionalInfoType Byte349

// 当前可用的认证模式，0代表无需认证模式 A从低位开始最后一位代表图片验证码，倒数第二位代表动态口令，倒数第三位代表短信验证码类型
type TThostFtdcCurrentAuthMethodType int32

// 图片验证信息长度类型
type TThostFtdcCaptchaInfoLenType int32

// 图片验证信息类型
type TThostFtdcCaptchaInfoType Byte2561

// 用户短信验证码的编号类型
type TThostFtdcUserTextSeqType int32

// 握手数据内容类型
type TThostFtdcHandshakeDataType Byte301

// 握手数据内容长度类型
type TThostFtdcHandshakeDataLenType int32

// api与front通信密钥版本号类型
type TThostFtdcCryptoKeyVersionType Byte31

// 公钥版本号类型
type TThostFtdcRsaKeyVersionType int32

// 交易软件商ID类型
type TThostFtdcSoftwareProviderIDType Byte22

// 信息采集时间类型
type TThostFtdcCollectTimeType Byte21

// 查询频率类型
type TThostFtdcQueryFreqType int32

// 应答类型类型
type TThostFtdcResponseValueType byte

const THOST_FTDC_RV_Right TThostFtdcResponseValueType = '0' // 检查成功

const THOST_FTDC_RV_Refuse TThostFtdcResponseValueType = '1' // 检查失败

var mpTThostFtdcResponseValueType = map[TThostFtdcResponseValueType]string{'0': "THOST_FTDC_RV_Right", '1': "THOST_FTDC_RV_Refuse"}

func (e TThostFtdcResponseValueType) String() string {
	if s, ok := mpTThostFtdcResponseValueType[e]; ok {
		return s[strings.LastIndex(s, "_")+1:]
	}
	return string(e) + "值未定义"
}

// OTC成交类型类型
type TThostFtdcOTCTradeTypeType byte

const THOST_FTDC_OTC_TRDT_Block TThostFtdcOTCTradeTypeType = '0' // 大宗交易

const THOST_FTDC_OTC_TRDT_EFP TThostFtdcOTCTradeTypeType = '1' // 期转现

var mpTThostFtdcOTCTradeTypeType = map[TThostFtdcOTCTradeTypeType]string{'0': "THOST_FTDC_OTC_TRDT_Block", '1': "THOST_FTDC_OTC_TRDT_EFP"}

func (e TThostFtdcOTCTradeTypeType) String() string {
	if s, ok := mpTThostFtdcOTCTradeTypeType[e]; ok {
		return s[strings.LastIndex(s, "_")+1:]
	}
	return string(e) + "值未定义"
}

// 期现风险匹配方式类型
type TThostFtdcMatchTypeType byte

const THOST_FTDC_OTC_MT_DV01 TThostFtdcMatchTypeType = '1' // 基点价值

const THOST_FTDC_OTC_MT_ParValue TThostFtdcMatchTypeType = '2' // 面值

var mpTThostFtdcMatchTypeType = map[TThostFtdcMatchTypeType]string{'1': "THOST_FTDC_OTC_MT_DV01", '2': "THOST_FTDC_OTC_MT_ParValue"}

func (e TThostFtdcMatchTypeType) String() string {
	if s, ok := mpTThostFtdcMatchTypeType[e]; ok {
		return s[strings.LastIndex(s, "_")+1:]
	}
	return string(e) + "值未定义"
}

// OTC交易员代码类型
type TThostFtdcOTCTraderIDType Byte31

// 期货风险值类型
type TThostFtdcRiskValueType Double

// 握手数据内容类型
type TThostFtdcIDBNameType Byte100

// 折扣率类型
type TThostFtdcDiscountRatioType Double

// 用户终端认证方式类型
type TThostFtdcAuthTypeType byte

const THOST_FTDC_AU_WHITE TThostFtdcAuthTypeType = '0' // 白名单校验

const THOST_FTDC_AU_BLACK TThostFtdcAuthTypeType = '1' // 黑名单校验

var mpTThostFtdcAuthTypeType = map[TThostFtdcAuthTypeType]string{'0': "THOST_FTDC_AU_WHITE", '1': "THOST_FTDC_AU_BLACK"}

func (e TThostFtdcAuthTypeType) String() string {
	if s, ok := mpTThostFtdcAuthTypeType[e]; ok {
		return s[strings.LastIndex(s, "_")+1:]
	}
	return string(e) + "值未定义"
}

// 合约分类方式类型
type TThostFtdcClassTypeType byte

const THOST_FTDC_INS_ALL TThostFtdcClassTypeType = '0' // 所有合约

const THOST_FTDC_INS_FUTURE TThostFtdcClassTypeType = '1' // 期货、即期、期转现、Tas、金属指数合约

const THOST_FTDC_INS_OPTION TThostFtdcClassTypeType = '2' // 期货、现货期权合约

const THOST_FTDC_INS_COMB TThostFtdcClassTypeType = '3' // 组合合约

var mpTThostFtdcClassTypeType = map[TThostFtdcClassTypeType]string{'0': "THOST_FTDC_INS_ALL", '1': "THOST_FTDC_INS_FUTURE", '2': "THOST_FTDC_INS_OPTION", '3': "THOST_FTDC_INS_COMB"}

func (e TThostFtdcClassTypeType) String() string {
	if s, ok := mpTThostFtdcClassTypeType[e]; ok {
		return s[strings.LastIndex(s, "_")+1:]
	}
	return string(e) + "值未定义"
}

// 合约交易状态分类方式类型
type TThostFtdcTradingTypeType byte

const THOST_FTDC_TD_ALL TThostFtdcTradingTypeType = '0' // 所有状态

const THOST_FTDC_TD_TRADE TThostFtdcTradingTypeType = '1' // 交易

const THOST_FTDC_TD_UNTRADE TThostFtdcTradingTypeType = '2' // 非交易

var mpTThostFtdcTradingTypeType = map[TThostFtdcTradingTypeType]string{'0': "THOST_FTDC_TD_ALL", '1': "THOST_FTDC_TD_TRADE", '2': "THOST_FTDC_TD_UNTRADE"}

func (e TThostFtdcTradingTypeType) String() string {
	if s, ok := mpTThostFtdcTradingTypeType[e]; ok {
		return s[strings.LastIndex(s, "_")+1:]
	}
	return string(e) + "值未定义"
}

// 产品状态类型
type TThostFtdcProductStatusType byte

const THOST_FTDC_PS_tradeable TThostFtdcProductStatusType = '1' // 可交易

const THOST_FTDC_PS_untradeable TThostFtdcProductStatusType = '2' // 不可交易

var mpTThostFtdcProductStatusType = map[TThostFtdcProductStatusType]string{'1': "THOST_FTDC_PS_tradeable", '2': "THOST_FTDC_PS_untradeable"}

func (e TThostFtdcProductStatusType) String() string {
	if s, ok := mpTThostFtdcProductStatusType[e]; ok {
		return s[strings.LastIndex(s, "_")+1:]
	}
	return string(e) + "值未定义"
}

// 追平状态类型
type TThostFtdcSyncDeltaStatusType byte

const THOST_FTDC_SDS_Readable TThostFtdcSyncDeltaStatusType = '1' // 交易可读

const THOST_FTDC_SDS_Reading TThostFtdcSyncDeltaStatusType = '2' // 交易在读

const THOST_FTDC_SDS_Readend TThostFtdcSyncDeltaStatusType = '3' // 交易读取完成

const THOST_FTDC_SDS_OptErr TThostFtdcSyncDeltaStatusType = 'e' // 追平失败 交易本地状态结算不存在

var mpTThostFtdcSyncDeltaStatusType = map[TThostFtdcSyncDeltaStatusType]string{'1': "THOST_FTDC_SDS_Readable", '2': "THOST_FTDC_SDS_Reading", '3': "THOST_FTDC_SDS_Readend", 'e': "THOST_FTDC_SDS_OptErr"}

func (e TThostFtdcSyncDeltaStatusType) String() string {
	if s, ok := mpTThostFtdcSyncDeltaStatusType[e]; ok {
		return s[strings.LastIndex(s, "_")+1:]
	}
	return string(e) + "值未定义"
}

// 操作标志类型
type TThostFtdcActionDirectionType byte

const THOST_FTDC_ACD_Add TThostFtdcActionDirectionType = '1' // 增加

const THOST_FTDC_ACD_Del TThostFtdcActionDirectionType = '2' // 删除

const THOST_FTDC_ACD_Upd TThostFtdcActionDirectionType = '3' // 更新

var mpTThostFtdcActionDirectionType = map[TThostFtdcActionDirectionType]string{'1': "THOST_FTDC_ACD_Add", '2': "THOST_FTDC_ACD_Del", '3': "THOST_FTDC_ACD_Upd"}

func (e TThostFtdcActionDirectionType) String() string {
	if s, ok := mpTThostFtdcActionDirectionType[e]; ok {
		return s[strings.LastIndex(s, "_")+1:]
	}
	return string(e) + "值未定义"
}

// 撤单时选择席位算法类型
type TThostFtdcOrderCancelAlgType byte

const THOST_FTDC_OAC_Balance TThostFtdcOrderCancelAlgType = '1' // 轮询席位撤单

const THOST_FTDC_OAC_OrigFirst TThostFtdcOrderCancelAlgType = '2' // 优先原报单席位撤单

var mpTThostFtdcOrderCancelAlgType = map[TThostFtdcOrderCancelAlgType]string{'1': "THOST_FTDC_OAC_Balance", '2': "THOST_FTDC_OAC_OrigFirst"}

func (e TThostFtdcOrderCancelAlgType) String() string {
	if s, ok := mpTThostFtdcOrderCancelAlgType[e]; ok {
		return s[strings.LastIndex(s, "_")+1:]
	}
	return string(e) + "值未定义"
}

// 追平描述类型
type TThostFtdcSyncDescriptionType Byte257

// 通用int类型类型
type TThostFtdcCommonIntType int32

// 系统版本类型
type TThostFtdcSysVersionType Byte41

// 开仓量限制粒度类型
type TThostFtdcOpenLimitControlLevelType byte

const THOST_FTDC_PLCL_None TThostFtdcOpenLimitControlLevelType = '0' // 不控制

const THOST_FTDC_PLCL_Product TThostFtdcOpenLimitControlLevelType = '1' // 产品级别

const THOST_FTDC_PLCL_Inst TThostFtdcOpenLimitControlLevelType = '2' // 合约级别

var mpTThostFtdcOpenLimitControlLevelType = map[TThostFtdcOpenLimitControlLevelType]string{'0': "THOST_FTDC_PLCL_None", '1': "THOST_FTDC_PLCL_Product", '2': "THOST_FTDC_PLCL_Inst"}

func (e TThostFtdcOpenLimitControlLevelType) String() string {
	if s, ok := mpTThostFtdcOpenLimitControlLevelType[e]; ok {
		return s[strings.LastIndex(s, "_")+1:]
	}
	return string(e) + "值未定义"
}

// 报单频率控制粒度类型
type TThostFtdcOrderFreqControlLevelType byte

const THOST_FTDC_OFCL_None TThostFtdcOrderFreqControlLevelType = '0' // 不控制

const THOST_FTDC_OFCL_Product TThostFtdcOrderFreqControlLevelType = '1' // 产品级别

const THOST_FTDC_OFCL_Inst TThostFtdcOrderFreqControlLevelType = '2' // 合约级别

var mpTThostFtdcOrderFreqControlLevelType = map[TThostFtdcOrderFreqControlLevelType]string{'0': "THOST_FTDC_OFCL_None", '1': "THOST_FTDC_OFCL_Product", '2': "THOST_FTDC_OFCL_Inst"}

func (e TThostFtdcOrderFreqControlLevelType) String() string {
	if s, ok := mpTThostFtdcOrderFreqControlLevelType[e]; ok {
		return s[strings.LastIndex(s, "_")+1:]
	}
	return string(e) + "值未定义"
}

// 枚举bool类型类型
type TThostFtdcEnumBoolType byte

const THOST_FTDC_EBL_False TThostFtdcEnumBoolType = '0' // false

const THOST_FTDC_EBL_True TThostFtdcEnumBoolType = '1' // true

var mpTThostFtdcEnumBoolType = map[TThostFtdcEnumBoolType]string{'0': "THOST_FTDC_EBL_False", '1': "THOST_FTDC_EBL_True"}

func (e TThostFtdcEnumBoolType) String() string {
	if s, ok := mpTThostFtdcEnumBoolType[e]; ok {
		return s[strings.LastIndex(s, "_")+1:]
	}
	return string(e) + "值未定义"
}

// 期货合约阶段标识类型
type TThostFtdcTimeRangeType byte

const THOST_FTDC_ETR_USUAL TThostFtdcTimeRangeType = '1' // 一般月份

const THOST_FTDC_ETR_FNSP TThostFtdcTimeRangeType = '2' // 交割月前一个月上半月

const THOST_FTDC_ETR_BNSP TThostFtdcTimeRangeType = '3' // 交割月前一个月下半月

const THOST_FTDC_ETR_SPOT TThostFtdcTimeRangeType = '4' // 交割月份

var mpTThostFtdcTimeRangeType = map[TThostFtdcTimeRangeType]string{'1': "THOST_FTDC_ETR_USUAL", '2': "THOST_FTDC_ETR_FNSP", '3': "THOST_FTDC_ETR_BNSP", '4': "THOST_FTDC_ETR_SPOT"}

func (e TThostFtdcTimeRangeType) String() string {
	if s, ok := mpTThostFtdcTimeRangeType[e]; ok {
		return s[strings.LastIndex(s, "_")+1:]
	}
	return string(e) + "值未定义"
}

// Delta类型类型
type TThostFtdcDeltaType Double

// 抵扣组优先级类型
type TThostFtdcSpreadIdType int32

// 新型组保算法类型
type TThostFtdcPortfolioType byte

const THOST_FTDC_EPF_None TThostFtdcPortfolioType = '0' // 不使用新型组保算法

const THOST_FTDC_EPF_SPBM TThostFtdcPortfolioType = '1' // SPBM算法

const THOST_FTDC_EPF_RULE TThostFtdcPortfolioType = '2' // RULE算法

const THOST_FTDC_EPF_SPMM TThostFtdcPortfolioType = '3' // SPMM算法

const THOST_FTDC_EPF_RCAMS TThostFtdcPortfolioType = '4' // RCAMS算法

var mpTThostFtdcPortfolioType = map[TThostFtdcPortfolioType]string{'0': "THOST_FTDC_EPF_None", '1': "THOST_FTDC_EPF_SPBM", '2': "THOST_FTDC_EPF_RULE", '3': "THOST_FTDC_EPF_SPMM", '4': "THOST_FTDC_EPF_RCAMS"}

func (e TThostFtdcPortfolioType) String() string {
	if s, ok := mpTThostFtdcPortfolioType[e]; ok {
		return s[strings.LastIndex(s, "_")+1:]
	}
	return string(e) + "值未定义"
}

// SPBM组合套餐ID类型
type TThostFtdcPortfolioDefIDType int32

// 可提参数代码类型
type TThostFtdcWithDrawParamIDType byte

const THOST_FTDC_WDPID_CashIn TThostFtdcWithDrawParamIDType = 'C' // 权利金收支是否可提 1 代表可提 0 不可提

var mpTThostFtdcWithDrawParamIDType = map[TThostFtdcWithDrawParamIDType]string{'C': "THOST_FTDC_WDPID_CashIn"}

func (e TThostFtdcWithDrawParamIDType) String() string {
	if s, ok := mpTThostFtdcWithDrawParamIDType[e]; ok {
		return s[strings.LastIndex(s, "_")+1:]
	}
	return string(e) + "值未定义"
}

// 可提控制参数内容类型
type TThostFtdcWithDrawParamValueType Byte41

// 投资者交易权限类型
type TThostFtdcInvstTradingRightType byte

const THOST_FTDC_ITR_CloseOnly TThostFtdcInvstTradingRightType = '1' // 只能平仓

const THOST_FTDC_ITR_Forbidden TThostFtdcInvstTradingRightType = '2' // 不能交易

var mpTThostFtdcInvstTradingRightType = map[TThostFtdcInvstTradingRightType]string{'1': "THOST_FTDC_ITR_CloseOnly", '2': "THOST_FTDC_ITR_Forbidden"}

func (e TThostFtdcInvstTradingRightType) String() string {
	if s, ok := mpTThostFtdcInvstTradingRightType[e]; ok {
		return s[strings.LastIndex(s, "_")+1:]
	}
	return string(e) + "值未定义"
}

// Thost终端功能代码类型
type TThostFtdcThostFunctionCodeType int32

// SPMM折扣率类型
type TThostFtdcSPMMDiscountRatioType Double

// SPMM模板描述类型
type TThostFtdcSPMMModelDescType Byte129

// SPMM模板ID类型
type TThostFtdcSPMMModelIDType Byte33

// SPMM商品群商品组ID类型
type TThostFtdcSPMMProductIDType Byte41

// SPMM合约保证金算法类型
type TThostFtdcInstMarginCalIDType byte

const THOST_FTDC_IMID_BothSide TThostFtdcInstMarginCalIDType = '1' // 标准算法收取双边

const THOST_FTDC_IMID_MMSA TThostFtdcInstMarginCalIDType = '2' // 单向大边

const THOST_FTDC_IMID_SPMM TThostFtdcInstMarginCalIDType = '3' // 新组保SPMM

var mpTThostFtdcInstMarginCalIDType = map[TThostFtdcInstMarginCalIDType]string{'1': "THOST_FTDC_IMID_BothSide", '2': "THOST_FTDC_IMID_MMSA", '3': "THOST_FTDC_IMID_SPMM"}

func (e TThostFtdcInstMarginCalIDType) String() string {
	if s, ok := mpTThostFtdcInstMarginCalIDType[e]; ok {
		return s[strings.LastIndex(s, "_")+1:]
	}
	return string(e) + "值未定义"
}

// 产品ID类型
type TThostFtdcProductIDType Byte41

// HedgeRate类型类型
type TThostFtdcHedgeRateType Double

// 优先级类型
type TThostFtdcRCAMSPriorityType int32

// 空头期权风险调整标准类型类型
type TThostFtdcAdjustValueType Double

// RCAMS组合类型类型
type TThostFtdcRCAMSCombinationTypeType byte

const THOST_FTDC_ERComb_BUC TThostFtdcRCAMSCombinationTypeType = '0' // 牛市看涨价差组合

const THOST_FTDC_ERComb_BEC TThostFtdcRCAMSCombinationTypeType = '1' // 熊市看涨价差组合

const THOST_FTDC_ERComb_BEP TThostFtdcRCAMSCombinationTypeType = '2' // 熊市看跌价差组合

const THOST_FTDC_ERComb_BUP TThostFtdcRCAMSCombinationTypeType = '3' // 牛市看跌价差组合

const THOST_FTDC_ERComb_CAS TThostFtdcRCAMSCombinationTypeType = '4' // 日历价差组合

var mpTThostFtdcRCAMSCombinationTypeType = map[TThostFtdcRCAMSCombinationTypeType]string{'0': "THOST_FTDC_ERComb_BUC", '1': "THOST_FTDC_ERComb_BEC", '2': "THOST_FTDC_ERComb_BEP", '3': "THOST_FTDC_ERComb_BUP", '4': "THOST_FTDC_ERComb_CAS"}

func (e TThostFtdcRCAMSCombinationTypeType) String() string {
	if s, ok := mpTThostFtdcRCAMSCombinationTypeType[e]; ok {
		return s[strings.LastIndex(s, "_")+1:]
	}
	return string(e) + "值未定义"
}

// 策略id类型
type TThostFtdcRuleIdType Byte51

// 新组保算法启用类型类型
type TThostFtdcPortfTypeType byte

const THOST_FTDC_EET_None TThostFtdcPortfTypeType = '0' // 使用初版交易所算法

const THOST_FTDC_EET_SPBM_AddOnHedge TThostFtdcPortfTypeType = '1' // SPBM算法V1.1.0_附加保证金调整

var mpTThostFtdcPortfTypeType = map[TThostFtdcPortfTypeType]string{'0': "THOST_FTDC_EET_None", '1': "THOST_FTDC_EET_SPBM_AddOnHedge"}

func (e TThostFtdcPortfTypeType) String() string {
	if s, ok := mpTThostFtdcPortfTypeType[e]; ok {
		return s[strings.LastIndex(s, "_")+1:]
	}
	return string(e) + "值未定义"
}

// 合约类型类型
type TThostFtdcInstrumentClassType byte

const THOST_FTDC_EIC_Usual TThostFtdcInstrumentClassType = '1' // 一般月份合约

const THOST_FTDC_EIC_Delivery TThostFtdcInstrumentClassType = '2' // 临近交割合约

const THOST_FTDC_EIC_NonComb TThostFtdcInstrumentClassType = '3' // 非组合合约

var mpTThostFtdcInstrumentClassType = map[TThostFtdcInstrumentClassType]string{'1': "THOST_FTDC_EIC_Usual", '2': "THOST_FTDC_EIC_Delivery", '3': "THOST_FTDC_EIC_NonComb"}

func (e TThostFtdcInstrumentClassType) String() string {
	if s, ok := mpTThostFtdcInstrumentClassType[e]; ok {
		return s[strings.LastIndex(s, "_")+1:]
	}
	return string(e) + "值未定义"
}

// 商品群号类型
type TThostFtdcCommodityGroupIDType int32

// 标准持仓类型类型
type TThostFtdcStdPositionType Double

// 品种记录改变状态类型
type TThostFtdcProdChangeFlagType byte

const THOST_FTDC_PCF_None TThostFtdcProdChangeFlagType = '0' // 持仓量和冻结量均无变化

const THOST_FTDC_PCF_OnlyFrozen TThostFtdcProdChangeFlagType = '1' // 持仓量无变化，冻结量有变化

const THOST_FTDC_PCF_PositionChange TThostFtdcProdChangeFlagType = '2' // 持仓量有变化

var mpTThostFtdcProdChangeFlagType = map[TThostFtdcProdChangeFlagType]string{'0': "THOST_FTDC_PCF_None", '1': "THOST_FTDC_PCF_OnlyFrozen", '2': "THOST_FTDC_PCF_PositionChange"}

func (e TThostFtdcProdChangeFlagType) String() string {
	if s, ok := mpTThostFtdcProdChangeFlagType[e]; ok {
		return s[strings.LastIndex(s, "_")+1:]
	}
	return string(e) + "值未定义"
}

// 历史密码来源类型
type TThostFtdcPwdRcdSrcType byte

const THOST_FTDC_PRS_Init TThostFtdcPwdRcdSrcType = '0' // 来源于Sync初始化数据

const THOST_FTDC_PRS_Sync TThostFtdcPwdRcdSrcType = '1' // 来源于实时上场数据

const THOST_FTDC_PRS_UserUpd TThostFtdcPwdRcdSrcType = '2' // 来源于用户修改

const THOST_FTDC_PRS_SuperUserUpd TThostFtdcPwdRcdSrcType = '3' // 来源于超户修改，很可能来自主席同步数据

const THOST_FTDC_PRS_SecUpd TThostFtdcPwdRcdSrcType = '4' // 来源于次席同步的修改

var mpTThostFtdcPwdRcdSrcType = map[TThostFtdcPwdRcdSrcType]string{'0': "THOST_FTDC_PRS_Init", '1': "THOST_FTDC_PRS_Sync", '2': "THOST_FTDC_PRS_UserUpd", '3': "THOST_FTDC_PRS_SuperUserUpd", '4': "THOST_FTDC_PRS_SecUpd"}

func (e TThostFtdcPwdRcdSrcType) String() string {
	if s, ok := mpTThostFtdcPwdRcdSrcType[e]; ok {
		return s[strings.LastIndex(s, "_")+1:]
	}
	return string(e) + "值未定义"
}

// 地址服务类型类型
type TThostFtdcAddrSrvModeType byte

const THOST_FTDC_ASM_Trade TThostFtdcAddrSrvModeType = '0' // 交易地址

const THOST_FTDC_ASM_MarketData TThostFtdcAddrSrvModeType = '1' // 行情地址

const THOST_FTDC_ASM_Other TThostFtdcAddrSrvModeType = '2' // 其他

var mpTThostFtdcAddrSrvModeType = map[TThostFtdcAddrSrvModeType]string{'0': "THOST_FTDC_ASM_Trade", '1': "THOST_FTDC_ASM_MarketData", '2': "THOST_FTDC_ASM_Other"}

func (e TThostFtdcAddrSrvModeType) String() string {
	if s, ok := mpTThostFtdcAddrSrvModeType[e]; ok {
		return s[strings.LastIndex(s, "_")+1:]
	}
	return string(e) + "值未定义"
}

// 地址版本类型
type TThostFtdcAddrVerType byte

const THOST_FTDC_ADV_V4 TThostFtdcAddrVerType = '0' // IPV4

const THOST_FTDC_ADV_V6 TThostFtdcAddrVerType = '1' // IPV6

var mpTThostFtdcAddrVerType = map[TThostFtdcAddrVerType]string{'0': "THOST_FTDC_ADV_V4", '1': "THOST_FTDC_ADV_V6"}

func (e TThostFtdcAddrVerType) String() string {
	if s, ok := mpTThostFtdcAddrVerType[e]; ok {
		return s[strings.LastIndex(s, "_")+1:]
	}
	return string(e) + "值未定义"
}

// 地址备注类型
type TThostFtdcAddrRemarkType Byte161

// 地址名称类型
type TThostFtdcAddrNameType Byte65

// 服务地址IP类型
type TThostFtdcIpAddrType Byte129

// TGATE会话查询状态类型
type TThostFtdcTGSessionQryStatusType byte

const THOST_FTDC_TGQS_QryIdle TThostFtdcTGSessionQryStatusType = '1' // 查询状态空闲

const THOST_FTDC_TGQS_QryBusy TThostFtdcTGSessionQryStatusType = '2' // 查询状态频繁

var mpTThostFtdcTGSessionQryStatusType = map[TThostFtdcTGSessionQryStatusType]string{'1': "THOST_FTDC_TGQS_QryIdle", '2': "THOST_FTDC_TGQS_QryBusy"}

func (e TThostFtdcTGSessionQryStatusType) String() string {
	if s, ok := mpTThostFtdcTGSessionQryStatusType[e]; ok {
		return s[strings.LastIndex(s, "_")+1:]
	}
	return string(e) + "值未定义"
}

// 对冲类型类型
type TThostFtdcOffsetTypeType byte

const THOST_FTDC_OT_OPT_OFFSET TThostFtdcOffsetTypeType = '0' // 期权对冲

const THOST_FTDC_OT_FUT_OFFSET TThostFtdcOffsetTypeType = '1' // 期货对冲

const THOST_FTDC_OT_EXEC_OFFSET TThostFtdcOffsetTypeType = '2' // 行权后期货对冲

const THOST_FTDC_OT_PERFORM_OFFSET TThostFtdcOffsetTypeType = '3' // 履约后期货对冲

var mpTThostFtdcOffsetTypeType = map[TThostFtdcOffsetTypeType]string{'0': "THOST_FTDC_OT_OPT_OFFSET", '1': "THOST_FTDC_OT_FUT_OFFSET", '2': "THOST_FTDC_OT_EXEC_OFFSET", '3': "THOST_FTDC_OT_PERFORM_OFFSET"}

func (e TThostFtdcOffsetTypeType) String() string {
	if s, ok := mpTThostFtdcOffsetTypeType[e]; ok {
		return s[strings.LastIndex(s, "_")+1:]
	}
	return string(e) + "值未定义"
}

// 站点类型
type TThostFtdcSiteType Byte51

// 网络运营商类型
type TThostFtdcNetOperatorType Byte9

// 申请来源类型
type TThostFtdcApplySrcType byte

const THOST_FTDC_AS_Trade TThostFtdcApplySrcType = '0' // 交易

const THOST_FTDC_AS_Member TThostFtdcApplySrcType = '1' // 会服

var mpTThostFtdcApplySrcType = map[TThostFtdcApplySrcType]string{'0': "THOST_FTDC_AS_Trade", '1': "THOST_FTDC_AS_Member"}

func (e TThostFtdcApplySrcType) String() string {
	if s, ok := mpTThostFtdcApplySrcType[e]; ok {
		return s[strings.LastIndex(s, "_")+1:]
	}
	return string(e) + "值未定义"
}

// 预留信息类型
type TThostFtdcReserveInfoType Byte65

// 套利套保申请状态类型
type TThostFtdcApplyStatusType byte

const THOST_FTDC_AS2_Unknown TThostFtdcApplyStatusType = 'a' // 未知

const THOST_FTDC_AS2_Canceled TThostFtdcApplyStatusType = '0' // 已撤销

const THOST_FTDC_AS2_Suspended TThostFtdcApplyStatusType = '1' // 已挂起

const THOST_FTDC_AS2_Accepted TThostFtdcApplyStatusType = '3' // 申请成功

var mpTThostFtdcApplyStatusType = map[TThostFtdcApplyStatusType]string{'a': "THOST_FTDC_AS2_Unknown", '0': "THOST_FTDC_AS2_Canceled", '1': "THOST_FTDC_AS2_Suspended", '3': "THOST_FTDC_AS2_Accepted"}

func (e TThostFtdcApplyStatusType) String() string {
	if s, ok := mpTThostFtdcApplyStatusType[e]; ok {
		return s[strings.LastIndex(s, "_")+1:]
	}
	return string(e) + "值未定义"
}

// 组合定单类型
type TThostFtdcCmbTypeType byte

const THOST_FTDC_ZCECT_SPZ TThostFtdcCmbTypeType = '0' // SPZ

const THOST_FTDC_ZCECT_SPD TThostFtdcCmbTypeType = '1' // SPD

const THOST_FTDC_ZCECT_IPS TThostFtdcCmbTypeType = '2' // IPS

const THOST_FTDC_ZCECT_BUL TThostFtdcCmbTypeType = '3' // BUL

const THOST_FTDC_ZCECT_BER TThostFtdcCmbTypeType = '4' // BER

const THOST_FTDC_ZCECT_BLT TThostFtdcCmbTypeType = '5' // BLT

const THOST_FTDC_ZCECT_BRT TThostFtdcCmbTypeType = '6' // BRT

const THOST_FTDC_ZCECT_STD TThostFtdcCmbTypeType = '7' // STD

const THOST_FTDC_ZCECT_STG TThostFtdcCmbTypeType = '8' // STG

const THOST_FTDC_ZCECT_PRT TThostFtdcCmbTypeType = '9' // PRT

var mpTThostFtdcCmbTypeType = map[TThostFtdcCmbTypeType]string{'0': "THOST_FTDC_ZCECT_SPZ", '1': "THOST_FTDC_ZCECT_SPD", '2': "THOST_FTDC_ZCECT_IPS", '3': "THOST_FTDC_ZCECT_BUL", '4': "THOST_FTDC_ZCECT_BER", '5': "THOST_FTDC_ZCECT_BLT", '6': "THOST_FTDC_ZCECT_BRT", '7': "THOST_FTDC_ZCECT_STD", '8': "THOST_FTDC_ZCECT_STG", '9': "THOST_FTDC_ZCECT_PRT"}

func (e TThostFtdcCmbTypeType) String() string {
	if s, ok := mpTThostFtdcCmbTypeType[e]; ok {
		return s[strings.LastIndex(s, "_")+1:]
	}
	return string(e) + "值未定义"
}

// 短信验证码类型
type TThostFtdcSMSCodeType Byte17

// 短信内容类型
type TThostFtdcSMSContentType Byte129

// 手机号类型
type TThostFtdcSMSPhoneType Byte17
