package main

func guess(num int) int

func guessNumber(n int) int {
	low, high := 1, n
	guessed := (high-low)/2 + low
	guessed_ret := guess(guessed)
	for guessed_ret != 0 {
		if guessed_ret > 0 {
			low = guessed + 1
		} else {
			high = guessed - 1
		}
		guessed = (high-low)/2 + low
		guessed_ret = guess(guessed)
	}
	return guessed
}
