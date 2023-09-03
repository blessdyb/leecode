package main

func checkIfExist(arr []int) bool {
	hashmap := map[int]bool{}
	for _, num := range arr {
		if hashmap[num*2] || num%2 == 0 && hashmap[num/2] {
			return true
		}
		hashmap[num] = true
	}
	return false
}
