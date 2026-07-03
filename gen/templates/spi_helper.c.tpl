{{- $className := .SpiClass.Name -}}
#include "spi_helper.h"

{{- range .SpiClass.Methods }}
extern {{ if .Rtn }}{{ CCaller .Rtn }}{{ else }}void{{ end }} Cgo{{ .Name }}(void *this{{ range .Params }}, {{ CCaller . }}{{ end }});
{{ end }}

{{- range .SpiClass.Methods }}
{{ if .Rtn }}{{ CCaller .Rtn }}{{ else }}void{{ end }} C{{ .Name }}(void *this{{ range .Params }}, {{ CCaller . }}{{ end }})
{
    return Cgo{{ .Name }}(this{{ range .Params }}, {{ CCallee . }}{{ end }});
}
{{ end }}