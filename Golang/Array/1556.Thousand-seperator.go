package main

import "fmt"

func thousandSeparator(n int) string {
	if n == 0 {
		return "0"
	}
	nums := ""
	count := 0
	for n > 0 {
		digit := n % 10
		nums = fmt.Sprint(digit) + nums
		n = n / 10
		count++
		if count == 3 && n != 0 {
			nums = "." + nums
			count = 0
		}
	}
	return nums
}
