package main

import (
	"sort"
	"strconv"
)

func findRelativeRanks(score []int) []string {
	cloned := append([]int{}, score...)
	sort.Sort(sort.Reverse(sort.IntSlice(cloned)))
	hashmap := map[int]string{}
	for i, v := range cloned {
		var result string
		switch i {
		case 0:
			result = "Gold Medal"
		case 1:
			result = "Silver Medal"
		case 2:
			result = "Bronze Medal"
		default:
			result = strconv.Itoa(i + 1)
		}
		hashmap[v] = result
	}
	ret := []string{}
	for _, v := range score {
		ret = append(ret, hashmap[v])
	}
	return ret
}
