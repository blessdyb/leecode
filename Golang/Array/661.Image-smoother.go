package main

func imageSmoother(img [][]int) [][]int {
	m, n := len(img), len(img[0])
	ret := make([][]int, m)
	for i := 0; i < m; i++ {
		ret[i] = make([]int, n)
		for j := 0; j < n; j++ {
			sum := img[i][j]
			count := 1
			if i-1 >= 0 {
				sum += img[i-1][j]
				count++
				if j-1 >= 0 {
					sum += img[i-1][j-1]
					count++
				}
				if j+1 < n {
					sum += img[i-1][j+1]
					count++
				}
			}
			if i+1 < m {
				sum += img[i+1][j]
				count++
				if j-1 >= 0 {
					sum += img[i+1][j-1]
					count++
				}
				if j+1 < n {
					sum += img[i+1][j+1]
					count++
				}
			}
			if j-1 >= 0 {
				sum += img[i][j-1]
				count++
			}
			if j+1 < n {
				sum += img[i][j+1]
				count++
			}
			ret[i][j] = sum / count
		}
	}
	return ret
}
