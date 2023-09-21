package main

import "math"

func furthestDistanceFromOrigin(moves string) int {
	hashmap := map[byte]int{}
	for i := 0; i < len(moves); i++ {
		hashmap[moves[i]]++
	}
	return int(math.Abs(float64(hashmap['L']-hashmap['R']))) + hashmap['_']
}
