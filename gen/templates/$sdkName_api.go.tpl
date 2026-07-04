{{- $sdk := .Sdk -}}
package {{ .Platform }}

type {{ .ApiClass.Name | TrimPrefix "CThostFtdc" }} interface {
{{- range .ApiClass.Methods }}
    {{- range .Comments.Summary }}
    // {{ . }}
    {{- end}}
    {{- range .Comments.ParamComment }}
    //   @{{ .Cmd }} {{ .ArgName }} {{ .Description }}
        {{- range .Values}}
    //     {{ . }}
        {{- end }}
    {{- end }}
    {{ .Name }}({{ range $idx, $ele := .Params }}
        {{- if gt $idx 0 }}, {{ end }}
        {{- if eq $ele.Type $sdk.SpiName }}
            {{- GoParamName . }} {{ $sdk.SpiName | TrimPrefix "CThostFtdc" }}
        {{- else }}
            {{- GoCaller . }}
        {{- end }}
    {{- end }}){{ if .Rtn}} {{ GoParamType .Rtn }}{{end}}
{{ end }}
}
