package main

// 0-1 backpack problem
// Given a backpack with volume target, the weight nums[i] with value[i]
// If we can fulfill the whole backpack
func canPartition2DDP(nums []int) bool {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	if sum%2 == 1 {
		return false
	}
	target := sum / 2
	length := len(nums)
	dp := make([][]bool, length+1)
	for i := 0; i <= length; i++ {
		dp[i] = make([]bool, target+1)
		dp[i][0] = true
	}
	for i := 1; i <= length; i++ {
		for j := 1; j <= target; j++ {
			if j < nums[i-1] {
				// nums[i-1] can't fit in
				dp[i][j] = dp[i-1][j]
			} else {
				// With nums[i-1] or without nums[i-1]
				dp[i][j] = dp[i-1][j] || dp[i-1][j-nums[i-1]]
			}
		}
	}
	return dp[length][target]
}

func canPartition1DDP(nums []int) bool {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	if sum%2 == 1 {
		return false
	}
	target := sum / 2
	length := len(nums)
	dp := make([]bool, target+1)
	dp[0] = true
	for i := 0; i < length; i++ {
		for j := target; j >= 0; j-- {
			if j >= nums[i] {
				// dp[i][j] = dp[i-1][j] || dp[i-1][j-nums[i-1]]
				dp[j] = dp[j] || dp[j-nums[i]]
			}
		}
	}
	return dp[target]
}
