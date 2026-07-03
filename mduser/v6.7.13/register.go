package v6_7_13

import (
	"fmt"

	"github.com/frozenpine/ctp4go/thost"
)

func sdkMaker(
	libPath string,
	params ...thost.Param,
) func() (thost.MdApi, error) {
	return func() (thost.MdApi, error) {
		if libPath == "" {
			return nil, fmt.Errorf(
				"%w: lib path is empty", thost.ErrInvalidArgs,
			)
		}

		var (
			FlowPath         string
			IsUsingUdp       bool
			IsMulticast      bool
			IsProductinoMode bool

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
			case thost.ParamIsProductionMode:
				if IsProductinoMode, ok = p.Value.(bool); !ok {
					return nil, fmt.Errorf(
						"%w: invalid %s value %+v",
						thost.ErrInvalidArgs, p.Key, p.Value,
					)
				}
			}
		}

		return CreateThostFtdcMdApi(
			libPath, FlowPath, IsUsingUdp, IsMulticast, IsProductinoMode,
		)
	}
}

func init() {
	if err := thost.SetMduserMaker("v6.7.13", sdkMaker); err != nil {
		panic(err)
	}
}
