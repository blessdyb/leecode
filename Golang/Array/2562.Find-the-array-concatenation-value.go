package main

func findTheArrayConcVal(nums []int) int64 {
	ret := 0
	for i, j := 0, len(nums)-1; i <= j; {
		if i == j {
			ret += nums[i]
		} else {
			temp := nums[j]
			for temp > 0 {
				nums[i] *= 10
				temp /= 10
			}
			ret += nums[i] + nums[j]
		}
		i++
		j--
	}
	return int64(ret)
}
