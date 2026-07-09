package state_test

import (
	"testing"

	"github.com/frozenpine/ctp4go/state"
	"github.com/frozenpine/ctp4go/thost/future"
	"github.com/frozenpine/ctp4go/thost/future/types"
	v6_7_13 "github.com/frozenpine/ctp4go/trader/future/v6.7.13"
)

func TestReqFactory(t *testing.T) {
	var api future.TraderApi = new(v6_7_13.ThostFtdcTraderApi)

	factory := state.NewRequestFactory(t.Context(), api)

	t.Logf("%+v", factory)

	req, err := state.MakeRequest(
		factory, &future.CThostFtdcReqUserLoginField{
			UserID: types.TThostFtdcUserIDType{'a'},
		},
	)
	if err != nil {
		t.Fatal(err)
	}

	req, err = state.MakeRequest(
		factory, &future.CThostFtdcReqUserLoginField{
			BrokerID: types.TThostFtdcBrokerIDType{'b'},
		},
	)
	if err != nil {
		t.Fatal(err)
	}

	if err = factory.DoRequest(req); err != nil {
		t.Fatal(err)
	}

	req, err = state.MakeRequest(
		factory, &future.CThostFtdcReqUserLoginField{},
	)
	if err != nil {
		t.Fatal(err)
	}
}
