package main

// Integer overflow error
func func getLucky(s string, k int) int {
	num := 0
	for i := 0; i < len(s); i++ {
		digit := int(s[i] - 'a') + 1
		if digit < 10 {
			num += digit
		} else {
			num += digit % 10
			num += digit / 10
		}
	}
	for k > 1 && num >= 10 {
		temp := 0
		for num > 0 {
			temp += num % 10
			num /= 10
		}
		num = temp
		k--
	}
	return num
}