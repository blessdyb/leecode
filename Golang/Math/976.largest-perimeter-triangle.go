package main

import "sort"

func largestPerimeter(nums []int) int {
	ret := 0
	sort.Ints(nums)
	for i := len(nums) - 1; i >= 2; i-- {
		if nums[i-2]+nums[i-1] > nums[i] {
			sum := nums[i-2] + nums[i-1] + nums[i]
			if sum > ret {
				ret = sum
			}
		}
	}
	return ret
}
