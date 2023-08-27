package main

import "sort"

func sortEvenOdd(nums []int) []int {
	odds := []int{}
	evens := []int{}
	for idx, num := range nums {
		if idx%2 == 0 {
			evens = append(evens, num)
		} else {
			odds = append(odds, num)
		}
	}
	ret := []int{}
	sort.Ints(evens)
	sort.Sort(sort.Reverse(sort.IntSlice(odds)))
	for i := 0; len(ret) < len(nums); i++ {
		if i < len(evens) {
			ret = append(ret, evens[i])
		}
		if i < len(odds) {
			ret = append(ret, odds[i])
		}
	}
	return ret
}
