package mduser

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/frozenpine/ctp4go/state"
	"github.com/frozenpine/ctp4go/thost"
	"github.com/frozenpine/ctp4go/thost/types"
)

type mduserCfg struct {
	lazyInit   bool
	isTest     bool
	isUdp      bool
	isMulti    bool
	libPath    string
	flowPath   string
	frontAddrs []string
	nameSvrs   []string
	fensMode   types.TThostFtdcLoginModeType
	brokerID   string
	userID     string
	userPass   string
}

type cfgOpt func(*mduserCfg) error

func WithTestMode() cfgOpt {
	return func(ac *mduserCfg) error {
		ac.isTest = true

		return nil
	}
}

func WithUdpMode() cfgOpt {
	return func(mc *mduserCfg) error {
		mc.isUdp = true
		return nil
	}
}

func WithMulticast() cfgOpt {
	return func(mc *mduserCfg) error {
		mc.isMulti = true
		return nil
	}
}

func WithLibPath(p string) cfgOpt {
	return func(tc *mduserCfg) error {
		if p == "" {
			return fmt.Errorf(
				"%w: invalid lib path", thost.ErrInvalidArgs,
			)
		}

		tc.libPath = p
		return nil
	}
}

func WithFensMode(mode types.TThostFtdcLoginModeType, nsSvrs ...string) cfgOpt {
	return func(tc *mduserCfg) error {
		switch mode {
		case types.THOST_FTDC_LM_Trade, types.THOST_FTDC_LM_Transfer:
			if len(nsSvrs) < 1 {
				return fmt.Errorf(
					"%w: no name servers specified", thost.ErrInvalidArgs,
				)
			}
		default:
			return fmt.Errorf(
				"%w: invalid fens mode: %s",
				thost.ErrInvalidArgs, mode.String(),
			)
		}

		tc.fensMode = mode
		tc.nameSvrs = append(tc.nameSvrs, nsSvrs...)
		return nil
	}
}

func WithFlowPath(flowPath string) cfgOpt {
	return func(ac *mduserCfg) error {
		if flowPath == "" {
			return fmt.Errorf(
				"%w: empty flow path", thost.ErrInvalidArgs,
			)
		}

		flowPath, err := filepath.Abs(flowPath)
		if err != nil {
			return errors.Join(thost.ErrInvalidArgs, err)
		}

		pathInfo, err := os.Stat(flowPath)

		if err != nil && errors.Is(err, os.ErrNotExist) {
			err = os.MkdirAll(flowPath, os.ModePerm)
		}

		if err != nil {
			return errors.Join(thost.ErrInvalidArgs, err)
		} else if !pathInfo.IsDir() {
			return fmt.Errorf(
				"%w: flow path should be DIR", thost.ErrInvalidArgs,
			)
		}

		if flowPath[len(flowPath)-1] != '/' {
			flowPath += "/"
		}

		ac.flowPath = flowPath

		return nil
	}
}

func WithBrokerID(id string) cfgOpt {
	return func(ac *mduserCfg) error {
		if id == "" {
			return fmt.Errorf("%w: empty broker id", thost.ErrInvalidArgs)
		}

		if len(id) >= len(types.TThostFtdcBrokerIDType{}) {
			return fmt.Errorf(
				"%w: broker id size exceeded", thost.ErrInvalidArgs,
			)
		}

		ac.brokerID = id
		return nil
	}
}

func WithBrokerIDEnvKey(env string) cfgOpt {
	return func(ac *mduserCfg) error {
		if env == "" {
			return fmt.Errorf(
				"%w: broker id env key empty", thost.ErrInvalidArgs,
			)
		}

		return WithBrokerID(os.Getenv(env))(ac)
	}
}

func WithUserID(id string) cfgOpt {
	return func(ac *mduserCfg) error {
		if id == "" {
			return fmt.Errorf("%w: empty user id", thost.ErrInvalidArgs)
		}

		if len(id) >= len(types.TThostFtdcUserIDType{}) {
			return fmt.Errorf(
				"%w: user id size exceeded", thost.ErrInvalidArgs,
			)
		}

		ac.userID = id
		return nil
	}
}

func WithUserIDEnvKey(env string) cfgOpt {
	return func(ac *mduserCfg) error {
		if env == "" {
			return fmt.Errorf(
				"%w: user id env key empty", thost.ErrInvalidArgs,
			)
		}

		return WithUserID(os.Getenv(env))(ac)
	}
}

func WithUserPass(pass string) cfgOpt {
	return func(ac *mduserCfg) error {
		if pass == "" {
			return fmt.Errorf("%w: empty user pass", thost.ErrInvalidArgs)
		}

		if len(pass) >= len(types.TThostFtdcPasswordType{}) {
			return fmt.Errorf(
				"%w: user pass size exceeded", thost.ErrInvalidArgs,
			)
		}

		ac.userPass = pass
		return nil
	}
}

func WithUserPassEnvKey(env string) cfgOpt {
	return func(ac *mduserCfg) error {
		if env == "" {
			return fmt.Errorf(
				"%w: user pass env key empty", thost.ErrInvalidArgs,
			)
		}

		return WithUserPass(os.Getenv(env))(ac)
	}
}

func WithFrontAddr(v string) cfgOpt {
	return func(tc *mduserCfg) error {
		if v == "" {
			return fmt.Errorf(
				"%w: front addr empty", thost.ErrInvalidArgs,
			)
		}

		tc.frontAddrs = append(tc.frontAddrs, v)
		return nil
	}
}

type traderOpt func(*MduserApi) error

func WithTraderState(stateOpts ...stateOpt) traderOpt {
	return func(tc *MduserApi) error {
		if len(stateOpts) < 1 {
			return errors.New("no state opt specified")
		}

		if tc.state == nil {
			return errors.New("trader state not initialized")
		}

		for _, opt := range stateOpts {
			if opt == nil {
				continue
			}

			if err := opt(tc.state); err != nil {
				return err
			}
		}

		return nil
	}
}

type stateOpt func(*state.FlagResponsor[mduserState]) error

func WithStateResponsor(
	s mduserState, hdl func() error, options ...state.HandlerOpt,
) stateOpt {
	return func(fr *state.FlagResponsor[mduserState]) error {
		hdl, err := state.NewHandler(hdl, options...)
		if err != nil {
			return err
		}

		return fr.AddHandler(s, hdl)
	}
}
