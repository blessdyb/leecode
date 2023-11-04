package main

func maxArea(height []int) int {
	ret := 0
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	for l, r := 0, len(height)-1; l < r; {
		area := min(height[l], height[r]) * (r - l)
		if area > ret {
			ret = area
		}

		// Move the pointer pointing to the shorter bar inward, as moving the pointer pointing to the taller bar will not increase the capacity, and it will only decrease the width.
		if height[l] < height[r] {
			l++
		} else {
			r--
		}
	}
	return ret
}
