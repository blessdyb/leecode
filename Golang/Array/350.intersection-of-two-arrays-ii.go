package main

func intersection(nums1 []int, nums2 []int) []int {
	hashmap := map[int]int{}
	for _, num := range nums1 {
		hashmap[num]++
	}
	result := []int{}
	for _, num := range nums2 {
		if hashmap[num] > 0 {
			hashmap[num]--
			result = append(result, num)
		}
	}
	return result
}

// If nums1 and nums2 are sorted in ascending order, we can use two pointers.
// If the length of nums1 is smaller than the length of nums2, we can use nums1 to construct the hashmap.
// If we can't load the whole nums2 into the memory at once, we can read elements from nums2 in smaller chunks and then use two pointers approach to find the intersection with the sorted nums1
