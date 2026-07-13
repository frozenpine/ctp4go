package future

import (
	"strconv"
	"strings"
	"time"

	"github.com/frozenpine/ctp4go"
	"github.com/frozenpine/ctp4go/thost/future/types"
)

func InvestorIdt(
	brk *types.TThostFtdcBrokerIDType,
	inv *types.TThostFtdcInvestorIDType,
) string {
	buff := strings.Builder{}
	buff.Grow(len(brk) + len(inv))

	buff.WriteString(brk.String())
	buff.WriteByte('.')
	buff.WriteString(inv.String())

	return buff.String()
}

func OrderRefIdt(
	ref *types.TThostFtdcOrderRefType,
	front types.TThostFtdcFrontIDType,
	session types.TThostFtdcSessionIDType,
) string {
	buff := strings.Builder{}
	buff.Grow(len(ref) + 8 + 8)

	buff.WriteString(ref.String())
	buff.WriteString("@[")
	buff.WriteString(strconv.Itoa(int(front)))
	buff.WriteString("]:")
	buff.WriteString(strconv.Itoa(int(session)))

	return buff.String()
}

func SymbolIdt(
	ex *types.TThostFtdcExchangeIDType,
	ins *types.TThostFtdcInstrumentIDType,
) string {
	buff := strings.Builder{}
	buff.Grow(len(ex) + len(ins))

	buff.WriteString(ex.String())
	buff.WriteByte('.')
	buff.WriteString(ins.String())

	return buff.String()
}

func PositionIdt(
	ex *types.TThostFtdcExchangeIDType,
	ins *types.TThostFtdcInstrumentIDType,
	direct types.TThostFtdcPosiDirectionType,
	offset types.TThostFtdcHedgeFlagType,
) string {
	buff := strings.Builder{}
	buff.Grow(len(ex) + len(ins) + 10)

	buff.WriteString(SymbolIdt(ex, ins))
	buff.WriteByte('@')
	buff.WriteString(direct.String())
	buff.WriteByte('[')
	buff.WriteString(offset.String())
	buff.WriteByte(']')

	return buff.String()
}

func TimestampIdt(
	dt *types.TThostFtdcDateType,
	tm *types.TThostFtdcTimeType, mill types.TThostFtdcMillisecType,
) string {
	buff := strings.Builder{}
	buff.Grow(len(dt) + len(tm) + 3)

	buff.WriteString(dt.String())
	buff.WriteByte(' ')
	buff.WriteString(tm.String())
	buff.WriteByte('.')
	v := strconv.Itoa(int(mill))
	if len(v) < 3 {
		buff.WriteString(strings.Repeat("0", 3-len(v)))
	}
	buff.WriteString(v)

	return buff.String()
}

func MarketDataIdt(
	ex *types.TThostFtdcExchangeIDType,
	ins *types.TThostFtdcInstrumentIDType,
	dt *types.TThostFtdcDateType,
	tm *types.TThostFtdcTimeType, mill types.TThostFtdcMillisecType,
) string {
	buff := strings.Builder{}
	buff.Grow(len(ex) + len(ins) + len(dt) + len(tm) + 3)

	buff.WriteString(SymbolIdt(ex, ins))
	buff.WriteByte('@')
	buff.WriteString(TimestampIdt(dt, tm, mill))

	return buff.String()
}

func TimestampTs(
	dt *types.TThostFtdcDateType,
	tm *types.TThostFtdcTimeType, mill types.TThostFtdcMillisecType,
) time.Time {
	buff := strings.Builder{}
	buff.Grow(len(dt) + len(tm))

	buff.WriteString(dt.String())
	buff.WriteString(tm.String())

	ts, _ := time.ParseInLocation("20060102150405", buff.String(), ctp4go.CST)
	ts.Add(time.Millisecond * time.Duration(mill))

	return ts
}
