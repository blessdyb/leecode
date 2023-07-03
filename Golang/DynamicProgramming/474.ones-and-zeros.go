package main

import "strings"

// Each string can be consider as a backpack item
// zeros and ones inside of the string can be the two criteria
// To get the maximum number, it's a DP problem (3D)
// We can go with rolloing array to downgrade it back to 2D
// dp[i][j] = max(dp[i][j], dp[i-zeros][j-ones] + 1)
func findMaxForm(strs []string, m int, n int) int {
	dp := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
	}
	for _, str := range strs {
		zeros, ones := strings.Count(str, "0"), strings.Count(str, "1")
		for i := m; i >= zeros; i-- {
			for j := n; j >= ones; j-- {
				temp := dp[i-zeros][j-ones] + 1
				if temp > dp[i][j] {
					dp[i][j] = temp
				}
			}
		}
	}
	return dp[m][n]
}
