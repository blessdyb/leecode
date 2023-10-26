package main

func countTime(time string) int {
	h := 1
	m := 1
	if time[0] == '?' && time[1] == '?' {
		h = 24
	} else if time[0] == '?' && time[1] != '?' {
		if time[1] <= '3' {
			h = 3
		} else {
			h = 2
		}
	} else if time[0] != '?' && time[1] == '?' {
		if time[0] < '2' {
			h = 10
		} else {
			h = 4
		}
	}
	if time[3] == '?' && time[4] == '?' {
		m = 60
	} else if time[3] == '?' && time[4] != '?' {
		m = 6
	} else if time[3] != '?' && time[4] == '?' {
		m = 10
	}

	return h * m
}
