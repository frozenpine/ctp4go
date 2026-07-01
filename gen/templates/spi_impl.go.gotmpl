{{- $sdk := .Sdk -}}
{{- $className := .SpiClass.Name -}}
package {{ $sdk.Version | ReplaceAll "." "_"}}

/*
#cgo CFLAGS: -I. -I${SRCDIR} -I${SRCDIR}/../../dependencies/{{ .Platform }}/{{ $sdk.Version }}/
#cgo LDFLAGS: -ldl

#include "spi_helper.h"
*/
import "C"
import (
	"context"
	"log/slog"
	"runtime"
	"unsafe"

	"github.com/frozenpine/ctp4go/thost"
)

var (
	// 全局回调函数虚表
	spiCVtablePtr *C.{{ $className }}VTable
)

func init() {
	// C端为虚表分配内存
	spiCVtablePtr = (*C.{{ $className }}VTable)(C.malloc(
		C.sizeof_{{ $className }}VTable))

	{{ range .SpiClass.Methods }}
	spiCVtablePtr.{{ $className }}VTable_{{ .Name }} = (C.{{ .Name }})(
		unsafe.Pointer(C.C{{ .Name }}))
	{{ end }}
}

type {{ $className | TrimPrefix "C" }} struct {
    runtime.Pinner
	callback thost.{{ $className | TrimPrefix "CThostFtdc" }}
}

{{ range .SpiClass.Methods }}
//export Cgo{{ .Name }}
func Cgo{{ .Name }}(
	this unsafe.Pointer, 
	{{- range .Params }}
	{{ CgoCaller . }},
	{{- end }}
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"Cgo{{ .Name }} called",
		slog.Any("this", this),
	)

	(*{{ $className | TrimPrefix "C" }})(
		(*C.{{ $sdk.SpiExtName }})(this).spi,
	).callback.{{ .Name }}(
		{{- range .Params }}
		{{ GoCallee . }},{{ end }}
	)
}
{{ end }}
