package main

func longestPalindromeSubseqDP(s string) int {
	length := len(s)
	dp := make([][]int, length)
	for i := 0; i < length; i++ {
		dp[i] = make([]int, length)
		dp[i][i] = 1
	}
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	for i := length - 1; i >= 0; i-- {
		for j := i + 1; j < length; j++ {
			if s[i] == s[j] {
				dp[i][j] = 2 + dp[i+1][j-1]
			} else {
				dp[i][j] = max(dp[i+1][j], dp[i][j-1])
			}
		}
	}
	return dp[0][length-1]
}
