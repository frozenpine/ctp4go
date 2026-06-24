package trader

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"sync"
	"sync/atomic"
	"time"

	"github.com/frozenpine/ctp4go/state"
	"github.com/frozenpine/ctp4go/thost"
)

type TraderApi struct {
	thost.TraderLogSpi

	ctx    context.Context
	cancel context.CancelFunc
	cfg    traderCfg

	initOnce  sync.Once
	finalOnce sync.Once
	done      chan struct{}

	state *state.Flag[traderState]

	api       thost.TraderApi
	requestID atomic.Int32
}

func NewTraderApi(
	ctx context.Context, libPath string, options ...tracerOpt,
) (*TraderApi, error) {
	if libPath == "" {
		return nil, fmt.Errorf(
			"%w: no lib path specified", thost.ErrApiCreateFailed,
		)
	}

	maker := thost.GetTraderMaker()
	if maker == nil {
		return nil, fmt.Errorf(
			"%w: no version maker found", thost.ErrCreatorMissing,
		)
	}

	if ctx == nil {
		ctx = context.Background()
	}

	trader := TraderApi{
		TraderLogSpi: thost.TraderLogSpi{
			Logger: slog.Default(),
		},
		cfg: traderCfg{
			flowMode: thost.DEFAULT_FLOW_MODE,
			flowPath: thost.DEFAULT_FLOW_PATH,
		},
		state: state.NewFlag[traderState]("traderState"),
		done:  make(chan struct{}),
	}

	for _, opt := range options {
		if opt == nil {
			continue
		}

		if err := opt(&trader.cfg); err != nil {
			return nil, errors.Join(thost.ErrApiCreateFailed, err)
		}
	}

	api, err := maker.TraderMaker(
		libPath, trader.cfg.flowPath,
		thost.Param{Key: thost.ParamRunMode, Value: !trader.cfg.isTest},
	)()
	if err != nil {
		return nil, errors.Join(thost.ErrApiCreateFailed, err)
	}

	trader.ctx, trader.cancel = context.WithCancel(ctx)
	trader.api = api

	return &trader, nil
}

func (td *TraderApi) Initialize(
	timeout time.Duration, options ...tracerOpt,
) (err error) {
	td.initOnce.Do(func() {
		for _, opt := range options {
			if opt == nil {
				continue
			}

			if err = opt(&td.cfg); err != nil {
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

		td.state.SetFlag(Initialized)
	})

	return
}

func (td *TraderApi) Finalize() (err error) {
	td.finalOnce.Do(func() {
		defer close(td.done)
		td.api.Release()
	})

	return
}

func (td *TraderApi) OnFrontConnected() {
	defer td.state.SetFlag(Connected)

	td.api.GetFrontInfo(&td.FrontInfo)

	td.TraderLogSpi.OnFrontConnected()

	auth := thost.CThostFtdcReqAuthenticateField{}
	auth.BrokerID.SetString(td.cfg.brokerID)
	auth.UserID.SetString(td.cfg.userID)
	auth.AppID.SetString(td.cfg.appID)
	auth.AuthCode.SetString(td.cfg.authCode)

	td.api.ReqAuthenticate(&auth, int(td.requestID.Add(1)))
}

func (td *TraderApi) OnFrontDisconnected(nReason int) {
	defer td.state.SetFlag(Disconnected)

	td.TraderLogSpi.OnFrontDisconnected(nReason)
}

func (td *TraderApi) OnRspAuthenticate(
	pRspAuthenticateField *thost.CThostFtdcRspAuthenticateField,
	pRspInfo *thost.CThostFtdcRspInfoField,
	nRequestID int, bIsLast bool,
) {
	defer func() {
		if pRspInfo.ErrorID == 0 {
			td.state.SetFlag(AuthSuccess)
		} else {
			td.state.SetFlag(AuthFailed)
		}
	}()

	td.TraderLogSpi.OnRspAuthenticate(
		pRspAuthenticateField, pRspInfo, nRequestID, bIsLast,
	)

	if pRspInfo != nil && pRspInfo.ErrorID == 0 {
		login := thost.CThostFtdcReqUserLoginField{}
		login.BrokerID.SetString(td.cfg.brokerID)
		login.UserID.SetString(td.cfg.userID)
		login.Password.SetString(td.cfg.userPass)

		td.api.ReqUserLogin(&login, int(td.requestID.Add(1)))
	}
}

func (td *TraderApi) OnRspUserLogin(
	pRspUserLogin *thost.CThostFtdcRspUserLoginField,
	pRspInfo *thost.CThostFtdcRspInfoField,
	nRequestID int, bIsLast bool,
) {
	defer func() {
		if pRspInfo.ErrorID == 0 {
			td.state.SetFlag(LoginSuccess)
		} else {
			td.state.SetFlag(LoginFailed)
		}
	}()

	td.TraderLogSpi.OnRspUserLogin(
		pRspUserLogin, pRspInfo, nRequestID, bIsLast,
	)
}
