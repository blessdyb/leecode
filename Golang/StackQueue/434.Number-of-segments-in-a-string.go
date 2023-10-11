package main

func countSegments(s string) int {
	ret := 0
	stack := []byte{}
	for i := 0; i < len(s); i++ {
		if s[i] != ' ' {
			stack = append(stack, s[i])
		} else if len(stack) > 0 {
			ret++
			stack = []byte{}
		}
	}
	if len(stack) > 0 {
		ret++
	}
	return ret
}
