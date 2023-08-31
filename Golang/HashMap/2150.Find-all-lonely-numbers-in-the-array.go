package main

func findLonely(nums []int) []int {
	hashmap := map[int]int{}
	for _, num := range nums {
		hashmap[num]++
	}
	ret := []int{}
	for key, value := range hashmap {
		if value == 1 && hashmap[key-1] == 0 && hashmap[key+1] == 0 {
			ret = append(ret, key)
		}
	}
	return ret
}
