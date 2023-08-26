package main

func calculateTax(brackets [][]int, income int) float64 {
	if income < brackets[0][0] {
		return float64(income) * float64(brackets[0][1]) / 100
	}
	var ret float64 = float64(brackets[0][0]) * float64(brackets[0][1]) / 100
	income -= brackets[0][0]
	for i := 1; i < len(brackets); i++ {
		start := brackets[i][0] - brackets[i-1][0]
		if income < start {
			ret += float64(income) * float64(brackets[i][1]) / 100
			break
		}
		ret += float64(start) * float64(brackets[i][1]) / 100
		income -= start
	}
	return ret
}
