package main

// Partially modeled after golang.org/x/tools/cmd/stringer/endtoendtest.go

// This file runs tests in the testdata/single-file-tests directory

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"runtime"
	"strings"
	"sync"
	"testing"

	"github.com/mailgun/godebug/Godeps/_workspace/src/github.com/kylelemons/godebug/diff"
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
	var exe string
	if runtime.GOOS == "windows" {
		exe = ".exe"
	}
	cmd := exec.Command("go", "build", "-o", godebug+exe, "-ldflags="+buildModeFlag)
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		os.Remove(godebug)
		t.Fatal("failed to build godebug:", err)
	}
	return godebug
}

// stripTestPrefix removes the note that godebug prints when it is built in test mode.
func stripTestPrefix(b []byte) []byte {
	prefix := []byte("godebug: test mode build\n")
	if !bytes.HasPrefix(b, prefix) {
		panic("Expected test mode note, but did not get one")
	}
	return b[len(prefix):]
}

func readDirNames(t *testing.T, dir string) (names []string) {
	fd, err := os.Open(dir)
	if err != nil {
		t.Fatal(err)
	}
	defer fd.Close()
	names, err = fd.Readdirnames(-1)
	if err != nil {
		t.Fatal("Readdirnames:", err)
	}
	return
}

func TestGoldenFiles(t *testing.T) {
	godebug := compileGodebug(t)
	defer os.Remove(godebug)

	dirname := filepath.FromSlash("testdata/single-file-tests")

	tests := make(map[string]bool)
	sessions := make(map[string][]string)
	if *files != "" {
		for _, name := range strings.Split(*files, ",") {
			tests[name] = true
		}
	}

	re := regexp.MustCompile(`(.*)-(?:out.go|in.go|(session(?:.*).txt))`)

	for _, name := range readDirNames(t, dirname) {
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
			fmt.Printf("Skipping golden test %q\n", prefix)
		}
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

var gopherjsAvailable = false

func init() {
	if _, err := exec.LookPath("gopherjs"); err == nil {
		gopherjsAvailable = true
	} else {
		fmt.Println("gopherjs is not in PATH. Skipping gopherjs tests.")
	}
}

func oneTest(t *testing.T, godebug, test, dirname string, goldenSessions []string) {
	compareGolden(t, godebug, test)
	if len(goldenSessions) == 0 {
		runGolden(t, test, "go", nil)
		if gopherjsAvailable {
			runGolden(t, test, "gopherjs", nil)
		}
	}
	var wg sync.WaitGroup
	wg.Add(len(goldenSessions))
	for _, filename := range goldenSessions {
		go func(filename string) {
			defer wg.Done()
			s := parseSession(t, filepath.Join(dirname, filename))
			runGolden(t, test, "go", s)
			if gopherjsAvailable {
				runGolden(t, test, "gopherjs", s)
			}
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
	golden = normalizeCRLF(golden)
	cmd := exec.Command(godebug, "output", testInput(test))
	var buf bytes.Buffer
	cmd.Stdout = &buf
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		fmt.Println(buf.String())
		t.Fatal(err)
	}
	output := stripTestPrefix(buf.Bytes())
	if !bytes.Equal(output, golden) {
		if *accept {
			if err = ioutil.WriteFile(goldenOutput(test), output, 0644); err != nil {
				t.Fatal(err)
			}
			return
		}
		t.Errorf("%s: want != got. Diff:\n%s", test, diff.Diff(string(golden), string(output)))
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

	// A comment at the top of the session file.
	comment string
}

func runGolden(t *testing.T, test, tool string, s *session) {
	var buf bytes.Buffer
	cmd := exec.Command(tool, "run", goldenOutput(test))
	if s != nil {
		cmd.Stdin = bytes.NewReader(s.input)
	}
	cmd.Stdout = &buf
	cmd.Stderr = &buf
	if err := cmd.Run(); err != nil {
		t.Errorf("Golden file %s-out.go failed to run under '%s run': %v\n%s", test, tool, err, buf.Bytes())
	}
	if s != nil {
		checkOutput(t, s, tool, buf.Bytes())
	}
}

func checkOutput(t *testing.T, want *session, tool string, output []byte) {
	testName := filepath.Base(want.filename)
	fmt.Printf("checking %s (%s)\n", testName, tool)
	got := interleaveCommands(want.input, output)
	if bytes.Equal(got, want.fullSession) {
		return
	}

	if *acceptSession {
		if want.comment != "" {
			got = append([]byte(want.comment+"\n"), got...)
		}
		if err := ioutil.WriteFile(want.filename, got, 0644); err != nil {
			t.Fatal(err)
		}
		return
	}

	t.Errorf("%s: Session did not match. Tool: %s, Diff:\n%v", testName, tool, diff.Diff(string(want.fullSession), string(got)))
}

var prompt = []byte("(godebug) ")
var newline = []byte("\n")

// interleaveCommands reconstructs what a terminal session would have looked like,
// given the bytes sent to stdin and the bytes received from stdout. It assumes
// input only happens after prompts.
func interleaveCommands(input, output []byte) (combined []byte) {
	linesIn := bytes.Split(input, newline)
	if len(input) == 0 {
		linesIn = nil
	} else if input[len(input)-1] == '\n' {
		linesIn = linesIn[:len(linesIn)-1]
	}
	chunks := bytes.Split(output, prompt)
	for i, chunk := range chunks {
		combined = append(combined, chunk...)
		if i != len(chunks)-1 && len(linesIn) > 0 {
			combined = append(combined, prompt...)
			combined = append(combined, linesIn[0]...)
			combined = append(combined, '\n')
			linesIn = linesIn[1:]
		}
	}
	for _, line := range linesIn {
		combined = append(combined, line...)
		combined = append(combined, '\n')
	}
	return combined
}

func parseSession(t *testing.T, filename string) *session {
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		t.Fatal(err)
	}
	s := parseSessionFromBytes(b)
	s.filename = filename
	return s
}

func parseSessionFromBytes(b []byte) *session {
	var s session

	b = normalizeCRLF(b)

	if bytes.HasSuffix(b, newline) {
		b = b[:len(b)-1]
	}

	lines := bytes.Split(b, newline)
	lines, comment := removeSessionComment(lines)
	s.comment = string(bytes.Join(comment, newline))

	for _, line := range lines {
		if bytes.HasSuffix(line, []byte{'\r'}) { // convert CRLF to LF
			line = line[:len(line)-1]
		}
		line = append(line, '\n')

		if bytes.HasPrefix(line, prompt) {
			s.input = append(s.input, line[len(prompt):]...)
		}
		s.fullSession = append(s.fullSession, line...)
	}

	return &s
}

// Scan past top of file comment. The top of file comment consists of any number of consecutive
// lines that are either blank or begin with the string "//".
func removeSessionComment(lines [][]byte) (content, comment [][]byte) {
	for i := range lines {
		if len(lines[i]) > 0 && !bytes.HasPrefix(lines[i], []byte("//")) {
			return lines[i:], lines[:i]
		}
	}
	return nil, lines
}

func normalizeCRLF(b []byte) []byte {
	return bytes.Replace(b, []byte("\r\n"), []byte("\n"), -1)
}
