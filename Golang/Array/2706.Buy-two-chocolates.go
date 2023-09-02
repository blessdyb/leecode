package main

import "sort"

func buyChoco(prices []int, money int) int {
	sort.Ints(prices)
	bill := prices[0] + prices[1]
	if money >= bill {
		return money - bill
	}
	return money
}
