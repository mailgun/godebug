package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"sync"
	"testing"

	"github.com/kylelemons/godebug/diff"
	"gopkg.in/yaml.v1"
)

// This file runs tests in the testdata directory, excluding those in testdata/single-file-tests

func TestCLISessions(t *testing.T) {
	godebug := compileGodebug(t)
	defer os.Remove(godebug)

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
	tests := make([]string, 0, len(names))
	for _, name := range names {
		if strings.HasSuffix(name, ".yaml") {
			tests = append(tests, name)
		}
	}

	// Run tests in parallel
	var wg sync.WaitGroup
	for _, test := range tests {
		for _, c := range parseCases(t, filepath.Join("testdata", test)) {
			s := parseSessionFromBytes([]byte(c.Transcript))
			for _, tt := range c.Invocations {
				wg.Add(1)
				go func(cmd, dir, filename, desc string, nze bool, s *session) {
					defer wg.Done()
					runTest(t, godebug, cmd, dir, filename, desc, nze, s)
				}(tt.Cmd, tt.Dir, test, c.Desc, c.NonzeroExit, s)
			}
		}
	}
	wg.Wait()
}

type testCase struct {
	Invocations []struct {
		Dir, Cmd string
	}
	Desc, Transcript string
	NonzeroExit      bool `yaml:"nonzero_exit"`
}

func parseCases(t *testing.T, filename string) []testCase {
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		t.Fatal(err)
	}
	bb := bytes.Split(b, []byte("\n---\n"))
	testCases := make([]testCase, len(bb))
	for i := range bb {
		if err = yaml.Unmarshal(bb[i], &testCases[i]); err != nil {
			fmt.Println(string(bb[i]))
			t.Fatal(err)
		}
	}
	return testCases
}

func runTest(t *testing.T, godebug, command, dir, filename, description string, wantNonzeroExit bool, session *session) {
	var buf bytes.Buffer
	cmd := exec.Command(godebug, strings.Split(command, " ")[1:]...)
	cmd.Dir = filepath.FromSlash("testdata/test-filesystem/" + dir)
	cmd.Stdout = &buf
	cmd.Stderr = &buf
	cmd.Stdin = bytes.NewReader(session.input)
	setTestGopath(t, cmd)

	prefix := fmt.Sprintf("File: %s\nDescription: %s\nWorking dir: %s\nCommand: %s\nFailure:", filename, description, dir, command)

	err := cmd.Run()
	switch err.(type) {
	case nil:
		if wantNonzeroExit {
			t.Errorf("%s got exit code == 0, wanted a nonzero exit code.", prefix)
			return
		}
	case *exec.ExitError:
		if !wantNonzeroExit {
			t.Errorf("%s %q failed to run: %v\n%s", prefix, command, err, buf.Bytes())
		}
	default:
		t.Errorf("%s %q failed to run: %v\n%s", prefix, command, err, buf.Bytes())
		return
	}

	got := interleaveCommands(session.input, buf.Bytes())
	if equivalent(got, session.fullSession) {
		return
	}
	t.Errorf("%s Golden transcript did not match actual transcript. Diff:\n\n%v", prefix, diff.Diff(string(session.fullSession), string(got)))
}

// equivalent does a linewise comparison of a and b.
// Each line must be exactly equal or the want line must end in "//substr"
// and be a substring of the got line.
// Otherwise equivalent returns false.
func equivalent(got, want []byte) bool {
	var (
		gotLines  = bytes.Split(got, newline)
		wantLines = bytes.Split(want, newline)
		substr    = []byte("//substr")
		gg, ww    []byte
	)

	if len(gotLines) != len(wantLines) {
		return false
	}

	for i := range gotLines {
		gg, ww = gotLines[i], wantLines[i]
		if !(bytes.Equal(gg, ww) || bytes.HasSuffix(ww, substr) && bytes.Contains(gg, ww[:len(ww)-len(substr)])) {
			return false
		}
	}
	return true
}

func setTestGopath(t *testing.T, cmd *exec.Cmd) {
	cwd, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	setGopath(cmd, filepath.Join(cwd, "testdata", "test-filesystem", "gopath"))
}
