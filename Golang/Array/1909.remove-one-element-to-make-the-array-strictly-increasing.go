package main

func canBeIncreasing(nums []int) bool {
	count := 0
	for i := 1; i < len(nums) && count < 2; i++ {
		if nums[i] <= nums[i-1] {
			count++
			if i > 1 && nums[i-2] >= nums[i] {
				nums[i] = nums[i-1]
			}
		}
	}
	return count < 2
}

func canBeIncreasing2(nums []int) bool {
	removed := false
	n := len(nums)
	for i := 1; i < n; i++ {
		if nums[i] <= nums[i-1] {
			if removed {
				return false
			}
			if i < 2 || nums[i] > num[i-2] {
				nums[i-1] = nums[i]
			} else {
				nums[i] = nums[i-1]
			}
			removed = true
		}
	}
	return true
}
