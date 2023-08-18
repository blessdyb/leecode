package main

func isPalindrome(s string) bool {
	transform := []byte{}
	for i := 0; i < len(s); i++ {
		if 'a' <= s[i] && 'z' >= s[i] {
			transform = append(transform, s[i])
		} else if 'A' <= s[i] && 'Z' >= s[i] {
			transform = append(transform, s[i]+('a'-'A'))
		} else if '0' <= s[i] && '9' >= s[i] {
			transform = append(transform, s[i])
		}
	}
	for i, j := 0, len(transform)-1; i < j; {
		if transform[i] != transform[j] {
			return false
		}
		i++
		j--
	}
	return true
}
