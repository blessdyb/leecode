package main

import "sort"

func findHSMapSolution(nums []int) int {
	hashmap := map[int]int{}
	for _, num := range nums {
		hashmap[num]++
	}
	ret := 0
	for num, freq := range hashmap {
		if v, ok := hashmap[num-1]; ok {
			if ret < v+freq {
				ret = v + freq
			}
		}
	}
	return ret
}

func findHSTwoPointsSolution(nums []int) int {
	sort.Ints(nums)
	ret := 0
	start, end := 0, 0
	for end < len(nums) {
		diff := nums[end] - nums[start]
		if diff == 1 {
			temp := end - start + 1
			if ret < temp {
				ret = temp
			}
			end++
		} else if diff > 1 {
			start++
		} else {
			end++
		}
	}
	return ret
}
