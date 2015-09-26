// +build go1.5

package main

// Prior to Go 1.5, the go tool would not rebuild stale package outside of the
// current GOPATH directory. This is a problem for godebug, because it creates
// temporary instrumented versions of packages in a separate directory from the
// user's regular GOPATH. For versions of Go before 1.5, godebug needs to
// work around this.
//
// See https://github.com/golang/go/issues/10509
const doGopathWorkaround = false

// New syntax for -X flag.
const buildModeFlag = "-X github.com/mailgun/godebug/lib.buildMode=test"
