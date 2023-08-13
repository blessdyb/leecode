package main

func maximumNumberOfStringPairs(words []string) int {
	reverse := func(str string) string {
		newStr := ""
		for i := len(str) - 1; i >= 0; i-- {
			newStr += string(str[i])
		}
		return newStr
	}
	hashmap := map[string]int{}
	ret := 0
	for _, word := range words {
		hashmap[word]++
	}
	for _, word := range words {
		reversed := reverse(word)
		if hashmap[word] > 0 && hashmap[reversed] > 0 {
			hashmap[word]--
			hashmap[reversed]--
			if hashmap[word] >= 0 && hashmap[reversed] >= 0 {
				ret++
			}
		}
	}
	return ret
}
