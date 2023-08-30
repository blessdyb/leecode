package main

import "sort"

func minimumAbsDifference(arr []int) [][]int {
	sort.Ints(arr)
	minValue := int(1e7)
	ret := [][]int{}
	for i := 0; i < len(arr)-1; i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[j]-arr[i] < minValue {
				ret = [][]int{
					[]int{arr[i], arr[j]},
				}
				minValue = arr[j] - arr[i]
			} else if arr[j]-arr[i] == minValue {
				ret = append(ret, []int{arr[i], arr[j]})
			} else {
				break
			}
		}
	}
	return ret
}
