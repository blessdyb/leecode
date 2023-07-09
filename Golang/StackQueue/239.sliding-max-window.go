package main

func maxSlidingWindow(nums []int, k int) []int {
	queue := []int{}
	ret := []int{}
	for i := 0; i < len(nums); i++ {
		for len(queue) > 0 && nums[i] > nums[queue[len(queue)-1]] {
			queue = queue[:len(queue)-1]
		}
		queue = append(queue, i)
		for queue[0] <= i-k {
			queue = queue[1:]
		}
		for i >= k-1 {
			ret = append(ret, nums[queue[0]])
		}
	}
	return ret
}
