package mduser

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

type MduserApi struct {
	thost.ThostLogSpi

	rootCtx   context.Context
	apiCtx    context.Context
	apiCancel context.CancelFunc
	cfg       mduserCfg

	initOnce  sync.Once
	initOpts  []traderOpt
	finalOnce sync.Once

	state *state.FlagResponsor[mduserState]

	api       thost.MdApi
	requestID atomic.Int32
}

func NewMduserApi(
	ctx context.Context, options ...cfgOpt,
) (*MduserApi, error) {
	if ctx == nil {
		ctx = context.Background()
	}

	var (
		libPath string
		err     error
	)
	switch runtime.GOOS {
	case "windows":
		libPath = filepath.Join(".", "libs", "thostmduserapi_se.dll")
	case "linux":
		libPath = filepath.Join(".", "libs", "thostmduserapi_se.so")
	default:
		return nil, fmt.Errorf(
			"%w: unsupported platform %s",
			thost.ErrApiCreateFailed, runtime.GOOS,
		)
	}
	if libPath, err = filepath.Abs(libPath); err != nil {
		return nil, errors.Join(thost.ErrApiCreateFailed, err)
	}

	mduser := MduserApi{
		rootCtx: ctx,
		ThostLogSpi: thost.ThostLogSpi{
			Logger: slog.Default(),
		},
		cfg: mduserCfg{
			libPath:  libPath,
			flowPath: thost.DEFAULT_FLOW_PATH,
		},
	}

	for _, opt := range options {
		if opt == nil {
			continue
		}

		if err := opt(&mduser.cfg); err != nil {
			return nil, errors.Join(thost.ErrApiCreateFailed, err)
		}
	}

	if err := mduser.createApi(); err != nil {
		return nil, err
	}

	return &mduser, nil
}

func (md *MduserApi) createApi() error {
	maker := thost.GetMduserMaker()
	if maker == nil {
		return fmt.Errorf(
			"%w: no version maker found", thost.ErrCreatorMissing,
		)
	}

	api, err := maker.MduserMaker(
		md.cfg.libPath,
		thost.Param{Key: thost.ParamFlowPath, Value: md.cfg.flowPath},
		thost.Param{Key: thost.ParamUseUDP, Value: md.cfg.isUdp},
		thost.Param{Key: thost.ParamUseMulticast, Value: md.cfg.isMulti},
		thost.Param{Key: thost.ParamRunMode, Value: !md.cfg.isTest},
	)()
	if err != nil {
		return errors.Join(thost.ErrApiCreateFailed, err)
	}

	md.state = state.NewFlagResponsor[mduserState]("state")
	md.apiCtx, md.apiCancel = context.WithCancel(md.rootCtx)
	md.api = api
	md.requestID.Store(0)

	return md.state.SetFlag(Created)
}

func (md *MduserApi) Initialize(options ...traderOpt) (err error) {
	md.initOnce.Do(func() {
		md.Info("initializing trader api")

		for _, opt := range options {
			if opt == nil {
				continue
			}

			if err = opt(md); err != nil {
				return
			}
		}

		md.initOpts = options

		md.api.RegisterSpi(md)

		md.Info(
			"initializing connection params",
			slog.Any("front_addrs", md.cfg.frontAddrs),
			slog.Any("name_svrs", md.cfg.nameSvrs),
		)

		if len(md.cfg.nameSvrs) > 0 {
			fens := thost.CThostFtdcFensUserInfoField{
				LoginMode: md.cfg.fensMode,
			}
			fens.BrokerID.SetString(md.cfg.brokerID)
			fens.UserID.SetString(md.cfg.userID)

			md.api.RegisterFensUserInfo(&fens)

			for _, v := range md.cfg.nameSvrs {
				md.api.RegisterNameServer(v)
			}
		} else if len(md.cfg.frontAddrs) < 1 {
			err = fmt.Errorf(
				"%w: no fronts or nameservers", thost.ErrInvalidArgs,
			)
			return
		}

		for _, v := range md.cfg.frontAddrs {
			md.api.RegisterFront(v)
		}

		md.api.Init()

		err = md.state.SetFlag(Initialized)
	})

	return
}

func (md *MduserApi) Finalize() (err error) {
	md.finalOnce.Do(func() {
		defer md.api.Release()

		md.Info("finalizing trader api")

		md.apiCancel()

		err = md.migrateState(Finalized)
	})

	return
}

func (md *MduserApi) Reset(options ...traderOpt) error {
	md.Finalize()

	md.Info("reseting trader api")
	md.initOnce = sync.Once{}
	md.finalOnce = sync.Once{}

	if err := md.createApi(); err != nil {
		return err
	}

	if len(options) > 0 {
		return md.Initialize(options...)
	}

	return md.Initialize(md.initOpts...)
}

func (md *MduserApi) Login() error {
	login := thost.CThostFtdcReqUserLoginField{}
	login.BrokerID.SetString(md.cfg.brokerID)
	login.UserID.SetString(md.cfg.userID)
	login.Password.SetString(md.cfg.userPass)

	rtn := md.api.ReqUserLogin(&login, int(md.requestID.Add(1)))

	return thost.Rtn{Code: rtn}.Error()
}

func (md *MduserApi) migrateState(v mduserState) error {
	err := md.state.SetFlag(v)
	if err != nil {
		md.Error(
			"migrate state failed",
			slog.Any("error", err),
		)
	}
	return err
}

func (md *MduserApi) OnFrontConnected() {
	defer md.migrateState(Connected)

	md.ThostLogSpi.OnFrontConnected()
}

func (md *MduserApi) OnFrontDisconnected(nReason int) {
	defer md.migrateState(Disconnected)

	md.ThostLogSpi.OnFrontDisconnected(nReason)
}

func (md *MduserApi) OnRspUserLogin(
	pRspUserLogin *thost.CThostFtdcRspUserLoginField,
	pRspInfo *thost.CThostFtdcRspInfoField,
	nRequestID int, bIsLast bool,
) {
	defer func() {
		if pRspInfo.ErrorID == 0 {
			md.migrateState(LoginSuccess)
		} else {
			md.migrateState(LoginFailed)
		}
	}()

	md.ThostLogSpi.OnRspUserLogin(
		pRspUserLogin, pRspInfo, nRequestID, bIsLast,
	)
}
