package main

import "fmt"

func readBinaryWatch(num int) []string {
	countBits := func(num int) int {
		bits := 0
		for num > 0 {
			bits += num & 1
			num = num >> 1
		}
		return bits
	}
	ret := []string{}
	for i := 0; i < 12; i++ {
		for j := 0; j < 60; j++ {
			if countBits(i)+countBits(j) == num {
				ret = append(ret, fmt.Sprintf("%d:%02d", i, j))
			}
		}
	}
	return ret
}
