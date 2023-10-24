package main

import "sort"

func splitNum(num int) int {
	digits := []int{}
	for num > 0 {
		digits = append(digits, num%10)
		num /= 10
	}
	sort.Ints(digits)
	base := 10
	a, b := 0, 0
	for i := 0; i < len(digits); {
		a = digits[i] + a*base
		i++
		if i < len(digits) {
			b = digits[i] + b*base
			i++
		}
	}
	return a + b
}
