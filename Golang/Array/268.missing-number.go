package main

import "sort"

func missingNumberSort(nums []int) int {
	sort.Ints(nums)
	for i := 0; i < len(nums); i++ {
		if nums[i] != i {
			return i
		}
	}
	return len(nums)
}

func missingNumberO1(nums []int) int {
	sum := len(nums) * (len(nums) + 1) / 2
	for i := 0; i < len(nums); i++ {
		sum -= nums[i]
	}
	return sum
}
