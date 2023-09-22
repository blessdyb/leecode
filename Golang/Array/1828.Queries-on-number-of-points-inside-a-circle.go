package main

import "math"

func countPoints(points [][]int, queries [][]int) []int {
	ret := make([]int, len(queries))
	distance := func(a, b []int) float64 {
		return math.Sqrt(float64((a[0]-b[0])*(a[0]-b[0]) + (a[1]-b[1])*(a[1]-b[1])))
	}
	for i, query := range queries {
		temp := 0
		for _, point := range points {
			if distance(point, query[:2]) <= float64(query[2]) {
				temp++
			}
		}
		ret[i] = temp
	}
	return ret
}
