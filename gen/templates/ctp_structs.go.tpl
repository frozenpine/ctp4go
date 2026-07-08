package {{ .Platform }}

import (
    "strings"
    "fmt"

    "github.com/frozenpine/ctp4go/thost/{{ .Platform }}/types"
)

{{ range .Structures }}
    {{- $buffLen := len .Name }}
    {{- range .Comments.Summary }}
// {{ . }}
    {{- end }}
    {{- range .Comments.ParamComment }}
//   @{{ .Cmd }} {{ .ArgName }} {{ .Description }}
        {{- range .Values }}
//     {{ . }}
        {{- end }}
    {{- end }}
type {{ .Name }} struct {
    {{- range .Fields }}
        {{- $buffLen = Add $buffLen (len .Name) 10 }}
        {{- range .Comments.Summary }}
    // {{ . }}
        {{- end }}
        {{- range .Comments.ParamComment }}
    //   @{{ .Cmd }} {{ .ArgName }} {{ .Description }}
            {{- range .Values }}
    //     {{ . }}
            {{- end }}
        {{- end}}
    {{ .Name }} {{ if .IsBaseType }}{{ GoTypeName .Type }}{{ else }}types.{{ .Type }}{{ end }}
    {{- end }}
}

func (d {{ .Name }}) Type() string { return "{{ .Name }}" }

func (d {{ .Name }}) String() string {
    var builder strings.Builder
    builder.Grow({{ $buffLen }})

    builder.WriteString("{{ .Name }}{")
    
    {{- range $idx, $field := .Fields }}
        {{- if gt $idx 0 }}
    fmt.Fprintf(&builder, ", {{ $field.Name }}=%+v", d.{{ $field.Name }})
        {{- else }}
    fmt.Fprintf(&builder, "{{ $field.Name }}=%+v", d.{{ $field.Name }})
        {{- end }}
    {{- end}}
    
    builder.WriteByte('}')

    return builder.String()
}
{{ end }}