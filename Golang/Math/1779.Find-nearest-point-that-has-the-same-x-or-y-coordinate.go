package main

func nearestValidPoint(x int, y int, points [][]int) int {
	ret := -1
	distance := 100000000
	for index, point := range points {
		if point[0] == x || point[1] == y {
			temp := (x-point[0])*(x-point[0]) + (y-point[1])*(y-point[1])
			if distance > temp {
				distance = temp
				ret = index
			}
		}
	}
	return ret
}
