package main

func totalMoney(n int) int {
	monday := 1
	ret := 0
	count := 0
	prev := monday - 1
	for i := 1; i <= n; i++ {
		ret += prev + 1
		prev += 1
		count++
		if count == 7 {
			monday++
			prev = monday - 1
			count = 0
		}
	}
	return ret
}
