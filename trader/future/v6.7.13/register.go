package v6_7_13

import (
	"fmt"

	"github.com/frozenpine/ctp4go/thost"
	"github.com/frozenpine/ctp4go/thost/future"
)

func sdkMaker(
	libPath string,
	params ...future.Param,
) func() (future.TraderApi, error) {
	return func() (future.TraderApi, error) {
		if libPath == "" {
			return nil, fmt.Errorf(
				"%w: lib path is empty", thost.ErrInvalidArgs,
			)
		}

		var (
			FlowPath         string
			IsProductionMode bool

			ok bool
		)

		for _, p := range params {
			switch p.Key {
			case future.ParamFlowPath:
				if FlowPath, ok = p.Value.(string); !ok {
					return nil, fmt.Errorf(
						"%w: invalid %s value %+v",
						thost.ErrInvalidArgs, p.Key, p.Value,
					)
				}
			case future.ParamIsProductionMode:
				if IsProductionMode, ok = p.Value.(bool); !ok {
					return nil, fmt.Errorf(
						"%w: invalid %s value %+v",
						thost.ErrInvalidArgs, p.Key, p.Value,
					)
				}
			}
		}

		return CreateThostFtdcTraderApi(
			libPath, FlowPath, IsProductionMode,
		)
	}
}

func init() {
	if err := future.SetTraderMaker("v6.7.13", sdkMaker); err != nil {
		panic(err)
	}
}
