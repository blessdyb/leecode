package main

func findSubarrays(nums []int) bool {
	hashmap := map[int]int{}
	hashmap[nums[0]+nums[1]] = 1
	for i := 1; i < len(nums)-1; i++ {
		hashmap[nums[i]+nums[i+1]]++
		if hashmap[nums[i]+nums[i+1]] > 1 {
			return true
		}
	}
	return false
}
