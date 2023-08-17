package main

func countCharacters(words []string, chars string) int {
	hashmap := map[byte]int{}
	for i := 0; i < len(chars); i++ {
		hashmap[chars[i]]++
	}
	ret := 0
	for _, w := range words {
		flag := true
		temp := map[byte]int{}
		for i := 0; i < len(w); i++ {
			if hashmap[w[i]] == 0 {
				flag = false
				break
			} else {
				temp[w[i]]++
				if hashmap[w[i]] < temp[w[i]] {
					flag = false
					break
				}
			}
		}
		if flag {
			ret += len(w)
		}
	}
	return ret
}
