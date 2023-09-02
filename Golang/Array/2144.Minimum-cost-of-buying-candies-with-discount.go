package main

import "sort"

func minimumCost(cost []int) int {
	sort.Ints(cost)
	ret := 0
	count := 0
	for i := len(cost) - 1; i >= 0; i-- {
		if count == 2 {
			count = 0
		} else {
			ret += cost[i]
			count++
		}
	}
	return ret
}
