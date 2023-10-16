package main

import "sort"

func maximunProduct(nums []int) int {
	sort.Ints(nums)
	a := nums[len(nums)-1] * nums[len(nums)-2] * nums[len(nums)-3]
	b := nums[0] * nums[1] * nums[len(nums)-1]
	if a > b {
		return a
	}
	return b
}
