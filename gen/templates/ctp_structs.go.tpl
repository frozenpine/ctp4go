package thost

import "github.com/frozenpine/ctp4go/thost/types"

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
    {{ .Name }} types.{{ .Type }}
    {{- end }}
}
{{ end }}