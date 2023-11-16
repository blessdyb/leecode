package main

import "fmt"

func stoneGameRecursive(piles []int) bool {
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	hashmap := map[string]int{}
	var helper func(l, r int) int
	helper = func(l, r int) int {
		key := fmt.Sprintf("%d-%d", l, r)
		if _, ok := hashmap[key]; ok {
			return hashmap[key]
		}
		if l > r {
			return 0
		} else if l == r {
			return piles[l]
		}
		value := max(piles[l]-helper(l+1, r), piles[r]-helper(r-1, l))
		hashmap[key] = value
		return value
	}
	return helper(0, len(piles)-1) > 0
}

func stoneGameDP(piles []int) bool {
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	n := len(piles)
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
		dp[i][i] = piles[i]
	}
	for l := 2; l <= n; l++ {
		for i := 0; i < n-(l-1); i++ {
			j := i + l - 1
			dp[i][j] = max(piles[i]-dp[i+1][j], piles[j]-dp[i][j-1])
		}
	}
	return dp[0][n-1] > 0
}
