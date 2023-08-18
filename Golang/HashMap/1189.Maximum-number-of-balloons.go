package main

func maxNumberOfBalloons(text string) int {
	hashmap := map[byte]int{}
	for i := 0; i < len(text); i++ {
		hashmap[text[i]]++
	}
	ret := 0
	for hashmap['b'] >= 1 && hashmap['a'] >= 1 && hashmap['l'] >= 2 && hashmap['o'] >= 2 && hashmap['n'] >= 1 {
		ret++
		hashmap['b'] -= 1
		hashmap['a'] -= 1
		hashmap['l'] -= 2
		hashmap['o'] -= 2
		hashmap['n'] -= 1
	}
	return ret
}
