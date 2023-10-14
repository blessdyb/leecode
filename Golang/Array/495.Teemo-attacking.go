package main

func findPoisonedDuration(timeSeries []int, duration int) int {
	ret := 0
	lastStop := 0
	for _, t := range timeSeries {
		end := t + duration - 1
		if lastStop == 0 {
			ret += duration
		} else {
			if lastStop < t {
				ret += duration
			} else {
				ret += duration - (lastStop - t + 1)
			}
		}
		lastStop = end
	}
	return ret
}
