package main

func leftRightDifference(nums []int) []int {
	runningSum := make([]int, len(nums)+1)
	runningSum[0] = 0
	for i := 1; i <= len(nums); i++ {
		runningSum[i] = runningSum[i-1] + nums[i-1]
	}
	// leftSum is runningSum[:len(nums)]
	// rightSum[i] = runningSum[len(nums)] - runningSum[i] - nums[i]
	ret := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		ret[i] = runningSum[i] - (runningSum[len(nums)] - runningSum[i] - nums[i])
		if ret[i] < 0 {
			ret[i] *= -1
		}
	}
	return ret
}
