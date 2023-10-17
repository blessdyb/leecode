package main

func minStartValue(nums []int) int {
	preSum := make([]int, len(nums))
	preSum[0] = nums[0]
	min := preSum[0]
	for i := 1; i < len(nums); i++ {
		preSum[i] = preSum[i-1] + nums[i]
		if min > preSum[i] {
			min = preSum[i]
		}
	}
	if min >= 0 {
		return 1
	}
	return 1 - min
}
