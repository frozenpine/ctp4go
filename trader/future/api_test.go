package future_test

import (
	"testing"
	"time"

	"github.com/frozenpine/ctp4go/state"
	thost_futer "github.com/frozenpine/ctp4go/thost/future"
	"github.com/frozenpine/ctp4go/trader/future"
)

func TestTraderApi(t *testing.T) {
	// slog.SetLogLoggerLevel(slog.LevelDebug - 2)

	libPath := "../../dependencies/future/v6.7.13/thosttraderapi_se.dll"
	// libPath := "../../dependencies/future/v6.7.13/thosttraderapi_se.so"

	//  SHZJ - CT
	// front := "tcp://180.166.6.245:51205"
	//  RDXM - CT
	//  front := "tcp://222.76.240.170:51205"

	// FZ md_port: 42213
	// front := "tcp://101.226.250.133:42205"

	// simnow
	// 主用：182.254.243.31:30001 行情30011
	// 备用: 182.254.243.31:30002 行情30012
	// 7*24：182.254.243.31:40001 行情40011
	front := "tcp://182.254.243.31:40001"

	td, err := future.NewTraderApi(
		t.Context(),
		future.WithLibPath(libPath),
		future.WithFrontAddr(front),
		// future.WithTestMode(),

		future.WithBrokerID("9999"),
		future.WithUserID("164889"),
		future.WithAppID("simnow_client_test"),
		future.WithAuthCode("0000000000000000"),
		future.WithUserPassEnvKey("SIMNOW_USER_PASS"),

		// future.WithBrokerID("0121"),
		// future.WithUserID("666666"),
		// future.WithAuthCodeEnvKey("RDQH_CTP_AUTH_CODE"),
		// future.WithUserPassEnvKey("RDQH_CTP_USER_PASS"),
		// future.WithUserPass("rdqh@123456"),
	)
	if err != nil {
		t.Fatal(err)
	}
	defer td.Finalize()

	// var resetOnce sync.Once
	done := make(chan struct{})

	if err = td.Initialize(
		future.WithTraderState(
			future.WithStateResponsor(future.Connected, td.Authenticate),
			future.WithStateResponsor(future.AuthSuccess, td.Login),
			// 重置会锁状态，Responsor执行也会锁状态
			// 此处需要异步以避免死锁
			// future.WithStateResponsor(future.LoginSuccess, func() error {
			// 	resetOnce.Do(func() {
			// 		go td.Reset()
			// 	})

			// 	return nil
			// }),
			future.WithStateResponsor(future.LoginSuccess, func() error {
				qry := thost_futer.CThostFtdcQryInstrumentField{}

				r, err := state.MakeRequest(td.RFactory, &qry)
				if err != nil {
					return err
				}

				return td.DoRequest(r)
			}),
			future.WithStateResponsor(future.LoginFailed, func() error {
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

	t.Log(td.GetCacheData(thost_futer.InsCache, "Symbol:SHFE.zn2611P21600"))

	for idx, v := range td.IterCacheData(
		thost_futer.InsCache, func(cfif state.Data) bool {
			if v, err := cfif.GetFieldString("ProductID"); err != nil {
				return false
			} else {
				return v == "ag"
			}
		},
	) {
		t.Log(idx, v)
	}
}
