package main

func isPowerOfTwoRecursion(n int) bool {
	if n < 1 {
		return false
	}
	return n == 1 || (n%2 == 0 && isPowerOfTwoRecursion(n/2))
}

// Given any n, the value of n & (n-1) will set the bit of the left most 1 in n (binary) to 0
func isPowerOfTwoMath(n int) bool {
	if n < 1 {
		return false
	}
	return n&(n-1) == 0
}
