package v1_7_5

import (
	"fmt"

	"github.com/frozenpine/ctp4go/thost"
	"github.com/frozenpine/ctp4go/thost/mini"
)

func sdkMaker(
	libPath string,
	params ...thost.Param,
) func() (mini.MdApi, error) {
	return func() (mini.MdApi, error) {
		if libPath == "" {
			return nil, fmt.Errorf(
				"%w: lib path is empty", thost.ErrInvalidArgs,
			)
		}

		var (
			FlowPath    string
			IsUsingUdp  bool
			IsMulticast bool

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
			case thost.ParamIsUsingUdp:
				if IsUsingUdp, ok = p.Value.(bool); !ok {
					return nil, fmt.Errorf(
						"%w: invalid %s value %+v",
						thost.ErrInvalidArgs, p.Key, p.Value,
					)
				}
			case thost.ParamIsMulticast:
				if IsMulticast, ok = p.Value.(bool); !ok {
					return nil, fmt.Errorf(
						"%w: invalid %s value %+v",
						thost.ErrInvalidArgs, p.Key, p.Value,
					)
				}
			}
		}

		return CreateThostFtdcMdApi(
			libPath, FlowPath, IsUsingUdp, IsMulticast,
		)
	}
}

func init() {
	if err := thost.SetSdkMaker(
		"mini", "mduser",
		"v1.7.5", sdkMaker,
	); err != nil {
		panic(err)
	}
}
