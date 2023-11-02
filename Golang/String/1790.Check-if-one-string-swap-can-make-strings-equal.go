package main

func areAlmostEqual(s1 string, s2 string) bool {
	idx := []int{}
	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			idx = append(idx, i)
		}
		if len(idx) > 2 {
			return false
		}
	}
	if len(idx) == 0 {
		return true
	} else if len(idx) == 1 {
		return false
	}
	if s1[idx[0]] == s2[idx[1]] && s1[idx[1]] == s2[idx[0]] {
		return true
	}
	return false
}
