package main

func unequalTriplets(nums []int) int {
	ret := 0
	hashmap := map[int]int{}
	for _, num := range nums {
		hashmap[num]++
	}
	if len(hashmap) >= 3 {
		values := []int{}
		for _, v := range hashmap {
			values = append(values, v)
		}
		for i := 0; i < len(values)-2; i++ {
			for j := i + 1; j < len(values)-1; j++ {
				for k := j + 1; k < len(values); k++ {
					ret += values[i] * values[j] * values[k]
				}
			}
		}
	}
	return ret
}
