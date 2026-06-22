package v6_7_13_test

import (
	"testing"

	"github.com/frozenpine/ctp4go/trader/v6_7_13"
)

func TestCTPApi(t *testing.T) {
	libPath := "../../dependencies/future/v6.7.13/thosttraderapi_se.dll"

	ins, err := v6_7_13.CreateThostFtdcTraderApi(libPath, "./", true)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(ins.GetApiVersion())
}
