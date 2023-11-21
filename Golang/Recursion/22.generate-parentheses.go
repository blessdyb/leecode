package main

func generateParenthesis(n int) []string {
	ret := []string{}
	var helper func(s string, left, right int)
	helper = func(s string, left, right int) {
		if len(s) == 2*n {
			ret = append(ret, s)
		} else {
			if left < n {
				helper(s+"(", left+1, right)
			}
			if right < left {
				// If the count of closed parentheses right is less than the count of open parentheses left, we can add a closing parenthesis to the combination and call the backtrack function recursively with updated parameters.
				helper(s+")", left, right+1)
			}
		}
	}
	helper("", 0, 0)
	return ret
}
