package main

import "math"

func findMaxK(nums []int) int {
	hashmap := map[int]bool{}
	for _, num := range nums {
		hashmap[num] = true
	}
	ret := 0
	for k, _ := range hashmap {
		if hashmap[-k] && (ret == 0 || ret < int(math.Abs(float64(k)))) {
			ret = int(math.Abs(float64(k)))
		}
	}
	if ret == 0 {
		return -1
	}
	return ret
}
