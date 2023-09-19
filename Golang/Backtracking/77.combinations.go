package main

func combine(n int, k int) [][]int {
	combinations := [][]int{}
	combination := []int{}
	var backtracking func(int)
	backtracking = func(i int) {
		if len(combination) == k {
			combinations = append(combinations, append([]int{}, combination...))
		} else {
			for j := i; j <= n; j++ {
				combination = append(combination, j)
				backtracking(j + 1)
				combination = combination[:len(combination)-1]
			}
		}
	}
	backtracking(1)
	return combinations
}
