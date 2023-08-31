package main

import "sort"

func findEvenNumbers(digits []int) []int {
	hashmap := map[int]bool{}
	n := len(digits)
	check := func(a, b, c int) {
		temp := digits[a]*100 + digits[b]*10 + digits[c]
		if temp >= 100 && temp%2 == 0 {
			hashmap[temp] = true
		}
	}
	for i := 0; i < n-2; i++ {
		for j := i + 1; j < n-1; j++ {
			for k := j + 1; k < n; k++ {
				check(i, j, k)
				check(i, k, j)
				check(j, k, i)
				check(j, i, k)
				check(k, i, j)
				check(k, j, i)
			}
		}
	}
	ret := []int{}
	for k, _ := range hashmap {
		ret = append(ret, k)
	}
	sort.Ints(ret)
	return ret
}
