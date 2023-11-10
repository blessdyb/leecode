package main

func findChampion(n int, edges [][]int) int {
	hashmap := make(map[int]bool, n)
	for _, e := range edges {
		hashmap[e[1]] = true
	}
	ret := -1
	for k, v := range hashmap {
		if !v {
			if ret == -1 {
				ret = k
			} else {
				return -1
			}
		}
	}
	return ret

}

func findChampionII(n int, edges [][]int) int {
	hashmap := map[int]int{}
	for i := 0; i < n; i++ {
		hashmap[i] = i
	}
	for i := 0; i < len(edges); i++ {
		hashmap[edges[i][1]] = edges[i][0]
	}
	champions := []int{}
	for k, v := range hashmap {
		if k == v {
			champions = append(champions, k)
		}
	}
	if len(champions) == 1 {
		return champions[0]
	}
	return -1
}
