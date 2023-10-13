package main

import "strings"

func repeatedSubstringPattern(s string) bool {
	length := len(s)
	for i := 1; i < length/2+1; i++ {
		if length%i == 0 {
			partial := s[:i]
			temp := ""
			for len(temp) < length {
				temp += partial
			}
			if temp == s {
				return true
			}
		}
	}
	return false
}

func repeatedSubstringPatternTrick(s string) bool {
	double := s + s
	return strings.Contains(double[1:len(double)-1], s)
}
