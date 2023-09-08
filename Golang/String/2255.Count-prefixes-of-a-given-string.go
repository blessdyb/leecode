package main

import "strings"

func countPrefixes(words []string, s string) int {
	ret := 0
	for _, word := range words {
		if strings.HasPrefix(s, word) {
			ret++
		}
	}
	return ret
}
