package main

// Partly modeled after golang.org/x/tools/cmd/stringer/endtoendtest.go

import (
	"bytes"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"testing"
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
	tests := []string{"example"}
	for _, test := range tests {
		compareGolden(t, godebug, test)
		runGolden(t, test)
	}
}

func compareGolden(t *testing.T, godebug, test string) {
	golden, err := ioutil.ReadFile(filepath.Join("golden_tests", test+"-out.go"))
	if err != nil {
		t.Fatal(err)
	}
	cmd := exec.Command(godebug, filepath.Join("golden_tests", test+"-in.go"))
	var buf bytes.Buffer
	cmd.Stdout = &buf
	//cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(buf.Bytes(), golden) {
		diff := getDiff(filepath.Join("golden_tests", test+"-out.go"), buf.Bytes())
		t.Errorf("%s: got != want. Diff:\n%s", test, diff)
	}
}

func runGolden(t *testing.T, test string) {
	if err := exec.Command("go", "run", filepath.Join("golden_tests", test+"-out.go")).Run(); err != nil {
		t.Error("Golden file failed to run:", err)
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
