package main

import "sort"

func fourSum(nums []int, target int) [][]int {
	ret := [][]int{}
	sort.Ints(nums)
	length := len(nums)
	for i := 0; i < length-3; i++ {
		n1 := nums[i]
		if target >= 0 && n1 > target {
			break
		} else if i > 0 && n1 == nums[i-1] {
			continue
		} else {
			for j := i + 1; j < length-2; j++ {
				n2 := nums[j]
				if target >= 0 && n1+n2 > target {
					break
				} else if j > i+1 && n2 == nums[j-1] {
					continue
				} else {
					k, m := j+1, length-1
					for k < m {
						n3 := nums[k]
						n4 := nums[m]
						t := n1 + n2 + n3 + n4
						if t == target {
							ret = append(ret, []int{n1, n2, n3, n4})
							for k < m && n3 == nums[k] {
								k++
							}
							for k < m && n4 == nums[m] {
								m--
							}
						} else if t > target {
							m--
						} else {
							k++
						}
					}
				}
			}
		}
	}
	return ret
}
