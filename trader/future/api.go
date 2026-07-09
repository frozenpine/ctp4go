package future

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"runtime"
	"sync"

	"github.com/frozenpine/ctp4go/state"
	"github.com/frozenpine/ctp4go/thost"
	"github.com/frozenpine/ctp4go/thost/future"
	"github.com/frozenpine/ctp4go/thost/future/types"
)

type TraderApi struct {
	future.ThostLogSpi

	rootCtx   context.Context
	apiCtx    context.Context
	apiCancel context.CancelFunc
	cfg       traderCfg

	initOnce  sync.Once
	initOpts  []traderOpt
	finalOnce sync.Once

	state       *state.FlagResponsor[traderState]
	instruments *state.DataCache[
		future.CThostFtdcInstrumentField,
		*future.CThostFtdcInstrumentField]

	requests *state.RequestFactory[future.TraderApi]
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
		ThostLogSpi: future.ThostLogSpi{
			Logger: slog.Default(),
		},
		cfg: traderCfg{
			libPath:  libPath,
			flowMode: types.THOST_TERT_QUICK,
			flowPath: "./flow/",
		},
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
	maker, err := thost.GetSdkMaker[future.TraderApi]("future", "trader")

	if err != nil {
		return err
	}

	if err = os.MkdirAll(td.cfg.flowPath, os.ModePerm); err != nil {
		return errors.Join(thost.ErrInvalidArgs, err)
	}

	api, err := maker.SdkMaker(
		td.cfg.libPath,
		thost.Param{Key: thost.ParamFlowPath, Value: td.cfg.flowPath},
		thost.Param{Key: thost.ParamIsProductionMode, Value: !td.cfg.isTest},
	)()
	if err != nil {
		return errors.Join(thost.ErrApiCreateFailed, err)
	}

	td.state = state.NewFlagResponsor[traderState]("state")
	td.apiCtx, td.apiCancel = context.WithCancel(td.rootCtx)

	td.requests = state.NewRequestFactory(api)

	return td.state.SetFlag(Created)
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

func (td *TraderApi) Initialize(options ...traderOpt) (err error) {
	td.initOnce.Do(func() {
		td.Info("initializing trader api")

		for _, opt := range options {
			if opt == nil {
				continue
			}

			if err = opt(td); err != nil {
				return
			}
		}

		td.initOpts = options

		td.requests.Api.RegisterSpi(td)

		td.Info(
			"initializing connection params",
			slog.String("flow_mode", td.cfg.flowMode.String()),
			slog.Int("flow_seq", td.cfg.flowSeq),
			slog.Any("front_addrs", td.cfg.frontAddrs),
			slog.Any("name_svrs", td.cfg.nameSvrs),
		)

		td.requests.Api.SubscribePrivateTopic(
			int(td.cfg.flowMode), td.cfg.flowSeq,
		)
		td.requests.Api.SubscribePublicTopic(int(td.cfg.flowMode))

		if len(td.cfg.nameSvrs) > 0 {
			fens := future.CThostFtdcFensUserInfoField{
				LoginMode: td.cfg.fensMode,
			}
			fens.BrokerID.SetString(td.cfg.brokerID)
			fens.UserID.SetString(td.cfg.userID)

			td.requests.Api.RegisterFensUserInfo(&fens)

			for _, v := range td.cfg.nameSvrs {
				td.requests.Api.RegisterNameServer(v)
			}
		} else if len(td.cfg.frontAddrs) < 1 {
			err = fmt.Errorf(
				"%w: no fronts or nameservers", thost.ErrInvalidArgs,
			)
			return
		}

		for _, v := range td.cfg.frontAddrs {
			td.requests.Api.RegisterFront(v)
		}

		td.requests.Api.Init()

		err = td.state.SetFlag(Initialized)
	})

	return
}

func (td *TraderApi) Finalize() (err error) {
	td.finalOnce.Do(func() {
		defer td.requests.Api.Release()

		td.Info("finalizing trader api")

		td.apiCancel()

		err = td.migrateState(Finalized)
	})

	return
}

func (td *TraderApi) Reset(options ...traderOpt) error {
	td.Finalize()

	td.Info("reseting trader api")
	td.initOnce = sync.Once{}
	td.finalOnce = sync.Once{}

	if err := td.createApi(); err != nil {
		return err
	}

	if len(options) > 0 {
		return td.Initialize(options...)
	}

	return td.Initialize(td.initOpts...)
}

func (td *TraderApi) Authenticate() error {
	auth := future.CThostFtdcReqAuthenticateField{}
	auth.BrokerID.SetString(td.cfg.brokerID)
	auth.UserID.SetString(td.cfg.userID)
	auth.AppID.SetString(td.cfg.appID)
	auth.AuthCode.SetString(td.cfg.authCode)

	req, err := state.MakeRequest(td.requests, &auth)
	if err != nil {
		return err
	}

	return td.requests.DoRequest(req)
}

func (td *TraderApi) Login() error {
	login := future.CThostFtdcReqUserLoginField{}
	login.BrokerID.SetString(td.cfg.brokerID)
	login.UserID.SetString(td.cfg.userID)
	login.Password.SetString(td.cfg.userPass)

	req, err := state.MakeRequest(td.requests, &login)
	if err != nil {
		return err
	}

	return td.requests.DoRequest(req)
}

func (td *TraderApi) QueryInstruments() error {
	qry := future.CThostFtdcQryInstrumentField{}

	req, err := state.MakeRequest(td.requests, &qry)
	if err != nil {
		return err
	}

	return td.requests.DoRequest(req)
}

func (td *TraderApi) OnFrontConnected() {
	td.requests.Reset()
	defer td.migrateState(Connected)

	td.requests.Api.GetFrontInfo(&td.FrontInfo)

	td.ThostLogSpi.OnFrontConnected()
}

func (td *TraderApi) OnFrontDisconnected(nReason int) {
	defer td.migrateState(Disconnected)

	td.ThostLogSpi.OnFrontDisconnected(nReason)
}

func (td *TraderApi) OnRspAuthenticate(
	pRspAuthenticateField *future.CThostFtdcRspAuthenticateField,
	pRspInfo *future.CThostFtdcRspInfoField,
	nRequestID int, bIsLast bool,
) {
	td.requests.Complete(nRequestID)

	defer func() {
		if pRspInfo.ErrorID == 0 {
			td.migrateState(AuthSuccess)
		} else {
			td.migrateState(AuthFailed)
		}
	}()

	td.ThostLogSpi.OnRspAuthenticate(
		pRspAuthenticateField, pRspInfo, nRequestID, bIsLast,
	)
}

func (td *TraderApi) OnRspUserLogin(
	pRspUserLogin *future.CThostFtdcRspUserLoginField,
	pRspInfo *future.CThostFtdcRspInfoField,
	nRequestID int, bIsLast bool,
) {
	td.requests.Complete(nRequestID)

	defer func() {
		if pRspInfo.ErrorID == 0 {
			td.migrateState(LoginSuccess)
		} else {
			td.migrateState(LoginFailed)
		}
	}()

	td.ThostLogSpi.OnRspUserLogin(
		pRspUserLogin, pRspInfo, nRequestID, bIsLast,
	)
}
