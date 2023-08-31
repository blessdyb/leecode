package main

func removeAnagrams(words []string) []string {
	isAnagram := func(a, b string) bool {
		hashmap := map[byte]int{}
		for j := 0; j < len(a); j++ {
			hashmap[a[j]]++
		}
		flag := true
		for k := 0; k < len(b); k++ {
			if hashmap[b[k]] == 0 {
				flag = false
				break
			}
			hashmap[b[k]]--
		}
		return flag
	}
	for i := 1; i < len(words); {
		if isAnagram(words[i], words[i-1]) && isAnagram(words[i-1], words[i]) {
			words = append(words[:i], words[i+1:]...)
		} else {
			i++
		}
	}
	return words
}
