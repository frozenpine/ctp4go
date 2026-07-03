{{- $sdk := .Sdk -}}
{{- $className := .ApiClass.Name -}}
package {{ $sdk.Version | ReplaceAll "." "_"}}

/*
#cgo CFLAGS: -I. -I${SRCDIR} -I${SRCDIR}/../../dependencies/{{ .Platform }}/{{ $sdk.Version }}/
#cgo LDFLAGS: -ldl

#include "api_helper.h"
#include "spi_helper.h"
*/
import "C"
import (
	"errors"
	"fmt"
	"log/slog"
	"path/filepath"
	"unsafe"

	"github.com/frozenpine/ctp4go/thost"
	"github.com/frozenpine/ctp4go/thost/types"
)

// var (
//     // 确保Api封装完整实现了thost中的接口签名
//     _ thost.{{ $className | TrimPrefix "CThostFtdc" }} = &{{ $className | TrimPrefix "C" }}{}
// )

func Create{{ $className | TrimPrefix "C" }}(
    libPath string, {{ range .CreateCall.Params }}{{ GoCaller . }}, {{end}}
) (*{{ $className | TrimPrefix "C" }}, error) {
    if libPath == "" {
		return nil, fmt.Errorf(
			"%w: invalid create args", thost.ErrInvalidArgs,
		)
	}

	libPath, err := filepath.Abs(libPath)
	if err != nil {
		return nil, errors.Join(thost.ErrInvalidArgs, err)
	}

	libCPath := C.CString(libPath)
	defer C.free(unsafe.Pointer(libCPath))

    slog.Info(
        "try to load thost {{ $sdk.Name }} api lib",
        slog.String("lib_path", libPath),
    )

    lib := C.dlopen(libCPath, C.RTLD_LAZY)
	if lib == nil {
		msg := C.dlerror()

		return nil, fmt.Errorf(
			"%w: %s", thost.ErrLibOpenFailed,
			types.DecodeGBK(([]byte)(C.GoString(msg))),
		)
	}

    slog.Info(
        "thost {{ $sdk.Name }} api lib opened",
        slog.String("lib_path", libPath),
    )

    var (
        createFnName = C.CString({{ $sdk.Name | ToUpper }}_CREATE_FN_NAME)
        {{ if .VersionCall }}versionFnName = C.CString({{ $sdk.Name | ToUpper }}_VERSION_FN_NAME){{ end }}

        apiVer string
    )
    defer func() {
        C.free(unsafe.Pointer(createFnName))
        {{ if .VersionCall }}C.free(unsafe.Pointer(versionFnName)){{ end }}
    }()

    creator := C.dlsym(lib, createFnName)
	if creator == nil {
		msg := C.dlerror()

		return nil, fmt.Errorf(
			"%w: %s", thost.ErrLibSymbolNotFound,
			types.DecodeGBK(([]byte)(C.GoString(msg))),
		)
	}

    {{ if .VersionCall -}}
    if versioner := C.dlsym(lib, versionFnName); versioner == nil {
		msg := C.dlerror()

		slog.Error(
			"thost mduser api version fn not found",
			slog.String("error", types.DecodeGBK(([]byte)(C.GoString(msg)))),
		)
	} else {
		apiVer = C.GoString(C.CallGetApiVersion(
			(C.GetApiVersion)(versioner),
		))
	}
    {{- end }}

    slog.Info(
		"thost trader api creator found, try to create api instance",
	)

    {{ range .CreateCall.Params -}}
		{{ if and (eq .Type "Char_S") .IsPointer }}
	{{ .Name }} := C.CString({{ GoParamName . }})
	defer C.free(unsafe.Pointer({{ .Name }}))
		{{ end }}
	{{- end }}

    instance := C.Call{{ .CreateCall.Name }}(
		C.{{ .CreateCall.Name }}(creator), 
		{{ range .CreateCall.Params }}{{ CgoCallee . }},{{ end }}
	)

    if instance == nil {
		return nil, fmt.Errorf(
			"%w: thost {{ $sdk.Name }} api[%s] create failed",
			thost.ErrApiCreateFailed, libPath,
		)
	}

    return &{{ $className | TrimPrefix "C" }}{
        apiPtr:  (*C.{{ $sdk.ApiExtName }})(instance),
        version: apiVer,
        lib:     lib,
    }, nil
}

type {{ $className | TrimPrefix "C" }} struct {
    lib unsafe.Pointer

    version     string
    apiPtr      *C.{{ $sdk.ApiExtName }}
    spiPtr      *{{ $sdk.SpiName | TrimPrefix "C" }}
}

func (api *{{ $className | TrimPrefix "C" }}) GetApiVersion() string {
    return api.version
}

{{ range .ApiClass.Methods }}
func (api *{{ $className | TrimPrefix "C" }}) {{ .Name }}({{ range .Params }}{{ if eq .Type $sdk.SpiName }}{{ GoParamName . }} thost.{{ $sdk.SpiName | TrimPrefix "CThostFtdc" }}{{ else }}{{ GoCaller . }}{{end}},{{ end }}) {{ GoCaller .Rtn }} {
	slog.Info("executing thost {{ $sdk.Name }} api {{ .Name }}")

	{{ if Contains .Name "Release" -}}defer func() {
		if api.spiPtr != nil {
			api.spiPtr.Pinner.Unpin()
			api.spiPtr = nil
		}

		C.dlclose(api.lib)
	}(){{- end }}

	{{- range .Params }}
		{{ if and (eq .Type "Char_S") .IsPointer }}
			{{ if .IsArray }}
	{{ .Name }} := make([]*C.char, len({{ GoParamName . }}))
	for idx, v := range {{ GoParamName . }} {
		{{ .Name }}[idx] = C.CString(v)
		defer C.free(unsafe.Pointer({{ .Name }}[idx]))
	}
			{{ else }}
	{{ .Name }} := C.CString({{ GoParamName . }})
	defer C.free(unsafe.Pointer({{ .Name }}))
			{{ end }}
		{{ end }}
		{{ if eq .Type $sdk.SpiName }}
	api.spiPtr = new({{ $sdk.SpiName | TrimPrefix "C" }})
	api.spiPtr.callback = {{ GoParamName . }}

	api.spiPtr.Pinner.Pin(unsafe.Pointer(api.spiPtr))
	slog.Info(
		"thost $sdk.Name spi created & pinned",
		slog.Any("spi", {{ GoParamName . }}),
		slog.Any("pinned", api.spiPtr),
	)

	{{ .Name }} := (*C.{{ $sdk.SpiExtName }})(C.malloc(
		C.sizeof_{{ $sdk.SpiExtName }},
	))
	{{ .Name }}.vtable = spiCVtablePtr
	{{ .Name }}.spi = unsafe.Pointer(api.spiPtr)
		{{ end }}
	{{ end }}

	{{ if .Rtn }} rtn := {{ end }}C.Call{{ .Name }}(
		api.apiPtr.vtable.{{ $className }}VTable_{{ .Name }},
		unsafe.Pointer(api.apiPtr),
		{{ range .Params }}{{ if eq .Type $sdk.SpiName }}unsafe.Pointer({{ .Name }}){{ else }}{{ CgoCallee . }}{{ end }},{{ end }}
	)

	slog.Info("thost {{ $sdk.Name }} api {{ .Name }} executed")

	{{ if .Rtn }}return {{ GoCallee .Rtn }}{{ end }}
}
{{ end }}
