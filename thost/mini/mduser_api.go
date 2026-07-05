package mini

type MdApi interface {
    // 删除接口对象本身
    //   @remark  不再使用本接口对象时,调用该函数删除接口对象
    Release()

    // 初始化
    //   @remark  初始化运行环境,只有调用后,接口才开始工作
    Init(Continuousm bool)

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
    RegisterSpi(Spi MdSpi)

    // 订阅行情。
    //   @param ppInstrumentID 合约ID
    //   @param nCount 要订阅退订行情的合约个数
    //   @remark  
    SubscribeMarketData(InstrumentID ...string) int

    // 取消订阅行情。
    UnSubscribeMarketData(InstrumentID ...string) int

    // 订阅询价。
    //   @param ppInstrumentID 合约ID
    //   @param nCount 要订阅退订行情的合约个数
    //   @remark  
    SubscribeForQuoteRsp(InstrumentID ...string) int

    // 退订询价。
    //   @param ppInstrumentID 合约ID
    //   @param nCount 要订阅退订行情的合约个数
    //   @remark  
    UnSubscribeForQuoteRsp(InstrumentID ...string) int

    // 用户登录请求
    ReqUserLogin(ReqUserLoginField *CThostFtdcReqUserLoginField, RequestID int) int

    // 用户加密登录请求
    ReqUserLoginEncrypt(ReqUserLoginField *CThostFtdcReqUserLoginField, RequestID int) int

    // 登出请求
    ReqUserLogout(UserLogout *CThostFtdcUserLogoutField, RequestID int) int

}
