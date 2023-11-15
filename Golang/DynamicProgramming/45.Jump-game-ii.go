package main

func canJump(nums []int) int {
	dp1, dp2 := make([]bool, len(nums)), make([]int, len(nums))
	dp1[0] = true
	dp2[0] = 0
	for i := 0; i < len(nums); i++ {
		if dp1[i] {
			for j := 1; j < nums[i] && i+j < len(nums); j++ {
				dp[i+j] = true
				if dp2[i+j] == 0 {
					dp2[i+j] = dp2[i] + 1
				} else if dp2[i+j] > dp2[i]+1 {
					dp2[i+j] = dp2[i] + 1
				}
			}
		}
	}
	return dp2[len(nums)-1]
}
