package main

func countSeniors(details []string) int {
	ret := 0
	for _, detail := range details {
		if detail[11] > '6' || (detail[11] == '6' && detail[12] > '0') {
			ret++
		}
	}
	return ret
}
