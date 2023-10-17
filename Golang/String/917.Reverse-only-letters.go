package main

func reverseOnlyLetters(s string) string {
	bytes := []byte(s)
	isEnglishLetter := func(c byte) bool {
		return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z')
	}
	for i, j := 0, len(bytes)-1; i < j; {
		if isEnglishLetter(bytes[i]) && isEnglishLetter(bytes[j]) {
			bytes[i], bytes[j] = bytes[j], bytes[i]
			i++
			j--
		} else if isEnglishLetter(bytes[i]) {
			j--
		} else if isEnglishLetter(bytes[j]) {
			i++
		} else {
			i++
			j--
		}
	}
	return string(bytes)
}
