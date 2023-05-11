package main

// A sorted array with target, so binary search.
// Since we need to find the range (left and right), it means when we get the target, we need to remember the index and continue the search.
func searchRange(nums []int, target int) []int {
	return []int{
		binarySearch(nums, target, true),
		binarySearch(nums, target, false),
	}
}

func binarySearch(nums []int, target int, direction bool) int {
	i, j := 0, len(nums)-1
	ret := -1
	for i <= j {
		h := i + (j-i)/2
		if nums[h] == target {
			ret = h
			if direction {
				j = h - 1
			} else {
				i = h + 1
			}
		} else if nums[h] > target {
			j = h - 1
		} else {
			i = h + 1
		}
	}
	return ret
}
