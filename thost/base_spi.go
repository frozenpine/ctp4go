package thost

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/frozenpine/ctp4go/thost/types"
)

var (
	_ TraderSpi = &TraderLogSpi{}
)

type TraderLogSpi struct {
	*slog.Logger

	FrontInfo CThostFtdcFrontInfoField
}

func (spi *TraderLogSpi) CheckRsp(rsp *CThostFtdcRspInfoField) error {
	if rsp == nil {
		return nil
	}

	if rsp.ErrorID != 0 {
		return fmt.Errorf(
			"[%d] %s", rsp.ErrorID, types.DecodeGBK(rsp.ErrorMsg[:]),
		)
	}

	return nil
}

func (spi *TraderLogSpi) OnFrontConnected() {
	if spi.Logger == nil {
		spi.Logger = slog.Default()
	}

	spi.Info(
		"thost trader [OnFrontConnected]",
		slog.Any("front", spi.FrontInfo),
	)
}

func (spi *TraderLogSpi) OnFrontDisconnected(nReason int) {
	spi.Info(
		"thost trader [OnFrontDisconnected]",
		slog.Any("front", spi.FrontInfo),
		slog.Int("reason", nReason),
	)
}

func (spi *TraderLogSpi) OnHeartBeatWarning(nTimeLapse int) {
	spi.Info(
		"thost trader [OnHeartBeatWarning]",
		slog.Any("front", spi.FrontInfo),
		slog.Int("time_lapse", nTimeLapse),
	)
}

