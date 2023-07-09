package main

import "strconv"

func evalRPN(tokens []string) int {
	stringToInt := func(s string) int {
		value, _ := strconv.Atoi(s)
		return value
	}
	if len(tokens) < 3 {
		return stringToInt(tokens[0])
	}
	stack := []int{stringToInt(tokens[0]), stringToInt(tokens[1])}
	result := 0
	for i := 2; i < len(tokens); i++ {
		switch tokens[i] {
		case "+":
			result = stack[len(stack)-1] + stack[len(stack)-2]
			stack = stack[:len(stack)-2]
		case "-":
			result = stack[len(stack)-1] - stack[len(stack)-2]
			stack = stack[:len(stack)-2]
		case "*":
			result = stack[len(stack)-1] * stack[len(stack)-2]
			stack = stack[:len(stack)-2]
		case "/":
			result = stack[len(stack)-2] / stack[len(stack)-1]
			stack = stack[:len(stack)-2]
		default:
			result = stringToInt(tokens[i])
		}
		stack = append(stack, result)
	}
	return result
}
