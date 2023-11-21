package main

func maxSumAfterPartitioning(arr []int, k int) int {
	n := len(arr)
	dp := make([]int, n+1)
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	for i := 1; i <= n; i++ {
		temp := -1
		for j := 1; j <= min(i, k); j++ {
			temp = max(temp, arr[i-j])
			dp[i] = max(dp[i], dp[i-j]+temp*j)
		}
	}
	return dp[n]
}
