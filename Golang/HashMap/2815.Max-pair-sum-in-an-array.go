package main

import "sort"

func maxSum(nums []int) int {
	maxDigit := func(num int) int {
		max := 0
		for num > 0 {
			digit := num % 10
			if digit > max {
				max = digit
			}
			num = num / 10
		}
		return max
	}
	hashmap := map[int][]int{}
	for _, num := range nums {
		md := maxDigit(num)
		if _, ok := hashmap[md]; !ok {
			hashmap[md] = []int{}
		}
		hashmap[md] = append(hashmap[md], num)
	}
	ret := -1
	for _, v := range hashmap {
		if len(v) > 1 {
			sort.Ints(v)
			temp := v[len(v)-1] + v[len(v)-2]
			if temp > ret {
				ret = temp
			}
		}
	}
	return ret
}
