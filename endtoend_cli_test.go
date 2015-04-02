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
				go func(cmd, dir, filename, desc string, s *session) {
					defer wg.Done()
					runTest(t, godebug, cmd, dir, filename, desc, s)
				}(tt.Cmd, tt.Dir, test, c.Desc, s)
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

func runTest(t *testing.T, godebug, command, dir, filename, description string, session *session) {
	var buf bytes.Buffer
	cmd := exec.Command(godebug, strings.Split(command, " ")[1:]...)
	cmd.Dir = filepath.FromSlash("testdata/test-filesystem/" + dir)
	cmd.Stdout = &buf
	cmd.Stderr = &buf
	cmd.Stdin = bytes.NewReader(session.input)
	setTestGopath(t, cmd)

	prefix := fmt.Sprintf("File: %s\nDescription: %s\nFailure:", filename, description)

	if err := cmd.Run(); err != nil {
		t.Errorf("%s %q failed to run: %v\n%s", prefix, command, err, buf.Bytes())
		return
	}

	got := interleaveCommands(session.input, buf.Bytes())
	if bytes.Equal(got, session.fullSession) {
		return
	}
	t.Errorf("%s Golden transcript did not match actual transcript. Diff:\n%v", prefix, diff.Diff(string(session.fullSession), string(got)))
}

func setTestGopath(t *testing.T, cmd *exec.Cmd) {
	cwd, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	setGopath(cmd, filepath.Join(cwd, "testdata", "test-filesystem", "gopath"))
}
