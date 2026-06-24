package trader

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/frozenpine/ctp4go/state"
	"github.com/frozenpine/ctp4go/thost"
	"github.com/frozenpine/ctp4go/thost/types"
)

type traderCfg struct {
	lazyInit   bool
	isTest     bool
	libPath    string
	flowPath   string
	flowMode   types.THOST_TE_RESUME_TYPE
	flowSeq    int
	frontAddrs []string
	nameSvrs   []string
	fensMode   types.TThostFtdcLoginModeType
	brokerID   string
	userID     string
	userPass   string
	appID      string
	authCode   string
}

type cfgOpt func(*traderCfg) error

func WithTestMode() cfgOpt {
	return func(ac *traderCfg) error {
		ac.isTest = true

		return nil
	}
}

func WithLibPath(p string) cfgOpt {
	return func(tc *traderCfg) error {
		if p == "" {
			return fmt.Errorf(
				"%w: invalid lib path", thost.ErrInvalidArgs,
			)
		}

		tc.libPath = p
		return nil
	}
}

func WithFlowMode(mode types.THOST_TE_RESUME_TYPE, seq ...int) cfgOpt {
	return func(tc *traderCfg) error {
		switch mode {
		case types.THOST_TERT_RESTART, types.THOST_TERT_RESUME,
			types.THOST_TERT_QUICK, types.THOST_TERT_NONE:
		case types.THOST_TERT_RESUME_FROM_SEQ_NO:
			if len(seq) < 1 || seq[0] <= 0 {
				return fmt.Errorf(
					"%w: invalid flow seq", thost.ErrInvalidArgs,
				)
			}
			tc.flowSeq = seq[0]
		default:
			return fmt.Errorf(
				"%w: invalid flow mode: %s",
				thost.ErrInvalidArgs,
				mode.String(),
			)
		}

		tc.flowMode = mode

		return nil
	}
}

func WithFensMode(mode types.TThostFtdcLoginModeType, nsSvrs ...string) cfgOpt {
	return func(tc *traderCfg) error {
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
	return func(ac *traderCfg) error {
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
	return func(ac *traderCfg) error {
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
	return func(ac *traderCfg) error {
		if env == "" {
			return fmt.Errorf(
				"%w: broker id env key empty", thost.ErrInvalidArgs,
			)
		}

		return WithBrokerID(os.Getenv(env))(ac)
	}
}

func WithUserID(id string) cfgOpt {
	return func(ac *traderCfg) error {
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
	return func(ac *traderCfg) error {
		if env == "" {
			return fmt.Errorf(
				"%w: user id env key empty", thost.ErrInvalidArgs,
			)
		}

		return WithUserID(os.Getenv(env))(ac)
	}
}

func WithUserPass(pass string) cfgOpt {
	return func(ac *traderCfg) error {
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
	return func(ac *traderCfg) error {
		if env == "" {
			return fmt.Errorf(
				"%w: user pass env key empty", thost.ErrInvalidArgs,
			)
		}

		return WithUserPass(os.Getenv(env))(ac)
	}
}

func WithAppID(id string) cfgOpt {
	return func(ac *traderCfg) error {
		if id == "" {
			return fmt.Errorf("%w: empty app id", thost.ErrInvalidArgs)
		}

		if len(id) >= len(types.TThostFtdcAppIDType{}) {
			return fmt.Errorf(
				"%w: app id size exceeded", thost.ErrInvalidArgs,
			)
		}

		ac.appID = id
		return nil
	}
}

func WithAppIDEnvKey(env string) cfgOpt {
	return func(ac *traderCfg) error {
		if env == "" {
			return fmt.Errorf("%w: app id env key empty", thost.ErrInvalidArgs)
		}

		return WithAppID(os.Getenv(env))(ac)
	}
}

func WithAuthCode(code string) cfgOpt {
	return func(ac *traderCfg) error {
		if code == "" {
			return fmt.Errorf("%w: empty auth code", thost.ErrInvalidArgs)
		}

		if len(code) >= len(types.TThostFtdcAuthCodeType{}) {
			return fmt.Errorf(
				"%w: auth code size exceeded", thost.ErrInvalidArgs,
			)
		}

		ac.authCode = code
		return nil
	}
}

func WithAuthCodeEnvKey(env string) cfgOpt {
	return func(ac *traderCfg) error {
		if env == "" {
			return fmt.Errorf(
				"%w: auth code env key empty", thost.ErrInvalidArgs)
		}

		return WithAuthCode(os.Getenv(env))(ac)
	}
}

func WithFrontAddr(v string) cfgOpt {
	return func(tc *traderCfg) error {
		if v == "" {
			return fmt.Errorf(
				"%w: front addr empty", thost.ErrInvalidArgs,
			)
		}

		tc.frontAddrs = append(tc.frontAddrs, v)
		return nil
	}
}

type traderOpt func(*TraderApi) error

func WithTraderState(stateOpts ...stateOpt) traderOpt {
	return func(tc *TraderApi) error {
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

type stateOpt func(*state.FlagResponsor[traderState]) error

func WithStateResponsor(
	s traderState, hdl func() error, options ...state.HandlerOpt,
) stateOpt {
	return func(fr *state.FlagResponsor[traderState]) error {
		hdl, err := state.NewHandler(hdl, options...)
		if err != nil {
			return err
		}

		return fr.AddHandler(s, hdl)
	}
}
