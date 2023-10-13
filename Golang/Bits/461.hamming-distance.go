package main

func hammingDistance(x int, y int) int {
	getBits := func(num int) []int {
		bits := []int{}
		for num > 0 {
			bits = append(bits, num&1)
			num >>= 1
		}
		return bits
	}
	xBits, yBits := getBits(x), getBits(y)
	ret := 0
	for i, j := 0, 0; i < len(xBits) || j < len(yBits); {
		if i < len(xBits) && j < len(yBits) {
			if xBits[i] != yBits[j] {
				ret++
			}
			i++
			j++
		} else if i < len(xBits) && xBits[i] == 1 {
			i++
			ret++
		} else if y < len(yBits) && yBits[j] == 1 {
			j++
			ret++
		} else {
			i++
			j++
		}
	}
	return ret
}

func hammingDistanceTrick(x, y int) int {
	diff := x ^ y
	ret := 0
	for diff > 0 {
		ret += diff & 1
		diff >>= 1
	}
	return ret
}
