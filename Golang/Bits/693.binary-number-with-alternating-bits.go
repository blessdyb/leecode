package main

func hasAlternatingBits(n int) bool {
	prev := -1
	for n > 0 {
		last := n & 1
		if prev == last {
			return false
		}
		prev = last
		n >>= 1
	}
	return true
}
