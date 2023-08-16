package main

func countBalls(lowLimit int, highLimit int) int {
	hashmap := map[int]int{}
	for i := lowLimit; i <= highLimit; i++ {
		temp := i
		sum := 0
		for temp > 0 {
			sum += temp % 10
			temp = temp / 10
		}
		hashmap[sum]++
	}
	max := -1
	for _, v := range hashmap {
		if max < v {
			max = v
		}
	}
	return max
}
