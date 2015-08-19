package main

import (
	"bytes"
	"flag"
	"fmt"
	"go/ast"
	"go/build"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/mailgun/godebug/Godeps/_workspace/src/bitbucket.org/JeremySchlatter/go-atexit"
	"github.com/mailgun/godebug/Godeps/_workspace/src/github.com/kisielk/gotool"
	"github.com/mailgun/godebug/Godeps/_workspace/src/golang.org/x/tools/go/loader"
	"github.com/mailgun/godebug/gen"
)

var (
	outputFlags flag.FlagSet
	w           = outputFlags.Bool("w", false, "write result to (source) file instead of stdout")

	runFlags   flag.FlagSet
	instrument = runFlags.String("instrument", "", "extra packages to enable for debugging")
	work       = runFlags.Bool("godebugwork", false, "print the name of the temporary work directory and do not delete it when exiting")
	tags       = runFlags.String("tags", "", "go build tags")

	buildFlags = runFlags
	o          = buildFlags.String("o", "", "output binary name")

	testFlags = buildFlags
	c         = testFlags.Bool("c", false, "compile the test binary but do not run it")
)

func init() {
	// Hack for godebug's CI system. The CI can't override PATH in its builders,
	// but it can set new environment variables.
	if p := os.Getenv("GODEBUG_GO_PATH"); p != "" {
		os.Setenv("PATH", p+string(filepath.ListSeparator)+os.Getenv("PATH"))
	}
}

func usage() {
	log.Print(
		`godebug is a tool for debugging Go programs.

Usage:

        godebug command [arguments]

The commands are:

    build     compile a debug-ready Go program
    run       compile, run, and debug a Go program
    test      compile, run, and debug Go package tests

Use "godebug help [command]" for more information about a command.
`)
	exit(0)
}

const commonArgsUsage = `
By default, godebug generates debugging code only for the named
Go source files, and not their dependencies. This means that in
the debugging session you will not be able to step into function
calls from imported packages. To instrument other packages,
pass the -instrument flag. Packages are comma-separated and
must not be relative.

If -godebugwork is set, godebug will print the name of the
temporary work directory and not delete it when exiting.

-tags works like in 'go help build'.
`

func runUsage() {
	log.Print(
		`usage: godebug run [-godebugwork] [-instrument pkgs...] [-tags 'tag list'] gofiles... [--] [arguments...]

Run emulates 'go run' behavior. It generates debugging code for the
named *.go files and then compiles and executes the result.

Like 'go run' it takes a list of go files which are treated as a
single main package. The rest of the arguments is passed to the
binary. Optionally, a '--' argument ends the list of gofiles.
` + commonArgsUsage)
}

func buildUsage() {
	log.Print(
		`usage: godebug build [-godebugwork] [-instrument pkgs...] [-tags 'tag list'] [-o output] [package]

Build is a wrapper around 'go build'. It generates debugging code for
the named target and builds the result.

Like 'go build' it takes a single main package. If arguments are a list
of *.go files they are treated as a single package. Relative packages are
not supported - which means you can't leave the package name out, too.

The output file naming if -o is not passed works like 'go build' (see
'go help build') with the addition of the suffix '.debug'.
` + commonArgsUsage)
}

func testUsage() {
	log.Print(
		`usage: godebug test [-godebugwork] [-instrument pkgs...] [-tags 'tag list'] [-c] [-o output] [packages] [flags for test binary]

Test is a wrapper around 'go test'. It generates debugging code for
the tests in the named packages and runs 'go test' on the result.

As with 'go test', by default godebug test needs no arguments.

Flags parsing, -c and -o work like for 'go test' - see 'go help test'.
The default binary name for -c has the suffix '.test.debug'.

See also: 'go help testflag'. Note that you have to use the 'test.'
prefix like '-test.v'.
` + commonArgsUsage)
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
	case "build":
		doBuild(os.Args[2:])
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
	case "build":
		buildUsage()
	case "test":
		testUsage()
	default:
		log.Printf("Unknown help topic `%s`. Run 'godebug help'.\n", args[0])
	}
}

func doBuild(args []string) {
	exitIfErr(buildFlags.Parse(args))
	goArgs, isPkg := parseBuildArguments(buildFlags.Args())

	conf := newLoader()
	if isPkg {
		conf.Import(goArgs[0])
	} else {
		exitIfErr(conf.CreateFromFilenames("main", goArgs...))
	}

	tmpDir := generateSourceFiles(&conf, "build")
	tmpFile := filepath.Join(tmpDir, "godebug.-i.a.out")

	if doGopathWorkaround {
		// Rebuild stale packages, since this version of Go will not do so by default.
		shellGo("", []string{"build", "-o", tmpFile, "-tags", *tags, "-i"}, goArgs)
	}

	if isPkg {
		goArgs = mapPkgsToTmpDir(goArgs)
	} else {
		goArgs = mapToTmpDir(tmpDir, goArgs)
	}

	bin := filepath.Base(strings.TrimSuffix(goArgs[0], ".go")) + ".debug"
	if *o != "" {
		bin = *o
	}

	shellGo(tmpDir, []string{"build", "-o", bin, "-tags", *tags}, goArgs)
}

