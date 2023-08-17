package main

import "sort"

func frequencySort(nums []int) []int {
	hashmap := map[int]int{}
	for _, num := range nums {
		hashmap[num]++
	}
	sort.Slice(nums, func(i, j int) bool {
		if hashmap[nums[i]] < hashmap[nums[j]] {
			return true
		} else if hashmap[nums[i]] == hashmap[nums[j]] {
			return nums[i] > nums[j]
		}
		return false
	})
	return nums
}
