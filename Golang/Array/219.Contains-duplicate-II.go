package main

func containsNearbyDuplicate(nums []int, k int) bool {
	hashmap := map[int]int{}
	for index, num := range nums {
		if _, ok := hashmap[num]; !ok {
			hashmap[num] = index
		} else {
			if index-hashmap[num] <= k {
				return true
			}
			hashmap[num] = index
		}

	}
	return false
}
