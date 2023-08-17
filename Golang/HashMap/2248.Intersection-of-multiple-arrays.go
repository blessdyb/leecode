package main

import "sort"

func intersection(nums [][]int) []int {
	hashmap := map[int]int{}
	ret := []int{}
	for _, row := range nums {
		for _, item := range row {
			hashmap[item]++
			if hashmap[item] == len(nums) {
				ret = append(ret, item)
			}
		}
	}
	sort.Ints(ret)
	return ret
}
