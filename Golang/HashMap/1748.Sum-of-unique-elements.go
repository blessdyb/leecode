package main

func sumOfUnique(nums []int) int {
	hashmap := map[int]int{}
	for _, num := range nums {
		hashmap[num]++
	}
	ret := 0
	for k, v := range hashmap {
		if v == 1 {
			ret += k
		}
	}
	return ret
}
