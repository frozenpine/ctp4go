package v6_7_13_test

import (
	"log/slog"
	"testing"
	"time"

	"github.com/frozenpine/ctp4go/thost"
	"github.com/frozenpine/ctp4go/thost/types"
	trader "github.com/frozenpine/ctp4go/trader/v6_7_13"
)

type testSpi struct {
	thost.ThostLogSpi

	api      *trader.ThostFtdcTraderApi
	brokerId string
	userId   string
	userPass string
	appId    string
	authCode string
}

func (s *testSpi) OnFrontConnected() {
	auth := thost.CThostFtdcReqAuthenticateField{}

	types.SetCString(auth.BrokerID[:], s.brokerId)
	types.SetCString(auth.UserID[:], s.userId)
	types.SetCString(auth.AppID[:], s.appId)
	types.SetCString(auth.AuthCode[:], s.authCode)

	s.api.ReqAuthenticate(&auth, 1)
}

func (s *testSpi) OnRspAuthenticate(
	pRspAuthenticateField *thost.CThostFtdcRspAuthenticateField,
	pRspInfo *thost.CThostFtdcRspInfoField,
	nRequestID int, bIsLast bool,
) {
	s.ThostLogSpi.OnRspAuthenticate(
		pRspAuthenticateField, pRspInfo, nRequestID, bIsLast,
	)

	login := thost.CThostFtdcReqUserLoginField{}
	types.SetCString(login.BrokerID[:], s.brokerId)
	types.SetCString(login.UserID[:], s.userId)
	types.SetCString(login.Password[:], s.userPass)

	s.api.ReqUserLogin(&login, nRequestID+1)
}

func TestCTPApi(t *testing.T) {
	slog.SetLogLoggerLevel(slog.LevelDebug - 2)

	libPath := "../../dependencies/future/v6.7.13/thosttraderapi_se.dll"
	// libPath := "../../dependencies/future/v6.7.13/thosttraderapi_se.so"

	//  SHZJ - CT
	front := "tcp://180.166.6.245:51205"
	//  RDXM - CT
	//  front := "tcp://222.76.240.170:51205"

	ins, err := trader.CreateThostFtdcTraderApi(libPath, "./flow", true)
	if err != nil {
		t.Fatal(err)
	}
	defer ins.Release()

	t.Log(ins.GetApiVersion())

	ins.RegisterSpi(&testSpi{
		ThostLogSpi: thost.ThostLogSpi{
			Logger: slog.Default(),
		},
		api: ins,

		brokerId: "5100",
		userId:   "000008",
		userPass: "xxxx",
		appId:    "client_dtprobe_1.0.0",
		authCode: "xxxx",
	})
	ins.SubscribePrivateTopic(types.THOST_TERT_QUICK, 1)
	ins.SubscribePublicTopic(types.THOST_TERT_QUICK)
	ins.RegisterFront(front)
	ins.Init()

	<-time.After(time.Second * 20)

	t.Log(ins.GetTradingDay())

	info := thost.CThostFtdcFrontInfoField{}
	ins.GetFrontInfo(&info)
	t.Logf("%+v", info)
}
