package main

func maxProfit(prices []int) int {
	var max = func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	length := len(prices)
	dp := make([][]int, length)
	for i := 0; i < length; i++ {
		dp[i] = make([]int, 2)
	}
	dp[0][0] = -prices[0]
	dp[0][1] = 0
	for i := 1; i < length; i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]-prices[i])
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]+prices[i])
	}
	return max(dp[length-1][0], dp[length-1][1])
}
