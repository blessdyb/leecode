package main

func reverseStr(s string, k int) string {
	ret := []byte(s)
	for i := 0; i < len(ret); i = i + 2*k {
		end := i + k
		if end >= len(ret) {
			end = len(ret)
		}
		for m, n := i, end-1; m < n; {
			s[m], s[n] = s[n], s[m]
			m++
			n--
		}
	}
	return string(ret)
}
