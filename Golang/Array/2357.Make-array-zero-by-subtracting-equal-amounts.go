package main

import "sort"

func minimumOperations(nums []int) int {
	sort.Ints(nums)
	ret := 0
	for {
		candidate := 0
		for _, num := range nums {
			if num != 0 {
				candidate = num
				break
			}
		}
		if candidate != 0 {
			for i := 0; i < len(nums); i++ {
				if nums[i] != 0 {
					nums[i] -= candidate
				}
			}
			ret++
		} else {
			break
		}
	}
	return ret
}
