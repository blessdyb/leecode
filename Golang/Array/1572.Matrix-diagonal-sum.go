package main

func diagonalSum(mat [][]int) int {
	n := len(mat)
	ret := 0
	for i := 0; i < n; i++ {
		ret += mat[i][i]
	}
	for i, j := n-1, 0; i >= 0 && j < n; {
		ret += mat[i][j]
		i--
		j++
	}
	if n%2 == 1 {
		center := n / 2
		ret -= mat[center][center]
	}
	return ret
}
