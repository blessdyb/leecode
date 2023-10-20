package main

func findRotation(mat [][]int, target [][]int) bool {
	n := len(mat)
	ret := [4]bool{true, true, true, true}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if mat[i][j] != target[i][j] {
				ret[0] = false
			}
			if mat[i][j] != target[j][n-i-1] {
				ret[1] = false
			}
			if mat[i][j] != target[n-i-1][n-j-1] {
				ret[2] = false
			}
			if mat[i][j] != target[n-j-1][i] {
				ret[3] = false
			}
			if !(ret[0] || ret[1] || ret[2] || ret[3]) {
				return false
			}
		}
	}
	return ret[0] || ret[1] || ret[2] || ret[3]
}
