package eval

// Wraps panic message in "eval: %s. Did you type check?"
func dytc(probablyNot string) string {
	return "eval: " + probablyNot + ". Did you type check?"
}
