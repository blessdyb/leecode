package main

import (
	"sort"
	"strconv"
)

func largestInteger(num int) int {
	odd, even, origin := []int{}, []int{}, []int{}
	nums := strconv.Itoa(num)
	length := len(nums)
	for i := 0; i < length; i++ {
		digit := int(nums[i] - '0')
		if digit%2 == 0 {
			even = append(even, digit)
		} else {
			odd = append(odd, digit)
		}
		origin = append(origin, digit)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(odd)))
	sort.Sort(sort.Reverse(sort.IntSlice(even)))
	ret := 0
	for _, digit := range origin {
		if digit%2 == 0 {
			ret = ret*10 + even[0]
			even = even[1:]
		} else {
			ret = ret*10 + odd[0]
			odd = odd[1:]
		}
	}
	return ret
}
