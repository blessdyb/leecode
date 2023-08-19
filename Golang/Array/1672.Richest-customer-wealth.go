package main

func maximumWealth(accounts [][]int) int {
	max := 0
	for _, acc := range accounts {
		temp := 0
		for _, a := range acc {
			temp += a
		}
		if max < temp {
			max = temp
		}
	}
	return max
}
