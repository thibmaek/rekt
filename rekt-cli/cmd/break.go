package cmd

import (
	"flag"
	"fmt"
	"os"
	"path"
	"rekt/analysis"
	"rekt/android"
	"rekt/rn"
)

const (
	BreakCmdName = "break"
)

type BreakCmd struct {
	name     string
	verbose  bool
	inputDir string
}

func NewBreakCommand(flags []string) *BreakCmd {
	cmd := &BreakCmd{
		name: BreakCmdName,
	}

	fs := flag.NewFlagSet(cmd.name, flag.ExitOnError)
	fs.BoolVar(&cmd.verbose, "verbose", false, "Verbose output")
	fs.StringVar(&cmd.inputDir, "inputDir", "", "Input directory to check. The output directory of a decompiled APK")

	err := fs.Parse(flags)
	if err != nil {
		fs.PrintDefaults()
		os.Exit(1)
	}

	return cmd
}

func (cmd *BreakCmd) Name() string {
	return cmd.name
}

func (cmd *BreakCmd) ArchiveType() string {
	return analysis.GetArchiveType(cmd.inputDir)
}

func (cmd *BreakCmd) Run() any {
	if cmd.inputDir == "" {
		fmt.Println("No input directory passed!\nPass the input directory with the -inputDir flag.")
		os.Exit(1)
	}

	if cmd.verbose {
		PrintAscii()
	}

	assetsDir := path.Join(cmd.inputDir, "resources/assets")
	sourcesDir := path.Join(cmd.inputDir, "sources")

	if cmd.ArchiveType() == "android" {
		_, extras := analysis.GetAndroidBundleId(cmd.inputDir)

		android.CheckBuildConfig(sourcesDir, extras.MainApplication)
		android.CheckPrivateKeys(assetsDir)
		android.CheckAppCenterConfig(assetsDir)
		android.CheckAirshipConfig(assetsDir)
	} else {
		fmt.Println("Break is not yet supported on iOS")
		return nil
	}

	if rn.IsRNApp(cmd.inputDir) {
		if cmd.ArchiveType() == "android" {
			rn.ScanReactNativeBundle(path.Join(assetsDir, "index.android.js"))
		} else {
			rn.ScanReactNativeBundle(path.Join(assetsDir, "index.ios.js"))
		}
	}

	return nil
}
