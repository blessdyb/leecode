package main

func findMatrix(nums []int) [][]int {
	hashmap := map[int]int{}
	for _, num := range nums {
		hashmap[num]++
	}
	ret := [][]int{}
	for len(hashmap) > 0 {
		temp := []int{}
		for k, v := range hashmap {
			if v > 0 {
				hashmap[k]--
				temp = append(temp, k)
			} else {
				delete(hashmap, k)
			}
		}
		if len(temp) > 0 {
			ret = append(ret, temp)
		}
	}
	return ret
}
