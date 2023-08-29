package main

func minMaxGame(nums []int) int {
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	for len(nums) > 1 {
		temp := []int{}
		i := 0
		for len(temp) < len(nums)/2 {
			if i%2 == 0 {
				temp = append(temp, min(nums[2*i], nums[2*i+1]))
			} else {
				temp = append(temp, max(nums[2*i], nums[2*i+1]))
			}
			i++
		}
		nums = temp
	}
	return nums[0]
}
