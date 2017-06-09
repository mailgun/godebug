### DEPRECATED! There will be no further development. Please use https://github.com/derekparker/delve. But if you want to keep the project going and ready to become its maintaner please contact us and we can make you one.


godebug

-------
[![Linux Build Status](https://img.shields.io/travis/mailgun/godebug/master.svg?label=linux)](https://travis-ci.org/mailgun/godebug)
[![Windows Build Status](https://img.shields.io/appveyor/ci/jeremyschlatter/godebug/master.svg?label=windows)](https://ci.appveyor.com/project/jeremyschlatter/godebug/branch/master)


A cross-platform debugger for Go.

### How?

`godebug` uses source code generation to instrument your program with debugging calls. [go tool cover](http://blog.golang.org/cover) takes a similar approach to code coverage. When you run `godebug`, it parses your program, instruments function calls, variable declarations, and statement lines, outputs the resulting code somewhere, and runs/compiles it. When this modified code runs, it stops at breakpoints and lets you step through the program and inspect variables.

For more detail, see the [end of this README](#how-it-works-more-detail).

### Status

`godebug` is still very new. [File an issue](https://github.com/mailgun/godebug/issues/new) or send me an email if you find any rough edges:

![contact](https://s3.amazonaws.com/f.cl.ly/items/1d0i0W2e3F1K0L3K0Y1N/contact.png)

### Installation:

    $ go get github.com/mailgun/godebug

### Getting started:

Insert a breakpoint anywhere in a source file you want to debug:

    _ = "breakpoint"

If the breakpoint is in package main and you don't want to examine any imported packages, you can just run:

    $ godebug run gofiles... [arguments...]

If you want to trace the program outside of the main package, list the packages to trace in the `-instrument` flag:

    $ godebug run -instrument=pkg1,pkg2,pkg3 gofiles... [arguments...]

If you are debugging a test, use 'godebug test' like you would use 'go test':

    $ godebug test [-instrument pkgs...] [packages]

Finally, you can [cross-]compile a debugging binary using 'godebug build':

    $ godebug build [-instrument pkgs...] [-o output] [package]

The compiled binary has no dependencies, so you can build it locally and then debug on i.e. a staging server.

That's it. See 'godebug help' for the full usage.

### Debugger commands:

The current commands are:

command              | result
---------------------|------------------------
h(elp)               | show help message
n(ext)               | run the next line
s(tep)               | run for one step
c(ontinue)           | run until the next breakpoint
l(ist)               | show the current line in context of the code around it
p(rint) [expression] | print a variable or any other Go expression
q(uit)               | exit the program

### Caveats

It is not currently possible to step into standard library packages. (Issue [#12](https://github.com/mailgun/godebug/issues/12))

### How it works (more detail)

Consider this program:

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, world!")
}
```

Now let's modify it a bit:

```go
package main

import (
    "bufio"
    "fmt"
    "os"
)

func main() {
    fmt.Println(`-> fmt.Println("Hello, world!")`)
    bufio.NewScanner(os.Stdin).Scan()
    fmt.Println("Hello, world!")
}
```

When we run this modified version, we see:

    -> fmt.Println("Hello, world!")

And then the program waits for input before proceeding.

We have just implemented a debugger for the first program! It may not seem like much, but this program implements two fundamental debugger behaviors: (1) display the current state of the program, and (2) do not proceed until instructed by the user. Furthermore, the changes we made were straightforward and easy to automate:

  * insert import statements for `bufio` and `os`, if not already present.
  * in `main()`, insert the statement `fmt.Println(<quote next line>)`
  * in `main()`, insert the statement `bufio.NewScanner(os.Stdin).Scan()`.

We could do exactly the same thing for any other program with a single-line main function. And it's not hard to see how to generalize this to multiple lines. This, in essence, is what godebug does. Parse source code, insert extra code that implements the behavior of a debugger for that program, output and run the result. godebug handles many more cases than this simple example and implements more interesting debugging behavior, but the principle is exactly the same.
