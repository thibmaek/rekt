package cmd

import (
	"flag"
	"fmt"
	"os"
	analysis "rekt/analysis"
)

const (
	ProbeCmdName = "probe"
)

type ProbeCmd struct {
	name     string
	verbose  bool
	inputDir string
}

func NewProbeCommand(flags []string) *ProbeCmd {
	cmd := &ProbeCmd{
		name: ProbeCmdName,
	}

	fs := flag.NewFlagSet(cmd.name, flag.ExitOnError)
	fs.BoolVar(&cmd.verbose, "verbose", false, "Verbose output")
	fs.StringVar(&cmd.inputDir, "inputDir", "", "Input directory to probe. The output directory of a decompiled APK")

	err := fs.Parse(flags)
	if err != nil {
		fs.PrintDefaults()
		os.Exit(1)
	}

	return cmd
}

func (cmd *ProbeCmd) Name() string {
	return cmd.name
}

func (cmd *ProbeCmd) Run() string {
	if cmd.verbose {
		PrintAscii()
	}

	appType := analysis.GetAppType(cmd.inputDir)
	bundleId, mainApplication := analysis.GetBundleId(cmd.inputDir)

	fmt.Println()
	fmt.Printf(`Analysis details:
  - Bundle ID: %s
  - MainApplication name: %s
  - App type: %s
  `, bundleId, mainApplication, appType)
	fmt.Println()

	return ""
}
