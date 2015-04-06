package main

import (
	"bytes"
	"flag"
	"fmt"
	"go/build"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"

	"bitbucket.org/JeremySchlatter/go-atexit"

	"github.com/mailgun/godebug/Godeps/_workspace/src/golang.org/x/tools/go/loader"
)

var (
	outputFlags flag.FlagSet
	w           = outputFlags.Bool("w", false, "write result to (source) file instead of stdout")

	runTestFlags flag.FlagSet
	instrument   = runTestFlags.String("instrument", "", "extra packages to enable for debugging")
	work         = runTestFlags.Bool("godebugwork", false, "print the name of the temporary work directory and do not delete it when exiting")
)

func usage() {
	log.Print(
		`godebug is a tool for debugging Go programs.

Usage:

        godebug command [arguments]

The commands are:

    run       compile, run, and debug a Go program
    test      compile, run, and debug Go package tests
    output    generate debug source code, but do not build or run it

Use "godebug help [command]" for more information about a command.
`)
	exit(0)
}

func runUsage() {
	log.Print(
		`usage: godebug run [-godebugwork] [-instrument pkgs...] gofiles... [--] [arguments...]

Run is a wrapper around 'go run'. It generates debugging code for
the named Go source files and runs 'go run' on the result.

Optionally, a '--' argument ends the list of gofiles.

By default, godebug generates debugging code only for the named
Go source files, and not their dependencies. This means that in
the debugging session you will not be able to step into function
calls from imported packages. To instrument other packages,
pass the -instrument flag. Packages are comma-separated and
must not be relative.

If -godebugwork is set, godebug will print the name of the
temporary work directory and not delete it when exiting.
`)
}

