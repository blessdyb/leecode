package main

func repeatedNTimes(nums []int) int {
	hashmap := map[int]int{}
	var ret int
	for _, num := range nums {
		hashmap[num]++
		if hashmap[num] > 1 {
			ret = num
			break
		}
	}
	return ret
}
