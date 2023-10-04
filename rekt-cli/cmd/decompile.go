package cmd

import (
	"flag"
	"fmt"
	"os"
	"path"
	"path/filepath"
	utils "rekt/utils"
	"strings"
)

const (
	DecompileCmdName = "decompile"
)

type DecompileCmd struct {
	name      string
	verbose   bool
	outputDir string
	inputApk  string
}

func NewDecompileCommand(flags []string) *DecompileCmd {
	cmd := &DecompileCmd{
		name: DecompileCmdName,
	}

	fs := flag.NewFlagSet(cmd.name, flag.ExitOnError)
	fs.BoolVar(&cmd.verbose, "verbose", false, "Verbose output")
	fs.StringVar(&cmd.outputDir, "outputDir", "", "Output directory where the decompiled app will be written to")
	fs.StringVar(&cmd.inputApk, "apk", "", "Input APK to decompile. Required.")

	err := fs.Parse(flags)
	if err != nil {
		fs.PrintDefaults()
		os.Exit(1)
	}

	return cmd
}

func (cmd *DecompileCmd) Name() string {
	return cmd.name
}

// Decompiles the app and returns the output directory
func (cmd *DecompileCmd) Run() string {
	if cmd.inputApk == "" {
		fmt.Println("APK must be passed: rekt decompile -apk=<path-to-apk>")
		os.Exit(1)
	}

	if cmd.verbose {
		PrintAscii()
	}

	baseName := strings.ReplaceAll(strings.TrimSuffix(filepath.Base(cmd.inputApk), ".apk"), " ", "_")
	outDir := cmd.outputDir
	if cmd.outputDir == "" {
		outDir = path.Join("./", "scan", baseName)
	}

	fmt.Println("Decompiling APK using jadx...")
	utils.ExecCommand("jadx", "-d", outDir, cmd.inputApk)

	assetDir := path.Join(outDir, "resources/assets")
	decompileRN(assetDir)

	fmt.Println("Decompiled app. Output directory:", outDir)
	return outDir
}

func decompileRN(assetDir string) {
	reactNativeBundle := path.Join(assetDir, "index.android.bundle")

	_, ok := utils.HasFile(reactNativeBundle)
	if ok {
		out, _ := utils.ExecOutput("file", reactNativeBundle)
		isHermesBundle := strings.Contains(out, "Hermes JavaScript bytecode")
		if isHermesBundle {
			fmt.Println("Decompiling Hermes bundle...")
			utils.ExecCommand("hbc-disassembler", reactNativeBundle, path.Join(assetDir, "index.android.hasm"))
			utils.ExecCommand("hbc-decompiler", reactNativeBundle, path.Join(assetDir, "index.android.js"))
		} else {
			utils.CopyFile(reactNativeBundle, path.Join(assetDir, "index.android.js"))
		}
	}
}
