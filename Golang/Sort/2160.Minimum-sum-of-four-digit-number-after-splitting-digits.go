package main

import "sort"

func minimumSum(num int) int {
	a := num / 1000
	b := (num - a*1000) / 100
	c := (num - a*1000 - b*100) / 10
	d := (num - a*1000 - b*100 - c*10)
	nums := []int{a, b, c, d}
	sort.Ints(nums)
	if nums[0] == 0 && nums[1] == 0 {
		return nums[2] + nums[3]
	} else if nums[0] == 0 {
		return nums[2] + nums[1]*10 + nums[3]
	}
	return nums[0]*10 + nums[2] + nums[1]*10 + nums[3]
}
