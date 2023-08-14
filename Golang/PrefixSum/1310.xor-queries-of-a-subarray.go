package main

func xorQueries(arr []int, queries [][]int) []int {
	presum := []int{arr[0]}
	for i := 1; i < len(arr); i++ {
		presum = append(presum, presum[i-1]^arr[i])
	}
	ret := []int{}
	for _, query := range queries {
		ret = append(ret, presum[query[0]]^presum[query[1]]^arr[query[0]])
	}
	return ret
}
