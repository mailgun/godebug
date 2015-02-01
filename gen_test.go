package main

// Modeled after golang.org/x/tools/cmd/stringer/endtoendtest.go

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
)

var (
	files  = flag.String("files", "", `Comma-separated list of files in the golden_tests directory to check. e.g. "example,name-conflicts". If not set, all of them will be checked.`)
	accept = flag.Bool("accept", false, "Accept the output of the program as the new golden file.")
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
	// Read the golden_tests directory
	fd, err := os.Open("golden_tests")
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
	for _, name := range names {
		if !strings.HasSuffix(name, "-out.go") && !strings.HasSuffix(name, "-in.go") {
			t.Fatal("Unexpected file in golden_tests directory:", name)
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
	golden, err := ioutil.ReadFile(filepath.Join("golden_tests", test+"-out.go"))
	if err != nil {
		t.Fatal(err)
	}
	cmd := exec.Command(godebug, filepath.Join("golden_tests", test+"-in.go"))
	var buf bytes.Buffer
	cmd.Stdout = &buf
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		fmt.Println(buf.String())
		t.Fatal(err)
	}
	if !bytes.Equal(buf.Bytes(), golden) {
		if *accept {
			if err = ioutil.WriteFile(filepath.Join("golden_tests", test+"-out.go"), buf.Bytes(), 0644); err != nil {
				t.Fatal(err)
			}
			return
		}
		diff := getDiff(filepath.Join("golden_tests", test+"-out.go"), buf.Bytes())
		fmt.Println(buf.String())
		t.Errorf("%s: got != want. Diff:\n%s", test, diff)
	}
}

func runGolden(t *testing.T, test string) {
	var buf bytes.Buffer
	cmd := exec.Command("go", "run", filepath.Join("golden_tests", test+"-out.go"))
	cmd.Stdout = &buf
	cmd.Stderr = &buf
	if err := cmd.Run(); err != nil {
		t.Errorf("Golden file failed to run: %v\n%s", err, buf.Bytes())
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
