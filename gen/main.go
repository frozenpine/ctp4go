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
	"strconv"
	"strings"
	"text/template"

	"github.com/frozenpine/ctp4go/gen/handlers"
	"github.com/frozenpine/ctp4go/gen/parser"
)

var (
	version, goVersion, gitVersion, buildTime string

	dep    string
	plat   string
	sdk    = sdkOpt{}
	debug  bool
	stdout bool
	output outputOpt
	clean  bool

	cTplMapper = map[string][]string{
		"api": {
			"api_helper.h.tpl:$platform/$version",
			"api_helper.c.tpl:$platform/$version",
			"api_impl.go.tpl:$platform/$version",
			"consts_linux.go.tpl:$platform/$version",
			"consts_windows.go.tpl:$platform/$version",
			"register.go.tpl:$platform/$version",
			"imp_$version.go.tpl:$platform",
			"$sdkName_api.go.tpl:../thost/$platform",
		},
		"spi": {
			"spi_helper.h.tpl:$platform/$version",
			"spi_helper.c.tpl:$platform/$version",
			"spi_impl.go.tpl:$platform/$version",
			"$sdkName_spi.go.tpl:../thost/$platform",
		},
		"thost": {
			"ctp_types.go.tpl:$platform/types:true",
			"ctp_structs.go.tpl:$platform",
		},
	}

	paramMapper = map[string]func(*parser.Entry) string{
		"$version":  func(e *parser.Entry) string { return e.Sdk().Version() },
		"$platform": func(e *parser.Entry) string { return e.Platform() },
		"$sdkName":  func(e *parser.Entry) string { return e.Sdk().Name() },
	}

	tplFuncs = template.FuncMap{
		"ToUpper": strings.ToUpper,
		"ToLower": strings.ToLower,
		"Title": func(in string) string {
			if in == "" {
				return ""
			}

			return strings.ToTitle(string(in[0])) + in[1:]
		},
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
		"GoParamType": handlers.GoParamType,
		"GoType":      handlers.GoType,
		"GoTypeName":  handlers.GoTypeName,
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

func (sdk sdkOpt) options() parser.EntryOptions {
	options := parser.SdkOptions{
		parser.WithVersion(sdk.ver),
	}

	if sdk.hdrFileName != "" {
		options = append(options, parser.WithHdrFileName(sdk.hdrFileName))
	}

	if sdk.apiName != "" {
		options = append(options, parser.WithApiName(sdk.apiName))
	}

	if sdk.spiName != "" {
		options = append(options, parser.WithSpiName(sdk.spiName))
	}

	return parser.EntryOptions{parser.WithSDK(sdk.name, options...)}
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
	flag.BoolVar(&clean, "clean", false, "Clean old DIR files")

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

type tplDefine struct {
	tpl      *template.Template
	dir      string
	dirClean bool
}

func (d *tplDefine) parseDefine(tplBase fs.FS, mapper string) error {
	mapValues := strings.SplitN(mapper, ":", 3)

	genTpl, err := template.New(
		mapValues[0],
	).Funcs(tplFuncs).ParseFS(
		tplBase, mapValues[0],
	)
	if err != nil {
		return fmt.Errorf(
			"parse %s template failed: %+v", output, err,
		)
	}

	d.tpl = genTpl

	switch len(mapValues) {
	case 3:
		d.dirClean, _ = strconv.ParseBool(mapValues[2])

		fallthrough
	case 2:
		if mapValues[1] != "" {
			d.dir = mapValues[1]
		}
	}

	return nil
}

func handleParam(in string, entry *parser.Entry) string {
	for param, to := range paramMapper {
		in = strings.ReplaceAll(in, param, to(entry))
	}

	return in
}

func (d tplDefine) execute(entry *parser.Entry) error {
	var wr io.Writer = os.Stdout

	fileName := strings.TrimSuffix(
		handleParam(d.tpl.Name(), entry), ".tpl",
	)

	if !stdout {
		if d.dir != "" {
			d.dir = handleParam(d.dir, entry)

			if d.dirClean && clean {
				if err := os.RemoveAll(d.dir); err != nil {
					return fmt.Errorf("clean package dir failed: %+v", err)
				}
			}

			if err := os.MkdirAll(d.dir, os.ModePerm); err != nil {
				return fmt.Errorf("mkdir package dir failed: %+v", err)
			}
		}

		outFilePath := filepath.Join(d.dir, fileName)

		outFile, err := os.OpenFile(
			outFilePath,
			os.O_CREATE|os.O_TRUNC|os.O_WRONLY,
			os.ModePerm,
		)
		if err != nil {
			return fmt.Errorf("open output file failed: %+v", err)
		}

		wr = outFile
		defer outFile.Close()
	}

	if err := d.tpl.Execute(wr, entry); err != nil {
		return fmt.Errorf("convert failed: %+v", err)
	}

	return nil
}

func main() {
	templates := map[string][]tplDefine{}

	tplBase, err := fs.Sub(tplFs, "templates")
	if err != nil {
		fmt.Fprintf(os.Stderr, "open embed templates dir failed: %+v\n", err)
		os.Exit(1)
	}

	for _, o := range output {
		for _, mapper := range cTplMapper[o] {
			var d tplDefine

			if err := d.parseDefine(tplBase, mapper); err != nil {
				fmt.Fprintf(os.Stderr, "%+v\n", err)
				os.Exit(1)
			}

			templates[o] = append(templates[o], d)
		}
	}

	options := sdk.options()
	if debug {
		options = append(options, parser.WithDebug())
	}

	entry, err := parser.NewEntry(plat, dep, options...)
	if err != nil {
		fmt.Fprintf(os.Stderr, "create parse entry failed: %+v\n", err)
		os.Exit(255)
	}
	defer entry.Release()

	if err := entry.Parse(); err != nil {
		fmt.Fprintf(os.Stderr, "parse failed: %+v\n", err)
		os.Exit(255)
	} else {
		fmt.Fprintf(
			os.Stdout, "entry file parsed: %s\n",
			entry.EntryFile(),
		)
	}

	for mod, tpls := range templates {
		for _, v := range tpls {
			fmt.Fprintf(os.Stdout, "converting %s %s\n", mod, v.tpl.Name())

			if err := v.execute(entry); err != nil {
				fmt.Fprintf(os.Stderr, "execute template failed: %+v", err)

				os.Exit(2)
			}
			fmt.Fprintf(
				os.Stdout, "converting done: %s %s\n", mod, v.tpl.Name(),
			)
		}
	}
}
