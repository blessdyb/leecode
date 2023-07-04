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
		dp[i] = make([]int, 5)
	}
	dp[0][1] = -prices[0]
	dp[0][3] = -prices[0]
	for i := 1; i < length; i++ {
		// 0: no action
		dp[i][0] = dp[i-1][0]
		// 1: hold the stock for the first time
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]-prices[i])
		// 2: sell the stock for the first time
		dp[i][2] = max(dp[i-1][2], dp[i-1][1]+prices[i])
		// 3: hold the stock for the second time
		dp[i][3] = max(dp[i-1][3], dp[i-1][2]-prices[i])
		// 4: sell the stock for the second time
		dp[i][4] = max(dp[i-1][4], dp[i-1][3]+prices[i])
	}
	return max(dp[length-1][2], dp[length-1][4])
}

func maxProfit1D(prices []int) int {
	var max = func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	length := len(prices)
	dp := make([]int, 5)
	dp[1] = -prices[0]
	dp[3] = -prices[0]
	for i := 1; i < length; i++ {
		// 0: no action
		// 1: hold the stock for the first time
		dp[1] = max(dp[1], dp[0]-prices[i])
		// 2: sell the stock for the first time
		dp[2] = max(dp[2], dp[1]+prices[i])
		// 3: hold the stock for the second time
		dp[3] = max(dp[3], dp[2]-prices[i])
		// 4: sell the stock for the second time
		dp[4] = max(dp[4], dp[3]+prices[i])
	}
	return max(dp[2], dp[4])
}
