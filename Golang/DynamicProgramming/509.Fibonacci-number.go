package main

func fib(n int) int {
	if n < 2 {
		return n
	}
	f0, f1, ret := 0, 1, 1
	for i := 2; i <= n; i++ {
		ret = f1 + f0
		f0 = f1
		f1 = ret
	}
	return ret
}
