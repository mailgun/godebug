package subfoo

import "github.com/mailgun/godebug/lib"

func SubFoo() {
	_ = "in subfoo"
}

func HasBreakpoint() {
	godebug.SetTrace()
	_ = "breakpoint"
}
