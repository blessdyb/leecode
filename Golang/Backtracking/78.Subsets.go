package main

func subsets(nums []int) [][]int {
	ret := [][]int{}
	combination := []int{}
	var backtracking func(i int)
	backtracking = func(i int) {
		ret = append(ret, append([]int{}, combination...))
		for j := i; j < len(nums); j++ {
			combination = append(combination, nums[j])
			backtracking(j + 1)
			combination = combination[:len(combination)-1]
		}
	}
	backtracking(0)
	return ret
}
