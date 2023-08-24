package main

import "sort"

func minNumber(nums1 []int, nums2 []int) int {
	sort.Ints(nums1)
	sort.Ints(nums2)
	nums := make([]int, 9)
	for _, num := range nums1 {
		nums[num-1] |= 1
	}
	for _, num := range nums2 {
		nums[num-1] |= 2
		if nums[num-1] == 3 {
			return num
		}
	}
	a := nums1[0]
	b := nums2[0]
	if a < b {
		a, b = b, a
	}
	return b*10 + a
}
