package thost

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strconv"

	"github.com/frozenpine/ctp4go/thost/types"
)

type apiCfg struct {
	isTest   bool
	flowPath string
	brokerID string
	userID   string
	userPass string
	appID    string
	authCode string
}

type apiOpt func(*apiCfg) error

func WithTestMode() apiOpt {
	return func(ac *apiCfg) error {
		ac.isTest = true

		return nil
	}
}

func WithModeEnvKey(env string) apiOpt {
	return func(ac *apiCfg) error {
		if env == "" {
			return fmt.Errorf("%w: test mode env key empty", ErrInvalidArgs)
		}

		flag := os.Getenv(env)
		if flag == "" {
			ac.isTest = false
			return nil
		}

		parsed, err := strconv.ParseBool(flag)
		if err != nil {
			return errors.Join(ErrInvalidArgs, err)
		}

		ac.isTest = parsed
		return nil
	}
}

func WithFlowPath(flowPath string) apiOpt {
	return func(ac *apiCfg) error {
		if flowPath == "" {
			return fmt.Errorf(
				"%w: empty flow path", ErrInvalidArgs,
			)
		}

		flowPath, err := filepath.Abs(flowPath)
		if err != nil {
			return errors.Join(ErrInvalidArgs, err)
		}

		pathInfo, err := os.Stat(flowPath)

		if err != nil && errors.Is(err, os.ErrNotExist) {
			err = os.MkdirAll(flowPath, os.ModePerm)
		}

		if err != nil {
			return errors.Join(ErrInvalidArgs, err)
		} else if !pathInfo.IsDir() {
			return fmt.Errorf(
				"%w: flow path should be DIR", ErrInvalidArgs,
			)
		}

		if flowPath[len(flowPath)-1] != '/' {
			flowPath += "/"
		}

		ac.flowPath = flowPath

		return nil
	}
}

func WithFlowPathEnvKey(env string) apiOpt {
	return func(ac *apiCfg) error {
		if env == "" {
			return fmt.Errorf("%w: flow path env key empty", ErrInvalidArgs)
		}

		return WithFlowPath(os.Getenv(env))(ac)
	}
}

func WithBrokerID(id string) apiOpt {
	return func(ac *apiCfg) error {
		if id == "" {
			return fmt.Errorf("%w: empty broker id", ErrInvalidArgs)
		}

		if len(id) >= len(types.TThostFtdcBrokerIDType{}) {
			return fmt.Errorf("%w: broker id size exceeded", ErrInvalidArgs)
		}

		ac.brokerID = id
		return nil
	}
}

func WithBrokerIDEnvKey(env string) apiOpt {
	return func(ac *apiCfg) error {
		if env == "" {
			return fmt.Errorf("%w: broker id env key empty", ErrInvalidArgs)
		}

		return WithBrokerID(os.Getenv(env))(ac)
	}
}

func WithUserID(id string) apiOpt {
	return func(ac *apiCfg) error {
		if id == "" {
			return fmt.Errorf("%w: empty user id", ErrInvalidArgs)
		}

		if len(id) >= len(types.TThostFtdcUserIDType{}) {
			return fmt.Errorf("%w: user id size exceeded", ErrInvalidArgs)
		}

		ac.userID = id
		return nil
	}
}

func WithUserIDEnvKey(env string) apiOpt {
	return func(ac *apiCfg) error {
		if env == "" {
			return fmt.Errorf("%w: user id env key empty", ErrInvalidArgs)
		}

		return WithUserID(os.Getenv(env))(ac)
	}
}

func WithUserPass(pass string) apiOpt {
	return func(ac *apiCfg) error {
		if pass == "" {
			return fmt.Errorf("%w: empty user pass", ErrInvalidArgs)
		}

		if len(pass) >= len(types.TThostFtdcPasswordType{}) {
			return fmt.Errorf("%w: user pass size exceeded", ErrInvalidArgs)
		}

		ac.userPass = pass
		return nil
	}
}

func WithUserPassEnvKey(env string) apiOpt {
	return func(ac *apiCfg) error {
		if env == "" {
			return fmt.Errorf("%w: user pass env key empty", ErrInvalidArgs)
		}

		return WithUserPass(os.Getenv(env))(ac)
	}
}

func WithAppID(id string) apiOpt {
	return func(ac *apiCfg) error {
		if id == "" {
			return fmt.Errorf("%w: empty app id", ErrInvalidArgs)
		}

		if len(id) >= len(types.TThostFtdcAppIDType{}) {
			return fmt.Errorf("%w: app id size exceeded", ErrInvalidArgs)
		}

		ac.appID = id
		return nil
	}
}

func WithAppIDEnvKey(env string) apiOpt {
	return func(ac *apiCfg) error {
		if env == "" {
			return fmt.Errorf("%w: app id env key empty", ErrInvalidArgs)
		}

		return WithAppID(os.Getenv(env))(ac)
	}
}

func WithAuthCode(code string) apiOpt {
	return func(ac *apiCfg) error {
		if code == "" {
			return fmt.Errorf("%w: empty auth code", ErrInvalidArgs)
		}

		if len(code) >= len(types.TThostFtdcAuthCodeType{}) {
			return fmt.Errorf("%w: auth code size exceeded", ErrInvalidArgs)
		}

		ac.authCode = code
		return nil
	}
}

func WithAuthCodeEnvKey(env string) apiOpt {
	return func(ac *apiCfg) error {
		if env == "" {
			return fmt.Errorf("%w: auth code env key empty", ErrInvalidArgs)
		}

		return WithAuthCode(os.Getenv(env))(ac)
	}
}
