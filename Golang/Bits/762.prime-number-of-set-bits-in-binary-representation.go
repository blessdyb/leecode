package main

func countPrimeSetBits(left int, right int) int {
	hashmap := map[int]int{}
	hashmap[1] = 0
	ret := 0
	for i := left; i <= right; i++ {
		bits := 0
		for j := i; j > 0; {
			bits += j & 1
			j >>= 1
		}
		if _, ok := hashmap[bits]; !ok {
			flag := 1
			for k := 2; k < bits; k++ {
				if bits%k == 0 {
					flag = 0
					break
				}
			}
			hashmap[bits] = flag
		}
		ret += hashmap[bits]
	}
	return ret
}
