package main

func findLength(nums1 []int, nums2 []int) int {
	ret := 0
	length1, length2 := len(nums1), len(nums2)
	dp := make([][]int, length1+1)
	for i := 0; i <= length1; i++ {
		dp[i] = make([]int, length2+1)
	}
	// dp[i][j] = dp[i-1][j-1] + 1 (when num1[i] == num2[j])
	for i := 1; i <= length1; i++ {
		for j := 1; j <= length2; j++ {
			if nums1[i-1] == nums2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			}
			if dp[i][j] > ret {
				ret = dp[i][j]
			}
		}
	}
	return ret
}
