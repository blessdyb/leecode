package main

func intersection(nums1 []int, nums2 []int) []int {
	hashmap := map[int]bool{}
	for _, num := range nums1 {
		hashmap[num] = true
	}
	retmap := map[int]bool{}
	for _, num := range nums2 {
		if _, ok := hashmap[num]; ok {
			retmap[num] = true
		}
	}
	ret := []int{}
	for k := range retmap {
		ret = append(ret, k)
	}
	return ret
}
