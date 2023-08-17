package main

func findFinalValue(nums []int, original int) int {
	hashmap := map[int]bool{}
	for _, num := range nums {
		hashmap[num] = true
	}
	for hashmap[original] {
		original *= 2
	}
	return original
}
