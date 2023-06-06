package main

// The problem is asking to get a subarray, so we could try sliding window
func minSubArrayLen(target int, nums []int) int {
	ret := 0
	length := len(nums)
	i, j := 0, 0
	sum := 0
	for j < length {
		sum += nums[j]
		for sum >= target {
			tempLength := j - i + 1
			if tempLength < ret || ret == 0 {
				ret = tempLength
			}
			sum -= nums[i]
			i++
		}
		j++
	}
	return ret
}
