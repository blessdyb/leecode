package main

func checkString(s string) bool {
	lastA, firstB := -1, 101
	for i := 0; i < len(s); i++ {
		if s[i] == 'b' && firstB == 101 {
			firstB = i
		} else if s[i] == 'a' {
			lastA = i
		}
	}
	return lastA < firstB
}
