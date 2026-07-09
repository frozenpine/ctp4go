package tests

import (
	"os"
	"testing"
	"text/template"

	"github.com/frozenpine/ctp4go/gen/handlers"
	"github.com/frozenpine/ctp4go/gen/parser"
)

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

	tpl, err := template.New("ctp_types").Funcs(
		handlers.TplFuncs,
	).Parse(string(content))
	if err != nil {
		t.Fatal(err)
	}

	if err := tpl.Execute(os.Stdout, entry); err != nil {
		t.Fatal(err)
	}
}
