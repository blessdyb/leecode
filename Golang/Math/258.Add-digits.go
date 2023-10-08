package main

func addDigits(num int) int {
	for num >= 10 {
		temp := 0
		for num > 0 {
			temp += num % 10
			num = num / 10
		}
		num = temp
	}
	return num
}
