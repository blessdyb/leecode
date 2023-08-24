package main

func distanceBetweenBusStops(distance []int, start int, destination int) int {
	sum := 0
	clockwise := 0
	if start > destination {
		start, destination = destination, start
	}
	for idx, dis := range distance {
		sum += dis
		if idx >= start && idx < destination {
			clockwise += dis
		}
	}
	if clockwise < sum-clockwise {
		return clockwise
	}
	return sum - clockwise
}
