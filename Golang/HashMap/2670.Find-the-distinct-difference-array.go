package main

func distinctDifferenceArray(nums []int) []int {
	hashmap1 := map[int]int{}
	hashmap2 := map[int]int{}
	for _, num := range nums {
		hashmap2[num]++
	}
	ret := make([]int, len(nums))
	for id, num := range nums {
		hashmap1[num]++
		hashmap2[num]--
		if hashmap2[num] == 0 {
			delete(hashmap2, num)
		}
		ret[id] = len(hashmap1) - len(hashmap2)
	}
	return ret
}
