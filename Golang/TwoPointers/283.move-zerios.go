package main

// In-place array modification, so we need to consider two pointers solution.
// Remove 0 first, then reset the rest of elements to 0
func moveZeroes(nums []int) {
	// In-place remove 0
	i, j := 0, 0
	length := len(nums)
	for j < length {
		if nums[j] != 0 {
			nums[i] = nums[j]
			i++
		}
		j++
	}
	for i < length {
		nums[i] = 0
		i++
	}
}
