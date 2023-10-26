package main

import "sort"

func arrayRankTransform(arr []int) []int {
	sortedArr := append([]int{}, arr...)
	sort.Ints(sortedArr)
	rank := 1
	rankMap := map[int]int{}
	for i := 0; i < len(sortedArr); i++ {
		if _, ok := rankMap[sortedArr[i]]; !ok {
			rankMap[sortedArr[i]] = rank
			rank++
		}
	}
	for i := 0; i < len(arr); i++ {
		arr[i] = rankMap[arr[i]]
	}
	return arr
}
