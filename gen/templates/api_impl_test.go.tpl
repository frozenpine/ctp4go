{{- $sdk := .Sdk -}}
{{- $className := .ApiClass.Name -}}
package {{ $sdk.Version | ReplaceAll "." "_"}}_test

import (
	"fmt"
	"go/types"
	"os"
	"strings"
	"testing"

	"golang.org/x/tools/go/packages"
)

func TestGenerate(t *testing.T) {
	cfg := &packages.Config{
		Mode: packages.NeedName | packages.NeedTypes |
			packages.NeedSyntax | packages.NeedTypesInfo,
	}

	pkgs, err := packages.Load(cfg, "../../thost")
	if err != nil {
		t.Fatal(err)
	}

	thost := pkgs[0]
    apiMethods := map[string]*types.Signature{}

	if len(thost.Errors) > 0 {
		t.Fatal(thost.Errors)
	}

    generatedMethods := []string{
        {{- range .ApiClass.Methods }}
        "{{ .Name }}",
        {{ end }}
    }

	scope := thost.Types.Scope()
	for _, name := range scope.Names() {
		obj := scope.Lookup(name)

		typeName, ok := obj.(*types.TypeName)

		if !ok {
			continue
		}

		iface, ok := typeName.Type().Underlying().(*types.Interface)
		if !ok {
			continue
		} else if !strings.HasSuffix(typeName.Name(), "TraderApi") {
			continue
		}

		iface.Complete()
		for method := range iface.Methods() {
            if !method.Exported() {
				continue
			}

            switch method.Name() {
            {{- range .ApiClass.Statics }}
            case "{{ .Name }}":
                continue
            {{ end }}
            }

			sig := method.Signature()
			apiMethods[method.Name()] = sig
		}
	}

    for _, k := range generatedMethods {
        delete(apiMethods, k)
    }

    if len(apiMethods) > 0 {
        complete, err := os.OpenFile(
			"api_complete.go", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, os.ModePerm,
		)
		if err != nil {
			t.Fatal(err)
		}
        defer complete.Close()

		fmt.Fprintf(
            complete, 
            "package {{ $sdk.Version | ReplaceAll "." "_"}}\n\n",
        )

        fmt.Fprintf(
            complete, "import \"github.com/frozenpine/ctp4go/thost\"\n\n",
        )

        for n, sig := range apiMethods {
            variadic := sig.Variadic()
            params := sig.Params()
            results := sig.Results()

            fmt.Fprintf(
                complete, 
                "func (api *{{ $className | TrimPrefix "C" }}) %s(\n", n,
            )

            for idx := 0; idx < params.Len(); idx++ {
                p := params.At(idx)
                var prefix string
                if variadic && idx == params.Len()-1 {
                    prefix = "..."
                }

                fmt.Fprintf(
                    complete, "\t%s%s,\n",
                    prefix, 
                    strings.ReplaceAll(
                        p.Type().String(), 
                        "github.com/frozenpine/ctp4go/", "",
                    ),
                )
            }

            switch results.Len() {
            case 0:
                fmt.Fprintf(complete, ") {}\n\n")
            case 1:
                rtn := results.At(0)
                rtnType := strings.ReplaceAll(
                    rtn.Type().String(),
                    "github.com/frozenpine/ctp4go/", "",
                )
                
                fmt.Fprintf(
                    complete, ") %s {\n\tvar r %s\n\treturn r\n}\n\n", 
                    rtnType, rtnType,
                )
            default:
                var (
                    rTypes []string
                    rNames []string
                    rDefines []string
                )
                
                for idx := 0; idx < results.Len(); idx++ {
                    rtn := results.At(idx)
                    rtnType := strings.ReplaceAll(
                        rtn.Type().String(),
                        "github.com/frozenpine/ctp4go/", "",
                    )

                    rTypes = append(rTypes, rtnType)
                    rNames = append(rNames, fmt.Sprintf("n%d", idx))
                    rDefines = append(
                        rDefines, fmt.Sprintf("n%d %s", idx, rtnType),
                    )
                }

                fmt.Fprintf(
                    complete, "(%s) {\n\tvar (\n%s\t)\n\t return %s}\n\n",
                    strings.Join(rTypes, ", "),
                    strings.Join(rDefines, "\n"),
                    strings.Join(rNames, ", "),
                )
            }
        }
    }
}

