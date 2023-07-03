package main

func findTargetSumWaysRecursive(nums []int, target int) int {
	ret := 0
	var recursive func(nums []int, t int)
	recursive = func(nums []int, t int) {
		if len(nums) == 0 {
			if t == 0 {
				ret++
			}
			return
		}
		recursive(nums[1:], t+nums[0])
		recursive(nums[1:], t-nums[0])
	}
	recursive(nums, target)
	return ret
}
