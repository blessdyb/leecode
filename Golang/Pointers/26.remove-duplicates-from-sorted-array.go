package main

// Since the problem asks for a O(1) in-place solution, two pointers is the only one in our toolkit can make it happen.
// When nums[i] == nums[j], we only need to move j forward
// Until nums[i] != num[j], we set nums[i + 1] = nums[j] and move both i, j forward
func removeDuplicates(nums []int) int {
	i, j := 0, 0
	length := len(nums)
	for j < length {
		if nums[i] != nums[j] {
			i++
			nums[i] = nums[j]
		}
		j++
	}
	return i + 1
}
