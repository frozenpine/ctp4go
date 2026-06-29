package main

import (
	"bytes"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/frozenpine/ctp4go/gen/internal"
)

var (
	version, goVersion, gitVersion, buildTime string

	dep  string
	plat string
	sdk  = sdkOpt{
		name: "mduser",
		ver:  "v6.7.13",
	}
	output string
)

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

func (sdk sdkOpt) options() internal.ParseOptions {
	options := internal.SdkOptions{
		internal.WithVersion(sdk.ver),
	}

	return internal.ParseOptions{internal.WithSDK(sdk.name, options...)}
}

func init() {
	showVersion := flag.Bool("version", false, "Show tool version")

	flag.StringVar(&dep, "dep", "", "CTP SDK dependencies base DIR")
	flag.StringVar(&plat, "plat", "future", "CTP platform: future | mini | etf")
	flag.Var(&sdk, "sdk", "CTP SDK catagory: trader | mduser")

	flag.StringVar(&output, "output", "", "Parsed data output DIR")

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
	if err := internal.CTPEntry.Parse(
		dep, append(sdk.options(), internal.WithPlatform(plat))...,
	); err != nil {
		fmt.Fprintf(os.Stderr, "parse failed: %+v\n", err)
	} else {
		fmt.Fprintf(
			os.Stdout, "entry file parsed: %s\n",
			internal.CTPEntry.EntryFile(),
		)
	}
}
