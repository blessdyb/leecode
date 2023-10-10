package main

func countBits(n int) []int {
	ret := make([]int, n+1)
	for i := 0; i <= n; i++ {
		ret[i] = ret[i<<1] + i&1
	}
	return ret
}
