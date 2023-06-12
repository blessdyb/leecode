package main

func spiralOrder(matrix [][]int) []int {
	row, column := len(matrix), len(matrix[0])
	ret := make([]int, row*column)
	up, right, down, left := 0, column-1, row-1, 0
	total := row * column
	for i := 0; i < total; {
		for j := left; j <= right; j++ {
			ret[i] = matrix[up][j]
			i++
		}
		for j := up + 1; j <= down; j++ {
			ret[i] = matrix[j][right]
			i++
		}
		if up != down {
			for j := right - 1; j >= left; j-- {
				ret[i] = matrix[down][j]
				i++
			}
		}
		if left != right {
			for j := down - 1; j > up; j-- {
				ret[i] = matrix[j][left]
				i++
			}
		}
		up++
		right--
		down--
		left++
	}
	return ret
}
