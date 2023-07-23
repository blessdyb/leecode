package main

import (
	"sort"
)

func smallerNumbersThanCurrent(nums []int) []int {
	ret := []int{}
	hashmap := map[int]int{}
	clone := append([]int{}, nums...)
	sort.Ints(clone)
	for id, num := range clone {
		if _, ok := hashmap[num]; !ok {
			hashmap[num] = id
		}
	}
	for _, num := range nums {
		ret = append(ret, hashmap[num])
	}
	return ret
}
