package main

func rotate(matrix [][]int) {
	for i := 0; i < len(matrix); i++ {
		for j := i + 1; j < len(matrix); j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}

	for i := 0; i < len(matrix); i++ {
		for j, k := 0, len(matrix)-1; j < k; j, k = j+1, k-1 {
			matrix[i][j], matrix[i][k] = matrix[j][k], matrix[i][j]
		}
	}
}
