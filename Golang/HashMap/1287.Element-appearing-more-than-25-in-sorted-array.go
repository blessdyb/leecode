package main

func findSpecialInteger(arr []int) int {
	target := len(arr) / 4
	hashmap := map[int]int{}
	for _, num := range arr {
		hashmap[num]++
		if hashmap[num] > target {
			return num
		}
	}
	return 0
}
