package main

func findClosestNumber(nums []int) int {
	min := 100001
	var num int
	for i := 0; i < len(nums); i++ {
		dist := nums[i]
		if nums[i] < 0 {
			dist = -dist
		}
		if min > dist {
			min = dist
			num = nums[i]
		} else if min == dist && num < nums[i] {
			num = nums[i]
		}
	}
	return num
}
