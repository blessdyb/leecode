package main

func distributeCandies(candyType []int) int {
	hashmap := map[int]int{}
	for _, c := range candyType {
		hashmap[c]++
	}
	max := len(candyType) / 2
	if len(hashmap) <= max {
		return len(hashmap)
	}
	return max
}
