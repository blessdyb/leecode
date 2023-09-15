package main

func sortArrayByParityII(nums []int) []int {
	ret := make([]int, len(nums))
	even, odd := 0, 1
	for i := 0; i < len(nums); i++ {
		if nums[i]%2 == 0 {
			ret[even] = nums[i]
			even += 2
		} else {
			ret[odd] = nums[i]
			odd += 2
		}
	}
	return ret
}
