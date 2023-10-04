package main

import (
	"fmt"
	"os"
	cmd "rekt/cmd"
)

var Version string

type Command interface {
	Name() string
	Run() string
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please pass a Command. Options:")
		os.Exit(1)
	}

	cmdName := os.Args[1]
	flags := os.Args[2:]

	var com Command
	switch cmdName {
	case cmd.ProbeCmdName:
		com = cmd.NewProbeCommand(flags)
	case cmd.BreakCmdName:
		com = cmd.NewBreakCommand(flags)
	case cmd.DecompileCmdName:
		com = cmd.NewDecompileCommand(flags)
	case "intro":
		cmd.PrintAscii()
		os.Exit(0)
	case "version":
		fmt.Println("rekt v" + Version + " (Time to get rekt!)")
		os.Exit(0)
	default:
		fmt.Println("Unknown Command. Options:")
		os.Exit(1)
	}

	com.Run()
}
