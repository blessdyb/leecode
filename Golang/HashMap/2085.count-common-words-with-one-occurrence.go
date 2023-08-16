package main

func countWords(words1 []string, words2 []string) int {
	hashmap1 := map[string]int{}
	hashmap2 := map[string]int{}
	hashmap3 := map[string]bool{}
	for _, w := range words1 {
		hashmap1[w]++
		hashmap3[w] = true
	}
	for _, w := range words2 {
		hashmap2[w]++
		hashmap3[w] = true
	}
	ret := 0
	for k, _ := range hashmap3 {
		if hashmap1[k] == 1 && hashmap2[k] == 1 {
			ret++
		}
	}
	return ret
}
