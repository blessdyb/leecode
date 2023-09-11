package main

func countHillValley(nums []int) int {
	ret := 0
	stack := []int{nums[0]}
	for i := 1; i < len(nums); i++ {
		if nums[i] == stack[len(stack)-1] {
			continue
		} else if len(stack) < 2 {
			stack = append(stack, nums[i])
		} else if (stack[1]-stack[0] > 0 && stack[1]-nums[i] > 0) || (stack[1]-stack[0] < 0 && stack[1]-nums[i] < 0) {
			ret++
			stack = append(stack, nums[i])
			stack = stack[1:]
		} else {
			stack = append(stack, nums[i])
			stack = stack[1:]
		}
	}
	return ret
}
