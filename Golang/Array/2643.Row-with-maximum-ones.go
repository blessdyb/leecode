package main

func rowAndMaximumOnes(mat [][]int) []int {
	ret := []int{0, 0}
	for i := 0; i < len(mat); i++ {
		sum := 0
		for j := 0; j < len(mat[i]); j++ {
			sum += mat[i][j]
		}
		if ret[1] < sum {
			ret = []int{i, sum}
		}
	}
	return ret
}
