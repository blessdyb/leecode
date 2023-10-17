package main

import "strings"

func capitalizeTitle(title string) string {
	toUpper := func(c byte) byte {
		if c >= 'a' && c <= 'z' {
			return c - ('a' - 'A')
		}
		return c
	}
	toLower := func(c byte) byte {
		if c >= 'A' && c <= 'Z' {
			return c + ('a' - 'A')
		}
		return c
	}
	words := strings.Split(title, " ")
	ret := []string{}
	for _, word := range words {
		if len(word) > 2 {
			temp := []byte(word)
			temp[0] = toUpper(word[0])
			for i := 1; i < len(word); i++ {
				temp[i] = toLower(word[i])
			}
			ret = append(ret, string(temp))
		} else {
			temp := []byte(word)
			for i := 0; i < len(word); i++ {
				temp[i] = toLower(word[i])
			}
			ret = append(ret, string(temp))
		}
	}
	return strings.Join(ret, " ")
}
