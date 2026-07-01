package main

import (
	"bytes"
	"embed"
	"errors"
	"flag"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"slices"
	"strings"
	"text/template"

	"github.com/frozenpine/ctp4go/gen/handlers"
	"github.com/frozenpine/ctp4go/gen/parser"
)

var (
	version, goVersion, gitVersion, buildTime string

	dep  string
	plat string
	sdk  = sdkOpt{
		name: "mduser",
		ver:  "v6.7.13",
	}
	debug  bool
	stdout bool
	output outputOpt

	cTplMapper = map[string][]string{
		"api": {
			"api_helper.h.gotmpl",
			"api_helper.c.gotmpl",
			"api_impl.go.gotmpl",
			// "api_impl_test.go.gotmpl",
			"consts_linux.go.gotmpl",
			"consts_windows.go.gotmpl",
		},
		"spi": {
			"spi_helper.h.gotmpl",
			"spi_helper.c.gotmpl",
			"spi_impl.go.gotmpl",
		},
	}

	tplFuncs = template.FuncMap{
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
		"CgoCaller":   handlers.CgoCaller,
		"CgoCallee":   handlers.CgoCallee,
	}
)

//go:embed templates
var tplFs embed.FS

type sdkOpt struct {
	name        string
	ver         string
	hdrFileName string
	apiName     string
	apiExtName  string
	spiName     string
	spiExtName  string
}

func (sdk sdkOpt) Type() string { return "SDK" }

func (sdk sdkOpt) String() string {
	buff := bytes.NewBufferString("SdkOpt{")

	fmt.Fprintf(buff, "Name=%s", sdk.name)
	fmt.Fprintf(buff, " Version=%s", sdk.ver)
	if sdk.hdrFileName != "" {
		fmt.Fprintf(buff, " HdrFile=%s", sdk.hdrFileName)
	}
	if sdk.apiName != "" {
		fmt.Fprintf(buff, " ApiName=%s", sdk.apiName)
	}
	if sdk.apiExtName != "" {
		fmt.Fprintf(buff, " ApiExtName=%s", sdk.apiExtName)
	}
	if sdk.spiName != "" {
		fmt.Fprintf(buff, " SpiName=%s", sdk.spiName)
	}
	if sdk.spiExtName != "" {
		fmt.Fprintf(buff, " SpiExtName=%s", sdk.spiExtName)
	}
	buff.WriteString("}")

	return buff.String()
}

func (sdk *sdkOpt) Set(v string) error {
	for d := range strings.SplitSeq(v, ",") {
		kvData := strings.Split(d, "=")

		if len(kvData) != 2 {
			return fmt.Errorf("invalid sdk option: %s", d)
		}

		switch strings.ToLower(strings.TrimSpace(kvData[0])) {
		case "name", "type":
			sdk.name = strings.TrimSpace(kvData[1])
		case "version", "ver":
			sdk.ver = strings.TrimSpace(kvData[1])
		case "hdr", "header":
			sdk.hdrFileName = strings.TrimSpace(kvData[1])
		case "api":
			if strings.Contains(kvData[1], ":") {
				apiValues := strings.SplitN(kvData[1], ":", 2)
				sdk.apiName = strings.TrimSpace(apiValues[0])
				sdk.apiExtName = strings.TrimSpace(apiValues[1])
			} else {
				sdk.apiName = strings.TrimSpace(kvData[1])
			}
		case "spi":
			if strings.Contains(kvData[1], ":") {
				spiValues := strings.SplitN(kvData[1], ":", 2)
				sdk.spiName = strings.TrimSpace(spiValues[0])
				sdk.spiExtName = strings.TrimSpace(spiValues[1])
			} else {
				sdk.spiName = strings.TrimSpace(kvData[1])
			}
		}
	}

	return nil
}

func (sdk sdkOpt) options() parser.ParseOptions {
	options := parser.SdkOptions{
		parser.WithVersion(sdk.ver),
	}

	return parser.ParseOptions{parser.WithSDK(sdk.name, options...)}
}

type outputOpt []string

func (out outputOpt) Type() string { return "OUTPUT" }

func (out outputOpt) String() string {
	return fmt.Sprintf("%+v", ([]string)(out))
}

func (out *outputOpt) Set(v string) error {
	for v := range strings.SplitSeq(v, ",") {
		module := strings.ToLower(strings.TrimSpace(v))

		switch module {
		case "api", "spi", "thost":
		default:
			return errors.New("invalid module name")
		}

		if slices.Contains(*out, module) {
			continue
		}

		*out = append(*out, module)
	}

	return nil
}

func init() {
	showVersion := flag.Bool("version", false, "Show tool version")

	flag.StringVar(&dep, "dep", "", "CTP SDK dependencies base DIR")
	flag.StringVar(&plat, "plat", "future", "CTP platform: future | mini | etf")
	flag.Var(&sdk, "sdk", "CTP SDK catagory: trader | mduser")
	flag.BoolVar(&debug, "debug", false, "Debug print AST parsing")

	flag.BoolVar(&stdout, "stdout", false, "Print converted moduels to STDOUT")
	flag.Var(&output, "output", "Convert output modules")

	flag.Parse()

	if *showVersion {
		execName := filepath.Base(os.Args[0])
		fmt.Printf(
			"%s: %s [%s]\nBuild: %s by %s\n",
			execName, version, gitVersion, buildTime, goVersion,
		)

		os.Exit(0)
	}
}

func main() {
	defer parser.CTPEntry.Release()

	templates := map[string][]*template.Template{}

	tplBase, err := fs.Sub(tplFs, "templates")
	if err != nil {
		fmt.Fprintf(os.Stderr, "open embed templates dir failed: %+v", err)
		os.Exit(1)
	}

	for _, o := range output {
		for _, f := range cTplMapper[o] {
			tpl, err := template.New(f).Funcs(tplFuncs).ParseFS(tplBase, f)
			if err != nil {
				fmt.Fprintf(
					os.Stderr, "parse %s template failed: %+v",
					output, err,
				)
				os.Exit(1)
			}

			templates[o] = append(templates[o], tpl)
		}
	}

	options := parser.ParseOptions{parser.WithPlatform(plat)}
	if debug {
		options = append(options, parser.WithDebug())
	}

	if err := parser.CTPEntry.Parse(
		dep, append(options, sdk.options()...)...,
	); err != nil {
		fmt.Fprintf(os.Stderr, "parse failed: %+v\n", err)
		os.Exit(255)
	} else {
		fmt.Fprintf(
			os.Stdout, "entry file parsed: %s\n",
			parser.CTPEntry.EntryFile(),
		)
	}

	for mod, tpls := range templates {
		for _, v := range tpls {
			fmt.Fprintf(os.Stdout, "converting %s %s\n", mod, v.Name())

			var wr io.Writer

			if stdout {
				wr = os.Stdout
			} else {
				outFilePath := strings.TrimSuffix(v.Name(), ".gotmpl")

				if outFile, err := os.OpenFile(
					outFilePath,
					os.O_CREATE|os.O_TRUNC|os.O_WRONLY,
					os.ModePerm,
				); err != nil {
					fmt.Fprintf(os.Stderr, "open output file failed: %+v", err)
					os.Exit(2)
				} else {
					wr = outFile
					defer outFile.Close()
				}
			}

			if err := v.Execute(wr, &parser.CTPEntry); err != nil {
				fmt.Fprintf(os.Stderr, "convert failed: %+v", err)
			}
		}
	}
}
