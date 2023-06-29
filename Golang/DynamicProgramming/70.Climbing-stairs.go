package main

func climbStairs(n int) int {
	dp := []int{0, 1, 2}
	// dp[n] = dp[n-1] + dp[n-2]
	for i := 3; i <= n; i++ {
		dp = append(dp, dp[i-1]+dp[i-2])
	}
	return dp[n]
}
