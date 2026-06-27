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

	ins, err := mduser.CreateThostFtdcMduserApi(
		libPath, "./flow", false, false, false,
	)
	if err != nil {
		t.Fatal(err)
	}
	defer ins.Release()

	ins.RegisterSpi(&thost.ThostLogSpi{})

	t.Log(ins.GetApiVersion())
}
