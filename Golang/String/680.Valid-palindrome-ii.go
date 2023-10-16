package main

func validPalindrome(s string) bool {
	var recursive func(str string, flag bool) bool
	recursive = func(str string, flag bool) bool {
		if len(str) <= 1 {
			return true
		}
		if str[0] == str[len(str)-1] {
			return recursive(str[1:len(str)-1], flag)
		} else if flag {
			return recursive(str[0:len(str)-1], false) || recursive(str[1:len(str)], false)
		}
		return false
	}
	return recursive(s, true)
}
