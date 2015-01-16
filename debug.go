package godebug

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"reflect"
	"runtime"
	"strings"
)

var vars = make(map[string]interface{})

// RecordVars records the mapping between variable names and their values.
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

// OutOfScope marks variables as having gone out of scope.
func OutOfScope(names ...string) {
	for _, n := range names {
		delete(vars, n)
	}
}

type state int

const (
	run state = iota
	next
	step
)

var currentState state
var currentDepth int
var debuggerDepth int

// EnterFunc marks the beginning of a function.
func EnterFunc() {
	currentDepth++
}

// ExitFunc marks the end of a function.
func ExitFunc() {
	currentDepth--
}

// Line marks a normal line where the debugger might pause.
func Line() {
	if currentState == run || (currentState == next && currentDepth != debuggerDepth) {
		return
	}
	debuggerDepth = currentDepth
	printLine()
	waitForInput()
}

var skipNextElseIfExpr bool

// ElseIfSimpleStmt marks a simple statement preceding an "else if" expression.
func ElseIfSimpleStmt(line string) {
	SLine(line)
	if currentState == next {
		skipNextElseIfExpr = true
	}
}

// ElseIfExpr marks an "else if" expression.
func ElseIfExpr(line string) {
	if skipNextElseIfExpr {
		skipNextElseIfExpr = false
		return
	}
	SLine(line)
}

// SLine is like Line, except that the debugger should print the provided line rather than
// reading the next line from the source code.
func SLine(line string) {
	if currentState == run || (currentState == next && currentDepth != debuggerDepth) {
		return
	}
	debuggerDepth = currentDepth
	fmt.Println("->", line)
	waitForInput()
}

// SetTrace is the entrypoint to the debugger.
func SetTrace() {
	currentState = step
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
			currentState = run
			return
		}
		s := input.Text()
		switch s {
		case "n", "next":
			currentState = next
			return
		case "s", "step":
			currentState = step
			return
		}
		if v, ok := vars[strings.TrimSpace(s)]; ok {
			fmt.Println(dereference(v))
			continue
		}
		var cmd, name string
		n, _ := fmt.Sscan(s, &cmd, &name)
		if n == 2 && (cmd == "p" || cmd == "print") {
			if v, ok := vars[strings.TrimSpace(name)]; ok {
				fmt.Println(dereference(v))
				continue
			}
		}
		fmt.Printf("Command not recognized, sorry! You typed: %q\n", s)
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
