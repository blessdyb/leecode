package main

import "math"

func minTimeToVisitAllPoints(points [][]int) int {
	ret := 0
	var distance func(a, b, c, d int) int
	distance = func(a, b, c, d int) int {
		if a == c {
			return int(math.Abs(float64(d - b)))
		} else if b == d {
			return int(math.Abs(float64(a - c)))
		}
		ratio := float64(d-b) / float64(c-a)
		if ratio == float64(1) || ratio == float64(-1) {
			return int(math.Abs(float64(c - a)))
		} else if c-a > 0 && d-b > 0 {
			return 1 + distance(a+1, b+1, c, d)
		} else if c-a > 0 && d-b < 0 {
			return 1 + distance(a+1, b-1, c, d)
		} else if c-a < 0 && d-b < 0 {
			return 1 + distance(a-1, b-1, c, d)
		}
		return 1 + distance(a-1, b+1, c, d)
	}
	for i := 0; i < len(points); i++ {
		ret += distance(points[i-1][0], points[i-1][1], points[i][0], points[i][1])
	}
	return ret
}
