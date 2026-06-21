package trader

import (
	"time"

	"github.com/frozenpine/ctp4go/thost"
)

type TraderBase interface {
	Init()
	Release()
	Join(timeout time.Duration) error

	GetApiVersion() string
	GetTradingDay() string
	GetFrontInfo() (*thost.CThostFtdcFrontInfoField, error)

	RegisterFront(conn string) error
	RegisterNameServer(conn string) error
	RegisterFensUserInfo(fens *thost.CThostFtdcFensUserInfoField) error

	SubscribePrivateTopic(resume int) error
	SubscribePublicTopic(resume int) error
}

type TraderAuth interface {
	ReqAuthenticate(appId, authCode string) error
	SubmitUserSystemInfo(sysInfo *thost.CThostFtdcUserSystemInfoField) error
	ReqUserLogin(user, pass string) error
	ReqUserLogout() error
	ReqUserPasswordUpdate(newPass *thost.CThostFtdcUserPasswordUpdateField) error
	ReqTradingAccountPasswordUpdate(newPass *thost.CThostFtdcTradingAccountPasswordUpdateField) error
	ReqUserAuthMethod()
}

type TraderQuery interface {
	ReqQryOrder() error
	ReqQryTrade() error
	ReqQryInvestorPosition() error
	ReqQryTradingAccount() error
	ReqQryInvestor() error
	ReqQryTradingCode() error
	ReqQryInstrumentMarginRate() error
	ReqQryInstrumentCommissionRate() error

	ReqQryExchange() error
	ReqQryProduct() error
	ReqQryInstrument(symbols ...string) error
	ReqQryDepthMarketData() error
}
