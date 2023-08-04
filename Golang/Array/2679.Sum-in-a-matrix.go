package main

import "sort"

func matrixSum(nums [][]int) int {
	for i := 0; i < len(nums); i++ {
		sort.Ints(nums[i])
	}
	ret := 0
	for len(nums[0]) > 0 {
		max := 0
		for i := 0; i < len(nums); i++ {
			if max < nums[i][len(nums[i])-1] {
				max = nums[i][len(nums[i])-1]
			}
			nums[i] = nums[i][:len(nums[i])-1]
		}
		ret += max
	}
	return ret
}
