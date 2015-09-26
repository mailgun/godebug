// +build !js

package godebug

import (
	"fmt"
	"strings"

	"github.com/mailgun/godebug/Godeps/_workspace/src/github.com/peterh/liner"
)

var (
	line              *liner.State
	origMode, rawMode liner.ModeApplier
)

func init() {
	// There are three paths to pay attention to here:
	//  1. Normal, happy path
	//	readline initialization succeeds, godebug uses it as intended.
	//  2. Failure path
	//	readline initialization fails. We fall back to reading stdin/stdout,
	//	which works but does not support readline features.
	//  3. Test path
	//	readline initialization would fail in tests if we tried to do it
	//	the normal way. If we just accepted that and used the above failure path,
	//	we would not test the readline library at all. Instead, we do an
	//	abbreviated initialization, which lets us test the library in a way that
	//	is at least somewhat similar to the normal path.
	if buildMode == "test" {
		fmt.Println("godebug: test mode build")
		line = liner.NewLiner()
		promptUser = promptUserReadline
		return
	}

	var err error
	origMode, err = liner.TerminalMode()
	if err != nil {
		return
	}
	line = liner.NewLiner()
	rawMode, err = liner.TerminalMode()
	if err != nil {
		line.Close()
		return
	}
	checkReadlineErr(origMode.ApplyMode())
	promptUser = promptUserReadline
}

var stopBugging = false

func checkReadlineErr(err error) {
	if err != nil && !stopBugging {
		fmt.Println("\nWhoops! You found a godebug issue. Could you report it at https://github.com/mailgun/godebug/issues/new ?\nWe failed to adjust the terminal mode because of this error:", err)
		stopBugging = true
	}
}

func promptUserReadline() (response string, ok bool) {
	if buildMode != "test" {
		checkReadlineErr(rawMode.ApplyMode())
		defer func() {
			checkReadlineErr(origMode.ApplyMode())
		}()
	}
	s, err := line.Prompt("(godebug) ")
	if err != nil {
		fmt.Println("readline error:", err)
		return "", false
	}
	if strings.TrimSpace(s) != "" {
		line.AppendHistory(s)
	}
	return s, true
}
