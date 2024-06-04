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
	fs.StringVar(&cmd.inputDir, "inputDir", "", "Input directory to probe. The output directory of a decompiled app archive")

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

func (cmd *ProbeCmd) ArchiveType() string {
	isIOS, _ := analysis.IsIOSApp(cmd.inputDir)
	if isIOS {
		return "ios"
	}
	return "android"
}

func (cmd *ProbeCmd) Run() any {
	if cmd.verbose {
		PrintAscii()
	}

	appType := analysis.GetAppType(cmd.inputDir)
	id, extras := analysis.GetBundleId(cmd.inputDir)

	if cmd.ArchiveType() == "android" {
		fmt.Println()
		fmt.Printf(`Analysis details:
		- Bundle ID: %s
		- MainApplication name: %s
		- App type: %s
		`, id, extras.MainApplication, appType)
		fmt.Println()
	} else {
		fmt.Println()
		fmt.Printf(`Analysis details:
		- Bundle ID: %s
		- App type: %s
		- App version: %s
		`, id, appType, extras.AppVersion)
		fmt.Println()
	}

	return nil
}
