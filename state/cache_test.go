package state_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/frozenpine/ctp4go/state"
	"github.com/frozenpine/ctp4go/thost/future"
	"github.com/frozenpine/ctp4go/thost/future/types"
)

func TestCache(t *testing.T) {
	wrapFn, err := state.MakeDataWrapper(
		state.WithIdentifier(func(v *future.CThostFtdcInputOrderField) string {
			return fmt.Sprintf(
				"%s.%s", v.ExchangeID.String(), v.InstrumentID.String(),
			)
		}),
		state.WithMerger(func(dst, src *future.CThostFtdcInputOrderField) error {
			if src == nil {
				return errors.New("src data empty")
			}

			return nil
		}),
	)
	if err != nil {
		t.Fatal(err)
	}

	data := wrapFn(&future.CThostFtdcInputOrderField{
		ExchangeID:          types.TThostFtdcExchangeIDType{'S', 'H', 'F', 'E'},
		InstrumentID:        types.TThostFtdcInstrumentIDType{'a', 'b'},
		Direction:           types.THOST_FTDC_D_Buy,
		VolumeTotalOriginal: 10,
		LimitPrice:          152364.45,
		CombOffsetFlag:      types.TThostFtdcCombOffsetFlagType{byte(types.THOST_FTDC_OF_Open)},
		CombHedgeFlag:       types.TThostFtdcCombHedgeFlagType{0, byte(types.THOST_FTDC_HF_Speculation)},
	})

	t.Log(data.GetIdentity())

	t.Log(data.GetFieldString("ExchangeID"))
	t.Log(data.GetFieldInt("ProductClass"))
	t.Log(data.GetFieldInt("Direction"))
	t.Log(data.GetFieldInt("VolumeTotalOriginal"))
	t.Log(data.GetFieldFloat("LimitPrice"))
	t.Log(data.GetFieldByte("CombOffsetFlag"))
	t.Log(data.GetFieldByte("CombHedgeFlag", 1))
	t.Log(data.GetFieldByte("CombHedgeFlag"))
}
