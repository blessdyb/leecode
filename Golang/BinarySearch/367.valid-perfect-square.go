package main

/** Given any integer, it's a sorted digits 1 ~ n and we want to find sqrt(x) which is a perfect case for the binary search
 */
func isPerfectSquare(num int) bool {
	i, j := 0, num
	for i <= j {
		h = i + (j-i)/2
		h2 = h * h
		if h2 == num {
			return true
		} else if h2 < num {
			i = h + 1
		} else {
			j = h - 1
		}
	}
	return false
}
