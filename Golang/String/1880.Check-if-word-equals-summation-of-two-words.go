package main

import (
	"strconv"
	"strings"
)

func isSumEqual(firstWord string, secondWord string, targetWord string) bool {
	getValue := func(str string) int {
		digits := []string{}
		for i := 0; i < len(str); i++ {
			digits = append(digits, strconv.Itoa(int(str[i]-'a')))
		}
		digitsStr := strings.Join(digits, "")
		ret := 0
		for i, j := len(digitsStr)-1, 1; i >= 0; {
			ret += int(digitsStr[i]-'0') * j
			i--
			j *= 10
		}
		return ret
	}
	return getValue(firstWord)+getValue(secondWord) == getValue(targetWord)
}


0, 0
0, 1

0, 0
1, 0


1, 0
0, 0

0, 1
0, 0