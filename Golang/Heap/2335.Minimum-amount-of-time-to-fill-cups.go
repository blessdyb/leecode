package main

import "sort"

func fillCups(amount []int) int {
	ret := 0
	fill := func(amount []int) []int {
		sort.Ints(amount)
		noneZeroIndex := len(amount)
		for i := 0; i < len(amount); i++ {
			if amount[i] != 0 && noneZeroIndex == len(amount) {
				noneZeroIndex = i
				break
			}
		}
		amount = amount[noneZeroIndex:]
		if len(amount) == 1 {
			amount[0] -= 1
			ret++
		} else if len(amount) > 1 {
			ret++
			amount[len(amount)-1] -= 1
			amount[len(amount)-2] -= 1
		}
		return amount
	}
	for len(amount) > 0 {
		amount = fill(amount)
	}
	return ret
}
