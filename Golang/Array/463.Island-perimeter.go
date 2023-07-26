package main

func islandPerimeter(grid [][]int) int {
	ret := 0
	row := len(grid)
	col := len(grid[0])
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if grid[i][j] == 1 {
				if i == 0 {
					ret++
				}
				if i == row-1 {
					ret++
				}
				if j == 0 {
					ret++
				}
				if j == col-1 {
					ret++
				}
				if i > 0 && grid[i-1][j] == 0 {
					ret++
				}
				if i < row-1 && grid[i+1][j] == 0 {
					ret++
				}
				if j > 0 && grid[i][j-1] == 0 {
					ret++
				}
				if j < col-1 && grid[i][j+1] == 0 {
					ret++
				}
			}
		}
	}
	return ret
}
