package main

func largestOddNumber(num string) string {
	index := -1
	for i := len(num) - 1; i >= 0; i-- {
		digit := int(num[i] - '0')
		if digit%2 == 1 {
			index = i
			break
		}
	}
	if index == -1 {
		return ""
	}
	return num[:index+1]
}
