package main

import "fmt"

func reverseBits(num uint32) uint32 {
	var ret uint32 = 0
	numStr := fmt.Sprintf("%b", num)
	for len(numStr) < 32 {
		numStr = "0" + numStr
	}
	for i := 31; i >= 0; i-- {
		if numStr[i] == '1' {
			ret = 2*ret + 1
		} else {
			ret = 2 * ret
		}
	}
	return ret
}

func reverseBitsBinary(num uint32) uint32 {
	var ret uint32 = 0
	for i := 0; i < 32; i++ {
		ret <<= 1         // Shift the result to the left by 1
		result |= num & 1 // Set the rightmost bit of the result to the current bit of num
		num >>= 1         // Shift the num to the right by 1
	}
	return ret
}
