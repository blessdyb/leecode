package main

import "sort"

func average(salary []int) float64 {
	sort.Ints(salary)
	var ret float64 = 0
	for i := 1; i < len(salary)-1; i++ {
		ret += float64(salary[i])
	}
	return ret / float64(len(salary)-2)
}
