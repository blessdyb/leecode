package main

func maxSubArray(nums []int) int {
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	dp := make([]int, len(nums))
	dp[0] = 0
	ret := dp[0]
	for i := 1; i < len(nums); i++ {
		// dp[i], the max subarray ends with i
		dp[i] = max(dp[i-1]+nums[i], nums[i])
		ret = max(ret, dp[i])
	}
	return ret
}
