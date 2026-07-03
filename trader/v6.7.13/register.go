package v6_7_13

import (
	"fmt"

	"github.com/frozenpine/ctp4go/thost"
	"github.com/frozenpine/ctp4go/thost/types"
)

type apiWrapper struct {
	*ThostFtdcTraderApi
}

func (api apiWrapper) SubscribePrivateTopic(
	nResumeType types.THOST_TE_RESUME_TYPE, nSeqNo ...int,
) {
	var seq int
	if len(nSeqNo) > 0 {
		seq = nSeqNo[0]
	}
	api.ThostFtdcTraderApi.SubscribePrivateTopic(
		int(nResumeType), seq,
	)
}

func (api apiWrapper) SubscribePublicTopic(
	nResumeType types.THOST_TE_RESUME_TYPE,
) {
	api.ThostFtdcTraderApi.SubscribePublicTopic(int(nResumeType))
}

func sdkMaker(
	libPath string,
	params ...thost.Param,
) func() (thost.TraderApi, error) {
	return func() (thost.TraderApi, error) {
		if libPath == "" {
			return nil, fmt.Errorf(
				"%w: lib path is empty", thost.ErrInvalidArgs,
			)
		}

		var (
			FlowPath         string
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

		api, err := CreateThostFtdcTraderApi(
			libPath, FlowPath, IsProductinoMode,
		)
		if err != nil {
			return nil, err
		}

		return apiWrapper{api}, nil
	}
}

func init() {
	if err := thost.SetTraderMaker("v6.7.13", sdkMaker); err != nil {
		panic(err)
	}
}
