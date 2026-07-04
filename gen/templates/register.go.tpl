{{- $sdk := .Sdk -}}
{{- $className := .ApiClass.Name -}}
package {{ $sdk.Version | ReplaceAll "." "_"}}

import (
	"fmt"

	"github.com/frozenpine/ctp4go/thost"
)

func sdkMaker(
	libPath string,
	params ...thost.Param,
) func() (thost.{{ $className | TrimPrefix "CThostFtdc" }}, error) {
	return func() (thost.{{ $className | TrimPrefix "CThostFtdc" }}, error) {
		if libPath == "" {
			return nil, fmt.Errorf(
				"%w: lib path is empty", thost.ErrInvalidArgs,
			)
		}

		var (
			{{- range .CreateCall.Params }}
			{{ GoCaller . }}
			{{ end }}

			ok bool
		)

		for _, p := range params {
			switch p.Key {
			{{- range .CreateCall.Params }}
			case thost.Param{{ GoParamName . }}:
				if {{ GoParamName }}, ok = p.Value.({{ GoParamType . }}); !ok {
					return nil, fmt.Errorf(
						"%w: invalid %s value %+v",
						thost.ErrInvalidArgs, p.Key, p.Value,
					)
				}
			{{ end }}
			}
		}

		return Create{{ $className | TrimPrefix "C" }}(
			libPath, {{ range .CreateCall.Params }}{{ GoCaller . }}, {{end}}
		)
	}
}

func init() {
	if err := thost.Set{{ $sdk.Name | Title }}Maker("{{ $sdk.Version }}", sdkMaker); err != nil {
		panic(err)
	}
}
