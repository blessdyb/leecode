package main

import "sort"

func sortPeople(names []string, heights []int) []string {
	hashmap := map[int]string{}
	for i := 0; i < len(heights); i++ {
		hashmap[heights[i]] = names[i]
	}
	sort.Ints(heights)
	ret := []string{}
	for i := len(heights) - 1; i >= 0; i-- {
		ret = append(ret, hashmap[heights[i]])
	}
	return ret
}
