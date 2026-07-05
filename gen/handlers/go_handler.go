package handlers

import (
	"bytes"
	"fmt"
	"os"
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

func GoParamType(p *parser.Param, prefix ...string) string {
	buff := bytes.NewBufferString("")

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
	case "LongLong":
		buff.WriteString("int64")
	default:
		if p.IsPointer {
			buff.WriteString("*")
		}

		if p.IsArray {
			buff.WriteString("[]")
		}

		if p.IsStruct {
			buff.WriteString(strings.Join(
				append(prefix, p.Type), ".",
			))
		}
	}

	return buff.String()
}

func GoCallee(p *parser.Param, prefix ...string) string {
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
			fmt.Fprintf(buff, "(%s)(unsafe.Pointer(%s))",
				GoParamType(p, prefix...), p.Name)
		}
	}

	return buff.String()
}

func GoCaller(p *parser.Param, prefix ...string) string {
	if p == nil {
		return ""
	}

	return fmt.Sprintf("%s %s", GoParamName(p), GoParamType(p, prefix...))
}

func GoTypeName(v string) string {
	switch v {
	case "Int":
		return "int32"
	case "Bool":
		return "bool"
	case "Char_S":
		return "byte"
	case "Double":
		return "float64"
	case "Short":
		return "int16"
	case "LongLong":
		return "int64"
	case "UInt":
		return "uint32"
	default:
		fmt.Fprintf(
			os.Stderr, "unsupported type for go: %s", v,
		)
	}

	return ""
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
		if p.Size > 0 {
			buff.WriteString("byte")
		} else {
			buff.WriteString("uint8")
		}
	case "Double":
		buff.WriteString("float64")
	case "Short":
		buff.WriteString("int16")
	case "LongLong":
		buff.WriteString("int64")
	case "Uint":
		buff.WriteString("uint32")
	default:
		fmt.Fprintf(
			os.Stderr, "unsupported type for go: %s", p.String(),
		)
	}

	return buff.String()
}
