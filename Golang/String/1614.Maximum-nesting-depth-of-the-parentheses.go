package main

func maxDepth(s string) int {
	stack := 0
	max := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			stack++
		} else if s[i] == ')' {
			stack--
		}
		if stack > max {
			max = stack
		}
	}
	return max
}
