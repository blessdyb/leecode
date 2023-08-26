package main

func checkXMatrix(grid [][]int) bool {
	n := len(grid)
	topLeft := []int{0, 0}
	topRight := []int{0, n - 1}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i == topLeft[0] && j == topLeft[1] {
				if grid[i][j] == 0 {
					return false
				}
			} else if i == topRight[0] && j == topRight[1] {
				if grid[i][j] == 0 {
					return false
				}
			} else if grid[i][j] != 0 {
				return false
			}
		}
		topLeft[0]++
		topLeft[1]++
		topRight[0]++
		topRight[1]--
	}
	return true
}
