package main

func reversePrefix(word string, ch byte) string {
	chars := []byte(word)
	for i := 0; i < len(chars); i++ {
		if chars[i] == ch {
			for j, k := 0, i; j < k; {
				chars[j], chars[k] = chars[k], chars[j]
				j++
				k--
			}
			break
		}
	}
	return string(chars)
}
