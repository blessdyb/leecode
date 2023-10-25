package main

import "math"

func findIndices(nums []int, indexDifference int, valueDifference int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + indexDifference; j < len(nums); j++ {
			if int(math.Abs(float64(nums[i]-nums[j]))) >= valueDifference {
				return []int{i, j}
			}
		}
	}
	return []int{-1, -1}
}
