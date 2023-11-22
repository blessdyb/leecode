package main

import "sort"

func getKth(lo int, hi int, k int) int {
	power := func(num int) int {
		count := 0
		for num != 1 {
			if num%2 == 0 {
				num /= 2
			} else {
				num = 3*num + 1
			}
			count++
		}
		return count
	}
	nums := make([][2]int, hi-lo+1)
	for i := lo; i <= hi; i++ {
		nums[i-lo] = [2]int{i, power(i)}
	}
	sort.Slice(nums, func(a, b int) bool {
		if nums[a][1] != nums[b][1] {
			return nums[a][1] < nums[b][1]
		}
		return nums[a][0] < nums[b][0]
	})
	return nums[k-1][0]
}
