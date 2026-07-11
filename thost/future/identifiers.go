package future

import (
	"fmt"
	"strconv"
	"strings"

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
	front types.TThostFtdcFrontIDType, session types.TThostFtdcSessionIDType,
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

func PositionIdt(
	ex *types.Exchengeid
)

func DotIdt(inputs ...fmt.Stringer) string {
	values := make([]string, len(inputs))

	for idx, o := range inputs {
		values[idx] = o.String()
	}

	return strings.Join(values, ".")
}

func FormatIdt(format string, others ...any) string {
	return fmt.Sprintf(format, others...)
}