func testUsage() {
	log.Print(
		`usage: godebug test [-godebugwork] [-instrument pkgs...] [packages] [flags for test binary]

Test is a wrapper around 'go test'. It generates debugging code for
the tests in the named packages and runs 'go test' on the result.

As with 'go test', by default godebug test needs no arguments.

By default, godebug generates debugging code only for the named
packages, and not their dependencies. This means that in the
debugging session you will not be able to step into function
calls from imported packages. To instrument other packages,
pass the -instrument flag. Packages are comma-separated and
must not be relative.

If -godebugwork is set, godebug will print the name of the
temporary work directory and not delete it when exiting.

See also: 'go help testflag'.
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
	atexit.TrapSignals()
	defer atexit.CallExitFuncs()
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
	case "test":
		doTest(os.Args[2:])
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
	case "test":
		testUsage()
	default:
		log.Printf("Unknown help topic `%s`. Run 'godebug help'.\n", args[0])
	}
}

func doRun(args []string) {
	// Parse arguments.
	exitIfErr(runTestFlags.Parse(args))

	// Separate the .go files from the arguments to the binary we're building.
	gofiles, rest := getGoFiles()
	if len(gofiles) == 0 {
		logFatal("godebug run: no go files listed")
	}

	// Build a loader.Config from the .go files.
	var conf loader.Config
	exitIfErr(conf.CreateFromFilenames("main", gofiles...))

	tmpDir := generateSourceFiles(&conf, "run")

	// Run 'go build -i' once without changing the GOPATH.
	// This will recompile and install any out-of-date packages.
	// When we modify the GOPATH in the next invocation of the go tool,
	// it will not check if any of the uninstrumented dependencies are out-of-date.
	shellGo("", []string{"build", "-o", os.DevNull, "-i"}, gofiles)

	// Run 'go build', then run the binary.
	// We do this rather than invoking 'go run' directly so we can implement the '--' argument,
	// which 'go run' does not have.
	bin := filepath.Join(tmpDir, "godebug.a.out")
	shellGo(tmpDir, []string{"build", "-o", bin}, mapToTmpDir(tmpDir, gofiles))
	shell("", bin, rest...)
}

func doTest(args []string) {
	// Parse arguments.
	packages, testFlags := parseTestArguments(args)

	// Default to the package in the current directory.
	if len(packages) == 0 {
		packages = []string{"."}
	}

	// Build a loader.Config from the provided packages.
	var conf loader.Config
	for _, pkg := range packages {
		exitIfErr(conf.ImportWithTests(pkg))
	}

	tmpDir := generateSourceFiles(&conf, "test")

	// Run 'go test -i' once without changing the GOPATH.
	// This will recompile and install any out-of-date packages.
	// When we modify the GOPATH in the next invocation of the go tool,
	// it will not check if any of the uninstrumented dependencies are out-of-date.
	shellGo("", []string{"test", "-i"}, packages)

	// First compile the test with -c and then run the binary directly.
	// This resolves some issues that came up with running 'go test' directly:
	//    (1) 'go test' changes the working directory to that of the source files of the test.
	//    (2) 'go test' does not forward stdin to the test binary.
	bin := filepath.Join(tmpDir, "godebug-test-bin.test")
	goArgs := []string{"test", "-c", "-o", bin}
	shellGo(tmpDir, goArgs, mapPkgsToTmpDir(packages))
	shell("", bin, testFlags...)
}

func generateSourceFiles(conf *loader.Config, subcommand string) (tmpDirPath string) {
	// Make a temp directory.
	tmpDir := makeTmpDir()
	if *work {
		// Print the name of the directory and don't clean it up on exit.
		fmt.Println(tmpDir)
	} else {
		// Clean up the directory on exit.
		atexit.Run(func() {
			removeDir(tmpDir)
		})
	}

	// Load the whole program from source files. This almost certainly causes more work than we need to do,
	// but it's an easy fix for a few problems I've encountered. Deleting this may be a good target for
	// future optimization work.
	conf.SourceImports = true

	// Mark the extra packages we want to instrument.
	var pkgs []string
	*instrument = strings.Trim(*instrument, ", ")
	if *instrument != "" {
		pkgs = strings.Split(*instrument, ",")
	}
	all := false
	var stdLib map[string]bool
	for _, pkg := range pkgs {
		// check for the special reserved package names: "main", "all", and "std"
		// see 'go help packages'
		switch pkg {
		case "main":
			switch subcommand {
			case "run": // The main package is always instrumented anyway. Carry on.
			case "test":
				logFatal(`godebug test: can't pass reserved name "main" in the -instrument flag.`)
			}
		case "all":
			if !all {
				all = true
				stdLib = getStdLibPkgs()
				fmt.Println(`godebug run: heads up: "all" means "all except std". godebug can't step into the standard library yet.` + "\n")
			}
		case "std":
			logFatalf("godebug %s: reserved name \"std\" cannot be passed in the -instrument flag."+
				"\ngodebug cannot currently instrument packages in the standard library."+
				"\nDo you wish it could? Chime in at https://github.com/mailgun/godebug/issues/12",
				subcommand)
		default:
			conf.Import(pkg)
		}
	}

	// Load the program.
	prog, err := conf.Load()
	exitIfErr(err)

	// If we're in "all" mode, mark all but the standard library packages and godebug itself for instrumenting.
	if all {
		markAlmostAllPackages(prog, stdLib)
	}

	// Generate debugging-enabled source files.
	wd := getwd()
	generate(prog, func(importPath, filename string) io.WriteCloser {
		if importPath == "main" {
			filename = filepath.Join(tmpDir, filepath.Base(filename))
		} else {
			importPath = findUnderGopath(wd, importPath)
			exitIfErr(os.MkdirAll(filepath.Join(tmpDir, "src", importPath), 0770))
			filename = filepath.Join(tmpDir, "src", importPath, filepath.Base(filename))
		}
		return createFileHook(filename, tmpDir)
	})
	return tmpDir
}

func markAlmostAllPackages(prog *loader.Program, stdLib map[string]bool) {
	for _, pkg := range prog.AllPackages {
		path := pkg.String()
		switch {

		// skip this package if...
		case stdLib[path]: // it's part of the standard library
		case prog.ImportMap[path] == nil: // it's a Created package
		case path == "github.com/mailgun/godebug/lib": // it's the godebug library
		case path == "github.com/jtolds/gls": // it's a dependency of the godebug library

		// otherwise include it
		default:
			prog.Imported[pkg.String()] = pkg
		}
	}
}

