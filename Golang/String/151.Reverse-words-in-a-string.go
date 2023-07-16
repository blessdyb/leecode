package main

func reverseWords(s string) string {
	ret := ""
	start := -1
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			if start != -1 {
				ret = s[start:i] + " " + ret
				start = -1
			}
		} else if start == -1 {
			start = i
		}
	}
	if start != -1 {
		ret = s[start:] + " " + ret
	}
	return ret[:len(ret)-1]
}
