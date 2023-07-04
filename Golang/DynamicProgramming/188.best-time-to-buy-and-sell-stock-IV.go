package main

func maxProfitDP3D(k int, prices []int) int {
	var max = func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	length := len(prices)
	// dp[i][j][0] on the ith day, we start the jth transaction and hold a stock
	// dp[i][j][1] on the ith day, we end the jth transaction and sell a stock
	dp := make([][][]int, length)
	for i := 0; i < length; i++ {
		dp[i] = make([][]int, k+1)
		for j := 0; j <= k; j++ {
			dp[i][j] = make([]int, 2)
		}
	}
	for j := 0; j <= k; j++ {
		dp[0][j][0] = -prices[0]
	}
	for i := 1; i < length; i++ {
		for j := 1; j <= k; j++ {
			// on ith day, we don't have the stock
			// a. on (i-1)th day, we held the stock, no actions today, so the same transaction k
			// b. on (i-1)th day, we sold the stock as transaction k-1, so today we buy it again as a new transaction k
			dp[i][j][0] = max(dp[i-1][j][0], dp[i-1][j-1][1]-prices[i])
			dp[i][j][1] = max(dp[i-1][j][1], dp[i-1][j][0]+prices[i])
		}
	}
	return dp[length-1][k][1]
}
