package main

// Get a list of perfect numbers, then we convert the problem to a full backpack
func numSquares(n int) int {
	perfectNums := []int{}
	for i := 1; i*i <= n; i++ {
		perfectNums = append(perfectNums, i*i)
	}
	dp := make([]int, n+1)
	dp[1] = 1
	for i := 2; i <= n; i++ {
		dp[i] = n
	}
	for i := 2; i <= n; i++ {
		for _, num := range perfectNums {
			if i >= num {
				pick := dp[i-num] + 1
				if pick < dp[i] {
					dp[i] = pick
				}
			}
		}
	}
	return dp[n]
}
