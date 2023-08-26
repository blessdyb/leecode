package main

import "strconv"

func maximumValue(strs []string) int {
	getValue := func(str string) int {
		for i := 0; i < len(str); i++ {
			if str[i] >= 'a' && str[i] <= 'z' {
				return len(str)
			}
		}
		ret, _ := strconv.Atoi(str)
		return ret
	}
	max := getValue(strs[0])
	for i := 1; i < len(strs); i++ {
		current := getValue(strs[i])
		if max < current {
			max = current
		}
	}
	return max
}
