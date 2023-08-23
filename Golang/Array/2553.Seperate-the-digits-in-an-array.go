package main

func separateDigits(nums []int) []int {
	ret := []int{}
	for _, num := range nums {
		temp := []int{}
		for num > 0 {
			temp = append([]int{num % 10}, temp...)
			num /= 10
		}
		ret = append(ret, temp...)
	}
	return ret
}
