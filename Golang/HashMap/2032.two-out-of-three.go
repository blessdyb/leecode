package main

func twoOutOfThree(nums1 []int, nums2 []int, nums3 []int) []int {
	hashset := map[int]bool{}
	hashmap := map[int]int{}
	ret := []int{}
	for _, nums := range [][]int{nums1, nums2, nums3} {
		temp := map[int]bool{}
		for _, num := range nums {
			if !temp[num] {
				temp[num] = true
				hashmap[num]++
				if hashmap[num] >= 2 {
					hashset[num] = true
				}
			}
		}
	}
	for key := range hashset {
		ret = append(ret, key)
	}
	return ret
}
