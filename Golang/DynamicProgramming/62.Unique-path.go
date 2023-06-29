package main

func uniquePaths(m int, n int) int {
	dp := make([][]int, m)
	//dp[m][n] = dp[m-1][n] + dp[m][n-1]
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
		for j := 0; j < n; j++ {
			if i == 0 {
				dp[i][j] = 1
			} else if j == 0 {
				dp[i][j] = 1
			} else {
				dp[i][j] = dp[i-1][j] + dp[i][j-1]
			}
		}
	}
	return dp[m-1][n-1]
}
