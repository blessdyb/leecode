package main

func check(nums []int) bool {
	rotateIndex := -1
	for i := 1; i < len(nums); i++ {
		if nums[i] < nums[i-1] {
			rotateIndex = i
			break
		}
	}
	for i := rotateIndex; i < len(nums)-1; i++ {
		if nums[i] > nums[i+1] {
			return false
		}
	}
	if rotateIndex != -1 {
		return nums[0] >= nums[len(nums)-1]
	}
	return true
}
