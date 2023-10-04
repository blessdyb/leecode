package main

func minimumRightShifts(nums []int) int {
	rotateIndex := -1
	for i := 1; i < len(nums); i++ {
		if nums[i] < nums[i-1] {
			rotateIndex = i
			break
		}
	}
	if rotateIndex == -1 {
		return 0
	}
	for i := rotateIndex; i < len(nums)-1; i++ {
		if nums[i] > nums[i+1] {
			return -1
		}
	}
	if nums[0] < nums[len(nums)-1] {
		return -1
	}
	return len(nums) - rotateIndex
}
