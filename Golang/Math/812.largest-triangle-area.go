package main

import "math"

func largestTriangleArea(points [][]int) float64 {
	ret := 0.0
	// Area = 0.5 * |x1y2 + x2y3 + x3y1 - x2y1 - x3y2 - x1y3|
	n := len(points)
	for i := 0; i < n-2; i++ {
		for j := i + 1; j < n-1; j++ {
			for k := j + 1; k < n; k++ {
				x1, y1 := points[i][0], points[i][1]
				x2, y2 := points[j][0], points[j][1]
				x3, y3 := points[k][0], points[k][1]
				area := 0.5 * math.Abs(float64(x1*y2+x2*y3+x3*y1-x2*y1-x3*y2-x1*y3))
				if area > ret {
					ret = area
				}
			}
		}
	}
	return ret
}
