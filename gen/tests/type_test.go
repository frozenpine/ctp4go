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
	typTpl := "../templates/ctp_types.go.gotmpl"

	defer parser.CTPEntry.Release()

	if err := parser.CTPEntry.Parse(
		dep, parser.WithPlatform("future"), parser.WithSDK(
			"mduser", parser.WithVersion("v6.5.1"),
		),
	); err != nil {
		t.Fatal(err)
	}

	tpl, err := template.New("ctp_types").Funcs(tplFuncs).ParseFiles(typTpl)
	if err != nil {
		t.Fatal(err)
	}

	if err := tpl.Execute(os.Stdout, &parser.CTPEntry); err != nil {
		t.Fatal(err)
	}
}
