package main

func maxAreaOfIsland(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	visited := make([][]bool, m)
	for i := 0; i < m; i++ {
		visited[i] = make([]bool, n)
	}
	var helper func(i, j int) int
	helper = func(i, j int) int {
		if i < 0 || i >= m || j < 0 || j >= n || grid[i][j] == 0 {
			return 0
		}
		if !visited[i][j] {
			visited[i][j] = true
			return 1 + helper(i+1, j) + helper(i-1, j) + helper(i, j+1) + helper(i, j-1)
		}
		return 0
	}
	max := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if !visited[i][j] {
				current := helper(i, j)
				if max < current {
					max = current
				}
			}
		}
	}
	return max
}
