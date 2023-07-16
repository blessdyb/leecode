package main

func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
	ret := 0
	hashmap := map[int]int{}
	for _, n1 := range nums1 {
		for _, n2 := range nums2 {
			hashmap[-n1-n2]++
		}
	}
	for _, n3 := range nums3 {
		for _, n4 := range nums4 {
			ret += hashmap[n3+n4]
		}
	}
	return ret
}
