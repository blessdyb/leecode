package main

// Given the range 1 ~ 2^31 as a sorted array, we want to find sqrt(x). So binary search

func isPerfectSquare(num int) bool {
	i, j := 1, num
	for i <= j {
		h := i + (j-i)/2
		sqrt := h * h
		if sqrt == num {
			return true
		} else if sqrt < num {
			i = h + 1
		} else {
			j = h - 1
		}
	}
	return false
}
