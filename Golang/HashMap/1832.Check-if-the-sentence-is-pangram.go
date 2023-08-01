package main

func checkIfPangram(sentence string) bool {
	hashmap := map[rune]int{}
	for i := 'a'; i <= 'z'; i++ {
		hashmap[i] = 0
	}
	for _, s := range sentence {
		hashmap[s]++
	}
	for _, v := range hashmap {
		if v == 0 {
			return false
		}
	}
	return true
}
