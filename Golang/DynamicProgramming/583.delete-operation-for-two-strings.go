package main

func minDistance(word1 string, word2 string) int {
	min := func(a, b int) int {
		if a > b {
			return b
		}
		return a
	}
	len1, len2 := len(word1), len(word2)
	dp := make([][]int, len1+1)
	for i := 0; i <= len1; i++ {
		dp[i] = make([]int, len2+1)
		dp[i][0] = i
		for j := 1; j <= len2; j++ {
			if i == 0 {
				dp[i][j] = j
			} else if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + 1
			}
		}
	}
	return dp[len1][len2]
}
