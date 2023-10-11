package main

import "sort"

func findContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)
	ret := 0
	for i, j := len(g)-1, len(s)-1; i >= 0 && j >= 0; {
		if g[i] <= s[j] {
			ret++
			i--
			j--
		} else {
			i--
		}
	}
	return ret
}
