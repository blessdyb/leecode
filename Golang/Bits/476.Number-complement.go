package main

func findComplement(num int) int {
	ret := 0
	base := 1
	for num > 0 {
		if num&1 == 0 {
			ret += base
		}
		num >>= 1
		base <<= 1
	}
	return ret
}
