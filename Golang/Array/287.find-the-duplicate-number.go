package main

import "math"

func findDuplicateNavigate(nums []int) int {
	duplicate := 0
	for i := 0; i < len(nums); i++ {
		duplicate = int(math.Abs(float64(nums[i])))
		if nums[duplicate] < 0 {
			break
		}
		nums[duplicate] *= -1
	}
	for i := 0; i < len(nums); i++ {
		nums[i] = int(math.Abs(float64(nums[i])))
	}
	return duplicate
}

func findDuplicateFloyd(nums []int) int {
	slow, fast := nums[0], nums[0]
	for {
		slow = nums[slow]
		fast = nums[nums[fast]]
		if slow == fast {
			break
		}
	}
	slow = nums[0]
	for slow != fast {
		slow = nums[slow]
		fast = nums[fast]
	}
	return fast
}

func findDuplicateRecursive(nums []int) int {
	var recursive func(current int) int
	recursive = func(current int) int {
		if current == nums[current] {
			return current
		}
		next := nums[current]
		nums[current] = current
		return recursive(next)
	}
	return recursive(0)
}

func findDuplicateArrayMap(nums []int) int {
	for nums[0] != nums[nums[0]] {
		nums[0], nums[nums[0]] = nums[nums[0]], nums[0]
	}
	return nums[0]
}
