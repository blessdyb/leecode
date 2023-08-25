package main

func findNumbers(nums []int) int {
	ret := 0
	for _, num := range nums {
		count := 0
		for num > 0 {
			count += 1
			num /= 10
		}
		if count%2 == 0 {
			ret++
		}
	}
	return ret
}
