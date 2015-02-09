godebug
-------

A debugger for Go.

### How it works

`godebug` uses source code generation to instrument your program with debugging calls. [go tool cover](http://blog.golang.org/cover) takes a similar approach to code coverage. When you run `godebug -w`, it parses your program, instruments function calls, variable declarations, and statement lines, and outputs the resulting code somewhere (currently either stdout or  in place over your original files). When you run this modified code, assuming you put a breakpoint somewhere, you can step through it and inspect variables. Coming later: evaluate arbitrary Go expressions and write to variables.


### Status

`godebug` is currently in alpha stage -- expect some problems.


### Installation:

    $ go get github.com/mailgun/godebug


### Use:

First, get your directory in a clean state. **The command below will overwrite your files, so make sure you have committed or stashed everything.**

In any file where you want a breakpoint, import `github.com/mailgun/godebug/lib` and insert this breakpoint anywhere in the code: `godebug.SetTrace()`. Then run:

    $ godebug -w .

Your code is now self-debugging. Congrats! Run it with `go run`, test it with `go test`, or build then run it with `go build`.


### Debugger commands:

The debugger is currently extremely limited. The commands are:

command       | result
--------------|------------------------
n(ext)        | run until the next line
s(tep)        | run for one step
p(rint) [var] | print a variable

The debugger will attempt to interpret any text that does not match the above commands as a variable name. If that variable exists, the debugger will print it.
