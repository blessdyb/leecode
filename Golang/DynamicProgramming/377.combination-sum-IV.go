package main

// Full-backpack problem given the target and items
// Since we can pick the same item repeatly, dp[i] += dp[i - num]
// The problem is asking for combination, so the sequence of picking the items matters
func combinationSum4(nums []int, target int) int {
	dp := make([]int, target+1)
	dp[0] = 1
	for i := 1; i <= target; i++ {
		for _, num := range nums {
			if i >= num {
				dp[i] += dp[i-num]
			}
		}
	}
	return dp[target]
}
