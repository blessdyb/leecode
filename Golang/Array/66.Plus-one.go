package main

func plusOne(digits []int) []int {
	pre := 1
	for i := len(digits) - 1; i >= 0; i-- {
		digits[i] += pre
		if digits[i] >= 10 {
			digits[i] = 0
			pre = 1
		} else {
			pre = 0
		}
	}
	if pre == 1 {
		digits = append([]int{1}, digits...)
	}
	return digits
}
