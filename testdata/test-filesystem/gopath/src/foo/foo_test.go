package foo

import (
	"testing"

	"github.com/mailgun/godebug/lib"
)

func TestFoo(t *testing.T) {
	godebug.SetTrace()
	t.Fail()
}
