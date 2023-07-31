package main

func pivotInteger(n int) int {
	preSum := make([]int, n+1)
	for i := 1; i <= n; i++ {
		preSum[i] = preSum[i-1] + i
	}
	for i := 1; i <= n; i++ {
		if preSum[i] == preSum[n]-preSum[i-1] {
			return i
		}
	}
	return -1
}
