package future_test

import (
	"log/slog"
	"testing"
	"time"

	"github.com/frozenpine/ctp4go/mduser/future"
)

func TestMdApi(t *testing.T) {
	slog.SetLogLoggerLevel(slog.LevelDebug - 2)

	libPath := "../../dependencies/future/v6.7.13/thostmduserapi_se.dll"
	// libPath := "../../dependencies/future/v6.7.13/thostmduserapi_se.so"

	//  SHZJ - CT
	// front := "tcp://180.166.6.245:51205"
	//  RDXM - CT
	//  front := "tcp://222.76.240.170:51205"

	// simnow
	// 主用：182.254.243.31:30001 行情30011
	// 备用: 182.254.243.31:30002 行情30012
	// 7*24：182.254.243.31:40001 行情40011
	front := "tcp://182.254.243.31:40001"

	md, err := future.NewMduserApi(
		t.Context(),
		future.WithLibPath(libPath),
		future.WithFrontAddr(front),
		future.WithBrokerID("9999"),
		future.WithUserID("164889"),
		future.WithUserPassEnvKey("RDQH_CTP_USER_PASS"),
	)
	if err != nil {
		t.Fatal(err)
	}
	defer md.Finalize()

	done := make(chan struct{})

	if err = md.Initialize(
		future.WithMduserState(
			future.WithStateResponsor(future.Connected, md.Login),
			future.WithStateResponsor(future.LoginSuccess, func() error {
				return md.Subscribe("ag2609")
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
