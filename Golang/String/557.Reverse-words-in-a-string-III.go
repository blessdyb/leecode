package main

import "strings"

func reverseWords(s string) string {
	words := strings.Split(s, " ")
	ret := []string{}
	for _, word := range words {
		temp := []byte(word)
		for i, j := 0, len(temp)-1; i < j; {
			temp[i], temp[j] = temp[j], temp[i]
			i++
			j--
		}
		ret = append(ret, string(temp))
	}
	return strings.Join(ret, " ")
}
