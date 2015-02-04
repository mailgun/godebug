package main

// Partially modeled after golang.org/x/tools/cmd/stringer/endtoendtest.go

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
	"testing"
)

var (
	files         = flag.String("files", "", `Comma-separated list of files in the testdata directory to check. e.g. "example,name-conflicts". If not set, all of them will be checked.`)
	accept        = flag.Bool("accept", false, "Accept the output of the program as the new golden file.")
	acceptSession = flag.Bool("accept-session", false, "If a *-session.txt file exists for a given test, accept any differences from running a new session and overwrite the file.")
)

func TestGoldenFiles(t *testing.T) {
	f, err := ioutil.TempFile("", "godebug")
	if err != nil {
		t.Fatal(err)
	}
	f.Close()
	godebug := f.Name()
	defer os.Remove(godebug)
	err = exec.Command("go", "build", "-o", godebug, "gen.go").Run()
	if err != nil {
		t.Fatal("failed to build godebug:", err)
	}
	// Read the testdata directory
	fd, err := os.Open("testdata")
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
	if *files != "" {
		for _, name := range strings.Split(*files, ",") {
			tests[name] = true
		}
	}
	re := regexp.MustCompile(".*-out.go|.*-in.go|.*-session.txt")
	for _, name := range names {
		if name == "README.md" {
			continue
		}
		if !re.MatchString(name) {
			t.Fatal("Unexpected file in testdata directory:", name)
		}
		prefix := name[:strings.LastIndex(name, "-")]
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
	for test := range tests {
		compareGolden(t, godebug, test)
		runGolden(t, test)
	}
}

func compareGolden(t *testing.T, godebug, test string) {
	golden, err := ioutil.ReadFile(filepath.Join("testdata", test+"-out.go"))
	if err != nil {
		t.Fatal(err)
	}
	cmd := exec.Command(godebug, filepath.Join("testdata", test+"-in.go"))
	var buf bytes.Buffer
	cmd.Stdout = &buf
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		fmt.Println(buf.String())
		t.Fatal(err)
	}
	if !bytes.Equal(buf.Bytes(), golden) {
		if *accept {
			if err = ioutil.WriteFile(filepath.Join("testdata", test+"-out.go"), buf.Bytes(), 0644); err != nil {
				t.Fatal(err)
			}
			return
		}
		diff := getDiff(filepath.Join("testdata", test+"-out.go"), buf.Bytes())
		fmt.Println(buf.String())
		t.Errorf("%s: got != want. Diff:\n%s", test, diff)
	}
}

var (
	firstDiff  = "first diff --> "
	whitespace = "               "
)

func runGolden(t *testing.T, test string) {
	input, goldenSession, checkSession := parseSession(t, test)
	var buf bytes.Buffer
	cmd := exec.Command("go", "run", filepath.Join("testdata", test+"-out.go"))
	if checkSession {
		cmd.Stdin = bytes.NewReader(input)
	}
	cmd.Stdout = &buf
	cmd.Stderr = &buf
	if err := cmd.Run(); err != nil {
		t.Errorf("Golden file failed to run: %v\n%s", err, buf.Bytes())
	}
	if !checkSession {
		return
	}
	session := interleaveCommands(input, buf.Bytes())
	if bytes.Equal(session, goldenSession) {
		return
	}
	if *acceptSession {
		if err := ioutil.WriteFile(filepath.Join("testdata", test+"-session.txt"), session, 0644); err != nil {
			t.Fatal(err)
		}
		return
	}
	goldLines, gotLines := bytes.Split(goldenSession, []byte("\n")), bytes.Split(session, []byte("\n"))
	var diff []byte
	i := 0
	for ; i < len(goldLines) && i < len(gotLines); i++ {
		if bytes.Equal(goldLines[i], gotLines[i]) {
			diff = append(diff, fmt.Sprintf("%s%s\n", whitespace, gotLines[i])...)
			continue
		}
		diff = append(diff, fmt.Sprintf("%s%s\n", firstDiff, gotLines[i])...)
		i++
		break
	}
	for ; i < len(gotLines); i++ {
		diff = append(diff, fmt.Sprintf("%s%s\n", whitespace, gotLines[i])...)
	}
	t.Errorf("%s: Session did not match. Got this session:\n%s", test, diff)
}

var prompt = []byte("(godebug) ")

// interleaveCommands reconstructs what a terminal session would have looked like,
// given the bytes sent to stdin and the bytes received from stdout. It assumes
// input only happens after prompts.
func interleaveCommands(input, output []byte) (combined []byte) {
	linesIn := bytes.Split(bytes.TrimSpace(input), []byte("\n"))
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

func parseSession(t *testing.T, test string) (input, fullSession []byte, ok bool) {
	f, err := os.Open(filepath.Join("testdata", test+"-session.txt"))
	if os.IsNotExist(err) {
		ok = false
		return
	}
	ok = true
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		b := append(scanner.Bytes(), '\n')
		if bytes.HasPrefix(b, prompt) {
			input = append(input, b[len(prompt):]...)
		}
		fullSession = append(fullSession, b...)
	}
	if err := scanner.Err(); err != nil {
		t.Fatal(err)
	}
	return
}

func getDiff(filename string, inBuf []byte) []byte {
	var buf bytes.Buffer
	cmd := exec.Command("diff", "-", filename)
	cmd.Stdin = bytes.NewReader(inBuf)
	cmd.Stdout = &buf
	cmd.Run()
	return buf.Bytes()
}
