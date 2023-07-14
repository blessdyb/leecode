package main

func isAnagram(s string, t string) bool {
	hashmap := map[rune]int{}
	for _, item := range s {
		if _, ok := hashmap[item]; !ok {
			hashmap[item] = 1
		} else {
			hashmap[item]++
		}
	}
	for _, item := range t {
		if _, ok := hashmap[item]; !ok {
			return false
		} else if hashmap[item] == 0 {
			return false
		} else {
			hashmap[item]--
		}
	}
	return len(s) == len(t)
}

func isAnagramArray(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	hashset := [26]int{}
	for _, item := range s {
		hashset[item-rune('a')]++
	}
	for _, item := range t {
		if hashset[item-rune('a')] == 0 {
			return false
		} else {
			hashset[item-rune('a')]--
		}
	}
	return true
}
