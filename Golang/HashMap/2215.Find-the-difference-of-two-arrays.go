package main

func findDifference(nums1 []int, nums2 []int) [][]int {
	hashmap1 := map[int]bool{}
	hashmap2 := map[int]bool{}
	for _, num := range nums1 {
		hashmap1[num] = true
	}
	for _, num := range nums2 {
		hashmap2[num] = true
	}
	hashmap3 := map[int]bool{}
	hashmap4 := map[int]bool{}
	for _, num := range nums1 {
		if !hashmap2[num] {
			hashmap3[num] = true
		}
	}
	for _, num := range nums2 {
		if !hashmap1[num] {
			hashmap4[num] = true
		}
	}
	a1, a2 := []int{}, []int{}
	for k, _ := range hashmap3 {
		a1 = append(a1, k)
	}
	for k, _ := range hashmap4 {
		a2 = append(a2, k)
	}
	return [][]int{a1, a2}
}
