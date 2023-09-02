package main

func maximumPopulation(logs [][]int) int {
	years := make([]int, 101)
	max := 0
	year := 1950
	for _, log := range logs {
		for i := log[0]; i < log[1]; i++ {
			years[i-1950]++
		}
	}
	for i, y := range years {
		if y > max {
			max = y
			year = i + 1950
		}
	}
	return year
}
