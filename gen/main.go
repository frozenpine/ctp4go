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
	sdk    sdkList
	debug  bool
	stdout bool
	output outputOpt
	clean  bool

	cleanUp []func()

	cTplMapper = map[string][]string{
		"api": {
			"api_helper.h.tpl:$platform/$version",
			"api_helper.c.tpl:$platform/$version",
			"api_impl.go.tpl:$platform/$version",
			"consts_linux.go.tpl:$platform/$version",
			"consts_windows.go.tpl:$platform/$version",
			"register.go.tpl:$platform/$version",
			"imp_$version.go.tpl:$platform",
		},
		"spi": {
			"spi_helper.h.tpl:$platform/$version",
			"spi_helper.c.tpl:$platform/$version",
			"spi_impl.go.tpl:$platform/$version",
		},
		"thost": {
			"ctp_types.go.tpl:$platform/types:true",
			"ctp_structs.go.tpl:$platform",
		},
		"extend": {
			"$sdkName_api.go.tpl:$platform",
			"$sdkName_spi.go.tpl:$platform",
			"log_spi.go.tpl:$platform:false:true",
		},
	}

	paramMapper = map[string]func(*parser.Entry) string{
		"$version":  func(e *parser.Entry) string { return e.Sdk().Version() },
		"$platform": func(e *parser.Entry) string { return e.Platform() },
		"$sdkName":  func(e *parser.Entry) string { return e.Sdk().Name() },
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

type sdkList []sdkOpt

func (l sdkList) Type() string { return "SDKList" }

func (l sdkList) String() string {
	buff := bytes.NewBufferString("[")

	for idx, sdk := range l {
		if idx > 0 {
			buff.WriteString(", ")
		}

		fmt.Fprintf(buff, "%+v", sdk)
	}

	buff.WriteByte(']')
	return buff.String()
}

func (l *sdkList) Set(v string) error {
	for d := range strings.SplitSeq(v, "|") {
		sdk := sdkOpt{}

		if err := sdk.Set(strings.TrimSpace(d)); err != nil {
			return err
		}

		*l = append(*l, sdk)
	}

	return nil
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
	combine  bool
}

func (d *tplDefine) parseDefine(tplBase fs.FS, mapper string) error {
	mapValues := strings.Split(mapper, ":")

	var err error

	switch len(mapValues) {
	case 4:
		d.combine, _ = strconv.ParseBool(mapValues[3])
		if d.combine && strings.Contains(mapper, "$sdkName") {
			return errors.New(
				"can not have $sdkName param in combined template",
			)
		}
		fallthrough
	case 3:
		d.dirClean, _ = strconv.ParseBool(mapValues[2])
		fallthrough
	case 2:
		if mapValues[1] != "" {
			d.dir = mapValues[1]
		}
	}

	d.tpl, err = template.New(
		mapValues[0],
	).Funcs(handlers.TplFuncs).ParseFS(
		tplBase, mapValues[0],
	)
	if err != nil {
		return fmt.Errorf(
			"parse %s template failed: %+v", output, err,
		)
	}

	return nil
}

func handleParam(in string, entry *parser.Entry) string {
	for param, to := range paramMapper {
		in = strings.ReplaceAll(in, param, to(entry))
	}

	return in
}

func prepareOutput(dir, fileName string, clean bool) (io.Writer, error) {
	if stdout {
		return os.Stdout, nil
	}

	if dir != "" {
		if clean {
			if err := os.RemoveAll(dir); err != nil {
				return nil, fmt.Errorf("clean package dir failed: %+v", err)
			}
		}

		if err := os.MkdirAll(dir, os.ModePerm); err != nil {
			return nil, fmt.Errorf("mkdir package dir failed: %+v", err)
		}
	}

	outFilePath := filepath.Join(dir, fileName)

	outFile, err := os.OpenFile(
		outFilePath,
		os.O_CREATE|os.O_TRUNC|os.O_WRONLY,
		os.ModePerm,
	)
	if err != nil {
		return nil, fmt.Errorf("open output file failed: %+v", err)
	}
	cleanUp = append(cleanUp, func() {
		outFile.Close()
	})

	return outFile, nil
}

func (d tplDefine) execute(entry *parser.Entry) error {
	wr, err := prepareOutput(
		handleParam(d.dir, entry),
		strings.TrimSuffix(
			handleParam(d.tpl.Name(), entry), ".tpl",
		),
		d.dirClean && clean,
	)
	if err != nil {
		return err
	}

	if err := d.tpl.Execute(wr, entry); err != nil {
		return fmt.Errorf("convert failed: %+v", err)
	}

	return nil
}

func (d tplDefine) execCombine(entries []*parser.Entry) error {
	if len(entries) < 1 {
		return errors.New("no entry found")
	}

	wr, err := prepareOutput(
		handleParam(d.dir, entries[0]),
		strings.TrimSuffix(
			handleParam(d.tpl.Name(), entries[0]), ".tpl",
		),
		d.dirClean && clean,
	)
	if err != nil {
		return err
	}

	if err := d.tpl.Execute(wr, map[string]any{
		"Platform": entries[0].Platform(),
		"Entries":  entries,
	}); err != nil {
		return fmt.Errorf("convert failed: %+v", err)
	}

	return nil
}

func main() {
	defer func() {
		for _, fn := range cleanUp {
			fn()
		}
	}()

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

	var entries []*parser.Entry

	for _, d := range sdk {
		options := d.options()
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
		}

		entries = append(entries, entry)
		fmt.Fprintf(
			os.Stdout, "entry file parsed: %s\n",
			entry.EntryFile(),
		)
	}

	var combined []tplDefine

	for _, entry := range entries {
		for mod, tpls := range templates {
			for _, v := range tpls {
				if v.combine {
					combined = append(combined, v)
					continue
				}

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

	for _, c := range combined {
		if err := c.execCombine(entries); err != nil {
			fmt.Fprintf(os.Stderr, "execute combine template failed: %+v", err)
			os.Exit(3)
		}
		fmt.Fprintf(
			os.Stdout, "converting combined done: %s\n", c.tpl.Name(),
		)
	}
}
