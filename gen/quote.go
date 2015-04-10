package gen

import "strings"

func rawQuote(s string) string {
	return "`" + strings.Replace(s, "`", "` + \"`\" + `", -1) + "`"
}
