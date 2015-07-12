package main

import (
	"os"
	"fmt"
	"github.com/0xfaded/go-testgen"
)

func main() {
	if err := testgen.Generate(&Test{}, os.Stdout); err != nil {
		panic(fmt.Sprintf("Test generation failed %v\n", err))
	}
}
