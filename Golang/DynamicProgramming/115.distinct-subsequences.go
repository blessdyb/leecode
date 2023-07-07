package main

func numDistinct(s string, t string) int {
	lenS, lenT := len(s), len(t)
	if lenS < lenT {
		return 0
	}
	dp := make([][]int, lenS+1)
	for i := 0; i <= lenS; i++ {
		dp[i] = make([]int, lenT+1)
		dp[i][0] = 1
	}
	for i := 1; i <= lenS; i++ {
		for j := 1; j <= lenT; j++ {
			if s[i-1] == t[j-1] {
				dp[i][j] = dp[i-1][j-1] + dp[i-1][j]
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}
	return dp[lenS][lenT]
}
