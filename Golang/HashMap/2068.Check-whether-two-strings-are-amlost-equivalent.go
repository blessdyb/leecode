package main

func checkAlmostEquivalent(word1 string, word2 string) bool {
	h1, h2 := map[byte]int{}, map[byte]bool{}
	for i := 0; i < len(word1); i++ {
		h1[word1[i]]++
	}
	for i := 0; i < len(word2); i++ {
		if h1[word2[i]] == 0 {
			h2[word2[i]] = true
		}
		if h2[word2[i]] {
			h1[word2[i]]++
		} else {
			h1[word2[i]]--
		}
	}
	for _, v := range h1 {
		if v > 3 {
			return false
		}
	}
	return true
}
