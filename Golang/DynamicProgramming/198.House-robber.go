package main

func rob(nums []int) int {
	dp := make([]int, len(nums)+1)
	dp[0] = 0
	dp[1] = nums[0]
	for i := 2; i <= len(nums); i++ {
		with := dp[i-2] + nums[i-1]
		without := dp[i-1]
		if with > without {
			dp[i] = with
		} else {
			dp[i] = without
		}
	}
	return dp[len(nums)]
}
