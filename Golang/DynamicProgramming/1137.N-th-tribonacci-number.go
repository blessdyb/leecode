package main

func tribonacci(n int) int {
	if n == 0 {
		return 0
	} else if n < 3 {
		return 1
	}
	t0, t1, t2, t3 := 0, 1, 1, 2
	for i := 3; i <= n; i++ {
		t3 = t2 + t1 + t0
		t2, t1, t0 = t3, t2, t1
	}
	return t3
}
