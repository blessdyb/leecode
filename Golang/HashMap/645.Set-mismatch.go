package main

func findErrorNums(nums []int) []int {
	hashmap := map[int]int{}
	duplicate := 0
	sum := 0
	for _, num := range nums {
		hashmap[num]++
		sum += num
		if hashmap[num] > 1 {
			duplicate = num
		}
	}
	n := len(nums)
	missing := (n * (n + 1) / 2) - (sum - duplicate)
	return []int{duplicate, missing}
}
