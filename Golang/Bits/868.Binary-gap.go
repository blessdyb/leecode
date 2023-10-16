package main

func binaryGap(n int) int {
	ret := 0
	lastOneIndex := -1
	index := 0
	for n > 0 {
		if n&1 == 1 {
			if lastOneIndex != -1 {
				diff := index - lastOneIndex
				if diff > ret {
					ret = diff
				}
			}
			lastOneIndex = index
		}
		index++
		n >>>= 1
	}
	return ret
}