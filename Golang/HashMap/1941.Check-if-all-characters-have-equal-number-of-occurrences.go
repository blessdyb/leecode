package main

func areOccurrencesEqual(s string) bool {
	hashmap := map[byte]int{}
	for i := 0; i < len(s); i++ {
		hashmap[s[i]]++
	}
	val := hashmap[s[0]]
	for _, v := range hashmap {
		if val != v {
			return false
		}
	}
	return true
}
