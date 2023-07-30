package main

import "sort"

func targetIndices(nums []int, target int) []int {
	sort.Ints(nums)
	ret := []int{}
	for i := 0; i < len(nums); i++ {
		if nums[i] == target {
			ret = append(ret, i)
		}
	}
	return ret
}
