package main

func hasGroupSizeX(deck []int) bool {
	hashmap := map[int]int{}
	for _, card := range deck {
		hashmap[card]++
	}
	minCounts := hashmap[deck[0]]
	for _, v := range hashmap {
		if v < minCounts {
			minCounts = v
		}
	}
	for i := 2; i < minCounts; i++ {
		flag := true
		for _, v := range hashmap {
			if v%i != 0 {
				flag = false
				break
			}
		}
		if flag {
			return true
		}
	}
	return false
}
