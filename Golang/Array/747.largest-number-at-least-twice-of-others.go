package main

func dominantIndex(nums []int) int {
	max := -1
	ret := -1
	for index, num := range nums {
		if num > max {
			max = num
			ret = index
		}
	}
	for _, num := range nums {
		if num != 0 && num < max && max/num < 2 {
			return -1
		}
	}
	return ret
}
