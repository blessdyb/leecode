package main

func numberOfEmployeesWhoMetTarget(hours []int, target int) int {
	ret := 0
	for _, h := range hours {
		if h >= target {
			ret++
		}
	}
	return ret
}
