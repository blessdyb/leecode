package main

func countDistinctIntegers(nums []int) int {
	hashmap := map[int]bool{}
	for _, num := range nums {
		hashmap[num] = true
		sum := 0
		for num > 0 {
			sum = num%10 + sum*10
			num = num / 10
		}
		hashmap[sum] = true
	}
	return len(hashmap)
}
