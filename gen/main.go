package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/frozenpine/ctp4go/gen/internal"
)

var (
	entry    string
	platform string
	version  string
	sdk      string
	output   string
)

func init() {
	flag.StringVar(&entry, "entry", "", "Api headers base DIR")
	flag.StringVar(&platform, "platform", "future", "CTP platform")
	flag.StringVar(&version, "version", "", "CTP version")
	flag.StringVar(&sdk, "sdk", "mduser", "CTP SDK catagory")

	flag.StringVar(&output, "output", "", "Parsed data output")

	flag.Parse()
}

func main() {
	if err := internal.CTPEntry.Parse(
		entry,
		internal.WithPlatform(platform),
		internal.WithVersion(version),
		internal.WithSDK(sdk),
	); err != nil {
		fmt.Fprintf(os.Stderr, "parse failed: %+v\n", err)
	} else {
		fmt.Fprintf(
			os.Stdout, "entry file parsed: %s\n",
			internal.CTPEntry.EntryFile(),
		)
	}
}
