package main

func canJump(nums []int) bool {
	dp := make([]bool, len(nums))
	dp[0] = nums[0] != 0 || len(nums) == 1
	for i := 0; i < len(nums); i++ {
		if dp[i] {
			for j := 1; j <= nums[i] && i+j < len(nums); j++ {
				dp[i+j] = true
			}
		}
	}
	return dp[len(nums)-1]
}
