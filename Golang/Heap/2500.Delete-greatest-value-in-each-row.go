package main

import "sort"

func deleteGreatestValue(grid [][]int) int {
	ret := 0
	for len(grid) > 0 {
		max := 0
		for i := 0; i < len(grid); i++ {
			sort.Ints(grid[i])
			if max < grid[i][len(grid[i])-1] {
				max = grid[i][len(grid[i])-1]
			}
			grid[i] = grid[i][:len(grid[i])-1]
			if len(grid[i]) == 0 && i == len(grid)-1 {
				grid = nil
			}
		}
		ret += max
	}
	return ret
}
