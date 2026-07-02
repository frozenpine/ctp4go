package handlers

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/frozenpine/ctp4go/gen/parser"
)

func GoParamName(p *parser.Param) string {
	if p != nil && p.Name != "" {
		return strings.TrimLeftFunc(p.Name, func(r rune) bool {
			return 'a' <= r && r <= 'z'
		})
	}

	return ""
}

func GoCallee(p *parser.Param) string {
	buff := bytes.NewBufferString("")

	switch p.Type {
	case "Int":
		if p.Name != "" {
			fmt.Fprintf(buff, "int(%s)", p.Name)
		} else {
			buff.WriteString("int(rtn)")
		}
	case "Bool":
		if p.Name != "" {
			fmt.Fprintf(buff, "bool(%s)", p.Name)
		} else {
			buff.WriteString("bool(rtn)")
		}
	case "Char_S":
		if p.Name != "" {
			fmt.Fprintf(buff, "C.GoString(%s)", p.Name)
		} else {
			buff.WriteString("C.GoString(rtn)")
		}
	default:
		if p.IsPointer {
			fmt.Fprintf(buff, "(*thost.%s)(unsafe.Pointer(%s))", p.Type, p.Name)
		}
	}

	return buff.String()
}

func GoCaller(p *parser.Param) string {
	if p == nil {
		return ""
	}

	buff := bytes.NewBufferString("")

	fmt.Fprintf(buff, "%s ", GoParamName(p))

	switch p.Type {
	case "Int":
		if p.IsPointer {
			buff.WriteString("*")
		}
		buff.WriteString("int")
	case "Bool":
		if p.IsPointer {
			buff.WriteString("*")
		}
		buff.WriteString("bool")
	case "Char_S":
		if p.IsPointer {
			if p.IsArray {
				buff.WriteString("...")
			}
			buff.WriteString("string")
		} else {
			buff.WriteString("byte")
		}
	case "Enum":
		buff.WriteString("int")
	default:
		if p.IsPointer {
			buff.WriteString("*")
		}

		if p.IsArray {
			buff.WriteString("[]")
		}

		if p.IsStruct {
			buff.WriteString("thost.")
		}
		buff.WriteString(p.Type)
	}

	return buff.String()
}

func GoType(p parser.UnderType) string {
	buff := bytes.NewBufferString("")

	if p.Size > 0 {
		fmt.Fprintf(buff, "[%d]", p.Size)
	}

	switch p.Name {
	case "Int":
		buff.WriteString("int32")
	case "Bool":
		buff.WriteString("bool")
	case "Char_S":
		buff.WriteString("byte")
	case "Double":
		buff.WriteString("float64")
	}

	return buff.String()
}
