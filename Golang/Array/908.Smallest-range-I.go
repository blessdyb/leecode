package main

import "sort"

func smallestRangeI(nums []int, k int) int {
	sort.Ints(nums)
	if nums[len(nums)-1]-nums[0] >= 2*k {
		return nums[len(nums)-1] - nums[0] - 2*k
	}
	return 0
}
