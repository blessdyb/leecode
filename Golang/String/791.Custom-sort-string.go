package main

import "sort"

func customSortString(order string, s string) string {
	hashmap := map[byte]int{}
	for i := 0; i < len(order); i++ {
		hashmap[order[i]] = i + 1
	}
	chars := []byte(s)
	sort.Slice(chars, func(a, b int) bool {
		return hashmap[chars[a]] < hashmap[chars[b]]
	})
	return string(chars)
}
