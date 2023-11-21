package main

func countVowelStrings(n int) int {
	ret := make([][]string, n)
	ret[0] = []string{"a", "e", "i", "o", "u"}
	for i := 1; i < n; i++ {
		ret[i] = []string{}
		for j := 0; j < len(ret[i-1]); j++ {
			if len(ret[i-1][j]) <= i {
				for k := 0; k < 5; k++ {
					if string(ret[i-1][j][len(ret[i-1][j])-1]) <= ret[0][k] {
						ret[i] = append(ret[i], ret[i-1][j]+ret[0][k])
					}
				}
			}
		}
	}
	return len(ret[n-1])
}

func countVowelStringsDP(n int) int {
	dp := []int{1, 1, 1, 1, 1}
	for n > 1 {
		temp := make([]int, 5)
		for i := range dp {
			for j := i; j < 5; j++ {
				temp[i] += dp[j]
			}
		}
		dp = temp
		n--
	}
	ret := 0
	for _, v := range dp {
		ret += v
	}
	return ret
}
