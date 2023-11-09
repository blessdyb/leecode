package main

func numberOfArithmeticSlices(nums []int) int {
	if len(nums) < 3 {
		return 0
	}
	ret, diff := 0, 0
	for i := 0; i < len(nums)-1; i++ {
		diff = nums[i+1] - nums[i]
		for j := i + 2; j < len(nums); j++ {
			if diff == nums[j]-nums[j-1] {
				ret++
			} else {
				break
			}
		}
	}
	return ret
}
