package future

import (
	"context"
	"errors"
	"fmt"
	"iter"
	"log/slog"

	"github.com/frozenpine/ctp4go"
	"github.com/frozenpine/ctp4go/state"
	"github.com/frozenpine/ctp4go/thost"
	"github.com/frozenpine/ctp4go/thost/future/types"
)

var (
	_ TraderSpi = &ThostFutureBase{}
	_ MdSpi     = &ThostFutureBase{}

	ErrCacheNotExist = errors.New("cache not exists")
)

type cacheName string

const (
	InvCache cacheName = "investors"   // 投资者缓存
	OrdCache cacheName = "orders"      // 委托缓存
	TrdCache cacheName = "trades"      // 成交缓存
	PosCache cacheName = "positions"   // 持仓缓存
	InsCache cacheName = "instruments" // 合约缓存
	MdCache  cacheName = "marketdatas" // 行情缓存
)

var cacheMakers = map[cacheName]func(*ThostFutureBase, int) error{
	InvCache: func(tls *ThostFutureBase, i int) error {
		return makeCache(
			tls, InvCache, state.DataOptions[
				CThostFtdcInvestorField,
				*CThostFtdcInvestorField,
			]{
				state.WithIdentifier(
					"Investor", func(inv *CThostFtdcInvestorField) string {
						return inv.String()
					},
				),
			},
		)
	},
	OrdCache: func(tls *ThostFutureBase, i int) error {
		return makeCache(
			tls, OrdCache, state.DataOptions[
				CThostFtdcOrderField,
				*CThostFtdcOrderField,
			]{
				state.WithIdentifier(
					"Order", func(ord *CThostFtdcOrderField) string {
						return ord.OrderSysID.String()
					},
				),
				state.WithIdentifier(
					"Ref", func(ord *CThostFtdcOrderField) string {
						return fmt.Sprintf(
							"%s@%d.%d",
							ord.OrderRef.String(), ord.FrontID, ord.SessionID,
						)
					},
				),
				state.WithMerger(func(
					dst, src *CThostFtdcOrderField,
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
	},
	TrdCache: func(tls *ThostFutureBase, i int) error {
		return makeCache(
			tls, TrdCache, state.DataOptions[
				CThostFtdcTradeField,
				*CThostFtdcTradeField,
			]{
				state.WithIdentifier(
					"Trade", func(td *CThostFtdcTradeField) string {
						return td.TradeID.String()
					},
				),
			},
		)
	},
	PosCache: func(tls *ThostFutureBase, i int) error {
		return makeCache(
			tls, PosCache, state.DataOptions[
				CThostFtdcInvestorPositionField,
				*CThostFtdcInvestorPositionField,
			]{
				state.WithIdentifier(
					"Position", func(pos *CThostFtdcInvestorPositionField) string {
						return fmt.Sprintf(
							"%s.%s.%s.%s",
							pos.ExchangeID.String(), pos.InstrumentID.String(),
							pos.PosiDirection.String(), pos.HedgeFlag.String(),
						)
					},
				),
			},
		)
	},
	InsCache: func(tls *ThostFutureBase, i int) error {
		return makeCache(
			tls, InsCache, state.DataOptions[
				CThostFtdcInstrumentField,
				*CThostFtdcInstrumentField,
			]{
				state.WithIdentifier(
					"Symbol", func(ins *CThostFtdcInstrumentField) string {
						return fmt.Sprintf(
							"%s.%s",
							ins.ExchangeID.String(),
							ins.InstrumentID.String(),
						)
					},
				),
				state.WithIdentifier(
					"InstrumentID", func(ins *CThostFtdcInstrumentField) string {
						return ins.InstrumentID.String()
					},
				),
			},
		)
	},
	MdCache: func(tfb *ThostFutureBase, i int) error {
		return makeCache(
			tfb, MdCache, state.DataOptions[
				CThostFtdcDepthMarketDataField,
				*CThostFtdcDepthMarketDataField,
			]{
				state.WithIdentifier(
					"Tick", func(md *CThostFtdcDepthMarketDataField) string {
						return fmt.Sprintf(
							"%s.%s@%s@%s.%03d",
							md.ExchangeID.String(), md.InstrumentID.String(),
							md.TradingDay.String(),
							md.UpdateTime.String(), md.UpdateMillisec,
						)
					},
				),
			},
		)
	},
}

const (
	DEFAULT_BUFF_SIZE = 10
)

type ThostFutureBase struct {
	*slog.Logger
	state.RFactory

	ctx    context.Context
	caches map[cacheName]state.Cache
}

func makeCache[
	T thost.ThostData, Ptr state.DataPtr[T],
](
	spi *ThostFutureBase, name cacheName,
	options state.DataOptions[T, Ptr],
) error {
	cache, err := state.NewDataCache(string(name), options...)
	if err != nil {
		return err
	}

	if _, exist := spi.caches[name]; exist {
		return errors.New("cache name dumplicated")
	}

	if rd, err := state.NewReadOnlyCache(cache); err != nil {
		return err
	} else {
		spi.caches[name] = rd
		return nil
	}
}

type initCfg struct {
	buffSize int
	caches   map[cacheName]struct{}
}

type initOpt func(*initCfg) error

func (spi *ThostFutureBase) Initialize(
	ctx context.Context, req state.RFactory, options ...initOpt,
) error {
	if req == nil {
		return errors.New("request factory is nil")
	}
	spi.RFactory = req

	if ctx == nil {
		ctx = context.Background()
	}
	spi.ctx = ctx

	cfg := initCfg{
		buffSize: 1 << 4,
		caches: map[cacheName]struct{}{
			InsCache: {},
		},
	}

	for _, opt := range options {
		if opt == nil {
			continue
		}

		if err := opt(&cfg); err != nil {
			return err
		}
	}

	spi.caches = make(map[cacheName]state.Cache)
	for c := range cfg.caches {
		maker, exist := cacheMakers[c]
		if !exist {
			return fmt.Errorf("%w: no maker for cache %s", ErrCacheNotExist, c)
		}

		if err := maker(spi, cfg.buffSize); err != nil {
			return err
		}
	}

	return nil
}

func (spi *ThostFutureBase) GetCacheData(
	name cacheName, idt string,
) (state.Data, error) {
	c, exist := spi.caches[name]

	if !exist {
		return nil, fmt.Errorf("%w: %s", ErrCacheNotExist, name)
	}
	return c.GetByKey(idt)
}

func (spi *ThostFutureBase) IterCacheData(
	name cacheName, filters ...func(state.Data) bool,
) iter.Seq2[int, state.Data] {
	c, exist := spi.caches[name]

	if !exist {
		return func(yield func(int, state.Data) bool) {}
	}

	return c.Iter(filters...)
}

func (spi *ThostFutureBase) CheckRsp(rsp *CThostFtdcRspInfoField) error {
	if rsp == nil {
		return nil
	}

	if rsp.ErrorID != 0 {
		return fmt.Errorf(
			"[%d] %s", rsp.ErrorID, ctp4go.DecodeGBK(rsp.ErrorMsg[:]),
		)
	}

	return nil
}

func (spi *ThostFutureBase) OnFrontConnected() {
	spi.Info("thost [OnFrontConnected]")
}

func (spi *ThostFutureBase) OnFrontDisconnected(nReason int) {
	spi.Info(
		"thost [OnFrontDisconnected]",
		slog.Int("reason", nReason),
	)
}

func (spi *ThostFutureBase) OnHeartBeatWarning(nTimeLapse int) {
	spi.Info(
		"thost [OnHeartBeatWarning]",
		slog.Int("time_lapse", nTimeLapse),
	)
}

func (spi *ThostFutureBase) OnRspAuthenticate(
	pRspAuthenticateField *CThostFtdcRspAuthenticateField,
	pRspInfo *CThostFtdcRspInfoField,
	nRequestID int, bIsLast bool,
) {
	err := spi.CheckRsp(pRspInfo)
	if err != nil {
		spi.Error(
			"thost trader rsp [OnRspAuthenticate] failed",
			slog.Any("error", err),
			slog.Int("request_id", nRequestID),
			slog.Bool("is_last", bIsLast),
		)
		return
	}

	spi.Info(
		"thost trader rsp [OnRspAuthenticate] succeeded",
		slog.Any("authenticate", pRspAuthenticateField),
	)
}

func (spi *ThostFutureBase) OnRtnPrivateSeqNo(nSeqNo int) {
	spi.Log(
		spi.ctx, slog.LevelDebug-2,
		"thost trader [OnRtnPrivateSeqNo]",
		slog.Int("seq_no", nSeqNo),
	)
}

func (spi *ThostFutureBase) OnRspUserLogin(
	pRspUserLogin *CThostFtdcRspUserLoginField,
	pRspInfo *CThostFtdcRspInfoField,
	nRequestID int, bIsLast bool,
) {
	err := spi.CheckRsp(pRspInfo)
	if err != nil {
		spi.Error(
			"thost [OnRspUserLogin] failed",
			slog.Any("error", err),
			slog.Int("request_id", nRequestID),
			slog.Bool("is_last", bIsLast),
		)
		return
	}

	spi.Info(
		"thost [OnRspUserLogin] succeeded",
		slog.Any("login", pRspUserLogin),
	)
}

func (spi *ThostFutureBase) OnRspUserLogout(
	pUserLogout *CThostFtdcUserLogoutField,
	pRspInfo *CThostFtdcRspInfoField,
	nRequestID int, bIsLast bool,
) {
	err := spi.CheckRsp(pRspInfo)
	if err != nil {
		spi.Error(
			"thost [OnRspUserLogout] failed",
			slog.Any("error", err),
			slog.Int("request_id", nRequestID),
			slog.Bool("is_last", bIsLast),
		)
		return
	}

	spi.Info(
		"thost [OnRspUserLogout] succeeded",
		slog.Any("logout", pUserLogout),
	)
}

func (spi *ThostFutureBase) OnRspUserPasswordUpdate(
	pUserPasswordUpdate *CThostFtdcUserPasswordUpdateField,
	pRspInfo *CThostFtdcRspInfoField,
	nRequestID int, bIsLast bool,
) {
	err := spi.CheckRsp(pRspInfo)
	if err != nil {
		spi.Error(
			"thost trader [OnRspUserPasswordUpdate] failed",
			slog.Any("error", err),
			slog.Int("request_id", nRequestID),
			slog.Bool("is_last", bIsLast),
		)
		return
	}

	spi.Info(
		"thost trader [OnRspUserPasswordUpdate] succeeded",
		slog.Any("data", pUserPasswordUpdate),
	)
}

func (spi *ThostFutureBase) OnRspTradingAccountPasswordUpdate(
	pTradingAccountPasswordUpdate *CThostFtdcTradingAccountPasswordUpdateField,
	pRspInfo *CThostFtdcRspInfoField,
	nRequestID int, bIsLast bool,
) {
	err := spi.CheckRsp(pRspInfo)
	if err != nil {
		spi.Error(
			"thost trader [OnRspTradingAccountPasswordUpdate] failed",
			slog.Any("error", err),
			slog.Int("request_id", nRequestID),
			slog.Bool("is_last", bIsLast),
		)
		return
	}

	spi.Info(
		"thost trader [OnRspTradingAccountPasswordUpdate] succeeded",
		slog.Any("data", pTradingAccountPasswordUpdate),
	)
}

func (spi *ThostFutureBase) OnRspUserAuthMethod(
	pRspUserAuthMethod *CThostFtdcRspUserAuthMethodField,
	pRspInfo *CThostFtdcRspInfoField,
	nRequestID int, bIsLast bool,
) {
	err := spi.CheckRsp(pRspInfo)
	if err != nil {
		spi.Error(
			"thost trader [OnRspUserAuthMethod] failed",
			slog.Any("error", err),
			slog.Int("request_id", nRequestID),
			slog.Bool("is_last", bIsLast),
		)
		return
	}

	spi.Info(
		"thost trader [OnRspUserAuthMethod] succeeded",
		slog.Any("data", pRspUserAuthMethod),
	)
}

func (spi *ThostFutureBase) OnRspGenUserCaptcha(
	pRspGenUserCaptcha *CThostFtdcRspGenUserCaptchaField,
	pRspInfo *CThostFtdcRspInfoField,
	nRequestID int, bIsLast bool,
) {
	err := spi.CheckRsp(pRspInfo)
	if err != nil {
		spi.Error(
			"thost trader [OnRspGenUserCaptcha] failed",
			slog.Any("error", err),
			slog.Int("request_id", nRequestID),
			slog.Bool("is_last", bIsLast),
		)
		return
	}

	spi.Info(
		"thost trader [OnRspGenUserCaptcha] succeeded",
		slog.Any("data", pRspGenUserCaptcha),
	)
}

func (spi *ThostFutureBase) OnRspGenUserText(
	pRspGenUserText *CThostFtdcRspGenUserTextField,
	pRspInfo *CThostFtdcRspInfoField,
	nRequestID int, bIsLast bool,
) {
	err := spi.CheckRsp(pRspInfo)
	if err != nil {
		spi.Error(
			"thost trader [OnRspGenUserText] failed",
			slog.Any("error", err),
			slog.Int("request_id", nRequestID),
			slog.Bool("is_last", bIsLast),
		)
		return
	}

	spi.Info(
		"thost trader [OnRspGenUserText] succeeded",
		slog.Any("data", pRspGenUserText),
	)
}

func (spi *ThostFutureBase) OnRspOrderInsert(
	pInputOrder *CThostFtdcInputOrderField,
	pRspInfo *CThostFtdcRspInfoField,
	nRequestID int, bIsLast bool,
) {
	err := spi.CheckRsp(pRspInfo)
	if err != nil {
		spi.Error(
			"thost trader [OnRspOrderInsert] failed",
			slog.Any("error", err),
			slog.Int("request_id", nRequestID),
			slog.Bool("is_last", bIsLast),
		)
		return
	}

	spi.Info(
		"thost trader [OnRspOrderInsert] succeeded",
		slog.Any("data", pInputOrder),
	)
}

func (spi *ThostFutureBase) OnRspParkedOrderInsert(
	pParkedOrder *CThostFtdcParkedOrderField,
	pRspInfo *CThostFtdcRspInfoField,
	nRequestID int, bIsLast bool,
) {
	err := spi.CheckRsp(pRspInfo)
	if err != nil {
		spi.Error(
			"thost trader [OnRspParkedOrderInsert] failed",
			slog.Any("error", err),
			slog.Int("request_id", nRequestID),
			slog.Bool("is_last", bIsLast),
		)
		return
	}

	spi.Info(
		"thost trader [OnRspParkedOrderInsert] succeeded",
		slog.Any("data", pParkedOrder),
	)
}

func (spi *ThostFutureBase) OnRspParkedOrderAction(
	pParkedOrderAction *CThostFtdcParkedOrderActionField,
	pRspInfo *CThostFtdcRspInfoField,
	nRequestID int, bIsLast bool,
) {
	err := spi.CheckRsp(pRspInfo)
	if err != nil {
		spi.Error(
			"thost trader [OnRspParkedOrderAction] failed",
			slog.Any("error", err),
			slog.Int("request_id", nRequestID),
			slog.Bool("is_last", bIsLast),
		)
		return
	}

	spi.Info(
		"thost trader [OnRspParkedOrderAction] succeeded",
		slog.Any("data", pParkedOrderAction),
	)
}

func (spi *ThostFutureBase) OnRspOrderAction(
	pInputOrderAction *CThostFtdcInputOrderActionField,
	pRspInfo *CThostFtdcRspInfoField,
	nRequestID int, bIsLast bool,
) {
	err := spi.CheckRsp(pRspInfo)
	if err != nil {
		spi.Error(
			"thost trader [OnRspOrderAction] failed",
			slog.Any("error", err),
			slog.Int("request_id", nRequestID),
			slog.Bool("is_last", bIsLast),
		)
		return
	}

	spi.Info(
		"thost trader [OnRspOrderAction] succeeded",
		slog.Any("data", pInputOrderAction),
	)
}

func (spi *ThostFutureBase) OnRspQryMaxOrderVolume(
	pQryMaxOrderVolume *CThostFtdcQryMaxOrderVolumeField,
	pRspInfo *CThostFtdcRspInfoField,
	nRequestID int, bIsLast bool,
) {
	err := spi.CheckRsp(pRspInfo)
	if err != nil {
		spi.Error(
			"thost trader [OnRspQryMaxOrderVolume] failed",
			slog.Any("error", err),
			slog.Int("request_id", nRequestID),
			slog.Bool("is_last", bIsLast),
		)
		return
	}

	spi.Log(
		spi.ctx, slog.LevelDebug-2,
		"thost trader [OnRspQryMaxOrderVolume] succeeded",
		slog.Any("data", pQryMaxOrderVolume),
	)
}

func (spi *ThostFutureBase) OnRspSettlementInfoConfirm(
	pSettlementInfoConfirm *CThostFtdcSettlementInfoConfirmField,
	pRspInfo *CThostFtdcRspInfoField,
	nRequestID int, bIsLast bool,
) {
	err := spi.CheckRsp(pRspInfo)
	if err != nil {
		spi.Error(
			"thost trader [OnRspSettlementInfoConfirm] failed",
			slog.Any("error", err),
			slog.Int("request_id", nRequestID),
			slog.Bool("is_last", bIsLast),
		)
		return
	}

	spi.Info(
		"thost trader [OnRspSettlementInfoConfirm] succeeded",
		slog.Any("data", pSettlementInfoConfirm),
	)
}

func (spi *ThostFutureBase) OnRspRemoveParkedOrder(
	pRemoveParkedOrder *CThostFtdcRemoveParkedOrderField,
	pRspInfo *CThostFtdcRspInfoField,
	nRequestID int, bIsLast bool,
) {
	err := spi.CheckRsp(pRspInfo)
	if err != nil {
		spi.Error(
			"thost trader [OnRspRemoveParkedOrder] failed",
			slog.Any("error", err),
			slog.Int("request_id", nRequestID),
			slog.Bool("is_last", bIsLast),
		)
		return
	}

	spi.Info(
		"thost trader [OnRspRemoveParkedOrder] succeeded",
		slog.Any("data", pRemoveParkedOrder),
	)
}

func (spi *ThostFutureBase) OnRspRemoveParkedOrderAction(
	pRemoveParkedOrderAction *CThostFtdcRemoveParkedOrderActionField,
	pRspInfo *CThostFtdcRspInfoField,
	nRequestID int, bIsLast bool,
) {
	err := spi.CheckRsp(pRspInfo)
	if err != nil {
		spi.Error(
			"thost trader [OnRspRemoveParkedOrderAction] failed",
			slog.Any("error", err),
			slog.Int("request_id", nRequestID),
			slog.Bool("is_last", bIsLast),
		)
		return
	}

	spi.Info(
		"thost trader [OnRspRemoveParkedOrderAction] succeeded",
		slog.Any("data", pRemoveParkedOrderAction),
	)
}

func (spi *ThostFutureBase) OnRspExecOrderInsert(
	pInputExecOrder *CThostFtdcInputExecOrderField,
	pRspInfo *CThostFtdcRspInfoField,
	nRequestID int, bIsLast bool,
) {
	err := spi.CheckRsp(pRspInfo)
	if err != nil {
		spi.Error(
			"thost trader [OnRspExecOrderInsert] failed",
			slog.Any("error", err),
			slog.Int("request_id", nRequestID),
			slog.Bool("is_last", bIsLast),
		)
		return
	}

	spi.Info(
		"thost trader [OnRspExecOrderInsert] succeeded",
		slog.Any("data", pInputExecOrder),
	)
}

func (spi *ThostFutureBase) OnRspExecOrderAction(
	pInputExecOrderAction *CThostFtdcInputExecOrderActionField,
	pRspInfo *CThostFtdcRspInfoField,
	nRequestID int, bIsLast bool,
) {
	err := spi.CheckRsp(pRspInfo)
	if err != nil {
		spi.Error(
			"thost trader [OnRspExecOrderAction] failed",
			slog.Any("error", err),
			slog.Int("request_id", nRequestID),
			slog.Bool("is_last", bIsLast),
		)
		return
	}

	spi.Info(
		"thost trader [OnRspExecOrderAction] succeeded",
		slog.Any("data", pInputExecOrderAction),
	)
}

func (spi *ThostFutureBase) OnRspForQuoteInsert(
	pInputForQuote *CThostFtdcInputForQuoteField,
	pRspInfo *CThostFtdcRspInfoField,
	nRequestID int, bIsLast bool,
) {
	err := spi.CheckRsp(pRspInfo)
	if err != nil {
		spi.Error(
			"thost trader [OnRspForQuoteInsert] failed",
			slog.Any("error", err),
			slog.Int("request_id", nRequestID),
			slog.Bool("is_last", bIsLast),
		)
		return
	}

	spi.Info(
		"thost trader [OnRspForQuoteInsert] succeeded",
		slog.Any("data", pInputForQuote),
	)
}

func (spi *ThostFutureBase) OnRspQuoteInsert(
	pInputQuote *CThostFtdcInputQuoteField,
	pRspInfo *CThostFtdcRspInfoField,
	nRequestID int, bIsLast bool,
) {
	err := spi.CheckRsp(pRspInfo)
	if err != nil {
		spi.Error(
			"thost trader [OnRspQuoteInsert] failed",
			slog.Any("error", err),
			slog.Int("request_id", nRequestID),
			slog.Bool("is_last", bIsLast),
		)
		return
	}

	spi.Info(
		"thost trader [OnRspQuoteInsert] succeeded",
		slog.Any("data", pInputQuote),
	)
}

func (spi *ThostFutureBase) OnRspQuoteAction(
	pInputQuoteAction *CThostFtdcInputQuoteActionField,
	pRspInfo *CThostFtdcRspInfoField,
	nRequestID int, bIsLast bool,
) {
	err := spi.CheckRsp(pRspInfo)
	if err != nil {
		spi.Error(
			"thost trader [OnRspQuoteAction] failed",
			slog.Any("error", err),
			slog.Int("request_id", nRequestID),
			slog.Bool("is_last", bIsLast),
		)
		return
	}

	spi.Info(
		"thost trader [OnRspQuoteAction] succeeded",
		slog.Any("data", pInputQuoteAction),
	)
}

func (spi *ThostFutureBase) OnRspBatchOrderAction(
	pInputBatchOrderAction *CThostFtdcInputBatchOrderActionField,
	pRspInfo *CThostFtdcRspInfoField,
	nRequestID int, bIsLast bool,
) {
	err := spi.CheckRsp(pRspInfo)
	if err != nil {
		spi.Error(
			"thost trader [OnRspBatchOrderAction] failed",
			slog.Any("error", err),
			slog.Int("request_id", nRequestID),
			slog.Bool("is_last", bIsLast),
		)
		return
	}

	spi.Info(
		"thost trader [OnRspBatchOrderAction] succeeded",
		slog.Any("data", pInputBatchOrderAction),
	)
}

func (spi *ThostFutureBase) OnRspOptionSelfCloseInsert(
	pInputOptionSelfClose *CThostFtdcInputOptionSelfCloseField,
	pRspInfo *CThostFtdcRspInfoField,
	nRequestID int, bIsLast bool,
) {
	err := spi.CheckRsp(pRspInfo)
	if err != nil {
		spi.Error(
			"thost trader [OnRspOptionSelfCloseInsert] failed",
			slog.Any("error", err),
			slog.Int("request_id", nRequestID),
			slog.Bool("is_last", bIsLast),
		)
		return
	}

	spi.Info(
		"thost trader [OnRspOptionSelfCloseInsert] succeeded",
		slog.Any("data", pInputOptionSelfClose),
	)
}

func (spi *ThostFutureBase) OnRspOptionSelfCloseAction(
	pInputOptionSelfCloseAction *CThostFtdcInputOptionSelfCloseActionField,
	pRspInfo *CThostFtdcRspInfoField,
	nRequestID int, bIsLast bool,
) {
	err := spi.CheckRsp(pRspInfo)
	if err != nil {
		spi.Error(
			"thost trader [OnRspOptionSelfCloseAction] failed",
			slog.Any("error", err),
			slog.Int("request_id", nRequestID),
			slog.Bool("is_last", bIsLast),
		)
		return
	}

	spi.Info(
		"thost trader [OnRspOptionSelfCloseAction] succeeded",
		slog.Any("data", pInputOptionSelfCloseAction),
	)
}

func (spi *ThostFutureBase) OnRspCombActionInsert(
	pInputCombAction *CThostFtdcInputCombActionField,
	pRspInfo *CThostFtdcRspInfoField,
	nRequestID int, bIsLast bool,
) {
	err := spi.CheckRsp(pRspInfo)
	if err != nil {
		spi.Error(
			"thost trader [OnRspCombActionInsert] failed",
			slog.Any("error", err),
			slog.Int("request_id", nRequestID),
			slog.Bool("is_last", bIsLast),
		)
		return
	}

	spi.Info(
		"thost trader [OnRspCombActionInsert] succeeded",
		slog.Any("data", pInputCombAction),
	)
}

func (spi *ThostFutureBase) OnRspQryOrder(
	pOrder *CThostFtdcOrderField,
	pRspInfo *CThostFtdcRspInfoField,
	nRequestID int, bIsLast bool,
) {
	err := spi.CheckRsp(pRspInfo)
	if err != nil {
		spi.Error(
			"thost trader [OnRspQryOrder] failed",
			slog.Any("error", err),
			slog.Int("request_id", nRequestID),
			slog.Bool("is_last", bIsLast),
		)
		return
	}

	spi.Log(
		spi.ctx, slog.LevelDebug-2,
		"thost trader [OnRspQryOrder] succeeded",
		slog.Any("data", pOrder),
	)
}

func (spi *ThostFutureBase) OnRspQryTrade(
	pTrade *CThostFtdcTradeField,
	pRspInfo *CThostFtdcRspInfoField,
	nRequestID int, bIsLast bool,
) {
	err := spi.CheckRsp(pRspInfo)
	if err != nil {
		spi.Error(
			"thost trader [OnRspQryTrade] failed",
			slog.Any("error", err),
			slog.Int("request_id", nRequestID),
			slog.Bool("is_last", bIsLast),
		)
		return
	}

	spi.Log(
		spi.ctx, slog.LevelDebug-2,
		"thost trader [OnRspQryTrade] succeeded",
		slog.Any("data", pTrade),
	)
}

func (spi *ThostFutureBase) OnRspQryInvestorPosition(
	pInvestorPosition *CThostFtdcInvestorPositionField,
	pRspInfo *CThostFtdcRspInfoField,
	nRequestID int, bIsLast bool,
) {
	err := spi.CheckRsp(pRspInfo)
	if err != nil {
		spi.Error(
			"thost trader [OnRspQryInvestorPosition] failed",
			slog.Any("error", err),
			slog.Int("request_id", nRequestID),
			slog.Bool("is_last", bIsLast),
		)
		return
	}

	spi.Log(
		spi.ctx, slog.LevelDebug-2,
		"thost trader [OnRspQryInvestorPosition] succeeded",
		slog.Any("data", pInvestorPosition),
	)
}

func (spi *ThostFutureBase) OnRspQryTradingAccount(
	pTradingAccount *CThostFtdcTradingAccountField,
	pRspInfo *CThostFtdcRspInfoField,
	nRequestID int, bIsLast bool,
) {
	err := spi.CheckRsp(pRspInfo)
	if err != nil {
		spi.Error(
			"thost trader [OnRspQryTradingAccount] failed",
			slog.Any("error", err),
			slog.Int("request_id", nRequestID),
			slog.Bool("is_last", bIsLast),
		)
		return
	}

	spi.Log(
		spi.ctx, slog.LevelDebug-2,
		"thost trader [OnRspQryTradingAccount] succeeded",
		slog.Any("data", pTradingAccount),
	)
}

func (spi *ThostFutureBase) OnRspQryInvestor(
	pInvestor *CThostFtdcInvestorField,
	pRspInfo *CThostFtdcRspInfoField,
	nRequestID int, bIsLast bool,
) {
	err := spi.CheckRsp(pRspInfo)
	defer func() {
		if bIsLast {
			spi.Complete(nRequestID, spi.caches[InvCache], err)
		}
	}()

	if err != nil {
		spi.Error(
			"thost trader [OnRspQryInvestor] failed",
			slog.Any("error", err),
			slog.Int("request_id", nRequestID),
			slog.Bool("is_last", bIsLast),
		)
		return
	}

	if c, exist := spi.caches[InvCache]; exist {
		cache, err := state.CastDataCache[CThostFtdcInvestorField](c)
		if err != nil {
			spi.Error(
				"cast investor cache failed",
				slog.Any("error", err),
			)
		} else {
			cache.AddOrUpdate(pInvestor)
		}
	}

	spi.Log(
		spi.ctx, slog.LevelDebug-2,
		"thost trader [OnRspQryInvestor] succeeded",
		slog.Any("data", pInvestor),
	)
}

func (spi *ThostFutureBase) OnRspQryTradingCode(
	pTradingCode *CThostFtdcTradingCodeField,
	pRspInfo *CThostFtdcRspInfoField,
	nRequestID int, bIsLast bool,
) {
	err := spi.CheckRsp(pRspInfo)
	if err != nil {
		spi.Error(
			"thost trader [OnRspQryTradingCode] failed",
			slog.Any("error", err),
			slog.Int("request_id", nRequestID),
			slog.Bool("is_last", bIsLast),
		)
		return
	}

	spi.Log(
		spi.ctx, slog.LevelDebug-2,
		"thost trader [OnRspQryTradingCode] succeeded",
		slog.Any("data", pTradingCode),
	)
}

func (spi *ThostFutureBase) OnRspQryInstrumentMarginRate(
	pInstrumentMarginRate *CThostFtdcInstrumentMarginRateField,
	pRspInfo *CThostFtdcRspInfoField,
	nRequestID int, bIsLast bool,
) {
	err := spi.CheckRsp(pRspInfo)
	if err != nil {
		spi.Error(
			"thost trader [OnRspQryInstrumentMarginRate] failed",
			slog.Any("error", err),
			slog.Int("request_id", nRequestID),
			slog.Bool("is_last", bIsLast),
		)
		return
	}

	spi.Log(
		spi.ctx, slog.LevelDebug-2,
		"thost trader [OnRspQryInstrumentMarginRate] succeeded",
		slog.Any("data", pInstrumentMarginRate),
	)
}

func (spi *ThostFutureBase) OnRspQryInstrumentCommissionRate(
	pInstrumentCommissionRate *CThostFtdcInstrumentCommissionRateField,
	pRspInfo *CThostFtdcRspInfoField,
	nRequestID int, bIsLast bool,
) {
	err := spi.CheckRsp(pRspInfo)
	if err != nil {
		spi.Error(
			"thost trader [OnRspQryInstrumentCommissionRate] failed",
			slog.Any("error", err),
			slog.Int("request_id", nRequestID),
			slog.Bool("is_last", bIsLast),
		)
		return
	}

	spi.Log(
		spi.ctx, slog.LevelDebug-2,
		"thost trader [OnRspQryInstrumentCommissionRate] succeeded",
		slog.Any("data", pInstrumentCommissionRate),
	)
}

func (spi *ThostFutureBase) OnRspQryUserSession(
	pUserSession *CThostFtdcUserSessionField,
	pRspInfo *CThostFtdcRspInfoField,
	nRequestID int, bIsLast bool,
) {
	err := spi.CheckRsp(pRspInfo)
	if err != nil {
		spi.Error(
			"thost trader [OnRspQryUserSession] failed",
			slog.Any("error", err),
			slog.Int("request_id", nRequestID),
			slog.Bool("is_last", bIsLast),
		)
		return
	}

	spi.Log(
		spi.ctx, slog.LevelDebug-2,
		"thost trader [OnRspQryUserSession] succeeded",
		slog.Any("data", pUserSession),
	)
}

func (spi *ThostFutureBase) OnRspQryExchange(
	pExchange *CThostFtdcExchangeField,
	pRspInfo *CThostFtdcRspInfoField,
	nRequestID int, bIsLast bool,
) {
	err := spi.CheckRsp(pRspInfo)
	if err != nil {
		spi.Error(
			"thost trader [OnRspQryExchange] failed",
			slog.Any("error", err),
			slog.Int("request_id", nRequestID),
			slog.Bool("is_last", bIsLast),
		)
		return
	}

	spi.Log(
		spi.ctx, slog.LevelDebug-2,
		"thost trader [OnRspQryExchange] succeeded",
		slog.Any("data", pExchange),
	)
}

func (spi *ThostFutureBase) OnRspQryProduct(
	pProduct *CThostFtdcProductField,
	pRspInfo *CThostFtdcRspInfoField,
	nRequestID int, bIsLast bool,
) {
	err := spi.CheckRsp(pRspInfo)
	if err != nil {
		spi.Error(
			"thost trader [OnRspQryProduct] failed",
			slog.Any("error", err),
			slog.Int("request_id", nRequestID),
			slog.Bool("is_last", bIsLast),
		)
		return
	}

	spi.Log(
		spi.ctx, slog.LevelDebug-2,
		"thost trader [OnRspQryProduct] succeeded",
		slog.Any("data", pProduct),
	)
}

func (spi *ThostFutureBase) OnRspQryInstrument(
	pInstrument *CThostFtdcInstrumentField,
	pRspInfo *CThostFtdcRspInfoField,
	nRequestID int, bIsLast bool,
) {
	err := spi.CheckRsp(pRspInfo)
	defer func() {
		if bIsLast {
			spi.Complete(nRequestID, spi.caches[InsCache], err)
		}
	}()

	if err != nil {
		spi.Error(
			"thost trader [OnRspQryInstrument] failed",
			slog.Any("error", err),
			slog.Int("request_id", nRequestID),
			slog.Bool("is_last", bIsLast),
		)
		return
	}

	if c, exist := spi.caches[InsCache]; exist {
		cache, err := state.CastDataCache[CThostFtdcInstrumentField](c)
		if err != nil {
			spi.Error(
				"cast instrument cache failed",
				slog.Any("error", err),
			)
		} else {
			cache.AddOrUpdate(pInstrument)
		}
	}

	spi.Log(
		spi.ctx, slog.LevelDebug-2,
		"thost trader [OnRspQryInstrument] succeeded",
		slog.Any("data", pInstrument),
	)
}

func (spi *ThostFutureBase) OnRspQryDepthMarketData(
	pDepthMarketData *CThostFtdcDepthMarketDataField,
	pRspInfo *CThostFtdcRspInfoField,
	nRequestID int, bIsLast bool,
) {
	err := spi.CheckRsp(pRspInfo)
	if err != nil {
		spi.Error(
			"thost trader [OnRspQryDepthMarketData] failed",
			slog.Any("error", err),
			slog.Int("request_id", nRequestID),
			slog.Bool("is_last", bIsLast),
		)
		return
	}

	spi.Log(
		spi.ctx, slog.LevelDebug-2,
		"thost trader [OnRspQryDepthMarketData] succeeded",
		slog.Any("data", pDepthMarketData),
	)
}

func (spi *ThostFutureBase) OnRspQryTraderOffer(
	pTraderOffer *CThostFtdcTraderOfferField,
	pRspInfo *CThostFtdcRspInfoField,
	nRequestID int, bIsLast bool,
) {
	err := spi.CheckRsp(pRspInfo)
	if err != nil {
		spi.Error(
			"thost trader [OnRspQryTraderOffer] failed",
			slog.Any("error", err),
			slog.Int("request_id", nRequestID),
			slog.Bool("is_last", bIsLast),
		)
		return
	}

	spi.Log(
		spi.ctx, slog.LevelDebug-2,
		"thost trader [OnRspQryTraderOffer] succeeded",
		slog.Any("data", pTraderOffer),
	)
}

func (spi *ThostFutureBase) OnRspQrySettlementInfo(
	pSettlementInfo *CThostFtdcSettlementInfoField,
	pRspInfo *CThostFtdcRspInfoField,
	nRequestID int, bIsLast bool,
) {
	err := spi.CheckRsp(pRspInfo)
	if err != nil {
		spi.Error(
			"thost trader [OnRspQrySettlementInfo] failed",
			slog.Any("error", err),
			slog.Int("request_id", nRequestID),
			slog.Bool("is_last", bIsLast),
		)
		return
	}

	spi.Info(
		"thost trader [OnRspQrySettlementInfo] succeeded",
		slog.Any("data", pSettlementInfo),
	)
}

func (spi *ThostFutureBase) OnRspQryTransferBank(
	pTransferBank *CThostFtdcTransferBankField,
	pRspInfo *CThostFtdcRspInfoField,
	nRequestID int, bIsLast bool,
) {
	err := spi.CheckRsp(pRspInfo)
	if err != nil {
		spi.Error(
			"thost trader [OnRspQryTransferBank] failed",
			slog.Any("error", err),
			slog.Int("request_id", nRequestID),
			slog.Bool("is_last", bIsLast),
		)
		return
	}

	spi.Log(
		spi.ctx, slog.LevelDebug-2,
		"thost trader [OnRspQryTransferBank] succeeded",
		slog.Any("data", pTransferBank),
	)
}

func (spi *ThostFutureBase) OnRspQryInvestorPositionDetail(
	pInvestorPositionDetail *CThostFtdcInvestorPositionDetailField,
	pRspInfo *CThostFtdcRspInfoField,
	nRequestID int, bIsLast bool,
) {
	err := spi.CheckRsp(pRspInfo)
	if err != nil {
		spi.Error(
			"thost trader [OnRspQryInvestorPositionDetail] failed",
			slog.Any("error", err),
			slog.Int("request_id", nRequestID),
			slog.Bool("is_last", bIsLast),
		)
		return
	}

	spi.Log(
		spi.ctx, slog.LevelDebug-2,
		"thost trader [OnRspQryInvestorPositionDetail] succeeded",
		slog.Any("data", pInvestorPositionDetail),
	)
}

func (spi *ThostFutureBase) OnRspQryNotice(
	pNotice *CThostFtdcNoticeField,
	pRspInfo *CThostFtdcRspInfoField,
	nRequestID int, bIsLast bool,
) {
	err := spi.CheckRsp(pRspInfo)
	if err != nil {
		spi.Error(
			"thost trader [OnRspQryNotice] failed",
			slog.Any("error", err),
			slog.Int("request_id", nRequestID),
			slog.Bool("is_last", bIsLast),
		)
		return
	}

	spi.Log(
		spi.ctx, slog.LevelDebug-2,
		"thost trader [OnRspQryNotice] succeeded",
		slog.Any("data", pNotice),
	)
}

func (spi *ThostFutureBase) OnRspQrySettlementInfoConfirm(
	pSettlementInfoConfirm *CThostFtdcSettlementInfoConfirmField,
	pRspInfo *CThostFtdcRspInfoField,
	nRequestID int, bIsLast bool,
) {
	err := spi.CheckRsp(pRspInfo)
	if err != nil {
		spi.Error(
			"thost trader [OnRspQrySettlementInfoConfirm] failed",
			slog.Any("error", err),
			slog.Int("request_id", nRequestID),
			slog.Bool("is_last", bIsLast),
		)
		return
	}

	spi.Info(
		"thost trader [OnRspQrySettlementInfoConfirm] succeeded",
		slog.Any("data", pSettlementInfoConfirm),
	)
}

func (spi *ThostFutureBase) OnRspQryInvestorPositionCombineDetail(
	pInvestorPositionCombineDetail *CThostFtdcInvestorPositionCombineDetailField,
	pRspInfo *CThostFtdcRspInfoField,
	nRequestID int, bIsLast bool,
) {
	err := spi.CheckRsp(pRspInfo)
	if err != nil {
		spi.Error(
			"thost trader [OnRspQryInvestorPositionCombineDetail] failed",
			slog.Any("error", err),
			slog.Int("request_id", nRequestID),
			slog.Bool("is_last", bIsLast),
		)
		return
	}

	spi.Log(
		spi.ctx, slog.LevelDebug-2,
		"thost trader [OnRspQryInvestorPositionCombineDetail] succeeded",
		slog.Any("data", pInvestorPositionCombineDetail),
	)
}

func (spi *ThostFutureBase) OnRspQryCFMMCTradingAccountKey(
	pCFMMCTradingAccountKey *CThostFtdcCFMMCTradingAccountKeyField,
	pRspInfo *CThostFtdcRspInfoField,
	nRequestID int, bIsLast bool,
) {
	err := spi.CheckRsp(pRspInfo)
	if err != nil {
		spi.Error(
			"thost trader [OnRspQryCFMMCTradingAccountKey] failed",
			slog.Any("error", err),
			slog.Int("request_id", nRequestID),
			slog.Bool("is_last", bIsLast),
		)
		return
	}

	spi.Log(
		spi.ctx, slog.LevelDebug-2,
		"thost trader [OnRspQryCFMMCTradingAccountKey] succeeded",
		slog.Any("data", pCFMMCTradingAccountKey),
	)
}

func (spi *ThostFutureBase) OnRspQryEWarrantOffset(
	pEWarrantOffset *CThostFtdcEWarrantOffsetField,
	pRspInfo *CThostFtdcRspInfoField,
	nRequestID int, bIsLast bool,
) {
	err := spi.CheckRsp(pRspInfo)
	if err != nil {
		spi.Error(
			"thost trader [OnRspQryEWarrantOffset] failed",
			slog.Any("error", err),
			slog.Int("request_id", nRequestID),
			slog.Bool("is_last", bIsLast),
		)
		return
	}

	spi.Log(
		spi.ctx, slog.LevelDebug-2,
		"thost trader [OnRspQryEWarrantOffset] succeeded",
		slog.Any("data", pEWarrantOffset),
	)
}

func (spi *ThostFutureBase) OnRspQryInvestorProductGroupMargin(
	pInvestorProductGroupMargin *CThostFtdcInvestorProductGroupMarginField,
	pRspInfo *CThostFtdcRspInfoField,
	nRequestID int, bIsLast bool,
) {
	err := spi.CheckRsp(pRspInfo)
	if err != nil {
		spi.Error(
			"thost trader [OnRspQryInvestorProductGroupMargin] failed",
			slog.Any("error", err),
			slog.Int("request_id", nRequestID),
			slog.Bool("is_last", bIsLast),
		)
		return
	}

	spi.Log(
		spi.ctx, slog.LevelDebug-2,
		"thost trader [OnRspQryInvestorProductGroupMargin] succeeded",
		slog.Any("data", pInvestorProductGroupMargin),
	)
}

func (spi *ThostFutureBase) OnRspQryExchangeMarginRate(
	pExchangeMarginRate *CThostFtdcExchangeMarginRateField,
	pRspInfo *CThostFtdcRspInfoField,
	nRequestID int, bIsLast bool,
) {
	err := spi.CheckRsp(pRspInfo)
	if err != nil {
		spi.Error(
			"thost trader [OnRspQryExchangeMarginRate] failed",
			slog.Any("error", err),
			slog.Int("request_id", nRequestID),
			slog.Bool("is_last", bIsLast),
		)
		return
	}

	spi.Log(
		spi.ctx, slog.LevelDebug-2,
		"thost trader [OnRspQryExchangeMarginRate] succeeded",
		slog.Any("data", pExchangeMarginRate),
	)
}

func (spi *ThostFutureBase) OnRspQryExchangeMarginRateAdjust(
	pExchangeMarginRateAdjust *CThostFtdcExchangeMarginRateAdjustField,
	pRspInfo *CThostFtdcRspInfoField,
	nRequestID int, bIsLast bool,
) {
	err := spi.CheckRsp(pRspInfo)
	if err != nil {
		spi.Error(
			"thost trader [OnRspQryExchangeMarginRateAdjust] failed",
			slog.Any("error", err),
			slog.Int("request_id", nRequestID),
			slog.Bool("is_last", bIsLast),
		)
		return
	}

	spi.Log(
		spi.ctx, slog.LevelDebug-2,
		"thost trader [OnRspQryExchangeMarginRateAdjust] succeeded",
		slog.Any("data", pExchangeMarginRateAdjust),
	)
}

func (spi *ThostFutureBase) OnRspQryExchangeRate(
	pExchangeRate *CThostFtdcExchangeRateField,
	pRspInfo *CThostFtdcRspInfoField,
	nRequestID int, bIsLast bool,
) {
	err := spi.CheckRsp(pRspInfo)
	if err != nil {
		spi.Error(
			"thost trader [OnRspQryExchangeRate] failed",
			slog.Any("error", err),
			slog.Int("request_id", nRequestID),
			slog.Bool("is_last", bIsLast),
		)
		return
	}

	spi.Log(
		spi.ctx, slog.LevelDebug-2,
		"thost trader [OnRspQryExchangeRate] succeeded",
		slog.Any("data", pExchangeRate),
	)
}

func (spi *ThostFutureBase) OnRspQrySecAgentACIDMap(
	pSecAgentACIDMap *CThostFtdcSecAgentACIDMapField,
	pRspInfo *CThostFtdcRspInfoField,
	nRequestID int, bIsLast bool,
) {
	err := spi.CheckRsp(pRspInfo)
	if err != nil {
		spi.Error(
			"thost trader [OnRspQrySecAgentACIDMap] failed",
			slog.Any("error", err),
			slog.Int("request_id", nRequestID),
			slog.Bool("is_last", bIsLast),
		)
		return
	}

	spi.Log(
		spi.ctx, slog.LevelDebug-2,
		"thost trader [OnRspQrySecAgentACIDMap] succeeded",
		slog.Any("data", pSecAgentACIDMap),
	)
}

func (spi *ThostFutureBase) OnRspQryProductExchRate(
	pProductExchRate *CThostFtdcProductExchRateField,
	pRspInfo *CThostFtdcRspInfoField,
	nRequestID int, bIsLast bool,
) {
	err := spi.CheckRsp(pRspInfo)
	if err != nil {
		spi.Error(
			"thost trader [OnRspQryProductExchRate] failed",
			slog.Any("error", err),
			slog.Int("request_id", nRequestID),
			slog.Bool("is_last", bIsLast),
		)
		return
	}

	spi.Log(
		spi.ctx, slog.LevelDebug-2,
		"thost trader [OnRspQryProductExchRate] succeeded",
		slog.Any("data", pProductExchRate),
	)
}

func (spi *ThostFutureBase) OnRspQryProductGroup(
	pProductGroup *CThostFtdcProductGroupField,
	pRspInfo *CThostFtdcRspInfoField,
	nRequestID int, bIsLast bool,
) {
	err := spi.CheckRsp(pRspInfo)
	if err != nil {
		spi.Error(
			"thost trader [OnRspQryProductGroup] failed",
			slog.Any("error", err),
			slog.Int("request_id", nRequestID),
			slog.Bool("is_last", bIsLast),
		)
		return
	}

	spi.Info(
		"thost trader [OnRspQryProductGroup] succeeded",
		slog.Any("data", pProductGroup),
	)
}

func (spi *ThostFutureBase) OnRspQryMMInstrumentCommissionRate(
	pMMInstrumentCommissionRate *CThostFtdcMMInstrumentCommissionRateField,
	pRspInfo *CThostFtdcRspInfoField,
	nRequestID int, bIsLast bool,
) {
	err := spi.CheckRsp(pRspInfo)
	if err != nil {
		spi.Error(
			"thost trader [OnRspQryMMInstrumentCommissionRate] failed",
			slog.Any("error", err),
			slog.Int("request_id", nRequestID),
			slog.Bool("is_last", bIsLast),
		)
		return
	}

	spi.Info(
		"thost trader [OnRspQryMMInstrumentCommissionRate] succeeded",
		slog.Any("data", pMMInstrumentCommissionRate),
	)
}

func (spi *ThostFutureBase) OnRspQryMMOptionInstrCommRate(
	pMMOptionInstrCommRate *CThostFtdcMMOptionInstrCommRateField,
	pRspInfo *CThostFtdcRspInfoField,
	nRequestID int, bIsLast bool,
) {
	err := spi.CheckRsp(pRspInfo)
	if err != nil {
		spi.Error(
			"thost trader [OnRspQryMMOptionInstrCommRate] failed",
			slog.Any("error", err),
			slog.Int("request_id", nRequestID),
			slog.Bool("is_last", bIsLast),
		)
		return
	}

	spi.Info(
		"thost trader [OnRspQryMMOptionInstrCommRate] succeeded",
		slog.Any("data", pMMOptionInstrCommRate),
	)
}

func (spi *ThostFutureBase) OnRspQryInstrumentOrderCommRate(
	pInstrumentOrderCommRate *CThostFtdcInstrumentOrderCommRateField,
	pRspInfo *CThostFtdcRspInfoField,
	nRequestID int, bIsLast bool,
) {
	err := spi.CheckRsp(pRspInfo)
	if err != nil {
		spi.Error(
			"thost trader [OnRspQryInstrumentOrderCommRate] failed",
			slog.Any("error", err),
			slog.Int("request_id", nRequestID),
			slog.Bool("is_last", bIsLast),
		)
		return
	}

	spi.Info(
		"thost trader [OnRspQryInstrumentOrderCommRate] succeeded",
		slog.Any("data", pInstrumentOrderCommRate),
	)
}

func (spi *ThostFutureBase) OnRspQrySecAgentTradingAccount(
	pTradingAccount *CThostFtdcTradingAccountField,
	pRspInfo *CThostFtdcRspInfoField,
	nRequestID int, bIsLast bool,
) {
	err := spi.CheckRsp(pRspInfo)
	if err != nil {
		spi.Error(
			"thost trader [OnRspQrySecAgentTradingAccount] failed",
			slog.Any("error", err),
			slog.Int("request_id", nRequestID),
			slog.Bool("is_last", bIsLast),
		)
		return
	}

	spi.Info(
		"thost trader [OnRspQrySecAgentTradingAccount] succeeded",
		slog.Any("data", pTradingAccount),
	)
}

func (spi *ThostFutureBase) OnRspQrySecAgentCheckMode(
	pSecAgentCheckMode *CThostFtdcSecAgentCheckModeField,
	pRspInfo *CThostFtdcRspInfoField,
	nRequestID int, bIsLast bool,
) {
	err := spi.CheckRsp(pRspInfo)
	if err != nil {
		spi.Error(
			"thost trader [OnRspQrySecAgentCheckMode] failed",
			slog.Any("error", err),
			slog.Int("request_id", nRequestID),
			slog.Bool("is_last", bIsLast),
		)
		return
	}

	spi.Info(
		"thost trader [OnRspQrySecAgentCheckMode] succeeded",
		slog.Any("data", pSecAgentCheckMode),
	)
}

func (spi *ThostFutureBase) OnRspQrySecAgentTradeInfo(
	pSecAgentTradeInfo *CThostFtdcSecAgentTradeInfoField,
	pRspInfo *CThostFtdcRspInfoField,
	nRequestID int, bIsLast bool,
) {
	err := spi.CheckRsp(pRspInfo)
	if err != nil {
		spi.Error(
			"thost trader [OnRspQrySecAgentTradeInfo] failed",
			slog.Any("error", err),
			slog.Int("request_id", nRequestID),
			slog.Bool("is_last", bIsLast),
		)
		return
	}

	spi.Info(
		"thost trader [OnRspQrySecAgentTradeInfo] succeeded",
		slog.Any("data", pSecAgentTradeInfo),
	)
}

func (spi *ThostFutureBase) OnRspQryOptionInstrTradeCost(
	pOptionInstrTradeCost *CThostFtdcOptionInstrTradeCostField,
	pRspInfo *CThostFtdcRspInfoField,
	nRequestID int, bIsLast bool,
) {
	err := spi.CheckRsp(pRspInfo)
	if err != nil {
		spi.Error(
			"thost trader [OnRspQryOptionInstrTradeCost] failed",
			slog.Any("error", err),
			slog.Int("request_id", nRequestID),
			slog.Bool("is_last", bIsLast),
		)
		return
	}

	spi.Info(
		"thost trader [OnRspQryOptionInstrTradeCost] succeeded",
		slog.Any("data", pOptionInstrTradeCost),
	)
}

func (spi *ThostFutureBase) OnRspQryOptionInstrCommRate(
	pOptionInstrCommRate *CThostFtdcOptionInstrCommRateField,
	pRspInfo *CThostFtdcRspInfoField,
	nRequestID int, bIsLast bool,
) {
	err := spi.CheckRsp(pRspInfo)
	if err != nil {
		spi.Error(
			"thost trader [OnRspQryOptionInstrCommRate] failed",
			slog.Any("error", err),
			slog.Int("request_id", nRequestID),
			slog.Bool("is_last", bIsLast),
		)
		return
	}

	spi.Info(
		"thost trader [OnRspQryOptionInstrCommRate] succeeded",
		slog.Any("data", pOptionInstrCommRate),
	)
}

func (spi *ThostFutureBase) OnRspQryExecOrder(
	pExecOrder *CThostFtdcExecOrderField,
	pRspInfo *CThostFtdcRspInfoField,
	nRequestID int, bIsLast bool,
) {
	err := spi.CheckRsp(pRspInfo)
	if err != nil {
		spi.Error(
			"thost trader [OnRspQryExecOrder] failed",
			slog.Any("error", err),
			slog.Int("request_id", nRequestID),
			slog.Bool("is_last", bIsLast),
		)
		return
	}

	spi.Info(
		"thost trader [OnRspQryExecOrder] succeeded",
		slog.Any("data", pExecOrder),
	)
}

func (spi *ThostFutureBase) OnRspQryForQuote(
	pForQuote *CThostFtdcForQuoteField,
	pRspInfo *CThostFtdcRspInfoField,
	nRequestID int, bIsLast bool,
) {
	err := spi.CheckRsp(pRspInfo)
	if err != nil {
		spi.Error(
			"thost trader [OnRspQryForQuote] failed",
			slog.Any("error", err),
			slog.Int("request_id", nRequestID),
			slog.Bool("is_last", bIsLast),
		)
		return
	}

	spi.Info(
		"thost trader [OnRspQryForQuote] succeeded",
		slog.Any("data", pForQuote),
	)
}

func (spi *ThostFutureBase) OnRspQryQuote(
	pQuote *CThostFtdcQuoteField,
	pRspInfo *CThostFtdcRspInfoField,
	nRequestID int, bIsLast bool,
) {
	err := spi.CheckRsp(pRspInfo)
	if err != nil {
		spi.Error(
			"thost trader [OnRspQryQuote] failed",
			slog.Any("error", err),
			slog.Int("request_id", nRequestID),
			slog.Bool("is_last", bIsLast),
		)
		return
	}

	spi.Info(
		"thost trader [OnRspQryQuote] succeeded",
		slog.Any("data", pQuote),
	)
}

func (spi *ThostFutureBase) OnRspQryOptionSelfClose(
	pOptionSelfClose *CThostFtdcOptionSelfCloseField,
	pRspInfo *CThostFtdcRspInfoField,
	nRequestID int, bIsLast bool,
) {
	err := spi.CheckRsp(pRspInfo)
	if err != nil {
		spi.Error(
			"thost trader [OnRspQryOptionSelfClose] failed",
			slog.Any("error", err),
			slog.Int("request_id", nRequestID),
			slog.Bool("is_last", bIsLast),
		)
		return
	}

	spi.Info(
		"thost trader [OnRspQryOptionSelfClose] succeeded",
		slog.Any("data", pOptionSelfClose),
	)
}

func (spi *ThostFutureBase) OnRspQryInvestUnit(
	pInvestUnit *CThostFtdcInvestUnitField,
	pRspInfo *CThostFtdcRspInfoField,
	nRequestID int, bIsLast bool,
) {
	err := spi.CheckRsp(pRspInfo)
	if err != nil {
		spi.Error(
			"thost trader [OnRspQryInvestUnit] failed",
			slog.Any("error", err),
			slog.Int("request_id", nRequestID),
			slog.Bool("is_last", bIsLast),
		)
		return
	}

	spi.Info(
		"thost trader [OnRspQryInvestUnit] succeeded",
		slog.Any("data", pInvestUnit),
	)
}

func (spi *ThostFutureBase) OnRspQryCombInstrumentGuard(
	pCombInstrumentGuard *CThostFtdcCombInstrumentGuardField,
	pRspInfo *CThostFtdcRspInfoField,
	nRequestID int, bIsLast bool,
) {
	err := spi.CheckRsp(pRspInfo)
	if err != nil {
		spi.Error(
			"thost trader [OnRspQryCombInstrumentGuard] failed",
			slog.Any("error", err),
			slog.Int("request_id", nRequestID),
			slog.Bool("is_last", bIsLast),
		)
		return
	}

	spi.Info(
		"thost trader [OnRspQryCombInstrumentGuard] succeeded",
		slog.Any("data", pCombInstrumentGuard),
	)
}

func (spi *ThostFutureBase) OnRspQryCombAction(
	pCombAction *CThostFtdcCombActionField,
	pRspInfo *CThostFtdcRspInfoField,
	nRequestID int, bIsLast bool,
) {
	err := spi.CheckRsp(pRspInfo)
	if err != nil {
		spi.Error(
			"thost trader [OnRspQryCombAction] failed",
			slog.Any("error", err),
			slog.Int("request_id", nRequestID),
			slog.Bool("is_last", bIsLast),
		)
		return
	}

	spi.Info(
		"thost trader [OnRspQryCombAction] succeeded",
		slog.Any("data", pCombAction),
	)
}

func (spi *ThostFutureBase) OnRspQryTransferSerial(
	pTransferSerial *CThostFtdcTransferSerialField,
	pRspInfo *CThostFtdcRspInfoField,
	nRequestID int, bIsLast bool,
) {
	err := spi.CheckRsp(pRspInfo)
	if err != nil {
		spi.Error(
			"thost trader [OnRspQryTransferSerial] failed",
			slog.Any("error", err),
			slog.Int("request_id", nRequestID),
			slog.Bool("is_last", bIsLast),
		)
		return
	}

	spi.Info(
		"thost trader [OnRspQryTransferSerial] succeeded",
		slog.Any("data", pTransferSerial),
	)
}

func (spi *ThostFutureBase) OnRspQryAccountregister(
	pAccountregister *CThostFtdcAccountregisterField,
	pRspInfo *CThostFtdcRspInfoField,
	nRequestID int, bIsLast bool,
) {
	err := spi.CheckRsp(pRspInfo)
	if err != nil {
		spi.Error(
			"thost trader [OnRspQryAccountregister] failed",
			slog.Any("error", err),
			slog.Int("request_id", nRequestID),
			slog.Bool("is_last", bIsLast),
		)
		return
	}

	spi.Info(
		"thost trader [OnRspQryAccountregister] succeeded",
		slog.Any("data", pAccountregister),
	)
}

func (spi *ThostFutureBase) OnRspError(pRspInfo *CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
	spi.Error(
		"thost trader [OnRspError]",
		slog.Any("error", spi.CheckRsp(pRspInfo)),
		slog.Int("request_id", nRequestID),
		slog.Bool("is_last", bIsLast),
	)
}

func (spi *ThostFutureBase) OnRtnOrder(pOrder *CThostFtdcOrderField) {
	spi.Log(
		spi.ctx, slog.LevelDebug-1,
		"thost trader [OnRtnOrder]",
		slog.Any("data", pOrder),
	)
}

func (spi *ThostFutureBase) OnRtnTrade(pTrade *CThostFtdcTradeField) {
	spi.Log(
		spi.ctx, slog.LevelDebug-1,
		"thost trader [OnRtnTrade]",
		slog.Any("data", pTrade),
	)
}

func (spi *ThostFutureBase) OnErrRtnOrderInsert(
	pInputOrder *CThostFtdcInputOrderField, pRspInfo *CThostFtdcRspInfoField,
) {
	spi.Error(
		"thost trader [OnRtnOrderInsert] error",
		slog.Any("error", spi.CheckRsp(pRspInfo)),
		slog.Any("data", pInputOrder),
	)
}

func (spi *ThostFutureBase) OnErrRtnOrderAction(
	pOrderAction *CThostFtdcOrderActionField, pRspInfo *CThostFtdcRspInfoField,
) {
	spi.Error(
		"thost trader [OnRtnOrderAction] error",
		slog.Any("error", spi.CheckRsp(pRspInfo)),
		slog.Any("data", pOrderAction),
	)
}

func (spi *ThostFutureBase) OnRtnInstrumentStatus(pInstrumentStatus *CThostFtdcInstrumentStatusField) {
	spi.Log(
		spi.ctx, slog.LevelDebug-1,
		"thost trader [OnRtnInstrumentStatus]",
		slog.Any("data", pInstrumentStatus),
	)
}

func (spi *ThostFutureBase) OnRtnBulletin(pBulletin *CThostFtdcBulletinField) {
	spi.Log(
		spi.ctx, slog.LevelDebug-1,
		"thost trader [OnRtnBulletin]",
		slog.Any("data", pBulletin),
	)
}

func (spi *ThostFutureBase) OnRtnTradingNotice(pTradingNoticeInfo *CThostFtdcTradingNoticeInfoField) {
	spi.Log(
		spi.ctx, slog.LevelDebug-1,
		"thost trader [OnRtnTradingNotice]",
		slog.Any("data", pTradingNoticeInfo),
	)
}

func (spi *ThostFutureBase) OnRtnErrorConditionalOrder(pErrorConditionalOrder *CThostFtdcErrorConditionalOrderField) {
	spi.Log(
		spi.ctx, slog.LevelDebug-1,
		"thost trader [OnRtnErrorConditionalOrder]",
		slog.Any("data", pErrorConditionalOrder),
	)
}

func (spi *ThostFutureBase) OnRtnExecOrder(pExecOrder *CThostFtdcExecOrderField) {
	spi.Log(
		spi.ctx, slog.LevelDebug-1,
		"thost trader [OnRtnExecOrder]",
		slog.Any("data", pExecOrder),
	)
}

func (spi *ThostFutureBase) OnErrRtnExecOrderInsert(
	pInputExecOrder *CThostFtdcInputExecOrderField, pRspInfo *CThostFtdcRspInfoField,
) {
	spi.Error(
		"thost trader [OnRtnExecOrderInsert] error",
		slog.Any("error", spi.CheckRsp(pRspInfo)),
		slog.Any("data", pInputExecOrder),
	)
}

func (spi *ThostFutureBase) OnErrRtnExecOrderAction(
	pExecOrderAction *CThostFtdcExecOrderActionField, pRspInfo *CThostFtdcRspInfoField,
) {
	spi.Error(
		"thost trader [OnRtnExecOrderAction] error",
		slog.Any("error", spi.CheckRsp(pRspInfo)),
		slog.Any("data", pExecOrderAction),
	)
}

func (spi *ThostFutureBase) OnErrRtnForQuoteInsert(
	pInputForQuote *CThostFtdcInputForQuoteField, pRspInfo *CThostFtdcRspInfoField,
) {
	spi.Error(
		"thost trader [OnRtnForQuoteInsert] error",
		slog.Any("error", spi.CheckRsp(pRspInfo)),
		slog.Any("data", pInputForQuote),
	)
}

func (spi *ThostFutureBase) OnRtnQuote(pQuote *CThostFtdcQuoteField) {
	spi.Log(
		spi.ctx, slog.LevelDebug-1,
		"thost trader [OnRtnQuote]",
		slog.Any("data", pQuote),
	)
}

func (spi *ThostFutureBase) OnErrRtnQuoteInsert(
	pInputQuote *CThostFtdcInputQuoteField, pRspInfo *CThostFtdcRspInfoField,
) {
	spi.Error(
		"thost trader [OnRtnQuoteInsert] error",
		slog.Any("error", spi.CheckRsp(pRspInfo)),
		slog.Any("data", pInputQuote),
	)
}

func (spi *ThostFutureBase) OnErrRtnQuoteAction(
	pQuoteAction *CThostFtdcQuoteActionField, pRspInfo *CThostFtdcRspInfoField,
) {
	spi.Error(
		"thost trader [OnRtnQuoteAction] error",
		slog.Any("error", spi.CheckRsp(pRspInfo)),
		slog.Any("data", pQuoteAction),
	)
}

func (spi *ThostFutureBase) OnRtnForQuoteRsp(pForQuoteRsp *CThostFtdcForQuoteRspField) {
	spi.Log(
		spi.ctx, slog.LevelDebug-1,
		"thost [OnRtnForQuoteRsp]",
		slog.Any("data", pForQuoteRsp),
	)
}

func (spi *ThostFutureBase) OnRtnCFMMCTradingAccountToken(pCFMMCTradingAccountToken *CThostFtdcCFMMCTradingAccountTokenField) {
	spi.Log(
		spi.ctx, slog.LevelDebug-1,
		"thost trader [OnRtnCFMMCTradingAccountToken]",
		slog.Any("data", pCFMMCTradingAccountToken),
	)
}

func (spi *ThostFutureBase) OnErrRtnBatchOrderAction(
	pBatchOrderAction *CThostFtdcBatchOrderActionField, pRspInfo *CThostFtdcRspInfoField,
) {
	spi.Error(
		"thost trader [OnRtnBatchOrderAction] error",
		slog.Any("error", spi.CheckRsp(pRspInfo)),
		slog.Any("data", pBatchOrderAction),
	)
}

func (spi *ThostFutureBase) OnRtnOptionSelfClose(pOptionSelfClose *CThostFtdcOptionSelfCloseField) {
	spi.Log(
		spi.ctx, slog.LevelDebug-1,
		"thost trader [OnRtnOptionSelfClose]",
		slog.Any("data", pOptionSelfClose),
	)
}

func (spi *ThostFutureBase) OnErrRtnOptionSelfCloseInsert(
	pInputOptionSelfClose *CThostFtdcInputOptionSelfCloseField, pRspInfo *CThostFtdcRspInfoField,
) {
	spi.Error(
		"thost trader [OnRtnOptionSelfCloseInsert] error",
		slog.Any("error", spi.CheckRsp(pRspInfo)),
		slog.Any("data", pInputOptionSelfClose),
	)
}

func (spi *ThostFutureBase) OnErrRtnOptionSelfCloseAction(
	pOptionSelfCloseAction *CThostFtdcOptionSelfCloseActionField, pRspInfo *CThostFtdcRspInfoField,
) {
	spi.Error(
		"thost trader [OnRtnOptionSelfCloseAction] error",
		slog.Any("error", spi.CheckRsp(pRspInfo)),
		slog.Any("data", pOptionSelfCloseAction),
	)
}

func (spi *ThostFutureBase) OnRtnCombAction(pCombAction *CThostFtdcCombActionField) {
	spi.Log(
		spi.ctx, slog.LevelDebug-1,
		"thost trader [OnRtnCombAction]",
		slog.Any("data", pCombAction),
	)
}

func (spi *ThostFutureBase) OnErrRtnCombActionInsert(
	pInputCombAction *CThostFtdcInputCombActionField, pRspInfo *CThostFtdcRspInfoField,
) {
	spi.Error(
		"thost trader [OnRtnCombActionInsert] error",
		slog.Any("error", spi.CheckRsp(pRspInfo)),
		slog.Any("data", pInputCombAction),
	)
}

func (spi *ThostFutureBase) OnRspQryContractBank(
	pContractBank *CThostFtdcContractBankField,
	pRspInfo *CThostFtdcRspInfoField,
	nRequestID int, bIsLast bool,
) {
	err := spi.CheckRsp(pRspInfo)
	if err != nil {
		spi.Error(
			"thost trader [OnRspQryContractBank] failed",
			slog.Any("error", err),
			slog.Int("request_id", nRequestID),
			slog.Bool("is_last", bIsLast),
		)
		return
	}

	spi.Info(
		"thost trader [OnRspQryContractBank] succeeded",
		slog.Any("data", pContractBank),
	)
}

func (spi *ThostFutureBase) OnRspQryParkedOrder(
	pParkedOrder *CThostFtdcParkedOrderField,
	pRspInfo *CThostFtdcRspInfoField,
	nRequestID int, bIsLast bool,
) {
	err := spi.CheckRsp(pRspInfo)
	if err != nil {
		spi.Error(
			"thost trader [OnRspQryParkedOrder] failed",
			slog.Any("error", err),
			slog.Int("request_id", nRequestID),
			slog.Bool("is_last", bIsLast),
		)
		return
	}

	spi.Info(
		"thost trader [OnRspQryParkedOrder] succeeded",
		slog.Any("data", pParkedOrder),
	)
}

func (spi *ThostFutureBase) OnRspQryParkedOrderAction(
	pParkedOrderAction *CThostFtdcParkedOrderActionField,
	pRspInfo *CThostFtdcRspInfoField,
	nRequestID int, bIsLast bool,
) {
	err := spi.CheckRsp(pRspInfo)
	if err != nil {
		spi.Error(
			"thost trader [OnRspQryParkedOrderAction] failed",
			slog.Any("error", err),
			slog.Int("request_id", nRequestID),
			slog.Bool("is_last", bIsLast),
		)
		return
	}

	spi.Info(
		"thost trader [OnRspQryParkedOrderAction] succeeded",
		slog.Any("data", pParkedOrderAction),
	)
}

func (spi *ThostFutureBase) OnRspQryTradingNotice(
	pTradingNotice *CThostFtdcTradingNoticeField,
	pRspInfo *CThostFtdcRspInfoField,
	nRequestID int, bIsLast bool,
) {
	err := spi.CheckRsp(pRspInfo)
	if err != nil {
		spi.Error(
			"thost trader [OnRspQryTradingNotice] failed",
			slog.Any("error", err),
			slog.Int("request_id", nRequestID),
			slog.Bool("is_last", bIsLast),
		)
		return
	}

	spi.Info(
		"thost trader [OnRspQryTradingNotice] succeeded",
		slog.Any("data", pTradingNotice),
	)
}

func (spi *ThostFutureBase) OnRspQryBrokerTradingParams(
	pBrokerTradingParams *CThostFtdcBrokerTradingParamsField,
	pRspInfo *CThostFtdcRspInfoField,
	nRequestID int, bIsLast bool,
) {
	err := spi.CheckRsp(pRspInfo)
	if err != nil {
		spi.Error(
			"thost trader [OnRspQryBrokerTradingParams] failed",
			slog.Any("error", err),
			slog.Int("request_id", nRequestID),
			slog.Bool("is_last", bIsLast),
		)
		return
	}

	spi.Info(
		"thost trader [OnRspQryBrokerTradingParams] succeeded",
		slog.Any("data", pBrokerTradingParams),
	)
}

func (spi *ThostFutureBase) OnRspQryBrokerTradingAlgos(
	pBrokerTradingAlgos *CThostFtdcBrokerTradingAlgosField,
	pRspInfo *CThostFtdcRspInfoField,
	nRequestID int, bIsLast bool,
) {
	err := spi.CheckRsp(pRspInfo)
	if err != nil {
		spi.Error(
			"thost trader [OnRspQryBrokerTradingAlgos] failed",
			slog.Any("error", err),
			slog.Int("request_id", nRequestID),
			slog.Bool("is_last", bIsLast),
		)
		return
	}

	spi.Info(
		"thost trader [OnRspQryBrokerTradingAlgos] succeeded",
		slog.Any("data", pBrokerTradingAlgos),
	)
}

func (spi *ThostFutureBase) OnRspQueryCFMMCTradingAccountToken(
	pQueryCFMMCTradingAccountToken *CThostFtdcQueryCFMMCTradingAccountTokenField,
	pRspInfo *CThostFtdcRspInfoField,
	nRequestID int, bIsLast bool,
) {
	err := spi.CheckRsp(pRspInfo)
	if err != nil {
		spi.Error(
			"thost trader [OnRspQueryCFMMCTradingAccountToken] failed",
			slog.Any("error", err),
			slog.Int("request_id", nRequestID),
			slog.Bool("is_last", bIsLast),
		)
		return
	}

	spi.Info(
		"thost trader [OnRspQueryCFMMCTradingAccountToken] succeeded",
		slog.Any("data", pQueryCFMMCTradingAccountToken),
	)
}

func (spi *ThostFutureBase) OnRtnFromBankToFutureByBank(pRspTransfer *CThostFtdcRspTransferField) {
	spi.Log(
		spi.ctx, slog.LevelDebug-1,
		"thost trader [OnRtnFromBankToFutureByBank]",
		slog.Any("data", pRspTransfer),
	)
}

func (spi *ThostFutureBase) OnRtnFromFutureToBankByBank(pRspTransfer *CThostFtdcRspTransferField) {
	spi.Log(
		spi.ctx, slog.LevelDebug-1,
		"thost trader [OnRtnFromFutureToBankByBank]",
		slog.Any("data", pRspTransfer),
	)
}

func (spi *ThostFutureBase) OnRtnRepealFromBankToFutureByBank(pRspRepeal *CThostFtdcRspRepealField) {
	spi.Log(
		spi.ctx, slog.LevelDebug-1,
		"thost trader [OnRtnRepealFromBankToFutureByBank]",
		slog.Any("data", pRspRepeal),
	)
}

func (spi *ThostFutureBase) OnRtnRepealFromFutureToBankByBank(pRspRepeal *CThostFtdcRspRepealField) {
	spi.Log(
		spi.ctx, slog.LevelDebug-1,
		"thost trader [OnRtnRepealFromFutureToBankByBank]",
		slog.Any("data", pRspRepeal),
	)
}

func (spi *ThostFutureBase) OnRtnFromBankToFutureByFuture(pRspTransfer *CThostFtdcRspTransferField) {
	spi.Log(
		spi.ctx, slog.LevelDebug-1,
		"thost trader [OnRtnFromBankToFutureByFuture]",
		slog.Any("data", pRspTransfer),
	)
}

func (spi *ThostFutureBase) OnRtnFromFutureToBankByFuture(pRspTransfer *CThostFtdcRspTransferField) {
	spi.Log(
		spi.ctx, slog.LevelDebug-1,
		"thost trader [OnRtnFromFutureToBankByFuture]",
		slog.Any("data", pRspTransfer),
	)
}

func (spi *ThostFutureBase) OnRtnRepealFromBankToFutureByFutureManual(pRspRepeal *CThostFtdcRspRepealField) {
	spi.Log(
		spi.ctx, slog.LevelDebug-1,
		"thost trader [OnRtnRepealFromBankToFutureByFutureManual]",
		slog.Any("data", pRspRepeal),
	)
}

func (spi *ThostFutureBase) OnRtnRepealFromFutureToBankByFutureManual(pRspRepeal *CThostFtdcRspRepealField) {
	spi.Log(
		spi.ctx, slog.LevelDebug-1,
		"thost trader [OnRtnRepealFromFutureToBankByFutureManual]",
		slog.Any("data", pRspRepeal),
	)
}

func (spi *ThostFutureBase) OnRtnQueryBankBalanceByFuture(pNotifyQueryAccount *CThostFtdcNotifyQueryAccountField) {
	spi.Log(
		spi.ctx, slog.LevelDebug-1,
		"thost trader [OnRtnQueryBankBalanceByFuture]",
		slog.Any("data", pNotifyQueryAccount),
	)
}

func (spi *ThostFutureBase) OnErrRtnBankToFutureByFuture(
	pReqTransfer *CThostFtdcReqTransferField, pRspInfo *CThostFtdcRspInfoField,
) {
	spi.Error(
		"thost trader [OnRtnBankToFutureByFuture] error",
		slog.Any("error", spi.CheckRsp(pRspInfo)),
		slog.Any("data", pReqTransfer),
	)
}

func (spi *ThostFutureBase) OnErrRtnFutureToBankByFuture(
	pReqTransfer *CThostFtdcReqTransferField, pRspInfo *CThostFtdcRspInfoField,
) {
	spi.Error(
		"thost trader [OnRtnFutureToBankByFuture] error",
		slog.Any("error", spi.CheckRsp(pRspInfo)),
		slog.Any("data", pReqTransfer),
	)
}

func (spi *ThostFutureBase) OnErrRtnRepealBankToFutureByFutureManual(
	pReqRepeal *CThostFtdcReqRepealField, pRspInfo *CThostFtdcRspInfoField,
) {
	spi.Error(
		"thost trader [OnRtnRepealBankToFutureByFutureManual] error",
		slog.Any("error", spi.CheckRsp(pRspInfo)),
		slog.Any("data", pReqRepeal),
	)
}

func (spi *ThostFutureBase) OnErrRtnRepealFutureToBankByFutureManual(
	pReqRepeal *CThostFtdcReqRepealField, pRspInfo *CThostFtdcRspInfoField,
) {
	spi.Error(
		"thost trader [OnRtnRepealFutureToBankByFutureManual] error",
		slog.Any("error", spi.CheckRsp(pRspInfo)),
		slog.Any("data", pReqRepeal),
	)
}

func (spi *ThostFutureBase) OnErrRtnQueryBankBalanceByFuture(
	pReqQueryAccount *CThostFtdcReqQueryAccountField, pRspInfo *CThostFtdcRspInfoField,
) {
	spi.Error(
		"thost trader [OnRtnQueryBankBalanceByFuture] error",
		slog.Any("error", spi.CheckRsp(pRspInfo)),
		slog.Any("data", pReqQueryAccount),
	)
}

func (spi *ThostFutureBase) OnRtnRepealFromBankToFutureByFuture(pRspRepeal *CThostFtdcRspRepealField) {
	spi.Log(
		spi.ctx, slog.LevelDebug-1,
		"thost trader [OnRtnRepealFromBankToFutureByFuture]",
		slog.Any("data", pRspRepeal),
	)
}

func (spi *ThostFutureBase) OnRtnRepealFromFutureToBankByFuture(pRspRepeal *CThostFtdcRspRepealField) {
	spi.Log(
		spi.ctx, slog.LevelDebug-1,
		"thost trader [OnRtnRepealFromFutureToBankByFuture]",
		slog.Any("data", pRspRepeal),
	)
}

func (spi *ThostFutureBase) OnRspFromBankToFutureByFuture(
	pReqTransfer *CThostFtdcReqTransferField,
	pRspInfo *CThostFtdcRspInfoField,
	nRequestID int, bIsLast bool,
) {
	err := spi.CheckRsp(pRspInfo)
	if err != nil {
		spi.Error(
			"thost trader [OnRspFromBankToFutureByFuture] failed",
			slog.Any("error", err),
			slog.Int("request_id", nRequestID),
			slog.Bool("is_last", bIsLast),
		)
		return
	}

	spi.Info(
		"thost trader [OnRspFromBankToFutureByFuture] succeeded",
		slog.Any("data", pReqTransfer),
	)
}

func (spi *ThostFutureBase) OnRspFromFutureToBankByFuture(
	pReqTransfer *CThostFtdcReqTransferField,
	pRspInfo *CThostFtdcRspInfoField,
	nRequestID int, bIsLast bool,
) {
	err := spi.CheckRsp(pRspInfo)
	if err != nil {
		spi.Error(
			"thost trader [OnRspFromFutureToBankByFuture] failed",
			slog.Any("error", err),
			slog.Int("request_id", nRequestID),
			slog.Bool("is_last", bIsLast),
		)
		return
	}

	spi.Info(
		"thost trader [OnRspFromFutureToBankByFuture] succeeded",
		slog.Any("data", pReqTransfer),
	)
}

func (spi *ThostFutureBase) OnRspQueryBankAccountMoneyByFuture(
	pReqQueryAccount *CThostFtdcReqQueryAccountField,
	pRspInfo *CThostFtdcRspInfoField,
	nRequestID int, bIsLast bool,
) {
	err := spi.CheckRsp(pRspInfo)
	if err != nil {
		spi.Error(
			"thost trader [OnRspQueryBankAccountMoneyByFuture] failed",
			slog.Any("error", err),
			slog.Int("request_id", nRequestID),
			slog.Bool("is_last", bIsLast),
		)
		return
	}

	spi.Info(
		"thost trader [OnRspQueryBankAccountMoneyByFuture] succeeded",
		slog.Any("data", pReqQueryAccount),
	)
}

func (spi *ThostFutureBase) OnRtnOpenAccountByBank(pOpenAccount *CThostFtdcOpenAccountField) {
	spi.Log(
		spi.ctx, slog.LevelDebug-1,
		"thost trader [OnRtnOpenAccountByBank]",
		slog.Any("data", pOpenAccount),
	)
}

func (spi *ThostFutureBase) OnRtnCancelAccountByBank(pCancelAccount *CThostFtdcCancelAccountField) {
	spi.Log(
		spi.ctx, slog.LevelDebug-1,
		"thost trader [OnRtnCancelAccountByBank]",
		slog.Any("data", pCancelAccount),
	)
}

func (spi *ThostFutureBase) OnRtnChangeAccountByBank(pChangeAccount *CThostFtdcChangeAccountField) {
	spi.Log(
		spi.ctx, slog.LevelDebug-1,
		"thost trader [OnRtnChangeAccountByBank]",
		slog.Any("data", pChangeAccount),
	)
}

func (spi *ThostFutureBase) OnRspQryClassifiedInstrument(
	pInstrument *CThostFtdcInstrumentField,
	pRspInfo *CThostFtdcRspInfoField,
	nRequestID int, bIsLast bool,
) {
	err := spi.CheckRsp(pRspInfo)
	if err != nil {
		spi.Error(
			"thost trader [OnRspQryClassifiedInstrument] failed",
			slog.Any("error", err),
			slog.Int("request_id", nRequestID),
			slog.Bool("is_last", bIsLast),
		)
		return
	}

	spi.Info(
		"thost trader [OnRspQryClassifiedInstrument] succeeded",
		slog.Any("data", pInstrument),
	)
}

func (spi *ThostFutureBase) OnRspQryCombPromotionParam(
	pCombPromotionParam *CThostFtdcCombPromotionParamField,
	pRspInfo *CThostFtdcRspInfoField,
	nRequestID int, bIsLast bool,
) {
	err := spi.CheckRsp(pRspInfo)
	if err != nil {
		spi.Error(
			"thost trader [OnRspQryCombPromotionParam] failed",
			slog.Any("error", err),
			slog.Int("request_id", nRequestID),
			slog.Bool("is_last", bIsLast),
		)
		return
	}

	spi.Info(
		"thost trader [OnRspQryCombPromotionParam] succeeded",
		slog.Any("data", pCombPromotionParam),
	)
}

func (spi *ThostFutureBase) OnRspQryRiskSettleInvstPosition(
	pRiskSettleInvstPosition *CThostFtdcRiskSettleInvstPositionField,
	pRspInfo *CThostFtdcRspInfoField,
	nRequestID int, bIsLast bool,
) {
	err := spi.CheckRsp(pRspInfo)
	if err != nil {
		spi.Error(
			"thost trader [OnRspQryRiskSettleInvstPosition] failed",
			slog.Any("error", err),
			slog.Int("request_id", nRequestID),
			slog.Bool("is_last", bIsLast),
		)
		return
	}

	spi.Info(
		"thost trader [OnRspQryRiskSettleInvstPosition] succeeded",
		slog.Any("data", pRiskSettleInvstPosition),
	)
}

func (spi *ThostFutureBase) OnRspQryRiskSettleProductStatus(
	pRiskSettleProductStatus *CThostFtdcRiskSettleProductStatusField,
	pRspInfo *CThostFtdcRspInfoField,
	nRequestID int, bIsLast bool,
) {
	err := spi.CheckRsp(pRspInfo)
	if err != nil {
		spi.Error(
			"thost trader [OnRspQryRiskSettleProductStatus] failed",
			slog.Any("error", err),
			slog.Int("request_id", nRequestID),
			slog.Bool("is_last", bIsLast),
		)
		return
	}

	spi.Info(
		"thost trader [OnRspQryRiskSettleProductStatus] succeeded",
		slog.Any("data", pRiskSettleProductStatus),
	)
}

func (spi *ThostFutureBase) OnRspQrySPBMFutureParameter(
	pSPBMFutureParameter *CThostFtdcSPBMFutureParameterField,
	pRspInfo *CThostFtdcRspInfoField,
	nRequestID int, bIsLast bool,
) {
	err := spi.CheckRsp(pRspInfo)
	if err != nil {
		spi.Error(
			"thost trader [OnRspQrySPBMFutureParameter] failed",
			slog.Any("error", err),
			slog.Int("request_id", nRequestID),
			slog.Bool("is_last", bIsLast),
		)
		return
	}

	spi.Info(
		"thost trader [OnRspQrySPBMFutureParameter] succeeded",
		slog.Any("data", pSPBMFutureParameter),
	)
}

func (spi *ThostFutureBase) OnRspQrySPBMOptionParameter(
	pSPBMOptionParameter *CThostFtdcSPBMOptionParameterField,
	pRspInfo *CThostFtdcRspInfoField,
	nRequestID int, bIsLast bool,
) {
	err := spi.CheckRsp(pRspInfo)
	if err != nil {
		spi.Error(
			"thost trader [OnRspQrySPBMOptionParameter] failed",
			slog.Any("error", err),
			slog.Int("request_id", nRequestID),
			slog.Bool("is_last", bIsLast),
		)
		return
	}

	spi.Info(
		"thost trader [OnRspQrySPBMOptionParameter] succeeded",
		slog.Any("data", pSPBMOptionParameter),
	)
}

func (spi *ThostFutureBase) OnRspQrySPBMIntraParameter(
	pSPBMIntraParameter *CThostFtdcSPBMIntraParameterField,
	pRspInfo *CThostFtdcRspInfoField,
	nRequestID int, bIsLast bool,
) {
	err := spi.CheckRsp(pRspInfo)
	if err != nil {
		spi.Error(
			"thost trader [OnRspQrySPBMIntraParameter] failed",
			slog.Any("error", err),
			slog.Int("request_id", nRequestID),
			slog.Bool("is_last", bIsLast),
		)
		return
	}

	spi.Info(
		"thost trader [OnRspQrySPBMIntraParameter] succeeded",
		slog.Any("data", pSPBMIntraParameter),
	)
}

func (spi *ThostFutureBase) OnRspQrySPBMInterParameter(
	pSPBMInterParameter *CThostFtdcSPBMInterParameterField,
	pRspInfo *CThostFtdcRspInfoField,
	nRequestID int, bIsLast bool,
) {
	err := spi.CheckRsp(pRspInfo)
	if err != nil {
		spi.Error(
			"thost trader [OnRspQrySPBMInterParameter] failed",
			slog.Any("error", err),
			slog.Int("request_id", nRequestID),
			slog.Bool("is_last", bIsLast),
		)
		return
	}

	spi.Info(
		"thost trader [OnRspQrySPBMInterParameter] succeeded",
		slog.Any("data", pSPBMInterParameter),
	)
}

func (spi *ThostFutureBase) OnRspQrySPBMPortfDefinition(
	pSPBMPortfDefinition *CThostFtdcSPBMPortfDefinitionField,
	pRspInfo *CThostFtdcRspInfoField,
	nRequestID int, bIsLast bool,
) {
	err := spi.CheckRsp(pRspInfo)
	if err != nil {
		spi.Error(
			"thost trader [OnRspQrySPBMPortfDefinition] failed",
			slog.Any("error", err),
			slog.Int("request_id", nRequestID),
			slog.Bool("is_last", bIsLast),
		)
		return
	}

	spi.Info(
		"thost trader [OnRspQrySPBMPortfDefinition] succeeded",
		slog.Any("data", pSPBMPortfDefinition),
	)
}

func (spi *ThostFutureBase) OnRspQrySPBMInvestorPortfDef(
	pSPBMInvestorPortfDef *CThostFtdcSPBMInvestorPortfDefField,
	pRspInfo *CThostFtdcRspInfoField,
	nRequestID int, bIsLast bool,
) {
	err := spi.CheckRsp(pRspInfo)
	if err != nil {
		spi.Error(
			"thost trader [OnRspQrySPBMInvestorPortfDef] failed",
			slog.Any("error", err),
			slog.Int("request_id", nRequestID),
			slog.Bool("is_last", bIsLast),
		)
		return
	}

	spi.Info(
		"thost trader [OnRspQrySPBMInvestorPortfDef] succeeded",
		slog.Any("data", pSPBMInvestorPortfDef),
	)
}

func (spi *ThostFutureBase) OnRspQryInvestorPortfMarginRatio(
	pInvestorPortfMarginRatio *CThostFtdcInvestorPortfMarginRatioField,
	pRspInfo *CThostFtdcRspInfoField,
	nRequestID int, bIsLast bool,
) {
	err := spi.CheckRsp(pRspInfo)
	if err != nil {
		spi.Error(
			"thost trader [OnRspQryInvestorPortfMarginRatio] failed",
			slog.Any("error", err),
			slog.Int("request_id", nRequestID),
			slog.Bool("is_last", bIsLast),
		)
		return
	}

	spi.Info(
		"thost trader [OnRspQryInvestorPortfMarginRatio] succeeded",
		slog.Any("data", pInvestorPortfMarginRatio),
	)
}

func (spi *ThostFutureBase) OnRspQryInvestorProdSPBMDetail(
	pInvestorProdSPBMDetail *CThostFtdcInvestorProdSPBMDetailField,
	pRspInfo *CThostFtdcRspInfoField,
	nRequestID int, bIsLast bool,
) {
	err := spi.CheckRsp(pRspInfo)
	if err != nil {
		spi.Error(
			"thost trader [OnRspQryInvestorProdSPBMDetail] failed",
			slog.Any("error", err),
			slog.Int("request_id", nRequestID),
			slog.Bool("is_last", bIsLast),
		)
		return
	}

	spi.Info(
		"thost trader [OnRspQryInvestorProdSPBMDetail] succeeded",
		slog.Any("data", pInvestorProdSPBMDetail),
	)
}

func (spi *ThostFutureBase) OnRspQryInvestorCommoditySPMMMargin(
	pInvestorCommoditySPMMMargin *CThostFtdcInvestorCommoditySPMMMarginField,
	pRspInfo *CThostFtdcRspInfoField,
	nRequestID int, bIsLast bool,
) {
	err := spi.CheckRsp(pRspInfo)
	if err != nil {
		spi.Error(
			"thost trader [OnRspQryInvestorCommoditySPMMMargin] failed",
			slog.Any("error", err),
			slog.Int("request_id", nRequestID),
			slog.Bool("is_last", bIsLast),
		)
		return
	}

	spi.Info(
		"thost trader [OnRspQryInvestorCommoditySPMMMargin] succeeded",
		slog.Any("data", pInvestorCommoditySPMMMargin),
	)
}

func (spi *ThostFutureBase) OnRspQryInvestorCommodityGroupSPMMMargin(
	pInvestorCommodityGroupSPMMMargin *CThostFtdcInvestorCommodityGroupSPMMMarginField,
	pRspInfo *CThostFtdcRspInfoField,
	nRequestID int, bIsLast bool,
) {
	err := spi.CheckRsp(pRspInfo)
	if err != nil {
		spi.Error(
			"thost trader [OnRspQryInvestorCommodityGroupSPMMMargin] failed",
			slog.Any("error", err),
			slog.Int("request_id", nRequestID),
			slog.Bool("is_last", bIsLast),
		)
		return
	}

	spi.Info(
		"thost trader [OnRspQryInvestorCommodityGroupSPMMMargin] succeeded",
		slog.Any("data", pInvestorCommodityGroupSPMMMargin),
	)
}

func (spi *ThostFutureBase) OnRspQrySPMMInstParam(
	pSPMMInstParam *CThostFtdcSPMMInstParamField,
	pRspInfo *CThostFtdcRspInfoField,
	nRequestID int, bIsLast bool,
) {
	err := spi.CheckRsp(pRspInfo)
	if err != nil {
		spi.Error(
			"thost trader [OnRspQrySPMMInstParam] failed",
			slog.Any("error", err),
			slog.Int("request_id", nRequestID),
			slog.Bool("is_last", bIsLast),
		)
		return
	}

	spi.Info(
		"thost trader [OnRspQrySPMMInstParam] succeeded",
		slog.Any("data", pSPMMInstParam),
	)
}

func (spi *ThostFutureBase) OnRspQrySPMMProductParam(
	pSPMMProductParam *CThostFtdcSPMMProductParamField,
	pRspInfo *CThostFtdcRspInfoField,
	nRequestID int, bIsLast bool,
) {
	err := spi.CheckRsp(pRspInfo)
	if err != nil {
		spi.Error(
			"thost trader [OnRspQrySPMMProductParam] failed",
			slog.Any("error", err),
			slog.Int("request_id", nRequestID),
			slog.Bool("is_last", bIsLast),
		)
		return
	}

	spi.Info(
		"thost trader [OnRspQrySPMMProductParam] succeeded",
		slog.Any("data", pSPMMProductParam),
	)
}

func (spi *ThostFutureBase) OnRspQrySPBMAddOnInterParameter(
	pSPBMAddOnInterParameter *CThostFtdcSPBMAddOnInterParameterField,
	pRspInfo *CThostFtdcRspInfoField,
	nRequestID int, bIsLast bool,
) {
	err := spi.CheckRsp(pRspInfo)
	if err != nil {
		spi.Error(
			"thost trader [OnRspQrySPBMAddOnInterParameter] failed",
			slog.Any("error", err),
			slog.Int("request_id", nRequestID),
			slog.Bool("is_last", bIsLast),
		)
		return
	}

	spi.Info(
		"thost trader [OnRspQrySPBMAddOnInterParameter] succeeded",
		slog.Any("data", pSPBMAddOnInterParameter),
	)
}

func (spi *ThostFutureBase) OnRspQryRCAMSCombProductInfo(
	pRCAMSCombProductInfo *CThostFtdcRCAMSCombProductInfoField,
	pRspInfo *CThostFtdcRspInfoField,
	nRequestID int, bIsLast bool,
) {
	err := spi.CheckRsp(pRspInfo)
	if err != nil {
		spi.Error(
			"thost trader [OnRspQryRCAMSCombProductInfo] failed",
			slog.Any("error", err),
			slog.Int("request_id", nRequestID),
			slog.Bool("is_last", bIsLast),
		)
		return
	}

	spi.Info(
		"thost trader [OnRspQryRCAMSCombProductInfo] succeeded",
		slog.Any("data", pRCAMSCombProductInfo),
	)
}

func (spi *ThostFutureBase) OnRspQryRCAMSInstrParameter(
	pRCAMSInstrParameter *CThostFtdcRCAMSInstrParameterField,
	pRspInfo *CThostFtdcRspInfoField,
	nRequestID int, bIsLast bool,
) {
	err := spi.CheckRsp(pRspInfo)
	if err != nil {
		spi.Error(
			"thost trader [OnRspQryRCAMSInstrParameter] failed",
			slog.Any("error", err),
			slog.Int("request_id", nRequestID),
			slog.Bool("is_last", bIsLast),
		)
		return
	}

	spi.Info(
		"thost trader [OnRspQryRCAMSInstrParameter] succeeded",
		slog.Any("data", pRCAMSInstrParameter),
	)
}

func (spi *ThostFutureBase) OnRspQryRCAMSIntraParameter(
	pRCAMSIntraParameter *CThostFtdcRCAMSIntraParameterField,
	pRspInfo *CThostFtdcRspInfoField,
	nRequestID int, bIsLast bool,
) {
	err := spi.CheckRsp(pRspInfo)
	if err != nil {
		spi.Error(
			"thost trader [OnRspQryRCAMSIntraParameter] failed",
			slog.Any("error", err),
			slog.Int("request_id", nRequestID),
			slog.Bool("is_last", bIsLast),
		)
		return
	}

	spi.Info(
		"thost trader [OnRspQryRCAMSIntraParameter] succeeded",
		slog.Any("data", pRCAMSIntraParameter),
	)
}

func (spi *ThostFutureBase) OnRspQryRCAMSInterParameter(
	pRCAMSInterParameter *CThostFtdcRCAMSInterParameterField,
	pRspInfo *CThostFtdcRspInfoField,
	nRequestID int, bIsLast bool,
) {
	err := spi.CheckRsp(pRspInfo)
	if err != nil {
		spi.Error(
			"thost trader [OnRspQryRCAMSInterParameter] failed",
			slog.Any("error", err),
			slog.Int("request_id", nRequestID),
			slog.Bool("is_last", bIsLast),
		)
		return
	}

	spi.Info(
		"thost trader [OnRspQryRCAMSInterParameter] succeeded",
		slog.Any("data", pRCAMSInterParameter),
	)
}

func (spi *ThostFutureBase) OnRspQryRCAMSShortOptAdjustParam(
	pRCAMSShortOptAdjustParam *CThostFtdcRCAMSShortOptAdjustParamField,
	pRspInfo *CThostFtdcRspInfoField,
	nRequestID int, bIsLast bool,
) {
	err := spi.CheckRsp(pRspInfo)
	if err != nil {
		spi.Error(
			"thost trader [OnRspQryRCAMSShortOptAdjustParam] failed",
			slog.Any("error", err),
			slog.Int("request_id", nRequestID),
			slog.Bool("is_last", bIsLast),
		)
		return
	}

	spi.Info(
		"thost trader [OnRspQryRCAMSShortOptAdjustParam] succeeded",
		slog.Any("data", pRCAMSShortOptAdjustParam),
	)
}

func (spi *ThostFutureBase) OnRspQryRCAMSInvestorCombPosition(
	pRCAMSInvestorCombPosition *CThostFtdcRCAMSInvestorCombPositionField,
	pRspInfo *CThostFtdcRspInfoField,
	nRequestID int, bIsLast bool,
) {
	err := spi.CheckRsp(pRspInfo)
	if err != nil {
		spi.Error(
			"thost trader [OnRspQryRCAMSInvestorCombPosition] failed",
			slog.Any("error", err),
			slog.Int("request_id", nRequestID),
			slog.Bool("is_last", bIsLast),
		)
		return
	}

	spi.Info(
		"thost trader [OnRspQryRCAMSInvestorCombPosition] succeeded",
		slog.Any("data", pRCAMSInvestorCombPosition),
	)
}

func (spi *ThostFutureBase) OnRspQryInvestorProdRCAMSMargin(
	pInvestorProdRCAMSMargin *CThostFtdcInvestorProdRCAMSMarginField,
	pRspInfo *CThostFtdcRspInfoField,
	nRequestID int, bIsLast bool,
) {
	err := spi.CheckRsp(pRspInfo)
	if err != nil {
		spi.Error(
			"thost trader [OnRspQryInvestorProdRCAMSMargin] failed",
			slog.Any("error", err),
			slog.Int("request_id", nRequestID),
			slog.Bool("is_last", bIsLast),
		)
		return
	}

	spi.Info(
		"thost trader [OnRspQryInvestorProdRCAMSMargin] succeeded",
		slog.Any("data", pInvestorProdRCAMSMargin),
	)
}

func (spi *ThostFutureBase) OnRspQryRULEInstrParameter(
	pRULEInstrParameter *CThostFtdcRULEInstrParameterField,
	pRspInfo *CThostFtdcRspInfoField,
	nRequestID int, bIsLast bool,
) {
	err := spi.CheckRsp(pRspInfo)
	if err != nil {
		spi.Error(
			"thost trader [OnRspQryRULEInstrParameter] failed",
			slog.Any("error", err),
			slog.Int("request_id", nRequestID),
			slog.Bool("is_last", bIsLast),
		)
		return
	}

	spi.Info(
		"thost trader [OnRspQryRULEInstrParameter] succeeded",
		slog.Any("data", pRULEInstrParameter),
	)
}

func (spi *ThostFutureBase) OnRspQryRULEIntraParameter(
	pRULEIntraParameter *CThostFtdcRULEIntraParameterField,
	pRspInfo *CThostFtdcRspInfoField,
	nRequestID int, bIsLast bool,
) {
	err := spi.CheckRsp(pRspInfo)
	if err != nil {
		spi.Error(
			"thost trader [OnRspQryRULEIntraParameter] failed",
			slog.Any("error", err),
			slog.Int("request_id", nRequestID),
			slog.Bool("is_last", bIsLast),
		)
		return
	}

	spi.Info(
		"thost trader [OnRspQryRULEIntraParameter] succeeded",
		slog.Any("data", pRULEIntraParameter),
	)
}

func (spi *ThostFutureBase) OnRspQryRULEInterParameter(
	pRULEInterParameter *CThostFtdcRULEInterParameterField,
	pRspInfo *CThostFtdcRspInfoField,
	nRequestID int, bIsLast bool,
) {
	err := spi.CheckRsp(pRspInfo)
	if err != nil {
		spi.Error(
			"thost trader [OnRspQryRULEInterParameter] failed",
			slog.Any("error", err),
			slog.Int("request_id", nRequestID),
			slog.Bool("is_last", bIsLast),
		)
		return
	}

	spi.Info(
		"thost trader [OnRspQryRULEInterParameter] succeeded",
		slog.Any("data", pRULEInterParameter),
	)
}

func (spi *ThostFutureBase) OnRspQryInvestorProdRULEMargin(
	pInvestorProdRULEMargin *CThostFtdcInvestorProdRULEMarginField,
	pRspInfo *CThostFtdcRspInfoField,
	nRequestID int, bIsLast bool,
) {
	err := spi.CheckRsp(pRspInfo)
	if err != nil {
		spi.Error(
			"thost trader [OnRspQryInvestorProdRULEMargin] failed",
			slog.Any("error", err),
			slog.Int("request_id", nRequestID),
			slog.Bool("is_last", bIsLast),
		)
		return
	}

	spi.Info(
		"thost trader [OnRspQryInvestorProdRULEMargin] succeeded",
		slog.Any("data", pInvestorProdRULEMargin),
	)
}

func (spi *ThostFutureBase) OnRspQryInvestorPortfSetting(
	pInvestorPortfSetting *CThostFtdcInvestorPortfSettingField,
	pRspInfo *CThostFtdcRspInfoField,
	nRequestID int, bIsLast bool,
) {
	err := spi.CheckRsp(pRspInfo)
	if err != nil {
		spi.Error(
			"thost trader [OnRspQryInvestorPortfSetting] failed",
			slog.Any("error", err),
			slog.Int("request_id", nRequestID),
			slog.Bool("is_last", bIsLast),
		)
		return
	}

	spi.Info(
		"thost trader [OnRspQryInvestorPortfSetting] succeeded",
		slog.Any("data", pInvestorPortfSetting),
	)
}

func (spi *ThostFutureBase) OnRspQryInvestorInfoCommRec(
	pInvestorInfoCommRec *CThostFtdcInvestorInfoCommRecField,
	pRspInfo *CThostFtdcRspInfoField,
	nRequestID int, bIsLast bool,
) {
	err := spi.CheckRsp(pRspInfo)
	if err != nil {
		spi.Error(
			"thost trader [OnRspQryInvestorInfoCommRec] failed",
			slog.Any("error", err),
			slog.Int("request_id", nRequestID),
			slog.Bool("is_last", bIsLast),
		)
		return
	}

	spi.Info(
		"thost trader [OnRspQryInvestorInfoCommRec] succeeded",
		slog.Any("data", pInvestorInfoCommRec),
	)
}

func (spi *ThostFutureBase) OnRspQryCombLeg(
	pCombLeg *CThostFtdcCombLegField,
	pRspInfo *CThostFtdcRspInfoField,
	nRequestID int, bIsLast bool,
) {
	err := spi.CheckRsp(pRspInfo)
	if err != nil {
		spi.Error(
			"thost trader [OnRspQryCombLeg] failed",
			slog.Any("error", err),
			slog.Int("request_id", nRequestID),
			slog.Bool("is_last", bIsLast),
		)
		return
	}

	spi.Info(
		"thost trader [OnRspQryCombLeg] succeeded",
		slog.Any("data", pCombLeg),
	)
}

func (spi *ThostFutureBase) OnRspOffsetSetting(
	pInputOffsetSetting *CThostFtdcInputOffsetSettingField,
	pRspInfo *CThostFtdcRspInfoField,
	nRequestID int, bIsLast bool,
) {
	err := spi.CheckRsp(pRspInfo)
	if err != nil {
		spi.Error(
			"thost trader [OnRspOffsetSetting] failed",
			slog.Any("error", err),
			slog.Int("request_id", nRequestID),
			slog.Bool("is_last", bIsLast),
		)
		return
	}

	spi.Info(
		"thost trader [OnRspOffsetSetting] succeeded",
		slog.Any("data", pInputOffsetSetting),
	)
}

func (spi *ThostFutureBase) OnRspCancelOffsetSetting(
	pInputOffsetSetting *CThostFtdcInputOffsetSettingField,
	pRspInfo *CThostFtdcRspInfoField,
	nRequestID int, bIsLast bool,
) {
	err := spi.CheckRsp(pRspInfo)
	if err != nil {
		spi.Error(
			"thost trader [OnRspCancelOffsetSetting] failed",
			slog.Any("error", err),
			slog.Int("request_id", nRequestID),
			slog.Bool("is_last", bIsLast),
		)
		return
	}

	spi.Info(
		"thost trader [OnRspCancelOffsetSetting] succeeded",
		slog.Any("data", pInputOffsetSetting),
	)
}

func (spi *ThostFutureBase) OnRtnOffsetSetting(pOffsetSetting *CThostFtdcOffsetSettingField) {
	spi.Log(
		spi.ctx, slog.LevelDebug-1,
		"thost trader [OnRtnOffsetSetting]",
		slog.Any("data", pOffsetSetting),
	)
}

func (spi *ThostFutureBase) OnErrRtnOffsetSetting(
	pInputOffsetSetting *CThostFtdcInputOffsetSettingField, pRspInfo *CThostFtdcRspInfoField,
) {
	spi.Error(
		"thost trader [OnRtnOffsetSetting] error",
		slog.Any("error", spi.CheckRsp(pRspInfo)),
		slog.Any("data", pInputOffsetSetting),
	)
}

func (spi *ThostFutureBase) OnErrRtnCancelOffsetSetting(
	pCancelOffsetSetting *CThostFtdcCancelOffsetSettingField, pRspInfo *CThostFtdcRspInfoField,
) {
	spi.Error(
		"thost trader [OnRtnCancelOffsetSetting] error",
		slog.Any("error", spi.CheckRsp(pRspInfo)),
		slog.Any("data", pCancelOffsetSetting),
	)
}

func (spi *ThostFutureBase) OnRspQryOffsetSetting(
	pOffsetSetting *CThostFtdcOffsetSettingField,
	pRspInfo *CThostFtdcRspInfoField,
	nRequestID int, bIsLast bool,
) {
	err := spi.CheckRsp(pRspInfo)
	if err != nil {
		spi.Error(
			"thost trader [OnRspQryOffsetSetting] failed",
			slog.Any("error", err),
			slog.Int("request_id", nRequestID),
			slog.Bool("is_last", bIsLast),
		)
		return
	}

	spi.Info(
		"thost trader [OnRspQryOffsetSetting] succeeded",
		slog.Any("data", pOffsetSetting),
	)
}

func (spi *ThostFutureBase) OnRspGenSMSCode(
	pRspGenSMSCode *CThostFtdcRspGenSMSCodeField,
	pRspInfo *CThostFtdcRspInfoField,
	nRequestID int, bIsLast bool,
) {
	err := spi.CheckRsp(pRspInfo)
	if err != nil {
		spi.Error(
			"thost trader [OnRspGenSMSCode] failed",
			slog.Any("error", err),
			slog.Int("request_id", nRequestID),
			slog.Bool("is_last", bIsLast),
		)
		return
	}

	spi.Info(
		"thost trader [OnRspGenSMSCode] succeeded",
		slog.Any("data", pRspGenSMSCode),
	)
}

func (spi *ThostFutureBase) OnRtnSMSVerifyInfoFromSec(pSMSVerifyInfoFromSec *CThostFtdcSMSVerifyInfoFromSecField) {
	spi.Log(
		spi.ctx, slog.LevelDebug-1,
		"thost trader [OnRtnSMSVerifyInfoFromSec]",
		slog.Any("data", pSMSVerifyInfoFromSec),
	)
}

func (spi *ThostFutureBase) OnRspSpdApply(
	pInputSpdApply *CThostFtdcInputSpdApplyField,
	pRspInfo *CThostFtdcRspInfoField,
	nRequestID int, bIsLast bool,
) {
	err := spi.CheckRsp(pRspInfo)
	if err != nil {
		spi.Error(
			"thost trader [OnRspSpdApply] failed",
			slog.Any("error", err),
			slog.Int("request_id", nRequestID),
			slog.Bool("is_last", bIsLast),
		)
		return
	}

	spi.Info(
		"thost trader [OnRspSpdApply] succeeded",
		slog.Any("data", pInputSpdApply),
	)
}

func (spi *ThostFutureBase) OnRspSpdApplyAction(
	pInputSpdApplyAction *CThostFtdcInputSpdApplyActionField,
	pRspInfo *CThostFtdcRspInfoField,
	nRequestID int, bIsLast bool,
) {
	err := spi.CheckRsp(pRspInfo)
	if err != nil {
		spi.Error(
			"thost trader [OnRspSpdApplyAction] failed",
			slog.Any("error", err),
			slog.Int("request_id", nRequestID),
			slog.Bool("is_last", bIsLast),
		)
		return
	}

	spi.Info(
		"thost trader [OnRspSpdApplyAction] succeeded",
		slog.Any("data", pInputSpdApplyAction),
	)
}

func (spi *ThostFutureBase) OnRspQrySpdApply(
	pSpdApply *CThostFtdcSpdApplyField,
	pRspInfo *CThostFtdcRspInfoField,
	nRequestID int, bIsLast bool,
) {
	err := spi.CheckRsp(pRspInfo)
	if err != nil {
		spi.Error(
			"thost trader [OnRspQrySpdApply] failed",
			slog.Any("error", err),
			slog.Int("request_id", nRequestID),
			slog.Bool("is_last", bIsLast),
		)
		return
	}

	spi.Log(
		spi.ctx, slog.LevelDebug-2,
		"thost trader [OnRspQrySpdApply] succeeded",
		slog.Any("data", pSpdApply),
	)
}

func (spi *ThostFutureBase) OnRtnSpdApply(pSpdApply *CThostFtdcSpdApplyField) {
	spi.Log(
		spi.ctx, slog.LevelDebug-1,
		"thost trader [OnRtnSpdApply]",
		slog.Any("data", pSpdApply),
	)
}

func (spi *ThostFutureBase) OnErrRtnSpdApply(
	pInputSpdApply *CThostFtdcInputSpdApplyField, pRspInfo *CThostFtdcRspInfoField,
) {
	spi.Error(
		"thost trader [OnRtnSpdApply] error",
		slog.Any("error", spi.CheckRsp(pRspInfo)),
		slog.Any("data", pInputSpdApply),
	)
}

func (spi *ThostFutureBase) OnErrRtnSpdApplyAction(
	pSpdApplyAction *CThostFtdcSpdApplyActionField, pRspInfo *CThostFtdcRspInfoField,
) {
	spi.Error(
		"thost trader [OnRtnSpdApplyAction] error",
		slog.Any("error", spi.CheckRsp(pRspInfo)),
		slog.Any("data", pSpdApplyAction),
	)
}

func (spi *ThostFutureBase) OnRspHedgeCfm(
	pInputHedgeCfm *CThostFtdcInputHedgeCfmField,
	pRspInfo *CThostFtdcRspInfoField,
	nRequestID int, bIsLast bool,
) {
	err := spi.CheckRsp(pRspInfo)
	if err != nil {
		spi.Error(
			"thost trader [OnRspHedgeCfm] failed",
			slog.Any("error", err),
			slog.Int("request_id", nRequestID),
			slog.Bool("is_last", bIsLast),
		)
		return
	}

	spi.Info(
		"thost trader [OnRspHedgeCfm] succeeded",
		slog.Any("data", pInputHedgeCfm),
	)
}

func (spi *ThostFutureBase) OnRspHedgeCfmAction(
	pInputHedgeCfmAction *CThostFtdcInputHedgeCfmActionField,
	pRspInfo *CThostFtdcRspInfoField,
	nRequestID int, bIsLast bool,
) {
	err := spi.CheckRsp(pRspInfo)
	if err != nil {
		spi.Error(
			"thost trader [OnRspHedgeCfmAction] failed",
			slog.Any("error", err),
			slog.Int("request_id", nRequestID),
			slog.Bool("is_last", bIsLast),
		)
		return
	}

	spi.Info(
		"thost trader [OnRspHedgeCfmAction] succeeded",
		slog.Any("data", pInputHedgeCfmAction),
	)
}

func (spi *ThostFutureBase) OnRspQryHedgeCfm(
	pHedgeCfm *CThostFtdcHedgeCfmField,
	pRspInfo *CThostFtdcRspInfoField,
	nRequestID int, bIsLast bool,
) {
	err := spi.CheckRsp(pRspInfo)
	if err != nil {
		spi.Error(
			"thost trader [OnRspQryHedgeCfm] failed",
			slog.Any("error", err),
			slog.Int("request_id", nRequestID),
			slog.Bool("is_last", bIsLast),
		)
		return
	}

	spi.Log(
		spi.ctx, slog.LevelDebug-2,
		"thost trader [OnRspQryHedgeCfm] succeeded",
		slog.Any("data", pHedgeCfm),
	)
}

func (spi *ThostFutureBase) OnRtnHedgeCfm(pHedgeCfm *CThostFtdcHedgeCfmField) {
	spi.Log(
		spi.ctx, slog.LevelDebug-1,
		"thost trader [OnRtnHedgeCfm]",
		slog.Any("data", pHedgeCfm),
	)
}

func (spi *ThostFutureBase) OnErrRtnHedgeCfm(
	pInputHedgeCfm *CThostFtdcInputHedgeCfmField, pRspInfo *CThostFtdcRspInfoField,
) {
	spi.Error(
		"thost trader [OnRtnHedgeCfm] error",
		slog.Any("error", spi.CheckRsp(pRspInfo)),
		slog.Any("data", pInputHedgeCfm),
	)
}

func (spi *ThostFutureBase) OnErrRtnHedgeCfmAction(
	pHedgeCfmAction *CThostFtdcHedgeCfmActionField, pRspInfo *CThostFtdcRspInfoField,
) {
	spi.Error(
		"thost trader [OnRtnHedgeCfmAction] error",
		slog.Any("error", spi.CheckRsp(pRspInfo)),
		slog.Any("data", pHedgeCfmAction),
	)
}

func (spi *ThostFutureBase) OnRspQryMulticastInstrument(
	pMulticastInstrument *CThostFtdcMulticastInstrumentField,
	pRspInfo *CThostFtdcRspInfoField,
	nRequestID int, bIsLast bool,
) {
	err := spi.CheckRsp(pRspInfo)
	if err != nil {
		spi.Error(
			"thost mduser [OnRspQryMulticastInstrument] failed",
			slog.Any("error", err),
			slog.Int("request_id", nRequestID),
			slog.Bool("is_last", bIsLast),
		)
		return
	}

	spi.Log(
		spi.ctx, slog.LevelDebug-2,
		"thost mduser [OnRspQryMulticastInstrument] succeeded",
		slog.Any("data", pMulticastInstrument),
	)
}

func (spi *ThostFutureBase) OnRspSubMarketData(
	pSpecificInstrument *CThostFtdcSpecificInstrumentField,
	pRspInfo *CThostFtdcRspInfoField,
	nRequestID int, bIsLast bool,
) {
	err := spi.CheckRsp(pRspInfo)
	if err != nil {
		spi.Error(
			"thost mduser [OnRspSubMarketData] failed",
			slog.Any("error", err),
			slog.Int("request_id", nRequestID),
			slog.Bool("is_last", bIsLast),
		)
		return
	}

	spi.Info(
		"thost mduser [OnRspSubMarketData] succeeded",
		slog.Any("data", pSpecificInstrument),
	)
}

func (spi *ThostFutureBase) OnRspUnSubMarketData(
	pSpecificInstrument *CThostFtdcSpecificInstrumentField,
	pRspInfo *CThostFtdcRspInfoField,
	nRequestID int, bIsLast bool,
) {
	err := spi.CheckRsp(pRspInfo)
	if err != nil {
		spi.Error(
			"thost mduser [OnRspUnSubMarketData] failed",
			slog.Any("error", err),
			slog.Int("request_id", nRequestID),
			slog.Bool("is_last", bIsLast),
		)
		return
	}

	spi.Info(
		"thost mduser [OnRspUnSubMarketData] succeeded",
		slog.Any("data", pSpecificInstrument),
	)
}

func (spi *ThostFutureBase) OnRspSubForQuoteRsp(
	pSpecificInstrument *CThostFtdcSpecificInstrumentField,
	pRspInfo *CThostFtdcRspInfoField,
	nRequestID int, bIsLast bool,
) {
	err := spi.CheckRsp(pRspInfo)
	if err != nil {
		spi.Error(
			"thost mduser [OnRspSubForQuoteRsp] failed",
			slog.Any("error", err),
			slog.Int("request_id", nRequestID),
			slog.Bool("is_last", bIsLast),
		)
		return
	}

	spi.Info(
		"thost mduser [OnRspSubForQuoteRsp] succeeded",
		slog.Any("data", pSpecificInstrument),
	)
}

func (spi *ThostFutureBase) OnRspUnSubForQuoteRsp(
	pSpecificInstrument *CThostFtdcSpecificInstrumentField,
	pRspInfo *CThostFtdcRspInfoField,
	nRequestID int, bIsLast bool,
) {
	err := spi.CheckRsp(pRspInfo)
	if err != nil {
		spi.Error(
			"thost mduser [OnRspUnSubForQuoteRsp] failed",
			slog.Any("error", err),
			slog.Int("request_id", nRequestID),
			slog.Bool("is_last", bIsLast),
		)
		return
	}

	spi.Info(
		"thost mduser [OnRspUnSubForQuoteRsp] succeeded",
		slog.Any("data", pSpecificInstrument),
	)
}

func (spi *ThostFutureBase) OnRtnDepthMarketData(
	pDepthMarketData *CThostFtdcDepthMarketDataField,
) {
	spi.Log(
		spi.ctx, slog.LevelDebug-2,
		"thost mduser [OnRtnDepthMarketData] succeeded",
		slog.Any("data", pDepthMarketData),
	)
}
