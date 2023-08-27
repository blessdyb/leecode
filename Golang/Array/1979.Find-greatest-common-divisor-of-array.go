package main

import "sort"

func findGCD(nums []int) int {
	sort.Ints(nums)
	ret := 1
	for i := 2; i <= nums[0]; i++ {
		if nums[0]%i == 0 && nums[len(nums)-1]%i == 0 {
			ret = i
		}
	}
	return ret
}
