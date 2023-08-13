package main

func arithmeticTriplets(nums []int, diff int) int {
	hashmap := map[int]bool{}
	for _, num := range nums {
		hashmap[num] = true
	}
	ret := 0
	for _, num := range nums {
		if hashmap[num+diff] && hashmap[num+diff+diff] {
			ret++
		}
	}
	return ret
}
