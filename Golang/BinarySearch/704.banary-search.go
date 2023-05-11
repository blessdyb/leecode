package main

// Classical Binary Search
func search(nums []int, target int) int {
	i, j := 0, len(nums)-1
	for i <= j {
		h := i + (j-i)/2
		if nums[h] > target {
			j = h - 1
		} else if nums[h] < target {
			i = h + 1
		} else {
			return h
		}
	}
	return -1
}
