package main

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	hashmap := map[int]int{}
	nums1map := map[int]bool{}
	for _, num := range nums1 {
		nums1map[num] = true
	}
	stack := []int{}
	for i := 0; i < len(nums2); i++ {
		hashmap[nums2[i]] = -1
		for len(stack) > 0 && nums2[stack[len(stack)-1]] < nums2[i] {
			hashmap[nums2[stack[len(stack)-1]]] = nums2[i]
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}
	ret := []int{}
	for _, num := range nums1 {
		ret = append(ret, hashmap[num])
	}
	return ret
}
