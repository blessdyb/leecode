package main

func countMaxOrSubsets(nums []int) int {
	max := 0
	hashmap := map[int]int{}
	var backtracking func(int, int)
	backtracking = func(i, bor int) {
		if i >= len(nums) {
			if max < bor {
				max = bor
			}
			hashmap[bor]++
		} else {
			backtracking(i+1, bor)
			backtracking(i+1, bor|nums[i])
		}
	}
	backtracking(0, 0)
	return hashmap[max]
}
