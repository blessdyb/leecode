package main

// Given a total number with a collection of items to get a maximum value, DP
// Since each item has infinitive amount, it's a full backpack problem
// We can use rolling array
// dp[i] += dp[i - coin] and dp[0] = 1 (only 1 way to get 0 amount)
func change1DDP(amount int, coins []int) int {
	dp := make([]int, amount+1)
	dp[0] = 1
	for _, coin := range coins {
		for i := coin; i <= amount; i++ {
			dp[i] += dp[i-amount]
		}
	}
	return dp[amount]
}

func change2DDP(amount int, coins []int) int {
	length := len(coins)
	dp := make([][]int, length+1)
	for i := 0; i <= length; i++ {
		dp[i] = make([]int, amount+1)
		dp[i][0] = 1
	}
	for i := 1; i <= length; i++ {
		for j := 1; j <= amount; j++ {
			if j < coins[i-1] {
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = dp[i-1][j] + dp[i][j-coins[i-1]]
			}
		}
	}
	return dp[length][amount]
}
