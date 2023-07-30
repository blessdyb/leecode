package main

func largestAltitude(gain []int) int {
	ret := 0
	prev := 0
	for _, g := range gain {
		current := prev + g
		if current > ret {
			ret = current
		}
		prev = current
	}
	return ret
}
