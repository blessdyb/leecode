package main

func minCostToMoveChips(position []int) int {
	odds, evens := 0, 0
	for _, idx := range position {
		if idx%2 == 0 {
			evens++
		} else {
			odds++
		}
	}
	if odds > evens {
		return evens
	}
	return odds
}
