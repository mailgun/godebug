package main

import "strings"

func rawQuote(s string) string {
	return "`" + strings.Replace(s, "`", "` + \"`\" + `", -1) + "`"
}
