package main

import "sort"

func combinationSum(candidates []int, target int) [][]int {
	ret := [][]int{}
	sort.Ints(candidates)
	combination := []int{}
	var backtracking func(int, int)
	backtracking = func(start, target int) {
		if target == 0 {
			ret = append(ret, append([]int{}, combination...))
		} else {
			for i := start; i < len(candidates); i++ {
				if candidates[i] > target {
					break
				}
				combination = append(combination, candidates[i])
				// for repeatable candidates backtracking, here we don't increase the index
				backtracking(i, target-candidates[i])
				combination = combination[:len(combination)-1]
			}
		}
	}
	backtracking(0, target)
	return ret
}
