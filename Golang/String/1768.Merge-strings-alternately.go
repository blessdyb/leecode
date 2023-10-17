package main

func mergeAlternately(word1 string, word2 string) string {
	ret := ""
	for i, j := 0, 0; i < len(word1) || j < len(word2); {
		if i < len(word1) {
			ret += string(word1[i])
			i++
		}
		if j < len(word2) {
			ret += string(word2[j])
			j++
		}
	}
	return ret
}
