package main

func canBeEqual(s1 string, s2 string) bool {
	if s1 == s2 {
		return true
	}
	b1, b2 := []byte(s1), []byte(s1)
	b1[0], b1[2] = b1[2], b1[0]
	if string(b1) == s2 {
		return true
	}
	b2[1], b2[3] = b2[3], b2[1]
	if string(b2) == s2 {
		return true
	}
	b2[0], b2[2] = b2[2], b2[0]
	if string(b2) == s2 {
		return true
	}
	return false
}
