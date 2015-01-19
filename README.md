godebug
-------

`godebug` is a code-generation-based debugger for Go. It is currently in alpha stage -- expect lots of problems.

Installation:

    $ go get github.com/mailgun/godebug


Use:

First, get your directory in a clean state. **The command below will overwrite your files, so make sure you have committed or stashed everything.**

Insert the following text anywhere in your code where you want a breakpoint: `godebug.SetTrace()` Then run:

    $ godebug -w .

Your code is now self-debugging. Congrats! Run it with `go run`, test it with `go test`, or build then run it with `go build`.


Debugger commands:

The debugger is currently extremely limited. The commands are:

n(ext): run until the next line
s(tep): run for one step
p(rint) <var>: print a variable

The debugger will attempt to interpret any text that does not match the above commands as a variable name. If that variable exists, the debugger will print it.
