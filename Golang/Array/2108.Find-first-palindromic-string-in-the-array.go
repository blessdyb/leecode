package main

func firstPalindrome(words []string) string {
	ret := ""
	for _, word := range words {
		flag := true
		for i, j := 0, len(word)-1; i < j; {
			if word[i] != word[j] {
				flag = false
				break
			}
			i++
			j--
		}
		if flag {
			ret = word
			break
		}
	}
	return ret
}
