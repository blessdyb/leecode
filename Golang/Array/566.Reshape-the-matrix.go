package main

func matrixReshape(mat [][]int, r int, c int) [][]int {
	m, n := len(mat), len(mat[0])
	if m*n != r*c {
		return mat
	}
	ret := make([][]int, r)
	s, t := 0, 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if t == c {
				t = 0
				s++
			}
			if t == 0 {
				ret[s] = make([]int, c)
			}
			ret[s][t] = mat[i][j]
			t++
		}
	}
	return ret
}
