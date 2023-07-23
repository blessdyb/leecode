package main

func numIdenticalPairs(nums []int) int {
	hashmap := map[int]int{}
	for _, num := range nums {
		hashmap[num]++
	}
	ret := 0
	for _, frequency := range hashmap {
		for i := frequency - 1; i > 0; i-- {
			ret += i
		}
	}
	return ret
}
