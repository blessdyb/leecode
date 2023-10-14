package main

func checkPerfectNumber(num int) bool {
	total := 0
	for i := 1; i < num; i++ {
		if num%i == 0 {
			total += i
		}
	}
	return total == num
}
