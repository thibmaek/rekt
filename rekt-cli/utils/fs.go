package utils

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func CopyFile(src string, dest string) {
	srcFile, err := os.Open(src)
	if err != nil {
		panic(err)
	}
	defer srcFile.Close()

	destFile, err := os.Create(dest)
	if err != nil {
		panic(err)
	}
	defer destFile.Close()

	_, err = io.Copy(destFile, srcFile)
	if err != nil {
		panic(err)
	}

	err = destFile.Sync()
	if err != nil {
		panic(err)
	}
}

func HasFile(glob string) (matches []string, ok bool) {
	matches, err := filepath.Glob(glob)
	if err != nil {
		log.Fatal(err)
	}
	return matches, len(matches) > 0
}

/*
*

	FindInFile searches for a regular expression in a file.
	If excludes is not nil, it will exclude any lines that match the excludes regex.
	A custom buffer 'buf' can be passed to allow bigger lines to be read.
	* maxSize := 64 * 1024
	* buf := make([]byte, maxSize)
	* FindInFile("test.txt", regexp.MustCompile(`\d+`), nil, &buf)
*/
func FindInFile(filepath string, rgx *regexp.Regexp, excludes *regexp.Regexp, buf *[]byte) (ok bool, matches string) {
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	foundMatch := false
	var matchingLines strings.Builder

	scanner := bufio.NewScanner(file)
	if buf != nil {
		scanner.Buffer(*buf, len(*buf))
	}
	lineNumber := 1
	for scanner.Scan() {
		lineText := scanner.Text()
		matches := rgx.FindStringSubmatch(lineText)
		excl := []string{}
		if excludes != nil {
			excl = excludes.FindStringSubmatch(lineText)
		}
		if len(matches) > 0 && len(excl) == 0 {
			matchingLines.WriteString(fmt.Sprintf("Line %d: %s\n", lineNumber, lineText))
			// fmt.Printf("Line %d: %s\n", lineNumber, lineText)
			foundMatch = true
		}
		lineNumber++
	}

	if err := scanner.Err(); err != nil {
		if err != bufio.ErrTooLong {
			log.Fatal(err)
		}
	}

	return foundMatch, matchingLines.String()
}
