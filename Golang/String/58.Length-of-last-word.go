package main

func lengthOfLastWord(s string) int {
	ret := 0
	flag := false
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == ' ' && flag {
			break
		} else if s[i] != ' ' {
			flag = true
			ret++
		}
	}
	return ret
}
