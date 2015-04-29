package main

import (
	"os"
)

func main() {
	if o := os.Getenv("O"); o != "" {
		println(o)
	} else if s := os.Getenv("S"); s != "" {
		println(s)
	}
}
