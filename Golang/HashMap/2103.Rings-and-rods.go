package main

func countPoints(rings string) int {
	hashmap := map[byte]map[byte]int{}
	for i := 0; i < len(rings); i += 2 {
		color := rings[i]
		position := rings[i+1]
		if hashmap[position] == nil {
			hashmap[position] = map[byte]int{}
		}
		hashmap[position][color]++
	}
	ret := 0
	for _, v := range hashmap {
		if len(v) == 3 {
			ret++
		}
	}
	return ret
}
