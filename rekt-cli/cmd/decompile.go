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
	name         string
	verbose      bool
	outputDir    string
	inputArchive string
}

func NewDecompileCommand(flags []string) *DecompileCmd {
	cmd := &DecompileCmd{
		name: DecompileCmdName,
	}

	fs := flag.NewFlagSet(cmd.name, flag.ExitOnError)
	fs.BoolVar(&cmd.verbose, "verbose", false, "Verbose output")
	fs.StringVar(&cmd.outputDir, "outputDir", "", "Output directory where the decompiled app will be written to")
	fs.StringVar(&cmd.inputArchive, "archive", "", "Input archive to decompile. Required.")

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
func (cmd *DecompileCmd) Run() any {
	if cmd.inputArchive == "" {
		fmt.Println("An input archive must be passed: rekt decompile -archive=<path-to-app-archive>")
		os.Exit(1)
	}

	if cmd.verbose {
		PrintAscii()
	}

	var outDir string

	if strings.HasSuffix(cmd.inputArchive, ".apk") {
		outDir = decompileAndroid(cmd)
	} else if strings.HasSuffix(cmd.inputArchive, ".ipa") {
		outDir = decompileIOS(cmd)
	} else {
		fmt.Println("Unsupported archive format. Must be APK or IPA")
		os.Exit(1)
	}

	return outDir
}

func decompileAndroid(cmd *DecompileCmd) string {
	baseName := strings.ReplaceAll(strings.TrimSuffix(filepath.Base(cmd.inputArchive), ".apk"), " ", "_")
	outDir := cmd.outputDir
	if cmd.outputDir == "" {
		outDir = path.Join("./", "scan", baseName)
	}

	fmt.Println("Decompiling APK using jadx...")
	utils.ExecCommand("jadx", "-d", outDir, cmd.inputArchive)

	assetDir := path.Join(outDir, "resources/assets")
	decompileRN(assetDir, "android")

	fmt.Println("Decompiled app. Output directory:", outDir)
	return outDir
}

func decompileIOS(cmd *DecompileCmd) string {
	baseName := strings.ReplaceAll(strings.TrimSuffix(filepath.Base(cmd.inputArchive), ".ipa"), " ", "_")
	outDir := cmd.outputDir
	if cmd.outputDir == "" {
		outDir = path.Join("./", "scan", baseName)
	}

	fmt.Println("Unpacking IPA...")
	utils.Unzip(cmd.inputArchive, outDir)
	os.RemoveAll(path.Join(outDir, "META-INF"))
	os.RemoveAll(path.Join(outDir, "iTunesMetadata.plist"))

	matches, _ := utils.HasFile(strings.Join([]string{outDir, "Payload/*.app"}, "/"))
	decompileRN(matches[0], "ios")

	fmt.Println("Decompiled app. Output directory:", outDir)
	return outDir
}

func decompileRN(assetDir string, platform string) {
	bundleSuffix := fmt.Sprintf("index.%s", platform)
	var reactNativeBundle string
	if platform == "ios" {
		reactNativeBundle = path.Join(assetDir, "main.jsbundle")
	} else if platform == "android" {
		reactNativeBundle = path.Join(assetDir, fmt.Sprintf("%s.bundle", bundleSuffix))
	}

	_, ok := utils.HasFile(reactNativeBundle)
	if ok {
		out, _ := utils.ExecOutput("file", reactNativeBundle)
		isHermesBundle := strings.Contains(out, "Hermes JavaScript bytecode")
		if isHermesBundle {
			fmt.Println("Decompiling Hermes bundle...")
			utils.ExecCommand("hbc-disassembler", reactNativeBundle, path.Join(assetDir, fmt.Sprintf("%s.hasm", bundleSuffix)))
			utils.ExecCommand("hbc-decompiler", reactNativeBundle, path.Join(assetDir, fmt.Sprintf("%s.js", bundleSuffix)))
		} else {
			utils.CopyFile(reactNativeBundle, path.Join(assetDir, fmt.Sprintf("%s.js", bundleSuffix)))
		}
	}
}
