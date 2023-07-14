package main

func isHappy(n int) bool {
	cache := map[int]bool{}
	helper := func(num int) int {
		sum := 0
		for num > 0 {
			t := num % 10
			sum += t * t
			num = num / 10
		}
		return sum
	}
	for n != 1 && !cache[n] {
		n, cache[n] = helper(n), true
	}
	return n == 1
}
