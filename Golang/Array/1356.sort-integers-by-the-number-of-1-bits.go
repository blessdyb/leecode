package main

import "sort"

func sortByBits(arr []int) []int {
	countOnes := func(n int) int {
		count := 0
		for n > 0 {
			n = n & (n - 1)
			count++
		}
		return count
	}
	sort.Slice(arr, func(i, j int) bool {
		a, b := countOnes(arr[i]), countOnes(arr[j])
		if a == b {
			return arr[i] < arr[j]
		}
		return a < b
	})
	return arr
}
