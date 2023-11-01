package main

import "strings"

func findOcurrences(text string, first string, second string) []string {
	words := strings.Split(text, " ")
	ret := []string{}
	for i := 2; i < len(words); i++ {
		if words[i-2] == first && words[i-1] == second {
			ret = append(ret, words[i])
		}
	}
	return ret
}
