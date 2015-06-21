package main

import (
	"fmt"
)

func main() {
	func() {
		var (
			x   int
			y   int
			err error
		)
		x = 2
		y = 3
		err = nil
		if err != nil {
			fmt.Printf("%d\n", x+y)
		}
	}()
}
