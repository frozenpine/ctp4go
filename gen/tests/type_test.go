package tests

import (
	"os"
	"strings"
	"testing"
	"text/template"

	"github.com/frozenpine/ctp4go/gen/handlers"
	"github.com/frozenpine/ctp4go/gen/parser"
)

var tplFuncs = template.FuncMap{
	"ToUpper": strings.ToUpper,
	"ToLower": strings.ToLower,
	"Title":   strings.ToTitle,
	"Replace": func(old, new string, n int, in string) string {
		return strings.Replace(in, old, new, n)
	},
	"ReplaceAll": func(old, new, in string) string {
		return strings.ReplaceAll(in, old, new)
	},
	"TrimSpace": strings.TrimSpace,
	"TrimPrefix": func(prefix, in string) string {
		return strings.TrimPrefix(in, prefix)
	},
	"TrimSuffix": func(suffix, in string) string {
		return strings.TrimSuffix(in, suffix)
	},
	"Contains": func(substr, in string) bool {
		return strings.Contains(in, substr)
	},
	"Index": func(substr, in string) int {
		return strings.Index(in, substr)
	},

	"CCaller":     handlers.CCaller,
	"CCallee":     handlers.CCallee,
	"GoCaller":    handlers.GoCaller,
	"GoCallee":    handlers.GoCallee,
	"GoParamName": handlers.GoParamName,
	"GoType":      handlers.GoType,
	"CgoCaller":   handlers.CgoCaller,
	"CgoCallee":   handlers.CgoCallee,
}

func TestTypes(t *testing.T) {
	dep := "../../dependencies"
	typTpl := "../templates/ctp_types.go.tpl"

	entry, err := parser.NewEntry(
		"mini", dep, parser.WithSDK("",
			parser.WithVersion("v1.7.5"),
			parser.WithHdrFileName("ThostFtdcUserApiStruct.h"),
		),
	)
	if err != nil {
		t.Fatal(err)
	}
	defer entry.Release()

	if err := entry.Parse(); err != nil {
		t.Fatal(err)
	}

	content, err := os.ReadFile(typTpl)
	if err != nil {
		t.Fatal(err)
	}

	tpl, err := template.New("ctp_types").Funcs(tplFuncs).Parse(string(content))
	if err != nil {
		t.Fatal(err)
	}

	if err := tpl.Execute(os.Stdout, entry); err != nil {
		t.Fatal(err)
	}
}
