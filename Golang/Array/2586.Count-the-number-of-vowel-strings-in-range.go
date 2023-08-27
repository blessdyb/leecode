package main

func vowelStrings(words []string, left int, right int) int {
	hashmap := map[byte]bool{
		'a': true,
		'e': true,
		'i': true,
		'o': true,
		'u': true,
	}
	ret := 0
	for i := left; i <= right && i < len(words); i++ {
		word := words[i]
		if hashmap[word[0]] && hashmap[word[len(word)-1]] {
			ret++
		}
	}
	return ret
}
