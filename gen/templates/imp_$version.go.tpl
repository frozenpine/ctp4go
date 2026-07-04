{{- $sdk := .Sdk -}}
//go:build {{ $sdk.Version }}

package {{ $sdk.Name }}

import (
    _ "github.com/frozenpine/ctp4go/{{ .Platform }}/{{ $sdk.Version }}"
)