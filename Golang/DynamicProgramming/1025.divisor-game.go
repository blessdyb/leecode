package main

func divisorGame(n int) bool {
	dp := make([]bool, n+1)
	dp[0] = false
	dp[1] = false
	for i := 2; i <= n; i++ {
		for j := 1; j < i; j++ {
			if i%j == 0 && !dp[i-j] {
				dp[i] = true
				break
			}
		}
	}
	return dp[n]
}
