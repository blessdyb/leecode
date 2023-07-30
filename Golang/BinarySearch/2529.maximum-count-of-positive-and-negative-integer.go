package main

func maximumCount(nums []int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] > 0 {
			right = mid - 1
		} else if nums[mid] < 0 {
			left = mid + 1
		} else {
			left = mid
			right = mid
			for right < len(nums) && nums[right] == 0 {
				right++
			}
			for left >= 0 && nums[left] == 0 {
				left--
			}
			left = left + 1
			right = len(nums) - right
			if left > right {
				return left
			}
			return right
		}
	}
	right = len(nums) - right - 1
	if left > right {
		return left
	}
	return right
}
