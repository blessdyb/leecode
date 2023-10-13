package main

func bitWiseComplement(num int) int {
	if n == 0 {
		return 1
	}
	ret := 0
	base := 1
	for num > 0 {
		if num&1 == 0 {
			ret += base
		}
		base <<= 1
		num >>= 1
	}
	return ret
}
