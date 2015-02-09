package main

func main() {
	for s := range []string{"foo"} {
		_ = s
	}
}
