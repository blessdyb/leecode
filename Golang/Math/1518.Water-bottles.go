package main

func numWaterBottles(numBottles int, numExchange int) int {
	ret := numBottles
	empty := 0
	for (numBottles + empty) < numExchange {
		newEmpty := numBottles + empty
		temp := newEmpty % numExchange
		numBottles = newEmpty / numExchange
		empty = temp
		ret += numBottles
	}
	return ret
}
