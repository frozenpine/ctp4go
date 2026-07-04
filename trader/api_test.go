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
	// front := "tcp://180.166.6.245:51205"
	//  RDXM - CT
	//  front := "tcp://222.76.240.170:51205"

	// simnow
	// 主用：182.254.243.31:30001 行情30011
	// 备用: 182.254.243.31:30002 行情30012
	// 7*24：182.254.243.31:40001 行情40011
	front := "tcp://182.254.243.31:40001"

	td, err := trader.NewTraderApi(
		t.Context(),
		trader.WithLibPath(libPath),
		trader.WithFrontAddr(front),
		trader.WithBrokerID("9999"),
		trader.WithUserID("164889"),
		trader.WithAppID("simnow_client_test"),
		// trader.WithAuthCodeEnvKey("RDQH_CTP_AUTH_CODE"),
		// trader.WithUserPassEnvKey("RDQH_CTP_USER_PASS"),
		trader.WithAuthCode("0000000000000000"),
		trader.WithUserPass("022010blue@safe"),
	)
	if err != nil {
		t.Fatal(err)
	}
	defer td.Finalize()

	// var resetOnce sync.Once
	done := make(chan struct{})

	if err = td.Initialize(
		trader.WithTraderState(
			trader.WithStateResponsor(trader.Connected, td.Authenticate),
			trader.WithStateResponsor(trader.AuthSuccess, td.Login),
			// 重置会锁状态，Responsor执行也会锁状态
			// 此处需要异步以避免死锁
			// trader.WithStateResponsor(trader.LoginSuccess, func() error {
			// 	resetOnce.Do(func() {
			// 		go td.Reset()
			// 	})

			// 	return nil
			// }),
			trader.WithStateResponsor(trader.LoginSuccess, func() error {
				close(done)
				return nil
			}),
			trader.WithStateResponsor(trader.LoginFailed, func() error {
				close(done)
				return nil
			}),
		),
	); err != nil {
		t.Fatal(err)
	}

	select {
	case <-done:
	case <-time.After(time.Second * 20):
	}
}
