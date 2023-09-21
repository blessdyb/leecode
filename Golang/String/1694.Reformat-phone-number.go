package main

import "strings"

func reformatNumber(number string) string {
	digits := []byte{}
	for i := 0; i < len(number); i++ {
		if number[i] >= '0' && number[i] <= '9' {
			digits = append(digits, number[i])
		}
	}
	blocks := []string{}
	for len(digits) > 4 {
		blocks = append(blocks, string(digits[:3]))
		digits = digits[3:]
	}
	if len(digits) == 4 {
		blocks = append(blocks, string(digits[:2]), string(digits[2:]))
	} else {
		blocks = append(blocks, string(digits))
	}
	return strings.Join(blocks, "-")
}
