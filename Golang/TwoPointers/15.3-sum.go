package main

import "sort"

func threeSum(nums []int) [][]int {
	ret := [][]int{}
	sort.Ints(nums)
	length := len(nums)
	for i := 0; i < length; i++ {
		n1 := nums[i]
		if n1 > 0 {
			break
		} else if i > 0 && n1 == nums[i-1] {
			continue
		} else {
			left := i + 1
			right := length - 1
			for left < right {
				n2 := nums[left]
				n3 := nums[right]
				t := n1 + n2 + n3
				if t == 0 {
					ret = append(ret, []int{n1, n2, n3})
					for left < right && n2 == nums[left] {
						left++
					}
					for left < right && n3 == nums[right] {
						right--
					}
				} else if t > 0 {
					right--
				} else {
					left++
				}
			}
		}
	}
	return ret
}
