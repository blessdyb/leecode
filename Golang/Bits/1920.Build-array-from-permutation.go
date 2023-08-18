package main

func buildArray(nums []int) []int {
	ret := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		ret[i] = nums[nums[i]]
	}
	return ret
}

// O(1) space solution
// 1. We need to store two values in one space
// 2. We use bit calculation to store new value to high end

func buildArrayO1(nums []int) []int {
	n := len(nums)
	mask := (1 << 10) - 1
	for i := 0; i < n; i++ {
		nums[i] |= (nums[nums[i]] & mask) << 10
	}
	for i := 0; i < n; i++ {
		nums[i] = nums[i] >> 10
	}
	return nums
}
