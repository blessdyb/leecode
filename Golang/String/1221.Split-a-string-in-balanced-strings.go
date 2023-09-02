package main

func balancedStringSplit(s string) int {
	count := 0
	hashmap := map[byte]int{
		'L': 0,
		'R': 0,
	}
	for i := 0; i < len(s); i++ {
		hashmap[s[i]]++
		if hashmap['L'] == hashmap['R'] {
			count++
			hashmap = map[byte]int{
				'L': 0,
				'R': 0,
			}
		}
	}
	return count
}
