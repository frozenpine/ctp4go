{{- $sdk := .Sdk -}}
//go:build {{ $sdk.Version }}

import (
    _ "github.com/frozenpine/ctp4go/{{ .Platform }}/{{ $sdk.Version | ReplaceAll "." "_" }}"
)