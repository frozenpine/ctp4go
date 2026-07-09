package {{ .Platform }}

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/frozenpine/ctp4go"
)

var (
{{- range .Entries }}
    _ {{ .SpiClass.Name | TrimPrefix "CThostFtdc" }} = &{{ .Platform | Title }}LogSpi{}
{{- end }}
)

type {{ .Platform | Title }}LogSpi struct {
    *slog.Logger
}

func (spi *{{ $.Platform | Title }}LogSpi) CheckRsp(rsp *CThostFtdcRspInfoField) error {
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

{{- range $_, $sdk := .Entries }}
    {{- range $sdk.SpiClass.Methods }}
func (spi *{{ $.Platform | Title }}LogSpi) {{ .Name }}({{ range $idx, $ele := .Params }}{{ if gt $idx 0 }}, {{ end }}{{ GoCaller . }}{{ end }}) {
    {{- if HasPrefix "OnRsp" .Name }}
    err := spi.CheckRsp(RspInfo)
    if err != nil {
        spi.Error(
            "thost {{ $sdk.Name }} [{{ .Name }}] failed",
            slog.Any("error", err),
            slog.Int("seq", RequestID),
            slog.Bool("last", IsLast),
        )
        return
    }
    {{ end }}
    {{- if HasPrefix "OnErrRtn" .Name }}
    spi.Error(
        "thost {{ $sdk.Name }} [{{ .Name }}] error",
        slog.Any("error", spi.CheckRsp(RspInfo)),
        slog.Any("data", {{ index .Params 0 | GoParamName }})
    )
    {{- else }}
    spi.{{ if Contains "Warn" .Name }}Warn{{ else }}Info{{ end }}(
        "thos [{{ .Name }}] called",
        {{- range .Params }}
        slog.Any("{{ .Name }}", {{ GoParamName . }})
        {{- end }}
    )
    {{- end }}
}
    {{- end}}
{{ end }}