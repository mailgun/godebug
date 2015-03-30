package main

// Partially modeled after golang.org/x/tools/cmd/stringer/endtoendtest.go

// This file runs tests in the testdata/single-file-tests directory

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
	"sync"
	"testing"

	"github.com/kylelemons/godebug/diff"
)

var (
	files         = flag.String("files", "", `Comma-separated list of files in the testdata/single-file-tests directory to check. e.g. "example,name-conflicts". If not set, all of them will be checked.`)
	accept        = flag.Bool("accept", false, "Accept the output of the program as the new golden file.")
	acceptSession = flag.Bool("accept-session", false, "If a *-session.txt file exists for a given test, accept any differences from running a new session and overwrite the file.")
)

func compileGodebug(t *testing.T) (filename string) {
	f, err := ioutil.TempFile("", "godebug")
	if err != nil {
		t.Fatal(err)
	}
	f.Close()
	godebug := f.Name()
	cmd := exec.Command("go", "build", "-o", godebug)
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		os.Remove(godebug)
		t.Fatal("failed to build godebug:", err)
	}
	return godebug
}

func TestGoldenFiles(t *testing.T) {
	godebug := compileGodebug(t)
	defer os.Remove(godebug)

	// Read the testdata directory
	dirname := filepath.FromSlash("testdata/single-file-tests")
	fd, err := os.Open(dirname)
	if err != nil {
		t.Fatal(err)
	}
	defer fd.Close()
	names, err := fd.Readdirnames(-1)
	if err != nil {
		t.Fatal("Readdirnames:", err)
	}
	tests := make(map[string]bool)
	skipped := make(map[string]bool)
	sessions := make(map[string][]string)
	if *files != "" {
		for _, name := range strings.Split(*files, ",") {
			tests[name] = true
		}
	}
	re := regexp.MustCompile(`(.*)-(?:out.go|in.go|(session(?:.*).txt))`)
	for _, name := range names {
		if name == "README.md" {
			continue
		}
		groups := re.FindStringSubmatch(name)
		if groups == nil {
			t.Fatal("Unexpected file in testdata directory:", name)
		}
		prefix, sessionName := groups[1], groups[2]
		if sessionName != "" {
			sessions[prefix] = append(sessions[prefix], name)
			continue
		}
		if *files == "" {
			tests[prefix] = true
			continue
		}
		if !tests[prefix] {
			skipped[prefix] = true
		}
	}
	for name := range skipped {
		fmt.Printf("Skipping golden test %q\n", name)
	}
	var wg sync.WaitGroup
	wg.Add(len(tests))
	for test := range tests {
		go func(test string) {
			defer wg.Done()
			oneTest(t, godebug, test, dirname, sessions[test])
		}(test)
	}
	wg.Wait()
}

func oneTest(t *testing.T, godebug, test, dirname string, goldenSessions []string) {
	compareGolden(t, godebug, test)
	if len(goldenSessions) == 0 {
		runGolden(t, test, nil)
	}
	var wg sync.WaitGroup
	wg.Add(len(goldenSessions))
	for _, filename := range goldenSessions {
		go func(filename string) {
			defer wg.Done()
			runGolden(t, test, parseSession(t, filepath.Join(dirname, filename)))
		}(filename)
	}
	wg.Wait()
}

func goldenOutput(testName string) (filename string) {
	return filepath.Join("testdata", "single-file-tests", testName+"-out.go")
}

func testInput(testName string) (filename string) {
	return filepath.Join("testdata", "single-file-tests", testName+"-in.go")
}

func compareGolden(t *testing.T, godebug, test string) {
	golden, err := ioutil.ReadFile(goldenOutput(test))
	if err != nil {
		t.Fatal(err)
	}
	cmd := exec.Command(godebug, "output", testInput(test))
	var buf bytes.Buffer
	cmd.Stdout = &buf
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		fmt.Println(buf.String())
		t.Fatal(err)
	}
	if !bytes.Equal(buf.Bytes(), golden) {
		if *accept {
			if err = ioutil.WriteFile(goldenOutput(test), buf.Bytes(), 0644); err != nil {
				t.Fatal(err)
			}
			return
		}
		diff := getDiff(goldenOutput(test), buf.Bytes())
		fmt.Println(buf.String())
		t.Errorf("%s: got != want. Diff:\n%s", test, diff)
	}
}

type session struct {
	// The bytes to send to stdin.
	input []byte

	// A transcript of the session as it would appear if run interactively in a terminal.
	fullSession []byte

	// The filename of this session inside testdata/single-file-tests/.
	filename string

	// The directory to change to before running cmd.
	workingDir string

	// The command to run. The first element must be "godebug".
	cmd []string
}

