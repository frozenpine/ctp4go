package v6_7_13_test

import (
	"testing"

	"github.com/frozenpine/ctp4go/thost"
	"github.com/frozenpine/ctp4go/trader/v6_7_13"
)

func TestCTPApi(t *testing.T) {
	libPath := "../../dependencies/future/v6.7.13/thosttraderapi_se.dll"

	//  SHZJ - CT
	front := "tcp://180.166.6.245:51205"
	//  RDXM - CT
	//  front := "tcp://222.76.240.170:51205"

	ins, err := v6_7_13.CreateThostFtdcTraderApi(libPath, "./flow", true)
	if err != nil {
		t.Fatal(err)
	}
	defer ins.Release()

	t.Log(ins.GetApiVersion())

	ins.RegisterFront(front)
	ins.SubscribePrivateTopic(thost.THOST_TERT_QUICK, -1)
	ins.SubscribePublicTopic(thost.THOST_TERT_QUICK)
	ins.Init()
}
