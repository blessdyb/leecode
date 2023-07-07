package main

func isSubsequenceRecursive(s string, t string) bool {
	lengthS, lengthT := len(s), len(t)
	if lengthS > lengthT {
		return false
	} else if lengthS == 0 {
		return true
	}
	if s[0] == t[0] {
		return isSubsequenceRecursive(s[1:], t[1:])
	}
	return isSubsequenceRecursive(s, t[1:])
}

func isSubsequenceDP(s string, t string) bool {
	lengthS, lengthT := len(s), len(t)
	if lengthS > lengthT {
		return false
	} else if lengthS == 0 {
		return true
	}
	dp := make([][]bool, lengthS+1)
	for i := 0; i <= lengthS; i++ {
		dp[i] = make([]bool, lengthT+1)
	}
	dp[0][0] = true
	for i := 0; i <= lengthS; i++ {
		for j := 1; j <= lengthT; j++ {
			if i == 0 {
				dp[i][j] = true
			} else if s[i-1] == t[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = dp[i][j-1]
			}
		}
	}
	return dp[lengthS][lengthT]
}
