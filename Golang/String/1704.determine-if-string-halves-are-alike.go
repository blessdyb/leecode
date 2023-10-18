package main

func halvesAreAlike(s string) bool {
	countVowels := func(str string) int {
		ret := 0
		for i := 0; i < len(str); i++ {
			if str[i] == 'a' || str[i] == 'e' || str[i] == 'i' || str[i] == 'o' || str[i] == 'u' || str[i] == 'A' || str[i] == 'E' || str[i] == 'I' || str[i] == 'O' || str[i] == 'U' {
				ret++
			}
		}
		return ret
	}
	return countVowels(s[:len(s)/2]) == countVowels(s[len(s)/2:])
}
