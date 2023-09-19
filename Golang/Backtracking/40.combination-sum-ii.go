package main

import "sort"

func combinationSum2(candidates []int, target int) [][]int {
	ret := [][]int{}
	sort.Ints(candidates)
	combination := []int{}
	var backtracking func(int, int)
	backtracking = func(index, target int) {
		if 0 == target {
			ret = append(ret, append(make([]int, 0, len(combination)), combination...))
		} else {
			for i := index; i < len(candidates); i++ {
				if candidates[i] > target {
					break
				}
				if i > index && candidates[i] == candidates[i-1] {
					continue
				}
				combination = append(combination, candidates[i])
				backtracking(i+1, target-candidates[i])
				combination = combination[:len(combination)-1]
			}
		}
	}
	backtracking(0, target)
	return ret
}
