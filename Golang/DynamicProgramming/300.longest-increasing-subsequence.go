package main

func lengthOfLIS(nums []int) int {
	var max = func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	length := len(nums)
	ret := 1
	dp := make([]int, length)
	for i := 0; i < length; i++ {
		dp[i] = 1
	}
	for i := 1; i < length; i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], dp[j]+1)
			}
			if ret < dp[i] {
				ret = dp[i]
			}
		}
	}
	return ret
}
