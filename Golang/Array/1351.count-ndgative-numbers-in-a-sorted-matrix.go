package main

func countNegatives(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	first := n - 1
	ret := 0
	for i := 0; i < m; i++ {
		for first >= 0 && grid[i][first] < 0 {
			first--
		}
		ret += n - (first + 1)
	}
	return ret
}
