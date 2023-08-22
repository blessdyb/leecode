package main

import "sort"

func maximizeSum(nums []int, k int) int {
	sort.Ints(nums)
	ret := 0
	for i := 0; i < k; i++ {
		ret += nums[len(nums)-1] + i
	}
	return ret
}
