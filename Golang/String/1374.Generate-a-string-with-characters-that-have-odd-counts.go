package main

func generateTheString(n int) string {
	ret := ""
	if n%2 == 1 {
		ret = "c"
		n = n - 1
	}
	if n > 0 {
		ret += "a"
		if n > 1 {
			for i := 1; i <= n-1; i++ {
				ret += "b"
			}
		}
	}
	return ret
}
