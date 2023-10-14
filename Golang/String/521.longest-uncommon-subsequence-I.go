package main

func findLUSlength(a string, b string) int {
	if len(a) > len(b) {
		a, b = b, a
	}
	if a == b {
		return -1
	}
	return len(b)
}
