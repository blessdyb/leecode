/** In place replacement, two pointers
 */
func moveZeroes(nums []int)  {
	i, j := 0, 0
	length := len(nums)
	for j < length {
		if nums[j] != 0 {
			// Here we want to remove the current item if it's 0, so override first and then move the slower pointer
			nums[i] = nums[j]
			i++
		}
		j++
	}
	for i < length {
		nums[i] = 0
	}
}