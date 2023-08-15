package main

import "sort"

func mergeSimilarItems(items1 [][]int, items2 [][]int) [][]int {
	hashmap := map[int]int{}
	for _, item := range items1 {
		hashmap[item[0]] += item[1]
	}
	for _, item := range items2 {
		hashmap[item[0]] += item[1]
	}
	sm := [][]int{}
	for k, v := range hashmap {
		sm = append(sm, []int{k, v})
	}
	sort.Slice(sm, func(i, j int) bool {
		return sm[i][0] < sm[j][0]
	})
	return sm
}
