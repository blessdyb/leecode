package main

func countDigits(num int) int {
	count := 0
	origin := num
	for num > 0 {
		digit := num % 10
		if origin%digit == 0 {
			count++
		}
		num /= 10
	}
	return count
}
