package main

func findChampion(grid [][]int) int {
	hashmap := map[int]int{}
	n := len(grid)
	for i := 0; i < n; i++ {
		hashmap[i] = i
	}
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if grid[i][j] == 1 {
				hashmap[j] = i
			} else {
				hashmap[i] = j
			}
		}
	}
	ret := -1
	for k, v := range hashmap {
		if k == v {
			ret = k
			break
		}
	}
	return ret
}
