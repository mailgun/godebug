package main

import (
	"fmt"

	"github.com/bradfitz/iter"
)

func main() {
	x := mul(2, 2)
	fmt.Println(mul(x, x))
}

func add(n, m int) int {
	if n == 0 {
		return m
	}
	if m == 0 {
		return n
	}
	return n + m
}

func mul(n, m int) int {
	var r int
	for range iter.N(m) {
		r += m
	}
	return r
}
