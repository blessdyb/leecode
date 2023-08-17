package main

func canBeEqual(target []int, arr []int) bool {
	hashmap := map[int]int{}
	for _, num := range target {
		hashmap[num]++
	}
	for _, num := range arr {
		hashmap[num]--
		if hashmap[num] == 0 {
			delete(hashmap, num)
		}
	}
	return len(hashmap) == 0
}
