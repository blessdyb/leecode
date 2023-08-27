package main

func sortArrayByParity(nums []int) []int {
	for i := 0; i < len(nums); i++ {
		if nums[i]%2 == 1 {
			for j := i + 1; j < len(nums); j++ {
				if nums[j]%2 == 0 {
					nums[i], nums[j] = nums[j], nums[i]
					break
				}
			}
		}
	}
	return nums
}
