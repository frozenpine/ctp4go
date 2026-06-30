package handlers

import (
	"bytes"
	"fmt"

	"github.com/frozenpine/ctp4go/gen/internal"
)

func CParamType(p *internal.Param) string {
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

	if p.Name != "" {
		buff.WriteString(" ")
	}

	if p.IsPointer {
		buff.WriteString("*")
	}
	if p.Name != "" {
		buff.WriteString(p.Name)
	}
	if p.IsArray {
		fmt.Fprintf(buff, "[], int %s", p.ArrSizeName)
	}

	return buff.String()
}

func CParamName(p *internal.Param) string {
	if !p.IsArray {
		return p.Name
	}

	return fmt.Sprintf("%s, %s", p.Name, p.ArrSizeName)
}
