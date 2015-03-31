package main

import (
	"bytes"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"sync"
	"testing"
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
		if strings.HasSuffix(name, ".txt") {
			tests = append(tests, name)
		}
	}

	// Run tests in parallel
	var wg sync.WaitGroup
	wg.Add(len(tests))
	for _, test := range tests {
		go func(filename string) {
			defer wg.Done()
			runTest(t, godebug, filename)
		}(filepath.Join("testdata", test))
	}
	wg.Wait()
}

func runTest(t *testing.T, godebug, filename string) {
	var buf bytes.Buffer
	session := parseSession(t, filename)
	cmd := exec.Command(godebug, session.cmd[1:]...)
	cmd.Dir = filepath.FromSlash("testdata/test-filesystem/" + session.workingDir)
	cmd.Stdout = &buf
	cmd.Stderr = &buf
	cmd.Stdin = bytes.NewReader(session.input)
	setTestGopath(t, cmd)
	if err := cmd.Run(); err != nil {
		t.Fatalf("From test %q, 'godebug %v' failed to run: %v\n%s", filepath.Base(filename), strings.Join(session.cmd[1:], " "), err, buf.Bytes())
	}
	checkOutput(t, session, buf.Bytes())
}

func setTestGopath(t *testing.T, cmd *exec.Cmd) {
	cwd, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	setGopath(cmd, filepath.Join(cwd, "testdata", "test-filesystem", "gopath"))
}
