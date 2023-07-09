package main

func longestPalindromeDP(s string) string {
	length := len(s)
	ret, max := "", 0
	dp := make([][]bool, length)
	for i := 0; i < length; i++ {
		dp[i] = make([]bool, length)
	}
	for i := length - 1; i >= 0; i-- {
		for j := i; j < length; j++ {
			if s[i] == s[j] {
				if j-i <= 1 {
					dp[i][j] = true
				} else {
					dp[i][j] = dp[i+1][j-1]
				}
				if dp[i][j] && max < j-i+1 {
					max = j - i + 1
					ret = s[i : j+1]
				}
			}
		}
	}
	return ret
}

func longestPalindromeTwoPointers(s string) string {
	ret, max, length := "", 0, len(s)
	palindrom := func(start, end int) int {
		for start >= 0 && end < length && s[start] == s[end] {
			start--
			end++
		}
		//return (end - 1) - (start + 1) + 1
		return end - start - 1
	}
	for i := 0; i < length; i++ {
		p1 := palindrom(i, i)
		p2 := palindrom(i, i+1)
		if p1 < p2 && max < p2 {
			max = p2
			start := i - (p2-2)/2
			ret = s[start : start+max]
		} else if p1 > p2 && max < p1 {
			max = p1
			start := i - (p1-1)/2
			ret = s[start : start+max]
		}
	}
	return ret
}
