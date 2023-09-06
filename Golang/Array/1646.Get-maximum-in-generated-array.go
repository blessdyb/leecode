package main

func getMaximumGenerated(n int) int {
	if n <= 1 {
		return n
	}

	data := make([]int, n+1)
	data[0] = 0
	data[1] = 1
	max := 1
	for i := 1; i <= n; i++ {
		if i*2 <= n {
			data[i*2] = data[i]
			if max < data[i*2] {
				max = data[i*2]
			}
		}
		if i*2+1 <= n {
			data[i*2+1] = data[i] + data[i+1]
			if max < data[i*2+1] {
				max = data[i*2+1]
			}
		}
	}
	return max
}
