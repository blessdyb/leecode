package main

func findMaxAverage(nums []int, k int) float64 {
	preSum := make([]int, len(nums))
	preSum[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		preSum[i] = preSum[i-1] + nums[i]
	}
	max := -100000 * 10000
	for i := 0; i < len(nums)-k+1; i++ {
		if max < preSum[i+k-1]-preSum[i]+nums[i] {
			max = preSum[i+k-1] - preSum[i] + nums[i]
		}
	}
	return float64(max) / float64(k)
}
