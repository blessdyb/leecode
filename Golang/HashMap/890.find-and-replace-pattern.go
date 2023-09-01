package main

func findAndReplacePattern(words []string, pattern string) []string {
	ret := []string{}
	match := func(a, b string) bool {
		hashmap := map[byte]byte{}
		flag := true
		for i := 0; i < len(a); i++ {
			if _, ok := hashmap[b[i]]; !ok {
				hashmap[b[i]] = a[i]
			}
			if hashmap[b[i]] != a[i] {
				flag = false
				break
			}
		}
		return flag
	}
	for _, word := range words {
		if match(word, pattern) && match(pattern, word) {
			ret = append(ret, word)
		}
	}
	return ret
}
