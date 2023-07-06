package main

func maxUncrossedLines(nums1 []int, nums2 []int) int {
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	ret, len1, len2 := 0, len(nums1), len(nums2)
	dp := make([]int, len2+1)
	for i := 1; i <= len1; i++ {
		temp := make([]int, len2+1)
		for j := 1; j <= len2; j++ {
			if nums1[i-1] == nums2[j-1] {
				temp[j] = dp[j-1] + 1
			} else {
				temp[j] = max(dp[j], temp[j-1])
			}
			ret = max(ret, temp[j])
		}
		dp = temp
	}
	return ret
}
