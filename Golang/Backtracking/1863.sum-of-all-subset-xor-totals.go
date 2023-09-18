package main

func subsetXORSum(nums []int) int {
	var backtracking func(i, sum int) int
	backtracking = func(i, sum int) int {
		if i >= len(nums) {
			return sum
		}
		return backtracking(i+1, sum^nums[i]) + backtracking(i+1, sum)
	}
	return backtracking(0, 0)
}
