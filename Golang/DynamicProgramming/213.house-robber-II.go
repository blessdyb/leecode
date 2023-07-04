package main

func rob(nums []int) int {
	var robInner = func(nums []int) int {
		if len(nums) == 0 {
			return 0
		}
		dp := make([]int, len(nums)+1)
		dp[0] = 0
		dp[1] = nums[0]
		for i := 2; i <= len(nums); i++ {
			with := dp[i-2] + nums[i-1]
			without := dp[i-1]
			if with > without {
				dp[i] = with
			} else {
				dp[i] = without
			}
		}
		return dp[len(nums)]
	}
	var max = func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	oneItem := nums[0]
	// withoutHeadTail is a sub case of below twos
	withoutHead := robInner(nums[1:])
	withoutTail := robInner(nums[:len(nums)-1])
	return max(withoutHead, max(withoutTail, oneItem))
}
