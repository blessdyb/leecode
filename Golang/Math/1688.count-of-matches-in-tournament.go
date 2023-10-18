package math

func numberOfMatches(n int) int {
	ret := 0
	for n > 1 {
		ret += n / 2
		carry := 0
		if n%2 == 1 {
			carry = 1
		}
		n = n/2 + carry
	}
	return ret
}
