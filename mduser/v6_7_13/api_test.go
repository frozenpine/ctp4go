package v6_7_13_test

import (
	"log/slog"
	"testing"

	mduser "github.com/frozenpine/ctp4go/mduser/v6_7_13"
	"github.com/frozenpine/ctp4go/thost"
)

func TestMduserApi(t *testing.T) {
	slog.SetLogLoggerLevel(slog.LevelDebug - 2)

	libPath := "../../dependencies/future/v6.7.13/thostmduserapi_se.dll"
	// libPath := "../../dependencies/future/v6.7.13/thostmduserapi_se.so"

	//  SHZJ - CT
	front := "tcp://180.166.6.245:51205"
	//  RDXM - CT
	//  front := "tcp://222.76.240.170:51205"

	ins, err := mduser.CreateThostFtdcMduserApi(
		libPath, "./flow", false, false, true,
	)
	if err != nil {
		t.Fatal(err)
	}
	defer ins.Release()

	t.Log(ins.GetApiVersion())

	ins.RegisterSpi(&thost.ThostLogSpi{})
	ins.RegisterFront(front)
	// ins.Init()

	// <-time.After(time.Second * 10)
}
