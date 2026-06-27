package v6_7_13

import (
	"fmt"

	"github.com/frozenpine/ctp4go/thost"
)

func thostCreator(
	libPath, flowPath string, params ...thost.Param,
) func() (thost.TraderApi, error) {
	var (
		isProduct bool

		ok bool
	)

NEXT:
	for _, v := range params {
		switch v.Key {
		case thost.ParamRunMode:
			if isProduct, ok = v.Value.(bool); ok {
				continue NEXT
			}
		}

		return func() (thost.TraderApi, error) {
			return nil, fmt.Errorf(
				"%w: invalid creator args[%s]: %+v",
				thost.ErrInvalidArgs, v.Key, v.Value,
			)
		}
	}

	return func() (thost.TraderApi, error) {
		return CreateThostFtdcTraderApi(
			libPath, flowPath, isProduct,
		)
	}
}

func init() {
	if err := thost.SetTraderMaker("v6.7.13", thostCreator); err != nil {
		panic(err)
	}
}
