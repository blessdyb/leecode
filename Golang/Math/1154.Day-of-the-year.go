package main

import "strconv"

func dayOfYear(date string) int {
	isLeapYear := func(year int) bool {
		if year%4 == 0 {
			if year%100 == 0 && year%400 != 0 {
				return false
			}
			return true
		}
		return false
	}
	monthDays := map[int]int{
		1:  31,
		2:  28,
		3:  31,
		4:  30,
		5:  31,
		6:  30,
		7:  31,
		8:  31,
		9:  30,
		10: 31,
		11: 30,
		12: 31,
	}
	year, _ := strconv.Atoi(date[:4])
	month, _ := strconv.Atoi(date[5:7])
	day, _ := strconv.Atoi(date[8:10])
	ret := day
	for i := 1; i < month; i++ {
		if i == 2 && isLeapYear(year) {
			ret += 29
		} else {
			ret += monthDays[i]
		}
	}
	return ret
}
