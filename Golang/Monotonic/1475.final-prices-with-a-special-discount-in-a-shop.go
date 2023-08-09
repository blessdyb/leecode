package main

func finalPrices(prices []int) []int {
	ret := append([]int{}, prices...)
	stack := []int{0}
	for i := 1; i < len(prices); i++ {
		for len(stack) > 0 && prices[stack[len(stack)-1]] >= prices[i] {
			ret[stack[len(stack)-1]] = prices[stack[len(stack)-1]] - prices[i]
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}
	return ret
}
