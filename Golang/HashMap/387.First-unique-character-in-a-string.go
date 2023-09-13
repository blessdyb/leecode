package main

func firstUniqChar(s string) int {
	hashmap := map[byte]int{}
	for i := 0; i < len(s); i++ {
		hashmap[s[i]]++
	}
	for i := 0; i < len(s); i++ {
		if hashmap[s[i]] == 1 {
			return i
		}
	}
	return -1
}
