package main

import (
	"bytes"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
	"time"
)

func TestUpdatedSource(t *testing.T) {
	// A note on GOPATH directories.
	//
	// The $GOPATH environment variable can point to multiple directories.
	// When the go tool looks for a package, it checks these directories
	// one by one until it finds the package or determines that the package
	// is not present in any of the directories.
	//
	// A fact about this process that surprised me is documented in the
	// following comment taken from $GOROOT/src/cmd/go/pkg.go:
	//
	//     If a package p is not in the same tree as any package named on the
	//     command-line, assume [the installed package file] is up-to-date no
	//     matter what the modification times on the source files indicate.
	//     ...
	//     See issue 3149.
	//
	// godebug sets GOPATH=<godebug-generated-packages>:$GOPATH before
	// running the go tool. Since <godebug-generated-packages> includes
	// all of the packages named on the command-line, this meant that
	// the go tool considered _everything_ in the normal $GOPATH to be
	// up-to-date, and didn't recompile any of it.
	//
	// This behavior has led to some surprising results, where
	// non-godebug-instrumented packages would run whatever code was last
	// installed rather than the code that was sitting on disk in src/.
	//
	// This test checks that godebug no longer has that behavior.

	// The package that will be named on the command-line.
	main := []byte(
		`package main

		import "foo"

		func main() { foo.Foo() }
		`)
	mainTest := []byte(
		`package main

		import (
			"testing"

			"foo"
		)

		func TestFoo(t *testing.T) { foo.Foo() }
		`)

	// The old package that will be installed in pkg/.
	foo1 := []byte(
		`package foo

		import "fmt"

		func Foo() { fmt.Println("foo v1") }
		`)

	// The new package that will be sitting in src/.
	foo2 := []byte(
		`package foo

		import "fmt"

		func Foo() { fmt.Println("foo v2") }
		`)

	// Initial filesystem state.
	tmpDir, err := ioutil.TempDir("", "godebug-test")
	checkErr(t, err)
	defer os.RemoveAll(tmpDir)
	checkErr(t, os.MkdirAll(filepath.Join(tmpDir, "src", "foo"), 0770))
	checkErr(t, os.MkdirAll(filepath.Join(tmpDir, "src", "bar"), 0770))
	checkErr(t, ioutil.WriteFile(filepath.Join(tmpDir, "src", "bar", "main.go"), main, 0660))
	checkErr(t, ioutil.WriteFile(filepath.Join(tmpDir, "src", "bar", "main_test.go"), mainTest, 0660))
	checkErr(t, ioutil.WriteFile(filepath.Join(tmpDir, "src", "foo", "foo.go"), foo1, 0660))
	copyFiles(t, filepath.Join(tmpDir, "src", "github.com", "mailgun", "godebug", "lib"), "lib")
	copyFiles(t,
		filepath.Join(tmpDir, "src", "github.com", "mailgun", "godebug", "Godeps", "_workspace", "src", "github.com", "jtolds", "gls"),
		filepath.Join(os.Getenv("GOPATH"), "src", "github.com", "mailgun", "godebug", "Godeps", "_workspace", "src", "github.com", "jtolds", "gls"))
	copyFiles(t,
		filepath.Join(tmpDir, "src", "github.com", "mailgun", "godebug", "Godeps", "_workspace", "src", "github.com", "0xfaded", "eval"),
		filepath.Join(os.Getenv("GOPATH"), "src", "github.com", "mailgun", "godebug", "Godeps", "_workspace", "src", "github.com", "0xfaded", "eval"))
	copyFiles(t,
		filepath.Join(tmpDir, "src", "github.com", "mailgun", "godebug", "Godeps", "_workspace", "src", "github.com", "peterh", "liner"),
		filepath.Join(os.Getenv("GOPATH"), "src", "github.com", "mailgun", "godebug", "Godeps", "_workspace", "src", "github.com", "peterh", "liner"))

	// Install the first version of foo.
	cmd := exec.Command("go", "install", "foo")
	cmd.Env = os.Environ()
	overrideVar(cmd, "GOPATH", tmpDir)
	checkErr(t, cmd.Run())

	// Save the package file for later. Also serves to verify that it was installed.
	pkgFile := filepath.Join(tmpDir, "pkg", runtime.GOOS+"_"+runtime.GOARCH, "foo.a")
	pkgBytes, err := ioutil.ReadFile(pkgFile)
	checkErr(t, err)

	// Update foo.go. Make sure its modtime is later than the installed package file's modtime.
	checkErr(t, ioutil.WriteFile(filepath.Join(tmpDir, "src", "foo", "foo.go"), foo2, 0660))
	checkErr(t, os.Chtimes(pkgFile, time.Now().Add(-time.Hour), time.Now().Add(-time.Hour)))

	// Check that godebug run uses the new version of foo.
	godebug := compileGodebug(t)
	defer os.Remove(godebug)
	cmd = exec.Command(godebug, "run", filepath.Join(tmpDir, "src", "bar", "main.go"))
	cmd.Env = os.Environ()
	overrideVar(cmd, "GOPATH", tmpDir)
	out, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("godebug run failed: %v\n\n%s", err, out)
	}
	out = stripTestPrefix(out)
	if g, w := string(bytes.TrimSpace(out)), "foo v2"; g != w {
		if g == "foo v1" {
			t.Error("godebug run failed to recompile an out-of-date package")
		} else {
			t.Errorf("godebug run: got output %q, wanted output %q", g, w)
		}
	}

	// Put the old package back with an old modtime, leaving the new source file in place.
	checkErr(t, ioutil.WriteFile(pkgFile, pkgBytes, 0660))
	checkErr(t, os.Chtimes(pkgFile, time.Now().Add(-time.Hour), time.Now().Add(-time.Hour)))

	// Check that godebug test also uses the new version of foo.
	cmd = exec.Command(godebug, "test", "bar")
	cmd.Env = os.Environ()
	overrideVar(cmd, "GOPATH", tmpDir)
	out, err = cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("godebug run failed: %v\n\n%s", err, out)
	}
	out = stripTestPrefix(out)
	if g, w := string(bytes.Split(out, newline)[0]), "foo v2"; g != w {
		if g == "foo v1" {
			t.Error("godebug test failed to recompile an out-of-date package")
		} else {
			t.Errorf("godebug test: got output %q, wanted output %q", g, w)
		}
	}
}

func copyFiles(t *testing.T, dst, src string) {
	fis, err := ioutil.ReadDir(src)
	checkErr(t, err)
	checkErr(t, os.MkdirAll(dst, 0770))
	for _, fi := range fis {
		dstName := filepath.Join(dst, fi.Name())
		srcName := filepath.Join(src, fi.Name())
		if fi.IsDir() {
			checkErr(t, os.Mkdir(dstName, fi.Mode()))
			copyFiles(t, dstName, srcName)
		} else {
			s, err := os.Open(srcName)
			checkErr(t, err)
			defer s.Close()
			d, err := os.Create(dstName)
			checkErr(t, err)
			defer d.Close()
			_, err = io.Copy(d, s)
			checkErr(t, err)
		}
	}
}

func checkErr(t *testing.T, err error) {
	if err != nil {
		buf := make([]byte, 1024)
		t.Fatalf("%v\n\n%s", err, buf[:runtime.Stack(buf, false)])
	}
}

func overrideVar(cmd *exec.Cmd, key, val string) {
	for i, env := range cmd.Env {
		if strings.SplitN(env, "=", 2)[0] == key {
			cmd.Env[i] = key + "=" + val
			return
		}
	}
	cmd.Env = append(cmd.Env, key+"="+val)
}
