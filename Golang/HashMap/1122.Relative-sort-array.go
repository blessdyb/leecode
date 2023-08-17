package main

import "sort"

func relativeSortArray(arr1 []int, arr2 []int) []int {
	hashmap := map[int]int{}
	for idx, num := range arr2 {
		hashmap[num] = idx - len(arr2)
	}
	sort.Slice(arr1, func(i, j int) bool {
		if hashmap[arr1[i]] == 0 && hashmap[arr1[j]] == 0 {
			return arr1[i] < arr1[j]
		}
		return hashmap[arr1[i]] < hashmap[arr1[j]]
	})
	return arr1
}
