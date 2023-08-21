package main

func decompressRLElist(nums []int) []int {
	ret := []int{}
	for i := 0; i < len(nums); i += 2 {
		for j := 0; j < nums[i]; j++ {
			ret = append(ret, nums[i+1])
		}
	}
	return ret
}
