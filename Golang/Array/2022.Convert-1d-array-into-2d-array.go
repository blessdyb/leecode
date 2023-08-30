package main

func construct2DArray(original []int, m int, n int) [][]int {
	ret := [][]int{}
	if len(original) == m*n {
		for i := 0; i < m; i++ {
			ret = append(ret, original[i*n:(i+1)*n])
		}
	}
	return ret
}
