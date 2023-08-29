package main

func mostFrequentEven(nums []int) int {
	hashmap := map[int]int{}
	ret, max := -1, 0
	for _, num := range nums {
		if num%2 == 0 {
			hashmap[num]++
		}
	}
	for key, value := range hashmap {
		if max < value {
			max = value
			ret = key
		} else if max == value && ret > key {
			ret = key
		}
	}
	return ret
}
