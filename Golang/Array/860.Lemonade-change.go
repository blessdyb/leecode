package main

func lemonadeChange(bills []int) bool {
	changes := [3]int{0, 0, 0}
	for _, bill := range bills {
		if bill == 20 {
			if changes[1] > 0 && changes[0] > 0 {
				changes[1]--
				changes[0]--
			} else if changes[0] >= 3 {
				changes[0] -= 3
			} else {
				return false
			}
			changes[2]++
		} else if bill == 10 {
			if changes[0] > 0 {
				changes[0]--
			} else {
				return false
			}
			changes[1]++
		} else {
			changes[0]++
		}
	}
	return true
}
