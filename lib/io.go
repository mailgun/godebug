// +build !js

package godebug

import (
	"bufio"
	"os"
)

var input = bufio.NewScanner(os.Stdin)

var outputW = os.Stdout
