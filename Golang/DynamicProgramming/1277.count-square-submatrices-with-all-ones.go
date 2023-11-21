func countSquares(matrix [][]int) int {
	m, n := len(matrix), len(matrix[0])
	dp := make([][]int, m)
	ret := 0
	for i := range dp {
		dp[i] = make([]int, n)
		dp[i][0] = matrix[i][0]
		ret += dp[i][0]
	}
	for j := 1; j < n; j++ {
		dp[0][j] = matrix[0][j]
		ret += dp[0][j]
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if matrix[i][j] == 1 {
				dp[i][j] = min(dp[i-1][j-1], min(dp[i-1][j], dp[i][j-1])) + 1
				ret += dp[i][j]
			}
		}
	}
	return ret
}