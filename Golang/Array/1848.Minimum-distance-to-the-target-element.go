package main

import "math"

func getMinDistance(nums []int, target int, start int) int {
	ret := len(nums) * 10
	for i := 0; i < len(nums); i++ {
		if nums[i] == target {
			absDiff := int(math.Abs(float64(i - start)))
			if absDiff < ret {
				ret = absDiff
			}
		}
	}
	return ret
}
