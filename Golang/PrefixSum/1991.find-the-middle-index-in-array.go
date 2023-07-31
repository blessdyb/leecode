package main

func findMiddleIndex(nums []int) int {
	length := len(nums)
	preSum := make([]int, length+1)
	for i := 1; i <= length; i++ {
		preSum[i] = preSum[i-1] + nums[i-1]
	}
	for i := 0; i < length; i++ {
		if preSum[i] == preSum[length]-preSum[i+1] {
			return i
		}
	}
	return -1
}