func (spi *TraderLogSpi) OnRspAuthenticate(
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

func (spi *TraderLogSpi) OnRtnPrivateSeqNo(nSeqNo int) {
	spi.Log(
		context.Background(), slog.LevelDebug-2,
		"thost trader [OnRtnPrivateSeqNo]",
		slog.Int("seq_no", nSeqNo),
	)
}

func (spi *TraderLogSpi) OnRspUserLogin(
	pRspUserLogin *CThostFtdcRspUserLoginField,
	pRspInfo *CThostFtdcRspInfoField,
	nRequestID int, bIsLast bool,
) {
	err := spi.CheckRsp(pRspInfo)
	if err != nil {
		spi.Error(
			"thost trader [OnRspUserLogin] failed",
			slog.Any("error", err),
			slog.Int("request_id", nRequestID),
			slog.Bool("is_last", bIsLast),
		)
		return
	}

	spi.Info(
		"thost trader [OnRspUserLogin] succeeded",
		slog.Any("login", pRspUserLogin),
	)
}

func (spi *TraderLogSpi) OnRspUserLogout(
	pUserLogout *CThostFtdcUserLogoutField,
	pRspInfo *CThostFtdcRspInfoField,
	nRequestID int, bIsLast bool,
) {
	err := spi.CheckRsp(pRspInfo)
	if err != nil {
		spi.Error(
			"thost trader [OnRspUserLogout] failed",
			slog.Any("error", err),
			slog.Int("request_id", nRequestID),
			slog.Bool("is_last", bIsLast),
		)
		return
	}

	spi.Info(
		"thost trader [OnRspUserLogout] succeeded",
		slog.Any("logout", pUserLogout),
	)
}

func (spi *TraderLogSpi) OnRspUserPasswordUpdate(
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

func (spi *TraderLogSpi) OnRspTradingAccountPasswordUpdate(
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

func (spi *TraderLogSpi) OnRspUserAuthMethod(
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

func (spi *TraderLogSpi) OnRspGenUserCaptcha(
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

func (spi *TraderLogSpi) OnRspGenUserText(
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

func (spi *TraderLogSpi) OnRspOrderInsert(
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

func (spi *TraderLogSpi) OnRspParkedOrderInsert(
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

func (spi *TraderLogSpi) OnRspParkedOrderAction(
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

func (spi *TraderLogSpi) OnRspOrderAction(
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

func (spi *TraderLogSpi) OnRspQryMaxOrderVolume(
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

	spi.Info(
		"thost trader [OnRspQryMaxOrderVolume] succeeded",
		slog.Any("data", pQryMaxOrderVolume),
	)
}

func (spi *TraderLogSpi) OnRspSettlementInfoConfirm(
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

func (spi *TraderLogSpi) OnRspRemoveParkedOrder(
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

func (spi *TraderLogSpi) OnRspRemoveParkedOrderAction(
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

func (spi *TraderLogSpi) OnRspExecOrderInsert(
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

func (spi *TraderLogSpi) OnRspExecOrderAction(
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

func (spi *TraderLogSpi) OnRspForQuoteInsert(
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

func (spi *TraderLogSpi) OnRspQuoteInsert(
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

func (spi *TraderLogSpi) OnRspQuoteAction(
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

func (spi *TraderLogSpi) OnRspBatchOrderAction(
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

func (spi *TraderLogSpi) OnRspOptionSelfCloseInsert(
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

func (spi *TraderLogSpi) OnRspOptionSelfCloseAction(
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

func (spi *TraderLogSpi) OnRspCombActionInsert(
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

func (spi *TraderLogSpi) OnRspQryOrder(
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

	spi.Info(
		"thost trader [OnRspQryOrder] succeeded",
		slog.Any("data", pOrder),
	)
}

func (spi *TraderLogSpi) OnRspQryTrade(
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

	spi.Info(
		"thost trader [OnRspQryTrade] succeeded",
		slog.Any("data", pTrade),
	)
}

func (spi *TraderLogSpi) OnRspQryInvestorPosition(
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

	spi.Info(
		"thost trader [OnRspQryInvestorPosition] succeeded",
		slog.Any("data", pInvestorPosition),
	)
}

func (spi *TraderLogSpi) OnRspQryTradingAccount(
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

	spi.Info(
		"thost trader [OnRspQryTradingAccount] succeeded",
		slog.Any("data", pTradingAccount),
	)
}

func (spi *TraderLogSpi) OnRspQryInvestor(
	pInvestor *CThostFtdcInvestorField,
	pRspInfo *CThostFtdcRspInfoField,
	nRequestID int, bIsLast bool,
) {
	err := spi.CheckRsp(pRspInfo)
	if err != nil {
		spi.Error(
			"thost trader [OnRspQryInvestor] failed",
			slog.Any("error", err),
			slog.Int("request_id", nRequestID),
			slog.Bool("is_last", bIsLast),
		)
		return
	}

	spi.Info(
		"thost trader [OnRspQryInvestor] succeeded",
		slog.Any("data", pInvestor),
	)
}

func (spi *TraderLogSpi) OnRspQryTradingCode(
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

	spi.Info(
		"thost trader [OnRspQryTradingCode] succeeded",
		slog.Any("data", pTradingCode),
	)
}

func (spi *TraderLogSpi) OnRspQryInstrumentMarginRate(
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

	spi.Info(
		"thost trader [OnRspQryInstrumentMarginRate] succeeded",
		slog.Any("data", pInstrumentMarginRate),
	)
}

func (spi *TraderLogSpi) OnRspQryInstrumentCommissionRate(
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

	spi.Info(
		"thost trader [OnRspQryInstrumentCommissionRate] succeeded",
		slog.Any("data", pInstrumentCommissionRate),
	)
}

func (spi *TraderLogSpi) OnRspQryUserSession(
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

	spi.Info(
		"thost trader [OnRspQryUserSession] succeeded",
		slog.Any("data", pUserSession),
	)
}

func (spi *TraderLogSpi) OnRspQryExchange(
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

	spi.Info(
		"thost trader [OnRspQryExchange] succeeded",
		slog.Any("data", pExchange),
	)
}

func (spi *TraderLogSpi) OnRspQryProduct(
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

	spi.Info(
		"thost trader [OnRspQryProduct] succeeded",
		slog.Any("data", pProduct),
	)
}

func (spi *TraderLogSpi) OnRspQryInstrument(
	pInstrument *CThostFtdcInstrumentField,
	pRspInfo *CThostFtdcRspInfoField,
	nRequestID int, bIsLast bool,
) {
	err := spi.CheckRsp(pRspInfo)
	if err != nil {
		spi.Error(
			"thost trader [OnRspQryInstrument] failed",
			slog.Any("error", err),
			slog.Int("request_id", nRequestID),
			slog.Bool("is_last", bIsLast),
		)
		return
	}

	spi.Info(
		"thost trader [OnRspQryInstrument] succeeded",
		slog.Any("data", pInstrument),
	)
}

func (spi *TraderLogSpi) OnRspQryDepthMarketData(
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

	spi.Info(
		"thost trader [OnRspQryDepthMarketData] succeeded",
		slog.Any("data", pDepthMarketData),
	)
}

func (spi *TraderLogSpi) OnRspQryTraderOffer(
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

	spi.Info(
		"thost trader [OnRspQryTraderOffer] succeeded",
		slog.Any("data", pTraderOffer),
	)
}

func (spi *TraderLogSpi) OnRspQrySettlementInfo(
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

func (spi *TraderLogSpi) OnRspQryTransferBank(
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

	spi.Info(
		"thost trader [OnRspQryTransferBank] succeeded",
		slog.Any("data", pTransferBank),
	)
}

func (spi *TraderLogSpi) OnRspQryInvestorPositionDetail(
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

	spi.Info(
		"thost trader [OnRspQryInvestorPositionDetail] succeeded",
		slog.Any("data", pInvestorPositionDetail),
	)
}

func (spi *TraderLogSpi) OnRspQryNotice(
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

	spi.Info(
		"thost trader [OnRspQryNotice] succeeded",
		slog.Any("data", pNotice),
	)
}

func (spi *TraderLogSpi) OnRspQrySettlementInfoConfirm(
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

func (spi *TraderLogSpi) OnRspQryInvestorPositionCombineDetail(
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

	spi.Info(
		"thost trader [OnRspQryInvestorPositionCombineDetail] succeeded",
		slog.Any("data", pInvestorPositionCombineDetail),
	)
}

func (spi *TraderLogSpi) OnRspQryCFMMCTradingAccountKey(
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

	spi.Info(
		"thost trader [OnRspQryCFMMCTradingAccountKey] succeeded",
		slog.Any("data", pCFMMCTradingAccountKey),
	)
}

func (spi *TraderLogSpi) OnRspQryEWarrantOffset(
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

	spi.Info(
		"thost trader [OnRspQryEWarrantOffset] succeeded",
		slog.Any("data", pEWarrantOffset),
	)
}

func (spi *TraderLogSpi) OnRspQryInvestorProductGroupMargin(
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

	spi.Info(
		"thost trader [OnRspQryInvestorProductGroupMargin] succeeded",
		slog.Any("data", pInvestorProductGroupMargin),
	)
}

func (spi *TraderLogSpi) OnRspQryExchangeMarginRate(
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

	spi.Info(
		"thost trader [OnRspQryExchangeMarginRate] succeeded",
		slog.Any("data", pExchangeMarginRate),
	)
}

func (spi *TraderLogSpi) OnRspQryExchangeMarginRateAdjust(
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

	spi.Info(
		"thost trader [OnRspQryExchangeMarginRateAdjust] succeeded",
		slog.Any("data", pExchangeMarginRateAdjust),
	)
}

func (spi *TraderLogSpi) OnRspQryExchangeRate(
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

	spi.Info(
		"thost trader [OnRspQryExchangeRate] succeeded",
		slog.Any("data", pExchangeRate),
	)
}

func (spi *TraderLogSpi) OnRspQrySecAgentACIDMap(
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

	spi.Info(
		"thost trader [OnRspQrySecAgentACIDMap] succeeded",
		slog.Any("data", pSecAgentACIDMap),
	)
}

func (spi *TraderLogSpi) OnRspQryProductExchRate(
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

	spi.Info(
		"thost trader [OnRspQryProductExchRate] succeeded",
		slog.Any("data", pProductExchRate),
	)
}

func (spi *TraderLogSpi) OnRspQryProductGroup(
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

func (spi *TraderLogSpi) OnRspQryMMInstrumentCommissionRate(
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

func (spi *TraderLogSpi) OnRspQryMMOptionInstrCommRate(
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

func (spi *TraderLogSpi) OnRspQryInstrumentOrderCommRate(
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

func (spi *TraderLogSpi) OnRspQrySecAgentTradingAccount(
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

func (spi *TraderLogSpi) OnRspQrySecAgentCheckMode(
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

func (spi *TraderLogSpi) OnRspQrySecAgentTradeInfo(
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

func (spi *TraderLogSpi) OnRspQryOptionInstrTradeCost(
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

func (spi *TraderLogSpi) OnRspQryOptionInstrCommRate(
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

func (spi *TraderLogSpi) OnRspQryExecOrder(
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

func (spi *TraderLogSpi) OnRspQryForQuote(
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

func (spi *TraderLogSpi) OnRspQryQuote(
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

func (spi *TraderLogSpi) OnRspQryOptionSelfClose(
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

func (spi *TraderLogSpi) OnRspQryInvestUnit(
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

func (spi *TraderLogSpi) OnRspQryCombInstrumentGuard(
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

func (spi *TraderLogSpi) OnRspQryCombAction(
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

func (spi *TraderLogSpi) OnRspQryTransferSerial(
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

func (spi *TraderLogSpi) OnRspQryAccountregister(
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

func (spi *TraderLogSpi) OnRspError(pRspInfo *CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
	spi.Error(
		"thost trader [OnRspError]",
		slog.Any("error", spi.CheckRsp(pRspInfo)),
		slog.Int("request_id", nRequestID),
		slog.Bool("is_last", bIsLast),
	)
}

func (spi *TraderLogSpi) OnRtnOrder(pOrder *CThostFtdcOrderField) {
	spi.Log(
		context.Background(), slog.LevelDebug-1,
		"thost trader [OnRtnOrder]",
		slog.Any("data", pOrder),
	)
}

func (spi *TraderLogSpi) OnRtnTrade(pTrade *CThostFtdcTradeField) {
	spi.Log(
		context.Background(), slog.LevelDebug-1,
		"thost trader [OnRtnTrade]",
		slog.Any("data", pTrade),
	)
}

func (spi *TraderLogSpi) OnErrRtnOrderInsert(
	pInputOrder *CThostFtdcInputOrderField, pRspInfo *CThostFtdcRspInfoField,
) {
	spi.Error(
		"thost trader [OnRtnOrderInsert] error",
		slog.Any("error", spi.CheckRsp(pRspInfo)),
		slog.Any("data", pInputOrder),
	)
}

func (spi *TraderLogSpi) OnErrRtnOrderAction(
	pOrderAction *CThostFtdcOrderActionField, pRspInfo *CThostFtdcRspInfoField,
) {
	spi.Error(
		"thost trader [OnRtnOrderAction] error",
		slog.Any("error", spi.CheckRsp(pRspInfo)),
		slog.Any("data", pOrderAction),
	)
}

func (spi *TraderLogSpi) OnRtnInstrumentStatus(pInstrumentStatus *CThostFtdcInstrumentStatusField) {
	spi.Log(
		context.Background(), slog.LevelDebug-1,
		"thost trader [OnRtnInstrumentStatus]",
		slog.Any("data", pInstrumentStatus),
	)
}

func (spi *TraderLogSpi) OnRtnBulletin(pBulletin *CThostFtdcBulletinField) {
	spi.Log(
		context.Background(), slog.LevelDebug-1,
		"thost trader [OnRtnBulletin]",
		slog.Any("data", pBulletin),
	)
}

func (spi *TraderLogSpi) OnRtnTradingNotice(pTradingNoticeInfo *CThostFtdcTradingNoticeInfoField) {
	spi.Log(
		context.Background(), slog.LevelDebug-1,
		"thost trader [OnRtnTradingNotice]",
		slog.Any("data", pTradingNoticeInfo),
	)
}

func (spi *TraderLogSpi) OnRtnErrorConditionalOrder(pErrorConditionalOrder *CThostFtdcErrorConditionalOrderField) {
	spi.Log(
		context.Background(), slog.LevelDebug-1,
		"thost trader [OnRtnErrorConditionalOrder]",
		slog.Any("data", pErrorConditionalOrder),
	)
}

func (spi *TraderLogSpi) OnRtnExecOrder(pExecOrder *CThostFtdcExecOrderField) {
	spi.Log(
		context.Background(), slog.LevelDebug-1,
		"thost trader [OnRtnExecOrder]",
		slog.Any("data", pExecOrder),
	)
}

func (spi *TraderLogSpi) OnErrRtnExecOrderInsert(
	pInputExecOrder *CThostFtdcInputExecOrderField, pRspInfo *CThostFtdcRspInfoField,
) {
	spi.Error(
		"thost trader [OnRtnExecOrderInsert] error",
		slog.Any("error", spi.CheckRsp(pRspInfo)),
		slog.Any("data", pInputExecOrder),
	)
}

func (spi *TraderLogSpi) OnErrRtnExecOrderAction(
	pExecOrderAction *CThostFtdcExecOrderActionField, pRspInfo *CThostFtdcRspInfoField,
) {
	spi.Error(
		"thost trader [OnRtnExecOrderAction] error",
		slog.Any("error", spi.CheckRsp(pRspInfo)),
		slog.Any("data", pExecOrderAction),
	)
}

func (spi *TraderLogSpi) OnErrRtnForQuoteInsert(
	pInputForQuote *CThostFtdcInputForQuoteField, pRspInfo *CThostFtdcRspInfoField,
) {
	spi.Error(
		"thost trader [OnRtnForQuoteInsert] error",
		slog.Any("error", spi.CheckRsp(pRspInfo)),
		slog.Any("data", pInputForQuote),
	)
}

func (spi *TraderLogSpi) OnRtnQuote(pQuote *CThostFtdcQuoteField) {
	spi.Log(
		context.Background(), slog.LevelDebug-1,
		"thost trader [OnRtnQuote]",
		slog.Any("data", pQuote),
	)
}

func (spi *TraderLogSpi) OnErrRtnQuoteInsert(
	pInputQuote *CThostFtdcInputQuoteField, pRspInfo *CThostFtdcRspInfoField,
) {
	spi.Error(
		"thost trader [OnRtnQuoteInsert] error",
		slog.Any("error", spi.CheckRsp(pRspInfo)),
		slog.Any("data", pInputQuote),
	)
}

func (spi *TraderLogSpi) OnErrRtnQuoteAction(
	pQuoteAction *CThostFtdcQuoteActionField, pRspInfo *CThostFtdcRspInfoField,
) {
	spi.Error(
		"thost trader [OnRtnQuoteAction] error",
		slog.Any("error", spi.CheckRsp(pRspInfo)),
		slog.Any("data", pQuoteAction),
	)
}

func (spi *TraderLogSpi) OnRtnForQuoteRsp(pForQuoteRsp *CThostFtdcForQuoteRspField) {
	spi.Log(
		context.Background(), slog.LevelDebug-1,
		"thost trader [OnRtnForQuoteRsp]",
		slog.Any("data", pForQuoteRsp),
	)
}

func (spi *TraderLogSpi) OnRtnCFMMCTradingAccountToken(pCFMMCTradingAccountToken *CThostFtdcCFMMCTradingAccountTokenField) {
	spi.Log(
		context.Background(), slog.LevelDebug-1,
		"thost trader [OnRtnCFMMCTradingAccountToken]",
		slog.Any("data", pCFMMCTradingAccountToken),
	)
}

func (spi *TraderLogSpi) OnErrRtnBatchOrderAction(
	pBatchOrderAction *CThostFtdcBatchOrderActionField, pRspInfo *CThostFtdcRspInfoField,
) {
	spi.Error(
		"thost trader [OnRtnBatchOrderAction] error",
		slog.Any("error", spi.CheckRsp(pRspInfo)),
		slog.Any("data", pBatchOrderAction),
	)
}

func (spi *TraderLogSpi) OnRtnOptionSelfClose(pOptionSelfClose *CThostFtdcOptionSelfCloseField) {
	spi.Log(
		context.Background(), slog.LevelDebug-1,
		"thost trader [OnRtnOptionSelfClose]",
		slog.Any("data", pOptionSelfClose),
	)
}

func (spi *TraderLogSpi) OnErrRtnOptionSelfCloseInsert(
	pInputOptionSelfClose *CThostFtdcInputOptionSelfCloseField, pRspInfo *CThostFtdcRspInfoField,
) {
	spi.Error(
		"thost trader [OnRtnOptionSelfCloseInsert] error",
		slog.Any("error", spi.CheckRsp(pRspInfo)),
		slog.Any("data", pInputOptionSelfClose),
	)
}

func (spi *TraderLogSpi) OnErrRtnOptionSelfCloseAction(
	pOptionSelfCloseAction *CThostFtdcOptionSelfCloseActionField, pRspInfo *CThostFtdcRspInfoField,
) {
	spi.Error(
		"thost trader [OnRtnOptionSelfCloseAction] error",
		slog.Any("error", spi.CheckRsp(pRspInfo)),
		slog.Any("data", pOptionSelfCloseAction),
	)
}

func (spi *TraderLogSpi) OnRtnCombAction(pCombAction *CThostFtdcCombActionField) {
	spi.Log(
		context.Background(), slog.LevelDebug-1,
		"thost trader [OnRtnCombAction]",
		slog.Any("data", pCombAction),
	)
}

func (spi *TraderLogSpi) OnErrRtnCombActionInsert(
	pInputCombAction *CThostFtdcInputCombActionField, pRspInfo *CThostFtdcRspInfoField,
) {
	spi.Error(
		"thost trader [OnRtnCombActionInsert] error",
		slog.Any("error", spi.CheckRsp(pRspInfo)),
		slog.Any("data", pInputCombAction),
	)
}

func (spi *TraderLogSpi) OnRspQryContractBank(
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

func (spi *TraderLogSpi) OnRspQryParkedOrder(
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

func (spi *TraderLogSpi) OnRspQryParkedOrderAction(
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

func (spi *TraderLogSpi) OnRspQryTradingNotice(
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

func (spi *TraderLogSpi) OnRspQryBrokerTradingParams(
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

func (spi *TraderLogSpi) OnRspQryBrokerTradingAlgos(
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

func (spi *TraderLogSpi) OnRspQueryCFMMCTradingAccountToken(
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

func (spi *TraderLogSpi) OnRtnFromBankToFutureByBank(pRspTransfer *CThostFtdcRspTransferField) {
	spi.Log(
		context.Background(), slog.LevelDebug-1,
		"thost trader [OnRtnFromBankToFutureByBank]",
		slog.Any("data", pRspTransfer),
	)
}

func (spi *TraderLogSpi) OnRtnFromFutureToBankByBank(pRspTransfer *CThostFtdcRspTransferField) {
	spi.Log(
		context.Background(), slog.LevelDebug-1,
		"thost trader [OnRtnFromFutureToBankByBank]",
		slog.Any("data", pRspTransfer),
	)
}

func (spi *TraderLogSpi) OnRtnRepealFromBankToFutureByBank(pRspRepeal *CThostFtdcRspRepealField) {
	spi.Log(
		context.Background(), slog.LevelDebug-1,
		"thost trader [OnRtnRepealFromBankToFutureByBank]",
		slog.Any("data", pRspRepeal),
	)
}

func (spi *TraderLogSpi) OnRtnRepealFromFutureToBankByBank(pRspRepeal *CThostFtdcRspRepealField) {
	spi.Log(
		context.Background(), slog.LevelDebug-1,
		"thost trader [OnRtnRepealFromFutureToBankByBank]",
		slog.Any("data", pRspRepeal),
	)
}

func (spi *TraderLogSpi) OnRtnFromBankToFutureByFuture(pRspTransfer *CThostFtdcRspTransferField) {
	spi.Log(
		context.Background(), slog.LevelDebug-1,
		"thost trader [OnRtnFromBankToFutureByFuture]",
		slog.Any("data", pRspTransfer),
	)
}

func (spi *TraderLogSpi) OnRtnFromFutureToBankByFuture(pRspTransfer *CThostFtdcRspTransferField) {
	spi.Log(
		context.Background(), slog.LevelDebug-1,
		"thost trader [OnRtnFromFutureToBankByFuture]",
		slog.Any("data", pRspTransfer),
	)
}

func (spi *TraderLogSpi) OnRtnRepealFromBankToFutureByFutureManual(pRspRepeal *CThostFtdcRspRepealField) {
	spi.Log(
		context.Background(), slog.LevelDebug-1,
		"thost trader [OnRtnRepealFromBankToFutureByFutureManual]",
		slog.Any("data", pRspRepeal),
	)
}

func (spi *TraderLogSpi) OnRtnRepealFromFutureToBankByFutureManual(pRspRepeal *CThostFtdcRspRepealField) {
	spi.Log(
		context.Background(), slog.LevelDebug-1,
		"thost trader [OnRtnRepealFromFutureToBankByFutureManual]",
		slog.Any("data", pRspRepeal),
	)
}

func (spi *TraderLogSpi) OnRtnQueryBankBalanceByFuture(pNotifyQueryAccount *CThostFtdcNotifyQueryAccountField) {
	spi.Log(
		context.Background(), slog.LevelDebug-1,
		"thost trader [OnRtnQueryBankBalanceByFuture]",
		slog.Any("data", pNotifyQueryAccount),
	)
}

func (spi *TraderLogSpi) OnErrRtnBankToFutureByFuture(
	pReqTransfer *CThostFtdcReqTransferField, pRspInfo *CThostFtdcRspInfoField,
) {
	spi.Error(
		"thost trader [OnRtnBankToFutureByFuture] error",
		slog.Any("error", spi.CheckRsp(pRspInfo)),
		slog.Any("data", pReqTransfer),
	)
}

func (spi *TraderLogSpi) OnErrRtnFutureToBankByFuture(
	pReqTransfer *CThostFtdcReqTransferField, pRspInfo *CThostFtdcRspInfoField,
) {
	spi.Error(
		"thost trader [OnRtnFutureToBankByFuture] error",
		slog.Any("error", spi.CheckRsp(pRspInfo)),
		slog.Any("data", pReqTransfer),
	)
}

func (spi *TraderLogSpi) OnErrRtnRepealBankToFutureByFutureManual(
	pReqRepeal *CThostFtdcReqRepealField, pRspInfo *CThostFtdcRspInfoField,
) {
	spi.Error(
		"thost trader [OnRtnRepealBankToFutureByFutureManual] error",
		slog.Any("error", spi.CheckRsp(pRspInfo)),
		slog.Any("data", pReqRepeal),
	)
}

func (spi *TraderLogSpi) OnErrRtnRepealFutureToBankByFutureManual(
	pReqRepeal *CThostFtdcReqRepealField, pRspInfo *CThostFtdcRspInfoField,
) {
	spi.Error(
		"thost trader [OnRtnRepealFutureToBankByFutureManual] error",
		slog.Any("error", spi.CheckRsp(pRspInfo)),
		slog.Any("data", pReqRepeal),
	)
}

func (spi *TraderLogSpi) OnErrRtnQueryBankBalanceByFuture(
	pReqQueryAccount *CThostFtdcReqQueryAccountField, pRspInfo *CThostFtdcRspInfoField,
) {
	spi.Error(
		"thost trader [OnRtnQueryBankBalanceByFuture] error",
		slog.Any("error", spi.CheckRsp(pRspInfo)),
		slog.Any("data", pReqQueryAccount),
	)
}

func (spi *TraderLogSpi) OnRtnRepealFromBankToFutureByFuture(pRspRepeal *CThostFtdcRspRepealField) {
	spi.Log(
		context.Background(), slog.LevelDebug-1,
		"thost trader [OnRtnRepealFromBankToFutureByFuture]",
		slog.Any("data", pRspRepeal),
	)
}

func (spi *TraderLogSpi) OnRtnRepealFromFutureToBankByFuture(pRspRepeal *CThostFtdcRspRepealField) {
	spi.Log(
		context.Background(), slog.LevelDebug-1,
		"thost trader [OnRtnRepealFromFutureToBankByFuture]",
		slog.Any("data", pRspRepeal),
	)
}

func (spi *TraderLogSpi) OnRspFromBankToFutureByFuture(
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

func (spi *TraderLogSpi) OnRspFromFutureToBankByFuture(
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

func (spi *TraderLogSpi) OnRspQueryBankAccountMoneyByFuture(
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

func (spi *TraderLogSpi) OnRtnOpenAccountByBank(pOpenAccount *CThostFtdcOpenAccountField) {
	spi.Log(
		context.Background(), slog.LevelDebug-1,
		"thost trader [OnRtnOpenAccountByBank]",
		slog.Any("data", pOpenAccount),
	)
}

func (spi *TraderLogSpi) OnRtnCancelAccountByBank(pCancelAccount *CThostFtdcCancelAccountField) {
	spi.Log(
		context.Background(), slog.LevelDebug-1,
		"thost trader [OnRtnCancelAccountByBank]",
		slog.Any("data", pCancelAccount),
	)
}

func (spi *TraderLogSpi) OnRtnChangeAccountByBank(pChangeAccount *CThostFtdcChangeAccountField) {
	spi.Log(
		context.Background(), slog.LevelDebug-1,
		"thost trader [OnRtnChangeAccountByBank]",
		slog.Any("data", pChangeAccount),
	)
}

func (spi *TraderLogSpi) OnRspQryClassifiedInstrument(
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

func (spi *TraderLogSpi) OnRspQryCombPromotionParam(
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

func (spi *TraderLogSpi) OnRspQryRiskSettleInvstPosition(
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

func (spi *TraderLogSpi) OnRspQryRiskSettleProductStatus(
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

func (spi *TraderLogSpi) OnRspQrySPBMFutureParameter(
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

func (spi *TraderLogSpi) OnRspQrySPBMOptionParameter(
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

func (spi *TraderLogSpi) OnRspQrySPBMIntraParameter(
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

func (spi *TraderLogSpi) OnRspQrySPBMInterParameter(
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

func (spi *TraderLogSpi) OnRspQrySPBMPortfDefinition(
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

func (spi *TraderLogSpi) OnRspQrySPBMInvestorPortfDef(
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

func (spi *TraderLogSpi) OnRspQryInvestorPortfMarginRatio(
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

func (spi *TraderLogSpi) OnRspQryInvestorProdSPBMDetail(
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

func (spi *TraderLogSpi) OnRspQryInvestorCommoditySPMMMargin(
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

func (spi *TraderLogSpi) OnRspQryInvestorCommodityGroupSPMMMargin(
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

func (spi *TraderLogSpi) OnRspQrySPMMInstParam(
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

func (spi *TraderLogSpi) OnRspQrySPMMProductParam(
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

func (spi *TraderLogSpi) OnRspQrySPBMAddOnInterParameter(
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

func (spi *TraderLogSpi) OnRspQryRCAMSCombProductInfo(
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

func (spi *TraderLogSpi) OnRspQryRCAMSInstrParameter(
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

func (spi *TraderLogSpi) OnRspQryRCAMSIntraParameter(
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

func (spi *TraderLogSpi) OnRspQryRCAMSInterParameter(
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

func (spi *TraderLogSpi) OnRspQryRCAMSShortOptAdjustParam(
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

func (spi *TraderLogSpi) OnRspQryRCAMSInvestorCombPosition(
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

func (spi *TraderLogSpi) OnRspQryInvestorProdRCAMSMargin(
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

func (spi *TraderLogSpi) OnRspQryRULEInstrParameter(
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

func (spi *TraderLogSpi) OnRspQryRULEIntraParameter(
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

func (spi *TraderLogSpi) OnRspQryRULEInterParameter(
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

func (spi *TraderLogSpi) OnRspQryInvestorProdRULEMargin(
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

func (spi *TraderLogSpi) OnRspQryInvestorPortfSetting(
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

func (spi *TraderLogSpi) OnRspQryInvestorInfoCommRec(
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

func (spi *TraderLogSpi) OnRspQryCombLeg(
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

func (spi *TraderLogSpi) OnRspOffsetSetting(
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

func (spi *TraderLogSpi) OnRspCancelOffsetSetting(
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

func (spi *TraderLogSpi) OnRtnOffsetSetting(pOffsetSetting *CThostFtdcOffsetSettingField) {
	spi.Log(
		context.Background(), slog.LevelDebug-1,
		"thost trader [OnRtnOffsetSetting]",
		slog.Any("data", pOffsetSetting),
	)
}

func (spi *TraderLogSpi) OnErrRtnOffsetSetting(
	pInputOffsetSetting *CThostFtdcInputOffsetSettingField, pRspInfo *CThostFtdcRspInfoField,
) {
	spi.Error(
		"thost trader [OnRtnOffsetSetting] error",
		slog.Any("error", spi.CheckRsp(pRspInfo)),
		slog.Any("data", pInputOffsetSetting),
	)
}

func (spi *TraderLogSpi) OnErrRtnCancelOffsetSetting(
	pCancelOffsetSetting *CThostFtdcCancelOffsetSettingField, pRspInfo *CThostFtdcRspInfoField,
) {
	spi.Error(
		"thost trader [OnRtnCancelOffsetSetting] error",
		slog.Any("error", spi.CheckRsp(pRspInfo)),
		slog.Any("data", pCancelOffsetSetting),
	)
}

func (spi *TraderLogSpi) OnRspQryOffsetSetting(
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

func (spi *TraderLogSpi) OnRspGenSMSCode(
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

func (spi *TraderLogSpi) OnRtnSMSVerifyInfoFromSec(pSMSVerifyInfoFromSec *CThostFtdcSMSVerifyInfoFromSecField) {
	spi.Log(
		context.Background(), slog.LevelDebug-1,
		"thost trader [OnRtnSMSVerifyInfoFromSec]",
		slog.Any("data", pSMSVerifyInfoFromSec),
	)
}

func (spi *TraderLogSpi) OnRspSpdApply(
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

func (spi *TraderLogSpi) OnRspSpdApplyAction(
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

func (spi *TraderLogSpi) OnRspQrySpdApply(
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

	spi.Info(
		"thost trader [OnRspQrySpdApply] succeeded",
		slog.Any("data", pSpdApply),
	)
}

func (spi *TraderLogSpi) OnRtnSpdApply(pSpdApply *CThostFtdcSpdApplyField) {
	spi.Log(
		context.Background(), slog.LevelDebug-1,
		"thost trader [OnRtnSpdApply]",
		slog.Any("data", pSpdApply),
	)
}

func (spi *TraderLogSpi) OnErrRtnSpdApply(
	pInputSpdApply *CThostFtdcInputSpdApplyField, pRspInfo *CThostFtdcRspInfoField,
) {
	spi.Error(
		"thost trader [OnRtnSpdApply] error",
		slog.Any("error", spi.CheckRsp(pRspInfo)),
		slog.Any("data", pInputSpdApply),
	)
}

func (spi *TraderLogSpi) OnErrRtnSpdApplyAction(
	pSpdApplyAction *CThostFtdcSpdApplyActionField, pRspInfo *CThostFtdcRspInfoField,
) {
	spi.Error(
		"thost trader [OnRtnSpdApplyAction] error",
		slog.Any("error", spi.CheckRsp(pRspInfo)),
		slog.Any("data", pSpdApplyAction),
	)
}

func (spi *TraderLogSpi) OnRspHedgeCfm(
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

func (spi *TraderLogSpi) OnRspHedgeCfmAction(
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

func (spi *TraderLogSpi) OnRspQryHedgeCfm(
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

	spi.Info(
		"thost trader [OnRspQryHedgeCfm] succeeded",
		slog.Any("data", pHedgeCfm),
	)
}

func (spi *TraderLogSpi) OnRtnHedgeCfm(pHedgeCfm *CThostFtdcHedgeCfmField) {
	spi.Log(
		context.Background(), slog.LevelDebug-1,
		"thost trader [OnRtnHedgeCfm]",
		slog.Any("data", pHedgeCfm),
	)
}

func (spi *TraderLogSpi) OnErrRtnHedgeCfm(
	pInputHedgeCfm *CThostFtdcInputHedgeCfmField, pRspInfo *CThostFtdcRspInfoField,
) {
	spi.Error(
		"thost trader [OnRtnHedgeCfm] error",
		slog.Any("error", spi.CheckRsp(pRspInfo)),
		slog.Any("data", pInputHedgeCfm),
	)
}

func (spi *TraderLogSpi) OnErrRtnHedgeCfmAction(
	pHedgeCfmAction *CThostFtdcHedgeCfmActionField, pRspInfo *CThostFtdcRspInfoField,
) {
	spi.Error(
		"thost trader [OnRtnHedgeCfmAction] error",
		slog.Any("error", spi.CheckRsp(pRspInfo)),
		slog.Any("data", pHedgeCfmAction),
	)
}
