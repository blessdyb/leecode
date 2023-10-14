package main

import "math"

func constructRectangle(area int) []int {
	sqrt := int(math.Sqrt(float64(area)))
	var ret []int
	if sqrt*sqrt == area {
		ret = []int{sqrt, sqrt}
	} else {
		for i := sqrt + 1; i <= area; i++ {
			if area%i == 0 {
				ret = []int{i, area / i}
				break
			}
		}
	}
	return ret
}
