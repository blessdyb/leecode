package main

func checkValid(matrix [][]int) bool {
	n := len(matrix)
	for i := 0; i < n; i++ {
		h1, h2 := map[int]int{}, map[int]int{}
		for j := 0; j < n; j++ {
			h1[matrix[i][j]]++
			h2[matrix[j][i]]++
		}
		if len(h1) != n || len(h2) != n {
			return false
		}
	}
	return true
}
