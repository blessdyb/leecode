package main

func detectCapitalUse(word string) bool {
	result := 0
	for i := 1; i < len(word); i++ {
		if word[i] >= 'A' && word[i] <= 'Z' {
			result |= 1
		} else {
			result |= 2
		}
	}
	if word[0] >= 'A' && word[0] <= 'Z' {
		return result != 3
	}
	return result == 2 || result == 0
}
