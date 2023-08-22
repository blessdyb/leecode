package main

import "math"

func findDisappearedNumbers(nums []int) []int {
	for i := 0; i < len(nums); i++ {
		idx := int(math.Abs(nums[i])) - 1
		if nums[idx] > 0 {
			nums[idx] *= -1
		}
	}
	ret := []int{}
	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			ret = append(ret, i+1)
		}
	}
	return ret
}
