package main

func removeTrailingZeros(num string) string {
	index := len(num)
	for i := len(num) - 1; i >= 0; i-- {
		if num[i] != '0' {
			break
		}
		index = i
	}
	return num[:index]
}
