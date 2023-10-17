package main

func getNoZeroIntegers(n int) []int {
	isValid := func(num int) bool {
		for num > 0 {
			if num%10 == 0 {
				return false
			}
			num /= 10
		}
		return true
	}
	var ret []int
	for i := 1; i < n; i++ {
		if isValid(i) && isValid(n-i) {
			ret = []int{i, n - i}
			break
		}
	}
	return ret
}
