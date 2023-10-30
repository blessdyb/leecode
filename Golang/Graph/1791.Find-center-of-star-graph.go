package main

func findCenter(edges [][]int) int {
	hash := map[int]int{}
	n := len(edges)
	for i := 0; i < n; i++ {
		hash[edges[i][0]]++
		hash[edges[i][1]]++
		if hash[edges[i][0]] == n {
			return edges[i][0]
		}
		if hash[edges[i][1]] == n {
			return edges[i][1]
		}
	}
	return 0
}
