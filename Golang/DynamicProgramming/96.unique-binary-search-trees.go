package main

func numTrees(n int) int {
	// dp[0] = 1
	// dp[1] = 1
	// dp[2] = 2
	// dp[3] = dp[2] * dp[0] + dp[1] * dp[1] + dp[0] * dp[2]
	// dp[n] += dp[n - i] * dp[i - 1]
	dp := make([]int, n+1)
	dp[0] = 1
	for i := 1; i <= n; i++ {
		for j := 1; j <= i; j++ {
			dp[i] += dp[j-1] * dp[i-j]
		}
	}
	return dp[n]
}
