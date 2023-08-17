package main

func countLargestGroup(n int) int {
	hashmap := map[int]int{}
	max := 0
	ret := 0
	for i := 1; i <= n; i++ {
		sum := 0
		temp := i
		for temp > 0 {
			sum += temp % 10
			temp = temp / 10
		}
		hashmap[sum]++
		if hashmap[sum] > max {
			max = hashmap[sum]
			ret = 1
		} else if hashmap[sum] == max {
			ret++
		}
	}
	return ret
}
