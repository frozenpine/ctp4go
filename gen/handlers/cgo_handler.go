package handlers

import (
	"bytes"
	"fmt"

	"github.com/frozenpine/ctp4go/gen/parser"
)

func CgoParamName(p *parser.Param) string {
	buff := bytes.NewBufferString("")

	switch p.Type {
	case "Int":
		fmt.Fprintf(buff, "C.int(%s)", GoParamName(p))
	case "Bool":
		fmt.Fprintf(buff, "C.bool(%s)", GoParamName(p))
	case "Char_S":
		if p.IsPointer {
			// 需要在模板中预处理
			if p.IsArray {
				fmt.Fprintf(
					buff, "(**C.char)(%s), C.int(cnt%s)",
					p.Name, GoParamName(p),
				)
			} else {
				buff.WriteString(p.Name)
			}
		} else {
			fmt.Fprintf(buff, "C.char(%s)", GoParamName(p))
		}
	default:
		if p.IsStruct && p.IsPointer {
			fmt.Fprintf(
				buff, "(*C.struct_%s)(unsafe.Pointer(%s))",
				p.Type, GoParamName(p),
			)
		} else {
			fmt.Fprintf(buff, "C.%s(%s)", p.Type, GoParamName(p))
		}
	}

	// 数组指针
	// if p.IsArray {
	// 	fmt.Fprintf(buff, ", C.int %s", p.ArrSizeName)
	// }

	return buff.String()
}
