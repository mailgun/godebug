package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"runtime"
)

// SetTrace is the entrypoint to the debugger.
func SetTrace() {
	printLine()
	waitForInput()
}

var input *bufio.Scanner

func init() {
	input = bufio.NewScanner(os.Stdin)
}

func waitForInput() {
	for {
		fmt.Print("(godebug) ")
		if !input.Scan() {
			fmt.Println("quitting session")
			return
		}
		fmt.Printf("Nothing is implemented yet. But you typed: %q\n", input.Text())
	}
}

func printLine() {
	_, file, line, ok := runtime.Caller(2)
	if !ok {
		fmt.Println("Hmm, something is broken. Failed to identify current source line.")
		return
	}
	printLineFromFile(line, file)
}

var parsedFiles map[string][]string

func init() {
	parsedFiles = make(map[string][]string)
}

func printLineFromFile(line int, file string) {
	f, ok := parsedFiles[file]
	if !ok {
		f = parseFile(file)
		parsedFiles[file] = f
	}
	if line >= len(f) {
		fmt.Printf("Hmm, something is broken. Current source line = %v, current file = %v, length of file = %v\n", line, file, len(f))
		return
	}
	fmt.Println("->", f[line])
}

func parseFile(file string) []string {
	f, err := os.Open(file)
	if err != nil {
		fmt.Println("Failed to open current source file:", err)
		return nil
	}
	var lines []string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lines = append(lines, string(bytes.TrimSpace(scanner.Bytes())))
	}
	if err = scanner.Err(); err != nil {
		fmt.Println("Error reading current source file:", err)
	}
	return lines
}
