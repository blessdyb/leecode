package main

func mostFrequent(nums []int, key int) int {
	n := len(nums)
	hashmap := map[int]int{}
	for i := 0; i < n-1; i++ {
		if nums[i] == key {
			hashmap[nums[i+1]]++
		}
	}
	max := 0
	ret := 0
	for k, v := range hashmap {
		if v > max {
			max = v
			ret = k
		}
	}
	return ret
}
