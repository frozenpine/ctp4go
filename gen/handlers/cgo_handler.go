package handlers

import (
	"bytes"
	"fmt"

	"github.com/frozenpine/ctp4go/gen/parser"
)

func CgoCaller(p *parser.Param) string {
	if p == nil {
		return ""
	}

	buff := bytes.NewBufferString("")

	if p.Name != "" {
		fmt.Fprintf(buff, "%s ", p.Name)
	}

	switch p.Type {
	case "Int":
		buff.WriteString("C.int")
	case "Bool":
		buff.WriteString("C.bool")
	case "Char_S":
		if p.IsPointer {
			buff.WriteString("*")
		}

		if p.IsArray {
			buff.WriteString("*")
		}

		buff.WriteString("C.char")
	default:
		if p.IsPointer {
			buff.WriteString("*")
		}

		fmt.Fprintf(buff, "C.struct_%s", p.Type)
	}

	return buff.String()
}

func CgoCallee(p *parser.Param) string {
	if p == nil {
		return ""
	}

	buff := bytes.NewBufferString("")

	switch p.Type {
	case "Int", "Enum":
		fmt.Fprintf(buff, "C.int(%s)", GoParamName(p))
	case "Bool":
		fmt.Fprintf(buff, "C.bool(%s)", GoParamName(p))
	case "Char_S":
		if p.IsPointer {
			// 需要在模板中预处理
			if p.IsArray {
				fmt.Fprintf(
					buff, "(**C.char)(unsafe.Pointer(&%s[0])), C.int(len(%s))",
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
