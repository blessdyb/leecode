package main

func nextGreaterElements(nums []int) []int {
	newNums := append(nums, nums[:len(nums)-1]...)
	ret := make([]int, len(newNums))
	stack := []int{}
	for i := 0; i < len(newNums); i++ {
		ret[i] = -1
		for len(stack) > 0 && newNums[stack[len(stack)-1]] < newNums[i] {
			ret[stack[len(stack)-1]] = newNums[i]
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}
	return ret[:len(nums)]
}
