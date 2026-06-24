package trader

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"path/filepath"
	"runtime"
	"sync"
	"sync/atomic"

	"github.com/frozenpine/ctp4go/state"
	"github.com/frozenpine/ctp4go/thost"
)

type TraderApi struct {
	thost.TraderLogSpi

	rootCtx   context.Context
	apiCtx    context.Context
	apiCancel context.CancelFunc
	cfg       traderCfg

	initOnce  sync.Once
	finalOnce sync.Once

	state *state.FlagResponsor[traderState]

	api       thost.TraderApi
	requestID atomic.Int32
}

func NewTraderApi(
	ctx context.Context, options ...cfgOpt,
) (*TraderApi, error) {
	if ctx == nil {
		ctx = context.Background()
	}

	var (
		libPath string
		err     error
	)
	switch runtime.GOOS {
	case "windows":
		libPath = filepath.Join(".", "libs", "thosttraderapi_se.dll")
	case "linux":
		libPath = filepath.Join(".", "libs", "thosttraderapi_se.so")
	default:
		return nil, fmt.Errorf(
			"%w: unsupported platform %s",
			thost.ErrApiCreateFailed, runtime.GOOS,
		)
	}
	if libPath, err = filepath.Abs(libPath); err != nil {
		return nil, errors.Join(thost.ErrApiCreateFailed, err)
	}

	trader := TraderApi{
		rootCtx: ctx,
		TraderLogSpi: thost.TraderLogSpi{
			Logger: slog.Default(),
		},
		cfg: traderCfg{
			libPath:  libPath,
			flowMode: thost.DEFAULT_FLOW_MODE,
			flowPath: thost.DEFAULT_FLOW_PATH,
		},
		state: state.NewFlagResponsor[traderState]("state"),
	}

	for _, opt := range options {
		if opt == nil {
			continue
		}

		if err := opt(&trader.cfg); err != nil {
			return nil, errors.Join(thost.ErrApiCreateFailed, err)
		}
	}

	if err := trader.createApi(); err != nil {
		return nil, err
	}

	return &trader, nil
}

func (td *TraderApi) createApi() error {
	maker := thost.GetTraderMaker()
	if maker == nil {
		return fmt.Errorf(
			"%w: no version maker found", thost.ErrCreatorMissing,
		)
	}

	api, err := maker.TraderMaker(
		td.cfg.libPath, td.cfg.flowPath,
		thost.Param{Key: thost.ParamRunMode, Value: !td.cfg.isTest},
	)()
	if err != nil {
		return errors.Join(thost.ErrApiCreateFailed, err)
	}

	td.apiCtx, td.apiCancel = context.WithCancel(td.rootCtx)
	td.api = api

	return nil
}

func (td *TraderApi) Initialize(options ...traderOpt) (err error) {
	td.initOnce.Do(func() {
		for _, opt := range options {
			if opt == nil {
				continue
			}

			if err = opt(td); err != nil {
				return
			}
		}

		td.api.RegisterSpi(td)

		td.Info(
			"initializing connection params",
			slog.String("flow_mode", td.cfg.flowMode.String()),
			slog.Int("flow_seq", td.cfg.flowSeq),
			slog.Any("front_addrs", td.cfg.frontAddrs),
			slog.Any("name_svrs", td.cfg.nameSvrs),
		)

		td.api.SubscribePrivateTopic(td.cfg.flowMode, td.cfg.flowSeq)
		td.api.SubscribePublicTopic(td.cfg.flowMode)

		if len(td.cfg.nameSvrs) > 0 {
			fens := thost.CThostFtdcFensUserInfoField{
				LoginMode: td.cfg.fensMode,
			}
			fens.BrokerID.SetString(td.cfg.brokerID)
			fens.UserID.SetString(td.cfg.userID)

			td.api.RegisterFensUserInfo(&fens)

			for _, v := range td.cfg.nameSvrs {
				td.api.RegisterNameServer(v)
			}
		} else if len(td.cfg.frontAddrs) < 1 {
			err = fmt.Errorf(
				"%w: no fronts or nameservers", thost.ErrInvalidArgs,
			)
			return
		}

		for _, v := range td.cfg.frontAddrs {
			td.api.RegisterFront(v)
		}

		td.api.Init()

		err = td.state.SetFlag(Initialized)
	})

	return
}

func (td *TraderApi) Finalize() (err error) {
	td.finalOnce.Do(func() {
		td.api.Release()

		err = td.migrateState(Finalized)
	})

	return
}

func (td *TraderApi) Reset() error {
	td.Finalize()

	return nil
}

func (td *TraderApi) Authenticate() error {
	auth := thost.CThostFtdcReqAuthenticateField{}
	auth.BrokerID.SetString(td.cfg.brokerID)
	auth.UserID.SetString(td.cfg.userID)
	auth.AppID.SetString(td.cfg.appID)
	auth.AuthCode.SetString(td.cfg.authCode)

	rtn := td.api.ReqAuthenticate(&auth, int(td.requestID.Add(1)))

	return thost.Rtn{Code: rtn}.Error()
}

func (td *TraderApi) Login() error {
	login := thost.CThostFtdcReqUserLoginField{}
	login.BrokerID.SetString(td.cfg.brokerID)
	login.UserID.SetString(td.cfg.userID)
	login.Password.SetString(td.cfg.userPass)

	rtn := td.api.ReqUserLogin(&login, int(td.requestID.Add(1)))

	return thost.Rtn{Code: rtn}.Error()
}

func (td *TraderApi) migrateState(v traderState) error {
	err := td.state.SetFlag(v)
	if err != nil {
		td.Error(
			"migrate state failed",
			slog.Any("error", err),
		)
	}
	return err
}

func (td *TraderApi) OnFrontConnected() {
	defer td.migrateState(Connected)

	td.api.GetFrontInfo(&td.FrontInfo)

	td.TraderLogSpi.OnFrontConnected()
}

func (td *TraderApi) OnFrontDisconnected(nReason int) {
	defer td.migrateState(Disconnected)

	td.TraderLogSpi.OnFrontDisconnected(nReason)
}

func (td *TraderApi) OnRspAuthenticate(
	pRspAuthenticateField *thost.CThostFtdcRspAuthenticateField,
	pRspInfo *thost.CThostFtdcRspInfoField,
	nRequestID int, bIsLast bool,
) {
	defer func() {
		if pRspInfo.ErrorID == 0 {
			td.migrateState(AuthSuccess)
		} else {
			td.migrateState(AuthFailed)
		}
	}()

	td.TraderLogSpi.OnRspAuthenticate(
		pRspAuthenticateField, pRspInfo, nRequestID, bIsLast,
	)
}

func (td *TraderApi) OnRspUserLogin(
	pRspUserLogin *thost.CThostFtdcRspUserLoginField,
	pRspInfo *thost.CThostFtdcRspInfoField,
	nRequestID int, bIsLast bool,
) {
	defer func() {
		if pRspInfo.ErrorID == 0 {
			td.migrateState(LoginSuccess)
		} else {
			td.migrateState(LoginFailed)
		}
	}()

	td.TraderLogSpi.OnRspUserLogin(
		pRspUserLogin, pRspInfo, nRequestID, bIsLast,
	)
}