func doRun(args []string) {
	// Parse arguments.
	exitIfErr(runFlags.Parse(args))

	// Separate the .go files from the arguments to the binary we're building.
	gofiles, rest := getGoFiles()
	if len(gofiles) == 0 {
		logFatal("godebug run: no go files listed")
	}

	// Build a loader.Config from the .go files.
	conf := newLoader()
	exitIfErr(conf.CreateFromFilenames("main", gofiles...))

	tmpDir := generateSourceFiles(&conf, "run")

	if doGopathWorkaround {
		// Rebuild stale packages, since this version of Go will not do so by default.
		shellGo("", []string{"build", "-o", os.DevNull, "-tags", *tags, "-i"},
			gofiles)
	}

	// Run 'go build', then run the binary.
	// We do this rather than invoking 'go run' directly so we can implement
	// the '--' argument, which 'go run' does not have.
	bin := filepath.Join(tmpDir, "godebug.a.out")
	shellGo(tmpDir, []string{"build", "-tags", *tags, "-o", bin},
		mapToTmpDir(tmpDir, gofiles))
	shell("", bin, rest...)
}

func doTest(args []string) {
	// Parse arguments.
	packages, testFlags := parseTestArguments(args)

	// Default to the package in the current directory.
	if len(packages) == 0 {
		packages = []string{"."}
	}

	// Expand ...
	packages = gotool.ImportPaths(packages)

	if len(packages) > 1 && (*c || *o != "") {
		logFatal("godebug test: cannot use -c or -o flag with multiple packages")
	}

	// Build a loader.Config from the provided packages.
	conf := newLoader()
	for _, pkg := range packages {
		exitIfErr(conf.ImportWithTests(pkg))
	}

	tmpDir := generateSourceFiles(&conf, "test")
	wd := getwd()

	// Run 'go test -i' once without changing the GOPATH.
	// This will recompile and install any out-of-date packages.
	// When we modify the GOPATH in the next invocation of the go tool,
	// it will not check if any of the uninstrumented dependencies are out-of-date.
	shellGo("", []string{"test", "-tags", *tags, "-i"}, packages)

	// The target binary goes to -o if specified, otherwise to the default name
	// if -c is specified, otherwise to the temporary directory.
	bin := filepath.Join(tmpDir, "godebug-test-bin.test")
	if *c {
		bin = filepath.Base(mapPkgsToTmpDir(packages)[0]) + ".test.debug"
	}
	if *o != "" {
		bin = abs(*o)
	}

	// First compile the test with -c and then run the binary directly.
	// This resolves some issues that came up with running 'go test' directly:
	//    (1) 'go test' changes the working directory to that of the source files of the test.
	//    (2) 'go test' does not forward stdin to the test binary.
	// Do it once for each package since we can't use -c with multiple packages.
	for _, pkg := range mapPkgsToTmpDir(packages) {
		if len(packages) > 1 {
			fmt.Println("===", pkg)
			os.Remove(bin)
		}
		goArgs := []string{"test", "-tags", *tags, "-c", "-o", bin}
		shellGo(tmpDir, goArgs, []string{pkg})
		// Skip execution if no binary was generated (no test files) or -c was passed
		if _, err := os.Stat(bin); err == nil && !*c {
			_, dir := findUnderGopath(wd, pkg)
			os.Chdir(dir)
			shell("", bin, testFlags...)
			os.Chdir(wd)
		}
	}
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
				fmt.Println(`godebug run: heads up: "all" means "all except std". godebug can't step into the standard library yet.` + "\n")
			}
		case "std":
			logFatalf("godebug %s: reserved name \"std\" cannot be passed in the -instrument flag."+
				"\ngodebug cannot currently instrument packages in the standard library."+
				"\nDo you wish it could? Chime in at https://github.com/mailgun/godebug/issues/12",
				subcommand)
		default:
			for _, path := range gotool.ImportPaths([]string{pkg}) { // wildcard "..." expansion
				conf.Import(path)
			}
		}
	}

	// Load the program.
	prog, err := conf.Load()
	exitIfErr(err)

	// If we're in "all" mode, mark all but the standard library packages and godebug itself for instrumenting.
	stdLib := getStdLibPkgs()
	if all {
		markAlmostAllPackages(prog, stdLib)
	}

	// Warn the user if they have breakpoints set in files that we are not instrumenting.
	checkForUnusedBreakpoints(subcommand, prog, stdLib)

	// Generate debugging-enabled source files.
	wd := getwd()
	gen.Generate(prog, ioutil.ReadFile, func(importPath, filename string) io.WriteCloser {
		if importPath == "main" {
			filename = filepath.Join(tmpDir, filepath.Base(filename))
		} else {
			importPath, _ = findUnderGopath(wd, importPath)
			exitIfErr(os.MkdirAll(filepath.Join(tmpDir, "src", importPath), 0770))
			filename = filepath.Join(tmpDir, "src", importPath, filepath.Base(filename))
		}
		return createFileHook(filename, tmpDir)
	})
	return tmpDir
}

