package main

import "math"

func fairCandySwap(aliceSizes []int, bobSizes []int) []int {
	sum := func(sizes []int) (int, map[int]bool) {
		t := 0
		hashmap := map[int]bool{}
		for _, size := range sizes {
			t += size
			hashmap[size] = true
		}
		return t, hashmap
	}
	aTotal, _ := sum(aliceSizes)
	bTotal, bMap := sum(bobSizes)
	expect := (aTotal + bTotal) / 2
	var ret []int
	for _, item := range aliceSizes {
		fromBob := int(math.Abs(float64(expect - (aTotal - item))))
		if bMap[fromBob] && bTotal-fromBob+item == expect {
			ret = []int{item, fromBob}
			break
		}
	}
	return ret
}
