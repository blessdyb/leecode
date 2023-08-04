package main

func containsDuplicate(nums []int) bool {
	hashmap := map[int]int{}
	for _, num := range nums {
		hashmap[num]++
		if hashmap[num] > 1 {
			return true
		}
	}
	return false
}
