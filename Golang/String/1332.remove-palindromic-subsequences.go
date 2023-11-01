package main

func removePalindromeSub(s string) int {
	if s == "" {
		return 0
	}
	for i, j := 0, len(s)-1; i < j; {
		if s[i] != s[j] {
			// If the given string is not a palindrome, you can remove all the 'a's and 'b's from it separately. This is because all 'a's can be removed in one step, and all 'b's can be removed in one step. In the end, you'll have an empty string.
			return 2
		}
		i++
		j--
	}
	return 1
}
