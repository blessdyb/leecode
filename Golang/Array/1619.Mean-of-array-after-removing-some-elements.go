package main

import "sort"

func trimMean(arr []int) float64 {
	n := len(arr)
	removal := n / 20
	sort.Ints(arr)
	arr = arr[removal : n-removal]
	sum := 0
	for _, num := range arr {
		sum += num
	}
	return float64(sum) / float64(len(arr))
}
