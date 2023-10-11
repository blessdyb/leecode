package main

func arrangeCoins(n int) int {
	i, j := 1, 1
	for j*(j+1) <= n*2 {
		i = j
		j++
	}
	return i
}
