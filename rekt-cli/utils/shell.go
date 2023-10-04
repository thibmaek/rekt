package utils

import (
	"log"
	"os/exec"
	"strings"
)

func ExecOutput(cmdName string, arg ...string) (string, error) {
	cmd := exec.Command(cmdName, arg...)
	stdout, err := cmd.CombinedOutput()
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(stdout)), nil
}

func ExecCommand(cmdName string, args ...string) {
	cmd := exec.Command(cmdName, args...)
	err := cmd.Run()
	if err != nil {
		log.Fatalf("Command failed with error: %v", err)
	}
}
