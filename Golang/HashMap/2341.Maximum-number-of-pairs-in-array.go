package main

func numberOfPairs(nums []int) []int {
	hashmap := map[int]int{}
	for _, num := range nums {
		hashmap[num]++
	}
	frequencies := []int{}
	for _, v := range hashmap {
		frequencies = append(frequencies, v)
	}
	pairs := 0
	left := 0
	for i := 0; i < len(frequencies); {
		if frequencies[i] >= 2 {
			pairs++
			frequencies[i] -= 2
		} else if frequencies[i] == 1 {
			left++
			i++
		} else {
			i++
		}
	}
	return []int{pairs, left}
}
