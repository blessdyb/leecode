package main

import "strconv"

func calPoints(operations []string) int {
	ret := 0
	records := []int{}
	for i := 0; i < len(operations); i++ {
		if operations[i] == "D" {
			records = append(records, 2*records[len(records)-1])
			ret += records[len(records)-1]
		} else if operations[i] == "C" {
			ret -= records[len(records)-1]
			records = records[:len(records)-1]
		} else if operations[i] == "+" {
			score := records[len(records)-1] + records[len(records)-2]
			records = append(records, score)
			ret += records[len(records)-1]
		} else {
			score, _ := strconv.Atoi(operations[i])
			records = append(records, score)
			ret += records[len(records)-1]
		}
	}
	return ret
}
