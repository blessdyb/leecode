package main

func checkOnesSegment(s string) bool {
	startOne := false
	startZero := false
	for i := 0; i < len(s); i++ {
		if s[i] == '1' {
			if startZero {
				return false
			}
			if !startOne {
				startOne = true
			}
		} else if startOne {
			startZero = true
		}
	}
	return true
}
