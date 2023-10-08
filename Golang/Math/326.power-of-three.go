package main

func isPoswerOfThree(n int) bool {
	if n < 1 {
		return false
	}
	for n%3 == 0 && n >= 3 {
		n /= 3
	}
	return n == 1
}

func isPoswerOfThreeNoRecursion(n int) bool {
	// 1162261467 is pow(3, 19)
	return n > 0 && 1162261467%n == 0
}
