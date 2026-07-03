{{- $sdk := .Sdk -}}
//go:build windows

package {{ $sdk.Version | ReplaceAll "." "_"}}

const (
    {{ $sdk.Name | ToUpper }}_CREATE_FN_NAME = "{{ .CreateCall.WinMangledName }}"
    {{ $sdk.Name | ToUpper }}_VERSION_FN_NAME = "{{ .VersionCall.WinMangledName }}"
)