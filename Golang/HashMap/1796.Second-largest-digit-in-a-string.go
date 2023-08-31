package main

func secondHighest(s string) int {
	digits := make([]int, 10)
	for i := 0; i < len(s); i++ {
		if s[i] >= '0' && s[i] <= '9' {
			digits[s[i]-'0']++
		}
	}
	count := 0
	ret := -1
	for i := 9; i >= 0; i-- {
		if digits[i] > 0 {
			count++
		}
		if count == 2 {
			return i
		}
	}
	return ret
}
