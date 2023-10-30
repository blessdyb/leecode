package main

func maxAscendingSum(nums []int) int {
	ret := nums[0]
	temp := num[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			temp += nums[i]
		} else {
			if ret < temp {
				ret = temp
			}
			temp = nums[i]
		}
	}
	if ret < temp {
		ret = temp
	}
	return ret
}
