package main

import "math"

func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i := 1; i <= amount; i++ {
		dp[i] = math.MaxInt32
	}
	for _, coin := range coins {
		for i := coin; i <= amount; i++ {
			pick := dp[i-coin] + 1
			if pick < dp[i] {
				dp[i] = pick
			}
		}
	}
	if dp[amount] == math.MaxInt32 {
		return -1
	}
	return dp[amount]
}
