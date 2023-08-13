package main

func countKDifference(nums []int, k int) int {
	hashmap := map[int]int{}
	for _, num := range nums {
		hashmap[num]++
	}
	ret := 0
	for _, num := range nums {
		ret += hashmap[num+k]
	}
	return ret
}
