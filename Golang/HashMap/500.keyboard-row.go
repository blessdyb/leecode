package main

func findWords(words []string) []string {
	ret := []string{}
	hashmap := map[rune]int{}
	// Note: uppercase && lowercase together
	for index, row := range []string{"qwertyuiopQWERTYUIOP", "asdfghjklASDFGHJKL", "zxcvbnmZXCVBNM"} {
		for _, k := range row {
			hashmap[k] = index
		}
	}
	for _, word := range words {
		prev := -1
		flag := true
		for _, w := range word {
			if prev == -1 {
				prev = hashmap[w]
			} else if prev != hashmap[w] {
				flag = false
				break
			}
		}
		if flag {
			ret = append(ret, word)
		}
	}
	return ret
}
