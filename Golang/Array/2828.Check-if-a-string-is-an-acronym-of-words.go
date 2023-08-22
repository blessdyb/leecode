package main

func isAcronym(words []string, s string) bool {
	if len(s) != len(words) {
		return false
	}
	for i, word := range words {
		if i >= len(s) || s[i] != word[0] {
			return false
		}
	}
	return true
}
