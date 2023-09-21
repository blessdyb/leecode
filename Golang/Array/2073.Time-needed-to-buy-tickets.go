package main

func timeRequiredToBuy(tickets []int, k int) int {
	ret := 0
	for tickets[k] != 0 {
		for i := 0; i < len(tickets); i++ {
			if tickets[i] != 0 {
				tickets[i]--
				ret++
			}
			if tickets[k] == 0 {
				break
			}
		}
	}
	return ret
}
