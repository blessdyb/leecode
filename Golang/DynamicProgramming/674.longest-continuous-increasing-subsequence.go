package main

func findLengthOfLCISDP(nums []int) int {
	length := len(nums)
	dp := make([]int, length)
	for i := 0; i < length; i++ {
		dp[i] = 1
	}
	ret := 1
	for i := 1; i < length; i++ {
		if nums[i] > nums[i-1] {
			dp[i] = dp[i-1] + 1
		}
		if ret < dp[i] {
			ret = dp[i]
		}
	}
	return ret
}

func findLengthOfLCIS(nums []int) int {
	ret := 1
	count := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			count++
		} else {
			count = 1
		}
		if ret < count {
			ret = count
		}
	}
	return ret
}
