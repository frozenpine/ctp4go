package v6_7_13

import (
	"fmt"

	"github.com/frozenpine/ctp4go/thost"
)

func thostCreator(
	libPath string, params ...thost.Param,
) func() (thost.MdApi, error) {
	var (
		flowPath                  = "./flow/"
		isUdp, isMulti, isProduct bool

		ok bool
	)

NEXT:
	for _, v := range params {
		switch v.Key {
		case thost.ParamFlowPath:
			if flowPath, ok = v.Value.(string); ok {
				continue NEXT
			}
		case thost.ParamRunMode:
			if isProduct, ok = v.Value.(bool); ok {
				continue NEXT
			}
		case thost.ParamUseUDP:
			if isUdp, ok = v.Value.(bool); ok {
				continue NEXT
			}
		case thost.ParamUseMulticast:
			if isMulti, ok = v.Value.(bool); ok {
				continue NEXT
			}
		}

		return func() (thost.MdApi, error) {
			return nil, fmt.Errorf(
				"%w: invalid creator args[%s]: %+v",
				thost.ErrInvalidArgs, v.Key, v.Value,
			)
		}
	}

	return func() (thost.MdApi, error) {
		return CreateThostFtdcMdApi(
			libPath, flowPath, isUdp, isMulti, isProduct,
		)
	}
}

func init() {
	if err := thost.SetMduserMaker("v6.7.13", thostCreator); err != nil {
		panic(err)
	}
}
