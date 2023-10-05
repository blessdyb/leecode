package main

import "fmt"

func summaryRanges(nums []int) []string {
	stack := [][]int{}
	temp := []int{}
	for i := 0; i < len(nums); i++ {
		if len(temp) == 0 {
			temp = append(temp, nums[i])
		} else if temp[len(temp)-1]+1 == nums[i] {
			temp = append(temp, nums[i])
		} else {
			stack = append(stack, append([]int{}, temp...))
			temp = []int{nums[i]}
		}
	}
	if len(temp) > 0 {
		stack = append(stack, temp)
	}
	ret := []string{}
	for i := 0; i < len(stack); i++ {
		if len(stack[i]) == 1 {
			ret = append(ret, fmt.Sprint(stack[i][0]))
		} else {
			ret = append(ret, fmt.Sprintf("%d->%d", stack[i][0], stack[i][len(stack[i])-1]))
		}
	}
	return ret
}
