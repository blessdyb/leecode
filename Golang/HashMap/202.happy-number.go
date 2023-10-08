package main

func isHappy(n int) bool {
	hashmap := map[int]bool{}
	for n != 1 && !hashmap[n] {
		hashmap[n] = true
		temp := 0
		for n > 0 {
			digit := n % 10
			n = n / 10
			temp += digit * digit
		}
		n = temp
	}
	return n == 1
}
