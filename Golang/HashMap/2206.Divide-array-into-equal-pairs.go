package main

func divideArray(nums []int) bool {
	hashmap := map[int]int{}
	for _, num := range nums {
		hashmap[num]++
	}
	for _, v := range hashmap {
		if v%2 == 1 {
			return false
		}
	}
	return true
}
