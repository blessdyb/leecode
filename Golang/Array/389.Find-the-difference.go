package main

func findTheDifference(s string, t string) byte {
	hashmap := map[byte]int{}
	for i := 0; i < len(s); i++ {
		hashmap[s[i]]++
	}
	var ret byte
	for i := 0; i < len(t); i++ {
		if hashmap[t[i]] == 0 {
			ret = t[i]
			break
		}
		hashmap[t[i]]--
	}
	return ret
}
