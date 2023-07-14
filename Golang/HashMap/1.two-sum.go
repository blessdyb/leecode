package main

func twoSum(nums []int, target int) []int {
	hashmap := map[int]int{}
	var ret []int
	for idx, num := range nums {
		if hashmap[num] > 0 {
			ret = []int{idx, hashmap[num] - 1}
			break
		}
		hashmap[target-num] = idx + 1
	}
	return ret
}
