package main

func projectionArea(grid [][]int) int {
	n := len(grid)
	top, front, side := 0, 0, 0
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	for i := 0; i < n; i++ {
		row, col := 0, 0
		for j := 0; j < n; j++ {
			if grid[i][j] > 0 {
				top++
			}
			row = max(row, grid[i][j])
			col = max(col, grid[j][i])
		}
		front += row
		side += col
	}
	return top + front + side
}
