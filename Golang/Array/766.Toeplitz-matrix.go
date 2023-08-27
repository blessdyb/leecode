package main

import "fmt"

func isToeplitzMatrix(matrix [][]int) bool {
	m, n := len(matrix), len(matrix[0])
	diagnals := [][2]int{}
	for i := 0; i < m; i++ {
		diagnals = append(diagnals, [2]int{i, 0})
	}
	for j := 1; j < n; j++ {
		diagnals = append(diagnals, [2]int{0, j})
	}
	fmt.Println(diagnals)
	for _, diagnal := range diagnals {
		start := matrix[diagnal[0]][diagnal[1]]
		for i, j := diagnal[0], diagnal[1]; i < m && j < n; {
			if matrix[i][j] != start {
				return false
			}
			i++
			j++
		}
	}
	return true
}
