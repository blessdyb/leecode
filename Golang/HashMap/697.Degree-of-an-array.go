package main

func findShortestSubArray(nums []int) int {
	maxFrequency := 0
	h1, h2, h3 := map[int]int{}, map[int]int{}, map[int]int{}
	for i, num := range nums {
		h1[num]++
		if _, ok := h2[num]; !ok {
			h2[num] = i
		}
		h3[num] = i
		if maxFrequency < h1[num] {
			maxFrequency = h1[num]
		}
	}
	ret := len(nums)
	for k, v := range h1 {
		if v == maxFrequency {
			if ret > h3[k]-h2[k]+1 {
				ret = h3[k] - h2[k] + 1
			}
		}
	}
	return ret
}
