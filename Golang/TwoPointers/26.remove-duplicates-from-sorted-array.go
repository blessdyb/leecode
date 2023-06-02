/**
 * In place search and replace, two pointers might be a good solution to try
 */
func removeDuplicates(nums []int) int {
	i, j := 0, 0
	length = len(nums)
	for j < length {
		if nums[j] != nums[i] {
			// Here we want to keep the existing one, so move the slower pointer first and then override the value
			i++
			nums[i] = nums[j]
		}
		j++
	}
	return i
}