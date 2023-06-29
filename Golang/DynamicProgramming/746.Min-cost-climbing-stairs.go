package main

// Note: reach to the top means dp[len(cost)]
func minCostClimbingStairs(cost []int) int {
	top := len(cost)
	dp := make([]int, top+1)
	dp[0] = 0
	dp[1] = 0
	for i := 2; i <= top; i++ {
		oneStep := dp[i-1] + cost[i-1]
		twoSteps := dp[i-2] + cost[i-2]
		if oneStep > twoSteps {
			dp[i] = twoSteps
		} else {
			dp[i] = oneStep
		}
	}
	return dp[top]
}
