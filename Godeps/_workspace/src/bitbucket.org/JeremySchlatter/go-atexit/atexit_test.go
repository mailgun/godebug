package atexit

import "fmt"

func Example() {
	defer CallExitFuncs()
	Run(func() {
		fmt.Println("End")
	})
	Run(func() {
		fmt.Println("Middle")
	})
	fmt.Println("Beginning")
	// Output:
	// Beginning
	// Middle
	// End
}
