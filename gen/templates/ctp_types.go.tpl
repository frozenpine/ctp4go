package types

import (
    "strconv"
    "math"

	"github.com/frozenpine/ctp4go"
)

{{ range $_, $enum := .Enums }}
	{{- range $enum.Comments.Summary }}
// {{ . }}
    {{- end }}
    {{- range .Comments.ParamComment }}
//   @{{ .Cmd }} {{ .ArgName }} {{ .Description }}
        {{- range .Values }}
//     {{ . }}
        {{- end }}
    {{- end }}
//go:generate stringer -type {{ $enum.Name }} -linecomment
type {{ .Name }} int32

const (
	{{- range .Members }}
	{{ .Name }} {{ $enum.Name }} = {{ .Value }} // {{ if gt (len .Comments.Summary) 0 -}}{{ index .Comments.Summary 0 }}{{ else }}{{ .Name | TrimPrefix "THOST_TERT_" }}{{ end }}
	{{- end }}
)
{{ end }}

{{ range $_, $def := .Types }}
    {{- range $def.Comments.Summary }}
// {{ . }}
    {{- end }}
    {{- range .Comments.ParamComment }}
//   @{{ .Cmd }} {{ .ArgName }} {{ .Description }}
        {{- range .Values }}
//     {{ . }}
        {{- end }}
    {{- end }}
type {{ .Name }} {{ GoType .Underlying }}
    {{ if .MacroDefine }}
//go:generate stringer -type {{ .Name }} -linecomment
const(
        {{- range .MacroDefine.Defines }}
    {{ .Name }} {{ $def.Name }} = {{ .Token }} // {{ index .Comments.Summary 0 }}
        {{- end }}
)
    {{ end }}
    {{- if eq .Underlying.Name "Double" }}
func (t {{ .Name }}) IsValid() bool {
    return float64(t) != math.MaxFloat64
}

func (t {{ .Name }}) String() string {
    if t.IsValid() {
        return strconv.FormatFloat(float64(t), 'f', 6, 64)
    } else {
        return "DBL_MAX"
    }
}
    {{ end }}
    {{- if and (eq .Underlying.Name "Char_S") (gt .Underlying.Size 0) }}
func (t {{ .Name }}) String() string { return ctp4go.DecodeGBK(t[:]) }

func (t *{{ .Name }}) SetString(v string) int {
    return ctp4go.SetCString(([]byte)((*t)[:]), v)
}
    {{ end }}
{{ end }}