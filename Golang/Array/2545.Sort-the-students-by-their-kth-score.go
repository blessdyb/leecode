package main

import "sort"

func sortTheStudents(score [][]int, k int) [][]int {
	sort.Slice(score, func(a, b int) bool {
		return score[a][k] > score[b][k]
	})
	return score
}
