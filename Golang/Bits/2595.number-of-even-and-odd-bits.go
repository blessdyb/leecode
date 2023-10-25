package main

func evenOddBits(n int) []int {
	base := 1
	index := 0
	ret := []int{0, 0}
	for base < 1024 {
		if n&base == base {
			if index%2 == 0 {
				ret[0]++
			} else {
				ret[1]++
			}
		}
		index++
		base *= 2
	}
	return ret
}
