package main

import "sort"

func heightChecker(heights []int) int {
	expected := append([]int{}, heights...)
	sort.Ints(expected)
	ret := 0
	for i := 0; i < len(heights); i++ {
		if heights[i] != expected[i] {
			ret++
		}
	}
	return ret
}
