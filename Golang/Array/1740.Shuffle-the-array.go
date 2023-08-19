package main

func shuffle(nums []int, n int) []int {
	ret := []int{}
	for i, j := 0, n; i < n && j < 2*n; {
		ret = append(ret, nums[i])
		ret = append(ret, nums[j])
		i++
		j++
	}
	return ret
}
