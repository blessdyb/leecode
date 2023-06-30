package main

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func integerBreak(n int) int {
	// dp[n] = max(dp[i] * (n - i), i * (n - i))
	dp := make([]int, n+1)
	dp[0] = 0
	dp[1] = 1
	dp[2] = 1
	for i := 3; i <= n; i++ {
		for j := 1; j < i; j++ {
			dp[i] = max(dp[i], max(dp[i-j]*j, (i-j)*j))
		}
	}
	return dp[n]
}
