package v1_7_5

import (
	"fmt"

	"github.com/frozenpine/ctp4go/thost"
	"github.com/frozenpine/ctp4go/thost/mini"
)

func sdkMaker(
	libPath string,
	params ...thost.Param,
) func() (mini.TraderApi, error) {
	return func() (mini.TraderApi, error) {
		if libPath == "" {
			return nil, fmt.Errorf(
				"%w: lib path is empty", thost.ErrInvalidArgs,
			)
		}

		var (
			FlowPath string

			ok bool
		)

		for _, p := range params {
			switch p.Key {
			case thost.ParamFlowPath:
				if FlowPath, ok = p.Value.(string); !ok {
					return nil, fmt.Errorf(
						"%w: invalid %s value %+v",
						thost.ErrInvalidArgs, p.Key, p.Value,
					)
				}
			}
		}

		return CreateThostFtdcTraderApi(
			libPath, FlowPath,
		)
	}
}

func init() {
	if err := thost.SetSdkMaker(
		"mini", "trader",
		"v1.7.5", sdkMaker,
	); err != nil {
		panic(err)
	}
}
