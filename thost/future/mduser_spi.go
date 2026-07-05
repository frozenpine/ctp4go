package future

type MdSpi interface {
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

	// 登录请求响应
	OnRspUserLogin(RspUserLogin *CThostFtdcRspUserLoginField, RspInfo *CThostFtdcRspInfoField, RequestID int, IsLast bool)

	// 登出请求响应
	OnRspUserLogout(UserLogout *CThostFtdcUserLogoutField, RspInfo *CThostFtdcRspInfoField, RequestID int, IsLast bool)

	// 请求查询组播合约响应
	OnRspQryMulticastInstrument(MulticastInstrument *CThostFtdcMulticastInstrumentField, RspInfo *CThostFtdcRspInfoField, RequestID int, IsLast bool)

	// 错误应答
	OnRspError(RspInfo *CThostFtdcRspInfoField, RequestID int, IsLast bool)

	// 订阅行情应答
	OnRspSubMarketData(SpecificInstrument *CThostFtdcSpecificInstrumentField, RspInfo *CThostFtdcRspInfoField, RequestID int, IsLast bool)

	// 取消订阅行情应答
	OnRspUnSubMarketData(SpecificInstrument *CThostFtdcSpecificInstrumentField, RspInfo *CThostFtdcRspInfoField, RequestID int, IsLast bool)

	// 订阅询价应答
	OnRspSubForQuoteRsp(SpecificInstrument *CThostFtdcSpecificInstrumentField, RspInfo *CThostFtdcRspInfoField, RequestID int, IsLast bool)

	// 取消订阅询价应答
	OnRspUnSubForQuoteRsp(SpecificInstrument *CThostFtdcSpecificInstrumentField, RspInfo *CThostFtdcRspInfoField, RequestID int, IsLast bool)

	// 深度行情通知
	OnRtnDepthMarketData(DepthMarketData *CThostFtdcDepthMarketDataField)

	// 询价通知
	OnRtnForQuoteRsp(ForQuoteRsp *CThostFtdcForQuoteRspField)
}
