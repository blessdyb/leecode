package main

func maximum69Number(num int) int {
	sixIndex := 0
	count := 1
	originNum := num
	for originNum > 0 {
		if originNum%10 == 6 {
			sixIndex = count
		}
		originNum /= 10
		count *= 10
	}
	num += sixIndex * 3
	return num
}
