package main

func surfaceArea(grid [][]int) int {
	ret := 0
	n := len(grid)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] > 0 {
				cubes := grid[i][j] * 6
				if grid[i][j] > 1 {
					cubes -= (grid[i][j] - 1) * 2
				}
				if i-1 >= 0 {
					if grid[i-1][j] >= grid[i][j] {
						cubes -= grid[i][j]
					} else {
						cubes -= grid[i-1][j]
					}
				}
				if j-1 >= 0 {
					if grid[i][j-1] >= grid[i][j] {
						cubes -= grid[i][j]
					} else {
						cubes -= grid[i][j-1]
					}
				}
				if i+1 < n {
					if grid[i+1][j] >= grid[i][j] {
						cubes -= grid[i][j]
					} else {
						cubes -= grid[i+1][j]
					}
				}
				if j+1 < n {
					if grid[i][j+1] >= grid[i][j] {
						cubes -= grid[i][j]
					} else {
						cubes -= grid[i][j+1]
					}
				}
				ret += cubes
			}
		}
	}
	return ret
}
