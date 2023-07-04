package main

func maxProfit3Status(prices []int) int {
	var max = func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	length := len(prices)
	// dp[i][0], on the ith day, we hold the stock
	// dp[i][1], on the ith day, we don't have any stocks
	// dp[i][2], on the ith day, we cooldown
	dp := make([][]int, length)
	for i := 0; i < length; i++ {
		dp[i] = make([]int, 3)
	}
	dp[0][1] = -prices[0]
	for i := 1; i < length; i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][2]-prices[i])
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]+prices[i])
		dp[i][2] = max(dp[i-1][2], dp[i-1][1])
	}
	return max(dp[length-1][1], dp[length-1][2])
}

func maxProfit2Status(prices []int) int {
	length := len(prices)
	if length < 2 {
		return 0
	}
	var max = func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	// dp[i][0], on the ith day, we hold the stock
	// dp[i][1], on the ith day, we don't have any stocks
	dp := make([][]int, length)
	for i := 0; i < length; i++ {
		dp[i] = make([]int, 2)
	}
	dp[0][0] = -prices[0]
	dp[0][1] = 0
	dp[1][0] = max(dp[0][0], -prices[1])
	dp[1][1] = max(dp[0][1], dp[0][0]+prices[1])
	for i := 2; i < length; i++ {
		// On day i we have stock on hands, a) no action , b) buy stock - previous day is a cooldown day
		dp[i][0] = max(dp[i-1][0], dp[i-2][1]-prices[i])
		// On day i we don't have stock, a) no action, b) sell stock
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]+prices[i])
	}
	return max(dp[length-1][0], dp[length-1][1])
}
