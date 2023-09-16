package main

func isCovered(ranges [][]int, left int, right int) bool {
	data := make([]int, 51)
	for _, r := range ranges {
		for i := r[0]; i <= r[1]; i++ {
			data[i] = 1
		}
	}
	ret := 1
	for i := left; i <= right; i++ {
		ret *= data[i]
	}
	return ret == 1
}
