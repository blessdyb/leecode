package main

func sumIndicesWithKSetBits(nums []int, k int) int {
	count := 0
	for i := 0; i < len(nums); i++ {
		base := 1
		c := 0
		for base <= i {
			if i&base == base {
				c++
			}
			base *= 2
		}
		if c == k {
			count += nums[i]
		}
	}
	return count
}
