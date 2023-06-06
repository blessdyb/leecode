package main

// Given 0 ~ 2^31 - 1 range sorted array, to find a target sqrt(x), the only solution is binary search
// To find the nearst value, we need to remember the smaller value before the next binary search round in case there's no exact match
func mySqrt(x int) int {
	i, j := 0, x
	ret := 0
	for i <= j {
		h := i + (j-i)/2
		sqrt := h * h
		if sqrt == x {
			return h
		} else if sqrt < x {
			ret = h
			i = h + 1
		} else {
			j = h - 1
		}
	}
	return ret
}
