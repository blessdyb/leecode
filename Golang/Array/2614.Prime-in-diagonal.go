package main

import "sort"

func diagonalPrime(nums [][]int) int {
	diagnals := []int{}
	for i, j := 0, len(nums)-1; i < len(nums) && j >= 0; {
		diagnals = append(diagnals, nums[i][i])
		diagnals = append(diagnals, nums[i][j])
		i++
		j--
	}
	sort.Ints(diagnals)
	isPrime := func(n int) bool {
		for i := 2; i < n; i++ {
			if n%i == 0 {
				return false
			}
		}
		return n > 1
	}
	for i := len(diagnals) - 1; i >= 0; i-- {
		if isPrime(diagnals[i]) {
			return diagnals[i]
		}
	}
	return 0
}
