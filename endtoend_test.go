package main

// Partially modeled after golang.org/x/tools/cmd/stringer/endtoendtest.go

// This file runs tests in the testdata/single-file-tests directory

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"runtime"
	"strings"
	"sync"
	"syscall"
	"testing"
	"unsafe"

	"github.com/mailgun/godebug/Godeps/_workspace/src/github.com/kr/pty"
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
	cmd := exec.Command("go", "build", "-o", godebug+exe)
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		os.Remove(godebug)
		t.Fatal("failed to build godebug:", err)
	}
	return godebug
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
	lim := make(chan bool, *parallel)
	for test := range tests {
		lim <- true
		go func(test string) {
			defer func() {
				<-lim
			}()
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
	if !bytes.Equal(buf.Bytes(), golden) {
		if *accept {
			if err = ioutil.WriteFile(goldenOutput(test), buf.Bytes(), 0644); err != nil {
				t.Fatal(err)
			}
			return
		}
		t.Errorf("%s: want != got. Diff:\n%s", test, diff.Diff(string(golden), buf.String()))
	}
}

type session struct {
	// Inputs to send the debugger.
	inputs [][]byte

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
	cmd := exec.Command(tool, "run", goldenOutput(test))
	b, err := runGodebugSession(cmd, s)
	if err != nil {
		t.Logf("%s-out.go failed under '%s run': %v", test, tool, err)
		t.Fatalf("%s\n", b)
	}
	if s != nil {
		checkOutput(t, s, tool, b)
	}
}

// readBytes is like bufio.Reader.ReadBytes, except it takes a slice of bytes
// as the delimiter instead of a single byte.
func readBytes(buf *bufio.Reader, delim []byte) (line []byte, err error) {
	if len(delim) == 0 {
		panic("readBytes: zero-length delim")
	}
	for {
		tmp, err := buf.ReadBytes(delim[len(delim)-1])
		line = append(line, tmp...)
		if bytes.HasSuffix(line, delim) || err != nil {
			return line, err
		}
	}
}

// start assigns a pseudo-terminal tty os.File to c.Stdin, c.Stdout,
// and c.Stderr, calls c.Start, and returns the File of the tty's
// corresponding pty.
//
// It is copied from github.com/kr/pty, with a modification to set
// the window size before starting the command. This is done so the
// readline library will believe it is writing to a terminal.
func start(c *exec.Cmd) (*os.File, error) {
	p, tty, err := pty.Open()
	if err != nil {
		return nil, err
	}
	defer tty.Close()
	c.Stdout = tty
	c.Stdin = tty
	c.Stderr = tty

	// Set the window size of the terminal before starting the command.
	c.SysProcAttr = &syscall.SysProcAttr{Setctty: true, Setsid: true}
	_, _, errno := syscall.Syscall(syscall.SYS_IOCTL, p.Fd(), syscall.TIOCSWINSZ,
		uintptr(unsafe.Pointer(&struct {
			row, col       uint16
			xpixel, ypixel uint16
		}{
			// These numbers are arbitrary, except that col needs to be greater than the length of the prompt.
			5, 80, 50, 600,
		})))
	if errno != 0 {
		p.Close()
		return nil, errno
	}

	err = c.Start()
	if err != nil {
		p.Close()
		return nil, err
	}
	return p, err
}

var ptyLock sync.Mutex

func runGodebugSession(cmd *exec.Cmd, s *session) ([]byte, error) {
	ptyLock.Lock()
	p, err := start(cmd)
	ptyLock.Unlock()
	if err != nil {
		return nil, fmt.Errorf("command failed to start: %v", err)
	}
	defer p.Close()
	var b []byte
	buf := bufio.NewReader(p)
	if s != nil && len(s.inputs) != 0 {
		for _, cmd := range s.inputs {
			data, err := readBytes(buf, append([]byte("\r\n"), prompt...))
			b = append(b, data...)
			switch {
			case err == io.EOF, isPtmxError(err):
				break
			case err == nil:
				fmt.Fprint(p, string(cmd))
			default:
				return nil, fmt.Errorf("error reading from subprocess: %v", err)
			}
		}
	}
	data, readErr := ioutil.ReadAll(buf)
	err = cmd.Wait()
	if readErr != nil && !isPtmxError(readErr) {
		return nil, fmt.Errorf("error reading from subprocess: %v, %T", readErr)
	}
	b = append(b, data...)

	return normalizeCRLF(b), err
}

func isPtmxError(err error) bool {
	// There is a known issue with the tty library we are using where,
	// on Linux, we get "read /dev/ptmx: input/output error" instead of EOF
	// when a command exits.
	//
	// See https://github.com/kr/pty/issues/21
	pathErr, ok := err.(*os.PathError)
	return ok && pathErr.Op == "read" && pathErr.Path == "/dev/ptmx" && pathErr.Err == syscall.Errno(5)
}

func checkOutput(t *testing.T, want *session, tool string, got []byte) {
	testName := filepath.Base(want.filename)
	fmt.Printf("checking %s (%s)\n", testName, tool)
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

	lines := bytes.Split(b, newline)
	lines, comment := removeSessionComment(lines)
	s.comment = string(bytes.Join(comment, newline))

	for i, line := range lines {
		if bytes.HasSuffix(line, []byte{'\r'}) { // convert CRLF to LF
			line = line[:len(line)-1]
		}
		if i+1 < len(lines) {
			line = append(line, '\n')
		}

		if bytes.HasPrefix(line, prompt) {
			s.inputs = append(s.inputs, line[len(prompt):])
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
