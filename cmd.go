package main

import (
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"bitbucket.org/JeremySchlatter/go-atexit"

	"github.com/mailgun/godebug/Godeps/_workspace/src/golang.org/x/tools/go/loader"
)

var w = flag.Bool("w", false, "write result to (source) file instead of stdout")

func usage() {
	log.Print(
		`godebug is a tool for debugging Go programs.

Usage:

        godebug command [arguments]

The commands are:

    run       compile, run, and debug a Go program
    output    generate debug source code, but do not build or run it

Use "godebug help [command]" for more information about a command.
`)
	exit(0)
}

func runUsage() {
	log.Print(
		`usage: godebug run gofiles... [arguments...]

Run is a wrapper around 'go run'. It generates debugging code for
the named Go source files and runs 'go run' on the result.
`)
}

func outputUsage() {
	log.Print(
		`usage: godebug output [-w] <packages>

Output outputs debugging code for <packages>.

By default, output will print the resulting code to stdout.
If the -w flag is given, output will overwrite the original
source files. Use with caution.

<packages> may take one of two forms:

    1. A list of *.go source files.

       All of the specified files are loaded, parsed and type-checked
       as a single package.  All the files must belong to the same directory.

    2. A list of import paths, each denoting a package.

       The package's directory is found relative to the $GOROOT and
       $GOPATH using similar logic to 'go build', and the *.go files in
       that directory are loaded, parsed and type-checked as a single
       package.

       In addition, all *_test.go files in the directory are then loaded
       and parsed.  Those files whose package declaration equals that of
       the non-*_test.go files are included in the primary package.  Test
       files whose package declaration ends with "_test" are type-checked
       as another package, the 'external' test package, so that a single
       import path may denote two packages.
`)
}

func main() {
	log.SetFlags(0)
	flag.Parse()
	flag.Usage = usage
	if flag.NArg() == 0 {
		usage()
	}
	switch flag.Arg(0) {
	case "help":
		doHelp()
	case "output":
		doOutput()
	case "run":
		doRun()
	default:
		usage()
	}
}

func doHelp() {
	if flag.NArg() < 2 {
		usage()
	}
	switch flag.Arg(1) {
	case "output":
		outputUsage()
	case "run":
		runUsage()
	default:
		log.Printf("Unknown help topic `%s`. Run 'godebug help'.\n", flag.Arg(1))
	}
}

func doRun() {
	atexit.TrapSignals()
	defer atexit.CallExitFuncs()

	tmpDir := makeTmpDir()
	atexit.Run(func() {
		removeDir(tmpDir)
	})

	var gofiles []string
	for _, arg := range flag.Args()[1:] {
		if !strings.HasSuffix(arg, ".go") {
			break
		}
		gofiles = append(gofiles, arg)
	}
	if len(gofiles) == 0 {
		logFatal("godebug run: no go files listed")
	}
	var conf loader.Config
	conf.SourceImports = true
	if err := conf.CreateFromFilenames("main", gofiles...); err != nil {
		logFatal(err)
	}
	prog, err := conf.Load()
	if err != nil {
		logFatal(err)
	}
	generate(prog, func(filename string) io.WriteCloser {
		f, err := os.Create(filepath.Join(tmpDir, filepath.Base(filename)))
		if err != nil {
			logFatal(err)
		}
		return f
	})
	args := []string{"run"}
	args = append(args, mapToTmpDir(tmpDir, gofiles)...)
	shell("go", args...)
}

func shell(command string, args ...string) {
	cmd := exec.Command(command, args...)
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	switch err.(type) {
	case nil:
	case *exec.ExitError:
		exit(1)
	default:
		log.Fatal(err)
	}
}

func mapToTmpDir(tmpDir string, gofiles []string) []string {
	result := make([]string, len(gofiles))
	for i := range gofiles {
		result[i] = filepath.Join(tmpDir, filepath.Base(gofiles[i]))
	}
	return result
}

func makeTmpDir() (dirname string) {
	tmp, err := ioutil.TempDir("", "godebug")
	if err != nil {
		logFatal("Failed to create temporary directory:", err)
	}
	return tmp
}

func removeDir(dir string) {
	if err := os.RemoveAll(dir); err != nil {
		log.Print("Failed to clean up temporary directory:", err)
	}
}

func doOutput() {
	var conf loader.Config
	rest, err := conf.FromArgs(flag.Args()[1:], true)
	if len(rest) > 0 {
		fmt.Fprintf(os.Stderr, "Unrecognized arguments:\n%v\n\n", strings.Join(rest, "\n"))
	}
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error identifying packages: %v\n\n", err)
	}
	if len(rest) > 0 || err != nil {
		flag.Usage()
	}
	conf.SourceImports = true
	prog, err := conf.Load()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error loading packages: %v\n\n", err)
		flag.Usage()
	}
	generate(prog, func(filename string) io.WriteCloser {
		if *w {
			file, err := os.Create(filename)
			if err != nil {
				logFatal(err)
			}
			return file
		}
		return nopCloser{os.Stdout}
	})
}

type nopCloser struct {
	io.Writer
}

func (nopCloser) Close() error {
	return nil
}

func logFatal(v ...interface{}) {
	atexit.CallExitFuncs()
	log.Fatal(v...)
}

func exit(n int) {
	atexit.CallExitFuncs()
	os.Exit(n)
}
