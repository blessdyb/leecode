package main

import "math"

func countGoodTriplets(arr []int, a int, b int, c int) int {
	ret := 0
	n := len(arr)
	compare := func(i, j, k int) bool {
		return int(math.Abs(float64(arr[i]-arr[j]))) <= k
	}
	for i := 0; i < n-2; i++ {
		for j := i + 1; j < n-1; j++ {
			for k := j + 1; k < n; k++ {
				if compare(i, j, a) && compare(j, k, b) && compare(i, k, c) {
					ret++
				}
			}
		}
	}
	return ret
}
