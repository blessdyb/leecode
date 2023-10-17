package main

func subtractProductAndSum(n int) int {
	sum, product := 0, 1
	for n > 0 {
		digit := n % 10
		n /= 10
		sum += digit
		product *= digit
	}
	return product - sum
}
