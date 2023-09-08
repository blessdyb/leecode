package main

func minOperations(nums []int) int {
	ret := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] <= nums[i-1] {
			ret += nums[i-1] + 1 - nums[i]
			nums[i] = nums[i-1] + 1
		}
	}
	return ret
}
