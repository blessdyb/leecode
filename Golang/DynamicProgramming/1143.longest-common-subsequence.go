package main

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func longestCommonSubsequenceDP2D(text1 string, text2 string) int {
	ret, length1, length2 := 0, len(text1), len(text2)
	dp := make([][]int, length1+1)
	for i := 0; i <= length1; i++ {
		dp[i] = make([]int, length2+1)
	}
	for i := 1; i <= length1; i++ {
		for j := 1; j <= length2; j++ {
			if text1[i-1] == text2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
			ret = max(ret, dp[i][j])
		}
	}
	return ret
}

func longestCommonSubsequenceDP1D(text1 string, text2 string) int {
	ret, length1, length2 := 0, len(text1), len(text2)
	dp := make([]int, length2+1)
	for i := 1; i <= length1; i++ {
		temp := make([]int, length2+1)
		for j := 1; j <= length2; j++ {
			if text1[i-1] == text2[j-1] {
				temp[j] = dp[j-1] + 1
			} else {
				temp[j] = max(dp[j], temp[j-1])
			}
			ret = max(ret, temp[j])
		}
		dp = temp
	}
	return ret
}
