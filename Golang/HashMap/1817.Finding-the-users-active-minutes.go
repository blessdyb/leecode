package main

func findingUsersActiveMinutes(logs [][]int, k int) []int {
	hashmap := map[int]map[int]bool{}
	for _, log := range logs {
		if _, ok := hashmap[log[0]]; !ok {
			hashmap[log[0]] = map[int]bool{}
		}
		hashmap[log[0]][log[1]] = true
	}
	ret := make([]int, k)
	for _, v := range hashmap {
		ret[len(v)-1]++
	}
	return ret
}
