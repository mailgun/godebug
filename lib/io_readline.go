// +build !js

package godebug

import (
	"fmt"

	"github.com/mailgun/godebug/Godeps/_workspace/src/github.com/peterh/liner"
)

var (
	line              *liner.State
	origMode, rawMode liner.ModeApplier
)

func init() {
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
	checkReadlineErr(rawMode.ApplyMode())
	defer func() {
		checkReadlineErr(origMode.ApplyMode())
	}()
	s, err := line.Prompt("(godebug) ")
	if err != nil {
		fmt.Println("readline error:", err)
		return "", false
	}
	line.AppendHistory(s)
	return s, true
}
