package main

import "sort"

func mergeArrays(nums1 [][]int, nums2 [][]int) [][]int {
	hashmap := map[int]int{}
	for _, num := range nums1 {
		hashmap[num[0]] = num[1]
	}
	for _, num := range nums2 {
		hashmap[num[0]] += num[1]
	}
	ret := [][]int{}
	for k, v := range hashmap {
		ret = append(ret, []int{k, v})
	}
	sort.Slice(ret, func(i, j int) bool {
		return ret[i][0] < ret[j][0]
	})
	return ret
}
