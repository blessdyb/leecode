package main

func canBeTypedWords(text string, brokenLetters string) int {
	hashmap := map[byte]bool{}
	for i := 0; i < len(brokenLetters); i++ {
		hashmap[brokenLetters[i]] = true
	}
	ret, start, end := 0, 0, 0
	for end < len(text) {
		// Jump to next word
		if hashmap[text[end]] {
			for end < len(text) && text[end] == ' ' {
				end++
			}
			start = end + 1
		} else if text[end] == ' ' {
			ret++
		}
		end++
	}
	// last word edge case
	if start != end {
		ret++
	}
	return ret
}
