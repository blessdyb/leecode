package main

import "sort"

func uniqueOccurrences(arr []int) bool {
	hashmap := map[int]int{}
	for _, num := range arr {
		hashmap[num]++
	}
	occurrencies := []int{}
	for _, v := range hashmap {
		occurrencies = append(occurrencies, v)
	}
	sort.Ints(occurrencies)
	for i := 1; i < len(occurrencies); i++ {
		if occurrencies[i] == occurrencies[i-1] {
			return false
		}
	}
	return true
}
