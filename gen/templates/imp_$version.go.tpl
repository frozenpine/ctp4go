{{- $sdk := .Sdk -}}
//go:build {{ $sdk.Version }}
package {{ .Platform }}

import (
    _ "github.com/frozenpine/ctp4go/{{ $sdk.Name }}/{{ .Platform }}/{{ $sdk.Version }}"
)