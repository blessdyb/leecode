package main

func minimizedStringLength(s string) int {
	hashmap := map[byte]int{}
	for i := 0; i < len(s); i++ {
		hashmap[s[i]]++
		if hashmap[s[i]] >= 3 {
			hashmap[s[i]] -= 2
		} else if hashmap[s[i]] == 2 {
			hashmap[s[i]]--
		}
		if hashmap[s[i]] == 0 {
			delete(hashmap, s[i])
		}
	}
	ret := 0
	for _, v := range hashmap {
		ret += v
	}
	return ret
}
