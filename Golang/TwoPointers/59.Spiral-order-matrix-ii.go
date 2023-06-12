package main

func generateMatrix(n int) [][]int {
	ret := make([][]int, n)
	for i := 0; i < n; i++ {
		ret[i] = make([]int, n)
	}
	total := n * n
	up, right, down, left := 0, n-1, n-1, 0
	for i := 0; i < total; {
		for j := left; j <= right; j++ {
			ret[up][j] = i
			i++
		}
		for j := up + 1; j <= down; j++ {
			ret[j][right] = i
			i++
		}
		if up != down {
			for j := right - 1; j >= left; j-- {
				ret[down][j] = i
				i++
			}
		}
		if left != right {
			for j := down - 1; j > up; j-- {
				ret[j][left] = i
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
