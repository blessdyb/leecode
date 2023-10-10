package main

func reverseVowels(s string) string {
	chars := []byte(s)
	for i, j := 0, len(s)-1; i < j; {
		if chars[i] != 'a' && chars[i] != 'e' && chars[i] != 'i' && chars[i] != 'o' && chars[i] != 'u' && chars[i] != 'A' && chars[i] != 'E' && chars[i] != 'I' && chars[i] != 'O' && chars[i] != 'U' {
			i++
		} else if chars[j] != 'a' && chars[j] != 'e' && chars[j] != 'i' && chars[j] != 'o' && chars[j] != 'u' && chars[j] != 'A' && chars[j] != 'E' && chars[j] != 'I' && chars[j] != 'O' && chars[j] != 'U' {
			j--
		} else {
			chars[i], chars[j] = chars[j], chars[i]
			i++
			j--
		}
	}
	return string(chars)
}
