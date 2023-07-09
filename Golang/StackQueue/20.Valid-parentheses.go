package main

func isValid(s string) bool {
	parenthesMap := map[byte]byte{
		')': '(',
		'}': '{',
		']': '[',
	}
	stack := []byte{s[0]}
	for i := 1; i < len(s); i++ {
		if len(stack) == 0 {
			stack = append(stack, s[i])
		} else if sV, ok := parenthesMap[s[i]]; ok && sV == stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, s[i])
		}
	}
	return len(stack) == 0
}
