package main

import (
	"math"
	"sort"
)

// Sort + two pointers (left end -> <- right end to quickly bypass the invalid cases)
func threeSumClosestBruteN2(nums []int, target int) int {
	sort.Ints(nums)
	ret := 300000
	for i := 0; i < len(nums)-2; i++ {
		left, right := i+1, len(nums)-1
		for left < right {
			total := nums[i] + nums[left] + nums[right]
			if math.Abs(float64(ret-target)) > math.Abs(float64(total-target)) {
				ret = total
			}
			if ret > target {
				right--
			} else {
				left++
			}
		}
	}
	return ret
}

func threeSumClosestBruteForce(nums []int, target int) int {
	distance := 300000
	ret := 0
	for i := 0; i < len(nums)-2; i++ {
		for j := i + 1; j < len(nums)-1; j++ {
			for k := j + 1; k < len(nums); k++ {
				total := int(math.Abs(float64(nums[i] + nums[j] + nums[k] - target)))
				if total < distance {
					ret = nums[i] + nums[j] + nums[k]
					distance = total
				}
			}
		}
	}
	return ret
}
