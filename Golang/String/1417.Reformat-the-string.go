package main

import "math"

func reformat(s string) string {
	letters, digits := "", ""
	for i := 0; i < len(s); i++ {
		if s[i] >= '0' && s[i] <= '9' {
			digits += string(s[i])
		} else {
			letters += string(s[i])
		}
	}
	if int(math.Abs(float64(len(digits)-len(letters)))) <= 1 {
		temp := ""
		if len(digits) < len(letters) {
			digits, letters = letters, digits
		}
		i := 0
		for i < len(letters) {
			temp += string(digits[i]) + string(letters[i])
			i++
		}
		if i < len(digits) {
			temp += string(digits[i])
		}
		return temp
	}
	return ""
}
