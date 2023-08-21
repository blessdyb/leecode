package main

func largestLocal(grid [][]int) [][]int {
	ret := make([][]int, len(grid)-2)
	for i := 0; i < len(grid)-2; i++ {
		ret[i] = make([]int, len(grid)-2)
	}
	for i := 0; i < len(ret); i++ {
		for j := 0; j < len(ret); j++ {
			max := 0
			for m := i; m < i+3; m++ {
				for n := j; n < j+3; n++ {
					if max < grid[m][n] {
						max = grid[m][n]
					}
				}
			}
			ret[i][j] = max
		}
	}
	return ret
}
