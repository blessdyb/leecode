package main

func averageValue(nums []int) int {
	ret := 0
	count := 0
	for _, num := range nums {
		if num%2 == 0 && num%3 == 0 {
			ret += num
			count++
		}
	}
	if count != 0 {
		return ret / count
	}
	return 0
}
