package main

import "fmt"

func isPalindrome(x int) bool {
	str := fmt.Sprintf("%d", x)
	for i, j := 0, len(str)-1; i <= j; {
		if str[i] == str[j] {
			i++
			j--
		} else {
			return false
		}
	}
	return true
}
