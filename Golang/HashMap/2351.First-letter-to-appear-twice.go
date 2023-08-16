package main

func repeatedCharacter(s string) byte {
	hashmap := map[byte]int{}
	var result byte
	for i := 0; i < len(s); i++ {
		hashmap[s[i]]++
		if hashmap[s[i]] == 2 {
			result = s[i]
			break
		}
	}
	return result
}
