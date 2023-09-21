package main

import "strconv"

func areNumbersAscending(s string) bool {
	last := -1
	nums := []byte{}
	for i := 0; i < len(s); i++ {
		if s[i] >= '0' && s[i] <= '9' {
			nums = append(nums, s[i])
		} else if len(nums) > 0 {
			temp, _ := strconv.Atoi(nums)
			if temp > last {
				last = temp
			} else {
				return false
			}
			nums = []byte{}
		}
	}
	if len(nums) > 0 {
		temp, _ := strconv.Atoi(nums)
		if temp <= last {
			return false
		}
	}
	return true
}
