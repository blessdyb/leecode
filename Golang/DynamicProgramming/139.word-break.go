package main

import "strings"

func wordBreakRecursive(s string, wordDict []string) bool {
	cache := make(map[int]bool)
	var dfs func(start int) bool
	dfs = func(start int) bool {
		if value, ok := cache[start]; ok {
			return value
		}
		ret := false
		for _, w := range wordDict {
			if strings.HasPrefix(s[start:], w) {
				ret = ret || dfs(start+len(w))
			}
		}
		cache[start] = ret
		return ret
	}
	return dfs(0)
}

/*
 * If you want to calculate combinations, the outer for loop iterates over the items, and the inner for loop iterates over the backpack.
 * If you want to calculate permutations, the outer for loop iterates over the backpack, and the inner for loop iterates over the items.
 */

func wordBreakDP(s string, wordDict []string) bool {
	dp := make([]bool, len(s)+1)
	dp[0] = true
	for i := 0; i < len(s); i++ {
		for _, w := range wordDict {
			wordLength := len(w)
			if wordLength <= i && w == s[i-wordLength:i] {
				dp[i] = dp[i] || dp[i-wordLength]
			}
		}
	}
	return dp[len(s)]
}
