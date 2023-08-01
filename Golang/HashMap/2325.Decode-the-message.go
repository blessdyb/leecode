package main

func decodeMessage(key string, message string) string {
	hashmap := map[rune]rune{}
	start := 'a'
	for _, s := range key {
		if _, ok := hashmap[s]; !ok && s != ' ' {
			hashmap[s] = start
			start++
		}
	}
	result := []rune{}
	for _, m := range message {
		if m == ' ' {
			result = append(result, m)
		} else {
			result = append(result, hashmap[m])
		}
	}
	return string(result)
}
