package main

func isIsomorphic(s string, t string) bool {
	hashmap := map[byte]byte{}
	for i := 0; i < len(s); i++ {
		if v, ok := hashmap[s[i]]; ok {
			if v != t[i] {
				return false
			}
		} else {
			hashmap[s[i]] = t[i]
		}
	}
	hashmap = map[byte]byte{}
	for i := 0; i < len(t); i++ {
		if v, ok := hashmap[t[i]]; ok {
			if v != s[i] {
				return false
			}
		} else {
			hashmap[t[i]] = s[i]
		}
	}
	return true
}
