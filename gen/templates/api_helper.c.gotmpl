{{- $sdk := .Sdk -}}
{{- $className := .ApiClass.Name -}}
#include "api_helper.h"
{{ range .ApiClass.Statics }}
{{ if .Rtn }}{{ if eq .Rtn.Type $className }}void*{{ else }}{{ CCaller .Rtn }}{{ end }}{{ else }}void{{ end }} Call{{ .Name }}({{ .Name }} fn{{ range .Params }}, {{ if eq .Type $sdk.SpiName }}void *{{ .Name }}{{ else }}{{ CCaller . }}{{ end }}{{ end }})
{
    return fn({{ range $idx, $ele := .Params }}{{ if ne $idx 0 }}, {{ end }}{{ CCallee $ele }}{{ end }});
}
{{ end }}

{{- range .ApiClass.Methods }}
{{ if .Rtn }}{{ CCaller .Rtn }}{{ else }}void{{ end }} Call{{ .Name }}({{ .Name }} fn, void *this{{ range .Params }}, {{ if eq .Type $sdk.SpiName }}void *{{ .Name }}{{ else }}{{ CCaller . }}{{ end }}{{ end }})
{
    return fn(this{{ range .Params }}, {{ CCallee . }}{{ end }});
}
{{ end }}