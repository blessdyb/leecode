package main

func makeEqual(words []string) bool {
	hashmap := map[byte]int{}
	for _, word := range words {
		for i := 0; i < len(word); i++ {
			hashmap[word[i]]++
		}
	}
	for _, v := range hashmap {
		if v%len(words) != 0 {
			return false
		}
	}
	return true
}
