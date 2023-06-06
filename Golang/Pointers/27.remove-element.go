package main

// For array in-place operation , we have to consider if two pointers solution is applicable.
// Given (i, j) with initial value (0, 0)
// If nums[j] != target, we can move forward both i and j and assign nums[i] = nums[j] (to override the value when i !== j && nums[i] == target)
// Otherwise, we can only move j forward

func removeElement(nums []int, val int) int {
	i, j := 0, 0
	length := len(nums)
	for j < length {
		if val != nums[j] {
			nums[i] = nums[j]
			i++
		}
		j++
	}
	return i
}
