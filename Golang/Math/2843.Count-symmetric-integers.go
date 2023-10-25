package main

func countSymmetricIntegers(low int, high int) int {
	ret := 0
	if high <= 99 {
		start := 10
		if low > 10 {
			start = low
		}
		for i := start; i <= high; i++ {
			if i/10 == i%10 {
				ret++
			}
		}
	} else if high <= 10000 {
		start := 10
		if low > 10 {
			start = low
		}
		for i := start; i <= 99; i++ {
			if i/10 == i%10 {
				ret++
			}
		}
		start = 1000
		if low > 1000 {
			start = low
		}
		end := 9999
		if high < 9999 {
			end = high
		}
		for i := start; i <= end; i++ {
			thousand := i / 1000
			hundred := i/100 - thousand*10
			ten := i/10 - i/100*10
			one := i % 10
			if thousand+hundred == ten+one {
				ret++
			}
		}
	}
	return ret
}
