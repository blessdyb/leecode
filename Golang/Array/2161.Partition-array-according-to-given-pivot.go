package main

func pivotArray(nums []int, pivot int) []int {
	a, b, c := []int{}, []int{}, []int{}
	for _, num := range nums {
		if num < pivot {
			a = append(a, num)
		} else if num == pivot {
			b = append(b, num)
		} else {
			c = append(c, num)
		}
	}
	return append(a, append(b, c...)...)
}
