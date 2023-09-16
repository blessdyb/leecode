package main

func maxCount(m int, n int, ops [][]int) int {
	if len(ops) == 0 {
		return m * n
	}
	a, b := ops[0][0], ops[0][1]
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	for _, op := range ops {
		a, b = min(a, op[0]), min(b, op[1])
	}
	return a * b
}
