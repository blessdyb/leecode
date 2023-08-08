package main

func dailyTemperatures(temperatures []int) []int {
	stack := []int{0}
	ret := make([]int, len(temperatures))
	for i := 1; i < len(temperatures); i++ {
		if temperatures[stack[len(stack)-1]] >= temperatures[i] {
			stack = append(stack, i)
		} else {
			for len(stack) > 0 && temperatures[stack[len(stack)-1]] < temperatures[i] {
				ret[stack[len(stack)-1]] = i - stack[len(stack)-1]
				stack = stack[:len(stack)-1]
			}
			stack = append(stack, i)
		}
	}
	return ret
}
