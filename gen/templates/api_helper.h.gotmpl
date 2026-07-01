{{- $sdk := .Sdk -}}
{{- $className := .ApiClass.Name -}}
#pragma once
#ifndef {{ $sdk.Name | ToUpper }}_{{ $sdk.Version | ReplaceAll "." "" | ToUpper }}_API_HELPER_H
#define {{ $sdk.Name | ToUpper }}_{{ $sdk.Version | ReplaceAll "." "" | ToUpper }}_API_HELPER_H

#ifdef __cplusplus
extern "C"
{
#endif

#include <stdbool.h>
#include <stdint.h>
#include <stdlib.h>
#include <stdio.h>
#include <string.h>
#include <dlfcn.h>

#include "ThostFtdcUserApiStruct.h"
// converted from {{ .EntryFile }}
// function call ptr
{{- range .ApiClass.Methods }}
typedef {{ if .Rtn }}{{ CCaller .Rtn }}{{ else }}void{{ end }} (*{{ .Name }})(void *this{{ range .Params }}, {{ if eq .Type $sdk.SpiName }}void *{{ .Name }}{{ else }}{{ CCaller . }}{{ end }}{{ end }});
{{ end }}

// api methods all vtable
typedef struct
{
{{- range .ApiClass.Methods }}
    {{ .Name }} {{ $className }}VTable_{{ .Name }};
{{ end }}
} {{ $className }}VTable;

// cpp api instance C wrapper
typedef struct
{
    {{ $className }}VTable *vtable;
} {{ $sdk.ApiExtName }};

// static functions to create cpp api instance & get version
{{ range .ApiClass.Statics }}
typedef {{ if .Rtn }}{{ if eq .Rtn.Type $className }}void*{{ else }}{{ CCaller .Rtn }}{{ end }}{{ else }}void{{ end }} (*{{ .Name }})({{ range $idx, $ele := .Params }}{{ if ne $idx 0}}, {{ end }}{{ CCaller $ele }}{{ end }});
{{ if .Rtn }}{{ if eq .Rtn.Type $className }}void*{{ else }}{{ CCaller .Rtn }}{{ end }}{{ else }}void{{ end }} Call{{ .Name }}({{ .Name }} fn{{ range .Params }}, {{ CCaller . }}{{ end }});
{{ end }}

// helper function call for cgo
{{- range .ApiClass.Methods }}
{{ if .Rtn }}{{ CCaller .Rtn }}{{ else }}void{{ end }} Call{{ .Name }}({{ .Name }} fn, void *this{{ range .Params }}, {{ if eq .Type $sdk.SpiName }}void *{{ .Name }}{{ else }}{{ CCaller . }}{{ end }}{{ end }});
{{ end }}

#ifdef __cplusplus
}
#endif

#endif