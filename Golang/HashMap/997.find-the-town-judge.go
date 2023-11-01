package main

func findJudge(n int, trust [][]int) int {
	if n == 1 && len(trust) == 0 {
		return 1
	}
	hashmap := map[int]int{}
	trustmap := map[int]bool{}
	candidate := [2]int{0, -1}
	for i := 0; i < len(trust); i++ {
		hashmap[trust[i][1]]++
		trustmap[trust[i][0]] = true
		if hashmap[trust[i][1]] > candidate[0] {
			candidate = [2]int{hashmap[trust[i][1]], trust[i][1]}
		}
	}
	if candidate[0] == n-1 && !trustmap[candidate[1]] {
		return candidate[1]
	}
	return -1
}
