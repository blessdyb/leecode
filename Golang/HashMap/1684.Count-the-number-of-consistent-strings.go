package main

func countConsistentStrings(allowed string, words []string) int {
	hashmap := map[rune]bool{}
	for _, a := range allowed {
		hashmap[a] = true
	}
	ret := 0
	for _, word := range words {
		flag := true
		for _, w := range word {
			if !hashmap[w] {
				flag = false
				break
			}
		}
		if flag {
			ret++
		}
	}
	return ret
}
