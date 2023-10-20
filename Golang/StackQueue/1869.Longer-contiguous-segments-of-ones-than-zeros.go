package main

func checkZeroOnes(s string) bool {
	longestOne, longestZero := 0, 0
	stack := []byte{}
	for i := 0; i < len(s); i++ {
		if len(stack) > 0 && stack[len(stack)-1] != s[i] {
			if stack[len(stack)-1] == '1' && len(stack) > longestOne {
				longestOne = len(stack)
			} else if stack[len(stack)-1] == '0' && len(stack) > longestZero {
				longestZero = len(stack)
			}
			stack = []byte{s[i]}
		} else {
			stack = append(stack, s[i])
		}
	}
	if len(stack) > 0 {
		if stack[len(stack)-1] == '1' && len(stack) > longestOne {
			longestOne = len(stack)
		} else if stack[len(stack)-1] == '0' && len(stack) > longestZero {
			longestZero = len(stack)
		}
	}
	return longestOne > longestZero
}
