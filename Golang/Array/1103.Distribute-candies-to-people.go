package main

func distributeCandies(candies int, num_people int) []int {
	ret := make([]int, num_people)
	index := 0
	candy := 1
	for candies > 0 {
		if candies >= candy {
			candies -= candy
			ret[index] += candy
		} else {
			ret[index] += candies
			candies = 0
		}
		candy++
		index++
		if index == num_people {
			index = 0
		}
	}
	return ret
}
