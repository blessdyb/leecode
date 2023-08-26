package main

func arraySign(nums []int) int {
	negative := 0
	for _, num := range nums {
		if num == 0 {
			return 0
		} else if num < 0 {
			negative++
		}
	}
	if negative%2 == 0 {
		return 1
	}
	return -1
}
