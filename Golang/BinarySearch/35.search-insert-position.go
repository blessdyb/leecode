package main

// Given a sorted array and target, so binary search is the solution
// * Out range case: 0 or len
// * Exist: h
// * In the range: j + 1 or i
func searchInsert(nums []int, target int) int {
	i, j := 0, len(nums)-1
	if target < nums[i] {
		return 0
	} else if target > nums[j] {
		return j + 1
	}
	for i <= j {
		h := i + (j-i)/2
		if target == nums[h] {
			return h
		} else if target > nums[h] {
			i = h + 1
		} else {
			j = h - 1
		}
	}
	return j + 1
}