func runGolden(t *testing.T, test string, s *session) {
	var buf bytes.Buffer
	cmd := exec.Command("go", "run", goldenOutput(test))
	if s != nil {
		cmd.Stdin = bytes.NewReader(s.input)
	}
	cmd.Stdout = &buf
	cmd.Stderr = &buf
	if err := cmd.Run(); err != nil {
		t.Errorf("Golden file %s-out.go failed to run: %v\n%s", test, err, buf.Bytes())
	}
	if s != nil {
		checkOutput(t, s, buf.Bytes())
	}
}

func checkOutput(t *testing.T, want *session, output []byte) {
	testName := filepath.Base(want.filename)
	fmt.Println("checking", testName)
	got := interleaveCommands(want.input, output)
	if bytes.Equal(got, want.fullSession) {
		return
	}

	if *acceptSession {
		if err := ioutil.WriteFile(want.filename, got, 0644); err != nil {
			t.Fatal(err)
		}
		return
	}

	t.Errorf("%s: Session did not match. Diff:\n%v", testName, diff.Diff(string(want.fullSession), string(got)))
}

var prompt = []byte("(godebug) ")

// interleaveCommands reconstructs what a terminal session would have looked like,
// given the bytes sent to stdin and the bytes received from stdout. It assumes
// input only happens after prompts.
func interleaveCommands(input, output []byte) (combined []byte) {
	linesIn := bytes.Split(bytes.TrimSpace(input), []byte("\n"))
	if len(input) == 0 {
		linesIn = nil
	}
	linesOut := bytes.Split(output, []byte("\n"))
	var in, out int
	for ; out < len(linesOut); out++ {
		if bytes.HasPrefix(linesOut[out], prompt) {
			if in >= len(linesIn) {
				break
			}
			combined = append(combined, fmt.Sprintf("%s%s\n%s\n", prompt, linesIn[in], linesOut[out][len(prompt):])...)
			in++
		} else {
			combined = append(combined, linesOut[out]...)
			combined = append(combined, '\n')
		}
	}
	lengthCheck := func(index int, lines [][]byte, message string) {
		if index < len(lines) {
			combined = append(combined, message...)
			for ; index < len(lines); index++ {
				combined = append(combined, lines[index]...)
				combined = append(combined, '\n')
			}
		}
	}
	lengthCheck(in, linesIn, "<<< Input more lines than prompted for. Extra input lines: >>>\n")
	lengthCheck(out, linesOut, "<<< Session continued after our input stopped. Extra output: >>>\n")
	return combined
}

func parseSession(t *testing.T, filename string) *session {
	f, err := os.Open(filename)
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)

	var s session
	s.filename = filename

	scanSessionComment(t, scanner, &s, filename)

	for {
		b := append(scanner.Bytes(), '\n')

		if bytes.HasPrefix(b, prompt) {
			s.input = append(s.input, b[len(prompt):]...)
		}
		s.fullSession = append(s.fullSession, b...)

		if !scanner.Scan() {
			break
		}
	}
	if err := scanner.Err(); err != nil {
		t.Fatal(err)
	}
	return &s
}

// Scan past top of file comment. The top of file comment consists of any number of consecutive
// lines that are either blank or begin with the string "//".
// The comment may include "dir:" or "cmd:" directives, which we parse here.
func scanSessionComment(t *testing.T, scanner *bufio.Scanner, s *session, filename string) {
	i := 0
	for scanner.Scan() {
		i++
		b := scanner.Bytes()
		if len(b) == 0 {
			continue
		}
		if !bytes.HasPrefix(b, []byte("//")) {
			break
		}
		b = bytes.TrimSpace(b[2:])
		if bytes.HasPrefix(b, []byte("dir:")) {
			s.workingDir = string(bytes.TrimSpace(b[4:]))
		}
		if bytes.HasPrefix(b, []byte("cmd:")) {
			s.cmd = strings.Split(strings.TrimSpace(string(b[4:])), " ")
			if len(s.cmd) < 1 || s.cmd[0] != "godebug" {
				t.Fatalf("%s:%d The command listed in the session file must start with 'godebug'.", filename, i)
			}
		}
	}
	if err := scanner.Err(); err != nil {
		t.Fatal(err)
	}
}

func getDiff(filename string, inBuf []byte) []byte {
	var buf bytes.Buffer
	cmd := exec.Command("diff", "-", filename)
	cmd.Stdin = bytes.NewReader(inBuf)
	cmd.Stdout = &buf
	cmd.Run()
	return buf.Bytes()
}
