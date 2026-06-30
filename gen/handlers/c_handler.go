package handlers

import (
	"bytes"
	"fmt"

	"github.com/frozenpine/ctp4go/gen/parser"
)

func CParamType(p *parser.Param) string {
	buff := bytes.NewBufferString("")

	if p.IsConst {
		buff.WriteString("const ")
	}

	switch p.Type {
	case "Int":
		buff.WriteString("int")
	case "Bool":
		buff.WriteString("bool")
	case "Char_S":
		buff.WriteString("char")
	default:
		if p.IsStruct {
			buff.WriteString("struct ")
		}
		fmt.Fprintf(buff, "%s", p.Type)
	}

	if p.IsPointer {
		buff.WriteString("*")
	}

	if p.IsArray {
		fmt.Fprintf(buff, "[], int %s", p.ArrSizeName)
	}

	return buff.String()
}

func CParamName(p *parser.Param) string {
	if !p.IsArray {
		return p.Name
	}

	return fmt.Sprintf("%s, %s", p.Name, p.ArrSizeName)
}
