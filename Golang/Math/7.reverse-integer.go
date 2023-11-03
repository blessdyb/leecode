package main

func reverse(x int) int {
	ret := 0
	minValue := -1 << 31  //-2147483648
	maxValue := 1<<31 - 1 // 2147483647
	for x != 0 {
		temp := x % 10
		if ret > maxValue/10 || (ret == maxValue/10 && temp > 7) {
			return 0
		} else if ret < minValue/10 || (ret == minValue/10 && temp < -8) {
			return 0
		}
		ret = ret*10 + temp
	}
	return ret
}
