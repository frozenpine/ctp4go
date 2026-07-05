package {{ .Platform }}

import "github.com/frozenpine/ctp4go/thost/{{ .Platform }}/types"

{{ range .Structures }}
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
{{ end }}