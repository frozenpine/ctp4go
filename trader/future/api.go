package future

import (
	"context"
	"errors"
	"fmt"
	"iter"
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

	front future.CThostFtdcFrontInfoField
	state *state.FlagResponsor[traderState]

	investors *state.DataCache[
		future.CThostFtdcInvestorField,
		*future.CThostFtdcInvestorField,
	]
	accounts *state.DataCache[
		future.CThostFtdcInvestorAccountField,
		*future.CThostFtdcInvestorAccountField,
	]
	orders *state.DataCache[
		future.CThostFtdcOrderField,
		*future.CThostFtdcOrderField,
	]
	trades *state.DataCache[
		future.CThostFtdcTradeField,
		*future.CThostFtdcTradeField,
	]
	positions *state.DataCache[
		future.CThostFtdcInvestorPositionField,
		*future.CThostFtdcInvestorPositionField,
	]
	instruments *state.DataCache[
		future.CThostFtdcInstrumentField,
		*future.CThostFtdcInstrumentField]

	caches map[string]state.Cache

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

func makeCache[
	T thost.ThostData, Ptr state.DataPtr[T],
](
	td *TraderApi, name string,
	options state.DataOptions[T, Ptr],
) (*state.DataCache[T, Ptr], error) {
	cache, err := state.NewDataCache(name, options...)
	if err != nil {
		return nil, err
	}

	if _, exist := td.caches[name]; exist {
		return nil, errors.New("cache name dumplicated")
	}

	if rd, err := state.NewReadOnlyCache(cache); err != nil {
		return nil, err
	} else {
		td.caches[name] = rd
		return cache, nil
	}
}

func (td *TraderApi) makeInvestorCache(name string) (err error) {
	td.investors, err = makeCache(
		td, name, state.DataOptions[
			future.CThostFtdcInvestorField,
			*future.CThostFtdcInvestorField,
		]{
			state.WithIdentifier(
				"Investor", func(inv *future.CThostFtdcInvestorField) string {
					return inv.String()
				},
			),
		},
	)

	return
}

func (td *TraderApi) makeOrderCache(name string) (err error) {
	td.orders, err = makeCache(
		td, name, state.DataOptions[
			future.CThostFtdcOrderField,
			*future.CThostFtdcOrderField,
		]{
			state.WithIdentifier(
				"Order", func(ord *future.CThostFtdcOrderField) string {
					return ord.OrderSysID.String()
				},
			),
			state.WithIdentifier(
				"Ref", func(ord *future.CThostFtdcOrderField) string {
					return fmt.Sprintf(
						"%s@%d.%d",
						ord.OrderRef.String(), ord.FrontID, ord.SessionID,
					)
				},
			),
			state.WithMerger(func(
				dst, src *future.CThostFtdcOrderField,
			) error {
				if src.OrderSysID != dst.OrderSysID {
					return fmt.Errorf(
						"%w: dst[%s] src[%s]",
						state.ErrCacheDataMismatch,
						dst.OrderSysID.String(),
						src.OrderSysID.String(),
					)
				}

				switch dst.OrderStatus {
				case types.THOST_FTDC_OST_AllTraded,
					types.THOST_FTDC_OST_PartTradedNotQueueing,
					types.THOST_FTDC_OST_NoTradeNotQueueing,
					types.THOST_FTDC_OST_Canceled:
					slog.Warn(
						"cached order already in final state",
						slog.Any("order", dst),
					)
					return nil
				}

				dst.OrderStatus = src.OrderStatus
				dst.VolumeTotal = src.VolumeTotal
				dst.VolumeTraded = src.VolumeTraded
				dst.ForceCloseReason = src.ForceCloseReason
				dst.OrderSource = src.OrderSource
				dst.CancelTime = src.CancelTime
				dst.ActiveTraderID = src.ActiveTraderID
				dst.ActiveUserID = src.ActiveUserID
				dst.ZCETotalTradedVolume = src.ZCETotalTradedVolume

				return nil
			}),
		},
	)

	return
}

func (td *TraderApi) makeTradeCache(name string) (err error) {
	td.trades, err = makeCache(
		td, name, state.DataOptions[
			future.CThostFtdcTradeField,
			*future.CThostFtdcTradeField,
		]{
			state.WithIdentifier(
				"Trade", func(td *future.CThostFtdcTradeField) string {
					return td.TradeID.String()
				},
			),
		},
	)

	return
}

func (td *TraderApi) makePositionCache(name string) (err error) {
	td.positions, err = makeCache(
		td, name, state.DataOptions[
			future.CThostFtdcInvestorPositionField,
			*future.CThostFtdcInvestorPositionField,
		]{
			state.WithIdentifier(
				"Position", func(pos *future.CThostFtdcInvestorPositionField) string {
					return fmt.Sprintf(
						"%s.%s.%s.%s",
						pos.ExchangeID.String(), pos.InstrumentID.String(),
						pos.PosiDirection.String(), pos.HedgeFlag.String(),
					)
				},
			),
		},
	)

	return
}

func (td *TraderApi) makeInstrumentCache(name string) (err error) {
	td.instruments, err = makeCache(
		td, name, state.DataOptions[
			future.CThostFtdcInstrumentField,
			*future.CThostFtdcInstrumentField,
		]{
			state.WithIdentifier(
				"Symbol", func(ins *future.CThostFtdcInstrumentField) string {
					return fmt.Sprintf(
						"%s.%s",
						ins.ExchangeID.String(),
						ins.InstrumentID.String(),
					)
				},
			),
		},
	)

	return
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
	td.caches = make(map[string]state.Cache)

	if err = td.makeInvestorCache("investors"); err != nil {
		return err
	}

	if err = td.makeOrderCache("orders"); err != nil {
		return err
	}

	if err = td.makeTradeCache("trades"); err != nil {
		return err
	}

	if err = td.makePositionCache("positions"); err != nil {
		return err
	}

	if err = td.makeInstrumentCache("instruments"); err != nil {
		return err
	}

	td.requests = state.NewRequestFactory(td.apiCtx, api)

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
		td.requests.Api.SubscribePublicTopic(int(types.THOST_TERT_QUICK))

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

func (td *TraderApi) GetInstrument(idt string) (
	*state.DataContainer[
		future.CThostFtdcInstrumentField,
		*future.CThostFtdcInstrumentField,
	], error,
) {
	return td.instruments.GetByKey(idt)
}

func (td *TraderApi) IterInstruments(
	filters ...func(*state.DataContainer[
		future.CThostFtdcInstrumentField,
		*future.CThostFtdcInstrumentField,
	]) bool,
) iter.Seq2[int, *state.DataContainer[
	future.CThostFtdcInstrumentField,
	*future.CThostFtdcInstrumentField,
]] {
	return func(yield func(int, *state.DataContainer[
		future.CThostFtdcInstrumentField,
		*future.CThostFtdcInstrumentField,
	]) bool) {
		for idx, v := range td.instruments.Iter(filters...) {
			if !yield(idx, v) {
				return
			}
		}
	}
}

func (td *TraderApi) OnFrontConnected() {
	td.requests.Reset()
	defer td.migrateState(Connected)

	td.requests.Api.GetFrontInfo(&td.front)
	td.Info("thost trader front", slog.Any("front", td.front))
	// 设置查询流控
	td.requests.SetQryLimit(int(td.front.QryFreq))

	td.ThostLogSpi.OnFrontConnected()
}

func (td *TraderApi) OnFrontDisconnected(nReason int) {
	defer td.migrateState(Disconnected)

	td.Info("thost trader front", slog.Any("front", td.front))
	td.ThostLogSpi.OnFrontDisconnected(nReason)
}

func (td *TraderApi) OnRspAuthenticate(
	pRspAuthenticateField *future.CThostFtdcRspAuthenticateField,
	pRspInfo *future.CThostFtdcRspInfoField,
	nRequestID int, bIsLast bool,
) {
	defer func() {
		td.requests.Complete(nRequestID, nil, td.CheckRsp(pRspInfo))

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
	defer func() {
		td.requests.Complete(nRequestID, nil, td.CheckRsp(pRspInfo))

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

func (td *TraderApi) OnRspQryInstrument(
	pInstrument *future.CThostFtdcInstrumentField,
	pRspInfo *future.CThostFtdcRspInfoField,
	nRequestID int, bIsLast bool,
) {
	td.ThostLogSpi.OnRspQryInstrument(
		pInstrument, pRspInfo, nRequestID, bIsLast,
	)

	td.instruments.AddOrUpdate(pInstrument)

	if bIsLast {
		td.requests.Complete(
			nRequestID,
			td.caches["instruments"],
			td.CheckRsp(pRspInfo),
		)
	}
}
