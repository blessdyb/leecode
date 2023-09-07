package main

import (
	"fmt"
	"strconv"
)

func divisorSubstrings(num int, k int) int {
	ret := 0
	numStr := fmt.Sprint(num)
	for i := 0; i < len(numStr); i++ {
		temp := 0
		for j := i; j < i+k && i+k < len(numStr); j++ {
			m, _ := strconv.Atoi(numStr[j : j+1])
			temp = temp*10 + m
		}
		if temp > 0 && num%temp == 0 {
			ret++
		}
	}
	return ret
}
