package main

import "fmt"

func digitSum(s string, k int) string {
	if len(s) <= k {
		return s
	}
	nums := []byte{}
	ret := ""
	for i := 0; i < len(s); i++ {
		if len(nums) < k {
			nums = append(nums, s[i])
		} else if len(nums) == k {
			sum := 0
			for j := 0; j < k; j++ {
				sum += int(nums[j] - '0')
			}
			ret += fmt.Sprint(sum)
		}
	}
	if len(nums) > 0 {
		sum := 0
		for i := 0; i < len(nums); i++ {
			sum += int(nums[i] - '0')
		}
		ret += fmt.Sprint(sum)
	}
	return digitSum(ret, k)
}