func getGoFiles() (gofiles, rest []string) {
	for i, arg := range runTestFlags.Args() {
		if arg == "--" {
			rest = runTestFlags.Args()[i+1:]
			break
		}
		if !strings.HasSuffix(arg, ".go") {
			rest = runTestFlags.Args()[i:]
			break
		}
		gofiles = append(gofiles, arg)
	}
	return gofiles, rest
}

func shellGo(tmpDir string, goArgs, packages []string) {
	shell(tmpDir, "go", append(goArgs, packages...)...)
}

func shell(gopath, command string, args ...string) {
	cmd := exec.Command(command, args...)
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	if gopath != "" {
		setGopath(cmd, gopath)
	}
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

func getwd() string {
	cwd, err := os.Getwd()
	if err != nil {
		logFatal("godebug needs to know the current working directory, but failed to determine it:", err)
	}
	return cwd
}

func mapPkgsToTmpDir(pkgs []string) []string {
	result := make([]string, len(pkgs))
	cwd := getwd()
	for i, pkg := range pkgs {
		result[i] = findUnderGopath(cwd, pkg)
	}
	return result
}

func findUnderGopath(cwd, pkg string) string {
	found, err := build.Import(pkg, cwd, build.FindOnly)
	if err != nil {
		logFatalf("Failed to find package %q in findUnderGopath. This is probably a bug -- please report it at https://github.com/mailgun/godebug/issues/new. Thanks!", pkg)
	}
	if found.SrcRoot == "" || found.ImportPath == "" {
		logFatalf("Looks like package %q is not in a GOPATH workspace. godebug doesn't support it right now, but if you open a ticket at https://github.com/mailgun/godebug/issues/new we'll fix it soon. Thanks!", pkg)
	}
	return found.ImportPath
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
	exitIfErr(outputFlags.Parse(args))

	var conf loader.Config
	rest, err := conf.FromArgs(outputFlags.Args(), true)
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
			return createFileHook(filename, "")
		}
		return nopCloser{os.Stdout}
	})
}

func parseTestArguments(args []string) (packages, testFlags []string) {
	// format: [-godebugwork] [-instrument pkgs...] [packages] [testFlags]

	// Find first unrecognized flag.
	sep := len(args)
	for i, arg := range args {
		if strings.HasPrefix(arg, "--") {
			arg = arg[1:]
		}
		if strings.HasPrefix(arg, "-") &&
			!strings.HasPrefix(arg, "-instrument") &&
			!strings.HasPrefix(arg, "-godebugwork") {
			sep = i
			break
		}
	}

	runTestFlags.Parse(args[:sep])
	return runTestFlags.Args(), args[sep:]
}

var (
	// For communicating with tests.
	logCreatedFiles bool
	logFileEnvVar   = "GODEBUG_LOG_CREATED_FILES"
	logFilePrefix   = "godebug created file: "
)

func init() {
	// This is only intended for tests, and so is not documented anywhere.
	if v := os.Getenv(logFileEnvVar); v != "" {
		logCreatedFiles, _ = strconv.ParseBool(v)
	}
}

// createFileHook is intended to capture all calls to os.Create.
// When we run under test, the tests can check if we are creating
// all and only the files we expect to.
func createFileHook(filename, tmpDir string) *os.File {
	if logCreatedFiles {
		if strings.HasPrefix(filename, tmpDir) {
			log.Println(logFilePrefix + "$TMP" + filename[len(tmpDir):])
		} else {
			log.Println(logFilePrefix + filename)
		}
	}
	file, err := os.Create(filename)
	exitIfErr(err)
	return file
}

func getStdLibPkgs() map[string]bool {
	pkgs := make(map[string]bool)
	b, err := exec.Command("go", "list", "std").Output()
	if err != nil {
		fmt.Printf("Failed to identify standard library packages. Here's the error from 'go list std':\n%s\n\nTry running again without passing \"all\" in the -instrument flag.", b)
	}
	b = bytes.TrimSpace(b)
	for _, pkg := range bytes.Split(b, []byte{'\n'}) {
		pkgs[string(pkg)] = true
	}
	return pkgs
}

type nopCloser struct {
	io.Writer
}

func (nopCloser) Close() error {
	return nil
}

func exitIfErr(err error) {
	if err != nil {
		logFatal(err)
	}
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
