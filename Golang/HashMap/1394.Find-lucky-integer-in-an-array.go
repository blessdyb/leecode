package main

func findLucky(arr []int) int {
	ret := -1
	hashmap := map[int]int{}
	for _, num := range arr {
		hashmap[num]++
	}
	for k, v := range hashmap {
		if k == v && (ret == -1 || k > ret) {
			ret = k
		}
	}
	return ret
}
