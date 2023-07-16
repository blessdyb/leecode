package main

func canConstruct(ransomNote string, magazine string) bool {
	hashmap := map[rune]int{}
	for _, c := range magazine {
		hashmap[c]++
	}
	for _, c := range ransomNote {
		if hashmap[c] == 0 {
			return false
		} else {
			hashmap[c]--
		}

	}
	return true
}
