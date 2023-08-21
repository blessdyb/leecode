package main

import "math"

func differenceOfSum(nums []int) int {
	x, y := 0, 0
	for _, num := range nums {
		x += num
		for num > 0 {
			y += num % 10
			num = num / 10
		}
	}
	return int(math.Abs(float64(x - y)))
}
