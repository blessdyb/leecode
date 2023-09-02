package main

func countAsterisks(s string) int {
	flag := true
	count := 0
	for i := 0; i < len(s); i++ {
		if flag && s[i] == '*' {
			count++
		} else if flag && s[i] == '|' {
			flag = false
		} else if !flag && s[i] == '|' {
			flag = true
		}
	}
	return count

}
