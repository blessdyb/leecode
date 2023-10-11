package main

func longestPalindrome(s string) int {
	hashmap := map[byte]int{}
	for i := 0; i < len(s); i++ {
		hashmap[s[i]]++
	}
	ret := 0
	for k, v := range hashmap {
		if v%2 == 0 {
			ret += v
			delete(hashmap, k)
		} else {
			ret += v - 1
		}
	}
	if len(hashmap) > 0 {
		ret++
	}
	return ret
}
