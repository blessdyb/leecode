package main

func checkRecord(s string) bool {
	absent := 0
	late := 0
	for i := 0; i < len(s); i++ {
		if s[i] == 'A' {
			absent++
		} else if s[i] == 'L' {
			if i+2 < len(s) {
				if s[i+1] == 'L' && s[i+2] == 'L' {
					late++
				}
			}
		}
	}
	return absent < 2 && late == 0
}
