package main

func isPowerOfFour(n int) bool {
	if n == 1 {
		return true
	} else if n < 4 {
		return false
	}
	temp := 4
	for temp < n {
		temp *= 4
	}
	return temp == n
}

func isPowerOfFourBits(num uint) bool {
	// Check if the number is positive and a power of 2
	if num > 0 && (num&(num-1)) == 0 {
		// Check if the set bit is in an even position
		return (num & 0xAAAAAAAA) == 0
	}
	return false
}
