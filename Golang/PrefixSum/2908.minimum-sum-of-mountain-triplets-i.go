package main

func minimumSum(nums []int) int {
	ret, minNum := 200, 100
	n := len(nums)
	preMin, sufMin := make([]int, n), make([]int, n)
	min := func(a, b int) int {
		if a > b {
			return b
		}
		return a
	}
	for i := 1; i < n-1; i++ {
		preMin[i] = min(nums[i-1], minNum)
		minNum = preMin[i]
	}
	minNum = 100
	for i := n - 2; i >= 0; i-- {
		sufMin[i] = min(nums[i+1], minNum)
		minNum = sufMin[i]
	}
	for i := 1; i < n-1; i++ {
		if nums[i] > preMin[i-1] && nums[i] > sufMin[i] {
			ret = min(ret, nums[i]+sufMin[i]+preMin[i])
		}
	}
	if ret == 200 {
		return -1
	}
	return ret
}
