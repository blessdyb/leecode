package main

func minOperations(s string) int {
	s0, s1 := 0, 0
	b0, b1 := []byte(s), []byte(s)
	for i := 0; i < len(s); i++ {
		if i%2 == 0 {
			if b0[i] == '1' {
				s0++
			}
			if b1[i] == '0' {
				s1++
			}
		} else {
			if b0[i] == '0' {
				s0++
			}
			if b1[i] == '1' {
				s1++
			}
		}
	}
	if s0 > s1 {
		return s1
	}
	return s0
}
