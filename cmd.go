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

var (
	outputFlags flag.FlagSet
	w           = outputFlags.Bool("w", false, "write result to (source) file instead of stdout")

	runFlags   flag.FlagSet
	instrument = runFlags.String("instrument", "", "extra packages to enable for debugging")
)

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
		`usage: godebug run [-instrument pkgs...] gofiles... [arguments...]

Run is a wrapper around 'go run'. It generates debugging code for
the named Go source files and runs 'go run' on the result.

By default, godebug generates debugging code only for the named
Go source files, and not their dependencies. This means that in
the debugging session you will not be able to step into function
calls from imported packages. To instrument other packages,
pass the -instrument flag. Packages are comma-separated and
must not be relative.
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
	if len(os.Args) == 1 {
		usage()
	}
	switch os.Args[1] {
	case "help":
		doHelp(os.Args[2:])
	case "output":
		doOutput(os.Args[2:])
	case "run":
		doRun(os.Args[2:])
	default:
		usage()
	}
}

func doHelp(args []string) {
	if len(args) == 0 {
		usage()
	}
	switch args[0] {
	case "output":
		outputUsage()
	case "run":
		runUsage()
	default:
		log.Printf("Unknown help topic `%s`. Run 'godebug help'.\n", args[0])
	}
}

func doRun(args []string) {
	// Parse arguments.
	if err := runFlags.Parse(args); err != nil {
		logFatal(err)
	}

	// Make a temp directory.
	tmpDir := makeTmpDir()

	// Make sure we clean up the temporary directory if we get killed early.
	atexit.TrapSignals()
	defer atexit.CallExitFuncs()
	atexit.Run(func() {
		removeDir(tmpDir)
	})

	// Separate the .go files from the arguments to the binary we're building.
	gofiles, rest := getGoFiles()
	if len(gofiles) == 0 {
		logFatal("godebug run: no go files listed")
	}

	var conf loader.Config
	// SourceImports added because this was breaking when a package was not installed in GOPATH/pkg/.
	// There is probably a better solution.
	conf.SourceImports = true

	// Pass .go files to configuration.
	if err := conf.CreateFromFilenames("main", gofiles...); err != nil {
		logFatal(err)
	}

	// Mark the extra packages we want to instrument.
	var pkgs []string
	*instrument = strings.Trim(*instrument, ", ")
	if *instrument != "" {
		pkgs = strings.Split(*instrument, ",")
	}
	for _, pkg := range pkgs {
		if pkg == "all" || pkg == "std" {
			logFatalf("Special package %q not is supported in the -instrument flag.", pkg)
		}
		conf.Import(pkg)
	}

	// Generate debugging-enabled source files.
	prog, err := conf.Load()
	if err != nil {
		logFatal(err)
	}
	generate(prog, func(importPath, filename string) io.WriteCloser {
		if importPath == "main" {
			filename = filepath.Join(tmpDir, filepath.Base(filename))
		} else {
			switch strings.Split(filepath.Clean(importPath), string(filepath.Separator))[0] {
			case ".", "..":
				logFatalf("Package %q not supported because it has a relative import path. Sorry about that! File an issue at https://github.com/mailgun/godebug/issues/new and we'll come up with a fix for you.")
			}
			if err := os.MkdirAll(filepath.Join(tmpDir, "src", importPath), 0770); err != nil {
				logFatal(err)
			}
			filename = filepath.Join(tmpDir, "src", importPath, filepath.Base(filename))
		}
		f, err := os.Create(filename)
		if err != nil {
			logFatal(err)
		}
		return f
	})

	// Run 'go run'.
	goArgs := []string{"run"}
	goArgs = append(goArgs, mapToTmpDir(tmpDir, gofiles)...)
	goArgs = append(goArgs, rest...)
	shell(tmpDir, "go", goArgs...)
}

func getGoFiles() (gofiles, rest []string) {
	for i, arg := range runFlags.Args() {
		if !strings.HasSuffix(arg, ".go") {
			rest = runFlags.Args()[i:]
			break
		}
		gofiles = append(gofiles, arg)
	}
	return gofiles, rest
}

func shell(gopath, command string, args ...string) {
	cmd := exec.Command(command, args...)
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	setGopath(cmd, gopath)
	err := cmd.Run()
	switch err.(type) {
	case nil:
	case *exec.ExitError:
		exit(1)
	default:
		log.Fatal(err)
	}
}

func setGopath(cmd *exec.Cmd, gopath string) {
	cmd.Env = os.Environ()
	sawGopath := false
	for i := range cmd.Env {
		keyVal := strings.SplitN(cmd.Env[i], "=", 2)
		if keyVal[0] == "GOPATH" {
			cmd.Env[i] = "GOPATH=" + gopath + string(filepath.ListSeparator) + keyVal[1]
		}
	}
	if !sawGopath {
		cmd.Env = append(cmd.Env, "GOPATH="+gopath)
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

func doOutput(args []string) {
	if err := outputFlags.Parse(args); err != nil {
		logFatal(err)
	}
	var conf loader.Config
	rest, err := conf.FromArgs(args, true)
	if len(rest) > 0 {
		fmt.Fprintf(os.Stderr, "Unrecognized arguments:\n%v\n\n", strings.Join(rest, "\n"))
	}
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error identifying packages: %v\n\n", err)
	}
	if len(rest) > 0 || err != nil {
		usage()
	}
	conf.SourceImports = true
	prog, err := conf.Load()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error loading packages: %v\n\n", err)
		usage()
	}
	generate(prog, func(importPath, filename string) io.WriteCloser {
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

func logFatalf(format string, v ...interface{}) {
	atexit.CallExitFuncs()
	log.Fatalf(format, v...)
}

func exit(n int) {
	atexit.CallExitFuncs()
	os.Exit(n)
}
