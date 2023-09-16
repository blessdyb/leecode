package main

func cellsInRange(s string) []string {
	ret := []string{}
	for i := s[0]; i <= s[3]; i++ {
		for j := s[1]; j <= s[4]; j++ {
			ret = append(ret, string([]byte{i, j}))
		}
	}
	return ret
}
