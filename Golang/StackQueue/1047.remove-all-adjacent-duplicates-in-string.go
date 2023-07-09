package main

func removeDuplicates(s string) string {
	stack := []byte{s[0]}
	for i := 1; i < len(s); i++ {
		length := len(stack)
		if length != 0 && stack[length-1] == s[i] {
			stack = stack[:length-1]
		} else {
			stack = append(stack, s[i])
		}
	}
	return string(stack)
}
