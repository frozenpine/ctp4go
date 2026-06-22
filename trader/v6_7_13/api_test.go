package v6_7_13_test

import (
	"log/slog"
	"testing"
	"time"

	"github.com/frozenpine/ctp4go/thost"
	"github.com/frozenpine/ctp4go/trader/v6_7_13"
)

type testSpi struct {
	thost.BaseLogSpi

	api      *v6_7_13.ThostFtdcTraderApi
	brokerId string
	userId   string
	userPass string
	appId    string
	authCode string
}

func (s *testSpi) OnFrontConnected() {
	auth := thost.CThostFtdcReqAuthenticateField{}

	thost.SetCString(auth.BrokerID[:], s.brokerId)
	thost.SetCString(auth.UserID[:], s.userId)
	thost.SetCString(auth.AppID[:], s.appId)
	thost.SetCString(auth.AuthCode[:], s.authCode)

	s.api.ReqAuthenticate(&auth, 1)
}

func (s *testSpi) OnRspAuthenticate(
	pRspAuthenticateField *thost.CThostFtdcRspAuthenticateField,
	pRspInfo *thost.CThostFtdcRspInfoField,
	nRequestID int, bIsLast bool,
) {
	s.BaseLogSpi.OnRspAuthenticate(
		pRspAuthenticateField, pRspInfo, nRequestID, bIsLast,
	)

	login := thost.CThostFtdcReqUserLoginField{}
	thost.SetCString(login.BrokerID[:], s.brokerId)
	thost.SetCString(login.UserID[:], s.userId)
	thost.SetCString(login.Password[:], s.userPass)

	s.api.ReqUserLogin(&login, nRequestID+1)
}

func TestCTPApi(t *testing.T) {
	slog.SetLogLoggerLevel(slog.LevelDebug - 2)

	// libPath := "../../dependencies/future/v6.7.13/thosttraderapi_se.dll"
	libPath := "../../dependencies/future/v6.7.13/thosttraderapi_se.so"

	//  SHZJ - CT
	front := "tcp://180.166.6.245:51205"
	//  RDXM - CT
	//  front := "tcp://222.76.240.170:51205"

	ins, err := v6_7_13.CreateThostFtdcTraderApi(libPath, "./flow", false)
	if err != nil {
		t.Fatal(err)
	}
	defer ins.Release()

	t.Log(ins.GetApiVersion())

	ins.RegisterSpi(&testSpi{
		BaseLogSpi: thost.BaseLogSpi{
			Logger:    slog.Default(),
			FrontAddr: front,
		},
		api: ins,

		brokerId: "5100",
		userId:   "000008",
		userPass: "xxxx",
		appId:    "client_dtprobe_1.0.0",
		authCode: "xxxxx",
	})
	ins.SubscribePrivateTopic(thost.THOST_TERT_QUICK, 1)
	ins.SubscribePublicTopic(thost.THOST_TERT_QUICK)
	ins.RegisterFront(front)
	ins.Init()

	t.Log(ins.GetTradingDay())
	info := thost.CThostFtdcFrontInfoField{}
	ins.GetFrontInfo(&info)
	t.Logf("%+v", info)

	<-time.After(time.Second * 20)
}
