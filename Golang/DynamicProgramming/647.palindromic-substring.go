package main

func countSubstringsDP(s string) int {
	length := len(s)
	ret := 0
	dp := make([][]bool, length)
	// dp[i][j] defines if s[i:j+1] is a panlindromic string
	// When s[i] == s[j], dp[i][j] = dp[i+1][j-1]
	// So we have to iterate i from right to left and j from bottom to top
	for i := 0; i < length; i++ {
		dp[i] = make([]bool, length)
	}
	for i := length - 1; i >= 0; i-- {
		for j := i; j < length; j++ {
			if s[i] == s[j] {
				if j-i <= 1 {
					dp[i][j] = true
					ret++
				} else if dp[i+1][j-1] {
					dp[i][j] = true
					ret++
				}
			}
		}
	}
	return ret
}

func countSubstringsTwoPointers(s string) int {
	ret, length := 0, len(s)
	checkPanlindromic := func(start, end int) int {
		ret := 0
		for start >= 0 && end < length && s[start] == s[end] {
			start--
			end++
			ret++
		}
		return ret
	}
	for i := 0; i < length; i++ {
		ret += checkPanlindromic(i, i)
		ret += checkPanlindromic(i, i+1)
	}
	return ret
}
