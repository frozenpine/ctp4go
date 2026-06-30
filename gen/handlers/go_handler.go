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

func GoParamType(p *parser.Param) string {
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
