package main

func subarraySumBruteForce(nums []int, k int) int {
	ret := 0
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j <= len(nums); j++ {
			sum := 0
			for m := i; m < j; m++ {
				sum += nums[m]
			}
			if sum == k {
				ret++
			}
		}
	}
	return ret
}

func subarraySumPrefixSum1(nums []int, k int) int {
	ret := 0
	preSum := make([]int, len(nums)+1)
	preSum[0] = 0
	// preSum[i] is the sum of items in range [0, i)
	for i := 1; i <= len(nums); i++ {
		preSum[i] = preSum[i-1] + nums[i-1]
	}
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j <= len(nums); j++ {
			// The sum of items in range [i, j)
			if preSum[j]-preSum[i] == k {
				ret++
			}
		}
	}
	return ret
}

func subarraySumPrefixSum2(nums []int, k int) int {
	ret, sum := 0, 0
	hashmap := map[int]int{0: 1}
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		ret += hashmap[sum-k]
		hashmap[sum]++
	}
	return ret
}