func newLoader() loader.Config {
	var conf loader.Config
	b := build.Default
	b.BuildTags = append(b.BuildTags, strings.Split(*tags, " ")...)
	conf.Build = &b
	return conf
}

func checkForUnusedBreakpoints(subcommand string, prog *loader.Program, stdLib map[string]bool) {
	initialPkgs := make(map[*loader.PackageInfo]bool)
	for _, pkg := range prog.InitialPackages() {
		initialPkgs[pkg] = true
	}
	// For now we'll look at all of the non-stdlib-source files.
	// As an optimization, we could just look at files that have been changed.
	for _, pkg := range prog.AllPackages {
		if stdLib[pkg.String()] || initialPkgs[pkg] {
			continue
		}
		for _, f := range pkg.Files {
			ast.Inspect(f, func(node ast.Node) bool {
				if gen.IsBreakpoint(node) {
					pos := prog.Fset.Position(node.Pos())
					fmt.Printf("godebug %s: Ignoring breakpoint at %s:%d because package %q has not been flagged for instrumentation. See 'godebug help %s'.\n\n",
						subcommand, filepath.Join(pkg.String(), filepath.Base(pos.Filename)), pos.Line, pkg.Pkg.Name(), subcommand)
				}
				return true
			})
		}
	}
}

func markAlmostAllPackages(prog *loader.Program, stdLib map[string]bool) {
	for _, pkg := range prog.AllPackages {
		path := pkg.String()
		switch {

		// skip this package if...
		case stdLib[path]: // it's part of the standard library
		case prog.ImportMap[path] == nil: // it's a Created package
		case strings.HasPrefix(path, "github.com/mailgun/godebug"):
			// it's the godebug library or one of its dependecies

		// otherwise include it
		default:
			prog.Imported[pkg.String()] = pkg
		}
	}
}

func getGoFiles() (gofiles, rest []string) {
	for i, arg := range runFlags.Args() {
		if arg == "--" {
			rest = runFlags.Args()[i+1:]
			break
		}
		if !strings.HasSuffix(arg, ".go") {
			rest = runFlags.Args()[i:]
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

func abs(s string) string {
	res, err := filepath.Abs(s)
	if err != nil {
		logFatal("failed to make output path absolute")
	}
	return res
}

func mapPkgsToTmpDir(pkgs []string) []string {
	result := make([]string, len(pkgs))
	cwd := getwd()
	for i, pkg := range pkgs {
		result[i], _ = findUnderGopath(cwd, pkg)
	}
	return result
}

func findUnderGopath(cwd, pkg string) (string, string) {
	found, err := build.Import(pkg, cwd, build.FindOnly)
	if err != nil {
		logFatalf("Failed to find package %q in findUnderGopath. This is probably a bug -- please report it at https://github.com/mailgun/godebug/issues/new. Thanks!", pkg)
	}
	if found.SrcRoot == "" || found.ImportPath == "" {
		logFatalf("Looks like package %q is not in a GOPATH workspace. godebug doesn't support it right now, but if you open a ticket at https://github.com/mailgun/godebug/issues/new we'll fix it soon. Thanks!", pkg)
	}
	return found.ImportPath, found.Dir
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
	gen.Generate(prog, ioutil.ReadFile, func(importPath, filename string) io.WriteCloser {
		if *w {
			return createFileHook(filename, "")
		}
		return nopCloser{os.Stdout}
	})
}

func parseBuildArguments(args []string) ([]string, bool) {
	if len(args) == 0 {
		return []string{"."}, true
	}
	if len(args) == 1 && !strings.HasSuffix(args[0], ".go") {
		return args, true
	}
	for _, a := range args {
		if !strings.HasSuffix(a, ".go") {
			logFatal("you can only build a set of files or a single package")
		}
	}
	return args, false
}

func isFlag(arg, name string) bool {
	return arg == name || strings.HasPrefix(arg, name+"=")
}

func parseTestArguments(args []string) (packages, otherFlags []string) {
	// format: [-godebugwork] [-instrument pkgs...] [-o=...] [-c] [packages] [testFlags]

	// Find first unrecognized flag.
	sep := len(args)
	for i, arg := range args {
		if strings.HasPrefix(arg, "--") {
			arg = arg[1:]
		}
		if strings.HasPrefix(arg, "-") &&
			!isFlag(arg, "-instrument") &&
			!isFlag(arg, "-godebugwork") &&
			!isFlag(arg, "-tags") &&
			!isFlag(arg, "-o") &&
			!isFlag(arg, "-c") {
			sep = i
			break
		}
	}

	exitIfErr(testFlags.Parse(args[:sep]))
	return testFlags.Args(), args[sep:]
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
	var (
		pkgs   = make(map[string]bool)
		cmd    = exec.Command("go", "list", "std")
		stdout = bytes.NewBuffer(nil)
		stderr = bytes.NewBuffer(nil)
	)
	cmd.Stdout = stdout
	cmd.Stderr = stderr
	if err := cmd.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to identify standard library packages. godebug should still work, but instrumentation might take longer.\nHere's the error from running 'go list std':\n%s\n", stderr.Bytes())
	}
	b := bytes.TrimSpace(stdout.Bytes())
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
