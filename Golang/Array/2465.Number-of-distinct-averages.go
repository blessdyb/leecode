package main

import "sort"

func distinctAverages(nums []int) int {
	hashmap := map[float32]bool{}
	sort.Ints(nums)
	for len(nums) > 0 {
		avg := float32(nums[0]+nums[len(nums)-1]) / 2.0
		hashmap[avg] = true
		nums = nums[1 : len(nums)-1]
	}
	return len(hashmap)
}
