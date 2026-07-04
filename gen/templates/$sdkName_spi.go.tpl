{{- $sdk := .Sdk -}}
package {{ .Platform }}

type {{ .SpiClass.Name | TrimPrefix "CThostFtdc" }} interface {
{{- range .SpiClass.Methods }}
    {{- range .Comments.Summary }}
    // {{ . }}
    {{- end}}
    {{- range .Comments.ParamComment }}
    //   @{{ .Cmd }} {{ .ArgName }} {{ .Description }}
        {{- range .Values}}
    //     {{ . }}
        {{- end }}
    {{- end }}
    {{ .Name }}({{ range $idx, $ele := .Params }}{{ if gt $idx 0 }}, {{ end }}{{ GoCaller . }}{{ end }})
{{ end }}
}
