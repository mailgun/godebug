package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"reflect"
	"runtime"
)

var vars = make(map[string]interface{})

func RecordVars(varmaps ...interface{}) {
	var i int
	for i = 0; i+1 < len(varmaps); i += 2 {
		s, ok := varmaps[i+1].(string)
		if !ok {
			panic("programming error: got even-numbered argument to RecordVars that was not a string")
		}
		vars[s] = varmaps[i]
	}
	if i != len(varmaps) {
		panic("programming error: called RecordVars with odd number of arguments")
	}
}

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
	var in []byte
	for {
		fmt.Print("(godebug) ")
		if !input.Scan() {
			fmt.Println("quitting session")
			return
		}
		in = input.Bytes()
		if v, ok := vars[string(bytes.TrimSpace(in))]; ok {
			fmt.Println(dereference(v))
			continue
		}
		fmt.Printf("Only variable printing is implemented. You typed: %q\n", input.Text())
	}
}

func dereference(i interface{}) interface{} {
	return reflect.ValueOf(i).Elem().Interface()
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
