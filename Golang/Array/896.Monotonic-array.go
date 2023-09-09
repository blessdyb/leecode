package main

func isMonotonic(nums []int) bool {
	if len(nums) < 3 {
		return true
	}
	a, b := 0, 0
	for i := 1; i < len(nums); i++ {
		if nums[i-1] < nums[i] {
			a++
		} else if nums[i-1] > nums[i] {
			b++
		}
		if a > 0 && b > 0 {
			return false
		}
	}
	return true
}
