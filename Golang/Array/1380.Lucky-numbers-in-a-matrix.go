package main

func luckyNumbers(matrix [][]int) []int {
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	max := func(a, b int) int {
		if a < b {
			return b
		}
		return a
	}
	mins := []int{}
	maxs := []int{}
	m, n := len(matrix), len(matrix[0])
	for i := 0; i < m; i++ {
		temp := matrix[i][0]
		for j := 1; j < n; j++ {
			temp = min(temp, matrix[i][j])
		}
		mins = append(mins, temp)
	}
	for j := 0; j < n; j++ {
		temp := matrix[0][j]
		for i := 1; i < m; i++ {
			temp = max(temp, matrix[i][j])
		}
		maxs = append(maxs, temp)
	}
	ret := []int{}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == mins[i] && matrix[i][j] == maxs[j] {
				ret = append(ret, matrix[i][j])
			}
		}
	}
	return ret
}
