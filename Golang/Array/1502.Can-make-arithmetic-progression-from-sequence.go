package main

import "sort"

func canMakeArithmeticProgression(arr []int) bool {
	sort.Ints(arr)
	diff := arr[1] - arr[0]
	for i := 2; i < len(arr); i++ {
		if diff != arr[i]-arr[i-1] {
			return false
		}
	}
	return true
}
