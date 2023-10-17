package main

func isAlienSorted(words []string, order string) bool {
	hashmap := map[byte]int{}
	for i := 0; i < len(order); i++ {
		hashmap[order[i]] = i
	}
	for i := 1; i < len(words); i++ {
		m, n := 0, 0
		for m < len(words[i-1]) && n < len(words[i]) {
			if hashmap[words[i-1][m]] > hashmap[words[i][n]] {
				return false
			} else if hashmap[words[i-1][m]] < hashmap[words[i][n]] {
				break
			}
			m++
			n++
		}
		if m < len(words[i-1]) && n == len(words[i]) {
			return false
		}
	}
	return true
}
