package main

func numberOfPoints(nums [][]int) int {
	data := make([]int, 101)
	for _, value := range nums {
		for i := value[0]; i <= value[1]; i++ {
			data[i] = 1
		}
	}
	count := 0
	for _, v := range data {
		if v == 1 {
			count++
		}
	}
	return count
}
