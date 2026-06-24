package trader_test

import (
	"log/slog"
	"testing"
	"time"

	"github.com/frozenpine/ctp4go/trader"
)

func TestTraderApi(t *testing.T) {
	slog.SetLogLoggerLevel(slog.LevelDebug - 2)

	libPath := "../dependencies/future/v6.7.13/thosttraderapi_se.dll"
	// libPath := "../dependencies/future/v6.7.13/thosttraderapi_se.so"

	//  SHZJ - CT
	front := "tcp://180.166.6.245:51205"
	//  RDXM - CT
	//  front := "tcp://222.76.240.170:51205"

	td, err := trader.NewTraderApi(
		t.Context(), libPath,
		trader.WithFrontAddr(front),
		trader.WithBrokerID("5100"),
		trader.WithUserID("000008"),
		trader.WithAppID("client_dtprobe_1.0.0"),
		trader.WithAuthCodeEnvKey("RDQH_CTP_AUTH_CODE"),
	)
	if err != nil {
		t.Fatal(err)
	}
	defer td.Finalize()

	if err = td.Initialize(
		trader.WithStateResponsor(trader.Connected, td.Authenticate),
		trader.WithStateResponsor(trader.AuthSuccess, td.Login),
	); err != nil {
		t.Fatal(err)
	}

	<-time.After(time.Second * 15)
}
