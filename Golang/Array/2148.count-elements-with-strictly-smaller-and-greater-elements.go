package main

func countElements(nums []int) int {
	hashmap := map[int]int{}
	min, max := nums[0], nums[0]
	for _, num := range nums {
		if num < min {
			min = num
		}
		if num > max {
			max = num
		}
		hashmap[num]++
	}
	ret := len(nums) - hashmap[min] - hashmap[max]
	if ret < 0 {
		return 0
	}
	return ret
}
