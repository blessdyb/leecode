package main

import (
	"math"
	"sort"
)

func minMovesToSeat(seats []int, students []int) int {
	sort.Ints(seats)
	sort.Ints(students)
	ret := 0
	for i := 0; i < len(seats); i++ {
		ret += int(math.Abs(float64(seats[i] - students[i])))
	}
	return ret
}
