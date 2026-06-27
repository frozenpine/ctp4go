package v6_7_13

import (
	"fmt"

	"github.com/frozenpine/ctp4go/thost"
)

func thostCreator(
	libPath, flowPath string, params ...thost.Param,
) func() (thost.TraderApi, error) {
	for _, v := range params {
		if v.Key != thost.ParamRunMode {
			continue
		}

		mode, ok := v.Value.(bool)
		if !ok {
			return func() (thost.TraderApi, error) {
				return nil, fmt.Errorf(
					"%w: invalid run mode value: %+v",
					thost.ErrInvalidArgs, v.Value,
				)
			}
		}

		return func() (thost.TraderApi, error) {
			return CreateThostFtdcTraderApi(
				libPath, flowPath, mode,
			)
		}
	}

	return func() (thost.TraderApi, error) {
		return nil, fmt.Errorf(
			"%w: no run mode specified", thost.ErrInvalidArgs,
		)
	}
}

func init() {
	if err := thost.SetTraderMaker("v6.7.13", thostCreator); err != nil {
		panic(err)
	}
}
