package main

func numTilePossibilities(tiles string) int {
	hashmap := map[byte]int{}
	for i := 0; i < len(tiles); i++ {
		hashmap[tiles[i]]++
	}
	ret := 0
	var backtracking func()
	backtracking = func() {
		for tile, count := range hashmap {
			if count > 0 {
				ret++
				hashmap[tile]--
				backtracking()
				hashmap[tile]++
			}
		}
	}
	backtracking()
	return ret
}
